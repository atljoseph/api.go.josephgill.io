package responder

import (
	"encoding/json"
	"net/http"

	"github.com/atljoseph/api.josephgill.io/apierr"
)

// SendJSON sends a JSON response
func SendJSON(w http.ResponseWriter, responseObj interface{}) { // []byte) {
	funcTag := "SendJSON"
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

	sendJSONResponse(w, response, funcTag)
}

// resBytes := new(bytes.Buffer)
// json.NewEncoder(resBytes).Encode(res)

// responder.SendJSON(w, resBytes.Bytes())

func sendJSONResponse(w http.ResponseWriter, response BaseResponse, funcTag string) {
	// no job is done until the paperwork is finished
	resJSON, err := json.Marshal(response)
	if err != nil {
		err = apierr.Errorf(err, funcTag, "error marshalling response to json")
		SendJSONHttpError(w, http.StatusBadRequest, err)
		return
	}

	// flush
	w.WriteHeader(response.StatusCode)
	w.Write(resJSON)
}
