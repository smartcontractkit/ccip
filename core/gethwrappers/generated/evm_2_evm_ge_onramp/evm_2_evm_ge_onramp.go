// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_ge_onramp

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

type CommonEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

type GEConsumerEVM2AnyGEMessage struct {
	Receiver         []byte
	Data             []byte
	TokensAndAmounts []CommonEVMTokenAndAmount
	FeeToken         common.Address
	ExtraArgs        []byte
}

type GEEVM2EVMGEMessage struct {
	SourceChainId    uint64
	SequenceNumber   uint64
	FeeTokenAmount   *big.Int
	Sender           common.Address
	Nonce            uint64
	GasLimit         *big.Int
	Strict           bool
	Receiver         common.Address
	Data             []byte
	TokensAndAmounts []CommonEVMTokenAndAmount
	FeeToken         common.Address
	MessageId        [32]byte
}

type IAggregateRateLimiterRateLimiterConfig struct {
	Rate     *big.Int
	Capacity *big.Int
}

type IAggregateRateLimiterTokenBucket struct {
	Rate        *big.Int
	Capacity    *big.Int
	Tokens      *big.Int
	LastUpdated *big.Int
}

type IBaseOnRampOnRampConfig struct {
	CommitFeeJuels  uint64
	MaxDataSize     uint64
	MaxTokensLength uint64
	MaxGasLimit     uint64
}

type IEVM2EVMGEOnRampDynamicFeeConfig struct {
	LinkToken       common.Address
	FeeAmount       *big.Int
	DestGasOverhead *big.Int
	Multiplier      *big.Int
	FeeManager      common.Address
	DestChainId     uint64
}

var EVM2EVMGEOnRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destinationChainId\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"commitFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structIBaseOnRamp.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"},{\"internalType\":\"contractIGERouter\",\"name\":\"router\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destGasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMGEOnRamp.DynamicFeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"expected\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"got\",\"type\":\"bytes4\"}],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageGasLimitTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeCalledByRouter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterMustSetOriginalSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chain\",\"type\":\"uint64\"}],\"name\":\"TokenOrChainNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"AllowListSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structGE.EVM2EVMGEMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"CCIPSendRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeAdmin\",\"type\":\"address\"}],\"name\":\"FeeAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destGasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMGEOnRamp.DynamicFeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"FeeConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"commitFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIBaseOnRamp.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OnRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structGEConsumer.EVM2AnyGEMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"}],\"name\":\"forwardFromRouter\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowlistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationChainId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedNextSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structGEConsumer.EVM2AnyGEMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOnRampConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"commitFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structIBaseOnRamp.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"}],\"name\":\"setAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setAllowlistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAdmin\",\"type\":\"address\"}],\"name\":\"setFeeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"linkToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destGasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMGEOnRamp.DynamicFeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"setFeeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"commitFeeJuels\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxGasLimit\",\"type\":\"uint64\"}],\"internalType\":\"structIBaseOnRamp.OnRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOnRampConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620045d9380380620045d9833981016040819052620000349162000a9e565b6000805460ff191681558b908b908b908b908b908b908b908b908b908b9083908390889088903390819081620000b15760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000eb57620000eb8162000654565b5050506001600160a01b0381166200011657604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580511562000168576002805460ff60a01b1916600160a01b17905580516200016690600390602084019062000705565b505b60005b8151811015620001d55760016004600084848151811062000190576200019062000bdb565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055620001cd8162000bf1565b90506200016b565b505080600560006101000a8154816001600160a01b0302191690836001600160a01b031602179055506040518060800160405280836000015181526020018360200151815260200183602001518152602001428152506008600082015181600001556020820151816001015560408201518160020155606082015181600301559050505050896001600160401b03166080816001600160401b031681525050886001600160401b031660a0816001600160401b03168152505083600d60008201518160000160006101000a8154816001600160401b0302191690836001600160401b0316021790555060208201518160000160086101000a8154816001600160401b0302191690836001600160401b0316021790555060408201518160000160106101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160000160186101000a8154816001600160401b0302191690836001600160401b0316021790555090505080600e60006101000a8154816001600160a01b0302191690836001600160a01b031602179055506000600c60006101000a8154816001600160401b0302191690836001600160401b031602179055508651885114620003b95760405162d8548360e71b815260040160405180910390fd5b8751620003ce9060109060208b019062000705565b5060005b885181101562000495576040518060400160405280898381518110620003fc57620003fc62000bdb565b60200260200101516001600160a01b0316815260200160011515815250600f60008b848151811062000432576200043262000bdb565b6020908102919091018101516001600160a01b0390811683528282019390935260409091016000208351815494909201511515600160a01b026001600160a81b031990941691909216179190911790556200048d8162000bf1565b9050620003d2565b505050505050505050505080601260008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060a08201518160040160146101000a8154816001600160401b0302191690836001600160401b031602179055509050507fba22a5847647789e6efe1840c86bc66129ac89e03d7b95e0eebdf7fa43763fdd8b8b30604051602001620005ac94939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b60408051808303601f19018152828252805160209182012060c090815284516001600160a01b03908116855285830151928501929092528483015192840192909252606080850151908401526080808501519091169083015260a0808401516001600160401b0316908301527f92c1d3bc951c7322787dfee144ba08e1dbc36535f211ae5fbdbbc3c27c3b26be910160405180910390a1505050505050505050505062000c19565b336001600160a01b03821603620006ae5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a8565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200075d579160200282015b828111156200075d57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000726565b506200076b9291506200076f565b5090565b5b808211156200076b576000815560010162000770565b80516001600160401b03811681146200079e57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620007e457620007e4620007a3565b604052919050565b60006001600160401b03821115620008085762000808620007a3565b5060051b60200190565b6001600160a01b03811681146200082857600080fd5b50565b80516200079e8162000812565b600082601f8301126200084a57600080fd5b81516020620008636200085d83620007ec565b620007b9565b82815260059290921b840181019181810190868411156200088357600080fd5b8286015b84811015620008ab5780516200089d8162000812565b835291830191830162000887565b509695505050505050565b600082601f830112620008c857600080fd5b81516020620008db6200085d83620007ec565b82815260059290921b84018101918181019086841115620008fb57600080fd5b8286015b84811015620008ab578051620009158162000812565b8352918301918301620008ff565b6000608082840312156200093657600080fd5b604051608081016001600160401b03811182821017156200095b576200095b620007a3565b6040529050806200096c8362000786565b81526200097c6020840162000786565b60208201526200098f6040840162000786565b6040820152620009a26060840162000786565b60608201525092915050565b600060408284031215620009c157600080fd5b604080519081016001600160401b0381118282101715620009e657620009e6620007a3565b604052825181526020928301519281019290925250919050565b600060c0828403121562000a1357600080fd5b60405160c081016001600160401b038111828210171562000a385762000a38620007a3565b8060405250809150825162000a4d8162000812565b80825250602083015160208201526040830151604082015260608301516060820152608083015162000a7f8162000812565b608082015262000a9260a0840162000786565b60a08201525092915050565b60008060008060008060008060008060006102808c8e03121562000ac157600080fd5b62000acc8c62000786565b9a5062000adc60208d0162000786565b60408d0151909a506001600160401b0381111562000af957600080fd5b62000b078e828f0162000838565b60608e0151909a5090506001600160401b0381111562000b2657600080fd5b62000b348e828f01620008b6565b60808e015190995090506001600160401b0381111562000b5357600080fd5b62000b618e828f0162000838565b97505062000b7260a08d016200082b565b955062000b838d60c08e0162000923565b945062000b958d6101408e01620009ae565b935062000ba66101808d016200082b565b925062000bb76101a08d016200082b565b915062000bc98d6101c08e0162000a00565b90509295989b509295989b9093969950565b634e487b7160e01b600052603260045260246000fd5b60006001820162000c1257634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05161398962000c506000396000611860015260006103c8015260008181610316015261167801526139896000f3fe608060405234801561001057600080fd5b506004361061025c5760003560e01c80636eb2d03111610145578063a7d3e02f116100bd578063d3c7c2c71161008c578063eb511dd411610071578063eb511dd414610653578063f2fde38b14610666578063f78faa321461067957600080fd5b8063d3c7c2c714610638578063d7644ba21461064057600080fd5b8063a7d3e02f146105ec578063b0f479a1146105ff578063c0d7865514610610578063c5eff3d01461062357600080fd5b8063856c82471161011457806390c2339b116100f957806390c2339b1461058b57806391872543146105c6578063a11855be146105d957600080fd5b8063856c8247146105425780638da5cb5b1461057557600080fd5b80636eb2d0311461050c578063744b92e21461051f57806379ba5097146105325780638456cb591461053a57600080fd5b80634120fccd116101d8578063499bb0e2116101a7578063599f64311161018c578063599f6431146104dd5780635c975abb146104ee5780635d86f141146104f957600080fd5b8063499bb0e2146103ec578063552b818b146104ca57600080fd5b80634120fccd1461038b5780634352fa9f146103935780634741062e146103a65780634894b536146103c657600080fd5b80632222dd421161022f57806338724a951161021457806338724a951461034f57806339aa9264146103705780633f4ba83a1461038357600080fd5b80632222dd42146102ef5780633408e4701461031457600080fd5b8063105f7d9314610261578063108ee5fc14610276578063147809b314610289578063181f5a77146102a6575b600080fd5b61027461026f366004612c2a565b61068b565b005b610274610284366004612c57565b6106dd565b610291610794565b60405190151581526020015b60405180910390f35b6102e26040518060400160405280601581526020017f45564d3245564d47454f6e52616d7020312e302e30000000000000000000000081525081565b60405161029d9190612cd0565b6002546001600160a01b03165b6040516001600160a01b03909116815260200161029d565b7f00000000000000000000000000000000000000000000000000000000000000005b60405167ffffffffffffffff909116815260200161029d565b61036261035d366004612cf5565b610821565b60405190815260200161029d565b61027461037e366004612c57565b6109d2565b610274610a09565b610336610a1b565b6102746103a1366004612e4b565b610a3b565b6103b96103b4366004612f06565b610c90565b60405161029d9190612f3b565b7f0000000000000000000000000000000000000000000000000000000000000000610336565b6104866040805160808101825260008082526020820181905291810182905260608101919091525060408051608081018252600d5467ffffffffffffffff808216835268010000000000000000820481166020840152700100000000000000000000000000000000820481169383019390935278010000000000000000000000000000000000000000000000009004909116606082015290565b60405161029d9190815167ffffffffffffffff9081168252602080840151821690830152604080840151821690830152606092830151169181019190915260800190565b6102746104d8366004612f7f565b610d58565b6005546001600160a01b03166102fc565b60005460ff16610291565b6102fc610507366004612c57565b610ee2565b61027461051a366004612c57565b610ef3565b61027461052d366004612ff4565b610f56565b6102746110a6565b610274611189565b610336610550366004612c57565b6001600160a01b031660009081526011602052604090205467ffffffffffffffff1690565b60005461010090046001600160a01b03166102fc565b610593611199565b60405161029d91908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6102746105d436600461302d565b61123a565b6102746105e736600461305f565b611366565b6103626105fa366004613071565b6113ad565b600e546001600160a01b03166102fc565b61027461061e366004612c57565b6118d2565b61062b611935565b60405161029d91906130cc565b61062b611997565b61027461064e36600461311b565b6119a1565b610274610661366004612ff4565b611a11565b610274610674366004612c57565b611bb2565b600254600160a01b900460ff16610291565b610693611bc6565b80600d6106a0828261314e565b9050507f0447ae479bc793c12cd12089f932a4c0b4ac50f1da17f1379c3d420af34407b8816040516106d2919061325f565b60405180910390a150565b6106e5611bc6565b6001600160a01b038116610725576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa1580156107f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081b91906132c9565b15905090565b60008061083961083460808501856132e6565b611c25565b516016549091506000906001600160a01b0316638e160ef46108616080870160608801612c57565b60165460405160e084901b7fffffffff000000000000000000000000000000000000000000000000000000001681526001600160a01b039092166004830152600160a01b900467ffffffffffffffff166024820152604401602060405180830381865afa1580156108d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fa9190613352565b6fffffffffffffffffffffffffffffffff16905080600003610982576109266080850160608601612c57565b6016546040517f102e3c280000000000000000000000000000000000000000000000000000000081526001600160a01b039092166004830152600160a01b900467ffffffffffffffff1660248201526044015b60405180910390fd5b601554601454670de0b6b3a76400009190839061099f90866133b3565b6109a991906133cb565b6109b391906133cb565b6109bd9190613408565b6013546109ca91906133b3565b949350505050565b6109da611bc6565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b610a11611bc6565b610a19611d8b565b565b600c54600090610a369067ffffffffffffffff166001613443565b905090565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015610a7857506005546001600160a01b03163314155b15610aaf576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610aeb576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610b45576006600060078381548110610b1057610b1061346f565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610b3e8161349e565b9050610af1565b5060005b82811015610c75576000858281518110610b6557610b6561346f565b6020026020010151905060006001600160a01b0316816001600160a01b031603610bbb576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610bcd57610bcd61346f565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610c3257610c3261346f565b6020026020010151604051610c5c9291906001600160a01b03929092168252602082015260400190565b60405180910390a150610c6e8161349e565b9050610b49565b508351610c89906007906020870190612b43565b5050505050565b80516060908067ffffffffffffffff811115610cae57610cae612d2a565b604051908082528060200260200182016040528015610cd7578160200160208202803683370190505b50915060005b81811015610d515760066000858381518110610cfb57610cfb61346f565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610d3657610d3661346f565b6020908102919091010152610d4a8161349e565b9050610cdd565b5050919050565b610d60611bc6565b60006003805480602002602001604051908101604052809291908181526020018280548015610db857602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d9a575b5050505050905060005b8151811015610e2557600060046000848481518110610de357610de361346f565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055610e1e8161349e565b9050610dc2565b50610e3260038484612bb5565b5060005b82811015610ea357600160046000868685818110610e5657610e5661346f565b9050602002016020810190610e6b9190612c57565b6001600160a01b031681526020810191909152604001600020805460ff1916911515919091179055610e9c8161349e565b9050610e36565b507ff8adc5fee247b62a85f63294cb46e4be61da815e4756bc57672a83b24faf0dda8383604051610ed59291906134d6565b60405180910390a1505050565b6000610eed82611e27565b92915050565b610efb611bc6565b6017805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f34efeb1f04731080ec2147b3b9c7e38f9b884e035020914e40269450f53b4815906020016106d2565b610f5e611bc6565b6001600160a01b038281166000908152600f6020908152604091829020825180840190935254928316808352600160a01b90930460ff1615159082015290610fdd576040517f73913ebd0000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401610979565b816001600160a01b031681600001516001600160a01b03161461102c576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038381166000818152600f602090815260409182902080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690558151928352928516928201929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c9101610ed5565b6001546001600160a01b031633146111005760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610979565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b611191611bc6565b610a19611eca565b6111c46040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b546060820181905242906000906111fe9083613519565b6020840151845191925061122a9161121690846133cb565b856040015161122591906133b3565b611f52565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561127757506005546001600160a01b03163314155b156112ae576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff11611302576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61130c6008611f68565b602081015160098190558151600855600a546113289190611f52565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91016106d2565b61136e611bc6565b80601261137b8282613530565b9050507f92c1d3bc951c7322787dfee144ba08e1dbc36535f211ae5fbdbbc3c27c3b26be816040516106d29190613616565b6000805460ff16156114015760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610979565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611454573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061147891906132c9565b156114ae576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600e546001600160a01b031633146114f2576040517f1c0a352900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6012546001600160a01b031661150e6080860160608701612c57565b6001600160a01b0316036115ba5760006115316105076080870160608801612c57565b90506001600160a01b03811661158f576115516080860160608701612c57565b6040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610979565b6115b481856115a46080890160608a01612c57565b6001600160a01b03169190612015565b506115db565b6016546115db906001600160a01b0316846115a46080880160608901612c57565b60006115ed61083460808701876132e6565b90506116686115ff60208701876132e6565b8351909150611611604089018961368b565b808060200260200160405190810160405280939291908181526020016000905b8282101561165d5761164e604083028601368190038101906136f3565b81526020019060010190611631565b50505050508661209a565b60006040518061018001604052807f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168152602001600c600081819054906101000a900467ffffffffffffffff166116c89061372d565b825467ffffffffffffffff9182166101009390930a8381029083021990911617909255825260208083018990526001600160a01b03881660408085018290526000918252601190925290812080546060909401939092611728911661372d565b825467ffffffffffffffff9182166101009390930a8381029202191617909155815283516020808301919091528401511515604082015260600161176c88806132e6565b8101906117799190612c57565b6001600160a01b0316815260200187806020019061179791906132e6565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506020016117de604089018961368b565b808060200260200160405190810160405280939291908181526020016000905b8282101561182a5761181b604083028601368190038101906136f3565b815260200190600101906117fe565b50505091835250506020016118456080890160608a01612c57565b6001600160a01b0316815260006020909101529050611884817f00000000000000000000000000000000000000000000000000000000000000006123ea565b6101608201526040517faffc45517195d6499808c643bd4a7b0ffeedf95bea5852840d7bfcf63f59e821906118ba9083906137a3565b60405180910390a161016001519150505b9392505050565b6118da611bc6565b600e805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527fc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15906020016106d2565b6060600380548060200260200160405190810160405280929190818152602001828054801561198d57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161196f575b5050505050905090565b6060610a366124f4565b6119a9611bc6565b60028054821515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff9091161790556040517fccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df032906106d290831515815260200190565b611a19611bc6565b6001600160a01b0382161580611a3657506001600160a01b038116155b15611a6d576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382166000908152600f6020526040902054600160a01b900460ff1615611ac7576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820182526001600160a01b03838116808352600160208085018281528885166000818152600f84528881209751885493511515600160a01b027fffffffffffffffffffffff000000000000000000000000000000000000000000909416971696909617919091179095556010805492830181559093527f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae67201805473ffffffffffffffffffffffffffffffffffffffff1916841790558351928352908201527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c9101610788565b611bba611bc6565b611bc38161268f565b50565b60005461010090046001600160a01b03163314610a195760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610979565b60408051808201909152600080825260208201526000829003611c5f57506040805180820190915262030d40815260006020820152610eed565b7f97a657c900000000000000000000000000000000000000000000000000000000611c8e6004600085876138c2565b611c97916138ec565b7fffffffff000000000000000000000000000000000000000000000000000000001614611d54577f97a657c900000000000000000000000000000000000000000000000000000000611ced6004600085876138c2565b611cf6916138ec565b6040517f55a0e02c0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000928316600482015291166024820152604401610979565b6040805180820190915280611d6d6024600486886138c2565b810190611d7a9190613934565b815260006020909101529392505050565b60005460ff16611ddd5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610979565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6001600160a01b038181166000908152600f602090815260408083208151808301909252549384168152600160a01b90930460ff1615801591840191909152909190611e8d5750506001600160a01b039081166000908152600f60205260409020541690565b6040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401610979565b60005460ff1615611f1d5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610979565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258611e0a3390565b6000818310611f6157816118cb565b5090919050565b6001810154600282015442911480611f835750808260030154145b15611f8c575050565b816001015482600201541115611fce576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482611fe09190613519565b6001840154845491925061200791611ff890846133cb565b856002015461122591906133b3565b600284015550600390910155565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261209590849061274b565b505050565b600e546001600160a01b03166120dc576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03811661211c576040517fa4ec747900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d5468010000000000000000900467ffffffffffffffff1684111561219057600d546040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090910467ffffffffffffffff16600482015260248101859052604401610979565b600d547801000000000000000000000000000000000000000000000000900467ffffffffffffffff168311156121f2576040517f4c4fc93a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151600d54700100000000000000000000000000000000900467ffffffffffffffff1681111561224e576040517f4c056b6a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254600160a01b900460ff16801561228057506001600160a01b03821660009081526004602052604090205460ff16155b156122c2576040517fd0d259760000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610979565b6122cb83612830565b60005b818110156123e25760008482815181106122ea576122ea61346f565b60200260200101519050600081600001519050600061230882610ee2565b90506001600160a01b038116612355576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610979565b60208301516040517f503c285800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0382169063503c285890602401600060405180830381600087803b1580156123b657600080fd5b505af11580156123ca573d6000803e3d6000fd5b50505050505050806123db9061349e565b90506122ce565b505050505050565b60008060001b828460200151856080015186606001518760e0015188610100015180519060200120896101200151604051602001612428919061394d565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d604001516040516020016124d69c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b60606000805b60105481101561256e57600f60006010838154811061251b5761251b61346f565b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff600160a01b909104161561255e578161255a8161349e565b9250505b6125678161349e565b90506124fa565b5060008167ffffffffffffffff81111561258a5761258a612d2a565b6040519080825280602002602001820160405280156125b3578160200160208202803683370190505b5090506000805b60105481101561268657600f6000601083815481106125db576125db61346f565b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff600160a01b909104161561267657601081815481106126235761262361346f565b6000918252602090912001546001600160a01b031683836126438161349e565b9450815181106126555761265561346f565b60200260200101906001600160a01b031690816001600160a01b0316815250505b61267f8161349e565b90506125ba565b50909392505050565b336001600160a01b038216036126e75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610979565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006127a0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316612a389092919063ffffffff16565b80519091501561209557808060200190518101906127be91906132c9565b6120955760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610979565b6000805b825181101561292f576000600660008584815181106128555761285561346f565b6020026020010151600001516001600160a01b03166001600160a01b03168152602001908152602001600020549050806000036128e85783828151811061289e5761289e61346f565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610979565b8382815181106128fa576128fa61346f565b6020026020010151602001518161291191906133cb565b61291b90846133b3565b925050806129289061349e565b9050612834565b508015612a34576129406008611f68565b60095481111561298a576009546040517f688ccf77000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610979565b600a548111156129ea57600854600a54600091906129a89084613519565b6129b29190613408565b9050806040517fe31e0f3200000000000000000000000000000000000000000000000000000000815260040161097991815260200190565b80600860020160008282546129ff9190613519565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba37890602001610788565b5050565b60606109ca848460008585843b612a915760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610979565b600080866001600160a01b03168587604051612aad9190613960565b60006040518083038185875af1925050503d8060008114612aea576040519150601f19603f3d011682016040523d82523d6000602084013e612aef565b606091505b5091509150612aff828286612b0a565b979650505050505050565b60608315612b195750816118cb565b825115612b295782518084602001fd5b8160405162461bcd60e51b81526004016109799190612cd0565b828054828255906000526020600020908101928215612ba5579160200282015b82811115612ba5578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190612b63565b50612bb1929150612c15565b5090565b828054828255906000526020600020908101928215612ba5579160200282015b82811115612ba557815473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03843516178255602090920191600190910190612bd5565b5b80821115612bb15760008155600101612c16565b600060808284031215612c3c57600080fd5b50919050565b6001600160a01b0381168114611bc357600080fd5b600060208284031215612c6957600080fd5b81356118cb81612c42565b60005b83811015612c8f578181015183820152602001612c77565b83811115612c9e576000848401525b50505050565b60008151808452612cbc816020860160208601612c74565b601f01601f19169290920160200192915050565b6020815260006118cb6020830184612ca4565b600060a08284031215612c3c57600080fd5b600060208284031215612d0757600080fd5b813567ffffffffffffffff811115612d1e57600080fd5b6109ca84828501612ce3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715612d7c57612d7c612d2a565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715612dab57612dab612d2a565b604052919050565b600067ffffffffffffffff821115612dcd57612dcd612d2a565b5060051b60200190565b600082601f830112612de857600080fd5b81356020612dfd612df883612db3565b612d82565b82815260059290921b84018101918181019086841115612e1c57600080fd5b8286015b84811015612e40578035612e3381612c42565b8352918301918301612e20565b509695505050505050565b60008060408385031215612e5e57600080fd5b823567ffffffffffffffff80821115612e7657600080fd5b612e8286838701612dd7565b9350602091508185013581811115612e9957600080fd5b85019050601f81018613612eac57600080fd5b8035612eba612df882612db3565b81815260059190911b82018301908381019088831115612ed957600080fd5b928401925b82841015612ef757833582529284019290840190612ede565b80955050505050509250929050565b600060208284031215612f1857600080fd5b813567ffffffffffffffff811115612f2f57600080fd5b6109ca84828501612dd7565b6020808252825182820181905260009190848201906040850190845b81811015612f7357835183529284019291840191600101612f57565b50909695505050505050565b60008060208385031215612f9257600080fd5b823567ffffffffffffffff80821115612faa57600080fd5b818501915085601f830112612fbe57600080fd5b813581811115612fcd57600080fd5b8660208260051b8501011115612fe257600080fd5b60209290920196919550909350505050565b6000806040838503121561300757600080fd5b823561301281612c42565b9150602083013561302281612c42565b809150509250929050565b60006040828403121561303f57600080fd5b613047612d59565b82358152602083013560208201528091505092915050565b600060c08284031215612c3c57600080fd5b60008060006060848603121561308657600080fd5b833567ffffffffffffffff81111561309d57600080fd5b6130a986828701612ce3565b9350506020840135915060408401356130c181612c42565b809150509250925092565b6020808252825182820181905260009190848201906040850190845b81811015612f735783516001600160a01b0316835292840192918401916001016130e8565b8015158114611bc357600080fd5b60006020828403121561312d57600080fd5b81356118cb8161310d565b67ffffffffffffffff81168114611bc357600080fd5b813561315981613138565b67ffffffffffffffff811690508154817fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008216178355602084013561319d81613138565b6fffffffffffffffff00000000000000008160401b16905080837fffffffffffffffffffffffffffffffff0000000000000000000000000000000084161717845560408501356131ec81613138565b77ffffffffffffffff000000000000000000000000000000008160801b1690507fffffffffffffffff00000000000000000000000000000000000000000000000081858286161784171786556060870135935061324884613138565b808460c01b16858417831717865550505050505050565b60808101823561326e81613138565b67ffffffffffffffff908116835260208401359061328b82613138565b90811660208401526040840135906132a282613138565b90811660408401526060840135906132b982613138565b8082166060850152505092915050565b6000602082840312156132db57600080fd5b81516118cb8161310d565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261331b57600080fd5b83018035915067ffffffffffffffff82111561333657600080fd5b60200191503681900382131561334b57600080fd5b9250929050565b60006020828403121561336457600080fd5b81516fffffffffffffffffffffffffffffffff811681146118cb57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156133c6576133c6613384565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561340357613403613384565b500290565b60008261343e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600067ffffffffffffffff80831681851680830382111561346657613466613384565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036134cf576134cf613384565b5060010190565b60208082528181018390526000908460408401835b86811015612e405782356134fe81612c42565b6001600160a01b0316825291830191908301906001016134eb565b60008282101561352b5761352b613384565b500390565b813561353b81612c42565b815473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0382161782555060208201356001820155604082013560028201556060820135600382015560048101608083013561359281612c42565b815473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0382161782555060a08301356135c781613138565b81547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a09190911b7bffffffffffffffff0000000000000000000000000000000000000000161790555050565b60c08101823561362581612c42565b6001600160a01b0380821684526020850135602085015260408501356040850152606085013560608501526080850135915061366082612c42565b16608083015260a083013561367481613138565b67ffffffffffffffff811660a08401525092915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126136c057600080fd5b83018035915067ffffffffffffffff8211156136db57600080fd5b6020019150600681901b360382131561334b57600080fd5b60006040828403121561370557600080fd5b61370d612d59565b823561371881612c42565b81526020928301359281019290925250919050565b600067ffffffffffffffff80831681810361374a5761374a613384565b6001019392505050565b600081518084526020808501945080840160005b8381101561379857815180516001600160a01b031688528301518388015260409096019590820190600101613768565b509495945050505050565b602081526137be60208201835167ffffffffffffffff169052565b600060208301516137db604084018267ffffffffffffffff169052565b5060408301516060830152606083015161380060808401826001600160a01b03169052565b50608083015167ffffffffffffffff811660a08401525060a083015160c083015260c083015161383460e084018215159052565b5060e0830151610100613851818501836001600160a01b03169052565b8085015191505061018061012081818601526138716101a0860184612ca4565b9250808601519050610140601f1986850301818701526138918483613754565b9350808701519150506101606138b1818701836001600160a01b03169052565b959095015193019290925250919050565b600080858511156138d257600080fd5b838611156138df57600080fd5b5050820193919092039150565b7fffffffff00000000000000000000000000000000000000000000000000000000813581811691600485101561392c5780818660040360031b1b83161692505b505092915050565b60006020828403121561394657600080fd5b5035919050565b6020815260006118cb6020830184613754565b60008251613972818460208701612c74565b919091019291505056fea164736f6c634300080f000a",
}

var EVM2EVMGEOnRampABI = EVM2EVMGEOnRampMetaData.ABI

var EVM2EVMGEOnRampBin = EVM2EVMGEOnRampMetaData.Bin

func DeployEVM2EVMGEOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, chainId uint64, destinationChainId uint64, tokens []common.Address, pools []common.Address, allowlist []common.Address, afn common.Address, config IBaseOnRampOnRampConfig, rateLimiterConfig IAggregateRateLimiterRateLimiterConfig, tokenLimitsAdmin common.Address, router common.Address, feeConfig IEVM2EVMGEOnRampDynamicFeeConfig) (common.Address, *types.Transaction, *EVM2EVMGEOnRamp, error) {
	parsed, err := EVM2EVMGEOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMGEOnRampBin), backend, chainId, destinationChainId, tokens, pools, allowlist, afn, config, rateLimiterConfig, tokenLimitsAdmin, router, feeConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMGEOnRamp{EVM2EVMGEOnRampCaller: EVM2EVMGEOnRampCaller{contract: contract}, EVM2EVMGEOnRampTransactor: EVM2EVMGEOnRampTransactor{contract: contract}, EVM2EVMGEOnRampFilterer: EVM2EVMGEOnRampFilterer{contract: contract}}, nil
}

type EVM2EVMGEOnRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMGEOnRampCaller
	EVM2EVMGEOnRampTransactor
	EVM2EVMGEOnRampFilterer
}

type EVM2EVMGEOnRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMGEOnRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMGEOnRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMGEOnRampSession struct {
	Contract     *EVM2EVMGEOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMGEOnRampCallerSession struct {
	Contract *EVM2EVMGEOnRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMGEOnRampTransactorSession struct {
	Contract     *EVM2EVMGEOnRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMGEOnRampRaw struct {
	Contract *EVM2EVMGEOnRamp
}

type EVM2EVMGEOnRampCallerRaw struct {
	Contract *EVM2EVMGEOnRampCaller
}

type EVM2EVMGEOnRampTransactorRaw struct {
	Contract *EVM2EVMGEOnRampTransactor
}

func NewEVM2EVMGEOnRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMGEOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMGEOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMGEOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRamp{address: address, abi: abi, EVM2EVMGEOnRampCaller: EVM2EVMGEOnRampCaller{contract: contract}, EVM2EVMGEOnRampTransactor: EVM2EVMGEOnRampTransactor{contract: contract}, EVM2EVMGEOnRampFilterer: EVM2EVMGEOnRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMGEOnRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMGEOnRampCaller, error) {
	contract, err := bindEVM2EVMGEOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampCaller{contract: contract}, nil
}

func NewEVM2EVMGEOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMGEOnRampTransactor, error) {
	contract, err := bindEVM2EVMGEOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampTransactor{contract: contract}, nil
}

func NewEVM2EVMGEOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMGEOnRampFilterer, error) {
	contract, err := bindEVM2EVMGEOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampFilterer{contract: contract}, nil
}

func bindEVM2EVMGEOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMGEOnRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMGEOnRamp.Contract.EVM2EVMGEOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.EVM2EVMGEOnRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.EVM2EVMGEOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMGEOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(IAggregateRateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(IAggregateRateLimiterTokenBucket)).(*IAggregateRateLimiterTokenBucket)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMGEOnRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMGEOnRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetAFN() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetAFN(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetAFN(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetAllowlist(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getAllowlist")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetAllowlist() ([]common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetAllowlist(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetAllowlist() ([]common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetAllowlist(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getAllowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetAllowlistEnabled() (bool, error) {
	return _EVM2EVMGEOnRamp.Contract.GetAllowlistEnabled(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetAllowlistEnabled() (bool, error) {
	return _EVM2EVMGEOnRamp.Contract.GetAllowlistEnabled(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetChainId() (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetChainId(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetChainId() (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetChainId(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetDestinationChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getDestinationChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetDestinationChainId() (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetDestinationChainId(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetDestinationChainId() (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetDestinationChainId(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getExpectedNextSequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetExpectedNextSequenceNumber() (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetExpectedNextSequenceNumber(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetFee(opts *bind.CallOpts, message GEConsumerEVM2AnyGEMessage) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getFee", message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetFee(message GEConsumerEVM2AnyGEMessage) (*big.Int, error) {
	return _EVM2EVMGEOnRamp.Contract.GetFee(&_EVM2EVMGEOnRamp.CallOpts, message)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetFee(message GEConsumerEVM2AnyGEMessage) (*big.Int, error) {
	return _EVM2EVMGEOnRamp.Contract.GetFee(&_EVM2EVMGEOnRamp.CallOpts, message)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetOnRampConfig(opts *bind.CallOpts) (IBaseOnRampOnRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getOnRampConfig")

	if err != nil {
		return *new(IBaseOnRampOnRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IBaseOnRampOnRampConfig)).(*IBaseOnRampOnRampConfig)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetOnRampConfig() (IBaseOnRampOnRampConfig, error) {
	return _EVM2EVMGEOnRamp.Contract.GetOnRampConfig(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetOnRampConfig() (IBaseOnRampOnRampConfig, error) {
	return _EVM2EVMGEOnRamp.Contract.GetOnRampConfig(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getPoolBySourceToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetPoolBySourceToken(&_EVM2EVMGEOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetPoolBySourceToken(&_EVM2EVMGEOnRamp.CallOpts, sourceToken)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMGEOnRamp.Contract.GetPricesForTokens(&_EVM2EVMGEOnRamp.CallOpts, tokens)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMGEOnRamp.Contract.GetPricesForTokens(&_EVM2EVMGEOnRamp.CallOpts, tokens)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetRouter() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetRouter(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetRouter(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getSenderNonce", sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetSenderNonce(&_EVM2EVMGEOnRamp.CallOpts, sender)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMGEOnRamp.Contract.GetSenderNonce(&_EVM2EVMGEOnRamp.CallOpts, sender)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetSupportedTokens(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetSupportedTokens(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMGEOnRamp.Contract.IsAFNHealthy(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMGEOnRamp.Contract.IsAFNHealthy(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) Owner() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.Owner(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMGEOnRamp.Contract.Owner(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) Paused() (bool, error) {
	return _EVM2EVMGEOnRamp.Contract.Paused(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMGEOnRamp.Contract.Paused(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMGEOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMGEOnRamp.Contract.TypeAndVersion(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMGEOnRamp.Contract.TypeAndVersion(&_EVM2EVMGEOnRamp.CallOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.AcceptOwnership(&_EVM2EVMGEOnRamp.TransactOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.AcceptOwnership(&_EVM2EVMGEOnRamp.TransactOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.AddPool(&_EVM2EVMGEOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.AddPool(&_EVM2EVMGEOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) ForwardFromRouter(opts *bind.TransactOpts, message GEConsumerEVM2AnyGEMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "forwardFromRouter", message, feeTokenAmount, originalSender)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) ForwardFromRouter(message GEConsumerEVM2AnyGEMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.ForwardFromRouter(&_EVM2EVMGEOnRamp.TransactOpts, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) ForwardFromRouter(message GEConsumerEVM2AnyGEMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.ForwardFromRouter(&_EVM2EVMGEOnRamp.TransactOpts, message, feeTokenAmount, originalSender)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.Pause(&_EVM2EVMGEOnRamp.TransactOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.Pause(&_EVM2EVMGEOnRamp.TransactOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.RemovePool(&_EVM2EVMGEOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.RemovePool(&_EVM2EVMGEOnRamp.TransactOpts, token, pool)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetAFN(&_EVM2EVMGEOnRamp.TransactOpts, afn)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetAFN(&_EVM2EVMGEOnRamp.TransactOpts, afn)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setAllowlist", allowlist)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetAllowlist(&_EVM2EVMGEOnRamp.TransactOpts, allowlist)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetAllowlist(allowlist []common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetAllowlist(&_EVM2EVMGEOnRamp.TransactOpts, allowlist)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setAllowlistEnabled", enabled)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetAllowlistEnabled(&_EVM2EVMGEOnRamp.TransactOpts, enabled)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetAllowlistEnabled(enabled bool) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetAllowlistEnabled(&_EVM2EVMGEOnRamp.TransactOpts, enabled)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetFeeAdmin(opts *bind.TransactOpts, feeAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setFeeAdmin", feeAdmin)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetFeeAdmin(feeAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetFeeAdmin(&_EVM2EVMGEOnRamp.TransactOpts, feeAdmin)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetFeeAdmin(feeAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetFeeAdmin(&_EVM2EVMGEOnRamp.TransactOpts, feeAdmin)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetFeeConfig(opts *bind.TransactOpts, feeConfig IEVM2EVMGEOnRampDynamicFeeConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setFeeConfig", feeConfig)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetFeeConfig(feeConfig IEVM2EVMGEOnRampDynamicFeeConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetFeeConfig(&_EVM2EVMGEOnRamp.TransactOpts, feeConfig)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetFeeConfig(feeConfig IEVM2EVMGEOnRampDynamicFeeConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetFeeConfig(&_EVM2EVMGEOnRamp.TransactOpts, feeConfig)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetOnRampConfig(opts *bind.TransactOpts, config IBaseOnRampOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setOnRampConfig", config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetOnRampConfig(config IBaseOnRampOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetOnRampConfig(&_EVM2EVMGEOnRamp.TransactOpts, config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetOnRampConfig(config IBaseOnRampOnRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetOnRampConfig(&_EVM2EVMGEOnRamp.TransactOpts, config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetPrices(&_EVM2EVMGEOnRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetPrices(&_EVM2EVMGEOnRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetRateLimiterConfig(&_EVM2EVMGEOnRamp.TransactOpts, config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetRateLimiterConfig(&_EVM2EVMGEOnRamp.TransactOpts, config)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetRouter(&_EVM2EVMGEOnRamp.TransactOpts, router)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetRouter(&_EVM2EVMGEOnRamp.TransactOpts, router)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMGEOnRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMGEOnRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.TransferOwnership(&_EVM2EVMGEOnRamp.TransactOpts, to)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.TransferOwnership(&_EVM2EVMGEOnRamp.TransactOpts, to)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.Unpause(&_EVM2EVMGEOnRamp.TransactOpts)
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMGEOnRamp.Contract.Unpause(&_EVM2EVMGEOnRamp.TransactOpts)
}

type EVM2EVMGEOnRampAFNSetIterator struct {
	Event *EVM2EVMGEOnRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampAFNSet)
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
		it.Event = new(EVM2EVMGEOnRampAFNSet)
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

func (it *EVM2EVMGEOnRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampAFNSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampAFNSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseAFNSet(log types.Log) (*EVM2EVMGEOnRampAFNSet, error) {
	event := new(EVM2EVMGEOnRampAFNSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampAllowListEnabledSetIterator struct {
	Event *EVM2EVMGEOnRampAllowListEnabledSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampAllowListEnabledSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampAllowListEnabledSet)
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
		it.Event = new(EVM2EVMGEOnRampAllowListEnabledSet)
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

func (it *EVM2EVMGEOnRampAllowListEnabledSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampAllowListEnabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampAllowListEnabledSet struct {
	Enabled bool
	Raw     types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterAllowListEnabledSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampAllowListEnabledSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "AllowListEnabledSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampAllowListEnabledSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "AllowListEnabledSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchAllowListEnabledSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampAllowListEnabledSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "AllowListEnabledSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampAllowListEnabledSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "AllowListEnabledSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseAllowListEnabledSet(log types.Log) (*EVM2EVMGEOnRampAllowListEnabledSet, error) {
	event := new(EVM2EVMGEOnRampAllowListEnabledSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "AllowListEnabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampAllowListSetIterator struct {
	Event *EVM2EVMGEOnRampAllowListSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampAllowListSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampAllowListSet)
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
		it.Event = new(EVM2EVMGEOnRampAllowListSet)
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

func (it *EVM2EVMGEOnRampAllowListSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampAllowListSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampAllowListSet struct {
	Allowlist []common.Address
	Raw       types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterAllowListSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampAllowListSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "AllowListSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampAllowListSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "AllowListSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchAllowListSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampAllowListSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "AllowListSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampAllowListSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "AllowListSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseAllowListSet(log types.Log) (*EVM2EVMGEOnRampAllowListSet, error) {
	event := new(EVM2EVMGEOnRampAllowListSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "AllowListSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampCCIPSendRequestedIterator struct {
	Event *EVM2EVMGEOnRampCCIPSendRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampCCIPSendRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampCCIPSendRequested)
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
		it.Event = new(EVM2EVMGEOnRampCCIPSendRequested)
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

func (it *EVM2EVMGEOnRampCCIPSendRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampCCIPSendRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampCCIPSendRequested struct {
	Message GEEVM2EVMGEMessage
	Raw     types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMGEOnRampCCIPSendRequestedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampCCIPSendRequestedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "CCIPSendRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampCCIPSendRequested) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampCCIPSendRequested)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseCCIPSendRequested(log types.Log) (*EVM2EVMGEOnRampCCIPSendRequested, error) {
	event := new(EVM2EVMGEOnRampCCIPSendRequested)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampConfigChangedIterator struct {
	Event *EVM2EVMGEOnRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampConfigChanged)
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
		it.Event = new(EVM2EVMGEOnRampConfigChanged)
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

func (it *EVM2EVMGEOnRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMGEOnRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampConfigChangedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampConfigChanged)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMGEOnRampConfigChanged, error) {
	event := new(EVM2EVMGEOnRampConfigChanged)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampFeeAdminSetIterator struct {
	Event *EVM2EVMGEOnRampFeeAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampFeeAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampFeeAdminSet)
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
		it.Event = new(EVM2EVMGEOnRampFeeAdminSet)
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

func (it *EVM2EVMGEOnRampFeeAdminSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampFeeAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampFeeAdminSet struct {
	FeeAdmin common.Address
	Raw      types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterFeeAdminSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampFeeAdminSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "FeeAdminSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampFeeAdminSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "FeeAdminSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchFeeAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampFeeAdminSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "FeeAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampFeeAdminSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "FeeAdminSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseFeeAdminSet(log types.Log) (*EVM2EVMGEOnRampFeeAdminSet, error) {
	event := new(EVM2EVMGEOnRampFeeAdminSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "FeeAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampFeeConfigSetIterator struct {
	Event *EVM2EVMGEOnRampFeeConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampFeeConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampFeeConfigSet)
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
		it.Event = new(EVM2EVMGEOnRampFeeConfigSet)
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

func (it *EVM2EVMGEOnRampFeeConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampFeeConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampFeeConfigSet struct {
	FeeConfig IEVM2EVMGEOnRampDynamicFeeConfig
	Raw       types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampFeeConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "FeeConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampFeeConfigSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "FeeConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampFeeConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "FeeConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampFeeConfigSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "FeeConfigSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseFeeConfigSet(log types.Log) (*EVM2EVMGEOnRampFeeConfigSet, error) {
	event := new(EVM2EVMGEOnRampFeeConfigSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "FeeConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampOnRampConfigSetIterator struct {
	Event *EVM2EVMGEOnRampOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampOnRampConfigSet)
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
		it.Event = new(EVM2EVMGEOnRampOnRampConfigSet)
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

func (it *EVM2EVMGEOnRampOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampOnRampConfigSet struct {
	Config IBaseOnRampOnRampConfig
	Raw    types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterOnRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampOnRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampOnRampConfigSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "OnRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "OnRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampOnRampConfigSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseOnRampConfigSet(log types.Log) (*EVM2EVMGEOnRampOnRampConfigSet, error) {
	event := new(EVM2EVMGEOnRampOnRampConfigSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "OnRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMGEOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMGEOnRampOwnershipTransferRequested)
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

func (it *EVM2EVMGEOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMGEOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampOwnershipTransferRequestedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampOwnershipTransferRequested)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMGEOnRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMGEOnRampOwnershipTransferRequested)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampOwnershipTransferredIterator struct {
	Event *EVM2EVMGEOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMGEOnRampOwnershipTransferred)
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

func (it *EVM2EVMGEOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMGEOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampOwnershipTransferredIterator{contract: _EVM2EVMGEOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampOwnershipTransferred)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMGEOnRampOwnershipTransferred, error) {
	event := new(EVM2EVMGEOnRampOwnershipTransferred)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampPausedIterator struct {
	Event *EVM2EVMGEOnRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampPaused)
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
		it.Event = new(EVM2EVMGEOnRampPaused)
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

func (it *EVM2EVMGEOnRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMGEOnRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampPausedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampPaused)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParsePaused(log types.Log) (*EVM2EVMGEOnRampPaused, error) {
	event := new(EVM2EVMGEOnRampPaused)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampPoolAddedIterator struct {
	Event *EVM2EVMGEOnRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampPoolAdded)
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
		it.Event = new(EVM2EVMGEOnRampPoolAdded)
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

func (it *EVM2EVMGEOnRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMGEOnRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampPoolAddedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampPoolAdded)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMGEOnRampPoolAdded, error) {
	event := new(EVM2EVMGEOnRampPoolAdded)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampPoolRemovedIterator struct {
	Event *EVM2EVMGEOnRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampPoolRemoved)
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
		it.Event = new(EVM2EVMGEOnRampPoolRemoved)
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

func (it *EVM2EVMGEOnRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMGEOnRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampPoolRemovedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampPoolRemoved)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMGEOnRampPoolRemoved, error) {
	event := new(EVM2EVMGEOnRampPoolRemoved)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampRouterSetIterator struct {
	Event *EVM2EVMGEOnRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampRouterSet)
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
		it.Event = new(EVM2EVMGEOnRampRouterSet)
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

func (it *EVM2EVMGEOnRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterRouterSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampRouterSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampRouterSetIterator{contract: _EVM2EVMGEOnRamp.contract, event: "RouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampRouterSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "RouterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampRouterSet)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseRouterSet(log types.Log) (*EVM2EVMGEOnRampRouterSet, error) {
	event := new(EVM2EVMGEOnRampRouterSet)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "RouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampTokenPriceChangedIterator struct {
	Event *EVM2EVMGEOnRampTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampTokenPriceChanged)
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
		it.Event = new(EVM2EVMGEOnRampTokenPriceChanged)
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

func (it *EVM2EVMGEOnRampTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMGEOnRampTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampTokenPriceChangedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampTokenPriceChanged)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMGEOnRampTokenPriceChanged, error) {
	event := new(EVM2EVMGEOnRampTokenPriceChanged)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMGEOnRampTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMGEOnRampTokensRemovedFromBucket)
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

func (it *EVM2EVMGEOnRampTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMGEOnRampTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampTokensRemovedFromBucketIterator{contract: _EVM2EVMGEOnRamp.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampTokensRemovedFromBucket)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMGEOnRampTokensRemovedFromBucket, error) {
	event := new(EVM2EVMGEOnRampTokensRemovedFromBucket)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOnRampUnpausedIterator struct {
	Event *EVM2EVMGEOnRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOnRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOnRampUnpaused)
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
		it.Event = new(EVM2EVMGEOnRampUnpaused)
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

func (it *EVM2EVMGEOnRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOnRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOnRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMGEOnRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOnRampUnpausedIterator{contract: _EVM2EVMGEOnRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOnRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOnRampUnpaused)
				if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMGEOnRampUnpaused, error) {
	event := new(EVM2EVMGEOnRampUnpaused)
	if err := _EVM2EVMGEOnRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMGEOnRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMGEOnRamp.ParseAFNSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["AllowListEnabledSet"].ID:
		return _EVM2EVMGEOnRamp.ParseAllowListEnabledSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["AllowListSet"].ID:
		return _EVM2EVMGEOnRamp.ParseAllowListSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["CCIPSendRequested"].ID:
		return _EVM2EVMGEOnRamp.ParseCCIPSendRequested(log)
	case _EVM2EVMGEOnRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMGEOnRamp.ParseConfigChanged(log)
	case _EVM2EVMGEOnRamp.abi.Events["FeeAdminSet"].ID:
		return _EVM2EVMGEOnRamp.ParseFeeAdminSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["FeeConfigSet"].ID:
		return _EVM2EVMGEOnRamp.ParseFeeConfigSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["OnRampConfigSet"].ID:
		return _EVM2EVMGEOnRamp.ParseOnRampConfigSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMGEOnRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMGEOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMGEOnRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMGEOnRamp.abi.Events["Paused"].ID:
		return _EVM2EVMGEOnRamp.ParsePaused(log)
	case _EVM2EVMGEOnRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMGEOnRamp.ParsePoolAdded(log)
	case _EVM2EVMGEOnRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMGEOnRamp.ParsePoolRemoved(log)
	case _EVM2EVMGEOnRamp.abi.Events["RouterSet"].ID:
		return _EVM2EVMGEOnRamp.ParseRouterSet(log)
	case _EVM2EVMGEOnRamp.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMGEOnRamp.ParseTokenPriceChanged(log)
	case _EVM2EVMGEOnRamp.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMGEOnRamp.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMGEOnRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMGEOnRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMGEOnRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMGEOnRampAllowListEnabledSet) Topic() common.Hash {
	return common.HexToHash("0xccf4daf6ab6430389f26b970595dab82a5881ad454770907e415ede27c8df032")
}

func (EVM2EVMGEOnRampAllowListSet) Topic() common.Hash {
	return common.HexToHash("0xf8adc5fee247b62a85f63294cb46e4be61da815e4756bc57672a83b24faf0dda")
}

func (EVM2EVMGEOnRampCCIPSendRequested) Topic() common.Hash {
	return common.HexToHash("0xaffc45517195d6499808c643bd4a7b0ffeedf95bea5852840d7bfcf63f59e821")
}

func (EVM2EVMGEOnRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMGEOnRampFeeAdminSet) Topic() common.Hash {
	return common.HexToHash("0x34efeb1f04731080ec2147b3b9c7e38f9b884e035020914e40269450f53b4815")
}

func (EVM2EVMGEOnRampFeeConfigSet) Topic() common.Hash {
	return common.HexToHash("0x92c1d3bc951c7322787dfee144ba08e1dbc36535f211ae5fbdbbc3c27c3b26be")
}

func (EVM2EVMGEOnRampOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x0447ae479bc793c12cd12089f932a4c0b4ac50f1da17f1379c3d420af34407b8")
}

func (EVM2EVMGEOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMGEOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMGEOnRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMGEOnRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMGEOnRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMGEOnRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0xc6b438e6a8a59579ce6a4406cbd203b740e0d47b458aae6596339bcd40c40d15")
}

func (EVM2EVMGEOnRampTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMGEOnRampTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMGEOnRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMGEOnRamp *EVM2EVMGEOnRamp) Address() common.Address {
	return _EVM2EVMGEOnRamp.address
}

type EVM2EVMGEOnRampInterface interface {
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetAllowlist(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	GetChainId(opts *bind.CallOpts) (uint64, error)

	GetDestinationChainId(opts *bind.CallOpts) (uint64, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetFee(opts *bind.CallOpts, message GEConsumerEVM2AnyGEMessage) (*big.Int, error)

	GetOnRampConfig(opts *bind.CallOpts) (IBaseOnRampOnRampConfig, error)

	GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error)

	GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, message GEConsumerEVM2AnyGEMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetAllowlist(opts *bind.TransactOpts, allowlist []common.Address) (*types.Transaction, error)

	SetAllowlistEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error)

	SetFeeAdmin(opts *bind.TransactOpts, feeAdmin common.Address) (*types.Transaction, error)

	SetFeeConfig(opts *bind.TransactOpts, feeConfig IEVM2EVMGEOnRampDynamicFeeConfig) (*types.Transaction, error)

	SetOnRampConfig(opts *bind.TransactOpts, config IBaseOnRampOnRampConfig) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMGEOnRampAFNSet, error)

	FilterAllowListEnabledSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampAllowListEnabledSetIterator, error)

	WatchAllowListEnabledSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampAllowListEnabledSet) (event.Subscription, error)

	ParseAllowListEnabledSet(log types.Log) (*EVM2EVMGEOnRampAllowListEnabledSet, error)

	FilterAllowListSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampAllowListSetIterator, error)

	WatchAllowListSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampAllowListSet) (event.Subscription, error)

	ParseAllowListSet(log types.Log) (*EVM2EVMGEOnRampAllowListSet, error)

	FilterCCIPSendRequested(opts *bind.FilterOpts) (*EVM2EVMGEOnRampCCIPSendRequestedIterator, error)

	WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampCCIPSendRequested) (event.Subscription, error)

	ParseCCIPSendRequested(log types.Log) (*EVM2EVMGEOnRampCCIPSendRequested, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMGEOnRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMGEOnRampConfigChanged, error)

	FilterFeeAdminSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampFeeAdminSetIterator, error)

	WatchFeeAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampFeeAdminSet) (event.Subscription, error)

	ParseFeeAdminSet(log types.Log) (*EVM2EVMGEOnRampFeeAdminSet, error)

	FilterFeeConfigSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampFeeConfigSetIterator, error)

	WatchFeeConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampFeeConfigSet) (event.Subscription, error)

	ParseFeeConfigSet(log types.Log) (*EVM2EVMGEOnRampFeeConfigSet, error)

	FilterOnRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampOnRampConfigSetIterator, error)

	WatchOnRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampOnRampConfigSet) (event.Subscription, error)

	ParseOnRampConfigSet(log types.Log) (*EVM2EVMGEOnRampOnRampConfigSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMGEOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMGEOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMGEOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMGEOnRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMGEOnRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMGEOnRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMGEOnRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMGEOnRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMGEOnRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMGEOnRampPoolRemoved, error)

	FilterRouterSet(opts *bind.FilterOpts) (*EVM2EVMGEOnRampRouterSetIterator, error)

	WatchRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampRouterSet) (event.Subscription, error)

	ParseRouterSet(log types.Log) (*EVM2EVMGEOnRampRouterSet, error)

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMGEOnRampTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMGEOnRampTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMGEOnRampTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMGEOnRampTokensRemovedFromBucket, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMGEOnRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOnRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMGEOnRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
