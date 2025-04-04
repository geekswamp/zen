package user

import "github.com/google/uuid"

type UserCreateRequest struct {
	FullName string `json:"full_name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"omitempty,e164"`
	Gender   int    `json:"gender" validate:"required,oneof=0 1"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}

type UserUpdateInfoRequest struct {
	FullName string `json:"full_name" validate:"min=3,max=100"`
	Email    string `json:"email" validate:"email"`
	Phone    string `json:"phone" validate:"omitempty,e164"`
	Gender   int    `json:"gender" validate:"oneof=0 1"`
}

type UserInfoResponse struct {
	ID            uuid.UUID `json:"id"`
	FullName      string    `json:"full_name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Active        bool      `json:"active"`
	Gender        int       `json:"gender"`
	ActivatedTime int64     `json:"activated_time"`
	CreatedTime   int64     `json:"created_time"`
	UpdateTime    int64     `json:"updated_time"`
	DeletedTime   *int64    `json:"deleted_time"`
}
