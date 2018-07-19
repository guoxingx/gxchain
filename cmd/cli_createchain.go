package cmd

import (
	"fmt"
	"log"

	"github.com/guoxingx/gxchain/accounts"
	"github.com/guoxingx/gxchain/core"
)

func (cli *CLI) createChain(address string) {
	if !accounts.ValidateAddress(address) {
		log.Panic("Error: Address is not valid")
	}

	bc := core.CreateBlockchain(address)
	// defer bc.db.Close()

	UTXOSet := core.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
