package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/menggggggg/go-web-template/pkg/util"
)

func VersionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/" || c.Request.URL.Path == "version" {
			c.Data(http.StatusOK, "application/json; charset=utf-8", util.VersionInfo())
			return
		}
		c.Next()
	}
}
