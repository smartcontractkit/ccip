package relay

import relaytypes "github.com/smartcontractkit/chainlink-relay/pkg/types"

type Network string

var (
	EVM             Network = "evm"
	Solana          Network = "solana"
	StarkNet        Network = "starknet"
	SupportedRelays         = map[Network]struct{}{
		EVM:      {},
		Solana:   {},
		StarkNet: {},
	}
)

type CCIPRelayer interface {
	relaytypes.Relayer
	NewCCIPRelayProvider(rargs relaytypes.RelayArgs, transmitterID string) (CCIPRelayProvider, error)
	NewCCIPExecutionProvider(rargs relaytypes.RelayArgs, transmitterID string) (CCIPExecutionProvider, error)
}

// CCIPRelayProvider provides all components needed for a CCIP Relay OCR2 plugin.
type CCIPRelayProvider interface {
	relaytypes.Plugin
}

// CCIPExecutionProvider provides all components needed for a CCIP Execution OCR2 plugin.
type CCIPExecutionProvider interface {
	relaytypes.Plugin
}
