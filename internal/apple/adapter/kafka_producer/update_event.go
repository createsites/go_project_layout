package kafka_producer

import (
	"context"
	"github.com/golang-school/layout/internal/apple/entity/apple"

	"github.com/golang-school/layout/pkg/tracer"
)

func (p *KafkaProducer) CreateEvent(ctx context.Context, s apple.CreateEvent) error {
	ctx, span := tracer.Start(ctx, "kafka_producer UpdateEvent")
	defer tracer.End(span)

	return nil
}
