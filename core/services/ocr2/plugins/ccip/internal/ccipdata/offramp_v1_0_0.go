package ccipdata

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/logpollerutil"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

const (
	EXEC_EXECUTION_STATE_CHANGES = "Exec execution state changes"
	EXEC_TOKEN_POOL_ADDED        = "Token pool added"
	EXEC_TOKEN_POOL_REMOVED      = "Token pool removed"
)

type OffRampV1_0_0 struct {
	offRamp *evm_2_evm_offramp.EVM2EVMOffRamp
	addr    common.Address
	lp      logpoller.LogPoller
	lggr    logger.Logger
}

func (o OffRampV1_0_0) GetDestinationTokens(ctx context.Context) ([]common.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (o OffRampV1_0_0) Close(qopts ...pg.QOpt) error {
	//TODO implement me
	panic("implement me")
}

func (o OffRampV1_0_0) GetExecutionStateChangesBetweenSeqNums(ctx context.Context, seqNumMin, seqNumMax uint64, confs int) ([]Event[ExecutionStateChanged], error) {
	logs, err := o.lp.IndexedLogsTopicRange(
		abihelpers.EventSignatures.ExecutionStateChanged,
		o.addr,
		abihelpers.EventSignatures.ExecutionStateChangedSequenceNumberIndex,
		logpoller.EvmWord(seqNumMin),
		logpoller.EvmWord(seqNumMax),
		confs,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, err
	}

	return parseLogs[ExecutionStateChanged](
		logs,
		o.lggr,
		func(log types.Log) (*ExecutionStateChanged, error) {
			sc, err := o.offRamp.ParseExecutionStateChanged(log)
			if err != nil {
				return nil, err
			}
			return &ExecutionStateChanged{SequenceNumber: sc.SequenceNumber}, nil
		},
	)
}

func (o OffRampV1_0_0) EncodeExecutionReport(report ExecReport) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func NewOffRampV1_0_0(lggr logger.Logger, addr common.Address, ec client.Client, lp logpoller.LogPoller) (*OffRampV1_0_0, error) {
	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(addr, ec)
	if err != nil {
		return nil, err
	}
	var filters = []logpoller.Filter{
		{
			Name:      logpoller.FilterName(EXEC_EXECUTION_STATE_CHANGES, addr.String()),
			EventSigs: []common.Hash{abihelpers.EventSignatures.ExecutionStateChanged},
			Addresses: []common.Address{addr},
		},
		{
			Name:      logpoller.FilterName(EXEC_TOKEN_POOL_ADDED, addr.String()),
			EventSigs: []common.Hash{abihelpers.EventSignatures.PoolAdded},
			Addresses: []common.Address{addr},
		},
		{
			Name:      logpoller.FilterName(EXEC_TOKEN_POOL_REMOVED, addr.String()),
			EventSigs: []common.Hash{abihelpers.EventSignatures.PoolRemoved},
			Addresses: []common.Address{addr},
		},
	}
	if err := logpollerutil.RegisterLpFilters(lp, filters); err != nil {
		return nil, err
	}
	return &OffRampV1_0_0{offRamp: offRamp, addr: addr, lggr: lggr, lp: lp}, nil
}
