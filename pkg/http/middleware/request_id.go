package middleware

import (
	"github.com/geekswamp/zen/internal/core"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const _HeaderXRequestIDKey = "X-Request-ID"

// RequestID is a Gin middleware function that ensures each request has a unique identifier.
// If the request already contains an X-Request-ID header, it uses that value.
// Otherwise, it generates a new UUID and sets it as the X-Request-ID header.
// The request ID is also added to the response headers and stored in the context.
func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := core.NewContext(ctx)
		reqID := ctx.Request.Header.Get(_HeaderXRequestIDKey)

		if reqID == "" {
			reqID = uuid.NewString()
		}

		ctx.Request.Header.Set(_HeaderXRequestIDKey, reqID)
		ctx.Writer.Header().Set(_HeaderXRequestIDKey, reqID)
		c.SetRequestID(uuid.MustParse(reqID))
		ctx.Next()
	}
}
