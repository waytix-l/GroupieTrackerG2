package main

import (
	"GroupieTracker/api"
	"GroupieTracker/config"
	"GroupieTracker/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	var appConfig config.Config

	APIdata := api.ImportAPI()

	fmt.Println(APIdata)
	fmt.Println()
	fmt.Println(APIdata[0])
	fmt.Println()
	fmt.Println(api.Search("S", APIdata))

	templateCache, err := handlers.CreateTemplateCache()

	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8080"

	handlers.CreateTemplates(&appConfig)

	http.HandleFunc("/", handlers.Home)

	fmt.Println("(http://localhost:8080) - Server started on port ", appConfig.Port)
	http.ListenAndServe(appConfig.Port, nil)
}
