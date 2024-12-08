package v1

import (
	"encoding/json"
	"errors"
	"github.com/golang-school/layout/internal/apple/dto"
	"github.com/golang-school/layout/internal/apple/entity"
	"github.com/golang-school/layout/pkg/otel/tracer"
	"github.com/golang-school/layout/pkg/render"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	"net/http"
)

func (h *Handlers) CreateApple(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "http/v1 CreateApple")
	defer span.End()

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

			span.SetAttributes(attribute.KeyValue{Key: "error", Value: attribute.StringValue(err.Error())})
			tracer.SetStatus(span, err)

			return
		}
	}

	render.JSON(w, output)
}
