package raft

import "github.com/ethereum/go-ethereum/core/types"

type SignedBlock struct {
	Signature []byte
	Block *types.Block
}
