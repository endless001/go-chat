package config

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"log"
	"os"
)

var (
	Conf = &OAuth2Config{}
	ClientID     = os.Getenv("GOOGLE_OAUTH2_CLIENT_ID")
	ClientSecret = os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET")
)


type OAuth2Config struct {
	Config   *oauth2.Config
	Provider *oidc.Provider

}

func init() {
	provider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://127.0.0.1:5556/auth/google/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	c := OAuth2Config{}
	c.Config = &config
	c.Provider = provider
	Conf = &c
}
