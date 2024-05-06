// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_onramp

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
	TokenAdminRegistry                common.Address
	DefaultTokenFeeUSDCents           uint16
	DefaultTokenDestGasOverhead       uint32
	DefaultTokenDestBytesOverhead     uint32
	EnforceAllowOutOfOrderDefault     bool
	DefaultAllowOutOfOrder            bool
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
	LinkToken         common.Address
	ChainSelector     uint64
	DestChainSelector uint64
	DefaultTxGasLimit uint64
	MaxNopFeesJuels   *big.Int
	PrevOnRamp        common.Address
	ArmProxy          common.Address
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceAllowOutOfOrderDefault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"defaultAllowOutOfOrder\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadARMSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSendZeroTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetSupportedTokensFunctionalityRemoved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChainSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"expected\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"actual\",\"type\":\"bool\"}],\"name\":\"InvalidExtraArgOutOfOrderValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"}],\"name\":\"InvalidNopAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkBalanceNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeBalanceReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeesToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNopsToPay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NotAFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdminOrNop\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SourceTokenDataTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyNops\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceAllowOutOfOrderDefault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"defaultAllowOutOfOrder\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeConfig\",\"type\":\"tuple[]\"}],\"name\":\"FeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NopPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nopWeightsTotal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"NopsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokenTransferFeeConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"transferFeeConfig\",\"type\":\"tuple[]\"}],\"name\":\"TokenTransferFeeConfigSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceAllowOutOfOrderDefault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"defaultAllowOutOfOrder\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeeTokenConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfig\",\"name\":\"feeTokenConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNopFeesJuels\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNops\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"weightsTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenTransferFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceAllowOutOfOrderDefault\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"defaultAllowOutOfOrder\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"setNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"address[]\"}],\"name\":\"setTokenTransferFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawNonLinkFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b50604051620077a9380380620077a9833981016040819052620000359162001ae3565b8333806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c081620002e8565b50506040805160a081018252602084810180516001600160801b039081168085524263ffffffff169385018490528751151585870181905292518216606086018190529790950151166080909301839052600380546001600160a01b031916909417600160801b9283021760ff60a01b1916600160a01b90910217909255029091176004555085516001600160a01b0316158062000169575060208601516001600160401b0316155b8062000180575060408601516001600160401b0316155b8062000197575060608601516001600160401b0316155b80620001ae575060c08601516001600160a01b0316155b15620001cd576040516306b7c75960e31b815260040160405180910390fd5b60208087015160408089015181517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3948101949094526001600160401b039283169184019190915216606082015230608082015260a00160408051601f198184030181529181528151602092830120608090815288516001600160a01b0390811660e052928901516001600160401b039081166101005291890151821661012052606089015190911660a0908152908801516001600160601b031660c0908152908801518216610140528701511661016052620002aa8562000393565b620002b583620006ee565b604080516000815260208101909152620002d190839062000828565b620002dc8162000a91565b505050505050620021f9565b336001600160a01b03821603620003425760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60e08101516001600160a01b0316620003bf576040516306b7c75960e31b815260040160405180910390fd5b80600560008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160000160146101000a81548161ffff021916908361ffff16021790555060408201518160000160166101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001601a6101000a81548161ffff021916908361ffff160217905550608082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160010160006101000a81548161ffff021916908361ffff16021790555060c08201518160010160026101000a81548161ffff021916908361ffff16021790555060e08201518160010160046101000a8154816001600160a01b0302191690836001600160a01b031602179055506101008201518160010160186101000a81548163ffffffff021916908363ffffffff16021790555061012082015181600101601c6101000a81548163ffffffff021916908363ffffffff1602179055506101408201518160020160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506101608201518160020160146101000a81548161ffff021916908361ffff1602179055506101808201518160020160166101000a81548163ffffffff021916908363ffffffff1602179055506101a082015181600201601a6101000a81548163ffffffff021916908363ffffffff1602179055506101c082015181600201601e6101000a81548160ff0219169083151502179055506101e082015181600201601f6101000a81548160ff0219169083151502179055509050507f528296870c043e937fa44fa60a8a0c7ebf07685bbb857f5e2b054b573961b6226040518060e0016040528060e0516001600160a01b03168152602001610100516001600160401b03168152602001610120516001600160401b0316815260200160a0516001600160401b0316815260200160c0516001600160601b03168152602001610140516001600160a01b03168152602001610160516001600160a01b031681525082604051620006e392919062001d99565b60405180910390a150565b60005b8151811015620007f657600082828151811062000712576200071262001e23565b60209081029190910181015160408051608080820183528385015163ffffffff9081168352838501516001600160401b03908116848801908152606080880151831686880190815294880151151590860190815296516001600160a01b03166000908152600b90985294909620925183549451925195511515600160a01b0260ff60a01b199688166c010000000000000000000000000296909616600160601b600160a81b031993909716640100000000026001600160601b031990951691161792909217919091169290921717905550620007ee8162001e4f565b9050620006f1565b507f067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e81604051620006e3919062001e6b565b60005b8251811015620009a65760008382815181106200084c576200084c62001e23565b6020908102919091018101516040805160e0810182528284015163ffffffff908116825282840151811682860190815260608086015161ffff908116858701908152608080890151861693870193845260a0808a0151871691880191825260c0808b01511515918901918252600190890190815299516001600160a01b03166000908152600c909b5297909920955186549451915193519951975198511515600160981b0260ff60981b19991515600160901b0260ff60901b19998816600160701b029990991664ffffffffff60701b199b88166a01000000000000000000000263ffffffff60501b199690951668010000000000000000029590951665ffffffffffff60401b19938816640100000000026001600160401b031990971692909716919091179490941716939093179290921795909516179190911791909116919091179055506200099e8162001e4f565b90506200082b565b507ff5791bc457b3bb990493cf5f655db46c25ccf5764c9b99b8969b4c72ea7df9d082604051620009d8919062001efa565b60405180910390a160005b815181101562000a4a57600c600083838151811062000a065762000a0662001e23565b6020908102919091018101516001600160a01b0316825281019190915260400160002080546001600160a01b031916905562000a428162001e4f565b9050620009e3565b5080511562000a8d577ffb95a0042158e60a33e7b5bec100f3d95407b1a71bee6633bd54b8887449750b8160405162000a84919062001f8f565b60405180910390a15b5050565b8051604081111562000ab657604051635ad0867d60e11b815260040160405180910390fd5b600e546c01000000000000000000000000900463ffffffff161580159062000b005750600e5463ffffffff6c010000000000000000000000008204166001600160601b0390911610155b1562000b105762000b1062000cb3565b600062000b1e600862000eb3565b90505b801562000b6a57600062000b4462000b3b60018462001fde565b60089062000ec6565b50905062000b5460088262000ee4565b50508062000b629062001ff4565b905062000b21565b506000805b8281101562000c4a57600084828151811062000b8f5762000b8f62001e23565b6020026020010151600001519050600085838151811062000bb45762000bb462001e23565b602002602001015160200151905060e0516001600160a01b0316826001600160a01b0316148062000bec57506001600160a01b038216155b1562000c0f5781604051634de938d160e01b81526004016200008491906200200e565b62000c2160088361ffff841662000f02565b5062000c3261ffff82168562002022565b935050508062000c429062001e4f565b905062000b6f565b50600e805463ffffffff60601b19166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd249062000ca6908390869062002042565b60405180910390a1505050565b6000546001600160a01b0316331480159062000cda57506002546001600160a01b03163314155b801562000cf1575062000cef60083362000f22565b155b1562000d105760405163032bb72b60e31b815260040160405180910390fd5b600e546c01000000000000000000000000900463ffffffff16600081900362000d4c5760405163990e30bf60e01b815260040160405180910390fd5b600e546001600160601b03168181101562000d7a576040516311a1ee3b60e31b815260040160405180910390fd5b600062000d8662000f39565b121562000da657604051631e9acf1760e31b815260040160405180910390fd5b80600062000db5600862000eb3565b905060005b8181101562000e8d5760008062000dd360088462000ec6565b909250905060008762000df0836001600160601b038a16620020af565b62000dfc9190620020c9565b905062000e0a8187620020ec565b60e05190965062000e2f906001600160a01b0316846001600160601b03841662000fce565b6040516001600160601b03821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a25050508062000e859062001e4f565b905062000dba565b5050600e80546001600160601b0319166001600160601b03929092169190911790555050565b600062000ec0826200102b565b92915050565b600080808062000ed7868662001038565b9097909650945050505050565b600062000efb836001600160a01b03841662001065565b9392505050565b600062000f1a846001600160a01b0385168462001084565b949350505050565b600062000efb836001600160a01b038416620010a3565b600e5460e0516040516370a0823160e01b81526000926001600160601b0316916001600160a01b0316906370a082319062000f799030906004016200200e565b602060405180830381865afa15801562000f97573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000fbd91906200210f565b62000fc9919062002129565b905090565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b0390811663a9059cbb60e01b1790915262001026918591620010b116565b505050565b600062000ec08262001182565b600080806200104885856200118d565b600081815260029690960160205260409095205494959350505050565b6000818152600283016020526040812081905562000efb83836200119b565b6000828152600284016020526040812082905562000f1a8484620011a9565b600062000efb8383620011b7565b6040805180820190915260208082527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65649082015260009062001100906001600160a01b038516908490620011c5565b8051909150156200102657808060200190518101906200112191906200214c565b620010265760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840162000084565b600062000ec0825490565b600062000efb8383620011d6565b600062000efb838362001203565b600062000efb83836200130e565b600062000efb83836200135d565b606062000f1a848460008562001375565b6000826000018281548110620011f057620011f062001e23565b9060005260206000200154905092915050565b60008181526001830160205260408120548015620012fc5760006200122a60018362001fde565b8554909150600090620012409060019062001fde565b9050818114620012ac57600086600001828154811062001264576200126462001e23565b90600052602060002001549050808760000184815481106200128a576200128a62001e23565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080620012c057620012c06200216a565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000ec0565b600091505062000ec0565b5092915050565b60006200131c83836200135d565b620013545750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000ec0565b50600062000ec0565b60009081526001919091016020526040902054151590565b606082471015620013d85760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840162000084565b600080866001600160a01b03168587604051620013f69190620021a6565b60006040518083038185875af1925050503d806000811462001435576040519150601f19603f3d011682016040523d82523d6000602084013e6200143a565b606091505b5090925090506200144e8783838762001459565b979650505050505050565b60608315620014cd578251600003620014c5576001600160a01b0385163b620014c55760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640162000084565b508162000f1a565b62000f1a8383815115620014e45781518083602001fd5b8060405162461bcd60e51b8152600401620000849190620021c4565b634e487b7160e01b600052604160045260246000fd5b60405161020081016001600160401b03811182821017156200153c576200153c62001500565b60405290565b60405160a081016001600160401b03811182821017156200153c576200153c62001500565b60405160e081016001600160401b03811182821017156200153c576200153c62001500565b604080519081016001600160401b03811182821017156200153c576200153c62001500565b604051601f8201601f191681016001600160401b0381118282101715620015dc57620015dc62001500565b604052919050565b80516001600160a01b0381168114620015fc57600080fd5b919050565b80516001600160401b0381168114620015fc57600080fd5b805161ffff81168114620015fc57600080fd5b805163ffffffff81168114620015fc57600080fd5b80518015158114620015fc57600080fd5b600061020082840312156200166657600080fd5b6200167062001516565b90506200167d82620015e4565b81526200168d6020830162001619565b6020820152620016a0604083016200162c565b6040820152620016b36060830162001619565b6060820152620016c6608083016200162c565b6080820152620016d960a0830162001619565b60a0820152620016ec60c0830162001619565b60c0820152620016ff60e08301620015e4565b60e0820152610100620017148184016200162c565b90820152610120620017288382016200162c565b908201526101406200173c838201620015e4565b908201526101606200175083820162001619565b90820152610180620017648382016200162c565b908201526101a0620017788382016200162c565b908201526101c06200178c83820162001641565b908201526101e0620017a083820162001641565b9082015292915050565b80516001600160801b0381168114620015fc57600080fd5b600060608284031215620017d557600080fd5b604051606081016001600160401b0381118282101715620017fa57620017fa62001500565b6040529050806200180b8362001641565b81526200181b60208401620017aa565b60208201526200182e60408401620017aa565b60408201525092915050565b60006001600160401b0382111562001856576200185662001500565b5060051b60200190565b600082601f8301126200187257600080fd5b815160206200188b62001885836200183a565b620015b1565b82815260a09283028501820192828201919087851115620018ab57600080fd5b8387015b85811015620019385781818a031215620018c95760008081fd5b620018d362001542565b620018de82620015e4565b8152620018ed8683016200162c565b8682015260406200190081840162001601565b9082015260606200191383820162001601565b9082015260806200192683820162001641565b908201528452928401928101620018af565b5090979650505050505050565b600082601f8301126200195757600080fd5b815160206200196a62001885836200183a565b82815260e092830285018201928282019190878511156200198a57600080fd5b8387015b85811015620019385781818a031215620019a85760008081fd5b620019b262001567565b620019bd82620015e4565b8152620019cc8683016200162c565b868201526040620019df8184016200162c565b908201526060620019f283820162001619565b90820152608062001a058382016200162c565b9082015260a062001a188382016200162c565b9082015260c062001a2b83820162001641565b9082015284529284019281016200198e565b600082601f83011262001a4f57600080fd5b8151602062001a6262001885836200183a565b82815260069290921b8401810191818101908684111562001a8257600080fd5b8286015b8481101562001ad8576040818903121562001aa15760008081fd5b62001aab6200158c565b62001ab682620015e4565b815262001ac585830162001619565b8186015283529183019160400162001a86565b509695505050505050565b6000806000806000808688036103a081121562001aff57600080fd5b60e081121562001b0e57600080fd5b5062001b1962001567565b62001b2488620015e4565b815262001b346020890162001601565b602082015262001b476040890162001601565b604082015262001b5a6060890162001601565b606082015260808801516001600160601b038116811462001b7a57600080fd5b608082015262001b8d60a08901620015e4565b60a082015262001ba060c08901620015e4565b60c0820152955062001bb68860e0890162001652565b945062001bc8886102e08901620017c2565b6103408801519094506001600160401b038082111562001be757600080fd5b62001bf58a838b0162001860565b945061036089015191508082111562001c0d57600080fd5b62001c1b8a838b0162001945565b935061038089015191508082111562001c3357600080fd5b5062001c4289828a0162001a3d565b9150509295509295509295565b6001600160a01b03169052565b62001c6982825162001c4f565b602081015162001c7f602084018261ffff169052565b50604081015162001c98604084018263ffffffff169052565b50606081015162001caf606084018261ffff169052565b50608081015162001cc8608084018263ffffffff169052565b5060a081015162001cdf60a084018261ffff169052565b5060c081015162001cf660c084018261ffff169052565b5060e081015162001d0b60e084018262001c4f565b506101008181015163ffffffff81168483015250506101208181015163ffffffff81168483015250506101408082015162001d498285018262001c4f565b50506101608181015161ffff16908301526101808082015163ffffffff908116918401919091526101a080830151909116908301526101c0808201511515908301526101e0908101511515910152565b82516001600160a01b0390811682526020808501516001600160401b0390811691840191909152604080860151821690840152606080860151909116908301526080808501516001600160601b03169083015260a0848101519091169082015260c0808401516102e083019162001e139084018262001c4f565b5062000efb60e083018462001c5c565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820162001e645762001e6462001e39565b5060010190565b602080825282518282018190526000919060409081850190868401855b8281101562001eed57815180516001600160a01b031685528681015163ffffffff1687860152858101516001600160401b03908116878701526060808301519091169086015260809081015115159085015260a0909301929085019060010162001e88565b5091979650505050505050565b602080825282518282018190526000919060409081850190868401855b8281101562001eed57815180516001600160a01b031685528681015163ffffffff908116888701528682015181168787015260608083015161ffff169087015260808083015182169087015260a0808301519091169086015260c09081015115159085015260e0909301929085019060010162001f17565b6020808252825182820181905260009190848201906040850190845b8181101562001fd25783516001600160a01b03168352928401929184019160010162001fab565b50909695505050505050565b8181038181111562000ec05762000ec062001e39565b60008162002006576200200662001e39565b506000190190565b6001600160a01b0391909116815260200190565b63ffffffff81811683821601908082111562001307576200130762001e39565b6000604080830163ffffffff8616845260208281860152818651808452606087019150828801935060005b81811015620020a157845180516001600160a01b0316845284015161ffff168484015293830193918501916001016200206d565b509098975050505050505050565b808202811582820484141762000ec05762000ec062001e39565b600082620020e757634e487b7160e01b600052601260045260246000fd5b500490565b6001600160601b0382811682821603908082111562001307576200130762001e39565b6000602082840312156200212257600080fd5b5051919050565b818103600083128015838313168383128216171562001307576200130762001e39565b6000602082840312156200215f57600080fd5b62000efb8262001641565b634e487b7160e01b600052603160045260246000fd5b60005b838110156200219d57818101518382015260200162002183565b50506000910152565b60008251620021ba81846020870162002180565b9190910192915050565b6020815260008251806020840152620021e581604085016020870162002180565b601f01601f19169190910160400192915050565b60805160a05160c05160e0516101005161012051610140516101605161549b6200230e600039600081816102b10152818161138d0152613166015260008181610282015281816110bf01528181611106015281816118520152818161189901526131370152600081816101f5015281816108ed0152818161147f01528181611d4a0152818161240e01526130aa0152600081816101c60152818161195b015261307b01526000818161019701528181610df701528181611309015281816116530152818161172c015281816120be01528181612c3b015261304c015260008181610253015281816117e401526131080152600081816102240152818161221701526130d901526000611ef4015261549b6000f3fe608060405234801561001057600080fd5b50600436106101495760003560e01c806306285c691461014e5780631772047e146102f7578063181f5a771461044057806320487ded146104835780634120fccd146104a45780634816f4f7146104b957806348a98aa4146104ce578063546719cd146104ee578063549e946f1461054957806354b714681461055c578063599f643114610575578063704b6c02146105865780637437ff9f1461059957806376f6ae761461072657806379ba50971461073957806383b1f9e514610741578063856c8247146107545780638da5cb5b146107675780639a113c361461076f578063b06d41bc14610857578063c92b28321461086d578063d09dc33914610880578063df0aa9e914610888578063eff7cc481461089b578063f25561fd146108a3578063f2fde38b146108b6578063fbca3b74146108c9575b600080fd5b6102e16040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526040518060e001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160401b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160401b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160401b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160601b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b6040516102ee9190614003565b60405180910390f35b6103d4610305366004614036565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152506001600160a01b03166000908152600c6020908152604091829020825160e081018452905463ffffffff8082168352600160201b820481169383019390935261ffff600160401b82041693820193909352600160501b830482166060820152600160701b8304909116608082015260ff600160901b83048116151560a0830152600160981b909204909116151560c082015290565b6040516102ee9190600060e08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015260c0830151151560c083015292915050565b6104766040518060400160405280601781526020017622ab269922ab26a7b72930b6b81018971a971816b232bb60491b81525081565b6040516102ee91906140a3565b6104966104913660046140e3565b6108e9565b6040519081526020016102ee565b6104ac610c7e565b6040516102ee9190614132565b6104cc6104c73660046142fa565b610ca4565b005b6104e16104dc366004614436565b610cba565b6040516102ee919061446f565b6104f6610d33565b6040516102ee919081516001600160801b03908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b6104cc610557366004614483565b610d99565b600e546001600160601b03166040516102ee91906144a1565b6002546001600160a01b03166104e1565b6104cc610594366004614036565b610ec9565b6107196040805161020081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c081018290526101e08101919091525060408051610200810182526005546001600160a01b03808216835261ffff600160a01b8084048216602086015263ffffffff600160b01b808604821697870197909752600160d01b80860484166060880152600160e01b958690048216608088015260065480851660a0890152620100008104851660c0890152600160201b8104861660e0890152600160c01b8104831661010089015295909504811661012087015260075493841661014087015290830490911661016085015293810484166101808401529081049092166101a082015260ff600160f01b8304811615156101c0830152600160f81b90920490911615156101e082015290565b6040516102ee91906145eb565b6104cc6107343660046145fa565b610f71565b6104cc610fd4565b6104cc61074f36600461466e565b61107e565b6104ac610762366004614036565b611092565b6104e161117c565b61080e61077d366004614036565b604080516080810182526000808252602082018190529181018290526060810191909152506001600160a01b03166000908152600b60209081526040918290208251608081018452905463ffffffff811682526001600160401b03600160201b8204811693830193909352600160601b81049092169281019290925260ff600160a01b909104161515606082015290565b60408051825163ffffffff1681526020808401516001600160401b03908116918301919091528383015116918101919091526060918201511515918101919091526080016102ee565b61085f61118b565b6040516102ee9291906147f5565b6104cc61087b36600461482e565b611285565b6104966112e2565b61049661089636600461489b565b611389565b6104cc611f68565b6104cc6108b1366004614906565b612169565b6104cc6108c4366004614036565b61217a565b6108dc6108d7366004614a04565b61218b565b6040516102ee9190614a21565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160401b0316836001600160401b0316146109485782604051631b3539ad60e31b815260040161093f9190614132565b60405180910390fd5b600061095f61095a6080850185614a6e565b6121fd565b5190506109896109726020850185614a6e565b9050826109826040870187614ab4565b9050612324565b6000600b8161099e6080870160608801614036565b6001600160a01b0316815260208082019290925260409081016000208151608081018352905463ffffffff81168252600160201b81046001600160401b0390811694830194909452600160601b810490931691810191909152600160a01b90910460ff16151560608201819052909150610a3c57610a226080850160608601614036565b60405163053a4ce960e51b815260040161093f919061446f565b6006546000908190600160201b90046001600160a01b031663ffdb4b37610a696080890160608a01614036565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526001600160401b038a1660248201526044016040805180830381865afa158015610abb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610adf9190614b14565b90925090506000808080610af660408b018b614ab4565b90501115610b3057610b24610b1160808b0160608c01614036565b86610b1f60408d018d614ab4565b6123c3565b91945092509050610b4c565b8551610b499063ffffffff16662386f26fc10000614b5d565b92505b60065460009062010000900461ffff1615610b9957610b966001600160701b03607087901c16610b7f60208d018d614a6e565b9050610b8e60408e018e614ab4565b905085612729565b90505b6020878101516005546000926001600160401b039092169163ffffffff871691600160d01b900461ffff1690610bd1908f018f614a6e565b610bdc929150614b5d565b600554610bf690600160b01b900463ffffffff168d614b74565b610c009190614b74565b610c0a9190614b74565b610c1d906001600160701b038916614b5d565b610c279190614b5d565b9050866001600160e01b031682828a604001516001600160401b031688610c4e9190614b5d565b610c589190614b74565b610c629190614b74565b610c6c9190614b87565b99505050505050505050505b92915050565b600e54600090610c9f90600160801b90046001600160401b03166001614ba9565b905090565b610cac6127d9565b610cb6828261282d565b5050565b60075460405163bbe4f6db60e01b81526000916001600160a01b03169063bbe4f6db90610ceb90859060040161446f565b602060405180830381865afa158015610d08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2c9190614bc9565b9392505050565b610d3b613f46565b6040805160a0810182526003546001600160801b038082168352600160801b80830463ffffffff166020850152600160a01b90920460ff161515938301939093526004548084166060840152049091166080820152610c9f90612a75565b610da16127d9565b6001600160a01b038116610dc85760405163232cb97f60e01b815260040160405180910390fd5b6000610dd26112e2565b90506000811215610df557604051620103af60e91b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b031603610e4757610e426001600160a01b0384168383612ae0565b505050565b610e4282846001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401610e77919061446f565b602060405180830381865afa158015610e94573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb89190614be6565b6001600160a01b0386169190612ae0565b610ed161117c565b6001600160a01b0316336001600160a01b031614158015610efd57506002546001600160a01b03163314155b15610f1b576040516307b66ab160e51b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0383161790556040517f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c90610f6690839061446f565b60405180910390a150565b610f796127d9565b610cb68282808060200260200160405190810160405280939291908181526020016000905b82821015610fca57610fbb60408302860136819003810190614bff565b81526020019060010190610f9e565b5050505050612b32565b6001546001600160a01b031633146110275760405162461bcd60e51b815260206004820152601660248201527526bab9ba10313290383937b837b9b2b21037bbb732b960511b604482015260640161093f565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b611086612d32565b61108f81612d85565b50565b6001600160a01b0381166000908152600d60205260408120546001600160401b0316801580156110ea57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615155b15610c785760405163856c824760e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063856c82479061113b90869060040161446f565b602060405180830381865afa158015611158573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2c9190614c3e565b6000546001600160a01b031690565b606060008061119a60086131a0565b9050806001600160401b038111156111b4576111b4614146565b6040519080825280602002602001820160405280156111f957816020015b60408051808201909152600080825260208201528152602001906001900390816111d25790505b50925060005b8181101561126b576000806112156008846131ab565b915091506040518060400160405280836001600160a01b031681526020018261ffff1681525086848151811061124d5761124d614c5b565b602002602001018190525050508061126490614c71565b90506111ff565b5050600e549192600160601b90920463ffffffff16919050565b61128d61117c565b6001600160a01b0316336001600160a01b0316141580156112b957506002546001600160a01b03163314155b156112d7576040516307b66ab160e51b815260040160405180910390fd5b61108f6003826131c9565b600e546040516370a0823160e01b81526000916001600160601b0316906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a082319061133e90309060040161446f565b602060405180830381865afa15801561135b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061137f9190614be6565b610c9f9190614c8a565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663397796f76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061140d9190614caa565b1561142b5760405163c148371560e01b815260040160405180910390fd5b6001600160a01b0382166114525760405163a4ec747960e01b815260040160405180910390fd5b6005546001600160a01b0316331461147d57604051631c0a352960e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160401b0316856001600160401b0316146114d15784604051631b3539ad60e31b815260040161093f9190614132565b60006114e361095a6080870187614a6e565b905060006114f46040870187614ab4565b915061151390506115086020880188614a6e565b845190915083612324565b8015611649576000805b82811015611637576115326040890189614ab4565b8281811061154257611542614c5b565b9050604002016020013560000361156c57604051635cf0444960e01b815260040160405180910390fd5b600c600061157d60408b018b614ab4565b8481811061158d5761158d614c5b565b6115a39260206040909202019081019150614036565b6001600160a01b03168152602081019190915260400160002054600160901b900460ff16156116275761161a6115dc60408a018a614ab4565b838181106115ec576115ec614c5b565b9050604002018036038101906116029190614cc7565b600654600160201b90046001600160a01b0316613303565b6116249083614b74565b91505b61163081614c71565b905061151d565b50801561164757611647816133be565b505b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166116836080880160608901614036565b6001600160a01b0316036116d857600e80548691906000906116af9084906001600160601b0316614d01565b92506101000a8154816001600160601b0302191690836001600160601b031602179055506117d7565b600654600160201b90046001600160a01b03166241e5be6116ff6080890160608a01614036565b60405160e083901b6001600160e01b03191681526001600160a01b039182166004820152602481018990527f00000000000000000000000000000000000000000000000000000000000000009091166044820152606401602060405180830381865afa158015611773573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117979190614be6565b600e80546000906117b29084906001600160601b0316614d01565b92506101000a8154816001600160601b0302191690836001600160601b031602179055505b600e546001600160601b037f00000000000000000000000000000000000000000000000000000000000000008116911611156118265760405163e5c7a49160e01b815260040160405180910390fd5b6001600160a01b0384166000908152600d60205260409020546001600160401b031615801561187d57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615155b156119465760405163856c824760e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063856c8247906118ce90879060040161446f565b602060405180830381865afa1580156118eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061190f9190614c3e565b6001600160a01b0385166000908152600d6020526040902080546001600160401b0319166001600160401b03929092169190911790555b604080516101a0810182526001600160401b037f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b038616602082015260009181016119d861199e8a80614a6e565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506133cb92505050565b6001600160a01b03168152602001600e601081819054906101000a90046001600160401b0316611a0790614d21565b91906101000a8154816001600160401b0302191690836001600160401b0316021790556001600160401b03168152602001846000015181526020016000151581526020018460200151611aab576001600160a01b0387166000908152600d602052604081208054909190611a83906001600160401b0316614d21565b91906101000a8154816001600160401b0302191690836001600160401b031602179055611aae565b60005b6001600160401b03168152602001611acc60808a0160608b01614036565b6001600160a01b03168152602001878152602001888060200190611af09190614a6e565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250602001611b3760408a018a614ab4565b808060200260200160405190810160405280939291908181526020016000905b82821015611b8357611b7460408302860136819003810190614cc7565b81526020019060010190611b57565b50505050508152602001836001600160401b03811115611ba557611ba5614146565b604051908082528060200260200182016040528015611bd857816020015b6060815260200190600190039081611bc35790505b508152600060209091018190529091505b82811015611eed576000611c0060408a018a614ab4565b83818110611c1057611c10614c5b565b905060400201803603810190611c269190614cc7565b90506000611c388b8360000151610cba565b90506001600160a01b0381161580611cbc57506040516301ffc9a760e01b815263aff2afbf60e01b60048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa158015611c96573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cba9190614caa565b155b15611cdd578151604051635f8b555b60e11b815261093f919060040161446f565b6000816001600160a01b0316637a5c972d60405180608001604052808e8060000190611d099190614a6e565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506001600160401b037f0000000000000000000000000000000000000000000000000000000000000000166020808301919091526001600160a01b038e1660408084019190915290880151606090920191909152516001600160e01b031960e084901b168152611db29190600401614d4f565b6000604051808303816000875af1158015611dd1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611df99190810190614e0e565b83516001600160a01b03166000908152600c60209081526040909120549082015151919250600160701b900463ffffffff161015611e4d578251604051631b7a9b6560e11b815261093f919060040161446f565b8051611e58906133cb565b50604080516060810190915280611e72846080830161446f565b6040516020818303038152906040528152602001826000015181526020018260200151815250604051602001611ea89190614e9e565b6040516020818303038152906040528561016001518581518110611ece57611ece614c5b565b602002602001018190525050505080611ee690614c71565b9050611be9565b50611f18817f000000000000000000000000000000000000000000000000000000000000000061340d565b6101808201526040517fd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd90611f4e908390614f92565b60405180910390a16101800151925050505b949350505050565b611f7061117c565b6001600160a01b0316336001600160a01b031614158015611f9c57506002546001600160a01b03163314155b8015611fb05750611fae600833613567565b155b15611fce5760405163032bb72b60e31b815260040160405180910390fd5b600e54600160601b900463ffffffff1660008190036120005760405163990e30bf60e01b815260040160405180910390fd5b600e546001600160601b03168181101561202d576040516311a1ee3b60e31b815260040160405180910390fd5b60006120376112e2565b121561205657604051631e9acf1760e31b815260040160405180910390fd5b80600061206360086131a0565b905060005b818110156121435760008061207e6008846131ab565b9092509050600087612099836001600160601b038a16614b5d565b6120a39190614b87565b90506120af81876150b2565b95506120ee6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016846001600160601b038416612ae0565b826001600160a01b03167f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f8260405161212791906144a1565b60405180910390a25050508061213c90614c71565b9050612068565b5050600e80546001600160601b0319166001600160601b03929092169190911790555050565b6121716127d9565b61108f8161357c565b612182612d32565b61108f816136a3565b60075460408051633adea29560e21b815290516060926001600160a01b03169163eb7a8a549160048083019260009291908290030181865afa1580156121d5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c7891908101906150d2565b612205613f74565b60008290036122565760075461224f907f00000000000000000000000000000000000000000000000000000000000000006001600160401b031690600160f81b900460ff16613746565b9050610c78565b60006122628385615160565b9050636859a83760e01b6001600160e01b03198216016122bc57600061228b8460048188615190565b81019061229891906151ba565b80516007549192506122b391600160f81b900460ff16613746565b92505050610c78565b630e7e230f60e41b6001600160e01b03198216016123045760006122e38460048188615190565b8101906122f091906151fb565b90506122b381600001518260200151613746565b604051632923fee760e11b815260040160405180910390fd5b5092915050565b600654600160c01b900463ffffffff168084111561235f57604051638693378960e01b8152600481018290526024810185905260440161093f565b600654600160e01b900463ffffffff1683111561238f57604051632627e49d60e11b815260040160405180910390fd5b600554600160a01b900461ffff168211156123bd57604051632602b5b560e11b815260040160405180910390fd5b50505050565b6000808083815b8181101561271d5760008787838181106123e6576123e6614c5b565b9050604002018036038101906123fc9190614cc7565b905060006001600160a01b03166124377f00000000000000000000000000000000000000000000000000000000000000008360000151610cba565b6001600160a01b031603612461578051604051635f8b555b60e11b815261093f919060040161446f565b80516001600160a01b03166000908152600c6020908152604091829020825160e081018452905463ffffffff8082168352600160201b820481169383019390935261ffff600160401b82041693820193909352600160501b830482166060820152600160701b8304909116608082015260ff600160901b83048116151560a0830152600160981b909204909116151560c082018190526125675760075461251a90600160a01b900461ffff16662386f26fc10000614b5d565b6125249088614b74565b60075490975061254190600160b01b900463ffffffff1687615227565b60075490965061255e90600160d01b900463ffffffff1686615227565b9450505061270d565b604081015160009061ffff161561265d5760008c6001600160a01b031684600001516001600160a01b031614612615576006548451604051634ab35b0b60e01b8152600160201b9092046001600160a01b031691634ab35b0b916125cd9160040161446f565b602060405180830381865afa1580156125ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061260e9190615244565b9050612618565b508a5b620186a0836040015161ffff166126458660200151846001600160e01b03166137c690919063ffffffff16565b61264f9190614b5d565b6126599190614b87565b9150505b606082015161266c9088615227565b965081608001518661267e9190615227565b825190965060009061269d9063ffffffff16662386f26fc10000614b5d565b9050808210156126bc576126b1818a614b74565b98505050505061270d565b6000836020015163ffffffff16662386f26fc100006126db9190614b5d565b9050808311156126fb576126ef818b614b74565b9950505050505061270d565b612705838b614b74565b995050505050505b61271681614c71565b90506123ca565b50509450945094915050565b60008063ffffffff831661273e608086614b5d565b61274a87610220614b74565b6127549190614b74565b61275e9190614b74565b600554600654919250600091600160e01b90910463ffffffff16906127879061ffff1684614b5d565b6127919190614b74565b60065490915062010000900461ffff166127b46001600160701b03891683614b5d565b6127be9190614b5d565b6127ce90655af3107a4000614b5d565b979650505050505050565b6127e161117c565b6001600160a01b0316336001600160a01b03161415801561280d57506002546001600160a01b03163314155b1561282b57604051637dedc72b60e11b815260040160405180910390fd5b565b60005b825181101561299757600083828151811061284d5761284d614c5b565b6020908102919091018101516040805160e0810182528284015163ffffffff908116825282840151811682860190815260608086015161ffff908116858701908152608080890151861693870193845260a0808a0151871691880191825260c0808b01511515918901918252600190890190815299516001600160a01b03166000908152600c909b5297909920955186549451915193519951975198511515600160981b0260ff60981b19991515600160901b0260ff60901b19998816600160701b029990991664ffffffffff60701b199b8816600160501b0263ffffffff60501b1996909516600160401b029590951665ffffffffffff60401b19938816600160201b026001600160401b0319909716929097169190911794909417169390931792909217959095161791909117919091169190911790555061299081614c71565b9050612830565b507ff5791bc457b3bb990493cf5f655db46c25ccf5764c9b99b8969b4c72ea7df9d0826040516129c7919061525f565b60405180910390a160005b8151811015612a3257600c60008383815181106129f1576129f1614c5b565b6020908102919091018101516001600160a01b0316825281019190915260400160002080546001600160a01b0319169055612a2b81614c71565b90506129d2565b50805115610cb6577ffb95a0042158e60a33e7b5bec100f3d95407b1a71bee6633bd54b8887449750b81604051612a699190614a21565b60405180910390a15050565b612a7d613f46565b612ac582606001516001600160801b031683600001516001600160801b0316846020015163ffffffff1642612ab291906152ff565b85608001516001600160801b03166137ee565b6001600160801b031682525063ffffffff4216602082015290565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610e42908490613816565b80516040811115612b5657604051635ad0867d60e11b815260040160405180910390fd5b600e54600160601b900463ffffffff1615801590612b8d5750600e5463ffffffff600160601b8204166001600160601b0390911610155b15612b9a57612b9a611f68565b6000612ba660086131a0565b90505b8015612be8576000612bc7612bbf6001846152ff565b6008906131ab565b509050612bd56008826138e8565b505080612be190615312565b9050612ba9565b506000805b82811015612cd4576000848281518110612c0957612c09614c5b565b60200260200101516000015190506000858381518110612c2b57612c2b614c5b565b60200260200101516020015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b03161480612c8057506001600160a01b038216155b15612ca05781604051634de938d160e01b815260040161093f919061446f565b612cb060088361ffff84166138fd565b50612cbf61ffff821685615227565b9350505080612ccd90614c71565b9050612bed565b50600e805463ffffffff60601b1916600160601b63ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd2490612d259083908690615329565b60405180910390a1505050565b6000546001600160a01b0316331461282b5760405162461bcd60e51b815260206004820152601660248201527527b7363c9031b0b63630b1363290313c9037bbb732b960511b604482015260640161093f565b60e08101516001600160a01b0316612db0576040516306b7c75960e31b815260040160405180910390fd5b80600560008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160000160146101000a81548161ffff021916908361ffff16021790555060408201518160000160166101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001601a6101000a81548161ffff021916908361ffff160217905550608082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160010160006101000a81548161ffff021916908361ffff16021790555060c08201518160010160026101000a81548161ffff021916908361ffff16021790555060e08201518160010160046101000a8154816001600160a01b0302191690836001600160a01b031602179055506101008201518160010160186101000a81548163ffffffff021916908363ffffffff16021790555061012082015181600101601c6101000a81548163ffffffff021916908363ffffffff1602179055506101408201518160020160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506101608201518160020160146101000a81548161ffff021916908361ffff1602179055506101808201518160020160166101000a81548163ffffffff021916908363ffffffff1602179055506101a082015181600201601a6101000a81548163ffffffff021916908363ffffffff1602179055506101c082015181600201601e6101000a81548160ff0219169083151502179055506101e082015181600201601f6101000a81548160ff0219169083151502179055509050507f528296870c043e937fa44fa60a8a0c7ebf07685bbb857f5e2b054b573961b6226040518060e001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160401b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160401b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160401b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160601b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681525082604051610f66929190615348565b6000610c7882613913565b60008080806131ba868661391e565b909450925050505b9250929050565b81546000906131e590600160801b900463ffffffff16426152ff565b905080156132435760018301548354613217916001600160801b03808216928116918591600160801b909104166137ee565b83546001600160801b03919091166001600160a01b031990911617600160801b4263ffffffff16021783555b60208201518354613260916001600160801b039081169116613949565b835483511515600160a01b02600164ff0000000160801b03199091166001600160801b039283161717845560208301516040808501518316600160801b0291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990612d259084908151151581526020808301516001600160801b0390811691830191909152604092830151169181019190915260600190565b8151604051630681320d60e51b815260009182916001600160a01b0385169163d02641a091613335919060040161446f565b6040805180830381865afa158015613351573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133759190615364565b5190506001600160e01b0381166000036133a5578351604051639a655f7b60e01b815261093f919060040161446f565b6020840151611f60906001600160e01b038316906137c6565b61108f600382600061395f565b600081516020146133f1578160405163046b337b60e51b815260040161093f91906140a3565b610c78828060200190518101906134089190614be6565b613b72565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b61010001516040516020016134a29897969594939291906001600160a01b03988916815296881660208801526001600160401b0395861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b60405160208183030381529060405280519060200120856101200151805190602001208661014001516040516020016134db9190615397565b6040516020818303038152906040528051906020012087610160015160405160200161350791906153aa565b60408051601f198184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b6000610d2c836001600160a01b038416613bc5565b60005b815181101561367357600082828151811061359c5761359c614c5b565b60209081029190910181015160408051608080820183528385015163ffffffff9081168352838501516001600160401b03908116848801908152606080880151831686880190815294880151151590860190815296516001600160a01b03166000908152600b90985294909620925183549451925195511515600160a01b0260ff60a01b19968816600160601b0296909616600160601b600160a81b031993909716600160201b026001600160601b03199095169116179290921791909116929092171790555061366c81614c71565b905061357f565b507f067924bf9277d905a9a4631a06d959bc032ace86b3caa835ae7e403d4f39010e81604051610f6691906153bd565b336001600160a01b038216036136f55760405162461bcd60e51b815260206004820152601760248201527621b0b73737ba103a3930b739b332b9103a379039b2b63360491b604482015260640161093f565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b61374e613f74565b600754600160f01b900460ff1680156137775750600754600160f81b900460ff16151582151514155b156137ae57600754604051636ce6a28960e01b8152600160f81b90910460ff1615156004820152821515602482015260440161093f565b50604080518082019091529182521515602082015290565b6000670de0b6b3a76400006137e4836001600160e01b038616614b5d565b610d2c9190614b87565b600061380d856137fe8486614b5d565b6138089087614b74565b613949565b95945050505050565b600061386b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613bd19092919063ffffffff16565b805190915015610e4257808060200190518101906138899190614caa565b610e425760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161093f565b6000610d2c836001600160a01b038416613be0565b6000611f60846001600160a01b03851684613bfd565b6000610c7882613c1a565b6000808061392c8585613c24565b600081815260029690960160205260409095205494959350505050565b60008183106139585781610d2c565b5090919050565b8254600160a01b900460ff161580613975575081155b1561397f57505050565b825460018401546001600160801b03808316929116906000906139af90600160801b900463ffffffff16426152ff565b90508015613a1b57818311156139d857604051634b92ca1560e11b815260040160405180910390fd5b60018601546139fc90839085908490600160801b90046001600160801b03166137ee565b865463ffffffff60801b1916600160801b4263ffffffff160217875592505b84821015613a73576001600160a01b038416613a545760405163f94ebcd160e01b8152600481018390526024810186905260440161093f565b818585604051630d3b2b9560e11b815260040161093f9392919061543d565b84831015613b1157600186810154600160801b90046001600160801b0316906000908290613aa190826152ff565b613aab878a6152ff565b613ab59190614b74565b613abf9190614b87565b90506001600160a01b038616613af2576040516302a4f38160e31b8152600481018290526024810186905260440161093f565b808587604051636864691d60e11b815260040161093f9392919061543d565b613b1b85846152ff565b86546001600160801b0319166001600160801b0382161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b60006001600160a01b03821180613b895750600a82105b15613bc15760408051602081018490520160408051601f198184030181529082905263046b337b60e51b825261093f916004016140a3565b5090565b6000610d2c8383613c30565b6060611f608484600085613c3c565b60008181526002830160205260408120819055610d2c8383613d0c565b60008281526002840160205260408120829055611f608484613d18565b6000610c78825490565b6000610d2c8383613d24565b6000610d2c8383613d4e565b606082471015613c9d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161093f565b600080866001600160a01b03168587604051613cb9919061545c565b60006040518083038185875af1925050503d8060008114613cf6576040519150601f19603f3d011682016040523d82523d6000602084013e613cfb565b606091505b50915091506127ce87838387613d66565b6000610d2c8383613ddf565b6000610d2c8383613ed2565b6000826000018281548110613d3b57613d3b614c5b565b9060005260206000200154905092915050565b60009081526001919091016020526040902054151590565b60608315613dd5578251600003613dce576001600160a01b0385163b613dce5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161093f565b5081611f60565b611f608383613f1c565b60008181526001830160205260408120548015613ec8576000613e036001836152ff565b8554909150600090613e17906001906152ff565b9050818114613e7c576000866000018281548110613e3757613e37614c5b565b9060005260206000200154905080876000018481548110613e5a57613e5a614c5b565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080613e8d57613e8d615478565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610c78565b6000915050610c78565b6000613ede8383613d4e565b613f1457508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610c78565b506000610c78565b815115613f2c5781518083602001fd5b8060405162461bcd60e51b815260040161093f91906140a3565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b604080518082019091526000808252602082015290565b6001600160a01b03169052565b80516001600160a01b0390811683526020808301516001600160401b0390811691850191909152604080840151821690850152606080840151909116908401526080808301516001600160601b03169084015260a08281015182169084015260c09182015116910152565b60e08101610c788284613f98565b6001600160a01b038116811461108f57600080fd5b803561403181614011565b919050565b60006020828403121561404857600080fd5b8135610d2c81614011565b60005b8381101561406e578181015183820152602001614056565b50506000910152565b6000815180845261408f816020860160208601614053565b601f01601f19169290920160200192915050565b602081526000610d2c6020830184614077565b6001600160401b038116811461108f57600080fd5b600060a082840312156140dd57600080fd5b50919050565b600080604083850312156140f657600080fd5b8235614101816140b6565b915060208301356001600160401b0381111561411c57600080fd5b614128858286016140cb565b9150509250929050565b6001600160401b0391909116815260200190565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b038111828210171561417e5761417e614146565b60405290565b60405161020081016001600160401b038111828210171561417e5761417e614146565b60405160a081016001600160401b038111828210171561417e5761417e614146565b604080519081016001600160401b038111828210171561417e5761417e614146565b604051601f8201601f191681016001600160401b038111828210171561421357614213614146565b604052919050565b60006001600160401b0382111561423457614234614146565b5060051b60200190565b63ffffffff8116811461108f57600080fd5b80356140318161423e565b803561ffff8116811461403157600080fd5b801515811461108f57600080fd5b80356140318161426d565b600082601f83011261429757600080fd5b813560206142ac6142a78361421b565b6141eb565b82815260059290921b840181019181810190868411156142cb57600080fd5b8286015b848110156142ef5780356142e281614011565b83529183019183016142cf565b509695505050505050565b600080604080848603121561430e57600080fd5b83356001600160401b038082111561432557600080fd5b818601915086601f83011261433957600080fd5b813560206143496142a78361421b565b82815260e0928302850182019282820191908b85111561436857600080fd5b958301955b848710156144115780878d0312156143855760008081fd5b61438d61415c565b873561439881614011565b8152878501356143a78161423e565b81860152878901356143b88161423e565b818a015260606143c989820161425b565b908201526080888101356143dc8161423e565b9082015260a06143ed898201614250565b9082015260c06143fe89820161427b565b908201528352958601959183019161436d565b509750508701359350508083111561442857600080fd5b505061412885828601614286565b6000806040838503121561444957600080fd5b8235614454816140b6565b9150602083013561446481614011565b809150509250929050565b6001600160a01b0391909116815260200190565b6000806040838503121561449657600080fd5b823561445481614011565b6001600160601b0391909116815260200190565b6144c0828251613f8b565b60208101516144d5602084018261ffff169052565b5060408101516144ed604084018263ffffffff169052565b506060810151614503606084018261ffff169052565b50608081015161451b608084018263ffffffff169052565b5060a081015161453160a084018261ffff169052565b5060c081015161454760c084018261ffff169052565b5060e081015161455a60e0840182613f8b565b506101008181015163ffffffff81168483015250506101208181015163ffffffff81168483015250506101408082015161459682850182613f8b565b50506101608181015161ffff16908301526101808082015163ffffffff908116918401919091526101a080830151909116908301526101c0808201511515908301526101e080820151801515828501526123bd565b6102008101610c7882846144b5565b6000806020838503121561460d57600080fd5b82356001600160401b038082111561462457600080fd5b818501915085601f83011261463857600080fd5b81358181111561464757600080fd5b8660208260061b850101111561465c57600080fd5b60209290920196919550909350505050565b6000610200828403121561468157600080fd5b614689614184565b61469283614026565b81526146a06020840161425b565b60208201526146b160408401614250565b60408201526146c26060840161425b565b60608201526146d360808401614250565b60808201526146e460a0840161425b565b60a08201526146f560c0840161425b565b60c082015261470660e08401614026565b60e0820152610100614719818501614250565b9082015261012061472b848201614250565b9082015261014061473d848201614026565b9082015261016061474f84820161425b565b90820152610180614761848201614250565b908201526101a0614773848201614250565b908201526101c061478584820161427b565b908201526101e061479784820161427b565b908201529392505050565b600081518084526020808501945080840160005b838110156147ea57815180516001600160a01b0316885283015161ffff1683880152604090960195908201906001016147b6565b509495945050505050565b60408152600061480860408301856147a2565b90508260208301529392505050565b80356001600160801b038116811461403157600080fd5b60006060828403121561484057600080fd5b604051606081016001600160401b038111828210171561486257614862614146565b60405282356148708161426d565b815261487e60208401614817565b602082015261488f60408401614817565b60408201529392505050565b600080600080608085870312156148b157600080fd5b84356148bc816140b6565b935060208501356001600160401b038111156148d757600080fd5b6148e3878288016140cb565b9350506040850135915060608501356148fb81614011565b939692955090935050565b6000602080838503121561491957600080fd5b82356001600160401b0381111561492f57600080fd5b8301601f8101851361494057600080fd5b803561494e6142a78261421b565b81815260a0918202830184019184820191908884111561496d57600080fd5b938501935b838510156149f85780858a03121561498a5760008081fd5b6149926141a7565b853561499d81614011565b8152858701356149ac8161423e565b818801526040868101356149bf816140b6565b908201526060868101356149d2816140b6565b908201526080868101356149e58161426d565b9082015283529384019391850191614972565b50979650505050505050565b600060208284031215614a1657600080fd5b8135610d2c816140b6565b6020808252825182820181905260009190848201906040850190845b81811015614a625783516001600160a01b031683529284019291840191600101614a3d565b50909695505050505050565b6000808335601e19843603018112614a8557600080fd5b8301803591506001600160401b03821115614a9f57600080fd5b6020019150368190038213156131c257600080fd5b6000808335601e19843603018112614acb57600080fd5b8301803591506001600160401b03821115614ae557600080fd5b6020019150600681901b36038213156131c257600080fd5b80516001600160e01b038116811461403157600080fd5b60008060408385031215614b2757600080fd5b614b3083614afd565b9150614b3e60208401614afd565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b8082028115828204841417610c7857610c78614b47565b80820180821115610c7857610c78614b47565b600082614ba457634e487b7160e01b600052601260045260246000fd5b500490565b6001600160401b0381811683821601908082111561231d5761231d614b47565b600060208284031215614bdb57600080fd5b8151610d2c81614011565b600060208284031215614bf857600080fd5b5051919050565b600060408284031215614c1157600080fd5b614c196141c9565b8235614c2481614011565b8152614c326020840161425b565b60208201529392505050565b600060208284031215614c5057600080fd5b8151610d2c816140b6565b634e487b7160e01b600052603260045260246000fd5b600060018201614c8357614c83614b47565b5060010190565b818103600083128015838313168383128216171561231d5761231d614b47565b600060208284031215614cbc57600080fd5b8151610d2c8161426d565b600060408284031215614cd957600080fd5b614ce16141c9565b8235614cec81614011565b81526020928301359281019290925250919050565b6001600160601b0381811683821601908082111561231d5761231d614b47565b60006001600160401b038281166002600160401b03198101614d4557614d45614b47565b6001019392505050565b602081526000825160806020840152614d6b60a0840182614077565b60208501516001600160401b03166040858101919091528501516001600160a01b0316606080860191909152909401516080909301929092525090919050565b600082601f830112614dbc57600080fd5b81516001600160401b03811115614dd557614dd5614146565b614de8601f8201601f19166020016141eb565b818152846020838601011115614dfd57600080fd5b611f60826020830160208701614053565b600060208284031215614e2057600080fd5b81516001600160401b0380821115614e3757600080fd5b9083019060408286031215614e4b57600080fd5b614e536141c9565b825182811115614e6257600080fd5b614e6e87828601614dab565b825250602083015182811115614e8357600080fd5b614e8f87828601614dab565b60208301525095945050505050565b602081526000825160606020840152614eba6080840182614077565b90506020840151601f1980858403016040860152614ed88383614077565b925060408601519150808584030160608601525061380d8282614077565b600081518084526020808501945080840160005b838110156147ea57815180516001600160a01b031688528301518388015260409096019590820190600101614f0a565b600082825180855260208086019550808260051b84010181860160005b84811015614f8557601f19868403018952614f73838351614077565b98840198925090830190600101614f57565b5090979650505050505050565b60208152614fac6020820183516001600160401b03169052565b60006020830151614fc06040840182613f8b565b506040830151614fd36060840182613f8b565b5060608301516001600160401b038116608084015250608083015160a083015260a083015161500660c084018215159052565b5060c08301516001600160401b03811660e08401525060e083015161010061503081850183613f8b565b840151610120848101919091528401516101a06101408086018290529192509061505e6101c0860184614077565b9250808601519050601f1961016081878603018188015261507f8584614ef6565b94508088015192505061018081878603018188015261509e8584614f3a565b970151959092019490945250929392505050565b6001600160601b0382811682821603908082111561231d5761231d614b47565b600060208083850312156150e557600080fd5b82516001600160401b038111156150fb57600080fd5b8301601f8101851361510c57600080fd5b805161511a6142a78261421b565b81815260059190911b8201830190838101908783111561513957600080fd5b928401925b828410156127ce57835161515181614011565b8252928401929084019061513e565b6001600160e01b031981358181169160048510156151885780818660040360031b1b83161692505b505092915050565b600080858511156151a057600080fd5b838611156151ad57600080fd5b5050820193919092039150565b6000602082840312156151cc57600080fd5b604051602081016001600160401b03811182821017156151ee576151ee614146565b6040529135825250919050565b60006040828403121561520d57600080fd5b6152156141c9565b823581526020830135614c328161426d565b63ffffffff81811683821601908082111561231d5761231d614b47565b60006020828403121561525657600080fd5b610d2c82614afd565b602080825282518282018190526000919060409081850190868401855b828110156152f257815180516001600160a01b031685528681015163ffffffff908116888701528682015181168787015260608083015161ffff169087015260808083015182169087015260a0808301519091169086015260c09081015115159085015260e0909301929085019060010161527c565b5091979650505050505050565b81810381811115610c7857610c78614b47565b60008161532157615321614b47565b506000190190565b63ffffffff83168152604060208201526000611f6060408301846147a2565b6102e081016153578285613f98565b610d2c60e08301846144b5565b60006040828403121561537657600080fd5b61537e6141c9565b61538783614afd565b81526020830151614c328161423e565b602081526000610d2c6020830184614ef6565b602081526000610d2c6020830184614f3a565b602080825282518282018190526000919060409081850190868401855b828110156152f257815180516001600160a01b031685528681015163ffffffff1687860152858101516001600160401b03908116878701526060808301519091169086015260809081015115159085015260a090930192908501906001016153da565b92835260208301919091526001600160a01b0316604082015260600190565b6000825161546e818460208701614053565b9190910192915050565b634e487b7160e01b600052603160045260246000fdfea164736f6c6343000813000a",
}

var EVM2EVMOnRampABI = EVM2EVMOnRampMetaData.ABI

var EVM2EVMOnRampBin = EVM2EVMOnRampMetaData.Bin

func DeployEVM2EVMOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMOnRampStaticConfig, dynamicConfig EVM2EVMOnRampDynamicConfig, rateLimiterConfig RateLimiterConfig, feeTokenConfigs []EVM2EVMOnRampFeeTokenConfigArgs, tokenTransferFeeConfigArgs []EVM2EVMOnRampTokenTransferFeeConfigArgs, nopsAndWeights []EVM2EVMOnRampNopAndWeight) (common.Address, *types.Transaction, *EVM2EVMOnRamp, error) {
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
	return address, tx, &EVM2EVMOnRamp{address: address, abi: *parsed, EVM2EVMOnRampCaller: EVM2EVMOnRampCaller{contract: contract}, EVM2EVMOnRampTransactor: EVM2EVMOnRampTransactor{contract: contract}, EVM2EVMOnRampFilterer: EVM2EVMOnRampFilterer{contract: contract}}, nil
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

func (EVM2EVMOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x528296870c043e937fa44fa60a8a0c7ebf07685bbb857f5e2b054b573961b622")
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

func (_EVM2EVMOnRamp *EVM2EVMOnRamp) Address() common.Address {
	return _EVM2EVMOnRamp.address
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

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
