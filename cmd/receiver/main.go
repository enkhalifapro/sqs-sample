package main

import (
	"fmt"
	"sqs-example/pkg/aws"
	"time"
)

func main() {
	snsPublisher := aws.NewSQSClient()
	for {
		msgs, err := snsPublisher.Receive("sqs-example-queue", 20)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s have been received successfully\n", len(msgs))

		time.Sleep(time.Second)
	}
}
