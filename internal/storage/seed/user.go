package seed

import (
	"time"

	"github.com/geekswamp/zen/internal/base"
	"github.com/geekswamp/zen/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userSeeder struct{}

func init() {
	RegisterSeeder(userSeeder{})
}

func (s userSeeder) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.User{}, &model.UserPassHash{})
}

func (s userSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.User{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	log.Info("Seeding user data")
	userID := uuid.New()
	user := model.User{
		FullName:      "John Doe",
		Email:         "john@doe.com",
		Gender:        model.Male,
		Active:        true,
		ActivatedTime: time.Now().UnixMilli(),
		Phone:         "+6281234567890",
		Model: base.Model{
			ID:          userID,
			CreatedTime: time.Now().UnixMilli(),
		},
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	pass := model.UserPassHash{
		UserID:   userID,
		PassHash: "$argon2id$v=19$m=12288,t=3,p=1$CiI6u9dtw8jTTkjFCfqxVD3JX6n4kNoF+sWUMwp9z28$qc2gq3BDYWFQb4v24cjnJDNgQuI9eUrOUxjvL9wUezw",
	}

	if err := db.Create(&pass).Error; err != nil {
		return err
	}

	return nil
}
