package evm

import chainselectors "github.com/smartcontractkit/chain-selectors"

const (
	// DefaultExecTxGasOverhead is the default gas overhead for exec transactions. This value plus the offchainconfig.BatchGasLimit
	// should be less than the gas limit of the chain. This is enforced by the offramp reader.
	DefaultExecTxGasOverhead = 1_000_000
)

// GetDefaultCommitTransactionGasLimit returns the default gas limit for a transaction on a given chain.
// Updating a price in the price registry costs <10k gas if the asset is already in the registry. If it's a new
// asset it will be <25k gas. There will be an OCR overhead of ~50k gas and a calldata/tx overhead of at most 50k.
// Given these values we would be able to handle about 36 price updates for new assets or 90 for existing assets.
func GetDefaultCommitTransactionGasLimit(chainSelector uint64) uint32 {
	switch chainSelector {
	case chainselectors.ETHEREUM_MAINNET_ARBITRUM_1.Selector,
		chainselectors.ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1.Selector,
		chainselectors.ETHEREUM_TESTNET_GOERLI_ARBITRUM_1.Selector:
		return 15_000_000
	default:
		return 1_000_000 // Default for all chains
	}
}

// GetDefaultExecTransactionGasLimit returns the default gas limit for a transaction on a given chain.
// The values defined here should be in sync with the offchainConfig.BatchGasLimit, which the value in
// the offchain config always being lower than the value defined here + the gas overhead of the transaction.
func GetDefaultExecTransactionGasLimit(chainSelector uint64) uint32 {
	switch chainSelector {
	case chainselectors.ETHEREUM_MAINNET_ARBITRUM_1.Selector,
		chainselectors.ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1.Selector,
		chainselectors.ETHEREUM_TESTNET_GOERLI_ARBITRUM_1.Selector:
		return 100_000_000
	default:

		return 8_000_000 // Default for all chains
	}
}
