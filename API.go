package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Groupie struct {
	ID int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
}

func importAPI() []Groupie {
	API_url := "https://groupietrackers.herokuapp.com/api/artists"
	req, err := http.NewRequest("GET", API_url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	// defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr.Error())
	}

	var groupes []Groupie
	json.Unmarshal(body, &groupes)

	//formattedData := formatJSON(body)

	return groupes
}



