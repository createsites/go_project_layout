package app

import (
	"context"
	"fmt"
	"github.com/golang-school/layout/config"
	"github.com/golang-school/layout/internal/apple/adapter/kafka_producer"
	"github.com/golang-school/layout/internal/apple/adapter/redis"
	"github.com/golang-school/layout/internal/apple/usecase"
	"github.com/golang-school/layout/pkg/postgres"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-school/layout/pkg/http_server"
)

func Run(c config.Config) error {
	ctx := context.Background()

	postgres, err := postgres.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}

	defer postgres.Close()

	kafka, err := kafka_producer.New()
	if err != nil {
		return fmt.Errorf("kafka_producer.New: %w", err)
	}

	defer kafka.Close()

	redis, err := redis.New()
	if err != nil {
		return fmt.Errorf("redis.New: %w", err)
	}

	defer redis.Close()

	// AppleDomain
	appleUseCases := usecase.New(postgres, kafka, redis)

	router := NewHTTPRouter(appleUseCases)

	s := http_server.New(router, c.HTTP.Port)
	defer s.Close()

	waiting()

	return nil
}

func waiting() {
	log.Info().Msg("App started!")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	<-wait

	log.Info().Msg("App is stopping...")
}
