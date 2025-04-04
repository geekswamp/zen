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

func NewUserRepo(repo base.Repository) UserRepository {
	return UserQueryBuilder{repo: repo}
}

func (q UserQueryBuilder) Create(user model.User, passHash string) error {
	err := q.repo.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		passHashModel := model.UserPassHash{UserID: user.ID, PassHash: passHash}
		if err := tx.Create(&passHashModel).Error; err != nil {
			return err
		}
		return nil
	})

	if err := q.repo.IsDuplicateKey(err); err != nil {
		return err
	}

	return err
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
	qr := q.repo.DB().Model(&model.User{}).Where("id = ?", id).Updates(userMap)

	if qr.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return qr.Error
}

func (q UserQueryBuilder) Delete(id uuid.UUID) error {
	qr := q.repo.DB().Unscoped().Model(&model.User{}).Where("id = ?", id).Delete(&model.User{})

	if qr.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return qr.Error
}
