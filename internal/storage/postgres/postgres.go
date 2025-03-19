package postgres

import (
	"github.com/geekswamp/zen/internal/logger"
	"gorm.io/gorm"
)

var log = logger.New()

type Postgres interface {
	Connect() (db *gorm.DB, err error)
}
