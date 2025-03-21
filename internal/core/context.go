package core

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	_RequestIDKey   string = "__request_id__"
	_UserSessionKey string = "__user_session__"
)

// Context holds request-scoped data for use throughout the application.
type Context struct {
	ctx *gin.Context
}

// NewContext create a new context instance.
func NewContext(ctx *gin.Context) Context {
	return Context{ctx: ctx}
}

// SetRequestID store the request id in the context.
func (c *Context) SetRequestID(id uuid.UUID) {
	c.ctx.Set(_RequestIDKey, id.String())
}

// GetRequestID generates and returns a new UUID as a string to uniquely identify a request.
func (c *Context) GetRequestID() *string {
	u, ok := c.ctx.Get(_RequestIDKey)
	if u == nil || !ok {
		return nil
	}

	if requestID, valid := u.(string); valid {
		return &([]string{requestID}[0])
	}

	return nil
}

// SetUserSession store the user session in the context.
func (c *Context) SetUserSession(u UserSession) {
	c.ctx.Set(_UserSessionKey, u)
}

// GetUserSession retrieves the user session from the context.
func (c *Context) GetUserSession() *UserSession {
	u, ok := c.ctx.Get(_UserSessionKey)
	if u == nil || !ok {
		return nil
	}

	if userSession, valid := u.(UserSession); valid {
		return &([]UserSession{userSession}[0])
	}

	return nil
}
