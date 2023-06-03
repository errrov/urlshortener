package shorten_test

import (
	"testing"

	"github.com/errrov/linkshortener/internal/shorten"
)

func TestShortenValid(t *testing.T) {
	var tests = []struct {
		input uint32
		want  string
	}{
		{0, ""},
		{1024, "jj"},
	}
	for _, test := range tests {
		if got := shorten.Shorten(test.input); got != test.want {
			t.Errorf("TestShortenValid(%v) = %v", test.input, got)
		}
	}
}

func TestShortenIdempotent(t *testing.T) {
	for i := 0; i < 100; i++ {
		if got := shorten.Shorten(1024); got != "jj" {
			t.Errorf("TestShortenValid(%v) = %v", 1024, got)
		}
	}
}

