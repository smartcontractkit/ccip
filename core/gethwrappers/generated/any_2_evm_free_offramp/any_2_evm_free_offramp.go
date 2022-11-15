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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitStore\",\"outputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"}],\"name\":\"setCommitStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005d7238038062005d7283398101604081905262000034916200063c565b6000805460ff191681556001908a908a908a908a908a908a908a908a908a90829082908690869089903390819081620000b45760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ee57620000ee8162000385565b5050506001600160a01b0381166200011957604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015b5760405162d8548360e71b815260040160405180910390fd5b81516200017090600490602085019062000436565b5060005b82518110156200023b57600082828151811062000195576200019562000796565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001df57620001df62000796565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b0292909116919091179055506200023381620007ac565b905062000174565b5050600580546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600855600993909355600a55600b91909155875116620002b9576040516342bcdf7f60e11b815260040160405180910390fd5b5050506080958652505060a0929092528051600e8054602084015160408501516001600160401b03908116600160c01b026001600160c01b0363ffffffff909316600160a01b026001600160c01b03199094166001600160a01b03968716179390931791909116919091179091556060830151600f80549490960151821668010000000000000000026001600160801b031990941691161791909117909255600d8054919092166001600160a01b031991909116179055151560c05250620007d4975050505050505050565b336001600160a01b03821603620003df5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ab565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200048e579160200282015b828111156200048e57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000457565b506200049c929150620004a0565b5090565b5b808211156200049c5760008155600101620004a1565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715620004f257620004f2620004b7565b60405290565b6001600160a01b03811681146200050e57600080fd5b50565b80516200051e81620004f8565b919050565b80516001600160401b03811681146200051e57600080fd5b600082601f8301126200054d57600080fd5b815160206001600160401b03808311156200056c576200056c620004b7565b8260051b604051601f19603f83011681018181108482111715620005945762000594620004b7565b604052938452858101830193838101925087851115620005b357600080fd5b83870191505b84821015620005df578151620005cf81620004f8565b83529183019190830190620005b9565b979650505050505050565b600060408284031215620005fd57600080fd5b604080519081016001600160401b0381118282101715620006225762000622620004b7565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000898b036101c08112156200065d57600080fd5b8a51995060208b0151985060a0603f19820112156200067b57600080fd5b5062000686620004cd565b60408b01516200069681620004f8565b815260608b015163ffffffff81168114620006b057600080fd5b6020820152620006c360808c0162000523565b6040820152620006d660a08c0162000523565b6060820152620006e960c08c0162000523565b60808201529650620006fe60e08b0162000511565b95506200070f6101008b0162000511565b6101208b01519095506001600160401b03808211156200072e57600080fd5b6200073c8d838e016200053b565b95506101408c01519150808211156200075457600080fd5b50620007638c828d016200053b565b935050620007768b6101608c01620005ea565b9150620007876101a08b0162000511565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b600060018201620007cd57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05161555962000819600039600061158a0152600081816102f001526130c60152600081816102cd015281816130a1015261388301526155596000f3fe608060405234801561001057600080fd5b50600436106102ad5760003560e01c80638456cb591161017b578063b66f0efb116100d8578063c90332841161008c578063e3d0e71211610071578063e3d0e71214610797578063eb511dd4146107aa578063f2fde38b146107bd57600080fd5b8063c903328414610771578063d30a364b1461078457600080fd5b8063c0d78655116100bd578063c0d7865514610633578063c3f909d414610646578063c9029f6a1461075e57600080fd5b8063b66f0efb146105f6578063bbe4f6db1461060757600080fd5b80639438ff631161012f578063b0f479a111610114578063b0f479a1146105bf578063b1dc65a4146105d0578063b4069b31146105e357600080fd5b80639438ff6314610591578063afcb95d71461059f57600080fd5b80638da5cb5b116101605780638da5cb5b1461052d57806390c2339b14610543578063918725431461057e57600080fd5b80638456cb591461051d57806389c065681461052557600080fd5b80633f4ba83a11610229578063681fba16116101dd57806379ba5097116101c257806379ba5097146104d057806381411834146104d857806381ff7048146104ed57600080fd5b8063681fba16146104a8578063744b92e2146104bd57600080fd5b80634741062e1161020e5780634741062e1461046c578063599f64311461048c5780635c975abb1461049d57600080fd5b80633f4ba83a146104515780634352fa9f1461045957600080fd5b8063147809b3116102805780632222dd42116102655780632222dd42146103cd5780632d0335ab146103f257806339aa92641461043e57600080fd5b8063147809b31461036c578063181f5a771461038457600080fd5b806307a22a07146102b2578063087ae6df146102c7578063108ee5fc1461031f578063142a98fc14610332575b600080fd5b6102c56102c0366004614140565b6107d0565b005b604080517f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006020820152015b60405180910390f35b6102c561032d366004614241565b610844565b61035f61034036600461427f565b67ffffffffffffffff1660009081526010602052604090205460ff1690565b60405161031691906142cb565b6103746108fb565b6040519015158152602001610316565b6103c06040518060400160405280602081526020017f45564d3245564d537562736372697074696f6e4f666652616d7020312e302e3081525081565b6040516103169190614364565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610316565b610425610400366004614241565b6001600160a01b031660009081526017602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff9091168152602001610316565b6102c561044c366004614241565b610988565b6102c56109bf565b6102c5610467366004614436565b6109d1565b61047f61047a36600461449a565b610c26565b60405161031691906144d7565b6005546001600160a01b03166103da565b60005460ff16610374565b6104b0610cee565b604051610316919061451b565b6102c56104cb36600461455c565b610db3565b6102c56110d9565b6104e06111c1565b60405161031691906145d9565b6013546011546040805163ffffffff80851682526401000000009094049093166020840152820152606001610316565b6102c5611223565b6104b0611233565b60005461010090046001600160a01b03166103da565b61054b611293565b60405161031691908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6102c561058c3660046145ec565b611334565b6102c56102ad36600461463b565b604080516001815260006020820181905291810191909152606001610316565b600c546001600160a01b03166103da565b6102c56105de3660046146c2565b611467565b6103da6105f1366004614241565b611a3a565b600d546001600160a01b03166103da565b6103da610615366004614241565b6001600160a01b039081166000908152600360205260409020541690565b6102c5610641366004614241565b611b28565b6107026040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a081018252600e546001600160a01b038116825263ffffffff74010000000000000000000000000000000000000000820416602083015267ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811692820192909252600f548083166060830152680100000000000000009004909116608082015290565b6040805182516001600160a01b0316815260208084015163ffffffff16908201528282015167ffffffffffffffff908116928201929092526060808401518316908201526080928301519091169181019190915260a001610316565b6102c561076c3660046147a7565b611b87565b6102c561077f366004614241565b611d26565b6102c561079236600461492c565b611d5d565b6102c56107a5366004614a76565b611d68565b6102c56107b836600461455c565b612621565b6102c56107cb366004614241565b6127f9565b333014610809576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a08101515115610838576108218160a0015161280a565b61083881608001518260a0015183604001516129a4565b61084181612a42565b50565b61084c612b1a565b6001600160a01b03811661088c576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa15801561095e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109829190614b43565b15905090565b610990612b1a565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6109c7612b1a565b6109cf612b79565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015610a0e57506005546001600160a01b03163314155b15610a45576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610a81576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610adb576006600060078381548110610aa657610aa6614b65565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610ad481614bc3565b9050610a87565b5060005b82811015610c0b576000858281518110610afb57610afb614b65565b6020026020010151905060006001600160a01b0316816001600160a01b031603610b51576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610b6357610b63614b65565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610bc857610bc8614b65565b6020026020010151604051610bf29291906001600160a01b03929092168252602082015260400190565b60405180910390a150610c0481614bc3565b9050610adf565b508351610c1f906007906020870190613e15565b5050505050565b80516060908067ffffffffffffffff811115610c4457610c44613eb7565b604051908082528060200260200182016040528015610c6d578160200160208202803683370190505b50915060005b81811015610ce75760066000858381518110610c9157610c91614b65565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610ccc57610ccc614b65565b6020908102919091010152610ce081614bc3565b9050610c73565b5050919050565b60045460609067ffffffffffffffff811115610d0c57610d0c613eb7565b604051908082528060200260200182016040528015610d35578160200160208202803683370190505b50905060005b600454811015610daf57610d7560048281548110610d5b57610d5b614b65565b6000918252602090912001546001600160a01b0316611a3a565b828281518110610d8757610d87614b65565b6001600160a01b0390921660209283029190910190910152610da881614bc3565b9050610d3b565b5090565b610dbb612b1a565b6004546000819003610df9576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610e87576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614610ed6576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610ee5600185614bdd565b81548110610ef557610ef5614b65565b9060005260206000200160009054906101000a90046001600160a01b03169050600482602001516bffffffffffffffffffffffff1681548110610f3a57610f3a614b65565b6000918252602090912001546001600160a01b03166004610f5c600186614bdd565b81548110610f6c57610f6c614b65565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483602001516bffffffffffffffffffffffff1681548110610fc057610fc0614b65565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000291909216179055600480548061104a5761104a614bf4565b600082815260208082206000199084018101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03878116808452600383526040808520949094558351908152908716918101919091527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146111385760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060601680548060200260200160405190810160405280929190818152602001828054801561121957602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116111fb575b5050505050905090565b61122b612b1a565b6109cf612c15565b60606004805480602002602001604051908101604052809291908181526020018280548015611219576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116111fb575050505050905090565b6112be6040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b546060820181905242906000906112f89083614bdd565b60208401518451919250611324916113109084614c23565b856040015161131f9190614c42565b612c9d565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561137157506005546001600160a01b03163314155b156113a8576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff116113fc576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114066008612cb3565b602081015160098190558151600855600a546114229190612c9d565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916114bd91849163ffffffff851691908e908e9081908401838280828437600092019190915250612d6092505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260115480825260125460ff808216602085015261010090910416928201929092529083146115785760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d617463680000000000000000000000604482015260640161112f565b6115868b8b8b8b8b8b612d83565b60007f0000000000000000000000000000000000000000000000000000000000000000156115e3576002826020015183604001516115c49190614c5a565b6115ce9190614c7f565b6115d9906001614c5a565b60ff1690506115f9565b60208201516115f3906001614c5a565b60ff1690505b8881146116485760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604482015260640161112f565b8887146116975760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604482015260640161112f565b3360009081526014602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156116da576116da61429c565b60028111156116eb576116eb61429c565b90525090506002816020015160028111156117085761170861429c565b14801561174257506016816000015160ff168154811061172a5761172a614b65565b6000918252602090912001546001600160a01b031633145b61178e5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604482015260640161112f565b5050505050600088886040516117a5929190614cc8565b6040519081900381206117bc918c90602001614cd8565b6040516020818303038152906040528051906020012090506117dc613e83565b604080518082019091526000808252602082015260005b88811015611a1857600060018588846020811061181257611812614b65565b61181f91901a601b614c5a565b8d8d8681811061183157611831614b65565b905060200201358c8c8781811061184a5761184a614b65565b9050602002013560405160008152602001604052604051611887949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156118a9573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526014602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156118fe576118fe61429c565b600281111561190f5761190f61429c565b905250925060018360200151600281111561192c5761192c61429c565b146119795760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e0000604482015260640161112f565b8251849060ff16601f811061199057611990614b65565b6020020151156119e25760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e6174757265000000000000000000000000604482015260640161112f565b600184846000015160ff16601f81106119fd576119fd614b65565b9115156020909202015250611a1181614bc3565b90506117f3565b5050505063ffffffff8110611a2f57611a2f614cf4565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611a8e576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611afd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b219190614d23565b9392505050565b611b30612b1a565b600c805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b611b8f612b1a565b80516001600160a01b0316611bd0576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600e805460208085018051604080880180516001600160a01b039098167fffffffffffffffff00000000000000000000000000000000000000000000000090961686177401000000000000000000000000000000000000000063ffffffff948516021777ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff998a16021790965560608089018051600f80546080808e018051948e167fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169290921768010000000000000000948e1694909402939093179091558451988952955190941695870195909552955187169085015291518516938301939093529151909216908201527fd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa09060a00161145c565b611d2e612b1a565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b610841816001612e13565b855185518560ff16601f831115611ddb576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e65727300000000000000000000000000000000604482015260640161112f565b80600003611e45576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f7369746976650000000000000000000000000000604482015260640161112f565b818314611ed3576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e00000000000000000000000000000000000000000000000000000000606482015260840161112f565b611ede816003614c23565b8311611f46576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604482015260640161112f565b611f4e612b1a565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b601554156120d557601554600090611fa690600190614bdd565b9050600060158281548110611fbd57611fbd614b65565b6000918252602082200154601680546001600160a01b0390921693509084908110611fea57611fea614b65565b60009182526020808320909101546001600160a01b0385811684526014909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009081169091559290911680845292208054909116905560158054919250908061205d5761205d614bf4565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff19169055019055601680548061209d5761209d614bf4565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff1916905501905550611f8c915050565b60005b81515181101561249a57600060146000846000015184815181106120fe576120fe614b65565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff16600281111561213b5761213b61429c565b146121a2576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e65722061646472657373000000000000000000604482015260640161112f565b6040805180820190915260ff821681526001602082015282518051601491600091859081106121d3576121d3614b65565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156122495761224961429c565b0217905550600091506122599050565b601460008460200151848151811061227357612273614b65565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156122b0576122b061429c565b14612317576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161112f565b6040805180820190915260ff82168152602081016002815250601460008460200151848151811061234a5761234a614b65565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156123c0576123c061429c565b0217905550508251805160159250839081106123de576123de614b65565b6020908102919091018101518254600181018455600093845292829020909201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909316929092179091558201518051601691908390811061244257612442614b65565b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390921691909117905561249381614bc3565b90506120d8565b5060408101516012805460ff191660ff909216919091179055601380547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff43811682029283178555908304811693600193909260009261250e928692908216911617614d40565b92506101000a81548163ffffffff021916908363ffffffff16021790555061256d4630601360009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a001516134ea565b6011819055825180516012805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560135460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e059861260c988b98919763ffffffff909216969095919491939192614d68565b60405180910390a15050505050505050505050565b612629612b1a565b6001600160a01b038216158061264657506001600160a01b038116155b1561267d576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152901561270c576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201805473ffffffffffffffffffffffffffffffffffffffff1916821790558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612801612b1a565b61084181613577565b6000805b82518110156129095760006006600085848151811061282f5761282f614b65565b6020026020010151600001516001600160a01b03166001600160a01b03168152602001908152602001600020549050806000036128c25783828151811061287857612878614b65565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260240161112f565b8382815181106128d4576128d4614b65565b602002602001015160200151816128eb9190614c23565b6128f59084614c42565b9250508061290290614bc3565b905061280e565b5080156129a05761291a6008612cb3565b600a54811115612956576040517f3bfa6f3800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806008600201600082825461296b9190614bdd565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378906020016108ef565b5050565b81518351146129df576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015612a3c57612a2c848281518110612a0057612a00614b65565b6020026020010151848381518110612a1a57612a1a614b65565b60200260200101516020015184613633565b612a3581614bc3565b90506129e2565b50505050565b60408101516001600160a01b03163b612a585750565b600c546040517fda52b4c40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063da52b4c490612aa1908490600401614e42565b6020604051808303816000875af1158015612ac0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ae49190614b43565b610841576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005461010090046001600160a01b031633146109cf5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161112f565b60005460ff16612bcb5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161112f565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60005460ff1615612c685760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161112f565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612bf83390565b6000818310612cac5781611b21565b5090919050565b6001810154600282015442911480612cce5750808260030154145b15612cd7575050565b816001015482600201541115612d19576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482612d2b9190614bdd565b60018401548454919250612d5291612d439084614c23565b856002015461131f9190614c42565b600284015550600390910155565b612d7e81806020019051810190612d7791906150ea565b6000612e13565b505050565b6000612d90826020614c23565b612d9b856020614c23565b612da788610144614c42565b612db19190614c42565b612dbb9190614c42565b612dc6906000614c42565b9050368114612e0a576040517f8e1192e10000000000000000000000000000000000000000000000000000000081526004810182905236602482015260440161112f565b50505050505050565b60005460ff1615612e665760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161112f565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612eb9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612edd9190614b43565b15612f13576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c546001600160a01b0316612f55576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060820151516000819003612f96576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115612fb157612fb1613eb7565b60405190808252806020026020018201604052801561301f57816020015b60408051610100810182526000808252602080830182905292820181905260608083018290526080830182905260a0830181905260c083015260e08201528252600019909201910181612fcf5790505b50905060008267ffffffffffffffff81111561303d5761303d613eb7565b604051908082528060200260200182016040528015613066578160200160208202803683370190505b50905060006131177f3997e2cfd3ccacf768662bd35c3dbf323724407d75aae3019c04f4aa59b1193f600e54604080516020808201949094527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201526001600160a01b039092166080808401919091528151808403909101815260a09092019052805191012090565b905060005b848110156131b15760008760600151828151811061313c5761313c614b65565b602002602001015180602001905181019061315791906152ab565b905061316381846136aa565b84838151811061317557613175614b65565b6020026020010181815250508085838151811061319457613194614b65565b602002602001018190525050806131aa90614bc3565b905061311c565b5060006131d18388608001518960a001518a60c001518b60e00151613787565b5090508580156132095750600e5474010000000000000000000000000000000000000000900463ffffffff166132078242614bdd565b105b15613240576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b858110156134e057600085828151811061325f5761325f614b65565b602002602001015190506000613292826020015167ffffffffffffffff1660009081526010602052604090205460ff1690565b905060028160038111156132a8576132a861429c565b036132f15760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161112f565b608082015160608301516001600160a01b0316600090815260176020526040812054909167ffffffffffffffff9081169161332e9116600161537c565b67ffffffffffffffff161490508080613358575060038260038111156133565761335661429c565b145b6133a05760808301516040517fb0241f4a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161112f565b6133a98361387f565b60208381015167ffffffffffffffff166000908152601090915260408120805460ff191660011790556133e36133de856139b1565b613c7f565b60208086015167ffffffffffffffff1660009081526010909152604090208054919250829160ff191660018360038111156134205761342061429c565b0217905550836020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf8260405161346391906142cb565b60405180910390a281156134cb5760608401516001600160a01b03166000908152601760205260408120805467ffffffffffffffff16916134a38361539f565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b50505050806134d990614bc3565b9050613243565b5050505050505050565b6000808a8a8a8a8a8a8a8a8a60405160200161350e999897969594939291906153c6565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b038216036135cf5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161112f565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b15801561369657600080fd5b505af1158015612e0a573d6000803e3d6000fd5b60008060001b828460200151856040015186606001518760a00151805190602001208860c001516040516020016136e1919061544e565b604051602081830303815290604052805190602001208960e001518a6080015160405160200161376999989796959493929190988952602089019790975267ffffffffffffffff95861660408901526001600160a01b03948516606089015292909316608087015260a086015260c085019190915260e0840152166101008201526101200190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600d546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce906137e4908c908c908c908c908c90600401615491565b6020604051808303816000875af1158015613803573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061382791906154e3565b905060008111613863576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a61386f9084614bdd565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146138df5780516040517fd44bc9eb000000000000000000000000000000000000000000000000000000008152600481019190915260240161112f565b600f5460c0820151516801000000000000000090910467ffffffffffffffff1610156139495760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161112f565b600f5460a08201515167ffffffffffffffff909116101561084157600f5460a0820151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9092166004830152602482015260440161112f565b6139fa6040518060e00160405280600081526020016060815260200160006001600160a01b03168152602001606081526020016060815260200160608152602001600081525090565b60c08201515160008167ffffffffffffffff811115613a1b57613a1b613eb7565b604051908082528060200260200182016040528015613a6057816020015b6040805180820190915260008082526020820152815260200190600190039081613a395790505b50905060008267ffffffffffffffff811115613a7e57613a7e613eb7565b604051908082528060200260200182016040528015613aa7578160200160208202803683370190505b50905060005b83811015613bf6576000613ae18760c001518381518110613ad057613ad0614b65565b602002602001015160000151613db3565b905080838381518110613af657613af6614b65565b60200260200101906001600160a01b031690816001600160a01b031681525050806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613b54573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b789190614d23565b848381518110613b8a57613b8a614b65565b60209081029190910101516001600160a01b03909116905260c0870151805183908110613bb957613bb9614b65565b602002602001015160200151848381518110613bd757613bd7614b65565b602090810291909101810151015250613bef81614bc3565b9050613aad565b506040518060e00160405280866000015181526020018660400151604051602001613c3091906001600160a01b0391909116815260200190565b604051602081830303815290604052815260200186606001516001600160a01b031681526020018660a0015181526020018281526020018381526020018660e001518152509350505050919050565b6040517f07a22a0700000000000000000000000000000000000000000000000000000000815260009030906307a22a0790613cbe908590600401614e42565b600060405180830381600087803b158015613cd857600080fd5b505af1925050508015613ce9575060015b613dab573d808015613d17576040519150601f19603f3d011682016040523d82523d6000602084013e613d1c565b606091505b50613d26816154fc565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003613d775750600392915050565b806040517fcf19edfd00000000000000000000000000000000000000000000000000000000815260040161112f9190614364565b506002919050565b6001600160a01b038181166000908152600360205260409020541680613e10576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161112f565b919050565b828054828255906000526020600020908101928215613e77579160200282015b82811115613e77578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190613e35565b50610daf929150613ea2565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610daf5760008155600101613ea3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613f0957613f09613eb7565b60405290565b60405160e0810167ffffffffffffffff81118282101715613f0957613f09613eb7565b604051610100810167ffffffffffffffff81118282101715613f0957613f09613eb7565b604051601f8201601f1916810167ffffffffffffffff81118282101715613f7f57613f7f613eb7565b604052919050565b600067ffffffffffffffff821115613fa157613fa1613eb7565b50601f01601f191660200190565b600082601f830112613fc057600080fd5b8135613fd3613fce82613f87565b613f56565b818152846020838601011115613fe857600080fd5b816020850160208301376000918101602001919091529392505050565b6001600160a01b038116811461084157600080fd5b8035613e1081614005565b600067ffffffffffffffff82111561403f5761403f613eb7565b5060051b60200190565b600082601f83011261405a57600080fd5b8135602061406a613fce83614025565b82815260059290921b8401810191818101908684111561408957600080fd5b8286015b848110156140ad5780356140a081614005565b835291830191830161408d565b509695505050505050565b600082601f8301126140c957600080fd5b813560206140d9613fce83614025565b82815260069290921b840181019181810190868411156140f857600080fd5b8286015b848110156140ad57604081890312156141155760008081fd5b61411d613ee6565b813561412881614005565b815281850135858201528352918301916040016140fc565b60006020828403121561415257600080fd5b813567ffffffffffffffff8082111561416a57600080fd5b9083019060e0828603121561417e57600080fd5b614186613f0f565b8235815260208301358281111561419c57600080fd5b6141a887828601613faf565b6020830152506141ba6040840161401a565b60408201526060830135828111156141d157600080fd5b6141dd87828601613faf565b6060830152506080830135828111156141f557600080fd5b61420187828601614049565b60808301525060a08301358281111561421957600080fd5b614225878286016140b8565b60a08301525060c083013560c082015280935050505092915050565b60006020828403121561425357600080fd5b8135611b2181614005565b67ffffffffffffffff8116811461084157600080fd5b8035613e108161425e565b60006020828403121561429157600080fd5b8135611b218161425e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310614306577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b8381101561432757818101518382015260200161430f565b83811115612a3c5750506000910152565b6000815180845261435081602086016020860161430c565b601f01601f19169290920160200192915050565b602081526000611b216020830184614338565b600082601f83011261438857600080fd5b81356020614398613fce83614025565b82815260059290921b840181019181810190868411156143b757600080fd5b8286015b848110156140ad5780356143ce81614005565b83529183019183016143bb565b600082601f8301126143ec57600080fd5b813560206143fc613fce83614025565b82815260059290921b8401810191818101908684111561441b57600080fd5b8286015b848110156140ad578035835291830191830161441f565b6000806040838503121561444957600080fd5b823567ffffffffffffffff8082111561446157600080fd5b61446d86838701614377565b9350602085013591508082111561448357600080fd5b50614490858286016143db565b9150509250929050565b6000602082840312156144ac57600080fd5b813567ffffffffffffffff8111156144c357600080fd5b6144cf84828501614377565b949350505050565b6020808252825182820181905260009190848201906040850190845b8181101561450f578351835292840192918401916001016144f3565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561450f5783516001600160a01b031683529284019291840191600101614537565b6000806040838503121561456f57600080fd5b823561457a81614005565b9150602083013561458a81614005565b809150509250929050565b600081518084526020808501945080840160005b838110156145ce5781516001600160a01b0316875295820195908201906001016145a9565b509495945050505050565b602081526000611b216020830184614595565b6000604082840312156145fe57600080fd5b6040516040810181811067ffffffffffffffff8211171561462157614621613eb7565b604052823581526020928301359281019290925250919050565b60006020828403121561464d57600080fd5b813567ffffffffffffffff81111561466457600080fd5b820160e08185031215611b2157600080fd5b60008083601f84011261468857600080fd5b50813567ffffffffffffffff8111156146a057600080fd5b6020830191508360208260051b85010111156146bb57600080fd5b9250929050565b60008060008060008060008060e0898b0312156146de57600080fd5b606089018a8111156146ef57600080fd5b8998503567ffffffffffffffff8082111561470957600080fd5b818b0191508b601f83011261471d57600080fd5b81358181111561472c57600080fd5b8c602082850101111561473e57600080fd5b6020830199508098505060808b013591508082111561475c57600080fd5b6147688c838d01614676565b909750955060a08b013591508082111561478157600080fd5b5061478e8b828c01614676565b999c989b50969995989497949560c00135949350505050565b600060a082840312156147b957600080fd5b60405160a0810181811067ffffffffffffffff821117156147dc576147dc613eb7565b60405282356147ea81614005565b8152602083013563ffffffff8116811461480357600080fd5b602082015260408301356148168161425e565b604082015260608301356148298161425e565b6060820152608083013561483c8161425e565b60808201529392505050565b600082601f83011261485957600080fd5b81356020614869613fce83614025565b82815260059290921b8401810191818101908684111561488857600080fd5b8286015b848110156140ad57803561489f8161425e565b835291830191830161488c565b600082601f8301126148bd57600080fd5b813560206148cd613fce83614025565b82815260059290921b840181019181810190868411156148ec57600080fd5b8286015b848110156140ad57803567ffffffffffffffff8111156149105760008081fd5b61491e8986838b0101613faf565b8452509183019183016148f0565b60006020828403121561493e57600080fd5b813567ffffffffffffffff8082111561495657600080fd5b90830190610100828603121561496b57600080fd5b614973613f32565b82358281111561498257600080fd5b61498e87828601614848565b8252506020830135828111156149a357600080fd5b6149af87828601614049565b6020830152506040830135828111156149c757600080fd5b6149d3878286016143db565b6040830152506060830135828111156149eb57600080fd5b6149f7878286016148ac565b606083015250608083013582811115614a0f57600080fd5b614a1b878286016143db565b60808301525060a083013560a082015260c083013582811115614a3d57600080fd5b614a49878286016143db565b60c08301525060e083013560e082015280935050505092915050565b803560ff81168114613e1057600080fd5b60008060008060008060c08789031215614a8f57600080fd5b863567ffffffffffffffff80821115614aa757600080fd5b614ab38a838b01614049565b97506020890135915080821115614ac957600080fd5b614ad58a838b01614049565b9650614ae360408a01614a65565b95506060890135915080821115614af957600080fd5b614b058a838b01613faf565b9450614b1360808a01614274565b935060a0890135915080821115614b2957600080fd5b50614b3689828a01613faf565b9150509295509295509295565b600060208284031215614b5557600080fd5b81518015158114611b2157600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006000198203614bd657614bd6614b94565b5060010190565b600082821015614bef57614bef614b94565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000816000190483118215151615614c3d57614c3d614b94565b500290565b60008219821115614c5557614c55614b94565b500190565b600060ff821660ff84168060ff03821115614c7757614c77614b94565b019392505050565b600060ff831680614cb9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600060208284031215614d3557600080fd5b8151611b2181614005565b600063ffffffff808316818516808303821115614d5f57614d5f614b94565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614d988184018a614595565b90508281036080840152614dac8189614595565b905060ff871660a084015282810360c0840152614dc98187614338565b905067ffffffffffffffff851660e0840152828103610100840152614dee8185614338565b9c9b505050505050505050505050565b600081518084526020808501945080840160005b838110156145ce57815180516001600160a01b031688528301518388015260409096019590820190600101614e12565b60208152815160208201526000602083015160e06040840152614e69610100840182614338565b90506001600160a01b0360408501511660608401526060840151601f1980858403016080860152614e9a8383614338565b925060808601519150808584030160a0860152614eb78383614595565b925060a08601519150808584030160c086015250614ed58282614dfe565b91505060c084015160e08401528091505092915050565b8051613e108161425e565b600082601f830112614f0857600080fd5b81516020614f18613fce83614025565b82815260059290921b84018101918181019086841115614f3757600080fd5b8286015b848110156140ad578051614f4e8161425e565b8352918301918301614f3b565b8051613e1081614005565b600082601f830112614f7757600080fd5b81516020614f87613fce83614025565b82815260059290921b84018101918181019086841115614fa657600080fd5b8286015b848110156140ad578051614fbd81614005565b8352918301918301614faa565b600082601f830112614fdb57600080fd5b81516020614feb613fce83614025565b82815260059290921b8401810191818101908684111561500a57600080fd5b8286015b848110156140ad578051835291830191830161500e565b600082601f83011261503657600080fd5b8151615044613fce82613f87565b81815284602083860101111561505957600080fd5b6144cf82602083016020870161430c565b600082601f83011261507b57600080fd5b8151602061508b613fce83614025565b82815260059290921b840181019181810190868411156150aa57600080fd5b8286015b848110156140ad57805167ffffffffffffffff8111156150ce5760008081fd5b6150dc8986838b0101615025565b8452509183019183016150ae565b6000602082840312156150fc57600080fd5b815167ffffffffffffffff8082111561511457600080fd5b90830190610100828603121561512957600080fd5b615131613f32565b82518281111561514057600080fd5b61514c87828601614ef7565b82525060208301518281111561516157600080fd5b61516d87828601614f66565b60208301525060408301518281111561518557600080fd5b61519187828601614fca565b6040830152506060830151828111156151a957600080fd5b6151b58782860161506a565b6060830152506080830151828111156151cd57600080fd5b6151d987828601614fca565b60808301525060a083015160a082015260c0830151828111156151fb57600080fd5b61520787828601614fca565b60c08301525060e083015160e082015280935050505092915050565b600082601f83011261523457600080fd5b81516020615244613fce83614025565b82815260069290921b8401810191818101908684111561526357600080fd5b8286015b848110156140ad57604081890312156152805760008081fd5b615288613ee6565b815161529381614005565b81528185015185820152835291830191604001615267565b6000602082840312156152bd57600080fd5b815167ffffffffffffffff808211156152d557600080fd5b9083019061010082860312156152ea57600080fd5b6152f2613f32565b8251815261530260208401614eec565b602082015261531360408401614f5b565b604082015261532460608401614f5b565b606082015261533560808401614eec565b608082015260a08301518281111561534c57600080fd5b61535887828601615025565b60a08301525060c08301518281111561537057600080fd5b61520787828601615223565b600067ffffffffffffffff808316818516808303821115614d5f57614d5f614b94565b600067ffffffffffffffff8083168181036153bc576153bc614b94565b6001019392505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526154008285018b614595565b91508382036080850152615414828a614595565b915060ff881660a085015283820360c08501526154318288614338565b90861660e08501528381036101008501529050614dee8185614338565b602081526000611b216020830184614dfe565b600081518084526020808501945080840160005b838110156145ce57815187529582019590820190600101615475565b60a0815260006154a460a0830188615461565b82810360208401526154b68188615461565b905085604084015282810360608401526154d08186615461565b9150508260808301529695505050505050565b6000602082840312156154f557600080fd5b5051919050565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156155445780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
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
