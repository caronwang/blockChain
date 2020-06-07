# 说明
我的第一个区块链项目。

# 快速开始
go run main.go
- GET http://127.0.0.1:8080/  查看区块链数据
- POST http://127.0.0.1:8080/trans {from:["li",...],to:["wang",...],amount:["1",...]}  用户交易
- POST http://127.0.0.1:8080/balance?addr=wang 查看用户拥有的token


# 区块链本地持久化
三方包github.com/boltdb/bolt
1.先将区块序列化为字节数组
2.以当前区块的hash为key，以序列化以后的字节数组为value
3.最新的区块在bucket中key以l存储，value值为block hash


# 区块中的交易信息
包含3部分，交易Hash，交易input,交易的output

# 钱包
创建一个钱包地址:
1. 生层一对公钥和私钥
2. 想获取地址，可以通过公钥进行Base58编码
3. 想要别人给我转账，把地址给别人，别人将地址反编码变成公钥，将公钥和数据进行签名
4. 通过私钥进行解密，只有用户私钥的人才能解密

