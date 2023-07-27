package handlers

import (
	"net/http"

	"github.com/rahmanazizf/basicwgo/cmd/pkg/renderer"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplate(w, "about.page.tmpl")
}
