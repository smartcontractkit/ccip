package factory

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	type_and_version "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/type_and_version_interface_wrapper"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_2_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_3_0"
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
	ec                  client.Client
	lp                  logpoller.LogPoller
	evmBatchCaller      rpclib.EvmBatchCaller
}

//go:generate mockery --quiet --name TokenPoolFactoryInterface --filename token_pool_factory_mock.go --case=underscor
type TokenPoolFactoryInterface interface {
	NewTokenPools(ctx context.Context, tokenPoolAddresses []common.Address) ([]ccipdata.TokenPoolReader, error)
}

var _ TokenPoolFactoryInterface = (*TokenPoolFactory)(nil)

func NewTokenPoolFactory(lggr logger.Logger, remoteChainSelector uint64, offRampAddress common.Address, ec client.Client, lp logpoller.LogPoller) TokenPoolFactory {
	return TokenPoolFactory{
		lggr:                lggr,
		remoteChainSelector: remoteChainSelector,
		offRampAddress:      offRampAddress,
		ec:                  ec,
		lp:                  lp,
		evmBatchCaller:      rpclib.NewDynamicLimitedBatchCaller(lggr, ec, rpclib.DefaultRpcBatchSizeLimit, rpclib.DefaultRpcBatchBackOffMultiplier),
	}
}

// NewTokenPools returns a slice of token pool readers for the given token pool addresses.
func (f TokenPoolFactory) NewTokenPools(ctx context.Context, tokenPoolAddresses []common.Address) ([]ccipdata.TokenPoolReader, error) {
	// Return early to avoid the logPoller call
	if len(tokenPoolAddresses) == 0 {
		return []ccipdata.TokenPoolReader{}, nil
	}

	var evmCalls []rpclib.EvmCall
	for _, poolAddress := range tokenPoolAddresses {
		evmCalls = append(evmCalls, rpclib.NewEvmCall(
			typeAndVersionABI,
			"typeAndVersion",
			poolAddress,
		))
	}

	latestBlock, err := f.lp.LatestBlock(pg.WithParentCtx(ctx))
	if err != nil {
		return nil, fmt.Errorf("get latest block: %w", err)
	}

	results, err := f.evmBatchCaller.BatchCall(ctx, uint64(latestBlock.FinalizedBlockNumber), evmCalls)
	if err != nil {
		return nil, fmt.Errorf("batch call limit: %w", err)
	}

	typeAndVersions, err := rpclib.ParseOutputs[string](results, func(d rpclib.DataAndErr) (string, error) {
		return rpclib.ParseOutput[string](d, 0)
	})

	var tokenPoolReaders []ccipdata.TokenPoolReader

	for i, tokenPoolAddress := range tokenPoolAddresses {
		typeAndVersion := typeAndVersions[i]
		poolType, version, err := ccipconfig.ParseTypeAndVersion(typeAndVersion)
		if err != nil {
			return nil, err
		}
		switch version {
		case ccipdata.V1_0_0, ccipdata.V1_1_0, ccipdata.V1_2_0:
			tokenPoolReaders = append(tokenPoolReaders, v1_2_0.NewTokenPool(poolType, tokenPoolAddress, f.offRampAddress, f.ec, f.lp, f.evmBatchCaller))
		case ccipdata.V1_3_0:
			tokenPoolReaders = append(tokenPoolReaders, v1_3_0.NewTokenPool(poolType, tokenPoolAddress, f.remoteChainSelector, f.ec, f.lp, f.evmBatchCaller))
		default:
			return nil, fmt.Errorf("unsupported token pool version %v", version)
		}
	}
	return tokenPoolReaders, nil
}
