package photoDB

import "fmt"

// GetPhotosByAlbumKey gets a photo album and returns it or an error
func GetPhotosByAlbumKey(key string) ([]*Photo, error) {
	errTag := "photoDB.GetPhotosByAlbumKey"

	// build the query
	var ps []*Photo
	query := `
SELECT p.* 
	FROM photo p 
	INNER JOIN album a on a.album_id = p.album_id
	WHERE a.album_key = ?
	`

	err = dbx.Select(&ps, query, key)
	if err != nil {
		return nil, fmt.Errorf("%s: Select: %s", errTag, err)
	}

	// TODO: wrap this in a proper error
	if len(ps) == 0 {
		return nil, fmt.Errorf("%s: %s", errTag, "No rows returned from query")
	}

	return ps, nil
}

// CreatePhoto is just a test right now
func CreatePhoto(txo *TxO, photo *Photo) (*Photo, error) {
	errTag := "photoDB.CreatePhoto"

	// build the query
	query := `
	INSERT INTO photo (album_id, photo_title, photo_description, photo_src) VALUES (
		:album_id
		, :photo_title
		, :photo_description
		, :photo_src
	);
	`

	// get the result
	result, err := txo.NamedExec(query, photo)
	if err != nil {
		return nil, fmt.Errorf("%s: NamedExec: %s", errTag, err)
	}

	// last inserted id
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%s: LastInsertId: %s", errTag, err)
	}
	photo.PhotoID = id

	return photo, nil
}
