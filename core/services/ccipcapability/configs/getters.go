package configs

import (
	"fmt"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	configstypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/configs/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/configs/v1_6_0"
	v1_6_0_evm "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/configs/v1_6_0/evm"
	evmrelaytypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

func GetCCIPReaderContractReaderConfig(
	ccipCapVersion string,
	chainSelector cciptypes.ChainSelector,
	chainSide configstypes.ChainSide,
) ([]byte, error) {
	switch ccipCapVersion {
	case "v1.6.0":
		return v1_6_0.GetCCIPReaderContractReaderConfig(chainSelector, chainSide)
	default:
		return nil, fmt.Errorf("unsupported ccip capability version: %s", ccipCapVersion)
	}
}

func GetHomeChainContractReaderConfig(
	ccipCapVersion string,
) (evmrelaytypes.ChainReaderConfig, error) {
	switch ccipCapVersion {
	case "v1.6.0":
		return v1_6_0_evm.HomeChainReaderConfigRaw(), nil
	default:
		return evmrelaytypes.ChainReaderConfig{}, fmt.Errorf("unsupported ccip capability version: %s", ccipCapVersion)
	}
}

func GetChainWriterConfig(
	ccipCapVersion string,
	chainSelector cciptypes.ChainSelector,
	fromAddress string,
	maxGasPrice *assets.Wei,
) ([]byte, error) {
	switch ccipCapVersion {
	case "v1.6.0":
		return v1_6_0.GetChainWriterConfig(chainSelector, fromAddress, maxGasPrice)
	default:
		return nil, fmt.Errorf("unsupported ccip capability version: %s", ccipCapVersion)
	}
}
