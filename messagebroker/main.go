package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

var brokers = []string{"localhost:9092"}

func main() {
	go publish()
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:       brokers,
			ConsumerGroup: "test_consumer_group",
		},
		saramaSubscriberConfig,
		kafka.DefaultMarshaler{},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), "go-training")
	if err != nil {
		panic(err)
	}

	for m := range messages {
		fmt.Println(m)
	}
}

func publish() {
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	saramaSubscriberConfig.Producer.Return.Successes = true

	publisher, err := kafka.NewPublisher(
		brokers,
		kafka.DefaultMarshaler{},
		saramaSubscriberConfig,
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))
	err = publisher.Publish("go-training", msg)
	if err != nil {
		panic(err)
	}

	log.Print("sent to the topic")
}
