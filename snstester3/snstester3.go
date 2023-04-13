package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

var svc snsiface.SNSAPI

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{}))
	if svc == nil {
		fmt.Println("svc is nil in global, initializing")
		svc = sns.New(sess)
	}
	fmt.Printf("sess is %T\nsvc is %T %v\n", sess, svc, svc)

	input := &sns.PublishInput{
		Message:  aws.String(string("message")),
		TopicArn: aws.String("outputSNSARN"),
	}

	snsPublishOutput, error := svc.Publish(input)
	fmt.Printf("snsPublishOutput Type [%T] Output Contents [%v], Error Type [%T] Error Contents [%v]\n", snsPublishOutput, snsPublishOutput, error, error)
}
