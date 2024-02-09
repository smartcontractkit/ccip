package evm

const (
	// DefaultExecTxGasOverhead is the default gas overhead for exec transactions. This value plus the offchainconfig.BatchGasLimit
	// should be less than the gas limit of the chain. This is enforced by the offramp reader.
	DefaultExecTxGasOverhead = 1_000_000
)

// GetDefaultCommitTransactionGasLimit returns the default gas limit for a transaction on a given chain.
// Updating a price in the price registry costs <10k gas if the asset is already in the registry. If it's a new
// asset it will be <25k gas. There will be an OCR overhead of ~50k gas and a calldata/tx overhead of at most 50k.
// Given these values we would be able to handle about 36 price updates for new assets or 90 for existing assets.
func GetDefaultCommitTransactionGasLimit(chainId uint64) uint32 {
	switch chainId {
	case 6101244977088475029, 3478487238524512106, 4949039107694359620:
		return 15_000_000 // Arbitrum and its testnets
	default:
		return 1_000_000 // Default for all chains
	}
}

// GetDefaultExecTransactionGasLimit returns the default gas limit for a transaction on a given chain.
// The values defined here should be in sync with the offchainConfig.BatchGasLimit, which the value in
// the offchain config always being lower than the value defined here + the gas overhead of the transaction.
func GetDefaultExecTransactionGasLimit(chainId uint64) uint32 {
	switch chainId {
	case 6101244977088475029, 3478487238524512106, 4949039107694359620:
		return 100_000_000 // Arbitrum and its testnets
	default:
		return 8_000_000 // Default for all chains
	}
}
