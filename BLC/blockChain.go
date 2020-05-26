package BLC

type BlockChain struct {
	Chain []*Block
}

var blockChain *BlockChain

func NewBlockChain() *BlockChain {
	if blockChain == nil {
		blockChain = &BlockChain{}
		blc := CreateBlock(1, "", "genesys block")
		blockChain.Chain = append(blockChain.Chain, blc)
	}

	return blockChain
}

func GetChain() *BlockChain {
	return blockChain
}

func (blc *BlockChain) Add(data string) (*BlockChain, error) {
	idx := blc.Chain[len(blc.Chain)-1].Index + 1
	prehash := blc.Chain[len(blc.Chain)-1].Hash

	block := CreateBlock(idx, prehash, data)
	blc.Chain = append(blc.Chain, block)
	return blc, nil
}
