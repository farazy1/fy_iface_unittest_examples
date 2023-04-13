package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var svc *sns.SNS

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{}))
	if svc == nil {
		fmt.Println("svc is nil in global, initializing")
		svc = sns.New(sess)
	}
	fmt.Printf("sess is %T\nsvc is %T %v\n", sess, svc, svc)
}
