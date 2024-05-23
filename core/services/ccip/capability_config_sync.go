package ccip

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

type CapabilityRegistrySyncer interface {
	Listen(ctx context.Context) <-chan CapabilityRegistryDiff
}

type CapabilityRegistryDiff struct {
	Init          bool
	AddedDons     []DonAddedDetails
	RemovedDons   []DonRemovedDetails
	ConfigChanges []DONConfigChange
}

type DonAddedDetails struct {
	DonID string
}

type DonRemovedDetails struct {
	DonID string
}

type DONConfigChange struct {
	DonID string
}

type BaseCapabilityConfigSync struct {
	lggr         logger.Logger
	registry     CapabilityRegistry
	pollInterval time.Duration
	lastPollHash string
	capabilityID string
}

func NewBaseCapabilityConfigSyncer(lggr logger.Logger, registry CapabilityRegistry, pollInterval time.Duration, capabilityID string) *BaseCapabilityConfigSync {
	return &BaseCapabilityConfigSync{
		lggr:         lggr,
		registry:     registry,
		pollInterval: pollInterval,
		capabilityID: capabilityID,
		lastPollHash: "",
	}
}

func (s *BaseCapabilityConfigSync) Listen(ctx context.Context) <-chan CapabilityRegistryDiff {
	ch := make(chan CapabilityRegistryDiff)

	go func() {
		defer close(ch)
		tick := time.NewTicker(s.pollInterval)

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				diff, pollHash, err := s.poll(ctx)
				if err != nil {
					s.lggr.Error("error polling capability registry", "err", err)
					continue
				}

				if pollHash == s.lastPollHash {
					continue
				}

				s.lastPollHash = pollHash
				ch <- *diff
			}
		}
	}()

	return ch
}

// poll queries the capability registry to find changes and returns a list of events.
func (s *BaseCapabilityConfigSync) poll(ctx context.Context) (*CapabilityRegistryDiff, string, error) {
	dons, err := s.registry.GetDONsWithCapability(ctx, s.capabilityID)
	if err != nil {
		return nil, "", fmt.Errorf("get dons with capability: %w", err)
	}

	encodedDons, err := json.Marshal(dons)
	if err != nil {
		return nil, "", fmt.Errorf("marshal dons: %w", err)
	}
	hashBytes := sha256.Sum256(encodedDons)
	donsHash := hex.EncodeToString(hashBytes[:])

	if s.lastPollHash == "" {
		return &CapabilityRegistryDiff{Init: true}, donsHash, nil
	}

	// ...
	for donID, cfg := range dons {
		fmt.Println(donID, cfg)
	}
	return nil, donsHash, nil
}
