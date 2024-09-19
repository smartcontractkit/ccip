package mantle

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	evmClient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas/rollups"
)

const (
	// tokenRatio is not volatile and can be requested not often.
	tokenRatioUpdateInterval = 60 * time.Minute
	// tokenRatio fetches the tokenRatio used for Mantle's gas price calculation
	// tokenRatio is a hex encoded call to:
	tokenRatioMethod          = "tokenRatio"
	mantleTokenRatioAbiString = `[{"inputs":[],"name":"tokenRatio","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
)

type Interceptor struct {
	client               evmClient.Client
	callData             []byte
	tokenRatio           *big.Int
	tokenRatioLastUpdate time.Time
}

func NewInterceptor(ctx context.Context, client evmClient.Client) (*Interceptor, error) {
	// Encode calldata for tokenRatio method
	tokenRatioMethodAbi, err := abi.JSON(strings.NewReader(mantleTokenRatioAbiString))
	if err != nil {
		return nil, fmt.Errorf("failed to parse GasPriceOracle %s() method ABI for Mantle; %w", tokenRatioMethod, err)
	}
	tokenRatioCallData, err := tokenRatioMethodAbi.Pack(tokenRatioMethod)
	if err != nil {
		return nil, fmt.Errorf("failed to parse GasPriceOracle %s() calldata for Mantle; %w", tokenRatioMethod, err)
	}

	interceptor := &Interceptor{
		client:   client,
		callData: tokenRatioCallData,
	}

	interceptor.tokenRatio, err = interceptor.getMantleGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get token ratio from the Mantle oracle %v", err)
	}

	return interceptor, nil
}

// ModifyGasPriceComponents returns modified gasPrice.
func (i *Interceptor) ModifyGasPriceComponents(ctx context.Context, gasPrice, daGasPrice *big.Int) (*big.Int, *big.Int, error) {
	if time.Since(i.tokenRatioLastUpdate) > tokenRatioUpdateInterval {
		var err error
		if i.tokenRatio, err = i.getMantleGasPrice(ctx); err != nil {
			return nil, nil, err
		}
	}

	newGasPrice := new(big.Int).Add(gasPrice, daGasPrice)

	return new(big.Int).Mul(newGasPrice, i.tokenRatio), daGasPrice, nil
}

// Request and returns the token ratio for Mantle.
func (i *Interceptor) getMantleGasPrice(ctx context.Context) (*big.Int, error) {
	// call oracle to get l1BaseFee and tokenRatio
	rpcBatchCalls := []rpc.BatchElem{
		{
			Method: "eth_call",
			Args: []any{
				map[string]interface{}{
					"from": common.Address{},
					"to":   rollups.OPGasOracleAddress,
					"data": hexutil.Bytes(i.callData),
				},
				"latest",
			},
			Result: new(string),
		},
	}

	err := i.client.BatchCallContext(ctx, rpcBatchCalls)
	if err != nil {
		return nil, fmt.Errorf("fetch gas price parameters batch call failed: %w", err)
	}
	if rpcBatchCalls[0].Error != nil {
		return nil, fmt.Errorf("%s call failed in a batch: %w", tokenRatioMethod, err)
	}

	// Extract values from responses
	tokenRatioResult := *(rpcBatchCalls[0].Result.(*string))

	// Decode the responses into bytes
	tokenRatioBytes, err := hexutil.Decode(tokenRatioResult)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s rpc result: %w", tokenRatioMethod, err)
	}

	// Convert bytes to big int for calculations
	tokenRatio := new(big.Int).SetBytes(tokenRatioBytes)

	// multiply l1BaseFee and tokenRatio and return
	return tokenRatio, nil
}
