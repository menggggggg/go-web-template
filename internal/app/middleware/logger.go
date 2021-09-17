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
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)

		requestLogger := logger.Entry{}

		// 记录业务信息
		requestLogger.WithFields(logger.Fields{
			"requestId": c.Value(RequestId),
		})

		// 记录请求信息
		requestLogger.WithFields(logger.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       path,
			"query":      query,
			"ip":         c.ClientIP(),
			"user-agent": c.Request.UserAgent(),
			"errors":     c.Errors.ByType(gin.ErrorTypePrivate).String(),
			"cost":       cost,
		}).Info(path)
	}
}
