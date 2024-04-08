package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

type LogLevel int

const (
	Info LogLevel = iota
	Warn
	Error
)

type Logger struct {
	*log.Logger
	level LogLevel
}

func NewLogger(prefix string, flag int, level LogLevel, logFile *os.File) *Logger {
	logger := log.New(os.Stdout, prefix, flag)
	if logFile != nil {
		logger.SetOutput(logFile)
	}

	return &Logger{
		Logger: logger,
		level:  level,
	}
}

func (l *Logger) logf(level LogLevel, format string, v ...interface{}) {
	if level >= l.level {
		_, file, line, ok := runtime.Caller(2) // 2 表示获取调用者的信息
		if ok {
			format = fmt.Sprintf("[%s] %s:%d %s", levelString(level), file, line, format)
			l.Printf(format, v...)
		}
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.logf(Info, format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.logf(Warn, format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logf(Error, format, v...)
}

func (l *Logger) logln(level LogLevel, v ...interface{}) {
	if level >= l.level {
		_, file, line, ok := runtime.Caller(2) // 2 表示获取调用者的信息
		if ok {
			v = append([]interface{}{fmt.Sprintf("[%s] %s:%d", levelString(level), file, line)}, v...)
		}
		l.Println(v...)
	}
}

func (l *Logger) Infoln(v ...interface{}) {
	l.logln(Info, v...)
}

func (l *Logger) Warnln(v ...interface{}) {
	l.logln(Warn, v...)
}

func (l *Logger) Errorln(v ...interface{}) {
	l.logln(Error, v...)
}

func levelString(level LogLevel) string {
	switch level {
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
