package mantle

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

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

func NewInterceptor(ctx context.Context, client evmClient.Client) *Interceptor {
	// Encode calldata for tokenRatio method
	tokenRatioMethodAbi, err := abi.JSON(strings.NewReader(mantleTokenRatioAbiString))
	if err != nil {
		log.Panicf("failed to parse GasPriceOracle %s() method ABI for Mantle; %v", tokenRatioMethod, err)
	}
	tokenRatioCallData, err := tokenRatioMethodAbi.Pack(tokenRatioMethod)
	if err != nil {
		log.Panicf("failed to parse GasPriceOracle %s() calldata for Mantle; %v", tokenRatioMethod, err)
	}

	return &Interceptor{
		client:   client,
		callData: tokenRatioCallData,
	}
}

// ModifyGasPriceComponents returns modified gasPrice.
func (i *Interceptor) ModifyGasPriceComponents(ctx context.Context, execGasPrice, daGasPrice *big.Int) (*big.Int, *big.Int, error) {
	if time.Since(i.tokenRatioLastUpdate) > tokenRatioUpdateInterval {
		mantleTokenRatio, err := i.getMantleGasPrice(ctx)
		if err != nil {
			return nil, nil, err
		}

		i.tokenRatio, i.tokenRatioLastUpdate = mantleTokenRatio, time.Now()
	}

	newExecGasPrice := new(big.Int).Mul(execGasPrice, i.tokenRatio)
	newDAGasPrice := new(big.Int).Mul(daGasPrice, i.tokenRatio)
	return newExecGasPrice, newDAGasPrice, nil
}

// getMantleGasPrice Requests and returns the token ratio for Mantle.
func (i *Interceptor) getMantleGasPrice(ctx context.Context) (*big.Int, error) {
	// call oracle to get l1BaseFee and tokenRatio
	precompile := common.HexToAddress(rollups.OPGasOracleAddress)
	tokenRatio, err := i.client.CallContract(ctx, ethereum.CallMsg{
		To:   &precompile,
		Data: i.callData,
	}, nil)

	if err != nil {
		return nil, fmt.Errorf("fetch gas price parameters call failed: %w", err)
	}

	// Convert bytes to big int for calculations and return
	return new(big.Int).SetBytes(tokenRatio), nil
}
