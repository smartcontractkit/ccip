package reader

import (
	"context"
	"math/big"
	"testing"

	"github.com/smartcontractkit/ccipocr3/internal/mocks"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	ETH_ACC = ocr2types.Account("ETH")
	OP_ACC  = ocr2types.Account("OP")
	ARB_ACC = ocr2types.Account("ARB")
)

var (
	ETH_PRICE = big.NewInt(100)
	OP_PRICE  = big.NewInt(10)
	ARB_PRICE = big.NewInt(1)
)

func TestOnchainTokenPricesReader_GetTokenPricesUSD(t *testing.T) {

	testCases := []struct {
		name         string
		staticPrices map[ocr2types.Account]big.Int
		inputTokens  []ocr2types.Account
		mockPrices   map[ocr2types.Account]*big.Int
		want         []*big.Int
		wantErr      bool
	}{
		{
			name:         "Static price only",
			staticPrices: map[ocr2types.Account]big.Int{ETH_ACC: *ETH_PRICE, OP_ACC: *OP_PRICE},
			inputTokens:  []ocr2types.Account{ETH_ACC, OP_ACC},
			mockPrices:   map[ocr2types.Account]*big.Int{},
			want:         []*big.Int{ETH_PRICE, OP_PRICE},
		},
		{
			name:         "On-chain price only",
			staticPrices: map[ocr2types.Account]big.Int{},
			inputTokens:  []ocr2types.Account{ARB_ACC, OP_ACC, ETH_ACC},
			mockPrices:   map[ocr2types.Account]*big.Int{OP_ACC: OP_PRICE, ARB_ACC: ARB_PRICE, ETH_ACC: ETH_PRICE},
			want:         []*big.Int{ARB_PRICE, OP_PRICE, ETH_PRICE},
		},
		{
			name:         "Mix of static price and onchain price",
			staticPrices: map[ocr2types.Account]big.Int{ETH_ACC: *ETH_PRICE},
			inputTokens:  []ocr2types.Account{ETH_ACC, OP_ACC, ARB_ACC},
			mockPrices:   map[ocr2types.Account]*big.Int{OP_ACC: OP_PRICE, ARB_ACC: ARB_PRICE},
			want:         []*big.Int{ETH_PRICE, OP_PRICE, ARB_PRICE},
		},
	}

	for _, tc := range testCases {
		contractReader := createMockReader(tc.mockPrices)
		tokenPricesReader := OnchainTokenPricesReader{
			TokenPriceConfig: TokenPriceConfig{StaticPrices: tc.staticPrices},
			ContractReader:   contractReader,
		}
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			result, err := tokenPricesReader.GetTokenPricesUSD(ctx, tc.inputTokens)

			if tc.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.want, result)
		})
	}

}

func createMockReader(mockPrices map[ocr2types.Account]*big.Int) *mocks.ContractReaderMock {
	reader := mocks.NewContractReaderMock()
	for acc, price := range mockPrices {
		reader.On("GetLatestValue", mock.Anything, "PriceAggregator", "getTokenPrice", acc, mock.Anything).Run(
			func(args mock.Arguments) {
				arg := args.Get(4).(*big.Int)
				arg.Set(price)
			}).Return(nil)
	}

	return reader
}
