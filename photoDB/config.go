package photoDB

import (
	"fmt"
	"strings"
	"time"

	// sqlx plus specific drivers
	"github.com/atljoseph/api.josephgill.io/apierr"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	// other drivers supported go here
)

// Config describes the database reference wrapper for the photoDB package
type Config struct {
	// MaxIdleConns represents the maximum number of idle db connections to maintain at any given time
	MaxIdleConns int
	// MaxOpenConns represents the maximum number of open db connections to maintain at any given time
	MaxOpenConns int
	// ConnMaxLifetime is in minutes
	ConnMaxLifetimeMinutes int
	// Username for the database
	Username string
	// Password for the database
	Password string
	// Connection Host for the Databsae
	Host string
	// Connection Port for the Databsae
	Port int
	// DefaultDatabase for the connection
	DefaultDatabase string
	// PopulateOnStart tells db initialization whether to populate or not on start
	PopulateOnStart string
	// ------------------------------------------------
	// connString is the filepath to the database, or the connection string.
	connString string
	// ConnType is the database driver to use when connecting (Note: Requires manual importation of required driver)
	connType string
	// populateOnStartBool is the boolean version of PopulateOnStart
	populateOnStartBool bool
}

// MergeWithDefaults merges the passed in config with the default options
func (config *Config) MergeWithDefaults() *Config {
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
	if strings.EqualFold(config.DefaultDatabase, "") {
		config.DefaultDatabase = "photos"
	}
	config.connType = "mysql"
	if strings.EqualFold(config.connString, "") {
		config.connString = fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", config.Username, config.Password, config.Host, config.Port, config.DefaultDatabase)
	}
	if config.PopulateOnStart == "1" || strings.EqualFold(config.PopulateOnStart, "true") {
		config.populateOnStartBool = true
	}
	return config
}

// configDB configures the Photos DB and returns an error if error ocurrs
func configDB(config *Config) (*sqlx.DB, error) {
	funcTag := "configDB"

	// open the dbx object
	dbx, err := sqlx.Open(config.connType, config.connString)
	if err != nil {
		return nil, apierr.Errorf(err, funcTag, "opening db connection")
	}

	// ping the database
	err = dbx.Ping()
	if err != nil {
		return nil, apierr.Errorf(err, funcTag, "pinging db")
	}

	// set connection info
	dbx.SetMaxIdleConns(0)
	dbx.SetMaxOpenConns(1)
	dbx.SetConnMaxLifetime(time.Minute * time.Duration(config.ConnMaxLifetimeMinutes))

	return dbx, nil
}
