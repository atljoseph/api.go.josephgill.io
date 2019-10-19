package responder

import (
	"fmt"
	"log"
	"net/http"
)

// SendError sends an error back to the client in a structured fashion
func SendError(w http.ResponseWriter, errText string, statusCode int) {

	// print message
	logTxt := fmt.Sprintf(
		"ERROR: %s\n",
		errText,
	)
	log.Println(logTxt)

	// respond with error
	http.Error(w, errText, statusCode)
}
