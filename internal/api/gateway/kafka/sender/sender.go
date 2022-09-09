package sender

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"homework/config/gateway"
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

func (s *Sender) SendMessage(ctx context.Context, topic, key string, payload []byte) {
	var span trace.Span
	ctx, span = otel.Tracer(gateway.SpanTraceName).Start(ctx, "SendMessage")
	defer span.End()

	_, _, err := s.kafkaProducer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(payload),
	})

	if err != nil {
		span.RecordError(err)
		return
	}
}