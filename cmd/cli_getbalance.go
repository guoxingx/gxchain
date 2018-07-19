package cmd

import (
	"fmt"
	"log"

	"github.com/guoxingx/gxchain/accounts"
	"github.com/guoxingx/gxchain/core"
)

func (cli *CLI) getBalance(address string) {
	if !accounts.ValidateAddress(address) {
		log.Panic("ERROR: Address is not Valid")
	}

	bc := core.NewBlockchain()
	u := &core.UTXOSet{bc}
	// defer bc.db.Close()

	balance := 0
	UTXOs := u.FindUTXO(accounts.AddressToPubKeyHash(address))

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("getBalance of '%s': %d\n", address, balance)
}
