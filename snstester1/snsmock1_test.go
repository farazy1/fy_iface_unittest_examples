package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

// Define a mock struct to be used in your unit tests of myFunc.
type mockSNSClient struct {
	snsiface.SNSAPI
}

func (m *mockSNSClient) Publish(input *sns.PublishInput) (*sns.PublishOutput, error) {
	return new(sns.PublishOutput), nil
}

func TestSNSPublish(t *testing.T) {
	mockSvc := &mockSNSClient{}
	sNSPublish(mockSvc)
}
