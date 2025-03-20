package core

import "github.com/google/uuid"

// Context holds request-scoped data for use throughout the application.
type Context struct {
	RequestID string
}

// GetRequestID generates and returns a new UUID as a string to uniquely identify a request.
// This method ensures each request has a unique identifier for tracking and logging purposes.
func (c *Context) GetRequestID() string {
	c.RequestID = uuid.NewString()

	return c.RequestID
}
