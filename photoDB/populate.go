package photoDB

import (
	"github.com/atljoseph/api.josephgill.io/apierr"
)

// PopulateDB populates the DB with data using the photoDB business logic
func PopulateDB() error {
	errTag := "photoDB.PopulateDB"

	// create a transaction
	txo, err := NewTxO("Test User")
	if err != nil {
		return apierr.Errorf(err, errTag, "open db transaction")
	}

	// TODO populate initial data with populateDB instead of migrate

	// commit transaction
	err = txo.Commit()
	if err != nil {
		return apierr.Errorf(err, errTag, "commit db transaction")
	}

	return nil
}
