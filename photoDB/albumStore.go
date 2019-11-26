package photoDB

import (
	"fmt"

	"github.com/atljoseph/api.josephgill.io/apierr"
)

// CreatePhotoAlbum is just a test right now
func CreatePhotoAlbum(txo *TxO, album *Album) (*Album, error) {
	funcTag := "CreatePhotoAlbum"

	pkgLog.WithFunc(funcTag).WithMessage("insert new photo album").Info()

	// build the query
	query := `
	INSERT INTO album (album_title, album_description, album_key, album_photo_src) VALUES (
		:album_title
		, :album_description
		, :album_key
		, :album_photo_src
	);
	`

	// get the result
	result, err := txo.NamedExec(query, album)
	if err != nil {
		return nil, apierr.Errorf(err, funcTag, "NamedExec")
	}

	// last inserted id
	id, err := result.LastInsertId()
	if err != nil {
		return nil, apierr.Errorf(err, funcTag, "LastInsertId")
	}
	album.ID = id

	return album, nil
}

// GetPhotoAlbums will eventually grab the photo albums the database, or error
func GetPhotoAlbums() ([]*Album, error) {
	funcTag := "GetPhotoAlbums"

	pkgLog.WithFunc(funcTag).WithMessage("get all photo albums").Info()

	// query with the dbx object
	var pas []*Album
	query := `
SELECT * from album
	`

	// run the query
	err = dbx.Select(&pas, query)
	if err != nil {
		return nil, apierr.Errorf(err, funcTag, "Select")
	}

	// must have rows returned
	if len(pas) == 0 {
		return nil, apierr.Errorf(fmt.Errorf("No rows returned from query"), funcTag, "could not find any photo albums")
	}

	return pas, nil
}
