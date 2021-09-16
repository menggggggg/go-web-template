package app

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/menggggggg/go-web-template/internal/app/config"
	"github.com/menggggggg/go-web-template/internal/app/middleware"
	"github.com/menggggggg/go-web-template/internal/app/router"
	"github.com/menggggggg/go-web-template/pkg/logger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitGinEngine(r router.IRouter) *gin.Engine {
	gin.SetMode(config.C.RunMode)

	app := gin.Default()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		logger.WithFields(logger.Fields{"method": httpMethod, "path": absolutePath, "handerName": handlerName}).Debug("http")
	}

	// 中间件配置
	// Version
	app.Use(middleware.VersionMiddleware())

	// RequestId
	app.Use(middleware.RequestIdMiddleware())

	// CORS
	if config.C.CORS.Enable {
		app.Use(middleware.CORSMiddleware())
	}

	// GZIP
	if config.C.GZIP.Enable {
		app.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// Router register
	r.Register(app)

	// Swagger
	if config.C.Swagger {
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return app
}
