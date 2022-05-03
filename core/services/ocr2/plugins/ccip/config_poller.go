package ccip

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	confighelper "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/logger"
)

func ContractConfigFromConfigSetEvent(changed offramp.OffRampConfigSet) ocrtypes.ContractConfig {
	var transmitAccounts []ocrtypes.Account
	for _, addr := range changed.Transmitters {
		transmitAccounts = append(transmitAccounts, ocrtypes.Account(addr.Hex()))
	}
	var signers []ocrtypes.OnchainPublicKey
	for _, addr := range changed.Signers {
		addr := addr
		signers = append(signers, addr[:])
	}
	return ocrtypes.ContractConfig{
		ConfigDigest:          changed.ConfigDigest,
		ConfigCount:           changed.ConfigCount,
		Signers:               signers,
		Transmitters:          transmitAccounts,
		F:                     changed.F,
		OnchainConfig:         changed.OnchainConfig,
		OffchainConfigVersion: changed.OffchainConfigVersion,
		OffchainConfig:        changed.OffchainConfig,
	}
}

type ConfigPoller struct {
	lggr               logger.Logger
	destChainLogPoller *logpoller.LogPoller
	offRamp            *offramp.OffRamp

	offchainConfigMu sync.RWMutex
	offchainConfig   OffchainConfig
	configDigest     [32]byte

	pollPeriod time.Duration
	done       chan struct{}
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewConfigPoller(lggr logger.Logger, destChainPoller *logpoller.LogPoller, offRamp *offramp.OffRamp, pollPeriod time.Duration) *ConfigPoller {
	destChainPoller.MergeFilter([]common.Hash{ConfigSet}, offRamp.Address())
	return &ConfigPoller{
		lggr:               lggr,
		destChainLogPoller: destChainPoller,
		offRamp:            offRamp,
		offchainConfigMu:   sync.RWMutex{},
		offchainConfig: OffchainConfig{
			// Consider everything unconfirmed until we get the onchain config
			// and update this.
			SourceIncomingConfirmations: 10000,
			DestIncomingConfirmations:   10000,
		},
		pollPeriod: pollPeriod,
		done:       make(chan struct{}),
	}
}

func (lp *ConfigPoller) sourceConfs() int {
	lp.offchainConfigMu.RLock()
	defer lp.offchainConfigMu.RUnlock()
	return int(lp.offchainConfig.SourceIncomingConfirmations)
}

// TODO: may not need
func (lp *ConfigPoller) destConfs() int {
	lp.offchainConfigMu.RLock()
	defer lp.offchainConfigMu.RUnlock()
	return int(lp.offchainConfig.DestIncomingConfirmations)
}

func (lp *ConfigPoller) Start(ctx context.Context) error {
	// Don't use start context, no async calls in startup.
	ctx, cancel := context.WithCancel(context.Background())
	lp.ctx = ctx
	lp.cancel = cancel
	go lp.run()
	return nil
}

func (lp *ConfigPoller) Close() error {
	lp.cancel()
	<-lp.done
	return nil
}

func (lp *ConfigPoller) run() {
	defer close(lp.done)
	tick := time.After(0)
	lp.lggr.Infow("Starting config poller")
	for {
		select {
		case <-lp.ctx.Done():
			lp.lggr.Infow("Closing config poller")
			return
		case <-tick:
			tick = time.After(lp.pollPeriod)
			if err := lp.processConfigSet(); err != nil {
				lp.lggr.Errorw("Error processing config set", "err", err)
			}
		}
	}
}

func (lp *ConfigPoller) processConfigSet() error {
	lp.offchainConfigMu.Lock()
	defer lp.offchainConfigMu.Unlock()
	configDetails, err := lp.offRamp.LatestConfigDetails(nil)
	if err != nil {
		return err
	}
	if configDetails.ConfigDigest == lp.configDigest {
		return nil
	}
	// Otherwise config digest has changed.
	// We don't expect reorgs to change contents of config set.
	latestConfigSetLog, err := lp.destChainLogPoller.LatestLogByEventSigWithConfs(ConfigSet, lp.offRamp.Address(), 1)
	if err != nil {
		// Its possible we haven't got any logs yet.
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}
	latestConfigSet, err := lp.offRamp.ParseConfigSet(types.Log{
		Data:   latestConfigSetLog.Data,
		Topics: latestConfigSetLog.GetTopics(),
	})
	if err != nil {
		return err
	}
	if latestConfigSet.ConfigDigest == configDetails.ConfigDigest {
		// Only update once we have the log
		publicConfig, err := confighelper.PublicConfigFromContractConfig(false, ContractConfigFromConfigSetEvent(*latestConfigSet))
		if err != nil {
			return err
		}
		lp.lggr.Infow("config change detected", "new digest", configDetails.ConfigDigest, "old digest", lp.configDigest)
		newConfig, err := Decode(publicConfig.ReportingPluginConfig)
		if err != nil {
			lp.lggr.Errorw("cannot decode config", "err", err, "offchainConfig", hexutil.Encode(latestConfigSet.OffchainConfig))
			return err
		}
		lp.configDigest = latestConfigSet.ConfigDigest
		lp.offchainConfig = newConfig
	}
	// Don't have the latest log yet, keep waiting.
	return nil
}
