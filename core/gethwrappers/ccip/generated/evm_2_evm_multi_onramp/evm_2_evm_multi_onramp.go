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
	Router             common.Address
	PriceRegistry      common.Address
	TokenAdminRegistry common.Address
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
	LinkToken       common.Address
	ChainSelector   uint64
	MaxNopFeesJuels *big.Int
	RmnProxy        common.Address
	NonceManager    common.Address
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSendZeroTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"DestinationChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidChainSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"}],\"name\":\"InvalidDestBytesOverhead\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidDestChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"}],\"name\":\"InvalidNopAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkBalanceNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeBalanceReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeesToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNopsToPay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NotAFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdminOrNop\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SourceTokenDataTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyNops\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainDynamicConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NopPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nopWeightsTotal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"NopsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"name\":\"PremiumMultiplierWeiPerEthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destChainSelector\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenTransferFeeConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"name\":\"TokenTransferFeeConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensConsumed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyDestChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyPremiumMultiplierWeiPerEthUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigRemoveArgs[]\",\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\"}],\"name\":\"applyTokenTransferFeeConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestChainConfig\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNopFeesJuels\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNops\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"weightsTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPremiumMultiplierWeiPerEth\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenTransferFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"aggregateRateLimitEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"setNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawNonLinkFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162008ae838038062008ae8833981016040819052620000359162002084565b8333806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c08162000282565b50506040805160a081018252602084810180516001600160801b039081168085524263ffffffff169385018490528751151585870181905292518216606086018190529790950151166080909301839052600380546001600160a01b031916909417600160801b9283021760ff60a01b1916600160a01b90910217909255029091176004555086516001600160a01b0316158062000169575060208701516001600160401b0316155b8062000180575060608701516001600160a01b0316155b8062000197575060808701516001600160a01b0316155b15620001b6576040516306b7c75960e31b815260040160405180910390fd5b86516001600160a01b0390811660a05260208801516001600160401b031660c05260408801516001600160601b031660809081526060890151821660e052880151166101005262000207866200032d565b620002128562000499565b6200021d83620009d9565b604080516000808252602082019092526200026a9184919062000263565b60408051808201909152600080825260208201528152602001906001900390816200023b5790505b5062000aa5565b620002758162000e02565b505050505050506200267c565b336001600160a01b03821603620002dc5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60208101516001600160a01b031662000359576040516306b7c75960e31b815260040160405180910390fd5b8051600580546001600160a01b039283166001600160a01b03199182161790915560208084015160068054918516918416919091179055604080850151600780549186169190941617909255815160a08082018452518416815260c0516001600160401b031691810191909152608080516001600160601b03168284015260e051841660608301526101005190931692810192909252517f7d6ca86e1084c492475f59ce7b74c37ad96a7440eb9c217ff5b27c183957b9b8916200048e91849082516001600160a01b0390811682526020808501516001600160401b0316818401526040808601516001600160601b0316818501526060808701518416908501526080958601518316958401959095528351821660a0840152830151811660c083015291909201511660e08201526101000190565b60405180910390a150565b60005b8151811015620009d5576000828281518110620004bd57620004bd6200217f565b602002602001015190506000838381518110620004de57620004de6200217f565b6020026020010151600001519050806001600160401b0316600003620005235760405163c35aa79d60e01b81526001600160401b038216600482015260240162000084565b6000600b6000836001600160401b03166001600160401b0316815260200190815260200160002090506000836040015190506000604051806080016040528086602001518152602001836001600160a01b031681526020018460020160149054906101000a90046001600160401b03166001600160401b031681526020018460030154815250905080600001518360000160008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a8154816001600160401b0302191690836001600160401b031602179055506101a082015181600101600c6101000a8154816001600160401b0302191690836001600160401b031602179055506101c08201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555090505082600301546000801b03620008f25760c051604080517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b360208201526001600160401b0392831691810191909152908516606082015230608082015260a00160408051601f1981840301815291905280516020909101206060820181905260038401556001600160a01b03821615620008a9576002830180546001600160a01b0319166001600160a01b0384161790555b836001600160401b03167f7a70081ee29c1fc27898089ba2a5fc35ac0106b043c82ccecd24c6fd48f6ca8684604051620008e4919062002195565b60405180910390a2620009c4565b60028301546001600160a01b038381169116146200092f5760405163c35aa79d60e01b81526001600160401b038516600482015260240162000084565b60208560200151610160015163ffffffff1610156200097c57602085015161016001516040516312766e0160e11b81526000600482015263ffffffff909116602482015260440162000084565b836001600160401b03167f944eb884a589931130671ee4a7379fbe5fe65ed605a048ba99c454582f2460b08660200151604051620009bb919062002329565b60405180910390a25b50505050508060010190506200049c565b5050565b60005b8151811015620009d5576000828281518110620009fd57620009fd6200217f565b6020026020010151600001519050600083838151811062000a225762000a226200217f565b6020908102919091018101518101516001600160a01b0384166000818152600c845260409081902080546001600160401b0319166001600160401b0385169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a25050600101620009dc565b60005b825181101562000d3757600083828151811062000ac95762000ac96200217f565b6020026020010151905060008160000151905060005b82602001515181101562000d285760008360200151828151811062000b085762000b086200217f565b602002602001015160200151905060008460200151838151811062000b315762000b316200217f565b60200260200101516000015190506020826080015163ffffffff16101562000b8a5760808201516040516312766e0160e11b81526001600160a01b038316600482015263ffffffff909116602482015260440162000084565b6001600160401b0384166000818152600d602090815260408083206001600160a01b0386168085529083529281902086518154938801518389015160608a015160808b015160a08c015160c08d01511515600160981b0260ff60981b19911515600160901b0260ff60901b1963ffffffff948516600160701b021664ffffffffff60701b199585166a01000000000000000000000263ffffffff60501b1961ffff90981668010000000000000000029790971665ffffffffffff60401b19988616640100000000026001600160401b0319909d1695909916949094179a909a179590951695909517929092171617949094171692909217909155519091907f16a6faa936552870f38ad6586ca4ae10b5d085667b357895aebb320becccf8d49062000d15908690600060e08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015260c0830151151560c083015292915050565b60405180910390a3505060010162000adf565b50505080600101905062000aa8565b5060005b815181101562000dfd57600082828151811062000d5c5762000d5c6200217f565b6020026020010151600001519050600083838151811062000d815762000d816200217f565b6020908102919091018101518101516001600160401b0384166000818152600d845260408082206001600160a01b038516808452955280822080546001600160a01b03191690555192945090917ffa22e84f9c809b5b7e94f084eb45cf17a5e4703cecef8f27ed35e54b719bffcd9190a3505060010162000d3b565b505050565b8051604081111562000e2757604051635ad0867d60e11b815260040160405180910390fd5b600e546c01000000000000000000000000900463ffffffff161580159062000e715750600e5463ffffffff6c010000000000000000000000008204166001600160601b0390911610155b1562000e815762000e8162001024565b600062000e8f60086200121c565b90505b801562000edb57600062000eb562000eac60018462002472565b6008906200122f565b50905062000ec56008826200124d565b50508062000ed39062002488565b905062000e92565b506000805b8281101562000fbb57600084828151811062000f005762000f006200217f565b6020026020010151600001519050600085838151811062000f255762000f256200217f565b602002602001015160200151905060a0516001600160a01b0316826001600160a01b0316148062000f5d57506001600160a01b038216155b1562000f8857604051634de938d160e01b81526001600160a01b038316600482015260240162000084565b62000f9a60088361ffff84166200126b565b5062000fab61ffff821685620024a2565b9350505080600101905062000ee0565b50600e805463ffffffff60601b19166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd2490620010179083908690620024c2565b60405180910390a1505050565b6000546001600160a01b031633148015906200104b57506002546001600160a01b03163314155b8015620010625750620010606008336200128b565b155b15620010815760405163032bb72b60e31b815260040160405180910390fd5b600e546c01000000000000000000000000900463ffffffff166000819003620010bd5760405163990e30bf60e01b815260040160405180910390fd5b600e546001600160601b031681811015620010eb576040516311a1ee3b60e31b815260040160405180910390fd5b6000620010f7620012a2565b12156200111757604051631e9acf1760e31b815260040160405180910390fd5b8060006200112660086200121c565b905060005b81811015620011f657600080620011446008846200122f565b909250905060008762001161836001600160601b038a1662002532565b6200116d91906200254c565b90506200117b81876200256f565b60a051909650620011a0906001600160a01b0316846001600160601b03841662001330565b6040516001600160601b03821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a25050508060010190506200112b565b5050600e80546001600160601b0319166001600160601b03929092169190911790555050565b6000620012298262001388565b92915050565b600080808062001240868662001395565b9097909650945050505050565b600062001264836001600160a01b038416620013c2565b9392505050565b600062001283846001600160a01b03851684620013e1565b949350505050565b600062001264836001600160a01b03841662001400565b600e5460a0516040516370a0823160e01b81523060048201526000926001600160601b0316916001600160a01b0316906370a0823190602401602060405180830381865afa158015620012f9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200131f919062002592565b6200132b9190620025ac565b905090565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b0390811663a9059cbb60e01b1790915262000dfd9185916200140e16565b60006200122982620014df565b60008080620013a58585620014ea565b600081815260029690960160205260409095205494959350505050565b60008181526002830160205260408120819055620012648383620014f8565b6000828152600284016020526040812082905562001283848462001506565b600062001264838362001514565b6040805180820190915260208082527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564908201526000906200145d906001600160a01b0385169084906200152d565b80519091501562000dfd57808060200190518101906200147e9190620025cf565b62000dfd5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840162000084565b600062001229825490565b60006200126483836200153e565b60006200126483836200156b565b600062001264838362001676565b6000818152600183016020526040812054151562001264565b6060620012838484600085620016c8565b60008260000182815481106200155857620015586200217f565b9060005260206000200154905092915050565b60008181526001830160205260408120548015620016645760006200159260018362002472565b8554909150600090620015a89060019062002472565b905081811462001614576000866000018281548110620015cc57620015cc6200217f565b9060005260206000200154905080876000018481548110620015f257620015f26200217f565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080620016285762001628620025ed565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062001229565b600091505062001229565b5092915050565b6000818152600183016020526040812054620016bf5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562001229565b50600062001229565b6060824710156200172b5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840162000084565b600080866001600160a01b0316858760405162001749919062002629565b60006040518083038185875af1925050503d806000811462001788576040519150601f19603f3d011682016040523d82523d6000602084013e6200178d565b606091505b509092509050620017a187838387620017ac565b979650505050505050565b606083156200182057825160000362001818576001600160a01b0385163b620018185760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640162000084565b508162001283565b620012838383815115620018375781518083602001fd5b8060405162461bcd60e51b815260040162000084919062002647565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156200188e576200188e62001853565b60405290565b6040516101e081016001600160401b03811182821017156200188e576200188e62001853565b604080519081016001600160401b03811182821017156200188e576200188e62001853565b60405160e081016001600160401b03811182821017156200188e576200188e62001853565b604051601f8201601f191681016001600160401b03811182821017156200192f576200192f62001853565b604052919050565b80516001600160a01b03811681146200194f57600080fd5b919050565b80516001600160401b03811681146200194f57600080fd5b600060a082840312156200197f57600080fd5b60405160a081016001600160401b0381118282101715620019a457620019a462001853565b604052905080620019b58362001937565b8152620019c56020840162001954565b602082015260408301516001600160601b0381168114620019e557600080fd5b6040820152620019f86060840162001937565b606082015262001a0b6080840162001937565b60808201525092915050565b60006060828403121562001a2a57600080fd5b62001a3462001869565b905062001a418262001937565b815262001a516020830162001937565b602082015262001a646040830162001937565b604082015292915050565b60006001600160401b0382111562001a8b5762001a8b62001853565b5060051b60200190565b805180151581146200194f57600080fd5b805161ffff811681146200194f57600080fd5b805163ffffffff811681146200194f57600080fd5b600082601f83011262001ae057600080fd5b8151602062001af962001af38362001a6f565b62001904565b828152610220928302850182019282820191908785111562001b1a57600080fd5b8387015b8581101562001cc4578089038281121562001b395760008081fd5b62001b4362001869565b62001b4e8362001954565b81526101e080601f198401121562001b665760008081fd5b62001b7062001894565b925062001b7f88850162001a95565b8352604062001b9081860162001aa6565b89850152606062001ba381870162001ab9565b82860152608062001bb681880162001ab9565b8287015260a0915062001bcb82880162001ab9565b9086015260c062001bde87820162001aa6565b8287015260e0915062001bf382880162001ab9565b9086015261010062001c0787820162001aa6565b82870152610120915062001c1d82880162001aa6565b9086015261014062001c3187820162001aa6565b82870152610160915062001c4782880162001ab9565b9086015261018062001c5b87820162001ab9565b828701526101a0915062001c7182880162001954565b908601526101c062001c8587820162001954565b8287015262001c9684880162001ab9565b818701525050838984015262001cb0610200860162001937565b908301525085525092840192810162001b1e565b5090979650505050505050565b80516001600160801b03811681146200194f57600080fd5b60006060828403121562001cfc57600080fd5b62001d0662001869565b905062001d138262001a95565b815262001d236020830162001cd1565b602082015262001a646040830162001cd1565b600082601f83011262001d4857600080fd5b8151602062001d5b62001af38362001a6f565b82815260069290921b8401810191818101908684111562001d7b57600080fd5b8286015b8481101562001dd1576040818903121562001d9a5760008081fd5b62001da4620018ba565b62001daf8262001937565b815262001dbe85830162001954565b8186015283529183019160400162001d7f565b509695505050505050565b600082601f83011262001dee57600080fd5b8151602062001e0162001af38362001a6f565b82815260059290921b8401810191818101908684111562001e2157600080fd5b8286015b8481101562001dd15780516001600160401b038082111562001e475760008081fd5b908801906040601f19838c03810182131562001e635760008081fd5b62001e6d620018ba565b62001e7a89860162001954565b8152828501518481111562001e8f5760008081fd5b8086019550508c603f86011262001ea857600093508384fd5b88850151935062001ebd62001af38562001a6f565b84815260089490941b8501830193898101908e86111562001ede5760008081fd5b958401955b8587101562001fd257868f0361010081121562001f005760008081fd5b62001f0a620018ba565b62001f158962001937565b815260e080878401121562001f2a5760008081fd5b62001f34620018df565b925062001f438e8b0162001ab9565b835262001f52888b0162001ab9565b8e840152606062001f65818c0162001aa6565b89850152608062001f78818d0162001ab9565b8286015260a0915062001f8d828d0162001ab9565b9085015260c062001fa08c820162001a95565b8286015262001fb1838d0162001a95565b908501525050808d019190915282526101009690960195908a019062001ee3565b828b01525087525050509284019250830162001e25565b600082601f83011262001ffb57600080fd5b815160206200200e62001af38362001a6f565b82815260069290921b840181019181810190868411156200202e57600080fd5b8286015b8481101562001dd157604081890312156200204d5760008081fd5b62002057620018ba565b620020628262001937565b81526200207185830162001aa6565b8186015283529183019160400162002032565b60008060008060008060006101e0888a031215620020a157600080fd5b620020ad89896200196c565b9650620020be8960a08a0162001a17565b6101008901519096506001600160401b0380821115620020dd57600080fd5b620020eb8b838c0162001ace565b9650620020fd8b6101208c0162001ce9565b95506101808a01519150808211156200211557600080fd5b620021238b838c0162001d36565b94506101a08a01519150808211156200213b57600080fd5b620021498b838c0162001ddc565b93506101c08a01519150808211156200216157600080fd5b50620021708a828b0162001fe9565b91505092959891949750929550565b634e487b7160e01b600052603260045260246000fd5b815460ff81161515825261024082019061ffff600882901c8116602085015263ffffffff601883901c81166040860152620021dd60608601828560381c1663ffffffff169052565b620021f560808601828560581c1663ffffffff169052565b6200220b60a08601838560781c1661ffff169052565b6200222360c08601828560881c1663ffffffff169052565b6200223960e08601838560a81c1661ffff169052565b620022506101008601838560b81c1661ffff169052565b620022676101208601838560c81c1661ffff169052565b620022806101408601828560d81c1663ffffffff169052565b600186015463ffffffff8282161661016087015292506001600160401b03602084901c81166101808701529150620022c96101a08601838560601c166001600160401b03169052565b620022e26101c08601828560a01c1663ffffffff169052565b5060028501546001600160a01b0381166101e08601529150620023166102008501828460a01c166001600160401b03169052565b5050600383015461022083015292915050565b8151151581526101e0810160208301516200234a602084018261ffff169052565b50604083015162002363604084018263ffffffff169052565b5060608301516200237c606084018263ffffffff169052565b50608083015162002395608084018263ffffffff169052565b5060a0830151620023ac60a084018261ffff169052565b5060c0830151620023c560c084018263ffffffff169052565b5060e0830151620023dc60e084018261ffff169052565b506101008381015161ffff9081169184019190915261012080850151909116908301526101408084015163ffffffff9081169184019190915261016080850151821690840152610180808501516001600160401b03908116918501919091526101a080860151909116908401526101c09384015116929091019190915290565b634e487b7160e01b600052601160045260246000fd5b818103818111156200122957620012296200245c565b6000816200249a576200249a6200245c565b506000190190565b63ffffffff8181168382160190808211156200166f576200166f6200245c565b6000604080830163ffffffff8616845260206040602086015281865180845260608701915060208801935060005b818110156200252457845180516001600160a01b0316845284015161ffff16848401529383019391850191600101620024f0565b509098975050505050505050565b80820281158282048414176200122957620012296200245c565b6000826200256a57634e487b7160e01b600052601260045260246000fd5b500490565b6001600160601b038281168282160390808211156200166f576200166f6200245c565b600060208284031215620025a557600080fd5b5051919050565b81810360008312801583831316838312821617156200166f576200166f6200245c565b600060208284031215620025e257600080fd5b620012648262001a95565b634e487b7160e01b600052603160045260246000fd5b60005b838110156200262057818101518382015260200162002606565b50506000910152565b600082516200263d81846020870162002603565b9190910192915050565b60208152600082518060208401526200266881604085016020870162002603565b601f01601f19169190910160400192915050565b60805160a05160c05160e051610100516163b7620027316000396000818161031701528181612540015261377c0152600081816102e80152818161251801526130f9015260008181610284015281816124b30152818161368a0152613eee015260008181610255015281816110570152818161151a01528181611ba50152818161248e01528181612c8201528181613454015261354d0152600081816102b4015281816124e5015261361901526163b76000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c806376f6ae76116100f9578063c92b283211610097578063e080bcba11610071578063e080bcba146109e5578063eff7cc48146109f8578063f2fde38b14610a00578063fbca3b7414610a1357600080fd5b8063c92b2832146109b7578063d09dc339146109ca578063df0aa9e9146109d257600080fd5b80638da5cb5b116100d35780638da5cb5b1461096a5780639041be3d1461097b578063a69c64c01461098e578063b06d41bc146109a157600080fd5b806376f6ae76146107b957806379ba5097146107cc57806382b49eb0146107d457600080fd5b8063546719cd11610166578063599f643111610140578063599f6431146104a85780636def4ce7146104b9578063704b6c02146107465780637437ff9f1461075957600080fd5b8063546719cd14610411578063549e946f1461047557806354b714681461048857600080fd5b806320487ded116101a257806320487ded1461039d57806334d560e4146103be5780634510d293146103d357806348a98aa4146103e657600080fd5b8063061877e3146101c957806306285c691461021a578063181f5a7714610354575b600080fd5b6101fc6101d7366004614be5565b6001600160a01b03166000908152600c602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff90911681526020015b60405180910390f35b6103476040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040518060a001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b6040516102119190614c02565b6103906040518060400160405280601c81526020017f45564d3245564d4d756c74694f6e52616d7020312e362e302d6465760000000081525081565b6040516102119190614cb2565b6103b06103ab366004614cfe565b610a33565b604051908152602001610211565b6103d16103cc366004614e28565b610e65565b005b6103d16103e1366004614f89565b610e79565b6103f96103f4366004615228565b610e8f565b6040516001600160a01b039091168152602001610211565b610419610f1e565b604051610211919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b6103d1610483366004615261565b610fc6565b600e546040516bffffffffffffffffffffffff9091168152602001610211565b6002546001600160a01b03166103f9565b6107396104c736600461527f565b604080516102608101825260006080820181815260a0830182905260c0830182905260e08301829052610100830182905261012083018290526101408301829052610160830182905261018083018290526101a083018290526101c083018290526101e0830182905261020083018290526102208301829052610240830182905282526020820181905291810182905260608101919091525067ffffffffffffffff9081166000908152600b6020908152604091829020825161026081018452815460ff811615156080830190815261ffff610100808404821660a086015263ffffffff63010000008504811660c08701526701000000000000008504811660e08701526b01000000000000000000000085048116918601919091526f01000000000000000000000000000000840482166101208601527101000000000000000000000000000000000084048116610140860152750100000000000000000000000000000000000000000084048216610160860152770100000000000000000000000000000000000000000000008404821661018086015279010000000000000000000000000000000000000000000000000084049091166101a08501527b0100000000000000000000000000000000000000000000000000000090920482166101c084015260018401548083166101e0850152640100000000810488166102008501526c01000000000000000000000000810488166102208501527401000000000000000000000000000000000000000090819004909216610240840152825260028301546001600160a01b0381169483019490945290920490931691810191909152600390910154606082015290565b60405161021191906153be565b6103d1610754366004614be5565b61113f565b6107ac604080516060810182526000808252602082018190529181019190915250604080516060810182526005546001600160a01b03908116825260065481166020830152600754169181019190915290565b604051610211919061540b565b6103d16107c736600461543b565b6111fe565b6103d1611261565b6108fe6107e2366004615228565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091525067ffffffffffffffff82166000908152600d602090815260408083206001600160a01b0385168452825291829020825160e081018452905463ffffffff8082168352640100000000820481169383019390935261ffff68010000000000000000820416938201939093526a01000000000000000000008304821660608201526e0100000000000000000000000000008304909116608082015260ff720100000000000000000000000000000000000083048116151560a0830152730100000000000000000000000000000000000000909204909116151560c082015292915050565b6040516102119190600060e08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015260c0830151151560c083015292915050565b6000546001600160a01b03166103f9565b6101fc61098936600461527f565b61131f565b6103d161099c3660046154b0565b611363565b6109a9611374565b6040516102119291906155c2565b6103d16109c5366004615604565b61146f565b6103b06114d7565b6103b06109e0366004615648565b611597565b6103d16109f33660046156b4565b6119d5565b6103d16119e6565b6103d1610a0e366004614be5565b611c77565b610a26610a2136600461527f565b611c88565b60405161021191906158ae565b67ffffffffffffffff82166000908152600b60205260408120805460ff16610a98576040517f99ac52f200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526024015b60405180910390fd5b6000610aa760808501856158fb565b159050610ac857610ac3610abe60808601866158fb565b611cbc565b610ae0565b6001820154640100000000900467ffffffffffffffff165b9050610b0a85610af360208701876158fb565b905083610b036040890189615942565b9050611d64565b6000600c81610b1f6080880160608901614be5565b6001600160a01b03168152602081019190915260400160009081205467ffffffffffffffff169150819003610b9c57610b5e6080860160608701614be5565b6040517fa7499d200000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a8f565b60065460009081906001600160a01b031663ffdb4b37610bc260808a0160608b01614be5565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015267ffffffffffffffff8b1660248201526044016040805180830381865afa158015610c2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c5191906159b8565b90925090506000808080610c6860408c018c615942565b90501115610ca357610c978b610c8460808d0160608e01614be5565b87610c9260408f018f615942565b611e74565b91945092509050610cda565b6001880154610cd79074010000000000000000000000000000000000000000900463ffffffff16662386f26fc100006159f8565b92505b875460009077010000000000000000000000000000000000000000000000900461ffff1615610d4657610d438c6dffffffffffffffffffffffffffff607088901c16610d2960208f018f6158fb565b90508e8060400190610d3b9190615942565b905086612287565b90505b600089600101600c9054906101000a900467ffffffffffffffff1667ffffffffffffffff168463ffffffff168b600001600f9054906101000a900461ffff1661ffff168e8060200190610d9991906158fb565b610da49291506159f8565b8c54610dc5906b010000000000000000000000900463ffffffff168d615a0f565b610dcf9190615a0f565b610dd99190615a0f565b610df3906dffffffffffffffffffffffffffff89166159f8565b610dfd91906159f8565b90507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff87168282610e3467ffffffffffffffff8c16896159f8565b610e3e9190615a0f565b610e489190615a0f565b610e529190615a22565b9a50505050505050505050505b92915050565b610e6d61238d565b610e76816123e9565b50565b610e81612595565b610e8b82826125f2565b5050565b6007546040517fbbe4f6db0000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152600092169063bbe4f6db90602401602060405180830381865afa158015610ef3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f179190615a44565b9392505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526003546fffffffffffffffffffffffffffffffff8082168352600160801b80830463ffffffff1660208501527401000000000000000000000000000000000000000090920460ff161515938301939093526004548084166060840152049091166080820152610fc190612a17565b905090565b610fce612595565b6001600160a01b03811661100e576040517f232cb97f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006110186114d7565b90506000811215611055576040517f02075e0000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316036110a7576110a26001600160a01b0384168383612ac9565b505050565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526110a29083906001600160a01b038616906370a0823190602401602060405180830381865afa15801561110a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061112e9190615a61565b6001600160a01b0386169190612ac9565b6000546001600160a01b0316331480159061116557506002546001600160a01b03163314155b1561119c576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c906020015b60405180910390a150565b611206612595565b610e8b8282808060200260200160405190810160405280939291908181526020016000905b828210156112575761124860408302860136819003810190615a7a565b8152602001906001019061122b565b5050505050612b49565b6001546001600160a01b031633146112bb5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610a8f565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b67ffffffffffffffff8082166000908152600b60205260408120600201549091610e5f91740100000000000000000000000000000000000000009004166001615ab9565b61136b612595565b610e7681612db6565b60606000806113836008612e93565b90508067ffffffffffffffff81111561139e5761139e614d4e565b6040519080825280602002602001820160405280156113e357816020015b60408051808201909152600080825260208201528152602001906001900390816113bc5790505b50925060005b8181101561144c576000806113ff600884612e9e565b915091506040518060400160405280836001600160a01b031681526020018261ffff1681525086848151811061143757611437615ada565b602090810291909101015250506001016113e9565b5050600e5491926c0100000000000000000000000090920463ffffffff16919050565b6000546001600160a01b0316331480159061149557506002546001600160a01b03163314155b156114cc576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e76600382612ebc565b600e546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000916bffffffffffffffffffffffff16907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa158015611569573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061158d9190615a61565b610fc19190615af0565b67ffffffffffffffff84166000908152600b60205260408120816115be8288888888613055565b905060005b816101400151518110156119775760006115e06040890189615942565b838181106115f0576115f0615ada565b9050604002018036038101906116069190615b10565b905060006116188a8360000151610e8f565b90506001600160a01b03811615806116ce57506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa1580156116a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116cc9190615b4a565b155b156117135781516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a8f565b6000816001600160a01b0316639a4575b96040518060a001604052808d806000019061173f91906158fb565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525067ffffffffffffffff8f166020808301919091526001600160a01b03808e16604080850191909152918901516060840152885116608090920191909152517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526117ea9190600401615b67565b6000604051808303816000875af1158015611809573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526118319190810190615c34565b90506020816020015151118015611892575067ffffffffffffffff8b166000908152600d6020908152604080832086516001600160a01b0316845282529091205490820151516e01000000000000000000000000000090910463ffffffff16105b156118d75782516040517f36f536ca0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a8f565b80516118e29061397c565b5060408051606081019091526001600160a01b03831660808201528060a0810160405160208183030381529060405281526020018260000151815260200182602001518152506040516020016119389190615cc5565b604051602081830303815290604052856101600151858151811061195e5761195e615ada565b60200260200101819052505050508060010190506115c3565b506119868183600301546139d7565b6101808201526040517fd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd906119bc908390615dbc565b60405180910390a161018001519150505b949350505050565b6119dd612595565b610e7681613b32565b6000546001600160a01b03163314801590611a0c57506002546001600160a01b03163314155b8015611a205750611a1e6008336140dc565b155b15611a57576040517f195db95800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e546c01000000000000000000000000900463ffffffff166000819003611aab576040517f990e30bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e546bffffffffffffffffffffffff1681811015611af6576040517f8d0f71d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611b006114d7565b1215611b38576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000611b456008612e93565b905060005b81811015611c3457600080611b60600884612e9e565b9092509050600087611b80836bffffffffffffffffffffffff8a166159f8565b611b8a9190615a22565b9050611b968187615ef1565b9550611bda6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016846bffffffffffffffffffffffff8416612ac9565b6040516bffffffffffffffffffffffff821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a2505050806001019050611b4a565b5050600e80547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff929092169190911790555050565b611c7f61238d565b610e76816140f1565b60606040517f9e7177c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f97a657c900000000000000000000000000000000000000000000000000000000611ce98385615f16565b7fffffffff000000000000000000000000000000000000000000000000000000001614611d42576040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611d4f8260048186615f5e565b810190611d5c9190615f88565b519392505050565b67ffffffffffffffff84166000908152600b6020526040902080546301000000900463ffffffff16841115611ddd5780546040517f86933789000000000000000000000000000000000000000000000000000000008152630100000090910463ffffffff16600482015260248101859052604401610a8f565b8054670100000000000000900463ffffffff16831115611e29576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8054610100900461ffff16821115611e6d576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b6000808083815b8181101561227a576000878783818110611e9757611e97615ada565b905060400201803603810190611ead9190615b10565b905060006001600160a01b0316611ec88c8360000151610e8f565b6001600160a01b031603611f165780516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a8f565b67ffffffffffffffff8b166000908152600d6020908152604080832084516001600160a01b03168452825291829020825160e081018452905463ffffffff8082168352640100000000820481169383019390935261ffff68010000000000000000820416938201939093526a01000000000000000000008304821660608201526e0100000000000000000000000000008304909116608082015260ff720100000000000000000000000000000000000083048116151560a0830152730100000000000000000000000000000000000000909204909116151560c082018190526120a65767ffffffffffffffff8c166000908152600b60205260409020805461204690790100000000000000000000000000000000000000000000000000900461ffff16662386f26fc100006159f8565b6120509089615a0f565b8154909850612084907b01000000000000000000000000000000000000000000000000000000900463ffffffff1688615fca565b600182015490975061209c9063ffffffff1687615fca565b9550505050612272565b604081015160009061ffff16156121c25760008c6001600160a01b031684600001516001600160a01b0316146121655760065484516040517f4ab35b0b0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911690634ab35b0b90602401602060405180830381865afa15801561213a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061215e9190615fe7565b9050612168565b508a5b620186a0836040015161ffff166121aa8660200151847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166141a790919063ffffffff16565b6121b491906159f8565b6121be9190615a22565b9150505b60608201516121d19088615fca565b96508160800151866121e39190615fca565b82519096506000906122029063ffffffff16662386f26fc100006159f8565b90508082101561222157612216818a615a0f565b985050505050612272565b6000836020015163ffffffff16662386f26fc1000061224091906159f8565b90508083111561226057612254818b615a0f565b99505050505050612272565b61226a838b615a0f565b995050505050505b600101611e7b565b5050955095509592505050565b60008063ffffffff831661229c6080866159f8565b6122a887610220615a0f565b6122b29190615a0f565b6122bc9190615a0f565b67ffffffffffffffff88166000908152600b6020526040812080549293509171010000000000000000000000000000000000810463ffffffff169061231e907501000000000000000000000000000000000000000000900461ffff16856159f8565b6123289190615a0f565b825490915077010000000000000000000000000000000000000000000000900461ffff166123666dffffffffffffffffffffffffffff8a16836159f8565b61237091906159f8565b61238090655af3107a40006159f8565b9998505050505050505050565b6000546001600160a01b031633146123e75760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610a8f565b565b60208101516001600160a01b031661242d576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516005805473ffffffffffffffffffffffffffffffffffffffff199081166001600160a01b03938416179091556020808401516006805484169185169190911790556040808501516007805490941690851617909255815160a0810183527f0000000000000000000000000000000000000000000000000000000000000000841681527f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16918101919091527f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff16818301527f0000000000000000000000000000000000000000000000000000000000000000831660608201527f00000000000000000000000000000000000000000000000000000000000000009092166080830152517f7d6ca86e1084c492475f59ce7b74c37ad96a7440eb9c217ff5b27c183957b9b8916111f3918490616002565b6000546001600160a01b031633148015906125bb57506002546001600160a01b03163314155b156123e7576040517ffbdb8e5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b825181101561294b57600083828151811061261257612612615ada565b6020026020010151905060008160000151905060005b82602001515181101561293d5760008360200151828151811061264d5761264d615ada565b602002602001015160200151905060008460200151838151811061267357612673615ada565b60200260200101516000015190506020826080015163ffffffff1610156126e35760808201516040517f24ecdc020000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015263ffffffff9091166024820152604401610a8f565b67ffffffffffffffff84166000818152600d602090815260408083206001600160a01b0386168085529083529281902086518154938801518389015160608a015160808b015160a08c015160c08d01511515730100000000000000000000000000000000000000027fffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffff9115157201000000000000000000000000000000000000027fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff63ffffffff9485166e01000000000000000000000000000002167fffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffffff9585166a0100000000000000000000027fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff61ffff9098166801000000000000000002979097167fffffffffffffffffffffffffffffffffffff000000000000ffffffffffffffff988616640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000909d1695909916949094179a909a179590951695909517929092171617949094171692909217909155519091907f16a6faa936552870f38ad6586ca4ae10b5d085667b357895aebb320becccf8d49061292b908690600060e08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015260c0830151151560c083015292915050565b60405180910390a35050600101612628565b5050508060010190506125f5565b5060005b81518110156110a257600082828151811061296c5761296c615ada565b6020026020010151600001519050600083838151811061298e5761298e615ada565b60209081029190910181015181015167ffffffffffffffff84166000818152600d845260408082206001600160a01b0385168084529552808220805473ffffffffffffffffffffffffffffffffffffffff191690555192945090917ffa22e84f9c809b5b7e94f084eb45cf17a5e4703cecef8f27ed35e54b719bffcd9190a3505060010161294f565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152612aa582606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff1642612a899190616091565b85608001516fffffffffffffffffffffffffffffffff166141e4565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526110a290849061420c565b80516040811115612b86576040517fb5a10cfa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e546c01000000000000000000000000900463ffffffff1615801590612bd45750600e5463ffffffff6c010000000000000000000000008204166bffffffffffffffffffffffff90911610155b15612be157612be16119e6565b6000612bed6008612e93565b90505b8015612c2f576000612c0e612c06600184616091565b600890612e9e565b509050612c1c6008826142f1565b505080612c28906160a4565b9050612bf0565b506000805b82811015612d37576000848281518110612c5057612c50615ada565b60200260200101516000015190506000858381518110612c7257612c72615ada565b60200260200101516020015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b03161480612cc757506001600160a01b038216155b15612d09576040517f4de938d10000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610a8f565b612d1960088361ffff8416614306565b50612d2861ffff821685615fca565b93505050806001019050612c34565b50600e80547fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd2490612da990839086906160d9565b60405180910390a1505050565b60005b8151811015610e8b576000828281518110612dd657612dd6615ada565b60200260200101516000015190506000838381518110612df857612df8615ada565b6020908102919091018101518101516001600160a01b0384166000818152600c845260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff85169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a25050600101612db9565b6000610e5f8261431c565b6000808080612ead8686614327565b909450925050505b9250929050565b8154600090612ed890600160801b900463ffffffff1642616091565b90508015612f555760018301548354612f13916fffffffffffffffffffffffffffffffff808216928116918591600160801b909104166141e4565b83546fffffffffffffffffffffffffffffffff9190911673ffffffffffffffffffffffffffffffffffffffff1990911617600160801b4263ffffffff16021783555b60208201518354612f7b916fffffffffffffffffffffffffffffffff9081169116614352565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff9283161717845560208301516040808501518316600160801b0291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990612da99084908151151581526020808301516fffffffffffffffffffffffffffffffff90811691830191909152604092830151169181019190915260600190565b604080516101a08101825260008082526020820181905281830181905260608083018290526080830182905260a0830182905260c0830182905260e0830182905261010083018290526101208301819052610140830181905261016083015261018082015290517f58babe3300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff861660048201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906358babe3390602401602060405180830381865afa158015613140573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131649190615b4a565b156131a7576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610a8f565b6001600160a01b0382166131e7576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005546001600160a01b0316331461322b576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855460ff16613272576040517f99ac52f200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610a8f565b600061328160808601866158fb565b15905061329d57613298610abe60808701876158fb565b6132b5565b6001870154640100000000900467ffffffffffffffff165b905060006132c66040870187615942565b91506132e49050876132db60208901896158fb565b90508484611d64565b801561344a576000805b82811015613438576133036040890189615942565b8281811061331357613313615ada565b90506040020160200135600003613356576040517f5cf0444900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff89166000908152600d6020526040808220919061337e908b018b615942565b8481811061338e5761338e615ada565b6133a49260206040909202019081019150614be5565b6001600160a01b031681526020810191909152604001600020547201000000000000000000000000000000000000900460ff1615613430576134236133ec60408a018a615942565b838181106133fc576133fc615ada565b9050604002018036038101906134129190615b10565b6006546001600160a01b0316614368565b61342d9083615a0f565b91505b6001016132ee565b5080156134485761344881614489565b505b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166134846080880160608901614be5565b6001600160a01b0316036134e857600e80548691906000906134b59084906bffffffffffffffffffffffff166160f8565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff160217905550613607565b6006546001600160a01b03166241e5be6135086080890160608a01614be5565b60405160e083901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b039182166004820152602481018990527f00000000000000000000000000000000000000000000000000000000000000009091166044820152606401602060405180830381865afa158015613594573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135b89190615a61565b600e80546000906135d89084906bffffffffffffffffffffffff166160f8565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055505b600e546bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811691161115613674576040517fe5c7a49100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516101a08101825267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b03861660208201529081016137056136cb89806158fb565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061397c92505050565b6001600160a01b0316815260200189600201601481819054906101000a900467ffffffffffffffff166137379061611d565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff1681526020018381526020016000151581526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316637a2dbcbc8a886040516020016137ca91906001600160a01b0391909116815260200190565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016137f6929190616144565b6020604051808303816000875af1158015613815573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138399190616167565b67ffffffffffffffff1681526020016138586080890160608a01614be5565b6001600160a01b0316815260200186815260200187806020019061387c91906158fb565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506020016138c36040890189615942565b808060200260200160405190810160405280939291908181526020016000905b8282101561390f5761390060408302860136819003810190615b10565b815260200190600101906138e3565b505050505081526020018267ffffffffffffffff81111561393257613932614d4e565b60405190808252806020026020018201604052801561396557816020015b60608152602001906001900390816139505790505b508152600060209091015298975050505050505050565b600081516020146139bb57816040517f8d666f60000000000000000000000000000000000000000000000000000000008152600401610a8f9190614cb2565b610e5f828060200190518101906139d29190615a61565b614496565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b6101000151604051602001613a6d9897969594939291906001600160a01b039889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b6040516020818303038152906040528051906020012085610120015180519060200120866101400151604051602001613aa69190616184565b60405160208183030381529060405280519060200120876101600151604051602001613ad29190616197565b60408051601f198184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b60005b8151811015610e8b576000828281518110613b5257613b52615ada565b602002602001015190506000838381518110613b7057613b70615ada565b60200260200101516000015190508067ffffffffffffffff16600003613bce576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610a8f565b6000600b60008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002090506000836040015190506000604051806080016040528086602001518152602001836001600160a01b031681526020018460020160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020018460030154815250905080600001518360000160008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101a082015181600101600c6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101c08201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555090505082600301546000801b03613fcc57604080517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3602082015267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811692820192909252908516606082015230608082015260a00160408051601f1981840301815291905280516020909101206060820181905260038401556001600160a01b03821615613f855760028301805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0384161790555b8367ffffffffffffffff167f7a70081ee29c1fc27898089ba2a5fc35ac0106b043c82ccecd24c6fd48f6ca8684604051613fbf91906161aa565b60405180910390a26140cc565b60028301546001600160a01b03838116911614614021576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610a8f565b60208560200151610160015163ffffffff16101561408557602085015161016001516040517f24ecdc020000000000000000000000000000000000000000000000000000000081526000600482015263ffffffff9091166024820152604401610a8f565b8367ffffffffffffffff167f944eb884a589931130671ee4a7379fbe5fe65ed605a048ba99c454582f2460b086602001516040516140c39190616336565b60405180910390a25b5050505050806001019050613b35565b6000610f17836001600160a01b038416614502565b336001600160a01b038216036141495760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610a8f565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000670de0b6b3a76400006141da837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff86166159f8565b610f179190615a22565b6000614203856141f484866159f8565b6141fe9087615a0f565b614352565b95945050505050565b6000614261826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661450e9092919063ffffffff16565b8051909150156110a2578080602001905181019061427f9190615b4a565b6110a25760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610a8f565b6000610f17836001600160a01b03841661451d565b60006119cd846001600160a01b0385168461453a565b6000610e5f82614557565b600080806143358585614561565b600081815260029690960160205260409095205494959350505050565b60008183106143615781610f17565b5090919050565b81516040517fd02641a00000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009182919084169063d02641a0906024016040805180830381865afa1580156143ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906143f29190616345565b5190507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811660000361445b5783516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a8f565b60208401516119cd907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8316906141a7565b610e76600382600061456d565b60006001600160a01b038211806144ad5750600a82105b156144fe5760408051602081018490520160408051601f19818403018152908290527f8d666f60000000000000000000000000000000000000000000000000000000008252610a8f91600401614cb2565b5090565b6000610f178383614888565b60606119cd84846000856148a0565b60008181526002830160205260408120819055610f178383614992565b600082815260028401602052604081208290556119cd848461499e565b6000610e5f825490565b6000610f1783836149aa565b825474010000000000000000000000000000000000000000900460ff161580614594575081155b1561459e57505050565b825460018401546fffffffffffffffffffffffffffffffff808316929116906000906145d790600160801b900463ffffffff1642616091565b9050801561467d5781831115614619576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600186015461464690839085908490600160801b90046fffffffffffffffffffffffffffffffff166141e4565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16600160801b4263ffffffff160217875592505b8482101561471a576001600160a01b0384166146cf576040517ff94ebcd10000000000000000000000000000000000000000000000000000000081526004810183905260248101869052604401610a8f565b6040517f1a76572a00000000000000000000000000000000000000000000000000000000815260048101839052602481018690526001600160a01b0385166044820152606401610a8f565b8483101561480657600186810154600160801b90046fffffffffffffffffffffffffffffffff169060009082906147519082616091565b61475b878a616091565b6147659190615a0f565b61476f9190615a22565b90506001600160a01b0386166147bb576040517f15279c080000000000000000000000000000000000000000000000000000000081526004810182905260248101869052604401610a8f565b6040517fd0c8d23a00000000000000000000000000000000000000000000000000000000815260048101829052602481018690526001600160a01b0387166044820152606401610a8f565b6148108584616091565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b60008181526001830160205260408120541515610f17565b6060824710156149185760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610a8f565b600080866001600160a01b031685876040516149349190616378565b60006040518083038185875af1925050503d8060008114614971576040519150601f19603f3d011682016040523d82523d6000602084013e614976565b606091505b5091509150614987878383876149d4565b979650505050505050565b6000610f178383614a4d565b6000610f178383614b47565b60008260000182815481106149c1576149c1615ada565b9060005260206000200154905092915050565b60608315614a43578251600003614a3c576001600160a01b0385163b614a3c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610a8f565b50816119cd565b6119cd8383614b96565b60008181526001830160205260408120548015614b36576000614a71600183616091565b8554909150600090614a8590600190616091565b9050818114614aea576000866000018281548110614aa557614aa5615ada565b9060005260206000200154905080876000018481548110614ac857614ac8615ada565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080614afb57614afb616394565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610e5f565b6000915050610e5f565b5092915050565b6000818152600183016020526040812054614b8e57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610e5f565b506000610e5f565b815115614ba65781518083602001fd5b8060405162461bcd60e51b8152600401610a8f9190614cb2565b6001600160a01b0381168114610e7657600080fd5b8035614be081614bc0565b919050565b600060208284031215614bf757600080fd5b8135610f1781614bc0565b60a08101610e5f82846001600160a01b0380825116835267ffffffffffffffff60208301511660208401526bffffffffffffffffffffffff6040830151166040840152806060830151166060840152806080830151166080840152505050565b60005b83811015614c7d578181015183820152602001614c65565b50506000910152565b60008151808452614c9e816020860160208601614c62565b601f01601f19169290920160200192915050565b602081526000610f176020830184614c86565b67ffffffffffffffff81168114610e7657600080fd5b8035614be081614cc5565b600060a08284031215614cf857600080fd5b50919050565b60008060408385031215614d1157600080fd5b8235614d1c81614cc5565b9150602083013567ffffffffffffffff811115614d3857600080fd5b614d4485828601614ce6565b9150509250929050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715614d8757614d87614d4e565b60405290565b6040805190810167ffffffffffffffff81118282101715614d8757614d87614d4e565b60405160e0810167ffffffffffffffff81118282101715614d8757614d87614d4e565b6040516101e0810167ffffffffffffffff81118282101715614d8757614d87614d4e565b604051601f8201601f1916810167ffffffffffffffff81118282101715614e2057614e20614d4e565b604052919050565b600060608284031215614e3a57600080fd5b614e42614d64565b8235614e4d81614bc0565b81526020830135614e5d81614bc0565b60208201526040830135614e7081614bc0565b60408201529392505050565b600067ffffffffffffffff821115614e9657614e96614d4e565b5060051b60200190565b63ffffffff81168114610e7657600080fd5b8035614be081614ea0565b803561ffff81168114614be057600080fd5b8015158114610e7657600080fd5b8035614be081614ecf565b600082601f830112614ef957600080fd5b81356020614f0e614f0983614e7c565b614df7565b82815260069290921b84018101918181019086841115614f2d57600080fd5b8286015b84811015614f7e5760408189031215614f4a5760008081fd5b614f52614d8d565b8135614f5d81614cc5565b815281850135614f6c81614bc0565b81860152835291830191604001614f31565b509695505050505050565b60008060408385031215614f9c57600080fd5b67ffffffffffffffff83351115614fb257600080fd5b83601f843585010112614fc457600080fd5b614fd4614f098435850135614e7c565b8335840180358083526020808401939260059290921b90910101861015614ffa57600080fd5b602085358601015b85358601803560051b016020018110156151f25767ffffffffffffffff8135111561502c57600080fd5b6040601f1982358835890101890301121561504657600080fd5b61504e614d8d565b6150616020833589358a01010135614cc5565b863587018235016020810135825267ffffffffffffffff604090910135111561508957600080fd5b86358701823501604081013501603f810189136150a557600080fd5b6150b5614f096020830135614e7c565b602082810135808352908201919060081b83016040018b10156150d757600080fd5b604083015b6040602085013560081b8501018110156151d957610100818d03121561510157600080fd5b615109614d8d565b6151138235614bc0565b8135815260e0601f19838f0301121561512b57600080fd5b615133614db0565b6151406020840135614ea0565b602083013581526151546040840135614ea0565b6040830135602082015261516a60608401614ebd565b604082015261517c6080840135614ea0565b6080830135606082015261519360a0840135614ea0565b60a083013560808201526151a960c08401614edd565b60a08201526151ba60e08401614edd565b60c08201526020828101919091529084529290920191610100016150dc565b5060208481019190915292865250509283019201615002565b5092505067ffffffffffffffff6020840135111561520f57600080fd5b61521f8460208501358501614ee8565b90509250929050565b6000806040838503121561523b57600080fd5b823561524681614cc5565b9150602083013561525681614bc0565b809150509250929050565b6000806040838503121561527457600080fd5b823561524681614bc0565b60006020828403121561529157600080fd5b8135610f1781614cc5565b80511515825260208101516152b7602084018261ffff169052565b5060408101516152cf604084018263ffffffff169052565b5060608101516152e7606084018263ffffffff169052565b5060808101516152ff608084018263ffffffff169052565b5060a081015161531560a084018261ffff169052565b5060c081015161532d60c084018263ffffffff169052565b5060e081015161534360e084018261ffff169052565b506101008181015161ffff9081169184019190915261012080830151909116908301526101408082015163ffffffff90811691840191909152610160808301518216908401526101808083015167ffffffffffffffff908116918501919091526101a080840151909116908401526101c09182015116910152565b6000610240820190506153d282845161529c565b60208301516001600160a01b03166101e0830152604083015167ffffffffffffffff166102008301526060909201516102209091015290565b60608101610e5f828480516001600160a01b03908116835260208083015182169084015260409182015116910152565b6000806020838503121561544e57600080fd5b823567ffffffffffffffff8082111561546657600080fd5b818501915085601f83011261547a57600080fd5b81358181111561548957600080fd5b8660208260061b850101111561549e57600080fd5b60209290920196919550909350505050565b600060208083850312156154c357600080fd5b823567ffffffffffffffff8111156154da57600080fd5b8301601f810185136154eb57600080fd5b80356154f9614f0982614e7c565b81815260069190911b8201830190838101908783111561551857600080fd5b928401925b8284101561498757604084890312156155365760008081fd5b61553e614d8d565b843561554981614bc0565b81528486013561555881614cc5565b818701528252604093909301929084019061551d565b60008151808452602080850194506020840160005b838110156155b757815180516001600160a01b0316885283015161ffff168388015260409096019590820190600101615583565b509495945050505050565b6040815260006155d5604083018561556e565b90508260208301529392505050565b80356fffffffffffffffffffffffffffffffff81168114614be057600080fd5b60006060828403121561561657600080fd5b61561e614d64565b823561562981614ecf565b8152615637602084016155e4565b6020820152614e70604084016155e4565b6000806000806080858703121561565e57600080fd5b843561566981614cc5565b9350602085013567ffffffffffffffff81111561568557600080fd5b61569187828801614ce6565b9350506040850135915060608501356156a981614bc0565b939692955090935050565b600060208083850312156156c757600080fd5b823567ffffffffffffffff8111156156de57600080fd5b8301601f810185136156ef57600080fd5b80356156fd614f0982614e7c565b818152610220918202830184019184820191908884111561571d57600080fd5b938501935b838510156158a2578489038181121561573b5760008081fd5b615743614d64565b863561574e81614cc5565b81526101e0601f1983018113156157655760008081fd5b61576d614dd3565b925061577a898901614edd565b83526040615789818a01614ebd565b8a850152606061579a818b01614eb2565b8286015260806157ab818c01614eb2565b8287015260a091506157be828c01614eb2565b9086015260c06157cf8b8201614ebd565b8287015260e091506157e2828c01614eb2565b908601526101006157f48b8201614ebd565b828701526101209150615808828c01614ebd565b9086015261014061581a8b8201614ebd565b82870152610160915061582e828c01614eb2565b908601526101806158408b8201614eb2565b828701526101a09150615854828c01614cdb565b908601526101c06158668b8201614cdb565b82870152615875848c01614eb2565b818701525050838a84015261588d6102008a01614bd5565b90830152508452509384019391850191615722565b50979650505050505050565b6020808252825182820181905260009190848201906040850190845b818110156158ef5783516001600160a01b0316835292840192918401916001016158ca565b50909695505050505050565b6000808335601e1984360301811261591257600080fd5b83018035915067ffffffffffffffff82111561592d57600080fd5b602001915036819003821315612eb557600080fd5b6000808335601e1984360301811261595957600080fd5b83018035915067ffffffffffffffff82111561597457600080fd5b6020019150600681901b3603821315612eb557600080fd5b80517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81168114614be057600080fd5b600080604083850312156159cb57600080fd5b6159d48361598c565b915061521f6020840161598c565b634e487b7160e01b600052601160045260246000fd5b8082028115828204841417610e5f57610e5f6159e2565b80820180821115610e5f57610e5f6159e2565b600082615a3f57634e487b7160e01b600052601260045260246000fd5b500490565b600060208284031215615a5657600080fd5b8151610f1781614bc0565b600060208284031215615a7357600080fd5b5051919050565b600060408284031215615a8c57600080fd5b615a94614d8d565b8235615a9f81614bc0565b8152615aad60208401614ebd565b60208201529392505050565b67ffffffffffffffff818116838216019080821115614b4057614b406159e2565b634e487b7160e01b600052603260045260246000fd5b8181036000831280158383131683831282161715614b4057614b406159e2565b600060408284031215615b2257600080fd5b615b2a614d8d565b8235615b3581614bc0565b81526020928301359281019290925250919050565b600060208284031215615b5c57600080fd5b8151610f1781614ecf565b602081526000825160a06020840152615b8360c0840182614c86565b905067ffffffffffffffff602085015116604084015260408401516001600160a01b038082166060860152606086015160808601528060808701511660a086015250508091505092915050565b600082601f830112615be157600080fd5b815167ffffffffffffffff811115615bfb57615bfb614d4e565b615c0e6020601f19601f84011601614df7565b818152846020838601011115615c2357600080fd5b6119cd826020830160208701614c62565b600060208284031215615c4657600080fd5b815167ffffffffffffffff80821115615c5e57600080fd5b9083019060408286031215615c7257600080fd5b615c7a614d8d565b825182811115615c8957600080fd5b615c9587828601615bd0565b825250602083015182811115615caa57600080fd5b615cb687828601615bd0565b60208301525095945050505050565b602081526000825160606020840152615ce16080840182614c86565b90506020840151601f1980858403016040860152615cff8383614c86565b92506040860151915080858403016060860152506142038282614c86565b60008151808452602080850194506020840160005b838110156155b757815180516001600160a01b031688528301518388015260409096019590820190600101615d32565b60008282518085526020808601955060208260051b8401016020860160005b84811015615daf57601f19868403018952615d9d838351614c86565b98840198925090830190600101615d81565b5090979650505050505050565b60208152615dd760208201835167ffffffffffffffff169052565b60006020830151615df360408401826001600160a01b03169052565b5060408301516001600160a01b038116606084015250606083015167ffffffffffffffff8116608084015250608083015160a083015260a0830151615e3c60c084018215159052565b5060c083015167ffffffffffffffff811660e08401525060e0830151610100615e6f818501836001600160a01b03169052565b840151610120848101919091528401516101a061014080860182905291925090615e9d6101c0860184614c86565b9250808601519050601f19610160818786030181880152615ebe8584615d1d565b945080880151925050610180818786030181880152615edd8584615d62565b970151959092019490945250929392505050565b6bffffffffffffffffffffffff828116828216039080821115614b4057614b406159e2565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015615f565780818660040360031b1b83161692505b505092915050565b60008085851115615f6e57600080fd5b83861115615f7b57600080fd5b5050820193919092039150565b600060208284031215615f9a57600080fd5b6040516020810181811067ffffffffffffffff82111715615fbd57615fbd614d4e565b6040529135825250919050565b63ffffffff818116838216019080821115614b4057614b406159e2565b600060208284031215615ff957600080fd5b610f178261598c565b610100810161606382856001600160a01b0380825116835267ffffffffffffffff60208301511660208401526bffffffffffffffffffffffff6040830151166040840152806060830151166060840152806080830151166080840152505050565b82516001600160a01b0390811660a08401526020840151811660c084015260408401511660e0830152610f17565b81810381811115610e5f57610e5f6159e2565b6000816160b3576160b36159e2565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b63ffffffff831681526040602082015260006119cd604083018461556e565b6bffffffffffffffffffffffff818116838216019080821115614b4057614b406159e2565b600067ffffffffffffffff80831681810361613a5761613a6159e2565b6001019392505050565b67ffffffffffffffff831681526040602082015260006119cd6040830184614c86565b60006020828403121561617957600080fd5b8151610f1781614cc5565b602081526000610f176020830184615d1d565b602081526000610f176020830184615d62565b815460ff81161515825261024082019061ffff600882901c8116602085015263ffffffff601883901c811660408601526161f160608601828560381c1663ffffffff169052565b61620860808601828560581c1663ffffffff169052565b61621d60a08601838560781c1661ffff169052565b61623460c08601828560881c1663ffffffff169052565b61624960e08601838560a81c1661ffff169052565b61625f6101008601838560b81c1661ffff169052565b6162756101208601838560c81c1661ffff169052565b61628d6101408601828560d81c1663ffffffff169052565b600186015463ffffffff82821616610160870152925067ffffffffffffffff602084901c811661018087015291506162d76101a08601838560601c1667ffffffffffffffff169052565b6162ef6101c08601828560a01c1663ffffffff169052565b5060028501546001600160a01b0381166101e086015291506163236102008501828460a01c1667ffffffffffffffff169052565b5050600383015461022083015292915050565b6101e08101610e5f828461529c565b60006040828403121561635757600080fd5b61635f614d8d565b6163688361598c565b81526020830151615aad81614ea0565b6000825161638a818460208701614c62565b9190910192915050565b634e487b7160e01b600052603160045260246000fdfea164736f6c6343000818000a",
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
	return common.HexToHash("0xd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd")
}

func (EVM2EVMMultiOnRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (EVM2EVMMultiOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x7d6ca86e1084c492475f59ce7b74c37ad96a7440eb9c217ff5b27c183957b9b8")
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
