package controller

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	// ? security - What is the purpose of the 'state' parameter in OAuth authorization request - Stack Overflow
	// ? https://stackoverflow.com/questions/26132066/what-is-the-purpose-of-the-state-parameter-in-oauth-authorization-request
	authState = "5T4t3"
	ctx       = context.Background()
	conf      = &oauth2.Config{
		ClientID:     "399227874160-36utc7jnd8ap1563t5k63r6vn98ijnhe.apps.googleusercontent.com",
		ClientSecret: "5Q15Z36nGJabKfen2vAv2WZ2",
		// ? list all scopes: https://developers.google.com/identity/protocols/oauth2/scopes
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint:    google.Endpoint,
		RedirectURL: "http://localhost:57902/callback-gl",
	}
)
