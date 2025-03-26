package model

import (
	"github.com/geekswamp/zen/internal/base"
	"github.com/google/uuid"
)

type Gender int

const (
	Female Gender = iota
	Male
)

type User struct {
	base.Model
	FullName      string       `gorm:"column:full_name;type:varchar;not null"`
	Email         string       `gorm:"column:email;type:varchar;uniqueIndex:idx_users_contact;not null"`
	Phone         *string      `gorm:"column:phone;type:varchar;uniqueIndex:idx_users_contact"`
	Active        bool         `gorm:"column:active;type:boolean;default:false;not null"`
	Gender        Gender       `gorm:"column:gender;type:smallint;not null"`
	ActivatedTime int64        `gorm:"column:activated_time"`
	PassHash      UserPassHash `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

type UserPassHash struct {
	UserID   uuid.UUID `gorm:"column:user_id;primaryKey;type:uuid"`
	PassHash string    `gorm:"column:pass_hash;type:varchar;not null"`
}
