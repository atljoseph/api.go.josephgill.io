package routes

import (
	"fmt"

	"github.com/atljoseph/api.josephgill.io/handlers"
	"github.com/atljoseph/api.josephgill.io/requester"
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
			Name:        "GetPhotoAlbum",
			Method:      "GET",
			Pattern:     fmt.Sprintf("%s/album/{%s}", BaseURL, requester.PhotoAlbumIDKey),
			HandlerFunc: handlers.GetPhotoAlbumsHandler,
		},
	}
}
