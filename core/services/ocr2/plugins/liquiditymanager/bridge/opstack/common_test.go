package opstack

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/liquiditymanager"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/optimism_standard_bridge"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	bridgecommon "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/bridge/common"
)

func Test_TopicIndexes(t *testing.T) {
	var (
		rebalancerABI     = abihelpers.MustParseABI(liquiditymanager.LiquidityManagerMetaData.ABI)
		standardBridgeABI = abihelpers.MustParseABI(optimism_standard_bridge.OptimismStandardBridgeMetaData.ABI)
	)
	t.Run("liquidity transferred to chain selector idx", func(t *testing.T) {
		ltEvent, ok := rebalancerABI.Events["LiquidityTransferred"]
		require.True(t, ok)

		var toChainSelectorArg abi.Argument
		var topicIndex = 0
		for _, arg := range ltEvent.Inputs {
			if arg.Indexed {
				topicIndex++
			}
			if arg.Name == "toChainSelector" {
				toChainSelectorArg = arg
				break
			}
		}

		require.True(t, toChainSelectorArg.Indexed)
		require.Equal(t, bridgecommon.LiquidityTransferredToChainSelectorTopicIndex, topicIndex)
	})

	t.Run("liquidity transferred from chain selector idx", func(t *testing.T) {
		ltEvent, ok := rebalancerABI.Events["LiquidityTransferred"]
		require.True(t, ok)

		var fromChainSelectorArg abi.Argument
		var topicIndex = 0
		for _, arg := range ltEvent.Inputs {
			if arg.Indexed {
				topicIndex++
			}
			if arg.Name == "fromChainSelector" {
				fromChainSelectorArg = arg
				break
			}
		}

		require.True(t, fromChainSelectorArg.Indexed)
		require.Equal(t, bridgecommon.LiquidityTransferredFromChainSelectorTopicIndex, topicIndex)
	})

	t.Run("ERC20 bridge finalized to address idx", func(t *testing.T) {
		bfEvent, ok := standardBridgeABI.Events["ERC20BridgeFinalized"]
		require.True(t, ok)

		var fromAddressArg abi.Argument
		var topicIndex = 0
		for _, arg := range bfEvent.Inputs {
			if arg.Indexed {
				topicIndex++
			}
			if arg.Name == "from" {
				fromAddressArg = arg
				break
			}
		}

		require.True(t, fromAddressArg.Indexed)
		require.Equal(t, ERC20BridgeFinalizedFromAddressTopicIndex, topicIndex)
	})
}
