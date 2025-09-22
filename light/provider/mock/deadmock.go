package mock

import (
	"context"

	"github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/light/provider"
	"github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/types"
)

type deadMock struct {
	chainID string
}

// NewDeadMock creates a mock provider that always errors.
func NewDeadMock(chainID string) provider.Provider {
	return &deadMock{chainID: chainID}
}

func (p *deadMock) ChainID() string { return p.chainID }

func (p *deadMock) String() string { return "deadMock" }

func (p *deadMock) LightBlock(context.Context, int64) (*types.LightBlock, error) {
	return nil, provider.ErrNoResponse
}

func (p *deadMock) ReportEvidence(context.Context, types.Evidence) error {
	return provider.ErrNoResponse
}
