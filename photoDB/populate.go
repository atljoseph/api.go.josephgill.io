package photoDB

import "fmt"

// PopulateDB populates the DB with data using the photoDB business logic
func PopulateDB() error {
	errTag := "photoDB.PopulateDB"

	// create a transaction
	txo, err := NewTxO("Test User")
	if err != nil {
		return fmt.Errorf("%s: %s", errTag, err)
	}

	// TODO populate initial data with populateDB instead of migrate

	// commit transaction
	err = txo.Commit()
	if err != nil {
		return fmt.Errorf("%s: %s", errTag, err)
	}

	return nil
}
