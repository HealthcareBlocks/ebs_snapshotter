// Package ebs implements functionality for managing Amazon EBS snapshots
//
// See https://github.com/HealthcareBlocks/ebs_snapshotter for sample usage in a command line app
//
// Note: this package relies on the AWS SDK, thus the host environment should
// either have an associated IAM role or user with the following IAM permissions:
// * ec2:CopySnapshot
// * ec2:CreateSnapshot
// * ec2:CreateTags
// * ec2:DeleteSnapshot
// * ec2:DeleteTags
// * ec2:DescribeSnapshotAttribute
// * ec2:DescribeSnapshots
// * ec2:DescribeTags
// * ec2:DescribeVolumes
// * ec2:ModifySnapshotAttribute
// * ec2:ResetSnapshotAttribute
package ebs

import (
	"fmt"
	"sort"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/healthcareblocks/ebs_snapshotter/awserror"
)

// MaxRetries is the number of AWS service requests that can be retried
const MaxRetries = 10

// SnapshotManager manages EBS snapshots via the AWS SDK
type SnapshotManager struct {
	// The EC2 region containg the EBS volumes to snapshot. This parameter is required.
	// See http://docs.aws.amazon.com/general/latest/gr/rande.html#ec2_region
	Region string

	// An optional endpoint URL (hostname only or fully qualified URI)
	// that overrides the default generated endpoint for a client. Set this
	// to `""` to use the default generated endpoint.
	Endpoint string

	// Whether EBS volume tags should be copied to the EBS snapshots
	CopyVolumeTags bool

	// How many EBS snapshots to retain after the current snapshot is generated
	NumSnapshotsToRetain int

	// Internal reference to EC2 Client
	ec2 *ec2.EC2
}

// NewSnapshotManager returns a new SnapshotManager pointer. The input parameters are based
// on the SnapshotManager struct fields.
//
//     mgr := ebs.NewSnapshotManager("us-west-2", "", true, 5, false)
//
func NewSnapshotManager(region string, endpoint string, copyVolumeTags bool, numSnapshotsToRetain int, debug bool) *SnapshotManager {
	if region == "" {
		log.Fatal("region is required")
	}

	if numSnapshotsToRetain < 1 {
		log.Fatal("numSnapshotsToRetain should be great than 0")
	}

	config := aws.NewConfig().WithRegion(region).WithMaxRetries(MaxRetries)
	if endpoint != "" {
		config = config.WithEndpoint(endpoint)
	}

	if debug {
		config = config.WithLogLevel(aws.LogDebugWithHTTPBody)
	}

	return &SnapshotManager{
		Region:               region,
		Endpoint:             endpoint,
		CopyVolumeTags:       copyVolumeTags,
		NumSnapshotsToRetain: numSnapshotsToRetain,
		ec2:                  ec2.New(session.New(), config),
	}
}

// SnapshotVolumes is a helper method that wraps several operations. It queries for attached
// EBS volumes in the SnapshotManager's region, generates new snapshots, optionally
// copying any volume tags, and keeps the last X snapshots as specified by retainCount.
func (mgr *SnapshotManager) SnapshotVolumes() (snapshotsGenerated int) {
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

	resp, err := mgr.ec2.DescribeVolumes(params)
	awserror.HandleError(err)

	for _, volume := range resp.Volumes {
		mgr.CreateSnapshot(volume)
		mgr.DestroySnapshots(volume)
		snapshotsGenerated++
	}

	return snapshotsGenerated
}

// CreateSnapshot creates an EBS snapshot for a specific EBS volume, optionally copying
// any volume tags depending on SnapshotManager's CopyVolumeTags. If the volume has a Name tag,
// it is used as the snapshot's description.
func (mgr *SnapshotManager) CreateSnapshot(volume *ec2.Volume) {
	description := fmt.Sprintf("Snapshot for volume %s", *volume.VolumeId)
	for _, tag := range volume.Tags {
		if *tag.Key == "Name" {
			description = *tag.Value
		}
	}

	params := &ec2.CreateSnapshotInput{
		Description: aws.String(description),
		VolumeId:    aws.String(*volume.VolumeId),
	}

	log.Printf("Starting snapshot for %s in region %s", *volume.VolumeId, mgr.Region)

	snapshot, err := mgr.ec2.CreateSnapshot(params)
	awserror.HandleError(err)

	if mgr.CopyVolumeTags && len(volume.Tags) > 0 {
		mgr.TagResource(snapshot.SnapshotId, volume.Tags)
	}
}

// TagResource creates tags for an EBS snapshot
func (mgr *SnapshotManager) TagResource(id *string, tags []*ec2.Tag) {
	// remove tags containing "aws:", as these are reserved by AWS
	// and cannot be duplicated for other resources
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

	_, err := mgr.ec2.CreateTags(params)
	awserror.HandleError(err)
}

// DestroySnapshots deletes snapshots greater than SnapshotManager's NumSnapshotsToRetain for a given volume
func (mgr *SnapshotManager) DestroySnapshots(volume *ec2.Volume) (snapshotsDestroyed int) {
	if mgr.NumSnapshotsToRetain < 1 {
		log.Fatal("NumSnapshotsToRetain should be great than 0")
	}

	params := &ec2.DescribeSnapshotsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("volume-id"),
				Values: []*string{
					aws.String(*volume.VolumeId),
				},
			},
		},
	}
	resp, err := mgr.ec2.DescribeSnapshots(params)
	awserror.HandleError(err)

	snapshots := resp.Snapshots
	sort.Sort(ByStartTime(snapshots))

	numberSnapshotsToDelete := len(snapshots) - mgr.NumSnapshotsToRetain
	for numberSnapshotsToDelete > 0 {
		params := &ec2.DeleteSnapshotInput{
			SnapshotId: aws.String(*resp.Snapshots[numberSnapshotsToDelete-1].SnapshotId),
		}

		_, err := mgr.ec2.DeleteSnapshot(params)
		if err != nil && err.(awserr.Error).Code() != "InvalidSnapshot.InUse" {
			awserror.HandleError(err)
		}

		numberSnapshotsToDelete--
		snapshotsDestroyed++
	}

	return snapshotsDestroyed
}
