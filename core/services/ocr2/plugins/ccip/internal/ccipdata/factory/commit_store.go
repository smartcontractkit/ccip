package factory

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_2_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

func NewCommitStoreReader(lggr logger.Logger, versionFinder VersionFinder, address common.Address, ec client.Client, lp logpoller.LogPoller, estimator gas.EvmFeeEstimator, pgOpts ...pg.QOpt) (ccipdata.CommitStoreReader, error) {
	return initOrCloseCommitStoreReader(lggr, versionFinder, address, ec, lp, estimator, false, pgOpts...)
}

func CloseCommitStoreReader(lggr logger.Logger, versionFinder VersionFinder, address common.Address, ec client.Client, lp logpoller.LogPoller, estimator gas.EvmFeeEstimator, pgOpts ...pg.QOpt) error {
	_, err := initOrCloseCommitStoreReader(lggr, versionFinder, address, ec, lp, estimator, true, pgOpts...)
	return err
}

func initOrCloseCommitStoreReader(lggr logger.Logger, versionFinder VersionFinder, address common.Address, ec client.Client, lp logpoller.LogPoller, estimator gas.EvmFeeEstimator, closeReader bool, pgOpts ...pg.QOpt) (ccipdata.CommitStoreReader, error) {
	contractType, version, err := versionFinder.TypeAndVersion(address, ec)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read type and version")
	}
	if contractType != ccipconfig.CommitStore {
		return nil, errors.Errorf("expected %v got %v", ccipconfig.CommitStore, contractType)
	}
	switch version.String() {
	case ccipdata.V1_0_0, ccipdata.V1_1_0: // Versions are identical
		cs, err := v1_0_0.NewCommitStore(lggr, address, ec, lp, estimator)
		if err != nil {
			return nil, err
		}
		if closeReader {
			return nil, cs.Close(pgOpts...)
		}
		return cs, cs.RegisterFilters(pgOpts...)
	case ccipdata.V1_2_0, ccipdata.V1_3_0:
		cs, err := v1_2_0.NewCommitStore(lggr, address, ec, lp, estimator)
		if err != nil {
			return nil, err
		}
		if closeReader {
			return nil, cs.Close(pgOpts...)
		}
		return cs, cs.RegisterFilters(pgOpts...)
	default:
		return nil, errors.Errorf("unsupported commit store version %v", version.String())
	}
}

func CommitReportToEthTxMeta(typ ccipconfig.ContractType, ver semver.Version) (func(report []byte) (*txmgr.TxMeta, error), error) {
	if typ != ccipconfig.CommitStore {
		return nil, errors.Errorf("expected %v got %v", ccipconfig.CommitStore, typ)
	}
	switch ver.String() {
	case ccipdata.V1_0_0, ccipdata.V1_1_0:
		commitStoreABI := abihelpers.MustParseABI(commit_store_1_0_0.CommitStoreABI)
		return func(report []byte) (*txmgr.TxMeta, error) {
			commitReport, err := v1_0_0.DecodeCommitReport(abihelpers.MustGetEventInputs(v1_0_0.ReportAccepted, commitStoreABI), report)
			if err != nil {
				return nil, err
			}
			return commitReportToEthTxMeta(commitReport)
		}, nil
	case ccipdata.V1_2_0, ccipdata.V1_3_0:
		commitStoreABI := abihelpers.MustParseABI(commit_store.CommitStoreABI)
		return func(report []byte) (*txmgr.TxMeta, error) {
			commitReport, err := v1_2_0.DecodeCommitReport(abihelpers.MustGetEventInputs(v1_0_0.ReportAccepted, commitStoreABI), report)
			if err != nil {
				return nil, err
			}
			return commitReportToEthTxMeta(commitReport)
		}, nil
	default:
		return nil, errors.Errorf("got unexpected version %v", ver.String())
	}
}

// CommitReportToEthTxMeta generates a txmgr.EthTxMeta from the given commit report.
// sequence numbers of the committed messages will be added to tx metadata
func commitReportToEthTxMeta(commitReport ccipdata.CommitStoreReport) (*txmgr.TxMeta, error) {
	n := (commitReport.Interval.Max - commitReport.Interval.Min) + 1
	seqRange := make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		seqRange[i] = i + commitReport.Interval.Min
	}
	return &txmgr.TxMeta{
		SeqNumbers: seqRange,
	}, nil
}
