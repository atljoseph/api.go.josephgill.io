package routes

import (
	"fmt"

	"github.com/atljoseph/api.josephgill.io/handlers"
	"github.com/atljoseph/api.josephgill.io/requester"
)

func getRoutes(routesConfig Config) Routes {
	routes := Routes{
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
			Name:        "GetPopulateDBHandler",
			Method:      "GET",
			Pattern:     fmt.Sprintf("%s/populate/photos", BaseURL),
			HandlerFunc: handlers.GetPopulateDBHandler,
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
	}
	if !routesConfig.IsProd {
		routes = append(routes,
			Route{
				Name:        "PostPhotoAlbum",
				Method:      "POST",
				Pattern:     fmt.Sprintf("%s/albums", BaseURL),
				HandlerFunc: handlers.PostPhotoAlbumHandler,
			},
			Route{
				Name:        "PostPhotoAlbumPhotoHandler",
				Method:      "POST",
				Pattern:     fmt.Sprintf("%s/albums/{%s}", BaseURL, requester.PhotoAlbumIDKey),
				HandlerFunc: handlers.PostPhotoAlbumPhotoHandler,
			},
		)
	}
	return routes
}
