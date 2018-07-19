package core

import (
	"bytes"
	"encoding/gob"
	"log"
	"math/big"
	"time"

	"github.com/guoxingx/gxchain/common"
	"github.com/guoxingx/gxchain/consensus"
	"github.com/guoxingx/gxchain/core/types"
)

// 获取一个新区块
// @param: miner: []byte: 挖出区块的矿工
// @param: parent: *Block: 上一个区块
// @param: transactions: []*Transaction: 待写入的交易
// @return: *Block
func NewBlock(miner string, parent *types.Block, transactions []*types.Transaction) *types.Block {
	var parentHash common.Hash
	var blockNumber big.Int
	if parent != nil {
		parentHash = parent.Hash
		blockNumber = *new(big.Int).Add(parent.Number(), big.NewInt(1))
	}

	header := &types.Header{parentHash, common.HexToAddress(miner), common.Hash{}, &blockNumber, big.NewInt(time.Now().Unix()), types.BlockNonce{}}
	block := &types.Block{header, transactions, common.Hash{}}

	if len(transactions) > 0 {
		block.HashTransactions()
	}
	pow := consensus.NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Header.Nonce = types.EncodeNonce(uint64(nonce))
	block.Hash.SetBytes(hash)

	return block
}

// 获取创世块
// rewardTx 矿工的奖励交易，不需要引用之前交易。
// @return: *Block
// func NewGenesisBlock(miner common.Address, rewardTx *Transaction) *Block {
func NewGenesisBlock(miner string, rewardTx *types.Transaction) *types.Block {
	// return NewBlock(miner, nil, []*Transaction{})
	return NewBlock(miner, nil, []*types.Transaction{rewardTx})
}

func DeserializeBlock(d []byte) *types.Block {
	var block types.Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
