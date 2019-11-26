package photoDB

import (
	"fmt"

	"github.com/atljoseph/api.josephgill.io/apierr"
)

// GetPhotosByAlbumKey gets a photo album and returns it or an error
func GetPhotosByAlbumKey(key string) ([]*Photo, error) {
	funcTag := "GetPhotosByAlbumKey"

	pkgLog.WithFunc(funcTag).WithMessage("get all photos for album").Info()

	// build the query
	var ps []*Photo
	query := `
SELECT * from photo where album_key = ?
	`
	// SELECT p.*
	// 	FROM photo p
	// 	INNER JOIN album a on a.album_id = p.album_id
	// 	WHERE a.album_key = ?

	err = dbx.Select(&ps, query, key)
	if err != nil {
		return nil, apierr.Errorf(err, funcTag, "Select")
	}

	// TODO: wrap this in a proper error for no rows returned
	if len(ps) == 0 {
		return nil, apierr.Errorf(fmt.Errorf("No rows returned from query"), funcTag, "could not find any photos")
	}

	return ps, nil
}

// CreatePhoto is just a test right now
func CreatePhoto(txo *TxO, photo *Photo) (*Photo, error) {
	funcTag := "CreatePhoto"

	pkgLog.WithFunc(funcTag).WithMessage("insert new album photo").Info()

	// build the query
	query := `
INSERT INTO photo (
	album_id
	, album_key
	, photo_title
	, photo_description
	, photo_src
) VALUES (
	:album_id
	, :album_key
	, :photo_title
	, :photo_description
	, :photo_src
);
	`

	// get the result
	result, err := txo.NamedExec(query, photo)
	if err != nil {
		return nil, apierr.Errorf(err, funcTag, "NamedExec")
	}

	// last inserted id
	id, err := result.LastInsertId()
	if err != nil {
		return nil, apierr.Errorf(err, funcTag, "LastInsertId")
	}
	photo.PhotoID = id

	return photo, nil
}
