package ccip

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
)

type PriceGetter interface {
	TokensPerFeeCoin(ctx context.Context, tokens []common.Address) (map[common.Address]*big.Int, error)
}

var _ PriceGetter = &priceGetter{}

type priceGetter struct {
	source        string
	runner        pipeline.Runner
	jobID         int32
	externalJobID uuid.UUID
	name          string
	lggr          logger.Logger
}

func NewPriceGetter(source string, runner pipeline.Runner, jobID int32, externalJobID uuid.UUID, name string, lggr logger.Logger) (*priceGetter, error) {
	_, err := pipeline.Parse(source)
	if err != nil {
		return nil, err
	}
	return &priceGetter{
		source:        source,
		runner:        runner,
		jobID:         jobID,
		externalJobID: externalJobID,
		name:          name,
		lggr:          lggr,
	}, nil
}

func (d *priceGetter) TokensPerFeeCoin(ctx context.Context, tokens []common.Address) (map[common.Address]*big.Int, error) {
	_, trrs, err := d.runner.ExecuteRun(ctx, pipeline.Spec{
		ID:           d.jobID,
		DotDagSource: d.source,
		CreatedAt:    time.Now(),
		JobID:        d.jobID,
		JobName:      d.name,
		JobType:      "",
	}, pipeline.NewVarsFrom(map[string]interface{}{}), d.lggr)
	if err != nil {
		return nil, err
	}
	finalResult := trrs.FinalResult(d.lggr)
	if finalResult.HasErrors() {
		return nil, errors.Errorf("error getting prices %v", finalResult.AllErrors)
	}
	if len(finalResult.Values) != 1 {
		return nil, errors.Errorf("invalid number of price results, expected 1 got %v", len(finalResult.Values))
	}
	prices, ok := finalResult.Values[0].(map[string]interface{})
	if !ok {
		return nil, errors.Errorf("expected map output of price pipeline, got %T", finalResult.Values[0])
	}
	priceMap := make(map[common.Address]*big.Int)
	for addr, priceI := range prices {
		addr := common.HexToAddress(addr)
		priceStr, _ := priceI.(string)
		priceBigInt, _ := big.NewInt(0).SetString(priceStr, 10)
		priceMap[addr] = priceBigInt
	}
	// The mapping of token address to source of token price has to live offchain.
	// Best we can do is sanity check that the token price spec covers all our desired execution token prices.
	for _, token := range tokens {
		if _, ok = priceMap[token]; !ok {
			return nil, errors.Errorf("missing token %s from tokensForFeeCoin spec", token)
		}
	}
	return priceMap, nil
}

var _ PriceGetter = &fakePriceGetter{}

type fakePriceGetter struct{}

// TokensPerFeeCoin should fetch the price of an asset on the destination chain.
func (d fakePriceGetter) TokensPerFeeCoin(ctx context.Context, tokens []common.Address) (map[common.Address]*big.Int, error) {
	// Just returns a juels/eth value for all tokens.
	// As the feed is in wei/link and not juels/eth we need to transform it
	//  0.005 eth/link or 5e15 wei/link or 2e20 juels/eth
	weiPerLink := big.NewInt(5e15)
	precision := big.NewInt(0)
	// 1e18 * 1e18 = 1e36
	precision.SetString("1000000000000000000000000000000000000", 10)
	juelsPerFeeCoin := big.NewInt(0).Div(precision, weiPerLink)
	tokenPrices := make(map[common.Address]*big.Int)
	for _, token := range tokens {
		tokenPrices[token] = juelsPerFeeCoin
	}
	return tokenPrices, nil
}
