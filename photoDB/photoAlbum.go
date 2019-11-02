package photoDB

// PhotoAlbum is the struct which describes photo albums
type PhotoAlbum struct {
	AlbumID       int64    `json:"album_id" db:"album_id"`
	Title         string   `json:"album_title" db:"album_title"`
	Description   string   `json:"album_description" db:"album_description"`
	Key           string   `json:"album_key" db:"album_key"`
	CoverPhotoSrc string   `json:"album_photo_src" db:"album_photo_src"`
	Photos        []*Photo `json:"album_photos" db:"photos"`
}
