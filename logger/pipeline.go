package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// TODO: pipe the logger data to external source

type externalLoggingHook struct {
}

func (elh externalLoggingHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (elh externalLoggingHook) Fire(entry *logrus.Entry) error {
	// non-blocking call
	go func() {
		// add external logging source fields
		// be sure they do not conflict with existing logrus fields
		// ...
		// entry = entry.WithFields(LogFields{"other": "fields"})

		// send to external logging source
		fmt.Printf("SENDING TO DATADOG ==> \n%+v\n", entry)
	}()
	return nil
}

// func getLoggingHook() *.logrus.Hook  {
// 	return new logrus.Hook{}
// }

// sendToPipeline will send data to the logger pipeline
// will be hooked into a logrus hook
func sendToPipeline(l *Log) {
	// send to data dog
}
