package postgres

import (
	"context"
	"github.com/golang-school/layout/internal/apple/entity"
	"github.com/golang-school/layout/pkg/otel/tracer"
)

func (p *Postgres) CreatePineApple(ctx context.Context, _ entity.PineApple) (err error) {
	ctx, span := tracer.Start(ctx, "postgres CreatePineApple")
	defer span.End()

	return nil
}
