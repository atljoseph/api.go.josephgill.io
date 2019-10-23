package responder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendJSON sends a JSON response
func SendJSON(w http.ResponseWriter, responseObj interface{}) { // []byte) {
	errTag := "responder.SendJSON"
	statusCode := http.StatusOK

	// set response headers to be json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// initialize the response
	response := BaseResponse{
		Message:    "It's all good, man!",
		Data:       responseObj,
		StatusCode: statusCode,
	}

	// // test error condition
	// SendError(w, "Test Error", http.StatusBadRequest)
	// return

	sendJSONResponse(w, response, errTag)
}

// resBytes := new(bytes.Buffer)
// json.NewEncoder(resBytes).Encode(res)

// responder.SendJSON(w, resBytes.Bytes())

func sendJSONResponse(w http.ResponseWriter, response BaseResponse, errTag string) {
	// no job is done until the paperwork is finished
	resJSON, err := json.Marshal(response)
	if err != nil {
		SendJSONHttpError(w, http.StatusBadRequest, fmt.Sprintf("%s: error marshalling response to json: %s", errTag, err))
		return
	}

	// flush
	w.WriteHeader(response.StatusCode)
	w.Write(resJSON)
}
