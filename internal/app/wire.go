//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"
	"github.com/menggggggg/go-web-template/internal/app/api"
	"github.com/menggggggg/go-web-template/internal/app/router"
)

// BuildInjector 生成注入器
func BuildInjector() *Injector {
	wire.Build(
		router.RouterSet,
		api.APISet,
		InitGinEngine,
		InjectorSet,
	)
	return new(Injector)
}
