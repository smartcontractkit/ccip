package tokenprice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	libocr "github.com/smartcontractkit/libocr/offchainreporting2plus"

	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink/v2/core/config/env"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/tokenprice/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/v2/plugins"
)

type TokenPriceConfig interface {
	plugins.RegistrarConfig
}

// concrete implementation of TokenPriceConfig
type tokenPriceConfig struct {
	plugins.RegistrarConfig
}

func NewTokenPriceConfig(pluginProcessCfg plugins.RegistrarConfig) TokenPriceConfig {
	return &tokenPriceConfig{
		RegistrarConfig: pluginProcessCfg,
	}
}

// This wrapper avoids the need to modify the signature of NewTokenPriceFactory in all of the non-evm
// relay repos as well as its primary definition in chainlink-common. Once ChainReader is implemented
// and working on all 4 blockchain families, we can remove the original TokenPriceContract() method from
// TokenPriceProvider and pass TokenPriceContract as a separate param to NewTokenPriceFactory
type tokenPriceProviderWrapper struct {
	types.TokenPriceProvider
	contract TokenPriceContract
}

// Override relay's implementation of TokenPriceContract with product plugin's implementation of
// TokenPriceContract, making use of product-agnostic ChainReader to read the contract instead of relay TokenPriceContract
func (m tokenPriceProviderWrapper) TokenPriceContract() TokenPriceContract {
	return m.contract
}

func NewTokenPriceServices(
	ctx context.Context,
	jb job.Job,
	isNewlyCreatedJob bool,
	relayer loop.Relayer,
	pipelineRunner pipeline.Runner,
	lggr logger.Logger,
	argsNoPlugin libocr.OCR3OracleArgs[meta],
	cfg TokenPriceConfig,
	chEnhancedTelem chan ocrcommon.EnhancedTelemetryData,
	errorLog loop.ErrorLog,
) (srvs []job.ServiceCtx, err error) {
	var pluginConfig config.PluginConfig
	err = json.Unmarshal(jb.OCR2OracleSpec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return
	}
	err = config.ValidatePluginConfig(pluginConfig)
	if err != nil {
		return
	}
	spec := jb.OCR2OracleSpec

	provider, err := relayer.NewPluginProvider(ctx, types.RelayArgs{
		ExternalJobID: jb.ExternalJobID,
		JobID:         jb.ID,
		ContractID:    spec.ContractID,
		New:           isNewlyCreatedJob,
		RelayConfig:   spec.RelayConfig.Bytes(),
		ProviderType:  string(spec.PluginType),
	}, types.PluginArgs{
		TransmitterID: spec.TransmitterID.String,
		PluginConfig:  spec.PluginConfig.Bytes(),
	})
	if err != nil {
		return
	}

	tokenPriceProvider, ok := provider.(types.TokenPriceProvider)
	if !ok {
		return nil, errors.New("could not coerce PluginProvider to TokenPriceProvider")
	}

	srvs = append(srvs, provider)
	argsNoPlugin.ContractTransmitter = provider.ContractTransmitter()
	argsNoPlugin.ContractConfigTracker = provider.ContractConfigTracker()
	argsNoPlugin.OffchainConfigDigester = provider.OffchainConfigDigester()

	abort := func() {
		if cerr := services.MultiCloser(srvs).Close(); err != nil {
			lggr.Errorw("Error closing unused services", "err", cerr)
		}
	}

	tokenPricePluginCmd := env.TokenPricePluginCmd.Get()
	tokenPriceLoopEnabled := tokenPricePluginCmd != ""

	// TODO BCF-2821 handle this properly as this blocks Solana chain reader dev
	if !tokenPriceLoopEnabled && tokenPriceProvider.ChainReader() != nil {
		lggr.Info("Chain Reader enabled")
		tokenPriceProvider = tokenPriceProviderWrapper{
			tokenPriceProvider, // attach newer MedianContract which uses ChainReader
			newTokenPriceContract(provider.ChainReader(), common.HexToAddress(spec.ContractID)),
		}
	} else {
		lggr.Info("Chain Reader disabled")
	}

	if tokenPriceLoopEnabled {
		// use unique logger names so we can use it to register a loop
		medianLggr := lggr.Named("TokenPrice").Named(spec.ContractID).Named(spec.GetID())
		cmdFn, telem, err2 := cfg.RegisterLOOP(medianLggr.Name(), tokenPricePluginCmd)
		if err2 != nil {
			err = fmt.Errorf("failed to register loop: %w", err2)
			abort()
			return
		}
		median := loop.NewTokenPriceService(lggr, telem, cmdFn, tokenPriceProvider, errorLog)
		argsNoPlugin.ReportingPluginFactory = median
		srvs = append(srvs, median)
	} else {
		argsNoPlugin.ReportingPluginFactory, err = NewPlugin(lggr).NewTokenPriceFactory(ctx, tokenPriceProvider, errorLog)
		if err != nil {
			err = fmt.Errorf("failed to create median factory: %w", err)
			abort()
			return
		}
	}

	var oracle libocr.Oracle
	oracle, err = libocr.NewOracle(argsNoPlugin)
	if err != nil {
		abort()
		return
	}
	srvs = append(srvs, runSaver, job.NewServiceAdapter(oracle))
	if !jb.OCR2OracleSpec.CaptureEATelemetry {
		lggr.Infof("Enhanced EA telemetry is disabled for job %s", jb.Name.ValueOrZero())
	}
	return
}
