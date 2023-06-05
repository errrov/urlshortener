package shorten_test

import (
	"log"
	"testing"

	"github.com/errrov/urlshortenerozon/internal/model"
	"github.com/errrov/urlshortenerozon/internal/shorten"
	"github.com/errrov/urlshortenerozon/internal/storage/in_memory"
)

func TestShortenService(t *testing.T) {
	shortenService := shorten.NewService(in_memory.NewInMemory())
	input := model.UserInput{OriginalURL: "google.com"}
	shortening, err := shortenService.AddShortenLinkToStorage(input)
	if err != nil {
		t.Errorf("Test Shorten Service adding (%v) = %v", "google.com", shortening)
	}
	if shortening.Identifier == "" {
		t.Errorf("Identifier of string is empty %v", shortening.Identifier)
	}
	if shortening.Original != "google.com" {
		t.Errorf("Original URL is empty")
	}
}

// strange test ?
func TestAlreadyExistingID(t *testing.T) {
	shortenService := shorten.NewService(in_memory.NewInMemory())
	input := model.UserInput{OriginalURL: "https://www.google.com", Identifier: "abcde"}
	_, err := shortenService.AddShortenLinkToStorage(input)
	if err != nil {
		t.Errorf("Test Shorten Service adding (%v)", "google.com")
	}
	_, err = shortenService.AddShortenLinkToStorage(input)
	log.Print(err)
	if err != model.ErrIdExist {
		t.Errorf("Test adding already existing id (%v)", "google.com")
	}
}

func TestRedirectService(t *testing.T) {
	inMemoryStorage := in_memory.NewInMemory()
	shortenService := shorten.NewService(inMemoryStorage)
	identifier := "google"
	input := model.UserInput{OriginalURL: "https://www.google.com", Identifier: identifier}
	_, err := shortenService.AddShortenLinkToStorage(input)
	if err != nil {
		t.Errorf("Error adding link to storage")
	}
	redirectURL, err := shortenService.Redirect(identifier)
	if err != nil {
		t.Errorf("Error creating redirect Link")
	}

	if redirectURL != "https://www.google.com" {
		t.Errorf("Error creating redirect URL")
	}
}

func TestRedirectNotExist(t *testing.T) {
	shortenService := shorten.NewService(in_memory.NewInMemory())
	_, err := shortenService.Redirect("google")
	if err != model.ErrNotFound {
		t.Errorf("No error for missing id")
	}
}
