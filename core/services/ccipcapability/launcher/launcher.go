package launcher

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.uber.org/multierr"

	ocr3reader "github.com/smartcontractkit/ccipocr3/pkg/reader"
	"github.com/smartcontractkit/chainlink-common/pkg/utils"

	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"
	p2ptypes "github.com/smartcontractkit/chainlink/v2/core/services/p2p/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/registrysyncer"
)

var (
	_ job.ServiceCtx          = (*launcher)(nil)
	_ registrysyncer.Launcher = (*launcher)(nil)
)

const (
	tickInterval = 10 * time.Second
)

func New(
	capabilityVersion,
	capabilityLabelledName string,
	p2pID p2pkey.KeyV2,
	lggr logger.Logger,
	homeChainReader ocr3reader.HomeChain,
	oracleCreator cctypes.OracleCreator,
) *launcher {
	return &launcher{
		capabilityVersion:      capabilityVersion,
		capabilityLabelledName: capabilityLabelledName,
		p2pID:                  p2pID,
		lggr:                   lggr,
		homeChainReader:        homeChainReader,
		regState: registrysyncer.State{
			IDsToDONs:         make(map[registrysyncer.DonID]kcr.CapabilitiesRegistryDONInfo),
			IDsToNodes:        make(map[p2ptypes.PeerID]kcr.CapabilitiesRegistryNodeInfo),
			IDsToCapabilities: make(map[registrysyncer.HashedCapabilityID]kcr.CapabilitiesRegistryCapabilityInfo),
		},
		oracleCreator: oracleCreator,
		dons:          make(map[registrysyncer.DonID]*ccipDeployment),
	}
}

// launcher manages the lifecycles of the CCIP capability on all chains.
type launcher struct {
	utils.StartStopOnce

	capabilityVersion      string
	capabilityLabelledName string
	p2pID                  p2pkey.KeyV2
	lggr                   logger.Logger
	homeChainReader        ocr3reader.HomeChain
	stopChan               chan struct{}
	// latestState is the latest capability registry state received from the syncer.
	latestState registrysyncer.State
	// regState is the latest capability registry state that we have successfully processed.
	regState      registrysyncer.State
	oracleCreator cctypes.OracleCreator
	lock          sync.RWMutex

	// dons is a map of CCIP DON IDs to the OCR instances that are running on them.
	// we can have up to two OCR instances per CCIP plugin, since we are running two plugins,
	// thats four OCR instances per CCIP DON maximum.
	dons map[registrysyncer.DonID]*ccipDeployment
}

// Launch implements registrysyncer.Launcher.
func (l *launcher) Launch(ctx context.Context, state registrysyncer.State) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.latestState = state
	return nil
}

func (l *launcher) getLatestState() registrysyncer.State {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.latestState
}

// Close implements job.ServiceCtx.
func (l *launcher) Close() error {
	return l.StartStopOnce.StopOnce("launcher", func() error {
		// shut down the monitor goroutine.
		close(l.stopChan)

		// shut down all running oracles.
		var err error
		for _, ceDep := range l.dons {
			err = multierr.Append(err, ceDep.Close())
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
	if ready := l.homeChainReader.Ready(); ready != nil {
		return fmt.Errorf("home chain reader is not ready: %w", ready)
	}

	// Fetch the latest state from the capability registry and determine if we need to
	// launch or update any OCR instances.
	latestState := l.getLatestState()

	diffRes, err := diff(l.capabilityVersion, l.capabilityLabelledName, l.regState, latestState)
	if err != nil {
		return fmt.Errorf("failed to diff capability registry states: %w", err)
	}

	err = l.processDiff(diffRes)
	if err != nil {
		return fmt.Errorf("failed to process diff: %w", err)
	}

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

		delete(l.regState.IDsToDONs, id)
	}

	var addedDeployments = make(map[registrysyncer.DonID]*ccipDeployment)
	for _, don := range diff.added {
		dep, err := l.addDON(don)
		if err != nil {
			return err
		}
		addedDeployments[registrysyncer.DonID(don.Id)] = dep
	}

	for donID, dep := range addedDeployments {
		if err := dep.StartBlue(); err != nil {
			if shutdownErr := dep.CloseBlue(); shutdownErr != nil {
				l.lggr.Errorw("Failed to shutdown blue instance after failed start", "donId", donID, "err", shutdownErr)
			}
			return fmt.Errorf("failed to start oracles for CCIP DON %d: %w", donID, err)
		}
		// update state.
		l.dons[donID] = dep
		// update the state with the latest config.
		// this way if one of the starts errors, we don't retry all of them.
		l.regState.IDsToDONs[donID] = diff.added[donID]
	}

	var updatedDeployments = make(map[registrysyncer.DonID]struct {
		depBefore, depAfter *ccipDeployment
	})
	for _, don := range diff.updated {
		depBefore, depAfter, err := l.updateDON(don)
		if err != nil {
			return err
		}
		updatedDeployments[registrysyncer.DonID(don.Id)] = struct {
			depBefore, depAfter *ccipDeployment
		}{
			depBefore: depBefore,
			depAfter:  depAfter,
		}
	}

	for donID, depPair := range updatedDeployments {
		if err := depPair.depAfter.HandleBlueGreen(depPair.depBefore); err != nil {
			// TODO: how to handle a failed blue-green deployment?
			return fmt.Errorf("failed to handle blue-green deployment for CCIP DON %d: %w", donID, err)
		}

		// update state.
		l.dons[donID] = depPair.depAfter
		// update the state with the latest config.
		// this way if one of the starts errors, we don't retry all of them.
		l.regState.IDsToDONs[donID] = diff.updated[donID]
	}

	return nil
}

func (l *launcher) removeDON(id registrysyncer.DonID) error {
	ceDep, ok := l.dons[id]
	if !ok {
		// not running this particular DON.
		return nil
	}

	if err := ceDep.Close(); err != nil {
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
func (l *launcher) updateDON(don kcr.CapabilitiesRegistryDONInfo) (depBefore, depAfter *ccipDeployment, err error) {
	if !isMemberOfDON(don, l.p2pID) {
		l.lggr.Infow("Not a member of this DON, skipping", "donId", don.Id, "p2pId", l.p2pID.ID())
		return nil, nil, nil
	}

	var ok bool
	depBefore, ok = l.dons[registrysyncer.DonID(don.Id)]
	if !ok {
		// This should never happen.
		return nil, nil, fmt.Errorf("no deployment found for CCIP DON %d", don.Id)
	}

	// this should be a retryable error.
	commitOCRConfigs, err := l.homeChainReader.GetOCRConfigs(don.Id, uint8(cctypes.PluginTypeCCIPCommit))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch OCR configs for CCIP commit plugin (don id: %d) from home chain config contract: %w",
			don.Id, err)
	}

	execOCRConfigs, err := l.homeChainReader.GetOCRConfigs(don.Id, uint8(cctypes.PluginTypeCCIPExec))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch OCR configs for CCIP exec plugin (don id: %d) from home chain config contract: %w",
			don.Id, err)
	}

	// valid cases:
	// a) len(commitOCRConfigs) == 2 && depBefore.NumCommitInstances() == 1: this is a new green instance.
	// b) len(commitOCRConfigs) == 1 && depBefore.NumCommitInstances() == 2: this is a promotion of green->blue.
	// invalid cases (enforced in the config contract):
	// a) len(commitOCRConfigs) == 2 && depBefore.NumCommitInstances() == 2: this is an invariant violation.
	// b) len(commitOCRConfigs) == 1 && depBefore.NumCommitInstances() == 1: this is an invariant violation.
	// same thing applies to exec.
	depAfter = &ccipDeployment{}
	if len(commitOCRConfigs) == 2 && !depBefore.HasGreenCommitInstance() {
		// this is a new green instance.
		greenOracle, err := l.oracleCreator.CreateCommitOracle(commitOCRConfigs[1])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create CCIP commit oracle: %w", err)
		}

		depAfter.commit.blue = depBefore.commit.blue
		depAfter.commit.green = greenOracle
	} else if len(commitOCRConfigs) == 1 && depBefore.HasGreenCommitInstance() {
		// this is a promotion of green->blue.
		depAfter.commit.blue = depBefore.commit.green
	} else {
		return nil, nil, fmt.Errorf("invariant violation: expected 1 or 2 OCR configs for CCIP commit plugin (don id: %d), got %d", don.Id, len(commitOCRConfigs))
	}

	if len(execOCRConfigs) == 2 && !depBefore.HasGreenExecInstance() {
		// this is a new green instance.
		greenOracle, err := l.oracleCreator.CreateExecOracle(execOCRConfigs[1])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create CCIP exec oracle: %w", err)
		}

		depAfter.exec.blue = depBefore.exec.blue
		depAfter.exec.green = greenOracle
	} else if len(execOCRConfigs) == 1 && depBefore.HasGreenExecInstance() {
		// this is a promotion of green->blue.
		depAfter.exec.blue = depBefore.exec.green
	} else {
		return nil, nil, fmt.Errorf("invariant violation: expected 1 or 2 OCR configs for CCIP exec plugin (don id: %d), got %d", don.Id, len(execOCRConfigs))
	}

	return depBefore, depAfter, nil
}

func (l *launcher) addDON(don kcr.CapabilitiesRegistryDONInfo) (*ccipDeployment, error) {
	if !isMemberOfDON(don, l.p2pID) {
		l.lggr.Infow("Not a member of this DON, skipping", "donId", don.Id, "p2pId", l.p2pID.ID())
		return nil, nil
	}

	// this should be a retryable error.
	commitOCRConfigs, err := l.homeChainReader.GetOCRConfigs(don.Id, uint8(cctypes.PluginTypeCCIPCommit))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch OCR configs for CCIP commit plugin (don id: %d) from home chain config contract: %w",
			don.Id, err)
	}

	execOCRConfigs, err := l.homeChainReader.GetOCRConfigs(don.Id, uint8(cctypes.PluginTypeCCIPExec))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch OCR configs for CCIP exec plugin (don id: %d) from home chain config contract: %w",
			don.Id, err)
	}

	// upon creation we should only have one OCR config per plugin type.
	if len(commitOCRConfigs) != 1 {
		return nil, fmt.Errorf("expected exactly one OCR config for CCIP commit plugin (don id: %d), got %d", don.Id, len(commitOCRConfigs))
	}

	if len(execOCRConfigs) != 1 {
		return nil, fmt.Errorf("expected exactly one OCR config for CCIP exec plugin (don id: %d), got %d", don.Id, len(execOCRConfigs))
	}

	commitOracle, err := l.oracleCreator.CreateCommitOracle(commitOCRConfigs[0])
	if err != nil {
		return nil, fmt.Errorf("failed to create CCIP commit oracle: %w", err)
	}

	var commitBootstrap cctypes.CCIPOracle
	if isMemberOfBootstrapSubcommittee(commitOCRConfigs[0].Config.BootstrapP2PIds, l.p2pID) {
		commitBootstrap, err = l.oracleCreator.CreateBootstrapOracle(commitOCRConfigs[0])
		if err != nil {
			return nil, fmt.Errorf("failed to create CCIP bootstrap oracle: %w", err)
		}
	}

	execOracle, err := l.oracleCreator.CreateExecOracle(execOCRConfigs[0])
	if err != nil {
		return nil, fmt.Errorf("failed to create CCIP exec oracle: %w", err)
	}

	var execBootstrap cctypes.CCIPOracle
	if isMemberOfBootstrapSubcommittee(execOCRConfigs[0].Config.BootstrapP2PIds, l.p2pID) {
		execBootstrap, err = l.oracleCreator.CreateBootstrapOracle(execOCRConfigs[0])
		if err != nil {
			return nil, fmt.Errorf("failed to create CCIP bootstrap oracle: %w", err)
		}
	}

	return &ccipDeployment{
		commit: blueGreenDeployment{
			blue:          commitOracle,
			bootstrapBlue: commitBootstrap,
		},
		exec: blueGreenDeployment{
			blue:          execOracle,
			bootstrapBlue: execBootstrap,
		},
	}, nil
}
