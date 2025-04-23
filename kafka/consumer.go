package kafka

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

type KafkaConsumerGroup struct {
	consumerGroup sarama.ConsumerGroup
	topics        []string
	ctx           context.Context
	cancel        context.CancelFunc
}

// NewKafkaConsumerGroup creates a new KafkaConsumerGroup instance
func NewKafkaConsumerGroup(brokers []string, groupID string, topics []string) (*KafkaConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	cg, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &KafkaConsumerGroup{
		consumerGroup: cg,
		topics:        topics,
		ctx:           ctx,
		cancel:        cancel,
	}, nil
}

// StartConsuming starts the consumer loop using the given handler
func (k *KafkaConsumerGroup) StartConsuming(handler sarama.ConsumerGroupHandler) {
	go func() {
		for {
			err := k.consumerGroup.Consume(k.ctx, k.topics, handler)
			if err != nil {
				log.Printf("Error from consumer: %v", err)
			}
			if k.ctx.Err() != nil {
				return
			}
		}
	}()

	go func() {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, os.Interrupt)
		<-sigterm
		k.Stop()
	}()
}

// Stop gracefully shuts down the consumer group
func (k *KafkaConsumerGroup) Stop() {
	k.cancel()
	_ = k.consumerGroup.Close()
}