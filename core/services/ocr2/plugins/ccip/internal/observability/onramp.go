package observability

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

type ObservedOnRampReader struct {
	ccipdata.OnRampReader
	metric metricDetails
}

func NewObservedOnRampReader(origin ccipdata.OnRampReader, chainID int64, pluginName string) *ObservedOnRampReader {
	return &ObservedOnRampReader{
		OnRampReader: origin,
		metric: metricDetails{
			histogram:  onRampHistogram,
			pluginName: pluginName,
			chainId:    chainID,
		},
	}
}

func (o ObservedOnRampReader) GetSendRequestsGteSeqNum(ctx context.Context, seqNum uint64, confs int) ([]ccipdata.Event[internal.EVM2EVMMessage], error) {
	return withObservedContract(o.metric, "GetSendRequestsGteSeqNum", func() ([]ccipdata.Event[internal.EVM2EVMMessage], error) {
		return o.OnRampReader.GetSendRequestsGteSeqNum(ctx, seqNum, confs)
	})
}

func (o ObservedOnRampReader) GetSendRequestsBetweenSeqNums(ctx context.Context, seqNumMin, seqNumMax uint64, confs int) ([]ccipdata.Event[internal.EVM2EVMMessage], error) {
	return withObservedContract(o.metric, "GetSendRequestsBetweenSeqNums", func() ([]ccipdata.Event[internal.EVM2EVMMessage], error) {
		return o.OnRampReader.GetSendRequestsBetweenSeqNums(ctx, seqNumMin, seqNumMax, confs)
	})
}

func (o ObservedOnRampReader) RouterAddress() (common.Address, error) {
	return withObservedContract(o.metric, "RouterAddress", func() (common.Address, error) {
		return o.OnRampReader.RouterAddress()
	})
}

func (o ObservedOnRampReader) GetOnRampAddress() (common.Address, error) {
	return withObservedContract(o.metric, "GetOnRampAddress", func() (common.Address, error) {
		return o.OnRampReader.GetOnRampAddress()
	})
}

func (o ObservedOnRampReader) GetOnRampDynamicConfig() (evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig, error) {
	return withObservedContract(o.metric, "GetOnRampDynamicConfig", func() (evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig, error) {
		return o.OnRampReader.GetOnRampDynamicConfig()
	})
}
