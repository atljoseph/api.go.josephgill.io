package apierr

// StackedError represents an error with somewhat of a stack trace
type StackedError struct {
	errorMessage string
	errorCode    int
}

// Error allows StackedError to satisfy the error interface
func (se StackedError) Error() string {
	return se.errorMessage
}

// ErrorCode returns the ErrorCode to be able to compare against errors being returned
func (se StackedError) ErrorCode() int {
	return se.errorCode
}
