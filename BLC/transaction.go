package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
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

func (tx *Transaction) IsCoinbaseTransaction() bool {
	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1
}

/*
	创建创世交易
*/
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
	设置Transaction的hash值
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

/*

 */
func GetTXInputs(txs []*Transaction, addr string, amount int) ([]*TXInput, int) {
	var result []*TXInput
	balance := 0
	for i, tx := range txs {
		if balance >= amount {
			break
		}

		for index, vout := range tx.Vouts {
			if vout.ScriptPubKey == addr {
				balance += vout.Value
				result = append(result, &TXInput{TxHash: txs[i].Txhash, Vout: index, ScriptSign: addr})
			}
		}
	}

	return result, balance
}

/*
	创建交易

*/
func NewTransaction(from string, to string, amount int) *Transaction {
	log.Printf("[transaction] %s -> %s : %v\n", from, to, amount)

	//检查from address的余额
	balance, txs := GetBalanceByAddress(from)
	if balance < amount {
		log.Printf("[%v]余额不足,当前余额%v", from, balance)
		return nil
	}

	inputs, tBalance := GetTXInputs(txs, from, amount)
	// var inputs []*TXInput
	// txInput := &TXInput{
	// 	TxHash:     []byte{},
	// 	Vout:       -1,
	// 	ScriptSign: from,
	// }
	// inputs = append(inputs, txInput)

	var outputs []*TXOutput
	txOuput_from := &TXOutput{
		Value:        tBalance - amount,
		ScriptPubKey: from,
	}
	txOuput_to := &TXOutput{
		Value:        amount,
		ScriptPubKey: to,
	}
	outputs = append(outputs, txOuput_from, txOuput_to)

	tx := &Transaction{
		Txhash: []byte{},
		Vins:   inputs,
		Vouts:  outputs,
	}

	tx.HashTransaction()
	//log.Println(tx)
	return tx
}

/*
	查看TXOutput是否已经被消费
*/
func (tx *Transaction) IsSpend(addr string, spendTXOutputs map[string][]int) bool {
	for index, out := range tx.Vouts {
		if out.UnlockWithAddress(addr) {
			if spendTXOutputs != nil {
				for txHash, indexArray := range spendTXOutputs {
					if txHash == hex.EncodeToString(tx.Txhash) {
						for _, i := range indexArray {
							if index == i {
								return true
							}
						}
					}
				}
			}
		}
	}

	return false
}

/*
	通过Address获取账户余额相关交易
*/
func GetValidTxInputsByAddress(addr string) ([]*Transaction, error) {
	var txs []*Transaction
	spendTXOutputs := make(map[string][]int)

	for it := GetChain().Iterator(); it.HasNext(); it.Next() {
		//log.Println(it.Value.Index)
		block := it.Value

		for _, tx := range block.Txs {
			/*
				判断是否为创世区块
			*/
			if !tx.IsCoinbaseTransaction() {
				for _, in := range tx.Vins {
					if in.UnlockWithAddress(addr) {
						key := hex.EncodeToString(in.TxHash)
						spendTXOutputs[key] = append(spendTXOutputs[key], in.Vout)
					}
				}
			}

			if !tx.IsSpend(addr, spendTXOutputs) {
				txs = append(txs, tx)
			}
		}
	}
	txs = ReverseTransArray(txs)
	//spew.Dump(txs)
	return txs, nil
}

/*
	通过Address获取账户余额
*/
func GetBalanceByAddress(addr string) (int, []*Transaction) {
	var balance int

	//log.Println("地址:", addr)

	txs, err := GetValidTxInputsByAddress(addr)
	if err != nil {
		log.Panic(err)
	}

	for _, tx := range txs {
		for _, out := range tx.Vouts {
			if out.ScriptPubKey == addr {
				//spew.Dump(tx)
				log.Println("[rebalance]+", out.Value)
				balance += out.Value
			}
		}
	}

	return balance, txs
}
