package logrot

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type Formatter struct {
	LogName         string
	TimestampFormat string
	MessageFormat   string
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 呼叫者
	var caller string
	if entry.HasCaller() {
		caller = fmt.Sprintf("%v:%v", entry.Caller.File, entry.Caller.Line)
	}

	// 日誌等級
	var level string
	upperLevel := strings.ToUpper(entry.Level.String())
	if len(upperLevel) > 4 {
		level = upperLevel[:4]
	} else {
		level = upperLevel
	}

	// 時間
	var time string
	time = entry.Time.Format(f.TimestampFormat)

	// 格式化
	msg := f.MessageFormat
	msg = strings.Replace(msg, "<time>", time, 1)
	msg = strings.Replace(msg, "<level>", level, 1)
	msg = strings.Replace(msg, "<name>", f.LogName, 1)
	msg = strings.Replace(msg, "<caller>", caller, 1)
	msg = strings.Replace(msg, "<message>", entry.Message, 1)
	msg += "\n"

	return []byte(msg), nil
}
