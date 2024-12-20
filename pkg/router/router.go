package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Get("/live", probe)
	r.Get("/ready", probe)

	return r
}

func probe(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
