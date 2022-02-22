package ilog

import (
	"github.com/sirupsen/logrus"
	"os"
)

type formatType string

const (
	TEXT formatType = "text"
	JSON formatType = "json"
)

const (
	Panic = logrus.PanicLevel
	Fatal = logrus.FatalLevel
	Error = logrus.ErrorLevel
	Warn  = logrus.WarnLevel
	Info  = logrus.InfoLevel
	Debug = logrus.DebugLevel
	Trace = logrus.TraceLevel
)

var (
	Logger *logrus.Logger
)

//Init 初始化日志库
func Init(format formatType, report bool, level logrus.Level) {
	Logger = logrus.New()

	if format == TEXT {
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	} else {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	}

	Logger.SetReportCaller(report)

	Logger.SetLevel(level)
	Logger.SetOutput(os.Stdout)
}
