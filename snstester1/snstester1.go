package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

func sNSPublish(svc snsiface.SNSAPI) bool {
	input := &sns.PublishInput{
		Message:  aws.String(string("message")),
		TopicArn: aws.String("outputSNSARN"),
	}

	snsPublishOutput, error := svc.Publish(input)
	fmt.Printf("mySNSPublish Output Type [%T] Output Contents [%v], Error Type [%T] Error Contents [%v]\n", snsPublishOutput, snsPublishOutput, error, error)
	return true
}

func main() {
	//	sess := session.New()
	sess := session.Must(session.NewSessionWithOptions(session.Options{}))
	svc := sns.New(sess)
	sNSPublish(svc)
}
