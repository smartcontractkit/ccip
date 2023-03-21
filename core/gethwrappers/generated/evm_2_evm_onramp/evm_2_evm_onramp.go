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
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated"
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

type EVM2EVMOnRampTokenAndPool struct {
	Token common.Address
	Pool  common.Address
}

type IAggregateRateLimiterRateLimiterConfig struct {
	Admin    common.Address
	Rate     *big.Int
	Capacity *big.Int
}

type IAggregateRateLimiterTokenBucket struct {
	Rate        *big.Int
	Capacity    *big.Int
	Tokens      *big.Int
	LastUpdated *big.Int
}

type IEVM2EVMOnRampDynamicConfig struct {
	Router          common.Address
	PriceRegistry   common.Address
	MaxDataSize     uint32
	MaxGasLimit     uint64
	MaxTokensLength uint16
	Afn             common.Address
}

type IEVM2EVMOnRampFeeTokenConfig struct {
	FeeAmount       *big.Int
	Multiplier      uint64
	DestGasOverhead uint32
}

type IEVM2EVMOnRampFeeTokenConfigArgs struct {
	Token           common.Address
	Multiplier      uint64
	FeeAmount       *big.Int
	DestGasOverhead uint32
}

type IEVM2EVMOnRampNopAndWeight struct {
	Nop    common.Address
	Weight uint16
}

type IEVM2EVMOnRampStaticConfig struct {
	LinkToken         common.Address
	ChainId           uint64
	DestChainId       uint64
	DefaultTxGasLimit uint64
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

var EVM2EVMOnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"}],\"internalType\":\"structIEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOnRamp.TokenAndPool[]\",\"name\":\"tokensAndPools\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"multiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structIEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structIEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"expected\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"got\",\"type\":\"bytes4\"}],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InvalidFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"InvalidWithdrawalAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeesToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNopsToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrFeeAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrFeeAdminOrNop\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chain\",\"type\":\"uint64\"}],\"name\":\"TokenOrChainNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"multiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeConfig\",\"type\":\"tuple[]\"}],\"name\":\"FeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NopPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nopWeightsTotal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"NopsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"}],\"name\":\"applyPoolUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"}],\"internalType\":\"structIEVM2EVMOnRamp.DynamicConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"multiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structIEVM2EVMOnRamp.FeeTokenConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNopFeesJuels\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNops\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structIEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"weightsTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAllowListEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"}],\"internalType\":\"structIEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"multiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structIEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structIEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"setNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawNonLinkFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b50604051620065a6380380620065a6833981016040819052620000359162001261565b6000805460ff1916815583903390819081620000985760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000d257620000d281620003d9565b50508151600280546001600160a01b0319166001600160a01b03928316179055604080516080810182526020808601516001600160d01b03168083529583018051918301829052519282018390524260609092018290526005959095556006949094556007556008929092555087511615806200015a575060208701516001600160401b0316155b8062000171575060408701516001600160401b0316155b8062000188575060608701516001600160401b0316155b15620001a7576040516306b7c75960e31b815260040160405180910390fd5b6020808801516040808a015181517fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a948101949094526001600160401b039283169184019190915216606082015230608082015260a00160408051601f19818403018152918152815160209283012060805288516001600160a01b031660c052908801516001600160401b0390811660e0529088015181166101005260608801511660a05262000257866200048a565b620002628262000687565b6200026d8162000811565b60005b85518110156200037e5760006001600160a01b03168682815181106200029a576200029a6200135c565b6020026020010151600001516001600160a01b03161480620002ee575060006001600160a01b0316868281518110620002d757620002d76200135c565b6020026020010151602001516001600160a01b0316145b156200030d576040516306b7c75960e31b815260040160405180910390fd5b6200036a8682815181106200032657620003266200135c565b6020026020010151600001518783815181106200034757620003476200135c565b602002602001015160200151600f6200094e60201b6200248c179092919060201c565b50620003768162001388565b905062000270565b50835115620003cc576012805460ff19166001179055620003cc846000604051908082528060200260200182016040528015620003c5578160200160208202803683370190505b506200097a565b50505050505050620014f3565b336001600160a01b03821603620004335760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200008f565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b03161580620004ad575060208101516001600160a01b0316155b80620004c4575060a08101516001600160a01b0316155b15620004e3576040516306b7c75960e31b815260040160405180910390fd5b8051600980546001600160a01b039283166001600160a01b0319909116179055602080830151600a80546040808701516060808901516001600160401b03908116600160c01b026001600160c01b0363ffffffff909416600160a01b026001600160c01b0319909616978a1697909717949094179190911694909417909255608080870151600b805460a0808b01518a1662010000026001600160b01b031990921661ffff909416939093171790558351918201845260c051909616815260e0518216948101949094526101005181168483015293519093169082015290517ff7736cb9ac6dabe6bc85a757f7dd584df95f4e92afdebd0cb6994b4aef746534916200067c91849082516001600160a01b0390811682526020808501516001600160401b03908116828501526040808701518216818601526060968701518216878601528551841660808087019190915292860151841660a0808701919091529086015163ffffffff1660c0860152958501511660e084015283015161ffff166101008301529190920151166101208201526101400190565b60405180910390a150565b60005b8151811015620007df576040518060600160405280838381518110620006b457620006b46200135c565b6020026020010151604001516001600160601b03168152602001838381518110620006e357620006e36200135c565b6020026020010151602001516001600160401b031681526020018383815181106200071257620007126200135c565b60200260200101516060015163ffffffff16815250601560008484815181106200074057620007406200135c565b602090810291909101810151516001600160a01b031682528181019290925260409081016000208351815493850151949092015163ffffffff16600160a01b0263ffffffff60a01b196001600160401b039095166c01000000000000000000000000026001600160a01b03199094166001600160601b03909316929092179290921792909216919091179055620007d78162001388565b90506200068a565b507f1026720e82621bb6d100957028cafff5ccf24add051c05d1bd65fa1d71e0fae6816040516200067c9190620013a4565b600c600081818181818162000827828262000cee565b50505050505050506000805b8251811015620008e6576200089e8382815181106200085657620008566200135c565b6020026020010151600001518483815181106200087757620008776200135c565b60200260200101516020015161ffff16600c62000af660201b620024aa179092919060201c565b50828181518110620008b457620008b46200135c565b60200260200101516020015161ffff1682620008d191906200142b565b9150620008de8162001388565b905062000833565b506017805463ffffffff60601b19166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd249062000942908390859062001456565b60405180910390a15050565b600062000972846001600160a01b0385168462000b0e602090811b620024c017901c565b949350505050565b60005b815181101562000a3557620009be828281518110620009a057620009a06200135c565b6020026020010151601362000b2660201b620024d61790919060201c565b507f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566828281518110620009f557620009f56200135c565b602002602001015160405162000a1a91906001600160a01b0391909116815260200190565b60405180910390a162000a2d8162001388565b90506200097d565b5060005b825181101562000af15762000a7a83828151811062000a5c5762000a5c6200135c565b6020026020010151601362000b4660201b620024eb1790919060201c565b507f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d883828151811062000ab15762000ab16200135c565b602002602001015160405162000ad691906001600160a01b0391909116815260200190565b60405180910390a162000ae98162001388565b905062000a39565b505050565b600062000972846001600160a01b0385168462000b5d565b60006200097284846001600160a01b03851662000b5d565b600062000b3d836001600160a01b03841662000b8a565b90505b92915050565b600062000b3d836001600160a01b03841662000c8e565b6000828152600284016020908152604082208390556200097290859085906200250062000ce0821b17901c565b6000818152600183016020526040812054801562000c8357600062000bb1600183620014c3565b855490915060009062000bc790600190620014c3565b905081811462000c3357600086600001828154811062000beb5762000beb6200135c565b906000526020600020015490508087600001848154811062000c115762000c116200135c565b6000918252602080832090910192909255918252600188019052604090208390555b855486908062000c475762000c47620014dd565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000b40565b600091505062000b40565b600081815260018301602052604081205462000cd75750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000b40565b50600062000b40565b600062000b3d838362000c8e565b508054600082559060005260206000209081019062000d0e919062000d11565b50565b5b8082111562000d28576000815560010162000d12565b5090565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b038111828210171562000d675762000d6762000d2c565b60405290565b604080519081016001600160401b038111828210171562000d675762000d6762000d2c565b604051601f8201601f191681016001600160401b038111828210171562000dbd5762000dbd62000d2c565b604052919050565b80516001600160a01b038116811462000ddd57600080fd5b919050565b80516001600160401b038116811462000ddd57600080fd5b60006080828403121562000e0d57600080fd5b62000e1762000d42565b905062000e248262000dc5565b815262000e346020830162000de2565b602082015262000e476040830162000de2565b604082015262000e5a6060830162000de2565b606082015292915050565b805163ffffffff8116811462000ddd57600080fd5b805161ffff8116811462000ddd57600080fd5b600060c0828403121562000ea057600080fd5b60405160c081016001600160401b038111828210171562000ec55762000ec562000d2c565b60405290508062000ed68362000dc5565b815262000ee66020840162000dc5565b602082015262000ef96040840162000e65565b604082015262000f0c6060840162000de2565b606082015262000f1f6080840162000e7a565b608082015262000f3260a0840162000dc5565b60a08201525092915050565b60006001600160401b0382111562000f5a5762000f5a62000d2c565b5060051b60200190565b600082601f83011262000f7657600080fd5b8151602062000f8f62000f898362000f3e565b62000d92565b82815260069290921b8401810191818101908684111562000faf57600080fd5b8286015b8481101562001005576040818903121562000fce5760008081fd5b62000fd862000d6d565b62000fe38262000dc5565b815262000ff285830162000dc5565b8186015283529183019160400162000fb3565b509695505050505050565b600082601f8301126200102257600080fd5b815160206200103562000f898362000f3e565b82815260059290921b840181019181810190868411156200105557600080fd5b8286015b8481101562001005576200106d8162000dc5565b835291830191830162001059565b6000606082840312156200108e57600080fd5b604051606081016001600160401b0381118282101715620010b357620010b362000d2c565b604052905080620010c48362000dc5565b815260208301516001600160d01b0381168114620010e157600080fd5b6020820152604092830151920191909152919050565b600082601f8301126200110957600080fd5b815160206200111c62000f898362000f3e565b82815260079290921b840181019181810190868411156200113c57600080fd5b8286015b848110156200100557608081890312156200115b5760008081fd5b6200116562000d42565b620011708262000dc5565b81526200117f85830162000de2565b818601526040828101516001600160601b0381168114620011a05760008081fd5b908201526060620011b383820162000e65565b9082015283529183019160800162001140565b600082601f830112620011d857600080fd5b81516020620011eb62000f898362000f3e565b82815260069290921b840181019181810190868411156200120b57600080fd5b8286015b848110156200100557604081890312156200122a5760008081fd5b6200123462000d6d565b6200123f8262000dc5565b81526200124e85830162000e7a565b818601528352918301916040016200120f565b6000806000806000806000610220888a0312156200127e57600080fd5b6200128a898962000dfa565b96506200129b8960808a0162000e8d565b6101408901519096506001600160401b0380821115620012ba57600080fd5b620012c88b838c0162000f64565b96506101608a0151915080821115620012e057600080fd5b620012ee8b838c0162001010565b9550620013008b6101808c016200107b565b94506101e08a01519150808211156200131857600080fd5b620013268b838c01620010f7565b93506102008a01519150808211156200133e57600080fd5b506200134d8a828b01620011c6565b91505092959891949750929550565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016200139d576200139d62001372565b5060010190565b602080825282518282018190526000919060409081850190868401855b828110156200141e57815180516001600160a01b03168552868101516001600160401b031687860152858101516001600160601b03168686015260609081015163ffffffff169085015260809093019290850190600101620013c1565b5091979650505050505050565b600063ffffffff8083168185168083038211156200144d576200144d62001372565b01949350505050565b6000604080830163ffffffff8616845260208281860152818651808452606087019150828801935060005b81811015620014b557845180516001600160a01b0316845284015161ffff1684840152938301939185019160010162001481565b509098975050505050505050565b600082821015620014d857620014d862001372565b500390565b634e487b7160e01b600052603160045260246000fd5b60805160a05160c05160e0516101005161501362001593600039600081816102d9015281816108ff015281816109c301526129260152600081816102a901528181611b9301526128fc01526000818161027a01528181611160015281816118ca0152818161197101528181612285015281816123a301526128d60152600081816103090152818161253f015261294d01526000611d7501526150136000f3fe608060405234801561001057600080fd5b50600436106102415760003560e01c80635d86f14111610145578063a7cd63b7116100bd578063d3c7c2c71161008c578063efeadb6d11610071578063efeadb6d146107c8578063eff7cc48146107db578063f2fde38b146107e357600080fd5b8063d3c7c2c7146107b5578063e0351e13146107bd57600080fd5b8063a7cd63b714610764578063a7d3e02f14610779578063b06d41bc1461078c578063b3a18a3e146107a257600080fd5b806379ba509711610114578063856c8247116100f9578063856c8247146106e05780638da5cb5b1461071357806390c2339b1461072957600080fd5b806379ba5097146106d05780638456cb59146106d857600080fd5b80635d86f141146105c057806368d8f142146105d35780637437ff9f146105e657806376f6ae76146106bd57600080fd5b80633f4ba83a116101d8578063549e946f116101a757806354c8a4f31161018c57806354c8a4f31461057d578063599f6431146105905780635c975abb146105b557600080fd5b8063549e946f1461054a57806354b714681461055d57600080fd5b80633f4ba83a146104ee5780634120fccd146104f65780634352fa9f146105175780634741062e1461052a57600080fd5b806338724a951161021457806338724a951461049257806339aa9264146104b35780633a87ac53146104c85780633b1a1403146104db57600080fd5b806306285c6914610246578063147809b314610350578063181f5a77146103685780631e442b55146103b1575b600080fd5b61033a60408051608081018252600080825260208201819052918101829052606081019190915260405180608001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16815250905090565b6040516103479190613de9565b60405180910390f35b6103586107f6565b6040519015158152602001610347565b6103a46040518060400160405280601381526020017f45564d3245564d4f6e52616d7020312e302e300000000000000000000000000081525081565b6040516103479190613ea9565b6104526103bf366004613ed1565b6040805160608082018352600080835260208084018290529284018190526001600160a01b03949094168452601582529282902082519384018352546bffffffffffffffffffffffff811684526c01000000000000000000000000810467ffffffffffffffff169184019190915274010000000000000000000000000000000000000000900463ffffffff169082015290565b6040805182516bffffffffffffffffffffffff16815260208084015167ffffffffffffffff16908201529181015163ffffffff1690820152606001610347565b6104a56104a0366004613f00565b610878565b604051908152602001610347565b6104c66104c1366004613ed1565b610adf565b005b6104c66104d63660046140c4565b610b21565b6104c66104e936600461416b565b610ddf565b6104c6610df3565b6104fe610e05565b60405167ffffffffffffffff9091168152602001610347565b6104c6610525366004614274565b610e39565b61053d61053836600461432f565b61108e565b6040516103479190614364565b6104c66105583660046143a8565b611156565b6017546040516bffffffffffffffffffffffff9091168152602001610347565b6104c661058b366004614426565b6112bb565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610347565b60005460ff16610358565b61059d6105ce366004613ed1565b611336565b6104c66105e13660046144ac565b611347565b6106b06040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c0810182526009546001600160a01b039081168252600a54808216602084015274010000000000000000000000000000000000000000810463ffffffff1693830193909352780100000000000000000000000000000000000000000000000090920467ffffffffffffffff166060820152600b5461ffff8116608083015262010000900490911660a082015290565b604051610347919061458c565b6104c66106cb3660046145f2565b6113c4565b6104c6611427565b6104c661152f565b6104fe6106ee366004613ed1565b6001600160a01b031660009081526016602052604090205467ffffffffffffffff1690565b60005461010090046001600160a01b031661059d565b61073161153f565b60405161034791908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b61076c6115df565b6040516103479190614667565b6104a56107873660046146a8565b611690565b610794611de8565b604051610347929190614756565b6104c66107b0366004614778565b611eec565b61076c612008565b60125460ff16610358565b6104c66107d6366004614810565b6120b3565b6104c661211a565b6104c66107f1366004613ed1565b61247b565b6000600960020160029054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561084e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610872919061482d565b15905090565b60008061089061088b608085018561484a565b61250c565b51600a549091506000906001600160a01b031663268e5d486108b86080870160608801613ed1565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166024820152604401602060405180830381865afa158015610944573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096891906148af565b9050806000036109f5576109826080850160608601613ed1565b6040517f102e3c280000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660248201526044015b60405180910390fd5b6000601581610a0a6080880160608901613ed1565b6001600160a01b031681526020808201929092526040908101600020815160608101835290546bffffffffffffffffffffffff811682526c01000000000000000000000000810467ffffffffffffffff1693820184905274010000000000000000000000000000000000000000900463ffffffff169181018290529250670de0b6b3a764000091908490610a9e90876148f7565b610aa8919061490f565b610ab2919061490f565b610abc919061494c565b8151610ad691906bffffffffffffffffffffffff166148f7565b95945050505050565b610ae761267c565b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b610b2961267c565b60005b8251811015610c7d576000838281518110610b4957610b49614987565b60200260200101516000015190506000848381518110610b6b57610b6b614987565b602002602001015160200151905060006001600160a01b0316826001600160a01b03161480610ba157506001600160a01b038116155b15610bd8576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610be3600f836126f5565b15610c1a576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c26600f838361248c565b50604080516001600160a01b038085168252831660208201527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505080610c76906149b6565b9050610b2c565b5060005b8151811015610dda576000828281518110610c9e57610c9e614987565b60200260200101516000015190506000838381518110610cc057610cc0614987565b6020026020010151602001519050610ce282600f6126f590919063ffffffff16565b610d23576040517f73913ebd0000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016109ec565b6001600160a01b038116610d38600f8461270a565b6001600160a01b031614610d78576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d83600f8361271f565b50604080516001600160a01b038085168252831660208201527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a1505080610dd3906149b6565b9050610c81565b505050565b610de761267c565b610df081612734565b50565b610dfb61267c565b610e036129a4565b565b601754600090610e3490700100000000000000000000000000000000900467ffffffffffffffff1660016149ee565b905090565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015610e7657506002546001600160a01b03163314155b15610ead576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610ee9576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60045460005b81811015610f43576003600060048381548110610f0e57610f0e614987565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610f3c816149b6565b9050610eef565b5060005b82811015611073576000858281518110610f6357610f63614987565b6020026020010151905060006001600160a01b0316816001600160a01b031603610fb9576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610fcb57610fcb614987565b602002602001015160036000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f8186848151811061103057611030614987565b602002602001015160405161105a9291906001600160a01b03929092168252602082015260400190565b60405180910390a15061106c816149b6565b9050610f47565b508351611087906004906020870190613d3d565b5050505050565b80516060908067ffffffffffffffff8111156110ac576110ac613f35565b6040519080825280602002602001820160405280156110d5578160200160208202803683370190505b50915060005b8181101561114f57600360008583815181106110f9576110f9614987565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000205483828151811061113457611134614987565b6020908102919091010152611148816149b6565b90506110db565b5050919050565b61115e61267c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316036111d4576040517feddf07f50000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016109ec565b6001600160a01b03811661121f576040517f21680a040000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016109ec565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526112b79082906001600160a01b038516906370a0823190602401602060405180830381865afa158015611282573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112a691906148af565b6001600160a01b0385169190612a78565b5050565b6112c361267c565b61133084848080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808802828101820190935287825290935087925086918291850190849080828437600092019190915250612af892505050565b50505050565b600061134182612c4b565b92915050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561138457506002546001600160a01b03163314155b156113bb576040517fdf938f4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610df081612ca4565b6113cc61267c565b6112b78282808060200260200160405190810160405280939291908181526020016000905b8282101561141d5761140e60408302860136819003810190614a1a565b815260200190600101906113f1565b5050505050612e69565b6001546001600160a01b0316331461149b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016109ec565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61153761267c565b610e03612fa4565b61156a6040518060800160405280600081526020016000815260200160008152602001600081525090565b604080516080810182526005548152600654602082015260075491810191909152600854606082018190526000906115a29042614a59565b602083015183519192506115ce916115ba908461490f565b84604001516115c991906148f7565b613064565b604083015250426060820152919050565b606060006115ed601361307a565b67ffffffffffffffff81111561160557611605613f35565b60405190808252806020026020018201604052801561162e578160200160208202803683370190505b50905060005b61163e601361307a565b81101561168a57611650601382613084565b82828151811061166257611662614987565b6001600160a01b0390921660209283029190910190910152611683816149b6565b9050611634565b50919050565b6000805460ff16156116fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016109ec565b600960020160029054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611754573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611778919061482d565b156117ae576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006117c061088b608087018761484a565b905061183b6117d2602087018761484a565b83519091506117e46040890189614a70565b808060200260200160405190810160405280939291908181526020016000905b828210156118305761182160408302860136819003810190614ad8565b81526020019060010190611804565b505050505086613090565b611845858061484a565b905060201461188c57611858858061484a565b6040517f370d875f0000000000000000000000000000000000000000000000000000000081526004016109ec929190614b12565b6000611898868061484a565b8101906118a59190614b5f565b90506001600160a01b038111156118c057611858868061484a565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166118fa6080880160608901613ed1565b6001600160a01b03160361195e576017805486919060009061192b9084906bffffffffffffffffffffffff16614b78565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff160217905550611a7f565b600a546001600160a01b031663a5817a957f00000000000000000000000000000000000000000000000000000000000000006119a060808a0160608b01613ed1565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03928316600482015291166024820152604481018890526064016020604051808303816000875af1158015611a0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a309190614b9f565b60178054600090611a509084906bffffffffffffffffffffffff16614b78565b92506101000a8154816bffffffffffffffffffffffff02191690836bffffffffffffffffffffffff1602179055505b60005b611a8f6040880188614a70565b9050811015611b82576000611aa76040890189614a70565b83818110611ab757611ab7614987565b905060400201803603810190611acd9190614ad8565b90506000611ade8260000151612c4b565b60208301516040517fe2e59b3e0000000000000000000000000000000000000000000000000000000081529192506001600160a01b0383169163e2e59b3e91611b3d918b906004019182526001600160a01b0316602082015260400190565b600060405180830381600087803b158015611b5757600080fd5b505af1158015611b6b573d6000803e3d6000fd5b50505050505080611b7b906149b6565b9050611a82565b5060006040518061018001604052807f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020016017601081819054906101000a900467ffffffffffffffff16611be390614bbc565b825467ffffffffffffffff9182166101009390930a8381029083021990911617909255825260208083018a90526001600160a01b03891660408085018290526000918252601690925290812080546060909401939092611c439116614bbc565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff16815260200184600001518152602001846020015115158152602001836001600160a01b03168152602001888060200190611cac919061484a565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250602001611cf360408a018a614a70565b808060200260200160405190810160405280939291908181526020016000905b82821015611d3f57611d3060408302860136819003810190614ad8565b81526020019060010190611d13565b5050509183525050602001611d5a60808a0160608b01613ed1565b6001600160a01b0316815260006020909101529050611d99817f00000000000000000000000000000000000000000000000000000000000000006132a2565b6101608201526040517faffc45517195d6499808c643bd4a7b0ffeedf95bea5852840d7bfcf63f59e82190611dcf908390614c27565b60405180910390a16101600151925050505b9392505050565b6060600080611df7600c6133ac565b90508067ffffffffffffffff811115611e1257611e12613f35565b604051908082528060200260200182016040528015611e5757816020015b6040805180820190915260008082526020820152815260200190600190039081611e305790505b50925060005b81811015611ec957600080611e73600c846133b7565b915091506040518060400160405280836001600160a01b031681526020018261ffff16815250868481518110611eab57611eab614987565b6020026020010181905250505080611ec2906149b6565b9050611e5d565b505060175491926c0100000000000000000000000090920463ffffffff16919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015611f2957506002546001600160a01b03163314155b15611f60576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611f6a60056133d5565b60408101516006819055602082015179ffffffffffffffffffffffffffffffffffffffffffffffffffff16600555600754611fa59190613064565b600755604081810151602080840151835192835279ffffffffffffffffffffffffffffffffffffffffffffffffffff16908201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b60606000612016600f613480565b67ffffffffffffffff81111561202e5761202e613f35565b604051908082528060200260200182016040528015612057578160200160208202803683370190505b50905060005b815181101561168a57612071600f8261348b565b5082828151811061208457612084614987565b60200260200101816001600160a01b03166001600160a01b031681525050806120ac906149b6565b905061205d565b6120bb61267c565b601280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168215159081179091556040519081527fccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df03290602001611ffd565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561215757506002546001600160a01b03163314155b801561216b5750612169600c3361349a565b155b156121a2576040517f63baed9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6017546c01000000000000000000000000900463ffffffff1660008190036121f6576040517f990e30bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6017546bffffffffffffffffffffffff1663ffffffff8216811015612247576040517f8d0f71d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526bffffffffffffffffffffffff8216907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa1580156122d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122f891906148af565b1015612330576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600061233d600c6133ac565b905060005b8181101561243857600080612358600c846133b7565b9092509050600063ffffffff881661237e836bffffffffffffffffffffffff8a1661490f565b612388919061494c565b90506123948187614d64565b95506123d86001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016846bffffffffffffffffffffffff8416612a78565b6040516bffffffffffffffffffffffff821681526001600160a01b038416907f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f9060200160405180910390a250505080612431906149b6565b9050612342565b5050601780547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff929092169190911790555050565b61248361267c565b610df0816134af565b60006124a2846001600160a01b038516846124c0565b949350505050565b60006124a2846001600160a01b03851684613590565b60006124a284846001600160a01b038516613590565b6000611de1836001600160a01b0384166135ad565b6000611de1836001600160a01b0384166136a0565b6000611de183836136a0565b6040805180820190915260008082526020820152600082900361256d57506040805180820190915267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815260006020820152611341565b7f97a657c90000000000000000000000000000000000000000000000000000000061259c600460008587614d91565b6125a591614dbb565b7fffffffff000000000000000000000000000000000000000000000000000000001614612662577f97a657c9000000000000000000000000000000000000000000000000000000006125fb600460008587614d91565b61260491614dbb565b6040517f55a0e02c0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000009283166004820152911660248201526044016109ec565b61266f8260048186614d91565b810190611de19190614e03565b60005461010090046001600160a01b03163314610e03576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016109ec565b6000611de1836001600160a01b0384166136ef565b6000611de1836001600160a01b0384166136fb565b6000611de1836001600160a01b038416613707565b80516001600160a01b03161580612756575060208101516001600160a01b0316155b8061276c575060a08101516001600160a01b0316155b156127a3576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600980547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316179055602080830151600a80546040808701516060808901519588167fffffffffffffffff000000000000000000000000000000000000000000000000909416939093177401000000000000000000000000000000000000000063ffffffff909216919091021777ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff9586160217909255608080870151600b805460a08a015161ffff9093167fffffffffffffffffffff00000000000000000000000000000000000000000000909116176201000092891692909202919091179055825190810183527f000000000000000000000000000000000000000000000000000000000000000090951685527f00000000000000000000000000000000000000000000000000000000000000008316938501939093527f00000000000000000000000000000000000000000000000000000000000000008216848201527f000000000000000000000000000000000000000000000000000000000000000090911691830191909152517ff7736cb9ac6dabe6bc85a757f7dd584df95f4e92afdebd0cb6994b4aef74653491611ffd918490614e2f565b60005460ff16612a10576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016109ec565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610dda908490613713565b60005b8151811015612ba157612b31828281518110612b1957612b19614987565b602002602001015160136124d690919063ffffffff16565b507f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566828281518110612b6557612b65614987565b6020026020010151604051612b8991906001600160a01b0391909116815260200190565b60405180910390a1612b9a816149b6565b9050612afb565b5060005b8251811015610dda57612bdb838281518110612bc357612bc3614987565b602002602001015160136124eb90919063ffffffff16565b507f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8838281518110612c0f57612c0f614987565b6020026020010151604051612c3391906001600160a01b0391909116815260200190565b60405180910390a1612c44816149b6565b9050612ba5565b6000612c58600f836126f5565b612c99576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016109ec565b611341600f8361270a565b60005b8151811015612e39576040518060600160405280838381518110612ccd57612ccd614987565b6020026020010151604001516bffffffffffffffffffffffff168152602001838381518110612cfe57612cfe614987565b60200260200101516020015167ffffffffffffffff168152602001838381518110612d2b57612d2b614987565b60200260200101516060015163ffffffff1681525060156000848481518110612d5657612d56614987565b602090810291909101810151516001600160a01b031682528181019290925260409081016000208351815493850151949092015163ffffffff1674010000000000000000000000000000000000000000027fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff9095166c01000000000000000000000000027fffffffffffffffffffffffff00000000000000000000000000000000000000009094166bffffffffffffffffffffffff909316929092179290921792909216919091179055612e32816149b6565b9050612ca7565b507f1026720e82621bb6d100957028cafff5ccf24add051c05d1bd65fa1d71e0fae681604051611ffd9190614edf565b600c6000818181818181612e7d8282613dba565b50505050505050506000805b8251811015612f2657612ee6838281518110612ea757612ea7614987565b602002602001015160000151848381518110612ec557612ec5614987565b60200260200101516020015161ffff16600c6124aa9092919063ffffffff16565b50828181518110612ef957612ef9614987565b60200260200101516020015161ffff1682612f149190614f6a565b9150612f1f816149b6565b9050612e89565b50601780547fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff166c0100000000000000000000000063ffffffff8416021790556040517f8c337bff38141c507abd25c547606bdde78fe8c12e941ab613f3a565fea6cd2490612f989083908590614f89565b60405180910390a15050565b60005460ff1615613011576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016109ec565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612a5b3390565b60008183106130735781611de1565b5090919050565b6000611341825490565b6000611de18383613812565b6009546001600160a01b031633146130d4576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038116613114576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a5474010000000000000000000000000000000000000000900463ffffffff1684111561319857600a546040517f869337890000000000000000000000000000000000000000000000000000000081527401000000000000000000000000000000000000000090910463ffffffff166004820152602481018590526044016109ec565b600a547801000000000000000000000000000000000000000000000000900467ffffffffffffffff168311156131fa576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b54825161ffff909116101561323d576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60125460ff168015613257575061325560138261383c565b155b15613299576040517fd0d259760000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016109ec565b6113308261385e565b60008060001b828460200151856080015186606001518760e00151886101000151805190602001208961012001516040516020016132e09190614fa8565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d6040015160405160200161338e9c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b600061134182613a5b565b60008080806133c68686613a66565b909450925050505b9250929050565b8060010154816002015414806133ee5750428160030154145b156133f65750565b806001015481600201541115613438576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081600301544261344a9190614a59565b6001830154835491925061347191613462908461490f565b84600201546115c991906148f7565b60028301555042600390910155565b6000611341826133ac565b60008080806133c686866133b7565b6000611de1836001600160a01b038416613a91565b336001600160a01b03821603613521576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016109ec565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600082815260028401602052604081208290556124a28484612500565b600081815260018301602052604081205480156136965760006135d1600183614a59565b85549091506000906135e590600190614a59565b905081811461364a57600086600001828154811061360557613605614987565b906000526020600020015490508087600001848154811061362857613628614987565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061365b5761365b614fbb565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050611341565b6000915050611341565b60008181526001830160205260408120546136e757508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155611341565b506000611341565b6000611de18383613a91565b6000611de18383613a9d565b6000611de18383613b27565b6000613768826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613b449092919063ffffffff16565b805190915015610dda5780806020019051810190613786919061482d565b610dda576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016109ec565b600082600001828154811061382957613829614987565b9060005260206000200154905092915050565b6001600160a01b03811660009081526001830160205260408120541515611de1565b6000805b825181101561395d5760006003600085848151811061388357613883614987565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600003613916578382815181106138cc576138cc614987565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201526024016109ec565b83828151811061392857613928614987565b6020026020010151602001518161393f919061490f565b61394990846148f7565b92505080613956906149b6565b9050613862565b5080156112b75761396e60056133d5565b6006548111156139b8576006546040517f688ccf770000000000000000000000000000000000000000000000000000000081526004810191909152602481018290526044016109ec565b600754811115613a11576005546007546139d29083614a59565b6139dc919061494c565b6040517fe31e0f320000000000000000000000000000000000000000000000000000000081526004016109ec91815260200190565b8060056002016000828254613a269190614a59565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba37890602001612f98565b60006113418261307a565b60008080613a748585613084565b600081815260029690960160205260409095205494959350505050565b6000611de18383613b53565b600081815260028301602052604081205480151580613ac15750613ac18484613a91565b611de1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b6579000060448201526064016109ec565b60008181526002830160205260408120819055611de18383613b6b565b60606124a28484600085613b77565b60008181526001830160205260408120541515611de1565b6000611de183836135ad565b606082471015613c09576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016109ec565b843b613c71576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016109ec565b600080866001600160a01b03168587604051613c8d9190614fea565b60006040518083038185875af1925050503d8060008114613cca576040519150601f19603f3d011682016040523d82523d6000602084013e613ccf565b606091505b5091509150613cdf828286613cea565b979650505050505050565b60608315613cf9575081611de1565b825115613d095782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ec9190613ea9565b828054828255906000526020600020908101928215613daa579160200282015b82811115613daa57825182547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03909116178255602090920191600190910190613d5d565b50613db6929150613dd4565b5090565b5080546000825590600052602060002090810190610df091905b5b80821115613db65760008155600101613dd5565b6080810161134182846001600160a01b038151168252602081015167ffffffffffffffff808216602085015280604084015116604085015280606084015116606085015250505050565b60005b83811015613e4e578181015183820152602001613e36565b838111156113305750506000910152565b60008151808452613e77816020860160208601613e33565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611de16020830184613e5f565b6001600160a01b0381168114610df057600080fd5b600060208284031215613ee357600080fd5b8135611de181613ebc565b600060a0828403121561168a57600080fd5b600060208284031215613f1257600080fd5b813567ffffffffffffffff811115613f2957600080fd5b6124a284828501613eee565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613f8757613f87613f35565b60405290565b6040516080810167ffffffffffffffff81118282101715613f8757613f87613f35565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613ff757613ff7613f35565b604052919050565b600067ffffffffffffffff82111561401957614019613f35565b5060051b60200190565b600082601f83011261403457600080fd5b8135602061404961404483613fff565b613fb0565b82815260069290921b8401810191818101908684111561406857600080fd5b8286015b848110156140b957604081890312156140855760008081fd5b61408d613f64565b813561409881613ebc565b8152818501356140a781613ebc565b8186015283529183019160400161406c565b509695505050505050565b600080604083850312156140d757600080fd5b823567ffffffffffffffff808211156140ef57600080fd5b6140fb86838701614023565b9350602085013591508082111561411157600080fd5b5061411e85828601614023565b9150509250929050565b803563ffffffff8116811461413c57600080fd5b919050565b803567ffffffffffffffff8116811461413c57600080fd5b803561ffff8116811461413c57600080fd5b600060c0828403121561417d57600080fd5b60405160c0810181811067ffffffffffffffff821117156141a0576141a0613f35565b60405282356141ae81613ebc565b815260208301356141be81613ebc565b60208201526141cf60408401614128565b60408201526141e060608401614141565b60608201526141f160808401614159565b608082015260a083013561420481613ebc565b60a08201529392505050565b600082601f83011261422157600080fd5b8135602061423161404483613fff565b82815260059290921b8401810191818101908684111561425057600080fd5b8286015b848110156140b957803561426781613ebc565b8352918301918301614254565b6000806040838503121561428757600080fd5b823567ffffffffffffffff8082111561429f57600080fd5b6142ab86838701614210565b93506020915081850135818111156142c257600080fd5b85019050601f810186136142d557600080fd5b80356142e361404482613fff565b81815260059190911b8201830190838101908883111561430257600080fd5b928401925b8284101561432057833582529284019290840190614307565b80955050505050509250929050565b60006020828403121561434157600080fd5b813567ffffffffffffffff81111561435857600080fd5b6124a284828501614210565b6020808252825182820181905260009190848201906040850190845b8181101561439c57835183529284019291840191600101614380565b50909695505050505050565b600080604083850312156143bb57600080fd5b82356143c681613ebc565b915060208301356143d681613ebc565b809150509250929050565b60008083601f8401126143f357600080fd5b50813567ffffffffffffffff81111561440b57600080fd5b6020830191508360208260051b85010111156133ce57600080fd5b6000806000806040858703121561443c57600080fd5b843567ffffffffffffffff8082111561445457600080fd5b614460888389016143e1565b9096509450602087013591508082111561447957600080fd5b50614486878288016143e1565b95989497509550505050565b6bffffffffffffffffffffffff81168114610df057600080fd5b600060208083850312156144bf57600080fd5b823567ffffffffffffffff8111156144d657600080fd5b8301601f810185136144e757600080fd5b80356144f561404482613fff565b81815260079190911b8201830190838101908783111561451457600080fd5b928401925b82841015613cdf57608084890312156145325760008081fd5b61453a613f8d565b843561454581613ebc565b8152614552858701614141565b8682015260408086013561456581614492565b908201526060614576868201614128565b9082015282526080939093019290840190614519565b60c0810161134182846001600160a01b0380825116835280602083015116602084015263ffffffff604083015116604084015267ffffffffffffffff606083015116606084015261ffff60808301511660808401528060a08301511660a0840152505050565b6000806020838503121561460557600080fd5b823567ffffffffffffffff8082111561461d57600080fd5b818501915085601f83011261463157600080fd5b81358181111561464057600080fd5b8660208260061b850101111561465557600080fd5b60209290920196919550909350505050565b6020808252825182820181905260009190848201906040850190845b8181101561439c5783516001600160a01b031683529284019291840191600101614683565b6000806000606084860312156146bd57600080fd5b833567ffffffffffffffff8111156146d457600080fd5b6146e086828701613eee565b9350506020840135915060408401356146f881613ebc565b809150509250925092565b600081518084526020808501945080840160005b8381101561474b57815180516001600160a01b0316885283015161ffff168388015260409096019590820190600101614717565b509495945050505050565b6040815260006147696040830185614703565b90508260208301529392505050565b60006060828403121561478a57600080fd5b6040516060810181811067ffffffffffffffff821117156147ad576147ad613f35565b60405282356147bb81613ebc565b8152602083013579ffffffffffffffffffffffffffffffffffffffffffffffffffff811681146147ea57600080fd5b60208201526040928301359281019290925250919050565b8015158114610df057600080fd5b60006020828403121561482257600080fd5b8135611de181614802565b60006020828403121561483f57600080fd5b8151611de181614802565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261487f57600080fd5b83018035915067ffffffffffffffff82111561489a57600080fd5b6020019150368190038213156133ce57600080fd5b6000602082840312156148c157600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561490a5761490a6148c8565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614947576149476148c8565b500290565b600082614982577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036149e7576149e76148c8565b5060010190565b600067ffffffffffffffff808316818516808303821115614a1157614a116148c8565b01949350505050565b600060408284031215614a2c57600080fd5b614a34613f64565b8235614a3f81613ebc565b8152614a4d60208401614159565b60208201529392505050565b600082821015614a6b57614a6b6148c8565b500390565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112614aa557600080fd5b83018035915067ffffffffffffffff821115614ac057600080fd5b6020019150600681901b36038213156133ce57600080fd5b600060408284031215614aea57600080fd5b614af2613f64565b8235614afd81613ebc565b81526020928301359281019290925250919050565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b600060208284031215614b7157600080fd5b5035919050565b60006bffffffffffffffffffffffff808316818516808303821115614a1157614a116148c8565b600060208284031215614bb157600080fd5b8151611de181614492565b600067ffffffffffffffff808316818103614bd957614bd96148c8565b6001019392505050565b600081518084526020808501945080840160005b8381101561474b57815180516001600160a01b031688528301518388015260409096019590820190600101614bf7565b60208152614c4260208201835167ffffffffffffffff169052565b60006020830151614c5f604084018267ffffffffffffffff169052565b50604083015160608301526060830151614c8460808401826001600160a01b03169052565b50608083015167ffffffffffffffff811660a08401525060a083015160c083015260c0830151614cb860e084018215159052565b5060e0830151610100614cd5818501836001600160a01b03169052565b808501519150506101806101208181860152614cf56101a0860184613e5f565b92508086015190506101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08685030181870152614d338483614be3565b935080870151915050610160614d53818701836001600160a01b03169052565b959095015193019290925250919050565b60006bffffffffffffffffffffffff83811690831681811015614d8957614d896148c8565b039392505050565b60008085851115614da157600080fd5b83861115614dae57600080fd5b5050820193919092039150565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015614dfb5780818660040360031b1b83161692505b505092915050565b600060408284031215614e1557600080fd5b614e1d613f64565b823581526020830135614a4d81614802565b6101408101614e7a82856001600160a01b038151168252602081015167ffffffffffffffff808216602085015280604084015116604085015280606084015116606085015250505050565b611de160808301846001600160a01b0380825116835280602083015116602084015263ffffffff604083015116604084015267ffffffffffffffff606083015116606084015261ffff60808301511660808401528060a08301511660a0840152505050565b602080825282518282018190526000919060409081850190868401855b82811015614f5d57815180516001600160a01b031685528681015167ffffffffffffffff1687860152858101516bffffffffffffffffffffffff168686015260609081015163ffffffff169085015260809093019290850190600101614efc565b5091979650505050505050565b600063ffffffff808316818516808303821115614a1157614a116148c8565b63ffffffff831681526040602082015260006124a26040830184614703565b602081526000611de16020830184614be3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008251614ffc818460208701613e33565b919091019291505056fea164736f6c634300080f000a",
}

var EVM2EVMOnRampABI = EVM2EVMOnRampMetaData.ABI

var EVM2EVMOnRampBin = EVM2EVMOnRampMetaData.Bin

func DeployEVM2EVMOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig IEVM2EVMOnRampStaticConfig, dynamicConfig IEVM2EVMOnRampDynamicConfig, tokensAndPools []EVM2EVMOnRampTokenAndPool, allowlist []common.Address, rateLimiterConfig IAggregateRateLimiterRateLimiterConfig, feeTokenConfigs []IEVM2EVMOnRampFeeTokenConfigArgs, nopsAndWeights []IEVM2EVMOnRampNopAndWeight) (common.Address, *types.Transaction, *EVM2EVMOnRamp, error) {
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
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMOnRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(IAggregateRateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(IAggregateRateLimiterTokenBucket)).(*IAggregateRateLimiterTokenBucket)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMOnRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMOnRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMOnRamp.CallOpts)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetDynamicConfig(opts *bind.CallOpts) (IEVM2EVMOnRampDynamicConfig, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(IEVM2EVMOnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IEVM2EVMOnRampDynamicConfig)).(*IEVM2EVMOnRampDynamicConfig)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetDynamicConfig() (IEVM2EVMOnRampDynamicConfig, error) {
	return _EVM2EVMOnRamp.Contract.GetDynamicConfig(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetDynamicConfig() (IEVM2EVMOnRampDynamicConfig, error) {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetFeeConfig(opts *bind.CallOpts, token common.Address) (IEVM2EVMOnRampFeeTokenConfig, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getFeeConfig", token)

	if err != nil {
		return *new(IEVM2EVMOnRampFeeTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IEVM2EVMOnRampFeeTokenConfig)).(*IEVM2EVMOnRampFeeTokenConfig)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetFeeConfig(token common.Address) (IEVM2EVMOnRampFeeTokenConfig, error) {
	return _EVM2EVMOnRamp.Contract.GetFeeConfig(&_EVM2EVMOnRamp.CallOpts, token)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetFeeConfig(token common.Address) (IEVM2EVMOnRampFeeTokenConfig, error) {
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

	outstruct.NopsAndWeights = *abi.ConvertType(out[0], new([]IEVM2EVMOnRampNopAndWeight)).(*[]IEVM2EVMOnRampNopAndWeight)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMOnRamp.Contract.GetPricesForTokens(&_EVM2EVMOnRamp.CallOpts, tokens)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMOnRamp.Contract.GetPricesForTokens(&_EVM2EVMOnRamp.CallOpts, tokens)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampCaller) GetStaticConfig(opts *bind.CallOpts) (IEVM2EVMOnRampStaticConfig, error) {
	var out []interface{}
	err := _EVM2EVMOnRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(IEVM2EVMOnRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IEVM2EVMOnRampStaticConfig)).(*IEVM2EVMOnRampStaticConfig)

	return out0, err

}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) GetStaticConfig() (IEVM2EVMOnRampStaticConfig, error) {
	return _EVM2EVMOnRamp.Contract.GetStaticConfig(&_EVM2EVMOnRamp.CallOpts)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampCallerSession) GetStaticConfig() (IEVM2EVMOnRampStaticConfig, error) {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, adds []common.Address, removes []common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "applyAllowListUpdates", adds, removes)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) ApplyAllowListUpdates(adds []common.Address, removes []common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.ApplyAllowListUpdates(&_EVM2EVMOnRamp.TransactOpts, adds, removes)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) ApplyAllowListUpdates(adds []common.Address, removes []common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.ApplyAllowListUpdates(&_EVM2EVMOnRamp.TransactOpts, adds, removes)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) ApplyPoolUpdates(opts *bind.TransactOpts, adds []InternalPoolUpdate, removes []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "applyPoolUpdates", adds, removes)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) ApplyPoolUpdates(adds []InternalPoolUpdate, removes []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.ApplyPoolUpdates(&_EVM2EVMOnRamp.TransactOpts, adds, removes)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) ApplyPoolUpdates(adds []InternalPoolUpdate, removes []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.ApplyPoolUpdates(&_EVM2EVMOnRamp.TransactOpts, adds, removes)
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetAllowListEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setAllowListEnabled", enabled)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetAllowListEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetAllowListEnabled(&_EVM2EVMOnRamp.TransactOpts, enabled)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetAllowListEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetAllowListEnabled(&_EVM2EVMOnRamp.TransactOpts, enabled)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig IEVM2EVMOnRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetDynamicConfig(dynamicConfig IEVM2EVMOnRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetDynamicConfig(&_EVM2EVMOnRamp.TransactOpts, dynamicConfig)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetDynamicConfig(dynamicConfig IEVM2EVMOnRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetDynamicConfig(&_EVM2EVMOnRamp.TransactOpts, dynamicConfig)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetFeeConfig(opts *bind.TransactOpts, feeTokenConfigs []IEVM2EVMOnRampFeeTokenConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setFeeConfig", feeTokenConfigs)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetFeeConfig(feeTokenConfigs []IEVM2EVMOnRampFeeTokenConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetFeeConfig(&_EVM2EVMOnRamp.TransactOpts, feeTokenConfigs)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetFeeConfig(feeTokenConfigs []IEVM2EVMOnRampFeeTokenConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetFeeConfig(&_EVM2EVMOnRamp.TransactOpts, feeTokenConfigs)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetNops(opts *bind.TransactOpts, nopsAndWeights []IEVM2EVMOnRampNopAndWeight) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setNops", nopsAndWeights)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetNops(nopsAndWeights []IEVM2EVMOnRampNopAndWeight) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetNops(&_EVM2EVMOnRamp.TransactOpts, nopsAndWeights)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetNops(nopsAndWeights []IEVM2EVMOnRampNopAndWeight) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetNops(&_EVM2EVMOnRamp.TransactOpts, nopsAndWeights)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetPrices(&_EVM2EVMOnRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetPrices(&_EVM2EVMOnRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetRateLimiterConfig(&_EVM2EVMOnRamp.TransactOpts, config)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetRateLimiterConfig(&_EVM2EVMOnRamp.TransactOpts, config)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMOnRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOnRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMOnRamp.TransactOpts, newAdmin)
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

type EVM2EVMOnRampConfigChangedIterator struct {
	Event *EVM2EVMOnRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampConfigChanged)
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
		it.Event = new(EVM2EVMOnRampConfigChanged)
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

func (it *EVM2EVMOnRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOnRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampConfigChangedIterator{contract: _EVM2EVMOnRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampConfigChanged)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMOnRampConfigChanged, error) {
	event := new(EVM2EVMOnRampConfigChanged)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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
	StaticConfig  IEVM2EVMOnRampStaticConfig
	DynamicConfig IEVM2EVMOnRampDynamicConfig
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
	FeeConfig []IEVM2EVMOnRampFeeTokenConfigArgs
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
	NopsAndWeights  []IEVM2EVMOnRampNopAndWeight
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

type EVM2EVMOnRampTokenPriceChangedIterator struct {
	Event *EVM2EVMOnRampTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampTokenPriceChanged)
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
		it.Event = new(EVM2EVMOnRampTokenPriceChanged)
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

func (it *EVM2EVMOnRampTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMOnRampTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampTokenPriceChangedIterator{contract: _EVM2EVMOnRamp.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampTokenPriceChanged)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMOnRampTokenPriceChanged, error) {
	event := new(EVM2EVMOnRampTokenPriceChanged)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOnRampTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMOnRampTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMOnRampTokensRemovedFromBucket)
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

func (it *EVM2EVMOnRampTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMOnRampTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampTokensRemovedFromBucketIterator{contract: _EVM2EVMOnRamp.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampTokensRemovedFromBucket)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMOnRampTokensRemovedFromBucket, error) {
	event := new(EVM2EVMOnRampTokensRemovedFromBucket)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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
	NopsAndWeights []IEVM2EVMOnRampNopAndWeight
	WeightsTotal   *big.Int
}

func (_EVM2EVMOnRamp *EVM2EVMOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMOnRamp.abi.Events["AllowListAdd"].ID:
		return _EVM2EVMOnRamp.ParseAllowListAdd(log)
	case _EVM2EVMOnRamp.abi.Events["AllowListEnabledSet"].ID:
		return _EVM2EVMOnRamp.ParseAllowListEnabledSet(log)
	case _EVM2EVMOnRamp.abi.Events["AllowListRemove"].ID:
		return _EVM2EVMOnRamp.ParseAllowListRemove(log)
	case _EVM2EVMOnRamp.abi.Events["CCIPSendRequested"].ID:
		return _EVM2EVMOnRamp.ParseCCIPSendRequested(log)
	case _EVM2EVMOnRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMOnRamp.ParseConfigChanged(log)
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
	case _EVM2EVMOnRamp.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMOnRamp.ParseTokenPriceChanged(log)
	case _EVM2EVMOnRamp.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMOnRamp.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMOnRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMOnRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
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

func (EVM2EVMOnRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
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

func (EVM2EVMOnRampTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMOnRampTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMOnRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMOnRamp *EVM2EVMOnRamp) Address() common.Address {
	return _EVM2EVMOnRamp.address
}

type EVM2EVMOnRampInterface interface {
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetDynamicConfig(opts *bind.CallOpts) (IEVM2EVMOnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetFee(opts *bind.CallOpts, message ClientEVM2AnyMessage) (*big.Int, error)

	GetFeeConfig(opts *bind.CallOpts, token common.Address) (IEVM2EVMOnRampFeeTokenConfig, error)

	GetNopFeesJuels(opts *bind.CallOpts) (*big.Int, error)

	GetNops(opts *bind.CallOpts) (GetNops,

		error)

	GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)

	GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (IEVM2EVMOnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, adds []common.Address, removes []common.Address) (*types.Transaction, error)

	ApplyPoolUpdates(opts *bind.TransactOpts, adds []InternalPoolUpdate, removes []InternalPoolUpdate) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	PayNops(opts *bind.TransactOpts) (*types.Transaction, error)

	SetAllowListEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig IEVM2EVMOnRampDynamicConfig) (*types.Transaction, error)

	SetFeeConfig(opts *bind.TransactOpts, feeTokenConfigs []IEVM2EVMOnRampFeeTokenConfigArgs) (*types.Transaction, error)

	SetNops(opts *bind.TransactOpts, nopsAndWeights []IEVM2EVMOnRampNopAndWeight) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error)

	SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawNonLinkFees(opts *bind.TransactOpts, feeToken common.Address, to common.Address) (*types.Transaction, error)

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

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOnRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMOnRampConfigChanged, error)

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

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMOnRampTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMOnRampTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMOnRampTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMOnRampTokensRemovedFromBucket, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMOnRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMOnRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
