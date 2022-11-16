// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package any_2_evm_free_offramp

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

type AggregateRateLimiterInterfaceRateLimiterConfig struct {
	Rate     *big.Int
	Capacity *big.Int
}

type AggregateRateLimiterInterfaceTokenBucket struct {
	Rate        *big.Int
	Capacity    *big.Int
	Tokens      *big.Int
	LastUpdated *big.Int
}

type BaseOffRampInterfaceOffRampConfig struct {
	OnRampAddress                           common.Address
	PermissionLessExecutionThresholdSeconds uint32
	ExecutionDelaySeconds                   uint64
	MaxDataSize                             uint64
	MaxTokensLength                         uint64
}

type CCIPAny2EVMMessageFromSender struct {
	SourceChainId        *big.Int
	Sender               []byte
	Receiver             common.Address
	Data                 []byte
	DestPools            []common.Address
	DestTokensAndAmounts []CCIPEVMTokenAndAmount
	GasLimit             *big.Int
}

type CCIPEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

type CCIPExecutionReport struct {
	SequenceNumbers          []uint64
	TokenPerFeeCoinAddresses []common.Address
	TokenPerFeeCoin          []*big.Int
	EncodedMessages          [][]byte
	InnerProofs              [][32]byte
	InnerProofFlagBits       *big.Int
	OuterProofs              [][32]byte
	OuterProofFlagBits       *big.Int
}

var EVM2EVMFreeOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitStore\",\"outputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"}],\"name\":\"setCommitStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005dfc38038062005dfc83398101604081905262000034916200063c565b6000805460ff191681556001908a908a908a908a908a908a908a908a908a90829082908690869089903390819081620000b45760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ee57620000ee8162000385565b5050506001600160a01b0381166200011957604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015b5760405162d8548360e71b815260040160405180910390fd5b81516200017090600490602085019062000436565b5060005b82518110156200023b57600082828151811062000195576200019562000796565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001df57620001df62000796565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b0292909116919091179055506200023381620007ac565b905062000174565b5050600580546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600855600993909355600a55600b91909155875116620002b9576040516342bcdf7f60e11b815260040160405180910390fd5b5050506080958652505060a0929092528051600e8054602084015160408501516001600160401b03908116600160c01b026001600160c01b0363ffffffff909316600160a01b026001600160c01b03199094166001600160a01b03968716179390931791909116919091179091556060830151600f80549490960151821668010000000000000000026001600160801b031990941691161791909117909255600d8054919092166001600160a01b031991909116179055151560c05250620007d4975050505050505050565b336001600160a01b03821603620003df5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ab565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200048e579160200282015b828111156200048e57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000457565b506200049c929150620004a0565b5090565b5b808211156200049c5760008155600101620004a1565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715620004f257620004f2620004b7565b60405290565b6001600160a01b03811681146200050e57600080fd5b50565b80516200051e81620004f8565b919050565b80516001600160401b03811681146200051e57600080fd5b600082601f8301126200054d57600080fd5b815160206001600160401b03808311156200056c576200056c620004b7565b8260051b604051601f19603f83011681018181108482111715620005945762000594620004b7565b604052938452858101830193838101925087851115620005b357600080fd5b83870191505b84821015620005df578151620005cf81620004f8565b83529183019190830190620005b9565b979650505050505050565b600060408284031215620005fd57600080fd5b604080519081016001600160401b0381118282101715620006225762000622620004b7565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000898b036101c08112156200065d57600080fd5b8a51995060208b0151985060a0603f19820112156200067b57600080fd5b5062000686620004cd565b60408b01516200069681620004f8565b815260608b015163ffffffff81168114620006b057600080fd5b6020820152620006c360808c0162000523565b6040820152620006d660a08c0162000523565b6060820152620006e960c08c0162000523565b60808201529650620006fe60e08b0162000511565b95506200070f6101008b0162000511565b6101208b01519095506001600160401b03808211156200072e57600080fd5b6200073c8d838e016200053b565b95506101408c01519150808211156200075457600080fd5b50620007638c828d016200053b565b935050620007768b6101608c01620005ea565b9150620007876101a08b0162000511565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b600060018201620007cd57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c0516155e362000819600039600061158a0152600081816102f001526131340152600081816102cd0152818161310f01526138f101526155e36000f3fe608060405234801561001057600080fd5b50600436106102ad5760003560e01c80638456cb591161017b578063b66f0efb116100d8578063c90332841161008c578063e3d0e71211610071578063e3d0e71214610797578063eb511dd4146107aa578063f2fde38b146107bd57600080fd5b8063c903328414610771578063d30a364b1461078457600080fd5b8063c0d78655116100bd578063c0d7865514610633578063c3f909d414610646578063c9029f6a1461075e57600080fd5b8063b66f0efb146105f6578063bbe4f6db1461060757600080fd5b80639438ff631161012f578063b0f479a111610114578063b0f479a1146105bf578063b1dc65a4146105d0578063b4069b31146105e357600080fd5b80639438ff6314610591578063afcb95d71461059f57600080fd5b80638da5cb5b116101605780638da5cb5b1461052d57806390c2339b14610543578063918725431461057e57600080fd5b80638456cb591461051d57806389c065681461052557600080fd5b80633f4ba83a11610229578063681fba16116101dd57806379ba5097116101c257806379ba5097146104d057806381411834146104d857806381ff7048146104ed57600080fd5b8063681fba16146104a8578063744b92e2146104bd57600080fd5b80634741062e1161020e5780634741062e1461046c578063599f64311461048c5780635c975abb1461049d57600080fd5b80633f4ba83a146104515780634352fa9f1461045957600080fd5b8063147809b3116102805780632222dd42116102655780632222dd42146103cd5780632d0335ab146103f257806339aa92641461043e57600080fd5b8063147809b31461036c578063181f5a771461038457600080fd5b806307a22a07146102b2578063087ae6df146102c7578063108ee5fc1461031f578063142a98fc14610332575b600080fd5b6102c56102c03660046141ae565b6107d0565b005b604080517f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006020820152015b60405180910390f35b6102c561032d3660046142af565b610844565b61035f6103403660046142ed565b67ffffffffffffffff1660009081526010602052604090205460ff1690565b6040516103169190614339565b6103746108fb565b6040519015158152602001610316565b6103c06040518060400160405280602081526020017f45564d3245564d537562736372697074696f6e4f666652616d7020312e302e3081525081565b60405161031691906143d2565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610316565b6104256104003660046142af565b6001600160a01b031660009081526017602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff9091168152602001610316565b6102c561044c3660046142af565b610988565b6102c56109bf565b6102c56104673660046144a4565b6109d1565b61047f61047a366004614508565b610c26565b6040516103169190614545565b6005546001600160a01b03166103da565b60005460ff16610374565b6104b0610cee565b6040516103169190614589565b6102c56104cb3660046145ca565b610db3565b6102c56110d9565b6104e06111c1565b6040516103169190614647565b6013546011546040805163ffffffff80851682526401000000009094049093166020840152820152606001610316565b6102c5611223565b6104b0611233565b60005461010090046001600160a01b03166103da565b61054b611293565b60405161031691908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6102c561058c36600461465a565b611334565b6102c56102ad3660046146a9565b604080516001815260006020820181905291810191909152606001610316565b600c546001600160a01b03166103da565b6102c56105de366004614730565b611467565b6103da6105f13660046142af565b611a3a565b600d546001600160a01b03166103da565b6103da6106153660046142af565b6001600160a01b039081166000908152600360205260409020541690565b6102c56106413660046142af565b611b28565b6107026040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a081018252600e546001600160a01b038116825263ffffffff74010000000000000000000000000000000000000000820416602083015267ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811692820192909252600f548083166060830152680100000000000000009004909116608082015290565b6040805182516001600160a01b0316815260208084015163ffffffff16908201528282015167ffffffffffffffff908116928201929092526060808401518316908201526080928301519091169181019190915260a001610316565b6102c561076c366004614815565b611b87565b6102c561077f3660046142af565b611d26565b6102c561079236600461499a565b611d5d565b6102c56107a5366004614ae4565b611d68565b6102c56107b83660046145ca565b612621565b6102c56107cb3660046142af565b6127f9565b333014610809576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a08101515115610838576108218160a0015161280a565b61083881608001518260a001518360400151612a12565b61084181612ab0565b50565b61084c612b88565b6001600160a01b03811661088c576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa15801561095e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109829190614bb1565b15905090565b610990612b88565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6109c7612b88565b6109cf612be7565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015610a0e57506005546001600160a01b03163314155b15610a45576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610a81576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610adb576006600060078381548110610aa657610aa6614bd3565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610ad481614c31565b9050610a87565b5060005b82811015610c0b576000858281518110610afb57610afb614bd3565b6020026020010151905060006001600160a01b0316816001600160a01b031603610b51576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610b6357610b63614bd3565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610bc857610bc8614bd3565b6020026020010151604051610bf29291906001600160a01b03929092168252602082015260400190565b60405180910390a150610c0481614c31565b9050610adf565b508351610c1f906007906020870190613e83565b5050505050565b80516060908067ffffffffffffffff811115610c4457610c44613f25565b604051908082528060200260200182016040528015610c6d578160200160208202803683370190505b50915060005b81811015610ce75760066000858381518110610c9157610c91614bd3565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610ccc57610ccc614bd3565b6020908102919091010152610ce081614c31565b9050610c73565b5050919050565b60045460609067ffffffffffffffff811115610d0c57610d0c613f25565b604051908082528060200260200182016040528015610d35578160200160208202803683370190505b50905060005b600454811015610daf57610d7560048281548110610d5b57610d5b614bd3565b6000918252602090912001546001600160a01b0316611a3a565b828281518110610d8757610d87614bd3565b6001600160a01b0390921660209283029190910190910152610da881614c31565b9050610d3b565b5090565b610dbb612b88565b6004546000819003610df9576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610e87576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614610ed6576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610ee5600185614c4b565b81548110610ef557610ef5614bd3565b9060005260206000200160009054906101000a90046001600160a01b03169050600482602001516bffffffffffffffffffffffff1681548110610f3a57610f3a614bd3565b6000918252602090912001546001600160a01b03166004610f5c600186614c4b565b81548110610f6c57610f6c614bd3565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483602001516bffffffffffffffffffffffff1681548110610fc057610fc0614bd3565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000291909216179055600480548061104a5761104a614c62565b600082815260208082206000199084018101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03878116808452600383526040808520949094558351908152908716918101919091527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146111385760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060601680548060200260200160405190810160405280929190818152602001828054801561121957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116111fb575b5050505050905090565b61122b612b88565b6109cf612c83565b60606004805480602002602001604051908101604052809291908181526020018280548015611219576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116111fb575050505050905090565b6112be6040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b546060820181905242906000906112f89083614c4b565b60208401518451919250611324916113109084614c91565b856040015161131f9190614cb0565b612d0b565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561137157506005546001600160a01b03163314155b156113a8576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff116113fc576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114066008612d21565b602081015160098190558151600855600a546114229190612d0b565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916114bd91849163ffffffff851691908e908e9081908401838280828437600092019190915250612dce92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260115480825260125460ff808216602085015261010090910416928201929092529083146115785760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d617463680000000000000000000000604482015260640161112f565b6115868b8b8b8b8b8b612df1565b60007f0000000000000000000000000000000000000000000000000000000000000000156115e3576002826020015183604001516115c49190614cc8565b6115ce9190614d1c565b6115d9906001614cc8565b60ff1690506115f9565b60208201516115f3906001614cc8565b60ff1690505b8881146116485760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604482015260640161112f565b8887146116975760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604482015260640161112f565b3360009081526014602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156116da576116da61430a565b60028111156116eb576116eb61430a565b90525090506002816020015160028111156117085761170861430a565b14801561174257506016816000015160ff168154811061172a5761172a614bd3565b6000918252602090912001546001600160a01b031633145b61178e5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604482015260640161112f565b5050505050600088886040516117a5929190614d3e565b6040519081900381206117bc918c90602001614d4e565b6040516020818303038152906040528051906020012090506117dc613ef1565b604080518082019091526000808252602082015260005b88811015611a1857600060018588846020811061181257611812614bd3565b61181f91901a601b614cc8565b8d8d8681811061183157611831614bd3565b905060200201358c8c8781811061184a5761184a614bd3565b9050602002013560405160008152602001604052604051611887949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156118a9573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526014602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156118fe576118fe61430a565b600281111561190f5761190f61430a565b905250925060018360200151600281111561192c5761192c61430a565b146119795760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e0000604482015260640161112f565b8251849060ff16601f811061199057611990614bd3565b6020020151156119e25760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e6174757265000000000000000000000000604482015260640161112f565b600184846000015160ff16601f81106119fd576119fd614bd3565b9115156020909202015250611a1181614c31565b90506117f3565b5050505063ffffffff8110611a2f57611a2f614d6a565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611a8e576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611afd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b219190614d99565b9392505050565b611b30612b88565b600c805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b611b8f612b88565b80516001600160a01b0316611bd0576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600e805460208085018051604080880180516001600160a01b039098167fffffffffffffffff00000000000000000000000000000000000000000000000090961686177401000000000000000000000000000000000000000063ffffffff948516021777ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff998a16021790965560608089018051600f80546080808e018051948e167fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169290921768010000000000000000948e1694909402939093179091558451988952955190941695870195909552955187169085015291518516938301939093529151909216908201527fd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa09060a00161145c565b611d2e612b88565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b610841816001612e81565b855185518560ff16601f831115611ddb576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e65727300000000000000000000000000000000604482015260640161112f565b80600003611e45576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f7369746976650000000000000000000000000000604482015260640161112f565b818314611ed3576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e00000000000000000000000000000000000000000000000000000000606482015260840161112f565b611ede816003614c91565b8311611f46576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604482015260640161112f565b611f4e612b88565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b601554156120d557601554600090611fa690600190614c4b565b9050600060158281548110611fbd57611fbd614bd3565b6000918252602082200154601680546001600160a01b0390921693509084908110611fea57611fea614bd3565b60009182526020808320909101546001600160a01b0385811684526014909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009081169091559290911680845292208054909116905560158054919250908061205d5761205d614c62565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff19169055019055601680548061209d5761209d614c62565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff1916905501905550611f8c915050565b60005b81515181101561249a57600060146000846000015184815181106120fe576120fe614bd3565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff16600281111561213b5761213b61430a565b146121a2576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e65722061646472657373000000000000000000604482015260640161112f565b6040805180820190915260ff821681526001602082015282518051601491600091859081106121d3576121d3614bd3565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156122495761224961430a565b0217905550600091506122599050565b601460008460200151848151811061227357612273614bd3565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156122b0576122b061430a565b14612317576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161112f565b6040805180820190915260ff82168152602081016002815250601460008460200151848151811061234a5761234a614bd3565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156123c0576123c061430a565b0217905550508251805160159250839081106123de576123de614bd3565b6020908102919091018101518254600181018455600093845292829020909201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909316929092179091558201518051601691908390811061244257612442614bd3565b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390921691909117905561249381614c31565b90506120d8565b5060408101516012805460ff191660ff909216919091179055601380547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff43811682029283178555908304811693600193909260009261250e928692908216911617614db6565b92506101000a81548163ffffffff021916908363ffffffff16021790555061256d4630601360009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613558565b6011819055825180516012805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560135460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e059861260c988b98919763ffffffff909216969095919491939192614dde565b60405180910390a15050505050505050505050565b612629612b88565b6001600160a01b038216158061264657506001600160a01b038116155b1561267d576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152901561270c576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201805473ffffffffffffffffffffffffffffffffffffffff1916821790558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612801612b88565b610841816135e5565b6000805b82518110156129095760006006600085848151811061282f5761282f614bd3565b6020026020010151600001516001600160a01b03166001600160a01b03168152602001908152602001600020549050806000036128c25783828151811061287857612878614bd3565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260240161112f565b8382815181106128d4576128d4614bd3565b602002602001015160200151816128eb9190614c91565b6128f59084614cb0565b9250508061290290614c31565b905061280e565b508015612a0e5761291a6008612d21565b600954811115612964576009546040517f688ccf7700000000000000000000000000000000000000000000000000000000815260048101919091526024810182905260440161112f565b600a548111156129c457600854600a54600091906129829084614c4b565b61298c9190614e74565b9050806040517fe31e0f3200000000000000000000000000000000000000000000000000000000815260040161112f91815260200190565b80600860020160008282546129d99190614c4b565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378906020016108ef565b5050565b8151835114612a4d576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015612aaa57612a9a848281518110612a6e57612a6e614bd3565b6020026020010151848381518110612a8857612a88614bd3565b602002602001015160200151846136a1565b612aa381614c31565b9050612a50565b50505050565b60408101516001600160a01b03163b612ac65750565b600c546040517fda52b4c40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063da52b4c490612b0f908490600401614ecc565b6020604051808303816000875af1158015612b2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b529190614bb1565b610841576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005461010090046001600160a01b031633146109cf5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161112f565b60005460ff16612c395760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161112f565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60005460ff1615612cd65760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161112f565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612c663390565b6000818310612d1a5781611b21565b5090919050565b6001810154600282015442911480612d3c5750808260030154145b15612d45575050565b816001015482600201541115612d87576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482612d999190614c4b565b60018401548454919250612dc091612db19084614c91565b856002015461131f9190614cb0565b600284015550600390910155565b612dec81806020019051810190612de59190615174565b6000612e81565b505050565b6000612dfe826020614c91565b612e09856020614c91565b612e1588610144614cb0565b612e1f9190614cb0565b612e299190614cb0565b612e34906000614cb0565b9050368114612e78576040517f8e1192e10000000000000000000000000000000000000000000000000000000081526004810182905236602482015260440161112f565b50505050505050565b60005460ff1615612ed45760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161112f565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f4b9190614bb1565b15612f81576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c546001600160a01b0316612fc3576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060820151516000819003613004576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff81111561301f5761301f613f25565b60405190808252806020026020018201604052801561308d57816020015b60408051610100810182526000808252602080830182905292820181905260608083018290526080830182905260a0830181905260c083015260e0820152825260001990920191018161303d5790505b50905060008267ffffffffffffffff8111156130ab576130ab613f25565b6040519080825280602002602001820160405280156130d4578160200160208202803683370190505b50905060006131857f3997e2cfd3ccacf768662bd35c3dbf323724407d75aae3019c04f4aa59b1193f600e54604080516020808201949094527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201526001600160a01b039092166080808401919091528151808403909101815260a09092019052805191012090565b905060005b8481101561321f576000876060015182815181106131aa576131aa614bd3565b60200260200101518060200190518101906131c59190615335565b90506131d18184613718565b8483815181106131e3576131e3614bd3565b6020026020010181815250508085838151811061320257613202614bd3565b6020026020010181905250508061321890614c31565b905061318a565b50600061323f8388608001518960a001518a60c001518b60e001516137f5565b5090508580156132775750600e5474010000000000000000000000000000000000000000900463ffffffff166132758242614c4b565b105b156132ae576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8581101561354e5760008582815181106132cd576132cd614bd3565b602002602001015190506000613300826020015167ffffffffffffffff1660009081526010602052604090205460ff1690565b905060028160038111156133165761331661430a565b0361335f5760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161112f565b608082015160608301516001600160a01b0316600090815260176020526040812054909167ffffffffffffffff9081169161339c91166001615406565b67ffffffffffffffff1614905080806133c6575060038260038111156133c4576133c461430a565b145b61340e5760808301516040517fb0241f4a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161112f565b613417836138ed565b60208381015167ffffffffffffffff166000908152601090915260408120805460ff1916600117905561345161344c85613a1f565b613ced565b60208086015167ffffffffffffffff1660009081526010909152604090208054919250829160ff1916600183600381111561348e5761348e61430a565b0217905550836020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf826040516134d19190614339565b60405180910390a281156135395760608401516001600160a01b03166000908152601760205260408120805467ffffffffffffffff169161351183615429565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b505050508061354790614c31565b90506132b1565b5050505050505050565b6000808a8a8a8a8a8a8a8a8a60405160200161357c99989796959493929190615450565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b0382160361363d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161112f565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b15801561370457600080fd5b505af1158015612e78573d6000803e3d6000fd5b60008060001b828460200151856040015186606001518760a00151805190602001208860c0015160405160200161374f91906154d8565b604051602081830303815290604052805190602001208960e001518a608001516040516020016137d799989796959493929190988952602089019790975267ffffffffffffffff95861660408901526001600160a01b03948516606089015292909316608087015260a086015260c085019190915260e0840152166101008201526101200190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600d546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce90613852908c908c908c908c908c9060040161551b565b6020604051808303816000875af1158015613871573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613895919061556d565b9050600081116138d1576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a6138dd9084614c4b565b9350935050509550959350505050565b80517f00000000000000000000000000000000000000000000000000000000000000001461394d5780516040517fd44bc9eb000000000000000000000000000000000000000000000000000000008152600481019190915260240161112f565b600f5460c0820151516801000000000000000090910467ffffffffffffffff1610156139b75760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161112f565b600f5460a08201515167ffffffffffffffff909116101561084157600f5460a0820151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9092166004830152602482015260440161112f565b613a686040518060e00160405280600081526020016060815260200160006001600160a01b03168152602001606081526020016060815260200160608152602001600081525090565b60c08201515160008167ffffffffffffffff811115613a8957613a89613f25565b604051908082528060200260200182016040528015613ace57816020015b6040805180820190915260008082526020820152815260200190600190039081613aa75790505b50905060008267ffffffffffffffff811115613aec57613aec613f25565b604051908082528060200260200182016040528015613b15578160200160208202803683370190505b50905060005b83811015613c64576000613b4f8760c001518381518110613b3e57613b3e614bd3565b602002602001015160000151613e21565b905080838381518110613b6457613b64614bd3565b60200260200101906001600160a01b031690816001600160a01b031681525050806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613bc2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613be69190614d99565b848381518110613bf857613bf8614bd3565b60209081029190910101516001600160a01b03909116905260c0870151805183908110613c2757613c27614bd3565b602002602001015160200151848381518110613c4557613c45614bd3565b602090810291909101810151015250613c5d81614c31565b9050613b1b565b506040518060e00160405280866000015181526020018660400151604051602001613c9e91906001600160a01b0391909116815260200190565b604051602081830303815290604052815260200186606001516001600160a01b031681526020018660a0015181526020018281526020018381526020018660e001518152509350505050919050565b6040517f07a22a0700000000000000000000000000000000000000000000000000000000815260009030906307a22a0790613d2c908590600401614ecc565b600060405180830381600087803b158015613d4657600080fd5b505af1925050508015613d57575060015b613e19573d808015613d85576040519150601f19603f3d011682016040523d82523d6000602084013e613d8a565b606091505b50613d9481615586565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003613de55750600392915050565b806040517fcf19edfd00000000000000000000000000000000000000000000000000000000815260040161112f91906143d2565b506002919050565b6001600160a01b038181166000908152600360205260409020541680613e7e576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161112f565b919050565b828054828255906000526020600020908101928215613ee5579160200282015b82811115613ee5578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190613ea3565b50610daf929150613f10565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610daf5760008155600101613f11565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613f7757613f77613f25565b60405290565b60405160e0810167ffffffffffffffff81118282101715613f7757613f77613f25565b604051610100810167ffffffffffffffff81118282101715613f7757613f77613f25565b604051601f8201601f1916810167ffffffffffffffff81118282101715613fed57613fed613f25565b604052919050565b600067ffffffffffffffff82111561400f5761400f613f25565b50601f01601f191660200190565b600082601f83011261402e57600080fd5b813561404161403c82613ff5565b613fc4565b81815284602083860101111561405657600080fd5b816020850160208301376000918101602001919091529392505050565b6001600160a01b038116811461084157600080fd5b8035613e7e81614073565b600067ffffffffffffffff8211156140ad576140ad613f25565b5060051b60200190565b600082601f8301126140c857600080fd5b813560206140d861403c83614093565b82815260059290921b840181019181810190868411156140f757600080fd5b8286015b8481101561411b57803561410e81614073565b83529183019183016140fb565b509695505050505050565b600082601f83011261413757600080fd5b8135602061414761403c83614093565b82815260069290921b8401810191818101908684111561416657600080fd5b8286015b8481101561411b57604081890312156141835760008081fd5b61418b613f54565b813561419681614073565b8152818501358582015283529183019160400161416a565b6000602082840312156141c057600080fd5b813567ffffffffffffffff808211156141d857600080fd5b9083019060e082860312156141ec57600080fd5b6141f4613f7d565b8235815260208301358281111561420a57600080fd5b6142168782860161401d565b60208301525061422860408401614088565b604082015260608301358281111561423f57600080fd5b61424b8782860161401d565b60608301525060808301358281111561426357600080fd5b61426f878286016140b7565b60808301525060a08301358281111561428757600080fd5b61429387828601614126565b60a08301525060c083013560c082015280935050505092915050565b6000602082840312156142c157600080fd5b8135611b2181614073565b67ffffffffffffffff8116811461084157600080fd5b8035613e7e816142cc565b6000602082840312156142ff57600080fd5b8135611b21816142cc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310614374577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b8381101561439557818101518382015260200161437d565b83811115612aaa5750506000910152565b600081518084526143be81602086016020860161437a565b601f01601f19169290920160200192915050565b602081526000611b2160208301846143a6565b600082601f8301126143f657600080fd5b8135602061440661403c83614093565b82815260059290921b8401810191818101908684111561442557600080fd5b8286015b8481101561411b57803561443c81614073565b8352918301918301614429565b600082601f83011261445a57600080fd5b8135602061446a61403c83614093565b82815260059290921b8401810191818101908684111561448957600080fd5b8286015b8481101561411b578035835291830191830161448d565b600080604083850312156144b757600080fd5b823567ffffffffffffffff808211156144cf57600080fd5b6144db868387016143e5565b935060208501359150808211156144f157600080fd5b506144fe85828601614449565b9150509250929050565b60006020828403121561451a57600080fd5b813567ffffffffffffffff81111561453157600080fd5b61453d848285016143e5565b949350505050565b6020808252825182820181905260009190848201906040850190845b8181101561457d57835183529284019291840191600101614561565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561457d5783516001600160a01b0316835292840192918401916001016145a5565b600080604083850312156145dd57600080fd5b82356145e881614073565b915060208301356145f881614073565b809150509250929050565b600081518084526020808501945080840160005b8381101561463c5781516001600160a01b031687529582019590820190600101614617565b509495945050505050565b602081526000611b216020830184614603565b60006040828403121561466c57600080fd5b6040516040810181811067ffffffffffffffff8211171561468f5761468f613f25565b604052823581526020928301359281019290925250919050565b6000602082840312156146bb57600080fd5b813567ffffffffffffffff8111156146d257600080fd5b820160e08185031215611b2157600080fd5b60008083601f8401126146f657600080fd5b50813567ffffffffffffffff81111561470e57600080fd5b6020830191508360208260051b850101111561472957600080fd5b9250929050565b60008060008060008060008060e0898b03121561474c57600080fd5b606089018a81111561475d57600080fd5b8998503567ffffffffffffffff8082111561477757600080fd5b818b0191508b601f83011261478b57600080fd5b81358181111561479a57600080fd5b8c60208285010111156147ac57600080fd5b6020830199508098505060808b01359150808211156147ca57600080fd5b6147d68c838d016146e4565b909750955060a08b01359150808211156147ef57600080fd5b506147fc8b828c016146e4565b999c989b50969995989497949560c00135949350505050565b600060a0828403121561482757600080fd5b60405160a0810181811067ffffffffffffffff8211171561484a5761484a613f25565b604052823561485881614073565b8152602083013563ffffffff8116811461487157600080fd5b60208201526040830135614884816142cc565b60408201526060830135614897816142cc565b606082015260808301356148aa816142cc565b60808201529392505050565b600082601f8301126148c757600080fd5b813560206148d761403c83614093565b82815260059290921b840181019181810190868411156148f657600080fd5b8286015b8481101561411b57803561490d816142cc565b83529183019183016148fa565b600082601f83011261492b57600080fd5b8135602061493b61403c83614093565b82815260059290921b8401810191818101908684111561495a57600080fd5b8286015b8481101561411b57803567ffffffffffffffff81111561497e5760008081fd5b61498c8986838b010161401d565b84525091830191830161495e565b6000602082840312156149ac57600080fd5b813567ffffffffffffffff808211156149c457600080fd5b9083019061010082860312156149d957600080fd5b6149e1613fa0565b8235828111156149f057600080fd5b6149fc878286016148b6565b825250602083013582811115614a1157600080fd5b614a1d878286016140b7565b602083015250604083013582811115614a3557600080fd5b614a4187828601614449565b604083015250606083013582811115614a5957600080fd5b614a658782860161491a565b606083015250608083013582811115614a7d57600080fd5b614a8987828601614449565b60808301525060a083013560a082015260c083013582811115614aab57600080fd5b614ab787828601614449565b60c08301525060e083013560e082015280935050505092915050565b803560ff81168114613e7e57600080fd5b60008060008060008060c08789031215614afd57600080fd5b863567ffffffffffffffff80821115614b1557600080fd5b614b218a838b016140b7565b97506020890135915080821115614b3757600080fd5b614b438a838b016140b7565b9650614b5160408a01614ad3565b95506060890135915080821115614b6757600080fd5b614b738a838b0161401d565b9450614b8160808a016142e2565b935060a0890135915080821115614b9757600080fd5b50614ba489828a0161401d565b9150509295509295509295565b600060208284031215614bc357600080fd5b81518015158114611b2157600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006000198203614c4457614c44614c02565b5060010190565b600082821015614c5d57614c5d614c02565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000816000190483118215151615614cab57614cab614c02565b500290565b60008219821115614cc357614cc3614c02565b500190565b600060ff821660ff84168060ff03821115614ce557614ce5614c02565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff831680614d2f57614d2f614ced565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600060208284031215614dab57600080fd5b8151611b2181614073565b600063ffffffff808316818516808303821115614dd557614dd5614c02565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614e0e8184018a614603565b90508281036080840152614e228189614603565b905060ff871660a084015282810360c0840152614e3f81876143a6565b905067ffffffffffffffff851660e0840152828103610100840152614e6481856143a6565b9c9b505050505050505050505050565b600082614e8357614e83614ced565b500490565b600081518084526020808501945080840160005b8381101561463c57815180516001600160a01b031688528301518388015260409096019590820190600101614e9c565b60208152815160208201526000602083015160e06040840152614ef36101008401826143a6565b90506001600160a01b0360408501511660608401526060840151601f1980858403016080860152614f2483836143a6565b925060808601519150808584030160a0860152614f418383614603565b925060a08601519150808584030160c086015250614f5f8282614e88565b91505060c084015160e08401528091505092915050565b8051613e7e816142cc565b600082601f830112614f9257600080fd5b81516020614fa261403c83614093565b82815260059290921b84018101918181019086841115614fc157600080fd5b8286015b8481101561411b578051614fd8816142cc565b8352918301918301614fc5565b8051613e7e81614073565b600082601f83011261500157600080fd5b8151602061501161403c83614093565b82815260059290921b8401810191818101908684111561503057600080fd5b8286015b8481101561411b57805161504781614073565b8352918301918301615034565b600082601f83011261506557600080fd5b8151602061507561403c83614093565b82815260059290921b8401810191818101908684111561509457600080fd5b8286015b8481101561411b5780518352918301918301615098565b600082601f8301126150c057600080fd5b81516150ce61403c82613ff5565b8181528460208386010111156150e357600080fd5b61453d82602083016020870161437a565b600082601f83011261510557600080fd5b8151602061511561403c83614093565b82815260059290921b8401810191818101908684111561513457600080fd5b8286015b8481101561411b57805167ffffffffffffffff8111156151585760008081fd5b6151668986838b01016150af565b845250918301918301615138565b60006020828403121561518657600080fd5b815167ffffffffffffffff8082111561519e57600080fd5b9083019061010082860312156151b357600080fd5b6151bb613fa0565b8251828111156151ca57600080fd5b6151d687828601614f81565b8252506020830151828111156151eb57600080fd5b6151f787828601614ff0565b60208301525060408301518281111561520f57600080fd5b61521b87828601615054565b60408301525060608301518281111561523357600080fd5b61523f878286016150f4565b60608301525060808301518281111561525757600080fd5b61526387828601615054565b60808301525060a083015160a082015260c08301518281111561528557600080fd5b61529187828601615054565b60c08301525060e083015160e082015280935050505092915050565b600082601f8301126152be57600080fd5b815160206152ce61403c83614093565b82815260069290921b840181019181810190868411156152ed57600080fd5b8286015b8481101561411b576040818903121561530a5760008081fd5b615312613f54565b815161531d81614073565b815281850151858201528352918301916040016152f1565b60006020828403121561534757600080fd5b815167ffffffffffffffff8082111561535f57600080fd5b90830190610100828603121561537457600080fd5b61537c613fa0565b8251815261538c60208401614f76565b602082015261539d60408401614fe5565b60408201526153ae60608401614fe5565b60608201526153bf60808401614f76565b608082015260a0830151828111156153d657600080fd5b6153e2878286016150af565b60a08301525060c0830151828111156153fa57600080fd5b615291878286016152ad565b600067ffffffffffffffff808316818516808303821115614dd557614dd5614c02565b600067ffffffffffffffff80831681810361544657615446614c02565b6001019392505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b16604085015281606085015261548a8285018b614603565b9150838203608085015261549e828a614603565b915060ff881660a085015283820360c08501526154bb82886143a6565b90861660e08501528381036101008501529050614e6481856143a6565b602081526000611b216020830184614e88565b600081518084526020808501945080840160005b8381101561463c578151875295820195908201906001016154ff565b60a08152600061552e60a08301886154eb565b828103602084015261554081886154eb565b9050856040840152828103606084015261555a81866154eb565b9150508260808301529695505050505050565b60006020828403121561557f57600080fd5b5051919050565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156155ce5780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
}

var EVM2EVMFreeOffRampABI = EVM2EVMFreeOffRampMetaData.ABI

var EVM2EVMFreeOffRampBin = EVM2EVMFreeOffRampMetaData.Bin

func DeployEVM2EVMFreeOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, commitStore common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address) (common.Address, *types.Transaction, *EVM2EVMFreeOffRamp, error) {
	parsed, err := EVM2EVMFreeOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMFreeOffRampBin), backend, sourceChainId, chainId, offRampConfig, commitStore, afn, sourceTokens, pools, rateLimiterConfig, tokenLimitsAdmin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMFreeOffRamp{EVM2EVMFreeOffRampCaller: EVM2EVMFreeOffRampCaller{contract: contract}, EVM2EVMFreeOffRampTransactor: EVM2EVMFreeOffRampTransactor{contract: contract}, EVM2EVMFreeOffRampFilterer: EVM2EVMFreeOffRampFilterer{contract: contract}}, nil
}

type EVM2EVMFreeOffRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMFreeOffRampCaller
	EVM2EVMFreeOffRampTransactor
	EVM2EVMFreeOffRampFilterer
}

type EVM2EVMFreeOffRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMFreeOffRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMFreeOffRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMFreeOffRampSession struct {
	Contract     *EVM2EVMFreeOffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMFreeOffRampCallerSession struct {
	Contract *EVM2EVMFreeOffRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMFreeOffRampTransactorSession struct {
	Contract     *EVM2EVMFreeOffRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMFreeOffRampRaw struct {
	Contract *EVM2EVMFreeOffRamp
}

type EVM2EVMFreeOffRampCallerRaw struct {
	Contract *EVM2EVMFreeOffRampCaller
}

type EVM2EVMFreeOffRampTransactorRaw struct {
	Contract *EVM2EVMFreeOffRampTransactor
}

func NewEVM2EVMFreeOffRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMFreeOffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMFreeOffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMFreeOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRamp{address: address, abi: abi, EVM2EVMFreeOffRampCaller: EVM2EVMFreeOffRampCaller{contract: contract}, EVM2EVMFreeOffRampTransactor: EVM2EVMFreeOffRampTransactor{contract: contract}, EVM2EVMFreeOffRampFilterer: EVM2EVMFreeOffRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMFreeOffRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMFreeOffRampCaller, error) {
	contract, err := bindEVM2EVMFreeOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampCaller{contract: contract}, nil
}

func NewEVM2EVMFreeOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMFreeOffRampTransactor, error) {
	contract, err := bindEVM2EVMFreeOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampTransactor{contract: contract}, nil
}

func NewEVM2EVMFreeOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMFreeOffRampFilterer, error) {
	contract, err := bindEVM2EVMFreeOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampFilterer{contract: contract}, nil
}

func bindEVM2EVMFreeOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMFreeOffRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMFreeOffRamp.Contract.EVM2EVMFreeOffRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.EVM2EVMFreeOffRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.EVM2EVMFreeOffRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMFreeOffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(AggregateRateLimiterInterfaceTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(AggregateRateLimiterInterfaceTokenBucket)).(*AggregateRateLimiterInterfaceTokenBucket)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMFreeOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMFreeOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMFreeOffRamp.Contract.CcipReceive(&_EVM2EVMFreeOffRamp.CallOpts, arg0)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMFreeOffRamp.Contract.CcipReceive(&_EVM2EVMFreeOffRamp.CallOpts, arg0)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetAFN() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetAFN(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetAFN(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

	error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getChainIDs")

	outstruct := new(GetChainIDs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMFreeOffRamp.Contract.GetChainIDs(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMFreeOffRamp.Contract.GetChainIDs(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetCommitStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getCommitStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetCommitStore() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetCommitStore(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetCommitStore() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetCommitStore(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOffRampInterfaceOffRampConfig)).(*BaseOffRampInterfaceOffRampConfig)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetConfig(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetConfig(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetDestinationToken(&_EVM2EVMFreeOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetDestinationToken(&_EVM2EVMFreeOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetDestinationTokens(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetDestinationTokens(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetExecutionState(&_EVM2EVMFreeOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetExecutionState(&_EVM2EVMFreeOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetNonce(opts *bind.CallOpts, receiver common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getNonce", receiver)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetNonce(receiver common.Address) (uint64, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetNonce(&_EVM2EVMFreeOffRamp.CallOpts, receiver)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetNonce(receiver common.Address) (uint64, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetNonce(&_EVM2EVMFreeOffRamp.CallOpts, receiver)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetPool(&_EVM2EVMFreeOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetPool(&_EVM2EVMFreeOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetPoolTokens(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetPoolTokens(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetPricesForTokens(&_EVM2EVMFreeOffRamp.CallOpts, tokens)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetPricesForTokens(&_EVM2EVMFreeOffRamp.CallOpts, tokens)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetRouter() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetRouter(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetRouter(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMFreeOffRamp.Contract.IsAFNHealthy(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMFreeOffRamp.Contract.IsAFNHealthy(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMFreeOffRamp.Contract.LatestConfigDetails(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMFreeOffRamp.Contract.LatestConfigDetails(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMFreeOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMFreeOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) Owner() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.Owner(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.Owner(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) Paused() (bool, error) {
	return _EVM2EVMFreeOffRamp.Contract.Paused(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMFreeOffRamp.Contract.Paused(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.Transmitters(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMFreeOffRamp.Contract.Transmitters(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMFreeOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMFreeOffRamp.Contract.TypeAndVersion(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMFreeOffRamp.Contract.TypeAndVersion(&_EVM2EVMFreeOffRamp.CallOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.AcceptOwnership(&_EVM2EVMFreeOffRamp.TransactOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.AcceptOwnership(&_EVM2EVMFreeOffRamp.TransactOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.AddPool(&_EVM2EVMFreeOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.AddPool(&_EVM2EVMFreeOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "executeSingleMessage", message)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMFreeOffRamp.TransactOpts, message)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMFreeOffRamp.TransactOpts, message)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) ManuallyExecute(opts *bind.TransactOpts, report CCIPExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "manuallyExecute", report)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) ManuallyExecute(report CCIPExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.ManuallyExecute(&_EVM2EVMFreeOffRamp.TransactOpts, report)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) ManuallyExecute(report CCIPExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.ManuallyExecute(&_EVM2EVMFreeOffRamp.TransactOpts, report)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Pause(&_EVM2EVMFreeOffRamp.TransactOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Pause(&_EVM2EVMFreeOffRamp.TransactOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.RemovePool(&_EVM2EVMFreeOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.RemovePool(&_EVM2EVMFreeOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetAFN(&_EVM2EVMFreeOffRamp.TransactOpts, afn)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetAFN(&_EVM2EVMFreeOffRamp.TransactOpts, afn)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetCommitStore(opts *bind.TransactOpts, commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setCommitStore", commitStore)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetCommitStore(commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetCommitStore(&_EVM2EVMFreeOffRamp.TransactOpts, commitStore)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetCommitStore(commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetCommitStore(&_EVM2EVMFreeOffRamp.TransactOpts, commitStore)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetConfig(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetConfig(&_EVM2EVMFreeOffRamp.TransactOpts, config)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetConfig(&_EVM2EVMFreeOffRamp.TransactOpts, config)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetConfig0(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setConfig0", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetConfig0(&_EVM2EVMFreeOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetConfig0(&_EVM2EVMFreeOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetPrices(&_EVM2EVMFreeOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetPrices(&_EVM2EVMFreeOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMFreeOffRamp.TransactOpts, config)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMFreeOffRamp.TransactOpts, config)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetRouter(&_EVM2EVMFreeOffRamp.TransactOpts, router)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetRouter(&_EVM2EVMFreeOffRamp.TransactOpts, router)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMFreeOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMFreeOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.TransferOwnership(&_EVM2EVMFreeOffRamp.TransactOpts, to)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.TransferOwnership(&_EVM2EVMFreeOffRamp.TransactOpts, to)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Transmit(&_EVM2EVMFreeOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Transmit(&_EVM2EVMFreeOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Unpause(&_EVM2EVMFreeOffRamp.TransactOpts)
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMFreeOffRamp.Contract.Unpause(&_EVM2EVMFreeOffRamp.TransactOpts)
}

type EVM2EVMFreeOffRampAFNSetIterator struct {
	Event *EVM2EVMFreeOffRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampAFNSet)
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
		it.Event = new(EVM2EVMFreeOffRampAFNSet)
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

func (it *EVM2EVMFreeOffRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampAFNSetIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampAFNSet)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseAFNSet(log types.Log) (*EVM2EVMFreeOffRampAFNSet, error) {
	event := new(EVM2EVMFreeOffRampAFNSet)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampConfigChangedIterator struct {
	Event *EVM2EVMFreeOffRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampConfigChanged)
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
		it.Event = new(EVM2EVMFreeOffRampConfigChanged)
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

func (it *EVM2EVMFreeOffRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampConfigChangedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampConfigChanged)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMFreeOffRampConfigChanged, error) {
	event := new(EVM2EVMFreeOffRampConfigChanged)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampConfigSetIterator struct {
	Event *EVM2EVMFreeOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampConfigSet)
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
		it.Event = new(EVM2EVMFreeOffRampConfigSet)
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

func (it *EVM2EVMFreeOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampConfigSetIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampConfigSet)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseConfigSet(log types.Log) (*EVM2EVMFreeOffRampConfigSet, error) {
	event := new(EVM2EVMFreeOffRampConfigSet)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampExecutionStateChangedIterator struct {
	Event *EVM2EVMFreeOffRampExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampExecutionStateChanged)
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
		it.Event = new(EVM2EVMFreeOffRampExecutionStateChanged)
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

func (it *EVM2EVMFreeOffRampExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampExecutionStateChanged struct {
	SequenceNumber uint64
	State          uint8
	Raw            types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMFreeOffRampExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampExecutionStateChangedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampExecutionStateChanged)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseExecutionStateChanged(log types.Log) (*EVM2EVMFreeOffRampExecutionStateChanged, error) {
	event := new(EVM2EVMFreeOffRampExecutionStateChanged)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampOffRampConfigSetIterator struct {
	Event *EVM2EVMFreeOffRampOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampOffRampConfigSet)
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
		it.Event = new(EVM2EVMFreeOffRampOffRampConfigSet)
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

func (it *EVM2EVMFreeOffRampOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampOffRampConfigSet struct {
	Config BaseOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampOffRampConfigSetIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampOffRampConfigSet)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseOffRampConfigSet(log types.Log) (*EVM2EVMFreeOffRampOffRampConfigSet, error) {
	event := new(EVM2EVMFreeOffRampOffRampConfigSet)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampOffRampRouterSetIterator struct {
	Event *EVM2EVMFreeOffRampOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampOffRampRouterSet)
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
		it.Event = new(EVM2EVMFreeOffRampOffRampRouterSet)
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

func (it *EVM2EVMFreeOffRampOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampOffRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMFreeOffRampOffRampRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampOffRampRouterSetIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampOffRampRouterSet)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseOffRampRouterSet(log types.Log) (*EVM2EVMFreeOffRampOffRampRouterSet, error) {
	event := new(EVM2EVMFreeOffRampOffRampRouterSet)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMFreeOffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMFreeOffRampOwnershipTransferRequested)
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

func (it *EVM2EVMFreeOffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMFreeOffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampOwnershipTransferRequestedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampOwnershipTransferRequested)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMFreeOffRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMFreeOffRampOwnershipTransferRequested)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampOwnershipTransferredIterator struct {
	Event *EVM2EVMFreeOffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMFreeOffRampOwnershipTransferred)
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

func (it *EVM2EVMFreeOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMFreeOffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampOwnershipTransferredIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampOwnershipTransferred)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMFreeOffRampOwnershipTransferred, error) {
	event := new(EVM2EVMFreeOffRampOwnershipTransferred)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampPausedIterator struct {
	Event *EVM2EVMFreeOffRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampPaused)
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
		it.Event = new(EVM2EVMFreeOffRampPaused)
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

func (it *EVM2EVMFreeOffRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampPausedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampPaused)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParsePaused(log types.Log) (*EVM2EVMFreeOffRampPaused, error) {
	event := new(EVM2EVMFreeOffRampPaused)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampPoolAddedIterator struct {
	Event *EVM2EVMFreeOffRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampPoolAdded)
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
		it.Event = new(EVM2EVMFreeOffRampPoolAdded)
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

func (it *EVM2EVMFreeOffRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampPoolAddedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampPoolAdded)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMFreeOffRampPoolAdded, error) {
	event := new(EVM2EVMFreeOffRampPoolAdded)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampPoolRemovedIterator struct {
	Event *EVM2EVMFreeOffRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampPoolRemoved)
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
		it.Event = new(EVM2EVMFreeOffRampPoolRemoved)
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

func (it *EVM2EVMFreeOffRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampPoolRemovedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampPoolRemoved)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMFreeOffRampPoolRemoved, error) {
	event := new(EVM2EVMFreeOffRampPoolRemoved)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampTokenPriceChangedIterator struct {
	Event *EVM2EVMFreeOffRampTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampTokenPriceChanged)
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
		it.Event = new(EVM2EVMFreeOffRampTokenPriceChanged)
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

func (it *EVM2EVMFreeOffRampTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampTokenPriceChangedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampTokenPriceChanged)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMFreeOffRampTokenPriceChanged, error) {
	event := new(EVM2EVMFreeOffRampTokenPriceChanged)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMFreeOffRampTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMFreeOffRampTokensRemovedFromBucket)
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

func (it *EVM2EVMFreeOffRampTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampTokensRemovedFromBucketIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampTokensRemovedFromBucket)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMFreeOffRampTokensRemovedFromBucket, error) {
	event := new(EVM2EVMFreeOffRampTokensRemovedFromBucket)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampTransmittedIterator struct {
	Event *EVM2EVMFreeOffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampTransmitted)
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
		it.Event = new(EVM2EVMFreeOffRampTransmitted)
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

func (it *EVM2EVMFreeOffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampTransmittedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampTransmittedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampTransmitted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampTransmitted)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseTransmitted(log types.Log) (*EVM2EVMFreeOffRampTransmitted, error) {
	event := new(EVM2EVMFreeOffRampTransmitted)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMFreeOffRampUnpausedIterator struct {
	Event *EVM2EVMFreeOffRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMFreeOffRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMFreeOffRampUnpaused)
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
		it.Event = new(EVM2EVMFreeOffRampUnpaused)
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

func (it *EVM2EVMFreeOffRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMFreeOffRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMFreeOffRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMFreeOffRampUnpausedIterator{contract: _EVM2EVMFreeOffRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMFreeOffRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMFreeOffRampUnpaused)
				if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMFreeOffRampUnpaused, error) {
	event := new(EVM2EVMFreeOffRampUnpaused)
	if err := _EVM2EVMFreeOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetChainIDs struct {
	SourceChainId *big.Int
	ChainId       *big.Int
}
type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}
type LatestConfigDigestAndEpoch struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMFreeOffRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMFreeOffRamp.ParseAFNSet(log)
	case _EVM2EVMFreeOffRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMFreeOffRamp.ParseConfigChanged(log)
	case _EVM2EVMFreeOffRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMFreeOffRamp.ParseConfigSet(log)
	case _EVM2EVMFreeOffRamp.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMFreeOffRamp.ParseExecutionStateChanged(log)
	case _EVM2EVMFreeOffRamp.abi.Events["OffRampConfigSet"].ID:
		return _EVM2EVMFreeOffRamp.ParseOffRampConfigSet(log)
	case _EVM2EVMFreeOffRamp.abi.Events["OffRampRouterSet"].ID:
		return _EVM2EVMFreeOffRamp.ParseOffRampRouterSet(log)
	case _EVM2EVMFreeOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMFreeOffRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMFreeOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMFreeOffRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMFreeOffRamp.abi.Events["Paused"].ID:
		return _EVM2EVMFreeOffRamp.ParsePaused(log)
	case _EVM2EVMFreeOffRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMFreeOffRamp.ParsePoolAdded(log)
	case _EVM2EVMFreeOffRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMFreeOffRamp.ParsePoolRemoved(log)
	case _EVM2EVMFreeOffRamp.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMFreeOffRamp.ParseTokenPriceChanged(log)
	case _EVM2EVMFreeOffRamp.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMFreeOffRamp.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMFreeOffRamp.abi.Events["Transmitted"].ID:
		return _EVM2EVMFreeOffRamp.ParseTransmitted(log)
	case _EVM2EVMFreeOffRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMFreeOffRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMFreeOffRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMFreeOffRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMFreeOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMFreeOffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf")
}

func (EVM2EVMFreeOffRampOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa0")
}

func (EVM2EVMFreeOffRampOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
}

func (EVM2EVMFreeOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMFreeOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMFreeOffRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMFreeOffRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMFreeOffRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMFreeOffRampTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMFreeOffRampTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMFreeOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (EVM2EVMFreeOffRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMFreeOffRamp *EVM2EVMFreeOffRamp) Address() common.Address {
	return _EVM2EVMFreeOffRamp.address
}

type EVM2EVMFreeOffRampInterface interface {
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error)

	CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

		error)

	GetCommitStore(opts *bind.CallOpts) (common.Address, error)

	GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetNonce(opts *bind.CallOpts, receiver common.Address) (uint64, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report CCIPExecutionReport) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetCommitStore(opts *bind.TransactOpts, commitStore common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error)

	SetConfig0(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMFreeOffRampAFNSet, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMFreeOffRampConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMFreeOffRampConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMFreeOffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMFreeOffRampExecutionStateChanged, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*EVM2EVMFreeOffRampOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMFreeOffRampOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*EVM2EVMFreeOffRampOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMFreeOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMFreeOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMFreeOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMFreeOffRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMFreeOffRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMFreeOffRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMFreeOffRampPoolRemoved, error)

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMFreeOffRampTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMFreeOffRampTokensRemovedFromBucket, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMFreeOffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMFreeOffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMFreeOffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMFreeOffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
