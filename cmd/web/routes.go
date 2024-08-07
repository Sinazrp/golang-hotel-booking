package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang-hotel-booking/pkg/config"
	"golang-hotel-booking/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//mux.Get("/home", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Use(WriteToConsole)
	mux.Get("/home", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/roof", handlers.Repo.Roof)
	mux.Get("/luxury", handlers.Repo.Luxury)
	mux.Get("/search", handlers.Repo.Search)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/makeres", handlers.Repo.Reservation)
	mux.Post("/search", handlers.Repo.PostSearch)
	mux.Get("/search-json", handlers.Repo.SearchJson)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	return mux

}
