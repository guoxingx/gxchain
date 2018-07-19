package core

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/guoxingx/gxchain/accounts"
	"github.com/guoxingx/gxchain/core/types"
)

const subsidy = 26

// 即区块的奖励交易
func NewRewardTx(to, data string) *types.Transaction {
	// 奖励交易没有输入 也不会被校验
	// 因此 TXInput.Signature = nil, TXInput.PubKey 随机生成
	// 根据 当前时间 和 随机数 生成 PubKey
	if data == "" {
		var ts bytes.Buffer
		binary.Write(&ts, binary.BigEndian, time.Now().UnixNano())

		randData := make([]byte, 20)
		_, err := rand.Read(randData)
		if err != nil {
			log.Panic(err)
		}

		randData = append(ts.Bytes(), randData...)
		data = fmt.Sprintf("%v", randData)
	}

	txin := types.TXInput{[]byte{}, -1, nil, []byte(data)}

	txout := types.NewTXOutput(subsidy, to)
	tx := types.Transaction{nil, []types.TXInput{txin}, []types.TXOutput{*txout}}
	tx.ID = tx.Hash()

	return &tx
}

// 发起交易
func NewUTXOTransaction(from, to string, amount int, UTXOSet *UTXOSet) *types.Transaction {
	var inputs []types.TXInput
	var outputs []types.TXOutput

	// 获取的未花费输出 总额 & UTXOs
	wallets, err := accounts.NewWallets()
	if err != nil {
		log.Panic(err)
	}

	// 从wallet获取address对应的pubKeyHash
	wallet := wallets.GetWallet(from)
	acc, validOutputs := UTXOSet.FindSpendableOutputs(wallet.GetHashPubKey(), amount)

	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}

	// 花费：将获取到的每一个输出都引用并创建一个新的输入
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := types.TXInput{txID, out, nil, wallet.PublicKey}
			inputs = append(inputs, input)
		}
	}

	// 转账 amount 到 to 的输出
	outputs = append(outputs, *types.NewTXOutput(amount, to))

	// 转账 acc - amount 的 from 的输出，即找零
	if acc > amount {
		outputs = append(outputs, *types.NewTXOutput(acc-amount, from)) // a change
	}

	tx := types.Transaction{nil, inputs, outputs}

	// 交易签名
	tx.ID = tx.Hash()
	UTXOSet.Blockchain.SignTransaction(&tx, wallet.PrivateKey)

	return &tx
}
