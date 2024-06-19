package launcher

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/multierr"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-common/pkg/utils"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/keystone_capability_registry"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"
)

var (
	_ job.ServiceCtx = (*launcher)(nil)
)

const (
	tickInterval = 10 * time.Second
)

func New(
	capabilityVersion,
	capabilityLabelledName string,
	p2pID p2pkey.KeyV2,
	capRegistry cctypes.CapabilityRegistry,
	lggr logger.Logger,
	homeChainReader cctypes.HomeChainReader,
	oracleCreator cctypes.OracleCreator,
) job.ServiceCtx {
	return &launcher{
		capabilityVersion:      capabilityVersion,
		capabilityLabelledName: capabilityLabelledName,
		p2pID:                  p2pID,
		capRegistry:            capRegistry,
		lggr:                   lggr,
		homeChainReader:        homeChainReader,
		regState:               cctypes.RegistryState{},
		oracleCreator:          oracleCreator,
		dons:                   make(map[uint32]*ccipDeployment),
	}
}

// launcher manages the lifecycles of the CCIP capability on all chains.
type launcher struct {
	utils.StartStopOnce

	capabilityVersion      string
	capabilityLabelledName string
	p2pID                  p2pkey.KeyV2
	capRegistry            cctypes.CapabilityRegistry
	lggr                   logger.Logger
	homeChainReader        cctypes.HomeChainReader
	stopChan               chan struct{}
	regState               cctypes.RegistryState
	oracleCreator          cctypes.OracleCreator

	// dons is a map of CCIP DON IDs to the OCR instances that are running on them.
	// we can have up to two OCR instances per CCIP plugin, since we are running two plugins,
	// thats four OCR instances per CCIP DON maximum.
	dons map[uint32]*ccipDeployment
}

// Close implements job.ServiceCtx.
func (l *launcher) Close() error {
	return l.StartStopOnce.StopOnce("launcher", func() error {
		// shut down the monitor goroutine.
		close(l.stopChan)

		// shut down all running oracles.
		var err error
		for _, ceDep := range l.dons {
			err = multierr.Append(err, ceDep.Shutdown())
		}

		return err
	})
}

// Start implements job.ServiceCtx.
func (l *launcher) Start(context.Context) error {
	return l.StartOnce("launcher", func() error {
		l.stopChan = make(chan struct{})
		go l.monitor()
		return nil
	})
}

func (l *launcher) monitor() {
	ticker := time.NewTicker(tickInterval)
	for {
		select {
		case <-l.stopChan:
			return
		case <-ticker.C:
			if err := l.tick(); err != nil {
				l.lggr.Errorw("Failed to tick", "err", err)
			}
		}
	}
}

func (l *launcher) tick() error {
	// Ensure that the home chain reader is healthy.
	// For new jobs it may be possible that the home chain reader is not yet ready
	// so we won't be able to fetch configs and start any OCR instances.
	if !l.homeChainReader.IsHealthy() {
		return fmt.Errorf("home chain reader is unhealthy")
	}

	// Fetch the latest state from the capability registry and determine if we need to
	// launch or update any OCR instances.
	latestState, err := l.capRegistry.LatestState()
	if err != nil {
		return fmt.Errorf("failed to fetch latest state from capability registry: %w", err)
	}

	diffRes, err := diff(l.capabilityVersion, l.capabilityLabelledName, l.regState, latestState)
	if err != nil {
		return fmt.Errorf("failed to diff capability registry states: %w", err)
	}

	err = l.processDiff(diffRes)
	if err != nil {
		return fmt.Errorf("failed to process diff: %w", err)
	}

	// if diff is correctly processed, update latest state.
	l.regState = latestState
	return nil
}

// processDiff processes the diff between the current and latest capability registry states.
// for any added OCR instances, it will launch them.
// for any removed OCR instances, it will shut them down.
// for any updated OCR instances, it will restart them with the new configuration.
func (l *launcher) processDiff(diff diffResult) error {
	for id := range diff.removed {
		if err := l.removeDON(id); err != nil {
			return err
		}
	}

	for _, don := range diff.added {
		if err := l.addDON(don); err != nil {
			return err
		}
	}

	for _, don := range diff.updated {
		if err := l.updateDON(don); err != nil {
			return err
		}
	}

	return nil
}

func (l *launcher) removeDON(id uint32) error {
	ceDep, ok := l.dons[id]
	if !ok {
		// not running this particular DON.
		return nil
	}

	if err := ceDep.Shutdown(); err != nil {
		return fmt.Errorf("failed to shutdown oracles for CCIP DON %d: %w", id, err)
	}

	// after a successful shutdown we can safely remove the DON deployment from the map.
	delete(l.dons, id)
	return nil
}

// updateDON handles the case where a DON in the capability registry has received a new configuration.
// In the case of CCIP, which follows blue-green deployment, we either:
// 1. Create a new oracle (the green instance) and start it.
// 2. Shut down the blue instance, making the green instance the new blue instance.
func (l *launcher) updateDON(don keystone_capability_registry.CapabilityRegistryDONInfo) error {
	if !isMemberOfDON(don, l.p2pID) {
		l.lggr.Infow("Not a member of this DON, skipping", "donId", don.Id, "p2pId", l.p2pID.ID())
		return nil
	}

	ceDep, ok := l.dons[don.Id]
	if !ok {
		// This should never happen.
		return fmt.Errorf("no deployment found for CCIP DON %d", don.Id)
	}

	// this should be a retryable error.
	commitOCRConfigs, err := l.homeChainReader.GetOCRConfigs(context.Background(), don.Id, cctypes.PluginTypeCCIPCommit)
	if err != nil {
		return fmt.Errorf("failed to fetch OCR configs for CCIP commit plugin (don id: %d) from home chain config contract: %w",
			don.Id, err)
	}

	execOCRConfigs, err := l.homeChainReader.GetOCRConfigs(context.Background(), don.Id, cctypes.PluginTypeCCIPExec)
	if err != nil {
		return fmt.Errorf("failed to fetch OCR configs for CCIP exec plugin (don id: %d) from home chain config contract: %w",
			don.Id, err)
	}

	// valid cases:
	// a) len(commitOCRConfigs) == 2 && ceDep.NumCommitInstances() == 1: this is a new green instance.
	// b) len(commitOCRConfigs) == 1 && ceDep.NumCommitInstances() == 2: this is a promotion of green->blue.
	// invalid cases (enforced in the config contract):
	// a) len(commitOCRConfigs) == 2 && ceDep.NumCommitInstances() == 2: this is an invariant violation.
	// b) len(commitOCRConfigs) == 1 && ceDep.NumCommitInstances() == 1: this is an invariant violation.
	// same thing applies to exec.
	if len(commitOCRConfigs) == 2 && !ceDep.HasGreenCommitInstance() {
		// this is a new green instance.
		greenOracle, err := l.oracleCreator.CreateCommitOracle(commitOCRConfigs[1])
		if err != nil {
			return fmt.Errorf("failed to create CCIP commit oracle: %w", err)
		}

		if err := greenOracle.Start(); err != nil {
			return fmt.Errorf("failed to start green commit oracle: %w", err)
		}
		ceDep.commit.green = greenOracle
		l.lggr.Infow("Started green commit oracle",
			"donId", don.Id, "ocrConfig", commitOCRConfigs[1].String())
	} else if len(commitOCRConfigs) == 1 && ceDep.HasGreenCommitInstance() {
		// this is a promotion of green->blue.
		// swap the green oracle with the blue oracle in the ceDep struct.
		oldBlue := ceDep.commit.blue
		ceDep.commit.blue = ceDep.commit.green
		ceDep.commit.green = nil

		// shut down blue oracle.
		if err := oldBlue.Shutdown(); err != nil {
			// we can't really roll back here, so we just log the error.
			l.lggr.Errorw("Failed to shutdown blue oracle", "err", err)
		}
	} else {
		return fmt.Errorf("invariant violation: expected 1 or 2 OCR configs for CCIP commit plugin (don id: %d), got %d", don.Id, len(commitOCRConfigs))
	}

	if len(execOCRConfigs) == 2 && !ceDep.HasGreenExecInstance() {
		// this is a new green instance.
		greenOracle, err := l.oracleCreator.CreateExecOracle(execOCRConfigs[1])
		if err != nil {
			return fmt.Errorf("failed to create CCIP exec oracle: %w", err)
		}

		if err := greenOracle.Start(); err != nil {
			return fmt.Errorf("failed to start green exec oracle: %w", err)
		}
		ceDep.exec.green = greenOracle
		l.lggr.Infow("Started green exec oracle",
			"donId", don.Id, "ocrConfig", execOCRConfigs[1].String())
	} else if len(execOCRConfigs) == 1 && ceDep.HasGreenExecInstance() {
		// this is a promotion of green->blue.
		// swap the green oracle with the blue oracle in the ceDep struct.
		oldBlue := ceDep.exec.blue
		ceDep.exec.blue = ceDep.exec.green
		ceDep.exec.green = nil

		// shut down blue oracle.
		if err := oldBlue.Shutdown(); err != nil {
			// we can't really roll back here, so we just log the error.
			l.lggr.Errorw("Failed to shutdown blue oracle", "err", err)
		}
	} else {
		return fmt.Errorf("invariant violation: expected 1 or 2 OCR configs for CCIP exec plugin (don id: %d), got %d", don.Id, len(execOCRConfigs))
	}

	return nil
}

func (l *launcher) addDON(don keystone_capability_registry.CapabilityRegistryDONInfo) error {
	if !isMemberOfDON(don, l.p2pID) {
		l.lggr.Infow("Not a member of this DON, skipping", "donId", don.Id, "p2pId", l.p2pID.ID())
		return nil
	}

	// this should be a retryable error.
	commitOCRConfigs, err := l.homeChainReader.GetOCRConfigs(context.Background(), don.Id, cctypes.PluginTypeCCIPCommit)
	if err != nil {
		return fmt.Errorf("failed to fetch OCR configs for CCIP commit plugin (don id: %d) from home chain config contract: %w",
			don.Id, err)
	}

	execOCRConfigs, err := l.homeChainReader.GetOCRConfigs(context.Background(), don.Id, cctypes.PluginTypeCCIPExec)
	if err != nil {
		return fmt.Errorf("failed to fetch OCR configs for CCIP exec plugin (don id: %d) from home chain config contract: %w",
			don.Id, err)
	}

	// upon creation we should only have one OCR config per plugin type.
	if len(commitOCRConfigs) != 1 {
		return fmt.Errorf("expected exactly one OCR config for CCIP commit plugin (don id: %d), got %d", don.Id, len(commitOCRConfigs))
	}

	if len(execOCRConfigs) != 1 {
		return fmt.Errorf("expected exactly one OCR config for CCIP exec plugin (don id: %d), got %d", don.Id, len(execOCRConfigs))
	}

	commitOracle, err := l.oracleCreator.CreateCommitOracle(commitOCRConfigs[0])
	if err != nil {
		return fmt.Errorf("failed to create CCIP commit oracle: %w", err)
	}

	execOracle, err := l.oracleCreator.CreateExecOracle(execOCRConfigs[0])
	if err != nil {
		return fmt.Errorf("failed to create CCIP exec oracle: %w", err)
	}

	var errGroup errgroup.Group
	errGroup.Go(func() error {
		return commitOracle.Start()
	})
	errGroup.Go(func() error {
		return execOracle.Start()
	})
	err = errGroup.Wait()
	if err != nil {
		// shut down the oracles if we failed to start them.
		errShutdown := commitOracle.Shutdown()
		if errShutdown != nil {
			l.lggr.Errorw("Failed to shutdown commit oracle", "err", err)
		}
		errShutdown = execOracle.Shutdown()
		if errShutdown != nil {
			l.lggr.Errorw("Failed to shutdown exec oracle", "err", err)
		}
		return fmt.Errorf("failed to start oracles for CCIP DON %d: %w, err shutdown (could be nil): %w", don.Id, err, errShutdown)
	}

	// update the dons map with the new deployment.
	l.dons[don.Id] = &ccipDeployment{
		commit: blueGreenDeployment{
			blue: commitOracle,
		},
		exec: blueGreenDeployment{
			blue: execOracle,
		},
	}

	return nil
}
