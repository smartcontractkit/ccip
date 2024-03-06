package cache

import (
	"context"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

//go:generate mockery --quiet --name ArmChainState --filename chain_state_mock.go --case=underscore
type ArmChainState interface {
	ValidateNotCursed(ctx context.Context) error
	ForceValidateNotCursed(ctx context.Context) error
}

const (
	refreshInterval = 20 * time.Second
	cacheKey        = "armChainState"
)

type armChainState struct {
	cache    *cache.Cache
	cacheKey string

	lggr        logger.Logger
	onRamp      ccipdata.OnRampReader
	commitStore ccipdata.CommitStoreReader
}

func NewArmChainState(
	lggr logger.Logger,
	onRamp ccipdata.OnRampReader,
	commitStore ccipdata.CommitStoreReader,
) ArmChainState {
	return &armChainState{
		cache:    cache.New(refreshInterval, 0),
		cacheKey: cacheKey,

		lggr:        lggr,
		onRamp:      onRamp,
		commitStore: commitStore,
	}
}

func newArmChainWithCustomEviction(
	lggr logger.Logger,
	onRamp ccipdata.OnRampReader,
	commitStore ccipdata.CommitStoreReader,
	eviction time.Duration,
) ArmChainState {
	return &armChainState{
		cache:    cache.New(eviction, 0),
		cacheKey: cacheKey,

		lggr:        lggr,
		onRamp:      onRamp,
		commitStore: commitStore,
	}
}

func (a armChainState) ValidateNotCursed(ctx context.Context) error {
	if err, found := a.cache.Get(a.cacheKey); found {
		if err != nil {
			return err.(error)
		}
		return nil
	}

	err := a.fetch(ctx)
	a.cache.Set(a.cacheKey, err, cache.DefaultExpiration)
	return err
}

func (a armChainState) ForceValidateNotCursed(ctx context.Context) error {
	err := a.fetch(ctx)
	a.cache.Set(a.cacheKey, err, cache.DefaultExpiration)
	return err
}

func (a armChainState) fetch(ctx context.Context) error {
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
		return err
	}

	if isDown || isCursed {
		a.lggr.Errorf("Source chain is cursed or CommitStore is down", "isDown", isDown, "isCursed", isCursed)
		return ccip.ErrChainPausedOrCursed
	}
	return nil
}
