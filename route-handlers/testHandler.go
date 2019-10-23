package routeHandlers

import (
	"fmt"
	"net/http"

	"github.com/atljoseph/api.josephgill.io/requester"
	"github.com/atljoseph/api.josephgill.io/responder"
)

// GetTestResponse is the response returned by GetTestHandler
type GetTestResponse struct {
	Test string `json:"test"`
}

// GetTestPathParamHandler is just a test endpoint
func GetTestPathParamHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "handlers.GetTestPathParamHandler"

	// process request params
	mp, err := requester.Process(r, nil, requester.TestNameKey)
	if err != nil {
		http.Error(w, fmt.Errorf("%s: %s", errTag, err).Error(), http.StatusBadRequest)
		return
	}

	// respond
	res := &GetTestResponse{}
	res.Test = fmt.Sprintf("Hello, %s\n", mp[requester.TestNameKey])
	responder.SendJSON(w, res)
}

// GetTestQueryParamHandler is just a test endpoint
func GetTestQueryParamHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "handlers.GetTestQueryParamHandler"

	pName := &requester.QueryParam{Name: requester.TestNameKey, DefaultValue: "Guest"}
	err := requester.GetQueryParams(r, pName)
	if err != nil {
		http.Error(w, fmt.Errorf("%s: %s", errTag, err).Error(), http.StatusBadRequest)
		return
	}

	// respond
	res := &GetTestResponse{}
	res.Test = fmt.Sprintf("Hello, %s\n", pName.Value)
	responder.SendJSON(w, res)
}

// GetTestErrorHandler is just a test endpoint that returns an error
func GetTestErrorHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "handlers.GetTestErrorHandler"

	responder.SendJSONHttpError(w, http.StatusBadRequest, fmt.Errorf("%s: TEST", errTag).Error())
	return
}
