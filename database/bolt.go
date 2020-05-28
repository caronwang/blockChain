package database

import (
	"log"

	"github.com/boltdb/bolt"
)

var (
	DB *bolt.DB
)

func init() {
	var err error
	DB, err = bolt.Open("blkchain.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("连接数据库成功！")

	//更新表数据
	err = DB.Update(func(tx *bolt.Tx) error {

		//创建表 BlockBucket
		// b, err := tx.CreateBucket([]byte("BlockBucket"))
		// if err != nil {
		// 	return err
		// }

		b := tx.Bucket([]byte("BlockBucket"))

		//存储数据
		if b != nil {
			err := b.Put([]byte("l"), []byte("testing"))
			if err != nil {
				log.Panic("数据存储失败！")
			}
		}

		return nil
	})
}
