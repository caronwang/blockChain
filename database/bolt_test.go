package database

import (
	"fmt"
	"testing"
)

// func TestBasic(t *testing.T) {
// 	var err error
// 	DB, err = bolt.Open("blkchain.db", 0600, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer DB.Close()
// 	log.Println("连接数据库成功！")

// 	err = DB.Update(func(tx *bolt.Tx) error {

// 		//创建表 BlockBucket
// 		b, err := tx.CreateBucket([]byte("BlockBucket"))
// 		if err != nil {
// 			return err
// 		}

// 		//存储数据
// 		if b != nil {
// 			err := b.Put([]byte("l"), []byte("testing"))
// 			if err != nil {
// 				log.Panic("数据存储失败！")
// 			}
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		log.Panic(err)
// 	}
// }

func TestNewBlockChainDB(t *testing.T) {
	bc := NewBlockChainDB("test.db", "info")
	bc.Write([]byte("test"), []byte("nihao"))
	fmt.Println(string(bc.Read([]byte("test"))))
}
