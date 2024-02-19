package handlers

// import (
// 	"encoding/gob"
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"time"

// 	"github.com/ZhijiunY/booking-web/internal/config"
// 	"github.com/ZhijiunY/booking-web/internal/models"
// 	"github.com/ZhijiunY/booking-web/internal/render"
// 	"github.com/alexedwards/scs/v2"
// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// 	"github.com/justinas/nosurf"
// )

// var (
// 	app             config.AppConfig
// 	session         *scs.SessionManager
// 	pathToTemplates = "./../../templates"
// 	functions       = template.FuncMap{}
// )

// // from main.go file
// func getRoutes() http.Handler {
// 	// what am I going to put in the session
// 	// reservation-summary
// 	gob.Register(models.Reservation{})

// 	// Change this to true when in production
// 	app.InProduction = false

// 	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
// 	app.InfoLog = infoLog

// 	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
// 	app.ErrorLog = errorLog

// 	session = scs.New()
// 	session.Lifetime = 24 * time.Hour
// 	// cookie persist after the browser window is closed by the end user
// 	session.Cookie.Persist = true
// 	session.Cookie.SameSite = http.SameSiteLaxMode
// 	session.Cookie.Secure = app.InProduction

// 	app.Session = session

// 	tc, err := CreateTestTemplateCache()
// 	if err != nil {
// 		log.Fatal("cannot create template cache")
// 	}

// 	app.TemplateCache = tc
// 	app.UseCache = true

// 	repo := NewRepo(&app)
// 	NewHandlers(repo)

// 	render.NewRenderer(&app)

// 	// from routes.go file
// 	mux := chi.NewRouter()

// 	// go get 	"github.com/go-chi/chi/v5/middleware"
// 	mux.Use(middleware.Recoverer)

// 	// go get "github.com/justinas/nosurf"
// 	// mux.Use(NoSurf)
// 	mux.Use(SessionLoad)

// 	mux.Get("/", Repo.Home)
// 	mux.Get("/about", Repo.About)
// 	mux.Get("/generals", Repo.Generals)
// 	mux.Get("/majors", Repo.Majors)

// 	mux.Get("/search-availability", Repo.Availability)
// 	mux.Post("/search-availability", Repo.PostAvailability)
// 	mux.Post("/search-availability-json", Repo.AvailabilityJSON)

// 	mux.Get("/contact", Repo.Contact)

// 	mux.Get("/make-reservation", Repo.Reservation)
// 	mux.Post("/make-reservation", Repo.PostReservation)
// 	mux.Get("/reservation-summary", Repo.ReservationSummary)

// 	fileServer := http.FileServer(http.Dir("./static/"))
// 	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

// 	return mux

// }

// // from middleware.go file
// // NoSurf adds CSRF protection to all POST requests
// func NoSurf(next http.Handler) http.Handler {
// 	csrfHandler := nosurf.New(next)

// 	csrfHandler.SetBaseCookie(http.Cookie{
// 		HttpOnly: true,
// 		Path:     "/",
// 		//	put "var app config.AppConfig" outside of the main function
// 		Secure:   app.InProduction,
// 		SameSite: http.SameSiteLaxMode,
// 	})
// 	return csrfHandler
// }

// // SessionLoad loads and saves the session on every request
// func SessionLoad(next http.Handler) http.Handler {
// 	return session.LoadAndSave(next)
// }

// // from render.go file
// // CreateTestTemplateCache creates a template cache as a map
// func CreateTestTemplateCache() (map[string]*template.Template, error) {

// 	myCache := map[string]*template.Template{}

// 	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
// 	if err != nil {
// 		return myCache, err
// 	}

// 	for _, page := range pages {
// 		name := filepath.Base(page)
// 		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
// 		if err != nil {
// 			return myCache, err
// 		}

// 		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
// 		if err != nil {
// 			return myCache, err
// 		}

// 		if len(matches) > 0 {
// 			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
// 			if err != nil {
// 				return myCache, err
// 			}
// 		}

// 		myCache[name] = ts
// 	}

// 	return myCache, nil
// }
