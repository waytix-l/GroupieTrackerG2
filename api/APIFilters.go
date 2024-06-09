package api

import "strconv"

func CreationDate(date1 int, date2 int, api []Groupie) []Groupie {
	tab := []Groupie{}
	for _, groupie := range api {
		if date1 < date2 {
			if (groupie.CreationDate >= date1 && groupie.CreationDate <= date2) {
				tab = append(tab, groupie)
			}
		} else if date1 > date2 {
			if (groupie.CreationDate <= date1 && groupie.CreationDate >= date2) {
				tab = append(tab, groupie)
			}
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