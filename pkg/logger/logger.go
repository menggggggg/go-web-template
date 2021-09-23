package logger

import (
	"context"
	"io"

	"github.com/sirupsen/logrus"
)

var (
	Trace = logrus.Trace
	Debug = logrus.Debug
	Info  = logrus.Info
	Warn  = logrus.Warn
	Error = logrus.Error
	Fatal = logrus.Fatal
	Panic = logrus.Panic
	Print = logrus.Print

	Tracef = logrus.Tracef
	Debugf = logrus.Debugf
	Infof  = logrus.Infof
	Warnf  = logrus.Warnf
	Errorf = logrus.Errorf
	Fatalf = logrus.Fatalf
	Panicf = logrus.Panicf
	Printf = logrus.Printf
)

type Logger = logrus.Logger

type Entry = logrus.Entry

type Fields = logrus.Fields

func SetLevel(level string) {
	le, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
		return
	}
	logrus.SetLevel(le)
}

func SetFormat(format string) {
	switch format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	}
}

func SetOutput(out io.Writer) {
	logrus.SetOutput(out)
}

func SetReportCaller(caller bool) {
	logrus.SetReportCaller(caller)
}

func WithContext(ctx context.Context) *Entry {
	return logrus.WithContext(ctx).WithField("X-Request-Id", ctx.Value("X-Request-Id"))
}

func WithError(err error) *Entry {
	return logrus.WithError(err)
}

func WithFields(fields Fields) *Entry {
	return logrus.WithFields(fields)
}
