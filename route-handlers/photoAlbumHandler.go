package routeHandlers

import (
	"fmt"
	"net/http"

	"github.com/atljoseph/api.josephgill.io/photoDB"
	"github.com/atljoseph/api.josephgill.io/requester"

	"github.com/atljoseph/api.josephgill.io/responder"
)

// PostPhotoAlbumResponse is the response returned by PostPhotoAlbumsHandler
type PostPhotoAlbumResponse struct {
	PhotoAlbum *photoDB.PhotoAlbum `json:"album"`
}

// PostPhotoAlbumsHandler gets all photo albums
func PostPhotoAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "handlers.PostPhotoAlbumsHandler"

	album := &photoDB.PhotoAlbum{
		Title:         "Yeah",
		Description:   "Man",
		Key:           "main",
		CoverPhotoSrc: "sam-shortline-candler-grandy-papa-daddy-with-train-12.jpg",
	}

	// get the albums
	album, err := photoDB.CreatePhotoAlbum(album)
	if err != nil {
		responder.SendJSONHttpError(w, http.StatusBadRequest, fmt.Sprintf("%s: %s", errTag, err))
		return
	}

	// build the return data
	res := &PostPhotoAlbumResponse{}
	res.PhotoAlbum = album
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

// GetPhotosResponse is the response returned by GetPhotoAlbumsHandler
type GetPhotosResponse struct {
	Photos []*photoDB.Photo `json:"photos"`
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
