package ccipexec

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/smartcontractkit/chainlink/v2/plugins"

	"github.com/smartcontractkit/chainlink-common/pkg/loop"

	"github.com/smartcontractkit/chainlink/v2/core/config/env"

	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/Masterminds/semver/v3"
	"go.uber.org/multierr"

	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2plus"

	commonlogger "github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/factory"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/oraclelib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/promwrapper"
)

var (
	// tokenDataWorkerTimeout defines 1) The timeout while waiting for a bg call to the token data 3P provider.
	// 2) When a client requests token data and does not specify a timeout this value is used as a default.
	// 5 seconds is a reasonable value for a timeout.
	// At this moment, minimum OCR Delta Round is set to 30s and deltaGrace to 5s. Based on this configuration
	// 5s for token data worker timeout is a reasonable default.
	tokenDataWorkerTimeout = 5 * time.Second
	// tokenDataWorkerNumWorkers is the number of workers that will be processing token data in parallel.
	tokenDataWorkerNumWorkers = 5
)

var defaultNewReportingPluginRetryConfig = ccipdata.RetryConfig{
	InitialDelay: time.Second,
	MaxDelay:     10 * time.Minute,
	// Retry for approximately 4hrs (MaxDelay of 10m = 6 times per hour, times 4 hours, plus 10 because the first
	// 10 retries only take 20 minutes due to an initial retry of 1s and exponential backoff)
	MaxRetries: (6 * 4) + 10,
}

func NewExecServices(ctx context.Context, lggr logger.Logger, cfg plugins.RegistrarConfig, jb job.Job, sourceTokenAddress string, srcProvider types.CCIPExecProvider, dstProvider types.CCIPExecProvider, srcChainID int64, dstChainID int64, new bool, argsNoPlugin libocr2.OCR2OracleArgs, logError func(string)) ([]job.ServiceCtx, error) {
	lggr = lggr.Named("ccip-exec").Named(string(jb.ID))

	loopCmd := env.CCIPExecPlugin.Cmd.Get()
	loopEnabled := loopCmd != ""

	var pluginFactory types.ReportingPluginFactory
	if loopEnabled {
		// find loop command
		envVars, err := plugins.ParseEnvFile(env.CCIPExecPlugin.Env.Get())
		if err != nil {
			return nil, fmt.Errorf("failed to parse ccip exec env file: %w", err)
		}
		cmdFn, grpcOpts, err := cfg.RegisterLOOP(plugins.CmdConfig{
			ID:  lggr.Name(),
			Cmd: loopCmd,
			Env: envVars,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to register ccip exec plugin: %w", err)
		}
		// get reporting plugin factory from loop
		factoryServer := loop.NewExecutionService(lggr, grpcOpts, cmdFn, srcProvider, dstProvider, uint32(srcChainID), uint32(dstChainID), sourceTokenAddress)
		pluginFactory = factoryServer
	} else {
		var err2 error
		pluginFactory, err2 = NewExecutionReportingPluginFactoryV2(ctx, lggr, sourceTokenAddress, srcChainID, dstChainID, srcProvider, dstProvider)
		if err2 != nil {
			return nil, err2
		}
	}

	argsNoPlugin.ReportingPluginFactory = promwrapper.NewPromFactory(pluginFactory, "CCIPExecution", jb.OCR2OracleSpec.Relay, big.NewInt(0).SetInt64(dstChainID))
	argsNoPlugin.Logger = commonlogger.NewOCRWrapper(lggr, true, logError)
	oracle, err := libocr2.NewOracle(argsNoPlugin)
	if err != nil {
		return nil, err
	}
	// If this is a brand-new job, then we make use of the start blocks. If not then we're rebooting and log poller will pick up where we left off.
	if new {
		return []job.ServiceCtx{
			oraclelib.NewChainAgnosticBackFilledOracle(
				lggr,
				srcProvider,
				dstProvider,
				job.NewServiceAdapter(oracle),
			),
		}, nil
	}
	return []job.ServiceCtx{
		job.NewServiceAdapter(oracle),
	}, nil
}

// UnregisterExecPluginLpFilters unregisters all the registered filters for both source and dest chains.
// See comment in UnregisterCommitPluginLpFilters
// It MUST mirror the filters registered in NewExecServices.
// This currently works because the filters registered by the created custom providers when the job is first added
// are stored in the db. Those same filters are unregistered (i.e. deleted from the db) by the newly created providers
// that are passed in from cleanupEVM, as while the providers have no knowledge of each other, they are created
// on the same source and dest relayer.
func UnregisterExecPluginLpFilters(srcProvider types.CCIPExecProvider, dstProvider types.CCIPExecProvider) error {
	unregisterFuncs := []func() error{
		func() error {
			return srcProvider.Close()
		},
		func() error {
			return dstProvider.Close()
		},
	}

	var multiErr error
	for _, fn := range unregisterFuncs {
		if err := fn(); err != nil {
			multiErr = multierr.Append(multiErr, err)
		}
	}
	return multiErr
}

// ExecReportToEthTxMeta generates a txmgr.EthTxMeta from the given report.
// Only MessageIDs will be populated in the TxMeta.
func ExecReportToEthTxMeta(ctx context.Context, typ ccipconfig.ContractType, ver semver.Version) (func(report []byte) (*txmgr.TxMeta, error), error) {
	return factory.ExecReportToEthTxMeta(ctx, typ, ver)
}
