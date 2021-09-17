package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/menggggggg/go-web-template/pkg/logger"
)

var UserSet = wire.NewSet(wire.Struct(new(UserAPI), "*"))

// UserAPI ...
type UserAPI struct {
	//UserSrv *service.UserSrv
}

// Get ...
func (a *UserAPI) Get(c *gin.Context) {
	//	ctx := c.Request.Context()
	logger.WithContext(c).Info("get")
	c.JSON(200, "ok")
}
