package apiutil

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/blend/go-sdk/async"
	"github.com/blend/go-sdk/ex"
	"github.com/blend/go-sdk/logger"
	"github.com/blend/go-sdk/r2"
)

// New creates a new client.
//
// The purpose of this client is to wrap r2 calls with common asks such as retries.
func New(addr string, opts ...Option) Client {
	client := Client{
		Addr:           addr,
		Transport:      new(http.Transport),
		ResponseFilter: ResponseFilterInvalidStatusAsError,
	}
	for _, opt := range opts {
		opt(&client)
	}
	return client
}

// Option mutates an api client.
type Option func(*Client)

// OptLog sets the logger.
func OptLog(log logger.Log) Option {
	return func(c *Client) { c.Log = log }
}

// OptDebug sets the debugging flag.
func OptDebug(debug bool) Option {
	return func(c *Client) { c.Debug = debug }
}

// OptTracer sets the tracer.
func OptTracer(tracer r2.Tracer) Option {
	return func(c *Client) { c.Tracer = tracer }
}

// OptInterceptors sets the client interceptor to a chained variadic set of interceptors.
func OptInterceptors(interceptors ...async.Interceptor) Option {
	return func(c *Client) {
		c.Interceptor = async.Interceptors(append([]async.Interceptor{c.Interceptor}, interceptors...)...)
	}
}

// OptDefaults sets the default r2 options.
func OptDefaults(opts ...r2.Option) Option {
	return func(c *Client) {
		c.Defaults = append(c.Defaults, opts...)
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
		c.Transport = t
	}
}

// Client is a base type for implementing common api practices like retries.
type Client struct {
	// Addr is the remote address of the host.
	// It is passed to `r2.New(c.Addr)` and should be
	// the basis for all calls made by the client.
	Addr string
	// Log adds logging messages for calls made.
	Log logger.Log
	// Debug controls if we show verbose logging output
	// including the request and response bodies,
	// in logging output.
	Debug bool
	// Tracer integrates distributed tracing into calls.
	Tracer r2.Tracer
	// Defaults add base options to all outgoing requests.
	Defaults []r2.Option
	// Interceptor is an indirect before the output of Do ...
	Interceptor async.Interceptor
	// ResponseFilter is used to help manage retries, specifically to
	// change bad status codes (i.e. non-200s) to errors so that
	// they will be retried.
	ResponseFilter ResponseFilter
	// Transport holds the client transport that will be re-used
	// between requests.
	Transport *http.Transport
}

// Action implements async.Actioner and can be used to chain actioner calls.
//
// Primarily this lets the Client be used as part of a circuit breaker chain or
// for rate limiting etc.
func (c Client) Action(ctx context.Context, args interface{}) (interface{}, error) {
	var allOpts []r2.Option
	allOpts = append(allOpts, c.Defaults...)
	if c.Tracer != nil {
		allOpts = append(allOpts, r2.OptTracer(c.Tracer))
	}
	if c.Log != nil {
		if c.Debug {
			allOpts = append(allOpts, r2.OptLogResponseWithBody(c.Log))
		} else {
			allOpts = append(allOpts, r2.OptLogResponse(c.Log))
		}
	}
	if c.Transport != nil {
		allOpts = append(allOpts, r2.OptTransport(c.Transport))
	}
	allOpts = append(allOpts, r2.OptContext(ctx))
	if opts, ok := args.([]r2.Option); ok {
		allOpts = append(allOpts, opts...)
	}
	return c.filterResponse(r2.New(c.Addr, allOpts...).Do())
}

// Do sends the request.
func (c Client) Do(ctx context.Context, opts ...r2.Option) (*http.Response, error) {
	var res interface{}
	var err error
	if c.Interceptor != nil {
		res, err = c.Interceptor.Intercept(c).Action(ctx, opts)
	} else {
		res, err = c.Action(ctx, opts)
	}
	typedRes, ok := res.(*http.Response)
	if !ok {
		return nil, ex.New(ErrNonResponseFromRetry)
	}
	return c.filterResponse(typedRes, err)
}

// Discard returns the metadata for the response but discards the response body.
func (c Client) Discard(ctx context.Context, opts ...r2.Option) (*http.Response, error) {
	res, err := c.Do(ctx, opts...)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	_, err = io.Copy(io.Discard, res.Body)
	if err != nil {
		return nil, ex.New(err)
	}
	return res, nil
}

// Bytes returns the raw bytes of a given api call.
func (c Client) Bytes(ctx context.Context, opts ...r2.Option) (*http.Response, []byte, error) {
	res, err := c.Do(ctx, opts...)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	contents, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, ex.New(err)
	}
	return res, contents, nil
}

// JSON sends the request and reads the response into a given output variable as json.
func (c Client) JSON(ctx context.Context, output interface{}, opts ...r2.Option) (*http.Response, error) {
	res, err := c.Do(ctx, opts...)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(output); err != nil {
		return nil, ex.New(err)
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
