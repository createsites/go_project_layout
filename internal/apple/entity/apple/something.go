package apple

import "github.com/google/uuid"

type Something struct {
	ID     uuid.UUID `json:"id"`
	Action string    `json:"action"`
}
