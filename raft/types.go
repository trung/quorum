package raft

import "github.com/ethereum/go-ethereum/core/types"

type SignedBlock struct {
	Signature string
	Block *types.Block
}
