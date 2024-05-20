package services_test

import (
	"testing"

	pservices "github.com/guilycst/shortee/internal/ports/services"
	"github.com/guilycst/shortee/internal/services"
)

func getShortener() *services.SQLiteShortener {
	db := getDB()

	g, err := services.NewSQLiteBigIntGenerator(db)
	if err != nil {
		panic(err)
	}
	short, err := services.NewShortener(db, g)
	if err != nil {
		panic(err)
	}
	return short
}

func TestShortenShouldNotErr(t *testing.T) {
	short := getShortener()
	_, err := short.Shorten("https://google.com")
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestShortenShouldErrEmptyURL(t *testing.T) {
	short := getShortener()
	_, err := short.Shorten("")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if _, ok := err.(*pservices.ErrEmptyURL); !ok {
		t.Fatalf("expected ErrEmptyURL, got %v", err)
	}
}

func TestShortenShouldErrURLTooLong(t *testing.T) {
	short := getShortener()
	s, err := short.Shorten("https://google.com/?" + string(make([]byte, 2083)))
	if err == nil {
		t.Fatalf("expected err, got %v", s)
	}

	if _, ok := err.(*pservices.ErrURLTooLong); !ok {
		t.Fatalf("expected ErrURLTooLong, got %v", err)
	}
}
