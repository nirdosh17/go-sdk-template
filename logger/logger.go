// Package logger provides the logging interface.
//
// Default logger can be overridden with a custom logger in the Config object.
package logger

import (
	"log"
	"os"
)

const (
	LevelInfo  = "INFO:"
	LevelError = "ERROR:"
)

// Logger is default log writer for the sdk. Custom loggers must satisfy this interface
type Logger interface {
	Log(...interface{})
}

// useful for the consumers to provide a logger function
type LoggerFunc func(...interface{})

func (l LoggerFunc) Log(args ...interface{}) {
	// runs provided logger function with all arguments
	l(args...)
}

// DefaultLogger is a minimal logger
type SimpleLogger struct {
	Logger *log.Logger
	Debug  bool
}

// NewDefaultLogger returns a minimal logger
func NewDefaultLogger() *SimpleLogger {
	return &SimpleLogger{
		// using empty prefix here
		Logger: log.New(os.Stdout, "", log.LstdFlags),
		Debug:  false,
	}
}

func (l *SimpleLogger) Log(args ...interface{}) {
	l.Logger.Println(args...)
}
