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

type CapabilityRegistryEvent struct {
	CapabilityID string
	DonID        DonID
}

type CapabilityRegistrySync interface {
	Listen(ctx context.Context) <-chan CapabilityRegistryEvent
}

type BaseCapabilityConfigSync struct {
	lggr         logger.Logger
	registry     CapabilityRegistry
	pollInterval time.Duration
	lastPollHash string
	capabilityID string
}

func NewBaseCapabilityConfigSync(lggr logger.Logger, registry CapabilityRegistry, pollInterval time.Duration, capabilityID string) *BaseCapabilityConfigSync {
	return &BaseCapabilityConfigSync{
		lggr:         lggr,
		registry:     registry,
		pollInterval: pollInterval,
		capabilityID: capabilityID,
		lastPollHash: "",
	}
}

func (s *BaseCapabilityConfigSync) Listen(ctx context.Context) <-chan CapabilityRegistryEvent {
	ch := make(chan CapabilityRegistryEvent)

	go func() {
		defer close(ch)
		tick := time.NewTicker(s.pollInterval)

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				events, pollHash, err := s.poll(ctx)
				if err != nil {
					s.lggr.Error("error polling capability registry", "err", err)
					continue
				}

				if len(events) == 0 {
					continue
				}

				if pollHash == s.lastPollHash {
					continue
				}

				s.lastPollHash = pollHash
				for _, event := range events {
					select {
					case ch <- event:
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}()

	return ch
}

// poll queries the capability registry to find changes and returns a list of events.
func (s *BaseCapabilityConfigSync) poll(ctx context.Context) ([]CapabilityRegistryEvent, string, error) {
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

	for donID, cfg := range dons {
		fmt.Println(donID, cfg)
	}

	return nil, donsHash, nil
}
