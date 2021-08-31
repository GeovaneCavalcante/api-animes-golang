package presenter

// Anime data
type Anime struct {
	ID     string    `json:"id"`
	Banner string    `json:"banner"`
	Name   string    `json:"name"`
	Note   string    `json:"note"`
	Genres []string  `json:"genres"`
	Year   string    `json:"year"`
}
