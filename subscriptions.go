package twitchkithelix

import (
	"context"
	"fmt"
)

// CheckSubscriptionRequest represents the data required to request if
// a user is subscribed or not to a broadcaster.
type CheckSubscriptionRequest struct {
	// Required
	BroadcasterID string `url:"broadcaster_id" json:"broadcaster_id"`

	// UserID is the unique identifier of the user you want to check.
	//
	// Required
	UserID string `url:"user_id" json:"user_id"`
}

// CheckSubscriptionResponse represents the response returned
type CheckSubscriptionResponse struct {
	// Data is a slice of SubscriptionData
	Data []SubscriptionData `json:"data"`
}

// SubscriptionData represents the data about a users subscription status to a broadcaster
type SubscriptionData struct {
	// BroadcasterID is the id of the braudcaster the user is subscribed to
	BroadcasterID string `json:"broadcaster_id"`

	// BroadcasterLogin is the login name of the braudcaster the user is subscribed to
	BroadcasterLogin string `json:"broadcaster_login"`

	// BroadcasterName is the username of the braudcaster the user is subscribed to
	BroadcasterName string `json:"broadcaster_name"`

	// Recieved the subscription as a gift from another user
	IsGift bool `json:"is_gift"`

	// Tier of the current subscription
	//
	// Tier 1 = "1000"
	// Tier 2 = "2000"
	// Tier 3 = "3000"
	Tier string `json:"tier"`

	// GifterID is the id of the user wo gifted the sub IF it was gifted
	//
	// Possibly Empty
	GifterID *string `json:"gifter_id,omitempty"`

	// GifterLogin is the login of the user wo gifted the sub IF it was gifted
	//
	// Possibly Empty
	GifterLogin *string `json:"gifter_login,omitempty"`

	// GifterName is the username of the user wo gifted the sub IF it was gifted
	//
	// Possibly Empty
	GifterName *string `json:"gifter_name,omitempty"`
}

func (c *Client) CheckSubscription(ctx context.Context, req CheckSubscriptionRequest) (*CheckSubscriptionResponse, error) {
	var resp CheckSubscriptionResponse
	// Example: 'https://api.twitch.tv/helix/subscriptions/user?broadcaster_id=149747285&user_id=141981764'
	query := fmt.Sprintf("subscriptions/user?broadcaster_id=%s&user_id=%s", req.BroadcasterID, req.UserID)
	err := c.doRequest(ctx, "GET", query, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
