package evm

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/v2/common/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	evmrelaytypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

func MustChainWriterConfig(fromAddress common.Address) []byte {
	rawConfig := ChainWriterConfigRaw(fromAddress)
	encoded, err := json.Marshal(rawConfig)
	if err != nil {
		panic(fmt.Errorf("failed to marshal ChainWriterConfig: %w", err))
	}

	return encoded
}

// ChainWriterConfigRaw returns a ChainWriterConfig that can be used to transmit commit and execute reports.
func ChainWriterConfigRaw(fromAddress common.Address) evmrelaytypes.ChainWriterConfig {
	return evmrelaytypes.ChainWriterConfig{
		Contracts: map[string]*evmrelaytypes.ContractConfig{
			"offRamp": {
				ContractABI: evm_2_evm_multi_offramp.EVM2EVMMultiOffRampABI,
				Configs: map[string]*evmrelaytypes.ChainWriterDefinition{
					"commit": {
						ChainSpecificName: "commit",
						FromAddress:       fromAddress,
						// TODO: probably need to fetch this from home chain config?
						GasLimit: 500_000,
					},
					"execute": {
						ChainSpecificName: "execute",
						FromAddress:       fromAddress,
						// TODO: probably need to fetch this from home chain config?
						GasLimit: 6_500_000,
					},
				},
			},
		},
		SendStrategy: txmgr.NewSendEveryStrategy(),
	}
}
