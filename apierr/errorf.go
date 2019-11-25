package apierr

import (
	"fmt"
)

// Errorf composes or adds to a StackedError
func Errorf(err error, funcTag, errString string, args ...interface{}) error {
	errString = formatErrorString(errString, args)

	// assert error to type StackedError
	// if failure, then return new stacked error
	stackedError, ok := err.(StackedError)
	if !ok {
		stackedError = StackedError{
			errorMessage: composeStackMsg(err, funcTag, errString),
		}
	} else {
		// if error is already stacked error, then compose a new message
		stackedError.errorMessage = composeStackMsg(fmt.Errorf(stackedError.errorMessage), funcTag, errString)
	}

	// log error
	pkgLog.WithFunc(funcTag).WithError(err).Warn()

	return stackedError
}

func formatErrorString(errString string, args []interface{}) string {
	es := errString
	if len(args) != 0 {
		// this if check is to make sure that the error string does not
		// get mangled, which happens when you run sprintf without any
		// args.
		es = fmt.Sprintf(es, args)
	}
	return es
}

func composeStackMsg(err error, funcTag, msg string) string {
	if err == nil {
		err = fmt.Errorf("UnknownError")
	}
	if len(msg) == 0 {
		msg = "UnspecifiedMessage"
	}
	if len(funcTag) == 0 {
		funcTag = "UnspecifiedFunc"
	}
	// err always needs to come last
	return fmt.Sprintf("%s (%s) --> %s", funcTag, msg, err)
}
