package services

import (
	"database/sql"
	"math/big"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteBigIntGenerator struct {
	db *sql.DB
}

func NewSQLiteBigIntGenerator(db *sql.DB) (*SQLiteBigIntGenerator, error) {
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

	rows.Next()
	var value int64
	err = rows.Scan(&value)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return big.NewInt(value), nil
}
