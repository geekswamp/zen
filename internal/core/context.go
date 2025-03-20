package core

import "github.com/google/uuid"

type Context struct {
	RequestID string
}

func (c *Context) GetRequestID() string {
	c.RequestID = uuid.NewString()

	return c.RequestID
}
