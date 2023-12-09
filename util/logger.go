package util

import (
	"sync"
)

type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

var (
	logs      Logger
	logsMutex sync.Mutex
)

func SetLogger(l Logger) {
	logsMutex.Lock()
	defer logsMutex.Unlock()
	logs = l
}

func Logs() Logger {
	if logs == nil {
		return defaultLogger()
	}
	return logs
}
