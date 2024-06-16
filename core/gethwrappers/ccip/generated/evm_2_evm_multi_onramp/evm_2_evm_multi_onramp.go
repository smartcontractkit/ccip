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
	GasMultiplierWeiPerEth            uint64
	NetworkFeeUSDCents                uint32
}

type EVM2EVMMultiOnRampDynamicConfig struct {
	Router        common.Address
	PriceRegistry common.Address
}

type EVM2EVMMultiOnRampNopAndWeight struct {
	Nop    common.Address
	Weight uint16
}

type EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthArgs struct {
	Token                      common.Address
	PremiumMultiplierWeiPerEth uint64
}

type EVM2EVMMultiOnRampStaticConfig struct {
	LinkToken          common.Address
	ChainSelector      uint64
	MaxNopFeesJuels    *big.Int
	RmnProxy           common.Address
	NonceManager       common.Address
	TokenAdminRegistry common.Address
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
	DestChainSelector       uint64
	TokenTransferFeeConfigs []EVM2EVMMultiOnRampTokenTransferFeeConfigSingleTokenArgs
}

type EVM2EVMMultiOnRampTokenTransferFeeConfigRemoveArgs struct {
	DestChainSelector uint64
	Token             common.Address
}

type EVM2EVMMultiOnRampTokenTransferFeeConfigSingleTokenArgs struct {
	Token                  common.Address
	TokenTransferFeeConfig EVM2EVMMultiOnRampTokenTransferFeeConfig
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSendZeroTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"DestinationChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChainSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"}],\"name\":\"InvalidDestBytesOverhead\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidDestChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"}],\"name\":\"InvalidNopAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkBalanceNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeBalanceReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeesToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNopsToPay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NotAFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdminOrNop\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SourceTokenDataTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyNops\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainDynamicConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NopPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nopWeightsTotal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"NopsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"name\":\"PremiumMultiplierWeiPerEthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destChainSelector\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenTransferFeeConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"name\":\"TokenTransferFeeConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensConsumed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyDestChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyPremiumMultiplierWeiPerEthUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigRemoveArgs[]\",\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\"}],\"name\":\"applyTokenTransferFeeConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestChainConfig\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNopFeesJuels\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNops\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"weightsTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPremiumMultiplierWeiPerEth\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenTransferFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"setNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawNonLinkFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b5060405162008bc138038062008bc18339810160408190526200003591620020c8565b8333806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c081620002a4565b50506040805160a081018252602084810180516001600160801b039081168085524263ffffffff169385018490528751151585870181905292518216606086018190529790950151166080909301839052600380546001600160a01b031916909417600160801b9283021760ff60a01b1916600160a01b90910217909255029091176004555086516001600160a01b0316158062000169575060208701516001600160401b0316155b8062000180575060608701516001600160a01b0316155b8062000197575060808701516001600160a01b0316155b80620001ae575060a08701516001600160a01b0316155b15620001cd576040516306b7c75960e31b815260040160405180910390fd5b86516001600160a01b0390811660a090815260208901516001600160401b031660c05260408901516001600160601b0316608090815260608a0151831660e052890151821661010052880151166101205262000229866200034f565b6200023485620004b5565b6200023f8362000a12565b604080516000808252602082019092526200028c9184919062000285565b60408051808201909152600080825260208201528152602001906001900390816200025d5790505b5062000ade565b620002978162000e3b565b50505050505050620026c0565b336001600160a01b03821603620002fe5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60208101516001600160a01b03166200037b576040516306b7c75960e31b815260040160405180910390fd5b8051600580546001600160a01b039283166001600160a01b0319918216179091556020808401516006805491851691909316179091556040805160c0808201835260a080518616835290516001600160401b031693820193909352608080516001600160601b03168284015260e05185166060830152610100518516908201526101205190931691830191909152517fbbf7a36713bfcf4a7869fbf276c760a1c96d364781b5516aa806b8dc1e7362e491620004aa91849082516001600160a01b0390811682526020808501516001600160401b0316818401526040808601516001600160601b03169084015260608086015183169084015260808086015183169084015260a0948501518216948301949094528251811660c083015291909201511660e08201526101000190565b60405180910390a150565b60005b815181101562000a0e576000828281518110620004d957620004d9620021c3565b602002602001015190506000838381518110620004fa57620004fa620021c3565b6020026020010151600001519050806001600160401b031660001480620005315750602082015161018001516001600160401b0316155b156200055c5760405163c35aa79d60e01b81526001600160401b038216600482015260240162000084565b6000600a6000836001600160401b03166001600160401b0316815260200190815260200160002090506000836040015190506000604051806080016040528086602001518152602001836001600160a01b031681526020018460020160149054906101000a90046001600160401b03166001600160401b031681526020018460030154815250905080600001518360000160008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a8154816001600160401b0302191690836001600160401b031602179055506101a082015181600101600c6101000a8154816001600160401b0302191690836001600160401b031602179055506101c08201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555090505082600301546000801b036200092b5760c051604080517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b360208201526001600160401b0392831691810191909152908516606082015230608082015260a00160408051601f1981840301815291905280516020909101206060820181905260038401556001600160a01b03821615620008e2576002830180546001600160a01b0319166001600160a01b0384161790555b836001600160401b03167f7a70081ee29c1fc27898089ba2a5fc35ac0106b043c82ccecd24c6fd48f6ca86846040516200091d9190620021d9565b60405180910390a2620009fd565b60028301546001600160a01b03838116911614620009685760405163c35aa79d60e01b81526001600160401b038516600482015260240162000084565b60208560200151610160015163ffffffff161015620009b557602085015161016001516040516312766e0160e11b81526000600482015263ffffffff909116602482015260440162000084565b836001600160401b03167f944eb884a589931130671ee4a7379fbe5fe65ed605a048ba99c454582f2460b08660200151604051620009f491906200236d565b60405180910390a25b5050505050806001019050620004b8565b5050565b60005b815181101562000a0e57600082828151811062000a365762000a36620021c3565b6020026020010151600001519050600083838151811062000a5b5762000a5b620021c3565b6020908102919091018101518101516001600160a01b0384166000818152600b845260409081902080546001600160401b0319166001600160401b0385169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a2505060010162000a15565b60005b825181101562000d7057600083828151811062000b025762000b02620021c3565b6020026020010151905060008160000151905060005b82602001515181101562000d615760008360200151828151811062000b415762000b41620021c3565b602002602001015160200151905060008460200151838151811062000b6a5762000b6a620021c3565b60200260200101516000015190506020826080015163ffffffff16101562000bc35760808201516040516312766e0160e11b81526001600160a01b038316600482015263ffffffff909116602482015260440162000084565b6001600160401b0384166000818152600c602090815260408083206001600160a01b0386168085529083529281902086518154938801518389015160608a015160808b015160a08c015160c08d01511515600160981b0260ff60981b19911515600160901b0260ff60901b1963ffffffff948516600160701b021664ffffffffff60701b199585166a01000000000000000000000263ffffffff60501b1961ffff90981668010000000000000000029790971665ffffffffffff60401b19988616640100000000026001600160401b0319909d1695909916949094179a909a179590951695909517929092171617949094171692909217909155519091907f16a6faa936552870f38ad6586ca4ae10b5d085667b357895aebb320becccf8d49062000d4e908690600060e08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015260c0830151151560c083015292915050565b60405180910390a3505060010162000b18565b50505080600101905062000ae1565b5060005b815181101562000e3657600082828151811062000d955762000d95620021c3565b6020026020010151600001519050600083838151811062000dba5762000dba620021c3565b6020908102919091018101518101516001600160401b0384166000818152600c845260408082206001600160a01b038516808452955280822080546001600160a01b03191690555192945090917ffa22e84f9c809b5b7e94f084eb45cf17a5e4703cecef8f27ed35e54b719bffcd9190a3505060010162000d74565b505050565b8051604081111562000e6057604051635ad0867d60e11b815260040160405180910390fd5b600d546c01000000000000000000000000900463ffffffff161580159062000eaa5750600d5463ffffffff6c010000000000000000000000008204166001600160601b0390911610155b1562000eba5762000eba6200105d565b600062000ec8600762001255565b90505b801562000f1457600062000eee62000ee5600184620024b6565b60079062001268565b50905062000efe60078262001286565b50508062000f0c90620024cc565b905062000ecb565b506000805b8281101562000ff457600084828151811062000f395762000f39620021c3565b6020026020010151600001519050600085838151811062000f5e5762000f5e620021c3565b602002602001015160200151905060a0516001600160a01b0316826001600160a01b0316148062000f9657506001600160a01b038216155b1562000fc157604051634de938d160e01b81526001600160a01b038316600482015260240162000084565b62000fd360078361ffff8416620012a4565b5062000fe461ffff821685620024e6565b9350505080600101905062000f19565b50600d805463ffffffff60601b19166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd249062001050908390869062002506565b60405180910390a1505050565b6000546001600160a01b031633148015906200108457506002546001600160a01b03163314155b80156200109b575062001099600733620012c4565b155b15620010ba5760405163032bb72b60e31b815260040160405180910390fd5b600d546c01000000000000000000000000900463ffffffff166000819003620010f65760405163990e30bf60e01b815260040160405180910390fd5b600d546001600160601b03168181101562001124576040516311a1ee3b60e31b815260040160405180910390fd5b600062001130620012db565b12156200115057604051631e9acf1760e31b815260040160405180910390fd5b8060006200115f600762001255565b905060005b818110156200122f576000806200117d60078462001268565b90925090506000876200119a836001600160601b038a1662002576565b620011a6919062002590565b9050620011b48187620025b3565b60a051909650620011d9906001600160a01b0316846001600160601b03841662001369565b6040516001600160601b03821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a250505080600101905062001164565b5050600d80546001600160601b0319166001600160601b03929092169190911790555050565b60006200126282620013c1565b92915050565b6000808080620012798686620013ce565b9097909650945050505050565b60006200129d836001600160a01b038416620013fb565b9392505050565b6000620012bc846001600160a01b038516846200141a565b949350505050565b60006200129d836001600160a01b03841662001439565b600d5460a0516040516370a0823160e01b81523060048201526000926001600160601b0316916001600160a01b0316906370a0823190602401602060405180830381865afa15801562001332573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620013589190620025d6565b620013649190620025f0565b905090565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b0390811663a9059cbb60e01b1790915262000e369185916200144716565b6000620012628262001518565b60008080620013de858562001523565b600081815260029690960160205260409095205494959350505050565b600081815260028301602052604081208190556200129d838362001531565b60008281526002840160205260408120829055620012bc84846200153f565b60006200129d83836200154d565b6040805180820190915260208082527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65649082015260009062001496906001600160a01b03851690849062001566565b80519091501562000e365780806020019051810190620014b7919062002613565b62000e365760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840162000084565b600062001262825490565b60006200129d838362001577565b60006200129d8383620015a4565b60006200129d8383620016af565b600081815260018301602052604081205415156200129d565b6060620012bc848460008562001701565b6000826000018281548110620015915762001591620021c3565b9060005260206000200154905092915050565b600081815260018301602052604081205480156200169d576000620015cb600183620024b6565b8554909150600090620015e190600190620024b6565b90508181146200164d576000866000018281548110620016055762001605620021c3565b90600052602060002001549050808760000184815481106200162b576200162b620021c3565b6000918252602080832090910192909255918252600188019052604090208390555b855486908062001661576200166162002631565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062001262565b600091505062001262565b5092915050565b6000818152600183016020526040812054620016f85750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562001262565b50600062001262565b606082471015620017645760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840162000084565b600080866001600160a01b031685876040516200178291906200266d565b60006040518083038185875af1925050503d8060008114620017c1576040519150601f19603f3d011682016040523d82523d6000602084013e620017c6565b606091505b509092509050620017da87838387620017e5565b979650505050505050565b606083156200185957825160000362001851576001600160a01b0385163b620018515760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640162000084565b5081620012bc565b620012bc8383815115620018705781518083602001fd5b8060405162461bcd60e51b81526004016200008491906200268b565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715620018c757620018c76200188c565b60405290565b604051606081016001600160401b0381118282101715620018c757620018c76200188c565b6040516101e081016001600160401b0381118282101715620018c757620018c76200188c565b60405160e081016001600160401b0381118282101715620018c757620018c76200188c565b604051601f8201601f191681016001600160401b03811182821017156200196857620019686200188c565b604052919050565b80516001600160a01b03811681146200198857600080fd5b919050565b80516001600160401b03811681146200198857600080fd5b600060c08284031215620019b857600080fd5b60405160c081016001600160401b0381118282101715620019dd57620019dd6200188c565b604052905080620019ee8362001970565b8152620019fe602084016200198d565b602082015260408301516001600160601b038116811462001a1e57600080fd5b604082015262001a316060840162001970565b606082015262001a446080840162001970565b608082015262001a5760a0840162001970565b60a08201525092915050565b60006040828403121562001a7657600080fd5b62001a80620018a2565b905062001a8d8262001970565b815262001a9d6020830162001970565b602082015292915050565b60006001600160401b0382111562001ac45762001ac46200188c565b5060051b60200190565b805180151581146200198857600080fd5b805161ffff811681146200198857600080fd5b805163ffffffff811681146200198857600080fd5b600082601f83011262001b1957600080fd5b8151602062001b3262001b2c8362001aa8565b6200193d565b828152610220928302850182019282820191908785111562001b5357600080fd5b8387015b8581101562001cfd578089038281121562001b725760008081fd5b62001b7c620018cd565b62001b87836200198d565b81526101e080601f198401121562001b9f5760008081fd5b62001ba9620018f2565b925062001bb888850162001ace565b8352604062001bc981860162001adf565b89850152606062001bdc81870162001af2565b82860152608062001bef81880162001af2565b8287015260a0915062001c0482880162001af2565b9086015260c062001c1787820162001adf565b8287015260e0915062001c2c82880162001af2565b9086015261010062001c4087820162001adf565b82870152610120915062001c5682880162001adf565b9086015261014062001c6a87820162001adf565b82870152610160915062001c8082880162001af2565b9086015261018062001c9487820162001af2565b828701526101a0915062001caa8288016200198d565b908601526101c062001cbe8782016200198d565b8287015262001ccf84880162001af2565b818701525050838984015262001ce9610200860162001970565b908301525085525092840192810162001b57565b5090979650505050505050565b80516001600160801b03811681146200198857600080fd5b60006060828403121562001d3557600080fd5b62001d3f620018cd565b905062001d4c8262001ace565b815262001d5c6020830162001d0a565b602082015262001d6f6040830162001d0a565b604082015292915050565b600082601f83011262001d8c57600080fd5b8151602062001d9f62001b2c8362001aa8565b82815260069290921b8401810191818101908684111562001dbf57600080fd5b8286015b8481101562001e15576040818903121562001dde5760008081fd5b62001de8620018a2565b62001df38262001970565b815262001e028583016200198d565b8186015283529183019160400162001dc3565b509695505050505050565b600082601f83011262001e3257600080fd5b8151602062001e4562001b2c8362001aa8565b82815260059290921b8401810191818101908684111562001e6557600080fd5b8286015b8481101562001e155780516001600160401b038082111562001e8b5760008081fd5b908801906040601f19838c03810182131562001ea75760008081fd5b62001eb1620018a2565b62001ebe8986016200198d565b8152828501518481111562001ed35760008081fd5b8086019550508c603f86011262001eec57600093508384fd5b88850151935062001f0162001b2c8562001aa8565b84815260089490941b8501830193898101908e86111562001f225760008081fd5b958401955b858710156200201657868f0361010081121562001f445760008081fd5b62001f4e620018a2565b62001f598962001970565b815260e080878401121562001f6e5760008081fd5b62001f7862001918565b925062001f878e8b0162001af2565b835262001f96888b0162001af2565b8e840152606062001fa9818c0162001adf565b89850152608062001fbc818d0162001af2565b8286015260a0915062001fd1828d0162001af2565b9085015260c062001fe48c820162001ace565b8286015262001ff5838d0162001ace565b908501525050808d019190915282526101009690960195908a019062001f27565b828b01525087525050509284019250830162001e69565b600082601f8301126200203f57600080fd5b815160206200205262001b2c8362001aa8565b82815260069290921b840181019181810190868411156200207257600080fd5b8286015b8481101562001e155760408189031215620020915760008081fd5b6200209b620018a2565b620020a68262001970565b8152620020b585830162001adf565b8186015283529183019160400162002076565b60008060008060008060006101e0888a031215620020e557600080fd5b620020f18989620019a5565b9650620021028960c08a0162001a63565b6101008901519096506001600160401b03808211156200212157600080fd5b6200212f8b838c0162001b07565b9650620021418b6101208c0162001d22565b95506101808a01519150808211156200215957600080fd5b620021678b838c0162001d7a565b94506101a08a01519150808211156200217f57600080fd5b6200218d8b838c0162001e20565b93506101c08a0151915080821115620021a557600080fd5b50620021b48a828b016200202d565b91505092959891949750929550565b634e487b7160e01b600052603260045260246000fd5b815460ff81161515825261024082019061ffff600882901c8116602085015263ffffffff601883901c811660408601526200222160608601828560381c1663ffffffff169052565b6200223960808601828560581c1663ffffffff169052565b6200224f60a08601838560781c1661ffff169052565b6200226760c08601828560881c1663ffffffff169052565b6200227d60e08601838560a81c1661ffff169052565b620022946101008601838560b81c1661ffff169052565b620022ab6101208601838560c81c1661ffff169052565b620022c46101408601828560d81c1663ffffffff169052565b600186015463ffffffff8282161661016087015292506001600160401b03602084901c811661018087015291506200230d6101a08601838560601c166001600160401b03169052565b620023266101c08601828560a01c1663ffffffff169052565b5060028501546001600160a01b0381166101e086015291506200235a6102008501828460a01c166001600160401b03169052565b5050600383015461022083015292915050565b8151151581526101e0810160208301516200238e602084018261ffff169052565b506040830151620023a7604084018263ffffffff169052565b506060830151620023c0606084018263ffffffff169052565b506080830151620023d9608084018263ffffffff169052565b5060a0830151620023f060a084018261ffff169052565b5060c08301516200240960c084018263ffffffff169052565b5060e08301516200242060e084018261ffff169052565b506101008381015161ffff9081169184019190915261012080850151909116908301526101408084015163ffffffff9081169184019190915261016080850151821690840152610180808501516001600160401b03908116918501919091526101a080860151909116908401526101c09384015116929091019190915290565b634e487b7160e01b600052601160045260246000fd5b81810381811115620012625762001262620024a0565b600081620024de57620024de620024a0565b506000190190565b63ffffffff818116838216019080821115620016a857620016a8620024a0565b6000604080830163ffffffff8616845260206040602086015281865180845260608701915060208801935060005b818110156200256857845180516001600160a01b0316845284015161ffff1684840152938301939185019160010162002534565b509098975050505050505050565b8082028115828204841417620012625762001262620024a0565b600082620025ae57634e487b7160e01b600052601260045260246000fd5b500490565b6001600160601b03828116828216039080821115620016a857620016a8620024a0565b600060208284031215620025e957600080fd5b5051919050565b8181036000831280158383131683831282161715620016a857620016a8620024a0565b6000602082840312156200262657600080fd5b6200129d8262001ace565b634e487b7160e01b600052603160045260246000fd5b60005b83811015620026645781810151838201526020016200264a565b50506000910152565b600082516200268181846020870162002647565b9190910192915050565b6020815260008251806020840152620026ac81604085016020870162002647565b601f01601f19169190910160400192915050565b60805160a05160c05160e05161010051610120516164336200278e6000396000818161034d01528181610ecd0152612db901526000818161031e01528181612d9101526137ed0152600081816102ef01528181612d69015261316201526000818161028b01528181612d04015281816136fb0152613f7c01526000818161025c0152818161107a0152818161155101528181611be801528181612abf01528181612cdf015281816134c501526135be0152600081816102bb01528181612d36015261368a01526164336000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c806379ba5097116100f9578063c92b283211610097578063e080bcba11610071578063e080bcba146109fc578063eff7cc4814610a0f578063f2fde38b14610a17578063fbca3b7414610a2a57600080fd5b8063c92b2832146109ce578063d09dc339146109e1578063df0aa9e9146109e957600080fd5b80638da5cb5b116100d35780638da5cb5b146109815780639041be3d14610992578063a69c64c0146109a5578063b06d41bc146109b857600080fd5b806379ba5097146107d057806382b49eb0146107d8578063869b7f621461096e57600080fd5b8063549e946f116101665780636def4ce7116101405780636def4ce7146104dc578063704b6c02146107695780637437ff9f1461077c57806376f6ae76146107bd57600080fd5b8063549e946f1461049857806354b71468146104ab578063599f6431146104cb57600080fd5b806320487ded116101a257806320487ded146103d35780634510d293146103f457806348a98aa414610409578063546719cd1461043457600080fd5b8063061877e3146101c957806306285c691461021a578063181f5a771461038a575b600080fd5b6101fc6101d7366004614c74565b6001600160a01b03166000908152600b602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff90911681526020015b60405180910390f35b61037d6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a08101919091526040518060c001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b6040516102119190614c91565b6103c66040518060400160405280601c81526020017f45564d3245564d4d756c74694f6e52616d7020312e362e302d6465760000000081525081565b6040516102119190614d4d565b6103e66103e1366004614d99565b610a4a565b604051908152602001610211565b610407610402366004614fd0565b610e7c565b005b61041c61041736600461526f565b610e92565b6040516001600160a01b039091168152602001610211565b61043c610f41565b604051610211919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b6104076104a63660046152a8565b610fe9565b600d546040516bffffffffffffffffffffffff9091168152602001610211565b6002546001600160a01b031661041c565b61075c6104ea3660046152c6565b604080516102608101825260006080820181815260a0830182905260c0830182905260e08301829052610100830182905261012083018290526101408301829052610160830182905261018083018290526101a083018290526101c083018290526101e0830182905261020083018290526102208301829052610240830182905282526020820181905291810182905260608101919091525067ffffffffffffffff9081166000908152600a6020908152604091829020825161026081018452815460ff811615156080830190815261ffff610100808404821660a086015263ffffffff63010000008504811660c08701526701000000000000008504811660e08701526b01000000000000000000000085048116918601919091526f01000000000000000000000000000000840482166101208601527101000000000000000000000000000000000084048116610140860152750100000000000000000000000000000000000000000084048216610160860152770100000000000000000000000000000000000000000000008404821661018086015279010000000000000000000000000000000000000000000000000084049091166101a08501527b0100000000000000000000000000000000000000000000000000000090920482166101c084015260018401548083166101e0850152640100000000810488166102008501526c01000000000000000000000000810488166102208501527401000000000000000000000000000000000000000090819004909216610240840152825260028301546001600160a01b0381169483019490945290920490931691810191909152600390910154606082015290565b6040516102119190615405565b610407610777366004614c74565b611162565b604080518082018252600080825260209182015281518083019092526005546001600160a01b03908116835260065416908201526040516102119190615452565b6104076107cb366004615476565b611221565b610407611284565b6109026107e636600461526f565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091525067ffffffffffffffff82166000908152600c602090815260408083206001600160a01b0385168452825291829020825160e081018452905463ffffffff8082168352640100000000820481169383019390935261ffff68010000000000000000820416938201939093526a01000000000000000000008304821660608201526e0100000000000000000000000000008304909116608082015260ff720100000000000000000000000000000000000083048116151560a0830152730100000000000000000000000000000000000000909204909116151560c082015292915050565b6040516102119190600060e08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015260c0830151151560c083015292915050565b61040761097c3660046154eb565b611342565b6000546001600160a01b031661041c565b6101fc6109a03660046152c6565b611356565b6104076109b336600461552c565b61139a565b6109c06113ab565b60405161021192919061563e565b6104076109dc366004615680565b6114a6565b6103e661150e565b6103e66109f73660046156d0565b6115ce565b610407610a0a36600461573c565b611a18565b610407611a29565b610407610a25366004614c74565b611cba565b610a3d610a383660046152c6565b611ccb565b6040516102119190615936565b67ffffffffffffffff82166000908152600a60205260408120805460ff16610aaf576040517f99ac52f200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526024015b60405180910390fd5b6000610abe6080850185615983565b159050610adf57610ada610ad56080860186615983565b611cff565b610af7565b6001820154640100000000900467ffffffffffffffff165b9050610b2185610b0a6020870187615983565b905083610b1a60408901896159ca565b9050611da7565b6000600b81610b366080880160608901614c74565b6001600160a01b03168152602081019190915260400160009081205467ffffffffffffffff169150819003610bb357610b756080860160608701614c74565b6040517fa7499d200000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610aa6565b60065460009081906001600160a01b031663ffdb4b37610bd960808a0160608b01614c74565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015267ffffffffffffffff8b1660248201526044016040805180830381865afa158015610c44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c689190615a40565b90925090506000808080610c7f60408c018c6159ca565b90501115610cba57610cae8b610c9b60808d0160608e01614c74565b87610ca960408f018f6159ca565b611eb7565b91945092509050610cf1565b6001880154610cee9074010000000000000000000000000000000000000000900463ffffffff16662386f26fc10000615a80565b92505b875460009077010000000000000000000000000000000000000000000000900461ffff1615610d5d57610d5a8c6dffffffffffffffffffffffffffff607088901c16610d4060208f018f615983565b90508e8060400190610d5291906159ca565b9050866122ca565b90505b600089600101600c9054906101000a900467ffffffffffffffff1667ffffffffffffffff168463ffffffff168b600001600f9054906101000a900461ffff1661ffff168e8060200190610db09190615983565b610dbb929150615a80565b8c54610ddc906b010000000000000000000000900463ffffffff168d615a97565b610de69190615a97565b610df09190615a97565b610e0a906dffffffffffffffffffffffffffff8916615a80565b610e149190615a80565b90507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff87168282610e4b67ffffffffffffffff8c1689615a80565b610e559190615a97565b610e5f9190615a97565b610e699190615aaa565b9a50505050505050505050505b92915050565b610e846123d0565b610e8e828261242f565b5050565b6040517fbbe4f6db0000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063bbe4f6db90602401602060405180830381865afa158015610f16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f3a9190615acc565b9392505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526003546fffffffffffffffffffffffffffffffff8082168352600160801b80830463ffffffff1660208501527401000000000000000000000000000000000000000090920460ff161515938301939093526004548084166060840152049091166080820152610fe490612854565b905090565b610ff16123d0565b6001600160a01b038116611031576040517f232cb97f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061103b61150e565b90506000811215611078576040517f02075e0000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316036110ca576110c56001600160a01b0384168383612906565b505050565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526110c59083906001600160a01b038616906370a0823190602401602060405180830381865afa15801561112d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111519190615ae9565b6001600160a01b0386169190612906565b6000546001600160a01b0316331480159061118857506002546001600160a01b03163314155b156111bf576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c906020015b60405180910390a150565b6112296123d0565b610e8e8282808060200260200160405190810160405280939291908181526020016000905b8282101561127a5761126b60408302860136819003810190615b02565b8152602001906001019061124e565b5050505050612986565b6001546001600160a01b031633146112de5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610aa6565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61134a612bf3565b61135381612c4d565b50565b67ffffffffffffffff8082166000908152600a60205260408120600201549091610e7691740100000000000000000000000000000000000000009004166001615b35565b6113a26123d0565b61135381612e0f565b60606000806113ba6007612eec565b90508067ffffffffffffffff8111156113d5576113d5614de9565b60405190808252806020026020018201604052801561141a57816020015b60408051808201909152600080825260208201528152602001906001900390816113f35790505b50925060005b8181101561148357600080611436600784612ef7565b915091506040518060400160405280836001600160a01b031681526020018261ffff1681525086848151811061146e5761146e615b56565b60209081029190910101525050600101611420565b5050600d5491926c0100000000000000000000000090920463ffffffff16919050565b6000546001600160a01b031633148015906114cc57506002546001600160a01b03163314155b15611503576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611353600382612f15565b600d546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000916bffffffffffffffffffffffff16907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa1580156115a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115c49190615ae9565b610fe49190615b6c565b67ffffffffffffffff84166000908152600a60205260408120816115f582888888886130ae565b905060005b816101400151518110156119ae57600061161760408901896159ca565b8381811061162757611627615b56565b90506040020180360381019061163d9190615b8c565b9050600061164f8a8360000151610e92565b90506001600160a01b038116158061170557506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa1580156116df573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117039190615bc6565b155b1561174a5781516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610aa6565b6000816001600160a01b0316639a4575b96040518060a001604052808d80600001906117769190615983565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525067ffffffffffffffff8f166020808301919091526001600160a01b03808e16604080850191909152918901516060840152885116608090920191909152517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526118219190600401615be3565b6000604051808303816000875af1158015611840573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526118689190810190615cb0565b905060208160200151511180156118c9575067ffffffffffffffff8b166000908152600c6020908152604080832086516001600160a01b0316845282529091205490820151516e01000000000000000000000000000090910463ffffffff16105b1561190e5782516040517f36f536ca0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610aa6565b8051611919906139ed565b5060408051606081019091526001600160a01b03831660808201528060a08101604051602081830303815290604052815260200182600001518152602001826020015181525060405160200161196f9190615d41565b604051602081830303815290604052856101600151858151811061199557611995615b56565b60200260200101819052505050508060010190506115fa565b506119bd818360030154613a48565b61018082015260405167ffffffffffffffff8816907fc79f9c3e610deac14de4e704195fe17eab0983ee9916866bc04d16a00f54daa6906119ff908490615e38565b60405180910390a261018001519150505b949350505050565b611a206123d0565b61135381613ba3565b6000546001600160a01b03163314801590611a4f57506002546001600160a01b03163314155b8015611a635750611a6160073361416a565b155b15611a9a576040517f195db95800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d546c01000000000000000000000000900463ffffffff166000819003611aee576040517f990e30bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d546bffffffffffffffffffffffff1681811015611b39576040517f8d0f71d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611b4361150e565b1215611b7b576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000611b886007612eec565b905060005b81811015611c7757600080611ba3600784612ef7565b9092509050600087611bc3836bffffffffffffffffffffffff8a16615a80565b611bcd9190615aaa565b9050611bd98187615f6d565b9550611c1d6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016846bffffffffffffffffffffffff8416612906565b6040516bffffffffffffffffffffffff821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a2505050806001019050611b8d565b5050600d80547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff929092169190911790555050565b611cc2612bf3565b6113538161417f565b60606040517f9e7177c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f97a657c900000000000000000000000000000000000000000000000000000000611d2c8385615f92565b7fffffffff000000000000000000000000000000000000000000000000000000001614611d85576040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611d928260048186615fda565b810190611d9f9190616004565b519392505050565b67ffffffffffffffff84166000908152600a6020526040902080546301000000900463ffffffff16841115611e205780546040517f86933789000000000000000000000000000000000000000000000000000000008152630100000090910463ffffffff16600482015260248101859052604401610aa6565b8054670100000000000000900463ffffffff16831115611e6c576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8054610100900461ffff16821115611eb0576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b6000808083815b818110156122bd576000878783818110611eda57611eda615b56565b905060400201803603810190611ef09190615b8c565b905060006001600160a01b0316611f0b8c8360000151610e92565b6001600160a01b031603611f595780516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610aa6565b67ffffffffffffffff8b166000908152600c6020908152604080832084516001600160a01b03168452825291829020825160e081018452905463ffffffff8082168352640100000000820481169383019390935261ffff68010000000000000000820416938201939093526a01000000000000000000008304821660608201526e0100000000000000000000000000008304909116608082015260ff720100000000000000000000000000000000000083048116151560a0830152730100000000000000000000000000000000000000909204909116151560c082018190526120e95767ffffffffffffffff8c166000908152600a60205260409020805461208990790100000000000000000000000000000000000000000000000000900461ffff16662386f26fc10000615a80565b6120939089615a97565b81549098506120c7907b01000000000000000000000000000000000000000000000000000000900463ffffffff1688616046565b60018201549097506120df9063ffffffff1687616046565b95505050506122b5565b604081015160009061ffff16156122055760008c6001600160a01b031684600001516001600160a01b0316146121a85760065484516040517f4ab35b0b0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911690634ab35b0b90602401602060405180830381865afa15801561217d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121a19190616063565b90506121ab565b508a5b620186a0836040015161ffff166121ed8660200151847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1661423590919063ffffffff16565b6121f79190615a80565b6122019190615aaa565b9150505b60608201516122149088616046565b96508160800151866122269190616046565b82519096506000906122459063ffffffff16662386f26fc10000615a80565b90508082101561226457612259818a615a97565b9850505050506122b5565b6000836020015163ffffffff16662386f26fc100006122839190615a80565b9050808311156122a357612297818b615a97565b995050505050506122b5565b6122ad838b615a97565b995050505050505b600101611ebe565b5050955095509592505050565b60008063ffffffff83166122df608086615a80565b6122eb87610220615a97565b6122f59190615a97565b6122ff9190615a97565b67ffffffffffffffff88166000908152600a6020526040812080549293509171010000000000000000000000000000000000810463ffffffff1690612361907501000000000000000000000000000000000000000000900461ffff1685615a80565b61236b9190615a97565b825490915077010000000000000000000000000000000000000000000000900461ffff166123a96dffffffffffffffffffffffffffff8a1683615a80565b6123b39190615a80565b6123c390655af3107a4000615a80565b9998505050505050505050565b6000546001600160a01b031633148015906123f657506002546001600160a01b03163314155b1561242d576040517ffbdb8e5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60005b825181101561278857600083828151811061244f5761244f615b56565b6020026020010151905060008160000151905060005b82602001515181101561277a5760008360200151828151811061248a5761248a615b56565b60200260200101516020015190506000846020015183815181106124b0576124b0615b56565b60200260200101516000015190506020826080015163ffffffff1610156125205760808201516040517f24ecdc020000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015263ffffffff9091166024820152604401610aa6565b67ffffffffffffffff84166000818152600c602090815260408083206001600160a01b0386168085529083529281902086518154938801518389015160608a015160808b015160a08c015160c08d01511515730100000000000000000000000000000000000000027fffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffff9115157201000000000000000000000000000000000000027fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff63ffffffff9485166e01000000000000000000000000000002167fffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffffff9585166a0100000000000000000000027fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff61ffff9098166801000000000000000002979097167fffffffffffffffffffffffffffffffffffff000000000000ffffffffffffffff988616640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000909d1695909916949094179a909a179590951695909517929092171617949094171692909217909155519091907f16a6faa936552870f38ad6586ca4ae10b5d085667b357895aebb320becccf8d490612768908690600060e08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015260c0830151151560c083015292915050565b60405180910390a35050600101612465565b505050806001019050612432565b5060005b81518110156110c55760008282815181106127a9576127a9615b56565b602002602001015160000151905060008383815181106127cb576127cb615b56565b60209081029190910181015181015167ffffffffffffffff84166000818152600c845260408082206001600160a01b0385168084529552808220805473ffffffffffffffffffffffffffffffffffffffff191690555192945090917ffa22e84f9c809b5b7e94f084eb45cf17a5e4703cecef8f27ed35e54b719bffcd9190a3505060010161278c565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526128e282606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff16426128c6919061607e565b85608001516fffffffffffffffffffffffffffffffff16614272565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526110c590849061429a565b805160408111156129c3576040517fb5a10cfa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d546c01000000000000000000000000900463ffffffff1615801590612a115750600d5463ffffffff6c010000000000000000000000008204166bffffffffffffffffffffffff90911610155b15612a1e57612a1e611a29565b6000612a2a6007612eec565b90505b8015612a6c576000612a4b612a4360018461607e565b600790612ef7565b509050612a5960078261437f565b505080612a6590616091565b9050612a2d565b506000805b82811015612b74576000848281518110612a8d57612a8d615b56565b60200260200101516000015190506000858381518110612aaf57612aaf615b56565b60200260200101516020015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b03161480612b0457506001600160a01b038216155b15612b46576040517f4de938d10000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610aa6565b612b5660078361ffff8416614394565b50612b6561ffff821685616046565b93505050806001019050612a71565b50600d80547fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd2490612be690839086906160c6565b60405180910390a1505050565b6000546001600160a01b0316331461242d5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610aa6565b60208101516001600160a01b0316612c91576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516005805473ffffffffffffffffffffffffffffffffffffffff199081166001600160a01b039384161790915560208084015160068054909316908416179091556040805160c0810182527f0000000000000000000000000000000000000000000000000000000000000000841681527f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16928101929092527f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff16828201527f0000000000000000000000000000000000000000000000000000000000000000831660608301527f0000000000000000000000000000000000000000000000000000000000000000831660808301527f000000000000000000000000000000000000000000000000000000000000000090921660a082015290517fbbf7a36713bfcf4a7869fbf276c760a1c96d364781b5516aa806b8dc1e7362e4916112169184906160e5565b60005b8151811015610e8e576000828281518110612e2f57612e2f615b56565b60200260200101516000015190506000838381518110612e5157612e51615b56565b6020908102919091018101518101516001600160a01b0384166000818152600b845260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff85169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a25050600101612e12565b6000610e76826143aa565b6000808080612f0686866143b5565b909450925050505b9250929050565b8154600090612f3190600160801b900463ffffffff164261607e565b90508015612fae5760018301548354612f6c916fffffffffffffffffffffffffffffffff808216928116918591600160801b90910416614272565b83546fffffffffffffffffffffffffffffffff9190911673ffffffffffffffffffffffffffffffffffffffff1990911617600160801b4263ffffffff16021783555b60208201518354612fd4916fffffffffffffffffffffffffffffffff90811691166143e0565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff9283161717845560208301516040808501518316600160801b0291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990612be69084908151151581526020808301516fffffffffffffffffffffffffffffffff90811691830191909152604092830151169181019190915260600190565b604080516101a08101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820183905260e082018390526101008201839052610120820181905261014082018190526101608201526101808101919091526040517f2cbc26bb000000000000000000000000000000000000000000000000000000008152608086901b77ffffffffffffffff000000000000000000000000000000001660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632cbc26bb90602401602060405180830381865afa1580156131b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131d59190615bc6565b15613218576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610aa6565b6001600160a01b038216613258576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005546001600160a01b0316331461329c576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855460ff166132e3576040517f99ac52f200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610aa6565b60006132f26080860186615983565b15905061330e57613309610ad56080870187615983565b613326565b6001870154640100000000900467ffffffffffffffff165b9050600061333760408701876159ca565b915061335590508761334c6020890189615983565b90508484611da7565b80156134bb576000805b828110156134a95761337460408901896159ca565b8281811061338457613384615b56565b905060400201602001356000036133c7576040517f5cf0444900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff89166000908152600c602052604080822091906133ef908b018b6159ca565b848181106133ff576133ff615b56565b6134159260206040909202019081019150614c74565b6001600160a01b031681526020810191909152604001600020547201000000000000000000000000000000000000900460ff16156134a15761349461345d60408a018a6159ca565b8381811061346d5761346d615b56565b9050604002018036038101906134839190615b8c565b6006546001600160a01b03166143f6565b61349e9083615a97565b91505b60010161335f565b5080156134b9576134b981614517565b505b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166134f56080880160608901614c74565b6001600160a01b03160361355957600d80548691906000906135269084906bffffffffffffffffffffffff16616174565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff160217905550613678565b6006546001600160a01b03166241e5be6135796080890160608a01614c74565b60405160e083901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b039182166004820152602481018990527f00000000000000000000000000000000000000000000000000000000000000009091166044820152606401602060405180830381865afa158015613605573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136299190615ae9565b600d80546000906136499084906bffffffffffffffffffffffff16616174565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055505b600d546bffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116911611156136e5576040517fe5c7a49100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516101a08101825267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b038616602082015290810161377661373c8980615983565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506139ed92505050565b6001600160a01b0316815260200189600201601481819054906101000a900467ffffffffffffffff166137a890616199565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff1681526020018381526020016000151581526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637a2dbcbc8a8860405160200161383b91906001600160a01b0391909116815260200190565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016138679291906161c0565b6020604051808303816000875af1158015613886573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138aa91906161e3565b67ffffffffffffffff1681526020016138c96080890160608a01614c74565b6001600160a01b031681526020018681526020018780602001906138ed9190615983565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200161393460408901896159ca565b808060200260200160405190810160405280939291908181526020016000905b828210156139805761397160408302860136819003810190615b8c565b81526020019060010190613954565b505050505081526020018267ffffffffffffffff8111156139a3576139a3614de9565b6040519080825280602002602001820160405280156139d657816020015b60608152602001906001900390816139c15790505b508152600060209091015298975050505050505050565b60008151602014613a2c57816040517f8d666f60000000000000000000000000000000000000000000000000000000008152600401610aa69190614d4d565b610e7682806020019051810190613a439190615ae9565b614524565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b6101000151604051602001613ade9897969594939291906001600160a01b039889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b6040516020818303038152906040528051906020012085610120015180519060200120866101400151604051602001613b179190616200565b60405160208183030381529060405280519060200120876101600151604051602001613b439190616213565b60408051601f198184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b60005b8151811015610e8e576000828281518110613bc357613bc3615b56565b602002602001015190506000838381518110613be157613be1615b56565b60200260200101516000015190508067ffffffffffffffff1660001480613c1957506020820151610180015167ffffffffffffffff16155b15613c5c576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610aa6565b6000600a60008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002090506000836040015190506000604051806080016040528086602001518152602001836001600160a01b031681526020018460020160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020018460030154815250905080600001518360000160008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101a082015181600101600c6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101c08201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555090505082600301546000801b0361405a57604080517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3602082015267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811692820192909252908516606082015230608082015260a00160408051601f1981840301815291905280516020909101206060820181905260038401556001600160a01b038216156140135760028301805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0384161790555b8367ffffffffffffffff167f7a70081ee29c1fc27898089ba2a5fc35ac0106b043c82ccecd24c6fd48f6ca868460405161404d9190616226565b60405180910390a261415a565b60028301546001600160a01b038381169116146140af576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610aa6565b60208560200151610160015163ffffffff16101561411357602085015161016001516040517f24ecdc020000000000000000000000000000000000000000000000000000000081526000600482015263ffffffff9091166024820152604401610aa6565b8367ffffffffffffffff167f944eb884a589931130671ee4a7379fbe5fe65ed605a048ba99c454582f2460b0866020015160405161415191906163b2565b60405180910390a25b5050505050806001019050613ba6565b6000610f3a836001600160a01b038416614591565b336001600160a01b038216036141d75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610aa6565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000670de0b6b3a7640000614268837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8616615a80565b610f3a9190615aaa565b6000614291856142828486615a80565b61428c9087615a97565b6143e0565b95945050505050565b60006142ef826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661459d9092919063ffffffff16565b8051909150156110c5578080602001905181019061430d9190615bc6565b6110c55760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610aa6565b6000610f3a836001600160a01b0384166145ac565b6000611a10846001600160a01b038516846145c9565b6000610e76826145e6565b600080806143c385856145f0565b600081815260029690960160205260409095205494959350505050565b60008183106143ef5781610f3a565b5090919050565b81516040517fd02641a00000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009182919084169063d02641a0906024016040805180830381865afa15801561445c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061448091906163c1565b5190507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81166000036144e95783516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610aa6565b6020840151611a10907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690614235565b61135360038260006145fc565b60006001600160a01b0382118061453c575061040082105b1561458d5760408051602081018490520160408051601f19818403018152908290527f8d666f60000000000000000000000000000000000000000000000000000000008252610aa691600401614d4d565b5090565b6000610f3a8383614917565b6060611a10848460008561492f565b60008181526002830160205260408120819055610f3a8383614a21565b60008281526002840160205260408120829055611a108484614a2d565b6000610e76825490565b6000610f3a8383614a39565b825474010000000000000000000000000000000000000000900460ff161580614623575081155b1561462d57505050565b825460018401546fffffffffffffffffffffffffffffffff8083169291169060009061466690600160801b900463ffffffff164261607e565b9050801561470c57818311156146a8576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018601546146d590839085908490600160801b90046fffffffffffffffffffffffffffffffff16614272565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16600160801b4263ffffffff160217875592505b848210156147a9576001600160a01b03841661475e576040517ff94ebcd10000000000000000000000000000000000000000000000000000000081526004810183905260248101869052604401610aa6565b6040517f1a76572a00000000000000000000000000000000000000000000000000000000815260048101839052602481018690526001600160a01b0385166044820152606401610aa6565b8483101561489557600186810154600160801b90046fffffffffffffffffffffffffffffffff169060009082906147e0908261607e565b6147ea878a61607e565b6147f49190615a97565b6147fe9190615aaa565b90506001600160a01b03861661484a576040517f15279c080000000000000000000000000000000000000000000000000000000081526004810182905260248101869052604401610aa6565b6040517fd0c8d23a00000000000000000000000000000000000000000000000000000000815260048101829052602481018690526001600160a01b0387166044820152606401610aa6565b61489f858461607e565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b60008181526001830160205260408120541515610f3a565b6060824710156149a75760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610aa6565b600080866001600160a01b031685876040516149c391906163f4565b60006040518083038185875af1925050503d8060008114614a00576040519150601f19603f3d011682016040523d82523d6000602084013e614a05565b606091505b5091509150614a1687838387614a63565b979650505050505050565b6000610f3a8383614adc565b6000610f3a8383614bd6565b6000826000018281548110614a5057614a50615b56565b9060005260206000200154905092915050565b60608315614ad2578251600003614acb576001600160a01b0385163b614acb5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610aa6565b5081611a10565b611a108383614c25565b60008181526001830160205260408120548015614bc5576000614b0060018361607e565b8554909150600090614b149060019061607e565b9050818114614b79576000866000018281548110614b3457614b34615b56565b9060005260206000200154905080876000018481548110614b5757614b57615b56565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080614b8a57614b8a616410565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610e76565b6000915050610e76565b5092915050565b6000818152600183016020526040812054614c1d57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610e76565b506000610e76565b815115614c355781518083602001fd5b8060405162461bcd60e51b8152600401610aa69190614d4d565b6001600160a01b038116811461135357600080fd5b8035614c6f81614c4f565b919050565b600060208284031215614c8657600080fd5b8135610f3a81614c4f565b60c08101610e7682846001600160a01b0380825116835267ffffffffffffffff60208301511660208401526bffffffffffffffffffffffff60408301511660408401528060608301511660608401528060808301511660808401528060a08301511660a0840152505050565b60005b83811015614d18578181015183820152602001614d00565b50506000910152565b60008151808452614d39816020860160208601614cfd565b601f01601f19169290920160200192915050565b602081526000610f3a6020830184614d21565b67ffffffffffffffff8116811461135357600080fd5b8035614c6f81614d60565b600060a08284031215614d9357600080fd5b50919050565b60008060408385031215614dac57600080fd5b8235614db781614d60565b9150602083013567ffffffffffffffff811115614dd357600080fd5b614ddf85828601614d81565b9150509250929050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715614e2257614e22614de9565b60405290565b60405160e0810167ffffffffffffffff81118282101715614e2257614e22614de9565b6040516060810167ffffffffffffffff81118282101715614e2257614e22614de9565b6040516101e0810167ffffffffffffffff81118282101715614e2257614e22614de9565b604051601f8201601f1916810167ffffffffffffffff81118282101715614ebb57614ebb614de9565b604052919050565b600067ffffffffffffffff821115614edd57614edd614de9565b5060051b60200190565b63ffffffff8116811461135357600080fd5b8035614c6f81614ee7565b803561ffff81168114614c6f57600080fd5b801515811461135357600080fd5b8035614c6f81614f16565b600082601f830112614f4057600080fd5b81356020614f55614f5083614ec3565b614e92565b82815260069290921b84018101918181019086841115614f7457600080fd5b8286015b84811015614fc55760408189031215614f915760008081fd5b614f99614dff565b8135614fa481614d60565b815281850135614fb381614c4f565b81860152835291830191604001614f78565b509695505050505050565b60008060408385031215614fe357600080fd5b67ffffffffffffffff83351115614ff957600080fd5b83601f84358501011261500b57600080fd5b61501b614f508435850135614ec3565b8335840180358083526020808401939260059290921b9091010186101561504157600080fd5b602085358601015b85358601803560051b016020018110156152395767ffffffffffffffff8135111561507357600080fd5b6040601f1982358835890101890301121561508d57600080fd5b615095614dff565b6150a86020833589358a01010135614d60565b863587018235016020810135825267ffffffffffffffff60409091013511156150d057600080fd5b86358701823501604081013501603f810189136150ec57600080fd5b6150fc614f506020830135614ec3565b602082810135808352908201919060081b83016040018b101561511e57600080fd5b604083015b6040602085013560081b85010181101561522057610100818d03121561514857600080fd5b615150614dff565b61515a8235614c4f565b8135815260e0601f19838f0301121561517257600080fd5b61517a614e28565b6151876020840135614ee7565b6020830135815261519b6040840135614ee7565b604083013560208201526151b160608401614f04565b60408201526151c36080840135614ee7565b608083013560608201526151da60a0840135614ee7565b60a083013560808201526151f060c08401614f24565b60a082015261520160e08401614f24565b60c0820152602082810191909152908452929092019161010001615123565b5060208481019190915292865250509283019201615049565b5092505067ffffffffffffffff6020840135111561525657600080fd5b6152668460208501358501614f2f565b90509250929050565b6000806040838503121561528257600080fd5b823561528d81614d60565b9150602083013561529d81614c4f565b809150509250929050565b600080604083850312156152bb57600080fd5b823561528d81614c4f565b6000602082840312156152d857600080fd5b8135610f3a81614d60565b80511515825260208101516152fe602084018261ffff169052565b506040810151615316604084018263ffffffff169052565b50606081015161532e606084018263ffffffff169052565b506080810151615346608084018263ffffffff169052565b5060a081015161535c60a084018261ffff169052565b5060c081015161537460c084018263ffffffff169052565b5060e081015161538a60e084018261ffff169052565b506101008181015161ffff9081169184019190915261012080830151909116908301526101408082015163ffffffff90811691840191909152610160808301518216908401526101808083015167ffffffffffffffff908116918501919091526101a080840151909116908401526101c09182015116910152565b6000610240820190506154198284516152e3565b60208301516001600160a01b03166101e0830152604083015167ffffffffffffffff166102008301526060909201516102209091015290565b60408101610e76828480516001600160a01b03908116835260209182015116910152565b6000806020838503121561548957600080fd5b823567ffffffffffffffff808211156154a157600080fd5b818501915085601f8301126154b557600080fd5b8135818111156154c457600080fd5b8660208260061b85010111156154d957600080fd5b60209290920196919550909350505050565b6000604082840312156154fd57600080fd5b615505614dff565b823561551081614c4f565b8152602083013561552081614c4f565b60208201529392505050565b6000602080838503121561553f57600080fd5b823567ffffffffffffffff81111561555657600080fd5b8301601f8101851361556757600080fd5b8035615575614f5082614ec3565b81815260069190911b8201830190838101908783111561559457600080fd5b928401925b82841015614a1657604084890312156155b25760008081fd5b6155ba614dff565b84356155c581614c4f565b8152848601356155d481614d60565b8187015282526040939093019290840190615599565b60008151808452602080850194506020840160005b8381101561563357815180516001600160a01b0316885283015161ffff1683880152604090960195908201906001016155ff565b509495945050505050565b60408152600061565160408301856155ea565b90508260208301529392505050565b80356fffffffffffffffffffffffffffffffff81168114614c6f57600080fd5b60006060828403121561569257600080fd5b61569a614e4b565b82356156a581614f16565b81526156b360208401615660565b60208201526156c460408401615660565b60408201529392505050565b600080600080608085870312156156e657600080fd5b84356156f181614d60565b9350602085013567ffffffffffffffff81111561570d57600080fd5b61571987828801614d81565b93505060408501359150606085013561573181614c4f565b939692955090935050565b6000602080838503121561574f57600080fd5b823567ffffffffffffffff81111561576657600080fd5b8301601f8101851361577757600080fd5b8035615785614f5082614ec3565b81815261022091820283018401918482019190888411156157a557600080fd5b938501935b8385101561592a57848903818112156157c35760008081fd5b6157cb614e4b565b86356157d681614d60565b81526101e0601f1983018113156157ed5760008081fd5b6157f5614e6e565b9250615802898901614f24565b83526040615811818a01614f04565b8a8501526060615822818b01614ef9565b828601526080615833818c01614ef9565b8287015260a09150615846828c01614ef9565b9086015260c06158578b8201614f04565b8287015260e0915061586a828c01614ef9565b9086015261010061587c8b8201614f04565b828701526101209150615890828c01614f04565b908601526101406158a28b8201614f04565b8287015261016091506158b6828c01614ef9565b908601526101806158c88b8201614ef9565b828701526101a091506158dc828c01614d76565b908601526101c06158ee8b8201614d76565b828701526158fd848c01614ef9565b818701525050838a8401526159156102008a01614c64565b908301525084525093840193918501916157aa565b50979650505050505050565b6020808252825182820181905260009190848201906040850190845b818110156159775783516001600160a01b031683529284019291840191600101615952565b50909695505050505050565b6000808335601e1984360301811261599a57600080fd5b83018035915067ffffffffffffffff8211156159b557600080fd5b602001915036819003821315612f0e57600080fd5b6000808335601e198436030181126159e157600080fd5b83018035915067ffffffffffffffff8211156159fc57600080fd5b6020019150600681901b3603821315612f0e57600080fd5b80517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81168114614c6f57600080fd5b60008060408385031215615a5357600080fd5b615a5c83615a14565b915061526660208401615a14565b634e487b7160e01b600052601160045260246000fd5b8082028115828204841417610e7657610e76615a6a565b80820180821115610e7657610e76615a6a565b600082615ac757634e487b7160e01b600052601260045260246000fd5b500490565b600060208284031215615ade57600080fd5b8151610f3a81614c4f565b600060208284031215615afb57600080fd5b5051919050565b600060408284031215615b1457600080fd5b615b1c614dff565b8235615b2781614c4f565b815261552060208401614f04565b67ffffffffffffffff818116838216019080821115614bcf57614bcf615a6a565b634e487b7160e01b600052603260045260246000fd5b8181036000831280158383131683831282161715614bcf57614bcf615a6a565b600060408284031215615b9e57600080fd5b615ba6614dff565b8235615bb181614c4f565b81526020928301359281019290925250919050565b600060208284031215615bd857600080fd5b8151610f3a81614f16565b602081526000825160a06020840152615bff60c0840182614d21565b905067ffffffffffffffff602085015116604084015260408401516001600160a01b038082166060860152606086015160808601528060808701511660a086015250508091505092915050565b600082601f830112615c5d57600080fd5b815167ffffffffffffffff811115615c7757615c77614de9565b615c8a6020601f19601f84011601614e92565b818152846020838601011115615c9f57600080fd5b611a10826020830160208701614cfd565b600060208284031215615cc257600080fd5b815167ffffffffffffffff80821115615cda57600080fd5b9083019060408286031215615cee57600080fd5b615cf6614dff565b825182811115615d0557600080fd5b615d1187828601615c4c565b825250602083015182811115615d2657600080fd5b615d3287828601615c4c565b60208301525095945050505050565b602081526000825160606020840152615d5d6080840182614d21565b90506020840151601f1980858403016040860152615d7b8383614d21565b92506040860151915080858403016060860152506142918282614d21565b60008151808452602080850194506020840160005b8381101561563357815180516001600160a01b031688528301518388015260409096019590820190600101615dae565b60008282518085526020808601955060208260051b8401016020860160005b84811015615e2b57601f19868403018952615e19838351614d21565b98840198925090830190600101615dfd565b5090979650505050505050565b60208152615e5360208201835167ffffffffffffffff169052565b60006020830151615e6f60408401826001600160a01b03169052565b5060408301516001600160a01b038116606084015250606083015167ffffffffffffffff8116608084015250608083015160a083015260a0830151615eb860c084018215159052565b5060c083015167ffffffffffffffff811660e08401525060e0830151610100615eeb818501836001600160a01b03169052565b840151610120848101919091528401516101a061014080860182905291925090615f196101c0860184614d21565b9250808601519050601f19610160818786030181880152615f3a8584615d99565b945080880151925050610180818786030181880152615f598584615dde565b970151959092019490945250929392505050565b6bffffffffffffffffffffffff828116828216039080821115614bcf57614bcf615a6a565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015615fd25780818660040360031b1b83161692505b505092915050565b60008085851115615fea57600080fd5b83861115615ff757600080fd5b5050820193919092039150565b60006020828403121561601657600080fd5b6040516020810181811067ffffffffffffffff8211171561603957616039614de9565b6040529135825250919050565b63ffffffff818116838216019080821115614bcf57614bcf615a6a565b60006020828403121561607557600080fd5b610f3a82615a14565b81810381811115610e7657610e76615a6a565b6000816160a0576160a0615a6a565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b63ffffffff83168152604060208201526000611a1060408301846155ea565b610100810161615282856001600160a01b0380825116835267ffffffffffffffff60208301511660208401526bffffffffffffffffffffffff60408301511660408401528060608301511660608401528060808301511660808401528060a08301511660a0840152505050565b82516001600160a01b0390811660c084015260208401511660e0830152610f3a565b6bffffffffffffffffffffffff818116838216019080821115614bcf57614bcf615a6a565b600067ffffffffffffffff8083168181036161b6576161b6615a6a565b6001019392505050565b67ffffffffffffffff83168152604060208201526000611a106040830184614d21565b6000602082840312156161f557600080fd5b8151610f3a81614d60565b602081526000610f3a6020830184615d99565b602081526000610f3a6020830184615dde565b815460ff81161515825261024082019061ffff600882901c8116602085015263ffffffff601883901c8116604086015261626d60608601828560381c1663ffffffff169052565b61628460808601828560581c1663ffffffff169052565b61629960a08601838560781c1661ffff169052565b6162b060c08601828560881c1663ffffffff169052565b6162c560e08601838560a81c1661ffff169052565b6162db6101008601838560b81c1661ffff169052565b6162f16101208601838560c81c1661ffff169052565b6163096101408601828560d81c1663ffffffff169052565b600186015463ffffffff82821616610160870152925067ffffffffffffffff602084901c811661018087015291506163536101a08601838560601c1667ffffffffffffffff169052565b61636b6101c08601828560a01c1663ffffffff169052565b5060028501546001600160a01b0381166101e0860152915061639f6102008501828460a01c1667ffffffffffffffff169052565b5050600383015461022083015292915050565b6101e08101610e7682846152e3565b6000604082840312156163d357600080fd5b6163db614dff565b6163e483615a14565b8152602083015161552081614ee7565b60008251616406818460208701614cfd565b9190910192915050565b634e487b7160e01b600052603160045260246000fdfea164736f6c6343000818000a",
}

var EVM2EVMMultiOnRampABI = EVM2EVMMultiOnRampMetaData.ABI

var EVM2EVMMultiOnRampBin = EVM2EVMMultiOnRampMetaData.Bin

func DeployEVM2EVMMultiOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMMultiOnRampStaticConfig, dynamicConfig EVM2EVMMultiOnRampDynamicConfig, destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs, rateLimiterConfig RateLimiterConfig, premiumMultiplierWeiPerEthArgs []EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthArgs, tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, nopsAndWeights []EVM2EVMMultiOnRampNopAndWeight) (common.Address, *types.Transaction, *EVM2EVMMultiOnRamp, error) {
	parsed, err := EVM2EVMMultiOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMMultiOnRampBin), backend, staticConfig, dynamicConfig, destChainConfigArgs, rateLimiterConfig, premiumMultiplierWeiPerEthArgs, tokenTransferFeeConfigArgs, nopsAndWeights)
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetPremiumMultiplierWeiPerEth(opts *bind.CallOpts, token common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getPremiumMultiplierWeiPerEth", token)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetPremiumMultiplierWeiPerEth(token common.Address) (uint64, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetPremiumMultiplierWeiPerEth(&_EVM2EVMMultiOnRamp.CallOpts, token)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetPremiumMultiplierWeiPerEth(token common.Address) (uint64, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetPremiumMultiplierWeiPerEth(&_EVM2EVMMultiOnRamp.CallOpts, token)
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (EVM2EVMMultiOnRampTokenTransferFeeConfig, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getTokenTransferFeeConfig", destChainSelector, token)

	if err != nil {
		return *new(EVM2EVMMultiOnRampTokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMMultiOnRampTokenTransferFeeConfig)).(*EVM2EVMMultiOnRampTokenTransferFeeConfig)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetTokenTransferFeeConfig(destChainSelector uint64, token common.Address) (EVM2EVMMultiOnRampTokenTransferFeeConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetTokenTransferFeeConfig(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector, token)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetTokenTransferFeeConfig(destChainSelector uint64, token common.Address) (EVM2EVMMultiOnRampTokenTransferFeeConfig, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetTokenTransferFeeConfig(&_EVM2EVMMultiOnRamp.CallOpts, destChainSelector, token)
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) ApplyDestChainConfigUpdates(destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyDestChainConfigUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, destChainConfigArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyDestChainConfigUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, destChainConfigArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) ApplyPremiumMultiplierWeiPerEthUpdates(opts *bind.TransactOpts, premiumMultiplierWeiPerEthArgs []EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "applyPremiumMultiplierWeiPerEthUpdates", premiumMultiplierWeiPerEthArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) ApplyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs []EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyPremiumMultiplierWeiPerEthUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, premiumMultiplierWeiPerEthArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) ApplyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs []EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyPremiumMultiplierWeiPerEthUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, premiumMultiplierWeiPerEthArgs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []EVM2EVMMultiOnRampTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []EVM2EVMMultiOnRampTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyTokenTransferFeeConfigUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []EVM2EVMMultiOnRampTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.ApplyTokenTransferFeeConfigUpdates(&_EVM2EVMMultiOnRamp.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
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
	DestChainSelector uint64
	Message           InternalEVM2EVMMessage
	Raw               types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterCCIPSendRequested(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampCCIPSendRequestedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "CCIPSendRequested", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampCCIPSendRequestedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "CCIPSendRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampCCIPSendRequested, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "CCIPSendRequested", destChainSelectorRule)
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

type EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdatedIterator struct {
	Event *EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdated)
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
		it.Event = new(EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdated)
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

func (it *EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdatedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdated struct {
	Token                      common.Address
	PremiumMultiplierWeiPerEth uint64
	Raw                        types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterPremiumMultiplierWeiPerEthUpdated(opts *bind.FilterOpts, token []common.Address) (*EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "PremiumMultiplierWeiPerEthUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdatedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "PremiumMultiplierWeiPerEthUpdated", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchPremiumMultiplierWeiPerEthUpdated(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "PremiumMultiplierWeiPerEthUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdated)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "PremiumMultiplierWeiPerEthUpdated", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParsePremiumMultiplierWeiPerEthUpdated(log types.Log) (*EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdated, error) {
	event := new(EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdated)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "PremiumMultiplierWeiPerEthUpdated", log); err != nil {
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
	DestChainSelector *big.Int
	Token             common.Address
	Raw               types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []*big.Int, token []common.Address) (*EVM2EVMMultiOnRampTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampTokenTransferFeeConfigDeletedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted, destChainSelector []*big.Int, token []common.Address) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule, tokenRule)
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

type EVM2EVMMultiOnRampTokenTransferFeeConfigUpdatedIterator struct {
	Event *EVM2EVMMultiOnRampTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampTokenTransferFeeConfigUpdated)
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
		it.Event = new(EVM2EVMMultiOnRampTokenTransferFeeConfigUpdated)
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

func (it *EVM2EVMMultiOnRampTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	Token                  common.Address
	TokenTransferFeeConfig EVM2EVMMultiOnRampTokenTransferFeeConfig
	Raw                    types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*EVM2EVMMultiOnRampTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampTokenTransferFeeConfigUpdatedIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampTokenTransferFeeConfigUpdated, destChainSelector []uint64, token []common.Address) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampTokenTransferFeeConfigUpdated)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*EVM2EVMMultiOnRampTokenTransferFeeConfigUpdated, error) {
	event := new(EVM2EVMMultiOnRampTokenTransferFeeConfigUpdated)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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
	case _EVM2EVMMultiOnRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMMultiOnRamp.ParseConfigChanged(log)
	case _EVM2EVMMultiOnRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseConfigSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["DestChainAdded"].ID:
		return _EVM2EVMMultiOnRamp.ParseDestChainAdded(log)
	case _EVM2EVMMultiOnRamp.abi.Events["DestChainDynamicConfigUpdated"].ID:
		return _EVM2EVMMultiOnRamp.ParseDestChainDynamicConfigUpdated(log)
	case _EVM2EVMMultiOnRamp.abi.Events["NopPaid"].ID:
		return _EVM2EVMMultiOnRamp.ParseNopPaid(log)
	case _EVM2EVMMultiOnRamp.abi.Events["NopsSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseNopsSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMMultiOnRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMMultiOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMMultiOnRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMMultiOnRamp.abi.Events["PremiumMultiplierWeiPerEthUpdated"].ID:
		return _EVM2EVMMultiOnRamp.ParsePremiumMultiplierWeiPerEthUpdated(log)
	case _EVM2EVMMultiOnRamp.abi.Events["TokenTransferFeeConfigDeleted"].ID:
		return _EVM2EVMMultiOnRamp.ParseTokenTransferFeeConfigDeleted(log)
	case _EVM2EVMMultiOnRamp.abi.Events["TokenTransferFeeConfigUpdated"].ID:
		return _EVM2EVMMultiOnRamp.ParseTokenTransferFeeConfigUpdated(log)
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
	return common.HexToHash("0xc79f9c3e610deac14de4e704195fe17eab0983ee9916866bc04d16a00f54daa6")
}

func (EVM2EVMMultiOnRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (EVM2EVMMultiOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xbbf7a36713bfcf4a7869fbf276c760a1c96d364781b5516aa806b8dc1e7362e4")
}

func (EVM2EVMMultiOnRampDestChainAdded) Topic() common.Hash {
	return common.HexToHash("0x7a70081ee29c1fc27898089ba2a5fc35ac0106b043c82ccecd24c6fd48f6ca86")
}

func (EVM2EVMMultiOnRampDestChainDynamicConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x944eb884a589931130671ee4a7379fbe5fe65ed605a048ba99c454582f2460b0")
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

func (EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdated) Topic() common.Hash {
	return common.HexToHash("0xbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d")
}

func (EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0xfa22e84f9c809b5b7e94f084eb45cf17a5e4703cecef8f27ed35e54b719bffcd")
}

func (EVM2EVMMultiOnRampTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x16a6faa936552870f38ad6586ca4ae10b5d085667b357895aebb320becccf8d4")
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

	GetNopFeesJuels(opts *bind.CallOpts) (*big.Int, error)

	GetNops(opts *bind.CallOpts) (GetNops,

		error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetPremiumMultiplierWeiPerEth(opts *bind.CallOpts, token common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMMultiOnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (EVM2EVMMultiOnRampTokenTransferFeeConfig, error)

	LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error)

	ApplyPremiumMultiplierWeiPerEthUpdates(opts *bind.TransactOpts, premiumMultiplierWeiPerEthArgs []EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []EVM2EVMMultiOnRampTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	PayNops(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMMultiOnRampDynamicConfig) (*types.Transaction, error)

	SetNops(opts *bind.TransactOpts, nopsAndWeights []EVM2EVMMultiOnRampNopAndWeight) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawNonLinkFees(opts *bind.TransactOpts, feeToken common.Address, to common.Address) (*types.Transaction, error)

	FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampAdminSetIterator, error)

	WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAdminSet) (event.Subscription, error)

	ParseAdminSet(log types.Log) (*EVM2EVMMultiOnRampAdminSet, error)

	FilterCCIPSendRequested(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampCCIPSendRequestedIterator, error)

	WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampCCIPSendRequested, destChainSelector []uint64) (event.Subscription, error)

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

	FilterPremiumMultiplierWeiPerEthUpdated(opts *bind.FilterOpts, token []common.Address) (*EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdatedIterator, error)

	WatchPremiumMultiplierWeiPerEthUpdated(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdated, token []common.Address) (event.Subscription, error)

	ParsePremiumMultiplierWeiPerEthUpdated(log types.Log) (*EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthUpdated, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []*big.Int, token []common.Address) (*EVM2EVMMultiOnRampTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted, destChainSelector []*big.Int, token []common.Address) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*EVM2EVMMultiOnRampTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*EVM2EVMMultiOnRampTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampTokenTransferFeeConfigUpdated, destChainSelector []uint64, token []common.Address) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*EVM2EVMMultiOnRampTokenTransferFeeConfigUpdated, error)

	FilterTokensConsumed(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampTokensConsumedIterator, error)

	WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampTokensConsumed) (event.Subscription, error)

	ParseTokensConsumed(log types.Log) (*EVM2EVMMultiOnRampTokensConsumed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
