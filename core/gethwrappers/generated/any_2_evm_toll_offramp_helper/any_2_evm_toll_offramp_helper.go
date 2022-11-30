// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package any_2_evm_toll_offramp_helper

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

var EVM2EVMTollOffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"linkPerUnitGas\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"rep\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeTaken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitStore\",\"outputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"linkPerUnitGas\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"merkleGasShare\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount\",\"name\":\"feeTokenAndAmount\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVM2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"overheadGasToll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"}],\"name\":\"setCommitStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200685538038062006855833981016040819052620000359162000768565b8989898989898989898960018a8a89898989898989818185858833806000806000806101000a81548160ff02191690831515021790555060006001600160a01b0316826001600160a01b031603620000d45760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b0319909216919091179091558116156200010e576200010e81620003b0565b5050506001600160a01b0381166200013957604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200017b5760405162d8548360e71b815260040160405180910390fd5b81516200019090600490602085019062000461565b5060005b82518110156200025b576000828281518110620001b557620001b56200085d565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001ff57620001ff6200085d565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b029290911691909117905550620002538162000873565b905062000194565b5050600580546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600855600993909355600a55600b919091558716620002d8576040516342bcdf7f60e11b815260040160405180910390fd5b5050506080959095525060a0929092526001600160a01b0390811660c052600d80546001600160a01b0319169190921617905550151560e0525050855160168054602089015160408a01516060909a015163ffffffff9094166001600160601b0319909216919091176401000000006001600160401b039283160217600160601b600160e01b0319166c0100000000000000000000000099821699909902600160a01b600160e01b03191698909817600160a01b989092169790970217909555506200089b9f50505050505050505050505050505050565b336001600160a01b038216036200040a5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000cb565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620004b9579160200282015b82811115620004b957825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000482565b50620004c7929150620004cb565b5090565b5b80821115620004c75760008155600101620004cc565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620005235762000523620004e2565b604052919050565b80516001600160401b03811681146200054357600080fd5b919050565b6000608082840312156200055b57600080fd5b604051608081016001600160401b0381118282101715620005805762000580620004e2565b8060405250809150825163ffffffff811681146200059d57600080fd5b8152620005ad602084016200052b565b6020820152620005c0604084016200052b565b6040820152620005d3606084016200052b565b60608201525092915050565b6001600160a01b0381168114620005f557600080fd5b50565b80516200054381620005df565b60006001600160401b03821115620006215762000621620004e2565b5060051b60200190565b600082601f8301126200063d57600080fd5b8151602062000656620006508362000605565b620004f8565b82815260059290921b840181019181810190868411156200067657600080fd5b8286015b848110156200069e5780516200069081620005df565b83529183019183016200067a565b509695505050505050565b600082601f830112620006bb57600080fd5b81516020620006ce620006508362000605565b82815260059290921b84018101918181019086841115620006ee57600080fd5b8286015b848110156200069e5780516200070881620005df565b8352918301918301620006f2565b6000604082840312156200072957600080fd5b604080519081016001600160401b03811182821017156200074e576200074e620004e2565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000806101c08b8d0312156200078957600080fd5b8a51995060208b01519850620007a38c60408d0162000548565b9750620007b360c08c01620005f8565b9650620007c360e08c01620005f8565b9550620007d46101008c01620005f8565b6101208c01519095506001600160401b0380821115620007f357600080fd5b620008018e838f016200062b565b95506101408d01519150808211156200081957600080fd5b50620008288d828e01620006a9565b9350506200083b8c6101608d0162000716565b91506200084c6101a08c01620005f8565b90509295989b9194979a5092959850565b634e487b7160e01b600052603260045260246000fd5b6000600182016200089457634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e051615f56620008ff6000396000611696015260008181611e55015261311e01526000818161032c01528181611e2f01526130f801526000818161030901528181611e0a015281816130d30152613b000152615f566000f3fe608060405234801561001057600080fd5b50600436106102e95760003560e01c806389c0656811610191578063b5767166116100e3578063c3f909d411610097578063e3d0e71211610071578063e3d0e712146107bb578063eb511dd4146107ce578063f2fde38b146107e157600080fd5b8063c3f909d4146106c3578063c5a1d7f0146107a0578063c9033284146107a857600080fd5b8063bbe4f6db116100c8578063bbe4f6db14610671578063bc29705a1461069d578063c0d78655146106b057600080fd5b8063b57671661461064d578063b66f0efb1461066057600080fd5b8063a8b640c111610145578063b0f479a11161011f578063b0f479a114610616578063b1dc65a414610627578063b4069b311461063a57600080fd5b8063a8b640c1146105c3578063a8e91321146105e3578063afcb95d7146105f657600080fd5b806390c2339b1161017657806390c2339b1461056757806391872543146105a25780639438ff63146105b557600080fd5b806389c06568146105495780638da5cb5b1461055157600080fd5b80634741062e1161024a578063681fba16116101fe57806381411834116101d857806381411834146104fc57806381ff7048146105115780638456cb591461054157600080fd5b8063681fba16146104cc578063744b92e2146104e157806379ba5097146104f457600080fd5b80635b4dc8121161022f5780635b4dc8121461048d5780635c975abb146104ae578063609ec69c146104b957600080fd5b80634741062e1461045c578063599f64311461047c57600080fd5b8063181f5a77116102a157806339aa92641161028657806339aa92641461042e5780633f4ba83a146104415780634352fa9f1461044957600080fd5b8063181f5a77146103c05780632222dd421461040957600080fd5b8063108ee5fc116102d2578063108ee5fc1461035b578063142a98fc1461036e578063147809b3146103a857600080fd5b806307a22a07146102ee578063087ae6df14610303575b600080fd5b6103016102fc366004614895565b6107f4565b005b604080517f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006020820152015b60405180910390f35b610301610369366004614996565b610868565b61039b61037c3660046149d4565b67ffffffffffffffff166000908152600e602052604090205460ff1690565b6040516103529190614a20565b6103b061091f565b6040519015158152602001610352565b6103fc6040518060400160405280601881526020017f45564d3245564d546f6c6c4f666652616d7020312e302e30000000000000000081525081565b6040516103529190614ab9565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610352565b61030161043c366004614996565b6109ac565b6103016109e3565b610301610457366004614b27565b6109f5565b61046f61046a366004614b8b565b610c4a565b6040516103529190614bc8565b6005546001600160a01b0316610416565b6104a061049b366004614c0c565b610d12565b604051908152602001610352565b60005460ff166103b0565b6103016104c7366004614fc9565b610dc3565b6104d4610dd1565b604051610352919061501b565b6103016104ef36600461505c565b610e96565b6103016111da565b6105046112c2565b60405161035291906150ce565b601154600f546040805163ffffffff80851682526401000000009094049093166020840152820152606001610352565b610301611324565b6104d4611334565b60005461010090046001600160a01b0316610416565b61056f611394565b60405161035291908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6103016105b03660046150e1565b611435565b6103016102e93660046150fd565b6104a06105d1366004615138565b60156020526000908152604090205481565b6103016105f1366004615151565b611568565b604080516001815260006020820181905291810191909152606001610352565b600c546001600160a01b0316610416565b6103016106353660046151d2565b611573565b610416610648366004614996565b611b46565b61030161065b3660046152b7565b611c34565b600d546001600160a01b0316610416565b61041661067f366004614996565b6001600160a01b039081166000908152600360205260409020541690565b6103016106ab3660046152ec565b611c40565b6103016106be366004614996565b611d76565b610755604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260165463ffffffff8116825267ffffffffffffffff6401000000008204811660208401526c010000000000000000000000008204811693830193909352740100000000000000000000000000000000000000009004909116606082015290565b6040516103529190815163ffffffff16815260208083015167ffffffffffffffff90811691830191909152604080840151821690830152606092830151169181019190915260800190565b6104a0611dd5565b6103016107b6366004614996565b611ea5565b6103016107c936600461538b565b611edc565b6103016107dc36600461505c565b6127d1565b6103016107ef366004614996565b6129a9565b33301461082d576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a0810151511561085c576108458160a001516129ba565b61085c81608001518260a001518360400151612bbe565b61086581612c5c565b50565b610870612d34565b6001600160a01b0381166108b0576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa158015610982573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a69190615458565b15905090565b6109b4612d34565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6109eb612d34565b6109f3612d93565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015610a3257506005546001600160a01b03163314155b15610a69576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610aa5576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610aff576006600060078381548110610aca57610aca615475565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610af8816154d3565b9050610aab565b5060005b82811015610c2f576000858281518110610b1f57610b1f615475565b6020026020010151905060006001600160a01b0316816001600160a01b031603610b75576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610b8757610b87615475565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610bec57610bec615475565b6020026020010151604051610c169291906001600160a01b03929092168252602082015260400190565b60405180910390a150610c28816154d3565b9050610b03565b508351610c4390600790602087019061452f565b5050505050565b80516060908067ffffffffffffffff811115610c6857610c686145d1565b604051908082528060200260200182016040528015610c91578160200160208202803683370190505b50915060005b81811015610d0b5760066000858381518110610cb557610cb5615475565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610cf057610cf0615475565b6020908102919091010152610d04816154d3565b9050610c97565b5050919050565b6000808260800151518360a001515160206014610d2f919061550b565b610d399190615523565b610d4490608661550b565b610d4e919061550b565b90506000610d5d601083615523565b9050610a28611bbc8560a00151516001610d77919061550b565b610d8390618aac615523565b6156b8610d90898661550b565b610d9a919061550b565b610da4919061550b565b610dae919061550b565b610db8919061550b565b925050505b92915050565b610dcd8282612e2f565b5050565b60045460609067ffffffffffffffff811115610def57610def6145d1565b604051908082528060200260200182016040528015610e18578160200160208202803683370190505b50905060005b600454811015610e9257610e5860048281548110610e3e57610e3e615475565b6000918252602090912001546001600160a01b0316611b46565b828281518110610e6a57610e6a615475565b6001600160a01b0390921660209283029190910190910152610e8b816154d3565b9050610e1e565b5090565b610e9e612d34565b6004546000819003610edc576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610f6a576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614610fb9576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610fc8600185615560565b81548110610fd857610fd8615475565b9060005260206000200160009054906101000a90046001600160a01b03169050600482602001516bffffffffffffffffffffffff168154811061101d5761101d615475565b6000918252602090912001546001600160a01b0316600461103f600186615560565b8154811061104f5761104f615475565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483602001516bffffffffffffffffffffffff16815481106110a3576110a3615475565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000291909216179055600480548061112d5761112d615577565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9084018101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03878116808452600383526040808520949094558351908152908716918101919091527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146112395760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060601480548060200260200160405190810160405280929190818152602001828054801561131a57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116112fc575b5050505050905090565b61132c612d34565b6109f361355e565b6060600480548060200260200160405190810160405280929190818152602001828054801561131a576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116112fc575050505050905090565b6113bf6040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b546060820181905242906000906113f99083615560565b60208401518451919250611425916114119084615523565b8560400151611420919061550b565b6135e6565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561147257506005546001600160a01b03163314155b156114a9576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff116114fd576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61150760086135fc565b602081015160098190558151600855600a5461152391906135e6565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b610865816001612e2f565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916115c991849163ffffffff851691908e908e90819084018382808284376000920191909152506136a992505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600f5480825260105460ff808216602085015261010090910416928201929092529083146116845760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401611230565b6116928b8b8b8b8b8b6136cc565b60007f0000000000000000000000000000000000000000000000000000000000000000156116ef576002826020015183604001516116d091906155a6565b6116da91906155fa565b6116e59060016155a6565b60ff169050611705565b60208201516116ff9060016155a6565b60ff1690505b8881146117545760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401611230565b8887146117a35760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401611230565b3360009081526012602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156117e6576117e66149f1565b60028111156117f7576117f76149f1565b9052509050600281602001516002811115611814576118146149f1565b14801561184e57506014816000015160ff168154811061183657611836615475565b6000918252602090912001546001600160a01b031633145b61189a5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401611230565b5050505050600088886040516118b192919061561c565b6040519081900381206118c8918c9060200161562c565b6040516020818303038152906040528051906020012090506118e861459d565b604080518082019091526000808252602082015260005b88811015611b2457600060018588846020811061191e5761191e615475565b61192b91901a601b6155a6565b8d8d8681811061193d5761193d615475565b905060200201358c8c8781811061195657611956615475565b9050602002013560405160008152602001604052604051611993949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156119b5573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526012602090815290849020838501909452835460ff80821685529296509294508401916101009004166002811115611a0a57611a0a6149f1565b6002811115611a1b57611a1b6149f1565b9052509250600183602001516002811115611a3857611a386149f1565b14611a855760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401611230565b8251849060ff16601f8110611a9c57611a9c615475565b602002015115611aee5760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401611230565b600184846000015160ff16601f8110611b0957611b09615475565b9115156020909202015250611b1d816154d3565b90506118ff565b5050505063ffffffff8110611b3b57611b3b615648565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611b9a576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611c09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c2d9190615677565b9392505050565b610865600080836136a9565b611c48612d34565b80516016805460208085018051604080880180516060808b01805167ffffffffffffffff90811674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff9482166c0100000000000000000000000002949094167fffffffff00000000000000000000000000000000ffffffffffffffffffffffff978216640100000000027fffffffffffffffffffffffffffffffffffffffff000000000000000000000000909b1663ffffffff909d169c8d179a909a17969096169890981791909117909755815197885292518216938701939093529051811691850191909152905116908201527f1a9ff9caccd597226ff8c393ab44e6b57669905d9b16f1cdb2ac267253dbf27c9060800161155d565b611d7e612d34565b600c805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b6000611ea07fb9b8993db34ae003b2aacdae4cdef2888717531ab95157174f8f0dbf076b5e58604080516020808201939093527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166080808301919091528251808303909101815260a0909101909152805191012090565b905090565b611ead612d34565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b855185518560ff16601f831115611f4f576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401611230565b80600003611fb9576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401611230565b818314612047576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401611230565b612052816003615523565b83116120ba576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401611230565b6120c2612d34565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b601354156122855760135460009061211a90600190615560565b905060006013828154811061213157612131615475565b6000918252602082200154601480546001600160a01b039092169350908490811061215e5761215e615475565b60009182526020808320909101546001600160a01b0385811684526012909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000908116909155929091168084529220805490911690556013805491925090806121d1576121d1615577565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908101805473ffffffffffffffffffffffffffffffffffffffff19169055019055601480548061222f5761222f615577565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908101805473ffffffffffffffffffffffffffffffffffffffff1916905501905550612100915050565b60005b81515181101561264a57600060126000846000015184815181106122ae576122ae615475565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156122eb576122eb6149f1565b14612352576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401611230565b6040805180820190915260ff8216815260016020820152825180516012916000918590811061238357612383615475565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156123f9576123f96149f1565b0217905550600091506124099050565b601260008460200151848151811061242357612423615475565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115612460576124606149f1565b146124c7576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401611230565b6040805180820190915260ff8216815260208101600281525060126000846020015184815181106124fa576124fa615475565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612570576125706149f1565b02179055505082518051601392508390811061258e5761258e615475565b6020908102919091018101518254600181018455600093845292829020909201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390931692909217909155820151805160149190839081106125f2576125f2615475565b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909216919091179055612643816154d3565b9050612288565b5060408101516010805460ff191660ff909216919091179055601180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff4381168202928317855590830481169360019390926000926126be928692908216911617615694565b92506101000a81548163ffffffff021916908363ffffffff16021790555061271d4630601160009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a0015161375c565b600f819055825180516010805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560115460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05986127bc988b98919763ffffffff9092169690959194919391926156bc565b60405180910390a15050505050505050505050565b6127d9612d34565b6001600160a01b03821615806127f657506001600160a01b038116155b1561282d576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156128bc576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201805473ffffffffffffffffffffffffffffffffffffffff1916821790558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b6129b1612d34565b610865816137e9565b6000805b8251811015612ab9576000600660008584815181106129df576129df615475565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600003612a7257838281518110612a2857612a28615475565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401611230565b838281518110612a8457612a84615475565b60200260200101516020015181612a9b9190615523565b612aa5908461550b565b92505080612ab2906154d3565b90506129be565b508015610dcd57612aca60086135fc565b600954811115612b14576009546040517f688ccf77000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401611230565b600a54811115612b7457600854600a5460009190612b329084615560565b612b3c9190615752565b9050806040517fe31e0f3200000000000000000000000000000000000000000000000000000000815260040161123091815260200190565b8060086002016000828254612b899190615560565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba37890602001610913565b8151835114612bf9576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015612c5657612c46848281518110612c1a57612c1a615475565b6020026020010151848381518110612c3457612c34615475565b602002602001015160200151846138a5565b612c4f816154d3565b9050612bfc565b50505050565b60408101516001600160a01b03163b612c725750565b600c546040517fda52b4c40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063da52b4c490612cbb9084906004016157b3565b6020604051808303816000875af1158015612cda573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cfe9190615458565b610865576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005461010090046001600160a01b031633146109f35760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401611230565b60005460ff16612de55760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401611230565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60005460ff1615612e825760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611230565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ed5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ef99190615458565b15612f2f576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c546001600160a01b0316612f71576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6080820151516000819003612fb2576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115612fcd57612fcd6145d1565b60405190808252806020026020018201604052801561305457816020015b61304160408051610100810182526000808252602080830182905282840182905260608084018390526080840181905260a084015283518085019094528184528301529060c08201908152602001600081525090565b815260200190600190039081612feb5790505b50905060008267ffffffffffffffff811115613072576130726145d1565b60405190808252806020026020018201604052801561309b578160200160208202803683370190505b50905060006131697fb9b8993db34ae003b2aacdae4cdef2888717531ab95157174f8f0dbf076b5e58604080516020808201939093527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166080808301919091528251808303909101815260a0909101909152805191012090565b905060005b848110156132035760008760800151828151811061318e5761318e615475565b60200260200101518060200190518101906131a99190615957565b90506131b5818461391c565b8483815181106131c7576131c7615475565b602002602001018181525050808583815181106131e6576131e6615475565b602002602001018190525050806131fc906154d3565b905061316e565b50600080613225848960a001518a60c001518b60e001518c6101000151613a04565b601654919350915060009063ffffffff166132408442615560565b11905060005b8781101561355257600087828151811061326257613262615475565b602002602001015190506000613295826020015167ffffffffffffffff166000908152600e602052604090205460ff1690565b905060028160038111156132ab576132ab6149f1565b036132f45760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611230565b8a15806132fe5750835b8061331a57506003816003811115613318576133186149f1565b145b613350576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61335982613afc565b60008082600381111561336e5761336e6149f1565b14801561337957508b155b156133d2576133958a518761338e9190615752565b8e85613c58565b905060006133aa8460c0015160000151613e45565b90506133b78183306138a5565b818460c001516020018181516133cd9190615560565b905250505b60008260038111156133e6576133e66149f1565b146134225760208084015167ffffffffffffffff16600090815260158252604090205460c0850151909101805161341e908390615560565b9052505b60208381015167ffffffffffffffff166000908152600e90915260408120805460ff1916600117905561345c61345785613ea7565b614170565b60208086015167ffffffffffffffff166000908152600e909152604090208054919250829160ff19166001836003811115613499576134996149f1565b021790555060008360038111156134b2576134b26149f1565b1480156134d0575060038160038111156134ce576134ce6149f1565b145b156134f75760208085015167ffffffffffffffff1660009081526015909152604090208290555b836020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf826040516135359190614a20565b60405180910390a2505050508061354b906154d3565b9050613246565b50505050505050505050565b60005460ff16156135b15760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611230565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612e123390565b60008183106135f55781611c2d565b5090919050565b60018101546002820154429114806136175750808260030154145b15613620575050565b816001015482600201541115613662576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260030154826136749190615560565b6001840154845491925061369b9161368c9084615523565b8560020154611420919061550b565b600284015550600390910155565b6136c7818060200190518101906136c09190615c66565b6000612e2f565b505050565b60006136d9826020615523565b6136e4856020615523565b6136f08861014461550b565b6136fa919061550b565b613704919061550b565b61370f90600061550b565b9050368114613753576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401611230565b50505050505050565b6000808a8a8a8a8a8a8a8a8a60405160200161378099989796959493929190615dc3565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b038216036138415760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401611230565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b15801561390857600080fd5b505af1158015613753573d6000803e3d6000fd5b60008060001b828460200151856040015186606001518760800151805190602001208860a001516040516020016139539190615e4b565b604051602081830303815290604052805190602001208960e001518a60c001516040516020016139e6999897969594939291909889526020808a019890985267ffffffffffffffff9690961660408901526001600160a01b039485166060890152928416608088015260a087019190915260c086015260e085015281511661010084015201516101208201526101400190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600d546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce90613a61908c908c908c908c908c90600401615e8e565b6020604051808303816000875af1158015613a80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613aa49190615ee0565b905060008111613ae0576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a613aec9084615560565b9350935050509550959350505050565b80517f000000000000000000000000000000000000000000000000000000000000000014613b5c5780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401611230565b60165460a0820151517401000000000000000000000000000000000000000090910467ffffffffffffffff161015613bd25760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611230565b6016546080820151516c0100000000000000000000000090910467ffffffffffffffff161015610865576016546080820151516040517f869337890000000000000000000000000000000000000000000000000000000081526c0100000000000000000000000090920467ffffffffffffffff1660048301526024820152604401611230565b6000806000613c6e8460c0015160000151613e45565b6001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613cab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ccf9190615677565b905060005b856020015151811015613d4857816001600160a01b031686602001518281518110613d0157613d01615475565b60200260200101516001600160a01b031603613d385785604001518181518110613d2d57613d2d615475565b602002602001015192505b613d41816154d3565b9050613cd4565b5081613d8b576040517fce480bcc0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401611230565b6000670de0b6b3a7640000833a8760e00151613da78b8a610d12565b613db1919061550b565b613dbb9190615523565b613dc59190615523565b613dcf9190615752565b90508460c0015160200151811115613e3b5760208086015160c0870151909101516040517f394a2c2700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9092166004830152602482018390526044820152606401611230565b9695505050505050565b6001600160a01b038181166000908152600360205260409020541680613ea2576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401611230565b919050565b613ef06040518060e00160405280600081526020016060815260200160006001600160a01b03168152602001606081526020016060815260200160608152602001600081525090565b6000613f048360a001518460c001516142a4565b805190915060008167ffffffffffffffff811115613f2457613f246145d1565b604051908082528060200260200182016040528015613f6957816020015b6040805180820190915260008082526020820152815260200190600190039081613f425790505b50905060008267ffffffffffffffff811115613f8757613f876145d1565b604051908082528060200260200182016040528015613fb0578160200160208202803683370190505b50905060005b838110156140e6576000613fe6868381518110613fd557613fd5615475565b602002602001015160000151613e45565b905080838381518110613ffb57613ffb615475565b60200260200101906001600160a01b031690816001600160a01b0316815250506040518060400160405280826001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015614064573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906140889190615677565b6001600160a01b031681526020018784815181106140a8576140a8615475565b6020026020010151602001518152508483815181106140c9576140c9615475565b602002602001018190525050806140df906154d3565b9050613fb6565b506040518060e0016040528087600001518152602001876040015160405160200161412091906001600160a01b0391909116815260200190565b604051602081830303815290604052815260200187606001516001600160a01b03168152602001876080015181526020018281526020018381526020018760e00151815250945050505050919050565b6040517f07a22a0700000000000000000000000000000000000000000000000000000000815260009030906307a22a07906141af9085906004016157b3565b600060405180830381600087803b1580156141c957600080fd5b505af19250505080156141da575060015b61429c573d808015614208576040519150601f19603f3d011682016040523d82523d6000602084013e61420d565b606091505b5061421781615ef9565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da800000000000000000000000000000000000000000000000000000000036142685750600392915050565b806040517fcf19edfd0000000000000000000000000000000000000000000000000000000081526004016112309190614ab9565b506002919050565b606060005b83518110156144435782600001516001600160a01b03168482815181106142d2576142d2615475565b6020026020010151600001516001600160a01b031603614433576000845167ffffffffffffffff811115614308576143086145d1565b60405190808252806020026020018201604052801561434d57816020015b60408051808201909152600080825260208201528152602001906001900390816143265790505b50905060005b85518110156143a45785818151811061436e5761436e615475565b602002602001015182828151811061438857614388615475565b60200260200101819052508061439d906154d3565b9050614353565b5060405180604001604052808284815181106143c2576143c2615475565b6020026020010151600001516001600160a01b0316815260200185602001518385815181106143f3576143f3615475565b602002602001015160200151614409919061550b565b81525081838151811061441e5761441e615475565b60200260200101819052508092505050610dbd565b61443c816154d3565b90506142a9565b50600083516001614454919061550b565b67ffffffffffffffff81111561446c5761446c6145d1565b6040519080825280602002602001820160405280156144b157816020015b604080518082019091526000808252602082015281526020019060019003908161448a5790505b50905060005b8451811015614508578481815181106144d2576144d2615475565b60200260200101518282815181106144ec576144ec615475565b602002602001018190525080614501906154d3565b90506144b7565b50828185518151811061451d5761451d615475565b60209081029190910101529392505050565b828054828255906000526020600020908101928215614591579160200282015b82811115614591578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390911617825560209092019160019091019061454f565b50610e929291506145bc565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610e9257600081556001016145bd565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715614623576146236145d1565b60405290565b60405160e0810167ffffffffffffffff81118282101715614623576146236145d1565b604051610100810167ffffffffffffffff81118282101715614623576146236145d1565b604051610120810167ffffffffffffffff81118282101715614623576146236145d1565b604051601f8201601f1916810167ffffffffffffffff811182821017156146bd576146bd6145d1565b604052919050565b600067ffffffffffffffff8211156146df576146df6145d1565b50601f01601f191660200190565b600082601f8301126146fe57600080fd5b813561471161470c826146c5565b614694565b81815284602083860101111561472657600080fd5b816020850160208301376000918101602001919091529392505050565b6001600160a01b038116811461086557600080fd5b8035613ea281614743565b600067ffffffffffffffff82111561477d5761477d6145d1565b5060051b60200190565b600082601f83011261479857600080fd5b813560206147a861470c83614763565b82815260059290921b840181019181810190868411156147c757600080fd5b8286015b848110156147eb5780356147de81614743565b83529183019183016147cb565b509695505050505050565b60006040828403121561480857600080fd5b614810614600565b9050813561481d81614743565b808252506020820135602082015292915050565b600082601f83011261484257600080fd5b8135602061485261470c83614763565b82815260069290921b8401810191818101908684111561487157600080fd5b8286015b848110156147eb5761488788826147f6565b835291830191604001614875565b6000602082840312156148a757600080fd5b813567ffffffffffffffff808211156148bf57600080fd5b9083019060e082860312156148d357600080fd5b6148db614629565b823581526020830135828111156148f157600080fd5b6148fd878286016146ed565b60208301525061490f60408401614758565b604082015260608301358281111561492657600080fd5b614932878286016146ed565b60608301525060808301358281111561494a57600080fd5b61495687828601614787565b60808301525060a08301358281111561496e57600080fd5b61497a87828601614831565b60a08301525060c083013560c082015280935050505092915050565b6000602082840312156149a857600080fd5b8135611c2d81614743565b67ffffffffffffffff8116811461086557600080fd5b8035613ea2816149b3565b6000602082840312156149e657600080fd5b8135611c2d816149b3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310614a5b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b83811015614a7c578181015183820152602001614a64565b83811115612c565750506000910152565b60008151808452614aa5816020860160208601614a61565b601f01601f19169290920160200192915050565b602081526000611c2d6020830184614a8d565b600082601f830112614add57600080fd5b81356020614aed61470c83614763565b82815260059290921b84018101918181019086841115614b0c57600080fd5b8286015b848110156147eb5780358352918301918301614b10565b60008060408385031215614b3a57600080fd5b823567ffffffffffffffff80821115614b5257600080fd5b614b5e86838701614787565b93506020850135915080821115614b7457600080fd5b50614b8185828601614acc565b9150509250929050565b600060208284031215614b9d57600080fd5b813567ffffffffffffffff811115614bb457600080fd5b614bc084828501614787565b949350505050565b6020808252825182820181905260009190848201906040850190845b81811015614c0057835183529284019291840191600101614be4565b50909695505050505050565b60008060408385031215614c1f57600080fd5b82359150602083013567ffffffffffffffff80821115614c3e57600080fd5b908401906101208287031215614c5357600080fd5b614c5b61464c565b82358152614c6b602084016149c9565b6020820152614c7c60408401614758565b6040820152614c8d60608401614758565b6060820152608083013582811115614ca457600080fd5b614cb0888286016146ed565b60808301525060a083013582811115614cc857600080fd5b614cd488828601614831565b60a083015250614ce78760c085016147f6565b60c082015261010083013560e08201528093505050509250929050565b600082601f830112614d1557600080fd5b81356020614d2561470c83614763565b82815260059290921b84018101918181019086841115614d4457600080fd5b8286015b848110156147eb578035614d5b816149b3565b8352918301918301614d48565b600060408284031215614d7a57600080fd5b614d82614600565b9050813581526020820135602082015292915050565b600082601f830112614da957600080fd5b81356020614db961470c83614763565b82815260069290921b84018101918181019086841115614dd857600080fd5b8286015b848110156147eb57614dee8882614d68565b835291830191604001614ddc565b600082601f830112614e0d57600080fd5b81356020614e1d61470c83614763565b82815260059290921b84018101918181019086841115614e3c57600080fd5b8286015b848110156147eb57803567ffffffffffffffff811115614e605760008081fd5b614e6e8986838b01016146ed565b845250918301918301614e40565b60006101208284031215614e8f57600080fd5b614e97614670565b9050813567ffffffffffffffff80821115614eb157600080fd5b614ebd85838601614d04565b83526020840135915080821115614ed357600080fd5b614edf85838601614787565b60208401526040840135915080821115614ef857600080fd5b614f0485838601614acc565b60408401526060840135915080821115614f1d57600080fd5b614f2985838601614d98565b60608401526080840135915080821115614f4257600080fd5b614f4e85838601614dfc565b608084015260a0840135915080821115614f6757600080fd5b614f7385838601614acc565b60a084015260c084013560c084015260e0840135915080821115614f9657600080fd5b50614fa384828501614acc565b60e08301525061010080830135818301525092915050565b801515811461086557600080fd5b60008060408385031215614fdc57600080fd5b823567ffffffffffffffff811115614ff357600080fd5b614fff85828601614e7c565b925050602083013561501081614fbb565b809150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614c005783516001600160a01b031683529284019291840191600101615037565b6000806040838503121561506f57600080fd5b823561507a81614743565b9150602083013561501081614743565b600081518084526020808501945080840160005b838110156150c35781516001600160a01b03168752958201959082019060010161509e565b509495945050505050565b602081526000611c2d602083018461508a565b6000604082840312156150f357600080fd5b611c2d8383614d68565b60006020828403121561510f57600080fd5b813567ffffffffffffffff81111561512657600080fd5b820160e08185031215611c2d57600080fd5b60006020828403121561514a57600080fd5b5035919050565b60006020828403121561516357600080fd5b813567ffffffffffffffff81111561517a57600080fd5b614bc084828501614e7c565b60008083601f84011261519857600080fd5b50813567ffffffffffffffff8111156151b057600080fd5b6020830191508360208260051b85010111156151cb57600080fd5b9250929050565b60008060008060008060008060e0898b0312156151ee57600080fd5b606089018a8111156151ff57600080fd5b8998503567ffffffffffffffff8082111561521957600080fd5b818b0191508b601f83011261522d57600080fd5b81358181111561523c57600080fd5b8c602082850101111561524e57600080fd5b6020830199508098505060808b013591508082111561526c57600080fd5b6152788c838d01615186565b909750955060a08b013591508082111561529157600080fd5b5061529e8b828c01615186565b999c989b50969995989497949560c00135949350505050565b6000602082840312156152c957600080fd5b813567ffffffffffffffff8111156152e057600080fd5b614bc0848285016146ed565b6000608082840312156152fe57600080fd5b6040516080810181811067ffffffffffffffff82111715615321576153216145d1565b604052823563ffffffff8116811461533857600080fd5b81526020830135615348816149b3565b6020820152604083013561535b816149b3565b6040820152606083013561536e816149b3565b60608201529392505050565b803560ff81168114613ea257600080fd5b60008060008060008060c087890312156153a457600080fd5b863567ffffffffffffffff808211156153bc57600080fd5b6153c88a838b01614787565b975060208901359150808211156153de57600080fd5b6153ea8a838b01614787565b96506153f860408a0161537a565b9550606089013591508082111561540e57600080fd5b61541a8a838b016146ed565b945061542860808a016149c9565b935060a089013591508082111561543e57600080fd5b5061544b89828a016146ed565b9150509295509295509295565b60006020828403121561546a57600080fd5b8151611c2d81614fbb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203615504576155046154a4565b5060010190565b6000821982111561551e5761551e6154a4565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561555b5761555b6154a4565b500290565b600082821015615572576155726154a4565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff038211156155c3576155c36154a4565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff83168061560d5761560d6155cb565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60006020828403121561568957600080fd5b8151611c2d81614743565b600063ffffffff8083168185168083038211156156b3576156b36154a4565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526156ec8184018a61508a565b90508281036080840152615700818961508a565b905060ff871660a084015282810360c084015261571d8187614a8d565b905067ffffffffffffffff851660e08401528281036101008401526157428185614a8d565b9c9b505050505050505050505050565b600082615761576157616155cb565b500490565b600081518084526020808501945080840160005b838110156150c3576157a087835180516001600160a01b03168252602090810151910152565b604096909601959082019060010161577a565b60208152815160208201526000602083015160e060408401526157da610100840182614a8d565b90506001600160a01b0360408501511660608401526060840151601f198085840301608086015261580b8383614a8d565b925060808601519150808584030160a0860152615828838361508a565b925060a08601519150808584030160c0860152506158468282615766565b91505060c084015160e08401528091505092915050565b8051613ea2816149b3565b8051613ea281614743565b600082601f83011261588457600080fd5b815161589261470c826146c5565b8181528460208386010111156158a757600080fd5b614bc0826020830160208701614a61565b6000604082840312156158ca57600080fd5b6158d2614600565b905081516158df81614743565b808252506020820151602082015292915050565b600082601f83011261590457600080fd5b8151602061591461470c83614763565b82815260069290921b8401810191818101908684111561593357600080fd5b8286015b848110156147eb5761594988826158b8565b835291830191604001615937565b60006020828403121561596957600080fd5b815167ffffffffffffffff8082111561598157600080fd5b90830190610120828603121561599657600080fd5b61599e61464c565b825181526159ae6020840161585d565b60208201526159bf60408401615868565b60408201526159d060608401615868565b60608201526080830151828111156159e757600080fd5b6159f387828601615873565b60808301525060a083015182811115615a0b57600080fd5b615a17878286016158f3565b60a083015250615a2a8660c085016158b8565b60c0820152610100929092015160e0830152509392505050565b600082601f830112615a5557600080fd5b81516020615a6561470c83614763565b82815260059290921b84018101918181019086841115615a8457600080fd5b8286015b848110156147eb578051615a9b816149b3565b8352918301918301615a88565b600082601f830112615ab957600080fd5b81516020615ac961470c83614763565b82815260059290921b84018101918181019086841115615ae857600080fd5b8286015b848110156147eb578051615aff81614743565b8352918301918301615aec565b600082601f830112615b1d57600080fd5b81516020615b2d61470c83614763565b82815260059290921b84018101918181019086841115615b4c57600080fd5b8286015b848110156147eb5780518352918301918301615b50565b600082601f830112615b7857600080fd5b81516020615b8861470c83614763565b82815260069290921b84018101918181019086841115615ba757600080fd5b8286015b848110156147eb5760408189031215615bc45760008081fd5b615bcc614600565b815181528482015185820152835291830191604001615bab565b600082601f830112615bf757600080fd5b81516020615c0761470c83614763565b82815260059290921b84018101918181019086841115615c2657600080fd5b8286015b848110156147eb57805167ffffffffffffffff811115615c4a5760008081fd5b615c588986838b0101615873565b845250918301918301615c2a565b600060208284031215615c7857600080fd5b815167ffffffffffffffff80821115615c9057600080fd5b908301906101208286031215615ca557600080fd5b615cad614670565b825182811115615cbc57600080fd5b615cc887828601615a44565b825250602083015182811115615cdd57600080fd5b615ce987828601615aa8565b602083015250604083015182811115615d0157600080fd5b615d0d87828601615b0c565b604083015250606083015182811115615d2557600080fd5b615d3187828601615b67565b606083015250608083015182811115615d4957600080fd5b615d5587828601615be6565b60808301525060a083015182811115615d6d57600080fd5b615d7987828601615b0c565b60a08301525060c083015160c082015260e083015182811115615d9b57600080fd5b615da787828601615b0c565b60e0830152506101009283015192810192909252509392505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b166040850152816060850152615dfd8285018b61508a565b91508382036080850152615e11828a61508a565b915060ff881660a085015283820360c0850152615e2e8288614a8d565b90861660e085015283810361010085015290506157428185614a8d565b602081526000611c2d6020830184615766565b600081518084526020808501945080840160005b838110156150c357815187529582019590820190600101615e72565b60a081526000615ea160a0830188615e5e565b8281036020840152615eb38188615e5e565b90508560408401528281036060840152615ecd8186615e5e565b9150508260808301529695505050505050565b600060208284031215615ef257600080fd5b5051919050565b6000815160208301517fffffffff0000000000000000000000000000000000000000000000000000000080821693506004831015615f415780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
}

var EVM2EVMTollOffRampHelperABI = EVM2EVMTollOffRampHelperMetaData.ABI

var EVM2EVMTollOffRampHelperBin = EVM2EVMTollOffRampHelperMetaData.Bin

func DeployEVM2EVMTollOffRampHelper(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, onRampAddress common.Address, commitStore common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address) (common.Address, *types.Transaction, *EVM2EVMTollOffRampHelper, error) {
	parsed, err := EVM2EVMTollOffRampHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMTollOffRampHelperBin), backend, sourceChainId, chainId, offRampConfig, onRampAddress, commitStore, afn, sourceTokens, pools, rateLimiterConfig, tokenLimitsAdmin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMTollOffRampHelper{EVM2EVMTollOffRampHelperCaller: EVM2EVMTollOffRampHelperCaller{contract: contract}, EVM2EVMTollOffRampHelperTransactor: EVM2EVMTollOffRampHelperTransactor{contract: contract}, EVM2EVMTollOffRampHelperFilterer: EVM2EVMTollOffRampHelperFilterer{contract: contract}}, nil
}

type EVM2EVMTollOffRampHelper struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMTollOffRampHelperCaller
	EVM2EVMTollOffRampHelperTransactor
	EVM2EVMTollOffRampHelperFilterer
}

type EVM2EVMTollOffRampHelperCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampHelperTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampHelperFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampHelperSession struct {
	Contract     *EVM2EVMTollOffRampHelper
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOffRampHelperCallerSession struct {
	Contract *EVM2EVMTollOffRampHelperCaller
	CallOpts bind.CallOpts
}

type EVM2EVMTollOffRampHelperTransactorSession struct {
	Contract     *EVM2EVMTollOffRampHelperTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOffRampHelperRaw struct {
	Contract *EVM2EVMTollOffRampHelper
}

type EVM2EVMTollOffRampHelperCallerRaw struct {
	Contract *EVM2EVMTollOffRampHelperCaller
}

type EVM2EVMTollOffRampHelperTransactorRaw struct {
	Contract *EVM2EVMTollOffRampHelperTransactor
}

func NewEVM2EVMTollOffRampHelper(address common.Address, backend bind.ContractBackend) (*EVM2EVMTollOffRampHelper, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMTollOffRampHelperABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMTollOffRampHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelper{address: address, abi: abi, EVM2EVMTollOffRampHelperCaller: EVM2EVMTollOffRampHelperCaller{contract: contract}, EVM2EVMTollOffRampHelperTransactor: EVM2EVMTollOffRampHelperTransactor{contract: contract}, EVM2EVMTollOffRampHelperFilterer: EVM2EVMTollOffRampHelperFilterer{contract: contract}}, nil
}

func NewEVM2EVMTollOffRampHelperCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMTollOffRampHelperCaller, error) {
	contract, err := bindEVM2EVMTollOffRampHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperCaller{contract: contract}, nil
}

func NewEVM2EVMTollOffRampHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMTollOffRampHelperTransactor, error) {
	contract, err := bindEVM2EVMTollOffRampHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperTransactor{contract: contract}, nil
}

func NewEVM2EVMTollOffRampHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMTollOffRampHelperFilterer, error) {
	contract, err := bindEVM2EVMTollOffRampHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperFilterer{contract: contract}, nil
}

func bindEVM2EVMTollOffRampHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMTollOffRampHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOffRampHelper.Contract.EVM2EVMTollOffRampHelperCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.EVM2EVMTollOffRampHelperTransactor.contract.Transfer(opts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.EVM2EVMTollOffRampHelperTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOffRampHelper.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.contract.Transfer(opts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(AggregateRateLimiterInterfaceTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(AggregateRateLimiterInterfaceTokenBucket)).(*AggregateRateLimiterInterfaceTokenBucket)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMTollOffRampHelper.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMTollOffRampHelper.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMTollOffRampHelper.Contract.CcipReceive(&_EVM2EVMTollOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMTollOffRampHelper.Contract.CcipReceive(&_EVM2EVMTollOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) FeeTaken(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "feeTaken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) FeeTaken(arg0 *big.Int) (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.FeeTaken(&_EVM2EVMTollOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) FeeTaken(arg0 *big.Int) (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.FeeTaken(&_EVM2EVMTollOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetAFN(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetAFN(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getChainIDs")

	outstruct := new(GetChainIDs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetChainIDs(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetChainIDs(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetCommitStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getCommitStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetCommitStore() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetCommitStore(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetCommitStore() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetCommitStore(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOffRampInterfaceOffRampConfig)).(*BaseOffRampInterfaceOffRampConfig)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetConfig(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetConfig(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetDestinationToken(&_EVM2EVMTollOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetDestinationToken(&_EVM2EVMTollOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetDestinationTokens(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetDestinationTokens(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetExecutionState(&_EVM2EVMTollOffRampHelper.CallOpts, sequenceNumber)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetExecutionState(&_EVM2EVMTollOffRampHelper.CallOpts, sequenceNumber)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPool(&_EVM2EVMTollOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPool(&_EVM2EVMTollOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPoolTokens(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPoolTokens(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPricesForTokens(&_EVM2EVMTollOffRampHelper.CallOpts, tokens)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPricesForTokens(&_EVM2EVMTollOffRampHelper.CallOpts, tokens)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetRouter(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetRouter(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetTokenLimitAdmin(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetTokenLimitAdmin(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMTollOffRampHelper.Contract.IsAFNHealthy(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMTollOffRampHelper.Contract.IsAFNHealthy(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMTollOffRampHelper.Contract.LatestConfigDetails(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMTollOffRampHelper.Contract.LatestConfigDetails(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMTollOffRampHelper.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMTollOffRampHelper.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) MetadataHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "metadataHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) MetadataHash() ([32]byte, error) {
	return _EVM2EVMTollOffRampHelper.Contract.MetadataHash(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) MetadataHash() ([32]byte, error) {
	return _EVM2EVMTollOffRampHelper.Contract.MetadataHash(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) OverheadGasToll(opts *bind.CallOpts, merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "overheadGasToll", merkleGasShare, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) OverheadGasToll(merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.OverheadGasToll(&_EVM2EVMTollOffRampHelper.CallOpts, merkleGasShare, message)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) OverheadGasToll(merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.OverheadGasToll(&_EVM2EVMTollOffRampHelper.CallOpts, merkleGasShare, message)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Owner(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Owner(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Paused() (bool, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Paused(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) Paused() (bool, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Paused(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Transmitters(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Transmitters(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOffRampHelper.Contract.TypeAndVersion(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOffRampHelper.Contract.TypeAndVersion(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.AcceptOwnership(&_EVM2EVMTollOffRampHelper.TransactOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.AcceptOwnership(&_EVM2EVMTollOffRampHelper.TransactOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.AddPool(&_EVM2EVMTollOffRampHelper.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.AddPool(&_EVM2EVMTollOffRampHelper.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) Execute(opts *bind.TransactOpts, rep CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "execute", rep, manualExecution)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Execute(rep CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Execute(&_EVM2EVMTollOffRampHelper.TransactOpts, rep, manualExecution)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) Execute(rep CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Execute(&_EVM2EVMTollOffRampHelper.TransactOpts, rep, manualExecution)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "executeSingleMessage", message)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRampHelper.TransactOpts, message)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRampHelper.TransactOpts, message)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) ManuallyExecute(opts *bind.TransactOpts, report CCIPExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "manuallyExecute", report)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) ManuallyExecute(report CCIPExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.ManuallyExecute(&_EVM2EVMTollOffRampHelper.TransactOpts, report)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) ManuallyExecute(report CCIPExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.ManuallyExecute(&_EVM2EVMTollOffRampHelper.TransactOpts, report)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "pause")
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Pause(&_EVM2EVMTollOffRampHelper.TransactOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Pause(&_EVM2EVMTollOffRampHelper.TransactOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.RemovePool(&_EVM2EVMTollOffRampHelper.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.RemovePool(&_EVM2EVMTollOffRampHelper.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "report", executableMessages)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Report(executableMessages []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Report(&_EVM2EVMTollOffRampHelper.TransactOpts, executableMessages)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) Report(executableMessages []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Report(&_EVM2EVMTollOffRampHelper.TransactOpts, executableMessages)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetAFN(&_EVM2EVMTollOffRampHelper.TransactOpts, afn)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetAFN(&_EVM2EVMTollOffRampHelper.TransactOpts, afn)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetCommitStore(opts *bind.TransactOpts, commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setCommitStore", commitStore)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetCommitStore(commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetCommitStore(&_EVM2EVMTollOffRampHelper.TransactOpts, commitStore)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetCommitStore(commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetCommitStore(&_EVM2EVMTollOffRampHelper.TransactOpts, commitStore)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetConfig(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig(&_EVM2EVMTollOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig(&_EVM2EVMTollOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetConfig0(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setConfig0", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig0(&_EVM2EVMTollOffRampHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig0(&_EVM2EVMTollOffRampHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetPrices(&_EVM2EVMTollOffRampHelper.TransactOpts, tokens, prices)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetPrices(&_EVM2EVMTollOffRampHelper.TransactOpts, tokens, prices)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetRateLimiterConfig(&_EVM2EVMTollOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetRateLimiterConfig(&_EVM2EVMTollOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetRouter(&_EVM2EVMTollOffRampHelper.TransactOpts, router)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetRouter(&_EVM2EVMTollOffRampHelper.TransactOpts, router)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetTokenLimitAdmin(&_EVM2EVMTollOffRampHelper.TransactOpts, newAdmin)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetTokenLimitAdmin(&_EVM2EVMTollOffRampHelper.TransactOpts, newAdmin)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.TransferOwnership(&_EVM2EVMTollOffRampHelper.TransactOpts, to)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.TransferOwnership(&_EVM2EVMTollOffRampHelper.TransactOpts, to)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Transmit(&_EVM2EVMTollOffRampHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Transmit(&_EVM2EVMTollOffRampHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "unpause")
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Unpause(&_EVM2EVMTollOffRampHelper.TransactOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Unpause(&_EVM2EVMTollOffRampHelper.TransactOpts)
}

type EVM2EVMTollOffRampHelperAFNSetIterator struct {
	Event *EVM2EVMTollOffRampHelperAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperAFNSet)
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
		it.Event = new(EVM2EVMTollOffRampHelperAFNSet)
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

func (it *EVM2EVMTollOffRampHelperAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperAFNSetIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperAFNSet)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseAFNSet(log types.Log) (*EVM2EVMTollOffRampHelperAFNSet, error) {
	event := new(EVM2EVMTollOffRampHelperAFNSet)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperConfigChangedIterator struct {
	Event *EVM2EVMTollOffRampHelperConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperConfigChanged)
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
		it.Event = new(EVM2EVMTollOffRampHelperConfigChanged)
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

func (it *EVM2EVMTollOffRampHelperConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperConfigChangedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperConfigChanged)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMTollOffRampHelperConfigChanged, error) {
	event := new(EVM2EVMTollOffRampHelperConfigChanged)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperConfigSetIterator struct {
	Event *EVM2EVMTollOffRampHelperConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperConfigSet)
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
		it.Event = new(EVM2EVMTollOffRampHelperConfigSet)
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

func (it *EVM2EVMTollOffRampHelperConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperConfigSet struct {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperConfigSetIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperConfigSet)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseConfigSet(log types.Log) (*EVM2EVMTollOffRampHelperConfigSet, error) {
	event := new(EVM2EVMTollOffRampHelperConfigSet)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperExecutionStateChangedIterator struct {
	Event *EVM2EVMTollOffRampHelperExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperExecutionStateChanged)
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
		it.Event = new(EVM2EVMTollOffRampHelperExecutionStateChanged)
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

func (it *EVM2EVMTollOffRampHelperExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperExecutionStateChanged struct {
	SequenceNumber uint64
	State          uint8
	Raw            types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMTollOffRampHelperExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperExecutionStateChangedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperExecutionStateChanged)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseExecutionStateChanged(log types.Log) (*EVM2EVMTollOffRampHelperExecutionStateChanged, error) {
	event := new(EVM2EVMTollOffRampHelperExecutionStateChanged)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperOffRampConfigSetIterator struct {
	Event *EVM2EVMTollOffRampHelperOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperOffRampConfigSet)
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
		it.Event = new(EVM2EVMTollOffRampHelperOffRampConfigSet)
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

func (it *EVM2EVMTollOffRampHelperOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperOffRampConfigSet struct {
	Config BaseOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperOffRampConfigSetIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperOffRampConfigSet)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseOffRampConfigSet(log types.Log) (*EVM2EVMTollOffRampHelperOffRampConfigSet, error) {
	event := new(EVM2EVMTollOffRampHelperOffRampConfigSet)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperOffRampRouterSetIterator struct {
	Event *EVM2EVMTollOffRampHelperOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperOffRampRouterSet)
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
		it.Event = new(EVM2EVMTollOffRampHelperOffRampRouterSet)
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

func (it *EVM2EVMTollOffRampHelperOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperOffRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMTollOffRampHelperOffRampRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperOffRampRouterSetIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOffRampRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperOffRampRouterSet)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseOffRampRouterSet(log types.Log) (*EVM2EVMTollOffRampHelperOffRampRouterSet, error) {
	event := new(EVM2EVMTollOffRampHelperOffRampRouterSet)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMTollOffRampHelperOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMTollOffRampHelperOwnershipTransferRequested)
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

func (it *EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperOwnershipTransferRequested)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOffRampHelperOwnershipTransferRequested, error) {
	event := new(EVM2EVMTollOffRampHelperOwnershipTransferRequested)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperOwnershipTransferredIterator struct {
	Event *EVM2EVMTollOffRampHelperOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperOwnershipTransferred)
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
		it.Event = new(EVM2EVMTollOffRampHelperOwnershipTransferred)
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

func (it *EVM2EVMTollOffRampHelperOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampHelperOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperOwnershipTransferredIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperOwnershipTransferred)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOffRampHelperOwnershipTransferred, error) {
	event := new(EVM2EVMTollOffRampHelperOwnershipTransferred)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperPausedIterator struct {
	Event *EVM2EVMTollOffRampHelperPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperPaused)
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
		it.Event = new(EVM2EVMTollOffRampHelperPaused)
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

func (it *EVM2EVMTollOffRampHelperPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperPausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperPausedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperPaused)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParsePaused(log types.Log) (*EVM2EVMTollOffRampHelperPaused, error) {
	event := new(EVM2EVMTollOffRampHelperPaused)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperPoolAddedIterator struct {
	Event *EVM2EVMTollOffRampHelperPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperPoolAdded)
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
		it.Event = new(EVM2EVMTollOffRampHelperPoolAdded)
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

func (it *EVM2EVMTollOffRampHelperPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperPoolAddedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperPoolAdded)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMTollOffRampHelperPoolAdded, error) {
	event := new(EVM2EVMTollOffRampHelperPoolAdded)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperPoolRemovedIterator struct {
	Event *EVM2EVMTollOffRampHelperPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperPoolRemoved)
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
		it.Event = new(EVM2EVMTollOffRampHelperPoolRemoved)
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

func (it *EVM2EVMTollOffRampHelperPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperPoolRemovedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperPoolRemoved)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMTollOffRampHelperPoolRemoved, error) {
	event := new(EVM2EVMTollOffRampHelperPoolRemoved)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperTokenPriceChangedIterator struct {
	Event *EVM2EVMTollOffRampHelperTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperTokenPriceChanged)
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
		it.Event = new(EVM2EVMTollOffRampHelperTokenPriceChanged)
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

func (it *EVM2EVMTollOffRampHelperTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperTokenPriceChangedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperTokenPriceChanged)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMTollOffRampHelperTokenPriceChanged, error) {
	event := new(EVM2EVMTollOffRampHelperTokenPriceChanged)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMTollOffRampHelperTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMTollOffRampHelperTokensRemovedFromBucket)
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

func (it *EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperTokensRemovedFromBucket)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMTollOffRampHelperTokensRemovedFromBucket, error) {
	event := new(EVM2EVMTollOffRampHelperTokensRemovedFromBucket)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperTransmittedIterator struct {
	Event *EVM2EVMTollOffRampHelperTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperTransmitted)
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
		it.Event = new(EVM2EVMTollOffRampHelperTransmitted)
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

func (it *EVM2EVMTollOffRampHelperTransmittedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperTransmittedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperTransmittedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperTransmitted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperTransmitted)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseTransmitted(log types.Log) (*EVM2EVMTollOffRampHelperTransmitted, error) {
	event := new(EVM2EVMTollOffRampHelperTransmitted)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperUnpausedIterator struct {
	Event *EVM2EVMTollOffRampHelperUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperUnpaused)
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
		it.Event = new(EVM2EVMTollOffRampHelperUnpaused)
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

func (it *EVM2EVMTollOffRampHelperUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperUnpausedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperUnpaused)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseUnpaused(log types.Log) (*EVM2EVMTollOffRampHelperUnpaused, error) {
	event := new(EVM2EVMTollOffRampHelperUnpaused)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelper) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMTollOffRampHelper.abi.Events["AFNSet"].ID:
		return _EVM2EVMTollOffRampHelper.ParseAFNSet(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMTollOffRampHelper.ParseConfigChanged(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["ConfigSet"].ID:
		return _EVM2EVMTollOffRampHelper.ParseConfigSet(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMTollOffRampHelper.ParseExecutionStateChanged(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["OffRampConfigSet"].ID:
		return _EVM2EVMTollOffRampHelper.ParseOffRampConfigSet(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["OffRampRouterSet"].ID:
		return _EVM2EVMTollOffRampHelper.ParseOffRampRouterSet(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMTollOffRampHelper.ParseOwnershipTransferRequested(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMTollOffRampHelper.ParseOwnershipTransferred(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["Paused"].ID:
		return _EVM2EVMTollOffRampHelper.ParsePaused(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["PoolAdded"].ID:
		return _EVM2EVMTollOffRampHelper.ParsePoolAdded(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMTollOffRampHelper.ParsePoolRemoved(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMTollOffRampHelper.ParseTokenPriceChanged(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMTollOffRampHelper.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["Transmitted"].ID:
		return _EVM2EVMTollOffRampHelper.ParseTransmitted(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["Unpaused"].ID:
		return _EVM2EVMTollOffRampHelper.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMTollOffRampHelperAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMTollOffRampHelperConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMTollOffRampHelperConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMTollOffRampHelperExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf")
}

func (EVM2EVMTollOffRampHelperOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1a9ff9caccd597226ff8c393ab44e6b57669905d9b16f1cdb2ac267253dbf27c")
}

func (EVM2EVMTollOffRampHelperOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
}

func (EVM2EVMTollOffRampHelperOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMTollOffRampHelperOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMTollOffRampHelperPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMTollOffRampHelperPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMTollOffRampHelperPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMTollOffRampHelperTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMTollOffRampHelperTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMTollOffRampHelperTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (EVM2EVMTollOffRampHelperUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelper) Address() common.Address {
	return _EVM2EVMTollOffRampHelper.address
}

type EVM2EVMTollOffRampHelperInterface interface {
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

	MetadataHash(opts *bind.CallOpts) ([32]byte, error)

	OverheadGasToll(opts *bind.CallOpts, merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, rep CCIPExecutionReport, manualExecution bool) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report CCIPExecutionReport) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error)

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

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMTollOffRampHelperAFNSet, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMTollOffRampHelperConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMTollOffRampHelperConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMTollOffRampHelperExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMTollOffRampHelperExecutionStateChanged, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*EVM2EVMTollOffRampHelperOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMTollOffRampHelperOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOffRampRouterSet, router []common.Address) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*EVM2EVMTollOffRampHelperOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOffRampHelperOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampHelperOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOffRampHelperOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMTollOffRampHelperPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMTollOffRampHelperPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMTollOffRampHelperPoolRemoved, error)

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMTollOffRampHelperTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMTollOffRampHelperTokensRemovedFromBucket, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMTollOffRampHelperTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMTollOffRampHelperUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
