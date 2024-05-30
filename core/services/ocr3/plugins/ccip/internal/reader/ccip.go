package reader

import "C"
import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/smartcontractkit/ccipocr3/internal/model"
)

var (
	ErrChainReaderNotFound = errors.New("chain reader not found")
)

type CCIP interface {
	// CommitReportsGTETimestamp reads the requested chain starting at a given timestamp
	// and finds all ReportAccepted up to the provided limit.
	CommitReportsGTETimestamp(ctx context.Context, dest model.ChainSelector, ts time.Time, limit int) ([]model.CommitPluginReportWithMeta, error)

	// MsgsAfterTimestamp reads the provided chains.
	// Finds and returns ccip messages submitted after the target time.
	// Messages are sorted ascending based on their timestamp and limited up to the provided limit.
	// TODO: unused.
	MsgsAfterTimestamp(ctx context.Context, chains []model.ChainSelector, ts time.Time, limit int) ([]model.CCIPMsg, error)

	// MsgsBetweenSeqNums reads the provided chains.
	// Finds and returns ccip messages submitted between the provided sequence numbers.
	// Messages are sorted ascending based on their timestamp and limited up to the provided limit.
	// TODO: a slice of chain selectors and a single seqNumRange doesn't make sense. Either have one
	//       chain to read or a slice of sequence number ranges.
	MsgsBetweenSeqNums(ctx context.Context, chains []model.ChainSelector, seqNumRange model.SeqNumRange) ([]model.CCIPMsg, error)

	// NextSeqNum reads the destination chain.
	// Returns the next expected sequence number for each one of the provided chains.
	// TODO: if destination was a parameter, this could be a capability reused across plugin instances.
	NextSeqNum(ctx context.Context, chains []model.ChainSelector) (seqNum []model.SeqNum, err error)

	// GasPrices reads the provided chains gas prices.
	GasPrices(ctx context.Context, chains []model.ChainSelector) ([]model.BigInt, error)

	// Close closes any open resources.
	Close(ctx context.Context) error
}

type ChainReader interface{} // TODO: Imported from chainlink-common

type CCIPChainReader struct {
	chainReaders map[model.ChainSelector]ChainReader
	destChain    model.ChainSelector
}

func (r *CCIPChainReader) CommitReportsGTETimestamp(ctx context.Context, dest model.ChainSelector, ts time.Time, limit int) ([]model.CommitPluginReport, error) {
	if err := r.validateReaderExistence(dest); err != nil {
		return nil, err
	}
	panic("implement me")
}

func (r *CCIPChainReader) MsgsAfterTimestamp(ctx context.Context, chains []model.ChainSelector, ts time.Time, limit int) ([]model.CCIPMsg, error) {
	if err := r.validateReaderExistence(chains...); err != nil {
		return nil, err
	}
	panic("implement me")
}

func (r *CCIPChainReader) MsgsBetweenSeqNums(ctx context.Context, chains []model.ChainSelector, seqNumRange model.SeqNumRange) ([]model.CCIPMsg, error) {
	if err := r.validateReaderExistence(chains...); err != nil {
		return nil, err
	}
	panic("implement me")
}

func (r *CCIPChainReader) NextSeqNum(ctx context.Context, chains []model.ChainSelector) (seqNum []model.SeqNum, err error) {
	if err := r.validateReaderExistence(r.destChain); err != nil {
		return nil, err
	}
	panic("implement me")
}

func (r *CCIPChainReader) GasPrices(ctx context.Context, chains []model.ChainSelector) ([]model.BigInt, error) {
	if err := r.validateReaderExistence(chains...); err != nil {
		return nil, err
	}
	panic("implement me")
}

func (r *CCIPChainReader) Close(ctx context.Context) error {
	return nil
}

func (r *CCIPChainReader) validateReaderExistence(chains ...model.ChainSelector) error {
	for _, ch := range chains {
		_, exists := r.chainReaders[ch]
		if !exists {
			return fmt.Errorf("chain %d: %w", ch, ErrChainReaderNotFound)
		}
	}
	return nil
}
