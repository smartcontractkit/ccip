package commitplugin

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/cache"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

type Factory struct {
	// Configuration derived from the job spec which does not change
	// between plugin instances (ie between SetConfigs onchain)
	config StaticConfig

	// Dynamic readers
	readersMu          *sync.Mutex
	destPriceRegReader ccipdata.PriceRegistryReader
	destPriceRegAddr   common.Address
}

// NewReportingPlugin returns the ccip CommitReportingPlugin and satisfies the ReportingPluginFactory interface.
func (rf *Factory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	destPriceReg, err := rf.config.commitStore.ChangeConfig(config.OnchainConfig, config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	if err = rf.updateDynamicReaders(destPriceReg); err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}

	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}

	return &CommitReportingPlugin{
			sourceChainSelector:     rf.config.sourceChainSelector,
			sourceNative:            rf.config.sourceNative,
			onRampReader:            rf.config.onRampReader,
			commitStoreReader:       rf.config.commitStore,
			priceGetter:             rf.config.priceGetter,
			F:                       config.F,
			lggr:                    rf.config.lggr.Named("CommitReportingPlugin"),
			inflightReports:         newInflightCommitReportsContainer(rf.config.commitStore.OffchainConfig().InflightCacheExpiry),
			destPriceRegistryReader: rf.destPriceRegReader,
			tokenDecimalsCache: cache.NewTokenToDecimals(
				rf.config.lggr,
				rf.config.destLP,
				rf.config.offRamp,
				rf.destPriceRegReader,
				rf.config.destClient,
				int64(rf.config.commitStore.OffchainConfig().DestFinalityDepth),
			),
			gasPriceEstimator: rf.config.commitStore.GasPriceEstimator(),
		},
		types.ReportingPluginInfo{
			Name:          "CCIPCommit",
			UniqueReports: false, // See comment in CommitStore constructor.
			Limits: types.ReportingPluginLimits{
				MaxQueryLength:       0,
				MaxObservationLength: ccip.MaxObservationLength,
				MaxReportLength:      MaxCommitReportLength,
			},
		}, nil
}

// newFactory return a new CommitReportingPluginFactory.
func newFactory(config StaticConfig) *Factory {
	return &Factory{
		config:    config,
		readersMu: &sync.Mutex{},
	}
}

func (rf *Factory) updateDynamicReaders(newPriceRegAddr common.Address) error {
	rf.readersMu.Lock()
	defer rf.readersMu.Unlock()
	// TODO: Investigate use of Close() to cleanup.
	// TODO: a true price registry upgrade on an existing lane may want some kind of start block in its config? Right now we
	// essentially assume that plugins don't care about historical price reg logs.
	if rf.destPriceRegAddr == newPriceRegAddr {
		// No-op
		return nil
	}
	// Close old reader if present and open new reader if address changed
	if rf.destPriceRegReader != nil {
		if err := rf.destPriceRegReader.Close(); err != nil {
			return err
		}
	}
	destPriceRegistryReader, err := ccipdata.NewPriceRegistryReader(rf.config.lggr, newPriceRegAddr, rf.config.destLP, rf.config.destClient)
	if err != nil {
		return err
	}
	rf.destPriceRegReader = destPriceRegistryReader
	rf.destPriceRegAddr = newPriceRegAddr
	return nil
}
