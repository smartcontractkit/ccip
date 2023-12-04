package observability

import (
	"context"
	"time"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/offramp"
)

type ObservedCommitStoreReader struct {
	commit_store.CommitStoreReader
	metric metricDetails
}

func NewObservedCommitStoreReader(origin commit_store.CommitStoreReader, chainID int64, pluginName string) *ObservedCommitStoreReader {
	return &ObservedCommitStoreReader{
		CommitStoreReader: origin,
		metric: metricDetails{
			interactionDuration: readerHistogram,
			resultSetSize:       readerDatasetSize,
			pluginName:          pluginName,
			readerName:          "CommitStoreReader",
			chainId:             chainID,
		},
	}
}

func (o *ObservedCommitStoreReader) GetExpectedNextSequenceNumber(context context.Context) (uint64, error) {
	return withObservedInteraction(o.metric, "GetExpectedNextSequenceNumber", func() (uint64, error) {
		return o.CommitStoreReader.GetExpectedNextSequenceNumber(context)
	})
}

func (o *ObservedCommitStoreReader) GetLatestPriceEpochAndRound(context context.Context) (uint64, error) {
	return withObservedInteraction(o.metric, "GetLatestPriceEpochAndRound", func() (uint64, error) {
		return o.CommitStoreReader.GetLatestPriceEpochAndRound(context)
	})
}

func (o *ObservedCommitStoreReader) GetCommitReportMatchingSeqNum(ctx context.Context, seqNum uint64, confs int) ([]ccipdata.Event[commit_store.CommitStoreReport], error) {
	return withObservedInteractionAndResults(o.metric, "GetCommitReportMatchingSeqNum", func() ([]ccipdata.Event[commit_store.CommitStoreReport], error) {
		return o.CommitStoreReader.GetCommitReportMatchingSeqNum(ctx, seqNum, confs)
	})
}

func (o *ObservedCommitStoreReader) GetAcceptedCommitReportsGteTimestamp(ctx context.Context, ts time.Time, confs int) ([]ccipdata.Event[commit_store.CommitStoreReport], error) {
	return withObservedInteractionAndResults(o.metric, "GetAcceptedCommitReportsGteTimestamp", func() ([]ccipdata.Event[commit_store.CommitStoreReport], error) {
		return o.CommitStoreReader.GetAcceptedCommitReportsGteTimestamp(ctx, ts, confs)
	})
}

func (o *ObservedCommitStoreReader) IsDown(ctx context.Context) (bool, error) {
	return withObservedInteraction(o.metric, "IsDown", func() (bool, error) {
		return o.CommitStoreReader.IsDown(ctx)
	})
}

func (o *ObservedCommitStoreReader) IsBlessed(ctx context.Context, root [32]byte) (bool, error) {
	return withObservedInteraction(o.metric, "IsBlessed", func() (bool, error) {
		return o.CommitStoreReader.IsBlessed(ctx, root)
	})
}

func (o *ObservedCommitStoreReader) EncodeCommitReport(report commit_store.CommitStoreReport) ([]byte, error) {
	return withObservedInteraction(o.metric, "EncodeCommitReport", func() ([]byte, error) {
		return o.CommitStoreReader.EncodeCommitReport(report)
	})
}

func (o *ObservedCommitStoreReader) DecodeCommitReport(report []byte) (commit_store.CommitStoreReport, error) {
	return withObservedInteraction(o.metric, "DecodeCommitReport", func() (commit_store.CommitStoreReport, error) {
		return o.CommitStoreReader.DecodeCommitReport(report)
	})
}

func (o *ObservedCommitStoreReader) VerifyExecutionReport(ctx context.Context, report offramp.ExecReport) (bool, error) {
	return withObservedInteraction(o.metric, "VerifyExecutionReport", func() (bool, error) {
		return o.CommitStoreReader.VerifyExecutionReport(ctx, report)
	})
}

func (o *ObservedCommitStoreReader) GetCommitStoreStaticConfig(ctx context.Context) (commit_store.CommitStoreStaticConfig, error) {
	return withObservedInteraction(o.metric, "GetCommitStoreStaticConfig", func() (commit_store.CommitStoreStaticConfig, error) {
		return o.CommitStoreReader.GetCommitStoreStaticConfig(ctx)
	})
}
