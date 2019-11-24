package photoDB

import (
	"github.com/atljoseph/api.josephgill.io/apierr"
)

// PopulateDB populates the DB with data using the photoDB business logic
func PopulateDB() error {
	funcTag := "PopulateDB"

	logMessage(funcTag, "populate initial data")

	// create a transaction
	txo, err := NewTxO("Test User")
	if err != nil {
		return apierr.Errorf(err, funcTag, "open db transaction")
	}

	// TODO populate initial data with populateDB instead of migrate

	// commit transaction
	err = txo.Commit()
	if err != nil {
		return apierr.Errorf(err, funcTag, "commit db transaction")
	}

	return nil
}
