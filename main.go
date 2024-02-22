package main

import (
	"context"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
	log "github.com/sirupsen/logrus"
)

func main() {
	connectionString := os.Getenv("CONNECTION_STRING")
	queueName := os.Getenv("QUEUE_NAME")

	// create a client with the provided connection string
	client, err := azqueue.NewServiceClientFromConnectionString(connectionString, nil)
	if err != nil {
		log.Fatalf("unable to connect to queue: %w", err)
	}
	queueClient := client.NewQueueClient(queueName)
	resp, err := queueClient.DequeueMessage(context.Background(), &azqueue.DequeueMessageOptions{})
	if err != nil {
		log.Fatalf("error getting messages: %w", err)
	}
	for _, m := range resp.Messages {
		log.Infof("message text: %s", *m.MessageText)
	}
}
