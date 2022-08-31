package consumer

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"homework/config/kafka"
	"homework/internal/api/storage/kafka/consumer/handler"
)

type consumerHandler func(context.Context, []byte, []byte)

type Consumer struct {
	initialOffset int64

	consumerKafka sarama.Consumer
	handlersMap map[string]consumerHandler
	handler *handler.Handler
}

func (c *Consumer) subscribe(topic string, handler consumerHandler) error{
	partitionList, err := c.consumerKafka.Partitions(topic)
	if err != nil {
		err = fmt.Errorf("error retrieving partitionList: %v", err.Error())
		logrus.Error(err.Error())
		return err
	}

	for _, partition := range partitionList {
		pc, _ := c.consumerKafka.ConsumePartition(topic, partition, c.initialOffset)

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				handler(context.Background(), message.Key, message.Value)
			}
		}(pc)
	}
	return nil
}

func (c *Consumer) registerHandler(topic string, handler consumerHandler) {
	c.handlersMap[topic] = handler
}

func (c *Consumer) Serve() error {
	for topic, handler := range c.handlersMap {
		if err := c.subscribe(topic, handler); err != nil {
			logrus.Errorf("could not subscribe topic '%v': %v", topic, err)
			return err
		}
		logrus.Infof("subscribed to topic '%v'", topic)
	}

	return nil
}

func NewConsumer(brokers []string, handler *handler.Handler) (*Consumer, error){
	consumerKafka, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return nil, err
	}

	consumer := &Consumer{
		consumerKafka: consumerKafka,
		initialOffset: sarama.OffsetNewest,
		handlersMap: make(map[string]consumerHandler),
		handler: handler,
	}

	consumer.registerHandler(kafka.TopicCreate, handler.Create)
	consumer.registerHandler(kafka.TopicUpdate, handler.Update)
	consumer.registerHandler(kafka.TopicDelete, handler.Delete)

	return consumer, nil
}
