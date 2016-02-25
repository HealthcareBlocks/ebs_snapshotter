// EBS Snapshotter creates snapshots from existing EBS volumes via the AWS SDK.
// It can generate snapshots for more than one region, retain a specific number of
// historical snapshots, copy tags from volumes, and send an optional SNS alert.
//
// Source: https://github.com/HealthcareBlocks/ebs_snapshotter
//
// Note: this package relies on the AWS SDK, thus the host environment should
// either have an associated IAM role or user with the following IAM permissions:
// - ec2:CopySnapshot
// - ec2:CreateSnapshot
// - ec2:CreateTags
// - ec2:DeleteSnapshot
// - ec2:DeleteTags
// - ec2:DescribeSnapshotAttribute
// - ec2:DescribeSnapshots
// - ec2:DescribeTags
// - ec2:DescribeVolumes
// - ec2:ModifySnapshotAttribute
// - ec2:ResetSnapshotAttribute
// - SNS:Publish (optional)

package main // import "github.com/healthcareblocks/ebs_snapshotter"

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/healthcareblocks/ebs_snapshotter/ebs"
	"github.com/healthcareblocks/ebs_snapshotter/sns"
	"github.com/healthcareblocks/ec2_metrics_publisher/metadata"
)

var (
	appVersion = flag.Bool("v", false, "Prints version of this app and exits")
	debug      = flag.Bool("d", false, "Turns on AWS request profiling")

	// EBS snapshot flags
	regions     = flag.String("regions", "", "AWS EC2 regions (comma delimited) to include in EBS snapshots. If not set, this value is determined using the host machine's EC2 metadata.")
	copyTags    = flag.Bool("copytags", true, "Copy tags from volume")
	retainCount = flag.Int("retain", 7, "Keep x number of snapshots per each volume")

	// SNS alert flags
	snsTopic   = flag.String("sns_topic", "", "Optional SNS ARN topic. Triggers an alert for each completed region.")
	snsRegion  = flag.String("sns_region", "", "AWS region for SNS topic. If not set, this value is determined using the host machine's EC2 metadata.")
	snsSubject = flag.String("sns_subject", "", "SNS subject. If set, it overrides the default subject.")
	snsMessage = flag.String("sns_message", "", "SNS message. If set, it overrides the default message.")
)

func init() {
	// log output in JSON formatt
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	flag.Parse()

	if *appVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	var machine metadata.Machine

	if *regions == "" || (*snsTopic != "" && *snsRegion == "") {
		machine = metadata.Machine{}
		if err := machine.LoadFromMetadata(); err != nil {
			log.Fatal("can't get EC2 metadata, must set -regions (and -sns_region) explicitly")
		}
		if *regions == "" {
			*regions = machine.Region
		}
		if *snsRegion == "" {
			*snsRegion = machine.Region
		}
	}

	log.Print("Starting Snapshot Process On " + time.Now().Format(time.RFC822))

	var wg sync.WaitGroup
	for _, region := range strings.Split(*regions, ",") {
		wg.Add(1)
		go func(region string) {
			defer wg.Done()

			mgr := ebs.NewSnapshotManager(region, "", *copyTags, *retainCount, *debug)
			snapshotCount := mgr.SnapshotVolumes()

			if *snsTopic != "" {
				if *snsSubject == "" {
					*snsSubject = fmt.Sprintf("EBS Snapshots Completed (%s)", region)
				}

				if *snsMessage == "" {
					*snsMessage = fmt.Sprintf("%d snapshots completed at %s", snapshotCount, time.Now().Format(time.RFC822))
				}
				sns.SendMessage(*snsRegion, *snsTopic, *snsSubject, *snsMessage)
			}
		}(region)
	}

	wg.Wait()
}
