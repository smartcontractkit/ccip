package ccipevents

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

// LogPollerClient implements the Client interface by using a logPoller instance to fetch the events.
type LogPollerClient struct {
	lp     logpoller.LogPoller
	lggr   logger.Logger
	client evmclient.Client
}

func NewLogPollerClient(lp logpoller.LogPoller, lggr logger.Logger, client evmclient.Client) *LogPollerClient {
	return &LogPollerClient{
		lp:     lp,
		lggr:   lggr,
		client: client,
	}
}

func (c *LogPollerClient) GetSendRequestsAfterNextMin(ctx context.Context, onRampAddress common.Address, nextMin uint64, confs int, checkFinalityTags bool) (sendReqs []Event[evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested], err error) {
	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddress, c.client)
	if err != nil {
		return nil, err
	}

	var logs []logpoller.Log

	if !checkFinalityTags {
		logs, err = c.lp.LogsDataWordGreaterThan(
			abihelpers.EventSignatures.SendRequested,
			onRampAddress,
			abihelpers.EventSignatures.SendRequestedSequenceNumberWord,
			abihelpers.EvmWord(nextMin),
			confs,
			pg.WithParentCtx(ctx),
		)
		if err != nil {
			return nil, fmt.Errorf("logs data word greater than: %w", err)
		}
	} else {
		// If the chain is based on explicit finality we only examine logs less than or equal to the latest finalized block number.
		// NOTE: there appears to be a bug in ethclient whereby BlockByNumber fails with "unsupported txtype" when trying to parse the block
		// when querying L2s, headers however work.
		// TODO (CCIP-778): Migrate to core finalized tags, below doesn't work for some chains e.g. Celo.
		latestFinalizedHeader, err := c.client.HeaderByNumber(
			ctx,
			big.NewInt(rpc.FinalizedBlockNumber.Int64()),
		)
		if err != nil {
			return nil, err
		}

		if latestFinalizedHeader == nil {
			return nil, errors.New("latest finalized header is nil")
		}
		if latestFinalizedHeader.Number == nil {
			return nil, errors.New("latest finalized number is nil")
		}
		logs, err = c.lp.LogsUntilBlockHashDataWordGreaterThan(
			abihelpers.EventSignatures.SendRequested,
			onRampAddress,
			abihelpers.EventSignatures.SendRequestedSequenceNumberWord,
			abihelpers.EvmWord(nextMin),
			latestFinalizedHeader.Hash(),
			pg.WithParentCtx(ctx),
		)
		if err != nil {
			return nil, fmt.Errorf("logs until block hash data word greater than: %w", err)
		}
	}

	return convertLogsToRequests[evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested](
		logs,
		c.lggr,
		func(log types.Log) (*evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested, error) {
			return onRamp.ParseCCIPSendRequested(log)
		},
	)
}

func (c *LogPollerClient) GetSendRequestsInSeqNumRange(ctx context.Context, onRampAddress common.Address, rangeMin, rangeMax uint64, confs int) ([]Event[evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested], error) {
	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddress, c.client)
	if err != nil {
		return nil, err
	}

	logs, err := c.lp.LogsDataWordRange(
		abihelpers.EventSignatures.SendRequested,
		onRampAddress,
		abihelpers.EventSignatures.SendRequestedSequenceNumberWord,
		logpoller.EvmWord(rangeMin),
		logpoller.EvmWord(rangeMax),
		confs,
		pg.WithParentCtx(ctx))
	if err != nil {
		return nil, err
	}

	return convertLogsToRequests[evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested](
		logs,
		c.lggr,
		func(log types.Log) (*evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested, error) {
			return onRamp.ParseCCIPSendRequested(log)
		},
	)
}

func (c *LogPollerClient) GetTokenPriceUpdatesCreatedAfter(ctx context.Context, priceRegistryAddress common.Address, ts time.Time, confs int) ([]Event[price_registry.PriceRegistryUsdPerTokenUpdated], error) {
	priceRegistry, err := price_registry.NewPriceRegistry(priceRegistryAddress, c.client)
	if err != nil {
		return nil, err
	}

	logs, err := c.lp.LogsCreatedAfter(
		abihelpers.EventSignatures.UsdPerTokenUpdated,
		priceRegistryAddress,
		ts,
		confs,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, err
	}

	return convertLogsToRequests[price_registry.PriceRegistryUsdPerTokenUpdated](
		logs,
		c.lggr,
		func(log types.Log) (*price_registry.PriceRegistryUsdPerTokenUpdated, error) {
			return priceRegistry.ParseUsdPerTokenUpdated(log)
		},
	)
}

func (c *LogPollerClient) GetGasPriceUpdatesCreatedAfter(ctx context.Context, priceRegistryAddress common.Address, chainSelector uint64, ts time.Time, confs int) ([]Event[price_registry.PriceRegistryUsdPerUnitGasUpdated], error) {
	priceRegistry, err := price_registry.NewPriceRegistry(priceRegistryAddress, c.client)
	if err != nil {
		return nil, err
	}

	logs, err := c.lp.IndexedLogsCreatedAfter(
		abihelpers.EventSignatures.UsdPerUnitGasUpdated,
		priceRegistryAddress,
		1,
		[]common.Hash{abihelpers.EvmWord(chainSelector)},
		ts,
		confs,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, err
	}

	return convertLogsToRequests[price_registry.PriceRegistryUsdPerUnitGasUpdated](
		logs,
		c.lggr,
		func(log types.Log) (*price_registry.PriceRegistryUsdPerUnitGasUpdated, error) {
			return priceRegistry.ParseUsdPerUnitGasUpdated(log)
		},
	)
}

func (c *LogPollerClient) GetExecutionStateChangesInRange(ctx context.Context, offRampAddress common.Address, rangeMin, rangeMax uint64, confs int) ([]Event[evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged], error) {
	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(offRampAddress, c.client)
	if err != nil {
		return nil, err
	}

	logs, err := c.lp.IndexedLogsTopicRange(
		abihelpers.EventSignatures.ExecutionStateChanged,
		offRampAddress,
		abihelpers.EventSignatures.ExecutionStateChangedSequenceNumberIndex,
		logpoller.EvmWord(rangeMin),
		logpoller.EvmWord(rangeMax),
		confs,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, err
	}

	return convertLogsToRequests[evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged](
		logs,
		c.lggr,
		func(log types.Log) (*evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged, error) {
			return offRamp.ParseExecutionStateChanged(log)
		},
	)
}

func (c *LogPollerClient) LatestBlock(ctx context.Context) (int64, error) {
	return c.lp.LatestBlock(pg.WithParentCtx(ctx))
}

func convertLogsToRequests[T any](logs []logpoller.Log, lggr logger.Logger, parse func(log types.Log) (*T, error)) ([]Event[T], error) {
	reqs := make([]Event[T], 0, len(logs))
	for _, log := range logs {
		data, err := parse(log.ToGethLog())
		if err == nil {
			reqs = append(reqs, Event[T]{
				Data: *data,
				BlockMeta: BlockMeta{
					BlockTimestamp: log.BlockTimestamp,
					BlockNumber:    log.BlockNumber,
				},
			})
		}
	}
	if len(logs) != len(reqs) {
		lggr.Warnw("Some logs were not parsed", "logs", len(logs), "requests", len(reqs))
	}

	return reqs, nil
}
