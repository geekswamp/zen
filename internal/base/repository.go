package base

import "gorm.io/gorm"

type Repository struct{ db *gorm.DB }

func New(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) DB() *gorm.DB {
	return r.db
}
