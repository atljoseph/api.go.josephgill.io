package handlers

import (
	"fmt"
	"net/http"

	"github.com/atljoseph/api.josephgill.io/apierr"

	"github.com/atljoseph/api.josephgill.io/requester"
	"github.com/atljoseph/api.josephgill.io/responder"
)

// GetTestResponse is the response returned by GetTestHandler
type GetTestResponse struct {
	Test string `json:"test"`
}

// GetTestPathParamHandler is just a test endpoint
func GetTestPathParamHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "GetTestPathParamHandler"

	// process request params
	mp, err := requester.Process(r, nil, requester.TestNameKey)
	if err != nil {
		err = apierr.Errorf(err, errTag, "processing request params")
		responder.SendJSONHttpError(w, http.StatusBadRequest, err)
		return
	}

	// respond
	res := &GetTestResponse{}
	res.Test = fmt.Sprintf("Hello, %s\n", mp[requester.TestNameKey])

	// return
	responder.SendJSON(w, res)
}

// GetTestQueryParamHandler is just a test endpoint
func GetTestQueryParamHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "GetTestQueryParamHandler"

	pName := &requester.QueryParam{Name: requester.TestNameKey, DefaultValue: "Guest"}
	err := requester.GetQueryParams(r, pName)
	if err != nil {
		err = apierr.Errorf(err, errTag, "getting query params")
		responder.SendJSONHttpError(w, http.StatusBadRequest, err)
		return
	}

	// respond
	res := &GetTestResponse{}
	res.Test = fmt.Sprintf("Hello, %s\n", pName.Value)

	// return
	responder.SendJSON(w, res)
}

// GetTestErrorHandler is just a test endpoint that returns an error
func GetTestErrorHandler(w http.ResponseWriter, r *http.Request) {
	errTag := "GetTestErrorHandler"

	err := apierr.Errorf(nil, errTag, "TEST")
	responder.SendJSONHttpError(w, http.StatusBadRequest, err)
	return
}
