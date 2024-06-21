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
	Router           common.Address
	PriceRegistry    common.Address
	MessageValidator common.Address
	FeeAggregator    common.Address
}

type EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthArgs struct {
	Token                      common.Address
	PremiumMultiplierWeiPerEth uint64
}

type EVM2EVMMultiOnRampStaticConfig struct {
	LinkToken          common.Address
	ChainSelector      uint64
	MaxFeeJuelsPerMsg  *big.Int
	RmnProxy           common.Address
	TokenAdminRegistry common.Address
}

type EVM2EVMMultiOnRampTokenTransferFeeConfig struct {
	MinFeeUSDCents    uint32
	MaxFeeUSDCents    uint32
	DeciBps           uint16
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	IsEnabled         bool
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

var EVM2EVMMultiOnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CannotSendZeroTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"DestinationChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"}],\"name\":\"InvalidDestBytesOverhead\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidDestChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgFeeJuels\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint256\"}],\"name\":\"MessageFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"errorReason\",\"type\":\"bytes\"}],\"name\":\"MessageValidationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NotAFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SourceTokenDataTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainDynamicConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeValueJuels\",\"type\":\"uint256\"}],\"name\":\"FeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeTokenWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"name\":\"PremiumMultiplierWeiPerEthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destChainSelector\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenTransferFeeConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"name\":\"TokenTransferFeeConfigUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyDestChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyPremiumMultiplierWeiPerEthUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfigRemoveArgs[]\",\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\"}],\"name\":\"applyTokenTransferFeeConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestChainConfig\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DestChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPoolV1\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPremiumMultiplierWeiPerEth\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenTransferFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeAggregator\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b50604051620069d8380380620069d8833981016040819052620000359162001463565b33806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf81620001f2565b505085516001600160a01b031615905080620000e6575060208501516001600160401b0316155b80620000fd575060608501516001600160a01b0316155b8062000114575060808501516001600160a01b0316155b1562000133576040516306b7c75960e31b815260040160405180910390fd5b84516001600160a01b0390811660a05260208601516001600160401b031660c05260408601516001600160601b031660809081526060870151821660e052860151166101005262000184846200029d565b6200018f8362000447565b6200019a82620009a4565b60408051600080825260208201909252620001e791839190620001e0565b6040805180820190915260008082526020820152815260200190600190039081620001b85790505b5062000a70565b505050505062001873565b336001600160a01b038216036200024c5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60208101516001600160a01b03161580620002c3575060608101516001600160a01b0316155b15620002e2576040516306b7c75960e31b815260040160405180910390fd5b8051600280546001600160a01b039283166001600160a01b0319918216179091556020808401516003805491851691841691909117905560408085015160048054918616918516919091179055606080860151600580549187169190951617909355805160a08082018352518516815260c0516001600160401b031692810192909252608080516001600160601b03168383015260e051851693830193909352610100519093169181019190915290517f823ad577168c60d73b7419edfe54229b4737c81cc2eb4db07b2b6604c040429f916200043c91849082516001600160a01b0390811682526020808501516001600160401b0316818401526040808601516001600160601b0316818501526060808701518416818601526080968701518416968501969096528451831660a085015290840151821660c0840152830151811660e08301529190920151166101008201526101200190565b60405180910390a150565b60005b8151811015620009a05760008282815181106200046b576200046b62001596565b6020026020010151905060008383815181106200048c576200048c62001596565b6020026020010151600001519050806001600160401b031660001480620004c35750602082015161018001516001600160401b0316155b15620004ee5760405163c35aa79d60e01b81526001600160401b038216600482015260240162000083565b600060066000836001600160401b03166001600160401b0316815260200190815260200160002090506000836040015190506000604051806080016040528086602001518152602001836001600160a01b031681526020018460020160149054906101000a90046001600160401b03166001600160401b031681526020018460030154815250905080600001518360000160008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a8154816001600160401b0302191690836001600160401b031602179055506101a082015181600101600c6101000a8154816001600160401b0302191690836001600160401b031602179055506101c08201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555090505082600301546000801b03620008bd5760c051604080517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b360208201526001600160401b0392831691810191909152908516606082015230608082015260a00160408051601f1981840301815291905280516020909101206060820181905260038401556001600160a01b0382161562000874576002830180546001600160a01b0319166001600160a01b0384161790555b836001600160401b03167f7a70081ee29c1fc27898089ba2a5fc35ac0106b043c82ccecd24c6fd48f6ca8684604051620008af9190620015ac565b60405180910390a26200098f565b60028301546001600160a01b03838116911614620008fa5760405163c35aa79d60e01b81526001600160401b038516600482015260240162000083565b60208560200151610160015163ffffffff1610156200094757602085015161016001516040516312766e0160e11b81526000600482015263ffffffff909116602482015260440162000083565b836001600160401b03167f944eb884a589931130671ee4a7379fbe5fe65ed605a048ba99c454582f2460b0866020015160405162000986919062001740565b60405180910390a25b50505050508060010190506200044a565b5050565b60005b8151811015620009a0576000828281518110620009c857620009c862001596565b60200260200101516000015190506000838381518110620009ed57620009ed62001596565b6020908102919091018101518101516001600160a01b03841660008181526007845260409081902080546001600160401b0319166001600160401b0385169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a25050600101620009a7565b60005b825181101562000cdd57600083828151811062000a945762000a9462001596565b6020026020010151905060008160000151905060005b82602001515181101562000cce5760008360200151828151811062000ad35762000ad362001596565b602002602001015160200151905060008460200151838151811062000afc5762000afc62001596565b60200260200101516000015190506020826080015163ffffffff16101562000b555760808201516040516312766e0160e11b81526001600160a01b038316600482015263ffffffff909116602482015260440162000083565b6001600160401b03841660008181526008602090815260408083206001600160a01b0386168085529083529281902086518154938801518389015160608a015160808b015160a08c01511515600160901b0260ff60901b1963ffffffff928316600160701b021664ffffffffff60701b199383166a01000000000000000000000263ffffffff60501b1961ffff90961668010000000000000000029590951665ffffffffffff60401b19968416640100000000026001600160401b0319909b16939097169290921798909817939093169390931717919091161792909217909155519091907f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b59062000cbb908690600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60405180910390a3505060010162000aaa565b50505080600101905062000a73565b5060005b815181101562000da357600082828151811062000d025762000d0262001596565b6020026020010151600001519050600083838151811062000d275762000d2762001596565b6020908102919091018101518101516001600160401b03841660008181526008845260408082206001600160a01b038516808452955280822080546001600160981b03191690555192945090917ffa22e84f9c809b5b7e94f084eb45cf17a5e4703cecef8f27ed35e54b719bffcd9190a3505060010162000ce1565b505050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b038111828210171562000de35762000de362000da8565b60405290565b6040516101e081016001600160401b038111828210171562000de35762000de362000da8565b604080519081016001600160401b038111828210171562000de35762000de362000da8565b60405160c081016001600160401b038111828210171562000de35762000de362000da8565b60405160a081016001600160401b038111828210171562000de35762000de362000da8565b604051601f8201601f191681016001600160401b038111828210171562000ea95762000ea962000da8565b604052919050565b80516001600160a01b038116811462000ec957600080fd5b919050565b80516001600160401b038116811462000ec957600080fd5b60006080828403121562000ef957600080fd5b604051608081016001600160401b038111828210171562000f1e5762000f1e62000da8565b60405290508062000f2f8362000eb1565b815262000f3f6020840162000eb1565b602082015262000f526040840162000eb1565b604082015262000f656060840162000eb1565b60608201525092915050565b60006001600160401b0382111562000f8d5762000f8d62000da8565b5060051b60200190565b8051801515811462000ec957600080fd5b805161ffff8116811462000ec957600080fd5b805163ffffffff8116811462000ec957600080fd5b600082601f83011262000fe257600080fd5b8151602062000ffb62000ff58362000f71565b62000e7e565b82815261022092830285018201928282019190878511156200101c57600080fd5b8387015b85811015620011c657808903828112156200103b5760008081fd5b6200104562000dbe565b620010508362000ece565b81526101e080601f1984011215620010685760008081fd5b6200107262000de9565b92506200108188850162000f97565b835260406200109281860162000fa8565b898501526060620010a581870162000fbb565b828601526080620010b881880162000fbb565b8287015260a09150620010cd82880162000fbb565b9086015260c0620010e087820162000fa8565b8287015260e09150620010f582880162000fbb565b908601526101006200110987820162000fa8565b8287015261012091506200111f82880162000fa8565b908601526101406200113387820162000fa8565b8287015261016091506200114982880162000fbb565b908601526101806200115d87820162000fbb565b828701526101a091506200117382880162000ece565b908601526101c06200118787820162000ece565b828701526200119884880162000fbb565b8187015250508389840152620011b2610200860162000eb1565b908301525085525092840192810162001020565b5090979650505050505050565b600082601f830112620011e557600080fd5b81516020620011f862000ff58362000f71565b82815260069290921b840181019181810190868411156200121857600080fd5b8286015b848110156200126e5760408189031215620012375760008081fd5b6200124162000e0f565b6200124c8262000eb1565b81526200125b85830162000ece565b818601528352918301916040016200121c565b509695505050505050565b600082601f8301126200128b57600080fd5b815160206200129e62000ff58362000f71565b82815260059290921b84018101918181019086841115620012be57600080fd5b8286015b848110156200126e5780516001600160401b0380821115620012e357600080fd5b908801906040601f19838c038101821315620012fe57600080fd5b6200130862000e0f565b6200131589860162000ece565b815282850151848111156200132957600080fd5b8086019550508c603f8601126200133f57600080fd5b8885015193506200135462000ff58562000f71565b84815260e09094028501830193898101908e8611156200137357600080fd5b958401955b858710156200144c57868f0360e08112156200139357600080fd5b6200139d62000e0f565b620013a88962000eb1565b815260c08683011215620013bb57600080fd5b620013c562000e34565b9150620013d48d8a0162000fbb565b8252620013e3878a0162000fbb565b8d830152620013f560608a0162000fa8565b878301526200140760808a0162000fbb565b60608301526200141a60a08a0162000fbb565b60808301526200142d60c08a0162000f97565b60a0830152808d0191909152825260e09690960195908a019062001378565b828b015250875250505092840192508301620012c2565b60008060008060008587036101808112156200147e57600080fd5b60a08112156200148d57600080fd5b506200149862000e59565b620014a38762000eb1565b8152620014b36020880162000ece565b602082015260408701516001600160601b0381168114620014d357600080fd5b6040820152620014e66060880162000eb1565b6060820152620014f96080880162000eb1565b608082015294506200150f8760a0880162000ee6565b6101208701519094506001600160401b03808211156200152e57600080fd5b6200153c89838a0162000fd0565b94506101408801519150808211156200155457600080fd5b6200156289838a01620011d3565b93506101608801519150808211156200157a57600080fd5b50620015898882890162001279565b9150509295509295909350565b634e487b7160e01b600052603260045260246000fd5b815460ff81161515825261024082019061ffff600882901c8116602085015263ffffffff601883901c81166040860152620015f460608601828560381c1663ffffffff169052565b6200160c60808601828560581c1663ffffffff169052565b6200162260a08601838560781c1661ffff169052565b6200163a60c08601828560881c1663ffffffff169052565b6200165060e08601838560a81c1661ffff169052565b620016676101008601838560b81c1661ffff169052565b6200167e6101208601838560c81c1661ffff169052565b620016976101408601828560d81c1663ffffffff169052565b600186015463ffffffff8282161661016087015292506001600160401b03602084901c81166101808701529150620016e06101a08601838560601c166001600160401b03169052565b620016f96101c08601828560a01c1663ffffffff169052565b5060028501546001600160a01b0381166101e086015291506200172d6102008501828460a01c166001600160401b03169052565b5050600383015461022083015292915050565b8151151581526101e08101602083015162001761602084018261ffff169052565b5060408301516200177a604084018263ffffffff169052565b50606083015162001793606084018263ffffffff169052565b506080830151620017ac608084018263ffffffff169052565b5060a0830151620017c360a084018261ffff169052565b5060c0830151620017dc60c084018263ffffffff169052565b5060e0830151620017f360e084018261ffff169052565b506101008381015161ffff9081169184019190915261012080850151909116908301526101408084015163ffffffff9081169184019190915261016080850151821690840152610180808501516001600160401b03908116918501919091526101a080860151909116908401526101c09384015116929091019190915290565b60805160a05160c05160e051610100516150c562001913600039600081816102d501528181610f3a01526125290152600081816102a6015281816124ff01526126320152600081816102420152818161249a01528181612b940152613359015260008181610213015281816124750152818161292501526129cf015260008181610272015281816124cc01528181612a910152612b0101526150c56000f3fe608060405234801561001057600080fd5b50600436106101825760003560e01c806379ba5097116100d8578063a69c64c01161008c578063e080bcba11610066578063e080bcba146108c3578063f2fde38b146108d6578063fbca3b74146108e957600080fd5b8063a69c64c01461088a578063a6f3ab6c1461089d578063df0aa9e9146108b057600080fd5b80638b364334116100bd5780638b364334146108535780638da5cb5b146108665780639041be3d1461087757600080fd5b806379ba5097146106e857806382b49eb0146106f057600080fd5b806348a98aa41161013a578063704b6c0211610114578063704b6c021461064f5780637437ff9f14610662578063770e2dc4146106d557600080fd5b806348a98aa4146103865780636def4ce7146103b15780636e9960c31461063e57600080fd5b8063181f5a771161016b578063181f5a771461031257806320487ded1461035b5780633a0199401461037c57600080fd5b8063061877e31461018757806306285c69146101d8575b600080fd5b6101ba6101953660046138d4565b6001600160a01b031660009081526007602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff90911681526020015b60405180910390f35b6103056040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040518060a001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b6040516101cf91906138f1565b61034e6040518060400160405280601c81526020017f45564d3245564d4d756c74694f6e52616d7020312e362e302d6465760000000081525081565b6040516101cf91906139a1565b61036e6103693660046139ed565b610909565b6040519081526020016101cf565b610384610d3b565b005b610399610394366004613a3d565b610eff565b6040516001600160a01b0390911681526020016101cf565b6106316103bf366004613a76565b604080516102608101825260006080820181815260a0830182905260c0830182905260e08301829052610100830182905261012083018290526101408301829052610160830182905261018083018290526101a083018290526101c083018290526101e0830182905261020083018290526102208301829052610240830182905282526020820181905291810182905260608101919091525067ffffffffffffffff908116600090815260066020908152604091829020825161026081018452815460ff811615156080830190815261ffff610100808404821660a086015263ffffffff63010000008504811660c08701526701000000000000008504811660e08701526b01000000000000000000000085048116918601919091526f01000000000000000000000000000000840482166101208601527101000000000000000000000000000000000084048116610140860152750100000000000000000000000000000000000000000084048216610160860152770100000000000000000000000000000000000000000000008404821661018086015279010000000000000000000000000000000000000000000000000084049091166101a08501527b0100000000000000000000000000000000000000000000000000000090920482166101c084015260018401548083166101e0850152640100000000810488166102008501526c01000000000000000000000000810488166102208501527401000000000000000000000000000000000000000090819004909216610240840152825260028301546001600160a01b0381169483019490945290920490931691810191909152600390910154606082015290565b6040516101cf9190613bb5565b600a546001600160a01b0316610399565b61038461065d3660046138d4565b610fae565b6106c860408051608081018252600080825260208201819052918101829052606081019190915250604080516080810182526002546001600160a01b03908116825260035481166020830152600454811692820192909252600554909116606082015290565b6040516101cf9190613c02565b6103846106e3366004613e35565b611018565b61038461102e565b6107f36106fe366004613a3d565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a08101919091525067ffffffffffffffff9190911660009081526008602090815260408083206001600160a01b0394909416835292815290829020825160c081018452905463ffffffff8082168352640100000000820481169383019390935268010000000000000000810461ffff16938201939093526a01000000000000000000008304821660608201526e01000000000000000000000000000083049091166080820152720100000000000000000000000000000000000090910460ff16151560a082015290565b6040516101cf9190600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b6101ba610861366004613a3d565b6110ec565b6000546001600160a01b0316610399565b6101ba610885366004613a76565b6111e5565b6103846108983660046140af565b611229565b6103846108ab36600461416d565b61123d565b61036e6108be3660046141f2565b61124e565b6103846108d136600461425e565b6116d6565b6103846108e43660046138d4565b6116e7565b6108fc6108f7366004613a76565b6116f8565b6040516101cf9190614458565b67ffffffffffffffff82166000908152600660205260408120805460ff1661096e576040517f99ac52f200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526024015b60405180910390fd5b600061097d60808501856144a5565b15905061099e5761099961099460808601866144a5565b61172c565b6109b6565b6001820154640100000000900467ffffffffffffffff165b90506109e0856109c960208701876144a5565b9050836109d960408901896144f3565b90506117d4565b60006007816109f560808801606089016138d4565b6001600160a01b03168152602081019190915260400160009081205467ffffffffffffffff169150819003610a7257610a3460808601606087016138d4565b6040517fa7499d200000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610965565b60035460009081906001600160a01b031663ffdb4b37610a9860808a0160608b016138d4565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015267ffffffffffffffff8b1660248201526044016040805180830381865afa158015610b03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b279190614569565b90925090506000808080610b3e60408c018c6144f3565b90501115610b7957610b6d8b610b5a60808d0160608e016138d4565b87610b6860408f018f6144f3565b6118e4565b91945092509050610bb0565b6001880154610bad9074010000000000000000000000000000000000000000900463ffffffff16662386f26fc100006145c2565b92505b875460009077010000000000000000000000000000000000000000000000900461ffff1615610c1c57610c198c6dffffffffffffffffffffffffffff607088901c16610bff60208f018f6144a5565b90508e8060400190610c1191906144f3565b905086611cd5565b90505b600089600101600c9054906101000a900467ffffffffffffffff1667ffffffffffffffff168463ffffffff168b600001600f9054906101000a900461ffff1661ffff168e8060200190610c6f91906144a5565b610c7a9291506145c2565b8c54610c9b906b010000000000000000000000900463ffffffff168d6145d9565b610ca591906145d9565b610caf91906145d9565b610cc9906dffffffffffffffffffffffffffff89166145c2565b610cd391906145c2565b90507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff87168282610d0a67ffffffffffffffff8c16896145c2565b610d1491906145d9565b610d1e91906145d9565b610d2891906145ec565b9a50505050505050505050505b92915050565b600354604080517fcdc73d5100000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163cdc73d5191600480830192869291908290030181865afa158015610d9d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610dc59190810190614627565b6005549091506001600160a01b031660005b8251811015610efa576000838281518110610df457610df46146b6565b60209081029190910101516040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529091506000906001600160a01b038316906370a0823190602401602060405180830381865afa158015610e62573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8691906146e5565b90508015610ef057610ea26001600160a01b0383168583611ddb565b816001600160a01b0316846001600160a01b03167f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e83604051610ee791815260200190565b60405180910390a35b5050600101610dd7565b505050565b6040517fbbe4f6db0000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063bbe4f6db90602401602060405180830381865afa158015610f83573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa791906146fe565b9392505050565b610fb6611e5b565b600a805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c906020015b60405180910390a150565b611020611e5b565b61102a8282611eba565b5050565b6001546001600160a01b031633146110885760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610965565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b67ffffffffffffffff80831660009081526009602090815260408083206001600160a01b0386168452909152812054909116808203610fa75767ffffffffffffffff84166000908152600660205260409020600201546001600160a01b031680156111dd576040517f856c82470000000000000000000000000000000000000000000000000000000081526001600160a01b03858116600483015282169063856c824790602401602060405180830381865afa1580156111b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111d4919061471b565b92505050610d35565b509392505050565b67ffffffffffffffff8082166000908152600660205260408120600201549091610d3591740100000000000000000000000000000000000000009004166001614738565b611231611e5b565b61123a81612283565b50565b611245612349565b61123a816123a3565b67ffffffffffffffff8416600090815260066020526040812081611275828888888861257e565b905060005b8161014001515181101561166c57600061129760408901896144f3565b838181106112a7576112a76146b6565b9050604002018036038101906112bd9190614760565b905080602001516000036112fd576040517f5cf0444900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061130d8a8360000151610eff565b90506001600160a01b03811615806113c357506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa15801561139d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113c1919061479a565b155b156114085781516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610965565b6000816001600160a01b0316639a4575b96040518060a001604052808d806000019061143491906144a5565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525067ffffffffffffffff8f166020808301919091526001600160a01b03808e16604080850191909152918901516060840152885116608090920191909152517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526114df91906004016147b7565b6000604051808303816000875af11580156114fe573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115269190810190614884565b90506020816020015151118015611587575067ffffffffffffffff8b16600090815260086020908152604080832086516001600160a01b0316845282529091205490820151516e01000000000000000000000000000090910463ffffffff16105b156115cc5782516040517f36f536ca0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610965565b80516115d790612dca565b5060408051606081019091526001600160a01b03831660808201528060a08101604051602081830303815290604052815260200182600001518152602001826020015181525060405160200161162d9190614915565b6040516020818303038152906040528561016001518581518110611653576116536146b6565b602002602001018190525050505080600101905061127a565b5061167b818360030154612e25565b61018082015260405167ffffffffffffffff8816907fc79f9c3e610deac14de4e704195fe17eab0983ee9916866bc04d16a00f54daa6906116bd908490614a20565b60405180910390a261018001519150505b949350505050565b6116de611e5b565b61123a81612f80565b6116ef612349565b61123a81613547565b60606040517f9e7177c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f97a657c9000000000000000000000000000000000000000000000000000000006117598385614b55565b7fffffffff0000000000000000000000000000000000000000000000000000000016146117b2576040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117bf8260048186614b9d565b8101906117cc9190614bc7565b519392505050565b67ffffffffffffffff8416600090815260066020526040902080546301000000900463ffffffff1684111561184d5780546040517f86933789000000000000000000000000000000000000000000000000000000008152630100000090910463ffffffff16600482015260248101859052604401610965565b8054670100000000000000900463ffffffff16831115611899576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8054610100900461ffff168211156118dd576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b6000808083815b81811015611cc8576000878783818110611907576119076146b6565b90506040020180360381019061191d9190614760565b905060006001600160a01b03166119388c8360000151610eff565b6001600160a01b0316036119865780516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610965565b67ffffffffffffffff8b16600090815260086020908152604080832084516001600160a01b03168452825291829020825160c081018452905463ffffffff8082168352640100000000820481169383019390935268010000000000000000810461ffff16938201939093526a01000000000000000000008304821660608201526e01000000000000000000000000000083049091166080820152720100000000000000000000000000000000000090910460ff16151560a08201819052611af45767ffffffffffffffff8c1660009081526006602052604090208054611a9490790100000000000000000000000000000000000000000000000000900461ffff16662386f26fc100006145c2565b611a9e90896145d9565b8154909850611ad2907b01000000000000000000000000000000000000000000000000000000900463ffffffff1688614c09565b6001820154909750611aea9063ffffffff1687614c09565b9550505050611cc0565b604081015160009061ffff1615611c105760008c6001600160a01b031684600001516001600160a01b031614611bb35760035484516040517f4ab35b0b0000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152911690634ab35b0b90602401602060405180830381865afa158015611b88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bac9190614c26565b9050611bb6565b508a5b620186a0836040015161ffff16611bf88660200151847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166135fd90919063ffffffff16565b611c0291906145c2565b611c0c91906145ec565b9150505b6060820151611c1f9088614c09565b9650816080015186611c319190614c09565b8251909650600090611c509063ffffffff16662386f26fc100006145c2565b905080821015611c6f57611c64818a6145d9565b985050505050611cc0565b6000836020015163ffffffff16662386f26fc10000611c8e91906145c2565b905080831115611cae57611ca2818b6145d9565b99505050505050611cc0565b611cb8838b6145d9565b995050505050505b6001016118eb565b5050955095509592505050565b60008063ffffffff8316611cea6080866145c2565b611cf6876102206145d9565b611d0091906145d9565b611d0a91906145d9565b67ffffffffffffffff8816600090815260066020526040812080549293509171010000000000000000000000000000000000810463ffffffff1690611d6c907501000000000000000000000000000000000000000000900461ffff16856145c2565b611d7691906145d9565b825490915077010000000000000000000000000000000000000000000000900461ffff16611db46dffffffffffffffffffffffffffff8a16836145c2565b611dbe91906145c2565b611dce90655af3107a40006145c2565b9998505050505050505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610efa90849061363a565b6000546001600160a01b03163314801590611e815750600a546001600160a01b03163314155b15611eb8576040517ffbdb8e5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60005b82518110156121ac576000838281518110611eda57611eda6146b6565b6020026020010151905060008160000151905060005b82602001515181101561219e57600083602001518281518110611f1557611f156146b6565b6020026020010151602001519050600084602001518381518110611f3b57611f3b6146b6565b60200260200101516000015190506020826080015163ffffffff161015611fab5760808201516040517f24ecdc020000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015263ffffffff9091166024820152604401610965565b67ffffffffffffffff841660008181526008602090815260408083206001600160a01b0386168085529083529281902086518154938801518389015160608a015160808b015160a08c015115157201000000000000000000000000000000000000027fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff63ffffffff9283166e01000000000000000000000000000002167fffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffffff9383166a0100000000000000000000027fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff61ffff9096166801000000000000000002959095167fffffffffffffffffffffffffffffffffffff000000000000ffffffffffffffff9684166401000000000267ffffffffffffffff19909b16939097169290921798909817939093169390931717919091161792909217909155519091907f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b59061218c908690600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60405180910390a35050600101611ef0565b505050806001019050611ebd565b5060005b8151811015610efa5760008282815181106121cd576121cd6146b6565b602002602001015160000151905060008383815181106121ef576121ef6146b6565b60209081029190910181015181015167ffffffffffffffff841660008181526008845260408082206001600160a01b038516808452955280822080547fffffffffffffffffffffffffff000000000000000000000000000000000000001690555192945090917ffa22e84f9c809b5b7e94f084eb45cf17a5e4703cecef8f27ed35e54b719bffcd9190a350506001016121b0565b60005b815181101561102a5760008282815181106122a3576122a36146b6565b602002602001015160000151905060008383815181106122c5576122c56146b6565b6020908102919091018101518101516001600160a01b038416600081815260078452604090819020805467ffffffffffffffff191667ffffffffffffffff85169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a25050600101612286565b6000546001600160a01b03163314611eb85760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610965565b60208101516001600160a01b031615806123c8575060608101516001600160a01b0316155b156123ff576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516002805473ffffffffffffffffffffffffffffffffffffffff199081166001600160a01b03938416179091556020808401516003805484169185169190911790556040808501516004805485169186169190911790556060808601516005805490951690861617909355805160a0810182527f0000000000000000000000000000000000000000000000000000000000000000851681527f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16928101929092527f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff16828201527f00000000000000000000000000000000000000000000000000000000000000008416928201929092527f00000000000000000000000000000000000000000000000000000000000000009092166080830152517f823ad577168c60d73b7419edfe54229b4737c81cc2eb4db07b2b6604c040429f9161100d918490614c41565b604080516101a08101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820183905260e082018390526101008201839052610120820181905261014082018190526101608201526101808101919091526040517f2cbc26bb000000000000000000000000000000000000000000000000000000008152608086901b77ffffffffffffffff000000000000000000000000000000001660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632cbc26bb90602401602060405180830381865afa158015612681573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126a5919061479a565b156126e8576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610965565b6001600160a01b038216612728576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546001600160a01b0316331461276c576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855460ff166127b3576040517f99ac52f200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610965565b60006127c260808601866144a5565b1590506127de576127d961099460808701876144a5565b6127f6565b6001870154640100000000900467ffffffffffffffff165b9050600061280760408701876144f3565b915061282590508761281c60208901896144a5565b905084846117d4565b8015612919576004546001600160a01b03168015612917576040517facdfc7890000000000000000000000000000000000000000000000000000000081526001600160a01b0382169063acdfc78990612884908b908b90600401614d96565b600060405180830381600087803b15801561289e57600080fd5b505af19250505080156128af575060015b612917573d8080156128dd576040519150601f19603f3d011682016040523d82523d6000602084013e6128e2565b606091505b50806040517f09c2532500000000000000000000000000000000000000000000000000000000815260040161096591906139a1565b505b60006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166129556080890160608a016138d4565b6001600160a01b03160361296a575084612a3d565b6003546001600160a01b03166241e5be61298a60808a0160608b016138d4565b60405160e083901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b039182166004820152602481018a90527f00000000000000000000000000000000000000000000000000000000000000009091166044820152606401602060405180830381865afa158015612a16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a3a91906146e5565b90505b612a4d60808801606089016138d4565b6001600160a01b03167f075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f82604051612a8791815260200190565b60405180910390a27f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff16811115612b2e576040517f6a92a483000000000000000000000000000000000000000000000000000000008152600481018290526bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166024820152604401610965565b6000612b3a89876110ec565b612b45906001614738565b67ffffffffffffffff8a811660009081526009602090815260408083206001600160a01b038c1680855290835292819020805467ffffffffffffffff191686861617905580516101a0810182527f0000000000000000000000000000000000000000000000000000000000000000909416845290830191909152919250908101612c0c612bd28b806144a5565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612dca92505050565b6001600160a01b031681526020018b600201601481819054906101000a900467ffffffffffffffff16612c3e90614eb4565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff1681526020018581526020016000151581526020018267ffffffffffffffff168152602001896060016020810190612ca491906138d4565b6001600160a01b03168152602001888152602001898060200190612cc891906144a5565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250602001612d0f60408b018b6144f3565b808060200260200160405190810160405280939291908181526020016000905b82821015612d5b57612d4c60408302860136819003810190614760565b81526020019060010190612d2f565b505050505081526020018467ffffffffffffffff811115612d7e57612d7e613c3e565b604051908082528060200260200182016040528015612db157816020015b6060815260200190600190039081612d9c5790505b50815260006020909101529a9950505050505050505050565b60008151602014612e0957816040517f8d666f6000000000000000000000000000000000000000000000000000000000815260040161096591906139a1565b610d3582806020019051810190612e2091906146e5565b61371f565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b6101000151604051602001612ebb9897969594939291906001600160a01b039889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b6040516020818303038152906040528051906020012085610120015180519060200120866101400151604051602001612ef49190614edb565b60405160208183030381529060405280519060200120876101600151604051602001612f209190614eee565b60408051601f198184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b60005b815181101561102a576000828281518110612fa057612fa06146b6565b602002602001015190506000838381518110612fbe57612fbe6146b6565b60200260200101516000015190508067ffffffffffffffff1660001480612ff657506020820151610180015167ffffffffffffffff16155b15613039576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610965565b6000600660008367ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002090506000836040015190506000604051806080016040528086602001518152602001836001600160a01b031681526020018460020160149054906101000a900467ffffffffffffffff1667ffffffffffffffff1681526020018460030154815250905080600001518360000160008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101a082015181600101600c6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101c08201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555090505082600301546000801b0361343757604080517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3602082015267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811692820192909252908516606082015230608082015260a00160408051601f1981840301815291905280516020909101206060820181905260038401556001600160a01b038216156133f05760028301805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0384161790555b8367ffffffffffffffff167f7a70081ee29c1fc27898089ba2a5fc35ac0106b043c82ccecd24c6fd48f6ca868460405161342a9190614f01565b60405180910390a2613537565b60028301546001600160a01b0383811691161461348c576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610965565b60208560200151610160015163ffffffff1610156134f057602085015161016001516040517f24ecdc020000000000000000000000000000000000000000000000000000000081526000600482015263ffffffff9091166024820152604401610965565b8367ffffffffffffffff167f944eb884a589931130671ee4a7379fbe5fe65ed605a048ba99c454582f2460b0866020015160405161352e919061508d565b60405180910390a25b5050505050806001019050612f83565b336001600160a01b0382160361359f5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610965565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000670de0b6b3a7640000613630837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff86166145c2565b610fa791906145ec565b600061368f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661378c9092919063ffffffff16565b805190915015610efa57808060200190518101906136ad919061479a565b610efa5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610965565b60006001600160a01b03821180613737575061040082105b156137885760408051602081018490520160408051601f19818403018152908290527f8d666f60000000000000000000000000000000000000000000000000000000008252610965916004016139a1565b5090565b60606116ce848460008585600080866001600160a01b031685876040516137b3919061509c565b60006040518083038185875af1925050503d80600081146137f0576040519150601f19603f3d011682016040523d82523d6000602084013e6137f5565b606091505b509150915061380687838387613811565b979650505050505050565b60608315613880578251600003613879576001600160a01b0385163b6138795760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610965565b50816116ce565b6116ce83838151156138955781518083602001fd5b8060405162461bcd60e51b815260040161096591906139a1565b6001600160a01b038116811461123a57600080fd5b80356138cf816138af565b919050565b6000602082840312156138e657600080fd5b8135610fa7816138af565b60a08101610d3582846001600160a01b0380825116835267ffffffffffffffff60208301511660208401526bffffffffffffffffffffffff6040830151166040840152806060830151166060840152806080830151166080840152505050565b60005b8381101561396c578181015183820152602001613954565b50506000910152565b6000815180845261398d816020860160208601613951565b601f01601f19169290920160200192915050565b602081526000610fa76020830184613975565b67ffffffffffffffff8116811461123a57600080fd5b80356138cf816139b4565b600060a082840312156139e757600080fd5b50919050565b60008060408385031215613a0057600080fd5b8235613a0b816139b4565b9150602083013567ffffffffffffffff811115613a2757600080fd5b613a33858286016139d5565b9150509250929050565b60008060408385031215613a5057600080fd5b8235613a5b816139b4565b91506020830135613a6b816138af565b809150509250929050565b600060208284031215613a8857600080fd5b8135610fa7816139b4565b8051151582526020810151613aae602084018261ffff169052565b506040810151613ac6604084018263ffffffff169052565b506060810151613ade606084018263ffffffff169052565b506080810151613af6608084018263ffffffff169052565b5060a0810151613b0c60a084018261ffff169052565b5060c0810151613b2460c084018263ffffffff169052565b5060e0810151613b3a60e084018261ffff169052565b506101008181015161ffff9081169184019190915261012080830151909116908301526101408082015163ffffffff90811691840191909152610160808301518216908401526101808083015167ffffffffffffffff908116918501919091526101a080840151909116908401526101c09182015116910152565b600061024082019050613bc9828451613a93565b60208301516001600160a01b03166101e0830152604083015167ffffffffffffffff166102008301526060909201516102209091015290565b60808101610d35828480516001600160a01b03908116835260208083015182169084015260408083015182169084015260609182015116910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613c9057613c90613c3e565b60405290565b60405160c0810167ffffffffffffffff81118282101715613c9057613c90613c3e565b6040516060810167ffffffffffffffff81118282101715613c9057613c90613c3e565b6040516101e0810167ffffffffffffffff81118282101715613c9057613c90613c3e565b604051601f8201601f1916810167ffffffffffffffff81118282101715613d2957613d29613c3e565b604052919050565b600067ffffffffffffffff821115613d4b57613d4b613c3e565b5060051b60200190565b803563ffffffff811681146138cf57600080fd5b803561ffff811681146138cf57600080fd5b801515811461123a57600080fd5b80356138cf81613d7b565b600082601f830112613da557600080fd5b81356020613dba613db583613d31565b613d00565b82815260069290921b84018101918181019086841115613dd957600080fd5b8286015b84811015613e2a5760408189031215613df65760008081fd5b613dfe613c6d565b8135613e09816139b4565b815281850135613e18816138af565b81860152835291830191604001613ddd565b509695505050505050565b60008060408385031215613e4857600080fd5b67ffffffffffffffff83351115613e5e57600080fd5b83601f843585010112613e7057600080fd5b613e80613db58435850135613d31565b8335840180358083526020808401939260059290921b90910101861015613ea657600080fd5b602085358601015b85358601803560051b016020018110156140795767ffffffffffffffff81351115613ed857600080fd5b6040601f19823588358901018903011215613ef257600080fd5b613efa613c6d565b613f0d6020833589358a010101356139b4565b863587018235016020810135825267ffffffffffffffff6040909101351115613f3557600080fd5b86358701823501604081013501603f81018913613f5157600080fd5b613f61613db56020830135613d31565b602082810135808352908201919060e00283016040018b1015613f8357600080fd5b604083015b604060e06020860135028501018110156140605760e0818d031215613fac57600080fd5b613fb4613c6d565b613fbe82356138af565b8135815260c0601f19838f03011215613fd657600080fd5b613fde613c96565b613fea60208401613d55565b8152613ff860408401613d55565b602082015261400960608401613d69565b604082015261401a60808401613d55565b606082015261402b60a08401613d55565b608082015261403d60c0840135613d7b565b60c083013560a0820152602082810191909152908452929092019160e001613f88565b5060208481019190915292865250509283019201613eae565b5092505067ffffffffffffffff6020840135111561409657600080fd5b6140a68460208501358501613d94565b90509250929050565b600060208083850312156140c257600080fd5b823567ffffffffffffffff8111156140d957600080fd5b8301601f810185136140ea57600080fd5b80356140f8613db582613d31565b81815260069190911b8201830190838101908783111561411757600080fd5b928401925b8284101561380657604084890312156141355760008081fd5b61413d613c6d565b8435614148816138af565b815284860135614157816139b4565b818701528252604093909301929084019061411c565b60006080828403121561417f57600080fd5b6040516080810181811067ffffffffffffffff821117156141a2576141a2613c3e565b60405282356141b0816138af565b815260208301356141c0816138af565b602082015260408301356141d3816138af565b604082015260608301356141e6816138af565b60608201529392505050565b6000806000806080858703121561420857600080fd5b8435614213816139b4565b9350602085013567ffffffffffffffff81111561422f57600080fd5b61423b878288016139d5565b935050604085013591506060850135614253816138af565b939692955090935050565b6000602080838503121561427157600080fd5b823567ffffffffffffffff81111561428857600080fd5b8301601f8101851361429957600080fd5b80356142a7613db582613d31565b81815261022091820283018401918482019190888411156142c757600080fd5b938501935b8385101561444c57848903818112156142e55760008081fd5b6142ed613cb9565b86356142f8816139b4565b81526101e0601f19830181131561430f5760008081fd5b614317613cdc565b9250614324898901613d89565b83526040614333818a01613d69565b8a8501526060614344818b01613d55565b828601526080614355818c01613d55565b8287015260a09150614368828c01613d55565b9086015260c06143798b8201613d69565b8287015260e0915061438c828c01613d55565b9086015261010061439e8b8201613d69565b8287015261012091506143b2828c01613d69565b908601526101406143c48b8201613d69565b8287015261016091506143d8828c01613d55565b908601526101806143ea8b8201613d55565b828701526101a091506143fe828c016139ca565b908601526101c06144108b82016139ca565b8287015261441f848c01613d55565b818701525050838a8401526144376102008a016138c4565b908301525084525093840193918501916142cc565b50979650505050505050565b6020808252825182820181905260009190848201906040850190845b818110156144995783516001600160a01b031683529284019291840191600101614474565b50909695505050505050565b6000808335601e198436030181126144bc57600080fd5b83018035915067ffffffffffffffff8211156144d757600080fd5b6020019150368190038213156144ec57600080fd5b9250929050565b6000808335601e1984360301811261450a57600080fd5b83018035915067ffffffffffffffff82111561452557600080fd5b6020019150600681901b36038213156144ec57600080fd5b80517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811681146138cf57600080fd5b6000806040838503121561457c57600080fd5b6145858361453d565b91506140a66020840161453d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082028115828204841417610d3557610d35614593565b80820180821115610d3557610d35614593565b600082614622577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000602080838503121561463a57600080fd5b825167ffffffffffffffff81111561465157600080fd5b8301601f8101851361466257600080fd5b8051614670613db582613d31565b81815260059190911b8201830190838101908783111561468f57600080fd5b928401925b828410156138065783516146a7816138af565b82529284019290840190614694565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156146f757600080fd5b5051919050565b60006020828403121561471057600080fd5b8151610fa7816138af565b60006020828403121561472d57600080fd5b8151610fa7816139b4565b67ffffffffffffffff81811683821601908082111561475957614759614593565b5092915050565b60006040828403121561477257600080fd5b61477a613c6d565b8235614785816138af565b81526020928301359281019290925250919050565b6000602082840312156147ac57600080fd5b8151610fa781613d7b565b602081526000825160a060208401526147d360c0840182613975565b905067ffffffffffffffff602085015116604084015260408401516001600160a01b038082166060860152606086015160808601528060808701511660a086015250508091505092915050565b600082601f83011261483157600080fd5b815167ffffffffffffffff81111561484b5761484b613c3e565b61485e6020601f19601f84011601613d00565b81815284602083860101111561487357600080fd5b6116ce826020830160208701613951565b60006020828403121561489657600080fd5b815167ffffffffffffffff808211156148ae57600080fd5b90830190604082860312156148c257600080fd5b6148ca613c6d565b8251828111156148d957600080fd5b6148e587828601614820565b8252506020830151828111156148fa57600080fd5b61490687828601614820565b60208301525095945050505050565b6020815260008251606060208401526149316080840182613975565b90506020840151601f198085840301604086015261494f8383613975565b925060408601519150808584030160608601525061496d8282613975565b95945050505050565b60008151808452602080850194506020840160005b838110156149bb57815180516001600160a01b03168852830151838801526040909601959082019060010161498b565b509495945050505050565b60008282518085526020808601955060208260051b8401016020860160005b84811015614a1357601f19868403018952614a01838351613975565b988401989250908301906001016149e5565b5090979650505050505050565b60208152614a3b60208201835167ffffffffffffffff169052565b60006020830151614a5760408401826001600160a01b03169052565b5060408301516001600160a01b038116606084015250606083015167ffffffffffffffff8116608084015250608083015160a083015260a0830151614aa060c084018215159052565b5060c083015167ffffffffffffffff811660e08401525060e0830151610100614ad3818501836001600160a01b03169052565b840151610120848101919091528401516101a061014080860182905291925090614b016101c0860184613975565b9250808601519050601f19610160818786030181880152614b228584614976565b945080880151925050610180818786030181880152614b4185846149c6565b970151959092019490945250929392505050565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015614b955780818660040360031b1b83161692505b505092915050565b60008085851115614bad57600080fd5b83861115614bba57600080fd5b5050820193919092039150565b600060208284031215614bd957600080fd5b6040516020810181811067ffffffffffffffff82111715614bfc57614bfc613c3e565b6040529135825250919050565b63ffffffff81811683821601908082111561475957614759614593565b600060208284031215614c3857600080fd5b610fa78261453d565b6101208101614ca282856001600160a01b0380825116835267ffffffffffffffff60208301511660208401526bffffffffffffffffffffffff6040830151166040840152806060830151166060840152806080830151166080840152505050565b82516001600160a01b0390811660a08401526020840151811660c08401526040840151811660e0840152606084015116610100830152610fa7565b6000808335601e19843603018112614cf457600080fd5b830160208101925035905067ffffffffffffffff811115614d1457600080fd5b8036038213156144ec57600080fd5b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b8183526000602080850194508260005b858110156149bb578135614d71816138af565b6001600160a01b03168752818301358388015260409687019690910190600101614d5e565b600067ffffffffffffffff808516835260406020840152614db78485614cdd565b60a06040860152614dcc60e086018284614d23565b915050614ddc6020860186614cdd565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc080878503016060880152614e12848385614d23565b935060408801359250601e19883603018312614e2d57600080fd5b60209288019283019235915084821115614e4657600080fd5b8160061b3603831315614e5857600080fd5b80878503016080880152614e6d848385614d4e565b9450614e7b606089016138c4565b6001600160a01b03811660a08901529350614e996080890189614cdd565b94509250808786030160c08801525050613806838383614d23565b600067ffffffffffffffff808316818103614ed157614ed1614593565b6001019392505050565b602081526000610fa76020830184614976565b602081526000610fa760208301846149c6565b815460ff81161515825261024082019061ffff600882901c8116602085015263ffffffff601883901c81166040860152614f4860608601828560381c1663ffffffff169052565b614f5f60808601828560581c1663ffffffff169052565b614f7460a08601838560781c1661ffff169052565b614f8b60c08601828560881c1663ffffffff169052565b614fa060e08601838560a81c1661ffff169052565b614fb66101008601838560b81c1661ffff169052565b614fcc6101208601838560c81c1661ffff169052565b614fe46101408601828560d81c1663ffffffff169052565b600186015463ffffffff82821616610160870152925067ffffffffffffffff602084901c8116610180870152915061502e6101a08601838560601c1667ffffffffffffffff169052565b6150466101c08601828560a01c1663ffffffff169052565b5060028501546001600160a01b0381166101e0860152915061507a6102008501828460a01c1667ffffffffffffffff169052565b5050600383015461022083015292915050565b6101e08101610d358284613a93565b600082516150ae818460208701613951565b919091019291505056fea164736f6c6343000818000a",
}

var EVM2EVMMultiOnRampABI = EVM2EVMMultiOnRampMetaData.ABI

var EVM2EVMMultiOnRampBin = EVM2EVMMultiOnRampMetaData.Bin

func DeployEVM2EVMMultiOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMMultiOnRampStaticConfig, dynamicConfig EVM2EVMMultiOnRampDynamicConfig, destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs, premiumMultiplierWeiPerEthArgs []EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthArgs, tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs) (common.Address, *types.Transaction, *EVM2EVMMultiOnRamp, error) {
	parsed, err := EVM2EVMMultiOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMMultiOnRampBin), backend, staticConfig, dynamicConfig, destChainConfigArgs, premiumMultiplierWeiPerEthArgs, tokenTransferFeeConfigArgs)
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMMultiOnRamp.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) GetAdmin() (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetAdmin(&_EVM2EVMMultiOnRamp.CallOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampCallerSession) GetAdmin() (common.Address, error) {
	return _EVM2EVMMultiOnRamp.Contract.GetAdmin(&_EVM2EVMMultiOnRamp.CallOpts)
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.TransferOwnership(&_EVM2EVMMultiOnRamp.TransactOpts, to)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.TransferOwnership(&_EVM2EVMMultiOnRamp.TransactOpts, to)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactor) WithdrawFeeTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.contract.Transact(opts, "withdrawFeeTokens")
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampSession) WithdrawFeeTokens() (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.WithdrawFeeTokens(&_EVM2EVMMultiOnRamp.TransactOpts)
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampTransactorSession) WithdrawFeeTokens() (*types.Transaction, error) {
	return _EVM2EVMMultiOnRamp.Contract.WithdrawFeeTokens(&_EVM2EVMMultiOnRamp.TransactOpts)
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

type EVM2EVMMultiOnRampFeePaidIterator struct {
	Event *EVM2EVMMultiOnRampFeePaid

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampFeePaidIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampFeePaid)
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
		it.Event = new(EVM2EVMMultiOnRampFeePaid)
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

func (it *EVM2EVMMultiOnRampFeePaidIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampFeePaid struct {
	FeeToken      common.Address
	FeeValueJuels *big.Int
	Raw           types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterFeePaid(opts *bind.FilterOpts, feeToken []common.Address) (*EVM2EVMMultiOnRampFeePaidIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "FeePaid", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampFeePaidIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "FeePaid", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchFeePaid(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampFeePaid, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "FeePaid", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampFeePaid)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "FeePaid", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseFeePaid(log types.Log) (*EVM2EVMMultiOnRampFeePaid, error) {
	event := new(EVM2EVMMultiOnRampFeePaid)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "FeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOnRampFeeTokenWithdrawnIterator struct {
	Event *EVM2EVMMultiOnRampFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOnRampFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOnRampFeeTokenWithdrawn)
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
		it.Event = new(EVM2EVMMultiOnRampFeeTokenWithdrawn)
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

func (it *EVM2EVMMultiOnRampFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOnRampFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOnRampFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*EVM2EVMMultiOnRampFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOnRampFeeTokenWithdrawnIterator{contract: _EVM2EVMMultiOnRamp.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _EVM2EVMMultiOnRamp.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOnRampFeeTokenWithdrawn)
				if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRampFilterer) ParseFeeTokenWithdrawn(log types.Log) (*EVM2EVMMultiOnRampFeeTokenWithdrawn, error) {
	event := new(EVM2EVMMultiOnRampFeeTokenWithdrawn)
	if err := _EVM2EVMMultiOnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMMultiOnRamp.abi.Events["AdminSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseAdminSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["CCIPSendRequested"].ID:
		return _EVM2EVMMultiOnRamp.ParseCCIPSendRequested(log)
	case _EVM2EVMMultiOnRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMMultiOnRamp.ParseConfigSet(log)
	case _EVM2EVMMultiOnRamp.abi.Events["DestChainAdded"].ID:
		return _EVM2EVMMultiOnRamp.ParseDestChainAdded(log)
	case _EVM2EVMMultiOnRamp.abi.Events["DestChainDynamicConfigUpdated"].ID:
		return _EVM2EVMMultiOnRamp.ParseDestChainDynamicConfigUpdated(log)
	case _EVM2EVMMultiOnRamp.abi.Events["FeePaid"].ID:
		return _EVM2EVMMultiOnRamp.ParseFeePaid(log)
	case _EVM2EVMMultiOnRamp.abi.Events["FeeTokenWithdrawn"].ID:
		return _EVM2EVMMultiOnRamp.ParseFeeTokenWithdrawn(log)
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

func (EVM2EVMMultiOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x823ad577168c60d73b7419edfe54229b4737c81cc2eb4db07b2b6604c040429f")
}

func (EVM2EVMMultiOnRampDestChainAdded) Topic() common.Hash {
	return common.HexToHash("0x7a70081ee29c1fc27898089ba2a5fc35ac0106b043c82ccecd24c6fd48f6ca86")
}

func (EVM2EVMMultiOnRampDestChainDynamicConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x944eb884a589931130671ee4a7379fbe5fe65ed605a048ba99c454582f2460b0")
}

func (EVM2EVMMultiOnRampFeePaid) Topic() common.Hash {
	return common.HexToHash("0x075a2720282fdf622141dae0b048ef90a21a7e57c134c76912d19d006b3b3f6f")
}

func (EVM2EVMMultiOnRampFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
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
	return common.HexToHash("0x94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b5")
}

func (_EVM2EVMMultiOnRamp *EVM2EVMMultiOnRamp) Address() common.Address {
	return _EVM2EVMMultiOnRamp.address
}

type EVM2EVMMultiOnRampInterface interface {
	GetAdmin(opts *bind.CallOpts) (common.Address, error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (EVM2EVMMultiOnRampDestChainConfig, error)

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMMultiOnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetPremiumMultiplierWeiPerEth(opts *bind.CallOpts, token common.Address) (uint64, error)

	GetSenderNonce(opts *bind.CallOpts, destChainSelector uint64, sender common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMMultiOnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (EVM2EVMMultiOnRampTokenTransferFeeConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []EVM2EVMMultiOnRampDestChainConfigArgs) (*types.Transaction, error)

	ApplyPremiumMultiplierWeiPerEthUpdates(opts *bind.TransactOpts, premiumMultiplierWeiPerEthArgs []EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []EVM2EVMMultiOnRampTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMMultiOnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampAdminSetIterator, error)

	WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampAdminSet) (event.Subscription, error)

	ParseAdminSet(log types.Log) (*EVM2EVMMultiOnRampAdminSet, error)

	FilterCCIPSendRequested(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampCCIPSendRequestedIterator, error)

	WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampCCIPSendRequested, destChainSelector []uint64) (event.Subscription, error)

	ParseCCIPSendRequested(log types.Log) (*EVM2EVMMultiOnRampCCIPSendRequested, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMMultiOnRampConfigSet, error)

	FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampDestChainAddedIterator, error)

	WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampDestChainAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainAdded(log types.Log) (*EVM2EVMMultiOnRampDestChainAdded, error)

	FilterDestChainDynamicConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*EVM2EVMMultiOnRampDestChainDynamicConfigUpdatedIterator, error)

	WatchDestChainDynamicConfigUpdated(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampDestChainDynamicConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainDynamicConfigUpdated(log types.Log) (*EVM2EVMMultiOnRampDestChainDynamicConfigUpdated, error)

	FilterFeePaid(opts *bind.FilterOpts, feeToken []common.Address) (*EVM2EVMMultiOnRampFeePaidIterator, error)

	WatchFeePaid(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampFeePaid, feeToken []common.Address) (event.Subscription, error)

	ParseFeePaid(log types.Log) (*EVM2EVMMultiOnRampFeePaid, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*EVM2EVMMultiOnRampFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*EVM2EVMMultiOnRampFeeTokenWithdrawn, error)

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

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
