package kafka

import (
	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(broker []string) (*KafkaProducer, error) {
	producer, err := sarama.NewSyncProducer(broker, nil)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		producer: producer,
	}, nil
}

func (kp *KafkaProducer) SendMessage(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	_, _, err := kp.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

func (kp *KafkaProducer) SendMessageWithKey(topic string, key string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(message),
	}

	_, _, err := kp.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

func (kp *KafkaProducer) Close() error {
	err := kp.producer.Close()
	if err != nil {
		return err
	}
	return nil
}