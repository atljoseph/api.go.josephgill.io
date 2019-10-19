package responder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendJSON sends a JSON response
func SendJSON(w http.ResponseWriter, responseObj interface{}) { // []byte) {
	errTag := "responder.SendJSON"

	// set response headers to be json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// initialize the response
	response := BaseResponse{
		Data: responseObj,
	}

	// // test error condition
	// SendError(w, "Test Error", http.StatusBadRequest)
	// return

	// no job is done until the paperwork is finished
	resJSON, err := json.Marshal(response.Data)
	if err != nil {
		SendError(w, fmt.Sprintf("%s: error marshalling response to json: %s", errTag, err), http.StatusBadRequest)
		return
	}

	// flush
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

// resBytes := new(bytes.Buffer)
// json.NewEncoder(resBytes).Encode(res)

// responder.SendJSON(w, resBytes.Bytes())
