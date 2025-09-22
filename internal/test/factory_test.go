package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/akshiiitt/CometBFT-Quantum-Dilithium2-/types"
)

func TestMakeHeader(t *testing.T) {
	header := MakeHeader(t, &types.Header{})
	require.NotNil(t, header)

	require.NoError(t, header.ValidateBasic())
}
