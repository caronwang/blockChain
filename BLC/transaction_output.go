package BLC

/*

 */
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

/*
	判断当前消费是否属于当前Address
*/
func (txOutput *TXOutput) UnlockWithAddress(addr string) bool {
	return txOutput.ScriptPubKey == addr
}
