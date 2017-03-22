// Package sns implements functionality for sending an Amazon SNS message
//
// Note: this package relies on the AWS SDK, thus the host environment should
// either have an associated IAM role or user with the SNS:Publish permission.
package sns

import (
	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/healthcareblocks/ebs_snapshotter/awserror"
)

// SendMessage sends an SNS message to an SNS region.
// See http://docs.aws.amazon.com/sdk-for-go/api/service/sns.html#type-PublishInput
func SendMessage(region string, topic string, subject string, message string) {
	if region == "" {
		log.Fatal("SNS region is required")
	}

	if topic == "" {
		log.Fatal("SNS topic is required")
	}

	if subject == "" {
		log.Fatal("SNS subject is required")
	}

	if message == "" {
		log.Fatal("SNS message is required")
	}

	params := &sns.PublishInput{
		Subject:   aws.String(subject),
		Message:   aws.String(message),
		TargetArn: aws.String(topic),
	}

	sess, sessErr := session.NewSession(aws.NewConfig().WithRegion(region))
	awserror.HandleError(sessErr)

	sns := sns.New(sess)

	_, publishErr := sns.Publish(params)
	awserror.HandleError(publishErr)
}
