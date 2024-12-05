package app

import (
	"github.com/golang-school/layout/internal/apple/controller/http_router"
	"github.com/golang-school/layout/internal/apple/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHTTPRouter(uc *usecase.UseCases) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Get("/live", probe)
	r.Get("/ready", probe)

	r.Group(func(r chi.Router) {
		r.Use(
		// Middlewares: OpenTelemetry, Prometheus
		)

		r.Mount("/", http_router.AppleRouter(uc))
	})

	return r
}

func probe(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
