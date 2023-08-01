package handlers

import (
	"net/http"

	"github.com/rahmanazizf/basicwgo/pkg/config"
	"github.com/rahmanazizf/basicwgo/pkg/renderer"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplate(w, "about.page.tmpl")
}
