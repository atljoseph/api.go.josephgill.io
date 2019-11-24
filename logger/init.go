package logger

import (
	"fmt"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

// private vars
var err error
var once sync.Once

// Initialize initializes this singleton package
func Initialize(c *Config) error {
	funcTag := "Initialize"

	once.Do(func() {
		// merge config with defaults
		c = c.MergeWithDefaults()

		if err == nil {
			// Log as JSON instead of the default ASCII formatter.
			// logrus.SetFormatter(&logrus.JSONFormatter{})
			logrus.SetFormatter(&logrus.TextFormatter{})

			// Output to stdout instead of the default stderr
			// Can be any io.Writer, see below for File example
			logrus.SetOutput(os.Stdout)

			// Only log the warning severity or above.
			logrus.SetLevel(logrus.DebugLevel)
		}

		// log things AFTER initialized (above)
		logMessage(funcTag, "init completed")
	})

	// return error if any, each and every time this function is called
	if err != nil {
		return fmt.Errorf("%s (%s) --> %s", funcTag, "initializing logger", err)
	}

	return nil
}
