package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ZhijiunY/booking-web/pkg/config"
	"github.com/ZhijiunY/booking-web/pkg/models"
	"github.com/justinas/nosurf"
)

var (
	// "undefined app = a" is most likely due to the variable "app" not being defined in the package-level scope.
	app       *config.AppConfig
	functions = template.FuncMap{}
)

// NewTemplate sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a

}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {

		tc, _ = CreateTemplateCache()
		// var err error
		// tc, err = CreateTemplateCache()
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get the template from the template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
