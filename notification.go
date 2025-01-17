package pagerduty

import (
	"context"

	"github.com/google/go-querystring/query"
)

// Notification is a message containing the details of the incident.
type Notification struct {
	ID        string `json:"id"`
	Type      string
	StartedAt string `json:"started_at"`
	Address   string
	User      APIObject
}

// ListNotificationOptions is the data structure used when calling the ListNotifications API endpoint.
type ListNotificationOptions struct {
	APIListObject
	TimeZone string   `url:"time_zone,omitempty"`
	Since    string   `url:"since,omitempty"`
	Until    string   `url:"until,omitempty"`
	Filter   string   `url:"filter,omitempty"`
	Includes []string `url:"include,omitempty"`
}

// ListNotificationsResponse is the data structure returned from the ListNotifications API endpoint.
type ListNotificationsResponse struct {
	APIListObject
	Notifications []Notification
}

// ListNotifications lists notifications for a given time range, optionally
// filtered by type (sms_notification, email_notification, phone_notification,
// or push_notification).
//
// Deprecated: Use ListNotificationsWithContext instead.
func (c *Client) ListNotifications(o ListNotificationOptions) (*ListNotificationsResponse, error) {
	return c.ListNotificationsWithContext(context.Background(), o)
}

// ListNotificationsWithContext lists notifications for a given time range,
// optionally filtered by type (sms_notification, email_notification,
// phone_notification, or push_notification).
func (c *Client) ListNotificationsWithContext(ctx context.Context, o ListNotificationOptions) (*ListNotificationsResponse, error) {
	v, err := query.Values(o)
	if err != nil {
		return nil, err
	}

	resp, err := c.get(ctx, "/notifications?"+v.Encode())
	if err != nil {
		return nil, err
	}

	var result ListNotificationsResponse
	if err := c.decodeJSON(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
