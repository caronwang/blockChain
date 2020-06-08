package BLC

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Wallets struct {
	WMap map[string]*Wallet
}

const walletFile = "wallet.dat"

/*
	从本地文件读取wallets信息
*/
func NewWallets() *Wallets {
	wallets := &Wallets{WMap: map[string]*Wallet{}}

	if err := wallets.LoadFromFile(); err != nil {
		fmt.Println(err)
	}
	return wallets
}

func (ws *Wallets) CreateWallet() string {
	w := NewWallet()
	fmt.Printf("Address:%v\n", w.GetAddress())
	ws.WMap[w.GetAddress()] = w
	ws.SaveToFile()
	return w.GetAddress()
}

/*
	获取钱包里所有地址
*/
func (ws *Wallets) GetAddresses() []string {
	var addresses []string

	for address := range ws.WMap {
		addresses = append(addresses, address)
	}
	return addresses

}

/*
	根据地址获取钱包对象
*/
func (ws *Wallets) GetWallet(address string) Wallet {
	return *ws.WMap[address]
}

/*
	加载钱包文件
*/
func (ws *Wallets) LoadFromFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Fatal(err)
	}

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&ws)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

/*
	保存钱包数据到本地文件
*/
func (ws *Wallets) SaveToFile() {
	var content bytes.Buffer

	//注册的目的为了序列化任何类型
	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)
	if err != nil {
		log.Fatal(err)
	}
	//序列化以后的数据写入文件，原来的文件会被覆盖
	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}

}

func (ws *Wallets) Serialize() []byte {
	data, err := json.Marshal(ws)
	if err != nil {
		log.Panic(err)
	}
	return data
}
