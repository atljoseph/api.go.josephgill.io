package photoDB

// GetDB gets the *sqlx.DB reference, or returns an error if Initialize failed
// Should call this method both internally or externally if connection reference is needed
// func GetDB() (*sqlx.DB, error) {
// 	errTag := "photoDB.GetDB"

// 	// if error, return the error from Initialize
// 	if err != nil {
// 		return nil, err
// 	}

// 	// if initialization succeeded for some reason, but the dbx object is empty, return error
// 	if dbx == nil {
// 		return nil, fmt.Errorf("%s: %s", errTag, "photoDB is not initialized")
// 	}

// 	return dbx, nil
// }
