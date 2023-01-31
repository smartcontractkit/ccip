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

type CommonAny2EVMMessage struct {
	SourceChainId        uint64
	Sender               []byte
	Data                 []byte
	DestTokensAndAmounts []CommonEVMTokenAndAmount
}

type CommonEVMTokenAndAmount struct {
	Token  common.Address
	Amount *big.Int
}

type IAggregateRateLimiterRateLimiterConfig struct {
	Admin    common.Address
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
	ExecutionDelaySeconds                   uint32
	MaxDataSize                             uint32
	MaxTokensLength                         uint16
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

var EVM2EVMTollOffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"executionDelaySeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"}],\"internalType\":\"structIBaseOffRamp.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractICommitStore\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"executionDelaySeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIBaseOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCommon.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structToll.ExecutionReport\",\"name\":\"rep\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount\",\"name\":\"feeTokenAndAmount\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structToll.EVM2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeTaken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitStore\",\"outputs\":[{\"internalType\":\"contractICommitStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"executionDelaySeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"}],\"internalType\":\"structIBaseOffRamp.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolByDestToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractIAny2EVMOffRampRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structToll.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"merkleGasShare\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount\",\"name\":\"feeTokenAndAmount\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structToll.EVM2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"overheadGasToll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICommitStore\",\"name\":\"commitStore\",\"type\":\"address\"}],\"name\":\"setCommitStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"executionDelaySeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"}],\"internalType\":\"structIBaseOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAny2EVMOffRampRouter\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620065a3380380620065a38339810160408190526200003491620007e5565b88888888888888888888888787878787878083838633806000806000806101000a81548160ff02191690831515021790555060006001600160a01b0316826001600160a01b031603620000ce5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620001085762000108816200046a565b5050506001600160a01b0381166200013357604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03929092169190911790558051825114620001755760405162d8548360e71b815260040160405180910390fd5b81516200018a9060059060208501906200051b565b5060005b8251811015620003135760006040518060400160405280848481518110620001ba57620001ba62000939565b60200260200101516001600160a01b03168152602001836001600160601b031681525090508060036000868581518110620001f957620001f962000939565b6020908102919091018101516001600160a01b0390811683528282019390935260409091016000908120845194909201516001600160601b0316600160a01b02939092169290921790915581518451909160049186908690811062000262576200026262000939565b60200260200101516001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015620002a8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002ce91906200094f565b6001600160a01b039081168252602082019290925260400160002080546001600160a01b03191692909116919091179055506200030b8162000976565b90506200018e565b50508151600680546001600160a01b0319166001600160a01b0392831617905560408051608081018252602080860151808352958301805191830182905251928201839052426060909201829052600995909555600a94909455600b55600c9290925550861662000397576040516342bcdf7f60e11b815260040160405180910390fd5b5050506001600160401b03948516608052509190921660a0526001600160a01b0391821660c052600e8054919092166001600160a01b031990911617905550508451601780546020880151604089015160609099015161ffff166c010000000000000000000000000261ffff60601b1963ffffffff9a8b1668010000000000000000021665ffffffffffff60401b19928b16640100000000026001600160401b03199094169a9095169990991791909117169190911795909517909455506200099e9d5050505050505050505050505050565b336001600160a01b03821603620004c45760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000c5565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b82805482825590600052602060002090810192821562000573579160200282015b828111156200057357825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200053c565b506200058192915062000585565b5090565b5b8082111562000581576000815560010162000586565b80516001600160401b0381168114620005b457600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715620005f457620005f4620005b9565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620006255762000625620005b9565b604052919050565b805163ffffffff81168114620005b457600080fd5b6001600160a01b03811681146200065857600080fd5b50565b8051620005b48162000642565b60006001600160401b03821115620006845762000684620005b9565b5060051b60200190565b600082601f830112620006a057600080fd5b81516020620006b9620006b38362000668565b620005fa565b82815260059290921b84018101918181019086841115620006d957600080fd5b8286015b8481101562000701578051620006f38162000642565b8352918301918301620006dd565b509695505050505050565b600082601f8301126200071e57600080fd5b8151602062000731620006b38362000668565b82815260059290921b840181019181810190868411156200075157600080fd5b8286015b84811015620007015780516200076b8162000642565b835291830191830162000755565b6000606082840312156200078c57600080fd5b604051606081016001600160401b0381118282101715620007b157620007b1620005b9565b80604052508091508251620007c68162000642565b8082525060208301516020820152604083015160408201525092915050565b6000806000806000806000806000898b036101c08112156200080657600080fd5b620008118b6200059c565b99506200082160208c016200059c565b98506080603f19820112156200083657600080fd5b5062000841620005cf565b6200084f60408c016200062d565b81526200085f60608c016200062d565b60208201526200087260808c016200062d565b604082015260a08b015161ffff811681146200088d57600080fd5b60608201529650620008a260c08b016200065b565b9550620008b260e08b016200065b565b9450620008c36101008b016200065b565b6101208b01519094506001600160401b0380821115620008e257600080fd5b620008f08d838e016200068e565b94506101408c01519150808211156200090857600080fd5b50620009178c828d016200070c565b9250506200092a8b6101608c0162000779565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156200096257600080fd5b81516200096f8162000642565b9392505050565b6000600182016200099757634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c051615bb2620009f16000396000818161220e015261326301526000818161032d0152613242015260008181610308015281816121ea015281816132210152613a6e0152615bb26000f3fe608060405234801561001057600080fd5b50600436106102f45760003560e01c80638da5cb5b11610191578063c0d78655116100e3578063d3c7c2c711610097578063eb511dd411610071578063eb511dd4146107e1578063f2fde38b146107f4578063f358426f1461080757600080fd5b8063d3c7c2c71461079a578063d7e2bb50146107a2578063e4381140146107ce57600080fd5b8063c5a1d7f0116100c8578063c5a1d7f01461076c578063c903328414610774578063d30a364b1461078757600080fd5b8063c0d7865514610699578063c3f909d4146106ac57600080fd5b8063b0f479a111610145578063b57671661161011f578063b576716614610662578063b66f0efb14610675578063be9b03f11461068657600080fd5b8063b0f479a11461062b578063b1dc65a41461063c578063b4069b311461064f57600080fd5b80639129badf116101765780639129badf146105ca578063a8b640c1146105eb578063afcb95d71461060b57600080fd5b80638da5cb5b1461057957806390c2339b1461058f57600080fd5b80634352fa9f1161024a578063666cab8d116101fe57806379ba5097116101d857806379ba50971461053957806381ff7048146105415780638456cb591461057157600080fd5b8063666cab8d146104fc578063681fba1614610511578063744b92e21461052657600080fd5b8063599f64311161022f578063599f6431146104b45780635c975abb146104c55780635d86f141146104d057600080fd5b80634352fa9f146104815780634741062e1461049457600080fd5b8063181f5a77116102ac5780633015b91c116102865780633015b91c1461045857806339aa9264146104665780633f4ba83a1461047957600080fd5b8063181f5a77146103d75780631ef38174146104205780632222dd421461043357600080fd5b8063142a98fc116102dd578063142a98fc14610372578063147809b3146103ac57806317332f9b146103c457600080fd5b8063087ae6df146102f9578063108ee5fc1461035d575b600080fd5b6040805167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811682527f0000000000000000000000000000000000000000000000000000000000000000166020820152015b60405180910390f35b61037061036b3660046143bf565b61081a565b005b61039f6103803660046143fd565b67ffffffffffffffff166000908152600f602052604090205460ff1690565b6040516103549190614430565b6103b46108d1565b6040519015158152602001610354565b6103706103d23660046144f7565b61095e565b6104136040518060400160405280601881526020017f45564d3245564d546f6c6c4f666652616d7020312e302e30000000000000000081525081565b60405161035491906145b5565b61037061042e3660046146ea565b610a99565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610354565b6103706102f43660046147b7565b6103706104743660046143bf565b6110ff565b610370611136565b61037061048f3660046148b1565b611148565b6104a76104a2366004614915565b61139d565b6040516103549190614952565b6006546001600160a01b0316610440565b60005460ff166103b4565b6104406104de3660046143bf565b6001600160a01b039081166000908152600360205260409020541690565b610504611465565b60405161035491906149da565b6105196114c7565b60405161035491906149ed565b610370610534366004614a2e565b61158c565b610370611904565b6012546010546040805163ffffffff80851682526401000000009094049093166020840152820152606001610354565b6103706119e7565b60005461010090046001600160a01b0316610440565b6105976119f7565b60405161035491908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6105dd6105d8366004614bd7565b611a98565b604051908152602001610354565b6105dd6105f9366004614c14565b60166020526000908152604090205481565b604080516001815260006020820181905291810191909152606001610354565b600d546001600160a01b0316610440565b61037061064a366004614c79565b611b49565b61044061065d3660046143bf565b61209f565b610370610670366004614d5e565b61218d565b600e546001600160a01b0316610440565b610370610694366004614f9d565b612199565b6103706106a73660046143bf565b6121a7565b61072b604080516080810182526000808252602082018190529181018290526060810191909152506040805160808101825260175463ffffffff808216835264010000000082048116602084015268010000000000000000820416928201929092526c0100000000000000000000000090910461ffff16606082015290565b60408051825163ffffffff908116825260208085015182169083015283830151169181019190915260609182015161ffff1691810191909152608001610354565b6105dd612265565b6103706107823660046143bf565b612295565b610370610795366004614fe4565b6122cc565b6105196122d7565b6104406107b03660046143bf565b6001600160a01b039081166000908152600460205260409020541690565b6103706107dc36600461502d565b612337565b6103706107ef366004614a2e565b612458565b6103706108023660046143bf565b6126d4565b6103706108153660046150b3565b6126e5565b6108226128a8565b6001600160a01b038116610862576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa158015610934573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095891906150e9565b15905090565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561099b57506006546001600160a01b03163314155b156109d2576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602081015179ffffffffffffffffffffffffffffffffffffffffffffffffffff11610a29576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a336009612907565b6040810151600a8190556020820151600955600b54610a5291906129b4565b600b556040818101516020808401518351928352908201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b855185518560ff16601f831115610b11576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064015b60405180910390fd5b80600003610b7b576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610b08565b818314610c09576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610b08565b610c1481600361511c565b8311610c7c576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610b08565b610c846128a8565b60145460005b81811015610d2c576013600060148381548110610ca957610ca9615159565b60009182526020808320909101546001600160a01b031683528201929092526040018120805461ffff1916905560158054601392919084908110610cef57610cef615159565b60009182526020808320909101546001600160a01b031683528201929092526040019020805461ffff19169055610d258161516f565b9050610c8a565b50895160005b81811015610fc15760008c8281518110610d4e57610d4e615159565b6020026020010151905060006002811115610d6b57610d6b61441a565b6001600160a01b038216600090815260136020526040902054610100900460ff166002811115610d9d57610d9d61441a565b14610e04576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610b08565b6040805180820190915260ff8316815260208101600190526001600160a01b03821660009081526013602090815260409091208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610e6c57610e6c61441a565b021790555090505060008c8381518110610e8857610e88615159565b6020026020010151905060006002811115610ea557610ea561441a565b6001600160a01b038216600090815260136020526040902054610100900460ff166002811115610ed757610ed761441a565b14610f3e576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610b08565b6040805180820190915260ff8416815260208101600290526001600160a01b03821660009081526013602090815260409091208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610fa657610fa661441a565b0217905550905050505080610fba9061516f565b9050610d32565b508a51610fd59060149060208e0190614308565b508951610fe99060159060208d0190614308565b506011805460ff8381166101000261ffff19909216908c1617179055601280546110529146913091906000906110249063ffffffff166151a7565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168e8e8e8e8e8e6129ca565b6010600001819055506000601260049054906101000a900463ffffffff16905043601260046101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581601060000154601260009054906101000a900463ffffffff168f8f8f8f8f8f6040516110e9999897969594939291906151ca565b60405180910390a1505050505050505050505050565b6111076128a8565b6006805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b61113e6128a8565b611146612a57565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561118557506006546001600160a01b03163314155b156111bc576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151815181146111f8576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085460005b8181101561125257600760006008838154811061121d5761121d615159565b60009182526020808320909101546001600160a01b0316835282019290925260400181205561124b8161516f565b90506111fe565b5060005b8281101561138257600085828151811061127257611272615159565b6020026020010151905060006001600160a01b0316816001600160a01b0316036112c8576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8482815181106112da576112da615159565b602002602001015160076000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f8186848151811061133f5761133f615159565b60200260200101516040516113699291906001600160a01b03929092168252602082015260400190565b60405180910390a15061137b8161516f565b9050611256565b508351611396906008906020870190614308565b5050505050565b80516060908067ffffffffffffffff8111156113bb576113bb614458565b6040519080825280602002602001820160405280156113e4578160200160208202803683370190505b50915060005b8181101561145e576007600085838151811061140857611408615159565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000205483828151811061144357611443615159565b60209081029190910101526114578161516f565b90506113ea565b5050919050565b606060158054806020026020016040519081016040528092919081815260200182805480156114bd57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161149f575b5050505050905090565b60055460609067ffffffffffffffff8111156114e5576114e5614458565b60405190808252806020026020018201604052801561150e578160200160208202803683370190505b50905060005b6005548110156115885761154e6005828154811061153457611534615159565b6000918252602090912001546001600160a01b031661209f565b82828151811061156057611560615159565b6001600160a01b03909216602092830291909101909101526115818161516f565b9050611514565b5090565b6115946128a8565b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290611622576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816001600160a01b031681600001516001600160a01b031614611671576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005805490600090611684600184615260565b8154811061169457611694615159565b9060005260206000200160009054906101000a90046001600160a01b03169050600583602001516bffffffffffffffffffffffff16815481106116d9576116d9615159565b6000918252602090912001546001600160a01b031660056116fb600185615260565b8154811061170b5761170b615159565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600584602001516bffffffffffffffffffffffff168154811061175f5761175f615159565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558581015184841683526003909152604090912080546bffffffffffffffffffffffff90921674010000000000000000000000000000000000000000029190921617905560058054806117e9576117e9615277565b6001900381819060005260206000200160006101000a8154906001600160a01b030219169055905560046000856001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611853573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611877919061528d565b6001600160a01b03908116825260208083019390935260409182016000908120805473ffffffffffffffffffffffffffffffffffffffff1916905588821680825260038552838220919091558251908152908716928101929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b0316331461195e5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610b08565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6119ef6128a8565b611146612af3565b611a226040518060800160405280600081526020016000815260200160008152602001600081525090565b604080516080810182526009548152600a546020820152600b5491810191909152600c54606082018190524290600090611a5c9083615260565b60208401518451919250611a8891611a74908461511c565b8560400151611a8391906152aa565b6129b4565b6040840152506060820152919050565b6000808260800151518360a001515160206014611ab591906152aa565b611abf919061511c565b611aca9060866152aa565b611ad491906152aa565b90506000611ae360108361511c565b9050610a28611bbc8560a00151516001611afd91906152aa565b611b0990618aac61511c565b6156b8611b1689866152aa565b611b2091906152aa565b611b2a91906152aa565b611b3491906152aa565b611b3e91906152aa565b925050505b92915050565b60005a9050611b8d88888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612b7b92505050565b6040805160608101825260105480825260115460ff808216602085015261010090910416928201929092528a35918214611c005780516040517f93df584c000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610b08565b6040805183815260208d81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a1600281602001518260400151611c5b91906152c2565b611c6591906152fd565b611c709060016152c2565b60ff168714611cab576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b868514611ce4576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526013602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115611d2757611d2761441a565b6002811115611d3857611d3861441a565b9052509050600281602001516002811115611d5557611d5561441a565b148015611d8f57506015816000015160ff1681548110611d7757611d77615159565b6000918252602090912001546001600160a01b031633145b611dc5576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506000611dd386602061511c565b611dde89602061511c565b611dea8c6101446152aa565b611df491906152aa565b611dfe91906152aa565b9050368114611e42576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401610b08565b5060008a8a604051611e5592919061531f565b604051908190038120611e6c918e9060200161532f565b604051602081830303815290604052805190602001209050611e8c614376565b8860005b8181101561208e5760006001858a8460208110611eaf57611eaf615159565b611ebc91901a601b6152c2565b8f8f86818110611ece57611ece615159565b905060200201358e8e87818110611ee757611ee7615159565b9050602002013560405160008152602001604052604051611f24949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611f46573d6000803e3d6000fd5b505060408051601f198101516001600160a01b038116600090815260136020908152848220848601909552845460ff8082168652939750919550929392840191610100909104166002811115611f9e57611f9e61441a565b6002811115611faf57611faf61441a565b9052509050600181602001516002811115611fcc57611fcc61441a565b14612003576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051859060ff16601f811061201a5761201a615159565b602002015115612056576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600185826000015160ff16601f811061207157612071615159565b911515602090920201525061208790508161516f565b9050611e90565b505050505050505050505050505050565b6001600160a01b03808216600090815260036020526040812054909116806120f3576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015612162573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612186919061528d565b9392505050565b61219681612b7b565b50565b6121a38282612b95565b5050565b6121af6128a8565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038381169182179092556040805167ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681527f0000000000000000000000000000000000000000000000000000000000000000909316602084015290917f052b5907be1d3ac35d571862117562e80ee743c01251e388dafb7dc4e92a726c910160405180910390a250565b60006122907fb9b8993db34ae003b2aacdae4cdef2888717531ab95157174f8f0dbf076b5e5861321c565b905090565b61229d6128a8565b600e805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b612196816001612b95565b606060058054806020026020016040519081016040528092919081815260200182805480156114bd576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161149f575050505050905090565b61233f6128a8565b80516017805460208085018051604080880180516060808b01805161ffff9081166c01000000000000000000000000027fffffffffffffffffffffffffffffffffffff0000ffffffffffffffffffffffff63ffffffff9586166801000000000000000002167fffffffffffffffffffffffffffffffffffff000000000000ffffffffffffffff988616640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000909c169d86169d8e179b909b17979097169990991795909517909855825198895293518416948801949094529251909116918501919091525116908201527f8f362c1cfd3071646996aaf74f584c630b3859adcd2ee3a6393c460e1467567e90608001610a8e565b6124606128a8565b6001600160a01b038216158061247d57506001600160a01b038116155b156124b4576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529015612543576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038083168083526005546bffffffffffffffffffffffff90811660208086019182528785166000908152600382526040808220885194519095167401000000000000000000000000000000000000000002939096169290921790925583517f21df0da70000000000000000000000000000000000000000000000000000000081529351869460049492936321df0da79282870192819003870181865afa1580156125f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061261c919061528d565b6001600160a01b03908116825260208083019390935260409182016000908120805495831673ffffffffffffffffffffffffffffffffffffffff199687161790556005805460018101825591527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001805488831695168517905581519384528516918301919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b6126dc6128a8565b612196816132dc565b33301461271e576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516000808252602082019092528161275b565b60408051808201909152600080825260208201528152602001906001900390816127345790505b5060a0840151519091501561278c5761278961277f8460a001518560c00151613398565b8460600151613623565b90505b60608301516001600160a01b03163b15806127dc575060608301516127da906001600160a01b03167f3015b91c000000000000000000000000000000000000000000000000000000006137ce565b155b156127e657505050565b600d546001600160a01b0316624b61bb61280085846137ea565b848660e0015187606001516040518563ffffffff1660e01b815260040161282a9493929190615398565b6020604051808303816000875af1158015612849573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061286d91906150e9565b6128a3576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b60005461010090046001600160a01b031633146111465760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610b08565b60018101546002820154429114806129225750808260030154145b1561292b575050565b81600101548260020154111561296d576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082600301548261297f9190615260565b600184015484549192506129a691612997908461511c565b8560020154611a8391906152aa565b600284015550600390910155565b60008183106129c35781612186565b5090919050565b6000808a8a8a8a8a8a8a8a8a6040516020016129ee9998979695949392919061544a565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b60005460ff16612aa95760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610b08565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60005460ff1615612b465760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610b08565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612ad63390565b61219681806020019051810190612b9291906156d0565b60005b60005460ff1615612be85760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610b08565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612c3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c5f91906150e9565b15612c95576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d546001600160a01b0316612cd7576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060820151516000819003612d18576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115612d3357612d33614458565b604051908082528060200260200182016040528015612dba57816020015b612da760408051610100810182526000808252602080830182905282840182905260608084018390526080840181905260a084015283518085019094528184528301529060c08201908152602001600081525090565b815260200190600190039081612d515790505b50905060008267ffffffffffffffff811115612dd857612dd8614458565b604051908082528060200260200182016040528015612e01578160200160208202803683370190505b5090506000612e2f7fb9b8993db34ae003b2aacdae4cdef2888717531ab95157174f8f0dbf076b5e5861321c565b905060005b84811015612ec957600087606001518281518110612e5457612e54615159565b6020026020010151806020019051810190612e6f91906158a8565b9050612e7b818461388c565b848381518110612e8d57612e8d615159565b60200260200101818152505080858381518110612eac57612eac615159565b60200260200101819052505080612ec29061516f565b9050612e34565b50600080612eea8489608001518a60a001518b60c001518c60e00151613974565b601754919350915060009063ffffffff16612f058442615260565b11905060005b87811015613210576000878281518110612f2757612f27615159565b602002602001015190506000612f5a826020015167ffffffffffffffff166000908152600f602052604090205460ff1690565b90506002816003811115612f7057612f7061441a565b03612fb95760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610b08565b8a1580612fc35750835b80612fdf57506003816003811115612fdd57612fdd61441a565b145b613015576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61301e82613a6c565b6000808260038111156130335761303361441a565b14801561303e57508b155b156130975761305a8a5187613053919061599c565b8e85613bca565b9050600061306f8460c0015160000151613db7565b905061307c818330613e19565b818460c001516020018181516130929190615260565b905250505b60008260038111156130ab576130ab61441a565b146130e75760208084015167ffffffffffffffff16600090815260168252604090205460c085015190910180516130e3908390615260565b9052505b60208381015167ffffffffffffffff166000908152600f90915260408120805460ff1916600117905561311a848e613e99565b60208086015167ffffffffffffffff166000908152600f909152604090208054919250829160ff191660018360038111156131575761315761441a565b021790555060008360038111156131705761317061441a565b14801561318e5750600381600381111561318c5761318c61441a565b145b156131b55760208085015167ffffffffffffffff1660009081526016909152604090208290555b836020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf826040516131f39190614430565b60405180910390a250505050806132099061516f565b9050612f0b565b50505050505050505050565b6000817f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000006040516020016132bf949392919093845267ffffffffffffffff9283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b336001600160a01b038216036133345760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610b08565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b606060005b83518110156135375782600001516001600160a01b03168482815181106133c6576133c6615159565b6020026020010151600001516001600160a01b031603613527576000845167ffffffffffffffff8111156133fc576133fc614458565b60405190808252806020026020018201604052801561344157816020015b604080518082019091526000808252602082015281526020019060019003908161341a5790505b50905060005b85518110156134985785818151811061346257613462615159565b602002602001015182828151811061347c5761347c615159565b6020026020010181905250806134919061516f565b9050613447565b5060405180604001604052808284815181106134b6576134b6615159565b6020026020010151600001516001600160a01b0316815260200185602001518385815181106134e7576134e7615159565b6020026020010151602001516134fd91906152aa565b81525081838151811061351257613512615159565b60200260200101819052508092505050611b43565b6135308161516f565b905061339d565b5060008351600161354891906152aa565b67ffffffffffffffff81111561356057613560614458565b6040519080825280602002602001820160405280156135a557816020015b604080518082019091526000808252602082015281526020019060019003908161357e5790505b50905060005b84518110156135fc578481815181106135c6576135c6615159565b60200260200101518282815181106135e0576135e0615159565b6020026020010181905250806135f59061516f565b90506135ab565b50828185518151811061361157613611615159565b60209081029190910101529392505050565b60606000835167ffffffffffffffff81111561364157613641614458565b60405190808252806020026020018201604052801561368657816020015b604080518082019091526000808252602082015281526020019060019003908161365f5790505b50905060005b84518110156137c45760006136bd8683815181106136ac576136ac615159565b602002602001015160000151613db7565b90506136e7818784815181106136d5576136d5615159565b60200260200101516020015187613e19565b806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613725573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613749919061528d565b83838151811061375b5761375b615159565b60209081029190910101516001600160a01b039091169052855186908390811061378757613787615159565b6020026020010151602001518383815181106137a5576137a5615159565b6020908102919091018101510152506137bd8161516f565b905061368c565b5061218681613fd1565b60006137d9836141d5565b801561218657506121868383614239565b61381f6040518060800160405280600067ffffffffffffffff1681526020016060815260200160608152602001606081525090565b6040518060800160405280846000015167ffffffffffffffff168152602001846040015160405160200161386291906001600160a01b0391909116815260200190565b60405160208183030381529060405281526020018460800151815260200183815250905092915050565b60008060001b828460200151856040015186606001518760800151805190602001208860a001516040516020016138c391906159b0565b604051602081830303815290604052805190602001208960e001518a60c00151604051602001613956999897969594939291909889526020808a019890985267ffffffffffffffff9690961660408901526001600160a01b039485166060890152928416608088015260a087019190915260c086015260e085015281511661010084015201516101208201526101400190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600e546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce906139d1908c908c908c908c908c906004016159f3565b6020604051808303816000875af11580156139f0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a149190615a45565b905060008111613a50576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a613a5c9084615260565b9350935050509550959350505050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16816000015167ffffffffffffffff1614613aec5780516040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610b08565b60175460a0820151516c0100000000000000000000000090910461ffff161015613b545760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610b08565b6017546080820151516801000000000000000090910463ffffffff161015612196576017546080820151516040517f869337890000000000000000000000000000000000000000000000000000000081526801000000000000000090920463ffffffff1660048301526024820152604401610b08565b6000806000613be08460c0015160000151613db7565b6001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613c1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c41919061528d565b905060005b856020015151811015613cba57816001600160a01b031686602001518281518110613c7357613c73615159565b60200260200101516001600160a01b031603613caa5785604001518181518110613c9f57613c9f615159565b602002602001015192505b613cb38161516f565b9050613c46565b5081613cfd576040517fce480bcc0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401610b08565b6000670de0b6b3a7640000833a8760e00151613d198b8a611a98565b613d2391906152aa565b613d2d919061511c565b613d37919061511c565b613d41919061599c565b90508460c0015160200151811115613dad5760208086015160c0870151909101516040517f3cab2f4d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9092166004830152602482018390526044820152606401610b08565b9695505050505050565b6001600160a01b038181166000908152600360205260409020541680613e14576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610b08565b919050565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b158015613e7c57600080fd5b505af1158015613e90573d6000803e3d6000fd5b50505050505050565b6040517ff358426f000000000000000000000000000000000000000000000000000000008152600090309063f358426f90613eda9086908690600401615a5e565b600060405180830381600087803b158015613ef457600080fd5b505af1925050508015613f05575060015b613fc8573d808015613f33576040519150601f19603f3d011682016040523d82523d6000602084013e613f38565b606091505b50613f4281615b55565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003613f94576003915050611b43565b806040517fcf19edfd000000000000000000000000000000000000000000000000000000008152600401610b0891906145b5565b50600292915050565b6000805b82518110156140d057600060076000858481518110613ff657613ff6615159565b6020026020010151600001516001600160a01b03166001600160a01b03168152602001908152602001600020549050806000036140895783828151811061403f5761403f615159565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610b08565b83828151811061409b5761409b615159565b602002602001015160200151816140b2919061511c565b6140bc90846152aa565b925050806140c99061516f565b9050613fd5565b5080156121a3576140e16009612907565b600a5481111561412b57600a546040517f688ccf77000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610b08565b600b5481111561418b57600954600b54600091906141499084615260565b614153919061599c565b9050806040517fe31e0f32000000000000000000000000000000000000000000000000000000008152600401610b0891815260200190565b80600960020160008282546141a09190615260565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378906020016108c5565b6000614201827f01ffc9a700000000000000000000000000000000000000000000000000000000614239565b8015611b435750614232827fffffffff00000000000000000000000000000000000000000000000000000000614239565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d915060005190508280156142f1575060208210155b80156142fd5750600081115b979650505050505050565b82805482825590600052602060002090810192821561436a579160200282015b8281111561436a578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190614328565b50611588929150614395565b604051806103e00160405280601f906020820280368337509192915050565b5b808211156115885760008155600101614396565b6001600160a01b038116811461219657600080fd5b6000602082840312156143d157600080fd5b8135612186816143aa565b67ffffffffffffffff8116811461219657600080fd5b8035613e14816143dc565b60006020828403121561440f57600080fd5b8135612186816143dc565b634e487b7160e01b600052602160045260246000fd5b602081016004831061445257634e487b7160e01b600052602160045260246000fd5b91905290565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561449157614491614458565b60405290565b604051610100810167ffffffffffffffff8111828210171561449157614491614458565b604051601f8201601f1916810167ffffffffffffffff811182821017156144e4576144e4614458565b604052919050565b8035613e14816143aa565b60006060828403121561450957600080fd5b6040516060810181811067ffffffffffffffff8211171561452c5761452c614458565b604052823561453a816143aa565b8152602083810135908201526040928301359281019290925250919050565b60005b8381101561457457818101518382015260200161455c565b83811115614583576000848401525b50505050565b600081518084526145a1816020860160208601614559565b601f01601f19169290920160200192915050565b6020815260006121866020830184614589565b600067ffffffffffffffff8211156145e2576145e2614458565b5060051b60200190565b600082601f8301126145fd57600080fd5b8135602061461261460d836145c8565b6144bb565b82815260059290921b8401810191818101908684111561463157600080fd5b8286015b84811015614655578035614648816143aa565b8352918301918301614635565b509695505050505050565b803560ff81168114613e1457600080fd5b600067ffffffffffffffff82111561468b5761468b614458565b50601f01601f191660200190565b600082601f8301126146aa57600080fd5b81356146b861460d82614671565b8181528460208386010111156146cd57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c0878903121561470357600080fd5b863567ffffffffffffffff8082111561471b57600080fd5b6147278a838b016145ec565b9750602089013591508082111561473d57600080fd5b6147498a838b016145ec565b965061475760408a01614660565b9550606089013591508082111561476d57600080fd5b6147798a838b01614699565b945061478760808a016143f2565b935060a089013591508082111561479d57600080fd5b506147aa89828a01614699565b9150509295509295509295565b6000602082840312156147c957600080fd5b813567ffffffffffffffff8111156147e057600080fd5b82016080818503121561218657600080fd5b600082601f83011261480357600080fd5b8135602061481361460d836145c8565b82815260059290921b8401810191818101908684111561483257600080fd5b8286015b84811015614655578035614849816143aa565b8352918301918301614836565b600082601f83011261486757600080fd5b8135602061487761460d836145c8565b82815260059290921b8401810191818101908684111561489657600080fd5b8286015b84811015614655578035835291830191830161489a565b600080604083850312156148c457600080fd5b823567ffffffffffffffff808211156148dc57600080fd5b6148e8868387016147f2565b935060208501359150808211156148fe57600080fd5b5061490b85828601614856565b9150509250929050565b60006020828403121561492757600080fd5b813567ffffffffffffffff81111561493e57600080fd5b61494a848285016147f2565b949350505050565b6020808252825182820181905260009190848201906040850190845b8181101561498a5783518352928401929184019160010161496e565b50909695505050505050565b600081518084526020808501945080840160005b838110156149cf5781516001600160a01b0316875295820195908201906001016149aa565b509495945050505050565b6020815260006121866020830184614996565b6020808252825182820181905260009190848201906040850190845b8181101561498a5783516001600160a01b031683529284019291840191600101614a09565b60008060408385031215614a4157600080fd5b8235614a4c816143aa565b91506020830135614a5c816143aa565b809150509250929050565b600060408284031215614a7957600080fd5b614a8161446e565b90508135614a8e816143aa565b808252506020820135602082015292915050565b600082601f830112614ab357600080fd5b81356020614ac361460d836145c8565b82815260069290921b84018101918181019086841115614ae257600080fd5b8286015b8481101561465557614af88882614a67565b835291830191604001614ae6565b60006101208284031215614b1957600080fd5b614b21614497565b9050614b2c826143f2565b8152614b3a602083016143f2565b6020820152614b4b604083016144ec565b6040820152614b5c606083016144ec565b6060820152608082013567ffffffffffffffff80821115614b7c57600080fd5b614b8885838601614699565b608084015260a0840135915080821115614ba157600080fd5b50614bae84828501614aa2565b60a083015250614bc18360c08401614a67565b60c082015261010082013560e082015292915050565b60008060408385031215614bea57600080fd5b82359150602083013567ffffffffffffffff811115614c0857600080fd5b61490b85828601614b06565b600060208284031215614c2657600080fd5b5035919050565b60008083601f840112614c3f57600080fd5b50813567ffffffffffffffff811115614c5757600080fd5b6020830191508360208260051b8501011115614c7257600080fd5b9250929050565b60008060008060008060008060e0898b031215614c9557600080fd5b606089018a811115614ca657600080fd5b8998503567ffffffffffffffff80821115614cc057600080fd5b818b0191508b601f830112614cd457600080fd5b813581811115614ce357600080fd5b8c6020828501011115614cf557600080fd5b6020830199508098505060808b0135915080821115614d1357600080fd5b614d1f8c838d01614c2d565b909750955060a08b0135915080821115614d3857600080fd5b50614d458b828c01614c2d565b999c989b50969995989497949560c00135949350505050565b600060208284031215614d7057600080fd5b813567ffffffffffffffff811115614d8757600080fd5b61494a84828501614699565b600082601f830112614da457600080fd5b81356020614db461460d836145c8565b82815260059290921b84018101918181019086841115614dd357600080fd5b8286015b84811015614655578035614dea816143dc565b8352918301918301614dd7565b600082601f830112614e0857600080fd5b81356020614e1861460d836145c8565b82815260059290921b84018101918181019086841115614e3757600080fd5b8286015b8481101561465557803567ffffffffffffffff811115614e5b5760008081fd5b614e698986838b0101614699565b845250918301918301614e3b565b60006101008284031215614e8a57600080fd5b614e92614497565b9050813567ffffffffffffffff80821115614eac57600080fd5b614eb885838601614d93565b83526020840135915080821115614ece57600080fd5b614eda858386016145ec565b60208401526040840135915080821115614ef357600080fd5b614eff85838601614856565b60408401526060840135915080821115614f1857600080fd5b614f2485838601614df7565b60608401526080840135915080821115614f3d57600080fd5b614f4985838601614856565b608084015260a084013560a084015260c0840135915080821115614f6c57600080fd5b50614f7984828501614856565b60c08301525060e082013560e082015292915050565b801515811461219657600080fd5b60008060408385031215614fb057600080fd5b823567ffffffffffffffff811115614fc757600080fd5b614fd385828601614e77565b9250506020830135614a5c81614f8f565b600060208284031215614ff657600080fd5b813567ffffffffffffffff81111561500d57600080fd5b61494a84828501614e77565b803563ffffffff81168114613e1457600080fd5b60006080828403121561503f57600080fd5b6040516080810181811067ffffffffffffffff8211171561506257615062614458565b60405261506e83615019565b815261507c60208401615019565b602082015261508d60408401615019565b6040820152606083013561ffff811681146150a757600080fd5b60608201529392505050565b600080604083850312156150c657600080fd5b823567ffffffffffffffff8111156150dd57600080fd5b614fd385828601614b06565b6000602082840312156150fb57600080fd5b815161218681614f8f565b634e487b7160e01b600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561515457615154615106565b500290565b634e487b7160e01b600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036151a0576151a0615106565b5060010190565b600063ffffffff8083168181036151c0576151c0615106565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526151fa8184018a614996565b9050828103608084015261520e8189614996565b905060ff871660a084015282810360c084015261522b8187614589565b905067ffffffffffffffff851660e08401528281036101008401526152508185614589565b9c9b505050505050505050505050565b60008282101561527257615272615106565b500390565b634e487b7160e01b600052603160045260246000fd5b60006020828403121561529f57600080fd5b8151612186816143aa565b600082198211156152bd576152bd615106565b500190565b600060ff821660ff84168060ff038211156152df576152df615106565b019392505050565b634e487b7160e01b600052601260045260246000fd5b600060ff831680615310576153106152e7565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b600081518084526020808501945080840160005b838110156149cf5761538587835180516001600160a01b03168252602090810151910152565b604096909601959082019060010161535f565b6080815267ffffffffffffffff855116608082015260006020860151608060a08401526153c9610100840182614589565b905060408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80808584030160c08601526154058383614589565b925060608901519150808584030160e086015250615423828261534b565b961515602085015250505060408101929092526001600160a01b0316606090910152919050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526154848285018b614996565b91508382036080850152615498828a614996565b915060ff881660a085015283820360c08501526154b58288614589565b90861660e085015283810361010085015290506152508185614589565b8051613e14816143dc565b600082601f8301126154ee57600080fd5b815160206154fe61460d836145c8565b82815260059290921b8401810191818101908684111561551d57600080fd5b8286015b84811015614655578051615534816143dc565b8352918301918301615521565b8051613e14816143aa565b600082601f83011261555d57600080fd5b8151602061556d61460d836145c8565b82815260059290921b8401810191818101908684111561558c57600080fd5b8286015b848110156146555780516155a3816143aa565b8352918301918301615590565b600082601f8301126155c157600080fd5b815160206155d161460d836145c8565b82815260059290921b840181019181810190868411156155f057600080fd5b8286015b8481101561465557805183529183019183016155f4565b600082601f83011261561c57600080fd5b815161562a61460d82614671565b81815284602083860101111561563f57600080fd5b61494a826020830160208701614559565b600082601f83011261566157600080fd5b8151602061567161460d836145c8565b82815260059290921b8401810191818101908684111561569057600080fd5b8286015b8481101561465557805167ffffffffffffffff8111156156b45760008081fd5b6156c28986838b010161560b565b845250918301918301615694565b6000602082840312156156e257600080fd5b815167ffffffffffffffff808211156156fa57600080fd5b90830190610100828603121561570f57600080fd5b615717614497565b82518281111561572657600080fd5b615732878286016154dd565b82525060208301518281111561574757600080fd5b6157538782860161554c565b60208301525060408301518281111561576b57600080fd5b615777878286016155b0565b60408301525060608301518281111561578f57600080fd5b61579b87828601615650565b6060830152506080830151828111156157b357600080fd5b6157bf878286016155b0565b60808301525060a083015160a082015260c0830151828111156157e157600080fd5b6157ed878286016155b0565b60c08301525060e083015160e082015280935050505092915050565b60006040828403121561581b57600080fd5b61582361446e565b90508151615830816143aa565b808252506020820151602082015292915050565b600082601f83011261585557600080fd5b8151602061586561460d836145c8565b82815260069290921b8401810191818101908684111561588457600080fd5b8286015b848110156146555761589a8882615809565b835291830191604001615888565b6000602082840312156158ba57600080fd5b815167ffffffffffffffff808211156158d257600080fd5b9083019061012082860312156158e757600080fd5b6158ef614497565b6158f8836154d2565b8152615906602084016154d2565b602082015261591760408401615541565b604082015261592860608401615541565b606082015260808301518281111561593f57600080fd5b61594b8782860161560b565b60808301525060a08301518281111561596357600080fd5b61596f87828601615844565b60a0830152506159828660c08501615809565b60c0820152610100929092015160e0830152509392505050565b6000826159ab576159ab6152e7565b500490565b602081526000612186602083018461534b565b600081518084526020808501945080840160005b838110156149cf578151875295820195908201906001016159d7565b60a081526000615a0660a08301886159c3565b8281036020840152615a1881886159c3565b90508560408401528281036060840152615a3281866159c3565b9150508260808301529695505050505050565b600060208284031215615a5757600080fd5b5051919050565b6040815267ffffffffffffffff835116604082015260006020840151615a90606084018267ffffffffffffffff169052565b5060408401516001600160a01b03811660808401525060608401516001600160a01b03811660a084015250608084015161012060c0840152615ad6610160840182614589565b905060a08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08483030160e0850152615b11828261534b565b91505060c0850151615b3a61010085018280516001600160a01b03168252602090810151910152565b5060e085015161014084015283151560208401529050612186565b6000815160208301517fffffffff0000000000000000000000000000000000000000000000000000000080821693506004831015615b9d5780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
}

var EVM2EVMTollOffRampHelperABI = EVM2EVMTollOffRampHelperMetaData.ABI

var EVM2EVMTollOffRampHelperBin = EVM2EVMTollOffRampHelperMetaData.Bin

func DeployEVM2EVMTollOffRampHelper(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId uint64, chainId uint64, offRampConfig IBaseOffRampOffRampConfig, onRampAddress common.Address, commitStore common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig IAggregateRateLimiterRateLimiterConfig) (common.Address, *types.Transaction, *EVM2EVMTollOffRampHelper, error) {
	parsed, err := EVM2EVMTollOffRampHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMTollOffRampHelperBin), backend, sourceChainId, chainId, offRampConfig, onRampAddress, commitStore, afn, sourceTokens, pools, rateLimiterConfig)
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(IAggregateRateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(IAggregateRateLimiterTokenBucket)).(*IAggregateRateLimiterTokenBucket)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMTollOffRampHelper.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMTollOffRampHelper.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) CcipReceive(opts *bind.CallOpts, arg0 CommonAny2EVMMessage) error {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) CcipReceive(arg0 CommonAny2EVMMessage) error {
	return _EVM2EVMTollOffRampHelper.Contract.CcipReceive(&_EVM2EVMTollOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) CcipReceive(arg0 CommonAny2EVMMessage) error {
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

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ChainId = *abi.ConvertType(out[1], new(uint64)).(*uint64)

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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetConfig(opts *bind.CallOpts) (IBaseOffRampOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(IBaseOffRampOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IBaseOffRampOffRampConfig)).(*IBaseOffRampOffRampConfig)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetConfig() (IBaseOffRampOffRampConfig, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetConfig(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetConfig() (IBaseOffRampOffRampConfig, error) {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getPoolByDestToken", destToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetPoolByDestToken(destToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPoolByDestToken(&_EVM2EVMTollOffRampHelper.CallOpts, destToken)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetPoolByDestToken(destToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPoolByDestToken(&_EVM2EVMTollOffRampHelper.CallOpts, destToken)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getPoolBySourceToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPoolBySourceToken(&_EVM2EVMTollOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPoolBySourceToken(&_EVM2EVMTollOffRampHelper.CallOpts, sourceToken)
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetSupportedTokens(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetSupportedTokens(&_EVM2EVMTollOffRampHelper.CallOpts)
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetTransmitters() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetTransmitters(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetTransmitters() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetTransmitters(&_EVM2EVMTollOffRampHelper.CallOpts)
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) OverheadGasToll(opts *bind.CallOpts, merkleGasShare *big.Int, message TollEVM2EVMTollMessage) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "overheadGasToll", merkleGasShare, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) OverheadGasToll(merkleGasShare *big.Int, message TollEVM2EVMTollMessage) (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.OverheadGasToll(&_EVM2EVMTollOffRampHelper.CallOpts, merkleGasShare, message)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) OverheadGasToll(merkleGasShare *big.Int, message TollEVM2EVMTollMessage) (*big.Int, error) {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) Execute(opts *bind.TransactOpts, rep TollExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "execute", rep, manualExecution)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Execute(rep TollExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Execute(&_EVM2EVMTollOffRampHelper.TransactOpts, rep, manualExecution)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) Execute(rep TollExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Execute(&_EVM2EVMTollOffRampHelper.TransactOpts, rep, manualExecution)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message TollEVM2EVMTollMessage, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "executeSingleMessage", message, manualExecution)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) ExecuteSingleMessage(message TollEVM2EVMTollMessage, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRampHelper.TransactOpts, message, manualExecution)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) ExecuteSingleMessage(message TollEVM2EVMTollMessage, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRampHelper.TransactOpts, message, manualExecution)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) ManuallyExecute(opts *bind.TransactOpts, report TollExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "manuallyExecute", report)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) ManuallyExecute(report TollExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.ManuallyExecute(&_EVM2EVMTollOffRampHelper.TransactOpts, report)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) ManuallyExecute(report TollExecutionReport) (*types.Transaction, error) {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetConfig(opts *bind.TransactOpts, config IBaseOffRampOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetConfig(config IBaseOffRampOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig(&_EVM2EVMTollOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetConfig(config IBaseOffRampOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig(&_EVM2EVMTollOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setOCR2Config", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetOCR2Config(&_EVM2EVMTollOffRampHelper.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetOCR2Config(&_EVM2EVMTollOffRampHelper.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetRateLimiterConfig(&_EVM2EVMTollOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
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
	Config IBaseOffRampOffRampConfig
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
	Router        common.Address
	SourceChainId uint64
	OnRampAddress common.Address
	Raw           types.Log
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
	return common.HexToHash("0x8f362c1cfd3071646996aaf74f584c630b3859adcd2ee3a6393c460e1467567e")
}

func (EVM2EVMTollOffRampHelperOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x052b5907be1d3ac35d571862117562e80ee743c01251e388dafb7dc4e92a726c")
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
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error)

	CcipReceive(opts *bind.CallOpts, arg0 CommonAny2EVMMessage) error

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

	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	MetadataHash(opts *bind.CallOpts) ([32]byte, error)

	OverheadGasToll(opts *bind.CallOpts, merkleGasShare *big.Int, message TollEVM2EVMTollMessage) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, rep TollExecutionReport, manualExecution bool) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message TollEVM2EVMTollMessage, manualExecution bool) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report TollExecutionReport) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetCommitStore(opts *bind.TransactOpts, commitStore common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, config IBaseOffRampOffRampConfig) (*types.Transaction, error)

	SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error)

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
