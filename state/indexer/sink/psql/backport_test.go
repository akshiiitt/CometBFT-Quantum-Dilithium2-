package psql

import (
	"github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/state/indexer"
	"github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
