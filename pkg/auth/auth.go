package auth

import (
	"crypto/rand"
	"fmt"
	"memoriesbot/pkg/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var conf *oauth2.Config = &oauth2.Config{}

func init() {
	c := config.Get()

	conf.ClientID = c.Auth.GooglePhotos.ClientID
	conf.ClientSecret = c.Auth.GooglePhotos.ClientSecret
	conf.RedirectURL = c.Auth.GooglePhotos.RedirectURL
	conf.Scopes = []string{
		"https://www.googleapis.com/auth/photoslibrary.readonly",
	}
	conf.Endpoint = google.Endpoint
}

// GenerateID - generate a unique ID to use as state
func GenerateID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	id := fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return id, nil
}

// RequestToken - to request a token
func RequestToken(state string) string {

	url := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return url
}

// ExchangeToken - exchange token with authorization code
func ExchangeToken(authCode string) (*oauth2.Token, error) {
	token, err := conf.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		return nil, err
	}

	return token, nil
}
