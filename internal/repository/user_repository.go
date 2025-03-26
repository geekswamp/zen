package repository

import (
	"github.com/geekswamp/zen/internal/base"
	"github.com/geekswamp/zen/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user model.User, passHash string) error
	FindByID(id uuid.UUID) (*model.User, error)
	IsExist(user *model.User) (bool, error)
	Update(id uuid.UUID, userMap base.UpdateMap) error
	Delete(id uuid.UUID) error
}

type UserQueryBuilder struct{ repo base.Repository }

func New(repo base.Repository) UserRepository {
	return UserQueryBuilder{repo: repo}
}

func (q UserQueryBuilder) Create(user model.User, passHash string) error {
	return q.repo.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		passHash := model.UserPassHash{UserID: user.ID, PassHash: passHash}
		if err := tx.Create(&passHash).Error; err != nil {
			return err
		}

		return nil
	})
}

func (q UserQueryBuilder) FindByID(id uuid.UUID) (*model.User, error) {
	user := model.User{}
	if err := q.repo.DB().First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (q UserQueryBuilder) IsExist(user *model.User) (bool, error) {
	err := q.repo.DB().Where("email = ?", user.Email).Or("phone = ?", user.Phone).First(&model.User{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (q UserQueryBuilder) Update(id uuid.UUID, userMap base.UpdateMap) error {
	user, err := q.FindByID(id)
	if err != nil {
		return err
	}

	return q.repo.DB().Model(user).Updates(userMap).Error
}

func (q UserQueryBuilder) Delete(id uuid.UUID) error {
	return q.repo.DB().Delete(&model.User{}, id).Error
}
