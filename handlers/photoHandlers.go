package handlers

import (
	"fmt"
	"net/http"

	"github.com/atljoseph/api.josephgill.io/apierr"
	"github.com/atljoseph/api.josephgill.io/aws"
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
	funcTag := "PostPhotoAlbumPhotoHandler"

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
		err = apierr.Errorf(err, funcTag, "open db transaction")
		responder.SendJSONError(w, http.StatusBadRequest, err)
		return
	}

	// get the albums
	photo, err = photoDB.CreatePhoto(txo, photo)
	if errTxo := txo.RollbackOnError(err); errTxo != nil {
		err = apierr.Errorf(err, funcTag, "create photo")
		responder.SendJSONError(w, http.StatusBadRequest, err)
		return
	}

	// commit transaction
	err = txo.Commit()
	if err != nil {
		err = apierr.Errorf(err, funcTag, "commit db transaction")
		responder.SendJSONError(w, http.StatusBadRequest, err)
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
	funcTag := "GetPhotosByAlbumKeyHandler"

	// process request params
	mp, err := requester.GetRequestParams(r, nil, requester.PhotoAlbumIDKey)
	if err != nil {
		err = apierr.Errorf(err, funcTag, "process request params")
		responder.SendJSONError(w, http.StatusBadRequest, err)
		return
	}

	// get the photos
	ps, err := photoDB.GetPhotosByAlbumKey(mp[requester.PhotoAlbumIDKey])
	if err != nil {
		err = apierr.Errorf(err, funcTag, "get photos by album key")
		responder.SendJSONError(w, http.StatusBadRequest, err)
		return
	}

	// give the photos their url from s3
	// TODO: have the client pass in a quality filter via query params ("1024" below)
	for _, p := range ps {
		relativePath := fmt.Sprintf("%s/%s", "1024", p.Src)
		p.Src = aws.S3PublicAssetURL(relativePath)
	}

	// build the return data
	res := &GetPhotosResponse{}
	res.Photos = ps

	// return
	responder.SendJSON(w, res)
}
