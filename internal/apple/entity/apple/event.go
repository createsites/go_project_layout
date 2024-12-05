package apple

import "github.com/google/uuid"

type CreateEvent struct {
	ID   uuid.UUID
	Name string
}
