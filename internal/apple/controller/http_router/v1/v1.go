package v1

import "github.com/golang-school/layout/internal/apple/usecase"

type Handlers struct {
	uc *usecase.UseCases
}

func New(uc *usecase.UseCases) *Handlers {
	return &Handlers{
		uc: uc,
	}
}
