package base

import (
	"time"

	"github.com/google/uuid"
)

// ID represents a custom type based on UUID for unique identification.
// It provides type safety and distinct handling of UUID-based identifiers.
type ID uuid.UUID

// Model represents a base model structure with common fields used across entities.
// It includes fields for unique identification, creation timestamp, last update timestamp,
// and soft deletion timestamp. This struct is designed to be embedded in other models
// to provide standard database record functionality.
type Model struct {
	ID        ID         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
