package pubsub

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (c *CapsulePubSub) RunConsumer(serverUrl string, groupId string, topics []string) {
	c.Logger.Infof("Consumer is starting...")

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": serverUrl,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		c.Logger.Errorf("Failed to create consumer: %v\n", err)
		return
	}

	defer consumer.Close()

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		c.Logger.Errorf("Error subscribing to topic: %v\n", err)
		return
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			c.Logger.Errorf("Consumer error: %v\n", err)
			continue
		}
		c.Logger.Infof("Received message: %s\n", string(msg.Value))
	}
}
