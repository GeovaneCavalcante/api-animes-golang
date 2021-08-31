package anime

import "github.com/GeovaneCavalcante/api-anime-golang/entity"

// Writer interface
type Writer interface {
	Get(id string) (*entity.Anime, error)
	Search(query string) ([]*entity.Anime, error)
	List() ([]*entity.Anime, error)
}

// Render interface
type Render interface {
	Create(e *entity.Anime) (*entity.Anime, error)
	Update(e *entity.Anime) error
	Delete(id *entity.ID) error
}

// Repository interface
type Repository interface {
	//Render
	//Writer
	Get(id string) (*entity.Anime, error)
}

// UseCase interfaces
type UseCase interface {
	GetAnime(id string) (*entity.Anime, error)
}
