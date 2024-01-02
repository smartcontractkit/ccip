package tokenprice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	utilsbig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/token_price_ocr"
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
	contract       TokenPriceContract
}

// Close implements ocr3types.ReportingPlugin.
func (t *tokenPricePlugin) Close() error {
	return nil
}

// Observation implements ocr3types.ReportingPlugin.
func (t *tokenPricePlugin) Observation(ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query) (types.Observation, error) {
	// Fetch the tokens that need to have their price posted
	tokensNeeded, err := t.infoer.EnabledTokens(ctx)
	if err != nil {
		return nil, err
	}

	// Fetch the symbols for each of the enabled tokens
	// and map the symbol of the token to its address
	symbols, err := t.infoer.Symbols(ctx, tokensNeeded)
	if err != nil {
		return nil, err
	}

	if len(symbols) != len(tokensNeeded) {
		return nil, fmt.Errorf(
			"number of symbols (%d) does not match number of tokens (%d)",
			len(symbols), len(tokensNeeded))
	}

	enabledTokens := make(map[string]common.Address)
	for i, symbol := range symbols {
		enabledTokens[symbol] = tokensNeeded[i]
	}

	// Fetch the latest prices from the price service as well
	// as the prices that are posted onchain.
	prices, err := t.client.LatestPrices(ctx, symbols)
	if err != nil {
		return nil, err
	}

	onchainPrices, err := t.infoer.OnchainTokenPrices(ctx, tokensNeeded)
	if err != nil {
		return nil, err
	}

	// Get the timestamps for the last time each token price was updated.
	_, timestamps, err := t.contract.GetTokenPriceUpdates(ctx, tokensNeeded)
	if err != nil {
		return nil, err
	}

	heartbeats := make(map[string]int64)
	for i, symbol := range symbols {
		heartbeats[symbol] = timestamps[i]
	}

	o := Observation{
		TokenPrices:        prices,
		OnchainTokenPrices: onchainPrices,
		EnabledTokens:      enabledTokens,
		Heartbeats:         heartbeats,
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

	// obtain quorum on observed enabled tokens
	type pair struct {
		symbol string
		token  common.Address
	}
	var (
		tokenCounts = make(map[pair]int)
	)
	for _, o := range obs {
		for symb, token := range o.EnabledTokens {
			tokenCounts[pair{symb, token}]++
		}
	}

	quorumTokens := make(map[string]common.Address)
	for p, count := range tokenCounts {
		if count >= (t.f + 1) {
			quorumTokens[p.symbol] = p.token
		}
	}

	if len(quorumTokens) == 0 {
		return nil, errors.New("no quorum on tokens")
	}

	// medianize new prices for each token
	medians := make(map[string]*utilsbig.Big)
	for symb, _ := range quorumTokens {
		var prices []*utilsbig.Big
		for _, o := range obs {
			price, ok := o.TokenPrices[symb]
			if !ok {
				continue
			}
			prices = append(prices, price)
		}
		if len(prices) == 0 {
			// Should be impossible
			return nil, fmt.Errorf("no prices on quorum token %s", symb)
		}
		medians[symb] = median(prices)
	}

	// medianize onchain prices for each token
	onchainMedians := make(map[string]*utilsbig.Big)
	for symb, _ := range quorumTokens {
		var prices []*utilsbig.Big
		for _, o := range obs {
			price, ok := o.OnchainTokenPrices[symb]
			if !ok {
				continue
			}
			prices = append(prices, price)
		}
		// its possible there are no onchain prices if price was never posted onchain
		if len(prices) > 0 {
			onchainMedians[symb] = median(prices)
		}
	}

	// medianize last updated timestamps for each token
	medianTimestamps := make(map[string]int64)
	for symb, _ := range quorumTokens {
		var timestamps []int64
		for _, o := range obs {
			ts, ok := o.Heartbeats[symb]
			if !ok {
				continue
			}
			timestamps = append(timestamps, ts)
		}
		// its possible there is no heartbeat if price was never posted onchain
		if len(timestamps) > 0 {
			medianTimestamps[symb] = median(func() []*utilsbig.Big {
				var ret []*utilsbig.Big
				for _, ts := range timestamps {
					ret = append(ret, utilsbig.NewI(ts))
				}
				return ret
			}()).Int64()
		}
	}

	out := Outcome{
		MedianTokenPrices:        medians,
		MedianOnchainTokenPrices: onchainMedians,
		MedianTimestamps:         medianTimestamps,
		QuorumTokens:             quorumTokens,
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
func (t *tokenPricePlugin) Reports(seqNr uint64, outcome ocr3types.Outcome) ([]ocr3types.ReportWithInfo[meta], error) {
	// determine deviation from the onchain median
	// if deviation is exceeded, then we want to report that
	// if deviation is not exceeded, but the heartbeat period is exceeded, then we want to report that
	// otherwise, we don't want to report anything
	var o Outcome
	err := json.Unmarshal(outcome, &o)
	if err != nil {
		return nil, err
	}

	// Only want to append a new report if the size limit of the current
	// report is exceeded. Otherwise, we want to update the current report.
	var reports []ocr3types.ReportWithInfo[meta]
	var currentReport token_price_ocr.TokenPriceOCRReport
	for symb, token := range o.QuorumTokens {
		// an offchain median must be present if the symbol is present
		offchainMedian := o.MedianTokenPrices[symb]
		// an onchain median may or may not be present, if the token price was never posted
		// onchain.
		onchainMedian, onchainMedianPresent := o.MedianOnchainTokenPrices[symb]
		if !onchainMedianPresent {
			if len(currentReport.PriceUpdates) >= int(t.offchainConfig.MaxPriceUpdates) {
				encoded, err := ABIEncodeReportData(currentReport)
				if err != nil {
					return nil, err
				}
				reports = append(reports, ocr3types.ReportWithInfo[meta]{
					Report: encoded,
				})
				currentReport = token_price_ocr.TokenPriceOCRReport{}
			} else {
				currentReport.PriceUpdates = append(currentReport.PriceUpdates, token_price_ocr.InternalTokenPriceUpdate{
					SourceToken: token,
					UsdPerToken: offchainMedian.ToInt(),
				})
			}
			continue
		}
		// onchain median is present, so compare that to the offchain median
		// if there is sufficient deviation then update the price onchain.
		if deviates(t.offchainConfig.ReportPPB, onchainMedian.ToInt(), offchainMedian.ToInt()) {
			if len(currentReport.PriceUpdates) >= int(t.offchainConfig.MaxPriceUpdates) {
				encoded, err := ABIEncodeReportData(currentReport)
				if err != nil {
					return nil, err
				}
				reports = append(reports, ocr3types.ReportWithInfo[meta]{
					Report: encoded,
				})
				currentReport = token_price_ocr.TokenPriceOCRReport{}
			} else {
				currentReport.PriceUpdates = append(currentReport.PriceUpdates, token_price_ocr.InternalTokenPriceUpdate{
					SourceToken: token,
					UsdPerToken: offchainMedian.ToInt(),
				})
			}
			continue
		}
		// onchain median is present, deviation is not sufficient
		// check if the heartbeat period has passed, and if it has, update
		// the price onchain.
		heartbeatTs, heartbeatPresent := o.MedianTimestamps[symb]
		if heartbeatPresent && (heartbeatTs+int64(t.offchainConfig.DeltaC.Seconds())) > time.Now().UTC().Unix() {
			if len(currentReport.PriceUpdates) >= int(t.offchainConfig.MaxPriceUpdates) {
				encoded, err := ABIEncodeReportData(currentReport)
				if err != nil {
					return nil, err
				}
				reports = append(reports, ocr3types.ReportWithInfo[meta]{
					Report: encoded,
				})
				currentReport = token_price_ocr.TokenPriceOCRReport{}
			} else {
				currentReport.PriceUpdates = append(currentReport.PriceUpdates, token_price_ocr.InternalTokenPriceUpdate{
					SourceToken: token,
					UsdPerToken: offchainMedian.ToInt(),
				})
			}
		}
	}

	return reports, nil
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

func deviates(thresholdPPB uint64, old *big.Int, new *big.Int) bool {
	if old.Cmp(big.NewInt(0)) == 0 {
		if new.Cmp(big.NewInt(0)) == 0 { //nolint:gosimple
			return false // Both values are zero; no deviation
		}
		return true // Any deviation from 0 is significant
	}
	// ||new - old|| / ||old||, approximated by a float
	change := &big.Rat{}
	change.SetFrac(big.NewInt(0).Sub(new, old), old)
	change.Abs(change)
	threshold := &big.Rat{}
	threshold.SetFrac(
		(&big.Int{}).SetUint64(thresholdPPB),
		(&big.Int{}).SetUint64(1e9),
	)
	return change.Cmp(threshold) >= 0
}
