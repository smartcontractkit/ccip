package prices

import (
	"context"
	"fmt"
	"math/big"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcalc"
)

type DAGasPriceEstimator struct {
	execEstimator       ExecGasPriceEstimator
	priceEncodingLength uint
}

func (g DAGasPriceEstimator) GetGasPrice(ctx context.Context) (GasPrice, error) {
	execGasPrice, err := g.execEstimator.GetGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	var gasPrice *big.Int = execGasPrice
	if gasPrice.BitLen() > int(g.priceEncodingLength) {
		return nil, fmt.Errorf("native gas price exceeded max range %+v", gasPrice)
	}

	if l1Oracle := g.execEstimator.estimator.L1Oracle(); l1Oracle != nil {
		daGasPriceWei, err := l1Oracle.GasPrice(ctx)
		if err != nil {
			return nil, err
		}

		if daGasPrice := daGasPriceWei.ToInt(); daGasPrice.Cmp(big.NewInt(0)) > 0 {
			if daGasPrice.BitLen() > int(g.priceEncodingLength) {
				return nil, fmt.Errorf("data availability gas price exceeded max range %+v", daGasPrice)
			}

			daGasPrice := new(big.Int).Lsh(daGasPrice, g.priceEncodingLength)
			gasPrice = new(big.Int).Add(gasPrice, daGasPrice)
		}
	}

	return gasPrice, nil
}

func (g DAGasPriceEstimator) DenoteInUSD(p GasPrice, wrappedNativePrice *big.Int) (GasPrice, error) {
	daGasPrice, execGasPrice, err := g.parseEncodedGasPrice(p)
	if err != nil {
		return nil, err
	}

	// This assumes l1GasPrice is priced using the same native token as l2 native
	daUSD := ccipcalc.CalculateUsdPerUnitGas(daGasPrice, wrappedNativePrice)
	if daUSD.BitLen() > int(g.priceEncodingLength) {
		return nil, fmt.Errorf("data availability gas price USD exceeded max range %+v", daUSD)
	}
	execUSD := ccipcalc.CalculateUsdPerUnitGas(execGasPrice, wrappedNativePrice)
	if execUSD.BitLen() > int(g.priceEncodingLength) {
		return nil, fmt.Errorf("exec gas price USD exceeded max range %+v", execUSD)
	}

	daUSD = new(big.Int).Lsh(daUSD, g.priceEncodingLength)
	return new(big.Int).Add(daUSD, execUSD), nil
}

func (g DAGasPriceEstimator) Median(gasPrices []GasPrice) (GasPrice, error) {
	daPrices := make([]*big.Int, len(gasPrices))
	execPrices := make([]*big.Int, len(gasPrices))

	for i := range gasPrices {
		daGasPrice, execGasPrice, err := g.parseEncodedGasPrice(gasPrices[i])
		if err != nil {
			return nil, err
		}

		daPrices[i] = daGasPrice
		execPrices[i] = execGasPrice
	}

	daMedian := ccipcalc.BigIntMedian(daPrices)
	execMedian := ccipcalc.BigIntMedian(execPrices)

	daMedian = new(big.Int).Lsh(daMedian, g.priceEncodingLength)
	return new(big.Int).Add(daMedian, execMedian), nil
}

func (g DAGasPriceEstimator) Deviates(p1 GasPrice, p2 GasPrice, opt GasPriceDeviationOptions) (bool, error) {
	p1DAGasPrice, p1ExecGasPrice, err := g.parseEncodedGasPrice(p1)
	if err != nil {
		return false, err
	}
	p2DAGasPrice, p2ExecGasPrice, err := g.parseEncodedGasPrice(p2)
	if err != nil {
		return false, err
	}

	deviated := ccipcalc.Deviates(p1DAGasPrice, p2DAGasPrice, opt.DADeviationPPB) || ccipcalc.Deviates(p1ExecGasPrice, p2ExecGasPrice, opt.ExecDeviationPPB)
	return deviated, nil
}

func (g DAGasPriceEstimator) EstimateMsgCostUSD(p GasPrice, wrappedNativePrice *big.Int, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta, opt MsgCostOptions) (*big.Int, error) {
	daGasPrice, execGasPrice, err := g.parseEncodedGasPrice(p)
	if err != nil {
		return nil, err
	}

	execCostUSD, err := g.execEstimator.EstimateMsgCostUSD(execGasPrice, wrappedNativePrice, msg, opt)
	if err != nil {
		return nil, err
	}

	// If there is data availability price component, then include data availability cost in fee estimation
	if daGasPrice.Cmp(big.NewInt(0)) > 0 {
		daGasCostUSD := g.estimateDACostUSD(daGasPrice, wrappedNativePrice, msg, opt)
		execCostUSD = new(big.Int).Add(daGasCostUSD, execCostUSD)
	}
	return execCostUSD, nil
}

func (g DAGasPriceEstimator) String(p GasPrice) string {
	daGasPrice, execGasPrice, err := g.parseEncodedGasPrice(p)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("DA Price: %s, Exec Price: %s", daGasPrice, execGasPrice)
}

func (g DAGasPriceEstimator) parseEncodedGasPrice(p *big.Int) (*big.Int, *big.Int, error) {
	if p.BitLen() > int(g.priceEncodingLength*2) {
		return nil, nil, fmt.Errorf("encoded gas price exceeded max range %+v", p)
	}

	daGasPrice := new(big.Int).Rsh(p, g.priceEncodingLength)

	daStart := new(big.Int).Lsh(big.NewInt(1), g.priceEncodingLength)
	execGasPrice := new(big.Int).Mod(p, daStart)

	return daGasPrice, execGasPrice, nil
}

func (g DAGasPriceEstimator) estimateDACostUSD(daGasPrice GasPrice, wrappedNativePrice *big.Int, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta, opt MsgCostOptions) *big.Int {
	var sourceTokenDataLen int
	for _, tokenData := range msg.SourceTokenData {
		sourceTokenDataLen += len(tokenData)
	}

	dataLen := EvmMessageFixedBytes + len(msg.Data) + len(msg.TokenAmounts)*EvmMessageBytesPerToken + sourceTokenDataLen
	dataGas := big.NewInt(int64(dataLen)*opt.GasPerDAByte + opt.DAOverheadGas)

	dataGasEstimate := new(big.Int).Mul(dataGas, daGasPrice)
	dataGasEstimate = new(big.Int).Div(new(big.Int).Mul(dataGasEstimate, big.NewInt(opt.DAMultiplier)), big.NewInt(DAMultiplierBase))

	return ccipcalc.CalculateUsdPerUnitGas(dataGasEstimate, wrappedNativePrice)
}
