package apiutil

import (
	"net/http"
)

// ResponseFilter mutates a response, handling things like status codes as errors.
type ResponseFilter func(*http.Response, error) (*http.Response, error)

// ResponseFilterInvalidStatusAsError translates status codes into errors if they're outside the acceptable range (200-299).
//
// This can be used as a Client `ResponseFilter` to retry on bad status codes.
func ResponseFilterInvalidStatusAsError(res *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return res, err
	}
	if statusCode := res.StatusCode; statusCode < 200 || statusCode > 299 {
		if res.Request != nil && res.Request.URL != nil {
			return res, ErrNon200FromRemote
		}
		return res, ErrNon200FromRemote
	}
	return res, nil
}
