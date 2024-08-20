package logkit

import (
	"fmt"
	"strings"
)

const (
	red    = "\033[1;31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	reset  = "\033[0m"

	prefix_color = "\033[1;33m"
)

type Logger struct {
	prefix string
}

// ==== CONSTRUCTOR ====
func New(prefix string) *Logger {

	return &Logger{
		prefix: strings.ToUpper(prefix),
	}
}

//==== DEFAULT FUNCTIONS ====

func (l *Logger) Debug(input ...interface{}) {
	all := append([]interface{}{cyan, prefix_color, l.prefix, reset}, input...)
	fmt.Printf("%s[DEBUG] %s[%s] %s%v\n", all...)
}

func (l *Logger) Info(input ...interface{}) {
	all := append([]interface{}{green, prefix_color, l.prefix, reset}, input...)
	fmt.Printf("%s[INFO]  %s[%s] %s%v\n", all...)
}

func (l *Logger) Warn(input ...interface{}) {
	all := append([]interface{}{yellow, prefix_color, l.prefix, reset}, input...)
	fmt.Printf("%s[WARN]  %s[%s] %s%v\n", all...)
}

func (l *Logger) Error(input ...interface{}) {
	all := append([]interface{}{red, prefix_color, l.prefix, reset}, input...)
	fmt.Printf("%s[ERROR] %s[%s] %s%v\n", all...)
}

//==== FUNCTIONS WITH FORMATTING ====

func (l *Logger) Debugf(format string, values ...interface{}) {
	l.Debug(fmt.Sprintf(format, values...))
}

func (l *Logger) Infof(format string, values ...interface{}) {
	l.Info(fmt.Sprintf(format, values...))
}

func (l *Logger) Warnf(format string, values ...interface{}) {
	l.Warn(fmt.Sprintf(format, values...))
}

func (l *Logger) Errorf(format string, values ...interface{}) {
	l.Error(fmt.Sprintf(format, values...))
}
