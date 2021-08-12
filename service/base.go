package service

import (
	"net/http"
	"oauth2/pages"

	"github.com/kataras/golog"
)

func Base(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(pages.Index)); err != nil {
		golog.Error(err)
	}
}
