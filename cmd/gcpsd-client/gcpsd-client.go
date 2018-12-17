package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
)

func publish(projectID, topicID, imageID string) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}

	data := []byte(imageID)

	topic := client.Topic(topicID)
	_, err = topic.Publish(ctx, &pubsub.Message{Data: data}).Get(ctx)
	if err != nil {
		panic(err)
	}
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
	imageID := flag.Arg(0)

	fmt.Printf("publishing %s to topic=%s project=%s\n", imageID, topicID, projectID)
	publish(projectID, topicID, imageID)
}
