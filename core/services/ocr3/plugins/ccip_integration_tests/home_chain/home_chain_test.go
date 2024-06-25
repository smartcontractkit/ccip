package home_chain

import (
	"context"
	_ "embed"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	capcfg "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ccip_capability_configuration"
	helpers "github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/ccip_integration_tests"

	evmtypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
	"github.com/stretchr/testify/assert"
)

const (
	chainID = 1337
	chainA  = ccipocr3.ChainSelector(1)
	chainB  = ccipocr3.ChainSelector(2)
	chainC  = ccipocr3.ChainSelector(3)
)

func TestHomeChainReader(t *testing.T) {
	deployFunc := func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *capcfg.CCIPCapabilityConfiguration, error) {
		return capcfg.DeployCCIPCapabilityConfiguration(auth, backend, common.Address{})
	}
	// Initialize chainReader
	cfg := evmtypes.ChainReaderConfig{
		Contracts: map[string]evmtypes.ChainContractReader{
			"CCIPCapabilityConfiguration": {
				ContractABI: capcfg.CCIPCapabilityConfigurationMetaData.ABI,
				Configs: map[string]*evmtypes.ChainReaderDefinition{
					"ChainConfigSet": {
						ChainSpecificName:       "ChainConfigSet",
						ReadType:                evmtypes.Event,
						ConfidenceConfirmations: map[string]int{"0.0": 0, "1.0": 0},
					},
					"getAllChainConfigs": {
						ChainSpecificName: "getAllChainConfigs",
					},
				},
			},
		},
	}

	d := helpers.SetupChainReaderTest[capcfg.CCIPCapabilityConfiguration](t, context.Background(), deployFunc, capcfg.NewCCIPCapabilityConfiguration, cfg)
	chainReader := *d.ChainReader

	// Apply chain configs to the contract
	inputConfig := setupConfigInfo()
	_, err := d.Contract.ApplyChainConfigUpdates(d.Auth, nil, inputConfig)
	d.SimulatedBE.Commit()
	assert.NoError(t, err)

	// Now read the contract using chain reader
	var allConfigs []capcfg.CCIPCapabilityConfigurationChainConfigInfo
	err = chainReader.GetLatestValue(
		context.Background(),
		"CCIPCapabilityConfiguration",
		"getAllChainConfigs",
		map[string]interface{}{},
		&allConfigs,
	)
	assert.NoError(t, err)
	assert.Equal(t, inputConfig, allConfigs)
}

func setupConfigInfo() []capcfg.CCIPCapabilityConfigurationChainConfigInfo {
	return []capcfg.CCIPCapabilityConfigurationChainConfigInfo{
		{
			ChainSelector: chainID,
			ChainConfig: capcfg.CCIPCapabilityConfigurationChainConfig{
				Readers: [][32]byte{},
				FChain:  2,
				Config:  []byte{1, 2, 3},
			},
		},
	}
}
