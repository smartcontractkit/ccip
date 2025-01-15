package price_registry_1_2_0

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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessThreshold\",\"type\":\"uint32\"}],\"internalType\":\"structPriceRegistry.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"priceUpdaters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIPriceRegistry.TokenPriceFeedConfig\",\"name\":\"feedConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.TokenPriceFeedUpdate[]\",\"name\":\"tokenPriceFeeds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structPriceRegistry.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"internalType\":\"structPriceRegistry.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chain\",\"type\":\"uint64\"}],\"name\":\"ChainNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataFeedValueOutOfUint224Range\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"DestinationChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtraArgOutOfOrderExecutionMustBeTrue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"}],\"name\":\"InvalidDestBytesOverhead\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidDestChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStaticConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgFeeJuels\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint256\"}],\"name\":\"MessageFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SourceTokenDataTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePassed\",\"type\":\"uint256\"}],\"name\":\"StaleGasPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"UnauthorizedCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedCallerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedCallerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"indexed\":false,\"internalType\":\"structPriceRegistry.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"indexed\":false,\"internalType\":\"structPriceRegistry.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"FeeTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"FeeTokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"name\":\"PremiumMultiplierWeiPerEthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIPriceRegistry.TokenPriceFeedConfig\",\"name\":\"priceFeedConfig\",\"type\":\"tuple\"}],\"name\":\"PriceFeedPerTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenTransferFeeConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structPriceRegistry.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"name\":\"TokenTransferFeeConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UsdPerTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChain\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UsdPerUnitGasUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"addedCallers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removedCallers\",\"type\":\"address[]\"}],\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\"}],\"name\":\"applyAuthorizedCallerUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"internalType\":\"structPriceRegistry.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyDestChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"feeTokensToAdd\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"feeTokensToRemove\",\"type\":\"address[]\"}],\"name\":\"applyFeeTokensUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structPriceRegistry.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyPremiumMultiplierWeiPerEthUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfigRemoveArgs[]\",\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\"}],\"name\":\"applyTokenTransferFeeConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"}],\"name\":\"convertTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAuthorizedCallers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"internalType\":\"structPriceRegistry.DestChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestinationChainGasPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.TimestampedPackedUint224\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPremiumMultiplierWeiPerEth\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stalenessThreshold\",\"type\":\"uint32\"}],\"internalType\":\"structPriceRegistry.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getTokenAndGasPrices\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"tokenPrice\",\"type\":\"uint224\"},{\"internalType\":\"uint224\",\"name\":\"gasPriceValue\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.TimestampedPackedUint224\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPriceFeedConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIPriceRegistry.TokenPriceFeedConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getTokenPrices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.TimestampedPackedUint224[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenTransferFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structPriceRegistry.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getValidatedFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getValidatedTokenPrice\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"name\":\"processMessageArgs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"msgFeeJuels\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOutOfOrderExecution\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"convertedExtraArgs\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint224\",\"name\":\"usdPerToken\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint224\",\"name\":\"usdPerUnitGas\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.GasPriceUpdate[]\",\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.PriceUpdates\",\"name\":\"priceUpdates\",\"type\":\"tuple\"}],\"name\":\"updatePrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIPriceRegistry.TokenPriceFeedConfig\",\"name\":\"feedConfig\",\"type\":\"tuple\"}],\"internalType\":\"structPriceRegistry.TokenPriceFeedUpdate[]\",\"name\":\"tokenPriceFeedUpdates\",\"type\":\"tuple[]\"}],\"name\":\"updateTokenPriceFeeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.RampTokenAmount[]\",\"name\":\"rampTokenAmounts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\"}],\"name\":\"validatePoolReturnData\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162006a4238038062006a4283398101604081905262000034916200188e565b8533806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000207565b5050604080518082018252838152815160008152602080820190935291810191909152620000ee9150620002b2565b5060208701516001600160a01b0316158062000112575086516001600160601b0316155b80620001265750604087015163ffffffff16155b15620001455760405163d794ef9560e01b815260040160405180910390fd5b6020878101516001600160a01b031660a05287516001600160601b031660805260408089015163ffffffff1660c05280516000815291820190526200018c90869062000401565b620001978462000549565b620001a2816200061a565b620001ad8262000a9f565b60408051600080825260208201909252620001fa91859190620001f3565b6040805180820190915260008082526020820152815260200190600190039081620001cb5790505b5062000b6b565b5050505050505062001b4c565b336001600160a01b03821603620002615760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b602081015160005b815181101562000342576000828281518110620002db57620002db620019ad565b60209081029190910101519050620002f560028262000ea4565b1562000338576040516001600160a01b03821681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda775809060200160405180910390a15b50600101620002ba565b50815160005b8151811015620003fb576000828281518110620003695762000369620019ad565b6020026020010151905060006001600160a01b0316816001600160a01b031603620003a7576040516342bcdf7f60e11b815260040160405180910390fd5b620003b460028262000ec4565b506040516001600160a01b03821681527feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef9060200160405180910390a15060010162000348565b50505050565b60005b8251811015620004a25762000440838281518110620004275762000427620019ad565b6020026020010151600a62000ec460201b90919060201c565b1562000499578281815181106200045b576200045b620019ad565b60200260200101516001600160a01b03167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba2360405160405180910390a25b60010162000404565b5060005b81518110156200054457620004e2828281518110620004c957620004c9620019ad565b6020026020010151600a62000edb60201b90919060201c565b156200053b57818181518110620004fd57620004fd620019ad565b60200260200101516001600160a01b03167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f9160405160405180910390a25b600101620004a6565b505050565b60005b8151811015620006165760008282815181106200056d576200056d620019ad565b6020908102919091018101518051818301516001600160a01b0380831660008181526006875260409081902084518154868a018051929096166001600160a81b03199091168117600160a01b60ff9384160217909255825191825293519093169683019690965293955091939092917f08a5f7f5bb38a81d8e43aca13ecd76431dbf8816ae4699affff7b00b2fc1c464910160405180910390a25050508060010190506200054c565b5050565b60005b8151811015620006165760008282815181106200063e576200063e620019ad565b6020026020010151905060008383815181106200065f576200065f620019ad565b6020026020010151600001519050600082602001519050816001600160401b03166000148062000698575061018081015163ffffffff16155b80620006ba57506102008101516001600160e01b031916630a04b54b60e21b14155b80620006d75750602063ffffffff1681610160015163ffffffff16105b80620006f75750806060015163ffffffff1681610180015163ffffffff16115b15620007225760405163c35aa79d60e01b81526001600160401b038316600482015260240162000083565b6001600160401b038216600090815260086020526040812060010154600160a81b900460e01b6001600160e01b0319169003620007a257816001600160401b03167fa937382a486d993de71c220bc8b559242deb4e286a353fa732330b4aa7d1357782604051620007949190620019c3565b60405180910390a2620007e6565b816001600160401b03167fa7b607fc10d28a1caf39ab7d27f4c94945db708a576d572781a455c5894fad9382604051620007dd9190620019c3565b60405180910390a25b8060086000846001600160401b03166001600160401b0316815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a81548163ffffffff021916908363ffffffff1602179055506101a08201518160010160086101000a8154816001600160401b0302191690836001600160401b031602179055506101c08201518160010160106101000a81548163ffffffff021916908363ffffffff1602179055506101e08201518160010160146101000a81548160ff0219169083151502179055506102008201518160010160156101000a81548163ffffffff021916908360e01c02179055509050505050508060010190506200061d565b60005b81518110156200061657600082828151811062000ac35762000ac3620019ad565b6020026020010151600001519050600083838151811062000ae85762000ae8620019ad565b6020908102919091018101518101516001600160a01b03841660008181526007845260409081902080546001600160401b0319166001600160401b0385169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a2505060010162000aa2565b60005b825181101562000dde57600083828151811062000b8f5762000b8f620019ad565b6020026020010151905060008160000151905060005b82602001515181101562000dcf5760008360200151828151811062000bce5762000bce620019ad565b602002602001015160200151905060008460200151838151811062000bf75762000bf7620019ad565b6020026020010151600001519050602063ffffffff16826080015163ffffffff16101562000c565760808201516040516312766e0160e11b81526001600160a01b038316600482015263ffffffff909116602482015260440162000083565b6001600160401b03841660008181526009602090815260408083206001600160a01b0386168085529083529281902086518154938801518389015160608a015160808b015160a08c01511515600160901b0260ff60901b1963ffffffff928316600160701b021664ffffffffff60701b199383166a01000000000000000000000263ffffffff60501b1961ffff90961668010000000000000000029590951665ffffffffffff60401b19968416640100000000026001600160401b0319909b16939097169290921798909817939093169390931717919091161792909217909155519091907f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b59062000dbc908690600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60405180910390a3505060010162000ba5565b50505080600101905062000b6e565b5060005b81518110156200054457600082828151811062000e035762000e03620019ad565b6020026020010151600001519050600083838151811062000e285762000e28620019ad565b6020908102919091018101518101516001600160401b03841660008181526009845260408082206001600160a01b038516808452955280822080546001600160981b03191690555192945090917f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b9190a3505060010162000de2565b600062000ebb836001600160a01b03841662000ef2565b90505b92915050565b600062000ebb836001600160a01b03841662000ff6565b600062000ebb836001600160a01b03841662001048565b6000818152600183016020526040812054801562000feb57600062000f1960018362001b14565b855490915060009062000f2f9060019062001b14565b905081811462000f9b57600086600001828154811062000f535762000f53620019ad565b906000526020600020015490508087600001848154811062000f795762000f79620019ad565b6000918252602080832090910192909255918252600188019052604090208390555b855486908062000faf5762000faf62001b36565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000ebe565b600091505062000ebe565b60008181526001830160205260408120546200103f5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000ebe565b50600062000ebe565b6000818152600183016020526040812054801562000feb5760006200106f60018362001b14565b8554909150600090620010859060019062001b14565b905080821462000f9b57600086600001828154811062000f535762000f53620019ad565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715620010e457620010e4620010a9565b60405290565b60405160c081016001600160401b0381118282101715620010e457620010e4620010a9565b60405161022081016001600160401b0381118282101715620010e457620010e4620010a9565b604051601f8201601f191681016001600160401b0381118282101715620011605762001160620010a9565b604052919050565b80516001600160a01b03811681146200118057600080fd5b919050565b805163ffffffff811681146200118057600080fd5b600060608284031215620011ad57600080fd5b604051606081016001600160401b0381118282101715620011d257620011d2620010a9565b604052825190915081906001600160601b0381168114620011f257600080fd5b8152620012026020840162001168565b6020820152620012156040840162001185565b60408201525092915050565b60006001600160401b038211156200123d576200123d620010a9565b5060051b60200190565b600082601f8301126200125957600080fd5b81516020620012726200126c8362001221565b62001135565b8083825260208201915060208460051b8701019350868411156200129557600080fd5b602086015b84811015620012bc57620012ae8162001168565b83529183019183016200129a565b509695505050505050565b600082601f830112620012d957600080fd5b81516020620012ec6200126c8362001221565b828152606092830285018201928282019190878511156200130c57600080fd5b8387015b858110156200139f57808903828112156200132b5760008081fd5b62001335620010bf565b620013408362001168565b8152604080601f1984011215620013575760008081fd5b62001361620010bf565b92506200137088850162001168565b835283015160ff81168114620013865760008081fd5b8288015280870191909152845292840192810162001310565b5090979650505050505050565b80516001600160401b03811681146200118057600080fd5b805161ffff811681146200118057600080fd5b805180151581146200118057600080fd5b600082601f830112620013fa57600080fd5b815160206200140d6200126c8362001221565b82815260059290921b840181019181810190868411156200142d57600080fd5b8286015b84811015620012bc5780516001600160401b03808211156200145257600080fd5b908801906040601f19838c0381018213156200146d57600080fd5b62001477620010bf565b62001484898601620013ac565b815282850151848111156200149857600080fd5b8086019550508c603f860112620014ae57600080fd5b888501519350620014c36200126c8562001221565b84815260e09094028501830193898101908e861115620014e257600080fd5b958401955b85871015620015bb57868f0360e08112156200150257600080fd5b6200150c620010bf565b620015178962001168565b815260c086830112156200152a57600080fd5b62001534620010ea565b9150620015438d8a0162001185565b825262001552878a0162001185565b8d8301526200156460608a01620013c4565b878301526200157660808a0162001185565b60608301526200158960a08a0162001185565b60808301526200159c60c08a01620013d7565b60a0830152808d0191909152825260e09690960195908a0190620014e7565b828b01525087525050509284019250830162001431565b600082601f830112620015e457600080fd5b81516020620015f76200126c8362001221565b82815260069290921b840181019181810190868411156200161757600080fd5b8286015b84811015620012bc5760408189031215620016365760008081fd5b62001640620010bf565b6200164b8262001168565b81526200165a858301620013ac565b818601528352918301916040016200161b565b80516001600160e01b0319811681146200118057600080fd5b600082601f8301126200169857600080fd5b81516020620016ab6200126c8362001221565b8281526102409283028501820192828201919087851115620016cc57600080fd5b8387015b858110156200139f5780890382811215620016eb5760008081fd5b620016f5620010bf565b6200170083620013ac565b815261022080601f1984011215620017185760008081fd5b620017226200110f565b925062001731888501620013d7565b8352604062001742818601620013c4565b8985015260606200175581870162001185565b82860152608091506200176a82870162001185565b9085015260a06200177d86820162001185565b8286015260c0915062001792828701620013c4565b9085015260e0620017a586820162001185565b828601526101009150620017bb828701620013c4565b90850152610120620017cf868201620013c4565b828601526101409150620017e5828701620013c4565b90850152610160620017f986820162001185565b8286015261018091506200180f82870162001185565b908501526101a06200182386820162001185565b828601526101c0915062001839828701620013ac565b908501526101e06200184d86820162001185565b82860152610200915062001863828701620013d7565b90850152620018748583016200166d565b9084015250808701919091528452928401928101620016d0565b6000806000806000806000610120888a031215620018ab57600080fd5b620018b789896200119a565b60608901519097506001600160401b0380821115620018d557600080fd5b620018e38b838c0162001247565b975060808a0151915080821115620018fa57600080fd5b620019088b838c0162001247565b965060a08a01519150808211156200191f57600080fd5b6200192d8b838c01620012c7565b955060c08a01519150808211156200194457600080fd5b620019528b838c01620013e8565b945060e08a01519150808211156200196957600080fd5b620019778b838c01620015d2565b93506101008a01519150808211156200198f57600080fd5b506200199e8a828b0162001686565b91505092959891949750929550565b634e487b7160e01b600052603260045260246000fd5b81511515815261022081016020830151620019e4602084018261ffff169052565b506040830151620019fd604084018263ffffffff169052565b50606083015162001a16606084018263ffffffff169052565b50608083015162001a2f608084018263ffffffff169052565b5060a083015162001a4660a084018261ffff169052565b5060c083015162001a5f60c084018263ffffffff169052565b5060e083015162001a7660e084018261ffff169052565b506101008381015161ffff9081169184019190915261012080850151909116908301526101408084015163ffffffff9081169184019190915261016080850151821690840152610180808501518216908401526101a0808501516001600160401b0316908401526101c080850151909116908301526101e080840151151590830152610200928301516001600160e01b031916929091019190915290565b8181038181111562000ebe57634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60805160a05160c051614e9c62001ba6600039600081816102d901528181611ad30152611b3c01526000818161029d0152818161104e01526110ae015260008181610269015281816110d701526111470152614e9c6000f3fe608060405234801561001057600080fd5b50600436106101b85760003560e01c80637afac322116100f9578063cc88924c11610097578063d8694ccd11610071578063d8694ccd14610a5c578063f2fde38b14610a6f578063f700042a14610a82578063ffdb4b3714610a9557600080fd5b8063cc88924c14610a2e578063cdc73d5114610a41578063d02641a014610a4957600080fd5b806391a2749a116100d357806391a2749a14610939578063a69c64c01461094c578063bf78e03f1461095f578063c4276bfc14610a0c57600080fd5b80637afac3221461078e57806382b49eb0146107a15780638da5cb5b1461091157600080fd5b8063407e108611610166578063514e8cff11610140578063514e8cff146104385780636def4ce7146104db578063770e2dc41461077357806379ba50971461078657600080fd5b8063407e1086146103c557806345ac924d146103d85780634ab35b0b146103f857600080fd5b8063181f5a7711610197578063181f5a77146103525780632451a6271461039b5780633937306f146103b057600080fd5b806241e5be146101bd578063061877e3146101e357806306285c691461023c575b600080fd5b6101d06101cb366004613889565b610add565b6040519081526020015b60405180910390f35b6102236101f13660046138c5565b73ffffffffffffffffffffffffffffffffffffffff1660009081526007602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff90911681526020016101da565b610306604080516060810182526000808252602082018190529181019190915260405180606001604052807f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000063ffffffff16815250905090565b6040805182516bffffffffffffffffffffffff16815260208084015173ffffffffffffffffffffffffffffffffffffffff16908201529181015163ffffffff16908201526060016101da565b61038e6040518060400160405280601781526020017f5072696365526567697374727920312e362e302d64657600000000000000000081525081565b6040516101da9190613944565b6103a3610b4b565b6040516101da9190613957565b6103c36103be3660046139b1565b610b5c565b005b6103c36103d3366004613b0d565b610e11565b6103eb6103e6366004613c6b565b610e25565b6040516101da9190613cad565b61040b6104063660046138c5565b610ef0565b6040517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90911681526020016101da565b6104ce610446366004613d40565b60408051808201909152600080825260208201525067ffffffffffffffff166000908152600460209081526040918290208251808401909352547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811683527c0100000000000000000000000000000000000000000000000000000000900463ffffffff169082015290565b6040516101da9190613d5b565b6107666104e9366004613d40565b6040805161022081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c081018290526101e081018290526102008101919091525067ffffffffffffffff908116600090815260086020908152604091829020825161022081018452815460ff8082161515835261ffff61010080840482169685019690965263ffffffff630100000084048116978501979097526701000000000000008304871660608501526b0100000000000000000000008304871660808501526f010000000000000000000000000000008304811660a0850152710100000000000000000000000000000000008304871660c08501527501000000000000000000000000000000000000000000808404821660e08087019190915277010000000000000000000000000000000000000000000000850483169786019790975279010000000000000000000000000000000000000000000000000084049091166101208501527b010000000000000000000000000000000000000000000000000000009092048616610140840152600190930154808616610160840152640100000000810486166101808401526801000000000000000081049096166101a083015270010000000000000000000000000000000086049094166101c082015274010000000000000000000000000000000000000000850490911615156101e08201527fffffffff0000000000000000000000000000000000000000000000000000000092909304901b1661020082015290565b6040516101da9190613d96565b6103c3610781366004613fd3565b610efb565b6103c3610f11565b6103c361079c3660046142ed565b611013565b6108b16107af366004614351565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a08101919091525067ffffffffffffffff91909116600090815260096020908152604080832073ffffffffffffffffffffffffffffffffffffffff94909416835292815290829020825160c081018452905463ffffffff8082168352640100000000820481169383019390935268010000000000000000810461ffff16938201939093526a01000000000000000000008304821660608201526e01000000000000000000000000000083049091166080820152720100000000000000000000000000000000000090910460ff16151560a082015290565b6040516101da9190600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101da565b6103c361094736600461437b565b611025565b6103c361095a36600461440c565b611036565b6109d861096d3660046138c5565b6040805180820182526000808252602091820181905273ffffffffffffffffffffffffffffffffffffffff93841681526006825282902082518084019093525492831682527401000000000000000000000000000000000000000090920460ff169181019190915290565b60408051825173ffffffffffffffffffffffffffffffffffffffff16815260209283015160ff1692810192909252016101da565b610a1f610a1a3660046144d1565b611047565b6040516101da9392919061456c565b6103c3610a3c366004614596565b611245565b6103a361141b565b6104ce610a573660046138c5565b611427565b6101d0610a6a366004614631565b611523565b6103c3610a7d3660046138c5565b6119dd565b6103c3610a903660046146b6565b6119ee565b610aa8610aa33660046148d6565b6119ff565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9384168152929091166020830152016101da565b6000610ae882611b8a565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16610b0f85611b8a565b610b37907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168561492f565b610b419190614946565b90505b9392505050565b6060610b576002611c24565b905090565b610b64611c31565b6000610b708280614981565b9050905060005b81811015610cba576000610b8b8480614981565b83818110610b9b57610b9b6149e9565b905060400201803603810190610bb19190614a44565b604080518082018252602080840180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff908116845263ffffffff42818116858701908152885173ffffffffffffffffffffffffffffffffffffffff9081166000908152600590975295889020965190519092167c010000000000000000000000000000000000000000000000000000000002919092161790935584519051935194955016927f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a92610ca99290917bffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a250600101610b77565b506000610cca6020840184614981565b9050905060005b81811015610e0b576000610ce86020860186614981565b83818110610cf857610cf86149e9565b905060400201803603810190610d0e9190614a81565b604080518082018252602080840180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff908116845263ffffffff42818116858701908152885167ffffffffffffffff9081166000908152600490975295889020965190519092167c010000000000000000000000000000000000000000000000000000000002919092161790935584519051935194955016927fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e92610dfa9290917bffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a250600101610cd1565b50505050565b610e19611c76565b610e2281611cf7565b50565b60608160008167ffffffffffffffff811115610e4357610e436139ec565b604051908082528060200260200182016040528015610e8857816020015b6040805180820190915260008082526020820152815260200190600190039081610e615790505b50905060005b82811015610ee557610ec0868683818110610eab57610eab6149e9565b9050602002016020810190610a5791906138c5565b828281518110610ed257610ed26149e9565b6020908102919091010152600101610e8e565b509150505b92915050565b6000610eea82611b8a565b610f03611c76565b610f0d8282611df5565b5050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610f97576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61101b611c76565b610f0d8282612207565b61102d611c76565b610e228161234e565b61103e611c76565b610e22816124da565b60008060607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16036110a7578592506110d5565b6110d287877f0000000000000000000000000000000000000000000000000000000000000000610add565b92505b7f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff16831115611174576040517f6a92a483000000000000000000000000000000000000000000000000000000008152600481018490526bffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166024820152604401610f8e565b67ffffffffffffffff8816600090815260086020526040812060010154640100000000900463ffffffff16906111ab8787846125c4565b9050806020015193508484611232836040805182516024820152602092830151151560448083019190915282518083039091018152606490910190915290810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f181dcf100000000000000000000000000000000000000000000000000000000017905290565b9450945094505050955095509592505050565b67ffffffffffffffff85166000908152600860205260408120600101547501000000000000000000000000000000000000000000900460e01b905b8481101561141257600084848381811061129c5761129c6149e9565b6112b292602060409092020190810191506138c5565b905060008787848181106112c8576112c86149e9565b90506020028101906112da9190614aa4565b6112e8906040810190614ae2565b91505060208111156113985767ffffffffffffffff8916600090815260096020908152604080832073ffffffffffffffffffffffffffffffffffffffff861684529091529020546e010000000000000000000000000000900463ffffffff16811115611398576040517f36f536ca00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610f8e565b611408848989868181106113ae576113ae6149e9565b90506020028101906113c09190614aa4565b6113ce906020810190614ae2565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061276d92505050565b5050600101611280565b50505050505050565b6060610b57600a611c24565b604080518082019091526000808252602082015273ffffffffffffffffffffffffffffffffffffffff8281166000908152600660209081526040918290208251808401909352549283168083527401000000000000000000000000000000000000000090930460ff16908201529061151a57505073ffffffffffffffffffffffffffffffffffffffff166000908152600560209081526040918290208251808401909352547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811683527c0100000000000000000000000000000000000000000000000000000000900463ffffffff169082015290565b610b44816127bf565b67ffffffffffffffff8083166000908152600860209081526040808320815161022081018352815460ff808216151580845261ffff61010080850482169886019890985263ffffffff630100000085048116978601979097526701000000000000008404871660608601526b0100000000000000000000008404871660808601526f010000000000000000000000000000008404811660a0860152710100000000000000000000000000000000008404871660c08601527501000000000000000000000000000000000000000000808504821660e08088019190915277010000000000000000000000000000000000000000000000860483169987019990995279010000000000000000000000000000000000000000000000000085049091166101208601527b010000000000000000000000000000000000000000000000000000009093048616610140850152600190940154808616610160850152640100000000810486166101808501526801000000000000000081049098166101a084015270010000000000000000000000000000000088049094166101c083015274010000000000000000000000000000000000000000870490931615156101e08201527fffffffff000000000000000000000000000000000000000000000000000000009290950490921b16610200840152909190611759576040517f99ac52f200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610f8e565b60006117686040850185614981565b91506117c490508261177d6020870187614ae2565b90508361178a8880614ae2565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612a0292505050565b60006007816117d960808801606089016138c5565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160009081205467ffffffffffffffff169150806118286118226080890160608a016138c5565b896119ff565b909250905060008080861561186e57611862888c61184c60808e0160608f016138c5565b888e806040019061185d9190614981565b612aac565b9194509250905061188e565b6101c088015161188b9063ffffffff16662386f26fc1000061492f565b92505b61010088015160009061ffff16156118d2576118cf896dffffffffffffffffffffffffffff607088901c166118c660208f018f614ae2565b90508b86612d8a565b90505b6101a089015160009067ffffffffffffffff166118fb6118f560808f018f614ae2565b8d612e3a565b600001518563ffffffff168c60a0015161ffff168f806020019061191f9190614ae2565b61192a92915061492f565b8d6080015163ffffffff1661193f9190614b47565b6119499190614b47565b6119539190614b47565b61196d906dffffffffffffffffffffffffffff891661492f565b611977919061492f565b90507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff871682826119ae67ffffffffffffffff8c168961492f565b6119b89190614b47565b6119c29190614b47565b6119cc9190614946565b9d9c50505050505050505050505050565b6119e5611c76565b610e2281612efb565b6119f6611c76565b610e2281612ff0565b67ffffffffffffffff811660009081526004602090815260408083208151808301909252547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811682527c0100000000000000000000000000000000000000000000000000000000900463ffffffff1691810182905282918203611ab7576040517f2e59db3a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610f8e565b6000816020015163ffffffff1642611acf9190614b5a565b90507f000000000000000000000000000000000000000000000000000000000000000063ffffffff16811115611b70576040517ff08bcb3e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8616600482015263ffffffff7f000000000000000000000000000000000000000000000000000000000000000016602482015260448101829052606401610f8e565b611b7986611b8a565b9151919350909150505b9250929050565b600080611b9683611427565b9050806020015163ffffffff1660001480611bce575080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16155b15611c1d576040517f06439c6b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610f8e565b5192915050565b60606000610b44836134de565b611c3c60023361353a565b611c74576040517fd86ad9cf000000000000000000000000000000000000000000000000000000008152336004820152602401610f8e565b565b60005473ffffffffffffffffffffffffffffffffffffffff163314611c74576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610f8e565b60005b8151811015610f0d576000828281518110611d1757611d176149e9565b60209081029190910181015180518183015173ffffffffffffffffffffffffffffffffffffffff80831660008181526006875260409081902084518154868a018051929096167fffffffffffffffffffffff00000000000000000000000000000000000000000090911681177401000000000000000000000000000000000000000060ff9384160217909255825191825293519093169683019690965293955091939092917f08a5f7f5bb38a81d8e43aca13ecd76431dbf8816ae4699affff7b00b2fc1c464910160405180910390a2505050806001019050611cfa565b60005b825181101561211e576000838281518110611e1557611e156149e9565b6020026020010151905060008160000151905060005b82602001515181101561211057600083602001518281518110611e5057611e506149e9565b6020026020010151602001519050600084602001518381518110611e7657611e766149e9565b6020026020010151600001519050602063ffffffff16826080015163ffffffff161015611ef95760808201516040517f24ecdc0200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015263ffffffff9091166024820152604401610f8e565b67ffffffffffffffff8416600081815260096020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168085529083529281902086518154938801518389015160608a015160808b015160a08c015115157201000000000000000000000000000000000000027fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff63ffffffff9283166e01000000000000000000000000000002167fffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffffff9383166a0100000000000000000000027fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff61ffff9096166801000000000000000002959095167fffffffffffffffffffffffffffffffffffff000000000000ffffffffffffffff968416640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000909b16939097169290921798909817939093169390931717919091161792909217909155519091907f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b5906120fe908690600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60405180910390a35050600101611e2b565b505050806001019050611df8565b5060005b815181101561220257600082828151811061213f5761213f6149e9565b60200260200101516000015190506000838381518110612161576121616149e9565b60209081029190910181015181015167ffffffffffffffff8416600081815260098452604080822073ffffffffffffffffffffffffffffffffffffffff8516808452955280822080547fffffffffffffffffffffffffff000000000000000000000000000000000000001690555192945090917f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b9190a35050600101612122565b505050565b60005b82518110156122aa57612240838281518110612228576122286149e9565b6020026020010151600a61356990919063ffffffff16565b156122a257828181518110612257576122576149e9565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba2360405160405180910390a25b60010161220a565b5060005b8151811015612202576122e48282815181106122cc576122cc6149e9565b6020026020010151600a61358b90919063ffffffff16565b15612346578181815181106122fb576122fb6149e9565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f9160405160405180910390a25b6001016122ae565b602081015160005b81518110156123e9576000828281518110612373576123736149e9565b602002602001015190506123918160026135ad90919063ffffffff16565b156123e05760405173ffffffffffffffffffffffffffffffffffffffff821681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda775809060200160405180910390a15b50600101612356565b50815160005b8151811015610e0b57600082828151811061240c5761240c6149e9565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361247c576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612487600282613569565b5060405173ffffffffffffffffffffffffffffffffffffffff821681527feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef9060200160405180910390a1506001016123ef565b60005b8151811015610f0d5760008282815181106124fa576124fa6149e9565b6020026020010151600001519050600083838151811061251c5761251c6149e9565b60209081029190910181015181015173ffffffffffffffffffffffffffffffffffffffff841660008181526007845260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff85169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a250506001016124dd565b6040805180820190915260008082526020820152600083900361260557506040805180820190915267ffffffffffffffff8216815260006020820152610b44565b60006126118486614b6d565b905060006126228560048189614bb3565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293505050507fffffffff0000000000000000000000000000000000000000000000000000000082167fe7e230f000000000000000000000000000000000000000000000000000000000016126bf57808060200190518101906126b69190614bdd565b92505050610b44565b7f6859a837000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000083160161273b576040518060400160405280828060200190518101906127279190614c09565b815260006020909101529250610b44915050565b6040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fd7ed2ad4000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831601610f0d57612202816135cf565b604080518082019091526000808252602082015260008260000151905060008173ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa158015612829573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061284d9190614c3c565b505050915050600081121561288e576040517f10cb51d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000819050600085602001518473ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129099190614c8c565b6129139190614ca9565b905060248160ff1611156129485761292c602482614cc2565b61293790600a614dfb565b6129419083614946565b915061296b565b612953816024614cc2565b61295e90600a614dfb565b612968908361492f565b91505b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8211156129c1576040517f10cb51d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50604080518082019091527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff909116815263ffffffff42166020820152949350505050565b836040015163ffffffff16831115612a5b5760408085015190517f8693378900000000000000000000000000000000000000000000000000000000815263ffffffff909116600482015260248101849052604401610f8e565b836020015161ffff16821115612a9d576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e0b8461020001518261276d565b6000808083815b81811015612d7c576000878783818110612acf57612acf6149e9565b905060400201803603810190612ae59190614e0a565b67ffffffffffffffff8c166000908152600960209081526040808320845173ffffffffffffffffffffffffffffffffffffffff168452825291829020825160c081018452905463ffffffff8082168352640100000000820481169383019390935268010000000000000000810461ffff16938201939093526a01000000000000000000008304821660608201526e01000000000000000000000000000083049091166080820152720100000000000000000000000000000000000090910460ff16151560a0820181905291925090612c0b576101208d0151612bd29061ffff16662386f26fc1000061492f565b612bdc9088614b47565b96508c610140015186612bef9190614e43565b95508c610160015185612c029190614e43565b94505050612d74565b604081015160009061ffff1615612cc45760008c73ffffffffffffffffffffffffffffffffffffffff16846000015173ffffffffffffffffffffffffffffffffffffffff1614612c67578351612c6090611b8a565b9050612c6a565b508a5b620186a0836040015161ffff16612cac8660200151847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1661368290919063ffffffff16565b612cb6919061492f565b612cc09190614946565b9150505b6060820151612cd39088614e43565b9650816080015186612ce59190614e43565b8251909650600090612d049063ffffffff16662386f26fc1000061492f565b905080821015612d2357612d18818a614b47565b985050505050612d74565b6000836020015163ffffffff16662386f26fc10000612d42919061492f565b905080831115612d6257612d56818b614b47565b99505050505050612d74565b612d6c838b614b47565b995050505050505b600101612ab3565b505096509650969350505050565b60008063ffffffff8316612da06101408661492f565b612dac876101c0614b47565b612db69190614b47565b612dc09190614b47565b905060008760c0015163ffffffff168860e0015161ffff1683612de3919061492f565b612ded9190614b47565b61010089015190915061ffff16612e146dffffffffffffffffffffffffffff89168361492f565b612e1e919061492f565b612e2e90655af3107a400061492f565b98975050505050505050565b60408051808201909152600080825260208201526000612e66858585610180015163ffffffff166125c4565b9050826060015163ffffffff1681600001511115612eb0576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826101e001518015612ec457508060200151155b15610b41576040517fee433e9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff821603612f7a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610f8e565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60005b8151811015610f0d576000828281518110613010576130106149e9565b60200260200101519050600083838151811061302e5761302e6149e9565b60200260200101516000015190506000826020015190508167ffffffffffffffff1660001480613067575061018081015163ffffffff16155b806130b957506102008101517fffffffff00000000000000000000000000000000000000000000000000000000167f2812d52c0000000000000000000000000000000000000000000000000000000014155b806130d55750602063ffffffff1681610160015163ffffffff16105b806130f45750806060015163ffffffff1681610180015163ffffffff16115b15613137576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff83166004820152602401610f8e565b67ffffffffffffffff82166000908152600860205260408120600101547501000000000000000000000000000000000000000000900460e01b7fffffffff000000000000000000000000000000000000000000000000000000001690036131df578167ffffffffffffffff167fa937382a486d993de71c220bc8b559242deb4e286a353fa732330b4aa7d13577826040516131d29190613d96565b60405180910390a2613222565b8167ffffffffffffffff167fa7b607fc10d28a1caf39ab7d27f4c94945db708a576d572781a455c5894fad93826040516132199190613d96565b60405180910390a25b80600860008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a81548163ffffffff021916908363ffffffff1602179055506101a08201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101c08201518160010160106101000a81548163ffffffff021916908363ffffffff1602179055506101e08201518160010160146101000a81548160ff0219169083151502179055506102008201518160010160156101000a81548163ffffffff021916908360e01c0217905550905050505050806001019050612ff3565b60608160000180548060200260200160405190810160405280929190818152602001828054801561352e57602002820191906000526020600020905b81548152602001906001019080831161351a575b50505050509050919050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515610b44565b6000610b448373ffffffffffffffffffffffffffffffffffffffff84166136bf565b6000610b448373ffffffffffffffffffffffffffffffffffffffff841661370e565b6000610b448373ffffffffffffffffffffffffffffffffffffffff8416613808565b6000815160201461360e57816040517f8d666f60000000000000000000000000000000000000000000000000000000008152600401610f8e9190613944565b6000828060200190518101906136249190614c09565b905073ffffffffffffffffffffffffffffffffffffffff811180613649575061040081105b15610eea57826040517f8d666f60000000000000000000000000000000000000000000000000000000008152600401610f8e9190613944565b6000670de0b6b3a76400006136b5837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff861661492f565b610b449190614946565b600081815260018301602052604081205461370657508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610eea565b506000610eea565b600081815260018301602052604081205480156137f7576000613732600183614b5a565b855490915060009061374690600190614b5a565b90508082146137ab576000866000018281548110613766576137666149e9565b9060005260206000200154905080876000018481548110613789576137896149e9565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806137bc576137bc614e60565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610eea565b6000915050610eea565b5092915050565b600081815260018301602052604081205480156137f757600061382c600183614b5a565b855490915060009061384090600190614b5a565b90508181146137ab576000866000018281548110613766576137666149e9565b803573ffffffffffffffffffffffffffffffffffffffff8116811461388457600080fd5b919050565b60008060006060848603121561389e57600080fd5b6138a784613860565b9250602084013591506138bc60408501613860565b90509250925092565b6000602082840312156138d757600080fd5b610b4482613860565b6000815180845260005b81811015613906576020818501810151868301820152016138ea565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610b4460208301846138e0565b6020808252825182820181905260009190848201906040850190845b818110156139a557835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101613973565b50909695505050505050565b6000602082840312156139c357600080fd5b813567ffffffffffffffff8111156139da57600080fd5b820160408185031215610b4457600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613a3e57613a3e6139ec565b60405290565b60405160c0810167ffffffffffffffff81118282101715613a3e57613a3e6139ec565b604051610220810167ffffffffffffffff81118282101715613a3e57613a3e6139ec565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613ad257613ad26139ec565b604052919050565b600067ffffffffffffffff821115613af457613af46139ec565b5060051b60200190565b60ff81168114610e2257600080fd5b60006020808385031215613b2057600080fd5b823567ffffffffffffffff811115613b3757600080fd5b8301601f81018513613b4857600080fd5b8035613b5b613b5682613ada565b613a8b565b81815260609182028301840191848201919088841115613b7a57600080fd5b938501935b83851015613c1a5784890381811215613b985760008081fd5b613ba0613a1b565b613ba987613860565b81526040807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084011215613bdd5760008081fd5b613be5613a1b565b9250613bf2898901613860565b8352870135613c0081613afe565b828901528088019190915283529384019391850191613b7f565b50979650505050505050565b60008083601f840112613c3857600080fd5b50813567ffffffffffffffff811115613c5057600080fd5b6020830191508360208260051b8501011115611b8357600080fd5b60008060208385031215613c7e57600080fd5b823567ffffffffffffffff811115613c9557600080fd5b613ca185828601613c26565b90969095509350505050565b602080825282518282018190526000919060409081850190868401855b82811015613d1b57613d0b84835180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16825260209081015163ffffffff16910152565b9284019290850190600101613cca565b5091979650505050505050565b803567ffffffffffffffff8116811461388457600080fd5b600060208284031215613d5257600080fd5b610b4482613d28565b81517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815260208083015163ffffffff169082015260408101610eea565b81511515815261022081016020830151613db6602084018261ffff169052565b506040830151613dce604084018263ffffffff169052565b506060830151613de6606084018263ffffffff169052565b506080830151613dfe608084018263ffffffff169052565b5060a0830151613e1460a084018261ffff169052565b5060c0830151613e2c60c084018263ffffffff169052565b5060e0830151613e4260e084018261ffff169052565b506101008381015161ffff9081169184019190915261012080850151909116908301526101408084015163ffffffff9081169184019190915261016080850151821690840152610180808501518216908401526101a08085015167ffffffffffffffff16908401526101c080850151909116908301526101e080840151151590830152610200808401517fffffffff000000000000000000000000000000000000000000000000000000008116828501525b505092915050565b803563ffffffff8116811461388457600080fd5b803561ffff8116811461388457600080fd5b8015158114610e2257600080fd5b803561388481613f22565b600082601f830112613f4c57600080fd5b81356020613f5c613b5683613ada565b82815260069290921b84018101918181019086841115613f7b57600080fd5b8286015b84811015613fc85760408189031215613f985760008081fd5b613fa0613a1b565b613fa982613d28565b8152613fb6858301613860565b81860152835291830191604001613f7f565b509695505050505050565b60008060408385031215613fe657600080fd5b67ffffffffffffffff83351115613ffc57600080fd5b83601f84358501011261400e57600080fd5b61401e613b568435850135613ada565b8335840180358083526020808401939260059290921b9091010186101561404457600080fd5b602085358601015b85358601803560051b016020018110156142515767ffffffffffffffff8135111561407657600080fd5b8035863587010160407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828a030112156140af57600080fd5b6140b7613a1b565b6140c360208301613d28565b815267ffffffffffffffff604083013511156140de57600080fd5b88603f6040840135840101126140f357600080fd5b614109613b566020604085013585010135613ada565b6020604084810135850182810135808552928401939260e00201018b101561413057600080fd5b6040808501358501015b6040858101358601602081013560e00201018110156142325760e0818d03121561416357600080fd5b61416b613a1b565b61417482613860565b815260c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838f030112156141a857600080fd5b6141b0613a44565b6141bc60208401613efc565b81526141ca60408401613efc565b60208201526141db60608401613f10565b60408201526141ec60808401613efc565b60608201526141fd60a08401613efc565b608082015261420f60c0840135613f22565b60c083013560a0820152602082810191909152908452929092019160e00161413a565b508060208401525050808552505060208301925060208101905061404c565b5092505067ffffffffffffffff6020840135111561426e57600080fd5b61427e8460208501358501613f3b565b90509250929050565b600082601f83011261429857600080fd5b813560206142a8613b5683613ada565b8083825260208201915060208460051b8701019350868411156142ca57600080fd5b602086015b84811015613fc8576142e081613860565b83529183019183016142cf565b6000806040838503121561430057600080fd5b823567ffffffffffffffff8082111561431857600080fd5b61432486838701614287565b9350602085013591508082111561433a57600080fd5b5061434785828601614287565b9150509250929050565b6000806040838503121561436457600080fd5b61436d83613d28565b915061427e60208401613860565b60006020828403121561438d57600080fd5b813567ffffffffffffffff808211156143a557600080fd5b90830190604082860312156143b957600080fd5b6143c1613a1b565b8235828111156143d057600080fd5b6143dc87828601614287565b8252506020830135828111156143f157600080fd5b6143fd87828601614287565b60208301525095945050505050565b6000602080838503121561441f57600080fd5b823567ffffffffffffffff81111561443657600080fd5b8301601f8101851361444757600080fd5b8035614455613b5682613ada565b81815260069190911b8201830190838101908783111561447457600080fd5b928401925b828410156144c657604084890312156144925760008081fd5b61449a613a1b565b6144a385613860565b81526144b0868601613d28565b8187015282526040939093019290840190614479565b979650505050505050565b6000806000806000608086880312156144e957600080fd5b6144f286613d28565b945061450060208701613860565b935060408601359250606086013567ffffffffffffffff8082111561452457600080fd5b818801915088601f83011261453857600080fd5b81358181111561454757600080fd5b89602082850101111561455957600080fd5b9699959850939650602001949392505050565b838152821515602082015260606040820152600061458d60608301846138e0565b95945050505050565b6000806000806000606086880312156145ae57600080fd5b6145b786613d28565b9450602086013567ffffffffffffffff808211156145d457600080fd5b6145e089838a01613c26565b909650945060408801359150808211156145f957600080fd5b818801915088601f83011261460d57600080fd5b81358181111561461c57600080fd5b8960208260061b850101111561455957600080fd5b6000806040838503121561464457600080fd5b61464d83613d28565b9150602083013567ffffffffffffffff81111561466957600080fd5b830160a0818603121561467b57600080fd5b809150509250929050565b80357fffffffff000000000000000000000000000000000000000000000000000000008116811461388457600080fd5b600060208083850312156146c957600080fd5b823567ffffffffffffffff8111156146e057600080fd5b8301601f810185136146f157600080fd5b80356146ff613b5682613ada565b818152610240918202830184019184820191908884111561471f57600080fd5b938501935b83851015613c1a578489038181121561473d5760008081fd5b614745613a1b565b61474e87613d28565b8152610220807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0840112156147835760008081fd5b61478b613a67565b9250614798898901613f30565b835260406147a7818a01613f10565b8a85015260606147b8818b01613efc565b82860152608091506147cb828b01613efc565b9085015260a06147dc8a8201613efc565b8286015260c091506147ef828b01613f10565b9085015260e06148008a8201613efc565b828601526101009150614814828b01613f10565b908501526101206148268a8201613f10565b82860152610140915061483a828b01613f10565b9085015261016061484c8a8201613efc565b828601526101809150614860828b01613efc565b908501526101a06148728a8201613efc565b828601526101c09150614886828b01613d28565b908501526101e06148988a8201613efc565b8286015261020091506148ac828b01613f30565b908501526148bb898301614686565b90840152508088019190915283529384019391850191614724565b600080604083850312156148e957600080fd5b6148f283613860565b915061427e60208401613d28565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082028115828204841417610eea57610eea614900565b60008261497c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126149b657600080fd5b83018035915067ffffffffffffffff8211156149d157600080fd5b6020019150600681901b3603821315611b8357600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116811461388457600080fd5b600060408284031215614a5657600080fd5b614a5e613a1b565b614a6783613860565b8152614a7560208401614a18565b60208201529392505050565b600060408284031215614a9357600080fd5b614a9b613a1b565b614a6783613d28565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112614ad857600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112614b1757600080fd5b83018035915067ffffffffffffffff821115614b3257600080fd5b602001915036819003821315611b8357600080fd5b80820180821115610eea57610eea614900565b81810381811115610eea57610eea614900565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015613ef45760049490940360031b84901b1690921692915050565b60008085851115614bc357600080fd5b83861115614bd057600080fd5b5050820193919092039150565b600060408284031215614bef57600080fd5b614bf7613a1b565b825181526020830151614a7581613f22565b600060208284031215614c1b57600080fd5b5051919050565b805169ffffffffffffffffffff8116811461388457600080fd5b600080600080600060a08688031215614c5457600080fd5b614c5d86614c22565b9450602086015193506040860151925060608601519150614c8060808701614c22565b90509295509295909350565b600060208284031215614c9e57600080fd5b8151610b4481613afe565b60ff8181168382160190811115610eea57610eea614900565b60ff8281168282160390811115610eea57610eea614900565b600181815b80851115614d3457817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115614d1a57614d1a614900565b80851615614d2757918102915b93841c9390800290614ce0565b509250929050565b600082614d4b57506001610eea565b81614d5857506000610eea565b8160018114614d6e5760028114614d7857614d94565b6001915050610eea565b60ff841115614d8957614d89614900565b50506001821b610eea565b5060208310610133831016604e8410600b8410161715614db7575081810a610eea565b614dc18383614cdb565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115614df357614df3614900565b029392505050565b6000610b4460ff841683614d3c565b600060408284031215614e1c57600080fd5b614e24613a1b565b614e2d83613860565b8152602083013560208201528091505092915050565b63ffffffff81811683821601908082111561380157613801614900565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000818000a",
}

var PriceRegistryABI = PriceRegistryMetaData.ABI

var PriceRegistryBin = PriceRegistryMetaData.Bin

func DeployPriceRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig PriceRegistryStaticConfig, priceUpdaters []common.Address, feeTokens []common.Address, tokenPriceFeeds []PriceRegistryTokenPriceFeedUpdate, tokenTransferFeeConfigArgs []PriceRegistryTokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs []PriceRegistryPremiumMultiplierWeiPerEthArgs, destChainConfigArgs []PriceRegistryDestChainConfigArgs) (common.Address, *generated.Transaction, *PriceRegistry, error) {
	parsed, err := PriceRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}
	if generated.IsZKSync(backend) {
		address, ethTx, contractBind, _ := generated.DeployContract(auth, parsed, common.FromHex(PriceRegistryZKBin), backend, staticConfig, priceUpdaters, feeTokens, tokenPriceFeeds, tokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs, destChainConfigArgs)
		contractReturn := &PriceRegistry{address: address, abi: *parsed, PriceRegistryCaller: PriceRegistryCaller{contract: contractBind}, PriceRegistryTransactor: PriceRegistryTransactor{contract: contractBind}, PriceRegistryFilterer: PriceRegistryFilterer{contract: contractBind}}
		return address, ethTx, contractReturn, err
	}
	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PriceRegistryBin), backend, staticConfig, priceUpdaters, feeTokens, tokenPriceFeeds, tokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs, destChainConfigArgs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, &generated.Transaction{Transaction: tx, HashZks: tx.Hash()}, &PriceRegistry{address: address, abi: *parsed, PriceRegistryCaller: PriceRegistryCaller{contract: contract}, PriceRegistryTransactor: PriceRegistryTransactor{contract: contract}, PriceRegistryFilterer: PriceRegistryFilterer{contract: contract}}, nil
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

var PriceRegistryZKBin = ("0x0003000000000002002c000000000002000200000001035500000000030100190000006003300270000008430030019d000008430330019700000001002001900000007a0000c13d0000008004000039000000400040043f000000040030008c0000009b0000413d000000000201043b000000e0022002700000088a0020009c000000a70000a13d0000088b0020009c000001110000213d000008950020009c000002e30000a13d000008960020009c000003250000213d000008990020009c000004c50000613d0000089a0020009c0000009b0000c13d000000240030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000402100370000000000502043b000008490050009c0000009b0000213d0000002302500039000000000032004b0000009b0000813d0000000402500039000000000221034f000000000202043b000008490020009c000000a10000213d00000005042002100000003f044000390000084a04400197000008b00040009c000000a10000213d0000008004400039000000400040043f000000800020043f000000060220021000000024042001bf0000000002540019000000000032004b0000009b0000213d000000240040008c00000da70000c13d000000000100041a00000848011001970000000002000411000000000012004b00000ba30000c13d000000800100043d000000000001004b000003230000613d002400000000001d0000002402000029000000000021004b00001ad60000a13d0000000501200210000000a0011000390000000001010433000000002101043400000848011001970000000002020433002200000002001d002300000001001d000000000010043f0000000701000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d00000022020000290000084902200197000000000101043b000000000301041a0000087803300197000000000323019f000000000031041b000000400100043d0000000000210435000008430010009c000008430100804100000040011002100000000002000414000008430020009c0000084302008041000000c002200210000000000112019f00000855011001c70000800d02000039000000020300003900000879040000410000002305000029210820fe0000040f00000001002001900000009b0000613d0000002402000029002400010020003d000000800100043d000000240010006b000000440000413d000003230000013d000000e004000039000000400040043f0000000002000416000000000002004b0000009b0000c13d0000001f023000390000084402200197000000e002200039000000400020043f0000001f0530018f0000084506300198000000e0026000390000008c0000613d000000000701034f000000007807043c0000000004840436000000000024004b000000880000c13d000000000005004b000000990000613d000000000161034f0000000304500210000000000502043300000000054501cf000000000545022f000000000101043b0000010004400089000000000141022f00000000014101cf000000000151019f0000000000120435000001200030008c0000009d0000813d00000000010000190000210a00010430000000400100043d002400000001001d000008460010009c000001330000a13d000008d701000041000000000010043f0000004101000039000000040010043f00000877010000410000210a000104300000089e0020009c000002630000a13d0000089f0020009c000002cc0000a13d000008a00020009c000003040000213d000008a30020009c000004090000613d000008a40020009c0000009b0000c13d000000240030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000401100370000000000101043b000008490010009c0000009b0000213d000002a002000039000000400020043f000000800000043f000000a00000043f000000c00000043f000000e00000043f000001000000043f000001200000043f000001400000043f000001600000043f000001800000043f000001a00000043f000001c00000043f000001e00000043f000002000000043f000002200000043f000002400000043f000002600000043f000002800000043f000000000010043f0000000801000039000000200010043f210820d30000040f002400000001001d000002a00100003921081b9e0000040f0000002403000029000000000103041a000000ff001001900000000002000039000000010200c039000002a00020043f00000008021002700000ffff0220018f000002c00020043f00000018021002700000084302200197000002e00020043f00000038021002700000084302200197000003000020043f00000058021002700000084302200197000003200020043f00000078021002700000ffff0220018f000003400020043f00000088021002700000084302200197000003600020043f000000a8021002700000ffff0220018f000003800020043f000000b8021002700000ffff0220018f000003a00020043f000000c8021002700000ffff0220018f000003c00020043f000000d8011002700000084301100197000003e00010043f0000000101300039000000000101041a0000084302100197000004000020043f00000020021002700000084302200197000004200020043f00000040021002700000084902200197000004400020043f00000080021002700000084302200197000004600020043f0000085f001001980000000002000039000000010200c039000004800020043f00000038011002100000086101100197000004a00010043f000000400100043d002400000001001d000002a00200003921081b3c0000040f00000a8d0000013d0000088c0020009c000002f10000a13d0000088d0020009c0000038d0000213d000008900020009c000005910000613d000008910020009c0000009b0000c13d000000240030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000401100370000000000601043b000008480060009c0000009b0000213d000000000100041a00000848011001970000000005000411000000000015004b00000a970000c13d000000000056004b00000aa10000c13d0000088801000041000000800010043f0000002001000039000000840010043f0000001701000039000000a40010043f000008b501000041000000c40010043f000008b3010000410000210a0001043000000024010000290000006001100039000000400010043f000000e00100043d000008470010009c0000009b0000213d00000024020000290000000001120436002300000001001d000001000100043d000008480010009c0000009b0000213d00000023020000290000000000120435000001200100043d000008430010009c0000009b0000213d00000024020000290000004002200039002200000002001d0000000000120435000001400200043d000008490020009c0000009b0000213d000000e001300039000000ff04200039000000000014004b0000009b0000813d000000e0042000390000000004040433000008490040009c000000a10000213d00000005054002100000003f065000390000084a06600197000000400700043d0000000006670019002100000007001d000000000076004b00000000070000390000000107004039000008490060009c000000a10000213d0000000100700190000000a10000c13d000000400060043f0000002106000029000000000046043500000100022000390000000005250019000000000015004b0000009b0000213d000000000004004b000001710000613d00000021040000290000000026020434000008480060009c0000009b0000213d00000020044000390000000000640435000000000052004b0000016a0000413d000001600200043d000008490020009c0000009b0000213d0000001f04200039000000000034004b00000000050000190000084b050080410000084b04400197000000000004004b00000000060000190000084b060040410000084b0040009c000000000605c019000000000006004b0000009b0000c13d000000e0042000390000000004040433000008490040009c000000a10000213d00000005054002100000003f065000390000084a06600197000000400700043d0000000006670019002000000007001d000000000076004b00000000070000390000000107004039000008490060009c000000a10000213d0000000100700190000000a10000c13d000000400060043f00000020060000290000000006460436001c00000006001d00000100022000390000000005250019000000000015004b0000009b0000213d000000000004004b000001a30000613d00000020040000290000000026020434000008480060009c0000009b0000213d00000020044000390000000000640435000000000052004b0000019c0000413d000001800200043d000008490020009c0000009b0000213d0000001f04200039000000000034004b00000000050000190000084b050080410000084b04400197000000000004004b00000000060000190000084b060040410000084b0040009c000000000605c019000000000006004b0000009b0000c13d000000e0042000390000000005040433000008490050009c000000a10000213d00000005045002100000003f044000390000084a04400197000000400600043d0000000004460019001b00000006001d000000000064004b00000000060000390000000106004039000008490040009c000000a10000213d0000000100600190000000a10000c13d000000400040043f0000001b040000290000000004540436001a00000004001d000001000220003900000060045000c90000000004240019000000000014004b0000009b0000213d000000000005004b000011df0000c13d000001a00200043d000008490020009c0000009b0000213d0000001f04200039000000000034004b00000000050000190000084b050080410000084b04400197000000000004004b00000000060000190000084b060040410000084b0040009c000000000605c019000000000006004b0000009b0000c13d000000e004200039001f00000004001d0000000007040433000008490070009c000000a10000213d00000005067002100000003f046000390000084a05400197000000400400043d0000000005540019000000000045004b00000000080000390000000108004039000008490050009c000000a10000213d0000000100800190000000a10000c13d000000400050043f002c00000004001d00000000007404350000010005200039001e00000056001d0000001e0010006b0000009b0000213d000000000007004b000012d60000c13d000001c00200043d000008490020009c0000009b0000213d0000001f04200039000000000034004b00000000050000190000084b050040410000084b04400197000000000004004b00000000060000190000084b060020410000084b0040009c000000000605c019000000000006004b0000009b0000613d000000e0042000390000000005040433000008490050009c000000a10000213d00000005045002100000003f044000390000084a04400197000000400600043d0000000004460019001800000006001d000000000064004b00000000060000390000000106004039000008490040009c000000a10000213d0000000100600190000000a10000c13d000000400040043f0000001804000029002b00000004001d0000000000540435000001000220003900000006045002100000000004240019000000000014004b0000009b0000213d000000000005004b000015d30000c13d000001e00200043d000008490020009c0000009b0000213d0000001f04200039000000000034004b00000000030000190000084b030040410000084b04400197000000000004004b00000000050000190000084b050020410000084b0040009c000000000503c019000000000005004b0000009b0000613d000000e0032000390000000004030433000008490040009c000000a10000213d00000005034002100000003f033000390000084a03300197000000400500043d0000000003350019001700000005001d000000000053004b00000000050000390000000105004039000008490030009c000000a10000213d0000000100500190000000a10000c13d000000400030043f00000017030000290000000003430436001600000003001d000001000220003900000240034000c90000000003230019000000000013004b0000009b0000213d000000000004004b000015ec0000c13d000000400100043d001d00000001001d0000000001000411000000000001004b0000166f0000c13d0000001d0300002900000044013000390000088702000041000000000021043500000024013000390000001802000039000000000021043500000888010000410000000000130435000000040130003900000020020000390000000000210435000008430030009c0000084303008041000000400130021000000889011001c70000210a00010430000008a80020009c000002a00000213d000008ac0020009c000003d10000613d000008ad0020009c000003ab0000613d000008ae0020009c0000009b0000c13d0000000001000416000000000001004b0000009b0000c13d000000800000043f000000a00000043f000000c00000043f0000014001000039000000400010043f0000000001000412002a00000001001d002900000000003d0000800501000039000000440300003900000000040004150000002a0440008a0000000504400210000008d102000041210820e00000040f0000084701100197002400000001001d000000e00010043f0000000001000412002800000001001d002700200000003d0000000004000415000000280440008a00000005044002100000800501000039000008d1020000410000004403000039210820e00000040f0000084801100197002300000001001d000001000010043f0000000001000412002600000001001d002500400000003d0000000004000415000000260440008a00000005044002100000800501000039000008d1020000410000004403000039210820e00000040f0000084301100197000001200010043f0000002402000029000001400020043f0000002302000029000001600020043f000001800010043f000008e001000041000021090001042e000008a90020009c000003f50000613d000008aa0020009c000003bd0000613d000008ab0020009c0000009b0000c13d000000240030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000401100370000000000101043b001e00000001001d000008490010009c0000009b0000213d0000001e0130006a0000084c0010009c0000009b0000213d000000440010008c0000009b0000413d0000000001000411000000000010043f0000000301000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000101041a000000000001004b00000ab40000c13d000000400100043d000008dc02000041000000000021043500000004021000390000000003000411000006090000013d000008a50020009c000007fc0000613d000008a60020009c000007780000613d000008a70020009c0000009b0000c13d000000240030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000401100370000000000101043b000008480010009c0000009b0000213d21081f560000040f000000400200043d0000000000120435000008430020009c00000843020080410000004001200210000008c8011001c7000021090001042e0000089b0020009c000008630000613d0000089c0020009c000007b70000613d0000089d0020009c0000009b0000c13d0000000001000416000000000001004b0000009b0000c13d000000000100041a0000084801100197000000800010043f000008d601000041000021090001042e000008920020009c000009340000613d000008930020009c000007f10000613d000008940020009c0000009b0000c13d000000240030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000401100370000000000101043b000008480010009c0000009b0000213d21081c150000040f000000400200043d002400000002001d000004260000013d000008a10020009c0000042d0000613d000008a20020009c0000009b0000c13d0000000001000416000000000001004b0000009b0000c13d0000000101000039000000000201041a00000848032001970000000006000411000000000036004b00000a720000c13d000000000300041a0000085104300197000000000464019f000000000040041b0000085102200197000000000021041b00000000010004140000084805300197000008430010009c0000084301008041000000c0011002100000085a011001c70000800d020000390000000303000039000008d904000041210820fe0000040f00000001002001900000009b0000613d0000000001000019000021090001042e000008970020009c0000060f0000613d000008980020009c0000009b0000c13d000000840030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000402100370000000000202043b002400000002001d000008490020009c0000009b0000213d0000002402100370000000000202043b002300000002001d000008480020009c0000009b0000213d0000004402100370000000000202043b002200000002001d0000006402100370000000000202043b000008490020009c0000009b0000213d0000002304200039000000000034004b0000009b0000813d002000040020003d0000002001100360000000000101043b002100000001001d000008490010009c0000009b0000213d0000002101200029001f00240010003d0000001f0030006b0000009b0000213d000008d101000041000000000010044300000000010004120000000400100443000000200100003900000024001004430000000001000414000008430010009c0000084301008041000000c001100210000008d2011001c70000800502000039210821030000040f0000000100200190000015770000613d000000000201043b001e00000002001d000000230120014f00000848001001980000036e0000613d000000230100002921081f560000040f00230022001000bd000000220000006b000003690000613d000000230300002900000022023000fa000000000012004b000015cd0000c13d0000001e0100002921081f560000040f000000000001004b000003ef0000613d0022002300100101000008d10100004100000000001004430000000001000412000000040010044300000024000004430000000001000414000008430010009c0000084301008041000000c001100210000008d2011001c70000800502000039210821030000040f0000000100200190000015770000613d000000000101043b0000084701100197000000220010006b00000f200000a13d000000400200043d00000024032000390000000000130435000008d4010000410000000000120435000000040120003900000022030000290000000000310435000008430020009c0000084302008041000000400120021000000882011001c70000210a000104300000088e0020009c0000062f0000613d0000088f0020009c0000009b0000c13d000000440030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000402100370000000000302043b000008480030009c0000009b0000213d0000002401100370000000000201043b000008490020009c0000009b0000213d000000000103001921081d820000040f0000085002200197000000400300043d0000002004300039000000000024043500000850011001970000000000130435000008430030009c00000843030080410000004001300210000008af011001c7000021090001042e000000240030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000401100370000000000101043b000008480010009c0000009b0000213d000000000010043f0000000701000039000000200010043f210820d30000040f000000000101041a0000084901100197000000800010043f000008d601000041000021090001042e0000000001000416000000000001004b0000009b0000c13d0000000202000039000000000102041a000000800010043f000000000020043f000000000001004b000007fa0000613d000000a004000039000008dd0200004100000000030000190000000005040019000000000402041a000000000445043600000001022000390000000103300039000000000013004b000003c90000413d00000a860000013d000000640030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000402100370000000000202043b000008480020009c0000009b0000213d0000004401100370000000000101043b002400000001001d000008480010009c0000009b0000213d000000000102001921081f560000040f00000024020000390000000202200367000000000202043b00000000031200a9000000000002004b000003ea0000613d00000000022300d9000000000012004b000015cd0000c13d002300000003001d000000240100002921081f560000040f000000000001004b00000ab20000c13d000008d701000041000000000010043f0000001201000039000000040010043f00000877010000410000210a000104300000000001000416000000000001004b0000009b0000c13d000000c001000039000000400010043f0000001701000039000000800010043f000008de01000041000000a00010043f0000002001000039000000c00010043f0000008001000039000000e00200003921081bc60000040f000000c00110008a000008430010009c00000843010080410000006001100210000008df011001c7000021090001042e000000240030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000401100370000000000101043b000008490010009c0000009b0000213d000000c002000039000000400020043f000000800000043f000000a00000043f000000000010043f0000000401000039000000200010043f210820d30000040f002400000001001d000000c00100003921081b930000040f0000002401000029000000000101041a0000085002100197000000c00020043f000000e001100270000000e00010043f000000400200043d002400000002001d000000c00100003921081be80000040f0000002401000029000008430010009c00000843010080410000004001100210000008af011001c7000021090001042e000000440030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000402100370000000000202043b002300000002001d000008490020009c0000009b0000213d00000023020000290000002302200039000000000032004b0000009b0000813d00000023020000290000000402200039000000000221034f000000000202043b000008490020009c000000a10000213d00000005052002100000003f045000390000084a04400197000008b00040009c000000a10000213d0000008004400039000000400040043f000000800020043f00000023040000290000002404400039002200000045001d000000220030006b0000009b0000213d000000000002004b00000c6d0000c13d0000002402100370000000000202043b000008490020009c0000009b0000213d0000002304200039000000000034004b00000000050000190000084b050040410000084b04400197000000000004004b00000000060000190000084b060020410000084b0040009c000000000605c019000000000006004b0000009b0000613d0000000404200039000000000441034f000000000504043b000008490050009c000000a10000213d00000005045002100000003f044000390000084a04400197000000400600043d0000000004460019001e00000006001d000000000064004b00000000060000390000000106004039000008490040009c000000a10000213d0000000100600190000000a10000c13d000000400040043f0000001e040000290000000004540436001d00000004001d000000240220003900000006045002100000000004240019000000000034004b0000009b0000213d000000000005004b00000ff70000c13d000000000100041a00000848011001970000000002000411000000000012004b00000ba30000c13d000000800100043d000000000001004b000010a40000c13d0000001e010000290000000001010433000000000001004b000003230000613d002400000000001d000000240100002900000005011002100000001d01100029000000000101043300000020021000390000000002020433002300000002001d00000000010104330000084901100197002200000001001d000000000010043f0000000901000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d00000023020000290000084802200197000000000101043b002300000002001d000000000020043f000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000001041b0000000001000414000008430010009c0000084301008041000000c0011002100000085a011001c70000800d020000390000000303000039000008830400004100000022050000290000002306000029210820fe0000040f00000001002001900000009b0000613d0000002402000029002400010020003d0000001e010000290000000001010433000000240010006b0000048a0000413d000003230000013d000000240030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000402100370000000000202043b000008490020009c0000009b0000213d00000000042300490000084c0040009c0000009b0000213d000000440040008c0000009b0000413d000000c005000039000000400050043f0000000404200039000000000641034f000000000606043b000008490060009c0000009b0000213d00000000062600190000002307600039000000000037004b0000009b0000813d0000000407600039000000000771034f000000000807043b000008490080009c000000a10000213d00000005078002100000003f097000390000084a099001970000084e0090009c000000a10000213d000000c009900039000000400090043f000000c00080043f00000024066000390000000007670019000000000037004b0000009b0000213d000000000008004b000004fa0000613d000000000861034f000000000808043b000008480080009c0000009b0000213d000000200550003900000000008504350000002006600039000000000076004b000004f10000413d000000c005000039000000800050043f0000002004400039000000000441034f000000000404043b000008490040009c0000009b0000213d00000000022400190000002304200039000000000034004b00000000050000190000084b050080410000084b04400197000000000004004b00000000060000190000084b060040410000084b0040009c000000000605c019000000000006004b0000009b0000c13d0000000404200039000000000441034f000000000404043b000008490040009c000000a10000213d00000005054002100000003f065000390000084a06600197000000400700043d0000000006670019002200000007001d000000000076004b00000000070000390000000107004039000008490060009c000000a10000213d0000000100700190000000a10000c13d000000400060043f00000022060000290000000004460436002100000004001d00000024022000390000000004250019000000000034004b0000009b0000213d000000000042004b000005340000813d0000002203000029000000000521034f000000000505043b000008480050009c0000009b0000213d000000200330003900000000005304350000002002200039000000000042004b0000052b0000413d0000002201000029000000a00010043f000000000100041a00000848011001970000000002000411000000000012004b00000ba30000c13d00000022010000290000000001010433000000000001004b000011730000c13d000000800100043d002100000001001d0000000021010434002200000002001d000000000001004b000003230000613d002400000000001d00000024010000290000000501100210000000220110002900000000010104330000084801100198000017a00000613d002300000001001d000000000010043f0000000301000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000101041a000000000001004b000005780000c13d0000000201000039000000000101041a000008490010009c000000a10000213d00000001021000390000000203000039000000000023041b000008570110009a0000002302000029000000000021041b000000000103041a002000000001001d000000000020043f0000000301000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b0000002002000029000000000021041b000000400100043d00000023020000290000000000210435000008430010009c000008430100804100000040011002100000000002000414000008430020009c0000084302008041000000c002200210000000000112019f00000855011001c70000800d0200003900000001030000390000085804000041210820fe0000040f00000001002001900000009b0000613d0000002402000029002400010020003d00000021010000290000000001010433000000240010006b000005460000413d000003230000013d000000440030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000402100370000000000202043b002400000002001d000008490020009c0000009b0000213d0000002401100370000000000101043b000008490010009c0000009b0000213d00000000011300490000084c0010009c0000009b0000213d000000a40010008c0000009b0000413d0000002401000029000000000010043f0000000801000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000400500043d0000084f0050009c000000a10000213d000000000101043b0000022002500039000000400020043f000000000401041a000000d80240027000000843022001970000014003500039002300000003001d0000000000230435000000c8024002700000ffff0220018f0000012003500039002200000003001d0000000000230435000000b8024002700000ffff0220018f000001000f50003900000000002f0435000000a8024002700000ffff0220018f000000e00e50003900000000002e043500000088024002700000084302200197000000c00d50003900000000002d043500000078024002700000ffff0220018f000000a00b50003900000000002b043500000058024002700000084302200197000000800c50003900000000002c04350000003802400270000008430220019700000060095000390000000000290435000000180240027000000843022001970000004003500039000000000023043500000008024002700000ffff0620018f00000020025000390000000000620435000000ff064001900000000004000039000000010400c03900000000004504350000000101100039000000000101041a00000160075000390000084304100197002100000007001d000000000047043500000038041002100000086107400197000002000450003900000000007404350000085f001001980000000007000039000000010700c039000001e00a50003900000000007a043500000080071002700000084307700197000001c008500039002000000008001d000000000078043500000040071002700000084907700197000001a00850003900000000007804350000018005500039000000200110027000000843011001970000000000150435000000000006004b00000bb40000c13d000000400100043d000008cd020000410000000000210435000000040210003900000024030000290000000000320435000008430010009c0000084301008041000000400110021000000877011001c70000210a00010430000000240030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000401100370000000000101043b000008480010009c0000009b0000213d000000c002000039000000400020043f000000800000043f000000a00000043f000000000010043f0000000601000039000000200010043f210820d30000040f002400000001001d000000c00100003921081b930000040f0000002401000029000000000101041a0000084802100197000000c00020043f000000a001100270000000ff0110018f000000e00010043f000000400100043d002400000001001d000000c00200003921081c0d0000040f00000a8d0000013d000000240030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000402100370000000000502043b000008490050009c0000009b0000213d0000002302500039000000000032004b0000009b0000813d0000000402500039000000000221034f000000000202043b000008490020009c000000a10000213d00000005042002100000003f044000390000084a04400197000008b00040009c000000a10000213d0000008004400039000000400040043f000000800020043f00000240022000c900000024042001bf0000000002540019000000000032004b0000009b0000213d000000240040008c00000cff0000c13d000000000100041a00000848011001970000000002000411000000000012004b00000ba30000c13d000000800100043d000000000001004b000003230000613d002400000000001d000000240010006c000000240100002900001ad60000a13d0000000501100210000000a00110003900000000010104330000000012010434000008490420019800000f1a0000613d00000000030104330000018001300039002300000001001d0000000001010433000008430110019800000f1a0000613d0000020002300039002200000002001d00000000020204330000086102200197000008620020009c00000f1a0000c13d0000016002300039002100000002001d0000000002020433000008450020019800000f1a0000613d0000006002300039002000000002001d00000000020204330000084302200197000000000021004b00000f1a0000213d001e00000003001d000000000040043f0000000801000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039001f00000004001d210821030000040f00000001002001900000009b0000613d0000001f06000029000000000101043b0000000101100039000000000101041a0000001e050000290000000042050434000000000002004b0000000003000039000000010300c039000000400200043d0000000003320436001d00000004001d00000000040404330000ffff0440018f00000000004304350000004003500039001c00000003001d0000000003030433000008430330019700000040042000390000000000340435000000200300002900000000030304330000084303300197000000600420003900000000003404350000008003500039001b00000003001d0000000003030433000008430330019700000080042000390000000000340435000000a003500039001a00000003001d00000000030304330000ffff0330018f000000a0042000390000000000340435000000c003500039001900000003001d00000000030304330000084303300197000000c0042000390000000000340435000000e003500039001800000003001d00000000030304330000ffff0330018f000000e00420003900000000003404350000010003500039001700000003001d00000000030304330000ffff0330018f000001000420003900000000003404350000012003500039001500000003001d00000000030304330000ffff0330018f000001200420003900000000003404350000014003500039001300000003001d00000000030304330000084303300197000001400420003900000000003404350000002103000029000000000303043300000843033001970000016004200039000000000034043500000023030000290000000003030433000008430330019700000180042000390000000000340435000001a003500039001600000003001d00000000030304330000084903300197000001a0042000390000000000340435000001c003500039001400000003001d00000000030304330000084303300197000001c0042000390000000000340435000001e003500039001200000003001d0000000003030433000000000003004b0000000003000039000000010300c039000001e004200039000000000034043500000022030000290000000003030433000008610330019700000200042000390000000000340435000008430020009c000008430200804100000040022002100000086300100198000006fc0000613d0000000001000414000008430010009c0000084301008041000000c001100210000000000121019f000008b1011001c70000800d0200003900000002030000390000086404000041000007050000013d0000000001000414000008430010009c0000084301008041000000c001100210000000000121019f000008b1011001c70000800d02000039000000020300003900000865040000410000000005060019210820fe0000040f00000001002001900000009b0000613d0000001f01000029000000000010043f0000000801000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d0000001e020000290000000002020433000000000002004b000000000101043b000000000201041a0000086602200197000000010220c1bf0000001d03000029000000000303043300000008033002100000086703300197000000000232019f0000001c03000029000000000303043300000018033002100000086803300197000000000232019f0000002003000029000000000303043300000038033002100000086903300197000000000232019f0000001b03000029000000000303043300000058033002100000086a03300197000000000232019f0000001a03000029000000000303043300000078033002100000086b03300197000000000232019f0000001903000029000000000303043300000088033002100000086c03300197000000000232019f00000018030000290000000003030433000000a8033002100000086d03300197000000000232019f00000017030000290000000003030433000000b8033002100000086e03300197000000000232019f00000015030000290000000003030433000000c8033002100000086f03300197000000000232019f00000013030000290000000003030433000000d8033002100000087003300197000000000232019f000000000021041b0000000101100039000000210200002900000000020204330000084302200197000000000301041a0000087103300197000000000223019f0000002303000029000000000303043300000020033002100000087203300197000000000232019f0000001603000029000000000303043300000040033002100000087303300197000000000232019f0000001403000029000000000303043300000080033002100000087403300197000000000232019f00000012030000290000000003030433000000000003004b00000875030000410000000003006019000000000232019f0000002203000029000000000303043300000038033002700000086303300197000000000232019f000000000021041b0000002402000029002400010020003d000000800100043d000000240010006b000006580000413d000003230000013d000000240030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000402100370000000000202043b000008490020009c0000009b0000213d0000002304200039000000000034004b0000009b0000813d0000000404200039000000000141034f000000000101043b002100000001001d000008490010009c0000009b0000213d002000240020003d000000210100002900000005021002100000002001200029000000000031004b0000009b0000213d0000003f012000390000084a03100197000008b00030009c000000a10000213d0000008001300039000000400010043f0000002104000029000000800040043f000000000004004b00000d960000c13d00000020020000390000000002210436000000800300043d00000000003204350000004002100039000000000003004b000007ae0000613d000000a0040000390000000005000019000000004604043400000000760604340000085006600197000000000662043600000000070704330000084307700197000000000076043500000040022000390000000105500039000000000035004b000007a30000413d0000000002120049000008430020009c00000843020080410000006002200210000008430010009c00000843010080410000004001100210000000000112019f000021090001042e000000440030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000402100370000000000202043b000008490020009c0000009b0000213d0000002401100370000000000101043b002400000001001d000008480010009c0000009b0000213d0000014001000039000000400010043f000000800000043f000000a00000043f000000c00000043f000000e00000043f000001000000043f000001200000043f000000000020043f0000000901000039000000200010043f210820d30000040f0000002402000029000000000020043f000000200010043f210820d30000040f002400000001001d000001400100003921081ba90000040f0000002401000029000000000101041a0000084302100197000001400020043f00000020021002700000084302200197000001600020043f00000040021002700000ffff0220018f000001800020043f00000050021002700000084302200197000001a00020043f00000070021002700000084302200197000001c00020043f000008b8001001980000000001000039000000010100c039000001e00010043f000000400100043d002400000001001d000001400200003921081bef0000040f00000a8d0000013d0000000001000416000000000001004b0000009b0000c13d0000000a02000039000000000102041a000000800010043f000000000020043f000000000001004b00000a7c0000c13d000000200200003900000a870000013d000000240030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000402100370000000000202043b000008490020009c0000009b0000213d0000002304200039000000000034004b0000009b0000813d0000000404200039000000000441034f000000000504043b000008490050009c000000a10000213d00000005045002100000003f044000390000084a04400197000008b00040009c000000a10000213d0000008004400039000000400040043f000000800050043f000000240220003900000060045000c90000000004240019000000000034004b0000009b0000213d000000000005004b00000dc30000c13d000000000100041a00000848011001970000000002000411000000000012004b00000ba30000c13d000000800100043d000000000001004b000003230000613d002400000000001d00000024010000290000000501100210000000a001100039000000000101043300000020021000390000000002020433002200000002001d00000000010104330000084801100197002300000001001d000000000010043f0000000601000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000201041a0000085e02200197000000220500002900000020035000390000000004030433000000a0044002100000085f04400197000000000242019f00000000040504330000084804400197000000000242019f000000000021041b000000400100043d00000000024104360000000003030433000000ff0330018f0000000000320435000008430010009c000008430100804100000040011002100000000002000414000008430020009c0000084302008041000000c002200210000000000112019f00000853011001c70000800d02000039000000020300003900000860040000410000002305000029210820fe0000040f00000001002001900000009b0000613d0000002402000029002400010020003d000000800100043d000000240010006b000008250000413d000003230000013d000000440030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000402100370000000000202043b000008490020009c0000009b0000213d0000002305200039000000000035004b0000009b0000813d0000000405200039000000000551034f000000000605043b000008490060009c000000a10000213d00000005056002100000003f075000390000084a07700197000008b00070009c000000a10000213d0000008007700039000000400070043f000000800060043f00000024022000390000000005250019000000000035004b0000009b0000213d000000000006004b0000088b0000613d000000000621034f000000000606043b000008480060009c0000009b0000213d000000200440003900000000006404350000002002200039000000000052004b000008820000413d0000002402100370000000000202043b000008490020009c0000009b0000213d0000002304200039000000000034004b00000000050000190000084b050080410000084b04400197000000000004004b00000000060000190000084b060040410000084b0040009c000000000605c019000000000006004b0000009b0000c13d0000000404200039000000000441034f000000000404043b000008490040009c000000a10000213d00000005054002100000003f065000390000084a06600197000000400700043d0000000006670019002000000007001d000000000076004b00000000070000390000000107004039000008490060009c000000a10000213d0000000100700190000000a10000c13d000000400060043f00000020060000290000000006460436001f00000006001d00000024022000390000000005250019000000000035004b0000009b0000213d000000000004004b000008c10000613d0000002003000029000000000421034f000000000404043b000008480040009c0000009b0000213d000000200330003900000000004304350000002002200039000000000052004b000008b80000413d000000000100041a00000848011001970000000002000411000000000012004b00000ba30000c13d000000800100043d000000000001004b000010120000c13d00000020010000290000000001010433000000000001004b000003230000613d002400000000001d000008d50000013d0000002402000029002400010020003d00000020010000290000000001010433000000240010006b000003230000813d000000240100002900000005011002100000001f01100029002200000001001d00000000010104330000084801100197002300000001001d000000000010043f0000000b01000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000301041a000000000003004b000008cf0000613d0000000a01000039000000000201041a000000000002004b000015cd0000613d000000010130008a000000000023004b0000090e0000613d000000000012004b00001ad60000a13d0000085c0130009a0000085c0220009a000000000202041a000000000021041b000000000020043f0000000b01000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039002100000003001d210821030000040f000000210300002900000001002001900000009b0000613d000000000101043b000000000031041b0000000a01000039000000000301041a000000000003004b00001adc0000613d000000010130008a0000085c0230009a000000000002041b0000000a02000039000000000012041b0000002301000029000000000010043f0000000b01000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000001041b00000020010000290000000001010433000000240010006c00001ad60000a13d0000002201000029000000000101043300000000020004140000084805100197000008430020009c0000084302008041000000c0012002100000085a011001c70000800d0200003900000002030000390000085d04000041210820fe0000040f0000000100200190000008cf0000c13d0000009b0000013d000000640030008c0000009b0000413d0000000002000416000000000002004b0000009b0000c13d0000000402100370000000000202043b001b00000002001d000008490020009c0000009b0000213d0000002402100370000000000202043b002400000002001d000008490020009c0000009b0000213d00000024020000290000002302200039000000000032004b0000009b0000813d00000024020000290000000402200039000000000221034f000000000202043b002100000002001d000008490020009c0000009b0000213d00000024020000290000002405200039000000210200002900000005022002100000000002520019000000000032004b0000009b0000213d0000004402100370000000000202043b000008490020009c0000009b0000213d0000002304200039000000000034004b0000009b0000813d0000000404200039000000000141034f000000000101043b002000000001001d000008490010009c0000009b0000213d001f00240020003d000000200100002900000006011002100000001f01100029000000000031004b0000009b0000213d0000001b01000029000000000010043f0000000801000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039001a00000005001d210821030000040f0000001a0b00002900000001002001900000009b0000613d000000210000006b000003230000613d000000000101043b0000000101100039000000000101041a001e08630010019b000000200c00008a000000000d000019000009830000013d000000010dd000390000002100d0006c000003230000813d0000002000d0006c00001ad60000813d0000000601d002100000001f011000290000000202000367000000000112034f000000000101043b002300000001001d000008480010009c0000009b0000213d0000000501d002100022000000b1001d0000002201200360000000000101043b0000000003000031000000240430006a000000a30440008a0000084b054001970000084b06100197000000000756013f000000000056004b00000000050000190000084b05004041000000000041004b00000000060000190000084b060080410000084b0070009c000000000506c019000000000005004b0000009b0000c13d0000000005b100190000004006500039000000000662034f000000000606043b00000000075300490000001f0770008a0000084b087001970000084b09600197000000000a89013f000000000089004b00000000080000190000084b08004041000000000076004b00000000070000190000084b070080410000084b00a0009c000000000807c019000000000008004b0000009b0000c13d0000000005560019000000000652034f000000000906043b000008490090009c0000009b0000213d00000000069300490000002005500039000000000065004b00000000070000190000084b070020410000084b066001970000084b05500197000000000865013f000000000065004b00000000050000190000084b050040410000084b0080009c000000000507c019000000000005004b0000009b0000c13d000000210090008c000009f60000413d0000001b01000029000000000010043f0000000901000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039001d0000000d001d001c00000009001d210821030000040f00000001002001900000009b0000613d000000000101043b0000002302000029000000000020043f000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f0000001c030000290000001d0d000029000000200c00008a0000001a0b00002900000001002001900000009b0000613d000000000101043b000000000101041a00000070011002700000084301100197000000000013004b000011d90000213d0000000003000031000000240130006a00000002020003670000002205200360000000a30410008a000000000105043b000000000041004b00000000050000190000084b050080410000084b044001970000084b06100197000000000746013f000000000046004b00000000040000190000084b040040410000084b0070009c000000000405c019000000000004004b0000009b0000c13d0000000001b100190000002004100039000000000442034f000000000404043b00000000051300490000001f0550008a0000084b065001970000084b07400197000000000867013f000000000067004b00000000060000190000084b06004041000000000054004b00000000050000190000084b050080410000084b0080009c000000000605c019000000000006004b0000009b0000c13d0000000001140019000000000412034f000000000404043b000008490040009c0000009b0000213d000000000543004900000020061000390000084b015001970000084b07600197000000000817013f000000000017004b00000000010000190000084b01004041000000000056004b00000000050000190000084b050020410000084b0080009c000000000105c019000000000001004b0000009b0000c13d0000001f014000390000000001c1016f0000003f011000390000000005c1016f000000400100043d0000000005510019000000000015004b00000000070000390000000107004039000008490050009c000000a10000213d0000000100700190000000a10000c13d000000400050043f00000000054104360000000007640019000000000037004b0000009b0000213d000000000362034f0000000006c40170000000000265001900000a460000613d000000000703034f0000000008050019000000007907043c0000000008980436000000000028004b00000a420000c13d0000001f0740019000000a530000613d000000000363034f0000000306700210000000000702043300000000076701cf000000000767022f000000000303043b0000010006600089000000000363022f00000000036301cf000000000373019f0000000000320435000000000245001900000000000204350000001e02000029000008cf0020009c000009800000c13d0000000002010433000000200020008c00000a5f0000c13d0000000002050433000004000220008a000008b60020009c000009800000a13d000000400400043d002400000004001d000008b7020000410000000000240435000000040240003900000020030000390000000000320435000000240240003921081bc60000040f00000024020000290000000001210049000008430010009c0000084301008041000008430020009c000008430200804100000060011002100000004002200210000000000121019f0000210a000104300000088801000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f000008d801000041000000c40010043f000008b3010000410000210a00010430000000a004000039000008ce0200004100000000030000190000000005040019000000000402041a000000000445043600000001022000390000000103300039000000000013004b00000a7f0000413d000000600250008a000000800100003921081bb40000040f000000400100043d002400000001001d000000800200003921081bd80000040f00000024020000290000000001210049000008430010009c00000843010080410000006001100210000008430020009c00000843020080410000004002200210000000000121019f000021090001042e0000088801000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f000008b201000041000000c40010043f000008b3010000410000210a000104300000000101000039000000000201041a0000085102200197000000000262019f000000000021041b0000000001000414000008430010009c0000084301008041000000c0011002100000085a011001c70000800d020000390000000303000039000008b404000041210820fe0000040f00000001002001900000009b0000613d000003230000013d00000023011000f9000002dc0000013d0000001e04000029002400040040003d00000002020003670000002401200360000000000301043b00000000010000310000000004410049000000230440008a0000084b054001970000084b06300197000000000756013f000000000056004b00000000050000190000084b05004041000000000043004b00000000040000190000084b040080410000084b0070009c000000000504c019000000000005004b0000009b0000c13d0000002403300029000000000232034f000000000202043b001b00000002001d000008490020009c0000009b0000213d0000001b02000029000000060220021000000000012100490000002002300039000000000012004b00000000030000190000084b030020410000084b011001970000084b02200197000000000412013f000000000012004b00000000010000190000084b010040410000084b0040009c000000000103c019000000000001004b0000009b0000c13d000008be0100004100000000001004430000000001000414000008430010009c0000084301008041000000c001100210000008bf011001c70000800b02000039210821030000040f0000000100200190000015770000613d000000000101043b001c00000001001d001d08430010019b0000001b0000006b00000f720000c13d0000002401000029001b00200010003d00000002020003670000001b01200360000000000301043b00000000010000310000001e0410006a000000230440008a0000084b054001970000084b06300197000000000756013f000000000056004b00000000050000190000084b05004041000000000043004b00000000040000190000084b040080410000084b0070009c000000000504c019000000000005004b0000009b0000c13d0000002403300029000000000232034f000000000202043b001a00000002001d000008490020009c0000009b0000213d0000001a02000029000000060220021000000000012100490000002002300039000000000012004b00000000030000190000084b030020410000084b011001970000084b02200197000000000412013f000000000012004b00000000010000190000084b010040410000084b0040009c000000000103c019000000000001004b0000009b0000c13d0000001a0000006b000003230000613d000000000900001900000002010003670000001b02100360000000000302043b00000000020000310000001e0420006a000000230440008a0000084b054001970000084b06300197000000000756013f000000000056004b00000000050000190000084b05004041000000000043004b00000000040000190000084b040080410000084b0070009c000000000504c019000000000005004b0000009b0000c13d0000002404300029000000000341034f000000000303043b000008490030009c0000009b0000213d0000000605300210000000000552004900000020044000390000084b065001970000084b07400197000000000867013f000000000067004b00000000060000190000084b06004041000000000054004b00000000050000190000084b050020410000084b0080009c000000000605c019000000000006004b0000009b0000c13d000000000039004b00001ad60000813d0000000603900210000000000334001900000000023200490000084c0020009c0000009b0000213d000000400020008c0000009b0000413d000000400400043d0000084d0040009c000000a10000213d0000004002400039000000400020043f000000000231034f000000000202043b000008490020009c0000009b0000213d00000000052404360000002002300039000000000121034f000000000101043b000008500010009c0000009b0000213d002300000005001d0000000000150435000000400300043d0000084d0030009c000000a10000213d0000004002300039000000400020043f002000000003001d00000000021304360000001d01000029001f00000002001d000000000012043500000000010404330000084901100197000000000010043f0000000401000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039002100000009001d002200000004001d210821030000040f0000002204000029000000010020019000000023050000290000009b0000613d0000002002000029000000000202043300000850022001970000001f030000290000000003030433000000e003300210000000000223019f000000000101043b000000000021041b00000000020404330000000001050433000000400300043d00000020043000390000001c05000029000000000054043500000850011001970000000000130435000008430030009c000008430300804100000040013002100000000003000414000008430030009c0000084303008041000000c003300210000000000113019f00000853011001c700000849052001970000800d020000390000000203000039000008db04000041210820fe0000040f000000210900002900000001002001900000009b0000613d00000001099000390000001a0090006c00000b1f0000413d000003230000013d000000400100043d0000004402100039000008b203000041000000000032043500000024021000390000001603000039000000000032043500000888020000410000000000210435000000040210003900000020030000390000000000320435000008430010009c0000084301008041000000400110021000000889011001c70000210a00010430001f0000000f001d001b0000000e001d001c0000000d001d001e0000000c001d001d0000000b001d001700000005001d00180000000a001d001900000009001d001a00000008001d00000002050003670000002401500370000000000801043b0000004407800039000000000175034f000000000b01043b00000000060000310000000001860049000000230110008a0000084b091001970000084b0ab00197000000000c9a013f00000000009a004b000000000a0000190000084b0a00404100000000001b004b000000000d0000190000084b0d0080410000084b00c0009c000000000a0dc01900000000000a004b0000009b0000c13d000000040a8000390000000008ab0019000000000b85034f000000000b0b043b00160000000b001d0000084900b0009c0000009b0000213d000000160b000029000000060bb00210000000000bb6004900000020088000390000000000b8004b000000000c0000190000084b0c0020410000084b0bb001970000084b08800197000000000db8013f0000000000b8004b00000000080000190000084b080040410000084b00d0009c00000000080cc019000000000008004b0000009b0000c13d000000200870008a000000000785034f000000000707043b0000084b0b700197000000000c9b013f00000000009b004b000000000b0000190000084b0b004041000000000017004b000000000d0000190000084b0d0080410000084b00c0009c000000000b0dc01900000000000b004b0000009b0000c13d000000000ba700190000000007b5034f000000000707043b000008490070009c0000009b0000213d000000000c760049000000200bb000390000000000cb004b000000000d0000190000084b0d0020410000084b0cc001970000084b0bb00197000000000ecb013f0000000000cb004b000000000b0000190000084b0b0040410000084b00e0009c000000000b0dc01900000000000b004b0000009b0000c13d000000200880008a000000000b85034f000000000b0b043b0000084b0cb00197000000000d9c013f00000000009c004b00000000090000190000084b0900404100000000001b004b00000000010000190000084b010080410000084b00d0009c000000000901c019000000000009004b0000009b0000c13d0000000001ab0019000000000915034f000000000909043b000008490090009c0000009b0000213d000000000a960049000000200b1000390000084b01a001970000084b0cb00197000000000d1c013f00000000001c004b00000000010000190000084b010040410000000000ab004b000000000a0000190000084b0a0020410000084b00d0009c00000000010ac019000000000001004b0000009b0000c13d0000001f01900039000008e1011001970000003f01100039000008e10a100197000000400100043d000000000aa1001900000000001a004b000000000c000039000000010c0040390000084900a0009c000000a10000213d0000000100c00190000000a10000c13d0000004000a0043f000000000a910436000000000cb9001900000000006c004b0000009b0000213d0015000000b50353000008e10c9001980000001f0d90018f0000000006ca001900000c4e0000613d000000150e00035f000000000f0a001900000000eb0e043c000000000fbf043600000000006f004b00000c4a0000c13d00000000000d004b00000c5b0000613d000000150bc0035f000000030cd00210000000000d060433000000000dcd01cf000000000dcd022f000000000b0b043b000001000cc00089000000000bcb022f000000000bcb01cf000000000bdb019f0000000000b6043500000000069a0019000000000006043500000000030304330000084303300197000000000037004b000012160000a13d000000400100043d00000024021000390000000000720435000008cc02000041000000000021043500000004021000390000000000320435000008430010009c0000084301008041000000400110021000000882011001c70000210a00010430000000a006000039002100240030009200000c760000013d00000024020000290000000000a2043500000000068604360000002004400039000000220040006c000004500000813d000000000241034f000000000202043b000008490020009c0000009b0000213d000000230220002900000021052000690000084c0050009c0000009b0000213d000000400050008c0000009b0000413d000000400800043d0000084d0080009c000000a10000213d0000004005800039000000400050043f0000002405200039000000000751034f000000000707043b000008490070009c0000009b0000213d0000000007780436002400000007001d0000002005500039000000000551034f000000000505043b000008490050009c0000009b0000213d00000000022500190000004305200039000000000035004b00000000070000190000084b070080410000084b05500197000000000005004b00000000090000190000084b090040410000084b0050009c000000000907c019000000000009004b0000009b0000c13d0000002405200039000000000551034f000000000c05043b0000084900c0009c000000a10000213d0000000505c002100000003f055000390000084a05500197000000400a00043d000000000b5a00190000000000ab004b000000000500003900000001050040390000084900b0009c000000a10000213d0000000100500190000000a10000c13d0000004000b0043f0000000000ca0435000000440b200039000000e002c000c9000000000cb2001900000000003c004b0000009b0000213d0000000000cb004b00000c700000813d000000000d0a00190000000002b300490000084c0020009c0000009b0000213d000000e00020008c0000009b0000413d000000400e00043d0000084d00e0009c000000a10000213d0000004002e00039000000400020043f0000000002b1034f000000000202043b000008480020009c0000009b0000213d000000000f2e0436000000400200043d0000084e0020009c000000a10000213d000000c005200039000000400050043f0000002005b00039000000000751034f000000000707043b000008430070009c0000009b0000213d00000000077204360000002005500039000000000951034f000000000909043b000008430090009c0000009b0000213d00000000009704350000002005500039000000000751034f000000000707043b0000ffff0070008c0000009b0000213d000000400920003900000000007904350000002005500039000000000751034f000000000707043b000008430070009c0000009b0000213d000000600920003900000000007904350000002005500039000000000751034f000000000707043b000008430070009c0000009b0000213d000000800920003900000000007904350000002005500039000000000551034f000000000505043b000000000005004b0000000007000039000000010700c039000000000075004b0000009b0000c13d000000200dd00039000000a007200039000000000057043500000000002f04350000000000ed0435000000e00bb000390000000000cb004b00000cb90000413d00000c700000013d000000a004000039000000240550003900000000065300490000084c0060009c0000009b0000213d000002400060008c0000009b0000413d000000400600043d0000084d0060009c000000a10000213d0000004007600039000000400070043f000000000751034f000000000707043b000008490070009c0000009b0000213d0000000007760436000000400800043d0000084f0080009c000000a10000213d0000022009800039000000400090043f0000002009500039000000000a91034f000000000a0a043b00000000000a004b000000000b000039000000010b00c0390000000000ba004b0000009b0000c13d000000000aa804360000002009900039000000000b91034f000000000b0b043b0000ffff00b0008c0000009b0000213d0000000000ba04350000002009900039000000000a91034f000000000a0a043b0000084300a0009c0000009b0000213d000000400b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000084300a0009c0000009b0000213d000000600b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000084300a0009c0000009b0000213d000000800b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000ffff00a0008c0000009b0000213d000000a00b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000084300a0009c0000009b0000213d000000c00b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000ffff00a0008c0000009b0000213d000000e00b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000ffff00a0008c0000009b0000213d000001000b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000ffff00a0008c0000009b0000213d000001200b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000084300a0009c0000009b0000213d000001400b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000084300a0009c0000009b0000213d000001600b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000084300a0009c0000009b0000213d000001800b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000084900a0009c0000009b0000213d000001a00b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000084300a0009c0000009b0000213d000001c00b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b00000000000a004b000000000b000039000000010b00c0390000000000ba004b0000009b0000c13d000001e00b8000390000000000ab04350000002009900039000000000991034f000000000909043b00000850009001980000009b0000c13d000002000a80003900000000009a0435000000000087043500000000046404360000024005500039000000000025004b00000d010000413d0000064f0000013d0000084e0030009c000000a10000213d00000000030000190000004004100039000000400040043f000000200410003900000000000404350000000000010435000000a00430003900000000001404350000002003300039000000000023004b00000dea0000813d000000400100043d0000084d0010009c00000d990000a13d000000a10000013d000000a004000039000000240550003900000000065300490000084c0060009c0000009b0000213d000000400060008c0000009b0000413d000000400600043d0000084d0060009c000000a10000213d0000004007600039000000400070043f000000000751034f000000000707043b000008480070009c0000009b0000213d00000000077604360000002008500039000000000881034f000000000808043b000008490080009c0000009b0000213d000000000087043500000000046404360000004005500039000000000025004b00000da90000413d0000003b0000013d000000a00500003900000000062300490000084c0060009c0000009b0000213d000000600060008c0000009b0000413d000000400600043d0000084d0060009c000000a10000213d0000004007600039000000400070043f000000000721034f000000000707043b000008480070009c0000009b0000213d0000000007760436000000400800043d0000084d0080009c000000a10000213d0000004009800039000000400090043f0000002009200039000000000a91034f000000000a0a043b0000084800a0009c0000009b0000213d000000000aa804360000002009900039000000000991034f000000000909043b000000ff0090008c0000009b0000213d00000000009a0435000000000087043500000000056504360000006002200039000000000042004b00000dc40000413d0000081c0000013d0000000002000019002300000002001d0000000502200210002200000002001d00000020012000290000000201100367000000000101043b002400000001001d000008480010009c0000009b0000213d000000400100043d0000084d0010009c000000a10000213d0000004002100039000000400020043f0000002002100039000000000002043500000000000104350000002401000029000000000010043f0000000601000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000400200043d0000084d0020009c000000a10000213d000000000101043b0000004003200039000000400030043f000000000101041a00000848031001980000000003320436000000a001100270000000ff0110018f000000000013043500000e2c0000613d002400000003001d000000400100043d0000084d0010009c000000a10000213d0000004003100039000000400030043f0000002003100039000000000003043500000000000104350000000002020433000000400b00043d000008ba0100004100000000051b043600000000010004140000084802200197000000040020008c00000e440000c13d0000000103000031000000a00030008c000000a004000039000000000403401900000e730000013d0000002401000029000000000010043f0000000501000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000400300043d0000084d0030009c000000a10000213d000000000101043b0000004002300039000000400020043f000000000101041a00000850021001970000000000230435000000e00110027000000f060000013d001d00000005001d0000084300b0009c000008430300004100000000030b40190000004003300210000008430010009c0000084301008041000000c001100210000000000131019f00000886011001c7001e00000002001d001f0000000b001d210821030000040f0000001f0b000029000000000301001900000060033002700000084303300197000000a00030008c000000a0040000390000000004034019000000e00640019000000000056b001900000e610000613d000000000701034f00000000080b0019000000007907043c0000000008980436000000000058004b00000e5d0000c13d0000001f0740019000000e6e0000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f0000000100200190000011410000613d0000001e020000290000001d050000290000001f01400039000001e00110018f000000000ab1001900000000001a004b000000000100003900000001010040390000084900a0009c000000a10000213d0000000100100190000000a10000c13d0000004000a0043f000000a00030008c0000009b0000413d00000000010b0433000008bb0010009c0000009b0000213d0000008001b000390000000001010433000008bb0010009c0000009b0000213d00000000060504330000084c0060009c000011330000213d000008bc0100004100000000001a04350000000001000414000000040020008c000000200400003900000ebd0000613d001e00000006001d0000084300a0009c000008430300004100000000030a40190000004003300210000008430010009c0000084301008041000000c001100210000000000131019f00000886011001c7001f0000000a001d210821030000040f0000001f0a000029000000000301001900000060033002700000084303300197000000200030008c00000020040000390000000004034019000000200640019000000000056a001900000eac0000613d000000000701034f00000000080a0019000000007907043c0000000008980436000000000058004b00000ea80000c13d0000001f0740019000000eb90000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f00000001002001900000114d0000613d0000001e060000290000001f01400039000000600110018f0000000005a10019000008490050009c000000a10000213d000000400050043f000000200030008c0000009b0000413d00000000010a0433000000ff0010008c0000009b0000213d00000024020000290000000002020433000000ff0220018f0000000002120019000000ff0020008c000015cd0000213d000000240120008c00000ee20000213d000000010100003900000edb0000613d00000024022000890000000a03000039000000010020019000000000043300a9000000010300603900000000011300a90000000102200272000000000304001900000ed40000c13d000000000006004b00000f160000613d00000000026100a900000000036200d9000000000013004b00000ef00000613d000015cd0000013d0000004d0010008c000015cd0000213d00000001020000390000000a03000039000000010010019000000000043300a9000000010300603900000000022300a90000000101100272000000000304001900000ee60000c13d000000000002004b000003ef0000613d00000000022600d9000008500020009c0000113a0000213d0000084d0050009c000000a10000213d0000004001500039000000400010043f002400000005001d0000000000250435000008be0100004100000000001004430000000001000414000008430010009c0000084301008041000000c001100210000008bf011001c70000800b02000039210821030000040f0000000100200190000015770000613d000000000101043b0000084301100197000000240300002900000020023000390000000000120435000000800100043d0000002302000029000000000021004b00001ad60000a13d0000002201000029000000a0011000390000000000310435000000800100043d000000000021004b00001ad60000a13d0000000102200039000000210020006c00000deb0000413d0000105c0000013d00000000020000190000084d0050009c00000ef40000a13d000000a10000013d000000400100043d00000876020000410000000000210435000000040210003900000000004204350000060a0000013d0000002401000029000000000010043f0000000801000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000400200043d0000084d0020009c000000a10000213d0000000101100039000000000101041a0000004003200039000000400030043f000000200320003900000000000304350000000000020435000000210000006b0000105e0000c13d000000400300043d0000084d0030009c000000a10000213d000000200110027000000843011001970000004002300039000000400020043f00000000001304350000000001000019000000000200001900000020043000390000000000240435000000400600043d0000002002600039000008c60500004100000000005204350000000002030433000000240360003900000000002304350000000002040433000000000002004b0000000002000039000000010200c03900000044036000390000000000230435000000440200003900000000002604350000000005060019000008b00060009c000000a10000213d002300000005001d0000008004500039002400000004001d000000400040043f000000c00250003900000060030000390000000000320435000000010110018f000000a002500039000000000012043500000022010000290000000000140435000000e002500039000000000105001921081bc60000040f000000230110006a000000800110008a0000002402000029000008430020009c0000084302008041000008430010009c000008430100804100000040022002100000006001100210000000000121019f000021090001042e000000000900001900000002010003670000002402100360000000000302043b00000000020000310000001e0420006a000000230440008a0000084b054001970000084b06300197000000000756013f000000000056004b00000000050000190000084b05004041000000000043004b00000000040000190000084b040080410000084b0070009c000000000504c019000000000005004b0000009b0000c13d0000002404300029000000000341034f000000000303043b000008490030009c0000009b0000213d0000000605300210000000000552004900000020044000390000084b065001970000084b07400197000000000867013f000000000067004b00000000060000190000084b06004041000000000054004b00000000050000190000084b050020410000084b0080009c000000000605c019000000000006004b0000009b0000c13d000000000039004b00001ad60000813d0000000603900210000000000334001900000000023200490000084c0020009c0000009b0000213d000000400020008c0000009b0000413d000000400400043d0000084d0040009c000000a10000213d0000004002400039000000400020043f000000000231034f000000000202043b000008480020009c0000009b0000213d00000000052404360000002002300039000000000121034f000000000101043b000008500010009c0000009b0000213d002300000005001d0000000000150435000000400300043d0000084d0030009c000000a10000213d0000004002300039000000400020043f002000000003001d00000000021304360000001d01000029001f00000002001d000000000012043500000000010404330000084801100197000000000010043f0000000501000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039002100000009001d002200000004001d210821030000040f0000002204000029000000010020019000000023050000290000009b0000613d0000002002000029000000000202043300000850022001970000001f030000290000000003030433000000e003300210000000000223019f000000000101043b000000000021041b00000000020404330000000001050433000000400300043d00000020043000390000001c05000029000000000054043500000850011001970000000000130435000008430030009c000008430300804100000040013002100000000003000414000008430030009c0000084303008041000000c003300210000000000113019f00000853011001c700000848052001970000800d020000390000000203000039000008da04000041210820fe0000040f000000210900002900000001002001900000009b0000613d00000001099000390000001b0090006c00000f730000413d00000af00000013d0000001d0500002900000000062300490000084c0060009c0000009b0000213d000000400060008c0000009b0000413d000000400600043d0000084d0060009c000000a10000213d0000004007600039000000400070043f000000000721034f000000000707043b000008490070009c0000009b0000213d00000000077604360000002008200039000000000881034f000000000808043b000008480080009c0000009b0000213d000000000087043500000000056504360000004002200039000000000042004b00000ff80000413d0000047d0000013d0000000002000019000010190000013d00000024020000290000000102200039000000800100043d000000000012004b000008c90000813d002400000002001d0000000501200210000000a001100039002200000001001d00000000010104330000084801100197002300000001001d000000000010043f0000000b01000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000101041a000000000001004b000010140000c13d0000000a03000039000000000103041a000008490010009c000000a10000213d0000000102100039000000000023041b000008590110009a0000002302000029000000000021041b000000000103041a002100000001001d000000000020043f0000000b01000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b0000002102000029000000000021041b000000800100043d000000240010006c00001ad60000a13d0000002201000029000000000101043300000000020004140000084805100197000008430020009c0000084302008041000000c0012002100000085a011001c70000800d0200003900000002030000390000085b04000041210820fe0000040f0000000100200190000010140000c13d0000009b0000013d000000400100043d0000079a0000013d000000200100002900000020051000390000000204000367000000000154034f000000000301043b0000002101000029000000040610008c0000009b0000413d00000021010000290000001b01100039000008e1011001970000003f01100039000008e102100197000000400100043d0000000002210019000000000012004b00000000080000390000000108004039000008490020009c000000a10000213d0000000100800190000000a10000c13d000000400020043f00000000026104360000001f09000029000000000090007c0000009b0000213d0000000405500039000000000554034f000008e1076001980000001f0660018f0000000004720019000010850000613d000000000805034f0000000009020019000000008a08043c0000000009a90436000000000049004b000010810000c13d0000086103300197000000000006004b000010930000613d000000000575034f0000000306600210000000000704043300000000076701cf000000000767022f000000000505043b0000010006600089000000000565022f00000000056501cf000000000575019f000000000054043500000021041000290000001c044000390000000000040435000008c60030009c000012020000613d000008c70030009c000016e40000c13d00000000010104330000084c0010009c0000009b0000213d000000200010008c0000009b0000413d000000400300043d0000084d0030009c000000a10000213d000000000102043300000f3f0000013d0000000005000019000010a90000013d0000000105500039000000000015004b000004850000813d0000000502500210000000a00220003900000000030204330000002002300039001f00000002001d00000000020204330000000004020433000000000004004b000010a60000613d001c00000005001d0000000001030433002308490010019b0000000003000019002000000003001d00000005013002100000000001120019000000200110003900000000010104330000000021010434002408480010019b0000000001020433002100000001001d0000008001100039002200000001001d000000000101043300000843011001970000001f0010008c0000116b0000a13d0000002301000029000000000010043f0000000901000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f000000010020019000000024030000290000009b0000613d000000000101043b000000000030043f000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f000000240600002900000001002001900000009b0000613d000000000101043b000000000201041a0000087a02200197000000210a0000290000002003a00039000000000403043300000020044002100000087204400197000000000242019f0000004004a00039000000000504043300000040055002100000087b05500197000000000252019f0000006005a00039000000000705043300000050077002100000087c07700197000000000272019f0000002209000029000000000709043300000070077002100000087d07700197000000000272019f000000a008a000390000000007080433000000000007004b0000087e070000410000000007006019000000000272019f00000000070a04330000084307700197000000000272019f000000000021041b000000400100043d000000000271043600000000030304330000084303300197000000000032043500000000020404330000ffff0220018f00000040031000390000000000230435000000000205043300000843022001970000006003100039000000000023043500000000020904330000084302200197000000800310003900000000002304350000000002080433000000000002004b0000000002000039000000010200c039000000a0031000390000000000230435000008430010009c000008430100804100000040011002100000000002000414000008430020009c0000084302008041000000c002200210000000000112019f0000087f011001c70000800d02000039000000030300003900000880040000410000002305000029210820fe0000040f00000001002001900000009b0000613d000000200300002900000001033000390000001f0100002900000000020104330000000001020433000000000013004b000010b60000413d000000800100043d0000001c05000029000010a60000013d000008bd0100004100000000001a04350000084300a0009c000008430a0080410000004001a0021000000886011001c70000210a00010430000008bd010000410000000000150435000008430050009c0000084305008041000000400150021000000886011001c70000210a000104300000001f0530018f0000084506300198000000400200043d0000000004620019000011580000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000011480000c13d000011580000013d0000001f0530018f0000084506300198000000400200043d0000000004620019000011580000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000011540000c13d000000000005004b000011650000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000008430020009c00000843020080410000004002200210000000000112019f0000210a00010430000000400200043d000000240320003900000000001304350000088101000041000000000012043500000004012000390000002403000029000003870000013d00000000020000190000117b0000013d0000002302000029000000010220003900000022010000290000000001010433000000000012004b0000053f0000813d002300000002001d0000000501200210000000210110002900000000010104330000084801100197002400000001001d000000000010043f0000000301000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000301041a000000000003004b000011750000613d0000000201000039000000000201041a000000000002004b000015cd0000613d000000010130008a000000000032004b000011b30000613d000000000012004b00001ad60000a13d000008540130009a000008540220009a000000000202041a000000000021041b000000000020043f0000000301000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039002000000003001d210821030000040f000000200300002900000001002001900000009b0000613d000000000101043b000000000031041b0000000201000039000000000301041a000000000003004b00001adc0000613d000000010130008a000008540230009a000000000002041b0000000202000039000000000012041b0000002401000029000000000010043f0000000301000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000001041b000000400100043d00000024020000290000000000210435000008430010009c000008430100804100000040011002100000000002000414000008430020009c0000084302008041000000c002200210000000000112019f00000855011001c70000800d0200003900000001030000390000085604000041210820fe0000040f0000000100200190000011750000c13d0000009b0000013d000000400100043d000008d002000041000000000021043500000004021000390000002303000029000006090000013d0000001a0500002900000000062100490000084c0060009c0000009b0000213d000000600060008c0000009b0000413d000000400600043d0000084d0060009c000000a10000213d0000004007600039000000400070043f0000000097020434000008480070009c0000009b0000213d0000000007760436000000400800043d0000084d0080009c000000a10000213d000000400a8000390000004000a0043f0000000009090433000008480090009c0000009b0000213d0000000009980436000000400a200039000000000a0a0433000000ff00a0008c0000009b0000213d0000000000a90435000000000087043500000000056504360000006002200039000000000042004b000011e00000413d000001ce0000013d00000000030104330000084c0030009c0000009b0000213d000000400030008c0000009b0000413d000000400300043d0000084d0030009c000000a10000213d0000004004300039000000400040043f0000000002020433000000000023043500000040011000390000000002010433000000000002004b0000000001000039000000010100c039000000000012004b00000f440000613d0000009b0000013d00000000020204330000ffff0220018f000000160020006b0000121d0000a13d000000400100043d000008cb02000041000017a20000013d00000000020404330000086102200197000008620020009c000012280000c13d0000000002010433000000200020008c00000a5f0000c13d00000000020a0433000004000220008a000008b60020009c00000a5f0000213d0000006001800039000000000115034f000000000101043b000008480010009c0000009b0000213d000000000010043f0000000701000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d00000002020003670000002403200370000000000303043b0000006403300039000000000232034f000000000101043b000000000101041a000200000001001d000000000102043b000008480010009c0000009b0000213d000000240200002921081d820000040f000700000001001d000100000002001d000000160000006b000013590000c13d000000200100002900000000010104330000084301100197001508b9001000d5001100000000001d001200000000001d00000002020003670000002401200370000000000701043b0000002401700039000000000112034f00000000040000310000000003740049000000230630008a000000000501043b0000001f0100002900000000010104330000ffff031001900000000001000019000015780000c13d000000000065004b00000000030000190000084b030080410000084b086001970000084b09500197000000000a89013f000000000089004b00000000090000190000084b090040410000084b00a0009c000000000903c019000000000009004b0000009b0000c13d0000001e03000029000000000303043300000004097000390000000005950019000000000a52034f000000000a0a043b0000084900a0009c0000009b0000213d000000000ba4004900000020055000390000000000b5004b000000000c0000190000084b0c0020410000084b0bb001970000084b05500197000000000db5013f0000000000b5004b00000000050000190000084b050040410000084b00d0009c00000000050cc019000000000005004b0000009b0000c13d0000001d0500002900000000050504330000ffff0b50018f0000000005ab00a900000000000a004b0000128d0000613d000008c50aa00197000008c50c500197000000000aac00d90000000000ab004b000015cd0000c13d0000008407700039000000000772034f000000000707043b0000084b0a700197000000000b8a013f00000000008a004b00000000080000190000084b08004041000000000067004b00000000060000190000084b060080410000084b00b0009c000000000806c019000000000008004b0000009b0000c13d0000000007970019000000000672034f000000000606043b000008490060009c0000009b0000213d000000000864004900000020097000390000084b078001970000084b0a900197000000000b7a013f00000000007a004b00000000070000190000084b07004041000000000089004b00000000080000190000084b080020410000084b00b0009c000000000708c019000000000007004b0000009b0000c13d000000400700043d0000084d0070009c000000a10000213d0000004008700039000000400080043f000000200870003900000000000804350000000000070435000000400800043d0000084d0080009c000000a10000213d00000017070000290000000007070433000000400a8000390000004000a0043f000000200a80003900000000000a04350000000000080435000000000006004b000016a30000c13d000000400400043d0000084d0040009c000000a10000213d00000843027001970000004006400039000000400060043f0000000000240435000000000600001900000020044000390000000000640435000000190400002900000000040404330000084304400197000000000042004b000017b30000a13d000000400100043d000008ca02000041000017a20000013d001d00c00030003d000012dd0000013d00000020044000390000000000a9043500000000008404350000001e0050006c000001f70000813d0000000052050434000008490020009c0000009b0000213d0000001f022000290000001d062000690000084c0060009c0000009b0000213d000000400060008c0000009b0000413d000000400800043d0000084d0080009c000000a10000213d0000004006800039000000400060043f00000020062000390000000006060433000008490060009c0000009b0000213d000000000968043600000040062000390000000006060433000008490060009c0000009b0000213d00000000022600190000003f06200039000000000016004b00000000070000190000084b070080410000084b06600197000000000006004b000000000a0000190000084b0a0040410000084b0060009c000000000a07c01900000000000a004b0000009b0000c13d0000002006200039000000000c0604330000084900c0009c000000a10000213d0000000506c002100000003f066000390000084a06600197000000400a00043d00000000066a00190000000000a6004b00000000070000390000000107004039000008490060009c000000a10000213d0000000100700190000000a10000c13d000000400060043f0000000000ca0435000000400b200039000000e002c000c9000000000cb2001900000000001c004b0000009b0000213d0000000000cb004b000012d80000813d000000000d0a00190000000002b100490000084c0020009c0000009b0000213d000000e00020008c0000009b0000413d000000400e00043d0000084d00e0009c000000a10000213d0000004002e00039000000400020043f00000000620b0434000008480020009c0000009b0000213d000000000f2e0436000000400200043d0000084e0020009c000000a10000213d000000c007200039000000400070043f0000000006060433000008430060009c0000009b0000213d00000000066204360000004007b000390000000007070433000008430070009c0000009b0000213d00000000007604350000006006b0003900000000060604330000ffff0060008c0000009b0000213d000000400720003900000000006704350000008006b000390000000006060433000008430060009c0000009b0000213d00000060072000390000000000670435000000a006b000390000000006060433000008430060009c0000009b0000213d00000080072000390000000000670435000000c006b000390000000006060433000000000006004b0000000007000039000000010700c039000000000076004b0000009b0000c13d000000200dd00039000000a007200039000000000067043500000000002f04350000000000ed0435000000e00bb000390000000000cb004b0000131b0000413d000012d80000013d00000002010003670000002402100370000000000202043b0000006403200039000000000431034f000000000404043b000600000004001d000008480040009c0000009b0000213d000000200330008a000000000331034f000000000403043b00000000030000310000000005230049000000230550008a0000084b065001970000084b07400197000000000867013f000000000067004b00000000060000190000084b06004041000000000054004b00000000050000190000084b050080410000084b0080009c000000000605c019000000000006004b0000009b0000c13d00000000024200190000000402200039000000000121034f000000000101043b000b00000001001d000008490010009c0000009b0000213d0000000b010000290000000601100210000000000113004900000020052000390000084b021001970000084b03500197000000000423013f000000000023004b00000000020000190000084b02004041000a00000005001d000000000015004b00000000010000190000084b010020410000084b0040009c000000000201c019000000000002004b0000009b0000c13d0000000b0000006b001100000000001d001200000000001d001500000000001d000012500000613d001400000000001d001500000000001d001100000000001d001200000000001d000013a00000013d000000150010002a000015cd0000413d001500150010002d00000014020000290000000102200039001400000002001d0000000b0020006c000012500000813d000000140100002900000006011002100000000a0110002900000000021000790000084c0020009c0000009b0000213d000000400020008c0000009b0000413d000000400200043d002000000002001d0000084d0020009c000000a10000213d00000020020000290000004002200039000000400020043f0000000202000367000000000312034f000000000303043b000008480030009c0000009b0000213d000000200400002900000000033404360000002001100039000000000112034f000000000101043b000f00000003001d00000000001304350000002401000029000000000010043f0000000901000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000200200002900000000020204330000084802200197000000000020043f000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000400200043d001300000002001d0000084e0020009c000000a10000213d000000000101043b0000001304000029000000c002400039000000400020043f000000000101041a000000a002400039000008b8001001980000000003000039000000010300c039000000000032043500000843021001970000000003240436000000700210027000000843022001970000008005400039001000000005001d000000000025043500000020021002700000084302200197000c00000003001d0000000000230435000000500210027000000843022001970000006003400039000e00000003001d0000000000230435000000400340003900000040011002700000ffff0110018f000d00000003001d00000000001304350000143f0000613d000000000001004b000014580000613d000000200100002900000000010104330000084802100197000800000002001d000000060020006c0000000701000029000015420000613d000000400100043d0000084d0010009c000000a10000213d0000004002100039000000400020043f0000002002100039000000000002043500000000000104350000000801000029000000000010043f0000000601000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000400200043d0000084d0020009c000000a10000213d000000000101043b0000004003200039000000400030043f000000000101041a00000848031001980000000003320436000000a001100270000000ff0110018f000400000003001d00000000001304350000145a0000613d000000400100043d0000084d0010009c000000a10000213d0000004003100039000000400030043f0000002003100039000000000003043500000000000104350000000002020433000000400300043d000008ba01000041002000000003001d0000000001130436000300000001001d00000000010004140000084802200197000500000002001d000000040020008c000014740000c13d0000000103000031000000a00030008c000000a00400003900000000040340190000149e0000013d000000220100002900000000010104330000ffff0110018f000008b9011000d1000000150010002a000015cd0000413d000000110200002900000843022001970000002303000029000000000303043300000843033001970000000002230019001100000002001d000008430020009c000015cd0000213d000000120200002900000843022001970000002103000029000000000303043300000843033001970000000002230019001200000002001d000008430020009c0000139a0000a13d000015cd0000013d0000000001000019000015520000013d0000000801000029000000000010043f0000000501000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000400200043d002000000002001d0000084d0020009c000000a10000213d000000000101043b00000020030000290000004002300039000000400020043f000000000101041a00000850021001970000000000230435000000e001100270000015390000013d0000002002000029000008430020009c00000843020080410000004002200210000008430010009c0000084301008041000000c001100210000000000121019f00000886011001c70000000502000029210821030000040f000000000301001900000060033002700000084303300197000000a00030008c000000a0040000390000000004034019000000e00640019000000020056000290000148e0000613d000000000701034f0000002008000029000000007907043c0000000008980436000000000058004b0000148a0000c13d0000001f074001900000149b0000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f0000000100200190000017bc0000613d0000001f01400039000001e00110018f0000002002100029000000000012004b00000000010000390000000101004039000900000002001d000008490020009c000000a10000213d0000000100100190000000a10000c13d0000000901000029000000400010043f000000a00030008c0000009b0000413d00000020010000290000000001010433000008bb0010009c0000009b0000213d000000200100002900000080011000390000000001010433000008bb0010009c0000009b0000213d00000003010000290000000001010433000300000001001d0000084c0010009c000017a80000213d000008bc010000410000000902000029000000000012043500000000010004140000000502000029000000040020008c0000002004000039000014ed0000613d0000000902000029000008430020009c00000843020080410000004002200210000008430010009c0000084301008041000000c001100210000000000121019f00000886011001c70000000502000029210821030000040f000000000301001900000060033002700000084303300197000000200030008c0000002004000039000000000403401900000020064001900000000905600029000014dd0000613d000000000701034f0000000908000029000000007907043c0000000008980436000000000058004b000014d90000c13d0000001f07400190000014ea0000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f0000000100200190000017c80000613d0000001f01400039000000600110018f0000000901100029002000000001001d000008490010009c000000a10000213d0000002001000029000000400010043f000000200030008c0000009b0000413d00000009010000290000000001010433000000ff0010008c0000009b0000213d00000004020000290000000002020433000000ff0220018f0000000002120019000000ff0020008c000015cd0000213d000000240120008c000015150000213d00000001010000390000150e0000613d00000024022000890000000a03000039000000010020019000000000043300a9000000010300603900000000011300a900000001022002720000000003040019000015070000c13d000000030000006b000015720000613d00000003021000b900000003032000fa000000000013004b000015230000613d000015cd0000013d0000004d0010008c000015cd0000213d00000001020000390000000a03000039000000010010019000000000043300a9000000010300603900000000022300a900000001011002720000000003040019000015190000c13d000000000002004b000003ef0000613d00000003022000f9000008500020009c000017ab0000213d00000020010000290000084d0010009c000000a10000213d00000020030000290000004001300039000000400010043f0000000000230435000008be0100004100000000001004430000000001000414000008430010009c0000084301008041000000c001100210000008bf011001c70000800b02000039210821030000040f0000000100200190000015770000613d000000000101043b0000084301100197000000200200002900000020022000390000000000120435000000000001004b0000169d0000613d0000002001000029000000000101043300000850011001980000169d0000613d0000000f020000290000000002020433000008500310019800000000013200a90000154a0000613d00000000033100d9000000000023004b000015cd0000c13d000008c10110012a0000000d0200002900000000020204330000ffff0220018f00000000011200a90000000e020000290000000002020433000008c20110012a0000001103000029000008430330019700000843022001970000000002320019001100000002001d000008430020009c000015cd0000213d000000120200002900000843022001970000001003000029000000000303043300000843033001970000000002230019001200000002001d000008430020009c000015cd0000213d000000130200002900000000020204330000084302200197000008b9022000d1000000000021004b0000156e0000413d0000000c0200002900000000020204330000084302200197000008b9022000d1000000000021004b000013980000a13d000000150020002a000015cd0000413d001500150020002d0000139b0000013d000000000200001900000020010000290000084d0010009c000015280000a13d000000a10000013d000000000001042f000000000065004b00000000010000190000084b010080410000084b086001970000084b09500197000000000a89013f000000000089004b00000000080000190000084b080040410000084b00a0009c000000000801c019000000000008004b0000009b0000c13d00000000015700190000000408100039000000000182034f000000000101043b000008490010009c0000009b0000213d00000000091400490000002008800039000000000098004b000000000a0000190000084b0a0020410000084b099001970000084b08800197000000000b98013f000000000098004b00000000080000190000084b080040410000084b00b0009c00000000080ac019000000000008004b0000009b0000c13d000000160900002900000140089000c9000000000009004b000015a10000613d00000016098000fa000001400090008c000015cd0000c13d000001c001100039000000000081001a000015cd0000413d0000000001810019000000120010002a000015cd0000413d000000120810002a0000001b0100002900000000010104330000ffff0910018f00000000018900a9000015b00000613d00000000088100d9000000000098004b000015cd0000c13d0000001c0800002900000000080804330000084308800197000000000018001a000015cd0000413d000000000918001a00000000010000190000125e0000613d00000001010000290000007001100270000008c30110019700000000081900a900000000099800d9000000000019004b000015cd0000c13d000000000008004b00000000010000190000125e0000613d00000000093800a900000000018900d9000000000031004b000015cd0000c13d000000000009004b00000000010000190000125e0000613d000008c4019000d100000000039100d9000008c40030009c0000125e0000613d000008d701000041000000000010043f0000001101000039000000040010043f00000877010000410000210a00010430000000180500002900000000062100490000084c0060009c0000009b0000213d000000400060008c0000009b0000413d000000400600043d0000084d0060009c000000a10000213d0000004007600039000000400070043f0000000087020434000008480070009c0000009b0000213d00000000077604360000000008080433000008490080009c0000009b0000213d0000002005500039000000000087043500000000006504350000004002200039000000000042004b000015d40000413d000002220000013d000000160400002900000000052100490000084c0050009c0000009b0000213d000002400050008c0000009b0000413d000000400500043d0000084d0050009c000000a10000213d0000004006500039000000400060043f0000000086020434000008490060009c0000009b0000213d0000000006650436000000400700043d0000084f0070009c000000a10000213d0000022009700039000000400090043f0000000008080433000000000008004b0000000009000039000000010900c039000000000098004b0000009b0000c13d0000000008870436000000400920003900000000090904330000ffff0090008c0000009b0000213d000000000098043500000060082000390000000008080433000008430080009c0000009b0000213d0000004009700039000000000089043500000080082000390000000008080433000008430080009c0000009b0000213d00000060097000390000000000890435000000a0082000390000000008080433000008430080009c0000009b0000213d00000080097000390000000000890435000000c00820003900000000080804330000ffff0080008c0000009b0000213d000000a0097000390000000000890435000000e0082000390000000008080433000008430080009c0000009b0000213d000000c0097000390000000000890435000001000820003900000000080804330000ffff0080008c0000009b0000213d000000e0097000390000000000890435000001200820003900000000080804330000ffff0080008c0000009b0000213d00000100097000390000000000890435000001400820003900000000080804330000ffff0080008c0000009b0000213d0000012009700039000000000089043500000160082000390000000008080433000008430080009c0000009b0000213d0000014009700039000000000089043500000180082000390000000008080433000008430080009c0000009b0000213d00000160097000390000000000890435000001a0082000390000000008080433000008430080009c0000009b0000213d00000180097000390000000000890435000001c0082000390000000008080433000008490080009c0000009b0000213d000001a0097000390000000000890435000001e0082000390000000008080433000008430080009c0000009b0000213d000001c009700039000000000089043500000200082000390000000008080433000000000008004b0000000009000039000000010900c039000000000098004b0000009b0000c13d000001e00970003900000000008904350000022008200039000000000808043300000850008001980000009b0000c13d00000200097000390000000000890435000000000076043500000000045404360000024002200039000000000032004b000015ed0000413d0000024d0000013d000000000200041a0000085102200197000000000112019f000000000010041b0000001d01000029000008520010009c000000a10000213d0000001d010000290000002002100039001900000002001d000000400020043f0000000000010435000000400100043d001500000001001d0000084d0010009c000000a10000213d00000015030000290000004001300039000000400010043f00000020013000390000001d020000290000000000210435000000210100002900000000001304350000000001020433000000000001004b000016e70000c13d00000021010000290000000001010433000000000001004b000017520000c13d0000002301000029000000000101043300000848011001980000169a0000613d0000002402000029000000000202043300000847002001980000169a0000613d000000220200002900000000020204330000084300200198000018030000c13d000000400100043d0000088502000041000017a20000013d000000400100043d000008c002000041000000000021043500000004021000390000000803000029000006090000013d000000000792034f000000000a07043b000000040b60008c0000009b0000413d0000001b07600039000008e1077001970000003f07700039000008e108700197000000400700043d0000000008870019000000000078004b000000000c000039000000010c004039000008490080009c000000a10000213d0000000100c00190000000a10000c13d000000400080043f0000000008b70436000000000c96001900000000004c004b0000009b0000213d0000000404900039000000000442034f000008e109b001980000001f0bb0018f0000000002980019000016c50000613d000000000c04034f000000000d08001900000000ce0c043c000000000ded043600000000002d004b000016c10000c13d000008610aa0019700000000000b004b000016d30000613d000000000494034f0000000309b00210000000000b020433000000000b9b01cf000000000b9b022f000000000404043b0000010009900089000000000494022f00000000049401cf0000000004b4019f000000000042043500000000027600190000001c022000390000000000020435000008c600a0009c000018660000613d000008c700a0009c000016e40000c13d00000000020704330000084c0020009c0000009b0000213d000000200020008c0000009b0000413d000000400400043d0000084d0040009c000000a10000213d0000000002080433000012c80000013d000000400100043d000008d302000041000017a20000013d001f00000000001d000016ef0000013d0000001f02000029001f00010020003d0000001d0100002900000000010104330000001f0010006b000017d40000813d0000001f010000290000000501100210000000190110002900000000010104330000084801100197001e00000001001d000000000010043f0000000301000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000101041a002100000001001d000000000001004b000016e90000613d0000000201000039000000000201041a000000000002004b000015cd0000613d0000002103000029000000010130008a000000000032004b0000172b0000613d000000000012004b00001ad60000a13d0000002101000029000008540110009a000008540220009a000000000202041a000000000021041b000000000020043f0000000301000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b0000002102000029000000000021041b0000000201000039000000000101041a002100000001001d000000000001004b00001adc0000613d0000002101000029000000010110008a0000002102000029000008540220009a000000000002041b0000000202000039000000000012041b0000001e01000029000000000010043f0000000301000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000001041b000000400100043d0000001e020000290000000000210435000008430010009c000008430100804100000040011002100000000002000414000008430020009c0000084302008041000000c002200210000000000112019f00000855011001c70000800d0200003900000001030000390000085604000041210820fe0000040f0000000100200190000016e90000c13d0000009b0000013d0000002101000029001d00200010003d001e00000000001d0000001e0100002900000005011002100000001d011000290000000001010433001f08480010019c000017a00000613d0000001f01000029000000000010043f0000000301000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000101041a000000000001004b000017870000c13d0000000201000039000000000101041a000008490010009c000000a10000213d00000001021000390000000203000039000000000023041b000008570110009a0000001f02000029000000000021041b000000000103041a001900000001001d000000000020043f0000000301000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b0000001902000029000000000021041b000000400100043d0000001f020000290000000000210435000008430010009c000008430100804100000040011002100000000002000414000008430020009c0000084302008041000000c002200210000000000112019f00000855011001c70000800d0200003900000001030000390000085804000041210820fe0000040f00000001002001900000009b0000613d0000001e02000029001e00010020003d000000210100002900000000010104330000001e0010006b000017550000413d0000168e0000013d000000400100043d000008d5020000410000000000210435000008430010009c0000084301008041000000400110021000000886011001c70000210a00010430000008bd010000410000000902000029000017ad0000013d000008bd0100004100000020020000290000000000120435000008430020009c0000084302008041000000400120021000000886011001c70000210a00010430000000000006004b000017d80000c13d00000018040000290000000004040433000000000004004b000017d80000613d000000400100043d000008c902000041000017a20000013d0000001f0530018f0000084506300198000000400200043d0000000004620019000011580000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000017c30000c13d000011580000013d0000001f0530018f0000084506300198000000400200043d0000000004620019000011580000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000017cf0000c13d000011580000013d00000015010000290000000001010433002100000001001d0000168a0000013d000008430330019700000011033000290000000004530019000000000042001a000015cd0000413d0000000103000029000008c3053001980000000003000019000017f00000613d000000000342001900000000025300a900000000045200d9000000000034004b000015cd0000c13d000000000002004b0000000003000019000017f00000613d0000001a030000290000000003030433000008490430019700000000032400a900000000022300d9000000000042004b000015cd0000c13d0000000202000029000008490420019700000015024000b9000000150000006b000017f80000613d00000015052000fa000000000045004b000015cd0000c13d000000000023001a000015cd0000413d0000000002230019000000000012001a000015cd0000413d00000007030000290000085003300198000003ef0000613d000000000112001900000000013100d9000002dc0000013d000000a00010043f000000240100002900000000010104330000084701100197000000800010043f000000220100002900000000010104330000084301100197000000c00010043f000000400100043d001f00000001001d000008520010009c000000a10000213d0000001f010000290000002002100039001e00000002001d000000400020043f000000000001043500000020010000290000000001010433000000000001004b0000187e0000613d002400000000001d000018210000013d0000002402000029002400010020003d00000020010000290000000001010433000000240010006b0000187a0000813d000000240100002900000005011002100000001c01100029002200000001001d00000000010104330000084801100197002300000001001d000000000010043f0000000b01000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000101041a000000000001004b0000181b0000c13d0000000a01000039000000000101041a000008490010009c000000a10000213d00000001021000390000000a03000039000000000023041b000008590110009a0000002302000029000000000021041b000000000103041a002100000001001d000000000020043f0000000b01000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b0000002102000029000000000021041b00000020010000290000000001010433000000240010006c00001ad60000a13d0000002201000029000000000101043300000000020004140000084805100197000008430020009c0000084302008041000000c0012002100000085a011001c70000800d0200003900000002030000390000085b04000041210820fe0000040f00000001002001900000181b0000c13d0000009b0000013d00000000020704330000084c0020009c0000009b0000213d000000400020008c0000009b0000413d000000400400043d0000084d0040009c000000a10000213d0000004002400039000000400020043f0000000002080433000000000024043500000040067000390000000006060433000000000006004b0000000007000039000000010700c039000000000076004b000012cc0000613d0000009b0000013d0000001f010000290000000001010433000000000001004b00001a6a0000c13d0000001b010000290000000001010433000000000001004b000018c10000613d002400000000001d000000240100002900000005011002100000001a01100029000000000101043300000020021000390000000002020433002200000002001d00000000010104330000084801100197002300000001001d000000000010043f0000000601000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000201041a0000085e02200197000000220500002900000020035000390000000004030433000000a0044002100000085f04400197000000000242019f00000000040504330000084804400197000000000242019f000000000021041b000000400100043d00000000024104360000000003030433000000ff0330018f0000000000320435000008430010009c000008430100804100000040011002100000000002000414000008430020009c0000084302008041000000c002200210000000000112019f00000853011001c70000800d02000039000000020300003900000860040000410000002305000029210820fe0000040f00000001002001900000009b0000613d0000002402000029002400010020003d0000001b010000290000000001010433000000240010006b000018830000413d00000017010000290000000001010433000000000001004b0000198f0000613d002200000000001d00000022010000290000000501100210000000160110002900000000010104330000000012010434002308490020019c00001ae20000613d0000000001010433002400000001001d0000018001100039002000000001001d0000000001010433000008430110019800001ae20000613d00000024020000290000020002200039002100000002001d00000000020204330000086102200197000008620020009c00001ae20000c13d00000024020000290000016002200039001f00000002001d0000000002020433000008450020019800001ae20000613d00000024020000290000006002200039001e00000002001d00000000020204330000084302200197000000000021004b00001ae20000213d0000002301000029000000000010043f0000000801000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b0000000101100039000000000101041a001d00000001001d000000400100043d001c00000001001d000000240200002921081b3c0000040f0000001c020000290000000001210049000008430020009c0000084302008041000008430010009c000008430100804100000040022002100000006001100210000000000121019f0000001d020000290000086300200198000019130000613d0000000002000414000008430020009c0000084302008041000000c002200210000000000112019f0000085a011001c70000800d02000039000000020300003900000864040000410000191c0000013d0000000002000414000008430020009c0000084302008041000000c002200210000000000112019f0000085a011001c70000800d02000039000000020300003900000865040000410000002305000029210820fe0000040f00000001002001900000009b0000613d0000002301000029000000000010043f0000000801000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d00000024040000290000000032040434000000000002004b000000000101043b000000000201041a0000086602200197000000010220c1bf000000000303043300000008033002100000086703300197000000000232019f0000004003400039000000000303043300000018033002100000086803300197000000000232019f0000001e03000029000000000303043300000038033002100000086903300197000000000232019f0000008003400039000000000303043300000058033002100000086a03300197000000000232019f000000a003400039000000000303043300000078033002100000086b03300197000000000232019f000000c003400039000000000303043300000088033002100000086c03300197000000000232019f000000e0034000390000000003030433000000a8033002100000086d03300197000000000232019f00000100034000390000000003030433000000b8033002100000086e03300197000000000232019f00000120034000390000000003030433000000c8033002100000086f03300197000000000232019f00000140034000390000000003030433000000d8033002100000087003300197000000000232019f000000000021041b00000001011000390000001f0200002900000000020204330000084302200197000000000301041a0000087103300197000000000223019f0000002003000029000000000303043300000020033002100000087203300197000000000232019f000001a003400039000000000303043300000040033002100000087303300197000000000232019f000001c003400039000000000303043300000080033002100000087403300197000000000232019f000001e0034000390000000003030433000000000003004b00000875030000410000000003006019000000000232019f0000002103000029000000000303043300000038033002700000086303300197000000000232019f000000000021041b0000002202000029002200010020003d00000017010000290000000001010433000000220010006b000018c60000413d0018002b0000002d00000018010000290000000001010433000000000001004b000019cb0000613d0000001801000029002100200010003d002400000000001d000000240100002900000005011002100000002101100029000000000101043300000020021000390000000002020433002200000002001d00000000010104330000084801100197002300000001001d000000000010043f0000000701000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d00000022020000290000084902200197000000000101043b000000000301041a0000087803300197000000000323019f000000000031041b000000400100043d0000000000210435000008430010009c000008430100804100000040011002100000000002000414000008430020009c0000084302008041000000c002200210000000000112019f00000855011001c70000800d02000039000000020300003900000879040000410000002305000029210820fe0000040f00000001002001900000009b0000613d0000002402000029002400010020003d00000018010000290000000001010433000000240010006b000019960000413d000000400100043d001b00000001001d000008520010009c000000a10000213d0000001b010000290000002002100039001a00000002001d000000400020043f00000000000104350000002c01000029001c00000001001d0000000021010434001d00000002001d000000000001004b00001b240000613d001e00000000001d000019e00000013d0000001e02000029001e00010020003d0000001e0010006b00001ae50000813d0000001e0200002900000005022002100000001d0220002900000000030204330000002002300039001f00000002001d00000000020204330000000004020433000000000004004b000019dc0000613d0000000001030433002208490010019b002400000000001d000000240100002900000005011002100000000001120019000000200110003900000000010104330000000021010434002308480010019b0000000001020433002000000001001d0000008001100039002100000001001d000000000101043300000843011001970000001f0010008c00001b340000a13d0000002201000029000000000010043f0000000901000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b0000002302000029000000000020043f000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000201041a0000087a0220019700000020090000290000002003900039000000000403043300000020044002100000087204400197000000000242019f0000004004900039000000000504043300000040055002100000087b05500197000000000252019f0000006005900039000000000605043300000050066002100000087c06600197000000000262019f0000002108000029000000000608043300000070066002100000087d06600197000000000262019f000000a0069000390000000007060433000000000007004b0000087e070000410000000007006019000000000272019f00000000070904330000084307700197000000000272019f000000000021041b000000400100043d000000000271043600000000030304330000084303300197000000000032043500000000020404330000ffff0220018f00000040031000390000000000230435000000000205043300000843022001970000006003100039000000000023043500000000020804330000084302200197000000800310003900000000002304350000000002060433000000000002004b0000000002000039000000010200c039000000a0031000390000000000230435000008430010009c000008430100804100000040011002100000000002000414000008430020009c0000084302008041000000c002200210000000000112019f0000087f011001c70000800d020000390000000303000039000008800400004100000022050000290000002306000029210820fe0000040f00000001002001900000009b0000613d0000002403000029002400010030003d0000001f0100002900000000020104330000000001020433000000240010006b000019ed0000413d0000001c010000290000000001010433000019dc0000013d002400000000001d00001a720000013d0000002402000029002400010020003d0000001f010000290000000001010433000000240010006b0000187e0000813d000000240100002900000005011002100000001e01100029002200000001001d00000000010104330000084801100197002100000001001d000000000010043f0000000b01000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000101041a002300000001001d000000000001004b00001a6c0000613d0000000a01000039000000000201041a000000000002004b000015cd0000613d0000002303000029000000010130008a000000000023004b00001aaf0000613d000000000012004b00001ad60000a13d00000023010000290000085c0110009a0000085c0220009a000000000202041a000000000021041b000000000020043f0000000b01000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b0000002302000029000000000021041b0000000a01000039000000000101041a002300000001001d000000000001004b00001adc0000613d0000002301000029000000010110008a00000023020000290000085c0220009a000000000002041b0000000a02000039000000000012041b0000002101000029000000000010043f0000000b01000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000001041b0000001f010000290000000001010433000000240010006c00001ad60000a13d0000002201000029000000000101043300000000020004140000084805100197000008430020009c0000084302008041000000c0012002100000085a011001c70000800d0200003900000002030000390000085d04000041210820fe0000040f000000010020019000001a6c0000c13d0000009b0000013d000008d701000041000000000010043f0000003201000039000000040010043f00000877010000410000210a00010430000008d701000041000000000010043f0000003101000039000000040010043f00000877010000410000210a00010430000000400100043d0000087602000041000011db0000013d0000001b010000290000000001010433000000000001004b00001b240000613d002400000000001d000000240100002900000005011002100000001a01100029000000000101043300000020021000390000000002020433002300000002001d00000000010104330000084901100197002200000001001d000000000010043f0000000901000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d00000023020000290000084802200197000000000101043b002300000002001d000000000020043f000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f00000001002001900000009b0000613d000000000101043b000000000001041b0000000001000414000008430010009c0000084301008041000000c0011002100000085a011001c70000800d020000390000000303000039000008830400004100000022050000290000002306000029210820fe0000040f00000001002001900000009b0000613d0000002402000029002400010020003d0000001b010000290000000001010433000000240010006b00001aea0000413d000000800100043d00000140000004430000016000100443000000a00100043d00000020020000390000018000200443000001a000100443000000c00100043d0000004003000039000001c000300443000001e0001004430000010000200443000000030100003900000120001004430000088401000041000021090001042e000000400200043d000000240320003900000000001304350000088101000041000000000012043500000004012000390000002303000029000003870000013d0000000043020434000000000003004b0000000003000039000000010300c039000000000331043600000000040404330000ffff0440018f0000000000430435000000400320003900000000030304330000084303300197000000400410003900000000003404350000006003200039000000000303043300000843033001970000006004100039000000000034043500000080032000390000000003030433000008430330019700000080041000390000000000340435000000a00320003900000000030304330000ffff0330018f000000a0041000390000000000340435000000c00320003900000000030304330000084303300197000000c0041000390000000000340435000000e00320003900000000030304330000ffff0330018f000000e0041000390000000000340435000001000320003900000000030304330000ffff0330018f00000100041000390000000000340435000001200320003900000000030304330000ffff0330018f00000120041000390000000000340435000001400320003900000000030304330000084303300197000001400410003900000000003404350000016003200039000000000303043300000843033001970000016004100039000000000034043500000180032000390000000003030433000008430330019700000180041000390000000000340435000001a00320003900000000030304330000084903300197000001a0041000390000000000340435000001c00320003900000000030304330000084303300197000001c0041000390000000000340435000001e0032000390000000003030433000000000003004b0000000003000039000000010300c039000001e0041000390000000000340435000002000220003900000000020204330000086102200197000002000310003900000000002304350000022001100039000000000001042d000008e20010009c00001b980000813d0000004001100039000000400010043f000000000001042d000008d701000041000000000010043f0000004101000039000000040010043f00000877010000410000210a00010430000008e30010009c00001ba30000813d0000022001100039000000400010043f000000000001042d000008d701000041000000000010043f0000004101000039000000040010043f00000877010000410000210a00010430000008e40010009c00001bae0000813d000000c001100039000000400010043f000000000001042d000008d701000041000000000010043f0000004101000039000000040010043f00000877010000410000210a000104300000001f02200039000008e1022001970000000001120019000000000021004b00000000020000390000000102004039000008490010009c00001bc00000213d000000010020019000001bc00000c13d000000400010043f000000000001042d000008d701000041000000000010043f0000004101000039000000040010043f00000877010000410000210a0001043000000000430104340000000001320436000000000003004b00001bd20000613d000000000200001900000000052100190000000006240019000000000606043300000000006504350000002002200039000000000032004b00001bcb0000413d000000000231001900000000000204350000001f02300039000008e1022001970000000001210019000000000001042d00000020030000390000000004310436000000000302043300000000003404350000004001100039000000000003004b00001be70000613d000000000400001900000020022000390000000005020433000008480550019700000000015104360000000104400039000000000034004b00001be00000413d000000000001042d000000003101043400000850011001970000000001120436000000000203043300000843022001970000000000210435000000000001042d000000004302043400000843033001970000000003310436000000000404043300000843044001970000000000430435000000400320003900000000030304330000ffff0330018f000000400410003900000000003404350000006003200039000000000303043300000843033001970000006004100039000000000034043500000080032000390000000003030433000008430330019700000080041000390000000000340435000000a0022000390000000002020433000000000002004b0000000002000039000000010200c039000000a0031000390000000000230435000000c001100039000000000001042d0000000032020434000008480220019700000000022104360000000003030433000000ff0330018f00000000003204350000004001100039000000000001042d0004000000000002000000400200043d000008e20020009c00001d350000813d0000004003200039000000400030043f0000002003200039000000000003043500000000000204350000084801100197000400000001001d000000000010043f0000000601000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f000000010020019000001d3b0000613d000000400200043d0000084d0020009c00001d350000213d000000000101043b0000004003200039000000400030043f000000000101041a00000848031001980000000003320436000000a001100270000000ff0110018f000000000013043500001c4f0000613d000400000003001d000000400100043d0000084d0010009c00001d350000213d0000004003100039000000400030043f0000002003100039000000000003043500000000000104350000000002020433000000400c00043d000008ba0100004100000000051c043600000000010004140000084802200197000000040020008c00001c670000c13d0000000103000031000000a00030008c000000a004000039000000000403401900001c970000013d0000000401000029000000000010043f0000000501000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f000000010020019000001d3b0000613d000000400300043d0000084d0030009c00001d350000213d000000000101043b0000004002300039000000400020043f000000000101041a00000850021001970000000000230435000000e00110027000001d2e0000013d000100000005001d0000084300c0009c000008430300004100000000030c40190000004003300210000008430010009c0000084301008041000000c001100210000000000131019f00000886011001c7000200000002001d00030000000c001d210821030000040f000000030c000029000000000301001900000060033002700000084303300197000000a00030008c000000a00400003900000000040340190000001f0640018f000000e00740019000000000057c001900001c850000613d000000000801034f00000000090c0019000000008a08043c0000000009a90436000000000059004b00001c810000c13d000000000006004b00001c920000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f000000010020019000001d520000613d000000020200002900000001050000290000001f01400039000001e00110018f000000000bc1001900000000001b004b000000000100003900000001010040390000084900b0009c00001d350000213d000000010010019000001d350000c13d0000004000b0043f000000a00030008c00001d3b0000413d00000000010c0433000008bb0010009c00001d3b0000213d0000008001c000390000000001010433000008bb0010009c00001d3b0000213d00000000050504330000084c0050009c00001d430000213d000008bc0100004100000000001b04350000000001000414000000040020008c00001cb50000c13d000000200400003900001ce30000013d000200000005001d0000084300b0009c000008430300004100000000030b40190000004003300210000008430010009c0000084301008041000000c001100210000000000131019f00000886011001c700030000000b001d210821030000040f000000030b000029000000000301001900000060033002700000084303300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b001900001cd20000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b00001cce0000c13d000000000006004b00001cdf0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f000000010020019000001d5e0000613d00000002050000290000001f01400039000000600110018f0000000006b10019000008490060009c00001d350000213d000000400060043f000000200030008c00001d3b0000413d00000000010b0433000000ff0010008c00001d3b0000213d00000004020000290000000002020433000000ff0220018f0000000002120019000000ff0020008c00001d3d0000213d000000240120008c00001cf90000213d00001d080000c13d000000010100003900001d120000013d0000004d0010008c00001d3d0000213d0000000a030000390000000102000039000000010010019000000000043300a9000000010300603900000000022300a90000000101100272000000000304001900001cfd0000c13d000000000002004b00001d7c0000613d00000000022500d900001d180000013d0000000a0300003900000001010000390000002402200089000000010020019000000000043300a9000000010300603900000000011300a90000000102200272000000000304001900001d0b0000c13d000000000005004b00001d320000613d00000000025100a900000000035200d9000000000013004b00001d3d0000c13d000008500020009c00001d4b0000213d0000084d0060009c00001d350000213d0000004001600039000000400010043f000400000006001d0000000000260435000008be0100004100000000001004430000000001000414000008430010009c0000084301008041000000c001100210000008bf011001c70000800b02000039210821030000040f000000010020019000001d4a0000613d000000000101043b00000843011001970000000403000029000000200230003900000000001204350000000001030019000000000001042d00000000020000190000084d0060009c00001d1c0000a13d000008d701000041000000000010043f0000004101000039000000040010043f00000877010000410000210a0001043000000000010000190000210a00010430000008d701000041000000000010043f0000001101000039000000040010043f00000877010000410000210a00010430000008bd0100004100000000001b04350000084300b0009c000008430b0080410000004001b0021000000886011001c70000210a00010430000000000001042f000008bd010000410000000000160435000008430060009c0000084306008041000000400160021000000886011001c70000210a000104300000001f0530018f0000084506300198000000400200043d000000000462001900001d690000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001d590000c13d00001d690000013d0000001f0530018f0000084506300198000000400200043d000000000462001900001d690000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001d650000c13d000000000005004b00001d760000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000008430020009c00000843020080410000004002200210000000000121019f0000210a00010430000008d701000041000000000010043f0000001201000039000000040010043f00000877010000410000210a000104300006000000000002000400000001001d0000084901200197000300000001001d000000000010043f0000000401000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f000000010020019000001ef00000613d000000400200043d000600000002001d000008e20020009c00001eea0000813d000000000101043b00000006030000290000004002300039000000400020043f000000000101041a00000850021001970000000002230436000000e001100272000500000001001d000000000012043500001eff0000613d000008be0100004100000000001004430000000001000414000008430010009c0000084301008041000000c001100210000008bf011001c70000800b02000039210821030000040f000000010020019000001ef20000613d000000000101043b000500050010007400001ef30000413d000008d101000041000000000010044300000000010004120000000400100443000000400100003900000024001004430000000001000414000008430010009c0000084301008041000000c001100210000008d2011001c70000800502000039210821030000040f000000010020019000001ef20000613d000000400200043d000000000101043b00000843011001970000000504000029000000000014004b00001f0a0000213d0000084d0020009c00001eea0000213d0000004001200039000000400010043f00000020012000390000000000010435000000000002043500000004010000290000084801100197000500000001001d000000000010043f0000000601000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f000000010020019000001ef00000613d000000400200043d0000084d0020009c00001eea0000213d000000000101043b0000004003200039000000400030043f000000000101041a00000848031001980000000003320436000000a001100270000000ff0110018f000000000013043500001dfd0000613d000400000003001d000000400100043d0000084d0010009c00001eea0000213d0000004003100039000000400030043f0000002003100039000000000003043500000000000104350000000002020433000000400c00043d000008ba0100004100000000051c043600000000010004140000084802200197000000040020008c00001e150000c13d0000000103000031000000a00030008c000000a004000039000000000403401900001e450000013d0000000501000029000000000010043f0000000501000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f000000010020019000001ef00000613d000000400300043d0000084d0030009c00001eea0000213d000000000101043b0000004002300039000000400020043f000000000101041a00000850021001970000000000230435000000e00110027000001edc0000013d000100000005001d0000084300c0009c000008430300004100000000030c40190000004003300210000008430010009c0000084301008041000000c001100210000000000131019f00000886011001c7000200000002001d00030000000c001d210821030000040f000000030c000029000000000301001900000060033002700000084303300197000000a00030008c000000a00400003900000000040340190000001f0640018f000000e00740019000000000057c001900001e330000613d000000000801034f00000000090c0019000000008a08043c0000000009a90436000000000059004b00001e2f0000c13d000000000006004b00001e400000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f000000010020019000001f260000613d000000020200002900000001050000290000001f01400039000001e00110018f000000000bc1001900000000001b004b000000000100003900000001010040390000084900b0009c00001eea0000213d000000010010019000001eea0000c13d0000004000b0043f000000a00030008c00001ef00000413d00000000010c0433000008bb0010009c00001ef00000213d0000008001c000390000000001010433000008bb0010009c00001ef00000213d00000000050504330000084c0050009c00001f180000213d000008bc0100004100000000001b04350000000001000414000000040020008c00001e630000c13d000000200400003900001e910000013d000200000005001d0000084300b0009c000008430300004100000000030b40190000004003300210000008430010009c0000084301008041000000c001100210000000000131019f00000886011001c700030000000b001d210821030000040f000000030b000029000000000301001900000060033002700000084303300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b001900001e800000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b00001e7c0000c13d000000000006004b00001e8d0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f000000010020019000001f320000613d00000002050000290000001f01400039000000600110018f0000000006b10019000008490060009c00001eea0000213d000000400060043f000000200030008c00001ef00000413d00000000010b0433000000ff0010008c00001ef00000213d00000004020000290000000002020433000000ff0220018f0000000002120019000000ff0020008c00001ef30000213d000000240120008c00001ea70000213d00001eb60000c13d000000010100003900001ec00000013d0000004d0010008c00001ef30000213d0000000a030000390000000102000039000000010010019000000000043300a9000000010300603900000000022300a90000000101100272000000000304001900001eab0000c13d000000000002004b00001f500000613d00000000022500d900001ec60000013d0000000a0300003900000001010000390000002402200089000000010020019000000000043300a9000000010300603900000000011300a90000000102200272000000000304001900001eb90000c13d000000000005004b00001ee70000613d00000000025100a900000000035200d9000000000013004b00001ef30000c13d000008500020009c00001f1f0000213d0000084d0060009c00001eea0000213d0000004001600039000000400010043f000400000006001d0000000000260435000008be0100004100000000001004430000000001000414000008430010009c0000084301008041000000c001100210000008bf011001c70000800b02000039210821030000040f000000010020019000001ef20000613d000000000101043b0000084301100197000000040300002900000020023000390000000000120435000000000001004b00001ef90000613d0000000001030433000008500110019800001ef90000613d000000060200002900000000020204330000085002200197000000000001042d00000000020000190000084d0060009c00001eca0000a13d000008d701000041000000000010043f0000004101000039000000040010043f00000877010000410000210a0001043000000000010000190000210a00010430000000000001042f000008d701000041000000000010043f0000001101000039000000040010043f00000877010000410000210a00010430000000400100043d000008c00200004100000000002104350000000402100039000000050300002900001f040000013d000000400100043d000008e6020000410000000000210435000000040210003900000003030000290000000000320435000008430010009c0000084301008041000000400110021000000877011001c70000210a000104300000004403200039000000000043043500000024032000390000000000130435000008e5010000410000000000120435000000040120003900000003030000290000000000310435000008430020009c0000084302008041000000400120021000000889011001c70000210a00010430000008bd0100004100000000001b04350000084300b0009c000008430b0080410000004001b0021000000886011001c70000210a00010430000008bd010000410000000000160435000008430060009c0000084306008041000000400160021000000886011001c70000210a000104300000001f0530018f0000084506300198000000400200043d000000000462001900001f3d0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001f2d0000c13d00001f3d0000013d0000001f0530018f0000084506300198000000400200043d000000000462001900001f3d0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001f390000c13d000000000005004b00001f4a0000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000008430020009c00000843020080410000004002200210000000000112019f0000210a00010430000008d701000041000000000010043f0000001201000039000000040010043f00000877010000410000210a000104300005000000000002000000400200043d000008e20020009c0000207a0000813d0000004003200039000000400030043f0000002003200039000000000003043500000000000204350000084801100197000500000001001d000000000010043f0000000601000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f0000000100200190000020800000613d000000400200043d0000084d0020009c0000207a0000213d000000000101043b0000004003200039000000400030043f000000000101041a00000848031001980000000003320436000000a001100270000000ff0110018f000000000013043500001f900000613d000400000003001d000000400100043d0000084d0010009c0000207a0000213d0000004003100039000000400030043f0000002003100039000000000003043500000000000104350000000002020433000000400c00043d000008ba0100004100000000051c043600000000010004140000084802200197000000040020008c00001fa80000c13d0000000103000031000000a00030008c000000a004000039000000000403401900001fd80000013d0000000501000029000000000010043f0000000501000039000000200010043f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f0000000100200190000020800000613d000000400300043d0000084d0030009c0000207a0000213d000000000101043b0000004002300039000000400020043f000000000101041a00000850021001970000000000230435000000e0011002700000206f0000013d000100000005001d0000084300c0009c000008430300004100000000030c40190000004003300210000008430010009c0000084301008041000000c001100210000000000131019f00000886011001c7000200000002001d00030000000c001d210821030000040f000000030c000029000000000301001900000060033002700000084303300197000000a00030008c000000a00400003900000000040340190000001f0640018f000000e00740019000000000057c001900001fc60000613d000000000801034f00000000090c0019000000008a08043c0000000009a90436000000000059004b00001fc20000c13d000000000006004b00001fd30000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000020a20000613d000000020200002900000001050000290000001f01400039000001e00110018f000000000bc1001900000000001b004b000000000100003900000001010040390000084900b0009c0000207a0000213d00000001001001900000207a0000c13d0000004000b0043f000000a00030008c000020800000413d00000000010c0433000008bb0010009c000020800000213d0000008001c000390000000001010433000008bb0010009c000020800000213d00000000050504330000084c0050009c000020930000213d000008bc0100004100000000001b04350000000001000414000000040020008c00001ff60000c13d0000002004000039000020240000013d000200000005001d0000084300b0009c000008430300004100000000030b40190000004003300210000008430010009c0000084301008041000000c001100210000000000131019f00000886011001c700030000000b001d210821030000040f000000030b000029000000000301001900000060033002700000084303300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b0019000020130000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b0000200f0000c13d000000000006004b000020200000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000020ae0000613d00000002050000290000001f01400039000000600110018f0000000006b10019000008490060009c0000207a0000213d000000400060043f000000200030008c000020800000413d00000000010b0433000000ff0010008c000020800000213d00000004020000290000000002020433000000ff0220018f0000000002120019000000ff0020008c0000208d0000213d000000240120008c0000203a0000213d000020490000c13d0000000101000039000020530000013d0000004d0010008c0000208d0000213d0000000a030000390000000102000039000000010010019000000000043300a9000000010300603900000000022300a9000000010110027200000000030400190000203e0000c13d000000000002004b000020cc0000613d00000000022500d9000020590000013d0000000a0300003900000001010000390000002402200089000000010020019000000000043300a9000000010300603900000000011300a9000000010220027200000000030400190000204c0000c13d000000000005004b000020770000613d00000000025100a900000000035200d9000000000013004b0000208d0000c13d000008500020009c0000209b0000213d0000084d0060009c0000207a0000213d0000004001600039000000400010043f000400000006001d0000000000260435000008be0100004100000000001004430000000001000414000008430010009c0000084301008041000000c001100210000008bf011001c70000800b02000039210821030000040f00000001002001900000209a0000613d000000000101043b0000084301100197000000040300002900000020023000390000000000120435000000000001004b000020820000613d00000000010304330000085001100198000020820000613d000000000001042d00000000020000190000084d0060009c0000205d0000a13d000008d701000041000000000010043f0000004101000039000000040010043f00000877010000410000210a0001043000000000010000190000210a00010430000000400100043d000008c0020000410000000000210435000000040210003900000005030000290000000000320435000008430010009c0000084301008041000000400110021000000877011001c70000210a00010430000008d701000041000000000010043f0000001101000039000000040010043f00000877010000410000210a00010430000008bd0100004100000000001b04350000084300b0009c000008430b0080410000004001b0021000000886011001c70000210a00010430000000000001042f000008bd010000410000000000160435000008430060009c0000084306008041000000400160021000000886011001c70000210a000104300000001f0530018f0000084506300198000000400200043d0000000004620019000020b90000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000020a90000c13d000020b90000013d0000001f0530018f0000084506300198000000400200043d0000000004620019000020b90000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000020b50000c13d000000000005004b000020c60000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f00000000001404350000006001300210000008430020009c00000843020080410000004002200210000000000112019f0000210a00010430000008d701000041000000000010043f0000001201000039000000040010043f00000877010000410000210a00010430000000000001042f0000000001000414000008430010009c0000084301008041000000c00110021000000853011001c70000801002000039210821030000040f0000000100200190000020de0000613d000000000101043b000000000001042d00000000010000190000210a0001043000000000050100190000000000200443000000050030008c000020ee0000413d000000040100003900000000020000190000000506200210000000000664001900000005066002700000000006060031000000000161043a0000000102200039000000000031004b000020e60000413d000008430030009c000008430300804100000060013002100000000002000414000008430020009c0000084302008041000000c002200210000000000112019f000008e7011001c70000000002050019210821030000040f0000000100200190000020fd0000613d000000000101043b000000000001042d000000000001042f00002101002104210000000102000039000000000001042d0000000002000019000000000001042d00002106002104230000000102000039000000000001042d0000000002000019000000000001042d0000210800000432000021090001042e0000210a00010430000000000000000000000000000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000000000000000000000000001ffffffe000000000000000000000000000000000000000000000000000000000ffffffe0000000000000000000000000000000000000000000000000ffffffffffffff9f0000000000000000000000000000000000000000ffffffffffffffffffffffff000000000000000000000000ffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe080000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffffbf000000000000000000000000000000000000000000000000ffffffffffffff3f000000000000000000000000000000000000000000000000fffffffffffffddf00000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffdf0200000000000000000000000000000000000040000000000000000000000000bfa87805ed57dc1f0d489ce33be4c4577d74ccde357eeeee058a32c55c44a5330200000000000000000000000000000000000020000000000000000000000000c3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580bfa87805ed57dc1f0d489ce33be4c4577d74ccde357eeeee058a32c55c44a532eb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef39a5844729cae3e308f36a5ce933956d7c6367997d26743ca06a70b77c062d580200000000000000000000000000000000000000000000000000000000000000df1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba2339a5844729cae3e308f36a5ce933956d7c6367997d26743ca06a70b77c062d591795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91ffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000000000000ff000000000000000000000000000000000000000008a5f7f5bb38a81d8e43aca13ecd76431dbf8816ae4699affff7b00b2fc1c464ffffffff000000000000000000000000000000000000000000000000000000002812d52c0000000000000000000000000000000000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000a7b607fc10d28a1caf39ab7d27f4c94945db708a576d572781a455c5894fad93a937382a486d993de71c220bc8b559242deb4e286a353fa732330b4aa7d13577ff000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffff0000000000000000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000000000ffffffff0000000000000000000000000000000000000000000000000000ffff0000000000000000000000000000000000000000000000000000ffffffff0000000000000000000000000000000000000000000000000000ffff00000000000000000000000000000000000000000000000000000000ffff00000000000000000000000000000000000000000000000000000000ffff0000000000000000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000000000000000ffffffffffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffff0000000000000000000000000000000000000000ffffffffffffffff0000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000c35aa79d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000bb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97dffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000000000000000000000000000000ffff0000000000000000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000000000ffffffff0000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000002000000000000000000000000000000000000c000000000000000000000000094967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b524ecdc020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000440000000000000000000000004de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b0000000200000000000000000000000000000100000001000000000000000000d794ef9500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000043616e6e6f7420736574206f776e657220746f207a65726f000000000000000008c379a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000000000000000000000007afac32100000000000000000000000000000000000000000000000000000000cc88924b00000000000000000000000000000000000000000000000000000000d8694ccc00000000000000000000000000000000000000000000000000000000f700042900000000000000000000000000000000000000000000000000000000f700042a00000000000000000000000000000000000000000000000000000000ffdb4b3700000000000000000000000000000000000000000000000000000000d8694ccd00000000000000000000000000000000000000000000000000000000f2fde38b00000000000000000000000000000000000000000000000000000000cc88924c00000000000000000000000000000000000000000000000000000000cdc73d5100000000000000000000000000000000000000000000000000000000d02641a00000000000000000000000000000000000000000000000000000000091a2749900000000000000000000000000000000000000000000000000000000bf78e03e00000000000000000000000000000000000000000000000000000000bf78e03f00000000000000000000000000000000000000000000000000000000c4276bfc0000000000000000000000000000000000000000000000000000000091a2749a00000000000000000000000000000000000000000000000000000000a69c64c0000000000000000000000000000000000000000000000000000000007afac3220000000000000000000000000000000000000000000000000000000082b49eb0000000000000000000000000000000000000000000000000000000008da5cb5b00000000000000000000000000000000000000000000000000000000407e108500000000000000000000000000000000000000000000000000000000514e8cfe00000000000000000000000000000000000000000000000000000000770e2dc300000000000000000000000000000000000000000000000000000000770e2dc40000000000000000000000000000000000000000000000000000000079ba509700000000000000000000000000000000000000000000000000000000514e8cff000000000000000000000000000000000000000000000000000000006def4ce700000000000000000000000000000000000000000000000000000000407e10860000000000000000000000000000000000000000000000000000000045ac924d000000000000000000000000000000000000000000000000000000004ab35b0b00000000000000000000000000000000000000000000000000000000181f5a7600000000000000000000000000000000000000000000000000000000181f5a77000000000000000000000000000000000000000000000000000000002451a627000000000000000000000000000000000000000000000000000000003937306f000000000000000000000000000000000000000000000000000000000041e5be00000000000000000000000000000000000000000000000000000000061877e30000000000000000000000000000000000000000000000000000000006285c690000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffff7f02000000000000000000000000000000000002200000000000000000000000004f6e6c792063616c6c61626c65206279206f776e6572000000000000000000000000000000000000000000000000000000000064000000800000000000000000ed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127843616e6e6f74207472616e7366657220746f2073656c66000000000000000000000000000000000000000000fffffffffffffffffffffffffffffffffffffbff8d666f600000000000000000000000000000000000000000000000000000000000000000000000000000000000ff000000000000000000000000000000000000000000000000000000000000000000000000000000000000002386f26fc10000feaf968c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffff313ce5670000000000000000000000000000000000000000000000000000000010cb51d100000000000000000000000000000000000000000000000000000000796b89b91644bc98cd93958e4c9038275d622183e25ac5af08cc6b5d95539132020000020000000000000000000000000000000400000000000000000000000006439c6b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000de0b6b3a764000000000000000000000000000000000000000000000000000000000000000186a0000000000000000000000000000000000000ffffffffffffffffffffffffffff00000000000000000000000000000000000000000000000000005af3107a400000000000000000000000000000000000ffffffffffffffffffffffffffffffff181dcf100000000000000000000000000000000000000000000000000000000097a657c9000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000ee433e99000000000000000000000000000000000000000000000000000000004c4fc93a000000000000000000000000000000000000000000000000000000004c056b6a00000000000000000000000000000000000000000000000000000000869337890000000000000000000000000000000000000000000000000000000099ac52f200000000000000000000000000000000000000000000000000000000c65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a8000000000000002812d52c00000000000000000000000000000000000000000036f536ca00000000000000000000000000000000000000000000000000000000310ab089e4439a4c15d089f94afb7896ff553aecb10793d0ab882de59d99a32e02000002000000000000000000000000000000440000000000000000000000005247fdce000000000000000000000000000000000000000000000000000000006a92a483000000000000000000000000000000000000000000000000000000008579befe0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000008000000000000000004e487b71000000000000000000000000000000000000000000000000000000004d7573742062652070726f706f736564206f776e6572000000000000000000008be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e052f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148add84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6ed86ad9cf00000000000000000000000000000000000000000000000000000000405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace5072696365526567697374727920312e362e302d6465760000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000060000001400000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0000000000000000000000000000000000000000000000000ffffffffffffffc0000000000000000000000000000000000000000000000000fffffffffffffde0000000000000000000000000000000000000000000000000ffffffffffffff40f08bcb3e000000000000000000000000000000000000000000000000000000002e59db3a0000000000000000000000000000000000000000000000000000000002000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
