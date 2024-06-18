package opstack

import (
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/optimism_l1_standard_bridge"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/optimism_standard_bridge"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
)

const (
	// ERC20BridgeFinalizedFromAddressTopicIndex is the index of the topic in the ERC20BridgeFinalized event
	// that contains the "from" address. In the case of an L1 to L2 transfer, this event will be emitted by the OP
	// StandardBridge on L2 and the "from" address should be the L1 bridge adapter contract address.
	ERC20BridgeFinalizedFromAddressTopicIndex = 3

	// Optimism stages
	// StageRebalanceConfirmed is set as the transfer stage when the rebalanceLiquidity tx is confirmed onchain, but
	// when it has not yet been finalized.
	StageRebalanceConfirmed = 1
	// StageFinalizeReady is set as the transfer stage when the finalization is ready to execute onchain.
	StageFinalizeReady = 2
	// StageFinalizeConfirmed is set as the transfer stage when the finalization is confirmed onchain.
	// This is a terminal stage.
	StageFinalizeConfirmed = 3

	// Function calls
	DepositETHToFunction = "depositETHTo"
)

var (
	// Optimism events emitted on L2
	ERC20BridgeFinalizedTopic = optimism_standard_bridge.OptimismStandardBridgeERC20BridgeFinalized{}.Topic()

	// ABIs
	l1standardBridgeABI = abihelpers.MustParseABI(optimism_l1_standard_bridge.OptimismL1StandardBridgeMetaData.ABI)
)
