package kafka_producer

import (
	"context"
	"fmt"
	"github.com/golang-school/layout/internal/apple/entity"
	"github.com/golang-school/layout/pkg/tracer"
	"github.com/segmentio/kafka-go"
)

func (p *Producer) CreateEvent(ctx context.Context, e entity.CreateEvent) error {
	ctx, span := tracer.Start(ctx, "kafka_producer UpdateEvent")
	defer tracer.End(span)

	m := kafka.Message{
		Key:   []byte(e.ID.String()),
		Value: []byte(e.Name),
	}

	err := p.writer.WriteMessages(ctx, m)
	if err != nil {
		return fmt.Errorf("p.writer.WriteMessages: %w", err)
	}

	return nil
}
