package main

import (
	"net/http"

	"github.com/ZhijiunY/booking-web/pkg/config"
	"github.com/ZhijiunY/booking-web/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	//
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// use Chi
	mux := chi.NewRouter()

	// go get 	"github.com/go-chi/chi/v5/middleware"
	mux.Use(middleware.Recoverer)

	// go get "github.com/justinas/nosurf"
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
