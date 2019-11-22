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
		// TODO // Remove for PROD until ready
		Route{
			Name:        "PostPhotoAlbum",
			Method:      "POST",
			Pattern:     fmt.Sprintf("%s/albums", BaseURL),
			HandlerFunc: handlers.PostPhotoAlbumHandler,
		},
		// TODO // Remove for PROD until ready
		Route{
			Name:        "PostPhotoAlbumPhotoHandler",
			Method:      "POST",
			Pattern:     fmt.Sprintf("%s/albums/{%s}", BaseURL, requester.PhotoAlbumIDKey),
			HandlerFunc: handlers.PostPhotoAlbumPhotoHandler,
		},
	}
}
