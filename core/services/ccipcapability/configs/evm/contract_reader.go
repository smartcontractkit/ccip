package evm

import (
	"encoding/json"
	"fmt"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_onramp"
	evmrelaytypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

// MustSourceReaderConfig returns a ChainReaderConfig that can be used to read from the CCIP contracts
// on both the source and destination chains. The configuration is marshaled into JSON so that it can be passed
// to the relayer NewContractReader() method.
func MustSourceReaderConfig() []byte {
	rawConfig := SourceReaderConfig()
	encoded, err := json.Marshal(rawConfig)
	if err != nil {
		panic(fmt.Errorf("failed to marshal ChainReaderConfig into JSON: %w", err))
	}

	return encoded
}

// DestReaderConfig returns a ChainReaderConfig that can be used to read from the offramp.
func DestReaderConfig() evmrelaytypes.ChainReaderConfig {
	return evmrelaytypes.ChainReaderConfig{
		Contracts: map[string]evmrelaytypes.ChainContractReader{
			"offRamp": {
				ContractABI: evm_2_evm_multi_offramp.EVM2EVMMultiOffRampABI,
				ContractPollingFilter: evmrelaytypes.ContractPollingFilter{
					GenericEventNames: []string{
						"ExecutionStateChanged",
						"CommitReportAccepted",
					},
				},
				Configs: map[string]*evmrelaytypes.ChainReaderDefinition{
					"getExecutionState": {
						ChainSpecificName: "getExecutionState",
						ReadType:          evmrelaytypes.Method,
					},
					"getMerkleRoot": {
						ChainSpecificName: "getMerkleRoot",
						ReadType:          evmrelaytypes.Method,
					},
					"isBlessed": {
						ChainSpecificName: "isBlessed",
						ReadType:          evmrelaytypes.Method,
					},
					"getLatestPriceSequenceNumber": {
						ChainSpecificName: "getLatestPriceSequenceNumber",
						ReadType:          evmrelaytypes.Method,
					},
					"getStaticConfig": {
						ChainSpecificName: "getStaticConfig",
						ReadType:          evmrelaytypes.Method,
					},
					"getDynamicConfig": {
						ChainSpecificName: "getDynamicConfig",
						ReadType:          evmrelaytypes.Method,
					},
					"getSourceChainConfig": {
						ChainSpecificName: "getSourceChainConfig",
						ReadType:          evmrelaytypes.Method,
					},
				},
			},
		},
	}
}

// SourceReaderConfig returns a ChainReaderConfig that can be used to read from the onramp.
func SourceReaderConfig() evmrelaytypes.ChainReaderConfig {
	return evmrelaytypes.ChainReaderConfig{
		Contracts: map[string]evmrelaytypes.ChainContractReader{
			"onRamp": {
				ContractABI: evm_2_evm_multi_onramp.EVM2EVMMultiOnRampABI,
				ContractPollingFilter: evmrelaytypes.ContractPollingFilter{
					GenericEventNames: []string{
						"CCIPSendRequested",
					},
				},
				Configs: map[string]*evmrelaytypes.ChainReaderDefinition{
					// all "{external|public} view" functions in the onramp except for getFee and getPoolBySourceToken are here.
					// getFee is not expected to get called offchain and is only called by end-user contracts.
					"getExpectedNextSequenceNumber": {
						ChainSpecificName: "getExpectedNextSequenceNumber",
						ReadType:          evmrelaytypes.Method,
					},
					"getStaticConfig": {
						ChainSpecificName: "getStaticConfig",
						ReadType:          evmrelaytypes.Method,
					},
					"getDynamicConfig": {
						ChainSpecificName: "getDynamicConfig",
						ReadType:          evmrelaytypes.Method,
					},
					"getDestChainConfig": {
						ChainSpecificName: "getDestChainConfig",
						ReadType:          evmrelaytypes.Method,
					},
					"getPremiumMultiplierWeiPerEth": {
						ChainSpecificName: "getPremiumMultiplierWeiPerEth",
						ReadType:          evmrelaytypes.Method,
					},
					"getTokenTransferFeeConfig": {
						ChainSpecificName: "getTokenTransferFeeConfig",
						ReadType:          evmrelaytypes.Method,
					},
				},
			},
		},
	}
}
