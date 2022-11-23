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
	FeeUpdates               []CCIPFeeUpdate
	EncodedMessages          [][]byte
	InnerProofs              [][32]byte
	InnerProofFlagBits       *big.Int
	OuterProofs              [][32]byte
	OuterProofFlagBits       *big.Int
}

type CCIPFeeUpdate struct {
	ChainId  *big.Int
	GasPrice *big.Int
}

var EVM2EVMFreeOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitStore\",\"outputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"}],\"name\":\"setCommitStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162005f8538038062005f8583398101604081905262000035916200073f565b6000805460ff191681556001908b908b908a908a908a908a908a908a908a90829082908690869089903390819081620000b55760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ef57620000ef8162000387565b5050506001600160a01b0381166200011a57604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015c5760405162d8548360e71b815260040160405180910390fd5b81516200017190600490602085019062000438565b5060005b82518110156200023c57600082828151811062000196576200019662000834565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001e057620001e062000834565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b02929091169190911790555062000234816200084a565b905062000175565b5050600580546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600855600993909355600a55600b919091558716620002b9576040516342bcdf7f60e11b815260040160405180910390fd5b5050506080959095525060a0929092526001600160a01b0390811660c052600d80546001600160a01b0319169190921617905550151560e0525050855160168054602089015160408a01516060909a015163ffffffff9094166001600160601b0319909216919091176401000000006001600160401b039283160217600160601b600160e01b0319166c0100000000000000000000000099821699909902600160a01b600160e01b03191698909817600160a01b989092169790970217909555506200087295505050505050565b336001600160a01b03821603620003e15760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ac565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b82805482825590600052602060002090810192821562000490579160200282015b828111156200049057825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000459565b506200049e929150620004a2565b5090565b5b808211156200049e5760008155600101620004a3565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620004fa57620004fa620004b9565b604052919050565b80516001600160401b03811681146200051a57600080fd5b919050565b6000608082840312156200053257600080fd5b604051608081016001600160401b0381118282101715620005575762000557620004b9565b8060405250809150825163ffffffff811681146200057457600080fd5b8152620005846020840162000502565b6020820152620005976040840162000502565b6040820152620005aa6060840162000502565b60608201525092915050565b6001600160a01b0381168114620005cc57600080fd5b50565b80516200051a81620005b6565b60006001600160401b03821115620005f857620005f8620004b9565b5060051b60200190565b600082601f8301126200061457600080fd5b815160206200062d6200062783620005dc565b620004cf565b82815260059290921b840181019181810190868411156200064d57600080fd5b8286015b84811015620006755780516200066781620005b6565b835291830191830162000651565b509695505050505050565b600082601f8301126200069257600080fd5b81516020620006a56200062783620005dc565b82815260059290921b84018101918181019086841115620006c557600080fd5b8286015b8481101562000675578051620006df81620005b6565b8352918301918301620006c9565b6000604082840312156200070057600080fd5b604080519081016001600160401b0381118282101715620007255762000725620004b9565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000806101c08b8d0312156200076057600080fd5b8a51995060208b015198506200077a8c60408d016200051f565b97506200078a60c08c01620005cf565b96506200079a60e08c01620005cf565b9550620007ab6101008c01620005cf565b6101208c01519095506001600160401b0380821115620007ca57600080fd5b620007d88e838f0162000602565b95506101408d0151915080821115620007f057600080fd5b50620007ff8d828e0162000680565b935050620008128c6101608d01620006ed565b9150620008236101a08c01620005cf565b90509295989b9194979a5092959850565b634e487b7160e01b600052603260045260246000fd5b6000600182016200086b57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e0516156c4620008c1600039600061155a015260006130000152600081816102f00152612fda0152600081816102cd01528181612fb5015261385301526156c46000f3fe608060405234801561001057600080fd5b50600436106102ad5760003560e01c80638456cb591161017b578063b4069b31116100d8578063c3f909d41161008c578063e3d0e71211610071578063e3d0e7121461075c578063eb511dd41461076f578063f2fde38b1461078257600080fd5b8063c3f909d41461066c578063c90332841461074957600080fd5b8063bbe4f6db116100bd578063bbe4f6db1461061a578063bc29705a14610646578063c0d786551461065957600080fd5b8063b4069b31146105f6578063b66f0efb1461060957600080fd5b80639438ff631161012f578063afcb95d711610114578063afcb95d7146105b2578063b0f479a1146105d2578063b1dc65a4146105e357600080fd5b80639438ff6314610591578063a8e913211461059f57600080fd5b80638da5cb5b116101605780638da5cb5b1461052d57806390c2339b14610543578063918725431461057e57600080fd5b80638456cb591461051d57806389c065681461052557600080fd5b80633f4ba83a11610229578063681fba16116101dd57806379ba5097116101c257806379ba5097146104d057806381411834146104d857806381ff7048146104ed57600080fd5b8063681fba16146104a8578063744b92e2146104bd57600080fd5b80634741062e1161020e5780634741062e1461046c578063599f64311461048c5780635c975abb1461049d57600080fd5b80633f4ba83a146104515780634352fa9f1461045957600080fd5b8063147809b3116102805780632222dd42116102655780632222dd42146103cd5780632d0335ab146103f257806339aa92641461043e57600080fd5b8063147809b31461036c578063181f5a771461038457600080fd5b806307a22a07146102b2578063087ae6df146102c7578063108ee5fc1461031f578063142a98fc14610332575b600080fd5b6102c56102c036600461415e565b610795565b005b604080517f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006020820152015b60405180910390f35b6102c561032d36600461425f565b610809565b61035f61034036600461429d565b67ffffffffffffffff166000908152600e602052604090205460ff1690565b60405161031691906142e9565b6103746108c0565b6040519015158152602001610316565b6103c06040518060400160405280602081526020017f45564d3245564d537562736372697074696f6e4f666652616d7020312e302e3081525081565b6040516103169190614382565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610316565b61042561040036600461425f565b6001600160a01b031660009081526015602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff9091168152602001610316565b6102c561044c36600461425f565b61094d565b6102c5610984565b6102c5610467366004614454565b610996565b61047f61047a3660046144b8565b610beb565b60405161031691906144f5565b6005546001600160a01b03166103da565b60005460ff16610374565b6104b0610cb3565b6040516103169190614539565b6102c56104cb36600461457a565b610d78565b6102c561109e565b6104e0611186565b60405161031691906145f7565b601154600f546040805163ffffffff80851682526401000000009094049093166020840152820152606001610316565b6102c56111e8565b6104b06111f8565b60005461010090046001600160a01b03166103da565b61054b611258565b60405161031691908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6102c561058c36600461463a565b6112f9565b6102c56102ad366004614656565b6102c56105ad3660046147d9565b61142c565b604080516001815260006020820181905291810191909152606001610316565b600c546001600160a01b03166103da565b6102c56105f1366004614982565b611437565b6103da61060436600461425f565b611a0a565b600d546001600160a01b03166103da565b6103da61062836600461425f565b6001600160a01b039081166000908152600360205260409020541690565b6102c5610654366004614a67565b611af8565b6102c561066736600461425f565b611c2e565b6106fe604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260165463ffffffff8116825267ffffffffffffffff6401000000008204811660208401526c010000000000000000000000008204811693830193909352740100000000000000000000000000000000000000009004909116606082015290565b6040516103169190815163ffffffff16815260208083015167ffffffffffffffff90811691830191909152604080840151821690830152606092830151169181019190915260800190565b6102c561075736600461425f565b611c8d565b6102c561076a366004614b06565b611cc4565b6102c561077d36600461457a565b61257d565b6102c561079036600461425f565b612755565b3330146107ce576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a081015151156107fd576107e68160a00151612766565b6107fd81608001518260a00151836040015161296e565b61080681612a0c565b50565b610811612ae4565b6001600160a01b038116610851576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa158015610923573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109479190614bd3565b15905090565b610955612ae4565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b61098c612ae4565b610994612b43565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b0316141580156109d357506005546001600160a01b03163314155b15610a0a576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610a46576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610aa0576006600060078381548110610a6b57610a6b614bf5565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610a9981614c53565b9050610a4c565b5060005b82811015610bd0576000858281518110610ac057610ac0614bf5565b6020026020010151905060006001600160a01b0316816001600160a01b031603610b16576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610b2857610b28614bf5565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610b8d57610b8d614bf5565b6020026020010151604051610bb79291906001600160a01b03929092168252602082015260400190565b60405180910390a150610bc981614c53565b9050610aa4565b508351610be4906007906020870190613e0f565b5050505050565b80516060908067ffffffffffffffff811115610c0957610c09613eb1565b604051908082528060200260200182016040528015610c32578160200160208202803683370190505b50915060005b81811015610cac5760066000858381518110610c5657610c56614bf5565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610c9157610c91614bf5565b6020908102919091010152610ca581614c53565b9050610c38565b5050919050565b60045460609067ffffffffffffffff811115610cd157610cd1613eb1565b604051908082528060200260200182016040528015610cfa578160200160208202803683370190505b50905060005b600454811015610d7457610d3a60048281548110610d2057610d20614bf5565b6000918252602090912001546001600160a01b0316611a0a565b828281518110610d4c57610d4c614bf5565b6001600160a01b0390921660209283029190910190910152610d6d81614c53565b9050610d00565b5090565b610d80612ae4565b6004546000819003610dbe576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610e4c576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614610e9b576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610eaa600185614c6d565b81548110610eba57610eba614bf5565b9060005260206000200160009054906101000a90046001600160a01b03169050600482602001516bffffffffffffffffffffffff1681548110610eff57610eff614bf5565b6000918252602090912001546001600160a01b03166004610f21600186614c6d565b81548110610f3157610f31614bf5565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483602001516bffffffffffffffffffffffff1681548110610f8557610f85614bf5565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000291909216179055600480548061100f5761100f614c84565b600082815260208082206000199084018101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03878116808452600383526040808520949094558351908152908716918101919091527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146110fd5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b606060148054806020026020016040519081016040528092919081815260200182805480156111de57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116111c0575b5050505050905090565b6111f0612ae4565b610994612bdf565b606060048054806020026020016040519081016040528092919081815260200182805480156111de576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116111c0575050505050905090565b6112836040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b546060820181905242906000906112bd9083614c6d565b602084015184519192506112e9916112d59084614cb3565b85604001516112e49190614cd2565b612c67565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561133657506005546001600160a01b03163314155b1561136d576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff116113c1576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113cb6008612c7d565b602081015160098190558151600855600a546113e79190612c67565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b610806816001612d2a565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161148d91849163ffffffff851691908e908e908190840183828082843760009201919091525061340792505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600f5480825260105460ff808216602085015261010090910416928201929092529083146115485760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d61746368000000000000000000000060448201526064016110f4565b6115568b8b8b8b8b8b61342a565b60007f0000000000000000000000000000000000000000000000000000000000000000156115b3576002826020015183604001516115949190614cea565b61159e9190614d3e565b6115a9906001614cea565b60ff1690506115c9565b60208201516115c3906001614cea565b60ff1690505b8881146116185760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e61747572657300000000000060448201526064016110f4565b8887146116675760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e000060448201526064016110f4565b3360009081526012602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156116aa576116aa6142ba565b60028111156116bb576116bb6142ba565b90525090506002816020015160028111156116d8576116d86142ba565b14801561171257506014816000015160ff16815481106116fa576116fa614bf5565b6000918252602090912001546001600160a01b031633145b61175e5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d6974746572000000000000000060448201526064016110f4565b505050505060008888604051611775929190614d60565b60405190819003812061178c918c90602001614d70565b6040516020818303038152906040528051906020012090506117ac613e7d565b604080518082019091526000808252602082015260005b888110156119e85760006001858884602081106117e2576117e2614bf5565b6117ef91901a601b614cea565b8d8d8681811061180157611801614bf5565b905060200201358c8c8781811061181a5761181a614bf5565b9050602002013560405160008152602001604052604051611857949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611879573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526012602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156118ce576118ce6142ba565b60028111156118df576118df6142ba565b90525092506001836020015160028111156118fc576118fc6142ba565b146119495760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e000060448201526064016110f4565b8251849060ff16601f811061196057611960614bf5565b6020020151156119b25760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e617475726500000000000000000000000060448201526064016110f4565b600184846000015160ff16601f81106119cd576119cd614bf5565b91151560209092020152506119e181614c53565b90506117c3565b5050505063ffffffff81106119ff576119ff614d8c565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611a5e576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611acd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611af19190614dbb565b9392505050565b611b00612ae4565b80516016805460208085018051604080880180516060808b01805167ffffffffffffffff90811674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff9482166c0100000000000000000000000002949094167fffffffff00000000000000000000000000000000ffffffffffffffffffffffff978216640100000000027fffffffffffffffffffffffffffffffffffffffff000000000000000000000000909b1663ffffffff909d169c8d179a909a17969096169890981791909117909755815197885292518216938701939093529051811691850191909152905116908201527f1a9ff9caccd597226ff8c393ab44e6b57669905d9b16f1cdb2ac267253dbf27c90608001611421565b611c36612ae4565b600c805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b611c95612ae4565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b855185518560ff16601f831115611d37576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064016110f4565b80600003611da1576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016110f4565b818314611e2f576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084016110f4565b611e3a816003614cb3565b8311611ea2576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016110f4565b611eaa612ae4565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b6013541561203157601354600090611f0290600190614c6d565b9050600060138281548110611f1957611f19614bf5565b6000918252602082200154601480546001600160a01b0390921693509084908110611f4657611f46614bf5565b60009182526020808320909101546001600160a01b0385811684526012909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090811690915592909116808452922080549091169055601380549192509080611fb957611fb9614c84565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff191690550190556014805480611ff957611ff9614c84565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff1916905501905550611ee8915050565b60005b8151518110156123f6576000601260008460000151848151811061205a5761205a614bf5565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff166002811115612097576120976142ba565b146120fe576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016110f4565b6040805180820190915260ff8216815260016020820152825180516012916000918590811061212f5761212f614bf5565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156121a5576121a56142ba565b0217905550600091506121b59050565b60126000846020015184815181106121cf576121cf614bf5565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff16600281111561220c5761220c6142ba565b14612273576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016110f4565b6040805180820190915260ff8216815260208101600281525060126000846020015184815181106122a6576122a6614bf5565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561231c5761231c6142ba565b02179055505082518051601392508390811061233a5761233a614bf5565b6020908102919091018101518254600181018455600093845292829020909201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909316929092179091558201518051601491908390811061239e5761239e614bf5565b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039092169190911790556123ef81614c53565b9050612034565b5060408101516010805460ff191660ff909216919091179055601180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff43811682029283178555908304811693600193909260009261246a928692908216911617614dd8565b92506101000a81548163ffffffff021916908363ffffffff1602179055506124c94630601160009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a001516134ba565b600f819055825180516010805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560115460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612568988b98919763ffffffff909216969095919491939192614e00565b60405180910390a15050505050505050505050565b612585612ae4565b6001600160a01b03821615806125a257506001600160a01b038116155b156125d9576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612668576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201805473ffffffffffffffffffffffffffffffffffffffff1916821790558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b61275d612ae4565b61080681613547565b6000805b82518110156128655760006006600085848151811061278b5761278b614bf5565b6020026020010151600001516001600160a01b03166001600160a01b031681526020019081526020016000205490508060000361281e578382815181106127d4576127d4614bf5565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201526024016110f4565b83828151811061283057612830614bf5565b602002602001015160200151816128479190614cb3565b6128519084614cd2565b9250508061285e90614c53565b905061276a565b50801561296a576128766008612c7d565b6009548111156128c0576009546040517f688ccf770000000000000000000000000000000000000000000000000000000081526004810191909152602481018290526044016110f4565b600a5481111561292057600854600a54600091906128de9084614c6d565b6128e89190614e96565b9050806040517fe31e0f320000000000000000000000000000000000000000000000000000000081526004016110f491815260200190565b80600860020160008282546129359190614c6d565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378906020016108b4565b5050565b81518351146129a9576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015612a06576129f68482815181106129ca576129ca614bf5565b60200260200101518483815181106129e4576129e4614bf5565b60200260200101516020015184613603565b6129ff81614c53565b90506129ac565b50505050565b60408101516001600160a01b03163b612a225750565b600c546040517fda52b4c40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063da52b4c490612a6b908490600401614eee565b6020604051808303816000875af1158015612a8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612aae9190614bd3565b610806576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005461010090046001600160a01b031633146109945760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016110f4565b60005460ff16612b955760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016110f4565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60005460ff1615612c325760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016110f4565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612bc23390565b6000818310612c765781611af1565b5090919050565b6001810154600282015442911480612c985750808260030154145b15612ca1575050565b816001015482600201541115612ce3576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482612cf59190614c6d565b60018401548454919250612d1c91612d0d9084614cb3565b85600201546112e49190614cd2565b600284015550600390910155565b60005460ff1615612d7d5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016110f4565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612dd0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612df49190614bd3565b15612e2a576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c546001600160a01b0316612e6c576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6080820151516000819003612ead576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115612ec857612ec8613eb1565b604051908082528060200260200182016040528015612f3657816020015b60408051610100810182526000808252602080830182905292820181905260608083018290526080830182905260a0830181905260c083015260e08201528252600019909201910181612ee65790505b50905060008267ffffffffffffffff811115612f5457612f54613eb1565b604051908082528060200260200182016040528015612f7d578160200160208202803683370190505b509050600061304b7f3997e2cfd3ccacf768662bd35c3dbf323724407d75aae3019c04f4aa59b1193f604080516020808201939093527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166080808301919091528251808303909101815260a0909101909152805191012090565b905060005b848110156130e55760008760800151828151811061307057613070614bf5565b602002602001015180602001905181019061308b919061507b565b9050613097818461367a565b8483815181106130a9576130a9614bf5565b602002602001018181525050808583815181106130c8576130c8614bf5565b602002602001018190525050806130de90614c53565b9050613050565b506000613106838860a001518960c001518a60e001518b6101000151613757565b509050858015613126575060165463ffffffff166131248242614c6d565b105b1561315d576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b858110156133fd57600085828151811061317c5761317c614bf5565b6020026020010151905060006131af826020015167ffffffffffffffff166000908152600e602052604090205460ff1690565b905060028160038111156131c5576131c56142ba565b0361320e5760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016110f4565b608082015160608301516001600160a01b0316600090815260156020526040812054909167ffffffffffffffff9081169161324b91166001615168565b67ffffffffffffffff16149050808061327557506003826003811115613273576132736142ba565b145b6132bd5760808301516040517fb0241f4a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016110f4565b6132c68361384f565b60208381015167ffffffffffffffff166000908152600e90915260408120805460ff191660011790556133006132fb856139ab565b613c79565b60208086015167ffffffffffffffff166000908152600e909152604090208054919250829160ff1916600183600381111561333d5761333d6142ba565b0217905550836020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf8260405161338091906142e9565b60405180910390a281156133e85760608401516001600160a01b03166000908152601560205260408120805467ffffffffffffffff16916133c08361518b565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b50505050806133f690614c53565b9050613160565b5050505050505050565b6134258180602001905181019061341e91906153d4565b6000612d2a565b505050565b6000613437826020614cb3565b613442856020614cb3565b61344e88610144614cd2565b6134589190614cd2565b6134629190614cd2565b61346d906000614cd2565b90503681146134b1576040517f8e1192e1000000000000000000000000000000000000000000000000000000008152600481018290523660248201526044016110f4565b50505050505050565b6000808a8a8a8a8a8a8a8a8a6040516020016134de99989796959493929190615531565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b0382160361359f5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016110f4565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b15801561366657600080fd5b505af11580156134b1573d6000803e3d6000fd5b60008060001b828460200151856040015186606001518760a00151805190602001208860c001516040516020016136b191906155b9565b604051602081830303815290604052805190602001208960e001518a6080015160405160200161373999989796959493929190988952602089019790975267ffffffffffffffff95861660408901526001600160a01b03948516606089015292909316608087015260a086015260c085019190915260e0840152166101008201526101200190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600d546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce906137b4908c908c908c908c908c906004016155fc565b6020604051808303816000875af11580156137d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137f7919061564e565b905060008111613833576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a61383f9084614c6d565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146138af5780516040517fd44bc9eb00000000000000000000000000000000000000000000000000000000815260048101919091526024016110f4565b60165460c0820151517401000000000000000000000000000000000000000090910467ffffffffffffffff1610156139255760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016110f4565b60165460a0820151516c0100000000000000000000000090910467ffffffffffffffff1610156108065760165460a0820151516040517f869337890000000000000000000000000000000000000000000000000000000081526c0100000000000000000000000090920467ffffffffffffffff16600483015260248201526044016110f4565b6139f46040518060e00160405280600081526020016060815260200160006001600160a01b03168152602001606081526020016060815260200160608152602001600081525090565b60c08201515160008167ffffffffffffffff811115613a1557613a15613eb1565b604051908082528060200260200182016040528015613a5a57816020015b6040805180820190915260008082526020820152815260200190600190039081613a335790505b50905060008267ffffffffffffffff811115613a7857613a78613eb1565b604051908082528060200260200182016040528015613aa1578160200160208202803683370190505b50905060005b83811015613bf0576000613adb8760c001518381518110613aca57613aca614bf5565b602002602001015160000151613dad565b905080838381518110613af057613af0614bf5565b60200260200101906001600160a01b031690816001600160a01b031681525050806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613b4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b729190614dbb565b848381518110613b8457613b84614bf5565b60209081029190910101516001600160a01b03909116905260c0870151805183908110613bb357613bb3614bf5565b602002602001015160200151848381518110613bd157613bd1614bf5565b602090810291909101810151015250613be981614c53565b9050613aa7565b506040518060e00160405280866000015181526020018660400151604051602001613c2a91906001600160a01b0391909116815260200190565b604051602081830303815290604052815260200186606001516001600160a01b031681526020018660a0015181526020018281526020018381526020018660e001518152509350505050919050565b6040517f07a22a0700000000000000000000000000000000000000000000000000000000815260009030906307a22a0790613cb8908590600401614eee565b600060405180830381600087803b158015613cd257600080fd5b505af1925050508015613ce3575060015b613da5573d808015613d11576040519150601f19603f3d011682016040523d82523d6000602084013e613d16565b606091505b50613d2081615667565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003613d715750600392915050565b806040517fcf19edfd0000000000000000000000000000000000000000000000000000000081526004016110f49190614382565b506002919050565b6001600160a01b038181166000908152600360205260409020541680613e0a576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016110f4565b919050565b828054828255906000526020600020908101928215613e71579160200282015b82811115613e71578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190613e2f565b50610d74929150613e9c565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610d745760008155600101613e9d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613f0357613f03613eb1565b60405290565b60405160e0810167ffffffffffffffff81118282101715613f0357613f03613eb1565b604051610120810167ffffffffffffffff81118282101715613f0357613f03613eb1565b604051610100810167ffffffffffffffff81118282101715613f0357613f03613eb1565b604051601f8201601f1916810167ffffffffffffffff81118282101715613f9d57613f9d613eb1565b604052919050565b600067ffffffffffffffff821115613fbf57613fbf613eb1565b50601f01601f191660200190565b600082601f830112613fde57600080fd5b8135613ff1613fec82613fa5565b613f74565b81815284602083860101111561400657600080fd5b816020850160208301376000918101602001919091529392505050565b6001600160a01b038116811461080657600080fd5b8035613e0a81614023565b600067ffffffffffffffff82111561405d5761405d613eb1565b5060051b60200190565b600082601f83011261407857600080fd5b81356020614088613fec83614043565b82815260059290921b840181019181810190868411156140a757600080fd5b8286015b848110156140cb5780356140be81614023565b83529183019183016140ab565b509695505050505050565b600082601f8301126140e757600080fd5b813560206140f7613fec83614043565b82815260069290921b8401810191818101908684111561411657600080fd5b8286015b848110156140cb57604081890312156141335760008081fd5b61413b613ee0565b813561414681614023565b8152818501358582015283529183019160400161411a565b60006020828403121561417057600080fd5b813567ffffffffffffffff8082111561418857600080fd5b9083019060e0828603121561419c57600080fd5b6141a4613f09565b823581526020830135828111156141ba57600080fd5b6141c687828601613fcd565b6020830152506141d860408401614038565b60408201526060830135828111156141ef57600080fd5b6141fb87828601613fcd565b60608301525060808301358281111561421357600080fd5b61421f87828601614067565b60808301525060a08301358281111561423757600080fd5b614243878286016140d6565b60a08301525060c083013560c082015280935050505092915050565b60006020828403121561427157600080fd5b8135611af181614023565b67ffffffffffffffff8116811461080657600080fd5b8035613e0a8161427c565b6000602082840312156142af57600080fd5b8135611af18161427c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160048310614324577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b8381101561434557818101518382015260200161432d565b83811115612a065750506000910152565b6000815180845261436e81602086016020860161432a565b601f01601f19169290920160200192915050565b602081526000611af16020830184614356565b600082601f8301126143a657600080fd5b813560206143b6613fec83614043565b82815260059290921b840181019181810190868411156143d557600080fd5b8286015b848110156140cb5780356143ec81614023565b83529183019183016143d9565b600082601f83011261440a57600080fd5b8135602061441a613fec83614043565b82815260059290921b8401810191818101908684111561443957600080fd5b8286015b848110156140cb578035835291830191830161443d565b6000806040838503121561446757600080fd5b823567ffffffffffffffff8082111561447f57600080fd5b61448b86838701614395565b935060208501359150808211156144a157600080fd5b506144ae858286016143f9565b9150509250929050565b6000602082840312156144ca57600080fd5b813567ffffffffffffffff8111156144e157600080fd5b6144ed84828501614395565b949350505050565b6020808252825182820181905260009190848201906040850190845b8181101561452d57835183529284019291840191600101614511565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561452d5783516001600160a01b031683529284019291840191600101614555565b6000806040838503121561458d57600080fd5b823561459881614023565b915060208301356145a881614023565b809150509250929050565b600081518084526020808501945080840160005b838110156145ec5781516001600160a01b0316875295820195908201906001016145c7565b509495945050505050565b602081526000611af160208301846145b3565b60006040828403121561461c57600080fd5b614624613ee0565b9050813581526020820135602082015292915050565b60006040828403121561464c57600080fd5b611af1838361460a565b60006020828403121561466857600080fd5b813567ffffffffffffffff81111561467f57600080fd5b820160e08185031215611af157600080fd5b600082601f8301126146a257600080fd5b813560206146b2613fec83614043565b82815260059290921b840181019181810190868411156146d157600080fd5b8286015b848110156140cb5780356146e88161427c565b83529183019183016146d5565b600082601f83011261470657600080fd5b81356020614716613fec83614043565b82815260069290921b8401810191818101908684111561473557600080fd5b8286015b848110156140cb5761474b888261460a565b835291830191604001614739565b600082601f83011261476a57600080fd5b8135602061477a613fec83614043565b82815260059290921b8401810191818101908684111561479957600080fd5b8286015b848110156140cb57803567ffffffffffffffff8111156147bd5760008081fd5b6147cb8986838b0101613fcd565b84525091830191830161479d565b6000602082840312156147eb57600080fd5b813567ffffffffffffffff8082111561480357600080fd5b90830190610120828603121561481857600080fd5b614820613f2c565b82358281111561482f57600080fd5b61483b87828601614691565b82525060208301358281111561485057600080fd5b61485c87828601614067565b60208301525060408301358281111561487457600080fd5b614880878286016143f9565b60408301525060608301358281111561489857600080fd5b6148a4878286016146f5565b6060830152506080830135828111156148bc57600080fd5b6148c887828601614759565b60808301525060a0830135828111156148e057600080fd5b6148ec878286016143f9565b60a08301525060c083013560c082015260e08301358281111561490e57600080fd5b61491a878286016143f9565b60e0830152506101009283013592810192909252509392505050565b60008083601f84011261494857600080fd5b50813567ffffffffffffffff81111561496057600080fd5b6020830191508360208260051b850101111561497b57600080fd5b9250929050565b60008060008060008060008060e0898b03121561499e57600080fd5b606089018a8111156149af57600080fd5b8998503567ffffffffffffffff808211156149c957600080fd5b818b0191508b601f8301126149dd57600080fd5b8135818111156149ec57600080fd5b8c60208285010111156149fe57600080fd5b6020830199508098505060808b0135915080821115614a1c57600080fd5b614a288c838d01614936565b909750955060a08b0135915080821115614a4157600080fd5b50614a4e8b828c01614936565b999c989b50969995989497949560c00135949350505050565b600060808284031215614a7957600080fd5b6040516080810181811067ffffffffffffffff82111715614a9c57614a9c613eb1565b604052823563ffffffff81168114614ab357600080fd5b81526020830135614ac38161427c565b60208201526040830135614ad68161427c565b60408201526060830135614ae98161427c565b60608201529392505050565b803560ff81168114613e0a57600080fd5b60008060008060008060c08789031215614b1f57600080fd5b863567ffffffffffffffff80821115614b3757600080fd5b614b438a838b01614067565b97506020890135915080821115614b5957600080fd5b614b658a838b01614067565b9650614b7360408a01614af5565b95506060890135915080821115614b8957600080fd5b614b958a838b01613fcd565b9450614ba360808a01614292565b935060a0890135915080821115614bb957600080fd5b50614bc689828a01613fcd565b9150509295509295509295565b600060208284031215614be557600080fd5b81518015158114611af157600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006000198203614c6657614c66614c24565b5060010190565b600082821015614c7f57614c7f614c24565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000816000190483118215151615614ccd57614ccd614c24565b500290565b60008219821115614ce557614ce5614c24565b500190565b600060ff821660ff84168060ff03821115614d0757614d07614c24565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff831680614d5157614d51614d0f565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600060208284031215614dcd57600080fd5b8151611af181614023565b600063ffffffff808316818516808303821115614df757614df7614c24565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614e308184018a6145b3565b90508281036080840152614e4481896145b3565b905060ff871660a084015282810360c0840152614e618187614356565b905067ffffffffffffffff851660e0840152828103610100840152614e868185614356565b9c9b505050505050505050505050565b600082614ea557614ea5614d0f565b500490565b600081518084526020808501945080840160005b838110156145ec57815180516001600160a01b031688528301518388015260409096019590820190600101614ebe565b60208152815160208201526000602083015160e06040840152614f15610100840182614356565b90506001600160a01b0360408501511660608401526060840151601f1980858403016080860152614f468383614356565b925060808601519150808584030160a0860152614f6383836145b3565b925060a08601519150808584030160c086015250614f818282614eaa565b91505060c084015160e08401528091505092915050565b8051613e0a8161427c565b8051613e0a81614023565b600082601f830112614fbf57600080fd5b8151614fcd613fec82613fa5565b818152846020838601011115614fe257600080fd5b6144ed82602083016020870161432a565b600082601f83011261500457600080fd5b81516020615014613fec83614043565b82815260069290921b8401810191818101908684111561503357600080fd5b8286015b848110156140cb57604081890312156150505760008081fd5b615058613ee0565b815161506381614023565b81528185015185820152835291830191604001615037565b60006020828403121561508d57600080fd5b815167ffffffffffffffff808211156150a557600080fd5b9083019061010082860312156150ba57600080fd5b6150c2613f50565b825181526150d260208401614f98565b60208201526150e360408401614fa3565b60408201526150f460608401614fa3565b606082015261510560808401614f98565b608082015260a08301518281111561511c57600080fd5b61512887828601614fae565b60a08301525060c08301518281111561514057600080fd5b61514c87828601614ff3565b60c08301525060e083015160e082015280935050505092915050565b600067ffffffffffffffff808316818516808303821115614df757614df7614c24565b600067ffffffffffffffff8083168181036151a8576151a8614c24565b6001019392505050565b600082601f8301126151c357600080fd5b815160206151d3613fec83614043565b82815260059290921b840181019181810190868411156151f257600080fd5b8286015b848110156140cb5780516152098161427c565b83529183019183016151f6565b600082601f83011261522757600080fd5b81516020615237613fec83614043565b82815260059290921b8401810191818101908684111561525657600080fd5b8286015b848110156140cb57805161526d81614023565b835291830191830161525a565b600082601f83011261528b57600080fd5b8151602061529b613fec83614043565b82815260059290921b840181019181810190868411156152ba57600080fd5b8286015b848110156140cb57805183529183019183016152be565b600082601f8301126152e657600080fd5b815160206152f6613fec83614043565b82815260069290921b8401810191818101908684111561531557600080fd5b8286015b848110156140cb57604081890312156153325760008081fd5b61533a613ee0565b815181528482015185820152835291830191604001615319565b600082601f83011261536557600080fd5b81516020615375613fec83614043565b82815260059290921b8401810191818101908684111561539457600080fd5b8286015b848110156140cb57805167ffffffffffffffff8111156153b85760008081fd5b6153c68986838b0101614fae565b845250918301918301615398565b6000602082840312156153e657600080fd5b815167ffffffffffffffff808211156153fe57600080fd5b90830190610120828603121561541357600080fd5b61541b613f2c565b82518281111561542a57600080fd5b615436878286016151b2565b82525060208301518281111561544b57600080fd5b61545787828601615216565b60208301525060408301518281111561546f57600080fd5b61547b8782860161527a565b60408301525060608301518281111561549357600080fd5b61549f878286016152d5565b6060830152506080830151828111156154b757600080fd5b6154c387828601615354565b60808301525060a0830151828111156154db57600080fd5b6154e78782860161527a565b60a08301525060c083015160c082015260e08301518281111561550957600080fd5b6155158782860161527a565b60e0830152506101009283015192810192909252509392505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b16604085015281606085015261556b8285018b6145b3565b9150838203608085015261557f828a6145b3565b915060ff881660a085015283820360c085015261559c8288614356565b90861660e08501528381036101008501529050614e868185614356565b602081526000611af16020830184614eaa565b600081518084526020808501945080840160005b838110156145ec578151875295820195908201906001016155e0565b60a08152600061560f60a08301886155cc565b828103602084015261562181886155cc565b9050856040840152828103606084015261563b81866155cc565b9150508260808301529695505050505050565b60006020828403121561566057600080fd5b5051919050565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156156af5780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
}

var EVM2EVMFreeOffRampABI = EVM2EVMFreeOffRampMetaData.ABI

var EVM2EVMFreeOffRampBin = EVM2EVMFreeOffRampMetaData.Bin

func DeployEVM2EVMFreeOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, onRampAddress common.Address, commitStore common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address) (common.Address, *types.Transaction, *EVM2EVMFreeOffRamp, error) {
	parsed, err := EVM2EVMFreeOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMFreeOffRampBin), backend, sourceChainId, chainId, offRampConfig, onRampAddress, commitStore, afn, sourceTokens, pools, rateLimiterConfig, tokenLimitsAdmin)
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
	return common.HexToHash("0x1a9ff9caccd597226ff8c393ab44e6b57669905d9b16f1cdb2ac267253dbf27c")
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
