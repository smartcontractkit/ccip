package mocks

import (
	"context"
	"time"

	"github.com/smartcontractkit/ccipocr3/internal/model"
	"github.com/smartcontractkit/ccipocr3/internal/reader"
	"github.com/stretchr/testify/mock"
)

type CCIPReader struct {
	*mock.Mock
}

func NewCCIPReader() *CCIPReader {
	return &CCIPReader{
		Mock: &mock.Mock{},
	}
}

func (r CCIPReader) ReportsFromBlockNum(ctx context.Context, chain model.ChainSelector, blockNum uint64, limit int) (map[model.ChainSelector][]model.CommitReport, error) {
	args := r.Called(ctx, chain, blockNum, limit)
	return args.Get(0).(map[model.ChainSelector][]model.CommitReport), args.Error(1)
}

func (r CCIPReader) MsgsAfterTimestamp(ctx context.Context, chains []model.ChainSelector, ts time.Time, limit int) ([]model.CCIPMsg, error) {
	args := r.Called(ctx, chains, ts, limit)
	return args.Get(0).([]model.CCIPMsg), args.Error(1)
}

func (r CCIPReader) MsgsBetweenSeqNums(ctx context.Context, chains []model.ChainSelector, seqNumRange model.SeqNumRange) ([]model.CCIPMsg, error) {
	args := r.Called(ctx, chains, seqNumRange)
	return args.Get(0).([]model.CCIPMsg), args.Error(1)
}

func (r CCIPReader) NextSeqNum(ctx context.Context, chains []model.ChainSelector) (seqNum []model.SeqNum, err error) {
	args := r.Called(ctx, chains)
	return args.Get(0).([]model.SeqNum), args.Error(1)
}

func (r CCIPReader) GasPrices(ctx context.Context, chains []model.ChainSelector) ([]model.BigInt, error) {
	args := r.Called(ctx, chains)
	return args.Get(0).([]model.BigInt), args.Error(1)
}

func (r CCIPReader) Close(ctx context.Context) error {
	args := r.Called(ctx)
	return args.Error(0)
}

// Interface compatibility check.
var _ reader.CCIP = (*CCIPReader)(nil)
