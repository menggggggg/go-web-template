package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/menggggggg/go-web-template/internal/app/api"
)

var _ IRouter = (*Router)(nil)

// RouterSet 注入router
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

// IRouter 注册路由
type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

// Router 路由管理器
type Router struct {
	UserAPI *api.UserAPI
} // end

// Register 注册路由
func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

// Prefixes 路由前缀列表
func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")

	v1 := g.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("", a.UserAPI.Create)
			user.GET("", a.UserAPI.GetByName)
			user.GET("/:userId", a.UserAPI.Get)
			user.PUT("/:userId", a.UserAPI.Update)
			user.DELETE("/:userId", a.UserAPI.Delete)
		}
	} // v1 end
}
