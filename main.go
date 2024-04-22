package main

import (
	"fmt"
	"net/http"
	"github.com/waytix-l/GroupieTracker/internal/handlers"
)

func main() {
	var appConfig config.Config

	templateCache, err := handlers.CreateTemplateCache()

	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8877"

	handlers.CreateTemplates(&appConfig)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/contact", handlers.Contact)

	http.HandleFunc("/algerie1", handlers.Algerie1)
	http.HandleFunc("/algerie2", handlers.Algerie2)
	http.HandleFunc("/algerie3", handlers.Algerie3)
	http.HandleFunc("/algerie4", handlers.Algerie4)

	http.HandleFunc("/hoth", handlers.Hoth)
	http.HandleFunc("/corruscant", handlers.Corruscant)
	http.HandleFunc("/korriban", handlers.Korriban)
	http.HandleFunc("/mustafar", handlers.Mustafar)

	http.HandleFunc("/hothvideo", handlers.HothVideo)
	http.HandleFunc("/corruscantvideo", handlers.CorruscantVideo)
	http.HandleFunc("/korribanvideo", handlers.KorribanVideo)
	http.HandleFunc("/mustafarvideo", handlers.MustafarVideo)


	fmt.Println("(htpp://localhost:8877) - Server started on port ", appConfig.Port)
	http.ListenAndServe(appConfig.Port, nil)
}
