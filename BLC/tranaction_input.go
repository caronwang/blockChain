package BLC

type TXInput struct {
	//交易hash
	TxHash []byte
	// 存储Txoutput在Vout里面的索引
	Vout int
	//用户数字签名
	ScriptSign string
}

/*
	判断当前消费是否属于当前Address
*/
func (txIntput *TXInput) UnlockWithAddress(addr string) bool {
	return txIntput.ScriptSign == addr
}
