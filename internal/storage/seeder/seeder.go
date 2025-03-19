package seeder

import (
	"github.com/geekswamp/zen/internal/errors"
	"github.com/geekswamp/zen/internal/logger"
	"gorm.io/gorm"
)

var log = logger.New()

// SeederFunc represents a function type that performs database seeding operations.
type SeederFunc func(db *gorm.DB) error

// Manager handles a collection of seed functions for populating database tables.
type Manager struct {
	seeders []SeederFunc
}

// New creates and returns a new instance of Manager.
func New() *Manager {
	return &Manager{}
}

// Add appends one or more seeder functions to the Manager's collection of seeders.
func (s *Manager) Add(seeders ...SeederFunc) {
	s.seeders = append(s.seeders, seeders...)
}

// Run executes all registered seeders sequentially using the provided database connection.
func (s *Manager) Run(db *gorm.DB) error {
	for _, seed := range s.seeders {
		if err := seed(db); err != nil {
			log.Fatal(errors.ErrFailedRunningSeeder.Error(), logger.ErrDetails(err))
			return err
		}
	}

	return nil
}
