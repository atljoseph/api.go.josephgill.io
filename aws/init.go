package aws

import (
	"sync"

	"github.com/atljoseph/api.josephgill.io/apierr"
)

// private vars
var err error
var once sync.Once

// config for package
var config *Config

// Initialize initializes all a new aws connector package
// Call this function first!
func Initialize(c *Config) error {
	funcTag := "Initialize"

	// only do this the first time
	once.Do(func() {
		logMessage(funcTag, "start")

		// merge config with defaults
		c = c.MergeWithDefaults()

		// set the aws config
		if err == nil {
			config = c
		}

		// log it
		logMessage(funcTag, "end")

		S3PublicAssetList()
	})

	// return error if any, each and every time this function is called
	if err != nil {
		return apierr.Errorf(err, funcTag, "initializing package")
	}

	return nil
}
