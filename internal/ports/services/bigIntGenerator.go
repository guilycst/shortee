package services

import "math/big"

type BigIntGenerator interface {
	Generate() (*big.Int, error)
}
