package app

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/menggggggg/go-web-template/internal/app/config"
	"github.com/menggggggg/go-web-template/pkg/logger"
)

func Init(ctx context.Context) (func(), error) {
	// 初始化配置文件
	config.LoadConfig()

	// 初始化日志
	InitLogger()

	// 初始化服务运行监控
	monitorCleanFunc := InitMonitor(ctx)

	// 初始化依赖注入器
	injector := BuildInjector()

	// 初始化HTTP服务
	httpServerCleanFunc := InitHTTPServer(ctx, injector.Engine)

	return func() {
		httpServerCleanFunc()
		monitorCleanFunc()
	}, nil
}

// InitHTTPServer 初始化http服务
func InitHTTPServer(ctx context.Context, handler http.Handler) func() {
	cfg := config.C.HTTP
	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		logger.WithContext(ctx).Printf("HTTP server is running at %s.", cfg.Addr)

		var err error
		if cfg.CertFile != "" && cfg.KeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}

	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.WithContext(ctx).Errorf(err.Error())
		}
	}
}

// Run 运行服务
func Run(ctx context.Context) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		logger.WithContext(ctx).Infof("Receive signal[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	logger.WithContext(ctx).Infof("Server exit")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
