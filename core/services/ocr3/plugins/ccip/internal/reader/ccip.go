package reader

import (
	"context"
	"time"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/ccip/internal/model"
)

type CCIP interface {
	// MsgsAfterTimestamp returns ccip messages submitted in the provided chains after the target timestamp.
	// Messages are sorted ascending based on their timestamp and limited up to the provided limit.
	MsgsAfterTimestamp(ctx context.Context, chains []model.ChainSelector, ts time.Time, limit int) ([]model.CCIPMsg, error)

	// MsgsBetweenSeqNums returns ccip messages submitted in the provided chains with sequence number between
	// the given range (inclusive).
	MsgsBetweenSeqNums(ctx context.Context, chains []model.ChainSelector, seqNumRange model.SeqNumRange) ([]model.CCIPMsg, error)

	// NextSeqNum returns the next expected message sequence number.
	NextSeqNum(ctx context.Context) (seqNum model.SeqNum, err error)
}
