package routes

import "net/http"

// Route is a struct representing a specific publicly-available route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	// PathParams  []string
	// VerifyJWT    bool
	// VerifyPerms  bool
	// VerifyAPIKey bool
	// VerifyServer bool
	// PostHandler  *PostHandler
}

// type PostHandler struct {
// 	Func  func(errImpl.LoggingResponseWriter, *http.Request, interface{})
// 	Extra interface{}
// }

// Routes is a type
type Routes []Route
