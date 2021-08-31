package anime

import "github.com/GeovaneCavalcante/api-anime-golang/entity"

type inmem struct {
	m map[entity.IDHash]*entity.Anime
}

func newInmem() *inmem {
	var m = map[entity.IDHash]*entity.Anime{}
	return &inmem{
		m: m,
	}
}
