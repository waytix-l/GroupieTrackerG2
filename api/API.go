package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
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

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr.Error())
	}

	var groupes []Groupie
	json.Unmarshal(body, &groupes)

	return groupes
}

func search(word string, api []Groupie) []string {
	regex := regexp.MustCompile(word)
	tab := []string{}
	for index, groupe := range api {
		for _, match := range regex.FindAllString(groupe.Name, -1) {
			a := match + " found at index " + strconv.Itoa(index)
			tab = append(tab, a)
		}
	}
	return tab
}

