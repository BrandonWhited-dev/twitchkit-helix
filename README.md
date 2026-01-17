# twitchkit-helix
"twitchkit-helix" is a Go library that enables users to interact with Twitch's Helix API, providing access to features such as sending chat messages, starting polls, creating clips, and more.

For detailed field info, see the [pkg.go.dev documentation](https://pkg.go.dev/github.com/brandonwhited-dev/twitchkit-helix).
# Features
- Strongly-typed Go wrappers for Twitch Helix endpoints
- Chat messaging 
- Poll creation
- Clip creation
- Search streams, channels, users, games, and categories
- Manage channel points & rewards
# Installation 
```bash
go get github.com/brandonwhited-dev/twitchkit-helix
```
# Example
See [`twitchkit-examples`](https://github.com/brandonwhited-dev/twitchkit-examples) for more examples of twitckhit.
```go
package main
import (
    "log"

    "github.com/brandonwhited-dev/twitchkit-helix"
)

func main () {
    oauth := "YOUR_TWITCH_OAUTH"       // no OAuth Prefix
    clientID := "YOUR_CLIENT_ID"       // Client ID linked with the oauth

	// ===Create Helix API Client===
	helixClient := twitchkithelix.NewClient(&clientID, &oauth, nil)

	// Send a hello message to twitch chat
	sourceOnly := true
	req := twitchkithelix.SendMessageRequest{
		BroadcasterID: "BroadcasterID",
		SenderID:      "SenderID",
		Message:       "Hello, Twitch!",
		ForSourceOnly: &sourceOnly,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = helixClient.SendMessage(ctx, req)
	if err != nil {
		log.Println("Send Message Erorr: ", err)
	}
}
```
