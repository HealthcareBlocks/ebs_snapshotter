// Package awserror handles AWS SDK error responses
package awserror

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

// HandleError parses an AWS error and exits the program
func HandleError(err error) {
	if err == nil {
		return
	}

	if awsErr, ok := err.(awserr.Error); ok {
		log.Print("Code: " + awsErr.Code())
		log.Print("Message: " + awsErr.Message())

		if awsErr.OrigErr() != nil {
			log.Printf("Orginal Error: %v", awsErr.OrigErr())
		}

		if reqErr, ok := err.(awserr.RequestFailure); ok {
			log.Printf("Status Code: %d", reqErr.StatusCode())

			if reqErr.RequestID() != "" {
				log.Print("Request ID: " + reqErr.RequestID())
			}
		}
	} else {
		log.Print(err.Error())
	}
	os.Exit(1)
}
