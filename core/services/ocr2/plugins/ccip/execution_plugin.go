package ccip

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/smartcontractkit/sqlx"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/onramp"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/config"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins"
	ccipconfig "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/core/services/relay/types"
)

// Config contains OCR2 configurations for a job.
type Config interface {
	config.OCR2Config
	Dev() bool
	JobPipelineResultWriteQueueDepth() uint64
	LogSQL() bool
}

type CCIPExecution struct {
	db           *sqlx.DB
	lggr         logger.Logger
	ocr2Provider types.OCR2ProviderCtx
	cfg          Config

	jobID  int32
	spec   *job.OCR2OracleSpec
	config ccipconfig.ExecutionPluginConfig

	sourceChain  evm.Chain
	destChain    evm.Chain
	offRamp      *offramp.OffRamp
	onRamp       *onramp.OnRamp
	configPoller *ConfigPoller
}

var _ plugins.OraclePlugin = &CCIPExecution{}

func NewCCIPExecution(jobID int32, spec *job.OCR2OracleSpec, chainSet evm.ChainSet, db *sqlx.DB, ocr2Provider types.OCR2ProviderCtx, cfg Config, lggr logger.Logger) (*CCIPExecution, error) {
	var pluginConfig ccipconfig.ExecutionPluginConfig
	err := json.Unmarshal(spec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return &CCIPExecution{}, err
	}
	err = pluginConfig.ValidateExecutionPluginConfig()
	if err != nil {
		return &CCIPExecution{}, err
	}
	lggr.Infof("CCIP execution plugin initialized with config: %+v", pluginConfig)

	sourceChain, err := chainSet.Get(big.NewInt(pluginConfig.SourceChainID))
	if err != nil {
		return nil, errors.Wrap(err, "unable to open source chain")
	}
	destChain, err := chainSet.Get(big.NewInt(pluginConfig.DestChainID))
	if err != nil {
		return nil, errors.Wrap(err, "unable to open destination chain")
	}

	if !common.IsHexAddress(spec.ContractID) {
		return nil, errors.Wrap(err, "spec.OffRampID is not a valid hex address")
	}
	offRamp, err := offramp.NewOffRamp(common.HexToAddress(pluginConfig.OffRampId), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed creating a new onramp")
	}
	onRamp, err := onramp.NewOnRamp(common.HexToAddress(string(pluginConfig.OnRampID)), sourceChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed creating a new onramp")
	}

	// Subscribe to all relevant execution logs.
	sourceChain.LogPoller().MergeFilter([]common.Hash{CrossChainSendRequested}, onRamp.Address())
	destChain.LogPoller().MergeFilter([]common.Hash{ReportAccepted}, offRamp.Address())
	destChain.LogPoller().MergeFilter([]common.Hash{CrossChainMessageExecuted}, offRamp.Address())

	configPoller := NewConfigPoller(
		lggr.Named("CCIP_LogListener").With("jobID", jobID),
		destChain.LogPoller(),
		offRamp,
		pluginConfig.PollPeriod.Duration())
	return &CCIPExecution{
		db:           db,
		lggr:         lggr,
		ocr2Provider: ocr2Provider,
		cfg:          cfg,
		jobID:        jobID,
		spec:         spec,
		config:       pluginConfig,
		offRamp:      offRamp,
		onRamp:       onRamp,
		sourceChain:  sourceChain,
		destChain:    destChain,
		configPoller: configPoller,
	}, nil
}

func (c *CCIPExecution) GetPluginFactory() (plugin ocrtypes.ReportingPluginFactory, err error) {
	return NewExecutionReportingPluginFactory(
		c.lggr,
		c.onRamp,
		c.offRamp,
		c.sourceChain.LogPoller(),
		c.destChain.LogPoller(),
		common.HexToAddress(c.spec.ContractID),
		c.offRamp,
		c.configPoller,
	), nil
}

// GetServices returns the log listener service.
func (c *CCIPExecution) GetServices() ([]job.ServiceCtx, error) {
	return []job.ServiceCtx{c.configPoller}, nil
}
