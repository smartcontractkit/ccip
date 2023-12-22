package tokenprice

import (
	"context"
	"encoding/json"
	"sort"

	utilsbig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

type meta struct {
}

type tokenPricePlugin struct {
	infoer         TokenInfoer
	client         PriceServiceClient
	f              int
	offchainConfig OffchainConfig
}

// Close implements ocr3types.ReportingPlugin.
func (t *tokenPricePlugin) Close() error {
	panic("unimplemented")
}

// Observation implements ocr3types.ReportingPlugin.
func (t *tokenPricePlugin) Observation(ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query) (types.Observation, error) {
	tokensNeeded, err := t.infoer.EnabledTokens(ctx)
	if err != nil {
		return nil, err
	}

	var symbols []string
	for _, token := range tokensNeeded {
		symbol, err := t.infoer.Symbol(ctx, token)
		if err != nil {
			return nil, err
		}
		symbols = append(symbols, symbol)
	}

	prices, err := t.client.LatestPrices(ctx, symbols)
	if err != nil {
		return nil, err
	}

	onchainPrices, err := t.infoer.OnchainPrices(ctx, tokensNeeded)
	if err != nil {
		return nil, err
	}

	o := Observation{
		Prices:        prices,
		OnchainPrices: onchainPrices,
		EnabledTokens: symbols,
	}
	marshaled, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	return types.Observation(marshaled), nil
}

// ObservationQuorum implements ocr3types.ReportingPlugin.
func (*tokenPricePlugin) ObservationQuorum(outctx ocr3types.OutcomeContext, query types.Query) (ocr3types.Quorum, error) {
	return ocr3types.QuorumFPlusOne, nil
}

// Outcome implements ocr3types.ReportingPlugin.
func (t *tokenPricePlugin) Outcome(outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation) (ocr3types.Outcome, error) {
	// parse observations from other nodes
	var obs []Observation
	for _, ao := range aos {
		var o Observation
		err := json.Unmarshal(ao.Observation, &o)
		if err != nil {
			return nil, err
		}
		obs = append(obs, o)
	}

	// obtrain quorum on observed tokens
	tokenSymbolCounts := make(map[string]int)
	for _, o := range obs {
		for _, token := range o.EnabledTokens {
			tokenSymbolCounts[token]++
		}
	}

	var quorumTokens []string
	for token, count := range tokenSymbolCounts {
		if count >= (t.f + 1) {
			quorumTokens = append(quorumTokens, token)
		}
	}

	// medianize new prices for each token
	medians := make(map[string]*utilsbig.Big)
	for _, token := range quorumTokens {
		var prices []*utilsbig.Big
		for _, o := range obs {
			price, ok := o.Prices[token]
			if !ok {
				continue
			}
			prices = append(prices, price)
		}
		medians[token] = median(prices)
	}

	// medianize onchain prices for each token
	onchainMedians := make(map[string]*utilsbig.Big)
	for _, token := range quorumTokens {
		var prices []*utilsbig.Big
		for _, o := range obs {
			price, ok := o.OnchainPrices[token]
			if !ok {
				continue
			}
			prices = append(prices, price)
		}
		onchainMedians[token] = median(prices)
	}

	out := Outcome{
		MedianPrices:        medians,
		MedianOnchainPrices: onchainMedians,
	}
	outBytes, err := json.Marshal(out)
	if err != nil {
		return nil, err
	}

	return ocr3types.Outcome(outBytes), nil
}

// Query implements ocr3types.ReportingPlugin.
func (*tokenPricePlugin) Query(ctx context.Context, outctx ocr3types.OutcomeContext) (types.Query, error) {
	return types.Query{}, nil
}

// Reports implements ocr3types.ReportingPlugin.
func (*tokenPricePlugin) Reports(seqNr uint64, outcome ocr3types.Outcome) ([]ocr3types.ReportWithInfo[meta], error) {
	panic("unimplemented")
}

// ShouldAcceptAttestedReport implements ocr3types.ReportingPlugin.
func (*tokenPricePlugin) ShouldAcceptAttestedReport(context.Context, uint64, ocr3types.ReportWithInfo[meta]) (bool, error) {
	// TODO: figure out
	return true, nil
}

// ShouldTransmitAcceptedReport implements ocr3types.ReportingPlugin.
func (*tokenPricePlugin) ShouldTransmitAcceptedReport(context.Context, uint64, ocr3types.ReportWithInfo[meta]) (bool, error) {
	// TODO: figure out
	return true, nil
}

// ValidateObservation implements ocr3types.ReportingPlugin.
func (*tokenPricePlugin) ValidateObservation(outctx ocr3types.OutcomeContext, query types.Query, ao types.AttributedObservation) error {
	return json.Unmarshal(ao.Observation, &Observation{})
}

var _ ocr3types.ReportingPlugin[meta] = (*tokenPricePlugin)(nil)

func median(vals []*utilsbig.Big) *utilsbig.Big {
	if len(vals) == 0 {
		return nil
	}

	valsCopy := make([]*utilsbig.Big, len(vals))
	copy(valsCopy[:], vals[:])
	sort.Slice(valsCopy, func(i, j int) bool {
		return valsCopy[i].Cmp(valsCopy[j]) == -1
	})
	return valsCopy[len(valsCopy)/2]
}
