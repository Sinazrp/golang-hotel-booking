package handlers

import (
	"encoding/json"
	"fmt"
	"golang-hotel-booking/internal/config"
	"golang-hotel-booking/internal/forms"
	"golang-hotel-booking/internal/models"
	"golang-hotel-booking/internal/render"
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
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(res, "makeres.page.gohtml", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	}, req)
}

// posting reservation Form
func (m *Repository) PostReservation(res http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: req.Form.Get("first_name"),
		LastName:  req.Form.Get("last_name"),
		Email:     req.Form.Get("email"),
		Phone:     req.Form.Get("phone"),
	}
	form := forms.New(req.PostForm)
	//form.Has("first_name", req)
	//form.Has("last_name", req)
	//form.Has("email", req)
	//form.Has("phone", req)
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3)
	form.MaxLength("phone", 12)
	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(res, "makeres.page.gohtml", &models.TemplateData{
			Form: form,
			Data: data,
		}, req)
		return
	}
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
	start := req.Form.Get("start")
	out, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))
	log.Println(string(start))
	res.Header().Set("Content-Type", "application/json")
	_, err = res.Write(out)
	if err != nil {
		return
	}
}
