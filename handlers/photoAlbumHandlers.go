package handlers

import (
	"net/http"

	"github.com/atljoseph/api.josephgill.io/apierr"
	"github.com/atljoseph/api.josephgill.io/photoDB"

	"github.com/atljoseph/api.josephgill.io/responder"
)

// PostPhotoAlbumHandler gets all photo albums
func PostPhotoAlbumHandler(w http.ResponseWriter, r *http.Request) {
	funcTag := "PostPhotoAlbumHandler"

	// TODO: Get data from body of the request
	album := &photoDB.Album{
		Title:       "Yeah",
		Description: "Man",
		Key:         "main",
		CoverPhoto:  "sam-shortline-candler-grandy-papa-daddy-with-train-12.jpg",
	}

	// create a transaction
	// TODO: Parse JWT and Get user from JWT Request
	txo, err := photoDB.NewTxO("Test User")
	if err != nil {
		err = apierr.Errorf(err, funcTag, "open db transaction")
		responder.SendJSONError(w, http.StatusBadRequest, err)
		return
	}

	// create the album
	album, err = photoDB.CreatePhotoAlbum(txo, album)
	// TODO: Handle this with apierr http implementation - apierr.TerminateTxIfError()?
	if err = txo.RollbackOnError(err); err != nil {
		err = apierr.Errorf(err, funcTag, "create photo album")
		responder.SendJSONError(w, http.StatusBadRequest, err)
		return
	}
	//

	// commit transaction
	err = txo.Commit()
	if err != nil {
		err = apierr.Errorf(err, funcTag, "commit db transaction")
		responder.SendJSONError(w, http.StatusBadRequest, err)
		return
	}

	// build the return data
	res := &GetPhotoAlbumsResponse{}
	res.PhotoAlbums = []*photoDB.Album{album}

	// return
	responder.SendJSON(w, res)
}

// GetPhotoAlbumsResponse is the response returned by GetPhotoAlbumsHandler
type GetPhotoAlbumsResponse struct {
	PhotoAlbums []*photoDB.Album `json:"albums"`
}

// GetPhotoAlbumsHandler gets all photo albums
func GetPhotoAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	funcTag := "GetPhotoAlbumsHandler"

	// get the albums
	albums, err := photoDB.GetPhotoAlbums()
	if err != nil {
		err = apierr.Errorf(err, funcTag, "get photo albums")
		responder.SendJSONError(w, http.StatusBadRequest, err)
		return
	}

	// build the return data
	res := &GetPhotoAlbumsResponse{}
	res.PhotoAlbums = albums

	// return
	responder.SendJSON(w, res)
}
