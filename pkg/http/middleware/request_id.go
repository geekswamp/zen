package middleware

import (
	"github.com/geekswamp/zen/internal/core"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const _HeaderXRequestIDKey = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := core.NewContext(ctx)
		reqID := ctx.Request.Header.Get(_HeaderXRequestIDKey)
		ID := uuid.New()

		if reqID == "" {

			ctx.Request.Header.Set(_HeaderXRequestIDKey, ID.String())
			c.SetRequestID(ID)
		}
		ctx.Writer.Header().Set(_HeaderXRequestIDKey, ID.String())
		ctx.Next()
	}
}
