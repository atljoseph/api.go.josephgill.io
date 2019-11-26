package photoDB

// Photo is the struct which describes a single photo
type Photo struct {
	PhotoID     int64  `json:"id" db:"photo_id"`
	AlbumKey    string `json:"album_key" db:"album_key"`
	AlbumID     int64  `json:"album_id" db:"album_id"`
	Title       string `json:"title" db:"photo_title"`
	Description string `json:"description" db:"photo_description"`
	Src         string `json:"src" db:"photo_src"`
}
