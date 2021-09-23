package app

import (
	"github.com/menggggggg/go-web-template/internal/app/config"
	"github.com/menggggggg/go-web-template/pkg/logger"
	"github.com/menggggggg/go-web-template/pkg/mysql"
	"gorm.io/gorm"
)

func InitGormDB() *gorm.DB {
	cfg := config.C.MySQL
	db, err := mysql.New(&mysql.Config{
		DSN:          cfg.DSN(),
		MaxIdleConns: cfg.MaxIdleConns,
		MaxLifetime:  cfg.MaxLifetime,
		MaxOpenConns: cfg.MaxOpenConns,
	})
	if err != nil {
		logger.Panicf("init grom db fail, dsn: %s error: %s", cfg.DSN(), err.Error())
	}
	return db
}
