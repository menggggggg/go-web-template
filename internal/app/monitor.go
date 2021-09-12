package app

import (
	"context"

	"github.com/google/gops/agent"
	"github.com/menggggggg/go-web-template/internal/app/config"
)

// InitMonitor 初始化服务监控
func InitMonitor(ctx context.Context) func() {
	if c := config.C.Monitor; c.Enable {
		err := agent.Listen(agent.Options{Addr: c.Addr, ConfigDir: c.ConfigDir, ShutdownCleanup: false})
		if err != nil {
			//logger.WithContext(ctx).Errorf("Agent monitor error: %s", err.Error())
		}
		return func() {
			agent.Close()
		}
	}
	return func() {}
}
