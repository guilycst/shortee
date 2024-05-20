package services_test

import (
	"database/sql"
	"math/big"
	"sync"
	"testing"

	"github.com/guilycst/shortee/internal/services"
)

func getDB() *sql.DB {
	var err error
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)
	// db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	return db
}

func TestNewSQLiteBigIntGeneratorShouldNotErrValidDBPath(t *testing.T) {
	s, err := services.NewSQLiteBigIntGenerator(getDB())
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	if s == nil {
		t.Fatalf("expected not nil, got nil")
	}
}

func TestGenerateIncrementsAtomicCounterSingleThread(t *testing.T) {
	s, err := services.NewSQLiteBigIntGenerator(getDB())
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	num, err := s.Generate()
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	if num.Int64() < 1 {
		t.Fatalf("expected > 1, got %d", num.Int64())
	}

	num2, err := s.Generate()
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	if num2.Int64() < num.Int64() {
		t.Fatalf("expected > %d, got %d", num.Int64(), num2.Int64())
	}

}

func IgnoreGenerateIncrementsAtomicCounterConcurrent(t *testing.T) {
	s, _ := services.NewSQLiteBigIntGenerator(getDB())

	prev, _ := s.Generate()
	gens := make(chan *big.Int, 1000)

	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			num, err := s.Generate()
			if err != nil {
				panic(err)
			}
			gens <- num
			wg.Done()
		}()
	}

	wg.Wait()
	close(gens)

	for current := range gens {
		if current.Int64() > prev.Int64() {
			t.Fatalf("expected prev(%d) < current(%d)", prev.Int64(), current.Int64())
		}
		prev = current
	}
}
