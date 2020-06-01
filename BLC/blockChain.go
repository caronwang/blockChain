package BLC

import (
	"blockChain/database"
	"errors"
	"fmt"
	"log"
	"strconv"
)

var (
	tableName = "BlockChainBucket"
	dbName    = "blkchain.db"
)

type BlockChain struct {
	LastBlock *Block
	Db        database.BlockChainDB
}

type Iterator struct {
	Value *Block
	Db    database.BlockChainDB
}

func (bc *BlockChain) Iterator() *Iterator {
	return &Iterator{
		Value: bc.LastBlock,
		Db:    bc.Db,
	}
}

func (i *Iterator) HasNext() bool {

	if i.Value != nil {
		//log.Println("### HasNext ###", i.Value.Hash)
		return true
	}

	return false
}

func (i *Iterator) Next() *Block {
	data := i.Db.Read([]byte(i.Value.PrevHash))
	if data != nil {
		i.Value = DeserializeBlock(data)
		return i.Value
	}
	i.Value = nil
	return nil
}

var blockChain *BlockChain

func init() {

	var lastBlock *Block
	db := database.NewBlockChainDB(dbName, tableName)
	// if db.Read([]byte("l")) == nil {
	// 	log.Println("创建创世区块")

	// 	txs := NewCoinbaseTransaction("user")
	// 	blc := CreateBlock(1, "", []*Transaction{txs})

	// 	err := db.Write([]byte("l"), []byte(blc.Hash))
	// 	if err != nil {
	// 		log.Panic("last hash数据存储失败！")
	// 	}

	// 	err = db.Write([]byte(blc.Hash), blc.Serialize())
	// 	if err != nil {
	// 		log.Panic("区块数据存储失败！")
	// 	}

	// 	lastBlock = blc
	// } else {
	// 	lb := db.Read([]byte(db.Read([]byte("l"))))
	// 	if lb != nil {
	// 		tbc := DeserializeBlock(lb)
	// 		if tbc != nil {
	// 			lastBlock = tbc
	// 		}
	// 	}
	// }
	lHash := db.Read([]byte("l"))
	if lHash != nil {
		lb := db.Read([]byte(lHash))
		if lb != nil {
			tbc := DeserializeBlock(lb)
			if tbc != nil {
				lastBlock = tbc
			}
		}
	}

	blockChain = &BlockChain{
		LastBlock: lastBlock,
		Db:        *db,
	}

	log.Println("区块链初始化完毕!")
}

func NewGenesysBlock(usr string) *Block {
	log.Println("创建创世区块")

	txs := NewCoinbaseTransaction(usr)
	blc := CreateBlock(1, "", []*Transaction{txs})

	err := blockChain.Db.Write([]byte("l"), []byte(blc.Hash))
	if err != nil {
		log.Panic("last hash数据存储失败！")
	}

	err = blockChain.Db.Write([]byte(blc.Hash), blc.Serialize())
	if err != nil {
		log.Panic("区块数据存储失败！")
	}

	blockChain.LastBlock = blc

	return blc
}

func GetBlockList() []*Block {
	var bc []*Block
	//fmt.Println(string(blockChain.Db.Read([]byte("l"))))

	for it := blockChain.Iterator(); it.HasNext(); it.Next() {
		//fmt.Println(it.Value)
		bc = append(bc, it.Value)
	}
	fmt.Println("block chain 长度:", len(bc))
	return bc
}

func GetChain() *BlockChain {
	return blockChain
}

/*
	挖掘新的区块
*/
func MineNewBlock(from []string, to []string, amount []string) error {
	fmt.Println("from:", from)
	fmt.Println("to:", to)
	fmt.Println("amount:", amount)

	//1.通过相关算法建立Transaction数组
	var txs []*Transaction

	for i, _ := range from {
		if i <= len(to) && i < len(amount) {
			t_amount, err := strconv.Atoi(amount[i])
			if err != nil {
				log.Panic(err)
			}
			ntx := NewTransaction(from[i], to[i], t_amount)
			if ntx == nil {
				return errors.New(fmt.Sprintf("%v转账%v : %v失败！", from[i], to[i], t_amount))
			}
			txs = append(txs, ntx)
		} else {
			return errors.New("输入参数有误！")
		}

	}

	//

	//2 建立新区块
	block := CreateBlock(blockChain.LastBlock.Index+1, blockChain.LastBlock.Hash, txs)
	dataBlock := block.Serialize()
	blockChain.Db.Write([]byte("l"), []byte(block.Hash))
	blockChain.Db.Write([]byte(block.Hash), dataBlock)
	blockChain.LastBlock = block
	return nil
}

func (blc *BlockChain) Add(txs []*Transaction) (*BlockChain, error) {
	idx := blc.LastBlock.Index + 1
	prehash := blc.LastBlock.Hash

	block := CreateBlock(idx, prehash, txs)
	blc.Db.Write([]byte(block.Hash), block.Serialize())
	blc.Db.Write([]byte("l"), []byte(block.Hash))

	blc.LastBlock = block
	return blc, nil
}
