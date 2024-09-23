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
	"github.com/shopspring/decimal"

	evmClient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas/rollups"
)

const (
	// tokenRatio is not volatile and can be requested not often.
	tokenRatioUpdateInterval = 60 * time.Minute
	// tokenRatio fetches the tokenRatio used for Mantle's gas price calculation
	tokenRatioMethod          = "tokenRatio"
	mantleTokenRatioAbiString = `[{"inputs":[],"name":"tokenRatio","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
	// decimals fetches the number of chainDecimals used in the scalar for gas price calculation
	decimalsMethod = "decimals"
)

type Interceptor struct {
	client               evmClient.Client
	tokenRatioCallData   []byte
	decimalsCallData     []byte
	tokenRatio           decimal.Decimal
	chainDecimals        *big.Int
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

	// Encode calldata for decimals method
	decimalsMethodAbi, err := abi.JSON(strings.NewReader(rollups.OPDecimalsAbiString))
	if err != nil {
		log.Panicf("failed to parse GasPriceOracle %s() method ABI for Mantle; %v", decimalsMethod, err)
	}
	decimalsCallData, err := decimalsMethodAbi.Pack(decimalsMethod)
	if err != nil {
		log.Panicf("failed to parse GasPriceOracle %s() calldata for Mantle; %v", decimalsMethod, err)
	}

	return &Interceptor{
		client:             client,
		tokenRatioCallData: tokenRatioCallData,
		decimalsCallData:   decimalsCallData,
	}
}

// ModifyGasPriceComponents returns modified gasPrice.
func (i *Interceptor) ModifyGasPriceComponents(ctx context.Context, execGasPrice, daGasPrice *big.Int) (*big.Int, *big.Int, error) {
	if time.Since(i.tokenRatioLastUpdate) > tokenRatioUpdateInterval {
		mantleTokenRatio, err := i.getMantleTokenRatio(ctx)
		if err != nil {
			return nil, nil, err
		}

		i.tokenRatio, i.tokenRatioLastUpdate = mantleTokenRatio, time.Now()
	}

	// multiply daGasPrice and execGas price by tokenRatio
	dExecGasPrice := decimal.NewFromBigInt(execGasPrice, 0)
	newExecGasPrice := dExecGasPrice.Mul(i.tokenRatio).BigInt()

	dDAGasPrice := decimal.NewFromBigInt(daGasPrice, 0)
	newDAGasPrice := dDAGasPrice.Mul(i.tokenRatio).BigInt()

	return newExecGasPrice, newDAGasPrice, nil
}

// getMantleTokenRatio Requests and returns the token ratio value for the Mantle chain.
func (i *Interceptor) getMantleTokenRatio(ctx context.Context) (decimal.Decimal, error) {
	precompile := common.HexToAddress(rollups.OPGasOracleAddress)
	tokenRatio, err := i.client.CallContract(ctx, ethereum.CallMsg{
		To:   &precompile,
		Data: i.tokenRatioCallData,
	}, nil)

	if err != nil {
		return decimal.Zero, fmt.Errorf("getMantleTokenRatio call failed: %w", err)
	}

	bigIntTokenRatio := new(big.Int).SetBytes(tokenRatio)

	// request chainDecimals value once, it rarely changed and use cached value during the app lifecycle
	if i.chainDecimals == nil {
		decimals, err := i.getMantleDecimals(ctx)
		if err != nil {
			return decimal.Zero, err
		}

		i.chainDecimals = decimals
	}

	// convert bigInt token ratio to the decimal format
	// rawTokenRatio = bigIntTokenRatio * 10 ^ -i.chainDecimals
	exp := int32(-1 * i.chainDecimals.Int64())
	rawTokenRatio := decimal.NewFromBigInt(bigIntTokenRatio, exp)

	return rawTokenRatio, nil
}

// getMantleDecimals Requests and returns the decimals value for the Mantle chain.
func (i *Interceptor) getMantleDecimals(ctx context.Context) (*big.Int, error) {
	precompile := common.HexToAddress(rollups.OPGasOracleAddress)
	decimals, err := i.client.CallContract(ctx, ethereum.CallMsg{
		To:   &precompile,
		Data: i.decimalsCallData,
	}, nil)

	if err != nil {
		return nil, fmt.Errorf("getMantleDecimals call failed: %w", err)
	}

	return new(big.Int).SetBytes(decimals), nil
}
