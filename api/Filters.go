package api

type Filters struct {
	MinCreationDate int
	MaxCreationDate int

	MinFirstAlbumDate int
	MaxFirstAlbumDate int

	Err error

	Image string
	Name string

	Groupies []Groupie
	ArtistName []string
}

func (f *Filters) InitFilters() {
	f.MinCreationDate = 1950
	f.MaxCreationDate = 2024
	f.MinFirstAlbumDate = 1950
	f.MaxFirstAlbumDate = 2024
	f.Err = nil
	f.Image = ""
	f.Groupies = []Groupie{}
	f.ArtistName = []string{}
}
