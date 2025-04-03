package middleware

import (
	"github.com/geekswamp/zen/internal/core"
	"github.com/geekswamp/zen/internal/http"
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

		ID, err := uuid.Parse(reqID)
		if err != nil {
			// If the request ID is invalid, generate a new one and set it in the context and headers.
			// Also, return a 400 Bad Request response with an error message.
			ID = uuid.New()
			c.SetRequestID(ID)
			ctx.Request.Header.Set(_HeaderXRequestIDKey, ID.String())
			ctx.Writer.Header().Set(_HeaderXRequestIDKey, ID.String())
			http.New().BadRequest(ctx, http.Error{
				Code:   http.InvalidRequestID.Code(),
				Reason: http.InvalidRequestID.Detail(),
			})
			ctx.Abort()
			return
		}

		c.SetRequestID(ID)
		ctx.Request.Header.Set(_HeaderXRequestIDKey, ID.String())
		ctx.Writer.Header().Set(_HeaderXRequestIDKey, ID.String())
		ctx.Next()
	}
}
