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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"InvalidOffRampConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedGasPriceUpdate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRampConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolByDestToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceFeeToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"feeTokenBaseUnitsPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structInternal.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOffRampConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620063a0380380620063a083398101604081905262000035916200091e565b6000805460ff19168155819084908490879033908190816200009e5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620000d857620000d881620003cf565b5050506001600160a01b0381166200010357604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03929092169190911790558051825114620001455760405162d8548360e71b815260040160405180910390fd5b81516200015a906005906020850190620006fe565b5060005b8251811015620002e357600060405180604001604052808484815181106200018a576200018a62000aa1565b60200260200101516001600160a01b03168152602001836001600160601b031681525090508060036000868581518110620001c957620001c962000aa1565b6020908102919091018101516001600160a01b0390811683528282019390935260409091016000908120845194909201516001600160601b0316600160a01b02939092169290921790915581518451909160049186908690811062000232576200023262000aa1565b60200260200101516001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000278573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200029e919062000ab7565b6001600160a01b039081168252602082019290925260400160002080546001600160a01b0319169290911691909117905550620002db8162000ade565b90506200015e565b50508151600680546001600160a01b0319166001600160a01b0392831617905560408051608081018252602080860151808352958301805191830182905251928201839052426060909201829052600995909555600a94909455600b55600c9290925550861662000367576040516342bcdf7f60e11b815260040160405180910390fd5b620003728562000480565b6001600160401b03808916608052871660a0526001600160a01b03861660c052620003bd7fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a62000698565b60e0525062000b069650505050505050565b336001600160a01b03821603620004295760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000095565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60608101516001600160a01b03161580620004a6575060a08101516001600160a01b0316155b80620004ba575080516001600160a01b0316155b156200053f5760408051630e48a3c160e01b815282516001600160a01b039081166004830152602084015163ffffffff166024830152918301516001600160401b03908116604483015260608401518316606483015260808401518116608483015260a084015190921660a482015260c083015190911660c482015260e40162000095565b80516013805460208401516040808601516001600160401b03908116600160c01b026001600160c01b0363ffffffff909416600160a01b9081026001600160c01b03199096166001600160a01b03988916179590951793909316929092179093556060850151601480546080880151841685026001600160e01b03199182169388169390931792909217905560a08601516015805460c08901519094169094029290911694169390931792909217909155517f08fd01f552eb5d1aaca89b8c44ca4f2c0693ba70658968601e9c8de5714d8062906200068d90839081516001600160a01b03908116825260208084015163ffffffff16908301526040808401516001600160401b039081169184019190915260608085015183169084015260808085015182169084015260a0848101519092169183019190915260c092830151169181019190915260e00190565b60405180910390a150565b60008160805160a05160c051604051602001620006e194939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b82805482825590600052602060002090810192821562000756579160200282015b828111156200075657825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906200071f565b506200076492915062000768565b5090565b5b8082111562000764576000815560010162000769565b80516001600160401b03811681146200079757600080fd5b919050565b6001600160a01b0381168114620007b257600080fd5b50565b805162000797816200079c565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b0381118282101715620007fd57620007fd620007c2565b60405290565b600082601f8301126200081557600080fd5b815160206001600160401b0380831115620008345762000834620007c2565b8260051b604051601f19603f830116810181811084821117156200085c576200085c620007c2565b6040529384528581018301938381019250878511156200087b57600080fd5b83870191505b84821015620008a757815162000897816200079c565b8352918301919083019062000881565b979650505050505050565b600060608284031215620008c557600080fd5b604051606081016001600160401b0381118282101715620008ea57620008ea620007c2565b80604052508091508251620008ff816200079c565b8082525060208301516020820152604083015160408201525092915050565b600080600080600080600080888a036102008112156200093d57600080fd5b620009488a6200077f565b98506200095860208b016200077f565b975060408a01516200096a816200079c565b965060e0605f19820112156200097f57600080fd5b506200098a620007d8565b60608a01516200099a816200079c565b815260808a015163ffffffff81168114620009b457600080fd5b6020820152620009c760a08b016200077f565b6040820152620009da60c08b01620007b5565b6060820152620009ed60e08b016200077f565b608082015262000a016101008b01620007b5565b60a082015262000a156101208b016200077f565b60c0820152945062000a2b6101408a01620007b5565b6101608a01519094506001600160401b038082111562000a4a57600080fd5b62000a588c838d0162000803565b94506101808b015191508082111562000a7057600080fd5b5062000a7f8b828c0162000803565b92505062000a928a6101a08b01620008b2565b90509295985092959890939650565b634e487b7160e01b600052603260045260246000fd5b60006020828403121562000aca57600080fd5b815162000ad7816200079c565b9392505050565b60006001820162000aff57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e05161585c62000b4460003960006134ba01526000505060006102950152600081816102700152613ba9015261585c6000f3fe608060405234801561001057600080fd5b506004361061025c5760003560e01c806379ba509711610145578063addb9e19116100bd578063cb3c8c011161008c578063d7e2bb5011610071578063d7e2bb5014610748578063eb511dd414610774578063f2fde38b1461078757600080fd5b8063cb3c8c011461072d578063d3c7c2c71461074057600080fd5b8063addb9e19146106d4578063afcb95d7146106e7578063b1dc65a414610707578063b4069b311461071a57600080fd5b8063856c82471161011457806390c2339b116100f957806390c2339b14610543578063a8ebd0f41461057e578063abc39f1f146106c157600080fd5b8063856c8247146104e15780638da5cb5b1461052d57600080fd5b806379ba50971461049357806381ff70481461049b5780638456cb59146104cb57806385572ffb146104d357600080fd5b80633f4ba83a116101d85780635c975abb116101a7578063666cab8d1161018c578063666cab8d14610456578063681fba161461046b578063744b92e21461048057600080fd5b80635c975abb1461041f5780635d86f1411461042a57600080fd5b80633f4ba83a146103d35780634352fa9f146103db5780634741062e146103ee578063599f64311461040e57600080fd5b806317332f9b1161022f5780631ef38174116102145780631ef38174146103885780632222dd421461039b57806339aa9264146103c057600080fd5b806317332f9b1461032c578063181f5a771461033f57600080fd5b8063087ae6df14610261578063108ee5fc146102c5578063142a98fc146102da578063147809b314610314575b600080fd5b6040805167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811682527f0000000000000000000000000000000000000000000000000000000000000000166020820152015b60405180910390f35b6102d86102d3366004613ef2565b61079a565b005b6103076102e8366004613f35565b67ffffffffffffffff1660009081526017602052604090205460ff1690565b6040516102bc9190613f68565b61031c610851565b60405190151581526020016102bc565b6102d861033a366004614098565b6108de565b61037b6040518060400160405280601481526020017f45564d3245564d4f666652616d7020312e302e3000000000000000000000000081525081565b6040516102bc9190614156565b6102d861039636600461428b565b610a19565b6002546001600160a01b03165b6040516001600160a01b0390911681526020016102bc565b6102d86103ce366004613ef2565b61107f565b6102d86110b6565b6102d86103e93660046143bc565b6110c8565b6104016103fc366004614477565b61131d565b6040516102bc91906144b4565b6006546001600160a01b03166103a8565b60005460ff1661031c565b6103a8610438366004613ef2565b6001600160a01b039081166000908152600360205260409020541690565b61045e6113e5565b6040516102bc919061453c565b610473611447565b6040516102bc919061454f565b6102d861048e366004614590565b61150c565b6102d8611862565b600f54600d546040805163ffffffff808516825264010000000090940490931660208401528201526060016102bc565b6102d8611945565b6102d861025c3660046145c9565b6105146104ef366004613ef2565b6001600160a01b031660009081526016602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff90911681526020016102bc565b60005461010090046001600160a01b03166103a8565b61054b611955565b6040516102bc91908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6106446040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152506040805160e0810182526013546001600160a01b038082168352600160a01b80830463ffffffff166020850152780100000000000000000000000000000000000000000000000090920467ffffffffffffffff9081169484019490945260145480821660608501528290048416608084015260155490811660a08401520490911660c082015290565b6040516102bc9190600060e0820190506001600160a01b0380845116835263ffffffff6020850151166020840152604084015167ffffffffffffffff80821660408601528260608701511660608601528060808701511660808601528260a08701511660a08601528060c08701511660c086015250505092915050565b6102d86106cf3660046146a5565b6119f5565b6102d86106e23660046147e9565b611bae565b6040805160018152600060208201819052918101919091526060016102bc565b6102d86107153660046148d4565b611bc2565b6103a8610728366004613ef2565b612112565b6102d861073b366004614bc5565b612200565b61047361220b565b6103a8610756366004613ef2565b6001600160a01b039081166000908152600460205260409020541690565b6102d8610782366004614590565b61226b565b6102d8610795366004613ef2565b6124c5565b6107a26124d6565b6001600160a01b0381166107e2576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa1580156108b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d89190614cb6565b15905090565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561091b57506006546001600160a01b03163314155b15610952576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602081015179ffffffffffffffffffffffffffffffffffffffffffffffffffff116109a9576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109b36009612535565b6040810151600a8190556020820151600955600b546109d291906125e0565b600b556040818101516020808401518351928352908201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b855185518560ff16601f831115610a91576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064015b60405180910390fd5b80600003610afb576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610a88565b818314610b89576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610a88565b610b94816003614ce9565b8311610bfc576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610a88565b610c046124d6565b60115460005b81811015610cac576010600060118381548110610c2957610c29614d08565b60009182526020808320909101546001600160a01b031683528201929092526040018120805461ffff1916905560128054601092919084908110610c6f57610c6f614d08565b60009182526020808320909101546001600160a01b031683528201929092526040019020805461ffff19169055610ca581614d1e565b9050610c0a565b50895160005b81811015610f415760008c8281518110610cce57610cce614d08565b6020026020010151905060006002811115610ceb57610ceb613f52565b6001600160a01b038216600090815260106020526040902054610100900460ff166002811115610d1d57610d1d613f52565b14610d84576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610a88565b6040805180820190915260ff8316815260208101600190526001600160a01b03821660009081526010602090815260409091208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610dec57610dec613f52565b021790555090505060008c8381518110610e0857610e08614d08565b6020026020010151905060006002811115610e2557610e25613f52565b6001600160a01b038216600090815260106020526040902054610100900460ff166002811115610e5757610e57613f52565b14610ebe576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610a88565b6040805180820190915260ff8416815260208101600290526001600160a01b03821660009081526010602090815260409091208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115610f2657610f26613f52565b0217905550905050505080610f3a90614d1e565b9050610cb2565b508a51610f559060119060208e0190613e3b565b508951610f699060129060208d0190613e3b565b50600e805460ff8381166101000261ffff19909216908c1617179055600f8054610fd2914691309190600090610fa49063ffffffff16614d38565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168e8e8e8e8e8e6125fa565b600d600001819055506000600f60049054906101000a900463ffffffff16905043600f60046101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581600d60000154600f60009054906101000a900463ffffffff168f8f8f8f8f8f60405161106999989796959493929190614d5b565b60405180910390a1505050505050505050505050565b6110876124d6565b6006805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6110be6124d6565b6110c6612687565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561110557506006546001600160a01b03163314155b1561113c576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815181518114611178576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085460005b818110156111d257600760006008838154811061119d5761119d614d08565b60009182526020808320909101546001600160a01b031683528201929092526040018120556111cb81614d1e565b905061117e565b5060005b828110156113025760008582815181106111f2576111f2614d08565b6020026020010151905060006001600160a01b0316816001600160a01b031603611248576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84828151811061125a5761125a614d08565b602002602001015160076000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f818684815181106112bf576112bf614d08565b60200260200101516040516112e99291906001600160a01b03929092168252602082015260400190565b60405180910390a1506112fb81614d1e565b90506111d6565b508351611316906008906020870190613e3b565b5050505050565b80516060908067ffffffffffffffff81111561133b5761133b613f90565b604051908082528060200260200182016040528015611364578160200160208202803683370190505b50915060005b818110156113de576007600085838151811061138857611388614d08565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020548382815181106113c3576113c3614d08565b60209081029190910101526113d781614d1e565b905061136a565b5050919050565b6060601280548060200260200160405190810160405280929190818152602001828054801561143d57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161141f575b5050505050905090565b60055460609067ffffffffffffffff81111561146557611465613f90565b60405190808252806020026020018201604052801561148e578160200160208202803683370190505b50905060005b600554811015611508576114ce600582815481106114b4576114b4614d08565b6000918252602090912001546001600160a01b0316612112565b8282815181106114e0576114e0614d08565b6001600160a01b039092166020928302919091019091015261150181614d1e565b9050611494565b5090565b6115146124d6565b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352600160a01b9093046bffffffffffffffffffffffff169082015290611591576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816001600160a01b031681600001516001600160a01b0316146115e0576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058054906000906115f3600184614df1565b8154811061160357611603614d08565b9060005260206000200160009054906101000a90046001600160a01b03169050600583602001516bffffffffffffffffffffffff168154811061164857611648614d08565b6000918252602090912001546001600160a01b0316600561166a600185614df1565b8154811061167a5761167a614d08565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600584602001516bffffffffffffffffffffffff16815481106116ce576116ce614d08565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558581015184841683526003909152604090912080546bffffffffffffffffffffffff909216600160a01b0291909216179055600580548061174757611747614e08565b6001900381819060005260206000200160006101000a8154906001600160a01b030219169055905560046000856001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117d59190614e1e565b6001600160a01b03908116825260208083019390935260409182016000908120805473ffffffffffffffffffffffffffffffffffffffff1916905588821680825260038552838220919091558251908152908716928101929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146118bc5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610a88565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61194d6124d6565b6110c6612723565b6119806040518060800160405280600081526020016000815260200160008152602001600081525090565b604080516080810182526009548152600a546020820152600b5491810191909152600c54606082018190526000906119b89042614df1565b602083015183519192506119e4916119d09084614ce9565b84604001516119df9190614e3b565b6125e0565b604083015250426060820152919050565b333014611a2e576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160008082526020820190925281611a6b565b6040805180820190915260008082526020820152815260200190600190039081611a445790505b506101208401515190915015611a9157611a8e8361012001518460e001516127ab565b90505b60e08301516001600160a01b03163b1580611ae1575060e0830151611adf906001600160a01b03167f85572ffb000000000000000000000000000000000000000000000000000000006129d9565b155b15611aeb57505050565b6014546001600160a01b0316635607b375611b0685846129f5565b848660a001518760e001516040518563ffffffff1660e01b8152600401611b309493929190614e97565b6020604051808303816000875af1158015611b4f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b739190614cb6565b611ba9576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b611bb66124d6565b611bbf81612a98565b50565b611c0187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612d1392505050565b60408051606081018252600d54808252600e5460ff808216602085015261010090910416928201929092528935918214611c745780516040517f93df584c000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610a88565b6040805183815260208c81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a1600281602001518260400151611ccf9190614f63565b611cd99190614f9e565b611ce4906001614f63565b60ff168614611d1f576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b858414611d58576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526010602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115611d9b57611d9b613f52565b6002811115611dac57611dac613f52565b9052509050600281602001516002811115611dc957611dc9613f52565b148015611e0357506012816000015160ff1681548110611deb57611deb614d08565b6000918252602090912001546001600160a01b031633145b611e39576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506000611e47856020614ce9565b611e52886020614ce9565b611e5e8b610144614e3b565b611e689190614e3b565b611e729190614e3b565b9050368114611eb6576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401610a88565b5060008989604051611ec9929190614fc0565b604051908190038120611ee0918d90602001614fd0565b604051602081830303815290604052805190602001209050611f00613ea9565b8760005b81811015612102576000600185898460208110611f2357611f23614d08565b611f3091901a601b614f63565b8e8e86818110611f4257611f42614d08565b905060200201358d8d87818110611f5b57611f5b614d08565b9050602002013560405160008152602001604052604051611f98949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611fba573d6000803e3d6000fd5b505060408051601f198101516001600160a01b038116600090815260106020908152848220848601909552845460ff808216865293975091955092939284019161010090910416600281111561201257612012613f52565b600281111561202357612023613f52565b905250905060018160200151600281111561204057612040613f52565b14612077576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051859060ff16601f811061208e5761208e614d08565b6020020151156120ca576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600185826000015160ff16601f81106120e5576120e5614d08565b91151560209092020152506120fb905081614d1e565b9050611f04565b5050505050505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680612166576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa1580156121d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121f99190614e1e565b9392505050565b611bbf816001612d2d565b6060600580548060200260200160405190810160405280929190818152602001828054801561143d576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161141f575050505050905090565b6122736124d6565b6001600160a01b038216158061229057506001600160a01b038116155b156122c7576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352600160a01b9093046bffffffffffffffffffffffff16908201529015612345576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038083168083526005546bffffffffffffffffffffffff9081166020808601918252878516600090815260038252604080822088519451909516600160a01b02939096169290921790925583517f21df0da70000000000000000000000000000000000000000000000000000000081529351869460049492936321df0da79282870192819003870181865afa1580156123e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061240d9190614e1e565b6001600160a01b03908116825260208083019390935260409182016000908120805495831673ffffffffffffffffffffffffffffffffffffffff199687161790556005805460018101825591527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001805488831695168517905581519384528516918301919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b6124cd6124d6565b611bbf81612f08565b60005461010090046001600160a01b031633146110c65760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610a88565b80600101548160020154148061254e5750428160030154145b156125565750565b806001015481600201541115612598576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008160030154426125aa9190614df1565b600183015483549192506125d1916125c29084614ce9565b84600201546119df9190614e3b565b60028301555042600390910155565b60008183106125ef57816125f1565b825b90505b92915050565b6000808a8a8a8a8a8a8a8a8a60405160200161261e99989796959493929190614fec565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b60005460ff166126d95760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610a88565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60005460ff16156127765760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610a88565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586127063390565b60606000835167ffffffffffffffff8111156127c9576127c9613f90565b60405190808252806020026020018201604052801561280e57816020015b60408051808201909152600080825260208201528152602001906001900390816127e75790505b50905060005b84518110156129cf57600061285e86838151811061283457612834614d08565b6020026020010151600001516001600160a01b039081166000908152600360205260409020541690565b90506001600160a01b0381166128ca5785828151811061288057612880614d08565b6020908102919091010151516040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a88565b6128f2818784815181106128e0576128e0614d08565b60200260200101516020015187612fc4565b806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612930573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129549190614e1e565b83838151811061296657612966614d08565b60209081029190910101516001600160a01b039091169052855186908390811061299257612992614d08565b6020026020010151602001518383815181106129b0576129b0614d08565b6020908102919091018101510152506129c881614d1e565b9050612814565b506125f181613044565b60006129e483613248565b80156125f157506125f183836132ac565b6040805160a08101825260008082526020820152606091810182905281810182905260808101919091526040518060a001604052808461016001518152602001846000015167ffffffffffffffff1681526020018460600151604051602001612a6d91906001600160a01b0391909116815260200190565b6040516020818303038152906040528152602001846101000151815260200183815250905092915050565b60608101516001600160a01b03161580612abd575060a08101516001600160a01b0316155b80612ad0575080516001600160a01b0316155b15612b6d57604080517f0e48a3c100000000000000000000000000000000000000000000000000000000815282516001600160a01b039081166004830152602084015163ffffffff1660248301529183015167ffffffffffffffff908116604483015260608401518316606483015260808401518116608483015260a084015190921660a482015260c083015190911660c482015260e401610a88565b805160138054602084015160408086015167ffffffffffffffff90811678010000000000000000000000000000000000000000000000000277ffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff909416600160a01b9081027fffffffffffffffff0000000000000000000000000000000000000000000000009096166001600160a01b03988916179590951793909316929092179093556060850151601480546080880151841685027fffffffff000000000000000000000000000000000000000000000000000000009182169388169390931792909217905560a08601516015805460c08901519094169094029290911694169390931792909217909155517f08fd01f552eb5d1aaca89b8c44ca4f2c0693ba70658968601e9c8de5714d806290610a0e908390600060e0820190506001600160a01b0380845116835263ffffffff6020850151166020840152604084015167ffffffffffffffff80821660408601528260608701511660608601528060808701511660808601528260a08701511660a08601528060c08701511660c086015250505092915050565b611bbf81806020019051810190612d2a91906152b0565b60005b60005460ff1615612d805760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610a88565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612dd3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612df79190614cb6565b15612e2d576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208201515115612eef578015612e70576040517f198753d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60135460208301516040517f9086658e0000000000000000000000000000000000000000000000000000000081526001600160a01b0390921691639086658e91612ebc91600401615396565b600060405180830381600087803b158015612ed657600080fd5b505af1158015612eea573d6000803e3d6000fd5b505050505b60408201515115612f0457612f04828261337b565b5050565b336001600160a01b03821603612f605760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610a88565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b15801561302757600080fd5b505af115801561303b573d6000803e3d6000fd5b50505050505050565b6000805b82518110156131435760006007600085848151811061306957613069614d08565b6020026020010151600001516001600160a01b03166001600160a01b03168152602001908152602001600020549050806000036130fc578382815181106130b2576130b2614d08565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610a88565b83828151811061310e5761310e614d08565b602002602001015160200151816131259190614ce9565b61312f9084614e3b565b9250508061313c90614d1e565b9050613048565b508015612f04576131546009612535565b600a5481111561319e57600a546040517f688ccf77000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610a88565b600b548111156131fe57600954600b54600091906131bc9084614df1565b6131c69190615414565b9050806040517fe31e0f32000000000000000000000000000000000000000000000000000000008152600401610a8891815260200190565b80600960020160008282546132139190614df1565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba37890602001610845565b6000613274827f01ffc9a7000000000000000000000000000000000000000000000000000000006132ac565b80156125f457506132a5827fffffffff000000000000000000000000000000000000000000000000000000006132ac565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015613364575060208210155b80156133705750600081115b979650505050505050565b60408201515160008167ffffffffffffffff81111561339c5761339c613f90565b6040519080825280602002602001820160405280156133c5578160200160208202803683370190505b50905060008267ffffffffffffffff8111156133e3576133e3613f90565b60405190808252806020026020018201604052801561347157816020015b60408051610180810182526000808252602080830182905292820181905260608083018290526080830182905260a0830182905260c0830182905260e083018290526101008301819052610120830152610140820181905261016082015282526000199092019101816134015790505b50905060005b8381101561352c5760008660400151828151811061349757613497614d08565b60200260200101518060200190518101906134b291906154b0565b90506134de817f0000000000000000000000000000000000000000000000000000000000000000613a9d565b8483815181106134f0576134f0614d08565b6020026020010181815250508083838151811061350f5761350f614d08565b6020026020010181905250508061352590614d1e565b9050613477565b50601554606086015160808701516040517f320488750000000000000000000000000000000000000000000000000000000081526000936001600160a01b0316926332048875926135839288929190600401615613565b6020604051808303816000875af11580156135a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135c69190615649565b905060008111613602576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8481101561303b57600083828151811061362157613621614d08565b602002602001015190506000613654826020015167ffffffffffffffff1660009081526017602052604090205460ff1690565b9050600081600381111561366a5761366a613f52565b14806136875750600381600381111561368557613685613f52565b145b6136cf5760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a88565b871561374e57601354600090600160a01b900463ffffffff166136f28642614df1565b11905080806137125750600382600381111561371057613710613f52565b145b613748576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506137ab565b600081600381111561376257613762613f52565b146137ab5760208201516040517f67d9ba0f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a88565b60008160038111156137bf576137bf613f52565b0361385e57608082015160608301516001600160a01b031660009081526016602052604090205467ffffffffffffffff918216916137ff91166001615662565b67ffffffffffffffff161461385e5781606001516001600160a01b0316826080015167ffffffffffffffff167fd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf4123760405160405180910390a35050613a8d565b61386782613ba7565b60208281015167ffffffffffffffff166000908152601790915260408120805460ff1916600117905561389a838a613d03565b60208085015167ffffffffffffffff1660009081526017909152604090208054919250829160ff191660018360038111156138d7576138d7613f52565b021790555088156139bd578260c0015180156139045750600382600381111561390257613902613f52565b145b80156139215750600281600381111561391f5761391f613f52565b145b806139595750600082600381111561393b5761393b613f52565b1480156139595750600281600381111561395757613957613f52565b145b156139b85760608301516001600160a01b03166000908152601660205260408120805467ffffffffffffffff16916139908361568e565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b613a3d565b8260c0015180156139df575060038160038111156139dd576139dd613f52565b145b613a3d5760608301516001600160a01b03166000908152601660205260408120805467ffffffffffffffff1691613a158361568e565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b826101600151836020015167ffffffffffffffff167f5e04b4755a4460aa6de4f3a906c4324a025c7449c02b52f5466659b5bfdfba5f83604051613a819190613f68565b60405180910390a35050505b613a9681614d1e565b9050613605565b60008060001b828460200151856080015186606001518760e0015188610100015180519060200120896101200151604051602001613adb91906156ab565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d60400151604051602001613b899c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16816000015167ffffffffffffffff1614613c275780516040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a88565b60155461012082015151600160a01b90910467ffffffffffffffff161015613c8d5760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a88565b60145461010082015151600160a01b90910467ffffffffffffffff161015611bbf57601454610100820151516040517f86933789000000000000000000000000000000000000000000000000000000008152600160a01b90920467ffffffffffffffff1660048301526024820152604401610a88565b6040517fabc39f1f000000000000000000000000000000000000000000000000000000008152600090309063abc39f1f90613d4490869086906004016156be565b600060405180830381600087803b158015613d5e57600080fd5b505af1925050508015613d6f575060015b613e32573d808015613d9d576040519150601f19603f3d011682016040523d82523d6000602084013e613da2565b606091505b50613dac816157ff565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003613dfe5760039150506125f4565b806040517fcf19edfd000000000000000000000000000000000000000000000000000000008152600401610a889190614156565b50600292915050565b828054828255906000526020600020908101928215613e9d579160200282015b82811115613e9d578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190613e5b565b50611508929150613ec8565b604051806103e00160405280601f906020820280368337509192915050565b5b808211156115085760008155600101613ec9565b6001600160a01b0381168114611bbf57600080fd5b600060208284031215613f0457600080fd5b81356125f181613edd565b67ffffffffffffffff81168114611bbf57600080fd5b8035613f3081613f0f565b919050565b600060208284031215613f4757600080fd5b81356125f181613f0f565b634e487b7160e01b600052602160045260246000fd5b6020810160048310613f8a57634e487b7160e01b600052602160045260246000fd5b91905290565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613fc957613fc9613f90565b60405290565b604051610180810167ffffffffffffffff81118282101715613fc957613fc9613f90565b60405160e0810167ffffffffffffffff81118282101715613fc957613fc9613f90565b6040516060810167ffffffffffffffff81118282101715613fc957613fc9613f90565b60405160a0810167ffffffffffffffff81118282101715613fc957613fc9613f90565b604051601f8201601f1916810167ffffffffffffffff8111828210171561408557614085613f90565b604052919050565b8035613f3081613edd565b6000606082840312156140aa57600080fd5b6040516060810181811067ffffffffffffffff821117156140cd576140cd613f90565b60405282356140db81613edd565b8152602083810135908201526040928301359281019290925250919050565b60005b838110156141155781810151838201526020016140fd565b83811115614124576000848401525b50505050565b600081518084526141428160208601602086016140fa565b601f01601f19169290920160200192915050565b6020815260006125f1602083018461412a565b600067ffffffffffffffff82111561418357614183613f90565b5060051b60200190565b600082601f83011261419e57600080fd5b813560206141b36141ae83614169565b61405c565b82815260059290921b840181019181810190868411156141d257600080fd5b8286015b848110156141f65780356141e981613edd565b83529183019183016141d6565b509695505050505050565b803560ff81168114613f3057600080fd5b600067ffffffffffffffff82111561422c5761422c613f90565b50601f01601f191660200190565b600082601f83011261424b57600080fd5b81356142596141ae82614212565b81815284602083860101111561426e57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c087890312156142a457600080fd5b863567ffffffffffffffff808211156142bc57600080fd5b6142c88a838b0161418d565b975060208901359150808211156142de57600080fd5b6142ea8a838b0161418d565b96506142f860408a01614201565b9550606089013591508082111561430e57600080fd5b61431a8a838b0161423a565b945061432860808a01613f25565b935060a089013591508082111561433e57600080fd5b5061434b89828a0161423a565b9150509295509295509295565b600082601f83011261436957600080fd5b813560206143796141ae83614169565b82815260059290921b8401810191818101908684111561439857600080fd5b8286015b848110156141f65780356143af81613edd565b835291830191830161439c565b600080604083850312156143cf57600080fd5b823567ffffffffffffffff808211156143e757600080fd5b6143f386838701614358565b935060209150818501358181111561440a57600080fd5b85019050601f8101861361441d57600080fd5b803561442b6141ae82614169565b81815260059190911b8201830190838101908883111561444a57600080fd5b928401925b828410156144685783358252928401929084019061444f565b80955050505050509250929050565b60006020828403121561448957600080fd5b813567ffffffffffffffff8111156144a057600080fd5b6144ac84828501614358565b949350505050565b6020808252825182820181905260009190848201906040850190845b818110156144ec578351835292840192918401916001016144d0565b50909695505050505050565b600081518084526020808501945080840160005b838110156145315781516001600160a01b03168752958201959082019060010161450c565b509495945050505050565b6020815260006125f160208301846144f8565b6020808252825182820181905260009190848201906040850190845b818110156144ec5783516001600160a01b03168352928401929184019160010161456b565b600080604083850312156145a357600080fd5b82356145ae81613edd565b915060208301356145be81613edd565b809150509250929050565b6000602082840312156145db57600080fd5b813567ffffffffffffffff8111156145f257600080fd5b820160a081850312156125f157600080fd5b8015158114611bbf57600080fd5b8035613f3081614604565b600082601f83011261462e57600080fd5b8135602061463e6141ae83614169565b82815260069290921b8401810191818101908684111561465d57600080fd5b8286015b848110156141f6576040818903121561467a5760008081fd5b614682613fa6565b813561468d81613edd565b81528185013585820152835291830191604001614661565b600080604083850312156146b857600080fd5b823567ffffffffffffffff808211156146d057600080fd5b9084019061018082870312156146e557600080fd5b6146ed613fcf565b6146f683613f25565b815261470460208401613f25565b60208201526040830135604082015261471f6060840161408d565b606082015261473060808401613f25565b608082015260a083013560a082015261474b60c08401614612565b60c082015261475c60e0840161408d565b60e0820152610100808401358381111561477557600080fd5b6147818982870161423a565b828401525050610120808401358381111561479b57600080fd5b6147a78982870161461d565b82840152505061014091506147bd82840161408d565b82820152610160915081830135828201528094505050506147e060208401614612565b90509250929050565b600060e082840312156147fb57600080fd5b614803613ff3565b823561480e81613edd565b8152602083013563ffffffff8116811461482757600080fd5b602082015261483860408401613f25565b60408201526148496060840161408d565b606082015261485a60808401613f25565b608082015261486b60a0840161408d565b60a082015261487c60c08401613f25565b60c08201529392505050565b60008083601f84011261489a57600080fd5b50813567ffffffffffffffff8111156148b257600080fd5b6020830191508360208260051b85010111156148cd57600080fd5b9250929050565b60008060008060008060008060e0898b0312156148f057600080fd5b606089018a81111561490157600080fd5b8998503567ffffffffffffffff8082111561491b57600080fd5b818b0191508b601f83011261492f57600080fd5b81358181111561493e57600080fd5b8c602082850101111561495057600080fd5b6020830199508098505060808b013591508082111561496e57600080fd5b61497a8c838d01614888565b909750955060a08b013591508082111561499357600080fd5b506149a08b828c01614888565b999c989b50969995989497949560c00135949350505050565b600082601f8301126149ca57600080fd5b813560206149da6141ae83614169565b82815260059290921b840181019181810190868411156149f957600080fd5b8286015b848110156141f6578035614a1081613f0f565b83529183019183016149fd565b6fffffffffffffffffffffffffffffffff81168114611bbf57600080fd5b600082601f830112614a4c57600080fd5b81356020614a5c6141ae83614169565b82815260609283028501820192828201919087851115614a7b57600080fd5b8387015b85811015614add5781818a031215614a975760008081fd5b614a9f614016565b8135614aaa81613edd565b815281860135614ab981613f0f565b81870152604082810135614acc81614a1d565b908201528452928401928101614a7f565b5090979650505050505050565b600082601f830112614afb57600080fd5b81356020614b0b6141ae83614169565b82815260059290921b84018101918181019086841115614b2a57600080fd5b8286015b848110156141f657803567ffffffffffffffff811115614b4e5760008081fd5b614b5c8986838b010161423a565b845250918301918301614b2e565b600082601f830112614b7b57600080fd5b81356020614b8b6141ae83614169565b82815260059290921b84018101918181019086841115614baa57600080fd5b8286015b848110156141f65780358352918301918301614bae565b600060208284031215614bd757600080fd5b813567ffffffffffffffff80821115614bef57600080fd5b9083019060a08286031215614c0357600080fd5b614c0b614039565b823582811115614c1a57600080fd5b614c26878286016149b9565b825250602083013582811115614c3b57600080fd5b614c4787828601614a3b565b602083015250604083013582811115614c5f57600080fd5b614c6b87828601614aea565b604083015250606083013582811115614c8357600080fd5b614c8f87828601614b6a565b6060830152506080830135608082015280935050505092915050565b8051613f3081614604565b600060208284031215614cc857600080fd5b81516125f181614604565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615614d0357614d03614cd3565b500290565b634e487b7160e01b600052603260045260246000fd5b60006000198203614d3157614d31614cd3565b5060010190565b600063ffffffff808316818103614d5157614d51614cd3565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614d8b8184018a6144f8565b90508281036080840152614d9f81896144f8565b905060ff871660a084015282810360c0840152614dbc818761412a565b905067ffffffffffffffff851660e0840152828103610100840152614de1818561412a565b9c9b505050505050505050505050565b600082821015614e0357614e03614cd3565b500390565b634e487b7160e01b600052603160045260246000fd5b600060208284031215614e3057600080fd5b81516125f181613edd565b60008219821115614e4e57614e4e614cd3565b500190565b600081518084526020808501945080840160005b8381101561453157815180516001600160a01b031688528301518388015260409096019590820190600101614e67565b608081528451608082015267ffffffffffffffff60208601511660a08201526000604086015160a060c0840152614ed261012084018261412a565b905060608701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80808584030160e0860152614f0e838361412a565b925060808901519150808584030161010086015250614f2d8282614e53565b92505050614f3f602083018615159052565b836040830152614f5a60608301846001600160a01b03169052565b95945050505050565b600060ff821660ff84168060ff03821115614f8057614f80614cd3565b019392505050565b634e487b7160e01b600052601260045260246000fd5b600060ff831680614fb157614fb1614f88565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526150268285018b6144f8565b9150838203608085015261503a828a6144f8565b915060ff881660a085015283820360c0850152615057828861412a565b90861660e08501528381036101008501529050614de1818561412a565b8051613f3081613f0f565b600082601f83011261509057600080fd5b815160206150a06141ae83614169565b82815260059290921b840181019181810190868411156150bf57600080fd5b8286015b848110156141f65780516150d681613f0f565b83529183019183016150c3565b8051613f3081613edd565b600082601f8301126150ff57600080fd5b8151602061510f6141ae83614169565b8281526060928302850182019282820191908785111561512e57600080fd5b8387015b85811015614add5781818a03121561514a5760008081fd5b615152614016565b815161515d81613edd565b81528186015161516c81613f0f565b8187015260408281015161517f81614a1d565b908201528452928401928101615132565b600082601f8301126151a157600080fd5b81516151af6141ae82614212565b8181528460208386010111156151c457600080fd5b6144ac8260208301602087016140fa565b600082601f8301126151e657600080fd5b815160206151f66141ae83614169565b82815260059290921b8401810191818101908684111561521557600080fd5b8286015b848110156141f657805167ffffffffffffffff8111156152395760008081fd5b6152478986838b0101615190565b845250918301918301615219565b600082601f83011261526657600080fd5b815160206152766141ae83614169565b82815260059290921b8401810191818101908684111561529557600080fd5b8286015b848110156141f65780518352918301918301615299565b6000602082840312156152c257600080fd5b815167ffffffffffffffff808211156152da57600080fd5b9083019060a082860312156152ee57600080fd5b6152f6614039565b82518281111561530557600080fd5b6153118782860161507f565b82525060208301518281111561532657600080fd5b615332878286016150ee565b60208301525060408301518281111561534a57600080fd5b615356878286016151d5565b60408301525060608301518281111561536e57600080fd5b61537a87828601615255565b6060830152506080830151608082015280935050505092915050565b602080825282518282018190526000919060409081850190868401855b8281101561540757815180516001600160a01b031685528681015167ffffffffffffffff16878601528501516fffffffffffffffffffffffffffffffff1685850152606090930192908501906001016153b3565b5091979650505050505050565b60008261542357615423614f88565b500490565b600082601f83011261543957600080fd5b815160206154496141ae83614169565b82815260069290921b8401810191818101908684111561546857600080fd5b8286015b848110156141f657604081890312156154855760008081fd5b61548d613fa6565b815161549881613edd565b8152818501518582015283529183019160400161546c565b6000602082840312156154c257600080fd5b815167ffffffffffffffff808211156154da57600080fd5b9083019061018082860312156154ef57600080fd5b6154f7613fcf565b61550083615074565b815261550e60208401615074565b602082015260408301516040820152615529606084016150e3565b606082015261553a60808401615074565b608082015260a083015160a082015261555560c08401614cab565b60c082015261556660e084016150e3565b60e0820152610100808401518381111561557f57600080fd5b61558b88828701615190565b82840152505061012080840151838111156155a557600080fd5b6155b188828701615428565b82840152505061014091506155c78284016150e3565b9181019190915261016091820151918101919091529392505050565b600081518084526020808501945080840160005b83811015614531578151875295820195908201906001016155f7565b60608152600061562660608301866155e3565b828103602084015261563881866155e3565b915050826040830152949350505050565b60006020828403121561565b57600080fd5b5051919050565b600067ffffffffffffffff80831681851680830382111561568557615685614cd3565b01949350505050565b600067ffffffffffffffff808316818103614d5157614d51614cd3565b6020815260006125f16020830184614e53565b604081526156d960408201845167ffffffffffffffff169052565b600060208401516156f6606084018267ffffffffffffffff169052565b506040840151608083015260608401516001600160a01b03811660a084015250608084015167ffffffffffffffff811660c08401525060a084015160e083015260c084015161010061574b8185018315159052565b60e08601519150610120615769818601846001600160a01b03169052565b8187015192506101809150610140828187015261578a6101c087018561412a565b93508188015191506101607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301818801526157c88584614e53565b94508189015192506157e4848801846001600160a01b03169052565b8801516101a0870152505050831515602084015290506121f9565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156158475780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
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
	Arg0 IEVM2EVMOffRampOffRampConfig
	Raw  types.Log
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
	return common.HexToHash("0x08fd01f552eb5d1aaca89b8c44ca4f2c0693ba70658968601e9c8de5714d8062")
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
