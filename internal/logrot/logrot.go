package logrot

import (
	"fmt"
	"io"
	"mygame/config"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	Environment_local        = "local"
	Default_Timestamp_Format = "2006-01-02 15:04:05.000"
	Default_Log_Extension    = ".log"
)

var (
	defaultLogConfig = LogConfig{
		Environment:     config.Config.System.Environment,
		Name:            config.Config.System.ServiceName,
		Directory:       config.Config.System.ProjectRootPath + "/logs/",
		TimestampFormat: Default_Timestamp_Format,
		MessageFormat:   fmt.Sprintf("[%v] [%v] %v - %v", FORMAT_STRING_TIME, FORMAT_STRING_LEVEL, FORMAT_STRING_NAME, FORMAT_STRING_MESSAGE),
		Extension:       Default_Log_Extension,
		LocalTime:       true,
		MaxSize:         5,
		MaxBackups:      5,
		MaxAge:          10,
		Compress:        false,
		ReportCaller:    true,
	}
	// Example: logrot.Log.Errorln("err:", err)
	Log = New(defaultLogConfig)
)

type LogConfig struct {
	Environment     string
	Name            string
	Directory       string
	TimestampFormat string
	MessageFormat   string
	Extension       string
	LocalTime       bool
	MaxSize         int // MB
	MaxBackups      int // Num
	MaxAge          int // Day
	Compress        bool
	ReportCaller    bool
}

type logrot struct {
	*logrus.Logger
	MultiWriter io.Writer
}

func New(config LogConfig) *logrot {
	// Console log:
	// Instance
	log := &logrot{Logger: logrus.New()}
	// Level
	log.SetLevel(logrus.TraceLevel)
	// Caller
	if config.ReportCaller {
		log.SetReportCaller(true)
	}
	// Formatter
	switch config.Environment {
	case Environment_local:
		// Local log with color
		log.SetFormatter(NewFormatter(config.Name, config.TimestampFormat, config.MessageFormat, true))
	default:
		// Other Environment log without color
		log.SetFormatter(NewFormatter(config.Name, config.TimestampFormat, config.MessageFormat, false))
	}
	// Output
	log.SetOutput(os.Stdout)
	// File log:
	// Writer
	writer := &Writer{
		LogDirectory: config.Directory,
		LogName:      config.Name,
		LogExtension: config.Extension,
		Logger: &lumberjack.Logger{
			Filename:   config.Directory + config.Name + config.Extension,
			LocalTime:  config.LocalTime,
			MaxSize:    config.MaxSize,    // MB
			MaxBackups: config.MaxBackups, // Num
			MaxAge:     config.MaxAge,     // Days
			Compress:   config.Compress,
		},
	}
	log.MultiWriter = io.MultiWriter(os.Stdout, writer)
	// Hook
	log.AddHook(&RotateFileHook{
		logWriter: writer,
		Level:     logrus.TraceLevel,
		// Log file without color
		Formatter: NewFormatter(config.Name, config.TimestampFormat, config.MessageFormat, false),
	})
	return log
}
