package commit

import (
	"context"
	"math/big"

	"github.com/smartcontractkit/ccipocr3/internal/reader"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"google.golang.org/grpc"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

// PluginFactoryConstructor implements common OCR3ReportingPluginClient and is used for initializing a plugin factory
// and a validation service.
type PluginFactoryConstructor struct{}

func NewPluginFactoryConstructor() *PluginFactoryConstructor {
	return &PluginFactoryConstructor{}
}
func (p PluginFactoryConstructor) NewReportingPluginFactory(
	ctx context.Context,
	config core.ReportingPluginServiceConfig,
	grpcProvider grpc.ClientConnInterface,
	pipelineRunner core.PipelineRunnerService,
	telemetry core.TelemetryService,
	errorLog core.ErrorLog,
	capRegistry core.CapabilitiesRegistry,
	keyValueStore core.KeyValueStore,
	relayerSet core.RelayerSet,
) (core.OCR3ReportingPluginFactory, error) {
	return nil, nil
}

func (p PluginFactoryConstructor) NewValidationService(ctx context.Context) (core.ValidationService, error) {
	panic("implement me")
}

// PluginFactory implements common ReportingPluginFactory and is used for (re-)initializing commit plugin instances.
type PluginFactory struct {
	contractReaders map[cciptypes.ChainSelector]types.ContractReader
	contractWriters map[cciptypes.ChainSelector]types.ChainWriter
	destChain       cciptypes.ChainSelector
	destChainCodec  cciptypes.CommitPluginCodec
	destChainHasher cciptypes.MessageHasher
	lggr            logger.Logger
}

func NewPluginFactory(
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
	contractWriters map[cciptypes.ChainSelector]types.ChainWriter,
	destChain cciptypes.ChainSelector,
	destChainCodec cciptypes.CommitPluginCodec,
	destChainHasher cciptypes.MessageHasher,
	lggr logger.Logger,
) *PluginFactory {
	return &PluginFactory{
		contractReaders: contractReaders,
		contractWriters: contractWriters,
		destChain:       destChain,
		destChainCodec:  destChainCodec,
		destChainHasher: destChainHasher,
		lggr:            lggr,
	}
}

func (p PluginFactory) NewReportingPlugin(config ocr3types.ReportingPluginConfig,
) (ocr3types.ReportingPlugin[[]byte], ocr3types.ReportingPluginInfo, error) {
	onChainTokenPricesReader := reader.NewOnchainTokenPricesReader(
		reader.TokenPriceConfig{ // TODO: Inject config
			StaticPrices: map[ocr2types.Account]big.Int{},
		},
		nil, // TODO: Inject this
	)

	return NewPlugin(
			context.Background(),
			config.OracleID,
			cciptypes.CommitPluginConfig{},
			reader.NewCCIPChainReader(
				p.contractReaders,
				p.contractWriters,
				p.destChain,
			),
			onChainTokenPricesReader,
			p.destChainCodec,
			p.destChainHasher,
			p.lggr,
		), ocr3types.ReportingPluginInfo{
			Name: "CCIPCommitOCR3",
			Limits: ocr3types.ReportingPluginLimits{
				MaxQueryLength:       0, // no query for commit (yet)
				MaxObservationLength: 2048,
				MaxOutcomeLength:     2048,
				MaxReportLength:      2048,
				MaxReportCount:       50,
			},
		}, nil
}

func (p PluginFactory) Name() string {
	panic("implement me")
}

func (p PluginFactory) Start(ctx context.Context) error {
	panic("implement me")
}

func (p PluginFactory) Close() error {
	panic("implement me")
}

func (p PluginFactory) Ready() error {
	panic("implement me")
}

func (p PluginFactory) HealthReport() map[string]error {
	panic("implement me")
}

// Interface compatibility checks.
var _ core.OCR3ReportingPluginClient = &PluginFactoryConstructor{}
var _ core.OCR3ReportingPluginFactory = &PluginFactory{}
