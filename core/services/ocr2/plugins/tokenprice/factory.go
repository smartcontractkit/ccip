package tokenprice

import (
	"context"

	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type Plugin struct {
	loop.Plugin
	stop services.StopChan
}

func NewPlugin(lggr logger.Logger) *Plugin {
	return &Plugin{Plugin: loop.Plugin{Logger: lggr}, stop: make(services.StopChan)}
}

func (p *Plugin) NewTokenPriceFactory(ctx context.Context, provider types.TokenPriceProvider, errorLog loop.ErrorLog) (loop.ReportingPluginFactory, error) {
	var ctxVals loop.ContextValues
	ctxVals.SetValues(ctx)
	lggr := logger.With(p.Logger, ctxVals.Args()...)
	factory := &TokenPriceFactory{
		Logger: logger.NewOCRWrapper(lggr, true, func(msg string) {
			ctx, cancelFn := p.stop.NewCtx()
			defer cancelFn()
			if err := errorLog.SaveError(ctx, msg); err != nil {
				lggr.Errorw("Unable to save error", "err", msg)
			}
		}),
	}
	if cr := provider.ChainReader(); cr != nil {
		factory.ContractTransmitter = &chainReaderContract{cr, types.BoundContract{Name: "tokenprice"}}
	} else {
		factory.ContractTransmitter = provider.TokenPriceContract()
	}
	s := &reportingPluginFactoryService{lggr: logger.Named(lggr, "ReportingPluginFactory"), ReportingPluginFactory: factory}

	p.SubService(s)

	return s, nil
}

type reportingPluginFactoryService struct {
	services.StateMachine
	lggr logger.Logger
	ocrtypes.ReportingPluginFactory
}

func (r *reportingPluginFactoryService) Name() string { return r.lggr.Name() }

func (r *reportingPluginFactoryService) Start(ctx context.Context) error {
	return r.StartOnce("ReportingPluginFactory", func() error { return nil })
}

func (r *reportingPluginFactoryService) Close() error {
	return r.StopOnce("ReportingPluginFactory", func() error { return nil })
}

func (r *reportingPluginFactoryService) HealthReport() map[string]error {
	return map[string]error{r.Name(): r.Healthy()}
}

type chainReaderContract struct {
	types.ChainReader
	types.BoundContract
}

// FromAccount implements types.ContractTransmitter.
func (*chainReaderContract) FromAccount() (ocrtypes.Account, error) {
	panic("unimplemented")
}

// LatestConfigDigestAndEpoch implements types.ContractTransmitter.
func (*chainReaderContract) LatestConfigDigestAndEpoch(ctx context.Context) (configDigest ocrtypes.ConfigDigest, epoch uint32, err error) {
	panic("unimplemented")
}

// Transmit implements types.ContractTransmitter.
func (*chainReaderContract) Transmit(context.Context, ocrtypes.ReportContext, ocrtypes.Report, []ocrtypes.AttributedOnchainSignature) error {
	panic("unimplemented")
}
