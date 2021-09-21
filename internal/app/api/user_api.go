package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/menggggggg/go-web-template/internal/app/schema"
	"github.com/menggggggg/go-web-template/pkg/logger"
)

var UserSet = wire.NewSet(wire.Struct(new(UserAPI), "*"))

// UserAPI ...
type UserAPI struct {
	//UserSrv *service.UserSrv
}

// Get ...
func (a *UserAPI) Get(c *gin.Context) {
	request := schema.UserGetRequest{}
	if c.ShouldBind(&request) != nil {
		c.AbortWithError(400, errors.New("param"))
		return
	}

	logger.WithContext(c).Info(request)
	c.JSON(200, "ok")
}
