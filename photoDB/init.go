package photoDB

import (
	"sync"
	"time"

	"github.com/atljoseph/api.josephgill.io/apierr"
	"github.com/jmoiron/sqlx"
)

// private vars
var err error
var once sync.Once

// dbx is the *sqlx.DB object for the photos database
var dbx *sqlx.DB

// Initialize initializes all db connections used in the app
// Call this function first!
func Initialize(c *Config) error {
	funcTag := "Initialize"

	// sleep X seconds to give the db time to warm up if needed
	logMessage(funcTag, "Sleeping to give db time to warm up")
	time.Sleep(10 * time.Second)

	// only do this the first time
	once.Do(func() {
		// merge config with defaults
		c = c.MergeWithDefaults()

		// log the config
		logMessage(funcTag, "start")

		// config the db
		if err == nil {
			dbx, err = configDB(c)
		}

		// apply migrations
		if err == nil {
			err = migrateDB(c)
		}

		logMessage(funcTag, "end")
	})

	// return error if any, each and every time this function is called
	if err != nil {
		return apierr.Errorf(err, funcTag, "initializing db")
	}

	return nil
}
