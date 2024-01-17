package factory

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_2_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

func NewOffRampReader(lggr logger.Logger, versionFinder VersionFinder, addr common.Address, destClient client.Client, lp logpoller.LogPoller, estimator gas.EvmFeeEstimator, registerFilters bool, pgOpts ...pg.QOpt) (ccipdata.OffRampReader, error) {
	return initOrCloseOffRampReader(lggr, versionFinder, addr, destClient, lp, estimator, false, registerFilters, pgOpts...)
}

func CloseOffRampReader(lggr logger.Logger, versionFinder VersionFinder, addr common.Address, destClient client.Client, lp logpoller.LogPoller, estimator gas.EvmFeeEstimator, pgOpts ...pg.QOpt) error {
	_, err := initOrCloseOffRampReader(lggr, versionFinder, addr, destClient, lp, estimator, true, false, pgOpts...)
	return err
}

func initOrCloseOffRampReader(lggr logger.Logger, versionFinder VersionFinder, addr common.Address, destClient client.Client, lp logpoller.LogPoller, estimator gas.EvmFeeEstimator, closeReader bool, registerFilters bool, pgOpts ...pg.QOpt) (ccipdata.OffRampReader, error) {
	contractType, version, err := versionFinder.TypeAndVersion(addr, destClient)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read type and version")
	}
	if contractType != ccipconfig.EVM2EVMOffRamp {
		return nil, errors.Errorf("expected %v got %v", ccipconfig.EVM2EVMOffRamp, contractType)
	}
	switch version.String() {
	case ccipdata.V1_0_0, ccipdata.V1_1_0:
		offRamp, err := v1_0_0.NewOffRamp(lggr, addr, destClient, lp, estimator)
		if err != nil {
			return nil, err
		}
		if closeReader {
			return nil, offRamp.Close(pgOpts...)
		}
		return offRamp, offRamp.RegisterFilters(pgOpts...)
	case ccipdata.V1_2_0, ccipdata.V1_3_0:
		offRamp, err := v1_2_0.NewOffRamp(lggr, addr, destClient, lp, estimator)
		if err != nil {
			return nil, err
		}
		if closeReader {
			return nil, offRamp.Close(pgOpts...)
		}
		return offRamp, offRamp.RegisterFilters(pgOpts...)
	default:
		return nil, errors.Errorf("unsupported offramp version %v", version.String())
	}
	// TODO can validate it pointing to the correct version
}

func ExecReportToEthTxMeta(typ ccipconfig.ContractType, ver semver.Version) (func(report []byte) (*txmgr.TxMeta, error), error) {
	if typ != ccipconfig.EVM2EVMOffRamp {
		return nil, errors.Errorf("expected %v got %v", ccipconfig.EVM2EVMOffRamp, typ)
	}
	switch ver.String() {
	case ccipdata.V1_0_0, ccipdata.V1_1_0:
		offRampABI := abihelpers.MustParseABI(evm_2_evm_offramp_1_0_0.EVM2EVMOffRampABI)
		return func(report []byte) (*txmgr.TxMeta, error) {
			execReport, err := v1_0_0.DecodeExecReport(abihelpers.MustGetMethodInputs(ccipdata.ManuallyExecute, offRampABI)[:1], report)
			if err != nil {
				return nil, err
			}
			return execReportToEthTxMeta(execReport)
		}, nil
	case ccipdata.V1_2_0, ccipdata.V1_3_0:
		offRampABI := abihelpers.MustParseABI(evm_2_evm_offramp.EVM2EVMOffRampABI)
		return func(report []byte) (*txmgr.TxMeta, error) {
			execReport, err := v1_2_0.DecodeExecReport(abihelpers.MustGetMethodInputs(ccipdata.ManuallyExecute, offRampABI)[:1], report)
			if err != nil {
				return nil, err
			}
			return execReportToEthTxMeta(execReport)
		}, nil
	default:
		return nil, errors.Errorf("got unexpected version %v", ver.String())
	}
}

func execReportToEthTxMeta(execReport ccipdata.ExecReport) (*txmgr.TxMeta, error) {
	msgIDs := make([]string, len(execReport.Messages))
	for i, msg := range execReport.Messages {
		msgIDs[i] = hexutil.Encode(msg.MessageId[:])
	}

	return &txmgr.TxMeta{
		MessageIDs: msgIDs,
	}, nil
}
