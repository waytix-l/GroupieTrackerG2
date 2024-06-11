package main

import (
	"GroupieTracker/api"
	"GroupieTracker/config"
	"fmt"
	"net/http"
)

func main() {
	var appConfig config.Config

	APIdata := api.ImportAPI()
	ConcertData := api.ImportConcert()

	// for _, v := range APIdata {
	// 	fmt.Println(v)
	// }

	fmt.Println(APIdata)
	fmt.Println()
	fmt.Println(ConcertData)

	templateCache, err := CreateTemplateCache()

	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8080"

	CreateTemplates(&appConfig)

	http.HandleFunc("/", Home)
	http.HandleFunc("/artist", Artist)
	http.HandleFunc("/concert", Concert)
	http.HandleFunc("/date", Date)

	fmt.Println("(http://localhost:8080) - Server started on port ", appConfig.Port)
	http.ListenAndServe(appConfig.Port, nil)
}
