package apiutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// RequestOption mutates a request.
type RequestOption func(*http.Request) error

// OptMethod sets the method for the request.
func OptMethod(method string) RequestOption {
	return func(req *http.Request) error {
		req.Method = method
		return nil
	}
}

// OptPath sets the path for the request.
func OptPath(path string) RequestOption {
	return func(req *http.Request) error {
		req.URL.Path = path
		return nil
	}
}

// OptPathf sets the path for the request by a given format and args.
func OptPathf(format string, args ...interface{}) RequestOption {
	return func(req *http.Request) error {
		req.URL.Path = fmt.Sprintf(format, args...)
		return nil
	}
}

// OptBasicAuth sets the basic auth header.
func OptBasicAuth(username, password string) RequestOption {
	return func(req *http.Request) error {
		req.SetBasicAuth(username, password)
		return nil
	}
}

// OptHeader sets a header value.
func OptHeader(key, value string) RequestOption {
	return func(req *http.Request) error {
		if req.Header == nil {
			req.Header = make(http.Header)
		}
		req.Header.Set(key, value)
		return nil
	}
}

// OptJSONBody sets the json body.
func OptJSONBody(body interface{}) RequestOption {
	return func(req *http.Request) error {
		contents, err := json.Marshal(body)
		if err != nil {
			return err
		}
		if req.Header == nil {
			req.Header = make(http.Header)
		}
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		req.Body = io.NopCloser(bytes.NewReader(contents))
		return nil
	}
}
