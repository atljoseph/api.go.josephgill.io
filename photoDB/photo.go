package photoDB

// Photo is the struct which describes a single photo
type Photo struct {
	PhotoID     int64  `json:"photo_id" db:"photo_id"`
	AlbumID     int64  `json:"album_id" db:"album_id"`
	Title       string `json:"photo_title" db:"photo_title"`
	Description string `json:"photo_description" db:"photo_description"`
	Src         string `json:"photo_src" db:"photo_src"`
}
