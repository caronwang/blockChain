package BLC

import (
	"testing"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)


func printChain(chain *BlockChain){
	for _, block := range chain.Chain {
		data := block.Serialize()
		fmt.Printf("Serialize:\n%x\n",data)

		bk := DeserializeBlock(data)
		fmt.Printf("Deserialize:\n")
		spew.Dump(bk)
	}
}

func TestSerialize(t *testing.T){
	chain:= NewBlockChain()
	
	printChain(chain)
}