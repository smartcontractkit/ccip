package pricegetter

import (
	"context"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/parseutil"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
)

var _ PriceGetter = &PipelineGetter{}

type PipelineGetter struct {
	source        string
	runner        pipeline.Runner
	jobID         int32
	externalJobID uuid.UUID
	name          string
	lggr          logger.Logger
}

func NewPipelineGetter(source string, runner pipeline.Runner, jobID int32, externalJobID uuid.UUID, name string, lggr logger.Logger) (*PipelineGetter, error) {
	_, err := pipeline.Parse(source)
	if err != nil {
		return nil, err
	}

	return &PipelineGetter{
		source:        source,
		runner:        runner,
		jobID:         jobID,
		externalJobID: externalJobID,
		name:          name,
		lggr:          lggr,
	}, nil
}

func (d *PipelineGetter) TokenPricesUSD(ctx context.Context, tokens []common.Address) (map[common.Address]TokenPriceResult, error) {
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

	providedTokensSet := mapset.NewSet[common.Address](tokens...)
	tokenPrices := make(map[common.Address]TokenPriceResult)
	for tokenAddressStr, rawPrice := range prices {
		tokenAddress := common.HexToAddress(tokenAddressStr)

		if providedTokensSet.Contains(tokenAddress) {
			castedPrice, err1 := parseutil.ParseBigIntFromAny(rawPrice)

			tokenPrices[tokenAddress] = TokenPriceResult{
				Price: castedPrice,
				Error: err1,
			}
		}
	}

	// Make sure that not found tokens are also returned with an error.
	for _, token := range tokens {
		if _, exist := tokenPrices[token]; exist {
			continue
		}
		tokenPrices[token] = TokenPriceResult{
			Error: errors.Errorf("token price not found in spec"),
		}
	}

	return tokenPrices, nil
}
