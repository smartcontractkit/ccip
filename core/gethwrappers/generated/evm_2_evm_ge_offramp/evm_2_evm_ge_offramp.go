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

type CommonEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
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

type IBaseOffRampOffRampConfig struct {
	PermissionLessExecutionThresholdSeconds uint32
	ExecutionDelaySeconds                   uint64
	MaxDataSize                             uint64
	MaxTokensLength                         uint64
}

type IEVM2EVMGEOffRampGEOffRampConfig struct {
	GasOverhead                             *big.Int
	GasFeeCache                             common.Address
	PermissionLessExecutionThresholdSeconds uint32
	ExecutionDelaySeconds                   uint64
	MaxDataSize                             uint64
	MaxTokensLength                         uint64
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasFeeCache\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMGEOffRamp.GEOffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractICommitStore\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"seqNum\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedGasPriceUpdate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasFeeCache\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMGEOffRamp.GEOffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"GEOffRampConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIBaseOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitStore\",\"outputs\":[{\"internalType\":\"contractICommitStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGEConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasFeeCache\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMGEOffRamp.GEOffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"}],\"name\":\"getNopBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolByDestToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractIAny2EVMOffRampRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"linkPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structGE.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structGE.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICommitStore\",\"name\":\"commitStore\",\"type\":\"address\"}],\"name\":\"setCommitStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasFeeCache\",\"name\":\"gasFeeCache\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMGEOffRamp.GEOffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setGEConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAny2EVMOffRampRouter\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b50604051620067ce380380620067ce8339810160408190526200003591620008e8565b6000805460ff191681556001908c908c908b908b908b908b908b908b908b90829082908690869089903390819081620000b55760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ef57620000ef8162000540565b5050506001600160a01b0381166200011a57604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015c5760405162d8548360e71b815260040160405180910390fd5b81516200017190600590602085019062000657565b5060005b8251811015620002fa5760006040518060400160405280848481518110620001a157620001a162000a0b565b60200260200101516001600160a01b03168152602001836001600160601b031681525090508060036000868581518110620001e057620001e062000a0b565b6020908102919091018101516001600160a01b0390811683528282019390935260409091016000908120845194909201516001600160601b0316600160a01b02939092169290921790915581518451909160049186908690811062000249576200024962000a0b565b60200260200101516001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200028f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002b5919062000a21565b6001600160a01b039081168252602082019290925260400160002080546001600160a01b0319169290911691909117905550620002f28162000a48565b905062000175565b5050600680546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600955600a93909355600b55600c91909155871662000377576040516342bcdf7f60e11b815260040160405180910390fd5b886001600160401b03166080816001600160401b031681525050876001600160401b031660a0816001600160401b031681525050866001600160a01b031660c0816001600160a01b03168152505085600e60006101000a8154816001600160a01b0302191690836001600160a01b0316021790555050505050505050505080151560e081151581525050508860186000820151816000015560208201518160010160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555060608201518160010160186101000a8154816001600160401b0302191690836001600160401b0316021790555060808201518160020160006101000a8154816001600160401b0302191690836001600160401b0316021790555060a08201518160020160086101000a8154816001600160401b0302191690836001600160401b031602179055509050506200051e7fba22a5847647789e6efe1840c86bc66129ac89e03d7b95e0eebdf7fa43763fdd620005f160201b60201c565b610100526001600160a01b0316610120525062000a7098505050505050505050565b336001600160a01b038216036200059a5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ac565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008160805160a05160c0516040516020016200063a94939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b828054828255906000526020600020908101928215620006af579160200282015b82811115620006af57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000678565b50620006bd929150620006c1565b5090565b5b80821115620006bd5760008155600101620006c2565b80516001600160401b0381168114620006f057600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146200072157600080fd5b50565b600060c082840312156200073757600080fd5b60405160c081016001600160401b03811182821017156200075c576200075c620006f5565b806040525080915082518152602083015162000778816200070b565b6020820152604083015163ffffffff811681146200079557600080fd5b6040820152620007a860608401620006d8565b6060820152620007bb60808401620006d8565b6080820152620007ce60a08401620006d8565b60a08201525092915050565b8051620006f0816200070b565b600082601f830112620007f957600080fd5b815160206001600160401b0380831115620008185762000818620006f5565b8260051b604051601f19603f83011681018181108482111715620008405762000840620006f5565b6040529384528581018301938381019250878511156200085f57600080fd5b83870191505b848210156200088b5781516200087b816200070b565b8352918301919083019062000865565b979650505050505050565b600060408284031215620008a957600080fd5b604080519081016001600160401b0381118282101715620008ce57620008ce620006f5565b604052825181526020928301519281019290925250919050565b60008060008060008060008060008060006102208c8e0312156200090b57600080fd5b620009168c620006d8565b9a506200092660208d01620006d8565b9950620009378d60408e0162000724565b9850620009486101008d01620007da565b9750620009596101208d01620007da565b96506200096a6101408d01620007da565b6101608d01519096506001600160401b038111156200098857600080fd5b620009968e828f01620007e7565b6101808e015190965090506001600160401b03811115620009b657600080fd5b620009c48e828f01620007e7565b945050620009d78d6101a08e0162000896565b9250620009e86101e08d01620007da565b9150620009f96102008d01620007da565b90509295989b509295989b9093969950565b634e487b7160e01b600052603260045260246000fd5b60006020828403121562000a3457600080fd5b815162000a41816200070b565b9392505050565b60006001820162000a6957634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e0516101005161012051615d0062000ace6000396000613b71015260006136250152600061186a01526000611e6f0152600061030c0152600081816102e701528181611e4b0152613dbb0152615d006000f3fe608060405234801561001057600080fd5b50600436106102d35760003560e01c806381ff704811610186578063b4069b31116100e3578063d1d8a83d11610097578063ea03d1b511610071578063ea03d1b514610850578063eb511dd41461085e578063f2fde38b1461087157600080fd5b8063d1d8a83d146107da578063d7e2bb5014610811578063e3d0e7121461083d57600080fd5b8063c0d78655116100c8578063c0d78655146107a1578063c9033284146107b4578063ca911034146107c757600080fd5b8063b4069b311461077d578063b66f0efb1461079057600080fd5b806390c2339b1161013a578063afcb95d71161011f578063afcb95d714610739578063b0f479a114610759578063b1dc65a41461076a57600080fd5b806390c2339b146106eb578063918725431461072657600080fd5b8063856c82471161016b578063856c82471461068157806389c06568146106cd5780638da5cb5b146106d557600080fd5b806381ff7048146106495780638456cb591461067957600080fd5b80634352fa9f11610234578063681fba16116101e857806379ba5097116101cd57806379ba5097146106195780637f738dc814610621578063814118341461063457600080fd5b8063681fba16146105f1578063744b92e21461060657600080fd5b8063599f643111610219578063599f6431146105a95780635c975abb146105ba5780635d86f141146105c557600080fd5b80634352fa9f146105765780634741062e1461058957600080fd5b80631628b6a71161028b5780632222dd42116102705780632222dd421461053657806339aa92641461055b5780633f4ba83a1461056e57600080fd5b80631628b6a7146103b6578063181f5a77146104ed57600080fd5b8063142a98fc116102bc578063142a98fc14610351578063147809b31461038b57806315fcd8c1146103a357600080fd5b8063087ae6df146102d8578063108ee5fc1461033c575b600080fd5b6040805167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811682527f0000000000000000000000000000000000000000000000000000000000000000166020820152015b60405180910390f35b61034f61034a366004614434565b610884565b005b61037e61035f366004614472565b67ffffffffffffffff166000908152600f602052604090205460ff1690565b60405161033391906144a5565b61039361093b565b6040519015158152602001610333565b61034f6103b13660046145a8565b6109c8565b6104806040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260185481526019546001600160a01b038116602083015274010000000000000000000000000000000000000000810463ffffffff1692820192909252780100000000000000000000000000000000000000000000000090910467ffffffffffffffff9081166060830152601a5480821660808401526801000000000000000090041660a082015290565b6040516103339190600060c082019050825182526001600160a01b03602084015116602083015263ffffffff6040840151166040830152606083015167ffffffffffffffff80821660608501528060808601511660808501528060a08601511660a0850152505092915050565b6105296040518060400160405280601681526020017f45564d3245564d47454f666652616d7020312e302e300000000000000000000081525081565b60405161033391906146a7565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610333565b61034f610569366004614434565b610b6d565b61034f610ba4565b61034f6105843660046147b8565b610bb6565b61059c61059736600461481c565b610e0b565b6040516103339190614859565b6006546001600160a01b0316610543565b60005460ff16610393565b6105436105d3366004614434565b6001600160a01b039081166000908152600360205260409020541690565b6105f9610ed3565b604051610333919061489d565b61034f6106143660046148de565b610f98565b61034f61134a565b61034f61062f366004614a95565b611432565b61063c6114a8565b6040516103339190614bf1565b6012546010546040805163ffffffff80851682526401000000009094049093166020840152820152606001610333565b61034f61150a565b6106b461068f366004614434565b6001600160a01b031660009081526017602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff9091168152602001610333565b6105f961151a565b60005461010090046001600160a01b0316610543565b6106f361157a565b60405161033391908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b61034f610734366004614c04565b61161b565b604080516001815260006020820181905291810191909152606001610333565b600d546001600160a01b0316610543565b61034f610778366004614c9f565b611747565b61054361078b366004614434565b611d1a565b600e546001600160a01b0316610543565b61034f6107af366004614434565b611e08565b61034f6107c2366004614434565b611ec6565b61034f6107d5366004614f17565b611efd565b6108036107e8366004614434565b6001600160a01b031660009081526016602052604090205490565b604051908152602001610333565b61054361081f366004614434565b6001600160a01b039081166000908152600460205260409020541690565b61034f61084b366004615085565b611f0b565b61034f6102d3366004615152565b61034f61086c3660046148de565b61276d565b61034f61087f366004614434565b6129e9565b61088c6129fa565b6001600160a01b0381166108cc576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa15801561099e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c29190615198565b15905090565b6109d06129fa565b8051601855602081015160198054604080850151606086015167ffffffffffffffff90811678010000000000000000000000000000000000000000000000000277ffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff90931674010000000000000000000000000000000000000000027fffffffffffffffff0000000000000000000000000000000000000000000000009095166001600160a01b03909716969096179390931716939093179091556080830151601a805460a0860151841668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009091169290931691909117919091179055517fda126dc369452deb5c9e98a0522d5a2312155d7d64f0b41fa77adcddf459850190610b62908390600060c082019050825182526001600160a01b03602084015116602083015263ffffffff6040840151166040830152606083015167ffffffffffffffff80821660608501528060808601511660808501528060a08601511660a0850152505092915050565b60405180910390a150565b610b756129fa565b6006805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b610bac6129fa565b610bb4612a59565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015610bf357506006546001600160a01b03163314155b15610c2a576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610c66576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085460005b81811015610cc0576007600060088381548110610c8b57610c8b6151b5565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610cb9816151e1565b9050610c6c565b5060005b82811015610df0576000858281518110610ce057610ce06151b5565b6020026020010151905060006001600160a01b0316816001600160a01b031603610d36576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610d4857610d486151b5565b602002602001015160076000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610dad57610dad6151b5565b6020026020010151604051610dd79291906001600160a01b03929092168252602082015260400190565b60405180910390a150610de9816151e1565b9050610cc4565b508351610e0490600890602087019061437d565b5050505050565b80516060908067ffffffffffffffff811115610e2957610e296144cd565b604051908082528060200260200182016040528015610e52578160200160208202803683370190505b50915060005b81811015610ecc5760076000858381518110610e7657610e766151b5565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610eb157610eb16151b5565b6020908102919091010152610ec5816151e1565b9050610e58565b5050919050565b60055460609067ffffffffffffffff811115610ef157610ef16144cd565b604051908082528060200260200182016040528015610f1a578160200160208202803683370190505b50905060005b600554811015610f9457610f5a60058281548110610f4057610f406151b5565b6000918252602090912001546001600160a01b0316611d1a565b828281518110610f6c57610f6c6151b5565b6001600160a01b0390921660209283029190910190910152610f8d816151e1565b9050610f20565b5090565b610fa06129fa565b6005546000819003610fde576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529061106c576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b0316146110bb576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060056110ca6001856151fb565b815481106110da576110da6151b5565b9060005260206000200160009054906101000a90046001600160a01b03169050600582602001516bffffffffffffffffffffffff168154811061111f5761111f6151b5565b6000918252602090912001546001600160a01b031660056111416001866151fb565b81548110611151576111516151b5565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600583602001516bffffffffffffffffffffffff16815481106111a5576111a56151b5565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000291909216179055600580548061122f5761122f615212565b6001900381819060005260206000200160006101000a8154906001600160a01b030219169055905560046000856001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611299573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112bd9190615228565b6001600160a01b03908116825260208083019390935260409182016000908120805473ffffffffffffffffffffffffffffffffffffffff1916905588821680825260038552838220919091558251908152908716928101929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146113a95760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b33301461146b576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a0820151511561149a576114838260a00151612af5565b61149a82608001518360a001518460400151612cf9565b6114a48282612d97565b5050565b6060601580548060200260200160405190810160405280929190818152602001828054801561150057602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116114e2575b5050505050905090565b6115126129fa565b610bb4612e72565b60606005805480602002602001604051908101604052809291908181526020018280548015611500576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116114e2575050505050905090565b6115a56040518060800160405280600081526020016000815260200160008152602001600081525090565b604080516080810182526009548152600a546020820152600b5491810191909152600c546060820181905242906000906115df90836151fb565b6020840151845191925061160b916115f79084615245565b85604001516116069190615264565b612efa565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561165857506006546001600160a01b03163314155b1561168f576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff116116e3576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6116ed6009612f14565b6020810151600a8190558151600955600b546117099190612efa565b600b55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d9101610b62565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161179d91849163ffffffff851691908e908e9081908401838280828437600092019190915250612fc192505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260105480825260115460ff808216602085015261010090910416928201929092529083146118585760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d61746368000000000000000000000060448201526064016113a0565b6118668b8b8b8b8b8b612fe4565b60007f0000000000000000000000000000000000000000000000000000000000000000156118c3576002826020015183604001516118a4919061527c565b6118ae91906152b7565b6118b990600161527c565b60ff1690506118d9565b60208201516118d390600161527c565b60ff1690505b8881146119285760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e61747572657300000000000060448201526064016113a0565b8887146119775760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e000060448201526064016113a0565b3360009081526013602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156119ba576119ba61448f565b60028111156119cb576119cb61448f565b90525090506002816020015160028111156119e8576119e861448f565b148015611a2257506015816000015160ff1681548110611a0a57611a0a6151b5565b6000918252602090912001546001600160a01b031633145b611a6e5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d6974746572000000000000000060448201526064016113a0565b505050505060008888604051611a859291906152d9565b604051908190038120611a9c918c906020016152e9565b604051602081830303815290604052805190602001209050611abc6143eb565b604080518082019091526000808252602082015260005b88811015611cf8576000600185888460208110611af257611af26151b5565b611aff91901a601b61527c565b8d8d86818110611b1157611b116151b5565b905060200201358c8c87818110611b2a57611b2a6151b5565b9050602002013560405160008152602001604052604051611b67949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611b89573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526013602090815290849020838501909452835460ff80821685529296509294508401916101009004166002811115611bde57611bde61448f565b6002811115611bef57611bef61448f565b9052509250600183602001516002811115611c0c57611c0c61448f565b14611c595760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e000060448201526064016113a0565b8251849060ff16601f8110611c7057611c706151b5565b602002015115611cc25760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e617475726500000000000000000000000060448201526064016113a0565b600184846000015160ff16601f8110611cdd57611cdd6151b5565b9115156020909202015250611cf1816151e1565b9050611ad3565b5050505063ffffffff8110611d0f57611d0f615305565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611d6e576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611ddd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e019190615228565b9392505050565b611e106129fa565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038381169182179092556040805167ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681527f0000000000000000000000000000000000000000000000000000000000000000909316602084015290917f052b5907be1d3ac35d571862117562e80ee743c01251e388dafb7dc4e92a726c910160405180910390a250565b611ece6129fa565b600e805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b611f08816001613074565b50565b855185518560ff16601f831115611f7e576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064016113a0565b80600003611fe8576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064016113a0565b818314612076576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084016113a0565b612081816003615245565b83116120e9576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f2068696768000000000000000060448201526064016113a0565b6120f16129fa565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b6014541561225b57601454600090612149906001906151fb565b9050600060148281548110612160576121606151b5565b6000918252602082200154601580546001600160a01b039092169350908490811061218d5761218d6151b5565b60009182526020808320909101546001600160a01b03858116845260139092526040808420805461ffff19908116909155929091168084529220805490911690556014805491925090806121e3576121e3615212565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff19169055019055601580548061222357612223615212565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff191690550190555061212f915050565b60005b8151518110156125e65760006013600084600001518481518110612284576122846151b5565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156122c1576122c161448f565b14612328576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e6572206164647265737300000000000000000060448201526064016113a0565b6040805180820190915260ff82168152600160208201528251805160139160009185908110612359576123596151b5565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff1916176101008360028111156123b2576123b261448f565b0217905550600091506123c29050565b60136000846020015184815181106123dc576123dc6151b5565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156124195761241961448f565b14612480576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016113a0565b6040805180820190915260ff8216815260208101600281525060136000846020015184815181106124b3576124b36151b5565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff19161761010083600281111561250c5761250c61448f565b02179055505082518051601492508390811061252a5761252a6151b5565b6020908102919091018101518254600181018455600093845292829020909201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909316929092179091558201518051601591908390811061258e5761258e6151b5565b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039092169190911790556125df816151e1565b905061225e565b5060408101516011805460ff191660ff909216919091179055601280547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff43811682029283178555908304811693600193909260009261265a92869290821691161761531b565b92506101000a81548163ffffffff021916908363ffffffff1602179055506126b94630601260009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613318565b6010819055825180516011805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560125460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612758988b98919763ffffffff909216969095919491939192615343565b60405180910390a15050505050505050505050565b6127756129fa565b6001600160a01b038216158061279257506001600160a01b038116155b156127c9576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612858576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038083168083526005546bffffffffffffffffffffffff90811660208086019182528785166000908152600382526040808220885194519095167401000000000000000000000000000000000000000002939096169290921790925583517f21df0da70000000000000000000000000000000000000000000000000000000081529351869460049492936321df0da79282870192819003870181865afa15801561290d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129319190615228565b6001600160a01b03908116825260208083019390935260409182016000908120805495831673ffffffffffffffffffffffffffffffffffffffff199687161790556005805460018101825591527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001805488831695168517905581519384528516918301919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b6129f16129fa565b611f08816133a5565b60005461010090046001600160a01b03163314610bb45760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016113a0565b60005460ff16612aab5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016113a0565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6000805b8251811015612bf457600060076000858481518110612b1a57612b1a6151b5565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600003612bad57838281518110612b6357612b636151b5565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201526024016113a0565b838281518110612bbf57612bbf6151b5565b60200260200101516020015181612bd69190615245565b612be09084615264565b92505080612bed906151e1565b9050612af9565b5080156114a457612c056009612f14565b600a54811115612c4f57600a546040517f688ccf770000000000000000000000000000000000000000000000000000000081526004810191909152602481018290526044016113a0565b600b54811115612caf57600954600b5460009190612c6d90846151fb565b612c7791906153d9565b9050806040517fe31e0f320000000000000000000000000000000000000000000000000000000081526004016113a091815260200190565b8060096002016000828254612cc491906151fb565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba3789060200161092f565b8151835114612d34576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015612d9157612d81848281518110612d5557612d556151b5565b6020026020010151848381518110612d6f57612d6f6151b5565b60200260200101516020015184613461565b612d8a816151e1565b9050612d37565b50505050565b60408201516001600160a01b03163b612dae575050565b600d546040517facd754d40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063acd754d490612df99085908590600401615431565b6020604051808303816000875af1158015612e18573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e3c9190615198565b6114a4576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005460ff1615612ec55760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016113a0565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612ad83390565b6000818310612f095781612f0b565b825b90505b92915050565b6001810154600282015442911480612f2f5750808260030154145b15612f38575050565b816001015482600201541115612f7a576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482612f8c91906151fb565b60018401548454919250612fb391612fa49084615245565b85600201546116069190615264565b600284015550600390910155565b612fdf81806020019051810190612fd891906157a0565b6000613074565b505050565b6000612ff1826020615245565b612ffc856020615245565b61300888610144615264565b6130129190615264565b61301c9190615264565b613027906000615264565b905036811461306b576040517f8e1192e1000000000000000000000000000000000000000000000000000000008152600481018290523660248201526044016113a0565b50505050505050565b60005460ff16156130c75760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016113a0565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561311a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061313e9190615198565b15613174576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005a600d549091506001600160a01b03166131bc576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060830151511561327e5781156131ff576040517f198753d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60195460608401516040517f371708310000000000000000000000000000000000000000000000000000000081526001600160a01b039092169163371708319161324b916004016158fd565b600060405180830381600087803b15801561326557600080fd5b505af1158015613279573d6000803e3d6000fd5b505050505b61328883836134d8565b81612fdf57670de0b6b3a764000083604001516000815181106132ad576132ad6151b5565b60200260200101513a6018600001545a6132c790866151fb565b6132d19190615264565b6132db9190615245565b6132e59190615245565b6132ef91906153d9565b336000908152601660205260408120805490919061330e908490615264565b9091555050505050565b6000808a8a8a8a8a8a8a8a8a60405160200161333c99989796959493929190615968565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b038216036133fd5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016113a0565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b1580156134c457600080fd5b505af115801561306b573d6000803e3d6000fd5b60808201515160008190036134ec57505050565b60008167ffffffffffffffff811115613507576135076144cd565b604051908082528060200260200182016040528015613530578160200160208202803683370190505b50905060008267ffffffffffffffff81111561354e5761354e6144cd565b6040519080825280602002602001820160405280156135dc57816020015b60408051610180810182526000808252602080830182905292820181905260608083018290526080830182905260a0830182905260c0830182905260e0830182905261010083018190526101208301526101408201819052610160820152825260001990920191018161356c5790505b50905060005b8381101561369757600086608001518281518110613602576136026151b5565b602002602001015180602001905181019061361d9190615a78565b9050613649817f0000000000000000000000000000000000000000000000000000000000000000613bb7565b84838151811061365b5761365b6151b5565b6020026020010181815250508083838151811061367a5761367a6151b5565b60200260200101819052505080613690906151e1565b90506135e2565b5060006136b8838760a001518860c001518960e001518a6101000151613cc1565b5060195490915060009074010000000000000000000000000000000000000000900463ffffffff166136ea83426151fb565b1190506000805b86811015613b6657600085828151811061370d5761370d6151b5565b602002602001015190506000613740826020015167ffffffffffffffff166000908152600f602052604090205460ff1690565b905060028160038111156137565761375661448f565b0361379f5760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016113a0565b89156137fd5784806137c2575060038160038111156137c0576137c061448f565b145b6137f8576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61385a565b60008160038111156138115761381161448f565b1461385a5760208201516040517f67d9ba0f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016113a0565b600081600381111561386e5761386e61448f565b0361391f57608082015160608301516001600160a01b031660009081526017602052604090205467ffffffffffffffff918216916138ae91166001615bab565b67ffffffffffffffff161461390d5781606001516001600160a01b0316826080015167ffffffffffffffff167fd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf4123760405160405180910390a35050613b56565b604082015161391c9085615264565b93505b61392882613db9565b60208281015167ffffffffffffffff166000908152600f90915260408120805460ff1916600117905561396361395d84613f0e565b8c6141e3565b60208085015167ffffffffffffffff166000908152600f909152604090208054919250829160ff191660018360038111156139a0576139a061448f565b02179055508a15613a86578260c0015180156139cd575060038260038111156139cb576139cb61448f565b145b80156139ea575060028160038111156139e8576139e861448f565b145b80613a2257506000826003811115613a0457613a0461448f565b148015613a2257506002816003811115613a2057613a2061448f565b145b15613a815760608301516001600160a01b03166000908152601760205260408120805467ffffffffffffffff1691613a5983615bce565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b613b06565b8260c001518015613aa857506003816003811115613aa657613aa661448f565b145b613b065760608301516001600160a01b03166000908152601760205260408120805467ffffffffffffffff1691613ade83615bce565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b826101600151836020015167ffffffffffffffff167f5e04b4755a4460aa6de4f3a906c4324a025c7449c02b52f5466659b5bfdfba5f83604051613b4a91906144a5565b60405180910390a35050505b613b5f816151e1565b90506136f1565b506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116600090815260046020526040902054613bad91168230613461565b5050505050505050565b60008060001b828460200151856080015186606001518760e0015188610100015180519060200120896101200151604051602001613bf59190615bf5565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d60400151604051602001613ca39c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600e546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce90613d1e908c908c908c908c908c90600401615c38565b6020604051808303816000875af1158015613d3d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d619190615c8a565b905060008111613d9d576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a613da990846151fb565b9350935050509550959350505050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16816000015167ffffffffffffffff1614613e395780516040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016113a0565b601a54610120820151516801000000000000000090910467ffffffffffffffff161015613ea45760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016113a0565b601a546101008201515167ffffffffffffffff9091161015611f0857601a54610100820151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909216600483015260248201526044016113a0565b613f616040518060e00160405280600067ffffffffffffffff1681526020016060815260200160006001600160a01b03168152602001606081526020016060815260200160608152602001600081525090565b6101208201515160008167ffffffffffffffff811115613f8357613f836144cd565b604051908082528060200260200182016040528015613fc857816020015b6040805180820190915260008082526020820152815260200190600190039081613fa15790505b50905060008267ffffffffffffffff811115613fe657613fe66144cd565b60405190808252806020026020018201604052801561400f578160200160208202803683370190505b50905060005b8381101561414f57600061404a8761012001518381518110614039576140396151b5565b60200260200101516000015161431b565b90508083838151811061405f5761405f6151b5565b60200260200101906001600160a01b031690816001600160a01b0316815250506040518060400160405280826001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156140c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906140ec9190615228565b6001600160a01b031681526020018861012001518481518110614111576141116151b5565b602002602001015160200151815250848381518110614132576141326151b5565b60200260200101819052505080614148906151e1565b9050614015565b506040518060e00160405280866000015167ffffffffffffffff168152602001866060015160405160200161419391906001600160a01b0391909116815260200190565b60405160208183030381529060405281526020018660e001516001600160a01b0316815260200186610100015181526020018281526020018381526020018660a001518152509350505050919050565b6040517f7f738dc80000000000000000000000000000000000000000000000000000000081526000903090637f738dc8906142249086908690600401615431565b600060405180830381600087803b15801561423e57600080fd5b505af192505050801561424f575060015b614312573d80801561427d576040519150601f19603f3d011682016040523d82523d6000602084013e614282565b606091505b5061428c81615ca3565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da800000000000000000000000000000000000000000000000000000000036142de576003915050612f0e565b806040517fcf19edfd0000000000000000000000000000000000000000000000000000000081526004016113a091906146a7565b50600292915050565b6001600160a01b038181166000908152600360205260409020541680614378576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024016113a0565b919050565b8280548282559060005260206000209081019282156143df579160200282015b828111156143df578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390911617825560209092019160019091019061439d565b50610f9492915061440a565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610f94576000815560010161440b565b6001600160a01b0381168114611f0857600080fd5b60006020828403121561444657600080fd5b8135611e018161441f565b67ffffffffffffffff81168114611f0857600080fd5b803561437881614451565b60006020828403121561448457600080fd5b8135611e0181614451565b634e487b7160e01b600052602160045260246000fd5b60208101600483106144c757634e487b7160e01b600052602160045260246000fd5b91905290565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715614506576145066144cd565b60405290565b60405160e0810167ffffffffffffffff81118282101715614506576145066144cd565b604051610120810167ffffffffffffffff81118282101715614506576145066144cd565b604051610180810167ffffffffffffffff81118282101715614506576145066144cd565b604051601f8201601f1916810167ffffffffffffffff811182821017156145a0576145a06144cd565b604052919050565b600060c082840312156145ba57600080fd5b60405160c0810181811067ffffffffffffffff821117156145dd576145dd6144cd565b6040528235815260208301356145f28161441f565b6020820152604083013563ffffffff8116811461460e57600080fd5b6040820152606083013561462181614451565b606082015261463260808401614467565b608082015261464360a08401614467565b60a08201529392505050565b60005b8381101561466a578181015183820152602001614652565b83811115612d915750506000910152565b6000815180845261469381602086016020860161464f565b601f01601f19169290920160200192915050565b602081526000612f0b602083018461467b565b80356143788161441f565b600067ffffffffffffffff8211156146df576146df6144cd565b5060051b60200190565b600082601f8301126146fa57600080fd5b8135602061470f61470a836146c5565b614577565b82815260059290921b8401810191818101908684111561472e57600080fd5b8286015b848110156147525780356147458161441f565b8352918301918301614732565b509695505050505050565b600082601f83011261476e57600080fd5b8135602061477e61470a836146c5565b82815260059290921b8401810191818101908684111561479d57600080fd5b8286015b8481101561475257803583529183019183016147a1565b600080604083850312156147cb57600080fd5b823567ffffffffffffffff808211156147e357600080fd5b6147ef868387016146e9565b9350602085013591508082111561480557600080fd5b506148128582860161475d565b9150509250929050565b60006020828403121561482e57600080fd5b813567ffffffffffffffff81111561484557600080fd5b614851848285016146e9565b949350505050565b6020808252825182820181905260009190848201906040850190845b8181101561489157835183529284019291840191600101614875565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156148915783516001600160a01b0316835292840192918401916001016148b9565b600080604083850312156148f157600080fd5b82356148fc8161441f565b9150602083013561490c8161441f565b809150509250929050565b600067ffffffffffffffff821115614931576149316144cd565b50601f01601f191660200190565b600082601f83011261495057600080fd5b813561495e61470a82614917565b81815284602083860101111561497357600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126149a157600080fd5b813560206149b161470a836146c5565b82815260059290921b840181019181810190868411156149d057600080fd5b8286015b848110156147525780356149e78161441f565b83529183019183016149d4565b600082601f830112614a0557600080fd5b81356020614a1561470a836146c5565b82815260069290921b84018101918181019086841115614a3457600080fd5b8286015b848110156147525760408189031215614a515760008081fd5b614a596144e3565b8135614a648161441f565b81528185013585820152835291830191604001614a38565b8015158114611f0857600080fd5b803561437881614a7c565b60008060408385031215614aa857600080fd5b823567ffffffffffffffff80821115614ac057600080fd5b9084019060e08287031215614ad457600080fd5b614adc61450c565b614ae583614467565b8152602083013582811115614af957600080fd5b614b058882860161493f565b602083015250614b17604084016146ba565b6040820152606083013582811115614b2e57600080fd5b614b3a8882860161493f565b606083015250608083013582811115614b5257600080fd5b614b5e88828601614990565b60808301525060a083013582811115614b7657600080fd5b614b82888286016149f4565b60a08301525060c083013560c0820152809450505050614ba460208401614a8a565b90509250929050565b600081518084526020808501945080840160005b83811015614be65781516001600160a01b031687529582019590820190600101614bc1565b509495945050505050565b602081526000612f0b6020830184614bad565b600060408284031215614c1657600080fd5b6040516040810181811067ffffffffffffffff82111715614c3957614c396144cd565b604052823581526020928301359281019290925250919050565b60008083601f840112614c6557600080fd5b50813567ffffffffffffffff811115614c7d57600080fd5b6020830191508360208260051b8501011115614c9857600080fd5b9250929050565b60008060008060008060008060e0898b031215614cbb57600080fd5b606089018a811115614ccc57600080fd5b8998503567ffffffffffffffff80821115614ce657600080fd5b818b0191508b601f830112614cfa57600080fd5b813581811115614d0957600080fd5b8c6020828501011115614d1b57600080fd5b6020830199508098505060808b0135915080821115614d3957600080fd5b614d458c838d01614c53565b909750955060a08b0135915080821115614d5e57600080fd5b50614d6b8b828c01614c53565b999c989b50969995989497949560c00135949350505050565b600082601f830112614d9557600080fd5b81356020614da561470a836146c5565b82815260059290921b84018101918181019086841115614dc457600080fd5b8286015b84811015614752578035614ddb81614451565b8352918301918301614dc8565b6fffffffffffffffffffffffffffffffff81168114611f0857600080fd5b600082601f830112614e1757600080fd5b81356020614e2761470a836146c5565b82815260069290921b84018101918181019086841115614e4657600080fd5b8286015b848110156147525760408189031215614e635760008081fd5b614e6b6144e3565b8135614e7681614451565b815281850135614e8581614de8565b81860152835291830191604001614e4a565b600082601f830112614ea857600080fd5b81356020614eb861470a836146c5565b82815260059290921b84018101918181019086841115614ed757600080fd5b8286015b8481101561475257803567ffffffffffffffff811115614efb5760008081fd5b614f098986838b010161493f565b845250918301918301614edb565b600060208284031215614f2957600080fd5b813567ffffffffffffffff80821115614f4157600080fd5b908301906101208286031215614f5657600080fd5b614f5e61452f565b823582811115614f6d57600080fd5b614f7987828601614d84565b825250602083013582811115614f8e57600080fd5b614f9a87828601614990565b602083015250604083013582811115614fb257600080fd5b614fbe8782860161475d565b604083015250606083013582811115614fd657600080fd5b614fe287828601614e06565b606083015250608083013582811115614ffa57600080fd5b61500687828601614e97565b60808301525060a08301358281111561501e57600080fd5b61502a8782860161475d565b60a08301525060c083013560c082015260e08301358281111561504c57600080fd5b6150588782860161475d565b60e0830152506101009283013592810192909252509392505050565b803560ff8116811461437857600080fd5b60008060008060008060c0878903121561509e57600080fd5b863567ffffffffffffffff808211156150b657600080fd5b6150c28a838b01614990565b975060208901359150808211156150d857600080fd5b6150e48a838b01614990565b96506150f260408a01615074565b9550606089013591508082111561510857600080fd5b6151148a838b0161493f565b945061512260808a01614467565b935060a089013591508082111561513857600080fd5b5061514589828a0161493f565b9150509295509295509295565b60006020828403121561516457600080fd5b813567ffffffffffffffff81111561517b57600080fd5b820160e08185031215611e0157600080fd5b805161437881614a7c565b6000602082840312156151aa57600080fd5b8151611e0181614a7c565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060001982036151f4576151f46151cb565b5060010190565b60008282101561520d5761520d6151cb565b500390565b634e487b7160e01b600052603160045260246000fd5b60006020828403121561523a57600080fd5b8151611e018161441f565b600081600019048311821515161561525f5761525f6151cb565b500290565b60008219821115615277576152776151cb565b500190565b600060ff821660ff84168060ff03821115615299576152996151cb565b019392505050565b634e487b7160e01b600052601260045260246000fd5b600060ff8316806152ca576152ca6152a1565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b634e487b7160e01b600052600160045260246000fd5b600063ffffffff80831681851680830382111561533a5761533a6151cb565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526153738184018a614bad565b905082810360808401526153878189614bad565b905060ff871660a084015282810360c08401526153a4818761467b565b905067ffffffffffffffff851660e08401528281036101008401526153c9818561467b565b9c9b505050505050505050505050565b6000826153e8576153e86152a1565b500490565b600081518084526020808501945080840160005b83811015614be657815180516001600160a01b031688528301518388015260409096019590820190600101615401565b6040815267ffffffffffffffff83511660408201526000602084015160e0606084015261546261012084018261467b565b9050604085015161547e60808501826001600160a01b03169052565b5060608501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808584030160a08601526154b9838361467b565b925060808701519150808584030160c08601526154d68383614bad565b925060a08701519150808584030160e0860152506154f482826153ed565b60c087015161010086015285151560208601529250611e01915050565b805161437881614451565b600082601f83011261552d57600080fd5b8151602061553d61470a836146c5565b82815260059290921b8401810191818101908684111561555c57600080fd5b8286015b8481101561475257805161557381614451565b8352918301918301615560565b80516143788161441f565b600082601f83011261559c57600080fd5b815160206155ac61470a836146c5565b82815260059290921b840181019181810190868411156155cb57600080fd5b8286015b848110156147525780516155e28161441f565b83529183019183016155cf565b600082601f83011261560057600080fd5b8151602061561061470a836146c5565b82815260059290921b8401810191818101908684111561562f57600080fd5b8286015b848110156147525780518352918301918301615633565b600082601f83011261565b57600080fd5b8151602061566b61470a836146c5565b82815260069290921b8401810191818101908684111561568a57600080fd5b8286015b8481101561475257604081890312156156a75760008081fd5b6156af6144e3565b81516156ba81614451565b8152818501516156c981614de8565b8186015283529183019160400161568e565b600082601f8301126156ec57600080fd5b81516156fa61470a82614917565b81815284602083860101111561570f57600080fd5b61485182602083016020870161464f565b600082601f83011261573157600080fd5b8151602061574161470a836146c5565b82815260059290921b8401810191818101908684111561576057600080fd5b8286015b8481101561475257805167ffffffffffffffff8111156157845760008081fd5b6157928986838b01016156db565b845250918301918301615764565b6000602082840312156157b257600080fd5b815167ffffffffffffffff808211156157ca57600080fd5b9083019061012082860312156157df57600080fd5b6157e761452f565b8251828111156157f657600080fd5b6158028782860161551c565b82525060208301518281111561581757600080fd5b6158238782860161558b565b60208301525060408301518281111561583b57600080fd5b615847878286016155ef565b60408301525060608301518281111561585f57600080fd5b61586b8782860161564a565b60608301525060808301518281111561588357600080fd5b61588f87828601615720565b60808301525060a0830151828111156158a757600080fd5b6158b3878286016155ef565b60a08301525060c083015160c082015260e0830151828111156158d557600080fd5b6158e1878286016155ef565b60e0830152506101009283015192810192909252509392505050565b602080825282518282018190526000919060409081850190868401855b8281101561595b578151805167ffffffffffffffff1685528601516fffffffffffffffffffffffffffffffff1686850152928401929085019060010161591a565b5091979650505050505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526159a28285018b614bad565b915083820360808501526159b6828a614bad565b915060ff881660a085015283820360c08501526159d3828861467b565b90861660e085015283810361010085015290506153c9818561467b565b600082601f830112615a0157600080fd5b81516020615a1161470a836146c5565b82815260069290921b84018101918181019086841115615a3057600080fd5b8286015b848110156147525760408189031215615a4d5760008081fd5b615a556144e3565b8151615a608161441f565b81528185015185820152835291830191604001615a34565b600060208284031215615a8a57600080fd5b815167ffffffffffffffff80821115615aa257600080fd5b908301906101808286031215615ab757600080fd5b615abf614553565b615ac883615511565b8152615ad660208401615511565b602082015260408301516040820152615af160608401615580565b6060820152615b0260808401615511565b608082015260a083015160a0820152615b1d60c0840161518d565b60c0820152615b2e60e08401615580565b60e08201526101008084015183811115615b4757600080fd5b615b53888287016156db565b8284015250506101208084015183811115615b6d57600080fd5b615b79888287016159f0565b8284015250506101409150615b8f828401615580565b9181019190915261016091820151918101919091529392505050565b600067ffffffffffffffff80831681851680830382111561533a5761533a6151cb565b600067ffffffffffffffff808316818103615beb57615beb6151cb565b6001019392505050565b602081526000612f0b60208301846153ed565b600081518084526020808501945080840160005b83811015614be657815187529582019590820190600101615c1c565b60a081526000615c4b60a0830188615c08565b8281036020840152615c5d8188615c08565b90508560408401528281036060840152615c778186615c08565b9150508260808301529695505050505050565b600060208284031215615c9c57600080fd5b5051919050565b6000815160208301517fffffffff0000000000000000000000000000000000000000000000000000000080821693506004831015615ceb5780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
}

var EVM2EVMGEOffRampABI = EVM2EVMGEOffRampMetaData.ABI

var EVM2EVMGEOffRampBin = EVM2EVMGEOffRampMetaData.Bin

func DeployEVM2EVMGEOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId uint64, chainId uint64, offRampConfig IEVM2EVMGEOffRampGEOffRampConfig, onRampAddress common.Address, commitStore common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig IAggregateRateLimiterRateLimiterConfig, tokenLimitsAdmin common.Address, feeToken common.Address) (common.Address, *types.Transaction, *EVM2EVMGEOffRamp, error) {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(IAggregateRateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(IAggregateRateLimiterTokenBucket)).(*IAggregateRateLimiterTokenBucket)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMGEOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetGEConfig(opts *bind.CallOpts) (IEVM2EVMGEOffRampGEOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getGEConfig")

	if err != nil {
		return *new(IEVM2EVMGEOffRampGEOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IEVM2EVMGEOffRampGEOffRampConfig)).(*IEVM2EVMGEOffRampGEOffRampConfig)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetGEConfig() (IEVM2EVMGEOffRampGEOffRampConfig, error) {
	return _EVM2EVMGEOffRamp.Contract.GetGEConfig(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetGEConfig() (IEVM2EVMGEOffRampGEOffRampConfig, error) {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMMessageFromSender, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "executeSingleMessage", message, manualExecution)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) ExecuteSingleMessage(message InternalAny2EVMMessageFromSender, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMGEOffRamp.TransactOpts, message, manualExecution)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) ExecuteSingleMessage(message InternalAny2EVMMessageFromSender, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMGEOffRamp.TransactOpts, message, manualExecution)
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) SetGEConfig(opts *bind.TransactOpts, config IEVM2EVMGEOffRampGEOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "setGEConfig", config)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) SetGEConfig(config IEVM2EVMGEOffRampGEOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetGEConfig(&_EVM2EVMGEOffRamp.TransactOpts, config)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) SetGEConfig(config IEVM2EVMGEOffRampGEOffRampConfig) (*types.Transaction, error) {
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMGEOffRamp.TransactOpts, config)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
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
	Arg0 IEVM2EVMGEOffRampGEOffRampConfig
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
	Config IBaseOffRampOffRampConfig
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
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error)

	CcipReceive(opts *bind.CallOpts, arg0 InternalAny2EVMMessageFromSender) error

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

		error)

	GetCommitStore(opts *bind.CallOpts) (common.Address, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetGEConfig(opts *bind.CallOpts) (IEVM2EVMGEOffRampGEOffRampConfig, error)

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

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMMessageFromSender, manualExecution bool) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report GEExecutionReport) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetCommitStore(opts *bind.TransactOpts, commitStore common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetGEConfig(opts *bind.TransactOpts, config IEVM2EVMGEOffRampGEOffRampConfig) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error)

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
