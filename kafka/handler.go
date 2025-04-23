/*
	This file is just a placeholder for the Kafka consumer group handler.
	You can implement your own handler by creating a struct that implements the sarama.ConsumerGroupHandler interface.
*/

package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

type MyConsumerHandler struct{}

func (h MyConsumerHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h MyConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (h MyConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("Consumed message: %s\n", string(message.Value))
		session.MarkMessage(message, "")
	}
	return nil
}