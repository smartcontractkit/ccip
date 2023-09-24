package prices

import (
	"context"
	"fmt"
	"math/big"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcalc"
)

type ExecGasPriceEstimator struct {
	estimator    gas.EvmFeeEstimator
	maxGasPrice  *big.Int
	deviationPPB *int64
}

func (g ExecGasPriceEstimator) GetGasPrice(ctx context.Context) (GasPrice, error) {
	gasPriceWei, _, err := g.estimator.GetFee(ctx, nil, 0, assets.NewWei(g.maxGasPrice))
	if err != nil {
		return nil, err
	}
	// Use legacy if no dynamic is available.
	gasPrice := gasPriceWei.Legacy.ToInt()
	if gasPriceWei.DynamicFeeCap != nil {
		gasPrice = gasPriceWei.DynamicFeeCap.ToInt()
	}
	if gasPrice == nil {
		return nil, fmt.Errorf("missing gas price %+v", gasPriceWei)
	}

	return gasPrice, nil
}

func (g ExecGasPriceEstimator) DenoteInUSD(p GasPrice, wrappedNativePrice *big.Int) (GasPrice, error) {
	return ccipcalc.CalculateUsdPerUnitGas(p, wrappedNativePrice), nil
}

func (g ExecGasPriceEstimator) Median(gasPrices []GasPrice) (GasPrice, error) {
	var prices []*big.Int
	for _, p := range gasPrices {
		if p != nil {
			prices = append(prices, p)
		}
	}

	return ccipcalc.BigIntMedian(prices), nil
}

func (g ExecGasPriceEstimator) Deviates(p1 GasPrice, p2 GasPrice) (bool, error) {
	if g.deviationPPB == nil {
		return false, errors.New("missing gas price deviation ppb")
	}
	return ccipcalc.Deviates(p1, p2, *g.deviationPPB), nil
}

// EstimateMsgCostUSD calculates the costs for next execution, and converts to USD value scaled by 1e18 (e.g. 5$ = 5e18).
func (g ExecGasPriceEstimator) EstimateMsgCostUSD(p GasPrice, wrappedNativePrice *big.Int, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta, _ MsgCostConfig) (*big.Int, error) {
	execGasAmount := new(big.Int).Add(big.NewInt(FeeBoostingOverheadGas), msg.GasLimit)
	execGasAmount = new(big.Int).Add(execGasAmount, new(big.Int).Mul(big.NewInt(int64(len(msg.Data))), big.NewInt(ExecGasPerPayloadByte)))
	execGasAmount = new(big.Int).Add(execGasAmount, new(big.Int).Mul(big.NewInt(int64(len(msg.TokenAmounts))), big.NewInt(ExecGasPerToken)))

	execGasCost := new(big.Int).Mul(execGasAmount, p)

	return ccipcalc.CalculateUsdPerUnitGas(execGasCost, wrappedNativePrice), nil
}

func (g ExecGasPriceEstimator) String(p GasPrice) string {
	var pi *big.Int = p
	return pi.String()
}
