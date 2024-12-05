package postgres

import (
	"context"
	"github.com/golang-school/layout/internal/apple/entity/pineapple"

	"github.com/golang-school/layout/pkg/tracer"
)

func (p *Postgres) CreatePineApple(ctx context.Context, a pineapple.PineApple) (err error) {
	ctx, span := tracer.Start(ctx, "postgres CreatePineApple")
	defer tracer.End(span)

	return nil
}
