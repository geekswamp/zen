package base

import "gorm.io/gorm"

type UpdateMap map[string]any

type Repository struct{ db *gorm.DB }

func NewRepo(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) DB() *gorm.DB {
	return r.db
}
