package api

type Filters struct {
	MinCreationDate int
	MaxCreationDate int

	MinFirstAlbumDate int
	MaxFirstAlbumDate int

	MembersAmount []int

	Err error

	Image string
	Name string
}

func (f *Filters) InitFilters() {
	f.MinCreationDate = 1950
	f.MaxCreationDate = 2024
	f.MinFirstAlbumDate = 1950
	f.MaxFirstAlbumDate = 2024
	f.MembersAmount = []int{}
	f.Err = nil
	f.Image = ""
}
