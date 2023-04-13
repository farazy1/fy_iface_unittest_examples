package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func TestMyFunc(t *testing.T) {
	sess := session.New()
	svc = sns.New(sess)

	main()

}
