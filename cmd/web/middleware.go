package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Writing to Console handler")
		next.ServeHTTP(w, r)
	})
}

// NoSurf is a middleware function that protects against Cross-Site Request Forgery (CSRF) attacks
// by wrapping the next handler with a CSRF protection handler.
func NoSurf(next http.Handler) http.Handler {
	// Create a new CSRF handler wrapping the next handler
	csrfHandler := nosurf.New(next)
	// Set the base cookie options for the CSRF token
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,                 // The cookie is accessible only through the HTTP protocol, not via JavaScript
		Path:     "/",                  // The cookie is valid for the entire site
		Secure:   app.IsProduction,     // The cookie is sent only over HTTPS if the app is in production mode
		SameSite: http.SameSiteLaxMode, // The cookie is sent with same-site requests and top-level navigations
	})
	// Return the CSRF handler
	return csrfHandler
}

// SessionLoad is a middleware function that loads and saves session data
// before and after calling the next handler in the chain.
func SessionLoad(next http.Handler) http.Handler {
	// Return the next handler wrapped with session loading and saving
	return session.LoadAndSave(next)
}
