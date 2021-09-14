package app

import (
	"os"

	"github.com/menggggggg/go-web-template/internal/app/config"
	"github.com/menggggggg/go-web-template/pkg/logger"
)

func InitLogger() {
	c := config.C.Log
	logger.SetLevel(c.Level)
	logger.SetFormat(c.Format)
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
}
