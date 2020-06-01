package BLC

import "encoding/hex"

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

/*
	查看TXOutput是否已经被消费
*/
func (txOutput *TXOutput) IsSpend(addr string, spendTXOutputs map[string][]int) bool {
	for index, out := range txOutput {
		if out.UnlockWithAddress(addr) {
			if spendTXOutputs != nil {
				for txHash, indexArray := range spendTXOutputs {
					if txHash == hex.EncodeToString(hash) {
						for _, i := range indexArray {
							if index == i {
								return true
							}
						}
					}
				}
			}
		}
	}

	return false
}
