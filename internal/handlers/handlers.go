package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/waytix-l/HangmanWeb/config"
	"github.com/waytix-l/HangmanWeb/hangman"
)

var h hangman.HangmanData

func Home(w http.ResponseWriter, r *http.Request) {
	h.HangmanInit()
	renderTemplate(w, r, "home", nil)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "contact", nil)
}
func Algerie1(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "algerie1", nil)
}
func Algerie2(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "algerie2", nil)
}
func Algerie3(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "algerie3", nil)
}
func Algerie4(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "algerie4", nil)
}

func Hoth(w http.ResponseWriter, r *http.Request) {
	h.Gessed_letters = r.FormValue("letter")
	fmt.Println(h.Gessed_letters)
	h.Hangman()
	renderTemplate(w, r, "hoth", hangman.HangmanData{Word: h.Word, Mot_secret: h.Mot_secret, ListMotSecret: h.ListMotSecret, Win: h.Win, Lost: h.Lost, Try :h.Try})
}
func HothVideo(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "hothVideo", nil)
}

func Corruscant(w http.ResponseWriter, r *http.Request) {
	h.Gessed_letters = r.FormValue("letter")
	fmt.Println(h.Gessed_letters)
	h.Hangman()
	renderTemplate(w, r, "corruscant", hangman.HangmanData{Word: h.Word, Mot_secret: h.Mot_secret, ListMotSecret: h.ListMotSecret, Win: h.Win, Lost: h.Lost, Try :h.Try})
}
func CorruscantVideo(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "corruscantVideo", nil)
}

func Korriban(w http.ResponseWriter, r *http.Request) {
	h.Gessed_letters = r.FormValue("letter")
	fmt.Println(h.Gessed_letters)
	h.Hangman()
	renderTemplate(w, r, "korriban", hangman.HangmanData{Word: h.Word, Mot_secret: h.Mot_secret, ListMotSecret: h.ListMotSecret, Win: h.Win, Lost: h.Lost, Try :h.Try})
}
func KorribanVideo(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "korribanVideo", nil)
}

func Mustafar(w http.ResponseWriter, r *http.Request) {
	h.Gessed_letters = r.FormValue("letter")
	fmt.Println(h.Gessed_letters)
	h.Hangman()
	renderTemplate(w, r, "mustafar", hangman.HangmanData{Word: h.Word, Mot_secret: h.Mot_secret, ListMotSecret: h.ListMotSecret, Win: h.Win, Lost: h.Lost, Try :h.Try})
}
func MustafarVideo(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "mustafarVideo", nil)
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
