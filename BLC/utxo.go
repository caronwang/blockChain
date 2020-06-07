package BLC

type UTXO struct {
	Txhash   []byte
	Index    int
	TxOutPut TXOutput
}
