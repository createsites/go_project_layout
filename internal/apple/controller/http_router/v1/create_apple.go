package v1

import (
	"encoding/json"
	"errors"
	"github.com/golang-school/layout/internal/apple/dto"
	"github.com/golang-school/layout/internal/apple/entity"
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/golang-school/layout/pkg/tracer"
)

func (h *Handlers) CreateApple(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "http/v1 CreateApple")
	defer tracer.End(span)

	input := dto.CreateAppleInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Error().Err(err).Msg("json.NewDecoder")
		http.Error(w, "json error", http.StatusBadRequest)

		return
	}

	err = input.Validate()
	if err != nil {
		log.Error().Err(err).Msg("input.Validate")
		http.Error(w, "validate error", http.StatusBadRequest)

		return
	}

	output, err := h.uc.CreateApple(ctx, input)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			http.Error(w, "not found", http.StatusNotFound)

			return

		case errors.Is(err, entity.ErrUUIDInvalid), errors.Is(err, entity.ErrStatusInvalid):
			http.Error(w, "validate error", http.StatusBadRequest)

			return

		default:
			log.Error().Err(err).Msg("uc.CreateApple")
			http.Error(w, "internal error", http.StatusInternalServerError)

			return
		}
	}

	_ = output

	w.WriteHeader(http.StatusOK)
}
