package cmd

import (
	"fmt"

	"github.com/guoxingx/gxchain/core"
	"github.com/guoxingx/gxchain/core/types"
)

func (cli *CLI) send(from, to string, amount int) {
	bc := core.NewBlockchain()
	u := &core.UTXOSet{bc}
	// defer u.Blockchain.db.Close()

	tx := core.NewUTXOTransaction(from, to, amount, u)

	newBlock := bc.MineBlock(from, []*types.Transaction{tx})
	u.Update(newBlock)
	fmt.Println("success!")
}
