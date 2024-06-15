package db

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"sort"
	"testing"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	cciporm "github.com/smartcontractkit/chainlink/v2/core/services/ccip"
	ccipmocks "github.com/smartcontractkit/chainlink/v2/core/services/ccip/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcalc"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcommon"
	ccipdatamocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/pricegetter"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/prices"
)

func Test_priceWrite(t *testing.T) {
	lggr := logger.TestLogger(t)

	gasPrice := big.NewInt(1e18)
	tokenPrices := map[cciptypes.Address]*big.Int{
		"0x123": big.NewInt(2e18),
		"0x234": big.NewInt(3e18),
	}
	jobId := int32(1)
	destChainSelector := uint64(12345)
	sourceChainSelector := uint64(67890)

	expectedGasPriceUpdate := []cciporm.GasPriceUpdate{
		{
			SourceChainSelector: sourceChainSelector,
			GasPrice:            assets.NewWei(gasPrice),
		},
	}
	expectedTokenPriceUpdate := []cciporm.TokenPriceUpdate{
		{
			TokenAddr:  "0x123",
			TokenPrice: assets.NewWei(big.NewInt(2e18)),
		},
		{
			TokenAddr:  "0x234",
			TokenPrice: assets.NewWei(big.NewInt(3e18)),
		},
	}

	testCases := []struct {
		name            string
		gasPriceError   bool
		tokenPriceError bool
		expectedErr     bool
	}{
		{
			name:            "ORM called successfully",
			gasPriceError:   false,
			tokenPriceError: false,
			expectedErr:     false,
		},
		{
			name:            "gasPrice clear failed",
			gasPriceError:   true,
			tokenPriceError: false,
			expectedErr:     true,
		},
		{
			name:            "tokenPrice clear failed",
			gasPriceError:   false,
			tokenPriceError: true,
			expectedErr:     true,
		},
		{
			name:            "both ORM calls failed",
			gasPriceError:   true,
			tokenPriceError: true,
			expectedErr:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tests.Context(t)

			var gasPricesError error
			var tokenPricesError error
			if tc.gasPriceError {
				gasPricesError = fmt.Errorf("gas prices error")
			}
			if tc.tokenPriceError {
				tokenPricesError = fmt.Errorf("token prices error")
			}

			mockOrm := ccipmocks.NewORM(t)
			mockOrm.On("InsertGasPricesForDestChain", ctx, destChainSelector, jobId, expectedGasPriceUpdate).Return(gasPricesError).Once()
			mockOrm.On("InsertTokenPricesForDestChain", ctx, destChainSelector, jobId, expectedTokenPriceUpdate).Return(tokenPricesError).Once()

			priceWrite := NewPriceWrite(
				lggr,
				mockOrm,
				jobId,
				destChainSelector,
				sourceChainSelector,
				"",
				nil,
				nil,
				nil,
				nil,
			).(*priceWrite)
			err := priceWrite.writePricesToDB(ctx, gasPrice, tokenPrices)
			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_generatePriceUpdates(t *testing.T) {
	lggr := logger.TestLogger(t)
	jobId := int32(1)
	destChainSelector := uint64(12345)
	sourceChainSelector := uint64(67890)

	const nTokens = 10
	tokens := make([]cciptypes.Address, nTokens)
	for i := range tokens {
		tokens[i] = cciptypes.Address(utils.RandomAddress().String())
	}
	sort.Slice(tokens, func(i, j int) bool { return tokens[i] < tokens[j] })

	testCases := []struct {
		name                 string
		tokenDecimals        map[cciptypes.Address]uint8
		sourceNativeToken    cciptypes.Address
		priceGetterRespData  map[cciptypes.Address]*big.Int
		priceGetterRespErr   error
		feeEstimatorRespFee  *big.Int
		feeEstimatorRespErr  error
		maxGasPrice          uint64
		expSourceGasPriceUSD *big.Int
		expTokenPricesUSD    map[cciptypes.Address]*big.Int
		expErr               bool
	}{
		{
			name: "base",
			tokenDecimals: map[cciptypes.Address]uint8{
				tokens[0]: 18,
				tokens[1]: 18,
			},
			sourceNativeToken: tokens[0],
			priceGetterRespData: map[cciptypes.Address]*big.Int{
				tokens[0]: val1e18(100),
				tokens[1]: val1e18(200),
				tokens[2]: val1e18(300), // price getter returned a price for this token even though we didn't request it (should be skipped)
			},
			priceGetterRespErr:   nil,
			feeEstimatorRespFee:  big.NewInt(10),
			feeEstimatorRespErr:  nil,
			maxGasPrice:          1e18,
			expSourceGasPriceUSD: big.NewInt(1000),
			expTokenPricesUSD: map[cciptypes.Address]*big.Int{
				tokens[0]: val1e18(100),
				tokens[1]: val1e18(200),
			},
			expErr: false,
		},
		{
			name: "price getter returned an error",
			tokenDecimals: map[cciptypes.Address]uint8{
				tokens[0]: 18,
				tokens[1]: 18,
			},
			sourceNativeToken:   tokens[0],
			priceGetterRespData: nil,
			priceGetterRespErr:  fmt.Errorf("some random network error"),
			expErr:              true,
		},
		{
			name: "price getter skipped a requested price",
			tokenDecimals: map[cciptypes.Address]uint8{
				tokens[0]: 18,
				tokens[1]: 18,
			},
			sourceNativeToken: tokens[0],
			priceGetterRespData: map[cciptypes.Address]*big.Int{
				tokens[0]: val1e18(100),
			},
			priceGetterRespErr: nil,
			expErr:             true,
		},
		{
			name: "price getter skipped source native price",
			tokenDecimals: map[cciptypes.Address]uint8{
				tokens[0]: 18,
				tokens[1]: 18,
			},
			sourceNativeToken: tokens[2],
			priceGetterRespData: map[cciptypes.Address]*big.Int{
				tokens[0]: val1e18(100),
				tokens[1]: val1e18(200),
			},
			priceGetterRespErr: nil,
			expErr:             true,
		},
		{
			name: "base",
			tokenDecimals: map[cciptypes.Address]uint8{
				tokens[0]: 18,
				tokens[1]: 18,
			},
			sourceNativeToken: tokens[0],
			priceGetterRespData: map[cciptypes.Address]*big.Int{
				tokens[0]: val1e18(100),
				tokens[1]: val1e18(200),
				tokens[2]: val1e18(300), // price getter returned a price for this token even though we didn't request it
			},
			priceGetterRespErr:   nil,
			feeEstimatorRespFee:  big.NewInt(10),
			feeEstimatorRespErr:  nil,
			maxGasPrice:          1e18,
			expSourceGasPriceUSD: big.NewInt(1000),
			expTokenPricesUSD: map[cciptypes.Address]*big.Int{
				tokens[0]: val1e18(100),
				tokens[1]: val1e18(200),
			},
			expErr: false,
		},
		{
			name: "dynamic fee cap overrides legacy",
			tokenDecimals: map[cciptypes.Address]uint8{
				tokens[0]: 18,
				tokens[1]: 18,
			},
			sourceNativeToken: tokens[0],
			priceGetterRespData: map[cciptypes.Address]*big.Int{
				tokens[0]: val1e18(100),
				tokens[1]: val1e18(200),
				tokens[2]: val1e18(300), // price getter returned a price for this token even though we didn't request it (should be skipped)
			},
			priceGetterRespErr:   nil,
			feeEstimatorRespFee:  big.NewInt(20),
			feeEstimatorRespErr:  nil,
			maxGasPrice:          1e18,
			expSourceGasPriceUSD: big.NewInt(2000),
			expTokenPricesUSD: map[cciptypes.Address]*big.Int{
				tokens[0]: val1e18(100),
				tokens[1]: val1e18(200),
			},
			expErr: false,
		},
		{
			name: "nil gas price",
			tokenDecimals: map[cciptypes.Address]uint8{
				tokens[0]: 18,
				tokens[1]: 18,
			},
			sourceNativeToken: tokens[0],
			priceGetterRespData: map[cciptypes.Address]*big.Int{
				tokens[0]: val1e18(100),
				tokens[1]: val1e18(200),
				tokens[2]: val1e18(300), // price getter returned a price for this token even though we didn't request it (should be skipped)
			},
			feeEstimatorRespFee: nil,
			maxGasPrice:         1e18,
			expErr:              true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			priceGetter := pricegetter.NewMockPriceGetter(t)
			defer priceGetter.AssertExpectations(t)

			gasPriceEstimator := prices.NewMockGasPriceEstimatorCommit(t)
			defer gasPriceEstimator.AssertExpectations(t)

			var destTokens []cciptypes.Address
			for tk := range tc.tokenDecimals {
				destTokens = append(destTokens, tk)
			}
			sort.Slice(destTokens, func(i, j int) bool {
				return destTokens[i] < destTokens[j]
			})
			var destDecimals []uint8
			for _, token := range destTokens {
				destDecimals = append(destDecimals, tc.tokenDecimals[token])
			}

			queryTokens := ccipcommon.FlattenUniqueSlice([]cciptypes.Address{tc.sourceNativeToken}, destTokens)

			if len(queryTokens) > 0 {
				priceGetter.On("TokenPricesUSD", mock.Anything, queryTokens).Return(tc.priceGetterRespData, tc.priceGetterRespErr)
			}

			if tc.maxGasPrice > 0 {
				gasPriceEstimator.On("GetGasPrice", mock.Anything).Return(tc.feeEstimatorRespFee, tc.feeEstimatorRespErr)
				if tc.feeEstimatorRespFee != nil {
					pUSD := ccipcalc.CalculateUsdPerUnitGas(tc.feeEstimatorRespFee, tc.expTokenPricesUSD[tc.sourceNativeToken])
					gasPriceEstimator.On("DenoteInUSD", mock.Anything, mock.Anything).Return(pUSD, nil)
				}
			}

			priceWrite := NewPriceWrite(
				lggr,
				nil,
				jobId,
				destChainSelector,
				sourceChainSelector,
				tc.sourceNativeToken,
				priceGetter,
				nil,
				nil,
				nil,
			).(*priceWrite)

			destPriceReg := ccipdatamocks.NewPriceRegistryReader(t)
			destPriceReg.On("GetTokensDecimals", mock.Anything, destTokens).Return(destDecimals, nil).Maybe()
			err := priceWrite.UpdateDynamicConfig(gasPriceEstimator, destPriceReg)
			assert.NoError(t, err)

			sourceGasPriceUSD, tokenPricesUSD, err := priceWrite.generatePriceUpdates(context.Background(), lggr, destTokens)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.True(t, tc.expSourceGasPriceUSD.Cmp(sourceGasPriceUSD) == 0)
			assert.True(t, reflect.DeepEqual(tc.expTokenPricesUSD, tokenPricesUSD))
		})
	}
}

func Test_priceWriteInBackground(t *testing.T) {
	lggr := logger.TestLogger(t)
	jobId := int32(1)
	destChainSelector := uint64(12345)
	sourceChainSelector := uint64(67890)

	sourceNative := cciptypes.Address("0x123")
	feeTokens := []cciptypes.Address{"0x234"}
	rampTokens := []cciptypes.Address{"0x345", "0x456"}
	rampFilteredTokens := []cciptypes.Address{"0x345"}
	rampFilterOutTokens := []cciptypes.Address{"0x456"}

	laneTokens := []cciptypes.Address{"0x234", "0x345"}
	laneTokenDecimals := []uint8{18, 18}

	tokens := []cciptypes.Address{sourceNative, "0x234", "0x345"}
	tokenPrices := []int64{2, 3, 4}

	gasPrice := big.NewInt(10)
	// gas price is written in dollar terms, value is calculated from evm gas price * price of source native
	expectedGasPriceUpdate := []cciporm.GasPriceUpdate{
		{
			SourceChainSelector: sourceChainSelector,
			GasPrice:            assets.NewWei(big.NewInt(10 * 2)),
		},
	}
	expectedTokenPriceUpdate := []cciporm.TokenPriceUpdate{
		{
			TokenAddr:  "0x234",
			TokenPrice: assets.NewWei(val1e18(3)),
		},
		{
			TokenAddr:  "0x345",
			TokenPrice: assets.NewWei(val1e18(4)),
		},
	}

	// run write task every 1 second
	expectedPriceUpdateInterval := 1 * time.Second

	mockOrm := ccipmocks.NewORM(t)

	priceGetter := pricegetter.NewMockPriceGetter(t)
	defer priceGetter.AssertExpectations(t)

	gasPriceEstimator := prices.NewMockGasPriceEstimatorCommit(t)
	defer gasPriceEstimator.AssertExpectations(t)

	priceGetter.On("TokenPricesUSD", mock.Anything, tokens).Return(map[cciptypes.Address]*big.Int{
		tokens[0]: val1e18(tokenPrices[0]),
		tokens[1]: val1e18(tokenPrices[1]),
		tokens[2]: val1e18(tokenPrices[2]),
	}, nil)
	priceGetter.On("FilterConfiguredTokens", mock.Anything, rampTokens).Return(rampFilteredTokens, rampFilterOutTokens, nil)

	offRampReader := ccipdatamocks.NewOffRampReader(t)
	offRampReader.On("GetTokens", mock.Anything).Return(cciptypes.OffRampTokens{
		DestinationTokens: rampTokens,
	}, nil).Maybe()

	gasPriceEstimator.On("GetGasPrice", mock.Anything).Return(gasPrice, nil)
	pUSD := ccipcalc.CalculateUsdPerUnitGas(gasPrice, val1e18(tokenPrices[0]))
	gasPriceEstimator.On("DenoteInUSD", mock.Anything, mock.Anything).Return(pUSD, nil)

	destPriceReg := ccipdatamocks.NewPriceRegistryReader(t)
	destPriceReg.On("GetTokensDecimals", mock.Anything, laneTokens).Return(laneTokenDecimals, nil).Maybe()
	destPriceReg.On("GetFeeTokens", mock.Anything).Return(feeTokens, nil).Maybe()

	priceWrite := NewPriceWrite(
		lggr,
		mockOrm,
		jobId,
		destChainSelector,
		sourceChainSelector,
		tokens[0],
		priceGetter,
		offRampReader,
		nil,
		nil,
	).(*priceWrite)
	priceWrite.updateInterval = expectedPriceUpdateInterval

	err := priceWrite.UpdateDynamicConfig(gasPriceEstimator, destPriceReg)
	assert.NoError(t, err)

	assert.NoError(t, priceWrite.Start(tests.Context(t)))

	for i := 0; i < 10; i++ {
		mockOrm.On("InsertGasPricesForDestChain", priceWrite.backgroundCtx, destChainSelector, jobId, expectedGasPriceUpdate).Return(nil).Once()
		mockOrm.On("InsertTokenPricesForDestChain", priceWrite.backgroundCtx, destChainSelector, jobId, expectedTokenPriceUpdate).Return(nil).Once()
		time.Sleep(expectedPriceUpdateInterval)
	}

	assert.NoError(t, priceWrite.Close())
	// No more calls after closing
	time.Sleep(expectedPriceUpdateInterval)
}

func Test_calculateUsdPer1e18TokenAmount(t *testing.T) {
	testCases := []struct {
		name       string
		price      *big.Int
		decimal    uint8
		wantResult *big.Int
	}{
		{
			name:       "18-decimal token, $6.5 per token",
			price:      big.NewInt(65e17),
			decimal:    18,
			wantResult: big.NewInt(65e17),
		},
		{
			name:       "6-decimal token, $1 per token",
			price:      big.NewInt(1e18),
			decimal:    6,
			wantResult: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e12)), // 1e30
		},
		{
			name:       "0-decimal token, $1 per token",
			price:      big.NewInt(1e18),
			decimal:    0,
			wantResult: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e18)), // 1e36
		},
		{
			name:       "36-decimal token, $1 per token",
			price:      big.NewInt(1e18),
			decimal:    36,
			wantResult: big.NewInt(1),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateUsdPer1e18TokenAmount(tt.price, tt.decimal)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}

func val1e18(val int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(1e18), big.NewInt(val))
}
