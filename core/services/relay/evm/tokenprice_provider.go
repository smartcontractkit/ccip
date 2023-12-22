package evm

import (
	"context"
	"errors"

	"github.com/smartcontractkit/chainlink-common/pkg/services"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

type TokenPriceProvider interface {
	commontypes.PluginProvider
}

var _ TokenPriceProvider = (*tokenPriceProvider)(nil)

type tokenPriceProvider struct {
	configWatcher       *configWatcher
	contractTransmitter *contractTransmitter
	lggr                logger.Logger
	services.StateMachine
}

func NewTokenPriceProvider(configWatcher *configWatcher, contractTransmitter *contractTransmitter, lggr logger.Logger) *tokenPriceProvider {
	return &tokenPriceProvider{
		configWatcher:       configWatcher,
		contractTransmitter: contractTransmitter,
		lggr:                lggr,
	}
}

func (p *tokenPriceProvider) Start(ctx context.Context) error {
	return p.StartOnce("TokenPriceProvider", func() error {
		return nil
	})
}

func (p *tokenPriceProvider) Close() error {
	return p.StopOnce("TokenPriceProvider", func() error {
		return nil
	})
}

func (p *tokenPriceProvider) Ready() error {
	return errors.Join(p.configWatcher.Ready())
}

// ChainReader implements TokenPriceProvider.
func (*tokenPriceProvider) ChainReader() commontypes.ChainReader {
	return nil
}

// ContractConfigTracker implements TokenPriceProvider.
func (t *tokenPriceProvider) ContractConfigTracker() types.ContractConfigTracker {
	return t.configWatcher.ContractConfigTracker()
}

// ContractTransmitter implements TokenPriceProvider.
func (t *tokenPriceProvider) ContractTransmitter() types.ContractTransmitter {
	return t.contractTransmitter
}

// HealthReport implements TokenPriceProvider.
func (t *tokenPriceProvider) HealthReport() map[string]error {
	report := map[string]error{}
	services.CopyHealth(report, t.configWatcher.HealthReport())
	return report
}

// Name implements TokenPriceProvider.
func (t *tokenPriceProvider) Name() string {
	return t.lggr.Name()
}

// OffchainConfigDigester implements TokenPriceProvider.
func (t *tokenPriceProvider) OffchainConfigDigester() types.OffchainConfigDigester {
	t.configWatcher.OffchainConfigDigester()
}
