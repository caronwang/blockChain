package BLC

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"
)

/*
   hash有256位
*/

const tartgetBit = 16

type POW struct {
	block  *Block   //当前需要验证的区块
	target *big.Int //大数据存储，代表挖矿难度
}

func NewPOW(blk *Block) *POW {

	//创建big.Int对象
	target := big.NewInt(1)
	target.Lsh(target, 256-tartgetBit)

	return &POW{block: blk, target: target}
}

func (pow *POW) PrepareData(nonce int64) []byte {

	txsString, err := json.Marshal(pow.block.Txs)
	if err != nil {
		log.Fatal(err)
	}

	data := strconv.FormatInt(pow.block.Index, 10) +
		strconv.FormatInt(pow.block.Timestamp, 10) +
		strconv.FormatInt(nonce, 10) + pow.block.PrevHash +
		string(txsString)

	return []byte(data)

}

/*

 */
func (pow *POW) IsValid() bool {
	var hashInt big.Int
	var hash [32]byte
	data := pow.PrepareData(pow.block.Nonce)

	hash = sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	if pow.target.Cmp(&hashInt) == 1 {
		return true
	}

	return false
}

func (pow *POW) Run() ([]byte, int64) {
	// 将block属性拼接为字节数组

	// 生成hash

	// 判断hash有效性，满足条件，跳出循环
	var nonce int64 = 0
	var hashInt big.Int
	var hash [32]byte
	t := time.Now()
	for {
		data := pow.PrepareData(nonce)

		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		fmt.Printf("\r")
		fmt.Printf("try:%v hash:%x", nonce, hash[:])
		if pow.target.Cmp(&hashInt) == 1 {
			fmt.Printf("\n挖矿完成!耗时%s\n", time.Since(t).String())
			break
		}

		nonce++
	}
	return hash[:], nonce
}
