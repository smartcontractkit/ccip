package v1_2_0

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/burn_mint_token_pool_1_2_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

var (
	poolABI = abihelpers.MustParseABI(burn_mint_token_pool_1_2_0.BurnMintTokenPoolABI)
)

var _ ccipdata.TokenPoolReader = &TokenPool{}

type TokenPool struct {
	addr           common.Address
	offRampAddress common.Address
	poolType       string
	lp             logpoller.LogPoller
	evmBatchCaller rpclib.EvmBatchCaller
}

func NewTokenPool(poolType string, addr common.Address, offRampAddress common.Address, lp logpoller.LogPoller, evmBatchCaller rpclib.EvmBatchCaller) *TokenPool {
	return &TokenPool{
		addr:           addr,
		offRampAddress: offRampAddress,
		poolType:       poolType,
		lp:             lp,
		evmBatchCaller: evmBatchCaller,
	}
}

func (p *TokenPool) Address() common.Address {
	return p.addr
}

func (p *TokenPool) Type() string {
	return p.poolType
}

func (p *TokenPool) GetInboundTokenPoolRateLimits(ctx context.Context, tokenPoolReaders []ccipdata.TokenPoolReader) ([]ccipdata.TokenBucketRateLimit, error) {
	if len(tokenPoolReaders) == 0 {
		return []ccipdata.TokenBucketRateLimit{}, nil
	}

	evmCalls := make([]rpclib.EvmCall, 0, len(tokenPoolReaders))
	for _, poolReader := range tokenPoolReaders {
		call, err := poolReader.GetInboundTokenPoolRateLimitCall()
		if err != nil {
			return nil, fmt.Errorf("get inbound token pool rate limit call: %w", err)
		}
		evmCalls = append(evmCalls, call)
	}

	latestBlock, err := p.lp.LatestBlock(pg.WithParentCtx(ctx))
	if err != nil {
		return nil, fmt.Errorf("get latest block: %w", err)
	}

	results, err := p.evmBatchCaller.BatchCall(ctx, uint64(latestBlock.BlockNumber), evmCalls)
	if err != nil {
		return nil, fmt.Errorf("batch call limit: %w", err)
	}

	rateLimits, err := rpclib.ParseOutputs[ccipdata.TokenBucketRateLimit](results, func(d rpclib.DataAndErr) (ccipdata.TokenBucketRateLimit, error) {
		return rpclib.ParseOutput[ccipdata.TokenBucketRateLimit](d, 0)
	})
	if err != nil {
		return nil, fmt.Errorf("parse outputs: %w", err)
	}

	return rateLimits, nil
}

func (p *TokenPool) GetInboundTokenPoolRateLimitCall() (rpclib.EvmCall, error) {
	return rpclib.NewEvmCall(
		poolABI,
		"currentOffRampRateLimiterState",
		p.addr,
		p.offRampAddress,
	), nil
}
