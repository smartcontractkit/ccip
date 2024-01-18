package ocr3impls

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/multierr"

	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/services"

	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/no_op_ocr3"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditymanager"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay"
)

var (
	// See https://github.com/smartcontractkit/ccip/compare/ccip-develop...CCIP-1438-op-stack-bridge-adapter-l-1#diff-2fe14bb9d1ecbc62f43cef26daff5d1f86275f16e1296fc9827b934a518d3f4cR20
	ConfigSet common.Hash

	defaultABI abi.ABI

	_ ocrtypes.ContractConfigTracker = &multichainConfigTracker{}

	defaultTimeout = 1 * time.Minute
)

func init() {
	var err error
	tabi, err := no_op_ocr3.NoOpOCR3MetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	defaultABI = *tabi
	ConfigSet = defaultABI.Events["ConfigSet"].ID
}

type CombinerFn func(masterChain relay.ID, contractConfigs map[relay.ID]ocrtypes.ContractConfig) (ocrtypes.ContractConfig, error)

type multichainConfigTracker struct {
	services.StateMachine

	// masterChain is the chain that contains the "master" OCR3 configuration
	// contract.
	masterChain       relay.ID
	lggr              logger.Logger
	logPollers        map[relay.ID]logpoller.LogPoller
	masterClient      evmclient.Client
	contractAddresses map[relay.ID]common.Address
	masterContract    no_op_ocr3.NoOpOCR3Interface
	combiner          CombinerFn
	fromBlocks        map[string]int64
	replaying         atomic.Bool
}

func NewMultichainConfigTracker(
	masterChain relay.ID,
	lggr logger.Logger,
	logPollers map[relay.ID]logpoller.LogPoller,
	masterClient evmclient.Client,
	masterContract common.Address,
	lmFactory liquiditymanager.Factory,
	combiner CombinerFn,
	fromBlocks map[string]int64,
) (*multichainConfigTracker, error) {
	// Ensure master chain is in the log pollers
	if _, ok := logPollers[masterChain]; !ok {
		return nil, fmt.Errorf("master chain %s not in log pollers", masterChain)
	}

	// Ensure combiner is not nil
	if combiner == nil {
		return nil, fmt.Errorf("provide non-nil combiner")
	}

	// Ensure factory is not nil
	if lmFactory == nil {
		return nil, fmt.Errorf("provide non-nil liquidity manager factory")
	}

	// before we register filters we need to discover all the available liquidity managers
	// on all the chains
	masterChainID, err := strconv.ParseInt(masterChain.ChainID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse network ID %s: %w", masterChain, err)
	}
	masterLM, err := lmFactory.NewRebalancer(
		models.NetworkSelector(masterChainID), // todo: probably need to find selector from chain id first
		models.Address(masterContract))
	if err != nil {
		return nil, fmt.Errorf("failed to create master liquidity manager: %w", err)
	}

	// Discover all the liquidity managers
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	lms, _, err := masterLM.Discover(ctx, lmFactory)
	if err != nil {
		return nil, fmt.Errorf("failed to discover liquidity managers: %w", err)
	}
	all := lms.GetAll()

	// sanity check, there should be only one liquidity manager per-chain per-asset
	if len(all) != len(logPollers) {
		return nil, fmt.Errorf("expected %d liquidity managers but found %d", len(logPollers), len(all))
	}

	// Register filters on all log pollers
	contracts := make(map[relay.ID]common.Address)
	for id, lp := range logPollers {
		nid, err2 := strconv.ParseInt(id.ChainID, 10, 64)
		if err2 != nil {
			return nil, fmt.Errorf("failed to parse network ID %s: %w", id, err2)
		}
		address, ok := all[models.NetworkSelector(nid)]
		if !ok {
			return nil, fmt.Errorf("no liquidity manager found for network ID %d", nid)
		}
		fName := configTrackerFilterName(id, common.Address(address))
		err2 = lp.RegisterFilter(logpoller.Filter{
			Name:      fName,
			EventSigs: []common.Hash{ConfigSet},
			Addresses: []common.Address{common.Address(address)},
		})
		if err2 != nil {
			return nil, err2
		}
		contracts[id] = common.Address(address)
	}

	wrapper, err := no_op_ocr3.NewNoOpOCR3(masterContract, masterClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create wrapper: %w", err)
	}

	return &multichainConfigTracker{
		lggr:              lggr,
		logPollers:        logPollers,
		masterChain:       masterChain,
		combiner:          combiner,
		masterClient:      masterClient,
		contractAddresses: contracts,
		masterContract:    wrapper,
		fromBlocks:        fromBlocks,
	}, nil
}

func (m *multichainConfigTracker) GetContractAddresses() map[relay.ID]common.Address {
	return m.contractAddresses
}

func (m *multichainConfigTracker) Start() {
	_ = m.StartOnce("MultichainConfigTracker", func() error {
		if m.fromBlocks != nil {
			m.replaying.Store(true)
			defer m.replaying.Store(false)

			// TODO: replay multiple chains in parallel?
			var errs error
			for id, fromBlock := range m.fromBlocks {
				err := m.ReplayChain(context.Background(), relay.NewID("evm", id), fromBlock)
				if err != nil {
					m.lggr.Errorw("failed to replay chain", "chain", id, "fromBlock", fromBlock, "err", err)
					errs = multierr.Append(errs, err)
				} else {
					m.lggr.Infow("successfully replayed chain", "chain", id, "fromBlock", fromBlock)
				}
			}

			if errs != nil {
				m.lggr.Errorw("failed to replay some chains", "err", errs)
				return errs
			}
		}
		return nil
	})
}

func (m *multichainConfigTracker) Close() error {
	return nil
}

// Notify noop method
func (m *multichainConfigTracker) Notify() <-chan struct{} {
	return nil
}

// ReplayChain replays the log poller for the provided chain
func (m *multichainConfigTracker) ReplayChain(ctx context.Context, id relay.ID, fromBlock int64) error {
	if _, ok := m.logPollers[id]; !ok {
		return fmt.Errorf("chain %s not in log pollers", id)
	}
	return m.logPollers[id].Replay(ctx, fromBlock)
}

// Replay replays the log poller for the master chain
func (m *multichainConfigTracker) Replay(ctx context.Context, fromBlock int64) error {
	return m.logPollers[m.masterChain].Replay(ctx, fromBlock)
}

// LatestBlockHeight implements types.ContractConfigTracker.
// Returns the block height of the master chain.
func (m *multichainConfigTracker) LatestBlockHeight(ctx context.Context) (blockHeight uint64, err error) {
	latestBlock, err := m.logPollers[m.masterChain].LatestBlock(pg.WithParentCtx(ctx))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}
	return uint64(latestBlock.BlockNumber), nil
}

// LatestConfig implements types.ContractConfigTracker.
// LatestConfig fetches the config from the master chain and then fetches the
// remaining configurations from all the other chains.
func (m *multichainConfigTracker) LatestConfig(ctx context.Context, changedInBlock uint64) (ocrtypes.ContractConfig, error) {
	// if we're still replaying the follower chains we won't have their configs
	if m.replaying.Load() {
		return ocrtypes.ContractConfig{}, errors.New("cannot call LatestConfig while replaying")
	}

	lgs, err := m.logPollers[m.masterChain].Logs(int64(changedInBlock), int64(changedInBlock), ConfigSet, m.contractAddresses[m.masterChain], pg.WithParentCtx(ctx))
	if err != nil {
		return ocrtypes.ContractConfig{}, err
	}
	if len(lgs) == 0 {
		return ocrtypes.ContractConfig{}, fmt.Errorf("no logs found for config on contract %s (chain %s) at block %d",
			m.contractAddresses[m.masterChain].Hex(), m.masterChain.String(), changedInBlock)
	}
	masterConfig, err := configFromLog(lgs[len(lgs)-1].Data)
	if err != nil {
		return ocrtypes.ContractConfig{}, err
	}
	m.lggr.Infow("LatestConfig from master chain", "latestConfig", masterConfig)

	// check all other chains for their config
	contractConfigs := map[relay.ID]ocrtypes.ContractConfig{
		m.masterChain: masterConfig,
	}
	for id, lp := range m.logPollers {
		if id == m.masterChain {
			continue
		}

		lg, err2 := lp.LatestLogByEventSigWithConfs(ConfigSet, m.contractAddresses[id], 1, pg.WithParentCtx(ctx))
		if err2 != nil {
			return ocrtypes.ContractConfig{}, err2
		}

		configSet, err2 := configFromLog(lg.Data)
		if err2 != nil {
			return ocrtypes.ContractConfig{}, err2
		}
		contractConfigs[id] = configSet
	}

	// at this point we can combine the configs into a single one
	combined, err := m.combiner(m.masterChain, contractConfigs)
	if err != nil {
		return ocrtypes.ContractConfig{}, fmt.Errorf("error combining configs: %w", err)
	}

	return combined, nil
}

// LatestConfigDetails implements types.ContractConfigTracker.
func (m *multichainConfigTracker) LatestConfigDetails(ctx context.Context) (changedInBlock uint64, configDigest ocrtypes.ConfigDigest, err error) {
	latest, err := m.logPollers[m.masterChain].LatestLogByEventSigWithConfs(ConfigSet, m.contractAddresses[m.masterChain], 1, pg.WithParentCtx(ctx))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return m.callLatestConfigDetails(ctx)
		}
		return 0, ocrtypes.ConfigDigest{}, err
	}
	masterConfig, err := configFromLog(latest.Data)
	if err != nil {
		return 0, ocrtypes.ConfigDigest{}, fmt.Errorf("failed to unpack latest config details: %w", err)
	}

	return uint64(latest.BlockNumber), masterConfig.ConfigDigest, nil
}

func (m *multichainConfigTracker) callLatestConfigDetails(ctx context.Context) (changedInBlock uint64, configDigest ocrtypes.ConfigDigest, err error) {
	lcd, err := m.masterContract.LatestConfigDetails(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return 0, ocrtypes.ConfigDigest{}, fmt.Errorf("failed to get latest config details: %w", err)
	}
	return uint64(lcd.BlockNumber), lcd.ConfigDigest, nil
}
