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
func migrateDB(dbx *sqlx.DB) error {
	errTag := "migrateDB"

	// sum up migrations in a slice
	// TODO: get migrations in from somewhere else
	migrations := &migrate.MemoryMigrationSource{
		Migrations: migrations,
	}

	// apply migrations
	n, err := migrate.Exec(dbx.DB, dbxType, migrations, migrate.Up)
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
		Id: "2019-001-create-table-album",
		Up: []string{`
CREATE TABLE IF NOT EXISTS album (
	album_id INTEGER PRIMARY KEY AUTOINCREMENT
	, title text NOT NULL UNIQUE
	, description text NOT NULL
	, key text NOT NULL UNIQUE
	, cover_photo_src text NOT NULL UNIQUE
);
		`},
	},
	&migrate.Migration{
		Id: "2019-002-populate-table-album",
		Up: []string{`
INSERT INTO album (title, description, key, cover_photo_src) VALUES (
	'Riding the SAM Shortline Train'
	, 'From Cordele, GA to Plains, GA'
	, 'sam-shortline'
	, 'sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg'
);
INSERT INTO album (title, description, key, cover_photo_src) VALUES (
	'8', '7', '6', '5'
);
		`},
	},
	&migrate.Migration{
		Id: "2019-003-create-table-photo",
		Up: []string{`
CREATE TABLE IF NOT EXISTS photo (
	photo_id INTEGER PRIMARY KEY AUTOINCREMENT
	, album_id INTEGER 
	, title text NOT NULL
	, description text NOT NULL
	, src text NOT NULL UNIQUE
);
		`},
	},
	&migrate.Migration{
		Id: "2019-004-populate-table-photo",
		Up: []string{`
INSERT INTO photo (album_id, title, description, src) VALUES (
	1
	, 'So much fun!'
	, 'Look at that smile'
	, 'sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg'
);
		`},
	},
}
