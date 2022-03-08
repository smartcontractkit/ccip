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
	"github.com/smartcontractkit/chainlink/core/services/pg"
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
	ccipORM      ORM
	ocr2Provider types.OCR2ProviderCtx
	cfg          Config

	jobID  int32
	spec   *job.OCR2OracleSpec
	config ccipconfig.ExecutionPluginConfig

	sourceChain evm.Chain
	destChain   evm.Chain
	offRamp     *offramp.OffRamp
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

	return &CCIPExecution{
		db:           db,
		lggr:         lggr,
		ocr2Provider: ocr2Provider,
		cfg:          cfg,
		jobID:        jobID,
		spec:         spec,
		ccipORM:      NewORM(db, lggr, cfg),
		config:       pluginConfig,

		offRamp:     offRamp,
		sourceChain: sourceChain,
		destChain:   destChain,
	}, nil
}

func (c *CCIPExecution) GetPluginFactory() (plugin ocrtypes.ReportingPluginFactory, err error) {
	return NewExecutionReportingPluginFactory(
		c.lggr,
		c.ccipORM,
		big.NewInt(c.config.SourceChainID),
		big.NewInt(c.config.DestChainID),
		common.HexToAddress(c.spec.ContractID),
		c.offRamp,
	), nil
}

// GetServices returns the log listener service.
func (c *CCIPExecution) GetServices() ([]job.ServiceCtx, error) {
	singleTokenOnRamp, err := onramp.NewOnRamp(common.HexToAddress(string(c.config.OnRampID)), c.sourceChain.Client())
	if err != nil {
		return nil, err
	}

	ccipConfig, err := GetOffchainConfig(c.ocr2Provider.ContractConfigTracker())
	if err != nil {
		return nil, errors.Wrap(err, "could not get the latest encoded config")
	}

	logListener := NewLogListener(c.lggr,
		c.sourceChain.LogBroadcaster(),
		c.destChain.LogBroadcaster(),
		singleTokenOnRamp,
		c.offRamp,
		ccipConfig,
		c.ccipORM,
		c.jobID,
		pg.NewQ(c.db, c.lggr, c.cfg))

	return []job.ServiceCtx{logListener}, nil
}
