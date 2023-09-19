package pricegetter_test

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/bridges"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/pricegetter"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"

	pipelinemocks "github.com/smartcontractkit/chainlink/v2/core/services/pipeline/mocks"

	"github.com/smartcontractkit/chainlink/v2/core/store/models"

	config "github.com/smartcontractkit/chainlink/v2/core/internal/testutils/configtest/v2"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
)

func TestDataSource(t *testing.T) {
	linkEth := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(`{"JuelsPerETH": "200000000000000000000"}`))
		require.NoError(t, err)
	}))
	defer linkEth.Close()
	usdcEth := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(`{"USDCWeiPerETH": "1000000000000000000000"}`)) // 1000 USDC / ETH
		require.NoError(t, err)
	}))
	defer usdcEth.Close()
	linkTokenAddress := common.HexToAddress("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
	usdcTokenAddress := common.HexToAddress("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e10")
	source := fmt.Sprintf(`
	// Price 1
	link [type=http method=GET url="%s"];
	link_parse [type=jsonparse path="JuelsPerETH"];
	link->link_parse;
	// Price 2
	usdc [type=http method=GET url="%s"];
	usdc_parse [type=jsonparse path="USDCWeiPerETH"];
	usdc->usdc_parse;
	merge [type=merge left="{}" right="{\"%s\":$(link_parse), \"%s\":$(usdc_parse)}"];
`, linkEth.URL, usdcEth.URL, linkTokenAddress, usdcTokenAddress)

	priceGetter := newTestPipelineGetter(t, source)
	// Ask for all prices present in spec.
	prices, err := priceGetter.TokenPricesUSD(context.Background(), []common.Address{linkTokenAddress, usdcTokenAddress})
	require.NoError(t, err)
	assert.Equal(t, prices, map[common.Address]*big.Int{
		linkTokenAddress: big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1000000000000000000)),
		usdcTokenAddress: big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(1000000000000000000)),
	})
	// Ask a non-existent price.
	_, err = priceGetter.TokenPricesUSD(context.Background(), []common.Address{common.HexToAddress("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e11")})
	require.Error(t, err)
}

func TestParsingDifferentFormats(t *testing.T) {
	tests := []struct {
		name          string
		inputValue    string
		expectedValue *big.Int
		expectedError bool
	}{
		{
			name:          "number as string",
			inputValue:    "\"200000000000000000000\"",
			expectedValue: new(big.Int).Mul(big.NewInt(200), big.NewInt(1e18)),
		},
		{
			name:          "number as big number",
			inputValue:    "500000000000000000000",
			expectedValue: new(big.Int).Mul(big.NewInt(500), big.NewInt(1e18)),
		},
		{
			name:          "number as int64",
			inputValue:    "150",
			expectedValue: big.NewInt(150),
		},
		{
			name:          "number in scientific notation",
			inputValue:    "3e22",
			expectedValue: new(big.Int).Mul(big.NewInt(30000), big.NewInt(1e18)),
		},
		{
			name:          "number as string in scientific notation returns error",
			inputValue:    "\"3e22\"",
			expectedError: true,
		},
		{
			name:          "invalid value should return error",
			inputValue:    "\"NaN\"",
			expectedError: true,
		},
		{
			name:          "null should return error",
			inputValue:    "null",
			expectedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				_, err := w.Write([]byte(fmt.Sprintf(`{"MyCoin": %s}`, tt.inputValue)))
				require.NoError(t, err)
			}))
			defer token.Close()

			address := common.HexToAddress("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
			source := fmt.Sprintf(`
			// Price 1
			coin [type=http method=GET url="%s"];
			coin_parse [type=jsonparse path="MyCoin"];
			coin->coin_parse;
			merge [type=merge left="{}" right="{\"%s\":$(coin_parse)}"];
			`, token.URL, address)

			prices, err := newTestPipelineGetter(t, source).
				TokenPricesUSD(context.Background(), []common.Address{address})

			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, prices[address], tt.expectedValue)
			}
		})
	}
}

func newTestPipelineGetter(t *testing.T, source string) *pricegetter.PipelineGetter {
	lggr, _ := logger.NewLogger()
	cfg := pipelinemocks.NewConfig(t)
	cfg.On("MaxRunDuration").Return(time.Second)
	cfg.On("DefaultHTTPTimeout").Return(models.MakeDuration(time.Second))
	cfg.On("DefaultHTTPLimit").Return(int64(1024 * 10))
	db := pgtest.NewSqlxDB(t)
	bridgeORM := bridges.NewORM(db, lggr, config.NewTestGeneralConfig(t).Database())
	runner := pipeline.NewRunner(pipeline.NewORM(db, lggr, config.NewTestGeneralConfig(t).Database(), config.NewTestGeneralConfig(t).JobPipeline().MaxSuccessfulRuns()),
		bridgeORM, cfg, nil, nil, nil, nil, lggr, &http.Client{}, &http.Client{})
	ds, err := pricegetter.NewPipelineGetter(source, runner, 1, uuid.New(), "test", lggr)
	require.NoError(t, err)
	return ds
}
