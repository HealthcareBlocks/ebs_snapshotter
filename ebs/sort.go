package ebs

import "github.com/aws/aws-sdk-go/service/ec2"

// ByStartTime implements sort.Interface for ec2.Snapshot based on the StartTime field
type ByStartTime []*ec2.Snapshot

func (t ByStartTime) Len() int           { return len(t) }
func (t ByStartTime) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ByStartTime) Less(i, j int) bool { return t[i].StartTime.Before(*t[j].StartTime) }
