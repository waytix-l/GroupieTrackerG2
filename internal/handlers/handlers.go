package handlers

import (
	"GroupieTracker/api"
	"GroupieTracker/config"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

var Filters api.Filters

func Home(w http.ResponseWriter, r *http.Request) {
	Filters.MinCreationDate, Filters.Err = strconv.Atoi(r.FormValue("CreationDateMin"))
	Filters.MaxCreationDate, Filters.Err = strconv.Atoi(r.FormValue("CreationDateMax"))
	renderTemplate(w, r, "home", api.Filters{MinCreationDate: Filters.MinCreationDate, MaxCreationDate: Filters.MaxCreationDate})
}

var appConfig *config.Config

func CreateTemplates(app *config.Config) {
	appConfig = app
}

func renderTemplate(w http.ResponseWriter, r *http.Request, tmplName string, data interface{}) {
	templateCache := appConfig.TemplateCache

	tmpl, ok := templateCache[tmplName+".page.html"]
	fmt.Println(tmpl)

	if !ok {
		http.Error(w, "Le template n'existe pas", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("assets/templates_HTML/*.page.html")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layouts, err := filepath.Glob("assets/templates_HTML/layouts/*.layout.html")

		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			tmpl.ParseGlob("assets/templates_HTML/layouts/*.layout.html")
		}

		cache[name] = tmpl
	}

	return cache, nil
}
