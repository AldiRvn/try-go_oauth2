package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"oauth2/util"

	"github.com/kataras/golog"
)

func Google(w http.ResponseWriter, r *http.Request) {
	golog.Debug("Callback-gl is called")
	golog.Debug(r.URL)

	// ? See 'state' description in globalVar.go
	state := r.FormValue("state")
	if authState != state {
		golog.Debug(state)
		golog.Warn("Invalid oauth state.")

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// * Prepare the code
	// ? The authorization code is a one-time code that your server can exchange for an
	// ? access token. This access token is passed to the Gmail API to grant your
	// ? application access to user data for a limited time. ... This will
	// ? cause the user to see a dialog to grant permission to your
	// ? application again.
	// ? source: Authorization Code Grant - OAuth 2.0 Simplified
	// ? https://www.oauth.com/oauth2-servers/server-side-apps/authorization-code/
	code := r.FormValue("code")
	golog.Debug(code)
	if code == "" {
		golog.Warn("Code not found")

		util.WriterOneLine(w, "Code not found to provide AccessToken...\n")

		reason := r.FormValue("error_reason")
		if reason == "user_denied" {
			util.WriterOneLine(w, "User has denied Permission...\n")
		}
	} else {
		// * Get token using the code
		token, err := conf.Exchange(ctx, code)
		if err != nil {
			golog.Errorf("conf.Exchange:\n %v", err)
			return
		}

		// ? Print token content
		mapToken := map[string]interface{}{
			"access_token":    token.AccessToken,
			"expiration_time": token.Expiry,
			"refresh_token":   token.RefreshToken,
		}
		util.PrettyPrint(mapToken)

		// * Request user data using token
		getUserInfoURL := "https://www.googleapis.com/oauth2/v3/userinfo?"
		afterSetAccessToken := fmt.Sprintf(
			"%saccess_token=%s", getUserInfoURL, url.QueryEscape(token.AccessToken))
		resp, err := http.Get(afterSetAccessToken)
		if err != nil {
			golog.Errorf("http.Get:\n %v", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		defer resp.Body.Close()

		// * Read the response body
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			golog.Errorf("ioutil.ReadAll:\n %v", err)
		}

		mapUserData := util.JsonUnmarshalOneLine(response)
		util.PrettyPrint(mapUserData)
		util.WriteJsonByMap(w, mapUserData)
		return
	}
}
