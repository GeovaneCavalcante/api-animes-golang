package entity_test

import (
	"testing"

	"github.com/GeovaneCavalcante/api-anime-golang/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewAnime(t *testing.T) {
	genres := []string{"ficao", "aventura"}
	a, err := entity.NewAnime(
		"https://animesonline.cc/wp-content/uploads/2019/06/l9G0xQ5Wc3EvLslzDDrNXlfdB8m-185x278.jpg",
		"Sakura Trick",
		"",
		"2020",
		genres,
	)
	assert.Nil(t, err)
	assert.Equal(t, a.Name, "Sakura Trick")
}

func TestAnimeValidate(t *testing.T) {
	genres := []string{"ficao", "aventura"}
	type test struct {
		banner string
		name   string
		note   string
		year   string
		genres []string
		want   error
	}
	tests := []test{
		{
			banner: "https://animesonline.cc/wp-content/uploads/2019/06/l9G0xQ5Wc3EvLslzDDrNXlfdB8m-185x278.jpg",
			name:   "Sakura Trick",
			note:   "",
			year:   "2020",
			genres: genres,
			want:   nil,
		},
		{
			banner: "",
			name:   "Sakura Trick",
			note:   "",
			year:   "2020",
			genres: genres,
			want:   entity.ErrInvalidEntity,
		},
		{
			banner: "https://animesonline.cc/wp-content/uploads/2019/06/l9G0xQ5Wc3EvLslzDDrNXlfdB8m-185x278.jpg",
			name:   "",
			note:   "",
			year:   "2020",
			genres: genres,
			want:   entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewAnime(tc.banner, tc.name, tc.note, tc.year, tc.genres)
		assert.Equal(t, err, tc.want)
	}
}
