package util

import (
	"log"
	"os"
	"sync"
)

type DefaultLogger struct {
	l *log.Logger
}

var (
	defaultLog         *DefaultLogger
	defaultLogInitOnce sync.Once
)

func defaultLogger() Logger {
	defaultLogInitOnce.Do(func() {
		defaultLog = &DefaultLogger{l: log.New(os.Stdout, "logger: ", log.Lshortfile)}
	})
	return defaultLog
}

func (l *DefaultLogger) Debugf(format string, v ...interface{}) {
	l.l.Printf("Debug "+format, v...)
}
func (l *DefaultLogger) Infof(format string, v ...interface{}) {
	l.l.Printf("Info "+format, v...)
}
func (l *DefaultLogger) Warnf(format string, v ...interface{}) {
	l.l.Printf("Warn "+format, v...)
}
func (l *DefaultLogger) Errorf(format string, v ...interface{}) {
	l.l.Printf("Error "+format, v...)
}
