package ccip

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/smartcontractkit/sqlx"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/onramp"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins"
	ccipconfig "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/core/services/relay/types"
)

type CCIPRelay struct {
	db           *sqlx.DB
	lggr         logger.Logger
	ocr2Provider types.OCR2ProviderCtx
	cfg          Config
	configPoller *ConfigPoller

	jobID  int32
	spec   *job.OCR2OracleSpec
	config ccipconfig.RelayPluginConfig

	sourceChain evm.Chain
	destChain   evm.Chain
	offRamp     *offramp.OffRamp
	onRamp      *onramp.OnRamp
	afn         *afn_contract.AFNContract
}

var _ plugins.OraclePlugin = &CCIPRelay{}

func NewCCIPRelay(jobID int32, spec *job.OCR2OracleSpec, chainSet evm.ChainSet, db *sqlx.DB, ocr2Provider types.OCR2ProviderCtx, cfg Config, lggr logger.Logger) (*CCIPRelay, error) {
	var pluginConfig ccipconfig.RelayPluginConfig
	err := json.Unmarshal(spec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return &CCIPRelay{}, err
	}
	err = pluginConfig.ValidateRelayPluginConfig()
	if err != nil {
		return &CCIPRelay{}, err
	}
	lggr.Infof("CCIP relay plugin initialized with config: %+v", pluginConfig)

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
	offRamp, err := offramp.NewOffRamp(common.HexToAddress(spec.ContractID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed creating a new offramp")
	}
	onRamp, err := onramp.NewOnRamp(common.HexToAddress(string(pluginConfig.OnRampID)), sourceChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed creating a new onramp")
	}

	// Subscribe to all relevant relay logs.
	sourceChain.LogPoller().MergeFilter([]common.Hash{CrossChainSendRequested}, onRamp.Address())

	configPoller := NewConfigPoller(
		lggr.Named("CCIP_LogListener").With("jobID", jobID),
		destChain.LogPoller(),
		offRamp,
		pluginConfig.PollPeriod.Duration())
	return &CCIPRelay{
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

func (c *CCIPRelay) GetPluginFactory() (plugin ocrtypes.ReportingPluginFactory, err error) {
	return NewRelayReportingPluginFactory(c.lggr, c.sourceChain.LogPoller(), c.offRamp, c.onRamp, c.configPoller), nil
}

// GetServices returns the log listener service.
func (c *CCIPRelay) GetServices() ([]job.ServiceCtx, error) {
	return []job.ServiceCtx{c.configPoller}, nil
}
