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

type InternalGasPriceUpdate struct {
	DestChainSelector uint64
	UsdPerUnitGas     *big.Int
}

type InternalPriceUpdates struct {
	TokenPriceUpdates []InternalTokenPriceUpdate
	GasPriceUpdates   []InternalGasPriceUpdate
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

type KeystoneFeedsPermissionHandlerPermission struct {
	Forwarder     common.Address
	WorkflowName  [10]byte
	ReportName    [2]byte
	WorkflowOwner common.Address
	IsAllowed     bool
}

type PriceRegistryDestChainConfig struct {
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
	DefaultTxGasLimit                 uint32
	GasMultiplierWeiPerEth            uint64
	NetworkFeeUSDCents                uint32
	EnforceOutOfOrder                 bool
	ChainFamilySelector               [4]byte
}

type PriceRegistryDestChainConfigArgs struct {
	DestChainSelector uint64
	DestChainConfig   PriceRegistryDestChainConfig
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessThreshold\",\"type\":\"uint32\"}],\"internalType\":\"structPriceRegistry.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"priceUpdaters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIPriceRegistry.TokenPriceFeedConfig\",\"name\":\"feedConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.TokenPriceFeedUpdate[]\",\"name\":\"tokenPriceFeeds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structPriceRegistry.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"internalType\":\"structPriceRegistry.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chain\",\"type\":\"uint64\"}],\"name\":\"ChainNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataFeedValueOutOfUint224Range\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"DestinationChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtraArgOutOfOrderExecutionMustBeTrue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"}],\"name\":\"InvalidDestBytesOverhead\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidDestChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStaticConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgFeeJuels\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint256\"}],\"name\":\"MessageFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SourceTokenDataTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePassed\",\"type\":\"uint256\"}],\"name\":\"StaleGasPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"bytes10\",\"name\":\"workflowName\",\"type\":\"bytes10\"},{\"internalType\":\"bytes2\",\"name\":\"reportName\",\"type\":\"bytes2\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"UnauthorizedCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedCallerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedCallerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"indexed\":false,\"internalType\":\"structPriceRegistry.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"indexed\":false,\"internalType\":\"structPriceRegistry.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"FeeTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"FeeTokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"name\":\"PremiumMultiplierWeiPerEthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIPriceRegistry.TokenPriceFeedConfig\",\"name\":\"priceFeedConfig\",\"type\":\"tuple\"}],\"name\":\"PriceFeedPerTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceUpdater\",\"type\":\"address\"}],\"name\":\"PriceUpdaterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"priceUpdater\",\"type\":\"address\"}],\"name\":\"PriceUpdaterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reportId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bytes10\",\"name\":\"workflowName\",\"type\":\"bytes10\"},{\"internalType\":\"bytes2\",\"name\":\"reportName\",\"type\":\"bytes2\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structKeystoneFeedsPermissionHandler.Permission\",\"name\":\"permission\",\"type\":\"tuple\"}],\"name\":\"ReportPermissionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenTransferFeeConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structPriceRegistry.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"name\":\"TokenTransferFeeConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UsdPerTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChain\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UsdPerUnitGasUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"addedCallers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removedCallers\",\"type\":\"address[]\"}],\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\"}],\"name\":\"applyAuthorizedCallerUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"internalType\":\"structPriceRegistry.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyDestChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"feeTokensToAdd\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"feeTokensToRemove\",\"type\":\"address[]\"}],\"name\":\"applyFeeTokensUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structPriceRegistry.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyPremiumMultiplierWeiPerEthUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigRemoveArgs[]\",\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\"}],\"name\":\"applyTokenTransferFeeConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"}],\"name\":\"convertTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAuthorizedCallers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"internalType\":\"structPriceRegistry.DestChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestinationChainGasPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.TimestampedPackedUint224\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPremiumMultiplierWeiPerEth\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessThreshold\",\"type\":\"uint32\"}],\"internalType\":\"structPriceRegistry.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getTokenAndGasPrices\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"tokenPrice\",\"type\":\"uint224\"},{\"internalType\":\"uint224\",\"name\":\"gasPriceValue\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.TimestampedPackedUint224\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPriceFeedConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIPriceRegistry.TokenPriceFeedConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getTokenPrices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.TimestampedPackedUint224[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenTransferFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getValidatedFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getValidatedTokenPrice\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"onReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"name\":\"processMessageArgs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"msgFeeJuels\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOutOfOrderExecution\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"convertedExtraArgs\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bytes10\",\"name\":\"workflowName\",\"type\":\"bytes10\"},{\"internalType\":\"bytes2\",\"name\":\"reportName\",\"type\":\"bytes2\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"internalType\":\"structKeystoneFeedsPermissionHandler.Permission[]\",\"name\":\"permissions\",\"type\":\"tuple[]\"}],\"name\":\"setReportPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint224\",\"name\":\"usdPerToken\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint224\",\"name\":\"usdPerUnitGas\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.GasPriceUpdate[]\",\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.PriceUpdates\",\"name\":\"priceUpdates\",\"type\":\"tuple\"}],\"name\":\"updatePrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIPriceRegistry.TokenPriceFeedConfig\",\"name\":\"feedConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.TokenPriceFeedUpdate[]\",\"name\":\"tokenPriceFeedUpdates\",\"type\":\"tuple[]\"}],\"name\":\"updateTokenPriceFeeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.RampTokenAmount[]\",\"name\":\"rampTokenAmounts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\"}],\"name\":\"validatePoolReturnData\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200723d3803806200723d83398101604081905262000034916200180a565b8533806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000207565b5050604080518082018252838152815160008152602080820190935291810191909152620000ee9150620002b2565b5060208701516001600160a01b0316158062000112575086516001600160601b0316155b80620001265750604087015163ffffffff16155b15620001455760405163d794ef9560e01b815260040160405180910390fd5b6020878101516001600160a01b031660a05287516001600160601b031660805260408089015163ffffffff1660c05280516000815291820190526200018c90869062000401565b620001978462000549565b620001a2816200061a565b620001ad8262000a99565b60408051600080825260208201909252620001fa91859190620001f3565b6040805180820190915260008082526020820152815260200190600190039081620001cb5790505b5062000b65565b5050505050505062001ac8565b336001600160a01b03821603620002615760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b602081015160005b815181101562000342576000828281518110620002db57620002db62001929565b60209081029190910101519050620002f560028262000e98565b1562000338576040516001600160a01b03821681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda775809060200160405180910390a15b50600101620002ba565b50815160005b8151811015620003fb57600082828151811062000369576200036962001929565b6020026020010151905060006001600160a01b0316816001600160a01b031603620003a7576040516342bcdf7f60e11b815260040160405180910390fd5b620003b460028262000eb8565b506040516001600160a01b03821681527feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef9060200160405180910390a15060010162000348565b50505050565b60005b8251811015620004a2576200044083828151811062000427576200042762001929565b6020026020010151600d62000eb860201b90919060201c565b1562000499578281815181106200045b576200045b62001929565b60200260200101516001600160a01b03167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba2360405160405180910390a25b60010162000404565b5060005b81518110156200054457620004e2828281518110620004c957620004c962001929565b6020026020010151600d62000e9860201b90919060201c565b156200053b57818181518110620004fd57620004fd62001929565b60200260200101516001600160a01b03167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f9160405160405180910390a25b600101620004a6565b505050565b60005b8151811015620006165760008282815181106200056d576200056d62001929565b6020908102919091018101518051818301516001600160a01b0380831660008181526007875260409081902084518154868a018051929096166001600160a81b03199091168117600160a01b60ff9384160217909255825191825293519093169683019690965293955091939092917f08a5f7f5bb38a81d8e43aca13ecd76431dbf8816ae4699affff7b00b2fc1c464910160405180910390a25050508060010190506200054c565b5050565b60005b8151811015620006165760008282815181106200063e576200063e62001929565b6020026020010151905060008383815181106200065f576200065f62001929565b6020026020010151600001519050600082602001519050816001600160401b03166000148062000698575061018081015163ffffffff16155b80620006ba57506102008101516001600160e01b031916630a04b54b60e21b14155b80620006d15750602081610160015163ffffffff16105b80620006f15750806060015163ffffffff1681610180015163ffffffff16115b156200071c5760405163c35aa79d60e01b81526001600160401b038316600482015260240162000083565b6001600160401b038216600090815260096020526040812060010154600160a81b900460e01b6001600160e01b03191690036200079c57816001600160401b03167fa937382a486d993de71c220bc8b559242deb4e286a353fa732330b4aa7d13577826040516200078e91906200193f565b60405180910390a2620007e0565b816001600160401b03167fa7b607fc10d28a1caf39ab7d27f4c94945db708a576d572781a455c5894fad9382604051620007d791906200193f565b60405180910390a25b8060096000846001600160401b03166001600160401b0316815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a81548163ffffffff021916908363ffffffff1602179055506101a08201518160010160086101000a8154816001600160401b0302191690836001600160401b031602179055506101c08201518160010160106101000a81548163ffffffff021916908363ffffffff1602179055506101e08201518160010160146101000a81548160ff0219169083151502179055506102008201518160010160156101000a81548163ffffffff021916908360e01c02179055509050505050508060010190506200061d565b60005b81518110156200061657600082828151811062000abd5762000abd62001929565b6020026020010151600001519050600083838151811062000ae25762000ae262001929565b6020908102919091018101518101516001600160a01b03841660008181526008845260409081902080546001600160401b0319166001600160401b0385169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a2505060010162000a9c565b60005b825181101562000dd257600083828151811062000b895762000b8962001929565b6020026020010151905060008160000151905060005b82602001515181101562000dc35760008360200151828151811062000bc85762000bc862001929565b602002602001015160200151905060008460200151838151811062000bf15762000bf162001929565b60200260200101516000015190506020826080015163ffffffff16101562000c4a5760808201516040516312766e0160e11b81526001600160a01b038316600482015263ffffffff909116602482015260440162000083565b6001600160401b0384166000818152600a602090815260408083206001600160a01b0386168085529083529281902086518154938801518389015160608a015160808b015160a08c01511515600160901b0260ff60901b1963ffffffff928316600160701b021664ffffffffff60701b199383166a01000000000000000000000263ffffffff60501b1961ffff90961668010000000000000000029590951665ffffffffffff60401b19968416640100000000026001600160401b0319909b16939097169290921798909817939093169390931717919091161792909217909155519091907f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b59062000db0908690600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60405180910390a3505060010162000b9f565b50505080600101905062000b68565b5060005b81518110156200054457600082828151811062000df75762000df762001929565b6020026020010151600001519050600083838151811062000e1c5762000e1c62001929565b6020908102919091018101518101516001600160401b0384166000818152600a845260408082206001600160a01b038516808452955280822080546001600160981b03191690555192945090917f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b9190a3505060010162000dd6565b600062000eaf836001600160a01b03841662000ecf565b90505b92915050565b600062000eaf836001600160a01b03841662000fd3565b6000818152600183016020526040812054801562000fc857600062000ef660018362001a90565b855490915060009062000f0c9060019062001a90565b905081811462000f7857600086600001828154811062000f305762000f3062001929565b906000526020600020015490508087600001848154811062000f565762000f5662001929565b6000918252602080832090910192909255918252600188019052604090208390555b855486908062000f8c5762000f8c62001ab2565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000eb2565b600091505062000eb2565b60008181526001830160205260408120546200101c5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000eb2565b50600062000eb2565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171562001060576200106062001025565b60405290565b60405160c081016001600160401b038111828210171562001060576200106062001025565b60405161022081016001600160401b038111828210171562001060576200106062001025565b604051601f8201601f191681016001600160401b0381118282101715620010dc57620010dc62001025565b604052919050565b80516001600160a01b0381168114620010fc57600080fd5b919050565b805163ffffffff81168114620010fc57600080fd5b6000606082840312156200112957600080fd5b604051606081016001600160401b03811182821017156200114e576200114e62001025565b604052825190915081906001600160601b03811681146200116e57600080fd5b81526200117e60208401620010e4565b6020820152620011916040840162001101565b60408201525092915050565b60006001600160401b03821115620011b957620011b962001025565b5060051b60200190565b600082601f830112620011d557600080fd5b81516020620011ee620011e8836200119d565b620010b1565b8083825260208201915060208460051b8701019350868411156200121157600080fd5b602086015b8481101562001238576200122a81620010e4565b835291830191830162001216565b509695505050505050565b600082601f8301126200125557600080fd5b8151602062001268620011e8836200119d565b828152606092830285018201928282019190878511156200128857600080fd5b8387015b858110156200131b5780890382811215620012a75760008081fd5b620012b16200103b565b620012bc83620010e4565b8152604080601f1984011215620012d35760008081fd5b620012dd6200103b565b9250620012ec888501620010e4565b835283015160ff81168114620013025760008081fd5b828801528087019190915284529284019281016200128c565b5090979650505050505050565b80516001600160401b0381168114620010fc57600080fd5b805161ffff81168114620010fc57600080fd5b80518015158114620010fc57600080fd5b600082601f8301126200137657600080fd5b8151602062001389620011e8836200119d565b82815260059290921b84018101918181019086841115620013a957600080fd5b8286015b84811015620012385780516001600160401b0380821115620013ce57600080fd5b908801906040601f19838c038101821315620013e957600080fd5b620013f36200103b565b6200140089860162001328565b815282850151848111156200141457600080fd5b8086019550508c603f8601126200142a57600080fd5b8885015193506200143f620011e8856200119d565b84815260e09094028501830193898101908e8611156200145e57600080fd5b958401955b858710156200153757868f0360e08112156200147e57600080fd5b620014886200103b565b6200149389620010e4565b815260c08683011215620014a657600080fd5b620014b062001066565b9150620014bf8d8a0162001101565b8252620014ce878a0162001101565b8d830152620014e060608a0162001340565b87830152620014f260808a0162001101565b60608301526200150560a08a0162001101565b60808301526200151860c08a0162001353565b60a0830152808d0191909152825260e09690960195908a019062001463565b828b015250875250505092840192508301620013ad565b600082601f8301126200156057600080fd5b8151602062001573620011e8836200119d565b82815260069290921b840181019181810190868411156200159357600080fd5b8286015b84811015620012385760408189031215620015b25760008081fd5b620015bc6200103b565b620015c782620010e4565b8152620015d685830162001328565b8186015283529183019160400162001597565b80516001600160e01b031981168114620010fc57600080fd5b600082601f8301126200161457600080fd5b8151602062001627620011e8836200119d565b82815261024092830285018201928282019190878511156200164857600080fd5b8387015b858110156200131b5780890382811215620016675760008081fd5b620016716200103b565b6200167c8362001328565b815261022080601f1984011215620016945760008081fd5b6200169e6200108b565b9250620016ad88850162001353565b83526040620016be81860162001340565b898501526060620016d181870162001101565b8286015260809150620016e682870162001101565b9085015260a0620016f986820162001101565b8286015260c091506200170e82870162001340565b9085015260e06200172186820162001101565b8286015261010091506200173782870162001340565b908501526101206200174b86820162001340565b8286015261014091506200176182870162001340565b908501526101606200177586820162001101565b8286015261018091506200178b82870162001101565b908501526101a06200179f86820162001101565b828601526101c09150620017b582870162001328565b908501526101e0620017c986820162001101565b828601526102009150620017df82870162001353565b90850152620017f0858301620015e9565b90840152508087019190915284529284019281016200164c565b6000806000806000806000610120888a0312156200182757600080fd5b62001833898962001116565b60608901519097506001600160401b03808211156200185157600080fd5b6200185f8b838c01620011c3565b975060808a01519150808211156200187657600080fd5b620018848b838c01620011c3565b965060a08a01519150808211156200189b57600080fd5b620018a98b838c0162001243565b955060c08a0151915080821115620018c057600080fd5b620018ce8b838c0162001364565b945060e08a0151915080821115620018e557600080fd5b620018f38b838c016200154e565b93506101008a01519150808211156200190b57600080fd5b506200191a8a828b0162001602565b91505092959891949750929550565b634e487b7160e01b600052603260045260246000fd5b8151151581526102208101602083015162001960602084018261ffff169052565b50604083015162001979604084018263ffffffff169052565b50606083015162001992606084018263ffffffff169052565b506080830151620019ab608084018263ffffffff169052565b5060a0830151620019c260a084018261ffff169052565b5060c0830151620019db60c084018263ffffffff169052565b5060e0830151620019f260e084018261ffff169052565b506101008381015161ffff9081169184019190915261012080850151909116908301526101408084015163ffffffff9081169184019190915261016080850151821690840152610180808501518216908401526101a0808501516001600160401b0316908401526101c080850151909116908301526101e080840151151590830152610200928301516001600160e01b031916929091019190915290565b8181038181111562000eb257634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60805160a05160c05161571b62001b22600039600081816102ef01528181611db70152611e200152600081816102b301528181611332015261139201526000818161027f015281816113bb015261142b015261571b6000f3fe608060405234801561001057600080fd5b50600436106101ce5760003560e01c80637afac32211610104578063c4276bfc116100a2578063d8694ccd11610071578063d8694ccd14610a98578063f2fde38b14610aab578063f700042a14610abe578063ffdb4b3714610ad157600080fd5b8063c4276bfc14610a48578063cc88924c14610a6a578063cdc73d5114610a7d578063d02641a014610a8557600080fd5b80638da5cb5b116100de5780638da5cb5b1461094d57806391a2749a14610975578063a69c64c014610988578063bf78e03f1461099b57600080fd5b80637afac322146107b7578063805f2132146107ca57806382b49eb0146107dd57600080fd5b806341ed29e711610171578063514e8cff1161014b578063514e8cff146104615780636def4ce714610504578063770e2dc41461079c57806379ba5097146107af57600080fd5b806341ed29e7146103ee57806345ac924d146104015780634ab35b0b1461042157600080fd5b8063181f5a77116101ad578063181f5a77146103685780632451a627146103b15780633937306f146103c6578063407e1086146103db57600080fd5b806241e5be146101d3578063061877e3146101f957806306285c6914610252575b600080fd5b6101e66101e1366004613e2b565b610b19565b6040519081526020015b60405180910390f35b610239610207366004613e67565b73ffffffffffffffffffffffffffffffffffffffff1660009081526008602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff90911681526020016101f0565b61031c604080516060810182526000808252602082018190529181019190915260405180606001604052807f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000063ffffffff16815250905090565b6040805182516bffffffffffffffffffffffff16815260208084015173ffffffffffffffffffffffffffffffffffffffff16908201529181015163ffffffff16908201526060016101f0565b6103a46040518060400160405280601781526020017f5072696365526567697374727920312e362e302d64657600000000000000000081525081565b6040516101f09190613ee6565b6103b9610b87565b6040516101f09190613ef9565b6103d96103d4366004613f53565b610b98565b005b6103d96103e93660046140f5565b610e4d565b6103d96103fc366004614227565b610e61565b61041461040f3660046143a7565b610ea3565b6040516101f091906143e9565b61043461042f366004613e67565b610f6e565b6040517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90911681526020016101f0565b6104f761046f36600461447c565b60408051808201909152600080825260208201525067ffffffffffffffff166000908152600560209081526040918290208251808401909352547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811683527c0100000000000000000000000000000000000000000000000000000000900463ffffffff169082015290565b6040516101f09190614497565b61078f61051236600461447c565b6040805161022081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c081018290526101e081018290526102008101919091525067ffffffffffffffff908116600090815260096020908152604091829020825161022081018452815460ff8082161515835261ffff61010080840482169685019690965263ffffffff630100000084048116978501979097526701000000000000008304871660608501526b0100000000000000000000008304871660808501526f010000000000000000000000000000008304811660a0850152710100000000000000000000000000000000008304871660c08501527501000000000000000000000000000000000000000000808404821660e08087019190915277010000000000000000000000000000000000000000000000850483169786019790975279010000000000000000000000000000000000000000000000000084049091166101208501527b010000000000000000000000000000000000000000000000000000009092048616610140840152600190930154808616610160840152640100000000810486166101808401526801000000000000000081049096166101a083015270010000000000000000000000000000000086049094166101c082015274010000000000000000000000000000000000000000850490911615156101e08201527fffffffff0000000000000000000000000000000000000000000000000000000092909304901b1661020082015290565b6040516101f091906144d2565b6103d96107aa3660046146f6565b610f79565b6103d9610f8b565b6103d96107c5366004614a10565b61108d565b6103d96107d8366004614ab6565b61109f565b6108ed6107eb366004614b22565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a08101919091525067ffffffffffffffff919091166000908152600a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff94909416835292815290829020825160c081018452905463ffffffff8082168352640100000000820481169383019390935268010000000000000000810461ffff16938201939093526a01000000000000000000008304821660608201526e01000000000000000000000000000083049091166080820152720100000000000000000000000000000000000090910460ff16151560a082015290565b6040516101f09190600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101f0565b6103d9610983366004614b4c565b611309565b6103d9610996366004614bdd565b61131a565b610a146109a9366004613e67565b6040805180820182526000808252602091820181905273ffffffffffffffffffffffffffffffffffffffff93841681526007825282902082518084019093525492831682527401000000000000000000000000000000000000000090920460ff169181019190915290565b60408051825173ffffffffffffffffffffffffffffffffffffffff16815260209283015160ff1692810192909252016101f0565b610a5b610a56366004614ca2565b61132b565b6040516101f093929190614d11565b6103d9610a78366004614d3b565b611529565b6103b96116ff565b6104f7610a93366004613e67565b61170b565b6101e6610aa6366004614de9565b611807565b6103d9610ab9366004613e67565b611cc1565b6103d9610acc366004614e6e565b611cd2565b610ae4610adf36600461508e565b611ce3565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9384168152929091166020830152016101f0565b6000610b2482611e6e565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16610b4b85611e6e565b610b73907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16856150e7565b610b7d91906150fe565b90505b9392505050565b6060610b936002611f08565b905090565b610ba0611f15565b6000610bac8280615139565b9050905060005b81811015610cf6576000610bc78480615139565b83818110610bd757610bd76151a1565b905060400201803603810190610bed91906151fc565b604080518082018252602080840180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff908116845263ffffffff42818116858701908152885173ffffffffffffffffffffffffffffffffffffffff9081166000908152600690975295889020965190519092167c010000000000000000000000000000000000000000000000000000000002919092161790935584519051935194955016927f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a92610ce59290917bffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a250600101610bb3565b506000610d066020840184615139565b9050905060005b81811015610e47576000610d246020860186615139565b83818110610d3457610d346151a1565b905060400201803603810190610d4a9190615239565b604080518082018252602080840180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff908116845263ffffffff42818116858701908152885167ffffffffffffffff9081166000908152600590975295889020965190519092167c010000000000000000000000000000000000000000000000000000000002919092161790935584519051935194955016927fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e92610e369290917bffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a250600101610d0d565b50505050565b610e55611f5a565b610e5e81611fdb565b50565b610e69611f5a565b60005b8151811015610e9f57610e97828281518110610e8a57610e8a6151a1565b60200260200101516120d9565b600101610e6c565b5050565b60608160008167ffffffffffffffff811115610ec157610ec1613f8e565b604051908082528060200260200182016040528015610f0657816020015b6040805180820190915260008082526020820152815260200190600190039081610edf5790505b50905060005b82811015610f6357610f3e868683818110610f2957610f296151a1565b9050602002016020810190610a939190613e67565b828281518110610f5057610f506151a1565b6020908102919091010152600101610f0c565b509150505b92915050565b6000610f6882611e6e565b610f81611f5a565b610e9f82826122ab565b60015473ffffffffffffffffffffffffffffffffffffffff163314611011576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b611095611f5a565b610e9f82826126b7565b60008060006110e387878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506127fe92505050565b9250925092506110f533838584612819565b60006111038587018761525c565b905060005b81518110156112fe57604051806040016040528083838151811061112e5761112e6151a1565b6020026020010151602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815260200183838151811061116f5761116f6151a1565b60200260200101516040015163ffffffff168152506006600084848151811061119a5761119a6151a1565b6020908102919091018101515173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251929091015163ffffffff167c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790558151829082908110611232576112326151a1565b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff167f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a838381518110611287576112876151a1565b6020026020010151602001518484815181106112a5576112a56151a1565b6020026020010151604001516040516112ee9291907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff92909216825263ffffffff16602082015260400190565b60405180910390a2600101611108565b505050505050505050565b611311611f5a565b610e5e81612971565b611322611f5a565b610e5e81612afd565b60008060607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff160361138b578592506113b9565b6113b687877f0000000000000000000000000000000000000000000000000000000000000000610b19565b92505b7f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff16831115611458576040517f6a92a483000000000000000000000000000000000000000000000000000000008152600481018490526bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166024820152604401611008565b67ffffffffffffffff8816600090815260096020526040812060010154640100000000900463ffffffff169061148f878784612be7565b9050806020015193508484611516836040805182516024820152602092830151151560448083019190915282518083039091018152606490910190915290810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f181dcf100000000000000000000000000000000000000000000000000000000017905290565b9450945094505050955095509592505050565b67ffffffffffffffff85166000908152600960205260408120600101547501000000000000000000000000000000000000000000900460e01b905b848110156116f6576000848483818110611580576115806151a1565b6115969260206040909202019081019150613e67565b905060008787848181106115ac576115ac6151a1565b90506020028101906115be9190615323565b6115cc906040810190615361565b915050602081111561167c5767ffffffffffffffff89166000908152600a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff861684529091529020546e010000000000000000000000000000900463ffffffff1681111561167c576040517f36f536ca00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401611008565b6116ec84898986818110611692576116926151a1565b90506020028101906116a49190615323565b6116b2906020810190615361565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612d9092505050565b5050600101611564565b50505050505050565b6060610b93600d611f08565b604080518082019091526000808252602082015273ffffffffffffffffffffffffffffffffffffffff8281166000908152600760209081526040918290208251808401909352549283168083527401000000000000000000000000000000000000000090930460ff1690820152906117fe57505073ffffffffffffffffffffffffffffffffffffffff166000908152600660209081526040918290208251808401909352547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811683527c0100000000000000000000000000000000000000000000000000000000900463ffffffff169082015290565b610b8081612de2565b67ffffffffffffffff8083166000908152600960209081526040808320815161022081018352815460ff808216151580845261ffff61010080850482169886019890985263ffffffff630100000085048116978601979097526701000000000000008404871660608601526b0100000000000000000000008404871660808601526f010000000000000000000000000000008404811660a0860152710100000000000000000000000000000000008404871660c08601527501000000000000000000000000000000000000000000808504821660e08088019190915277010000000000000000000000000000000000000000000000860483169987019990995279010000000000000000000000000000000000000000000000000085049091166101208601527b010000000000000000000000000000000000000000000000000000009093048616610140850152600190940154808616610160850152640100000000810486166101808501526801000000000000000081049098166101a084015270010000000000000000000000000000000088049094166101c083015274010000000000000000000000000000000000000000870490931615156101e08201527fffffffff000000000000000000000000000000000000000000000000000000009290950490921b16610200840152909190611a3d576040517f99ac52f200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401611008565b6000611a4c6040850185615139565b9150611aa8905082611a616020870187615361565b905083611a6e8880615361565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061302592505050565b6000600881611abd6080880160608901613e67565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160009081205467ffffffffffffffff16915080611b0c611b066080890160608a01613e67565b89611ce3565b9092509050600080808615611b5257611b46888c611b3060808e0160608f01613e67565b888e8060400190611b419190615139565b6130cf565b91945092509050611b72565b6101c0880151611b6f9063ffffffff16662386f26fc100006150e7565b92505b61010088015160009061ffff1615611bb657611bb3896dffffffffffffffffffffffffffff607088901c16611baa60208f018f615361565b90508b866133ad565b90505b6101a089015160009067ffffffffffffffff16611bdf611bd960808f018f615361565b8d61345c565b600001518563ffffffff168c60a0015161ffff168f8060200190611c039190615361565b611c0e9291506150e7565b8d6080015163ffffffff16611c2391906153c6565b611c2d91906153c6565b611c3791906153c6565b611c51906dffffffffffffffffffffffffffff89166150e7565b611c5b91906150e7565b90507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff87168282611c9267ffffffffffffffff8c16896150e7565b611c9c91906153c6565b611ca691906153c6565b611cb091906150fe565b9d9c50505050505050505050505050565b611cc9611f5a565b610e5e8161351d565b611cda611f5a565b610e5e81613612565b67ffffffffffffffff811660009081526005602090815260408083208151808301909252547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811682527c0100000000000000000000000000000000000000000000000000000000900463ffffffff1691810182905282918203611d9b576040517f2e59db3a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401611008565b6000816020015163ffffffff1642611db391906153d9565b90507f000000000000000000000000000000000000000000000000000000000000000063ffffffff16811115611e54576040517ff08bcb3e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8616600482015263ffffffff7f000000000000000000000000000000000000000000000000000000000000000016602482015260448101829052606401611008565b611e5d86611e6e565b9151919350909150505b9250929050565b600080611e7a8361170b565b9050806020015163ffffffff1660001480611eb2575080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16155b15611f01576040517f06439c6b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401611008565b5192915050565b60606000610b8083613afa565b611f20600233613b56565b611f58576040517fd86ad9cf000000000000000000000000000000000000000000000000000000008152336004820152602401611008565b565b60005473ffffffffffffffffffffffffffffffffffffffff163314611f58576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401611008565b60005b8151811015610e9f576000828281518110611ffb57611ffb6151a1565b60209081029190910181015180518183015173ffffffffffffffffffffffffffffffffffffffff80831660008181526007875260409081902084518154868a018051929096167fffffffffffffffffffffff00000000000000000000000000000000000000000090911681177401000000000000000000000000000000000000000060ff9384160217909255825191825293519093169683019690965293955091939092917f08a5f7f5bb38a81d8e43aca13ecd76431dbf8816ae4699affff7b00b2fc1c464910160405180910390a2505050806001019050611fde565b600061219282600001518360600151846020015185604001516040805173ffffffffffffffffffffffffffffffffffffffff80871660208301528516918101919091527fffffffffffffffffffff00000000000000000000000000000000000000000000831660608201527fffff0000000000000000000000000000000000000000000000000000000000008216608082015260009060a001604051602081830303815290604052805190602001209050949350505050565b60808301516000828152600460205260409081902080549215157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909316929092179091555190915081907f32a4ba3fa3351b11ad555d4c8ec70a744e8705607077a946807030d64b6ab1a39061229f908590600060a08201905073ffffffffffffffffffffffffffffffffffffffff8084511683527fffffffffffffffffffff0000000000000000000000000000000000000000000060208501511660208401527fffff00000000000000000000000000000000000000000000000000000000000060408501511660408401528060608501511660608401525060808301511515608083015292915050565b60405180910390a25050565b60005b82518110156125ce5760008382815181106122cb576122cb6151a1565b6020026020010151905060008160000151905060005b8260200151518110156125c057600083602001518281518110612306576123066151a1565b602002602001015160200151905060008460200151838151811061232c5761232c6151a1565b60200260200101516000015190506020826080015163ffffffff1610156123a95760808201516040517f24ecdc0200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015263ffffffff9091166024820152604401611008565b67ffffffffffffffff84166000818152600a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168085529083529281902086518154938801518389015160608a015160808b015160a08c015115157201000000000000000000000000000000000000027fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff63ffffffff9283166e01000000000000000000000000000002167fffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffffff9383166a0100000000000000000000027fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff61ffff9096166801000000000000000002959095167fffffffffffffffffffffffffffffffffffff000000000000ffffffffffffffff968416640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000909b16939097169290921798909817939093169390931717919091161792909217909155519091907f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b5906125ae908690600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60405180910390a350506001016122e1565b5050508060010190506122ae565b5060005b81518110156126b25760008282815181106125ef576125ef6151a1565b60200260200101516000015190506000838381518110612611576126116151a1565b60209081029190910181015181015167ffffffffffffffff84166000818152600a8452604080822073ffffffffffffffffffffffffffffffffffffffff8516808452955280822080547fffffffffffffffffffffffffff000000000000000000000000000000000000001690555192945090917f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b9190a350506001016125d2565b505050565b60005b825181101561275a576126f08382815181106126d8576126d86151a1565b6020026020010151600d613b8590919063ffffffff16565b1561275257828181518110612707576127076151a1565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba2360405160405180910390a25b6001016126ba565b5060005b81518110156126b25761279482828151811061277c5761277c6151a1565b6020026020010151600d613ba790919063ffffffff16565b156127f6578181815181106127ab576127ab6151a1565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f9160405160405180910390a25b60010161275e565b6040810151604a820151605e90920151909260609290921c91565b6040805173ffffffffffffffffffffffffffffffffffffffff868116602080840191909152908616828401527fffffffffffffffffffff00000000000000000000000000000000000000000000851660608301527fffff00000000000000000000000000000000000000000000000000000000000084166080808401919091528351808403909101815260a09092018352815191810191909120600081815260049092529190205460ff1661296a576040517f9589cb9d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8087166004830152851660248201527fffffffffffffffffffff00000000000000000000000000000000000000000000841660448201527fffff00000000000000000000000000000000000000000000000000000000000083166064820152608401611008565b5050505050565b602081015160005b8151811015612a0c576000828281518110612996576129966151a1565b602002602001015190506129b4816002613ba790919063ffffffff16565b15612a035760405173ffffffffffffffffffffffffffffffffffffffff821681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda775809060200160405180910390a15b50600101612979565b50815160005b8151811015610e47576000828281518110612a2f57612a2f6151a1565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612a9f576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612aaa600282613b85565b5060405173ffffffffffffffffffffffffffffffffffffffff821681527feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef9060200160405180910390a150600101612a12565b60005b8151811015610e9f576000828281518110612b1d57612b1d6151a1565b60200260200101516000015190506000838381518110612b3f57612b3f6151a1565b60209081029190910181015181015173ffffffffffffffffffffffffffffffffffffffff841660008181526008845260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff85169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a25050600101612b00565b60408051808201909152600080825260208201526000839003612c2857506040805180820190915267ffffffffffffffff8216815260006020820152610b80565b6000612c3484866153ec565b90506000612c458560048189615432565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293505050507fffffffff0000000000000000000000000000000000000000000000000000000082167fe7e230f00000000000000000000000000000000000000000000000000000000001612ce25780806020019051810190612cd9919061545c565b92505050610b80565b7f6859a837000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831601612d5e57604051806040016040528082806020019051810190612d4a9190615488565b815260006020909101529250610b80915050565b6040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fd7ed2ad4000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831601610e9f576126b281613bc9565b604080518082019091526000808252602082015260008260000151905060008173ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa158015612e4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e7091906154bb565b5050509150506000811215612eb1576040517f10cb51d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000819050600085602001518473ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f2c919061550b565b612f369190615528565b905060248160ff161115612f6b57612f4f602482615541565b612f5a90600a61567a565b612f6490836150fe565b9150612f8e565b612f76816024615541565b612f8190600a61567a565b612f8b90836150e7565b91505b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff821115612fe4576040517f10cb51d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50604080518082019091527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff909116815263ffffffff42166020820152949350505050565b836040015163ffffffff1683111561307e5760408085015190517f8693378900000000000000000000000000000000000000000000000000000000815263ffffffff909116600482015260248101849052604401611008565b836020015161ffff168211156130c0576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e4784610200015182612d90565b6000808083815b8181101561339f5760008787838181106130f2576130f26151a1565b9050604002018036038101906131089190615689565b67ffffffffffffffff8c166000908152600a60209081526040808320845173ffffffffffffffffffffffffffffffffffffffff168452825291829020825160c081018452905463ffffffff8082168352640100000000820481169383019390935268010000000000000000810461ffff16938201939093526a01000000000000000000008304821660608201526e01000000000000000000000000000083049091166080820152720100000000000000000000000000000000000090910460ff16151560a082018190529192509061322e576101208d01516131f59061ffff16662386f26fc100006150e7565b6131ff90886153c6565b96508c61014001518661321291906156c2565b95508c61016001518561322591906156c2565b94505050613397565b604081015160009061ffff16156132e75760008c73ffffffffffffffffffffffffffffffffffffffff16846000015173ffffffffffffffffffffffffffffffffffffffff161461328a57835161328390611e6e565b905061328d565b508a5b620186a0836040015161ffff166132cf8660200151847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16613c7c90919063ffffffff16565b6132d991906150e7565b6132e391906150fe565b9150505b60608201516132f690886156c2565b965081608001518661330891906156c2565b82519096506000906133279063ffffffff16662386f26fc100006150e7565b9050808210156133465761333b818a6153c6565b985050505050613397565b6000836020015163ffffffff16662386f26fc1000061336591906150e7565b90508083111561338557613379818b6153c6565b99505050505050613397565b61338f838b6153c6565b995050505050505b6001016130d6565b505096509650969350505050565b60008063ffffffff83166133c260e0866150e7565b6133ce876101c06153c6565b6133d891906153c6565b6133e291906153c6565b905060008760c0015163ffffffff168860e0015161ffff168361340591906150e7565b61340f91906153c6565b61010089015190915061ffff166134366dffffffffffffffffffffffffffff8916836150e7565b61344091906150e7565b61345090655af3107a40006150e7565b98975050505050505050565b60408051808201909152600080825260208201526000613488858585610180015163ffffffff16612be7565b9050826060015163ffffffff16816000015111156134d2576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826101e0015180156134e657508060200151155b15610b7d576040517fee433e9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff82160361359c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401611008565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60005b8151811015610e9f576000828281518110613632576136326151a1565b602002602001015190506000838381518110613650576136506151a1565b60200260200101516000015190506000826020015190508167ffffffffffffffff1660001480613689575061018081015163ffffffff16155b806136db57506102008101517fffffffff00000000000000000000000000000000000000000000000000000000167f2812d52c0000000000000000000000000000000000000000000000000000000014155b806136f15750602081610160015163ffffffff16105b806137105750806060015163ffffffff1681610180015163ffffffff16115b15613753576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff83166004820152602401611008565b67ffffffffffffffff82166000908152600960205260408120600101547501000000000000000000000000000000000000000000900460e01b7fffffffff000000000000000000000000000000000000000000000000000000001690036137fb578167ffffffffffffffff167fa937382a486d993de71c220bc8b559242deb4e286a353fa732330b4aa7d13577826040516137ee91906144d2565b60405180910390a261383e565b8167ffffffffffffffff167fa7b607fc10d28a1caf39ab7d27f4c94945db708a576d572781a455c5894fad938260405161383591906144d2565b60405180910390a25b80600960008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a81548163ffffffff021916908363ffffffff1602179055506101a08201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101c08201518160010160106101000a81548163ffffffff021916908363ffffffff1602179055506101e08201518160010160146101000a81548160ff0219169083151502179055506102008201518160010160156101000a81548163ffffffff021916908360e01c0217905550905050505050806001019050613615565b606081600001805480602002602001604051908101604052809291908181526020018280548015613b4a57602002820191906000526020600020905b815481526020019060010190808311613b36575b50505050509050919050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515610b80565b6000610b808373ffffffffffffffffffffffffffffffffffffffff8416613cb9565b6000610b808373ffffffffffffffffffffffffffffffffffffffff8416613d08565b60008151602014613c0857816040517f8d666f600000000000000000000000000000000000000000000000000000000081526004016110089190613ee6565b600082806020019051810190613c1e9190615488565b905073ffffffffffffffffffffffffffffffffffffffff811180613c43575061040081105b15610f6857826040517f8d666f600000000000000000000000000000000000000000000000000000000081526004016110089190613ee6565b6000670de0b6b3a7640000613caf837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff86166150e7565b610b8091906150fe565b6000818152600183016020526040812054613d0057508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610f68565b506000610f68565b60008181526001830160205260408120548015613df1576000613d2c6001836153d9565b8554909150600090613d40906001906153d9565b9050818114613da5576000866000018281548110613d6057613d606151a1565b9060005260206000200154905080876000018481548110613d8357613d836151a1565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080613db657613db66156df565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610f68565b6000915050610f68565b5092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114613e2657600080fd5b919050565b600080600060608486031215613e4057600080fd5b613e4984613e02565b925060208401359150613e5e60408501613e02565b90509250925092565b600060208284031215613e7957600080fd5b610b8082613e02565b6000815180845260005b81811015613ea857602081850181015186830182015201613e8c565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610b806020830184613e82565b6020808252825182820181905260009190848201906040850190845b81811015613f4757835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101613f15565b50909695505050505050565b600060208284031215613f6557600080fd5b813567ffffffffffffffff811115613f7c57600080fd5b820160408185031215610b8057600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613fe057613fe0613f8e565b60405290565b60405160a0810167ffffffffffffffff81118282101715613fe057613fe0613f8e565b60405160c0810167ffffffffffffffff81118282101715613fe057613fe0613f8e565b604051610220810167ffffffffffffffff81118282101715613fe057613fe0613f8e565b6040516060810167ffffffffffffffff81118282101715613fe057613fe0613f8e565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156140ba576140ba613f8e565b604052919050565b600067ffffffffffffffff8211156140dc576140dc613f8e565b5060051b60200190565b60ff81168114610e5e57600080fd5b6000602080838503121561410857600080fd5b823567ffffffffffffffff81111561411f57600080fd5b8301601f8101851361413057600080fd5b803561414361413e826140c2565b614073565b8181526060918202830184019184820191908884111561416257600080fd5b938501935b8385101561420257848903818112156141805760008081fd5b614188613fbd565b61419187613e02565b81526040807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0840112156141c55760008081fd5b6141cd613fbd565b92506141da898901613e02565b83528701356141e8816140e6565b828901528088019190915283529384019391850191614167565b50979650505050505050565b8015158114610e5e57600080fd5b8035613e268161420e565b6000602080838503121561423a57600080fd5b823567ffffffffffffffff81111561425157600080fd5b8301601f8101851361426257600080fd5b803561427061413e826140c2565b81815260a0918202830184019184820191908884111561428f57600080fd5b938501935b838510156142025780858a0312156142ac5760008081fd5b6142b4613fe6565b6142bd86613e02565b8152868601357fffffffffffffffffffff00000000000000000000000000000000000000000000811681146142f25760008081fd5b818801526040868101357fffff0000000000000000000000000000000000000000000000000000000000008116811461432b5760008081fd5b90820152606061433c878201613e02565b9082015260808681013561434f8161420e565b9082015283529384019391850191614294565b60008083601f84011261437457600080fd5b50813567ffffffffffffffff81111561438c57600080fd5b6020830191508360208260051b8501011115611e6757600080fd5b600080602083850312156143ba57600080fd5b823567ffffffffffffffff8111156143d157600080fd5b6143dd85828601614362565b90969095509350505050565b602080825282518282018190526000919060409081850190868401855b828110156144575761444784835180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16825260209081015163ffffffff16910152565b9284019290850190600101614406565b5091979650505050505050565b803567ffffffffffffffff81168114613e2657600080fd5b60006020828403121561448e57600080fd5b610b8082614464565b81517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815260208083015163ffffffff169082015260408101610f68565b815115158152610220810160208301516144f2602084018261ffff169052565b50604083015161450a604084018263ffffffff169052565b506060830151614522606084018263ffffffff169052565b50608083015161453a608084018263ffffffff169052565b5060a083015161455060a084018261ffff169052565b5060c083015161456860c084018263ffffffff169052565b5060e083015161457e60e084018261ffff169052565b506101008381015161ffff9081169184019190915261012080850151909116908301526101408084015163ffffffff9081169184019190915261016080850151821690840152610180808501518216908401526101a08085015167ffffffffffffffff16908401526101c080850151909116908301526101e080840151151590830152610200808401517fffffffff000000000000000000000000000000000000000000000000000000008116828501525b505092915050565b803563ffffffff81168114613e2657600080fd5b803561ffff81168114613e2657600080fd5b600082601f83011261466f57600080fd5b8135602061467f61413e836140c2565b82815260069290921b8401810191818101908684111561469e57600080fd5b8286015b848110156146eb57604081890312156146bb5760008081fd5b6146c3613fbd565b6146cc82614464565b81526146d9858301613e02565b818601528352918301916040016146a2565b509695505050505050565b6000806040838503121561470957600080fd5b67ffffffffffffffff8335111561471f57600080fd5b83601f84358501011261473157600080fd5b61474161413e84358501356140c2565b8335840180358083526020808401939260059290921b9091010186101561476757600080fd5b602085358601015b85358601803560051b016020018110156149745767ffffffffffffffff8135111561479957600080fd5b8035863587010160407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828a030112156147d257600080fd5b6147da613fbd565b6147e660208301614464565b815267ffffffffffffffff6040830135111561480157600080fd5b88603f60408401358401011261481657600080fd5b61482c61413e60206040850135850101356140c2565b6020604084810135850182810135808552928401939260e00201018b101561485357600080fd5b6040808501358501015b6040858101358601602081013560e00201018110156149555760e0818d03121561488657600080fd5b61488e613fbd565b61489782613e02565b815260c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838f030112156148cb57600080fd5b6148d3614009565b6148df60208401614638565b81526148ed60408401614638565b60208201526148fe6060840161464c565b604082015261490f60808401614638565b606082015261492060a08401614638565b608082015261493260c084013561420e565b60c083013560a0820152602082810191909152908452929092019160e00161485d565b508060208401525050808552505060208301925060208101905061476f565b5092505067ffffffffffffffff6020840135111561499157600080fd5b6149a1846020850135850161465e565b90509250929050565b600082601f8301126149bb57600080fd5b813560206149cb61413e836140c2565b8083825260208201915060208460051b8701019350868411156149ed57600080fd5b602086015b848110156146eb57614a0381613e02565b83529183019183016149f2565b60008060408385031215614a2357600080fd5b823567ffffffffffffffff80821115614a3b57600080fd5b614a47868387016149aa565b93506020850135915080821115614a5d57600080fd5b50614a6a858286016149aa565b9150509250929050565b60008083601f840112614a8657600080fd5b50813567ffffffffffffffff811115614a9e57600080fd5b602083019150836020828501011115611e6757600080fd5b60008060008060408587031215614acc57600080fd5b843567ffffffffffffffff80821115614ae457600080fd5b614af088838901614a74565b90965094506020870135915080821115614b0957600080fd5b50614b1687828801614a74565b95989497509550505050565b60008060408385031215614b3557600080fd5b614b3e83614464565b91506149a160208401613e02565b600060208284031215614b5e57600080fd5b813567ffffffffffffffff80821115614b7657600080fd5b9083019060408286031215614b8a57600080fd5b614b92613fbd565b823582811115614ba157600080fd5b614bad878286016149aa565b825250602083013582811115614bc257600080fd5b614bce878286016149aa565b60208301525095945050505050565b60006020808385031215614bf057600080fd5b823567ffffffffffffffff811115614c0757600080fd5b8301601f81018513614c1857600080fd5b8035614c2661413e826140c2565b81815260069190911b82018301908381019087831115614c4557600080fd5b928401925b82841015614c975760408489031215614c635760008081fd5b614c6b613fbd565b614c7485613e02565b8152614c81868601614464565b8187015282526040939093019290840190614c4a565b979650505050505050565b600080600080600060808688031215614cba57600080fd5b614cc386614464565b9450614cd160208701613e02565b935060408601359250606086013567ffffffffffffffff811115614cf457600080fd5b614d0088828901614a74565b969995985093965092949392505050565b8381528215156020820152606060408201526000614d326060830184613e82565b95945050505050565b600080600080600060608688031215614d5357600080fd5b614d5c86614464565b9450602086013567ffffffffffffffff80821115614d7957600080fd5b614d8589838a01614362565b90965094506040880135915080821115614d9e57600080fd5b818801915088601f830112614db257600080fd5b813581811115614dc157600080fd5b8960208260061b8501011115614dd657600080fd5b9699959850939650602001949392505050565b60008060408385031215614dfc57600080fd5b614e0583614464565b9150602083013567ffffffffffffffff811115614e2157600080fd5b830160a08186031215614e3357600080fd5b809150509250929050565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114613e2657600080fd5b60006020808385031215614e8157600080fd5b823567ffffffffffffffff811115614e9857600080fd5b8301601f81018513614ea957600080fd5b8035614eb761413e826140c2565b8181526102409182028301840191848201919088841115614ed757600080fd5b938501935b838510156142025784890381811215614ef55760008081fd5b614efd613fbd565b614f0687614464565b8152610220807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084011215614f3b5760008081fd5b614f4361402c565b9250614f5089890161421c565b83526040614f5f818a0161464c565b8a8501526060614f70818b01614638565b8286015260809150614f83828b01614638565b9085015260a0614f948a8201614638565b8286015260c09150614fa7828b0161464c565b9085015260e0614fb88a8201614638565b828601526101009150614fcc828b0161464c565b90850152610120614fde8a820161464c565b828601526101409150614ff2828b0161464c565b908501526101606150048a8201614638565b828601526101809150615018828b01614638565b908501526101a061502a8a8201614638565b828601526101c0915061503e828b01614464565b908501526101e06150508a8201614638565b828601526102009150615064828b0161421c565b90850152615073898301614e3e565b90840152508088019190915283529384019391850191614edc565b600080604083850312156150a157600080fd5b6150aa83613e02565b91506149a160208401614464565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082028115828204841417610f6857610f686150b8565b600082615134577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261516e57600080fd5b83018035915067ffffffffffffffff82111561518957600080fd5b6020019150600681901b3603821315611e6757600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81168114613e2657600080fd5b60006040828403121561520e57600080fd5b615216613fbd565b61521f83613e02565b815261522d602084016151d0565b60208201529392505050565b60006040828403121561524b57600080fd5b615253613fbd565b61521f83614464565b6000602080838503121561526f57600080fd5b823567ffffffffffffffff81111561528657600080fd5b8301601f8101851361529757600080fd5b80356152a561413e826140c2565b818152606091820283018401918482019190888411156152c457600080fd5b938501935b838510156142025780858a0312156152e15760008081fd5b6152e9614050565b6152f286613e02565b81526152ff8787016151d0565b878201526040615310818801614638565b90820152835293840193918501916152c9565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8183360301811261535757600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261539657600080fd5b83018035915067ffffffffffffffff8211156153b157600080fd5b602001915036819003821315611e6757600080fd5b80820180821115610f6857610f686150b8565b81810381811115610f6857610f686150b8565b7fffffffff0000000000000000000000000000000000000000000000000000000081358181169160048510156146305760049490940360031b84901b1690921692915050565b6000808585111561544257600080fd5b8386111561544f57600080fd5b5050820193919092039150565b60006040828403121561546e57600080fd5b615476613fbd565b82518152602083015161522d8161420e565b60006020828403121561549a57600080fd5b5051919050565b805169ffffffffffffffffffff81168114613e2657600080fd5b600080600080600060a086880312156154d357600080fd5b6154dc866154a1565b94506020860151935060408601519250606086015191506154ff608087016154a1565b90509295509295909350565b60006020828403121561551d57600080fd5b8151610b80816140e6565b60ff8181168382160190811115610f6857610f686150b8565b60ff8281168282160390811115610f6857610f686150b8565b600181815b808511156155b357817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115615599576155996150b8565b808516156155a657918102915b93841c939080029061555f565b509250929050565b6000826155ca57506001610f68565b816155d757506000610f68565b81600181146155ed57600281146155f757615613565b6001915050610f68565b60ff841115615608576156086150b8565b50506001821b610f68565b5060208310610133831016604e8410600b8410161715615636575081810a610f68565b615640838361555a565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115615672576156726150b8565b029392505050565b6000610b8060ff8416836155bb565b60006040828403121561569b57600080fd5b6156a3613fbd565b6156ac83613e02565b8152602083013560208201528091505092915050565b63ffffffff818116838216019080821115613dfb57613dfb6150b8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000818000a",
}

var PriceRegistryABI = PriceRegistryMetaData.ABI

var PriceRegistryBin = PriceRegistryMetaData.Bin

func DeployPriceRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig PriceRegistryStaticConfig, priceUpdaters []common.Address, feeTokens []common.Address, tokenPriceFeeds []PriceRegistryTokenPriceFeedUpdate, tokenTransferFeeConfigArgs []PriceRegistryTokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs []PriceRegistryPremiumMultiplierWeiPerEthArgs, destChainConfigArgs []PriceRegistryDestChainConfigArgs) (common.Address, *types.Transaction, *PriceRegistry, error) {
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

func (_PriceRegistry *PriceRegistryCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (PriceRegistryDestChainConfig, error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	if err != nil {
		return *new(PriceRegistryDestChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(PriceRegistryDestChainConfig)).(*PriceRegistryDestChainConfig)

	return out0, err

}

func (_PriceRegistry *PriceRegistrySession) GetDestChainConfig(destChainSelector uint64) (PriceRegistryDestChainConfig, error) {
	return _PriceRegistry.Contract.GetDestChainConfig(&_PriceRegistry.CallOpts, destChainSelector)
}

func (_PriceRegistry *PriceRegistryCallerSession) GetDestChainConfig(destChainSelector uint64) (PriceRegistryDestChainConfig, error) {
	return _PriceRegistry.Contract.GetDestChainConfig(&_PriceRegistry.CallOpts, destChainSelector)
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

func (_PriceRegistry *PriceRegistryCaller) ProcessMessageArgs(opts *bind.CallOpts, destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte) (ProcessMessageArgs,

	error) {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "processMessageArgs", destChainSelector, feeToken, feeTokenAmount, extraArgs)

	outstruct := new(ProcessMessageArgs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MsgFeeJuels = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsOutOfOrderExecution = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.ConvertedExtraArgs = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

func (_PriceRegistry *PriceRegistrySession) ProcessMessageArgs(destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte) (ProcessMessageArgs,

	error) {
	return _PriceRegistry.Contract.ProcessMessageArgs(&_PriceRegistry.CallOpts, destChainSelector, feeToken, feeTokenAmount, extraArgs)
}

func (_PriceRegistry *PriceRegistryCallerSession) ProcessMessageArgs(destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte) (ProcessMessageArgs,

	error) {
	return _PriceRegistry.Contract.ProcessMessageArgs(&_PriceRegistry.CallOpts, destChainSelector, feeToken, feeTokenAmount, extraArgs)
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

func (_PriceRegistry *PriceRegistryCaller) ValidatePoolReturnData(opts *bind.CallOpts, destChainSelector uint64, rampTokenAmounts []InternalRampTokenAmount, sourceTokenAmounts []ClientEVMTokenAmount) error {
	var out []interface{}
	err := _PriceRegistry.contract.Call(opts, &out, "validatePoolReturnData", destChainSelector, rampTokenAmounts, sourceTokenAmounts)

	if err != nil {
		return err
	}

	return err

}

func (_PriceRegistry *PriceRegistrySession) ValidatePoolReturnData(destChainSelector uint64, rampTokenAmounts []InternalRampTokenAmount, sourceTokenAmounts []ClientEVMTokenAmount) error {
	return _PriceRegistry.Contract.ValidatePoolReturnData(&_PriceRegistry.CallOpts, destChainSelector, rampTokenAmounts, sourceTokenAmounts)
}

func (_PriceRegistry *PriceRegistryCallerSession) ValidatePoolReturnData(destChainSelector uint64, rampTokenAmounts []InternalRampTokenAmount, sourceTokenAmounts []ClientEVMTokenAmount) error {
	return _PriceRegistry.Contract.ValidatePoolReturnData(&_PriceRegistry.CallOpts, destChainSelector, rampTokenAmounts, sourceTokenAmounts)
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

func (_PriceRegistry *PriceRegistryTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []PriceRegistryDestChainConfigArgs) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_PriceRegistry *PriceRegistrySession) ApplyDestChainConfigUpdates(destChainConfigArgs []PriceRegistryDestChainConfigArgs) (*types.Transaction, error) {
	return _PriceRegistry.Contract.ApplyDestChainConfigUpdates(&_PriceRegistry.TransactOpts, destChainConfigArgs)
}

func (_PriceRegistry *PriceRegistryTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []PriceRegistryDestChainConfigArgs) (*types.Transaction, error) {
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

func (_PriceRegistry *PriceRegistryTransactor) OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "onReport", metadata, report)
}

func (_PriceRegistry *PriceRegistrySession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _PriceRegistry.Contract.OnReport(&_PriceRegistry.TransactOpts, metadata, report)
}

func (_PriceRegistry *PriceRegistryTransactorSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _PriceRegistry.Contract.OnReport(&_PriceRegistry.TransactOpts, metadata, report)
}

func (_PriceRegistry *PriceRegistryTransactor) SetReportPermissions(opts *bind.TransactOpts, permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _PriceRegistry.contract.Transact(opts, "setReportPermissions", permissions)
}

func (_PriceRegistry *PriceRegistrySession) SetReportPermissions(permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _PriceRegistry.Contract.SetReportPermissions(&_PriceRegistry.TransactOpts, permissions)
}

func (_PriceRegistry *PriceRegistryTransactorSession) SetReportPermissions(permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _PriceRegistry.Contract.SetReportPermissions(&_PriceRegistry.TransactOpts, permissions)
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
	DestChainConfig   PriceRegistryDestChainConfig
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

type PriceRegistryDestChainConfigUpdatedIterator struct {
	Event *PriceRegistryDestChainConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryDestChainConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryDestChainConfigUpdated)
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
		it.Event = new(PriceRegistryDestChainConfigUpdated)
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

func (it *PriceRegistryDestChainConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryDestChainConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryDestChainConfigUpdated struct {
	DestChainSelector uint64
	DestChainConfig   PriceRegistryDestChainConfig
	Raw               types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterDestChainConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*PriceRegistryDestChainConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "DestChainConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryDestChainConfigUpdatedIterator{contract: _PriceRegistry.contract, event: "DestChainConfigUpdated", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchDestChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryDestChainConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "DestChainConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryDestChainConfigUpdated)
				if err := _PriceRegistry.contract.UnpackLog(event, "DestChainConfigUpdated", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseDestChainConfigUpdated(log types.Log) (*PriceRegistryDestChainConfigUpdated, error) {
	event := new(PriceRegistryDestChainConfigUpdated)
	if err := _PriceRegistry.contract.UnpackLog(event, "DestChainConfigUpdated", log); err != nil {
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

type PriceRegistryReportPermissionSetIterator struct {
	Event *PriceRegistryReportPermissionSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *PriceRegistryReportPermissionSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceRegistryReportPermissionSet)
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
		it.Event = new(PriceRegistryReportPermissionSet)
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

func (it *PriceRegistryReportPermissionSetIterator) Error() error {
	return it.fail
}

func (it *PriceRegistryReportPermissionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type PriceRegistryReportPermissionSet struct {
	ReportId   [32]byte
	Permission KeystoneFeedsPermissionHandlerPermission
	Raw        types.Log
}

func (_PriceRegistry *PriceRegistryFilterer) FilterReportPermissionSet(opts *bind.FilterOpts, reportId [][32]byte) (*PriceRegistryReportPermissionSetIterator, error) {

	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}

	logs, sub, err := _PriceRegistry.contract.FilterLogs(opts, "ReportPermissionSet", reportIdRule)
	if err != nil {
		return nil, err
	}
	return &PriceRegistryReportPermissionSetIterator{contract: _PriceRegistry.contract, event: "ReportPermissionSet", logs: logs, sub: sub}, nil
}

func (_PriceRegistry *PriceRegistryFilterer) WatchReportPermissionSet(opts *bind.WatchOpts, sink chan<- *PriceRegistryReportPermissionSet, reportId [][32]byte) (event.Subscription, error) {

	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}

	logs, sub, err := _PriceRegistry.contract.WatchLogs(opts, "ReportPermissionSet", reportIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(PriceRegistryReportPermissionSet)
				if err := _PriceRegistry.contract.UnpackLog(event, "ReportPermissionSet", log); err != nil {
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

func (_PriceRegistry *PriceRegistryFilterer) ParseReportPermissionSet(log types.Log) (*PriceRegistryReportPermissionSet, error) {
	event := new(PriceRegistryReportPermissionSet)
	if err := _PriceRegistry.contract.UnpackLog(event, "ReportPermissionSet", log); err != nil {
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
type ProcessMessageArgs struct {
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
	case _PriceRegistry.abi.Events["DestChainConfigUpdated"].ID:
		return _PriceRegistry.ParseDestChainConfigUpdated(log)
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
	case _PriceRegistry.abi.Events["ReportPermissionSet"].ID:
		return _PriceRegistry.ParseReportPermissionSet(log)
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
	return common.HexToHash("0xa937382a486d993de71c220bc8b559242deb4e286a353fa732330b4aa7d13577")
}

func (PriceRegistryDestChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xa7b607fc10d28a1caf39ab7d27f4c94945db708a576d572781a455c5894fad93")
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

func (PriceRegistryReportPermissionSet) Topic() common.Hash {
	return common.HexToHash("0x32a4ba3fa3351b11ad555d4c8ec70a744e8705607077a946807030d64b6ab1a3")
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

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (PriceRegistryDestChainConfig, error)

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

	GetValidatedTokenPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	ProcessMessageArgs(opts *bind.CallOpts, destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte) (ProcessMessageArgs,

		error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	ValidatePoolReturnData(opts *bind.CallOpts, destChainSelector uint64, rampTokenAmounts []InternalRampTokenAmount, sourceTokenAmounts []ClientEVMTokenAmount) error

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []PriceRegistryDestChainConfigArgs) (*types.Transaction, error)

	ApplyFeeTokensUpdates(opts *bind.TransactOpts, feeTokensToAdd []common.Address, feeTokensToRemove []common.Address) (*types.Transaction, error)

	ApplyPremiumMultiplierWeiPerEthUpdates(opts *bind.TransactOpts, premiumMultiplierWeiPerEthArgs []PriceRegistryPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []PriceRegistryTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []PriceRegistryTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error)

	OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error)

	SetReportPermissions(opts *bind.TransactOpts, permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error)

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

	FilterDestChainConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*PriceRegistryDestChainConfigUpdatedIterator, error)

	WatchDestChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *PriceRegistryDestChainConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigUpdated(log types.Log) (*PriceRegistryDestChainConfigUpdated, error)

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

	FilterReportPermissionSet(opts *bind.FilterOpts, reportId [][32]byte) (*PriceRegistryReportPermissionSetIterator, error)

	WatchReportPermissionSet(opts *bind.WatchOpts, sink chan<- *PriceRegistryReportPermissionSet, reportId [][32]byte) (event.Subscription, error)

	ParseReportPermissionSet(log types.Log) (*PriceRegistryReportPermissionSet, error)

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
