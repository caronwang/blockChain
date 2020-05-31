# 说明
我的第一个区块链项目。

# 快速开始
go run main.go
GET http://127.0.0.1:8080/  查看区块链数据
POST http://127.0.0.1:8080/  新增区块数据



# 区块链本地持久化
三方包github.com/boltdb/bolt
1.先将区块序列化为字节数组
2.以当前区块的hash为key，以序列化以后的字节数组为value
3.最新的区块在bucket中key以l存储，value值为block hash




