package handlers

import (
	"encoding/json"
	"fmt"
	"golang-hotel-booking/pkg/config"
	"golang-hotel-booking/pkg/models"
	"golang-hotel-booking/pkg/render"
	"log"
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
	render.RenderTemplate(res, "home.page.gohtml", &models.TemplateData{}, req)

}

func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "1.0.0"
	remoteIp := m.App.Sessions.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(res, "about.page.gohtml", &models.TemplateData{StringMap: stringMap}, req)

}
func HandlerICon(w http.ResponseWriter, r *http.Request) {}

func (m *Repository) Reservation(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "makeres.page.gohtml", &models.TemplateData{}, req)
}

func (m *Repository) Luxury(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "luxury.page.gohtml", &models.TemplateData{}, req)
}
func (m *Repository) Roof(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "roof.page.gohtml", &models.TemplateData{}, req)
}
func (m *Repository) Search(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "search.page.gohtml", &models.TemplateData{}, req)
}
func (m *Repository) Contact(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "contact.page.gohtml", &models.TemplateData{}, req)
}
func (m *Repository) PostSearch(res http.ResponseWriter, req *http.Request) {
	start := req.Form.Get("start_date")
	end := req.Form.Get("ending_date")

	res.Write([]byte(fmt.Sprintf("start: %s\nend: %s", start, end)))
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) SearchJson(res http.ResponseWriter, req *http.Request) {
	resp := jsonResponse{
		Ok:      false,
		Message: "available!",
	}
	out, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))
	res.Header().Set("Content-Type", "application/json")
	_, err = res.Write(out)
	if err != nil {
		return
	}
}
