package types

import (
	"bytes"

	"github.com/guoxingx/gxchain/accounts"
)

type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte // PubKey origin
}

func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := accounts.HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
