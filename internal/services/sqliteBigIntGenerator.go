package services

import (
	"database/sql"
	"math/big"

	"github.com/guilycst/shortee/internal/ports/services"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteBigIntGenerator struct {
	db *sql.DB
}

func NewSQLiteBigIntGenerator(db *sql.DB) (*SQLiteBigIntGenerator, error) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS counter (id varchar(255) PRIMARY KEY, `value` BIGINT DEFAULT 1);")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("INSERT INTO counter (id, value) VALUES ('default', 1) on conflict(id) do nothing;")
	if err != nil {
		return nil, err
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

	ok := rows.Next()
	if !ok {
		err := rows.Err()
		if err != nil {
			return nil, err
		}
		return nil, &services.ErrNoRows{}
	}
	var value int64
	err = rows.Scan(&value)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return big.NewInt(value), nil
}
