// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_offramp

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
	FeeManager                              common.Address
	PermissionLessExecutionThresholdSeconds uint32
	ExecutionDelaySeconds                   uint64
	Router                                  common.Address
	MaxDataSize                             uint64
	CommitStore                             common.Address
	MaxTokensLength                         uint64
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
	FeeUpdates      []InternalFeeUpdate
	EncodedMessages [][]byte
	Proofs          [][32]byte
	ProofFlagBits   *big.Int
}

type InternalFeeUpdate struct {
	SourceFeeToken              common.Address
	DestChainId                 uint64
	FeeTokenBaseUnitsPerUnitGas *big.Int
}

var EVM2EVMOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"InvalidOffRampConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedGasPriceUpdate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"name\":\"OffRampConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRampConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolByDestToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceFeeToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"feeTokenBaseUnitsPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structInternal.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOffRampConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620063f9380380620063f983398101604081905262000035916200084a565b6000805460ff19168155819084908490879033908190816200009e5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000d857620000d881620003ce565b5050506001600160a01b0381166200010357604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03929092169190911790558051825114620001455760405162d8548360e71b815260040160405180910390fd5b81516200015a9060059060208501906200062a565b5060005b8251811015620002e357600060405180604001604052808484815181106200018a576200018a620009cd565b60200260200101516001600160a01b03168152602001836001600160601b031681525090508060036000868581518110620001c957620001c9620009cd565b6020908102919091018101516001600160a01b0390811683528282019390935260409091016000908120845194909201516001600160601b0316600160a01b029390921692909217909155815184519091600491869086908110620002325762000232620009cd565b60200260200101516001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000278573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200029e9190620009e3565b6001600160a01b039081168252602082019290925260400160002080546001600160a01b0319169290911691909117905550620002db8162000a0a565b90506200015e565b50508151600680546001600160a01b0319166001600160a01b0392831617905560408051608081018252602080860151808352958301805191830182905251928201839052426060909201829052600995909555600a94909455600b55600c9290925550861662000367576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160401b03808916608052871660a0526001600160a01b03861660c052620003b27fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a6200047f565b60e052620003c085620004e5565b505050505050505062000b4c565b336001600160a01b03821603620004285760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000095565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008160805160a05160c051604051602001620004c894939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b60608101516001600160a01b031615806200050b575060a08101516001600160a01b0316155b806200051f575080516001600160a01b0316155b15620005425780604051630e48a3c160e01b815260040162000095919062000a32565b80516013805460208401516040808601516001600160a01b039586166001600160c01b031990941693909317600160a01b63ffffffff9093168302176001600160c01b0316600160c01b6001600160401b0394851602179093556060850151601480546080808901519388166001600160e01b03199283161793861685029390931790915560a08701516015805460c0808b015193909916931692909217941690920292909217905551915190517fb4284598c699c1ce7d9b6db1b8f73ac462f762dc04449dc34344b293a142f617926200061f92859262000aad565b60405180910390a150565b82805482825590600052602060002090810192821562000682579160200282015b828111156200068257825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200064b565b506200069092915062000694565b5090565b5b8082111562000690576000815560010162000695565b80516001600160401b0381168114620006c357600080fd5b919050565b6001600160a01b0381168114620006de57600080fd5b50565b8051620006c381620006c8565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b0381118282101715620007295762000729620006ee565b60405290565b600082601f8301126200074157600080fd5b815160206001600160401b0380831115620007605762000760620006ee565b8260051b604051601f19603f83011681018181108482111715620007885762000788620006ee565b604052938452858101830193838101925087851115620007a757600080fd5b83870191505b84821015620007d3578151620007c381620006c8565b83529183019190830190620007ad565b979650505050505050565b600060608284031215620007f157600080fd5b604051606081016001600160401b0381118282101715620008165762000816620006ee565b806040525080915082516200082b81620006c8565b8082525060208301516020820152604083015160408201525092915050565b600080600080600080600080888a036102008112156200086957600080fd5b620008748a620006ab565b98506200088460208b01620006ab565b975060408a01516200089681620006c8565b965060e0605f1982011215620008ab57600080fd5b50620008b662000704565b60608a0151620008c681620006c8565b815260808a015163ffffffff81168114620008e057600080fd5b6020820152620008f360a08b01620006ab565b60408201526200090660c08b01620006e1565b60608201526200091960e08b01620006ab565b60808201526200092d6101008b01620006e1565b60a0820152620009416101208b01620006ab565b60c08201529450620009576101408a01620006e1565b6101608a01519094506001600160401b03808211156200097657600080fd5b620009848c838d016200072f565b94506101808b01519150808211156200099c57600080fd5b50620009ab8b828c016200072f565b925050620009be8a6101a08b01620007de565b90509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b600060208284031215620009f657600080fd5b815162000a0381620006c8565b9392505050565b60006001820162000a2b57634e487b7160e01b600052601160045260246000fd5b5060010190565b60e0810162000aa7828480516001600160a01b03908116835260208083015163ffffffff16908401526040808301516001600160401b039081169185019190915260608084015183169085015260808084015182169085015260a0838101519092169184019190915260c09182015116910152565b92915050565b610120810162000b23828680516001600160a01b03908116835260208083015163ffffffff16908401526040808301516001600160401b039081169185019190915260608084015183169085015260808084015182169085015260a0838101519092169184019190915260c09182015116910152565b6001600160401b039390931660e08201526001600160a01b039190911661010090910152919050565b60805160a05160c05160e05161586562000b9460003960006133b201526000612be60152600061029501526000818161027001528181612bc40152613aa101526158656000f3fe608060405234801561001057600080fd5b506004361061025c5760003560e01c806379ba509711610145578063addb9e19116100bd578063cb3c8c011161008c578063d7e2bb5011610071578063d7e2bb50146106d8578063eb511dd414610704578063f2fde38b1461071757600080fd5b8063cb3c8c01146106bd578063d3c7c2c7146106d057600080fd5b8063addb9e1914610664578063afcb95d714610677578063b1dc65a414610697578063b4069b31146106aa57600080fd5b8063856c82471161011457806390c2339b116100f957806390c2339b14610543578063a8ebd0f41461057e578063abc39f1f1461065157600080fd5b8063856c8247146104e15780638da5cb5b1461052d57600080fd5b806379ba50971461049357806381ff70481461049b5780638456cb59146104cb57806385572ffb146104d357600080fd5b80633f4ba83a116101d85780635c975abb116101a7578063666cab8d1161018c578063666cab8d14610456578063681fba161461046b578063744b92e21461048057600080fd5b80635c975abb1461041f5780635d86f1411461042a57600080fd5b80633f4ba83a146103d35780634352fa9f146103db5780634741062e146103ee578063599f64311461040e57600080fd5b806317332f9b1161022f5780631ef38174116102145780631ef38174146103885780632222dd421461039b57806339aa9264146103c057600080fd5b806317332f9b1461032c578063181f5a771461033f57600080fd5b8063087ae6df14610261578063108ee5fc146102c5578063142a98fc146102da578063147809b314610314575b600080fd5b6040805167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811682527f0000000000000000000000000000000000000000000000000000000000000000166020820152015b60405180910390f35b6102d86102d3366004613dea565b61072a565b005b6103076102e8366004613e2d565b67ffffffffffffffff1660009081526017602052604090205460ff1690565b6040516102bc9190613e60565b61031c6107e1565b60405190151581526020016102bc565b6102d861033a366004613f90565b61086e565b61037b6040518060400160405280601481526020017f45564d3245564d4f666652616d7020312e302e3000000000000000000000000081525081565b6040516102bc919061404e565b6102d8610396366004614183565b6109a9565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016102bc565b6102d86103ce366004613dea565b61100f565b6102d8611046565b6102d86103e93660046142b4565b611058565b6104016103fc36600461436f565b6112ad565b6040516102bc91906143ac565b6006546001600160a01b03166103a8565b60005460ff1661031c565b6103a8610438366004613dea565b6001600160a01b039081166000908152600360205260409020541690565b61045e611375565b6040516102bc9190614434565b6104736113d7565b6040516102bc9190614447565b6102d861048e366004614488565b61149c565b6102d86117f2565b600f54600d546040805163ffffffff808516825264010000000090940490931660208401528201526060016102bc565b6102d86118d5565b6102d861025c3660046144c1565b6105146104ef366004613dea565b6001600160a01b031660009081526016602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff90911681526020016102bc565b60005461010090046001600160a01b03166103a8565b61054b6118e5565b6040516102bc91908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6106446040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152506040805160e0810182526013546001600160a01b038082168352600160a01b80830463ffffffff166020850152780100000000000000000000000000000000000000000000000090920467ffffffffffffffff9081169484019490945260145480821660608501528290048416608084015260155490811660a08401520490911660c082015290565b6040516102bc91906144fc565b6102d861065f366004614611565b611985565b6102d8610672366004614755565b611b3e565b6040805160018152600060208201819052918101919091526060016102bc565b6102d86106a5366004614840565b611b52565b6103a86106b8366004613dea565b6120a2565b6102d86106cb366004614b31565b612190565b61047361219b565b6103a86106e6366004613dea565b6001600160a01b039081166000908152600460205260409020541690565b6102d8610712366004614488565b6121fb565b6102d8610725366004613dea565b612455565b610732612466565b6001600160a01b038116610772576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa158015610844573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108689190614c22565b15905090565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b0316141580156108ab57506006546001600160a01b03163314155b156108e2576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602081015179ffffffffffffffffffffffffffffffffffffffffffffffffffff11610939576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61094360096124c5565b6040810151600a8190556020820151600955600b546109629190612570565b600b556040818101516020808401518351928352908201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b855185518560ff16601f831115610a21576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064015b60405180910390fd5b80600003610a8b576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610a18565b818314610b19576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610a18565b610b24816003614c55565b8311610b8c576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610a18565b610b94612466565b60115460005b81811015610c3c576010600060118381548110610bb957610bb9614c74565b60009182526020808320909101546001600160a01b031683528201929092526040018120805461ffff1916905560128054601092919084908110610bff57610bff614c74565b60009182526020808320909101546001600160a01b031683528201929092526040019020805461ffff19169055610c3581614c8a565b9050610b9a565b50895160005b81811015610ed15760008c8281518110610c5e57610c5e614c74565b6020026020010151905060006002811115610c7b57610c7b613e4a565b6001600160a01b038216600090815260106020526040902054610100900460ff166002811115610cad57610cad613e4a565b14610d14576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610a18565b6040805180820190915260ff8316815260208101600190526001600160a01b03821660009081526010602090815260409091208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610d7c57610d7c613e4a565b021790555090505060008c8381518110610d9857610d98614c74565b6020026020010151905060006002811115610db557610db5613e4a565b6001600160a01b038216600090815260106020526040902054610100900460ff166002811115610de757610de7613e4a565b14610e4e576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610a18565b6040805180820190915260ff8416815260208101600290526001600160a01b03821660009081526010602090815260409091208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610eb657610eb6613e4a565b0217905550905050505080610eca90614c8a565b9050610c42565b508a51610ee59060119060208e0190613d33565b508951610ef99060129060208d0190613d33565b50600e805460ff8381166101000261ffff19909216908c1617179055600f8054610f62914691309190600090610f349063ffffffff16614ca4565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168e8e8e8e8e8e61258a565b600d600001819055506000600f60049054906101000a900463ffffffff16905043600f60046101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581600d60000154600f60009054906101000a900463ffffffff168f8f8f8f8f8f604051610ff999989796959493929190614cc7565b60405180910390a1505050505050505050505050565b611017612466565b6006805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b61104e612466565b611056612617565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561109557506006546001600160a01b03163314155b156110cc576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114611108576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085460005b8181101561116257600760006008838154811061112d5761112d614c74565b60009182526020808320909101546001600160a01b0316835282019290925260400181205561115b81614c8a565b905061110e565b5060005b8281101561129257600085828151811061118257611182614c74565b6020026020010151905060006001600160a01b0316816001600160a01b0316036111d8576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8482815181106111ea576111ea614c74565b602002602001015160076000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f8186848151811061124f5761124f614c74565b60200260200101516040516112799291906001600160a01b03929092168252602082015260400190565b60405180910390a15061128b81614c8a565b9050611166565b5083516112a6906008906020870190613d33565b5050505050565b80516060908067ffffffffffffffff8111156112cb576112cb613e88565b6040519080825280602002602001820160405280156112f4578160200160208202803683370190505b50915060005b8181101561136e576007600085838151811061131857611318614c74565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000205483828151811061135357611353614c74565b602090810291909101015261136781614c8a565b90506112fa565b5050919050565b606060128054806020026020016040519081016040528092919081815260200182805480156113cd57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116113af575b5050505050905090565b60055460609067ffffffffffffffff8111156113f5576113f5613e88565b60405190808252806020026020018201604052801561141e578160200160208202803683370190505b50905060005b6005548110156114985761145e6005828154811061144457611444614c74565b6000918252602090912001546001600160a01b03166120a2565b82828151811061147057611470614c74565b6001600160a01b039092166020928302919091019091015261149181614c8a565b9050611424565b5090565b6114a4612466565b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352600160a01b9093046bffffffffffffffffffffffff169082015290611521576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816001600160a01b031681600001516001600160a01b031614611570576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005805490600090611583600184614d5d565b8154811061159357611593614c74565b9060005260206000200160009054906101000a90046001600160a01b03169050600583602001516bffffffffffffffffffffffff16815481106115d8576115d8614c74565b6000918252602090912001546001600160a01b031660056115fa600185614d5d565b8154811061160a5761160a614c74565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600584602001516bffffffffffffffffffffffff168154811061165e5761165e614c74565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558581015184841683526003909152604090912080546bffffffffffffffffffffffff909216600160a01b029190921617905560058054806116d7576116d7614d74565b6001900381819060005260206000200160006101000a8154906001600160a01b030219169055905560046000856001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611741573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117659190614d8a565b6001600160a01b03908116825260208083019390935260409182016000908120805473ffffffffffffffffffffffffffffffffffffffff1916905588821680825260038552838220919091558251908152908716928101929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b0316331461184c5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610a18565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6118dd612466565b6110566126b3565b6119106040518060800160405280600081526020016000815260200160008152602001600081525090565b604080516080810182526009548152600a546020820152600b5491810191909152600c54606082018190526000906119489042614d5d565b60208301518351919250611974916119609084614c55565b846040015161196f9190614da7565b612570565b604083015250426060820152919050565b3330146119be576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051600080825260208201909252816119fb565b60408051808201909152600080825260208201528152602001906001900390816119d45790505b506101208401515190915015611a2157611a1e8361012001518460e0015161273b565b90505b60e08301516001600160a01b03163b1580611a71575060e0830151611a6f906001600160a01b03167f85572ffb00000000000000000000000000000000000000000000000000000000612969565b155b15611a7b57505050565b6014546001600160a01b0316635607b375611a968584612985565b848660a001518760e001516040518563ffffffff1660e01b8152600401611ac09493929190614e03565b6020604051808303816000875af1158015611adf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b039190614c22565b611b39576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b611b46612466565b611b4f81612a28565b50565b611b9187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612c0b92505050565b60408051606081018252600d54808252600e5460ff808216602085015261010090910416928201929092528935918214611c045780516040517f93df584c000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610a18565b6040805183815260208c81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a1600281602001518260400151611c5f9190614ecf565b611c699190614f0a565b611c74906001614ecf565b60ff168614611caf576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b858414611ce8576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526010602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115611d2b57611d2b613e4a565b6002811115611d3c57611d3c613e4a565b9052509050600281602001516002811115611d5957611d59613e4a565b148015611d9357506012816000015160ff1681548110611d7b57611d7b614c74565b6000918252602090912001546001600160a01b031633145b611dc9576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506000611dd7856020614c55565b611de2886020614c55565b611dee8b610144614da7565b611df89190614da7565b611e029190614da7565b9050368114611e46576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401610a18565b5060008989604051611e59929190614f2c565b604051908190038120611e70918d90602001614f3c565b604051602081830303815290604052805190602001209050611e90613da1565b8760005b81811015612092576000600185898460208110611eb357611eb3614c74565b611ec091901a601b614ecf565b8e8e86818110611ed257611ed2614c74565b905060200201358d8d87818110611eeb57611eeb614c74565b9050602002013560405160008152602001604052604051611f28949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611f4a573d6000803e3d6000fd5b505060408051601f198101516001600160a01b038116600090815260106020908152848220848601909552845460ff8082168652939750919550929392840191610100909104166002811115611fa257611fa2613e4a565b6002811115611fb357611fb3613e4a565b9052509050600181602001516002811115611fd057611fd0613e4a565b14612007576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051859060ff16601f811061201e5761201e614c74565b60200201511561205a576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600185826000015160ff16601f811061207557612075614c74565b911515602090920201525061208b905081614c8a565b9050611e94565b5050505050505050505050505050565b6001600160a01b03808216600090815260036020526040812054909116806120f6576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015612165573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121899190614d8a565b9392505050565b611b4f816001612c25565b606060058054806020026020016040519081016040528092919081815260200182805480156113cd576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116113af575050505050905090565b612203612466565b6001600160a01b038216158061222057506001600160a01b038116155b15612257576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352600160a01b9093046bffffffffffffffffffffffff169082015290156122d5576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038083168083526005546bffffffffffffffffffffffff9081166020808601918252878516600090815260038252604080822088519451909516600160a01b02939096169290921790925583517f21df0da70000000000000000000000000000000000000000000000000000000081529351869460049492936321df0da79282870192819003870181865afa158015612379573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061239d9190614d8a565b6001600160a01b03908116825260208083019390935260409182016000908120805495831673ffffffffffffffffffffffffffffffffffffffff199687161790556005805460018101825591527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001805488831695168517905581519384528516918301919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b61245d612466565b611b4f81612e00565b60005461010090046001600160a01b031633146110565760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610a18565b8060010154816002015414806124de5750428160030154145b156124e65750565b806001015481600201541115612528576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081600301544261253a9190614d5d565b60018301548354919250612561916125529084614c55565b846002015461196f9190614da7565b60028301555042600390910155565b600081831061257f5781612581565b825b90505b92915050565b6000808a8a8a8a8a8a8a8a8a6040516020016125ae99989796959493929190614f58565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b60005460ff166126695760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610a18565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60005460ff16156127065760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610a18565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586126963390565b60606000835167ffffffffffffffff81111561275957612759613e88565b60405190808252806020026020018201604052801561279e57816020015b60408051808201909152600080825260208201528152602001906001900390816127775790505b50905060005b845181101561295f5760006127ee8683815181106127c4576127c4614c74565b6020026020010151600001516001600160a01b039081166000908152600360205260409020541690565b90506001600160a01b03811661285a5785828151811061281057612810614c74565b6020908102919091010151516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a18565b6128828187848151811061287057612870614c74565b60200260200101516020015187612ebc565b806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128e49190614d8a565b8383815181106128f6576128f6614c74565b60209081029190910101516001600160a01b039091169052855186908390811061292257612922614c74565b60200260200101516020015183838151811061294057612940614c74565b60209081029190910181015101525061295881614c8a565b90506127a4565b5061258181612f3c565b600061297483613140565b8015612581575061258183836131a4565b6040805160a08101825260008082526020820152606091810182905281810182905260808101919091526040518060a001604052808461016001518152602001846000015167ffffffffffffffff16815260200184606001516040516020016129fd91906001600160a01b0391909116815260200190565b6040516020818303038152906040528152602001846101000151815260200183815250905092915050565b60608101516001600160a01b03161580612a4d575060a08101516001600160a01b0316155b80612a60575080516001600160a01b0316155b15612a9957806040517f0e48a3c1000000000000000000000000000000000000000000000000000000008152600401610a1891906144fc565b80516013805460208401516040808601516001600160a01b039586167fffffffffffffffff00000000000000000000000000000000000000000000000090941693909317600160a01b63ffffffff90931683021777ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff948516021790935560608501516014805460808801519287167fffffffff0000000000000000000000000000000000000000000000000000000091821617928516840292909217905560a08601516015805460c0890151929097169690921695909517949092160292909217909155517fb4284598c699c1ce7d9b6db1b8f73ac462f762dc04449dc34344b293a142f6179061099e9083907f0000000000000000000000000000000000000000000000000000000000000000907f000000000000000000000000000000000000000000000000000000000000000090614fe0565b611b4f81806020019051810190612c2291906152b9565b60005b60005460ff1615612c785760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610a18565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ccb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cef9190614c22565b15612d25576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208201515115612de7578015612d68576040517f198753d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60135460208301516040517f9086658e0000000000000000000000000000000000000000000000000000000081526001600160a01b0390921691639086658e91612db49160040161539f565b600060405180830381600087803b158015612dce57600080fd5b505af1158015612de2573d6000803e3d6000fd5b505050505b60408201515115612dfc57612dfc8282613273565b5050565b336001600160a01b03821603612e585760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610a18565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b158015612f1f57600080fd5b505af1158015612f33573d6000803e3d6000fd5b50505050505050565b6000805b825181101561303b57600060076000858481518110612f6157612f61614c74565b6020026020010151600001516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600003612ff457838281518110612faa57612faa614c74565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a18565b83828151811061300657613006614c74565b6020026020010151602001518161301d9190614c55565b6130279084614da7565b9250508061303490614c8a565b9050612f40565b508015612dfc5761304c60096124c5565b600a5481111561309657600a546040517f688ccf77000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610a18565b600b548111156130f657600954600b54600091906130b49084614d5d565b6130be919061541d565b9050806040517fe31e0f32000000000000000000000000000000000000000000000000000000008152600401610a1891815260200190565b806009600201600082825461310b9190614d5d565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378906020016107d5565b600061316c827f01ffc9a7000000000000000000000000000000000000000000000000000000006131a4565b8015612584575061319d827fffffffff000000000000000000000000000000000000000000000000000000006131a4565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d9150600051905082801561325c575060208210155b80156132685750600081115b979650505050505050565b60408201515160008167ffffffffffffffff81111561329457613294613e88565b6040519080825280602002602001820160405280156132bd578160200160208202803683370190505b50905060008267ffffffffffffffff8111156132db576132db613e88565b60405190808252806020026020018201604052801561336957816020015b60408051610180810182526000808252602080830182905292820181905260608083018290526080830182905260a0830182905260c0830182905260e083018290526101008301819052610120830152610140820181905261016082015282526000199092019101816132f95790505b50905060005b838110156134245760008660400151828151811061338f5761338f614c74565b60200260200101518060200190518101906133aa91906154b9565b90506133d6817f0000000000000000000000000000000000000000000000000000000000000000613995565b8483815181106133e8576133e8614c74565b6020026020010181815250508083838151811061340757613407614c74565b6020026020010181905250508061341d90614c8a565b905061336f565b50601554606086015160808701516040517f320488750000000000000000000000000000000000000000000000000000000081526000936001600160a01b03169263320488759261347b928892919060040161561c565b6020604051808303816000875af115801561349a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134be9190615652565b9050600081116134fa576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b84811015612f3357600083828151811061351957613519614c74565b60200260200101519050600061354c826020015167ffffffffffffffff1660009081526017602052604090205460ff1690565b9050600081600381111561356257613562613e4a565b148061357f5750600381600381111561357d5761357d613e4a565b145b6135c75760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a18565b871561364657601354600090600160a01b900463ffffffff166135ea8642614d5d565b119050808061360a5750600382600381111561360857613608613e4a565b145b613640576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506136a3565b600081600381111561365a5761365a613e4a565b146136a35760208201516040517f67d9ba0f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a18565b60008160038111156136b7576136b7613e4a565b0361375657608082015160608301516001600160a01b031660009081526016602052604090205467ffffffffffffffff918216916136f79116600161566b565b67ffffffffffffffff16146137565781606001516001600160a01b0316826080015167ffffffffffffffff167fd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf4123760405160405180910390a35050613985565b61375f82613a9f565b60208281015167ffffffffffffffff166000908152601790915260408120805460ff19166001179055613792838a613bfb565b60208085015167ffffffffffffffff1660009081526017909152604090208054919250829160ff191660018360038111156137cf576137cf613e4a565b021790555088156138b5578260c0015180156137fc575060038260038111156137fa576137fa613e4a565b145b80156138195750600281600381111561381757613817613e4a565b145b806138515750600082600381111561383357613833613e4a565b1480156138515750600281600381111561384f5761384f613e4a565b145b156138b05760608301516001600160a01b03166000908152601660205260408120805467ffffffffffffffff169161388883615697565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b613935565b8260c0015180156138d7575060038160038111156138d5576138d5613e4a565b145b6139355760608301516001600160a01b03166000908152601660205260408120805467ffffffffffffffff169161390d83615697565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b826101600151836020015167ffffffffffffffff167f5e04b4755a4460aa6de4f3a906c4324a025c7449c02b52f5466659b5bfdfba5f836040516139799190613e60565b60405180910390a35050505b61398e81614c8a565b90506134fd565b60008060001b828460200151856080015186606001518760e00151886101000151805190602001208961012001516040516020016139d391906156b4565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d60400151604051602001613a819c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16816000015167ffffffffffffffff1614613b1f5780516040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a18565b60155461012082015151600160a01b90910467ffffffffffffffff161015613b855760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a18565b60145461010082015151600160a01b90910467ffffffffffffffff161015611b4f57601454610100820151516040517f86933789000000000000000000000000000000000000000000000000000000008152600160a01b90920467ffffffffffffffff1660048301526024820152604401610a18565b6040517fabc39f1f000000000000000000000000000000000000000000000000000000008152600090309063abc39f1f90613c3c90869086906004016156c7565b600060405180830381600087803b158015613c5657600080fd5b505af1925050508015613c67575060015b613d2a573d808015613c95576040519150601f19603f3d011682016040523d82523d6000602084013e613c9a565b606091505b50613ca481615808565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003613cf6576003915050612584565b806040517fcf19edfd000000000000000000000000000000000000000000000000000000008152600401610a18919061404e565b50600292915050565b828054828255906000526020600020908101928215613d95579160200282015b82811115613d95578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190613d53565b50611498929150613dc0565b604051806103e00160405280601f906020820280368337509192915050565b5b808211156114985760008155600101613dc1565b6001600160a01b0381168114611b4f57600080fd5b600060208284031215613dfc57600080fd5b813561258181613dd5565b67ffffffffffffffff81168114611b4f57600080fd5b8035613e2881613e07565b919050565b600060208284031215613e3f57600080fd5b813561258181613e07565b634e487b7160e01b600052602160045260246000fd5b6020810160048310613e8257634e487b7160e01b600052602160045260246000fd5b91905290565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613ec157613ec1613e88565b60405290565b604051610180810167ffffffffffffffff81118282101715613ec157613ec1613e88565b60405160e0810167ffffffffffffffff81118282101715613ec157613ec1613e88565b6040516060810167ffffffffffffffff81118282101715613ec157613ec1613e88565b60405160a0810167ffffffffffffffff81118282101715613ec157613ec1613e88565b604051601f8201601f1916810167ffffffffffffffff81118282101715613f7d57613f7d613e88565b604052919050565b8035613e2881613dd5565b600060608284031215613fa257600080fd5b6040516060810181811067ffffffffffffffff82111715613fc557613fc5613e88565b6040528235613fd381613dd5565b8152602083810135908201526040928301359281019290925250919050565b60005b8381101561400d578181015183820152602001613ff5565b8381111561401c576000848401525b50505050565b6000815180845261403a816020860160208601613ff2565b601f01601f19169290920160200192915050565b6020815260006125816020830184614022565b600067ffffffffffffffff82111561407b5761407b613e88565b5060051b60200190565b600082601f83011261409657600080fd5b813560206140ab6140a683614061565b613f54565b82815260059290921b840181019181810190868411156140ca57600080fd5b8286015b848110156140ee5780356140e181613dd5565b83529183019183016140ce565b509695505050505050565b803560ff81168114613e2857600080fd5b600067ffffffffffffffff82111561412457614124613e88565b50601f01601f191660200190565b600082601f83011261414357600080fd5b81356141516140a68261410a565b81815284602083860101111561416657600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c0878903121561419c57600080fd5b863567ffffffffffffffff808211156141b457600080fd5b6141c08a838b01614085565b975060208901359150808211156141d657600080fd5b6141e28a838b01614085565b96506141f060408a016140f9565b9550606089013591508082111561420657600080fd5b6142128a838b01614132565b945061422060808a01613e1d565b935060a089013591508082111561423657600080fd5b5061424389828a01614132565b9150509295509295509295565b600082601f83011261426157600080fd5b813560206142716140a683614061565b82815260059290921b8401810191818101908684111561429057600080fd5b8286015b848110156140ee5780356142a781613dd5565b8352918301918301614294565b600080604083850312156142c757600080fd5b823567ffffffffffffffff808211156142df57600080fd5b6142eb86838701614250565b935060209150818501358181111561430257600080fd5b85019050601f8101861361431557600080fd5b80356143236140a682614061565b81815260059190911b8201830190838101908883111561434257600080fd5b928401925b8284101561436057833582529284019290840190614347565b80955050505050509250929050565b60006020828403121561438157600080fd5b813567ffffffffffffffff81111561439857600080fd5b6143a484828501614250565b949350505050565b6020808252825182820181905260009190848201906040850190845b818110156143e4578351835292840192918401916001016143c8565b50909695505050505050565b600081518084526020808501945080840160005b838110156144295781516001600160a01b031687529582019590820190600101614404565b509495945050505050565b60208152600061258160208301846143f0565b6020808252825182820181905260009190848201906040850190845b818110156143e45783516001600160a01b031683529284019291840191600101614463565b6000806040838503121561449b57600080fd5b82356144a681613dd5565b915060208301356144b681613dd5565b809150509250929050565b6000602082840312156144d357600080fd5b813567ffffffffffffffff8111156144ea57600080fd5b820160a0818503121561258157600080fd5b60e0810161258482846001600160a01b0380825116835263ffffffff6020830151166020840152604082015167ffffffffffffffff80821660408601528260608501511660608601528060808501511660808601528260a08501511660a08601528060c08501511660c08601525050505050565b8015158114611b4f57600080fd5b8035613e2881614570565b600082601f83011261459a57600080fd5b813560206145aa6140a683614061565b82815260069290921b840181019181810190868411156145c957600080fd5b8286015b848110156140ee57604081890312156145e65760008081fd5b6145ee613e9e565b81356145f981613dd5565b815281850135858201528352918301916040016145cd565b6000806040838503121561462457600080fd5b823567ffffffffffffffff8082111561463c57600080fd5b90840190610180828703121561465157600080fd5b614659613ec7565b61466283613e1d565b815261467060208401613e1d565b60208201526040830135604082015261468b60608401613f85565b606082015261469c60808401613e1d565b608082015260a083013560a08201526146b760c0840161457e565b60c08201526146c860e08401613f85565b60e082015261010080840135838111156146e157600080fd5b6146ed89828701614132565b828401525050610120808401358381111561470757600080fd5b61471389828701614589565b8284015250506101409150614729828401613f85565b828201526101609150818301358282015280945050505061474c6020840161457e565b90509250929050565b600060e0828403121561476757600080fd5b61476f613eeb565b823561477a81613dd5565b8152602083013563ffffffff8116811461479357600080fd5b60208201526147a460408401613e1d565b60408201526147b560608401613f85565b60608201526147c660808401613e1d565b60808201526147d760a08401613f85565b60a08201526147e860c08401613e1d565b60c08201529392505050565b60008083601f84011261480657600080fd5b50813567ffffffffffffffff81111561481e57600080fd5b6020830191508360208260051b850101111561483957600080fd5b9250929050565b60008060008060008060008060e0898b03121561485c57600080fd5b606089018a81111561486d57600080fd5b8998503567ffffffffffffffff8082111561488757600080fd5b818b0191508b601f83011261489b57600080fd5b8135818111156148aa57600080fd5b8c60208285010111156148bc57600080fd5b6020830199508098505060808b01359150808211156148da57600080fd5b6148e68c838d016147f4565b909750955060a08b01359150808211156148ff57600080fd5b5061490c8b828c016147f4565b999c989b50969995989497949560c00135949350505050565b600082601f83011261493657600080fd5b813560206149466140a683614061565b82815260059290921b8401810191818101908684111561496557600080fd5b8286015b848110156140ee57803561497c81613e07565b8352918301918301614969565b6fffffffffffffffffffffffffffffffff81168114611b4f57600080fd5b600082601f8301126149b857600080fd5b813560206149c86140a683614061565b828152606092830285018201928282019190878511156149e757600080fd5b8387015b85811015614a495781818a031215614a035760008081fd5b614a0b613f0e565b8135614a1681613dd5565b815281860135614a2581613e07565b81870152604082810135614a3881614989565b9082015284529284019281016149eb565b5090979650505050505050565b600082601f830112614a6757600080fd5b81356020614a776140a683614061565b82815260059290921b84018101918181019086841115614a9657600080fd5b8286015b848110156140ee57803567ffffffffffffffff811115614aba5760008081fd5b614ac88986838b0101614132565b845250918301918301614a9a565b600082601f830112614ae757600080fd5b81356020614af76140a683614061565b82815260059290921b84018101918181019086841115614b1657600080fd5b8286015b848110156140ee5780358352918301918301614b1a565b600060208284031215614b4357600080fd5b813567ffffffffffffffff80821115614b5b57600080fd5b9083019060a08286031215614b6f57600080fd5b614b77613f31565b823582811115614b8657600080fd5b614b9287828601614925565b825250602083013582811115614ba757600080fd5b614bb3878286016149a7565b602083015250604083013582811115614bcb57600080fd5b614bd787828601614a56565b604083015250606083013582811115614bef57600080fd5b614bfb87828601614ad6565b6060830152506080830135608082015280935050505092915050565b8051613e2881614570565b600060208284031215614c3457600080fd5b815161258181614570565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615614c6f57614c6f614c3f565b500290565b634e487b7160e01b600052603260045260246000fd5b60006000198203614c9d57614c9d614c3f565b5060010190565b600063ffffffff808316818103614cbd57614cbd614c3f565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614cf78184018a6143f0565b90508281036080840152614d0b81896143f0565b905060ff871660a084015282810360c0840152614d288187614022565b905067ffffffffffffffff851660e0840152828103610100840152614d4d8185614022565b9c9b505050505050505050505050565b600082821015614d6f57614d6f614c3f565b500390565b634e487b7160e01b600052603160045260246000fd5b600060208284031215614d9c57600080fd5b815161258181613dd5565b60008219821115614dba57614dba614c3f565b500190565b600081518084526020808501945080840160005b8381101561442957815180516001600160a01b031688528301518388015260409096019590820190600101614dd3565b608081528451608082015267ffffffffffffffff60208601511660a08201526000604086015160a060c0840152614e3e610120840182614022565b905060608701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80808584030160e0860152614e7a8383614022565b925060808901519150808584030161010086015250614e998282614dbf565b92505050614eab602083018615159052565b836040830152614ec660608301846001600160a01b03169052565b95945050505050565b600060ff821660ff84168060ff03821115614eec57614eec614c3f565b019392505050565b634e487b7160e01b600052601260045260246000fd5b600060ff831680614f1d57614f1d614ef4565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b166040850152816060850152614f928285018b6143f0565b91508382036080850152614fa6828a6143f0565b915060ff881660a085015283820360c0850152614fc38288614022565b90861660e08501528381036101008501529050614d4d8185614022565b610120810161505582866001600160a01b0380825116835263ffffffff6020830151166020840152604082015167ffffffffffffffff80821660408601528260608501511660608601528060808501511660808601528260a08501511660a08601528060c08501511660c08601525050505050565b67ffffffffffffffff841660e08301526001600160a01b038316610100830152949350505050565b8051613e2881613e07565b600082601f83011261509957600080fd5b815160206150a96140a683614061565b82815260059290921b840181019181810190868411156150c857600080fd5b8286015b848110156140ee5780516150df81613e07565b83529183019183016150cc565b8051613e2881613dd5565b600082601f83011261510857600080fd5b815160206151186140a683614061565b8281526060928302850182019282820191908785111561513757600080fd5b8387015b85811015614a495781818a0312156151535760008081fd5b61515b613f0e565b815161516681613dd5565b81528186015161517581613e07565b8187015260408281015161518881614989565b90820152845292840192810161513b565b600082601f8301126151aa57600080fd5b81516151b86140a68261410a565b8181528460208386010111156151cd57600080fd5b6143a4826020830160208701613ff2565b600082601f8301126151ef57600080fd5b815160206151ff6140a683614061565b82815260059290921b8401810191818101908684111561521e57600080fd5b8286015b848110156140ee57805167ffffffffffffffff8111156152425760008081fd5b6152508986838b0101615199565b845250918301918301615222565b600082601f83011261526f57600080fd5b8151602061527f6140a683614061565b82815260059290921b8401810191818101908684111561529e57600080fd5b8286015b848110156140ee57805183529183019183016152a2565b6000602082840312156152cb57600080fd5b815167ffffffffffffffff808211156152e357600080fd5b9083019060a082860312156152f757600080fd5b6152ff613f31565b82518281111561530e57600080fd5b61531a87828601615088565b82525060208301518281111561532f57600080fd5b61533b878286016150f7565b60208301525060408301518281111561535357600080fd5b61535f878286016151de565b60408301525060608301518281111561537757600080fd5b6153838782860161525e565b6060830152506080830151608082015280935050505092915050565b602080825282518282018190526000919060409081850190868401855b8281101561541057815180516001600160a01b031685528681015167ffffffffffffffff16878601528501516fffffffffffffffffffffffffffffffff1685850152606090930192908501906001016153bc565b5091979650505050505050565b60008261542c5761542c614ef4565b500490565b600082601f83011261544257600080fd5b815160206154526140a683614061565b82815260069290921b8401810191818101908684111561547157600080fd5b8286015b848110156140ee576040818903121561548e5760008081fd5b615496613e9e565b81516154a181613dd5565b81528185015185820152835291830191604001615475565b6000602082840312156154cb57600080fd5b815167ffffffffffffffff808211156154e357600080fd5b9083019061018082860312156154f857600080fd5b615500613ec7565b6155098361507d565b81526155176020840161507d565b602082015260408301516040820152615532606084016150ec565b60608201526155436080840161507d565b608082015260a083015160a082015261555e60c08401614c17565b60c082015261556f60e084016150ec565b60e0820152610100808401518381111561558857600080fd5b61559488828701615199565b82840152505061012080840151838111156155ae57600080fd5b6155ba88828701615431565b82840152505061014091506155d08284016150ec565b9181019190915261016091820151918101919091529392505050565b600081518084526020808501945080840160005b8381101561442957815187529582019590820190600101615600565b60608152600061562f60608301866155ec565b828103602084015261564181866155ec565b915050826040830152949350505050565b60006020828403121561566457600080fd5b5051919050565b600067ffffffffffffffff80831681851680830382111561568e5761568e614c3f565b01949350505050565b600067ffffffffffffffff808316818103614cbd57614cbd614c3f565b6020815260006125816020830184614dbf565b604081526156e260408201845167ffffffffffffffff169052565b600060208401516156ff606084018267ffffffffffffffff169052565b506040840151608083015260608401516001600160a01b03811660a084015250608084015167ffffffffffffffff811660c08401525060a084015160e083015260c08401516101006157548185018315159052565b60e08601519150610120615772818601846001600160a01b03169052565b818701519250610180915061014082818701526157936101c0870185614022565b93508188015191506101607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301818801526157d18584614dbf565b94508189015192506157ed848801846001600160a01b03169052565b8801516101a087015250505083151560208401529050612189565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156158505780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
}

var EVM2EVMOffRampABI = EVM2EVMOffRampMetaData.ABI

var EVM2EVMOffRampBin = EVM2EVMOffRampMetaData.Bin

func DeployEVM2EVMOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId uint64, chainId uint64, onRampAddress common.Address, offRampConfig IEVM2EVMOffRampOffRampConfig, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig IAggregateRateLimiterRateLimiterConfig) (common.Address, *types.Transaction, *EVM2EVMOffRamp, error) {
	parsed, err := EVM2EVMOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMOffRampBin), backend, sourceChainId, chainId, onRampAddress, offRampConfig, afn, sourceTokens, pools, rateLimiterConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMOffRamp{EVM2EVMOffRampCaller: EVM2EVMOffRampCaller{contract: contract}, EVM2EVMOffRampTransactor: EVM2EVMOffRampTransactor{contract: contract}, EVM2EVMOffRampFilterer: EVM2EVMOffRampFilterer{contract: contract}}, nil
}

type EVM2EVMOffRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMOffRampCaller
	EVM2EVMOffRampTransactor
	EVM2EVMOffRampFilterer
}

type EVM2EVMOffRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMOffRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMOffRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMOffRampSession struct {
	Contract     *EVM2EVMOffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMOffRampCallerSession struct {
	Contract *EVM2EVMOffRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMOffRampTransactorSession struct {
	Contract     *EVM2EVMOffRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMOffRampRaw struct {
	Contract *EVM2EVMOffRamp
}

type EVM2EVMOffRampCallerRaw struct {
	Contract *EVM2EVMOffRampCaller
}

type EVM2EVMOffRampTransactorRaw struct {
	Contract *EVM2EVMOffRampTransactor
}

func NewEVM2EVMOffRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMOffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMOffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRamp{address: address, abi: abi, EVM2EVMOffRampCaller: EVM2EVMOffRampCaller{contract: contract}, EVM2EVMOffRampTransactor: EVM2EVMOffRampTransactor{contract: contract}, EVM2EVMOffRampFilterer: EVM2EVMOffRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMOffRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMOffRampCaller, error) {
	contract, err := bindEVM2EVMOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampCaller{contract: contract}, nil
}

func NewEVM2EVMOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMOffRampTransactor, error) {
	contract, err := bindEVM2EVMOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampTransactor{contract: contract}, nil
}

func NewEVM2EVMOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMOffRampFilterer, error) {
	contract, err := bindEVM2EVMOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampFilterer{contract: contract}, nil
}

func bindEVM2EVMOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMOffRampABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMOffRamp.Contract.EVM2EVMOffRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.EVM2EVMOffRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.EVM2EVMOffRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMOffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (IAggregateRateLimiterTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(IAggregateRateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(IAggregateRateLimiterTokenBucket)).(*IAggregateRateLimiterTokenBucket)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) CalculateCurrentTokenBucketState() (IAggregateRateLimiterTokenBucket, error) {
	return _EVM2EVMOffRamp.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _EVM2EVMOffRamp.Contract.CcipReceive(&_EVM2EVMOffRamp.CallOpts, arg0)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _EVM2EVMOffRamp.Contract.CcipReceive(&_EVM2EVMOffRamp.CallOpts, arg0)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetAFN() (common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetAFN(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetAFN(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

	error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getChainIDs")

	outstruct := new(GetChainIDs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ChainId = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMOffRamp.Contract.GetChainIDs(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMOffRamp.Contract.GetChainIDs(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetDestinationToken(&_EVM2EVMOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetDestinationToken(&_EVM2EVMOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetDestinationTokens(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetDestinationTokens(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMOffRamp.Contract.GetExecutionState(&_EVM2EVMOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMOffRamp.Contract.GetExecutionState(&_EVM2EVMOffRamp.CallOpts, sequenceNumber)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetOffRampConfig(opts *bind.CallOpts) (IEVM2EVMOffRampOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getOffRampConfig")

	if err != nil {
		return *new(IEVM2EVMOffRampOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IEVM2EVMOffRampOffRampConfig)).(*IEVM2EVMOffRampOffRampConfig)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetOffRampConfig() (IEVM2EVMOffRampOffRampConfig, error) {
	return _EVM2EVMOffRamp.Contract.GetOffRampConfig(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetOffRampConfig() (IEVM2EVMOffRampOffRampConfig, error) {
	return _EVM2EVMOffRamp.Contract.GetOffRampConfig(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getPoolByDestToken", destToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetPoolByDestToken(destToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetPoolByDestToken(&_EVM2EVMOffRamp.CallOpts, destToken)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetPoolByDestToken(destToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetPoolByDestToken(&_EVM2EVMOffRamp.CallOpts, destToken)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getPoolBySourceToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetPoolBySourceToken(&_EVM2EVMOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetPoolBySourceToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetPoolBySourceToken(&_EVM2EVMOffRamp.CallOpts, sourceToken)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMOffRamp.Contract.GetPricesForTokens(&_EVM2EVMOffRamp.CallOpts, tokens)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMOffRamp.Contract.GetPricesForTokens(&_EVM2EVMOffRamp.CallOpts, tokens)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getSenderNonce", sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMOffRamp.Contract.GetSenderNonce(&_EVM2EVMOffRamp.CallOpts, sender)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetSenderNonce(sender common.Address) (uint64, error) {
	return _EVM2EVMOffRamp.Contract.GetSenderNonce(&_EVM2EVMOffRamp.CallOpts, sender)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetSupportedTokens(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetSupportedTokens(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetTokenLimitAdmin(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetTransmitters() ([]common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetTransmitters(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetTransmitters() ([]common.Address, error) {
	return _EVM2EVMOffRamp.Contract.GetTransmitters(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMOffRamp.Contract.IsAFNHealthy(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMOffRamp.Contract.IsAFNHealthy(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMOffRamp.Contract.LatestConfigDetails(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMOffRamp.Contract.LatestConfigDetails(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) Owner() (common.Address, error) {
	return _EVM2EVMOffRamp.Contract.Owner(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMOffRamp.Contract.Owner(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) Paused() (bool, error) {
	return _EVM2EVMOffRamp.Contract.Paused(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMOffRamp.Contract.Paused(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMOffRamp.Contract.TypeAndVersion(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMOffRamp.Contract.TypeAndVersion(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.AcceptOwnership(&_EVM2EVMOffRamp.TransactOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.AcceptOwnership(&_EVM2EVMOffRamp.TransactOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.AddPool(&_EVM2EVMOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.AddPool(&_EVM2EVMOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "executeSingleMessage", message, manualExecution)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMOffRamp.TransactOpts, message, manualExecution)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMOffRamp.TransactOpts, message, manualExecution)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) ManuallyExecute(opts *bind.TransactOpts, report InternalExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "manuallyExecute", report)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) ManuallyExecute(report InternalExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.ManuallyExecute(&_EVM2EVMOffRamp.TransactOpts, report)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) ManuallyExecute(report InternalExecutionReport) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.ManuallyExecute(&_EVM2EVMOffRamp.TransactOpts, report)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.Pause(&_EVM2EVMOffRamp.TransactOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.Pause(&_EVM2EVMOffRamp.TransactOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.RemovePool(&_EVM2EVMOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.RemovePool(&_EVM2EVMOffRamp.TransactOpts, token, pool)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetAFN(&_EVM2EVMOffRamp.TransactOpts, afn)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetAFN(&_EVM2EVMOffRamp.TransactOpts, afn)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "setOCR2Config", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetOCR2Config(&_EVM2EVMOffRamp.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetOCR2Config(&_EVM2EVMOffRamp.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) SetOffRampConfig(opts *bind.TransactOpts, config IEVM2EVMOffRampOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "setOffRampConfig", config)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) SetOffRampConfig(config IEVM2EVMOffRampOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetOffRampConfig(&_EVM2EVMOffRamp.TransactOpts, config)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) SetOffRampConfig(config IEVM2EVMOffRampOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetOffRampConfig(&_EVM2EVMOffRamp.TransactOpts, config)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetPrices(&_EVM2EVMOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetPrices(&_EVM2EVMOffRamp.TransactOpts, tokens, prices)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMOffRamp.TransactOpts, config)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) SetRateLimiterConfig(config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMOffRamp.TransactOpts, config)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetTokenLimitAdmin(&_EVM2EVMOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.TransferOwnership(&_EVM2EVMOffRamp.TransactOpts, to)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.TransferOwnership(&_EVM2EVMOffRamp.TransactOpts, to)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.Transmit(&_EVM2EVMOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.Transmit(&_EVM2EVMOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.Unpause(&_EVM2EVMOffRamp.TransactOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.Unpause(&_EVM2EVMOffRamp.TransactOpts)
}

type EVM2EVMOffRampAFNSetIterator struct {
	Event *EVM2EVMOffRampAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampAFNSet)
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
		it.Event = new(EVM2EVMOffRampAFNSet)
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

func (it *EVM2EVMOffRampAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMOffRampAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampAFNSetIterator{contract: _EVM2EVMOffRamp.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampAFNSet)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseAFNSet(log types.Log) (*EVM2EVMOffRampAFNSet, error) {
	event := new(EVM2EVMOffRampAFNSet)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampConfigChangedIterator struct {
	Event *EVM2EVMOffRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampConfigChanged)
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
		it.Event = new(EVM2EVMOffRampConfigChanged)
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

func (it *EVM2EVMOffRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampConfigChangedIterator{contract: _EVM2EVMOffRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampConfigChanged)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMOffRampConfigChanged, error) {
	event := new(EVM2EVMOffRampConfigChanged)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampConfigSetIterator struct {
	Event *EVM2EVMOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampConfigSet)
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
		it.Event = new(EVM2EVMOffRampConfigSet)
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

func (it *EVM2EVMOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampConfigSet struct {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampConfigSetIterator{contract: _EVM2EVMOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampConfigSet)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseConfigSet(log types.Log) (*EVM2EVMOffRampConfigSet, error) {
	event := new(EVM2EVMOffRampConfigSet)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampExecutionStateChangedIterator struct {
	Event *EVM2EVMOffRampExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampExecutionStateChanged)
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
		it.Event = new(EVM2EVMOffRampExecutionStateChanged)
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

func (it *EVM2EVMOffRampExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampExecutionStateChanged struct {
	SequenceNumber uint64
	MessageId      [32]byte
	State          uint8
	Raw            types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64, messageId [][32]byte) (*EVM2EVMOffRampExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampExecutionStateChangedIterator{contract: _EVM2EVMOffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampExecutionStateChanged, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampExecutionStateChanged)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseExecutionStateChanged(log types.Log) (*EVM2EVMOffRampExecutionStateChanged, error) {
	event := new(EVM2EVMOffRampExecutionStateChanged)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampOffRampConfigChangedIterator struct {
	Event *EVM2EVMOffRampOffRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampOffRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampOffRampConfigChanged)
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
		it.Event = new(EVM2EVMOffRampOffRampConfigChanged)
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

func (it *EVM2EVMOffRampOffRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampOffRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampOffRampConfigChanged struct {
	Config        IEVM2EVMOffRampOffRampConfig
	SourceChainId uint64
	OnRamp        common.Address
	Raw           types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterOffRampConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampOffRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "OffRampConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampOffRampConfigChangedIterator{contract: _EVM2EVMOffRamp.contract, event: "OffRampConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchOffRampConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampOffRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "OffRampConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampOffRampConfigChanged)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "OffRampConfigChanged", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseOffRampConfigChanged(log types.Log) (*EVM2EVMOffRampOffRampConfigChanged, error) {
	event := new(EVM2EVMOffRampOffRampConfigChanged)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "OffRampConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMOffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMOffRampOwnershipTransferRequested)
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

func (it *EVM2EVMOffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampOwnershipTransferRequestedIterator{contract: _EVM2EVMOffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampOwnershipTransferRequested)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMOffRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMOffRampOwnershipTransferRequested)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampOwnershipTransferredIterator struct {
	Event *EVM2EVMOffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMOffRampOwnershipTransferred)
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

func (it *EVM2EVMOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampOwnershipTransferredIterator{contract: _EVM2EVMOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampOwnershipTransferred)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMOffRampOwnershipTransferred, error) {
	event := new(EVM2EVMOffRampOwnershipTransferred)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampPausedIterator struct {
	Event *EVM2EVMOffRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampPaused)
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
		it.Event = new(EVM2EVMOffRampPaused)
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

func (it *EVM2EVMOffRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMOffRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampPausedIterator{contract: _EVM2EVMOffRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampPaused)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParsePaused(log types.Log) (*EVM2EVMOffRampPaused, error) {
	event := new(EVM2EVMOffRampPaused)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampPoolAddedIterator struct {
	Event *EVM2EVMOffRampPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampPoolAdded)
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
		it.Event = new(EVM2EVMOffRampPoolAdded)
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

func (it *EVM2EVMOffRampPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMOffRampPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampPoolAddedIterator{contract: _EVM2EVMOffRamp.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampPoolAdded)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMOffRampPoolAdded, error) {
	event := new(EVM2EVMOffRampPoolAdded)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampPoolRemovedIterator struct {
	Event *EVM2EVMOffRampPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampPoolRemoved)
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
		it.Event = new(EVM2EVMOffRampPoolRemoved)
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

func (it *EVM2EVMOffRampPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMOffRampPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampPoolRemovedIterator{contract: _EVM2EVMOffRamp.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampPoolRemoved)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMOffRampPoolRemoved, error) {
	event := new(EVM2EVMOffRampPoolRemoved)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampSkippedIncorrectNonceIterator struct {
	Event *EVM2EVMOffRampSkippedIncorrectNonce

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampSkippedIncorrectNonceIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampSkippedIncorrectNonce)
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
		it.Event = new(EVM2EVMOffRampSkippedIncorrectNonce)
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

func (it *EVM2EVMOffRampSkippedIncorrectNonceIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampSkippedIncorrectNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampSkippedIncorrectNonce struct {
	Nonce  uint64
	Sender common.Address
	Raw    types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterSkippedIncorrectNonce(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampSkippedIncorrectNonceIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "SkippedIncorrectNonce", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampSkippedIncorrectNonceIterator{contract: _EVM2EVMOffRamp.contract, event: "SkippedIncorrectNonce", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchSkippedIncorrectNonce(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampSkippedIncorrectNonce, nonce []uint64, sender []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "SkippedIncorrectNonce", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampSkippedIncorrectNonce)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "SkippedIncorrectNonce", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseSkippedIncorrectNonce(log types.Log) (*EVM2EVMOffRampSkippedIncorrectNonce, error) {
	event := new(EVM2EVMOffRampSkippedIncorrectNonce)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "SkippedIncorrectNonce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampTokenPriceChangedIterator struct {
	Event *EVM2EVMOffRampTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampTokenPriceChanged)
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
		it.Event = new(EVM2EVMOffRampTokenPriceChanged)
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

func (it *EVM2EVMOffRampTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampTokenPriceChangedIterator{contract: _EVM2EVMOffRamp.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampTokenPriceChanged)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMOffRampTokenPriceChanged, error) {
	event := new(EVM2EVMOffRampTokenPriceChanged)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMOffRampTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMOffRampTokensRemovedFromBucket)
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

func (it *EVM2EVMOffRampTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMOffRampTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampTokensRemovedFromBucketIterator{contract: _EVM2EVMOffRamp.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampTokensRemovedFromBucket)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMOffRampTokensRemovedFromBucket, error) {
	event := new(EVM2EVMOffRampTokensRemovedFromBucket)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampTransmittedIterator struct {
	Event *EVM2EVMOffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampTransmitted)
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
		it.Event = new(EVM2EVMOffRampTransmitted)
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

func (it *EVM2EVMOffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMOffRampTransmittedIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampTransmittedIterator{contract: _EVM2EVMOffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTransmitted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampTransmitted)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseTransmitted(log types.Log) (*EVM2EVMOffRampTransmitted, error) {
	event := new(EVM2EVMOffRampTransmitted)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampUnpausedIterator struct {
	Event *EVM2EVMOffRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampUnpaused)
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
		it.Event = new(EVM2EVMOffRampUnpaused)
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

func (it *EVM2EVMOffRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMOffRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampUnpausedIterator{contract: _EVM2EVMOffRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampUnpaused)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMOffRampUnpaused, error) {
	event := new(EVM2EVMOffRampUnpaused)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMOffRamp.abi.Events["AFNSet"].ID:
		return _EVM2EVMOffRamp.ParseAFNSet(log)
	case _EVM2EVMOffRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMOffRamp.ParseConfigChanged(log)
	case _EVM2EVMOffRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMOffRamp.ParseConfigSet(log)
	case _EVM2EVMOffRamp.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMOffRamp.ParseExecutionStateChanged(log)
	case _EVM2EVMOffRamp.abi.Events["OffRampConfigChanged"].ID:
		return _EVM2EVMOffRamp.ParseOffRampConfigChanged(log)
	case _EVM2EVMOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMOffRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMOffRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMOffRamp.abi.Events["Paused"].ID:
		return _EVM2EVMOffRamp.ParsePaused(log)
	case _EVM2EVMOffRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMOffRamp.ParsePoolAdded(log)
	case _EVM2EVMOffRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMOffRamp.ParsePoolRemoved(log)
	case _EVM2EVMOffRamp.abi.Events["SkippedIncorrectNonce"].ID:
		return _EVM2EVMOffRamp.ParseSkippedIncorrectNonce(log)
	case _EVM2EVMOffRamp.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMOffRamp.ParseTokenPriceChanged(log)
	case _EVM2EVMOffRamp.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMOffRamp.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMOffRamp.abi.Events["Transmitted"].ID:
		return _EVM2EVMOffRamp.ParseTransmitted(log)
	case _EVM2EVMOffRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMOffRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMOffRampAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMOffRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMOffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x5e04b4755a4460aa6de4f3a906c4324a025c7449c02b52f5466659b5bfdfba5f")
}

func (EVM2EVMOffRampOffRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0xb4284598c699c1ce7d9b6db1b8f73ac462f762dc04449dc34344b293a142f617")
}

func (EVM2EVMOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMOffRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMOffRampPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMOffRampPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMOffRampSkippedIncorrectNonce) Topic() common.Hash {
	return common.HexToHash("0xd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf41237")
}

func (EVM2EVMOffRampTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMOffRampTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (EVM2EVMOffRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMOffRamp *EVM2EVMOffRamp) Address() common.Address {
	return _EVM2EVMOffRamp.address
}

type EVM2EVMOffRampInterface interface {
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

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, manualExecution bool) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report InternalExecutionReport) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetOffRampConfig(opts *bind.TransactOpts, config IEVM2EVMOffRampOffRampConfig) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error)

	SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMOffRampAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMOffRampAFNSet, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMOffRampConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMOffRampConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64, messageId [][32]byte) (*EVM2EVMOffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampExecutionStateChanged, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMOffRampExecutionStateChanged, error)

	FilterOffRampConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampOffRampConfigChangedIterator, error)

	WatchOffRampConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampOffRampConfigChanged) (event.Subscription, error)

	ParseOffRampConfigChanged(log types.Log) (*EVM2EVMOffRampOffRampConfigChanged, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMOffRampOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMOffRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMOffRampPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMOffRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMOffRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMOffRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMOffRampPoolRemoved, error)

	FilterSkippedIncorrectNonce(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampSkippedIncorrectNonceIterator, error)

	WatchSkippedIncorrectNonce(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampSkippedIncorrectNonce, nonce []uint64, sender []common.Address) (event.Subscription, error)

	ParseSkippedIncorrectNonce(log types.Log) (*EVM2EVMOffRampSkippedIncorrectNonce, error)

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMOffRampTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMOffRampTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMOffRampTokensRemovedFromBucket, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMOffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMOffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMOffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
