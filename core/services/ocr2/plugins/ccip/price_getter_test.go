package ccip_test

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/bridges"
	config "github.com/smartcontractkit/chainlink/core/internal/testutils/configtest/v2"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
	pipelinemocks "github.com/smartcontractkit/chainlink/core/services/pipeline/mocks"
	"github.com/smartcontractkit/chainlink/core/store/models"
)

func TestDataSource(t *testing.T) {
	linkEth := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"JuelsPerETH": "200000000000000000000"}`))
		require.NoError(t, err)
	}))
	defer linkEth.Close()
	usdcEth := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"USDCWeiPerETH": "1000000000000000000000"}`)) // 1000 USDC / ETH
		require.NoError(t, err)
	}))
	defer linkEth.Close()
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
	lggr, _ := logger.NewLogger()
	cfg := pipelinemocks.NewConfig(t)
	cfg.On("JobPipelineMaxRunDuration").Return(time.Second)
	cfg.On("DefaultHTTPTimeout").Return(models.MakeDuration(time.Second))
	cfg.On("DefaultHTTPLimit").Return(int64(1024 * 10))
	db := pgtest.NewSqlxDB(t)
	bridgeORM := bridges.NewORM(db, lggr, config.NewTestGeneralConfig(t))
	runner := pipeline.NewRunner(pipeline.NewORM(db, lggr, config.NewTestGeneralConfig(t)), bridgeORM, cfg, nil, nil, nil, lggr, &http.Client{}, &http.Client{})
	ds, err := ccip.NewPriceGetter(source, runner, 1, uuid.NewV1(), "test", lggr)
	require.NoError(t, err)

	// Ask for all prices present in spec.
	prices, err := ds.TokensPerFeeCoin(context.Background(), []common.Address{linkTokenAddress, usdcTokenAddress})
	require.NoError(t, err)
	assert.Equal(t, prices, map[common.Address]*big.Int{
		linkTokenAddress: big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1000000000000000000)),
		usdcTokenAddress: big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(1000000000000000000)),
	})

	// Ask a non-existent price.
	_, err = ds.TokensPerFeeCoin(context.Background(), []common.Address{common.HexToAddress("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e11")})
	require.Error(t, err)
}
