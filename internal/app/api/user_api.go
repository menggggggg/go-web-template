package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/menggggggg/go-web-template/internal/app/schema/user"
	"github.com/menggggggg/go-web-template/internal/app/service"
	"github.com/menggggggg/go-web-template/pkg/errors"
)

var UserSet = wire.NewSet(wire.Struct(new(UserAPI), "*"))

// UserAPI ...
type UserAPI struct {
	UserSrv *service.UserSrv
}

// Create ...
func (a *UserAPI) Create(c *gin.Context) {
	request := user.CreateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.WrapWithBadRequest(err))
		return
	}

	resp, err := a.UserSrv.Create(c, &request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.WrapWithInternalServerError(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetByName ...
func (a *UserAPI) GetByName(c *gin.Context) {
	request := user.GetByNameRequest{}
	if err := c.ShouldBindQuery(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.WrapWithBadRequest(err))
		return
	}

	resp, err := a.UserSrv.GetByName(c, &request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.WrapWithInternalServerError(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Get ...
func (a *UserAPI) Get(c *gin.Context) {
	request := user.GetRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.WrapWithBadRequest(err))
		return
	}

	resp, err := a.UserSrv.Get(c, &request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.WrapWithInternalServerError(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Update ...
func (a *UserAPI) Update(c *gin.Context) {
	request := user.UpdateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.WrapWithBadRequest(err))
		return
	}

	resp, err := a.UserSrv.Update(c, &request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.WrapWithInternalServerError(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Delete ...
func (a *UserAPI) Delete(c *gin.Context) {
	request := user.DeleteRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.WrapWithBadRequest(err))
		return
	}

	if err := a.UserSrv.Delete(c, &request); err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.WrapWithInternalServerError(err))
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
