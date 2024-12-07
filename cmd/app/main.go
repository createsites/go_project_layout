package main

import (
	"github.com/golang-school/layout/config"
	"github.com/golang-school/layout/internal/app"
	"github.com/golang-school/layout/pkg/logger"
	"github.com/golang-school/layout/pkg/sentry"
	"github.com/golang-school/layout/pkg/tracer"
	"github.com/rs/zerolog/log"
	_ "go.uber.org/automaxprocs"
)

func main() {
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

	err = tracer.Init(c.Tracer)
	if err != nil {
		log.Error().Err(err).Msg("tracer.Init")
	}

	defer tracer.Close()

	err = app.Run(c)
	if err != nil {
		log.Fatal().Err(err).Msg("app.Run")
	}

	log.Info().Msg("App stopped!")
}
