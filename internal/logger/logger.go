package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger use for logging
type Logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

// NewLogger creates a new logger
func NewLogger(domain string) *Logger {
	infoPrefix := fmt.Sprintf("Info [%s]:", domain)
	warnPrefix := fmt.Sprintf("Warn [%s]:", domain)
	errorPrefix := fmt.Sprintf("Error [%s]:", domain)

	infoLogger := log.New(os.Stdout, infoPrefix, log.Ldate|log.Ltime)
	warnLogger := log.New(os.Stdout, warnPrefix, log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stdout, errorPrefix, log.Ldate|log.Ltime)

	return &Logger{infoLogger: infoLogger, warnLogger: warnLogger, errorLogger: errorLogger}
}

// Info prints an info message to the logs
func (l *Logger) Info(msg string) {
	l.infoLogger.Println(msg)
}

// Warn prints an warn message to the logs
func (l *Logger) Warn(msg string) {
	l.warnLogger.Println(msg)
}

// Error prints an error message to the logs
func (l *Logger) Error(msg string) {
	l.infoLogger.Println(msg)
}
