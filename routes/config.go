package routes

import (
	"net/http"

	"github.com/atljoseph/api.josephgill.io/requester"
	"github.com/gorilla/mux"
)

// Configure sets up the router for the api
func Configure(isProd bool) (*mux.Router, error) {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range getRoutes(isProd) {
		var handler http.Handler

		// Last thing that will be executed is the actual handler function
		handler = route.HandlerFunc

		// if route.VerifyPerms {
		// Before that we will verify permissions
		// handler = VerifyPermissions(handler)
		// }

		// if route.VerifyJWT {
		// 	// Before THAT we will validate the token passed
		// 	handler = jwtMiddleware.Handler(handler)
		// }

		// // add context params list
		// handler = requester.HandlerWithContext(handler, route.PathParams)

		// And even before THAT we will start the logger to be able to calculate how everything takes.
		handler = requester.HandleWithLogging(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router, nil
}
