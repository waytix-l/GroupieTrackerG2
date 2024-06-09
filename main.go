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

	fmt.Println(APIdata)
	fmt.Println()
	fmt.Println(APIdata[0])
	fmt.Println()
	//fmt.Println(api.Search("S", APIdata))
	//fmt.Println()
	//fmt.Println(APIdata[0].Image)
	//fmt.Println()
	fmt.Println(api.CreationDate(1970, 1990, APIdata))
	fmt.Println()
	fmt.Println(api.FirstAlbum(1970, APIdata))
	fmt.Println()
	fmt.Println(api.MembersAmount(7, APIdata))

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

	fmt.Println("(http://localhost:8080) - Server started on port ", appConfig.Port)
	http.ListenAndServe(appConfig.Port, nil)
}
