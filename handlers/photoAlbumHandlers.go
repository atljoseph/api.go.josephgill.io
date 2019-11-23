package handlers

import (
	"net/http"

	"github.com/atljoseph/api.josephgill.io/apierr"
	"github.com/atljoseph/api.josephgill.io/photoDB"

	"github.com/atljoseph/api.josephgill.io/responder"
)

// PostPhotoAlbumHandler gets all photo albums
func PostPhotoAlbumHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "PostPhotoAlbumHandler"

	// TODO: Get data from body of the request
	album := &photoDB.PhotoAlbum{
		Title:         "Yeah",
		Description:   "Man",
		Key:           "main",
		CoverPhotoSrc: "sam-shortline-candler-grandy-papa-daddy-with-train-12.jpg",
	}

	// create a transaction
	// TODO: Parse JWT and Get user from JWT Request
	txo, err := photoDB.NewTxO("Test User")
	if err != nil {
		err = apierr.Errorf(err, errTag, "open db transaction")
		responder.SendJSONHttpError(w, http.StatusBadRequest, err)
		return
	}

	// get the albums
	album, err = photoDB.CreatePhotoAlbum(txo, album)
	// TODO: Handle this with apierr http implementation - apierr.TerminateTxIfError()?
	err = txo.RollbackOnError(err)
	if err != nil {
		err = apierr.Errorf(err, errTag, "create photo album")
		responder.SendJSONHttpError(w, http.StatusBadRequest, err)
		return
	}
	//

	// commit transaction
	err = txo.Commit()
	if err != nil {
		err = apierr.Errorf(err, errTag, "commit db transaction")
		responder.SendJSONHttpError(w, http.StatusBadRequest, err)
		return
	}

	// build the return data
	res := &GetPhotoAlbumsResponse{}
	res.PhotoAlbums = []*photoDB.PhotoAlbum{album}

	// return
	responder.SendJSON(w, res)
}

// GetPhotoAlbumsResponse is the response returned by GetPhotoAlbumsHandler
type GetPhotoAlbumsResponse struct {
	PhotoAlbums []*photoDB.PhotoAlbum `json:"albums"`
}

// GetPhotoAlbumsHandler gets all photo albums
func GetPhotoAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "GetPhotoAlbumsHandler"

	// get the albums
	albums, err := photoDB.GetPhotoAlbums()
	if err != nil {
		err = apierr.Errorf(err, errTag, "get photo albums")
		responder.SendJSONHttpError(w, http.StatusBadRequest, err)
		return
	}

	// build the return data
	res := &GetPhotoAlbumsResponse{}
	res.PhotoAlbums = albums

	// return
	responder.SendJSON(w, res)
}
