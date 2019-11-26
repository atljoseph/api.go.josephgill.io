package photoDB

// Album is the struct which describes photo albums
type Album struct {
	ID          int64    `json:"id" db:"album_id"`
	Title       string   `json:"title" db:"album_title"`
	Description string   `json:"description" db:"album_description"`
	Key         string   `json:"key" db:"album_key"`
	CoverPhoto  string   `json:"photo_src" db:"album_photo_src"`
	Photos      []*Photo `json:"photos" db:"-"`
}
