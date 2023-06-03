package model

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type Shortened struct {
	Identifier string `json:"identifier"`
	Original string `json:"original_url"`
}