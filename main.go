/*
* EBS Snapshotter creates snapshots from existing EBS volumes via the AWS SDK
*
* Configurable options:
* - regions
* - number of snapshots to retain
* - copying of volume-level tags
* - SNS alert when snapshots complete
*
* AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION should be set prior to execution.
 */

package main // import "github.com/atlashealth/ebs_snapshotter"

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sns"
)

// MaxRetries is the number of AWS service requests that can be retried
const MaxRetries = 10

var (
	regions     = flag.String("regions", "us-west-2", "AWS regions to include in EBS snapshots")
	copyTags    = flag.Bool("copytags", true, "Copy tags from volume")
	retainCount = flag.Int("retain", 7, "Keep x number of snapshots per each volume")
	snsRegion   = flag.String("sns_region", "us-west-2", "AWS region for SNS topic")
	snsTopic    = flag.String("sns_topic", "", "SNS ARN topic (optional)")
)

// ByStartTime implements sort.Interface for ec2.Snapshot based on the StartTime field
type ByStartTime []*ec2.Snapshot

func (t ByStartTime) Len() int           { return len(t) }
func (t ByStartTime) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ByStartTime) Less(i, j int) bool { return t[i].StartTime.Before(*t[j].StartTime) }

// ec2Service returns an EC2 client for use with API calls
func ec2Service(region string) *ec2.EC2 {
	return ec2.New(&aws.Config{
		MaxRetries: aws.Int(MaxRetries),
		Region:     aws.String(region),
	})
}

// snapshotVolumes retrieves attached volumes for a specific region, generates
// new snapshots, and deletes any snapshots outside of the retention period
func snapshotVolumes(region string) {
	svc := ec2Service(region)

	params := &ec2.DescribeVolumesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("attachment.status"),
				Values: []*string{
					aws.String("attached"),
				},
			},
		},
	}

	resp, err := svc.DescribeVolumes(params)
	if err != nil {
		handleResponseError(err)
	}

	for _, volume := range resp.Volumes {
		createSnapshot(region, volume)
		destroyOldSnapshots(region, volume)
	}

	sendCompletionAlert(region)
}

// createSnapshot creates a snapshot for a specific volume;
// if the volume has a Name tag, it is used as the snapshot's description
func createSnapshot(region string, volume *ec2.Volume) {
	svc := ec2Service(region)

	description := fmt.Sprintf("Snapshot for volume %s", *volume.VolumeID)
	for _, tag := range volume.Tags {
		if *tag.Key == "Name" {
			description = *tag.Value
		}
	}

	params := &ec2.CreateSnapshotInput{
		Description: aws.String(description),
		VolumeID:    aws.String(*volume.VolumeID),
	}

	fmt.Println("Starting snapshot for " + *volume.VolumeID + " in region " + region)
	snapshot, err := svc.CreateSnapshot(params)
	if err != nil {
		handleResponseError(err)
	}

	if *copyTags && len(volume.Tags) > 0 {
		tagResource(region, snapshot.SnapshotID, volume.Tags)
	}
}

// tagResource creates tags for an ec2 resource - in this case, an EBS snapshot
func tagResource(region string, id *string, tags []*ec2.Tag) {
	svc := ec2Service(region)

	// remove tags containing "aws:", as these are reserved by AWS
	// and cannot be duplicated to other resources
	filteredTags := tags[:0]
	for _, tag := range tags {
		if !strings.Contains(*tag.Key, "aws:") {
			filteredTags = append(filteredTags, tag)
		}
	}

	params := &ec2.CreateTagsInput{
		Resources: []*string{
			aws.String(*id),
		},
		Tags: filteredTags,
	}

	_, err := svc.CreateTags(params)
	if err != nil {
		handleResponseError(err)
	}
}

// destroyOldSnapshots deletes snapshots greater than retainCount for a given volume
func destroyOldSnapshots(region string, volume *ec2.Volume) {
	svc := ec2Service(region)

	params := &ec2.DescribeSnapshotsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("volume-id"),
				Values: []*string{
					aws.String(*volume.VolumeID),
				},
			},
		},
	}
	resp, err := svc.DescribeSnapshots(params)
	if err != nil {
		handleResponseError(err)
	}

	snapshots := resp.Snapshots
	sort.Sort(ByStartTime(snapshots))

	numberSnapshotsToDelete := len(snapshots) - *retainCount
	for numberSnapshotsToDelete > 0 {
		params := &ec2.DeleteSnapshotInput{
			SnapshotID: aws.String(*resp.Snapshots[numberSnapshotsToDelete-1].SnapshotID),
		}

		_, err := svc.DeleteSnapshot(params)
		if err != nil && err.(awserr.Error).Code() != "InvalidSnapshot.InUse" {
			handleResponseError(err)
		}

		numberSnapshotsToDelete--
	}
}

// sendCompletionAlert sends an SNS alert if snsTopic flag is set
func sendCompletionAlert(region string) {
	if *snsTopic == "" {
		return
	}

	svc := sns.New(&aws.Config{
		Region: aws.String(*snsRegion),
	})

	params := &sns.PublishInput{
		Subject:  aws.String(fmt.Sprintf("EBS Snapshots Completed (%s)", region)),
		Message:  aws.String("Completed on " + time.Now().Format(time.RFC822)),
		TopicARN: aws.String(*snsTopic),
	}

	_, err := svc.Publish(params)
	if err != nil {
		handleResponseError(err)
	}
}

// handleResponseError is a helper method for displaying AWS errors
func handleResponseError(err error) {
	if awsErr, ok := err.(awserr.Error); ok {
		fmt.Println("Code: " + awsErr.Code())
		fmt.Println("Message: " + awsErr.Message())

		if awsErr.OrigErr() != nil {
			fmt.Print("Orginal Error: ")
			fmt.Println(awsErr.OrigErr())
		}

		if reqErr, ok := err.(awserr.RequestFailure); ok {
			fmt.Printf("Status Code: %d\n", reqErr.StatusCode())

			if reqErr.RequestID() != "" {
				fmt.Println("Request ID: " + reqErr.RequestID())
			}
		}
	} else {
		fmt.Println(err.Error())
	}
	os.Exit(1)
}

func main() {
	flag.Parse()
	var wg sync.WaitGroup

	fmt.Println("Starting Snapshot Process On " + time.Now().Format(time.RFC822))

	for _, region := range strings.Split(*regions, ",") {
		wg.Add(1)
		go func(region string) {
			defer wg.Done()
			snapshotVolumes(region)
		}(region)
	}

	wg.Wait()
}
