package logrot

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Writer struct {
	mu sync.Mutex
	*lumberjack.Logger
	LogDirectory string
	LogName      string
	LogExtension string
	shouldRotate bool
}

type LogFile struct {
	Index int
	Name  string
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	// Inherit
	n, err = w.Logger.Write(p)
	if err != nil {
		fmt.Println("[Logger.Write]:", err)
		return n, err
	}
	// Get Files
	files, err := os.ReadDir(w.LogDirectory)
	if err != nil {
		fmt.Println("[ReadDir]:", err)
		return n, err
	}
	// Filter
	var list = make([]LogFile, 0)
	for _, file := range files {
		// Extension match
		extIndex := strings.Index(file.Name(), w.LogExtension) // .log
		if extIndex != -1 {
			name := file.Name()[:extIndex]
			logNameWithHyphen := w.LogName + "-" // logname-
			logNameWithPeriod := w.LogName + "." // logname.
			logNameIndexWithHyphen := strings.Index(name, logNameWithHyphen)
			logNameIndexWithPeriod := strings.Index(name, logNameWithPeriod)
			switch {
			case logNameIndexWithHyphen != -1:
				// logname-time.log
				surfix := file.Name()[logNameIndexWithHyphen+len(logNameWithHyphen) : extIndex]
				_, err := time.Parse("2006-01-02T15-04-05.000", surfix)
				if err == nil {
					// List append
					list = append(list, LogFile{Name: file.Name(), Index: 0})
					w.shouldRotate = true
				}
			case logNameIndexWithPeriod != -1:
				// logname.1.log
				surfix := file.Name()[logNameIndexWithPeriod+len(logNameWithPeriod) : extIndex]
				index, err := strconv.Atoi(surfix)
				if err == nil {
					// List append
					list = append(list, LogFile{Name: file.Name(), Index: index})
				}
			}
		}
	}
	// Rotate when it has index zero file
	if w.shouldRotate {
		// Selection Sort
		for i := range list {
			for j := i + 1; j < len(list); j++ {
				// Exchange
				if list[i].Index < list[j].Index {
					list[i], list[j] = list[j], list[i]
				}
			}
		}
		// Rename
		for _, v := range list {
			newIndex := v.Index + 1
			if newIndex > w.MaxBackups {
				// Remove out of limit
				os.Remove(w.LogDirectory + v.Name)
			} else {
				// Rename
				err := os.Rename(w.LogDirectory+v.Name, w.LogDirectory+w.LogName+"."+strconv.Itoa(newIndex)+w.LogExtension)
				if err != nil {
					log.Println("[Log rename error]:", w.LogDirectory+v.Name)
				}
			}
		}
		w.shouldRotate = false
	}
	return n, err
}
