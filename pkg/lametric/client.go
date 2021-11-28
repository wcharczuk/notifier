package lametric

import (
	"context"
)

// Client is the main interface for the api.
type Client interface {
	CreateNotification(context.Context, Notification) (*CreateNotificationOutput, error)
	GetNotifications(context.Context) ([]Notification, error)
}
