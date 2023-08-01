package handlers

import (
	"net/http"

	"github.com/rahmanazizf/basicwgo/pkg/config"
	"github.com/rahmanazizf/basicwgo/pkg/models"
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
	// implement some logic here
	stringMap := map[string]string{"random": "hello from handlers"}
	renderer.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	renderer.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
