package stores

// PhotoAlbum is the struct which describes photo albums
type PhotoAlbum struct {
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Key           string  `json:"key"`
	CoverPhotoSrc string  `json:"coverPhotoSrc"`
	Photos        []Photo `json:"photos"`
}

// Photo is the struct which describes a single photo
type Photo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Src         string `json:"src"`
}

// GetPhotoAlbums will eventually grab the photo albums from pg database, or other source
func GetPhotoAlbums() ([]PhotoAlbum, error) {

	albums := []PhotoAlbum{
		PhotoAlbum{
			Title:         "Riding the SAM Shortline Train",
			Key:           "sam-shortline",
			CoverPhotoSrc: "sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg",
			Photos: []Photo{
				Photo{
					Title: "Cordele to Plains, GA", Src: "sam-shortline-candler-grandy-papa-daddy-with-train-1.jpg",
				},
				Photo{
					Src: "sam-shortline-candler-pointing-1.jpg",
				},
			},
		},
	}

	return albums, nil
}
