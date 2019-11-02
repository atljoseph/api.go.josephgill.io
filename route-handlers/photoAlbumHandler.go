package routeHandlers

import (
	"fmt"
	"net/http"

	"github.com/atljoseph/api.josephgill.io/photoDB"

	"github.com/atljoseph/api.josephgill.io/responder"
)

// PostPhotoAlbumsHandler gets all photo albums
func PostPhotoAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "handlers.PostPhotoAlbumsHandler"

	// TODO: Get data from body of the request
	album := &photoDB.PhotoAlbum{
		Title:         "Yeah",
		Description:   "Man",
		Key:           "main",
		CoverPhotoSrc: "sam-shortline-candler-grandy-papa-daddy-with-train-12.jpg",
	}

	// create a transaction
	// TODO: User from JWT Request
	txo, err := photoDB.NewTxO("Test User")
	if err != nil {
		responder.SendJSONHttpError(w, http.StatusBadRequest, fmt.Sprintf("%s: %s", errTag, err))
	}

	// get the albums
	album, err = photoDB.CreatePhotoAlbum(txo, album)
	if errTxo := txo.RollbackOnError(err); errTxo != nil {
		responder.SendJSONHttpError(w, http.StatusBadRequest, fmt.Sprintf("%s: %s", errTag, err))
		return
	}

	// commit transaction
	err = txo.Commit()
	if err != nil {
		responder.SendJSONHttpError(w, http.StatusBadRequest, fmt.Sprintf("%s: %s", errTag, err))
	}

	// build the return data
	res := &GetPhotoAlbumsResponse{}
	res.PhotoAlbums = []*photoDB.PhotoAlbum{album}
	responder.SendJSON(w, res)
}

// GetPhotoAlbumsResponse is the response returned by GetPhotoAlbumsHandler
type GetPhotoAlbumsResponse struct {
	PhotoAlbums []*photoDB.PhotoAlbum `json:"albums"`
}

// GetPhotoAlbumsHandler gets all photo albums
func GetPhotoAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "handlers.GetPhotoAlbumsHandler"

	// get the albums
	albums, err := photoDB.GetPhotoAlbums()
	if err != nil {
		responder.SendJSONHttpError(w, http.StatusBadRequest, fmt.Sprintf("%s: %s", errTag, err))
		return
	}

	// build the return data
	res := &GetPhotoAlbumsResponse{}
	res.PhotoAlbums = albums
	responder.SendJSON(w, res)
}
