package api

import "strconv"

func CreationDate(date int, api []Groupie) []Groupie {
	tab := []Groupie{}
	for _, groupie := range api {
		if (groupie.CreationDate == date) {
			tab = append(tab, groupie)
		}
	}
	return tab;
}

func FirstAlbum(date int, api []Groupie) []Groupie {
	tab := []Groupie{}
	for _, groupie := range api {
		album, _ := strconv.Atoi(groupie.FirstAlbum[6:])
		if (album == date) {
			tab = append(tab, groupie)
		}
	}
	return tab;
}

func MembersAmount(number int, api []Groupie) []Groupie {
	tab := []Groupie{}
	for _, groupie := range api {
		if (len(groupie.Members) == number) {
			tab = append(tab, groupie)
		}
	}
	return tab
}

func Location() {
	
}