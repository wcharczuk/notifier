package lametric

import "context"

type Client interface {
	CreateNotification(context.Context, CreateNotificationInput) (*CreateNotificationOutput, error)
}

// CreateNotificationInput is the input for CreateNotification.
type CreateNotificationInput struct {
	Model DeviceModel `json:"model"`
}

// CreateNotificationOutput is the output for CreateNotification.
type CreateNotificationOutput struct {
	Success Identifier `json:"success"`
}

// Identifier is an identifier.
type Identifier struct {
	ID string `json:"id"`
}

// DeviceModel is a component.
type DeviceModel struct {
	Frames []Frame `json:"frames"`
}

// Frame is a component.
type Frame struct {
	Icon int    `json:"int"`
	Text string `json:"text"`
}
