package service

import (
	"time"

	"github.com/geekswamp/zen/configs"
	"github.com/geekswamp/zen/internal/base"
	"github.com/geekswamp/zen/internal/crypto/password"
	"github.com/geekswamp/zen/internal/model"
	"github.com/geekswamp/zen/internal/repository"
	"github.com/google/uuid"
)

type UserService interface {
	Create(fullName, email, passwordStr string, phone string, gender model.Gender) error
	Get(id uuid.UUID) (*model.User, error)
	Update(id uuid.UUID, userMap base.UpdateMap) error
	Delete(id uuid.UUID) error
	SoftDelete(id uuid.UUID) error
	SetToActive(id uuid.UUID) error
	SetToInactive(id uuid.UUID) error
}

type UserServiceRepo struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserServiceRepo{repo: repo}
}

func (s UserServiceRepo) Create(fullName, email, passwordStr string, phone string, gender model.Gender) error {
	user := model.User{
		FullName: fullName,
		Email:    email,
		Phone:    phone,
		Gender:   gender,
	}

	pc := password.NewFromConfig(configs.Get())
	hash, err := pc.Generate([]byte(passwordStr))
	if err != nil {
		return err
	}

	return s.repo.Create(user, hash)
}

func (s UserServiceRepo) Get(id uuid.UUID) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s UserServiceRepo) Update(id uuid.UUID, userMap base.UpdateMap) error {
	return s.repo.Update(id, userMap)
}

func (s UserServiceRepo) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s UserServiceRepo) SoftDelete(id uuid.UUID) error {
	return s.Update(id, base.UpdateMap{"Active": false, "DeletedTime": time.Now().Local().UnixMilli()})
}

func (s UserServiceRepo) SetToActive(id uuid.UUID) error {
	return s.Update(id, base.UpdateMap{"Active": true, "ActivatedTime": time.Now().Local().UnixMilli()})
}

func (s UserServiceRepo) SetToInactive(id uuid.UUID) error {
	return s.Update(id, base.UpdateMap{"Active": false})
}
