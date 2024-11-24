// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_onramp

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
	zkSyncAccounts "github.com/zksync-sdk/zksync2-go/accounts"
	zkSyncClient "github.com/zksync-sdk/zksync2-go/clients"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type ClientEVM2AnyMessage struct {
	Receiver     []byte
	Data         []byte
	TokenAmounts []ClientEVMTokenAmount
	FeeToken     common.Address
	ExtraArgs    []byte
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type EVM2EVMOnRampDynamicConfig struct {
	Router                            common.Address
	MaxNumberOfTokensPerMsg           uint16
	DestGasOverhead                   uint32
	DestGasPerPayloadByte             uint16
	DestDataAvailabilityOverheadGas   uint32
	DestGasPerDataAvailabilityByte    uint16
	DestDataAvailabilityMultiplierBps uint16
	PriceRegistry                     common.Address
	MaxDataBytes                      uint32
	MaxPerMsgGasLimit                 uint32
	DefaultTokenFeeUSDCents           uint16
	DefaultTokenDestGasOverhead       uint32
	EnforceOutOfOrder                 bool
}

type EVM2EVMOnRampFeeTokenConfig struct {
	NetworkFeeUSDCents         uint32
	GasMultiplierWeiPerEth     uint64
	PremiumMultiplierWeiPerEth uint64
	Enabled                    bool
}

type EVM2EVMOnRampFeeTokenConfigArgs struct {
	Token                      common.Address
	NetworkFeeUSDCents         uint32
	GasMultiplierWeiPerEth     uint64
	PremiumMultiplierWeiPerEth uint64
	Enabled                    bool
}

type EVM2EVMOnRampNopAndWeight struct {
	Nop    common.Address
	Weight uint16
}

type EVM2EVMOnRampStaticConfig struct {
	LinkToken          common.Address
	ChainSelector      uint64
	DestChainSelector  uint64
	DefaultTxGasLimit  uint64
	MaxNopFeesJuels    *big.Int
	PrevOnRamp         common.Address
	RmnProxy           common.Address
	TokenAdminRegistry common.Address
}

type EVM2EVMOnRampTokenTransferFeeConfig struct {
	MinFeeUSDCents            uint32
	MaxFeeUSDCents            uint32
	DeciBps                   uint16
	DestGasOverhead           uint32
	DestBytesOverhead         uint32
	AggregateRateLimitEnabled bool
	IsEnabled                 bool
}

type EVM2EVMOnRampTokenTransferFeeConfigArgs struct {
	Token                     common.Address
	MinFeeUSDCents            uint32
	MaxFeeUSDCents            uint32
	DeciBps                   uint16
	DestGasOverhead           uint32
	DestBytesOverhead         uint32
	AggregateRateLimitEnabled bool
}

type InternalEVM2EVMMessage struct {
	SourceChainSelector uint64
	Sender              common.Address
	Receiver            common.Address
	SequenceNumber      uint64
	GasLimit            *big.Int
	Strict              bool
	Nonce               uint64
	FeeToken            common.Address
	FeeTokenAmount      *big.Int
	Data                []byte
	TokenAmounts        []ClientEVMTokenAmount
	SourceTokenData     [][]byte
	MessageId           [32]byte
}

type RateLimiterConfig struct {
	IsEnabled bool
	Capacity  *big.Int
	Rate      *big.Int
}

type RateLimiterTokenBucket struct {
	Tokens      *big.Int
	LastUpdated uint32
	IsEnabled   bool
	Capacity    *big.Int
	Rate        *big.Int
}

var EVM2EVMOnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSendZeroTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtraArgOutOfOrderExecutionMustBeTrue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChainSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"}],\"name\":\"InvalidDestBytesOverhead\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"}],\"name\":\"InvalidNopAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkBalanceNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeBalanceReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeesToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNopsToPay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NotAFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdminOrNop\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SourceTokenDataTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyNops\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeConfig\",\"type\":\"tuple[]\"}],\"name\":\"FeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NopPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nopWeightsTotal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"NopsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokenTransferFeeConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"transferFeeConfig\",\"type\":\"tuple[]\"}],\"name\":\"TokenTransferFeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensConsumed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeeTokenConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfig\",\"name\":\"feeTokenConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNopFeesJuels\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNops\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"weightsTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPoolV1\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenTransferFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"setNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"address[]\"}],\"name\":\"setTokenTransferFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawNonLinkFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101a06040523480156200001257600080fd5b506040516200822d3803806200822d833981016040819052620000359162001a93565b8333806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c0816200030b565b50506040805160a081018252602084810180516001600160801b039081168085524263ffffffff169385018490528751151585870181905292518216606086018190529790950151166080909301839052600380546001600160a01b031916909417600160801b9283021760ff60a01b1916600160a01b90910217909255029091176004555085516001600160a01b0316158062000169575060208601516001600160401b0316155b8062000180575060408601516001600160401b0316155b8062000197575060608601516001600160401b0316155b80620001ae575060c08601516001600160a01b0316155b80620001c5575060e08601516001600160a01b0316155b15620001e4576040516306b7c75960e31b815260040160405180910390fd5b60208087015160408089015181517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3948101949094526001600160401b039283169184019190915216606082015230608082015260a00160408051601f198184030181529181528151602092830120608090815288516001600160a01b0390811660e0908152938a01516001600160401b0390811661010052928a015183166101205260608a015190921660a0908152908901516001600160601b031660c090815290890151821661014052880151811661016052908701511661018052620002cd85620003b6565b620002d883620006ad565b604080516000815260208101909152620002f4908390620007dd565b620002ff8162000a83565b5050505050506200218a565b336001600160a01b03821603620003655760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60e08101516001600160a01b0316620003e2576040516306b7c75960e31b815260040160405180910390fd5b80600560008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160000160146101000a81548161ffff021916908361ffff16021790555060408201518160000160166101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001601a6101000a81548161ffff021916908361ffff160217905550608082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160010160006101000a81548161ffff021916908361ffff16021790555060c08201518160010160026101000a81548161ffff021916908361ffff16021790555060e08201518160010160046101000a8154816001600160a01b0302191690836001600160a01b031602179055506101008201518160010160186101000a81548163ffffffff021916908363ffffffff16021790555061012082015181600101601c6101000a81548163ffffffff021916908363ffffffff1602179055506101408201518160020160006101000a81548161ffff021916908361ffff1602179055506101608201518160020160026101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160020160066101000a81548160ff0219169083151502179055509050507f45b5ad483aa608464c2c7f278bd413d284d7790cdc836e40652e23a02770822060405180610100016040528060e0516001600160a01b03168152602001610100516001600160401b03168152602001610120516001600160401b0316815260200160a0516001600160401b0316815260200160c0516001600160601b03168152602001610140516001600160a01b03168152602001610160516001600160a01b03168152602001610180516001600160a01b031681525082604051620006a292919062001d1c565b60405180910390a150565b60005b8151811015620007ab576000828281518110620006d157620006d162001de1565b60209081029190910181015160408051608080820183528385015163ffffffff9081168352838501516001600160401b03908116848801908152606080880151831686880190815294880151151590860190815296516001600160a01b03166000908152600b90985294909620925183549451925195511515600160a01b0260ff60a01b199688166c010000000000000000000000000296909616600160601b600160a81b031993909716640100000000026001600160601b031990951691161792909217919091169290921717905550600101620006b0565b507f067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e81604051620006a2919062001df7565b60005b8251811015620009a257600083828151811062000801576200080162001de1565b60200260200101519050602063ffffffff168160a0015163ffffffff1610156200085d57805160a08201516040516312766e0160e11b81526001600160a01b03909216600483015263ffffffff16602482015260440162000084565b6040805160e08101825260208381015163ffffffff908116835284840151811682840190815260608087015161ffff9081168688019081526080808a0151861693880193845260a0808b0151871691890191825260c0808c01511515918a019182526001908a018181529b516001600160a01b03166000908152600c9099529990972097518854955192519451915197519a519087166001600160401b031990961695909517640100000000928716929092029190911765ffffffffffff60401b191668010000000000000000939092169290920263ffffffff60501b1916176a0100000000000000000000918416919091021764ffffffffff60701b1916600160701b939092169290920260ff60901b191617600160901b941515949094029390931760ff60981b1916600160981b931515939093029290921790915501620007e0565b507ff5791bc457b3bb990493cf5f655db46c25ccf5764c9b99b8969b4c72ea7df9d082604051620009d4919062001e86565b60405180910390a160005b815181101562000a3c57600c600083838151811062000a025762000a0262001de1565b6020908102919091018101516001600160a01b0316825281019190915260400160002080546001600160a01b0319169055600101620009df565b5080511562000a7f577ffb95a0042158e60a33e7b5bec100f3d95407b1a71bee6633bd54b8887449750b8160405162000a76919062001f1b565b60405180910390a15b5050565b8051604081111562000aa857604051635ad0867d60e11b815260040160405180910390fd5b600e546c01000000000000000000000000900463ffffffff161562000afb57600e5463ffffffff6c010000000000000000000000008204166001600160601b039091161062000afb5762000afb62000c9e565b600062000b09600862000e8a565b90505b801562000b5557600062000b2f62000b2660018462001f80565b60089062000e9d565b50905062000b3f60088262000ebb565b50508062000b4d9062001f96565b905062000b0c565b506000805b8281101562000c3557600084828151811062000b7a5762000b7a62001de1565b6020026020010151600001519050600085838151811062000b9f5762000b9f62001de1565b602002602001015160200151905060e0516001600160a01b0316826001600160a01b0316148062000bd757506001600160a01b038216155b1562000c0257604051634de938d160e01b81526001600160a01b038316600482015260240162000084565b62000c1460088361ffff841662000ed9565b5062000c2561ffff82168562001fb0565b9350505080600101905062000b5a565b50600e805463ffffffff60601b19166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd249062000c91908390869062001fd0565b60405180910390a1505050565b6000546001600160a01b0316331462000cef576002546001600160a01b0316331462000cef5762000cd160083362000ef9565b62000cef5760405163032bb72b60e31b815260040160405180910390fd5b600e546c01000000000000000000000000900463ffffffff16600081900362000d2b5760405163990e30bf60e01b815260040160405180910390fd5b600e546001600160601b03168181101562000d59576040516311a1ee3b60e31b815260040160405180910390fd5b600062000d6562000f10565b121562000d8557604051631e9acf1760e31b815260040160405180910390fd5b80600062000d94600862000e8a565b905060005b8181101562000e645760008062000db260088462000e9d565b909250905060008762000dcf836001600160601b038a1662002040565b62000ddb91906200205a565b905062000de981876200207d565b60e05190965062000e0e906001600160a01b0316846001600160601b03841662000f9e565b6040516001600160601b03821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a250505080600101905062000d99565b5050600e80546001600160601b0319166001600160601b03929092169190911790555050565b600062000e978262000ffb565b92915050565b600080808062000eae868662001008565b9097909650945050505050565b600062000ed2836001600160a01b03841662001035565b9392505050565b600062000ef1846001600160a01b0385168462001054565b949350505050565b600062000ed2836001600160a01b03841662001073565b600e5460e0516040516370a0823160e01b81523060048201526000926001600160601b0316916001600160a01b0316906370a0823190602401602060405180830381865afa15801562000f67573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000f8d9190620020a0565b62000f999190620020ba565b905090565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b0390811663a9059cbb60e01b1790915262000ff69185916200108116565b505050565b600062000e978262001152565b600080806200101885856200115d565b600081815260029690960160205260409095205494959350505050565b6000818152600283016020526040812081905562000ed283836200116b565b6000828152600284016020526040812082905562000ef1848462001179565b600062000ed2838362001187565b6040805180820190915260208082527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656490820152600090620010d0906001600160a01b038516908490620011a0565b80519091501562000ff65780806020019051810190620010f19190620020dd565b62000ff65760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840162000084565b600062000e97825490565b600062000ed28383620011b1565b600062000ed28383620011de565b600062000ed28383620012e9565b6000818152600183016020526040812054151562000ed2565b606062000ef184846000856200133b565b6000826000018281548110620011cb57620011cb62001de1565b9060005260206000200154905092915050565b60008181526001830160205260408120548015620012d75760006200120560018362001f80565b85549091506000906200121b9060019062001f80565b9050808214620012875760008660000182815481106200123f576200123f62001de1565b906000526020600020015490508087600001848154811062001265576200126562001de1565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806200129b576200129b620020fb565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000e97565b600091505062000e97565b5092915050565b6000818152600183016020526040812054620013325750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000e97565b50600062000e97565b6060824710156200139e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840162000084565b600080866001600160a01b03168587604051620013bc919062002137565b60006040518083038185875af1925050503d8060008114620013fb576040519150601f19603f3d011682016040523d82523d6000602084013e62001400565b606091505b50909250905062001414878383876200141f565b979650505050505050565b60608315620014935782516000036200148b576001600160a01b0385163b6200148b5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640162000084565b508162000ef1565b62000ef18383815115620014aa5781518083602001fd5b8060405162461bcd60e51b815260040162000084919062002155565b634e487b7160e01b600052604160045260246000fd5b6040516101a081016001600160401b0381118282101715620015025762001502620014c6565b60405290565b60405160a081016001600160401b0381118282101715620015025762001502620014c6565b60405160e081016001600160401b0381118282101715620015025762001502620014c6565b604080519081016001600160401b0381118282101715620015025762001502620014c6565b60405161010081016001600160401b0381118282101715620015025762001502620014c6565b604051601f8201601f191681016001600160401b0381118282101715620015c857620015c8620014c6565b604052919050565b80516001600160a01b0381168114620015e857600080fd5b919050565b80516001600160401b0381168114620015e857600080fd5b805161ffff81168114620015e857600080fd5b805163ffffffff81168114620015e857600080fd5b80518015158114620015e857600080fd5b60006101a082840312156200165257600080fd5b6200165c620014dc565b90506200166982620015d0565b8152620016796020830162001605565b60208201526200168c6040830162001618565b60408201526200169f6060830162001605565b6060820152620016b26080830162001618565b6080820152620016c560a0830162001605565b60a0820152620016d860c0830162001605565b60c0820152620016eb60e08301620015d0565b60e08201526101006200170081840162001618565b908201526101206200171483820162001618565b908201526101406200172883820162001605565b908201526101606200173c83820162001618565b90820152610180620017508382016200162d565b9082015292915050565b80516001600160801b0381168114620015e857600080fd5b6000606082840312156200178557600080fd5b604051606081016001600160401b0381118282101715620017aa57620017aa620014c6565b604052905080620017bb836200162d565b8152620017cb602084016200175a565b6020820152620017de604084016200175a565b60408201525092915050565b60006001600160401b03821115620018065762001806620014c6565b5060051b60200190565b600082601f8301126200182257600080fd5b815160206200183b6200183583620017ea565b6200159d565b82815260a092830285018201928282019190878511156200185b57600080fd5b8387015b85811015620018e85781818a031215620018795760008081fd5b6200188362001508565b6200188e82620015d0565b81526200189d86830162001618565b868201526040620018b0818401620015ed565b908201526060620018c3838201620015ed565b908201526080620018d68382016200162d565b9082015284529284019281016200185f565b5090979650505050505050565b600082601f8301126200190757600080fd5b815160206200191a6200183583620017ea565b82815260e092830285018201928282019190878511156200193a57600080fd5b8387015b85811015620018e85781818a031215620019585760008081fd5b620019626200152d565b6200196d82620015d0565b81526200197c86830162001618565b8682015260406200198f81840162001618565b908201526060620019a283820162001605565b908201526080620019b583820162001618565b9082015260a0620019c883820162001618565b9082015260c0620019db8382016200162d565b9082015284529284019281016200193e565b600082601f830112620019ff57600080fd5b8151602062001a126200183583620017ea565b82815260069290921b8401810191818101908684111562001a3257600080fd5b8286015b8481101562001a88576040818903121562001a515760008081fd5b62001a5b62001552565b62001a6682620015d0565b815262001a7585830162001605565b8186015283529183019160400162001a36565b509695505050505050565b60008060008060008086880361036081121562001aaf57600080fd5b6101008082121562001ac057600080fd5b62001aca62001577565b915062001ad789620015d0565b825262001ae760208a01620015ed565b602083015262001afa60408a01620015ed565b604083015262001b0d60608a01620015ed565b606083015260808901516001600160601b038116811462001b2d57600080fd5b608083015262001b4060a08a01620015d0565b60a083015262001b5360c08a01620015d0565b60c083015262001b6660e08a01620015d0565b60e083015281975062001b7c8a828b016200163e565b9650505062001b90886102a0890162001772565b6103008801519094506001600160401b038082111562001baf57600080fd5b62001bbd8a838b0162001810565b945061032089015191508082111562001bd557600080fd5b62001be38a838b01620018f5565b935061034089015191508082111562001bfb57600080fd5b5062001c0a89828a01620019ed565b9150509295509295509295565b80516001600160a01b03168252602081015162001c3a602084018261ffff169052565b50604081015162001c53604084018263ffffffff169052565b50606081015162001c6a606084018261ffff169052565b50608081015162001c83608084018263ffffffff169052565b5060a081015162001c9a60a084018261ffff169052565b5060c081015162001cb160c084018261ffff169052565b5060e081015162001ccd60e08401826001600160a01b03169052565b506101008181015163ffffffff90811691840191909152610120808301518216908401526101408083015161ffff16908401526101608083015190911690830152610180908101511515910152565b82516001600160a01b031681526020808401516001600160401b0390811691830191909152604080850151821690830152606080850151918216908301526102a082019050608084015162001d7c60808401826001600160601b03169052565b5060a084015162001d9860a08401826001600160a01b03169052565b5060c084015162001db460c08401826001600160a01b03169052565b5060e084015162001dd060e08401826001600160a01b03169052565b5062000ed261010083018462001c17565b634e487b7160e01b600052603260045260246000fd5b602080825282518282018190526000919060409081850190868401855b8281101562001e7957815180516001600160a01b031685528681015163ffffffff1687860152858101516001600160401b03908116878701526060808301519091169086015260809081015115159085015260a0909301929085019060010162001e14565b5091979650505050505050565b602080825282518282018190526000919060409081850190868401855b8281101562001e7957815180516001600160a01b031685528681015163ffffffff908116888701528682015181168787015260608083015161ffff169087015260808083015182169087015260a0808301519091169086015260c09081015115159085015260e0909301929085019060010162001ea3565b6020808252825182820181905260009190848201906040850190845b8181101562001f5e5783516001600160a01b03168352928401929184019160010162001f37565b50909695505050505050565b634e487b7160e01b600052601160045260246000fd5b8181038181111562000e975762000e9762001f6a565b60008162001fa85762001fa862001f6a565b506000190190565b63ffffffff818116838216019080821115620012e257620012e262001f6a565b6000604080830163ffffffff8616845260206040602086015281865180845260608701915060208801935060005b818110156200203257845180516001600160a01b0316845284015161ffff1684840152938301939185019160010162001ffe565b509098975050505050505050565b808202811582820484141762000e975762000e9762001f6a565b6000826200207857634e487b7160e01b600052601260045260246000fd5b500490565b6001600160601b03828116828216039080821115620012e257620012e262001f6a565b600060208284031215620020b357600080fd5b5051919050565b8181036000831280158383131683831282161715620012e257620012e262001f6a565b600060208284031215620020f057600080fd5b62000ed2826200162d565b634e487b7160e01b600052603160045260246000fd5b60005b838110156200212e57818101518382015260200162002114565b50506000910152565b600082516200214b81846020870162002111565b9190910192915050565b60208152600082518060208401526200217681604085016020870162002111565b601f01601f19169190910160400192915050565b60805160a05160c05160e0516101005161012051610140516101605161018051615f75620022b86000396000818161036101528181610efa0152613786015260008181610332015281816116f80152613757015260008181610303015281816113ae0152818161141301528181611c7401528181611d02015261372801526000818161026f01528181610a30015281816118200152818161222201528181612b42015261369401526000818161023f01528181611dd30152613664015260008181610210015281816110a50152818161162401528181611a4101528181611b42015281816126dd015281816136350152613a240152600081816102cf01528181611c0e01526136f401526000818161029f0152818161283801526136c4015260006124b60152615f756000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c80637437ff9f116100f9578063c92b283211610097578063eff7cc4811610071578063eff7cc48146109de578063f25561fd146109e6578063f2fde38b146109f9578063fbca3b7414610a0c57600080fd5b8063c92b2832146109b0578063d09dc339146109c3578063df0aa9e9146109cb57600080fd5b8063856c8247116100d3578063856c8247146108825780638da5cb5b146108955780639a113c36146108a6578063b06d41bc1461099a57600080fd5b80637437ff9f146106c057806376f6ae761461086757806379ba50971461087a57600080fd5b806348a98aa411610166578063549e946f11610140578063549e946f1461066957806354b714681461067c578063599f64311461069c578063704b6c02146106ad57600080fd5b806348a98aa4146105c7578063504bffe0146105f2578063546719cd1461060557600080fd5b806320487ded1161019757806320487ded146105705780634120fccd146105915780634816f4f7146105b257600080fd5b806306285c69146101be5780631772047e146103a7578063181f5a7714610527575b600080fd5b6103916040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101919091526040518061010001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b60405161039e9190614b86565b60405180910390f35b6104bb6103b5366004614bba565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152506001600160a01b03166000908152600c6020908152604091829020825160e081018452905463ffffffff8082168352640100000000820481169383019390935261ffff68010000000000000000820416938201939093526a01000000000000000000008304821660608201526e0100000000000000000000000000008304909116608082015260ff720100000000000000000000000000000000000083048116151560a0830152730100000000000000000000000000000000000000909204909116151560c082015290565b60405161039e9190600060e08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015260c0830151151560c083015292915050565b6105636040518060400160405280601381526020017f45564d3245564d4f6e52616d7020312e352e300000000000000000000000000081525081565b60405161039e9190614c27565b61058361057e366004614c68565b610a2c565b60405190815260200161039e565b610599610e82565b60405167ffffffffffffffff909116815260200161039e565b6105c56105c0366004614e76565b610ea9565b005b6105da6105d5366004614fb3565b610ebf565b6040516001600160a01b03909116815260200161039e565b6105c5610600366004614fec565b610f6e565b61060d610f82565b60405161039e919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b6105c56106773660046150ea565b611014565b600e546040516bffffffffffffffffffffffff909116815260200161039e565b6002546001600160a01b03166105da565b6105c56106bb366004614bba565b61118d565b61085a604080516101a081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081019190915250604080516101a0810182526005546001600160a01b038082168352600160a01b820461ffff9081166020850152760100000000000000000000000000000000000000000000830463ffffffff908116958501959095527a0100000000000000000000000000000000000000000000000000008304811660608501527c0100000000000000000000000000000000000000000000000000000000928390048516608085015260065480821660a086015262010000808204831660c0870152640100000000820490931660e08601527801000000000000000000000000000000000000000000000000810486166101008601529290920484166101208401526007549182166101408401528104909216610160820152660100000000000090910460ff16151561018082015290565b60405161039e9190615206565b6105c5610875366004615215565b611257565b6105c56112ba565b610599610890366004614bba565b611383565b6000546001600160a01b03166105da565b6109506108b4366004614bba565b604080516080810182526000808252602082018190529181018290526060810191909152506001600160a01b03166000908152600b60209081526040918290208251608081018452905463ffffffff8116825267ffffffffffffffff64010000000082048116938301939093526c0100000000000000000000000081049092169281019290925260ff600160a01b909104161515606082015290565b60408051825163ffffffff16815260208084015167ffffffffffffffff9081169183019190915283830151169181019190915260609182015115159181019190915260800161039e565b6109a261147e565b60405161039e9291906152de565b6105c56109be366004615320565b611579565b6105836115e1565b6105836109d936600461538e565b6116a1565b6105c561252a565b6105c56109f43660046153fa565b6127af565b6105c5610a07366004614bba565b6127c0565b610a1f610a1a3660046154f9565b6127d1565b60405161039e9190615516565b60007f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168367ffffffffffffffff1614610aac576040517fd9a9cd6800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b6000610ac3610abe6080850185615563565b612805565b9050610af3610ad56020850185615563565b8351909150610ae760408701876155c8565b90508460200151612992565b6000600b81610b086080870160608801614bba565b6001600160a01b0316815260208082019290925260409081016000208151608081018352905463ffffffff81168252640100000000810467ffffffffffffffff908116948301949094526c01000000000000000000000000810490931691810191909152600160a01b90910460ff16151560608201819052909150610bd557610b976080850160608601614bba565b6040517fa7499d200000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610aa3565b600654600090819064010000000090046001600160a01b031663ffdb4b37610c036080890160608a01614bba565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015267ffffffffffffffff8a1660248201526044016040805180830381865afa158015610c6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c92919061565c565b90925090506000808080610ca960408b018b6155c8565b90501115610ce357610cd7610cc460808b0160608c01614bba565b86610cd260408d018d6155c8565b612af7565b91945092509050610cff565b8551610cfc9063ffffffff16662386f26fc100006156a5565b92505b60065460009062010000900461ffff1615610d5357610d506dffffffffffffffffffffffffffff607087901c16610d3960208d018d615563565b9050610d4860408e018e6155c8565b905085612ec6565b90505b60208781015160055460009267ffffffffffffffff9092169163ffffffff8716917a010000000000000000000000000000000000000000000000000000900461ffff1690610da3908f018f615563565b610dae9291506156a5565b6005548c51610ddd91760100000000000000000000000000000000000000000000900463ffffffff16906156bc565b610de791906156bc565b610df191906156bc565b610e0b906dffffffffffffffffffffffffffff89166156a5565b610e1591906156a5565b9050867bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1682828a6040015167ffffffffffffffff1688610e5291906156a5565b610e5c91906156bc565b610e6691906156bc565b610e7091906156cf565b99505050505050505050505b92915050565b600e54600090610ea490600160801b900467ffffffffffffffff1660016156f1565b905090565b610eb1612f97565b610ebb8282612fef565b5050565b6040517fbbe4f6db0000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063bbe4f6db90602401602060405180830381865afa158015610f43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f679190615712565b9392505050565b610f76613371565b610f7f816133cb565b50565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526003546fffffffffffffffffffffffffffffffff8082168352600160801b80830463ffffffff166020850152600160a01b90920460ff161515938301939093526004548084166060840152049091166080820152610ea4906137c0565b61101c612f97565b6001600160a01b03811661105c576040517f232cb97f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006110666115e1565b905060008112156110a3576040517f02075e0000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316036110f5576110f06001600160a01b0384168383613872565b505050565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526110f09083906001600160a01b038616906370a0823190602401602060405180830381865afa158015611158573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061117c919061572f565b6001600160a01b0386169190613872565b6000546001600160a01b031633148015906111b357506002546001600160a01b03163314155b156111ea576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c906020015b60405180910390a150565b61125f612f97565b610ebb8282808060200260200160405190810160405280939291908181526020016000905b828210156112b0576112a160408302860136819003810190615748565b81526020019060010190611284565b50505050506138f2565b6001546001600160a01b031633146113145760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610aa3565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6001600160a01b038082166000908152600d6020526040812054909167ffffffffffffffff909116907f00000000000000000000000000000000000000000000000000000000000000001615610e7c5780600003610e7c576040517f856c82470000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063856c824790602401602060405180830381865afa15801561145a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f679190615787565b606060008061148d6008613b58565b90508067ffffffffffffffff8111156114a8576114a8614cb8565b6040519080825280602002602001820160405280156114ed57816020015b60408051808201909152600080825260208201528152602001906001900390816114c65790505b50925060005b8181101561155657600080611509600884613b63565b915091506040518060400160405280836001600160a01b031681526020018261ffff16815250868481518110611541576115416157a4565b602090810291909101015250506001016114f3565b5050600e5491926c0100000000000000000000000090920463ffffffff16919050565b6000546001600160a01b0316331480159061159f57506002546001600160a01b03163314155b156115d6576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f7f600382613b81565b600e546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000916bffffffffffffffffffffffff16907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa158015611673573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611697919061572f565b610ea491906157ba565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608086901b1660048201526000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690632cbc26bb90602401602060405180830381865afa15801561173f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061176391906157da565b1561179a576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166117da576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005546001600160a01b0316331461181e576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168567ffffffffffffffff1614611897576040517fd9a9cd6800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610aa3565b60006118a9610abe6080870187615563565b905060006118ba60408701876155c8565b91506118e090506118ce6020880188615563565b90508360000151838560200151612992565b8015611a37576000805b82811015611a25576118ff60408901896155c8565b8281811061190f5761190f6157a4565b90506040020160200135600003611952576040517f5cf0444900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c600061196360408b018b6155c8565b84818110611973576119736157a4565b6119899260206040909202019081019150614bba565b6001600160a01b031681526020810191909152604001600020547201000000000000000000000000000000000000900460ff1615611a1d57611a106119d160408a018a6155c8565b838181106119e1576119e16157a4565b9050604002018036038101906119f791906157f7565b60065464010000000090046001600160a01b0316613d14565b611a1a90836156bc565b91505b6001016118ea565b508015611a3557611a3581613e35565b505b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016611a716080880160608901614bba565b6001600160a01b031603611ad557600e8054869190600090611aa29084906bffffffffffffffffffffffff16615831565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff160217905550611bfc565b60065464010000000090046001600160a01b03166241e5be611afd6080890160608a01614bba565b60405160e083901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b039182166004820152602481018990527f00000000000000000000000000000000000000000000000000000000000000009091166044820152606401602060405180830381865afa158015611b89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bad919061572f565b600e8054600090611bcd9084906bffffffffffffffffffffffff16615831565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055505b600e546bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811691161115611c69576040517fe5c7a49100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160200151611dbd577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615611dbd576001600160a01b0384166000908152600d602052604081205467ffffffffffffffff169003611dbd576040517f856c82470000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063856c824790602401602060405180830381865afa158015611d49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d6d9190615787565b6001600160a01b0385166000908152600d6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff929092169190911790555b604080516101a08101825267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b03861660208201526000918101611e50611e168a80615563565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613e4292505050565b6001600160a01b03168152602001600e601081819054906101000a900467ffffffffffffffff16611e8090615856565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff168152602001846000015181526020016000151581526020018460200151611f2a576001600160a01b0387166000908152600d602052604081208054909190611f009067ffffffffffffffff16615856565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055611f2d565b60005b67ffffffffffffffff168152602001611f4c60808a0160608b01614bba565b6001600160a01b03168152602001878152602001888060200190611f709190615563565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250602001611fb760408a018a6155c8565b808060200260200160405190810160405280939291908181526020016000905b8282101561200357611ff4604083028601368190038101906157f7565b81526020019060010190611fd7565b505050505081526020018367ffffffffffffffff81111561202657612026614cb8565b60405190808252806020026020018201604052801561205957816020015b60608152602001906001900390816120445790505b508152600060209091018190529091505b828110156124af57600061208160408a018a6155c8565b83818110612091576120916157a4565b9050604002018036038101906120a791906157f7565b905060006120b98b8360000151610ebf565b90506001600160a01b038116158061216f57506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa158015612149573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061216d91906157da565b155b156121b45781516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610aa3565b6000816001600160a01b0316639a4575b96040518060a001604052808e80600001906121e09190615563565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166020808301919091526001600160a01b03808f16604080850191909152918901516060840152885116608090920191909152517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526122ab919060040161587d565b6000604051808303816000875af11580156122ca573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526122f2919081019061594a565b9050602063ffffffff1681602001515111156123895782516001600160a01b03166000908152600c602090815260409091205490820151516e01000000000000000000000000000090910463ffffffff1610156123895782516040517f36f536ca0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610aa3565b805161239490613e42565b5060408051608081019091526001600160a01b03831660a08201528060c0810160408051808303601f190181529181529082528351602080840191909152808501518383015286516001600160a01b03166000908152600c9091522054606090910190730100000000000000000000000000000000000000900460ff166124295760075462010000900463ffffffff16612458565b84516001600160a01b03166000908152600c60205260409020546a0100000000000000000000900463ffffffff165b63ffffffff16905260405161247091906020016159db565b6040516020818303038152906040528561016001518581518110612496576124966157a4565b602002602001018190525050505080600101905061206a565b506124da817f0000000000000000000000000000000000000000000000000000000000000000613ee8565b6101808201526040517fd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd90612510908390615aef565b60405180910390a16101800151925050505b949350505050565b6000546001600160a01b0316331461258f576002546001600160a01b0316331461258f57612559600833614043565b61258f576040517f195db95800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e546c01000000000000000000000000900463ffffffff1660008190036125e3576040517f990e30bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e546bffffffffffffffffffffffff168181101561262e576040517f8d0f71d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006126386115e1565b1215612670576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600061267d6008613b58565b905060005b8181101561276c57600080612698600884613b63565b90925090506000876126b8836bffffffffffffffffffffffff8a166156a5565b6126c291906156cf565b90506126ce8187615c24565b95506127126001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016846bffffffffffffffffffffffff8416613872565b6040516bffffffffffffffffffffffff821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a2505050806001019050612682565b5050600e80547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff929092169190911790555050565b6127b7612f97565b610f7f81614058565b6127c8613371565b610f7f816141ca565b60606040517f9e7177c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260008082526020820152600082900361286657506040805180820190915267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815260006020820152610e7c565b60006128728385615c49565b90507fe7e230f0000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216016128df576128ca8360048187615c91565b8101906128d79190615cbb565b915050610e7c565b7f6859a837000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216016129605760408051808201909152806129408560048189615c91565b81019061294d9190615ce7565b815260006020909101529150610e7c9050565b6040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006547801000000000000000000000000000000000000000000000000900463ffffffff16808511156129fb576040517f869337890000000000000000000000000000000000000000000000000000000081526004810182905260248101869052604401610aa3565b6006547c0100000000000000000000000000000000000000000000000000000000900463ffffffff16841115612a5d576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600554600160a01b900461ffff16831115612aa4576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81612af0576007546601000000000000900460ff1615612af0576040517fee433e9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b6000808083815b81811015612eba576000878783818110612b1a57612b1a6157a4565b905060400201803603810190612b3091906157f7565b905060006001600160a01b0316612b6b7f00000000000000000000000000000000000000000000000000000000000000008360000151610ebf565b6001600160a01b031603612bb95780516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610aa3565b80516001600160a01b03166000908152600c6020908152604091829020825160e081018452905463ffffffff8082168352640100000000820481169383019390935261ffff68010000000000000000820416938201939093526a01000000000000000000008304821660608201526e0100000000000000000000000000008304909116608082015260ff720100000000000000000000000000000000000083048116151560a0830152730100000000000000000000000000000000000000909204909116151560c08201819052612cde57600754612ca29061ffff16662386f26fc100006156a5565b612cac90886156bc565b600754909750612cc89062010000900463ffffffff1687615d00565b9550612cd5602086615d00565b94505050612eb2565b604081015160009061ffff1615612e025760008c6001600160a01b031684600001516001600160a01b031614612da55760065484516040517f4ab35b0b0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526401000000009092041690634ab35b0b90602401602060405180830381865afa158015612d7a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d9e9190615d1d565b9050612da8565b508a5b620186a0836040015161ffff16612dea8660200151847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1661428b90919063ffffffff16565b612df491906156a5565b612dfe91906156cf565b9150505b6060820151612e119088615d00565b9650816080015186612e239190615d00565b8251909650600090612e429063ffffffff16662386f26fc100006156a5565b905080821015612e6157612e56818a6156bc565b985050505050612eb2565b6000836020015163ffffffff16662386f26fc10000612e8091906156a5565b905080831115612ea057612e94818b6156bc565b99505050505050612eb2565b612eaa838b6156bc565b995050505050505b600101612afe565b50509450945094915050565b60008063ffffffff8316612edc610180866156a5565b612ee8876102206156bc565b612ef291906156bc565b612efc91906156bc565b6005546006549192506000917c010000000000000000000000000000000000000000000000000000000090910463ffffffff1690612f3e9061ffff16846156a5565b612f4891906156bc565b60065490915062010000900461ffff16612f726dffffffffffffffffffffffffffff8916836156a5565b612f7c91906156a5565b612f8c90655af3107a40006156a5565b979650505050505050565b6000546001600160a01b03163314612fed576002546001600160a01b03163314612fed576040517ffbdb8e5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60005b825181101561328357600083828151811061300f5761300f6157a4565b60200260200101519050602063ffffffff168160a0015163ffffffff16101561308257805160a08201516040517f24ecdc020000000000000000000000000000000000000000000000000000000081526001600160a01b03909216600483015263ffffffff166024820152604401610aa3565b6040805160e08101825260208381015163ffffffff908116835284840151811682840190815260608087015161ffff9081168688019081526080808a0151861693880193845260a0808b0151871691890191825260c0808c01511515918a019182526001908a018181529b516001600160a01b03166000908152600c9099529990972097518854955192519451915197519a519087167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009096169590951764010000000092871692909202919091177fffffffffffffffffffffffffffffffffffff000000000000ffffffffffffffff166801000000000000000093909216929092027fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff16176a010000000000000000000091841691909102177fffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffffff166e01000000000000000000000000000093909216929092027fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff1617720100000000000000000000000000000000000094151594909402939093177fffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffff16730100000000000000000000000000000000000000931515939093029290921790915501612ff2565b507ff5791bc457b3bb990493cf5f655db46c25ccf5764c9b99b8969b4c72ea7df9d0826040516132b39190615d38565b60405180910390a160005b815181101561332e57600c60008383815181106132dd576132dd6157a4565b6020908102919091018101516001600160a01b0316825281019190915260400160002080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556001016132be565b50805115610ebb577ffb95a0042158e60a33e7b5bec100f3d95407b1a71bee6633bd54b8887449750b816040516133659190615516565b60405180910390a15050565b6000546001600160a01b03163314612fed5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610aa3565b60e08101516001600160a01b031661340f576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160000160146101000a81548161ffff021916908361ffff16021790555060408201518160000160166101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001601a6101000a81548161ffff021916908361ffff160217905550608082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160010160006101000a81548161ffff021916908361ffff16021790555060c08201518160010160026101000a81548161ffff021916908361ffff16021790555060e08201518160010160046101000a8154816001600160a01b0302191690836001600160a01b031602179055506101008201518160010160186101000a81548163ffffffff021916908363ffffffff16021790555061012082015181600101601c6101000a81548163ffffffff021916908363ffffffff1602179055506101408201518160020160006101000a81548161ffff021916908361ffff1602179055506101608201518160020160026101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160020160066101000a81548160ff0219169083151502179055509050507f45b5ad483aa608464c2c7f278bd413d284d7790cdc836e40652e23a0277082206040518061010001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152508260405161124c929190615dd8565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261384e82606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff16426138329190615df5565b85608001516fffffffffffffffffffffffffffffffff166142c8565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526110f09084906142f0565b8051604081111561392f576040517fb5a10cfa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e546c01000000000000000000000000900463ffffffff161561398357600e5463ffffffff6c010000000000000000000000008204166bffffffffffffffffffffffff909116106139835761398361252a565b600061398f6008613b58565b90505b80156139d15760006139b06139a8600184615df5565b600890613b63565b5090506139be6008826143d5565b5050806139ca90615e08565b9050613992565b506000805b82811015613ad95760008482815181106139f2576139f26157a4565b60200260200101516000015190506000858381518110613a1457613a146157a4565b60200260200101516020015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b03161480613a6957506001600160a01b038216155b15613aab576040517f4de938d10000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610aa3565b613abb60088361ffff84166143ea565b50613aca61ffff821685615d00565b935050508060010190506139d6565b50600e80547fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd2490613b4b9083908690615e3d565b60405180910390a1505050565b6000610e7c82614400565b6000808080613b72868661440b565b909450925050505b9250929050565b8154600090613b9d90600160801b900463ffffffff1642615df5565b90508015613c255760018301548354613bd8916fffffffffffffffffffffffffffffffff808216928116918591600160801b909104166142c8565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff000000000000000000000000000000000000000090911617600160801b4263ffffffff16021783555b60208201518354613c4b916fffffffffffffffffffffffffffffffff9081169116614436565b835483511515600160a01b027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff9283161717845560208301516040808501518316600160801b0291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990613b4b9084908151151581526020808301516fffffffffffffffffffffffffffffffff90811691830191909152604092830151169181019190915260600190565b81516040517fd02641a00000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009182919084169063d02641a0906024016040805180830381865afa158015613d7a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d9e9190615e5c565b5190507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116600003613e075783516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610aa3565b6020840151612522907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83169061428b565b610f7f600382600061444c565b60008151602014613e8157816040517f8d666f60000000000000000000000000000000000000000000000000000000008152600401610aa39190614c27565b600082806020019051810190613e97919061572f565b90506001600160a01b03811180613eaf575061040081105b15610e7c57826040517f8d666f60000000000000000000000000000000000000000000000000000000008152600401610aa39190614c27565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b6101000151604051602001613f7e9897969594939291906001600160a01b039889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b6040516020818303038152906040528051906020012085610120015180519060200120866101400151604051602001613fb79190615e8f565b60405160208183030381529060405280519060200120876101600151604051602001613fe39190615ea2565b60408051601f198184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b6000610f67836001600160a01b038416614756565b60005b815181101561419a576000828281518110614078576140786157a4565b60209081029190910181015160408051608080820183528385015163ffffffff90811683528385015167ffffffffffffffff908116848801908152606080880151831686880190815294880151151590860190815296516001600160a01b03166000908152600b90985294909620925183549451925195511515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9688166c0100000000000000000000000002969096167fffffffffffffffffffffff000000000000000000ffffffffffffffffffffffff93909716640100000000027fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009095169116179290921791909116929092171790555060010161405b565b507f067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e8160405161124c9190615eb5565b336001600160a01b038216036142225760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610aa3565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000670de0b6b3a76400006142be837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff86166156a5565b610f6791906156cf565b60006142e7856142d884866156a5565b6142e290876156bc565b614436565b95945050505050565b6000614345826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166147629092919063ffffffff16565b8051909150156110f0578080602001905181019061436391906157da565b6110f05760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610aa3565b6000610f67836001600160a01b038416614771565b6000612522846001600160a01b0385168461478e565b6000610e7c826147ab565b6000808061441985856147b5565b600081815260029690960160205260409095205494959350505050565b60008183106144455781610f67565b5090919050565b8254600160a01b900460ff161580614462575081155b1561446c57505050565b825460018401546fffffffffffffffffffffffffffffffff808316929116906000906144a590600160801b900463ffffffff1642615df5565b9050801561454b57818311156144e7576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600186015461451490839085908490600160801b90046fffffffffffffffffffffffffffffffff166142c8565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16600160801b4263ffffffff160217875592505b848210156145e8576001600160a01b03841661459d576040517ff94ebcd10000000000000000000000000000000000000000000000000000000081526004810183905260248101869052604401610aa3565b6040517f1a76572a00000000000000000000000000000000000000000000000000000000815260048101839052602481018690526001600160a01b0385166044820152606401610aa3565b848310156146d457600186810154600160801b90046fffffffffffffffffffffffffffffffff1690600090829061461f9082615df5565b614629878a615df5565b61463391906156bc565b61463d91906156cf565b90506001600160a01b038616614689576040517f15279c080000000000000000000000000000000000000000000000000000000081526004810182905260248101869052604401610aa3565b6040517fd0c8d23a00000000000000000000000000000000000000000000000000000000815260048101829052602481018690526001600160a01b0387166044820152606401610aa3565b6146de8584615df5565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b6000610f6783836147c1565b606061252284846000856147d9565b60008181526002830160205260408120819055610f6783836148c0565b6000828152600284016020526040812082905561252284846148cc565b6000610e7c825490565b6000610f6783836148d8565b60008181526001830160205260408120541515610f67565b6060824710156148515760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610aa3565b600080866001600160a01b0316858760405161486d9190615f36565b60006040518083038185875af1925050503d80600081146148aa576040519150601f19603f3d011682016040523d82523d6000602084013e6148af565b606091505b5091509150612f8c87838387614902565b6000610f67838361497b565b6000610f678383614a75565b60008260000182815481106148ef576148ef6157a4565b9060005260206000200154905092915050565b6060831561497157825160000361496a576001600160a01b0385163b61496a5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610aa3565b5081612522565b6125228383614ac4565b60008181526001830160205260408120548015614a6457600061499f600183615df5565b85549091506000906149b390600190615df5565b9050808214614a185760008660000182815481106149d3576149d36157a4565b90600052602060002001549050808760000184815481106149f6576149f66157a4565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080614a2957614a29615f52565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610e7c565b6000915050610e7c565b5092915050565b6000818152600183016020526040812054614abc57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610e7c565b506000610e7c565b815115614ad45781518083602001fd5b8060405162461bcd60e51b8152600401610aa39190614c27565b6001600160a01b03808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015280606085015116606086015250506bffffffffffffffffffffffff60808301511660808401528060a08301511660a08401525060c0810151614b6b60c08401826001600160a01b03169052565b5060e08101516110f060e08401826001600160a01b03169052565b6101008101610e7c8284614aee565b6001600160a01b0381168114610f7f57600080fd5b8035614bb581614b95565b919050565b600060208284031215614bcc57600080fd5b8135610f6781614b95565b60005b83811015614bf2578181015183820152602001614bda565b50506000910152565b60008151808452614c13816020860160208601614bd7565b601f01601f19169290920160200192915050565b602081526000610f676020830184614bfb565b67ffffffffffffffff81168114610f7f57600080fd5b600060a08284031215614c6257600080fd5b50919050565b60008060408385031215614c7b57600080fd5b8235614c8681614c3a565b9150602083013567ffffffffffffffff811115614ca257600080fd5b614cae85828601614c50565b9150509250929050565b634e487b7160e01b600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715614cf157614cf1614cb8565b60405290565b6040516101a0810167ffffffffffffffff81118282101715614cf157614cf1614cb8565b60405160a0810167ffffffffffffffff81118282101715614cf157614cf1614cb8565b6040805190810167ffffffffffffffff81118282101715614cf157614cf1614cb8565b604051601f8201601f1916810167ffffffffffffffff81118282101715614d8a57614d8a614cb8565b604052919050565b600067ffffffffffffffff821115614dac57614dac614cb8565b5060051b60200190565b63ffffffff81168114610f7f57600080fd5b8035614bb581614db6565b803561ffff81168114614bb557600080fd5b8015158114610f7f57600080fd5b8035614bb581614de5565b600082601f830112614e0f57600080fd5b81356020614e24614e1f83614d92565b614d61565b8083825260208201915060208460051b870101935086841115614e4657600080fd5b602086015b84811015614e6b578035614e5e81614b95565b8352918301918301614e4b565b509695505050505050565b6000806040808486031215614e8a57600080fd5b833567ffffffffffffffff80821115614ea257600080fd5b818601915086601f830112614eb657600080fd5b81356020614ec6614e1f83614d92565b82815260e0928302850182019282820191908b851115614ee557600080fd5b958301955b84871015614f8e5780878d031215614f025760008081fd5b614f0a614cce565b8735614f1581614b95565b815287850135614f2481614db6565b8186015287890135614f3581614db6565b818a01526060614f46898201614dd3565b90820152608088810135614f5981614db6565b9082015260a0614f6a898201614dc8565b9082015260c0614f7b898201614df3565b9082015283529586019591830191614eea565b5097505087013593505080831115614fa557600080fd5b5050614cae85828601614dfe565b60008060408385031215614fc657600080fd5b8235614fd181614c3a565b91506020830135614fe181614b95565b809150509250929050565b60006101a08284031215614fff57600080fd5b615007614cf7565b61501083614baa565b815261501e60208401614dd3565b602082015261502f60408401614dc8565b604082015261504060608401614dd3565b606082015261505160808401614dc8565b608082015261506260a08401614dd3565b60a082015261507360c08401614dd3565b60c082015261508460e08401614baa565b60e0820152610100615097818501614dc8565b908201526101206150a9848201614dc8565b908201526101406150bb848201614dd3565b908201526101606150cd848201614dc8565b908201526101806150df848201614df3565b908201529392505050565b600080604083850312156150fd57600080fd5b8235614fd181614b95565b80516001600160a01b03168252602081015161512a602084018261ffff169052565b506040810151615142604084018263ffffffff169052565b506060810151615158606084018261ffff169052565b506080810151615170608084018263ffffffff169052565b5060a081015161518660a084018261ffff169052565b5060c081015161519c60c084018261ffff169052565b5060e08101516151b760e08401826001600160a01b03169052565b506101008181015163ffffffff90811691840191909152610120808301518216908401526101408083015161ffff16908401526101608083015190911690830152610180908101511515910152565b6101a08101610e7c8284615108565b6000806020838503121561522857600080fd5b823567ffffffffffffffff8082111561524057600080fd5b818501915085601f83011261525457600080fd5b81358181111561526357600080fd5b8660208260061b850101111561527857600080fd5b60209290920196919550909350505050565b60008151808452602080850194506020840160005b838110156152d357815180516001600160a01b0316885283015161ffff16838801526040909601959082019060010161529f565b509495945050505050565b6040815260006152f1604083018561528a565b90508260208301529392505050565b80356fffffffffffffffffffffffffffffffff81168114614bb557600080fd5b60006060828403121561533257600080fd5b6040516060810181811067ffffffffffffffff8211171561535557615355614cb8565b604052823561536381614de5565b815261537160208401615300565b602082015261538260408401615300565b60408201529392505050565b600080600080608085870312156153a457600080fd5b84356153af81614c3a565b9350602085013567ffffffffffffffff8111156153cb57600080fd5b6153d787828801614c50565b9350506040850135915060608501356153ef81614b95565b939692955090935050565b6000602080838503121561540d57600080fd5b823567ffffffffffffffff81111561542457600080fd5b8301601f8101851361543557600080fd5b8035615443614e1f82614d92565b81815260a0918202830184019184820191908884111561546257600080fd5b938501935b838510156154ed5780858a03121561547f5760008081fd5b615487614d1b565b853561549281614b95565b8152858701356154a181614db6565b818801526040868101356154b481614c3a565b908201526060868101356154c781614c3a565b908201526080868101356154da81614de5565b9082015283529384019391850191615467565b50979650505050505050565b60006020828403121561550b57600080fd5b8135610f6781614c3a565b6020808252825182820181905260009190848201906040850190845b818110156155575783516001600160a01b031683529284019291840191600101615532565b50909695505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261559857600080fd5b83018035915067ffffffffffffffff8211156155b357600080fd5b602001915036819003821315613b7a57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126155fd57600080fd5b83018035915067ffffffffffffffff82111561561857600080fd5b6020019150600681901b3603821315613b7a57600080fd5b80517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81168114614bb557600080fd5b6000806040838503121561566f57600080fd5b61567883615630565b915061568660208401615630565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b8082028115828204841417610e7c57610e7c61568f565b80820180821115610e7c57610e7c61568f565b6000826156ec57634e487b7160e01b600052601260045260246000fd5b500490565b67ffffffffffffffff818116838216019080821115614a6e57614a6e61568f565b60006020828403121561572457600080fd5b8151610f6781614b95565b60006020828403121561574157600080fd5b5051919050565b60006040828403121561575a57600080fd5b615762614d3e565b823561576d81614b95565b815261577b60208401614dd3565b60208201529392505050565b60006020828403121561579957600080fd5b8151610f6781614c3a565b634e487b7160e01b600052603260045260246000fd5b8181036000831280158383131683831282161715614a6e57614a6e61568f565b6000602082840312156157ec57600080fd5b8151610f6781614de5565b60006040828403121561580957600080fd5b615811614d3e565b823561581c81614b95565b81526020928301359281019290925250919050565b6bffffffffffffffffffffffff818116838216019080821115614a6e57614a6e61568f565b600067ffffffffffffffff8083168181036158735761587361568f565b6001019392505050565b602081526000825160a0602084015261589960c0840182614bfb565b905067ffffffffffffffff602085015116604084015260408401516001600160a01b038082166060860152606086015160808601528060808701511660a086015250508091505092915050565b600082601f8301126158f757600080fd5b815167ffffffffffffffff81111561591157615911614cb8565b6159246020601f19601f84011601614d61565b81815284602083860101111561593957600080fd5b612522826020830160208701614bd7565b60006020828403121561595c57600080fd5b815167ffffffffffffffff8082111561597457600080fd5b908301906040828603121561598857600080fd5b615990614d3e565b82518281111561599f57600080fd5b6159ab878286016158e6565b8252506020830151828111156159c057600080fd5b6159cc878286016158e6565b60208301525095945050505050565b6020815260008251608060208401526159f760a0840182614bfb565b90506020840151601f1980858403016040860152615a158383614bfb565b9250604086015191508085840301606086015250615a338282614bfb565b91505063ffffffff60608501511660808401528091505092915050565b60008151808452602080850194506020840160005b838110156152d357815180516001600160a01b031688528301518388015260409096019590820190600101615a65565b60008282518085526020808601955060208260051b8401016020860160005b84811015615ae257601f19868403018952615ad0838351614bfb565b98840198925090830190600101615ab4565b5090979650505050505050565b60208152615b0a60208201835167ffffffffffffffff169052565b60006020830151615b2660408401826001600160a01b03169052565b5060408301516001600160a01b038116606084015250606083015167ffffffffffffffff8116608084015250608083015160a083015260a0830151615b6f60c084018215159052565b5060c083015167ffffffffffffffff811660e08401525060e0830151610100615ba2818501836001600160a01b03169052565b840151610120848101919091528401516101a061014080860182905291925090615bd06101c0860184614bfb565b9250808601519050601f19610160818786030181880152615bf18584615a50565b945080880151925050610180818786030181880152615c108584615a95565b970151959092019490945250929392505050565b6bffffffffffffffffffffffff828116828216039080821115614a6e57614a6e61568f565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015615c895780818660040360031b1b83161692505b505092915050565b60008085851115615ca157600080fd5b83861115615cae57600080fd5b5050820193919092039150565b600060408284031215615ccd57600080fd5b615cd5614d3e565b82358152602083013561577b81614de5565b600060208284031215615cf957600080fd5b5035919050565b63ffffffff818116838216019080821115614a6e57614a6e61568f565b600060208284031215615d2f57600080fd5b610f6782615630565b602080825282518282018190526000919060409081850190868401855b82811015615dcb57815180516001600160a01b031685528681015163ffffffff908116888701528682015181168787015260608083015161ffff169087015260808083015182169087015260a0808301519091169086015260c09081015115159085015260e09093019290850190600101615d55565b5091979650505050505050565b6102a08101615de78285614aee565b610f67610100830184615108565b81810381811115610e7c57610e7c61568f565b600081615e1757615e1761568f565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b63ffffffff83168152604060208201526000612522604083018461528a565b600060408284031215615e6e57600080fd5b615e76614d3e565b615e7f83615630565b8152602083015161577b81614db6565b602081526000610f676020830184615a50565b602081526000610f676020830184615a95565b602080825282518282018190526000919060409081850190868401855b82811015615dcb57815180516001600160a01b031685528681015163ffffffff16878601528581015167ffffffffffffffff908116878701526060808301519091169086015260809081015115159085015260a09093019290850190600101615ed2565b60008251615f48818460208701614bd7565b9190910192915050565b634e487b7160e01b600052603160045260246000fdfea164736f6c6343000818000a",
}

var EVM2EVMOnRampABI = EVM2EVMOnRampMetaData.ABI

var EVM2EVMOnRampBin = EVM2EVMOnRampMetaData.Bin

func DeployEVM2EVMOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMOnRampStaticConfig, dynamicConfig EVM2EVMOnRampDynamicConfig, rateLimiterConfig RateLimiterConfig, feeTokenConfigs []EVM2EVMOnRampFeeTokenConfigArgs, tokenTransferFeeConfigArgs []EVM2EVMOnRampTokenTransferFeeConfigArgs, nopsAndWeights []EVM2EVMOnRampNopAndWeight) (common.Address, *CustomTransaction, *EVM2EVMOnRamp, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	switch chainId.Uint64() {

	case 324, 280, 300:
		return DeployZkSyncEVM2EVMOnRamp(auth, backend, staticConfig, dynamicConfig, rateLimiterConfig, feeTokenConfigs, tokenTransferFeeConfigArgs, nopsAndWeights)
	}

	parsed, err := EVM2EVMOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMOnRampBin), backend, staticConfig, dynamicConfig, rateLimiterConfig, feeTokenConfigs, tokenTransferFeeConfigArgs, nopsAndWeights)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, &CustomTransaction{Transaction: tx, CustomHash: tx.Hash()}, &EVM2EVMOnRamp{address: address, abi: *parsed, EVM2EVMOnRampCaller: EVM2EVMOnRampCaller{contract: contract}, EVM2EVMOnRampTransactor: EVM2EVMOnRampTransactor{contract: contract}, EVM2EVMOnRampFilterer: EVM2EVMOnRampFilterer{contract: contract}}, nil
}

type EVM2EVMOnRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMOnRampCaller
	EVM2EVMOnRampTransactor
	EVM2EVMOnRampFilterer
}

type EVM2EVMOnRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMOnRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMOnRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMOnRampSession struct {
	Contract     *EVM2EVMOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMOnRampCallerSession struct {
	Contract *EVM2EVMOnRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMOnRampTransactorSession struct {
	Contract     *EVM2EVMOnRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMOnRampRaw struct {
	Contract *EVM2EVMOnRamp
}

type EVM2EVMOnRampCallerRaw struct {
	Contract *EVM2EVMOnRampCaller
}

type EVM2EVMOnRampTransactorRaw struct {
	Contract *EVM2EVMOnRampTransactor
}

func NewEVM2EVMOnRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRamp{address: address, abi: abi, EVM2EVMOnRampCaller: EVM2EVMOnRampCaller{contract: contract}, EVM2EVMOnRampTransactor: EVM2EVMOnRampTransactor{contract: contract}, EVM2EVMOnRampFilterer: EVM2EVMOnRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMOnRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMOnRampCaller, error) {
	contract, err := bindEVM2EVMOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampCaller{contract: contract}, nil
}

func NewEVM2EVMOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMOnRampTransactor, error) {
	contract, err := bindEVM2EVMOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampTransactor{contract: contract}, nil
}

func NewEVM2EVMOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMOnRampFilterer, error) {
	contract, err := bindEVM2EVMOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampFilterer{contract: contract}, nil
}

func bindEVM2EVMOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EVM2EVMOnRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMOnRamp.Contract.EVM2EVMOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.EVM2EVMOnRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.EVM2EVMOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "currentRateLimiterState")

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _EVM2EVMOnRamp.Contract.CurrentRateLimiterState(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _EVM2EVMOnRamp.Contract.CurrentRateLimiterState(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMOnRampDynamicConfig, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(EVM2EVMOnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMOnRampDynamicConfig)).(*EVM2EVMOnRampDynamicConfig)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetDynamicConfig() (EVM2EVMOnRampDynamicConfig, error) {
	return _EVM2EVMOnRamp.Contract.GetDynamicConfig(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetDynamicConfig() (EVM2EVMOnRampDynamicConfig, error) {
	return _EVM2EVMOnRamp.Contract.GetDynamicConfig(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getExpectedNextSequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _EVM2EVMOnRamp.Contract.GetFee(&_EVM2EVMOnRamp.CallOpts, destChainSelector, message)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _EVM2EVMOnRamp.Contract.GetFee(&_EVM2EVMOnRamp.CallOpts, destChainSelector, message)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetFeeTokenConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMOnRampFeeTokenConfig, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getFeeTokenConfig", token)

	if err != nil {
		return *new(EVM2EVMOnRampFeeTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMOnRampFeeTokenConfig)).(*EVM2EVMOnRampFeeTokenConfig)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetFeeTokenConfig(token common.Address) (EVM2EVMOnRampFeeTokenConfig, error) {
	return _EVM2EVMOnRamp.Contract.GetFeeTokenConfig(&_EVM2EVMOnRamp.CallOpts, token)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetFeeTokenConfig(token common.Address) (EVM2EVMOnRampFeeTokenConfig, error) {
	return _EVM2EVMOnRamp.Contract.GetFeeTokenConfig(&_EVM2EVMOnRamp.CallOpts, token)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetNopFeesJuels(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getNopFeesJuels")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetNopFeesJuels() (*big.Int, error) {
	return _EVM2EVMOnRamp.Contract.GetNopFeesJuels(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetNopFeesJuels() (*big.Int, error) {
	return _EVM2EVMOnRamp.Contract.GetNopFeesJuels(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetNops(opts *bind.CallOpts) (GetNops,

	error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getNops")

	outstruct := new(GetNops)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NopsAndWeights = *abi.ConvertType(out[0], new([]EVM2EVMOnRampNopAndWeight)).(*[]EVM2EVMOnRampNopAndWeight)
	outstruct.WeightsTotal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetNops() (GetNops,

	error) {
	return _EVM2EVMOnRamp.Contract.GetNops(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetNops() (GetNops,

	error) {
	return _EVM2EVMOnRamp.Contract.GetNops(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getPoolBySourceToken", arg0, sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOnRamp.Contract.GetPoolBySourceToken(&_EVM2EVMOnRamp.CallOpts, arg0, sourceToken)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOnRamp.Contract.GetPoolBySourceToken(&_EVM2EVMOnRamp.CallOpts, arg0, sourceToken)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getSenderNonce", sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMOnRamp.Contract.GetSenderNonce(&_EVM2EVMOnRamp.CallOpts, sender)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMOnRamp.Contract.GetSenderNonce(&_EVM2EVMOnRamp.CallOpts, sender)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetStaticConfig(opts *bind.CallOpts) (EVM2EVMOnRampStaticConfig, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(EVM2EVMOnRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMOnRampStaticConfig)).(*EVM2EVMOnRampStaticConfig)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetStaticConfig() (EVM2EVMOnRampStaticConfig, error) {
	return _EVM2EVMOnRamp.Contract.GetStaticConfig(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetStaticConfig() (EVM2EVMOnRampStaticConfig, error) {
	return _EVM2EVMOnRamp.Contract.GetStaticConfig(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getSupportedTokens", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _EVM2EVMOnRamp.Contract.GetSupportedTokens(&_EVM2EVMOnRamp.CallOpts, arg0)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _EVM2EVMOnRamp.Contract.GetSupportedTokens(&_EVM2EVMOnRamp.CallOpts, arg0)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMOnRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMOnRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMOnRampTokenTransferFeeConfig, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getTokenTransferFeeConfig", token)

	if err != nil {
		return *new(EVM2EVMOnRampTokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMOnRampTokenTransferFeeConfig)).(*EVM2EVMOnRampTokenTransferFeeConfig)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetTokenTransferFeeConfig(token common.Address) (EVM2EVMOnRampTokenTransferFeeConfig, error) {
	return _EVM2EVMOnRamp.Contract.GetTokenTransferFeeConfig(&_EVM2EVMOnRamp.CallOpts, token)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetTokenTransferFeeConfig(token common.Address) (EVM2EVMOnRampTokenTransferFeeConfig, error) {
	return _EVM2EVMOnRamp.Contract.GetTokenTransferFeeConfig(&_EVM2EVMOnRamp.CallOpts, token)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) LinkAvailableForPayment() (*big.Int, error) {
	return _EVM2EVMOnRamp.Contract.LinkAvailableForPayment(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _EVM2EVMOnRamp.Contract.LinkAvailableForPayment(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) Owner() (common.Address, error) {
	return _EVM2EVMOnRamp.Contract.Owner(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMOnRamp.Contract.Owner(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMOnRamp.Contract.TypeAndVersion(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMOnRamp.Contract.TypeAndVersion(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.AcceptOwnership(&_EVM2EVMOnRamp.TransactOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.AcceptOwnership(&_EVM2EVMOnRamp.TransactOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "forwardFromRouter", destChainSelector, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.ForwardFromRouter(&_EVM2EVMOnRamp.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.ForwardFromRouter(&_EVM2EVMOnRamp.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) PayNops(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "payNops")
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) PayNops() (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.PayNops(&_EVM2EVMOnRamp.TransactOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) PayNops() (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.PayNops(&_EVM2EVMOnRamp.TransactOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setAdmin", newAdmin)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetAdmin(&_EVM2EVMOnRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetAdmin(&_EVM2EVMOnRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMOnRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetDynamicConfig(dynamicConfig EVM2EVMOnRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetDynamicConfig(&_EVM2EVMOnRamp.TransactOpts, dynamicConfig)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetDynamicConfig(dynamicConfig EVM2EVMOnRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetDynamicConfig(&_EVM2EVMOnRamp.TransactOpts, dynamicConfig)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetFeeTokenConfig(opts *bind.TransactOpts, feeTokenConfigArgs []EVM2EVMOnRampFeeTokenConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setFeeTokenConfig", feeTokenConfigArgs)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetFeeTokenConfig(feeTokenConfigArgs []EVM2EVMOnRampFeeTokenConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetFeeTokenConfig(&_EVM2EVMOnRamp.TransactOpts, feeTokenConfigArgs)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetFeeTokenConfig(feeTokenConfigArgs []EVM2EVMOnRampFeeTokenConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetFeeTokenConfig(&_EVM2EVMOnRamp.TransactOpts, feeTokenConfigArgs)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetNops(opts *bind.TransactOpts, nopsAndWeights []EVM2EVMOnRampNopAndWeight) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setNops", nopsAndWeights)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetNops(nopsAndWeights []EVM2EVMOnRampNopAndWeight) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetNops(&_EVM2EVMOnRamp.TransactOpts, nopsAndWeights)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetNops(nopsAndWeights []EVM2EVMOnRampNopAndWeight) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetNops(&_EVM2EVMOnRamp.TransactOpts, nopsAndWeights)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetRateLimiterConfig(&_EVM2EVMOnRamp.TransactOpts, config)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetRateLimiterConfig(&_EVM2EVMOnRamp.TransactOpts, config)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetTokenTransferFeeConfig(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []EVM2EVMOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setTokenTransferFeeConfig", tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetTokenTransferFeeConfig(tokenTransferFeeConfigArgs []EVM2EVMOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetTokenTransferFeeConfig(&_EVM2EVMOnRamp.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetTokenTransferFeeConfig(tokenTransferFeeConfigArgs []EVM2EVMOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetTokenTransferFeeConfig(&_EVM2EVMOnRamp.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.TransferOwnership(&_EVM2EVMOnRamp.TransactOpts, to)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.TransferOwnership(&_EVM2EVMOnRamp.TransactOpts, to)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) WithdrawNonLinkFees(opts *bind.TransactOpts, feeToken common.Address, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "withdrawNonLinkFees", feeToken, to)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) WithdrawNonLinkFees(feeToken common.Address, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.WithdrawNonLinkFees(&_EVM2EVMOnRamp.TransactOpts, feeToken, to)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) WithdrawNonLinkFees(feeToken common.Address, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.WithdrawNonLinkFees(&_EVM2EVMOnRamp.TransactOpts, feeToken, to)
}

type EVM2EVMOnRampAdminSetIterator struct {
	Event *EVM2EVMOnRampAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampAdminSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOnRampAdminSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOnRampAdminSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampAdminSet struct {
	NewAdmin common.Address
	Raw      types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMOnRampAdminSetIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampAdminSetIterator{contract: _EVM2EVMOnRamp.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampAdminSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampAdminSet)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "AdminSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseAdminSet(log types.Log) (*EVM2EVMOnRampAdminSet, error) {
	event := new(EVM2EVMOnRampAdminSet)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "AdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampCCIPSendRequestedIterator struct {
	Event *EVM2EVMOnRampCCIPSendRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampCCIPSendRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampCCIPSendRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOnRampCCIPSendRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOnRampCCIPSendRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampCCIPSendRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampCCIPSendRequested struct {
	Message InternalEVM2EVMMessage
	Raw     types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMOnRampCCIPSendRequestedIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampCCIPSendRequestedIterator{contract: _EVM2EVMOnRamp.contract, event: "CCIPSendRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampCCIPSendRequested) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampCCIPSendRequested)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseCCIPSendRequested(log types.Log) (*EVM2EVMOnRampCCIPSendRequested, error) {
	event := new(EVM2EVMOnRampCCIPSendRequested)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampConfigChangedIterator struct {
	Event *EVM2EVMOnRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOnRampConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOnRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOnRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampConfigChangedIterator{contract: _EVM2EVMOnRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampConfigChanged)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMOnRampConfigChanged, error) {
	event := new(EVM2EVMOnRampConfigChanged)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampConfigSetIterator struct {
	Event *EVM2EVMOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOnRampConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampConfigSet struct {
	StaticConfig  EVM2EVMOnRampStaticConfig
	DynamicConfig EVM2EVMOnRampDynamicConfig
	Raw           types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMOnRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampConfigSetIterator{contract: _EVM2EVMOnRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampConfigSet)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseConfigSet(log types.Log) (*EVM2EVMOnRampConfigSet, error) {
	event := new(EVM2EVMOnRampConfigSet)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampFeeConfigSetIterator struct {
	Event *EVM2EVMOnRampFeeConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampFeeConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampFeeConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOnRampFeeConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOnRampFeeConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampFeeConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampFeeConfigSet struct {
	FeeConfig []EVM2EVMOnRampFeeTokenConfigArgs
	Raw       types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMOnRampFeeConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "FeeConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampFeeConfigSetIterator{contract: _EVM2EVMOnRamp.contract, event: "FeeConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampFeeConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "FeeConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampFeeConfigSet)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "FeeConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseFeeConfigSet(log types.Log) (*EVM2EVMOnRampFeeConfigSet, error) {
	event := new(EVM2EVMOnRampFeeConfigSet)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "FeeConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampNopPaidIterator struct {
	Event *EVM2EVMOnRampNopPaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampNopPaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampNopPaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOnRampNopPaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOnRampNopPaidIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampNopPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampNopPaid struct {
	Nop    common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterNopPaid(opts *bind.FilterOpts, nop []common.Address) (*EVM2EVMOnRampNopPaidIterator, error) {

	var nopRule []interface{}
	for _, nopItem := range nop {
		nopRule = append(nopRule, nopItem)
	}

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "NopPaid", nopRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampNopPaidIterator{contract: _EVM2EVMOnRamp.contract, event: "NopPaid", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchNopPaid(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampNopPaid, nop []common.Address) (event.Subscription, error) {

	var nopRule []interface{}
	for _, nopItem := range nop {
		nopRule = append(nopRule, nopItem)
	}

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "NopPaid", nopRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampNopPaid)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "NopPaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseNopPaid(log types.Log) (*EVM2EVMOnRampNopPaid, error) {
	event := new(EVM2EVMOnRampNopPaid)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "NopPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampNopsSetIterator struct {
	Event *EVM2EVMOnRampNopsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampNopsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampNopsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOnRampNopsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOnRampNopsSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampNopsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampNopsSet struct {
	NopWeightsTotal *big.Int
	NopsAndWeights  []EVM2EVMOnRampNopAndWeight
	Raw             types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterNopsSet(opts *bind.FilterOpts) (*EVM2EVMOnRampNopsSetIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "NopsSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampNopsSetIterator{contract: _EVM2EVMOnRamp.contract, event: "NopsSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchNopsSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampNopsSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "NopsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampNopsSet)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "NopsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseNopsSet(log types.Log) (*EVM2EVMOnRampNopsSet, error) {
	event := new(EVM2EVMOnRampNopsSet)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "NopsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOnRampOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampOwnershipTransferRequestedIterator{contract: _EVM2EVMOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampOwnershipTransferRequested)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMOnRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMOnRampOwnershipTransferRequested)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampOwnershipTransferredIterator struct {
	Event *EVM2EVMOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOnRampOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampOwnershipTransferredIterator{contract: _EVM2EVMOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampOwnershipTransferred)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMOnRampOwnershipTransferred, error) {
	event := new(EVM2EVMOnRampOwnershipTransferred)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampTokenTransferFeeConfigDeletedIterator struct {
	Event *EVM2EVMOnRampTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampTokenTransferFeeConfigDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOnRampTokenTransferFeeConfigDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOnRampTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampTokenTransferFeeConfigDeleted struct {
	Tokens []common.Address
	Raw    types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts) (*EVM2EVMOnRampTokenTransferFeeConfigDeletedIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampTokenTransferFeeConfigDeletedIterator{contract: _EVM2EVMOnRamp.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampTokenTransferFeeConfigDeleted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampTokenTransferFeeConfigDeleted)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*EVM2EVMOnRampTokenTransferFeeConfigDeleted, error) {
	event := new(EVM2EVMOnRampTokenTransferFeeConfigDeleted)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampTokenTransferFeeConfigSetIterator struct {
	Event *EVM2EVMOnRampTokenTransferFeeConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampTokenTransferFeeConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampTokenTransferFeeConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOnRampTokenTransferFeeConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOnRampTokenTransferFeeConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampTokenTransferFeeConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampTokenTransferFeeConfigSet struct {
	TransferFeeConfig []EVM2EVMOnRampTokenTransferFeeConfigArgs
	Raw               types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterTokenTransferFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMOnRampTokenTransferFeeConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "TokenTransferFeeConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampTokenTransferFeeConfigSetIterator{contract: _EVM2EVMOnRamp.contract, event: "TokenTransferFeeConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchTokenTransferFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampTokenTransferFeeConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "TokenTransferFeeConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampTokenTransferFeeConfigSet)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "TokenTransferFeeConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseTokenTransferFeeConfigSet(log types.Log) (*EVM2EVMOnRampTokenTransferFeeConfigSet, error) {
	event := new(EVM2EVMOnRampTokenTransferFeeConfigSet)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "TokenTransferFeeConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampTokensConsumedIterator struct {
	Event *EVM2EVMOnRampTokensConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampTokensConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampTokensConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(EVM2EVMOnRampTokensConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *EVM2EVMOnRampTokensConsumedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampTokensConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampTokensConsumed struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterTokensConsumed(opts *bind.FilterOpts) (*EVM2EVMOnRampTokensConsumedIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampTokensConsumedIterator{contract: _EVM2EVMOnRamp.contract, event: "TokensConsumed", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampTokensConsumed) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampTokensConsumed)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseTokensConsumed(log types.Log) (*EVM2EVMOnRampTokensConsumed, error) {
	event := new(EVM2EVMOnRampTokensConsumed)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var EVM2EVMOnRampZkBin string = ("0x000400000000000200330000000000020000006004100270000009b60340019700030000003103550002000000010355000009b60040019d0000000100200190000000200000c13d0000008002000039000000400020043f000000040030008c000000410000413d000000000201043b000000e00220027000000a050020009c0000004c0000a13d00000a060020009c0000005f0000213d00000a100020009c000003d20000a13d00000a110020009c0000044a0000213d00000a140020009c000006420000613d00000a150020009c000000410000c13d0000000001000416000000000001004b000000410000c13d000000000100041a0000063e0000013d000001a004000039000000400040043f0000000002000416000000000002004b000000410000c13d0000001f02300039000009b702200197000001a002200039000000400020043f0000001f0530018f000009b806300198000001a002600039000000320000613d000000000701034f000000007807043c0000000004840436000000000024004b0000002e0000c13d000000000005004b0000003f0000613d000000000161034f0000000304500210000000000502043300000000054501cf000000000545022f000000000101043b0000010004400089000000000141022f00000000014101cf000000000151019f0000000000120435000003600030008c000000430000813d0000000001000019000026d500010430000000400a00043d000009b900a0009c000001190000a13d00000a6201000041000000000010043f0000004101000039000000040010043f000009ec01000041000026d50001043000000a190020009c000002650000a13d00000a1a0020009c000003700000a13d00000a1b0020009c0000042f0000213d00000a1e0020009c000006190000613d00000a1f0020009c000000410000c13d0000000001000416000000000001004b000000410000c13d0000000e01000039000000000101041a000009bc01100197000000800010043f00000a5f01000041000026d40001042e00000a070020009c000003f10000a13d00000a080020009c0000047b0000213d00000a0b0020009c000006800000613d00000a0c0020009c000000410000c13d000000240030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000000402100370000000000202043b000009bb0020009c000000410000213d0000002304200039000000000034004b000000410000813d0000000404200039000000000441034f000000000504043b000009bb0050009c000000460000213d00000005045002100000003f04400039000009c004400197000009d90040009c000000460000213d0000008004400039000000400040043f000000800050043f000000a0045000c900000024022000390000000004420019000000000034004b000000410000213d000000000005004b0000095c0000c13d000000000100041a000009ba021001970000000001000411000000000021004b000000910000613d0000000202000039000000000202041a000009ba02200197000000000021004b00000a7c0000c13d000000800100043d000000000001004b000000e10000613d0000000002000019000000400700043d000009d90070009c000000460000213d0000000501200210000000a0011000390000000001010433001d00000002001d000000400210003900000000020204330000002003100039000000000303043300000060041000390000000004040433000000800510003900000000050504330000008006700039000000400060043f000000000005004b0000000005000039000000010500c0390000006006700039001b00000006001d0000000000560435000009bb044001970000004005700039001a00000005001d0000000000450435000009b603300197001c00000007001d0000000003370436000009bb02200197001900000003001d00000000002304350000000001010433000009ba01100197000000000010043f0000000b01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d0000001c020000290000000002020433000009b602200197000000000101043b000000000301041a000009cb03300197000000000223019f000000190300002900000000030304330000002003300210000009db03300197000000000232019f0000001a0300002900000000030304330000006003300210000009dc03300197000000000232019f0000001b030000290000000003030433000000000003004b000009c9030000410000000003006019000000000232019f000000000021041b0000001d020000290000000102200039000000800100043d000000000012004b000000950000413d000000400100043d00000020020000390000000002210436000000800300043d00000000003204350000004002100039000000000003004b000001070000613d000000a004000039000000000500001900000000460404340000000087060434000009ba0770019700000000077204360000000008080433000009b608800197000000000087043500000040076000390000000007070433000009bb077001970000004008200039000000000078043500000060076000390000000007070433000009bb077001970000006008200039000000000078043500000080066000390000000006060433000000000006004b0000000006000039000000010600c03900000080072000390000000000670435000000a0022000390000000105500039000000000035004b000000eb0000413d0000000002120049000009b60020009c000009b6020080410000006002200210000009b60010009c000009b6010080410000004001100210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000121019f000009cd011001c70000800d020000390000000103000039000009dd040000410000080e0000013d0000010001a00039000000400010043f000001a00100043d000009ba0010009c000000410000213d00000000011a0436001d00000001001d000001c00100043d000009bb0010009c000000410000213d0000001d020000290000000000120435000001e00100043d000009bb0010009c000000410000213d0000004002a00039001c00000002001d0000000000120435000002000100043d000009bb0010009c000000410000213d0000006002a00039001b00000002001d0000000000120435000002200100043d000009bc0010009c000000410000213d0000008002a00039001a00000002001d0000000000120435000002400100043d000009ba0010009c000000410000213d000000a002a00039001900000002001d0000000000120435000002600100043d000009ba0010009c000000410000213d000000c002a00039001800000002001d0000000000120435000002800100043d000009ba0010009c000000410000213d000000e002a00039001600000002001d0000000000120435000000400100043d001700000001001d000009bd0010009c000000460000213d0000001701000029000001a001100039000000400010043f000002a00100043d000009ba0010009c000000410000213d00000017020000290000000001120436001500000001001d000002c00100043d0000ffff0010008c000000410000213d00000015020000290000000000120435000002e00100043d000009b60010009c000000410000213d00000017020000290000004002200039001400000002001d0000000000120435000003000100043d0000ffff0010008c000000410000213d00000017020000290000006002200039001300000002001d0000000000120435000003200100043d000009b60010009c000000410000213d00000017020000290000008002200039001200000002001d0000000000120435000003400100043d0000ffff0010008c000000410000213d0000001702000029000000a002200039001100000002001d0000000000120435000003600100043d0000ffff0010008c000000410000213d0000001702000029000000c002200039001000000002001d0000000000120435000003800100043d000009ba0010009c000000410000213d0000001702000029000000e002200039000f00000002001d0000000000120435000003a00100043d000009b60010009c000000410000213d00000017020000290000010002200039000e00000002001d0000000000120435000003c00100043d000009b60010009c000000410000213d00000017020000290000012002200039000d00000002001d0000000000120435000003e00100043d0000ffff0010008c000000410000213d00000017020000290000014002200039000c00000002001d0000000000120435000004000100043d000009b60010009c000000410000213d00000017020000290000016002200039000b00000002001d0000000000120435000004200100043d000000000001004b0000000002000039000000010200c039000000000021004b000000410000c13d00000017020000290000018002200039000a00000002001d0000000000120435000000400100043d000009be0010009c000000460000213d0000006002100039000000400020043f000004400200043d000000000002004b0000000004000039000000010400c039000000000042004b000000410000c13d0000000002210436000004600400043d000009bf0040009c000000410000213d0000000000420435000004800500043d000009bf0050009c000000410000213d00000040041000390000000000540435000004a00600043d000009bb0060009c000000410000213d000001a005300039000001bf07600039000000000057004b000000410000813d000001a0076000390000000008070433000009bb0080009c000000460000213d00000005078002100000003f07700039000009c007700197000000400900043d0000000007790019000900000009001d000000000097004b00000000090000390000000109004039000009bb0070009c000000460000213d0000000100900190000000460000c13d000000400070043f00000009090000290000000007890436000800000007001d000001c006600039000000a0078000c90000000007670019003300200090003d000000000057004b000000410000213d00070000000a001d000000000008004b000011060000c13d000004c00700043d000009bb0070009c000000410000213d0000001f06700039000000000036004b0000000008000019000009c308004041000009c306600197000000000006004b0000000009000019000009c309002041000009c30060009c000000000908c019000000000009004b000000410000613d000001a0067000390000000009060433000009bb0090009c000000460000213d0000000506900210000000400a00043d00320000000a001d0000003f06600039000009c00660019700000000066a00190000000000a6004b00000000080000390000000108004039000009bb0060009c000000460000213d0000000100800190000000460000c13d000000400060043f00000000069a0436000001c007700039000000e0089000c900000000087800190031002000a0003d000000000058004b000000410000213d000000000009004b00000e550000c13d000004e00600043d000009bb0060009c000000410000213d0000001f07600039000000000037004b0000000003000019000009c303004041000009c307700197000000000007004b0000000008000019000009c308002041000009c30070009c000000000803c019000000000008004b000000410000613d000001a0036000390000000008030433000009bb0080009c000000460000213d0000000503800210000000400900043d003000000009001d0000003f03300039000009c0033001970000000003390019000000000093004b00000000070000390000000107004039000009bb0030009c000000460000213d0000000100700190000000460000c13d000000400030043f0000000003890436000001c00660003900000006078002100000000007670019002f00200090003d000000000057004b000000410000213d000000000008004b0000024f0000613d0000000008650049000009c10080009c000000410000213d000000400080008c000000410000413d000000400800043d000009c50080009c000000460000213d0000004009800039000000400090043f000000009a060434000009ba00a0009c000000410000213d000000000aa8043600000000090904330000ffff0090008c000000410000213d00000000009a043500000000038304360000004006600039000000000076004b000002390000413d000000400300043d000600000003001d0000000003000411000000000003004b00001b0a0000c13d0000000603000029000000440130003900000a04020000410000000000210435000000240130003900000018020000390000000000210435000009f5010000410000000000130435000000040130003900000020020000390000000000210435000009b60030009c000009b6030080410000004001300210000009fa011001c7000026d50001043000000a230020009c000002810000213d00000a270020009c000004f00000613d00000a280020009c0000048c0000613d00000a290020009c000000410000c13d0000000001000416000000000001004b000000410000c13d000000c001000039000000400010043f0000001301000039000000800010043f00000a7801000041000000a00010043f0000002001000039000000c00010043f0000008001000039000000e00200003926d321cc0000040f000000c00110008a000009b60010009c000009b601008041000000600110021000000a79011001c7000026d40001042e00000a240020009c000005620000613d00000a250020009c000004e10000613d00000a260020009c000000410000c13d000000440030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000000402100370000000000202043b000009bb0020009c000000410000213d0000002304200039000000000034004b000000410000813d0000000404200039000000000441034f000000000504043b000009bb0050009c000000460000213d00000005045002100000003f04400039000009c004400197000009d90040009c000000460000213d0000008004400039000000400040043f000000800050043f0000002402200039000000e0045000c90000000004240019000000000034004b000000410000213d000000000005004b000008a70000c13d0000002402100370000000000202043b000009bb0020009c000000410000213d0000002304200039000000000034004b0000000005000019000009c305004041000009c304400197000000000004004b0000000006000019000009c306002041000009c30040009c000000000605c019000000000006004b000000410000613d0000000404200039000000000441034f000000000404043b000009bb0040009c000000460000213d00000005054002100000003f06500039000009c006600197000000400700043d0000000006670019001500000007001d000000000076004b00000000070000390000000107004039000009bb0060009c000000460000213d0000000100700190000000460000c13d000000400060043f00000015060000290000000006460436001400000006001d00000024022000390000000005250019000000000035004b000000410000213d000000000004004b000002dc0000613d0000001403000029000000000421034f000000000404043b000009ba0040009c000000410000213d00000000034304360000002002200039000000000052004b000002d40000413d000000000100041a000009ba021001970000000001000411000000000021004b000002e60000613d0000000202000039000000000202041a000009ba02200197000000000021004b00000a7c0000c13d000000800100043d000000000001004b00000c450000c13d000000400100043d00000020020000390000000002210436000000800300043d00000000003204350000004002100039000000000003004b000003190000613d000000a004000039000000000500001900000000460404340000000087060434000009ba0770019700000000077204360000000008080433000009b608800197000000000087043500000040076000390000000007070433000009b60770019700000040082000390000000000780435000000600760003900000000070704330000ffff0770018f0000006008200039000000000078043500000080076000390000000007070433000009b60770019700000080082000390000000000780435000000a0076000390000000007070433000009b607700197000000a0082000390000000000780435000000c0066000390000000006060433000000000006004b0000000006000039000000010600c039000000c0072000390000000000670435000000e0022000390000000105500039000000000035004b000002f30000413d0000000002120049000009b60020009c000009b6020080410000006002200210000009b60010009c000009b6010080410000004001100210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000121019f000009cd011001c70000800d020000390000000103000039000009e70400004126d326c90000040f0000000100200190000000410000613d00000015010000290000000001010433000000000001004b0000083a0000613d0000000002000019001d00000002001d000000050120021000000014011000290000000001010433000009ba01100197000000000010043f0000000c01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000001041b0000001d02000029000000010220003900000015010000290000000001010433000000000012004b000003320000413d000000000001004b0000083a0000613d000000400100043d000000200200003900000000022104360000001503000029000000000303043300000000003204350000004002100039000000000003004b0000035e0000613d000000000400001900000014060000290000000065060434000009ba0550019700000000025204360000000104400039000000000034004b000003580000413d0000000002120049000009b60020009c000009b6020080410000006002200210000009b60010009c000009b6010080410000004001100210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000121019f000009cd011001c70000800d020000390000000103000039000009e8040000410000080e0000013d00000a200020009c0000076b0000613d00000a210020009c000006da0000613d00000a220020009c000000410000c13d0000000001000416000000000001004b000000410000c13d000000800000043f000000a00000043f000000c00000043f000000e00000043f000001000000043f0000000301000039000000000101041a000009bf02100197000001200020043f0000008002100270000009b602200197001d00000002001d000001400020043f00000a3e001001980000000001000039000000010100c039000001600010043f0000000401000039000000000101041a000009bf02100197001c00000002001d000001800020043f0000008001100270000001a00010043f0000026001000039000000400010043f000001c00000043f000001e00000043f000002000000043f000002200000043f000002400000043f000009c70100004100000000001004430000000001000414000009b60010009c000009b601008041000000c001100210000009c8011001c70000800b0200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b0000001d06000029000000000361004b000004ea0000413d000001a00200043d000009bf0520019700000000023500a9000001200400043d000000000061004b000003b00000613d00000000033200d9000000000053004b000004ea0000c13d000009bf03400197000000000032001a000004ea0000413d00000000023200190000001c0020006b00000000030200190000001c03004029000001200030043f000009b601100197000001400010043f000000400100043d0000000002310436000001400300043d000009b6033001970000000000320435000001600200043d000000000002004b0000000002000039000000010200c03900000040031000390000000000230435000001800200043d000009bf0220019700000060031000390000000000230435000001a00200043d000009bf0220019700000080031000390000000000230435000009b60010009c000009b601008041000000400110021000000a6a011001c7000026d40001042e00000a160020009c0000077a0000613d00000a170020009c000007300000613d00000a180020009c000000410000c13d0000000001000416000000000001004b000000410000c13d0000000101000039000000000201041a000009ba032001970000000006000411000000000036004b000007f80000c13d000000000300041a000009c604300197000000000464019f000000000040041b000009c602200197000000000021041b0000000001000414000009ba05300197000009b60010009c000009b601008041000000c001100210000009cd011001c70000800d02000039000000030300003900000a6104000041000008370000013d00000a0d0020009c000007b90000613d00000a0e0020009c000007600000613d00000a0f0020009c000000410000c13d000000840030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000000402100370000000000202043b001d00000002001d000009bb0020009c000000410000213d0000002402100370000000000202043b000009bb0020009c000000410000213d0000000002230049000009c10020009c000000410000213d000000a40020008c000000410000413d0000006401100370000000000101043b001e00000001001d001c00000001001d000009ba0010009c000000410000213d00000a3001000041000000800010043f0000001d01000029000000800110021000000a3101100197000000840010043f00000a3201000041000000000010044300000000010004120000000400100443000000e00100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000201043b0000000001000414000009ba02200197000000040020008c000008e80000c13d0000000103000031000000200030008c000000200400003900000000040340190000090d0000013d00000a1c0020009c000006390000613d00000a1d0020009c000000410000c13d000000240030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000000401100370000000000101043b000009ba0010009c000000410000213d0000000202000039000000000302041a000000000400041a000009ba054001970000000004000411000000000054004b000008020000613d000009ba05300197000000000054004b000008020000613d00000a5901000041000000800010043f00000a2b01000041000026d50001043000000a120020009c000006860000613d00000a130020009c000000410000c13d0000000001000416000000000001004b000000410000c13d0000000801000039000000000101041a001b00000001001d000009bb0010009c000000460000213d0000001b0100002900000005011002100000003f02100039000009c002200197000009d90020009c000000460000213d0000008004200039000000400040043f0000001b03000029000000800030043f000000000003004b000008190000c13d001b00000004001d0000000e01000039000000000101041a001d00000001001d00000040010000390000000001140436001c00000001001d0000004002400039000000800100003926d3221e0000040f0000001d020000290000006002200270000009b6022001970000001c0300002900000000002304350000001b020000290000000001210049000009b60010009c000009b601008041000009b60020009c000009b60200804100000060011002100000004002200210000000000121019f000026d40001042e00000a090020009c000006c00000613d00000a0a0020009c000000410000c13d000000240030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000000401100370000000000101043b000009bb0010009c000000410000213d00000a2a01000041000000800010043f00000a2b01000041000026d500010430000000240030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000000401100370000000000101043b000009ba0010009c000000410000213d0000016002000039000000400020043f000000800000043f000000a00000043f000000c00000043f000000e00000043f000001000000043f000001200000043f000001400000043f000000000010043f0000000c01000039000000200010043f0000004002000039000000000100001926d326960000040f001d00000001001d000001600100003926d321b60000040f0000001d01000029000000000101041a000009b602100197000001600020043f0000002003100270000009b603300197000001800030043f00000040031002700000ffff0330018f000001a00030043f0000005003100270000009b603300197000001c00030043f0000007003100270000009b603300197000001e00030043f00000a39001001980000000003000039000000010300c039000002000030043f00000a50001001980000000001000039000000010100c039000002200010043f000000400100043d0000000002210436000001800300043d000009b6033001970000000000320435000001a00200043d0000ffff0220018f00000040031000390000000000230435000001c00200043d000009b60220019700000060031000390000000000230435000001e00200043d000009b60220019700000080031000390000000000230435000002000200043d000000000002004b0000000002000039000000010200c039000000a0031000390000000000230435000002200200043d000000000002004b0000000002000039000000010200c039000000c0031000390000000000230435000009b60010009c000009b601008041000000400110021000000a7a011001c7000026d40001042e0000000001000416000000000001004b000000410000c13d0000000e01000039000000000101041a0000008001100270000009bb01100197000009bb0010009c000007f40000c13d00000a6201000041000000000010043f0000001101000039000000040010043f000009ec01000041000026d5000104300000000001000416000000000001004b000000410000c13d000000800000043f000000a00000043f000000c00000043f000000e00000043f000001000000043f000001200000043f000001400000043f000001600000043f0000028001000039000000400010043f0000000001000412002e00000001001d002d00600000003d0000800501000039000000440300003900000000040004150000002e0440008a000000050440021000000a320200004126d326ab0000040f000009ba01100197000001800010043f0000000001000412002c00000001001d002b00800000003d00000000040004150000002c0440008a0000000504400210000080050100003900000a3202000041000000440300003926d326ab0000040f000009bb01100197000001a00010043f0000000001000412002a00000001001d002900a00000003d00000000040004150000002a0440008a0000000504400210000080050100003900000a3202000041000000440300003926d326ab0000040f000009bb01100197000001c00010043f0000000001000412002800000001001d002700200000003d0000000004000415000000280440008a0000000504400210000080050100003900000a3202000041000000440300003926d326ab0000040f000009bb01100197000001e00010043f0000000001000412002600000001001d002500400000003d0000000004000415000000260440008a0000000504400210000080050100003900000a3202000041000000440300003926d326ab0000040f000009bc01100197000002000010043f0000000001000412002400000001001d002300c00000003d0000000004000415000000240440008a0000000504400210000080050100003900000a3202000041000000440300003926d326ab0000040f000009ba01100197000002200010043f0000000001000412002200000001001d002100e00000003d0000000004000415000000220440008a0000000504400210000080050100003900000a3202000041000000440300003926d326ab0000040f000009ba01100197000002400010043f0000000001000412002000000001001d001f01000000003d0000000004000415000000200440008a0000000504400210000080050100003900000a3202000041000000440300003926d326ab0000040f000009ba01100197000002600010043f0000018001000039000002800200003926d321910000040f00000a7b01000041000026d40001042e000000440030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000000402100370000000000202043b001d00000002001d000009bb0020009c000000410000213d0000002401100370000000000101043b000009bb0010009c000000410000213d0000000001130049000009c10010009c000000410000213d000000a40010008c000000410000413d00000a3201000041000000000010044300000000010004120000000400100443000000a00100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b0000001d03000029000000000131013f000009bb001001980000084b0000c13d00000002010003670000002402100370000000000202043b0000008403200039000000000331034f000000000403043b00000000030000310000000005230049000000230550008a000009c306500197000009c307400197000000000867013f000000000067004b0000000006000019000009c306004041000000000054004b0000000005000019000009c305008041000009c30080009c000000000605c019000000000006004b000000410000c13d00000000024200190000000404200039000000000141034f000000000201043b000009bb0020009c000000410000213d00000000032300490000002001400039000009c304300197000009c305100197000000000645013f000000000045004b0000000004000019000009c304004041000000000031004b0000000003000019000009c303002041000009c30060009c000000000403c019000000000004004b000000410000c13d26d325530000040f00000002020003670000002403200370000000000403043b0000002407400039000000000372034f000000000303043b0000000405400039000000000400003100000000065400490000001f0660008a000009c308600197000009c309300197000000000a89013f000000000089004b0000000009000019000009c309004041000000000063004b000000000b000019000009c30b008041000009c300a0009c00000000090bc019001c00000001001d000000000009004b000000410000c13d0000000001530019000000000312034f000000000303043b000009bb0030009c000000410000213d00000000093400490000002001100039000000000091004b000000000a000019000009c30a002041000009c309900197000009c301100197000000000b91013f000000000091004b0000000001000019000009c301004041000009c300b0009c00000000010ac019000000000001004b000000410000c13d0000002001700039000000000712034f000000000707043b000009c309700197000000000a89013f000000000089004b0000000008000019000009c308004041000000000067004b0000000006000019000009c306008041000009c300a0009c000000000806c019000000000008004b000000410000c13d0000001c0600002900000000060604330000000007570019000000000572034f000000000505043b000009bb0050009c000000410000213d000000060850021000000000048400490000002007700039000000000047004b0000000008000019000009c308002041000009c304400197000009c307700197000000000947013f000000000047004b0000000004000019000009c304004041000009c30090009c000000000408c019000000000004004b000000410000c13d0000000604000039000000000704041a000000c004700270000009b604400197000000000034004b00000cce0000813d000000400100043d0000002402100039000000000032043500000a5702000041000000000021043500000004021000390000000000420435000009b60010009c000009b6010080410000004001100210000009e6011001c7000026d500010430000000440030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000000402100370000000000202043b001d00000002001d000009ba0020009c000000410000213d0000002401100370000000000101043b001c00000001001d000009ba0010009c000000410000213d000000000100041a000009ba021001970000000001000411000000000021004b000006320000613d0000000202000039000000000202041a000009ba02200197000000000021004b000008470000c13d0000001c01000029000009ba001001980000083c0000c13d00000a6901000041000000800010043f00000a2b01000041000026d5000104300000000001000416000000000001004b000000410000c13d0000000201000039000000000101041a000009ba01100197000000800010043f00000a5f01000041000026d40001042e000000240030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000000401100370000000000101043b001d00000001001d000009ba0010009c000000410000213d0000001d01000029000000000010043f0000000d01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000101041a001c00000001001d00000a3201000041000000000010044300000000010004120000000400100443000000c00100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000400b00043d0000001c02000029000009bb03200197000000000101043b000009ba02100198000008550000613d000000000003004b000008550000c13d00000a470100004100000000001b04350000000401b000390000001d0300002900000000003104350000000001000414000000040020008c0000091f0000c13d0000000103000031000000200030008c000000200400003900000000040340190000094b0000013d0000000001000416000000000001004b000000410000c13d26d323500000040f0000000001000019000026d40001042e000000240030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000000401100370000000000101043b000009ba0010009c000000410000213d0000010002000039000000400020043f000000800000043f000000a00000043f000000c00000043f000000e00000043f000000000010043f0000000b01000039000000200010043f0000004002000039000000000100001926d326960000040f001d00000001001d000001000100003926d321c10000040f0000001d01000029000000000101041a000009b602100197000001000020043f0000002003100270000009bb03300197000001200030043f0000006003100270000009bb03300197000001400030043f00000a3e001001980000000001000039000000010100c039000001600010043f000000400100043d0000000002210436000001200300043d000009bb033001970000000000320435000001400200043d000009bb0220019700000040031000390000000000230435000001600200043d000000000002004b0000000002000039000000010200c03900000060031000390000000000230435000009b60010009c000009b601008041000000400110021000000a5e011001c7000026d40001042e000000240030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000000401100370000000000601043b000009ba0060009c000000410000213d000000000100041a000009ba011001970000000005000411000000000015004b0000080f0000c13d000000000056004b0000082a0000c13d000009f501000041000000800010043f0000002001000039000000840010043f0000001701000039000000a40010043f00000a2f01000041000000c40010043f00000a2d01000041000026d500010430000001a40030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000022002000039000000400020043f0000000402100370000000000202043b000009ba0020009c000000410000213d000000800020043f0000002403100370000000000303043b0000ffff0030008c000000410000213d000000a00030043f0000004404100370000000000404043b000009b60040009c000000410000213d000000c00040043f0000006405100370000000000605043b0000ffff0060008c000000410000213d000000e00060043f0000008405100370000000000705043b000009b60070009c000000410000213d000001000070043f000000a405100370000000000505043b0000ffff0050008c000000410000213d000001200050043f000000c408100370000000000808043b0000ffff0080008c000000410000213d000001400080043f000000e409100370000000000909043b000009ba0090009c000000410000213d000001600090043f000001040a100370000000000b0a043b000009b600b0009c000000410000213d0000018000b0043f000001240a100370000000000c0a043b000009b600c0009c000000410000213d000001a000c0043f000001440a100370000000000a0a043b0000ffff00a0008c000000410000213d000001c000a0043f000001640d100370000000000d0d043b000009b600d0009c000000410000213d000001e000d0043f0000018401100370000000000101043b000000000001004b000000000e000039000000010e00c0390000000000e1004b000000410000c13d000002000010043f000000000e00041a000009ba0ee00197000000000f0004110000000000ef004b00000cd40000c13d000000000009004b00000cde0000c13d00000a0301000041000002200010043f00000a6c01000041000026d500010430000000240030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000000402100370000000000402043b000009bb0040009c000000410000213d0000002302400039000000000032004b000000410000813d0000000402400039000000000221034f000000000202043b000009bb0020009c000000410000213d0000000008020019000000060220021000000024052001bf0000000002450019000000000032004b000000410000213d000000000600041a000009ba076001970000000006000411000000000076004b000007510000613d0000000207000039000000000707041a000009ba07700197000000000076004b000008470000c13d00000005068002100000003f06600039000009c006600197000009d90060009c000000460000213d0000008006600039000000400060043f000000800080043f000000240050008c000009930000c13d000000400080008c000009b10000a13d000000400100043d00000a6302000041000008410000013d0000000001000416000000000001004b000000410000c13d26d322b60000040f000000400200043d0000000000120435000009b60020009c000009b602008041000000400120021000000a55011001c7000026d40001042e000000440030008c000000410000413d0000000002000416000000000002004b000000410000c13d0000000402100370000000000202043b000009bb0020009c000000410000213d0000002401100370000000000101043b000009ba0010009c000000410000213d26d322310000040f000007640000013d0000000001000416000000000001004b000000410000c13d000000800000043f000000a00000043f000000c00000043f000000e00000043f000001000000043f000001200000043f000001400000043f000001600000043f000001800000043f000001a00000043f000001c00000043f000001e00000043f000002000000043f000003c002000039000000400020043f0000000501000039000000000101041a000009ba03100197000002200030043f000000a0031002700000ffff0330018f000002400030043f000000b003100270000009b603300197000002600030043f000000d0031002700000ffff0330018f000002800030043f000000e001100270000002a00010043f0000000601000039000000000101041a0000ffff0310018f000002c00030043f00000010031002700000ffff0330018f000002e00030043f0000002003100270000009ba03300197000003000030043f000000c003100270000009b603300197000003200030043f000000e001100270000003400010043f0000000701000039000000000101041a0000ffff0310018f000003600030043f0000001003100270000009b603300197000003800030043f00000a38001001980000000001000039000000010100c039000003a00010043f000002200100003926d321de0000040f00000a6401000041000026d40001042e000000640030008c000000410000413d0000000002000416000000000002004b000000410000c13d000000e002000039000000400020043f0000000402100370000000000202043b000000000002004b0000000003000039000000010300c039000000000032004b000000410000c13d000000800020043f0000002402100370000000000202043b000009bf0020009c000000410000213d000000a00020043f0000004401100370000000000101043b000009bf0010009c000000410000213d000000c00010043f000000000100041a000009ba021001970000000001000411000000000021004b000007dc0000613d0000000202000039000000000202041a000009ba02200197000000000021004b0000098f0000c13d0000000301000039000000000101041a001d00000001001d000009c70100004100000000001004430000000001000414000009b60010009c000009b601008041000000c001100210000009c8011001c70000800b0200003926d326ce0000040f00000001002001900000212f0000613d0000001d060000290000008002600270000009b602200197000000000101043b000000000421004b000004ea0000413d00000a3d0000c13d0000000304000039000000000104041a00000a520000013d0000000101100039000000800010043f00000a5f01000041000026d40001042e000009f501000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f00000a6001000041000000c40010043f00000a2d01000041000026d500010430000009c603300197000000000313019f000000000032041b000000800010043f0000000001000414000009b60010009c000009b601008041000000c00110021000000a65011001c70000800d02000039000000010300003900000a6604000041000008370000013d000009f501000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f00000a2c01000041000000c40010043f00000a2d01000041000026d500010430000009ef0020009c000000460000213d00000000020000190000004003400039000000400030043f000000200340003900000000000304350000000000040435000000a00320003900000000004304350000002002200039000000000012004b000008760000813d000000400400043d000009c50040009c0000081c0000a13d000000460000013d0000000101000039000000000201041a000009c602200197000000000262019f000000000021041b0000000001000414000009b60010009c000009b601008041000000c001100210000009cd011001c70000800d02000039000000030300003900000a2e0400004126d326c90000040f0000000100200190000000410000613d0000000001000019000026d40001042e26d322b60000040f000009c30010009c0000085c0000413d000000400100043d00000a68020000410000000000210435000009b60010009c000009b6010080410000004001100210000009ea011001c7000026d50001043000000a6701000041000000800010043f00000a2b01000041000026d500010430000000400100043d00000a3702000041000000000021043500000004021000390000000000320435000009b60010009c000009b6010080410000004001100210000009ec011001c7000026d50001043000000000010b00190000000000310435000009b60010009c000009b601008041000000400110021000000a55011001c7000026d40001042e001b00000001001d00000a3201000041000000000010044300000000010004120000000400100443000000600100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b000009ba021001970000001d01000029000000000021004b00000a2e0000c13d0000001c020000290000001b0300002926d325bb0000040f0000000001000019000026d40001042e00000000030000190000000801000039000000000101041a000000000031004b000021300000a13d001d00000003001d000009ed0130009a000000000101041a001c00000001001d000000000010043f0000000a01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000400200043d000009c50020009c000000460000213d000000000101043b000000000101041a0000004003200039000000400030043f0000ffff0110018f000000200320003900000000001304350000001c01000029000009ba011001970000000000120435000000800100043d0000001d03000029000000000031004b000021300000a13d0000000501300210000000a0011000390000000000210435000000800100043d000000000031004b000021300000a13d00000001033000390000001b0030006c000008770000413d000000400400043d000004620000013d000000a0050000390000000006230049000009c10060009c000000410000213d000000e00060008c000000410000413d000000400600043d000009c40060009c000000460000213d000000e007600039000000400070043f000000000721034f000000000707043b000009ba0070009c000000410000213d00000000087604360000002007200039000000000971034f000000000909043b000009b60090009c000000410000213d00000000009804350000002007700039000000000871034f000000000808043b000009b60080009c000000410000213d000000400960003900000000008904350000002007700039000000000871034f000000000808043b0000ffff0080008c000000410000213d000000600960003900000000008904350000002007700039000000000871034f000000000808043b000009b60080009c000000410000213d000000800960003900000000008904350000002007700039000000000871034f000000000808043b000009b60080009c000000410000213d000000a00960003900000000008904350000002007700039000000000771034f000000000707043b000000000007004b0000000008000039000000010800c039000000000087004b000000410000c13d000000c00860003900000000007804350000000005650436000000e002200039000000000042004b000008a80000413d000002a70000013d000009b60010009c000009b601008041000000c00110021000000a34011001c726d326ce0000040f0000006003100270000009b603300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000080057001bf000000800a000039000008fc0000613d000000000801034f000000008908043c000000000a9a043600000000005a004b000008f80000c13d000000000006004b000009090000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000000a040000613d0000001f01400039000000600110018f00000080011001bf000000400010043f000000200030008c000000410000413d000000800200043d000000000002004b0000000003000039000000010300c039000000000032004b000000410000c13d000000000002004b00000aef0000c13d0000001c0000006b00000b320000c13d00000a580200004100000af00000013d000009b600b0009c000009b60300004100000000030b40190000004003300210000009b60010009c000009b601008041000000c001100210000000000131019f000009ec011001c7001d0000000b001d26d326ce0000040f0000001d0b0000290000006003100270000009b603300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b00190000093a0000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b000009360000c13d000000000006004b000009470000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000000a100000613d0000001f01400039000000600210018f0000000001b20019000000000021004b00000000020000390000000102004039000009bb0010009c000000460000213d0000000100200190000000460000c13d000000400010043f000000200030008c000000410000413d00000000030b0433000009bb0030009c000008560000a13d000000410000013d000000a0050000390000000006230049000009c10060009c000000410000213d000000a00060008c000000410000413d000000400600043d000009c20060009c000000460000213d000000a007600039000000400070043f000000000721034f000000000707043b000009ba0070009c000000410000213d00000000087604360000002007200039000000000971034f000000000909043b000009b60090009c000000410000213d00000000009804350000002007700039000000000871034f000000000808043b000009bb0080009c000000410000213d000000400960003900000000008904350000002007700039000000000871034f000000000808043b000009bb0080009c000000410000213d000000600960003900000000008904350000002007700039000000000771034f000000000707043b000000000007004b0000000008000039000000010800c039000000000087004b000000410000c13d000000800860003900000000007804350000000005650436000000a002200039000000000042004b0000095d0000413d000000870000013d00000a5901000041000000e00010043f00000a5a01000041000026d500010430000000a00500003900000024044000390000000006430049000009c10060009c000000410000213d000000400060008c000000410000413d000000400600043d000009c50060009c000000460000213d0000004007600039000000400070043f000000000741034f000000000707043b000009ba0070009c000000410000213d00000000077604360000002008400039000000000881034f000000000808043b0000ffff0080008c000000410000213d000000000087043500000000056504360000004004400039000000000024004b000009950000413d000000800800043d000000400080008c0000075d0000213d001900000008001d0000000e01000039000000000101041a000009bc021001970000006001100270000009b601100197000000010110008a000000000021004b000009bb0000813d26d323500000040f0000000803000039000000000103041a000000000001004b00000a7f0000c13d00000a3201000041000000000010044300000000010004120000000400100443000000600100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b000000190000006b000000000400001900000be70000c13d0000000e03000039000000000103041a000009fe011001970000006002400210000009ff02200197000000000112019f000000000013041b000000400100043d00000020021000390000004003000039000000000032043500000000004104350000004002100039000000800300043d00000000003204350000006002100039000000000003004b000009f20000613d00000080040000390000000005000019000000200440003900000000060404330000000076060434000009ba06600197000000000662043600000000070704330000ffff0770018f000000000076043500000040022000390000000105500039000000000035004b000009e60000413d0000000002120049000009b60020009c000009b6020080410000006002200210000009b60010009c000009b6010080410000004001100210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000121019f000009cd011001c70000800d02000039000000010300003900000a00040000410000080e0000013d0000001f0530018f000009b806300198000000400200043d000000000462001900000a1b0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000a0b0000c13d00000a1b0000013d0000001f0530018f000009b806300198000000400200043d000000000462001900000a1b0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000a170000c13d000000000005004b00000a280000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000009b60020009c000009b6020080410000004002200210000000000112019f000026d500010430000000400300043d000009eb020000410000000000230435001b00000003001d0000000403300039000000000200041000000000002304350000000003000414000000040010008c00000af40000c13d0000000103000031000000200030008c0000002004000039000000000403401900000b200000013d0000000402000039000000000202041a000000800520027000000000034500a900000000044300d9000000000054004b000004ea0000c13d000009bf04600197000000000043001a000004ea0000413d000009bf022001970000000003430019000000000032004b00000000020380190000008001100210000009ca01100197000000000112019f0000000304000039000000000204041a000009c602200197000000000121019f000000a00200043d000009bf02200197000009bf03100197000000000032004b000000000302401900000a5b01100197000000000113019f000000800300043d000000000003004b0000000003000019000009c90300c041000000000131019f000000000014041b000000c00100043d0000008001100210000000000121019f0000000402000039000000000012041b0000000001000039000000010100c039000000400200043d0000000001120436000000a00300043d000009bf033001970000000000310435000000c00100043d000009bf0110019700000040032000390000000000130435000009b60020009c000009b60200804100000040012002100000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f00000a5c011001c70000800d02000039000000010300003900000a5d040000410000080e0000013d000000400100043d00000a6702000041000008410000013d000080100200003900000a860000013d0000001c01000029000000000001004b00000008030000390000801002000039000009bf0000613d001c000100100092000000000303041a0000001c0030006c000021300000a13d000009fc0110009a000000000101041a001d00000001001d000000000010043f0000000a01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c726d326ce0000040f0000000100200190000000410000613d0000001d01000029000009ba01100197001d00000001001d000000000010043f0000000a01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000001041b0000001d01000029000000000010043f0000000901000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000501041a000000000005004b00000a810000613d0000000803000039000000000403041a000000000004004b0000801002000039000004ea0000613d000000010150008a000000000045004b00000add0000613d000000000014004b000021300000a13d000009fc0150009a000009fc0340009a000000000303041a000000000031041b000000000030043f0000000901000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7001b00000005001d26d326ce0000040f0000000100200190000000410000613d000000000101043b0000001b02000029000000000021041b0000000803000039000000000503041a000000000005004b000021360000613d000000010150008a0000801002000039000009fc0450009a000000000004041b000000000013041b0000001d01000029000000000010043f0000000901000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c726d326ce0000040f0000000100200190000000410000613d000000000101043b000000000001041b00000a810000013d00000a350200004100000000002104350000004001100210000009ea011001c7000026d5000104300000001b01000029000009b60010009c000009b60200004100000000020140190000004002200210000009b60030009c000009b603008041000000c001300210000000000121019f000009ec011001c70000001d0200002926d326ce0000040f0000006003100270000009b603300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000001b0570002900000b0f0000613d000000000801034f0000001b09000029000000008a08043c0000000009a90436000000000059004b00000b0b0000c13d000000000006004b00000b1c0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000000bdb0000613d0000001f01400039000000600210018f0000001b01200029000000000021004b00000000020000390000000102004039000009bb0010009c000000460000213d0000000100200190000000460000c13d000000400010043f000000200030008c000000410000413d0000001b010000290000000001010433001b00000001001d0000001d01000029000008710000013d0000000502000039000000000202041a000009ba022001970000000003000411000000000023004b00000c430000c13d00000a3201000041000000000010044300000000010004120000000400100443000000a00100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b0000001d0110014f000009bb0010019800000cb50000c13d00000002010003670000002402100370000000000202043b0000008403200039000000000331034f000000000403043b00000000030000310000000005230049000000230550008a000009c306500197000009c307400197000000000867013f000000000067004b0000000006000019000009c306004041000000000054004b0000000005000019000009c305008041000009c30080009c000000000605c019000000000006004b000000410000c13d00000000024200190000000404200039000000000141034f000000000201043b000009bb0020009c000000410000213d00000000032300490000002001400039000009c304300197000009c305100197000000000645013f000000000045004b0000000004000019000009c304004041000000000031004b0000000003000019000009c303002041000009c30060009c000000000403c019000000000004004b000000410000c13d26d325530000040f00000002020003670000002403200370000000000703043b0000004405700039000000000352034f000000000803043b00000000030000310000000004730049000000230440008a000009c306400197000009c309800197000000000a69013f000000000069004b0000000009000019000009c309004041000000000048004b000000000b000019000009c30b008041000009c300a0009c00000000090bc019001500000001001d000000000009004b000000410000c13d00000004017000390000000007180019000000000872034f000000000808043b001600000008001d000009bb0080009c000000410000213d0000001608000029000000060880021000000000088300490000002007700039000000000087004b0000000009000019000009c309002041000009c308800197000009c307700197000000000a87013f000000000087004b0000000007000019000009c307004041000009c300a0009c000000000709c019000000000007004b000000410000c13d000000200550008a000000000552034f000000000505043b000009c307500197000000000867013f000000000067004b0000000006000019000009c306004041000000000045004b0000000004000019000009c304008041000009c30080009c000000000604c019000000000006004b000000410000c13d0000000004150019000000000142034f000000000101043b000009bb0010009c000000410000213d00000000021300490000002003400039000000000023004b0000000004000019000009c304002041000009c302200197000009c303300197000000000523013f000000000023004b0000000002000019000009c302004041000009c30050009c000000000204c019000000000002004b000000410000c13d0000000602000039000000000302041a000000c002300270000009b602200197000000000012004b00000e990000813d000000400300043d0000002404300039000000000014043500000a5701000041000000000013043500000004013000390000000000210435000009b60030009c000009b6030080410000004001300210000009e6011001c7000026d5000104300000001f0530018f000009b806300198000000400200043d000000000462001900000a1b0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000be20000c13d00000a1b0000013d001809ba0010019b000080100300003900000000020000190000000004000019000000800100043d000000000021004b000021300000a13d001b00000004001d001a00000002001d0000000501200210000000a00110003900000000010104330000000012010434000009ba04200197000000180040006c00000cbb0000613d000000000004004b00000cbb0000613d0000000001010433001c00000001001d000000000040043f0000000a01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c70000000002030019001d00000004001d26d326ce0000040f0000000100200190000000410000613d0000001c020000290000ffff0220018f000000000101043b001c00000002001d000000000021041b0000001d01000029000000000010043f0000000901000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d0000001d04000029000000000101043b000000000101041a000000000001004b00000c380000c13d0000000803000039000000000103041a000009bb0010009c000000460000213d0000000102100039000000000023041b000009ed0110009a000000000041041b000000000103041a001700000001001d000000000040043f0000000901000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b0000001702000029000000000021041b0000001b01000029000009b6011001970000001c04100029000009b60040009c000004ea0000213d0000001a020000290000000102200039000000190020006c000080100300003900000beb0000413d000009d20000013d00000a360200004100000af00000013d0000000002000019000000400900043d001d00000002001d0000000501200210000000a0011000390000000001010433000000a0021000390000000002020433000009b6022001970000001f0020008c00000cc10000a13d000009c40090009c000000460000213d00000040031000390000000003030433000000200410003900000000040404330000006005100039000000000505043300000080061000390000000006060433000000c0071000390000000007070433000000e008900039000000400080043f000000c00a9000390000000108000039001b0000000a001d00000000008a0435000000000007004b0000000007000039000000010700c039000000a008900039001a00000008001d00000000007804350000008007900039001900000007001d0000000000270435000009b6026001970000006006900039001800000006001d00000000002604350000ffff0250018f0000004005900039001700000005001d0000000000250435000009b602400197001c00000009001d0000000004290436000009b602300197001600000004001d00000000002404350000000001010433000009ba01100197000000000010043f0000000c01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d0000001c020000290000000002020433000009b602200197000000000101043b000000000301041a000009c603300197000000000223019f000000160300002900000000030304330000002003300210000009df03300197000000000232019f000000170300002900000000030304330000004003300210000009e003300197000000000232019f000000180300002900000000030304330000005003300210000009e103300197000000000232019f000000190300002900000000030304330000007003300210000009e203300197000000000232019f0000001a030000290000000003030433000000000003004b000009e3030000410000000003006019000000000232019f0000001b030000290000000003030433000000000003004b000009e4030000410000000003006019000000000232019f000000000021041b0000001d020000290000000102200039000000800100043d000000000012004b00000c460000413d000002e90000013d000000400100043d00000a3702000041000000000021043500000004021000390000001d030000290000084f0000013d000000400100043d000009fd02000041000000000021043500000004021000390000000000420435000008500000013d000000000101043300000024039000390000000000230435000009e5020000410000000000290435000009ba0110019700000004029000390000000000120435000009b60090009c000009b6090080410000004001900210000009e6011001c7000026d500010430000000e003700270000000000063004b00000df80000813d000000400100043d00000a7702000041000008410000013d000009f501000041000002200010043f0000002001000039000002240010043f0000001601000039000002440010043f00000a2c01000041000002640010043f00000a6b01000041000026d500010430000000a003300210000009ce03300197000000b004400210000009cf04400197000000000334019f000000d004600210000009d004400197000000000343019f000000e004700210000000000343019f000000000223019f0000000503000039000000000023041b0000001002800210000009d2022001970000002003900210000009d103300197000000000223019f000000c003b00210000009d303300197000000000232019f000000e003c00210000000000232019f000000000252019f0000000603000039000000000023041b000000000001004b0000000001000019000009d60100c0410000001002d00210000009d502200197000000000112019f0000000702000039000000000302041a000009d403300197000000000131019f0000000001a1019f000000000012041b0000032001000039000000400010043f00000a3201000041000000000010044300000000010004120000000400100443000000600100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b000009ba01100197000002200010043f00000a3201000041000000000010044300000000010004120000000400100443000000800100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b000009bb01100197000002400010043f00000a3201000041000000000010044300000000010004120000000400100443000000a00100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b000009bb01100197000002600010043f00000a3201000041000000000010044300000000010004120000000400100443000000200100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b000009bb01100197000002800010043f00000a3201000041000000000010044300000000010004120000000400100443000000400100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b000009bc01100197000002a00010043f00000a3201000041000000000010044300000000010004120000000400100443000000c00100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b000009ba01100197000002c00010043f00000a3201000041000000000010044300000000010004120000000400100443000000e00100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b000009ba01100197000002e00010043f00000a3201000041000000000010044300000000010004120000000400100443000001000100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b000009ba01100197000003000010043f000002200100043d000009ba02100197000000400100043d0000000002210436000002400300043d000009bb033001970000000000320435000002600200043d000009bb0220019700000040031000390000000000230435000002800200043d000009bb0220019700000060031000390000000000230435000002a00200043d000009bc0220019700000080031000390000000000230435000002c00200043d000009ba02200197000000a0031000390000000000230435000002e00200043d000009ba02200197000000c0031000390000000000230435000003000200043d000009ba02200197000000e0031000390000000000230435000000800200043d000009ba0220019700000100031000390000000000230435000000a00200043d0000ffff0220018f00000120031000390000000000230435000000c00200043d000009b60220019700000140031000390000000000230435000000e00200043d0000ffff0220018f00000160031000390000000000230435000001000200043d000009b60220019700000180031000390000000000230435000001200200043d0000ffff0220018f000001a0031000390000000000230435000001400200043d0000ffff0220018f000001c0031000390000000000230435000001600200043d000009ba02200197000001e0031000390000000000230435000001800200043d000009b60220019700000200031000390000000000230435000001a00200043d000009b60220019700000220031000390000000000230435000001c00200043d0000ffff0220018f00000240031000390000000000230435000001e00200043d000009b60220019700000260031000390000000000230435000002000200043d000000000002004b0000000002000039000000010200c03900000280031000390000000000230435000009b60010009c000009b60100804100000040011002100000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f000009d7011001c70000800d020000390000000103000039000009d8040000410000080e0000013d0000000503000039000000000303041a000000a0033002700000ffff0330018f000000000053004b00000e010000813d000000400100043d00000a7602000041000008410000013d0000001c0300002900000020033000390000000003030433000000000003004b00000e0a0000c13d0000000703000039000000000303041a00000a380030019800000ef10000c13d0000002001100039000000000112034f000000000101043b000009ba0010009c000000410000213d000000000010043f0000000b01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000400200043d001b00000002001d000009d90020009c000000460000213d000000000101043b0000001b050000290000008002500039000000400020043f000000000101041a000000600250003900000a3e001001980000000003000039000000010300c0390000000000320435000009b60210019700000000042504360000006002100270000009bb022001970000004003500039000c00000003001d00000000002304350000002001100270000009bb01100197000b00000004001d000000000014043500000002010003670000002402100370000000000202043b0000006402200039000000000121034f000000000101043b00000e8d0000613d000009ba0010009c000000410000213d0000000602000039000000000302041a000000400500043d00000024025000390000001d04000029000000000042043500000a6f020000410000000002250436001d00000002001d001a00000005001d000000040250003900000000001204350000000001000414000a00000003001d0000002002300270000009ba02200197000d00000002001d000000040020008c00000ef40000c13d0000000103000031000000400030008c0000004004000039000000000403401900000f1f0000013d0000000009750049000009c10090009c000000410000213d000000e00090008c000000410000413d000000400900043d000009c40090009c000000460000213d000000e00a9000390000004000a0043f00000000ab070434000009ba00b0009c000000410000213d000000000bb90436000000000a0a0433000009b600a0009c000000410000213d0000000000ab0435000000400a700039000000000a0a0433000009b600a0009c000000410000213d000000400b9000390000000000ab0435000000600a700039000000000a0a04330000ffff00a0008c000000410000213d000000600b9000390000000000ab0435000000800a700039000000000a0a0433000009b600a0009c000000410000213d000000800b9000390000000000ab0435000000a00a700039000000000a0a0433000009b600a0009c000000410000213d000000a00b9000390000000000ab0435000000c00a700039000000000a0a043300000000000a004b000000000b000039000000010b00c0390000000000ba004b000000410000c13d000000c00b9000390000000000ab04350000000006960436000000e007700039000000000087004b00000e550000413d0000020f0000013d000009ba0010009c000000410000213d000000400200043d00000a6e03000041000000000032043500000004032000390000000000130435000009b60020009c000009b6020080410000004001200210000009ec011001c7000026d50001043000000015010000290000000001010433000000e002300270000000000012004b00000cd10000413d0000000501000039000000000101041a000000a0011002700000ffff0110018f000000160010006c00000dfe0000413d00000015010000290000002001100039001400000001001d0000000001010433000000000001004b00000eae0000c13d0000000701000039000000000101041a00000a380010019800000ef10000c13d000000160000006b00000feb0000c13d00000002010003670000002402100370000000000202043b0000006402200039000000000121034f000000000101043b001d00000001001d000009ba0010009c000000410000213d00000a3201000041000000000010044300000000010004120000000400100443000000600100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d0000000203000367000000000101043b000009ba011001970000001d0010006b000010e70000c13d0000004401300370000000000101043b000009bc021001970000000e01000039000000000101041a000009bc031001970000000002320019001d00000002001d000009bc0020009c000004ea0000213d000009fb011001970000001d011001af0000000e02000039000000000012041b00000a3201000041000000000010044300000000010004120000000400100443000000400100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b000009bc011001970000001d0010006b000013110000a13d000000400100043d00000a5602000041000008410000013d000000400100043d00000a6d02000041000008410000013d0000001a02000029000009b60020009c000009b6020080410000004002200210000009b60010009c000009b601008041000000c001100210000000000121019f000009e6011001c70000000d0200002926d326ce0000040f0000006003100270000009b603300197000000400030008c000000400400003900000000040340190000001f0640018f00000060074001900000001a0570002900000f0e0000613d000000000801034f0000001a09000029000000008a08043c0000000009a90436000000000059004b00000f0a0000c13d000000000006004b00000f1b0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000000fdf0000613d0000001f01400039000000e00210018f0000001a01200029000000000021004b00000000020000390000000102004039000009bb0010009c000000460000213d0000000100200190000000460000c13d000000400010043f000000400030008c000000410000413d0000001a010000290000000001010433000f00000001001d00000a3b0010009c000000410000213d0000001d010000290000000001010433000900000001001d00000a3b0010009c000000410000213d00000002010003670000002402100370000000000202043b0000004403200039000000000331034f000000000403043b00000000030000310000000005230049000000230550008a000009c306500197000009c307400197000000000867013f000000000067004b0000000006000019000009c306004041000000000054004b0000000005000019000009c305008041000009c30080009c000000000605c019000000000006004b000000410000c13d00000000044200190000000404400039000000000541034f000000000505043b001100000005001d000009bb0050009c000000410000213d0000001105000029000000060550021000000000035300490000002007400039000009c304300197000009c305700197000000000645013f000000000045004b0000000004000019000009c304004041001000000007001d000000000037004b0000000003000019000009c303002041000009c30060009c000000000403c019000000000004004b000000410000c13d000000110000006b000011a40000c13d0000001b010000290000000001010433000009b601100197001a0a70001000d5001800000000001d001700000000001d0000000a0100002900000010011002700000ffff01100190000019980000c13d0000000501000039000000000201041a00000000010000190000001c030000290000000003030433000000b004200270000009b604400197000000000034001a000004ea0000413d00000002050003670000002406500370000000000706043b0000002406700039000000000665034f000000000806043b00000000060000310000000009760049000000230990008a000009c30a900197000009c30b800197000000000cab013f0000000000ab004b000000000a000019000009c30a004041000000000098004b0000000009000019000009c309008041000009c300c0009c000000000a09c01900000000000a004b000000410000c13d00000000078700190000000407700039000000000575034f000000000505043b000009bb0050009c000000410000213d00000000065600490000002007700039000000000067004b0000000008000019000009c308002041000009c306600197000009c307700197000000000967013f000000000067004b0000000006000019000009c306004041000009c30090009c000000000608c019000000000006004b000000410000c13d000000d0022002700000ffff0620018f00000000026500a9000000000005004b00000faf0000613d000009bf05500197000009bf0720019700000000055700d9000000000056004b000004ea0000c13d0000000003340019000000000032001a000004ea0000413d0000000003320019000000180030002a000004ea0000413d000000090200002900000a7404200198000000000200001900000fc80000613d000000180230002900000000034200a900000000044300d9000000000024004b000004ea0000c13d000000000003004b000000000200001900000fc80000613d0000000b020000290000000002020433000009bb0420019700000000023400a900000000033200d9000000000043004b000004ea0000c13d0000000c030000290000000003030433000009bb043001970000001a034000b90000001a0000006b00000fd10000613d0000001a053000fa000000000045004b000004ea0000c13d000000000023001a000004ea0000413d0000000002230019000000000012001a000004ea0000413d0000000f0300002900000a3b0330019800001a990000c13d00000a6201000041000000000010043f0000001201000039000000040010043f000009ec01000041000026d5000104300000001f0530018f000009b806300198000000400200043d000000000462001900000a1b0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00000fe60000c13d00000a1b0000013d001d00000000001d001a00000000001d00000ff30000013d0000001d020000290000000102200039001d00000002001d000000160020006c000011330000813d00000002010003670000002402100370000000000302043b0000004402300039000000000221034f000000000402043b00000000020000310000000005320049000000230550008a000009c306500197000009c307400197000000000867013f000000000067004b0000000006000019000009c306004041000000000054004b0000000005000019000009c305008041000009c30080009c000000000605c019000000000006004b000000410000c13d00000000034300190000000404300039000000000341034f000000000303043b000009bb0030009c000000410000213d000000060530021000000000055200490000002002400039000009c304500197000009c306200197000000000746013f000000000046004b0000000004000019000009c304004041000000000052004b0000000005000019000009c305002041000009c30070009c000000000405c019000000000004004b000000410000c13d0000001d0030006b000021300000813d0000001d03000029001b0006003002180000001b022000290000002003200039000000000331034f000000000303043b000000000003004b000011580000613d000000000121034f000000000101043b000009ba0010009c000000410000213d000000000010043f0000000c01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000101041a00000a390010019800000fee0000613d00000002010003670000002402100370000000000302043b0000004402300039000000000221034f000000000402043b00000000020000310000000005320049000000230550008a000009c306500197000009c307400197000000000867013f000000000067004b0000000006000019000009c306004041000000000054004b0000000005000019000009c305008041000009c30080009c000000000605c019000000000006004b000000410000c13d00000000034300190000000404300039000000000341034f000000000303043b000009bb0030009c000000410000213d000000060530021000000000055200490000002004400039000009c306500197000009c307400197000000000867013f000000000067004b0000000006000019000009c306004041000000000054004b0000000005000019000009c305002041000009c30080009c000000000605c019000000000006004b000000410000c13d0000001d0030006b000021300000813d0000001b034000290000000002320049000009c10020009c000000410000213d000000400020008c000000410000413d000000400200043d001900000002001d000009c50020009c000000460000213d0000000602000039000000000202041a00000019040000290000004004400039000000400040043f000000000431034f000000000404043b000009ba0040009c000000410000213d0000002002200270000009ba022001970000002003300039000000000131034f00000019030000290000000003430436000000000101043b001700000003001d0000000000130435000000400300043d00000a3a010000410000000001130436001800000001001d001b00000003001d000000040130003900000000004104350000000001000414000000040020008c000010960000c13d0000000103000031000000400030008c00000040040000390000000004034019000010bf0000013d0000001b03000029000009b60030009c000009b6030080410000004003300210000009b60010009c000009b601008041000000c001100210000000000131019f000009ec011001c726d326ce0000040f0000006003100270000009b603300197000000400030008c0000004004000039000000000403401900000060064001900000001b05600029000010ae0000613d000000000701034f0000001b08000029000000007907043c0000000008980436000000000058004b000010aa0000c13d0000001f07400190000010bb0000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f00030000000103550000000100200190000013050000613d0000001f01400039000000e00210018f0000001b01200029000000000021004b00000000020000390000000102004039000009bb0010009c000000460000213d0000000100200190000000460000c13d000000400010043f000000400030008c000000410000413d000009c50010009c000000460000213d0000004002100039000000400020043f0000001b02000029000000000202043300000a3b0020009c000000410000213d000000000121043600000018030000290000000003030433000009b60030009c000000410000213d0000000000310435000000000002004b0000115b0000613d0000001701000029000000000301043300000000012300a900000000022100d9000000000032004b000004ea0000c13d00000a3c0110012a0000001a0010002a000004ea0000413d001a001a0010002d00000fee0000013d0000002402300370000000000202043b0000006402200039000000000423034f0000000602000039000000000202041a000000000404043b000009ba0040009c000000410000213d0000002002200270000009ba02200197000000400600043d00000a46050000410000000000560435000000040560003900000000004504350000004403300370000000000303043b00000044046000390000000000140435001d00000006001d000000240160003900000000003104350000000001000414000000040020008c000011620000c13d0000000103000031000000200030008c000000200400003900000000040340190000118c0000013d00000008080000290000000009650049000009c10090009c000000410000213d000000a00090008c000000410000413d000000400900043d000009c20090009c000000460000213d000000a00a9000390000004000a0043f00000000ab060434000009ba00b0009c000000410000213d000000000bb90436000000000a0a0433000009b600a0009c000000410000213d0000000000ab0435000000400a600039000000000a0a0433000009bb00a0009c000000410000213d000000400b9000390000000000ab0435000000600a600039000000000a0a0433000009bb00a0009c000000410000213d000000600b9000390000000000ab0435000000800a600039000000000a0a043300000000000a004b000000000b000039000000010b00c0390000000000ba004b000000410000c13d000000800b9000390000000000ab04350000000008980436000000a006600039000000000076004b000011070000413d000001e50000013d0000001a0000006b00000eb00000613d0000000301000039000000000101041a001d00000001001d00000a3e0010019800000eb00000613d0000000401000039000000000101041a001b00000001001d000009c70100004100000000001004430000000001000414000009b60010009c000009b601008041000000c001100210000009c8011001c70000800b0200003926d326ce0000040f00000001002001900000212f0000613d0000001d020000290000008002200270000009b602200197000000000301043b000000000423004b000004ea0000413d0000001d01000029000009bf011001970000001b02000029000009bf0220019700001a520000613d000000000021004b00001a400000a13d000000400100043d00000a4002000041000008410000013d000000400100043d00000a4502000041000008410000013d00000019010000290000000001010433000000400200043d00000a3d030000410000000000320435000009ba0110019700000e920000013d0000001d03000029000009b60030009c000009b6030080410000004003300210000009b60010009c000009b601008041000000c001100210000000000131019f000009fa011001c726d326ce0000040f0000006003100270000009b603300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000001d057000290000117b0000613d000000000801034f0000001d09000029000000008a08043c0000000009a90436000000000059004b000011770000c13d000000000006004b000011880000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f000300000001035500000001002001900000198c0000613d0000001f01400039000000600210018f0000001d01200029000000000021004b00000000020000390000000102004039000009bb0010009c000000460000213d0000000100200190000000460000c13d000000400010043f000000200030008c000000410000413d0000001d010000290000000001010433000009bc021001970000000e01000039000000000101041a000009bc031001970000000002230019001d00000002001d000009bc0020009c000004ea0000213d00000ed70000013d0000006402200039000000000121034f000000000101043b000e00000001001d000009ba0010009c000000410000213d001900000000001d001a00000000001d001800000000001d001700000000001d000011b60000013d001a001a0010002d001700200020003d00000019020000290000000102200039001900000002001d000000110020006c00000f6d0000813d0000001901000029000000060110021000000010011000290000000002100079000009c10020009c000000410000213d000000400020008c000000410000413d000000400200043d001d00000002001d000009c50020009c000000460000213d0000001d020000290000004002200039000000400020043f0000000203000367000000000213034f000000000202043b000009ba0020009c000000410000213d0000002001100039000000000113034f0000001d030000290000000003230436000000000101043b001300000003001d0000000000130435000000400300043d00000a4c010000410000000000130435001b00000003001d0000000401300039000000000021043500000a3201000041000000000010044300000000010004120000000400100443000001000100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000201043b0000000001000414000009ba02200197000000040020008c000011f00000c13d0000000103000031000000200030008c00000020040000390000000004034019000012190000013d0000001b03000029000009b60030009c000009b6030080410000004003300210000009b60010009c000009b601008041000000c001100210000000000131019f000009ec011001c726d326ce0000040f0000006003100270000009b603300197000000200030008c0000002004000039000000000403401900000020064001900000001b05600029000012080000613d000000000701034f0000001b08000029000000007907043c0000000008980436000000000058004b000012040000c13d0000001f07400190000012150000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000001a8d0000613d0000001f01400039000000600210018f0000001b01200029000000000021004b00000000020000390000000102004039000009bb0010009c000000460000213d0000000100200190000000460000c13d000000400010043f000000200030008c000000410000413d0000001b020000290000000003020433000009ba0030009c000000410000213d0000001d020000290000000002020433000009ba02200197000000000003004b00001a3b0000613d000000000020043f0000000c01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000400200043d001b00000002001d000009c40020009c000000460000213d000000000101043b0000001b05000029000000e002500039000000400020043f000000000101041a000000c00250003900000a50031001980000000004000039000000010400c039000000000042043500000a39001001980000000002000039000000010200c039000000a0045000390000000000240435000009b60210019700000000042504360000007002100270000009b6022001970000008006500039001600000006001d00000000002604350000002002100270000009b602200197001400000004001d00000000002404350000005002100270000009b6022001970000006004500039001500000004001d0000000000240435000000400450003900000040011002700000ffff0110018f001200000004001d0000000000140435000000000003004b0000127c0000613d000000000001004b0000128f0000613d0000001d010000290000000001010433000009ba021001970000000e0020006c0000000f01000029000012cc0000613d000000400300043d00000a72010000410000000000130435001d00000003001d0000000401300039000000000021043500000000010004140000000d02000029000000040020008c000012910000c13d0000000103000031000000200030008c00000020040000390000000004034019000012bb0000013d0000000701000039000000000201041a0000ffff0120018f00000a70011000d10000001a0010002a000004ea0000413d0000001803000029000009b6033001970000001002200270000009b6022001970000000002320019001800000002001d000009b60020009c000004ea0000213d0000001702000029000009b60220019700000a710020009c000011af0000a13d000004ea0000013d0000000001000019000012e10000013d0000001d02000029000009b60020009c000009b6020080410000004002200210000009b60010009c000009b601008041000000c001100210000000000121019f000009ec011001c70000000d0200002926d326ce0000040f0000006003100270000009b603300197000000200030008c0000002004000039000000000403401900000020064001900000001d05600029000012aa0000613d000000000701034f0000001d08000029000000007907043c0000000008980436000000000058004b000012a60000c13d0000001f07400190000012b70000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000001aaf0000613d0000001f01400039000000600210018f0000001d01200029000000000021004b00000000020000390000000102004039000009bb0010009c000000460000213d0000000100200190000000460000c13d000000400010043f000000200030008c000000410000413d0000001d01000029000000000101043300000a3b0010009c000000410000213d00000a3b02100198000012dd0000613d0000001301000029000000000301043300000000012300a900000000022100d9000000000032004b000004ea0000c13d00000a3c0010009c0000000003000019000012de0000413d00000a3c0110012a000000120200002900000000020204330000ffff0220018f00000000031200a9000012de0000013d00000000030000190000001501000029000000000201043300000a730130012a0000001803000029000009b603300197000009b6022001970000000002320019001800000002001d000009b60020009c000004ea0000213d0000001702000029000009b60220019700000016030000290000000003030433000009b6033001970000000002230019001700000002001d000009b60020009c000004ea0000213d0000001b020000290000000002020433000009b60220019700000a70022000d1000000000021004b000012fd0000413d00000014020000290000000002020433000009b60220019700000a70022000d1000000000021004b000013010000a13d0000001a0020002a000004ea0000413d001a001a0020002d000011b10000013d0000001a0010002a000004ea0000413d001a001a0010002d000011b10000013d0000001f0530018f000009b806300198000000400200043d000000000462001900000a1b0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b0000130c0000c13d00000a1b0000013d00000014010000290000000001010433000000000001004b000013270000c13d00000a3201000041000000000010044300000000010004120000000400100443000000c00100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b000009ba0010019800001a0a0000c13d00000002020003670000002401200370000000000301043b0000000401300039000000000412034f000000000404043b00000000050000310000000006350049000000230660008a000009c307600197000009c308400197000000000978013f000000000078004b0000000007000019000009c307004041000000000064004b0000000006000019000009c306008041000009c30090009c000000000706c019000000000007004b000000410000c13d0000000001140019000000000412034f000000000404043b000009bb0040009c000000410000213d00000000064500490000002007100039000009c301600197000009c308700197000000000918013f000000000018004b0000000001000019000009c301004041000000000067004b0000000006000019000009c306002041000009c30090009c000000000106c019000000000001004b000000410000c13d0000001f0140003900000a7c011001970000003f0110003900000a7c06100197000000400100043d0000000006610019000000000016004b00000000080000390000000108004039000009bb0060009c000000460000213d0000000100800190000000460000c13d000000400060043f00000000064104360000000008740019000000000058004b000000410000213d000000000772034f00000a7c084001980000001f0940018f00000000058600190000136e0000613d000000000a07034f000000000b06001900000000ac0a043c000000000bcb043600000000005b004b0000136a0000c13d000000000009004b0000137b0000613d000000000787034f0000000308900210000000000905043300000000098901cf000000000989022f000000000707043b0000010008800089000000000787022f00000000078701cf000000000797019f0000000000750435000000000446001900000000000404350000000004010433000000200040008c00001a9c0000c13d0000000004060433001900000004001d000004000440008a00000a490040009c00001a9c0000213d0000000e01000039000000000101041a0000008004100270000009bb04400197000009bb0040009c000004ea0000613d00000a4b011001970000000104400039001100000004001d000000800440021000000a3104400197000000000114019f0000000e04000039000000000014041b00000015010000290000000001010433001000000001001d00000014010000290000000001010433000000000001004b001500000000001d000013b40000c13d0000001c01000029000000000010043f0000000d01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000201041a000009bb03200197000009bb0030009c000004ea0000613d00000a4802200197001500010030003d00000015022001af000000000021041b00000002020003670000002401200370000000000301043b0000006401300039000000000412034f000000000404043b001400000004001d000009ba0040009c000000410000213d000000400510008a000000000152034f000000000701043b00000000010000310000000004310049000000230440008a000009c306400197000009c308700197000000000968013f000000000068004b0000000008000019000009c308004041000000000047004b000000000a000019000009c30a008041000009c30090009c00000000080ac019000000000008004b000000410000c13d00000004033000390000000007370019000000000872034f000000000808043b001a00000008001d000009bb0080009c000000410000213d0000001a0810006a000000200b700039000009c307800197000009c309b00197000000000a79013f000000000079004b0000000007000019000009c307004041000c0000000b001d00000000008b004b0000000008000019000009c308002041000009c300a0009c000000000708c019000000000007004b000000410000c13d0000002005500039000000000552034f000000000505043b000009c307500197000000000867013f000000000067004b0000000006000019000009c306004041000000000045004b0000000004000019000009c304008041000009c30080009c000000000604c019000000000006004b000000410000c13d0000000003350019000000000232034f000000000202043b001700000002001d000009bb0020009c000000410000213d0000001702000029000b0006002002180000000b0110006a0000002005300039000009c302100197000009c303500197000000000423013f000000000023004b0000000002000019000009c302004041001b00000005001d000000000015004b0000000001000019000009c301002041000009c30040009c000000000201c019000000000002004b000000410000c13d000000160100002900000005011002100000003f02100039000009c002200197000000400300043d0000000002230019001d00000003001d000000000032004b00000000030000390000000103004039000009bb0020009c000000460000213d0000000100300190000000460000c13d000000400020043f00000016020000290000001d030000290000000000230435000000000002004b000014260000613d0000000002000019000000600400003900000020022000390000001d032000290000000000430435000000000012004b000014210000413d000000400100043d001800000001001d000009bd0010009c000000460000213d0000001801000029000001a001100039000000400010043f00000a3201000041000000000010044300000000010004120000000400100443000000800100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b0000001402000029000009ba022001970000001803000029000000e004300039001400000004001d00000000002404350000001502000029000009bb02200197000000c004300039001300000004001d00000000002404350000008002300039001200000002001d000000100400002900000000004204350000001102000029000009bb022001970000006004300039001000000004001d00000000002404350000001902000029000009ba022001970000004004300039000e00000004001d000000000024043500000020043000390000001c02000029000d00000004001d0000000000240435000009bb011001970000000000130435000000a001300039000f00000001001d000000000001043500000002010003670000004402100370000000000202043b0000010003300039001100000003001d00000000002304350000001a020000290000001f0220003900000a7c022001970000003f0220003900000a7c02200197000000400300043d0000000004230019000000000034004b00000000020000390000000102004039000009bb0040009c000000460000213d0000000100200190000000460000c13d0000000002000031000000400040043f0000001a0500002900000000045304360000000c05500029000000000025004b000000410000213d0000000c061003600000001a0800002900000a7c078001980000001f0880018f0000000005740019000014860000613d000000000906034f000000000a040019000000009b09043c000000000aba043600000000005a004b000014820000c13d000000000008004b000014930000613d000000000676034f0000000307800210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f00000000006504350000001a04400029000000000004043500000018040000290000012004400039000c00000004001d0000000000340435000000170300002900000005033002100000003f03300039000009c004300197000000400300043d0000000004430019000000000034004b00000000050000390000000105004039000009bb0040009c000000460000213d0000000100500190000000460000c13d000000400040043f000000170400002900000000004304350000000b050000290000001b04500029000000000024004b000000410000213d0000001b0040006b000014c90000813d00000000050300190000001b0620006a000009c10060009c000000410000213d000000400060008c000000410000413d000000400600043d000009c50060009c000000460000213d0000004007600039000000400070043f0000001b07100360000000000707043b000009ba0070009c000000410000213d000000200550003900000000077604360000001b090000290000002008900039000000000881034f000000000808043b00000000008704350000000000650435001b00400090003d0000001b0040006b000014b00000413d00000018010000290000016002100039001500000002001d0000001d0400002900000000004204350000014002100039000b00000002001d00000000003204350000018001100039001700000001001d0000000000010435000000160000006b000018000000613d001a00000000001d00000002010003670000002402100370000000000302043b0000004402300039000000000221034f000000000402043b00000000020000310000000005320049000000230550008a000009c306500197000009c307400197000000000867013f000000000067004b0000000006000019000009c306004041000000000054004b0000000005000019000009c305008041000009c30080009c000000000605c019000000000006004b000000410000c13d00000000034300190000000404300039000000000341034f000000000303043b000009bb0030009c000000410000213d000000060530021000000000055200490000002004400039000009c306500197000009c307400197000000000867013f000000000067004b0000000006000019000009c306004041000000000054004b0000000005000019000009c305002041000009c30080009c000000000605c019000000000006004b000000410000c13d0000001a0030006b000021300000813d0000001a03000029000000060330021000000000033400190000000002320049000009c10020009c000000410000213d000000400020008c000000410000413d000000400200043d001c00000002001d000009c50020009c000000460000213d0000001c020000290000004002200039000000400020043f000000000231034f000000000202043b000009ba0020009c000000410000213d0000002003300039000000000131034f0000001c030000290000000003230436000000000101043b000a00000003001d0000000000130435000000400300043d00000a4c010000410000000000130435001b00000003001d0000000401300039000000000021043500000a3201000041000000000010044300000000010004120000000400100443000001000100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000201043b0000000001000414000009ba02200197000000040020008c0000153e0000c13d0000000103000031000000200030008c00000020040000390000000004034019000015670000013d0000001b03000029000009b60030009c000009b6030080410000004003300210000009b60010009c000009b601008041000000c001100210000000000131019f000009ec011001c726d326ce0000040f0000006003100270000009b603300197000000200030008c0000002004000039000000000403401900000020064001900000001b05600029000015560000613d000000000701034f0000001b08000029000000007907043c0000000008980436000000000058004b000015520000c13d0000001f07400190000015630000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000001baf0000613d0000001f01400039000000600110018f0000001b02100029000000000012004b00000000010000390000000101004039001d00000002001d000009bb0020009c000000460000213d0000000100100190000000460000c13d0000001d01000029000000400010043f000000200030008c000000410000413d0000001b010000290000000001010433001900000001001d000009ba0010009c000000410000213d000000190000006b00001bc80000613d00000a4d010000410000001d020000290000000000120435000000040120003900000a4e02000041000000000021043500000000010004140000001902000029000000040020008c0000002004000039000015b20000613d0000001d02000029000009b60020009c000009b6020080410000004002200210000009b60010009c000009b601008041000000c001100210000000000121019f000009ec011001c7000000190200002926d326ce0000040f0000006003100270000009b603300197000000200030008c0000002004000039000000000403401900000020064001900000001d05600029000015a10000613d000000000701034f0000001d08000029000000007907043c0000000008980436000000000058004b0000159d0000c13d0000001f07400190000015ae0000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000001bbb0000613d0000001f01400039000000600110018f0000001d01100029001b00000001001d000009bb0010009c000000460000213d0000001b01000029000000400010043f000000200030008c000000410000413d0000001d010000290000000001010433000000000001004b0000000002000039000000010200c039000000000021004b000000410000c13d000000000001004b00001bc70000613d00000002020003670000002401200370000000000501043b0000000401500039000000000312034f000000000403043b00000000030000310000000005530049000000230550008a000009c306500197000009c307400197000000000867013f000000000067004b0000000006000019000009c306004041000000000054004b0000000005000019000009c305008041000009c30080009c000000000605c019000000000006004b000000410000c13d0000000004140019000000000142034f000000000101043b000009bb0010009c000000410000213d00000000061300490000002005400039000009c304600197000009c307500197000000000847013f000000000047004b0000000004000019000009c304004041000000000065004b0000000006000019000009c306002041000009c30080009c000000000406c019000000000004004b000000410000c13d0000001b04000029000009c20040009c000000460000213d0000000a040000290000000004040433000a00000004001d0000001c040000290000000004040433000900000004001d0000001f0410003900000a7c044001970000003f0440003900000a7c064001970000001b04000029000000a004400039000000400040043f0000000006460019000009bb0060009c000000460000213d000000400060043f00000000001404350000000006510019000000000036004b000000410000213d000000000552034f00000a7c061001980000001b02000029000000c0022000390000000003620019000016130000613d000000000705034f0000000008020019000000007907043c0000000008980436000000000038004b0000160f0000c13d0000001f07100190000016200000613d000000000565034f0000000306700210000000000703043300000000076701cf000000000767022f000000000505043b0000010006600089000000000565022f00000000056501cf000000000575019f0000000000530435000000000121001900000000000104350000001b010000290000000001410436001d00000001001d00000a3201000041000000000010044300000000010004120000000400100443000000a00100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d0000000902000029000009ba02200197000000000101043b000009bb011001970000001d030000290000000000130435000009ba010000410000001e0310017f0000001b060000290000008001600039000000000021043500000060026000390000000a04000029000000000042043500000040046000390000000000340435000000400700043d00000a4f03000041000000000037043500000004037000390000002005000039000000000053043500000000030604330000002405700039000000a0060000390000000000650435000000c40670003900000000530304340000000000360435001b00000007001d000000e406700039000000000003004b0000165d0000613d000000000700001900000000086700190000000009750019000000000909043300000000009804350000002007700039000000000037004b000016560000413d000000000563001900000000000504350000001d050000290000000005050433000009bb055001970000001b07000029000000440670003900000000005604350000000004040433000009ba0440019700000064057000390000000000450435000000000202043300000084047000390000000000240435000000a4027000390000000001010433000009ba01100197000000000012043500000000010004140000001902000029000000040020008c000016770000c13d000000030100036700000001030000310000168e0000013d0000001f0230003900000a7c02200197000000e402200039000009b60020009c000009b60200804100000060022002100000001b03000029000009b60030009c000009b6030080410000004003300210000000000232019f000009b60010009c000009b601008041000000c001100210000000000121019f000000190200002926d326c90000040f0000006003100270000109b60030019d000009b6033001970003000000010355000000010020019000001bd50000613d00000a7c043001980000001b02400029000016970000613d000000000501034f0000001b06000029000000005705043c0000000006760436000000000026004b000016930000c13d0000001f05300190000016a40000613d000000000141034f0000000304500210000000000502043300000000054501cf000000000545022f000000000101043b0000010004400089000000000141022f00000000014101cf000000000151019f00000000001204350000001f0130003900000a7c011001970000001b02100029000000000012004b00000000010000390000000101004039001d00000002001d000009bb0020009c000000460000213d0000000100100190000000460000c13d0000001d01000029000000400010043f000009c10030009c000000410000213d000000200030008c000000410000413d0000001b010000290000000002010433000009bb0020009c000000410000213d0000001b013000290000001b022000290000000003210049000009c10030009c000000410000213d000000400030008c000000410000413d0000001d03000029000009c50030009c000000460000213d0000001d030000290000004003300039000000400030043f0000000054020434000009bb0040009c000000410000213d00000000062400190000001f04600039000000000014004b0000000007000019000009c307008041000009c308400197000009c304100197000000000948013f000000000048004b0000000008000019000009c308004041000009c30090009c000000000807c019000000000008004b000000410000c13d0000000076060434000009bb0060009c000000460000213d0000001f0860003900000a7c088001970000003f0880003900000a7c088001970000000008380019000009bb0080009c000000460000213d000000400080043f00000000006304350000000008760019000000000018004b000000410000213d0000001d080000290000006008800039000000000006004b000016f30000613d0000000009000019000000000a890019000000000b790019000000000b0b04330000000000ba04350000002009900039000000000069004b000016ec0000413d000000000686001900000000000604350000001d060000290000000003360436001b00000003001d0000000003050433000009bb0030009c000000410000213d00000000022300190000001f03200039000000000013004b0000000005000019000009c305008041000009c303300197000000000643013f000000000043004b0000000003000019000009c303004041000009c30060009c000000000305c019000000000003004b000000410000c13d0000000032020434000009bb0020009c000000460000213d0000001f0420003900000a7c044001970000003f0440003900000a7c05400197000000400400043d0000000005540019000000000045004b00000000060000390000000106004039000009bb0050009c000000460000213d0000000100600190000000460000c13d000000400050043f00000000052404360000000006320019000000000016004b000000410000213d000000000002004b000017280000613d000000000100001900000000065100190000000007310019000000000707043300000000007604350000002001100039000000000021004b000017210000413d000000000125001900000000000104350000001b0100002900000000004104350000000001040433000a00000001001d000000200010008c000017450000a13d0000001c010000290000000001010433000009ba01100197000000000010043f0000000c01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000101041a0000007001100270000009b6011001970000000a0010006b00001be10000213d0000001d0100002900000000010104330000000023010434000000200030008c00001a9c0000c13d0000000002020433000004000220008a00000a490020009c00001a9c0000213d000000400300043d00000020013000390000001902000029000000000021043500000020010000390000000000130435001900000003001d000009c50030009c000000460000213d00000019010000290000004001100039000000400010043f0000001b010000290000000001010433001b00000001001d0000001d010000290000000001010433001d00000001001d0000001c010000290000000001010433000009ba01100197000000000010043f0000000c01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000101041a00000a500010019800000007020000390000001001000039000017840000613d0000001c010000290000000001010433000009ba01100197000000000010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000201043b0000005001000039000000400500043d000009d90050009c000000460000213d000000000202041a0000008003500039000000400030043f00000040035000390000001b04000029000000000043043500000020045000390000001d06000029000000000064043500000019060000290000000000650435000000000112022f000009b60110019700000060025000390000000000120435000000400100043d0000002006100039000000200700003900000000007604350000000005050433000000400610003900000080070000390000000000760435000000c00810003900000000760504340000000000680435000000e005100039000000000006004b000017ac0000613d00000000080000190000000009580019000000000a870019000000000a0a04330000000000a904350000002008800039000000000068004b000017a50000413d000000000756001900000000000704350000001f0660003900000a7c0660019700000000040404330000006007100039000000a0086000390000000000870435000000000756001900000000650404340000000004570436000000000005004b000017c10000613d000000000700001900000000084700190000000009760019000000000909043300000000009804350000002007700039000000000057004b000017ba0000413d000000000645001900000000000604350000001f0550003900000a7c0550019700000000044500190000000005140049000000400550008a00000000030304330000008006100039000000000056043500000000530304340000000004340436000000000003004b000017d70000613d000000000600001900000000074600190000000008650019000000000808043300000000008704350000002006600039000000000036004b000017d00000413d000000000543001900000000000504350000000002020433000009b602200197000000a005100039000000000025043500000000021400490000001f0330003900000a7c033001970000000002320019000000200320008a00000000003104350000001f0220003900000a7c032001970000000002130019000000000032004b00000000030000390000000103004039000009bb0020009c000000460000213d0000000100300190000000460000c13d000000400020043f0000001502000029000000000202043300000000030204330000001a0030006c000021300000a13d0000001a0400002900000005034002100000000003230019000000200330003900000000001304350000000001020433000000000041004b000021300000a13d0000001a020000290000000102200039001a00000002001d000000160020006c000014d70000413d0000000d0100002900000000010104330000000e020000290000000002020433000000100300002900000000030304330000000f04000029000000000404043300000013050000290000000005050433000000140600002900000000060604330000001207000029000000000707043300000011080000290000000008080433000000400a00043d0000010009a0003900000000008904350000008008a000390000000000780435000009ba06600197000000e007a000390000000000670435000009bb05500197000000c006a000390000000000560435000000000004004b0000000004000039000000010400c039000000a005a000390000000000450435000009bb033001970000006004a000390000000000340435000009ba022001970000004003a000390000000000230435000001000200003900000000022a0436000009ba011001970000000000120435001d0000000a001d00000a5300a0009c000000460000213d0000001d030000290000012001300039001c00000001001d000000400010043f000009b60020009c000009b60200804100000040012002100000000002030433000009b60020009c000009b6020080410000006002200210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f000009cd011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d0000000c0200002900000000020204330000002003200039000009b60030009c000009b60300804100000040033002100000000002020433000009b60020009c000009b6020080410000006002200210000000000232019f000000000101043b001b00000001001d0000000001000414000009b60010009c000009b601008041000000c001100210000000000121019f000009cd011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b001a00000001001d0000000b0100002900000000020104330000001d050000290000014001500039000000200300003900000000003104350000016003500039000000000402043300000000004304350000018003500039000000000004004b000018740000613d0000000005000019000000200220003900000000060204330000000076060434000009ba0660019700000000066304360000000007070433000000000076043500000040033000390000000105500039000000000045004b000018690000413d0000001d0230006a000001400320008a0000001c040000290000000000340435000001010220008a00000a7c032001970000000002430019000000000032004b00000000030000390000000103004039000009bb0020009c000000460000213d0000000100300190000000460000c13d000000400020043f000009b60010009c000009b60100804100000040011002100000001c020000290000000002020433000009b60020009c000009b6020080410000006002200210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f000009cd011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b001c00000001001d00000015010000290000000003010433000000400100043d000000200210003900000020040000390000000000420435000000000403043300000005054002100000000006510019000000400510003900000000004504350000006007600039000000000004004b00001be60000c13d0000000003170049000000200430008a00000000004104350000001f0330003900000a7c043001970000000003140019000000000043004b00000000040000390000000104004039000009bb0030009c000000460000213d0000000100400190000000460000c13d000000400030043f000009b60020009c000009b60200804100000040022002100000000001010433000009b60010009c000009b6010080410000006001100210000000000121019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f000009cd011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b001600000001001d000000400100043d001d00000001001d0000002001100039001900000001001d000000000001043500000a320100004100000000001004430000000001000412000000040010044300000024000004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b0000001d04000029000000c00240003900000016030000290000000000320435000000a0024000390000001c03000029000000000032043500000080024000390000001a03000029000000000032043500000060024000390000001b03000029000000000032043500000040024000390000000000120435000000c0010000390000000000140435000009c40040009c000000460000213d0000001d02000029000000e001200039000000400010043f0000001901000029000009b60010009c000009b60100804100000040011002100000000002020433000009b60020009c000009b6020080410000006002200210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f000009cd011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b00000017020000290000000000120435000000400100043d0000002002000039000000000221043600000018030000290000000003030433000009bb0330019700000000003204350000000d020000290000000002020433000009ba02200197000000400310003900000000002304350000000e020000290000000002020433000009ba022001970000006003100039000000000023043500000010020000290000000002020433000009bb022001970000008003100039000000000023043500000012020000290000000002020433000000a00310003900000000002304350000000f020000290000000002020433000000000002004b0000000002000039000000010200c039000000c003100039000000000023043500000013020000290000000002020433000009bb02200197000000e003100039000000000023043500000014020000290000000002020433000009ba022001970000010003100039000000000023043500000011020000290000000002020433000001200310003900000000002304350000000c0200002900000000020204330000014003100039000001a0040000390000000000430435000001c00510003900000000430204340000000000350435000001e002100039000000000003004b0000194a0000613d000000000500001900000000062500190000000007540019000000000707043300000000007604350000002005500039000000000035004b000019430000413d000000000423001900000000000404350000001f0330003900000a7c043001970000000b0300002900000000030304330000016005100039000001c0064000390000000000650435000000000224001900000000040304330000000002420436000000000004004b000019640000613d0000000005000019000000200330003900000000060304330000000076060434000009ba0660019700000000066204360000000007070433000000000076043500000040022000390000000105500039000000000045004b000019590000413d0000000003120049000000200430008a000000150300002900000000030304330000018005100039000000000045043500000000040304330000000000420435000000050540021000000000055200190000002007500039000000000004004b00001c030000c13d00000017020000290000000002020433000001a00310003900000000002304350000000002170049000009b60020009c000009b6020080410000006002200210000009b60010009c000009b6010080410000004001100210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f000009cd011001c70000800d02000039000000010300003900000a540400004126d326c90000040f0000000100200190000000410000613d00000017010000290000000001010433000007640000013d0000001f0530018f000009b806300198000000400200043d000000000462001900000a1b0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000019930000c13d00000a1b0000013d00000002030003670000002402300370000000000402043b0000002407400039000000000273034f000000000202043b0000000405400039000000000400003100000000065400490000001f0660008a000009c308600197000009c309200197000000000a89013f000000000089004b0000000009000019000009c309004041000000000062004b000000000b000019000009c30b008041000009c300a0009c00000000090bc019000000000009004b000000410000c13d0000000009520019000000000293034f000000000202043b000009bb0020009c000000410000213d000000000a24004900000020099000390000000000a9004b000000000b000019000009c30b002041000009c30aa00197000009c309900197000000000ca9013f0000000000a9004b0000000009000019000009c309004041000009c300c0009c00000000090bc019000000000009004b000000410000c13d0000002007700039000000000773034f000000000707043b000009c309700197000000000a89013f000000000089004b0000000008000019000009c308004041000000000067004b0000000006000019000009c306008041000009c300a0009c000000000806c019000000000008004b000000410000c13d0000000005570019000000000353034f000000000303043b000009bb0030009c000000410000213d000000060630021000000000046400490000002005500039000000000045004b0000000006000019000009c306002041000009c304400197000009c305500197000000000745013f000000000045004b0000000004000019000009c304004041000009c30070009c000000000406c019000000000004004b000000410000c13d00000180043000c9000000000003004b000019ef0000613d000009bf03300197000009bf0540019700000000033500d9000001800030008c000004ea0000c13d0000001702200029000000000242001900000220022000390000000a030000290000ffff0430018f000009bf0520019700000000034200a9000009bf0230019700000000025200d9000000000024004b000004ea0000c13d0000000502000039000000000202041a000000e004200270000000000334001a00000f730000613d0000000904000029000000700440027000000a740440019800000f730000613d00000000011400a900000000033100a900000a75013000d100000000033100d900000a750030009c000004ea0000c13d00000f740000013d0000001c01000029000000000010043f0000000d01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000101041a000009bb00100198000013270000c13d000000400200043d00000a47010000410000000000120435001d00000002001d00000004012000390000001c02000029000000000021043500000a3201000041000000000010044300000000010004120000000400100443000000c00100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f00000001002001900000212f0000613d000000000201043b0000000001000414000009ba02200197000000040020008c00001abb0000c13d0000000103000031000000200030008c0000002004000039000000000403401900001ae50000013d00000a5203000041000000000031043500000004031000390000000000230435000008500000013d0000001b05000029000000800650027000000000056400a900000000044500d9000000000064004b000004ea0000c13d000000000015001a000004ea0000413d00000000011500190000001d0400002900000a3f044001970000008003300210000009ca03300197000000000343019f0000000304000039000000000034041b000000000012004b00000000010240190000001a0020006c00001a5d0000813d000000400100043d00000024031000390000001a04000029000000000043043500000a4403000041000000000031043500000004031000390000000000230435000006140000013d0000001a0210006c00001a740000813d0000001b020000290000008002200272000004ea0000613d0000001a03100069000000010420008a000000000034001a000004ea0000413d0000000003340019000000400400043d0000002405400039000000000015043500000a4301000041000000000014043500000000012300d900000004024000390000000000120435000009b60040009c000009b6040080410000004001400210000009e6011001c7000026d500010430000009bf012001970000000303000039000000000203041a00000a4102200197000000000112019f000000000013041b000000400100043d0000001a020000290000000000210435000009b60010009c000009b60100804100000040011002100000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f000009f7011001c70000800d02000039000000010300003900000a420400004126d326c90000040f000000010020019000000eb00000c13d000000410000013d0000001f0530018f000009b806300198000000400200043d000000000462001900000a1b0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001a940000c13d00000a1b0000013d000000000112001900000000013100d9000007640000013d000000400400043d001d00000004001d00000a4a020000410000000000240435000000040240003900000020030000390000000000320435000000240240003926d321cc0000040f0000001d020000290000000001210049000009b60010009c000009b601008041000009b60020009c000009b60200804100000060011002100000004002200210000000000121019f000026d5000104300000001f0530018f000009b806300198000000400200043d000000000462001900000a1b0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001ab60000c13d00000a1b0000013d0000001d03000029000009b60030009c000009b6030080410000004003300210000009b60010009c000009b601008041000000c001100210000000000131019f000009ec011001c726d326ce0000040f0000006003100270000009b603300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000001d0570002900001ad40000613d000000000801034f0000001d09000029000000008a08043c0000000009a90436000000000059004b00001ad00000c13d000000000006004b00001ae10000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000001c210000613d0000001f01400039000000600210018f0000001d01200029000000000021004b00000000020000390000000102004039000009bb0010009c000000460000213d0000000100200190000000460000c13d000000400010043f000000200030008c000000410000413d0000001d010000290000000001010433001d00000001001d000009bb0010009c000000410000213d0000001c01000029000000000010043f0000000d01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000201041a00000a48022001970000001d022001af000000000021041b000013270000013d000000000300041a000009c6033001970000000005000411000000000353019f000000000030041b0000000001010433000300000001001d000000000001004b0000000001000039000000010100c039000200000001001d0000000601000029000009c20010009c000000460000213d0000000001040433000500000001001d0000000001020433000009bf031001970000000602000029000000a001200039000000400010043f000400000003001d0000000001320436000100000001001d000009c70100004100000000001004430000000001000414000009b60010009c000009b601008041000000c001100210000009c8011001c70000800b0200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b0000000505000029000009bf02500197000000060400002900000080034000390000000000230435000000400240003900000002030000290000000000320435000009b60210019700000001030000290000000000230435000000600240003900000004040000290000000000420435000000030000006b0000000002000019000009c90200c0410000008001100210000009ca01100197000000000112019f0000000302000039000000000302041a000009cb03300197000000000131019f000000000141019f000000000012041b0000008001500210000000000114019f0000000402000039000000000012041b00000007010000290000000001010433000009ba0010019800001bac0000613d0000001d010000290000000001010433000009bb0110019800001bac0000613d0000001c020000290000000002020433000009bb0320019800001bac0000613d0000001b020000290000000002020433000009bb0020019800001bac0000613d00000018020000290000000002020433000009ba0020019800001bac0000613d00000016020000290000000002020433000009ba0020019800001bac0000613d000000400200043d000000600420003900000000003404350000004003200039000000000013043500000080010000390000000001120436000000800320003900000000040004100000000000430435000009cc030000410000000000310435000009c20020009c000000460000213d000000a003200039000000400030043f000009b60010009c000009b60100804100000040011002100000000002020433000009b60020009c000009b6020080410000006002200210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f000009cd011001c7000080100200003926d326ce0000040f00000007030000290000000100200190000000410000613d000000000101043b000000800010043f0000000001030433000009ba01100197000000e00010043f0000001d020000290000000002020433000009bb02200197000001000020043f0000001c030000290000000003030433000009bb03300197000001200030043f0000001b040000290000000004040433000009bb04400197000000a00040043f0000001a050000290000000005050433000009bc05500197000000c00050043f00000019060000290000000006060433000009ba06600197000001400060043f00000018070000290000000007070433000009ba07700197000001600070043f00000016080000290000000008080433000009ba08800197000001800080043f0000000f090000290000000009090433000009ba0090019800001c2d0000c13d000000400100043d00000a0302000041000008410000013d0000001f0530018f000009b806300198000000400200043d000000000462001900000a1b0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001bb60000c13d00000a1b0000013d0000001f0530018f000009b806300198000000400200043d000000000462001900000a1b0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001bc20000c13d00000a1b0000013d001d001b0000002d0000001c01000029000000000101043300000a52020000410000001d030000290000000000230435000009ba0110019700000004023000390000000000120435000009b60030009c000009b6030080410000004001300210000009ec011001c7000026d5000104300000001f0530018f000009b806300198000000400200043d000000000462001900000a1b0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001bdc0000c13d00000a1b0000013d0000001c010000290000000001010433000000400200043d00000a51030000410000115f0000013d000000000600001900001bf00000013d000000000978001900000000000904350000001f0880003900000a7c0880019700000000077800190000000106600039000000000046004b000018a60000813d0000000008170049000000600880008a000000200550003900000000008504350000002003300039000000000803043300000000980804340000000007870436000000000008004b00001be80000613d000000000a000019000000000b7a0019000000000ca90019000000000c0c04330000000000cb0435000000200aa0003900000000008a004b00001bfb0000413d00001be80000013d0000000005000019000000000602001900001c0e0000013d000000000978001900000000000904350000001f0880003900000a7c0880019700000000077800190000000105500039000000000045004b000019710000813d0000000008270049000000200880008a000000200660003900000000008604350000002003300039000000000803043300000000980804340000000007870436000000000008004b00001c060000613d000000000a000019000000000b7a0019000000000ca90019000000000c0c04330000000000cb0435000000200aa0003900000000008a004b00001c190000413d00001c060000013d0000001f0530018f000009b806300198000000400200043d000000000462001900000a1b0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001c280000c13d00000a1b0000013d000000170a000029000000000a0a0433000009ba0aa00197000000150b000029000000000b0b0433000000a00bb00210000009ce0bb00197000000000aab019f000000140b000029000000000b0b0433000000b00bb00210000009cf0bb00197000000000aba019f000000130b000029000000000b0b0433000000d00bb00210000009d00bb00197000000000aba019f000000120b000029000000000b0b0433000000e00bb00210000000000aba019f000000050b0000390000000000ab041b0000002009900210000009d109900197000000110a000029000000000a0a04330000ffff0aa0018f00000000099a019f000000100a000029000000000a0a0433000000100aa00210000009d20aa001970000000009a9019f0000000e0a000029000000000a0a0433000000c00aa00210000009d30aa001970000000009a9019f0000000d0a000029000000000a0a0433000000e00aa002100000000009a9019f000000060a00003900000000009a041b0000000c0900002900000000090904330000ffff0990018f000000070a000039000000000b0a041a000009d40bb0019700000000099b019f0000000b0b000029000000000b0b0433000000100bb00210000009d50bb001970000000009b9019f0000000a0b000029000000000b0b043300000000000b004b000009d60b000041000000000b0060190000000009b9019f00000000009a041b000000400a00043d000009b900a0009c000000460000213d0000010009a00039000000400090043f000000e009a000390000000000890435000000c008a000390000000000780435000000a007a0003900000000006704350000008006a0003900000000005604350000006005a0003900000000004504350000004004a0003900000000003404350000002003a00039000000000023043500000000001a0435000000400200043d00000000011204360000000003030433000009bb0330019700000000003104350000000001040433000009bb01100197000000400320003900000000001304350000000001050433000009bb01100197000000600320003900000000001304350000000001060433000009bc01100197000000800320003900000000001304350000000001070433000009ba01100197000000a00320003900000000001304350000000001080433000009ba01100197000000c00320003900000000001304350000000001090433000009ba01100197000000e003200039000000000013043500000017010000290000000001010433000009ba0110019700000100032000390000000000130435000000150100002900000000010104330000ffff0110018f0000012003200039000000000013043500000014010000290000000001010433000009b60110019700000140032000390000000000130435000000130100002900000000010104330000ffff0110018f0000016003200039000000000013043500000012010000290000000001010433000009b60110019700000180032000390000000000130435000000110100002900000000010104330000ffff0110018f000001a0032000390000000000130435000000100100002900000000010104330000ffff0110018f000001c00320003900000000001304350000000f010000290000000001010433000009ba01100197000001e00320003900000000001304350000000e010000290000000001010433000009b601100197000002000320003900000000001304350000000d010000290000000001010433000009b601100197000002200320003900000000001304350000000c0100002900000000010104330000ffff0110018f000002400320003900000000001304350000000b010000290000000001010433000009b601100197000002600320003900000000001304350000000a010000290000000001010433000000000001004b0000000001000039000000010100c03900000280032000390000000000130435000009b60020009c000009b60200804100000040012002100000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f000009d7011001c70000800d020000390000000103000039000009d80400004126d326c90000040f0000000100200190000000410000613d00000009010000290000000001010433000000000001004b00001d440000613d001d00000000001d000000400100043d001c00000001001d000009d90010009c000000460000213d0000001d01000029000000050110021000000008011000290000000001010433000000400210003900000000020204330000002003100039000000000303043300000060041000390000000004040433000000800510003900000000050504330000001c070000290000008006700039000000400060043f000000000005004b0000000005000039000000010500c0390000006006700039001b00000006001d0000000000560435000009bb044001970000004005700039001a00000005001d0000000000450435000009b6033001970000000003370436000009bb02200197001900000003001d00000000002304350000000001010433000009ba01100197000000000010043f0000000b01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d0000001c020000290000000002020433000009b602200197000000000101043b000000000301041a000009cb03300197000000000223019f000000190300002900000000030304330000002003300210000009db03300197000000000232019f0000001a0300002900000000030304330000006003300210000009dc03300197000000000232019f0000001b030000290000000003030433000000000003004b000009c9030000410000000003006019000000000232019f000000000021041b0000001d02000029001d00010020003d000000090100002900000000010104330000001d0010006b00001cf60000413d000000400100043d000000200200003900000000022104360000000903000029000000000303043300000000003204350000004002100039000000000003004b00001d6b0000613d0000003304000029000000000500001900000000460404340000000087060434000009ba0770019700000000077204360000000008080433000009b608800197000000000087043500000040076000390000000007070433000009bb077001970000004008200039000000000078043500000060076000390000000007070433000009bb077001970000006008200039000000000078043500000080066000390000000006060433000000000006004b0000000006000039000000010600c03900000080072000390000000000670435000000a0022000390000000105500039000000000035004b00001d4f0000413d0000000002120049000009b60020009c000009b6020080410000006002200210000009b60010009c000009b6010080410000004001100210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000121019f000009cd011001c70000800d020000390000000103000039000009dd0400004126d326c90000040f0000000100200190000000410000613d000000400100043d001a00000001001d000009de0010009c000000460000213d0000001a020000290000002001200039001d00000001001d000000400010043f00000000000204350000003201000029001300000001001d0000000021010434001200000002001d000000000001004b00001e000000613d001b00000000001d0000001b0100002900000005011002100000001201100029000000400200043d001c00000002001d0000000001010433000000a0021000390000000002020433000009b6022001970000001f0020008c00001eba0000a13d0000001c03000029000009c40030009c000000460000213d00000040031000390000000003030433000000200410003900000000040404330000006005100039000000000505043300000080061000390000000006060433000000c00710003900000000070704330000001c09000029000000e008900039000000400080043f000000c00a900039000000010800003900190000000a001d00000000008a0435000000000007004b0000000007000039000000010700c039000000a008900039001800000008001d00000000007804350000008007900039001700000007001d0000000000270435000009b6026001970000006006900039001600000006001d00000000002604350000ffff0250018f0000004005900039001500000005001d0000000000250435000009b6024001970000000004290436000009b602300197001400000004001d00000000002404350000000001010433000009ba01100197000000000010043f0000000c01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d0000001c020000290000000002020433000009b602200197000000000101043b000000000301041a000009c603300197000000000223019f000000140300002900000000030304330000002003300210000009df03300197000000000232019f000000150300002900000000030304330000004003300210000009e003300197000000000232019f000000160300002900000000030304330000005003300210000009e103300197000000000232019f000000170300002900000000030304330000007003300210000009e203300197000000000232019f00000018030000290000000003030433000000000003004b000009e3030000410000000003006019000000000232019f00000019030000290000000003030433000000000003004b000009e4030000410000000003006019000000000232019f000000000021041b0000001b02000029001b00010020003d000000130100002900000000010104330000001b0010006b00001d8f0000413d000000400100043d000000200200003900000000022104360000001303000029000000000303043300000000003204350000004002100039000000000003004b00001e310000613d0000003104000029000000000500001900000000460404340000000087060434000009ba0770019700000000077204360000000008080433000009b608800197000000000087043500000040076000390000000007070433000009b60770019700000040082000390000000000780435000000600760003900000000070704330000ffff0770018f0000006008200039000000000078043500000080076000390000000007070433000009b60770019700000080082000390000000000780435000000a0076000390000000007070433000009b607700197000000a0082000390000000000780435000000c0066000390000000006060433000000000006004b0000000006000039000000010600c039000000c0072000390000000000670435000000e0022000390000000105500039000000000035004b00001e0b0000413d0000000002120049000009b60020009c000009b6020080410000006002200210000009b60010009c000009b6010080410000004001100210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000121019f000009cd011001c70000800d020000390000000103000039000009e70400004126d326c90000040f0000000100200190000000410000613d0000001a010000290000000001010433000000000001004b00001e8b0000613d001c00000000001d0000001c0100002900000005011002100000001d011000290000000001010433000009ba01100197000000000010043f0000000c01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000001041b0000001c02000029001c00010020003d0000001a0100002900000000010104330000001c0010006b00001e4a0000413d000000000001004b00001e8b0000613d000000400100043d000000200200003900000000022104360000001a03000029000000000303043300000000003204350000004002100039000000000003004b00001e770000613d00000000040000190000001d060000290000000065060434001d00000006001d000009ba0550019700000000025204360000000104400039000000000034004b00001e6f0000413d0000000002120049000009b60020009c000009b6020080410000006002200210000009b60010009c000009b6010080410000004001100210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000121019f000009cd011001c70000800d020000390000000103000039000009e80400004126d326c90000040f0000000100200190000000410000613d0000003001000029001700000001001d0000000021010434001300000002001d001600000001001d000000400010008c0000075d0000213d0000000e01000039000000000101041a000009bc021001970000006001100270000009b601100197001200000001001d000000010110008a001c00000002001d000000000021004b00001ffc0000813d000000000100041a000009ba011001970000000002000411000000000012004b00001ec20000613d0000000201000039000000000101041a000009ba01100197000000000012004b00001ec20000613d0000000001000411000000000010043f0000000901000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000101041a000000000001004b00001ec20000c13d000000400100043d000009e902000041000008410000013d00000000010104330000001c0400002900000024034000390000000000230435000009e5020000410000000000240435000009ba0110019700001a6d0000013d000000e00200043d000000400300043d000009eb010000410000000000130435001d00000003001d0000000401300039000000000300041000000000003104350000000001000414000009ba02200197000000040020008c00001ed30000c13d0000000103000031000000200030008c0000002004000039000000000403401900001efd0000013d0000001d03000029000009b60030009c000009b6030080410000004003300210000009b60010009c000009b601008041000000c001100210000000000131019f000009ec011001c726d326ce0000040f0000006003100270000009b603300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000001d0570002900001eec0000613d000000000801034f0000001d09000029000000008a08043c0000000009a90436000000000059004b00001ee80000c13d000000000006004b00001ef90000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000001f1e0000613d0000001f01400039000000600210018f0000001d01200029000000000021004b00000000020000390000000102004039000009bb0010009c000000460000213d0000000100200190000000460000c13d000000400010043f000000200030008c000000410000413d0000001d0200002900000000030204330000001c0230006a000009c304200197000009c305300197000000000654013f000000000054004b0000000004000019000009c304002041000000000032004b0000000003000019000009c30300a041000009c30060009c000000000403c019000000000004004b000004ea0000613d000009c30020009c00001f2a0000413d00000a0202000041000008410000013d0000001f0530018f000009b806300198000000400200043d000000000462001900000a1b0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001f250000c13d00000a1b0000013d0000000801000039000000000101041a001100000001001d000000000001004b00001f310000c13d001b001c0000002d00001ff70000013d001a00000000001d001b001c0000002d0000000801000039000000000101041a0000001a0010006c000021300000a13d0000001a01000029000009ed0110009a000000000101041a001d00000001001d000000000010043f0000000a01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000201041a0000001c012000b90000001c031000fa000000000032004b000004ea0000c13d00000012011000fa001509bc0010019b0000001b01000029000009bc01100197000000150110006a001b00000001001d000009bc0010009c000004ea0000213d0000001d01000029000009ba05100197000000400200043d0000004401200039000000e00300043d000000150400002900000000004104350000002001200039000009ee0400004100000000004104350000002404200039001000000005001d000000000054043500000044040000390000000000420435000009ef0020009c000000460000213d000009ba05300197000000c003200039000000400030043f000000a004200039000009f003000041001d00000004001d000000000034043500000080042000390000002003000039000f00000004001d000000000034043500000000030204330000000002000414001400000005001d000000040050008c00001f790000c13d0000000101000031000000010200003900001f8b0000013d000009b60010009c000009b6010080410000004001100210000009b60030009c000009b6030080410000006003300210000000000113019f000009b60020009c000009b602008041000000c002200210000000000121019f000000140200002926d326c90000040f000000010220018f00030000000103550000006001100270000109b60010019d000009b601100197000000000001004b001900800000003d001800600000003d00001fb80000613d0000001f0310003900000a7c033001970000003f0330003900000a7c03300197000000400400043d0000000003340019001800000004001d000000000043004b00000000040000390000000104004039000009bb0030009c000000460000213d0000000100400190000000460000c13d000000400030043f0000001803000029000000000513043600000a7c04100198001900000005001d0000000003450019000000030500036700001fab0000613d000000000605034f0000001907000029000000006806043c0000000007870436000000000037004b00001fa70000c13d0000001f0110019000001fb80000613d000000000445034f0000000301100210000000000503043300000000051501cf000000000515022f000000000404043b0000010001100089000000000414022f00000000011401cf000000000151019f000000000013043500000018010000290000000001010433000000000002004b0000213c0000613d000000000001004b00001fd20000c13d000009f1010000410000000000100443000000140100002900000004001004430000000001000414000009b60010009c000009b601008041000000c001100210000009f2011001c7000080020200003926d326ce0000040f00000001002001900000212f0000613d000000000101043b000000000001004b000021600000613d00000018010000290000000001010433000000000001004b00001fdf0000613d000009c10010009c000000410000213d000000200010008c000000410000413d00000019010000290000000001010433000000000001004b0000000002000039000000010200c039000000000021004b000000410000c13d000000000001004b000021710000613d000000400100043d00000015020000290000000000210435000009b60010009c000009b60100804100000040011002100000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f000009f7011001c70000800d020000390000000203000039000009f804000041000000100500002926d326c90000040f0000000100200190000000410000613d0000001a020000290000000102200039001a00000002001d000000110020006c00001f330000413d0000000e02000039000000000102041a000009fb011001970000001b011001af000000000012041b0000000801000039000000000101041a000000000001004b000020730000613d001b0001001000920000000802000039000000000202041a0000001b0020006c000021300000a13d000009fc0110009a000000000101041a001d00000001001d000000000010043f0000000a01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d0000001d01000029000009ba01100197001c00000001001d000000000010043f0000000a01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000001041b0000001c01000029000000000010043f0000000901000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000101041a001d00000001001d000000000001004b000020700000613d0000000801000039000000000201041a000000000002004b000004ea0000613d0000001d03000029000000010130008a000000000023004b0000205c0000613d000000000012004b000021300000a13d0000001d01000029000009fc0110009a000009fc0220009a000000000202041a000000000021041b000000000020043f0000000901000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b0000001d02000029000000000021041b0000000801000039000000000101041a001d00000001001d000000000001004b000021360000613d0000001d01000029000000010110008a0000001d02000029000009fc0220009a000000000002041b0000000802000039000000000012041b0000001c01000029000000000010043f0000000901000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000001041b0000001b01000029000000000001004b000020000000c13d000000160000006b001c00000000001d000020d20000613d001b00000000001d001c00000000001d000000170100002900000000010104330000001b0010006c000021300000a13d0000001b010000290000000501100210000000130110002900000000010104330000000012010434001d09ba0020019c000021850000613d000000e00200043d000009ba022001970000001d0020006b000021850000613d0000000001010433001a00000001001d0000001d01000029000000000010043f0000000a01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d0000001a020000290000ffff0220018f000000000101043b001a00000002001d000000000021041b0000001d01000029000000000010043f0000000901000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b000000000101041a000000000001004b000020c70000c13d0000000801000039000000000101041a000009bb0010009c000000460000213d00000001021000390000000803000039000000000023041b000009ed0110009a0000001d02000029000000000021041b000000000103041a001900000001001d000000000020043f0000000901000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000000410000613d000000000101043b0000001902000029000000000021041b0000001c01000029000009b6011001970000001a01100029001c00000001001d000009b60010009c000004ea0000213d0000001b020000290000000102200039001b00000002001d000000160020006c000020780000413d0000000e03000039000000000103041a000009fe011001970000001c040000290000006002400210000009ff02200197000000000112019f000000000013041b000000400100043d000000200210003900000040030000390000000000320435000000000041043500000017020000290000000003020433000000400210003900000000003204350000006002100039000000000003004b000020f30000613d0000002f04000029000000000500001900000000460404340000000076060434000009ba06600197000000000662043600000000070704330000ffff0770018f000000000076043500000040022000390000000105500039000000000035004b000020e80000413d0000000002120049000009b60020009c000009b6020080410000006002200210000009b60010009c000009b6010080410000004001100210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000121019f000009cd011001c70000800d02000039000000010300003900000a000400004126d326c90000040f0000000100200190000000410000613d000000800100043d00000140000004430000016000100443000000a00100043d00000020030000390000018000300443000001a000100443000000c00100043d0000004002000039000001c000200443000001e000100443000000e00100043d000000600200003900000200002004430000022000100443000001000100043d000000800200003900000240002004430000026000100443000000a001000039000001200200043d0000028000100443000002a000200443000000c001000039000001400200043d000002c000100443000002e000200443000000e001000039000001600200043d000003000010044300000320002004430000010001000039000001800200043d0000034000100443000003600020044300000100003004430000000901000039000001200010044300000a0101000041000026d40001042e000000000001042f00000a6201000041000000000010043f0000003201000039000000040010043f000009ec01000041000026d50001043000000a6201000041000000000010043f0000003101000039000000040010043f000009ec01000041000026d500010430000000000001004b000021880000c13d000000400100043d000009f50200004100000000002104350000000402100039000000200300003900000000003204350000000f020000290000000002020433000000240310003900000000002304350000004403100039000000000002004b000021530000613d000000000400001900000000053400190000001d06400029000000000606043300000000006504350000002004400039000000000024004b0000214c0000413d0000001f0420003900000a7c04400197000000000232001900000000000204350000004402400039000009b60020009c000009b6020080410000006002200210000009b60010009c000009b6010080410000004001100210000000000112019f000026d500010430000000400100043d0000004402100039000009f903000041000000000032043500000024021000390000001d030000390000000000320435000009f5020000410000000000210435000000040210003900000020030000390000000000320435000009b60010009c000009b6010080410000004001100210000009fa011001c7000026d500010430000000400100043d0000006402100039000009f30300004100000000003204350000004402100039000009f403000041000000000032043500000024021000390000002a030000390000000000320435000009f5020000410000000000210435000000040210003900000020030000390000000000320435000009b60010009c000009b6010080410000004001100210000009f6011001c7000026d500010430000000400100043d000009fd0200004100000cb70000013d0000001902000029000009b60020009c000009b6020080410000004002200210000009b60010009c000009b6010080410000006001100210000000000121019f000026d5000104300000000043010434000009ba0330019700000000033204360000000004040433000009bb04400197000000000043043500000040031000390000000003030433000009bb033001970000004004200039000000000034043500000060031000390000000003030433000009bb033001970000006004200039000000000034043500000080031000390000000003030433000009bc0330019700000080042000390000000000340435000000a0031000390000000003030433000009ba03300197000000a0042000390000000000340435000000c0031000390000000003030433000009ba03300197000000c0042000390000000000340435000000e002200039000000e0011000390000000001010433000009ba011001970000000000120435000000000001042d00000a7d0010009c000021bb0000813d000000e001100039000000400010043f000000000001042d00000a6201000041000000000010043f0000004101000039000000040010043f000009ec01000041000026d50001043000000a7e0010009c000021c60000813d0000008001100039000000400010043f000000000001042d00000a6201000041000000000010043f0000004101000039000000040010043f000009ec01000041000026d50001043000000000430104340000000001320436000000000003004b000021d80000613d000000000200001900000000051200190000000006240019000000000606043300000000006504350000002002200039000000000032004b000021d10000413d000000000213001900000000000204350000001f0230003900000a7c022001970000000001210019000000000001042d0000000043010434000009ba03300197000000000332043600000000040404330000ffff0440018f000000000043043500000040031000390000000003030433000009b60330019700000040042000390000000000340435000000600310003900000000030304330000ffff0330018f0000006004200039000000000034043500000080031000390000000003030433000009b60330019700000080042000390000000000340435000000a00310003900000000030304330000ffff0330018f000000a0042000390000000000340435000000c00310003900000000030304330000ffff0330018f000000c0042000390000000000340435000000e0031000390000000003030433000009ba03300197000000e004200039000000000034043500000100031000390000000003030433000009b6033001970000010004200039000000000034043500000120031000390000000003030433000009b60330019700000120042000390000000000340435000001400310003900000000030304330000ffff0330018f0000014004200039000000000034043500000160031000390000000003030433000009b60330019700000160042000390000000000340435000001800220003900000180011000390000000001010433000000000001004b0000000001000039000000010100c0390000000000120435000000000001042d000000000301001900000000040104330000000001420436000000000004004b000022300000613d0000000002000019000000200330003900000000050304330000000065050434000009ba05500197000000000551043600000000060604330000ffff0660018f000000000065043500000040011000390000000102200039000000000042004b000022240000413d000000000001042d0001000000000002000000400300043d00000a4c020000410000000000230435000009ba01100197000100000003001d0000000402300039000000000012043500000a3201000041000000000010044300000000010004120000000400100443000001000100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f0000000100200190000022910000613d000000000201043b0000000001000414000009ba02200197000000040020008c000022530000c13d0000000103000031000000200030008c00000020040000390000000004034019000000010b0000290000227e0000013d0000000103000029000009b60030009c000009b6030080410000004003300210000009b60010009c000009b601008041000000c001100210000000000131019f000009ec011001c726d326ce0000040f000000010b0000290000006003100270000009b603300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b00190000226d0000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b000022690000c13d000000000006004b0000227a0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f00030000000103550000000100200190000022980000613d0000001f01400039000000600210018f0000000001b20019000000000021004b00000000020000390000000102004039000009bb0010009c000022920000213d0000000100200190000022920000c13d000000400010043f0000001f0030008c0000228f0000a13d00000000010b0433000009ba0010009c0000228f0000213d000000000001042d0000000001000019000026d500010430000000000001042f00000a6201000041000000000010043f0000004101000039000000040010043f000009ec01000041000026d5000104300000001f0530018f000009b806300198000000400200043d0000000004620019000022a30000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b0000229f0000c13d000000000005004b000022b00000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000009b60020009c000009b6020080410000004002200210000000000112019f000026d5000104300001000000000002000000400200043d000009eb010000410000000000120435000100000002001d00000004012000390000000002000410000000000021043500000a3201000041000000000010044300000000010004120000000400100443000000600100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f0000000100200190000023230000613d000000000201043b0000000001000414000009ba02200197000000040020008c000022d80000c13d0000000103000031000000200030008c00000020040000390000000004034019000000010b000029000023030000013d0000000103000029000009b60030009c000009b6030080410000004003300210000009b60010009c000009b601008041000000c001100210000000000131019f000009ec011001c726d326ce0000040f000000010b0000290000006003100270000009b603300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b0019000022f20000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b000022ee0000c13d000000000006004b000022ff0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f00030000000103550000000100200190000023320000613d0000001f01400039000000600210018f0000000001b20019000000000021004b00000000020000390000000102004039000009bb0010009c000023240000213d0000000100200190000023240000c13d000000400010043f0000001f0030008c0000232a0000a13d00000000020b04330000000e01000039000000000101041a000009bc011001970000000001120049000009c303100197000009c304200197000000000543013f000000000043004b0000000003000019000009c303002041000000000021004b0000000002000019000009c30200a041000009c30050009c000000000302c019000000000003004b0000232c0000613d000000000001042d000000000001042f00000a6201000041000000000010043f0000004101000039000000040010043f000009ec01000041000026d5000104300000000001000019000026d50001043000000a6201000041000000000010043f0000001101000039000000040010043f000009ec01000041000026d5000104300000001f0530018f000009b806300198000000400200043d00000000046200190000233d0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000023390000c13d000000000005004b0000234a0000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000009b60020009c000009b6020080410000004002200210000000000112019f000026d500010430000a000000000002000000000100041a000009ba021001970000000001000411000000000021004b0000236b0000613d0000000202000039000000000202041a000009ba02200197000000000021004b0000236b0000613d000000000010043f0000000901000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f0000000100200190000024ca0000613d000000000101043b000000000101041a000000000001004b0000254b0000613d000000400300043d0000000e01000039000000000101041a0000006002100270000009b602200198000025220000613d000009bc01100197000300000002001d000500000001001d000000000021004b000025240000413d000009eb010000410000000000130435000a00000003001d00000004013000390000000002000410000000000021043500000a3201000041000000000010044300000000010004120000000400100443000000600100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f0000000100200190000025100000613d000000000201043b0000000001000414000009ba02200197000000040020008c000023960000c13d0000000103000031000000200030008c000000200400003900000000040340190000000a0b000029000023c10000013d0000000a03000029000009b60030009c000009b6030080410000004003300210000009b60010009c000009b601008041000000c001100210000000000131019f000009ec011001c726d326ce0000040f0000000a0b0000290000006003100270000009b603300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b0019000023b00000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b000023ac0000c13d000000000006004b000023bd0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f000300000001035500000001002001900000252d0000613d0000001f01400039000000600210018f0000000001b20019000000000021004b00000000020000390000000102004039000009bb0010009c000024cc0000213d0000000100200190000024cc0000c13d000000400010043f000000200030008c000024ca0000413d00000000030b04330000000e02000039000000000202041a000009bc022001970000000002230049000009c304200197000009c305300197000000000654013f000000000054004b0000000004000019000009c304002041000000000032004b0000000003000019000009c30300a041000009c30060009c000000000403c019000000000004004b000024d20000613d000009c30020009c0000252b0000813d0000000801000039000000000101041a000400000001001d00000a3201000041000000000010044300000000010004120000000400100443000000600100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f0000000100200190000025100000613d000000000101043b000000040000006b000024c30000613d000609ba0010019b0000000002000019000a00050000002d0000000801000039000000000101041a000000000021004b000024d80000a13d000800000002001d000009ed0120009a000000000101041a000900000001001d000000000010043f0000000a01000039000000200010043f0000000001000414000009b60010009c000009b601008041000000c001100210000009da011001c7000080100200003926d326ce0000040f00000001002001900000000609000029000024ca0000613d000000000101043b000000000201041a00000005012000b900000005031000fa000000000023004b000024d20000c13d00000003011000fa000009bc0b1001970000000a01000029000009bc011001970000000001b10049000a00000001001d000009bc0010009c000024d20000213d0000000901000029000009ba05100197000000400200043d00000044012000390000000000b104350000002001200039000009ee0300004100000000003104350000002403200039000000000053043500000044030000390000000000320435000009d90020009c000024cc0000213d000000800c2000390000004000c0043f000009ef0020009c000024cc0000213d000000c003200039000000400030043f000000200300003900000000003c0435000000a003200039000009f004000041000000000043043500000000030204330000000002000414000000040090008c0000243d0000c13d00000001010000310000000102000039000024560000013d000009b60010009c000009b6010080410000004001100210000009b60030009c000009b6030080410000006003300210000000000113019f000009b60020009c000009b602008041000000c002200210000000000121019f0000000002090019000900000005001d00070000000b001d00020000000c001d26d326c90000040f000000020c000029000000070b00002900000009050000290000000609000029000000010220018f00030000000103550000006001100270000109b60010019d000009b601100197000000000001004b000000800d000039000000600e000039000024820000613d000009bb0010009c000024cc0000213d0000001f0310003900000a7c033001970000003f0330003900000a7c03300197000000400e00043d00000000033e00190000000000e3004b00000000040000390000000104004039000009bb0030009c000024cc0000213d0000000100400190000024cc0000c13d000000400030043f000000000d1e043600000a7c0410019800000000034d0019000000030a000367000024750000613d00000000060a034f00000000070d0019000000006806043c0000000007870436000000000037004b000024710000c13d0000001f01100190000024820000613d00000000044a034f0000000301100210000000000603043300000000061601cf000000000616022f000000000404043b0000010001100089000000000414022f00000000011401cf000000000161019f000000000013043500000000030e0433000000000002004b000024de0000613d000000000003004b000024a10000c13d00020000000e001d00010000000d001d00070000000b001d000900000005001d000009f101000041000000000010044300000004009004430000000001000414000009b60010009c000009b601008041000000c001100210000009f2011001c7000080020200003926d326ce0000040f0000000100200190000025100000613d000000000101043b000000000001004b0000000201000029000025110000613d0000000003010433000000000003004b0000000905000029000000070b000029000000010d000029000024ad0000613d000009c10030009c000024ca0000213d000000200030008c000024ca0000413d00000000010d0433000000000001004b0000000002000039000000010200c039000000000021004b000024ca0000c13d000000000001004b000024f40000613d000000400100043d0000000000b10435000009b60010009c000009b60100804100000040011002100000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f000009f7011001c70000800d020000390000000203000039000009f80400004126d326c90000040f0000000100200190000024ca0000613d00000008020000290000000102200039000000040020006c000023fa0000413d000024c40000013d000a00050000002d0000000e02000039000000000102041a000009fb011001970000000a011001af000000000012041b000000000001042d0000000001000019000026d50001043000000a6201000041000000000010043f0000004101000039000000040010043f000009ec01000041000026d50001043000000a6201000041000000000010043f0000001101000039000000040010043f000009ec01000041000026d50001043000000a6201000041000000000010043f0000003201000039000000040010043f000009ec01000041000026d500010430000000000003004b000025080000c13d00000000010c0019000000400400043d000a00000004001d000009f5020000410000000000240435000000040340003900000020020000390000000000230435000000240240003926d321cc0000040f0000000a020000290000000001210049000009b60010009c000009b601008041000009b60020009c000009b60200804100000060011002100000004002200210000000000121019f000026d500010430000000400100043d0000006402100039000009f30300004100000000003204350000004402100039000009f403000041000000000032043500000024021000390000002a030000390000000000320435000009f5020000410000000000210435000000040210003900000020030000390000000000320435000009b60010009c000009b6010080410000004001100210000009f6011001c7000026d500010430000009b600d0009c000009b60d0080410000004002d00210000009b60030009c000009b6030080410000006001300210000000000121019f000026d500010430000000000001042f000000400100043d0000004402100039000009f903000041000000000032043500000024021000390000001d030000390000000000320435000009f5020000410000000000210435000000040210003900000020030000390000000000320435000009b60010009c000009b6010080410000004001100210000009fa011001c7000026d50001043000000a8001000041000025250000013d00000a7f010000410000000000130435000009b60030009c000009b6030080410000004001300210000009ea011001c7000026d50001043000000a02020000410000254d0000013d0000001f0530018f000009b806300198000000400200043d0000000004620019000025380000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000025340000c13d000000000005004b000025450000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000009b60020009c000009b6020080410000004002200210000000000112019f000026d500010430000000400100043d000009e9020000410000000000210435000009b60010009c000009b6010080410000004001100210000009ea011001c7000026d5000104300001000000000002000000400300043d00000a810030009c000025aa0000813d0000004004300039000000400040043f000000200430003900000000000404350000000000030435000000000002004b000025770000613d0000000303200210000000200330008900000a820430021f000000040020008c00000a82040080410000000203000367000000000513034f000000000505043b000000000445016f00000a830040009c000025920000613d00000a840040009c000025b30000c13d00000a850220009a00000a860020009c000025b00000413d000000400400043d000009c50040009c000025aa0000213d0000004002400039000000400020043f0000000401100039000000000113034f000000000101043b0000258f0000013d000000400100043d000009c50010009c000025aa0000213d000100000001001d0000004001100039000000400010043f00000a3201000041000000000010044300000000010004120000000400100443000000200100003900000024001004430000000001000414000009b60010009c000009b601008041000000c00110021000000a33011001c7000080050200003926d326ce0000040f0000000100200190000025b20000613d000000000101043b000009bb01100197000000010400002900000000001404350000000001000019000025a60000013d00000a850220009a00000a870020009c000025b00000413d000000400400043d000009c50040009c000025aa0000213d0000004002400039000000400020043f0000000402100039000000000223034f000000000202043b00000000002404350000002401100039000000000113034f000000000101043b000000000001004b0000000002000039000000010200c039000000000021004b000025b00000c13d000000200240003900000000001204350000000001040019000000000001042d00000a6201000041000000000010043f0000004101000039000000040010043f000009ec01000041000026d5000104300000000001000019000026d500010430000000000001042f000000400100043d00000a88020000410000000000210435000009b60010009c000009b6010080410000004001100210000009ea011001c7000026d5000104300002000000000002000000400400043d000000440540003900000000003504350000002003400039000009ee050000410000000000530435000009ba02200197000000240540003900000000002504350000004402000039000000000024043500000a7e0040009c000026490000813d0000008009400039000000400090043f000009ef0040009c000026490000213d000009ba0a100197000000c001400039000000400010043f00000020010000390000000000190435000000a001400039000009f0020000410000000000210435000000000204043300000000010004140000000400a0008c000026070000c13d00000001020000390000000101000031000000000001004b0000261f0000613d000009bb0010009c000026490000213d0000001f0410003900000a7c044001970000003f0440003900000a7c04400197000000400c00043d00000000044c00190000000000c4004b00000000050000390000000105004039000009bb0040009c000026490000213d0000000100500190000026490000c13d000000400040043f000000000b1c043600000a7c031001980000001f0410018f00000000013b00190000000305000367000025f90000613d000000000605034f00000000070b0019000000006806043c0000000007870436000000000017004b000025f50000c13d000000000004004b000026210000613d000000000335034f0000000304400210000000000501043300000000054501cf000000000545022f000000000303043b0000010004400089000000000343022f00000000034301cf000000000353019f0000000000310435000026210000013d000009b60030009c000009b6030080410000004003300210000009b60020009c000009b6020080410000006002200210000000000232019f000009b60010009c000009b601008041000000c001100210000000000112019f00000000020a0019000200000009001d00010000000a001d26d326c90000040f000000010a0000290000000209000029000000010220018f00030000000103550000006001100270000109b60010019d000009b601100197000000000001004b000025dd0000c13d000000600c000039000000800b00003900000000030c0433000000000002004b000026510000613d000000000003004b0000263c0000c13d00020000000c001d00010000000b001d000009f10100004100000000001004430000000400a004430000000001000414000009b60010009c000009b601008041000000c001100210000009f2011001c7000080020200003926d326ce0000040f0000000100200190000026830000613d000000000101043b000000000001004b0000000201000029000026840000613d0000000003010433000000000003004b000000010b000029000026480000613d000009c10030009c0000264f0000213d0000001f0030008c0000264f0000a13d00000000010b0433000000000001004b0000000002000039000000010200c039000000000021004b0000264f0000c13d000000000001004b000026670000613d000000000001042d00000a6201000041000000000010043f0000004101000039000000040010043f000009ec01000041000026d5000104300000000001000019000026d500010430000000000003004b0000267b0000c13d0000000001090019000000400400043d000200000004001d000009f5020000410000000000240435000000040340003900000020020000390000000000230435000000240240003926d321cc0000040f00000002020000290000000001210049000009b60010009c000009b601008041000009b60020009c000009b60200804100000060011002100000004002200210000000000121019f000026d500010430000000400100043d0000006402100039000009f30300004100000000003204350000004402100039000009f403000041000000000032043500000024021000390000002a030000390000000000320435000009f5020000410000000000210435000000040210003900000020030000390000000000320435000009b60010009c000009b6010080410000004001100210000009f6011001c7000026d500010430000009b600b0009c000009b60b0080410000004002b00210000009b60030009c000009b6030080410000006001300210000000000121019f000026d500010430000000000001042f000000400100043d0000004402100039000009f903000041000000000032043500000024021000390000001d030000390000000000320435000009f5020000410000000000210435000000040210003900000020030000390000000000320435000009b60010009c000009b6010080410000004001100210000009fa011001c7000026d500010430000000000001042f000009b60010009c000009b6010080410000004001100210000009b60020009c000009b6020080410000006002200210000000000112019f0000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f000009cd011001c7000080100200003926d326ce0000040f0000000100200190000026a90000613d000000000101043b000000000001042d0000000001000019000026d50001043000000000050100190000000000200443000000050030008c000026b90000413d000000040100003900000000020000190000000506200210000000000664001900000005066002700000000006060031000000000161043a0000000102200039000000000031004b000026b10000413d000009b60030009c000009b60300804100000060013002100000000002000414000009b60020009c000009b602008041000000c002200210000000000112019f00000a89011001c7000000000205001926d326ce0000040f0000000100200190000026c80000613d000000000101043b000000000001042d000000000001042f000026cc002104210000000102000039000000000001042d0000000002000019000000000001042d000026d1002104230000000102000039000000000001042d0000000002000019000000000001042d000026d300000432000026d40001042e000026d5000104300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000000000000000000000000001ffffffe000000000000000000000000000000000000000000000000000000000ffffffe0000000000000000000000000000000000000000000000000fffffffffffffeff000000000000000000000000ffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffffff0000000000000000000000000000000000000000ffffffffffffffffffffffff000000000000000000000000000000000000000000000000fffffffffffffe5f000000000000000000000000000000000000000000000000ffffffffffffff9f00000000000000000000000000000000ffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffff5f8000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffff1f000000000000000000000000000000000000000000000000ffffffffffffffbfffffffffffffffffffffffff0000000000000000000000000000000000000000796b89b91644bc98cd93958e4c9038275d622183e25ac5af08cc6b5d9553913202000002000000000000000000000000000000040000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000ffffffffffffffffffffff0000000000000000000000000000000000000000008acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3020000000000000000000000000000000000000000000000000000000000000000000000000000000000ffff0000000000000000000000000000000000000000000000000000ffffffff0000000000000000000000000000000000000000000000000000ffff00000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000000000000ffff000000000000ffffffff000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000000000000000000000ffffffff0000000000000000000000000000000000000000000000000000000100000000000002000000000000000000000000000000000002a000000000000000000000000045b5ad483aa608464c2c7f278bd413d284d7790cdc836e40652e23a027708220000000000000000000000000000000000000000000000000ffffffffffffff7f02000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000ffffffffffffffff00000000000000000000000000000000ffffffffffffffff000000000000000000000000067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e000000000000000000000000000000000000000000000000ffffffffffffffdf000000000000000000000000000000000000000000000000ffffffff0000000000000000000000000000000000000000000000000000ffff0000000000000000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000024ecdc02000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000044000000000000000000000000f5791bc457b3bb990493cf5f655db46c25ccf5764c9b99b8969b4c72ea7df9d0fb95a0042158e60a33e7b5bec100f3d95407b1a71bee6633bd54b8887449750b195db95800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000070a082310000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000240000000000000000000000000c085601c9b05546c4de925af5cdebeab0dd5f5d4bea4dc57b37e961749c911da9059cbb00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffff3f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65641806aa1896bbf26568e884a7374b41e002500962caba6a15023a8d90e8508b8302000002000000000000000000000000000000240000000000000000000000006f742073756363656564000000000000000000000000000000000000000000005361666545524332303a204552433230206f7065726174696f6e20646964206e08c379a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000084000000000000000000000000020000000000000000000000000000000000002000000000000000000000000055fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000000000000000000000000000000000000000000064000000000000000000000000ffffffffffffffffffffffffffffffffffffffff0000000000000000000000000c085601c9b05546c4de925af5cdebeab0dd5f5d4bea4dc57b37e961749c911e4de938d100000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff00000000000000000000000000000000ffffffff0000000000000000000000008c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd240000000200000000000000000000000000000280000001000000000000000000f4d678b80000000000000000000000000000000000000000000000000000000035be3ac80000000000000000000000000000000000000000000000000000000043616e6e6f7420736574206f776e657220746f207a65726f0000000000000000000000000000000000000000000000000000000000000000000000007437ff9e00000000000000000000000000000000000000000000000000000000c92b283100000000000000000000000000000000000000000000000000000000eff7cc4700000000000000000000000000000000000000000000000000000000f2fde38a00000000000000000000000000000000000000000000000000000000f2fde38b00000000000000000000000000000000000000000000000000000000fbca3b7400000000000000000000000000000000000000000000000000000000eff7cc4800000000000000000000000000000000000000000000000000000000f25561fd00000000000000000000000000000000000000000000000000000000c92b283200000000000000000000000000000000000000000000000000000000d09dc33900000000000000000000000000000000000000000000000000000000df0aa9e900000000000000000000000000000000000000000000000000000000856c8246000000000000000000000000000000000000000000000000000000009a113c35000000000000000000000000000000000000000000000000000000009a113c3600000000000000000000000000000000000000000000000000000000b06d41bc00000000000000000000000000000000000000000000000000000000856c8247000000000000000000000000000000000000000000000000000000008da5cb5b000000000000000000000000000000000000000000000000000000007437ff9f0000000000000000000000000000000000000000000000000000000076f6ae760000000000000000000000000000000000000000000000000000000079ba50970000000000000000000000000000000000000000000000000000000048a98aa300000000000000000000000000000000000000000000000000000000549e946e00000000000000000000000000000000000000000000000000000000599f643000000000000000000000000000000000000000000000000000000000599f643100000000000000000000000000000000000000000000000000000000704b6c0200000000000000000000000000000000000000000000000000000000549e946f0000000000000000000000000000000000000000000000000000000054b714680000000000000000000000000000000000000000000000000000000048a98aa400000000000000000000000000000000000000000000000000000000504bffe000000000000000000000000000000000000000000000000000000000546719cd0000000000000000000000000000000000000000000000000000000020487dec0000000000000000000000000000000000000000000000000000000020487ded000000000000000000000000000000000000000000000000000000004120fccd000000000000000000000000000000000000000000000000000000004816f4f70000000000000000000000000000000000000000000000000000000006285c69000000000000000000000000000000000000000000000000000000001772047e00000000000000000000000000000000000000000000000000000000181f5a779e7177c80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000008000000000000000004f6e6c792063616c6c61626c65206279206f776e6572000000000000000000000000000000000000000000000000000000000064000000800000000000000000ed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127843616e6e6f74207472616e7366657220746f2073656c660000000000000000002cbc26bb000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffff00000000000000000000000000000000310ab089e4439a4c15d089f94afb7896ff553aecb10793d0ab882de59d99a32e0200000200000000000000000000000000000044000000000000000000000000000000000000000000000000000000000000002400000080000000000000000053ad11d8000000000000000000000000000000000000000000000000000000001c0a352900000000000000000000000000000000000000000000000000000000d9a9cd680000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ff00000000000000000000000000000000000000ff000000000000000000000000000000000000d02641a00000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000de0b6b3a76400009a655f7b000000000000000000000000000000000000000000000000000000000000000000000000000000ff0000000000000000000000000000000000000000ffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9725942a00000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffff000000000000000000000000000000001871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a15279c0800000000000000000000000000000000000000000000000000000000f94ebcd1000000000000000000000000000000000000000000000000000000005cf04449000000000000000000000000000000000000000000000000000000000041e5be00000000000000000000000000000000000000000000000000000000856c824700000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000fffffffffffffffffffffffffffffffffffffbff8d666f6000000000000000000000000000000000000000000000000000000000ffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffffbbe4f6db0000000000000000000000000000000000000000000000000000000001ffc9a700000000000000000000000000000000000000000000000000000000aff2afbf000000000000000000000000000000000000000000000000000000009a4575b900000000000000000000000000000000000000000000000000000000000000000000000000000000ff0000000000000000000000000000000000000036f536ca00000000000000000000000000000000000000000000000000000000bf16aab600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000fffffffffffffedfd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd0000000000000000000000000000000000000020000000000000000000000000e5c7a491000000000000000000000000000000000000000000000000000000008693378900000000000000000000000000000000000000000000000000000000a4ec747900000000000000000000000000000000000000000000000000000000f6cd5620000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000e00000000000000000ffffffffffffffffffffff00ffffffff0000000000000000000000000000000002000000000000000000000000000000000000600000000000000000000000009ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000200000008000000000000000004d7573742062652070726f706f736564206f776e6572000000000000000000008be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e04e487b7100000000000000000000000000000000000000000000000000000000b5a10cfa0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001a0000003c0000000000000000002000000000000000000000000000000000000200000008000000000000000008fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329cfbdb8e560000000000000000000000000000000000000000000000000000000002075e0000000000000000000000000000000000000000000000000000000000232cb97f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000640000022000000000000000000000000000000000000000000000000000000004000002200000000000000000ee433e9900000000000000000000000000000000000000000000000000000000a7499d2000000000000000000000000000000000000000000000000000000000ffdb4b3700000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002386f26fc1000000000000000000000000000000000000000000000000000000000000ffffffdf4ab35b0b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000186a0000000000000000000000000000000000000ffffffffffffffffffffffffffff00000000000000000000000000000000000000000000000000005af3107a40004c056b6a000000000000000000000000000000000000000000000000000000004c4fc93a0000000000000000000000000000000000000000000000000000000045564d3245564d4f6e52616d7020312e352e30000000000000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000100000002800000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0000000000000000000000000000000000000000000000000ffffffffffffff20000000000000000000000000000000000000000000000000ffffffffffffff808d0f71d800000000000000000000000000000000000000000000000000000000990e30bf00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffc0ffffffff00000000000000000000000000000000000000000000000000000000181dcf100000000000000000000000000000000000000000000000000000000097a657c9000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000004800000000000000000000000000000000000000000000000000000000000002080000000000000000000000000000000000000000000000000000000000000405247fdce0000000000000000000000000000000000000000000000000000000002000002000000000000000000000000000000000000000000000000000000005e7b9793a663e708e92314ad76f4a3ef0640790771a3c2d7b93ee8e15f133b39")

type GetNops struct {
	NopsAndWeights []EVM2EVMOnRampNopAndWeight
	WeightsTotal   *big.Int
}

func (_EVM2EVMOnRamp *EVM2EVMOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMOnRamp.abi.Events["AdminSet"].ID:
		return _EVM2EVMOnRamp.ParseAdminSet(log)
	case _EVM2EVMOnRamp.abi.Events["CCIPSendRequested"].ID:
		return _EVM2EVMOnRamp.ParseCCIPSendRequested(log)
	case _EVM2EVMOnRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMOnRamp.ParseConfigChanged(log)
	case _EVM2EVMOnRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMOnRamp.ParseConfigSet(log)
	case _EVM2EVMOnRamp.abi.Events["FeeConfigSet"].ID:
		return _EVM2EVMOnRamp.ParseFeeConfigSet(log)
	case _EVM2EVMOnRamp.abi.Events["NopPaid"].ID:
		return _EVM2EVMOnRamp.ParseNopPaid(log)
	case _EVM2EVMOnRamp.abi.Events["NopsSet"].ID:
		return _EVM2EVMOnRamp.ParseNopsSet(log)
	case _EVM2EVMOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMOnRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMOnRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMOnRamp.abi.Events["TokenTransferFeeConfigDeleted"].ID:
		return _EVM2EVMOnRamp.ParseTokenTransferFeeConfigDeleted(log)
	case _EVM2EVMOnRamp.abi.Events["TokenTransferFeeConfigSet"].ID:
		return _EVM2EVMOnRamp.ParseTokenTransferFeeConfigSet(log)
	case _EVM2EVMOnRamp.abi.Events["TokensConsumed"].ID:
		return _EVM2EVMOnRamp.ParseTokensConsumed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMOnRampAdminSet) Topic() common.Hash {
	return common.HexToHash("0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c")
}

func (EVM2EVMOnRampCCIPSendRequested) Topic() common.Hash {
	return common.HexToHash("0xd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd")
}

func (EVM2EVMOnRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (EVM2EVMOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x45b5ad483aa608464c2c7f278bd413d284d7790cdc836e40652e23a027708220")
}

func (EVM2EVMOnRampFeeConfigSet) Topic() common.Hash {
	return common.HexToHash("0x067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e")
}

func (EVM2EVMOnRampNopPaid) Topic() common.Hash {
	return common.HexToHash("0x55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f")
}

func (EVM2EVMOnRampNopsSet) Topic() common.Hash {
	return common.HexToHash("0x8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd24")
}

func (EVM2EVMOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMOnRampTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0xfb95a0042158e60a33e7b5bec100f3d95407b1a71bee6633bd54b8887449750b")
}

func (EVM2EVMOnRampTokenTransferFeeConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf5791bc457b3bb990493cf5f655db46c25ccf5764c9b99b8969b4c72ea7df9d0")
}

func (EVM2EVMOnRampTokensConsumed) Topic() common.Hash {
	return common.HexToHash("0x1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a")
}

func (_EVM2EVMOnRamp *EVM2EVMOnRamp) Address() common.Address {
	return _EVM2EVMOnRamp.address
}

type CustomTransaction struct {
	*types.Transaction
	CustomHash common.Hash
}

func (tx *CustomTransaction) Hash() common.Hash {
	return tx.CustomHash
}

func ConvertToTransaction(resp zktypes.TransactionResponse) *CustomTransaction {
	dtx := &types.DynamicFeeTx{
		ChainID:   resp.ChainID.ToInt(),
		Nonce:     uint64(resp.Nonce),
		GasTipCap: resp.MaxPriorityFeePerGas.ToInt(),
		GasFeeCap: resp.MaxFeePerGas.ToInt(),
		To:        &resp.To,
		Value:     resp.Value.ToInt(),
		Data:      resp.Data,
		Gas:       uint64(resp.Gas),
	}

	tx := types.NewTx(dtx)
	customTransaction := CustomTransaction{Transaction: tx, CustomHash: resp.Hash}
	return &customTransaction
}

func DeployZkSyncEVM2EVMOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, params ...interface{}) (common.Address, *CustomTransaction, *EVM2EVMOnRamp, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	fmt.Println("Deploying zksync contract")
	zksyncClient := zkSyncClient.NewClient(client.Client())
	fmt.Println("getting wallet")
	wallet := auth.Context.Value("wallet").(*zkSyncAccounts.Wallet)
	fmt.Println("got wallet")
	fmt.Println("getting bytes")
	decodedBytes := common.FromHex(EVM2EVMOnRampZkBin)
	fmt.Println("deploying")
	EVM2EVMOnRampAbi, err := EVM2EVMOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	constructor, err := EVM2EVMOnRampAbi.Pack("", params...)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	hash, err := wallet.DeployWithCreate(nil, zkSyncAccounts.CreateTransaction{
		Bytecode: decodedBytes,
		Calldata: constructor,
	})
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	fmt.Println("hash of tx", hash)
	receipt, err := zksyncClient.WaitMined(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	tx, _, err := zksyncClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	fmt.Println("tx hash", tx.Hash)

	ethTx := ConvertToTransaction(*tx)
	address := receipt.ContractAddress

	parsed, err := EVM2EVMOnRampMetaData.GetAbi()
	contractBind := bind.NewBoundContract(address, *parsed, backend, backend, backend)

	contractReturn := &EVM2EVMOnRamp{address: address, abi: *parsed, EVM2EVMOnRampCaller: EVM2EVMOnRampCaller{contract: contractBind}, EVM2EVMOnRampTransactor: EVM2EVMOnRampTransactor{contract: contractBind}, EVM2EVMOnRampFilterer: EVM2EVMOnRampFilterer{contract: contractBind}}

	return address, ethTx, contractReturn, err
}

type EVM2EVMOnRampInterface interface {
	CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error)

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMOnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetFeeTokenConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMOnRampFeeTokenConfig, error)

	GetNopFeesJuels(opts *bind.CallOpts) (*big.Int, error)

	GetNops(opts *bind.CallOpts) (GetNops,

		error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMOnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMOnRampTokenTransferFeeConfig, error)

	LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	PayNops(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMOnRampDynamicConfig) (*types.Transaction, error)

	SetFeeTokenConfig(opts *bind.TransactOpts, feeTokenConfigArgs []EVM2EVMOnRampFeeTokenConfigArgs) (*types.Transaction, error)

	SetNops(opts *bind.TransactOpts, nopsAndWeights []EVM2EVMOnRampNopAndWeight) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	SetTokenTransferFeeConfig(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []EVM2EVMOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawNonLinkFees(opts *bind.TransactOpts, feeToken common.Address, to common.Address) (*types.Transaction, error)

	FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMOnRampAdminSetIterator, error)

	WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampAdminSet) (event.Subscription, error)

	ParseAdminSet(log types.Log) (*EVM2EVMOnRampAdminSet, error)

	FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMOnRampCCIPSendRequestedIterator, error)

	WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampCCIPSendRequested) (event.Subscription, error)

	ParseCCIPSendRequested(log types.Log) (*EVM2EVMOnRampCCIPSendRequested, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOnRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMOnRampConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMOnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMOnRampConfigSet, error)

	FilterFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMOnRampFeeConfigSetIterator, error)

	WatchFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampFeeConfigSet) (event.Subscription, error)

	ParseFeeConfigSet(log types.Log) (*EVM2EVMOnRampFeeConfigSet, error)

	FilterNopPaid(opts *bind.FilterOpts, nop []common.Address) (*EVM2EVMOnRampNopPaidIterator, error)

	WatchNopPaid(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampNopPaid, nop []common.Address) (event.Subscription, error)

	ParseNopPaid(log types.Log) (*EVM2EVMOnRampNopPaid, error)

	FilterNopsSet(opts *bind.FilterOpts) (*EVM2EVMOnRampNopsSetIterator, error)

	WatchNopsSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampNopsSet) (event.Subscription, error)

	ParseNopsSet(log types.Log) (*EVM2EVMOnRampNopsSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMOnRampOwnershipTransferred, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts) (*EVM2EVMOnRampTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampTokenTransferFeeConfigDeleted) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*EVM2EVMOnRampTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMOnRampTokenTransferFeeConfigSetIterator, error)

	WatchTokenTransferFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampTokenTransferFeeConfigSet) (event.Subscription, error)

	ParseTokenTransferFeeConfigSet(log types.Log) (*EVM2EVMOnRampTokenTransferFeeConfigSet, error)

	FilterTokensConsumed(opts *bind.FilterOpts) (*EVM2EVMOnRampTokensConsumedIterator, error)

	WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampTokensConsumed) (event.Subscription, error)

	ParseTokensConsumed(log types.Log) (*EVM2EVMOnRampTokensConsumed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
