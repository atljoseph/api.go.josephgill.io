package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// LogEntry is extended from logrus.Entry
type LogEntry = logrus.Entry

// Entry returns *LogEntry from an eventID
// eventID needs to be a unique string,
// which can be passed in as a package name
func Entry(eventID, action string) *LogEntry {
	return logrus.WithFields(LogFields{
		"source": eventID,
		"action": action,
	})
}

// EntryWithMessage logs *LogEntry with a message
func EntryWithMessage(eventID, action, message string) *LogEntry {
	return Entry(eventID, action).
		WithFields(LogFields{
			"message": message,
		})
}

// EntryWithError logs *LogEntry with a error
func EntryWithError(eventID, action string, err error) *LogEntry {
	if err == nil {
		err = fmt.Errorf("Unspecified Error")
	}
	return Entry(eventID, action).
		WithFields(LogFields{
			"error": err.Error(),
		})
}
