package entity

// Anime data
type Anime struct {
	ID      string
	Banner 	string
	Name 	string
	Note    string
	Genres  []string
	Year    string
}

// NewAnime create a new anime
func NewAnime(banner, name, note, year string, genres []string) (*Anime, error){
	a := &Anime{
		Banner: banner,
		Name: name,
		Note: note,
		Year: year,
		Genres: genres,
	}
	err := a.Validate()
	if err != nil{
		return nil, ErrInvalidEntity
	}
	return a, err
}

func (a *Anime) Validate() error{
	if a.Name == "" || a.Banner == ""{
		return ErrInvalidEntity
	}
	return nil
}