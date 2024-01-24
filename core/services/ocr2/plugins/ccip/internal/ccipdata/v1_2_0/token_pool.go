package v1_2_0

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/burn_mint_token_pool_1_2_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
)

var (
	poolABI = abihelpers.MustParseABI(burn_mint_token_pool_1_2_0.BurnMintTokenPoolABI)
)

var _ ccipdata.TokenPoolReader = &TokenPool{}

type TokenPool struct {
	addr           common.Address
	offRampAddress common.Address
	poolType       string
}

func NewTokenPool(poolType string, addr common.Address, offRampAddress common.Address) *TokenPool {
	return &TokenPool{
		addr:           addr,
		offRampAddress: offRampAddress,
		poolType:       poolType,
	}
}

func (p *TokenPool) Address() common.Address {
	return p.addr
}

func (p *TokenPool) GetInboundTokenPoolRateLimitCall() (rpclib.EvmCall, error) {
	return rpclib.NewEvmCall(
		poolABI,
		"currentOffRampRateLimiterState",
		p.addr,
		p.offRampAddress,
	), nil
}
