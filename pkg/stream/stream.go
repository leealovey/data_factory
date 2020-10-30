package stream

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// Config hanlde kafak message
type Config struct {
	Protocal string
	URI      string
}

func (s *Config) writer(ctx context.Context, topic string, msg kafka.Message) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{s.URI},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	err := w.WriteMessages(ctx, msg)
	if err != nil {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Config) read(ctx context.Context, topic string) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{s.URI},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	r.SetOffset(42)

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}

		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		return err
	}

	return nil
}
