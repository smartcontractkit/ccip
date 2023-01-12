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
	Token          common.Address
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
	FeeManager                              common.Address
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractIFeeManager\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMGEOffRamp.GEOffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractICommitStore\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"seqNum\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedGasPriceUpdate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractIFeeManager\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMGEOffRamp.GEOffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"GEOffRampConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIBaseOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitStore\",\"outputs\":[{\"internalType\":\"contractICommitStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGEConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractIFeeManager\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMGEOffRamp.GEOffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nop\",\"type\":\"address\"}],\"name\":\"getNopBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolByDestToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractIAny2EVMOffRampRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"linkPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structGE.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structGE.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICommitStore\",\"name\":\"commitStore\",\"type\":\"address\"}],\"name\":\"setCommitStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasOverhead\",\"type\":\"uint256\"},{\"internalType\":\"contractIFeeManager\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMGEOffRamp.GEOffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setGEConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAny2EVMOffRampRouter\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b506040516200648e3803806200648e8339810160408190526200003591620008d7565b6000805460ff191681558b908b908a908a908a908a908a908a908a90829082908690869089903390819081620000b25760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ec57620000ec816200052f565b5050506001600160a01b0381166200011757604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03929092169190911790558051825114620001595760405162d8548360e71b815260040160405180910390fd5b81516200016e90600590602085019062000646565b5060005b8251811015620002f757600060405180604001604052808484815181106200019e576200019e620009fa565b60200260200101516001600160a01b03168152602001836001600160601b031681525090508060036000868581518110620001dd57620001dd620009fa565b6020908102919091018101516001600160a01b0390811683528282019390935260409091016000908120845194909201516001600160601b0316600160a01b029390921692909217909155815184519091600491869086908110620002465762000246620009fa565b60200260200101516001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200028c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002b2919062000a10565b6001600160a01b039081168252602082019290925260400160002080546001600160a01b0319169290911691909117905550620002ef8162000a37565b905062000172565b5050600680546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600955600a93909355600b55600c91909155871662000374576040516342bcdf7f60e11b815260040160405180910390fd5b886001600160401b03166080816001600160401b031681525050876001600160401b031660a0816001600160401b031681525050866001600160a01b031660c0816001600160a01b03168152505085600e60006101000a8154816001600160a01b0302191690836001600160a01b031602179055505050505050505050508860186000820151816000015560208201518160010160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555060608201518160010160186101000a8154816001600160401b0302191690836001600160401b0316021790555060808201518160020160006101000a8154816001600160401b0302191690836001600160401b0316021790555060a08201518160020160086101000a8154816001600160401b0302191690836001600160401b031602179055509050506200050e7fba22a5847647789e6efe1840c86bc66129ac89e03d7b95e0eebdf7fa43763fdd620005e060201b60201c565b60e0526001600160a01b0316610100525062000a5f98505050505050505050565b336001600160a01b03821603620005895760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a9565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008160805160a05160c0516040516020016200062994939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b8280548282559060005260206000209081019282156200069e579160200282015b828111156200069e57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000667565b50620006ac929150620006b0565b5090565b5b80821115620006ac5760008155600101620006b1565b80516001600160401b0381168114620006df57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146200071057600080fd5b50565b600060c082840312156200072657600080fd5b60405160c081016001600160401b03811182821017156200074b576200074b620006e4565b80604052508091508251815260208301516200076781620006fa565b6020820152604083015163ffffffff811681146200078457600080fd5b60408201526200079760608401620006c7565b6060820152620007aa60808401620006c7565b6080820152620007bd60a08401620006c7565b60a08201525092915050565b8051620006df81620006fa565b600082601f830112620007e857600080fd5b815160206001600160401b0380831115620008075762000807620006e4565b8260051b604051601f19603f830116810181811084821117156200082f576200082f620006e4565b6040529384528581018301938381019250878511156200084e57600080fd5b83870191505b848210156200087a5781516200086a81620006fa565b8352918301919083019062000854565b979650505050505050565b6000604082840312156200089857600080fd5b604080519081016001600160401b0381118282101715620008bd57620008bd620006e4565b604052825181526020928301519281019290925250919050565b60008060008060008060008060008060006102208c8e031215620008fa57600080fd5b620009058c620006c7565b9a506200091560208d01620006c7565b9950620009268d60408e0162000713565b9850620009376101008d01620007c9565b9750620009486101208d01620007c9565b9650620009596101408d01620007c9565b6101608d01519096506001600160401b038111156200097757600080fd5b620009858e828f01620007d6565b6101808e015190965090506001600160401b03811115620009a557600080fd5b620009b38e828f01620007d6565b945050620009c68d6101a08e0162000885565b9250620009d76101e08d01620007c9565b9150620009e86102008d01620007c9565b90509295989b509295989b9093969950565b634e487b7160e01b600052603260045260246000fd5b60006020828403121562000a2357600080fd5b815162000a3081620006fa565b9392505050565b60006001820162000a5857634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e051610100516159dc62000ab260003960006137e80152600061329c01526000612461015260006103210152600081816102fc0152818161243d0152613ab201526159dc6000f3fe608060405234801561001057600080fd5b50600436106102d35760003560e01c806379ba509711610186578063b0f479a1116100e3578063c903328411610097578063ea03d1b511610071578063ea03d1b514610850578063eb511dd41461085e578063f2fde38b1461087157600080fd5b8063c9033284146107da578063d1d8a83d146107ed578063d7e2bb501461082457600080fd5b8063b4069b31116100c8578063b4069b31146107a3578063b66f0efb146107b6578063c0d78655146107c757600080fd5b8063b0f479a11461077f578063b1dc65a41461079057600080fd5b806389c065681161013a57806390c2339b1161011f57806390c2339b14610711578063918725431461074c578063afcb95d71461075f57600080fd5b806389c06568146106f35780638da5cb5b146106fb57600080fd5b806381ff70481161016b57806381ff70481461066f5780638456cb591461069f578063856c8247146106a757600080fd5b806379ba5097146106545780637f738dc81461065c57600080fd5b806339aa9264116102345780635c975abb116101e8578063666cab8d116101cd578063666cab8d14610617578063681fba161461062c578063744b92e21461064157600080fd5b80635c975abb146105e05780635d86f141146105eb57600080fd5b80634352fa9f116102195780634352fa9f1461059c5780634741062e146105af578063599f6431146105cf57600080fd5b806339aa9264146105815780633f4ba83a1461059457600080fd5b806315fcd8c11161028b578063181f5a7711610270578063181f5a77146105005780631ef38174146105495780632222dd421461055c57600080fd5b806315fcd8c1146103b65780631628b6a7146103c957600080fd5b8063108ee5fc116102bc578063108ee5fc14610351578063142a98fc14610364578063147809b31461039e57600080fd5b8063031c5c65146102d8578063087ae6df146102ed575b600080fd5b6102eb6102e6366004614572565b610884565b005b6040805167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811682527f0000000000000000000000000000000000000000000000000000000000000000166020820152015b60405180910390f35b6102eb61035f3660046146cf565b610892565b6103916103723660046146ec565b67ffffffffffffffff166000908152600f602052604090205460ff1690565b604051610348919061471f565b6103a6610949565b6040519015158152602001610348565b6102eb6103c4366004614747565b6109d6565b6104936040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260185481526019546001600160a01b038116602083015274010000000000000000000000000000000000000000810463ffffffff1692820192909252780100000000000000000000000000000000000000000000000090910467ffffffffffffffff9081166060830152601a5480821660808401526801000000000000000090041660a082015290565b6040516103489190600060c082019050825182526001600160a01b03602084015116602083015263ffffffff6040840151166040830152606083015167ffffffffffffffff80821660608501528060808601511660808501528060a08601511660a0850152505092915050565b61053c6040518060400160405280601681526020017f45564d3245564d47454f666652616d7020312e302e300000000000000000000081525081565b6040516103489190614846565b6102eb61055736600461486a565b610b7b565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610348565b6102eb61058f3660046146cf565b6111e1565b6102eb611218565b6102eb6105aa366004614937565b61122a565b6105c26105bd36600461499b565b61147f565b60405161034891906149d8565b6006546001600160a01b0316610569565b60005460ff166103a6565b6105696105f93660046146cf565b6001600160a01b039081166000908152600360205260409020541690565b61061f611547565b6040516103489190614a60565b6106346115a9565b6040516103489190614a73565b6102eb61064f366004614ab4565b61166e565b6102eb611a20565b6102eb61066a366004614b8e565b611b03565b6012546010546040805163ffffffff80851682526401000000009094049093166020840152820152606001610348565b6102eb611b79565b6106da6106b53660046146cf565b6001600160a01b031660009081526017602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff9091168152602001610348565b610634611b89565b60005461010090046001600160a01b0316610569565b610719611be9565b60405161034891908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6102eb61075a366004614ca6565b611c8a565b604080516001815260006020820181905291810191909152606001610348565b600d546001600160a01b0316610569565b6102eb61079e366004614d41565b611db6565b6105696107b13660046146cf565b61230c565b600e546001600160a01b0316610569565b6102eb6107d53660046146cf565b6123fa565b6102eb6107e83660046146cf565b6124b8565b6108166107fb3660046146cf565b6001600160a01b031660009081526016602052604090205490565b604051908152602001610348565b6105696108323660046146cf565b6001600160a01b039081166000908152600460205260409020541690565b6102eb6102d3366004614e26565b6102eb61086c366004614ab4565b6124ef565b6102eb61087f3660046146cf565b61276b565b61088f81600161277c565b50565b61089a612a21565b6001600160a01b0381166108da576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa1580156109ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d09190614e6c565b15905090565b6109de612a21565b8051601855602081015160198054604080850151606086015167ffffffffffffffff90811678010000000000000000000000000000000000000000000000000277ffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff90931674010000000000000000000000000000000000000000027fffffffffffffffff0000000000000000000000000000000000000000000000009095166001600160a01b03909716969096179390931716939093179091556080830151601a805460a0860151841668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009091169290931691909117919091179055517fda126dc369452deb5c9e98a0522d5a2312155d7d64f0b41fa77adcddf459850190610b70908390600060c082019050825182526001600160a01b03602084015116602083015263ffffffff6040840151166040830152606083015167ffffffffffffffff80821660608501528060808601511660808501528060a08601511660a0850152505092915050565b60405180910390a150565b855185518560ff16601f831115610bf3576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064015b60405180910390fd5b80600003610c5d576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610bea565b818314610ceb576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610bea565b610cf6816003614e9f565b8311610d5e576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610bea565b610d66612a21565b60145460005b81811015610e0e576013600060148381548110610d8b57610d8b614ebe565b60009182526020808320909101546001600160a01b031683528201929092526040018120805461ffff1916905560158054601392919084908110610dd157610dd1614ebe565b60009182526020808320909101546001600160a01b031683528201929092526040019020805461ffff19169055610e0781614ed4565b9050610d6c565b50895160005b818110156110a35760008c8281518110610e3057610e30614ebe565b6020026020010151905060006002811115610e4d57610e4d614709565b6001600160a01b038216600090815260136020526040902054610100900460ff166002811115610e7f57610e7f614709565b14610ee6576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610bea565b6040805180820190915260ff8316815260208101600190526001600160a01b03821660009081526013602090815260409091208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610f4e57610f4e614709565b021790555090505060008c8381518110610f6a57610f6a614ebe565b6020026020010151905060006002811115610f8757610f87614709565b6001600160a01b038216600090815260136020526040902054610100900460ff166002811115610fb957610fb9614709565b14611020576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610bea565b6040805180820190915260ff8416815260208101600290526001600160a01b03821660009081526013602090815260409091208251815460ff90911660ff19821681178355928401519192839161ffff19161761010083600281111561108857611088614709565b021790555090505050508061109c90614ed4565b9050610e14565b508a516110b79060149060208e0190614074565b5089516110cb9060159060208d0190614074565b506011805460ff8381166101000261ffff19909216908c1617179055601280546111349146913091906000906111069063ffffffff16614eee565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168e8e8e8e8e8e612a80565b6010600001819055506000601260049054906101000a900463ffffffff16905043601260046101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581601060000154601260009054906101000a900463ffffffff168f8f8f8f8f8f6040516111cb99989796959493929190614f11565b60405180910390a1505050505050505050505050565b6111e9612a21565b6006805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b611220612a21565b611228612b0d565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561126757506006546001600160a01b03163314155b1561129e576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151815181146112da576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085460005b818110156113345760076000600883815481106112ff576112ff614ebe565b60009182526020808320909101546001600160a01b0316835282019290925260400181205561132d81614ed4565b90506112e0565b5060005b8281101561146457600085828151811061135457611354614ebe565b6020026020010151905060006001600160a01b0316816001600160a01b0316036113aa576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8482815181106113bc576113bc614ebe565b602002602001015160076000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f8186848151811061142157611421614ebe565b602002602001015160405161144b9291906001600160a01b03929092168252602082015260400190565b60405180910390a15061145d81614ed4565b9050611338565b508351611478906008906020870190614074565b5050505050565b80516060908067ffffffffffffffff81111561149d5761149d614116565b6040519080825280602002602001820160405280156114c6578160200160208202803683370190505b50915060005b8181101561154057600760008583815181106114ea576114ea614ebe565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000205483828151811061152557611525614ebe565b602090810291909101015261153981614ed4565b90506114cc565b5050919050565b6060601580548060200260200160405190810160405280929190818152602001828054801561159f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611581575b5050505050905090565b60055460609067ffffffffffffffff8111156115c7576115c7614116565b6040519080825280602002602001820160405280156115f0578160200160208202803683370190505b50905060005b60055481101561166a576116306005828154811061161657611616614ebe565b6000918252602090912001546001600160a01b031661230c565b82828151811061164257611642614ebe565b6001600160a01b039092166020928302919091019091015261166381614ed4565b90506115f6565b5090565b611676612a21565b60055460008190036116b4576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290611742576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614611791576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060056117a0600185614fa7565b815481106117b0576117b0614ebe565b9060005260206000200160009054906101000a90046001600160a01b03169050600582602001516bffffffffffffffffffffffff16815481106117f5576117f5614ebe565b6000918252602090912001546001600160a01b03166005611817600186614fa7565b8154811061182757611827614ebe565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600583602001516bffffffffffffffffffffffff168154811061187b5761187b614ebe565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000291909216179055600580548061190557611905614fbe565b6001900381819060005260206000200160006101000a8154906001600160a01b030219169055905560046000856001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561196f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119939190614fd4565b6001600160a01b03908116825260208083019390935260409182016000908120805473ffffffffffffffffffffffffffffffffffffffff1916905588821680825260038552838220919091558251908152908716928101929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b03163314611a7a5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610bea565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b333014611b3c576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a08201515115611b6b57611b548260a00151612ba9565b611b6b82608001518360a001518460400151612dad565b611b758282612e4b565b5050565b611b81612a21565b611228612f26565b6060600580548060200260200160405190810160405280929190818152602001828054801561159f576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611581575050505050905090565b611c146040518060800160405280600081526020016000815260200160008152602001600081525090565b604080516080810182526009548152600a546020820152600b5491810191909152600c54606082018190524290600090611c4e9083614fa7565b60208401518451919250611c7a91611c669084614e9f565b8560400151611c759190614ff1565b612fae565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015611cc757506006546001600160a01b03163314155b15611cfe576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff11611d52576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611d5c6009612fc8565b6020810151600a8190558151600955600b54611d789190612fae565b600b55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d9101610b70565b60005a9050611dfa88888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061307592505050565b6040805160608101825260105480825260115460ff808216602085015261010090910416928201929092528a35918214611e6d5780516040517f93df584c000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610bea565b6040805183815260208d81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a1600281602001518260400151611ec89190615009565b611ed29190615044565b611edd906001615009565b60ff168714611f18576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b868514611f51576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526013602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115611f9457611f94614709565b6002811115611fa557611fa5614709565b9052509050600281602001516002811115611fc257611fc2614709565b148015611ffc57506015816000015160ff1681548110611fe457611fe4614ebe565b6000918252602090912001546001600160a01b031633145b612032576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506000612040866020614e9f565b61204b896020614e9f565b6120578c610144614ff1565b6120619190614ff1565b61206b9190614ff1565b90503681146120af576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401610bea565b5060008a8a6040516120c2929190615066565b6040519081900381206120d9918e90602001615076565b6040516020818303038152906040528051906020012090506120f96140e2565b8860005b818110156122fb5760006001858a846020811061211c5761211c614ebe565b61212991901a601b615009565b8f8f8681811061213b5761213b614ebe565b905060200201358e8e8781811061215457612154614ebe565b9050602002013560405160008152602001604052604051612191949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156121b3573d6000803e3d6000fd5b505060408051601f198101516001600160a01b038116600090815260136020908152848220848601909552845460ff808216865293975091955092939284019161010090910416600281111561220b5761220b614709565b600281111561221c5761221c614709565b905250905060018160200151600281111561223957612239614709565b14612270576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051859060ff16601f811061228757612287614ebe565b6020020151156122c3576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600185826000015160ff16601f81106122de576122de614ebe565b91151560209092020152506122f4905081614ed4565b90506120fd565b505050505050505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680612360576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa1580156123cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123f39190614fd4565b9392505050565b612402612a21565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038381169182179092556040805167ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681527f0000000000000000000000000000000000000000000000000000000000000000909316602084015290917f052b5907be1d3ac35d571862117562e80ee743c01251e388dafb7dc4e92a726c910160405180910390a250565b6124c0612a21565b600e805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6124f7612a21565b6001600160a01b038216158061251457506001600160a01b038116155b1561254b576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156125da576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038083168083526005546bffffffffffffffffffffffff90811660208086019182528785166000908152600382526040808220885194519095167401000000000000000000000000000000000000000002939096169290921790925583517f21df0da70000000000000000000000000000000000000000000000000000000081529351869460049492936321df0da79282870192819003870181865afa15801561268f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126b39190614fd4565b6001600160a01b03908116825260208083019390935260409182016000908120805495831673ffffffffffffffffffffffffffffffffffffffff199687161790556005805460018101825591527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001805488831695168517905581519384528516918301919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612773612a21565b61088f81613093565b60005460ff16156127cf5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610bea565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612822573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128469190614e6c565b1561287c576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005a600d549091506001600160a01b03166128c4576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608301515115612986578115612907576040517f198753d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60195460608401516040517f9086658e0000000000000000000000000000000000000000000000000000000081526001600160a01b0390921691639086658e9161295391600401615092565b600060405180830381600087803b15801561296d57600080fd5b505af1158015612981573d6000803e3d6000fd5b505050505b612990838361314f565b81612a1c57670de0b6b3a764000083604001516000815181106129b5576129b5614ebe565b60200260200101513a6018600001545a6129cf9086614fa7565b6129d99190614ff1565b6129e39190614e9f565b6129ed9190614e9f565b6129f79190615110565b3360009081526016602052604081208054909190612a16908490614ff1565b90915550505b505050565b60005461010090046001600160a01b031633146112285760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610bea565b6000808a8a8a8a8a8a8a8a8a604051602001612aa499989796959493929190615124565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b60005460ff16612b5f5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610bea565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6000805b8251811015612ca857600060076000858481518110612bce57612bce614ebe565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600003612c6157838281518110612c1757612c17614ebe565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610bea565b838281518110612c7357612c73614ebe565b60200260200101516020015181612c8a9190614e9f565b612c949084614ff1565b92505080612ca190614ed4565b9050612bad565b508015611b7557612cb96009612fc8565b600a54811115612d0357600a546040517f688ccf77000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610bea565b600b54811115612d6357600954600b5460009190612d219084614fa7565b612d2b9190615110565b9050806040517fe31e0f32000000000000000000000000000000000000000000000000000000008152600401610bea91815260200190565b8060096002016000828254612d789190614fa7565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba3789060200161093d565b8151835114612de8576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015612e4557612e35848281518110612e0957612e09614ebe565b6020026020010151848381518110612e2357612e23614ebe565b6020026020010151602001518461382e565b612e3e81614ed4565b9050612deb565b50505050565b60408201516001600160a01b03163b612e62575050565b600d546040517facd754d40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063acd754d490612ead90859085906004016151f0565b6020604051808303816000875af1158015612ecc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ef09190614e6c565b611b75576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005460ff1615612f795760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610bea565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612b8c3390565b6000818310612fbd5781612fbf565b825b90505b92915050565b6001810154600282015442911480612fe35750808260030154145b15612fec575050565b81600101548260020154111561302e576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260030154826130409190614fa7565b60018401548454919250613067916130589084614e9f565b8560020154611c759190614ff1565b600284015550600390910155565b61088f8180602001905181019061308c9190615570565b600061277c565b336001600160a01b038216036130eb5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610bea565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b608082015151600081900361316357505050565b60008167ffffffffffffffff81111561317e5761317e614116565b6040519080825280602002602001820160405280156131a7578160200160208202803683370190505b50905060008267ffffffffffffffff8111156131c5576131c5614116565b60405190808252806020026020018201604052801561325357816020015b60408051610180810182526000808252602080830182905292820181905260608083018290526080830182905260a0830182905260c0830182905260e083018290526101008301819052610120830152610140820181905261016082015282526000199092019101816131e35790505b50905060005b8381101561330e5760008660800151828151811061327957613279614ebe565b60200260200101518060200190518101906132949190615755565b90506132c0817f00000000000000000000000000000000000000000000000000000000000000006138ae565b8483815181106132d2576132d2614ebe565b602002602001018181525050808383815181106132f1576132f1614ebe565b6020026020010181905250508061330790614ed4565b9050613259565b50600061332f838760a001518860c001518960e001518a61010001516139b8565b5060195490915060009074010000000000000000000000000000000000000000900463ffffffff166133618342614fa7565b1190506000805b868110156137dd57600085828151811061338457613384614ebe565b6020026020010151905060006133b7826020015167ffffffffffffffff166000908152600f602052604090205460ff1690565b905060028160038111156133cd576133cd614709565b036134165760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610bea565b89156134745784806134395750600381600381111561343757613437614709565b145b61346f576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6134d1565b600081600381111561348857613488614709565b146134d15760208201516040517f67d9ba0f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610bea565b60008160038111156134e5576134e5614709565b0361359657608082015160608301516001600160a01b031660009081526017602052604090205467ffffffffffffffff9182169161352591166001615888565b67ffffffffffffffff16146135845781606001516001600160a01b0316826080015167ffffffffffffffff167fd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf4123760405160405180910390a350506137cd565b60408201516135939085614ff1565b93505b61359f82613ab0565b60208281015167ffffffffffffffff166000908152600f90915260408120805460ff191660011790556135da6135d484613c05565b8c613eda565b60208085015167ffffffffffffffff166000908152600f909152604090208054919250829160ff1916600183600381111561361757613617614709565b02179055508a156136fd578260c0015180156136445750600382600381111561364257613642614709565b145b80156136615750600281600381111561365f5761365f614709565b145b806136995750600082600381111561367b5761367b614709565b1480156136995750600281600381111561369757613697614709565b145b156136f85760608301516001600160a01b03166000908152601760205260408120805467ffffffffffffffff16916136d0836158b4565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b61377d565b8260c00151801561371f5750600381600381111561371d5761371d614709565b145b61377d5760608301516001600160a01b03166000908152601760205260408120805467ffffffffffffffff1691613755836158b4565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b826101600151836020015167ffffffffffffffff167f5e04b4755a4460aa6de4f3a906c4324a025c7449c02b52f5466659b5bfdfba5f836040516137c1919061471f565b60405180910390a35050505b6137d681614ed4565b9050613368565b506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166000908152600460205260409020546138249116823061382e565b5050505050505050565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b15801561389157600080fd5b505af11580156138a5573d6000803e3d6000fd5b50505050505050565b60008060001b828460200151856080015186606001518760e00151886101000151805190602001208961012001516040516020016138ec91906158d1565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d6040015160405160200161399a9c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600e546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce90613a15908c908c908c908c908c90600401615914565b6020604051808303816000875af1158015613a34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a589190615966565b905060008111613a94576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a613aa09084614fa7565b9350935050509550959350505050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16816000015167ffffffffffffffff1614613b305780516040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610bea565b601a54610120820151516801000000000000000090910467ffffffffffffffff161015613b9b5760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610bea565b601a546101008201515167ffffffffffffffff909116101561088f57601a54610100820151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301526024820152604401610bea565b613c586040518060e00160405280600067ffffffffffffffff1681526020016060815260200160006001600160a01b03168152602001606081526020016060815260200160608152602001600081525090565b6101208201515160008167ffffffffffffffff811115613c7a57613c7a614116565b604051908082528060200260200182016040528015613cbf57816020015b6040805180820190915260008082526020820152815260200190600190039081613c985790505b50905060008267ffffffffffffffff811115613cdd57613cdd614116565b604051908082528060200260200182016040528015613d06578160200160208202803683370190505b50905060005b83811015613e46576000613d418761012001518381518110613d3057613d30614ebe565b602002602001015160000151614012565b905080838381518110613d5657613d56614ebe565b60200260200101906001600160a01b031690816001600160a01b0316815250506040518060400160405280826001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613dbf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613de39190614fd4565b6001600160a01b031681526020018861012001518481518110613e0857613e08614ebe565b602002602001015160200151815250848381518110613e2957613e29614ebe565b60200260200101819052505080613e3f90614ed4565b9050613d0c565b506040518060e00160405280866000015167ffffffffffffffff1681526020018660600151604051602001613e8a91906001600160a01b0391909116815260200190565b60405160208183030381529060405281526020018660e001516001600160a01b0316815260200186610100015181526020018281526020018381526020018660a001518152509350505050919050565b6040517f7f738dc80000000000000000000000000000000000000000000000000000000081526000903090637f738dc890613f1b90869086906004016151f0565b600060405180830381600087803b158015613f3557600080fd5b505af1925050508015613f46575060015b614009573d808015613f74576040519150601f19603f3d011682016040523d82523d6000602084013e613f79565b606091505b50613f838161597f565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003613fd5576003915050612fc2565b806040517fcf19edfd000000000000000000000000000000000000000000000000000000008152600401610bea9190614846565b50600292915050565b6001600160a01b03818116600090815260036020526040902054168061406f576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610bea565b919050565b8280548282559060005260206000209081019282156140d6579160200282015b828111156140d6578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190614094565b5061166a929150614101565b604051806103e00160405280601f906020820280368337509192915050565b5b8082111561166a5760008155600101614102565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561414f5761414f614116565b60405290565b604051610120810167ffffffffffffffff8111828210171561414f5761414f614116565b6040805190810167ffffffffffffffff8111828210171561414f5761414f614116565b60405160e0810167ffffffffffffffff8111828210171561414f5761414f614116565b604051610180810167ffffffffffffffff8111828210171561414f5761414f614116565b604051601f8201601f1916810167ffffffffffffffff8111828210171561420c5761420c614116565b604052919050565b600067ffffffffffffffff82111561422e5761422e614116565b5060051b60200190565b67ffffffffffffffff8116811461088f57600080fd5b803561406f81614238565b600082601f83011261426a57600080fd5b8135602061427f61427a83614214565b6141e3565b82815260059290921b8401810191818101908684111561429e57600080fd5b8286015b848110156142c25780356142b581614238565b83529183019183016142a2565b509695505050505050565b6001600160a01b038116811461088f57600080fd5b803561406f816142cd565b600082601f8301126142fe57600080fd5b8135602061430e61427a83614214565b82815260059290921b8401810191818101908684111561432d57600080fd5b8286015b848110156142c2578035614344816142cd565b8352918301918301614331565b600082601f83011261436257600080fd5b8135602061437261427a83614214565b82815260059290921b8401810191818101908684111561439157600080fd5b8286015b848110156142c25780358352918301918301614395565b6fffffffffffffffffffffffffffffffff8116811461088f57600080fd5b600082601f8301126143db57600080fd5b813560206143eb61427a83614214565b8281526060928302850182019282820191908785111561440a57600080fd5b8387015b8581101561446c5781818a0312156144265760008081fd5b61442e61412c565b8135614439816142cd565b81528186013561444881614238565b8187015260408281013561445b816143ac565b90820152845292840192810161440e565b5090979650505050505050565b600067ffffffffffffffff82111561449357614493614116565b50601f01601f191660200190565b600082601f8301126144b257600080fd5b81356144c061427a82614479565b8181528460208386010111156144d557600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261450357600080fd5b8135602061451361427a83614214565b82815260059290921b8401810191818101908684111561453257600080fd5b8286015b848110156142c257803567ffffffffffffffff8111156145565760008081fd5b6145648986838b01016144a1565b845250918301918301614536565b60006020828403121561458457600080fd5b813567ffffffffffffffff8082111561459c57600080fd5b9083019061012082860312156145b157600080fd5b6145b9614155565b8235828111156145c857600080fd5b6145d487828601614259565b8252506020830135828111156145e957600080fd5b6145f5878286016142ed565b60208301525060408301358281111561460d57600080fd5b61461987828601614351565b60408301525060608301358281111561463157600080fd5b61463d878286016143ca565b60608301525060808301358281111561465557600080fd5b614661878286016144f2565b60808301525060a08301358281111561467957600080fd5b61468587828601614351565b60a08301525060c083013560c082015260e0830135828111156146a757600080fd5b6146b387828601614351565b60e0830152506101009283013592810192909252509392505050565b6000602082840312156146e157600080fd5b81356123f3816142cd565b6000602082840312156146fe57600080fd5b81356123f381614238565b634e487b7160e01b600052602160045260246000fd5b602081016004831061474157634e487b7160e01b600052602160045260246000fd5b91905290565b600060c0828403121561475957600080fd5b60405160c0810181811067ffffffffffffffff8211171561477c5761477c614116565b604052823581526020830135614791816142cd565b6020820152604083013563ffffffff811681146147ad57600080fd5b604082015260608301356147c081614238565b60608201526147d16080840161424e565b60808201526147e260a0840161424e565b60a08201529392505050565b60005b838110156148095781810151838201526020016147f1565b83811115612e455750506000910152565b600081518084526148328160208601602086016147ee565b601f01601f19169290920160200192915050565b602081526000612fbf602083018461481a565b803560ff8116811461406f57600080fd5b60008060008060008060c0878903121561488357600080fd5b863567ffffffffffffffff8082111561489b57600080fd5b6148a78a838b016142ed565b975060208901359150808211156148bd57600080fd5b6148c98a838b016142ed565b96506148d760408a01614859565b955060608901359150808211156148ed57600080fd5b6148f98a838b016144a1565b945061490760808a0161424e565b935060a089013591508082111561491d57600080fd5b5061492a89828a016144a1565b9150509295509295509295565b6000806040838503121561494a57600080fd5b823567ffffffffffffffff8082111561496257600080fd5b61496e868387016142ed565b9350602085013591508082111561498457600080fd5b5061499185828601614351565b9150509250929050565b6000602082840312156149ad57600080fd5b813567ffffffffffffffff8111156149c457600080fd5b6149d0848285016142ed565b949350505050565b6020808252825182820181905260009190848201906040850190845b81811015614a10578351835292840192918401916001016149f4565b50909695505050505050565b600081518084526020808501945080840160005b83811015614a555781516001600160a01b031687529582019590820190600101614a30565b509495945050505050565b602081526000612fbf6020830184614a1c565b6020808252825182820181905260009190848201906040850190845b81811015614a105783516001600160a01b031683529284019291840191600101614a8f565b60008060408385031215614ac757600080fd5b8235614ad2816142cd565b91506020830135614ae2816142cd565b809150509250929050565b600082601f830112614afe57600080fd5b81356020614b0e61427a83614214565b82815260069290921b84018101918181019086841115614b2d57600080fd5b8286015b848110156142c25760408189031215614b4a5760008081fd5b614b52614179565b8135614b5d816142cd565b81528185013585820152835291830191604001614b31565b801515811461088f57600080fd5b803561406f81614b75565b60008060408385031215614ba157600080fd5b823567ffffffffffffffff80821115614bb957600080fd5b9084019060e08287031215614bcd57600080fd5b614bd561419c565b614bde8361424e565b8152602083013582811115614bf257600080fd5b614bfe888286016144a1565b602083015250614c10604084016142e2565b6040820152606083013582811115614c2757600080fd5b614c33888286016144a1565b606083015250608083013582811115614c4b57600080fd5b614c57888286016142ed565b60808301525060a083013582811115614c6f57600080fd5b614c7b88828601614aed565b60a08301525060c083013560c0820152809450505050614c9d60208401614b83565b90509250929050565b600060408284031215614cb857600080fd5b6040516040810181811067ffffffffffffffff82111715614cdb57614cdb614116565b604052823581526020928301359281019290925250919050565b60008083601f840112614d0757600080fd5b50813567ffffffffffffffff811115614d1f57600080fd5b6020830191508360208260051b8501011115614d3a57600080fd5b9250929050565b60008060008060008060008060e0898b031215614d5d57600080fd5b606089018a811115614d6e57600080fd5b8998503567ffffffffffffffff80821115614d8857600080fd5b818b0191508b601f830112614d9c57600080fd5b813581811115614dab57600080fd5b8c6020828501011115614dbd57600080fd5b6020830199508098505060808b0135915080821115614ddb57600080fd5b614de78c838d01614cf5565b909750955060a08b0135915080821115614e0057600080fd5b50614e0d8b828c01614cf5565b999c989b50969995989497949560c00135949350505050565b600060208284031215614e3857600080fd5b813567ffffffffffffffff811115614e4f57600080fd5b820160e081850312156123f357600080fd5b805161406f81614b75565b600060208284031215614e7e57600080fd5b81516123f381614b75565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615614eb957614eb9614e89565b500290565b634e487b7160e01b600052603260045260246000fd5b60006000198203614ee757614ee7614e89565b5060010190565b600063ffffffff808316818103614f0757614f07614e89565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614f418184018a614a1c565b90508281036080840152614f558189614a1c565b905060ff871660a084015282810360c0840152614f72818761481a565b905067ffffffffffffffff851660e0840152828103610100840152614f97818561481a565b9c9b505050505050505050505050565b600082821015614fb957614fb9614e89565b500390565b634e487b7160e01b600052603160045260246000fd5b600060208284031215614fe657600080fd5b81516123f3816142cd565b6000821982111561500457615004614e89565b500190565b600060ff821660ff84168060ff0382111561502657615026614e89565b019392505050565b634e487b7160e01b600052601260045260246000fd5b600060ff8316806150575761505761502e565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b602080825282518282018190526000919060409081850190868401855b8281101561510357815180516001600160a01b031685528681015167ffffffffffffffff16878601528501516fffffffffffffffffffffffffffffffff1685850152606090930192908501906001016150af565b5091979650505050505050565b60008261511f5761511f61502e565b500490565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b16604085015281606085015261515e8285018b614a1c565b91508382036080850152615172828a614a1c565b915060ff881660a085015283820360c085015261518f828861481a565b90861660e08501528381036101008501529050614f97818561481a565b600081518084526020808501945080840160005b83811015614a5557815180516001600160a01b0316885283015183880152604090960195908201906001016151c0565b6040815267ffffffffffffffff83511660408201526000602084015160e0606084015261522161012084018261481a565b9050604085015161523d60808501826001600160a01b03169052565b5060608501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808584030160a0860152615278838361481a565b925060808701519150808584030160c08601526152958383614a1c565b925060a08701519150808584030160e0860152506152b382826151ac565b60c0870151610100860152851515602086015292506123f3915050565b805161406f81614238565b600082601f8301126152ec57600080fd5b815160206152fc61427a83614214565b82815260059290921b8401810191818101908684111561531b57600080fd5b8286015b848110156142c257805161533281614238565b835291830191830161531f565b805161406f816142cd565b600082601f83011261535b57600080fd5b8151602061536b61427a83614214565b82815260059290921b8401810191818101908684111561538a57600080fd5b8286015b848110156142c25780516153a1816142cd565b835291830191830161538e565b600082601f8301126153bf57600080fd5b815160206153cf61427a83614214565b82815260059290921b840181019181810190868411156153ee57600080fd5b8286015b848110156142c257805183529183019183016153f2565b600082601f83011261541a57600080fd5b8151602061542a61427a83614214565b8281526060928302850182019282820191908785111561544957600080fd5b8387015b8581101561446c5781818a0312156154655760008081fd5b61546d61412c565b8151615478816142cd565b81528186015161548781614238565b8187015260408281015161549a816143ac565b90820152845292840192810161544d565b600082601f8301126154bc57600080fd5b81516154ca61427a82614479565b8181528460208386010111156154df57600080fd5b6149d08260208301602087016147ee565b600082601f83011261550157600080fd5b8151602061551161427a83614214565b82815260059290921b8401810191818101908684111561553057600080fd5b8286015b848110156142c257805167ffffffffffffffff8111156155545760008081fd5b6155628986838b01016154ab565b845250918301918301615534565b60006020828403121561558257600080fd5b815167ffffffffffffffff8082111561559a57600080fd5b9083019061012082860312156155af57600080fd5b6155b7614155565b8251828111156155c657600080fd5b6155d2878286016152db565b8252506020830151828111156155e757600080fd5b6155f38782860161534a565b60208301525060408301518281111561560b57600080fd5b615617878286016153ae565b60408301525060608301518281111561562f57600080fd5b61563b87828601615409565b60608301525060808301518281111561565357600080fd5b61565f878286016154f0565b60808301525060a08301518281111561567757600080fd5b615683878286016153ae565b60a08301525060c083015160c082015260e0830151828111156156a557600080fd5b6156b1878286016153ae565b60e0830152506101009283015192810192909252509392505050565b600082601f8301126156de57600080fd5b815160206156ee61427a83614214565b82815260069290921b8401810191818101908684111561570d57600080fd5b8286015b848110156142c2576040818903121561572a5760008081fd5b615732614179565b815161573d816142cd565b81528185015185820152835291830191604001615711565b60006020828403121561576757600080fd5b815167ffffffffffffffff8082111561577f57600080fd5b90830190610180828603121561579457600080fd5b61579c6141bf565b6157a5836152d0565b81526157b3602084016152d0565b6020820152604083015160408201526157ce6060840161533f565b60608201526157df608084016152d0565b608082015260a083015160a08201526157fa60c08401614e61565b60c082015261580b60e0840161533f565b60e0820152610100808401518381111561582457600080fd5b615830888287016154ab565b828401525050610120808401518381111561584a57600080fd5b615856888287016156cd565b828401525050610140915061586c82840161533f565b9181019190915261016091820151918101919091529392505050565b600067ffffffffffffffff8083168185168083038211156158ab576158ab614e89565b01949350505050565b600067ffffffffffffffff808316818103614f0757614f07614e89565b602081526000612fbf60208301846151ac565b600081518084526020808501945080840160005b83811015614a55578151875295820195908201906001016158f8565b60a08152600061592760a08301886158e4565b828103602084015261593981886158e4565b9050856040840152828103606084015261595381866158e4565b9150508260808301529695505050505050565b60006020828403121561597857600080fd5b5051919050565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156159c75780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCaller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMGEOffRamp.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) GetTransmitters() ([]common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetTransmitters(&_EVM2EVMGEOffRamp.CallOpts)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampCallerSession) GetTransmitters() ([]common.Address, error) {
	return _EVM2EVMGEOffRamp.Contract.GetTransmitters(&_EVM2EVMGEOffRamp.CallOpts)
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

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) SetGEConfig(opts *bind.TransactOpts, config IEVM2EVMGEOffRampGEOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "setGEConfig", config)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) SetGEConfig(config IEVM2EVMGEOffRampGEOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetGEConfig(&_EVM2EVMGEOffRamp.TransactOpts, config)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) SetGEConfig(config IEVM2EVMGEOffRampGEOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetGEConfig(&_EVM2EVMGEOffRamp.TransactOpts, config)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactor) SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.contract.Transact(opts, "setOCR2Config", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetOCR2Config(&_EVM2EVMGEOffRamp.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMGEOffRamp *EVM2EVMGEOffRampTransactorSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMGEOffRamp.Contract.SetOCR2Config(&_EVM2EVMGEOffRamp.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
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

	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMMessageFromSender, manualExecution bool) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report GEExecutionReport) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetCommitStore(opts *bind.TransactOpts, commitStore common.Address) (*types.Transaction, error)

	SetGEConfig(opts *bind.TransactOpts, config IEVM2EVMGEOffRampGEOffRampConfig) (*types.Transaction, error)

	SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

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
