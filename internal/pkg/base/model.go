package base

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index;type:timestamp"`
}
