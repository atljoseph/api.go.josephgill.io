package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var once sync.Once
var instance *sql.DB
var err error

// Configure sets up and returns a sqlite3 database instance
func Configure(filepath string) (*sql.DB, error) {
	os.Remove(filepath)

	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil, nil
	}

	return db, nil
}

// Connect sets up the DB connection(s) for the app and returns an error if not successful
func Connect(filepath string) (*sql.DB, error) {
	errTag := "database.Setup"

	once.Do(func() {
		instance, err = Configure(filepath)
	})

	if err != nil {
		return nil, fmt.Errorf("%s: %s", errTag, err)
	}

	return instance, nil

}
