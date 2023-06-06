package shorten

import (
	"log"

	"github.com/errrov/urlshortenerozon/internal/model"
	"github.com/google/uuid"
)

type Storage interface {
	Add(shortened model.Shortened) (*model.Shortened, error)
	Get(identifier string) (*model.Shortened, error)
}

type Service struct {
	Storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{Storage: storage}
}

func (s *Service) AddShortenLinkToStorage(input model.UserInput) (*model.Shortened, error) {
	id := uuid.New().ID()
	var linkIndetifier string
	if input.Identifier == "" {
		linkIndetifier = Shorten(id)
	} else {
		linkIndetifier = input.Identifier
	}
	newUrl := model.Shortened{
		Identifier: linkIndetifier,
		Original:   input.OriginalURL,
	}
	log.Println("INSIDE ADD SHORTEN FUNCTION")
	shortened, err := s.Storage.Add(newUrl)
	if err != nil {
		return nil, err
	}
	return shortened, nil
}

func (s *Service) Get(identifier string) (*model.Shortened, error) {
	shortened, err := s.Storage.Get(identifier)
	if err != nil {
		return nil, err
	}
	return shortened, nil
}

func (s *Service) Redirect(identifier string) (string, error) {
	shortened, err := s.Storage.Get(identifier)
	if err != nil {
		return "", err
	}
	return shortened.Original, nil
}
