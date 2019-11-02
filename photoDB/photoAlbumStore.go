package photoDB

import "fmt"

// CreatePhotoAlbum is just a test right now
func CreatePhotoAlbum(txo *TxO, album *PhotoAlbum) (*PhotoAlbum, error) {
	errTag := "photoDB.CreatePhotoAlbum"

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
		return nil, fmt.Errorf("%s: NamedExec: %s", errTag, err)
	}

	// last inserted id
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%s: LastInsertId: %s", errTag, err)
	}
	album.AlbumID = id

	return album, nil
}

// GetPhotoAlbums will eventually grab the photo albums the database, or error
func GetPhotoAlbums() ([]*PhotoAlbum, error) {
	errTag := "photoDB.GetPhotoAlbums"

	// query with the dbx object
	var pas []*PhotoAlbum
	query := `
SELECT * from album
	`

	// run the query
	err = dbx.Select(&pas, query)
	if err != nil {
		return nil, fmt.Errorf("%s: Select: %s", errTag, err)
	}

	// must have rows returned
	if len(pas) == 0 {
		return nil, fmt.Errorf("%s: %s", errTag, "No rows returned from query")
	}

	return pas, nil
}
