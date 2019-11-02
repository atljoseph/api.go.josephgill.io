package routes

import (
	"fmt"

	"github.com/atljoseph/api.josephgill.io/requester"
	handlers "github.com/atljoseph/api.josephgill.io/route-handlers"
)

func getRoutes() Routes {
	return Routes{
		Route{
			Name:        "GetTestPathParamHandler",
			Method:      "GET",
			Pattern:     fmt.Sprintf("%s/test/{%s}", BaseURL, requester.TestNameKey),
			HandlerFunc: handlers.GetTestPathParamHandler,
		},
		Route{
			Name:        "GetTestQueryParamHandler",
			Method:      "GET",
			Pattern:     fmt.Sprintf("%s/test", BaseURL),
			HandlerFunc: handlers.GetTestQueryParamHandler,
		},
		Route{
			Name:        "GetTestError",
			Method:      "GET",
			Pattern:     fmt.Sprintf("%s/error", BaseURL),
			HandlerFunc: handlers.GetTestErrorHandler,
		},
		Route{
			Name:        "GetPhotoAlbums",
			Method:      "GET",
			Pattern:     fmt.Sprintf("%s/albums", BaseURL),
			HandlerFunc: handlers.GetPhotoAlbumsHandler,
		},
		Route{
			Name:        "GetPhotosByAlbumKey",
			Method:      "GET",
			Pattern:     fmt.Sprintf("%s/albums/{%s}", BaseURL, requester.PhotoAlbumIDKey),
			HandlerFunc: handlers.GetPhotosByAlbumKeyHandler,
		},
		Route{
			Name:        "PostPhotosByAlbumKey",
			Method:      "POST",
			Pattern:     fmt.Sprintf("%s/albums", BaseURL),
			HandlerFunc: handlers.PostPhotoAlbumsHandler,
		},
	}
}
