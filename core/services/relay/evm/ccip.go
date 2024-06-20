package evm

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
)

// [IncompleteCommitStoreReader] is an implementation of CommitStoreReader with the only valid method being
// GasPriceEstimator.
type IncompleteCommitStoreReader struct {
	sourceChain legacyevm.Chain
}

func NewIncompleteCommitStoreReader() *IncompleteCommitStoreReader {
	return &IncompleteCommitStoreReader{}
}

func (i IncompleteCommitStoreReader) ChangeConfig(ctx context.Context, onchainConfig []byte, offchainConfig []byte) (cciptypes.Address, error) {
	return "", fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}

func (i IncompleteCommitStoreReader) DecodeCommitReport(ctx context.Context, report []byte) (cciptypes.CommitStoreReport, error) {
	return cciptypes.CommitStoreReport{}, fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}

func (i IncompleteCommitStoreReader) EncodeCommitReport(ctx context.Context, report cciptypes.CommitStoreReport) ([]byte, error) {
	return []byte{}, fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}

func (i IncompleteCommitStoreReader) GasPriceEstimator(ctx context.Context) (cciptypes.GasPriceEstimatorCommit, error) {
	estimator := i.sourceChain.GasEstimator()
	return estimator, nil
}

func (i IncompleteCommitStoreReader) GetAcceptedCommitReportsGteTimestamp(ctx context.Context, ts time.Time, confirmations int) ([]cciptypes.CommitStoreReportWithTxMeta, error) {
	return nil, fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}

func (i IncompleteCommitStoreReader) GetCommitReportMatchingSeqNum(ctx context.Context, seqNum uint64, confirmations int) ([]cciptypes.CommitStoreReportWithTxMeta, error) {
	return nil, fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}

func (i IncompleteCommitStoreReader) GetCommitStoreStaticConfig(ctx context.Context) (cciptypes.CommitStoreStaticConfig, error) {
	return cciptypes.CommitStoreStaticConfig{}, fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}

func (i IncompleteCommitStoreReader) GetExpectedNextSequenceNumber(ctx context.Context) (uint64, error) {
	return 0, fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}

func (i IncompleteCommitStoreReader) GetLatestPriceEpochAndRound(ctx context.Context) (uint64, error) {
	return 0, fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}

func (i IncompleteCommitStoreReader) IsBlessed(ctx context.Context, root [32]byte) (bool, error) {
	return false, fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}

func (i IncompleteCommitStoreReader) IsDestChainHealthy(ctx context.Context) (bool, error) {
	return false, fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}

func (i IncompleteCommitStoreReader) IsDown(ctx context.Context) (bool, error) {
	return false, fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}

func (i IncompleteCommitStoreReader) OffchainConfig(ctx context.Context) (cciptypes.CommitOffchainConfig, error) {
	return cciptypes.CommitOffchainConfig{}, fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}

func (i IncompleteCommitStoreReader) VerifyExecutionReport(ctx context.Context, report cciptypes.ExecReport) (bool, error) {
	return false, fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}

func (i IncompleteCommitStoreReader) Close() error {
	return fmt.Errorf("invalid usage of IncompleteCommitStoreReader")
}
