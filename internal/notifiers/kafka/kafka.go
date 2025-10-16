package kafka

import (
	notifiers "xm/internal/notifiers"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	publisher *kafka.Producer
}

func InitKafka(publisher *kafka.Producer) notifiers.Notifier {
	return &Kafka{publisher: publisher}
}

func (k *Kafka) Notify(topic string, data []byte) error {
	return k.publisher.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)
}
