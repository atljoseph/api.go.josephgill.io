package handlers

import (
	"fmt"

	"github.com/atljoseph/api.josephgill.io/routes"
)

func Routes(isProd bool) routes.Routes {
	rs := routes.Routes{
		routes.Route{
			Name:        "GetTestPathParamHandler",
			Method:      "GET",
			Pattern:     fmt.Sprintf("%s/test/{%s}", routeBaseURL, routeKeyTestName),
			HandlerFunc: GetTestPathParamHandler,
		},
		routes.Route{
			Name:        "GetTestQueryParamHandler",
			Method:      "GET",
			Pattern:     fmt.Sprintf("%s/test", routeBaseURL),
			HandlerFunc: GetTestQueryParamHandler,
		},
		routes.Route{
			Name:        "GetTestError",
			Method:      "GET",
			Pattern:     fmt.Sprintf("%s/error", routeBaseURL),
			HandlerFunc: GetTestErrorHandler,
		},
		routes.Route{
			Name:        "GetPopulateDBHandler",
			Method:      "GET",
			Pattern:     fmt.Sprintf("%s/populate/photos", routeBaseURL),
			HandlerFunc: GetPopulateDBHandler,
		},
		routes.Route{
			Name:        "GetPhotoAlbums",
			Method:      "GET",
			Pattern:     fmt.Sprintf("%s/albums", routeBaseURL),
			HandlerFunc: GetPhotoAlbumsHandler,
		},
		routes.Route{
			Name:        "GetPhotosByAlbumKey",
			Method:      "GET",
			Pattern:     fmt.Sprintf("%s/albums/{%s}", routeBaseURL, routeKeyAlbumID),
			HandlerFunc: GetPhotosByAlbumKeyHandler,
		},
	}
	if !isProd {
		rs = append(rs,
			routes.Route{
				Name:        "PostPhotoAlbum",
				Method:      "POST",
				Pattern:     fmt.Sprintf("%s/albums", routeBaseURL),
				HandlerFunc: PostPhotoAlbumHandler,
			},
			routes.Route{
				Name:        "PostPhotoAlbumPhotoHandler",
				Method:      "POST",
				Pattern:     fmt.Sprintf("%s/albums/{%s}", routeBaseURL, routeKeyAlbumID),
				HandlerFunc: PostPhotoAlbumPhotoHandler,
			},
		)
	}
	return rs
}
