package pricegetter_test

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	ccipdatamocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/merclib"
	merclibmocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/merclib/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/pricegetter"
	v3report "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/v3/types"
)

func TestMercuryGetter_TokenPricesUSD(t *testing.T) {
	t.Run("happy path", func(tt *testing.T) {
		tokenAddresses := []common.Address{
			common.HexToAddress("0x1"),
			common.HexToAddress("0x2"),
			common.HexToAddress("0x3"),
		}
		feedIDs := [][32]byte{
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
		}
		expectedPrices := map[common.Address]*big.Int{
			tokenAddresses[0]: big.NewInt(1),
			tokenAddresses[1]: big.NewInt(2),
			tokenAddresses[2]: big.NewInt(3),
		}
		mc := &merclibmocks.MercuryClient{}
		mc.On("BatchFetchPrices", mock.Anything, feedIDs).Return([]*merclib.ReportWithContext{
			{
				FeedId: feedIDs[0],
				V3Report: &v3report.Report{
					BenchmarkPrice: expectedPrices[tokenAddresses[0]],
				},
			},
			{
				FeedId: feedIDs[1],
				V3Report: &v3report.Report{
					BenchmarkPrice: expectedPrices[tokenAddresses[1]],
				},
			},
			{
				FeedId: feedIDs[2],
				V3Report: &v3report.Report{
					BenchmarkPrice: expectedPrices[tokenAddresses[2]],
				},
			},
		}, nil)
		prr := &ccipdatamocks.PriceRegistryReader{}
		prr.On("GetFeedIDsForTokens", mock.Anything, tokenAddresses).Return(feedIDs, nil)
		mercGetter := pricegetter.NewMercuryGetter(mc, prr)
		actualPrices, err := mercGetter.TokenPricesUSD(testutils.Context(t), tokenAddresses)
		require.NoError(tt, err)
		require.Equal(tt, expectedPrices, actualPrices)
	})
	t.Run("can't get feed IDs", func(tt *testing.T) {
		tokenAddresses := []common.Address{
			common.HexToAddress("0x1"),
			common.HexToAddress("0x2"),
			common.HexToAddress("0x3"),
		}
		prr := &ccipdatamocks.PriceRegistryReader{}
		prr.On("GetFeedIDsForTokens", mock.Anything, tokenAddresses).Return(nil, errors.New("error getting feed IDs"))
		mercGetter := pricegetter.NewMercuryGetter(nil, prr)
		_, err := mercGetter.TokenPricesUSD(testutils.Context(t), tokenAddresses)
		require.Error(tt, err)
	})
	t.Run("error batch fetching prices", func(tt *testing.T) {
		tokenAddresses := []common.Address{
			common.HexToAddress("0x1"),
			common.HexToAddress("0x2"),
			common.HexToAddress("0x3"),
		}
		feedIDs := [][32]byte{
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
		}
		mc := &merclibmocks.MercuryClient{}
		mc.On("BatchFetchPrices", mock.Anything, feedIDs).Return(nil, errors.New("error fetching prices"))
		prr := &ccipdatamocks.PriceRegistryReader{}
		prr.On("GetFeedIDsForTokens", mock.Anything, tokenAddresses).Return(feedIDs, nil)
		mercGetter := pricegetter.NewMercuryGetter(mc, prr)
		_, err := mercGetter.TokenPricesUSD(testutils.Context(t), tokenAddresses)
		require.Error(tt, err)
	})
	t.Run("wrong length rwcs array returned by BatchFetchPrices", func(tt *testing.T) {
		tokenAddresses := []common.Address{
			common.HexToAddress("0x1"),
			common.HexToAddress("0x2"),
			common.HexToAddress("0x3"),
		}
		feedIDs := [][32]byte{
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
		}
		mc := &merclibmocks.MercuryClient{}
		mc.On("BatchFetchPrices", mock.Anything, feedIDs).Return([]*merclib.ReportWithContext{
			{
				FeedId: feedIDs[0],
				V3Report: &v3report.Report{
					BenchmarkPrice: big.NewInt(1),
				},
			},
			{
				FeedId: feedIDs[1],
				V3Report: &v3report.Report{
					BenchmarkPrice: big.NewInt(2),
				},
			},
		}, nil)
		prr := &ccipdatamocks.PriceRegistryReader{}
		prr.On("GetFeedIDsForTokens", mock.Anything, tokenAddresses).Return(feedIDs, nil)
		mercGetter := pricegetter.NewMercuryGetter(mc, prr)
		_, err := mercGetter.TokenPricesUSD(testutils.Context(t), tokenAddresses)
		require.Error(tt, err)
	})
}
