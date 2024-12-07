package v1

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/golang-school/layout/internal/apple/dto"
	"github.com/golang-school/layout/internal/apple/entity"
	"github.com/golang-school/layout/pkg/render"
	"github.com/golang-school/layout/pkg/tracer"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handlers) GetApple(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "http/v1 GetApple")
	defer tracer.End(span)

	var (
		input dto.GetAppleInput
		err   error
	)

	id := chi.URLParam(r, "id")

	input.ID, err = uuid.Parse(id)
	if err != nil {
		log.Error().Err(err).Msg("uuid.Parse")
		http.Error(w, "uuid validate error", http.StatusBadRequest)

		return
	}

	output, err := h.uc.GetApple(ctx, input)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			log.Error().Err(err).Msg("uc.CreateApple: not found")
			http.Error(w, "not found", http.StatusNotFound)

			return

		case errors.Is(err, entity.ErrUUIDInvalid), errors.Is(err, entity.ErrStatusInvalid):
			log.Error().Err(err).Msg("uc.CreateApple: validate error")
			http.Error(w, "validate error", http.StatusBadRequest)

			return

		default:
			log.Error().Err(err).Msg("uc.CreateApple: internal error")
			http.Error(w, "internal error", http.StatusInternalServerError)

			return
		}
	}

	render.JSON(w, output)
}
