package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/menggggggg/go-web-template/pkg/util"
)

type RequestIDCtx struct{}

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.GetHeader("X-Request-Id")
		if rid == "" {
			rid = util.NewRequestID()
		}

		c.Request = c.Request.WithContext(context.WithValue(c, RequestIDCtx{}, rid))
		c.Writer.Header().Set("X-Request-Id", rid)

		c.Next()
	}
}
