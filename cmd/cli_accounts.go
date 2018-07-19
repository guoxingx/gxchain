package cmd

import (
	"fmt"
	"log"

	"github.com/guoxingx/gxchain/accounts"
)

func (cli *CLI) accounts() {
	wallets, err := accounts.NewWallets()
	if err != nil {
		log.Panic(err)
	}

	accounts := wallets.GetAddresses()

	for _, account := range accounts {
		fmt.Printf("%v, ", account)
	}
	fmt.Println()
}
