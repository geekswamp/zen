package seed

import (
	"github.com/geekswamp/zen/internal/logger"
	"gorm.io/gorm"
)

var log = logger.New()

type Seeder interface {
	Migrate(db *gorm.DB) error
	Seed(db *gorm.DB) error
}

var seeders []Seeder

func RegisterSeeder(s Seeder) {
	seeders = append(seeders, s)
}

func RunSeeders(db *gorm.DB) error {
	for _, s := range seeders {
		if err := s.Migrate(db); err != nil {
			return err
		}
		if err := s.Seed(db); err != nil {
			return err
		}
	}
	return nil
}
