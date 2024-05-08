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

type EVM2EVMMultiOnRampDestChainConfigArgs struct {
	DestChainSelector                 uint64
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadARMSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSendZeroTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChainSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"}],\"name\":\"InvalidNopAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkBalanceNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeBalanceReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeesToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNopsToPay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NotAFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdminOrNop\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SourceTokenDataTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyNops\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeConfig\",\"type\":\"tuple[]\"}],\"name\":\"FeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NopPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nopWeightsTotal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"NopsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokenTransferFeeConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"transferFeeConfig\",\"type\":\"tuple[]\"}],\"name\":\"TokenTransferFeeConfigSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigs\",\"type\":\"tuple[]\"}],\"name\":\"applyDestChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeeTokenConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.FeeTokenConfig\",\"name\":\"feeTokenConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNopFeesJuels\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNops\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"weightsTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenTransferFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"setNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"address[]\"}],\"name\":\"setTokenTransferFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawNonLinkFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b506040516200855238038062008552833981016040819052620000359162001cf8565b8333806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c081620002f4565b50506040805160a081018252602084810180516001600160801b039081168085524263ffffffff169385018490528751151585870181905292518216606086018190529790950151166080909301839052600380546001600160a01b031916909417600160801b9283021760ff60a01b1916600160a01b90910217909255029091176004555086516001600160a01b0316158062000169575060208701516001600160401b0316155b8062000180575060408701516001600160401b0316155b8062000197575060608701516001600160401b0316155b80620001ae575060c08701516001600160a01b0316155b15620001cd576040516306b7c75960e31b815260040160405180910390fd5b6020808801516040808a015181517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3948101949094526001600160401b039283169184019190915216606082015230608082015260a00160408051601f198184030181529181528151602092830120608090815289516001600160a01b0390811660e052928a01516001600160401b0390811661010052918a015182166101205260608a015190911660a0908152908901516001600160601b031660c0908152908901518216610140528801511661016052620002aa866200039f565b620002b585620004b4565b620002c083620007f2565b604080516000815260208101909152620002dc9083906200092c565b620002e78162000b94565b5050505050505062002349565b336001600160a01b038216036200034e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60208101516001600160a01b0316620003cb576040516306b7c75960e31b815260040160405180910390fd5b8051600580546001600160a01b03199081166001600160a01b03938416179091556020808401516006805484169185169190911790556040808501516007805490941690851617909255815160e080820184525184168152610100516001600160401b03908116928201929092526101205182168184015260a08051909216606082015260c080516001600160601b03166080830152610140518516928201929092526101605190931690830152517ff98cd3cd66fb6e771ed76b7e0209a9c4d3a0eea42e3bce57226c92a0cdb56dff91620004a991849062001df3565b60405180910390a150565b60005b8151811015620007ee576000828281518110620004d857620004d862001e88565b602002602001015190506000604051806101800160405280836020015115158152602001836040015161ffff168152602001836060015163ffffffff168152602001836080015163ffffffff1681526020018360a0015163ffffffff1681526020018360c0015161ffff1681526020018360e0015163ffffffff16815260200183610100015161ffff16815260200183610120015161ffff16815260200183610140015161ffff16815260200183610160015163ffffffff16815260200183610180015163ffffffff16815250905080600b600084600001516001600160401b03166001600160401b0316815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff16021790555090505081600001516001600160401b03167fbc576b498b3bc03bf0dd433e6675f56dfbc6c438fc7f5363038bd1aceae1505182604051620007d0919062001e9e565b60405180910390a2505080620007e69062001fb3565b9050620004b7565b5050565b60005b8151811015620008fa57600082828151811062000816576200081662001e88565b60209081029190910181015160408051608080820183528385015163ffffffff9081168352838501516001600160401b03908116848801908152606080880151831686880190815294880151151590860190815296516001600160a01b03166000908152600c90985294909620925183549451925195511515600160a01b0260ff60a01b199688166c010000000000000000000000000296909616600160601b600160a81b031993909716640100000000026001600160601b031990951691161792909217919091169290921717905550620008f28162001fb3565b9050620007f5565b507f067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e81604051620004a9919062001fcf565b60005b825181101562000aaa57600083828151811062000950576200095062001e88565b6020908102919091018101516040805160e0810182528284015163ffffffff908116825282840151811682860190815260608086015161ffff908116858701908152608080890151861693870193845260a0808a0151871691880191825260c0808b01511515918901918252600190890190815299516001600160a01b03166000908152600d909b5297909920955186549451915193519951975198511515600160981b0260ff60981b19991515600160901b0260ff60901b19998816600160701b029990991664ffffffffff60701b199b88166a01000000000000000000000263ffffffff60501b199690951668010000000000000000029590951665ffffffffffff60401b19938816640100000000026001600160401b0319909716929097169190911794909417169390931792909217959095161791909117919091169190911790555062000aa28162001fb3565b90506200092f565b507ff5791bc457b3bb990493cf5f655db46c25ccf5764c9b99b8969b4c72ea7df9d08260405162000adc91906200205e565b60405180910390a160005b815181101562000b4e57600d600083838151811062000b0a5762000b0a62001e88565b6020908102919091018101516001600160a01b0316825281019190915260400160002080546001600160a01b031916905562000b468162001fb3565b905062000ae7565b50805115620007ee577ffb95a0042158e60a33e7b5bec100f3d95407b1a71bee6633bd54b8887449750b8160405162000b889190620020f3565b60405180910390a15050565b8051604081111562000bb957604051635ad0867d60e11b815260040160405180910390fd5b600f546c01000000000000000000000000900463ffffffff161580159062000c035750600f5463ffffffff6c010000000000000000000000008204166001600160601b0390911610155b1562000c135762000c1362000dbe565b600062000c21600862000fbe565b90505b801562000c6d57600062000c4762000c3e60018462002142565b60089062000fd1565b50905062000c5760088262000fef565b50508062000c659062002158565b905062000c24565b506000805b8281101562000d5557600084828151811062000c925762000c9262001e88565b6020026020010151600001519050600085838151811062000cb75762000cb762001e88565b602002602001015160200151905060e0516001600160a01b0316826001600160a01b0316148062000cef57506001600160a01b038216155b1562000d1a57604051634de938d160e01b81526001600160a01b038316600482015260240162000084565b62000d2c60088361ffff84166200100d565b5062000d3d61ffff82168562002172565b935050508062000d4d9062001fb3565b905062000c72565b50600f805463ffffffff60601b19166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd249062000db1908390869062002192565b60405180910390a1505050565b6000546001600160a01b0316331480159062000de557506002546001600160a01b03163314155b801562000dfc575062000dfa6008336200102d565b155b1562000e1b5760405163032bb72b60e31b815260040160405180910390fd5b600f546c01000000000000000000000000900463ffffffff16600081900362000e575760405163990e30bf60e01b815260040160405180910390fd5b600f546001600160601b03168181101562000e85576040516311a1ee3b60e31b815260040160405180910390fd5b600062000e9162001044565b121562000eb157604051631e9acf1760e31b815260040160405180910390fd5b80600062000ec0600862000fbe565b905060005b8181101562000f985760008062000ede60088462000fd1565b909250905060008762000efb836001600160601b038a16620021ff565b62000f07919062002219565b905062000f1581876200223c565b60e05190965062000f3a906001600160a01b0316846001600160601b038416620010d2565b6040516001600160601b03821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a25050508062000f909062001fb3565b905062000ec5565b5050600f80546001600160601b0319166001600160601b03929092169190911790555050565b600062000fcb826200112f565b92915050565b600080808062000fe286866200113c565b9097909650945050505050565b600062001006836001600160a01b03841662001169565b9392505050565b600062001025846001600160a01b0385168462001188565b949350505050565b600062001006836001600160a01b038416620011a7565b600f5460e0516040516370a0823160e01b81523060048201526000926001600160601b0316916001600160a01b0316906370a0823190602401602060405180830381865afa1580156200109b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620010c191906200225f565b620010cd919062002279565b905090565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b0390811663a9059cbb60e01b179091526200112a918591620011b516565b505050565b600062000fcb8262001286565b600080806200114c858562001291565b600081815260029690960160205260409095205494959350505050565b600081815260028301602052604081208190556200100683836200129f565b60008281526002840160205260408120829055620010258484620012ad565b6000620010068383620012bb565b6040805180820190915260208082527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65649082015260009062001204906001600160a01b038516908490620012d4565b8051909150156200112a57808060200190518101906200122591906200229c565b6200112a5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840162000084565b600062000fcb825490565b6000620010068383620012e5565b600062001006838362001312565b60006200100683836200141d565b6000818152600183016020526040812054151562001006565b60606200102584846000856200146f565b6000826000018281548110620012ff57620012ff62001e88565b9060005260206000200154905092915050565b600081815260018301602052604081205480156200140b5760006200133960018362002142565b85549091506000906200134f9060019062002142565b9050818114620013bb57600086600001828154811062001373576200137362001e88565b906000526020600020015490508087600001848154811062001399576200139962001e88565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080620013cf57620013cf620022ba565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000fcb565b600091505062000fcb565b5092915050565b6000818152600183016020526040812054620014665750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000fcb565b50600062000fcb565b606082471015620014d25760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840162000084565b600080866001600160a01b03168587604051620014f09190620022f6565b60006040518083038185875af1925050503d80600081146200152f576040519150601f19603f3d011682016040523d82523d6000602084013e62001534565b606091505b509092509050620015488783838762001553565b979650505050505050565b60608315620015c7578251600003620015bf576001600160a01b0385163b620015bf5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640162000084565b508162001025565b620010258383815115620015de5781518083602001fd5b8060405162461bcd60e51b815260040162000084919062002314565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b0381118282101715620016355762001635620015fa565b60405290565b604051606081016001600160401b0381118282101715620016355762001635620015fa565b6040516101a081016001600160401b0381118282101715620016355762001635620015fa565b60405160a081016001600160401b0381118282101715620016355762001635620015fa565b604080519081016001600160401b0381118282101715620016355762001635620015fa565b604051601f8201601f191681016001600160401b0381118282101715620016fb57620016fb620015fa565b604052919050565b80516001600160a01b03811681146200171b57600080fd5b919050565b80516001600160401b03811681146200171b57600080fd5b600060e082840312156200174b57600080fd5b6200175562001610565b9050620017628262001703565b8152620017726020830162001720565b6020820152620017856040830162001720565b6040820152620017986060830162001720565b606082015260808201516001600160601b0381168114620017b857600080fd5b6080820152620017cb60a0830162001703565b60a0820152620017de60c0830162001703565b60c082015292915050565b600060608284031215620017fc57600080fd5b620018066200163b565b9050620018138262001703565b8152620018236020830162001703565b6020820152620018366040830162001703565b604082015292915050565b60006001600160401b038211156200185d576200185d620015fa565b5060051b60200190565b805180151581146200171b57600080fd5b805161ffff811681146200171b57600080fd5b805163ffffffff811681146200171b57600080fd5b600082601f830112620018b257600080fd5b81516020620018cb620018c58362001841565b620016d0565b8281526101a09283028501820192828201919087851115620018ec57600080fd5b8387015b8581101562001a165781818a0312156200190a5760008081fd5b6200191462001660565b6200191f8262001720565b81526200192e86830162001867565b8682015260406200194181840162001878565b908201526060620019548382016200188b565b908201526080620019678382016200188b565b9082015260a06200197a8382016200188b565b9082015260c06200198d83820162001878565b9082015260e0620019a08382016200188b565b90820152610100620019b483820162001878565b90820152610120620019c883820162001878565b90820152610140620019dc83820162001878565b90820152610160620019f08382016200188b565b9082015261018062001a048382016200188b565b908201528452928401928101620018f0565b5090979650505050505050565b80516001600160801b03811681146200171b57600080fd5b60006060828403121562001a4e57600080fd5b62001a586200163b565b905062001a658262001867565b815262001a756020830162001a23565b6020820152620018366040830162001a23565b600082601f83011262001a9a57600080fd5b8151602062001aad620018c58362001841565b82815260a0928302850182019282820191908785111562001acd57600080fd5b8387015b8581101562001a165781818a03121562001aeb5760008081fd5b62001af562001686565b62001b008262001703565b815262001b0f8683016200188b565b86820152604062001b2281840162001720565b90820152606062001b3583820162001720565b90820152608062001b4883820162001867565b90820152845292840192810162001ad1565b600082601f83011262001b6c57600080fd5b8151602062001b7f620018c58362001841565b82815260e0928302850182019282820191908785111562001b9f57600080fd5b8387015b8581101562001a165781818a03121562001bbd5760008081fd5b62001bc762001610565b62001bd28262001703565b815262001be18683016200188b565b86820152604062001bf48184016200188b565b90820152606062001c0783820162001878565b90820152608062001c1a8382016200188b565b9082015260a062001c2d8382016200188b565b9082015260c062001c4083820162001867565b90820152845292840192810162001ba3565b600082601f83011262001c6457600080fd5b8151602062001c77620018c58362001841565b82815260069290921b8401810191818101908684111562001c9757600080fd5b8286015b8481101562001ced576040818903121562001cb65760008081fd5b62001cc0620016ab565b62001ccb8262001703565b815262001cda85830162001878565b8186015283529183019160400162001c9b565b509695505050505050565b6000806000806000806000610220888a03121562001d1557600080fd5b62001d21898962001738565b965062001d328960e08a01620017e9565b6101408901519096506001600160401b038082111562001d5157600080fd5b62001d5f8b838c01620018a0565b965062001d718b6101608c0162001a3b565b95506101c08a015191508082111562001d8957600080fd5b62001d978b838c0162001a88565b94506101e08a015191508082111562001daf57600080fd5b62001dbd8b838c0162001b5a565b93506102008a015191508082111562001dd557600080fd5b5062001de48a828b0162001c52565b91505092959891949750929550565b82516001600160a01b0390811682526020808501516001600160401b0390811682850152604080870151821681860152606080880151909216918501919091526080808701516001600160601b03169085015260a08087015184169085015260c0808701518416908501528451831660e085015290840151821661010084015283015116610120820152610140810162001006565b634e487b7160e01b600052603260045260246000fd5b8151151581526101808101602083015162001ebf602084018261ffff169052565b50604083015162001ed8604084018263ffffffff169052565b50606083015162001ef1606084018263ffffffff169052565b50608083015162001f0a608084018263ffffffff169052565b5060a083015162001f2160a084018261ffff169052565b5060c083015162001f3a60c084018263ffffffff169052565b5060e083015162001f5160e084018261ffff169052565b506101008381015161ffff9081169184019190915261012080850151909116908301526101408084015163ffffffff908116918401919091526101609384015116929091019190915290565b634e487b7160e01b600052601160045260246000fd5b60006001820162001fc85762001fc862001f9d565b5060010190565b602080825282518282018190526000919060409081850190868401855b828110156200205157815180516001600160a01b031685528681015163ffffffff1687860152858101516001600160401b03908116878701526060808301519091169086015260809081015115159085015260a0909301929085019060010162001fec565b5091979650505050505050565b602080825282518282018190526000919060409081850190868401855b828110156200205157815180516001600160a01b031685528681015163ffffffff908116888701528682015181168787015260608083015161ffff169087015260808083015182169087015260a0808301519091169086015260c09081015115159085015260e090930192908501906001016200207b565b6020808252825182820181905260009190848201906040850190845b81811015620021365783516001600160a01b0316835292840192918401916001016200210f565b50909695505050505050565b8181038181111562000fcb5762000fcb62001f9d565b6000816200216a576200216a62001f9d565b506000190190565b63ffffffff81811683821601908082111562001416576200141662001f9d565b6000604080830163ffffffff8616845260208281860152818651808452606087019150828801935060005b81811015620021f157845180516001600160a01b0316845284015161ffff16848401529383019391850191600101620021bd565b509098975050505050505050565b808202811582820484141762000fcb5762000fcb62001f9d565b6000826200223757634e487b7160e01b600052601260045260246000fd5b500490565b6001600160601b0382811682821603908082111562001416576200141662001f9d565b6000602082840312156200227257600080fd5b5051919050565b818103600083128015838313168383128216171562001416576200141662001f9d565b600060208284031215620022af57600080fd5b620010068262001867565b634e487b7160e01b600052603160045260246000fd5b60005b83811015620022ed578181015183820152602001620022d3565b50506000910152565b600082516200230a818460208701620022d0565b9190910192915050565b602081526000825180602084015262002335816040850160208701620022d0565b601f01601f19169190910160400192915050565b60805160a05160c05160e051610100516101205161014051610160516160fb620024576000396000818161033f0152818161173101526130a70152600081816103100152818161142e0152818161149601528181611cd501528181611d3d015261307f01526000818161027c01528181610abd0152818161186e015281816122250152612ffb01526000818161024c01528181611e0e0152612fc701526000818161021d01528181611138015281816116b001528181611a8801528181611b810152818161264f01528181612fa201526136e30152600081816102dc01528181611c4d015261304b0152600081816102ac015281816127bd01526130220152600061241c01526160fb6000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80637437ff9f11610104578063c92b2832116100a2578063f25561fd11610071578063f25561fd14610a60578063f2fde38b14610a73578063faee138314610a86578063fbca3b7414610a9957600080fd5b8063c92b283214610a2a578063d09dc33914610a3d578063df0aa9e914610a45578063eff7cc4814610a5857600080fd5b8063856c8247116100de578063856c8247146108eb5780638da5cb5b146108fe5780639a113c361461090f578063b06d41bc14610a1457600080fd5b80637437ff9f1461087057806376f6ae76146108d057806379ba5097146108e357600080fd5b806348a98aa41161017157806354b714681161014b57806354b714681461065a578063599f64311461067a5780636def4ce71461068b578063704b6c021461085d57600080fd5b806348a98aa4146105b8578063546719cd146105e3578063549e946f1461064757600080fd5b806320487ded116101ad57806320487ded1461054e57806334d560e41461056f5780634120fccd146105845780634816f4f7146105a557600080fd5b806306285c69146101d45780631772047e14610385578063181f5a7714610505575b600080fd5b61036f6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526040518060e001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b60405161037c9190614b38565b60405180910390f35b610499610393366004614bc9565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152506001600160a01b03166000908152600d6020908152604091829020825160e081018452905463ffffffff8082168352640100000000820481169383019390935261ffff68010000000000000000820416938201939093526a01000000000000000000008304821660608201526e0100000000000000000000000000008304909116608082015260ff720100000000000000000000000000000000000083048116151560a0830152730100000000000000000000000000000000000000909204909116151560c082015290565b60405161037c9190600060e08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015260c0830151151560c083015292915050565b6105416040518060400160405280601c81526020017f45564d3245564d4d756c74694f6e52616d7020312e362e302d6465760000000081525081565b60405161037c9190614c36565b61056161055c366004614c87565b610ab9565b60405190815260200161037c565b61058261057d366004614dd4565b610f24565b005b61058c610f38565b60405167ffffffffffffffff909116815260200161037c565b6105826105b3366004614f08565b610f5f565b6105cb6105c6366004615045565b610f75565b6040516001600160a01b03909116815260200161037c565b6105eb611004565b60405161037c919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b61058261065536600461507e565b6110a7565b600f546040516bffffffffffffffffffffffff909116815260200161037c565b6002546001600160a01b03166105cb565b61085061069936600461509c565b6040805161018081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081018290526101608101919091525067ffffffffffffffff166000908152600b6020908152604091829020825161018081018452815460ff81161515825261010080820461ffff908116958401959095526301000000820463ffffffff908116968401969096526701000000000000008204861660608401526b0100000000000000000000008204861660808401526f010000000000000000000000000000008204851660a0840152710100000000000000000000000000000000008204861660c084015275010000000000000000000000000000000000000000008204851660e084015277010000000000000000000000000000000000000000000000820485169083015279010000000000000000000000000000000000000000000000000081049093166101208201527b0100000000000000000000000000000000000000000000000000000090920483166101408301526001015490911661016082015290565b60405161037c91906150b9565b61058261086b366004614bc9565b611220565b6108c3604080516060810182526000808252602082018190529181019190915250604080516060810182526005546001600160a01b03908116825260065481166020830152600754169181019190915290565b60405161037c91906151b9565b6105826108de3660046151e9565b6112df565b610582611342565b61058c6108f9366004614bc9565b611400565b6000546001600160a01b03166105cb565b6109ca61091d366004614bc9565b604080516080810182526000808252602082018190529181018290526060810191909152506001600160a01b03166000908152600c60209081526040918290208251608081018452905463ffffffff8116825267ffffffffffffffff64010000000082048116938301939093526c0100000000000000000000000081049092169281019290925260ff74010000000000000000000000000000000000000000909104161515606082015290565b60408051825163ffffffff16815260208084015167ffffffffffffffff9081169183019190915283830151169181019190915260609182015115159181019190915260800161037c565b610a1c611501565b60405161037c9291906152b1565b610582610a383660046152f3565b611605565b61056161166d565b610561610a53366004615337565b61172d565b610582612490565b610582610a6e3660046153a3565b612727565b610582610a81366004614bc9565b612738565b610582610a943660046154a2565b612749565b610aac610aa736600461509c565b61275a565b60405161037c9190615619565b60007f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168367ffffffffffffffff1614610b39576040517fd9a9cd6800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b6000610b50610b4b6080850185615666565b61278e565b519050610b7b84610b646020860186615666565b905083610b7460408801886156cb565b9050612882565b6000600c81610b906080870160608801614bc9565b6001600160a01b0316815260208082019290925260409081016000208151608081018352905463ffffffff81168252640100000000810467ffffffffffffffff908116948301949094526c010000000000000000000000008104909316918101919091527401000000000000000000000000000000000000000090910460ff16151560608201819052909150610c6e57610c306080850160608601614bc9565b6040517fa7499d200000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b30565b60065460009081906001600160a01b031663ffdb4b37610c946080890160608a01614bc9565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015267ffffffffffffffff8a1660248201526044016040805180830381865afa158015610cff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d23919061575f565b90925090506000808080610d3a60408b018b6156cb565b90501115610d7557610d698a610d5660808c0160608d01614bc9565b87610d6460408e018e6156cb565b612992565b91945092509050610d91565b8551610d8e9063ffffffff16662386f26fc100006157a8565b92505b67ffffffffffffffff8a166000908152600b60205260408120805477010000000000000000000000000000000000000000000000900461ffff1615610e1357610e108c6dffffffffffffffffffffffffffff607089901c16610df660208f018f615666565b90508e8060400190610e0891906156cb565b905087612d9b565b91505b6000886020015167ffffffffffffffff168563ffffffff1683600001600f9054906101000a900461ffff1661ffff168e8060200190610e529190615666565b610e5d9291506157a8565b8454610e7e906b010000000000000000000000900463ffffffff168e6157bf565b610e8891906157bf565b610e9291906157bf565b610eac906dffffffffffffffffffffffffffff8a166157a8565b610eb691906157a8565b9050877bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1683828b6040015167ffffffffffffffff1689610ef391906157a8565b610efd91906157bf565b610f0791906157bf565b610f1191906157d2565b9a50505050505050505050505b92915050565b610f2c612ea1565b610f3581612efd565b50565b600f54600090610f5a90600160801b900467ffffffffffffffff1660016157f4565b905090565b610f676130fc565b610f718282613159565b5050565b6007546040517fbbe4f6db0000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152600092169063bbe4f6db90602401602060405180830381865afa158015610fd9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ffd9190615815565b9392505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526003546fffffffffffffffffffffffffffffffff8082168352600160801b80830463ffffffff1660208501527401000000000000000000000000000000000000000090920460ff161515938301939093526004548084166060840152049091166080820152610f5a90613478565b6110af6130fc565b6001600160a01b0381166110ef576040517f232cb97f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006110f961166d565b90506000811215611136576040517f02075e0000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b031603611188576111836001600160a01b038416838361352a565b505050565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526111839083906001600160a01b038616906370a0823190602401602060405180830381865afa1580156111eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061120f9190615832565b6001600160a01b038616919061352a565b6000546001600160a01b0316331480159061124657506002546001600160a01b03163314155b1561127d576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c906020015b60405180910390a150565b6112e76130fc565b610f718282808060200260200160405190810160405280939291908181526020016000905b82821015611338576113296040830286013681900381019061584b565b8152602001906001019061130c565b50505050506135aa565b6001546001600160a01b0316331461139c5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610b30565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6001600160a01b0381166000908152600e602052604081205467ffffffffffffffff168015801561145957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615155b15610f1e576040517f856c82470000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063856c824790602401602060405180830381865afa1580156114dd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ffd919061588a565b6060600080611510600861381d565b90508067ffffffffffffffff81111561152b5761152b614cd7565b60405190808252806020026020018201604052801561157057816020015b60408051808201909152600080825260208201528152602001906001900390816115495790505b50925060005b818110156115e25760008061158c600884613828565b915091506040518060400160405280836001600160a01b031681526020018261ffff168152508684815181106115c4576115c46158a7565b60200260200101819052505050806115db906158bd565b9050611576565b5050600f5491926c0100000000000000000000000090920463ffffffff16919050565b6000546001600160a01b0316331480159061162b57506002546001600160a01b03163314155b15611662576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f35600382613846565b600f546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000916bffffffffffffffffffffffff16907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa1580156116ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117239190615832565b610f5a91906158f5565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663397796f76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561178d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117b19190615915565b156117e8576040517fc148371500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216611828576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005546001600160a01b0316331461186c576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168567ffffffffffffffff16146118e5576040517fd9a9cd6800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610b30565b60006118f7610b4b6080870187615666565b519050600061190960408701876156cb565b915061192790508761191e6020890189615666565b90508484612882565b8015611a7e576000805b82811015611a6c5761194660408901896156cb565b82818110611956576119566158a7565b90506040020160200135600003611999576040517f5cf0444900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d60006119aa60408b018b6156cb565b848181106119ba576119ba6158a7565b6119d09260206040909202019081019150614bc9565b6001600160a01b031681526020810191909152604001600020547201000000000000000000000000000000000000900460ff1615611a5c57611a4f611a1860408a018a6156cb565b83818110611a2857611a286158a7565b905060400201803603810190611a3e9190615932565b6006546001600160a01b03166139df565b611a5990836157bf565b91505b611a65816158bd565b9050611931565b508015611a7c57611a7c81613b00565b505b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016611ab86080880160608901614bc9565b6001600160a01b031603611b1c57600f8054869190600090611ae99084906bffffffffffffffffffffffff1661596c565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff160217905550611c3b565b6006546001600160a01b03166241e5be611b3c6080890160608a01614bc9565b60405160e083901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b039182166004820152602481018990527f00000000000000000000000000000000000000000000000000000000000000009091166044820152606401602060405180830381865afa158015611bc8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bec9190615832565b600f8054600090611c0c9084906bffffffffffffffffffffffff1661596c565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055505b600f546bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811691161115611ca8576040517fe5c7a49100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384166000908152600e602052604090205467ffffffffffffffff16158015611d0057507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615155b15611df8576040517f856c82470000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063856c824790602401602060405180830381865afa158015611d84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611da8919061588a565b6001600160a01b0385166000908152600e6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff929092169190911790555b604080516101a08101825267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b03861660208201526000918101611e8b611e518a80615666565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613b0d92505050565b6001600160a01b03168152602001600f601081819054906101000a900467ffffffffffffffff16611ebb90615991565b825467ffffffffffffffff9182166101009390930a838102908302199091161790925582526020808301879052600060408085018290526001600160a01b038b168252600e90925290812080546060909401939092611f1a9116615991565b825467ffffffffffffffff9182166101009390930a83810292021916179091558152602001611f4f60808a0160608b01614bc9565b6001600160a01b03168152602001878152602001888060200190611f739190615666565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250602001611fba60408a018a6156cb565b808060200260200160405190810160405280939291908181526020016000905b8282101561200657611ff760408302860136819003810190615932565b81526020019060010190611fda565b505050505081526020018367ffffffffffffffff81111561202957612029614cd7565b60405190808252806020026020018201604052801561205c57816020015b60608152602001906001900390816120475790505b508152600060209091018190529091505b8281101561241557600061208460408a018a6156cb565b83818110612094576120946158a7565b9050604002018036038101906120aa9190615932565b905060006120bc8b8360000151610f75565b90506001600160a01b038116158061217257506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa15801561214c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121709190615915565b155b156121b75781516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b30565b6000816001600160a01b0316637a5c972d60405180608001604052808e80600001906121e39190615666565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166020808301919091526001600160a01b038e1660408084019190915290880151606090920191909152517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526122a591906004016159b8565b6000604051808303816000875af11580156122c4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526122ec9190810190615a75565b83516001600160a01b03166000908152600d602090815260409091205490820151519192506e010000000000000000000000000000900463ffffffff16101561236f5782516040517f36f536ca0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b30565b805161237a90613b0d565b5060408051606081019091526001600160a01b03831660808201528060a0810160405160208183030381529060405281526020018260000151815260200182602001518152506040516020016123d09190615b06565b60405160208183030381529060405285610160015185815181106123f6576123f66158a7565b60200260200101819052505050508061240e906158bd565b905061206d565b50612440817f0000000000000000000000000000000000000000000000000000000000000000613b68565b6101808201526040517fd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd90612476908390615bf7565b60405180910390a16101800151925050505b949350505050565b6000546001600160a01b031633148015906124b657506002546001600160a01b03163314155b80156124ca57506124c8600833613cc3565b155b15612501576040517f195db95800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600f546c01000000000000000000000000900463ffffffff166000819003612555576040517f990e30bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600f546bffffffffffffffffffffffff16818110156125a0576040517f8d0f71d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006125aa61166d565b12156125e2576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060006125ef600861381d565b905060005b818110156126e45760008061260a600884613828565b909250905060008761262a836bffffffffffffffffffffffff8a166157a8565b61263491906157d2565b90506126408187615d2c565b95506126846001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016846bffffffffffffffffffffffff841661352a565b6040516bffffffffffffffffffffffff821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a2505050806126dd906158bd565b90506125f4565b5050600f80547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff929092169190911790555050565b61272f6130fc565b610f3581613cd8565b612740612ea1565b610f3581613e63565b6127516130fc565b610f3581613f19565b60606040517f9e7177c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051602081019091526000815260008290036127e45750604080516020810190915267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152610f1e565b7f97a657c90000000000000000000000000000000000000000000000000000000061280f8385615d51565b7fffffffff000000000000000000000000000000000000000000000000000000001614612868576040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6128758260048186615d97565b810190610ffd9190615dc1565b67ffffffffffffffff84166000908152600b6020526040902080546301000000900463ffffffff168411156128fb5780546040517f86933789000000000000000000000000000000000000000000000000000000008152630100000090910463ffffffff16600482015260248101859052604401610b30565b8054670100000000000000900463ffffffff16831115612947576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8054610100900461ffff1682111561298b576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b6000808083815b81811015612d8e5760008787838181106129b5576129b56158a7565b9050604002018036038101906129cb9190615932565b905060006001600160a01b03166129e68c8360000151610f75565b6001600160a01b031603612a345780516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b30565b80516001600160a01b03166000908152600d6020908152604091829020825160e081018452905463ffffffff8082168352640100000000820481169383019390935261ffff68010000000000000000820416938201939093526a01000000000000000000008304821660608201526e0100000000000000000000000000008304909116608082015260ff720100000000000000000000000000000000000083048116151560a0830152730100000000000000000000000000000000000000909204909116151560c08201819052612bb25767ffffffffffffffff8c166000908152600b602052604090208054612b5290790100000000000000000000000000000000000000000000000000900461ffff16662386f26fc100006157a8565b612b5c90896157bf565b8154909850612b90907b01000000000000000000000000000000000000000000000000000000900463ffffffff1688615e03565b6001820154909750612ba89063ffffffff1687615e03565b9550505050612d7e565b604081015160009061ffff1615612cce5760008c6001600160a01b031684600001516001600160a01b031614612c715760065484516040517f4ab35b0b0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911690634ab35b0b90602401602060405180830381865afa158015612c46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c6a9190615e20565b9050612c74565b508a5b620186a0836040015161ffff16612cb68660200151847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1661424d90919063ffffffff16565b612cc091906157a8565b612cca91906157d2565b9150505b6060820151612cdd9088615e03565b9650816080015186612cef9190615e03565b8251909650600090612d0e9063ffffffff16662386f26fc100006157a8565b905080821015612d2d57612d22818a6157bf565b985050505050612d7e565b6000836020015163ffffffff16662386f26fc10000612d4c91906157a8565b905080831115612d6c57612d60818b6157bf565b99505050505050612d7e565b612d76838b6157bf565b995050505050505b612d87816158bd565b9050612999565b5050955095509592505050565b60008063ffffffff8316612db06080866157a8565b612dbc876102206157bf565b612dc691906157bf565b612dd091906157bf565b67ffffffffffffffff88166000908152600b6020526040812080549293509171010000000000000000000000000000000000810463ffffffff1690612e32907501000000000000000000000000000000000000000000900461ffff16856157a8565b612e3c91906157bf565b825490915077010000000000000000000000000000000000000000000000900461ffff16612e7a6dffffffffffffffffffffffffffff8a16836157a8565b612e8491906157a8565b612e9490655af3107a40006157a8565b9998505050505050505050565b6000546001600160a01b03163314612efb5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610b30565b565b60208101516001600160a01b0316612f41576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516005805473ffffffffffffffffffffffffffffffffffffffff199081166001600160a01b03938416179091556020808401516006805484169185169190911790556040808501516007805490941690851617909255815160e0810183527f0000000000000000000000000000000000000000000000000000000000000000841681527f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff908116928201929092527f00000000000000000000000000000000000000000000000000000000000000008216818401527f000000000000000000000000000000000000000000000000000000000000000090911660608201527f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1660808201527f0000000000000000000000000000000000000000000000000000000000000000831660a08201527f000000000000000000000000000000000000000000000000000000000000000090921660c0830152517ff98cd3cd66fb6e771ed76b7e0209a9c4d3a0eea42e3bce57226c92a0cdb56dff916112d4918490615e3b565b6000546001600160a01b0316331480159061312257506002546001600160a01b03163314155b15612efb576040517ffbdb8e5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b825181101561338d576000838281518110613179576131796158a7565b6020908102919091018101516040805160e0810182528284015163ffffffff908116825282840151811682860190815260608086015161ffff908116858701908152608080890151861693870193845260a0808a0151871691880191825260c0808b01511515918901918252600190890190815299516001600160a01b03166000908152600d909b5297909920955186549451915193519951975198511515730100000000000000000000000000000000000000027fffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffff9915157201000000000000000000000000000000000000027fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff9988166e01000000000000000000000000000002999099167fffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffffff9b88166a0100000000000000000000027fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff969095166801000000000000000002959095167fffffffffffffffffffffffffffffffffffff000000000000ffffffffffffffff938816640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009097169290971691909117949094171693909317929092179590951617919091179190911691909117905550613386816158bd565b905061315c565b507ff5791bc457b3bb990493cf5f655db46c25ccf5764c9b99b8969b4c72ea7df9d0826040516133bd9190615ee8565b60405180910390a160005b815181101561343557600d60008383815181106133e7576133e76158a7565b6020908102919091018101516001600160a01b03168252810191909152604001600020805473ffffffffffffffffffffffffffffffffffffffff1916905561342e816158bd565b90506133c8565b50805115610f71577ffb95a0042158e60a33e7b5bec100f3d95407b1a71bee6633bd54b8887449750b8160405161346c9190615619565b60405180910390a15050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261350682606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff16426134ea9190615f7b565b85608001516fffffffffffffffffffffffffffffffff1661428a565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526111839084906142b2565b805160408111156135e7576040517fb5a10cfa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600f546c01000000000000000000000000900463ffffffff16158015906136355750600f5463ffffffff6c010000000000000000000000008204166bffffffffffffffffffffffff90911610155b1561364257613642612490565b600061364e600861381d565b90505b801561369057600061366f613667600184615f7b565b600890613828565b50905061367d600882614397565b50508061368990615f8e565b9050613651565b506000805b8281101561379e5760008482815181106136b1576136b16158a7565b602002602001015160000151905060008583815181106136d3576136d36158a7565b60200260200101516020015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316148061372857506001600160a01b038216155b1561376a576040517f4de938d10000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610b30565b61377a60088361ffff84166143ac565b5061378961ffff821685615e03565b9350505080613797906158bd565b9050613695565b50600f80547fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd24906138109083908690615fc3565b60405180910390a1505050565b6000610f1e826143c2565b600080808061383786866143cd565b909450925050505b9250929050565b815460009061386290600160801b900463ffffffff1642615f7b565b905080156138df576001830154835461389d916fffffffffffffffffffffffffffffffff808216928116918591600160801b9091041661428a565b83546fffffffffffffffffffffffffffffffff9190911673ffffffffffffffffffffffffffffffffffffffff1990911617600160801b4263ffffffff16021783555b60208201518354613905916fffffffffffffffffffffffffffffffff90811691166143f8565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff9283161717845560208301516040808501518316600160801b0291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19906138109084908151151581526020808301516fffffffffffffffffffffffffffffffff90811691830191909152604092830151169181019190915260600190565b81516040517fd02641a00000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009182919084169063d02641a0906024016040805180830381865afa158015613a45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a699190615fe2565b5190507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116600003613ad25783516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b30565b6020840151612488907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83169061424d565b610f35600382600061440e565b60008151602014613b4c57816040517f8d666f60000000000000000000000000000000000000000000000000000000008152600401610b309190614c36565b610f1e82806020019051810190613b639190615832565b614729565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b6101000151604051602001613bfe9897969594939291906001600160a01b039889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b6040516020818303038152906040528051906020012085610120015180519060200120866101400151604051602001613c379190616015565b60405160208183030381529060405280519060200120876101600151604051602001613c639190616028565b60408051601f198184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b6000610ffd836001600160a01b038416614795565b60005b8151811015613e33576000828281518110613cf857613cf86158a7565b60209081029190910181015160408051608080820183528385015163ffffffff90811683528385015167ffffffffffffffff908116848801908152606080880151831686880190815294880151151590860190815296516001600160a01b03166000908152600c9098529490962092518354945192519551151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9688166c0100000000000000000000000002969096167fffffffffffffffffffffff000000000000000000ffffffffffffffffffffffff93909716640100000000027fffffffffffffffffffffffffffffffffffffffff00000000000000000000000090951691161792909217919091169290921717905550613e2c816158bd565b9050613cdb565b507f067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e816040516112d4919061603b565b336001600160a01b03821603613ebb5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610b30565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60005b8151811015610f71576000828281518110613f3957613f396158a7565b602002602001015190506000604051806101800160405280836020015115158152602001836040015161ffff168152602001836060015163ffffffff168152602001836080015163ffffffff1681526020018360a0015163ffffffff1681526020018360c0015161ffff1681526020018360e0015163ffffffff16815260200183610100015161ffff16815260200183610120015161ffff16815260200183610140015161ffff16815260200183610160015163ffffffff16815260200183610180015163ffffffff16815250905080600b6000846000015167ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff160217905550905050816000015167ffffffffffffffff167fbc576b498b3bc03bf0dd433e6675f56dfbc6c438fc7f5363038bd1aceae150518260405161423291906150b9565b60405180910390a2505080614246906158bd565b9050613f1c565b6000670de0b6b3a7640000614280837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff86166157a8565b610ffd91906157d2565b60006142a98561429a84866157a8565b6142a490876157bf565b6143f8565b95945050505050565b6000614307826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166147a19092919063ffffffff16565b80519091501561118357808060200190518101906143259190615915565b6111835760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610b30565b6000610ffd836001600160a01b0384166147b0565b6000612488846001600160a01b038516846147cd565b6000610f1e826147ea565b600080806143db85856147f4565b600081815260029690960160205260409095205494959350505050565b60008183106144075781610ffd565b5090919050565b825474010000000000000000000000000000000000000000900460ff161580614435575081155b1561443f57505050565b825460018401546fffffffffffffffffffffffffffffffff8083169291169060009061447890600160801b900463ffffffff1642615f7b565b9050801561451e57818311156144ba576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018601546144e790839085908490600160801b90046fffffffffffffffffffffffffffffffff1661428a565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16600160801b4263ffffffff160217875592505b848210156145bb576001600160a01b038416614570576040517ff94ebcd10000000000000000000000000000000000000000000000000000000081526004810183905260248101869052604401610b30565b6040517f1a76572a00000000000000000000000000000000000000000000000000000000815260048101839052602481018690526001600160a01b0385166044820152606401610b30565b848310156146a757600186810154600160801b90046fffffffffffffffffffffffffffffffff169060009082906145f29082615f7b565b6145fc878a615f7b565b61460691906157bf565b61461091906157d2565b90506001600160a01b03861661465c576040517f15279c080000000000000000000000000000000000000000000000000000000081526004810182905260248101869052604401610b30565b6040517fd0c8d23a00000000000000000000000000000000000000000000000000000000815260048101829052602481018690526001600160a01b0387166044820152606401610b30565b6146b18584615f7b565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b60006001600160a01b038211806147405750600a82105b156147915760408051602081018490520160408051601f19818403018152908290527f8d666f60000000000000000000000000000000000000000000000000000000008252610b3091600401614c36565b5090565b6000610ffd8383614800565b60606124888484600085614818565b60008181526002830160205260408120819055610ffd838361490a565b600082815260028401602052604081208290556124888484614916565b6000610f1e825490565b6000610ffd8383614922565b60008181526001830160205260408120541515610ffd565b6060824710156148905760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610b30565b600080866001600160a01b031685876040516148ac91906160bc565b60006040518083038185875af1925050503d80600081146148e9576040519150601f19603f3d011682016040523d82523d6000602084013e6148ee565b606091505b50915091506148ff8783838761494c565b979650505050505050565b6000610ffd83836149c5565b6000610ffd8383614abf565b6000826000018281548110614939576149396158a7565b9060005260206000200154905092915050565b606083156149bb5782516000036149b4576001600160a01b0385163b6149b45760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610b30565b5081612488565b6124888383614b0e565b60008181526001830160205260408120548015614aae5760006149e9600183615f7b565b85549091506000906149fd90600190615f7b565b9050818114614a62576000866000018281548110614a1d57614a1d6158a7565b9060005260206000200154905080876000018481548110614a4057614a406158a7565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080614a7357614a736160d8565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610f1e565b6000915050610f1e565b5092915050565b6000818152600183016020526040812054614b0657508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610f1e565b506000610f1e565b815115614b1e5781518083602001fd5b8060405162461bcd60e51b8152600401610b309190614c36565b60e08101610f1e82846001600160a01b03808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015280606085015116606086015250506bffffffffffffffffffffffff60808301511660808401528060a08301511660a08401528060c08301511660c0840152505050565b6001600160a01b0381168114610f3557600080fd5b600060208284031215614bdb57600080fd5b8135610ffd81614bb4565b60005b83811015614c01578181015183820152602001614be9565b50506000910152565b60008151808452614c22816020860160208601614be6565b601f01601f19169290920160200192915050565b602081526000610ffd6020830184614c0a565b67ffffffffffffffff81168114610f3557600080fd5b8035614c6a81614c49565b919050565b600060a08284031215614c8157600080fd5b50919050565b60008060408385031215614c9a57600080fd5b8235614ca581614c49565b9150602083013567ffffffffffffffff811115614cc157600080fd5b614ccd85828601614c6f565b9150509250929050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715614d1057614d10614cd7565b60405290565b60405160e0810167ffffffffffffffff81118282101715614d1057614d10614cd7565b60405160a0810167ffffffffffffffff81118282101715614d1057614d10614cd7565b6040516101a0810167ffffffffffffffff81118282101715614d1057614d10614cd7565b6040805190810167ffffffffffffffff81118282101715614d1057614d10614cd7565b604051601f8201601f1916810167ffffffffffffffff81118282101715614dcc57614dcc614cd7565b604052919050565b600060608284031215614de657600080fd5b614dee614ced565b8235614df981614bb4565b81526020830135614e0981614bb4565b60208201526040830135614e1c81614bb4565b60408201529392505050565b600067ffffffffffffffff821115614e4257614e42614cd7565b5060051b60200190565b63ffffffff81168114610f3557600080fd5b8035614c6a81614e4c565b803561ffff81168114614c6a57600080fd5b8015158114610f3557600080fd5b8035614c6a81614e7b565b600082601f830112614ea557600080fd5b81356020614eba614eb583614e28565b614da3565b82815260059290921b84018101918181019086841115614ed957600080fd5b8286015b84811015614efd578035614ef081614bb4565b8352918301918301614edd565b509695505050505050565b6000806040808486031215614f1c57600080fd5b833567ffffffffffffffff80821115614f3457600080fd5b818601915086601f830112614f4857600080fd5b81356020614f58614eb583614e28565b82815260e0928302850182019282820191908b851115614f7757600080fd5b958301955b848710156150205780878d031215614f945760008081fd5b614f9c614d16565b8735614fa781614bb4565b815287850135614fb681614e4c565b8186015287890135614fc781614e4c565b818a01526060614fd8898201614e69565b90820152608088810135614feb81614e4c565b9082015260a0614ffc898201614e5e565b9082015260c061500d898201614e89565b9082015283529586019591830191614f7c565b509750508701359350508083111561503757600080fd5b5050614ccd85828601614e94565b6000806040838503121561505857600080fd5b823561506381614c49565b9150602083013561507381614bb4565b809150509250929050565b6000806040838503121561509157600080fd5b823561506381614bb4565b6000602082840312156150ae57600080fd5b8135610ffd81614c49565b815115158152610180810160208301516150d9602084018261ffff169052565b5060408301516150f1604084018263ffffffff169052565b506060830151615109606084018263ffffffff169052565b506080830151615121608084018263ffffffff169052565b5060a083015161513760a084018261ffff169052565b5060c083015161514f60c084018263ffffffff169052565b5060e083015161516560e084018261ffff169052565b506101008381015161ffff81168483015250506101208381015161ffff81168483015250506101408381015163ffffffff81168483015250506101608381015163ffffffff8116848301525b505092915050565b60608101610f1e828480516001600160a01b03908116835260208083015182169084015260409182015116910152565b600080602083850312156151fc57600080fd5b823567ffffffffffffffff8082111561521457600080fd5b818501915085601f83011261522857600080fd5b81358181111561523757600080fd5b8660208260061b850101111561524c57600080fd5b60209290920196919550909350505050565b600081518084526020808501945080840160005b838110156152a657815180516001600160a01b0316885283015161ffff168388015260409096019590820190600101615272565b509495945050505050565b6040815260006152c4604083018561525e565b90508260208301529392505050565b80356fffffffffffffffffffffffffffffffff81168114614c6a57600080fd5b60006060828403121561530557600080fd5b61530d614ced565b823561531881614e7b565b8152615326602084016152d3565b6020820152614e1c604084016152d3565b6000806000806080858703121561534d57600080fd5b843561535881614c49565b9350602085013567ffffffffffffffff81111561537457600080fd5b61538087828801614c6f565b93505060408501359150606085013561539881614bb4565b939692955090935050565b600060208083850312156153b657600080fd5b823567ffffffffffffffff8111156153cd57600080fd5b8301601f810185136153de57600080fd5b80356153ec614eb582614e28565b81815260a0918202830184019184820191908884111561540b57600080fd5b938501935b838510156154965780858a0312156154285760008081fd5b615430614d39565b853561543b81614bb4565b81528587013561544a81614e4c565b8188015260408681013561545d81614c49565b9082015260608681013561547081614c49565b9082015260808681013561548381614e7b565b9082015283529384019391850191615410565b50979650505050505050565b600060208083850312156154b557600080fd5b823567ffffffffffffffff8111156154cc57600080fd5b8301601f810185136154dd57600080fd5b80356154eb614eb582614e28565b8181526101a0918202830184019184820191908884111561550b57600080fd5b938501935b838510156154965780858a0312156155285760008081fd5b615530614d5c565b61553986614c5f565b8152615546878701614e89565b878201526040615557818801614e69565b908201526060615568878201614e5e565b908201526080615579878201614e5e565b9082015260a061558a878201614e5e565b9082015260c061559b878201614e69565b9082015260e06155ac878201614e5e565b908201526101006155be878201614e69565b908201526101206155d0878201614e69565b908201526101406155e2878201614e69565b908201526101606155f4878201614e5e565b90820152610180615606878201614e5e565b9082015283529384019391850191615510565b6020808252825182820181905260009190848201906040850190845b8181101561565a5783516001600160a01b031683529284019291840191600101615635565b50909695505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261569b57600080fd5b83018035915067ffffffffffffffff8211156156b657600080fd5b60200191503681900382131561383f57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261570057600080fd5b83018035915067ffffffffffffffff82111561571b57600080fd5b6020019150600681901b360382131561383f57600080fd5b80517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81168114614c6a57600080fd5b6000806040838503121561577257600080fd5b61577b83615733565b915061578960208401615733565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b8082028115828204841417610f1e57610f1e615792565b80820180821115610f1e57610f1e615792565b6000826157ef57634e487b7160e01b600052601260045260246000fd5b500490565b67ffffffffffffffff818116838216019080821115614ab857614ab8615792565b60006020828403121561582757600080fd5b8151610ffd81614bb4565b60006020828403121561584457600080fd5b5051919050565b60006040828403121561585d57600080fd5b615865614d80565b823561587081614bb4565b815261587e60208401614e69565b60208201529392505050565b60006020828403121561589c57600080fd5b8151610ffd81614c49565b634e487b7160e01b600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036158ee576158ee615792565b5060010190565b8181036000831280158383131683831282161715614ab857614ab8615792565b60006020828403121561592757600080fd5b8151610ffd81614e7b565b60006040828403121561594457600080fd5b61594c614d80565b823561595781614bb4565b81526020928301359281019290925250919050565b6bffffffffffffffffffffffff818116838216019080821115614ab857614ab8615792565b600067ffffffffffffffff8083168181036159ae576159ae615792565b6001019392505050565b6020815260008251608060208401526159d460a0840182614c0a565b905067ffffffffffffffff60208501511660408401526001600160a01b036040850151166060840152606084015160808401528091505092915050565b600082601f830112615a2257600080fd5b815167ffffffffffffffff811115615a3c57615a3c614cd7565b615a4f6020601f19601f84011601614da3565b818152846020838601011115615a6457600080fd5b612488826020830160208701614be6565b600060208284031215615a8757600080fd5b815167ffffffffffffffff80821115615a9f57600080fd5b9083019060408286031215615ab357600080fd5b615abb614d80565b825182811115615aca57600080fd5b615ad687828601615a11565b825250602083015182811115615aeb57600080fd5b615af787828601615a11565b60208301525095945050505050565b602081526000825160606020840152615b226080840182614c0a565b90506020840151601f1980858403016040860152615b408383614c0a565b92506040860151915080858403016060860152506142a98282614c0a565b600081518084526020808501945080840160005b838110156152a657815180516001600160a01b031688528301518388015260409096019590820190600101615b72565b600081518084526020808501808196508360051b8101915082860160005b85811015615bea578284038952615bd8848351614c0a565b98850198935090840190600101615bc0565b5091979650505050505050565b60208152615c1260208201835167ffffffffffffffff169052565b60006020830151615c2e60408401826001600160a01b03169052565b5060408301516001600160a01b038116606084015250606083015167ffffffffffffffff8116608084015250608083015160a083015260a0830151615c7760c084018215159052565b5060c083015167ffffffffffffffff811660e08401525060e0830151610100615caa818501836001600160a01b03169052565b840151610120848101919091528401516101a061014080860182905291925090615cd86101c0860184614c0a565b9250808601519050601f19610160818786030181880152615cf98584615b5e565b945080880151925050610180818786030181880152615d188584615ba2565b970151959092019490945250929392505050565b6bffffffffffffffffffffffff828116828216039080821115614ab857614ab8615792565b7fffffffff0000000000000000000000000000000000000000000000000000000081358181169160048510156151b15760049490940360031b84901b1690921692915050565b60008085851115615da757600080fd5b83861115615db457600080fd5b5050820193919092039150565b600060208284031215615dd357600080fd5b6040516020810181811067ffffffffffffffff82111715615df657615df6614cd7565b6040529135825250919050565b63ffffffff818116838216019080821115614ab857614ab8615792565b600060208284031215615e3257600080fd5b610ffd82615733565b6101408101615eb882856001600160a01b03808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015280606085015116606086015250506bffffffffffffffffffffffff60808301511660808401528060a08301511660a08401528060c08301511660c0840152505050565b82516001600160a01b0390811660e084015260208401518116610100840152604084015116610120830152610ffd565b602080825282518282018190526000919060409081850190868401855b82811015615bea57815180516001600160a01b031685528681015163ffffffff908116888701528682015181168787015260608083015161ffff169087015260808083015182169087015260a0808301519091169086015260c09081015115159085015260e09093019290850190600101615f05565b81810381811115610f1e57610f1e615792565b600081615f9d57615f9d615792565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b63ffffffff83168152604060208201526000612488604083018461525e565b600060408284031215615ff457600080fd5b615ffc614d80565b61600583615733565b8152602083015161587e81614e4c565b602081526000610ffd6020830184615b5e565b602081526000610ffd6020830184615ba2565b602080825282518282018190526000919060409081850190868401855b82811015615bea57815180516001600160a01b031685528681015163ffffffff16878601528581015167ffffffffffffffff908116878701526060808301519091169086015260809081015115159085015260a09093019290850190600101616058565b600082516160ce818460208701614be6565b9190910192915050565b634e487b7160e01b600052603160045260246000fdfea164736f6c6343000813000a",
}

var EVM2EVMMultiOnRampABI = EVM2EVMMultiOnRampMetaData.ABI

var EVM2EVMMultiOnRampBin = EVM2EVMMultiOnRampMetaData.Bin

func DeployEVM2EVMMultiOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMMultiOnRampStaticConfig, dynamicConfig EVM2EVMMultiOnRampDynamicConfig, destChainConfigs []EVM2EVMMultiOnRampDestChainConfigArgs, rateLimiterConfig RateLimiterConfig, feeTokenConfigs []EVM2EVMMultiOnRampFeeTokenConfigArgs, tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, nopsAndWeights []EVM2EVMMultiOnRampNopAndWeight) (common.Address, *types.Transaction, *EVM2EVMMultiOnRamp, error) {
	parsed, err := EVM2EVMMultiOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMMultiOnRampBin), backend, staticConfig, dynamicConfig, destChainConfigs, rateLimiterConfig, feeTokenConfigs, tokenTransferFeeConfigArgs, nopsAndWeights)
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) ApplyDestChainConfigUpdates(destChainConfigs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyDestChainConfigUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, destChainConfigs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) ApplyDestChainConfigUpdates(destChainConfigs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyDestChainConfigUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, destChainConfigs)
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
	return common.HexToHash("0xbc576b498b3bc03bf0dd433e6675f56dfbc6c438fc7f5363038bd1aceae15051")
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

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error)

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
