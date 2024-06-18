package common

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/liquiditymanager"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/models"
)

const (
	DurationMonth = 720 * time.Hour

	// TODO: these index values might need to be updated when Ryan makes changes to the LM contract and event fields
	// LiquidityTransferredToChainSelectorTopicIndex is the index of the topic in the LiquidityTransferred event
	// that contains the "to" chain selector.
	LiquidityTransferredToChainSelectorTopicIndex = 3
	// LiquidityTransferredFromChainSelectorTopicIndex is the index of the topic in the LiquidityTransferred event
	// that contains the "from" chain selector.
	LiquidityTransferredFromChainSelectorTopicIndex = 2
	// DepositFinalizedToAddressTopicIndex is the index of the topic in the DepositFinalized event
	// that contains the "to" address.
	DepositFinalizedToAddressTopicIndex = 3
)

var (
	// LiquidityManager event - emitted on both L1 and L2
	LiquidityTransferredTopic = liquiditymanager.LiquidityManagerLiquidityTransferred{}.Topic()
)

func NetworkSelectorToHash(selector models.NetworkSelector) common.Hash {
	encoded := hexutil.EncodeUint64(uint64(selector))
	return common.HexToHash(encoded)
}

type LogKey struct {
	TxHash   common.Hash
	LogIndex int64
}

func ParseLiquidityTransferred(parseFunc func(gethtypes.Log) (*liquiditymanager.LiquidityManagerLiquidityTransferred, error), lgs []logpoller.Log) ([]*liquiditymanager.LiquidityManagerLiquidityTransferred, map[LogKey]logpoller.Log, error) {
	transferred := make([]*liquiditymanager.LiquidityManagerLiquidityTransferred, len(lgs))
	toLP := make(map[LogKey]logpoller.Log)
	for i, lg := range lgs {
		parsed, err := parseFunc(lg.ToGethLog())
		if err != nil {
			// should never happen
			return nil, nil, fmt.Errorf("parse LiquidityTransferred log: %w", err)
		}
		transferred[i] = parsed
		toLP[LogKey{
			TxHash:   lg.TxHash,
			LogIndex: lg.LogIndex,
		}] = lg
	}
	return transferred, toLP, nil
}

func GetBridgeFilterName(bridgeName, filterLayer string, liquidityManagerAddress common.Address, localChain, remoteChain, extra string) string {
	filterName := fmt.Sprintf("%s-%s_LiquidityManager:%s_LocalChain:%s_RemoteChain:%s",
		filterLayer,
		bridgeName,
		liquidityManagerAddress.Hex(),
		localChain,
		remoteChain,
	)
	if extra != "" {
		filterName = fmt.Sprintf("%s_%s", filterName, extra)
	}
	return filterName
}
