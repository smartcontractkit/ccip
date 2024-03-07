package cache

import (
	"context"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

//go:generate mockery --quiet --name ChainHealthcheck --filename chain_health_mock.go --case=underscore
type ChainHealthcheck interface {
	IsHealthy(ctx context.Context) (bool, error)
	ForceIsHealthy(ctx context.Context) (bool, error)
}

const (
	defaultRmnStatusDuration    = 20 * time.Second
	defaultGlobalStatusDuration = 30 * time.Minute

	globalStatusKey = "globalStatus"
	rmnStatusKey    = "rmnCurseCheck"
)

type chainHealthcheck struct {
	cache                  *cache.Cache
	globalStatusKey        string
	rmnStatusKey           string
	globalStatusExpiration time.Duration
	rmnStatusExpiration    time.Duration

	lggr        logger.Logger
	onRamp      ccipdata.OnRampReader
	commitStore ccipdata.CommitStoreReader
}

func NewChainHealthcheck(
	lggr logger.Logger,
	onRamp ccipdata.OnRampReader,
	commitStore ccipdata.CommitStoreReader,
) *chainHealthcheck {
	return &chainHealthcheck{
		cache:                  cache.New(defaultRmnStatusDuration, 0),
		globalStatusKey:        globalStatusKey,
		rmnStatusKey:           rmnStatusKey,
		globalStatusExpiration: defaultGlobalStatusDuration,
		rmnStatusExpiration:    defaultRmnStatusDuration,

		lggr:        lggr,
		onRamp:      onRamp,
		commitStore: commitStore,
	}
}

func newChainHealthcheckWithCustomEviction(
	lggr logger.Logger,
	onRamp ccipdata.OnRampReader,
	commitStore ccipdata.CommitStoreReader,
	globalStatusDuration time.Duration,
	rmnStatusDuration time.Duration,
) *chainHealthcheck {
	return &chainHealthcheck{
		cache:                  cache.New(rmnStatusDuration, 0),
		rmnStatusKey:           rmnStatusKey,
		globalStatusKey:        globalStatusKey,
		globalStatusExpiration: globalStatusDuration,
		rmnStatusExpiration:    rmnStatusDuration,

		lggr:        lggr,
		onRamp:      onRamp,
		commitStore: commitStore,
	}
}

func (a chainHealthcheck) IsHealthy(ctx context.Context) (bool, error) {
	return a.isHealthy(ctx, false)
}

func (a chainHealthcheck) ForceIsHealthy(ctx context.Context) (bool, error) {
	return a.isHealthy(ctx, true)
}

func (a chainHealthcheck) isHealthy(ctx context.Context, forceRmnRefresh bool) (bool, error) {
	// Verify if flag is raised to indicate that the chain is not healthy
	// If set to false then immediately return false without checking the chain
	if healthy, found := a.cache.Get(a.globalStatusKey); found && !healthy.(bool) {
		return false, nil
	}

	if healthy, err := a.checkIfReadersAreHealthy(ctx); err != nil {
		return false, err
	} else if !healthy {
		a.cache.Set(a.globalStatusKey, false, a.globalStatusExpiration)
		return healthy, nil
	}

	if healthy, err := a.checkIfRMNsAreHealthy(ctx, forceRmnRefresh); err != nil {
		return false, err
	} else if !healthy {
		a.cache.Set(a.globalStatusKey, false, a.globalStatusExpiration)
		return healthy, nil
	}
	return true, nil
}

// checkIfReadersAreHealthy checks if the source and destination chains are healthy by calling underlying LogPoller
// These calls are very cheap, because doesn't require any communication with database or RPC, so we don't have
// to cache the result of these calls.
func (a chainHealthcheck) checkIfReadersAreHealthy(ctx context.Context) (bool, error) {
	sourceChainHealthy, err := a.onRamp.IsSourceChainHealthy(ctx)
	if err != nil {
		return false, errors.Wrap(err, "onRamp IsSourceChainHealthy errored")
	}

	destChainHealthy, err := a.commitStore.IsDestChainHealthy(ctx)
	if err != nil {
		return false, errors.Wrap(err, "commitStore IsDestChainHealthy errored")
	}

	if !sourceChainHealthy || !destChainHealthy {
		a.lggr.Errorf("Source or destination chain is unhealthy", "sourceChainHealthy", sourceChainHealthy, "destChainHealthy", destChainHealthy)
	}
	return sourceChainHealthy && destChainHealthy, nil
}

func (a chainHealthcheck) checkIfRMNsAreHealthy(ctx context.Context, forceFetch bool) (bool, error) {
	if !forceFetch {
		if healthy, found := a.cache.Get(a.rmnStatusKey); found {
			return healthy.(bool), nil
		}
	}

	healthy, err := a.fetchRMNCurseState(ctx)
	if err != nil {
		return false, err
	}

	a.cache.Set(a.rmnStatusKey, healthy, a.rmnStatusExpiration)
	return healthy, nil
}

func (a chainHealthcheck) fetchRMNCurseState(ctx context.Context) (bool, error) {
	var (
		eg       = new(errgroup.Group)
		isDown   bool
		isCursed bool
	)

	eg.Go(func() error {
		var err error
		isDown, err = a.commitStore.IsDown(ctx)
		if err != nil {
			return errors.Wrap(err, "commitStore isDown check errored")
		}
		return nil
	})

	eg.Go(func() error {
		var err error
		isCursed, err = a.onRamp.IsSourceCursed(ctx)
		if err != nil {
			return errors.Wrap(err, "onRamp isSourceCursed errored")
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return false, err
	}

	if isDown || isCursed {
		a.lggr.Errorw("Source chain is cursed or CommitStore is down", "isDown", isDown, "isCursed", isCursed)
		return false, nil
	}
	return true, nil
}
