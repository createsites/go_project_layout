package main

import (
	"context"
	"github.com/golang-school/layout/config"
	"github.com/golang-school/layout/internal/app"
	"github.com/golang-school/layout/pkg/logger"
	"github.com/golang-school/layout/pkg/otel"
	"github.com/golang-school/layout/pkg/sentry"
	"github.com/rs/zerolog/log"
	_ "go.uber.org/automaxprocs"
)

func main() {
	ctx := context.Background()

	c, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}

	logger.Init(c.Logger)

	err = sentry.Init(c.Sentry)
	if err != nil {
		log.Error().Err(err).Msg("sentry.Init")
	}

	defer sentry.Close()

	err = otel.Init(ctx, c.OTEL)
	if err != nil {
		log.Error().Err(err).Msg("otel.Init")
	}

	defer otel.Close()

	err = app.Run(ctx, c)
	if err != nil {
		log.Fatal().Err(err).Msg("app.Run")
	}

	log.Info().Msg("App stopped!")
}
