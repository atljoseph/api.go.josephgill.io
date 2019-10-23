package responder

import (
	"fmt"
	"log"
	"net/http"
)

// TODO: Implement error tracing/handling for http responder

// SendHttpError sends an error back to the client in a structured fashion
func SendHttpError(w http.ResponseWriter, statusCode int, errText string) {

	// print message
	logTxt := fmt.Sprintf(
		"ERROR: %s\n",
		errText,
	)
	log.Println(logTxt)

	// respond with error
	http.Error(w, errText, statusCode)
}

// SendJSONHttpError sends an error back to the client in a structured fashion
func SendJSONHttpError(w http.ResponseWriter, statusCode int, errText string) {
	errTag := "responder.SendJSONError"

	// print message
	log.Println("ERROR", errText)

	// set response headers to be json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// initialize the response
	response := BaseResponse{
		Message:    "Whoops, something screwed up :/",
		StatusCode: statusCode,
		Error:      errText,
		IsError:    true,
	}

	sendJSONResponse(w, response, errTag)
}
