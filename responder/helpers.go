package responder

import (
	"encoding/json"
	"net/http"

	"github.com/atljoseph/api.josephgill.io/apierr"
)

// sendJSONResponse is used internall to send a JSON to the client
func sendJSONResponse(w http.ResponseWriter, response BaseResponse, funcTag string) {
	// no job is done until the paperwork is finished
	resJSON, err := json.Marshal(response)
	if err != nil {
		err = apierr.Errorf(err, funcTag, "error marshalling response to json")
		SendJSONError(w, http.StatusBadRequest, err)
		return
	}

	// flush
	w.WriteHeader(response.StatusCode)
	w.Write(resJSON)
}

// resBytes := new(bytes.Buffer)
// json.NewEncoder(resBytes).Encode(res)

// responder.SendJSON(w, resBytes.Bytes())
