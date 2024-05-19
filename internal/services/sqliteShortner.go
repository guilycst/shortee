package services

import (
	"database/sql"

	"github.com/guilycst/shortee/internal/ports/services"
	"github.com/guilycst/shortee/pkg/b62"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteShortener struct {
	bigIntGenerator services.BigIntGenerator
	db              *sql.DB
}

func NewShortener(db *sql.DB, bigIntGenerator services.BigIntGenerator) (*SQLiteShortener, error) {
	r, err := db.Exec("CREATE TABLE IF NOT EXISTS shortened (id varchar(11), `value` text)")
	if err != nil {
		return nil, err
	}

	ra, err := r.RowsAffected()
	if err != nil {
		return nil, err
	}

	if ra > 0 {
		_, err = db.Exec("CREATE INDEX idx_shortened_id ON shortened (id)")
		if err != nil {
			return nil, err
		}
	}

	return &SQLiteShortener{
		bigIntGenerator: bigIntGenerator,
		db:              db,
	}, nil
}

func (s *SQLiteShortener) Shorten(url string) (string, error) {
	id, err := s.bigIntGenerator.Generate()
	if err != nil {
		return "", err
	}

	b62ID, err := b62.Encode(id)
	if err != nil {
		return "", err
	}

	_, err = s.db.Exec("INSERT INTO shortened (id, value) VALUES (?, ?)", b62ID, url)
	if err != nil {
		return "", err
	}

	return b62ID, nil
}

func (s *SQLiteShortener) Resolve(id string) (string, error) {
	var url string
	err := s.db.QueryRow("SELECT value FROM shortened WHERE id = ?", id).Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}
