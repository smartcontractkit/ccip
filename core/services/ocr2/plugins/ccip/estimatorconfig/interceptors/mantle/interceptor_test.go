package mantle

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
)

func TestInterceptor(t *testing.T) {
	ethClient := mocks.NewClient(t)
	ctx := context.Background()

	tokenRatio := big.NewInt(10)
	interceptor := NewInterceptor(ctx, ethClient)

	ethClient.On("CallContract", ctx, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).
		Return(common.BigToHash(tokenRatio).Bytes(), nil)

	modExecGasPrice, modDAGasPrice, err := interceptor.ModifyGasPriceComponents(ctx, big.NewInt(1), big.NewInt(1))
	require.NoError(t, err)
	require.Equal(t, modExecGasPrice.Int64(), int64(20))
	require.Equal(t, modDAGasPrice.Int64(), int64(1))

	// second call won't invoke eth client
	modExecGasPrice, modDAGasPrice, err = interceptor.ModifyGasPriceComponents(ctx, big.NewInt(2), big.NewInt(1))
	require.NoError(t, err)
	require.Equal(t, modExecGasPrice.Int64(), int64(30))
	require.Equal(t, modDAGasPrice.Int64(), int64(1))
}

func TestModifyGasPriceComponents(t *testing.T) {
	testCases := map[string]struct {
		execGasPrice       *big.Int
		daGasPrice         *big.Int
		tokenRatio         *big.Int
		resultExecGasPrice *big.Int
	}{
		"regular": {
			execGasPrice:       big.NewInt(1),
			daGasPrice:         big.NewInt(1),
			tokenRatio:         big.NewInt(10),
			resultExecGasPrice: big.NewInt(20),
		},
		"zero DAGasPrice": {
			execGasPrice:       big.NewInt(1),
			daGasPrice:         big.NewInt(0),
			tokenRatio:         big.NewInt(1),
			resultExecGasPrice: big.NewInt(1),
		},
		"zero token ratio": {
			execGasPrice:       big.NewInt(15),
			daGasPrice:         big.NewInt(10),
			tokenRatio:         big.NewInt(0),
			resultExecGasPrice: big.NewInt(0),
		},
	}

	for tcName, tc := range testCases {
		t.Run(tcName, func(t *testing.T) {
			ethClient := mocks.NewClient(t)
			ctx := context.Background()

			interceptor := NewInterceptor(ctx, ethClient)

			ethClient.On("CallContract", ctx, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).
				Return(common.BigToHash(tc.tokenRatio).Bytes(), nil)

			modExecGasPrice, modDAGasPrice, err := interceptor.ModifyGasPriceComponents(ctx, tc.execGasPrice, tc.daGasPrice)
			require.NoError(t, err)
			require.Equal(t, modExecGasPrice, tc.resultExecGasPrice)
			require.Equal(t, modDAGasPrice, tc.daGasPrice)
		})
	}
}
