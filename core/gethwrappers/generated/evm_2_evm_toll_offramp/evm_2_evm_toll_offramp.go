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

type CommonEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
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

type InternalAny2EVMMessageFromSender struct {
	SourceChainId        uint64
	Sender               []byte
	Receiver             common.Address
	Data                 []byte
	DestPools            []common.Address
	DestTokensAndAmounts []CommonEVMTokenAndAmount
	GasLimit             *big.Int
}

type TollEVM2EVMTollMessage struct {
	SourceChainId     uint64
	SequenceNumber    uint64
	Sender            common.Address
	Receiver          common.Address
	Data              []byte
	TokensAndAmounts  []CommonEVMTokenAndAmount
	FeeTokenAndAmount CommonEVMTokenAndAmount
	GasLimit          *big.Int
}

type TollExecutionReport struct {
	SequenceNumbers          []uint64
	TokenPerFeeCoinAddresses []common.Address
	TokenPerFeeCoin          []*big.Int
	EncodedMessages          [][]byte
	InnerProofs              [][32]byte
	InnerProofFlagBits       *big.Int
	OuterProofs              [][32]byte
	OuterProofFlagBits       *big.Int
}

var EVM2EVMTollOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIBaseOffRamp.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractICommitStore\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIBaseOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeTaken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitStore\",\"outputs\":[{\"internalType\":\"contractICommitStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIBaseOffRamp.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolByDestToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractIAny2EVMOffRampRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structToll.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"merkleGasShare\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount\",\"name\":\"feeTokenAndAmount\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structToll.EVM2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"overheadGasToll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICommitStore\",\"name\":\"commitStore\",\"type\":\"address\"}],\"name\":\"setCommitStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIBaseOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAny2EVMOffRampRouter\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200674c3803806200674c8339810160408190526200003591620007fe565b6000805460ff191681556001908b908b908a908a908a908a908a908a908a90829082908690869089903390819081620000b55760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000ef57620000ef8162000446565b5050506001600160a01b0381166200011a57604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b039290921691909117905580518251146200015c5760405162d8548360e71b815260040160405180910390fd5b815162000171906005906020850190620004f7565b5060005b8251811015620002fa5760006040518060400160405280848481518110620001a157620001a162000905565b60200260200101516001600160a01b03168152602001836001600160601b031681525090508060036000868581518110620001e057620001e062000905565b6020908102919091018101516001600160a01b0390811683528282019390935260409091016000908120845194909201516001600160601b0316600160a01b02939092169290921790915581518451909160049186908690811062000249576200024962000905565b60200260200101516001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200028f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002b591906200091b565b6001600160a01b039081168252602082019290925260400160002080546001600160a01b0319169290911691909117905550620002f28162000942565b905062000175565b5050600680546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600955600a93909355600b55600c91909155871662000377576040516342bcdf7f60e11b815260040160405180910390fd5b505050506001600160401b039485166080525091831660a0526001600160a01b0390811660c052600e80546001600160a01b0319169190921617905590151560e05288516017805460208c015160408d01516060909d015163ffffffff9094166001600160601b0319909216919091176401000000009185169190910217600160601b600160e01b0319166c010000000000000000000000009b84169b909b02600160a01b600160e01b0319169a909a17600160a01b919092160217909755506200096a975050505050505050565b336001600160a01b03821603620004a05760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000ac565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8280548282559060005260206000209081019282156200054f579160200282015b828111156200054f57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000518565b506200055d92915062000561565b5090565b5b808211156200055d576000815560010162000562565b80516001600160401b03811681146200059057600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620005d657620005d662000595565b604052919050565b600060808284031215620005f157600080fd5b604051608081016001600160401b038111828210171562000616576200061662000595565b8060405250809150825163ffffffff811681146200063357600080fd5b8152620006436020840162000578565b6020820152620006566040840162000578565b6040820152620006696060840162000578565b60608201525092915050565b6001600160a01b03811681146200068b57600080fd5b50565b8051620005908162000675565b60006001600160401b03821115620006b757620006b762000595565b5060051b60200190565b600082601f830112620006d357600080fd5b81516020620006ec620006e6836200069b565b620005ab565b82815260059290921b840181019181810190868411156200070c57600080fd5b8286015b8481101562000734578051620007268162000675565b835291830191830162000710565b509695505050505050565b600082601f8301126200075157600080fd5b8151602062000764620006e6836200069b565b82815260059290921b840181019181810190868411156200078457600080fd5b8286015b84811015620007345780516200079e8162000675565b835291830191830162000788565b600060408284031215620007bf57600080fd5b604080519081016001600160401b0381118282101715620007e457620007e462000595565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000806101c08b8d0312156200081f57600080fd5b6200082a8b62000578565b99506200083a60208c0162000578565b98506200084b8c60408d01620005de565b97506200085b60c08c016200068e565b96506200086b60e08c016200068e565b95506200087c6101008c016200068e565b6101208c01519095506001600160401b03808211156200089b57600080fd5b620008a98e838f01620006c1565b95506101408d0151915080821115620008c157600080fd5b50620008d08d828e016200073f565b935050620008e38c6101608d01620007ac565b9150620008f46101a08c016200068e565b90509295989b9194979a5092959850565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156200092e57600080fd5b81516200093b8162000675565b9392505050565b6000600182016200096357634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e051615d85620009c760003960006116e1015260008181611e1c01526138b301526000818161030c01526138920152600081816102e701528181611df8015281816138710152613b0e0152615d856000f3fe608060405234801561001057600080fd5b50600436106102d35760003560e01c806389c0656811610186578063b66f0efb116100e3578063d30a364b11610097578063ea03d1b511610071578063ea03d1b5146107b4578063eb511dd4146107c2578063f2fde38b146107d557600080fd5b8063d30a364b14610762578063d7e2bb5014610775578063e3d0e712146107a157600080fd5b8063c0d78655116100c8578063c0d786551461065f578063c3f909d414610672578063c90332841461074f57600080fd5b8063b66f0efb1461063b578063bc29705a1461064c57600080fd5b8063a8b640c11161013a578063b0f479a11161011f578063b0f479a114610604578063b1dc65a414610615578063b4069b311461062857600080fd5b8063a8b640c1146105c4578063afcb95d7146105e457600080fd5b806390c2339b1161016b57806390c2339b146105555780639129badf1461059057806391872543146105b157600080fd5b806389c06568146105375780638da5cb5b1461053f57600080fd5b8063599f64311161023457806379ba5097116101e857806381411834116101cd57806381411834146104ea57806381ff7048146104ff5780638456cb591461052f57600080fd5b806379ba5097146104cf5780637f738dc8146104d757600080fd5b80635d86f141116102195780635d86f1411461047b578063681fba16146104a7578063744b92e2146104bc57600080fd5b8063599f64311461045f5780635c975abb1461047057600080fd5b80632222dd421161028b5780633f4ba83a116102705780633f4ba83a146104245780634352fa9f1461042c5780634741062e1461043f57600080fd5b80632222dd42146103ec57806339aa92641461041157600080fd5b8063142a98fc116102bc578063142a98fc14610351578063147809b31461038b578063181f5a77146103a357600080fd5b8063087ae6df146102d8578063108ee5fc1461033c575b600080fd5b6040805167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811682527f0000000000000000000000000000000000000000000000000000000000000000166020820152015b60405180910390f35b61034f61034a36600461462e565b6107e8565b005b61037e61035f36600461466c565b67ffffffffffffffff166000908152600f602052604090205460ff1690565b604051610333919061469f565b61039361089f565b6040519015158152602001610333565b6103df6040518060400160405280601881526020017f45564d3245564d546f6c6c4f666652616d7020312e302e30000000000000000081525081565b604051610333919061471f565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610333565b61034f61041f36600461462e565b61092c565b61034f610963565b61034f61043a3660046148e7565b610975565b61045261044d36600461494b565b610bca565b6040516103339190614988565b6006546001600160a01b03166103f9565b60005460ff16610393565b6103f961048936600461462e565b6001600160a01b039081166000908152600360205260409020541690565b6104af610c92565b60405161033391906149cc565b61034f6104ca366004614a0d565b610d57565b61034f611109565b61034f6104e5366004614bdb565b6111f1565b6104f2611267565b6040516103339190614d37565b6012546010546040805163ffffffff80851682526401000000009094049093166020840152820152606001610333565b61034f6112c9565b6104af6112d9565b60005461010090046001600160a01b03166103f9565b61055d611339565b60405161033391908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6105a361059e366004614d4a565b6113da565b604051908152602001610333565b61034f6105bf366004614e49565b61148b565b6105a36105d2366004614e98565b60166020526000908152604090205481565b604080516001815260006020820181905291810191909152606001610333565b600d546001600160a01b03166103f9565b61034f610623366004614efd565b6115be565b6103f961063636600461462e565b611b91565b600e546001600160a01b03166103f9565b61034f61065a366004614fe2565b611c7f565b61034f61066d36600461462e565b611db5565b610704604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260175463ffffffff8116825267ffffffffffffffff6401000000008204811660208401526c010000000000000000000000008204811693830193909352740100000000000000000000000000000000000000009004909116606082015290565b6040516103339190815163ffffffff16815260208083015167ffffffffffffffff90811691830191909152604080840151821690830152606092830151169181019190915260800190565b61034f61075d36600461462e565b611e73565b61034f610770366004615154565b611eaa565b6103f961078336600461462e565b6001600160a01b039081166000908152600460205260409020541690565b61034f6107af36600461529e565b611eb8565b61034f6102d336600461536b565b61034f6107d0366004614a0d565b61271a565b61034f6107e336600461462e565b612996565b6107f06129a7565b6001600160a01b038116610830576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa158015610902573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092691906153a6565b15905090565b6109346129a7565b6006805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b61096b6129a7565b610973612a06565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b0316141580156109b257506006546001600160a01b03163314155b156109e9576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114610a25576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085460005b81811015610a7f576007600060088381548110610a4a57610a4a6153c3565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610a78816153ef565b9050610a2b565b5060005b82811015610baf576000858281518110610a9f57610a9f6153c3565b6020026020010151905060006001600160a01b0316816001600160a01b031603610af5576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610b0757610b076153c3565b602002602001015160076000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610b6c57610b6c6153c3565b6020026020010151604051610b969291906001600160a01b03929092168252602082015260400190565b60405180910390a150610ba8816153ef565b9050610a83565b508351610bc3906008906020870190614577565b5050505050565b80516060908067ffffffffffffffff811115610be857610be861473d565b604051908082528060200260200182016040528015610c11578160200160208202803683370190505b50915060005b81811015610c8b5760076000858381518110610c3557610c356153c3565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610c7057610c706153c3565b6020908102919091010152610c84816153ef565b9050610c17565b5050919050565b60055460609067ffffffffffffffff811115610cb057610cb061473d565b604051908082528060200260200182016040528015610cd9578160200160208202803683370190505b50905060005b600554811015610d5357610d1960058281548110610cff57610cff6153c3565b6000918252602090912001546001600160a01b0316611b91565b828281518110610d2b57610d2b6153c3565b6001600160a01b0390921660209283029190910190910152610d4c816153ef565b9050610cdf565b5090565b610d5f6129a7565b6005546000819003610d9d576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290610e2b576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614610e7a576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006005610e89600185615409565b81548110610e9957610e996153c3565b9060005260206000200160009054906101000a90046001600160a01b03169050600582602001516bffffffffffffffffffffffff1681548110610ede57610ede6153c3565b6000918252602090912001546001600160a01b03166005610f00600186615409565b81548110610f1057610f106153c3565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600583602001516bffffffffffffffffffffffff1681548110610f6457610f646153c3565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff9092167401000000000000000000000000000000000000000002919092161790556005805480610fee57610fee615420565b6001900381819060005260206000200160006101000a8154906001600160a01b030219169055905560046000856001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611058573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061107c9190615436565b6001600160a01b03908116825260208083019390935260409182016000908120805473ffffffffffffffffffffffffffffffffffffffff1916905588821680825260038552838220919091558251908152908716928101929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146111685760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b33301461122a576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a08201515115611259576112428260a00151612aa2565b61125982608001518360a001518460400151612ca6565b6112638282612d44565b5050565b606060158054806020026020016040519081016040528092919081815260200182805480156112bf57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116112a1575b5050505050905090565b6112d16129a7565b610973612e1f565b606060058054806020026020016040519081016040528092919081815260200182805480156112bf576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116112a1575050505050905090565b6113646040518060800160405280600081526020016000815260200160008152602001600081525090565b604080516080810182526009548152600a546020820152600b5491810191909152600c5460608201819052429060009061139e9083615409565b602084015184519192506113ca916113b69084615453565b85604001516113c59190615472565b612ea7565b6040840152506060820152919050565b6000808260800151518360a0015151602060146113f79190615472565b6114019190615453565b61140c906086615472565b6114169190615472565b90506000611425601083615453565b9050610a28611bbc8560a0015151600161143f9190615472565b61144b90618aac615453565b6156b86114588986615472565b6114629190615472565b61146c9190615472565b6114769190615472565b6114809190615472565b925050505b92915050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b0316141580156114c857506006546001600160a01b03163314155b156114ff576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff11611553576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61155d6009612ebd565b6020810151600a8190558151600955600b546115799190612ea7565b600b55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c01359161161491849163ffffffff851691908e908e9081908401838280828437600092019190915250612f6a92505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260105480825260115460ff808216602085015261010090910416928201929092529083146116cf5760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d617463680000000000000000000000604482015260640161115f565b6116dd8b8b8b8b8b8b612f8d565b60007f00000000000000000000000000000000000000000000000000000000000000001561173a5760028260200151836040015161171b919061548a565b61172591906154c5565b61173090600161548a565b60ff169050611750565b602082015161174a90600161548a565b60ff1690505b88811461179f5760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e617475726573000000000000604482015260640161115f565b8887146117ee5760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e0000604482015260640161115f565b3360009081526013602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561183157611831614689565b600281111561184257611842614689565b905250905060028160200151600281111561185f5761185f614689565b14801561189957506015816000015160ff1681548110611881576118816153c3565b6000918252602090912001546001600160a01b031633145b6118e55760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d69747465720000000000000000604482015260640161115f565b5050505050600088886040516118fc9291906154e7565b604051908190038120611913918c906020016154f7565b6040516020818303038152906040528051906020012090506119336145e5565b604080518082019091526000808252602082015260005b88811015611b6f576000600185888460208110611969576119696153c3565b61197691901a601b61548a565b8d8d86818110611988576119886153c3565b905060200201358c8c878181106119a1576119a16153c3565b90506020020135604051600081526020016040526040516119de949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611a00573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526013602090815290849020838501909452835460ff80821685529296509294508401916101009004166002811115611a5557611a55614689565b6002811115611a6657611a66614689565b9052509250600183602001516002811115611a8357611a83614689565b14611ad05760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e0000604482015260640161115f565b8251849060ff16601f8110611ae757611ae76153c3565b602002015115611b395760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e6174757265000000000000000000000000604482015260640161115f565b600184846000015160ff16601f8110611b5457611b546153c3565b9115156020909202015250611b68816153ef565b905061194a565b5050505063ffffffff8110611b8657611b86615513565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611be5576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611c54573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c789190615436565b9392505050565b611c876129a7565b80516017805460208085018051604080880180516060808b01805167ffffffffffffffff90811674010000000000000000000000000000000000000000027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff9482166c0100000000000000000000000002949094167fffffffff00000000000000000000000000000000ffffffffffffffffffffffff978216640100000000027fffffffffffffffffffffffffffffffffffffffff000000000000000000000000909b1663ffffffff909d169c8d179a909a17969096169890981791909117909755815197885292518216938701939093529051811691850191909152905116908201527f1a9ff9caccd597226ff8c393ab44e6b57669905d9b16f1cdb2ac267253dbf27c906080016115b3565b611dbd6129a7565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038381169182179092556040805167ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681527f0000000000000000000000000000000000000000000000000000000000000000909316602084015290917f052b5907be1d3ac35d571862117562e80ee743c01251e388dafb7dc4e92a726c910160405180910390a250565b611e7b6129a7565b600e805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b611eb581600161301d565b50565b855185518560ff16601f831115611f2b576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e65727300000000000000000000000000000000604482015260640161115f565b80600003611f95576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f7369746976650000000000000000000000000000604482015260640161115f565b818314612023576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e00000000000000000000000000000000000000000000000000000000606482015260840161115f565b61202e816003615453565b8311612096576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f20686967680000000000000000604482015260640161115f565b61209e6129a7565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b60145415612208576014546000906120f690600190615409565b905060006014828154811061210d5761210d6153c3565b6000918252602082200154601580546001600160a01b039092169350908490811061213a5761213a6153c3565b60009182526020808320909101546001600160a01b03858116845260139092526040808420805461ffff199081169091559290911680845292208054909116905560148054919250908061219057612190615420565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff1916905501905560158054806121d0576121d0615420565b6000828152602090208101600019908101805473ffffffffffffffffffffffffffffffffffffffff19169055019055506120dc915050565b60005b8151518110156125935760006013600084600001518481518110612231576122316153c3565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff16600281111561226e5761226e614689565b146122d5576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e65722061646472657373000000000000000000604482015260640161115f565b6040805180820190915260ff82168152600160208201528251805160139160009185908110612306576123066153c3565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff19161761010083600281111561235f5761235f614689565b02179055506000915061236f9050565b6013600084602001518481518110612389576123896153c3565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff1660028111156123c6576123c6614689565b1461242d576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161115f565b6040805180820190915260ff821681526020810160028152506013600084602001518481518110612460576124606153c3565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff1916176101008360028111156124b9576124b9614689565b0217905550508251805160149250839081106124d7576124d76153c3565b6020908102919091018101518254600181018455600093845292829020909201805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909316929092179091558201518051601591908390811061253b5761253b6153c3565b602090810291909101810151825460018101845560009384529190922001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0390921691909117905561258c816153ef565b905061220b565b5060408101516011805460ff191660ff909216919091179055601280547fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff811664010000000063ffffffff438116820292831785559083048116936001939092600092612607928692908216911617615529565b92506101000a81548163ffffffff021916908363ffffffff1602179055506126664630601260009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a001516136ac565b6010819055825180516011805460ff909216610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff90921691909117905560125460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612705988b98919763ffffffff909216969095919491939192615551565b60405180910390a15050505050505050505050565b6127226129a7565b6001600160a01b038216158061273f57506001600160a01b038116155b15612776576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612805576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038083168083526005546bffffffffffffffffffffffff90811660208086019182528785166000908152600382526040808220885194519095167401000000000000000000000000000000000000000002939096169290921790925583517f21df0da70000000000000000000000000000000000000000000000000000000081529351869460049492936321df0da79282870192819003870181865afa1580156128ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128de9190615436565b6001600160a01b03908116825260208083019390935260409182016000908120805495831673ffffffffffffffffffffffffffffffffffffffff199687161790556005805460018101825591527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001805488831695168517905581519384528516918301919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b61299e6129a7565b611eb581613739565b60005461010090046001600160a01b031633146109735760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161115f565b60005460ff16612a585760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161115f565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6000805b8251811015612ba157600060076000858481518110612ac757612ac76153c3565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600003612b5a57838281518110612b1057612b106153c3565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b03909116600482015260240161115f565b838281518110612b6c57612b6c6153c3565b60200260200101516020015181612b839190615453565b612b8d9084615472565b92505080612b9a906153ef565b9050612aa6565b50801561126357612bb26009612ebd565b600a54811115612bfc57600a546040517f688ccf7700000000000000000000000000000000000000000000000000000000815260048101919091526024810182905260440161115f565b600b54811115612c5c57600954600b5460009190612c1a9084615409565b612c2491906155e7565b9050806040517fe31e0f3200000000000000000000000000000000000000000000000000000000815260040161115f91815260200190565b8060096002016000828254612c719190615409565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba37890602001610893565b8151835114612ce1576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015612d3e57612d2e848281518110612d0257612d026153c3565b6020026020010151848381518110612d1c57612d1c6153c3565b602002602001015160200151846137f5565b612d37816153ef565b9050612ce4565b50505050565b60408201516001600160a01b03163b612d5b575050565b600d546040517facd754d40000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063acd754d490612da69085908590600401615648565b6020604051808303816000875af1158015612dc5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612de991906153a6565b611263576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005460ff1615612e725760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161115f565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612a853390565b6000818310612eb65781611c78565b5090919050565b6001810154600282015442911480612ed85750808260030154145b15612ee1575050565b816001015482600201541115612f23576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482612f359190615409565b60018401548454919250612f5c91612f4d9084615453565b85600201546113c59190615472565b600284015550600390910155565b612f8881806020019051810190612f819190615926565b600061301d565b505050565b6000612f9a826020615453565b612fa5856020615453565b612fb188610144615472565b612fbb9190615472565b612fc59190615472565b612fd0906000615472565b9050368114613014576040517f8e1192e10000000000000000000000000000000000000000000000000000000081526004810182905236602482015260440161115f565b50505050505050565b60005460ff16156130705760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161115f565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156130c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130e791906153a6565b1561311d576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d546001600160a01b031661315f576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608201515160008190036131a0576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff8111156131bb576131bb61473d565b60405190808252806020026020018201604052801561324257816020015b61322f60408051610100810182526000808252602080830182905282840182905260608084018390526080840181905260a084015283518085019094528184528301529060c08201908152602001600081525090565b8152602001906001900390816131d95790505b50905060008267ffffffffffffffff8111156132605761326061473d565b604051908082528060200260200182016040528015613289578160200160208202803683370190505b50905060006132b77fb9b8993db34ae003b2aacdae4cdef2888717531ab95157174f8f0dbf076b5e5861386c565b905060005b84811015613351576000876060015182815181106132dc576132dc6153c3565b60200260200101518060200190518101906132f79190615afe565b9050613303818461392c565b848381518110613315576133156153c3565b60200260200101818152505080858381518110613334576133346153c3565b6020026020010181905250508061334a906153ef565b90506132bc565b506000806133728489608001518a60a001518b60c001518c60e00151613a14565b601754919350915060009063ffffffff1661338d8442615409565b11905060005b878110156136a05760008782815181106133af576133af6153c3565b6020026020010151905060006133e2826020015167ffffffffffffffff166000908152600f602052604090205460ff1690565b905060028160038111156133f8576133f8614689565b036134415760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161115f565b8a158061344b5750835b806134675750600381600381111561346557613465614689565b145b61349d576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6134a682613b0c565b6000808260038111156134bb576134bb614689565b1480156134c657508b155b1561351f576134e28a51876134db91906155e7565b8e85613c88565b905060006134f78460c0015160000151613e75565b90506135048183306137f5565b818460c0015160200181815161351a9190615409565b905250505b600082600381111561353357613533614689565b1461356f5760208084015167ffffffffffffffff16600090815260168252604090205460c0850151909101805161356b908390615409565b9052505b60208381015167ffffffffffffffff166000908152600f90915260408120805460ff191660011790556135aa6135a485613ed7565b8e6141b4565b60208086015167ffffffffffffffff166000908152600f909152604090208054919250829160ff191660018360038111156135e7576135e7614689565b0217905550600083600381111561360057613600614689565b14801561361e5750600381600381111561361c5761361c614689565b145b156136455760208085015167ffffffffffffffff1660009081526016909152604090208290555b836020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf82604051613683919061469f565b60405180910390a25050505080613699906153ef565b9050613393565b50505050505050505050565b6000808a8a8a8a8a8a8a8a8a6040516020016136d099989796959493929190615bf2565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b038216036137915760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161115f565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b15801561385857600080fd5b505af1158015613014573d6000803e3d6000fd5b6000817f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060405160200161390f949392919093845267ffffffffffffffff9283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b60008060001b828460200151856040015186606001518760800151805190602001208860a001516040516020016139639190615c7a565b604051602081830303815290604052805190602001208960e001518a60c001516040516020016139f6999897969594939291909889526020808a019890985267ffffffffffffffff9690961660408901526001600160a01b039485166060890152928416608088015260a087019190915260c086015260e085015281511661010084015201516101208201526101400190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600e546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce90613a71908c908c908c908c908c90600401615cbd565b6020604051808303816000875af1158015613a90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ab49190615d0f565b905060008111613af0576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a613afc9084615409565b9350935050509550959350505050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16816000015167ffffffffffffffff1614613b8c5780516040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161115f565b60175460a0820151517401000000000000000000000000000000000000000090910467ffffffffffffffff161015613c025760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161115f565b6017546080820151516c0100000000000000000000000090910467ffffffffffffffff161015611eb5576017546080820151516040517f869337890000000000000000000000000000000000000000000000000000000081526c0100000000000000000000000090920467ffffffffffffffff166004830152602482015260440161115f565b6000806000613c9e8460c0015160000151613e75565b6001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613cdb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613cff9190615436565b905060005b856020015151811015613d7857816001600160a01b031686602001518281518110613d3157613d316153c3565b60200260200101516001600160a01b031603613d685785604001518181518110613d5d57613d5d6153c3565b602002602001015192505b613d71816153ef565b9050613d04565b5081613dbb576040517fce480bcc0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161115f565b6000670de0b6b3a7640000833a8760e00151613dd78b8a6113da565b613de19190615472565b613deb9190615453565b613df59190615453565b613dff91906155e7565b90508460c0015160200151811115613e6b5760208086015160c0870151909101516040517f3cab2f4d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909216600483015260248201839052604482015260640161115f565b9695505050505050565b6001600160a01b038181166000908152600360205260409020541680613ed2576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260240161115f565b919050565b613f2a6040518060e00160405280600067ffffffffffffffff1681526020016060815260200160006001600160a01b03168152602001606081526020016060815260200160608152602001600081525090565b6000613f3e8360a001518460c001516142ec565b805190915060008167ffffffffffffffff811115613f5e57613f5e61473d565b604051908082528060200260200182016040528015613fa357816020015b6040805180820190915260008082526020820152815260200190600190039081613f7c5790505b50905060008267ffffffffffffffff811115613fc157613fc161473d565b604051908082528060200260200182016040528015613fea578160200160208202803683370190505b50905060005b8381101561412057600061402086838151811061400f5761400f6153c3565b602002602001015160000151613e75565b905080838381518110614035576140356153c3565b60200260200101906001600160a01b031690816001600160a01b0316815250506040518060400160405280826001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561409e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906140c29190615436565b6001600160a01b031681526020018784815181106140e2576140e26153c3565b602002602001015160200151815250848381518110614103576141036153c3565b60200260200101819052505080614119906153ef565b9050613ff0565b506040518060e00160405280876000015167ffffffffffffffff168152602001876040015160405160200161416491906001600160a01b0391909116815260200190565b604051602081830303815290604052815260200187606001516001600160a01b03168152602001876080015181526020018281526020018381526020018760e00151815250945050505050919050565b6040517f7f738dc80000000000000000000000000000000000000000000000000000000081526000903090637f738dc8906141f59086908690600401615648565b600060405180830381600087803b15801561420f57600080fd5b505af1925050508015614220575060015b6142e3573d80801561424e576040519150601f19603f3d011682016040523d82523d6000602084013e614253565b606091505b5061425d81615d28565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da800000000000000000000000000000000000000000000000000000000036142af576003915050611485565b806040517fcf19edfd00000000000000000000000000000000000000000000000000000000815260040161115f919061471f565b50600292915050565b606060005b835181101561448b5782600001516001600160a01b031684828151811061431a5761431a6153c3565b6020026020010151600001516001600160a01b03160361447b576000845167ffffffffffffffff8111156143505761435061473d565b60405190808252806020026020018201604052801561439557816020015b604080518082019091526000808252602082015281526020019060019003908161436e5790505b50905060005b85518110156143ec578581815181106143b6576143b66153c3565b60200260200101518282815181106143d0576143d06153c3565b6020026020010181905250806143e5906153ef565b905061439b565b50604051806040016040528082848151811061440a5761440a6153c3565b6020026020010151600001516001600160a01b03168152602001856020015183858151811061443b5761443b6153c3565b6020026020010151602001516144519190615472565b815250818381518110614466576144666153c3565b60200260200101819052508092505050611485565b614484816153ef565b90506142f1565b5060008351600161449c9190615472565b67ffffffffffffffff8111156144b4576144b461473d565b6040519080825280602002602001820160405280156144f957816020015b60408051808201909152600080825260208201528152602001906001900390816144d25790505b50905060005b84518110156145505784818151811061451a5761451a6153c3565b6020026020010151828281518110614534576145346153c3565b602002602001018190525080614549906153ef565b90506144ff565b508281855181518110614565576145656153c3565b60209081029190910101529392505050565b8280548282559060005260206000209081019282156145d9579160200282015b828111156145d9578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190614597565b50610d53929150614604565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610d535760008155600101614605565b6001600160a01b0381168114611eb557600080fd5b60006020828403121561464057600080fd5b8135611c7881614619565b67ffffffffffffffff81168114611eb557600080fd5b8035613ed28161464b565b60006020828403121561467e57600080fd5b8135611c788161464b565b634e487b7160e01b600052602160045260246000fd5b60208101600483106146c157634e487b7160e01b600052602160045260246000fd5b91905290565b60005b838110156146e25781810151838201526020016146ca565b83811115612d3e5750506000910152565b6000815180845261470b8160208601602086016146c7565b601f01601f19169290920160200192915050565b602081526000611c7860208301846146f3565b8035613ed281614619565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156147765761477661473d565b60405290565b60405160e0810167ffffffffffffffff811182821017156147765761477661473d565b604051610100810167ffffffffffffffff811182821017156147765761477661473d565b604051601f8201601f1916810167ffffffffffffffff811182821017156147ec576147ec61473d565b604052919050565b600067ffffffffffffffff82111561480e5761480e61473d565b5060051b60200190565b600082601f83011261482957600080fd5b8135602061483e614839836147f4565b6147c3565b82815260059290921b8401810191818101908684111561485d57600080fd5b8286015b8481101561488157803561487481614619565b8352918301918301614861565b509695505050505050565b600082601f83011261489d57600080fd5b813560206148ad614839836147f4565b82815260059290921b840181019181810190868411156148cc57600080fd5b8286015b8481101561488157803583529183019183016148d0565b600080604083850312156148fa57600080fd5b823567ffffffffffffffff8082111561491257600080fd5b61491e86838701614818565b9350602085013591508082111561493457600080fd5b506149418582860161488c565b9150509250929050565b60006020828403121561495d57600080fd5b813567ffffffffffffffff81111561497457600080fd5b61498084828501614818565b949350505050565b6020808252825182820181905260009190848201906040850190845b818110156149c0578351835292840192918401916001016149a4565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156149c05783516001600160a01b0316835292840192918401916001016149e8565b60008060408385031215614a2057600080fd5b8235614a2b81614619565b91506020830135614a3b81614619565b809150509250929050565b600067ffffffffffffffff821115614a6057614a6061473d565b50601f01601f191660200190565b600082601f830112614a7f57600080fd5b8135614a8d61483982614a46565b818152846020838601011115614aa257600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112614ad057600080fd5b81356020614ae0614839836147f4565b82815260059290921b84018101918181019086841115614aff57600080fd5b8286015b84811015614881578035614b1681614619565b8352918301918301614b03565b600060408284031215614b3557600080fd5b614b3d614753565b90508135614b4a81614619565b808252506020820135602082015292915050565b600082601f830112614b6f57600080fd5b81356020614b7f614839836147f4565b82815260069290921b84018101918181019086841115614b9e57600080fd5b8286015b8481101561488157614bb48882614b23565b835291830191604001614ba2565b8015158114611eb557600080fd5b8035613ed281614bc2565b60008060408385031215614bee57600080fd5b823567ffffffffffffffff80821115614c0657600080fd5b9084019060e08287031215614c1a57600080fd5b614c2261477c565b614c2b83614661565b8152602083013582811115614c3f57600080fd5b614c4b88828601614a6e565b602083015250614c5d60408401614732565b6040820152606083013582811115614c7457600080fd5b614c8088828601614a6e565b606083015250608083013582811115614c9857600080fd5b614ca488828601614abf565b60808301525060a083013582811115614cbc57600080fd5b614cc888828601614b5e565b60a08301525060c083013560c0820152809450505050614cea60208401614bd0565b90509250929050565b600081518084526020808501945080840160005b83811015614d2c5781516001600160a01b031687529582019590820190600101614d07565b509495945050505050565b602081526000611c786020830184614cf3565b60008060408385031215614d5d57600080fd5b82359150602083013567ffffffffffffffff80821115614d7c57600080fd5b908401906101208287031215614d9157600080fd5b614d9961479f565b614da283614661565b8152614db060208401614661565b6020820152614dc160408401614732565b6040820152614dd260608401614732565b6060820152608083013582811115614de957600080fd5b614df588828601614a6e565b60808301525060a083013582811115614e0d57600080fd5b614e1988828601614b5e565b60a083015250614e2c8760c08501614b23565b60c082015261010083013560e08201528093505050509250929050565b600060408284031215614e5b57600080fd5b6040516040810181811067ffffffffffffffff82111715614e7e57614e7e61473d565b604052823581526020928301359281019290925250919050565b600060208284031215614eaa57600080fd5b5035919050565b60008083601f840112614ec357600080fd5b50813567ffffffffffffffff811115614edb57600080fd5b6020830191508360208260051b8501011115614ef657600080fd5b9250929050565b60008060008060008060008060e0898b031215614f1957600080fd5b606089018a811115614f2a57600080fd5b8998503567ffffffffffffffff80821115614f4457600080fd5b818b0191508b601f830112614f5857600080fd5b813581811115614f6757600080fd5b8c6020828501011115614f7957600080fd5b6020830199508098505060808b0135915080821115614f9757600080fd5b614fa38c838d01614eb1565b909750955060a08b0135915080821115614fbc57600080fd5b50614fc98b828c01614eb1565b999c989b50969995989497949560c00135949350505050565b600060808284031215614ff457600080fd5b6040516080810181811067ffffffffffffffff821117156150175761501761473d565b604052823563ffffffff8116811461502e57600080fd5b8152602083013561503e8161464b565b602082015260408301356150518161464b565b604082015260608301356150648161464b565b60608201529392505050565b600082601f83011261508157600080fd5b81356020615091614839836147f4565b82815260059290921b840181019181810190868411156150b057600080fd5b8286015b848110156148815780356150c78161464b565b83529183019183016150b4565b600082601f8301126150e557600080fd5b813560206150f5614839836147f4565b82815260059290921b8401810191818101908684111561511457600080fd5b8286015b8481101561488157803567ffffffffffffffff8111156151385760008081fd5b6151468986838b0101614a6e565b845250918301918301615118565b60006020828403121561516657600080fd5b813567ffffffffffffffff8082111561517e57600080fd5b90830190610100828603121561519357600080fd5b61519b61479f565b8235828111156151aa57600080fd5b6151b687828601615070565b8252506020830135828111156151cb57600080fd5b6151d787828601614abf565b6020830152506040830135828111156151ef57600080fd5b6151fb8782860161488c565b60408301525060608301358281111561521357600080fd5b61521f878286016150d4565b60608301525060808301358281111561523757600080fd5b6152438782860161488c565b60808301525060a083013560a082015260c08301358281111561526557600080fd5b6152718782860161488c565b60c08301525060e083013560e082015280935050505092915050565b803560ff81168114613ed257600080fd5b60008060008060008060c087890312156152b757600080fd5b863567ffffffffffffffff808211156152cf57600080fd5b6152db8a838b01614abf565b975060208901359150808211156152f157600080fd5b6152fd8a838b01614abf565b965061530b60408a0161528d565b9550606089013591508082111561532157600080fd5b61532d8a838b01614a6e565b945061533b60808a01614661565b935060a089013591508082111561535157600080fd5b5061535e89828a01614a6e565b9150509295509295509295565b60006020828403121561537d57600080fd5b813567ffffffffffffffff81111561539457600080fd5b820160e08185031215611c7857600080fd5b6000602082840312156153b857600080fd5b8151611c7881614bc2565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006000198203615402576154026153d9565b5060010190565b60008282101561541b5761541b6153d9565b500390565b634e487b7160e01b600052603160045260246000fd5b60006020828403121561544857600080fd5b8151611c7881614619565b600081600019048311821515161561546d5761546d6153d9565b500290565b60008219821115615485576154856153d9565b500190565b600060ff821660ff84168060ff038211156154a7576154a76153d9565b019392505050565b634e487b7160e01b600052601260045260246000fd5b600060ff8316806154d8576154d86154af565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b634e487b7160e01b600052600160045260246000fd5b600063ffffffff808316818516808303821115615548576155486153d9565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526155818184018a614cf3565b905082810360808401526155958189614cf3565b905060ff871660a084015282810360c08401526155b281876146f3565b905067ffffffffffffffff851660e08401528281036101008401526155d781856146f3565b9c9b505050505050505050505050565b6000826155f6576155f66154af565b500490565b600081518084526020808501945080840160005b83811015614d2c5761563587835180516001600160a01b03168252602090810151910152565b604096909601959082019060010161560f565b6040815267ffffffffffffffff83511660408201526000602084015160e060608401526156796101208401826146f3565b9050604085015161569560808501826001600160a01b03169052565b5060608501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0808584030160a08601526156d083836146f3565b925060808701519150808584030160c08601526156ed8383614cf3565b925060a08701519150808584030160e08601525061570b82826155fb565b60c087015161010086015285151560208601529250611c78915050565b8051613ed28161464b565b600082601f83011261574457600080fd5b81516020615754614839836147f4565b82815260059290921b8401810191818101908684111561577357600080fd5b8286015b8481101561488157805161578a8161464b565b8352918301918301615777565b8051613ed281614619565b600082601f8301126157b357600080fd5b815160206157c3614839836147f4565b82815260059290921b840181019181810190868411156157e257600080fd5b8286015b848110156148815780516157f981614619565b83529183019183016157e6565b600082601f83011261581757600080fd5b81516020615827614839836147f4565b82815260059290921b8401810191818101908684111561584657600080fd5b8286015b84811015614881578051835291830191830161584a565b600082601f83011261587257600080fd5b815161588061483982614a46565b81815284602083860101111561589557600080fd5b6149808260208301602087016146c7565b600082601f8301126158b757600080fd5b815160206158c7614839836147f4565b82815260059290921b840181019181810190868411156158e657600080fd5b8286015b8481101561488157805167ffffffffffffffff81111561590a5760008081fd5b6159188986838b0101615861565b8452509183019183016158ea565b60006020828403121561593857600080fd5b815167ffffffffffffffff8082111561595057600080fd5b90830190610100828603121561596557600080fd5b61596d61479f565b82518281111561597c57600080fd5b61598887828601615733565b82525060208301518281111561599d57600080fd5b6159a9878286016157a2565b6020830152506040830151828111156159c157600080fd5b6159cd87828601615806565b6040830152506060830151828111156159e557600080fd5b6159f1878286016158a6565b606083015250608083015182811115615a0957600080fd5b615a1587828601615806565b60808301525060a083015160a082015260c083015182811115615a3757600080fd5b615a4387828601615806565b60c08301525060e083015160e082015280935050505092915050565b600060408284031215615a7157600080fd5b615a79614753565b90508151615a8681614619565b808252506020820151602082015292915050565b600082601f830112615aab57600080fd5b81516020615abb614839836147f4565b82815260069290921b84018101918181019086841115615ada57600080fd5b8286015b8481101561488157615af08882615a5f565b835291830191604001615ade565b600060208284031215615b1057600080fd5b815167ffffffffffffffff80821115615b2857600080fd5b908301906101208286031215615b3d57600080fd5b615b4561479f565b615b4e83615728565b8152615b5c60208401615728565b6020820152615b6d60408401615797565b6040820152615b7e60608401615797565b6060820152608083015182811115615b9557600080fd5b615ba187828601615861565b60808301525060a083015182811115615bb957600080fd5b615bc587828601615a9a565b60a083015250615bd88660c08501615a5f565b60c0820152610100929092015160e0830152509392505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b166040850152816060850152615c2c8285018b614cf3565b91508382036080850152615c40828a614cf3565b915060ff881660a085015283820360c0850152615c5d82886146f3565b90861660e085015283810361010085015290506155d781856146f3565b602081526000611c7860208301846155fb565b600081518084526020808501945080840160005b83811015614d2c57815187529582019590820190600101615ca1565b60a081526000615cd060a0830188615c8d565b8281036020840152615ce28188615c8d565b90508560408401528281036060840152615cfc8186615c8d565b9150508260808301529695505050505050565b600060208284031215615d2157600080fd5b5051919050565b6000815160208301517fffffffff0000000000000000000000000000000000000000000000000000000080821693506004831015615d705780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
}

var EVM2EVMTollOffRampABI = EVM2EVMTollOffRampMetaData.ABI

var EVM2EVMTollOffRampBin = EVM2EVMTollOffRampMetaData.Bin

func DeployEVM2EVMTollOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId uint64, chainId uint64, offRampConfig IBaseOffRampOffRampConfig, onRampAddress common.Address, commitStore common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig IAggregateRateLimiterRateLimiterConfig, tokenLimitsAdmin common.Address) (common.Address, *types.Transaction, *EVM2EVMTollOffRamp, error) {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(IAggregateRateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(IAggregateRateLimiterTokenBucket)).(*IAggregateRateLimiterTokenBucket)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMTollOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMTollOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 InternalAny2EVMMessageFromSender) error {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) CcipReceive(arg0 InternalAny2EVMMessageFromSender) error {
	return _EVM2EVMTollOffRamp.Contract.CcipReceive(&_EVM2EVMTollOffRamp.CallOpts, arg0)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) CcipReceive(arg0 InternalAny2EVMMessageFromSender) error {
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

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ChainId = *abi.ConvertType(out[1], new(uint64)).(*uint64)

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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetConfig(opts *bind.CallOpts) (IBaseOffRampOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(IBaseOffRampOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IBaseOffRampOffRampConfig)).(*IBaseOffRampOffRampConfig)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetConfig() (IBaseOffRampOffRampConfig, error) {
	return _EVM2EVMTollOffRamp.Contract.GetConfig(&_EVM2EVMTollOffRamp.CallOpts)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetConfig() (IBaseOffRampOffRampConfig, error) {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getPoolByDestToken", destToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetPoolByDestToken(destToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPoolByDestToken(&_EVM2EVMTollOffRamp.CallOpts, destToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetPoolByDestToken(destToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPoolByDestToken(&_EVM2EVMTollOffRamp.CallOpts, destToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "getPoolBySourceToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPoolBySourceToken(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRamp.Contract.GetPoolBySourceToken(&_EVM2EVMTollOffRamp.CallOpts, sourceToken)
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCaller) OverheadGasToll(opts *bind.CallOpts, merkleGasShare *big.Int, message TollEVM2EVMTollMessage) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRamp.contract.Call(opts, &out, "overheadGasToll", merkleGasShare, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) OverheadGasToll(merkleGasShare *big.Int, message TollEVM2EVMTollMessage) (*big.Int, error) {
	return _EVM2EVMTollOffRamp.Contract.OverheadGasToll(&_EVM2EVMTollOffRamp.CallOpts, merkleGasShare, message)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampCallerSession) OverheadGasToll(merkleGasShare *big.Int, message TollEVM2EVMTollMessage) (*big.Int, error) {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMMessageFromSender, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "executeSingleMessage", message, manualExecution)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) ExecuteSingleMessage(message InternalAny2EVMMessageFromSender, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRamp.TransactOpts, message, manualExecution)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) ExecuteSingleMessage(message InternalAny2EVMMessageFromSender, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRamp.TransactOpts, message, manualExecution)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) ManuallyExecute(opts *bind.TransactOpts, report TollExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "manuallyExecute", report)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) ManuallyExecute(report TollExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.ManuallyExecute(&_EVM2EVMTollOffRamp.TransactOpts, report)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) ManuallyExecute(report TollExecutionReport) (*types.Transaction, error) {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetConfig(opts *bind.TransactOpts, config IBaseOffRampOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetConfig(config IBaseOffRampOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetConfig(config IBaseOffRampOffRampConfig) (*types.Transaction, error) {
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

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMTollOffRamp.TransactOpts, config)
}

func (_EVM2EVMTollOffRamp *EVM2EVMTollOffRampTransactorSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
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
	Config IBaseOffRampOffRampConfig
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
	SourceChainId uint64
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
	return common.HexToHash("0x052b5907be1d3ac35d571862117562e80ee743c01251e388dafb7dc4e92a726c")
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
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error)

	CcipReceive(opts *bind.CallOpts, arg0 InternalAny2EVMMessageFromSender) error

	FeeTaken(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

		error)

	GetCommitStore(opts *bind.CallOpts) (common.Address, error)

	GetConfig(opts *bind.CallOpts) (IBaseOffRampOffRampConfig, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error)

	GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	OverheadGasToll(opts *bind.CallOpts, merkleGasShare *big.Int, message TollEVM2EVMTollMessage) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMMessageFromSender, manualExecution bool) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report TollExecutionReport) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetCommitStore(opts *bind.TransactOpts, commitStore common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, config IBaseOffRampOffRampConfig) (*types.Transaction, error)

	SetConfig0(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error)

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
