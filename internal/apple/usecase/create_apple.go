package usecase

import (
	"context"
	"fmt"
	"github.com/golang-school/layout/internal/apple/dto"
	"github.com/golang-school/layout/internal/apple/entity/apple"
	"github.com/google/uuid"

	"github.com/golang-school/layout/pkg/tracer"
)

func (u *UseCases) CreateApple(ctx context.Context, input dto.CreateAppleInput) (dto.CreateAppleOutput, error) {
	ctx, span := tracer.Start(ctx, "usecase CreateApple")
	defer tracer.End(span)

	var output dto.CreateAppleOutput

	a := apple.Apple{
		ID:     uuid.New(),
		Name:   input.Name,
		Status: apple.StatusNew,
	}

	err := u.postgres.CreateApple(ctx, a)
	if err != nil {
		return output, fmt.Errorf("u.postgres.CreateApple: %w", err)
	}

	err = u.redis.PutApple(ctx, a)
	if err != nil {
		return output, fmt.Errorf("u.redis.PutApple: %w", err)
	}

	event := apple.CreateEvent{
		ID:   a.ID,
		Name: input.Name,
	}

	err = u.kafka.CreateEvent(ctx, event)
	if err != nil {
		return output, fmt.Errorf("u.kafka.CreateEvent: %w", err)
	}

	output.ID = a.ID

	return output, nil
}
