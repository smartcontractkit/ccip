// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_offramp_helper

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

type ClientAny2EVMMessage struct {
	MessageId        [32]byte
	SourceChainId    uint64
	Sender           []byte
	Data             []byte
	DestTokenAmounts []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
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

type IEVM2EVMOffRampOffRampConfig struct {
	PermissionLessExecutionThresholdSeconds uint32
	ExecutionDelaySeconds                   uint64
	Router                                  common.Address
	MaxDataSize                             uint32
	CommitStore                             common.Address
	MaxTokensLength                         uint16
}

type InternalEVM2EVMMessage struct {
	SourceChainId  uint64
	SequenceNumber uint64
	FeeTokenAmount *big.Int
	Sender         common.Address
	Nonce          uint64
	GasLimit       *big.Int
	Strict         bool
	Receiver       common.Address
	Data           []byte
	TokenAmounts   []ClientEVMTokenAmount
	FeeToken       common.Address
	MessageId      [32]byte
}

type InternalExecutionReport struct {
	SequenceNumbers []uint64
	EncodedMessages [][]byte
	Proofs          [][32]byte
	ProofFlagBits   *big.Int
}

var EVM2EVMOffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"InvalidOffRampConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OffRampConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"rep\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRampConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolByDestToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"releaseOrMintToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"releaseOrMintTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"setExecutionState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOffRampConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b506040516200655538038062006555833981016040819052620000359162000a31565b6000805460ff191681558890889088908890889088908890889081908490849087903390819081620000ae5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000e857620000e88162000429565b5050506001600160a01b0381166200011357604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b038316908117909155604080516000815260208101929092527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28910160405180910390a15080518251146200018f5760405162d8548360e71b815260040160405180910390fd5b8151620001a490600590602085019062000753565b5060005b82518110156200032d5760006040518060400160405280848481518110620001d457620001d462000ba0565b60200260200101516001600160a01b03168152602001836001600160601b03168152509050806003600086858151811062000213576200021362000ba0565b6020908102919091018101516001600160a01b0390811683528282019390935260409091016000908120845194909201516001600160601b0316600160a01b0293909216929092179091558151845190916004918690869081106200027c576200027c62000ba0565b60200260200101516001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015620002c2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002e8919062000bb6565b6001600160a01b039081168252602082019290925260400160002080546001600160a01b0319169290911691909117905550620003258162000bdd565b9050620001a8565b50508151600680546001600160a01b0319166001600160a01b03928316179055604080516080810182526020808601516001600160d01b0316808352958301805191830182905251928201839052426060909201829052600995909555600a94909455600b55600c92909255508616620003ba576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160401b03808916608052871660a0526001600160a01b03861660c052620004057fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a620004da565b60e052620004138562000540565b5050505050505050505050505050505062000c05565b336001600160a01b03821603620004835760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a5565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008160805160a05160c0516040516020016200052394939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b60408101516001600160a01b0316158062000566575060808101516001600160a01b0316155b15620005e1576040805163cd4e29b960e01b8152825163ffffffff908116600483015260208401516001600160401b03166024830152918301516001600160a01b039081166044830152606084015190921660648201526080830151909116608482015260a082015161ffff1660a482015260c401620000a5565b805160138054602084015160408086015163ffffffff9586166001600160601b0319909416939093176401000000006001600160401b039093168302176001600160601b03166c010000000000000000000000006001600160a01b03948516021790935560608501516014805460808089015160a08a0151949098166001600160c01b031990921691909117969094169092029490941761ffff60c01b1916600160c01b61ffff9095169490940293909317909255905160c05191517fe48449493f29cd3dc290c645fdc0077d7533294b25426b653f53500b328114389262000748928592909190835163ffffffff90811682526020808601516001600160401b03908116918401919091526040808701516001600160a01b03908116918501919091526060808801519093169284019290925260808087015183169084015260a09586015161ffff16958301959095529290931660c08401521660e08201526101000190565b60405180910390a150565b828054828255906000526020600020908101928215620007ab579160200282015b82811115620007ab57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000774565b50620007b9929150620007bd565b5090565b5b80821115620007b95760008155600101620007be565b80516001600160401b0381168114620007ec57600080fd5b919050565b6001600160a01b03811681146200080757600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b60405160c081016001600160401b03811182821017156200084557620008456200080a565b60405290565b604051601f8201601f191681016001600160401b03811182821017156200087657620008766200080a565b604052919050565b805163ffffffff81168114620007ec57600080fd5b8051620007ec81620007f1565b60006001600160401b03821115620008bc57620008bc6200080a565b5060051b60200190565b600082601f830112620008d857600080fd5b81516020620008f1620008eb83620008a0565b6200084b565b82815260059290921b840181019181810190868411156200091157600080fd5b8286015b84811015620009395780516200092b81620007f1565b835291830191830162000915565b509695505050505050565b600082601f8301126200095657600080fd5b8151602062000969620008eb83620008a0565b82815260059290921b840181019181810190868411156200098957600080fd5b8286015b8481101562000939578051620009a381620007f1565b83529183019183016200098d565b600060608284031215620009c457600080fd5b604051606081016001600160401b0381118282101715620009e957620009e96200080a565b80604052508091508251620009fe81620007f1565b815260208301516001600160d01b038116811462000a1b57600080fd5b6020820152604092830151920191909152919050565b600080600080600080600080888a036101e081121562000a5057600080fd5b62000a5b8a620007d4565b985062000a6b60208b01620007d4565b975060408a015162000a7d81620007f1565b965060c0605f198201121562000a9257600080fd5b5062000a9d62000820565b62000aab60608b016200087e565b815262000abb60808b01620007d4565b602082015260a08a015162000ad081620007f1565b604082015262000ae360c08b016200087e565b606082015260e08a015162000af881620007f1565b60808201526101008a015161ffff8116811462000b1457600080fd5b60a0820152945062000b2a6101208a0162000893565b6101408a01519094506001600160401b038082111562000b4957600080fd5b62000b578c838d01620008c6565b94506101608b015191508082111562000b6f57600080fd5b5062000b7e8b828c0162000944565b92505062000b918a6101808b01620009b1565b90509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b60006020828403121562000bc957600080fd5b815162000bd681620007f1565b9392505050565b60006001820162000bfe57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e0516158f362000c626000396000612f79015260008181612866015261372b015260008181610317015261370a0152600081816102f201528181612844015281816136e90152613b7001526158f36000f3fe608060405234801561001057600080fd5b50600436106102de5760003560e01c80637ee5053b11610186578063abc39f1f116100e3578063b576716611610097578063d7e2bb5011610071578063d7e2bb50146107dd578063eb511dd414610809578063f2fde38b1461081c57600080fd5b8063b5767166146107ac578063c5a1d7f0146107bf578063d3c7c2c7146107d557600080fd5b8063b1dc65a4116100c8578063b1dc65a414610773578063b3a18a3e14610786578063b4069b311461079957600080fd5b8063abc39f1f14610740578063afcb95d71461075357600080fd5b80638da5cb5b1161013a578063945b49931161011f578063945b499314610646578063966991da14610659578063a8ebd0f41461066c57600080fd5b80638da5cb5b146105f557806390c2339b1461060b57600080fd5b80638456cb591161016b5780638456cb591461059357806385572ffb1461059b578063856c8247146105a957600080fd5b80637ee5053b1461055057806381ff70481461056357600080fd5b80633f4ba83a1161023f5780635d86f141116101f3578063744b92e2116101cd578063744b92e21461051557806379ba5097146105285780637c2973611461053057600080fd5b80635d86f141146104bf578063666cab8d146104eb578063681fba161461050057600080fd5b80634741062e116102245780634741062e14610483578063599f6431146104a35780635c975abb146104b457600080fd5b80633f4ba83a146104685780634352fa9f1461047057600080fd5b80631790c413116102965780631ef381741161027b5780631ef381741461041d5780632222dd421461043057806339aa92641461045557600080fd5b80631790c413146103c1578063181f5a77146103d457600080fd5b8063142a98fc116102c7578063142a98fc1461035c578063147809b314610396578063167e3447146103ae57600080fd5b8063087ae6df146102e3578063108ee5fc14610347575b600080fd5b6040805167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811682527f0000000000000000000000000000000000000000000000000000000000000000166020820152015b60405180910390f35b61035a610355366004613fe7565b61082f565b005b61038961036a36600461402a565b67ffffffffffffffff1660009081526016602052604090205460ff1690565b60405161033e919061405d565b61039e6108e6565b604051901515815260200161033e565b61035a6103bc36600461415b565b610973565b61035a6103cf366004614209565b610987565b6104106040518060400160405280601481526020017f45564d3245564d4f666652616d7020312e302e3000000000000000000000000081525081565b60405161033e91906142a2565b61035a61042b3660046143d7565b6109c6565b6002546001600160a01b03165b6040516001600160a01b03909116815260200161033e565b61035a610463366004613fe7565b61102c565b61035a611063565b61035a61047e3660046144a4565b611075565b61049661049136600461455f565b6112ca565b60405161033e919061459c565b6006546001600160a01b031661043d565b60005460ff1661039e565b61043d6104cd366004613fe7565b6001600160a01b039081166000908152600360205260409020541690565b6104f3611392565b60405161033e9190614624565b6105086113f4565b60405161033e9190614637565b61035a610523366004614678565b6114b9565b61035a611831565b61054361053e36600461472e565b611914565b60405161033e91906147b9565b61035a61055e3660046147cc565b611929565b600f54600d546040805163ffffffff8085168252640100000000909404909316602084015282015260600161033e565b61035a611939565b61035a6102de36600461480e565b6105dc6105b7366004613fe7565b6001600160a01b031660009081526015602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff909116815260200161033e565b60005461010090046001600160a01b031661043d565b610613611949565b60405161033e91908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b61035a610654366004614a20565b6119e9565b61035a610667366004614a6e565b6119f4565b6107336040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260135463ffffffff808216835264010000000080830467ffffffffffffffff1660208501526c010000000000000000000000009092046001600160a01b039081169484019490945260145490811660608401529081049092166080820152780100000000000000000000000000000000000000000000000090910461ffff1660a082015290565b60405161033e9190614ab5565b61035a61074e366004614b1f565b611a02565b60408051600181526000602082018190529181019190915260600161033e565b61035a610781366004614caf565b611bc6565b61035a610794366004614d94565b612116565b61043d6107a7366004613fe7565b612232565b61035a6107ba366004614e1e565b612320565b6107c7612329565b60405190815260200161033e565b610508612359565b61043d6107eb366004613fe7565b6001600160a01b039081166000908152600460205260409020541690565b61035a610817366004614678565b6123b9565b61035a61082a366004613fe7565b612635565b610837612646565b6001600160a01b038116610877576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa158015610949573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096d9190614e5e565b15905090565b61097b612646565b610984816126a5565b50565b67ffffffffffffffff82166000908152601660205260409020805482919060ff191660018360038111156109bd576109bd614047565b02179055505050565b855185518560ff16601f831115610a3e576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064015b60405180910390fd5b80600003610aa8576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610a35565b818314610b36576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610a35565b610b41816003614e91565b8311610ba9576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610a35565b610bb1612646565b60115460005b81811015610c59576010600060118381548110610bd657610bd6614eb0565b60009182526020808320909101546001600160a01b031683528201929092526040018120805461ffff1916905560128054601092919084908110610c1c57610c1c614eb0565b60009182526020808320909101546001600160a01b031683528201929092526040019020805461ffff19169055610c5281614ec6565b9050610bb7565b50895160005b81811015610eee5760008c8281518110610c7b57610c7b614eb0565b6020026020010151905060006002811115610c9857610c98614047565b6001600160a01b038216600090815260106020526040902054610100900460ff166002811115610cca57610cca614047565b14610d31576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610a35565b6040805180820190915260ff8316815260208101600190526001600160a01b03821660009081526010602090815260409091208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610d9957610d99614047565b021790555090505060008c8381518110610db557610db5614eb0565b6020026020010151905060006002811115610dd257610dd2614047565b6001600160a01b038216600090815260106020526040902054610100900460ff166002811115610e0457610e04614047565b14610e6b576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610a35565b6040805180820190915260ff8416815260208101600290526001600160a01b03821660009081526010602090815260409091208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610ed357610ed3614047565b0217905550905050505080610ee790614ec6565b9050610c5f565b508a51610f029060119060208e0190613f30565b508951610f169060129060208d0190613f30565b50600e805460ff8381166101000261ffff19909216908c1617179055600f8054610f7f914691309190600090610f519063ffffffff16614ee0565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168e8e8e8e8e8e61288b565b600d600001819055506000600f60049054906101000a900463ffffffff16905043600f60046101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581600d60000154600f60009054906101000a900463ffffffff168f8f8f8f8f8f60405161101699989796959493929190614f03565b60405180910390a1505050505050505050505050565b611034612646565b6006805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b61106b612646565b611073612918565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b0316141580156110b257506006546001600160a01b03163314155b156110e9576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114611125576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085460005b8181101561117f57600760006008838154811061114a5761114a614eb0565b60009182526020808320909101546001600160a01b0316835282019290925260400181205561117881614ec6565b905061112b565b5060005b828110156112af57600085828151811061119f5761119f614eb0565b6020026020010151905060006001600160a01b0316816001600160a01b0316036111f5576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84828151811061120757611207614eb0565b602002602001015160076000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f8186848151811061126c5761126c614eb0565b60200260200101516040516112969291906001600160a01b03929092168252602082015260400190565b60405180910390a1506112a881614ec6565b9050611183565b5083516112c3906008906020870190613f30565b5050505050565b80516060908067ffffffffffffffff8111156112e8576112e8614085565b604051908082528060200260200182016040528015611311578160200160208202803683370190505b50915060005b8181101561138b576007600085838151811061133557611335614eb0565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000205483828151811061137057611370614eb0565b602090810291909101015261138481614ec6565b9050611317565b5050919050565b606060128054806020026020016040519081016040528092919081815260200182805480156113ea57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113cc575b5050505050905090565b60055460609067ffffffffffffffff81111561141257611412614085565b60405190808252806020026020018201604052801561143b578160200160208202803683370190505b50905060005b6005548110156114b55761147b6005828154811061146157611461614eb0565b6000918252602090912001546001600160a01b0316612232565b82828151811061148d5761148d614eb0565b6001600160a01b03909216602092830291909101909101526114ae81614ec6565b9050611441565b5090565b6114c1612646565b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff16908201529061154f576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816001600160a01b031681600001516001600160a01b03161461159e576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058054906000906115b1600184614f99565b815481106115c1576115c1614eb0565b9060005260206000200160009054906101000a90046001600160a01b03169050600583602001516bffffffffffffffffffffffff168154811061160657611606614eb0565b6000918252602090912001546001600160a01b03166005611628600185614f99565b8154811061163857611638614eb0565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600584602001516bffffffffffffffffffffffff168154811061168c5761168c614eb0565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558581015184841683526003909152604090912080546bffffffffffffffffffffffff909216740100000000000000000000000000000000000000000291909216179055600580548061171657611716614fb0565b6001900381819060005260206000200160006101000a8154906001600160a01b030219169055905560046000856001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611780573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117a49190614fc6565b6001600160a01b03908116825260208083019390935260409182016000908120805473ffffffffffffffffffffffffffffffffffffffff1916905588821680825260038552838220919091558251908152908716928101929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b0316331461188b5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610a35565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b606061192083836129b4565b90505b92915050565b611934838383612be2565b505050565b611941612646565b611073612c62565b6119746040518060800160405280600081526020016000815260200160008152602001600081525090565b604080516080810182526009548152600a546020820152600b5491810191909152600c54606082018190526000906119ac9042614f99565b602083015183519192506119d8916119c49084614e91565b84604001516119d39190614fe3565b612cea565b604083015250426060820152919050565b610984816001612d00565b6119fe8282612d00565b5050565b333014611a3b576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160008082526020820190925281611a78565b6040805180820190915260008082526020820152815260200190600190039081611a515790505b506101208401515190915015611a9e57611a9b8361012001518460e001516129b4565b90505b60e08301516001600160a01b03163b1580611aee575060e0830151611aec906001600160a01b03167f85572ffb0000000000000000000000000000000000000000000000000000000061355c565b155b15611af857505050565b6013546c0100000000000000000000000090046001600160a01b0316635607b375611b238584613578565b848660a001518760e001516040518563ffffffff1660e01b8152600401611b4d9493929190614ffb565b6020604051808303816000875af1158015611b6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b909190614e5e565b611934576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c0587878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061361b92505050565b60408051606081018252600d54808252600e5460ff808216602085015261010090910416928201929092528935918214611c785780516040517f93df584c000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610a35565b6040805183815260208c81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a1600281602001518260400151611cd391906150c7565b611cdd9190615102565b611ce89060016150c7565b60ff168614611d23576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b858414611d5c576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526010602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115611d9f57611d9f614047565b6002811115611db057611db0614047565b9052509050600281602001516002811115611dcd57611dcd614047565b148015611e0757506012816000015160ff1681548110611def57611def614eb0565b6000918252602090912001546001600160a01b031633145b611e3d576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506000611e4b856020614e91565b611e56886020614e91565b611e628b610144614fe3565b611e6c9190614fe3565b611e769190614fe3565b9050368114611eba576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401610a35565b5060008989604051611ecd929190615124565b604051908190038120611ee4918d90602001615134565b604051602081830303815290604052805190602001209050611f04613f9e565b8760005b81811015612106576000600185898460208110611f2757611f27614eb0565b611f3491901a601b6150c7565b8e8e86818110611f4657611f46614eb0565b905060200201358d8d87818110611f5f57611f5f614eb0565b9050602002013560405160008152602001604052604051611f9c949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611fbe573d6000803e3d6000fd5b505060408051601f198101516001600160a01b038116600090815260106020908152848220848601909552845460ff808216865293975091955092939284019161010090910416600281111561201657612016614047565b600281111561202757612027614047565b905250905060018160200151600281111561204457612044614047565b1461207b576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051859060ff16601f811061209257612092614eb0565b6020020151156120ce576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600185826000015160ff16601f81106120e9576120e9614eb0565b91151560209092020152506120ff905081614ec6565b9050611f08565b5050505050505050505050505050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561215357506006546001600160a01b03163314155b1561218a576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6121946009613639565b6040810151600a819055602082015179ffffffffffffffffffffffffffffffffffffffffffffffffffff16600955600b546121cf9190612cea565b600b55604081810151602080840151835192835279ffffffffffffffffffffffffffffffffffffffffffffffffffff16908201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b6001600160a01b0380821660009081526003602052604081205490911680612286576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa1580156122f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123199190614fc6565b9392505050565b6109848161361b565b60006123547fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a6136e4565b905090565b606060058054806020026020016040519081016040528092919081815260200182805480156113ea576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113cc575050505050905090565b6123c1612646565b6001600160a01b03821615806123de57506001600160a01b038116155b15612415576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156124a4576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038083168083526005546bffffffffffffffffffffffff90811660208086019182528785166000908152600382526040808220885194519095167401000000000000000000000000000000000000000002939096169290921790925583517f21df0da70000000000000000000000000000000000000000000000000000000081529351869460049492936321df0da79282870192819003870181865afa158015612559573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061257d9190614fc6565b6001600160a01b03908116825260208083019390935260409182016000908120805495831673ffffffffffffffffffffffffffffffffffffffff199687161790556005805460018101825591527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001805488831695168517905581519384528516918301919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b61263d612646565b610984816137a4565b60005461010090046001600160a01b031633146110735760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610a35565b60408101516001600160a01b031615806126ca575060808101516001600160a01b0316155b1561270357806040517fcd4e29b9000000000000000000000000000000000000000000000000000000008152600401610a359190614ab5565b805160138054602084015160408086015163ffffffff9586167fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009094169390931764010000000067ffffffffffffffff9093168302176bffffffffffffffffffffffff166c010000000000000000000000006001600160a01b039485160217909355606085015160148054608088015160a0890151939097167fffffffffffffffff0000000000000000000000000000000000000000000000009091161795909316909102939093177fffffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000061ffff9094169390930292909217909155517fe48449493f29cd3dc290c645fdc0077d7533294b25426b653f53500b32811438906122279083907f0000000000000000000000000000000000000000000000000000000000000000907f000000000000000000000000000000000000000000000000000000000000000090615150565b6000808a8a8a8a8a8a8a8a8a6040516020016128af999897969594939291906151e2565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b60005460ff1661296a5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610a35565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60606000835167ffffffffffffffff8111156129d2576129d2614085565b604051908082528060200260200182016040528015612a1757816020015b60408051808201909152600080825260208201528152602001906001900390816129f05790505b50905060005b8451811015612bd8576000612a67868381518110612a3d57612a3d614eb0565b6020026020010151600001516001600160a01b039081166000908152600360205260409020541690565b90506001600160a01b038116612ad357858281518110612a8957612a89614eb0565b6020908102919091010151516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a35565b612afb81878481518110612ae957612ae9614eb0565b60200260200101516020015187612be2565b806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612b39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b5d9190614fc6565b838381518110612b6f57612b6f614eb0565b60209081029190910101516001600160a01b0390911690528551869083908110612b9b57612b9b614eb0565b602002602001015160200151838381518110612bb957612bb9614eb0565b602090810291909101810151015250612bd181614ec6565b9050612a1d565b5061192081613860565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b158015612c4557600080fd5b505af1158015612c59573d6000803e3d6000fd5b50505050505050565b60005460ff1615612cb55760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610a35565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586129973390565b6000818310612cf95781611920565b5090919050565b60005460ff1615612d535760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610a35565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612da6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dca9190614e5e565b15612e00576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020820151516000819003612e40576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115612e5b57612e5b614085565b604051908082528060200260200182016040528015612e84578160200160208202803683370190505b50905060008267ffffffffffffffff811115612ea257612ea2614085565b604051908082528060200260200182016040528015612f3057816020015b60408051610180810182526000808252602080830182905292820181905260608083018290526080830182905260a0830182905260c0830182905260e08301829052610100830181905261012083015261014082018190526101608201528252600019909201910181612ec05790505b50905060005b83811015612feb57600086602001518281518110612f5657612f56614eb0565b6020026020010151806020019051810190612f71919061534d565b9050612f9d817f0000000000000000000000000000000000000000000000000000000000000000613a64565b848381518110612faf57612faf614eb0565b60200260200101818152505080838381518110612fce57612fce614eb0565b60200260200101819052505080612fe490614ec6565b9050612f36565b50601454604086810151606088015191517f3204887500000000000000000000000000000000000000000000000000000000815260009364010000000090046001600160a01b031692633204887592613049928892906004016154b0565b6020604051808303816000875af1158015613068573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061308c91906154e6565b9050600081116130c8576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b84811015612c595760008382815181106130e7576130e7614eb0565b60200260200101519050600061311a826020015167ffffffffffffffff1660009081526016602052604090205460ff1690565b9050600081600381111561313057613130614047565b148061314d5750600381600381111561314b5761314b614047565b145b6131955760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a35565b871561320d5760135460009063ffffffff166131b18642614f99565b11905080806131d1575060038260038111156131cf576131cf614047565b145b613207576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061326a565b600081600381111561322157613221614047565b1461326a5760208201516040517f67d9ba0f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a35565b600081600381111561327e5761327e614047565b0361331d57608082015160608301516001600160a01b031660009081526015602052604090205467ffffffffffffffff918216916132be911660016154ff565b67ffffffffffffffff161461331d5781606001516001600160a01b0316826080015167ffffffffffffffff167fd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf4123760405160405180910390a3505061354c565b61332682613b6e565b60208281015167ffffffffffffffff166000908152601690915260408120805460ff19166001179055613359838a613cc5565b60208085015167ffffffffffffffff1660009081526016909152604090208054919250829160ff1916600183600381111561339657613396614047565b0217905550881561347c578260c0015180156133c3575060038260038111156133c1576133c1614047565b145b80156133e0575060028160038111156133de576133de614047565b145b80613418575060008260038111156133fa576133fa614047565b1480156134185750600281600381111561341657613416614047565b145b156134775760608301516001600160a01b03166000908152601560205260408120805467ffffffffffffffff169161344f8361552b565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b6134fc565b8260c00151801561349e5750600381600381111561349c5761349c614047565b145b6134fc5760608301516001600160a01b03166000908152601560205260408120805467ffffffffffffffff16916134d48361552b565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b826101600151836020015167ffffffffffffffff167f5e04b4755a4460aa6de4f3a906c4324a025c7449c02b52f5466659b5bfdfba5f83604051613540919061405d565b60405180910390a35050505b61355581614ec6565b90506130cb565b600061356783613dfd565b801561192057506119208383613e61565b6040805160a08101825260008082526020820152606091810182905281810182905260808101919091526040518060a001604052808461016001518152602001846000015167ffffffffffffffff16815260200184606001516040516020016135f091906001600160a01b0391909116815260200190565b6040516020818303038152906040528152602001846101000151815260200183815250905092915050565b610984818060200190518101906136329190615623565b6000612d00565b8060010154816002015414806136525750428160030154145b1561365a5750565b80600101548160020154111561369c576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008160030154426136ae9190614f99565b600183015483549192506136d5916136c69084614e91565b84600201546119d39190614fe3565b60028301555042600390910155565b6000817f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000604051602001613787949392919093845267ffffffffffffffff9283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b336001600160a01b038216036137fc5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610a35565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000805b825181101561395f5760006007600085848151811061388557613885614eb0565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600003613918578382815181106138ce576138ce614eb0565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a35565b83828151811061392a5761392a614eb0565b602002602001015160200151816139419190614e91565b61394b9084614fe3565b9250508061395890614ec6565b9050613864565b5080156119fe576139706009613639565b600a548111156139ba57600a546040517f688ccf77000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610a35565b600b54811115613a1a57600954600b54600091906139d89084614f99565b6139e29190615741565b9050806040517fe31e0f32000000000000000000000000000000000000000000000000000000008152600401610a3591815260200190565b8060096002016000828254613a2f9190614f99565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378906020016108da565b60008060001b828460200151856080015186606001518760e0015188610100015180519060200120896101200151604051602001613aa291906147b9565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d60400151604051602001613b509c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16816000015167ffffffffffffffff1614613bee5780516040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a35565b60145461012082015151780100000000000000000000000000000000000000000000000090910461ffff161015613c635760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a35565b6014546101008201515163ffffffff909116101561098457601454610100820151516040517f8693378900000000000000000000000000000000000000000000000000000000815263ffffffff90921660048301526024820152604401610a35565b6040517fabc39f1f000000000000000000000000000000000000000000000000000000008152600090309063abc39f1f90613d069086908690600401615755565b600060405180830381600087803b158015613d2057600080fd5b505af1925050508015613d31575060015b613df4573d808015613d5f576040519150601f19603f3d011682016040523d82523d6000602084013e613d64565b606091505b50613d6e81615896565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003613dc0576003915050611923565b806040517fcf19edfd000000000000000000000000000000000000000000000000000000008152600401610a3591906142a2565b50600292915050565b6000613e29827f01ffc9a700000000000000000000000000000000000000000000000000000000613e61565b80156119235750613e5a827fffffffff00000000000000000000000000000000000000000000000000000000613e61565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015613f19575060208210155b8015613f255750600081115b979650505050505050565b828054828255906000526020600020908101928215613f92579160200282015b82811115613f92578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190613f50565b506114b5929150613fbd565b604051806103e00160405280601f906020820280368337509192915050565b5b808211156114b55760008155600101613fbe565b6001600160a01b038116811461098457600080fd5b600060208284031215613ff957600080fd5b813561192081613fd2565b67ffffffffffffffff8116811461098457600080fd5b803561402581614004565b919050565b60006020828403121561403c57600080fd5b813561192081614004565b634e487b7160e01b600052602160045260246000fd5b602081016004831061407f57634e487b7160e01b600052602160045260246000fd5b91905290565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156140be576140be614085565b60405290565b6040516080810167ffffffffffffffff811182821017156140be576140be614085565b604051610180810167ffffffffffffffff811182821017156140be576140be614085565b604051601f8201601f1916810167ffffffffffffffff8111828210171561413457614134614085565b604052919050565b803563ffffffff8116811461402557600080fd5b803561402581613fd2565b600060c0828403121561416d57600080fd5b60405160c0810181811067ffffffffffffffff8211171561419057614190614085565b60405261419c8361413c565b815260208301356141ac81614004565b602082015260408301356141bf81613fd2565b60408201526141d06060840161413c565b606082015260808301356141e381613fd2565b608082015260a083013561ffff811681146141fd57600080fd5b60a08201529392505050565b6000806040838503121561421c57600080fd5b823561422781614004565b915060208301356004811061423b57600080fd5b809150509250929050565b60005b83811015614261578181015183820152602001614249565b83811115614270576000848401525b50505050565b6000815180845261428e816020860160208601614246565b601f01601f19169290920160200192915050565b6020815260006119206020830184614276565b600067ffffffffffffffff8211156142cf576142cf614085565b5060051b60200190565b600082601f8301126142ea57600080fd5b813560206142ff6142fa836142b5565b61410b565b82815260059290921b8401810191818101908684111561431e57600080fd5b8286015b8481101561434257803561433581613fd2565b8352918301918301614322565b509695505050505050565b803560ff8116811461402557600080fd5b600067ffffffffffffffff82111561437857614378614085565b50601f01601f191660200190565b600082601f83011261439757600080fd5b81356143a56142fa8261435e565b8181528460208386010111156143ba57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c087890312156143f057600080fd5b863567ffffffffffffffff8082111561440857600080fd5b6144148a838b016142d9565b9750602089013591508082111561442a57600080fd5b6144368a838b016142d9565b965061444460408a0161434d565b9550606089013591508082111561445a57600080fd5b6144668a838b01614386565b945061447460808a0161401a565b935060a089013591508082111561448a57600080fd5b5061449789828a01614386565b9150509295509295509295565b600080604083850312156144b757600080fd5b823567ffffffffffffffff808211156144cf57600080fd5b6144db868387016142d9565b93506020915081850135818111156144f257600080fd5b85019050601f8101861361450557600080fd5b80356145136142fa826142b5565b81815260059190911b8201830190838101908883111561453257600080fd5b928401925b8284101561455057833582529284019290840190614537565b80955050505050509250929050565b60006020828403121561457157600080fd5b813567ffffffffffffffff81111561458857600080fd5b614594848285016142d9565b949350505050565b6020808252825182820181905260009190848201906040850190845b818110156145d4578351835292840192918401916001016145b8565b50909695505050505050565b600081518084526020808501945080840160005b838110156146195781516001600160a01b0316875295820195908201906001016145f4565b509495945050505050565b60208152600061192060208301846145e0565b6020808252825182820181905260009190848201906040850190845b818110156145d45783516001600160a01b031683529284019291840191600101614653565b6000806040838503121561468b57600080fd5b823561469681613fd2565b9150602083013561423b81613fd2565b600082601f8301126146b757600080fd5b813560206146c76142fa836142b5565b82815260069290921b840181019181810190868411156146e657600080fd5b8286015b8481101561434257604081890312156147035760008081fd5b61470b61409b565b813561471681613fd2565b815281850135858201528352918301916040016146ea565b6000806040838503121561474157600080fd5b823567ffffffffffffffff81111561475857600080fd5b614764858286016146a6565b925050602083013561423b81613fd2565b600081518084526020808501945080840160005b8381101561461957815180516001600160a01b031688528301518388015260409096019590820190600101614789565b6020815260006119206020830184614775565b6000806000606084860312156147e157600080fd5b83356147ec81613fd2565b925060208401359150604084013561480381613fd2565b809150509250925092565b60006020828403121561482057600080fd5b813567ffffffffffffffff81111561483757600080fd5b820160a0818503121561192057600080fd5b600082601f83011261485a57600080fd5b8135602061486a6142fa836142b5565b82815260059290921b8401810191818101908684111561488957600080fd5b8286015b8481101561434257803567ffffffffffffffff8111156148ad5760008081fd5b6148bb8986838b0101614386565b84525091830191830161488d565b600082601f8301126148da57600080fd5b813560206148ea6142fa836142b5565b82815260059290921b8401810191818101908684111561490957600080fd5b8286015b84811015614342578035835291830191830161490d565b60006080828403121561493657600080fd5b61493e6140c4565b9050813567ffffffffffffffff8082111561495857600080fd5b818401915084601f83011261496c57600080fd5b8135602061497c6142fa836142b5565b82815260059290921b8401810191818101908884111561499b57600080fd5b948201945b838610156149c25785356149b381614004565b825294820194908201906149a0565b865250858101359350828411156149d857600080fd5b6149e487858801614849565b908501525060408401359150808211156149fd57600080fd5b50614a0a848285016148c9565b6040830152506060820135606082015292915050565b600060208284031215614a3257600080fd5b813567ffffffffffffffff811115614a4957600080fd5b61459484828501614924565b801515811461098457600080fd5b803561402581614a55565b60008060408385031215614a8157600080fd5b823567ffffffffffffffff811115614a9857600080fd5b614aa485828601614924565b925050602083013561423b81614a55565b60c08101611923828463ffffffff80825116835267ffffffffffffffff602083015116602084015260408201516001600160a01b03808216604086015282606085015116606086015280608085015116608086015250505061ffff60a08201511660a08301525050565b60008060408385031215614b3257600080fd5b823567ffffffffffffffff80821115614b4a57600080fd5b908401906101808287031215614b5f57600080fd5b614b676140e7565b614b708361401a565b8152614b7e6020840161401a565b602082015260408301356040820152614b9960608401614150565b6060820152614baa6080840161401a565b608082015260a083013560a0820152614bc560c08401614a63565b60c0820152614bd660e08401614150565b60e08201526101008084013583811115614bef57600080fd5b614bfb89828701614386565b8284015250506101208084013583811115614c1557600080fd5b614c21898287016146a6565b8284015250506101409150614c37828401614150565b8282015261016091508183013582820152809450505050614c5a60208401614a63565b90509250929050565b60008083601f840112614c7557600080fd5b50813567ffffffffffffffff811115614c8d57600080fd5b6020830191508360208260051b8501011115614ca857600080fd5b9250929050565b60008060008060008060008060e0898b031215614ccb57600080fd5b606089018a811115614cdc57600080fd5b8998503567ffffffffffffffff80821115614cf657600080fd5b818b0191508b601f830112614d0a57600080fd5b813581811115614d1957600080fd5b8c6020828501011115614d2b57600080fd5b6020830199508098505060808b0135915080821115614d4957600080fd5b614d558c838d01614c63565b909750955060a08b0135915080821115614d6e57600080fd5b50614d7b8b828c01614c63565b999c989b50969995989497949560c00135949350505050565b600060608284031215614da657600080fd5b6040516060810181811067ffffffffffffffff82111715614dc957614dc9614085565b6040528235614dd781613fd2565b8152602083013579ffffffffffffffffffffffffffffffffffffffffffffffffffff81168114614e0657600080fd5b60208201526040928301359281019290925250919050565b600060208284031215614e3057600080fd5b813567ffffffffffffffff811115614e4757600080fd5b61459484828501614386565b805161402581614a55565b600060208284031215614e7057600080fd5b815161192081614a55565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615614eab57614eab614e7b565b500290565b634e487b7160e01b600052603260045260246000fd5b60006000198203614ed957614ed9614e7b565b5060010190565b600063ffffffff808316818103614ef957614ef9614e7b565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614f338184018a6145e0565b90508281036080840152614f4781896145e0565b905060ff871660a084015282810360c0840152614f648187614276565b905067ffffffffffffffff851660e0840152828103610100840152614f898185614276565b9c9b505050505050505050505050565b600082821015614fab57614fab614e7b565b500390565b634e487b7160e01b600052603160045260246000fd5b600060208284031215614fd857600080fd5b815161192081613fd2565b60008219821115614ff657614ff6614e7b565b500190565b608081528451608082015267ffffffffffffffff60208601511660a08201526000604086015160a060c0840152615036610120840182614276565b905060608701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80808584030160e08601526150728383614276565b9250608089015191508085840301610100860152506150918282614775565b925050506150a3602083018615159052565b8360408301526150be60608301846001600160a01b03169052565b95945050505050565b600060ff821660ff84168060ff038211156150e4576150e4614e7b565b019392505050565b634e487b7160e01b600052601260045260246000fd5b600060ff831680615115576151156150ec565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b61010081016151bb828663ffffffff80825116835267ffffffffffffffff602083015116602084015260408201516001600160a01b03808216604086015282606085015116606086015280608085015116608086015250505061ffff60a08201511660a08301525050565b67ffffffffffffffff841660c08301526001600160a01b03831660e0830152949350505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b16604085015281606085015261521c8285018b6145e0565b91508382036080850152615230828a6145e0565b915060ff881660a085015283820360c085015261524d8288614276565b90861660e08501528381036101008501529050614f898185614276565b805161402581614004565b805161402581613fd2565b600082601f83011261529157600080fd5b815161529f6142fa8261435e565b8181528460208386010111156152b457600080fd5b614594826020830160208701614246565b600082601f8301126152d657600080fd5b815160206152e66142fa836142b5565b82815260069290921b8401810191818101908684111561530557600080fd5b8286015b8481101561434257604081890312156153225760008081fd5b61532a61409b565b815161533581613fd2565b81528185015185820152835291830191604001615309565b60006020828403121561535f57600080fd5b815167ffffffffffffffff8082111561537757600080fd5b90830190610180828603121561538c57600080fd5b6153946140e7565b61539d8361526a565b81526153ab6020840161526a565b6020820152604083015160408201526153c660608401615275565b60608201526153d76080840161526a565b608082015260a083015160a08201526153f260c08401614e53565b60c082015261540360e08401615275565b60e0820152610100808401518381111561541c57600080fd5b61542888828701615280565b828401525050610120808401518381111561544257600080fd5b61544e888287016152c5565b8284015250506101409150615464828401615275565b9181019190915261016091820151918101919091529392505050565b600081518084526020808501945080840160005b8381101561461957815187529582019590820190600101615494565b6060815260006154c36060830186615480565b82810360208401526154d58186615480565b915050826040830152949350505050565b6000602082840312156154f857600080fd5b5051919050565b600067ffffffffffffffff80831681851680830382111561552257615522614e7b565b01949350505050565b600067ffffffffffffffff808316818103614ef957614ef9614e7b565b600082601f83011261555957600080fd5b815160206155696142fa836142b5565b82815260059290921b8401810191818101908684111561558857600080fd5b8286015b8481101561434257805167ffffffffffffffff8111156155ac5760008081fd5b6155ba8986838b0101615280565b84525091830191830161558c565b600082601f8301126155d957600080fd5b815160206155e96142fa836142b5565b82815260059290921b8401810191818101908684111561560857600080fd5b8286015b84811015614342578051835291830191830161560c565b6000602080838503121561563657600080fd5b825167ffffffffffffffff8082111561564e57600080fd5b908401906080828703121561566257600080fd5b61566a6140c4565b82518281111561567957600080fd5b8301601f8101881361568a57600080fd5b80516156986142fa826142b5565b81815260059190911b8201860190868101908a8311156156b757600080fd5b928701925b828410156156de5783516156cf81614004565b825292870192908701906156bc565b845250505082840151828111156156f457600080fd5b61570088828601615548565b8583015250604083015193508184111561571957600080fd5b615725878585016155c8565b6040820152606083015160608201528094505050505092915050565b600082615750576157506150ec565b500490565b6040815261577060408201845167ffffffffffffffff169052565b6000602084015161578d606084018267ffffffffffffffff169052565b506040840151608083015260608401516001600160a01b03811660a084015250608084015167ffffffffffffffff811660c08401525060a084015160e083015260c08401516101006157e28185018315159052565b60e08601519150610120615800818601846001600160a01b03169052565b818701519250610180915061014082818701526158216101c0870185614276565b93508188015191506101607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018188015261585f8584614775565b945081890151925061587b848801846001600160a01b03169052565b8801516101a087015250505083151560208401529050612319565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156158de5780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
}

var EVM2EVMOffRampHelperABI = EVM2EVMOffRampHelperMetaData.ABI

var EVM2EVMOffRampHelperBin = EVM2EVMOffRampHelperMetaData.Bin

func DeployEVM2EVMOffRampHelper(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId uint64, chainId uint64, onRampAddress common.Address, offRampConfig IEVM2EVMOffRampOffRampConfig, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig IAggregateRateLimiterRateLimiterConfig) (common.Address, *types.Transaction, *EVM2EVMOffRampHelper, error) {
	parsed, err := EVM2EVMOffRampHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMOffRampHelperBin), backend, sourceChainId, chainId, onRampAddress, offRampConfig, afn, sourceTokens, pools, rateLimiterConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMOffRampHelper{EVM2EVMOffRampHelperCaller: EVM2EVMOffRampHelperCaller{contract: contract}, EVM2EVMOffRampHelperTransactor: EVM2EVMOffRampHelperTransactor{contract: contract}, EVM2EVMOffRampHelperFilterer: EVM2EVMOffRampHelperFilterer{contract: contract}}, nil
}

type EVM2EVMOffRampHelper struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMOffRampHelperCaller
	EVM2EVMOffRampHelperTransactor
	EVM2EVMOffRampHelperFilterer
}

type EVM2EVMOffRampHelperCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMOffRampHelperTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMOffRampHelperFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMOffRampHelperSession struct {
	Contract     *EVM2EVMOffRampHelper
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMOffRampHelperCallerSession struct {
	Contract *EVM2EVMOffRampHelperCaller
	CallOpts bind.CallOpts
}

type EVM2EVMOffRampHelperTransactorSession struct {
	Contract     *EVM2EVMOffRampHelperTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMOffRampHelperRaw struct {
	Contract *EVM2EVMOffRampHelper
}

type EVM2EVMOffRampHelperCallerRaw struct {
	Contract *EVM2EVMOffRampHelperCaller
}

type EVM2EVMOffRampHelperTransactorRaw struct {
	Contract *EVM2EVMOffRampHelperTransactor
}

func NewEVM2EVMOffRampHelper(address common.Address, backend bind.ContractBackend) (*EVM2EVMOffRampHelper, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMOffRampHelperABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMOffRampHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelper{address: address, abi: abi, EVM2EVMOffRampHelperCaller: EVM2EVMOffRampHelperCaller{contract: contract}, EVM2EVMOffRampHelperTransactor: EVM2EVMOffRampHelperTransactor{contract: contract}, EVM2EVMOffRampHelperFilterer: EVM2EVMOffRampHelperFilterer{contract: contract}}, nil
}

func NewEVM2EVMOffRampHelperCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMOffRampHelperCaller, error) {
	contract, err := bindEVM2EVMOffRampHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperCaller{contract: contract}, nil
}

func NewEVM2EVMOffRampHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMOffRampHelperTransactor, error) {
	contract, err := bindEVM2EVMOffRampHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperTransactor{contract: contract}, nil
}

func NewEVM2EVMOffRampHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMOffRampHelperFilterer, error) {
	contract, err := bindEVM2EVMOffRampHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperFilterer{contract: contract}, nil
}

func bindEVM2EVMOffRampHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMOffRampHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMOffRampHelper.Contract.EVM2EVMOffRampHelperCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.EVM2EVMOffRampHelperTransactor.contract.Transfer(opts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.EVM2EVMOffRampHelperTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMOffRampHelper.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.contract.Transfer(opts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(IAggregateRateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(IAggregateRateLimiterTokenBucket)).(*IAggregateRateLimiterTokenBucket)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMOffRampHelper.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMOffRampHelper.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _EVM2EVMOffRampHelper.Contract.CcipReceive(&_EVM2EVMOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _EVM2EVMOffRampHelper.Contract.CcipReceive(&_EVM2EVMOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetAFN() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetAFN(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetAFN(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

	error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getChainIDs")

	outstruct := new(GetChainIDs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ChainId = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMOffRampHelper.Contract.GetChainIDs(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMOffRampHelper.Contract.GetChainIDs(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetDestinationToken(&_EVM2EVMOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetDestinationToken(&_EVM2EVMOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetDestinationTokens(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetDestinationTokens(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMOffRampHelper.Contract.GetExecutionState(&_EVM2EVMOffRampHelper.CallOpts, sequenceNumber)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMOffRampHelper.Contract.GetExecutionState(&_EVM2EVMOffRampHelper.CallOpts, sequenceNumber)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetOffRampConfig(opts *bind.CallOpts) (IEVM2EVMOffRampOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getOffRampConfig")

	if err != nil {
		return *new(IEVM2EVMOffRampOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IEVM2EVMOffRampOffRampConfig)).(*IEVM2EVMOffRampOffRampConfig)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetOffRampConfig() (IEVM2EVMOffRampOffRampConfig, error) {
	return _EVM2EVMOffRampHelper.Contract.GetOffRampConfig(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetOffRampConfig() (IEVM2EVMOffRampOffRampConfig, error) {
	return _EVM2EVMOffRampHelper.Contract.GetOffRampConfig(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getPoolByDestToken", destToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetPoolByDestToken(destToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetPoolByDestToken(&_EVM2EVMOffRampHelper.CallOpts, destToken)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetPoolByDestToken(destToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetPoolByDestToken(&_EVM2EVMOffRampHelper.CallOpts, destToken)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getPoolBySourceToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetPoolBySourceToken(&_EVM2EVMOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetPoolBySourceToken(&_EVM2EVMOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMOffRampHelper.Contract.GetPricesForTokens(&_EVM2EVMOffRampHelper.CallOpts, tokens)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMOffRampHelper.Contract.GetPricesForTokens(&_EVM2EVMOffRampHelper.CallOpts, tokens)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getSenderNonce", sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMOffRampHelper.Contract.GetSenderNonce(&_EVM2EVMOffRampHelper.CallOpts, sender)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMOffRampHelper.Contract.GetSenderNonce(&_EVM2EVMOffRampHelper.CallOpts, sender)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetSupportedTokens(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetSupportedTokens(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetTokenLimitAdmin(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetTokenLimitAdmin(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetTransmitters() ([]common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetTransmitters(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetTransmitters() ([]common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetTransmitters(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMOffRampHelper.Contract.IsAFNHealthy(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMOffRampHelper.Contract.IsAFNHealthy(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMOffRampHelper.Contract.LatestConfigDetails(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMOffRampHelper.Contract.LatestConfigDetails(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMOffRampHelper.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMOffRampHelper.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) MetadataHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "metadataHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) MetadataHash() ([32]byte, error) {
	return _EVM2EVMOffRampHelper.Contract.MetadataHash(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) MetadataHash() ([32]byte, error) {
	return _EVM2EVMOffRampHelper.Contract.MetadataHash(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Owner() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.Owner(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.Owner(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Paused() (bool, error) {
	return _EVM2EVMOffRampHelper.Contract.Paused(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) Paused() (bool, error) {
	return _EVM2EVMOffRampHelper.Contract.Paused(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) TypeAndVersion() (string, error) {
	return _EVM2EVMOffRampHelper.Contract.TypeAndVersion(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMOffRampHelper.Contract.TypeAndVersion(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.AcceptOwnership(&_EVM2EVMOffRampHelper.TransactOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.AcceptOwnership(&_EVM2EVMOffRampHelper.TransactOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.AddPool(&_EVM2EVMOffRampHelper.TransactOpts, token, pool)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.AddPool(&_EVM2EVMOffRampHelper.TransactOpts, token, pool)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) Execute(opts *bind.TransactOpts, rep InternalExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "execute", rep, manualExecution)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Execute(rep InternalExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Execute(&_EVM2EVMOffRampHelper.TransactOpts, rep, manualExecution)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) Execute(rep InternalExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Execute(&_EVM2EVMOffRampHelper.TransactOpts, rep, manualExecution)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "executeSingleMessage", message, manualExecution)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMOffRampHelper.TransactOpts, message, manualExecution)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMOffRampHelper.TransactOpts, message, manualExecution)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ManuallyExecute(opts *bind.TransactOpts, report InternalExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "manuallyExecute", report)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ManuallyExecute(report InternalExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ManuallyExecute(&_EVM2EVMOffRampHelper.TransactOpts, report)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ManuallyExecute(report InternalExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ManuallyExecute(&_EVM2EVMOffRampHelper.TransactOpts, report)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "pause")
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Pause(&_EVM2EVMOffRampHelper.TransactOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Pause(&_EVM2EVMOffRampHelper.TransactOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ReleaseOrMintToken(opts *bind.TransactOpts, pool common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "releaseOrMintToken", pool, amount, receiver)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ReleaseOrMintToken(pool common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ReleaseOrMintToken(&_EVM2EVMOffRampHelper.TransactOpts, pool, amount, receiver)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ReleaseOrMintToken(pool common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ReleaseOrMintToken(&_EVM2EVMOffRampHelper.TransactOpts, pool, amount, receiver)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ReleaseOrMintTokens(opts *bind.TransactOpts, sourceTokenAmounts []ClientEVMTokenAmount, receiver common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "releaseOrMintTokens", sourceTokenAmounts, receiver)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ReleaseOrMintTokens(sourceTokenAmounts []ClientEVMTokenAmount, receiver common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ReleaseOrMintTokens(&_EVM2EVMOffRampHelper.TransactOpts, sourceTokenAmounts, receiver)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ReleaseOrMintTokens(sourceTokenAmounts []ClientEVMTokenAmount, receiver common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ReleaseOrMintTokens(&_EVM2EVMOffRampHelper.TransactOpts, sourceTokenAmounts, receiver)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.RemovePool(&_EVM2EVMOffRampHelper.TransactOpts, token, pool)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.RemovePool(&_EVM2EVMOffRampHelper.TransactOpts, token, pool)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "report", executableMessages)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Report(executableMessages []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Report(&_EVM2EVMOffRampHelper.TransactOpts, executableMessages)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) Report(executableMessages []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Report(&_EVM2EVMOffRampHelper.TransactOpts, executableMessages)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetAFN(&_EVM2EVMOffRampHelper.TransactOpts, afn)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetAFN(&_EVM2EVMOffRampHelper.TransactOpts, afn)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetExecutionState(opts *bind.TransactOpts, sequenceNumber uint64, state uint8) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setExecutionState", sequenceNumber, state)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetExecutionState(sequenceNumber uint64, state uint8) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetExecutionState(&_EVM2EVMOffRampHelper.TransactOpts, sequenceNumber, state)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetExecutionState(sequenceNumber uint64, state uint8) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetExecutionState(&_EVM2EVMOffRampHelper.TransactOpts, sequenceNumber, state)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setOCR2Config", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetOCR2Config(&_EVM2EVMOffRampHelper.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetOCR2Config(&_EVM2EVMOffRampHelper.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetOffRampConfig(opts *bind.TransactOpts, config IEVM2EVMOffRampOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setOffRampConfig", config)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetOffRampConfig(config IEVM2EVMOffRampOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetOffRampConfig(&_EVM2EVMOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetOffRampConfig(config IEVM2EVMOffRampOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetOffRampConfig(&_EVM2EVMOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetPrices(&_EVM2EVMOffRampHelper.TransactOpts, tokens, prices)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetPrices(&_EVM2EVMOffRampHelper.TransactOpts, tokens, prices)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetRateLimiterConfig(&_EVM2EVMOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetRateLimiterConfig(&_EVM2EVMOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetTokenLimitAdmin(&_EVM2EVMOffRampHelper.TransactOpts, newAdmin)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetTokenLimitAdmin(&_EVM2EVMOffRampHelper.TransactOpts, newAdmin)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.TransferOwnership(&_EVM2EVMOffRampHelper.TransactOpts, to)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.TransferOwnership(&_EVM2EVMOffRampHelper.TransactOpts, to)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Transmit(&_EVM2EVMOffRampHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Transmit(&_EVM2EVMOffRampHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "unpause")
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Unpause(&_EVM2EVMOffRampHelper.TransactOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Unpause(&_EVM2EVMOffRampHelper.TransactOpts)
}

type EVM2EVMOffRampHelperAFNSetIterator struct {
	Event *EVM2EVMOffRampHelperAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperAFNSet)
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
		it.Event = new(EVM2EVMOffRampHelperAFNSet)
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

func (it *EVM2EVMOffRampHelperAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperAFNSetIterator{contract: _EVM2EVMOffRampHelper.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperAFNSet)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseAFNSet(log types.Log) (*EVM2EVMOffRampHelperAFNSet, error) {
	event := new(EVM2EVMOffRampHelperAFNSet)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperConfigChangedIterator struct {
	Event *EVM2EVMOffRampHelperConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperConfigChanged)
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
		it.Event = new(EVM2EVMOffRampHelperConfigChanged)
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

func (it *EVM2EVMOffRampHelperConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperConfigChangedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperConfigChanged)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMOffRampHelperConfigChanged, error) {
	event := new(EVM2EVMOffRampHelperConfigChanged)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperConfigSetIterator struct {
	Event *EVM2EVMOffRampHelperConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperConfigSet)
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
		it.Event = new(EVM2EVMOffRampHelperConfigSet)
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

func (it *EVM2EVMOffRampHelperConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperConfigSet struct {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperConfigSetIterator{contract: _EVM2EVMOffRampHelper.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperConfigSet)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseConfigSet(log types.Log) (*EVM2EVMOffRampHelperConfigSet, error) {
	event := new(EVM2EVMOffRampHelperConfigSet)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperExecutionStateChangedIterator struct {
	Event *EVM2EVMOffRampHelperExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperExecutionStateChanged)
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
		it.Event = new(EVM2EVMOffRampHelperExecutionStateChanged)
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

func (it *EVM2EVMOffRampHelperExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperExecutionStateChanged struct {
	SequenceNumber uint64
	MessageId      [32]byte
	State          uint8
	Raw            types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64, messageId [][32]byte) (*EVM2EVMOffRampHelperExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperExecutionStateChangedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperExecutionStateChanged, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperExecutionStateChanged)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseExecutionStateChanged(log types.Log) (*EVM2EVMOffRampHelperExecutionStateChanged, error) {
	event := new(EVM2EVMOffRampHelperExecutionStateChanged)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperOffRampConfigChangedIterator struct {
	Event *EVM2EVMOffRampHelperOffRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperOffRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperOffRampConfigChanged)
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
		it.Event = new(EVM2EVMOffRampHelperOffRampConfigChanged)
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

func (it *EVM2EVMOffRampHelperOffRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperOffRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperOffRampConfigChanged struct {
	Config        IEVM2EVMOffRampOffRampConfig
	SourceChainId uint64
	OnRamp        common.Address
	Raw           types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterOffRampConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperOffRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "OffRampConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperOffRampConfigChangedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "OffRampConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchOffRampConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOffRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "OffRampConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperOffRampConfigChanged)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "OffRampConfigChanged", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseOffRampConfigChanged(log types.Log) (*EVM2EVMOffRampHelperOffRampConfigChanged, error) {
	event := new(EVM2EVMOffRampHelperOffRampConfigChanged)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "OffRampConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMOffRampHelperOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMOffRampHelperOwnershipTransferRequested)
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

func (it *EVM2EVMOffRampHelperOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampHelperOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperOwnershipTransferRequestedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperOwnershipTransferRequested)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMOffRampHelperOwnershipTransferRequested, error) {
	event := new(EVM2EVMOffRampHelperOwnershipTransferRequested)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperOwnershipTransferredIterator struct {
	Event *EVM2EVMOffRampHelperOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperOwnershipTransferred)
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
		it.Event = new(EVM2EVMOffRampHelperOwnershipTransferred)
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

func (it *EVM2EVMOffRampHelperOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampHelperOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperOwnershipTransferredIterator{contract: _EVM2EVMOffRampHelper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperOwnershipTransferred)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMOffRampHelperOwnershipTransferred, error) {
	event := new(EVM2EVMOffRampHelperOwnershipTransferred)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperPausedIterator struct {
	Event *EVM2EVMOffRampHelperPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperPaused)
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
		it.Event = new(EVM2EVMOffRampHelperPaused)
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

func (it *EVM2EVMOffRampHelperPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperPausedIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperPausedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperPaused)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParsePaused(log types.Log) (*EVM2EVMOffRampHelperPaused, error) {
	event := new(EVM2EVMOffRampHelperPaused)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperPoolAddedIterator struct {
	Event *EVM2EVMOffRampHelperPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperPoolAdded)
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
		it.Event = new(EVM2EVMOffRampHelperPoolAdded)
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

func (it *EVM2EVMOffRampHelperPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperPoolAddedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperPoolAdded)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMOffRampHelperPoolAdded, error) {
	event := new(EVM2EVMOffRampHelperPoolAdded)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperPoolRemovedIterator struct {
	Event *EVM2EVMOffRampHelperPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperPoolRemoved)
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
		it.Event = new(EVM2EVMOffRampHelperPoolRemoved)
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

func (it *EVM2EVMOffRampHelperPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperPoolRemovedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperPoolRemoved)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMOffRampHelperPoolRemoved, error) {
	event := new(EVM2EVMOffRampHelperPoolRemoved)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperSkippedIncorrectNonceIterator struct {
	Event *EVM2EVMOffRampHelperSkippedIncorrectNonce

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperSkippedIncorrectNonceIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperSkippedIncorrectNonce)
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
		it.Event = new(EVM2EVMOffRampHelperSkippedIncorrectNonce)
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

func (it *EVM2EVMOffRampHelperSkippedIncorrectNonceIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperSkippedIncorrectNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperSkippedIncorrectNonce struct {
	Nonce  uint64
	Sender common.Address
	Raw    types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterSkippedIncorrectNonce(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampHelperSkippedIncorrectNonceIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "SkippedIncorrectNonce", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperSkippedIncorrectNonceIterator{contract: _EVM2EVMOffRampHelper.contract, event: "SkippedIncorrectNonce", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchSkippedIncorrectNonce(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperSkippedIncorrectNonce, nonce []uint64, sender []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "SkippedIncorrectNonce", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperSkippedIncorrectNonce)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "SkippedIncorrectNonce", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseSkippedIncorrectNonce(log types.Log) (*EVM2EVMOffRampHelperSkippedIncorrectNonce, error) {
	event := new(EVM2EVMOffRampHelperSkippedIncorrectNonce)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "SkippedIncorrectNonce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperTokenPriceChangedIterator struct {
	Event *EVM2EVMOffRampHelperTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperTokenPriceChanged)
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
		it.Event = new(EVM2EVMOffRampHelperTokenPriceChanged)
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

func (it *EVM2EVMOffRampHelperTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperTokenPriceChangedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperTokenPriceChanged)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMOffRampHelperTokenPriceChanged, error) {
	event := new(EVM2EVMOffRampHelperTokenPriceChanged)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMOffRampHelperTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMOffRampHelperTokensRemovedFromBucket)
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

func (it *EVM2EVMOffRampHelperTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperTokensRemovedFromBucketIterator{contract: _EVM2EVMOffRampHelper.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperTokensRemovedFromBucket)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMOffRampHelperTokensRemovedFromBucket, error) {
	event := new(EVM2EVMOffRampHelperTokensRemovedFromBucket)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperTransmittedIterator struct {
	Event *EVM2EVMOffRampHelperTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperTransmitted)
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
		it.Event = new(EVM2EVMOffRampHelperTransmitted)
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

func (it *EVM2EVMOffRampHelperTransmittedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperTransmittedIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperTransmittedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperTransmitted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperTransmitted)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseTransmitted(log types.Log) (*EVM2EVMOffRampHelperTransmitted, error) {
	event := new(EVM2EVMOffRampHelperTransmitted)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampHelperUnpausedIterator struct {
	Event *EVM2EVMOffRampHelperUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperUnpaused)
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
		it.Event = new(EVM2EVMOffRampHelperUnpaused)
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

func (it *EVM2EVMOffRampHelperUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperUnpausedIterator{contract: _EVM2EVMOffRampHelper.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperUnpaused)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseUnpaused(log types.Log) (*EVM2EVMOffRampHelperUnpaused, error) {
	event := new(EVM2EVMOffRampHelperUnpaused)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelper) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMOffRampHelper.abi.Events["AFNSet"].ID:
		return _EVM2EVMOffRampHelper.ParseAFNSet(log)
	case _EVM2EVMOffRampHelper.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMOffRampHelper.ParseConfigChanged(log)
	case _EVM2EVMOffRampHelper.abi.Events["ConfigSet"].ID:
		return _EVM2EVMOffRampHelper.ParseConfigSet(log)
	case _EVM2EVMOffRampHelper.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMOffRampHelper.ParseExecutionStateChanged(log)
	case _EVM2EVMOffRampHelper.abi.Events["OffRampConfigChanged"].ID:
		return _EVM2EVMOffRampHelper.ParseOffRampConfigChanged(log)
	case _EVM2EVMOffRampHelper.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMOffRampHelper.ParseOwnershipTransferRequested(log)
	case _EVM2EVMOffRampHelper.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMOffRampHelper.ParseOwnershipTransferred(log)
	case _EVM2EVMOffRampHelper.abi.Events["Paused"].ID:
		return _EVM2EVMOffRampHelper.ParsePaused(log)
	case _EVM2EVMOffRampHelper.abi.Events["PoolAdded"].ID:
		return _EVM2EVMOffRampHelper.ParsePoolAdded(log)
	case _EVM2EVMOffRampHelper.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMOffRampHelper.ParsePoolRemoved(log)
	case _EVM2EVMOffRampHelper.abi.Events["SkippedIncorrectNonce"].ID:
		return _EVM2EVMOffRampHelper.ParseSkippedIncorrectNonce(log)
	case _EVM2EVMOffRampHelper.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMOffRampHelper.ParseTokenPriceChanged(log)
	case _EVM2EVMOffRampHelper.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMOffRampHelper.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMOffRampHelper.abi.Events["Transmitted"].ID:
		return _EVM2EVMOffRampHelper.ParseTransmitted(log)
	case _EVM2EVMOffRampHelper.abi.Events["Unpaused"].ID:
		return _EVM2EVMOffRampHelper.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMOffRampHelperAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMOffRampHelperConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMOffRampHelperConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMOffRampHelperExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x5e04b4755a4460aa6de4f3a906c4324a025c7449c02b52f5466659b5bfdfba5f")
}

func (EVM2EVMOffRampHelperOffRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0xe48449493f29cd3dc290c645fdc0077d7533294b25426b653f53500b32811438")
}

func (EVM2EVMOffRampHelperOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMOffRampHelperOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMOffRampHelperPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMOffRampHelperPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMOffRampHelperPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMOffRampHelperSkippedIncorrectNonce) Topic() common.Hash {
	return common.HexToHash("0xd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf41237")
}

func (EVM2EVMOffRampHelperTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMOffRampHelperTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMOffRampHelperTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (EVM2EVMOffRampHelperUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelper) Address() common.Address {
	return _EVM2EVMOffRampHelper.address
}

type EVM2EVMOffRampHelperInterface interface {
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error)

	CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

		error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetOffRampConfig(opts *bind.CallOpts) (IEVM2EVMOffRampOffRampConfig, error)

	GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error)

	GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)

	GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error)

	GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	MetadataHash(opts *bind.CallOpts) ([32]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, rep InternalExecutionReport, manualExecution bool) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, manualExecution bool) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report InternalExecutionReport) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	ReleaseOrMintToken(opts *bind.TransactOpts, pool common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error)

	ReleaseOrMintTokens(opts *bind.TransactOpts, sourceTokenAmounts []ClientEVMTokenAmount, receiver common.Address) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetExecutionState(opts *bind.TransactOpts, sequenceNumber uint64, state uint8) (*types.Transaction, error)

	SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetOffRampConfig(opts *bind.TransactOpts, config IEVM2EVMOffRampOffRampConfig) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error)

	SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMOffRampHelperAFNSet, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMOffRampHelperConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMOffRampHelperConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64, messageId [][32]byte) (*EVM2EVMOffRampHelperExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperExecutionStateChanged, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMOffRampHelperExecutionStateChanged, error)

	FilterOffRampConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperOffRampConfigChangedIterator, error)

	WatchOffRampConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOffRampConfigChanged) (event.Subscription, error)

	ParseOffRampConfigChanged(log types.Log) (*EVM2EVMOffRampHelperOffRampConfigChanged, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampHelperOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMOffRampHelperOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampHelperOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMOffRampHelperOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMOffRampHelperPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMOffRampHelperPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMOffRampHelperPoolRemoved, error)

	FilterSkippedIncorrectNonce(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampHelperSkippedIncorrectNonceIterator, error)

	WatchSkippedIncorrectNonce(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperSkippedIncorrectNonce, nonce []uint64, sender []common.Address) (event.Subscription, error)

	ParseSkippedIncorrectNonce(log types.Log) (*EVM2EVMOffRampHelperSkippedIncorrectNonce, error)

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMOffRampHelperTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMOffRampHelperTokensRemovedFromBucket, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMOffRampHelperTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMOffRampHelperUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
