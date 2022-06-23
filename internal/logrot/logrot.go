package logrot

import (
	"io"
	"mygame/config"
	"mygame/config/enum"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	default_TimeStampFormat            = "2006-01-02 15:04:05.000"
	default_LogMessageFormat           = "[<time>] [<level>] <name> - <message>"
	default_LogMessageFormatWithCaller = "[<time>] [<level>] <name> <caller> - <message>"
)

func init() {
	// Logger
	logrus.SetLevel(logrus.TraceLevel)
	// 是否顯示呼叫函式
	logrus.SetReportCaller(true)
	// SetFormatter
	switch config.Config.System.Environment {
	case enum.Environment_Local:
		logrus.SetFormatter(
			&Formatter{
				LogName:         config.Config.System.ServiceName,
				TimestampFormat: default_TimeStampFormat,
				MessageFormat:   default_LogMessageFormatWithCaller,
			},
		)
	default:
		logrus.SetFormatter(
			&Formatter{
				LogName:         config.Config.System.ServiceName,
				TimestampFormat: default_TimeStampFormat,
				MessageFormat:   default_LogMessageFormat,
			},
		)
	}
	// Rotate
	rotation := &Rotation{
		LogDirectory: config.Config.System.ProjectRootPath + "/logs/",
		LogName:      config.Config.System.ServiceName,
		LogExtension: ".log",
		Logger: lumberjack.Logger{
			Filename:   config.Config.System.ProjectRootPath + "/logs/" + config.Config.System.ServiceName + ".log",
			LocalTime:  true,
			MaxSize:    1,  // MB
			MaxBackups: 3,  // Number
			MaxAge:     28, // Days
			Compress:   false,
		},
	}
	mw := io.MultiWriter(os.Stdout, rotation)
	logrus.SetOutput(mw)
}
