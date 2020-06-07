package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"

	"golang.org/x/crypto/ripemd160"
)

const (
	version            = byte(0x00)
	addressChecksumLen = 4
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewWallet() *Wallet {
	pri, pub := newKeyPair()
	wallet := &Wallet{PrivateKey: pri, PublicKey: pub}
	return wallet
}

/*
	公钥hash (20字节)
*/
func HashPubKey(pubKey []byte) []byte {
	publicSha256 := sha256.Sum256(pubKey)

	Ripemd160Hasher := ripemd160.New()
	_, err := Ripemd160Hasher.Write(publicSha256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRipemd160 := Ripemd160Hasher.Sum(nil)

	return publicRipemd160
}

/*
	返回钱包地址
*/
func (w *Wallet) GetAddress() string {
	pubkeyhash := HashPubKey(w.PublicKey)

	versionPayload := append([]byte{version}, pubkeyhash...) //21字节
	checkSum := checksum(versionPayload)                     //4字节
	//fmt.Printf("versionPayload:%v\n", len(versionPayload))
	//fmt.Printf("checksum:%v\n", len(checkSum))
	fullPayload := append(versionPayload, checkSum...) //29字节
	//fmt.Printf("fullPayload:%v\n", len(fullPayload))
	address := Base58Encode(fullPayload)
	//fmt.Printf("address:%v\n", len(address))
	return hex.EncodeToString(address)
}

/*
	为一个公钥生成一个checksum （4字节）
*/
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}

/*
	验证地址
*/
func ValidateAddress(address string) bool {
	addressBytes, err := hex.DecodeString(address)
	if err != nil {
		log.Panic(err)
	}
	pubKeyHash := Base58Decode(addressBytes)
	acutalChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	//fmt.Printf("%v\n%v\n", hex.EncodeToString(pubKeyHash), hex.EncodeToString(acutalChecksum))
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))
	//fmt.Printf("%v\n", hex.EncodeToString(targetChecksum))
	return bytes.Compare(acutalChecksum, targetChecksum) == 0
}

/*
	创建私钥、公钥
*/
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}
