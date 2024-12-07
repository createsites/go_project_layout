package v1

import "github.com/golang-school/layout/internal/apple/usecase"

type Handlers struct {
	uc *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{
		uc: uc,
	}
}
