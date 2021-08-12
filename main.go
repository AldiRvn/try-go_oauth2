// ? list reference:
// ? https://pkg.go.dev/golang.org/x/oauth2#example-Config
// ? https://medium.com/@bnprashanth256/oauth2-with-google-account-gmail-in-go-golang-1372c237d25e
package main

import (
	"fmt"
	"net/http"
	"oauth2/controller"
	"oauth2/service"
	"os"

	"github.com/kataras/golog"
)

func init() {
	// ? run 'go run main.go debug'
	// ? to enable golog.SetLevel("debug")
	if len(os.Args) > 1 {
		golog.SetLevel(os.Args[1])
	}
}

func main() {
	const PORT = 57902

	http.HandleFunc("/", service.Base)
	http.HandleFunc("/login-gl", controller.LoginPage)
	http.HandleFunc("/callback-gl", controller.Google)

	golog.Debug("Debug mode is on.")
	golog.Infof("Run APP with {PORT:%d}", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		golog.Error(err)
	}
}
