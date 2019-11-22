package photoDB

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/atljoseph/api.josephgill.io/apierr"
	"github.com/jmoiron/sqlx"
)

// TxO is a transaction object that should always be passed along
// to database-interacting functions
type TxO struct {
	*sqlx.Tx
	Email string
	Pagination

	// Deadlock handling
	DeadlockErr   bool
	OriginHandler func(w http.ResponseWriter, r *http.Request)
	Req           *http.Request

	// Terminated flag
	Terminated bool
}

// Pagination helper
type Pagination struct {
	ApplyPagination bool
	Page            int
	PageRecords     int
	FoundRows       int
}

func handleDeadlockError(txo *TxO, err error) {
	if err != nil && strings.HasPrefix(err.Error(), "Error 1213") {
		txo.DeadlockErr = true
	}
}

// GetPagination looks at the *http.Request and verifies whether page and pageRecords are query parameters. If they are, pagination is applied
func GetPagination(r *http.Request) Pagination {

	var (
		page        int
		pageRecords int
		err         error
	)

	pagination := true

	// Get page
	if page, err = strconv.Atoi(r.FormValue("page")); err != nil {
		pagination = false
	}

	// Get number of result values to return
	if pageRecords, err = strconv.Atoi(r.FormValue("pageRecords")); err != nil {
		pagination = false
	}

	return Pagination{pagination, page, pageRecords, 0}
}

// NamedExec a named query within a transaction.
// Any named placeholder parameters are replaced with fields from arg.
func (txo *TxO) NamedExec(query string, arg interface{}) (r sql.Result, err error) {
	r, err = txo.Tx.NamedExec(query, arg)
	handleDeadlockError(txo, err)
	return
}

// Select within a transaction.
// Any placeholder parameters are replaced with supplied args.
func (txo *TxO) Select(dest interface{}, query string, args ...interface{}) (err error) {

	var newQ string
	if txo.Pagination.ApplyPagination {

		newQ = strings.Replace(query, "SELECT", "SELECT SQL_CALC_FOUND_ROWS", 1)
		newQ = newQ + fmt.Sprintf(" LIMIT %d,%d", (txo.Pagination.Page-1)*txo.Pagination.PageRecords, txo.Pagination.PageRecords)

	} else {
		newQ = query
	}

	err = txo.Tx.Select(dest, newQ, args...)

	handleDeadlockError(txo, err)

	if err != nil {
		return err
	}

	return txo.ApplyPagination()
}

// Query with a struct within a transaction.
// Any placeholder parameters are replaced with fields in the supplied arg
func (txo *TxO) NamedQuery(query string, arg interface{}) (result *sqlx.Rows, err error) {
	var newQ string

	if txo.Pagination.ApplyPagination {
		newQ = strings.Replace(query, "SELECT", "SELECT SQL_CALC_FOUND_ROWS", 1)
		newQ = newQ + fmt.Sprintf(" LIMIT %d,%d", (txo.Pagination.Page-1)*txo.Pagination.PageRecords, txo.Pagination.PageRecords)
	} else {
		newQ = query
	}
	result, err = txo.Tx.NamedQuery(newQ, arg)
	handleDeadlockError(txo, err)

	return result, txo.ApplyPagination()
}

func (txo *TxO) ApplyPagination() error {
	var err error
	if txo.Pagination.ApplyPagination {
		txo.Pagination.FoundRows, err = txo.GetFoundRows()
	}
	return err
}

// GetFoundRows gets all rows that were found without applying the LIMIT clause.
// Used for paginated queries
func (txo *TxO) GetFoundRows() (int, error) {
	q := "select FOUND_ROWS()"

	var rows int

	err := txo.Get(&rows, q)

	return rows, err
}

// Get within a transaction.
// Any placeholder parameters are replaced with supplied args.
// An error is returned if the result set is empty.
func (txo *TxO) Get(dest interface{}, query string, args ...interface{}) (err error) {
	err = txo.Tx.Get(dest, query, args...)
	handleDeadlockError(txo, err)
	return
}

// Exec a named query within a transaction.
func (txo *TxO) Exec(query string, args ...interface{}) (r sql.Result, err error) {
	r, err = txo.Tx.Exec(query, args...)
	handleDeadlockError(txo, err)
	return
}

// NewTxO initiates and returns a pointer to a TxO object.
// It takes in a *sqlx.DB object to begin the transaction from.
func NewTxO(email string) (*TxO, error) {
	errTag := "photoDB.NewTxO"

	tx, err := dbx.Beginx()
	if err != nil {
		return nil, apierr.Errorf(err, errTag, "error beginning transaction")
	}

	return &TxO{
		Tx:    tx,
		Email: email,
	}, nil
}

// NewTxOWithCTX returns a new TxO object that has a context baked into it
func NewTxOWithCTX(ctx context.Context, db *sqlx.DB, email string) (*TxO, error) {
	errTag := "photoDB.NewTxOWithCTX"

	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, apierr.Errorf(err, errTag, "unable to begin transaction")
	}

	return &TxO{
		Tx:    tx,
		Email: email,
	}, nil
}

// UpdateGetAffected executes a query and returns a slice of the values of the field of the rows affected.
func (txo *TxO) UpdateGetAffected(q, field, separator string, args ...interface{}) ([]string, error) {

	_, err := txo.Exec(`SET @uids := null;`)
	if err != nil {
		return nil, err
	}

	query := q + fmt.Sprintf(" AND ( SELECT @uids := CONCAT_WS('%s', %s, @uids) );", separator, field)
	_, err = txo.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	var v string
	err = txo.Get(&v, `SELECT CONVERT(@uids USING utf8)`)
	if err != nil {
		return nil, err
	}

	return strings.Split(v, separator), nil
}

// Rollback rolls back a transaction and sets the terminated flag to true on TxO
func (txo *TxO) Rollback() error {
	errTag := "TxO.Rollback"

	txo.Terminated = true
	err := txo.Tx.Rollback()
	if err != nil {
		return apierr.Errorf(err, errTag, "unable to rollback transaction")
	}
	return nil
}

// RollbackOnError rollsback the transaction if error
func (txo *TxO) RollbackOnError(err error) error {
	errTag := "TxO.RollbackOnError"

	// return if no error incoming
	if err == nil {
		return nil
	}

	// rollback, since error
	errR := txo.Rollback()
	if errR != nil {
		return apierr.Errorf(err, errTag, errR.Error())
	}

	// return original error if rollback was successful
	return apierr.Errorf(err, errTag, "rollback on query error")
}

// Commit commits a transaction and sets the terminated flag to true on TxO
func (txo *TxO) Commit() error {
	errTag := "TxO.Commit"

	txo.Terminated = true

	// commit
	// rollback on error
	err := txo.Tx.Commit()
	err = txo.RollbackOnError(err)
	if err != nil {
		return apierr.Errorf(err, errTag, "unable to commit transaction")
	}

	return nil
}
