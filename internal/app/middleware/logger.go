package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/menggggggg/go-web-template/pkg/logger"
)

// LoggerMiddleware 请求日志记录
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		cost := time.Since(start)

		// 记录请求信息
		fields := make(logger.Fields)
		fields["requestId"] = c.Value(RequestId)

		if len(c.Errors) > 0 {
			fields["errors"] = c.Errors.Errors()
		}

		fields["status"] = c.Writer.Status()
		fields["method"] = c.Request.Method
		fields["path"] = path
		fields["query"] = query
		fields["ip"] = c.ClientIP()
		fields["user-agent"] = c.Request.UserAgent()
		fields["cost"] = cost

		// 记录请求信息
		logger.WithFields(fields).Info("access")
	}
}
