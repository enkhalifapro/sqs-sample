package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"log"
	"os"
)

// SQSClient ...
type SQSClient struct {
	client *sqs.Client
}

// NewSQSClient ...
func NewSQSClient() *SQSClient {
	r := os.Getenv(region)
	if r == "" {
		log.Printf("No AWS Region found for env var AWS_REGION. setting defaultRegion=%s \n", defaultRegion)
		r = defaultRegion
	}
	awsCfg := aws.Config{
		Region: r,
	}
	if os.Getenv("AWS_ENDPOINT") != "" {
		awsCfg.EndpointResolver = aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           os.Getenv("AWS_ENDPOINT"), // usually should have localStack Endpoint "http://localhost:4566/"
				SigningRegion: r,
			}, nil
		})
	}
	return &SQSClient{
		client: sqs.NewFromConfig(awsCfg),
	}
}

func (c *SQSClient) Receive(queue string, timeout int) ([]types.Message, error) {
	input := &sqs.ReceiveMessageInput{
		MessageAttributeNames: []string{
			string(types.QueueAttributeNameAll),
		},
		QueueUrl:            &queue,
		MaxNumberOfMessages: 1,
		VisibilityTimeout:   int32(timeout),
	}
	res, err := c.client.ReceiveMessage(context.Background(), input)
	if err != nil {
		return []types.Message{}, err
	}
	return res.Messages, nil
}
