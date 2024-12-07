package usecase

import (
	"context"
	"fmt"
	"github.com/golang-school/layout/internal/apple/dto"
	"github.com/golang-school/layout/pkg/tracer"
)

func (u *UseCase) GetApple(ctx context.Context, input dto.GetAppleInput) (dto.GetAppleOutput, error) {
	ctx, span := tracer.Start(ctx, "usecase GetApple")
	defer tracer.End(span)

	var output dto.GetAppleOutput

	apple, err := u.redis.GetApple(ctx, input.ID)
	if err != nil {
		return output, fmt.Errorf("u.redis.GetApple: %w", err)
	}

	return dto.GetAppleOutput{
		ID:     apple.ID,
		Name:   apple.Name,
		Status: apple.Status,
	}, err
}
