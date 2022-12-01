// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_toll_offramp

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

type CCIPEVM2EVMTollMessage struct {
	SourceChainId     *big.Int
	SequenceNumber    uint64
	Sender            common.Address
	Receiver          common.Address
	Data              []byte
	TokensAndAmounts  []CCIPEVMTokenAndAmount
	FeeTokenAndAmount CCIPEVMTokenAndAmount
	GasLimit          *big.Int
}

type CCIPEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

type CCIPExecutionReport struct {
	SequenceNumbers          []uint64
	TokenPerFeeCoinAddresses []common.Address
	TokenPerFeeCoin          []*big.Int
	FeeUpdates               []CCIPFeeUpdate
	EncodedMessages          [][]byte
	InnerProofs              [][32]byte
	InnerProofFlagBits       *big.Int
	OuterProofs              [][32]byte
	OuterProofFlagBits       *big.Int
}

type CCIPFeeUpdate struct {
	ChainId        *big.Int
	LinkPerUnitGas *big.Int
}

var EVM2EVMTollOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeTaken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitStore\",\"outputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"linkPerUnitGas\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"merkleGasShare\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount\",\"name\":\"feeTokenAndAmount\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVM2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"overheadGasToll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"}],\"name\":\"setCommitStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620066a9380380620066a983398101604081905262000035916200073f565b6000805460ff191681556001908b908b908a908a908a908a908a908a908a90829082908690869089903390819081620000b55760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ef57620000ef8162000387565b5050506001600160a01b0381166200011a57604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015c5760405162d8548360e71b815260040160405180910390fd5b81516200017190600490602085019062000438565b5060005b82518110156200023c57600082828151811062000196576200019662000834565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001e057620001e062000834565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b02929091169190911790555062000234816200084a565b905062000175565b5050600580546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600855600993909355600a55600b919091558716620002b9576040516342bcdf7f60e11b815260040160405180910390fd5b5050506080959095525060a0929092526001600160a01b0390811660c052600d80546001600160a01b0319169190921617905550151560e0525050855160168054602089015160408a01516060909a015163ffffffff9094166001600160601b0319909216919091176401000000006001600160401b039283160217600160601b600160e01b0319166c0100000000000000000000000099821699909902600160a01b600160e01b03191698909817600160a01b989092169790970217909555506200087295505050505050565b336001600160a01b03821603620003e15760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ac565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b82805482825590600052602060002090810192821562000490579160200282015b828111156200049057825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000459565b506200049e929150620004a2565b5090565b5b808211156200049e5760008155600101620004a3565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620004fa57620004fa620004b9565b604052919050565b80516001600160401b03811681146200051a57600080fd5b919050565b6000608082840312156200053257600080fd5b604051608081016001600160401b0381118282101715620005575762000557620004b9565b8060405250809150825163ffffffff811681146200057457600080fd5b8152620005846020840162000502565b6020820152620005976040840162000502565b6040820152620005aa6060840162000502565b60608201525092915050565b6001600160a01b0381168114620005cc57600080fd5b50565b80516200051a81620005b6565b60006001600160401b03821115620005f857620005f8620004b9565b5060051b60200190565b600082601f8301126200061457600080fd5b815160206200062d6200062783620005dc565b620004cf565b82815260059290921b840181019181810190868411156200064d57600080fd5b8286015b84811015620006755780516200066781620005b6565b835291830191830162000651565b509695505050505050565b600082601f8301126200069257600080fd5b81516020620006a56200062783620005dc565b82815260059290921b84018101918181019086841115620006c557600080fd5b8286015b8481101562000675578051620006df81620005b6565b8352918301918301620006c9565b6000604082840312156200070057600080fd5b604080519081016001600160401b0381118282101715620007255762000725620004b9565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000806101c08b8d0312156200076057600080fd5b8a51995060208b015198506200077a8c60408d016200051f565b97506200078a60c08c01620005cf565b96506200079a60e08c01620005cf565b9550620007ab6101008c01620005cf565b6101208c01519095506001600160401b0380821115620007ca57600080fd5b620007d88e838f0162000602565b95506101408d0151915080821115620007f057600080fd5b50620007ff8d828e0162000680565b935050620008128c6101608d01620006ed565b9150620008236101a08c01620005cf565b90509295989b9194979a5092959850565b634e487b7160e01b600052603260045260246000fd5b6000600182016200086b57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e051615dda620008cf6000396000611639015260008181611d6a015261318901526000818161030b01526131630152600081816102e801528181611d470152818161313e0152613a200152615dda6000f3fe608060405234801561001057600080fd5b50600436106102c85760003560e01c806389c065681161017b578063b4069b31116100d8578063c3f909d41161008c578063e3d0e71211610071578063e3d0e7121461076c578063eb511dd41461077f578063f2fde38b1461079257600080fd5b8063c3f909d41461067c578063c90332841461075957600080fd5b8063bbe4f6db116100bd578063bbe4f6db1461062a578063bc29705a14610656578063c0d786551461066957600080fd5b8063b4069b3114610606578063b66f0efb1461061957600080fd5b8063a8b640c11161012f578063afcb95d711610114578063afcb95d7146105c2578063b0f479a1146105e2578063b1dc65a4146105f357600080fd5b8063a8b640c11461058f578063a8e91321146105af57600080fd5b806390c2339b1161016057806390c2339b14610533578063918725431461056e5780639438ff631461058157600080fd5b806389c06568146105155780638da5cb5b1461051d57600080fd5b80634741062e11610229578063744b92e2116101dd57806381411834116101c257806381411834146104c857806381ff7048146104dd5780638456cb591461050d57600080fd5b8063744b92e2146104ad57806379ba5097146104c057600080fd5b80635b4dc8121161020e5780635b4dc8121461046c5780635c975abb1461048d578063681fba161461049857600080fd5b80634741062e1461043b578063599f64311461045b57600080fd5b8063181f5a771161028057806339aa92641161026557806339aa92641461040d5780633f4ba83a146104205780634352fa9f1461042857600080fd5b8063181f5a771461039f5780632222dd42146103e857600080fd5b8063108ee5fc116102b1578063108ee5fc1461033a578063142a98fc1461034d578063147809b31461038757600080fd5b806307a22a07146102cd578063087ae6df146102e2575b600080fd5b6102e06102db3660046147b5565b6107a5565b005b604080517f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006020820152015b60405180910390f35b6102e06103483660046148b6565b610819565b61037a61035b3660046148f4565b67ffffffffffffffff166000908152600e602052604090205460ff1690565b6040516103319190614940565b61038f6108d0565b6040519015158152602001610331565b6103db6040518060400160405280601881526020017f45564d3245564d546f6c6c4f666652616d7020312e302e30000000000000000081525081565b60405161033191906149d9565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610331565b6102e061041b3660046148b6565b61095d565b6102e0610994565b6102e0610436366004614a47565b6109a6565b61044e610449366004614aab565b610bfb565b6040516103319190614ae8565b6005546001600160a01b03166103f5565b61047f61047a366004614b2c565b610cc3565b604051908152602001610331565b60005460ff1661038f565b6104a0610d74565b6040516103319190614c24565b6102e06104bb366004614c65565b610e39565b6102e061117d565b6104d0611265565b6040516103319190614ce2565b601154600f546040805163ffffffff80851682526401000000009094049093166020840152820152606001610331565b6102e06112c7565b6104a06112d7565b60005461010090046001600160a01b03166103f5565b61053b611337565b60405161033191908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6102e061057c366004614d25565b6113d8565b6102e06102c8366004614d41565b61047f61059d366004614d7c565b60156020526000908152604090205481565b6102e06105bd366004614edd565b61150b565b604080516001815260006020820181905291810191909152606001610331565b600c546001600160a01b03166103f5565b6102e0610601366004615086565b611516565b6103f56106143660046148b6565b611ae9565b600d546001600160a01b03166103f5565b6103f56106383660046148b6565b6001600160a01b039081166000908152600360205260409020541690565b6102e061066436600461516b565b611bd7565b6102e06106773660046148b6565b611d0d565b61070e604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260165463ffffffff8116825267ffffffffffffffff6401000000008204811660208401526c010000000000000000000000008204811693830193909352740100000000000000000000000000000000000000009004909116606082015290565b6040516103319190815163ffffffff16815260208083015167ffffffffffffffff90811691830191909152604080840151821690830152606092830151169181019190915260800190565b6102e06107673660046148b6565b611dc1565b6102e061077a36600461520a565b611df8565b6102e061078d366004614c65565b6126ed565b6102e06107a03660046148b6565b6128c5565b3330146107de576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a0810151511561080d576107f68160a001516128d6565b61080d81608001518260a001518360400151612ade565b61081681612b7c565b50565b610821612c54565b6001600160a01b038116610861576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa158015610933573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095791906152d7565b15905090565b610965612c54565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b61099c612c54565b6109a4612cb3565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b0316141580156109e357506005546001600160a01b03163314155b15610a1a576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610a56576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610ab0576006600060078381548110610a7b57610a7b6152f9565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610aa981615357565b9050610a5c565b5060005b82811015610be0576000858281518110610ad057610ad06152f9565b6020026020010151905060006001600160a01b0316816001600160a01b031603610b26576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610b3857610b386152f9565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610b9d57610b9d6152f9565b6020026020010151604051610bc79291906001600160a01b03929092168252602082015260400190565b60405180910390a150610bd981615357565b9050610ab4565b508351610bf490600790602087019061444f565b5050505050565b80516060908067ffffffffffffffff811115610c1957610c196144f1565b604051908082528060200260200182016040528015610c42578160200160208202803683370190505b50915060005b81811015610cbc5760066000858381518110610c6657610c666152f9565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610ca157610ca16152f9565b6020908102919091010152610cb581615357565b9050610c48565b5050919050565b6000808260800151518360a001515160206014610ce0919061538f565b610cea91906153a7565b610cf590608661538f565b610cff919061538f565b90506000610d0e6010836153a7565b9050610a28611bbc8560a00151516001610d28919061538f565b610d3490618aac6153a7565b6156b8610d41898661538f565b610d4b919061538f565b610d55919061538f565b610d5f919061538f565b610d69919061538f565b925050505b92915050565b60045460609067ffffffffffffffff811115610d9257610d926144f1565b604051908082528060200260200182016040528015610dbb578160200160208202803683370190505b50905060005b600454811015610e3557610dfb60048281548110610de157610de16152f9565b6000918252602090912001546001600160a01b0316611ae9565b828281518110610e0d57610e0d6152f9565b6001600160a01b0390921660209283029190910190910152610e2e81615357565b9050610dc1565b5090565b610e41612c54565b6004546000819003610e7f576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610f0d576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614610f5c576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610f6b6001856153e4565b81548110610f7b57610f7b6152f9565b9060005260206000200160009054906101000a90046001600160a01b03169050600482602001516bffffffffffffffffffffffff1681548110610fc057610fc06152f9565b6000918252602090912001546001600160a01b03166004610fe26001866153e4565b81548110610ff257610ff26152f9565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483602001516bffffffffffffffffffffffff1681548110611046576110466152f9565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff90921674010000000000000000000000000000000000000000029190921617905560048054806110d0576110d06153fb565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9084018101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03878116808452600383526040808520949094558351908152908716918101919091527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146111dc5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b606060148054806020026020016040519081016040528092919081815260200182805480156112bd57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161129f575b5050505050905090565b6112cf612c54565b6109a4612d4f565b606060048054806020026020016040519081016040528092919081815260200182805480156112bd576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161129f575050505050905090565b6113626040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b5460608201819052429060009061139c90836153e4565b602084015184519192506113c8916113b490846153a7565b85604001516113c3919061538f565b612dd7565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561141557506005546001600160a01b03163314155b1561144c576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff116114a0576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114aa6008612ded565b602081015160098190558151600855600a546114c69190612dd7565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b610816816001612e9a565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161156c91849163ffffffff851691908e908e90819084018382808284376000920191909152506135c992505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600f5480825260105460ff808216602085015261010090910416928201929092529083146116275760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d61746368000000000000000000000060448201526064016111d3565b6116358b8b8b8b8b8b6135ec565b60007f00000000000000000000000000000000000000000000000000000000000000001561169257600282602001518360400151611673919061542a565b61167d919061547e565b61168890600161542a565b60ff1690506116a8565b60208201516116a290600161542a565b60ff1690505b8881146116f75760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e61747572657300000000000060448201526064016111d3565b8887146117465760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e000060448201526064016111d3565b3360009081526012602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561178957611789614911565b600281111561179a5761179a614911565b90525090506002816020015160028111156117b7576117b7614911565b1480156117f157506014816000015160ff16815481106117d9576117d96152f9565b6000918252602090912001546001600160a01b031633145b61183d5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d6974746572000000000000000060448201526064016111d3565b5050505050600088886040516118549291906154a0565b60405190819003812061186b918c906020016154b0565b60405160208183030381529060405280519060200120905061188b6144bd565b604080518082019091526000808252602082015260005b88811015611ac75760006001858884602081106118c1576118c16152f9565b6118ce91901a601b61542a565b8d8d868181106118e0576118e06152f9565b905060200201358c8c878181106118f9576118f96152f9565b9050602002013560405160008152602001604052604051611936949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611958573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526012602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156119ad576119ad614911565b60028111156119be576119be614911565b90525092506001836020015160028111156119db576119db614911565b14611a285760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e000060448201526064016111d3565b8251849060ff16601f8110611a3f57611a3f6152f9565b602002015115611a915760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e617475726500000000000000000000000060448201526064016111d3565b600184846000015160ff16601f8110611aac57611aac6152f9565b9115156020909202015250611ac081615357565b90506118a2565b5050505063ffffffff8110611ade57611ade6154cc565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611b3d576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611bac573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bd091906154fb565b9392505050565b611bdf612c54565b80516016805460208085018051604080880180516060808b01805167ffffffffffffffff90811674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff9482166c0100000000000000000000000002949094167fffffffff00000000000000000000000000000000ffffffffffffffffffffffff978216640100000000027fffffffffffffffffffffffffffffffffffffffff000000000000000000000000909b1663ffffffff909d169c8d179a909a17969096169890981791909117909755815197885292518216938701939093529051811691850191909152905116908201527f1a9ff9caccd597226ff8c393ab44e6b57669905d9b16f1cdb2ac267253dbf27c90608001611500565b611d15612c54565b600c805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03838116918217909255604080517f000000000000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000909316602084015290917f1090eec34bfa4c1a6201bdd67d36d1e4527d87f93cd925fe8b6065a953e1e636910160405180910390a250565b611dc9612c54565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b855185518560ff16601f831115611e6b576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064016111d3565b80600003611ed5576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016111d3565b818314611f63576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084016111d3565b611f6e8160036153a7565b8311611fd6576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016111d3565b611fde612c54565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b601354156121a157601354600090612036906001906153e4565b905060006013828154811061204d5761204d6152f9565b6000918252602082200154601480546001600160a01b039092169350908490811061207a5761207a6152f9565b60009182526020808320909101546001600160a01b0385811684526012909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000908116909155929091168084529220805490911690556013805491925090806120ed576120ed6153fb565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908101805473ffffffffffffffffffffffffffffffffffffffff19169055019055601480548061214b5761214b6153fb565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908101805473ffffffffffffffffffffffffffffffffffffffff191690550190555061201c915050565b60005b81515181101561256657600060126000846000015184815181106121ca576121ca6152f9565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff16600281111561220757612207614911565b1461226e576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016111d3565b6040805180820190915260ff8216815260016020820152825180516012916000918590811061229f5761229f6152f9565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561231557612315614911565b0217905550600091506123259050565b601260008460200151848151811061233f5761233f6152f9565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff16600281111561237c5761237c614911565b146123e3576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016111d3565b6040805180820190915260ff821681526020810160028152506012600084602001518481518110612416576124166152f9565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561248c5761248c614911565b0217905550508251805160139250839081106124aa576124aa6152f9565b6020908102919091018101518254600181018455600093845292829020909201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909316929092179091558201518051601491908390811061250e5761250e6152f9565b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390921691909117905561255f81615357565b90506121a4565b5060408101516010805460ff191660ff909216919091179055601180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff4381168202928317855590830481169360019390926000926125da928692908216911617615518565b92506101000a81548163ffffffff021916908363ffffffff1602179055506126394630601160009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a0015161367c565b600f819055825180516010805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560115460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05986126d8988b98919763ffffffff909216969095919491939192615540565b60405180910390a15050505050505050505050565b6126f5612c54565b6001600160a01b038216158061271257506001600160a01b038116155b15612749576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156127d8576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201805473ffffffffffffffffffffffffffffffffffffffff1916821790558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b6128cd612c54565b61081681613709565b6000805b82518110156129d5576000600660008584815181106128fb576128fb6152f9565b6020026020010151600001516001600160a01b03166001600160a01b031681526020019081526020016000205490508060000361298e57838281518110612944576129446152f9565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201526024016111d3565b8382815181106129a0576129a06152f9565b602002602001015160200151816129b791906153a7565b6129c1908461538f565b925050806129ce90615357565b90506128da565b508015612ada576129e66008612ded565b600954811115612a30576009546040517f688ccf770000000000000000000000000000000000000000000000000000000081526004810191909152602481018290526044016111d3565b600a54811115612a9057600854600a5460009190612a4e90846153e4565b612a5891906155d6565b9050806040517fe31e0f320000000000000000000000000000000000000000000000000000000081526004016111d391815260200190565b8060086002016000828254612aa591906153e4565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378906020016108c4565b5050565b8151835114612b19576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015612b7657612b66848281518110612b3a57612b3a6152f9565b6020026020010151848381518110612b5457612b546152f9565b602002602001015160200151846137c5565b612b6f81615357565b9050612b1c565b50505050565b60408101516001600160a01b03163b612b925750565b600c546040517fda52b4c40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063da52b4c490612bdb908490600401615637565b6020604051808303816000875af1158015612bfa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c1e91906152d7565b610816576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005461010090046001600160a01b031633146109a45760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016111d3565b60005460ff16612d055760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016111d3565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60005460ff1615612da25760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016111d3565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612d323390565b6000818310612de65781611bd0565b5090919050565b6001810154600282015442911480612e085750808260030154145b15612e11575050565b816001015482600201541115612e53576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482612e6591906153e4565b60018401548454919250612e8c91612e7d90846153a7565b85600201546113c3919061538f565b600284015550600390910155565b60005460ff1615612eed5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016111d3565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f6491906152d7565b15612f9a576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c546001600160a01b0316612fdc576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608082015151600081900361301d576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115613038576130386144f1565b6040519080825280602002602001820160405280156130bf57816020015b6130ac60408051610100810182526000808252602080830182905282840182905260608084018390526080840181905260a084015283518085019094528184528301529060c08201908152602001600081525090565b8152602001906001900390816130565790505b50905060008267ffffffffffffffff8111156130dd576130dd6144f1565b604051908082528060200260200182016040528015613106578160200160208202803683370190505b50905060006131d47fb9b8993db34ae003b2aacdae4cdef2888717531ab95157174f8f0dbf076b5e58604080516020808201939093527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166080808301919091528251808303909101815260a0909101909152805191012090565b905060005b8481101561326e576000876080015182815181106131f9576131f96152f9565b602002602001015180602001905181019061321491906157db565b9050613220818461383c565b848381518110613232576132326152f9565b60200260200101818152505080858381518110613251576132516152f9565b6020026020010181905250508061326790615357565b90506131d9565b50600080613290848960a001518a60c001518b60e001518c6101000151613924565b601654919350915060009063ffffffff166132ab84426153e4565b11905060005b878110156135bd5760008782815181106132cd576132cd6152f9565b602002602001015190506000613300826020015167ffffffffffffffff166000908152600e602052604090205460ff1690565b9050600281600381111561331657613316614911565b0361335f5760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016111d3565b8a15806133695750835b806133855750600381600381111561338357613383614911565b145b6133bb576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6133c482613a1c565b6000808260038111156133d9576133d9614911565b1480156133e457508b155b1561343d576134008a51876133f991906155d6565b8e85613b78565b905060006134158460c0015160000151613d65565b90506134228183306137c5565b818460c0015160200181815161343891906153e4565b905250505b600082600381111561345157613451614911565b1461348d5760208084015167ffffffffffffffff16600090815260158252604090205460c085015190910180516134899083906153e4565b9052505b60208381015167ffffffffffffffff166000908152600e90915260408120805460ff191660011790556134c76134c285613dc7565b614090565b60208086015167ffffffffffffffff166000908152600e909152604090208054919250829160ff1916600183600381111561350457613504614911565b0217905550600083600381111561351d5761351d614911565b14801561353b5750600381600381111561353957613539614911565b145b156135625760208085015167ffffffffffffffff1660009081526015909152604090208290555b836020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf826040516135a09190614940565b60405180910390a250505050806135b690615357565b90506132b1565b50505050505050505050565b6135e7818060200190518101906135e09190615aea565b6000612e9a565b505050565b60006135f98260206153a7565b6136048560206153a7565b6136108861014461538f565b61361a919061538f565b613624919061538f565b61362f90600061538f565b9050368114613673576040517f8e1192e1000000000000000000000000000000000000000000000000000000008152600481018290523660248201526044016111d3565b50505050505050565b6000808a8a8a8a8a8a8a8a8a6040516020016136a099989796959493929190615c47565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b038216036137615760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016111d3565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b15801561382857600080fd5b505af1158015613673573d6000803e3d6000fd5b60008060001b828460200151856040015186606001518760800151805190602001208860a001516040516020016138739190615ccf565b604051602081830303815290604052805190602001208960e001518a60c00151604051602001613906999897969594939291909889526020808a019890985267ffffffffffffffff9690961660408901526001600160a01b039485166060890152928416608088015260a087019190915260c086015260e085015281511661010084015201516101208201526101400190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600d546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce90613981908c908c908c908c908c90600401615d12565b6020604051808303816000875af11580156139a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139c49190615d64565b905060008111613a00576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a613a0c90846153e4565b9350935050509550959350505050565b80517f000000000000000000000000000000000000000000000000000000000000000014613a7c5780516040517fd44bc9eb00000000000000000000000000000000000000000000000000000000815260048101919091526024016111d3565b60165460a0820151517401000000000000000000000000000000000000000090910467ffffffffffffffff161015613af25760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016111d3565b6016546080820151516c0100000000000000000000000090910467ffffffffffffffff161015610816576016546080820151516040517f869337890000000000000000000000000000000000000000000000000000000081526c0100000000000000000000000090920467ffffffffffffffff16600483015260248201526044016111d3565b6000806000613b8e8460c0015160000151613d65565b6001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613bcb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613bef91906154fb565b905060005b856020015151811015613c6857816001600160a01b031686602001518281518110613c2157613c216152f9565b60200260200101516001600160a01b031603613c585785604001518181518110613c4d57613c4d6152f9565b602002602001015192505b613c6181615357565b9050613bf4565b5081613cab576040517fce480bcc0000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016111d3565b6000670de0b6b3a7640000833a8760e00151613cc78b8a610cc3565b613cd1919061538f565b613cdb91906153a7565b613ce591906153a7565b613cef91906155d6565b90508460c0015160200151811115613d5b5760208086015160c0870151909101516040517f394a2c2700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301526024820183905260448201526064016111d3565b9695505050505050565b6001600160a01b038181166000908152600360205260409020541680613dc2576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016111d3565b919050565b613e106040518060e00160405280600081526020016060815260200160006001600160a01b03168152602001606081526020016060815260200160608152602001600081525090565b6000613e248360a001518460c001516141c4565b805190915060008167ffffffffffffffff811115613e4457613e446144f1565b604051908082528060200260200182016040528015613e8957816020015b6040805180820190915260008082526020820152815260200190600190039081613e625790505b50905060008267ffffffffffffffff811115613ea757613ea76144f1565b604051908082528060200260200182016040528015613ed0578160200160208202803683370190505b50905060005b83811015614006576000613f06868381518110613ef557613ef56152f9565b602002602001015160000151613d65565b905080838381518110613f1b57613f1b6152f9565b60200260200101906001600160a01b031690816001600160a01b0316815250506040518060400160405280826001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613f84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613fa891906154fb565b6001600160a01b03168152602001878481518110613fc857613fc86152f9565b602002602001015160200151815250848381518110613fe957613fe96152f9565b60200260200101819052505080613fff90615357565b9050613ed6565b506040518060e0016040528087600001518152602001876040015160405160200161404091906001600160a01b0391909116815260200190565b604051602081830303815290604052815260200187606001516001600160a01b03168152602001876080015181526020018281526020018381526020018760e00151815250945050505050919050565b6040517f07a22a0700000000000000000000000000000000000000000000000000000000815260009030906307a22a07906140cf908590600401615637565b600060405180830381600087803b1580156140e957600080fd5b505af19250505080156140fa575060015b6141bc573d808015614128576040519150601f19603f3d011682016040523d82523d6000602084013e61412d565b606091505b5061413781615d7d565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da800000000000000000000000000000000000000000000000000000000036141885750600392915050565b806040517fcf19edfd0000000000000000000000000000000000000000000000000000000081526004016111d391906149d9565b506002919050565b606060005b83518110156143635782600001516001600160a01b03168482815181106141f2576141f26152f9565b6020026020010151600001516001600160a01b031603614353576000845167ffffffffffffffff811115614228576142286144f1565b60405190808252806020026020018201604052801561426d57816020015b60408051808201909152600080825260208201528152602001906001900390816142465790505b50905060005b85518110156142c45785818151811061428e5761428e6152f9565b60200260200101518282815181106142a8576142a86152f9565b6020026020010181905250806142bd90615357565b9050614273565b5060405180604001604052808284815181106142e2576142e26152f9565b6020026020010151600001516001600160a01b031681526020018560200151838581518110614313576143136152f9565b602002602001015160200151614329919061538f565b81525081838151811061433e5761433e6152f9565b60200260200101819052508092505050610d6e565b61435c81615357565b90506141c9565b50600083516001614374919061538f565b67ffffffffffffffff81111561438c5761438c6144f1565b6040519080825280602002602001820160405280156143d157816020015b60408051808201909152600080825260208201528152602001906001900390816143aa5790505b50905060005b8451811015614428578481815181106143f2576143f26152f9565b602002602001015182828151811061440c5761440c6152f9565b60200260200101819052508061442190615357565b90506143d7565b50828185518151811061443d5761443d6152f9565b60209081029190910101529392505050565b8280548282559060005260206000209081019282156144b1579160200282015b828111156144b1578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390911617825560209092019160019091019061446f565b50610e359291506144dc565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610e3557600081556001016144dd565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715614543576145436144f1565b60405290565b60405160e0810167ffffffffffffffff81118282101715614543576145436144f1565b604051610100810167ffffffffffffffff81118282101715614543576145436144f1565b604051610120810167ffffffffffffffff81118282101715614543576145436144f1565b604051601f8201601f1916810167ffffffffffffffff811182821017156145dd576145dd6144f1565b604052919050565b600067ffffffffffffffff8211156145ff576145ff6144f1565b50601f01601f191660200190565b600082601f83011261461e57600080fd5b813561463161462c826145e5565b6145b4565b81815284602083860101111561464657600080fd5b816020850160208301376000918101602001919091529392505050565b6001600160a01b038116811461081657600080fd5b8035613dc281614663565b600067ffffffffffffffff82111561469d5761469d6144f1565b5060051b60200190565b600082601f8301126146b857600080fd5b813560206146c861462c83614683565b82815260059290921b840181019181810190868411156146e757600080fd5b8286015b8481101561470b5780356146fe81614663565b83529183019183016146eb565b509695505050505050565b60006040828403121561472857600080fd5b614730614520565b9050813561473d81614663565b808252506020820135602082015292915050565b600082601f83011261476257600080fd5b8135602061477261462c83614683565b82815260069290921b8401810191818101908684111561479157600080fd5b8286015b8481101561470b576147a78882614716565b835291830191604001614795565b6000602082840312156147c757600080fd5b813567ffffffffffffffff808211156147df57600080fd5b9083019060e082860312156147f357600080fd5b6147fb614549565b8235815260208301358281111561481157600080fd5b61481d8782860161460d565b60208301525061482f60408401614678565b604082015260608301358281111561484657600080fd5b6148528782860161460d565b60608301525060808301358281111561486a57600080fd5b614876878286016146a7565b60808301525060a08301358281111561488e57600080fd5b61489a87828601614751565b60a08301525060c083013560c082015280935050505092915050565b6000602082840312156148c857600080fd5b8135611bd081614663565b67ffffffffffffffff8116811461081657600080fd5b8035613dc2816148d3565b60006020828403121561490657600080fd5b8135611bd0816148d3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b602081016004831061497b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b8381101561499c578181015183820152602001614984565b83811115612b765750506000910152565b600081518084526149c5816020860160208601614981565b601f01601f19169290920160200192915050565b602081526000611bd060208301846149ad565b600082601f8301126149fd57600080fd5b81356020614a0d61462c83614683565b82815260059290921b84018101918181019086841115614a2c57600080fd5b8286015b8481101561470b5780358352918301918301614a30565b60008060408385031215614a5a57600080fd5b823567ffffffffffffffff80821115614a7257600080fd5b614a7e868387016146a7565b93506020850135915080821115614a9457600080fd5b50614aa1858286016149ec565b9150509250929050565b600060208284031215614abd57600080fd5b813567ffffffffffffffff811115614ad457600080fd5b614ae0848285016146a7565b949350505050565b6020808252825182820181905260009190848201906040850190845b81811015614b2057835183529284019291840191600101614b04565b50909695505050505050565b60008060408385031215614b3f57600080fd5b82359150602083013567ffffffffffffffff80821115614b5e57600080fd5b908401906101208287031215614b7357600080fd5b614b7b61456c565b82358152614b8b602084016148e9565b6020820152614b9c60408401614678565b6040820152614bad60608401614678565b6060820152608083013582811115614bc457600080fd5b614bd08882860161460d565b60808301525060a083013582811115614be857600080fd5b614bf488828601614751565b60a083015250614c078760c08501614716565b60c082015261010083013560e08201528093505050509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614b205783516001600160a01b031683529284019291840191600101614c40565b60008060408385031215614c7857600080fd5b8235614c8381614663565b91506020830135614c9381614663565b809150509250929050565b600081518084526020808501945080840160005b83811015614cd75781516001600160a01b031687529582019590820190600101614cb2565b509495945050505050565b602081526000611bd06020830184614c9e565b600060408284031215614d0757600080fd5b614d0f614520565b9050813581526020820135602082015292915050565b600060408284031215614d3757600080fd5b611bd08383614cf5565b600060208284031215614d5357600080fd5b813567ffffffffffffffff811115614d6a57600080fd5b820160e08185031215611bd057600080fd5b600060208284031215614d8e57600080fd5b5035919050565b600082601f830112614da657600080fd5b81356020614db661462c83614683565b82815260059290921b84018101918181019086841115614dd557600080fd5b8286015b8481101561470b578035614dec816148d3565b8352918301918301614dd9565b600082601f830112614e0a57600080fd5b81356020614e1a61462c83614683565b82815260069290921b84018101918181019086841115614e3957600080fd5b8286015b8481101561470b57614e4f8882614cf5565b835291830191604001614e3d565b600082601f830112614e6e57600080fd5b81356020614e7e61462c83614683565b82815260059290921b84018101918181019086841115614e9d57600080fd5b8286015b8481101561470b57803567ffffffffffffffff811115614ec15760008081fd5b614ecf8986838b010161460d565b845250918301918301614ea1565b600060208284031215614eef57600080fd5b813567ffffffffffffffff80821115614f0757600080fd5b908301906101208286031215614f1c57600080fd5b614f24614590565b823582811115614f3357600080fd5b614f3f87828601614d95565b825250602083013582811115614f5457600080fd5b614f60878286016146a7565b602083015250604083013582811115614f7857600080fd5b614f84878286016149ec565b604083015250606083013582811115614f9c57600080fd5b614fa887828601614df9565b606083015250608083013582811115614fc057600080fd5b614fcc87828601614e5d565b60808301525060a083013582811115614fe457600080fd5b614ff0878286016149ec565b60a08301525060c083013560c082015260e08301358281111561501257600080fd5b61501e878286016149ec565b60e0830152506101009283013592810192909252509392505050565b60008083601f84011261504c57600080fd5b50813567ffffffffffffffff81111561506457600080fd5b6020830191508360208260051b850101111561507f57600080fd5b9250929050565b60008060008060008060008060e0898b0312156150a257600080fd5b606089018a8111156150b357600080fd5b8998503567ffffffffffffffff808211156150cd57600080fd5b818b0191508b601f8301126150e157600080fd5b8135818111156150f057600080fd5b8c602082850101111561510257600080fd5b6020830199508098505060808b013591508082111561512057600080fd5b61512c8c838d0161503a565b909750955060a08b013591508082111561514557600080fd5b506151528b828c0161503a565b999c989b50969995989497949560c00135949350505050565b60006080828403121561517d57600080fd5b6040516080810181811067ffffffffffffffff821117156151a0576151a06144f1565b604052823563ffffffff811681146151b757600080fd5b815260208301356151c7816148d3565b602082015260408301356151da816148d3565b604082015260608301356151ed816148d3565b60608201529392505050565b803560ff81168114613dc257600080fd5b60008060008060008060c0878903121561522357600080fd5b863567ffffffffffffffff8082111561523b57600080fd5b6152478a838b016146a7565b9750602089013591508082111561525d57600080fd5b6152698a838b016146a7565b965061527760408a016151f9565b9550606089013591508082111561528d57600080fd5b6152998a838b0161460d565b94506152a760808a016148e9565b935060a08901359150808211156152bd57600080fd5b506152ca89828a0161460d565b9150509295509295509295565b6000602082840312156152e957600080fd5b81518015158114611bd057600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361538857615388615328565b5060010190565b600082198211156153a2576153a2615328565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156153df576153df615328565b500290565b6000828210156153f6576153f6615328565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff0382111561544757615447615328565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff8316806154915761549161544f565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60006020828403121561550d57600080fd5b8151611bd081614663565b600063ffffffff80831681851680830382111561553757615537615328565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526155708184018a614c9e565b905082810360808401526155848189614c9e565b905060ff871660a084015282810360c08401526155a181876149ad565b905067ffffffffffffffff851660e08401528281036101008401526155c681856149ad565b9c9b505050505050505050505050565b6000826155e5576155e561544f565b500490565b600081518084526020808501945080840160005b83811015614cd75761562487835180516001600160a01b03168252602090810151910152565b60409690960195908201906001016155fe565b60208152815160208201526000602083015160e0604084015261565e6101008401826149ad565b90506001600160a01b0360408501511660608401526060840151601f198085840301608086015261568f83836149ad565b925060808601519150808584030160a08601526156ac8383614c9e565b925060a08601519150808584030160c0860152506156ca82826155ea565b91505060c084015160e08401528091505092915050565b8051613dc2816148d3565b8051613dc281614663565b600082601f83011261570857600080fd5b815161571661462c826145e5565b81815284602083860101111561572b57600080fd5b614ae0826020830160208701614981565b60006040828403121561574e57600080fd5b615756614520565b9050815161576381614663565b808252506020820151602082015292915050565b600082601f83011261578857600080fd5b8151602061579861462c83614683565b82815260069290921b840181019181810190868411156157b757600080fd5b8286015b8481101561470b576157cd888261573c565b8352918301916040016157bb565b6000602082840312156157ed57600080fd5b815167ffffffffffffffff8082111561580557600080fd5b90830190610120828603121561581a57600080fd5b61582261456c565b82518152615832602084016156e1565b6020820152615843604084016156ec565b6040820152615854606084016156ec565b606082015260808301518281111561586b57600080fd5b615877878286016156f7565b60808301525060a08301518281111561588f57600080fd5b61589b87828601615777565b60a0830152506158ae8660c0850161573c565b60c0820152610100929092015160e0830152509392505050565b600082601f8301126158d957600080fd5b815160206158e961462c83614683565b82815260059290921b8401810191818101908684111561590857600080fd5b8286015b8481101561470b57805161591f816148d3565b835291830191830161590c565b600082601f83011261593d57600080fd5b8151602061594d61462c83614683565b82815260059290921b8401810191818101908684111561596c57600080fd5b8286015b8481101561470b57805161598381614663565b8352918301918301615970565b600082601f8301126159a157600080fd5b815160206159b161462c83614683565b82815260059290921b840181019181810190868411156159d057600080fd5b8286015b8481101561470b57805183529183019183016159d4565b600082601f8301126159fc57600080fd5b81516020615a0c61462c83614683565b82815260069290921b84018101918181019086841115615a2b57600080fd5b8286015b8481101561470b5760408189031215615a485760008081fd5b615a50614520565b815181528482015185820152835291830191604001615a2f565b600082601f830112615a7b57600080fd5b81516020615a8b61462c83614683565b82815260059290921b84018101918181019086841115615aaa57600080fd5b8286015b8481101561470b57805167ffffffffffffffff811115615ace5760008081fd5b615adc8986838b01016156f7565b845250918301918301615aae565b600060208284031215615afc57600080fd5b815167ffffffffffffffff80821115615b1457600080fd5b908301906101208286031215615b2957600080fd5b615b31614590565b825182811115615b4057600080fd5b615b4c878286016158c8565b825250602083015182811115615b6157600080fd5b615b6d8782860161592c565b602083015250604083015182811115615b8557600080fd5b615b9187828601615990565b604083015250606083015182811115615ba957600080fd5b615bb5878286016159eb565b606083015250608083015182811115615bcd57600080fd5b615bd987828601615a6a565b60808301525060a083015182811115615bf157600080fd5b615bfd87828601615990565b60a08301525060c083015160c082015260e083015182811115615c1f57600080fd5b615c2b87828601615990565b60e0830152506101009283015192810192909252509392505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b166040850152816060850152615c818285018b614c9e565b91508382036080850152615c95828a614c9e565b915060ff881660a085015283820360c0850152615cb282886149ad565b90861660e085015283810361010085015290506155c681856149ad565b602081526000611bd060208301846155ea565b600081518084526020808501945080840160005b83811015614cd757815187529582019590820190600101615cf6565b60a081526000615d2560a0830188615ce2565b8281036020840152615d378188615ce2565b90508560408401528281036060840152615d518186615ce2565b9150508260808301529695505050505050565b600060208284031215615d7657600080fd5b5051919050565b6000815160208301517fffffffff0000000000000000000000000000000000000000000000000000000080821693506004831015615dc55780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
}

var EVM2EVMTollOffRampABI = EVM2EVMTollOffRampMetaData.ABI

var EVM2EVMTollOffRampBin = EVM2EVMTollOffRampMetaData.Bin

func DeployEVM2EVMTollOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, onRampAddress common.Address, commitStore common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address) (common.Address, *types.Transaction, *EVM2EVMTollOffRamp, error) {
	parsed, err := EVM2EVMTollOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMTollOffRampBin), backend, sourceChainId, chainId, offRampConfig, onRampAddress, commitStore, afn, sourceTokens, pools, rateLimiterConfig, tokenLimitsAdmin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMTollOffRamp{EVM2EVMTollOffRampCaller: EVM2EVMTollOffRampCaller{contract: contract}, EVM2EVMTollOffRampTransactor: EVM2EVMTollOffRampTransactor{contract: contract}, EVM2EVMTollOffRampFilterer: EVM2EVMTollOffRampFilterer{contract: contract}}, nil
}

type EVM2EVMTollOffRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMTollOffRampCaller
	EVM2EVMTollOffRampTransactor
	EVM2EVMTollOffRampFilterer
}

type EVM2EVMTollOffRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampSession struct {
	Contract     *EVM2EVMTollOffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOffRampCallerSession struct {
	Contract *EVM2EVMTollOffRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMTollOffRampTransactorSession struct {
	Contract     *EVM2EVMTollOffRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOffRampRaw struct {
	Contract *EVM2EVMTollOffRamp
}

type EVM2EVMTollOffRampCallerRaw struct {
	Contract *EVM2EVMTollOffRampCaller
}

type EVM2EVMTollOffRampTransactorRaw struct {
	Contract *EVM2EVMTollOffRampTransactor
}

func NewEVM2EVMTollOffRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMTollOffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMTollOffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMTollOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRamp{address: address, abi: abi, EVM2EVMTollOffRampCaller: EVM2EVMTollOffRampCaller{contract: contract}, EVM2EVMTollOffRampTransactor: EVM2EVMTollOffRampTransactor{contract: contract}, EVM2EVMTollOffRampFilterer: EVM2EVMTollOffRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMTollOffRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMTollOffRampCaller, error) {
	contract, err := bindEVM2EVMTollOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampCaller{contract: contract}, nil
}

func NewEVM2EVMTollOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMTollOffRampTransactor, error) {
	contract, err := bindEVM2EVMTollOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampTransactor{contract: contract}, nil
}

func NewEVM2EVMTollOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMTollOffRampFilterer, error) {
	contract, err := bindEVM2EVMTollOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampFilterer{contract: contract}, nil
}

func bindEVM2EVMTollOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMTollOffRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOffRamp.Contract.EVM2EVMTollOffRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.EVM2EVMTollOffRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.EVM2EVMTollOffRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(AggregateRateLimiterInterfaceTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(AggregateRateLimiterInterfaceTokenBucket)).(*AggregateRateLimiterInterfaceTokenBucket)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMTollOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMTollOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMTollOffRamp.Contract.CcipReceive(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMTollOffRamp.Contract.CcipReceive(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) FeeTaken(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "feeTaken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) FeeTaken(arg0 *big.Int) (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.FeeTaken(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) FeeTaken(arg0 *big.Int) (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.FeeTaken(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetAFN(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetAFN(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getChainIDs")

	outstruct := new(GetChainIDs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMTollOffRamp.Contract.GetChainIDs(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMTollOffRamp.Contract.GetChainIDs(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetCommitStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getCommitStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetCommitStore() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetCommitStore(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetCommitStore() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetCommitStore(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOffRampInterfaceOffRampConfig)).(*BaseOffRampInterfaceOffRampConfig)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMTollOffRamp.Contract.GetConfig(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMTollOffRamp.Contract.GetConfig(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetDestinationToken(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetDestinationToken(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetDestinationTokens(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetDestinationTokens(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMTollOffRamp.Contract.GetExecutionState(&_EVM2EVMTollOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMTollOffRamp.Contract.GetExecutionState(&_EVM2EVMTollOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPool(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPool(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPoolTokens(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPoolTokens(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPricesForTokens(&_EVM2EVMTollOffRamp.CallOpts, tokens)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPricesForTokens(&_EVM2EVMTollOffRamp.CallOpts, tokens)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetRouter(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetRouter(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.IsAFNHealthy(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.IsAFNHealthy(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMTollOffRamp.Contract.LatestConfigDetails(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMTollOffRamp.Contract.LatestConfigDetails(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMTollOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMTollOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) OverheadGasToll(opts *bind.CallOpts, merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "overheadGasToll", merkleGasShare, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) OverheadGasToll(merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.OverheadGasToll(&_EVM2EVMTollOffRamp.CallOpts, merkleGasShare, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) OverheadGasToll(merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.OverheadGasToll(&_EVM2EVMTollOffRamp.CallOpts, merkleGasShare, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.Owner(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.Owner(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Paused() (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.Paused(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMTollOffRamp.Contract.Paused(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.Transmitters(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.Transmitters(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOffRamp.Contract.TypeAndVersion(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOffRamp.Contract.TypeAndVersion(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.AcceptOwnership(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.AcceptOwnership(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.AddPool(&_EVM2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.AddPool(&_EVM2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "executeSingleMessage", message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRamp.TransactOpts, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRamp.TransactOpts, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) ManuallyExecute(opts *bind.TransactOpts, report CCIPExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "manuallyExecute", report)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) ManuallyExecute(report CCIPExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.ManuallyExecute(&_EVM2EVMTollOffRamp.TransactOpts, report)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) ManuallyExecute(report CCIPExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.ManuallyExecute(&_EVM2EVMTollOffRamp.TransactOpts, report)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Pause(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Pause(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.RemovePool(&_EVM2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.RemovePool(&_EVM2EVMTollOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetAFN(&_EVM2EVMTollOffRamp.TransactOpts, afn)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetAFN(&_EVM2EVMTollOffRamp.TransactOpts, afn)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetCommitStore(opts *bind.TransactOpts, commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setCommitStore", commitStore)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetCommitStore(commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetCommitStore(&_EVM2EVMTollOffRamp.TransactOpts, commitStore)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetCommitStore(commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetCommitStore(&_EVM2EVMTollOffRamp.TransactOpts, commitStore)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetConfig(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetConfig0(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setConfig0", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig0(&_EVM2EVMTollOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig0(&_EVM2EVMTollOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetPrices(&_EVM2EVMTollOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetPrices(&_EVM2EVMTollOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetRouter(&_EVM2EVMTollOffRamp.TransactOpts, router)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetRouter(&_EVM2EVMTollOffRamp.TransactOpts, router)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMTollOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMTollOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.TransferOwnership(&_EVM2EVMTollOffRamp.TransactOpts, to)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.TransferOwnership(&_EVM2EVMTollOffRamp.TransactOpts, to)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Transmit(&_EVM2EVMTollOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Transmit(&_EVM2EVMTollOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Unpause(&_EVM2EVMTollOffRamp.TransactOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.Unpause(&_EVM2EVMTollOffRamp.TransactOpts)
}

type EVM2EVMTollOffRampAFNSetIterator struct {
	Event *EVM2EVMTollOffRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampAFNSet)
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
		it.Event = new(EVM2EVMTollOffRampAFNSet)
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

func (it *EVM2EVMTollOffRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampAFNSetIterator{contract: _EVM2EVMTollOffRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampAFNSet)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseAFNSet(log types.Log) (*EVM2EVMTollOffRampAFNSet, error) {
	event := new(EVM2EVMTollOffRampAFNSet)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampConfigChangedIterator struct {
	Event *EVM2EVMTollOffRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampConfigChanged)
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
		it.Event = new(EVM2EVMTollOffRampConfigChanged)
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

func (it *EVM2EVMTollOffRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampConfigChangedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampConfigChanged)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMTollOffRampConfigChanged, error) {
	event := new(EVM2EVMTollOffRampConfigChanged)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampConfigSetIterator struct {
	Event *EVM2EVMTollOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampConfigSet)
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
		it.Event = new(EVM2EVMTollOffRampConfigSet)
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

func (it *EVM2EVMTollOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampConfigSet struct {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampConfigSetIterator{contract: _EVM2EVMTollOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampConfigSet)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseConfigSet(log types.Log) (*EVM2EVMTollOffRampConfigSet, error) {
	event := new(EVM2EVMTollOffRampConfigSet)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampExecutionStateChangedIterator struct {
	Event *EVM2EVMTollOffRampExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampExecutionStateChanged)
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
		it.Event = new(EVM2EVMTollOffRampExecutionStateChanged)
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

func (it *EVM2EVMTollOffRampExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampExecutionStateChanged struct {
	SequenceNumber uint64
	State          uint8
	Raw            types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMTollOffRampExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampExecutionStateChangedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampExecutionStateChanged)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseExecutionStateChanged(log types.Log) (*EVM2EVMTollOffRampExecutionStateChanged, error) {
	event := new(EVM2EVMTollOffRampExecutionStateChanged)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampOffRampConfigSetIterator struct {
	Event *EVM2EVMTollOffRampOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampOffRampConfigSet)
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
		it.Event = new(EVM2EVMTollOffRampOffRampConfigSet)
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

func (it *EVM2EVMTollOffRampOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampOffRampConfigSet struct {
	Config BaseOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampOffRampConfigSetIterator{contract: _EVM2EVMTollOffRamp.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampOffRampConfigSet)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseOffRampConfigSet(log types.Log) (*EVM2EVMTollOffRampOffRampConfigSet, error) {
	event := new(EVM2EVMTollOffRampOffRampConfigSet)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampOffRampRouterSetIterator struct {
	Event *EVM2EVMTollOffRampOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampOffRampRouterSet)
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
		it.Event = new(EVM2EVMTollOffRampOffRampRouterSet)
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

func (it *EVM2EVMTollOffRampOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampOffRampRouterSet struct {
	Router        common.Address
	SourceChainId *big.Int
	OnRampAddress common.Address
	Raw           types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMTollOffRampOffRampRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampOffRampRouterSetIterator{contract: _EVM2EVMTollOffRamp.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampOffRampRouterSet)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseOffRampRouterSet(log types.Log) (*EVM2EVMTollOffRampOffRampRouterSet, error) {
	event := new(EVM2EVMTollOffRampOffRampRouterSet)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMTollOffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMTollOffRampOwnershipTransferRequested)
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

func (it *EVM2EVMTollOffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampOwnershipTransferRequestedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampOwnershipTransferRequested)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOffRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMTollOffRampOwnershipTransferRequested)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampOwnershipTransferredIterator struct {
	Event *EVM2EVMTollOffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMTollOffRampOwnershipTransferred)
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

func (it *EVM2EVMTollOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampOwnershipTransferredIterator{contract: _EVM2EVMTollOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampOwnershipTransferred)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOffRampOwnershipTransferred, error) {
	event := new(EVM2EVMTollOffRampOwnershipTransferred)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampPausedIterator struct {
	Event *EVM2EVMTollOffRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampPaused)
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
		it.Event = new(EVM2EVMTollOffRampPaused)
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

func (it *EVM2EVMTollOffRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampPausedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampPaused)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParsePaused(log types.Log) (*EVM2EVMTollOffRampPaused, error) {
	event := new(EVM2EVMTollOffRampPaused)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampPoolAddedIterator struct {
	Event *EVM2EVMTollOffRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampPoolAdded)
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
		it.Event = new(EVM2EVMTollOffRampPoolAdded)
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

func (it *EVM2EVMTollOffRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampPoolAddedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampPoolAdded)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMTollOffRampPoolAdded, error) {
	event := new(EVM2EVMTollOffRampPoolAdded)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampPoolRemovedIterator struct {
	Event *EVM2EVMTollOffRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampPoolRemoved)
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
		it.Event = new(EVM2EVMTollOffRampPoolRemoved)
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

func (it *EVM2EVMTollOffRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampPoolRemovedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampPoolRemoved)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMTollOffRampPoolRemoved, error) {
	event := new(EVM2EVMTollOffRampPoolRemoved)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampTokenPriceChangedIterator struct {
	Event *EVM2EVMTollOffRampTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampTokenPriceChanged)
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
		it.Event = new(EVM2EVMTollOffRampTokenPriceChanged)
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

func (it *EVM2EVMTollOffRampTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampTokenPriceChangedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampTokenPriceChanged)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMTollOffRampTokenPriceChanged, error) {
	event := new(EVM2EVMTollOffRampTokenPriceChanged)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMTollOffRampTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMTollOffRampTokensRemovedFromBucket)
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

func (it *EVM2EVMTollOffRampTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampTokensRemovedFromBucketIterator{contract: _EVM2EVMTollOffRamp.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampTokensRemovedFromBucket)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMTollOffRampTokensRemovedFromBucket, error) {
	event := new(EVM2EVMTollOffRampTokensRemovedFromBucket)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampTransmittedIterator struct {
	Event *EVM2EVMTollOffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampTransmitted)
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
		it.Event = new(EVM2EVMTollOffRampTransmitted)
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

func (it *EVM2EVMTollOffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTransmittedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampTransmittedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTransmitted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampTransmitted)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseTransmitted(log types.Log) (*EVM2EVMTollOffRampTransmitted, error) {
	event := new(EVM2EVMTollOffRampTransmitted)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampUnpausedIterator struct {
	Event *EVM2EVMTollOffRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampUnpaused)
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
		it.Event = new(EVM2EVMTollOffRampUnpaused)
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

func (it *EVM2EVMTollOffRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampUnpausedIterator{contract: _EVM2EVMTollOffRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampUnpaused)
				if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMTollOffRampUnpaused, error) {
	event := new(EVM2EVMTollOffRampUnpaused)
	if err := _EVM2EVMTollOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMTollOffRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMTollOffRamp.ParseAFNSet(log)
	case _EVM2EVMTollOffRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMTollOffRamp.ParseConfigChanged(log)
	case _EVM2EVMTollOffRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMTollOffRamp.ParseConfigSet(log)
	case _EVM2EVMTollOffRamp.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMTollOffRamp.ParseExecutionStateChanged(log)
	case _EVM2EVMTollOffRamp.abi.Events["OffRampConfigSet"].ID:
		return _EVM2EVMTollOffRamp.ParseOffRampConfigSet(log)
	case _EVM2EVMTollOffRamp.abi.Events["OffRampRouterSet"].ID:
		return _EVM2EVMTollOffRamp.ParseOffRampRouterSet(log)
	case _EVM2EVMTollOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMTollOffRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMTollOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMTollOffRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMTollOffRamp.abi.Events["Paused"].ID:
		return _EVM2EVMTollOffRamp.ParsePaused(log)
	case _EVM2EVMTollOffRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMTollOffRamp.ParsePoolAdded(log)
	case _EVM2EVMTollOffRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMTollOffRamp.ParsePoolRemoved(log)
	case _EVM2EVMTollOffRamp.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMTollOffRamp.ParseTokenPriceChanged(log)
	case _EVM2EVMTollOffRamp.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMTollOffRamp.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMTollOffRamp.abi.Events["Transmitted"].ID:
		return _EVM2EVMTollOffRamp.ParseTransmitted(log)
	case _EVM2EVMTollOffRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMTollOffRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMTollOffRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMTollOffRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMTollOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMTollOffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf")
}

func (EVM2EVMTollOffRampOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1a9ff9caccd597226ff8c393ab44e6b57669905d9b16f1cdb2ac267253dbf27c")
}

func (EVM2EVMTollOffRampOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x1090eec34bfa4c1a6201bdd67d36d1e4527d87f93cd925fe8b6065a953e1e636")
}

func (EVM2EVMTollOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMTollOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMTollOffRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMTollOffRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMTollOffRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMTollOffRampTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMTollOffRampTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMTollOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (EVM2EVMTollOffRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRamp) Address() common.Address {
	return _EVM2EVMTollOffRamp.address
}

type EVM2EVMTollOffRampInterface interface {
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error)

	CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error

	FeeTaken(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

		error)

	GetCommitStore(opts *bind.CallOpts) (common.Address, error)

	GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

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

	OverheadGasToll(opts *bind.CallOpts, merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error)

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

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMTollOffRampAFNSet, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMTollOffRampConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMTollOffRampConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMTollOffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMTollOffRampExecutionStateChanged, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*EVM2EVMTollOffRampOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMTollOffRampOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*EVM2EVMTollOffRampOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOffRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMTollOffRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMTollOffRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOffRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMTollOffRampPoolRemoved, error)

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMTollOffRampTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMTollOffRampTokensRemovedFromBucket, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMTollOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMTollOffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMTollOffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
