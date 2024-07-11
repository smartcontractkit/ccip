// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package price_registry

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

type AuthorizedCallersAuthorizedCallerArgs struct {
	AddedCallers   []common.Address
	RemovedCallers []common.Address
}

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

type IPriceRegistryTokenPriceFeedConfig struct {
	DataFeedAddress common.Address
	TokenDecimals   uint8
}

type InternalEVM2AnyRampMessage struct {
	Header         InternalRampMessageHeader
	Sender         common.Address
	Data           []byte
	Receiver       []byte
	ExtraArgs      []byte
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	TokenAmounts   []InternalRampTokenAmount
}

type InternalGasPriceUpdate struct {
	DestChainSelector uint64
	UsdPerUnitGas     *big.Int
}

type InternalPriceUpdates struct {
	TokenPriceUpdates []InternalTokenPriceUpdate
	GasPriceUpdates   []InternalGasPriceUpdate
}

type InternalRampMessageHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

type InternalRampTokenAmount struct {
	SourcePoolAddress []byte
	DestTokenAddress  []byte
	ExtraData         []byte
	Amount            *big.Int
}

type InternalTimestampedPackedUint224 struct {
	Value     *big.Int
	Timestamp uint32
}

type InternalTokenPriceUpdate struct {
	SourceToken common.Address
	UsdPerToken *big.Int
}

type PriceRegistryDestChainDynamicConfig struct {
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
	EnforceOutOfOrder                 bool
	ChainFamilySelector               [4]byte
}

type PriceRegistryDestChainDynamicConfigArgs struct {
	DestChainSelector uint64
	DynamicConfig     PriceRegistryDestChainDynamicConfig
}

type PriceRegistryPremiumMultiplierWeiPerEthArgs struct {
	Token                      common.Address
	PremiumMultiplierWeiPerEth uint64
}

type PriceRegistryStaticConfig struct {
	MaxFeeJuelsPerMsg  *big.Int
	LinkToken          common.Address
	StalenessThreshold uint32
}

type PriceRegistryTokenPriceFeedUpdate struct {
	SourceToken common.Address
	FeedConfig  IPriceRegistryTokenPriceFeedConfig
}

type PriceRegistryTokenTransferFeeConfig struct {
	MinFeeUSDCents    uint32
	MaxFeeUSDCents    uint32
	DeciBps           uint16
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	IsEnabled         bool
}

type PriceRegistryTokenTransferFeeConfigArgs struct {
	DestChainSelector       uint64
	TokenTransferFeeConfigs []PriceRegistryTokenTransferFeeConfigSingleTokenArgs
}

type PriceRegistryTokenTransferFeeConfigRemoveArgs struct {
	DestChainSelector uint64
	Token             common.Address
}

type PriceRegistryTokenTransferFeeConfigSingleTokenArgs struct {
	Token                  common.Address
	TokenTransferFeeConfig PriceRegistryTokenTransferFeeConfig
}

var PriceRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessThreshold\",\"type\":\"uint32\"}],\"internalType\":\"structPriceRegistry.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"priceUpdaters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIPriceRegistry.TokenPriceFeedConfig\",\"name\":\"feedConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.TokenPriceFeedUpdate[]\",\"name\":\"tokenPriceFeeds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structPriceRegistry.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"internalType\":\"structPriceRegistry.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.DestChainDynamicConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chain\",\"type\":\"uint64\"}],\"name\":\"ChainNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataFeedValueOutOfUint224Range\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"DestinationChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtraArgOutOfOrderExecutionMustBeTrue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"}],\"name\":\"InvalidDestBytesOverhead\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidDestChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStaticConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgFeeJuels\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint256\"}],\"name\":\"MessageFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NotAFeeToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SourceTokenDataTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePassed\",\"type\":\"uint256\"}],\"name\":\"StaleGasPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"UnauthorizedCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedCallerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedCallerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"indexed\":false,\"internalType\":\"structPriceRegistry.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"indexed\":false,\"internalType\":\"structPriceRegistry.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainDynamicConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"FeeTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"FeeTokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"name\":\"PremiumMultiplierWeiPerEthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIPriceRegistry.TokenPriceFeedConfig\",\"name\":\"priceFeedConfig\",\"type\":\"tuple\"}],\"name\":\"PriceFeedPerTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceUpdater\",\"type\":\"address\"}],\"name\":\"PriceUpdaterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceUpdater\",\"type\":\"address\"}],\"name\":\"PriceUpdaterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenTransferFeeConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structPriceRegistry.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"name\":\"TokenTransferFeeConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UsdPerTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChain\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UsdPerUnitGasUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"addedCallers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removedCallers\",\"type\":\"address[]\"}],\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\"}],\"name\":\"applyAuthorizedCallerUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"internalType\":\"structPriceRegistry.DestChainDynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.DestChainDynamicConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyDestChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"feeTokensToAdd\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"feeTokensToRemove\",\"type\":\"address[]\"}],\"name\":\"applyFeeTokensUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structPriceRegistry.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyPremiumMultiplierWeiPerEthUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigRemoveArgs[]\",\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\"}],\"name\":\"applyTokenTransferFeeConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"}],\"name\":\"convertTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAuthorizedCallers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestChainDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"internalType\":\"structPriceRegistry.DestChainDynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestinationChainGasPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.TimestampedPackedUint224\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPremiumMultiplierWeiPerEth\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessThreshold\",\"type\":\"uint32\"}],\"internalType\":\"structPriceRegistry.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getTokenAndGasPrices\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"tokenPrice\",\"type\":\"uint224\"},{\"internalType\":\"uint224\",\"name\":\"gasPriceValue\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.TimestampedPackedUint224\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPriceFeedConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIPriceRegistry.TokenPriceFeedConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getTokenPrices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.TimestampedPackedUint224[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenTransferFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getValidatedFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.RampTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.EVM2AnyRampMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\"}],\"name\":\"getValidatedRampMessageParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"msgFeeJuels\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOutOfOrderExecution\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"convertedExtraArgs\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getValidatedTokenPrice\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint224\",\"name\":\"usdPerToken\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint224\",\"name\":\"usdPerUnitGas\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.GasPriceUpdate[]\",\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.PriceUpdates\",\"name\":\"priceUpdates\",\"type\":\"tuple\"}],\"name\":\"updatePrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIPriceRegistry.TokenPriceFeedConfig\",\"name\":\"feedConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.TokenPriceFeedUpdate[]\",\"name\":\"tokenPriceFeedUpdates\",\"type\":\"tuple[]\"}],\"name\":\"updateTokenPriceFeeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200708b3803806200708b8339810160408190526200003491620017f3565b8533806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000207565b5050604080518082018252838152815160008152602080820190935291810191909152620000ee9150620002b2565b5060208701516001600160a01b0316158062000112575086516001600160601b0316155b80620001265750604087015163ffffffff16155b15620001455760405163d794ef9560e01b815260040160405180910390fd5b6020878101516001600160a01b031660a05287516001600160601b031660805260408089015163ffffffff1660c05280516000815291820190526200018c90869062000401565b620001978462000549565b620001a2816200061a565b620001ad8262000a82565b60408051600080825260208201909252620001fa91859190620001f3565b6040805180820190915260008082526020820152815260200190600190039081620001cb5790505b5062000b4e565b5050505050505062001ab7565b336001600160a01b03821603620002615760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b602081015160005b815181101562000342576000828281518110620002db57620002db62001912565b60209081029190910101519050620002f560028262000e81565b1562000338576040516001600160a01b03821681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda775809060200160405180910390a15b50600101620002ba565b50815160005b8151811015620003fb57600082828151811062000369576200036962001912565b6020026020010151905060006001600160a01b0316816001600160a01b031603620003a7576040516342bcdf7f60e11b815260040160405180910390fd5b620003b460028262000ea1565b506040516001600160a01b03821681527feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef9060200160405180910390a15060010162000348565b50505050565b60005b8251811015620004a2576200044083828151811062000427576200042762001912565b6020026020010151600c62000ea160201b90919060201c565b1562000499578281815181106200045b576200045b62001912565b60200260200101516001600160a01b03167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba2360405160405180910390a25b60010162000404565b5060005b81518110156200054457620004e2828281518110620004c957620004c962001912565b6020026020010151600c62000e8160201b90919060201c565b156200053b57818181518110620004fd57620004fd62001912565b60200260200101516001600160a01b03167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f9160405160405180910390a25b600101620004a6565b505050565b60005b8151811015620006165760008282815181106200056d576200056d62001912565b6020908102919091018101518051818301516001600160a01b0380831660008181526006875260409081902084518154868a018051929096166001600160a81b03199091168117600160a01b60ff9384160217909255825191825293519093169683019690965293955091939092917f08a5f7f5bb38a81d8e43aca13ecd76431dbf8816ae4699affff7b00b2fc1c464910160405180910390a25050508060010190506200054c565b5050565b60005b8151811015620006165760008282815181106200063e576200063e62001912565b6020026020010151905060008383815181106200065f576200065f62001912565b6020026020010151600001519050600082602001519050816001600160401b0316600014806200069b57506101808101516001600160401b0316155b80620006bd57506102008101516001600160e01b031916630a04b54b60e21b14155b80620006d45750602081610160015163ffffffff16105b15620006ff5760405163c35aa79d60e01b81526001600160401b038316600482015260240162000083565b6001600160401b038216600090815260086020526040812060010154600160c81b900460e01b6001600160e01b03191690036200077f57816001600160401b03167fd883f18d08e8ea5c2af22d13fb6cab86cc057d795d2735f83a8ff137d39d1b8f8260405162000771919062001928565b60405180910390a2620007c3565b816001600160401b03167fd42d3a670a4f1ab5d8703efa22bb17041446365cfcfa33980ced685a080cf7cb82604051620007ba919062001928565b60405180910390a25b8060086000846001600160401b03166001600160401b0316815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a8154816001600160401b0302191690836001600160401b031602179055506101a082015181600101600c6101000a8154816001600160401b0302191690836001600160401b031602179055506101c08201518160010160146101000a81548163ffffffff021916908363ffffffff1602179055506101e08201518160010160186101000a81548160ff0219169083151502179055506102008201518160010160196101000a81548163ffffffff021916908360e01c02179055509050505050508060010190506200061d565b60005b81518110156200061657600082828151811062000aa65762000aa662001912565b6020026020010151600001519050600083838151811062000acb5762000acb62001912565b6020908102919091018101518101516001600160a01b03841660008181526007845260409081902080546001600160401b0319166001600160401b0385169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a2505060010162000a85565b60005b825181101562000dbb57600083828151811062000b725762000b7262001912565b6020026020010151905060008160000151905060005b82602001515181101562000dac5760008360200151828151811062000bb15762000bb162001912565b602002602001015160200151905060008460200151838151811062000bda5762000bda62001912565b60200260200101516000015190506020826080015163ffffffff16101562000c335760808201516040516312766e0160e11b81526001600160a01b038316600482015263ffffffff909116602482015260440162000083565b6001600160401b03841660008181526009602090815260408083206001600160a01b0386168085529083529281902086518154938801518389015160608a015160808b015160a08c01511515600160901b0260ff60901b1963ffffffff928316600160701b021664ffffffffff60701b199383166a01000000000000000000000263ffffffff60501b1961ffff90961668010000000000000000029590951665ffffffffffff60401b19968416640100000000026001600160401b0319909b16939097169290921798909817939093169390931717919091161792909217909155519091907f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b59062000d99908690600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60405180910390a3505060010162000b88565b50505080600101905062000b51565b5060005b81518110156200054457600082828151811062000de05762000de062001912565b6020026020010151600001519050600083838151811062000e055762000e0562001912565b6020908102919091018101518101516001600160401b03841660008181526009845260408082206001600160a01b038516808452955280822080546001600160981b03191690555192945090917f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b9190a3505060010162000dbf565b600062000e98836001600160a01b03841662000eb8565b90505b92915050565b600062000e98836001600160a01b03841662000fbc565b6000818152600183016020526040812054801562000fb157600062000edf60018362001a7f565b855490915060009062000ef59060019062001a7f565b905081811462000f6157600086600001828154811062000f195762000f1962001912565b906000526020600020015490508087600001848154811062000f3f5762000f3f62001912565b6000918252602080832090910192909255918252600188019052604090208390555b855486908062000f755762000f7562001aa1565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000e9b565b600091505062000e9b565b6000818152600183016020526040812054620010055750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000e9b565b50600062000e9b565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156200104957620010496200100e565b60405290565b60405160c081016001600160401b03811182821017156200104957620010496200100e565b60405161022081016001600160401b03811182821017156200104957620010496200100e565b604051601f8201601f191681016001600160401b0381118282101715620010c557620010c56200100e565b604052919050565b80516001600160a01b0381168114620010e557600080fd5b919050565b805163ffffffff81168114620010e557600080fd5b6000606082840312156200111257600080fd5b604051606081016001600160401b03811182821017156200113757620011376200100e565b604052825190915081906001600160601b03811681146200115757600080fd5b81526200116760208401620010cd565b60208201526200117a60408401620010ea565b60408201525092915050565b60006001600160401b03821115620011a257620011a26200100e565b5060051b60200190565b600082601f830112620011be57600080fd5b81516020620011d7620011d18362001186565b6200109a565b8083825260208201915060208460051b870101935086841115620011fa57600080fd5b602086015b8481101562001221576200121381620010cd565b8352918301918301620011ff565b509695505050505050565b600082601f8301126200123e57600080fd5b8151602062001251620011d18362001186565b828152606092830285018201928282019190878511156200127157600080fd5b8387015b85811015620013045780890382811215620012905760008081fd5b6200129a62001024565b620012a583620010cd565b8152604080601f1984011215620012bc5760008081fd5b620012c662001024565b9250620012d5888501620010cd565b835283015160ff81168114620012eb5760008081fd5b8288015280870191909152845292840192810162001275565b5090979650505050505050565b80516001600160401b0381168114620010e557600080fd5b805161ffff81168114620010e557600080fd5b80518015158114620010e557600080fd5b600082601f8301126200135f57600080fd5b8151602062001372620011d18362001186565b82815260059290921b840181019181810190868411156200139257600080fd5b8286015b84811015620012215780516001600160401b0380821115620013b757600080fd5b908801906040601f19838c038101821315620013d257600080fd5b620013dc62001024565b620013e989860162001311565b81528285015184811115620013fd57600080fd5b8086019550508c603f8601126200141357600080fd5b88850151935062001428620011d18562001186565b84815260e09094028501830193898101908e8611156200144757600080fd5b958401955b858710156200152057868f0360e08112156200146757600080fd5b6200147162001024565b6200147c89620010cd565b815260c086830112156200148f57600080fd5b620014996200104f565b9150620014a88d8a01620010ea565b8252620014b7878a01620010ea565b8d830152620014c960608a0162001329565b87830152620014db60808a01620010ea565b6060830152620014ee60a08a01620010ea565b60808301526200150160c08a016200133c565b60a0830152808d0191909152825260e09690960195908a01906200144c565b828b01525087525050509284019250830162001396565b600082601f8301126200154957600080fd5b815160206200155c620011d18362001186565b82815260069290921b840181019181810190868411156200157c57600080fd5b8286015b848110156200122157604081890312156200159b5760008081fd5b620015a562001024565b620015b082620010cd565b8152620015bf85830162001311565b8186015283529183019160400162001580565b80516001600160e01b031981168114620010e557600080fd5b600082601f830112620015fd57600080fd5b8151602062001610620011d18362001186565b82815261024092830285018201928282019190878511156200163157600080fd5b8387015b85811015620013045780890382811215620016505760008081fd5b6200165a62001024565b620016658362001311565b815261022080601f19840112156200167d5760008081fd5b6200168762001074565b9250620016968885016200133c565b83526040620016a781860162001329565b898501526060620016ba818701620010ea565b8286015260809150620016cf828701620010ea565b9085015260a0620016e2868201620010ea565b8286015260c09150620016f782870162001329565b9085015260e06200170a868201620010ea565b8286015261010091506200172082870162001329565b908501526101206200173486820162001329565b8286015261014091506200174a82870162001329565b908501526101606200175e868201620010ea565b82860152610180915062001774828701620010ea565b908501526101a06200178886820162001311565b828601526101c091506200179e82870162001311565b908501526101e0620017b2868201620010ea565b828601526102009150620017c88287016200133c565b90850152620017d9858301620015d2565b908401525080870191909152845292840192810162001635565b6000806000806000806000610120888a0312156200181057600080fd5b6200181c8989620010ff565b60608901519097506001600160401b03808211156200183a57600080fd5b620018488b838c01620011ac565b975060808a01519150808211156200185f57600080fd5b6200186d8b838c01620011ac565b965060a08a01519150808211156200188457600080fd5b620018928b838c016200122c565b955060c08a0151915080821115620018a957600080fd5b620018b78b838c016200134d565b945060e08a0151915080821115620018ce57600080fd5b620018dc8b838c0162001537565b93506101008a0151915080821115620018f457600080fd5b50620019038a828b01620015eb565b91505092959891949750929550565b634e487b7160e01b600052603260045260246000fd5b8151151581526102208101602083015162001949602084018261ffff169052565b50604083015162001962604084018263ffffffff169052565b5060608301516200197b606084018263ffffffff169052565b50608083015162001994608084018263ffffffff169052565b5060a0830151620019ab60a084018261ffff169052565b5060c0830151620019c460c084018263ffffffff169052565b5060e0830151620019db60e084018261ffff169052565b506101008381015161ffff9081169184019190915261012080850151909116908301526101408084015163ffffffff9081169184019190915261016080850151821690840152610180808501516001600160401b03908116918501919091526101a080860151909116908401526101c080850151909116908301526101e080840151151590830152610200928301516001600160e01b031916929091019190915290565b8181038181111562000e9b57634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60805160a05160c05161557a62001b11600039600081816102ce0152818161224401526122ad0152600081816102920152818161126e01526112e501526000818161025e0152818161130e015261137e015261557a6000f3fe608060405234801561001057600080fd5b50600436106101ad5760003560e01c806382b49eb0116100ee578063ca64321311610097578063d8694ccd11610071578063d8694ccd146107b9578063eedeb982146107cc578063f2fde38b14610a70578063ffdb4b3714610a8357600080fd5b8063ca6432131461078b578063cdc73d511461079e578063d02641a0146107a657600080fd5b8063997fb26b116100c8578063997fb26b146106a9578063a69c64c0146106cb578063bf78e03f146106de57600080fd5b806382b49eb0146104fe5780638da5cb5b1461066e57806391a2749a1461069657600080fd5b8063407e10861161015b578063514e8cff11610135578063514e8cff1461042d578063770e2dc4146104d057806379ba5097146104e35780637afac322146104eb57600080fd5b8063407e1086146103ba57806345ac924d146103cd5780634ab35b0b146103ed57600080fd5b8063181f5a771161018c578063181f5a77146103475780632451a627146103905780633937306f146103a557600080fd5b806241e5be146101b2578063061877e3146101d857806306285c6914610231575b600080fd5b6101c56101c0366004613f9e565b610acb565b6040519081526020015b60405180910390f35b6102186101e6366004613fda565b73ffffffffffffffffffffffffffffffffffffffff1660009081526007602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff90911681526020016101cf565b6102fb604080516060810182526000808252602082018190529181019190915260405180606001604052807f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000063ffffffff16815250905090565b6040805182516bffffffffffffffffffffffff16815260208084015173ffffffffffffffffffffffffffffffffffffffff16908201529181015163ffffffff16908201526060016101cf565b6103836040518060400160405280601781526020017f5072696365526567697374727920312e362e302d64657600000000000000000081525081565b6040516101cf9190614059565b610398610b39565b6040516101cf919061406c565b6103b86103b33660046140c6565b610b4a565b005b6103b86103c8366004614222565b610e1e565b6103e06103db36600461433b565b610e32565b6040516101cf91906143b0565b6104006103fb366004613fda565b610efd565b6040517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90911681526020016101cf565b6104c361043b366004614443565b60408051808201909152600080825260208201525067ffffffffffffffff166000908152600460209081526040918290208251808401909352547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811683527c0100000000000000000000000000000000000000000000000000000000900463ffffffff169082015290565b6040516101cf919061445e565b6103b86104de366004614570565b610f08565b6103b8610f1e565b6103b86104f936600461488a565b611020565b61060e61050c3660046148ee565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a08101919091525067ffffffffffffffff91909116600090815260096020908152604080832073ffffffffffffffffffffffffffffffffffffffff94909416835292815290829020825160c081018452905463ffffffff8082168352640100000000820481169383019390935268010000000000000000810461ffff16938201939093526a01000000000000000000008304821660608201526e01000000000000000000000000000083049091166080820152720100000000000000000000000000000000000090910460ff16151560a082015290565b6040516101cf9190600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101cf565b6103b86106a4366004614918565b611032565b6106bc6106b73660046149a9565b611043565b6040516101cf93929190614a4c565b6103b86106d9366004614a76565b61142b565b6107576106ec366004613fda565b6040805180820182526000808252602091820181905273ffffffffffffffffffffffffffffffffffffffff93841681526006825282902082518084019093525492831682527401000000000000000000000000000000000000000090920460ff169181019190915290565b60408051825173ffffffffffffffffffffffffffffffffffffffff16815260209283015160ff1692810192909252016101cf565b6103b8610799366004614b6b565b61143c565b61039861144d565b6104c36107b4366004613fda565b611459565b6101c56107c7366004614d8b565b611555565b610a636107da366004614443565b6040805161022081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c081018290526101e081018290526102008101919091525067ffffffffffffffff908116600090815260086020908152604091829020825161022081018452815460ff8082161515835261ffff61010080840482169685019690965263ffffffff630100000084048116978501979097526701000000000000008304871660608501526b0100000000000000000000008304871660808501526f010000000000000000000000000000008304811660a0850152710100000000000000000000000000000000008304871660c085015275010000000000000000000000000000000000000000008304811660e0808601919091527701000000000000000000000000000000000000000000000084048216968501969096527901000000000000000000000000000000000000000000000000008084049091166101208501527b010000000000000000000000000000000000000000000000000000009092048616610140840152600190930154808616610160840152640100000000810487166101808401526c0100000000000000000000000081049096166101a08301527401000000000000000000000000000000000000000086049094166101c08201527801000000000000000000000000000000000000000000000000850490911615156101e08201527fffffffff0000000000000000000000000000000000000000000000000000000092909304901b1661020082015290565b6040516101cf9190614de0565b6103b8610a7e366004613fda565b61215f565b610a96610a91366004614f4c565b612170565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9384168152929091166020830152016101cf565b6000610ad6826122fb565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16610afd856122fb565b610b25907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1685614fa5565b610b2f9190614fbc565b90505b9392505050565b6060610b456002612395565b905090565b60005473ffffffffffffffffffffffffffffffffffffffff163314610b7157610b716123a2565b6000610b7d8280614ff7565b9050905060005b81811015610cc7576000610b988480614ff7565b83818110610ba857610ba861505f565b905060400201803603810190610bbe91906150ba565b604080518082018252602080840180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff908116845263ffffffff42818116858701908152885173ffffffffffffffffffffffffffffffffffffffff9081166000908152600590975295889020965190519092167c010000000000000000000000000000000000000000000000000000000002919092161790935584519051935194955016927f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a92610cb69290917bffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a250600101610b84565b506000610cd76020840184614ff7565b9050905060005b81811015610e18576000610cf56020860186614ff7565b83818110610d0557610d0561505f565b905060400201803603810190610d1b91906150f7565b604080518082018252602080840180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff908116845263ffffffff42818116858701908152885167ffffffffffffffff9081166000908152600490975295889020965190519092167c010000000000000000000000000000000000000000000000000000000002919092161790935584519051935194955016927fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e92610e079290917bffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a250600101610cde565b50505050565b610e266123e7565b610e2f81612468565b50565b60608160008167ffffffffffffffff811115610e5057610e50614101565b604051908082528060200260200182016040528015610e9557816020015b6040805180820190915260008082526020820152815260200190600190039081610e6e5790505b50905060005b82811015610ef257610ecd868683818110610eb857610eb861505f565b90506020020160208101906107b49190613fda565b828281518110610edf57610edf61505f565b6020908102919091010152600101610e9b565b509150505b92915050565b6000610ef7826122fb565b610f106123e7565b610f1a8282612566565b5050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610fa4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6110286123e7565b610f1a8282612972565b61103a6123e7565b610e2f81612ab9565b60008060608161105887830160408901614443565b67ffffffffffffffff811660009081526008602052604081209192505b6110836101608a018a61511a565b90508110156112565760008888838181106110a0576110a061505f565b6110b69260206040909202019081019150613fda565b905060006110c86101608c018c61511a565b848181106110d8576110d861505f565b90506020028101906110ea9190615182565b6110f89060408101906151c0565b91505060208111801561115c575067ffffffffffffffff8516600090815260096020908152604080832073ffffffffffffffffffffffffffffffffffffffff861684529091529020546e010000000000000000000000000000900463ffffffff1681115b156111ab576040517f36f536ca00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610f9b565b600184015461124c90790100000000000000000000000000000000000000000000000000900460e01b6111e26101608e018e61511a565b868181106111f2576111f261505f565b90506020028101906112049190615182565b6112129060208101906151c0565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612c4592505050565b5050600101611075565b5073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166112a06101408a016101208b01613fda565b73ffffffffffffffffffffffffffffffffffffffff16036112c857876101400135945061130c565b6113096112dd6101408a016101208b01613fda565b8961014001357f0000000000000000000000000000000000000000000000000000000000000000610acb565b94505b7f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff168511156113ab576040517f6a92a483000000000000000000000000000000000000000000000000000000008152600481018690526bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166024820152604401610f9b565b60006113da6113be6101008b018b6151c0565b6001850154640100000000900467ffffffffffffffff16612c97565b90508060200151945085858260405160200161140a91908151815260209182015115159181019190915260400190565b60405160208183030381529060405295509550955050505093509350939050565b6114336123e7565b610e2f81612e40565b6114446123e7565b610e2f81612f2a565b6060610b45600c612395565b604080518082019091526000808252602082015273ffffffffffffffffffffffffffffffffffffffff8281166000908152600660209081526040918290208251808401909352549283168083527401000000000000000000000000000000000000000090930460ff16908201529061154c57505073ffffffffffffffffffffffffffffffffffffffff166000908152600560209081526040918290208251808401909352547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811683527c0100000000000000000000000000000000000000000000000000000000900463ffffffff169082015290565b610b3281613403565b67ffffffffffffffff82166000908152600860205260408120805460ff166115b5576040517f99ac52f200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610f9b565b60006115c46040850185614ff7565b6040805161022081018252855460ff8082161515835261ffff610100808404821660208087019190915263ffffffff630100000086048116978701979097526701000000000000008504871660608701526b0100000000000000000000008504871660808701526f010000000000000000000000000000008504831660a0870152710100000000000000000000000000000000008504871660c087015275010000000000000000000000000000000000000000008504831660e0808801919091527701000000000000000000000000000000000000000000000086048416928701929092527901000000000000000000000000000000000000000000000000008086049093166101208701527b01000000000000000000000000000000000000000000000000000000909404861661014086015260018a015480871661016087015267ffffffffffffffff640100000000820481166101808801526c010000000000000000000000008204166101a08701527401000000000000000000000000000000000000000081049096166101c08601527801000000000000000000000000000000000000000000000000860490921615156101e08501527fffffffff000000000000000000000000000000000000000000000000000000009404901b9290921661020082015291935061180892506117c1908701876151c0565b9050836117ce88806151c0565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061364692505050565b600060078161181d6080880160608901613fda565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160009081205467ffffffffffffffff1691508190036118b4576118696080860160608701613fda565b6040517fa7499d2000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610f9b565b6000806118d06118ca6080890160608a01613fda565b89612170565b9092509050600080808615611b05576040805161022081018252895460ff8082161515835261010080830461ffff90811660208601526301000000840463ffffffff90811696860196909652670100000000000000840486166060808701919091526b010000000000000000000000850487166080808801919091526f010000000000000000000000000000008604831660a0880152710100000000000000000000000000000000008604881660c088015275010000000000000000000000000000000000000000008604831660e0808901919091527701000000000000000000000000000000000000000000000087048416948801949094527901000000000000000000000000000000000000000000000000008087049093166101208801527b01000000000000000000000000000000000000000000000000000000909504871661014087015260018f0154808816610160880152640100000000810467ffffffffffffffff9081166101808901526c010000000000000000000000008204166101a08801527401000000000000000000000000000000000000000081049097166101c08701527801000000000000000000000000000000000000000000000000870490931615156101e086015290940490931b7fffffffff0000000000000000000000000000000000000000000000000000000016610200830152611af9928e91611ae391908f01908f01613fda565b888e8060400190611af49190614ff7565b6136f0565b91945092509050611b3c565b6001880154611b399074010000000000000000000000000000000000000000900463ffffffff16662386f26fc10000614fa5565b92505b875460009077010000000000000000000000000000000000000000000000900461ffff1615611d805760408051610220810182528a5460ff8082161515835261ffff610100808404821660208087019190915263ffffffff630100000086048116978701979097526701000000000000008504871660608701526b0100000000000000000000008504871660808701526f010000000000000000000000000000008504831660a0870152710100000000000000000000000000000000008504871660c087015275010000000000000000000000000000000000000000008504831660e0808801919091527701000000000000000000000000000000000000000000000086048416928701929092527901000000000000000000000000000000000000000000000000008086049093166101208701527b01000000000000000000000000000000000000000000000000000000909404861661014086015260018f015480871661016087015267ffffffffffffffff640100000000820481166101808801526c010000000000000000000000008204166101a08701527401000000000000000000000000000000000000000081049096166101c08601527801000000000000000000000000000000000000000000000000860490921615156101e08501527fffffffff000000000000000000000000000000000000000000000000000000009404901b92909216610200820152611d7d916dffffffffffffffffffffffffffff607089901c1690611d74908f018f6151c0565b90508b866139ce565b90505b885460009063ffffffff8516906f01000000000000000000000000000000900461ffff16611db160208f018f6151c0565b611dbc929150614fa5565b8b54611ddd91906b010000000000000000000000900463ffffffff16615225565b611de79190615225565b90506120a4611df960808e018e6151c0565b8c604051806102200160405290816000820160009054906101000a900460ff161515151581526020016000820160019054906101000a900461ffff1661ffff1661ffff1681526020016000820160039054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016000820160079054906101000a900463ffffffff1663ffffffff1663ffffffff16815260200160008201600b9054906101000a900463ffffffff1663ffffffff1663ffffffff16815260200160008201600f9054906101000a900461ffff1661ffff1661ffff1681526020016000820160119054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016000820160159054906101000a900461ffff1661ffff1661ffff1681526020016000820160179054906101000a900461ffff1661ffff1661ffff1681526020016000820160199054906101000a900461ffff1661ffff1661ffff16815260200160008201601b9054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820160009054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820160049054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815260200160018201600c9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016001820160149054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820160189054906101000a900460ff161515151581526020016001820160199054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681525050613a7d565b516120af9082615225565b60018b01549091506000906c01000000000000000000000000900467ffffffffffffffff166120ee836dffffffffffffffffffffffffffff8a16614fa5565b6120f89190614fa5565b90507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8816838261212f67ffffffffffffffff8d168a614fa5565b6121399190615225565b6121439190615225565b61214d9190614fbc565b9e9d5050505050505050505050505050565b6121676123e7565b610e2f81613b38565b67ffffffffffffffff811660009081526004602090815260408083208151808301909252547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811682527c0100000000000000000000000000000000000000000000000000000000900463ffffffff1691810182905282918203612228576040517f2e59db3a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610f9b565b6000816020015163ffffffff16426122409190615238565b90507f000000000000000000000000000000000000000000000000000000000000000063ffffffff168111156122e1576040517ff08bcb3e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8616600482015263ffffffff7f000000000000000000000000000000000000000000000000000000000000000016602482015260448101829052606401610f9b565b6122ea866122fb565b9151919350909150505b9250929050565b60008061230783611459565b9050806020015163ffffffff166000148061233f575080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16155b1561238e576040517f06439c6b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610f9b565b5192915050565b60606000610b3283613c2d565b6123ad600233613c89565b6123e5576040517fd86ad9cf000000000000000000000000000000000000000000000000000000008152336004820152602401610f9b565b565b60005473ffffffffffffffffffffffffffffffffffffffff1633146123e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610f9b565b60005b8151811015610f1a5760008282815181106124885761248861505f565b60209081029190910181015180518183015173ffffffffffffffffffffffffffffffffffffffff80831660008181526006875260409081902084518154868a018051929096167fffffffffffffffffffffff00000000000000000000000000000000000000000090911681177401000000000000000000000000000000000000000060ff9384160217909255825191825293519093169683019690965293955091939092917f08a5f7f5bb38a81d8e43aca13ecd76431dbf8816ae4699affff7b00b2fc1c464910160405180910390a250505080600101905061246b565b60005b82518110156128895760008382815181106125865761258661505f565b6020026020010151905060008160000151905060005b82602001515181101561287b576000836020015182815181106125c1576125c161505f565b60200260200101516020015190506000846020015183815181106125e7576125e761505f565b60200260200101516000015190506020826080015163ffffffff1610156126645760808201516040517f24ecdc0200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015263ffffffff9091166024820152604401610f9b565b67ffffffffffffffff8416600081815260096020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168085529083529281902086518154938801518389015160608a015160808b015160a08c015115157201000000000000000000000000000000000000027fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff63ffffffff9283166e01000000000000000000000000000002167fffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffffff9383166a0100000000000000000000027fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff61ffff9096166801000000000000000002959095167fffffffffffffffffffffffffffffffffffff000000000000ffffffffffffffff968416640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000909b16939097169290921798909817939093169390931717919091161792909217909155519091907f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b590612869908690600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60405180910390a3505060010161259c565b505050806001019050612569565b5060005b815181101561296d5760008282815181106128aa576128aa61505f565b602002602001015160000151905060008383815181106128cc576128cc61505f565b60209081029190910181015181015167ffffffffffffffff8416600081815260098452604080822073ffffffffffffffffffffffffffffffffffffffff8516808452955280822080547fffffffffffffffffffffffffff000000000000000000000000000000000000001690555192945090917f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b9190a3505060010161288d565b505050565b60005b8251811015612a15576129ab8382815181106129935761299361505f565b6020026020010151600c613cb890919063ffffffff16565b15612a0d578281815181106129c2576129c261505f565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba2360405160405180910390a25b600101612975565b5060005b815181101561296d57612a4f828281518110612a3757612a3761505f565b6020026020010151600c613cda90919063ffffffff16565b15612ab157818181518110612a6657612a6661505f565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f9160405160405180910390a25b600101612a19565b602081015160005b8151811015612b54576000828281518110612ade57612ade61505f565b60200260200101519050612afc816002613cda90919063ffffffff16565b15612b4b5760405173ffffffffffffffffffffffffffffffffffffffff821681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda775809060200160405180910390a15b50600101612ac1565b50815160005b8151811015610e18576000828281518110612b7757612b7761505f565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612be7576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612bf2600282613cb8565b5060405173ffffffffffffffffffffffffffffffffffffffff821681527feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef9060200160405180910390a150600101612b5a565b7fd7ed2ad4000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831601610f1a5761296d81613cfc565b60408051808201909152600080825260208201526000839003612cd857506040805180820190915267ffffffffffffffff8216815260006020820152610b32565b6000612ce4848661524b565b90506000612cf58560048189615291565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293505050507fffffffff0000000000000000000000000000000000000000000000000000000082167fe7e230f00000000000000000000000000000000000000000000000000000000001612d925780806020019051810190612d8991906152bb565b92505050610b32565b7f6859a837000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831601612e0e57604051806040016040528082806020019051810190612dfa91906152e7565b815260006020909101529250610b32915050565b6040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8151811015610f1a576000828281518110612e6057612e6061505f565b60200260200101516000015190506000838381518110612e8257612e8261505f565b60209081029190910181015181015173ffffffffffffffffffffffffffffffffffffffff841660008181526007845260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff85169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a25050600101612e43565b60005b8151811015610f1a576000828281518110612f4a57612f4a61505f565b602002602001015190506000838381518110612f6857612f6861505f565b60200260200101516000015190506000826020015190508167ffffffffffffffff1660001480612fa5575061018081015167ffffffffffffffff16155b80612ff757506102008101517fffffffff00000000000000000000000000000000000000000000000000000000167f2812d52c0000000000000000000000000000000000000000000000000000000014155b8061300d5750602081610160015163ffffffff16105b15613050576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff83166004820152602401610f9b565b67ffffffffffffffff8216600090815260086020526040812060010154790100000000000000000000000000000000000000000000000000900460e01b7fffffffff000000000000000000000000000000000000000000000000000000001690036130fc578167ffffffffffffffff167fd883f18d08e8ea5c2af22d13fb6cab86cc057d795d2735f83a8ff137d39d1b8f826040516130ef9190614de0565b60405180910390a261313f565b8167ffffffffffffffff167fd42d3a670a4f1ab5d8703efa22bb17041446365cfcfa33980ced685a080cf7cb826040516131369190614de0565b60405180910390a25b80600860008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101a082015181600101600c6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101c08201518160010160146101000a81548163ffffffff021916908363ffffffff1602179055506101e08201518160010160186101000a81548160ff0219169083151502179055506102008201518160010160196101000a81548163ffffffff021916908360e01c0217905550905050505050806001019050612f2d565b604080518082019091526000808252602082015260008260000151905060008173ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa15801561346d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613491919061531a565b50505091505060008112156134d2576040517f10cb51d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000819050600085602001518473ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015613529573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061354d919061536a565b6135579190615387565b905060248160ff16111561358c576135706024826153a0565b61357b90600a6154d9565b6135859083614fbc565b91506135af565b6135978160246153a0565b6135a290600a6154d9565b6135ac9083614fa5565b91505b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff821115613605576040517f10cb51d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50604080518082019091527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff909116815263ffffffff42166020820152949350505050565b836040015163ffffffff1683111561369f5760408085015190517f8693378900000000000000000000000000000000000000000000000000000000815263ffffffff909116600482015260248101849052604401610f9b565b836020015161ffff168211156136e1576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e1884610200015182612c45565b6000808083815b818110156139c05760008787838181106137135761371361505f565b90506040020180360381019061372991906154e8565b67ffffffffffffffff8c166000908152600960209081526040808320845173ffffffffffffffffffffffffffffffffffffffff168452825291829020825160c081018452905463ffffffff8082168352640100000000820481169383019390935268010000000000000000810461ffff16938201939093526a01000000000000000000008304821660608201526e01000000000000000000000000000083049091166080820152720100000000000000000000000000000000000090910460ff16151560a082018190529192509061384f576101208d01516138169061ffff16662386f26fc10000614fa5565b6138209088615225565b96508c6101400151866138339190615521565b95508c6101600151856138469190615521565b945050506139b8565b604081015160009061ffff16156139085760008c73ffffffffffffffffffffffffffffffffffffffff16846000015173ffffffffffffffffffffffffffffffffffffffff16146138ab5783516138a4906122fb565b90506138ae565b508a5b620186a0836040015161ffff166138f08660200151847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16613d5790919063ffffffff16565b6138fa9190614fa5565b6139049190614fbc565b9150505b60608201516139179088615521565b96508160800151866139299190615521565b82519096506000906139489063ffffffff16662386f26fc10000614fa5565b9050808210156139675761395c818a615225565b9850505050506139b8565b6000836020015163ffffffff16662386f26fc100006139869190614fa5565b9050808311156139a65761399a818b615225565b995050505050506139b8565b6139b0838b615225565b995050505050505b6001016136f7565b505096509650969350505050565b60008063ffffffff83166139e360e086614fa5565b6139ef876101c0615225565b6139f99190615225565b613a039190615225565b905060008760c0015163ffffffff168860e0015161ffff1683613a269190614fa5565b613a309190615225565b61010089015190915061ffff16613a576dffffffffffffffffffffffffffff891683614fa5565b613a619190614fa5565b613a7190655af3107a4000614fa5565b98975050505050505050565b60408051808201909152600080825260208201526000613aa38585856101800151612c97565b9050826060015163ffffffff1681600001511115613aed576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826101e001518015613b0157508060200151155b15610b2f576040517fee433e9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff821603613bb7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610f9b565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b606081600001805480602002602001604051908101604052809291908181526020018280548015613c7d57602002820191906000526020600020905b815481526020019060010190808311613c69575b50505050509050919050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515610b32565b6000610b328373ffffffffffffffffffffffffffffffffffffffff8416613d94565b6000610b328373ffffffffffffffffffffffffffffffffffffffff8416613de3565b60008151602014613d3b57816040517f8d666f60000000000000000000000000000000000000000000000000000000008152600401610f9b9190614059565b610ef782806020019051810190613d5291906152e7565b613edd565b6000670de0b6b3a7640000613d8a837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8616614fa5565b610b329190614fbc565b6000818152600183016020526040812054613ddb57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610ef7565b506000610ef7565b60008181526001830160205260408120548015613ecc576000613e07600183615238565b8554909150600090613e1b90600190615238565b9050818114613e80576000866000018281548110613e3b57613e3b61505f565b9060005260206000200154905080876000018481548110613e5e57613e5e61505f565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080613e9157613e9161553e565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610ef7565b6000915050610ef7565b5092915050565b600073ffffffffffffffffffffffffffffffffffffffff821180613f02575061040082105b15613f7157604080516020810184905201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f8d666f60000000000000000000000000000000000000000000000000000000008252610f9b91600401614059565b5090565b803573ffffffffffffffffffffffffffffffffffffffff81168114613f9957600080fd5b919050565b600080600060608486031215613fb357600080fd5b613fbc84613f75565b925060208401359150613fd160408501613f75565b90509250925092565b600060208284031215613fec57600080fd5b610b3282613f75565b6000815180845260005b8181101561401b57602081850181015186830182015201613fff565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610b326020830184613ff5565b6020808252825182820181905260009190848201906040850190845b818110156140ba57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101614088565b50909695505050505050565b6000602082840312156140d857600080fd5b813567ffffffffffffffff8111156140ef57600080fd5b820160408185031215610b3257600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561415357614153614101565b60405290565b60405160c0810167ffffffffffffffff8111828210171561415357614153614101565b604051610220810167ffffffffffffffff8111828210171561415357614153614101565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156141e7576141e7614101565b604052919050565b600067ffffffffffffffff82111561420957614209614101565b5060051b60200190565b60ff81168114610e2f57600080fd5b6000602080838503121561423557600080fd5b823567ffffffffffffffff81111561424c57600080fd5b8301601f8101851361425d57600080fd5b803561427061426b826141ef565b6141a0565b8181526060918202830184019184820191908884111561428f57600080fd5b938501935b8385101561432f57848903818112156142ad5760008081fd5b6142b5614130565b6142be87613f75565b81526040807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0840112156142f25760008081fd5b6142fa614130565b9250614307898901613f75565b835287013561431581614213565b828901528088019190915283529384019391850191614294565b50979650505050505050565b6000806020838503121561434e57600080fd5b823567ffffffffffffffff8082111561436657600080fd5b818501915085601f83011261437a57600080fd5b81358181111561438957600080fd5b8660208260051b850101111561439e57600080fd5b60209290920196919550909350505050565b602080825282518282018190526000919060409081850190868401855b8281101561441e5761440e84835180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16825260209081015163ffffffff16910152565b92840192908501906001016143cd565b5091979650505050505050565b803567ffffffffffffffff81168114613f9957600080fd5b60006020828403121561445557600080fd5b610b328261442b565b81517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815260208083015163ffffffff169082015260408101610ef7565b803563ffffffff81168114613f9957600080fd5b803561ffff81168114613f9957600080fd5b8015158114610e2f57600080fd5b8035613f99816144bf565b600082601f8301126144e957600080fd5b813560206144f961426b836141ef565b82815260069290921b8401810191818101908684111561451857600080fd5b8286015b8481101561456557604081890312156145355760008081fd5b61453d614130565b6145468261442b565b8152614553858301613f75565b8186015283529183019160400161451c565b509695505050505050565b6000806040838503121561458357600080fd5b67ffffffffffffffff8335111561459957600080fd5b83601f8435850101126145ab57600080fd5b6145bb61426b84358501356141ef565b8335840180358083526020808401939260059290921b909101018610156145e157600080fd5b602085358601015b85358601803560051b016020018110156147ee5767ffffffffffffffff8135111561461357600080fd5b8035863587010160407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828a0301121561464c57600080fd5b614654614130565b6146606020830161442b565b815267ffffffffffffffff6040830135111561467b57600080fd5b88603f60408401358401011261469057600080fd5b6146a661426b60206040850135850101356141ef565b6020604084810135850182810135808552928401939260e00201018b10156146cd57600080fd5b6040808501358501015b6040858101358601602081013560e00201018110156147cf5760e0818d03121561470057600080fd5b614708614130565b61471182613f75565b815260c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838f0301121561474557600080fd5b61474d614159565b61475960208401614499565b815261476760408401614499565b6020820152614778606084016144ad565b604082015261478960808401614499565b606082015261479a60a08401614499565b60808201526147ac60c08401356144bf565b60c083013560a0820152602082810191909152908452929092019160e0016146d7565b50806020840152505080855250506020830192506020810190506145e9565b5092505067ffffffffffffffff6020840135111561480b57600080fd5b61481b84602085013585016144d8565b90509250929050565b600082601f83011261483557600080fd5b8135602061484561426b836141ef565b8083825260208201915060208460051b87010193508684111561486757600080fd5b602086015b848110156145655761487d81613f75565b835291830191830161486c565b6000806040838503121561489d57600080fd5b823567ffffffffffffffff808211156148b557600080fd5b6148c186838701614824565b935060208501359150808211156148d757600080fd5b506148e485828601614824565b9150509250929050565b6000806040838503121561490157600080fd5b61490a8361442b565b915061481b60208401613f75565b60006020828403121561492a57600080fd5b813567ffffffffffffffff8082111561494257600080fd5b908301906040828603121561495657600080fd5b61495e614130565b82358281111561496d57600080fd5b61497987828601614824565b82525060208301358281111561498e57600080fd5b61499a87828601614824565b60208301525095945050505050565b6000806000604084860312156149be57600080fd5b833567ffffffffffffffff808211156149d657600080fd5b9085019061018082880312156149eb57600080fd5b90935060208501359080821115614a0157600080fd5b818601915086601f830112614a1557600080fd5b813581811115614a2457600080fd5b8760208260061b8501011115614a3957600080fd5b6020830194508093505050509250925092565b8381528215156020820152606060408201526000614a6d6060830184613ff5565b95945050505050565b60006020808385031215614a8957600080fd5b823567ffffffffffffffff811115614aa057600080fd5b8301601f81018513614ab157600080fd5b8035614abf61426b826141ef565b81815260069190911b82018301908381019087831115614ade57600080fd5b928401925b82841015614b305760408489031215614afc5760008081fd5b614b04614130565b614b0d85613f75565b8152614b1a86860161442b565b8187015282526040939093019290840190614ae3565b979650505050505050565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114613f9957600080fd5b60006020808385031215614b7e57600080fd5b823567ffffffffffffffff811115614b9557600080fd5b8301601f81018513614ba657600080fd5b8035614bb461426b826141ef565b8181526102409182028301840191848201919088841115614bd457600080fd5b938501935b8385101561432f5784890381811215614bf25760008081fd5b614bfa614130565b614c038761442b565b8152610220807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084011215614c385760008081fd5b614c4061417c565b9250614c4d8989016144cd565b83526040614c5c818a016144ad565b8a8501526060614c6d818b01614499565b8286015260809150614c80828b01614499565b9085015260a0614c918a8201614499565b8286015260c09150614ca4828b016144ad565b9085015260e0614cb58a8201614499565b828601526101009150614cc9828b016144ad565b90850152610120614cdb8a82016144ad565b828601526101409150614cef828b016144ad565b90850152610160614d018a8201614499565b828601526101809150614d15828b01614499565b908501526101a0614d278a820161442b565b828601526101c09150614d3b828b0161442b565b908501526101e0614d4d8a8201614499565b828601526102009150614d61828b016144cd565b90850152614d70898301614b3b565b90840152508088019190915283529384019391850191614bd9565b60008060408385031215614d9e57600080fd5b614da78361442b565b9150602083013567ffffffffffffffff811115614dc357600080fd5b830160a08186031215614dd557600080fd5b809150509250929050565b81511515815261022081016020830151614e00602084018261ffff169052565b506040830151614e18604084018263ffffffff169052565b506060830151614e30606084018263ffffffff169052565b506080830151614e48608084018263ffffffff169052565b5060a0830151614e5e60a084018261ffff169052565b5060c0830151614e7660c084018263ffffffff169052565b5060e0830151614e8c60e084018261ffff169052565b506101008381015161ffff9081169184019190915261012080850151909116908301526101408084015163ffffffff90811691840191909152610160808501518216908401526101808085015167ffffffffffffffff908116918501919091526101a080860151909116908401526101c080850151909116908301526101e080840151151590830152610200808401517fffffffff000000000000000000000000000000000000000000000000000000008116828501525b505092915050565b60008060408385031215614f5f57600080fd5b614f6883613f75565b915061481b6020840161442b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082028115828204841417610ef757610ef7614f76565b600082614ff2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261502c57600080fd5b83018035915067ffffffffffffffff82111561504757600080fd5b6020019150600681901b36038213156122f457600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81168114613f9957600080fd5b6000604082840312156150cc57600080fd5b6150d4614130565b6150dd83613f75565b81526150eb6020840161508e565b60208201529392505050565b60006040828403121561510957600080fd5b615111614130565b6150dd8361442b565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261514f57600080fd5b83018035915067ffffffffffffffff82111561516a57600080fd5b6020019150600581901b36038213156122f457600080fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818336030181126151b657600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126151f557600080fd5b83018035915067ffffffffffffffff82111561521057600080fd5b6020019150368190038213156122f457600080fd5b80820180821115610ef757610ef7614f76565b81810381811115610ef757610ef7614f76565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015614f445760049490940360031b84901b1690921692915050565b600080858511156152a157600080fd5b838611156152ae57600080fd5b5050820193919092039150565b6000604082840312156152cd57600080fd5b6152d5614130565b8251815260208301516150eb816144bf565b6000602082840312156152f957600080fd5b5051919050565b805169ffffffffffffffffffff81168114613f9957600080fd5b600080600080600060a0868803121561533257600080fd5b61533b86615300565b945060208601519350604086015192506060860151915061535e60808701615300565b90509295509295909350565b60006020828403121561537c57600080fd5b8151610b3281614213565b60ff8181168382160190811115610ef757610ef7614f76565b60ff8281168282160390811115610ef757610ef7614f76565b600181815b8085111561541257817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156153f8576153f8614f76565b8085161561540557918102915b93841c93908002906153be565b509250929050565b60008261542957506001610ef7565b8161543657506000610ef7565b816001811461544c576002811461545657615472565b6001915050610ef7565b60ff84111561546757615467614f76565b50506001821b610ef7565b5060208310610133831016604e8410600b8410161715615495575081810a610ef7565b61549f83836153b9565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156154d1576154d1614f76565b029392505050565b6000610b3260ff84168361541a565b6000604082840312156154fa57600080fd5b615502614130565b61550b83613f75565b8152602083013560208201528091505092915050565b63ffffffff818116838216019080821115613ed657613ed6614f76565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000818000a",
}

var PriceRegistryABI = PriceRegistryMetaData.ABI

var PriceRegistryBin = PriceRegistryMetaData.Bin

func DeployPriceRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig PriceRegistryStaticConfig, priceUpdaters []common.Address, feeTokens []common.Address, tokenPriceFeeds []PriceRegistryTokenPriceFeedUpdate, tokenTransferFeeConfigArgs []PriceRegistryTokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs []PriceRegistryPremiumMultiplierWeiPerEthArgs, destChainConfigArgs []PriceRegistryDestChainDynamicConfigArgs) (common.Address, *types.Transaction, *PriceRegistry, error) {
	parsed, err := PriceRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PriceRegistryBin), backend, staticConfig, priceUpdaters, feeTokens, tokenPriceFeeds, tokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs, destChainConfigArgs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriceRegistry{address: address, abi: *parsed, PriceRegistryCaller: PriceRegistryCaller{contract: contract}, PriceRegistryTransactor: PriceRegistryTransactor{contract: contract}, PriceRegistryFilterer: PriceRegistryFilterer{contract: contract}}, nil
}

type PriceRegistry struct {
	address common.Address
	abi     abi.ABI
	PriceRegistryCaller
	PriceRegistryTransactor
	PriceRegistryFilterer
}

type PriceRegistryCaller struct {
	contract *bind.BoundContract
}

type PriceRegistryTransactor struct {
	contract *bind.BoundContract
}

type PriceRegistryFilterer struct {
	contract *bind.BoundContract
}

type PriceRegistrySession struct {
	Contract     *PriceRegistry
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type PriceRegistryCallerSession struct {
	Contract *PriceRegistryCaller
	CallOpts bind.CallOpts
}

type PriceRegistryTransactorSession struct {
	Contract     *PriceRegistryTransactor
	TransactOpts bind.TransactOpts
}

type PriceRegistryRaw struct {
	Contract *PriceRegistry
}

type PriceRegistryCallerRaw struct {
	Contract *PriceRegistryCaller
}

type PriceRegistryTransactorRaw struct {
	Contract *PriceRegistryTransactor
}

func NewPriceRegistry(address common.Address, backend bind.ContractBackend) (*PriceRegistry, error) {
	abi, err := abi.JSON(strings.NewReader(PriceRegistryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindPriceRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceRegistry{address: address, abi: abi, PriceRegistryCaller: PriceRegistryCaller{contract: contract}, PriceRegistryTransactor: PriceRegistryTransactor{contract: contract}, PriceRegistryFilterer: PriceRegistryFilterer{contract: contract}}, nil
}

func NewPriceRegistryCaller(address common.Address, caller bind.ContractCaller) (*PriceRegistryCaller, error) {
	contract, err := bindPriceRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryCaller{contract: contract}, nil
}

func NewPriceRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceRegistryTransactor, error) {
	contract, err := bindPriceRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryTransactor{contract: contract}, nil
}

func NewPriceRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceRegistryFilterer, error) {
	contract, err := bindPriceRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryFilterer{contract: contract}, nil
}

func bindPriceRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_PriceRegistry *PriceRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceRegistry.Contract.PriceRegistryCaller.contract.Call(opts, result, method, params...)
}

func (_PriceRegistry *PriceRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRegistry.Contract.PriceRegistryTransactor.contract.Transfer(opts)
}

func (_PriceRegistry *PriceRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceRegistry.Contract.PriceRegistryTransactor.contract.Transact(opts, method, params...)
}

func (_PriceRegistry *PriceRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PriceRegistry.Contract.contract.Call(opts, result, method, params...)
}

func (_PriceRegistry *PriceRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRegistry.Contract.contract.Transfer(opts)
}

func (_PriceRegistry *PriceRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceRegistry.Contract.contract.Transact(opts, method, params...)
}

func (_PriceRegistry *PriceRegistryCaller) ConvertTokenAmount(opts *bind.CallOpts, fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "convertTokenAmount", fromToken, fromTokenAmount, toToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) ConvertTokenAmount(fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error) {
	return _PriceRegistry.Contract.ConvertTokenAmount(&_PriceRegistry.CallOpts, fromToken, fromTokenAmount, toToken)
}

func (_PriceRegistry *PriceRegistryCallerSession) ConvertTokenAmount(fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error) {
	return _PriceRegistry.Contract.ConvertTokenAmount(&_PriceRegistry.CallOpts, fromToken, fromTokenAmount, toToken)
}

func (_PriceRegistry *PriceRegistryCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _PriceRegistry.Contract.GetAllAuthorizedCallers(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _PriceRegistry.Contract.GetAllAuthorizedCallers(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCaller) GetDestChainDynamicConfig(opts *bind.CallOpts, destChainSelector uint64) (PriceRegistryDestChainDynamicConfig, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getDestChainDynamicConfig", destChainSelector)

	if err != nil {
		return *new(PriceRegistryDestChainDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(PriceRegistryDestChainDynamicConfig)).(*PriceRegistryDestChainDynamicConfig)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetDestChainDynamicConfig(destChainSelector uint64) (PriceRegistryDestChainDynamicConfig, error) {
	return _PriceRegistry.Contract.GetDestChainDynamicConfig(&_PriceRegistry.CallOpts, destChainSelector)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetDestChainDynamicConfig(destChainSelector uint64) (PriceRegistryDestChainDynamicConfig, error) {
	return _PriceRegistry.Contract.GetDestChainDynamicConfig(&_PriceRegistry.CallOpts, destChainSelector)
}

func (_PriceRegistry *PriceRegistryCaller) GetDestinationChainGasPrice(opts *bind.CallOpts, destChainSelector uint64) (InternalTimestampedPackedUint224, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getDestinationChainGasPrice", destChainSelector)

	if err != nil {
		return *new(InternalTimestampedPackedUint224), err
	}

	out0 := *abi.ConvertType(out[0], new(InternalTimestampedPackedUint224)).(*InternalTimestampedPackedUint224)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetDestinationChainGasPrice(destChainSelector uint64) (InternalTimestampedPackedUint224, error) {
	return _PriceRegistry.Contract.GetDestinationChainGasPrice(&_PriceRegistry.CallOpts, destChainSelector)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetDestinationChainGasPrice(destChainSelector uint64) (InternalTimestampedPackedUint224, error) {
	return _PriceRegistry.Contract.GetDestinationChainGasPrice(&_PriceRegistry.CallOpts, destChainSelector)
}

func (_PriceRegistry *PriceRegistryCaller) GetFeeTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getFeeTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetFeeTokens() ([]common.Address, error) {
	return _PriceRegistry.Contract.GetFeeTokens(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetFeeTokens() ([]common.Address, error) {
	return _PriceRegistry.Contract.GetFeeTokens(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCaller) GetPremiumMultiplierWeiPerEth(opts *bind.CallOpts, token common.Address) (uint64, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getPremiumMultiplierWeiPerEth", token)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetPremiumMultiplierWeiPerEth(token common.Address) (uint64, error) {
	return _PriceRegistry.Contract.GetPremiumMultiplierWeiPerEth(&_PriceRegistry.CallOpts, token)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetPremiumMultiplierWeiPerEth(token common.Address) (uint64, error) {
	return _PriceRegistry.Contract.GetPremiumMultiplierWeiPerEth(&_PriceRegistry.CallOpts, token)
}

func (_PriceRegistry *PriceRegistryCaller) GetStaticConfig(opts *bind.CallOpts) (PriceRegistryStaticConfig, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(PriceRegistryStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(PriceRegistryStaticConfig)).(*PriceRegistryStaticConfig)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetStaticConfig() (PriceRegistryStaticConfig, error) {
	return _PriceRegistry.Contract.GetStaticConfig(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetStaticConfig() (PriceRegistryStaticConfig, error) {
	return _PriceRegistry.Contract.GetStaticConfig(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCaller) GetTokenAndGasPrices(opts *bind.CallOpts, token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

	error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getTokenAndGasPrices", token, destChainSelector)

	outstruct := new(GetTokenAndGasPrices)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasPriceValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_PriceRegistry *PriceRegistrySession) GetTokenAndGasPrices(token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

	error) {
	return _PriceRegistry.Contract.GetTokenAndGasPrices(&_PriceRegistry.CallOpts, token, destChainSelector)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetTokenAndGasPrices(token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

	error) {
	return _PriceRegistry.Contract.GetTokenAndGasPrices(&_PriceRegistry.CallOpts, token, destChainSelector)
}

func (_PriceRegistry *PriceRegistryCaller) GetTokenPrice(opts *bind.CallOpts, token common.Address) (InternalTimestampedPackedUint224, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getTokenPrice", token)

	if err != nil {
		return *new(InternalTimestampedPackedUint224), err
	}

	out0 := *abi.ConvertType(out[0], new(InternalTimestampedPackedUint224)).(*InternalTimestampedPackedUint224)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetTokenPrice(token common.Address) (InternalTimestampedPackedUint224, error) {
	return _PriceRegistry.Contract.GetTokenPrice(&_PriceRegistry.CallOpts, token)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetTokenPrice(token common.Address) (InternalTimestampedPackedUint224, error) {
	return _PriceRegistry.Contract.GetTokenPrice(&_PriceRegistry.CallOpts, token)
}

func (_PriceRegistry *PriceRegistryCaller) GetTokenPriceFeedConfig(opts *bind.CallOpts, token common.Address) (IPriceRegistryTokenPriceFeedConfig, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getTokenPriceFeedConfig", token)

	if err != nil {
		return *new(IPriceRegistryTokenPriceFeedConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPriceRegistryTokenPriceFeedConfig)).(*IPriceRegistryTokenPriceFeedConfig)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetTokenPriceFeedConfig(token common.Address) (IPriceRegistryTokenPriceFeedConfig, error) {
	return _PriceRegistry.Contract.GetTokenPriceFeedConfig(&_PriceRegistry.CallOpts, token)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetTokenPriceFeedConfig(token common.Address) (IPriceRegistryTokenPriceFeedConfig, error) {
	return _PriceRegistry.Contract.GetTokenPriceFeedConfig(&_PriceRegistry.CallOpts, token)
}

func (_PriceRegistry *PriceRegistryCaller) GetTokenPrices(opts *bind.CallOpts, tokens []common.Address) ([]InternalTimestampedPackedUint224, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getTokenPrices", tokens)

	if err != nil {
		return *new([]InternalTimestampedPackedUint224), err
	}

	out0 := *abi.ConvertType(out[0], new([]InternalTimestampedPackedUint224)).(*[]InternalTimestampedPackedUint224)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetTokenPrices(tokens []common.Address) ([]InternalTimestampedPackedUint224, error) {
	return _PriceRegistry.Contract.GetTokenPrices(&_PriceRegistry.CallOpts, tokens)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetTokenPrices(tokens []common.Address) ([]InternalTimestampedPackedUint224, error) {
	return _PriceRegistry.Contract.GetTokenPrices(&_PriceRegistry.CallOpts, tokens)
}

func (_PriceRegistry *PriceRegistryCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (PriceRegistryTokenTransferFeeConfig, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getTokenTransferFeeConfig", destChainSelector, token)

	if err != nil {
		return *new(PriceRegistryTokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(PriceRegistryTokenTransferFeeConfig)).(*PriceRegistryTokenTransferFeeConfig)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetTokenTransferFeeConfig(destChainSelector uint64, token common.Address) (PriceRegistryTokenTransferFeeConfig, error) {
	return _PriceRegistry.Contract.GetTokenTransferFeeConfig(&_PriceRegistry.CallOpts, destChainSelector, token)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetTokenTransferFeeConfig(destChainSelector uint64, token common.Address) (PriceRegistryTokenTransferFeeConfig, error) {
	return _PriceRegistry.Contract.GetTokenTransferFeeConfig(&_PriceRegistry.CallOpts, destChainSelector, token)
}

func (_PriceRegistry *PriceRegistryCaller) GetValidatedFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getValidatedFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetValidatedFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _PriceRegistry.Contract.GetValidatedFee(&_PriceRegistry.CallOpts, destChainSelector, message)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetValidatedFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _PriceRegistry.Contract.GetValidatedFee(&_PriceRegistry.CallOpts, destChainSelector, message)
}

func (_PriceRegistry *PriceRegistryCaller) GetValidatedRampMessageParams(opts *bind.CallOpts, message InternalEVM2AnyRampMessage, sourceTokenAmounts []ClientEVMTokenAmount) (GetValidatedRampMessageParams,

	error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getValidatedRampMessageParams", message, sourceTokenAmounts)

	outstruct := new(GetValidatedRampMessageParams)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MsgFeeJuels = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsOutOfOrderExecution = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.ConvertedExtraArgs = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

func (_PriceRegistry *PriceRegistrySession) GetValidatedRampMessageParams(message InternalEVM2AnyRampMessage, sourceTokenAmounts []ClientEVMTokenAmount) (GetValidatedRampMessageParams,

	error) {
	return _PriceRegistry.Contract.GetValidatedRampMessageParams(&_PriceRegistry.CallOpts, message, sourceTokenAmounts)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetValidatedRampMessageParams(message InternalEVM2AnyRampMessage, sourceTokenAmounts []ClientEVMTokenAmount) (GetValidatedRampMessageParams,

	error) {
	return _PriceRegistry.Contract.GetValidatedRampMessageParams(&_PriceRegistry.CallOpts, message, sourceTokenAmounts)
}

func (_PriceRegistry *PriceRegistryCaller) GetValidatedTokenPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getValidatedTokenPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetValidatedTokenPrice(token common.Address) (*big.Int, error) {
	return _PriceRegistry.Contract.GetValidatedTokenPrice(&_PriceRegistry.CallOpts, token)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetValidatedTokenPrice(token common.Address) (*big.Int, error) {
	return _PriceRegistry.Contract.GetValidatedTokenPrice(&_PriceRegistry.CallOpts, token)
}

func (_PriceRegistry *PriceRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) Owner() (common.Address, error) {
	return _PriceRegistry.Contract.Owner(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCallerSession) Owner() (common.Address, error) {
	return _PriceRegistry.Contract.Owner(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) TypeAndVersion() (string, error) {
	return _PriceRegistry.Contract.TypeAndVersion(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryCallerSession) TypeAndVersion() (string, error) {
	return _PriceRegistry.Contract.TypeAndVersion(&_PriceRegistry.CallOpts)
}

func (_PriceRegistry *PriceRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "acceptOwnership")
}

func (_PriceRegistry *PriceRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _PriceRegistry.Contract.AcceptOwnership(&_PriceRegistry.TransactOpts)
}

func (_PriceRegistry *PriceRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PriceRegistry.Contract.AcceptOwnership(&_PriceRegistry.TransactOpts)
}

func (_PriceRegistry *PriceRegistryTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_PriceRegistry *PriceRegistrySession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _PriceRegistry.Contract.ApplyAuthorizedCallerUpdates(&_PriceRegistry.TransactOpts, authorizedCallerArgs)
}

func (_PriceRegistry *PriceRegistryTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _PriceRegistry.Contract.ApplyAuthorizedCallerUpdates(&_PriceRegistry.TransactOpts, authorizedCallerArgs)
}

func (_PriceRegistry *PriceRegistryTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []PriceRegistryDestChainDynamicConfigArgs) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_PriceRegistry *PriceRegistrySession) ApplyDestChainConfigUpdates(destChainConfigArgs []PriceRegistryDestChainDynamicConfigArgs) (*types.Transaction, error) {
	return _PriceRegistry.Contract.ApplyDestChainConfigUpdates(&_PriceRegistry.TransactOpts, destChainConfigArgs)
}

func (_PriceRegistry *PriceRegistryTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []PriceRegistryDestChainDynamicConfigArgs) (*types.Transaction, error) {
	return _PriceRegistry.Contract.ApplyDestChainConfigUpdates(&_PriceRegistry.TransactOpts, destChainConfigArgs)
}

func (_PriceRegistry *PriceRegistryTransactor) ApplyFeeTokensUpdates(opts *bind.TransactOpts, feeTokensToAdd []common.Address, feeTokensToRemove []common.Address) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "applyFeeTokensUpdates", feeTokensToAdd, feeTokensToRemove)
}

func (_PriceRegistry *PriceRegistrySession) ApplyFeeTokensUpdates(feeTokensToAdd []common.Address, feeTokensToRemove []common.Address) (*types.Transaction, error) {
	return _PriceRegistry.Contract.ApplyFeeTokensUpdates(&_PriceRegistry.TransactOpts, feeTokensToAdd, feeTokensToRemove)
}

func (_PriceRegistry *PriceRegistryTransactorSession) ApplyFeeTokensUpdates(feeTokensToAdd []common.Address, feeTokensToRemove []common.Address) (*types.Transaction, error) {
	return _PriceRegistry.Contract.ApplyFeeTokensUpdates(&_PriceRegistry.TransactOpts, feeTokensToAdd, feeTokensToRemove)
}

func (_PriceRegistry *PriceRegistryTransactor) ApplyPremiumMultiplierWeiPerEthUpdates(opts *bind.TransactOpts, premiumMultiplierWeiPerEthArgs []PriceRegistryPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "applyPremiumMultiplierWeiPerEthUpdates", premiumMultiplierWeiPerEthArgs)
}

func (_PriceRegistry *PriceRegistrySession) ApplyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs []PriceRegistryPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _PriceRegistry.Contract.ApplyPremiumMultiplierWeiPerEthUpdates(&_PriceRegistry.TransactOpts, premiumMultiplierWeiPerEthArgs)
}

func (_PriceRegistry *PriceRegistryTransactorSession) ApplyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs []PriceRegistryPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _PriceRegistry.Contract.ApplyPremiumMultiplierWeiPerEthUpdates(&_PriceRegistry.TransactOpts, premiumMultiplierWeiPerEthArgs)
}

func (_PriceRegistry *PriceRegistryTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []PriceRegistryTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []PriceRegistryTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_PriceRegistry *PriceRegistrySession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []PriceRegistryTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []PriceRegistryTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _PriceRegistry.Contract.ApplyTokenTransferFeeConfigUpdates(&_PriceRegistry.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_PriceRegistry *PriceRegistryTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []PriceRegistryTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []PriceRegistryTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _PriceRegistry.Contract.ApplyTokenTransferFeeConfigUpdates(&_PriceRegistry.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_PriceRegistry *PriceRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "transferOwnership", to)
}

func (_PriceRegistry *PriceRegistrySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _PriceRegistry.Contract.TransferOwnership(&_PriceRegistry.TransactOpts, to)
}

func (_PriceRegistry *PriceRegistryTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _PriceRegistry.Contract.TransferOwnership(&_PriceRegistry.TransactOpts, to)
}

func (_PriceRegistry *PriceRegistryTransactor) UpdatePrices(opts *bind.TransactOpts, priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "updatePrices", priceUpdates)
}

func (_PriceRegistry *PriceRegistrySession) UpdatePrices(priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _PriceRegistry.Contract.UpdatePrices(&_PriceRegistry.TransactOpts, priceUpdates)
}

func (_PriceRegistry *PriceRegistryTransactorSession) UpdatePrices(priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _PriceRegistry.Contract.UpdatePrices(&_PriceRegistry.TransactOpts, priceUpdates)
}

func (_PriceRegistry *PriceRegistryTransactor) UpdateTokenPriceFeeds(opts *bind.TransactOpts, tokenPriceFeedUpdates []PriceRegistryTokenPriceFeedUpdate) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "updateTokenPriceFeeds", tokenPriceFeedUpdates)
}

func (_PriceRegistry *PriceRegistrySession) UpdateTokenPriceFeeds(tokenPriceFeedUpdates []PriceRegistryTokenPriceFeedUpdate) (*types.Transaction, error) {
	return _PriceRegistry.Contract.UpdateTokenPriceFeeds(&_PriceRegistry.TransactOpts, tokenPriceFeedUpdates)
}

func (_PriceRegistry *PriceRegistryTransactorSession) UpdateTokenPriceFeeds(tokenPriceFeedUpdates []PriceRegistryTokenPriceFeedUpdate) (*types.Transaction, error) {
	return _PriceRegistry.Contract.UpdateTokenPriceFeeds(&_PriceRegistry.TransactOpts, tokenPriceFeedUpdates)
}

type PriceRegistryAuthorizedCallerAddedIterator struct {
	Event *PriceRegistryAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryAuthorizedCallerAdded)
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
		it.Event = new(PriceRegistryAuthorizedCallerAdded)
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

func (it *PriceRegistryAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*PriceRegistryAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &PriceRegistryAuthorizedCallerAddedIterator{contract: _PriceRegistry.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *PriceRegistryAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryAuthorizedCallerAdded)
				if err := _PriceRegistry.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseAuthorizedCallerAdded(log types.Log) (*PriceRegistryAuthorizedCallerAdded, error) {
	event := new(PriceRegistryAuthorizedCallerAdded)
	if err := _PriceRegistry.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryAuthorizedCallerRemovedIterator struct {
	Event *PriceRegistryAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryAuthorizedCallerRemoved)
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
		it.Event = new(PriceRegistryAuthorizedCallerRemoved)
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

func (it *PriceRegistryAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*PriceRegistryAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &PriceRegistryAuthorizedCallerRemovedIterator{contract: _PriceRegistry.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *PriceRegistryAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryAuthorizedCallerRemoved)
				if err := _PriceRegistry.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*PriceRegistryAuthorizedCallerRemoved, error) {
	event := new(PriceRegistryAuthorizedCallerRemoved)
	if err := _PriceRegistry.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryDestChainAddedIterator struct {
	Event *PriceRegistryDestChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryDestChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryDestChainAdded)
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
		it.Event = new(PriceRegistryDestChainAdded)
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

func (it *PriceRegistryDestChainAddedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryDestChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryDestChainAdded struct {
	DestChainSelector uint64
	DynamicConfig     PriceRegistryDestChainDynamicConfig
	Raw               types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*PriceRegistryDestChainAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryDestChainAddedIterator{contract: _PriceRegistry.contract, event: "DestChainAdded", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *PriceRegistryDestChainAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryDestChainAdded)
				if err := _PriceRegistry.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseDestChainAdded(log types.Log) (*PriceRegistryDestChainAdded, error) {
	event := new(PriceRegistryDestChainAdded)
	if err := _PriceRegistry.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryDestChainDynamicConfigUpdatedIterator struct {
	Event *PriceRegistryDestChainDynamicConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryDestChainDynamicConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryDestChainDynamicConfigUpdated)
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
		it.Event = new(PriceRegistryDestChainDynamicConfigUpdated)
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

func (it *PriceRegistryDestChainDynamicConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryDestChainDynamicConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryDestChainDynamicConfigUpdated struct {
	DestChainSelector uint64
	DynamicConfig     PriceRegistryDestChainDynamicConfig
	Raw               types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterDestChainDynamicConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*PriceRegistryDestChainDynamicConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "DestChainDynamicConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryDestChainDynamicConfigUpdatedIterator{contract: _PriceRegistry.contract, event: "DestChainDynamicConfigUpdated", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchDestChainDynamicConfigUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryDestChainDynamicConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "DestChainDynamicConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryDestChainDynamicConfigUpdated)
				if err := _PriceRegistry.contract.UnpackLog(event, "DestChainDynamicConfigUpdated", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseDestChainDynamicConfigUpdated(log types.Log) (*PriceRegistryDestChainDynamicConfigUpdated, error) {
	event := new(PriceRegistryDestChainDynamicConfigUpdated)
	if err := _PriceRegistry.contract.UnpackLog(event, "DestChainDynamicConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryFeeTokenAddedIterator struct {
	Event *PriceRegistryFeeTokenAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryFeeTokenAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryFeeTokenAdded)
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
		it.Event = new(PriceRegistryFeeTokenAdded)
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

func (it *PriceRegistryFeeTokenAddedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryFeeTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryFeeTokenAdded struct {
	FeeToken common.Address
	Raw      types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterFeeTokenAdded(opts *bind.FilterOpts, feeToken []common.Address) (*PriceRegistryFeeTokenAddedIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "FeeTokenAdded", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryFeeTokenAddedIterator{contract: _PriceRegistry.contract, event: "FeeTokenAdded", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchFeeTokenAdded(opts *bind.WatchOpts, sink chan<- *PriceRegistryFeeTokenAdded, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "FeeTokenAdded", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryFeeTokenAdded)
				if err := _PriceRegistry.contract.UnpackLog(event, "FeeTokenAdded", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseFeeTokenAdded(log types.Log) (*PriceRegistryFeeTokenAdded, error) {
	event := new(PriceRegistryFeeTokenAdded)
	if err := _PriceRegistry.contract.UnpackLog(event, "FeeTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryFeeTokenRemovedIterator struct {
	Event *PriceRegistryFeeTokenRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryFeeTokenRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryFeeTokenRemoved)
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
		it.Event = new(PriceRegistryFeeTokenRemoved)
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

func (it *PriceRegistryFeeTokenRemovedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryFeeTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryFeeTokenRemoved struct {
	FeeToken common.Address
	Raw      types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterFeeTokenRemoved(opts *bind.FilterOpts, feeToken []common.Address) (*PriceRegistryFeeTokenRemovedIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "FeeTokenRemoved", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryFeeTokenRemovedIterator{contract: _PriceRegistry.contract, event: "FeeTokenRemoved", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchFeeTokenRemoved(opts *bind.WatchOpts, sink chan<- *PriceRegistryFeeTokenRemoved, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "FeeTokenRemoved", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryFeeTokenRemoved)
				if err := _PriceRegistry.contract.UnpackLog(event, "FeeTokenRemoved", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseFeeTokenRemoved(log types.Log) (*PriceRegistryFeeTokenRemoved, error) {
	event := new(PriceRegistryFeeTokenRemoved)
	if err := _PriceRegistry.contract.UnpackLog(event, "FeeTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryOwnershipTransferRequestedIterator struct {
	Event *PriceRegistryOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryOwnershipTransferRequested)
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
		it.Event = new(PriceRegistryOwnershipTransferRequested)
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

func (it *PriceRegistryOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PriceRegistryOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryOwnershipTransferRequestedIterator{contract: _PriceRegistry.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *PriceRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryOwnershipTransferRequested)
				if err := _PriceRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseOwnershipTransferRequested(log types.Log) (*PriceRegistryOwnershipTransferRequested, error) {
	event := new(PriceRegistryOwnershipTransferRequested)
	if err := _PriceRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryOwnershipTransferredIterator struct {
	Event *PriceRegistryOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryOwnershipTransferred)
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
		it.Event = new(PriceRegistryOwnershipTransferred)
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

func (it *PriceRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PriceRegistryOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryOwnershipTransferredIterator{contract: _PriceRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryOwnershipTransferred)
				if err := _PriceRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*PriceRegistryOwnershipTransferred, error) {
	event := new(PriceRegistryOwnershipTransferred)
	if err := _PriceRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryPremiumMultiplierWeiPerEthUpdatedIterator struct {
	Event *PriceRegistryPremiumMultiplierWeiPerEthUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryPremiumMultiplierWeiPerEthUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryPremiumMultiplierWeiPerEthUpdated)
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
		it.Event = new(PriceRegistryPremiumMultiplierWeiPerEthUpdated)
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

func (it *PriceRegistryPremiumMultiplierWeiPerEthUpdatedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryPremiumMultiplierWeiPerEthUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryPremiumMultiplierWeiPerEthUpdated struct {
	Token                      common.Address
	PremiumMultiplierWeiPerEth uint64
	Raw                        types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterPremiumMultiplierWeiPerEthUpdated(opts *bind.FilterOpts, token []common.Address) (*PriceRegistryPremiumMultiplierWeiPerEthUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "PremiumMultiplierWeiPerEthUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryPremiumMultiplierWeiPerEthUpdatedIterator{contract: _PriceRegistry.contract, event: "PremiumMultiplierWeiPerEthUpdated", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchPremiumMultiplierWeiPerEthUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryPremiumMultiplierWeiPerEthUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "PremiumMultiplierWeiPerEthUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryPremiumMultiplierWeiPerEthUpdated)
				if err := _PriceRegistry.contract.UnpackLog(event, "PremiumMultiplierWeiPerEthUpdated", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParsePremiumMultiplierWeiPerEthUpdated(log types.Log) (*PriceRegistryPremiumMultiplierWeiPerEthUpdated, error) {
	event := new(PriceRegistryPremiumMultiplierWeiPerEthUpdated)
	if err := _PriceRegistry.contract.UnpackLog(event, "PremiumMultiplierWeiPerEthUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryPriceFeedPerTokenUpdatedIterator struct {
	Event *PriceRegistryPriceFeedPerTokenUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryPriceFeedPerTokenUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryPriceFeedPerTokenUpdated)
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
		it.Event = new(PriceRegistryPriceFeedPerTokenUpdated)
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

func (it *PriceRegistryPriceFeedPerTokenUpdatedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryPriceFeedPerTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryPriceFeedPerTokenUpdated struct {
	Token           common.Address
	PriceFeedConfig IPriceRegistryTokenPriceFeedConfig
	Raw             types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterPriceFeedPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*PriceRegistryPriceFeedPerTokenUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "PriceFeedPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryPriceFeedPerTokenUpdatedIterator{contract: _PriceRegistry.contract, event: "PriceFeedPerTokenUpdated", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchPriceFeedPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryPriceFeedPerTokenUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "PriceFeedPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryPriceFeedPerTokenUpdated)
				if err := _PriceRegistry.contract.UnpackLog(event, "PriceFeedPerTokenUpdated", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParsePriceFeedPerTokenUpdated(log types.Log) (*PriceRegistryPriceFeedPerTokenUpdated, error) {
	event := new(PriceRegistryPriceFeedPerTokenUpdated)
	if err := _PriceRegistry.contract.UnpackLog(event, "PriceFeedPerTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryPriceUpdaterRemovedIterator struct {
	Event *PriceRegistryPriceUpdaterRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryPriceUpdaterRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryPriceUpdaterRemoved)
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
		it.Event = new(PriceRegistryPriceUpdaterRemoved)
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

func (it *PriceRegistryPriceUpdaterRemovedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryPriceUpdaterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryPriceUpdaterRemoved struct {
	PriceUpdater common.Address
	Raw          types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterPriceUpdaterRemoved(opts *bind.FilterOpts, priceUpdater []common.Address) (*PriceRegistryPriceUpdaterRemovedIterator, error) {

	var priceUpdaterRule []interface{}
	for _, priceUpdaterItem := range priceUpdater {
		priceUpdaterRule = append(priceUpdaterRule, priceUpdaterItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "PriceUpdaterRemoved", priceUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryPriceUpdaterRemovedIterator{contract: _PriceRegistry.contract, event: "PriceUpdaterRemoved", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchPriceUpdaterRemoved(opts *bind.WatchOpts, sink chan<- *PriceRegistryPriceUpdaterRemoved, priceUpdater []common.Address) (event.Subscription, error) {

	var priceUpdaterRule []interface{}
	for _, priceUpdaterItem := range priceUpdater {
		priceUpdaterRule = append(priceUpdaterRule, priceUpdaterItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "PriceUpdaterRemoved", priceUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryPriceUpdaterRemoved)
				if err := _PriceRegistry.contract.UnpackLog(event, "PriceUpdaterRemoved", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParsePriceUpdaterRemoved(log types.Log) (*PriceRegistryPriceUpdaterRemoved, error) {
	event := new(PriceRegistryPriceUpdaterRemoved)
	if err := _PriceRegistry.contract.UnpackLog(event, "PriceUpdaterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryPriceUpdaterSetIterator struct {
	Event *PriceRegistryPriceUpdaterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryPriceUpdaterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryPriceUpdaterSet)
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
		it.Event = new(PriceRegistryPriceUpdaterSet)
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

func (it *PriceRegistryPriceUpdaterSetIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryPriceUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryPriceUpdaterSet struct {
	PriceUpdater common.Address
	Raw          types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterPriceUpdaterSet(opts *bind.FilterOpts, priceUpdater []common.Address) (*PriceRegistryPriceUpdaterSetIterator, error) {

	var priceUpdaterRule []interface{}
	for _, priceUpdaterItem := range priceUpdater {
		priceUpdaterRule = append(priceUpdaterRule, priceUpdaterItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "PriceUpdaterSet", priceUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryPriceUpdaterSetIterator{contract: _PriceRegistry.contract, event: "PriceUpdaterSet", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchPriceUpdaterSet(opts *bind.WatchOpts, sink chan<- *PriceRegistryPriceUpdaterSet, priceUpdater []common.Address) (event.Subscription, error) {

	var priceUpdaterRule []interface{}
	for _, priceUpdaterItem := range priceUpdater {
		priceUpdaterRule = append(priceUpdaterRule, priceUpdaterItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "PriceUpdaterSet", priceUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryPriceUpdaterSet)
				if err := _PriceRegistry.contract.UnpackLog(event, "PriceUpdaterSet", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParsePriceUpdaterSet(log types.Log) (*PriceRegistryPriceUpdaterSet, error) {
	event := new(PriceRegistryPriceUpdaterSet)
	if err := _PriceRegistry.contract.UnpackLog(event, "PriceUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryTokenTransferFeeConfigDeletedIterator struct {
	Event *PriceRegistryTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryTokenTransferFeeConfigDeleted)
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
		it.Event = new(PriceRegistryTokenTransferFeeConfigDeleted)
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

func (it *PriceRegistryTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Token             common.Address
	Raw               types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*PriceRegistryTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryTokenTransferFeeConfigDeletedIterator{contract: _PriceRegistry.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *PriceRegistryTokenTransferFeeConfigDeleted, destChainSelector []uint64, token []common.Address) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryTokenTransferFeeConfigDeleted)
				if err := _PriceRegistry.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*PriceRegistryTokenTransferFeeConfigDeleted, error) {
	event := new(PriceRegistryTokenTransferFeeConfigDeleted)
	if err := _PriceRegistry.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryTokenTransferFeeConfigUpdatedIterator struct {
	Event *PriceRegistryTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryTokenTransferFeeConfigUpdated)
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
		it.Event = new(PriceRegistryTokenTransferFeeConfigUpdated)
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

func (it *PriceRegistryTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	Token                  common.Address
	TokenTransferFeeConfig PriceRegistryTokenTransferFeeConfig
	Raw                    types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*PriceRegistryTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryTokenTransferFeeConfigUpdatedIterator{contract: _PriceRegistry.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryTokenTransferFeeConfigUpdated, destChainSelector []uint64, token []common.Address) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryTokenTransferFeeConfigUpdated)
				if err := _PriceRegistry.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*PriceRegistryTokenTransferFeeConfigUpdated, error) {
	event := new(PriceRegistryTokenTransferFeeConfigUpdated)
	if err := _PriceRegistry.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryUsdPerTokenUpdatedIterator struct {
	Event *PriceRegistryUsdPerTokenUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryUsdPerTokenUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryUsdPerTokenUpdated)
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
		it.Event = new(PriceRegistryUsdPerTokenUpdated)
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

func (it *PriceRegistryUsdPerTokenUpdatedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryUsdPerTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryUsdPerTokenUpdated struct {
	Token     common.Address
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterUsdPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*PriceRegistryUsdPerTokenUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "UsdPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryUsdPerTokenUpdatedIterator{contract: _PriceRegistry.contract, event: "UsdPerTokenUpdated", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchUsdPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryUsdPerTokenUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "UsdPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryUsdPerTokenUpdated)
				if err := _PriceRegistry.contract.UnpackLog(event, "UsdPerTokenUpdated", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseUsdPerTokenUpdated(log types.Log) (*PriceRegistryUsdPerTokenUpdated, error) {
	event := new(PriceRegistryUsdPerTokenUpdated)
	if err := _PriceRegistry.contract.UnpackLog(event, "UsdPerTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PriceRegistryUsdPerUnitGasUpdatedIterator struct {
	Event *PriceRegistryUsdPerUnitGasUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryUsdPerUnitGasUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryUsdPerUnitGasUpdated)
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
		it.Event = new(PriceRegistryUsdPerUnitGasUpdated)
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

func (it *PriceRegistryUsdPerUnitGasUpdatedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryUsdPerUnitGasUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryUsdPerUnitGasUpdated struct {
	DestChain uint64
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterUsdPerUnitGasUpdated(opts *bind.FilterOpts, destChain []uint64) (*PriceRegistryUsdPerUnitGasUpdatedIterator, error) {

	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "UsdPerUnitGasUpdated", destChainRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryUsdPerUnitGasUpdatedIterator{contract: _PriceRegistry.contract, event: "UsdPerUnitGasUpdated", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchUsdPerUnitGasUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryUsdPerUnitGasUpdated, destChain []uint64) (event.Subscription, error) {

	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "UsdPerUnitGasUpdated", destChainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryUsdPerUnitGasUpdated)
				if err := _PriceRegistry.contract.UnpackLog(event, "UsdPerUnitGasUpdated", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseUsdPerUnitGasUpdated(log types.Log) (*PriceRegistryUsdPerUnitGasUpdated, error) {
	event := new(PriceRegistryUsdPerUnitGasUpdated)
	if err := _PriceRegistry.contract.UnpackLog(event, "UsdPerUnitGasUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetTokenAndGasPrices struct {
	TokenPrice    *big.Int
	GasPriceValue *big.Int
}
type GetValidatedRampMessageParams struct {
	MsgFeeJuels           *big.Int
	IsOutOfOrderExecution bool
	ConvertedExtraArgs    []byte
}

func (_PriceRegistry *PriceRegistry) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _PriceRegistry.abi.Events["AuthorizedCallerAdded"].ID:
		return _PriceRegistry.ParseAuthorizedCallerAdded(log)
	case _PriceRegistry.abi.Events["AuthorizedCallerRemoved"].ID:
		return _PriceRegistry.ParseAuthorizedCallerRemoved(log)
	case _PriceRegistry.abi.Events["DestChainAdded"].ID:
		return _PriceRegistry.ParseDestChainAdded(log)
	case _PriceRegistry.abi.Events["DestChainDynamicConfigUpdated"].ID:
		return _PriceRegistry.ParseDestChainDynamicConfigUpdated(log)
	case _PriceRegistry.abi.Events["FeeTokenAdded"].ID:
		return _PriceRegistry.ParseFeeTokenAdded(log)
	case _PriceRegistry.abi.Events["FeeTokenRemoved"].ID:
		return _PriceRegistry.ParseFeeTokenRemoved(log)
	case _PriceRegistry.abi.Events["OwnershipTransferRequested"].ID:
		return _PriceRegistry.ParseOwnershipTransferRequested(log)
	case _PriceRegistry.abi.Events["OwnershipTransferred"].ID:
		return _PriceRegistry.ParseOwnershipTransferred(log)
	case _PriceRegistry.abi.Events["PremiumMultiplierWeiPerEthUpdated"].ID:
		return _PriceRegistry.ParsePremiumMultiplierWeiPerEthUpdated(log)
	case _PriceRegistry.abi.Events["PriceFeedPerTokenUpdated"].ID:
		return _PriceRegistry.ParsePriceFeedPerTokenUpdated(log)
	case _PriceRegistry.abi.Events["PriceUpdaterRemoved"].ID:
		return _PriceRegistry.ParsePriceUpdaterRemoved(log)
	case _PriceRegistry.abi.Events["PriceUpdaterSet"].ID:
		return _PriceRegistry.ParsePriceUpdaterSet(log)
	case _PriceRegistry.abi.Events["TokenTransferFeeConfigDeleted"].ID:
		return _PriceRegistry.ParseTokenTransferFeeConfigDeleted(log)
	case _PriceRegistry.abi.Events["TokenTransferFeeConfigUpdated"].ID:
		return _PriceRegistry.ParseTokenTransferFeeConfigUpdated(log)
	case _PriceRegistry.abi.Events["UsdPerTokenUpdated"].ID:
		return _PriceRegistry.ParseUsdPerTokenUpdated(log)
	case _PriceRegistry.abi.Events["UsdPerUnitGasUpdated"].ID:
		return _PriceRegistry.ParseUsdPerUnitGasUpdated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (PriceRegistryAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (PriceRegistryAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (PriceRegistryDestChainAdded) Topic() common.Hash {
	return common.HexToHash("0xd883f18d08e8ea5c2af22d13fb6cab86cc057d795d2735f83a8ff137d39d1b8f")
}

func (PriceRegistryDestChainDynamicConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xd42d3a670a4f1ab5d8703efa22bb17041446365cfcfa33980ced685a080cf7cb")
}

func (PriceRegistryFeeTokenAdded) Topic() common.Hash {
	return common.HexToHash("0xdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23")
}

func (PriceRegistryFeeTokenRemoved) Topic() common.Hash {
	return common.HexToHash("0x1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91")
}

func (PriceRegistryOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (PriceRegistryOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (PriceRegistryPremiumMultiplierWeiPerEthUpdated) Topic() common.Hash {
	return common.HexToHash("0xbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d")
}

func (PriceRegistryPriceFeedPerTokenUpdated) Topic() common.Hash {
	return common.HexToHash("0x08a5f7f5bb38a81d8e43aca13ecd76431dbf8816ae4699affff7b00b2fc1c464")
}

func (PriceRegistryPriceUpdaterRemoved) Topic() common.Hash {
	return common.HexToHash("0xff7dbb85c77ca68ca1f894d6498570e3d5095cd19466f07ee8d222b337e4068c")
}

func (PriceRegistryPriceUpdaterSet) Topic() common.Hash {
	return common.HexToHash("0x34a02290b7920078c19f58e94b78c77eb9cc10195b20676e19bd3b82085893b8")
}

func (PriceRegistryTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b")
}

func (PriceRegistryTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b5")
}

func (PriceRegistryUsdPerTokenUpdated) Topic() common.Hash {
	return common.HexToHash("0x52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a")
}

func (PriceRegistryUsdPerUnitGasUpdated) Topic() common.Hash {
	return common.HexToHash("0xdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e")
}

func (_PriceRegistry *PriceRegistry) Address() common.Address {
	return _PriceRegistry.address
}

type PriceRegistryInterface interface {
	ConvertTokenAmount(opts *bind.CallOpts, fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error)

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetDestChainDynamicConfig(opts *bind.CallOpts, destChainSelector uint64) (PriceRegistryDestChainDynamicConfig, error)

	GetDestinationChainGasPrice(opts *bind.CallOpts, destChainSelector uint64) (InternalTimestampedPackedUint224, error)

	GetFeeTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetPremiumMultiplierWeiPerEth(opts *bind.CallOpts, token common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (PriceRegistryStaticConfig, error)

	GetTokenAndGasPrices(opts *bind.CallOpts, token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

		error)

	GetTokenPrice(opts *bind.CallOpts, token common.Address) (InternalTimestampedPackedUint224, error)

	GetTokenPriceFeedConfig(opts *bind.CallOpts, token common.Address) (IPriceRegistryTokenPriceFeedConfig, error)

	GetTokenPrices(opts *bind.CallOpts, tokens []common.Address) ([]InternalTimestampedPackedUint224, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (PriceRegistryTokenTransferFeeConfig, error)

	GetValidatedFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetValidatedRampMessageParams(opts *bind.CallOpts, message InternalEVM2AnyRampMessage, sourceTokenAmounts []ClientEVMTokenAmount) (GetValidatedRampMessageParams,

		error)

	GetValidatedTokenPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []PriceRegistryDestChainDynamicConfigArgs) (*types.Transaction, error)

	ApplyFeeTokensUpdates(opts *bind.TransactOpts, feeTokensToAdd []common.Address, feeTokensToRemove []common.Address) (*types.Transaction, error)

	ApplyPremiumMultiplierWeiPerEthUpdates(opts *bind.TransactOpts, premiumMultiplierWeiPerEthArgs []PriceRegistryPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []PriceRegistryTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []PriceRegistryTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdatePrices(opts *bind.TransactOpts, priceUpdates InternalPriceUpdates) (*types.Transaction, error)

	UpdateTokenPriceFeeds(opts *bind.TransactOpts, tokenPriceFeedUpdates []PriceRegistryTokenPriceFeedUpdate) (*types.Transaction, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*PriceRegistryAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *PriceRegistryAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*PriceRegistryAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*PriceRegistryAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *PriceRegistryAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*PriceRegistryAuthorizedCallerRemoved, error)

	FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*PriceRegistryDestChainAddedIterator, error)

	WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *PriceRegistryDestChainAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainAdded(log types.Log) (*PriceRegistryDestChainAdded, error)

	FilterDestChainDynamicConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*PriceRegistryDestChainDynamicConfigUpdatedIterator, error)

	WatchDestChainDynamicConfigUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryDestChainDynamicConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainDynamicConfigUpdated(log types.Log) (*PriceRegistryDestChainDynamicConfigUpdated, error)

	FilterFeeTokenAdded(opts *bind.FilterOpts, feeToken []common.Address) (*PriceRegistryFeeTokenAddedIterator, error)

	WatchFeeTokenAdded(opts *bind.WatchOpts, sink chan<- *PriceRegistryFeeTokenAdded, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenAdded(log types.Log) (*PriceRegistryFeeTokenAdded, error)

	FilterFeeTokenRemoved(opts *bind.FilterOpts, feeToken []common.Address) (*PriceRegistryFeeTokenRemovedIterator, error)

	WatchFeeTokenRemoved(opts *bind.WatchOpts, sink chan<- *PriceRegistryFeeTokenRemoved, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenRemoved(log types.Log) (*PriceRegistryFeeTokenRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PriceRegistryOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *PriceRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*PriceRegistryOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PriceRegistryOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PriceRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*PriceRegistryOwnershipTransferred, error)

	FilterPremiumMultiplierWeiPerEthUpdated(opts *bind.FilterOpts, token []common.Address) (*PriceRegistryPremiumMultiplierWeiPerEthUpdatedIterator, error)

	WatchPremiumMultiplierWeiPerEthUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryPremiumMultiplierWeiPerEthUpdated, token []common.Address) (event.Subscription, error)

	ParsePremiumMultiplierWeiPerEthUpdated(log types.Log) (*PriceRegistryPremiumMultiplierWeiPerEthUpdated, error)

	FilterPriceFeedPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*PriceRegistryPriceFeedPerTokenUpdatedIterator, error)

	WatchPriceFeedPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryPriceFeedPerTokenUpdated, token []common.Address) (event.Subscription, error)

	ParsePriceFeedPerTokenUpdated(log types.Log) (*PriceRegistryPriceFeedPerTokenUpdated, error)

	FilterPriceUpdaterRemoved(opts *bind.FilterOpts, priceUpdater []common.Address) (*PriceRegistryPriceUpdaterRemovedIterator, error)

	WatchPriceUpdaterRemoved(opts *bind.WatchOpts, sink chan<- *PriceRegistryPriceUpdaterRemoved, priceUpdater []common.Address) (event.Subscription, error)

	ParsePriceUpdaterRemoved(log types.Log) (*PriceRegistryPriceUpdaterRemoved, error)

	FilterPriceUpdaterSet(opts *bind.FilterOpts, priceUpdater []common.Address) (*PriceRegistryPriceUpdaterSetIterator, error)

	WatchPriceUpdaterSet(opts *bind.WatchOpts, sink chan<- *PriceRegistryPriceUpdaterSet, priceUpdater []common.Address) (event.Subscription, error)

	ParsePriceUpdaterSet(log types.Log) (*PriceRegistryPriceUpdaterSet, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*PriceRegistryTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *PriceRegistryTokenTransferFeeConfigDeleted, destChainSelector []uint64, token []common.Address) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*PriceRegistryTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*PriceRegistryTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryTokenTransferFeeConfigUpdated, destChainSelector []uint64, token []common.Address) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*PriceRegistryTokenTransferFeeConfigUpdated, error)

	FilterUsdPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*PriceRegistryUsdPerTokenUpdatedIterator, error)

	WatchUsdPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryUsdPerTokenUpdated, token []common.Address) (event.Subscription, error)

	ParseUsdPerTokenUpdated(log types.Log) (*PriceRegistryUsdPerTokenUpdated, error)

	FilterUsdPerUnitGasUpdated(opts *bind.FilterOpts, destChain []uint64) (*PriceRegistryUsdPerUnitGasUpdatedIterator, error)

	WatchUsdPerUnitGasUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryUsdPerUnitGasUpdated, destChain []uint64) (event.Subscription, error)

	ParseUsdPerUnitGasUpdated(log types.Log) (*PriceRegistryUsdPerUnitGasUpdated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
