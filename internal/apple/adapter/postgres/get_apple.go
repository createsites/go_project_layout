package postgres

import (
	"context"
	"github.com/golang-school/layout/internal/apple/entity/apple"
	"github.com/google/uuid"

	"github.com/golang-school/layout/pkg/tracer"
)

func (p *Postgres) GetApple(ctx context.Context, id uuid.UUID) (apple.Apple, error) {
	ctx, span := tracer.Start(ctx, "postgres GetApple")
	defer tracer.End(span)

	return apple.Apple{ID: id, Status: "from postgres"}, nil
}
