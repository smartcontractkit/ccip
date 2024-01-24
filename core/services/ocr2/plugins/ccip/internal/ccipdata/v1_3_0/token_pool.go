package v1_3_0

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
)

var (
	poolABI = abihelpers.MustParseABI(burn_mint_token_pool.BurnMintTokenPoolABI)
)

var _ ccipdata.TokenPoolReader = &TokenPool{}

type TokenPool struct {
	addr                common.Address
	remoteChainSelector uint64
	poolType            string
}

func NewTokenPool(poolType string, addr common.Address, remoteChainSelector uint64) *TokenPool {
	return &TokenPool{
		addr:                addr,
		remoteChainSelector: remoteChainSelector,
		poolType:            poolType,
	}
}

func (p *TokenPool) Address() common.Address {
	return p.addr
}

func (p *TokenPool) GetInboundTokenPoolRateLimitCall() (rpclib.EvmCall, error) {
	return rpclib.NewEvmCall(
		poolABI,
		"getCurrentInboundRateLimiterState",
		p.addr,
		p.remoteChainSelector,
	), nil
}
