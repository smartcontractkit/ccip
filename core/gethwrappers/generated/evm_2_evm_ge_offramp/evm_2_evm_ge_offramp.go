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

type EVM2EVMGEOffRampInterfaceGEOffRampConfig struct {
	FeeTokenDest                            common.Address
	GasOverhead                             *big.Int
	GasFeeCache                             common.Address
	PermissionLessExecutionThresholdSeconds uint32
	ExecutionDelaySeconds                   uint64
	MaxDataSize                             uint64
	MaxTokensLength                         uint64
}

var EVM2EVMGEOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenDest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractGasFeeCacheInterface\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMGEOffRampInterface.GEOffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"seqNum\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedGasPriceUpdate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenDest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractGasFeeCacheInterface\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMGEOffRampInterface.GEOffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"GEOffRampConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitStore\",\"outputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGEConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenDest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractGasFeeCacheInterface\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMGEOffRampInterface.GEOffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCommitStoreInterface\",\"name\":\"commitStore\",\"type\":\"address\"}],\"name\":\"setCommitStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeTokenDest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractGasFeeCacheInterface\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMGEOffRampInterface.GEOffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setGEConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200632638038062006326833981016040819052620000359162000799565b6000805460ff191681556001908b908b908a908a908a908a908a908a908a90829082908690869089903390819081620000b55760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ef57620000ef81620003a4565b5050506001600160a01b0381166200011a57604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015c5760405162d8548360e71b815260040160405180910390fd5b81516200017190600490602085019062000455565b5060005b82518110156200023c57600082828151811062000196576200019662000890565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001e057620001e062000890565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b0292909116919091179055506200023481620008a6565b905062000175565b5050600580546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600855600993909355600a55600b919091558716620002b9576040516342bcdf7f60e11b815260040160405180910390fd5b5050506080958652505060a09283526001600160a01b0391821660c0908152600d80549284166001600160a01b031993841617905594151560e0528c51601780549184169190921617905560208c015160185560408c01516019805460608f0151958f01516001600160401b03908116600160c01b026001600160c01b0363ffffffff909816600160a01b026001600160c01b031990931694909516939093171794909416919091179092558a0151601a80549b90930151821668010000000000000000026001600160801b0319909b169116179890981790975550620008ce975050505050505050565b336001600160a01b03821603620003fe5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ac565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b828054828255906000526020600020908101928215620004ad579160200282015b82811115620004ad57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000476565b50620004bb929150620004bf565b5090565b5b80821115620004bb5760008155600101620004c0565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b0381118282101715620005115762000511620004d6565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620005425762000542620004d6565b604052919050565b6001600160a01b03811681146200056057600080fd5b50565b805162000570816200054a565b919050565b80516001600160401b03811681146200057057600080fd5b600060e08284031215620005a057600080fd5b620005aa620004ec565b90508151620005b9816200054a565b8152602082810151908201526040820151620005d5816200054a565b6040820152606082015163ffffffff81168114620005f257600080fd5b6060820152620006056080830162000575565b60808201526200061860a0830162000575565b60a08201526200062b60c0830162000575565b60c082015292915050565b60006001600160401b03821115620006525762000652620004d6565b5060051b60200190565b600082601f8301126200066e57600080fd5b8151602062000687620006818362000636565b62000517565b82815260059290921b84018101918181019086841115620006a757600080fd5b8286015b84811015620006cf578051620006c1816200054a565b8352918301918301620006ab565b509695505050505050565b600082601f830112620006ec57600080fd5b81516020620006ff620006818362000636565b82815260059290921b840181019181810190868411156200071f57600080fd5b8286015b84811015620006cf57805162000739816200054a565b835291830191830162000723565b6000604082840312156200075a57600080fd5b604080519081016001600160401b03811182821017156200077f576200077f620004d6565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000806102208b8d031215620007ba57600080fd5b8a51995060208b01519850620007d48c60408d016200058d565b9750620007e56101208c0162000563565b9650620007f66101408c0162000563565b9550620008076101608c0162000563565b6101808c01519095506001600160401b03808211156200082657600080fd5b620008348e838f016200065c565b95506101a08d01519150808211156200084c57600080fd5b506200085b8d828e01620006da565b9350506200086e8c6101c08d0162000747565b91506200087f6102008c0162000563565b90509295989b9194979a5092959850565b634e487b7160e01b600052603260045260246000fd5b600060018201620008c757634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e051615a096200091d600039600061156e015260006133d10152600081816102d501526133ab0152600081816102b2015281816133860152613b320152615a096000f3fe608060405234801561001057600080fd5b50600436106102925760003560e01c80638456cb5911610160578063b1dc65a4116100d8578063c0d786551161008c578063e3d0e71211610071578063e3d0e71214610770578063eb511dd414610783578063f2fde38b1461079657600080fd5b8063c0d786551461074a578063c90332841461075d57600080fd5b8063b66f0efb116100bd578063b66f0efb146106fa578063bbe4f6db1461070b578063bc1213541461073757600080fd5b8063b1dc65a4146106d4578063b4069b31146106e757600080fd5b8063918725431161012f578063a8e9132111610114578063a8e9132114610690578063afcb95d7146106a3578063b0f479a1146106c357600080fd5b8063918725431461066f5780639438ff631461068257600080fd5b80638456cb591461060e57806389c06568146106165780638da5cb5b1461061e57806390c2339b1461063457600080fd5b80633f4ba83a1161020e578063681fba16116101c257806379ba5097116101a757806379ba5097146105c157806381411834146105c957806381ff7048146105de57600080fd5b8063681fba1614610599578063744b92e2146105ae57600080fd5b80634741062e116101f35780634741062e1461055d578063599f64311461057d5780635c975abb1461058e57600080fd5b80633f4ba83a146105425780634352fa9f1461054a57600080fd5b8063147809b311610265578063181f5a771161024a578063181f5a77146104c15780632222dd421461050a57806339aa92641461052f57600080fd5b8063147809b3146103515780631628b6a71461036957600080fd5b806307a22a0714610297578063087ae6df146102ac578063108ee5fc14610304578063142a98fc14610317575b600080fd5b6102aa6102a536600461440d565b6107a9565b005b604080517f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006020820152015b60405180910390f35b6102aa61031236600461450e565b61081d565b61034461032536600461454c565b67ffffffffffffffff166000908152600e602052604090205460ff1690565b6040516102fb9190614598565b6103596108d4565b60405190151581526020016102fb565b6104466040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152506040805160e0810182526017546001600160a01b03908116825260185460208301526019549081169282019290925274010000000000000000000000000000000000000000820463ffffffff166060820152780100000000000000000000000000000000000000000000000090910467ffffffffffffffff9081166080830152601a5480821660a08401526801000000000000000090041660c082015290565b6040516102fb9190600060e0820190506001600160a01b03808451168352602084015160208401528060408501511660408401525063ffffffff6060840151166060830152608083015167ffffffffffffffff80821660808501528060a08601511660a08501528060c08601511660c0850152505092915050565b6104fd6040518060400160405280601681526020017f45564d3245564d47454f666652616d7020312e302e300000000000000000000081525081565b6040516102fb9190614631565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016102fb565b6102aa61053d36600461450e565b610961565b6102aa610998565b6102aa610558366004614703565b6109aa565b61057061056b366004614767565b610bff565b6040516102fb91906147a4565b6005546001600160a01b0316610517565b60005460ff16610359565b6105a1610cc7565b6040516102fb91906147e8565b6102aa6105bc366004614829565b610d8c565b6102aa6110b2565b6105d161119a565b6040516102fb91906148a6565b601154600f546040805163ffffffff808516825264010000000090940490931660208401528201526060016102fb565b6102aa6111fc565b6105a161120c565b60005461010090046001600160a01b0316610517565b61063c61126c565b6040516102fb91908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6102aa61067d3660046148e9565b61130d565b6102aa610292366004614905565b6102aa61069e366004614a88565b611440565b6040805160018152600060208201819052918101919091526060016102fb565b600c546001600160a01b0316610517565b6102aa6106e2366004614c31565b61144b565b6105176106f536600461450e565b611a1e565b600d546001600160a01b0316610517565b61051761071936600461450e565b6001600160a01b039081166000908152600360205260409020541690565b6102aa610745366004614d16565b611b0c565b6102aa61075836600461450e565b611cde565b6102aa61076b36600461450e565b611d3d565b6102aa61077e366004614dc1565b611d74565b6102aa610791366004614829565b61262d565b6102aa6107a436600461450e565b612805565b3330146107e2576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a08101515115610811576107fa8160a00151612816565b61081181608001518260a001518360400151612a1e565b61081a81612abc565b50565b610825612b94565b6001600160a01b038116610865576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa158015610937573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095b9190614e9e565b15905090565b610969612b94565b6005805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6109a0612b94565b6109a8612bf3565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b0316141580156109e757506005546001600160a01b03163314155b15610a1e576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610a5a576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610ab4576006600060078381548110610a7f57610a7f614eb9565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610aad81614f17565b9050610a60565b5060005b82811015610be4576000858281518110610ad457610ad4614eb9565b6020026020010151905060006001600160a01b0316816001600160a01b031603610b2a576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610b3c57610b3c614eb9565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610ba157610ba1614eb9565b6020026020010151604051610bcb9291906001600160a01b03929092168252602082015260400190565b60405180910390a150610bdd81614f17565b9050610ab8565b508351610bf89060079060208701906140be565b5050505050565b80516060908067ffffffffffffffff811115610c1d57610c1d614160565b604051908082528060200260200182016040528015610c46578160200160208202803683370190505b50915060005b81811015610cc05760066000858381518110610c6a57610c6a614eb9565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610ca557610ca5614eb9565b6020908102919091010152610cb981614f17565b9050610c4c565b5050919050565b60045460609067ffffffffffffffff811115610ce557610ce5614160565b604051908082528060200260200182016040528015610d0e578160200160208202803683370190505b50905060005b600454811015610d8857610d4e60048281548110610d3457610d34614eb9565b6000918252602090912001546001600160a01b0316611a1e565b828281518110610d6057610d60614eb9565b6001600160a01b0390921660209283029190910190910152610d8181614f17565b9050610d14565b5090565b610d94612b94565b6004546000819003610dd2576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610e60576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614610eaf576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610ebe600185614f31565b81548110610ece57610ece614eb9565b9060005260206000200160009054906101000a90046001600160a01b03169050600482602001516bffffffffffffffffffffffff1681548110610f1357610f13614eb9565b6000918252602090912001546001600160a01b03166004610f35600186614f31565b81548110610f4557610f45614eb9565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483602001516bffffffffffffffffffffffff1681548110610f9957610f99614eb9565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000291909216179055600480548061102357611023614f48565b600082815260208082206000199084018101805473ffffffffffffffffffffffffffffffffffffffff191690559092019092556001600160a01b03878116808452600383526040808520949094558351908152908716918101919091527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146111115760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b606060148054806020026020016040519081016040528092919081815260200182805480156111f257602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116111d4575b5050505050905090565b611204612b94565b6109a8612c8f565b606060048054806020026020016040519081016040528092919081815260200182805480156111f2576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116111d4575050505050905090565b6112976040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b546060820181905242906000906112d19083614f31565b602084015184519192506112fd916112e99084614f77565b85604001516112f89190614f96565b612d17565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561134a57506005546001600160a01b03163314155b15611381576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff116113d5576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113df6008612d2d565b602081015160098190558151600855600a546113fb9190612d17565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b61081a816001612dda565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916114a191849163ffffffff851691908e908e908190840183828082843760009201919091525061308b92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a160408051606081018252600f5480825260105460ff8082166020850152610100909104169282019290925290831461155c5760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401611108565b61156a8b8b8b8b8b8b6130ae565b60007f0000000000000000000000000000000000000000000000000000000000000000156115c7576002826020015183604001516115a89190614fae565b6115b29190615002565b6115bd906001614fae565b60ff1690506115dd565b60208201516115d7906001614fae565b60ff1690505b88811461162c5760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401611108565b88871461167b5760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401611108565b3360009081526012602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156116be576116be614569565b60028111156116cf576116cf614569565b90525090506002816020015160028111156116ec576116ec614569565b14801561172657506014816000015160ff168154811061170e5761170e614eb9565b6000918252602090912001546001600160a01b031633145b6117725760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401611108565b505050505060008888604051611789929190615024565b6040519081900381206117a0918c90602001615034565b6040516020818303038152906040528051906020012090506117c061412c565b604080518082019091526000808252602082015260005b888110156119fc5760006001858884602081106117f6576117f6614eb9565b61180391901a601b614fae565b8d8d8681811061181557611815614eb9565b905060200201358c8c8781811061182e5761182e614eb9565b905060200201356040516000815260200160405260405161186b949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa15801561188d573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526012602090815290849020838501909452835460ff808216855292965092945084019161010090041660028111156118e2576118e2614569565b60028111156118f3576118f3614569565b905250925060018360200151600281111561191057611910614569565b1461195d5760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401611108565b8251849060ff16601f811061197457611974614eb9565b6020020151156119c65760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401611108565b600184846000015160ff16601f81106119e1576119e1614eb9565b91151560209092020152506119f581614f17565b90506117d7565b5050505063ffffffff8110611a1357611a13615050565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611a72576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611ae1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b05919061507f565b9392505050565b611b14612b94565b8051601780546001600160a01b0392831673ffffffffffffffffffffffffffffffffffffffff199091161790556020820151601855604080830151601980546060860151608087015167ffffffffffffffff90811678010000000000000000000000000000000000000000000000000277ffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff90931674010000000000000000000000000000000000000000027fffffffffffffffff000000000000000000000000000000000000000000000000909416959097169490941791909117169390931790925560a0830151601a805460c0860151851668010000000000000000027fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116929094169190911792909217909155517fd41dd65196ed6cf5fe4ec232c8b6f346c6db077c9e766c668a6862efc6ad488890611435908390600060e0820190506001600160a01b03808451168352602084015160208401528060408501511660408401525063ffffffff6060840151166060830152608083015167ffffffffffffffff80821660808501528060a08601511660a08501528060c08601511660c0850152505092915050565b611ce6612b94565b600c805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b611d45612b94565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b855185518560ff16601f831115611de7576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401611108565b80600003611e51576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401611108565b818314611edf576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401611108565b611eea816003614f77565b8311611f52576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401611108565b611f5a612b94565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b601354156120e157601354600090611fb290600190614f31565b9050600060138281548110611fc957611fc9614eb9565b6000918252602082200154601480546001600160a01b0390921693509084908110611ff657611ff6614eb9565b60009182526020808320909101546001600160a01b0385811684526012909252604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009081169091559290911680845292208054909116905560138054919250908061206957612069614f48565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff1916905501905560148054806120a9576120a9614f48565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff1916905501905550611f98915050565b60005b8151518110156124a6576000601260008460000151848151811061210a5761210a614eb9565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff16600281111561214757612147614569565b146121ae576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401611108565b6040805180820190915260ff821681526001602082015282518051601291600091859081106121df576121df614eb9565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561225557612255614569565b0217905550600091506122659050565b601260008460200151848151811061227f5761227f614eb9565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156122bc576122bc614569565b14612323576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401611108565b6040805180820190915260ff82168152602081016002815250601260008460200151848151811061235657612356614eb9565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff1982168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156123cc576123cc614569565b0217905550508251805160139250839081106123ea576123ea614eb9565b6020908102919091018101518254600181018455600093845292829020909201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909316929092179091558201518051601491908390811061244e5761244e614eb9565b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390921691909117905561249f81614f17565b90506120e4565b5060408101516010805460ff191660ff909216919091179055601180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff43811682029283178555908304811693600193909260009261251a92869290821691161761509c565b92506101000a81548163ffffffff021916908363ffffffff1602179055506125794630601160009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a0015161313e565b600f819055825180516010805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560115460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612618988b98919763ffffffff9092169690959194919391926150c4565b60405180910390a15050505050505050505050565b612635612b94565b6001600160a01b038216158061265257506001600160a01b038116155b15612689576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612718576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a519551909616740100000000000000000000000000000000000000000294909816939093179093558354600181018555939091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b909201805473ffffffffffffffffffffffffffffffffffffffff1916821790558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b61280d612b94565b61081a816131cb565b6000805b82518110156129155760006006600085848151811061283b5761283b614eb9565b6020026020010151600001516001600160a01b03166001600160a01b03168152602001908152602001600020549050806000036128ce5783828151811061288457612884614eb9565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401611108565b8382815181106128e0576128e0614eb9565b602002602001015160200151816128f79190614f77565b6129019084614f96565b9250508061290e90614f17565b905061281a565b508015612a1a576129266008612d2d565b600954811115612970576009546040517f688ccf77000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401611108565b600a548111156129d057600854600a546000919061298e9084614f31565b612998919061515a565b9050806040517fe31e0f3200000000000000000000000000000000000000000000000000000000815260040161110891815260200190565b80600860020160008282546129e59190614f31565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378906020016108c8565b5050565b8151835114612a59576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015612ab657612aa6848281518110612a7a57612a7a614eb9565b6020026020010151848381518110612a9457612a94614eb9565b60200260200101516020015184613287565b612aaf81614f17565b9050612a5c565b50505050565b60408101516001600160a01b03163b612ad25750565b600c546040517fda52b4c40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063da52b4c490612b1b9084906004016151b2565b6020604051808303816000875af1158015612b3a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b5e9190614e9e565b61081a576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005461010090046001600160a01b031633146109a85760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401611108565b60005460ff16612c455760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401611108565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60005460ff1615612ce25760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611108565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612c723390565b6000818310612d265781611b05565b5090919050565b6001810154600282015442911480612d485750808260030154145b15612d51575050565b816001015482600201541115612d93576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482612da59190614f31565b60018401548454919250612dcc91612dbd9084614f77565b85600201546112f89190614f96565b600284015550600390910155565b60005460ff1615612e2d5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611108565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ea49190614e9e565b15612eda576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005a600c549091506001600160a01b0316612f22576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606083015151828015612f3457508015155b15612f6b576040517f198753d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f7584846132fe565b8015612ffa5760195460608501516040517f45ef67060000000000000000000000000000000000000000000000000000000081526001600160a01b03909216916345ef670691612fc79160040161525c565b600060405180830381600087803b158015612fe157600080fd5b505af1158015612ff5573d6000803e3d6000fd5b505050505b82612ab657670de0b6b3a7640000846040015160008151811061301f5761301f614eb9565b60200260200101513a6017600101545a6130399087614f31565b6130439190614f96565b61304d9190614f77565b6130579190614f77565b613061919061515a565b3360009081526015602052604081208054909190613080908490614f96565b909155505050505050565b6130a9818060200190518101906130a29190615528565b6000612dda565b505050565b60006130bb826020614f77565b6130c6856020614f77565b6130d288610144614f96565b6130dc9190614f96565b6130e69190614f96565b6130f1906000614f96565b9050368114613135576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401611108565b50505050505050565b6000808a8a8a8a8a8a8a8a8a60405160200161316299989796959493929190615685565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b038216036132235760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401611108565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b1580156132ea57600080fd5b505af1158015613135573d6000803e3d6000fd5b60808201515180156130a95760008167ffffffffffffffff81111561332557613325614160565b60405190808252806020026020018201604052801561334e578160200160208202803683370190505b509050600061341c7f79d28ebf200d9eb053753c753cc1c6b1d6e0455a04272a9838f804c2f98e8389604080516020808201939093527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166080808301919091528251808303909101815260a0909101909152805191012090565b905060008367ffffffffffffffff81111561343957613439614160565b6040519080825280602002602001820160405280156134bf57816020015b60408051610160810182526000808252602080830182905292820181905260608083018290526080830182905260a0830182905260c0830182905260e08301829052610100830181905261012083015261014082015282526000199092019101816134575790505b50905060005b8481101561355a576000876080015182815181106134e5576134e5614eb9565b60200260200101518060200190518101906135009190615795565b905061350c818561392e565b85838151811061351e5761351e614eb9565b6020026020010181815250508083838151811061353d5761353d614eb9565b6020026020010181905250508061355390614f17565b90506134c5565b50600061357b848860a001518960c001518a60e001518b6101000151613a38565b5060195490915060009074010000000000000000000000000000000000000000900463ffffffff166135ad8342614f31565b11905060005b86811015611a135760008482815181106135cf576135cf614eb9565b602002602001015190506000613602826000015167ffffffffffffffff166000908152600e602052604090205460ff1690565b9050600281600381111561361857613618614569565b0361365e5781516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611108565b89156136bc5783806136815750600381600381111561367f5761367f614569565b145b6136b7576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613716565b60008160038111156136d0576136d0614569565b146137165781516040517f67d9ba0f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611108565b61371f82613b30565b600081600381111561373357613733614569565b036137e05760608201516040808401516001600160a01b031660009081526016602052205467ffffffffffffffff91821691613771911660016158b4565b67ffffffffffffffff16146137c45760608201516040517fb0241f4a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611108565b6137e06137d5836101400151613c67565b836020015130613287565b815167ffffffffffffffff166000908152600e60205260408120805460ff1916600117905561381661381184613cc9565b613f8a565b835167ffffffffffffffff166000908152600e602052604090208054919250829160ff1916600183600381111561384f5761384f614569565b02179055508260a0015180156138765750600381600381111561387457613874614569565b145b6138d4576040808401516001600160a01b031660009081526016602052908120805467ffffffffffffffff16916138ac836158d7565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b826000015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf826040516139129190614598565b60405180910390a25050508061392790614f17565b90506135b3565b60008060001b828460000151856060015186604001518760e001518861010001518051906020012089610120015160405160200161396c91906158fe565b604051602081830303815290604052805190602001208a608001518b60a001518c61014001518d60200151604051602001613a1a9c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600d546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce90613a95908c908c908c908c908c90600401615941565b6020604051808303816000875af1158015613ab4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ad89190615993565b905060008111613b14576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a613b209084614f31565b9350935050509550959350505050565b7f00000000000000000000000000000000000000000000000000000000000000008160c0015114613b95578060c001516040517fd44bc9eb00000000000000000000000000000000000000000000000000000000815260040161110891815260200190565b601a54610120820151516801000000000000000090910467ffffffffffffffff161015613bfd5780516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611108565b601a546101008201515167ffffffffffffffff909116101561081a57601a54610100820151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301526024820152604401611108565b6001600160a01b038181166000908152600360205260409020541680613cc4576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401611108565b919050565b613d126040518060e00160405280600081526020016060815260200160006001600160a01b03168152602001606081526020016060815260200160608152602001600081525090565b6101208201515160008167ffffffffffffffff811115613d3457613d34614160565b604051908082528060200260200182016040528015613d7957816020015b6040805180820190915260008082526020820152815260200190600190039081613d525790505b50905060008267ffffffffffffffff811115613d9757613d97614160565b604051908082528060200260200182016040528015613dc0578160200160208202803683370190505b50905060005b83811015613f00576000613dfb8761012001518381518110613dea57613dea614eb9565b602002602001015160000151613c67565b905080838381518110613e1057613e10614eb9565b60200260200101906001600160a01b031690816001600160a01b0316815250506040518060400160405280826001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613e79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e9d919061507f565b6001600160a01b031681526020018861012001518481518110613ec257613ec2614eb9565b602002602001015160200151815250848381518110613ee357613ee3614eb9565b60200260200101819052505080613ef990614f17565b9050613dc6565b506040518060e001604052808660c0015181526020018660400151604051602001613f3a91906001600160a01b0391909116815260200190565b60405160208183030381529060405281526020018660e001516001600160a01b03168152602001866101000151815260200182815260200183815260200186608001518152509350505050919050565b6040517f07a22a0700000000000000000000000000000000000000000000000000000000815260009030906307a22a0790613fc99085906004016151b2565b600060405180830381600087803b158015613fe357600080fd5b505af1925050508015613ff4575060015b6140b6573d808015614022576040519150601f19603f3d011682016040523d82523d6000602084013e614027565b606091505b50614031816159ac565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da800000000000000000000000000000000000000000000000000000000036140825750600392915050565b806040517fcf19edfd0000000000000000000000000000000000000000000000000000000081526004016111089190614631565b506002919050565b828054828255906000526020600020908101928215614120579160200282015b82811115614120578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039091161782556020909201916001909101906140de565b50610d8892915061414b565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610d88576000815560010161414c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156141b2576141b2614160565b60405290565b60405160e0810167ffffffffffffffff811182821017156141b2576141b2614160565b604051610120810167ffffffffffffffff811182821017156141b2576141b2614160565b604051610160810167ffffffffffffffff811182821017156141b2576141b2614160565b604051601f8201601f1916810167ffffffffffffffff8111828210171561424c5761424c614160565b604052919050565b600067ffffffffffffffff82111561426e5761426e614160565b50601f01601f191660200190565b600082601f83011261428d57600080fd5b81356142a061429b82614254565b614223565b8181528460208386010111156142b557600080fd5b816020850160208301376000918101602001919091529392505050565b6001600160a01b038116811461081a57600080fd5b8035613cc4816142d2565b600067ffffffffffffffff82111561430c5761430c614160565b5060051b60200190565b600082601f83011261432757600080fd5b8135602061433761429b836142f2565b82815260059290921b8401810191818101908684111561435657600080fd5b8286015b8481101561437a57803561436d816142d2565b835291830191830161435a565b509695505050505050565b600082601f83011261439657600080fd5b813560206143a661429b836142f2565b82815260069290921b840181019181810190868411156143c557600080fd5b8286015b8481101561437a57604081890312156143e25760008081fd5b6143ea61418f565b81356143f5816142d2565b815281850135858201528352918301916040016143c9565b60006020828403121561441f57600080fd5b813567ffffffffffffffff8082111561443757600080fd5b9083019060e0828603121561444b57600080fd5b6144536141b8565b8235815260208301358281111561446957600080fd5b6144758782860161427c565b602083015250614487604084016142e7565b604082015260608301358281111561449e57600080fd5b6144aa8782860161427c565b6060830152506080830135828111156144c257600080fd5b6144ce87828601614316565b60808301525060a0830135828111156144e657600080fd5b6144f287828601614385565b60a08301525060c083013560c082015280935050505092915050565b60006020828403121561452057600080fd5b8135611b05816142d2565b67ffffffffffffffff8116811461081a57600080fd5b8035613cc48161452b565b60006020828403121561455e57600080fd5b8135611b058161452b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60208101600483106145d3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b60005b838110156145f45781810151838201526020016145dc565b83811115612ab65750506000910152565b6000815180845261461d8160208601602086016145d9565b601f01601f19169290920160200192915050565b602081526000611b056020830184614605565b600082601f83011261465557600080fd5b8135602061466561429b836142f2565b82815260059290921b8401810191818101908684111561468457600080fd5b8286015b8481101561437a57803561469b816142d2565b8352918301918301614688565b600082601f8301126146b957600080fd5b813560206146c961429b836142f2565b82815260059290921b840181019181810190868411156146e857600080fd5b8286015b8481101561437a57803583529183019183016146ec565b6000806040838503121561471657600080fd5b823567ffffffffffffffff8082111561472e57600080fd5b61473a86838701614644565b9350602085013591508082111561475057600080fd5b5061475d858286016146a8565b9150509250929050565b60006020828403121561477957600080fd5b813567ffffffffffffffff81111561479057600080fd5b61479c84828501614644565b949350505050565b6020808252825182820181905260009190848201906040850190845b818110156147dc578351835292840192918401916001016147c0565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156147dc5783516001600160a01b031683529284019291840191600101614804565b6000806040838503121561483c57600080fd5b8235614847816142d2565b91506020830135614857816142d2565b809150509250929050565b600081518084526020808501945080840160005b8381101561489b5781516001600160a01b031687529582019590820190600101614876565b509495945050505050565b602081526000611b056020830184614862565b6000604082840312156148cb57600080fd5b6148d361418f565b9050813581526020820135602082015292915050565b6000604082840312156148fb57600080fd5b611b0583836148b9565b60006020828403121561491757600080fd5b813567ffffffffffffffff81111561492e57600080fd5b820160e08185031215611b0557600080fd5b600082601f83011261495157600080fd5b8135602061496161429b836142f2565b82815260059290921b8401810191818101908684111561498057600080fd5b8286015b8481101561437a5780356149978161452b565b8352918301918301614984565b600082601f8301126149b557600080fd5b813560206149c561429b836142f2565b82815260069290921b840181019181810190868411156149e457600080fd5b8286015b8481101561437a576149fa88826148b9565b8352918301916040016149e8565b600082601f830112614a1957600080fd5b81356020614a2961429b836142f2565b82815260059290921b84018101918181019086841115614a4857600080fd5b8286015b8481101561437a57803567ffffffffffffffff811115614a6c5760008081fd5b614a7a8986838b010161427c565b845250918301918301614a4c565b600060208284031215614a9a57600080fd5b813567ffffffffffffffff80821115614ab257600080fd5b908301906101208286031215614ac757600080fd5b614acf6141db565b823582811115614ade57600080fd5b614aea87828601614940565b825250602083013582811115614aff57600080fd5b614b0b87828601614316565b602083015250604083013582811115614b2357600080fd5b614b2f878286016146a8565b604083015250606083013582811115614b4757600080fd5b614b53878286016149a4565b606083015250608083013582811115614b6b57600080fd5b614b7787828601614a08565b60808301525060a083013582811115614b8f57600080fd5b614b9b878286016146a8565b60a08301525060c083013560c082015260e083013582811115614bbd57600080fd5b614bc9878286016146a8565b60e0830152506101009283013592810192909252509392505050565b60008083601f840112614bf757600080fd5b50813567ffffffffffffffff811115614c0f57600080fd5b6020830191508360208260051b8501011115614c2a57600080fd5b9250929050565b60008060008060008060008060e0898b031215614c4d57600080fd5b606089018a811115614c5e57600080fd5b8998503567ffffffffffffffff80821115614c7857600080fd5b818b0191508b601f830112614c8c57600080fd5b813581811115614c9b57600080fd5b8c6020828501011115614cad57600080fd5b6020830199508098505060808b0135915080821115614ccb57600080fd5b614cd78c838d01614be5565b909750955060a08b0135915080821115614cf057600080fd5b50614cfd8b828c01614be5565b999c989b50969995989497949560c00135949350505050565b600060e08284031215614d2857600080fd5b614d306141b8565b8235614d3b816142d2565b8152602083810135908201526040830135614d55816142d2565b6040820152606083013563ffffffff81168114614d7157600080fd5b6060820152614d8260808401614541565b6080820152614d9360a08401614541565b60a0820152614da460c08401614541565b60c08201529392505050565b803560ff81168114613cc457600080fd5b60008060008060008060c08789031215614dda57600080fd5b863567ffffffffffffffff80821115614df257600080fd5b614dfe8a838b01614316565b97506020890135915080821115614e1457600080fd5b614e208a838b01614316565b9650614e2e60408a01614db0565b95506060890135915080821115614e4457600080fd5b614e508a838b0161427c565b9450614e5e60808a01614541565b935060a0890135915080821115614e7457600080fd5b50614e8189828a0161427c565b9150509295509295509295565b80518015158114613cc457600080fd5b600060208284031215614eb057600080fd5b611b0582614e8e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006000198203614f2a57614f2a614ee8565b5060010190565b600082821015614f4357614f43614ee8565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000816000190483118215151615614f9157614f91614ee8565b500290565b60008219821115614fa957614fa9614ee8565b500190565b600060ff821660ff84168060ff03821115614fcb57614fcb614ee8565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600060ff83168061501557615015614fd3565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60006020828403121561509157600080fd5b8151611b05816142d2565b600063ffffffff8083168185168083038211156150bb576150bb614ee8565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526150f48184018a614862565b905082810360808401526151088189614862565b905060ff871660a084015282810360c08401526151258187614605565b905067ffffffffffffffff851660e084015282810361010084015261514a8185614605565b9c9b505050505050505050505050565b60008261516957615169614fd3565b500490565b600081518084526020808501945080840160005b8381101561489b57815180516001600160a01b031688528301518388015260409096019590820190600101615182565b60208152815160208201526000602083015160e060408401526151d9610100840182614605565b90506001600160a01b0360408501511660608401526060840151601f198085840301608086015261520a8383614605565b925060808601519150808584030160a08601526152278383614862565b925060a08601519150808584030160c086015250615245828261516e565b91505060c084015160e08401528091505092915050565b602080825282518282018190526000919060409081850190868401855b8281101561529e57815180518552860151868501529284019290850190600101615279565b5091979650505050505050565b8051613cc48161452b565b600082601f8301126152c757600080fd5b815160206152d761429b836142f2565b82815260059290921b840181019181810190868411156152f657600080fd5b8286015b8481101561437a57805161530d8161452b565b83529183019183016152fa565b8051613cc4816142d2565b600082601f83011261533657600080fd5b8151602061534661429b836142f2565b82815260059290921b8401810191818101908684111561536557600080fd5b8286015b8481101561437a57805161537c816142d2565b8352918301918301615369565b600082601f83011261539a57600080fd5b815160206153aa61429b836142f2565b82815260059290921b840181019181810190868411156153c957600080fd5b8286015b8481101561437a57805183529183019183016153cd565b600082601f8301126153f557600080fd5b8151602061540561429b836142f2565b82815260069290921b8401810191818101908684111561542457600080fd5b8286015b8481101561437a57604081890312156154415760008081fd5b61544961418f565b815181528482015185820152835291830191604001615428565b600082601f83011261547457600080fd5b815161548261429b82614254565b81815284602083860101111561549757600080fd5b61479c8260208301602087016145d9565b600082601f8301126154b957600080fd5b815160206154c961429b836142f2565b82815260059290921b840181019181810190868411156154e857600080fd5b8286015b8481101561437a57805167ffffffffffffffff81111561550c5760008081fd5b61551a8986838b0101615463565b8452509183019183016154ec565b60006020828403121561553a57600080fd5b815167ffffffffffffffff8082111561555257600080fd5b90830190610120828603121561556757600080fd5b61556f6141db565b82518281111561557e57600080fd5b61558a878286016152b6565b82525060208301518281111561559f57600080fd5b6155ab87828601615325565b6020830152506040830151828111156155c357600080fd5b6155cf87828601615389565b6040830152506060830151828111156155e757600080fd5b6155f3878286016153e4565b60608301525060808301518281111561560b57600080fd5b615617878286016154a8565b60808301525060a08301518281111561562f57600080fd5b61563b87828601615389565b60a08301525060c083015160c082015260e08301518281111561565d57600080fd5b61566987828601615389565b60e0830152506101009283015192810192909252509392505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526156bf8285018b614862565b915083820360808501526156d3828a614862565b915060ff881660a085015283820360c08501526156f08288614605565b90861660e0850152838103610100850152905061514a8185614605565b600082601f83011261571e57600080fd5b8151602061572e61429b836142f2565b82815260069290921b8401810191818101908684111561574d57600080fd5b8286015b8481101561437a576040818903121561576a5760008081fd5b61577261418f565b815161577d816142d2565b81528185015185820152835291830191604001615751565b6000602082840312156157a757600080fd5b815167ffffffffffffffff808211156157bf57600080fd5b9083019061016082860312156157d457600080fd5b6157dc6141ff565b6157e5836152ab565b8152602083015160208201526157fd6040840161531a565b604082015261580e606084016152ab565b60608201526080830151608082015261582960a08401614e8e565b60a082015260c083015160c082015261584460e0840161531a565b60e0820152610100808401518381111561585d57600080fd5b61586988828701615463565b828401525050610120808401518381111561588357600080fd5b61588f8882870161570d565b82840152505061014091506158a582840161531a565b91810191909152949350505050565b600067ffffffffffffffff8083168185168083038211156150bb576150bb614ee8565b600067ffffffffffffffff8083168181036158f4576158f4614ee8565b6001019392505050565b602081526000611b05602083018461516e565b600081518084526020808501945080840160005b8381101561489b57815187529582019590820190600101615925565b60a08152600061595460a0830188615911565b82810360208401526159668188615911565b905085604084015282810360608401526159808186615911565b9150508260808301529695505050505050565b6000602082840312156159a557600080fd5b5051919050565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156159f45780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
}

var EVM2EVMGEOffRampABI = EVM2EVMGEOffRampMetaData.ABI

var EVM2EVMGEOffRampBin = EVM2EVMGEOffRampMetaData.Bin

func DeployEVM2EVMGEOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig EVM2EVMGEOffRampInterfaceGEOffRampConfig, onRampAddress common.Address, commitStore common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address) (common.Address, *types.Transaction, *EVM2EVMGEOffRamp, error) {
	parsed, err := EVM2EVMGEOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMGEOffRampBin), backend, sourceChainId, chainId, offRampConfig, onRampAddress, commitStore, afn, sourceTokens, pools, rateLimiterConfig, tokenLimitsAdmin)
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMGEOffRamp.Contract.CcipReceive(&_EVM2EVMGEOffRamp.CallOpts, arg0)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
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

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetPool(&_EVM2EVMGEOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetPool(&_EVM2EVMGEOffRamp.CallOpts, sourceToken)
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "executeSingleMessage", message)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMGEOffRamp.TransactOpts, message)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMGEOffRamp.TransactOpts, message)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) ManuallyExecute(opts *bind.TransactOpts, report CCIPExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "manuallyExecute", report)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) ManuallyExecute(report CCIPExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.ManuallyExecute(&_EVM2EVMGEOffRamp.TransactOpts, report)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) ManuallyExecute(report CCIPExecutionReport) (*types.Transaction, error) {
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
	State          uint8
	Raw            types.Log
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMGEOffRampExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMGEOffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMGEOffRampExecutionStateChangedIterator{contract: _EVM2EVMGEOffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMGEOffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
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
	Router common.Address
	Raw    types.Log
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
	return common.HexToHash("0x06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf")
}

func (EVM2EVMGEOffRampGEOffRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0xd41dd65196ed6cf5fe4ec232c8b6f346c6db077c9e766c668a6862efc6ad4888")
}

func (EVM2EVMGEOffRampOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1a9ff9caccd597226ff8c393ab44e6b57669905d9b16f1cdb2ac267253dbf27c")
}

func (EVM2EVMGEOffRampOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
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

	CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

		error)

	GetCommitStore(opts *bind.CallOpts) (common.Address, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetGEConfig(opts *bind.CallOpts) (EVM2EVMGEOffRampInterfaceGEOffRampConfig, error)

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

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMGEOffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMGEOffRampExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error)

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
