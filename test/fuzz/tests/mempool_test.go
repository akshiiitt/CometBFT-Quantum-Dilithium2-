//go:build gofuzz || go1.21

package tests

import (
	"testing"

	abciclient "github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/abci/client"
	"github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/abci/example/kvstore"
	"github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/config"
	cmtsync "github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/libs/sync"
	mempool "github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/mempool"
)

func FuzzMempool(f *testing.F) {
	app := kvstore.NewInMemoryApplication()
	mtx := new(cmtsync.Mutex)
	conn := abciclient.NewLocalClient(mtx, app)
	err := conn.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false

	mp := mempool.NewCListMempool(cfg, conn, 0)

	f.Fuzz(func(t *testing.T, data []byte) {
		_ = mp.CheckTx(data, nil, mempool.TxInfo{})
	})
}
