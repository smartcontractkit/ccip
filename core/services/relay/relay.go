package relay

import relaytypes "github.com/smartcontractkit/chainlink-relay/pkg/types"

type Network string

var (
	EVM             Network = "evm"
	Cosmos          Network = "cosmos"
	Solana          Network = "solana"
	StarkNet        Network = "starknet"
	SupportedRelays         = map[Network]struct{}{
		EVM:      {},
		Cosmos:   {},
		Solana:   {},
		StarkNet: {},
	}
)

type CCIPRelayer interface {
	relaytypes.Relayer
	NewCCIPCommitProvider(rargs relaytypes.RelayArgs, transmitterID string) (CCIPCommitProvider, error)
	NewCCIPExecutionProvider(rargs relaytypes.RelayArgs, transmitterID string) (CCIPExecutionProvider, error)
}

// CCIPCommitProvider provides all components needed for a CCIP Relay OCR2 plugin.
type CCIPCommitProvider interface {
	relaytypes.Plugin
}

// CCIPExecutionProvider provides all components needed for a CCIP Execution OCR2 plugin.
type CCIPExecutionProvider interface {
	relaytypes.Plugin
}
