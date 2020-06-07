package BLC

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestMakeRipemd160(t *testing.T) {
	data := []byte("hello")
	fmt.Println(MakeRipemd160(data))
	fmt.Println(len(MakeRipemd160(data)))
}

func TestBase58Encode(t *testing.T) {
	binEncode := Base58Encode([]byte("hello world"))
	fmt.Println(hex.EncodeToString(binEncode))
	binDecode := Base58Decode(binEncode)
	fmt.Println(string(binDecode))
}
