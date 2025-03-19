package seeder

import (
	"github.com/geekswamp/zen/internal/logger"
	"github.com/geekswamp/zen/internal/pkg/errors"
	"gorm.io/gorm"
)

var log = logger.New()

// SeederFunc represents a function type that performs database seeding operations.
type SeederFunc func(db *gorm.DB) error

// SeederManager handles a collection of seed functions for populating database tables.
type SeederManager struct {
	seeders []SeederFunc
}

// New creates and returns a new instance of SeederManager.
func New() *SeederManager {
	return &SeederManager{}
}

// Add appends one or more seeder functions to the SeederManager's collection of seeders.
func (s *SeederManager) Add(seeders ...SeederFunc) {
	s.seeders = append(s.seeders, seeders...)
}

// Run executes all registered seeders sequentially using the provided database connection.
func (s *SeederManager) Run(db *gorm.DB) error {
	for _, seed := range s.seeders {
		if err := seed(db); err != nil {
			log.Fatal(errors.ErrFailedRunningSeeder.Error(), logger.ErrDetails(err))
			return err
		}
	}

	return nil
}
