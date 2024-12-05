package dto

import "github.com/google/uuid"

type GetAppleOutput struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type GetAppleInput struct {
	ID uuid.UUID `json:"id"`
}

func (i *GetAppleInput) Validate() error {

	return nil
}
