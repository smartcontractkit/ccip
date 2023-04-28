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
	Router          common.Address
	PriceRegistry   common.Address
	MaxDataSize     uint32
	MaxGasLimit     uint64
	MaxTokensLength uint16
	Afn             common.Address
}

type EVM2EVMOnRampFeeTokenConfig struct {
	FeeAmount       *big.Int
	Multiplier      uint64
	DestGasOverhead uint32
}

type EVM2EVMOnRampFeeTokenConfigArgs struct {
	Token           common.Address
	Multiplier      uint64
	FeeAmount       *big.Int
	DestGasOverhead uint32
}

type EVM2EVMOnRampNopAndWeight struct {
	Nop    common.Address
	Weight uint16
}

type EVM2EVMOnRampStaticConfig struct {
	LinkToken         common.Address
	ChainId           uint64
	DestChainId       uint64
	DefaultTxGasLimit uint64
}

type EVM2EVMOnRampTokenAndPool struct {
	Token common.Address
	Pool  common.Address
}

type InternalEVM2EVMMessage struct {
	SourceChainId  uint64
	SequenceNumber uint64
	FeeTokenAmount *big.Int
	Sender         common.Address
	Nonce          uint64
	GasLimit       *big.Int
	Strict         bool
	Receiver       common.Address
	Data           []byte
	TokenAmounts   []ClientEVMTokenAmount
	FeeToken       common.Address
	MessageId      [32]byte
}

type InternalPoolUpdate struct {
	Token common.Address
	Pool  common.Address
}

type RateLimiterConfig struct {
	IsEnabled bool
	Rate      *big.Int
	Capacity  *big.Int
}

type RateLimiterTokenBucket struct {
	Capacity    *big.Int
	Tokens      *big.Int
	Rate        *big.Int
	LastUpdated *big.Int
	IsEnabled   bool
}

var EVM2EVMOnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOnRamp.TokenAndPool[]\",\"name\":\"tokensAndPools\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"multiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ConsumingMoreThanMaxCapacity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"expected\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"got\",\"type\":\"bytes4\"}],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InvalidFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"InvalidWithdrawalAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkBalanceNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeesToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNopsToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrFeeAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrFeeAdminOrNop\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"RateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyNops\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"multiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeConfig\",\"type\":\"tuple[]\"}],\"name\":\"FeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NopPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nopWeightsTotal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"NopsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"}],\"name\":\"applyPoolUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdated\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"multiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfig\",\"name\":\"feeTokenConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNopFeesJuels\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNops\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"weightsTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAllowListEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"multiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"setNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawNonLinkFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162006d8f38038062006d8f8339810160408190526200003591620019bc565b6000805460ff1916815583903390819081620000985760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000d257620000d28162000416565b50506040805160a0810182528382018051808352905160208084018290528601516001600160d01b03169383018490524264ffffffffff1660608401819052955115156080909301839052600391909155600455600580547fff0000000000000000000000000000000000000000000000000000000000000016909217600160d01b909402939093176001600160f81b0316600160f81b909302929092179091555086516001600160a01b0316158062000197575060208701516001600160401b0316155b80620001ae575060408701516001600160401b0316155b80620001c5575060608701516001600160401b0316155b15620001e4576040516306b7c75960e31b815260040160405180910390fd5b6020808801516040808a015181517fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a948101949094526001600160401b039283169184019190915216606082015230608082015260a00160408051601f19818403018152918152815160209283012060805288516001600160a01b031660c052908801516001600160401b0390811660e0529088015181166101005260608801511660a0526200029486620004c7565b6200029f82620006c4565b620002aa816200084e565b60005b8551811015620003bb5760006001600160a01b0316868281518110620002d757620002d762001ab7565b6020026020010151600001516001600160a01b031614806200032b575060006001600160a01b031686828151811062000314576200031462001ab7565b6020026020010151602001516001600160a01b0316145b156200034a576040516306b7c75960e31b815260040160405180910390fd5b620003a786828151811062000363576200036362001ab7565b60200260200101516000015187838151811062000384576200038462001ab7565b602002602001015160200151600c62000a0960201b62002228179092919060201c565b50620003b38162001ae3565b9050620002ad565b508351156200040957600f805460ff191660011790556200040984600060405190808252806020026020018201604052801562000402578160200160208202803683370190505b5062000a37565b5050505050505062001d7c565b336001600160a01b03821603620004705760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200008f565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b03161580620004ea575060208101516001600160a01b0316155b8062000501575060a08101516001600160a01b0316155b1562000520576040516306b7c75960e31b815260040160405180910390fd5b8051600680546001600160a01b039283166001600160a01b0319909116179055602080830151600780546040808701516060808901516001600160401b03908116600160c01b026001600160c01b0363ffffffff909416600160a01b026001600160c01b0319909616978a16979097179490941791909116949094179092556080808701516008805460a0808b01518a1662010000026001600160b01b031990921661ffff909416939093171790558351918201845260c051909616815260e0518216948101949094526101005181168483015293519093169082015290517ff7736cb9ac6dabe6bc85a757f7dd584df95f4e92afdebd0cb6994b4aef74653491620006b991849082516001600160a01b0390811682526020808501516001600160401b03908116828501526040808701518216818601526060968701518216878601528551841660808087019190915292860151841660a0808701919091529086015163ffffffff1660c0860152958501511660e084015283015161ffff166101008301529190920151166101208201526101400190565b60405180910390a150565b60005b81518110156200081c576040518060600160405280838381518110620006f157620006f162001ab7565b6020026020010151604001516001600160601b0316815260200183838151811062000720576200072062001ab7565b6020026020010151602001516001600160401b031681526020018383815181106200074f576200074f62001ab7565b60200260200101516060015163ffffffff16815250601260008484815181106200077d576200077d62001ab7565b602090810291909101810151516001600160a01b031682528181019290925260409081016000208351815493850151949092015163ffffffff16600160a01b0263ffffffff60a01b196001600160401b039095166c01000000000000000000000000026001600160a01b03199094166001600160601b03909316929092179290921792909216919091179055620008148162001ae3565b9050620006c7565b507f1026720e82621bb6d100957028cafff5ccf24add051c05d1bd65fa1d71e0fae681604051620006b9919062001aff565b6014546c01000000000000000000000000900463ffffffff161580159062000897575060145463ffffffff6c010000000000000000000000008204166001600160601b03909116115b15620008a757620008a762000ba0565b80516040811115620008cc57604051635ad0867d60e11b815260040160405180910390fd5b60096000818181818181620008e2828262001438565b50505050505050506000805b82811015620009a0576200095884828151811062000910576200091062001ab7565b60200260200101516000015185838151811062000931576200093162001ab7565b60200260200101516020015161ffff16600962000e6e60201b62002246179092919060201c565b508381815181106200096e576200096e62001ab7565b60200260200101516020015161ffff16826200098b919062001b86565b9150620009988162001ae3565b9050620008ee565b506014805463ffffffff60601b19166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd2490620009fc908390869062001bb1565b60405180910390a1505050565b600062000a2d846001600160a01b0385168462000e86602090811b6200225c17901c565b90505b9392505050565b60005b825181101562000ad757600083828151811062000a5b5762000a5b62001ab7565b6020026020010151905062000a8081601062000e9e60201b620022721790919060201c565b1562000ac3576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b5062000acf8162001ae3565b905062000a3a565b5060005b815181101562000b9b57600082828151811062000afc5762000afc62001ab7565b6020026020010151905060006001600160a01b0316816001600160a01b03160362000b28575062000b88565b62000b4381601062000ebe60201b620022871790919060201c565b1562000b86576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b62000b938162001ae3565b905062000adb565b505050565b62000bb860005461010090046001600160a01b031690565b6001600160a01b0316336001600160a01b03161415801562000be557506002546001600160a01b03163314155b801562000c0a575062000c0833600962000ed560201b6200229c1790919060201c565b155b1562000c29576040516363baed9960e01b815260040160405180910390fd5b6014546c01000000000000000000000000900463ffffffff16600081900362000c655760405163990e30bf60e01b815260040160405180910390fd5b6014546001600160601b031663ffffffff821681101562000c99576040516311a1ee3b60e31b815260040160405180910390fd5b60c0516040516370a0823160e01b81523060048201526001600160601b038316916001600160a01b0316906370a0823190602401602060405180830381865afa15801562000ceb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000d11919062001c1e565b101562000d3157604051631e9acf1760e31b815260040160405180910390fd5b6000819050600062000d4f600962000eec60201b620022b11760201c565b905060005b8181101562000e485760008062000d7b83600962000ef960201b620022bc1790919060201c565b9092509050600063ffffffff881662000d9e836001600160601b038a1662001c38565b62000daa919062001c5a565b905062000db8818762001c7d565b955062000dea83826001600160601b031660c0516001600160a01b031662000f1760201b620022da179092919060201c565b6040516001600160601b03821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a25050508062000e409062001ae3565b905062000d54565b5050601480546001600160601b0319166001600160601b03929092169190911790555050565b600062000a2d846001600160a01b0385168462000f6f565b600062000a2d84846001600160a01b03851662000f6f565b600062000eb5836001600160a01b03841662000f9c565b90505b92915050565b600062000eb5836001600160a01b038416620010a0565b600062000eb5836001600160a01b038416620010f2565b600062000eb88262001111565b600080808062000f0a86866200112c565b9097909650945050505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b0390811663a9059cbb60e01b1790915262000b9b9185916200116b16565b60008281526002840160209081526040822083905562000a2d90859085906200235a62001249821b17901c565b600081815260018301602052604081205480156200109557600062000fc360018362001ca8565b855490915060009062000fd99060019062001ca8565b90508181146200104557600086600001828154811062000ffd5762000ffd62001ab7565b906000526020600020015490508087600001848154811062001023576200102362001ab7565b6000918252602080832090910192909255918252600188019052604090208390555b855486908062001059576200105962001cc2565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000eb8565b600091505062000eb8565b6000818152600183016020526040812054620010e95750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000eb8565b50600062000eb8565b600062000eb582846000016200125760201b620023661790919060201c565b600062000eb8826000016200127060201b6200237e1760201c565b60008060006200114e84866000016200127b60201b620023881790919060201c565b600081815260029690960160205260409095205494959350505050565b6000620011c7826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166200128960201b62002394179092919060201c565b80519091501562000b9b5780806020019051810190620011e8919062001cd8565b62000b9b5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016200008f565b600062000eb58383620010a0565b6000818152600183016020526040812054151562000eb5565b600062000eb8825490565b600062000eb583836200129a565b606062000a2d8484600085620012c7565b6000826000018281548110620012b457620012b462001ab7565b9060005260206000200154905092915050565b6060824710156200132a5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016200008f565b843b6200137a5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016200008f565b600080866001600160a01b0316858760405162001398919062001d29565b60006040518083038185875af1925050503d8060008114620013d7576040519150601f19603f3d011682016040523d82523d6000602084013e620013dc565b606091505b509092509050620013ef828286620013fa565b979650505050505050565b606083156200140b57508162000a30565b8251156200141c5782518084602001fd5b8160405162461bcd60e51b81526004016200008f919062001d47565b50805460008255906000526020600020908101906200145891906200145b565b50565b5b808211156200147257600081556001016200145c565b5090565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715620014b157620014b162001476565b60405290565b604080519081016001600160401b0381118282101715620014b157620014b162001476565b604051601f8201601f191681016001600160401b038111828210171562001507576200150762001476565b604052919050565b80516001600160a01b03811681146200152757600080fd5b919050565b80516001600160401b03811681146200152757600080fd5b6000608082840312156200155757600080fd5b620015616200148c565b90506200156e826200150f565b81526200157e602083016200152c565b602082015262001591604083016200152c565b6040820152620015a4606083016200152c565b606082015292915050565b805163ffffffff811681146200152757600080fd5b805161ffff811681146200152757600080fd5b600060c08284031215620015ea57600080fd5b60405160c081016001600160401b03811182821017156200160f576200160f62001476565b60405290508062001620836200150f565b815262001630602084016200150f565b60208201526200164360408401620015af565b604082015262001656606084016200152c565b60608201526200166960808401620015c4565b60808201526200167c60a084016200150f565b60a08201525092915050565b60006001600160401b03821115620016a457620016a462001476565b5060051b60200190565b600082601f830112620016c057600080fd5b81516020620016d9620016d38362001688565b620014dc565b82815260069290921b84018101918181019086841115620016f957600080fd5b8286015b848110156200174f5760408189031215620017185760008081fd5b62001722620014b7565b6200172d826200150f565b81526200173c8583016200150f565b81860152835291830191604001620016fd565b509695505050505050565b600082601f8301126200176c57600080fd5b815160206200177f620016d38362001688565b82815260059290921b840181019181810190868411156200179f57600080fd5b8286015b848110156200174f57620017b7816200150f565b8352918301918301620017a3565b805180151581146200152757600080fd5b600060608284031215620017e957600080fd5b604051606081016001600160401b03811182821017156200180e576200180e62001476565b6040529050806200181f83620017c5565b815260208301516001600160d01b03811681146200183c57600080fd5b6020820152604092830151920191909152919050565b600082601f8301126200186457600080fd5b8151602062001877620016d38362001688565b82815260079290921b840181019181810190868411156200189757600080fd5b8286015b848110156200174f5760808189031215620018b65760008081fd5b620018c06200148c565b620018cb826200150f565b8152620018da8583016200152c565b818601526040828101516001600160601b0381168114620018fb5760008081fd5b9082015260606200190e838201620015af565b908201528352918301916080016200189b565b600082601f8301126200193357600080fd5b8151602062001946620016d38362001688565b82815260069290921b840181019181810190868411156200196657600080fd5b8286015b848110156200174f5760408189031215620019855760008081fd5b6200198f620014b7565b6200199a826200150f565b8152620019a9858301620015c4565b818601528352918301916040016200196a565b6000806000806000806000610220888a031215620019d957600080fd5b620019e5898962001544565b9650620019f68960808a01620015d7565b6101408901519096506001600160401b038082111562001a1557600080fd5b62001a238b838c01620016ae565b96506101608a015191508082111562001a3b57600080fd5b62001a498b838c016200175a565b955062001a5b8b6101808c01620017d6565b94506101e08a015191508082111562001a7357600080fd5b62001a818b838c0162001852565b93506102008a015191508082111562001a9957600080fd5b5062001aa88a828b0162001921565b91505092959891949750929550565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820162001af85762001af862001acd565b5060010190565b602080825282518282018190526000919060409081850190868401855b8281101562001b7957815180516001600160a01b03168552868101516001600160401b031687860152858101516001600160601b03168686015260609081015163ffffffff16908501526080909301929085019060010162001b1c565b5091979650505050505050565b600063ffffffff80831681851680830382111562001ba85762001ba862001acd565b01949350505050565b6000604080830163ffffffff8616845260208281860152818651808452606087019150828801935060005b8181101562001c1057845180516001600160a01b0316845284015161ffff1684840152938301939185019160010162001bdc565b509098975050505050505050565b60006020828403121562001c3157600080fd5b5051919050565b600081600019048311821515161562001c555762001c5562001acd565b500290565b60008262001c7857634e487b7160e01b600052601260045260246000fd5b500490565b60006001600160601b038381169083168181101562001ca05762001ca062001acd565b039392505050565b60008282101562001cbd5762001cbd62001acd565b500390565b634e487b7160e01b600052603160045260246000fd5b60006020828403121562001ceb57600080fd5b62000eb582620017c5565b60005b8381101562001d1357818101518382015260200162001cf9565b8381111562001d23576000848401525b50505050565b6000825162001d3d81846020870162001cf6565b9190910192915050565b602081526000825180602084015262001d6881604085016020870162001cf6565b601f01601f19169190910160400192915050565b60805160a05160c05160e05161010051614f7362001e1c600039600081816102a30152818161094c01526128ee0152600081816102730152818161194801526128c401526000818161024401528181610e9901528181610f89015281816117830152818161187c015281816120210152818161213f015261289e0152600081816102d301528181612507015261291501526000611b2a0152614f736000f3fe608060405234801561001057600080fd5b506004361061020b5760003560e01c80635d86f1411161012a5780638da5cb5b116100bd578063d3c7c2c71161008c578063efeadb6d11610071578063efeadb6d14610793578063eff7cc48146107a6578063f2fde38b146107ae57600080fd5b8063d3c7c2c714610780578063e0351e131461078857600080fd5b80638da5cb5b1461072c578063a7cd63b714610742578063a7d3e02f14610757578063b06d41bc1461076a57600080fd5b806376f6ae76116100f957806376f6ae76146106d657806379ba5097146106e95780638456cb59146106f1578063856c8247146106f957600080fd5b80635d86f141146105c657806368d8f142146105d9578063704b6c02146105ec5780637437ff9f146105ff57600080fd5b80633f4ba83a116101a257806354b714681161017157806354b714681461056357806354c8a4f314610583578063599f6431146105965780635c975abb146105bb57600080fd5b80633f4ba83a146104b85780634120fccd146104c0578063546719cd146104e1578063549e946f1461055057600080fd5b80633091aee7116101de5780633091aee71461045c57806338724a95146104715780633a87ac53146104925780633b1a1403146104a557600080fd5b806306285c6914610210578063147809b31461031a578063181f5a77146103325780631e442b551461037b575b600080fd5b61030460408051608081018252600080825260208201819052918101829052606081019190915260405180608001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16815250905090565b6040516103119190613ea6565b60405180910390f35b6103226107c1565b6040519015158152602001610311565b61036e6040518060400160405280601381526020017f45564d3245564d4f6e52616d7020312e302e300000000000000000000000000081525081565b6040516103119190613f66565b61041c610389366004613f8e565b6040805160608082018352600080835260208084018290529284018190526001600160a01b03949094168452601282529282902082519384018352546bffffffffffffffffffffffff811684526c01000000000000000000000000810467ffffffffffffffff169184019190915274010000000000000000000000000000000000000000900463ffffffff169082015290565b6040805182516bffffffffffffffffffffffff16815260208084015167ffffffffffffffff16908201529181015163ffffffff1690820152606001610311565b61046f61046a366004614083565b610843565b005b61048461047f36600461411f565b6108c5565b604051908152602001610311565b61046f6104a0366004614219565b610aa1565b61046f6104b33660046142be565b610d6e565b61046f610d7f565b6104c8610d91565b60405167ffffffffffffffff9091168152602001610311565b6104e9610dc5565b6040516103119190600060a082019050825182526020830151602083015279ffffffffffffffffffffffffffffffffffffffffffffffffffff604084015116604083015264ffffffffff606084015116606083015260808301511515608083015292915050565b61046f61055e366004614365565b610e8f565b6014546040516bffffffffffffffffffffffff9091168152602001610311565b61046f6105913660046143e3565b6110e1565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610311565b60005460ff16610322565b6105a36105d4366004613f8e565b61115c565b61046f6105e736600461444f565b6111bb565b61046f6105fa366004613f8e565b611238565b6106c96040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c0810182526006546001600160a01b039081168252600754808216602084015274010000000000000000000000000000000000000000810463ffffffff1693830193909352780100000000000000000000000000000000000000000000000090920467ffffffffffffffff16606082015260085461ffff8116608083015262010000900490911660a082015290565b6040516103119190614543565b61046f6106e43660046145a9565b611319565b61046f611380565b61046f611488565b6104c8610707366004613f8e565b6001600160a01b031660009081526013602052604090205467ffffffffffffffff1690565b60005461010090046001600160a01b03166105a3565b61074a611498565b604051610311919061461e565b61048461076536600461466b565b611549565b610772611ca0565b604051610311929190614719565b61074a611da4565b600f5460ff16610322565b61046f6107a136600461473b565b611e4f565b61046f611eb6565b61046f6107bc366004613f8e565b612217565b6000600660020160029054906101000a90046001600160a01b03166001600160a01b031663397796f76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610819573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083d9190614758565b15905090565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561088057506002546001600160a01b03163314155b156108b7576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108c26003826123a3565b50565b6000806108dd6108d86080850185614775565b6124d4565b516007549091506000906001600160a01b031663268e5d486109056080870160608801613f8e565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166024820152604401602060405180830381865afa158015610991573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b591906147da565b905060006012816109cc6080880160608901613f8e565b6001600160a01b031681526020808201929092526040908101600020815160608101835290546bffffffffffffffffffffffff811682526c01000000000000000000000000810467ffffffffffffffff1693820184905274010000000000000000000000000000000000000000900463ffffffff169181018290529250670de0b6b3a764000091908490610a609087614822565b610a6a919061483a565b610a74919061483a565b610a7e9190614877565b8151610a9891906bffffffffffffffffffffffff16614822565b95945050505050565b610aa9612644565b60005b8251811015610c0f576000838281518110610ac957610ac96148b2565b60200260200101516000015190506000848381518110610aeb57610aeb6148b2565b6020026020010151602001519050610b0d82600c6126bd90919063ffffffff16565b610b53576040517f73913ebd0000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024015b60405180910390fd5b6001600160a01b038116610b68600c846126d2565b6001600160a01b031614610ba8576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610bb3600c836126e7565b15610bfc57604080516001600160a01b038085168252831660208201527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15b505080610c08906148e1565b9050610aac565b5060005b8151811015610d69576000828281518110610c3057610c306148b2565b60200260200101516000015190506000838381518110610c5257610c526148b2565b602002602001015160200151905060006001600160a01b0316826001600160a01b03161480610c8857506001600160a01b038116155b15610cbf576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cca600c836126bd565b15610d01576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d0d600c8383612228565b15610d5657604080516001600160a01b038085168252831660208201527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a15b505080610d62906148e1565b9050610c13565b505050565b610d76612644565b6108c2816126fc565b610d87612644565b610d8f61296c565b565b601454600090610dc090700100000000000000000000000000000000900467ffffffffffffffff166001614919565b905090565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526003548152600454602082015260055479ffffffffffffffffffffffffffffffffffffffffffffffffffff8116928201929092527a010000000000000000000000000000000000000000000000000000820464ffffffffff1660608201527f010000000000000000000000000000000000000000000000000000000000000090910460ff1615156080820152610dc090612a40565b610e97612644565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b031603610f0d576040517feddf07f50000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610b4a565b6001600160a01b038116610f58576040517f21680a040000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401610b4a565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa158015610fd8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ffc91906147da565b6014549091506bffffffffffffffffffffffff16811015611049576040517f02075e0000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152610d699083906001600160a01b038616906370a0823190602401602060405180830381865afa1580156110ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d091906147da565b6001600160a01b03861691906122da565b6110e9612644565b61115684848080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808802828101820190935287825290935087925086918291850190849080828437600092019190915250612ae392505050565b50505050565b6000611169600c836126bd565b6111aa576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610b4a565b6111b5600c836126d2565b92915050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b0316141580156111f857506002546001600160a01b03163314155b1561122f576040517fdf938f4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108c281612c1e565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561127557506002546001600160a01b03163314155b156112ac576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c906020015b60405180910390a150565b611321612644565b61137c8282808060200260200160405190810160405280939291908181526020016000905b828210156113725761136360408302860136819003810190614945565b81526020019060010190611346565b5050505050612de3565b5050565b6001546001600160a01b031633146113f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610b4a565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b611490612644565b610d8f612fb5565b606060006114a6601061237e565b67ffffffffffffffff8111156114be576114be613fab565b6040519080825280602002602001820160405280156114e7578160200160208202803683370190505b50905060005b6114f7601061237e565b81101561154357611509601082612388565b82828151811061151b5761151b6148b2565b6001600160a01b039092166020928302919091019091015261153c816148e1565b90506114ed565b50919050565b6000805460ff16156115b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610b4a565b600660020160029054906101000a90046001600160a01b03166001600160a01b031663397796f76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561160d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116319190614758565b15611667576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006116796108d86080870187614775565b90506116f461168b6020870187614775565b835190915061169d6040890189614984565b808060200260200160405190810160405280939291908181526020016000905b828210156116e9576116da604083028601368190038101906149ec565b815260200190600101906116bd565b505050505086613075565b6116fe8580614775565b9050602014611745576117118580614775565b6040517f370d875f000000000000000000000000000000000000000000000000000000008152600401610b4a929190614a26565b60006117518680614775565b81019061175e9190614a73565b90506001600160a01b03811115611779576117118680614775565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166117b36080880160608901613f8e565b6001600160a01b03160361181757601480548691906000906117e49084906bffffffffffffffffffffffff16614a8c565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff160217905550611938565b6007546001600160a01b03166241e5be6118376080890160608a01613f8e565b60405160e083901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b039182166004820152602481018990527f000000000000000000000000000000000000000000000000000000000000000090911660448201526064016020604051808303816000875af11580156118c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118e991906147da565b601480546000906119099084906bffffffffffffffffffffffff16614a8c565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055505b60006040518061018001604052807f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020016014601081819054906101000a900467ffffffffffffffff1661199890614ab3565b825467ffffffffffffffff9182166101009390930a8381029083021990911617909255825260208083018a90526001600160a01b038916604080850182905260009182526013909252908120805460609094019390926119f89116614ab3565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff16815260200184600001518152602001846020015115158152602001836001600160a01b03168152602001888060200190611a619190614775565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250602001611aa860408a018a614984565b808060200260200160405190810160405280939291908181526020016000905b82821015611af457611ae5604083028601368190038101906149ec565b81526020019060010190611ac8565b5050509183525050602001611b0f60808a0160608b01613f8e565b6001600160a01b0316815260006020909101529050611b4e817f0000000000000000000000000000000000000000000000000000000000000000613295565b61016082015260005b611b646040890189614984565b9050811015611c57576000611b7c60408a018a614984565b83818110611b8c57611b8c6148b2565b905060400201803603810190611ba291906149ec565b90506000611bb3826000015161115c565b60208301516040517fe2e59b3e0000000000000000000000000000000000000000000000000000000081529192506001600160a01b0383169163e2e59b3e91611c12918c906004019182526001600160a01b0316602082015260400190565b600060405180830381600087803b158015611c2c57600080fd5b505af1158015611c40573d6000803e3d6000fd5b50505050505080611c50906148e1565b9050611b57565b507faffc45517195d6499808c643bd4a7b0ffeedf95bea5852840d7bfcf63f59e82181604051611c879190614b1e565b60405180910390a16101600151925050505b9392505050565b6060600080611caf60096122b1565b90508067ffffffffffffffff811115611cca57611cca613fab565b604051908082528060200260200182016040528015611d0f57816020015b6040805180820190915260008082526020820152815260200190600190039081611ce85790505b50925060005b81811015611d8157600080611d2b6009846122bc565b915091506040518060400160405280836001600160a01b031681526020018261ffff16815250868481518110611d6357611d636148b2565b6020026020010181905250505080611d7a906148e1565b9050611d15565b505060145491926c0100000000000000000000000090920463ffffffff16919050565b60606000611db2600c61339f565b67ffffffffffffffff811115611dca57611dca613fab565b604051908082528060200260200182016040528015611df3578160200160208202803683370190505b50905060005b815181101561154357611e0d600c826133aa565b50828281518110611e2057611e206148b2565b60200260200101816001600160a01b03166001600160a01b03168152505080611e48906148e1565b9050611df9565b611e57612644565b600f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168215159081179091556040519081527fccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df0329060200161130e565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015611ef357506002546001600160a01b03163314155b8015611f075750611f0560093361229c565b155b15611f3e576040517f63baed9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6014546c01000000000000000000000000900463ffffffff166000819003611f92576040517f990e30bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6014546bffffffffffffffffffffffff1663ffffffff8216811015611fe3576040517f8d0f71d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526bffffffffffffffffffffffff8216907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa158015612070573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061209491906147da565b10156120cc576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060006120d960096122b1565b905060005b818110156121d4576000806120f46009846122bc565b9092509050600063ffffffff881661211a836bffffffffffffffffffffffff8a1661483a565b6121249190614877565b90506121308187614c5b565b95506121746001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016846bffffffffffffffffffffffff84166122da565b6040516bffffffffffffffffffffffff821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a2505050806121cd906148e1565b90506120de565b5050601480547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff929092169190911790555050565b61221f612644565b6108c2816133b9565b600061223e846001600160a01b0385168461225c565b949350505050565b600061223e846001600160a01b0385168461349a565b600061223e84846001600160a01b03851661349a565b6000611c99836001600160a01b0384166134b7565b6000611c99836001600160a01b0384166135aa565b6000611c99836001600160a01b0384166135f9565b60006111b582613605565b60008080806122cb8686613610565b909450925050505b9250929050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610d6990849061363b565b6000611c9983836135aa565b60008181526001830160205260408120541515611c99565b60006111b5825490565b6000611c99838361373a565b606061223e8484600085613764565b6123ac826138d7565b8051600283018054604084015180865560208501517effffffffff00000000000000000000000000000000000000000000000000009092167f0100000000000000000000000000000000000000000000000000000000000000941515949094027fffffffffffff0000000000000000000000000000000000000000000000000000169390931779ffffffffffffffffffffffffffffffffffffffffffffffffffff90911617905560018301546124629190613a2b565b60018301556040805182511515815260208084015179ffffffffffffffffffffffffffffffffffffffffffffffffffff169082015282820151918101919091527f44a2350342338075ac038f37b8d9e49e696e360492cb44cc6bc37fc117f19df8906060015b60405180910390a15050565b6040805180820190915260008082526020820152600082900361253557506040805180820190915267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152600060208201526111b5565b7f97a657c900000000000000000000000000000000000000000000000000000000612564600460008587614c88565b61256d91614cb2565b7fffffffff00000000000000000000000000000000000000000000000000000000161461262a577f97a657c9000000000000000000000000000000000000000000000000000000006125c3600460008587614c88565b6125cc91614cb2565b6040517f55a0e02c0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000928316600482015291166024820152604401610b4a565b6126378260048186614c88565b810190611c999190614cfa565b60005461010090046001600160a01b03163314610d8f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610b4a565b6000611c99836001600160a01b038416613a41565b6000611c99836001600160a01b038416613a4d565b6000611c99836001600160a01b038416613a59565b80516001600160a01b0316158061271e575060208101516001600160a01b0316155b80612734575060a08101516001600160a01b0316155b1561276b576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600680547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316179055602080830151600780546040808701516060808901519588167fffffffffffffffff000000000000000000000000000000000000000000000000909416939093177401000000000000000000000000000000000000000063ffffffff909216919091021777ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff95861602179092556080808701516008805460a08a015161ffff9093167fffffffffffffffffffff00000000000000000000000000000000000000000000909116176201000092891692909202919091179055825190810183527f000000000000000000000000000000000000000000000000000000000000000090951685527f00000000000000000000000000000000000000000000000000000000000000008316938501939093527f00000000000000000000000000000000000000000000000000000000000000008216848201527f000000000000000000000000000000000000000000000000000000000000000090911691830191909152517ff7736cb9ac6dabe6bc85a757f7dd584df95f4e92afdebd0cb6994b4aef7465349161130e918490614d26565b60005460ff166129d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610b4a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526000826060015164ffffffffff1642612a849190614dd6565b83516040850151919250612acc91612ab89079ffffffffffffffffffffffffffffffffffffffffffffffffffff168461483a565b8560200151612ac79190614822565b613a2b565b6020840152505064ffffffffff4216606082015290565b60005b8251811015612b74576000838281518110612b0357612b036148b2565b60200260200101519050612b2181601061227290919063ffffffff16565b15612b63576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b50612b6d816148e1565b9050612ae6565b5060005b8151811015610d69576000828281518110612b9557612b956148b2565b6020026020010151905060006001600160a01b0316816001600160a01b031603612bbf5750612c0e565b612bca601082612287565b15612c0c576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b612c17816148e1565b9050612b78565b60005b8151811015612db3576040518060600160405280838381518110612c4757612c476148b2565b6020026020010151604001516bffffffffffffffffffffffff168152602001838381518110612c7857612c786148b2565b60200260200101516020015167ffffffffffffffff168152602001838381518110612ca557612ca56148b2565b60200260200101516060015163ffffffff1681525060126000848481518110612cd057612cd06148b2565b602090810291909101810151516001600160a01b031682528181019290925260409081016000208351815493850151949092015163ffffffff1674010000000000000000000000000000000000000000027fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff9095166c01000000000000000000000000027fffffffffffffffffffffffff00000000000000000000000000000000000000009094166bffffffffffffffffffffffff909316929092179290921792909216919091179055612dac816148e1565b9050612c21565b507f1026720e82621bb6d100957028cafff5ccf24add051c05d1bd65fa1d71e0fae68160405161130e9190614ded565b6014546c01000000000000000000000000900463ffffffff1615801590612e30575060145463ffffffff6c010000000000000000000000008204166bffffffffffffffffffffffff909116115b15612e3d57612e3d611eb6565b80516040811115612e7a576040517fb5a10cfa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60096000818181818181612e8e8282613e74565b50505050505050506000805b82811015612f3657612ef6848281518110612eb757612eb76148b2565b602002602001015160000151858381518110612ed557612ed56148b2565b60200260200101516020015161ffff1660096122469092919063ffffffff16565b50838181518110612f0957612f096148b2565b60200260200101516020015161ffff1682612f249190614e78565b9150612f2f816148e1565b9050612e9a565b50601480547fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd2490612fa89083908690614e97565b60405180910390a1505050565b60005460ff1615613022576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610b4a565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612a233390565b6006546001600160a01b031633146130b9576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0381166130f9576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075474010000000000000000000000000000000000000000900463ffffffff1684111561317d576007546040517f869337890000000000000000000000000000000000000000000000000000000081527401000000000000000000000000000000000000000090910463ffffffff16600482015260248101859052604401610b4a565b6007547801000000000000000000000000000000000000000000000000900467ffffffffffffffff168311156131df576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600854825161ffff9091161015613222576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600f5460ff16801561323c575061323a601082613a65565b155b1561327e576040517fd0d259760000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401610b4a565b6007546111569083906001600160a01b0316613a87565b60008060001b828460200151856080015186606001518760e00151886101000151805190602001208961012001516040516020016132d39190614eb6565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d604001516040516020016133819c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b60006111b5826122b1565b60008080806122cb86866122bc565b336001600160a01b0382160361342b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610b4a565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000828152600284016020526040812082905561223e848461235a565b600081815260018301602052604081205480156135a05760006134db600183614dd6565b85549091506000906134ef90600190614dd6565b905081811461355457600086600001828154811061350f5761350f6148b2565b9060005260206000200154905080876000018481548110613532576135326148b2565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061356557613565614ec9565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506111b5565b60009150506111b5565b60008181526001830160205260408120546135f1575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556111b5565b5060006111b5565b6000611c998383612366565b60006111b58261237e565b6000808061361e8585612388565b600081815260029690960160205260409095205494959350505050565b6000613690826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166123949092919063ffffffff16565b805190915015610d6957808060200190518101906136ae9190614758565b610d69576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610b4a565b6000826000018281548110613751576137516148b2565b9060005260206000200154905092915050565b6060824710156137f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610b4a565b843b61385e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610b4a565b600080866001600160a01b0316858760405161387a9190614ef8565b60006040518083038185875af1925050503d80600081146138b7576040519150601f19603f3d011682016040523d82523d6000602084013e6138bc565b606091505b50915091506138cc828286613c18565b979650505050505050565b8054600182015414806139135750600281015464ffffffffff7a0100000000000000000000000000000000000000000000000000009091041642145b1561391b5750565b80546001820154111561395a576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002810154600090613991907a010000000000000000000000000000000000000000000000000000900464ffffffffff1642614dd6565b825460028401549192506139d4916139c59079ffffffffffffffffffffffffffffffffffffffffffffffffffff168461483a565b8460010154612ac79190614822565b60018301555060020180547fff0000000000ffffffffffffffffffffffffffffffffffffffffffffffffffff167a0100000000000000000000000000000000000000000000000000004264ffffffffff1602179055565b6000818310613a3a5781611c99565b5090919050565b6000611c9983836135f9565b6000611c998383613c6b565b6000611c998383613cf5565b6001600160a01b03811660009081526001830160205260408120541515611c99565b81516000805b82811015613c0c576000846001600160a01b031663d02641a0878481518110613ab857613ab86148b2565b6020908102919091010151516040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b0390911660048201526024016040805180830381865afa158015613b1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b439190614f14565b5177ffffffffffffffffffffffffffffffffffffffffffffffff1690506000819003613bc557858281518110613b7b57613b7b6148b2565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b4a565b80868381518110613bd857613bd86148b2565b602002602001015160200151613bee919061483a565b613bf89084614822565b92505080613c05906148e1565b9050613a8d565b50611156600382613d12565b60608315613c27575081611c99565b825115613c375782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4a9190613f66565b600081815260028301602052604081205480151580613c8f5750613c8f84846135f9565b611c99576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b657900006044820152606401610b4a565b60008181526002830160205260408120819055611c998383613e68565b60028201547f0100000000000000000000000000000000000000000000000000000000000000900460ff161580613d47575080155b15613d50575050565b613d59826138d7565b8154811115613da15781546040517f48369c43000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610b4a565b8082600101541015613e1f576002820154600183015479ffffffffffffffffffffffffffffffffffffffffffffffffffff90911690613de09083614dd6565b613dea9190614877565b6040517fdc96cefa000000000000000000000000000000000000000000000000000000008152600401610b4a91815260200190565b80826001016000828254613e339190614dd6565b90915550506040518181527f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a906020016124c8565b6000611c9983836134b7565b50805460008255906000526020600020908101906108c291905b80821115613ea25760008155600101613e8e565b5090565b608081016111b582846001600160a01b038151168252602081015167ffffffffffffffff808216602085015280604084015116604085015280606084015116606085015250505050565b60005b83811015613f0b578181015183820152602001613ef3565b838111156111565750506000910152565b60008151808452613f34816020860160208601613ef0565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611c996020830184613f1c565b6001600160a01b03811681146108c257600080fd5b600060208284031215613fa057600080fd5b8135611c9981613f79565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613ffd57613ffd613fab565b60405290565b6040516080810167ffffffffffffffff81118282101715613ffd57613ffd613fab565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561406d5761406d613fab565b604052919050565b80151581146108c257600080fd5b60006060828403121561409557600080fd5b6040516060810181811067ffffffffffffffff821117156140b8576140b8613fab565b60405282356140c681614075565b8152602083013579ffffffffffffffffffffffffffffffffffffffffffffffffffff811681146140f557600080fd5b60208201526040928301359281019290925250919050565b600060a0828403121561154357600080fd5b60006020828403121561413157600080fd5b813567ffffffffffffffff81111561414857600080fd5b61223e8482850161410d565b600067ffffffffffffffff82111561416e5761416e613fab565b5060051b60200190565b600082601f83011261418957600080fd5b8135602061419e61419983614154565b614026565b82815260069290921b840181019181810190868411156141bd57600080fd5b8286015b8481101561420e57604081890312156141da5760008081fd5b6141e2613fda565b81356141ed81613f79565b8152818501356141fc81613f79565b818601528352918301916040016141c1565b509695505050505050565b6000806040838503121561422c57600080fd5b823567ffffffffffffffff8082111561424457600080fd5b61425086838701614178565b9350602085013591508082111561426657600080fd5b5061427385828601614178565b9150509250929050565b803563ffffffff8116811461429157600080fd5b919050565b67ffffffffffffffff811681146108c257600080fd5b803561ffff8116811461429157600080fd5b600060c082840312156142d057600080fd5b60405160c0810181811067ffffffffffffffff821117156142f3576142f3613fab565b604052823561430181613f79565b8152602083013561431181613f79565b60208201526143226040840161427d565b6040820152606083013561433581614296565b6060820152614346608084016142ac565b608082015260a083013561435981613f79565b60a08201529392505050565b6000806040838503121561437857600080fd5b823561438381613f79565b9150602083013561439381613f79565b809150509250929050565b60008083601f8401126143b057600080fd5b50813567ffffffffffffffff8111156143c857600080fd5b6020830191508360208260051b85010111156122d357600080fd5b600080600080604085870312156143f957600080fd5b843567ffffffffffffffff8082111561441157600080fd5b61441d8883890161439e565b9096509450602087013591508082111561443657600080fd5b506144438782880161439e565b95989497509550505050565b6000602080838503121561446257600080fd5b823567ffffffffffffffff81111561447957600080fd5b8301601f8101851361448a57600080fd5b803561449861419982614154565b81815260079190911b820183019083810190878311156144b757600080fd5b928401925b828410156138cc57608084890312156144d55760008081fd5b6144dd614003565b84356144e881613f79565b8152848601356144f781614296565b818701526040858101356bffffffffffffffffffffffff8116811461451c5760008081fd5b90820152606061452d86820161427d565b90820152825260809390930192908401906144bc565b60c081016111b582846001600160a01b0380825116835280602083015116602084015263ffffffff604083015116604084015267ffffffffffffffff606083015116606084015261ffff60808301511660808401528060a08301511660a0840152505050565b600080602083850312156145bc57600080fd5b823567ffffffffffffffff808211156145d457600080fd5b818501915085601f8301126145e857600080fd5b8135818111156145f757600080fd5b8660208260061b850101111561460c57600080fd5b60209290920196919550909350505050565b6020808252825182820181905260009190848201906040850190845b8181101561465f5783516001600160a01b03168352928401929184019160010161463a565b50909695505050505050565b60008060006060848603121561468057600080fd5b833567ffffffffffffffff81111561469757600080fd5b6146a38682870161410d565b9350506020840135915060408401356146bb81613f79565b809150509250925092565b600081518084526020808501945080840160005b8381101561470e57815180516001600160a01b0316885283015161ffff1683880152604090960195908201906001016146da565b509495945050505050565b60408152600061472c60408301856146c6565b90508260208301529392505050565b60006020828403121561474d57600080fd5b8135611c9981614075565b60006020828403121561476a57600080fd5b8151611c9981614075565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126147aa57600080fd5b83018035915067ffffffffffffffff8211156147c557600080fd5b6020019150368190038213156122d357600080fd5b6000602082840312156147ec57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115614835576148356147f3565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614872576148726147f3565b500290565b6000826148ad577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614912576149126147f3565b5060010190565b600067ffffffffffffffff80831681851680830382111561493c5761493c6147f3565b01949350505050565b60006040828403121561495757600080fd5b61495f613fda565b823561496a81613f79565b8152614978602084016142ac565b60208201529392505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126149b957600080fd5b83018035915067ffffffffffffffff8211156149d457600080fd5b6020019150600681901b36038213156122d357600080fd5b6000604082840312156149fe57600080fd5b614a06613fda565b8235614a1181613f79565b81526020928301359281019290925250919050565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b600060208284031215614a8557600080fd5b5035919050565b60006bffffffffffffffffffffffff80831681851680830382111561493c5761493c6147f3565b600067ffffffffffffffff808316818103614ad057614ad06147f3565b6001019392505050565b600081518084526020808501945080840160005b8381101561470e57815180516001600160a01b031688528301518388015260409096019590820190600101614aee565b60208152614b3960208201835167ffffffffffffffff169052565b60006020830151614b56604084018267ffffffffffffffff169052565b50604083015160608301526060830151614b7b60808401826001600160a01b03169052565b50608083015167ffffffffffffffff811660a08401525060a083015160c083015260c0830151614baf60e084018215159052565b5060e0830151610100614bcc818501836001600160a01b03169052565b808501519150506101806101208181860152614bec6101a0860184613f1c565b92508086015190506101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08685030181870152614c2a8483614ada565b935080870151915050610160614c4a818701836001600160a01b03169052565b959095015193019290925250919050565b60006bffffffffffffffffffffffff83811690831681811015614c8057614c806147f3565b039392505050565b60008085851115614c9857600080fd5b83861115614ca557600080fd5b5050820193919092039150565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015614cf25780818660040360031b1b83161692505b505092915050565b600060408284031215614d0c57600080fd5b614d14613fda565b82358152602083013561497881614075565b6101408101614d7182856001600160a01b038151168252602081015167ffffffffffffffff808216602085015280604084015116604085015280606084015116606085015250505050565b611c9960808301846001600160a01b0380825116835280602083015116602084015263ffffffff604083015116604084015267ffffffffffffffff606083015116606084015261ffff60808301511660808401528060a08301511660a0840152505050565b600082821015614de857614de86147f3565b500390565b602080825282518282018190526000919060409081850190868401855b82811015614e6b57815180516001600160a01b031685528681015167ffffffffffffffff1687860152858101516bffffffffffffffffffffffff168686015260609081015163ffffffff169085015260809093019290850190600101614e0a565b5091979650505050505050565b600063ffffffff80831681851680830382111561493c5761493c6147f3565b63ffffffff8316815260406020820152600061223e60408301846146c6565b602081526000611c996020830184614ada565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008251614f0a818460208701613ef0565b9190910192915050565b600060408284031215614f2657600080fd5b614f2e613fda565b825177ffffffffffffffffffffffffffffffffffffffffffffffff81168114614f5657600080fd5b815260208301516149788161429656fea164736f6c634300080f000a",
}

var EVM2EVMOnRampABI = EVM2EVMOnRampMetaData.ABI

var EVM2EVMOnRampBin = EVM2EVMOnRampMetaData.Bin

func DeployEVM2EVMOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMOnRampStaticConfig, dynamicConfig EVM2EVMOnRampDynamicConfig, tokensAndPools []EVM2EVMOnRampTokenAndPool, allowlist []common.Address, rateLimiterConfig RateLimiterConfig, feeTokenConfigs []EVM2EVMOnRampFeeTokenConfigArgs, nopsAndWeights []EVM2EVMOnRampNopAndWeight) (common.Address, *types.Transaction, *EVM2EVMOnRamp, error) {
	parsed, err := EVM2EVMOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMOnRampBin), backend, staticConfig, dynamicConfig, tokensAndPools, allowlist, rateLimiterConfig, feeTokenConfigs, nopsAndWeights)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetAllowList() ([]common.Address, error) {
	return _EVM2EVMOnRamp.Contract.GetAllowList(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetAllowList() ([]common.Address, error) {
	return _EVM2EVMOnRamp.Contract.GetAllowList(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetAllowListEnabled() (bool, error) {
	return _EVM2EVMOnRamp.Contract.GetAllowListEnabled(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetAllowListEnabled() (bool, error) {
	return _EVM2EVMOnRamp.Contract.GetAllowListEnabled(&_EVM2EVMOnRamp.CallOpts)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetFeeConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMOnRampFeeTokenConfig, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getFeeConfig", token)

	if err != nil {
		return *new(EVM2EVMOnRampFeeTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMOnRampFeeTokenConfig)).(*EVM2EVMOnRampFeeTokenConfig)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetFeeConfig(token common.Address) (EVM2EVMOnRampFeeTokenConfig, error) {
	return _EVM2EVMOnRamp.Contract.GetFeeConfig(&_EVM2EVMOnRamp.CallOpts, token)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetFeeConfig(token common.Address) (EVM2EVMOnRampFeeTokenConfig, error) {
	return _EVM2EVMOnRamp.Contract.GetFeeConfig(&_EVM2EVMOnRamp.CallOpts, token)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMOnRamp.Contract.IsAFNHealthy(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMOnRamp.Contract.IsAFNHealthy(&_EVM2EVMOnRamp.CallOpts)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) Paused() (bool, error) {
	return _EVM2EVMOnRamp.Contract.Paused(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMOnRamp.Contract.Paused(&_EVM2EVMOnRamp.CallOpts)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.ApplyAllowListUpdates(&_EVM2EVMOnRamp.TransactOpts, removes, adds)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.ApplyAllowListUpdates(&_EVM2EVMOnRamp.TransactOpts, removes, adds)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.Pause(&_EVM2EVMOnRamp.TransactOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.Pause(&_EVM2EVMOnRamp.TransactOpts)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetAllowListEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setAllowListEnabled", enabled)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetAllowListEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetAllowListEnabled(&_EVM2EVMOnRamp.TransactOpts, enabled)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetAllowListEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetAllowListEnabled(&_EVM2EVMOnRamp.TransactOpts, enabled)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetFeeConfig(opts *bind.TransactOpts, feeTokenConfigs []EVM2EVMOnRampFeeTokenConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setFeeConfig", feeTokenConfigs)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetFeeConfig(feeTokenConfigs []EVM2EVMOnRampFeeTokenConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetFeeConfig(&_EVM2EVMOnRamp.TransactOpts, feeTokenConfigs)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetFeeConfig(feeTokenConfigs []EVM2EVMOnRampFeeTokenConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetFeeConfig(&_EVM2EVMOnRamp.TransactOpts, feeTokenConfigs)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.TransferOwnership(&_EVM2EVMOnRamp.TransactOpts, to)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.TransferOwnership(&_EVM2EVMOnRamp.TransactOpts, to)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.Unpause(&_EVM2EVMOnRamp.TransactOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.Unpause(&_EVM2EVMOnRamp.TransactOpts)
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

type EVM2EVMOnRampAllowListAddIterator struct {
	Event *EVM2EVMOnRampAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampAllowListAdd)
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
		it.Event = new(EVM2EVMOnRampAllowListAdd)
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

func (it *EVM2EVMOnRampAllowListAddIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*EVM2EVMOnRampAllowListAddIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampAllowListAddIterator{contract: _EVM2EVMOnRamp.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampAllowListAdd)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseAllowListAdd(log types.Log) (*EVM2EVMOnRampAllowListAdd, error) {
	event := new(EVM2EVMOnRampAllowListAdd)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampAllowListEnabledSetIterator struct {
	Event *EVM2EVMOnRampAllowListEnabledSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampAllowListEnabledSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampAllowListEnabledSet)
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
		it.Event = new(EVM2EVMOnRampAllowListEnabledSet)
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

func (it *EVM2EVMOnRampAllowListEnabledSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampAllowListEnabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampAllowListEnabledSet struct {
	Enabled bool
	Raw     types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterAllowListEnabledSet(opts *bind.FilterOpts) (*EVM2EVMOnRampAllowListEnabledSetIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "AllowListEnabledSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampAllowListEnabledSetIterator{contract: _EVM2EVMOnRamp.contract, event: "AllowListEnabledSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchAllowListEnabledSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampAllowListEnabledSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "AllowListEnabledSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampAllowListEnabledSet)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "AllowListEnabledSet", log); err != nil {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseAllowListEnabledSet(log types.Log) (*EVM2EVMOnRampAllowListEnabledSet, error) {
	event := new(EVM2EVMOnRampAllowListEnabledSet)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "AllowListEnabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampAllowListRemoveIterator struct {
	Event *EVM2EVMOnRampAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampAllowListRemove)
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
		it.Event = new(EVM2EVMOnRampAllowListRemove)
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

func (it *EVM2EVMOnRampAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*EVM2EVMOnRampAllowListRemoveIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampAllowListRemoveIterator{contract: _EVM2EVMOnRamp.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampAllowListRemove)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseAllowListRemove(log types.Log) (*EVM2EVMOnRampAllowListRemove, error) {
	event := new(EVM2EVMOnRampAllowListRemove)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

type EVM2EVMOnRampPausedIterator struct {
	Event *EVM2EVMOnRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampPaused)
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
		it.Event = new(EVM2EVMOnRampPaused)
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

func (it *EVM2EVMOnRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMOnRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampPausedIterator{contract: _EVM2EVMOnRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampPaused)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParsePaused(log types.Log) (*EVM2EVMOnRampPaused, error) {
	event := new(EVM2EVMOnRampPaused)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

type EVM2EVMOnRampUnpausedIterator struct {
	Event *EVM2EVMOnRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampUnpaused)
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
		it.Event = new(EVM2EVMOnRampUnpaused)
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

func (it *EVM2EVMOnRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMOnRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampUnpausedIterator{contract: _EVM2EVMOnRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampUnpaused)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMOnRampUnpaused, error) {
	event := new(EVM2EVMOnRampUnpaused)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
	case _EVM2EVMOnRamp.abi.Events["AllowListAdd"].ID:
		return _EVM2EVMOnRamp.ParseAllowListAdd(log)
	case _EVM2EVMOnRamp.abi.Events["AllowListEnabledSet"].ID:
		return _EVM2EVMOnRamp.ParseAllowListEnabledSet(log)
	case _EVM2EVMOnRamp.abi.Events["AllowListRemove"].ID:
		return _EVM2EVMOnRamp.ParseAllowListRemove(log)
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
	case _EVM2EVMOnRamp.abi.Events["Paused"].ID:
		return _EVM2EVMOnRamp.ParsePaused(log)
	case _EVM2EVMOnRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMOnRamp.ParsePoolAdded(log)
	case _EVM2EVMOnRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMOnRamp.ParsePoolRemoved(log)
	case _EVM2EVMOnRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMOnRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMOnRampAdminSet) Topic() common.Hash {
	return common.HexToHash("0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c")
}

func (EVM2EVMOnRampAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (EVM2EVMOnRampAllowListEnabledSet) Topic() common.Hash {
	return common.HexToHash("0xccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df032")
}

func (EVM2EVMOnRampAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (EVM2EVMOnRampCCIPSendRequested) Topic() common.Hash {
	return common.HexToHash("0xaffc45517195d6499808c643bd4a7b0ffeedf95bea5852840d7bfcf63f59e821")
}

func (EVM2EVMOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf7736cb9ac6dabe6bc85a757f7dd584df95f4e92afdebd0cb6994b4aef746534")
}

func (EVM2EVMOnRampFeeConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1026720e82621bb6d100957028cafff5ccf24add051c05d1bd65fa1d71e0fae6")
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

func (EVM2EVMOnRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMOnRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMOnRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMOnRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMOnRamp *EVM2EVMOnRamp) Address() common.Address {
	return _EVM2EVMOnRamp.address
}

type EVM2EVMOnRampInterface interface {
	CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMOnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetFee(opts *bind.CallOpts, message ClientEVM2AnyMessage) (*big.Int, error)

	GetFeeConfig(opts *bind.CallOpts, token common.Address) (EVM2EVMOnRampFeeTokenConfig, error)

	GetNopFeesJuels(opts *bind.CallOpts) (*big.Int, error)

	GetNops(opts *bind.CallOpts) (GetNops,

		error)

	GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMOnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyPoolUpdates(opts *bind.TransactOpts, removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	PayNops(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetAllowListEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMOnRampDynamicConfig) (*types.Transaction, error)

	SetFeeConfig(opts *bind.TransactOpts, feeTokenConfigs []EVM2EVMOnRampFeeTokenConfigArgs) (*types.Transaction, error)

	SetNops(opts *bind.TransactOpts, nopsAndWeights []EVM2EVMOnRampNopAndWeight) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawNonLinkFees(opts *bind.TransactOpts, feeToken common.Address, to common.Address) (*types.Transaction, error)

	FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMOnRampAdminSetIterator, error)

	WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampAdminSet) (event.Subscription, error)

	ParseAdminSet(log types.Log) (*EVM2EVMOnRampAdminSet, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*EVM2EVMOnRampAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*EVM2EVMOnRampAllowListAdd, error)

	FilterAllowListEnabledSet(opts *bind.FilterOpts) (*EVM2EVMOnRampAllowListEnabledSetIterator, error)

	WatchAllowListEnabledSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampAllowListEnabledSet) (event.Subscription, error)

	ParseAllowListEnabledSet(log types.Log) (*EVM2EVMOnRampAllowListEnabledSet, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*EVM2EVMOnRampAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*EVM2EVMOnRampAllowListRemove, error)

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

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMOnRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMOnRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMOnRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMOnRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMOnRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMOnRampPoolRemoved, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMOnRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMOnRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
