package tracer

import (
	"context"
	"github.com/rs/zerolog/log"
)

type Config struct {
	AppName    string  `envconfig:"APP_NAME" required:"true"`
	AppVersion string  `envconfig:"APP_VERSION" required:"true"`
	Endpoint   string  `envconfig:"OTEL_EXPORTER_OTLP_ENDPOINT"`
	SampleRate float64 `envconfig:"OTEL_EXPORTER_TRANSACTION_SAMPLE_RATE" default:"1"`
	PodName    string  `envconfig:"POD_NAME"`
}

type Span struct{}

func Init(c Config) error {
	if c.Endpoint == "" || c.PodName == "" {
		log.Info().Msg("Tracer is disabled")

		return nil
	}

	log.Info().Msg("Tracer initialized")

	return nil
}

func Start(ctx context.Context, spanName string) (context.Context, *Span) {
	return ctx, &Span{}
}

func End(span *Span) {}

func Close() {}
