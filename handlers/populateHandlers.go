package handlers

import (
	"net/http"

	"github.com/atljoseph/api.josephgill.io/photoDB"

	"github.com/atljoseph/api.josephgill.io/apierr"

	"github.com/atljoseph/api.josephgill.io/responder"
)

// GetPopulateResponse is the response returned by GetPopulateDBHandler
type GetPopulateResponse struct {
	Message string `json:"populate"`
}

// GetPopulateDBHandler is a test endpoint that populates the database
func GetPopulateDBHandler(w http.ResponseWriter, r *http.Request) {
	funcTag := "GetPopulateDBHandler"

	err := photoDB.PopulateDB()
	if err != nil {
		err = apierr.Errorf(err, funcTag, "populating db")
		responder.SendJSONError(w, http.StatusBadRequest, err)
		return
	}

	// respond
	res := &GetPopulateResponse{}
	res.Message = "Done! Time to party like it is 1699! :)"

	// return
	responder.SendJSON(w, res)
}
