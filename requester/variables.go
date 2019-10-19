package requester

const (

	// TestNameKey is a key used to send a query string prarameter for testing
	TestNameKey = "name"

	// PhotoAlbumIDKey is a key used to represent a variable-valued fragment of a request's URL
	PhotoAlbumIDKey = "album"
)

type contextKey string

func (c contextKey) String() string {
	return "mypackage context key " + string(c)
}

var (
	contextPathParamsKey = contextKey("path-params")
)
