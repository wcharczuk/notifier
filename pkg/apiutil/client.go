package apiutil

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// New creates a new client.
//
// The purpose of this client is to wrap r2 calls with common asks such as retries.
func New(addr string, opts ...Option) Client {
	client := Client{
		URL:            mustParseURL(addr),
		ResponseFilter: InvalidHTTPStatusAsError,
		Client:         &http.Client{},
	}
	for _, opt := range opts {
		opt(&client)
	}
	return client
}

// Option mutates an api client.
type Option func(*Client)

// OptLog sets the logger.
func OptLog(log Logger) Option {
	return func(c *Client) { c.Log = log }
}

// OptDefaults sets the default options for the client.
func OptDefaults(defaults ...RequestOption) Option {
	return func(c *Client) {
		c.Defaults = defaults
	}
}

// OptResponseFilter sets the client response filter.
func OptResponseFilter(rf ResponseFilter) Option {
	return func(c *Client) {
		c.ResponseFilter = rf
	}
}

// OptTransport sets the http transport for the client.
func OptTransport(t *http.Transport) Option {
	return func(c *Client) {
		c.Client.Transport = t
	}
}

// Client is a base type for implementing common api practices like retries.
type Client struct {
	URL            *url.URL
	Log            Logger
	ResponseFilter ResponseFilter
	Defaults       []RequestOption
	Client         *http.Client
}

// Do sends the request.
func (c Client) Do(ctx context.Context, opts ...RequestOption) (res *http.Response, err error) {
	req := &http.Request{
		Method:     http.MethodGet,
		URL:        c.URL,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       c.URL.Host,
	}
	req = req.Clone(ctx)
	for _, defaultOption := range c.Defaults {
		if err = defaultOption(req); err != nil {
			return nil, err
		}
	}
	for _, opt := range opts {
		if err = opt(req); err != nil {
			return nil, err
		}
	}
	if c.Client != nil {
		res, err = c.Client.Do(req)
	} else {
		res, err = http.DefaultClient.Do(req)
	}
	return c.filterResponse(res, err)
}

// Discard returns the metadata for the response but discards the response body.
func (c Client) Discard(ctx context.Context, opts ...RequestOption) (*http.Response, error) {
	res, err := c.Do(ctx, opts...)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	_, err = io.Copy(io.Discard, res.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Bytes returns the raw bytes of a given api call.
func (c Client) Bytes(ctx context.Context, opts ...RequestOption) (*http.Response, []byte, error) {
	res, err := c.Do(ctx, opts...)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	contents, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	return res, contents, nil
}

// JSON sends the request and reads the response into a given output variable as json.
func (c Client) JSON(ctx context.Context, output interface{}, opts ...RequestOption) (*http.Response, error) {
	res, err := c.Do(ctx, opts...)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(output); err != nil {
		return nil, err
	}
	return res, nil
}

func (c Client) filterResponse(res *http.Response, err error) (*http.Response, error) {
	if c.ResponseFilter != nil {
		return c.ResponseFilter(res, err)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}
