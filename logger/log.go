package logger

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// LogEntry is extended from logrus
type LogEntry = logrus.Entry

// LogFields is extended from logrus
type LogFields = logrus.Fields

// Log is a custom struct extending *LogEntry
type Log struct {
	entry *LogEntry // only extend from this package
	// *LogEntry // extend logrus to caller
}

// ForPackage returns a new package-level logger
func ForPackage(packageID string) *Log {
	return &Log{entry: logrus.WithFields(LogFields{
		"packageID": packageID,
	})}
}

// ForUnknownPackage returns a new logger for unknown log
func ForUnknownPackage() *Log {
	return ForPackage("unknown")
}

// ----------------------

// getEntry is an internal utility func
func (l *Log) getEntry() *LogEntry {
	if l.entry == nil {
		l = ForUnknownPackage()
	}
	return l.entry
}

// ----------------------

// WithFunc returns a logger with a funcTag
func (l *Log) WithFunc(funcTag string) *Log {
	l.entry = l.getEntry().WithFields(LogFields{
		"funcTag": funcTag,
	})
	return l
}

// WithMessage returns a logger with message
func (l *Log) WithMessage(message string) *Log {
	l.entry = l.getEntry().WithFields(LogFields{
		"message": message,
	})
	return l
}

// WithMessagef returns a logger with formatted message
func (l *Log) WithMessagef(formattedMessage string, args ...interface{}) *Log {
	if len(args) != 0 {
		// this if check is to make sure that the error string does not
		// get mangled, which happens when you run sprintf without any
		// args.
		formattedMessage = fmt.Sprintf(formattedMessage, args)
	}
	l.entry = l.getEntry().WithFields(LogFields{
		"message": formattedMessage,
	})
	return l
}

// WithError returns a logger with message from error
func (l *Log) WithError(err error) *Log {
	if err == nil {
		err = fmt.Errorf("Unspecified Error")
	}
	l.entry = l.getEntry().WithFields(LogFields{
		"error": err,
	})
	return l
}

// WithStruct returns a logger with message from a struct
func (l *Log) WithStruct(genericStruct interface{}) *Log {
	l.entry = l.getEntry().WithFields(LogFields{
		"struct": fmt.Sprintf("%+v", genericStruct),
	})
	return l
}

// WithErrorMessage returns a logger with message from error text
func (l *Log) WithErrorMessage(errText string) *Log {
	return l.WithError(fmt.Errorf(errText))
}

// WithStatusCode returns a logger with statusCode
func (l *Log) WithStatusCode(statusCode int) *Log {
	l.entry = l.getEntry().WithFields(LogFields{
		"statusCode": statusCode,
	})
	return l
}

// WithDuration returns a logger with duration
func (l *Log) WithDuration(duration time.Duration) *Log {
	l.entry = l.getEntry().WithFields(LogFields{
		"duration": duration,
	})
	return l
}

// WithFields returns a logger with fields specified by the caller
func (l *Log) WithFields(fs LogFields) *Log {
	l.entry = l.getEntry().WithFields(fs)
	return l
}

// Info actuates the logging at the Info level
// Please add all log info before calling this function
func (l *Log) Info() {
	l.getEntry().Info()
}

// Warn actuates the logging at the Warn level
// Please add all log info before calling this function
func (l *Log) Warn() {
	l.getEntry().Warn()
}

// Error actuates the logging at the Error level
// Please add all log info before calling this function
func (l *Log) Error() {
	l.getEntry().Error()
}

// Panic actuates the logging at the Panic level, then calls panic()
// Please add all log info before calling this function
func (l *Log) Panic() {
	l.getEntry().Fatal()
}
