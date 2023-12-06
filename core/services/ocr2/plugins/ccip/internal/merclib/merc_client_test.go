package merclib_test

import (
	"bytes"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	evmclient_mocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/models"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/merclib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/merclib/mocks"
)

func TestMercClient_BatchFetchPrices(t *testing.T) {
	t.Run("happy path", func(tt *testing.T) {
		feedIDs := [][32]byte{
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
		}
		prices := map[[32]byte]*big.Int{
			feedIDs[0]: big.NewInt(1),
			feedIDs[1]: big.NewInt(2),
			feedIDs[2]: big.NewInt(3),
		}
		response := genMercuryResponse(tt, feedIDs, prices)
		reader := io.NopCloser(bytes.NewReader(response))
		doer := &mocks.Doer{}
		doer.On("Do", mock.Anything).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       reader,
		}, nil)
		evmClient := &evmclient_mocks.Client{}
		evmClient.
			On("CallContract", mock.Anything, mock.AnythingOfType("ethereum.CallMsg"), mock.Anything).
			Return(nil, nil)
		mercClient := merclib.NewMercuryClient(
			&models.MercuryCredentials{
				Username: "testusername",
				Password: "testpassword",
				URL:      "https://testmercury.com",
			},
			doer,
			logger.TestLogger(t),
			evmClient,
			testutils.NewAddress(),
			testutils.NewAddress(),
			testutils.NewAddress())
		rwcs, err := mercClient.BatchFetchPrices(testutils.Context(t), feedIDs)
		require.NoError(tt, err)
		actualPrices := make(map[[32]byte]*big.Int)
		for _, rwc := range rwcs {
			actualPrices[rwc.FeedId] = rwc.V3Report.BenchmarkPrice
		}
		require.Equal(tt, prices, actualPrices)
	})
	t.Run("http Do error", func(tt *testing.T) {
		feedIDs := [][32]byte{
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
		}
		doer := &mocks.Doer{}
		doer.On("Do", mock.Anything).Return(nil, errors.New("http error"))
		evmClient := &evmclient_mocks.Client{}
		mercClient := merclib.NewMercuryClient(
			&models.MercuryCredentials{
				Username: "testusername",
				Password: "testpassword",
				URL:      "https://testmercury.com",
			},
			doer,
			logger.TestLogger(t),
			evmClient,
			testutils.NewAddress(),
			testutils.NewAddress(),
			testutils.NewAddress())
		_, err := mercClient.BatchFetchPrices(testutils.Context(t), feedIDs)
		require.Error(tt, err)
	})
	t.Run("non-200 status code", func(tt *testing.T) {
		feedIDs := [][32]byte{
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
		}
		doer := &mocks.Doer{}
		doer.On("Do", mock.Anything).Return(&http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       io.NopCloser(bytes.NewReader([]byte("blah"))),
		}, nil)
		evmClient := &evmclient_mocks.Client{}
		mercClient := merclib.NewMercuryClient(
			&models.MercuryCredentials{
				Username: "testusername",
				Password: "testpassword",
				URL:      "https://testmercury.com",
			},
			doer,
			logger.TestLogger(t),
			evmClient,
			testutils.NewAddress(),
			testutils.NewAddress(),
			testutils.NewAddress())
		_, err := mercClient.BatchFetchPrices(testutils.Context(t), feedIDs)
		require.Error(tt, err)
	})
	t.Run("json unmarshal error", func(tt *testing.T) {
		feedIDs := [][32]byte{
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
		}
		doer := &mocks.Doer{}
		doer.On("Do", mock.Anything).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader([]byte("invalid json"))),
		}, nil)
		evmClient := &evmclient_mocks.Client{}
		mercClient := merclib.NewMercuryClient(
			&models.MercuryCredentials{
				Username: "testusername",
				Password: "testpassword",
				URL:      "https://testmercury.com",
			},
			doer,
			logger.TestLogger(t),
			evmClient,
			testutils.NewAddress(),
			testutils.NewAddress(),
			testutils.NewAddress())
		_, err := mercClient.BatchFetchPrices(testutils.Context(t), feedIDs)
		require.Error(tt, err)
	})
	t.Run("report verification error", func(tt *testing.T) {
		feedIDs := [][32]byte{
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
			testutils.RandomFeedIDV3(),
		}
		prices := map[[32]byte]*big.Int{
			feedIDs[0]: big.NewInt(1),
			feedIDs[1]: big.NewInt(2),
			feedIDs[2]: big.NewInt(3),
		}
		response := genMercuryResponse(tt, feedIDs, prices)
		reader := io.NopCloser(bytes.NewReader(response))
		doer := &mocks.Doer{}
		doer.On("Do", mock.Anything).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       reader,
		}, nil)
		evmClient := &evmclient_mocks.Client{}
		evmClient.
			On("CallContract", mock.Anything, mock.AnythingOfType("ethereum.CallMsg"), mock.Anything).
			Return(nil, errors.New("call error"))
		mercClient := merclib.NewMercuryClient(
			&models.MercuryCredentials{
				Username: "testusername",
				Password: "testpassword",
				URL:      "https://testmercury.com",
			},
			doer,
			logger.TestLogger(t),
			evmClient,
			testutils.NewAddress(),
			testutils.NewAddress(),
			testutils.NewAddress())
		_, err := mercClient.BatchFetchPrices(testutils.Context(t), feedIDs)
		require.Error(tt, err)
	})
}

func randomReportContext() [3][32]byte {
	return [3][32]byte{
		testutils.Random32Byte(),
		testutils.Random32Byte(),
		testutils.Random32Byte(),
	}
}

func genMercuryResponse(
	t *testing.T,
	feedIDs [][32]byte,
	prices map[[32]byte]*big.Int) []byte {
	var v3reports []merclib.MercuryV03Report
	for _, feedID := range feedIDs {
		reportData, err := merclib.EncodeReportDataV3(
			feedID,
			1,
			1,
			big.NewInt(0),
			big.NewInt(0),
			1,
			prices[feedID],
			prices[feedID],
			prices[feedID],
		)
		require.NoError(t, err)
		fullReport, err := merclib.EncodeFullReport(
			randomReportContext(),
			reportData,
			randomSliceOf32ByteArrays(),
			randomSliceOf32ByteArrays(),
			testutils.Random32Byte(),
		)
		require.NoError(t, err)
		v3reports = append(v3reports, merclib.MercuryV03Report{
			FeedID:                hexutil.Encode(feedID[:]),
			ValidFromTimestamp:    1,
			ObservationsTimestamp: 1,
			FullReport:            hexutil.Encode(fullReport),
		})
	}
	v3Response := merclib.MercuryV03Response{
		Reports: v3reports,
	}
	jsonified, err := json.Marshal(v3Response)
	require.NoError(t, err)
	return jsonified
}

func randomSliceOf32ByteArrays() [][32]byte {
	return [][32]byte{
		testutils.Random32Byte(),
		testutils.Random32Byte(),
		testutils.Random32Byte(),
	}
}
