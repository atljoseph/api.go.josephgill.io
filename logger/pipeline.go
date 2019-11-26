package logger

import (
	"github.com/sirupsen/logrus"
)

// satisfy the logrus logging hook
type externalLoggingHook struct{}

// this is in the interface
func (elh externalLoggingHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// this is in the interface, too
func (elh externalLoggingHook) Fire(entry *logrus.Entry) error {
	// non-blocking call
	go func() {
		// add external logging source fields
		// be sure they do not conflict with existing logrus fields
		// ...
		// entry = entry.WithFields(LogFields{"other": "fields"})

		// send to external logging source
		// fmt.Printf("SENDING TO DATADOG ==> \n%+v\n", entry)
	}()
	return nil
}
