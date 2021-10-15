package main

import (
	"fmt"
	"sqs-example/pkg/aws"
	"time"
)

func main() {
	snsPublisher := aws.NewSNSClient()
	for {
		msgID, err := snsPublisher.Publish("sqs-example-topic", fmt.Sprintf("Test message %v", time.Now().Unix()))
		if err != nil {
			panic(err)
		}
		fmt.Printf("message %s has been sent successfully\n", msgID)

		time.Sleep(time.Second)
	}
}
