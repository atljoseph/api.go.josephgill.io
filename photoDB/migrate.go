package photoDB

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

// TODO: write schema
// TODO: formalize migrations a little bit more
// TODO: run certain migrations only when prod mode
// TODO: handle db nulls with special type 	"gopkg.in/guregu/null.v3"

// migrateDB migrates the photos DB, and returns error if failed
func migrateDB(dbx *sqlx.DB, cx *Config) error {
	errTag := "migrateDB"

	// sum up migrations in a slice
	// TODO: get migrations in from somewhere else
	migrations := &migrate.MemoryMigrationSource{
		Migrations: migrations,
	}

	// apply migrations
	n, err := migrate.Exec(dbx.DB, cx.ConnType, migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("%s: %s", errTag, err)
	}

	// once migrated, the migration will not be reapplied :)
	fmt.Printf("sqlite db migrated with %d migrations\n", n)

	return nil
}

// migrations is a list of database migrations
// in- memory migrations will be built into the static binary
// "Down" migrations are omitted
// this process should only ever move forward, never backwards
// even if something is to be dropped, it should be included in a NEW migration
var migrations = []*migrate.Migration{
	&migrate.Migration{
		Id: "2019-000-create-table-testy",
		Up: []string{`
CREATE TABLE IF NOT EXISTS testy(  
	test_id INT NOT NULL AUTO_INCREMENT,  
	test_name VARCHAR(100) NOT NULL,  
	PRIMARY KEY ( test_id )  
) 
		`},
	},
	&migrate.Migration{
		Id: "2019-001-create-table-album",
		Up: []string{`
CREATE TABLE IF NOT EXISTS album (
	album_id BIGINT NOT NULL AUTO_INCREMENT
		, PRIMARY KEY (album_id)
	, album_title VARCHAR(100) NOT NULL
		, UNIQUE KEY(album_title)
	, album_description VARCHAR(300) NOT NULL
	, album_key VARCHAR(100) NOT NULL
		, UNIQUE KEY(album_key)
	, album_photo_src VARCHAR(4000) NOT NULL
)
		`},
	},
	&migrate.Migration{
		Id: "2019-002-populate-table-album",
		Up: []string{`
INSERT INTO album (album_title, album_description, album_key, album_photo_src) VALUES (
	'Riding the SAM Shortline Train 1'
	, 'From Cordele, GA to Plains, GA 1'
	, 'sam-shortline'
	, 'sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg 1'
)`, `
INSERT INTO album (album_title, album_description, album_key, album_photo_src) VALUES (
	'Riding the SAM Shortline Train 2'
	, 'From Cordele, GA to Plains, GA 2'
	, 'sam-shortline 2'
	, 'sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg 2'
)
		`},
	},
	&migrate.Migration{
		Id: "2019-003-create-table-photo",
		Up: []string{`
CREATE TABLE IF NOT EXISTS photo (
	photo_id BIGINT NOT NULL AUTO_INCREMENT
		, PRIMARY KEY (photo_id)
	, album_id BIGINT NOT NULL
	, photo_title VARCHAR(100) NOT NULL
	, photo_description VARCHAR(300) NOT NULL
	, photo_src VARCHAR(4000) NOT NULL
)
			`},
	},
	&migrate.Migration{
		Id: "2019-004-populate-table-photo",
		Up: []string{`
INSERT INTO photo (album_id, photo_title, photo_description, photo_src) VALUES (
	1
	, 'So much fun!'
	, 'Look at that smile'
	, 'sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg'
)
			`},
	},
}
