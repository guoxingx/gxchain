package cmd

import (
	"fmt"
	"strconv"

	"github.com/guoxingx/gxchain/common"
	"github.com/guoxingx/gxchain/consensus"
	"github.com/guoxingx/gxchain/core"
)

// print each block and validate pow.
func (cli *CLI) printChain() {
	bc := core.NewBlockchain()
	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ Block %v %x ============\n", block.Number(), block.Hash)
		fmt.Printf("Parent hash: %x\n", block.ParentHash())
		pow := consensus.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("Transactions: ")
		for _, tx := range block.Transactions {
			fmt.Printf("%x, ", tx.ID)
		}
		fmt.Println()
		fmt.Println()

		if (block.ParentHash() == common.Hash{}) {
			break
		}
	}
}
