package lametric

import (
	"context"
	"fmt"

	"github.com/blend/go-sdk/r2"
	"github.com/blend/go-sdk/webutil"

	"github.com/wcharczuk/notifier/pkg/apiutil"
)

// New returns a new http client.
func New(addr, token string, opts ...apiutil.Option) *HTTPClient {
	hc := HTTPClient{
		Client: apiutil.New(fmt.Sprintf("http://%s:8080", addr),
			append(opts,
				apiutil.OptDebug(true),
				apiutil.OptDefaults(
					r2.OptBasicAuth("dev", token),
					r2.OptHeaderValue(webutil.HeaderAccept, "application/json"),
				),
			)...,
		),
	}
	return &hc
}

// HTTPClient is a concrete implementation of Client.
type HTTPClient struct {
	apiutil.Client
}

// CreateNotification creates a notification.
func (hc HTTPClient) CreateNotification(ctx context.Context, args Notification) (*CreateNotificationOutput, error) {
	var output CreateNotificationOutput
	if _, err := hc.Client.JSON(ctx, &output,
		r2.OptPost(),
		r2.OptPath("/api/v2/device/notifications"),
		r2.OptJSONBody(args),
	); err != nil {
		return nil, err
	}
	return &output, nil
}
