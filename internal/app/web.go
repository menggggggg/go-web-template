package app

import (
	"github.com/gin-gonic/gin"
	"github.com/menggggggg/go-web-template/internal/app/config"
	"github.com/menggggggg/go-web-template/internal/app/router"
)

func InitGinEngine(r router.IRouter) *gin.Engine {
	gin.SetMode(config.C.RunMode)

	app := gin.New()

	// CORS
	// if config.C.CORS.Enable {
	// 	app.Use(middleware.CORSMiddleware())
	// }

	// GZIP
	// if config.C.GZIP.Enable {
	// 	app.Use(gzip.Gzip(gzip.BestCompression,
	// 		gzip.WithExcludedExtensions(config.C.GZIP.ExcludedExtentions),
	// 		gzip.WithExcludedPaths(config.C.GZIP.ExcludedPaths),
	// 	))
	// }

	// Router register
	r.Register(app)

	// Swagger
	// if config.C.Swagger {
	// 	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// }
	return app
}
