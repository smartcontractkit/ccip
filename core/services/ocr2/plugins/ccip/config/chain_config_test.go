package config

import (
	"math/big"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	evmORMMocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

func TestGetChainFromSpec_success(t *testing.T) {
	spec := &job.OCR2OracleSpec{
		RelayConfig: job.JSONConfig{
			"chainID": float64(1337),
		},
	}

	mockChain := evmORMMocks.NewChain(t)
	mockChain.On("ID").Return(big.NewInt(1337))

	mockChainSet := evmORMMocks.NewLegacyChainContainer(t)
	mockChainSet.On("Get", "1337").Return(mockChain, nil)

	chain, chainID, err := GetChainFromSpec(spec, mockChainSet)
	require.NoError(t, err)
	require.Equal(t, mockChain, chain)
	require.Equal(t, int64(1337), chainID)
}

func TestGetChainFromSpec_missingChainID(t *testing.T) {
	spec := &job.OCR2OracleSpec{}

	mockChain := evmORMMocks.NewChain(t)
	mockChain.On("ID").Return(big.NewInt(1337)).Maybe()

	mockChainSet := evmORMMocks.NewLegacyChainContainer(t)
	mockChainSet.On("Get", "1337").Return(mockChain, nil).Maybe()

	_, _, err := GetChainFromSpec(spec, mockChainSet)
	require.Error(t, err)
	require.Contains(t, err.Error(), "chainID must be provided in relay config")
}

func TestGetChainByChainSelector_success(t *testing.T) {
	mockChain := evmORMMocks.NewChain(t)
	mockChain.On("ID").Return(big.NewInt(11155111))

	mockChainSet := evmORMMocks.NewLegacyChainContainer(t)
	mockChainSet.On("Get", "11155111").Return(mockChain, nil)

	// Ethereum Sepolia chain selector.
	chain, chainID, err := GetChainByChainSelector(mockChainSet, uint64(16015286601757825753))
	require.NoError(t, err)
	require.Equal(t, mockChain, chain)
	require.Equal(t, int64(11155111), chainID)
}

func TestGetChainByChainSelector_selectorNotFound(t *testing.T) {
	mockChainSet := evmORMMocks.NewLegacyChainContainer(t)

	_, _, err := GetChainByChainSelector(mockChainSet, uint64(444000444))
	require.Error(t, err)
}

func TestGetChainById_notFound(t *testing.T) {
	mockChainSet := evmORMMocks.NewLegacyChainContainer(t)
	mockChainSet.On("Get", "444").Return(nil, errors.New("test")).Maybe()

	_, _, err := GetChainByChainID(mockChainSet, uint64(444))
	require.Error(t, err)
	require.Contains(t, err.Error(), "chain not found in chainset")
}
