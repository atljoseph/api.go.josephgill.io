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

// GetTestPathParamHandler is a test endpoint that says hello
func GetTestPathParamHandler(w http.ResponseWriter, r *http.Request) {
	funcTag := "GetTestPathParamHandler"

	// process request params
	mp, err := requester.GetRequestParams(r, nil, requester.TestNameKey)
	if err != nil {
		err = apierr.Errorf(err, funcTag, "processing request params")
		responder.SendJSONError(w, http.StatusBadRequest, err)
		return
	}

	// respond
	res := &GetTestResponse{}
	res.Test = fmt.Sprintf("Hello, %s\n", mp[requester.TestNameKey])

	// return
	responder.SendJSON(w, res)
}

// GetTestQueryParamHandler is a test endpoint that says hello
func GetTestQueryParamHandler(w http.ResponseWriter, r *http.Request) {
	funcTag := "GetTestQueryParamHandler"

	pName := &requester.QueryParam{Name: requester.TestNameKey, DefaultValue: "Guest"}
	err := requester.GetQueryParams(r, pName)
	if err != nil {
		err = apierr.Errorf(err, funcTag, "getting query params")
		responder.SendJSONError(w, http.StatusBadRequest, err)
		return
	}

	// respond
	res := &GetTestResponse{}
	res.Test = fmt.Sprintf("Hello, %s\n", pName.Value)

	// return
	responder.SendJSON(w, res)
}

// GetTestErrorHandler is a test endpoint that returns an error
func GetTestErrorHandler(w http.ResponseWriter, r *http.Request) {
	funcTag := "GetTestErrorHandler"

	err := apierr.Errorf(nil, funcTag, "TEST")
	responder.SendJSONError(w, http.StatusBadRequest, err)
	return
}
