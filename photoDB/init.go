package photoDB

import (
	"sync"
	"time"

	"github.com/atljoseph/api.josephgill.io/apierr"
	"github.com/atljoseph/api.josephgill.io/logger"
	"github.com/jmoiron/sqlx"
)

// private vars
var err error
var once sync.Once
var pkgLog *logger.Log

// dbx is the *sqlx.DB object for the photos database
var dbx *sqlx.DB

// Initialize initializes all db connections used in the app
// Call this function first!
func Initialize(c *Config) error {
	funcTag := "Initialize"

	// only do this the first time
	once.Do(func() {

		// init the logger
		pkgLog = logger.ForPackage("aws")
		funcLog := pkgLog.WithFunc(funcTag)

		// sleep X seconds to give the db time to warm up if needed
		funcLog.WithMessage("Sleeping to give db time to warm up").Info()
		time.Sleep(10 * time.Second)

		// merge config with defaults
		c = c.MergeWithDefaults()

		// log the config
		funcLog.WithMessage("initializing db connector").Info()

		// config the db
		if err == nil {
			dbx, err = configDB(c)
		}

		// apply migrations
		if err == nil {
			err = migrateDB(c)
		}

		funcLog.WithMessage("db connector initialized").Info()
	})

	// return error if any, each and every time this function is called
	if err != nil {
		return apierr.Errorf(err, funcTag, "initializing db")
	}

	return nil
}
