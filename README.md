# EBS Snapshotter

Uses the AWS SDK to create EBS snapshots based on existing volumes.

## Installation

Download a binary from the releases section or build locally (see below).

## Usage

```
ebs_snapshotter -regions=us-west-2,us-east-1 -retain=5 -sns_topic="arn:aws:sns:us-west-2:123456789:BackupAlerts"
```

Run ```ebs_snapshotter -h``` for options.

Be sure you set AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION in your environment, required by the AWS SDK.

## Example Crontab

Backup at 1am nightly:
```
0 1 * * * /bin/ebs_snapshotter -regions=us-west-2,us-east-1 -retain=5 -sns_topic="arn:aws:sns:us-west-2:123456789:BackupAlerts" > /var/log/cron.log
```

## Building

Builds are handled through a Docker image, see [makefile](Makefile).
