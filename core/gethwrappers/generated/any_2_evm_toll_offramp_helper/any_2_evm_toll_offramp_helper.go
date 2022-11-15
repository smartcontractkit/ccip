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
	EncodedMessages          [][]byte
	InnerProofs              [][32]byte
	InnerProofFlagBits       *big.Int
	OuterProofs              [][32]byte
	OuterProofFlagBits       *big.Int
}

var EVM2EVMTollOffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"rep\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeTaken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitStore\",\"outputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"merkleGasShare\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount\",\"name\":\"feeTokenAndAmount\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVM2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"overheadGasToll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"}],\"name\":\"setCommitStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620065b7380380620065b78339810160408190526200003491620006d4565b8888888888888888886001898989898989898989818185858833806000806000806101000a81548160ff02191690831515021790555060006001600160a01b0316826001600160a01b031603620000d25760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b0319909216919091179091558116156200010c576200010c816200041d565b5050506001600160a01b0381166200013757604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03929092169190911790558051825114620001795760405162d8548360e71b815260040160405180910390fd5b81516200018e906004906020850190620004ce565b5060005b825181101562000259576000828281518110620001b357620001b36200082e565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001fd57620001fd6200082e565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b029290911691909117905550620002518162000844565b905062000192565b5050600580546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600855600993909355600a55600b91909155875116620002d7576040516342bcdf7f60e11b815260040160405180910390fd5b88608081815250508760a0818152505086600e60008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160000160146101000a81548163ffffffff021916908363ffffffff16021790555060408201518160000160186101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160010160006101000a8154816001600160401b0302191690836001600160401b0316021790555060808201518160010160086101000a8154816001600160401b0302191690836001600160401b0316021790555090505085600d60006101000a8154816001600160a01b0302191690836001600160a01b0316021790555050505050505050505080151560c081151581525050505050505050505050505050505050505050506200086c565b336001600160a01b03821603620004775760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000c9565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b82805482825590600052602060002090810192821562000526579160200282015b828111156200052657825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620004ef565b506200053492915062000538565b5090565b5b8082111562000534576000815560010162000539565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b03811182821017156200058a576200058a6200054f565b60405290565b6001600160a01b0381168114620005a657600080fd5b50565b8051620005b68162000590565b919050565b80516001600160401b0381168114620005b657600080fd5b600082601f830112620005e557600080fd5b815160206001600160401b03808311156200060457620006046200054f565b8260051b604051601f19603f830116810181811084821117156200062c576200062c6200054f565b6040529384528581018301938381019250878511156200064b57600080fd5b83870191505b8482101562000677578151620006678162000590565b8352918301919083019062000651565b979650505050505050565b6000604082840312156200069557600080fd5b604080519081016001600160401b0381118282101715620006ba57620006ba6200054f565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000898b036101c0811215620006f557600080fd5b8a51995060208b0151985060a0603f19820112156200071357600080fd5b506200071e62000565565b60408b01516200072e8162000590565b815260608b015163ffffffff811681146200074857600080fd5b60208201526200075b60808c01620005bb565b60408201526200076e60a08c01620005bb565b60608201526200078160c08c01620005bb565b608082015296506200079660e08b01620005a9565b9550620007a76101008b01620005a9565b6101208b01519095506001600160401b0380821115620007c657600080fd5b620007d48d838e01620005d3565b95506101408c0151915080821115620007ec57600080fd5b50620007fb8c828d01620005d3565b9350506200080e8b6101608c0162000682565b91506200081f6101a08b01620005a9565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b6000600182016200086557634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c051615cf8620008bf60003960006116b801526000818161032c01528181611d2c015261331201526000818161030901528181611d07015281816132ed0152613b130152615cf86000f3fe608060405234801561001057600080fd5b50600436106102e95760003560e01c80638da5cb5b11610191578063bbe4f6db116100e3578063c9029f6a11610097578063e3d0e71211610071578063e3d0e712146107f6578063eb511dd414610809578063f2fde38b1461081c57600080fd5b8063c9029f6a146107bd578063c9033284146107d0578063d30a364b146107e357600080fd5b8063c0d78655116100c8578063c0d786551461068a578063c3f909d41461069d578063c5a1d7f0146107b557600080fd5b8063bbe4f6db1461064b578063be9b03f11461067757600080fd5b8063afcb95d711610145578063b4069b311161011f578063b4069b3114610614578063b576716614610627578063b66f0efb1461063a57600080fd5b8063afcb95d7146105d0578063b0f479a1146105f0578063b1dc65a41461060157600080fd5b80639187254311610176578063918725431461058f5780639438ff63146105a2578063a8b640c1146105b057600080fd5b80638da5cb5b1461053e57806390c2339b1461055457600080fd5b80634741062e1161024a578063744b92e2116101fe57806381ff7048116101d857806381ff7048146104fe5780638456cb591461052e57806389c065681461053657600080fd5b8063744b92e2146104ce57806379ba5097146104e157806381411834146104e957600080fd5b80635b4dc8121161022f5780635b4dc8121461048d5780635c975abb146104ae578063681fba16146104b957600080fd5b80634741062e1461045c578063599f64311461047c57600080fd5b8063181f5a77116102a157806339aa92641161028657806339aa92641461042e5780633f4ba83a146104415780634352fa9f1461044957600080fd5b8063181f5a77146103c05780632222dd421461040957600080fd5b8063108ee5fc116102d2578063108ee5fc1461035b578063142a98fc1461036e578063147809b3146103a857600080fd5b806307a22a07146102ee578063087ae6df14610303575b600080fd5b6103016102fc36600461474f565b61082f565b005b604080517f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006020820152015b60405180910390f35b610301610369366004614850565b6108a3565b61039b61037c36600461488e565b67ffffffffffffffff1660009081526010602052604090205460ff1690565b60405161035291906148da565b6103b061095a565b6040519015158152602001610352565b6103fc6040518060400160405280601881526020017f45564d3245564d546f6c6c4f666652616d7020312e302e30000000000000000081525081565b6040516103529190614973565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610352565b61030161043c366004614850565b6109e7565b610301610a1e565b6103016104573660046149e1565b610a30565b61046f61046a366004614a45565b610c85565b6040516103529190614a82565b6005546001600160a01b0316610416565b6104a061049b366004614ac6565b610d4d565b604051908152602001610352565b60005460ff166103b0565b6104c1610dfe565b6040516103529190614bbe565b6103016104dc366004614bff565b610ec3565b610301611207565b6104f16112ef565b6040516103529190614c7c565b6013546011546040805163ffffffff80851682526401000000009094049093166020840152820152606001610352565b610301611351565b6104c1611361565b60005461010090046001600160a01b0316610416565b61055c6113c1565b60405161035291908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b61030161059d366004614c8f565b611462565b6103016102e9366004614cde565b6104a06105be366004614d19565b60176020526000908152604090205481565b604080516001815260006020820181905291810191909152606001610352565b600c546001600160a01b0316610416565b61030161060f366004614d7e565b611595565b610416610622366004614850565b611b68565b610301610635366004614e63565b611c56565b600d546001600160a01b0316610416565b610416610659366004614850565b6001600160a01b039081166000908152600360205260409020541690565b6103016106853660046150a2565b611c62565b610301610698366004614850565b611c70565b6107596040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a081018252600e546001600160a01b038116825263ffffffff74010000000000000000000000000000000000000000820416602083015267ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811692820192909252600f548083166060830152680100000000000000009004909116608082015290565b6040805182516001600160a01b0316815260208084015163ffffffff16908201528282015167ffffffffffffffff908116928201929092526060808401518316908201526080928301519091169181019190915260a001610352565b6104a0611ccf565b6103016107cb3660046150e9565b611d82565b6103016107de366004614850565b611f21565b6103016107f136600461518a565b611f58565b6103016108043660046151d0565b611f63565b610301610817366004614bff565b612858565b61030161082a366004614850565b612a30565b333014610868576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a08101515115610897576108808160a00151612a41565b61089781608001518260a001518360400151612bd7565b6108a081612c75565b50565b6108ab612d4d565b6001600160a01b0381166108eb576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa1580156109bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e1919061529d565b15905090565b6109ef612d4d565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b610a26612d4d565b610a2e612dac565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015610a6d57506005546001600160a01b03163314155b15610aa4576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610ae0576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610b3a576006600060078381548110610b0557610b056152ba565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610b3381615318565b9050610ae6565b5060005b82811015610c6a576000858281518110610b5a57610b5a6152ba565b6020026020010151905060006001600160a01b0316816001600160a01b031603610bb0576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610bc257610bc26152ba565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610c2757610c276152ba565b6020026020010151604051610c519291906001600160a01b03929092168252602082015260400190565b60405180910390a150610c6381615318565b9050610b3e565b508351610c7e90600790602087019061440d565b5050505050565b80516060908067ffffffffffffffff811115610ca357610ca36144af565b604051908082528060200260200182016040528015610ccc578160200160208202803683370190505b50915060005b81811015610d465760066000858381518110610cf057610cf06152ba565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610d2b57610d2b6152ba565b6020908102919091010152610d3f81615318565b9050610cd2565b5050919050565b6000808260800151518360a001515160206014610d6a9190615350565b610d749190615368565b610d7f906086615350565b610d899190615350565b90506000610d98601083615368565b9050610a28611bbc8560a00151516001610db29190615350565b610dbe90618aac615368565b6156b8610dcb8986615350565b610dd59190615350565b610ddf9190615350565b610de99190615350565b610df39190615350565b925050505b92915050565b60045460609067ffffffffffffffff811115610e1c57610e1c6144af565b604051908082528060200260200182016040528015610e45578160200160208202803683370190505b50905060005b600454811015610ebf57610e8560048281548110610e6b57610e6b6152ba565b6000918252602090912001546001600160a01b0316611b68565b828281518110610e9757610e976152ba565b6001600160a01b0390921660209283029190910190910152610eb881615318565b9050610e4b565b5090565b610ecb612d4d565b6004546000819003610f09576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610f97576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614610fe6576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610ff56001856153a5565b81548110611005576110056152ba565b9060005260206000200160009054906101000a90046001600160a01b03169050600482602001516bffffffffffffffffffffffff168154811061104a5761104a6152ba565b6000918252602090912001546001600160a01b0316600461106c6001866153a5565b8154811061107c5761107c6152ba565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483602001516bffffffffffffffffffffffff16815481106110d0576110d06152ba565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000291909216179055600480548061115a5761115a6153bc565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9084018101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03878116808452600383526040808520949094558351908152908716918101919091527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146112665760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060601680548060200260200160405190810160405280929190818152602001828054801561134757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611329575b5050505050905090565b611359612d4d565b610a2e612e48565b60606004805480602002602001604051908101604052809291908181526020018280548015611347576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611329575050505050905090565b6113ec6040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b5460608201819052429060009061142690836153a5565b602084015184519192506114529161143e9084615368565b856040015161144d9190615350565b612ed0565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561149f57506005546001600160a01b03163314155b156114d6576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff1161152a576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115346008612ee6565b602081015160098190558151600855600a546115509190612ed0565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916115eb91849163ffffffff851691908e908e9081908401838280828437600092019190915250612f9392505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260115480825260125460ff808216602085015261010090910416928201929092529083146116a65760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d617463680000000000000000000000604482015260640161125d565b6116b48b8b8b8b8b8b612fb6565b60007f000000000000000000000000000000000000000000000000000000000000000015611711576002826020015183604001516116f291906153eb565b6116fc919061543f565b6117079060016153eb565b60ff169050611727565b60208201516117219060016153eb565b60ff1690505b8881146117765760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604482015260640161125d565b8887146117c55760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604482015260640161125d565b3360009081526014602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115611808576118086148ab565b6002811115611819576118196148ab565b9052509050600281602001516002811115611836576118366148ab565b14801561187057506016816000015160ff1681548110611858576118586152ba565b6000918252602090912001546001600160a01b031633145b6118bc5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604482015260640161125d565b5050505050600088886040516118d3929190615461565b6040519081900381206118ea918c90602001615471565b60405160208183030381529060405280519060200120905061190a61447b565b604080518082019091526000808252602082015260005b88811015611b46576000600185888460208110611940576119406152ba565b61194d91901a601b6153eb565b8d8d8681811061195f5761195f6152ba565b905060200201358c8c87818110611978576119786152ba565b90506020020135604051600081526020016040526040516119b5949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156119d7573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526014602090815290849020838501909452835460ff80821685529296509294508401916101009004166002811115611a2c57611a2c6148ab565b6002811115611a3d57611a3d6148ab565b9052509250600183602001516002811115611a5a57611a5a6148ab565b14611aa75760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e0000604482015260640161125d565b8251849060ff16601f8110611abe57611abe6152ba565b602002015115611b105760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e6174757265000000000000000000000000604482015260640161125d565b600184846000015160ff16601f8110611b2b57611b2b6152ba565b9115156020909202015250611b3f81615318565b9050611921565b5050505063ffffffff8110611b5d57611b5d61548d565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611bbc576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611c2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c4f91906154bc565b9392505050565b6108a060008083612f93565b611c6c8282613046565b5050565b611c78612d4d565b600c805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b6000611d7d7f31a97e998befaf21ee85c8e3e1879003d67d46abb0b6b552882133e4c782986d600e54604080516020808201949094527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201526001600160a01b039092166080808401919091528151808403909101815260a09092019052805191012090565b905090565b611d8a612d4d565b80516001600160a01b0316611dcb576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600e805460208085018051604080880180516001600160a01b039098167fffffffffffffffff00000000000000000000000000000000000000000000000090961686177401000000000000000000000000000000000000000063ffffffff948516021777ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff998a16021790965560608089018051600f80546080808e018051948e167fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169290921768010000000000000000948e1694909402939093179091558451988952955190941695870195909552955187169085015291518516938301939093529151909216908201527fd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa09060a00161158a565b611f29612d4d565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6108a0816001613046565b855185518560ff16601f831115611fd6576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e65727300000000000000000000000000000000604482015260640161125d565b80600003612040576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f7369746976650000000000000000000000000000604482015260640161125d565b8183146120ce576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e00000000000000000000000000000000000000000000000000000000606482015260840161125d565b6120d9816003615368565b8311612141576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604482015260640161125d565b612149612d4d565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b6015541561230c576015546000906121a1906001906153a5565b90506000601582815481106121b8576121b86152ba565b6000918252602082200154601680546001600160a01b03909216935090849081106121e5576121e56152ba565b60009182526020808320909101546001600160a01b0385811684526014909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055601580549192509080612258576122586153bc565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908101805473ffffffffffffffffffffffffffffffffffffffff1916905501905560168054806122b6576122b66153bc565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908101805473ffffffffffffffffffffffffffffffffffffffff1916905501905550612187915050565b60005b8151518110156126d15760006014600084600001518481518110612335576123356152ba565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115612372576123726148ab565b146123d9576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e65722061646472657373000000000000000000604482015260640161125d565b6040805180820190915260ff8216815260016020820152825180516014916000918590811061240a5761240a6152ba565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115612480576124806148ab565b0217905550600091506124909050565b60146000846020015184815181106124aa576124aa6152ba565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156124e7576124e76148ab565b1461254e576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161125d565b6040805180820190915260ff821681526020810160028152506014600084602001518481518110612581576125816152ba565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156125f7576125f76148ab565b021790555050825180516015925083908110612615576126156152ba565b6020908102919091018101518254600181018455600093845292829020909201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039093169290921790915582015180516016919083908110612679576126796152ba565b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039092169190911790556126ca81615318565b905061230f565b5060408101516012805460ff191660ff909216919091179055601380547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff4381168202928317855590830481169360019390926000926127459286929082169116176154d9565b92506101000a81548163ffffffff021916908363ffffffff1602179055506127a44630601360009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a0015161376f565b6011819055825180516012805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560135460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612843988b98919763ffffffff909216969095919491939192615501565b60405180910390a15050505050505050505050565b612860612d4d565b6001600160a01b038216158061287d57506001600160a01b038116155b156128b4576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612943576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201805473ffffffffffffffffffffffffffffffffffffffff1916821790558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612a38612d4d565b6108a0816137fc565b6000805b8251811015612b4057600060066000858481518110612a6657612a666152ba565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600003612af957838281518110612aaf57612aaf6152ba565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260240161125d565b838281518110612b0b57612b0b6152ba565b60200260200101516020015181612b229190615368565b612b2c9084615350565b92505080612b3990615318565b9050612a45565b508015611c6c57612b516008612ee6565b600a54811115612b8d576040517f3bfa6f3800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060086002016000828254612ba291906153a5565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba3789060200161094e565b8151835114612c12576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015612c6f57612c5f848281518110612c3357612c336152ba565b6020026020010151848381518110612c4d57612c4d6152ba565b602002602001015160200151846138b8565b612c6881615318565b9050612c15565b50505050565b60408101516001600160a01b03163b612c8b5750565b600c546040517fda52b4c40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063da52b4c490612cd49084906004016155e4565b6020604051808303816000875af1158015612cf3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d17919061529d565b6108a0576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005461010090046001600160a01b03163314610a2e5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161125d565b60005460ff16612dfe5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161125d565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60005460ff1615612e9b5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161125d565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612e2b3390565b6000818310612edf5781611c4f565b5090919050565b6001810154600282015442911480612f015750808260030154145b15612f0a575050565b816001015482600201541115612f4c576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482612f5e91906153a5565b60018401548454919250612f8591612f769084615368565b856002015461144d9190615350565b600284015550600390910155565b612fb181806020019051810190612faa919061588c565b6000613046565b505050565b6000612fc3826020615368565b612fce856020615368565b612fda88610144615350565b612fe49190615350565b612fee9190615350565b612ff9906000615350565b905036811461303d576040517f8e1192e10000000000000000000000000000000000000000000000000000000081526004810182905236602482015260440161125d565b50505050505050565b60005460ff16156130995760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161125d565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156130ec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613110919061529d565b15613146576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c546001600160a01b0316613188576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608201515160008190036131c9576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff8111156131e4576131e46144af565b60405190808252806020026020018201604052801561326b57816020015b61325860408051610100810182526000808252602080830182905282840182905260608084018390526080840181905260a084015283518085019094528184528301529060c08201908152602001600081525090565b8152602001906001900390816132025790505b50905060008267ffffffffffffffff811115613289576132896144af565b6040519080825280602002602001820160405280156132b2578160200160208202803683370190505b50905060006133637f31a97e998befaf21ee85c8e3e1879003d67d46abb0b6b552882133e4c782986d600e54604080516020808201949094527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201526001600160a01b039092166080808401919091528151808403909101815260a09092019052805191012090565b905060005b848110156133fd57600087606001518281518110613388576133886152ba565b60200260200101518060200190518101906133a39190615a64565b90506133af818461392f565b8483815181106133c1576133c16152ba565b602002602001018181525050808583815181106133e0576133e06152ba565b602002602001018190525050806133f690615318565b9050613368565b5060008061341e8489608001518a60a001518b60c001518c60e00151613a17565b600e54919350915060009074010000000000000000000000000000000000000000900463ffffffff1661345184426153a5565b11905060005b87811015613763576000878281518110613473576134736152ba565b6020026020010151905060006134a6826020015167ffffffffffffffff1660009081526010602052604090205460ff1690565b905060028160038111156134bc576134bc6148ab565b036135055760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161125d565b8a158061350f5750835b8061352b57506003816003811115613529576135296148ab565b145b613561576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61356a82613b0f565b60008082600381111561357f5761357f6148ab565b14801561358a57508b155b156135e3576135a68a518761359f9190615b51565b8e85613c41565b905060006135bb8460c0015160000151613e2e565b90506135c88183306138b8565b818460c001516020018181516135de91906153a5565b905250505b60008260038111156135f7576135f76148ab565b146136335760208084015167ffffffffffffffff16600090815260178252604090205460c0850151909101805161362f9083906153a5565b9052505b60208381015167ffffffffffffffff166000908152601090915260408120805460ff1916600117905561366d61366885613e90565b614159565b60208086015167ffffffffffffffff1660009081526010909152604090208054919250829160ff191660018360038111156136aa576136aa6148ab565b021790555060008360038111156136c3576136c36148ab565b1480156136e1575060038160038111156136df576136df6148ab565b145b156137085760208085015167ffffffffffffffff1660009081526017909152604090208290555b836020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf8260405161374691906148da565b60405180910390a2505050508061375c90615318565b9050613457565b50505050505050505050565b6000808a8a8a8a8a8a8a8a8a60405160200161379399989796959493929190615b65565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b038216036138545760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161125d565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b15801561391b57600080fd5b505af115801561303d573d6000803e3d6000fd5b60008060001b828460200151856040015186606001518760800151805190602001208860a001516040516020016139669190615bed565b604051602081830303815290604052805190602001208960e001518a60c001516040516020016139f9999897969594939291909889526020808a019890985267ffffffffffffffff9690961660408901526001600160a01b039485166060890152928416608088015260a087019190915260c086015260e085015281511661010084015201516101208201526101400190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600d546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce90613a74908c908c908c908c908c90600401615c30565b6020604051808303816000875af1158015613a93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ab79190615c82565b905060008111613af3576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a613aff90846153a5565b9350935050509550959350505050565b80517f000000000000000000000000000000000000000000000000000000000000000014613b6f5780516040517fd44bc9eb000000000000000000000000000000000000000000000000000000008152600481019190915260240161125d565b600f5460a0820151516801000000000000000090910467ffffffffffffffff161015613bd95760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161125d565b600f5460808201515167ffffffffffffffff90911610156108a057600f546080820151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9092166004830152602482015260440161125d565b6000806000613c578460c0015160000151613e2e565b6001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613c94573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613cb891906154bc565b905060005b856020015151811015613d3157816001600160a01b031686602001518281518110613cea57613cea6152ba565b60200260200101516001600160a01b031603613d215785604001518181518110613d1657613d166152ba565b602002602001015192505b613d2a81615318565b9050613cbd565b5081613d74576040517fce480bcc0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161125d565b6000670de0b6b3a7640000833a8760e00151613d908b8a610d4d565b613d9a9190615350565b613da49190615368565b613dae9190615368565b613db89190615b51565b90508460c0015160200151811115613e245760208086015160c0870151909101516040517f394a2c2700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909216600483015260248201839052604482015260640161125d565b9695505050505050565b6001600160a01b038181166000908152600360205260409020541680613e8b576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161125d565b919050565b613ed96040518060e00160405280600081526020016060815260200160006001600160a01b03168152602001606081526020016060815260200160608152602001600081525090565b6000613eed8360a001518460c0015161428d565b805190915060008167ffffffffffffffff811115613f0d57613f0d6144af565b604051908082528060200260200182016040528015613f5257816020015b6040805180820190915260008082526020820152815260200190600190039081613f2b5790505b50905060008267ffffffffffffffff811115613f7057613f706144af565b604051908082528060200260200182016040528015613f99578160200160208202803683370190505b50905060005b838110156140cf576000613fcf868381518110613fbe57613fbe6152ba565b602002602001015160000151613e2e565b905080838381518110613fe457613fe46152ba565b60200260200101906001600160a01b031690816001600160a01b0316815250506040518060400160405280826001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561404d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061407191906154bc565b6001600160a01b03168152602001878481518110614091576140916152ba565b6020026020010151602001518152508483815181106140b2576140b26152ba565b602002602001018190525050806140c890615318565b9050613f9f565b506040518060e0016040528087600001518152602001876040015160405160200161410991906001600160a01b0391909116815260200190565b604051602081830303815290604052815260200187606001516001600160a01b03168152602001876080015181526020018281526020018381526020018760e00151815250945050505050919050565b6040517f07a22a0700000000000000000000000000000000000000000000000000000000815260009030906307a22a07906141989085906004016155e4565b600060405180830381600087803b1580156141b257600080fd5b505af19250505080156141c3575060015b614285573d8080156141f1576040519150601f19603f3d011682016040523d82523d6000602084013e6141f6565b606091505b5061420081615c9b565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da800000000000000000000000000000000000000000000000000000000036142515750600392915050565b806040517fcf19edfd00000000000000000000000000000000000000000000000000000000815260040161125d9190614973565b506002919050565b606060005b83518110156143215782600001516001600160a01b03168482815181106142bb576142bb6152ba565b6020026020010151600001516001600160a01b0316036143115782602001518482815181106142ec576142ec6152ba565b60200260200101516020018181516143049190615350565b905250839150610df89050565b61431a81615318565b9050614292565b506000835160016143329190615350565b67ffffffffffffffff81111561434a5761434a6144af565b60405190808252806020026020018201604052801561438f57816020015b60408051808201909152600080825260208201528152602001906001900390816143685790505b50905060005b84518110156143e6578481815181106143b0576143b06152ba565b60200260200101518282815181106143ca576143ca6152ba565b6020026020010181905250806143df90615318565b9050614395565b5082818551815181106143fb576143fb6152ba565b60209081029190910101529392505050565b82805482825590600052602060002090810192821561446f579160200282015b8281111561446f578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390911617825560209092019160019091019061442d565b50610ebf92915061449a565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610ebf576000815560010161449b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715614501576145016144af565b60405290565b60405160e0810167ffffffffffffffff81118282101715614501576145016144af565b604051610100810167ffffffffffffffff81118282101715614501576145016144af565b604051601f8201601f1916810167ffffffffffffffff81118282101715614577576145776144af565b604052919050565b600067ffffffffffffffff821115614599576145996144af565b50601f01601f191660200190565b600082601f8301126145b857600080fd5b81356145cb6145c68261457f565b61454e565b8181528460208386010111156145e057600080fd5b816020850160208301376000918101602001919091529392505050565b6001600160a01b03811681146108a057600080fd5b8035613e8b816145fd565b600067ffffffffffffffff821115614637576146376144af565b5060051b60200190565b600082601f83011261465257600080fd5b813560206146626145c68361461d565b82815260059290921b8401810191818101908684111561468157600080fd5b8286015b848110156146a5578035614698816145fd565b8352918301918301614685565b509695505050505050565b6000604082840312156146c257600080fd5b6146ca6144de565b905081356146d7816145fd565b808252506020820135602082015292915050565b600082601f8301126146fc57600080fd5b8135602061470c6145c68361461d565b82815260069290921b8401810191818101908684111561472b57600080fd5b8286015b848110156146a55761474188826146b0565b83529183019160400161472f565b60006020828403121561476157600080fd5b813567ffffffffffffffff8082111561477957600080fd5b9083019060e0828603121561478d57600080fd5b614795614507565b823581526020830135828111156147ab57600080fd5b6147b7878286016145a7565b6020830152506147c960408401614612565b60408201526060830135828111156147e057600080fd5b6147ec878286016145a7565b60608301525060808301358281111561480457600080fd5b61481087828601614641565b60808301525060a08301358281111561482857600080fd5b614834878286016146eb565b60a08301525060c083013560c082015280935050505092915050565b60006020828403121561486257600080fd5b8135611c4f816145fd565b67ffffffffffffffff811681146108a057600080fd5b8035613e8b8161486d565b6000602082840312156148a057600080fd5b8135611c4f8161486d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310614915577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b8381101561493657818101518382015260200161491e565b83811115612c6f5750506000910152565b6000815180845261495f81602086016020860161491b565b601f01601f19169290920160200192915050565b602081526000611c4f6020830184614947565b600082601f83011261499757600080fd5b813560206149a76145c68361461d565b82815260059290921b840181019181810190868411156149c657600080fd5b8286015b848110156146a557803583529183019183016149ca565b600080604083850312156149f457600080fd5b823567ffffffffffffffff80821115614a0c57600080fd5b614a1886838701614641565b93506020850135915080821115614a2e57600080fd5b50614a3b85828601614986565b9150509250929050565b600060208284031215614a5757600080fd5b813567ffffffffffffffff811115614a6e57600080fd5b614a7a84828501614641565b949350505050565b6020808252825182820181905260009190848201906040850190845b81811015614aba57835183529284019291840191600101614a9e565b50909695505050505050565b60008060408385031215614ad957600080fd5b82359150602083013567ffffffffffffffff80821115614af857600080fd5b908401906101208287031215614b0d57600080fd5b614b1561452a565b82358152614b2560208401614883565b6020820152614b3660408401614612565b6040820152614b4760608401614612565b6060820152608083013582811115614b5e57600080fd5b614b6a888286016145a7565b60808301525060a083013582811115614b8257600080fd5b614b8e888286016146eb565b60a083015250614ba18760c085016146b0565b60c082015261010083013560e08201528093505050509250929050565b6020808252825182820181905260009190848201906040850190845b81811015614aba5783516001600160a01b031683529284019291840191600101614bda565b60008060408385031215614c1257600080fd5b8235614c1d816145fd565b91506020830135614c2d816145fd565b809150509250929050565b600081518084526020808501945080840160005b83811015614c715781516001600160a01b031687529582019590820190600101614c4c565b509495945050505050565b602081526000611c4f6020830184614c38565b600060408284031215614ca157600080fd5b6040516040810181811067ffffffffffffffff82111715614cc457614cc46144af565b604052823581526020928301359281019290925250919050565b600060208284031215614cf057600080fd5b813567ffffffffffffffff811115614d0757600080fd5b820160e08185031215611c4f57600080fd5b600060208284031215614d2b57600080fd5b5035919050565b60008083601f840112614d4457600080fd5b50813567ffffffffffffffff811115614d5c57600080fd5b6020830191508360208260051b8501011115614d7757600080fd5b9250929050565b60008060008060008060008060e0898b031215614d9a57600080fd5b606089018a811115614dab57600080fd5b8998503567ffffffffffffffff80821115614dc557600080fd5b818b0191508b601f830112614dd957600080fd5b813581811115614de857600080fd5b8c6020828501011115614dfa57600080fd5b6020830199508098505060808b0135915080821115614e1857600080fd5b614e248c838d01614d32565b909750955060a08b0135915080821115614e3d57600080fd5b50614e4a8b828c01614d32565b999c989b50969995989497949560c00135949350505050565b600060208284031215614e7557600080fd5b813567ffffffffffffffff811115614e8c57600080fd5b614a7a848285016145a7565b600082601f830112614ea957600080fd5b81356020614eb96145c68361461d565b82815260059290921b84018101918181019086841115614ed857600080fd5b8286015b848110156146a5578035614eef8161486d565b8352918301918301614edc565b600082601f830112614f0d57600080fd5b81356020614f1d6145c68361461d565b82815260059290921b84018101918181019086841115614f3c57600080fd5b8286015b848110156146a557803567ffffffffffffffff811115614f605760008081fd5b614f6e8986838b01016145a7565b845250918301918301614f40565b60006101008284031215614f8f57600080fd5b614f9761452a565b9050813567ffffffffffffffff80821115614fb157600080fd5b614fbd85838601614e98565b83526020840135915080821115614fd357600080fd5b614fdf85838601614641565b60208401526040840135915080821115614ff857600080fd5b61500485838601614986565b6040840152606084013591508082111561501d57600080fd5b61502985838601614efc565b6060840152608084013591508082111561504257600080fd5b61504e85838601614986565b608084015260a084013560a084015260c084013591508082111561507157600080fd5b5061507e84828501614986565b60c08301525060e082013560e082015292915050565b80151581146108a057600080fd5b600080604083850312156150b557600080fd5b823567ffffffffffffffff8111156150cc57600080fd5b6150d885828601614f7c565b9250506020830135614c2d81615094565b600060a082840312156150fb57600080fd5b60405160a0810181811067ffffffffffffffff8211171561511e5761511e6144af565b604052823561512c816145fd565b8152602083013563ffffffff8116811461514557600080fd5b602082015260408301356151588161486d565b6040820152606083013561516b8161486d565b6060820152608083013561517e8161486d565b60808201529392505050565b60006020828403121561519c57600080fd5b813567ffffffffffffffff8111156151b357600080fd5b614a7a84828501614f7c565b803560ff81168114613e8b57600080fd5b60008060008060008060c087890312156151e957600080fd5b863567ffffffffffffffff8082111561520157600080fd5b61520d8a838b01614641565b9750602089013591508082111561522357600080fd5b61522f8a838b01614641565b965061523d60408a016151bf565b9550606089013591508082111561525357600080fd5b61525f8a838b016145a7565b945061526d60808a01614883565b935060a089013591508082111561528357600080fd5b5061529089828a016145a7565b9150509295509295509295565b6000602082840312156152af57600080fd5b8151611c4f81615094565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203615349576153496152e9565b5060010190565b60008219821115615363576153636152e9565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156153a0576153a06152e9565b500290565b6000828210156153b7576153b76152e9565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600060ff821660ff84168060ff03821115615408576154086152e9565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff83168061545257615452615410565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6000602082840312156154ce57600080fd5b8151611c4f816145fd565b600063ffffffff8083168185168083038211156154f8576154f86152e9565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526155318184018a614c38565b905082810360808401526155458189614c38565b905060ff871660a084015282810360c08401526155628187614947565b905067ffffffffffffffff851660e08401528281036101008401526155878185614947565b9c9b505050505050505050505050565b600081518084526020808501945080840160005b83811015614c71576155d187835180516001600160a01b03168252602090810151910152565b60409690960195908201906001016155ab565b60208152815160208201526000602083015160e0604084015261560b610100840182614947565b90506001600160a01b0360408501511660608401526060840151601f198085840301608086015261563c8383614947565b925060808601519150808584030160a08601526156598383614c38565b925060a08601519150808584030160c0860152506156778282615597565b91505060c084015160e08401528091505092915050565b8051613e8b8161486d565b600082601f8301126156aa57600080fd5b815160206156ba6145c68361461d565b82815260059290921b840181019181810190868411156156d957600080fd5b8286015b848110156146a55780516156f08161486d565b83529183019183016156dd565b8051613e8b816145fd565b600082601f83011261571957600080fd5b815160206157296145c68361461d565b82815260059290921b8401810191818101908684111561574857600080fd5b8286015b848110156146a557805161575f816145fd565b835291830191830161574c565b600082601f83011261577d57600080fd5b8151602061578d6145c68361461d565b82815260059290921b840181019181810190868411156157ac57600080fd5b8286015b848110156146a557805183529183019183016157b0565b600082601f8301126157d857600080fd5b81516157e66145c68261457f565b8181528460208386010111156157fb57600080fd5b614a7a82602083016020870161491b565b600082601f83011261581d57600080fd5b8151602061582d6145c68361461d565b82815260059290921b8401810191818101908684111561584c57600080fd5b8286015b848110156146a557805167ffffffffffffffff8111156158705760008081fd5b61587e8986838b01016157c7565b845250918301918301615850565b60006020828403121561589e57600080fd5b815167ffffffffffffffff808211156158b657600080fd5b9083019061010082860312156158cb57600080fd5b6158d361452a565b8251828111156158e257600080fd5b6158ee87828601615699565b82525060208301518281111561590357600080fd5b61590f87828601615708565b60208301525060408301518281111561592757600080fd5b6159338782860161576c565b60408301525060608301518281111561594b57600080fd5b6159578782860161580c565b60608301525060808301518281111561596f57600080fd5b61597b8782860161576c565b60808301525060a083015160a082015260c08301518281111561599d57600080fd5b6159a98782860161576c565b60c08301525060e083015160e082015280935050505092915050565b6000604082840312156159d757600080fd5b6159df6144de565b905081516159ec816145fd565b808252506020820151602082015292915050565b600082601f830112615a1157600080fd5b81516020615a216145c68361461d565b82815260069290921b84018101918181019086841115615a4057600080fd5b8286015b848110156146a557615a5688826159c5565b835291830191604001615a44565b600060208284031215615a7657600080fd5b815167ffffffffffffffff80821115615a8e57600080fd5b908301906101208286031215615aa357600080fd5b615aab61452a565b82518152615abb6020840161568e565b6020820152615acc604084016156fd565b6040820152615add606084016156fd565b6060820152608083015182811115615af457600080fd5b615b00878286016157c7565b60808301525060a083015182811115615b1857600080fd5b615b2487828601615a00565b60a083015250615b378660c085016159c5565b60c0820152610100929092015160e0830152509392505050565b600082615b6057615b60615410565b500490565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b166040850152816060850152615b9f8285018b614c38565b91508382036080850152615bb3828a614c38565b915060ff881660a085015283820360c0850152615bd08288614947565b90861660e085015283810361010085015290506155878185614947565b602081526000611c4f6020830184615597565b600081518084526020808501945080840160005b83811015614c7157815187529582019590820190600101615c14565b60a081526000615c4360a0830188615c00565b8281036020840152615c558188615c00565b90508560408401528281036060840152615c6f8186615c00565b9150508260808301529695505050505050565b600060208284031215615c9457600080fd5b5051919050565b6000815160208301517fffffffff0000000000000000000000000000000000000000000000000000000080821693506004831015615ce35780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
}

var EVM2EVMTollOffRampHelperABI = EVM2EVMTollOffRampHelperMetaData.ABI

var EVM2EVMTollOffRampHelperBin = EVM2EVMTollOffRampHelperMetaData.Bin

func DeployEVM2EVMTollOffRampHelper(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, commitStore common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address) (common.Address, *types.Transaction, *EVM2EVMTollOffRampHelper, error) {
	parsed, err := EVM2EVMTollOffRampHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMTollOffRampHelperBin), backend, sourceChainId, chainId, offRampConfig, commitStore, afn, sourceTokens, pools, rateLimiterConfig, tokenLimitsAdmin)
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
	return common.HexToHash("0xd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa0")
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
