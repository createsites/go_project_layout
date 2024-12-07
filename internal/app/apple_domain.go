package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/golang-school/layout/internal/apple/adapter/kafka_producer"
	"github.com/golang-school/layout/internal/apple/adapter/postgres"
	"github.com/golang-school/layout/internal/apple/adapter/redis"
	"github.com/golang-school/layout/internal/apple/controller/http_router"
	"github.com/golang-school/layout/internal/apple/usecase"
)

func AppleDomain(router *chi.Mux, d Dependencies) {
	appleUsecase := usecase.New(
		postgres.New(d.Postgres.Pool),
		kafka_producer.New(d.KafkaWriter.Writer),
		redis.New(d.Redis.Client),
	)

	http_router.AppleRouter(router, appleUsecase)
}
