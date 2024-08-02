package handlers

import (
	"golang-hotel-booking/pkg/config"
	"golang-hotel-booking/pkg/models"
	"golang-hotel-booking/pkg/render"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepository(app *config.AppConfig) {
	Repo = &Repository{App: app}
}

//func NewHandlers(r *Repository) {
//	Repo = r
//}

func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {

	remoteIp := req.RemoteAddr
	m.App.Sessions.Put(req.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(res, "home.page.gohtml", &models.TemplateData{})

}

func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "1.0.0"
	remoteIp := m.App.Sessions.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(res, "about.page.gohtml", &models.TemplateData{StringMap: stringMap})

}
func HandlerICon(w http.ResponseWriter, r *http.Request) {}

func (m *Repository) Reservation(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "makeres.page.gohtml", &models.TemplateData{})
}

func (m *Repository) Luxury(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "luxury.page.gohtml", &models.TemplateData{})
}
func (m *Repository) Roof(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "roof.page.gohtml", &models.TemplateData{})
}
func (m *Repository) Search(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "search.page.gohtml", &models.TemplateData{})
}
func (m *Repository) Contact(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "contact.page.gohtml", &models.TemplateData{})
}
