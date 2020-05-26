# 说明
我的第一个区块链项目。


# 实现区块的序列化和反序列化


# 区块链的存储
三方包github.com/boltdb/bolt
1.先将区块序列化为字节数组
2.以当前区块的hash为key，以序列化以后的字节数组为value
db.put(block.Hash,block.Seri())


# 实现命令行工具


