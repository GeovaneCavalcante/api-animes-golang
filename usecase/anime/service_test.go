package anime

import "github.com/GeovaneCavalcante/api-anime-golang/entity"

func newFixtureAnime() *entity.Anime {
	return &entity.Anime{
		Banner: "https://animesonline.cc/wp-content/uploads/2019/06/l9G0xQ5Wc3EvLslzDDrNXlfdB8m-185x278.jpg",
		Name:   "Sakura Trick",
		Note:   "6",
		Year:   "2020",
		Genres: []string{"Ficcao", "Acao"},
	}
}
