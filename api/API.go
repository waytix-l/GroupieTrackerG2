package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Groupie struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	ConcertLoc   ConcertLocation
}


type ConcertLocation struct {
	Loc []string `json:locations`
}

func ImportAPI() []Groupie {
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

func ImportConcert(api []Groupie) {
	ConcertURL := "https://groupietrackers.herokuapp.com/api/locations"
	req, err := http.NewRequest("GET", ConcertURL, nil)
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

	var concert []ConcertLocation
	json.Unmarshal(body, &concert)

	// for _, groupie := range api {
	// 	groupie.ConcertLoc = concert
	// }
}

func Search(word string, api []Groupie) []Groupie {
	regex := regexp.MustCompile(word)
	tab := []Groupie{}
	for _, groupe := range api {
		for i := 0; i < len(regex.FindAllString(groupe.Name, -1)); i++ {
			tab = append(tab, groupe)
		}
	}
	return tab
}
