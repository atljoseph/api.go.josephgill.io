package responder

import (
	"net/http"
)

// SendHTTPError sends an error back as text
func SendHTTPError(w http.ResponseWriter, statusCode int, errText string) {
	funcTag := "SendJSONError"

	// log message
	pkgLog.WithFunc(funcTag).WithErrorMessage(errText).WithStatusCode(statusCode).Error()

	// respond with error
	http.Error(w, errText, statusCode)
}

// SendJSONError sends an error back as a JSON
func SendJSONError(w http.ResponseWriter, statusCode int, stackedError error) {
	funcTag := "SendJSONError"

	// log message
	pkgLog.WithFunc(funcTag).WithError(stackedError).WithStatusCode(statusCode).Error()

	// set response headers to be json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// initialize the response
	response := BaseResponse{
		Message:    "Whoops, something screwed up :/",
		StatusCode: statusCode,
		Error:      stackedError.Error(),
		IsError:    true,
	}

	sendJSONResponse(w, response, funcTag)
}
