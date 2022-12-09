// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_ge_offramp

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

type CommonEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

type EVM2EVMGEOffRampInterfaceGEOffRampConfig struct {
	GasOverhead                             *big.Int
	GasFeeCache                             common.Address
	PermissionLessExecutionThresholdSeconds uint32
	ExecutionDelaySeconds                   uint64
	MaxDataSize                             uint64
	MaxTokensLength                         uint64
}

type GEExecutionReport struct {
	SequenceNumbers          []uint64
	TokenPerFeeCoinAddresses []common.Address
	TokenPerFeeCoin          []*big.Int
	FeeUpdates               []GEFeeUpdate
	EncodedMessages          [][]byte
	InnerProofs              [][32]byte
	InnerProofFlagBits       *big.Int
	OuterProofs              [][32]byte
	OuterProofFlagBits       *big.Int
}

type GEFeeUpdate struct {
	ChainId        uint64
	LinkPerUnitGas *big.Int
}

type InternalAny2EVMMessageFromSender struct {
	SourceChainId        uint64
	Sender               []byte
	Receiver             common.Address
	Data                 []byte
	DestPools            []common.Address
	DestTokensAndAmounts []CommonEVMTokenAndAmount
	GasLimit             *big.Int
}

var EVM2EVMGEOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractGasFeeCacheInterface\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMGEOffRampInterface.GEOffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"seqNum\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedGasPriceUpdate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractGasFeeCacheInterface\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMGEOffRampInterface.GEOffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"GEOffRampConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitStore\",\"outputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGEConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractGasFeeCacheInterface\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMGEOffRampInterface.GEOffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"}],\"name\":\"getNopBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolByDestToken\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"linkPerUnitGas\",\"type\":\"uint256\"}],\"internalType\":\"structGE.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structGE.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"}],\"name\":\"setCommitStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractGasFeeCacheInterface\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMGEOffRampInterface.GEOffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setGEConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b50604051620066cd380380620066cd8339810160408190526200003591620008e8565b6000805460ff191681556001908c908c908b908b908b908b908b908b908b90829082908690869089903390819081620000b55760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ef57620000ef8162000540565b5050506001600160a01b0381166200011a57604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015c5760405162d8548360e71b815260040160405180910390fd5b81516200017190600590602085019062000657565b5060005b8251811015620002fa5760006040518060400160405280848481518110620001a157620001a162000a0b565b60200260200101516001600160a01b03168152602001836001600160601b031681525090508060036000868581518110620001e057620001e062000a0b565b6020908102919091018101516001600160a01b0390811683528282019390935260409091016000908120845194909201516001600160601b0316600160a01b02939092169290921790915581518451909160049186908690811062000249576200024962000a0b565b60200260200101516001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200028f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002b5919062000a21565b6001600160a01b039081168252602082019290925260400160002080546001600160a01b0319169290911691909117905550620002f28162000a48565b905062000175565b5050600680546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600955600a93909355600b55600c91909155871662000377576040516342bcdf7f60e11b815260040160405180910390fd5b886001600160401b03166080816001600160401b031681525050876001600160401b031660a0816001600160401b031681525050866001600160a01b031660c0816001600160a01b03168152505085600e60006101000a8154816001600160a01b0302191690836001600160a01b0316021790555050505050505050505080151560e081151581525050508860186000820151816000015560208201518160010160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555060608201518160010160186101000a8154816001600160401b0302191690836001600160401b0316021790555060808201518160020160006101000a8154816001600160401b0302191690836001600160401b0316021790555060a08201518160020160086101000a8154816001600160401b0302191690836001600160401b031602179055509050506200051e7fba22a5847647789e6efe1840c86bc66129ac89e03d7b95e0eebdf7fa43763fdd620005f160201b60201c565b610100526001600160a01b0316610120525062000a7098505050505050505050565b336001600160a01b038216036200059a5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ac565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008160805160a05160c0516040516020016200063a94939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b828054828255906000526020600020908101928215620006af579160200282015b82811115620006af57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000678565b50620006bd929150620006c1565b5090565b5b80821115620006bd5760008155600101620006c2565b80516001600160401b0381168114620006f057600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146200072157600080fd5b50565b600060c082840312156200073757600080fd5b60405160c081016001600160401b03811182821017156200075c576200075c620006f5565b806040525080915082518152602083015162000778816200070b565b6020820152604083015163ffffffff811681146200079557600080fd5b6040820152620007a860608401620006d8565b6060820152620007bb60808401620006d8565b6080820152620007ce60a08401620006d8565b60a08201525092915050565b8051620006f0816200070b565b600082601f830112620007f957600080fd5b815160206001600160401b0380831115620008185762000818620006f5565b8260051b604051601f19603f83011681018181108482111715620008405762000840620006f5565b6040529384528581018301938381019250878511156200085f57600080fd5b83870191505b848210156200088b5781516200087b816200070b565b8352918301919083019062000865565b979650505050505050565b600060408284031215620008a957600080fd5b604080519081016001600160401b0381118282101715620008ce57620008ce620006f5565b604052825181526020928301519281019290925250919050565b60008060008060008060008060008060006102208c8e0312156200090b57600080fd5b620009168c620006d8565b9a506200092660208d01620006d8565b9950620009378d60408e0162000724565b9850620009486101008d01620007da565b9750620009596101208d01620007da565b96506200096a6101408d01620007da565b6101608d01519096506001600160401b038111156200098857600080fd5b620009968e828f01620007e7565b6101808e015190965090506001600160401b03811115620009b657600080fd5b620009c48e828f01620007e7565b945050620009d78d6101a08e0162000896565b9250620009e86101e08d01620007da565b9150620009f96102008d01620007da565b90509295989b509295989b9093969950565b634e487b7160e01b600052603260045260246000fd5b60006020828403121562000a3457600080fd5b815162000a41816200070b565b9392505050565b60006001820162000a6957634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e0516101005161012051615bff62000ace6000396000613b64015260006136190152600061187301526000611e780152600061030c0152600081816102e701528181611e540152613dae0152615bff6000f3fe608060405234801561001057600080fd5b50600436106102d35760003560e01c806381ff704811610186578063b1dc65a4116100e3578063d1d8a83d11610097578063ea03d1b511610071578063ea03d1b514610850578063eb511dd41461085e578063f2fde38b1461087157600080fd5b8063d1d8a83d146107da578063d7e2bb5014610811578063e3d0e7121461083d57600080fd5b8063b66f0efb116100c8578063b66f0efb146107a3578063c0d78655146107b4578063c9033284146107c757600080fd5b8063b1dc65a41461077d578063b4069b311461079057600080fd5b806390c2339b1161013a5780639b1d13e61161011f5780639b1d13e614610739578063afcb95d71461074c578063b0f479a11461076c57600080fd5b806390c2339b146106eb578063918725431461072657600080fd5b8063856c82471161016b578063856c82471461068157806389c06568146106cd5780638da5cb5b146106d557600080fd5b806381ff7048146106495780638456cb591461067957600080fd5b80634352fa9f11610234578063681fba16116101e8578063744b92e2116101cd578063744b92e21461061957806379ba50971461062c578063814118341461063457600080fd5b8063681fba16146105f15780636c51aba01461060657600080fd5b8063599f643111610219578063599f6431146105a95780635c975abb146105ba5780635d86f141146105c557600080fd5b80634352fa9f146105765780634741062e1461058957600080fd5b80631628b6a71161028b5780632222dd42116102705780632222dd421461053657806339aa92641461055b5780633f4ba83a1461056e57600080fd5b80631628b6a7146103b6578063181f5a77146104ed57600080fd5b8063142a98fc116102bc578063142a98fc14610351578063147809b31461038b57806315fcd8c1146103a357600080fd5b8063087ae6df146102d8578063108ee5fc1461033c575b600080fd5b6040805167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811682527f0000000000000000000000000000000000000000000000000000000000000000166020820152015b60405180910390f35b61034f61034a366004614423565b610884565b005b61037e61035f366004614461565b67ffffffffffffffff166000908152600f602052604090205460ff1690565b6040516103339190614494565b61039361093b565b6040519015158152602001610333565b61034f6103b1366004614597565b6109c8565b6104806040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260185481526019546001600160a01b038116602083015274010000000000000000000000000000000000000000810463ffffffff1692820192909252780100000000000000000000000000000000000000000000000090910467ffffffffffffffff9081166060830152601a5480821660808401526801000000000000000090041660a082015290565b6040516103339190600060c082019050825182526001600160a01b03602084015116602083015263ffffffff6040840151166040830152606083015167ffffffffffffffff80821660608501528060808601511660808501528060a08601511660a0850152505092915050565b6105296040518060400160405280601681526020017f45564d3245564d47454f666652616d7020312e302e300000000000000000000081525081565b6040516103339190614696565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610333565b61034f610569366004614423565b610b6d565b61034f610ba4565b61034f6105843660046147a7565b610bb6565b61059c61059736600461480b565b610e0b565b6040516103339190614848565b6006546001600160a01b0316610543565b60005460ff16610393565b6105436105d3366004614423565b6001600160a01b039081166000908152600360205260409020541690565b6105f9610ed3565b604051610333919061488c565b61034f6106143660046149ce565b610f98565b61034f610627366004614ad6565b61100c565b61034f6113be565b61063c6114a6565b6040516103339190614b53565b6012546010546040805163ffffffff80851682526401000000009094049093166020840152820152606001610333565b61034f611508565b6106b461068f366004614423565b6001600160a01b031660009081526017602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff9091168152602001610333565b6105f9611518565b60005461010090046001600160a01b0316610543565b6106f3611578565b60405161033391908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b61034f610734366004614b66565b611619565b61034f610747366004614d21565b611745565b604080516001815260006020820181905291810191909152606001610333565b600d546001600160a01b0316610543565b61034f61078b366004614eca565b611750565b61054361079e366004614423565b611d23565b600e546001600160a01b0316610543565b61034f6107c2366004614423565b611e11565b61034f6107d5366004614423565b611ecf565b6108036107e8366004614423565b6001600160a01b031660009081526016602052604090205490565b604051908152602001610333565b61054361081f366004614423565b6001600160a01b039081166000908152600460205260409020541690565b61034f61084b366004614fc0565b611f06565b61034f6102d336600461508d565b61034f61086c366004614ad6565b612768565b61034f61087f366004614423565b6129e4565b61088c6129f5565b6001600160a01b0381166108cc576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa15801561099e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c291906150d8565b15905090565b6109d06129f5565b8051601855602081015160198054604080850151606086015167ffffffffffffffff90811678010000000000000000000000000000000000000000000000000277ffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff90931674010000000000000000000000000000000000000000027fffffffffffffffff0000000000000000000000000000000000000000000000009095166001600160a01b03909716969096179390931716939093179091556080830151601a805460a0860151841668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009091169290931691909117919091179055517fda126dc369452deb5c9e98a0522d5a2312155d7d64f0b41fa77adcddf459850190610b62908390600060c082019050825182526001600160a01b03602084015116602083015263ffffffff6040840151166040830152606083015167ffffffffffffffff80821660608501528060808601511660808501528060a08601511660a0850152505092915050565b60405180910390a150565b610b756129f5565b6006805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b610bac6129f5565b610bb4612a54565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015610bf357506006546001600160a01b03163314155b15610c2a576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610c66576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085460005b81811015610cc0576007600060088381548110610c8b57610c8b6150f3565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610cb98161511f565b9050610c6c565b5060005b82811015610df0576000858281518110610ce057610ce06150f3565b6020026020010151905060006001600160a01b0316816001600160a01b031603610d36576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610d4857610d486150f3565b602002602001015160076000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610dad57610dad6150f3565b6020026020010151604051610dd79291906001600160a01b03929092168252602082015260400190565b60405180910390a150610de98161511f565b9050610cc4565b508351610e0490600890602087019061436c565b5050505050565b80516060908067ffffffffffffffff811115610e2957610e296144bc565b604051908082528060200260200182016040528015610e52578160200160208202803683370190505b50915060005b81811015610ecc5760076000858381518110610e7657610e766150f3565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610eb157610eb16150f3565b6020908102919091010152610ec58161511f565b9050610e58565b5050919050565b60055460609067ffffffffffffffff811115610ef157610ef16144bc565b604051908082528060200260200182016040528015610f1a578160200160208202803683370190505b50905060005b600554811015610f9457610f5a60058281548110610f4057610f406150f3565b6000918252602090912001546001600160a01b0316611d23565b828281518110610f6c57610f6c6150f3565b6001600160a01b0390921660209283029190910190910152610f8d8161511f565b9050610f20565b5090565b333014610fd1576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a0810151511561100057610fe98160a00151612af0565b61100081608001518260a001518360400151612cf8565b61100981612d96565b50565b6110146129f5565b6005546000819003611052576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906110e0576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b03161461112f576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600561113e600185615139565b8154811061114e5761114e6150f3565b9060005260206000200160009054906101000a90046001600160a01b03169050600582602001516bffffffffffffffffffffffff1681548110611193576111936150f3565b6000918252602090912001546001600160a01b031660056111b5600186615139565b815481106111c5576111c56150f3565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600583602001516bffffffffffffffffffffffff1681548110611219576112196150f3565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff90921674010000000000000000000000000000000000000000029190921617905560058054806112a3576112a3615150565b6001900381819060005260206000200160006101000a8154906001600160a01b030219169055905560046000856001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561130d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113319190615166565b6001600160a01b03908116825260208083019390935260409182016000908120805473ffffffffffffffffffffffffffffffffffffffff1916905588821680825260038552838220919091558251908152908716928101929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b0316331461141d5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b606060158054806020026020016040519081016040528092919081815260200182805480156114fe57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116114e0575b5050505050905090565b6115106129f5565b610bb4612e6e565b606060058054806020026020016040519081016040528092919081815260200182805480156114fe576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116114e0575050505050905090565b6115a36040518060800160405280600081526020016000815260200160008152602001600081525090565b604080516080810182526009548152600a546020820152600b5491810191909152600c546060820181905242906000906115dd9083615139565b60208401518451919250611609916115f59084615183565b856040015161160491906151a2565b612ef6565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561165657506006546001600160a01b03163314155b1561168d576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff116116e1576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6116eb6009612f0c565b6020810151600a8190558151600955600b546117079190612ef6565b600b55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d9101610b62565b611009816001612fb9565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916117a691849163ffffffff851691908e908e908190840183828082843760009201919091525061325e92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260105480825260115460ff808216602085015261010090910416928201929092529083146118615760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401611414565b61186f8b8b8b8b8b8b61327c565b60007f0000000000000000000000000000000000000000000000000000000000000000156118cc576002826020015183604001516118ad91906151ba565b6118b791906151f5565b6118c29060016151ba565b60ff1690506118e2565b60208201516118dc9060016151ba565b60ff1690505b8881146119315760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401611414565b8887146119805760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401611414565b3360009081526013602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156119c3576119c361447e565b60028111156119d4576119d461447e565b90525090506002816020015160028111156119f1576119f161447e565b148015611a2b57506015816000015160ff1681548110611a1357611a136150f3565b6000918252602090912001546001600160a01b031633145b611a775760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401611414565b505050505060008888604051611a8e929190615217565b604051908190038120611aa5918c90602001615227565b604051602081830303815290604052805190602001209050611ac56143da565b604080518082019091526000808252602082015260005b88811015611d01576000600185888460208110611afb57611afb6150f3565b611b0891901a601b6151ba565b8d8d86818110611b1a57611b1a6150f3565b905060200201358c8c87818110611b3357611b336150f3565b9050602002013560405160008152602001604052604051611b70949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611b92573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526013602090815290849020838501909452835460ff80821685529296509294508401916101009004166002811115611be757611be761447e565b6002811115611bf857611bf861447e565b9052509250600183602001516002811115611c1557611c1561447e565b14611c625760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401611414565b8251849060ff16601f8110611c7957611c796150f3565b602002015115611ccb5760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401611414565b600184846000015160ff16601f8110611ce657611ce66150f3565b9115156020909202015250611cfa8161511f565b9050611adc565b5050505063ffffffff8110611d1857611d18615243565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611d77576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611de6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e0a9190615166565b9392505050565b611e196129f5565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038381169182179092556040805167ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681527f0000000000000000000000000000000000000000000000000000000000000000909316602084015290917f052b5907be1d3ac35d571862117562e80ee743c01251e388dafb7dc4e92a726c910160405180910390a250565b611ed76129f5565b600e805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b855185518560ff16601f831115611f79576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401611414565b80600003611fe3576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401611414565b818314612071576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401611414565b61207c816003615183565b83116120e4576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401611414565b6120ec6129f5565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b601454156122565760145460009061214490600190615139565b905060006014828154811061215b5761215b6150f3565b6000918252602082200154601580546001600160a01b0390921693509084908110612188576121886150f3565b60009182526020808320909101546001600160a01b03858116845260139092526040808420805461ffff19908116909155929091168084529220805490911690556014805491925090806121de576121de615150565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff19169055019055601580548061221e5761221e615150565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff191690550190555061212a915050565b60005b8151518110156125e1576000601360008460000151848151811061227f5761227f6150f3565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156122bc576122bc61447e565b14612323576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401611414565b6040805180820190915260ff82168152600160208201528251805160139160009185908110612354576123546150f3565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff1916176101008360028111156123ad576123ad61447e565b0217905550600091506123bd9050565b60136000846020015184815181106123d7576123d76150f3565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156124145761241461447e565b1461247b576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401611414565b6040805180820190915260ff8216815260208101600281525060136000846020015184815181106124ae576124ae6150f3565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff1916176101008360028111156125075761250761447e565b021790555050825180516014925083908110612525576125256150f3565b6020908102919091018101518254600181018455600093845292829020909201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039093169290921790915582015180516015919083908110612589576125896150f3565b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039092169190911790556125da8161511f565b9050612259565b5060408101516011805460ff191660ff909216919091179055601280547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612655928692908216911617615259565b92506101000a81548163ffffffff021916908363ffffffff1602179055506126b44630601260009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a0015161330c565b6010819055825180516011805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560125460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612753988b98919763ffffffff909216969095919491939192615281565b60405180910390a15050505050505050505050565b6127706129f5565b6001600160a01b038216158061278d57506001600160a01b038116155b156127c4576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612853576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038083168083526005546bffffffffffffffffffffffff90811660208086019182528785166000908152600382526040808220885194519095167401000000000000000000000000000000000000000002939096169290921790925583517f21df0da70000000000000000000000000000000000000000000000000000000081529351869460049492936321df0da79282870192819003870181865afa158015612908573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061292c9190615166565b6001600160a01b03908116825260208083019390935260409182016000908120805495831673ffffffffffffffffffffffffffffffffffffffff199687161790556005805460018101825591527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001805488831695168517905581519384528516918301919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b6129ec6129f5565b61100981613399565b60005461010090046001600160a01b03163314610bb45760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401611414565b60005460ff16612aa65760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401611414565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6000805b8251811015612bef57600060076000858481518110612b1557612b156150f3565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600003612ba857838281518110612b5e57612b5e6150f3565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401611414565b838281518110612bba57612bba6150f3565b60200260200101516020015181612bd19190615183565b612bdb90846151a2565b92505080612be89061511f565b9050612af4565b508015612cf457612c006009612f0c565b600a54811115612c4a57600a546040517f688ccf77000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401611414565b600b54811115612caa57600954600b5460009190612c689084615139565b612c729190615317565b9050806040517fe31e0f3200000000000000000000000000000000000000000000000000000000815260040161141491815260200190565b8060096002016000828254612cbf9190615139565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba3789060200161092f565b5050565b8151835114612d33576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015612d9057612d80848281518110612d5457612d546150f3565b6020026020010151848381518110612d6e57612d6e6150f3565b60200260200101516020015184613455565b612d898161511f565b9050612d36565b50505050565b60408101516001600160a01b03163b612dac5750565b600d546040517f382ab03a0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063382ab03a90612df590849060040161536f565b6020604051808303816000875af1158015612e14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e3891906150d8565b611009576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005460ff1615612ec15760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611414565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612ad33390565b6000818310612f055781611e0a565b5090919050565b6001810154600282015442911480612f275750808260030154145b15612f30575050565b816001015482600201541115612f72576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482612f849190615139565b60018401548454919250612fab91612f9c9084615183565b856002015461160491906151a2565b600284015550600390910155565b60005460ff161561300c5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611414565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561305f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061308391906150d8565b156130b9576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005a600d549091506001600160a01b0316613101576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606083015151156131c3578115613144576040517f198753d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60195460608401516040517f6e42efb80000000000000000000000000000000000000000000000000000000081526001600160a01b0390921691636e42efb8916131909160040161542b565b600060405180830381600087803b1580156131aa57600080fd5b505af11580156131be573d6000803e3d6000fd5b505050505b6131cd83836134cc565b8161325957670de0b6b3a764000083604001516000815181106131f2576131f26150f3565b60200260200101513a6018600001545a61320c9086615139565b61321691906151a2565b6132209190615183565b61322a9190615183565b6132349190615317565b33600090815260166020526040812080549091906132539084906151a2565b90915550505b505050565b61325981806020019051810190613275919061570a565b6000612fb9565b6000613289826020615183565b613294856020615183565b6132a0886101446151a2565b6132aa91906151a2565b6132b491906151a2565b6132bf9060006151a2565b9050368114613303576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401611414565b50505050505050565b6000808a8a8a8a8a8a8a8a8a60405160200161333099989796959493929190615867565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b038216036133f15760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401611414565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b1580156134b857600080fd5b505af1158015613303573d6000803e3d6000fd5b60808201515160008190036134e057505050565b60008167ffffffffffffffff8111156134fb576134fb6144bc565b604051908082528060200260200182016040528015613524578160200160208202803683370190505b50905060008267ffffffffffffffff811115613542576135426144bc565b6040519080825280602002602001820160405280156135d057816020015b60408051610180810182526000808252602080830182905292820181905260608083018290526080830182905260a0830182905260c0830182905260e083018290526101008301819052610120830152610140820181905261016082015282526000199092019101816135605790505b50905060005b8381101561368b576000866080015182815181106135f6576135f66150f3565b60200260200101518060200190518101906136119190615977565b905061363d817f0000000000000000000000000000000000000000000000000000000000000000613baa565b84838151811061364f5761364f6150f3565b6020026020010181815250508083838151811061366e5761366e6150f3565b602002602001018190525050806136849061511f565b90506135d6565b5060006136ac838760a001518860c001518960e001518a6101000151613cb4565b5060195490915060009074010000000000000000000000000000000000000000900463ffffffff166136de8342615139565b1190506000805b86811015613b59576000858281518110613701576137016150f3565b602002602001015190506000613734826020015167ffffffffffffffff166000908152600f602052604090205460ff1690565b9050600281600381111561374a5761374a61447e565b036137935760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611414565b89156137f15784806137b6575060038160038111156137b4576137b461447e565b145b6137ec576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61384e565b60008160038111156138055761380561447e565b1461384e5760208201516040517f67d9ba0f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611414565b60008160038111156138625761386261447e565b0361391357608082015160608301516001600160a01b031660009081526017602052604090205467ffffffffffffffff918216916138a291166001615aaa565b67ffffffffffffffff16146139015781606001516001600160a01b0316826080015167ffffffffffffffff167fd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf4123760405160405180910390a35050613b49565b604082015161391090856151a2565b93505b61391c82613dac565b60208281015167ffffffffffffffff166000908152600f90915260408120805460ff1916600117905561395661395184613f01565b6141d6565b60208085015167ffffffffffffffff166000908152600f909152604090208054919250829160ff191660018360038111156139935761399361447e565b02179055508a15613a79578260c0015180156139c0575060038260038111156139be576139be61447e565b145b80156139dd575060028160038111156139db576139db61447e565b145b80613a15575060008260038111156139f7576139f761447e565b148015613a1557506002816003811115613a1357613a1361447e565b145b15613a745760608301516001600160a01b03166000908152601760205260408120805467ffffffffffffffff1691613a4c83615acd565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b613af9565b8260c001518015613a9b57506003816003811115613a9957613a9961447e565b145b613af95760608301516001600160a01b03166000908152601760205260408120805467ffffffffffffffff1691613ad183615acd565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b826101600151836020015167ffffffffffffffff167f5e04b4755a4460aa6de4f3a906c4324a025c7449c02b52f5466659b5bfdfba5f83604051613b3d9190614494565b60405180910390a35050505b613b528161511f565b90506136e5565b506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116600090815260046020526040902054613ba091168230613455565b5050505050505050565b60008060001b828460200151856080015186606001518760e0015188610100015180519060200120896101200151604051602001613be89190615af4565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d60400151604051602001613c969c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600e546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce90613d11908c908c908c908c908c90600401615b37565b6020604051808303816000875af1158015613d30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d549190615b89565b905060008111613d90576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a613d9c9084615139565b9350935050509550959350505050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16816000015167ffffffffffffffff1614613e2c5780516040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611414565b601a54610120820151516801000000000000000090910467ffffffffffffffff161015613e975760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611414565b601a546101008201515167ffffffffffffffff909116101561100957601a54610100820151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301526024820152604401611414565b613f546040518060e00160405280600067ffffffffffffffff1681526020016060815260200160006001600160a01b03168152602001606081526020016060815260200160608152602001600081525090565b6101208201515160008167ffffffffffffffff811115613f7657613f766144bc565b604051908082528060200260200182016040528015613fbb57816020015b6040805180820190915260008082526020820152815260200190600190039081613f945790505b50905060008267ffffffffffffffff811115613fd957613fd96144bc565b604051908082528060200260200182016040528015614002578160200160208202803683370190505b50905060005b8381101561414257600061403d876101200151838151811061402c5761402c6150f3565b60200260200101516000015161430a565b905080838381518110614052576140526150f3565b60200260200101906001600160a01b031690816001600160a01b0316815250506040518060400160405280826001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156140bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906140df9190615166565b6001600160a01b031681526020018861012001518481518110614104576141046150f3565b602002602001015160200151815250848381518110614125576141256150f3565b6020026020010181905250508061413b9061511f565b9050614008565b506040518060e00160405280866000015167ffffffffffffffff168152602001866060015160405160200161418691906001600160a01b0391909116815260200190565b60405160208183030381529060405281526020018660e001516001600160a01b0316815260200186610100015181526020018281526020018381526020018660a001518152509350505050919050565b6040517f6c51aba00000000000000000000000000000000000000000000000000000000081526000903090636c51aba09061421590859060040161536f565b600060405180830381600087803b15801561422f57600080fd5b505af1925050508015614240575060015b614302573d80801561426e576040519150601f19603f3d011682016040523d82523d6000602084013e614273565b606091505b5061427d81615ba2565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da800000000000000000000000000000000000000000000000000000000036142ce5750600392915050565b806040517fcf19edfd0000000000000000000000000000000000000000000000000000000081526004016114149190614696565b506002919050565b6001600160a01b038181166000908152600360205260409020541680614367576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401611414565b919050565b8280548282559060005260206000209081019282156143ce579160200282015b828111156143ce578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390911617825560209092019160019091019061438c565b50610f949291506143f9565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610f9457600081556001016143fa565b6001600160a01b038116811461100957600080fd5b60006020828403121561443557600080fd5b8135611e0a8161440e565b67ffffffffffffffff8116811461100957600080fd5b803561436781614440565b60006020828403121561447357600080fd5b8135611e0a81614440565b634e487b7160e01b600052602160045260246000fd5b60208101600483106144b657634e487b7160e01b600052602160045260246000fd5b91905290565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156144f5576144f56144bc565b60405290565b60405160e0810167ffffffffffffffff811182821017156144f5576144f56144bc565b604051610120810167ffffffffffffffff811182821017156144f5576144f56144bc565b604051610180810167ffffffffffffffff811182821017156144f5576144f56144bc565b604051601f8201601f1916810167ffffffffffffffff8111828210171561458f5761458f6144bc565b604052919050565b600060c082840312156145a957600080fd5b60405160c0810181811067ffffffffffffffff821117156145cc576145cc6144bc565b6040528235815260208301356145e18161440e565b6020820152604083013563ffffffff811681146145fd57600080fd5b6040820152606083013561461081614440565b606082015261462160808401614456565b608082015261463260a08401614456565b60a08201529392505050565b60005b83811015614659578181015183820152602001614641565b83811115612d905750506000910152565b6000815180845261468281602086016020860161463e565b601f01601f19169290920160200192915050565b602081526000611e0a602083018461466a565b80356143678161440e565b600067ffffffffffffffff8211156146ce576146ce6144bc565b5060051b60200190565b600082601f8301126146e957600080fd5b813560206146fe6146f9836146b4565b614566565b82815260059290921b8401810191818101908684111561471d57600080fd5b8286015b848110156147415780356147348161440e565b8352918301918301614721565b509695505050505050565b600082601f83011261475d57600080fd5b8135602061476d6146f9836146b4565b82815260059290921b8401810191818101908684111561478c57600080fd5b8286015b848110156147415780358352918301918301614790565b600080604083850312156147ba57600080fd5b823567ffffffffffffffff808211156147d257600080fd5b6147de868387016146d8565b935060208501359150808211156147f457600080fd5b506148018582860161474c565b9150509250929050565b60006020828403121561481d57600080fd5b813567ffffffffffffffff81111561483457600080fd5b614840848285016146d8565b949350505050565b6020808252825182820181905260009190848201906040850190845b8181101561488057835183529284019291840191600101614864565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156148805783516001600160a01b0316835292840192918401916001016148a8565b600067ffffffffffffffff8211156148e7576148e76144bc565b50601f01601f191660200190565b600082601f83011261490657600080fd5b81356149146146f9826148cd565b81815284602083860101111561492957600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261495757600080fd5b813560206149676146f9836146b4565b82815260069290921b8401810191818101908684111561498657600080fd5b8286015b8481101561474157604081890312156149a35760008081fd5b6149ab6144d2565b81356149b68161440e565b8152818501358582015283529183019160400161498a565b6000602082840312156149e057600080fd5b813567ffffffffffffffff808211156149f857600080fd5b9083019060e08286031215614a0c57600080fd5b614a146144fb565b614a1d83614456565b8152602083013582811115614a3157600080fd5b614a3d878286016148f5565b602083015250614a4f604084016146a9565b6040820152606083013582811115614a6657600080fd5b614a72878286016148f5565b606083015250608083013582811115614a8a57600080fd5b614a96878286016146d8565b60808301525060a083013582811115614aae57600080fd5b614aba87828601614946565b60a08301525060c083013560c082015280935050505092915050565b60008060408385031215614ae957600080fd5b8235614af48161440e565b91506020830135614b048161440e565b809150509250929050565b600081518084526020808501945080840160005b83811015614b485781516001600160a01b031687529582019590820190600101614b23565b509495945050505050565b602081526000611e0a6020830184614b0f565b600060408284031215614b7857600080fd5b6040516040810181811067ffffffffffffffff82111715614b9b57614b9b6144bc565b604052823581526020928301359281019290925250919050565b600082601f830112614bc657600080fd5b81356020614bd66146f9836146b4565b82815260059290921b84018101918181019086841115614bf557600080fd5b8286015b84811015614741578035614c0c81614440565b8352918301918301614bf9565b600082601f830112614c2a57600080fd5b81356020614c3a6146f9836146b4565b82815260069290921b84018101918181019086841115614c5957600080fd5b8286015b848110156147415760408189031215614c765760008081fd5b614c7e6144d2565b8135614c8981614440565b81528185013585820152835291830191604001614c5d565b600082601f830112614cb257600080fd5b81356020614cc26146f9836146b4565b82815260059290921b84018101918181019086841115614ce157600080fd5b8286015b8481101561474157803567ffffffffffffffff811115614d055760008081fd5b614d138986838b01016148f5565b845250918301918301614ce5565b600060208284031215614d3357600080fd5b813567ffffffffffffffff80821115614d4b57600080fd5b908301906101208286031215614d6057600080fd5b614d6861451e565b823582811115614d7757600080fd5b614d8387828601614bb5565b825250602083013582811115614d9857600080fd5b614da4878286016146d8565b602083015250604083013582811115614dbc57600080fd5b614dc88782860161474c565b604083015250606083013582811115614de057600080fd5b614dec87828601614c19565b606083015250608083013582811115614e0457600080fd5b614e1087828601614ca1565b60808301525060a083013582811115614e2857600080fd5b614e348782860161474c565b60a08301525060c083013560c082015260e083013582811115614e5657600080fd5b614e628782860161474c565b60e0830152506101009283013592810192909252509392505050565b60008083601f840112614e9057600080fd5b50813567ffffffffffffffff811115614ea857600080fd5b6020830191508360208260051b8501011115614ec357600080fd5b9250929050565b60008060008060008060008060e0898b031215614ee657600080fd5b606089018a811115614ef757600080fd5b8998503567ffffffffffffffff80821115614f1157600080fd5b818b0191508b601f830112614f2557600080fd5b813581811115614f3457600080fd5b8c6020828501011115614f4657600080fd5b6020830199508098505060808b0135915080821115614f6457600080fd5b614f708c838d01614e7e565b909750955060a08b0135915080821115614f8957600080fd5b50614f968b828c01614e7e565b999c989b50969995989497949560c00135949350505050565b803560ff8116811461436757600080fd5b60008060008060008060c08789031215614fd957600080fd5b863567ffffffffffffffff80821115614ff157600080fd5b614ffd8a838b016146d8565b9750602089013591508082111561501357600080fd5b61501f8a838b016146d8565b965061502d60408a01614faf565b9550606089013591508082111561504357600080fd5b61504f8a838b016148f5565b945061505d60808a01614456565b935060a089013591508082111561507357600080fd5b5061508089828a016148f5565b9150509295509295509295565b60006020828403121561509f57600080fd5b813567ffffffffffffffff8111156150b657600080fd5b820160e08185031215611e0a57600080fd5b8051801515811461436757600080fd5b6000602082840312156150ea57600080fd5b611e0a826150c8565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019820361513257615132615109565b5060010190565b60008282101561514b5761514b615109565b500390565b634e487b7160e01b600052603160045260246000fd5b60006020828403121561517857600080fd5b8151611e0a8161440e565b600081600019048311821515161561519d5761519d615109565b500290565b600082198211156151b5576151b5615109565b500190565b600060ff821660ff84168060ff038211156151d7576151d7615109565b019392505050565b634e487b7160e01b600052601260045260246000fd5b600060ff831680615208576152086151df565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b634e487b7160e01b600052600160045260246000fd5b600063ffffffff80831681851680830382111561527857615278615109565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526152b18184018a614b0f565b905082810360808401526152c58189614b0f565b905060ff871660a084015282810360c08401526152e2818761466a565b905067ffffffffffffffff851660e0840152828103610100840152615307818561466a565b9c9b505050505050505050505050565b600082615326576153266151df565b500490565b600081518084526020808501945080840160005b83811015614b4857815180516001600160a01b03168852830151838801526040909601959082019060010161533f565b6020815267ffffffffffffffff82511660208201526000602083015160e060408401526153a061010084018261466a565b905060408401516153bc60608501826001600160a01b03169052565b506060840151601f19808584030160808601526153d9838361466a565b925060808601519150808584030160a08601526153f68383614b0f565b925060a08601519150808584030160c086015250615414828261532b565b91505060c084015160e08401528091505092915050565b602080825282518282018190526000919060409081850190868401855b82811015615477578151805167ffffffffffffffff168552860151868501529284019290850190600101615448565b5091979650505050505050565b805161436781614440565b600082601f8301126154a057600080fd5b815160206154b06146f9836146b4565b82815260059290921b840181019181810190868411156154cf57600080fd5b8286015b848110156147415780516154e681614440565b83529183019183016154d3565b80516143678161440e565b600082601f83011261550f57600080fd5b8151602061551f6146f9836146b4565b82815260059290921b8401810191818101908684111561553e57600080fd5b8286015b848110156147415780516155558161440e565b8352918301918301615542565b600082601f83011261557357600080fd5b815160206155836146f9836146b4565b82815260059290921b840181019181810190868411156155a257600080fd5b8286015b8481101561474157805183529183019183016155a6565b600082601f8301126155ce57600080fd5b815160206155de6146f9836146b4565b82815260069290921b840181019181810190868411156155fd57600080fd5b8286015b84811015614741576040818903121561561a5760008081fd5b6156226144d2565b815161562d81614440565b81528185015185820152835291830191604001615601565b600082601f83011261565657600080fd5b81516156646146f9826148cd565b81815284602083860101111561567957600080fd5b61484082602083016020870161463e565b600082601f83011261569b57600080fd5b815160206156ab6146f9836146b4565b82815260059290921b840181019181810190868411156156ca57600080fd5b8286015b8481101561474157805167ffffffffffffffff8111156156ee5760008081fd5b6156fc8986838b0101615645565b8452509183019183016156ce565b60006020828403121561571c57600080fd5b815167ffffffffffffffff8082111561573457600080fd5b90830190610120828603121561574957600080fd5b61575161451e565b82518281111561576057600080fd5b61576c8782860161548f565b82525060208301518281111561578157600080fd5b61578d878286016154fe565b6020830152506040830151828111156157a557600080fd5b6157b187828601615562565b6040830152506060830151828111156157c957600080fd5b6157d5878286016155bd565b6060830152506080830151828111156157ed57600080fd5b6157f98782860161568a565b60808301525060a08301518281111561581157600080fd5b61581d87828601615562565b60a08301525060c083015160c082015260e08301518281111561583f57600080fd5b61584b87828601615562565b60e0830152506101009283015192810192909252509392505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526158a18285018b614b0f565b915083820360808501526158b5828a614b0f565b915060ff881660a085015283820360c08501526158d2828861466a565b90861660e08501528381036101008501529050615307818561466a565b600082601f83011261590057600080fd5b815160206159106146f9836146b4565b82815260069290921b8401810191818101908684111561592f57600080fd5b8286015b84811015614741576040818903121561594c5760008081fd5b6159546144d2565b815161595f8161440e565b81528185015185820152835291830191604001615933565b60006020828403121561598957600080fd5b815167ffffffffffffffff808211156159a157600080fd5b9083019061018082860312156159b657600080fd5b6159be614542565b6159c783615484565b81526159d560208401615484565b6020820152604083015160408201526159f0606084016154f3565b6060820152615a0160808401615484565b608082015260a083015160a0820152615a1c60c084016150c8565b60c0820152615a2d60e084016154f3565b60e08201526101008084015183811115615a4657600080fd5b615a5288828701615645565b8284015250506101208084015183811115615a6c57600080fd5b615a78888287016158ef565b8284015250506101409150615a8e8284016154f3565b9181019190915261016091820151918101919091529392505050565b600067ffffffffffffffff80831681851680830382111561527857615278615109565b600067ffffffffffffffff808316818103615aea57615aea615109565b6001019392505050565b602081526000611e0a602083018461532b565b600081518084526020808501945080840160005b83811015614b4857815187529582019590820190600101615b1b565b60a081526000615b4a60a0830188615b07565b8281036020840152615b5c8188615b07565b90508560408401528281036060840152615b768186615b07565b9150508260808301529695505050505050565b600060208284031215615b9b57600080fd5b5051919050565b6000815160208301517fffffffff0000000000000000000000000000000000000000000000000000000080821693506004831015615bea5780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
}

var EVM2EVMGEOffRampABI = EVM2EVMGEOffRampMetaData.ABI

var EVM2EVMGEOffRampBin = EVM2EVMGEOffRampMetaData.Bin

func DeployEVM2EVMGEOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId uint64, chainId uint64, offRampConfig EVM2EVMGEOffRampInterfaceGEOffRampConfig, onRampAddress common.Address, commitStore common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address, feeToken common.Address) (common.Address, *types.Transaction, *EVM2EVMGEOffRamp, error) {
	parsed, err := EVM2EVMGEOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMGEOffRampBin), backend, sourceChainId, chainId, offRampConfig, onRampAddress, commitStore, afn, sourceTokens, pools, rateLimiterConfig, tokenLimitsAdmin, feeToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMGEOffRamp{EVM2EVMGEOffRampCaller: EVM2EVMGEOffRampCaller{contract: contract}, EVM2EVMGEOffRampTransactor: EVM2EVMGEOffRampTransactor{contract: contract}, EVM2EVMGEOffRampFilterer: EVM2EVMGEOffRampFilterer{contract: contract}}, nil
}

type EVM2EVMGEOffRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMGEOffRampCaller
	EVM2EVMGEOffRampTransactor
	EVM2EVMGEOffRampFilterer
}

type EVM2EVMGEOffRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMGEOffRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMGEOffRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMGEOffRampSession struct {
	Contract     *EVM2EVMGEOffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMGEOffRampCallerSession struct {
	Contract *EVM2EVMGEOffRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMGEOffRampTransactorSession struct {
	Contract     *EVM2EVMGEOffRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMGEOffRampRaw struct {
	Contract *EVM2EVMGEOffRamp
}

type EVM2EVMGEOffRampCallerRaw struct {
	Contract *EVM2EVMGEOffRampCaller
}

type EVM2EVMGEOffRampTransactorRaw struct {
	Contract *EVM2EVMGEOffRampTransactor
}

func NewEVM2EVMGEOffRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMGEOffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMGEOffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMGEOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRamp{address: address, abi: abi, EVM2EVMGEOffRampCaller: EVM2EVMGEOffRampCaller{contract: contract}, EVM2EVMGEOffRampTransactor: EVM2EVMGEOffRampTransactor{contract: contract}, EVM2EVMGEOffRampFilterer: EVM2EVMGEOffRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMGEOffRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMGEOffRampCaller, error) {
	contract, err := bindEVM2EVMGEOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampCaller{contract: contract}, nil
}

func NewEVM2EVMGEOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMGEOffRampTransactor, error) {
	contract, err := bindEVM2EVMGEOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampTransactor{contract: contract}, nil
}

func NewEVM2EVMGEOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMGEOffRampFilterer, error) {
	contract, err := bindEVM2EVMGEOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampFilterer{contract: contract}, nil
}

func bindEVM2EVMGEOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMGEOffRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMGEOffRamp.Contract.EVM2EVMGEOffRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.EVM2EVMGEOffRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.EVM2EVMGEOffRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMGEOffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(AggregateRateLimiterInterfaceTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(AggregateRateLimiterInterfaceTokenBucket)).(*AggregateRateLimiterInterfaceTokenBucket)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMGEOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMGEOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 InternalAny2EVMMessageFromSender) error {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) CcipReceive(arg0 InternalAny2EVMMessageFromSender) error {
	return _EVM2EVMGEOffRamp.Contract.CcipReceive(&_EVM2EVMGEOffRamp.CallOpts, arg0)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) CcipReceive(arg0 InternalAny2EVMMessageFromSender) error {
	return _EVM2EVMGEOffRamp.Contract.CcipReceive(&_EVM2EVMGEOffRamp.CallOpts, arg0)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetAFN() (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetAFN(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetAFN(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

	error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getChainIDs")

	outstruct := new(GetChainIDs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ChainId = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMGEOffRamp.Contract.GetChainIDs(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMGEOffRamp.Contract.GetChainIDs(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetCommitStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getCommitStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetCommitStore() (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetCommitStore(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetCommitStore() (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetCommitStore(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetDestinationToken(&_EVM2EVMGEOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetDestinationToken(&_EVM2EVMGEOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetDestinationTokens(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetDestinationTokens(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMGEOffRamp.Contract.GetExecutionState(&_EVM2EVMGEOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMGEOffRamp.Contract.GetExecutionState(&_EVM2EVMGEOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetGEConfig(opts *bind.CallOpts) (EVM2EVMGEOffRampInterfaceGEOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getGEConfig")

	if err != nil {
		return *new(EVM2EVMGEOffRampInterfaceGEOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMGEOffRampInterfaceGEOffRampConfig)).(*EVM2EVMGEOffRampInterfaceGEOffRampConfig)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetGEConfig() (EVM2EVMGEOffRampInterfaceGEOffRampConfig, error) {
	return _EVM2EVMGEOffRamp.Contract.GetGEConfig(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetGEConfig() (EVM2EVMGEOffRampInterfaceGEOffRampConfig, error) {
	return _EVM2EVMGEOffRamp.Contract.GetGEConfig(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetNopBalance(opts *bind.CallOpts, nop common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getNopBalance", nop)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetNopBalance(nop common.Address) (*big.Int, error) {
	return _EVM2EVMGEOffRamp.Contract.GetNopBalance(&_EVM2EVMGEOffRamp.CallOpts, nop)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetNopBalance(nop common.Address) (*big.Int, error) {
	return _EVM2EVMGEOffRamp.Contract.GetNopBalance(&_EVM2EVMGEOffRamp.CallOpts, nop)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getPoolByDestToken", destToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetPoolByDestToken(destToken common.Address) (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetPoolByDestToken(&_EVM2EVMGEOffRamp.CallOpts, destToken)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetPoolByDestToken(destToken common.Address) (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetPoolByDestToken(&_EVM2EVMGEOffRamp.CallOpts, destToken)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getPoolBySourceToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetPoolBySourceToken(&_EVM2EVMGEOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetPoolBySourceToken(&_EVM2EVMGEOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetPoolTokens(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetPoolTokens(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMGEOffRamp.Contract.GetPricesForTokens(&_EVM2EVMGEOffRamp.CallOpts, tokens)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMGEOffRamp.Contract.GetPricesForTokens(&_EVM2EVMGEOffRamp.CallOpts, tokens)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetRouter() (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetRouter(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetRouter(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getSenderNonce", sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMGEOffRamp.Contract.GetSenderNonce(&_EVM2EVMGEOffRamp.CallOpts, sender)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMGEOffRamp.Contract.GetSenderNonce(&_EVM2EVMGEOffRamp.CallOpts, sender)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMGEOffRamp.Contract.IsAFNHealthy(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMGEOffRamp.Contract.IsAFNHealthy(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMGEOffRamp.Contract.LatestConfigDetails(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMGEOffRamp.Contract.LatestConfigDetails(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMGEOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMGEOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) Owner() (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.Owner(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.Owner(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) Paused() (bool, error) {
	return _EVM2EVMGEOffRamp.Contract.Paused(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMGEOffRamp.Contract.Paused(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.Transmitters(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.Transmitters(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMGEOffRamp.Contract.TypeAndVersion(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMGEOffRamp.Contract.TypeAndVersion(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.AcceptOwnership(&_EVM2EVMGEOffRamp.TransactOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.AcceptOwnership(&_EVM2EVMGEOffRamp.TransactOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.AddPool(&_EVM2EVMGEOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.AddPool(&_EVM2EVMGEOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "executeSingleMessage", message)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) ExecuteSingleMessage(message InternalAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMGEOffRamp.TransactOpts, message)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) ExecuteSingleMessage(message InternalAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMGEOffRamp.TransactOpts, message)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) ManuallyExecute(opts *bind.TransactOpts, report GEExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "manuallyExecute", report)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) ManuallyExecute(report GEExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.ManuallyExecute(&_EVM2EVMGEOffRamp.TransactOpts, report)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) ManuallyExecute(report GEExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.ManuallyExecute(&_EVM2EVMGEOffRamp.TransactOpts, report)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.Pause(&_EVM2EVMGEOffRamp.TransactOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.Pause(&_EVM2EVMGEOffRamp.TransactOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.RemovePool(&_EVM2EVMGEOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.RemovePool(&_EVM2EVMGEOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetAFN(&_EVM2EVMGEOffRamp.TransactOpts, afn)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetAFN(&_EVM2EVMGEOffRamp.TransactOpts, afn)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) SetCommitStore(opts *bind.TransactOpts, commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "setCommitStore", commitStore)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) SetCommitStore(commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetCommitStore(&_EVM2EVMGEOffRamp.TransactOpts, commitStore)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) SetCommitStore(commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetCommitStore(&_EVM2EVMGEOffRamp.TransactOpts, commitStore)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetConfig(&_EVM2EVMGEOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetConfig(&_EVM2EVMGEOffRamp.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) SetGEConfig(opts *bind.TransactOpts, config EVM2EVMGEOffRampInterfaceGEOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "setGEConfig", config)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) SetGEConfig(config EVM2EVMGEOffRampInterfaceGEOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetGEConfig(&_EVM2EVMGEOffRamp.TransactOpts, config)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) SetGEConfig(config EVM2EVMGEOffRampInterfaceGEOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetGEConfig(&_EVM2EVMGEOffRamp.TransactOpts, config)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetPrices(&_EVM2EVMGEOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetPrices(&_EVM2EVMGEOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMGEOffRamp.TransactOpts, config)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMGEOffRamp.TransactOpts, config)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetRouter(&_EVM2EVMGEOffRamp.TransactOpts, router)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetRouter(&_EVM2EVMGEOffRamp.TransactOpts, router)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMGEOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMGEOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.TransferOwnership(&_EVM2EVMGEOffRamp.TransactOpts, to)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.TransferOwnership(&_EVM2EVMGEOffRamp.TransactOpts, to)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.Transmit(&_EVM2EVMGEOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.Transmit(&_EVM2EVMGEOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.Unpause(&_EVM2EVMGEOffRamp.TransactOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.Unpause(&_EVM2EVMGEOffRamp.TransactOpts)
}

type EVM2EVMGEOffRampAFNSetIterator struct {
	Event *EVM2EVMGEOffRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampAFNSet)
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
		it.Event = new(EVM2EVMGEOffRampAFNSet)
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

func (it *EVM2EVMGEOffRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMGEOffRampAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampAFNSetIterator{contract: _EVM2EVMGEOffRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampAFNSet)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseAFNSet(log types.Log) (*EVM2EVMGEOffRampAFNSet, error) {
	event := new(EVM2EVMGEOffRampAFNSet)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampConfigChangedIterator struct {
	Event *EVM2EVMGEOffRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampConfigChanged)
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
		it.Event = new(EVM2EVMGEOffRampConfigChanged)
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

func (it *EVM2EVMGEOffRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMGEOffRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampConfigChangedIterator{contract: _EVM2EVMGEOffRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampConfigChanged)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMGEOffRampConfigChanged, error) {
	event := new(EVM2EVMGEOffRampConfigChanged)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampConfigSetIterator struct {
	Event *EVM2EVMGEOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampConfigSet)
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
		it.Event = new(EVM2EVMGEOffRampConfigSet)
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

func (it *EVM2EVMGEOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampConfigSet struct {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMGEOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampConfigSetIterator{contract: _EVM2EVMGEOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampConfigSet)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseConfigSet(log types.Log) (*EVM2EVMGEOffRampConfigSet, error) {
	event := new(EVM2EVMGEOffRampConfigSet)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampExecutionStateChangedIterator struct {
	Event *EVM2EVMGEOffRampExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampExecutionStateChanged)
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
		it.Event = new(EVM2EVMGEOffRampExecutionStateChanged)
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

func (it *EVM2EVMGEOffRampExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampExecutionStateChanged struct {
	SequenceNumber uint64
	MessageId      [32]byte
	State          uint8
	Raw            types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64, messageId [][32]byte) (*EVM2EVMGEOffRampExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampExecutionStateChangedIterator{contract: _EVM2EVMGEOffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampExecutionStateChanged, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampExecutionStateChanged)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseExecutionStateChanged(log types.Log) (*EVM2EVMGEOffRampExecutionStateChanged, error) {
	event := new(EVM2EVMGEOffRampExecutionStateChanged)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampGEOffRampConfigChangedIterator struct {
	Event *EVM2EVMGEOffRampGEOffRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampGEOffRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampGEOffRampConfigChanged)
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
		it.Event = new(EVM2EVMGEOffRampGEOffRampConfigChanged)
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

func (it *EVM2EVMGEOffRampGEOffRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampGEOffRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampGEOffRampConfigChanged struct {
	Arg0 EVM2EVMGEOffRampInterfaceGEOffRampConfig
	Raw  types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterGEOffRampConfigChanged(opts *bind.FilterOpts) (*EVM2EVMGEOffRampGEOffRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "GEOffRampConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampGEOffRampConfigChangedIterator{contract: _EVM2EVMGEOffRamp.contract, event: "GEOffRampConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchGEOffRampConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampGEOffRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "GEOffRampConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampGEOffRampConfigChanged)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "GEOffRampConfigChanged", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseGEOffRampConfigChanged(log types.Log) (*EVM2EVMGEOffRampGEOffRampConfigChanged, error) {
	event := new(EVM2EVMGEOffRampGEOffRampConfigChanged)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "GEOffRampConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampOffRampConfigSetIterator struct {
	Event *EVM2EVMGEOffRampOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampOffRampConfigSet)
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
		it.Event = new(EVM2EVMGEOffRampOffRampConfigSet)
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

func (it *EVM2EVMGEOffRampOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampOffRampConfigSet struct {
	Config BaseOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMGEOffRampOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampOffRampConfigSetIterator{contract: _EVM2EVMGEOffRamp.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampOffRampConfigSet)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseOffRampConfigSet(log types.Log) (*EVM2EVMGEOffRampOffRampConfigSet, error) {
	event := new(EVM2EVMGEOffRampOffRampConfigSet)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampOffRampRouterSetIterator struct {
	Event *EVM2EVMGEOffRampOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampOffRampRouterSet)
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
		it.Event = new(EVM2EVMGEOffRampOffRampRouterSet)
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

func (it *EVM2EVMGEOffRampOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampOffRampRouterSet struct {
	Router        common.Address
	SourceChainId uint64
	OnRampAddress common.Address
	Raw           types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMGEOffRampOffRampRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampOffRampRouterSetIterator{contract: _EVM2EVMGEOffRamp.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampOffRampRouterSet)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseOffRampRouterSet(log types.Log) (*EVM2EVMGEOffRampOffRampRouterSet, error) {
	event := new(EVM2EVMGEOffRampOffRampRouterSet)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMGEOffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMGEOffRampOwnershipTransferRequested)
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

func (it *EVM2EVMGEOffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMGEOffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampOwnershipTransferRequestedIterator{contract: _EVM2EVMGEOffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampOwnershipTransferRequested)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMGEOffRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMGEOffRampOwnershipTransferRequested)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampOwnershipTransferredIterator struct {
	Event *EVM2EVMGEOffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMGEOffRampOwnershipTransferred)
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

func (it *EVM2EVMGEOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMGEOffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampOwnershipTransferredIterator{contract: _EVM2EVMGEOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampOwnershipTransferred)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMGEOffRampOwnershipTransferred, error) {
	event := new(EVM2EVMGEOffRampOwnershipTransferred)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampPausedIterator struct {
	Event *EVM2EVMGEOffRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampPaused)
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
		it.Event = new(EVM2EVMGEOffRampPaused)
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

func (it *EVM2EVMGEOffRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMGEOffRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampPausedIterator{contract: _EVM2EVMGEOffRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampPaused)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParsePaused(log types.Log) (*EVM2EVMGEOffRampPaused, error) {
	event := new(EVM2EVMGEOffRampPaused)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampPoolAddedIterator struct {
	Event *EVM2EVMGEOffRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampPoolAdded)
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
		it.Event = new(EVM2EVMGEOffRampPoolAdded)
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

func (it *EVM2EVMGEOffRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMGEOffRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampPoolAddedIterator{contract: _EVM2EVMGEOffRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampPoolAdded)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMGEOffRampPoolAdded, error) {
	event := new(EVM2EVMGEOffRampPoolAdded)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampPoolRemovedIterator struct {
	Event *EVM2EVMGEOffRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampPoolRemoved)
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
		it.Event = new(EVM2EVMGEOffRampPoolRemoved)
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

func (it *EVM2EVMGEOffRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMGEOffRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampPoolRemovedIterator{contract: _EVM2EVMGEOffRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampPoolRemoved)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMGEOffRampPoolRemoved, error) {
	event := new(EVM2EVMGEOffRampPoolRemoved)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampSkippedIncorrectNonceIterator struct {
	Event *EVM2EVMGEOffRampSkippedIncorrectNonce

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampSkippedIncorrectNonceIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampSkippedIncorrectNonce)
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
		it.Event = new(EVM2EVMGEOffRampSkippedIncorrectNonce)
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

func (it *EVM2EVMGEOffRampSkippedIncorrectNonceIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampSkippedIncorrectNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampSkippedIncorrectNonce struct {
	Nonce  uint64
	Sender common.Address
	Raw    types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterSkippedIncorrectNonce(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMGEOffRampSkippedIncorrectNonceIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "SkippedIncorrectNonce", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampSkippedIncorrectNonceIterator{contract: _EVM2EVMGEOffRamp.contract, event: "SkippedIncorrectNonce", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchSkippedIncorrectNonce(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampSkippedIncorrectNonce, nonce []uint64, sender []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "SkippedIncorrectNonce", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampSkippedIncorrectNonce)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "SkippedIncorrectNonce", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseSkippedIncorrectNonce(log types.Log) (*EVM2EVMGEOffRampSkippedIncorrectNonce, error) {
	event := new(EVM2EVMGEOffRampSkippedIncorrectNonce)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "SkippedIncorrectNonce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampTokenPriceChangedIterator struct {
	Event *EVM2EVMGEOffRampTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampTokenPriceChanged)
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
		it.Event = new(EVM2EVMGEOffRampTokenPriceChanged)
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

func (it *EVM2EVMGEOffRampTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMGEOffRampTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampTokenPriceChangedIterator{contract: _EVM2EVMGEOffRamp.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampTokenPriceChanged)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMGEOffRampTokenPriceChanged, error) {
	event := new(EVM2EVMGEOffRampTokenPriceChanged)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMGEOffRampTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMGEOffRampTokensRemovedFromBucket)
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

func (it *EVM2EVMGEOffRampTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMGEOffRampTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampTokensRemovedFromBucketIterator{contract: _EVM2EVMGEOffRamp.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampTokensRemovedFromBucket)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMGEOffRampTokensRemovedFromBucket, error) {
	event := new(EVM2EVMGEOffRampTokensRemovedFromBucket)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampTransmittedIterator struct {
	Event *EVM2EVMGEOffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampTransmitted)
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
		it.Event = new(EVM2EVMGEOffRampTransmitted)
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

func (it *EVM2EVMGEOffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMGEOffRampTransmittedIterator, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampTransmittedIterator{contract: _EVM2EVMGEOffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampTransmitted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampTransmitted)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseTransmitted(log types.Log) (*EVM2EVMGEOffRampTransmitted, error) {
	event := new(EVM2EVMGEOffRampTransmitted)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMGEOffRampUnpausedIterator struct {
	Event *EVM2EVMGEOffRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMGEOffRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMGEOffRampUnpaused)
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
		it.Event = new(EVM2EVMGEOffRampUnpaused)
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

func (it *EVM2EVMGEOffRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMGEOffRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMGEOffRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMGEOffRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampUnpausedIterator{contract: _EVM2EVMGEOffRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMGEOffRampUnpaused)
				if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMGEOffRampUnpaused, error) {
	event := new(EVM2EVMGEOffRampUnpaused)
	if err := _EVM2EVMGEOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetChainIDs struct {
	SourceChainId uint64
	ChainId       uint64
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMGEOffRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMGEOffRamp.ParseAFNSet(log)
	case _EVM2EVMGEOffRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMGEOffRamp.ParseConfigChanged(log)
	case _EVM2EVMGEOffRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMGEOffRamp.ParseConfigSet(log)
	case _EVM2EVMGEOffRamp.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMGEOffRamp.ParseExecutionStateChanged(log)
	case _EVM2EVMGEOffRamp.abi.Events["GEOffRampConfigChanged"].ID:
		return _EVM2EVMGEOffRamp.ParseGEOffRampConfigChanged(log)
	case _EVM2EVMGEOffRamp.abi.Events["OffRampConfigSet"].ID:
		return _EVM2EVMGEOffRamp.ParseOffRampConfigSet(log)
	case _EVM2EVMGEOffRamp.abi.Events["OffRampRouterSet"].ID:
		return _EVM2EVMGEOffRamp.ParseOffRampRouterSet(log)
	case _EVM2EVMGEOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMGEOffRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMGEOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMGEOffRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMGEOffRamp.abi.Events["Paused"].ID:
		return _EVM2EVMGEOffRamp.ParsePaused(log)
	case _EVM2EVMGEOffRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMGEOffRamp.ParsePoolAdded(log)
	case _EVM2EVMGEOffRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMGEOffRamp.ParsePoolRemoved(log)
	case _EVM2EVMGEOffRamp.abi.Events["SkippedIncorrectNonce"].ID:
		return _EVM2EVMGEOffRamp.ParseSkippedIncorrectNonce(log)
	case _EVM2EVMGEOffRamp.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMGEOffRamp.ParseTokenPriceChanged(log)
	case _EVM2EVMGEOffRamp.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMGEOffRamp.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMGEOffRamp.abi.Events["Transmitted"].ID:
		return _EVM2EVMGEOffRamp.ParseTransmitted(log)
	case _EVM2EVMGEOffRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMGEOffRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMGEOffRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMGEOffRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMGEOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMGEOffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x5e04b4755a4460aa6de4f3a906c4324a025c7449c02b52f5466659b5bfdfba5f")
}

func (EVM2EVMGEOffRampGEOffRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0xda126dc369452deb5c9e98a0522d5a2312155d7d64f0b41fa77adcddf4598501")
}

func (EVM2EVMGEOffRampOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1a9ff9caccd597226ff8c393ab44e6b57669905d9b16f1cdb2ac267253dbf27c")
}

func (EVM2EVMGEOffRampOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x052b5907be1d3ac35d571862117562e80ee743c01251e388dafb7dc4e92a726c")
}

func (EVM2EVMGEOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMGEOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMGEOffRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMGEOffRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMGEOffRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMGEOffRampSkippedIncorrectNonce) Topic() common.Hash {
	return common.HexToHash("0xd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf41237")
}

func (EVM2EVMGEOffRampTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMGEOffRampTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMGEOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (EVM2EVMGEOffRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRamp) Address() common.Address {
	return _EVM2EVMGEOffRamp.address
}

type EVM2EVMGEOffRampInterface interface {
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error)

	CcipReceive(opts *bind.CallOpts, arg0 InternalAny2EVMMessageFromSender) error

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

		error)

	GetCommitStore(opts *bind.CallOpts) (common.Address, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetGEConfig(opts *bind.CallOpts) (EVM2EVMGEOffRampInterfaceGEOffRampConfig, error)

	GetNopBalance(opts *bind.CallOpts, nop common.Address) (*big.Int, error)

	GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error)

	GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error)

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

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMMessageFromSender) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report GEExecutionReport) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetCommitStore(opts *bind.TransactOpts, commitStore common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetGEConfig(opts *bind.TransactOpts, config EVM2EVMGEOffRampInterfaceGEOffRampConfig) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMGEOffRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMGEOffRampAFNSet, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMGEOffRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMGEOffRampConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMGEOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMGEOffRampConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64, messageId [][32]byte) (*EVM2EVMGEOffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampExecutionStateChanged, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMGEOffRampExecutionStateChanged, error)

	FilterGEOffRampConfigChanged(opts *bind.FilterOpts) (*EVM2EVMGEOffRampGEOffRampConfigChangedIterator, error)

	WatchGEOffRampConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampGEOffRampConfigChanged) (event.Subscription, error)

	ParseGEOffRampConfigChanged(log types.Log) (*EVM2EVMGEOffRampGEOffRampConfigChanged, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMGEOffRampOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*EVM2EVMGEOffRampOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMGEOffRampOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampOffRampRouterSet, router []common.Address) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*EVM2EVMGEOffRampOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMGEOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMGEOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMGEOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMGEOffRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMGEOffRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMGEOffRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMGEOffRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMGEOffRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMGEOffRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMGEOffRampPoolRemoved, error)

	FilterSkippedIncorrectNonce(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMGEOffRampSkippedIncorrectNonceIterator, error)

	WatchSkippedIncorrectNonce(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampSkippedIncorrectNonce, nonce []uint64, sender []common.Address) (event.Subscription, error)

	ParseSkippedIncorrectNonce(log types.Log) (*EVM2EVMGEOffRampSkippedIncorrectNonce, error)

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMGEOffRampTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMGEOffRampTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMGEOffRampTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMGEOffRampTokensRemovedFromBucket, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMGEOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMGEOffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMGEOffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMGEOffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
