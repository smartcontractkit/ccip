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
	Router                          common.Address
	MaxTokensLength                 uint16
	DestGasOverhead                 uint32
	DestGasPerPayloadByte           uint16
	DestDataAvailabilityOverheadGas uint32
	DestGasPerDataAvailabilityByte  uint16
	DestDataAvailabilityMultiplier  uint16
	PriceRegistry                   common.Address
	MaxDataSize                     *big.Int
	MaxGasLimit                     uint32
}

type EVM2EVMOnRampFeeTokenConfig struct {
	NetworkFeeUSD          uint32
	MinTokenTransferFeeUSD uint32
	MaxTokenTransferFeeUSD uint32
	GasMultiplier          uint64
	PremiumMultiplier      uint64
	Enabled                bool
}

type EVM2EVMOnRampFeeTokenConfigArgs struct {
	Token                  common.Address
	NetworkFeeUSD          uint32
	MinTokenTransferFeeUSD uint32
	MaxTokenTransferFeeUSD uint32
	GasMultiplier          uint64
	PremiumMultiplier      uint64
	Enabled                bool
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
	Ratio             uint16
	DestGasOverhead   uint32
	DestBytesOverhead uint32
}

type EVM2EVMOnRampTokenTransferFeeConfigArgs struct {
	Token             common.Address
	Ratio             uint16
	DestGasOverhead   uint32
	DestBytesOverhead uint32
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

type InternalPoolUpdate struct {
	Token common.Address
	Pool  common.Address
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplier\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"maxDataSize\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"tokensAndPools\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSD\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"minTokenTransferFeeUSD\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferFeeUSD\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplier\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"ratio\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadARMSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"}],\"name\":\"InvalidNopAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkBalanceNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeBalanceReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeesToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNopsToPay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NotAFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrAdminOrNop\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyNops\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplier\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"maxDataSize\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSD\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"minTokenTransferFeeUSD\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferFeeUSD\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplier\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeConfig\",\"type\":\"tuple[]\"}],\"name\":\"FeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NopPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nopWeightsTotal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"NopsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"ratio\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"transferFeeConfig\",\"type\":\"tuple[]\"}],\"name\":\"TokenTransferFeeConfigSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"}],\"name\":\"applyPoolUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplier\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"maxDataSize\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeeTokenConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"networkFeeUSD\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"minTokenTransferFeeUSD\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferFeeUSD\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplier\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfig\",\"name\":\"feeTokenConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNopFeesJuels\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNops\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"weightsTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"prevOnRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenTransferFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ratio\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOnRamp.TokenTransferFeeConfig\",\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkAvailableForPayment\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"destDataAvailabilityMultiplier\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"maxDataSize\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkFeeUSD\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"minTokenTransferFeeUSD\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferFeeUSD\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"gasMultiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"premiumMultiplier\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"setNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"ratio\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destBytesOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOnRamp.TokenTransferFeeConfigArgs[]\",\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"setTokenTransferFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawNonLinkFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b50604051620083ce380380620083ce833981016040819052620000359162001e2f565b8333806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c08162000323565b50506040805160a081018252602084810180516001600160801b039081168085524263ffffffff169385018490528751151585870181905292518216606086018190529790950151166080909301839052600380546001600160a01b031916909417600160801b9283021760ff60a01b1916600160a01b90910217909255029091176004555086516001600160a01b0316158062000169575060208701516001600160401b0316155b8062000180575060408701516001600160401b0316155b8062000197575060608701516001600160401b0316155b80620001ae575060c08701516001600160a01b0316155b15620001cd576040516306b7c75960e31b815260040160405180910390fd5b6020808801516040808a015181517fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a948101949094526001600160401b039283169184019190915216606082015230608082015260a00160408051601f198184030181529181528151602092830120608090815289516001600160a01b0390811660e052928a01516001600160401b0390811661010052918a015182166101205260608a015190911660a0908152908901516001600160601b031660c0908152908901518216610140528801511661016052620002aa86620003ce565b620002b58362000643565b620002c082620007de565b620002cb81620008e3565b6040805160008082526020820190925262000316916200030e565b6040805180820190915260008082526020820152815260200190600190039081620002e65790505b508662000b0d565b5050505050505062002419565b336001600160a01b038216036200037d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60e08101516001600160a01b0316620003fa576040516306b7c75960e31b815260040160405180910390fd5b80600560008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160000160146101000a81548161ffff021916908361ffff16021790555060408201518160000160166101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001601a6101000a81548161ffff021916908361ffff160217905550608082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160010160006101000a81548161ffff021916908361ffff16021790555060c08201518160010160026101000a81548161ffff021916908361ffff16021790555060e08201518160010160046101000a8154816001600160a01b0302191690836001600160a01b031602179055506101008201518160010160186101000a81548162ffffff021916908362ffffff16021790555061012082015181600101601b6101000a81548163ffffffff021916908363ffffffff1602179055509050507fbbe9920f263899c6c1294310cfb58d31969efcaca2afa033aeb4d6915f25f6fa6040518060e0016040528060e0516001600160a01b03168152602001610100516001600160401b03168152602001610120516001600160401b0316815260200160a0516001600160401b0316815260200160c0516001600160601b03168152602001610140516001600160a01b03168152602001610160516001600160a01b0316815250826040516200063892919062002003565b60405180910390a150565b60005b8151811015620007ac57600082828151811062000667576200066762002083565b6020908102919091018101516040805160c080820183528385015163ffffffff908116835283850151811683870190815260608087015183168587019081526080808901516001600160401b0390811693880193845260a0808b01518216928901928352968a0151151596880196875298516001600160a01b03166000908152600d909a529690982094518554925198519151965194511515600160e01b0260ff60e01b19958916600160a01b0295909516600160a01b600160e81b0319979098166c0100000000000000000000000002600160601b600160a01b0319928516680100000000000000000292909216600160401b600160a01b0319998516640100000000026001600160401b0319909416919094161791909117969096161794909417919091169190911791909117905550620007a481620020af565b905062000646565b507f2386f61ab5cafc3fed44f9f614f721ab53479ef64067fd16c1a2491b63ddf1a881604051620006389190620020cb565b60005b8151811015620008b157600082828151811062000802576200080262002083565b60209081029190910181015160408051606080820183528385015161ffff90811683528385015163ffffffff90811684880190815292860151811684860190815295516001600160a01b03166000908152600e909752939095209151825491519451841666010000000000000263ffffffff60301b1995909416620100000265ffffffffffff19909216951694909417939093179190911617905550620008a981620020af565b9050620007e1565b507f3df02211160f78e9bbed13e8bf0752035362a05f0651a5e25007acfb659c36208160405162000638919062002183565b805160408111156200090857604051635ad0867d60e11b815260040160405180910390fd5b6010546c01000000000000000000000000900463ffffffff161580159062000952575060105463ffffffff6c010000000000000000000000008204166001600160601b0390911610155b1562000962576200096262000e10565b600062000970600762001010565b90505b8015620009bc576000620009966200098d600184620021f2565b60079062001023565b509050620009a660078262001041565b505080620009b49062002208565b905062000973565b506000805b8281101562000aa4576000848281518110620009e157620009e162002083565b6020026020010151600001519050600085838151811062000a065762000a0662002083565b602002602001015160200151905060e0516001600160a01b0316826001600160a01b0316148062000a3e57506001600160a01b038216155b1562000a6957604051634de938d160e01b81526001600160a01b038316600482015260240162000084565b62000a7b60078361ffff84166200105f565b5062000a8c61ffff82168562002222565b935050508062000a9c90620020af565b9050620009c1565b506010805463ffffffff60601b19166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd249062000b00908390869062002242565b60405180910390a1505050565b60005b825181101562000c4957600083828151811062000b315762000b3162002083565b6020026020010151600001519050600084838151811062000b565762000b5662002083565b6020908102919091018101510151905062000b73600a836200107f565b62000b9d576040516373913ebd60e01b81526001600160a01b038316600482015260240162000084565b6001600160a01b03811662000bb4600a8462001096565b6001600160a01b03161462000bdc57604051630d98f73360e31b815260040160405180910390fd5b62000be9600a83620010ad565b1562000c3357604080516001600160a01b038085168252831660208201527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15b50508062000c4190620020af565b905062000b10565b5060005b815181101562000e0b57600082828151811062000c6e5762000c6e62002083565b6020026020010151600001519050600083838151811062000c935762000c9362002083565b602002602001015160200151905060006001600160a01b0316826001600160a01b0316148062000cca57506001600160a01b038116155b1562000ce85760405162d8548360e71b815260040160405180910390fd5b806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000d27573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000d4d9190620022af565b6001600160a01b0316826001600160a01b03161462000d7f57604051630d98f73360e31b815260040160405180910390fd5b62000d8d600a8383620010c4565b1562000ddc57604080516001600160a01b038085168252831660208201527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a162000df5565b604051633caf458560e01b815260040160405180910390fd5b50508062000e0390620020af565b905062000c4d565b505050565b6000546001600160a01b0316331480159062000e3757506002546001600160a01b03163314155b801562000e4e575062000e4c600733620010dc565b155b1562000e6d5760405163032bb72b60e31b815260040160405180910390fd5b6010546c01000000000000000000000000900463ffffffff16600081900362000ea95760405163990e30bf60e01b815260040160405180910390fd5b6010546001600160601b03168181101562000ed7576040516311a1ee3b60e31b815260040160405180910390fd5b600062000ee3620010f3565b121562000f0357604051631e9acf1760e31b815260040160405180910390fd5b80600062000f12600762001010565b905060005b8181101562000fea5760008062000f3060078462001023565b909250905060008762000f4d836001600160601b038a16620022cf565b62000f599190620022e9565b905062000f6781876200230c565b60e05190965062000f8c906001600160a01b0316846001600160601b03841662001181565b6040516001600160601b03821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a25050508062000fe290620020af565b905062000f17565b5050601080546001600160601b0319166001600160601b03929092169190911790555050565b60006200101d82620011d9565b92915050565b6000808080620010348686620011e6565b9097909650945050505050565b600062001058836001600160a01b03841662001213565b9392505050565b600062001077846001600160a01b0385168462001232565b949350505050565b600062001058836001600160a01b03841662001251565b600062001058836001600160a01b0384166200125f565b600062001058836001600160a01b0384166200126d565b600062001077846001600160a01b038516846200127b565b600062001058836001600160a01b03841662001293565b60105460e0516040516370a0823160e01b81523060048201526000926001600160601b0316916001600160a01b0316906370a0823190602401602060405180830381865afa1580156200114a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200117091906200232f565b6200117c919062002349565b905090565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b0390811663a9059cbb60e01b1790915262000e0b918591620012a116565b60006200101d8262001372565b60008080620011f685856200137d565b600081815260029690960160205260409095205494959350505050565b600081815260028301602052604081208190556200105883836200138b565b6000828152600284016020526040812082905562001077848462001399565b600062001058838362001293565b6000620010588383620013a7565b600062001058838362001213565b60006200107784846001600160a01b03851662001232565b60006200105883836200141c565b6040805180820190915260208082527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656490820152600090620012f0906001600160a01b03851690849062001435565b80519091501562000e0b57808060200190518101906200131191906200236c565b62000e0b5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840162000084565b60006200101d825490565b600062001058838362001446565b600062001058838362001473565b60006200105883836200157e565b600081815260028301602052604081205480151580620013ce5750620013ce848462001293565b620010585760405162461bcd60e51b815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b65790000604482015260640162000084565b6000818152600183016020526040812054151562001058565b6060620010778484600085620015d0565b600082600001828154811062001460576200146062002083565b9060005260206000200154905092915050565b600081815260018301602052604081205480156200156c5760006200149a600183620021f2565b8554909150600090620014b090600190620021f2565b90508181146200151c576000866000018281548110620014d457620014d462002083565b9060005260206000200154905080876000018481548110620014fa57620014fa62002083565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806200153057620015306200238a565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506200101d565b60009150506200101d565b5092915050565b6000818152600183016020526040812054620015c7575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556200101d565b5060006200101d565b606082471015620016335760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840162000084565b600080866001600160a01b03168587604051620016519190620023c6565b60006040518083038185875af1925050503d806000811462001690576040519150601f19603f3d011682016040523d82523d6000602084013e62001695565b606091505b509092509050620016a987838387620016b4565b979650505050505050565b606083156200172857825160000362001720576001600160a01b0385163b620017205760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640162000084565b508162001077565b6200107783838151156200173f5781518083602001fd5b8060405162461bcd60e51b8152600401620000849190620023e4565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b03811182821017156200179657620017966200175b565b60405290565b60405161014081016001600160401b03811182821017156200179657620017966200175b565b604080519081016001600160401b03811182821017156200179657620017966200175b565b604051608081016001600160401b03811182821017156200179657620017966200175b565b604051601f8201601f191681016001600160401b03811182821017156200183757620018376200175b565b604052919050565b6001600160a01b03811681146200185557600080fd5b50565b805162001865816200183f565b919050565b80516001600160401b03811681146200186557600080fd5b600060e082840312156200189557600080fd5b6200189f62001771565b90508151620018ae816200183f565b8152620018be602083016200186a565b6020820152620018d1604083016200186a565b6040820152620018e4606083016200186a565b606082015260808201516001600160601b03811681146200190457600080fd5b60808201526200191760a0830162001858565b60a08201526200192a60c0830162001858565b60c082015292915050565b805161ffff811681146200186557600080fd5b805163ffffffff811681146200186557600080fd5b805162ffffff811681146200186557600080fd5b600061014082840312156200198557600080fd5b6200198f6200179c565b90506200199c8262001858565b8152620019ac6020830162001935565b6020820152620019bf6040830162001948565b6040820152620019d26060830162001935565b6060820152620019e56080830162001948565b6080820152620019f860a0830162001935565b60a082015262001a0b60c0830162001935565b60c082015262001a1e60e0830162001858565b60e082015261010062001a338184016200195d565b9082015261012062001a4783820162001948565b9082015292915050565b60006001600160401b0382111562001a6d5762001a6d6200175b565b5060051b60200190565b600082601f83011262001a8957600080fd5b8151602062001aa262001a9c8362001a51565b6200180c565b82815260069290921b8401810191818101908684111562001ac257600080fd5b8286015b8481101562001b1c576040818903121562001ae15760008081fd5b62001aeb620017c2565b815162001af8816200183f565b81528185015162001b09816200183f565b8186015283529183019160400162001ac6565b509695505050505050565b805180151581146200186557600080fd5b80516001600160801b03811681146200186557600080fd5b60006060828403121562001b6357600080fd5b604051606081016001600160401b038111828210171562001b885762001b886200175b565b60405290508062001b998362001b27565b815262001ba96020840162001b38565b602082015262001bbc6040840162001b38565b60408201525092915050565b600082601f83011262001bda57600080fd5b8151602062001bed62001a9c8362001a51565b82815260e0928302850182019282820191908785111562001c0d57600080fd5b8387015b8581101562001cc25781818a03121562001c2b5760008081fd5b62001c3562001771565b815162001c42816200183f565b815262001c5182870162001948565b86820152604062001c6481840162001948565b90820152606062001c7783820162001948565b90820152608062001c8a8382016200186a565b9082015260a062001c9d8382016200186a565b9082015260c062001cb083820162001b27565b90820152845292840192810162001c11565b5090979650505050505050565b600082601f83011262001ce157600080fd5b8151602062001cf462001a9c8362001a51565b82815260079290921b8401810191818101908684111562001d1457600080fd5b8286015b8481101562001b1c576080818903121562001d335760008081fd5b62001d3d620017e7565b815162001d4a816200183f565b815262001d5982860162001935565b85820152604062001d6c81840162001948565b90820152606062001d7f83820162001948565b9082015283529183019160800162001d18565b600082601f83011262001da457600080fd5b8151602062001db762001a9c8362001a51565b82815260069290921b8401810191818101908684111562001dd757600080fd5b8286015b8481101562001b1c576040818903121562001df65760008081fd5b62001e00620017c2565b815162001e0d816200183f565b815262001e1c82860162001935565b8186015283529183019160400162001ddb565b6000806000806000806000610300888a03121562001e4c57600080fd5b62001e58898962001882565b965062001e698960e08a0162001971565b6102208901519096506001600160401b038082111562001e8857600080fd5b62001e968b838c0162001a77565b965062001ea88b6102408c0162001b50565b95506102a08a015191508082111562001ec057600080fd5b62001ece8b838c0162001bc8565b94506102c08a015191508082111562001ee657600080fd5b62001ef48b838c0162001ccf565b93506102e08a015191508082111562001f0c57600080fd5b5062001f1b8a828b0162001d92565b91505092959891949750929550565b80516001600160a01b03168252602081015162001f4d602084018261ffff169052565b50604081015162001f66604084018263ffffffff169052565b50606081015162001f7d606084018261ffff169052565b50608081015162001f96608084018263ffffffff169052565b5060a081015162001fad60a084018261ffff169052565b5060c081015162001fc460c084018261ffff169052565b5060e081015162001fe060e08401826001600160a01b03169052565b506101008181015162ffffff16908301526101209081015163ffffffff16910152565b82516001600160a01b0390811682526020808501516001600160401b0390811691840191909152604080860151821690840152606080860151909116908301526080808501516001600160601b03169083015260a08085015182169083015260c0808501519091169082015261022081016200105860e083018462001f2a565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201620020c457620020c462002099565b5060010190565b602080825282518282018190526000919060409081850190868401855b828110156200217657815180516001600160a01b031685528681015163ffffffff9081168887015286820151811687870152606080830151909116908601526080808201516001600160401b03169086015260a08082015162002155878301826001600160401b03169052565b505060c09081015115159085015260e09093019290850190600101620020e8565b5091979650505050505050565b602080825282518282018190526000919060409081850190868401855b828110156200217657815180516001600160a01b031685528681015161ffff16878601528581015163ffffffff90811687870152606091820151169085015260809093019290850190600101620021a0565b818103818111156200101d576200101d62002099565b6000816200221a576200221a62002099565b506000190190565b63ffffffff81811683821601908082111562001577576200157762002099565b6000604080830163ffffffff8616845260208281860152818651808452606087019150828801935060005b81811015620022a157845180516001600160a01b0316845284015161ffff168484015293830193918501916001016200226d565b509098975050505050505050565b600060208284031215620022c257600080fd5b815162001058816200183f565b80820281158282048414176200101d576200101d62002099565b6000826200230757634e487b7160e01b600052601260045260246000fd5b500490565b6001600160601b0382811682821603908082111562001577576200157762002099565b6000602082840312156200234257600080fd5b5051919050565b818103600083128015838313168383128216171562001577576200157762002099565b6000602082840312156200237f57600080fd5b620010588262001b27565b634e487b7160e01b600052603160045260246000fd5b60005b83811015620023bd578181015183820152602001620023a3565b50506000910152565b60008251620023da818460208701620023a0565b9190910192915050565b602081526000825180602084015262002405816040850160208701620023a0565b601f01601f19169190910160400192915050565b60805160a05160c05160e05161010051610120516101405161016051615eae62002520600039600081816103340152818161153b0152613e1a0152600081816103050152818161145d015281816114c501528181611a1601528181611a7e0152613deb01526000818161027101528181610be001528181611e300152613d5701526000818161024101528181611b490152613d2701526000818161021201528181610faf015281816117c1015281816118c201528181612360015281816130e9015281816134760152613cf80152600081816102d10152818161198e0152613db70152600081816102a1015281816126a70152613d8701526000611ef80152615eae6000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c806375294130116100f9578063b06d41bc11610097578063d09dc33911610071578063d09dc33914610955578063d3c7c2c71461095d578063eff7cc4814610972578063f2fde38b1461097a57600080fd5b8063b06d41bc14610919578063c92b28321461092f578063cf076f601461094257600080fd5b8063856c8247116100d3578063856c8247146107755780638da5cb5b146107885780639a113c3614610799578063a7d3e02f1461090657600080fd5b8063752941301461074757806376f6ae761461075a57806379ba50971461076d57600080fd5b8063546719cd11610166578063599f643111610140578063599f64311461056f5780635d86f14114610594578063704b6c02146105a75780637437ff9f146105ba57600080fd5b8063546719cd146104d8578063549e946f1461053c57806354b714681461054f57600080fd5b8063352e4bc8116101a2578063352e4bc81461046e57806338724a95146104835780633a87ac53146104a45780634120fccd146104b757600080fd5b806306285c69146101c95780631772047e1461037a578063181f5a7714610425575b600080fd5b6103646040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526040518060e001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b604051610371919061495d565b60405180910390f35b6103f16103883660046149fe565b6040805160608082018352600080835260208084018290529284018190526001600160a01b03949094168452600e825292829020825193840183525461ffff8116845263ffffffff62010000820481169285019290925266010000000000009004169082015290565b60408051825161ffff16815260208084015163ffffffff908116918301919091529282015190921690820152606001610371565b6104616040518060400160405280601381526020017f45564d3245564d4f6e52616d7020312e322e300000000000000000000000000081525081565b6040516103719190614a89565b61048161047c366004614c28565b61098d565b005b610496610491366004614d5e565b6109f6565b604051908152602001610371565b6104816104b2366004614e2f565b610e56565b6104bf610e6c565b60405167ffffffffffffffff9091168152602001610371565b6104e0610ea0565b604051610371919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b61048161054a366004614e93565b610f50565b6010546040516bffffffffffffffffffffffff9091168152602001610371565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610371565b61057c6105a23660046149fe565b611105565b6104816105b53660046149fe565b611164565b61073a6040805161014081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101919091525060408051610140810182526005546001600160a01b03808216835261ffff7401000000000000000000000000000000000000000083048116602085015263ffffffff76010000000000000000000000000000000000000000000084048116958501959095527a0100000000000000000000000000000000000000000000000000008304811660608501527c01000000000000000000000000000000000000000000000000000000009092048416608084015260065480831660a085015262010000810490921660c084015264010000000082041660e083015262ffffff78010000000000000000000000000000000000000000000000008204166101008301527b01000000000000000000000000000000000000000000000000000000900490911661012082015290565b6040516103719190614fa3565b610481610755366004614fc4565b61122e565b6104816107683660046150b1565b611294565b61048161134c565b6104bf6107833660046149fe565b61142f565b6000546001600160a01b031661057c565b61089c6107a73660046149fe565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506001600160a01b03166000908152600d6020908152604091829020825160c081018452905463ffffffff808216835264010000000082048116938301939093526801000000000000000081049092169281019290925267ffffffffffffffff6c0100000000000000000000000082048116606084015274010000000000000000000000000000000000000000820416608083015260ff7c010000000000000000000000000000000000000000000000000000000090910416151560a082015290565b6040516103719190600060c08201905063ffffffff80845116835280602085015116602084015280604085015116604084015250606083015167ffffffffffffffff8082166060850152806080860151166080850152505060a0830151151560a083015292915050565b610496610914366004615126565b611537565b610921611f69565b6040516103719291906151d4565b61048161093d366004615216565b61206d565b610481610950366004615297565b6120d5565b6104966120e6565b6109656120f0565b604051610371919061535f565b6104816121a1565b6104816109883660046149fe565b612438565b6000546001600160a01b031633148015906109b357506002546001600160a01b03163314155b156109ea576040517ffbdb8e5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109f381612449565b50565b600080610a0e610a0960808501856153ac565b612678565b9050610a39610a2060208501856153ac565b8351909150610a326040870187615411565b905061276c565b6000600d81610a4e60808701606088016149fe565b6001600160a01b031681526020808201929092526040908101600020815160c081018352905463ffffffff80821683526401000000008204811694830194909452680100000000000000008104909316918101919091526c01000000000000000000000000820467ffffffffffffffff90811660608301527401000000000000000000000000000000000000000083041660808201527c010000000000000000000000000000000000000000000000000000000090910460ff16151560a08201819052909150610b6b57610b2860808501606086016149fe565b6040517fa7499d200000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201526024015b60405180910390fd5b600654600090819064010000000090046001600160a01b031663ffdb4b37610b996080890160608a016149fe565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660248201526044016040805180830381865afa158015610c24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c4891906154a5565b9092509050806000808080610c6060408c018c615411565b90501115610c9b57610c8f610c7b60808c0160608d016149fe565b87610c8960408e018e615411565b8b612893565b91945092509050610cb7565b8651610cb49063ffffffff16662386f26fc10000615507565b92505b6080870151610cd09067ffffffffffffffff1684615507565b606088015160055491945060009167ffffffffffffffff9091169063ffffffff8516907a010000000000000000000000000000000000000000000000000000900461ffff16610d2260208f018f6153ac565b610d2d929150615507565b6005548c51610d5c91760100000000000000000000000000000000000000000000900463ffffffff169061551e565b610d66919061551e565b610d70919061551e565b610d7a9190615507565b610d94906dffffffffffffffffffffffffffff8716615507565b60065490915060009062010000900461ffff1615610e08576000607060ff16887bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16901c9050610e04818e8060200190610dea91906153ac565b90508f8060400190610dfc9190615411565b905087612b4c565b9150505b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff881681610e32848861551e565b610e3c919061551e565b610e469190615531565b9c9b505050505050505050505050565b610e5e612c1e565b610e688282612c94565b5050565b601054600090610e9b90700100000000000000000000000000000000900467ffffffffffffffff16600161556c565b905090565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526003546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff1660208501527401000000000000000000000000000000000000000090920460ff161515938301939093526004548084166060840152049091166080820152610e9b90612ff4565b6000546001600160a01b03163314801590610f7657506002546001600160a01b03163314155b15610fad576040517ffbdb8e5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b03161480610ff457506001600160a01b038116155b1561102b576040517f232cb97f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006110356130a6565b121561106d576040517f02075e0000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152610e689082906001600160a01b038516906370a0823190602401602060405180830381865afa1580156110d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110f4919061558d565b6001600160a01b0385169190613166565b6000611112600a836131e6565b611153576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610b62565b61115e600a836131fb565b92915050565b6000546001600160a01b0316331480159061118a57506002546001600160a01b03163314155b156111c1576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c906020015b60405180910390a150565b6000546001600160a01b0316331480159061125457506002546001600160a01b03163314155b1561128b576040517ffbdb8e5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109f381613210565b6000546001600160a01b031633148015906112ba57506002546001600160a01b03163314155b156112f1576040517ffbdb8e5600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e688282808060200260200160405190810160405280939291908181526020016000905b8282101561134257611333604083028601368190038101906155a6565b81526020019060010190611316565b505050505061333d565b6001546001600160a01b031633146113c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610b62565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6001600160a01b0381166000908152600f602052604081205467ffffffffffffffff168015801561148857507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615155b1561115e576040517f856c82470000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063856c824790602401602060405180830381865afa15801561150c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061153091906155e5565b9392505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663397796f76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611597573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115bb9190615602565b156115f2576040517fc148371500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038216611632576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005546001600160a01b03163314611676576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61168084806153ac565b90506020146116c75761169384806153ac565b6040517f370d875f000000000000000000000000000000000000000000000000000000008152600401610b62929190615668565b60006116d385806153ac565b8101906116e0919061567c565b90506001600160a01b038111806116f75750600a81105b156117065761169385806153ac565b6000611718610a0960808801886153ac565b51905061173b61172b60208801886153ac565b905082610a3260408a018a615411565b6117b761174b6040880188615411565b808060200260200160405190810160405280939291908181526020016000905b828210156117975761178860408302860136819003810190615695565b8152602001906001019061176b565b505060065464010000000090046001600160a01b031692506135b0915050565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166117f160808801606089016149fe565b6001600160a01b03160361185557601080548691906000906118229084906bffffffffffffffffffffffff166156cf565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff16021790555061197c565b60065464010000000090046001600160a01b03166241e5be61187d6080890160608a016149fe565b60405160e083901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b039182166004820152602481018990527f00000000000000000000000000000000000000000000000000000000000000009091166044820152606401602060405180830381865afa158015611909573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061192d919061558d565b6010805460009061194d9084906bffffffffffffffffffffffff166156cf565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055505b6010546bffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116911611156119e9576040517fe5c7a49100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0384166000908152600f602052604090205467ffffffffffffffff16158015611a4157507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615155b15611b39576040517f856c82470000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063856c824790602401602060405180830381865afa158015611ac5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ae991906155e5565b6001600160a01b0385166000908152600f6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff929092169190911790555b6000604051806101a001604052807f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168152602001866001600160a01b03168152602001846001600160a01b0316815260200160108081819054906101000a900467ffffffffffffffff16611bb6906156f4565b825467ffffffffffffffff9182166101009390930a838102908302199091161790925582526020808301869052600060408085018290526001600160a01b038b168252600f90925290812080546060909401939092611c1591166156f4565b825467ffffffffffffffff9182166101009390930a83810292021916179091558152602001611c4a60808a0160608b016149fe565b6001600160a01b03168152602001878152602001888060200190611c6e91906153ac565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250602001611cb560408a018a615411565b808060200260200160405190810160405280939291908181526020016000905b82821015611d0157611cf260408302860136819003810190615695565b81526020019060010190611cd5565b5050509183525050602001611d1960408a018a615411565b905067ffffffffffffffff811115611d3357611d33614a9c565b604051908082528060200260200182016040528015611d6657816020015b6060815260200190600190039081611d515790505b508152600060209091018190529091505b611d846040890189615411565b9050811015611ef1576000611d9c60408a018a615411565b83818110611dac57611dac61571b565b905060400201803603810190611dc29190615695565b9050611dd18160000151611105565b6001600160a01b0316639687544588611dea8c806153ac565b60208087015160408051928301815260008352517fffffffff0000000000000000000000000000000000000000000000000000000060e088901b168152611e58959493927f00000000000000000000000000000000000000000000000000000000000000009160040161574a565b6000604051808303816000875af1158015611e77573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611ebd91908101906157a2565b8361016001518381518110611ed457611ed461571b565b60200260200101819052505080611eea90615854565b9050611d77565b50611f1c817f000000000000000000000000000000000000000000000000000000000000000061376b565b6101808201526040517fd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd90611f52908390615925565b60405180910390a161018001519695505050505050565b6060600080611f7860076138fd565b90508067ffffffffffffffff811115611f9357611f93614a9c565b604051908082528060200260200182016040528015611fd857816020015b6040805180820190915260008082526020820152815260200190600190039081611fb15790505b50925060005b8181101561204a57600080611ff4600784613908565b915091506040518060400160405280836001600160a01b031681526020018261ffff1681525086848151811061202c5761202c61571b565b602002602001018190525050508061204390615854565b9050611fde565b505060105491926c0100000000000000000000000090920463ffffffff16919050565b6000546001600160a01b0316331480159061209357506002546001600160a01b03163314155b156120ca576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109f3600382613926565b6120dd612c1e565b6109f381613afe565b6000610e9b6130a6565b606060006120fe600a613e54565b67ffffffffffffffff81111561211657612116614a9c565b60405190808252806020026020018201604052801561213f578160200160208202803683370190505b50905060005b815181101561219b57612159600a82613e5f565b5082828151811061216c5761216c61571b565b60200260200101816001600160a01b03166001600160a01b0316815250508061219490615854565b9050612145565b50919050565b6000546001600160a01b031633148015906121c757506002546001600160a01b03163314155b80156121db57506121d9600733613e6e565b155b15612212576040517f195db95800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6010546c01000000000000000000000000900463ffffffff166000819003612266576040517f990e30bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6010546bffffffffffffffffffffffff16818110156122b1576040517f8d0f71d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006122bb6130a6565b12156122f3576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600061230060076138fd565b905060005b818110156123f55760008061231b600784613908565b909250905060008761233b836bffffffffffffffffffffffff8a16615507565b6123459190615531565b90506123518187615a78565b95506123956001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016846bffffffffffffffffffffffff8416613166565b6040516bffffffffffffffffffffffff821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a2505050806123ee90615854565b9050612305565b5050601080547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff929092169190911790555050565b612440612c1e565b6109f381613e83565b60005b81518110156126485760008282815181106124695761246961571b565b6020908102919091018101516040805160c080820183528385015163ffffffff9081168352838501518116838701908152606080870151831685870190815260808089015167ffffffffffffffff90811693880193845260a0808b01518216928901928352968a0151151596880196875298516001600160a01b03166000908152600d909a5296909820945185549251985191519651945115157c0100000000000000000000000000000000000000000000000000000000027fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff9589167401000000000000000000000000000000000000000002959095167fffffff000000000000000000ffffffffffffffffffffffffffffffffffffffff979098166c01000000000000000000000000027fffffffffffffffffffffffff0000000000000000ffffffffffffffffffffffff9285166801000000000000000002929092167fffffffffffffffffffffffff000000000000000000000000ffffffffffffffff998516640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000090941691909416179190911796909616179490941791909116919091179190911790555061264181615854565b905061244c565b507f2386f61ab5cafc3fed44f9f614f721ab53479ef64067fd16c1a2491b63ddf1a8816040516112239190615a9d565b60408051602081019091526000815260008290036126ce5750604080516020810190915267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261115e565b7f97a657c9000000000000000000000000000000000000000000000000000000006126f98385615b47565b7fffffffff000000000000000000000000000000000000000000000000000000001614612752576040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61275f8260048186615b8f565b8101906115309190615bb9565b6006547801000000000000000000000000000000000000000000000000900462ffffff16808411156127d4576040517f869337890000000000000000000000000000000000000000000000000000000081526004810182905260248101859052604401610b62565b6006547b01000000000000000000000000000000000000000000000000000000900463ffffffff16831115612835576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60055474010000000000000000000000000000000000000000900461ffff1682111561288d576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050565b6000808084815b81811015612ad75760008888838181106128b6576128b661571b565b9050604002018036038101906128cc9190615695565b80516001600160a01b03166000908152600e60209081526040918290208251606081018452905461ffff8116825263ffffffff620100008204811693830193909352660100000000000090048216928101929092528251929350909161293691600a91906131e616565b61297a5781516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b62565b805160009061ffff1615612a935760008d6001600160a01b031684600001516001600160a01b031614612a3e5760065484516040517f4ab35b0b0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526401000000009092041690634ab35b0b90602401602060405180830381865afa158015612a13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a379190615bfb565b9050612a41565b508b5b82516020850151620186a09161ffff1690612a7b907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff851690613f5e565b612a859190615507565b612a8f9190615531565b9150505b612a9d818961551e565b9750816020015187612aaf9190615c16565b9650816040015186612ac19190615c16565b955050505080612ad090615854565b905061289a565b506000856020015163ffffffff16662386f26fc10000612af79190615507565b905080851015612b0a579350612b419050565b6000866040015163ffffffff16662386f26fc10000612b299190615507565b905080861115612b3d579450612b41915050565b5050505b955095509592505050565b60008063ffffffff8316612b61608086615507565b612b6d8761022061551e565b612b77919061551e565b612b81919061551e565b6005546006549192506000917c010000000000000000000000000000000000000000000000000000000090910463ffffffff1690612bc39061ffff1684615507565b612bcd919061551e565b60065490915062010000900461ffff16612bf76dffffffffffffffffffffffffffff891683615507565b612c019190615507565b612c1190655af3107a4000615507565b925050505b949350505050565b6000546001600160a01b03163314612c92576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610b62565b565b60005b8251811015612df5576000838281518110612cb457612cb461571b565b60200260200101516000015190506000848381518110612cd657612cd661571b565b6020026020010151602001519050612cf882600a6131e690919063ffffffff16565b612d39576040517f73913ebd0000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610b62565b6001600160a01b038116612d4e600a846131fb565b6001600160a01b031614612d8e576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612d99600a83613f9b565b15612de257604080516001600160a01b038085168252831660208201527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15b505080612dee90615854565b9050612c97565b5060005b8151811015612fef576000828281518110612e1657612e1661571b565b60200260200101516000015190506000838381518110612e3857612e3861571b565b602002602001015160200151905060006001600160a01b0316826001600160a01b03161480612e6e57506001600160a01b038116155b15612ea5576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ee3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f079190615c33565b6001600160a01b0316826001600160a01b031614612f51576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f5d600a8383613fb0565b15612faa57604080516001600160a01b038085168252831660208201527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1612fdc565b6040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505080612fe890615854565b9050612df9565b505050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261308282606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff16426130669190615c50565b85608001516fffffffffffffffffffffffffffffffff16613fc6565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b6010546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000916bffffffffffffffffffffffff16907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa158015613138573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061315c919061558d565b610e9b9190615c63565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052612fef908490613fee565b6000611530836001600160a01b0384166140ed565b6000611530836001600160a01b0384166140f9565b60005b815181101561330d5760008282815181106132305761323061571b565b60209081029190910181015160408051606080820183528385015161ffff90811683528385015163ffffffff90811684880190815292860151811684860190815295516001600160a01b03166000908152600e90975293909520915182549151945184166601000000000000027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff9590941662010000027fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000000090921695169490941793909317919091161790555061330681615854565b9050613213565b507f3df02211160f78e9bbed13e8bf0752035362a05f0651a5e25007acfb659c3620816040516112239190615c83565b8051604081111561337a576040517fb5a10cfa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6010546c01000000000000000000000000900463ffffffff16158015906133c8575060105463ffffffff6c010000000000000000000000008204166bffffffffffffffffffffffff90911610155b156133d5576133d56121a1565b60006133e160076138fd565b90505b80156134235760006134026133fa600184615c50565b600790613908565b509050613410600782614105565b50508061341c90615cf0565b90506133e4565b506000805b828110156135315760008482815181106134445761344461571b565b602002602001015160000151905060008583815181106134665761346661571b565b60200260200101516020015190507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b031614806134bb57506001600160a01b038216155b156134fd576040517f4de938d10000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610b62565b61350d60078361ffff841661411a565b5061351c61ffff821685615c16565b935050508061352a90615854565b9050613428565b50601080547fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd24906135a39083908690615d25565b60405180910390a1505050565b81516000805b8281101561375d576000846001600160a01b031663d02641a08784815181106135e1576135e161571b565b6020908102919091010151516040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b0390911660048201526024016040805180830381865afa158015613648573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061366c9190615d44565b5190507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81166000036136f1578582815181106136a7576136a761571b565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b62565b61373f8683815181106137065761370661571b565b602002602001015160200151827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16613f5e90919063ffffffff16565b613749908461551e565b9250508061375690615854565b90506135b6565b5061288d6003826000614130565b60008060001b8284606001518560c0015186602001518760400151886101200151805190602001208961014001516040516020016137a99190615d77565b604051602081830303815290604052805190602001208a61016001516040516020016137d59190615d8a565b604051602081830303815290604052805190602001208b608001518c60a001516040516020016138739b9a999897969594939291909a8b5260208b019990995267ffffffffffffffff97881660408b01529590961660608901526001600160a01b0393841660808901529190921660a087015260c086019190915260e085015261010084019190915261012083015215156101408201526101600190565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815282825260e08601516101008701516001600160a01b03909116602085015283830152815180840383018152606084019092526138df92909190608001615d9d565b60405160208183030381529060405280519060200120905092915050565b600061115e8261447f565b6000808080613917868661448a565b909450925050505b9250929050565b815460009061394f90700100000000000000000000000000000000900463ffffffff1642615c50565b905080156139f15760018301548354613997916fffffffffffffffffffffffffffffffff80821692811691859170010000000000000000000000000000000090910416613fc6565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354613a17916fffffffffffffffffffffffffffffffff90811691166144b5565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19906135a39084908151151581526020808301516fffffffffffffffffffffffffffffffff90811691830191909152604092830151169181019190915260600190565b60e08101516001600160a01b0316613b42576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160000160146101000a81548161ffff021916908361ffff16021790555060408201518160000160166101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001601a6101000a81548161ffff021916908361ffff160217905550608082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160010160006101000a81548161ffff021916908361ffff16021790555060c08201518160010160026101000a81548161ffff021916908361ffff16021790555060e08201518160010160046101000a8154816001600160a01b0302191690836001600160a01b031602179055506101008201518160010160186101000a81548162ffffff021916908362ffffff16021790555061012082015181600101601b6101000a81548163ffffffff021916908363ffffffff1602179055509050507fbbe9920f263899c6c1294310cfb58d31969efcaca2afa033aeb4d6915f25f6fa6040518060e001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006bffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681525082604051611223929190615dcc565b600061115e826138fd565b60008080806139178686613908565b6000611530836001600160a01b0384166144cb565b336001600160a01b03821603613ef5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610b62565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000670de0b6b3a7640000613f91837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8616615507565b6115309190615531565b6000611530836001600160a01b0384166144d7565b6000612c16846001600160a01b038516846144e3565b6000613fe585613fd68486615507565b613fe0908761551e565b6144b5565b95945050505050565b6000614043826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166144f99092919063ffffffff16565b805190915015612fef57808060200190518101906140619190615602565b612fef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610b62565b600061153083836144cb565b60006115308383614508565b6000611530836001600160a01b038416614592565b6000612c16846001600160a01b038516846145af565b825474010000000000000000000000000000000000000000900460ff161580614157575081155b1561416157505050565b825460018401546fffffffffffffffffffffffffffffffff808316929116906000906141a790700100000000000000000000000000000000900463ffffffff1642615c50565b9050801561426757818311156141e9576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018601546142239083908590849070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16613fc6565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000004263ffffffff160217875592505b84821015614304576001600160a01b0384166142b9576040517ff94ebcd10000000000000000000000000000000000000000000000000000000081526004810183905260248101869052604401610b62565b6040517f1a76572a00000000000000000000000000000000000000000000000000000000815260048101839052602481018690526001600160a01b0385166044820152606401610b62565b848310156143fd5760018681015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff169060009082906143489082615c50565b614352878a615c50565b61435c919061551e565b6143669190615531565b90506001600160a01b0386166143b2576040517f15279c080000000000000000000000000000000000000000000000000000000081526004810182905260248101869052604401610b62565b6040517fd0c8d23a00000000000000000000000000000000000000000000000000000000815260048101829052602481018690526001600160a01b0387166044820152606401610b62565b6144078584615c50565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b600061115e826145cc565b6000808061449885856145d6565b600081815260029690960160205260409095205494959350505050565b60008183106144c45781611530565b5090919050565b600061153083836145e2565b60006115308383614592565b6000612c1684846001600160a01b0385166145af565b6060612c1684846000856145fa565b60008181526002830160205260408120548015158061452c575061452c84846144cb565b611530576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b657900006044820152606401610b62565b6000818152600283016020526040812081905561153083836146fb565b60008281526002840160205260408120829055612c168484614707565b600061115e825490565b60006115308383614713565b60008181526001830160205260408120541515611530565b60608247101561468c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610b62565b600080866001600160a01b031685876040516146a89190615e56565b60006040518083038185875af1925050503d80600081146146e5576040519150601f19603f3d011682016040523d82523d6000602084013e6146ea565b606091505b5091509150612c118783838761473d565b600061153083836147d0565b600061153083836148ca565b600082600001828154811061472a5761472a61571b565b9060005260206000200154905092915050565b606083156147c65782516000036147bf576001600160a01b0385163b6147bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610b62565b5081612c16565b612c168383614919565b600081815260018301602052604081205480156148b95760006147f4600183615c50565b855490915060009061480890600190615c50565b905081811461486d5760008660000182815481106148285761482861571b565b906000526020600020015490508087600001848154811061484b5761484b61571b565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061487e5761487e615e72565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061115e565b600091505061115e565b5092915050565b60008181526001830160205260408120546149115750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561115e565b50600061115e565b8151156149295781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b629190614a89565b60e0810161115e82846001600160a01b03808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015280606085015116606086015250506bffffffffffffffffffffffff60808301511660808401528060a08301511660a08401528060c08301511660c0840152505050565b6001600160a01b03811681146109f357600080fd5b80356149f9816149d9565b919050565b600060208284031215614a1057600080fd5b8135611530816149d9565b60005b83811015614a36578181015183820152602001614a1e565b50506000910152565b60008151808452614a57816020860160208601614a1b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006115306020830184614a3f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715614aee57614aee614a9c565b60405290565b6040805190810167ffffffffffffffff81118282101715614aee57614aee614a9c565b6040516080810167ffffffffffffffff81118282101715614aee57614aee614a9c565b604051610140810167ffffffffffffffff81118282101715614aee57614aee614a9c565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715614ba557614ba5614a9c565b604052919050565b600067ffffffffffffffff821115614bc757614bc7614a9c565b5060051b60200190565b63ffffffff811681146109f357600080fd5b80356149f981614bd1565b67ffffffffffffffff811681146109f357600080fd5b80356149f981614bee565b80151581146109f357600080fd5b80356149f981614c0f565b60006020808385031215614c3b57600080fd5b823567ffffffffffffffff811115614c5257600080fd5b8301601f81018513614c6357600080fd5b8035614c76614c7182614bad565b614b5e565b81815260e09182028301840191848201919088841115614c9557600080fd5b938501935b83851015614d405780858a031215614cb25760008081fd5b614cba614acb565b8535614cc5816149d9565b815285870135614cd481614bd1565b81880152604086810135614ce781614bd1565b90820152606086810135614cfa81614bd1565b908201526080614d0b878201614c04565b9082015260a0614d1c878201614c04565b9082015260c0614d2d878201614c1d565b9082015283529384019391850191614c9a565b50979650505050505050565b600060a0828403121561219b57600080fd5b600060208284031215614d7057600080fd5b813567ffffffffffffffff811115614d8757600080fd5b612c1684828501614d4c565b600082601f830112614da457600080fd5b81356020614db4614c7183614bad565b82815260069290921b84018101918181019086841115614dd357600080fd5b8286015b84811015614e245760408189031215614df05760008081fd5b614df8614af4565b8135614e03816149d9565b815281850135614e12816149d9565b81860152835291830191604001614dd7565b509695505050505050565b60008060408385031215614e4257600080fd5b823567ffffffffffffffff80821115614e5a57600080fd5b614e6686838701614d93565b93506020850135915080821115614e7c57600080fd5b50614e8985828601614d93565b9150509250929050565b60008060408385031215614ea657600080fd5b8235614eb1816149d9565b91506020830135614ec1816149d9565b809150509250929050565b80516001600160a01b031682526020810151614eee602084018261ffff169052565b506040810151614f06604084018263ffffffff169052565b506060810151614f1c606084018261ffff169052565b506080810151614f34608084018263ffffffff169052565b5060a0810151614f4a60a084018261ffff169052565b5060c0810151614f6060c084018261ffff169052565b5060e0810151614f7b60e08401826001600160a01b03169052565b506101008181015162ffffff16908301526101208082015163ffffffff81168285015261288d565b610140810161115e8284614ecc565b803561ffff811681146149f957600080fd5b60006020808385031215614fd757600080fd5b823567ffffffffffffffff811115614fee57600080fd5b8301601f81018513614fff57600080fd5b803561500d614c7182614bad565b81815260079190911b8201830190838101908783111561502c57600080fd5b928401925b828410156150a6576080848903121561504a5760008081fd5b615052614b17565b843561505d816149d9565b815261506a858701614fb2565b8682015260408086013561507d81614bd1565b9082015260608581013561509081614bd1565b9082015282526080939093019290840190615031565b979650505050505050565b600080602083850312156150c457600080fd5b823567ffffffffffffffff808211156150dc57600080fd5b818501915085601f8301126150f057600080fd5b8135818111156150ff57600080fd5b8660208260061b850101111561511457600080fd5b60209290920196919550909350505050565b60008060006060848603121561513b57600080fd5b833567ffffffffffffffff81111561515257600080fd5b61515e86828701614d4c565b935050602084013591506040840135615176816149d9565b809150509250925092565b600081518084526020808501945080840160005b838110156151c957815180516001600160a01b0316885283015161ffff168388015260409096019590820190600101615195565b509495945050505050565b6040815260006151e76040830185615181565b90508260208301529392505050565b80356fffffffffffffffffffffffffffffffff811681146149f957600080fd5b60006060828403121561522857600080fd5b6040516060810181811067ffffffffffffffff8211171561524b5761524b614a9c565b604052823561525981614c0f565b8152615267602084016151f6565b6020820152615278604084016151f6565b60408201529392505050565b803562ffffff811681146149f957600080fd5b600061014082840312156152aa57600080fd5b6152b2614b3a565b6152bb836149ee565b81526152c960208401614fb2565b60208201526152da60408401614be3565b60408201526152eb60608401614fb2565b60608201526152fc60808401614be3565b608082015261530d60a08401614fb2565b60a082015261531e60c08401614fb2565b60c082015261532f60e084016149ee565b60e0820152610100615342818501615284565b90820152610120615354848201614be3565b908201529392505050565b6020808252825182820181905260009190848201906040850190845b818110156153a05783516001600160a01b03168352928401929184019160010161537b565b50909695505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126153e157600080fd5b83018035915067ffffffffffffffff8211156153fc57600080fd5b60200191503681900382131561391f57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261544657600080fd5b83018035915067ffffffffffffffff82111561546157600080fd5b6020019150600681901b360382131561391f57600080fd5b80517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811681146149f957600080fd5b600080604083850312156154b857600080fd5b6154c183615479565b91506154cf60208401615479565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808202811582820484141761115e5761115e6154d8565b8082018082111561115e5761115e6154d8565b600082615567577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b67ffffffffffffffff8181168382160190808211156148c3576148c36154d8565b60006020828403121561559f57600080fd5b5051919050565b6000604082840312156155b857600080fd5b6155c0614af4565b82356155cb816149d9565b81526155d960208401614fb2565b60208201529392505050565b6000602082840312156155f757600080fd5b815161153081614bee565b60006020828403121561561457600080fd5b815161153081614c0f565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b602081526000612c1660208301848661561f565b60006020828403121561568e57600080fd5b5035919050565b6000604082840312156156a757600080fd5b6156af614af4565b82356156ba816149d9565b81526020928301359281019290925250919050565b6bffffffffffffffffffffffff8181168382160190808211156148c3576148c36154d8565b600067ffffffffffffffff808316818103615711576157116154d8565b6001019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6001600160a01b038716815260a06020820152600061576d60a08301878961561f565b85604084015267ffffffffffffffff8516606084015282810360808401526157958185614a3f565b9998505050505050505050565b6000602082840312156157b457600080fd5b815167ffffffffffffffff808211156157cc57600080fd5b818401915084601f8301126157e057600080fd5b8151818111156157f2576157f2614a9c565b61582360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601614b5e565b915080825285602082850101111561583a57600080fd5b61584b816020840160208601614a1b565b50949350505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203615885576158856154d8565b5060010190565b600081518084526020808501945080840160005b838110156151c957815180516001600160a01b0316885283015183880152604090960195908201906001016158a0565b600081518084526020808501808196508360051b8101915082860160005b85811015615918578284038952615906848351614a3f565b988501989350908401906001016158ee565b5091979650505050505050565b6020815261594060208201835167ffffffffffffffff169052565b6000602083015161595c60408401826001600160a01b03169052565b5060408301516001600160a01b038116606084015250606083015167ffffffffffffffff8116608084015250608083015160a083015260a08301516159a560c084018215159052565b5060c083015167ffffffffffffffff811660e08401525060e08301516101006159d8818501836001600160a01b03169052565b840151610120848101919091528401516101a061014080860182905291925090615a066101c0860184614a3f565b92508086015190507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610160818786030181880152615a45858461588c565b945080880151925050610180818786030181880152615a6485846158d0565b970151959092019490945250929392505050565b6bffffffffffffffffffffffff8281168282160390808211156148c3576148c36154d8565b602080825282518282018190526000919060409081850190868401855b8281101561591857815180516001600160a01b031685528681015163ffffffff90811688870152868201518116878701526060808301519091169086015260808082015167ffffffffffffffff169086015260a080820151615b278288018267ffffffffffffffff169052565b505060c09081015115159085015260e09093019290850190600101615aba565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015615b875780818660040360031b1b83161692505b505092915050565b60008085851115615b9f57600080fd5b83861115615bac57600080fd5b5050820193919092039150565b600060208284031215615bcb57600080fd5b6040516020810181811067ffffffffffffffff82111715615bee57615bee614a9c565b6040529135825250919050565b600060208284031215615c0d57600080fd5b61153082615479565b63ffffffff8181168382160190808211156148c3576148c36154d8565b600060208284031215615c4557600080fd5b8151611530816149d9565b8181038181111561115e5761115e6154d8565b81810360008312801583831316838312821617156148c3576148c36154d8565b602080825282518282018190526000919060409081850190868401855b8281101561591857815180516001600160a01b031685528681015161ffff16878601528581015163ffffffff90811687870152606091820151169085015260809093019290850190600101615ca0565b600081615cff57615cff6154d8565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b63ffffffff83168152604060208201526000612c166040830184615181565b600060408284031215615d5657600080fd5b615d5e614af4565b615d6783615479565b815260208301516155d981614bd1565b602081526000611530602083018461588c565b60208152600061153060208301846158d0565b60008351615daf818460208801614a1b565b835190830190615dc3818360208801614a1b565b01949350505050565b6102208101615e4982856001600160a01b03808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015280606085015116606086015250506bffffffffffffffffffffffff60808301511660808401528060a08301511660a08401528060c08301511660c0840152505050565b61153060e0830184614ecc565b60008251615e68818460208701614a1b565b9190910192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a",
}

var EVM2EVMOnRampABI = EVM2EVMOnRampMetaData.ABI

var EVM2EVMOnRampBin = EVM2EVMOnRampMetaData.Bin

func DeployEVM2EVMOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMOnRampStaticConfig, dynamicConfig EVM2EVMOnRampDynamicConfig, tokensAndPools []InternalPoolUpdate, rateLimiterConfig RateLimiterConfig, feeTokenConfigs []EVM2EVMOnRampFeeTokenConfigArgs, tokenTransferFeeConfigArgs []EVM2EVMOnRampTokenTransferFeeConfigArgs, nopsAndWeights []EVM2EVMOnRampNopAndWeight) (common.Address, *types.Transaction, *EVM2EVMOnRamp, error) {
	parsed, err := EVM2EVMOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMOnRampBin), backend, staticConfig, dynamicConfig, tokensAndPools, rateLimiterConfig, feeTokenConfigs, tokenTransferFeeConfigArgs, nopsAndWeights)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMOnRamp{EVM2EVMOnRampCaller: EVM2EVMOnRampCaller{contract: contract}, EVM2EVMOnRampTransactor: EVM2EVMOnRampTransactor{contract: contract}, EVM2EVMOnRampFilterer: EVM2EVMOnRampFilterer{contract: contract}}, nil
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetFee(opts *bind.CallOpts, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getFee", message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetFee(message ClientEVM2AnyMessage) (*big.Int, error) {
	return _EVM2EVMOnRamp.Contract.GetFee(&_EVM2EVMOnRamp.CallOpts, message)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetFee(message ClientEVM2AnyMessage) (*big.Int, error) {
	return _EVM2EVMOnRamp.Contract.GetFee(&_EVM2EVMOnRamp.CallOpts, message)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getPoolBySourceToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOnRamp.Contract.GetPoolBySourceToken(&_EVM2EVMOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOnRamp.Contract.GetPoolBySourceToken(&_EVM2EVMOnRamp.CallOpts, sourceToken)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMOnRamp.Contract.GetSupportedTokens(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMOnRamp.Contract.GetSupportedTokens(&_EVM2EVMOnRamp.CallOpts)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) ApplyPoolUpdates(opts *bind.TransactOpts, removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "applyPoolUpdates", removes, adds)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) ApplyPoolUpdates(removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.ApplyPoolUpdates(&_EVM2EVMOnRamp.TransactOpts, removes, adds)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) ApplyPoolUpdates(removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.ApplyPoolUpdates(&_EVM2EVMOnRamp.TransactOpts, removes, adds)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) ForwardFromRouter(opts *bind.TransactOpts, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "forwardFromRouter", message, feeTokenAmount, originalSender)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) ForwardFromRouter(message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.ForwardFromRouter(&_EVM2EVMOnRamp.TransactOpts, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) ForwardFromRouter(message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.ForwardFromRouter(&_EVM2EVMOnRamp.TransactOpts, message, feeTokenAmount, originalSender)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetTokenTransferFeeConfig(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []EVM2EVMOnRampTokenTransferFeeConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setTokenTransferFeeConfig", tokenTransferFeeConfigArgs)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetTokenTransferFeeConfig(tokenTransferFeeConfigArgs []EVM2EVMOnRampTokenTransferFeeConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetTokenTransferFeeConfig(&_EVM2EVMOnRamp.TransactOpts, tokenTransferFeeConfigArgs)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetTokenTransferFeeConfig(tokenTransferFeeConfigArgs []EVM2EVMOnRampTokenTransferFeeConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetTokenTransferFeeConfig(&_EVM2EVMOnRamp.TransactOpts, tokenTransferFeeConfigArgs)
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

type EVM2EVMOnRampPoolAddedIterator struct {
	Event *EVM2EVMOnRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampPoolAdded)
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
		it.Event = new(EVM2EVMOnRampPoolAdded)
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

func (it *EVM2EVMOnRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMOnRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampPoolAddedIterator{contract: _EVM2EVMOnRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampPoolAdded)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMOnRampPoolAdded, error) {
	event := new(EVM2EVMOnRampPoolAdded)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampPoolRemovedIterator struct {
	Event *EVM2EVMOnRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampPoolRemoved)
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
		it.Event = new(EVM2EVMOnRampPoolRemoved)
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

func (it *EVM2EVMOnRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMOnRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampPoolRemovedIterator{contract: _EVM2EVMOnRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampPoolRemoved)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMOnRampPoolRemoved, error) {
	event := new(EVM2EVMOnRampPoolRemoved)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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
	case _EVM2EVMOnRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMOnRamp.ParsePoolAdded(log)
	case _EVM2EVMOnRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMOnRamp.ParsePoolRemoved(log)
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
	return common.HexToHash("0xbbe9920f263899c6c1294310cfb58d31969efcaca2afa033aeb4d6915f25f6fa")
}

func (EVM2EVMOnRampFeeConfigSet) Topic() common.Hash {
	return common.HexToHash("0x2386f61ab5cafc3fed44f9f614f721ab53479ef64067fd16c1a2491b63ddf1a8")
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

func (EVM2EVMOnRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMOnRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMOnRampTokenTransferFeeConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3df02211160f78e9bbed13e8bf0752035362a05f0651a5e25007acfb659c3620")
}

func (_EVM2EVMOnRamp *EVM2EVMOnRamp) Address() common.Address {
	return _EVM2EVMOnRamp.address
}

type EVM2EVMOnRampInterface interface {
	CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error)

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMOnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetFee(opts *bind.CallOpts, message ClientEVM2AnyMessage) (*big.Int, error)

	GetFeeTokenConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMOnRampFeeTokenConfig, error)

	GetNopFeesJuels(opts *bind.CallOpts) (*big.Int, error)

	GetNops(opts *bind.CallOpts) (GetNops,

		error)

	GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMOnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMOnRampTokenTransferFeeConfig, error)

	LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyPoolUpdates(opts *bind.TransactOpts, removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	PayNops(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMOnRampDynamicConfig) (*types.Transaction, error)

	SetFeeTokenConfig(opts *bind.TransactOpts, feeTokenConfigArgs []EVM2EVMOnRampFeeTokenConfigArgs) (*types.Transaction, error)

	SetNops(opts *bind.TransactOpts, nopsAndWeights []EVM2EVMOnRampNopAndWeight) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	SetTokenTransferFeeConfig(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []EVM2EVMOnRampTokenTransferFeeConfigArgs) (*types.Transaction, error)

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

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMOnRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMOnRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMOnRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMOnRampPoolRemoved, error)

	FilterTokenTransferFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMOnRampTokenTransferFeeConfigSetIterator, error)

	WatchTokenTransferFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampTokenTransferFeeConfigSet) (event.Subscription, error)

	ParseTokenTransferFeeConfigSet(log types.Log) (*EVM2EVMOnRampTokenTransferFeeConfigSet, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
