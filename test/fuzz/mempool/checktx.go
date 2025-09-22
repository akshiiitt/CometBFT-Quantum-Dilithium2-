package reactor

import (
	"github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/abci/example/kvstore"
	"github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/config"
	mempl "github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/mempool"
	"github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/proxy"
)

var mempool mempl.Mempool

func init() {
	app := kvstore.NewInMemoryApplication()
	cc := proxy.NewLocalClientCreator(app)
	appConnMem, _ := cc.NewABCIClient()
	err := appConnMem.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false
	mempool = mempl.NewCListMempool(cfg, appConnMem, 0)
}

func Fuzz(data []byte) int {
	err := mempool.CheckTx(data, nil, mempl.TxInfo{})
	if err != nil {
		return 0
	}

	return 1
}
