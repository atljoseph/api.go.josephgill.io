package responder

import (
	"net/http"
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
