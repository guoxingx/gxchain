package cmd

import (
	"fmt"

	"github.com/guoxingx/gxchain/accounts"
)

// 创建新账号
func (cli *CLI) createWallet() {
	wallets, _ := accounts.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("Your new address: %s\n", address)
}
