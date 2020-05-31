package database

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockChainDB struct {
	Conn      *bolt.DB
	TableName string
}

func NewBlockChainDB(dbPath, tableName string) *BlockChainDB {

	conn, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	bc := &BlockChainDB{
		Conn:      conn,
		TableName: tableName,
	}

	if !bc.HasTable() {
		bc.createTable()
		log.Println("创建表成功")
	}

	return bc
}

func (db *BlockChainDB) HasTable() bool {
	res := false
	_ = db.Conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.TableName))
		if b != nil {
			res = true
			return nil
		}
		return nil
	})
	return res
}

func (db *BlockChainDB) createTable() error {
	_ = db.Conn.Update(func(tx *bolt.Tx) error {

		_, err := tx.CreateBucket([]byte(db.TableName))
		if err != nil {
			log.Println(err)
			return err
		}

		return nil
	})
	return nil
}

/*
	写入Key
*/
func (db *BlockChainDB) Write(key []byte, value []byte) error {

	err := db.Conn.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(db.TableName))
		//存储数据
		if b != nil {
			err := b.Put(key, value)
			if err != nil {
				log.Panic("数据存储失败！")
			}
		} else {
			log.Panic(db.TableName, "不存在")
		}

		return nil
	})

	return err
}

func (db *BlockChainDB) Read(key []byte) []byte {
	var data []byte
	//spew.Dump(db)
	_ = db.Conn.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(db.TableName))
		//存储数据
		if b != nil {
			data = b.Get(key)
			return nil
		} else {

			log.Panic(db.TableName, "不存在")
		}

		return nil
	})
	return data
}
