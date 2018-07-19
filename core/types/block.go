package types

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"math/big"

	"github.com/guoxingx/gxchain/common"
	"github.com/guoxingx/gxchain/trie"
)

type BlockNonce [8]byte

func EncodeNonce(i uint64) BlockNonce {
	var n BlockNonce
	binary.BigEndian.PutUint64(n[:], i)
	return n
}

type Header struct {
	ParentHash common.Hash
	Miner      common.Address
	TxHash     common.Hash
	Number     *big.Int
	Timestamp  *big.Int
	Nonce      BlockNonce
}

type Block struct {
	Header       *Header
	Transactions []*Transaction
	Hash         common.Hash
}

// func (block *Block) Transactions() []*Transaction { return block.transactions }
// func (block *Block) Hash() common.Hash            { return block.hash }
// func (block *Block) Header() *Header              { return block.header }
func (block *Block) ParentHash() common.Hash { return block.Header.ParentHash }
func (block *Block) Miner() common.Address   { return block.Header.Miner }
func (block *Block) TxHash() common.Hash     { return block.Header.TxHash }
func (block *Block) Number() *big.Int        { return new(big.Int).Set(block.Header.Number) }
func (block *Block) Timestamp() *big.Int     { return new(big.Int).Set(block.Header.Timestamp) }
func (block *Block) Nonce() uint64           { return binary.BigEndian.Uint64(block.Header.Nonce[:]) }

// 将一个区块序列化
// @param: b: *Block: 区块
// @return: []byte
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	// encodind/gob.NewEncoder(w io.Writer)
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// 一个区块所有交易的hash
func (b *Block) HashTransactions() {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := trie.NewMerkleTree(transactions)

	b.Hash.SetBytes(mTree.RootNode.Data)
}
