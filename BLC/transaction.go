package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

/*
	Transaction 创建分为两种情况
	1.创世区块创建时候的transaction
	2.转账时产生的tranaction
*/

type Transaction struct {
	//交易hash
	Txhash []byte
	//输入
	Vins []*TXInput
	//输出
	Vouts []*TXOutput
}

func NewCoinbaseTransaction(address string) *Transaction {
	//代表消费
	txInput := &TXInput{TxHash: []byte{}, Vout: -1, ScriptSign: "genesys data"}

	//未消费
	txOuput := &TXOutput{Value: 10, ScriptPubKey: address}

	txCoinbase := &Transaction{Txhash: []byte{}, Vins: []*TXInput{txInput}, Vouts: []*TXOutput{txOuput}}
	txCoinbase.HashTransaction()
	return txCoinbase
}

/*
	设置hash值
*/
func (tx *Transaction) HashTransaction() {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())
	tx.Txhash = hash[:]
	return
}

func Send(from []string, to []string, amount []string) {
	MineNewBlock(from, to, amount)
}
