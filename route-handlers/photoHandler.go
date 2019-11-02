package routeHandlers

import (
	"fmt"
	"net/http"

	"github.com/atljoseph/api.josephgill.io/photoDB"
	"github.com/atljoseph/api.josephgill.io/requester"
	"github.com/atljoseph/api.josephgill.io/responder"
)

// GetPhotosResponse is the response returned by GetPhotoAlbumsHandler
type GetPhotosResponse struct {
	Photos []*photoDB.Photo `json:"photos"`
}

// PostPhotoHandler gets all photo albums
func PostPhotoHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "handlers.PostPhotoHandler"

	// TODO: Get data from body of the request
	photo := &photoDB.Photo{
		AlbumID:     3,
		Title:       "Yeah",
		Description: "Man",
		Src:         "sam-shortline-candler-grandy-papa-daddy-with-train-12.jpg",
	}

	// create a transaction
	txo, err := photoDB.NewTxO("Test User")
	if err != nil {
		responder.SendJSONHttpError(w, http.StatusBadRequest, fmt.Sprintf("%s: %s", errTag, err))
	}

	// get the albums
	photo, err = photoDB.CreatePhoto(txo, photo)
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
	res := &GetPhotosResponse{}
	res.Photos = []*photoDB.Photo{photo}
	responder.SendJSON(w, res)
}

// GetPhotosByAlbumKeyHandler gets all photo albums
func GetPhotosByAlbumKeyHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "handlers.GetPhotosByAlbumKeyHandler"

	// process request params
	mp, err := requester.Process(r, nil, requester.PhotoAlbumIDKey)
	if err != nil {
		responder.SendJSONHttpError(w, http.StatusBadRequest, fmt.Sprintf("%s: %s", errTag, err))
		return
	}

	// get the photos
	ps, err := photoDB.GetPhotosByAlbumKey(mp[requester.PhotoAlbumIDKey])
	if err != nil {
		responder.SendJSONHttpError(w, http.StatusBadRequest, fmt.Sprintf("%s: %s", errTag, err))
		return
	}

	// build the return data
	res := &GetPhotosResponse{}
	res.Photos = ps
	responder.SendJSON(w, res)
}
