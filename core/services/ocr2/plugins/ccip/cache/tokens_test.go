package cache

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/link_token_interface"
	mock_contracts "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func Test_tokenToDecimals(t *testing.T) {
	tokenPriceMappings := map[common.Address]uint8{
		common.HexToAddress("0xA"): 10,
		common.HexToAddress("0xB"): 5,
		common.HexToAddress("0xC"): 2,
	}

	tests := []struct {
		name       string
		destTokens []common.Address
		feeTokens  []common.Address
		want       map[common.Address]uint8
	}{
		{
			name:       "empty map for empty tokens from origin",
			destTokens: []common.Address{},
			feeTokens:  []common.Address{},
			want:       map[common.Address]uint8{},
		},
		{
			name:       "separate destination and fee tokens",
			destTokens: []common.Address{common.HexToAddress("0xC")},
			feeTokens:  []common.Address{common.HexToAddress("0xB")},
			want: map[common.Address]uint8{
				common.HexToAddress("0xC"): 2,
				common.HexToAddress("0xB"): 5,
			},
		},
		{
			name:       "fee tokens and dest tokens are overlapping",
			destTokens: []common.Address{common.HexToAddress("0xA")},
			feeTokens:  []common.Address{common.HexToAddress("0xA")},
			want: map[common.Address]uint8{
				common.HexToAddress("0xA"): 10,
			},
		},
		{
			name:       "only fee tokens are returned",
			destTokens: []common.Address{},
			feeTokens:  []common.Address{common.HexToAddress("0xA"), common.HexToAddress("0xC")},
			want: map[common.Address]uint8{
				common.HexToAddress("0xA"): 10,
				common.HexToAddress("0xC"): 2,
			},
		},
		{
			name:       "missing tokens are skipped",
			destTokens: []common.Address{},
			feeTokens:  []common.Address{common.HexToAddress("0xD")},
			want:       map[common.Address]uint8{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offRamp := &mock_contracts.EVM2EVMOffRampInterface{}
			offRamp.On("GetDestinationTokens", mock.Anything).Return(tt.destTokens, nil)

			priceRegistry := &mock_contracts.PriceRegistryInterface{}
			priceRegistry.On("GetFeeTokens", mock.Anything).Return(tt.feeTokens, nil)

			tokenToDecimal := &tokenToDecimals{
				lggr:          logger.TestLogger(t),
				offRamp:       offRamp,
				priceRegistry: priceRegistry,
				tokenFactory:  createTokenFactory(tokenPriceMappings),
			}

			got, err := tokenToDecimal.CallOrigin(testutils.Context(t))
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func createTokenFactory(decimalMapping map[common.Address]uint8) func(address common.Address) (link_token_interface.LinkTokenInterface, error) {
	return func(address common.Address) (link_token_interface.LinkTokenInterface, error) {
		linkToken := &mock_contracts.LinkTokenInterface{}
		if decimals, found := decimalMapping[address]; found {
			// Make sure each token is fetched only once
			linkToken.On("Decimals", mock.Anything).Return(decimals, nil).Once()
		} else {
			linkToken.On("Decimals", mock.Anything).Return(uint8(0), errors.New("Error")).Once()
		}
		return linkToken, nil
	}
}
