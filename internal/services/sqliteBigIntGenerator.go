package services

import (
	"database/sql"
	"math/big"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteBigIntGenerator struct {
	db *sql.DB
}

func NewSQLiteBigIntGenerator(dbPath string) (*SQLiteBigIntGenerator, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	r, err := db.Exec("CREATE TABLE IF NOT EXISTS counter (id varchar(255) PRIMARY KEY, `value` BIGINT DEFAULT 1)")
	if err != nil {
		return nil, err
	}

	ra, err := r.RowsAffected()
	if err != nil {
		return nil, err
	}

	if ra > 0 {
		_, err = db.Exec("INSERT INTO counter (id, value) VALUES ('default', 1)")
		if err != nil {
			return nil, err
		}
	}

	return &SQLiteBigIntGenerator{
		db: db,
	}, nil
}

func (s *SQLiteBigIntGenerator) Generate() (*big.Int, error) {

	rows, err := s.db.Query("UPDATE counter set value = value + 1 where id = 'default' RETURNING value;")
	if err != nil {
		return nil, err
	}

	var value big.Int
	rows.Scan(&value)

	defer rows.Close()

	return &value, nil
}

func (s *SQLiteBigIntGenerator) Close() error {
	return s.db.Close()
}
