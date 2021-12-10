package lametric

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wcharczuk/lametric/pkg/apiutil"
)

// New returns a new http client.
func New(addr, token string, opts ...apiutil.Option) *HTTPClient {
	hc := HTTPClient{
		Client: apiutil.New(fmt.Sprintf("http://%s:8080", addr),
			append(opts,
				apiutil.OptDefaults(
					apiutil.OptBasicAuth("dev", token),
					apiutil.OptHeader("Accept", "application/json"),
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
		apiutil.OptMethod(http.MethodPost),
		apiutil.OptPath("/api/v2/device/notifications"),
		apiutil.OptJSONBody(args),
	); err != nil {
		return nil, err
	}
	return &output, nil
}
