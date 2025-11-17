package core

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = 1
	LogLevelInfo  LogLevel = 2
	LogLevelWarn  LogLevel = 3
	LogLevelError LogLevel = 4
)

type Logger interface {
	Debug(v ...any)
	Debugf(format string, v ...any)
	Info(v ...any)
	Infof(format string, v ...any)
	Warn(v ...any)
	Warnf(format string, v ...any)
	Error(v ...any)
	Errorf(format string, v ...any)
}

func newDefaultLogger() Logger {
	return &defaultLogger{loglevel: LogLevelInfo, logger: log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)}
}

type defaultLogger struct {
	loglevel LogLevel
	logger   *log.Logger
}

func (l defaultLogger) Debug(args ...any) {
	if l.loglevel <= LogLevelDebug {
		l.logger.SetPrefix("[Debug] ")
		l.logger.Print(args...)
	}
}

func (l defaultLogger) Debugf(format string, args ...any) {
	if l.loglevel <= LogLevelDebug {
		l.logger.SetPrefix("[Debug] ")
		l.logger.Printf(format, args...)
	}
}

func (l defaultLogger) Info(args ...any) {
	if l.loglevel <= LogLevelInfo {
		l.logger.SetPrefix("[Info] ")
		l.logger.Print(args...)
	}
}

func (l defaultLogger) Infof(format string, args ...any) {
	if l.loglevel <= LogLevelInfo {
		l.logger.SetPrefix("[Info] ")
		l.logger.Printf(format, args...)
	}
}

func (l defaultLogger) Warn(args ...any) {
	if l.loglevel <= LogLevelWarn {
		l.logger.SetPrefix("[Warn] ")
		l.logger.Print(args...)
	}
}

func (l defaultLogger) Warnf(format string, args ...any) {
	if l.loglevel <= LogLevelWarn {
		l.logger.SetPrefix("[Warn] ")
		l.logger.Printf(format, args...)
	}
}

func (l defaultLogger) Error(args ...any) {
	if l.loglevel <= LogLevelError {
		l.logger.SetPrefix("[Error] ")
		l.logger.Print(args...)
	}
}

func (l defaultLogger) Errorf(format string, args ...any) {
	if l.loglevel <= LogLevelError {
		l.logger.SetPrefix("[Error] ")
		l.logger.Printf(format, args...)
	}
}
