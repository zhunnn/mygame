package logrot

import (
	"fmt"
	"mygame/config/enum"
	"testing"
)

func TestLogLevel(t *testing.T) {
	Log.Trace("Trace")
	Log.Debug("Debug")
	Log.Info("Info")
	Log.Warn("Warn")
	Log.Error("Error")
}

func TestNewLog(t *testing.T) {
	mylog := New(LogConfig{
		Environment:     enum.Environment_Local,
		Name:            "mylog",
		Directory:       "./logs/",
		TimestampFormat: "2006-01-02 15:04:05.000",
		MessageFormat:   fmt.Sprintf("[%v] [%v] %v - %v", FORMAT_STRING_TIME, FORMAT_STRING_LEVEL, FORMAT_STRING_NAME, FORMAT_STRING_MESSAGE),
		Extension:       ".log",
		LocalTime:       true,
		MaxSize:         5,
		MaxBackups:      5,
		MaxAge:          10,
		Compress:        false,
		// ReportCaller:    true,
	})
	mylog.Error("Error")
}
