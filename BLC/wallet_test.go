package BLC

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestNewWallet(t *testing.T) {
	// bytes := []byte("haha wang asdasdasdas 12312312")
	// hasher := sha256.New()
	// hasher.Write(bytes)
	// hash := hasher.Sum(nil)
	// fmt.Println(len(hex.EncodeToString(hash)))
	// byte58 := Base58Encode(bytes)
	// fmt.Println(len(hex.EncodeToString(byte58)))
	w := NewWallet()
	addr := w.GetAddress()
	fmt.Println(addr)
	fmt.Println(len(addr))
	fmt.Println(ValidateAddress(addr))
}

func TestNewWallets(t *testing.T) {
	// bytes := []byte("haha wang asdasdasdas 12312312")
	// hasher := sha256.New()
	// hasher.Write(bytes)
	// hash := hasher.Sum(nil)
	// fmt.Println(len(hex.EncodeToString(hash)))
	// byte58 := Base58Encode(bytes)
	// fmt.Println(len(hex.EncodeToString(byte58)))
	ws := NewWallets()
	ws.CreateWallet()
	spew.Dump(ws)
}
