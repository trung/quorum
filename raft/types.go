package raft

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

type SignedBlock struct {
	R     []byte
	S     []byte
	Block *types.Block
}

func (sb *SignedBlock) String() string {
	return fmt.Sprintf(`Signature: %x, %x
%s
`, sb.R, sb.S, sb.Block.String())
}
