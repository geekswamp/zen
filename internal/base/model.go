package base

import "github.com/google/uuid"

// ID represents a custom type based on UUID for unique identification.
type ID uuid.UUID

// Model represents a base model structure with common fields used across entities.
type Model struct {
	ID          ID     `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedTime int64  `json:"created_itme" gorm:"column:created_time;autoCreateTime:milli"`
	UpdatedTime *int64 `json:"updated_itme" gorm:"column:updated_time;autoUpdateTime:milli"`
	DeletedTime *int64 `json:"deleted_itme" gorm:"column:deleted_time;index"`
}
