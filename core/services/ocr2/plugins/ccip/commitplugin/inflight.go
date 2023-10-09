package commitplugin

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

const (
	priceExpiryMultiplier = 3 // Keep price update cache longer and use it as source of truth
)

// inflightReport represents a commit report which has been submitted
// to the transaction manager, and we expect to be included in the chain.
// By keeping track of the inflight reports, we are able to build subsequent
// reports "on top" of the inflight ones for improved throughput - for example
// if seqNrs=[1,2] are inflight, we can build and send [3,4] while [1,2] is still confirming
// and optimistically assume they will complete in order. If for whatever reason (re-org or
// RPC timing) leads to [3,4] arriving before [1,2], we'll revert onchain. Once the cache
// expires we'll then build from the onchain state again and retries. In this manner,
// we are able to obtain high throughput during happy path yet still naturally recover
// if a reorg or issue causes onchain reverts.
type inflightReport struct {
	report    ccipdata.CommitStoreReport
	createdAt time.Time
}

type inflightPriceUpdate struct {
	gasPrices     []ccipdata.GasPrice
	tokenPrices   []ccipdata.TokenPrice
	createdAt     time.Time
	epochAndRound uint64
}

// inflightExecReportsContainer holds existing inflight reports.
// it provides a thread-safe access as it is called from multiple goroutines,
// e.g. reporting and transmission protocols.
type inflightCommitReportsContainer struct {
	locker               sync.RWMutex
	inFlight             map[[32]byte]inflightReport
	inFlightPriceUpdates []inflightPriceUpdate
	cacheExpiry          time.Duration
}

func newInflightCommitReportsContainer(inflightCacheExpiry time.Duration) *inflightCommitReportsContainer {
	return &inflightCommitReportsContainer{
		locker:               sync.RWMutex{},
		inFlight:             make(map[[32]byte]inflightReport),
		inFlightPriceUpdates: []inflightPriceUpdate{},
		cacheExpiry:          inflightCacheExpiry,
	}
}

func (c *inflightCommitReportsContainer) maxInflightSeqNr() uint64 {
	c.locker.RLock()
	defer c.locker.RUnlock()
	var maxVal uint64
	for _, report := range c.inFlight {
		if report.report.Interval.Max >= maxVal {
			maxVal = report.report.Interval.Max
		}
	}
	return maxVal
}

// getLatestInflightGasPriceUpdate returns the latest inflight gas price update, and bool flag on if update exists.
// Note we assume that reports contain either 1 or 0 gas prices.
// If this assumption is broken, we will need to update this logic.
func (c *inflightCommitReportsContainer) getLatestInflightGasPriceUpdate() (update, bool) {
	c.locker.RLock()
	defer c.locker.RUnlock()
	updateFound := false
	latestGasPriceUpdate := update{}
	var latestEpochAndRound uint64
	for _, inflight := range c.inFlightPriceUpdates {
		if len(inflight.gasPrices) == 0 {
			// Price updates did not include a gas price
			continue
		}
		if !updateFound || inflight.epochAndRound > latestEpochAndRound {
			// First price found or found later update, set it
			updateFound = true
			latestGasPriceUpdate = update{
				timestamp: inflight.createdAt,
				value:     inflight.gasPrices[0].Value,
			}
			latestEpochAndRound = inflight.epochAndRound
			continue
		}
	}
	return latestGasPriceUpdate, updateFound
}

// latestInflightTokenPriceUpdates returns a map of the latest token price updates
func (c *inflightCommitReportsContainer) latestInflightTokenPriceUpdates() map[common.Address]update {
	c.locker.RLock()
	defer c.locker.RUnlock()
	latestTokenPriceUpdates := make(map[common.Address]update)
	latestEpochAndRounds := make(map[common.Address]uint64)
	for _, inflight := range c.inFlightPriceUpdates {
		for _, inflightTokenUpdate := range inflight.tokenPrices {
			_, ok := latestTokenPriceUpdates[inflightTokenUpdate.Token]
			if !ok || inflight.epochAndRound > latestEpochAndRounds[inflightTokenUpdate.Token] {
				latestTokenPriceUpdates[inflightTokenUpdate.Token] = update{
					value:     inflightTokenUpdate.Value,
					timestamp: inflight.createdAt,
				}
				latestEpochAndRounds[inflightTokenUpdate.Token] = inflight.epochAndRound
			}
		}
	}
	return latestTokenPriceUpdates
}

func (c *inflightCommitReportsContainer) reset(lggr logger.Logger) {
	lggr.Infow("Inflight report reset")
	c.locker.Lock()
	defer c.locker.Unlock()
	c.inFlight = make(map[[32]byte]inflightReport)
	c.inFlightPriceUpdates = []inflightPriceUpdate{}
}

func (c *inflightCommitReportsContainer) expire(lggr logger.Logger) {
	c.locker.Lock()
	defer c.locker.Unlock()
	// Reap any expired entries from inflight.
	for root, inFlightReport := range c.inFlight {
		if time.Since(inFlightReport.createdAt) > c.cacheExpiry {
			// Happy path: inflight report was successfully transmitted onchain, we remove it from inflight and onchain state reflects inflight.
			// Sad path: inflight report reverts onchain, we remove it from inflight, onchain state does not reflect the chains, so we retry.
			lggr.Infow("Inflight report expired", "rootOfRoots", hexutil.Encode(inFlightReport.report.MerkleRoot[:]))
			delete(c.inFlight, root)
		}
	}

	lggr.Infow("Inflight expire with price count", "count", len(c.inFlightPriceUpdates))

	var stillInflight []inflightPriceUpdate
	for _, inFlightFeeUpdate := range c.inFlightPriceUpdates {
		timeSinceUpdate := time.Since(inFlightFeeUpdate.createdAt)
		// If time passed since the price update is greater than multiplier * cache expiry, we remove it from the inflight list.
		if timeSinceUpdate > c.cacheExpiry*priceExpiryMultiplier {
			// Happy path: inflight report was successfully transmitted onchain, we remove it from inflight and onchain state reflects inflight.
			// Sad path: inflight report reverts onchain, we remove it from inflight, onchain state does not reflect the chains, so we retry.
			lggr.Infow("Inflight price update expired", "gasPrices", inFlightFeeUpdate.gasPrices, "tokenPrices", inFlightFeeUpdate.tokenPrices)
		} else {
			// If the update is still valid, we keep it in the inflight list.
			stillInflight = append(stillInflight, inFlightFeeUpdate)
		}
	}
	c.inFlightPriceUpdates = stillInflight
}

func (c *inflightCommitReportsContainer) add(lggr logger.Logger, report ccipdata.CommitStoreReport, epochAndRound uint64) error {
	c.locker.Lock()
	defer c.locker.Unlock()

	if report.MerkleRoot != [32]byte{} {
		// Set new inflight ones as pending
		lggr.Infow("Adding to inflight report", "rootOfRoots", hexutil.Encode(report.MerkleRoot[:]))
		c.inFlight[report.MerkleRoot] = inflightReport{
			report:    report,
			createdAt: time.Now(),
		}
	}

	if len(report.GasPrices) != 0 || len(report.TokenPrices) != 0 {
		lggr.Infow("Adding to inflight fee updates", "gasPrices", report.GasPrices, "tokenPrices", report.TokenPrices)
		c.inFlightPriceUpdates = append(c.inFlightPriceUpdates, inflightPriceUpdate{
			gasPrices:     report.GasPrices,
			tokenPrices:   report.TokenPrices,
			createdAt:     time.Now(),
			epochAndRound: epochAndRound,
		})
	}
	return nil
}
