package core

import "github.com/google/uuid"

type UserSession struct {
	ID uuid.UUID `json:"id"`
}
