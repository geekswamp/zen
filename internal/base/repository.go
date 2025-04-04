package base

import (
	"strings"

	"gorm.io/gorm"
)

// UpdateMap represents a map used for dynamic updates in repository operations.
// The string key represents the field name to be updated, and the any value
// represents the new value to be set for that field.
type UpdateMap map[string]any

// Repository represents a base data access layer struct that holds a GORM database connection.
// It serves as a foundation for specific repository implementations by providing access to the database instance.
type Repository struct {
	db *gorm.DB
}

// NewRepo creates a new instance of Repository with the provided gorm database connection.
func NewRepo(db *gorm.DB) Repository {
	return Repository{db: db}
}

// DB returns the GORM database instance used by the repository.
func (r Repository) DB() *gorm.DB {
	return r.db
}

// IsDuplicateKey checks if the given error is a duplicate key constraint violation in the database.
// It returns gorm.ErrDuplicatedKey if the error contains a duplicate key violation message,
// otherwise returns nil.
func (r Repository) IsDuplicateKey(err error) error {
	if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return gorm.ErrDuplicatedKey
	}

	return nil
}
