package prices

import (
	"context"
	"math/big"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas/mocks"
)

func TestUSDCReader_GetGasPrice(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name                      string
		sourceFeeEstimatorRespFee gas.EvmFee
		sourceFeeEstimatorRespErr error
		maxGasPrice               *big.Int
		expPrice                  GasPrice
		expErr                    bool
	}{
		{
			name: "gets legacy gas price",
			sourceFeeEstimatorRespFee: gas.EvmFee{
				Legacy:        assets.NewWei(big.NewInt(10)),
				DynamicFeeCap: nil,
			},
			sourceFeeEstimatorRespErr: nil,
			maxGasPrice:               big.NewInt(1),
			expPrice:                  big.NewInt(10),
			expErr:                    false,
		},
		{
			name: "gets dynamic gas price",
			sourceFeeEstimatorRespFee: gas.EvmFee{
				Legacy:        nil,
				DynamicFeeCap: assets.NewWei(big.NewInt(20)),
			},
			sourceFeeEstimatorRespErr: nil,
			maxGasPrice:               big.NewInt(1),
			expPrice:                  big.NewInt(20),
			expErr:                    false,
		},
		{
			name: "gets dynamic gas price over legacy gas price",
			sourceFeeEstimatorRespFee: gas.EvmFee{
				Legacy:        assets.NewWei(big.NewInt(10)),
				DynamicFeeCap: assets.NewWei(big.NewInt(20)),
			},
			sourceFeeEstimatorRespErr: nil,
			maxGasPrice:               big.NewInt(1),
			expPrice:                  big.NewInt(20),
			expErr:                    false,
		},
		{
			name: "fee estimator error",
			sourceFeeEstimatorRespFee: gas.EvmFee{
				Legacy:        assets.NewWei(big.NewInt(10)),
				DynamicFeeCap: nil,
			},
			sourceFeeEstimatorRespErr: errors.New("fee estimator error"),
			maxGasPrice:               big.NewInt(1),
			expPrice:                  nil,
			expErr:                    true,
		},
		{
			name: "nil gas price error",
			sourceFeeEstimatorRespFee: gas.EvmFee{
				Legacy:        nil,
				DynamicFeeCap: nil,
			},
			sourceFeeEstimatorRespErr: nil,
			maxGasPrice:               big.NewInt(1),
			expPrice:                  nil,
			expErr:                    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sourceFeeEstimator := mocks.NewEvmFeeEstimator(t)
			sourceFeeEstimator.On("GetFee", ctx, []byte(nil), uint32(0), assets.NewWei(tc.maxGasPrice)).Return(
				tc.sourceFeeEstimatorRespFee, uint32(0), tc.sourceFeeEstimatorRespErr)

			g := ExecGasPriceEstimator{
				estimator:   sourceFeeEstimator,
				maxGasPrice: tc.maxGasPrice,
			}

			gasPrice, err := g.GetGasPrice(ctx)
			if tc.expErr {
				assert.Nil(t, gasPrice)
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expPrice, gasPrice)
		})
	}
}

func TestUSDCReader_DenoteInUSD(t *testing.T) {
	val1e18 := func(val int64) *big.Int { return new(big.Int).Mul(big.NewInt(1e18), big.NewInt(val)) }

	testCases := []struct {
		name        string
		gasPrice    GasPrice
		nativePrice *big.Int
		expPrice    GasPrice
	}{
		{
			name:        "base",
			gasPrice:    big.NewInt(1e9),
			nativePrice: val1e18(2000),
			expPrice:    big.NewInt(2000e9),
		},
		{
			name:        "low price truncates to 0",
			gasPrice:    big.NewInt(1e9),
			nativePrice: big.NewInt(1),
			expPrice:    big.NewInt(0),
		},
		{
			name:        "high price",
			gasPrice:    val1e18(1),
			nativePrice: val1e18(2000),
			expPrice:    val1e18(2000),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := ExecGasPriceEstimator{
				estimator:   nil,
				maxGasPrice: nil,
			}

			gasPrice, err := g.DenoteInUSD(tc.gasPrice, tc.nativePrice)
			assert.NoError(t, err)
			assert.True(t, ((*big.Int)(tc.expPrice)).Cmp(gasPrice) == 0)
		})
	}
}

func TestUSDCReader_Median(t *testing.T) {
	val1e18 := func(val int64) *big.Int { return new(big.Int).Mul(big.NewInt(1e18), big.NewInt(val)) }

	testCases := []struct {
		name      string
		gasPrices []GasPrice
		expMedian GasPrice
	}{
		{
			name:      "base",
			gasPrices: []GasPrice{big.NewInt(1), big.NewInt(2), big.NewInt(3)},
			expMedian: big.NewInt(2),
		},
		{
			name:      "median 1",
			gasPrices: []GasPrice{big.NewInt(1)},
			expMedian: big.NewInt(1),
		},
		{
			name:      "median 2",
			gasPrices: []GasPrice{big.NewInt(1), big.NewInt(2)},
			expMedian: big.NewInt(2),
		},
		{
			name:      "large values",
			gasPrices: []GasPrice{val1e18(5), val1e18(4), val1e18(3), val1e18(2), val1e18(1)},
			expMedian: val1e18(3),
		},
		{
			name:      "zeros",
			gasPrices: []GasPrice{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
			expMedian: big.NewInt(0),
		},
		{
			name:      "unsorted even number of prices",
			gasPrices: []GasPrice{big.NewInt(4), big.NewInt(2), big.NewInt(3), big.NewInt(1)},
			expMedian: big.NewInt(3),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := ExecGasPriceEstimator{
				estimator:   nil,
				maxGasPrice: nil,
			}

			gasPrice, err := g.Median(tc.gasPrices)
			assert.NoError(t, err)
			assert.True(t, ((*big.Int)(tc.expMedian)).Cmp(gasPrice) == 0)
		})
	}
}
