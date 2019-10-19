package handlers

import (
	"fmt"
	"net/http"

	"github.com/atljoseph/api.josephgill.io/stores"

	"github.com/atljoseph/api.josephgill.io/responder"
)

// GetPhotoAlbumsResponse is the response returned by GetPhotoAlbumsHandler
type GetPhotoAlbumsResponse struct {
	PhotoAlbums []stores.PhotoAlbum `json:"albums"`
}

// GetPhotoAlbumsHandler gets all photo albums
func GetPhotoAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "handlers.GetPhotoAlbumsHandler"

	res := &GetPhotoAlbumsResponse{}

	albums, err := stores.GetPhotoAlbums()
	if err != nil {
		http.Error(w, fmt.Errorf("%s: %s", errTag, err).Error(), http.StatusBadRequest)
		return
	}

	res.PhotoAlbums = albums
	responder.SendJSON(w, res)
}
