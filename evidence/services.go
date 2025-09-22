package evidence

import (
	"github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/types"
)

//go:generate ../scripts/mockery_generate.sh BlockStore

type BlockStore interface {
	LoadBlockMeta(height int64) *types.BlockMeta
	LoadBlockCommit(height int64) *types.Commit
	Height() int64
}
