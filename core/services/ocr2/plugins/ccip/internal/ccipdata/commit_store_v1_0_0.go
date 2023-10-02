package ccipdata

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/net/context"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/logpollerutil"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

const (
	EXEC_REPORT_ACCEPTS = "Exec report accepts"
)

var _ CommitStoreReader = &CommitStoreV1_0_0{}

type CommitStoreV1_0_0 struct {
	commit_store *commit_store.CommitStore
	lggr         logger.Logger
	lp           logpoller.LogPoller
	address      common.Address
}

func (c *CommitStoreV1_0_0) Close(qopts ...pg.QOpt) error {
	//TODO implement me
	panic("implement me")
}

func (c *CommitStoreV1_0_0) parseReport(log types.Log) (*CommitStoreReport, error) {
	repAccepted, err := c.commit_store.ParseReportAccepted(log)
	if err != nil {
		return nil, err
	}
	// Translate to common struct.
	var tokenPrices []TokenPrice
	for _, tpu := range repAccepted.Report.PriceUpdates.TokenPriceUpdates {
		tokenPrices = append(tokenPrices, TokenPrice{
			Token: tpu.SourceToken,
			Value: tpu.UsdPerToken,
		})
	}
	return &CommitStoreReport{
		TokenPrices: tokenPrices,
		GasPrices:   []GasPrice{{DestChain: repAccepted.Report.PriceUpdates.DestChainSelector, Value: repAccepted.Report.PriceUpdates.UsdPerUnitGas}},
		MerkleRoot:  repAccepted.Report.MerkleRoot,
		Interval:    Interval{Min: repAccepted.Report.Interval.Min, Max: repAccepted.Report.Interval.Max},
	}, nil
}

func (c *CommitStoreV1_0_0) GetAcceptedCommitReportsGteSeqNum(ctx context.Context, seqNum uint64, confs int) ([]Event[CommitStoreReport], error) {
	logs, err := c.lp.LogsDataWordGreaterThan(
		abihelpers.EventSignatures.ReportAccepted,
		c.address,
		abihelpers.EventSignatures.ReportAcceptedMaxSequenceNumberWord,
		logpoller.EvmWord(seqNum),
		confs,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, err
	}

	return parseLogs[CommitStoreReport](
		logs,
		c.lggr,
		c.parseReport,
	)
}

func (c *CommitStoreV1_0_0) GetAcceptedCommitReportsGteTimestamp(ctx context.Context, ts time.Time, confs int) ([]Event[CommitStoreReport], error) {
	logs, err := c.lp.LogsCreatedAfter(
		abihelpers.EventSignatures.ReportAccepted,
		c.address,
		ts,
		confs,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, err
	}

	return parseLogs[CommitStoreReport](
		logs,
		c.lggr,
		c.parseReport,
	)
}

func (c *CommitStoreV1_0_0) GetExpectedNextSequenceNumber(ctx context.Context) (uint64, error) {
	return c.commit_store.GetExpectedNextSequenceNumber(&bind.CallOpts{Context: ctx})
}

func (c *CommitStoreV1_0_0) GetLatestPriceEpochAndRound(ctx context.Context) (uint64, error) {
	return c.commit_store.GetLatestPriceEpochAndRound(&bind.CallOpts{Context: ctx})
}

func (c *CommitStoreV1_0_0) IsDown(ctx context.Context) bool {
	unPausedAndHealthy, err := c.commit_store.IsUnpausedAndARMHealthy(&bind.CallOpts{Context: ctx})
	if err != nil {
		// If we cannot read the state, assume the worst
		c.lggr.Errorw("Unable to read CommitStore IsUnpausedAndARMHealthy", "err", err)
		return true
	}
	return !unPausedAndHealthy
}

func (c *CommitStoreV1_0_0) Verify(ctx context.Context, report ExecReport) bool {
	var hashes [][32]byte
	for _, msg := range report.Messages {
		hashes = append(hashes, msg.Hash)
	}
	res, err := c.commit_store.Verify(&bind.CallOpts{Context: ctx}, hashes, report.Proofs, report.ProofFlagBits)
	if err != nil {
		c.lggr.Errorw("Unable to call verify", "messages", report.Messages, "err", err)
		return false
	}
	// No timestamp, means failed to verify root.
	if res.Cmp(big.NewInt(0)) == 0 {
		c.lggr.Errorw("Root does not verify", "messages", report.Messages)
		return false
	}
	return true
}

func NewCommitStoreV1_0_0(lggr logger.Logger, addr common.Address, ec client.Client, lp logpoller.LogPoller) (*CommitStoreV1_0_0, error) {
	commit_store, err := commit_store.NewCommitStore(addr, ec)
	if err != nil {
		return nil, err
	}
	var filters = []logpoller.Filter{
		{
			Name:      logpoller.FilterName(EXEC_REPORT_ACCEPTS, addr.String()),
			EventSigs: []common.Hash{abihelpers.EventSignatures.ReportAccepted},
			Addresses: []common.Address{addr},
		},
	}
	if err := logpollerutil.RegisterLpFilters(lp, filters); err != nil {
		return nil, err
	}
	return &CommitStoreV1_0_0{commit_store: commit_store, lggr: lggr, lp: lp}, nil
}
