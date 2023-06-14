package in_memory

import (
	"sync"

	"github.com/errrov/urlshortener/internal/model"
)

type inMemory struct {
	m sync.Map
}

func NewInMemory() *inMemory {
	return &inMemory{}
}

func (s *inMemory) Add(shortened model.Shortened) (*model.Shortened, error) {
	if _, exist := s.m.Load(shortened.Identifier); exist {
		return nil, model.ErrIdExist
	}
	s.m.Store(shortened.Identifier, shortened)
	return &shortened, nil
}

func (s *inMemory) Get(identifier string) (*model.Shortened, error) {
	v, ok := s.m.Load(identifier)
	if !ok {
		return nil, model.ErrNotFound
	}
	shortening := v.(model.Shortened)
	return &shortening, nil
}
