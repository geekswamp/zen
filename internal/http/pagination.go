package http

import "math"

// SortOption represents a string-based option for sorting operations.
// It is used to specify the sorting direction in query parameters.
type SortOption string

const (
	// Asc represents ascending sort order option for query parameters.
	Asc SortOption = "asc"

	// Desc represents descending sort order option for query paramters.
	Desc SortOption = "desc"
)

// Pagination defines the query parameters structure for pagination and sorting.
type Pagination struct {
	Page   int64      `form:"page" validate:"omitempty"`
	Limit  int64      `form:"limit" validate:"omitempty"`
	Sort   SortOption `form:"sort" validate:"omitempty"`
	Search string     `form:"search" validate:"omitempty"`
}

// GetPage returns the current page number from the query parameters.
func (p *Pagination) GetPage() int64 {
	if p.Page == 0 {
		p.Page = 1
	}

	return p.Page
}

// GetLimit returns the query limit value.
func (p *Pagination) GetLimit() int64 {
	if p.Limit == 0 {
		p.Limit = 10
	}

	return p.Limit
}

// GetOffset calculates and returns the offset value for pagination.
func (p *Pagination) GetOffset() int64 {
	return (p.GetPage() - 1) * p.GetLimit()
}

// GetSort returns the sort options for the query parameters
func (p *Pagination) GetSort() SortOption {
	return p.Sort
}

// GetTotalPages calculates the total number of pages based on the total number of items and the limit per page.
func (p *Pagination) GetTotalPages(totalItems int64) int64 {
	return int64(math.Ceil(float64(totalItems) / float64(p.Limit)))
}

// GetHasReachedMax checks if the current page has reached the maximum number of pages.
func (p *Pagination) GetHasReachedMax(totalItems int64) bool {
	return p.GetPage() >= p.GetTotalPages(totalItems)
}
