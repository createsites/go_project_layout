package redis

import (
	"context"
	"github.com/golang-school/layout/internal/apple/entity/apple"

	"github.com/golang-school/layout/pkg/tracer"
)

func (r *Redis) PutApple(ctx context.Context, a apple.Apple) error {
	ctx, span := tracer.Start(ctx, "redis PutApple")
	defer tracer.End(span)

	return nil
}
