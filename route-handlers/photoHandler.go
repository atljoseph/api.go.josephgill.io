package routeHandlers

import (
	"net/http"

	"github.com/atljoseph/api.josephgill.io/apierr"
	"github.com/atljoseph/api.josephgill.io/photoDB"
	"github.com/atljoseph/api.josephgill.io/requester"
	"github.com/atljoseph/api.josephgill.io/responder"
)

// GetPhotosResponse is the response returned by GetPhotoAlbumsHandler
type GetPhotosResponse struct {
	Photos []*photoDB.Photo `json:"photos"`
}

// PostPhotoAlbumPhotoHandler gets all photo albums
func PostPhotoAlbumPhotoHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "PostPhotoAlbumPhotoHandler"

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
		err = apierr.Errorf(err, errTag, "open db transaction")
		responder.SendJSONHttpError(w, http.StatusBadRequest, err)
		return
	}

	// get the albums
	photo, err = photoDB.CreatePhoto(txo, photo)
	if errTxo := txo.RollbackOnError(err); errTxo != nil {
		err = apierr.Errorf(err, errTag, "create photo")
		responder.SendJSONHttpError(w, http.StatusBadRequest, err)
		return
	}

	// commit transaction
	err = txo.Commit()
	if err != nil {
		err = apierr.Errorf(err, errTag, "commit db transaction")
		responder.SendJSONHttpError(w, http.StatusBadRequest, err)
		return
	}

	// build the return data
	res := &GetPhotosResponse{}
	res.Photos = []*photoDB.Photo{photo}

	// return
	responder.SendJSON(w, res)
}

// GetPhotosByAlbumKeyHandler gets all photo albums
func GetPhotosByAlbumKeyHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "GetPhotosByAlbumKeyHandler"

	// process request params
	mp, err := requester.Process(r, nil, requester.PhotoAlbumIDKey)
	if err != nil {
		err = apierr.Errorf(err, errTag, "process request params")
		responder.SendJSONHttpError(w, http.StatusBadRequest, err)
		return
	}

	// get the photos
	ps, err := photoDB.GetPhotosByAlbumKey(mp[requester.PhotoAlbumIDKey])
	if err != nil {
		err = apierr.Errorf(err, errTag, "get photos by album key")
		responder.SendJSONHttpError(w, http.StatusBadRequest, err)
		return
	}

	// build the return data
	res := &GetPhotosResponse{}
	res.Photos = ps

	// return
	responder.SendJSON(w, res)
}
