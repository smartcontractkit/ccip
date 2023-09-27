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
