package http_router

import (
	"github.com/go-chi/chi/v5"
	ver1 "github.com/golang-school/layout/internal/apple/controller/http_router/v1"
	"github.com/golang-school/layout/internal/apple/usecase"
	"net/http"
)

func AppleRouter(uc *usecase.UseCases) http.Handler {
	r := chi.NewRouter()

	r.Route("/api/apple", func(r chi.Router) {

		v1 := ver1.New(uc)

		r.Route("/v1", func(r chi.Router) {
			r.Post("/create_apple", v1.CreateApple)
			r.Get("/get_apple", v1.GetApple)
			r.Put("/update_apple", v1.DeleteApple)
			r.Delete("/delete_apple", v1.DeleteApple)
		})
	})

	return r
}
