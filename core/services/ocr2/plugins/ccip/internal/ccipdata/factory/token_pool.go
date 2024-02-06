package factory

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	type_and_version "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/type_and_version_interface_wrapper"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_2_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_4_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

var (
	typeAndVersionABI = abihelpers.MustParseABI(type_and_version.TypeAndVersionInterfaceABI)
)

type TokenPoolFactory struct {
	lggr                logger.Logger
	remoteChainSelector uint64
	offRampAddress      common.Address
	lp                  logpoller.LogPoller
	evmBatchCaller      rpclib.EvmBatchCaller

	tokenPoolReaders  map[common.Address]ccipdata.TokenPoolReader
	tokenPoolReaderMu sync.RWMutex
}

//go:generate mockery --quiet --name TokenPoolFactoryInterface --filename token_pool_factory_mock.go --case=underscor
type TokenPoolFactoryInterface interface {
	GetInboundTokenPoolRateLimits(ctx context.Context, tokenPoolReaders []common.Address) ([]ccipdata.TokenBucketRateLimit, error)
}

var _ TokenPoolFactoryInterface = (*TokenPoolFactory)(nil)

func NewTokenPoolFactory(lggr logger.Logger, remoteChainSelector uint64, offRampAddress common.Address, evmBatchCaller rpclib.EvmBatchCaller, lp logpoller.LogPoller) TokenPoolFactory {
	return TokenPoolFactory{
		lggr:                lggr,
		remoteChainSelector: remoteChainSelector,
		offRampAddress:      offRampAddress,
		lp:                  lp,
		evmBatchCaller:      evmBatchCaller,
		tokenPoolReaders:    make(map[common.Address]ccipdata.TokenPoolReader),
	}
}

func (f *TokenPoolFactory) GetInboundTokenPoolRateLimits(ctx context.Context, tokenPools []common.Address) ([]ccipdata.TokenBucketRateLimit, error) {
	if len(tokenPools) == 0 {
		return []ccipdata.TokenBucketRateLimit{}, nil
	}

	err := f.loadTokenPools(ctx, tokenPools)
	if err != nil {
		return nil, err
	}

	tokenPoolReaders := make([]ccipdata.TokenPoolReader, 0, len(tokenPools))
	for _, poolAddress := range tokenPools {
		f.tokenPoolReaderMu.RLock()
		tokenPoolReader, exists := f.tokenPoolReaders[poolAddress]
		f.tokenPoolReaderMu.RUnlock()
		if !exists {
			return nil, fmt.Errorf("token pool %s not found", poolAddress.Hex())
		}
		tokenPoolReaders = append(tokenPoolReaders, tokenPoolReader)
	}

	evmCalls := make([]rpclib.EvmCall, 0, len(tokenPoolReaders))
	for _, poolReader := range tokenPoolReaders {
		switch v := poolReader.(type) {
		case *v1_2_0.TokenPool:
			call, err := v1_2_0.GetInboundTokenPoolRateLimitCall(v.Address(), v.OffRampAddress)
			if err != nil {
				return nil, fmt.Errorf("get inbound token pool rate limit call: %w", err)
			}
			evmCalls = append(evmCalls, call)
		case *v1_4_0.TokenPool:
			call, err := v1_4_0.GetInboundTokenPoolRateLimitCall(v.Address(), v.RemoteChainSelector)
			if err != nil {
				return nil, fmt.Errorf("get inbound token pool rate limit call: %w", err)
			}
			evmCalls = append(evmCalls, call)
		}
	}

	latestBlock, err := f.lp.LatestBlock(pg.WithParentCtx(ctx))
	if err != nil {
		return nil, fmt.Errorf("get latest block: %w", err)
	}

	results, err := f.evmBatchCaller.BatchCall(ctx, uint64(latestBlock.BlockNumber), evmCalls)
	if err != nil {
		return nil, fmt.Errorf("batch call limit: %w", err)
	}

	rateLimits, err := rpclib.ParseOutputs[ccipdata.TokenBucketRateLimit](results, func(d rpclib.DataAndErr) (ccipdata.TokenBucketRateLimit, error) {
		return rpclib.ParseOutput[ccipdata.TokenBucketRateLimit](d, 0)
	})
	if err != nil {
		return nil, fmt.Errorf("parse outputs: %w", err)
	}

	if len(rateLimits) != len(tokenPoolReaders) {
		return nil, fmt.Errorf("expected %d rate limits, got %d", len(tokenPoolReaders), len(rateLimits))
	}

	return rateLimits, nil
}

// loadTokenPools loads the token pools into the factory's cache
func (f *TokenPoolFactory) loadTokenPools(ctx context.Context, tokenPoolAddresses []common.Address) error {
	var missingTokens []common.Address

	f.tokenPoolReaderMu.RLock()
	for _, poolAddress := range tokenPoolAddresses {
		if _, exists := f.tokenPoolReaders[poolAddress]; !exists {
			missingTokens = append(missingTokens, poolAddress)
		}
	}
	f.tokenPoolReaderMu.RUnlock()

	// Only continue if there are missing tokens
	if len(missingTokens) == 0 {
		return nil
	}

	typeAndVersions, err := getBatchedTypeAndVersion(ctx, f.lp, f.evmBatchCaller, missingTokens)
	if err != nil {
		return err
	}

	f.tokenPoolReaderMu.Lock()
	defer f.tokenPoolReaderMu.Unlock()
	for i, tokenPoolAddress := range missingTokens {
		typeAndVersion := typeAndVersions[i]
		poolType, version, err := ccipconfig.ParseTypeAndVersion(typeAndVersion)
		if err != nil {
			return err
		}
		switch version {
		case ccipdata.V1_0_0, ccipdata.V1_1_0, ccipdata.V1_2_0:
			f.tokenPoolReaders[tokenPoolAddress] = v1_2_0.NewTokenPool(poolType, tokenPoolAddress, f.offRampAddress)
		case ccipdata.V1_4_0:
			f.tokenPoolReaders[tokenPoolAddress] = v1_4_0.NewTokenPool(poolType, tokenPoolAddress, f.remoteChainSelector)
		default:
			return fmt.Errorf("unsupported token pool version %v", version)
		}
	}
	return nil
}

func getBatchedTypeAndVersion(ctx context.Context, lp logpoller.LogPoller, evmBatchCaller rpclib.EvmBatchCaller, poolAddresses []common.Address) ([]string, error) {
	var evmCalls []rpclib.EvmCall

	for _, poolAddress := range poolAddresses {
		// Add the typeAndVersion call to the batch
		evmCalls = append(evmCalls, rpclib.NewEvmCall(
			typeAndVersionABI,
			"typeAndVersion",
			poolAddress,
		))
	}

	latestBlock, err := lp.LatestBlock(pg.WithParentCtx(ctx))
	if err != nil {
		return nil, fmt.Errorf("get latest block: %w", err)
	}

	results, err := evmBatchCaller.BatchCall(ctx, uint64(latestBlock.FinalizedBlockNumber), evmCalls)
	if err != nil {
		return nil, fmt.Errorf("batch call limit: %w", err)
	}

	typeAndVersions, err := rpclib.ParseOutputs[string](results, func(d rpclib.DataAndErr) (string, error) {
		return rpclib.ParseOutput[string](d, 0)
	})
	if err != nil {
		return nil, fmt.Errorf("parse outputs: %w", err)
	}
	return typeAndVersions, nil
}
