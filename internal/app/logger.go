package app

import (
	"os"

	"github.com/menggggggg/go-web-template/internal/app/config"
	logger "github.com/sirupsen/logrus"
)

func InitLogger() {
	c := config.C.Log

	level, err := logger.ParseLevel(c.Level)
	if err != nil {
		logger.Warn("parse config log error: ", err)
		logger.SetLevel(logger.InfoLevel)
		return
	}
	logger.SetLevel(level)

	switch c.Format {
	case "json":
		logger.SetFormatter(&logger.JSONFormatter{})
	default:
		logger.SetFormatter(&logger.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	}

	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
}
