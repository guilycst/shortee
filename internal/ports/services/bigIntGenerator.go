package services

import "math/big"

type BigIntGenerator interface {
	Closeable
	Generate() (*big.Int, error)
}
