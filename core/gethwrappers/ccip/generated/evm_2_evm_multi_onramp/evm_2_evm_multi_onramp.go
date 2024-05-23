// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_multi_onramp

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
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

type EVM2EVMMultiOnRampDestChainConfig struct {
	DynamicConfig  EVM2EVMMultiOnRampDestChainDynamicConfig
	PrevOnRamp     common.Address
	SequenceNumber uint64
	MetadataHash   [32]byte
}

type EVM2EVMMultiOnRampDestChainConfigArgs struct {
	DestChainSelector uint64
	DynamicConfig     EVM2EVMMultiOnRampDestChainDynamicConfig
	PrevOnRamp        common.Address
}

type EVM2EVMMultiOnRampDestChainDynamicConfig struct {
	IsEnabled                         bool
	MaxNumberOfTokensPerMsg           uint16
	MaxDataBytes                      uint32
	MaxPerMsgGasLimit                 uint32
	DestGasOverhead                   uint32
	DestGasPerPayloadByte             uint16
	DestDataAvailabilityOverheadGas   uint32
	DestGasPerDataAvailabilityByte    uint16
	DestDataAvailabilityMultiplierBps uint16
	DefaultTokenFeeUSDCents           uint16
	DefaultTokenDestGasOverhead       uint32
	DefaultTokenDestBytesOverhead     uint32
	DefaultTxGasLimit                 uint64
}

type EVM2EVMMultiOnRampDynamicConfig struct {
	Router             common.Address
	PriceRegistry      common.Address
	TokenAdminRegistry common.Address
	FeeAggregator      common.Address
}

type EVM2EVMMultiOnRampFeeTokenConfig struct {
	NetworkFeeUSDCents         uint32
	GasMultiplierWeiPerEth     uint64
	PremiumMultiplierWeiPerEth uint64
	Enabled                    bool
}

type EVM2EVMMultiOnRampFeeTokenConfigArgs struct {
	Token                      common.Address
	NetworkFeeUSDCents         uint32
	GasMultiplierWeiPerEth     uint64
	PremiumMultiplierWeiPerEth uint64
	Enabled                    bool
}

type EVM2EVMMultiOnRampStaticConfig struct {
	LinkToken       common.Address
	ChainSelector   uint64
	MaxNopFeesJuels *big.Int
	RmnProxy        common.Address
}

type EVM2EVMMultiOnRampTokenTransferFeeConfig struct {
	MinFeeUSDCents            uint32
	MaxFeeUSDCents            uint32
	DeciBps                   uint16
	DestGasOverhead           uint32
	DestBytesOverhead         uint32
	AggregateRateLimitEnabled bool
	IsEnabled                 bool
}

type EVM2EVMMultiOnRampTokenTransferFeeConfigArgs struct {
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

var EVM2EVMMultiOnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSendZeroTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"DestinationChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChainSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"}],\"name\":\"InvalidDestBytesOverhead\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidDestChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkBalanceNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeBalanceReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeesToPay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NotAFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SourceTokenDataTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainDynamicConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeConfig\",\"type\":\"tuple[]\"}],\"name\":\"FeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NopsPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokenTransferFeeConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"transferFeeConfig\",\"type\":\"tuple[]\"}],\"name\":\"TokenTransferFeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensConsumed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyDestChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestChainConfig\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeeTokenConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.FeeTokenConfig\",\"name\":\"feeTokenConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNopFeesJuels\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenTransferFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"address[]\"}],\"name\":\"setTokenTransferFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162007b2438038062007b24833981016040819052620000359162001a31565b8233806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c08162000223565b50506040805160a081018252602084810180516001600160801b039081168085524263ffffffff169385018490528751151585870181905292518216606086018190529790950151166080909301839052600380546001600160a01b031916909417600160801b9283021760ff60a01b1916600160a01b90910217909255029091176004555085516001600160a01b0316158062000169575060208601516001600160401b0316155b8062000180575060608601516001600160a01b0316155b156200019f576040516306b7c75960e31b815260040160405180910390fd5b85516001600160a01b0390811660a05260208701516001600160401b031660c05260408701516001600160601b031660805260608701511660e052620001e585620002ce565b620001f08462000472565b620001fb826200095c565b6040805160008152602081019091526200021790829062000b83565b5050505050506200202f565b336001600160a01b038216036200027d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60208101516001600160a01b03161580620002f4575060408101516001600160a01b0316155b806200030b575060608101516001600160a01b0316155b156200032a576040516306b7c75960e31b815260040160405180910390fd5b8051600580546001600160a01b039283166001600160a01b031991821617909155602080840151600680549185169184169190911790556040808501516007805491861691851691909117905560608086015160088054918716919095161790935580516080808201835260a0518616825260c0516001600160401b03169382019390935291516001600160601b03168282015260e0519093169181019190915290517fa0d39dfd39d411e1d09ac113d962b2f31903b670d3d79ec23dc1a60b55238012916200046791849082516001600160a01b0390811682526020808501516001600160401b0316818401526040808601516001600160601b03168185015260609586015183168685015284518316608085015290840151821660a0840152830151811660c083015291909201511660e08201526101000190565b60405180910390a150565b60005b81518110156200095857600082828151811062000496576200049662001b64565b602002602001015190506000838381518110620004b757620004b762001b64565b6020026020010151600001519050806001600160401b0316600003620004fc5760405163c35aa79d60e01b81526001600160401b038216600482015260240162000084565b600060096000836001600160401b03166001600160401b0316815260200190815260200160002090506000836040015190506000604051806080016040528086602001518152602001836001600160a01b031681526020018460020160149054906101000a90046001600160401b03166001600160401b031681526020018460030154815250905080600001518360000160008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a8154816001600160401b0302191690836001600160401b0316021790555090505082600301546000801b03620008755760c051604080517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b360208201526001600160401b0392831691810191909152908516606082015230608082015260a00160408051601f1981840301815291905280516020909101206060820181905260038401556001600160a01b038216156200082c576002830180546001600160a01b0319166001600160a01b0384161790555b836001600160401b03167fb6031b30c90eb9236da8c32ffd47ca6307f4a0eaa17e78e12f20a24afd5805108460405162000867919062001b7a565b60405180910390a262000947565b60028301546001600160a01b03838116911614620008b25760405163c35aa79d60e01b81526001600160401b038516600482015260240162000084565b60208560200151610160015163ffffffff161015620008ff57602085015161016001516040516312766e0160e11b81526000600482015263ffffffff909116602482015260440162000084565b836001600160401b03167f8a71e142c5eeae8754825ba46ee7cd42468c1621a90b72dcad8bf7ee8c7111c386602001516040516200093e919062001cbe565b60405180910390a25b505050505080600101905062000475565b5050565b60005b815181101562000b5157600082828151811062000980576200098062001b64565b60209081029190910181015160408051608080820183528385015163ffffffff9081168352838501516001600160401b0390811684880190815260608088015183168688019081529488018051151591870191825288516001600160a01b03166000908152600a909a5296909820945185549151945198511515600160a01b0260ff60a01b199984166c010000000000000000000000000299909916600160601b600160a81b031995909316640100000000026001600160601b031990921693169290921791909117919091161793909317909255905190915015801562000a735750805162000a7390600d9062000e22565b1562000b1057805162000a8990600d9062000e49565b5080516040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa15801562000ad3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000af9919062001dd3565b111562000b0a5762000b0a62000e60565b62000b47565b8060800151801562000b2f5750805162000b2d90600d9062000e22565b155b1562000b4757805162000b4590600d9062000f91565b505b506001016200095f565b507f067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e8160405162000467919062001ded565b60005b825181101562000d4257600083828151811062000ba75762000ba762001b64565b6020026020010151905060208160a0015163ffffffff16101562000bfd57805160a08201516040516312766e0160e11b81526001600160a01b03909216600483015263ffffffff16602482015260440162000084565b6040805160e08101825260208381015163ffffffff908116835284840151811682840190815260608087015161ffff9081168688019081526080808a0151861693880193845260a0808b0151871691890191825260c0808c01511515918a019182526001908a018181529b516001600160a01b03166000908152600b9099529990972097518854955192519451915197519a519087166001600160401b031990961695909517640100000000928716929092029190911765ffffffffffff60401b191668010000000000000000939092169290920263ffffffff60501b1916176a0100000000000000000000918416919091021764ffffffffff60701b1916600160701b939092169290920260ff60901b191617600160901b941515949094029390931760ff60981b1916600160981b93151593909302929092179091550162000b86565b507ff5791bc457b3bb990493cf5f655db46c25ccf5764c9b99b8969b4c72ea7df9d08260405162000d74919062001e7c565b60405180910390a160005b815181101562000ddc57600b600083838151811062000da25762000da262001b64565b6020908102919091018101516001600160a01b0316825281019190915260400160002080546001600160a01b031916905560010162000d7f565b5080511562000958577ffb95a0042158e60a33e7b5bec100f3d95407b1a71bee6633bd54b8887449750b8160405162000e16919062001f11565b60405180910390a15050565b6001600160a01b038116600090815260018301602052604081205415155b90505b92915050565b600062000e40836001600160a01b03841662000fa8565b6008546001600160a01b0316600062000e7a600d620010ac565b600f80546001600160601b031981169091559091506001600160601b031660005b8281101562000f4457600062000eb3600d83620010b7565b6040516370a0823160e01b815230600482015290915062000f3a9086906001600160a01b038416906370a0823190602401602060405180830381865afa15801562000f02573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000f28919062001dd3565b6001600160a01b0384169190620010c5565b5060010162000e9b565b506040516001600160601b03821681526001600160a01b038416907f88db9a873cb1a9b6d3757fc0fc19ede741b35af2a9a497a5793e760cdd97966b9060200160405180910390a2505050565b600062000e40836001600160a01b03841662001122565b60008181526001830160205260408120548015620010a157600062000fcf60018362001f60565b855490915060009062000fe59060019062001f60565b90508181146200105157600086600001828154811062001009576200100962001b64565b90600052602060002001549050808760000184815481106200102f576200102f62001b64565b6000918252602080832090910192909255918252600188019052604090208390555b855486908062001065576200106562001f82565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000e43565b600091505062000e43565b600062000e43825490565b600062000e40838362001174565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b0390811663a9059cbb60e01b179091526200111d918591620011a116565b505050565b60008181526001830160205260408120546200116b5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000e43565b50600062000e43565b60008260000182815481106200118e576200118e62001b64565b9060005260206000200154905092915050565b6040805180820190915260208082527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656490820152600090620011f0906001600160a01b03851690849062001272565b8051909150156200111d578080602001905181019062001211919062001f98565b6200111d5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840162000084565b60606200128384846000856200128b565b949350505050565b606082471015620012ee5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840162000084565b600080866001600160a01b031685876040516200130c919062001fdc565b60006040518083038185875af1925050503d80600081146200134b576040519150601f19603f3d011682016040523d82523d6000602084013e62001350565b606091505b50909250905062001364878383876200136f565b979650505050505050565b60608315620013e3578251600003620013db576001600160a01b0385163b620013db5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640162000084565b508162001283565b620012838383815115620013fa5781518083602001fd5b8060405162461bcd60e51b815260040162000084919062001ffa565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b038111828210171562001451576200145162001416565b60405290565b604051606081016001600160401b038111828210171562001451576200145162001416565b6040516101a081016001600160401b038111828210171562001451576200145162001416565b60405160a081016001600160401b038111828210171562001451576200145162001416565b60405160e081016001600160401b038111828210171562001451576200145162001416565b604051601f8201601f191681016001600160401b038111828210171562001517576200151762001416565b604052919050565b80516001600160a01b03811681146200153757600080fd5b919050565b80516001600160401b03811681146200153757600080fd5b6000608082840312156200156757600080fd5b620015716200142c565b90506200157e826200151f565b81526200158e602083016200151f565b6020820152620015a1604083016200151f565b6040820152620015b4606083016200151f565b606082015292915050565b60006001600160401b03821115620015db57620015db62001416565b5060051b60200190565b805180151581146200153757600080fd5b805161ffff811681146200153757600080fd5b805163ffffffff811681146200153757600080fd5b600082601f8301126200163057600080fd5b81516020620016496200164383620015bf565b620014ec565b8281526101e092830285018201928282019190878511156200166a57600080fd5b8387015b85811015620017ea5780890382811215620016895760008081fd5b6200169362001457565b6200169e836200153c565b81526101a080601f1984011215620016b65760008081fd5b620016c06200147c565b9250620016cf888501620015e5565b83526040620016e0818601620015f6565b898501526060620016f381870162001609565b8286015260806200170681880162001609565b8287015260a091506200171b82880162001609565b9086015260c06200172e878201620015f6565b8287015260e091506200174382880162001609565b9086015261010062001757878201620015f6565b8287015261012091506200176d828801620015f6565b9086015261014062001781878201620015f6565b8287015261016091506200179782880162001609565b90860152610180620017ab87820162001609565b82870152620017bc8488016200153c565b8187015250508389840152620017d66101c086016200151f565b90830152508552509284019281016200166e565b5090979650505050505050565b80516001600160801b03811681146200153757600080fd5b6000606082840312156200182257600080fd5b6200182c62001457565b90506200183982620015e5565b81526200184960208301620017f7565b60208201526200185c60408301620017f7565b604082015292915050565b600082601f8301126200187957600080fd5b815160206200188c6200164383620015bf565b82815260a09283028501820192828201919087851115620018ac57600080fd5b8387015b85811015620017ea5781818a031215620018ca5760008081fd5b620018d4620014a2565b620018df826200151f565b8152620018ee86830162001609565b868201526040620019018184016200153c565b908201526060620019148382016200153c565b90820152608062001927838201620015e5565b908201528452928401928101620018b0565b600082601f8301126200194b57600080fd5b815160206200195e6200164383620015bf565b82815260e092830285018201928282019190878511156200197e57600080fd5b8387015b85811015620017ea5781818a0312156200199c5760008081fd5b620019a6620014c7565b620019b1826200151f565b8152620019c086830162001609565b868201526040620019d381840162001609565b908201526060620019e6838201620015f6565b908201526080620019f983820162001609565b9082015260a062001a0c83820162001609565b9082015260c062001a1f838201620015e5565b90820152845292840192810162001982565b6000806000806000808688036101c081121562001a4d57600080fd5b608081121562001a5c57600080fd5b5062001a676200142c565b62001a72886200151f565b815262001a82602089016200153c565b602082015260408801516001600160601b038116811462001aa257600080fd5b604082015262001ab5606089016200151f565b6060820152955062001acb886080890162001554565b6101008801519095506001600160401b038082111562001aea57600080fd5b62001af88a838b016200161e565b955062001b0a8a6101208b016200180f565b945061018089015191508082111562001b2257600080fd5b62001b308a838b0162001867565b93506101a089015191508082111562001b4857600080fd5b5062001b5789828a0162001939565b9150509295509295509295565b634e487b7160e01b600052603260045260246000fd5b815460ff81161515825261020082019061ffff600882901c8116602085015263ffffffff601883901c8116604086015262001bc260608601828560381c1663ffffffff169052565b62001bda60808601828560581c1663ffffffff169052565b62001bf060a08601838560781c1661ffff169052565b62001c0860c08601828560881c1663ffffffff169052565b62001c1e60e08601838560a81c1661ffff169052565b62001c356101008601838560b81c1661ffff169052565b62001c4c6101208601838560c81c1661ffff169052565b62001c656101408601828560d81c1663ffffffff169052565b600186015490811663ffffffff1661016086015260201c6001600160401b0390811661018086015260028601546001600160a01b0381166101a087015260a01c166101c085015250506003909201546101e09091015290565b8151151581526101a08101602083015162001cdf602084018261ffff169052565b50604083015162001cf8604084018263ffffffff169052565b50606083015162001d11606084018263ffffffff169052565b50608083015162001d2a608084018263ffffffff169052565b5060a083015162001d4160a084018261ffff169052565b5060c083015162001d5a60c084018263ffffffff169052565b5060e083015162001d7160e084018261ffff169052565b506101008381015161ffff9081169184019190915261012080850151909116908301526101408084015163ffffffff908116918401919091526101608085015190911690830152610180928301516001600160401b0316929091019190915290565b60006020828403121562001de657600080fd5b5051919050565b602080825282518282018190526000919060409081850190868401855b8281101562001e6f57815180516001600160a01b031685528681015163ffffffff1687860152858101516001600160401b03908116878701526060808301519091169086015260809081015115159085015260a0909301929085019060010162001e0a565b5091979650505050505050565b602080825282518282018190526000919060409081850190868401855b8281101562001e6f57815180516001600160a01b031685528681015163ffffffff908116888701528682015181168787015260608083015161ffff169087015260808083015182169087015260a0808301519091169086015260c09081015115159085015260e0909301929085019060010162001e99565b6020808252825182820181905260009190848201906040850190845b8181101562001f545783516001600160a01b03168352928401929184019160010162001f2d565b50909695505050505050565b8181038181111562000e4357634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60006020828403121562001fab57600080fd5b62000e4082620015e5565b60005b8381101562001fd357818101518382015260200162001fb9565b50506000910152565b6000825162001ff081846020870162001fb6565b9190910192915050565b60208152600082518060208401526200201b81604085016020870162001fb6565b601f01601f19169190910160400192915050565b60805160a05160c05160e051615a75620020af6000396000818161027a01528181612a160152612cb40152600081816102160152818161265c015281816129b1015261328d0152600081816101e70152818161298c01528181612ff801526130f1015260008181610246015281816129e301526131bd0152615a756000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c806379ba5097116100ee578063c92b283211610097578063eff7cc4811610071578063eff7cc48146109fa578063f25561fd14610a02578063f2fde38b14610a15578063fbca3b7414610a2857600080fd5b8063c92b2832146109bf578063cdc73d51146109d2578063df0aa9e9146109e757600080fd5b80639041be3d116100c85780639041be3d146108a55780639a113c36146108b8578063a6f3ab6c146109ac57600080fd5b806379ba5097146108605780638b364334146108685780638da5cb5b1461089457600080fd5b8063546719cd1161015b5780636def4ce7116101355780636def4ce71461057f578063704b6c02146107c757806370edbc2d146107da5780637437ff9f146107ed57600080fd5b8063546719cd146104ea57806354b714681461054e578063599f64311461056e57600080fd5b806320487ded1161018c57806320487ded146104895780634816f4f7146104aa57806348a98aa4146104bf57600080fd5b806306285c69146101b35780631772047e146102c0578063181f5a7714610440575b600080fd5b6102aa60408051608081018252600080825260208201819052918101829052606081019190915260405180608001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b6040516102b791906144db565b60405180910390f35b6103d46102ce366004614550565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152506001600160a01b03166000908152600b6020908152604091829020825160e081018452905463ffffffff8082168352640100000000820481169383019390935261ffff68010000000000000000820416938201939093526a01000000000000000000008304821660608201526e0100000000000000000000000000008304909116608082015260ff720100000000000000000000000000000000000083048116151560a0830152730100000000000000000000000000000000000000909204909116151560c082015290565b6040516102b79190600060e08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015260c0830151151560c083015292915050565b61047c6040518060400160405280601c81526020017f45564d3245564d4d756c74694f6e52616d7020312e362e302d6465760000000081525081565b6040516102b791906145bd565b61049c610497366004614609565b610a3b565b6040519081526020016102b7565b6104bd6104b836600461483a565b610ebf565b005b6104d26104cd366004614977565b610ed5565b6040516001600160a01b0390911681526020016102b7565b6104f2610f64565b6040516102b7919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b600f546040516bffffffffffffffffffffffff90911681526020016102b7565b6002546001600160a01b03166104d2565b6107ba61058d3660046149b0565b604080516102208101825260006080820181815260a0830182905260c0830182905260e08301829052610100830182905261012083018290526101408301829052610160830182905261018083018290526101a083018290526101c083018290526101e08301829052610200830182905282526020820181905291810182905260608101919091525067ffffffffffffffff908116600090815260096020908152604091829020825161022081018452815460ff811615156080830190815261ffff610100808404821660a086015263ffffffff63010000008504811660c08701526701000000000000008504811660e08701526b01000000000000000000000085048116918601919091526f01000000000000000000000000000000840482166101208601527101000000000000000000000000000000000084048116610140860152750100000000000000000000000000000000000000000084048216610160860152770100000000000000000000000000000000000000000000008404821661018086015279010000000000000000000000000000000000000000000000000084049091166101a08501527b0100000000000000000000000000000000000000000000000000000090920482166101c084015260018401549182166101e08401526401000000009091048616610200830152815260028201546001600160a01b03811693820193909352600160a01b90920490931691810191909152600390910154606082015290565b6040516102b79190614ad1565b6104bd6107d5366004614550565b610ffb565b6104bd6107e8366004614b1e565b6110ba565b61085360408051608081018252600080825260208201819052918101829052606081019190915250604080516080810182526005546001600160a01b03908116825260065481166020830152600754811692820192909252600854909116606082015290565b6040516102b79190614cf2565b6104bd6110ce565b61087b610876366004614977565b61118c565b60405167ffffffffffffffff90911681526020016102b7565b6000546001600160a01b03166104d2565b61087b6108b33660046149b0565b611276565b6109626108c6366004614550565b604080516080810182526000808252602082018190529181018290526060810191909152506001600160a01b03166000908152600a60209081526040918290208251608081018452905463ffffffff8116825267ffffffffffffffff64010000000082048116938301939093526c0100000000000000000000000081049092169281019290925260ff600160a01b909104161515606082015290565b60408051825163ffffffff16815260208084015167ffffffffffffffff908116918301919091528383015116918101919091526060918201511515918101919091526080016102b7565b6104bd6109ba366004614d2e565b6112a9565b6104bd6109cd366004614dd3565b6112ba565b6109da611322565b6040516102b79190614e23565b61049c6109f5366004614e70565b61132e565b6104bd611751565b6104bd610a10366004614edc565b611763565b6104bd610a23366004614550565b611774565b6109da610a363660046149b0565b611785565b67ffffffffffffffff82166000908152600960205260408120805460ff16610aa0576040517f99ac52f200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526024015b60405180910390fd5b6000610aaf6080850185614fcf565b159050610ad057610acb610ac66080860186614fcf565b6117b9565b610ae8565b6001820154640100000000900467ffffffffffffffff165b9050610b1285610afb6020870187614fcf565b905083610b0b604089018961501d565b9050611861565b6000600a81610b276080880160608901614550565b6001600160a01b0316815260208082019290925260409081016000208151608081018352905463ffffffff81168252640100000000810467ffffffffffffffff908116948301949094526c01000000000000000000000000810490931691810191909152600160a01b90910460ff16151560608201819052909150610bf457610bb66080860160608701614550565b6040517fa7499d200000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a97565b60065460009081906001600160a01b031663ffdb4b37610c1a60808a0160608b01614550565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015267ffffffffffffffff8b1660248201526044016040805180830381865afa158015610c85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca99190615093565b90925090506000808080610cc060408c018c61501d565b90501115610cfb57610cef8b610cdc60808d0160608e01614550565b87610cea60408f018f61501d565b611971565b91945092509050610d17565b8551610d149063ffffffff16662386f26fc100006150dc565b92505b67ffffffffffffffff8b166000908152600960205260408120805477010000000000000000000000000000000000000000000000900461ffff1615610dad57610daa8d607060ff16887bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16901c8e8060200190610d909190614fcf565b90508f8060400190610da2919061501d565b905087611d72565b91505b6000886020015167ffffffffffffffff168563ffffffff1683600001600f9054906101000a900461ffff1661ffff168f8060200190610dec9190614fcf565b610df79291506150dc565b8454610e18906b010000000000000000000000900463ffffffff168e6150f3565b610e2291906150f3565b610e2c91906150f3565b610e46906dffffffffffffffffffffffffffff8a166150dc565b610e5091906150dc565b9050877bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1683828b6040015167ffffffffffffffff1689610e8d91906150dc565b610e9791906150f3565b610ea191906150f3565b610eab9190615106565b9b5050505050505050505050505b92915050565b610ec7611e78565b610ed18282611ed5565b5050565b6007546040517fbbe4f6db0000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152600092169063bbe4f6db90602401602060405180830381865afa158015610f39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5d9190615128565b9392505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526003546fffffffffffffffffffffffffffffffff8082168352600160801b80830463ffffffff166020850152600160a01b90920460ff161515938301939093526004548084166060840152049091166080820152610ff690612246565b905090565b6000546001600160a01b0316331480159061102157506002546001600160a01b03163314155b15611058576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c906020015b60405180910390a150565b6110c2611e78565b6110cb816122f8565b50565b6001546001600160a01b031633146111285760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610a97565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6001600160a01b0381166000908152600c602052604081205467ffffffffffffffff16808203610f5d5767ffffffffffffffff84166000908152600960205260409020600201546001600160a01b0316801561126e576040517f856c82470000000000000000000000000000000000000000000000000000000081526001600160a01b03858116600483015282169063856c824790602401602060405180830381865afa158015611241573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112659190615145565b92505050610eb9565b509392505050565b67ffffffffffffffff8082166000908152600960205260408120600201549091610eb991600160a01b9004166001615162565b6112b161284a565b6110cb816128a4565b6000546001600160a01b031633148015906112e057506002546001600160a01b03163314155b15611317576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110cb600382612a6e565b6060610ff6600d612c03565b67ffffffffffffffff84166000908152600960205260408120816113558288888888612c10565b905060005b816101400151518110156116f3576000611377604089018961501d565b8381811061138757611387615183565b90506040020180360381019061139d9190615199565b905060006113af8a8360000151610ed5565b90506001600160a01b038116158061146557506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa15801561143f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061146391906151d3565b155b156114aa5781516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a97565b6000816001600160a01b0316637a5c972d60405180608001604052808d80600001906114d69190614fcf565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525067ffffffffffffffff8f166020808301919091526001600160a01b038d1660408084019190915290880151606090920191909152517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815261157891906004016151f0565b6000604051808303816000875af1158015611597573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115bf91908101906152ad565b9050602081602001515111801561160e575082516001600160a01b03166000908152600b602090815260409091205490820151516e01000000000000000000000000000090910463ffffffff16105b156116535782516040517f36f536ca0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a97565b805161165e906134c2565b5060408051606081019091526001600160a01b03831660808201528060a0810160405160208183030381529060405281526020018260000151815260200182602001518152506040516020016116b4919061533e565b60405160208183030381529060405285610160015185815181106116da576116da615183565b602002602001018190525050505080600101905061135a565b5061170281836003015461351d565b6101808201526040517fd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd90611738908390615440565b60405180910390a161018001519150505b949350505050565b611759611e78565b611761613678565b565b61176b611e78565b6110cb816137d9565b61177c61284a565b6110cb81613a48565b60606040517f9e7177c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f97a657c9000000000000000000000000000000000000000000000000000000006117e68385615575565b7fffffffff00000000000000000000000000000000000000000000000000000000161461183f576040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61184c82600481866155bd565b81019061185991906155e7565b519392505050565b67ffffffffffffffff8416600090815260096020526040902080546301000000900463ffffffff168411156118da5780546040517f86933789000000000000000000000000000000000000000000000000000000008152630100000090910463ffffffff16600482015260248101859052604401610a97565b8054670100000000000000900463ffffffff16831115611926576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8054610100900461ffff1682111561196a576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b6000808083815b81811015611d6557600087878381811061199457611994615183565b9050604002018036038101906119aa9190615199565b905060006001600160a01b03166119c58c8360000151610ed5565b6001600160a01b031603611a135780516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a97565b80516001600160a01b03166000908152600b6020908152604091829020825160e081018452905463ffffffff8082168352640100000000820481169383019390935261ffff68010000000000000000820416938201939093526a01000000000000000000008304821660608201526e0100000000000000000000000000008304909116608082015260ff720100000000000000000000000000000000000083048116151560a0830152730100000000000000000000000000000000000000909204909116151560c08201819052611b915767ffffffffffffffff8c1660009081526009602052604090208054611b3190790100000000000000000000000000000000000000000000000000900461ffff16662386f26fc100006150dc565b611b3b90896150f3565b8154909850611b6f907b01000000000000000000000000000000000000000000000000000000900463ffffffff1688615629565b6001820154909750611b879063ffffffff1687615629565b9550505050611d5d565b604081015160009061ffff1615611cad5760008c6001600160a01b031684600001516001600160a01b031614611c505760065484516040517f4ab35b0b0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911690634ab35b0b90602401602060405180830381865afa158015611c25573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c499190615646565b9050611c53565b508a5b620186a0836040015161ffff16611c958660200151847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16613afe90919063ffffffff16565b611c9f91906150dc565b611ca99190615106565b9150505b6060820151611cbc9088615629565b9650816080015186611cce9190615629565b8251909650600090611ced9063ffffffff16662386f26fc100006150dc565b905080821015611d0c57611d01818a6150f3565b985050505050611d5d565b6000836020015163ffffffff16662386f26fc10000611d2b91906150dc565b905080831115611d4b57611d3f818b6150f3565b99505050505050611d5d565b611d55838b6150f3565b995050505050505b600101611978565b5050955095509592505050565b60008063ffffffff8316611d876080866150dc565b611d93876102206150f3565b611d9d91906150f3565b611da791906150f3565b67ffffffffffffffff8816600090815260096020526040812080549293509171010000000000000000000000000000000000810463ffffffff1690611e09907501000000000000000000000000000000000000000000900461ffff16856150dc565b611e1391906150f3565b825490915077010000000000000000000000000000000000000000000000900461ffff16611e516dffffffffffffffffffffffffffff8a16836150dc565b611e5b91906150dc565b611e6b90655af3107a40006150dc565b9998505050505050505050565b6000546001600160a01b03163314801590611e9e57506002546001600160a01b03163314155b15611761576040517ffbdb8e5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8251811015612163576000838281518110611ef557611ef5615183565b6020026020010151905060208160a0015163ffffffff161015611f6257805160a08201516040517f24ecdc020000000000000000000000000000000000000000000000000000000081526001600160a01b03909216600483015263ffffffff166024820152604401610a97565b6040805160e08101825260208381015163ffffffff908116835284840151811682840190815260608087015161ffff9081168688019081526080808a0151861693880193845260a0808b0151871691890191825260c0808c01511515918a019182526001908a018181529b516001600160a01b03166000908152600b9099529990972097518854955192519451915197519a519087167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009096169590951764010000000092871692909202919091177fffffffffffffffffffffffffffffffffffff000000000000ffffffffffffffff166801000000000000000093909216929092027fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff16176a010000000000000000000091841691909102177fffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffffff166e01000000000000000000000000000093909216929092027fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff1617720100000000000000000000000000000000000094151594909402939093177fffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffff16730100000000000000000000000000000000000000931515939093029290921790915501611ed8565b507ff5791bc457b3bb990493cf5f655db46c25ccf5764c9b99b8969b4c72ea7df9d0826040516121939190615661565b60405180910390a160005b815181101561220357600b60008383815181106121bd576121bd615183565b6020908102919091018101516001600160a01b03168252810191909152604001600020805473ffffffffffffffffffffffffffffffffffffffff1916905560010161219e565b50805115610ed1577ffb95a0042158e60a33e7b5bec100f3d95407b1a71bee6633bd54b8887449750b8160405161223a9190614e23565b60405180910390a15050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526122d482606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff16426122b89190615701565b85608001516fffffffffffffffffffffffffffffffff16613b3b565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b60005b8151811015610ed157600082828151811061231857612318615183565b60200260200101519050600083838151811061233657612336615183565b60200260200101516000015190508067ffffffffffffffff16600003612394576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610a97565b6000600960008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002090506000836040015190506000604051806080016040528086602001518152602001836001600160a01b031681526020018460020160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020018460030154815250905080600001518360000160008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555090505082600301546000801b0361273a57604080517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3602082015267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811692820192909252908516606082015230608082015260a00160408051601f1981840301815291905280516020909101206060820181905260038401556001600160a01b038216156126f35760028301805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0384161790555b8367ffffffffffffffff167fb6031b30c90eb9236da8c32ffd47ca6307f4a0eaa17e78e12f20a24afd5805108460405161272d9190615714565b60405180910390a261283a565b60028301546001600160a01b0383811691161461278f576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610a97565b60208560200151610160015163ffffffff1610156127f357602085015161016001516040517f24ecdc020000000000000000000000000000000000000000000000000000000081526000600482015263ffffffff9091166024820152604401610a97565b8367ffffffffffffffff167f8a71e142c5eeae8754825ba46ee7cd42468c1621a90b72dcad8bf7ee8c7111c386602001516040516128319190615851565b60405180910390a25b50505050508060010190506122fb565b6000546001600160a01b031633146117615760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610a97565b60208101516001600160a01b031615806128c9575060408101516001600160a01b0316155b806128df575060608101516001600160a01b0316155b15612916576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516005805473ffffffffffffffffffffffffffffffffffffffff199081166001600160a01b0393841617909155602080840151600680548416918516919091179055604080850151600780548516918616919091179055606080860151600880549095169086161790935580516080810182527f0000000000000000000000000000000000000000000000000000000000000000851681527f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16928101929092527f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff16828201527f00000000000000000000000000000000000000000000000000000000000000009093169181019190915290517fa0d39dfd39d411e1d09ac113d962b2f31903b670d3d79ec23dc1a60b55238012916110af918490615860565b8154600090612a8a90600160801b900463ffffffff1642615701565b90508015612b075760018301548354612ac5916fffffffffffffffffffffffffffffffff808216928116918591600160801b90910416613b3b565b83546fffffffffffffffffffffffffffffffff9190911673ffffffffffffffffffffffffffffffffffffffff1990911617600160801b4263ffffffff16021783555b60208201518354612b2d916fffffffffffffffffffffffffffffffff9081169116613b63565b835483511515600160a01b027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff9283161717845560208301516040808501518316600160801b0291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990612bf69084908151151581526020808301516fffffffffffffffffffffffffffffffff90811691830191909152604092830151169181019190915260600190565b60405180910390a1505050565b60606000610f5d83613b79565b604080516101a08101825260008082526020820181905281830181905260608083018290526080830182905260a0830182905260c0830182905260e0830182905261010083018290526101208301819052610140830181905261016083015261018082015290517f58babe3300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff861660048201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906358babe3390602401602060405180830381865afa158015612cfb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d1f91906151d3565b15612d62576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610a97565b6001600160a01b038216612da2576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005546001600160a01b03163314612de6576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855460ff16612e2d576040517f99ac52f200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610a97565b6000612e3c6080860186614fcf565b159050612e5857612e53610ac66080870187614fcf565b612e70565b6001870154640100000000900467ffffffffffffffff165b90506000612e81604087018761501d565b9150612e9f905087612e966020890189614fcf565b90508484611861565b8015612fee576000805b82811015612fdc57612ebe604089018961501d565b82818110612ece57612ece615183565b90506040020160200135600003612f11576040517f5cf0444900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b6000612f2260408b018b61501d565b84818110612f3257612f32615183565b612f489260206040909202019081019150614550565b6001600160a01b031681526020810191909152604001600020547201000000000000000000000000000000000000900460ff1615612fd457612fc7612f9060408a018a61501d565b83818110612fa057612fa0615183565b905060400201803603810190612fb69190615199565b6006546001600160a01b0316613bd5565b612fd190836150f3565b91505b600101612ea9565b508015612fec57612fec81613cf6565b505b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166130286080880160608901614550565b6001600160a01b03160361308c57600f80548691906000906130599084906bffffffffffffffffffffffff166158eb565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055506131ab565b6006546001600160a01b03166241e5be6130ac6080890160608a01614550565b60405160e083901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b039182166004820152602481018990527f00000000000000000000000000000000000000000000000000000000000000009091166044820152606401602060405180830381865afa158015613138573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061315c9190615910565b600f805460009061317c9084906bffffffffffffffffffffffff166158eb565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055505b600f546bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811691161115613218576040517fe5c7a49100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000613224888661118c565b61322f906001615162565b6001600160a01b0386166000818152600c602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169190911790915582516101a0810184527f00000000000000000000000000000000000000000000000000000000000000009091168152908101929092529192509081016133056132cb8a80614fcf565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506134c292505050565b6001600160a01b031681526020018a600201601481819054906101000a900467ffffffffffffffff1661333790615929565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff1681526020018481526020016000151581526020018267ffffffffffffffff16815260200188606001602081019061339d9190614550565b6001600160a01b031681526020018781526020018880602001906133c19190614fcf565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200161340860408a018a61501d565b808060200260200160405190810160405280939291908181526020016000905b828210156134545761344560408302860136819003810190615199565b81526020019060010190613428565b505050505081526020018367ffffffffffffffff81111561347757613477614659565b6040519080825280602002602001820160405280156134aa57816020015b60608152602001906001900390816134955790505b50815260006020909101529998505050505050505050565b6000815160201461350157816040517f8d666f60000000000000000000000000000000000000000000000000000000008152600401610a9791906145bd565b610eb9828060200190518101906135189190615910565b613d03565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b61010001516040516020016135b39897969594939291906001600160a01b039889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b60405160208183030381529060405280519060200120856101200151805190602001208661014001516040516020016135ec9190615950565b604051602081830303815290604052805190602001208761016001516040516020016136189190615963565b60408051601f198184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b6008546001600160a01b03166000613690600d613d6f565b600f80547fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169091559091506bffffffffffffffffffffffff1660005b828110156137875760006136e3600d83613d79565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290915061377e9086906001600160a01b038416906370a0823190602401602060405180830381865afa158015613749573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061376d9190615910565b6001600160a01b0384169190613d85565b506001016136ce565b506040516bffffffffffffffffffffffff821681526001600160a01b038416907f88db9a873cb1a9b6d3757fc0fc19ede741b35af2a9a497a5793e760cdd97966b9060200160405180910390a2505050565b60005b8151811015613a185760008282815181106137f9576137f9615183565b60209081029190910181015160408051608080820183528385015163ffffffff90811683528385015167ffffffffffffffff90811684880190815260608088015183168688019081529488018051151591870191825288516001600160a01b03166000908152600a909a5296909820945185549151945198511515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9984166c0100000000000000000000000002999099167fffffffffffffffffffffff000000000000000000ffffffffffffffffffffffff95909316640100000000027fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009092169316929092179190911791909116179390931790925590519091501580156139325750805161393290600d90613e0a565b156139de57805161394590600d90613e2c565b5080516040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa1580156139a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139cb9190615910565b11156139d9576139d9613678565b613a0f565b806080015180156139fa575080516139f890600d90613e0a565b155b15613a0f578051613a0d90600d90613e41565b505b506001016137dc565b507f067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e816040516110af9190615976565b336001600160a01b03821603613aa05760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610a97565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000670de0b6b3a7640000613b31837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff86166150dc565b610f5d9190615106565b6000613b5a85613b4b84866150dc565b613b5590876150f3565b613b63565b95945050505050565b6000818310613b725781610f5d565b5090919050565b606081600001805480602002602001604051908101604052809291908181526020018280548015613bc957602002820191906000526020600020905b815481526020019060010190808311613bb5575b50505050509050919050565b81516040517fd02641a00000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009182919084169063d02641a0906024016040805180830381865afa158015613c3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c5f91906159f7565b5190507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116600003613cc85783516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a97565b6020840151611749907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690613afe565b6110cb6003826000613e56565b60006001600160a01b03821180613d1a5750600a82105b15613d6b5760408051602081018490520160408051601f19818403018152908290527f8d666f60000000000000000000000000000000000000000000000000000000008252610a97916004016145bd565b5090565b6000610eb9825490565b6000610f5d8383614160565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052613e0590849061418a565b505050565b6001600160a01b03811660009081526001830160205260408120541515610f5d565b6000610f5d836001600160a01b03841661426f565b6000610f5d836001600160a01b038416614369565b8254600160a01b900460ff161580613e6c575081155b15613e7657505050565b825460018401546fffffffffffffffffffffffffffffffff80831692911690600090613eaf90600160801b900463ffffffff1642615701565b90508015613f555781831115613ef1576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001860154613f1e90839085908490600160801b90046fffffffffffffffffffffffffffffffff16613b3b565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16600160801b4263ffffffff160217875592505b84821015613ff2576001600160a01b038416613fa7576040517ff94ebcd10000000000000000000000000000000000000000000000000000000081526004810183905260248101869052604401610a97565b6040517f1a76572a00000000000000000000000000000000000000000000000000000000815260048101839052602481018690526001600160a01b0385166044820152606401610a97565b848310156140de57600186810154600160801b90046fffffffffffffffffffffffffffffffff169060009082906140299082615701565b614033878a615701565b61403d91906150f3565b6140479190615106565b90506001600160a01b038616614093576040517f15279c080000000000000000000000000000000000000000000000000000000081526004810182905260248101869052604401610a97565b6040517fd0c8d23a00000000000000000000000000000000000000000000000000000000815260048101829052602481018690526001600160a01b0387166044820152606401610a97565b6140e88584615701565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b600082600001828154811061417757614177615183565b9060005260206000200154905092915050565b60006141df826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166143b89092919063ffffffff16565b805190915015613e0557808060200190518101906141fd91906151d3565b613e055760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610a97565b60008181526001830160205260408120548015614358576000614293600183615701565b85549091506000906142a790600190615701565b905081811461430c5760008660000182815481106142c7576142c7615183565b90600052602060002001549050808760000184815481106142ea576142ea615183565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061431d5761431d615a36565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610eb9565b6000915050610eb9565b5092915050565b60008181526001830160205260408120546143b057508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610eb9565b506000610eb9565b6060611749848460008585600080866001600160a01b031685876040516143df9190615a4c565b60006040518083038185875af1925050503d806000811461441c576040519150601f19603f3d011682016040523d82523d6000602084013e614421565b606091505b50915091506144328783838761443d565b979650505050505050565b606083156144ac5782516000036144a5576001600160a01b0385163b6144a55760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610a97565b5081611749565b61174983838151156144c15781518083602001fd5b8060405162461bcd60e51b8152600401610a9791906145bd565b60808101610eb9828480516001600160a01b03908116835260208083015167ffffffffffffffff16908401526040808301516bffffffffffffffffffffffff169084015260609182015116910152565b6001600160a01b03811681146110cb57600080fd5b803561454b8161452b565b919050565b60006020828403121561456257600080fd5b8135610f5d8161452b565b60005b83811015614588578181015183820152602001614570565b50506000910152565b600081518084526145a981602086016020860161456d565b601f01601f19169290920160200192915050565b602081526000610f5d6020830184614591565b67ffffffffffffffff811681146110cb57600080fd5b803561454b816145d0565b600060a0828403121561460357600080fd5b50919050565b6000806040838503121561461c57600080fd5b8235614627816145d0565b9150602083013567ffffffffffffffff81111561464357600080fd5b61464f858286016145f1565b9150509250929050565b634e487b7160e01b600052604160045260246000fd5b60405160e0810167ffffffffffffffff8111828210171561469257614692614659565b60405290565b6040516060810167ffffffffffffffff8111828210171561469257614692614659565b6040516101a0810167ffffffffffffffff8111828210171561469257614692614659565b60405160a0810167ffffffffffffffff8111828210171561469257614692614659565b6040805190810167ffffffffffffffff8111828210171561469257614692614659565b604051601f8201601f1916810167ffffffffffffffff8111828210171561474e5761474e614659565b604052919050565b600067ffffffffffffffff82111561477057614770614659565b5060051b60200190565b63ffffffff811681146110cb57600080fd5b803561454b8161477a565b803561ffff8116811461454b57600080fd5b80151581146110cb57600080fd5b803561454b816147a9565b600082601f8301126147d357600080fd5b813560206147e86147e383614756565b614725565b8083825260208201915060208460051b87010193508684111561480a57600080fd5b602086015b8481101561482f5780356148228161452b565b835291830191830161480f565b509695505050505050565b600080604080848603121561484e57600080fd5b833567ffffffffffffffff8082111561486657600080fd5b818601915086601f83011261487a57600080fd5b8135602061488a6147e383614756565b82815260e0928302850182019282820191908b8511156148a957600080fd5b958301955b848710156149525780878d0312156148c65760008081fd5b6148ce61466f565b87356148d98161452b565b8152878501356148e88161477a565b81860152878901356148f98161477a565b818a0152606061490a898201614797565b9082015260808881013561491d8161477a565b9082015260a061492e89820161478c565b9082015260c061493f8982016147b7565b90820152835295860195918301916148ae565b509750508701359350508083111561496957600080fd5b505061464f858286016147c2565b6000806040838503121561498a57600080fd5b8235614995816145d0565b915060208301356149a58161452b565b809150509250929050565b6000602082840312156149c257600080fd5b8135610f5d816145d0565b80511515825260208101516149e8602084018261ffff169052565b506040810151614a00604084018263ffffffff169052565b506060810151614a18606084018263ffffffff169052565b506080810151614a30608084018263ffffffff169052565b5060a0810151614a4660a084018261ffff169052565b5060c0810151614a5e60c084018263ffffffff169052565b5060e0810151614a7460e084018261ffff169052565b506101008181015161ffff9081169184019190915261012080830151909116908301526101408082015163ffffffff9081169184019190915261016080830151909116908301526101809081015167ffffffffffffffff16910152565b600061020082019050614ae58284516149cd565b60208301516001600160a01b03166101a0830152604083015167ffffffffffffffff166101c08301526060909201516101e09091015290565b60006020808385031215614b3157600080fd5b823567ffffffffffffffff811115614b4857600080fd5b8301601f81018513614b5957600080fd5b8035614b676147e382614756565b8181526101e09182028301840191848201919088841115614b8757600080fd5b938501935b83851015614ce65784890381811215614ba55760008081fd5b614bad614698565b8635614bb8816145d0565b81526101a0601f198301811315614bcf5760008081fd5b614bd76146bb565b9250614be48989016147b7565b83526040614bf3818a01614797565b8a8501526060614c04818b0161478c565b828601526080614c15818c0161478c565b8287015260a09150614c28828c0161478c565b9086015260c0614c398b8201614797565b8287015260e09150614c4c828c0161478c565b90860152610100614c5e8b8201614797565b828701526101209150614c72828c01614797565b90860152610140614c848b8201614797565b828701526101609150614c98828c0161478c565b90860152610180614caa8b820161478c565b82870152614cb9848c016145e6565b818701525050838a840152614cd16101c08a01614540565b90830152508452509384019391850191614b8c565b50979650505050505050565b60808101610eb9828480516001600160a01b03908116835260208083015182169084015260408083015182169084015260609182015116910152565b600060808284031215614d4057600080fd5b6040516080810181811067ffffffffffffffff82111715614d6357614d63614659565b6040528235614d718161452b565b81526020830135614d818161452b565b60208201526040830135614d948161452b565b60408201526060830135614da78161452b565b60608201529392505050565b80356fffffffffffffffffffffffffffffffff8116811461454b57600080fd5b600060608284031215614de557600080fd5b614ded614698565b8235614df8816147a9565b8152614e0660208401614db3565b6020820152614e1760408401614db3565b60408201529392505050565b6020808252825182820181905260009190848201906040850190845b81811015614e645783516001600160a01b031683529284019291840191600101614e3f565b50909695505050505050565b60008060008060808587031215614e8657600080fd5b8435614e91816145d0565b9350602085013567ffffffffffffffff811115614ead57600080fd5b614eb9878288016145f1565b935050604085013591506060850135614ed18161452b565b939692955090935050565b60006020808385031215614eef57600080fd5b823567ffffffffffffffff811115614f0657600080fd5b8301601f81018513614f1757600080fd5b8035614f256147e382614756565b81815260a09182028301840191848201919088841115614f4457600080fd5b938501935b83851015614ce65780858a031215614f615760008081fd5b614f696146df565b8535614f748161452b565b815285870135614f838161477a565b81880152604086810135614f96816145d0565b90820152606086810135614fa9816145d0565b90820152608086810135614fbc816147a9565b9082015283529384019391850191614f49565b6000808335601e19843603018112614fe657600080fd5b83018035915067ffffffffffffffff82111561500157600080fd5b60200191503681900382131561501657600080fd5b9250929050565b6000808335601e1984360301811261503457600080fd5b83018035915067ffffffffffffffff82111561504f57600080fd5b6020019150600681901b360382131561501657600080fd5b80517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116811461454b57600080fd5b600080604083850312156150a657600080fd5b6150af83615067565b91506150bd60208401615067565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b8082028115828204841417610eb957610eb96150c6565b80820180821115610eb957610eb96150c6565b60008261512357634e487b7160e01b600052601260045260246000fd5b500490565b60006020828403121561513a57600080fd5b8151610f5d8161452b565b60006020828403121561515757600080fd5b8151610f5d816145d0565b67ffffffffffffffff818116838216019080821115614362576143626150c6565b634e487b7160e01b600052603260045260246000fd5b6000604082840312156151ab57600080fd5b6151b3614702565b82356151be8161452b565b81526020928301359281019290925250919050565b6000602082840312156151e557600080fd5b8151610f5d816147a9565b60208152600082516080602084015261520c60a0840182614591565b905067ffffffffffffffff60208501511660408401526001600160a01b036040850151166060840152606084015160808401528091505092915050565b600082601f83011261525a57600080fd5b815167ffffffffffffffff81111561527457615274614659565b6152876020601f19601f84011601614725565b81815284602083860101111561529c57600080fd5b61174982602083016020870161456d565b6000602082840312156152bf57600080fd5b815167ffffffffffffffff808211156152d757600080fd5b90830190604082860312156152eb57600080fd5b6152f3614702565b82518281111561530257600080fd5b61530e87828601615249565b82525060208301518281111561532357600080fd5b61532f87828601615249565b60208301525095945050505050565b60208152600082516060602084015261535a6080840182614591565b90506020840151601f19808584030160408601526153788383614591565b9250604086015191508085840301606086015250613b5a8282614591565b60008151808452602080850194506020840160005b838110156153db57815180516001600160a01b0316885283015183880152604090960195908201906001016153ab565b509495945050505050565b60008282518085526020808601955060208260051b8401016020860160005b8481101561543357601f19868403018952615421838351614591565b98840198925090830190600101615405565b5090979650505050505050565b6020815261545b60208201835167ffffffffffffffff169052565b6000602083015161547760408401826001600160a01b03169052565b5060408301516001600160a01b038116606084015250606083015167ffffffffffffffff8116608084015250608083015160a083015260a08301516154c060c084018215159052565b5060c083015167ffffffffffffffff811660e08401525060e08301516101006154f3818501836001600160a01b03169052565b840151610120848101919091528401516101a0610140808601829052919250906155216101c0860184614591565b9250808601519050601f196101608187860301818801526155428584615396565b94508088015192505061018081878603018188015261556185846153e6565b970151959092019490945250929392505050565b7fffffffff0000000000000000000000000000000000000000000000000000000081358181169160048510156155b55780818660040360031b1b83161692505b505092915050565b600080858511156155cd57600080fd5b838611156155da57600080fd5b5050820193919092039150565b6000602082840312156155f957600080fd5b6040516020810181811067ffffffffffffffff8211171561561c5761561c614659565b6040529135825250919050565b63ffffffff818116838216019080821115614362576143626150c6565b60006020828403121561565857600080fd5b610f5d82615067565b602080825282518282018190526000919060409081850190868401855b828110156156f457815180516001600160a01b031685528681015163ffffffff908116888701528682015181168787015260608083015161ffff169087015260808083015182169087015260a0808301519091169086015260c09081015115159085015260e0909301929085019060010161567e565b5091979650505050505050565b81810381811115610eb957610eb96150c6565b815460ff81161515825261020082019061ffff600882901c8116602085015263ffffffff601883901c8116604086015261575b60608601828560381c1663ffffffff169052565b61577260808601828560581c1663ffffffff169052565b61578760a08601838560781c1661ffff169052565b61579e60c08601828560881c1663ffffffff169052565b6157b360e08601838560a81c1661ffff169052565b6157c96101008601838560b81c1661ffff169052565b6157df6101208601838560c81c1661ffff169052565b6157f76101408601828560d81c1663ffffffff169052565b600186015490811663ffffffff1661016086015260201c67ffffffffffffffff90811661018086015260028601546001600160a01b0381166101a087015260a01c166101c085015250506003909201546101e09091015290565b6101a08101610eb982846149cd565b61010081016158b1828580516001600160a01b03908116835260208083015167ffffffffffffffff16908401526040808301516bffffffffffffffffffffffff169084015260609182015116910152565b82516001600160a01b0390811660808401526020840151811660a08401526040840151811660c084015260608401511660e0830152610f5d565b6bffffffffffffffffffffffff818116838216019080821115614362576143626150c6565b60006020828403121561592257600080fd5b5051919050565b600067ffffffffffffffff808316818103615946576159466150c6565b6001019392505050565b602081526000610f5d6020830184615396565b602081526000610f5d60208301846153e6565b602080825282518282018190526000919060409081850190868401855b828110156156f457815180516001600160a01b031685528681015163ffffffff16878601528581015167ffffffffffffffff908116878701526060808301519091169086015260809081015115159085015260a09093019290850190600101615993565b600060408284031215615a0957600080fd5b615a11614702565b615a1a83615067565b81526020830151615a2a8161477a565b60208201529392505050565b634e487b7160e01b600052603160045260246000fd5b60008251615a5e81846020870161456d565b919091019291505056fea164736f6c6343000818000a",
}

var EVM2EVMMultiOnRampABI = EVM2EVMMultiOnRampMetaData.ABI

var EVM2EVMMultiOnRampBin = EVM2EVMMultiOnRampMetaData.Bin

func DeployEVM2EVMMultiOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMMultiOnRampStaticConfig, dynamicConfig EVM2EVMMultiOnRampDynamicConfig, destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs, rateLimiterConfig RateLimiterConfig, feeTokenConfigs []EVM2EVMMultiOnRampFeeTokenConfigArgs, tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs) (common.Address, *types.Transaction, *EVM2EVMMultiOnRamp, error) {
	parsed, err := EVM2EVMMultiOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMMultiOnRampBin), backend, staticConfig, dynamicConfig, destChainConfigArgs, rateLimiterConfig, feeTokenConfigs, tokenTransferFeeConfigArgs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMMultiOnRamp{address: address, abi: *parsed, EVM2EVMMultiOnRampCaller: EVM2EVMMultiOnRampCaller{contract: contract}, EVM2EVMMultiOnRampTransactor: EVM2EVMMultiOnRampTransactor{contract: contract}, EVM2EVMMultiOnRampFilterer: EVM2EVMMultiOnRampFilterer{contract: contract}}, nil
}

type EVM2EVMMultiOnRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMMultiOnRampCaller
	EVM2EVMMultiOnRampTransactor
	EVM2EVMMultiOnRampFilterer
}

type EVM2EVMMultiOnRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMMultiOnRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMMultiOnRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMMultiOnRampSession struct {
	Contract     *EVM2EVMMultiOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMMultiOnRampCallerSession struct {
	Contract *EVM2EVMMultiOnRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMMultiOnRampTransactorSession struct {
	Contract     *EVM2EVMMultiOnRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMMultiOnRampRaw struct {
	Contract *EVM2EVMMultiOnRamp
}

type EVM2EVMMultiOnRampCallerRaw struct {
	Contract *EVM2EVMMultiOnRampCaller
}

type EVM2EVMMultiOnRampTransactorRaw struct {
	Contract *EVM2EVMMultiOnRampTransactor
}

func NewEVM2EVMMultiOnRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMMultiOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMMultiOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMMultiOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRamp{address: address, abi: abi, EVM2EVMMultiOnRampCaller: EVM2EVMMultiOnRampCaller{contract: contract}, EVM2EVMMultiOnRampTransactor: EVM2EVMMultiOnRampTransactor{contract: contract}, EVM2EVMMultiOnRampFilterer: EVM2EVMMultiOnRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMMultiOnRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMMultiOnRampCaller, error) {
	contract, err := bindEVM2EVMMultiOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampCaller{contract: contract}, nil
}

func NewEVM2EVMMultiOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMMultiOnRampTransactor, error) {
	contract, err := bindEVM2EVMMultiOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampTransactor{contract: contract}, nil
}

func NewEVM2EVMMultiOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMMultiOnRampFilterer, error) {
	contract, err := bindEVM2EVMMultiOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampFilterer{contract: contract}, nil
}

func bindEVM2EVMMultiOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EVM2EVMMultiOnRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMMultiOnRamp.Contract.EVM2EVMMultiOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.EVM2EVMMultiOnRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.EVM2EVMMultiOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMMultiOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "currentRateLimiterState")

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _EVM2EVMMultiOnRamp.Contract.CurrentRateLimiterState(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _EVM2EVMMultiOnRamp.Contract.CurrentRateLimiterState(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (EVM2EVMMultiOnRampDestChainConfig, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	if err != nil {
		return *new(EVM2EVMMultiOnRampDestChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMMultiOnRampDestChainConfig)).(*EVM2EVMMultiOnRampDestChainConfig)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetDestChainConfig(destChainSelector uint64) (EVM2EVMMultiOnRampDestChainConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetDestChainConfig(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetDestChainConfig(destChainSelector uint64) (EVM2EVMMultiOnRampDestChainConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetDestChainConfig(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMMultiOnRampDynamicConfig, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(EVM2EVMMultiOnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMMultiOnRampDynamicConfig)).(*EVM2EVMMultiOnRampDynamicConfig)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetDynamicConfig() (EVM2EVMMultiOnRampDynamicConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetDynamicConfig(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetDynamicConfig() (EVM2EVMMultiOnRampDynamicConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetDynamicConfig(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getExpectedNextSequenceNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetFee(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector, message)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetFee(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector, message)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetFeeTokenConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMMultiOnRampFeeTokenConfig, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getFeeTokenConfig", token)

	if err != nil {
		return *new(EVM2EVMMultiOnRampFeeTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMMultiOnRampFeeTokenConfig)).(*EVM2EVMMultiOnRampFeeTokenConfig)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetFeeTokenConfig(token common.Address) (EVM2EVMMultiOnRampFeeTokenConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetFeeTokenConfig(&_EVM2EVMMultiOnRamp.CallOpts, token)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetFeeTokenConfig(token common.Address) (EVM2EVMMultiOnRampFeeTokenConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetFeeTokenConfig(&_EVM2EVMMultiOnRamp.CallOpts, token)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetFeeTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getFeeTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetFeeTokens() ([]common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetFeeTokens(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetFeeTokens() ([]common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetFeeTokens(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetNopFeesJuels(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getNopFeesJuels")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetNopFeesJuels() (*big.Int, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetNopFeesJuels(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetNopFeesJuels() (*big.Int, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetNopFeesJuels(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getPoolBySourceToken", arg0, sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetPoolBySourceToken(&_EVM2EVMMultiOnRamp.CallOpts, arg0, sourceToken)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetPoolBySourceToken(&_EVM2EVMMultiOnRamp.CallOpts, arg0, sourceToken)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetSenderNonce(opts *bind.CallOpts, destChainSelector uint64, sender common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getSenderNonce", destChainSelector, sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetSenderNonce(destChainSelector uint64, sender common.Address) (uint64, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetSenderNonce(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector, sender)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetSenderNonce(destChainSelector uint64, sender common.Address) (uint64, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetSenderNonce(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector, sender)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetStaticConfig(opts *bind.CallOpts) (EVM2EVMMultiOnRampStaticConfig, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(EVM2EVMMultiOnRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMMultiOnRampStaticConfig)).(*EVM2EVMMultiOnRampStaticConfig)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetStaticConfig() (EVM2EVMMultiOnRampStaticConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetStaticConfig(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetStaticConfig() (EVM2EVMMultiOnRampStaticConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetStaticConfig(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getSupportedTokens", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetSupportedTokens(&_EVM2EVMMultiOnRamp.CallOpts, arg0)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetSupportedTokens(&_EVM2EVMMultiOnRamp.CallOpts, arg0)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMMultiOnRampTokenTransferFeeConfig, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getTokenTransferFeeConfig", token)

	if err != nil {
		return *new(EVM2EVMMultiOnRampTokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMMultiOnRampTokenTransferFeeConfig)).(*EVM2EVMMultiOnRampTokenTransferFeeConfig)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetTokenTransferFeeConfig(token common.Address) (EVM2EVMMultiOnRampTokenTransferFeeConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetTokenTransferFeeConfig(&_EVM2EVMMultiOnRamp.CallOpts, token)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetTokenTransferFeeConfig(token common.Address) (EVM2EVMMultiOnRampTokenTransferFeeConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetTokenTransferFeeConfig(&_EVM2EVMMultiOnRamp.CallOpts, token)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) Owner() (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.Owner(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.Owner(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMMultiOnRamp.Contract.TypeAndVersion(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMMultiOnRamp.Contract.TypeAndVersion(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.AcceptOwnership(&_EVM2EVMMultiOnRamp.TransactOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.AcceptOwnership(&_EVM2EVMMultiOnRamp.TransactOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) ApplyDestChainConfigUpdates(destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyDestChainConfigUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, destChainConfigArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyDestChainConfigUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, destChainConfigArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "forwardFromRouter", destChainSelector, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ForwardFromRouter(&_EVM2EVMMultiOnRamp.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ForwardFromRouter(&_EVM2EVMMultiOnRamp.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) PayNops(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "payNops")
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) PayNops() (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.PayNops(&_EVM2EVMMultiOnRamp.TransactOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) PayNops() (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.PayNops(&_EVM2EVMMultiOnRamp.TransactOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "setAdmin", newAdmin)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetAdmin(&_EVM2EVMMultiOnRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetAdmin(&_EVM2EVMMultiOnRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMMultiOnRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) SetDynamicConfig(dynamicConfig EVM2EVMMultiOnRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetDynamicConfig(&_EVM2EVMMultiOnRamp.TransactOpts, dynamicConfig)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) SetDynamicConfig(dynamicConfig EVM2EVMMultiOnRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetDynamicConfig(&_EVM2EVMMultiOnRamp.TransactOpts, dynamicConfig)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) SetFeeTokenConfig(opts *bind.TransactOpts, feeTokenConfigArgs []EVM2EVMMultiOnRampFeeTokenConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "setFeeTokenConfig", feeTokenConfigArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) SetFeeTokenConfig(feeTokenConfigArgs []EVM2EVMMultiOnRampFeeTokenConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetFeeTokenConfig(&_EVM2EVMMultiOnRamp.TransactOpts, feeTokenConfigArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) SetFeeTokenConfig(feeTokenConfigArgs []EVM2EVMMultiOnRampFeeTokenConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetFeeTokenConfig(&_EVM2EVMMultiOnRamp.TransactOpts, feeTokenConfigArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetRateLimiterConfig(&_EVM2EVMMultiOnRamp.TransactOpts, config)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetRateLimiterConfig(&_EVM2EVMMultiOnRamp.TransactOpts, config)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) SetTokenTransferFeeConfig(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "setTokenTransferFeeConfig", tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) SetTokenTransferFeeConfig(tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetTokenTransferFeeConfig(&_EVM2EVMMultiOnRamp.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) SetTokenTransferFeeConfig(tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetTokenTransferFeeConfig(&_EVM2EVMMultiOnRamp.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.TransferOwnership(&_EVM2EVMMultiOnRamp.TransactOpts, to)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.TransferOwnership(&_EVM2EVMMultiOnRamp.TransactOpts, to)
}

type EVM2EVMMultiOnRampAdminSetIterator struct {
	Event *EVM2EVMMultiOnRampAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampAdminSet)
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
		it.Event = new(EVM2EVMMultiOnRampAdminSet)
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

func (it *EVM2EVMMultiOnRampAdminSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampAdminSet struct {
	NewAdmin common.Address
	Raw      types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampAdminSetIterator, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampAdminSetIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAdminSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampAdminSet)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseAdminSet(log types.Log) (*EVM2EVMMultiOnRampAdminSet, error) {
	event := new(EVM2EVMMultiOnRampAdminSet)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "AdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampCCIPSendRequestedIterator struct {
	Event *EVM2EVMMultiOnRampCCIPSendRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampCCIPSendRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampCCIPSendRequested)
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
		it.Event = new(EVM2EVMMultiOnRampCCIPSendRequested)
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

func (it *EVM2EVMMultiOnRampCCIPSendRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampCCIPSendRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampCCIPSendRequested struct {
	Message InternalEVM2EVMMessage
	Raw     types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampCCIPSendRequestedIterator, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampCCIPSendRequestedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "CCIPSendRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampCCIPSendRequested) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampCCIPSendRequested)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseCCIPSendRequested(log types.Log) (*EVM2EVMMultiOnRampCCIPSendRequested, error) {
	event := new(EVM2EVMMultiOnRampCCIPSendRequested)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampConfigChangedIterator struct {
	Event *EVM2EVMMultiOnRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampConfigChanged)
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
		it.Event = new(EVM2EVMMultiOnRampConfigChanged)
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

func (it *EVM2EVMMultiOnRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampConfigChangedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampConfigChanged)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMMultiOnRampConfigChanged, error) {
	event := new(EVM2EVMMultiOnRampConfigChanged)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampConfigSetIterator struct {
	Event *EVM2EVMMultiOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampConfigSet)
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
		it.Event = new(EVM2EVMMultiOnRampConfigSet)
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

func (it *EVM2EVMMultiOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampConfigSet struct {
	StaticConfig  EVM2EVMMultiOnRampStaticConfig
	DynamicConfig EVM2EVMMultiOnRampDynamicConfig
	Raw           types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampConfigSetIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampConfigSet)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseConfigSet(log types.Log) (*EVM2EVMMultiOnRampConfigSet, error) {
	event := new(EVM2EVMMultiOnRampConfigSet)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampDestChainAddedIterator struct {
	Event *EVM2EVMMultiOnRampDestChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampDestChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampDestChainAdded)
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
		it.Event = new(EVM2EVMMultiOnRampDestChainAdded)
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

func (it *EVM2EVMMultiOnRampDestChainAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampDestChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampDestChainAdded struct {
	DestChainSelector uint64
	DestChainConfig   EVM2EVMMultiOnRampDestChainConfig
	Raw               types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampDestChainAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampDestChainAddedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "DestChainAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampDestChainAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampDestChainAdded)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseDestChainAdded(log types.Log) (*EVM2EVMMultiOnRampDestChainAdded, error) {
	event := new(EVM2EVMMultiOnRampDestChainAdded)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampDestChainDynamicConfigUpdatedIterator struct {
	Event *EVM2EVMMultiOnRampDestChainDynamicConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampDestChainDynamicConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampDestChainDynamicConfigUpdated)
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
		it.Event = new(EVM2EVMMultiOnRampDestChainDynamicConfigUpdated)
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

func (it *EVM2EVMMultiOnRampDestChainDynamicConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampDestChainDynamicConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampDestChainDynamicConfigUpdated struct {
	DestChainSelector uint64
	DynamicConfig     EVM2EVMMultiOnRampDestChainDynamicConfig
	Raw               types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterDestChainDynamicConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampDestChainDynamicConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "DestChainDynamicConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampDestChainDynamicConfigUpdatedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "DestChainDynamicConfigUpdated", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchDestChainDynamicConfigUpdated(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampDestChainDynamicConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "DestChainDynamicConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampDestChainDynamicConfigUpdated)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "DestChainDynamicConfigUpdated", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseDestChainDynamicConfigUpdated(log types.Log) (*EVM2EVMMultiOnRampDestChainDynamicConfigUpdated, error) {
	event := new(EVM2EVMMultiOnRampDestChainDynamicConfigUpdated)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "DestChainDynamicConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampFeeConfigSetIterator struct {
	Event *EVM2EVMMultiOnRampFeeConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampFeeConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampFeeConfigSet)
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
		it.Event = new(EVM2EVMMultiOnRampFeeConfigSet)
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

func (it *EVM2EVMMultiOnRampFeeConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampFeeConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampFeeConfigSet struct {
	FeeConfig []EVM2EVMMultiOnRampFeeTokenConfigArgs
	Raw       types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampFeeConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "FeeConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampFeeConfigSetIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "FeeConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampFeeConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "FeeConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampFeeConfigSet)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "FeeConfigSet", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseFeeConfigSet(log types.Log) (*EVM2EVMMultiOnRampFeeConfigSet, error) {
	event := new(EVM2EVMMultiOnRampFeeConfigSet)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "FeeConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampNopsPaidIterator struct {
	Event *EVM2EVMMultiOnRampNopsPaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampNopsPaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampNopsPaid)
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
		it.Event = new(EVM2EVMMultiOnRampNopsPaid)
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

func (it *EVM2EVMMultiOnRampNopsPaidIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampNopsPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampNopsPaid struct {
	FeeAggregator common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterNopsPaid(opts *bind.FilterOpts, feeAggregator []common.Address) (*EVM2EVMMultiOnRampNopsPaidIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "NopsPaid", feeAggregatorRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampNopsPaidIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "NopsPaid", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchNopsPaid(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampNopsPaid, feeAggregator []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "NopsPaid", feeAggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampNopsPaid)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "NopsPaid", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseNopsPaid(log types.Log) (*EVM2EVMMultiOnRampNopsPaid, error) {
	event := new(EVM2EVMMultiOnRampNopsPaid)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "NopsPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMMultiOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMMultiOnRampOwnershipTransferRequested)
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

func (it *EVM2EVMMultiOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampOwnershipTransferRequestedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampOwnershipTransferRequested)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMMultiOnRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMMultiOnRampOwnershipTransferRequested)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampOwnershipTransferredIterator struct {
	Event *EVM2EVMMultiOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMMultiOnRampOwnershipTransferred)
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

func (it *EVM2EVMMultiOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampOwnershipTransferredIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampOwnershipTransferred)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMMultiOnRampOwnershipTransferred, error) {
	event := new(EVM2EVMMultiOnRampOwnershipTransferred)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampTokenTransferFeeConfigDeletedIterator struct {
	Event *EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted)
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
		it.Event = new(EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted)
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

func (it *EVM2EVMMultiOnRampTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted struct {
	Tokens []common.Address
	Raw    types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampTokenTransferFeeConfigDeletedIterator, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampTokenTransferFeeConfigDeletedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted, error) {
	event := new(EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampTokenTransferFeeConfigSetIterator struct {
	Event *EVM2EVMMultiOnRampTokenTransferFeeConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampTokenTransferFeeConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampTokenTransferFeeConfigSet)
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
		it.Event = new(EVM2EVMMultiOnRampTokenTransferFeeConfigSet)
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

func (it *EVM2EVMMultiOnRampTokenTransferFeeConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampTokenTransferFeeConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampTokenTransferFeeConfigSet struct {
	TransferFeeConfig []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs
	Raw               types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterTokenTransferFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampTokenTransferFeeConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "TokenTransferFeeConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampTokenTransferFeeConfigSetIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "TokenTransferFeeConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchTokenTransferFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampTokenTransferFeeConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "TokenTransferFeeConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampTokenTransferFeeConfigSet)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "TokenTransferFeeConfigSet", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseTokenTransferFeeConfigSet(log types.Log) (*EVM2EVMMultiOnRampTokenTransferFeeConfigSet, error) {
	event := new(EVM2EVMMultiOnRampTokenTransferFeeConfigSet)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "TokenTransferFeeConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampTokensConsumedIterator struct {
	Event *EVM2EVMMultiOnRampTokensConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampTokensConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampTokensConsumed)
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
		it.Event = new(EVM2EVMMultiOnRampTokensConsumed)
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

func (it *EVM2EVMMultiOnRampTokensConsumedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampTokensConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampTokensConsumed struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterTokensConsumed(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampTokensConsumedIterator, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampTokensConsumedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "TokensConsumed", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampTokensConsumed) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampTokensConsumed)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseTokensConsumed(log types.Log) (*EVM2EVMMultiOnRampTokensConsumed, error) {
	event := new(EVM2EVMMultiOnRampTokensConsumed)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMMultiOnRamp.abi.Events["AdminSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseAdminSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["CCIPSendRequested"].ID:
		return _EVM2EVMMultiOnRamp.ParseCCIPSendRequested(log)
	case _EVM2EVMMultiOnRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMMultiOnRamp.ParseConfigChanged(log)
	case _EVM2EVMMultiOnRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseConfigSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["DestChainAdded"].ID:
		return _EVM2EVMMultiOnRamp.ParseDestChainAdded(log)
	case _EVM2EVMMultiOnRamp.abi.Events["DestChainDynamicConfigUpdated"].ID:
		return _EVM2EVMMultiOnRamp.ParseDestChainDynamicConfigUpdated(log)
	case _EVM2EVMMultiOnRamp.abi.Events["FeeConfigSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseFeeConfigSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["NopsPaid"].ID:
		return _EVM2EVMMultiOnRamp.ParseNopsPaid(log)
	case _EVM2EVMMultiOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMMultiOnRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMMultiOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMMultiOnRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMMultiOnRamp.abi.Events["TokenTransferFeeConfigDeleted"].ID:
		return _EVM2EVMMultiOnRamp.ParseTokenTransferFeeConfigDeleted(log)
	case _EVM2EVMMultiOnRamp.abi.Events["TokenTransferFeeConfigSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseTokenTransferFeeConfigSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["TokensConsumed"].ID:
		return _EVM2EVMMultiOnRamp.ParseTokensConsumed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMMultiOnRampAdminSet) Topic() common.Hash {
	return common.HexToHash("0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c")
}

func (EVM2EVMMultiOnRampCCIPSendRequested) Topic() common.Hash {
	return common.HexToHash("0xd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd")
}

func (EVM2EVMMultiOnRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (EVM2EVMMultiOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xa0d39dfd39d411e1d09ac113d962b2f31903b670d3d79ec23dc1a60b55238012")
}

func (EVM2EVMMultiOnRampDestChainAdded) Topic() common.Hash {
	return common.HexToHash("0xb6031b30c90eb9236da8c32ffd47ca6307f4a0eaa17e78e12f20a24afd580510")
}

func (EVM2EVMMultiOnRampDestChainDynamicConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x8a71e142c5eeae8754825ba46ee7cd42468c1621a90b72dcad8bf7ee8c7111c3")
}

func (EVM2EVMMultiOnRampFeeConfigSet) Topic() common.Hash {
	return common.HexToHash("0x067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e")
}

func (EVM2EVMMultiOnRampNopsPaid) Topic() common.Hash {
	return common.HexToHash("0x88db9a873cb1a9b6d3757fc0fc19ede741b35af2a9a497a5793e760cdd97966b")
}

func (EVM2EVMMultiOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMMultiOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0xfb95a0042158e60a33e7b5bec100f3d95407b1a71bee6633bd54b8887449750b")
}

func (EVM2EVMMultiOnRampTokenTransferFeeConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf5791bc457b3bb990493cf5f655db46c25ccf5764c9b99b8969b4c72ea7df9d0")
}

func (EVM2EVMMultiOnRampTokensConsumed) Topic() common.Hash {
	return common.HexToHash("0x1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a")
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRamp) Address() common.Address {
	return _EVM2EVMMultiOnRamp.address
}

type EVM2EVMMultiOnRampInterface interface {
	CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (EVM2EVMMultiOnRampDestChainConfig, error)

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMMultiOnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetFeeTokenConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMMultiOnRampFeeTokenConfig, error)

	GetFeeTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetNopFeesJuels(opts *bind.CallOpts) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetSenderNonce(opts *bind.CallOpts, destChainSelector uint64, sender common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMMultiOnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMMultiOnRampTokenTransferFeeConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	PayNops(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMMultiOnRampDynamicConfig) (*types.Transaction, error)

	SetFeeTokenConfig(opts *bind.TransactOpts, feeTokenConfigArgs []EVM2EVMMultiOnRampFeeTokenConfigArgs) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	SetTokenTransferFeeConfig(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampAdminSetIterator, error)

	WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAdminSet) (event.Subscription, error)

	ParseAdminSet(log types.Log) (*EVM2EVMMultiOnRampAdminSet, error)

	FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampCCIPSendRequestedIterator, error)

	WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampCCIPSendRequested) (event.Subscription, error)

	ParseCCIPSendRequested(log types.Log) (*EVM2EVMMultiOnRampCCIPSendRequested, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMMultiOnRampConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMMultiOnRampConfigSet, error)

	FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampDestChainAddedIterator, error)

	WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampDestChainAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainAdded(log types.Log) (*EVM2EVMMultiOnRampDestChainAdded, error)

	FilterDestChainDynamicConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampDestChainDynamicConfigUpdatedIterator, error)

	WatchDestChainDynamicConfigUpdated(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampDestChainDynamicConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainDynamicConfigUpdated(log types.Log) (*EVM2EVMMultiOnRampDestChainDynamicConfigUpdated, error)

	FilterFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampFeeConfigSetIterator, error)

	WatchFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampFeeConfigSet) (event.Subscription, error)

	ParseFeeConfigSet(log types.Log) (*EVM2EVMMultiOnRampFeeConfigSet, error)

	FilterNopsPaid(opts *bind.FilterOpts, feeAggregator []common.Address) (*EVM2EVMMultiOnRampNopsPaidIterator, error)

	WatchNopsPaid(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampNopsPaid, feeAggregator []common.Address) (event.Subscription, error)

	ParseNopsPaid(log types.Log) (*EVM2EVMMultiOnRampNopsPaid, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMMultiOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMMultiOnRampOwnershipTransferred, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampTokenTransferFeeConfigSetIterator, error)

	WatchTokenTransferFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampTokenTransferFeeConfigSet) (event.Subscription, error)

	ParseTokenTransferFeeConfigSet(log types.Log) (*EVM2EVMMultiOnRampTokenTransferFeeConfigSet, error)

	FilterTokensConsumed(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampTokensConsumedIterator, error)

	WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampTokensConsumed) (event.Subscription, error)

	ParseTokensConsumed(log types.Log) (*EVM2EVMMultiOnRampTokensConsumed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
