package logrot

import (
	"io"

	"github.com/sirupsen/logrus"
)

type RotateFileHook struct {
	logWriter io.Writer
	Level     logrus.Level
	Formatter logrus.Formatter
}

func (hook *RotateFileHook) Levels() []logrus.Level {
	return logrus.AllLevels[:hook.Level+1]
}

func (hook *RotateFileHook) Fire(entry *logrus.Entry) (err error) {
	b, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.logWriter.Write(b)
	if err != nil {
		return err
	}
	return nil
}
