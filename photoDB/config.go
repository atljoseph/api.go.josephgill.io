package photoDB

import (
	"fmt"
	"strings"
	"time"

	// sqlx plus specific drivers
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	// other drivers supported go here
)

// Config describes the database reference wrapper for the photoDB package
type Config struct {
	// Username for the database
	Username string
	// Password for the database
	Password string
	// Connection Host for the Databsae
	Host string
	// Connection Port for the Databsae
	Port int
	// DefaultDetabase for the connection
	DefaultDetabase string
	// ------------------------------------------------
	// ConnString is the filepath to the database, or the connection string.
	ConnString string
	// ConnType is the database driver to use when connecting (Note: Requires manual importation of required driver)
	ConnType string
	// MaxIdleConns represents the maximum number of idle db connections to maintain at any given time
	MaxIdleConns int
	// MaxOpenConns represents the maximum number of open db connections to maintain at any given time
	MaxOpenConns int
	// ConnMaxLifetime is in minutes
	ConnMaxLifetimeMinutes int
}

// MergeWithDefaults merges the passed in config with the default options
func (config *Config) MergeWithDefaults() *Config {
	if strings.EqualFold(config.Username, "") {
		config.Username = "root"
	}
	if strings.EqualFold(config.Password, "") {
		config.Password = "password"
	}
	if strings.EqualFold(config.Host, "") {
		config.Host = "localhost"
	}
	if config.Port == 0 {
		config.Port = 3306
	}
	if strings.EqualFold(config.DefaultDetabase, "") {
		config.DefaultDetabase = "photos"
	}
	config.ConnType = "mysql"
	if strings.EqualFold(config.ConnString, "") {
		config.ConnString = fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", config.Username, config.Password, config.Host, config.Port, config.DefaultDetabase)
	}
	// ints will be 0 by default if not set
	// negatives have no meaning
	if config.MaxIdleConns < 0 {
		config.MaxIdleConns = 0
	}
	if config.MaxOpenConns < 0 || config.MaxOpenConns == 0 {
		config.MaxOpenConns = 1
	}
	if config.ConnMaxLifetimeMinutes < 0 || config.ConnMaxLifetimeMinutes == 0 {
		config.ConnMaxLifetimeMinutes = 5
	}
	fmt.Printf("Config [photoDB]: %+v\n", config)
	return config
}

// configDB configures the Photos DB and returns an error if error ocurrs
func configDB(config *Config) (*sqlx.DB, error) {
	errTag := "configDB"

	// merge config with defaults
	cx := config.MergeWithDefaults()

	// open the dbx object
	dbx, err := sqlx.Open(config.ConnType, config.ConnString)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errTag, err)
	}

	// if dbx.
	// // the problem of existance
	// // only sqlite supported for now
	// if strings.Contains(config.ConnType, "sqlite") {
	// 	// provide a clean database on each run
	// 	if config.ReplaceDB {
	// 		os.Remove(config.ConnString)
	// 	}

	// 	// create the file if not exists
	// 	_, err = os.Stat(config.ConnString)
	// 	if os.IsNotExist(err) {
	// 		dbx = sqlx.NewDb(dbx.DB, config.ConnType)
	// 	}
	// }

	// ping the database
	err = dbx.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errTag, err)
	}

	// set connection info
	dbx.SetMaxIdleConns(0)
	dbx.SetMaxOpenConns(1)
	dbx.SetConnMaxLifetime(time.Minute * time.Duration(config.ConnMaxLifetimeMinutes))

	// apply migrations
	err = migrateDB(dbx, cx)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", errTag, err)
	}
	fmt.Printf("%s: %s", errTag, "Migrated")

	return dbx, nil
}
