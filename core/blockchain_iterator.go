package core

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/guoxingx/gxchain/core/types"
)

// 区块链迭代
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// 其实是查找上一个区块
func (i *BlockchainIterator) Next() *types.Block {
	var block *types.Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.ParentHash().Bytes()

	return block
}
