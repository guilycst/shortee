package b62

import (
	"bytes"
	"errors"
	"math/big"
	"strings"
)

const base62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func EncodeString(data string) string {
	num := big.NewInt(0).SetBytes([]byte(data))
	e, _ := Encode(num)
	return e
}

func EncodeAny(data []byte) string {
	num := big.NewInt(0).SetBytes(data)
	e, _ := Encode(num)
	return e
}

func Encode(num *big.Int) (string, error) {
	if num == nil {
		return "", errors.New("nil input")
	}

	base := big.NewInt(62)
	result := make([]byte, 0, 11) // Initial capacity assuming average output length of 11 characters

	zero := big.NewInt(0)
	mod := &big.Int{}
	quotient := new(big.Int).Set(num)

	var buffer bytes.Buffer // Reuse buffer for string concatenation

	for quotient.Cmp(zero) > 0 {
		quotient.DivMod(quotient, base, mod)
		buffer.WriteByte(base62Alphabet[mod.Int64()])
	}

	// Reverse the buffer contents to get the final result
	for i := buffer.Len() - 1; i >= 0; i-- {
		result = append(result, buffer.Bytes()[i])
	}

	return string(result), nil
}

func Decode(encoded string) []byte {
	base := big.NewInt(62)
	result := big.NewInt(0)
	power := big.NewInt(1)

	for i := len(encoded) - 1; i >= 0; i-- {
		charIndex := strings.Index(base62Alphabet, string(encoded[i]))
		charValue := big.NewInt(int64(charIndex))
		temp := big.NewInt(0).Mul(charValue, power)
		result = result.Add(result, temp)
		power = big.NewInt(0).Mul(power, base)
	}

	return result.Bytes()
}
