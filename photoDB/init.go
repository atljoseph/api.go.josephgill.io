package photoDB

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
)

const (
	dbxType = "sqlite3"
)

// TODO: load db params from os.Getenv("key"), OR load from JSON

// private vars
var err error
var once sync.Once

// dbx is the *sqlx.DB object for the photos database
var dbx *sqlx.DB

// Initialize initializes all db connections used in the app
// Call this function first!
func Initialize(config *Config) error {
	errTag := "photoDB.Initialize"

	// only do this the first time
	once.Do(func() {
		// init
		if err == nil {
			dbx, err = configDB(config)
		}
	})

	// return error if any, each and every time this function is called
	if err != nil {
		err = fmt.Errorf("%s: %s", errTag, err)
		return err
	}

	return nil
}
