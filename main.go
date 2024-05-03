package main

import (
	"fmt"
	"net/http"
	"GroupieTracker/config"
	"GroupieTracker/internal/handlers"
	"GroupieTracker/api"
)

func main() {
	var appConfig config.Config

	APIdata := importAPI()

	fmt.Println(APIdata)
	fmt.Println()
	fmt.Println(APIdata[0])
	fmt.Println()
	fmt.Println(search("SO", APIdata))
	

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

	fmt.Println("(htpp://localhost:8877) - Server started on port ", appConfig.Port)
	http.ListenAndServe(appConfig.Port, nil)
}
