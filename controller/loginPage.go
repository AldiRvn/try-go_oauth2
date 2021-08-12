package controller

import (
	"net/http"

	"github.com/kataras/golog"
	"golang.org/x/oauth2"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	// * Redirect user to consent page to ask for permission
	// * for the scopes.
	// ? look globalVar.go(in &oauth2.Config{}) to see the
	// ? scopes or other config declaration.
	url := conf.AuthCodeURL(authState, oauth2.AccessTypeOffline)
	// ? google api - What does "offline" access in OAuth mean? - Stack Overflow
	// ? https://stackoverflow.com/questions/30637984/what-does-offline-access-in-oauth-mean

	golog.Debugf("Visiting the URL for the auth dialog:\n %v", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
