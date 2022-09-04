package handlers

import (
	"net/http"

	"github.com/go-rohan/go-bookings/pkg/config"
	"github.com/go-rohan/go-bookings/pkg/models"
	"github.com/go-rohan/go-bookings/pkg/render"
)

// holds the data sent from handlers and send it to the template

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

	remoteIp := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	// m.App.Session

	stringMap["test"] = "Hello again"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
