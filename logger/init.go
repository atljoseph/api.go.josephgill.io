package logger

import (
	"fmt"
	"log"
	"sync"

	"github.com/atljoseph/api.josephgill.io/apierr"
	"gopkg.in/natefinch/lumberjack.v2"
)

// private vars
var err error
var once sync.Once

// TODO: A real logging solution logrus?

// Initialize initializes this singleton package
func Initialize(c *Config) error {
	errTag := "logger.Initialize"

	once.Do(func() {
		// merge config with defaults
		c = c.MergeWithDefaults()

		// log the config
		fmt.Printf("Config [logger]: %+v\n", c)

		if err == nil {
			if c.Filename != "" {
				log.SetOutput(&lumberjack.Logger{
					Filename:   c.Filename,
					MaxSize:    500, // in MB
					MaxBackups: 3,
					MaxAge:     28,   // in days
					Compress:   true, // false by default
				})
			}
		}

	})

	// return error if any, each and every time this function is called
	if err != nil {
		return apierr.Errorf(err, errTag, "initializing logger")
	}

	return nil
}
