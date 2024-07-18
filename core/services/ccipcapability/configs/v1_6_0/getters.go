package v1_6_0

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	configstypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/configs/types"
	v1_6_0_evm "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/configs/v1_6_0/evm"
)

// GetCCIPReaderContractReaderConfig returns the marshaled JSON configuration for the CCIP reader contract reader.
// A different configuration is returned based on chain family and based on whether we are setting up the chain
// reader to read a source or a destination chain.
func GetCCIPReaderContractReaderConfig(chainSelector cciptypes.ChainSelector, chainSide configstypes.ChainSide) ([]byte, error) {
	switch family := getFamily(chainSelector); family {
	case "EVM":
		cfg, err := v1_6_0_evm.CCIPReaderContractReaderConfig(chainSide)
		if err != nil {
			return nil, fmt.Errorf("failed to get EVM CCIP reader contract reader config: %w", err)
		}
		return json.Marshal(cfg)
	default:
		return nil, fmt.Errorf("unsupported chain family: %s", family)
	}
}

func GetChainWriterConfig(chainSelector cciptypes.ChainSelector, fromAddress string, maxGasPrice *assets.Wei) ([]byte, error) {
	switch family := getFamily(chainSelector); family {
	case "EVM":
		return json.Marshal(v1_6_0_evm.ChainWriterConfigRaw(common.HexToAddress(fromAddress), maxGasPrice))
	default:
		return nil, fmt.Errorf("unsupported chain family: %s", family)
	}
}

// TODO: move this API to chain-selectors library.
func getFamily(_ cciptypes.ChainSelector) string {
	return "EVM"
}
