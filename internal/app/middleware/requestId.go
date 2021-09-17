package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/menggggggg/go-web-template/pkg/util"
)

var RequestId = "X-Request-Id"

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.GetHeader(RequestId)
		if rid == "" {
			rid = util.NewRequestID()
		}

		c.Set(RequestId, rid)
		c.Writer.Header().Set(RequestId, rid)

		c.Next()
	}
}
