package logrot

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	ANSI_COLOR_RESET   = "\x1b[0m"
	ANSI_COLOR_BLACK   = "\x1b[30m"
	ANSI_COLOR_RED     = "\x1b[31m"
	ANSI_COLOR_GREEN   = "\x1b[32m"
	ANSI_COLOR_YELLOW  = "\x1b[33m"
	ANSI_COLOR_BLUE    = "\x1b[34m"
	ANSI_COLOR_MAGENTA = "\x1b[35m"
	ANSI_COLOR_CYAN    = "\x1b[36m"
	ANSI_COLOR_WHITE   = "\x1b[37m"
)

const (
	FORMAT_STRING_TIME    = "<time>"
	FORMAT_STRING_LEVEL   = "<level>"
	FORMAT_STRING_NAME    = "<name>"
	FORMAT_STRING_MESSAGE = "<message>"
	FORMAT_STRING_CALLER  = "<caller>"
)

type Formatter struct {
	EnableColor     bool
	LogName         string
	TimestampFormat string
	MessageFormat   string
}

func NewFormatter(name string, timef string, msgf string, color bool) *Formatter {
	return &Formatter{
		EnableColor:     color,
		LogName:         name,
		TimestampFormat: timef,
		MessageFormat:   msgf,
	}
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	msg := f.MessageFormat
	// 呼叫者
	var caller string
	if entry.HasCaller() {
		msg = msg + " " + FORMAT_STRING_CALLER
		caller = fmt.Sprintf("%v:%v", entry.Caller.File, entry.Caller.Line)
	}

	// 顏色
	var color string
	color = f.getLevelColor(entry.Level)

	// 日誌等級
	var level string
	upperLevel := strings.ToUpper(entry.Level.String())
	level = upperLevel

	// 時間
	var time string
	time = entry.Time.Format(f.TimestampFormat)

	// 顏色
	if f.EnableColor {
		msgIndex := strings.Index(msg, FORMAT_STRING_MESSAGE)
		if msgIndex != 0 {
			msg = color + msg[:msgIndex] + ANSI_COLOR_RESET + msg[msgIndex:]
		}
		msg = strings.Replace(msg, FORMAT_STRING_CALLER, ANSI_COLOR_BLACK+FORMAT_STRING_CALLER+ANSI_COLOR_RESET, 1)
	}

	// 格式化替換
	msg = strings.Replace(msg, FORMAT_STRING_TIME, time, 1)
	msg = strings.Replace(msg, FORMAT_STRING_LEVEL, level, 1)
	msg = strings.Replace(msg, FORMAT_STRING_NAME, f.LogName, 1)
	msg = strings.Replace(msg, FORMAT_STRING_CALLER, caller, 1)
	msg = strings.Replace(msg, FORMAT_STRING_MESSAGE, entry.Message, 1)
	msg += "\n"

	return []byte(msg), nil
}

func (f *Formatter) getLevelColor(level logrus.Level) (ansiColor string) {
	switch level {
	case logrus.TraceLevel:
		ansiColor = ANSI_COLOR_WHITE
	case logrus.DebugLevel:
		ansiColor = ANSI_COLOR_CYAN
	case logrus.InfoLevel:
		ansiColor = ANSI_COLOR_GREEN
	case logrus.WarnLevel:
		ansiColor = ANSI_COLOR_YELLOW
	case logrus.ErrorLevel:
		ansiColor = ANSI_COLOR_RED
	case logrus.FatalLevel:
		ansiColor = ANSI_COLOR_BLUE
	case logrus.PanicLevel:
		ansiColor = ANSI_COLOR_BLUE
	default:
		ansiColor = ANSI_COLOR_WHITE
	}
	return ansiColor
}
