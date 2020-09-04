package log

import (
	"fmt"
	stdlog "log"
	"os"
	"time"
)

const (
	debugPrefix = "DEBUG"
	infoPrefix  = "INFO"
	warnPrefix  = "WARN"
	errPrefix   = "ERROR"
	fatalPrefix = "FATAL ERROR"
)

type Logger struct {
	stdout *stdlog.Logger
	stderr *stdlog.Logger
	debug  bool
}

func NewLogger(stdout *stdlog.Logger, stderr *stdlog.Logger, debug bool) *Logger {
	return &Logger{stdout, stderr, debug}
}

var (
	stdout = stdlog.New(os.Stdout, "", 0)
	stderr = stdlog.New(os.Stdout, "", 0)
	debug  = true

	defaultLogger = NewLogger(stdout, stderr, debug)
)

func (l *Logger) Debug(text string) {
	if l.debug {
		l.output(newMessage(debugPrefix, text), l.stdout)
	}
}

// Info logs given text as information level
func (l *Logger) Info(text string) {
	l.output(newMessage(infoPrefix, text), l.stdout)
}

// Warn logs given text as warning level
func (l *Logger) Warn(text string) {
	l.output(newMessage(warnPrefix, text), l.stderr)
}

// Error logs given text as error level
func (l *Logger) Error(text string) {
	l.output(newMessage(errPrefix, text), l.stderr)
}

// Fatal logs given text as fatal error level and calls os.Exit(1)
func (l *Logger) Fatal(text string) {
	l.output(newMessage(fatalPrefix, text), l.stderr)
	os.Exit(1)
}

func (l *Logger) output(m message, logger *stdlog.Logger) {
	var text string
	if l.debug {
		text = fmt.Sprintf("%v [%s] %s", m.Timestamp.Format("2006-01-02 15:04:05"), m.Prefix, m.Message)
	} else {
		text = fmt.Sprintf("time:%v\tlevel:%s\tmessage:%s", m.Timestamp.Format("2006-01-02 15:04:05"), m.Prefix, m.Message)
	}

	logger.Println(text)
}

type message struct {
	Timestamp time.Time
	Prefix    string
	Message   string
}

// newMessage creates message struct
func newMessage(prefix string, text string) message {
	return message{
		Timestamp: time.Now(),
		Prefix:    prefix,
		Message:   text,
	}
}

// SetDebug sets flag whether it is debug mode or not
func SetDefaultLoggerDebugMode(mode bool) {
	defaultLogger.debug = mode
}

// Debug logs debug message when it is debug mode
func Debug(text string) {
	defaultLogger.Debug(text)
}

func Debugf(format string, v ...interface{}) {
	defaultLogger.Debug(fmt.Sprintf(format, v...))
}

// Info logs given text as information level
func Info(text string) {
	defaultLogger.Info(text)
}

func Infof(format string, v ...interface{}) {
	defaultLogger.Info(fmt.Sprintf(format, v...))
}

// Warn logs given text as warning level
func Warn(text string) {
	defaultLogger.Warn(text)
}

func Warnf(format string, v ...interface{}) {
	defaultLogger.Warn(fmt.Sprintf(format, v...))
}

// Error logs given text as error level
func Error(text string) {
	defaultLogger.Error(text)
}

func Errorf(format string, v ...interface{}) {
	defaultLogger.Error(fmt.Sprintf(format, v...))
}

// Fatal logs given text as fatal error level and calls os.Exit(1)
func Fatal(text string) {
	defaultLogger.Fatal(text)
}

func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatal(fmt.Sprintf(format, v...))
}

// output calls Println to output log message to standard.
// If it is not debug mode log format will be ltsv.
func output(m message, logger *stdlog.Logger) {
	var text string
	if debug {
		text = fmt.Sprintf("%v [%s] %s", m.Timestamp.Format("2006-01-02 15:04:05"), m.Prefix, m.Message)
	} else {
		text = fmt.Sprintf("time:%v\tlevel:%s\tmessage:%s", m.Timestamp.Format("2006-01-02 15:04:05"), m.Prefix, m.Message)
	}

	logger.Println(text)
}
