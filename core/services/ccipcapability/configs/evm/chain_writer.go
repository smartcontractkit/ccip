package evm

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink/v2/common/txmgr"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	evmrelaytypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

var (
	offrampABI = evmtypes.MustGetABI(evm_2_evm_multi_offramp.EVM2EVMMultiOffRampABI)
)

func MustChainWriterConfig(fromAddress common.Address, maxGasPrice *assets.Wei) []byte {
	rawConfig := ChainWriterConfigRaw(fromAddress, maxGasPrice)
	encoded, err := json.Marshal(rawConfig)
	if err != nil {
		panic(fmt.Errorf("failed to marshal ChainWriterConfig: %w", err))
	}

	return encoded
}

// ChainWriterConfigRaw returns a ChainWriterConfig that can be used to transmit commit and execute reports.
func ChainWriterConfigRaw(fromAddress common.Address, maxGasPrice *assets.Wei) evmrelaytypes.ChainWriterConfig {
	return evmrelaytypes.ChainWriterConfig{
		Contracts: map[string]*evmrelaytypes.ContractConfig{
			consts.ContractNameOffRamp: {
				ContractABI: evm_2_evm_multi_offramp.EVM2EVMMultiOffRampABI,
				Configs: map[string]*evmrelaytypes.ChainWriterDefinition{
					consts.MethodCommit: {
						ChainSpecificName: mustGetMethodName("commit", offrampABI),
						FromAddress:       fromAddress,
						// TODO: probably need to fetch this from home chain config?
						GasLimit: 500_000,
					},
					consts.MethodExecute: {
						ChainSpecificName: mustGetMethodName("execute", offrampABI),
						FromAddress:       fromAddress,
						// TODO: probably need to fetch this from home chain config?
						GasLimit: 6_500_000,
					},
				},
			},
		},
		SendStrategy: txmgr.NewSendEveryStrategy(),
		MaxGasPrice:  maxGasPrice,
	}
}

func mustGetMethodName(name string, tabi abi.ABI) (methodName string) {
	m, ok := tabi.Methods[name]
	if !ok {
		panic(fmt.Sprintf("missing method %s in offrampABI", name))
	}
	return m.Name
}