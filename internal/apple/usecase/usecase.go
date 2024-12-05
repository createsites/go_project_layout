package usecase

import (
	"context"
	apple_postgres "github.com/golang-school/layout/internal/apple/adapter/postgres"
	"github.com/golang-school/layout/internal/apple/entity/apple"
	"github.com/golang-school/layout/internal/apple/entity/pineapple"
	"github.com/golang-school/layout/pkg/postgres"

	"github.com/google/uuid"
)

type Postgres interface {
	CreateApple(ctx context.Context, a apple.Apple) (err error)
	GetApple(ctx context.Context, id uuid.UUID) (apple.Apple, error)

	CreatePineApple(ctx context.Context, a pineapple.PineApple) (err error)
}

type Kafka interface {
	CreateEvent(ctx context.Context, s apple.CreateEvent) error
}

type Redis interface {
	GetApple(ctx context.Context, id uuid.UUID) (apple.Apple, error)
	PutApple(ctx context.Context, a apple.Apple) error
}

type UseCases struct {
	postgres Postgres
	kafka    Kafka
	redis    Redis
}

func New(p *postgres.Postgres, k Kafka, r Redis) *UseCases {
	return &UseCases{
		postgres: apple_postgres.New(p),
		kafka:    k,
		redis:    r,
	}
}
