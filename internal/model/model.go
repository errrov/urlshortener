package model

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrIdExist  = errors.New("identifier already exists")
)

type Shortened struct {
	Identifier string `json:"identifier"`
	Original   string `json:"original_url"`
}

type UserInput struct {
	OriginalURL string
	Identifier  string
}
