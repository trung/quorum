package raft

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

type SignedBlock struct {
	Signature []byte
	Block     *types.Block
}

func (sb *SignedBlock) String() string {
	return fmt.Sprintf(`Signature: %s
%s
`, string(sb.Signature), sb.Block.String())
}
