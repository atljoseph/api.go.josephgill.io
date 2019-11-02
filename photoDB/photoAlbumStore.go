package photoDB

import "fmt"

// GetPhotosByAlbumKey gets a photo album and returns it or an error
func GetPhotosByAlbumKey(key string) ([]*Photo, error) {
	errTag := "photoDB.GetPhotosByAlbumKey"

	// query with the dbx object
	var ps []*Photo
	query := `
SELECT p.* 
	FROM photo p 
	INNER JOIN album a on a.album_id = p.album_id
	WHERE a.album_key = ?
	`
	if err := dbx.Select(&ps, query, key); err != nil {
		return nil, fmt.Errorf("%s: %s", errTag, err)
	}

	// TODO: wrap this in a proper error
	if len(ps) == 0 {
		return nil, fmt.Errorf("%s: %s", errTag, "No rows returned from query")
	}

	return ps, nil
}

// CreatePhotoAlbum is just a test right now
func CreatePhotoAlbum(album *PhotoAlbum) (*PhotoAlbum, error) {
	errTag := "photoDB.CreatePhotoAlbum"

	// TODO: Get data from body of the request

	query := `
	INSERT INTO album (album_title, album_description, album_key, album_photo_src) VALUES (
		:album_title
		, :album_description
		, :album_key
		, :album_photo_src
	);
	`
	result, err := dbx.NamedExec(query, album)
	if err != nil {
		return nil, fmt.Errorf("%s: %s: NamedExec ", errTag, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%s: %s: LastInsertId", errTag, err)
	}

	album.AlbumID = id

	// TODO: wrap this in a proper error
	// if len(ps) == 0 {
	// 	return nil, fmt.Errorf("%s: %s", errTag, "No rows returned from query")
	// }

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
	if err := dbx.Select(&pas, query); err != nil {
		return nil, fmt.Errorf("%s: %s", errTag, err)
	}

	if len(pas) == 0 {
		return nil, fmt.Errorf("%s: %s", errTag, "No rows returned from query")
	}

	return pas, nil
}
