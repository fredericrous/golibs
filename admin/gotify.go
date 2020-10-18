package admin

import (
	"fmt"
	"net/http"
	liburl "net/url"

	"github.com/qdm12/golibs/format"

	"github.com/gotify/go-api-client/v2/auth"
	"github.com/gotify/go-api-client/v2/client"
	"github.com/gotify/go-api-client/v2/client/message"
	gotifylib "github.com/gotify/go-api-client/v2/gotify"
	"github.com/gotify/go-api-client/v2/models"
)

//go:generate mockgen -destination=mock_$GOPACKAGE/$GOFILE . Gotify

// Gotify is a Gotify client
type Gotify interface {
	// Ping obtains silently the version from the Gotify server and returns an error on failure
	Ping() error
	// Notify formats and sends a message to the Gotify server
	Notify(title string, priority int, args ...interface{}) error
}

type gotify struct {
	client *client.GotifyREST
	token  string
}

// NewGotify creates an API client with the token for the Gotify server
func NewGotify(url liburl.URL, token string, httpClient *http.Client) Gotify {
	client := gotifylib.NewClient(&url, httpClient)
	return &gotify{client: client, token: token}
}

// Ping obtains silently the version from the Gotify server and returns an error on failure
func (g *gotify) Ping() error {
	if _, err := g.client.Version.GetVersion(nil); err != nil {
		return fmt.Errorf("cannot communicate with Gotify server: %w", err)
	}
	return nil
}

// Notify formats and sends a message to the Gotify server
func (g *gotify) Notify(title string, priority int, args ...interface{}) error {
	content := format.ArgsToString(args...)
	params := message.NewCreateMessageParams()
	params.Body = &models.MessageExternal{
		Title:    title,
		Message:  content,
		Priority: priority,
	}
	_, err := g.client.Message.CreateMessage(params, auth.TokenAuth(g.token))
	if err != nil {
		return fmt.Errorf("cannot send message to Gotify: %w", err)
	}
	return nil
}
