package services

import "math/big"

type ErrNoRows struct {
}

func (e *ErrNoRows) Error() string {
	return "no rows"
}

type BigIntGenerator interface {
	Generate() (*big.Int, error)
}
