package app

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-school/layout/config"
	"github.com/golang-school/layout/pkg/http_server"
	"github.com/golang-school/layout/pkg/kafka_reader"
	"github.com/golang-school/layout/pkg/kafka_writer"
	"github.com/golang-school/layout/pkg/postgres"
	"github.com/golang-school/layout/pkg/redis"
	"github.com/golang-school/layout/pkg/router"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

type Dependencies struct {
	// Adapters
	Postgres    *postgres.Pool
	KafkaWriter *kafka_writer.Writer
	Redis       *redis.Client

	// Controllers
	RouterHTTP  *chi.Mux
	KafkaReader *kafka_reader.Reader
}

func Run(ctx context.Context, c config.Config) (err error) {
	var deps Dependencies

	// Adapters
	deps.Postgres, err = postgres.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}
	defer deps.Postgres.Close()

	deps.KafkaWriter, err = kafka_writer.New(c.KafkaWriter)
	if err != nil {
		return fmt.Errorf("kafka_writer.New: %w", err)
	}
	defer deps.KafkaWriter.Close()

	deps.Redis, err = redis.New(c.Redis)
	if err != nil {
		return fmt.Errorf("redis.New: %w", err)
	}
	defer deps.Redis.Close()

	// Controllers
	deps.RouterHTTP = router.New()

	deps.KafkaReader, err = kafka_reader.New(c.KafkaReader)
	if err != nil {
		return fmt.Errorf("kafka_reader.New: %w", err)
	}
	defer deps.KafkaReader.Close()

	// Domains
	AppleDomain(deps)

	httpServer := http_server.New(deps.RouterHTTP, c.HTTP.Port)
	defer httpServer.Close()

	waiting(httpServer)

	return nil
}

func waiting(httpServer *http_server.Server) {
	log.Info().Msg("App started!")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	select {
	case i := <-wait:
		log.Info().Msg("App got signal: " + i.String())
	case err := <-httpServer.Notify():
		log.Error().Err(err).Msg("App got notify: httpServer.Notify")
	}

	log.Info().Msg("App is stopping...")
}
