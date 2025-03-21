package base

import "github.com/google/uuid"

// Model represents a base model structure with common fields used across entities.
type Model struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedTime int64     `json:"created_time" gorm:"column:created_time;autoCreateTime:milli"`
	UpdatedTime *int64    `json:"updated_time" gorm:"column:updated_time;autoUpdateTime:milli"`
	DeletedTime *int64    `json:"deleted_time" gorm:"column:deleted_time;index"`
}
