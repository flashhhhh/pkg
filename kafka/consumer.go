package kafka

import (
	"github.com/IBM/sarama"
)

type KafkaConsumer struct {
	consumer sarama.Consumer
}

func NewKafkaConsumer(brokers []string) (*KafkaConsumer, error) {
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer: consumer,
	}, nil
}

func (kc *KafkaConsumer) ConsumePartition(topic string, partition int32, offset int64) (sarama.PartitionConsumer, error) {
	partitionConsumer, err := kc.consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		return nil, err
	}
	return partitionConsumer, nil
}

func (kc *KafkaConsumer) Close() error {
	err := kc.consumer.Close()
	if err != nil {
		return err
	}
	return nil
}