package core

import (
	"context"

	abci "github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/abci/types"
	"github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/libs/bytes"
	"github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/proxy"
	ctypes "github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/rpc/core/types"
	rpctypes "github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/rpc/jsonrpc/types"
)

// ABCIQuery queries the application for some information.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/ABCI/abci_query
func (env *Environment) ABCIQuery(
	_ *rpctypes.Context,
	path string,
	data bytes.HexBytes,
	height int64,
	prove bool,
) (*ctypes.ResultABCIQuery, error) {
	resQuery, err := env.ProxyAppQuery.Query(context.TODO(), &abci.RequestQuery{
		Path:   path,
		Data:   data,
		Height: height,
		Prove:  prove,
	})
	if err != nil {
		return nil, err
	}

	return &ctypes.ResultABCIQuery{Response: *resQuery}, nil
}

// ABCIInfo gets some info about the application.
// More: https://docs.cometbft.com/v0.38.x/rpc/#/ABCI/abci_info
func (env *Environment) ABCIInfo(_ *rpctypes.Context) (*ctypes.ResultABCIInfo, error) {
	resInfo, err := env.ProxyAppQuery.Info(context.TODO(), proxy.RequestInfo)
	if err != nil {
		return nil, err
	}

	return &ctypes.ResultABCIInfo{Response: *resInfo}, nil
}
