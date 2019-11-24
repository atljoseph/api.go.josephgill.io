package requester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/atljoseph/api.josephgill.io/apierr"
	"github.com/gorilla/mux"
)

// GetRequestParams is used to parse generic requests into a query string param map & pointer variable
func GetRequestParams(r *http.Request, ptrObjToPopulate interface{}, pathKeys ...string) (map[string]string, error) {
	funcTag := "GetRequestParams"

	logMessage(funcTag, "processing request paramaters")

	// get muxvars
	mp, err := buildReqVars(r, pathKeys)
	if err != nil {
		return nil, apierr.Errorf(err, funcTag, "error getting mux vars from request")
	}

	// decode json into passed in pointer object
	if _, err := decodeJSON(r, ptrObjToPopulate); err != nil {
		return nil, apierr.Errorf(err, funcTag, "cannot decode json")
	}

	return mp, nil
}

// decodeJSON decodes *http.Request body into obj which needs to be a pointer
func decodeJSON(r *http.Request, obj interface{}) ([]byte, error) {
	funcTag := "decodeJSON"

	if obj == nil {
		return nil, nil
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return nil, apierr.Errorf(err, funcTag, "cannot decode json")
	}

	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	/*
		if err := r.Body.Close(); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Warn("Error closing request body")

			return nil, pipoerr.NewCannotCloseRequestBodyError(tag, err)
		}
	*/

	// Unmarshal body into obj interface
	if err := json.Unmarshal(body, obj); err != nil {
		return nil, apierr.Errorf(err, funcTag, "error marshaling JSON to struct")
	}

	return body, nil
}

func buildReqVars(r *http.Request, keys []string) (map[string]string, error) {
	funcTag := "buildReqVars"

	// make a data recepticle
	mp := make(map[string]string, len(keys))

	// get the requested path params sent from the client
	muxVars := mux.Vars(r)

	// loop through all provided keys and build a map of their values
	for _, key := range keys {

		// get the value sent from the client
		newVal := muxVars[key]
		if newVal == "" {
			err := fmt.Errorf("no value for required path parameter")
			return nil, apierr.Errorf(err, funcTag, "key: %s", key)
		}

		// sanitized the value
		newValWithSlashes := strings.Replace(newVal, "_slash_", "/", -1)

		// save
		mp[key] = newValWithSlashes
	}

	// facilityID and email are added by default
	// mp[facilityIDKey] = muxVars[facilityIDKey]
	// mp[emailKey] = apiutils.GetUserEmail(r)

	return mp, nil
}
