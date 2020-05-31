package BLC

import (
	"blockChain/database"
	"fmt"
	"log"
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
	if db.Read([]byte("l")) == nil {
		log.Println("创建创世区块")
		blc := CreateBlock(1, "", nil)

		err := db.Write([]byte("l"), []byte(blc.Hash))
		if err != nil {
			log.Panic("last hash数据存储失败！")
		}

		err = db.Write([]byte(blc.Hash), blc.Serialize())
		if err != nil {
			log.Panic("区块数据存储失败！")
		}

		lastBlock = blc
	} else {
		lb := db.Read([]byte(db.Read([]byte("l"))))
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

	//spew.Dump(blockChain)
	blockChain.Db.Read([]byte("l"))
	log.Println("区块链初始化完毕!")
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

func (blc *BlockChain) Add(txs []*Transaction) (*BlockChain, error) {
	idx := blc.LastBlock.Index + 1
	prehash := blc.LastBlock.Hash

	block := CreateBlock(idx, prehash, txs)
	blc.Db.Write([]byte(block.Hash), block.Serialize())
	blc.Db.Write([]byte("l"), []byte(block.Hash))

	blc.LastBlock = block
	return blc, nil
}
