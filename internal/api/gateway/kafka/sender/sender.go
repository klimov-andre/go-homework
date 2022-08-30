package sender

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type Sender struct {
	kafkaProducer sarama.SyncProducer
}

func NewSender(brokers []string) (*Sender, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		return nil, fmt.Errorf("sync sender creation error: %v", err)
	}

	return &Sender{
		kafkaProducer: syncProducer,
	}, nil
}

func (s *Sender) SendMessage(topic, key string, payload []byte) {
	par, off, err := s.kafkaProducer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(payload),
	})

	_, _, _ = par, off, err
}