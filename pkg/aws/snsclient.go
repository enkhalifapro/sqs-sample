package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
	"os"
)

// SNSClient ...
type SNSClient struct {
	client *sns.Client
}

// NewSNSClient ...
func NewSNSClient() *SNSClient {
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
	return &SNSClient{
		client: sns.NewFromConfig(awsCfg),
	}
}

func (c *SNSClient) Publish(topicArn string, msg string) (string, error) {
	input := &sns.PublishInput{
		Message:  &msg,
		TopicArn: &topicArn,
	}
	res, err := c.client.Publish(context.Background(), input)
	if err != nil {
		return "", err
	}
	return *res.MessageId, nil
}
