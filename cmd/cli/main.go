package main

import (
	"math/big"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/guilycst/shortee/pkg/b62"
)

func main() {
	host := "localhost"
	var num atomic.Uint64
	num.Store(10000 + uint64(time.Now().UnixNano()))
	s, err := b62.Encode(big.NewInt(int64(num.Load())))
	if err != nil {
		panic(err)
	}

	path, err := url.JoinPath(host, s)
	if err != nil {
		panic(err)
	}

	println("https://" + path)
}
