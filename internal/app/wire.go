//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"
	"github.com/menggggggg/go-web-template/internal/app/api"
	"github.com/menggggggg/go-web-template/internal/app/dao"
	"github.com/menggggggg/go-web-template/internal/app/router"
	"github.com/menggggggg/go-web-template/internal/app/service"
)

// BuildInjector 生成注入器
func BuildInjector() *Injector {
	wire.Build(
		router.RouterSet,
		dao.DaoSet,
		service.ServiceSet,
		api.APISet,
		InitGinEngine,
		InitGormDB,
		InjectorSet,
	)
	return new(Injector)
}
