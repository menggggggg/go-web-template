package app

import (
	"context"

	"github.com/menggggggg/go-web-template/internal/app/config"
)

func Init() (func(), error) {
	rootCtx := context.Background()

	// 初始化配置文件
	config.LoadConfig()

	// 初始化日志
	InitLogger()

	// 初始化服务运行监控
	monitorCleanFunc := InitMonitor(rootCtx)

	return func() {
		monitorCleanFunc()
	}, nil
}
