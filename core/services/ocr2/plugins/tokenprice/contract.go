package tokenprice

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	clcommontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	clcommontokenprice "github.com/smartcontractkit/chainlink-common/pkg/types/tokenprice"
	utilsbig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
)

func newTokenPriceContract(chainReader clcommontypes.ChainReader, address common.Address) clcommontokenprice.TokenPriceContract {
	contract := clcommontypes.BoundContract{Address: address.String(), Name: "tokenprice", Pending: true}
	return &tokenPriceContract{chainReader, contract}
}

type tokenPriceContract struct {
	chainReader clcommontypes.ChainReader
	contract    clcommontypes.BoundContract
}

var _ clcommontokenprice.TokenPriceContract = &tokenPriceContract{}

func (t *tokenPriceContract) GetTokenPriceUpdates(ctx context.Context, addresses []string) (
	prices []*big.Int,
	timestamps []int64,
	err error,
) {
	var resp struct {
		Prices []*utilsbig.Big
		Times  []int64
	}
	err = t.chainReader.GetLatestValue(ctx, t.contract, "GetTokenPriceUpdates", addresses, &resp)
	if err != nil {
		return
	}
	var bigs []*big.Int
	for _, p := range resp.Prices {
		bigs = append(bigs, p.ToInt())
	}
	return bigs, resp.Times, nil
}
