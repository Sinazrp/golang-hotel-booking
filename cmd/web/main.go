package main

import (
	"fmt"
	"golang-hotel-booking/internal/config"
	"golang-hotel-booking/internal/handlers"
	"golang-hotel-booking/internal/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.IsProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction
	app.Sessions = session
	app.UseCache = true

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error initializing template cache")
	}
	app.TemplateCache = tc
	handlers.NewRepository(&app)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	//http.HandleFunc("/favicon.ico", handlers.HandlerICon)

	fmt.Println(fmt.Sprintf("Listening onn port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server")
	}

}
