package b62_test

import (
	"math/big"
	"testing"

	"github.com/guilycst/shortee/pkg/b62"
)

func TestEncodeNil(t *testing.T) {
	_, err := b62.Encode(nil)
	if err == nil {
		t.Error("Expected error")
	}
	if err.Error() != "nil input" {
		t.Errorf("Expected 'nil input', got %s", err.Error())
	}
}

func TestEncodeEmpty(t *testing.T) {
	e := b62.EncodeAny([]byte{})
	if e != "" {
		t.Errorf("Expected '', got %s", e)
	}
}

func TestEncodeDecode(t *testing.T) {
	data := []byte("hello")
	e := b62.EncodeAny(data)
	d := b62.Decode(e)
	if string(d) != "hello" {
		t.Errorf("Expected %v, got %v", data, d)
	}
}

func TestEncodeDecodeLong(t *testing.T) {
	data := []byte("hello world")
	e := b62.EncodeAny(data)
	d := b62.Decode(e)
	if string(d) != "hello world" {
		t.Errorf("Expected %v, got %v", data, d)
	}
}

func TestEncodeDecodeVeryLong(t *testing.T) {
	data := "some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode, some very long string that is not easy to encode."
	e := b62.EncodeString(data)
	d := b62.Decode(e)
	if string(d) != string(data) {
		t.Errorf("Expected %v, got %v", data, d)
	}
}

func TestEncodeDecodeZero(t *testing.T) {
	e, err := b62.Encode(big.NewInt(0))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	d := b62.Decode(e)
	if string(d) != string(big.NewInt(0).Bytes()) {
		t.Errorf("Expected '0', got %v", d)
	}
}

func TestEncodeDecodeZeroString(t *testing.T) {
	e := b62.EncodeAny([]byte("0"))
	d := b62.Decode(e)
	if string(d) != "0" {
		t.Errorf("Expected '0', got %v", d)
	}
}

func TestEncodeDecodeOne(t *testing.T) {
	e, err := b62.Encode(big.NewInt(1))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	d := b62.Decode(e)
	if string(d) != string(big.NewInt(1).Bytes()) {
		t.Errorf("Expected '1', got %v", d)
	}
}

func TestEncodeDecodeOneString(t *testing.T) {
	e := b62.EncodeAny([]byte("1"))
	d := b62.Decode(e)
	if string(d) != "1" {
		t.Errorf("Expected '1', got %v", d)
	}
}
