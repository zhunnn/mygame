package logrot

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Rotation struct {
	mu sync.Mutex
	lumberjack.Logger
	LogDirectory string
	LogName      string
	LogExtension string
	ShouldRotate bool
}

type LogFile struct {
	Index int
	Name  string
}

func (r *Rotation) Write(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Inherit
	n, err = r.Logger.Write(p)
	// Get Files
	files, err := os.ReadDir(r.LogDirectory)
	// Filter
	var list = make([]LogFile, 0)
	for _, file := range files {
		// Extension match
		extIndex := strings.Index(file.Name(), r.LogExtension) // .log
		if extIndex != -1 {
			name := file.Name()[:extIndex]
			logNameWithHyphen := r.LogName + "-" // logname-
			logNameWithPeriod := r.LogName + "." // logname.
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
					r.ShouldRotate = true
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
	if r.ShouldRotate {
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
			if newIndex > r.MaxBackups {
				// Remove out of limit
				os.Remove(r.LogDirectory + v.Name)
			} else {
				// Rename
				err := os.Rename(r.LogDirectory+v.Name, r.LogDirectory+r.LogName+"."+strconv.Itoa(newIndex)+r.LogExtension)
				if err != nil {
					log.Println("[Log rename error]:", r.LogDirectory+v.Name)
				}
			}
		}
		r.ShouldRotate = false
	}
	return n, err
}
