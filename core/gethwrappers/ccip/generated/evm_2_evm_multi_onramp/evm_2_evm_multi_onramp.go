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
	DynamicConfig EVM2EVMMultiOnRampDestChainDynamicConfig
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
}

type EVM2EVMMultiOnRampDestChainDynamicConfigArgs struct {
	DestChainSelector uint64
	DynamicConfig     EVM2EVMMultiOnRampDestChainDynamicConfig
}

type EVM2EVMMultiOnRampDynamicConfig struct {
	Router             common.Address
	PriceRegistry      common.Address
	TokenAdminRegistry common.Address
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

type EVM2EVMMultiOnRampNopAndWeight struct {
	Nop    common.Address
	Weight uint16
}

type EVM2EVMMultiOnRampStaticConfig struct {
	LinkToken         common.Address
	ChainSelector     uint64
	DestChainSelector uint64
	DefaultTxGasLimit uint64
	MaxNopFeesJuels   *big.Int
	PrevOnRamp        common.Address
	ArmProxy          common.Address
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfigArgs[]\",\"name\":\"destChainDynamicConfigsArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadARMSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSendZeroTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChainSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"}],\"name\":\"InvalidNopAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkBalanceNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeBalanceReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeesToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNopsToPay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NotAFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdminOrNop\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SourceTokenDataTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyNops\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeConfig\",\"type\":\"tuple[]\"}],\"name\":\"FeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NopPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nopWeightsTotal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"NopsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokenTransferFeeConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"transferFeeConfig\",\"type\":\"tuple[]\"}],\"name\":\"TokenTransferFeeConfigSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfigArgs[]\",\"name\":\"destChainDynamicConfigsArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyDestChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestChainConfig\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeeTokenConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.FeeTokenConfig\",\"name\":\"feeTokenConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNopFeesJuels\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNops\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"weightsTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenTransferFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"setNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"address[]\"}],\"name\":\"setTokenTransferFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawNonLinkFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b50604051620084a8380380620084a8833981016040819052620000359162001ca2565b8333806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c081620002f4565b50506040805160a081018252602084810180516001600160801b039081168085524263ffffffff169385018490528751151585870181905292518216606086018190529790950151166080909301839052600380546001600160a01b031916909417600160801b9283021760ff60a01b1916600160a01b90910217909255029091176004555086516001600160a01b0316158062000169575060208701516001600160401b0316155b8062000180575060408701516001600160401b0316155b8062000197575060608701516001600160401b0316155b80620001ae575060c08701516001600160a01b0316155b15620001cd576040516306b7c75960e31b815260040160405180910390fd5b6020808801516040808a015181517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3948101949094526001600160401b039283169184019190915216606082015230608082015260a00160408051601f198184030181529181528151602092830120608090815289516001600160a01b0390811660e052928a01516001600160401b0390811661010052918a015182166101205260608a015190911660a0908152908901516001600160601b031660c0908152908901518216610140528801511661016052620002aa866200039f565b620002b585620004b4565b620002c0836200076d565b604080516000815260208101909152620002dc908390620008a7565b620002e78162000b0f565b50505050505050620022f6565b336001600160a01b038216036200034e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60208101516001600160a01b0316620003cb576040516306b7c75960e31b815260040160405180910390fd5b8051600580546001600160a01b03199081166001600160a01b03938416179091556020808401516006805484169185169190911790556040808501516007805490941690851617909255815160e080820184525184168152610100516001600160401b03908116928201929092526101205182168184015260a08051909216606082015260c080516001600160601b03166080830152610140518516928201929092526101605190931690830152517ff98cd3cd66fb6e771ed76b7e0209a9c4d3a0eea42e3bce57226c92a0cdb56dff91620004a991849062001d9d565b60405180910390a150565b60005b815181101562000769576000828281518110620004d857620004d862001e32565b60200260200101516020015190506000838381518110620004fd57620004fd62001e32565b60200260200101516000015190506000604051806020016040528084815250905080600b6000846001600160401b03166001600160401b0316815260200190815260200160002060008201518160000160008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055505050905050816001600160401b03167f9fda653fbd9a0d9a27a82d8f481a3d81e76169fb54108f481b1da82a0875bdf6826040516200074a919062001e48565b60405180910390a250505080620007619062001f60565b9050620004b7565b5050565b60005b81518110156200087557600082828151811062000791576200079162001e32565b60209081029190910181015160408051608080820183528385015163ffffffff9081168352838501516001600160401b03908116848801908152606080880151831686880190815294880151151590860190815296516001600160a01b03166000908152600c90985294909620925183549451925195511515600160a01b0260ff60a01b199688166c010000000000000000000000000296909616600160601b600160a81b031993909716640100000000026001600160601b0319909516911617929092179190911692909217179055506200086d8162001f60565b905062000770565b507f067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e81604051620004a9919062001f7c565b60005b825181101562000a25576000838281518110620008cb57620008cb62001e32565b6020908102919091018101516040805160e0810182528284015163ffffffff908116825282840151811682860190815260608086015161ffff908116858701908152608080890151861693870193845260a0808a0151871691880191825260c0808b01511515918901918252600190890190815299516001600160a01b03166000908152600d909b5297909920955186549451915193519951975198511515600160981b0260ff60981b19991515600160901b0260ff60901b19998816600160701b029990991664ffffffffff60701b199b88166a01000000000000000000000263ffffffff60501b199690951668010000000000000000029590951665ffffffffffff60401b19938816640100000000026001600160401b0319909716929097169190911794909417169390931792909217959095161791909117919091169190911790555062000a1d8162001f60565b9050620008aa565b507ff5791bc457b3bb990493cf5f655db46c25ccf5764c9b99b8969b4c72ea7df9d08260405162000a5791906200200b565b60405180910390a160005b815181101562000ac957600d600083838151811062000a855762000a8562001e32565b6020908102919091018101516001600160a01b0316825281019190915260400160002080546001600160a01b031916905562000ac18162001f60565b905062000a62565b5080511562000769577ffb95a0042158e60a33e7b5bec100f3d95407b1a71bee6633bd54b8887449750b8160405162000b039190620020a0565b60405180910390a15050565b8051604081111562000b3457604051635ad0867d60e11b815260040160405180910390fd5b600f546c01000000000000000000000000900463ffffffff161580159062000b7e5750600f5463ffffffff6c010000000000000000000000008204166001600160601b0390911610155b1562000b8e5762000b8e62000d39565b600062000b9c600862000f39565b90505b801562000be857600062000bc262000bb9600184620020ef565b60089062000f4c565b50905062000bd260088262000f6a565b50508062000be09062002105565b905062000b9f565b506000805b8281101562000cd057600084828151811062000c0d5762000c0d62001e32565b6020026020010151600001519050600085838151811062000c325762000c3262001e32565b602002602001015160200151905060e0516001600160a01b0316826001600160a01b0316148062000c6a57506001600160a01b038216155b1562000c9557604051634de938d160e01b81526001600160a01b038316600482015260240162000084565b62000ca760088361ffff841662000f88565b5062000cb861ffff8216856200211f565b935050508062000cc89062001f60565b905062000bed565b50600f805463ffffffff60601b19166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd249062000d2c90839086906200213f565b60405180910390a1505050565b6000546001600160a01b0316331480159062000d6057506002546001600160a01b03163314155b801562000d77575062000d7560083362000fa8565b155b1562000d965760405163032bb72b60e31b815260040160405180910390fd5b600f546c01000000000000000000000000900463ffffffff16600081900362000dd25760405163990e30bf60e01b815260040160405180910390fd5b600f546001600160601b03168181101562000e00576040516311a1ee3b60e31b815260040160405180910390fd5b600062000e0c62000fbf565b121562000e2c57604051631e9acf1760e31b815260040160405180910390fd5b80600062000e3b600862000f39565b905060005b8181101562000f135760008062000e5960088462000f4c565b909250905060008762000e76836001600160601b038a16620021ac565b62000e829190620021c6565b905062000e908187620021e9565b60e05190965062000eb5906001600160a01b0316846001600160601b0384166200104d565b6040516001600160601b03821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a25050508062000f0b9062001f60565b905062000e40565b5050600f80546001600160601b0319166001600160601b03929092169190911790555050565b600062000f4682620010aa565b92915050565b600080808062000f5d8686620010b7565b9097909650945050505050565b600062000f81836001600160a01b038416620010e4565b9392505050565b600062000fa0846001600160a01b0385168462001103565b949350505050565b600062000f81836001600160a01b03841662001122565b600f5460e0516040516370a0823160e01b81523060048201526000926001600160601b0316916001600160a01b0316906370a0823190602401602060405180830381865afa15801562001016573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200103c91906200220c565b62001048919062002226565b905090565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b0390811663a9059cbb60e01b17909152620010a59185916200113016565b505050565b600062000f468262001201565b60008080620010c785856200120c565b600081815260029690960160205260409095205494959350505050565b6000818152600283016020526040812081905562000f8183836200121a565b6000828152600284016020526040812082905562000fa0848462001228565b600062000f81838362001236565b6040805180820190915260208082527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564908201526000906200117f906001600160a01b0385169084906200124f565b805190915015620010a55780806020019051810190620011a0919062002249565b620010a55760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840162000084565b600062000f46825490565b600062000f81838362001260565b600062000f8183836200128d565b600062000f81838362001398565b6000818152600183016020526040812054151562000f81565b606062000fa08484600085620013ea565b60008260000182815481106200127a576200127a62001e32565b9060005260206000200154905092915050565b6000818152600183016020526040812054801562001386576000620012b4600183620020ef565b8554909150600090620012ca90600190620020ef565b905081811462001336576000866000018281548110620012ee57620012ee62001e32565b906000526020600020015490508087600001848154811062001314576200131462001e32565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806200134a576200134a62002267565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000f46565b600091505062000f46565b5092915050565b6000818152600183016020526040812054620013e15750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000f46565b50600062000f46565b6060824710156200144d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840162000084565b600080866001600160a01b031685876040516200146b9190620022a3565b60006040518083038185875af1925050503d8060008114620014aa576040519150601f19603f3d011682016040523d82523d6000602084013e620014af565b606091505b509092509050620014c387838387620014ce565b979650505050505050565b60608315620015425782516000036200153a576001600160a01b0385163b6200153a5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640162000084565b508162000fa0565b62000fa08383815115620015595781518083602001fd5b8060405162461bcd60e51b8152600401620000849190620022c1565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b0381118282101715620015b057620015b062001575565b60405290565b604051606081016001600160401b0381118282101715620015b057620015b062001575565b604080519081016001600160401b0381118282101715620015b057620015b062001575565b60405161018081016001600160401b0381118282101715620015b057620015b062001575565b60405160a081016001600160401b0381118282101715620015b057620015b062001575565b604051601f8201601f191681016001600160401b038111828210171562001676576200167662001575565b604052919050565b80516001600160a01b03811681146200169657600080fd5b919050565b80516001600160401b03811681146200169657600080fd5b600060e08284031215620016c657600080fd5b620016d06200158b565b9050620016dd826200167e565b8152620016ed602083016200169b565b602082015262001700604083016200169b565b604082015262001713606083016200169b565b606082015260808201516001600160601b03811681146200173357600080fd5b60808201526200174660a083016200167e565b60a08201526200175960c083016200167e565b60c082015292915050565b6000606082840312156200177757600080fd5b62001781620015b6565b90506200178e826200167e565b81526200179e602083016200167e565b6020820152620017b1604083016200167e565b604082015292915050565b60006001600160401b03821115620017d857620017d862001575565b5060051b60200190565b805180151581146200169657600080fd5b805161ffff811681146200169657600080fd5b805163ffffffff811681146200169657600080fd5b600082601f8301126200182d57600080fd5b81516020620018466200184083620017bc565b6200164b565b8281526101a092830285018201928282019190878511156200186757600080fd5b8387015b85811015620019c05780890382811215620018865760008081fd5b62001890620015db565b6200189b836200169b565b815261018080601f1984011215620018b35760008081fd5b620018bd62001600565b9250620018cc888501620017e2565b83526040620018dd818601620017f3565b898501526060620018f081870162001806565b82860152608091506200190582870162001806565b9085015260a06200191886820162001806565b8286015260c091506200192d828701620017f3565b9085015260e06200194086820162001806565b82860152610100915062001956828701620017f3565b908501526101206200196a868201620017f3565b82860152610140915062001980828701620017f3565b908501526101606200199486820162001806565b82860152620019a583870162001806565b9085015250508087019190915284529284019281016200186b565b5090979650505050505050565b80516001600160801b03811681146200169657600080fd5b600060608284031215620019f857600080fd5b62001a02620015b6565b905062001a0f82620017e2565b815262001a1f60208301620019cd565b6020820152620017b160408301620019cd565b600082601f83011262001a4457600080fd5b8151602062001a576200184083620017bc565b82815260a0928302850182019282820191908785111562001a7757600080fd5b8387015b85811015620019c05781818a03121562001a955760008081fd5b62001a9f62001626565b62001aaa826200167e565b815262001ab986830162001806565b86820152604062001acc8184016200169b565b90820152606062001adf8382016200169b565b90820152608062001af2838201620017e2565b90820152845292840192810162001a7b565b600082601f83011262001b1657600080fd5b8151602062001b296200184083620017bc565b82815260e0928302850182019282820191908785111562001b4957600080fd5b8387015b85811015620019c05781818a03121562001b675760008081fd5b62001b716200158b565b62001b7c826200167e565b815262001b8b86830162001806565b86820152604062001b9e81840162001806565b90820152606062001bb1838201620017f3565b90820152608062001bc483820162001806565b9082015260a062001bd783820162001806565b9082015260c062001bea838201620017e2565b90820152845292840192810162001b4d565b600082601f83011262001c0e57600080fd5b8151602062001c216200184083620017bc565b82815260069290921b8401810191818101908684111562001c4157600080fd5b8286015b8481101562001c97576040818903121562001c605760008081fd5b62001c6a620015db565b62001c75826200167e565b815262001c84858301620017f3565b8186015283529183019160400162001c45565b509695505050505050565b6000806000806000806000610220888a03121562001cbf57600080fd5b62001ccb8989620016b3565b965062001cdc8960e08a0162001764565b6101408901519096506001600160401b038082111562001cfb57600080fd5b62001d098b838c016200181b565b965062001d1b8b6101608c01620019e5565b95506101c08a015191508082111562001d3357600080fd5b62001d418b838c0162001a32565b94506101e08a015191508082111562001d5957600080fd5b62001d678b838c0162001b04565b93506102008a015191508082111562001d7f57600080fd5b5062001d8e8a828b0162001bfc565b91505092959891949750929550565b82516001600160a01b0390811682526020808501516001600160401b0390811682850152604080870151821681860152606080880151909216918501919091526080808701516001600160601b03169085015260a08087015184169085015260c0808701518416908501528451831660e085015290840151821661010084015283015116610120820152610140810162000f81565b634e487b7160e01b600052603260045260246000fd5b8151805115158252610180820190602081015162001e6c602085018261ffff169052565b50604081015162001e85604085018263ffffffff169052565b50606081015162001e9e606085018263ffffffff169052565b50608081015162001eb7608085018263ffffffff169052565b5060a081015162001ece60a085018261ffff169052565b5060c081015162001ee760c085018263ffffffff169052565b5060e081015162001efe60e085018261ffff169052565b506101008181015161ffff9081169185019190915261012080830151909116908401526101408082015163ffffffff908116918501919091526101609182015116920191909152919050565b634e487b7160e01b600052601160045260246000fd5b60006001820162001f755762001f7562001f4a565b5060010190565b602080825282518282018190526000919060409081850190868401855b8281101562001ffe57815180516001600160a01b031685528681015163ffffffff1687860152858101516001600160401b03908116878701526060808301519091169086015260809081015115159085015260a0909301929085019060010162001f99565b5091979650505050505050565b602080825282518282018190526000919060409081850190868401855b8281101562001ffe57815180516001600160a01b031685528681015163ffffffff908116888701528682015181168787015260608083015161ffff169087015260808083015182169087015260a0808301519091169086015260c09081015115159085015260e0909301929085019060010162002028565b6020808252825182820181905260009190848201906040850190845b81811015620020e35783516001600160a01b031683529284019291840191600101620020bc565b50909695505050505050565b8181038181111562000f465762000f4662001f4a565b60008162002117576200211762001f4a565b506000190190565b63ffffffff81811683821601908082111562001391576200139162001f4a565b6000604080830163ffffffff8616845260208281860152818651808452606087019150828801935060005b818110156200219e57845180516001600160a01b0316845284015161ffff168484015293830193918501916001016200216a565b509098975050505050505050565b808202811582820484141762000f465762000f4662001f4a565b600082620021e457634e487b7160e01b600052601260045260246000fd5b500490565b6001600160601b0382811682821603908082111562001391576200139162001f4a565b6000602082840312156200221f57600080fd5b5051919050565b818103600083128015838313168383128216171562001391576200139162001f4a565b6000602082840312156200225c57600080fd5b62000f8182620017e2565b634e487b7160e01b600052603160045260246000fd5b60005b838110156200229a57818101518382015260200162002280565b50506000910152565b60008251620022b78184602087016200227d565b9190910192915050565b6020815260008251806020840152620022e28160408501602087016200227d565b601f01601f19169190910160400192915050565b60805160a05160c05160e051610100516101205161014051610160516160a4620024046000396000818161033f0152818161173b01526130b101526000818161031001528181611438015281816114a001528181611cdf01528181611d47015261308901526000818161027c01528181610ac7015281816118780152818161222f015261300501526000818161024c01528181611e180152612fd101526000818161021d01528181611142015281816116ba01528181611a9201528181611b8b0152818161266a01528181612fac01526136ed0152600081816102dc01528181611c5701526130550152600081816102ac015281816127c7015261302c0152600061242601526160a46000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80637437ff9f11610104578063c92b2832116100a2578063eff7cc4811610071578063eff7cc4814610a75578063f25561fd14610a7d578063f2fde38b14610a90578063fbca3b7414610aa357600080fd5b8063c92b283214610a34578063d09dc33914610a47578063df0aa9e914610a4f578063e1ee5b2514610a6257600080fd5b8063856c8247116100de578063856c8247146108f55780638da5cb5b146109085780639a113c3614610919578063b06d41bc14610a1e57600080fd5b80637437ff9f1461087a57806376f6ae76146108da57806379ba5097146108ed57600080fd5b806348a98aa41161017157806354b714681161014b57806354b714681461065a578063599f64311461067a5780636def4ce71461068b578063704b6c021461086757600080fd5b806348a98aa4146105b8578063546719cd146105e3578063549e946f1461064757600080fd5b806320487ded116101ad57806320487ded1461054e57806334d560e41461056f5780634120fccd146105845780634816f4f7146105a557600080fd5b806306285c69146101d45780631772047e14610385578063181f5a7714610505575b600080fd5b61036f6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526040518060e001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b60405161037c9190614aba565b60405180910390f35b610499610393366004614b4b565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152506001600160a01b03166000908152600d6020908152604091829020825160e081018452905463ffffffff8082168352640100000000820481169383019390935261ffff68010000000000000000820416938201939093526a01000000000000000000008304821660608201526e0100000000000000000000000000008304909116608082015260ff720100000000000000000000000000000000000083048116151560a0830152730100000000000000000000000000000000000000909204909116151560c082015290565b60405161037c9190600060e08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015260c0830151151560c083015292915050565b6105416040518060400160405280601c81526020017f45564d3245564d4d756c74694f6e52616d7020312e362e302d6465760000000081525081565b60405161037c9190614bb8565b61056161055c366004614bf9565b610ac3565b60405190815260200161037c565b61058261057d366004614d46565b610f2e565b005b61058c610f42565b60405167ffffffffffffffff909116815260200161037c565b6105826105b3366004614e7f565b610f69565b6105cb6105c6366004614fbc565b610f7f565b6040516001600160a01b03909116815260200161037c565b6105eb61100e565b60405161037c919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b610582610655366004614ff5565b6110b1565b600f546040516bffffffffffffffffffffffff909116815260200161037c565b6002546001600160a01b03166105cb565b61085a610699366004615013565b604080516101a081018252600060208201818152928201819052606082018190526080820181905260a0820181905260c0820181905260e0820181905261010082018190526101208201819052610140820181905261016082018190526101808201529081525067ffffffffffffffff166000908152600b602090815260409182902082516101a081018452815460ff8116151593820193845261ffff61010080830482169684019690965263ffffffff63010000008304811660608501526701000000000000008304811660808501526b0100000000000000000000008304811660a08501526f010000000000000000000000000000008304821660c0850152710100000000000000000000000000000000008304811660e08501527501000000000000000000000000000000000000000000830482169684019690965277010000000000000000000000000000000000000000000000820481166101208401527901000000000000000000000000000000000000000000000000008204166101408301527b0100000000000000000000000000000000000000000000000000000090048416610160820152600190910154909216610180830152815290565b60405161037c9190615030565b610582610875366004614b4b565b61122a565b6108cd604080516060810182526000808252602082018190529181019190915250604080516060810182526005546001600160a01b03908116825260065481166020830152600754169181019190915290565b60405161037c9190615134565b6105826108e8366004615164565b6112e9565b61058261134c565b61058c610903366004614b4b565b61140a565b6000546001600160a01b03166105cb565b6109d4610927366004614b4b565b604080516080810182526000808252602082018190529181018290526060810191909152506001600160a01b03166000908152600c60209081526040918290208251608081018452905463ffffffff8116825267ffffffffffffffff64010000000082048116938301939093526c0100000000000000000000000081049092169281019290925260ff74010000000000000000000000000000000000000000909104161515606082015290565b60408051825163ffffffff16815260208084015167ffffffffffffffff9081169183019190915283830151169181019190915260609182015115159181019190915260800161037c565b610a2661150b565b60405161037c92919061522c565b610582610a4236600461526e565b61160f565b610561611677565b610561610a5d3660046152b2565b611737565b610582610a7036600461531e565b61249a565b6105826124ab565b610582610a8b3660046154cf565b612742565b610582610a9e366004614b4b565b612753565b610ab6610ab1366004615013565b612764565b60405161037c91906155c2565b60007f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168367ffffffffffffffff1614610b43576040517fd9a9cd6800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b6000610b5a610b55608085018561560f565b612798565b519050610b8584610b6e602086018661560f565b905083610b7e6040880188615674565b905061288c565b6000600c81610b9a6080870160608801614b4b565b6001600160a01b0316815260208082019290925260409081016000208151608081018352905463ffffffff81168252640100000000810467ffffffffffffffff908116948301949094526c010000000000000000000000008104909316918101919091527401000000000000000000000000000000000000000090910460ff16151560608201819052909150610c7857610c3a6080850160608601614b4b565b6040517fa7499d200000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b3a565b60065460009081906001600160a01b031663ffdb4b37610c9e6080890160608a01614b4b565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015267ffffffffffffffff8a1660248201526044016040805180830381865afa158015610d09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2d9190615708565b90925090506000808080610d4460408b018b615674565b90501115610d7f57610d738a610d6060808c0160608d01614b4b565b87610d6e60408e018e615674565b61299c565b91945092509050610d9b565b8551610d989063ffffffff16662386f26fc10000615751565b92505b67ffffffffffffffff8a166000908152600b60205260408120805477010000000000000000000000000000000000000000000000900461ffff1615610e1d57610e1a8c6dffffffffffffffffffffffffffff607089901c16610e0060208f018f61560f565b90508e8060400190610e129190615674565b905087612da5565b91505b6000886020015167ffffffffffffffff168563ffffffff1683600001600f9054906101000a900461ffff1661ffff168e8060200190610e5c919061560f565b610e67929150615751565b8454610e88906b010000000000000000000000900463ffffffff168e615768565b610e929190615768565b610e9c9190615768565b610eb6906dffffffffffffffffffffffffffff8a16615751565b610ec09190615751565b9050877bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1683828b6040015167ffffffffffffffff1689610efd9190615751565b610f079190615768565b610f119190615768565b610f1b919061577b565b9a50505050505050505050505b92915050565b610f36612eab565b610f3f81612f07565b50565b600f54600090610f6490600160801b900467ffffffffffffffff16600161579d565b905090565b610f71613106565b610f7b8282613163565b5050565b6007546040517fbbe4f6db0000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152600092169063bbe4f6db90602401602060405180830381865afa158015610fe3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061100791906157be565b9392505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526003546fffffffffffffffffffffffffffffffff8082168352600160801b80830463ffffffff1660208501527401000000000000000000000000000000000000000090920460ff161515938301939093526004548084166060840152049091166080820152610f6490613482565b6110b9613106565b6001600160a01b0381166110f9576040517f232cb97f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611103611677565b90506000811215611140576040517f02075e0000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316036111925761118d6001600160a01b0384168383613534565b505050565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015261118d9083906001600160a01b038616906370a0823190602401602060405180830381865afa1580156111f5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061121991906157db565b6001600160a01b0386169190613534565b6000546001600160a01b0316331480159061125057506002546001600160a01b03163314155b15611287576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c906020015b60405180910390a150565b6112f1613106565b610f7b8282808060200260200160405190810160405280939291908181526020016000905b8282101561134257611333604083028601368190038101906157f4565b81526020019060010190611316565b50505050506135b4565b6001546001600160a01b031633146113a65760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610b3a565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6001600160a01b0381166000908152600e602052604081205467ffffffffffffffff168015801561146357507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615155b15610f28576040517f856c82470000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063856c824790602401602060405180830381865afa1580156114e7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110079190615833565b606060008061151a6008613827565b90508067ffffffffffffffff81111561153557611535614c49565b60405190808252806020026020018201604052801561157a57816020015b60408051808201909152600080825260208201528152602001906001900390816115535790505b50925060005b818110156115ec57600080611596600884613832565b915091506040518060400160405280836001600160a01b031681526020018261ffff168152508684815181106115ce576115ce615850565b60200260200101819052505050806115e590615866565b9050611580565b5050600f5491926c0100000000000000000000000090920463ffffffff16919050565b6000546001600160a01b0316331480159061163557506002546001600160a01b03163314155b1561166c576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f3f600382613850565b600f546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000916bffffffffffffffffffffffff16907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa158015611709573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061172d91906157db565b610f64919061589e565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663397796f76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611797573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117bb91906158be565b156117f2576040517fc148371500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216611832576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005546001600160a01b03163314611876576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168567ffffffffffffffff16146118ef576040517fd9a9cd6800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610b3a565b6000611901610b55608087018761560f565b51905060006119136040870187615674565b9150611931905087611928602089018961560f565b9050848461288c565b8015611a88576000805b82811015611a76576119506040890189615674565b8281811061196057611960615850565b905060400201602001356000036119a3576040517f5cf0444900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d60006119b460408b018b615674565b848181106119c4576119c4615850565b6119da9260206040909202019081019150614b4b565b6001600160a01b031681526020810191909152604001600020547201000000000000000000000000000000000000900460ff1615611a6657611a59611a2260408a018a615674565b83818110611a3257611a32615850565b905060400201803603810190611a4891906158db565b6006546001600160a01b03166139e9565b611a639083615768565b91505b611a6f81615866565b905061193b565b508015611a8657611a8681613b0a565b505b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016611ac26080880160608901614b4b565b6001600160a01b031603611b2657600f8054869190600090611af39084906bffffffffffffffffffffffff16615915565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff160217905550611c45565b6006546001600160a01b03166241e5be611b466080890160608a01614b4b565b60405160e083901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b039182166004820152602481018990527f00000000000000000000000000000000000000000000000000000000000000009091166044820152606401602060405180830381865afa158015611bd2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bf691906157db565b600f8054600090611c169084906bffffffffffffffffffffffff16615915565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055505b600f546bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811691161115611cb2576040517fe5c7a49100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384166000908152600e602052604090205467ffffffffffffffff16158015611d0a57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615155b15611e02576040517f856c82470000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063856c824790602401602060405180830381865afa158015611d8e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611db29190615833565b6001600160a01b0385166000908152600e6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff929092169190911790555b604080516101a08101825267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b03861660208201526000918101611e95611e5b8a8061560f565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613b1792505050565b6001600160a01b03168152602001600f601081819054906101000a900467ffffffffffffffff16611ec59061593a565b825467ffffffffffffffff9182166101009390930a838102908302199091161790925582526020808301879052600060408085018290526001600160a01b038b168252600e90925290812080546060909401939092611f24911661593a565b825467ffffffffffffffff9182166101009390930a83810292021916179091558152602001611f5960808a0160608b01614b4b565b6001600160a01b03168152602001878152602001888060200190611f7d919061560f565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250602001611fc460408a018a615674565b808060200260200160405190810160405280939291908181526020016000905b8282101561201057612001604083028601368190038101906158db565b81526020019060010190611fe4565b505050505081526020018367ffffffffffffffff81111561203357612033614c49565b60405190808252806020026020018201604052801561206657816020015b60608152602001906001900390816120515790505b508152600060209091018190529091505b8281101561241f57600061208e60408a018a615674565b8381811061209e5761209e615850565b9050604002018036038101906120b491906158db565b905060006120c68b8360000151610f7f565b90506001600160a01b038116158061217c57506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa158015612156573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061217a91906158be565b155b156121c15781516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b3a565b6000816001600160a01b0316637a5c972d60405180608001604052808e80600001906121ed919061560f565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166020808301919091526001600160a01b038e1660408084019190915290880151606090920191909152517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526122af9190600401615961565b6000604051808303816000875af11580156122ce573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526122f69190810190615a1e565b83516001600160a01b03166000908152600d602090815260409091205490820151519192506e010000000000000000000000000000900463ffffffff1610156123795782516040517f36f536ca0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b3a565b805161238490613b17565b5060408051606081019091526001600160a01b03831660808201528060a0810160405160208183030381529060405281526020018260000151815260200182602001518152506040516020016123da9190615aaf565b604051602081830303815290604052856101600151858151811061240057612400615850565b60200260200101819052505050508061241890615866565b9050612077565b5061244a817f0000000000000000000000000000000000000000000000000000000000000000613b72565b6101808201526040517fd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd90612480908390615ba0565b60405180910390a16101800151925050505b949350505050565b6124a2613106565b610f3f81613ccd565b6000546001600160a01b031633148015906124d157506002546001600160a01b03163314155b80156124e557506124e3600833613f79565b155b1561251c576040517f195db95800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600f546c01000000000000000000000000900463ffffffff166000819003612570576040517f990e30bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600f546bffffffffffffffffffffffff16818110156125bb576040517f8d0f71d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006125c5611677565b12156125fd576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600061260a6008613827565b905060005b818110156126ff57600080612625600884613832565b9092509050600087612645836bffffffffffffffffffffffff8a16615751565b61264f919061577b565b905061265b8187615cd5565b955061269f6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016846bffffffffffffffffffffffff8416613534565b6040516bffffffffffffffffffffffff821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a2505050806126f890615866565b905061260f565b5050600f80547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff929092169190911790555050565b61274a613106565b610f3f81613f8e565b61275b612eab565b610f3f81614119565b60606040517f9e7177c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051602081019091526000815260008290036127ee5750604080516020810190915267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152610f28565b7f97a657c9000000000000000000000000000000000000000000000000000000006128198385615cfa565b7fffffffff000000000000000000000000000000000000000000000000000000001614612872576040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61287f8260048186615d40565b8101906110079190615d6a565b67ffffffffffffffff84166000908152600b6020526040902080546301000000900463ffffffff168411156129055780546040517f86933789000000000000000000000000000000000000000000000000000000008152630100000090910463ffffffff16600482015260248101859052604401610b3a565b8054670100000000000000900463ffffffff16831115612951576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8054610100900461ffff16821115612995576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b6000808083815b81811015612d985760008787838181106129bf576129bf615850565b9050604002018036038101906129d591906158db565b905060006001600160a01b03166129f08c8360000151610f7f565b6001600160a01b031603612a3e5780516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b3a565b80516001600160a01b03166000908152600d6020908152604091829020825160e081018452905463ffffffff8082168352640100000000820481169383019390935261ffff68010000000000000000820416938201939093526a01000000000000000000008304821660608201526e0100000000000000000000000000008304909116608082015260ff720100000000000000000000000000000000000083048116151560a0830152730100000000000000000000000000000000000000909204909116151560c08201819052612bbc5767ffffffffffffffff8c166000908152600b602052604090208054612b5c90790100000000000000000000000000000000000000000000000000900461ffff16662386f26fc10000615751565b612b669089615768565b8154909850612b9a907b01000000000000000000000000000000000000000000000000000000900463ffffffff1688615dac565b6001820154909750612bb29063ffffffff1687615dac565b9550505050612d88565b604081015160009061ffff1615612cd85760008c6001600160a01b031684600001516001600160a01b031614612c7b5760065484516040517f4ab35b0b0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911690634ab35b0b90602401602060405180830381865afa158015612c50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c749190615dc9565b9050612c7e565b508a5b620186a0836040015161ffff16612cc08660200151847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166141cf90919063ffffffff16565b612cca9190615751565b612cd4919061577b565b9150505b6060820151612ce79088615dac565b9650816080015186612cf99190615dac565b8251909650600090612d189063ffffffff16662386f26fc10000615751565b905080821015612d3757612d2c818a615768565b985050505050612d88565b6000836020015163ffffffff16662386f26fc10000612d569190615751565b905080831115612d7657612d6a818b615768565b99505050505050612d88565b612d80838b615768565b995050505050505b612d9181615866565b90506129a3565b5050955095509592505050565b60008063ffffffff8316612dba608086615751565b612dc687610220615768565b612dd09190615768565b612dda9190615768565b67ffffffffffffffff88166000908152600b6020526040812080549293509171010000000000000000000000000000000000810463ffffffff1690612e3c907501000000000000000000000000000000000000000000900461ffff1685615751565b612e469190615768565b825490915077010000000000000000000000000000000000000000000000900461ffff16612e846dffffffffffffffffffffffffffff8a1683615751565b612e8e9190615751565b612e9e90655af3107a4000615751565b9998505050505050505050565b6000546001600160a01b03163314612f055760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610b3a565b565b60208101516001600160a01b0316612f4b576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516005805473ffffffffffffffffffffffffffffffffffffffff199081166001600160a01b03938416179091556020808401516006805484169185169190911790556040808501516007805490941690851617909255815160e0810183527f0000000000000000000000000000000000000000000000000000000000000000841681527f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff908116928201929092527f00000000000000000000000000000000000000000000000000000000000000008216818401527f000000000000000000000000000000000000000000000000000000000000000090911660608201527f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1660808201527f0000000000000000000000000000000000000000000000000000000000000000831660a08201527f000000000000000000000000000000000000000000000000000000000000000090921660c0830152517ff98cd3cd66fb6e771ed76b7e0209a9c4d3a0eea42e3bce57226c92a0cdb56dff916112de918490615de4565b6000546001600160a01b0316331480159061312c57506002546001600160a01b03163314155b15612f05576040517ffbdb8e5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b825181101561339757600083828151811061318357613183615850565b6020908102919091018101516040805160e0810182528284015163ffffffff908116825282840151811682860190815260608086015161ffff908116858701908152608080890151861693870193845260a0808a0151871691880191825260c0808b01511515918901918252600190890190815299516001600160a01b03166000908152600d909b5297909920955186549451915193519951975198511515730100000000000000000000000000000000000000027fffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffff9915157201000000000000000000000000000000000000027fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff9988166e01000000000000000000000000000002999099167fffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffffff9b88166a0100000000000000000000027fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff969095166801000000000000000002959095167fffffffffffffffffffffffffffffffffffff000000000000ffffffffffffffff938816640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000909716929097169190911794909417169390931792909217959095161791909117919091169190911790555061339081615866565b9050613166565b507ff5791bc457b3bb990493cf5f655db46c25ccf5764c9b99b8969b4c72ea7df9d0826040516133c79190615e91565b60405180910390a160005b815181101561343f57600d60008383815181106133f1576133f1615850565b6020908102919091018101516001600160a01b03168252810191909152604001600020805473ffffffffffffffffffffffffffffffffffffffff1916905561343881615866565b90506133d2565b50805115610f7b577ffb95a0042158e60a33e7b5bec100f3d95407b1a71bee6633bd54b8887449750b8160405161347691906155c2565b60405180910390a15050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261351082606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff16426134f49190615f24565b85608001516fffffffffffffffffffffffffffffffff1661420c565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261118d908490614234565b805160408111156135f1576040517fb5a10cfa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600f546c01000000000000000000000000900463ffffffff161580159061363f5750600f5463ffffffff6c010000000000000000000000008204166bffffffffffffffffffffffff90911610155b1561364c5761364c6124ab565b60006136586008613827565b90505b801561369a576000613679613671600184615f24565b600890613832565b509050613687600882614319565b50508061369390615f37565b905061365b565b506000805b828110156137a85760008482815181106136bb576136bb615850565b602002602001015160000151905060008583815181106136dd576136dd615850565b60200260200101516020015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316148061373257506001600160a01b038216155b15613774576040517f4de938d10000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610b3a565b61378460088361ffff841661432e565b5061379361ffff821685615dac565b93505050806137a190615866565b905061369f565b50600f80547fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd249061381a9083908690615f6c565b60405180910390a1505050565b6000610f2882614344565b6000808080613841868661434f565b909450925050505b9250929050565b815460009061386c90600160801b900463ffffffff1642615f24565b905080156138e957600183015483546138a7916fffffffffffffffffffffffffffffffff808216928116918591600160801b9091041661420c565b83546fffffffffffffffffffffffffffffffff9190911673ffffffffffffffffffffffffffffffffffffffff1990911617600160801b4263ffffffff16021783555b6020820151835461390f916fffffffffffffffffffffffffffffffff908116911661437a565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff9283161717845560208301516040808501518316600160801b0291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c199061381a9084908151151581526020808301516fffffffffffffffffffffffffffffffff90811691830191909152604092830151169181019190915260600190565b81516040517fd02641a00000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009182919084169063d02641a0906024016040805180830381865afa158015613a4f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a739190615f8b565b5190507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116600003613adc5783516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b3a565b6020840151612492907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8316906141cf565b610f3f6003826000614390565b60008151602014613b5657816040517f8d666f60000000000000000000000000000000000000000000000000000000008152600401610b3a9190614bb8565b610f2882806020019051810190613b6d91906157db565b6146ab565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b6101000151604051602001613c089897969594939291906001600160a01b039889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b6040516020818303038152906040528051906020012085610120015180519060200120866101400151604051602001613c419190615fbe565b60405160208183030381529060405280519060200120876101600151604051602001613c6d9190615fd1565b60408051601f198184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b60005b8151811015610f7b576000828281518110613ced57613ced615850565b60200260200101516020015190506000838381518110613d0f57613d0f615850565b60200260200101516000015190506000604051806020016040528084815250905080600b60008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff16021790555050509050508167ffffffffffffffff167f9fda653fbd9a0d9a27a82d8f481a3d81e76169fb54108f481b1da82a0875bdf682604051613f5d9190615030565b60405180910390a250505080613f7290615866565b9050613cd0565b6000611007836001600160a01b038416614717565b60005b81518110156140e9576000828281518110613fae57613fae615850565b60209081029190910181015160408051608080820183528385015163ffffffff90811683528385015167ffffffffffffffff908116848801908152606080880151831686880190815294880151151590860190815296516001600160a01b03166000908152600c9098529490962092518354945192519551151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9688166c0100000000000000000000000002969096167fffffffffffffffffffffff000000000000000000ffffffffffffffffffffffff93909716640100000000027fffffffffffffffffffffffffffffffffffffffff000000000000000000000000909516911617929092179190911692909217179055506140e281615866565b9050613f91565b507f067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e816040516112de9190615fe4565b336001600160a01b038216036141715760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610b3a565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000670de0b6b3a7640000614202837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8616615751565b611007919061577b565b600061422b8561421c8486615751565b6142269087615768565b61437a565b95945050505050565b6000614289826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166147239092919063ffffffff16565b80519091501561118d57808060200190518101906142a791906158be565b61118d5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610b3a565b6000611007836001600160a01b038416614732565b6000612492846001600160a01b0385168461474f565b6000610f288261476c565b6000808061435d8585614776565b600081815260029690960160205260409095205494959350505050565b60008183106143895781611007565b5090919050565b825474010000000000000000000000000000000000000000900460ff1615806143b7575081155b156143c157505050565b825460018401546fffffffffffffffffffffffffffffffff808316929116906000906143fa90600160801b900463ffffffff1642615f24565b905080156144a0578183111561443c576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600186015461446990839085908490600160801b90046fffffffffffffffffffffffffffffffff1661420c565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16600160801b4263ffffffff160217875592505b8482101561453d576001600160a01b0384166144f2576040517ff94ebcd10000000000000000000000000000000000000000000000000000000081526004810183905260248101869052604401610b3a565b6040517f1a76572a00000000000000000000000000000000000000000000000000000000815260048101839052602481018690526001600160a01b0385166044820152606401610b3a565b8483101561462957600186810154600160801b90046fffffffffffffffffffffffffffffffff169060009082906145749082615f24565b61457e878a615f24565b6145889190615768565b614592919061577b565b90506001600160a01b0386166145de576040517f15279c080000000000000000000000000000000000000000000000000000000081526004810182905260248101869052604401610b3a565b6040517fd0c8d23a00000000000000000000000000000000000000000000000000000000815260048101829052602481018690526001600160a01b0387166044820152606401610b3a565b6146338584615f24565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b60006001600160a01b038211806146c25750600a82105b156147135760408051602081018490520160408051601f19818403018152908290527f8d666f60000000000000000000000000000000000000000000000000000000008252610b3a91600401614bb8565b5090565b60006110078383614782565b6060612492848460008561479a565b60008181526002830160205260408120819055611007838361488c565b600082815260028401602052604081208290556124928484614898565b6000610f28825490565b600061100783836148a4565b60008181526001830160205260408120541515611007565b6060824710156148125760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610b3a565b600080866001600160a01b0316858760405161482e9190616065565b60006040518083038185875af1925050503d806000811461486b576040519150601f19603f3d011682016040523d82523d6000602084013e614870565b606091505b5091509150614881878383876148ce565b979650505050505050565b60006110078383614947565b60006110078383614a41565b60008260000182815481106148bb576148bb615850565b9060005260206000200154905092915050565b6060831561493d578251600003614936576001600160a01b0385163b6149365760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610b3a565b5081612492565b6124928383614a90565b60008181526001830160205260408120548015614a3057600061496b600183615f24565b855490915060009061497f90600190615f24565b90508181146149e457600086600001828154811061499f5761499f615850565b90600052602060002001549050808760000184815481106149c2576149c2615850565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806149f5576149f5616081565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610f28565b6000915050610f28565b5092915050565b6000818152600183016020526040812054614a8857508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610f28565b506000610f28565b815115614aa05781518083602001fd5b8060405162461bcd60e51b8152600401610b3a9190614bb8565b60e08101610f2882846001600160a01b03808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015280606085015116606086015250506bffffffffffffffffffffffff60808301511660808401528060a08301511660a08401528060c08301511660c0840152505050565b6001600160a01b0381168114610f3f57600080fd5b600060208284031215614b5d57600080fd5b813561100781614b36565b60005b83811015614b83578181015183820152602001614b6b565b50506000910152565b60008151808452614ba4816020860160208601614b68565b601f01601f19169290920160200192915050565b6020815260006110076020830184614b8c565b67ffffffffffffffff81168114610f3f57600080fd5b600060a08284031215614bf357600080fd5b50919050565b60008060408385031215614c0c57600080fd5b8235614c1781614bcb565b9150602083013567ffffffffffffffff811115614c3357600080fd5b614c3f85828601614be1565b9150509250929050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715614c8257614c82614c49565b60405290565b60405160e0810167ffffffffffffffff81118282101715614c8257614c82614c49565b6040805190810167ffffffffffffffff81118282101715614c8257614c82614c49565b604051610180810167ffffffffffffffff81118282101715614c8257614c82614c49565b60405160a0810167ffffffffffffffff81118282101715614c8257614c82614c49565b604051601f8201601f1916810167ffffffffffffffff81118282101715614d3e57614d3e614c49565b604052919050565b600060608284031215614d5857600080fd5b614d60614c5f565b8235614d6b81614b36565b81526020830135614d7b81614b36565b60208201526040830135614d8e81614b36565b60408201529392505050565b600067ffffffffffffffff821115614db457614db4614c49565b5060051b60200190565b63ffffffff81168114610f3f57600080fd5b8035614ddb81614dbe565b919050565b803561ffff81168114614ddb57600080fd5b8015158114610f3f57600080fd5b8035614ddb81614df2565b600082601f830112614e1c57600080fd5b81356020614e31614e2c83614d9a565b614d15565b82815260059290921b84018101918181019086841115614e5057600080fd5b8286015b84811015614e74578035614e6781614b36565b8352918301918301614e54565b509695505050505050565b6000806040808486031215614e9357600080fd5b833567ffffffffffffffff80821115614eab57600080fd5b818601915086601f830112614ebf57600080fd5b81356020614ecf614e2c83614d9a565b82815260e0928302850182019282820191908b851115614eee57600080fd5b958301955b84871015614f975780878d031215614f0b5760008081fd5b614f13614c88565b8735614f1e81614b36565b815287850135614f2d81614dbe565b8186015287890135614f3e81614dbe565b818a01526060614f4f898201614de0565b90820152608088810135614f6281614dbe565b9082015260a0614f73898201614dd0565b9082015260c0614f84898201614e00565b9082015283529586019591830191614ef3565b5097505087013593505080831115614fae57600080fd5b5050614c3f85828601614e0b565b60008060408385031215614fcf57600080fd5b8235614fda81614bcb565b91506020830135614fea81614b36565b809150509250929050565b6000806040838503121561500857600080fd5b8235614fda81614b36565b60006020828403121561502557600080fd5b813561100781614bcb565b81518051151582526101808201906020810151615053602085018261ffff169052565b50604081015161506b604085018263ffffffff169052565b506060810151615083606085018263ffffffff169052565b50608081015161509b608085018263ffffffff169052565b5060a08101516150b160a085018261ffff169052565b5060c08101516150c960c085018263ffffffff169052565b5060e08101516150df60e085018261ffff169052565b506101008181015161ffff81168583015250506101208181015161ffff81168583015250506101408181015163ffffffff81168583015250506101609081015163ffffffff811684830152905b505092915050565b60608101610f28828480516001600160a01b03908116835260208083015182169084015260409182015116910152565b6000806020838503121561517757600080fd5b823567ffffffffffffffff8082111561518f57600080fd5b818501915085601f8301126151a357600080fd5b8135818111156151b257600080fd5b8660208260061b85010111156151c757600080fd5b60209290920196919550909350505050565b600081518084526020808501945080840160005b8381101561522157815180516001600160a01b0316885283015161ffff1683880152604090960195908201906001016151ed565b509495945050505050565b60408152600061523f60408301856151d9565b90508260208301529392505050565b80356fffffffffffffffffffffffffffffffff81168114614ddb57600080fd5b60006060828403121561528057600080fd5b615288614c5f565b823561529381614df2565b81526152a16020840161524e565b6020820152614d8e6040840161524e565b600080600080608085870312156152c857600080fd5b84356152d381614bcb565b9350602085013567ffffffffffffffff8111156152ef57600080fd5b6152fb87828801614be1565b93505060408501359150606085013561531381614b36565b939692955090935050565b6000602080838503121561533157600080fd5b823567ffffffffffffffff81111561534857600080fd5b8301601f8101851361535957600080fd5b8035615367614e2c82614d9a565b8181526101a0918202830184019184820191908884111561538757600080fd5b938501935b838510156154c357848903818112156153a55760008081fd5b6153ad614cab565b86356153b881614bcb565b8152610180601f1983018113156153cf5760008081fd5b6153d7614cce565b92506153e4898901614e00565b835260406153f3818a01614de0565b8a8501526060615404818b01614dd0565b8286015260809150615417828b01614dd0565b9085015260a06154288a8201614dd0565b8286015260c0915061543b828b01614de0565b9085015260e061544c8a8201614dd0565b828601526101009150615460828b01614de0565b908501526101206154728a8201614de0565b828601526101409150615486828b01614de0565b908501526101606154988a8201614dd0565b828601526154a7838b01614dd0565b908501525050808801919091528352938401939185019161538c565b50979650505050505050565b600060208083850312156154e257600080fd5b823567ffffffffffffffff8111156154f957600080fd5b8301601f8101851361550a57600080fd5b8035615518614e2c82614d9a565b81815260a0918202830184019184820191908884111561553757600080fd5b938501935b838510156154c35780858a0312156155545760008081fd5b61555c614cf2565b853561556781614b36565b81528587013561557681614dbe565b8188015260408681013561558981614bcb565b9082015260608681013561559c81614bcb565b908201526080868101356155af81614df2565b908201528352938401939185019161553c565b6020808252825182820181905260009190848201906040850190845b818110156156035783516001600160a01b0316835292840192918401916001016155de565b50909695505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261564457600080fd5b83018035915067ffffffffffffffff82111561565f57600080fd5b60200191503681900382131561384957600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126156a957600080fd5b83018035915067ffffffffffffffff8211156156c457600080fd5b6020019150600681901b360382131561384957600080fd5b80517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81168114614ddb57600080fd5b6000806040838503121561571b57600080fd5b615724836156dc565b9150615732602084016156dc565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b8082028115828204841417610f2857610f2861573b565b80820180821115610f2857610f2861573b565b60008261579857634e487b7160e01b600052601260045260246000fd5b500490565b67ffffffffffffffff818116838216019080821115614a3a57614a3a61573b565b6000602082840312156157d057600080fd5b815161100781614b36565b6000602082840312156157ed57600080fd5b5051919050565b60006040828403121561580657600080fd5b61580e614cab565b823561581981614b36565b815261582760208401614de0565b60208201529392505050565b60006020828403121561584557600080fd5b815161100781614bcb565b634e487b7160e01b600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036158975761589761573b565b5060010190565b8181036000831280158383131683831282161715614a3a57614a3a61573b565b6000602082840312156158d057600080fd5b815161100781614df2565b6000604082840312156158ed57600080fd5b6158f5614cab565b823561590081614b36565b81526020928301359281019290925250919050565b6bffffffffffffffffffffffff818116838216019080821115614a3a57614a3a61573b565b600067ffffffffffffffff8083168181036159575761595761573b565b6001019392505050565b60208152600082516080602084015261597d60a0840182614b8c565b905067ffffffffffffffff60208501511660408401526001600160a01b036040850151166060840152606084015160808401528091505092915050565b600082601f8301126159cb57600080fd5b815167ffffffffffffffff8111156159e5576159e5614c49565b6159f86020601f19601f84011601614d15565b818152846020838601011115615a0d57600080fd5b612492826020830160208701614b68565b600060208284031215615a3057600080fd5b815167ffffffffffffffff80821115615a4857600080fd5b9083019060408286031215615a5c57600080fd5b615a64614cab565b825182811115615a7357600080fd5b615a7f878286016159ba565b825250602083015182811115615a9457600080fd5b615aa0878286016159ba565b60208301525095945050505050565b602081526000825160606020840152615acb6080840182614b8c565b90506020840151601f1980858403016040860152615ae98383614b8c565b925060408601519150808584030160608601525061422b8282614b8c565b600081518084526020808501945080840160005b8381101561522157815180516001600160a01b031688528301518388015260409096019590820190600101615b1b565b600081518084526020808501808196508360051b8101915082860160005b85811015615b93578284038952615b81848351614b8c565b98850198935090840190600101615b69565b5091979650505050505050565b60208152615bbb60208201835167ffffffffffffffff169052565b60006020830151615bd760408401826001600160a01b03169052565b5060408301516001600160a01b038116606084015250606083015167ffffffffffffffff8116608084015250608083015160a083015260a0830151615c2060c084018215159052565b5060c083015167ffffffffffffffff811660e08401525060e0830151610100615c53818501836001600160a01b03169052565b840151610120848101919091528401516101a061014080860182905291925090615c816101c0860184614b8c565b9250808601519050601f19610160818786030181880152615ca28584615b07565b945080880151925050610180818786030181880152615cc18584615b4b565b970151959092019490945250929392505050565b6bffffffffffffffffffffffff828116828216039080821115614a3a57614a3a61573b565b7fffffffff00000000000000000000000000000000000000000000000000000000813581811691600485101561512c5760049490940360031b84901b1690921692915050565b60008085851115615d5057600080fd5b83861115615d5d57600080fd5b5050820193919092039150565b600060208284031215615d7c57600080fd5b6040516020810181811067ffffffffffffffff82111715615d9f57615d9f614c49565b6040529135825250919050565b63ffffffff818116838216019080821115614a3a57614a3a61573b565b600060208284031215615ddb57600080fd5b611007826156dc565b6101408101615e6182856001600160a01b03808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015280606085015116606086015250506bffffffffffffffffffffffff60808301511660808401528060a08301511660a08401528060c08301511660c0840152505050565b82516001600160a01b0390811660e084015260208401518116610100840152604084015116610120830152611007565b602080825282518282018190526000919060409081850190868401855b82811015615b9357815180516001600160a01b031685528681015163ffffffff908116888701528682015181168787015260608083015161ffff169087015260808083015182169087015260a0808301519091169086015260c09081015115159085015260e09093019290850190600101615eae565b81810381811115610f2857610f2861573b565b600081615f4657615f4661573b565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b63ffffffff8316815260406020820152600061249260408301846151d9565b600060408284031215615f9d57600080fd5b615fa5614cab565b615fae836156dc565b8152602083015161582781614dbe565b6020815260006110076020830184615b07565b6020815260006110076020830184615b4b565b602080825282518282018190526000919060409081850190868401855b82811015615b9357815180516001600160a01b031685528681015163ffffffff16878601528581015167ffffffffffffffff908116878701526060808301519091169086015260809081015115159085015260a09093019290850190600101616001565b60008251616077818460208701614b68565b9190910192915050565b634e487b7160e01b600052603160045260246000fdfea164736f6c6343000813000a",
}

var EVM2EVMMultiOnRampABI = EVM2EVMMultiOnRampMetaData.ABI

var EVM2EVMMultiOnRampBin = EVM2EVMMultiOnRampMetaData.Bin

func DeployEVM2EVMMultiOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMMultiOnRampStaticConfig, dynamicConfig EVM2EVMMultiOnRampDynamicConfig, destChainDynamicConfigsArgs []EVM2EVMMultiOnRampDestChainDynamicConfigArgs, rateLimiterConfig RateLimiterConfig, feeTokenConfigs []EVM2EVMMultiOnRampFeeTokenConfigArgs, tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, nopsAndWeights []EVM2EVMMultiOnRampNopAndWeight) (common.Address, *types.Transaction, *EVM2EVMMultiOnRamp, error) {
	parsed, err := EVM2EVMMultiOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMMultiOnRampBin), backend, staticConfig, dynamicConfig, destChainDynamicConfigsArgs, rateLimiterConfig, feeTokenConfigs, tokenTransferFeeConfigArgs, nopsAndWeights)
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getExpectedNextSequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMMultiOnRamp.CallOpts)
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetNops(opts *bind.CallOpts) (GetNops,

	error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getNops")

	outstruct := new(GetNops)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NopsAndWeights = *abi.ConvertType(out[0], new([]EVM2EVMMultiOnRampNopAndWeight)).(*[]EVM2EVMMultiOnRampNopAndWeight)
	outstruct.WeightsTotal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetNops() (GetNops,

	error) {
	return _EVM2EVMMultiOnRamp.Contract.GetNops(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetNops() (GetNops,

	error) {
	return _EVM2EVMMultiOnRamp.Contract.GetNops(&_EVM2EVMMultiOnRamp.CallOpts)
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getSenderNonce", sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetSenderNonce(&_EVM2EVMMultiOnRamp.CallOpts, sender)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetSenderNonce(&_EVM2EVMMultiOnRamp.CallOpts, sender)
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) LinkAvailableForPayment() (*big.Int, error) {
	return _EVM2EVMMultiOnRamp.Contract.LinkAvailableForPayment(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _EVM2EVMMultiOnRamp.Contract.LinkAvailableForPayment(&_EVM2EVMMultiOnRamp.CallOpts)
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainDynamicConfigsArgs []EVM2EVMMultiOnRampDestChainDynamicConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "applyDestChainConfigUpdates", destChainDynamicConfigsArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) ApplyDestChainConfigUpdates(destChainDynamicConfigsArgs []EVM2EVMMultiOnRampDestChainDynamicConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyDestChainConfigUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, destChainDynamicConfigsArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) ApplyDestChainConfigUpdates(destChainDynamicConfigsArgs []EVM2EVMMultiOnRampDestChainDynamicConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyDestChainConfigUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, destChainDynamicConfigsArgs)
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) SetNops(opts *bind.TransactOpts, nopsAndWeights []EVM2EVMMultiOnRampNopAndWeight) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "setNops", nopsAndWeights)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) SetNops(nopsAndWeights []EVM2EVMMultiOnRampNopAndWeight) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetNops(&_EVM2EVMMultiOnRamp.TransactOpts, nopsAndWeights)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) SetNops(nopsAndWeights []EVM2EVMMultiOnRampNopAndWeight) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.SetNops(&_EVM2EVMMultiOnRamp.TransactOpts, nopsAndWeights)
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) WithdrawNonLinkFees(opts *bind.TransactOpts, feeToken common.Address, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "withdrawNonLinkFees", feeToken, to)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) WithdrawNonLinkFees(feeToken common.Address, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.WithdrawNonLinkFees(&_EVM2EVMMultiOnRamp.TransactOpts, feeToken, to)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) WithdrawNonLinkFees(feeToken common.Address, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.WithdrawNonLinkFees(&_EVM2EVMMultiOnRamp.TransactOpts, feeToken, to)
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

type EVM2EVMMultiOnRampDestChainConfigUpdatedIterator struct {
	Event *EVM2EVMMultiOnRampDestChainConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampDestChainConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampDestChainConfigUpdated)
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
		it.Event = new(EVM2EVMMultiOnRampDestChainConfigUpdated)
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

func (it *EVM2EVMMultiOnRampDestChainConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampDestChainConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampDestChainConfigUpdated struct {
	DestChainSelector uint64
	DestChainConfig   EVM2EVMMultiOnRampDestChainConfig
	Raw               types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterDestChainConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampDestChainConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "DestChainConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampDestChainConfigUpdatedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "DestChainConfigUpdated", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchDestChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampDestChainConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "DestChainConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampDestChainConfigUpdated)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "DestChainConfigUpdated", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseDestChainConfigUpdated(log types.Log) (*EVM2EVMMultiOnRampDestChainConfigUpdated, error) {
	event := new(EVM2EVMMultiOnRampDestChainConfigUpdated)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "DestChainConfigUpdated", log); err != nil {
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

type EVM2EVMMultiOnRampNopPaidIterator struct {
	Event *EVM2EVMMultiOnRampNopPaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampNopPaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampNopPaid)
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
		it.Event = new(EVM2EVMMultiOnRampNopPaid)
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

func (it *EVM2EVMMultiOnRampNopPaidIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampNopPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampNopPaid struct {
	Nop    common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterNopPaid(opts *bind.FilterOpts, nop []common.Address) (*EVM2EVMMultiOnRampNopPaidIterator, error) {

	var nopRule []interface{}
	for _, nopItem := range nop {
		nopRule = append(nopRule, nopItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "NopPaid", nopRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampNopPaidIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "NopPaid", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchNopPaid(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampNopPaid, nop []common.Address) (event.Subscription, error) {

	var nopRule []interface{}
	for _, nopItem := range nop {
		nopRule = append(nopRule, nopItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "NopPaid", nopRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampNopPaid)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "NopPaid", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseNopPaid(log types.Log) (*EVM2EVMMultiOnRampNopPaid, error) {
	event := new(EVM2EVMMultiOnRampNopPaid)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "NopPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampNopsSetIterator struct {
	Event *EVM2EVMMultiOnRampNopsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampNopsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampNopsSet)
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
		it.Event = new(EVM2EVMMultiOnRampNopsSet)
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

func (it *EVM2EVMMultiOnRampNopsSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampNopsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampNopsSet struct {
	NopWeightsTotal *big.Int
	NopsAndWeights  []EVM2EVMMultiOnRampNopAndWeight
	Raw             types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterNopsSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampNopsSetIterator, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "NopsSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampNopsSetIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "NopsSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchNopsSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampNopsSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "NopsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampNopsSet)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "NopsSet", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseNopsSet(log types.Log) (*EVM2EVMMultiOnRampNopsSet, error) {
	event := new(EVM2EVMMultiOnRampNopsSet)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "NopsSet", log); err != nil {
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

type GetNops struct {
	NopsAndWeights []EVM2EVMMultiOnRampNopAndWeight
	WeightsTotal   *big.Int
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMMultiOnRamp.abi.Events["AdminSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseAdminSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["CCIPSendRequested"].ID:
		return _EVM2EVMMultiOnRamp.ParseCCIPSendRequested(log)
	case _EVM2EVMMultiOnRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseConfigSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["DestChainConfigUpdated"].ID:
		return _EVM2EVMMultiOnRamp.ParseDestChainConfigUpdated(log)
	case _EVM2EVMMultiOnRamp.abi.Events["FeeConfigSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseFeeConfigSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["NopPaid"].ID:
		return _EVM2EVMMultiOnRamp.ParseNopPaid(log)
	case _EVM2EVMMultiOnRamp.abi.Events["NopsSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseNopsSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMMultiOnRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMMultiOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMMultiOnRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMMultiOnRamp.abi.Events["TokenTransferFeeConfigDeleted"].ID:
		return _EVM2EVMMultiOnRamp.ParseTokenTransferFeeConfigDeleted(log)
	case _EVM2EVMMultiOnRamp.abi.Events["TokenTransferFeeConfigSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseTokenTransferFeeConfigSet(log)

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

func (EVM2EVMMultiOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf98cd3cd66fb6e771ed76b7e0209a9c4d3a0eea42e3bce57226c92a0cdb56dff")
}

func (EVM2EVMMultiOnRampDestChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x9fda653fbd9a0d9a27a82d8f481a3d81e76169fb54108f481b1da82a0875bdf6")
}

func (EVM2EVMMultiOnRampFeeConfigSet) Topic() common.Hash {
	return common.HexToHash("0x067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e")
}

func (EVM2EVMMultiOnRampNopPaid) Topic() common.Hash {
	return common.HexToHash("0x55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f")
}

func (EVM2EVMMultiOnRampNopsSet) Topic() common.Hash {
	return common.HexToHash("0x8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd24")
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRamp) Address() common.Address {
	return _EVM2EVMMultiOnRamp.address
}

type EVM2EVMMultiOnRampInterface interface {
	CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (EVM2EVMMultiOnRampDestChainConfig, error)

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMMultiOnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetFeeTokenConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMMultiOnRampFeeTokenConfig, error)

	GetNopFeesJuels(opts *bind.CallOpts) (*big.Int, error)

	GetNops(opts *bind.CallOpts) (GetNops,

		error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMMultiOnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMMultiOnRampTokenTransferFeeConfig, error)

	LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainDynamicConfigsArgs []EVM2EVMMultiOnRampDestChainDynamicConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	PayNops(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMMultiOnRampDynamicConfig) (*types.Transaction, error)

	SetFeeTokenConfig(opts *bind.TransactOpts, feeTokenConfigArgs []EVM2EVMMultiOnRampFeeTokenConfigArgs) (*types.Transaction, error)

	SetNops(opts *bind.TransactOpts, nopsAndWeights []EVM2EVMMultiOnRampNopAndWeight) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	SetTokenTransferFeeConfig(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawNonLinkFees(opts *bind.TransactOpts, feeToken common.Address, to common.Address) (*types.Transaction, error)

	FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampAdminSetIterator, error)

	WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAdminSet) (event.Subscription, error)

	ParseAdminSet(log types.Log) (*EVM2EVMMultiOnRampAdminSet, error)

	FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampCCIPSendRequestedIterator, error)

	WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampCCIPSendRequested) (event.Subscription, error)

	ParseCCIPSendRequested(log types.Log) (*EVM2EVMMultiOnRampCCIPSendRequested, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMMultiOnRampConfigSet, error)

	FilterDestChainConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampDestChainConfigUpdatedIterator, error)

	WatchDestChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampDestChainConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigUpdated(log types.Log) (*EVM2EVMMultiOnRampDestChainConfigUpdated, error)

	FilterFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampFeeConfigSetIterator, error)

	WatchFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampFeeConfigSet) (event.Subscription, error)

	ParseFeeConfigSet(log types.Log) (*EVM2EVMMultiOnRampFeeConfigSet, error)

	FilterNopPaid(opts *bind.FilterOpts, nop []common.Address) (*EVM2EVMMultiOnRampNopPaidIterator, error)

	WatchNopPaid(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampNopPaid, nop []common.Address) (event.Subscription, error)

	ParseNopPaid(log types.Log) (*EVM2EVMMultiOnRampNopPaid, error)

	FilterNopsSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampNopsSetIterator, error)

	WatchNopsSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampNopsSet) (event.Subscription, error)

	ParseNopsSet(log types.Log) (*EVM2EVMMultiOnRampNopsSet, error)

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

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
