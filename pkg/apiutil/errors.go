package apiutil

// Error is an error constant.
type Error string

// Error implements error.
func (e Error) Error() string { return string(e) }

// Common Errors
const (
	ErrNon200FromRemote     Error = "non-200 status code from remote"
	ErrNonResponseFromRetry Error = "non-http response from retry"
)
