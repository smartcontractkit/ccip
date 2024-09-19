package mantle

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"

	evmClient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/config/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas/rollups"
)

const (
	// tokenRatio fetches the tokenRatio used for Mantle's gas price calculation
	// tokenRatio is a hex encoded call to:
	// `function tokenRatio() public pure returns (uint256);`
	tokenRatioMethod          = "tokenRatio"
	mantleTokenRatioAbiString = `[{"inputs":[],"name":"tokenRatio","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
)

type Interceptor struct {
	client     evmClient.Client
	callData   []byte
	tokenRatio *big.Int
}

func NewInterceptor(ctx context.Context, client evmClient.Client, chainType chaintype.ChainType) (*Interceptor, error) {
	if chainType != chaintype.ChainMantle { // TODO: change to mantle when it will be available from chainlink repo
		return nil, nil
	}
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

	go interceptor.processUpdate(ctx)

	return interceptor, nil
}

func (i *Interceptor) ModifyDAGasPrice(gasPrice *big.Int) *big.Int {
	return new(big.Int).Mul(gasPrice, i.tokenRatio)
}

func (i *Interceptor) processUpdate(ctx context.Context) {
	ticker := time.NewTicker(time.Minute * 10) // TODO: get from constant

	for {
		select {
		case <-ticker.C:
			tokenRatio, err := i.getMantleGasPrice(ctx)
			if err != nil {
				log.Printf("could not get token ratio from the Mantle oracle %v", err)
				continue
			}
			i.tokenRatio = tokenRatio
		case <-ctx.Done():
			return
		}
	}
}

// Returns the gas price for Mantle. The formula is the same as Optimism Bedrock (getV1GasPrice), but the tokenRatio parameter is multiplied
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
