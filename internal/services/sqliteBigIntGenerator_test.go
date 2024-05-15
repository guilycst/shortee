package services_test

import (
	"math/big"
	"sync"
	"testing"

	"github.com/guilycst/shortee/internal/services"
)

func TestNewSQLiteBigIntGeneratorShouldNotErrValidDBPath(t *testing.T) {
	s, err := services.NewSQLiteBigIntGenerator(":memory:")
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
	defer s.Close()
}

func TestGenerateIncrementsAtomicCounterSingleThread(t *testing.T) {
	s, _ := services.NewSQLiteBigIntGenerator(":memory:")
	defer s.Close()

	num, err := s.Generate()
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	if num.Int64() > 1 {
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

func TestGenerateIncrementsAtomicCounterConcurrent(t *testing.T) {
	s, _ := services.NewSQLiteBigIntGenerator("../../test.db")
	defer s.Close()

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
