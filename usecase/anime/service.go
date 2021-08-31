package anime

import (
	"github.com/GeovaneCavalcante/api-anime-golang/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetAnime(id string) (*entity.Anime, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return b, err
}
