package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/pubsub"
)

func subscribe(projectID, topicID, subscriptionID string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal("failed to create cloud pubsub client")
		return err
	}

	topic := client.Topic(topicID)
	sub, err := client.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 20 * time.Second,
	})
	if err != nil {
		sub = client.Subscription(subscriptionID)
	}

	var mu sync.Mutex
	received := 0
	cctx, cancel := context.WithCancel(ctx)

	log.Print("now, start receiving ...")
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		fmt.Printf("Got message: %q\n", string(msg.Data))
		mu.Lock()
		defer mu.Unlock()
		received++
		if received == 10 {
			cancel()
		}
	})

	return err
}

func main() {

	var projectID string
	var topicID string
	flag.StringVar(&projectID, "project", "", "project id")
	flag.StringVar(&topicID, "topic", "", "topic id")
	flag.Parse()

	if projectID == "" {
		projectID = os.Getenv("GCPSD_PROJECT_ID")
	}
	if topicID == "" {
		topicID = os.Getenv("GCPSD_TOPIC_ID")
	}
	subscriptionID := flag.Arg(0)

	log.Printf("topic=%s project=%s\n", topicID, projectID)
	err := subscribe(projectID, topicID, subscriptionID)
	if err != nil {
		panic(err)
	}
}
