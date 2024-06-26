package main

import (
	"GroupieTracker/api"
	"GroupieTracker/config"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

var Filters api.Filters

func Home(w http.ResponseWriter, r *http.Request) {

	APIdata := api.ImportAPI()

	Filters.MinCreationDate, Filters.Err = strconv.Atoi(r.FormValue("CreationDateMin"))
	Filters.MaxCreationDate, Filters.Err = strconv.Atoi(r.FormValue("CreationDateMax"))
	Filters.Image = APIdata[0].Image
	Filters.Name = APIdata[0].Name
	Filters.Name = strings.ToUpper(Filters.Name)
	renderTemplate(w, r, "home", api.Filters{MinCreationDate: Filters.MinCreationDate, MaxCreationDate: Filters.MaxCreationDate, Image: Filters.Image, Name: Filters.Name})
}

func Artist(w http.ResponseWriter, r *http.Request) {
	APIdata := api.ImportAPI()
	GroupieList := []api.Groupie{}

	for _, Groupie := range APIdata {
		GroupieList = append(GroupieList, Groupie)
	}

	firstAlbumDate := r.FormValue("FirstAlbum")
	fmt.Println(firstAlbumDate)
	membersAmountCheckbox := r.FormValue("membersAmountCheckbox")
	fmt.Println(membersAmountCheckbox)
	creationDateChecked := r.FormValue("CreationDateCheckbox")
	fmt.Println(creationDateChecked)

	if firstAlbumDate != "" {
		firstAlbum, _ := strconv.Atoi(firstAlbumDate)
		GroupieList = api.FirstAlbum(firstAlbum, GroupieList)
	}
	
	if membersAmountCheckbox == "on" {
		membersAmount := r.FormValue("members")
		if membersAmount != "" {
			members, _ := strconv.Atoi(membersAmount)
			GroupieList = api.MembersAmount(members, GroupieList)
		}
	}

	

	if creationDateChecked == "on" {
		minValueCD, _ := strconv.Atoi(r.FormValue("CreationDateMin"))
		maxValueCD, _ := strconv.Atoi(r.FormValue("CreationDateMax"))
		GroupieList = api.CreationDate(minValueCD, maxValueCD, GroupieList)
	}

	GroupieList = api.Search(r.FormValue("search"), GroupieList)
	
	renderTemplate(w, r, "artist", api.Filters{Groupies: GroupieList})
}

func Concert(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "concert", api.Filters{})
}

func Date(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "date", api.Filters{})
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
