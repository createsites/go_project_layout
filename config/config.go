package config

import (
	"fmt"
	"github.com/golang-school/layout/pkg/http_server"
	"github.com/golang-school/layout/pkg/logger"
	"github.com/golang-school/layout/pkg/postgres"
	"github.com/golang-school/layout/pkg/sentry"
	"github.com/golang-school/layout/pkg/tracer"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	Name    string `envconfig:"APP_NAME"    required:"true"`
	Version string `envconfig:"APP_VERSION" required:"true"`
}

type Config struct {
	App      App
	HTTP     http_server.Config
	Postgres postgres.Config
	Tracer   tracer.Config
	Logger   logger.Config
	Sentry   sentry.Config
}

func New() (Config, error) {
	var config Config

	err := godotenv.Load(".env")
	if err != nil {
		return config, fmt.Errorf("godotenv.Load: %w", err)
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return config, fmt.Errorf("envconfig.Process: %w", err)
	}

	return config, nil
}
