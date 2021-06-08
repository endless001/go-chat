package middleware

import (
	"github.com/coreos/go-oidc/v3/oidc"
	"go-chat/internal/config"
	"net/http"
)

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		oauth2Token, err := config.OAuth2Config.Exchange(ctx, r.URL.Query().Get("code"))
		var verifier = config.Provider.Verifier(&oidc.Config{ClientID: config.ClientID})
		if err != nil {
			// handle error
		}

		// Extract the ID Token from OAuth2 token.
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			// handle missing token
		}

		// Parse and verify ID Token payload.
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			// handle error
		}

		// Extract custom claims
		var claims struct {
			Email    string `json:"email"`
			Verified bool   `json:"email_verified"`
		}
		if err := idToken.Claims(&claims); err != nil {
			// handle error
		}
	})
}
