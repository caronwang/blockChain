package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

/*
	Index 是这个块在整个链中的位置
	Timestamp 显而易见就是块生成时的时间戳
	Hash 是这个块通过 SHA256 算法生成的散列值
	PrevHash 代表前一个块的 SHA256 散列值
*/

type Block struct {
	Index     int64
	Timestamp int64
	Txs       []*Transaction
	Hash      string
	PrevHash  string
	Nonce     int64
}

/*
	创建区块
*/
func CreateBlock(index int64, prevHash string, txs []*Transaction) *Block {
	blc := Block{
		Index:     index,
		Timestamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Txs:       txs,
	}
	//blc.SetHash()

	//调用工作量证明，返回有效的hash和Nance值
	pow := NewPOW(&blc)

	hash, nc := pow.Run()
	blc.Hash = hex.EncodeToString(hash)
	blc.Nonce = nc

	return &blc
}

func (blc *Block) SetHash() *Block {

	txsString, err := json.Marshal(blc.Txs)
	if err != nil {
		log.Fatal(err)
	}

	d := strconv.FormatInt(blc.Index, 10) + strconv.FormatInt(blc.Timestamp, 10) + blc.PrevHash + string(txsString)
	hash := sha256.Sum256([]byte(d))
	blc.Hash = hex.EncodeToString(hash[:])

	return blc
}

/*
	将区块序列化为字节数组
*/
func (blc *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(blc)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

/*
	反序列化字节数组为区块
*/
func DeserializeBlock(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

/*
	挖掘新的区块
*/
func MineNewBlock(from []string, to []string, amount []string) {
	fmt.Println("from:", from)
	fmt.Println("to:", to)
	fmt.Println("amount:", amount)

	//1.通过相关算法建立Transaction数组
	//var txs []*Transaction

}
