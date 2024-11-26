// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fee_quoter

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
	zkSyncAccounts "github.com/zksync-sdk/zksync2-go/accounts"
	zkSyncClient "github.com/zksync-sdk/zksync2-go/clients"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
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

type FeeQuoterDestChainConfig struct {
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
	DefaultTxGasLimit                 uint32
	GasMultiplierWeiPerEth            uint64
	NetworkFeeUSDCents                uint32
	GasPriceStalenessThreshold        uint32
	EnforceOutOfOrder                 bool
	ChainFamilySelector               [4]byte
}

type FeeQuoterDestChainConfigArgs struct {
	DestChainSelector uint64
	DestChainConfig   FeeQuoterDestChainConfig
}

type FeeQuoterPremiumMultiplierWeiPerEthArgs struct {
	Token                      common.Address
	PremiumMultiplierWeiPerEth uint64
}

type FeeQuoterStaticConfig struct {
	MaxFeeJuelsPerMsg            *big.Int
	LinkToken                    common.Address
	TokenPriceStalenessThreshold uint32
}

type FeeQuoterTokenPriceFeedConfig struct {
	DataFeedAddress common.Address
	TokenDecimals   uint8
}

type FeeQuoterTokenPriceFeedUpdate struct {
	SourceToken common.Address
	FeedConfig  FeeQuoterTokenPriceFeedConfig
}

type FeeQuoterTokenTransferFeeConfig struct {
	MinFeeUSDCents    uint32
	MaxFeeUSDCents    uint32
	DeciBps           uint16
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	IsEnabled         bool
}

type FeeQuoterTokenTransferFeeConfigArgs struct {
	DestChainSelector       uint64
	TokenTransferFeeConfigs []FeeQuoterTokenTransferFeeConfigSingleTokenArgs
}

type FeeQuoterTokenTransferFeeConfigRemoveArgs struct {
	DestChainSelector uint64
	Token             common.Address
}

type FeeQuoterTokenTransferFeeConfigSingleTokenArgs struct {
	Token                  common.Address
	TokenTransferFeeConfig FeeQuoterTokenTransferFeeConfig
}

type InternalEVM2AnyTokenTransfer struct {
	SourcePoolAddress common.Address
	DestTokenAddress  []byte
	ExtraData         []byte
	Amount            *big.Int
	DestExecData      []byte
}

type InternalGasPriceUpdate struct {
	DestChainSelector uint64
	UsdPerUnitGas     *big.Int
}

type InternalPriceUpdates struct {
	TokenPriceUpdates []InternalTokenPriceUpdate
	GasPriceUpdates   []InternalGasPriceUpdate
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

var FeeQuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"tokenPriceStalenessThreshold\",\"type\":\"uint32\"}],\"internalType\":\"structFeeQuoter.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"priceUpdaters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"feeTokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"name\":\"feedConfig\",\"type\":\"tuple\"}],\"internalType\":\"structFeeQuoter.TokenPriceFeedUpdate[]\",\"name\":\"tokenPriceFeeds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"internalType\":\"structFeeQuoter.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DataFeedValueOutOfUint224Range\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"DestinationChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtraArgOutOfOrderExecutionMustBeTrue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"FeeTokenNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"}],\"name\":\"InvalidDestBytesOverhead\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidDestChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStaticConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgFeeJuels\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint256\"}],\"name\":\"MessageFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"bytes10\",\"name\":\"workflowName\",\"type\":\"bytes10\"},{\"internalType\":\"bytes2\",\"name\":\"reportName\",\"type\":\"bytes2\"}],\"name\":\"ReportForwarderUnauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SourceTokenDataTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePassed\",\"type\":\"uint256\"}],\"name\":\"StaleGasPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storedTimeStamp\",\"type\":\"uint256\"}],\"name\":\"StaleKeystoneUpdate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"UnauthorizedCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedCallerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedCallerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"indexed\":false,\"internalType\":\"structFeeQuoter.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"name\":\"DestChainConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"FeeTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"FeeTokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"name\":\"PremiumMultiplierWeiPerEthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"name\":\"priceFeedConfig\",\"type\":\"tuple\"}],\"name\":\"PriceFeedPerTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reportId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bytes10\",\"name\":\"workflowName\",\"type\":\"bytes10\"},{\"internalType\":\"bytes2\",\"name\":\"reportName\",\"type\":\"bytes2\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structKeystoneFeedsPermissionHandler.Permission\",\"name\":\"permission\",\"type\":\"tuple\"}],\"name\":\"ReportPermissionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenTransferFeeConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"name\":\"TokenTransferFeeConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UsdPerTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"destChain\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UsdPerUnitGasUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_BASE_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KEYSTONE_PRICE_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"addedCallers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removedCallers\",\"type\":\"address[]\"}],\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\"}],\"name\":\"applyAuthorizedCallerUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"internalType\":\"structFeeQuoter.DestChainConfig\",\"name\":\"destChainConfig\",\"type\":\"tuple\"}],\"internalType\":\"structFeeQuoter.DestChainConfigArgs[]\",\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyDestChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"feeTokensToAdd\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"feeTokensToRemove\",\"type\":\"address[]\"}],\"name\":\"applyFeeTokensUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"internalType\":\"structFeeQuoter.PremiumMultiplierWeiPerEthArgs[]\",\"name\":\"premiumMultiplierWeiPerEthArgs\",\"type\":\"tuple[]\"}],\"name\":\"applyPremiumMultiplierWeiPerEthUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigSingleTokenArgs[]\",\"name\":\"tokenTransferFeeConfigs\",\"type\":\"tuple[]\"}],\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structFeeQuoter.TokenTransferFeeConfigRemoveArgs[]\",\"name\":\"tokensToUseDefaultFeeConfigs\",\"type\":\"tuple[]\"}],\"name\":\"applyTokenTransferFeeConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"}],\"name\":\"convertTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAuthorizedCallers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplierWeiPerEth\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasPriceStalenessThreshold\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enforceOutOfOrder\",\"type\":\"bool\"},{\"internalType\":\"bytes4\",\"name\":\"chainFamilySelector\",\"type\":\"bytes4\"}],\"internalType\":\"structFeeQuoter.DestChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getDestinationChainGasPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.TimestampedPackedUint224\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPremiumMultiplierWeiPerEth\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"premiumMultiplierWeiPerEth\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"tokenPriceStalenessThreshold\",\"type\":\"uint32\"}],\"internalType\":\"structFeeQuoter.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"}],\"name\":\"getTokenAndGasPrices\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"tokenPrice\",\"type\":\"uint224\"},{\"internalType\":\"uint224\",\"name\":\"gasPriceValue\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.TimestampedPackedUint224\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenPriceFeedConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getTokenPrices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"value\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"internalType\":\"structInternal.TimestampedPackedUint224[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenTransferFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"minFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxFeeUSDCents\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"deciBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structFeeQuoter.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getValidatedFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getValidatedTokenPrice\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"onReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourcePoolAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destExecData\",\"type\":\"bytes\"}],\"internalType\":\"structInternal.EVM2AnyTokenTransfer[]\",\"name\":\"onRampTokenTransfers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\"}],\"name\":\"processMessageArgs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"msgFeeJuels\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOutOfOrderExecution\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"convertedExtraArgs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"destExecDataPerToken\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bytes10\",\"name\":\"workflowName\",\"type\":\"bytes10\"},{\"internalType\":\"bytes2\",\"name\":\"reportName\",\"type\":\"bytes2\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"internalType\":\"structKeystoneFeedsPermissionHandler.Permission[]\",\"name\":\"permissions\",\"type\":\"tuple[]\"}],\"name\":\"setReportPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint224\",\"name\":\"usdPerToken\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint224\",\"name\":\"usdPerUnitGas\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.GasPriceUpdate[]\",\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.PriceUpdates\",\"name\":\"priceUpdates\",\"type\":\"tuple\"}],\"name\":\"updatePrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"dataFeedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structFeeQuoter.TokenPriceFeedConfig\",\"name\":\"feedConfig\",\"type\":\"tuple\"}],\"internalType\":\"structFeeQuoter.TokenPriceFeedUpdate[]\",\"name\":\"tokenPriceFeedUpdates\",\"type\":\"tuple[]\"}],\"name\":\"updateTokenPriceFeeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200787038038062007870833981016040819052620000349162001871565b8533806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000207565b5050604080518082018252838152815160008152602080820190935291810191909152620000ee9150620002b2565b5060208701516001600160a01b0316158062000112575086516001600160601b0316155b80620001265750604087015163ffffffff16155b15620001455760405163d794ef9560e01b815260040160405180910390fd5b6020878101516001600160a01b031660a05287516001600160601b031660805260408089015163ffffffff1660c05280516000815291820190526200018c90869062000401565b620001978462000549565b620001a2816200061a565b620001ad8262000a82565b60408051600080825260208201909252620001fa91859190620001f3565b6040805180820190915260008082526020820152815260200190600190039081620001cb5790505b5062000b4e565b5050505050505062001b2f565b336001600160a01b03821603620002615760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b602081015160005b815181101562000342576000828281518110620002db57620002db62001990565b60209081029190910101519050620002f560028262000e87565b1562000338576040516001600160a01b03821681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda775809060200160405180910390a15b50600101620002ba565b50815160005b8151811015620003fb57600082828151811062000369576200036962001990565b6020026020010151905060006001600160a01b0316816001600160a01b031603620003a7576040516342bcdf7f60e11b815260040160405180910390fd5b620003b460028262000ea7565b506040516001600160a01b03821681527feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef9060200160405180910390a15060010162000348565b50505050565b60005b8251811015620004a2576200044083828151811062000427576200042762001990565b6020026020010151600b62000ea760201b90919060201c565b1562000499578281815181106200045b576200045b62001990565b60200260200101516001600160a01b03167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba2360405160405180910390a25b60010162000404565b5060005b81518110156200054457620004e2828281518110620004c957620004c962001990565b6020026020010151600b62000ebe60201b90919060201c565b156200053b57818181518110620004fd57620004fd62001990565b60200260200101516001600160a01b03167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f9160405160405180910390a25b600101620004a6565b505050565b60005b8151811015620006165760008282815181106200056d576200056d62001990565b6020908102919091018101518051818301516001600160a01b0380831660008181526007875260409081902084518154868a018051929096166001600160a81b03199091168117600160a01b60ff9384160217909255825191825293519093169683019690965293955091939092917f08a5f7f5bb38a81d8e43aca13ecd76431dbf8816ae4699affff7b00b2fc1c464910160405180910390a25050508060010190506200054c565b5050565b60005b8151811015620006165760008282815181106200063e576200063e62001990565b6020026020010151905060008383815181106200065f576200065f62001990565b6020026020010151600001519050600082602001519050816001600160401b03166000148062000698575061016081015163ffffffff16155b80620006ba57506102008101516001600160e01b031916630a04b54b60e21b14155b80620006da5750806060015163ffffffff1681610160015163ffffffff16115b15620007055760405163c35aa79d60e01b81526001600160401b038316600482015260240162000083565b6001600160401b038216600090815260096020526040812060010154600160a81b900460e01b6001600160e01b03191690036200078557816001600160401b03167f525e3d4e0c31cef19cf9426af8d2c0ddd2d576359ca26bed92aac5fadda4626582604051620007779190620019a6565b60405180910390a2620007c9565b816001600160401b03167f283b699f411baff8f1c29fe49f32a828c8151596244b8e7e4c164edd6569a83582604051620007c09190620019a6565b60405180910390a25b8060096000846001600160401b03166001600160401b0316815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a8154816001600160401b0302191690836001600160401b031602179055506101a082015181600101600c6101000a81548163ffffffff021916908363ffffffff1602179055506101c08201518160010160106101000a81548163ffffffff021916908363ffffffff1602179055506101e08201518160010160146101000a81548160ff0219169083151502179055506102008201518160010160156101000a81548163ffffffff021916908360e01c02179055509050505050508060010190506200061d565b60005b81518110156200061657600082828151811062000aa65762000aa662001990565b6020026020010151600001519050600083838151811062000acb5762000acb62001990565b6020908102919091018101518101516001600160a01b03841660008181526008845260409081902080546001600160401b0319166001600160401b0385169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a2505060010162000a85565b60005b825181101562000dc157600083828151811062000b725762000b7262001990565b6020026020010151905060008160000151905060005b82602001515181101562000db25760008360200151828151811062000bb15762000bb162001990565b602002602001015160200151905060008460200151838151811062000bda5762000bda62001990565b6020026020010151600001519050602063ffffffff16826080015163ffffffff16101562000c395760808201516040516312766e0160e11b81526001600160a01b038316600482015263ffffffff909116602482015260440162000083565b6001600160401b0384166000818152600a602090815260408083206001600160a01b0386168085529083529281902086518154938801518389015160608a015160808b015160a08c01511515600160901b0260ff60901b1963ffffffff928316600160701b021664ffffffffff60701b199383166a01000000000000000000000263ffffffff60501b1961ffff90961668010000000000000000029590951665ffffffffffff60401b19968416640100000000026001600160401b0319909b16939097169290921798909817939093169390931717919091161792909217909155519091907f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b59062000d9f908690600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60405180910390a3505060010162000b88565b50505080600101905062000b51565b5060005b81518110156200054457600082828151811062000de65762000de662001990565b6020026020010151600001519050600083838151811062000e0b5762000e0b62001990565b6020908102919091018101518101516001600160401b0384166000818152600a845260408082206001600160a01b038516808452955280822080546001600160981b03191690555192945090917f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b9190a3505060010162000dc5565b600062000e9e836001600160a01b03841662000ed5565b90505b92915050565b600062000e9e836001600160a01b03841662000fd9565b600062000e9e836001600160a01b0384166200102b565b6000818152600183016020526040812054801562000fce57600062000efc60018362001af7565b855490915060009062000f129060019062001af7565b905081811462000f7e57600086600001828154811062000f365762000f3662001990565b906000526020600020015490508087600001848154811062000f5c5762000f5c62001990565b6000918252602080832090910192909255918252600188019052604090208390555b855486908062000f925762000f9262001b19565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000ea1565b600091505062000ea1565b6000818152600183016020526040812054620010225750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000ea1565b50600062000ea1565b6000818152600183016020526040812054801562000fce5760006200105260018362001af7565b8554909150600090620010689060019062001af7565b905080821462000f7e57600086600001828154811062000f365762000f3662001990565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715620010c757620010c76200108c565b60405290565b60405160c081016001600160401b0381118282101715620010c757620010c76200108c565b60405161022081016001600160401b0381118282101715620010c757620010c76200108c565b604051601f8201601f191681016001600160401b03811182821017156200114357620011436200108c565b604052919050565b80516001600160a01b03811681146200116357600080fd5b919050565b805163ffffffff811681146200116357600080fd5b6000606082840312156200119057600080fd5b604051606081016001600160401b0381118282101715620011b557620011b56200108c565b604052825190915081906001600160601b0381168114620011d557600080fd5b8152620011e5602084016200114b565b6020820152620011f86040840162001168565b60408201525092915050565b60006001600160401b038211156200122057620012206200108c565b5060051b60200190565b600082601f8301126200123c57600080fd5b81516020620012556200124f8362001204565b62001118565b8083825260208201915060208460051b8701019350868411156200127857600080fd5b602086015b848110156200129f5762001291816200114b565b83529183019183016200127d565b509695505050505050565b600082601f830112620012bc57600080fd5b81516020620012cf6200124f8362001204565b82815260609283028501820192828201919087851115620012ef57600080fd5b8387015b858110156200138257808903828112156200130e5760008081fd5b62001318620010a2565b62001323836200114b565b8152604080601f19840112156200133a5760008081fd5b62001344620010a2565b9250620013538885016200114b565b835283015160ff81168114620013695760008081fd5b82880152808701919091528452928401928101620012f3565b5090979650505050505050565b80516001600160401b03811681146200116357600080fd5b805161ffff811681146200116357600080fd5b805180151581146200116357600080fd5b600082601f830112620013dd57600080fd5b81516020620013f06200124f8362001204565b82815260059290921b840181019181810190868411156200141057600080fd5b8286015b848110156200129f5780516001600160401b03808211156200143557600080fd5b908801906040601f19838c0381018213156200145057600080fd5b6200145a620010a2565b620014678986016200138f565b815282850151848111156200147b57600080fd5b8086019550508c603f8601126200149157600080fd5b888501519350620014a66200124f8562001204565b84815260e09094028501830193898101908e861115620014c557600080fd5b958401955b858710156200159e57868f0360e0811215620014e557600080fd5b620014ef620010a2565b620014fa896200114b565b815260c086830112156200150d57600080fd5b62001517620010cd565b9150620015268d8a0162001168565b825262001535878a0162001168565b8d8301526200154760608a01620013a7565b878301526200155960808a0162001168565b60608301526200156c60a08a0162001168565b60808301526200157f60c08a01620013ba565b60a0830152808d0191909152825260e09690960195908a0190620014ca565b828b01525087525050509284019250830162001414565b600082601f830112620015c757600080fd5b81516020620015da6200124f8362001204565b82815260069290921b84018101918181019086841115620015fa57600080fd5b8286015b848110156200129f5760408189031215620016195760008081fd5b62001623620010a2565b6200162e826200114b565b81526200163d8583016200138f565b81860152835291830191604001620015fe565b80516001600160e01b0319811681146200116357600080fd5b600082601f8301126200167b57600080fd5b815160206200168e6200124f8362001204565b8281526102409283028501820192828201919087851115620016af57600080fd5b8387015b85811015620013825780890382811215620016ce5760008081fd5b620016d8620010a2565b620016e3836200138f565b815261022080601f1984011215620016fb5760008081fd5b62001705620010f2565b925062001714888501620013ba565b8352604062001725818601620013a7565b8985015260606200173881870162001168565b82860152608091506200174d82870162001168565b9085015260a06200176086820162001168565b8286015260c0915062001775828701620013a7565b9085015260e06200178886820162001168565b8286015261010091506200179e828701620013a7565b90850152610120620017b2868201620013a7565b828601526101409150620017c8828701620013a7565b90850152610160620017dc86820162001168565b828601526101809150620017f282870162001168565b908501526101a0620018068682016200138f565b828601526101c091506200181c82870162001168565b908501526101e06200183086820162001168565b82860152610200915062001846828701620013ba565b908501526200185785830162001650565b9084015250808701919091528452928401928101620016b3565b6000806000806000806000610120888a0312156200188e57600080fd5b6200189a89896200117d565b60608901519097506001600160401b0380821115620018b857600080fd5b620018c68b838c016200122a565b975060808a0151915080821115620018dd57600080fd5b620018eb8b838c016200122a565b965060a08a01519150808211156200190257600080fd5b620019108b838c01620012aa565b955060c08a01519150808211156200192757600080fd5b620019358b838c01620013cb565b945060e08a01519150808211156200194c57600080fd5b6200195a8b838c01620015b5565b93506101008a01519150808211156200197257600080fd5b50620019818a828b0162001669565b91505092959891949750929550565b634e487b7160e01b600052603260045260246000fd5b81511515815261022081016020830151620019c7602084018261ffff169052565b506040830151620019e0604084018263ffffffff169052565b506060830151620019f9606084018263ffffffff169052565b50608083015162001a12608084018263ffffffff169052565b5060a083015162001a2960a084018261ffff169052565b5060c083015162001a4260c084018263ffffffff169052565b5060e083015162001a5960e084018261ffff169052565b506101008381015161ffff9081169184019190915261012080850151909116908301526101408084015163ffffffff9081169184019190915261016080850151821690840152610180808501516001600160401b0316908401526101a0808501518216908401526101c080850151909116908301526101e080840151151590830152610200928301516001600160e01b031916929091019190915290565b8181038181111562000ea157634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60805160a05160c051615cee62001b82600039600081816102fa01526118750152600081816102be01528181610eb80152610f1801526000818161028a01528181610f410152610fb10152615cee6000f3fe608060405234801561001057600080fd5b50600436106101d95760003560e01c8063770e2dc411610104578063a69c64c0116100a2578063d63d3af211610071578063d63d3af214610ab0578063d8694ccd14610ab8578063f2fde38b14610acb578063ffdb4b3714610ade57600080fd5b8063a69c64c0146109d5578063bf78e03f146109e8578063cdc73d5114610a95578063d02641a014610a9d57600080fd5b8063805f2132116100de578063805f21321461081757806382b49eb01461082a5780638da5cb5b1461099a57806391a2749a146109c257600080fd5b8063770e2dc4146107e957806379ba5097146107fc5780637afac3221461080457600080fd5b8063407e10861161017c5780634ab35b0b1161014b5780634ab35b0b14610457578063514e8cff146104975780636cb5f3dd1461053a5780636def4ce71461054d57600080fd5b8063407e1086146103ee57806341ed29e714610401578063430d138c1461041457806345ac924d1461043757600080fd5b8063181f5a77116101b8578063181f5a77146103735780632451a627146103bc578063325c868e146103d15780633937306f146103d957600080fd5b806241e5be146101de578063061877e31461020457806306285c691461025d575b600080fd5b6101f16101ec3660046143ce565b610b26565b6040519081526020015b60405180910390f35b61024461021236600461440a565b73ffffffffffffffffffffffffffffffffffffffff1660009081526008602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff90911681526020016101fb565b610327604080516060810182526000808252602082018190529181019190915260405180606001604052807f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000063ffffffff16815250905090565b6040805182516bffffffffffffffffffffffff16815260208084015173ffffffffffffffffffffffffffffffffffffffff16908201529181015163ffffffff16908201526060016101fb565b6103af6040518060400160405280601381526020017f46656551756f74657220312e362e302d6465760000000000000000000000000081525081565b6040516101fb9190614489565b6103c4610b94565b6040516101fb919061449c565b6101f1602481565b6103ec6103e73660046144f6565b610ba5565b005b6103ec6103fc366004614698565b610e5a565b6103ec61040f3660046147ca565b610e6e565b6104276104223660046149a4565b610eb0565b6040516101fb9493929190614a98565b61044a610445366004614b37565b6110c0565b6040516101fb9190614b79565b61046a61046536600461440a565b61118b565b6040517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90911681526020016101fb565b61052d6104a5366004614bf4565b60408051808201909152600080825260208201525067ffffffffffffffff166000908152600560209081526040918290208251808401909352547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811683527c0100000000000000000000000000000000000000000000000000000000900463ffffffff169082015290565b6040516101fb9190614c0f565b6103ec610548366004614ca0565b611196565b6107dc61055b366004614bf4565b6040805161022081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c081018290526101e081018290526102008101919091525067ffffffffffffffff908116600090815260096020908152604091829020825161022081018452815460ff8082161515835261ffff61010080840482169685019690965263ffffffff630100000084048116978501979097526701000000000000008304871660608501526b0100000000000000000000008304871660808501526f010000000000000000000000000000008304811660a0850152710100000000000000000000000000000000008304871660c08501527501000000000000000000000000000000000000000000808404821660e08087019190915277010000000000000000000000000000000000000000000000850483169786019790975279010000000000000000000000000000000000000000000000000084049091166101208501527b01000000000000000000000000000000000000000000000000000000909204861661014084015260019093015480861661016084015264010000000081049096166101808301526c01000000000000000000000000860485166101a083015270010000000000000000000000000000000086049094166101c082015274010000000000000000000000000000000000000000850490911615156101e08201527fffffffff0000000000000000000000000000000000000000000000000000000092909304901b1661020082015290565b6040516101fb9190614ec0565b6103ec6107f73660046150be565b6111a7565b6103ec6111b9565b6103ec6108123660046153d8565b6112b6565b6103ec61082536600461543c565b6112c8565b61093a6108383660046154a8565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a08101919091525067ffffffffffffffff919091166000908152600a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff94909416835292815290829020825160c081018452905463ffffffff8082168352640100000000820481169383019390935268010000000000000000810461ffff16938201939093526a01000000000000000000008304821660608201526e01000000000000000000000000000083049091166080820152720100000000000000000000000000000000000090910460ff16151560a082015290565b6040516101fb9190600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101fb565b6103ec6109d03660046154d2565b6117b0565b6103ec6109e3366004615563565b6117c1565b610a616109f636600461440a565b6040805180820182526000808252602091820181905273ffffffffffffffffffffffffffffffffffffffff93841681526007825282902082518084019093525492831682527401000000000000000000000000000000000000000090920460ff169181019190915290565b60408051825173ffffffffffffffffffffffffffffffffffffffff16815260209283015160ff1692810192909252016101fb565b6103c46117d2565b61052d610aab36600461440a565b6117de565b6101f1601281565b6101f1610ac6366004615628565b611922565b6103ec610ad936600461440a565b611e5a565b610af1610aec36600461567d565b611e6b565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9384168152929091166020830152016101fb565b6000610b3182611f23565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16610b5885611f23565b610b80907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16856156d6565b610b8a91906156ed565b90505b9392505050565b6060610ba06002611fbd565b905090565b610bad611fca565b6000610bb98280615728565b9050905060005b81811015610d03576000610bd48480615728565b83818110610be457610be4615790565b905060400201803603810190610bfa91906157eb565b604080518082018252602080840180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff908116845263ffffffff42818116858701908152885173ffffffffffffffffffffffffffffffffffffffff9081166000908152600690975295889020965190519092167c010000000000000000000000000000000000000000000000000000000002919092161790935584519051935194955016927f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a92610cf29290917bffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a250600101610bc0565b506000610d136020840184615728565b9050905060005b81811015610e54576000610d316020860186615728565b83818110610d4157610d41615790565b905060400201803603810190610d579190615828565b604080518082018252602080840180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff908116845263ffffffff42818116858701908152885167ffffffffffffffff9081166000908152600590975295889020965190519092167c010000000000000000000000000000000000000000000000000000000002919092161790935584519051935194955016927fdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e92610e439290917bffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092168252602082015260400190565b60405180910390a250600101610d1a565b50505050565b610e6261200f565b610e6b81612090565b50565b610e7661200f565b60005b8151811015610eac57610ea4828281518110610e9757610e97615790565b602002602001015161218e565b600101610e79565b5050565b6000806060807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168c73ffffffffffffffffffffffffffffffffffffffff1603610f11578a9350610f3f565b610f3c8c8c7f0000000000000000000000000000000000000000000000000000000000000000610b26565b93505b7f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff16841115610fe3576040517f6a92a483000000000000000000000000000000000000000000000000000000008152600481018590526bffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660248201526044015b60405180910390fd5b67ffffffffffffffff8d1660009081526009602052604081206001015463ffffffff16906110128c8c84612360565b9050806020015194506110288f8b8b8b8b612509565b925085856110a8836040805182516024820152602092830151151560448083019190915282518083039091018152606490910190915290810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f181dcf100000000000000000000000000000000000000000000000000000000017905290565b95509550955050509950995099509995505050505050565b60608160008167ffffffffffffffff8111156110de576110de614531565b60405190808252806020026020018201604052801561112357816020015b60408051808201909152600080825260208201528152602001906001900390816110fc5790505b50905060005b828110156111805761115b86868381811061114657611146615790565b9050602002016020810190610aab919061440a565b82828151811061116d5761116d615790565b6020908102919091010152600101611129565b509150505b92915050565b600061118582611f23565b61119e61200f565b610e6b8161287a565b6111af61200f565b610eac8282612d4c565b60015473ffffffffffffffffffffffffffffffffffffffff16331461123a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610fda565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6112be61200f565b610eac828261315e565b600080600061130c87878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506132a592505050565b92509250925061131e338385846132c0565b600061132c8587018761584b565b905060005b81518110156117a55760006007600084848151811061135257611352615790565b6020908102919091018101515173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160009081205474010000000000000000000000000000000000000000900460ff169150819003611413578282815181106113bc576113bc615790565b6020908102919091010151516040517f06439c6b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610fda565b600061145c60128386868151811061142d5761142d615790565b6020026020010151602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16613418565b90506006600085858151811061147457611474615790565b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001601c9054906101000a900463ffffffff1663ffffffff168484815181106114e6576114e6615790565b60200260200101516040015163ffffffff1610156115f05783838151811061151057611510615790565b60200260200101516000015184848151811061152e5761152e615790565b6020026020010151604001516006600087878151811061155057611550615790565b6020908102919091018101515173ffffffffffffffffffffffffffffffffffffffff90811683529082019290925260409081016000205490517f191ec70600000000000000000000000000000000000000000000000000000000815293909116600484015263ffffffff91821660248401527c01000000000000000000000000000000000000000000000000000000009004166044820152606401610fda565b6040518060400160405280827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815260200185858151811061163157611631615790565b60200260200101516040015163ffffffff168152506006600086868151811061165c5761165c615790565b6020908102919091018101515173ffffffffffffffffffffffffffffffffffffffff168252818101929092526040016000208251929091015163ffffffff167c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90921691909117905583518490849081106116f4576116f4615790565b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff167f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a8286868151811061174a5761174a615790565b6020026020010151604001516040516117939291907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff92909216825263ffffffff16602082015260400190565b60405180910390a25050600101611331565b505050505050505050565b6117b861200f565b610e6b816134e4565b6117c961200f565b610e6b81613670565b6060610ba0600b611fbd565b604080518082019091526000808252602082015273ffffffffffffffffffffffffffffffffffffffff82166000908152600660209081526040918290208251808401909352547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116835263ffffffff7c010000000000000000000000000000000000000000000000000000000090910481169183018290527f000000000000000000000000000000000000000000000000000000000000000016906118a09042615912565b10156118ac5792915050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600760209081526040918290208251808401909352549283168083527401000000000000000000000000000000000000000090930460ff169082015290611911575092915050565b61191a8161375a565b949350505050565b67ffffffffffffffff8083166000908152600960209081526040808320815161022081018352815460ff808216151580845261ffff61010080850482169886019890985263ffffffff630100000085048116978601979097526701000000000000008404871660608601526b0100000000000000000000008404871660808601526f010000000000000000000000000000008404811660a0860152710100000000000000000000000000000000008404871660c08601527501000000000000000000000000000000000000000000808504821660e08088019190915277010000000000000000000000000000000000000000000000860483169987019990995279010000000000000000000000000000000000000000000000000085049091166101208601527b01000000000000000000000000000000000000000000000000000000909304861661014085015260019094015480861661016085015264010000000081049098166101808401526c01000000000000000000000000880485166101a084015270010000000000000000000000000000000088049094166101c083015274010000000000000000000000000000000000000000870490931615156101e08201527fffffffff000000000000000000000000000000000000000000000000000000009290950490921b16610200840152909190611b5c576040517f99ac52f200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610fda565b611b77611b6f608085016060860161440a565b600b906138e9565b611bd657611b8b608084016060850161440a565b6040517f2502348c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610fda565b6000611be56040850185615728565b9150611c41905082611bfa6020870187615925565b905083611c078880615925565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061391892505050565b6000611c5b611c56608087016060880161440a565b611f23565b90506000611c6e87856101c001516139c2565b9050600080808515611cae57611ca2878b611c8f60808d0160608e0161440a565b88611c9d60408f018f615728565b613ac2565b91945092509050611cce565b6101a0870151611ccb9063ffffffff16662386f26fc100006156d6565b92505b61010087015160009061ffff1615611d1257611d0f886dffffffffffffffffffffffffffff607088901c16611d0660208e018e615925565b90508a86613d9a565b90505b61018088015160009067ffffffffffffffff16611d3b611d3560808e018e615925565b8c613e4a565b600001518563ffffffff168b60a0015161ffff168e8060200190611d5f9190615925565b611d6a9291506156d6565b8c6080015163ffffffff16611d7f919061598a565b611d89919061598a565b611d93919061598a565b611dad906dffffffffffffffffffffffffffff89166156d6565b611db791906156d6565b9050867bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168282600860008f6060016020810190611df1919061440a565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002054611e2c9067ffffffffffffffff16896156d6565b611e36919061598a565b611e40919061598a565b611e4a91906156ed565b9c9b505050505050505050505050565b611e6261200f565b610e6b81613f0b565b67ffffffffffffffff8116600090815260096020526040812054819060ff16611ecc576040517f99ac52f200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610fda565b611ed584611f23565b67ffffffffffffffff8416600090815260096020526040902060010154611f17908590700100000000000000000000000000000000900463ffffffff166139c2565b915091505b9250929050565b600080611f2f836117de565b9050806020015163ffffffff1660001480611f67575080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16155b15611fb6576040517f06439c6b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610fda565b5192915050565b60606000610b8d83614000565b611fd56002336138e9565b61200d576040517fd86ad9cf000000000000000000000000000000000000000000000000000000008152336004820152602401610fda565b565b60005473ffffffffffffffffffffffffffffffffffffffff16331461200d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610fda565b60005b8151811015610eac5760008282815181106120b0576120b0615790565b60209081029190910181015180518183015173ffffffffffffffffffffffffffffffffffffffff80831660008181526007875260409081902084518154868a018051929096167fffffffffffffffffffffff00000000000000000000000000000000000000000090911681177401000000000000000000000000000000000000000060ff9384160217909255825191825293519093169683019690965293955091939092917f08a5f7f5bb38a81d8e43aca13ecd76431dbf8816ae4699affff7b00b2fc1c464910160405180910390a2505050806001019050612093565b600061224782600001518360600151846020015185604001516040805173ffffffffffffffffffffffffffffffffffffffff80871660208301528516918101919091527fffffffffffffffffffff00000000000000000000000000000000000000000000831660608201527fffff0000000000000000000000000000000000000000000000000000000000008216608082015260009060a001604051602081830303815290604052805190602001209050949350505050565b60808301516000828152600460205260409081902080549215157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909316929092179091555190915081907f32a4ba3fa3351b11ad555d4c8ec70a744e8705607077a946807030d64b6ab1a390612354908590600060a08201905073ffffffffffffffffffffffffffffffffffffffff8084511683527fffffffffffffffffffff0000000000000000000000000000000000000000000060208501511660208401527fffff00000000000000000000000000000000000000000000000000000000000060408501511660408401528060608501511660608401525060808301511515608083015292915050565b60405180910390a25050565b604080518082019091526000808252602082015260008390036123a157506040805180820190915267ffffffffffffffff8216815260006020820152610b8d565b60006123ad848661599d565b905060006123be85600481896159e3565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293505050507fffffffff0000000000000000000000000000000000000000000000000000000082167fe7e230f0000000000000000000000000000000000000000000000000000000000161245b57808060200190518101906124529190615a0d565b92505050610b8d565b7f6859a837000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316016124d7576040518060400160405280828060200190518101906124c39190615a39565b815260006020909101529250610b8d915050565b6040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff808616600090815260096020526040902060010154606091750100000000000000000000000000000000000000000090910460e01b90859081111561255957612559614531565b60405190808252806020026020018201604052801561258c57816020015b60608152602001906001900390816125775790505b50915060005b8581101561286f5760008585838181106125ae576125ae615790565b6125c4926020604090920201908101915061440a565b905060008888848181106125da576125da615790565b90506020028101906125ec9190615a52565b6125fa906040810190615925565b91505060208111156126aa5767ffffffffffffffff8a166000908152600a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff861684529091529020546e010000000000000000000000000000900463ffffffff168111156126aa576040517f36f536ca00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610fda565b61271a848a8a868181106126c0576126c0615790565b90506020028101906126d29190615a52565b6126e0906020810190615925565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061405c92505050565b67ffffffffffffffff8a166000818152600a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684528252808320815160c081018352905463ffffffff8082168352640100000000820481168386015268010000000000000000820461ffff16838501526a01000000000000000000008204811660608401526e010000000000000000000000000000820481166080840152720100000000000000000000000000000000000090910460ff16151560a08301908152958552600990935290832054935190937b0100000000000000000000000000000000000000000000000000000090049091169190612819578161281f565b82606001515b6040805163ffffffff831660208201529192500160405160208183030381529060405288878151811061285457612854615790565b60200260200101819052505050505050806001019050612592565b505095945050505050565b60005b8151811015610eac57600082828151811061289a5761289a615790565b6020026020010151905060008383815181106128b8576128b8615790565b60200260200101516000015190506000826020015190508167ffffffffffffffff16600014806128f1575061016081015163ffffffff16155b8061294357506102008101517fffffffff00000000000000000000000000000000000000000000000000000000167f2812d52c0000000000000000000000000000000000000000000000000000000014155b806129625750806060015163ffffffff1681610160015163ffffffff16115b156129a5576040517fc35aa79d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff83166004820152602401610fda565b67ffffffffffffffff82166000908152600960205260408120600101547501000000000000000000000000000000000000000000900460e01b7fffffffff00000000000000000000000000000000000000000000000000000000169003612a4d578167ffffffffffffffff167f525e3d4e0c31cef19cf9426af8d2c0ddd2d576359ca26bed92aac5fadda4626582604051612a409190614ec0565b60405180910390a2612a90565b8167ffffffffffffffff167f283b699f411baff8f1c29fe49f32a828c8151596244b8e7e4c164edd6569a83582604051612a879190614ec0565b60405180910390a25b80600960008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff02191690831515021790555060208201518160000160016101000a81548161ffff021916908361ffff16021790555060408201518160000160036101000a81548163ffffffff021916908363ffffffff16021790555060608201518160000160076101000a81548163ffffffff021916908363ffffffff160217905550608082015181600001600b6101000a81548163ffffffff021916908363ffffffff16021790555060a082015181600001600f6101000a81548161ffff021916908361ffff16021790555060c08201518160000160116101000a81548163ffffffff021916908363ffffffff16021790555060e08201518160000160156101000a81548161ffff021916908361ffff1602179055506101008201518160000160176101000a81548161ffff021916908361ffff1602179055506101208201518160000160196101000a81548161ffff021916908361ffff16021790555061014082015181600001601b6101000a81548163ffffffff021916908363ffffffff1602179055506101608201518160010160006101000a81548163ffffffff021916908363ffffffff1602179055506101808201518160010160046101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506101a082015181600101600c6101000a81548163ffffffff021916908363ffffffff1602179055506101c08201518160010160106101000a81548163ffffffff021916908363ffffffff1602179055506101e08201518160010160146101000a81548160ff0219169083151502179055506102008201518160010160156101000a81548163ffffffff021916908360e01c021790555090505050505080600101905061287d565b60005b8251811015613075576000838281518110612d6c57612d6c615790565b6020026020010151905060008160000151905060005b82602001515181101561306757600083602001518281518110612da757612da7615790565b6020026020010151602001519050600084602001518381518110612dcd57612dcd615790565b6020026020010151600001519050602063ffffffff16826080015163ffffffff161015612e505760808201516040517f24ecdc0200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015263ffffffff9091166024820152604401610fda565b67ffffffffffffffff84166000818152600a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168085529083529281902086518154938801518389015160608a015160808b015160a08c015115157201000000000000000000000000000000000000027fffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffff63ffffffff9283166e01000000000000000000000000000002167fffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffffff9383166a0100000000000000000000027fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff61ffff9096166801000000000000000002959095167fffffffffffffffffffffffffffffffffffff000000000000ffffffffffffffff968416640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000909b16939097169290921798909817939093169390931717919091161792909217909155519091907f94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b590613055908690600060c08201905063ffffffff80845116835280602085015116602084015261ffff60408501511660408401528060608501511660608401528060808501511660808401525060a0830151151560a083015292915050565b60405180910390a35050600101612d82565b505050806001019050612d4f565b5060005b815181101561315957600082828151811061309657613096615790565b602002602001015160000151905060008383815181106130b8576130b8615790565b60209081029190910181015181015167ffffffffffffffff84166000818152600a8452604080822073ffffffffffffffffffffffffffffffffffffffff8516808452955280822080547fffffffffffffffffffffffffff000000000000000000000000000000000000001690555192945090917f4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b9190a35050600101613079565b505050565b60005b82518110156132015761319783828151811061317f5761317f615790565b6020026020010151600b6140ae90919063ffffffff16565b156131f9578281815181106131ae576131ae615790565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167fdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba2360405160405180910390a25b600101613161565b5060005b81518110156131595761323b82828151811061322357613223615790565b6020026020010151600b6140d090919063ffffffff16565b1561329d5781818151811061325257613252615790565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167f1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f9160405160405180910390a25b600101613205565b6040810151604a820151605e90920151909260609290921c91565b6040805173ffffffffffffffffffffffffffffffffffffffff868116602080840191909152908616828401527fffffffffffffffffffff00000000000000000000000000000000000000000000851660608301527fffff00000000000000000000000000000000000000000000000000000000000084166080808401919091528351808403909101815260a09092018352815191810191909120600081815260049092529190205460ff16613411576040517f097e17ff00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8087166004830152851660248201527fffffffffffffffffffff00000000000000000000000000000000000000000000841660448201527fffff00000000000000000000000000000000000000000000000000000000000083166064820152608401610fda565b5050505050565b6000806134258486615a90565b9050600060248260ff16111561345f57613443602460ff8416615912565b61344e90600a615bc9565b61345890856156ed565b9050613485565b61346d60ff83166024615912565b61347890600a615bc9565b61348290856156d6565b90505b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8111156134db576040517f10cb51d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b95945050505050565b602081015160005b815181101561357f57600082828151811061350957613509615790565b602002602001015190506135278160026140f290919063ffffffff16565b156135765760405173ffffffffffffffffffffffffffffffffffffffff821681527fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda775809060200160405180910390a15b506001016134ec565b50815160005b8151811015610e545760008282815181106135a2576135a2615790565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613612576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61361d6002826140ae565b5060405173ffffffffffffffffffffffffffffffffffffffff821681527feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef9060200160405180910390a150600101613585565b60005b8151811015610eac57600082828151811061369057613690615790565b602002602001015160000151905060008383815181106136b2576136b2615790565b60209081029190910181015181015173ffffffffffffffffffffffffffffffffffffffff841660008181526008845260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff85169081179091559051908152919350917fbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d910160405180910390a25050600101613673565b604080518082019091526000808252602082015260008260000151905060008173ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa1580156137c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137e89190615bef565b5050509150506000811215613829576040517f10cb51d100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006138a88373ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015613879573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061389d9190615c3f565b866020015184613418565b604080518082019091527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff909116815263ffffffff4216602082015295945050505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515610b8d565b836040015163ffffffff168311156139715760408085015190517f8693378900000000000000000000000000000000000000000000000000000000815263ffffffff909116600482015260248101849052604401610fda565b836020015161ffff168211156139b3576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e548461020001518261405c565b67ffffffffffffffff821660009081526005602090815260408083208151808301909252547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116825263ffffffff7c010000000000000000000000000000000000000000000000000000000090910481169282019290925290831615613aba576000816020015163ffffffff1642613a579190615912565b90508363ffffffff16811115613ab8576040517ff08bcb3e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8616600482015263ffffffff8516602482015260448101829052606401610fda565b505b519392505050565b6000808083815b81811015613d8c576000878783818110613ae557613ae5615790565b905060400201803603810190613afb9190615c5c565b67ffffffffffffffff8c166000908152600a60209081526040808320845173ffffffffffffffffffffffffffffffffffffffff168452825291829020825160c081018452905463ffffffff8082168352640100000000820481169383019390935268010000000000000000810461ffff16938201939093526a01000000000000000000008304821660608201526e01000000000000000000000000000083049091166080820152720100000000000000000000000000000000000090910460ff16151560a0820181905291925090613c1b576101208d0151613be89061ffff16662386f26fc100006156d6565b613bf2908861598a565b96508c610140015186613c059190615c95565b9550613c12602086615c95565b94505050613d84565b604081015160009061ffff1615613cd45760008c73ffffffffffffffffffffffffffffffffffffffff16846000015173ffffffffffffffffffffffffffffffffffffffff1614613c77578351613c7090611f23565b9050613c7a565b508a5b620186a0836040015161ffff16613cbc8660200151847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1661411490919063ffffffff16565b613cc691906156d6565b613cd091906156ed565b9150505b6060820151613ce39088615c95565b9650816080015186613cf59190615c95565b8251909650600090613d149063ffffffff16662386f26fc100006156d6565b905080821015613d3357613d28818a61598a565b985050505050613d84565b6000836020015163ffffffff16662386f26fc10000613d5291906156d6565b905080831115613d7257613d66818b61598a565b99505050505050613d84565b613d7c838b61598a565b995050505050505b600101613ac9565b505096509650969350505050565b60008063ffffffff8316613db0610120866156d6565b613dbc876101c061598a565b613dc6919061598a565b613dd0919061598a565b905060008760c0015163ffffffff168860e0015161ffff1683613df391906156d6565b613dfd919061598a565b61010089015190915061ffff16613e246dffffffffffffffffffffffffffff8916836156d6565b613e2e91906156d6565b613e3e90655af3107a40006156d6565b98975050505050505050565b60408051808201909152600080825260208201526000613e76858585610160015163ffffffff16612360565b9050826060015163ffffffff1681600001511115613ec0576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826101e001518015613ed457508060200151155b15610b8a576040517fee433e9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff821603613f8a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610fda565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60608160000180548060200260200160405190810160405280929190818152602001828054801561405057602002820191906000526020600020905b81548152602001906001019080831161403c575b50505050509050919050565b7fd7ed2ad4000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831601610eac5761315981614151565b6000610b8d8373ffffffffffffffffffffffffffffffffffffffff8416614204565b6000610b8d8373ffffffffffffffffffffffffffffffffffffffff8416614253565b6000610b8d8373ffffffffffffffffffffffffffffffffffffffff841661434d565b6000670de0b6b3a7640000614147837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff86166156d6565b610b8d91906156ed565b6000815160201461419057816040517f8d666f60000000000000000000000000000000000000000000000000000000008152600401610fda9190614489565b6000828060200190518101906141a69190615a39565b905073ffffffffffffffffffffffffffffffffffffffff8111806141cb575061040081105b1561118557826040517f8d666f60000000000000000000000000000000000000000000000000000000008152600401610fda9190614489565b600081815260018301602052604081205461424b57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155611185565b506000611185565b6000818152600183016020526040812054801561433c576000614277600183615912565b855490915060009061428b90600190615912565b90508082146142f05760008660000182815481106142ab576142ab615790565b90600052602060002001549050808760000184815481106142ce576142ce615790565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061430157614301615cb2565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050611185565b6000915050611185565b5092915050565b6000818152600183016020526040812054801561433c576000614371600183615912565b855490915060009061438590600190615912565b90508181146142f05760008660000182815481106142ab576142ab615790565b803573ffffffffffffffffffffffffffffffffffffffff811681146143c957600080fd5b919050565b6000806000606084860312156143e357600080fd5b6143ec846143a5565b925060208401359150614401604085016143a5565b90509250925092565b60006020828403121561441c57600080fd5b610b8d826143a5565b6000815180845260005b8181101561444b5760208185018101518683018201520161442f565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610b8d6020830184614425565b6020808252825182820181905260009190848201906040850190845b818110156144ea57835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016144b8565b50909695505050505050565b60006020828403121561450857600080fd5b813567ffffffffffffffff81111561451f57600080fd5b820160408185031215610b8d57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561458357614583614531565b60405290565b60405160a0810167ffffffffffffffff8111828210171561458357614583614531565b604051610220810167ffffffffffffffff8111828210171561458357614583614531565b60405160c0810167ffffffffffffffff8111828210171561458357614583614531565b6040516060810167ffffffffffffffff8111828210171561458357614583614531565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561465d5761465d614531565b604052919050565b600067ffffffffffffffff82111561467f5761467f614531565b5060051b60200190565b60ff81168114610e6b57600080fd5b600060208083850312156146ab57600080fd5b823567ffffffffffffffff8111156146c257600080fd5b8301601f810185136146d357600080fd5b80356146e66146e182614665565b614616565b8181526060918202830184019184820191908884111561470557600080fd5b938501935b838510156147a557848903818112156147235760008081fd5b61472b614560565b614734876143a5565b81526040807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0840112156147685760008081fd5b614770614560565b925061477d8989016143a5565b835287013561478b81614689565b82890152808801919091528352938401939185019161470a565b50979650505050505050565b8015158114610e6b57600080fd5b80356143c9816147b1565b600060208083850312156147dd57600080fd5b823567ffffffffffffffff8111156147f457600080fd5b8301601f8101851361480557600080fd5b80356148136146e182614665565b81815260a0918202830184019184820191908884111561483257600080fd5b938501935b838510156147a55780858a03121561484f5760008081fd5b614857614589565b614860866143a5565b8152868601357fffffffffffffffffffff00000000000000000000000000000000000000000000811681146148955760008081fd5b818801526040868101357fffff000000000000000000000000000000000000000000000000000000000000811681146148ce5760008081fd5b9082015260606148df8782016143a5565b908201526080868101356148f2816147b1565b9082015283529384019391850191614837565b803567ffffffffffffffff811681146143c957600080fd5b60008083601f84011261492f57600080fd5b50813567ffffffffffffffff81111561494757600080fd5b602083019150836020828501011115611f1c57600080fd5b60008083601f84011261497157600080fd5b50813567ffffffffffffffff81111561498957600080fd5b6020830191508360208260051b8501011115611f1c57600080fd5b600080600080600080600080600060c08a8c0312156149c257600080fd5b6149cb8a614905565b98506149d960208b016143a5565b975060408a0135965060608a013567ffffffffffffffff808211156149fd57600080fd5b614a098d838e0161491d565b909850965060808c0135915080821115614a2257600080fd5b614a2e8d838e0161495f565b909650945060a08c0135915080821115614a4757600080fd5b818c0191508c601f830112614a5b57600080fd5b813581811115614a6a57600080fd5b8d60208260061b8501011115614a7f57600080fd5b6020830194508093505050509295985092959850929598565b848152600060208515158184015260806040840152614aba6080840186614425565b8381036060850152845180825282820190600581901b8301840184880160005b83811015614b26577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868403018552614b14838351614425565b94870194925090860190600101614ada565b50909b9a5050505050505050505050565b60008060208385031215614b4a57600080fd5b823567ffffffffffffffff811115614b6157600080fd5b614b6d8582860161495f565b90969095509350505050565b602080825282518282018190526000919060409081850190868401855b82811015614be757614bd784835180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16825260209081015163ffffffff16910152565b9284019290850190600101614b96565b5091979650505050505050565b600060208284031215614c0657600080fd5b610b8d82614905565b81517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815260208083015163ffffffff169082015260408101611185565b803561ffff811681146143c957600080fd5b803563ffffffff811681146143c957600080fd5b80357fffffffff00000000000000000000000000000000000000000000000000000000811681146143c957600080fd5b60006020808385031215614cb357600080fd5b823567ffffffffffffffff811115614cca57600080fd5b8301601f81018513614cdb57600080fd5b8035614ce96146e182614665565b8181526102409182028301840191848201919088841115614d0957600080fd5b938501935b838510156147a55784890381811215614d275760008081fd5b614d2f614560565b614d3887614905565b8152610220807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe084011215614d6d5760008081fd5b614d756145ac565b9250614d828989016147bf565b83526040614d91818a01614c4a565b8a8501526060614da2818b01614c5c565b8286015260809150614db5828b01614c5c565b9085015260a0614dc68a8201614c5c565b8286015260c09150614dd9828b01614c4a565b9085015260e0614dea8a8201614c5c565b828601526101009150614dfe828b01614c4a565b90850152610120614e108a8201614c4a565b828601526101409150614e24828b01614c4a565b90850152610160614e368a8201614c5c565b828601526101809150614e4a828b01614c5c565b908501526101a0614e5c8a8201614905565b828601526101c09150614e70828b01614c5c565b908501526101e0614e828a8201614c5c565b828601526102009150614e96828b016147bf565b90850152614ea5898301614c70565b90840152508088019190915283529384019391850191614d0e565b81511515815261022081016020830151614ee0602084018261ffff169052565b506040830151614ef8604084018263ffffffff169052565b506060830151614f10606084018263ffffffff169052565b506080830151614f28608084018263ffffffff169052565b5060a0830151614f3e60a084018261ffff169052565b5060c0830151614f5660c084018263ffffffff169052565b5060e0830151614f6c60e084018261ffff169052565b506101008381015161ffff9081169184019190915261012080850151909116908301526101408084015163ffffffff90811691840191909152610160808501518216908401526101808085015167ffffffffffffffff16908401526101a0808501518216908401526101c080850151909116908301526101e080840151151590830152610200808401517fffffffff000000000000000000000000000000000000000000000000000000008116828501525b505092915050565b600082601f83011261503757600080fd5b813560206150476146e183614665565b82815260069290921b8401810191818101908684111561506657600080fd5b8286015b848110156150b357604081890312156150835760008081fd5b61508b614560565b61509482614905565b81526150a18583016143a5565b8186015283529183019160400161506a565b509695505050505050565b600080604083850312156150d157600080fd5b67ffffffffffffffff833511156150e757600080fd5b83601f8435850101126150f957600080fd5b6151096146e18435850135614665565b8335840180358083526020808401939260059290921b9091010186101561512f57600080fd5b602085358601015b85358601803560051b0160200181101561533c5767ffffffffffffffff8135111561516157600080fd5b8035863587010160407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828a0301121561519a57600080fd5b6151a2614560565b6151ae60208301614905565b815267ffffffffffffffff604083013511156151c957600080fd5b88603f6040840135840101126151de57600080fd5b6151f46146e16020604085013585010135614665565b6020604084810135850182810135808552928401939260e00201018b101561521b57600080fd5b6040808501358501015b6040858101358601602081013560e002010181101561531d5760e0818d03121561524e57600080fd5b615256614560565b61525f826143a5565b815260c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838f0301121561529357600080fd5b61529b6145d0565b6152a760208401614c5c565b81526152b560408401614c5c565b60208201526152c660608401614c4a565b60408201526152d760808401614c5c565b60608201526152e860a08401614c5c565b60808201526152fa60c08401356147b1565b60c083013560a0820152602082810191909152908452929092019160e001615225565b5080602084015250508085525050602083019250602081019050615137565b5092505067ffffffffffffffff6020840135111561535957600080fd5b6153698460208501358501615026565b90509250929050565b600082601f83011261538357600080fd5b813560206153936146e183614665565b8083825260208201915060208460051b8701019350868411156153b557600080fd5b602086015b848110156150b3576153cb816143a5565b83529183019183016153ba565b600080604083850312156153eb57600080fd5b823567ffffffffffffffff8082111561540357600080fd5b61540f86838701615372565b9350602085013591508082111561542557600080fd5b5061543285828601615372565b9150509250929050565b6000806000806040858703121561545257600080fd5b843567ffffffffffffffff8082111561546a57600080fd5b6154768883890161491d565b9096509450602087013591508082111561548f57600080fd5b5061549c8782880161491d565b95989497509550505050565b600080604083850312156154bb57600080fd5b6154c483614905565b9150615369602084016143a5565b6000602082840312156154e457600080fd5b813567ffffffffffffffff808211156154fc57600080fd5b908301906040828603121561551057600080fd5b615518614560565b82358281111561552757600080fd5b61553387828601615372565b82525060208301358281111561554857600080fd5b61555487828601615372565b60208301525095945050505050565b6000602080838503121561557657600080fd5b823567ffffffffffffffff81111561558d57600080fd5b8301601f8101851361559e57600080fd5b80356155ac6146e182614665565b81815260069190911b820183019083810190878311156155cb57600080fd5b928401925b8284101561561d57604084890312156155e95760008081fd5b6155f1614560565b6155fa856143a5565b8152615607868601614905565b81870152825260409390930192908401906155d0565b979650505050505050565b6000806040838503121561563b57600080fd5b61564483614905565b9150602083013567ffffffffffffffff81111561566057600080fd5b830160a0818603121561567257600080fd5b809150509250929050565b6000806040838503121561569057600080fd5b615699836143a5565b915061536960208401614905565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082028115828204841417611185576111856156a7565b600082615723577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261575d57600080fd5b83018035915067ffffffffffffffff82111561577857600080fd5b6020019150600681901b3603821315611f1c57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811681146143c957600080fd5b6000604082840312156157fd57600080fd5b615805614560565b61580e836143a5565b815261581c602084016157bf565b60208201529392505050565b60006040828403121561583a57600080fd5b615842614560565b61580e83614905565b6000602080838503121561585e57600080fd5b823567ffffffffffffffff81111561587557600080fd5b8301601f8101851361588657600080fd5b80356158946146e182614665565b818152606091820283018401918482019190888411156158b357600080fd5b938501935b838510156147a55780858a0312156158d05760008081fd5b6158d86145f3565b6158e1866143a5565b81526158ee8787016157bf565b8782015260406158ff818801614c5c565b90820152835293840193918501916158b8565b81810381811115611185576111856156a7565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261595a57600080fd5b83018035915067ffffffffffffffff82111561597557600080fd5b602001915036819003821315611f1c57600080fd5b80820180821115611185576111856156a7565b7fffffffff00000000000000000000000000000000000000000000000000000000813581811691600485101561501e5760049490940360031b84901b1690921692915050565b600080858511156159f357600080fd5b83861115615a0057600080fd5b5050820193919092039150565b600060408284031215615a1f57600080fd5b615a27614560565b82518152602083015161581c816147b1565b600060208284031215615a4b57600080fd5b5051919050565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61833603018112615a8657600080fd5b9190910192915050565b60ff8181168382160190811115611185576111856156a7565b600181815b80851115615b0257817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115615ae857615ae86156a7565b80851615615af557918102915b93841c9390800290615aae565b509250929050565b600082615b1957506001611185565b81615b2657506000611185565b8160018114615b3c5760028114615b4657615b62565b6001915050611185565b60ff841115615b5757615b576156a7565b50506001821b611185565b5060208310610133831016604e8410600b8410161715615b85575081810a611185565b615b8f8383615aa9565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115615bc157615bc16156a7565b029392505050565b6000610b8d8383615b0a565b805169ffffffffffffffffffff811681146143c957600080fd5b600080600080600060a08688031215615c0757600080fd5b615c1086615bd5565b9450602086015193506040860151925060608601519150615c3360808701615bd5565b90509295509295909350565b600060208284031215615c5157600080fd5b8151610b8d81614689565b600060408284031215615c6e57600080fd5b615c76614560565b615c7f836143a5565b8152602083013560208201528091505092915050565b63ffffffff818116838216019080821115614346576143466156a7565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000818000a",
}

var FeeQuoterABI = FeeQuoterMetaData.ABI

var FeeQuoterBin = FeeQuoterMetaData.Bin

func DeployFeeQuoter(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig FeeQuoterStaticConfig, priceUpdaters []common.Address, feeTokens []common.Address, tokenPriceFeeds []FeeQuoterTokenPriceFeedUpdate, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (common.Address, *CustomTransaction, *FeeQuoter, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	switch chainId.Uint64() {

	case 324, 280, 300:
		return DeployZkSyncFeeQuoter(auth, backend, staticConfig, priceUpdaters, feeTokens, tokenPriceFeeds, tokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs, destChainConfigArgs)
	}

	parsed, err := FeeQuoterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeQuoterBin), backend, staticConfig, priceUpdaters, feeTokens, tokenPriceFeeds, tokenTransferFeeConfigArgs, premiumMultiplierWeiPerEthArgs, destChainConfigArgs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, &CustomTransaction{Transaction: tx, CustomHash: tx.Hash()}, &FeeQuoter{address: address, abi: *parsed, FeeQuoterCaller: FeeQuoterCaller{contract: contract}, FeeQuoterTransactor: FeeQuoterTransactor{contract: contract}, FeeQuoterFilterer: FeeQuoterFilterer{contract: contract}}, nil
}

type FeeQuoter struct {
	address common.Address
	abi     abi.ABI
	FeeQuoterCaller
	FeeQuoterTransactor
	FeeQuoterFilterer
}

type FeeQuoterCaller struct {
	contract *bind.BoundContract
}

type FeeQuoterTransactor struct {
	contract *bind.BoundContract
}

type FeeQuoterFilterer struct {
	contract *bind.BoundContract
}

type FeeQuoterSession struct {
	Contract     *FeeQuoter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type FeeQuoterCallerSession struct {
	Contract *FeeQuoterCaller
	CallOpts bind.CallOpts
}

type FeeQuoterTransactorSession struct {
	Contract     *FeeQuoterTransactor
	TransactOpts bind.TransactOpts
}

type FeeQuoterRaw struct {
	Contract *FeeQuoter
}

type FeeQuoterCallerRaw struct {
	Contract *FeeQuoterCaller
}

type FeeQuoterTransactorRaw struct {
	Contract *FeeQuoterTransactor
}

func NewFeeQuoter(address common.Address, backend bind.ContractBackend) (*FeeQuoter, error) {
	abi, err := abi.JSON(strings.NewReader(FeeQuoterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindFeeQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeQuoter{address: address, abi: abi, FeeQuoterCaller: FeeQuoterCaller{contract: contract}, FeeQuoterTransactor: FeeQuoterTransactor{contract: contract}, FeeQuoterFilterer: FeeQuoterFilterer{contract: contract}}, nil
}

func NewFeeQuoterCaller(address common.Address, caller bind.ContractCaller) (*FeeQuoterCaller, error) {
	contract, err := bindFeeQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterCaller{contract: contract}, nil
}

func NewFeeQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeQuoterTransactor, error) {
	contract, err := bindFeeQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterTransactor{contract: contract}, nil
}

func NewFeeQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeQuoterFilterer, error) {
	contract, err := bindFeeQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterFilterer{contract: contract}, nil
}

func bindFeeQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeQuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_FeeQuoter *FeeQuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeQuoter.Contract.FeeQuoterCaller.contract.Call(opts, result, method, params...)
}

func (_FeeQuoter *FeeQuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeQuoter.Contract.FeeQuoterTransactor.contract.Transfer(opts)
}

func (_FeeQuoter *FeeQuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeQuoter.Contract.FeeQuoterTransactor.contract.Transact(opts, method, params...)
}

func (_FeeQuoter *FeeQuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeQuoter.Contract.contract.Call(opts, result, method, params...)
}

func (_FeeQuoter *FeeQuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeQuoter.Contract.contract.Transfer(opts)
}

func (_FeeQuoter *FeeQuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeQuoter.Contract.contract.Transact(opts, method, params...)
}

func (_FeeQuoter *FeeQuoterCaller) FEEBASEDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "FEE_BASE_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) FEEBASEDECIMALS() (*big.Int, error) {
	return _FeeQuoter.Contract.FEEBASEDECIMALS(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) FEEBASEDECIMALS() (*big.Int, error) {
	return _FeeQuoter.Contract.FEEBASEDECIMALS(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) KEYSTONEPRICEDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "KEYSTONE_PRICE_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) KEYSTONEPRICEDECIMALS() (*big.Int, error) {
	return _FeeQuoter.Contract.KEYSTONEPRICEDECIMALS(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) KEYSTONEPRICEDECIMALS() (*big.Int, error) {
	return _FeeQuoter.Contract.KEYSTONEPRICEDECIMALS(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) ConvertTokenAmount(opts *bind.CallOpts, fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "convertTokenAmount", fromToken, fromTokenAmount, toToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) ConvertTokenAmount(fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error) {
	return _FeeQuoter.Contract.ConvertTokenAmount(&_FeeQuoter.CallOpts, fromToken, fromTokenAmount, toToken)
}

func (_FeeQuoter *FeeQuoterCallerSession) ConvertTokenAmount(fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error) {
	return _FeeQuoter.Contract.ConvertTokenAmount(&_FeeQuoter.CallOpts, fromToken, fromTokenAmount, toToken)
}

func (_FeeQuoter *FeeQuoterCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _FeeQuoter.Contract.GetAllAuthorizedCallers(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _FeeQuoter.Contract.GetAllAuthorizedCallers(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (FeeQuoterDestChainConfig, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	if err != nil {
		return *new(FeeQuoterDestChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeQuoterDestChainConfig)).(*FeeQuoterDestChainConfig)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetDestChainConfig(destChainSelector uint64) (FeeQuoterDestChainConfig, error) {
	return _FeeQuoter.Contract.GetDestChainConfig(&_FeeQuoter.CallOpts, destChainSelector)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetDestChainConfig(destChainSelector uint64) (FeeQuoterDestChainConfig, error) {
	return _FeeQuoter.Contract.GetDestChainConfig(&_FeeQuoter.CallOpts, destChainSelector)
}

func (_FeeQuoter *FeeQuoterCaller) GetDestinationChainGasPrice(opts *bind.CallOpts, destChainSelector uint64) (InternalTimestampedPackedUint224, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getDestinationChainGasPrice", destChainSelector)

	if err != nil {
		return *new(InternalTimestampedPackedUint224), err
	}

	out0 := *abi.ConvertType(out[0], new(InternalTimestampedPackedUint224)).(*InternalTimestampedPackedUint224)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetDestinationChainGasPrice(destChainSelector uint64) (InternalTimestampedPackedUint224, error) {
	return _FeeQuoter.Contract.GetDestinationChainGasPrice(&_FeeQuoter.CallOpts, destChainSelector)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetDestinationChainGasPrice(destChainSelector uint64) (InternalTimestampedPackedUint224, error) {
	return _FeeQuoter.Contract.GetDestinationChainGasPrice(&_FeeQuoter.CallOpts, destChainSelector)
}

func (_FeeQuoter *FeeQuoterCaller) GetFeeTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getFeeTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetFeeTokens() ([]common.Address, error) {
	return _FeeQuoter.Contract.GetFeeTokens(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetFeeTokens() ([]common.Address, error) {
	return _FeeQuoter.Contract.GetFeeTokens(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) GetPremiumMultiplierWeiPerEth(opts *bind.CallOpts, token common.Address) (uint64, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getPremiumMultiplierWeiPerEth", token)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetPremiumMultiplierWeiPerEth(token common.Address) (uint64, error) {
	return _FeeQuoter.Contract.GetPremiumMultiplierWeiPerEth(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetPremiumMultiplierWeiPerEth(token common.Address) (uint64, error) {
	return _FeeQuoter.Contract.GetPremiumMultiplierWeiPerEth(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCaller) GetStaticConfig(opts *bind.CallOpts) (FeeQuoterStaticConfig, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(FeeQuoterStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeQuoterStaticConfig)).(*FeeQuoterStaticConfig)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetStaticConfig() (FeeQuoterStaticConfig, error) {
	return _FeeQuoter.Contract.GetStaticConfig(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetStaticConfig() (FeeQuoterStaticConfig, error) {
	return _FeeQuoter.Contract.GetStaticConfig(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) GetTokenAndGasPrices(opts *bind.CallOpts, token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

	error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getTokenAndGasPrices", token, destChainSelector)

	outstruct := new(GetTokenAndGasPrices)
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GasPriceValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_FeeQuoter *FeeQuoterSession) GetTokenAndGasPrices(token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

	error) {
	return _FeeQuoter.Contract.GetTokenAndGasPrices(&_FeeQuoter.CallOpts, token, destChainSelector)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetTokenAndGasPrices(token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

	error) {
	return _FeeQuoter.Contract.GetTokenAndGasPrices(&_FeeQuoter.CallOpts, token, destChainSelector)
}

func (_FeeQuoter *FeeQuoterCaller) GetTokenPrice(opts *bind.CallOpts, token common.Address) (InternalTimestampedPackedUint224, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getTokenPrice", token)

	if err != nil {
		return *new(InternalTimestampedPackedUint224), err
	}

	out0 := *abi.ConvertType(out[0], new(InternalTimestampedPackedUint224)).(*InternalTimestampedPackedUint224)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetTokenPrice(token common.Address) (InternalTimestampedPackedUint224, error) {
	return _FeeQuoter.Contract.GetTokenPrice(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetTokenPrice(token common.Address) (InternalTimestampedPackedUint224, error) {
	return _FeeQuoter.Contract.GetTokenPrice(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCaller) GetTokenPriceFeedConfig(opts *bind.CallOpts, token common.Address) (FeeQuoterTokenPriceFeedConfig, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getTokenPriceFeedConfig", token)

	if err != nil {
		return *new(FeeQuoterTokenPriceFeedConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeQuoterTokenPriceFeedConfig)).(*FeeQuoterTokenPriceFeedConfig)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetTokenPriceFeedConfig(token common.Address) (FeeQuoterTokenPriceFeedConfig, error) {
	return _FeeQuoter.Contract.GetTokenPriceFeedConfig(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetTokenPriceFeedConfig(token common.Address) (FeeQuoterTokenPriceFeedConfig, error) {
	return _FeeQuoter.Contract.GetTokenPriceFeedConfig(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCaller) GetTokenPrices(opts *bind.CallOpts, tokens []common.Address) ([]InternalTimestampedPackedUint224, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getTokenPrices", tokens)

	if err != nil {
		return *new([]InternalTimestampedPackedUint224), err
	}

	out0 := *abi.ConvertType(out[0], new([]InternalTimestampedPackedUint224)).(*[]InternalTimestampedPackedUint224)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetTokenPrices(tokens []common.Address) ([]InternalTimestampedPackedUint224, error) {
	return _FeeQuoter.Contract.GetTokenPrices(&_FeeQuoter.CallOpts, tokens)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetTokenPrices(tokens []common.Address) ([]InternalTimestampedPackedUint224, error) {
	return _FeeQuoter.Contract.GetTokenPrices(&_FeeQuoter.CallOpts, tokens)
}

func (_FeeQuoter *FeeQuoterCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (FeeQuoterTokenTransferFeeConfig, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getTokenTransferFeeConfig", destChainSelector, token)

	if err != nil {
		return *new(FeeQuoterTokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FeeQuoterTokenTransferFeeConfig)).(*FeeQuoterTokenTransferFeeConfig)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetTokenTransferFeeConfig(destChainSelector uint64, token common.Address) (FeeQuoterTokenTransferFeeConfig, error) {
	return _FeeQuoter.Contract.GetTokenTransferFeeConfig(&_FeeQuoter.CallOpts, destChainSelector, token)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetTokenTransferFeeConfig(destChainSelector uint64, token common.Address) (FeeQuoterTokenTransferFeeConfig, error) {
	return _FeeQuoter.Contract.GetTokenTransferFeeConfig(&_FeeQuoter.CallOpts, destChainSelector, token)
}

func (_FeeQuoter *FeeQuoterCaller) GetValidatedFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getValidatedFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetValidatedFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _FeeQuoter.Contract.GetValidatedFee(&_FeeQuoter.CallOpts, destChainSelector, message)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetValidatedFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _FeeQuoter.Contract.GetValidatedFee(&_FeeQuoter.CallOpts, destChainSelector, message)
}

func (_FeeQuoter *FeeQuoterCaller) GetValidatedTokenPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "getValidatedTokenPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) GetValidatedTokenPrice(token common.Address) (*big.Int, error) {
	return _FeeQuoter.Contract.GetValidatedTokenPrice(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCallerSession) GetValidatedTokenPrice(token common.Address) (*big.Int, error) {
	return _FeeQuoter.Contract.GetValidatedTokenPrice(&_FeeQuoter.CallOpts, token)
}

func (_FeeQuoter *FeeQuoterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) Owner() (common.Address, error) {
	return _FeeQuoter.Contract.Owner(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) Owner() (common.Address, error) {
	return _FeeQuoter.Contract.Owner(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCaller) ProcessMessageArgs(opts *bind.CallOpts, destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte, onRampTokenTransfers []InternalEVM2AnyTokenTransfer, sourceTokenAmounts []ClientEVMTokenAmount) (ProcessMessageArgs,

	error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "processMessageArgs", destChainSelector, feeToken, feeTokenAmount, extraArgs, onRampTokenTransfers, sourceTokenAmounts)

	outstruct := new(ProcessMessageArgs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MsgFeeJuels = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IsOutOfOrderExecution = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.ConvertedExtraArgs = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.DestExecDataPerToken = *abi.ConvertType(out[3], new([][]byte)).(*[][]byte)

	return *outstruct, err

}

func (_FeeQuoter *FeeQuoterSession) ProcessMessageArgs(destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte, onRampTokenTransfers []InternalEVM2AnyTokenTransfer, sourceTokenAmounts []ClientEVMTokenAmount) (ProcessMessageArgs,

	error) {
	return _FeeQuoter.Contract.ProcessMessageArgs(&_FeeQuoter.CallOpts, destChainSelector, feeToken, feeTokenAmount, extraArgs, onRampTokenTransfers, sourceTokenAmounts)
}

func (_FeeQuoter *FeeQuoterCallerSession) ProcessMessageArgs(destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte, onRampTokenTransfers []InternalEVM2AnyTokenTransfer, sourceTokenAmounts []ClientEVMTokenAmount) (ProcessMessageArgs,

	error) {
	return _FeeQuoter.Contract.ProcessMessageArgs(&_FeeQuoter.CallOpts, destChainSelector, feeToken, feeTokenAmount, extraArgs, onRampTokenTransfers, sourceTokenAmounts)
}

func (_FeeQuoter *FeeQuoterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FeeQuoter.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_FeeQuoter *FeeQuoterSession) TypeAndVersion() (string, error) {
	return _FeeQuoter.Contract.TypeAndVersion(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterCallerSession) TypeAndVersion() (string, error) {
	return _FeeQuoter.Contract.TypeAndVersion(&_FeeQuoter.CallOpts)
}

func (_FeeQuoter *FeeQuoterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "acceptOwnership")
}

func (_FeeQuoter *FeeQuoterSession) AcceptOwnership() (*types.Transaction, error) {
	return _FeeQuoter.Contract.AcceptOwnership(&_FeeQuoter.TransactOpts)
}

func (_FeeQuoter *FeeQuoterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FeeQuoter.Contract.AcceptOwnership(&_FeeQuoter.TransactOpts)
}

func (_FeeQuoter *FeeQuoterTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_FeeQuoter *FeeQuoterSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyAuthorizedCallerUpdates(&_FeeQuoter.TransactOpts, authorizedCallerArgs)
}

func (_FeeQuoter *FeeQuoterTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyAuthorizedCallerUpdates(&_FeeQuoter.TransactOpts, authorizedCallerArgs)
}

func (_FeeQuoter *FeeQuoterTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_FeeQuoter *FeeQuoterSession) ApplyDestChainConfigUpdates(destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyDestChainConfigUpdates(&_FeeQuoter.TransactOpts, destChainConfigArgs)
}

func (_FeeQuoter *FeeQuoterTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyDestChainConfigUpdates(&_FeeQuoter.TransactOpts, destChainConfigArgs)
}

func (_FeeQuoter *FeeQuoterTransactor) ApplyFeeTokensUpdates(opts *bind.TransactOpts, feeTokensToAdd []common.Address, feeTokensToRemove []common.Address) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "applyFeeTokensUpdates", feeTokensToAdd, feeTokensToRemove)
}

func (_FeeQuoter *FeeQuoterSession) ApplyFeeTokensUpdates(feeTokensToAdd []common.Address, feeTokensToRemove []common.Address) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyFeeTokensUpdates(&_FeeQuoter.TransactOpts, feeTokensToAdd, feeTokensToRemove)
}

func (_FeeQuoter *FeeQuoterTransactorSession) ApplyFeeTokensUpdates(feeTokensToAdd []common.Address, feeTokensToRemove []common.Address) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyFeeTokensUpdates(&_FeeQuoter.TransactOpts, feeTokensToAdd, feeTokensToRemove)
}

func (_FeeQuoter *FeeQuoterTransactor) ApplyPremiumMultiplierWeiPerEthUpdates(opts *bind.TransactOpts, premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "applyPremiumMultiplierWeiPerEthUpdates", premiumMultiplierWeiPerEthArgs)
}

func (_FeeQuoter *FeeQuoterSession) ApplyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyPremiumMultiplierWeiPerEthUpdates(&_FeeQuoter.TransactOpts, premiumMultiplierWeiPerEthArgs)
}

func (_FeeQuoter *FeeQuoterTransactorSession) ApplyPremiumMultiplierWeiPerEthUpdates(premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyPremiumMultiplierWeiPerEthUpdates(&_FeeQuoter.TransactOpts, premiumMultiplierWeiPerEthArgs)
}

func (_FeeQuoter *FeeQuoterTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_FeeQuoter *FeeQuoterSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyTokenTransferFeeConfigUpdates(&_FeeQuoter.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_FeeQuoter *FeeQuoterTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error) {
	return _FeeQuoter.Contract.ApplyTokenTransferFeeConfigUpdates(&_FeeQuoter.TransactOpts, tokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs)
}

func (_FeeQuoter *FeeQuoterTransactor) OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "onReport", metadata, report)
}

func (_FeeQuoter *FeeQuoterSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _FeeQuoter.Contract.OnReport(&_FeeQuoter.TransactOpts, metadata, report)
}

func (_FeeQuoter *FeeQuoterTransactorSession) OnReport(metadata []byte, report []byte) (*types.Transaction, error) {
	return _FeeQuoter.Contract.OnReport(&_FeeQuoter.TransactOpts, metadata, report)
}

func (_FeeQuoter *FeeQuoterTransactor) SetReportPermissions(opts *bind.TransactOpts, permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "setReportPermissions", permissions)
}

func (_FeeQuoter *FeeQuoterSession) SetReportPermissions(permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _FeeQuoter.Contract.SetReportPermissions(&_FeeQuoter.TransactOpts, permissions)
}

func (_FeeQuoter *FeeQuoterTransactorSession) SetReportPermissions(permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _FeeQuoter.Contract.SetReportPermissions(&_FeeQuoter.TransactOpts, permissions)
}

func (_FeeQuoter *FeeQuoterTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "transferOwnership", to)
}

func (_FeeQuoter *FeeQuoterSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FeeQuoter.Contract.TransferOwnership(&_FeeQuoter.TransactOpts, to)
}

func (_FeeQuoter *FeeQuoterTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FeeQuoter.Contract.TransferOwnership(&_FeeQuoter.TransactOpts, to)
}

func (_FeeQuoter *FeeQuoterTransactor) UpdatePrices(opts *bind.TransactOpts, priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "updatePrices", priceUpdates)
}

func (_FeeQuoter *FeeQuoterSession) UpdatePrices(priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _FeeQuoter.Contract.UpdatePrices(&_FeeQuoter.TransactOpts, priceUpdates)
}

func (_FeeQuoter *FeeQuoterTransactorSession) UpdatePrices(priceUpdates InternalPriceUpdates) (*types.Transaction, error) {
	return _FeeQuoter.Contract.UpdatePrices(&_FeeQuoter.TransactOpts, priceUpdates)
}

func (_FeeQuoter *FeeQuoterTransactor) UpdateTokenPriceFeeds(opts *bind.TransactOpts, tokenPriceFeedUpdates []FeeQuoterTokenPriceFeedUpdate) (*types.Transaction, error) {
	return _FeeQuoter.contract.Transact(opts, "updateTokenPriceFeeds", tokenPriceFeedUpdates)
}

func (_FeeQuoter *FeeQuoterSession) UpdateTokenPriceFeeds(tokenPriceFeedUpdates []FeeQuoterTokenPriceFeedUpdate) (*types.Transaction, error) {
	return _FeeQuoter.Contract.UpdateTokenPriceFeeds(&_FeeQuoter.TransactOpts, tokenPriceFeedUpdates)
}

func (_FeeQuoter *FeeQuoterTransactorSession) UpdateTokenPriceFeeds(tokenPriceFeedUpdates []FeeQuoterTokenPriceFeedUpdate) (*types.Transaction, error) {
	return _FeeQuoter.Contract.UpdateTokenPriceFeeds(&_FeeQuoter.TransactOpts, tokenPriceFeedUpdates)
}

type FeeQuoterAuthorizedCallerAddedIterator struct {
	Event *FeeQuoterAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterAuthorizedCallerAdded)
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
		it.Event = new(FeeQuoterAuthorizedCallerAdded)
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

func (it *FeeQuoterAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*FeeQuoterAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &FeeQuoterAuthorizedCallerAddedIterator{contract: _FeeQuoter.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterAuthorizedCallerAdded)
				if err := _FeeQuoter.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseAuthorizedCallerAdded(log types.Log) (*FeeQuoterAuthorizedCallerAdded, error) {
	event := new(FeeQuoterAuthorizedCallerAdded)
	if err := _FeeQuoter.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterAuthorizedCallerRemovedIterator struct {
	Event *FeeQuoterAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterAuthorizedCallerRemoved)
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
		it.Event = new(FeeQuoterAuthorizedCallerRemoved)
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

func (it *FeeQuoterAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*FeeQuoterAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &FeeQuoterAuthorizedCallerRemovedIterator{contract: _FeeQuoter.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *FeeQuoterAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterAuthorizedCallerRemoved)
				if err := _FeeQuoter.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*FeeQuoterAuthorizedCallerRemoved, error) {
	event := new(FeeQuoterAuthorizedCallerRemoved)
	if err := _FeeQuoter.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterDestChainAddedIterator struct {
	Event *FeeQuoterDestChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterDestChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterDestChainAdded)
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
		it.Event = new(FeeQuoterDestChainAdded)
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

func (it *FeeQuoterDestChainAddedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterDestChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterDestChainAdded struct {
	DestChainSelector uint64
	DestChainConfig   FeeQuoterDestChainConfig
	Raw               types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*FeeQuoterDestChainAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterDestChainAddedIterator{contract: _FeeQuoter.contract, event: "DestChainAdded", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterDestChainAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterDestChainAdded)
				if err := _FeeQuoter.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseDestChainAdded(log types.Log) (*FeeQuoterDestChainAdded, error) {
	event := new(FeeQuoterDestChainAdded)
	if err := _FeeQuoter.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterDestChainConfigUpdatedIterator struct {
	Event *FeeQuoterDestChainConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterDestChainConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterDestChainConfigUpdated)
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
		it.Event = new(FeeQuoterDestChainConfigUpdated)
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

func (it *FeeQuoterDestChainConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterDestChainConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterDestChainConfigUpdated struct {
	DestChainSelector uint64
	DestChainConfig   FeeQuoterDestChainConfig
	Raw               types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterDestChainConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*FeeQuoterDestChainConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "DestChainConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterDestChainConfigUpdatedIterator{contract: _FeeQuoter.contract, event: "DestChainConfigUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchDestChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterDestChainConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "DestChainConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterDestChainConfigUpdated)
				if err := _FeeQuoter.contract.UnpackLog(event, "DestChainConfigUpdated", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseDestChainConfigUpdated(log types.Log) (*FeeQuoterDestChainConfigUpdated, error) {
	event := new(FeeQuoterDestChainConfigUpdated)
	if err := _FeeQuoter.contract.UnpackLog(event, "DestChainConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterFeeTokenAddedIterator struct {
	Event *FeeQuoterFeeTokenAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterFeeTokenAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterFeeTokenAdded)
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
		it.Event = new(FeeQuoterFeeTokenAdded)
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

func (it *FeeQuoterFeeTokenAddedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterFeeTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterFeeTokenAdded struct {
	FeeToken common.Address
	Raw      types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterFeeTokenAdded(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterFeeTokenAddedIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "FeeTokenAdded", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterFeeTokenAddedIterator{contract: _FeeQuoter.contract, event: "FeeTokenAdded", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchFeeTokenAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterFeeTokenAdded, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "FeeTokenAdded", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterFeeTokenAdded)
				if err := _FeeQuoter.contract.UnpackLog(event, "FeeTokenAdded", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseFeeTokenAdded(log types.Log) (*FeeQuoterFeeTokenAdded, error) {
	event := new(FeeQuoterFeeTokenAdded)
	if err := _FeeQuoter.contract.UnpackLog(event, "FeeTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterFeeTokenRemovedIterator struct {
	Event *FeeQuoterFeeTokenRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterFeeTokenRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterFeeTokenRemoved)
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
		it.Event = new(FeeQuoterFeeTokenRemoved)
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

func (it *FeeQuoterFeeTokenRemovedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterFeeTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterFeeTokenRemoved struct {
	FeeToken common.Address
	Raw      types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterFeeTokenRemoved(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterFeeTokenRemovedIterator, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "FeeTokenRemoved", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterFeeTokenRemovedIterator{contract: _FeeQuoter.contract, event: "FeeTokenRemoved", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchFeeTokenRemoved(opts *bind.WatchOpts, sink chan<- *FeeQuoterFeeTokenRemoved, feeToken []common.Address) (event.Subscription, error) {

	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "FeeTokenRemoved", feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterFeeTokenRemoved)
				if err := _FeeQuoter.contract.UnpackLog(event, "FeeTokenRemoved", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseFeeTokenRemoved(log types.Log) (*FeeQuoterFeeTokenRemoved, error) {
	event := new(FeeQuoterFeeTokenRemoved)
	if err := _FeeQuoter.contract.UnpackLog(event, "FeeTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterOwnershipTransferRequestedIterator struct {
	Event *FeeQuoterOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterOwnershipTransferRequested)
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
		it.Event = new(FeeQuoterOwnershipTransferRequested)
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

func (it *FeeQuoterOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterOwnershipTransferRequestedIterator{contract: _FeeQuoter.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FeeQuoterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterOwnershipTransferRequested)
				if err := _FeeQuoter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseOwnershipTransferRequested(log types.Log) (*FeeQuoterOwnershipTransferRequested, error) {
	event := new(FeeQuoterOwnershipTransferRequested)
	if err := _FeeQuoter.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterOwnershipTransferredIterator struct {
	Event *FeeQuoterOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterOwnershipTransferred)
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
		it.Event = new(FeeQuoterOwnershipTransferred)
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

func (it *FeeQuoterOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterOwnershipTransferredIterator{contract: _FeeQuoter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeQuoterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterOwnershipTransferred)
				if err := _FeeQuoter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseOwnershipTransferred(log types.Log) (*FeeQuoterOwnershipTransferred, error) {
	event := new(FeeQuoterOwnershipTransferred)
	if err := _FeeQuoter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator struct {
	Event *FeeQuoterPremiumMultiplierWeiPerEthUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterPremiumMultiplierWeiPerEthUpdated)
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
		it.Event = new(FeeQuoterPremiumMultiplierWeiPerEthUpdated)
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

func (it *FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterPremiumMultiplierWeiPerEthUpdated struct {
	Token                      common.Address
	PremiumMultiplierWeiPerEth uint64
	Raw                        types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterPremiumMultiplierWeiPerEthUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "PremiumMultiplierWeiPerEthUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator{contract: _FeeQuoter.contract, event: "PremiumMultiplierWeiPerEthUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchPremiumMultiplierWeiPerEthUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterPremiumMultiplierWeiPerEthUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "PremiumMultiplierWeiPerEthUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterPremiumMultiplierWeiPerEthUpdated)
				if err := _FeeQuoter.contract.UnpackLog(event, "PremiumMultiplierWeiPerEthUpdated", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParsePremiumMultiplierWeiPerEthUpdated(log types.Log) (*FeeQuoterPremiumMultiplierWeiPerEthUpdated, error) {
	event := new(FeeQuoterPremiumMultiplierWeiPerEthUpdated)
	if err := _FeeQuoter.contract.UnpackLog(event, "PremiumMultiplierWeiPerEthUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterPriceFeedPerTokenUpdatedIterator struct {
	Event *FeeQuoterPriceFeedPerTokenUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterPriceFeedPerTokenUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterPriceFeedPerTokenUpdated)
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
		it.Event = new(FeeQuoterPriceFeedPerTokenUpdated)
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

func (it *FeeQuoterPriceFeedPerTokenUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterPriceFeedPerTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterPriceFeedPerTokenUpdated struct {
	Token           common.Address
	PriceFeedConfig FeeQuoterTokenPriceFeedConfig
	Raw             types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterPriceFeedPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterPriceFeedPerTokenUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "PriceFeedPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterPriceFeedPerTokenUpdatedIterator{contract: _FeeQuoter.contract, event: "PriceFeedPerTokenUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchPriceFeedPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterPriceFeedPerTokenUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "PriceFeedPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterPriceFeedPerTokenUpdated)
				if err := _FeeQuoter.contract.UnpackLog(event, "PriceFeedPerTokenUpdated", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParsePriceFeedPerTokenUpdated(log types.Log) (*FeeQuoterPriceFeedPerTokenUpdated, error) {
	event := new(FeeQuoterPriceFeedPerTokenUpdated)
	if err := _FeeQuoter.contract.UnpackLog(event, "PriceFeedPerTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterReportPermissionSetIterator struct {
	Event *FeeQuoterReportPermissionSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterReportPermissionSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterReportPermissionSet)
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
		it.Event = new(FeeQuoterReportPermissionSet)
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

func (it *FeeQuoterReportPermissionSetIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterReportPermissionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterReportPermissionSet struct {
	ReportId   [32]byte
	Permission KeystoneFeedsPermissionHandlerPermission
	Raw        types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterReportPermissionSet(opts *bind.FilterOpts, reportId [][32]byte) (*FeeQuoterReportPermissionSetIterator, error) {

	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "ReportPermissionSet", reportIdRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterReportPermissionSetIterator{contract: _FeeQuoter.contract, event: "ReportPermissionSet", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchReportPermissionSet(opts *bind.WatchOpts, sink chan<- *FeeQuoterReportPermissionSet, reportId [][32]byte) (event.Subscription, error) {

	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "ReportPermissionSet", reportIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterReportPermissionSet)
				if err := _FeeQuoter.contract.UnpackLog(event, "ReportPermissionSet", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseReportPermissionSet(log types.Log) (*FeeQuoterReportPermissionSet, error) {
	event := new(FeeQuoterReportPermissionSet)
	if err := _FeeQuoter.contract.UnpackLog(event, "ReportPermissionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterTokenTransferFeeConfigDeletedIterator struct {
	Event *FeeQuoterTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterTokenTransferFeeConfigDeleted)
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
		it.Event = new(FeeQuoterTokenTransferFeeConfigDeleted)
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

func (it *FeeQuoterTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Token             common.Address
	Raw               types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*FeeQuoterTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterTokenTransferFeeConfigDeletedIterator{contract: _FeeQuoter.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *FeeQuoterTokenTransferFeeConfigDeleted, destChainSelector []uint64, token []common.Address) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterTokenTransferFeeConfigDeleted)
				if err := _FeeQuoter.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*FeeQuoterTokenTransferFeeConfigDeleted, error) {
	event := new(FeeQuoterTokenTransferFeeConfigDeleted)
	if err := _FeeQuoter.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterTokenTransferFeeConfigUpdatedIterator struct {
	Event *FeeQuoterTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterTokenTransferFeeConfigUpdated)
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
		it.Event = new(FeeQuoterTokenTransferFeeConfigUpdated)
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

func (it *FeeQuoterTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	Token                  common.Address
	TokenTransferFeeConfig FeeQuoterTokenTransferFeeConfig
	Raw                    types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*FeeQuoterTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterTokenTransferFeeConfigUpdatedIterator{contract: _FeeQuoter.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterTokenTransferFeeConfigUpdated, destChainSelector []uint64, token []common.Address) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterTokenTransferFeeConfigUpdated)
				if err := _FeeQuoter.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*FeeQuoterTokenTransferFeeConfigUpdated, error) {
	event := new(FeeQuoterTokenTransferFeeConfigUpdated)
	if err := _FeeQuoter.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterUsdPerTokenUpdatedIterator struct {
	Event *FeeQuoterUsdPerTokenUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterUsdPerTokenUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterUsdPerTokenUpdated)
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
		it.Event = new(FeeQuoterUsdPerTokenUpdated)
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

func (it *FeeQuoterUsdPerTokenUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterUsdPerTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterUsdPerTokenUpdated struct {
	Token     common.Address
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterUsdPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterUsdPerTokenUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "UsdPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterUsdPerTokenUpdatedIterator{contract: _FeeQuoter.contract, event: "UsdPerTokenUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchUsdPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterUsdPerTokenUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "UsdPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterUsdPerTokenUpdated)
				if err := _FeeQuoter.contract.UnpackLog(event, "UsdPerTokenUpdated", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseUsdPerTokenUpdated(log types.Log) (*FeeQuoterUsdPerTokenUpdated, error) {
	event := new(FeeQuoterUsdPerTokenUpdated)
	if err := _FeeQuoter.contract.UnpackLog(event, "UsdPerTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type FeeQuoterUsdPerUnitGasUpdatedIterator struct {
	Event *FeeQuoterUsdPerUnitGasUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *FeeQuoterUsdPerUnitGasUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeQuoterUsdPerUnitGasUpdated)
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
		it.Event = new(FeeQuoterUsdPerUnitGasUpdated)
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

func (it *FeeQuoterUsdPerUnitGasUpdatedIterator) Error() error {
	return it.fail
}

func (it *FeeQuoterUsdPerUnitGasUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type FeeQuoterUsdPerUnitGasUpdated struct {
	DestChain uint64
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log
}

func (_FeeQuoter *FeeQuoterFilterer) FilterUsdPerUnitGasUpdated(opts *bind.FilterOpts, destChain []uint64) (*FeeQuoterUsdPerUnitGasUpdatedIterator, error) {

	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _FeeQuoter.contract.FilterLogs(opts, "UsdPerUnitGasUpdated", destChainRule)
	if err != nil {
		return nil, err
	}
	return &FeeQuoterUsdPerUnitGasUpdatedIterator{contract: _FeeQuoter.contract, event: "UsdPerUnitGasUpdated", logs: logs, sub: sub}, nil
}

func (_FeeQuoter *FeeQuoterFilterer) WatchUsdPerUnitGasUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterUsdPerUnitGasUpdated, destChain []uint64) (event.Subscription, error) {

	var destChainRule []interface{}
	for _, destChainItem := range destChain {
		destChainRule = append(destChainRule, destChainItem)
	}

	logs, sub, err := _FeeQuoter.contract.WatchLogs(opts, "UsdPerUnitGasUpdated", destChainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(FeeQuoterUsdPerUnitGasUpdated)
				if err := _FeeQuoter.contract.UnpackLog(event, "UsdPerUnitGasUpdated", log); err != nil {
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

func (_FeeQuoter *FeeQuoterFilterer) ParseUsdPerUnitGasUpdated(log types.Log) (*FeeQuoterUsdPerUnitGasUpdated, error) {
	event := new(FeeQuoterUsdPerUnitGasUpdated)
	if err := _FeeQuoter.contract.UnpackLog(event, "UsdPerUnitGasUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

var FeeQuoterZkBin string = ("0x0003000000000002002c0000000000020002000000010355000000600310027000000a1a0030019d00000a1a033001970000000100200190000000bf0000c13d0000008002000039000000400020043f000000040030008c000000e00000413d000000000401043b000000e00440027000000a620040009c000000e20000213d00000a780040009c000002460000a13d00000a790040009c000002ab0000213d00000a7f0040009c000005910000213d00000a820040009c000007130000613d00000a830040009c000000e00000c13d000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000402100370000000000202043b00000a200020009c000000e00000213d0000002304200039000000000034004b000000e00000813d0000000404200039000000000441034f000000000504043b00000a200050009c000001140000213d00000005045002100000003f0440003900000a210440019700000ab10040009c000001140000213d0000008004400039000000400040043f000000800050043f0000002402200039000000a0045000c90000000004240019000000000034004b000000e00000213d000000000005004b00000ce60000c13d000000000100041a00000a1f011001970000000002000411000000000012004b00000af40000c13d000000800100043d000000000001004b00000a020000613d002400000000001d00000024010000290000000501100210000000a00110003900000000040104330000004001400039002100000001001d000000000101043300000ab3021001970000000053040434002300000004001d0000006001400039002200000001001d0000000004010433002000000005001d0000000005050433000000400100043d0000008006100039000000000026043500000ab4025001970000006005100039000000000025043500000a1f024001970000004004100039000000000024043500000a1f03300197000000200210003900000000003204350000008003000039000000000031043500000ab50010009c000001140000213d000000a003100039000000400030043f00000a1a0020009c00000a1a020080410000004002200210000000000101043300000a1a0010009c00000a1a010080410000006001100210000000000121019f000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a31011001c700008010020000392863285e0000040f00000001002001900000002303000029000000e00000613d000000000101043b0000008002300039001e00000002001d0000000002020433001d00000002001d001f00000001001d000000000010043f0000000401000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f00000023040000290000000100200190000000e00000613d000000000101043b000000000201041a00000acd022001970000001d0000006b000000010220c1bf000000000021041b000000000104043300000a1f01100197000000400200043d00000000011204360000002003000029000000000303043300000ab40330019700000000003104350000002101000029000000000101043300000ab301100197000000400320003900000000001304350000002201000029000000000101043300000a1f01100197000000600320003900000000001304350000001e010000290000000001010433000000000001004b0000000001000039000000010100c0390000008003200039000000000013043500000a1a0020009c00000a1a020080410000004001200210000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000ac5011001c70000800d02000039000000020300003900000ac6040000410000001f05000029286328590000040f0000000100200190000000e00000613d0000002402000029002400010020003d000000800100043d000000240010006b000000430000413d00000a020000013d000000e004000039000000400040043f0000000002000416000000000002004b000000e00000c13d0000001f0230003900000a1b02200197000000e002200039000000400020043f0000001f0530018f00000a1c06300198000000e002600039000000d10000613d000000000701034f000000007807043c0000000004840436000000000024004b000000cd0000c13d000000000005004b000000de0000613d000000000161034f0000000304500210000000000502043300000000054501cf000000000545022f000000000101043b0000010004400089000000000141022f00000000014101cf000000000151019f0000000000120435000001200030008c000001100000813d0000000001000019000028650001043000000a630040009c000002620000a13d00000a640040009c000002d10000213d00000a6a0040009c000005d90000213d00000a6d0040009c0000077a0000613d00000a6e0040009c000000e00000c13d000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000401100370000000000101043b002400000001001d00000a1f0010009c000000e00000213d286324230000040f0000002401000029000000000010043f0000000701000039000000200010043f00000040020000390000000001000019286328260000040f002400000001001d000000400100043d002300000001001d28631fe70000040f0000002401000029000000000101041a000000a002100270000000ff0220018f00000023040000290000002003400039000000000023043500000a1f0110019700000000001404350000000002040019000000400100043d002400000001001d286320b90000040f00000ac50000013d000000400100043d002400000001001d00000a1d0010009c0000011a0000a13d00000abb01000041000000000010043f0000004101000039000000040010043f00000a4e01000041000028650001043000000024010000290000006001100039000000400010043f000000e00100043d00000a1e0010009c000000e00000213d00000024020000290000000001120436002300000001001d000001000100043d00000a1f0010009c000000e00000213d00000023020000290000000000120435000001200100043d00000a1a0010009c000000e00000213d00000024020000290000004002200039002200000002001d0000000000120435000001400200043d00000a200020009c000000e00000213d000000e001300039000000ff04200039000000000014004b000000e00000813d000000e004200039000000000404043300000a200040009c000001140000213d00000005054002100000003f0650003900000a2106600197000000400700043d0000000006670019002100000007001d000000000076004b0000000007000039000000010700403900000a200060009c000001140000213d0000000100700190000001140000c13d000000400060043f0000002106000029000000000046043500000100022000390000000005250019000000000015004b000000e00000213d000000000004004b000001580000613d0000002104000029000000002602043400000a1f0060009c000000e00000213d00000020044000390000000000640435000000000052004b000001510000413d000001600200043d00000a200020009c000000e00000213d0000001f04200039000000000034004b000000000500001900000a220500804100000a2204400197000000000004004b000000000600001900000a220600404100000a220040009c000000000605c019000000000006004b000000e00000c13d000000e004200039000000000404043300000a200040009c000001140000213d00000005054002100000003f0650003900000a2106600197000000400700043d0000000006670019002000000007001d000000000076004b0000000007000039000000010700403900000a200060009c000001140000213d0000000100700190000001140000c13d000000400060043f00000020060000290000000006460436001c00000006001d00000100022000390000000005250019000000000015004b000000e00000213d000000000004004b0000018a0000613d0000002004000029000000002602043400000a1f0060009c000000e00000213d00000020044000390000000000640435000000000052004b000001830000413d000001800200043d00000a200020009c000000e00000213d0000001f04200039000000000034004b000000000500001900000a220500804100000a2204400197000000000004004b000000000600001900000a220600404100000a220040009c000000000605c019000000000006004b000000e00000c13d000000e004200039000000000504043300000a200050009c000001140000213d00000005045002100000003f0440003900000a2104400197000000400600043d0000000004460019001b00000006001d000000000064004b0000000006000039000000010600403900000a200040009c000001140000213d0000000100600190000001140000c13d000000400040043f0000001b040000290000000004540436001a00000004001d000001000220003900000060045000c90000000004240019000000000014004b000000e00000213d000000000005004b000013050000c13d000001a00200043d00000a200020009c000000e00000213d0000001f04200039000000000034004b000000000500001900000a220500804100000a2204400197000000000004004b000000000600001900000a220600404100000a220040009c000000000605c019000000000006004b000000e00000c13d000000e004200039001f00000004001d000000000704043300000a200070009c000001140000213d00000005067002100000003f0460003900000a2105400197000000400400043d0000000005540019000000000045004b0000000008000039000000010800403900000a200050009c000001140000213d0000000100800190000001140000c13d000000400050043f002c00000004001d00000000007404350000010005200039001e00000056001d0000001e0010006b000000e00000213d000000000007004b000013fe0000c13d000001c00200043d00000a200020009c000000e00000213d0000001f04200039000000000034004b000000000500001900000a220500404100000a2204400197000000000004004b000000000600001900000a220600204100000a220040009c000000000605c019000000000006004b000000e00000613d000000e004200039000000000504043300000a200050009c000001140000213d00000005045002100000003f0440003900000a2104400197000000400600043d0000000004460019001800000006001d000000000064004b0000000006000039000000010600403900000a200040009c000001140000213d0000000100600190000001140000c13d000000400040043f0000001804000029002b00000004001d0000000000540435000001000220003900000006045002100000000004240019000000000014004b000000e00000213d000000000005004b000016d60000c13d000001e00200043d00000a200020009c000000e00000213d0000001f04200039000000000034004b000000000300001900000a220300404100000a2204400197000000000004004b000000000500001900000a220500204100000a220040009c000000000503c019000000000005004b000000e00000613d000000e003200039000000000403043300000a200040009c000001140000213d00000005034002100000003f0330003900000a2103300197000000400500043d0000000003350019001700000005001d000000000053004b0000000005000039000000010500403900000a200030009c000001140000213d0000000100500190000001140000c13d000000400030043f00000017030000290000000003430436001600000003001d000001000220003900000240034000c90000000003230019000000000013004b000000e00000213d000000000004004b000017d40000c13d000000400100043d001d00000001001d0000000001000411000000000001004b00001a9c0000c13d0000001d03000029000000440130003900000a5f02000041000000000021043500000024013000390000001802000039000000000021043500000a60010000410000000000130435000000040130003900000020020000390000000000210435000017320000013d00000a840040009c0000037d0000a13d00000a850040009c000004970000213d00000a880040009c000006560000613d00000a890040009c000000e00000c13d0000000001000416000000000001004b000000e00000c13d0000000202000039000000000102041a000000800010043f000000000020043f000000000001004b000008740000613d000000a00400003900000ac90200004100000000030000190000000005040019000000000402041a000000000445043600000001022000390000000103300039000000000013004b0000025a0000413d00000abe0000013d00000a6f0040009c000003c00000a13d00000a700040009c000004c10000213d00000a730040009c0000066a0000613d00000a740040009c000000e00000c13d000000440030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000402100370000000000202043b00000a200020009c000000e00000213d0000002401100370000000000101043b002400000001001d00000a1f0010009c000000e00000213d0000014001000039000000400010043f000000800000043f000000a00000043f000000c00000043f000000e00000043f000001000000043f000001200000043f000000000020043f0000000a01000039000000200010043f00000040020000390000000001000019286328260000040f0000002402000029286323f50000040f002400000001001d000000400100043d002300000001001d28631ff20000040f0000002401000029000000000101041a00000a98001001980000000002000039000000010200c0390000002304000029000000a0034000390000000000230435000000700210027000000a1a0220019700000080034000390000000000230435000000500210027000000a1a022001970000006003400039000000000023043500000040021002700000ffff0220018f00000040034000390000000000230435000000200210027000000a1a022001970000002003400039000000000023043500000a1a0110019700000000001404350000000002040019000000400100043d002400000001001d2863209b0000040f00000ac50000013d00000a7a0040009c000005f00000213d00000a7d0040009c000007d80000613d00000a7e0040009c000000e00000c13d000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000401100370000000000101043b002400000001001d00000a200010009c000000e00000213d286324230000040f0000002401000029000000000010043f0000000501000039000000200010043f00000040020000390000000001000019286328260000040f002400000001001d000000400100043d002300000001001d28631fe70000040f0000002401000029000000000101041a00000023040000290000002002400039000000e003100270000000000032043500000a270110019700000000001404350000000001040019000005e70000013d00000a650040009c000006270000213d00000a680040009c000007e30000613d00000a690040009c000000e00000c13d000000440030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000402100370000000000202043b002400000002001d00000a200020009c000000e00000213d0000002401100370000000000101043b002300000001001d00000a200010009c000000e00000213d000000230130006a00000a230010009c000000e00000213d000000a40010008c000000e00000413d0000002401000029000000000010043f0000000901000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000400200043d00000a260020009c000001140000213d000000000101043b0000022003200039000000400030043f000000000301041a000000d80430027000000a1a044001970000014005200039002200000005001d0000000000450435000000c8043002700000ffff0440018f0000012005200039002100000005001d0000000000450435000000b8043002700000ffff0440018f0000010005200039001c00000005001d0000000000450435000000a8043002700000ffff0440018f000000e005200039001300000005001d0000000000450435000000880430027000000a1a04400197000000c005200039001400000005001d000000000045043500000078043002700000ffff0440018f000000a005200039001600000005001d0000000000450435000000580430027000000a1a044001970000008005200039001500000005001d0000000000450435000000380430027000000a1a044001970000006005200039001900000005001d0000000000450435000000180430027000000a1a044001970000004005200039002000000005001d000000000045043500000008043002700000ffff0440018f0000002005200039001f00000005001d0000000000450435000000ff033001900000000004000039000000010400c03900000000004204350000000101100039000000000101041a000001600520003900000a1a04100197001700000005001d0000000000450435000000380410021000000a38044001970000020005200039001e00000005001d000000000045043500000a36001001980000000004000039000000010400c039000001e005200039001800000005001d0000000000450435000000800410027000000a1a04400197000001c005200039001d00000005001d0000000000450435000000600410027000000a1a04400197000001a005200039001b00000005001d00000000004504350000018002200039000000200110027000000a2001100197001a00000002001d0000000000120435000000000003004b00000dbc0000613d00000023010000290000006401100039001200000001001d0000000201100367000000000101043b00000a1f0010009c000000e00000213d000000000010043f0000000c01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d0000000202000367000000000101043b000000000101041a000000000001004b00000f1d0000c13d0000001201200360000000000101043b00000a1f0010009c000000e00000213d000000400200043d00000aae0300004100000000003204350000000403200039000000000013043500000a1a0020009c00000a1a02008041000000400120021000000a4e011001c7000028650001043000000a8a0040009c00000a040000613d00000a8b0040009c000009d30000613d00000a8c0040009c000000e00000c13d0000000001000416000000000001004b000000e00000c13d000000e001000039000000400010043f000000800000043f000000a00000043f000000c00000043f28631fdc0000040f0000000001000412002a00000001001d002900000000003d0000800501000039000000440300003900000000040004150000002a0440008a000000050440021000000a9b020000412863283b0000040f00000a1e01100197002400000001001d000000e00010043f0000000001000412002800000001001d002700200000003d0000000004000415000000280440008a0000000504400210000080050100003900000a9b0200004100000044030000392863283b0000040f00000a1f01100197000001000010043f0000000001000412002600000001001d002500400000003d0000000004000415000000260440008a0000000504400210000080050100003900000a9b0200004100000044030000392863283b0000040f00000a1a01100197000001200010043f000000400100043d00000024020000290000000002210436000001000300043d00000a1f033001970000000000320435000001200200043d00000a1a022001970000004003100039000000000023043500000a1a0010009c00000a1a01008041000000400110021000000acc011001c7000028640001042e00000a750040009c00000a1c0000613d00000a760040009c000009e70000613d00000a770040009c000000e00000c13d000000440030008c000000e00000413d0000000004000416000000000004004b000000e00000c13d0000000404100370000000000404043b00000a200040009c000000e00000213d0000002305400039000000000035004b000000e00000813d0000000405400039000000000551034f000000000605043b00000a200060009c000001140000213d00000005056002100000003f0750003900000a210770019700000ab10070009c000001140000213d0000008007700039000000400070043f000000800060043f00000024044000390000000005450019000000000035004b000000e00000213d000000000006004b000003ee0000613d000000000641034f000000000606043b00000a1f0060009c000000e00000213d000000200220003900000000006204350000002004400039000000000054004b000003e50000413d0000002402100370000000000202043b00000a200020009c000000e00000213d0000002304200039000000000034004b000000000500001900000a220500804100000a2204400197000000000004004b000000000600001900000a220600404100000a220040009c000000000605c019000000000006004b000000e00000c13d0000000404200039000000000441034f000000000404043b00000a200040009c000001140000213d00000005054002100000003f0650003900000a2106600197000000400700043d0000000006670019002000000007001d000000000076004b0000000007000039000000010700403900000a200060009c000001140000213d0000000100700190000001140000c13d000000400060043f00000020060000290000000006460436001f00000006001d00000024022000390000000005250019000000000035004b000000e00000213d000000000004004b000004240000613d0000002003000029000000000421034f000000000404043b00000a1f0040009c000000e00000213d000000200330003900000000004304350000002002200039000000000052004b0000041b0000413d000000000100041a00000a1f011001970000000002000411000000000012004b00000af40000c13d000000800100043d000000000001004b000010700000c13d00000020010000290000000001010433000000000001004b00000a020000613d002400000000001d000004380000013d0000002402000029002400010020003d00000020010000290000000001010433000000240010006b00000a020000813d000000240100002900000005011002100000001f01100029002200000001001d000000000101043300000a1f01100197002300000001001d000000000010043f0000000c01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000301041a000000000003004b000004320000613d0000000b01000039000000000201041a000000000002004b000017680000613d000000010130008a000000000023004b000004710000613d000000000012004b00001ed50000a13d00000a330130009a00000a330220009a000000000202041a000000000021041b000000000020043f0000000c01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c70000801002000039002100000003001d2863285e0000040f00000021030000290000000100200190000000e00000613d000000000101043b000000000031041b0000000b01000039000000000301041a000000000003004b00001edb0000613d000000010130008a00000a330230009a000000000002041b0000000b02000039000000000012041b0000002301000029000000000010043f0000000c01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000001041b00000020010000290000000001010433000000240010006c00001ed50000a13d00000022010000290000000001010433000000000200041400000a1f0510019700000a1a0020009c00000a1a02008041000000c00120021000000a31011001c70000800d02000039000000020300003900000a3404000041286328590000040f0000000100200190000004320000c13d000000e00000013d00000a860040009c000007040000613d00000a870040009c000000e00000c13d000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000401100370000000000101043b001e00000001001d00000a200010009c000000e00000213d0000001e0130006a00000a230010009c000000e00000213d000000440010008c000000e00000413d0000000001000411000000000010043f0000000301000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000101041a000000000001004b00000b050000c13d000000400100043d00000ac802000041000000000021043500000004021000390000000003000411000006500000013d00000a710040009c0000070b0000613d00000a720040009c000000e00000c13d000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000402100370000000000202043b00000a200020009c000000e00000213d000000000423004900000a230040009c000000e00000213d000000440040008c000000e00000413d000000c005000039000000400050043f0000000404200039000000000641034f000000000606043b00000a200060009c000000e00000213d00000000062600190000002307600039000000000037004b000000e00000813d0000000407600039000000000771034f000000000807043b00000a200080009c000001140000213d00000005078002100000003f0970003900000a210990019700000a250090009c000001140000213d000000c009900039000000400090043f000000c00080043f00000024066000390000000007670019000000000037004b000000e00000213d000000000008004b000004fa0000613d000000000861034f000000000808043b00000a1f0080009c000000e00000213d000000200550003900000000008504350000002006600039000000000076004b000004f10000413d000000c005000039000000800050043f0000002004400039000000000441034f000000000404043b00000a200040009c000000e00000213d00000000022400190000002304200039000000000034004b000000000500001900000a220500804100000a2204400197000000000004004b000000000600001900000a220600404100000a220040009c000000000605c019000000000006004b000000e00000c13d0000000404200039000000000441034f000000000404043b00000a200040009c000001140000213d00000005054002100000003f0650003900000a2106600197000000400700043d0000000006670019002100000007001d000000000076004b0000000007000039000000010700403900000a200060009c000001140000213d0000000100700190000001140000c13d000000400060043f00000021060000290000000004460436002000000004001d00000024022000390000000004250019000000000034004b000000e00000213d000000000042004b000005340000813d0000002103000029000000000521034f000000000505043b00000a1f0050009c000000e00000213d000000200330003900000000005304350000002002200039000000000042004b0000052b0000413d0000002101000029000000a00010043f000000000100041a00000a1f011001970000000002000411000000000012004b00000af40000c13d00000021010000290000000001010433000000000001004b0000129e0000c13d000000800100043d002100000001001d0000000021010434002200000002001d000000000001004b00000a020000613d002400000000001d000000240100002900000005011002100000002201100029000000000101043300000a1f0110019800001b830000613d002300000001001d000000000010043f0000000301000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000101041a000000000001004b000005780000c13d0000000201000039000000000101041a00000a200010009c000001140000213d00000001021000390000000203000039000000000023041b00000a2e0110009a0000002302000029000000000021041b000000000103041a002000000001001d000000000020043f0000000301000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b0000002002000029000000000021041b000000400100043d0000002302000029000000000021043500000a1a0010009c00000a1a010080410000004001100210000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a2c011001c70000800d02000039000000010300003900000a2f04000041286328590000040f0000000100200190000000e00000613d0000002402000029002400010020003d00000021010000290000000001010433000000240010006b000005460000413d00000a020000013d00000a800040009c000007ea0000613d00000a810040009c000000e00000c13d000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000402100370000000000202043b00000a200020009c000000e00000213d0000002304200039000000000034004b000000e00000813d0000000404200039000000000141034f000000000101043b001f00000001001d00000a200010009c000000e00000213d001e00240020003d0000001f0100002900000005021002100000001e01200029000000000031004b000000e00000213d0000003f0120003900000a210310019700000ab10030009c000001140000213d0000008001300039000000400010043f0000001f04000029000000800040043f000000000004004b00000d190000c13d00000020020000390000000003210436000000800200043d00000000002304350000004004100039000000000002004b000005d00000613d000000800300003900000000050000190000000006010019000000000704001900000020033000390000000004030433000000008404043400000a270440019700000000004704350000006004600039000000000608043300000a1a06600197000000000064043500000040047000390000000105500039000000000025004b0000000006070019000005c10000413d000000000214004900000a1a0020009c00000a1a02008041000000600220021000000a1a0010009c00000a1a010080410000004001100210000000000112019f000028640001042e00000a6b0040009c0000086b0000613d00000a6c0040009c000000e00000c13d000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000401100370000000000101043b00000a1f0010009c000000e00000213d286324860000040f000000400200043d002400000002001d2863203d0000040f000000240100002900000a1a0010009c00000a1a01008041000000400110021000000a8d011001c7000028640001042e00000a7b0040009c000008760000613d00000a7c0040009c000000e00000c13d000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000401100370000000000101043b002400000001001d00000a200010009c000000e00000213d2863200f0000040f00000200021000390000000000020435000001e0021000390000000000020435000001c0021000390000000000020435000001a00210003900000000000204350000018002100039000000000002043500000160021000390000000000020435000001400210003900000000000204350000012002100039000000000002043500000100021000390000000000020435000000e0021000390000000000020435000000c0021000390000000000020435000000a0021000390000000000020435000000800210003900000000000204350000006002100039000000000002043500000040021000390000000000020435000000000101043600000000000104350000002401000029286324050000040f286324320000040f0000000002010019000000400100043d002400000001001d286320440000040f00000ac50000013d00000a660040009c000009b90000613d00000a670040009c000000e00000c13d000000440030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000402100370000000000202043b002400000002001d00000a1f0020009c000000e00000213d0000002401100370000000000101043b002300000001001d00000a200010009c000000e00000213d0000002301000029000000000010043f0000000901000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000101041a000000ff0010019000000bf40000c13d000000400100043d00000a8e02000041000000000021043500000004021000390000002303000029000000000032043500000a1a0010009c00000a1a01008041000000400110021000000a4e011001c700002865000104300000000001000416000000000001004b000000e00000c13d000000c001000039000000400010043f0000001301000039000000800010043f00000aca01000041000000a00010043f0000002001000039000000c00010043f0000008001000039000000e0020000392863201b0000040f000000c00110008a00000a1a0010009c00000a1a01008041000000600110021000000acb011001c7000028640001042e000000440030008c000000e00000413d0000000004000416000000000004004b000000e00000c13d0000000404100370000000000604043b00000a200060009c000000e00000213d0000002304600039000000000034004b000000e00000813d0000000405600039000000000451034f000000000404043b00000a200040009c000000e00000213d00000000064600190000002406600039000000000036004b000000e00000213d0000002406100370000000000606043b00000a200060009c000000e00000213d0000002307600039000000000037004b000000e00000813d002200040060003d0000002207100360000000000707043b002300000007001d00000a200070009c000000e00000213d0000002406600039002100000006001d002400230060002d000000240030006b000000e00000213d0000001f0340003900000ace033001970000003f0330003900000ace0330019700000ab10030009c000001140000213d0000008003300039000000400030043f0000002003500039000000000331034f000000800040043f00000ace054001980000001f0640018f000000a001500039000006a60000613d000000a007000039000000000803034f000000008908043c0000000007970436000000000017004b000006a20000c13d000000000006004b000006b30000613d000000000353034f0000000305600210000000000601043300000000065601cf000000000656022f000000000303043b0000010005500089000000000353022f00000000035301cf000000000363019f0000000000310435000000a0014000390000000000010435000000de0100043d002000000001001d00000ab303100197000000400100043d0000008004100039000000ca0500043d000000c00600043d0000000000340435001f00000006001d00000ab40360019700000060041000390000000000340435000000000221043600000060045002700000004003100039001e00000004001d0000000000430435000000000300041100000a1f03300197000000000032043500000ab50010009c000001140000213d000000a003100039000000400030043f00000a1a0020009c00000a1a020080410000004002200210000000000101043300000a1a0010009c00000a1a010080410000006001100210000000000121019f000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a31011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000010043f0000000401000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000101041a000000ff00100190000010bc0000c13d000000400200043d002400000002001d00000aba010000410000000000120435000000040120003900000000020004110000001e030000290000001f040000290000002005000029286327c30000040f0000002402000029000000000121004900000a1a0010009c00000a1a01008041000000600110021000000a1a0020009c00000a1a020080410000004002200210000000000121019f00002865000104300000000001000416000000000001004b000000e00000c13d0000002401000039000000800010043f00000aaf01000041000028640001042e0000000001000416000000000001004b000000e00000c13d000000000100041a00000a1f01100197000000800010043f00000aaf01000041000028640001042e000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000402100370000000000202043b00000a200020009c000000e00000213d0000002304200039000000000034004b000000e00000813d0000000404200039000000000441034f000000000504043b00000a200050009c000001140000213d00000005045002100000003f0440003900000a210440019700000ab10040009c000001140000213d0000008004400039000000400040043f000000800050043f000000240220003900000060045000c90000000004240019000000000034004b000000e00000213d000000000005004b00000c0e0000c13d000000000100041a00000a1f011001970000000002000411000000000012004b00000af40000c13d000000800100043d000000000001004b00000a020000613d002400000000001d00000024010000290000000501100210000000a001100039000000000101043300000020021000390000000002020433002200000002001d000000000101043300000a1f01100197002300000001001d000000000010043f0000000701000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000201041a00000a3502200197000000220500002900000020035000390000000004030433000000a00440021000000a3604400197000000000242019f000000000405043300000a1f04400197000000000242019f000000000021041b000000400100043d00000000024104360000000003030433000000ff0330018f000000000032043500000a1a0010009c00000a1a010080410000004001100210000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a2a011001c70000800d02000039000000020300003900000a37040000410000002305000029286328590000040f0000000100200190000000e00000613d0000002402000029002400010020003d000000800100043d000000240010006b0000073c0000413d00000a020000013d000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000402100370000000000202043b00000a200020009c000000e00000213d0000002304200039000000000034004b000000e00000813d0000000404200039000000000441034f000000000504043b00000a200050009c000001140000213d00000005045002100000003f0440003900000a210440019700000ab10040009c000001140000213d0000008004400039000000400040043f000000800050043f000000240220003900000006045002100000000004240019000000000034004b000000e00000213d000000000005004b00000c350000c13d000000000100041a00000a1f011001970000000002000411000000000012004b00000af40000c13d000000800100043d000000000001004b00000a020000613d002400000000001d00000024010000290000000501100210000000a001100039000000000101043300000020021000390000000002020433002200000002001d000000000101043300000a1f01100197002300000001001d000000000010043f0000000801000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000220200002900000a2002200197000000000101043b000000000301041a00000a4f03300197000000000323019f000000000031041b000000400100043d000000000021043500000a1a0010009c00000a1a010080410000004001100210000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a2c011001c70000800d02000039000000020300003900000a50040000410000002305000029286328590000040f0000000100200190000000e00000613d0000002402000029002400010020003d000000800100043d000000240010006b000007a30000413d00000a020000013d000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000401100370000000000101043b00000a1f0010009c000000e00000213d286326190000040f00000a150000013d0000000001000416000000000001004b000000e00000c13d0000001201000039000000800010043f00000aaf01000041000028640001042e000000c40030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000402100370000000000202043b002400000002001d00000a200020009c000000e00000213d0000002402100370000000000202043b002300000002001d00000a1f0020009c000000e00000213d0000006402100370000000000202043b00000a200020009c000000e00000213d0000002304200039000000000034004b000000e00000813d0000000404200039000000000441034f000000000404043b002200000004001d00000a200040009c000000e00000213d0000002404200039001f00000004001d001e00220040002d0000001e0030006b000000e00000213d0000008402100370000000000202043b002100000002001d00000a200020009c000000e00000213d00000021020000290000002302200039000000000032004b000000e00000813d00000021020000290000000402200039000000000221034f000000000202043b001d00000002001d00000a200020009c000000e00000213d000000210200002900000024042000390000001d020000290000000502200210002000000004001d001c00000002001d0000000002420019000000000032004b000000e00000213d000000a402100370000000000202043b00000a200020009c000000e00000213d0000002304200039000000000034004b000000e00000813d0000000404200039000000000441034f000000000404043b001b00000004001d00000a200040009c000000e00000213d001a00240020003d0000001b0200002900000006022002100000001a02200029000000000032004b000000e00000213d0000004401100370000000000101043b001600000001001d00000a9b0100004100000000001004430000000001000412000000040010044300000020010000390000002400100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a9c011001c700008005020000392863285e0000040f000000010020019000001a9b0000613d000000000301043b000000230130014f00000a1f00100198000008510000613d00000023010000290000001602000029286320c10000040f001600000001001d00000a9b010000410000000000100443000000000100041200000004001004430000002400000443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a9c011001c700008005020000392863285e0000040f000000010020019000001a9b0000613d000000000101043b00000a1e01100197000000160010006b000013280000a13d000000400200043d0000002403200039000000000013043500000ac201000041000000000012043500000004012000390000001603000029000012980000013d0000000001000416000000000001004b000000e00000c13d0000000b02000039000000000102041a000000800010043f000000000020043f000000000001004b00000ab40000c13d000000200200003900000abf0000013d000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000402100370000000000202043b00000a200020009c000000e00000213d0000002304200039000000000034004b000000e00000813d0000000404200039000000000441034f000000000504043b00000a200050009c000001140000213d00000005045002100000003f0440003900000a210440019700000ab10040009c000001140000213d0000008004400039000000400040043f000000800050043f000000240220003900000240045000c90000000004240019000000000034004b000000e00000213d000000000005004b00000c500000c13d000000000100041a00000a1f011001970000000002000411000000000012004b00000af40000c13d000000800100043d000000000001004b00000a020000613d002400000000001d00000024010000290000000501100210000000a0011000390000000001010433000000001201043400000a200420019800000f170000613d00000000030104330000016001300039002300000001001d000000000101043300000a1a0110019800000f170000613d0000020002300039002200000002001d000000000202043300000a380220019700000a390020009c00000f170000c13d0000006002300039002000000002001d000000000202043300000a1a02200197000000000021004b00000f170000213d001f00000003001d000000000040043f0000000901000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c70000801002000039002100000004001d2863285e0000040f0000000100200190000000e00000613d0000002106000029000000000101043b0000000101100039000000000101041a0000001f050000290000000042050434000000000002004b0000000003000039000000010300c039000000400200043d0000000003320436001e00000004001d00000000040404330000ffff0440018f00000000004304350000004003500039001d00000003001d000000000303043300000a1a03300197000000400420003900000000003404350000002003000029000000000303043300000a1a03300197000000600420003900000000003404350000008003500039001c00000003001d000000000303043300000a1a0330019700000080042000390000000000340435000000a003500039001b00000003001d00000000030304330000ffff0330018f000000a0042000390000000000340435000000c003500039001a00000003001d000000000303043300000a1a03300197000000c0042000390000000000340435000000e003500039001900000003001d00000000030304330000ffff0330018f000000e00420003900000000003404350000010003500039001800000003001d00000000030304330000ffff0330018f000001000420003900000000003404350000012003500039001600000003001d00000000030304330000ffff0330018f000001200420003900000000003404350000014003500039001400000003001d000000000303043300000a1a03300197000001400420003900000000003404350000002303000029000000000303043300000a1a03300197000001600420003900000000003404350000018003500039001700000003001d000000000303043300000a200330019700000180042000390000000000340435000001a003500039001500000003001d000000000303043300000a1a03300197000001a0042000390000000000340435000001c003500039001300000003001d000000000303043300000a1a03300197000001c0042000390000000000340435000001e003500039001200000003001d0000000003030433000000000003004b0000000003000039000000010300c039000001e00420003900000000003404350000002203000029000000000303043300000a38033001970000020004200039000000000034043500000a1a0020009c00000a1a02008041000000400220021000000a3a001001980000093d0000613d000000000100041400000a1a0010009c00000a1a01008041000000c001100210000000000121019f00000abe011001c70000800d02000039000000020300003900000a3b04000041000009460000013d000000000100041400000a1a0010009c00000a1a01008041000000c001100210000000000121019f00000abe011001c70000800d02000039000000020300003900000a3c040000410000000005060019286328590000040f0000000100200190000000e00000613d0000002101000029000000000010043f0000000901000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d0000001f020000290000000002020433000000000002004b000000000101043b000000000201041a00000a3d02200197000000010220c1bf0000001e030000290000000003030433000000080330021000000a3e03300197000000000232019f0000001d030000290000000003030433000000180330021000000a3f03300197000000000232019f00000020030000290000000003030433000000380330021000000a4003300197000000000232019f0000001c030000290000000003030433000000580330021000000a4103300197000000000232019f0000001b030000290000000003030433000000780330021000000a4203300197000000000232019f0000001a030000290000000003030433000000880330021000000a4303300197000000000232019f00000019030000290000000003030433000000a80330021000000a4403300197000000000232019f00000018030000290000000003030433000000b80330021000000a4503300197000000000232019f00000016030000290000000003030433000000c80330021000000a4603300197000000000232019f00000014030000290000000003030433000000d80330021000000a4703300197000000000232019f000000000021041b00000001011000390000002302000029000000000202043300000a1a02200197000000000301041a00000a4803300197000000000223019f00000017030000290000000003030433000000200330021000000a4903300197000000000232019f00000015030000290000000003030433000000600330021000000a4a03300197000000000232019f00000013030000290000000003030433000000800330021000000a4b03300197000000000232019f00000012030000290000000003030433000000000003004b00000a4c030000410000000003006019000000000232019f00000022030000290000000003030433000000380330027000000a3a03300197000000000232019f000000000021041b0000002402000029002400010020003d000000800100043d000000240010006b0000089f0000413d00000a020000013d000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000401100370000000000601043b00000a1f0060009c000000e00000213d000000000100041a00000a1f011001970000000005000411000000000015004b00000ad90000c13d000000000056004b00000ae30000c13d00000a6001000041000000800010043f0000002001000039000000840010043f0000001701000039000000a40010043f00000a9201000041000000c40010043f00000a90010000410000286500010430000000240030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000401100370000000000101043b00000a1f0010009c000000e00000213d000000000010043f0000000801000039000000200010043f00000040020000390000000001000019286328260000040f000000000101041a00000a2001100197000000800010043f00000aaf01000041000028640001042e0000000001000416000000000001004b000000e00000c13d0000000101000039000000000201041a00000a1f032001970000000006000411000000000036004b00000acf0000c13d000000000300041a00000a2804300197000000000464019f000000000040041b00000a2802200197000000000021041b000000000100041400000a1f0530019700000a1a0010009c00000a1a01008041000000c00110021000000a31011001c70000800d02000039000000030300003900000abd04000041286328590000040f0000000100200190000000e00000613d0000000001000019000028640001042e000000640030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000402100370000000000402043b00000a1f0040009c000000e00000213d0000004402100370000000000302043b00000a1f0030009c000000e00000213d0000002401100370000000000201043b0000000001040019286320c10000040f000000400200043d000000000012043500000a1a0020009c00000a1a02008041000000400120021000000aa9011001c7000028640001042e000000440030008c000000e00000413d0000000002000416000000000002004b000000e00000c13d0000000402100370000000000202043b002300000002001d00000a200020009c000000e00000213d00000023020000290000002302200039000000000032004b000000e00000813d00000023020000290000000402200039000000000221034f000000000202043b00000a200020009c000001140000213d00000005052002100000003f0450003900000a210440019700000ab10040009c000001140000213d0000008004400039000000400040043f000000800020043f00000023040000290000002404400039002200000045001d000000220030006b000000e00000213d000000000002004b00000d2a0000c13d0000002402100370000000000202043b00000a200020009c000000e00000213d0000002304200039000000000034004b000000000500001900000a220500404100000a2204400197000000000004004b000000000600001900000a220600204100000a220040009c000000000605c019000000000006004b000000e00000613d0000000404200039000000000441034f000000000504043b00000a200050009c000001140000213d00000005045002100000003f0440003900000a2104400197000000400600043d0000000004460019001e00000006001d000000000064004b0000000006000039000000010600403900000a200040009c000001140000213d0000000100600190000001140000c13d000000400040043f0000001e040000290000000004540436001d00000004001d000000240220003900000006045002100000000004240019000000000034004b000000e00000213d000000000005004b000010540000c13d000000000100041a00000a1f011001970000000002000411000000000012004b00000af40000c13d000000800100043d000000000001004b000011c90000c13d0000001e010000290000000001010433000000000001004b00000a020000613d002400000000001d000000240100002900000005011002100000001d01100029000000000101043300000020021000390000000002020433002300000002001d000000000101043300000a2001100197002200000001001d000000000010043f0000000a01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000230200002900000a1f02200197000000000101043b002300000002001d000000000020043f000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000001041b000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a31011001c70000800d02000039000000030300003900000a5b0400004100000022050000290000002306000029286328590000040f0000000100200190000000e00000613d0000002402000029002400010020003d0000001e010000290000000001010433000000240010006b00000a790000413d00000a020000013d000000a00400003900000ab00200004100000000030000190000000005040019000000000402041a000000000445043600000001022000390000000103300039000000000013004b00000ab70000413d000000600250008a000000800100003928631ffd0000040f000000400100043d002400000001001d00000080020000392863202d0000040f0000002402000029000000000121004900000a1a0010009c00000a1a01008041000000600110021000000a1a0020009c00000a1a020080410000004002200210000000000121019f000028640001042e00000a6001000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f00000abc01000041000000c40010043f00000a9001000041000028650001043000000a6001000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f00000a8f01000041000000c40010043f00000a900100004100002865000104300000000101000039000000000201041a00000a2802200197000000000262019f000000000021041b000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a31011001c70000800d02000039000000030300003900000a9104000041286328590000040f0000000100200190000000e00000613d00000a020000013d000000400100043d000000440210003900000a8f03000041000000000032043500000024021000390000001603000039000000000032043500000a6002000041000000000021043500000004021000390000002003000039000000000032043500000a1a0010009c00000a1a01008041000000400110021000000a61011001c700002865000104300000001e04000029002400040040003d00000002020003670000002401200360000000000301043b00000000010000310000000004410049000000230440008a00000a220540019700000a2206300197000000000756013f000000000056004b000000000500001900000a2205004041000000000043004b000000000400001900000a220400804100000a220070009c000000000504c019000000000005004b000000e00000c13d0000002403300029000000000232034f000000000202043b001b00000002001d00000a200020009c000000e00000213d0000001b02000029000000060220021000000000012100490000002002300039000000000012004b000000000300001900000a220300204100000a220110019700000a2202200197000000000412013f000000000012004b000000000100001900000a220100404100000a220040009c000000000103c019000000000001004b000000e00000c13d00000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f000000010020019000001a9b0000613d000000000101043b001c00000001001d001d0a1a0010019b0000001b0000006b00000fcf0000c13d0000002401000029001b00200010003d00000002020003670000001b01200360000000000301043b00000000010000310000001e0410006a000000230440008a00000a220540019700000a2206300197000000000756013f000000000056004b000000000500001900000a2205004041000000000043004b000000000400001900000a220400804100000a220070009c000000000504c019000000000005004b000000e00000c13d0000002403300029000000000232034f000000000202043b001a00000002001d00000a200020009c000000e00000213d0000001a02000029000000060220021000000000012100490000002002300039000000000012004b000000000300001900000a220300204100000a220110019700000a2202200197000000000412013f000000000012004b000000000100001900000a220100404100000a220040009c000000000103c019000000000001004b000000e00000c13d0000001a0000006b00000a020000613d000000000900001900000002010003670000001b02100360000000000302043b00000000020000310000001e0420006a000000230440008a00000a220540019700000a2206300197000000000756013f000000000056004b000000000500001900000a2205004041000000000043004b000000000400001900000a220400804100000a220070009c000000000504c019000000000005004b000000e00000c13d0000002404300029000000000341034f000000000303043b00000a200030009c000000e00000213d00000006053002100000000005520049000000200440003900000a220650019700000a2207400197000000000867013f000000000067004b000000000600001900000a2206004041000000000054004b000000000500001900000a220500204100000a220080009c000000000605c019000000000006004b000000e00000c13d000000000039004b00001ed50000813d00000006039002100000000003340019000000000232004900000a230020009c000000e00000213d000000400020008c000000e00000413d000000400400043d00000a240040009c000001140000213d0000004002400039000000400020043f000000000231034f000000000202043b00000a200020009c000000e00000213d00000000052404360000002002300039000000000121034f000000000101043b00000a270010009c000000e00000213d002300000005001d0000000000150435000000400300043d00000a240030009c000001140000213d0000004002300039000000400020043f002000000003001d00000000021304360000001d01000029001f00000002001d0000000000120435000000000104043300000a2001100197000000000010043f0000000501000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c70000801002000039002100000009001d002200000004001d2863285e0000040f000000220400002900000001002001900000002305000029000000e00000613d0000002002000029000000000202043300000a27022001970000001f030000290000000003030433000000e003300210000000000223019f000000000101043b000000000021041b00000000020404330000000001050433000000400300043d00000020043000390000001c05000029000000000054043500000a2701100197000000000013043500000a1a0030009c00000a1a030080410000004001300210000000000300041400000a1a0030009c00000a1a03008041000000c003300210000000000113019f00000a2a011001c700000a20052001970000800d02000039000000020300003900000ac704000041286328590000040f00000021090000290000000100200190000000e00000613d00000001099000390000001a0090006c00000b700000413d00000a020000013d0000002401000029286326190000040f0000002302000029000000000020043f0000000902000039000000200020043f002400000001001d00000040020000390000000001000019286328260000040f0000000101100039000000000101041a000000800110027000000a1a021001970000002301000029286327d00000040f000000400200043d000000200320003900000000001304350000002401000029000000000012043500000a1a0020009c00000a1a02008041000000400120021000000a8d011001c7000028640001042e000000a005000039000000000623004900000a230060009c000000e00000213d000000600060008c000000e00000413d000000400600043d00000a240060009c000001140000213d0000004007600039000000400070043f000000000721034f000000000707043b00000a1f0070009c000000e00000213d0000000007760436000000400800043d00000a240080009c000001140000213d0000004009800039000000400090043f0000002009200039000000000a91034f000000000a0a043b00000a1f00a0009c000000e00000213d000000000aa804360000002009900039000000000991034f000000000909043b000000ff0090008c000000e00000213d00000000009a0435000000000087043500000000056504360000006002200039000000000042004b00000c0f0000413d000007330000013d000000a005000039000000000623004900000a230060009c000000e00000213d000000400060008c000000e00000413d000000400600043d00000a240060009c000001140000213d0000004007600039000000400070043f000000000721034f000000000707043b00000a1f0070009c000000e00000213d00000000077604360000002008200039000000000881034f000000000808043b00000a200080009c000000e00000213d000000000087043500000000056504360000004002200039000000000042004b00000c360000413d0000079a0000013d000000a005000039000000000623004900000a230060009c000000e00000213d000002400060008c000000e00000413d000000400600043d00000a240060009c000001140000213d0000004007600039000000400070043f000000000721034f000000000707043b00000a200070009c000000e00000213d0000000007760436000000400800043d00000a260080009c000001140000213d0000022009800039000000400090043f0000002009200039000000000a91034f000000000a0a043b00000000000a004b000000000b000039000000010b00c0390000000000ba004b000000e00000c13d000000000aa804360000002009900039000000000b91034f000000000b0b043b0000ffff00b0008c000000e00000213d0000000000ba04350000002009900039000000000a91034f000000000a0a043b00000a1a00a0009c000000e00000213d000000400b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b00000a1a00a0009c000000e00000213d000000600b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b00000a1a00a0009c000000e00000213d000000800b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000ffff00a0008c000000e00000213d000000a00b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b00000a1a00a0009c000000e00000213d000000c00b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000ffff00a0008c000000e00000213d000000e00b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000ffff00a0008c000000e00000213d000001000b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b0000ffff00a0008c000000e00000213d000001200b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b00000a1a00a0009c000000e00000213d000001400b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b00000a1a00a0009c000000e00000213d000001600b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b00000a2000a0009c000000e00000213d000001800b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b00000a1a00a0009c000000e00000213d000001a00b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b00000a1a00a0009c000000e00000213d000001c00b8000390000000000ab04350000002009900039000000000a91034f000000000a0a043b00000000000a004b000000000b000039000000010b00c0390000000000ba004b000000e00000c13d000001e00b8000390000000000ab04350000002009900039000000000991034f000000000909043b00000a2700900198000000e00000c13d000002000a80003900000000009a0435000000000087043500000000056504360000024002200039000000000042004b00000c510000413d000008960000013d000000a005000039000000000623004900000a230060009c000000e00000213d000000a00060008c000000e00000413d000000400600043d00000ab50060009c000001140000213d000000a007600039000000400070043f000000000721034f000000000707043b00000a1f0070009c000000e00000213d00000000087604360000002007200039000000000971034f000000000909043b00000ac300900198000000e00000c13d00000000009804350000002007700039000000000871034f000000000808043b00000ac400800198000000e00000c13d000000400960003900000000008904350000002007700039000000000871034f000000000808043b00000a1f0080009c000000e00000213d000000600960003900000000008904350000002007700039000000000771034f000000000707043b000000000007004b0000000008000039000000010800c039000000000087004b000000e00000c13d000000800860003900000000007804350000000005650436000000a002200039000000000042004b00000ce70000413d0000003a0000013d00000a250030009c000001140000213d00000000030000190000004004100039000000400040043f000000200410003900000000000404350000000000010435000000a00430003900000000001404350000002003300039000000000023004b00000dc20000813d000000400100043d00000a240010009c00000d1c0000a13d000001140000013d000000a006000039002100240030009200000d330000013d00000024020000290000000000a2043500000000068604360000002004400039000000220040006c00000a3f0000813d000000000241034f000000000202043b00000a200020009c000000e00000213d0000002302200029000000210520006900000a230050009c000000e00000213d000000400050008c000000e00000413d000000400800043d00000a240080009c000001140000213d0000004005800039000000400050043f0000002405200039000000000751034f000000000707043b00000a200070009c000000e00000213d0000000007780436002400000007001d0000002005500039000000000551034f000000000505043b00000a200050009c000000e00000213d00000000022500190000004305200039000000000035004b000000000700001900000a220700804100000a2205500197000000000005004b000000000900001900000a220900404100000a220050009c000000000907c019000000000009004b000000e00000c13d0000002405200039000000000551034f000000000c05043b00000a2000c0009c000001140000213d0000000505c002100000003f0550003900000a2105500197000000400a00043d000000000b5a00190000000000ab004b0000000005000039000000010500403900000a2000b0009c000001140000213d0000000100500190000001140000c13d0000004000b0043f0000000000ca0435000000440b200039000000e002c000c9000000000cb2001900000000003c004b000000e00000213d0000000000cb004b00000d2d0000813d000000000d0a00190000000002b3004900000a230020009c000000e00000213d000000e00020008c000000e00000413d000000400e00043d00000a2400e0009c000001140000213d0000004002e00039000000400020043f0000000002b1034f000000000202043b00000a1f0020009c000000e00000213d000000000f2e0436000000400200043d00000a250020009c000001140000213d000000c005200039000000400050043f0000002005b00039000000000751034f000000000707043b00000a1a0070009c000000e00000213d00000000077204360000002005500039000000000951034f000000000909043b00000a1a0090009c000000e00000213d00000000009704350000002005500039000000000751034f000000000707043b0000ffff0070008c000000e00000213d000000400920003900000000007904350000002005500039000000000751034f000000000707043b00000a1a0070009c000000e00000213d000000600920003900000000007904350000002005500039000000000751034f000000000707043b00000a1a0070009c000000e00000213d000000800920003900000000007904350000002005500039000000000551034f000000000505043b000000000005004b0000000007000039000000010700c039000000000075004b000000e00000c13d000000200dd00039000000a007200039000000000057043500000000002f04350000000000ed0435000000e00bb000390000000000cb004b00000d760000413d00000d2d0000013d000000400100043d00000a8e02000041000000000021043500000004021000390000002403000029000006500000013d0000000002000019002100000002001d0000000502200210002000000002001d0000001e012000290000000201100367000000000101043b002300000001001d00000a1f0010009c000000e00000213d000000400100043d00000a240010009c000001140000213d0000004002100039000000400020043f0000002002100039000000000002043500000000000104350000002301000029000000000010043f0000000601000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000400300043d00000a240030009c000001140000213d000000000101043b0000004002300039000000400020043f000000000101041a00000a2702100197002400000003001d0000000002230436000000e001100270002200000001001d000000000012043500000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f000000010020019000001a9b0000613d000000000101043b0022002200100074000017680000413d00000a9b0100004100000000001004430000000001000412000000040010044300000040010000390000002400100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a9c011001c700008005020000392863285e0000040f000000010020019000001a9b0000613d000000000101043b00000a1a01100197000000220010006b00000f030000413d0000002301000029000000000010043f0000000701000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000400200043d00000a240020009c0000002404000029000001140000213d000000000101043b0000004003200039000000400030043f000000000101041a00000a1f031001980000000003320436000000a001100270000000ff0110018f000000000013043500000f040000613d002300000003001d000000400100043d00000a240010009c000001140000213d0000004003100039000000400030043f0000002003100039000000000003043500000000000104350000000002020433000000400b00043d00000a9d0100004100000000051b0436000000000100041400000a1f02200197000000040020008c00000e400000c13d0000000103000031000000a00030008c000000a004000039000000000403401900000e6e0000013d001d00000005001d00000a1a00b0009c00000a1a0300004100000000030b4019000000400330021000000a1a0010009c00000a1a01008041000000c001100210000000000131019f00000a5e011001c7002200000002001d00240000000b001d2863285e0000040f000000240b000029000000600310027000000a1a03300197000000a00030008c000000a0040000390000000004034019000000e00640019000000000056b001900000e5c0000613d000000000701034f00000000080b0019000000007907043c0000000008980436000000000058004b00000e580000c13d0000001f0740019000000e690000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f0000000100200190000012670000613d00000022020000290000001d050000290000001f01400039000001e00110018f000000000ab1001900000000001a004b0000000001000039000000010100403900000a2000a0009c000001140000213d0000000100100190000001140000c13d0000004000a0043f000000a00030008c000000e00000413d00000000010b043300000a9e0010009c000000e00000213d0000008001b00039000000000101043300000a9e0010009c000000e00000213d000000000505043300000a220050009c000012580000813d00000a9f0100004100000000001a04350000000001000414000000040020008c000000200400003900000eb70000613d002200000005001d00000a1a00a0009c00000a1a0300004100000000030a4019000000400330021000000a1a0010009c00000a1a01008041000000c001100210000000000131019f00000a5e011001c700240000000a001d2863285e0000040f000000240a000029000000600310027000000a1a03300197000000200030008c00000020040000390000000004034019000000200640019000000000056a001900000ea60000613d000000000701034f00000000080a0019000000007907043c0000000008980436000000000058004b00000ea20000c13d0000001f0740019000000eb30000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f0000000100200190000012730000613d00000022050000290000001f01400039000000600110018f0000000001a1001900000a200010009c000001140000213d002400000001001d000000400010043f000000200030008c000000e00000413d00000000010a0433000000ff0010008c000000e00000213d00000023020000290000000002020433000000ff0220018f0000000002120019000000ff0020008c000017680000213d000000240120008c00000edd0000213d000000010100003900000ed60000613d00000024022000890000000a03000039000000010020019000000000043300a9000000010300603900000000011300a90000000102200272000000000304001900000ecf0000c13d000000000005004b00000f120000613d00000000025100a900000000035200d9000000000013004b00000eeb0000613d000017680000013d0000004d0010008c000017680000213d00000001020000390000000a03000039000000010010019000000000043300a9000000010300603900000000022300a90000000101100272000000000304001900000ee10000c13d000000000002004b00001f7c0000613d00000000022500d900000a270020009c0000125f0000213d000000240300002900000a240030009c000001140000213d0000004001300039000000400010043f0000000001230436002300000001001d00000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f000000010020019000001a9b0000613d000000000101043b00000a1a01100197000000230200002900000000001204350000002404000029000000800100043d0000002102000029000000000021004b00001ed50000a13d0000002001000029000000a0011000390000000000410435000000800100043d000000000021004b00001ed50000a13d00000001022000390000001f0020006c00000dc30000413d000010ba0000013d0000000002000019000000240300002900000a240030009c00000ef00000a13d000001140000013d000000400100043d00000a4d02000041000000000021043500000004021000390000000000420435000006510000013d00000023010000290000000405100039000000120100002900100020001000920000001001200360000000000401043b0000000003000031001100000005001d00000000015300490000001f0110008a00000a220510019700000a2206400197000000000756013f000000000056004b000000000600001900000a2206004041000000000014004b000000000800001900000a220800804100000a220070009c000000000608c019000000000006004b000000e00000c13d0000001104400029000000000642034f000000000606043b000f00000006001d00000a200060009c000000e00000213d0000000f06000029000000060660021000000000066300490000002004400039000000000064004b000000000700001900000a220700204100000a220660019700000a2204400197000000000864013f000000000064004b000000000400001900000a220400404100000a220080009c000000000407c019000000000004004b000000e00000c13d0000001004000029000000200640008a000000000462034f000000000404043b00000a2207400197000000000857013f000000000057004b000000000700001900000a2207004041000000000014004b000000000900001900000a220900804100000a220080009c000000000709c019000000000007004b000000e00000c13d0000001107400029000000000472034f000000000404043b00000a200040009c000000e00000213d00000000084300490000002007700039000000000087004b000000000900001900000a220900204100000a220880019700000a2207700197000000000a87013f000000000087004b000000000700001900000a220700404100000a2200a0009c000000000709c019000000000007004b000000e00000c13d000000200660008a000000000662034f000000000606043b00000a2207600197000000000857013f000000000057004b000000000500001900000a2205004041000000000016004b000000000100001900000a220100804100000a220080009c000000000501c019000000000005004b000000e00000c13d0000001101600029000000000512034f000000000605043b00000a200060009c000000e00000213d0000000005630049000000200710003900000a220150019700000a2208700197000000000918013f000000000018004b000000000100001900000a2201004041000000000057004b000000000500001900000a220500204100000a220090009c000000000105c019000000000001004b000000e00000c13d0000001f0160003900000ace011001970000003f0110003900000ace05100197000000400100043d0000000005510019000000000015004b0000000008000039000000010800403900000a200050009c000001140000213d0000000100800190000001140000c13d000000400050043f00000000056104360000000008760019000000000038004b000000e00000213d000000000772034f00000ace086001980000001f0960018f000000000385001900000faf0000613d000000000a07034f000000000b05001900000000ac0a043c000000000bcb043600000000003b004b00000fab0000c13d000000000009004b00000fbc0000613d000000000787034f0000000308900210000000000903043300000000098901cf000000000989022f000000000707043b0000010008800089000000000787022f00000000078701cf000000000797019f0000000000730435000000000365001900000000000304350000002003000029000000000303043300000a1a03300197000000000043004b000015d10000813d000000400100043d0000002402100039000000000042043500000aad0200004100000000002104350000000402100039000000000032043500000a1a0010009c00000a1a01008041000000400110021000000a5a011001c70000286500010430000000000900001900000002010003670000002402100360000000000302043b00000000020000310000001e0420006a000000230440008a00000a220540019700000a2206300197000000000756013f000000000056004b000000000500001900000a2205004041000000000043004b000000000400001900000a220400804100000a220070009c000000000504c019000000000005004b000000e00000c13d0000002404300029000000000341034f000000000303043b00000a200030009c000000e00000213d00000006053002100000000005520049000000200440003900000a220650019700000a2207400197000000000867013f000000000067004b000000000600001900000a2206004041000000000054004b000000000500001900000a220500204100000a220080009c000000000605c019000000000006004b000000e00000c13d000000000039004b00001ed50000813d00000006039002100000000003340019000000000232004900000a230020009c000000e00000213d000000400020008c000000e00000413d000000400400043d00000a240040009c000001140000213d0000004002400039000000400020043f000000000231034f000000000202043b00000a1f0020009c000000e00000213d00000000052404360000002002300039000000000121034f000000000101043b00000a270010009c000000e00000213d002300000005001d0000000000150435000000400300043d00000a240030009c000001140000213d0000004002300039000000400020043f002000000003001d00000000021304360000001d01000029001f00000002001d0000000000120435000000000104043300000a1f01100197000000000010043f0000000601000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c70000801002000039002100000009001d002200000004001d2863285e0000040f000000220400002900000001002001900000002305000029000000e00000613d0000002002000029000000000202043300000a27022001970000001f030000290000000003030433000000e003300210000000000223019f000000000101043b000000000021041b00000000020404330000000001050433000000400300043d00000020043000390000001c05000029000000000054043500000a2701100197000000000013043500000a1a0030009c00000a1a030080410000004001300210000000000300041400000a1a0030009c00000a1a03008041000000c003300210000000000113019f00000a2a011001c700000a1f052001970000800d02000039000000020300003900000ab804000041286328590000040f00000021090000290000000100200190000000e00000613d00000001099000390000001b0090006c00000fd00000413d00000b410000013d0000001e05000029000000000623004900000a230060009c000000e00000213d000000400060008c000000e00000413d000000400600043d00000a240060009c000001140000213d0000004007600039000000400070043f000000000721034f000000000707043b00000a200070009c000000e00000213d00000000077604360000002008200039000000000881034f000000000808043b00000a1f0080009c000000e00000213d0000002005500039000000000087043500000000006504350000004002200039000000000042004b000010550000413d00000a6c0000013d0000000002000019000010770000013d00000024020000290000000102200039000000800100043d000000000012004b0000042c0000813d002400000002001d0000000501200210000000a001100039002200000001001d000000000101043300000a1f01100197002300000001001d000000000010043f0000000c01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000101041a000000000001004b000010720000c13d0000000b03000039000000000103041a00000a200010009c000001140000213d0000000102100039000000000023041b00000a300110009a0000002302000029000000000021041b000000000103041a002100000001001d000000000020043f0000000c01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b0000002102000029000000000021041b000000800100043d000000240010006c00001ed50000a13d00000022010000290000000001010433000000000200041400000a1f0510019700000a1a0020009c00000a1a02008041000000c00120021000000a31011001c70000800d02000039000000020300003900000a3204000041286328590000040f0000000100200190000010720000c13d000000e00000013d000000400100043d000005b70000013d0000002301000029000000200010008c000000e00000413d000000220100002900000020021000390000000201000367000000000221034f000000000202043b00000a200020009c000000e00000213d00000021022000290000001f03200039000000240030006c000000e00000813d000000000321034f000000000403043b00000a200040009c000001140000213d00000005034002100000003f0330003900000a2103300197000000400500043d0000000003350019002300000005001d000000000053004b0000000005000039000000010500403900000a200030009c000001140000213d0000000100500190000001140000c13d000000400030043f00000023030000290000000003430436001e00000003001d000000200220003900000060034000c90000000003230019000000240030006c000000e00000213d000000000004004b00000a020000613d0000002304000029000000240520006900000a230050009c000000e00000213d000000600050008c000000e00000413d000000400500043d00000a1d0050009c000001140000213d0000006006500039000000400060043f000000000621034f000000000606043b00000a1f0060009c000000e00000213d00000000076504360000002006200039000000000861034f000000000808043b00000a270080009c000000e00000213d00000000008704350000002006600039000000000661034f000000000606043b00000a1a0060009c000000e00000213d00000020044000390000004007500039000000000067043500000000005404350000006002200039000000000032004b000010e70000413d00000023010000290000000001010433000000000001004b00000a020000613d002400000000001d000000240100002900000005011002100000001e01100029002200000001001d0000000001010433000000000101043300000a1f01100197000000000010043f0000000701000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000101041a00000a3600100198000017040000613d00000023020000290000000002020433000000240020006c00001ed50000a13d00000ab60210019700000ab70020009c000017680000213d000000a003100270000000220100002900000000010104330000002002100039000000000202043300000a27022001970000001203300039000000ff0430018f000000240340008c000011470000213d0000000103000039000011400000613d00000024044000890000000a05000039000000010040019000000000065500a9000000010500603900000000033500a900000001044002720000000005060019000011390000c13d000000000002004b000011c70000613d00000000042300a900000000022400d9000000000032004b000011550000613d000017680000013d0000004d0030008c000017680000213d00000001040000390000000a05000039000000010030019000000000065500a9000000010500603900000000044500a9000000010330027200000000050600190000114b0000c13d000000000004004b00001f7c0000613d00000000044200d9002100000004001d00000a270040009c000017370000213d00000040021000390000000002020433002000000002001d000000000101043300000a1f01100197000000000010043f0000000601000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000200200002900000a1a02200197000000000101043b000000000101041a000000e001100270000000000012004b0000170e0000413d00000023010000290000000001010433000000240010006c00001ed50000a13d000000400100043d002000000001001d00000a240010009c000001140000213d000000220100002900000000010104330000004001100039000000000101043300000020030000290000004002300039000000400020043f0000002102000029000000000223043600000a1a01100197001f00000002001d000000000012043500000023010000290000000001010433000000240010006c00001ed50000a13d00000022010000290000000001010433000000000101043300000a1f01100197000000000010043f0000000601000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d0000002002000029000000000202043300000a27022001970000001f030000290000000003030433000000e003300210000000000223019f000000000101043b000000000021041b00000023010000290000000001010433000000240010006c00001ed50000a13d0000002201000029000000000101043300000000020104330000004001100039000000000101043300000a1a01100197000000400300043d000000200430003900000000001404350000002101000029000000000013043500000a1a0030009c00000a1a030080410000004001300210000000000300041400000a1a0030009c00000a1a03008041000000c003300210000000000113019f00000a1f0520019700000a2a011001c70000800d02000039000000020300003900000ab804000041286328590000040f0000000100200190000000e00000613d0000002402000029002400010020003d00000023010000290000000001010433000000240010006b0000110d0000413d00000a020000013d002100000000001d000011580000013d0000000005000019000011ce0000013d0000000105500039000000000015004b00000a740000813d0000000502500210000000a00220003900000000030204330000002002300039001f00000002001d00000000020204330000000004020433000000000004004b000011cb0000613d001c00000005001d000000000103043300230a200010019b0000000003000019002000000003001d0000000501300210000000000112001900000020011000390000000001010433000000002101043400240a1f0010019b0000000001020433002100000001001d0000008001100039002200000001001d000000000101043300000a1a011001970000001f0010008c000012910000a13d0000002301000029000000000010043f0000000a01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f00000001002001900000002403000029000000e00000613d000000000101043b000000000030043f000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f00000024060000290000000100200190000000e00000613d000000000101043b000000000201041a00000a5102200197000000210a0000290000002003a000390000000004030433000000200440021000000a5204400197000000000242019f0000004004a000390000000005040433000000400550021000000a5305500197000000000252019f0000006005a000390000000007050433000000500770021000000a5407700197000000000272019f00000022090000290000000007090433000000700770021000000a5507700197000000000272019f000000a008a000390000000007080433000000000007004b00000a56070000410000000007006019000000000272019f00000000070a043300000a1a07700197000000000272019f000000000021041b000000400100043d0000000002710436000000000303043300000a1a03300197000000000032043500000000020404330000ffff0220018f00000040031000390000000000230435000000000205043300000a1a0220019700000060031000390000000000230435000000000209043300000a1a02200197000000800310003900000000002304350000000002080433000000000002004b0000000002000039000000010200c039000000a003100039000000000023043500000a1a0010009c00000a1a010080410000004001100210000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a57011001c70000800d02000039000000030300003900000a58040000410000002305000029286328590000040f0000000100200190000000e00000613d000000200300002900000001033000390000001f0100002900000000020104330000000001020433000000000013004b000011db0000413d000000800100043d0000001c05000029000011cb0000013d00000aa00100004100000000001a043500000a1a00a0009c00000a1a0a0080410000004001a0021000000a5e011001c7000028650001043000000aa0010000410000002402000029000000000012043500000a1a0020009c00000a1a02008041000000400120021000000a5e011001c700002865000104300000001f0530018f00000a1c06300198000000400200043d00000000046200190000127e0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b0000126e0000c13d0000127e0000013d0000001f0530018f00000a1c06300198000000400200043d00000000046200190000127e0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b0000127a0000c13d000000000005004b0000128b0000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f0000000000140435000000600130021000000a1a0020009c00000a1a020080410000004002200210000000000112019f0000286500010430000000400200043d0000002403200039000000000013043500000a5901000041000000000012043500000004012000390000002403000029000000000031043500000a1a0020009c00000a1a02008041000000400120021000000a5a011001c700002865000104300000000002000019000012a60000013d0000002302000029000000010220003900000021010000290000000001010433000000000012004b0000053f0000813d002300000002001d00000005012002100000002001100029000000000101043300000a1f01100197002400000001001d000000000010043f0000000301000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000101041a002200000001001d000000000001004b000012a00000613d0000000201000039000000000201041a000000000002004b000017680000613d0000002203000029000000010130008a000000000032004b000012df0000613d000000000012004b00001ed50000a13d00000a2b0130009a00000a2b0220009a000000000202041a000000000021041b000000000020043f0000000301000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b0000002202000029000000000021041b0000000201000039000000000301041a000000000003004b00001edb0000613d000000010130008a00000a2b0230009a000000000002041b0000000202000039000000000012041b0000002401000029000000000010043f0000000301000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000001041b000000400100043d0000002402000029000000000021043500000a1a0010009c00000a1a010080410000004001100210000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a2c011001c70000800d02000039000000010300003900000a2d04000041286328590000040f0000000100200190000012a00000c13d000000e00000013d0000001a05000029000000000621004900000a230060009c000000e00000213d000000600060008c000000e00000413d000000400600043d00000a240060009c000001140000213d0000004007600039000000400070043f000000009702043400000a1f0070009c000000e00000213d0000000007760436000000400800043d00000a240080009c000001140000213d000000400a8000390000004000a0043f000000000909043300000a1f0090009c000000e00000213d0000000009980436000000400a200039000000000a0a0433000000ff00a0008c000000e00000213d0000000000a90435000000000087043500000000056504360000006002200039000000000042004b000013060000413d000001b50000013d0000002401000029000000000010043f0000000901000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000400200043d00000a240020009c000001140000213d0000000101100039000000000101041a0000004003200039000000400030043f000000200320003900000000000304350000000000020435000000220000006b000013b30000c13d000000400200043d001500000002001d00000a240020009c000001140000213d00000a1a0110019700000015030000290000004002300039000000400020043f0000000000130435001300000000001d000000000100001900000015020000290000002002200039001400000002001d00000000001204350000002401000029000000000010043f0000000901000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b0000001c020000290000003f0220003900000a2102200197000000400300043d0000000002230019001900000003001d000000000032004b0000000003000039000000010300403900000a200020009c000001140000213d0000000100300190000001140000c13d0000000101100039000000000101041a000000400020043f0000001d0200002900000019030000290000000003230436001800000003001d000000000002004b000014810000c13d000000400400043d000000200240003900000aa6010000410000000000120435000000150100002900000000010104330000002403400039000000000013043500000014010000290000000001010433000000000001004b0000000001000039000000010100c039000000440340003900000000001304350000004401000039000000000014043500000ab10040009c000001140000213d0000008001400039000000400010043f000000c003400039000000800500003900000000005304350000001303000029000000010330018f000000a0054000390000000000350435000000160300002900000000003104350000010003400039000000000504043300000000005304350000012003400039000000000005004b000013a10000613d000000000600001900000000073600190000000008260019000000000808043300000000008704350000002006600039000000000056004b0000139a0000413d000000000235001900000000000204350000001f0550003900000ace0550019700000000033500190000000005130049000000e0044000390000000000540435000000190400002900000000040404330000000000430435000000050540021000000000055300190000002007500039000000000004004b000016a40000c13d0000000002170049000005d10000013d00000022020000290000000301200210000000200110008900000a380310021f000000040020008c00000a380300804100000002050003670000001f01500360000000000401043b000000040620008c000000e00000413d00000022010000290000001b0110003900000ace011001970000003f0110003900000ace02100197000000400100043d0000000002210019000000000012004b0000000007000039000000010700403900000a200020009c000001140000213d0000000100700190000001140000c13d000000400020043f00000000026104360000001e09000029000000000090007c000000e00000213d0000001f070000290000000407700039000000000775034f00000ace086001980000001f0660018f0000000005820019000013de0000613d000000000907034f000000000a020019000000009b09043c000000000aba043600000000005a004b000013da0000c13d000000000334016f000000000006004b000013ec0000613d000000000487034f0000000306600210000000000705043300000000076701cf000000000767022f000000000404043b0000010006600089000000000464022f00000000046401cf000000000474019f000000000045043500000022041000290000001c04400039000000000004043500000aa60030009c000015d90000613d00000aa70030009c00001c410000c13d000000000101043300000a230010009c000000e00000213d000000200010008c000000e00000413d000000400100043d001500000001001d00000a240010009c000001140000213d0000000001020433000013470000013d001d00c00030003d000014050000013d00000020044000390000000000a9043500000000008404350000001e0050006c000001de0000813d000000005205043400000a200020009c000000e00000213d0000001f022000290000001d0620006900000a230060009c000000e00000213d000000400060008c000000e00000413d000000400800043d00000a240080009c000001140000213d0000004006800039000000400060043f0000002006200039000000000606043300000a200060009c000000e00000213d00000000096804360000004006200039000000000606043300000a200060009c000000e00000213d00000000022600190000003f06200039000000000016004b000000000700001900000a220700804100000a2206600197000000000006004b000000000a00001900000a220a00404100000a220060009c000000000a07c01900000000000a004b000000e00000c13d0000002006200039000000000c06043300000a2000c0009c000001140000213d0000000506c002100000003f0660003900000a2106600197000000400a00043d00000000066a00190000000000a6004b0000000007000039000000010700403900000a200060009c000001140000213d0000000100700190000001140000c13d000000400060043f0000000000ca0435000000400b200039000000e002c000c9000000000cb2001900000000001c004b000000e00000213d0000000000cb004b000014000000813d000000000d0a00190000000002b1004900000a230020009c000000e00000213d000000e00020008c000000e00000413d000000400e00043d00000a2400e0009c000001140000213d0000004002e00039000000400020043f00000000620b043400000a1f0020009c000000e00000213d000000000f2e0436000000400200043d00000a250020009c000001140000213d000000c007200039000000400070043f000000000606043300000a1a0060009c000000e00000213d00000000066204360000004007b00039000000000707043300000a1a0070009c000000e00000213d00000000007604350000006006b0003900000000060604330000ffff0060008c000000e00000213d000000400720003900000000006704350000008006b00039000000000606043300000a1a0060009c000000e00000213d00000060072000390000000000670435000000a006b00039000000000606043300000a1a0060009c000000e00000213d00000080072000390000000000670435000000c006b000390000000006060433000000000006004b0000000007000039000000010700c039000000000076004b000000e00000c13d000000200dd00039000000a007200039000000000067043500000000002f04350000000000ed0435000000e00bb000390000000000cb004b000014430000413d000014000000013d0000006002000039000000000300001900000018050000290000001c06000029000000000453001900000000002404350000002003300039000000000063004b000014850000413d00170a3a0010019b00000000040000190000001b0040006c00001ed50000813d002200000004001d00000006014002100000001a011000290000000202000367000000000112034f000000000101043b001f00000001001d00000a1f0010009c000000e00000213d00000022010000290000000501100210001e00000001001d001c00200010002d0000001c01200360000000000101043b0000000003000031000000200430006a0000009f0440008a00000a220540019700000a2206100197000000000756013f000000000056004b000000000500001900000a2205004041000000000041004b000000000600001900000a220600804100000a220070009c000000000506c019000000000005004b000000e00000c13d00000020051000290000004006500039000000000662034f000000000606043b00000000075300490000001f0770008a00000a220870019700000a2209600197000000000a89013f000000000089004b000000000800001900000a2208004041000000000076004b000000000700001900000a220700804100000a2200a0009c000000000807c019000000000008004b000000e00000c13d0000000005560019000000000652034f000000000606043b002300000006001d00000a200060009c000000e00000213d000000230630006a0000002005500039000000000065004b000000000700001900000a220700204100000a220660019700000a2205500197000000000865013f000000000065004b000000000500001900000a220500404100000a220080009c000000000507c019000000000005004b000000e00000c13d0000002305000029000000200050008c000014fe0000a13d0000002401000029000000000010043f0000000a01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b0000001f02000029000000000020043f000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000101041a000000700110027000000a1a01100197000000230010006b0000173a0000213d0000000003000031000000210130006a00000002020003670000001c05200360000000c30410008a000000000105043b000000000041004b000000000500001900000a220500804100000a220440019700000a2206100197000000000746013f000000000046004b000000000400001900000a220400404100000a220070009c000000000405c019000000000004004b000000e00000c13d00000020011000290000002004100039000000000442034f000000000404043b00000000051300490000001f0550008a00000a220650019700000a2207400197000000000867013f000000000067004b000000000600001900000a2206004041000000000054004b000000000500001900000a220500804100000a220080009c000000000605c019000000000006004b000000e00000c13d0000000001140019000000000412034f000000000504043b00000a200050009c000000e00000213d0000000004530049000000200610003900000a220140019700000a2207600197000000000817013f000000000017004b000000000100001900000a2201004041000000000046004b000000000400001900000a220400204100000a220080009c000000000104c019000000000001004b000000e00000c13d0000001f0150003900000ace011001970000003f0110003900000ace04100197000000400100043d0000000004410019000000000014004b0000000007000039000000010700403900000a200040009c000001140000213d0000000100700190000001140000c13d000000400040043f00000000045104360000000007650019000000000037004b000000e00000213d000000000362034f00000ace0650019800000000026400190000154e0000613d000000000703034f0000000008040019000000007907043c0000000008980436000000000028004b0000154a0000c13d0000001f075001900000155b0000613d000000000363034f0000000306700210000000000702043300000000076701cf000000000767022f000000000303043b0000010006600089000000000363022f00000000036301cf000000000373019f000000000032043500000000025400190000000000020435000000170200002900000ac00020009c000015670000c13d0000000002010433000000200020008c000016c30000c13d0000000002040433000004000220008a00000a940020009c000017400000813d0000002401000029000000000010043f0000000a01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b0000001f02000029000000000020043f000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000400200043d00000a250020009c000001140000213d000000000101043b000000c003200039000000400030043f000000000101041a000000700310027000000a1a033001970000008004200039000000000034043500000040031002700000ffff0330018f00000040042000390000000000340435000000200310027000000a1a033001970000002004200039000000000034043500000a1a03100197000000000032043500000a98001001980000000003000039000000010300c039000000a004200039002300000004001d00000000003404350000006002200039000000500110027000000a1a01100197001f00000002001d00000000001204350000002401000029000000000010043f0000000901000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b00000023020000290000000002020433000000000002004b000015b60000613d0000001f010000290000000001010433000015b80000013d000000000101041a000000d80110027000000a1a02100197000000400100043d000000200310003900000000002304350000002002000039000000000021043500000a240010009c000001140000213d0000004002100039000000400020043f000000190300002900000000020304330000002204000029000000000042004b00001ed50000a13d0000001e05000029000000180250002900000000001204350000000001030433000000000041004b00001ed50000a13d00000001044000390000001d0040006c0000148c0000413d000013750000013d0000001f0300002900000000030304330000ffff0330018f0000000f0030006c000015f00000813d000000400100043d00000aac0200004100001b850000013d000000000301043300000a230030009c000000e00000213d000000400030008c000000e00000413d000000400300043d001500000003001d00000a240030009c000001140000213d00000015040000290000004003400039000000400030043f0000000002020433000000000024043500000040011000390000000001010433000000000001004b0000000002000039000000010200c039001300000002001d000000000021004b000000e00000c13d0000134d0000013d0000001e03000029000000000303043300000a380330019700000a390030009c000015fc0000c13d0000000003010433000000200030008c000016c30000c13d0000000003050433000004000330008a00000a940030009c000016ef0000813d0000001201200360000000000101043b00000a1f0010009c000000e00000213d286326190000040f0000001d020000290000000002020433002000000002001d0000002402000029000000000020043f0000000502000039000000200020043f000400000001001d000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000400200043d001f00000002001d00000a240020009c000001140000213d0000002002000029001e0a1a0020019c000000000101043b0000001f030000290000004002300039000000400020043f000000000101041a0000002002300039000000e004100270002000000004001d000000000042043500000a2701100197000100000001001d00000000001304350000175a0000c13d0000000f0000006b0000185c0000c13d0000001b01000029000000000101043300000a1a0110019700200a99001000d5001b00000000001d001d00000000001d0000000002000031000000110320006a0000001201000029000000400610008a0000000201000367000000000461034f0000001f0530008a000000000404043b0000001c0300002900000000030304330000ffff03300190002400000000001d0000176e0000c13d000000000054004b000000000300001900000a220300804100000a220750019700000a2208400197000000000978013f000000000078004b000000000800001900000a220800404100000a220090009c000000000803c019000000000008004b000000e00000c13d000000150300002900000000030304330000001104400029000000000841034f000000000808043b00000a200080009c000000e00000213d00000000098200490000002004400039000000000094004b000000000a00001900000a220a00204100000a220990019700000a2204400197000000000b94013f000000000094004b000000000400001900000a220400404100000a2200b0009c00000000040ac019000000000004004b000000e00000c13d000000160400002900000000040404330000ffff0940018f00000000048900a9000000000008004b000016680000613d00000aa50880019700000aa50a40019700000000088a00d9000000000089004b000017680000c13d0000006006600039000000000661034f000000000606043b00000a2208600197000000000978013f000000000078004b000000000700001900000a2207004041000000000056004b000000000500001900000a220500804100000a220090009c000000000705c019000000000007004b000000e00000c13d0000001106600029000000000561034f000000000505043b00000a200050009c000000e00000213d0000000007520049000000200860003900000a220670019700000a2209800197000000000a69013f000000000069004b000000000600001900000a2206004041000000000078004b000000000700001900000a220700204100000a2200a0009c000000000607c019000000000006004b000000e00000c13d000000400600043d00000a240060009c000001140000213d0000004007600039000000400070043f000000200760003900000000000704350000000000060435000000400700043d00000a240070009c000001140000213d000000170600002900000000060604330000004009700039000000400090043f000000200970003900000000000904350000000000070435000000000005004b00001bf80000c13d000000400500043d00000a240050009c000001140000213d00000a1a0260019700001c3c0000013d00000000050000190000000006030019000000190d000029000016b00000013d000000000987001900000000000904350000001f0880003900000ace0880019700000000078700190000000105500039000000000045004b000013b10000813d0000000008370049000000200880008a00000020066000390000000000860435000000200dd0003900000000080d043300000000980804340000000007870436000000000008004b000016a80000613d000000000a000019000000000b7a0019000000000ca90019000000000c0c04330000000000cb0435000000200aa0003900000000008a004b000016bb0000413d000016a80000013d000000400400043d002400000004001d00000a9302000041000000000024043500000004024000390000002003000039000000000032043500000024024000392863201b0000040f0000002402000029000000000121004900000a1a0010009c00000a1a0100804100000a1a0020009c00000a1a0200804100000060011002100000004002200210000000000121019f00002865000104300000001805000029000000000621004900000a230060009c000000e00000213d000000400060008c000000e00000413d000000400600043d00000a240060009c000001140000213d0000004007600039000000400070043f000000008702043400000a1f0070009c000000e00000213d0000000007760436000000000808043300000a200080009c000000e00000213d0000002005500039000000000087043500000000006504350000004002200039000000000042004b000016d70000413d000002090000013d000000400200043d00000a930300004100000000003204350000000403200039000000200400003900000000004304350000000001010433000000240320003900000000001304350000004403200039000000000001004b000017540000613d000000000400001900000000063400190000000007540019000000000707043300000000007604350000002004400039000000000014004b000016fc0000413d000017540000013d00000023010000290000002402000029286324160000040f00000000010104330000000001010433000000400200043d00000aa103000041000000000032043500000a1f01100197000003760000013d00000023010000290000002402000029002400000002001d286324160000040f00000000010104330000000001010433002200000001001d00000023010000290000002402000029286324160000040f000000000101043300000040011000390000000001010433002100000001001d00000023010000290000002402000029286324160000040f0000000001010433000000000101043300000a1f01100197286323e40000040f000000210200002900000a1a02200197000000000101041a000000400300043d00000024043000390000000000240435000000e0011002700000004402300039000000000012043500000ab9010000410000000000130435000000220100002900000a1f011001970000000402300039000000000012043500000a1a0030009c00000a1a03008041000000400130021000000a61011001c70000286500010430000000400100043d00000aa00200004100001b850000013d000000400100043d00000ac102000041000000000021043500000004021000390000001f03000029000006500000013d000000400200043d00000a930300004100000000003204350000000403200039000000200500003900000000005304350000000001010433000000240320003900000000001304350000004403200039000000000001004b000017540000613d000000000500001900000000063500190000000007450019000000000707043300000000007604350000002005500039000000000015004b0000174d0000413d0000001f0410003900000ace04400197000000000131001900000000000104350000004401400039000006fc0000013d00000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f000000010020019000001a9b0000613d000000000101043b000000200110006c000017c20000813d00000abb01000041000000000010043f0000001101000039000000040010043f00000a4e010000410000286500010430000000000054004b000000000700001900000a220700804100000a220850019700000a2209400197000000000a89013f000000000089004b000000000800001900000a220800404100000a2200a0009c000000000807c019000000000008004b000000e00000c13d0000001108400029000000000781034f000000000707043b00000a200070009c000000e00000213d00000000097200490000002008800039000000000098004b000000000a00001900000a220a00204100000a220990019700000a2208800197000000000b98013f000000000098004b000000000800001900000a220800404100000a2200b0009c00000000080ac019000000000008004b000000e00000c13d0000000f0900002900000120089000c9000000000009004b000017960000613d0000000f098000fa000001200090008c000017680000c13d000001c007700039000000000087001a000017680000413d00000000078700190000001d0070002a000017680000413d0000001d0870002a000000130700002900000000070704330000ffff0970018f00000000078900a9000017a50000613d00000000088700d9000000000098004b000017680000c13d0000001408000029000000000808043300000a1a08800197000000000078001a000017680000413d000000000878001a002400000000001d0000163a0000613d0000000107000029000000700970027000000000079800a900000000088700d9000000000098004b000017680000c13d000000000007004b002400000000001d0000163a0000613d00000000083700a900000000077800d9000000000037004b000017680000c13d000000000008004b002400000000001d0000163a0000613d00240aa4008000d500000024038000f900000aa40030009c0000163a0000613d000017680000013d0000001e0010006c000018570000a13d000000400200043d0000004403200039000000000013043500000024012000390000001e03000029000000000031043500000a9701000041000000000012043500000004012000390000002403000029000000000031043500000a1a0020009c00000a1a02008041000000400120021000000a61011001c700002865000104300000001604000029000000000521004900000a230050009c000000e00000213d000002400050008c000000e00000413d000000400500043d00000a240050009c000001140000213d0000004006500039000000400060043f000000008602043400000a200060009c000000e00000213d0000000006650436000000400700043d00000a260070009c000001140000213d0000022009700039000000400090043f0000000008080433000000000008004b0000000009000039000000010900c039000000000098004b000000e00000c13d0000000008870436000000400920003900000000090904330000ffff0090008c000000e00000213d00000000009804350000006008200039000000000808043300000a1a0080009c000000e00000213d000000400970003900000000008904350000008008200039000000000808043300000a1a0080009c000000e00000213d00000060097000390000000000890435000000a008200039000000000808043300000a1a0080009c000000e00000213d00000080097000390000000000890435000000c00820003900000000080804330000ffff0080008c000000e00000213d000000a0097000390000000000890435000000e008200039000000000808043300000a1a0080009c000000e00000213d000000c0097000390000000000890435000001000820003900000000080804330000ffff0080008c000000e00000213d000000e0097000390000000000890435000001200820003900000000080804330000ffff0080008c000000e00000213d00000100097000390000000000890435000001400820003900000000080804330000ffff0080008c000000e00000213d000001200970003900000000008904350000016008200039000000000808043300000a1a0080009c000000e00000213d000001400970003900000000008904350000018008200039000000000808043300000a1a0080009c000000e00000213d00000160097000390000000000890435000001a008200039000000000808043300000a200080009c000000e00000213d00000180097000390000000000890435000001c008200039000000000808043300000a1a0080009c000000e00000213d000001a0097000390000000000890435000001e008200039000000000808043300000a1a0080009c000000e00000213d000001c009700039000000000089043500000200082000390000000008080433000000000008004b0000000009000039000000010900c039000000000098004b000000e00000c13d000001e00970003900000000008904350000022008200039000000000808043300000a2700800198000000e00000c13d00000200097000390000000000890435000000000076043500000000045404360000024002200039000000000032004b000017d50000413d000002340000013d0000001f01000029000000000101043300010a270010019b0000000f0000006b000016270000613d00000002010003670000001202100360000000000202043b000600000002001d00000a1f0020009c000000e00000213d0000001002100360000000000302043b0000000002000031000000230420006a000000230440008a00000a220540019700000a2206300197000000000756013f000000000056004b000000000500001900000a2205004041000000000043004b000000000400001900000a220400804100000a220070009c000000000504c019000000000005004b000000e00000c13d0000001103300029000000000131034f000000000101043b000a00000001001d00000a200010009c000000e00000213d0000000a0100002900000006011002100000000001120049000000200530003900000a220210019700000a2203500197000000000423013f000000000023004b000000000200001900000a2202004041000900000005001d000000000015004b000000000100001900000a220100204100000a220040009c000000000201c019000000000002004b000000e00000c13d0000000a0000006b001b00000000001d001d00000000001d002000000000001d0000162d0000613d001f00000000001d001d00000000001d001b00000000001d002000000000001d0000189b0000013d0000001f020000290000000102200039001f00000002001d0000000a0020006c0000162d0000813d0000001f0100002900000006011002100000000901100029000000000210007900000a230020009c000000e00000213d000000400020008c000000e00000413d000000400200043d002300000002001d00000a240020009c000001140000213d00000023020000290000004002200039000000400020043f0000000202000367000000000312034f000000000303043b00000a1f0030009c000000e00000213d000000230400002900000000033404360000002001100039000000000112034f000000000101043b000e00000003001d00000000001304350000002401000029000000000010043f0000000a01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b0000002302000029000000000202043300000a1f02200197000000000020043f000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000400200043d001e00000002001d00000a250020009c000001140000213d000000000101043b0000001e04000029000000c002400039000000400020043f000000000201041a000000a00140003900000a98002001980000000003000039000000010300c039000000000031043500000a1a012001970000000003140436000000700120027000000a1a011001970000008005400039001000000005001d0000000000150435000000200120027000000a1a01100197000b00000003001d0000000000130435000000500120027000000a1a011001970000006003400039000d00000003001d0000000000130435000000400440003900000040022002700000ffff0320018f000c00000004001d0000000000340435000019040000613d000000000003004b0000191a0000613d0000002302000029000000000202043300000a1f04200197000800000004001d000000060040006c0000191c0000c13d000000040000006b000019940000613d0000000e010000290000000001010433000000040300002900001a660000013d000000210100002900000000010104330000ffff0110018f00000a99011000d1000000200010002a000017680000413d0000001b0200002900000a1a022001970000002203000029000000000303043300000a1a033001970000000002230019001b00000002001d00000a1a0020009c000017680000213d0000001d0200002900000a1a0220019700000a9a0020009c000017680000213d002000200010002d001d00200020003d000018960000013d000000000200001900001a720000013d000000400100043d00000a240010009c000001140000213d0000004002100039000000400020043f0000002002100039000000000002043500000000000104350000000801000029000000000010043f0000000601000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000400200043d002300000002001d00000a240020009c000001140000213d000000000101043b00000023030000290000004002300039000000400020043f000000000101041a00000a27021001970000000002230436000000e001100270000700000002001d000500000001001d000000000012043500000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f000000010020019000001a9b0000613d000000000101043b0005000500100074000017680000413d00000a9b0100004100000000001004430000000001000412000000040010044300000040010000390000002400100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a9c011001c700008005020000392863285e0000040f000000010020019000001a9b0000613d000000000101043b00000a1a01100197000000050010006b00001a5c0000413d0000000801000029000000000010043f0000000701000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000400200043d00000a240020009c000001140000213d000000000101043b0000004003200039000000400030043f000000000101041a00000a1f031001980000000003320436000000a001100270000000ff0110018f000300000003001d000000000013043500001a5c0000613d000000400100043d00000a240010009c000001140000213d0000004003100039000000400030043f0000002003100039000000000003043500000000000104350000000002020433000000400300043d00000a9d01000041002300000003001d0000000001130436000200000001001d000000000100041400000a1f02200197000500000002001d000000040020008c000019960000c13d0000000103000031000000a00030008c000000a0040000390000000004034019000019bf0000013d000000000200001900001a6e0000013d000000230200002900000a1a0020009c00000a1a02008041000000400220021000000a1a0010009c00000a1a01008041000000c001100210000000000121019f00000a5e011001c700000005020000292863285e0000040f000000600310027000000a1a03300197000000a00030008c000000a0040000390000000004034019000000e0064001900000002305600029000019af0000613d000000000701034f0000002308000029000000007907043c0000000008980436000000000058004b000019ab0000c13d0000001f07400190000019bc0000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f000000010020019000001c440000613d0000001f01400039000001e00110018f0000002302100029000000000012004b00000000010000390000000101004039000700000002001d00000a200020009c000001140000213d0000000100100190000001140000c13d0000000701000029000000400010043f000000a00030008c000000e00000413d0000002301000029000000000101043300000a9e0010009c000000e00000213d00000023010000290000008001100039000000000101043300000a9e0010009c000000e00000213d00000002010000290000000001010433000200000001001d00000a220010009c00001c500000813d00000a9f010000410000000702000029000000000012043500000000010004140000000502000029000000040020008c000000200400003900001a0d0000613d000000070200002900000a1a0020009c00000a1a02008041000000400220021000000a1a0010009c00000a1a01008041000000c001100210000000000121019f00000a5e011001c700000005020000292863285e0000040f000000600310027000000a1a03300197000000200030008c0000002004000039000000000403401900000020064001900000000705600029000019fd0000613d000000000701034f0000000708000029000000007907043c0000000008980436000000000058004b000019f90000c13d0000001f0740019000001a0a0000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f000000010020019000001c530000613d0000001f01400039000000600110018f0000000701100029002300000001001d00000a200010009c000001140000213d0000002301000029000000400010043f000000200030008c000000e00000413d00000007010000290000000001010433000000ff0010008c000000e00000213d00000003020000290000000002020433000000ff0220018f0000000002120019000000ff0020008c000017680000213d000000240120008c00001a350000213d000000010100003900001a2e0000613d00000024022000890000000a03000039000000010020019000000000043300a9000000010300603900000000011300a90000000102200272000000000304001900001a270000c13d000000020000006b00001a960000613d00000002021000b900000002032000fa000000000013004b00001a430000613d000017680000013d0000004d0010008c000017680000213d00000001020000390000000a03000039000000010010019000000000043300a9000000010300603900000000022300a90000000101100272000000000304001900001a390000c13d000000000002004b00001f7c0000613d00000002022000f900000a270020009c00001c5f0000213d000000230100002900000a240010009c000001140000213d00000023030000290000004001300039000000400010043f0000000001230436000700000001001d00000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f000000010020019000001a9b0000613d000000000101043b00000a1a01100197000000070200002900000000001204350000000701000029000000000101043300000a1a0010019800001bf20000613d0000002301000029000000000101043300000a270310019800001bf20000613d0000000e01000029000000000101043300000000023100a900000000033200d9000000000013004b000017680000c13d0000000d0100002900000000010104330000000c0300002900000000030304330000ffff0330018f00000aa20220012a00000000023200a900000aa30220012a0000001b0300002900000a1a0330019700000a1a011001970000000001310019001b00000001001d00000a1a0010009c000017680000213d0000001d0100002900000a1a011001970000001003000029000000000303043300000a1a033001970000000001130019001d00000001001d00000a1a0010009c000017680000213d0000001e01000029000000000101043300000a1a0110019700000a99011000d1000000000012004b00001a8e0000413d0000000b01000029000000000101043300000a1a0110019700000a99011000d1000000000012004b00001a920000a13d000000200010002a000017680000413d002000200010002d000018960000013d000000200020002a000017680000413d002000200020002d000018960000013d0000000002000019000000230100002900000a240010009c00001a480000a13d000001140000013d000000000001042f000000000200041a00000a2802200197000000000112019f000000000010041b0000001d0100002900000a290010009c000001140000213d0000001d010000290000002002100039001900000002001d000000400020043f0000000000010435000000400100043d001500000001001d00000a240010009c000001140000213d00000015030000290000004001300039000000400010043f00000020013000390000001d020000290000000000210435000000210100002900000000001304350000000001020433000000000001004b00001aca0000c13d00000021010000290000000001010433000000000001004b00001b350000c13d0000002301000029000000000101043300000a1f0110019800001ac70000613d0000002402000029000000000202043300000a1e0020019800001ac70000613d0000002202000029000000000202043300000a1a0020019800001b8f0000c13d000000400100043d00000a5d0200004100001b850000013d001f00000000001d00001ad20000013d0000001f02000029001f00010020003d0000001d0100002900000000010104330000001f0010006b00001b8b0000813d0000001f0100002900000005011002100000001901100029000000000101043300000a1f01100197001e00000001001d000000000010043f0000000301000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000101041a002100000001001d000000000001004b00001acc0000613d0000000201000039000000000201041a000000000002004b000017680000613d0000002103000029000000010130008a000000000032004b00001b0e0000613d000000000012004b00001ed50000a13d000000210100002900000a2b0110009a00000a2b0220009a000000000202041a000000000021041b000000000020043f0000000301000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b0000002102000029000000000021041b0000000201000039000000000101041a002100000001001d000000000001004b00001edb0000613d0000002101000029000000010110008a000000210200002900000a2b0220009a000000000002041b0000000202000039000000000012041b0000001e01000029000000000010043f0000000301000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000001041b000000400100043d0000001e02000029000000000021043500000a1a0010009c00000a1a010080410000004001100210000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a2c011001c70000800d02000039000000010300003900000a2d04000041286328590000040f000000010020019000001acc0000c13d000000e00000013d0000002101000029001d00200010003d001e00000000001d0000001e0100002900000005011002100000001d011000290000000001010433001f0a1f0010019c00001b830000613d0000001f01000029000000000010043f0000000301000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000101041a000000000001004b00001b6a0000c13d0000000201000039000000000101041a00000a200010009c000001140000213d00000001021000390000000203000039000000000023041b00000a2e0110009a0000001f02000029000000000021041b000000000103041a001900000001001d000000000020043f0000000301000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b0000001902000029000000000021041b000000400100043d0000001f02000029000000000021043500000a1a0010009c00000a1a010080410000004001100210000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a2c011001c70000800d02000039000000010300003900000a2f04000041286328590000040f0000000100200190000000e00000613d0000001e02000029001e00010020003d000000210100002900000000010104330000001e0010006b00001b380000413d00001abb0000013d000000400100043d00000ab202000041000000000021043500000a1a0010009c00000a1a01008041000000400110021000000a5e011001c7000028650001043000000015010000290000000001010433002100000001001d00001ab70000013d000000a00010043f0000002401000029000000000101043300000a1e01100197000000800010043f0000002201000029000000000101043300000a1a01100197000000c00010043f000000400100043d001f00000001001d00000a290010009c000001140000213d0000001f010000290000002002100039001e00000002001d000000400020043f000000000001043500000020010000290000000001010433000000000001004b00001c660000613d002400000000001d00001bad0000013d0000002402000029002400010020003d00000020010000290000000001010433000000240010006b00001c620000813d000000240100002900000005011002100000001c01100029002200000001001d000000000101043300000a1f01100197002300000001001d000000000010043f0000000c01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000101041a000000000001004b00001ba70000c13d0000000b01000039000000000101041a00000a200010009c000001140000213d00000001021000390000000b03000039000000000023041b00000a300110009a0000002302000029000000000021041b000000000103041a002100000001001d000000000020043f0000000c01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b0000002102000029000000000021041b00000020010000290000000001010433000000240010006c00001ed50000a13d00000022010000290000000001010433000000000200041400000a1f0510019700000a1a0020009c00000a1a02008041000000c00120021000000a31011001c70000800d02000039000000020300003900000a3204000041286328590000040f000000010020019000001ba70000c13d000000e00000013d000000400100043d00000aa102000041000000000021043500000004021000390000000803000029000006500000013d000000000681034f0000000307500210000000200770008900000a380970021f000000040b50008c00000a3809008041000000000a06043b000000e00000413d0000001b0650003900000ace066001970000003f0660003900000ace07600197000000400600043d0000000007760019000000000067004b000000000c000039000000010c00403900000a200070009c000001140000213d0000000100c00190000001140000c13d000000400070043f0000000007b60436000000000c85001900000000002c004b000000e00000213d0000000402800039000000000821034f00000ace0cb001980000001f0bb0018f0000000002c7001900001c1e0000613d000000000d08034f000000000e07001900000000df0d043c000000000efe043600000000002e004b00001c1a0000c13d00000000099a016f00000000000b004b00001c2c0000613d0000000008c8034f000000030ab00210000000000b020433000000000bab01cf000000000bab022f000000000808043b000001000aa000890000000008a8022f0000000008a801cf0000000008b8019f000000000082043500000000026500190000001c02200039000000000002043500000aa60090009c00001e4c0000613d00000aa70090009c00001c410000c13d000000000206043300000a230020009c000000e00000213d000000200020008c000000e00000413d000000400500043d00000a240050009c000001140000213d00000000020704330000004006500039000000400060043f0000000000250435000000000600001900001e5f0000013d000000400100043d00000abf0200004100001b850000013d0000001f0530018f00000a1c06300198000000400200043d00000000046200190000127e0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001c4b0000c13d0000127e0000013d00000aa0010000410000000702000029000012610000013d0000001f0530018f00000a1c06300198000000400200043d00000000046200190000127e0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001c5a0000c13d0000127e0000013d00000aa0010000410000002302000029000012610000013d0000001f010000290000000001010433000000000001004b00001e690000c13d0000001b010000290000000001010433000000000001004b00001ca90000613d002400000000001d000000240100002900000005011002100000001a01100029000000000101043300000020021000390000000002020433002200000002001d000000000101043300000a1f01100197002300000001001d000000000010043f0000000701000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000201041a00000a3502200197000000220500002900000020035000390000000004030433000000a00440021000000a3604400197000000000242019f000000000405043300000a1f04400197000000000242019f000000000021041b000000400100043d00000000024104360000000003030433000000ff0330018f000000000032043500000a1a0010009c00000a1a010080410000004001100210000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a2a011001c70000800d02000039000000020300003900000a37040000410000002305000029286328590000040f0000000100200190000000e00000613d0000002402000029002400010020003d0000001b010000290000000001010433000000240010006b00001c6b0000413d00000017010000290000000001010433000000000001004b00001d710000613d002200000000001d0000002201000029000000050110021000000016011000290000000001010433000000001201043400230a200020019c00001ee10000613d0000000001010433002400000001001d0000016001100039002000000001001d000000000101043300000a1a0110019800001ee10000613d00000024020000290000020002200039002100000002001d000000000202043300000a380220019700000a390020009c00001ee10000c13d00000024020000290000006002200039001f00000002001d000000000202043300000a1a02200197000000000021004b00001ee10000213d0000002301000029000000000010043f0000000901000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b0000000101100039000000000101041a001e00000001001d000000400100043d001d00000001001d000000240200002928631f850000040f0000001d02000029000000000121004900000a1a0020009c00000a1a0200804100000a1a0010009c00000a1a0100804100000040022002100000006001100210000000000121019f0000001e0200002900000a3a0020019800001cf50000613d000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a31011001c70000800d02000039000000020300003900000a3b0400004100001cfe0000013d000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a31011001c70000800d02000039000000020300003900000a3c040000410000002305000029286328590000040f0000000100200190000000e00000613d0000002301000029000000000010043f0000000901000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d00000024040000290000000032040434000000000002004b000000000101043b000000000201041a00000a3d02200197000000010220c1bf0000000003030433000000080330021000000a3e03300197000000000232019f00000040034000390000000003030433000000180330021000000a3f03300197000000000232019f0000001f030000290000000003030433000000380330021000000a4003300197000000000232019f00000080034000390000000003030433000000580330021000000a4103300197000000000232019f000000a0034000390000000003030433000000780330021000000a4203300197000000000232019f000000c0034000390000000003030433000000880330021000000a4303300197000000000232019f000000e0034000390000000003030433000000a80330021000000a4403300197000000000232019f00000100034000390000000003030433000000b80330021000000a4503300197000000000232019f00000120034000390000000003030433000000c80330021000000a4603300197000000000232019f00000140034000390000000003030433000000d80330021000000a4703300197000000000232019f000000000021041b00000001011000390000002002000029000000000202043300000a1a02200197000000000301041a00000a4803300197000000000223019f00000180034000390000000003030433000000200330021000000a4903300197000000000232019f000001a0034000390000000003030433000000600330021000000a4a03300197000000000232019f000001c0034000390000000003030433000000800330021000000a4b03300197000000000232019f000001e0034000390000000003030433000000000003004b00000a4c030000410000000003006019000000000232019f00000021030000290000000003030433000000380330027000000a3a03300197000000000232019f000000000021041b0000002202000029002200010020003d00000017010000290000000001010433000000220010006b00001cae0000413d0018002b0000002d00000018010000290000000001010433000000000001004b00001dad0000613d0000001801000029002100200010003d002400000000001d000000240100002900000005011002100000002101100029000000000101043300000020021000390000000002020433002200000002001d000000000101043300000a1f01100197002300000001001d000000000010043f0000000801000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000220200002900000a2002200197000000000101043b000000000301041a00000a4f03300197000000000323019f000000000031041b000000400100043d000000000021043500000a1a0010009c00000a1a010080410000004001100210000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a2c011001c70000800d02000039000000020300003900000a50040000410000002305000029286328590000040f0000000100200190000000e00000613d0000002402000029002400010020003d00000018010000290000000001010433000000240010006b00001d780000413d000000400100043d001b00000001001d00000a290010009c000001140000213d0000001b010000290000002002100039001a00000002001d000000400020043f00000000000104350000002c01000029001c00000001001d0000000021010434001d00000002001d000000000001004b00001f230000613d001e00000000001d00001dc20000013d0000001e02000029001e00010020003d0000001e0010006b00001ee40000813d0000001e0200002900000005022002100000001d0220002900000000030204330000002002300039001f00000002001d00000000020204330000000004020433000000000004004b00001dbe0000613d000000000103043300220a200010019b002400000000001d00000024010000290000000501100210000000000112001900000020011000390000000001010433000000002101043400230a1f0010019b0000000001020433002000000001001d0000008001100039002100000001001d000000000101043300000a1a011001970000001f0010008c00001f330000a13d0000002201000029000000000010043f0000000a01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b0000002302000029000000000020043f000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000201041a00000a5102200197000000200900002900000020039000390000000004030433000000200440021000000a5204400197000000000242019f00000040049000390000000005040433000000400550021000000a5305500197000000000252019f00000060059000390000000006050433000000500660021000000a5406600197000000000262019f00000021080000290000000006080433000000700660021000000a5506600197000000000262019f000000a0069000390000000007060433000000000007004b00000a56070000410000000007006019000000000272019f000000000709043300000a1a07700197000000000272019f000000000021041b000000400100043d0000000002710436000000000303043300000a1a03300197000000000032043500000000020404330000ffff0220018f00000040031000390000000000230435000000000205043300000a1a0220019700000060031000390000000000230435000000000208043300000a1a02200197000000800310003900000000002304350000000002060433000000000002004b0000000002000039000000010200c039000000a003100039000000000023043500000a1a0010009c00000a1a010080410000004001100210000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a57011001c70000800d02000039000000030300003900000a580400004100000022050000290000002306000029286328590000040f0000000100200190000000e00000613d0000002403000029002400010030003d0000001f0100002900000000020104330000000001020433000000240010006b00001dcf0000413d0000001c01000029000000000101043300001dbe0000013d000000000206043300000a230020009c000000e00000213d000000400020008c000000e00000413d000000400500043d00000a240050009c000001140000213d0000004002500039000000400020043f0000000002070433000000000025043500000040066000390000000006060433000000000006004b0000000007000039000000010700c039000000000076004b000000e00000c13d000000200550003900000000006504350000001905000029000000000505043300000a1a05500197000000000052004b00001f3b0000a13d000000400100043d00000aab0200004100001b850000013d002400000000001d00001e710000013d0000002402000029002400010020003d0000001f010000290000000001010433000000240010006b00001c660000813d000000240100002900000005011002100000001e01100029002200000001001d000000000101043300000a1f01100197002100000001001d000000000010043f0000000c01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000101041a002300000001001d000000000001004b00001e6b0000613d0000000b01000039000000000201041a000000000002004b000017680000613d0000002303000029000000010130008a000000000023004b00001eae0000613d000000000012004b00001ed50000a13d000000230100002900000a330110009a00000a330220009a000000000202041a000000000021041b000000000020043f0000000c01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b0000002302000029000000000021041b0000000b01000039000000000101041a002300000001001d000000000001004b00001edb0000613d0000002301000029000000010110008a000000230200002900000a330220009a000000000002041b0000000b02000039000000000012041b0000002101000029000000000010043f0000000c01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000001041b0000001f010000290000000001010433000000240010006c00001ed50000a13d00000022010000290000000001010433000000000200041400000a1f0510019700000a1a0020009c00000a1a02008041000000c00120021000000a31011001c70000800d02000039000000020300003900000a3404000041286328590000040f000000010020019000001e6b0000c13d000000e00000013d00000abb01000041000000000010043f0000003201000039000000040010043f00000a4e01000041000028650001043000000abb01000041000000000010043f0000003101000039000000040010043f00000a4e010000410000286500010430000000400100043d00000a4d020000410000064d0000013d0000001b010000290000000001010433000000000001004b00001f230000613d002400000000001d000000240100002900000005011002100000001a01100029000000000101043300000020021000390000000002020433002300000002001d000000000101043300000a2001100197002200000001001d000000000010043f0000000a01000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000230200002900000a1f02200197000000000101043b002300000002001d000000000020043f000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000001041b000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a31011001c70000800d02000039000000030300003900000a5b0400004100000022050000290000002306000029286328590000040f0000000100200190000000e00000613d0000002402000029002400010020003d0000001b010000290000000001010433000000240010006b00001ee90000413d000000800100043d00000140000004430000016000100443000000a00100043d00000020020000390000018000200443000001a000100443000000c00100043d0000004003000039000001c000300443000001e00010044300000100002004430000000301000039000001200010044300000a5c01000041000028640001042e000000400200043d0000002403200039000000000013043500000a5901000041000000000012043500000004012000390000002303000029000012980000013d000000000006004b00001f440000c13d00000018050000290000000005050433000000000005004b00001f440000613d000000400100043d00000aaa0200004100001b850000013d00000a1a033001970000001b033000290000000003430019000000000032001a000017680000413d000000010400002900000aa804400198002300000000001d00001f5c0000613d000000000332001900000000024300a900000000044200d9000000000034004b000017680000c13d000000000002004b002300000000001d00001f5c0000613d0000001a03000029000000000303043300000a200330019700230000002300ad00000023022000f9000000000032004b000017680000c13d0000001201100360000000000101043b00000a1f0010009c000000e00000213d000000000010043f0000000801000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000000e00000613d000000000101043b000000000101041a00000a200210019700000020012000b9000000200000006b00001f750000613d00000020031000fa000000000023004b000017680000c13d000000230010002a000017680000413d0000002301100029000000240010002a000017680000413d000000040000006b00001f820000c13d00000abb01000041000000000010043f0000001201000039000000040010043f00000a4e010000410000286500010430000000240110002900000004011000fa00000a150000013d0000000043020434000000000003004b0000000003000039000000010300c039000000000331043600000000040404330000ffff0440018f00000000004304350000004003200039000000000303043300000a1a03300197000000400410003900000000003404350000006003200039000000000303043300000a1a03300197000000600410003900000000003404350000008003200039000000000303043300000a1a0330019700000080041000390000000000340435000000a00320003900000000030304330000ffff0330018f000000a0041000390000000000340435000000c003200039000000000303043300000a1a03300197000000c0041000390000000000340435000000e00320003900000000030304330000ffff0330018f000000e0041000390000000000340435000001000320003900000000030304330000ffff0330018f00000100041000390000000000340435000001200320003900000000030304330000ffff0330018f000001200410003900000000003404350000014003200039000000000303043300000a1a03300197000001400410003900000000003404350000016003200039000000000303043300000a1a03300197000001600410003900000000003404350000018003200039000000000303043300000a200330019700000180041000390000000000340435000001a003200039000000000303043300000a1a03300197000001a0041000390000000000340435000001c003200039000000000303043300000a1a03300197000001c0041000390000000000340435000001e0032000390000000003030433000000000003004b0000000003000039000000010300c039000001e00410003900000000003404350000020002200039000000000202043300000a3802200197000002000310003900000000002304350000022001100039000000000001042d00000acf0010009c00001fe10000813d0000006001100039000000400010043f000000000001042d00000abb01000041000000000010043f0000004101000039000000040010043f00000a4e01000041000028650001043000000ad00010009c00001fec0000813d0000004001100039000000400010043f000000000001042d00000abb01000041000000000010043f0000004101000039000000040010043f00000a4e01000041000028650001043000000ad10010009c00001ff70000813d000000c001100039000000400010043f000000000001042d00000abb01000041000000000010043f0000004101000039000000040010043f00000a4e0100004100002865000104300000001f0220003900000ace022001970000000001120019000000000021004b0000000002000039000000010200403900000a200010009c000020090000213d0000000100200190000020090000c13d000000400010043f000000000001042d00000abb01000041000000000010043f0000004101000039000000040010043f00000a4e010000410000286500010430000000400100043d00000ad20010009c000020150000813d0000022002100039000000400020043f000000000001042d00000abb01000041000000000010043f0000004101000039000000040010043f00000a4e01000041000028650001043000000000430104340000000001320436000000000003004b000020270000613d000000000200001900000000052100190000000006240019000000000606043300000000006504350000002002200039000000000032004b000020200000413d000000000231001900000000000204350000001f0230003900000ace022001970000000001210019000000000001042d00000020030000390000000004310436000000000302043300000000003404350000004001100039000000000003004b0000203c0000613d00000000040000190000002002200039000000000502043300000a1f0550019700000000015104360000000104400039000000000034004b000020350000413d000000000001042d000000003101043400000a27011001970000000001120436000000000203043300000a1a022001970000000000210435000000000001042d0000000043020434000000000003004b0000000003000039000000010300c039000000000331043600000000040404330000ffff0440018f00000000004304350000004003200039000000000303043300000a1a03300197000000400410003900000000003404350000006003200039000000000303043300000a1a03300197000000600410003900000000003404350000008003200039000000000303043300000a1a0330019700000080041000390000000000340435000000a00320003900000000030304330000ffff0330018f000000a0041000390000000000340435000000c003200039000000000303043300000a1a03300197000000c0041000390000000000340435000000e00320003900000000030304330000ffff0330018f000000e0041000390000000000340435000001000320003900000000030304330000ffff0330018f00000100041000390000000000340435000001200320003900000000030304330000ffff0330018f000001200410003900000000003404350000014003200039000000000303043300000a1a03300197000001400410003900000000003404350000016003200039000000000303043300000a1a03300197000001600410003900000000003404350000018003200039000000000303043300000a200330019700000180041000390000000000340435000001a003200039000000000303043300000a1a03300197000001a0041000390000000000340435000001c003200039000000000303043300000a1a03300197000001c0041000390000000000340435000001e0032000390000000003030433000000000003004b0000000003000039000000010300c039000001e00410003900000000003404350000020002200039000000000202043300000a3802200197000002000310003900000000002304350000022001100039000000000001042d000000004302043400000a1a033001970000000003310436000000000404043300000a1a044001970000000000430435000000400320003900000000030304330000ffff0330018f000000400410003900000000003404350000006003200039000000000303043300000a1a03300197000000600410003900000000003404350000008003200039000000000303043300000a1a0330019700000080041000390000000000340435000000a0022000390000000002020433000000000002004b0000000002000039000000010200c039000000a0031000390000000000230435000000c001100039000000000001042d000000003202043400000a1f0220019700000000022104360000000003030433000000ff0330018f00000000003204350000004001100039000000000001042d0007000000000002000300000003001d000600000002001d000000400200043d00000ad00020009c0000236d0000813d0000004003200039000000400030043f00000020032000390000000000030435000000000002043500000a1f01100197000200000001001d000000000010043f0000000601000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000023730000613d000000400300043d00000a240030009c0000236d0000213d000000000101043b0000004002300039000000400020043f000000000101041a00000a2702100197000700000003001d0000000002230436000000e001100270000500000002001d000400000001001d000000000012043500000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f0000000100200190000023750000613d000000000101043b0004000400100074000023760000413d00000a9b0100004100000000001004430000000001000412000000040010044300000040010000390000002400100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a9c011001c700008005020000392863285e0000040f0000000100200190000023750000613d000000000101043b00000a1a01100197000000040010006b0000210c0000813d00000007040000290000000501000029000022090000013d0000000201000029000000000010043f0000000701000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000023730000613d000000400200043d00000a240020009c00000007040000290000236d0000213d000000000101043b0000004003200039000000400030043f000000000101041a00000a1f031001980000000003320436000000a001100270000000ff0110018f00000000001304350000213d0000613d000500000003001d000000400100043d00000a240010009c0000236d0000213d0000004003100039000000400030043f0000002003100039000000000003043500000000000104350000000002020433000000400c00043d00000a9d0100004100000000051c0436000000000100041400000a1f02200197000000040020008c0000213f0000c13d0000000103000031000000a00030008c000000a00400003900000000040340190000216e0000013d0000000501000029000022090000013d000100000005001d00000a1a00c0009c00000a1a0300004100000000030c4019000000400330021000000a1a0010009c00000a1a01008041000000c001100210000000000131019f00000a5e011001c7000400000002001d00070000000c001d2863285e0000040f000000070c000029000000600310027000000a1a03300197000000a00030008c000000a00400003900000000040340190000001f0640018f000000e00740019000000000057c00190000215c0000613d000000000801034f00000000090c0019000000008a08043c0000000009a90436000000000059004b000021580000c13d000000000006004b000021690000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000023a20000613d000000040200002900000001050000290000001f01400039000001e00110018f000000000bc1001900000000001b004b0000000001000039000000010100403900000a2000b0009c0000236d0000213d00000001001001900000236d0000c13d0000004000b0043f000000a00030008c000023730000413d00000000010c043300000a9e0010009c000023730000213d0000008001c00039000000000101043300000a9e0010009c000023730000213d000000000505043300000a220050009c0000238d0000813d00000a9f0100004100000000001b04350000000001000414000000040020008c0000218c0000c13d0000002004000039000021b90000013d000400000005001d00000a1a00b0009c00000a1a0300004100000000030b4019000000400330021000000a1a0010009c00000a1a01008041000000c001100210000000000131019f00000a5e011001c700070000000b001d2863285e0000040f000000070b000029000000600310027000000a1a03300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b0019000021a80000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b000021a40000c13d000000000006004b000021b50000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000023ae0000613d00000004050000290000001f01400039000000600110018f0000000001b1001900000a200010009c0000236d0000213d000700000001001d000000400010043f000000200030008c000023730000413d00000000010b0433000000ff0010008c000023730000213d00000005020000290000000002020433000000ff0220018f0000000002120019000000ff0020008c000023760000213d000000240120008c000021d00000213d000021df0000c13d0000000101000039000021e90000013d0000004d0010008c000023760000213d0000000a030000390000000102000039000000010010019000000000043300a9000000010300603900000000022300a900000001011002720000000003040019000021d40000c13d000000000002004b0000239c0000613d00000000022500d9000021ef0000013d0000000a0300003900000001010000390000002402200089000000010020019000000000043300a9000000010300603900000000011300a900000001022002720000000003040019000021e20000c13d000000000005004b000023640000613d00000000025100a900000000035200d9000000000013004b000023760000c13d00000a270020009c000023940000213d000000070300002900000a240030009c0000236d0000213d0000004001300039000000400010043f0000000001230436000500000001001d00000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f0000000100200190000023750000613d000000000101043b00000a1a011001970000000502000029000000000012043500000000010200190000000704000029000000000101043300000a1a001001980000237c0000613d000000000104043300000a27011001980000237c0000613d00050006001000bd000000060000006b000022160000613d000000060300002900000005023000f9000000000012004b000023760000c13d000000400100043d00000a240010009c0000236d0000213d0000004002100039000000400020043f000000200210003900000000000204350000000000010435000000030100002900000a1f01100197000300000001001d000000000010043f0000000601000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000023730000613d000000400300043d00000a240030009c0000236d0000213d000000000101043b0000004002300039000000400020043f000000000101041a00000a2702100197000700000003001d0000000002230436000000e001100270000600000002001d000400000001001d000000000012043500000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f0000000100200190000023750000613d000000000101043b0004000400100074000023760000413d00000a9b0100004100000000001004430000000001000412000000040010044300000040010000390000002400100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a9c011001c700008005020000392863285e0000040f0000000100200190000023750000613d000000000101043b00000a1a01100197000000040010006b0000225f0000813d000000070400002900000006010000290000235c0000013d0000000301000029000000000010043f0000000701000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000023730000613d000000400200043d00000a240020009c00000007040000290000236d0000213d000000000101043b0000004003200039000000400030043f000000000101041a00000a1f031001980000000003320436000000a001100270000000ff0110018f0000000000130435000022900000613d000600000003001d000000400100043d00000a240010009c0000236d0000213d0000004003100039000000400030043f0000002003100039000000000003043500000000000104350000000002020433000000400c00043d00000a9d0100004100000000051c0436000000000100041400000a1f02200197000000040020008c000022920000c13d0000000103000031000000a00030008c000000a0040000390000000004034019000022c10000013d00000006010000290000235c0000013d000200000005001d00000a1a00c0009c00000a1a0300004100000000030c4019000000400330021000000a1a0010009c00000a1a01008041000000c001100210000000000131019f00000a5e011001c7000400000002001d00070000000c001d2863285e0000040f000000070c000029000000600310027000000a1a03300197000000a00030008c000000a00400003900000000040340190000001f0640018f000000e00740019000000000057c0019000022af0000613d000000000801034f00000000090c0019000000008a08043c0000000009a90436000000000059004b000022ab0000c13d000000000006004b000022bc0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000023ba0000613d000000040200002900000002050000290000001f01400039000001e00110018f000000000bc1001900000000001b004b0000000001000039000000010100403900000a2000b0009c0000236d0000213d00000001001001900000236d0000c13d0000004000b0043f000000a00030008c000023730000413d00000000010c043300000a9e0010009c000023730000213d0000008001c00039000000000101043300000a9e0010009c000023730000213d000000000505043300000a220050009c0000238d0000813d00000a9f0100004100000000001b04350000000001000414000000040020008c000022df0000c13d00000020040000390000230c0000013d000400000005001d00000a1a00b0009c00000a1a0300004100000000030b4019000000400330021000000a1a0010009c00000a1a01008041000000c001100210000000000131019f00000a5e011001c700070000000b001d2863285e0000040f000000070b000029000000600310027000000a1a03300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b0019000022fb0000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b000022f70000c13d000000000006004b000023080000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000023c60000613d00000004050000290000001f01400039000000600110018f0000000001b1001900000a200010009c0000236d0000213d000700000001001d000000400010043f000000200030008c000023730000413d00000000010b0433000000ff0010008c000023730000213d00000006020000290000000002020433000000ff0220018f0000000002120019000000ff0020008c000023760000213d000000240120008c000023230000213d000023320000c13d00000001010000390000233c0000013d0000004d0010008c000023760000213d0000000a030000390000000102000039000000010010019000000000043300a9000000010300603900000000022300a900000001011002720000000003040019000023270000c13d000000000002004b0000239c0000613d00000000022500d9000023420000013d0000000a0300003900000001010000390000002402200089000000010020019000000000043300a9000000010300603900000000011300a900000001022002720000000003040019000023350000c13d000000000005004b000023690000613d00000000025100a900000000035200d9000000000013004b000023760000c13d00000a270020009c000023940000213d000000070300002900000a240030009c0000236d0000213d0000004001300039000000400010043f0000000001230436000600000001001d00000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f0000000100200190000023750000613d000000000101043b00000a1a011001970000000602000029000000000012043500000000010200190000000704000029000000000101043300000a1a00100198000023820000613d000000000104043300000a2701100198000023820000613d00000005011000f9000000000001042d0000000002000019000000070300002900000a240030009c000021f40000a13d0000236d0000013d0000000002000019000000070300002900000a240030009c000023470000a13d00000abb01000041000000000010043f0000004101000039000000040010043f00000a4e01000041000028650001043000000000010000190000286500010430000000000001042f00000abb01000041000000000010043f0000001101000039000000040010043f00000a4e010000410000286500010430000000400100043d00000aa102000041000000000021043500000004021000390000000203000029000023870000013d000000400100043d00000aa102000041000000000021043500000004021000390000000303000029000000000032043500000a1a0010009c00000a1a01008041000000400110021000000a4e011001c7000028650001043000000aa00100004100000000001b043500000a1a00b0009c00000a1a0b0080410000004001b0021000000a5e011001c7000028650001043000000aa0010000410000000702000029000000000012043500000a1a0020009c00000a1a02008041000000400120021000000a5e011001c7000028650001043000000abb01000041000000000010043f0000001201000039000000040010043f00000a4e0100004100002865000104300000001f0530018f00000a1c06300198000000400200043d0000000004620019000023d10000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000023a90000c13d000023d10000013d0000001f0530018f00000a1c06300198000000400200043d0000000004620019000023d10000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000023b50000c13d000023d10000013d0000001f0530018f00000a1c06300198000000400200043d0000000004620019000023d10000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000023c10000c13d000023d10000013d0000001f0530018f00000a1c06300198000000400200043d0000000004620019000023d10000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000023cd0000c13d000000000005004b000023de0000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f0000000000140435000000600130021000000a1a0020009c00000a1a020080410000004002200210000000000112019f000028650001043000000a1f01100197000000000010043f0000000601000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000023f30000613d000000000101043b000000000001042d0000000001000019000028650001043000000a1f02200197000000000020043f000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000024030000613d000000000101043b000000000001042d0000000001000019000028650001043000000a2001100197000000000010043f0000000901000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000024140000613d000000000101043b000000000001042d000000000100001900002865000104300000000003010433000000000023004b0000241d0000a13d000000050220021000000000012100190000002001100039000000000001042d00000abb01000041000000000010043f0000003201000039000000040010043f00000a4e010000410000286500010430000000400100043d00000ad00010009c0000242c0000813d0000004002100039000000400020043f000000200210003900000000000204350000000000010435000000000001042d00000abb01000041000000000010043f0000004101000039000000040010043f00000a4e0100004100002865000104300000000002010019000000400100043d00000ad20010009c000024800000813d0000022003100039000000400030043f000000000302041a000000d80430027000000a1a0440019700000140051000390000000000450435000000c8043002700000ffff0440018f00000120051000390000000000450435000000b8043002700000ffff0440018f00000100051000390000000000450435000000a8043002700000ffff0440018f000000e0051000390000000000450435000000880430027000000a1a04400197000000c005100039000000000045043500000078043002700000ffff0440018f000000a0051000390000000000450435000000580430027000000a1a0440019700000080051000390000000000450435000000380430027000000a1a0440019700000060051000390000000000450435000000180430027000000a1a044001970000004005100039000000000045043500000008043002700000ffff0440018f00000020051000390000000000450435000000ff003001900000000003000039000000010300c03900000000003104350000000102200039000000000202041a000001600310003900000a1a042001970000000000430435000000380320021000000a38033001970000020004100039000000000034043500000a36002001980000000003000039000000010300c039000001e0041000390000000000340435000000800320027000000a1a03300197000001c0041000390000000000340435000000600320027000000a1a03300197000001a0041000390000000000340435000000200220027000000a200220019700000180031000390000000000230435000000000001042d00000abb01000041000000000010043f0000004101000039000000040010043f00000a4e0100004100002865000104300004000000000002000000400200043d00000ad00020009c000025cb0000813d0000004003200039000000400030043f00000020032000390000000000030435000000000002043500000a1f01100197000200000001001d000000000010043f0000000601000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000025d10000613d000000400300043d00000a240030009c000025cb0000213d000000000101043b0000004002300039000000400020043f000000000101041a00000a2702100197000400000003001d0000000002230436000000e001100270000300000001001d000000000012043500000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f0000000100200190000025d30000613d000000000101043b0003000300100074000025d40000413d00000a9b0100004100000000001004430000000001000412000000040010044300000040010000390000002400100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a9c011001c700008005020000392863285e0000040f0000000100200190000025d30000613d000000000101043b00000a1a01100197000000030010006b000025c40000413d0000000201000029000000000010043f0000000701000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000025d10000613d000000400200043d00000a240020009c0000000404000029000025cb0000213d000000000101043b0000004003200039000000400030043f000000000101041a00000a1f031001980000000003320436000000a001100270000000ff0110018f0000000000130435000025c50000613d000300000003001d000000400100043d00000a240010009c000025cb0000213d0000004003100039000000400030043f0000002003100039000000000003043500000000000104350000000002020433000000400c00043d00000a9d0100004100000000051c0436000000000100041400000a1f02200197000000040020008c000024fc0000c13d0000000103000031000000a00030008c000000a00400003900000000040340190000252b0000013d000100000005001d00000a1a00c0009c00000a1a0300004100000000030c4019000000400330021000000a1a0010009c00000a1a01008041000000c001100210000000000131019f00000a5e011001c7000200000002001d00040000000c001d2863285e0000040f000000040c000029000000600310027000000a1a03300197000000a00030008c000000a00400003900000000040340190000001f0640018f000000e00740019000000000057c0019000025190000613d000000000801034f00000000090c0019000000008a08043c0000000009a90436000000000059004b000025150000c13d000000000006004b000025260000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000025e90000613d000000020200002900000001050000290000001f01400039000001e00110018f000000000bc1001900000000001b004b0000000001000039000000010100403900000a2000b0009c000025cb0000213d0000000100100190000025cb0000c13d0000004000b0043f000000a00030008c000025d10000413d00000000010c043300000a9e0010009c000025d10000213d0000008001c00039000000000101043300000a9e0010009c000025d10000213d000000000505043300000a220050009c000025da0000813d00000a9f0100004100000000001b04350000000001000414000000040020008c000025490000c13d0000002004000039000025760000013d000200000005001d00000a1a00b0009c00000a1a0300004100000000030b4019000000400330021000000a1a0010009c00000a1a01008041000000c001100210000000000131019f00000a5e011001c700040000000b001d2863285e0000040f000000040b000029000000600310027000000a1a03300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b0019000025650000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b000025610000c13d000000000006004b000025720000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000025f50000613d00000002050000290000001f01400039000000600110018f0000000001b1001900000a200010009c000025cb0000213d000400000001001d000000400010043f000000200030008c000025d10000413d00000000010b0433000000ff0010008c000025d10000213d00000003020000290000000002020433000000ff0220018f0000000002120019000000ff0020008c000025d40000213d000000240120008c0000258d0000213d0000259c0000c13d0000000101000039000025a60000013d0000004d0010008c000025d40000213d0000000a030000390000000102000039000000010010019000000000043300a9000000010300603900000000022300a900000001011002720000000003040019000025910000c13d000000000002004b000026130000613d00000000022500d9000025ac0000013d0000000a0300003900000001010000390000002402200089000000010020019000000000043300a9000000010300603900000000011300a9000000010220027200000000030400190000259f0000c13d000000000005004b000025c70000613d00000000025100a900000000035200d9000000000013004b000025d40000c13d00000a270020009c000025e10000213d000000040300002900000a240030009c000025cb0000213d0000004001300039000000400010043f0000000001230436000300000001001d00000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f0000000100200190000025d30000613d000000000101043b00000a1a011001970000000302000029000000000012043500000004040000290000000001040019000000000001042d0000000002000019000000040300002900000a240030009c000025b10000a13d00000abb01000041000000000010043f0000004101000039000000040010043f00000a4e01000041000028650001043000000000010000190000286500010430000000000001042f00000abb01000041000000000010043f0000001101000039000000040010043f00000a4e01000041000028650001043000000aa00100004100000000001b043500000a1a00b0009c00000a1a0b0080410000004001b0021000000a5e011001c7000028650001043000000aa0010000410000000402000029000000000012043500000a1a0020009c00000a1a02008041000000400120021000000a5e011001c700002865000104300000001f0530018f00000a1c06300198000000400200043d0000000004620019000026000000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000025f00000c13d000026000000013d0000001f0530018f00000a1c06300198000000400200043d0000000004620019000026000000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000025fc0000c13d000000000005004b0000260d0000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f0000000000140435000000600130021000000a1a0020009c00000a1a020080410000004002200210000000000112019f000028650001043000000abb01000041000000000010043f0000001201000039000000040010043f00000a4e0100004100002865000104300005000000000002000000400200043d00000ad00020009c0000276a0000813d0000004003200039000000400030043f00000020032000390000000000030435000000000002043500000a1f01100197000200000001001d000000000010043f0000000601000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000027700000613d000000400300043d00000a240030009c0000276a0000213d000000000101043b0000004002300039000000400020043f000000000101041a00000a2702100197000500000003001d0000000002230436000000e001100270000400000002001d000300000001001d000000000012043500000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f0000000100200190000027720000613d000000000101043b00030003001000740000277e0000413d00000a9b0100004100000000001004430000000001000412000000040010044300000040010000390000002400100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a9c011001c700008005020000392863285e0000040f0000000100200190000027720000613d000000000101043b00000a1a01100197000000030010006b000026620000813d000000050400002900000004010000290000275f0000013d0000000201000029000000000010043f0000000701000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000027700000613d000000400200043d00000a240020009c00000005040000290000276a0000213d000000000101043b0000004003200039000000400030043f000000000101041a00000a1f031001980000000003320436000000a001100270000000ff0110018f0000000000130435000026930000613d000400000003001d000000400100043d00000a240010009c0000276a0000213d0000004003100039000000400030043f0000002003100039000000000003043500000000000104350000000002020433000000400c00043d00000a9d0100004100000000051c0436000000000100041400000a1f02200197000000040020008c000026950000c13d0000000103000031000000a00030008c000000a0040000390000000004034019000026c40000013d00000004010000290000275f0000013d000100000005001d00000a1a00c0009c00000a1a0300004100000000030c4019000000400330021000000a1a0010009c00000a1a01008041000000c001100210000000000131019f00000a5e011001c7000300000002001d00050000000c001d2863285e0000040f000000050c000029000000600310027000000a1a03300197000000a00030008c000000a00400003900000000040340190000001f0640018f000000e00740019000000000057c0019000026b20000613d000000000801034f00000000090c0019000000008a08043c0000000009a90436000000000059004b000026ae0000c13d000000000006004b000026bf0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f0000000100200190000027930000613d000000030200002900000001050000290000001f01400039000001e00110018f000000000bc1001900000000001b004b0000000001000039000000010100403900000a2000b0009c0000276a0000213d00000001001001900000276a0000c13d0000004000b0043f000000a00030008c000027700000413d00000000010c043300000a9e0010009c000027700000213d0000008001c00039000000000101043300000a9e0010009c000027700000213d000000000505043300000a220050009c000027840000813d00000a9f0100004100000000001b04350000000001000414000000040020008c000026e20000c13d00000020040000390000270f0000013d000300000005001d00000a1a00b0009c00000a1a0300004100000000030b4019000000400330021000000a1a0010009c00000a1a01008041000000c001100210000000000131019f00000a5e011001c700050000000b001d2863285e0000040f000000050b000029000000600310027000000a1a03300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b0019000026fe0000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b000026fa0000c13d000000000006004b0000270b0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f00000001002001900000279f0000613d00000003050000290000001f01400039000000600110018f0000000001b1001900000a200010009c0000276a0000213d000500000001001d000000400010043f000000200030008c000027700000413d00000000010b0433000000ff0010008c000027700000213d00000004020000290000000002020433000000ff0220018f0000000002120019000000ff0020008c0000277e0000213d000000240120008c000027260000213d000027350000c13d00000001010000390000273f0000013d0000004d0010008c0000277e0000213d0000000a030000390000000102000039000000010010019000000000043300a9000000010300603900000000022300a9000000010110027200000000030400190000272a0000c13d000000000002004b000027bd0000613d00000000022500d9000027450000013d0000000a0300003900000001010000390000002402200089000000010020019000000000043300a9000000010300603900000000011300a900000001022002720000000003040019000027380000c13d000000000005004b000027660000613d00000000025100a900000000035200d9000000000013004b0000277e0000c13d00000a270020009c0000278b0000213d000000050300002900000a240030009c0000276a0000213d0000004001300039000000400010043f0000000001230436000400000001001d00000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f0000000100200190000027720000613d000000000101043b00000a1a011001970000000402000029000000000012043500000000010200190000000504000029000000000101043300000a1a00100198000027730000613d000000000104043300000a2701100198000027730000613d000000000001042d0000000002000019000000050300002900000a240030009c0000274a0000a13d00000abb01000041000000000010043f0000004101000039000000040010043f00000a4e01000041000028650001043000000000010000190000286500010430000000000001042f000000400100043d00000aa102000041000000000021043500000004021000390000000203000029000000000032043500000a1a0010009c00000a1a01008041000000400110021000000a4e011001c7000028650001043000000abb01000041000000000010043f0000001101000039000000040010043f00000a4e01000041000028650001043000000aa00100004100000000001b043500000a1a00b0009c00000a1a0b0080410000004001b0021000000a5e011001c7000028650001043000000aa0010000410000000502000029000000000012043500000a1a0020009c00000a1a02008041000000400120021000000a5e011001c700002865000104300000001f0530018f00000a1c06300198000000400200043d0000000004620019000027aa0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b0000279a0000c13d000027aa0000013d0000001f0530018f00000a1c06300198000000400200043d0000000004620019000027aa0000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000027a60000c13d000000000005004b000027b70000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f0000000000140435000000600130021000000a1a0020009c00000a1a020080410000004002200210000000000112019f000028650001043000000abb01000041000000000010043f0000001201000039000000040010043f00000a4e01000041000028650001043000000ab3055001970000006006100039000000000056043500000ab4044001970000004005100039000000000045043500000a1f033001970000002004100039000000000034043500000a1f0220019700000000002104350000008001100039000000000001042d0004000000000002000400000002001d00000a2001100197000100000001001d000000000010043f0000000501000039000000200010043f000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a2a011001c700008010020000392863285e0000040f0000000100200190000028070000613d000000400300043d00000ad00030009c000028090000813d000000000101043b0000004002300039000000400020043f000000000101041a0000002002300039000000e004100270000000000042043500000a27011001970000000000130435000000040200002900000a1a02200198000028060000613d000300000004001d000400000002001d000200000003001d00000a95010000410000000000100443000000000100041400000a1a0010009c00000a1a01008041000000c00110021000000a96011001c70000800b020000392863285e0000040f00000001002001900000280f0000613d000000000101043b000000030110006c0000000404000029000028100000413d000000000041004b0000000202000029000028160000213d000000000102043300000a2701100197000000000001042d0000000001000019000028650001043000000abb01000041000000000010043f0000004101000039000000040010043f00000a4e010000410000286500010430000000000001042f00000abb01000041000000000010043f0000001101000039000000040010043f00000a4e010000410000286500010430000000400200043d000000440320003900000000001304350000002401200039000000000041043500000a9701000041000000000012043500000004012000390000000103000029000000000031043500000a1a0020009c00000a1a02008041000000400120021000000a61011001c70000286500010430000000000001042f00000a1a0010009c00000a1a01008041000000400110021000000a1a0020009c00000a1a020080410000006002200210000000000112019f000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000a31011001c700008010020000392863285e0000040f0000000100200190000028390000613d000000000101043b000000000001042d0000000001000019000028650001043000000000050100190000000000200443000000050030008c000028490000413d000000040100003900000000020000190000000506200210000000000664001900000005066002700000000006060031000000000161043a0000000102200039000000000031004b000028410000413d00000a1a0030009c00000a1a030080410000006001300210000000000200041400000a1a0020009c00000a1a02008041000000c002200210000000000112019f00000ad3011001c700000000020500192863285e0000040f0000000100200190000028580000613d000000000101043b000000000001042d000000000001042f0000285c002104210000000102000039000000000001042d0000000002000019000000000001042d00002861002104230000000102000039000000000001042d0000000002000019000000000001042d0000286300000432000028640001042e00002865000104300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000000000000000000000000001ffffffe000000000000000000000000000000000000000000000000000000000ffffffe0000000000000000000000000000000000000000000000000ffffffffffffff9f0000000000000000000000000000000000000000ffffffffffffffffffffffff000000000000000000000000ffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffffff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe080000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffffbf000000000000000000000000000000000000000000000000ffffffffffffff3f000000000000000000000000000000000000000000000000fffffffffffffddf00000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffdf0200000000000000000000000000000000000040000000000000000000000000bfa87805ed57dc1f0d489ce33be4c4577d74ccde357eeeee058a32c55c44a5330200000000000000000000000000000000000020000000000000000000000000c3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580bfa87805ed57dc1f0d489ce33be4c4577d74ccde357eeeee058a32c55c44a532eb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdeffe8a4859c7bd88fc0f24184464406785daae8e84cb1860cc4a4eff72e05fe2470200000000000000000000000000000000000000000000000000000000000000df1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23fe8a4859c7bd88fc0f24184464406785daae8e84cb1860cc4a4eff72e05fe2481795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91ffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000000000000ff000000000000000000000000000000000000000008a5f7f5bb38a81d8e43aca13ecd76431dbf8816ae4699affff7b00b2fc1c464ffffffff000000000000000000000000000000000000000000000000000000002812d52c0000000000000000000000000000000000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000283b699f411baff8f1c29fe49f32a828c8151596244b8e7e4c164edd6569a835525e3d4e0c31cef19cf9426af8d2c0ddd2d576359ca26bed92aac5fadda46265ff000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffff0000000000000000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000000000ffffffff0000000000000000000000000000000000000000000000000000ffff0000000000000000000000000000000000000000000000000000ffffffff0000000000000000000000000000000000000000000000000000ffff00000000000000000000000000000000000000000000000000000000ffff00000000000000000000000000000000000000000000000000000000ffff0000000000000000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000000000000000ffffffffffffff000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffff0000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000c35aa79d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000bb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97dffffffffffffffffffffffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffff0000000000000000000000000000000000000000000000000000ffff0000000000000000000000000000000000000000000000000000ffffffff000000000000000000000000000000000000000000000000ffffffff0000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000002000000000000000000000000000000000000c000000000000000000000000094967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b524ecdc020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000440000000000000000000000004de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b0000000200000000000000000000000000000100000001000000000000000000d794ef9500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000043616e6e6f7420736574206f776e657220746f207a65726f000000000000000008c379a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000000000000000000000000770e2dc300000000000000000000000000000000000000000000000000000000a69c64bf00000000000000000000000000000000000000000000000000000000d63d3af100000000000000000000000000000000000000000000000000000000f2fde38a00000000000000000000000000000000000000000000000000000000f2fde38b00000000000000000000000000000000000000000000000000000000ffdb4b3700000000000000000000000000000000000000000000000000000000d63d3af200000000000000000000000000000000000000000000000000000000d8694ccd00000000000000000000000000000000000000000000000000000000cdc73d5000000000000000000000000000000000000000000000000000000000cdc73d5100000000000000000000000000000000000000000000000000000000d02641a000000000000000000000000000000000000000000000000000000000a69c64c000000000000000000000000000000000000000000000000000000000bf78e03f00000000000000000000000000000000000000000000000000000000805f2131000000000000000000000000000000000000000000000000000000008da5cb5a000000000000000000000000000000000000000000000000000000008da5cb5b0000000000000000000000000000000000000000000000000000000091a2749a00000000000000000000000000000000000000000000000000000000805f21320000000000000000000000000000000000000000000000000000000082b49eb000000000000000000000000000000000000000000000000000000000770e2dc40000000000000000000000000000000000000000000000000000000079ba5097000000000000000000000000000000000000000000000000000000007afac32200000000000000000000000000000000000000000000000000000000407e1085000000000000000000000000000000000000000000000000000000004ab35b0a000000000000000000000000000000000000000000000000000000006cb5f3dc000000000000000000000000000000000000000000000000000000006cb5f3dd000000000000000000000000000000000000000000000000000000006def4ce7000000000000000000000000000000000000000000000000000000004ab35b0b00000000000000000000000000000000000000000000000000000000514e8cff00000000000000000000000000000000000000000000000000000000430d138b00000000000000000000000000000000000000000000000000000000430d138c0000000000000000000000000000000000000000000000000000000045ac924d00000000000000000000000000000000000000000000000000000000407e10860000000000000000000000000000000000000000000000000000000041ed29e700000000000000000000000000000000000000000000000000000000181f5a7600000000000000000000000000000000000000000000000000000000325c868d00000000000000000000000000000000000000000000000000000000325c868e000000000000000000000000000000000000000000000000000000003937306f00000000000000000000000000000000000000000000000000000000181f5a77000000000000000000000000000000000000000000000000000000002451a627000000000000000000000000000000000000000000000000000000000041e5be00000000000000000000000000000000000000000000000000000000061877e30000000000000000000000000000000000000000000000000000000006285c69000000000000000000000000000000000000004000000000000000000000000099ac52f2000000000000000000000000000000000000000000000000000000004f6e6c792063616c6c61626c65206279206f776e6572000000000000000000000000000000000000000000000000000000000064000000800000000000000000ed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127843616e6e6f74207472616e7366657220746f2073656c660000000000000000008d666f6000000000000000000000000000000000000000000000000000000000000000000000000000000000fffffffffffffffffffffffffffffffffffffc00796b89b91644bc98cd93958e4c9038275d622183e25ac5af08cc6b5d955391320200000200000000000000000000000000000004000000000000000000000000f08bcb3e0000000000000000000000000000000000000000000000000000000000000000000000000000000000ff000000000000000000000000000000000000000000000000000000000000000000000000000000000000002386f26fc1000000000000000000000000000000000000000000000000000000000000ffffffdf310ab089e4439a4c15d089f94afb7896ff553aecb10793d0ab882de59d99a32e0200000200000000000000000000000000000044000000000000000000000000feaf968c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffff313ce5670000000000000000000000000000000000000000000000000000000010cb51d10000000000000000000000000000000000000000000000000000000006439c6b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000de0b6b3a764000000000000000000000000000000000000000000000000000000000000000186a000000000000000000000000000000000000000000000000000005af3107a400000000000000000000000000000000000ffffffffffffffffffffffffffffffff181dcf100000000000000000000000000000000000000000000000000000000097a657c900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffff0000000000000000000000000000000000000020000000000000000000000000ee433e99000000000000000000000000000000000000000000000000000000004c4fc93a000000000000000000000000000000000000000000000000000000004c056b6a0000000000000000000000000000000000000000000000000000000086933789000000000000000000000000000000000000000000000000000000002502348c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000008000000000000000000175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db9000000000000000000000000000000000000000000000000ffffffffffffff7f8579befe00000000000000000000000000000000000000000000000000000000ffff000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffff5f0000000000000000000000fe00000000000000000000000000000000000000000000000000000000000000ed000000000000000000000000000000000000000052f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a191ec70600000000000000000000000000000000000000000000000000000000097e17ff000000000000000000000000000000000000000000000000000000004e487b71000000000000000000000000000000000000000000000000000000004d7573742062652070726f706f736564206f776e6572000000000000000000008be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e002000000000000000000000000000000000002200000000000000000000000005247fdce00000000000000000000000000000000000000000000000000000000000000000000002812d52c00000000000000000000000000000000000000000036f536ca000000000000000000000000000000000000000000000000000000006a92a4830000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff02000000000000000000000000000000000000a000000000000000000000000032a4ba3fa3351b11ad555d4c8ec70a744e8705607077a946807030d64b6ab1a3dd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6ed86ad9cf00000000000000000000000000000000000000000000000000000000405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace46656551756f74657220312e362e302d646576000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000060000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0000000000000000000000000000000000000000000000000ffffffffffffffa0000000000000000000000000000000000000000000000000ffffffffffffffc0000000000000000000000000000000000000000000000000ffffffffffffff40000000000000000000000000000000000000000000000000fffffffffffffde00200000200000000000000000000000000000000000000000000000000000000d854199a705ec63a2437c5f9f0eaf7eadb6f0a9abf94f7dec75ca1de173d55d7")

type GetTokenAndGasPrices struct {
	TokenPrice    *big.Int
	GasPriceValue *big.Int
}
type ProcessMessageArgs struct {
	MsgFeeJuels           *big.Int
	IsOutOfOrderExecution bool
	ConvertedExtraArgs    []byte
	DestExecDataPerToken  [][]byte
}

func (_FeeQuoter *FeeQuoter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _FeeQuoter.abi.Events["AuthorizedCallerAdded"].ID:
		return _FeeQuoter.ParseAuthorizedCallerAdded(log)
	case _FeeQuoter.abi.Events["AuthorizedCallerRemoved"].ID:
		return _FeeQuoter.ParseAuthorizedCallerRemoved(log)
	case _FeeQuoter.abi.Events["DestChainAdded"].ID:
		return _FeeQuoter.ParseDestChainAdded(log)
	case _FeeQuoter.abi.Events["DestChainConfigUpdated"].ID:
		return _FeeQuoter.ParseDestChainConfigUpdated(log)
	case _FeeQuoter.abi.Events["FeeTokenAdded"].ID:
		return _FeeQuoter.ParseFeeTokenAdded(log)
	case _FeeQuoter.abi.Events["FeeTokenRemoved"].ID:
		return _FeeQuoter.ParseFeeTokenRemoved(log)
	case _FeeQuoter.abi.Events["OwnershipTransferRequested"].ID:
		return _FeeQuoter.ParseOwnershipTransferRequested(log)
	case _FeeQuoter.abi.Events["OwnershipTransferred"].ID:
		return _FeeQuoter.ParseOwnershipTransferred(log)
	case _FeeQuoter.abi.Events["PremiumMultiplierWeiPerEthUpdated"].ID:
		return _FeeQuoter.ParsePremiumMultiplierWeiPerEthUpdated(log)
	case _FeeQuoter.abi.Events["PriceFeedPerTokenUpdated"].ID:
		return _FeeQuoter.ParsePriceFeedPerTokenUpdated(log)
	case _FeeQuoter.abi.Events["ReportPermissionSet"].ID:
		return _FeeQuoter.ParseReportPermissionSet(log)
	case _FeeQuoter.abi.Events["TokenTransferFeeConfigDeleted"].ID:
		return _FeeQuoter.ParseTokenTransferFeeConfigDeleted(log)
	case _FeeQuoter.abi.Events["TokenTransferFeeConfigUpdated"].ID:
		return _FeeQuoter.ParseTokenTransferFeeConfigUpdated(log)
	case _FeeQuoter.abi.Events["UsdPerTokenUpdated"].ID:
		return _FeeQuoter.ParseUsdPerTokenUpdated(log)
	case _FeeQuoter.abi.Events["UsdPerUnitGasUpdated"].ID:
		return _FeeQuoter.ParseUsdPerUnitGasUpdated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (FeeQuoterAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (FeeQuoterAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (FeeQuoterDestChainAdded) Topic() common.Hash {
	return common.HexToHash("0x525e3d4e0c31cef19cf9426af8d2c0ddd2d576359ca26bed92aac5fadda46265")
}

func (FeeQuoterDestChainConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x283b699f411baff8f1c29fe49f32a828c8151596244b8e7e4c164edd6569a835")
}

func (FeeQuoterFeeTokenAdded) Topic() common.Hash {
	return common.HexToHash("0xdf1b1bd32a69711488d71554706bb130b1fc63a5fa1a2cd85e8440f84065ba23")
}

func (FeeQuoterFeeTokenRemoved) Topic() common.Hash {
	return common.HexToHash("0x1795838dc8ab2ffc5f431a1729a6afa0b587f982f7b2be0b9d7187a1ef547f91")
}

func (FeeQuoterOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (FeeQuoterOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (FeeQuoterPremiumMultiplierWeiPerEthUpdated) Topic() common.Hash {
	return common.HexToHash("0xbb77da6f7210cdd16904228a9360133d1d7dfff99b1bc75f128da5b53e28f97d")
}

func (FeeQuoterPriceFeedPerTokenUpdated) Topic() common.Hash {
	return common.HexToHash("0x08a5f7f5bb38a81d8e43aca13ecd76431dbf8816ae4699affff7b00b2fc1c464")
}

func (FeeQuoterReportPermissionSet) Topic() common.Hash {
	return common.HexToHash("0x32a4ba3fa3351b11ad555d4c8ec70a744e8705607077a946807030d64b6ab1a3")
}

func (FeeQuoterTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x4de5b1bcbca6018c11303a2c3f4a4b4f22a1c741d8c4ba430d246ac06c5ddf8b")
}

func (FeeQuoterTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x94967ae9ea7729ad4f54021c1981765d2b1d954f7c92fbec340aa0a54f46b8b5")
}

func (FeeQuoterUsdPerTokenUpdated) Topic() common.Hash {
	return common.HexToHash("0x52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a")
}

func (FeeQuoterUsdPerUnitGasUpdated) Topic() common.Hash {
	return common.HexToHash("0xdd84a3fa9ef9409f550d54d6affec7e9c480c878c6ab27b78912a03e1b371c6e")
}

func (_FeeQuoter *FeeQuoter) Address() common.Address {
	return _FeeQuoter.address
}

type CustomTransaction struct {
	*types.Transaction
	CustomHash common.Hash
}

func (tx *CustomTransaction) Hash() common.Hash {
	return tx.CustomHash
}

func ConvertToTransaction(resp zktypes.TransactionResponse) *CustomTransaction {
	dtx := &types.DynamicFeeTx{
		ChainID:   resp.ChainID.ToInt(),
		Nonce:     uint64(resp.Nonce),
		GasTipCap: resp.MaxPriorityFeePerGas.ToInt(),
		GasFeeCap: resp.MaxFeePerGas.ToInt(),
		To:        &resp.To,
		Value:     resp.Value.ToInt(),
		Data:      resp.Data,
		Gas:       uint64(resp.Gas),
	}

	tx := types.NewTx(dtx)
	customTransaction := CustomTransaction{Transaction: tx, CustomHash: resp.Hash}
	return &customTransaction
}

func DeployZkSyncFeeQuoter(auth *bind.TransactOpts, backend bind.ContractBackend, params ...interface{}) (common.Address, *CustomTransaction, *FeeQuoter, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	
	zksyncClient := zkSyncClient.NewClient(client.Client())
	
	wallet := auth.Context.Value("wallet").(*zkSyncAccounts.Wallet)
	
	
	decodedBytes := common.FromHex(FeeQuoterZkBin)
	
	FeeQuoterAbi, err := FeeQuoterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	constructor, err := FeeQuoterAbi.Pack("", params...)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	hash, err := wallet.DeployWithCreate(nil, zkSyncAccounts.CreateTransaction{
		Bytecode: decodedBytes,
		Calldata: constructor,
	})
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	
	receipt, err := zksyncClient.WaitMined(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	tx, _, err := zksyncClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	

	ethTx := ConvertToTransaction(*tx)
	address := receipt.ContractAddress

	parsed, err := FeeQuoterMetaData.GetAbi()
	contractBind := bind.NewBoundContract(address, *parsed, backend, backend, backend)

	contractReturn := &FeeQuoter{address: address, abi: *parsed, FeeQuoterCaller: FeeQuoterCaller{contract: contractBind}, FeeQuoterTransactor: FeeQuoterTransactor{contract: contractBind}, FeeQuoterFilterer: FeeQuoterFilterer{contract: contractBind}}

	return address, ethTx, contractReturn, err
}

type FeeQuoterInterface interface {
	FEEBASEDECIMALS(opts *bind.CallOpts) (*big.Int, error)

	KEYSTONEPRICEDECIMALS(opts *bind.CallOpts) (*big.Int, error)

	ConvertTokenAmount(opts *bind.CallOpts, fromToken common.Address, fromTokenAmount *big.Int, toToken common.Address) (*big.Int, error)

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (FeeQuoterDestChainConfig, error)

	GetDestinationChainGasPrice(opts *bind.CallOpts, destChainSelector uint64) (InternalTimestampedPackedUint224, error)

	GetFeeTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetPremiumMultiplierWeiPerEth(opts *bind.CallOpts, token common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (FeeQuoterStaticConfig, error)

	GetTokenAndGasPrices(opts *bind.CallOpts, token common.Address, destChainSelector uint64) (GetTokenAndGasPrices,

		error)

	GetTokenPrice(opts *bind.CallOpts, token common.Address) (InternalTimestampedPackedUint224, error)

	GetTokenPriceFeedConfig(opts *bind.CallOpts, token common.Address) (FeeQuoterTokenPriceFeedConfig, error)

	GetTokenPrices(opts *bind.CallOpts, tokens []common.Address) ([]InternalTimestampedPackedUint224, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, destChainSelector uint64, token common.Address) (FeeQuoterTokenTransferFeeConfig, error)

	GetValidatedFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetValidatedTokenPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	ProcessMessageArgs(opts *bind.CallOpts, destChainSelector uint64, feeToken common.Address, feeTokenAmount *big.Int, extraArgs []byte, onRampTokenTransfers []InternalEVM2AnyTokenTransfer, sourceTokenAmounts []ClientEVMTokenAmount) (ProcessMessageArgs,

		error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []FeeQuoterDestChainConfigArgs) (*types.Transaction, error)

	ApplyFeeTokensUpdates(opts *bind.TransactOpts, feeTokensToAdd []common.Address, feeTokensToRemove []common.Address) (*types.Transaction, error)

	ApplyPremiumMultiplierWeiPerEthUpdates(opts *bind.TransactOpts, premiumMultiplierWeiPerEthArgs []FeeQuoterPremiumMultiplierWeiPerEthArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []FeeQuoterTokenTransferFeeConfigArgs, tokensToUseDefaultFeeConfigs []FeeQuoterTokenTransferFeeConfigRemoveArgs) (*types.Transaction, error)

	OnReport(opts *bind.TransactOpts, metadata []byte, report []byte) (*types.Transaction, error)

	SetReportPermissions(opts *bind.TransactOpts, permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdatePrices(opts *bind.TransactOpts, priceUpdates InternalPriceUpdates) (*types.Transaction, error)

	UpdateTokenPriceFeeds(opts *bind.TransactOpts, tokenPriceFeedUpdates []FeeQuoterTokenPriceFeedUpdate) (*types.Transaction, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*FeeQuoterAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*FeeQuoterAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*FeeQuoterAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *FeeQuoterAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*FeeQuoterAuthorizedCallerRemoved, error)

	FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*FeeQuoterDestChainAddedIterator, error)

	WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterDestChainAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainAdded(log types.Log) (*FeeQuoterDestChainAdded, error)

	FilterDestChainConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*FeeQuoterDestChainConfigUpdatedIterator, error)

	WatchDestChainConfigUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterDestChainConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigUpdated(log types.Log) (*FeeQuoterDestChainConfigUpdated, error)

	FilterFeeTokenAdded(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterFeeTokenAddedIterator, error)

	WatchFeeTokenAdded(opts *bind.WatchOpts, sink chan<- *FeeQuoterFeeTokenAdded, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenAdded(log types.Log) (*FeeQuoterFeeTokenAdded, error)

	FilterFeeTokenRemoved(opts *bind.FilterOpts, feeToken []common.Address) (*FeeQuoterFeeTokenRemovedIterator, error)

	WatchFeeTokenRemoved(opts *bind.WatchOpts, sink chan<- *FeeQuoterFeeTokenRemoved, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenRemoved(log types.Log) (*FeeQuoterFeeTokenRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FeeQuoterOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*FeeQuoterOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FeeQuoterOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeQuoterOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*FeeQuoterOwnershipTransferred, error)

	FilterPremiumMultiplierWeiPerEthUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterPremiumMultiplierWeiPerEthUpdatedIterator, error)

	WatchPremiumMultiplierWeiPerEthUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterPremiumMultiplierWeiPerEthUpdated, token []common.Address) (event.Subscription, error)

	ParsePremiumMultiplierWeiPerEthUpdated(log types.Log) (*FeeQuoterPremiumMultiplierWeiPerEthUpdated, error)

	FilterPriceFeedPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterPriceFeedPerTokenUpdatedIterator, error)

	WatchPriceFeedPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterPriceFeedPerTokenUpdated, token []common.Address) (event.Subscription, error)

	ParsePriceFeedPerTokenUpdated(log types.Log) (*FeeQuoterPriceFeedPerTokenUpdated, error)

	FilterReportPermissionSet(opts *bind.FilterOpts, reportId [][32]byte) (*FeeQuoterReportPermissionSetIterator, error)

	WatchReportPermissionSet(opts *bind.WatchOpts, sink chan<- *FeeQuoterReportPermissionSet, reportId [][32]byte) (event.Subscription, error)

	ParseReportPermissionSet(log types.Log) (*FeeQuoterReportPermissionSet, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*FeeQuoterTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *FeeQuoterTokenTransferFeeConfigDeleted, destChainSelector []uint64, token []common.Address) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*FeeQuoterTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64, token []common.Address) (*FeeQuoterTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterTokenTransferFeeConfigUpdated, destChainSelector []uint64, token []common.Address) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*FeeQuoterTokenTransferFeeConfigUpdated, error)

	FilterUsdPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*FeeQuoterUsdPerTokenUpdatedIterator, error)

	WatchUsdPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterUsdPerTokenUpdated, token []common.Address) (event.Subscription, error)

	ParseUsdPerTokenUpdated(log types.Log) (*FeeQuoterUsdPerTokenUpdated, error)

	FilterUsdPerUnitGasUpdated(opts *bind.FilterOpts, destChain []uint64) (*FeeQuoterUsdPerUnitGasUpdatedIterator, error)

	WatchUsdPerUnitGasUpdated(opts *bind.WatchOpts, sink chan<- *FeeQuoterUsdPerUnitGasUpdated, destChain []uint64) (event.Subscription, error)

	ParseUsdPerUnitGasUpdated(log types.Log) (*FeeQuoterUsdPerUnitGasUpdated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
