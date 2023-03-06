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
	FeeAdmin        common.Address
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
	Weight *big.Int
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOnRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"feeAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"}],\"internalType\":\"structIEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOnRamp.TokenAndPool[]\",\"name\":\"tokensAndPools\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"multiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structIEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structIEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"expected\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"got\",\"type\":\"bytes4\"}],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"InvalidFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"InvalidWithdrawalAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeesToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNopsToPay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrFeeAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwnerOrFeeAdminOrNop\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chain\",\"type\":\"uint64\"}],\"name\":\"TokenOrChainNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"feeAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMOnRamp.DynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"DynamicConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"multiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeConfig\",\"type\":\"tuple[]\"}],\"name\":\"FeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NopPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nopWeightsTotal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"NopsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"StaticConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"}],\"name\":\"applyPoolUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"feeAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"}],\"internalType\":\"structIEVM2EVMOnRamp.DynamicConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structClient.EVM2AnyMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFeeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"multiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structIEVM2EVMOnRamp.FeeTokenConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNopFeesJuels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNops\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structIEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"weightsTotal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOnRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAllowListEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"feeAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"}],\"internalType\":\"structIEVM2EVMOnRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"multiplier\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"feeAmount\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"destGasOverhead\",\"type\":\"uint32\"}],\"internalType\":\"structIEVM2EVMOnRamp.FeeTokenConfigArgs[]\",\"name\":\"feeTokenConfigs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structIEVM2EVMOnRamp.NopAndWeight[]\",\"name\":\"nopsAndWeights\",\"type\":\"tuple[]\"}],\"name\":\"setNops\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawNonLinkFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b50604051620060c7380380620060c783398101604081905262000035916200126f565b6000805460ff1916815583903390819081620000985760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000d257620000d2816200044e565b50508151600280546001600160a01b0319166001600160a01b03928316179055604080516080810182526020808601516001600160d01b03168083529583018051918301829052519282018390524260609092018290526005959095556006949094556007556008929092555087511615806200015a575060208701516001600160401b0316155b8062000171575060408701516001600160401b0316155b8062000188575060608701516001600160401b0316155b15620001a7576040516306b7c75960e31b815260040160405180910390fd5b6020808801516040808a015181517fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a948101949094526001600160401b039283169184019190915216606082015230608082015260a00160408051601f19818403018152828252805160209182012060805289516001600160a01b031660c0528901516001600160401b0390811660e0529089015181166101005260608901511660a0527f88ef61c99020a9e2d75ad86191c6f07f7b6ce93ccebda542554b2b1588204ae790620002b990899081516001600160a01b031681526020808301516001600160401b0390811691830191909152604080840151821690830152606092830151169181019190915260800190565b60405180910390a1620002cc86620004ff565b620002d782620006bb565b620002e28162000845565b60005b8551811015620003f35760006001600160a01b03168682815181106200030f576200030f6200136a565b6020026020010151600001516001600160a01b0316148062000363575060006001600160a01b03168682815181106200034c576200034c6200136a565b6020026020010151602001516001600160a01b0316145b1562000382576040516306b7c75960e31b815260040160405180910390fd5b620003df8682815181106200039b576200039b6200136a565b602002602001015160000151878381518110620003bc57620003bc6200136a565b60200260200101516020015160106200095860201b620023f9179092919060201c565b50620003eb8162001396565b9050620002e5565b5083511562000441576013805460ff19166001179055620004418460006040519080825280602002602001820160405280156200043a578160200160208202803683370190505b5062000984565b50505050505050620014e7565b336001600160a01b03821603620004a85760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200008f565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0316158062000522575060208101516001600160a01b0316155b8062000539575060c08101516001600160a01b0316155b1562000558576040516306b7c75960e31b815260040160405180910390fd5b8051600980546001600160a01b039283166001600160a01b0319918216179091556020830151600a805460408087015160608801516001600160401b0316600160c01b026001600160c01b0363ffffffff909216600160a01b026001600160c01b03199094169588169590951792909217919091169290921790556080840151600b805460a0870151861662010000026001600160b01b031990911661ffff9093169290921791909117905560c0840151600c805491909416921691909117909155517f0161b2e931b61fb04ae29bb225d36a3248c2b2abb7348385e3e74e5740490b8a90620006b090839081516001600160a01b03908116825260208084015182169083015260408084015163ffffffff16908301526060808401516001600160401b03169083015260808084015161ffff169083015260a08381015182169083015260c092830151169181019190915260e00190565b60405180910390a150565b60005b815181101562000813576040518060600160405280838381518110620006e857620006e86200136a565b6020026020010151604001516001600160601b031681526020018383815181106200071757620007176200136a565b6020026020010151602001516001600160401b031681526020018383815181106200074657620007466200136a565b60200260200101516060015163ffffffff16815250601660008484815181106200077457620007746200136a565b602090810291909101810151516001600160a01b031682528181019290925260409081016000208351815493850151949092015163ffffffff16600160a01b0263ffffffff60a01b196001600160401b039095166c01000000000000000000000000026001600160a01b03199094166001600160601b039093169290921792909217929092169190911790556200080b8162001396565b9050620006be565b507f1026720e82621bb6d100957028cafff5ccf24add051c05d1bd65fa1d71e0fae681604051620006b09190620013b2565b600d60008181818181816200085b828262000cf8565b50505050505050506000805b82518110156200091257620008ce8382815181106200088a576200088a6200136a565b602002602001015160000151848381518110620008ab57620008ab6200136a565b602002602001015160200151600d62000b0060201b62002417179092919060201c565b50828181518110620008e457620008e46200136a565b60200260200101516020015182620008fd919062001439565b91506200090a8162001396565b905062000867565b5060188190556040517f5c5cf993d6e35edd9d958829f41cc0961071d6f5ac714b1ce86fc44c73da1549906200094c908390859062001454565b60405180910390a15050565b60006200097c846001600160a01b0385168462000b18602090811b6200242d17901c565b949350505050565b60005b815181101562000a3f57620009c8828281518110620009aa57620009aa6200136a565b6020026020010151601462000b3060201b620024431790919060201c565b507f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566828281518110620009ff57620009ff6200136a565b602002602001015160405162000a2491906001600160a01b0391909116815260200190565b60405180910390a162000a378162001396565b905062000987565b5060005b825181101562000afb5762000a8483828151811062000a665762000a666200136a565b6020026020010151601462000b5060201b620024581790919060201c565b507f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d883828151811062000abb5762000abb6200136a565b602002602001015160405162000ae091906001600160a01b0391909116815260200190565b60405180910390a162000af38162001396565b905062000a43565b505050565b60006200097c846001600160a01b0385168462000b67565b60006200097c84846001600160a01b03851662000b67565b600062000b47836001600160a01b03841662000b94565b90505b92915050565b600062000b47836001600160a01b03841662000c98565b6000828152600284016020908152604082208390556200097c90859085906200246d62000cea821b17901c565b6000818152600183016020526040812054801562000c8d57600062000bbb600183620014b7565b855490915060009062000bd190600190620014b7565b905081811462000c3d57600086600001828154811062000bf55762000bf56200136a565b906000526020600020015490508087600001848154811062000c1b5762000c1b6200136a565b6000918252602080832090910192909255918252600188019052604090208390555b855486908062000c515762000c51620014d1565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062000b4a565b600091505062000b4a565b600081815260018301602052604081205462000ce15750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000b4a565b50600062000b4a565b600062000b47838362000c98565b508054600082559060005260206000209081019062000d18919062000d1b565b50565b5b8082111562000d32576000815560010162000d1c565b5090565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b038111828210171562000d715762000d7162000d36565b60405290565b60405160e081016001600160401b038111828210171562000d715762000d7162000d36565b604080519081016001600160401b038111828210171562000d715762000d7162000d36565b604051601f8201601f191681016001600160401b038111828210171562000dec5762000dec62000d36565b604052919050565b80516001600160a01b038116811462000e0c57600080fd5b919050565b80516001600160401b038116811462000e0c57600080fd5b60006080828403121562000e3c57600080fd5b62000e4662000d4c565b905062000e538262000df4565b815262000e636020830162000e11565b602082015262000e766040830162000e11565b604082015262000e896060830162000e11565b606082015292915050565b805163ffffffff8116811462000e0c57600080fd5b600060e0828403121562000ebc57600080fd5b62000ec662000d77565b905062000ed38262000df4565b815262000ee36020830162000df4565b602082015262000ef66040830162000e94565b604082015262000f096060830162000e11565b6060820152608082015161ffff8116811462000f2457600080fd5b608082015262000f3760a0830162000df4565b60a082015262000f4a60c0830162000df4565b60c082015292915050565b60006001600160401b0382111562000f715762000f7162000d36565b5060051b60200190565b600082601f83011262000f8d57600080fd5b8151602062000fa662000fa08362000f55565b62000dc1565b82815260069290921b8401810191818101908684111562000fc657600080fd5b8286015b848110156200101c576040818903121562000fe55760008081fd5b62000fef62000d9c565b62000ffa8262000df4565b81526200100985830162000df4565b8186015283529183019160400162000fca565b509695505050505050565b600082601f8301126200103957600080fd5b815160206200104c62000fa08362000f55565b82815260059290921b840181019181810190868411156200106c57600080fd5b8286015b848110156200101c57620010848162000df4565b835291830191830162001070565b600060608284031215620010a557600080fd5b604051606081016001600160401b0381118282101715620010ca57620010ca62000d36565b604052905080620010db8362000df4565b815260208301516001600160d01b0381168114620010f857600080fd5b6020820152604092830151920191909152919050565b600082601f8301126200112057600080fd5b815160206200113362000fa08362000f55565b82815260079290921b840181019181810190868411156200115357600080fd5b8286015b848110156200101c5760808189031215620011725760008081fd5b6200117c62000d4c565b620011878262000df4565b81526200119685830162000e11565b818601526040828101516001600160601b0381168114620011b75760008081fd5b908201526060620011ca83820162000e94565b9082015283529183019160800162001157565b600082601f830112620011ef57600080fd5b815160206200120262000fa08362000f55565b82815260069290921b840181019181810190868411156200122257600080fd5b8286015b848110156200101c5760408189031215620012415760008081fd5b6200124b62000d9c565b620012568262000df4565b8152818501518582015283529183019160400162001226565b6000806000806000806000610240888a0312156200128c57600080fd5b62001298898962000e29565b9650620012a98960808a0162000ea9565b6101608901519096506001600160401b0380821115620012c857600080fd5b620012d68b838c0162000f7b565b96506101808a0151915080821115620012ee57600080fd5b620012fc8b838c0162001027565b95506200130e8b6101a08c0162001092565b94506102008a01519150808211156200132657600080fd5b620013348b838c016200110e565b93506102208a01519150808211156200134c57600080fd5b506200135b8a828b01620011dd565b91505092959891949750929550565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201620013ab57620013ab62001380565b5060010190565b602080825282518282018190526000919060409081850190868401855b828110156200142c57815180516001600160a01b03168552868101516001600160401b031687860152858101516001600160601b03168686015260609081015163ffffffff169085015260809093019290850190600101620013cf565b5091979650505050505050565b600082198211156200144f576200144f62001380565b500190565b6000604080830185845260208281860152818651808452606087019150828801935060005b81811015620014a957845180516001600160a01b0316845284015184840152938301939185019160010162001479565b509098975050505050505050565b600082821015620014cc57620014cc62001380565b500390565b634e487b7160e01b600052603160045260246000fd5b60805160a05160c05160e05161010051614b5c6200156b600039600081816102d901528181610a190152610add0152600081816102a90152611b8f01526000818161027a0152818161125201528181611937015281816119a5015281816122300152612355015260008181610309015261260401526000611d710152614b5c6000f3fe608060405234801561001057600080fd5b50600436106102415760003560e01c80635d86f14111610145578063a7d3e02f116100bd578063d3c7c2c71161008c578063efeadb6d11610071578063efeadb6d14610870578063eff7cc4814610883578063f2fde38b1461088b57600080fd5b8063d3c7c2c71461085d578063e0351e131461086557600080fd5b8063a7d3e02f1461080e578063b06d41bc14610821578063b3a18a3e14610837578063cb18476d1461084a57600080fd5b80638456cb59116101145780638da5cb5b116100f95780638da5cb5b146107a857806390c2339b146107be578063a7cd63b7146107f957600080fd5b80638456cb591461076d578063856c82471461077557600080fd5b80635d86f141146105e957806368d8f142146105fc5780637437ff9f1461060f57806379ba50971461076557600080fd5b80633f4ba83a116101d8578063549e946f116101a757806354c8a4f31161018c57806354c8a4f3146105a6578063599f6431146105b95780635c975abb146105de57600080fd5b8063549e946f1461058b57806354b714681461059e57600080fd5b80633f4ba83a1461052f5780634120fccd146105375780634352fa9f146105585780634741062e1461056b57600080fd5b80631e442b55116102145780631e442b551461040757806338724a95146104e857806339aa9264146105095780633a87ac531461051c57600080fd5b806306285c6914610246578063147809b314610391578063181f5a77146103a95780631abacb7f146103f2575b600080fd5b61033a60408051608081018252600080825260208201819052918101829052606081019190915260405180608001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16815250905090565b604051610388919081516001600160a01b0316815260208083015167ffffffffffffffff90811691830191909152604080840151821690830152606092830151169181019190915260800190565b60405180910390f35b61039961089e565b6040519015158152602001610388565b6103e56040518060400160405280601381526020017f45564d3245564d4f6e52616d7020312e302e300000000000000000000000000081525081565b6040516103889190613c2a565b610405610400366004613c3d565b61092b565b005b6104a8610415366004613cd7565b6040805160608082018352600080835260208084018290529284018190526001600160a01b03949094168452601682529282902082519384018352546bffffffffffffffffffffffff811684526c01000000000000000000000000810467ffffffffffffffff169184019190915274010000000000000000000000000000000000000000900463ffffffff169082015290565b6040805182516bffffffffffffffffffffffff16815260208084015167ffffffffffffffff16908201529181015163ffffffff1690820152606001610388565b6104fb6104f6366004613d06565b610992565b604051908152602001610388565b610405610517366004613cd7565b610bf9565b61040561052a366004613ecf565b610c3b565b610405610ef9565b61053f610f0b565b60405167ffffffffffffffff9091168152602001610388565b610405610566366004613f97565b610f2b565b61057e610579366004614052565b611180565b6040516103889190614087565b6104056105993660046140cb565b611248565b6017546104fb565b6104056105b4366004614149565b6113a9565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610388565b60005460ff16610399565b6105c66105f7366004613cd7565b611424565b61040561060a3660046141e1565b611435565b6106ea6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152506040805160e0810182526009546001600160a01b039081168252600a54808216602084015274010000000000000000000000000000000000000000810463ffffffff1693830193909352780100000000000000000000000000000000000000000000000090920467ffffffffffffffff166060820152600b5461ffff81166080830152620100009004821660a0820152600c5490911660c082015290565b6040516103889190600060e0820190506001600160a01b0380845116835280602085015116602084015263ffffffff604085015116604084015267ffffffffffffffff606085015116606084015261ffff60808501511660808401528060a08501511660a08401528060c08501511660c08401525092915050565b6104056114bb565b6104056115a9565b61053f610783366004613cd7565b6001600160a01b031660009081526019602052604090205467ffffffffffffffff1690565b60005461010090046001600160a01b03166105c6565b6107c66115b9565b60405161038891908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b610801611659565b60405161038891906142d3565b6104fb61081c366004614314565b61170a565b610829611de4565b6040516103889291906143be565b6104056108453660046143e0565b611ecc565b61040561085836600461446a565b611fe8565b610801611ff9565b60135460ff16610399565b61040561087e366004614517565b6120a4565b6104056120ed565b610405610899366004613cd7565b6123e8565b600c54604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa158015610901573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109259190614534565b15905090565b610933612479565b61098e8282808060200260200160405190810160405280939291908181526020016000905b82821015610984576109756040830286013681900381019061458c565b81526020019060010190610958565b50505050506124d8565b5050565b6000806109aa6109a560808501856145a8565b6125d1565b51600a549091506000906001600160a01b031663268e5d486109d26080870160608801613cd7565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015267ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166024820152604401602060405180830381865afa158015610a5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a82919061460d565b905080600003610b0f57610a9c6080850160608601613cd7565b6040517f102e3c280000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660248201526044015b60405180910390fd5b6000601681610b246080880160608901613cd7565b6001600160a01b031681526020808201929092526040908101600020815160608101835290546bffffffffffffffffffffffff811682526c01000000000000000000000000810467ffffffffffffffff1693820184905274010000000000000000000000000000000000000000900463ffffffff169181018290529250670de0b6b3a764000091908490610bb89087614655565b610bc2919061466d565b610bcc919061466d565b610bd691906146aa565b8151610bf091906bffffffffffffffffffffffff16614655565b95945050505050565b610c01612479565b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b610c43612479565b60005b8251811015610d97576000838281518110610c6357610c636146e5565b60200260200101516000015190506000848381518110610c8557610c856146e5565b602002602001015160200151905060006001600160a01b0316826001600160a01b03161480610cbb57506001600160a01b038116155b15610cf2576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cfd601083612741565b15610d34576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d40601083836123f9565b50604080516001600160a01b038085168252831660208201527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505080610d9090614714565b9050610c46565b5060005b8151811015610ef4576000828281518110610db857610db86146e5565b60200260200101516000015190506000838381518110610dda57610dda6146e5565b6020026020010151602001519050610dfc82601061274190919063ffffffff16565b610e3d576040517f73913ebd0000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610b06565b6001600160a01b038116610e52601084612756565b6001600160a01b031614610e92576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e9d60108361276b565b50604080516001600160a01b038085168252831660208201527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a1505080610eed90614714565b9050610d9b565b505050565b610f01612479565b610f09612780565b565b601a54600090610f269067ffffffffffffffff16600161474c565b905090565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015610f6857506002546001600160a01b03163314155b15610f9f576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610fdb576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60045460005b81811015611035576003600060048381548110611000576110006146e5565b60009182526020808320909101546001600160a01b0316835282019290925260400181205561102e81614714565b9050610fe1565b5060005b82811015611165576000858281518110611055576110556146e5565b6020026020010151905060006001600160a01b0316816001600160a01b0316036110ab576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8482815181106110bd576110bd6146e5565b602002602001015160036000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110611122576111226146e5565b602002602001015160405161114c9291906001600160a01b03929092168252602082015260400190565b60405180910390a15061115e81614714565b9050611039565b508351611179906004906020870190613b26565b5050505050565b80516060908067ffffffffffffffff81111561119e5761119e613d3b565b6040519080825280602002602001820160405280156111c7578160200160208202803683370190505b50915060005b8181101561124157600360008583815181106111eb576111eb6146e5565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110611226576112266146e5565b602090810291909101015261123a81614714565b90506111cd565b5050919050565b611250612479565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316826001600160a01b0316036112c6576040517feddf07f50000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610b06565b6001600160a01b038116611311576040517f21680a040000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401610b06565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015261098e9082906001600160a01b038516906370a0823190602401602060405180830381865afa158015611374573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611398919061460d565b6001600160a01b038516919061281c565b6113b1612479565b61141e8484808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152505060408051602080880282810182019093528782529093508792508691829185019084908082843760009201919091525061289c92505050565b50505050565b600061142f826129ef565b92915050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b0316141580156114785750600b546201000090046001600160a01b03163314155b156114af576040517fdf938f4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114b881612a48565b50565b6001546001600160a01b031633146115155760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610b06565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6115b1612479565b610f09612c0d565b6115e46040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260055481526006546020820152600754918101919091526008546060820181905260009061161c9042614778565b6020830151835191925061164891611634908461466d565b84604001516116439190614655565b612c95565b604083015250426060820152919050565b606060006116676014612cab565b67ffffffffffffffff81111561167f5761167f613d3b565b6040519080825280602002602001820160405280156116a8578160200160208202803683370190505b50905060005b6116b86014612cab565b811015611704576116ca601482612cb5565b8282815181106116dc576116dc6146e5565b6001600160a01b03909216602092830291909101909101526116fd81614714565b90506116ae565b50919050565b6000805460ff161561175e5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610b06565b600c54604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516001600160a01b03909216916346f8e6d7916004808201926020929091908290030181865afa1580156117c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117e59190614534565b1561181b576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061182d6109a560808701876145a8565b90506118a861183f60208701876145a8565b8351909150611851604089018961478f565b808060200260200160405190810160405280939291908181526020016000905b8282101561189d5761188e6040830286013681900381019061458c565b81526020019060010190611871565b505050505086612cc1565b6118b285806145a8565b90506020146118f9576118c585806145a8565b6040517f370d875f000000000000000000000000000000000000000000000000000000008152600401610b069291906147f7565b600061190586806145a8565b8101906119129190614826565b90506001600160a01b0381111561192d576118c586806145a8565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166119676080880160608901613cd7565b6001600160a01b0316036119925784601760008282546119879190614655565b90915550611a7b9050565b600a546001600160a01b031663a5817a957f00000000000000000000000000000000000000000000000000000000000000006119d460808a0160608b01613cd7565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03928316600482015291166024820152604481018890526064016020604051808303816000875af1158015611a40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a64919061460d565b60176000828254611a759190614655565b90915550505b60005b611a8b604088018861478f565b9050811015611b7e576000611aa3604089018961478f565b83818110611ab357611ab36146e5565b905060400201803603810190611ac9919061458c565b90506000611ada82600001516129ef565b60208301516040517fe2e59b3e0000000000000000000000000000000000000000000000000000000081529192506001600160a01b0383169163e2e59b3e91611b39918b906004019182526001600160a01b0316602082015260400190565b600060405180830381600087803b158015611b5357600080fd5b505af1158015611b67573d6000803e3d6000fd5b50505050505080611b7790614714565b9050611a7e565b5060006040518061018001604052807f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168152602001601a600081819054906101000a900467ffffffffffffffff16611bdf9061483f565b825467ffffffffffffffff9182166101009390930a8381029083021990911617909255825260208083018a90526001600160a01b03891660408085018290526000918252601990925290812080546060909401939092611c3f911661483f565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905567ffffffffffffffff16815260200184600001518152602001846020015115158152602001836001600160a01b03168152602001888060200190611ca891906145a8565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250602001611cef60408a018a61478f565b808060200260200160405190810160405280939291908181526020016000905b82821015611d3b57611d2c6040830286013681900381019061458c565b81526020019060010190611d0f565b5050509183525050602001611d5660808a0160608b01613cd7565b6001600160a01b0316815260006020909101529050611d95817f0000000000000000000000000000000000000000000000000000000000000000612ed3565b6101608201526040517faffc45517195d6499808c643bd4a7b0ffeedf95bea5852840d7bfcf63f59e82190611dcb908390614866565b60405180910390a16101600151925050505b9392505050565b6060600080611df3600d612fdd565b90508067ffffffffffffffff811115611e0e57611e0e613d3b565b604051908082528060200260200182016040528015611e5357816020015b6040805180820190915260008082526020820152815260200190600190039081611e2c5790505b50925060005b81811015611ec157600080611e6f600d84612fe8565b915091506040518060400160405280836001600160a01b0316815260200182815250868481518110611ea357611ea36146e5565b6020026020010181905250505080611eba90614714565b9050611e59565b506018549150509091565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015611f0957506002546001600160a01b03163314155b15611f40576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611f4a6005613006565b60408101516006819055602082015179ffffffffffffffffffffffffffffffffffffffffffffffffffff16600555600754611f859190612c95565b600755604081810151602080840151835192835279ffffffffffffffffffffffffffffffffffffffffffffffffffff16908201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b611ff0612479565b6114b8816130b1565b6060600061200760106132fe565b67ffffffffffffffff81111561201f5761201f613d3b565b604051908082528060200260200182016040528015612048578160200160208202803683370190505b50905060005b815181101561170457612062601082613309565b50828281518110612075576120756146e5565b60200260200101816001600160a01b03166001600160a01b0316815250508061209d90614714565b905061204e565b6120ac612479565b6013805460ff19168215159081179091556040519081527fccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df03290602001611fdd565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b0316141580156121305750600b546201000090046001600160a01b03163314155b80156121445750612142600d33613318565b155b1561217b576040517f63baed9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60185460008190036121b9576040517f990e30bf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6017548015806121c857508181105b156121ff576040517f8d0f71d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906370a0823190602401602060405180830381865afa15801561227f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122a3919061460d565b9050818110156122df576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000836122f484670de0b6b3a764000061466d565b6122fe91906146aa565b905060005b61230d600d612fdd565b8110156123df57600080612322600d84612fe8565b90925090506000670de0b6b3a764000061233c838761466d565b61234691906146aa565b905061237c6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016848361281c565b6123868188614778565b9650826001600160a01b03167f55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f826040516123c391815260200190565b60405180910390a2505050806123d890614714565b9050612303565b50505060175550565b6123f0612479565b6114b88161332d565b600061240f846001600160a01b0385168461242d565b949350505050565b600061240f846001600160a01b038516846133f4565b600061240f84846001600160a01b0385166133f4565b6000611ddd836001600160a01b038416613411565b6000611ddd836001600160a01b038416613504565b6000611ddd8383613504565b60005461010090046001600160a01b03163314610f095760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610b06565b600d60008181818181816124ec8282613ba3565b50505050505050506000805b825181101561258d57612551838281518110612516576125166146e5565b602002602001015160000151848381518110612534576125346146e5565b602002602001015160200151600d6124179092919063ffffffff16565b50828181518110612564576125646146e5565b6020026020010151602001518261257b9190614655565b915061258681614714565b90506124f8565b5060188190556040517f5c5cf993d6e35edd9d958829f41cc0961071d6f5ac714b1ce86fc44c73da1549906125c59083908590614985565b60405180910390a15050565b6040805180820190915260008082526020820152600082900361263257506040805180820190915267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526000602082015261142f565b7f97a657c90000000000000000000000000000000000000000000000000000000061266160046000858761499e565b61266a916149c8565b7fffffffff000000000000000000000000000000000000000000000000000000001614612727577f97a657c9000000000000000000000000000000000000000000000000000000006126c060046000858761499e565b6126c9916149c8565b6040517f55a0e02c0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000928316600482015291166024820152604401610b06565b612734826004818661499e565b810190611ddd9190614a10565b6000611ddd836001600160a01b038416613553565b6000611ddd836001600160a01b03841661355f565b6000611ddd836001600160a01b03841661356b565b60005460ff166127d25760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610b06565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610ef4908490613577565b60005b8151811015612945576128d58282815181106128bd576128bd6146e5565b6020026020010151601461244390919063ffffffff16565b507f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566828281518110612909576129096146e5565b602002602001015160405161292d91906001600160a01b0391909116815260200190565b60405180910390a161293e81614714565b905061289f565b5060005b8251811015610ef45761297f838281518110612967576129676146e5565b6020026020010151601461245890919063ffffffff16565b507f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d88382815181106129b3576129b36146e5565b60200260200101516040516129d791906001600160a01b0391909116815260200190565b60405180910390a16129e881614714565b9050612949565b60006129fc601083612741565b612a3d576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610b06565b61142f601083612756565b60005b8151811015612bdd576040518060600160405280838381518110612a7157612a716146e5565b6020026020010151604001516bffffffffffffffffffffffff168152602001838381518110612aa257612aa26146e5565b60200260200101516020015167ffffffffffffffff168152602001838381518110612acf57612acf6146e5565b60200260200101516060015163ffffffff1681525060166000848481518110612afa57612afa6146e5565b602090810291909101810151516001600160a01b031682528181019290925260409081016000208351815493850151949092015163ffffffff1674010000000000000000000000000000000000000000027fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff9095166c01000000000000000000000000027fffffffffffffffffffffffff00000000000000000000000000000000000000009094166bffffffffffffffffffffffff909316929092179290921792909216919091179055612bd681614714565b9050612a4b565b507f1026720e82621bb6d100957028cafff5ccf24add051c05d1bd65fa1d71e0fae681604051611fdd9190614a66565b60005460ff1615612c605760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610b06565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586127ff3390565b6000818310612ca45781611ddd565b5090919050565b600061142f825490565b6000611ddd838361365c565b6009546001600160a01b03163314612d05576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038116612d45576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a5474010000000000000000000000000000000000000000900463ffffffff16841115612dc957600a546040517f869337890000000000000000000000000000000000000000000000000000000081527401000000000000000000000000000000000000000090910463ffffffff16600482015260248101859052604401610b06565b600a547801000000000000000000000000000000000000000000000000900467ffffffffffffffff16831115612e2b576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b54825161ffff9091161015612e6e576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60135460ff168015612e885750612e86601482613686565b155b15612eca576040517fd0d259760000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401610b06565b61141e826136a8565b60008060001b828460200151856080015186606001518760e0015188610100015180519060200120896101200151604051602001612f119190614af1565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d60400151604051602001612fbf9c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b600061142f826138ac565b6000808080612ff786866138b7565b909450925050505b9250929050565b80600101548160020154148061301f5750428160030154145b156130275750565b806001015481600201541115613069576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081600301544261307b9190614778565b600183015483549192506130a291613093908461466d565b84600201546116439190614655565b60028301555042600390910155565b80516001600160a01b031615806130d3575060208101516001600160a01b0316155b806130e9575060c08101516001600160a01b0316155b15613120576040517f35be3ac800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600980546001600160a01b039283167fffffffffffffffffffffffff0000000000000000000000000000000000000000918216179091556020830151600a8054604080870151606088015167ffffffffffffffff1678010000000000000000000000000000000000000000000000000277ffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff90921674010000000000000000000000000000000000000000027fffffffffffffffff0000000000000000000000000000000000000000000000009094169588169590951792909217919091169290921790556080840151600b805460a0870151861662010000027fffffffffffffffffffff0000000000000000000000000000000000000000000090911661ffff9093169290921791909117905560c0840151600c805491909416921691909117909155517f0161b2e931b61fb04ae29bb225d36a3248c2b2abb7348385e3e74e5740490b8a90611fdd908390600060e0820190506001600160a01b0380845116835280602085015116602084015263ffffffff604085015116604084015267ffffffffffffffff606085015116606084015261ffff60808501511660808401528060a08501511660a08401528060c08501511660c08401525092915050565b600061142f82612fdd565b6000808080612ff78686612fe8565b6000611ddd836001600160a01b0384166138e2565b336001600160a01b038216036133855760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610b06565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000828152600284016020526040812082905561240f848461246d565b600081815260018301602052604081205480156134fa576000613435600183614778565b855490915060009061344990600190614778565b90508181146134ae576000866000018281548110613469576134696146e5565b906000526020600020015490508087600001848154811061348c5761348c6146e5565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806134bf576134bf614b04565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061142f565b600091505061142f565b600081815260018301602052604081205461354b5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561142f565b50600061142f565b6000611ddd83836138e2565b6000611ddd83836138ee565b6000611ddd838361395e565b60006135cc826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661397b9092919063ffffffff16565b805190915015610ef457808060200190518101906135ea9190614534565b610ef45760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610b06565b6000826000018281548110613673576136736146e5565b9060005260206000200154905092915050565b6001600160a01b03811660009081526001830160205260408120541515611ddd565b6000805b82518110156137a7576000600360008584815181106136cd576136cd6146e5565b6020026020010151600001516001600160a01b03166001600160a01b031681526020019081526020016000205490508060000361376057838281518110613716576137166146e5565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b06565b838281518110613772576137726146e5565b60200260200101516020015181613789919061466d565b6137939084614655565b925050806137a090614714565b90506136ac565b50801561098e576137b86005613006565b600654811115613802576006546040517f688ccf77000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610b06565b60075481111561386257600554600754600091906138209084614778565b61382a91906146aa565b9050806040517fe31e0f32000000000000000000000000000000000000000000000000000000008152600401610b0691815260200190565b80600560020160008282546138779190614778565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378906020016125c5565b600061142f82612cab565b600080806138c58585612cb5565b600081815260029690960160205260409095205494959350505050565b6000611ddd838361398a565b600081815260028301602052604081205480151580613912575061391284846138e2565b611ddd5760405162461bcd60e51b815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b657900006044820152606401610b06565b60008181526002830160205260408120819055611ddd83836139a2565b606061240f84846000856139ae565b60008181526001830160205260408120541515611ddd565b6000611ddd8383613411565b606082471015613a265760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610b06565b843b613a745760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610b06565b600080866001600160a01b03168587604051613a909190614b33565b60006040518083038185875af1925050503d8060008114613acd576040519150601f19603f3d011682016040523d82523d6000602084013e613ad2565b606091505b5091509150613ae2828286613aed565b979650505050505050565b60608315613afc575081611ddd565b825115613b0c5782518084602001fd5b8160405162461bcd60e51b8152600401610b069190613c2a565b828054828255906000526020600020908101928215613b93579160200282015b82811115613b9357825182547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03909116178255602090920191600190910190613b46565b50613b9f929150613bbd565b5090565b50805460008255906000526020600020908101906114b891905b5b80821115613b9f5760008155600101613bbe565b60005b83811015613bed578181015183820152602001613bd5565b8381111561141e5750506000910152565b60008151808452613c16816020860160208601613bd2565b601f01601f19169290920160200192915050565b602081526000611ddd6020830184613bfe565b60008060208385031215613c5057600080fd5b823567ffffffffffffffff80821115613c6857600080fd5b818501915085601f830112613c7c57600080fd5b813581811115613c8b57600080fd5b8660208260061b8501011115613ca057600080fd5b60209290920196919550909350505050565b6001600160a01b03811681146114b857600080fd5b8035613cd281613cb2565b919050565b600060208284031215613ce957600080fd5b8135611ddd81613cb2565b600060a0828403121561170457600080fd5b600060208284031215613d1857600080fd5b813567ffffffffffffffff811115613d2f57600080fd5b61240f84828501613cf4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613d8d57613d8d613d3b565b60405290565b6040516080810167ffffffffffffffff81118282101715613d8d57613d8d613d3b565b60405160e0810167ffffffffffffffff81118282101715613d8d57613d8d613d3b565b604051601f8201601f1916810167ffffffffffffffff81118282101715613e0257613e02613d3b565b604052919050565b600067ffffffffffffffff821115613e2457613e24613d3b565b5060051b60200190565b600082601f830112613e3f57600080fd5b81356020613e54613e4f83613e0a565b613dd9565b82815260069290921b84018101918181019086841115613e7357600080fd5b8286015b84811015613ec45760408189031215613e905760008081fd5b613e98613d6a565b8135613ea381613cb2565b815281850135613eb281613cb2565b81860152835291830191604001613e77565b509695505050505050565b60008060408385031215613ee257600080fd5b823567ffffffffffffffff80821115613efa57600080fd5b613f0686838701613e2e565b93506020850135915080821115613f1c57600080fd5b50613f2985828601613e2e565b9150509250929050565b600082601f830112613f4457600080fd5b81356020613f54613e4f83613e0a565b82815260059290921b84018101918181019086841115613f7357600080fd5b8286015b84811015613ec4578035613f8a81613cb2565b8352918301918301613f77565b60008060408385031215613faa57600080fd5b823567ffffffffffffffff80821115613fc257600080fd5b613fce86838701613f33565b9350602091508185013581811115613fe557600080fd5b85019050601f81018613613ff857600080fd5b8035614006613e4f82613e0a565b81815260059190911b8201830190838101908883111561402557600080fd5b928401925b828410156140435783358252928401929084019061402a565b80955050505050509250929050565b60006020828403121561406457600080fd5b813567ffffffffffffffff81111561407b57600080fd5b61240f84828501613f33565b6020808252825182820181905260009190848201906040850190845b818110156140bf578351835292840192918401916001016140a3565b50909695505050505050565b600080604083850312156140de57600080fd5b82356140e981613cb2565b915060208301356140f981613cb2565b809150509250929050565b60008083601f84011261411657600080fd5b50813567ffffffffffffffff81111561412e57600080fd5b6020830191508360208260051b8501011115612fff57600080fd5b6000806000806040858703121561415f57600080fd5b843567ffffffffffffffff8082111561417757600080fd5b61418388838901614104565b9096509450602087013591508082111561419c57600080fd5b506141a987828801614104565b95989497509550505050565b803567ffffffffffffffff81168114613cd257600080fd5b803563ffffffff81168114613cd257600080fd5b600060208083850312156141f457600080fd5b823567ffffffffffffffff81111561420b57600080fd5b8301601f8101851361421c57600080fd5b803561422a613e4f82613e0a565b81815260079190911b8201830190838101908783111561424957600080fd5b928401925b82841015613ae257608084890312156142675760008081fd5b61426f613d93565b843561427a81613cb2565b81526142878587016141b5565b868201526040808601356bffffffffffffffffffffffff811681146142ac5760008081fd5b9082015260606142bd8682016141cd565b908201528252608093909301929084019061424e565b6020808252825182820181905260009190848201906040850190845b818110156140bf5783516001600160a01b0316835292840192918401916001016142ef565b60008060006060848603121561432957600080fd5b833567ffffffffffffffff81111561434057600080fd5b61434c86828701613cf4565b93505060208401359150604084013561436481613cb2565b809150509250925092565b600081518084526020808501945080840160005b838110156143b357815180516001600160a01b031688528301518388015260409096019590820190600101614383565b509495945050505050565b6040815260006143d1604083018561436f565b90508260208301529392505050565b6000606082840312156143f257600080fd5b6040516060810181811067ffffffffffffffff8211171561441557614415613d3b565b604052823561442381613cb2565b8152602083013579ffffffffffffffffffffffffffffffffffffffffffffffffffff8116811461445257600080fd5b60208201526040928301359281019290925250919050565b600060e0828403121561447c57600080fd5b614484613db6565b823561448f81613cb2565b8152602083013561449f81613cb2565b60208201526144b0604084016141cd565b60408201526144c1606084016141b5565b6060820152608083013561ffff811681146144db57600080fd5b60808201526144ec60a08401613cc7565b60a08201526144fd60c08401613cc7565b60c08201529392505050565b80151581146114b857600080fd5b60006020828403121561452957600080fd5b8135611ddd81614509565b60006020828403121561454657600080fd5b8151611ddd81614509565b60006040828403121561456357600080fd5b61456b613d6a565b9050813561457881613cb2565b808252506020820135602082015292915050565b60006040828403121561459e57600080fd5b611ddd8383614551565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126145dd57600080fd5b83018035915067ffffffffffffffff8211156145f857600080fd5b602001915036819003821315612fff57600080fd5b60006020828403121561461f57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561466857614668614626565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156146a5576146a5614626565b500290565b6000826146e0577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361474557614745614626565b5060010190565b600067ffffffffffffffff80831681851680830382111561476f5761476f614626565b01949350505050565b60008282101561478a5761478a614626565b500390565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126147c457600080fd5b83018035915067ffffffffffffffff8211156147df57600080fd5b6020019150600681901b3603821315612fff57600080fd5b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60006020828403121561483857600080fd5b5035919050565b600067ffffffffffffffff80831681810361485c5761485c614626565b6001019392505050565b6020815261488160208201835167ffffffffffffffff169052565b6000602083015161489e604084018267ffffffffffffffff169052565b506040830151606083015260608301516148c360808401826001600160a01b03169052565b50608083015167ffffffffffffffff811660a08401525060a083015160c083015260c08301516148f760e084018215159052565b5060e0830151610100614914818501836001600160a01b03169052565b8085015191505061018061012081818601526149346101a0860184613bfe565b9250808601519050610140601f198685030181870152614954848361436f565b935080870151915050610160614974818701836001600160a01b03169052565b959095015193019290925250919050565b82815260406020820152600061240f604083018461436f565b600080858511156149ae57600080fd5b838611156149bb57600080fd5b5050820193919092039150565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015614a085780818660040360031b1b83161692505b505092915050565b600060408284031215614a2257600080fd5b6040516040810181811067ffffffffffffffff82111715614a4557614a45613d3b565b604052823581526020830135614a5a81614509565b60208201529392505050565b602080825282518282018190526000919060409081850190868401855b82811015614ae457815180516001600160a01b031685528681015167ffffffffffffffff1687860152858101516bffffffffffffffffffffffff168686015260609081015163ffffffff169085015260809093019290850190600101614a83565b5091979650505050505050565b602081526000611ddd602083018461436f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008251614b45818460208701613bd2565b919091019291505056fea164736f6c634300080f000a",
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

type EVM2EVMOnRampDynamicConfigSetIterator struct {
	Event *EVM2EVMOnRampDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampDynamicConfigSet)
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
		it.Event = new(EVM2EVMOnRampDynamicConfigSet)
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

func (it *EVM2EVMOnRampDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampDynamicConfigSet struct {
	Arg0 IEVM2EVMOnRampDynamicConfig
	Raw  types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*EVM2EVMOnRampDynamicConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampDynamicConfigSetIterator{contract: _EVM2EVMOnRamp.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampDynamicConfigSet)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseDynamicConfigSet(log types.Log) (*EVM2EVMOnRampDynamicConfigSet, error) {
	event := new(EVM2EVMOnRampDynamicConfigSet)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

type EVM2EVMOnRampStaticConfigSetIterator struct {
	Event *EVM2EVMOnRampStaticConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOnRampStaticConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOnRampStaticConfigSet)
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
		it.Event = new(EVM2EVMOnRampStaticConfigSet)
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

func (it *EVM2EVMOnRampStaticConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOnRampStaticConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOnRampStaticConfigSet struct {
	Arg0 IEVM2EVMOnRampStaticConfig
	Raw  types.Log
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) FilterStaticConfigSet(opts *bind.FilterOpts) (*EVM2EVMOnRampStaticConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.FilterLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOnRampStaticConfigSetIterator{contract: _EVM2EVMOnRamp.contract, event: "StaticConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampStaticConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOnRamp.contract.WatchLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOnRampStaticConfigSet)
				if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
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

func (_EVM2EVMOnRamp *EVM2EVMOnRampFilterer) ParseStaticConfigSet(log types.Log) (*EVM2EVMOnRampStaticConfigSet, error) {
	event := new(EVM2EVMOnRampStaticConfigSet)
	if err := _EVM2EVMOnRamp.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
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
	case _EVM2EVMOnRamp.abi.Events["DynamicConfigSet"].ID:
		return _EVM2EVMOnRamp.ParseDynamicConfigSet(log)
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
	case _EVM2EVMOnRamp.abi.Events["StaticConfigSet"].ID:
		return _EVM2EVMOnRamp.ParseStaticConfigSet(log)
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

func (EVM2EVMOnRampDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x0161b2e931b61fb04ae29bb225d36a3248c2b2abb7348385e3e74e5740490b8a")
}

func (EVM2EVMOnRampFeeConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1026720e82621bb6d100957028cafff5ccf24add051c05d1bd65fa1d71e0fae6")
}

func (EVM2EVMOnRampNopPaid) Topic() common.Hash {
	return common.HexToHash("0x55fdec2aab60a41fa5abb106670eb1006f5aeaee1ba7afea2bc89b5b3ec7678f")
}

func (EVM2EVMOnRampNopsSet) Topic() common.Hash {
	return common.HexToHash("0x5c5cf993d6e35edd9d958829f41cc0961071d6f5ac714b1ce86fc44c73da1549")
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

func (EVM2EVMOnRampStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0x88ef61c99020a9e2d75ad86191c6f07f7b6ce93ccebda542554b2b1588204ae7")
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

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*EVM2EVMOnRampDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*EVM2EVMOnRampDynamicConfigSet, error)

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

	FilterStaticConfigSet(opts *bind.FilterOpts) (*EVM2EVMOnRampStaticConfigSetIterator, error)

	WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOnRampStaticConfigSet) (event.Subscription, error)

	ParseStaticConfigSet(log types.Log) (*EVM2EVMOnRampStaticConfigSet, error)

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
