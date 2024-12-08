package postgres

import (
	"context"
	"github.com/golang-school/layout/internal/apple/entity"
	"github.com/golang-school/layout/pkg/otel/tracer"
	"github.com/google/uuid"
)

func (p *Postgres) GetApple(ctx context.Context, id uuid.UUID) (entity.Apple, error) {
	ctx, span := tracer.Start(ctx, "postgres GetApple")
	defer span.End()

	return entity.Apple{ID: id, Status: "from postgres"}, nil
}
