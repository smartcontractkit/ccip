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

type IEVM2EVMOffRampOffRampConfig struct {
	FeeManager                              common.Address
	PermissionLessExecutionThresholdSeconds uint32
	ExecutionDelaySeconds                   uint64
	MaxDataSize                             uint64
	MaxTokensLength                         uint64
}

type InternalEVM2EVMMessage struct {
	SourceChainId    uint64
	SequenceNumber   uint64
	FeeTokenAmount   *big.Int
	Sender           common.Address
	Nonce            uint64
	GasLimit         *big.Int
	Strict           bool
	Receiver         common.Address
	Data             []byte
	TokensAndAmounts []CommonEVMTokenAndAmount
	FeeToken         common.Address
	MessageId        [32]byte
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

var EVM2EVMOffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"contractIFeeManager\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"contractICommitStore\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"seqNum\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedGasPriceUpdate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ValueExceedsCapacity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIAFN\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"contractIFeeManager\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"destTokensAndAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCommon.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceFeeToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"feeTokenBaseUnitsPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structInternal.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"rep\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCommon.EVMTokenAndAmount[]\",\"name\":\"tokensAndAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitStore\",\"outputs\":[{\"internalType\":\"contractICommitStore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOffRampConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIFeeManager\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolByDestToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractIAny2EVMOffRampRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceFeeToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"destChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"feeTokenBaseUnitsPerUnitGas\",\"type\":\"uint128\"}],\"internalType\":\"structInternal.FeeUpdate[]\",\"name\":\"feeUpdates\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAFN\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICommitStore\",\"name\":\"commitStore\",\"type\":\"address\"}],\"name\":\"setCommitStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIFeeManager\",\"name\":\"feeManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structIEVM2EVMOffRamp.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setOffRampConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structIAggregateRateLimiter.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAny2EVMOffRampRouter\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162006514380380620065148339810160408190526200003591620007d5565b88888888888888888888888787878787878083838633806000806000806101000a81548160ff02191690831515021790555060006001600160a01b0316826001600160a01b031603620000cf5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b031990921691909117909155811615620001095762000109816200049e565b5050506001600160a01b0381166200013457604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03929092169190911790558051825114620001765760405162d8548360e71b815260040160405180910390fd5b81516200018b906005906020850190620005b5565b5060005b8251811015620003145760006040518060400160405280848481518110620001bb57620001bb62000941565b60200260200101516001600160a01b03168152602001836001600160601b031681525090508060036000868581518110620001fa57620001fa62000941565b6020908102919091018101516001600160a01b0390811683528282019390935260409091016000908120845194909201516001600160601b0316600160a01b02939092169290921790915581518451909160049186908690811062000263576200026362000941565b60200260200101516001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015620002a9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002cf919062000957565b6001600160a01b039081168252602082019290925260400160002080546001600160a01b03191692909116919091179055506200030c816200097e565b90506200018f565b50508151600680546001600160a01b0319166001600160a01b0392831617905560408051608081018252602080860151808352958301805191830182905251928201839052426060909201829052600995909555600a94909455600b55600c9290925550861662000398576040516342bcdf7f60e11b815260040160405180910390fd5b5050506001600160401b03948516608090815293851660a052506001600160a01b0391821660c052600e80549183166001600160a01b031990921691909117905589516017805460208d015160408e01518716600160c01b026001600160c01b0363ffffffff909216600160a01b026001600160c01b031990931694909516939093171791909116919091179055606089015160188054928b0151841668010000000000000000026001600160801b031990931691909316171790556200047f7fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a6200054f565b60e08181525050505050505050505050505050505050505050620009a6565b336001600160a01b03821603620004f85760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000c6565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008160805160a05160c0516040516020016200059894939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b8280548282559060005260206000209081019282156200060d579160200282015b828111156200060d57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620005d6565b506200061b9291506200061f565b5090565b5b808211156200061b576000815560010162000620565b80516001600160401b03811681146200064e57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b03811182821017156200068e576200068e62000653565b60405290565b6001600160a01b0381168114620006aa57600080fd5b50565b80516200064e8162000694565b600082601f830112620006cc57600080fd5b815160206001600160401b0380831115620006eb57620006eb62000653565b8260051b604051601f19603f8301168101818110848211171562000713576200071362000653565b6040529384528581018301938381019250878511156200073257600080fd5b83870191505b848210156200075e5781516200074e8162000694565b8352918301919083019062000738565b979650505050505050565b6000606082840312156200077c57600080fd5b604051606081016001600160401b0381118282101715620007a157620007a162000653565b80604052508091508251620007b68162000694565b8082525060208301516020820152604083015160408201525092915050565b6000806000806000806000806000898b036101e0811215620007f657600080fd5b620008018b62000636565b99506200081160208c0162000636565b985060a0603f19820112156200082657600080fd5b506200083162000669565b60408b0151620008418162000694565b815260608b015163ffffffff811681146200085b57600080fd5b60208201526200086e60808c0162000636565b60408201526200088160a08c0162000636565b60608201526200089460c08c0162000636565b60808201529650620008a960e08b01620006ad565b9550620008ba6101008b01620006ad565b9450620008cb6101208b01620006ad565b6101408b01519094506001600160401b0380821115620008ea57600080fd5b620008f88d838e01620006ba565b94506101608c01519150808211156200091057600080fd5b506200091f8c828d01620006ba565b925050620009328b6101808c0162000769565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156200096a57600080fd5b8151620009778162000694565b9392505050565b6000600182016200099f57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e051615b1162000a0360003960006132bf0152600081816124d3015261303d015260008181610322015261301c0152600081816102fd015281816124af01528181612ffb0152613e170152615b116000f3fe608060405234801561001057600080fd5b50600436106102e95760003560e01c806379ba509711610191578063b4069b31116100e3578063c903328411610097578063d7e2bb5011610071578063d7e2bb501461082e578063eb511dd41461085a578063f2fde38b1461086d57600080fd5b8063c903328414610800578063cb3c8c0114610813578063d3c7c2c71461082657600080fd5b8063b66f0efb116100c8578063b66f0efb146107c6578063c0d78655146107d7578063c5a1d7f0146107ea57600080fd5b8063b4069b31146107a0578063b5767166146107b357600080fd5b806390c2339b11610145578063afcb95d71161011f578063afcb95d71461075c578063b0f479a11461077c578063b1dc65a41461078d57600080fd5b806390c2339b146105f6578063a8ebd0f414610631578063abc39f1f1461074957600080fd5b80638456cb59116101765780638456cb591461058c578063856c8247146105945780638da5cb5b146105e057600080fd5b806379ba50971461055457806381ff70481461055c57600080fd5b80633015b91c1161024a578063599f6431116101fe578063666cab8d116101d8578063666cab8d14610517578063681fba161461052c578063744b92e21461054157600080fd5b8063599f6431146104cf5780635c975abb146104e05780635d86f141146104eb57600080fd5b80633f4ba83a1161022f5780633f4ba83a146104945780634352fa9f1461049c5780634741062e146104af57600080fd5b80633015b91c1461047357806339aa92641461048157600080fd5b806317332f9b116102a15780631a4b60dd116102865780631a4b60dd146104285780631ef381741461043b5780632222dd421461044e57600080fd5b806317332f9b146103cc578063181f5a77146103df57600080fd5b80631130ab7b116102d25780631130ab7b14610367578063142a98fc1461037a578063147809b3146103b457600080fd5b8063087ae6df146102ee578063108ee5fc14610352575b600080fd5b6040805167ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811682527f0000000000000000000000000000000000000000000000000000000000000000166020820152015b60405180910390f35b610365610360366004614159565b610880565b005b610365610375366004614611565b610937565b6103a7610388366004614663565b67ffffffffffffffff166000908152600f602052604090205460ff1690565b6040516103499190614696565b6103bc610945565b6040519015158152602001610349565b6103656103da3660046146be565b6109d2565b61041b6040518060400160405280601481526020017f45564d3245564d4f666652616d7020312e302e3000000000000000000000000081525081565b604051610349919061477c565b61036561043636600461478f565b610b0d565b6103656104493660046148a5565b610c6b565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610349565b6103656102e9366004614972565b61036561048f366004614159565b6112d1565b610365611308565b6103656104aa366004614a11565b61131a565b6104c26104bd366004614acc565b61156f565b6040516103499190614b09565b6006546001600160a01b031661045b565b60005460ff166103bc565b61045b6104f9366004614159565b6001600160a01b039081166000908152600360205260409020541690565b61051f611637565b6040516103499190614b91565b610534611699565b6040516103499190614ba4565b61036561054f366004614be5565b61175e565b610365611ad6565b6012546010546040805163ffffffff80851682526401000000009094049093166020840152820152606001610349565b610365611bb9565b6105c76105a2366004614159565b6001600160a01b031660009081526016602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff9091168152602001610349565b60005461010090046001600160a01b031661045b565b6105fe611bc9565b60405161034991908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b6106ed6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a0810182526017546001600160a01b038116825263ffffffff74010000000000000000000000000000000000000000820416602083015267ffffffffffffffff78010000000000000000000000000000000000000000000000009091048116928201929092526018548083166060830152680100000000000000009004909116608082015290565b6040805182516001600160a01b0316815260208084015163ffffffff16908201528282015167ffffffffffffffff908116928201929092526060808401518316908201526080928301519091169181019190915260a001610349565b610365610757366004614c9b565b611c6a565b604080516001815260006020820181905291810191909152606001610349565b600d546001600160a01b031661045b565b61036561079b366004614e2b565b611e22565b61045b6107ae366004614159565b612372565b6103656107c1366004614f10565b612460565b600e546001600160a01b031661045b565b6103656107e5366004614159565b61246c565b6107f261252a565b604051908152602001610349565b61036561080e366004614159565b61255a565b610365610821366004614f45565b612591565b61053461259c565b61045b61083c366004614159565b6001600160a01b039081166000908152600460205260409020541690565b610365610868366004614be5565b6125fc565b61036561087b366004614159565b612878565b610888612889565b6001600160a01b0381166108c8576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff1983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd2891015b60405180910390a15050565b61094182826128e8565b5050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa1580156109a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109cc9190614f85565b15905090565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b031614158015610a0f57506006546001600160a01b03163314155b15610a46576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602081015179ffffffffffffffffffffffffffffffffffffffffffffffffffff11610a9d576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610aa76009612af6565b6040810151600a8190556020820151600955600b54610ac69190612ba3565b600b556040818101516020808401518351928352908201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b610b15612889565b80516017805460208085018051604080880180516001600160a01b039098167fffffffffffffffff00000000000000000000000000000000000000000000000090961686177401000000000000000000000000000000000000000063ffffffff948516021777ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff998a16021790965560608089018051601880546080808e018051948e167fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169290921768010000000000000000948e1694909402939093179091558451988952955190941695870195909552955187169085015291518516938301939093529151909216908201527f688cd64fc184dc976a5b5dda59c14123b4526a25ed379b766a8570f40b0257099060a001610b02565b855185518560ff16601f831115610ce3576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f746f6f206d616e79207369676e6572730000000000000000000000000000000060448201526064015b60405180910390fd5b80600003610d4d576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610cda565b818314610ddb576040517f89a61989000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401610cda565b610de6816003614fb8565b8311610e4e576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401610cda565b610e56612889565b60145460005b81811015610efe576013600060148381548110610e7b57610e7b614fd7565b60009182526020808320909101546001600160a01b031683528201929092526040018120805461ffff1916905560158054601392919084908110610ec157610ec1614fd7565b60009182526020808320909101546001600160a01b031683528201929092526040019020805461ffff19169055610ef781614fed565b9050610e5c565b50895160005b818110156111935760008c8281518110610f2057610f20614fd7565b6020026020010151905060006002811115610f3d57610f3d614680565b6001600160a01b038216600090815260136020526040902054610100900460ff166002811115610f6f57610f6f614680565b14610fd6576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401610cda565b6040805180820190915260ff8316815260208101600190526001600160a01b03821660009081526013602090815260409091208251815460ff90911660ff19821681178355928401519192839161ffff19161761010083600281111561103e5761103e614680565b021790555090505060008c838151811061105a5761105a614fd7565b602002602001015190506000600281111561107757611077614680565b6001600160a01b038216600090815260136020526040902054610100900460ff1660028111156110a9576110a9614680565b14611110576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610cda565b6040805180820190915260ff8416815260208101600290526001600160a01b03821660009081526013602090815260409091208251815460ff90911660ff19821681178355928401519192839161ffff19161761010083600281111561117857611178614680565b021790555090505050508061118c90614fed565b9050610f04565b508a516111a79060149060208e01906140a2565b5089516111bb9060159060208d01906140a2565b506011805460ff8381166101000261ffff19909216908c1617179055601280546112249146913091906000906111f69063ffffffff16615007565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168e8e8e8e8e8e612bbd565b6010600001819055506000601260049054906101000a900463ffffffff16905043601260046101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581601060000154601260009054906101000a900463ffffffff168f8f8f8f8f8f6040516112bb9998979695949392919061502a565b60405180910390a1505050505050505050505050565b6112d9612889565b6006805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b611310612889565b611318612c4a565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561135757506006546001600160a01b03163314155b1561138e576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151815181146113ca576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085460005b818110156114245760076000600883815481106113ef576113ef614fd7565b60009182526020808320909101546001600160a01b0316835282019290925260400181205561141d81614fed565b90506113d0565b5060005b8281101561155457600085828151811061144457611444614fd7565b6020026020010151905060006001600160a01b0316816001600160a01b03160361149a576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8482815181106114ac576114ac614fd7565b602002602001015160076000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f8186848151811061151157611511614fd7565b602002602001015160405161153b9291906001600160a01b03929092168252602082015260400190565b60405180910390a15061154d81614fed565b9050611428565b5083516115689060089060208701906140a2565b5050505050565b80516060908067ffffffffffffffff81111561158d5761158d614176565b6040519080825280602002602001820160405280156115b6578160200160208202803683370190505b50915060005b8181101561163057600760008583815181106115da576115da614fd7565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000205483828151811061161557611615614fd7565b602090810291909101015261162981614fed565b90506115bc565b5050919050565b6060601580548060200260200160405190810160405280929190818152602001828054801561168f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611671575b5050505050905090565b60055460609067ffffffffffffffff8111156116b7576116b7614176565b6040519080825280602002602001820160405280156116e0578160200160208202803683370190505b50905060005b60055481101561175a576117206005828154811061170657611706614fd7565b6000918252602090912001546001600160a01b0316612372565b82828151811061173257611732614fd7565b6001600160a01b039092166020928302919091019091015261175381614fed565b90506116e6565b5090565b611766612889565b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff1690820152906117f4576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816001600160a01b031681600001516001600160a01b031614611843576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058054906000906118566001846150c0565b8154811061186657611866614fd7565b9060005260206000200160009054906101000a90046001600160a01b03169050600583602001516bffffffffffffffffffffffff16815481106118ab576118ab614fd7565b6000918252602090912001546001600160a01b031660056118cd6001856150c0565b815481106118dd576118dd614fd7565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600584602001516bffffffffffffffffffffffff168154811061193157611931614fd7565b6000918252602080832090910180546001600160a01b0394851673ffffffffffffffffffffffffffffffffffffffff199091161790558581015184841683526003909152604090912080546bffffffffffffffffffffffff90921674010000000000000000000000000000000000000000029190921617905560058054806119bb576119bb6150d7565b6001900381819060005260206000200160006101000a8154906001600160a01b030219169055905560046000856001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a25573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a4991906150ed565b6001600160a01b03908116825260208083019390935260409182016000908120805473ffffffffffffffffffffffffffffffffffffffff1916905588821680825260038552838220919091558251908152908716928101929092527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b03163314611b305760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610cda565b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff84161784556001805473ffffffffffffffffffffffffffffffffffffffff191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b611bc1612889565b611318612ce6565b611bf46040518060800160405280600081526020016000815260200160008152602001600081525090565b604080516080810182526009548152600a546020820152600b5491810191909152600c54606082018190524290600090611c2e90836150c0565b60208401518451919250611c5a91611c469084614fb8565b8560400151611c55919061510a565b612ba3565b6040840152506060820152919050565b333014611ca3576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160008082526020820190925281611ce0565b6040805180820190915260008082526020820152815260200190600190039081611cb95790505b506101208401515190915015611d0657611d038361012001518460e00151612d6e565b90505b60e08301516001600160a01b03163b1580611d56575060e0830151611d54906001600160a01b03167f3015b91c00000000000000000000000000000000000000000000000000000000612f19565b155b15611d6057505050565b600d546001600160a01b0316624b61bb611d7a8584612f35565b848660a001518760e001516040518563ffffffff1660e01b8152600401611da49493929190615166565b6020604051808303816000875af1158015611dc3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611de79190614f85565b611e1d576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b611e6187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612fd892505050565b6040805160608101825260105480825260115460ff808216602085015261010090910416928201929092528935918214611ed45780516040517f93df584c000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610cda565b6040805183815260208c81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a1600281602001518260400151611f2f9190615218565b611f399190615253565b611f44906001615218565b60ff168614611f7f576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b858414611fb8576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526013602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115611ffb57611ffb614680565b600281111561200c5761200c614680565b905250905060028160200151600281111561202957612029614680565b14801561206357506015816000015160ff168154811061204b5761204b614fd7565b6000918252602090912001546001600160a01b031633145b612099576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060006120a7856020614fb8565b6120b2886020614fb8565b6120be8b61014461510a565b6120c8919061510a565b6120d2919061510a565b9050368114612116576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401610cda565b5060008989604051612129929190615275565b604051908190038120612140918d90602001615285565b604051602081830303815290604052805190602001209050612160614110565b8760005b8181101561236257600060018589846020811061218357612183614fd7565b61219091901a601b615218565b8e8e868181106121a2576121a2614fd7565b905060200201358d8d878181106121bb576121bb614fd7565b90506020020135604051600081526020016040526040516121f8949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa15801561221a573d6000803e3d6000fd5b505060408051601f198101516001600160a01b038116600090815260136020908152848220848601909552845460ff808216865293975091955092939284019161010090910416600281111561227257612272614680565b600281111561228357612283614680565b90525090506001816020015160028111156122a0576122a0614680565b146122d7576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051859060ff16601f81106122ee576122ee614fd7565b60200201511561232a576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600185826000015160ff16601f811061234557612345614fd7565b911515602090920201525061235b905081614fed565b9050612164565b5050505050505050505050505050565b6001600160a01b03808216600090815260036020526040812054909116806123c6576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015612435573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061245991906150ed565b9392505050565b61246981612fd8565b50565b612474612889565b600d805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b038381169182179092556040805167ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681527f0000000000000000000000000000000000000000000000000000000000000000909316602084015290917f052b5907be1d3ac35d571862117562e80ee743c01251e388dafb7dc4e92a726c910160405180910390a250565b60006125557fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a612ff6565b905090565b612562612889565b600e805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6124698160016128e8565b6060600580548060200260200160405190810160405280929190818152602001828054801561168f576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611671575050505050905090565b612604612889565b6001600160a01b038216158061262157506001600160a01b038116155b15612658576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352740100000000000000000000000000000000000000009093046bffffffffffffffffffffffff169082015290156126e7576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038083168083526005546bffffffffffffffffffffffff90811660208086019182528785166000908152600382526040808220885194519095167401000000000000000000000000000000000000000002939096169290921790925583517f21df0da70000000000000000000000000000000000000000000000000000000081529351869460049492936321df0da79282870192819003870181865afa15801561279c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127c091906150ed565b6001600160a01b03908116825260208083019390935260409182016000908120805495831673ffffffffffffffffffffffffffffffffffffffff199687161790556005805460018101825591527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001805488831695168517905581519384528516918301919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505050565b612880612889565b612469816130b6565b60005461010090046001600160a01b031633146113185760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610cda565b60005460ff161561293b5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610cda565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561298e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129b29190614f85565b156129e8576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600d546001600160a01b0316612a2a576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208201515115612aec578015612a6d576040517f198753d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60175460208301516040517f9086658e0000000000000000000000000000000000000000000000000000000081526001600160a01b0390921691639086658e91612ab9916004016152a1565b600060405180830381600087803b158015612ad357600080fd5b505af1158015612ae7573d6000803e3d6000fd5b505050505b6109418282613172565b6001810154600282015442911480612b115750808260030154145b15612b1a575050565b816001015482600201541115612b5c576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826003015482612b6e91906150c0565b60018401548454919250612b9591612b869084614fb8565b8560020154611c55919061510a565b600284015550600390910155565b6000818310612bb25781612bb4565b825b90505b92915050565b6000808a8a8a8a8a8a8a8a8a604051602001612be19998979695949392919061531f565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b60005460ff16612c9c5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610cda565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b60005460ff1615612d395760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610cda565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612cc93390565b60606000835167ffffffffffffffff811115612d8c57612d8c614176565b604051908082528060200260200182016040528015612dd157816020015b6040805180820190915260008082526020820152815260200190600190039081612daa5790505b50905060005b8451811015612f0f576000612e08868381518110612df757612df7614fd7565b602002602001015160000151613800565b9050612e3281878481518110612e2057612e20614fd7565b60200260200101516020015187613862565b806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e9491906150ed565b838381518110612ea657612ea6614fd7565b60209081029190910101516001600160a01b0390911690528551869083908110612ed257612ed2614fd7565b602002602001015160200151838381518110612ef057612ef0614fd7565b602090810291909101810151015250612f0881614fed565b9050612dd7565b50612bb4816138e2565b6000612f2483613ae6565b8015612bb45750612bb48383613b4a565b612f6a6040518060800160405280600067ffffffffffffffff1681526020016060815260200160608152602001606081525090565b6040518060800160405280846000015167ffffffffffffffff1681526020018460600151604051602001612fad91906001600160a01b0391909116815260200190565b6040516020818303038152906040528152602001846101000151815260200183815250905092915050565b61246981806020019051810190612fef91906155e3565b60006128e8565b6000817f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000604051602001613099949392919093845267ffffffffffffffff9283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b336001600160a01b0382160361310e5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610cda565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b604082015151600081900361318657505050565b60008167ffffffffffffffff8111156131a1576131a1614176565b6040519080825280602002602001820160405280156131ca578160200160208202803683370190505b50905060008267ffffffffffffffff8111156131e8576131e8614176565b60405190808252806020026020018201604052801561327657816020015b60408051610180810182526000808252602080830182905292820181905260608083018290526080830182905260a0830182905260c0830182905260e083018290526101008301819052610120830152610140820181905261016082015282526000199092019101816132065790505b50905060005b838110156133315760008660400151828151811061329c5761329c614fd7565b60200260200101518060200190518101906132b79190615751565b90506132e3817f0000000000000000000000000000000000000000000000000000000000000000613c19565b8483815181106132f5576132f5614fd7565b6020026020010181815250508083838151811061331457613314614fd7565b6020026020010181905250508061332a90614fed565b905061327c565b5060006133478387606001518860800151613d23565b5060175490915060009074010000000000000000000000000000000000000000900463ffffffff1661337983426150c0565b11905060005b858110156137f657600084828151811061339b5761339b614fd7565b6020026020010151905060006133ce826020015167ffffffffffffffff166000908152600f602052604090205460ff1690565b905060008160038111156133e4576133e4614680565b1480613401575060038160038111156133ff576133ff614680565b145b6134495760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610cda565b88156134a757838061346c5750600381600381111561346a5761346a614680565b145b6134a2576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613504565b60008160038111156134bb576134bb614680565b146135045760208201516040517f67d9ba0f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610cda565b600081600381111561351857613518614680565b036135b757608082015160608301516001600160a01b031660009081526016602052604090205467ffffffffffffffff9182169161355891166001615884565b67ffffffffffffffff16146135b75781606001516001600160a01b0316826080015167ffffffffffffffff167fd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf4123760405160405180910390a350506137e6565b6135c082613e15565b60208281015167ffffffffffffffff166000908152600f90915260408120805460ff191660011790556135f3838b613f6a565b60208085015167ffffffffffffffff166000908152600f909152604090208054919250829160ff1916600183600381111561363057613630614680565b02179055508915613716578260c00151801561365d5750600382600381111561365b5761365b614680565b145b801561367a5750600281600381111561367857613678614680565b145b806136b25750600082600381111561369457613694614680565b1480156136b2575060028160038111156136b0576136b0614680565b145b156137115760608301516001600160a01b03166000908152601660205260408120805467ffffffffffffffff16916136e9836158b0565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b613796565b8260c0015180156137385750600381600381111561373657613736614680565b145b6137965760608301516001600160a01b03166000908152601660205260408120805467ffffffffffffffff169161376e836158b0565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b826101600151836020015167ffffffffffffffff167f5e04b4755a4460aa6de4f3a906c4324a025c7449c02b52f5466659b5bfdfba5f836040516137da9190614696565b60405180910390a35050505b6137ef81614fed565b905061337f565b5050505050505050565b6001600160a01b03818116600090815260036020526040902054168061385d576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610cda565b919050565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401600060405180830381600087803b1580156138c557600080fd5b505af11580156138d9573d6000803e3d6000fd5b50505050505050565b6000805b82518110156139e15760006007600085848151811061390757613907614fd7565b6020026020010151600001516001600160a01b03166001600160a01b031681526020019081526020016000205490508060000361399a5783828151811061395057613950614fd7565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610cda565b8382815181106139ac576139ac614fd7565b602002602001015160200151816139c39190614fb8565b6139cd908461510a565b925050806139da90614fed565b90506138e6565b508015610941576139f26009612af6565b600a54811115613a3c57600a546040517f688ccf77000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610cda565b600b54811115613a9c57600954600b5460009190613a5a90846150c0565b613a6491906158cd565b9050806040517fe31e0f32000000000000000000000000000000000000000000000000000000008152600401610cda91815260200190565b8060096002016000828254613ab191906150c0565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba3789060200161092b565b6000613b12827f01ffc9a700000000000000000000000000000000000000000000000000000000613b4a565b8015612bb75750613b43827fffffffff00000000000000000000000000000000000000000000000000000000613b4a565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015613c02575060208210155b8015613c0e5750600081115b979650505050505050565b60008060001b828460200151856080015186606001518760e0015188610100015180519060200120896101200151604051602001613c5791906158e1565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d60400151604051602001613d059c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600e546040517f320488750000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b0390911690633204887590613d7c908a908a908a90600401615924565b6020604051808303816000875af1158015613d9b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613dbf919061595a565b905060008111613dfb576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a613e0790846150c0565b935093505050935093915050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16816000015167ffffffffffffffff1614613e955780516040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610cda565b601854610120820151516801000000000000000090910467ffffffffffffffff161015613f005760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610cda565b6018546101008201515167ffffffffffffffff909116101561246957601854610100820151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301526024820152604401610cda565b6040517fabc39f1f000000000000000000000000000000000000000000000000000000008152600090309063abc39f1f90613fab9086908690600401615973565b600060405180830381600087803b158015613fc557600080fd5b505af1925050508015613fd6575060015b614099573d808015614004576040519150601f19603f3d011682016040523d82523d6000602084013e614009565b606091505b5061401381615ab4565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003614065576003915050612bb7565b806040517fcf19edfd000000000000000000000000000000000000000000000000000000008152600401610cda919061477c565b50600292915050565b828054828255906000526020600020908101928215614104579160200282015b82811115614104578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039091161782556020909201916001909101906140c2565b5061175a92915061412f565b604051806103e00160405280601f906020820280368337509192915050565b5b8082111561175a5760008155600101614130565b6001600160a01b038116811461246957600080fd5b60006020828403121561416b57600080fd5b8135612bb481614144565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156141af576141af614176565b60405290565b60405160a0810167ffffffffffffffff811182821017156141af576141af614176565b6040805190810167ffffffffffffffff811182821017156141af576141af614176565b604051610180810167ffffffffffffffff811182821017156141af576141af614176565b604051601f8201601f1916810167ffffffffffffffff8111828210171561424857614248614176565b604052919050565b600067ffffffffffffffff82111561426a5761426a614176565b5060051b60200190565b67ffffffffffffffff8116811461246957600080fd5b803561385d81614274565b600082601f8301126142a657600080fd5b813560206142bb6142b683614250565b61421f565b82815260059290921b840181019181810190868411156142da57600080fd5b8286015b848110156142fe5780356142f181614274565b83529183019183016142de565b509695505050505050565b803561385d81614144565b6fffffffffffffffffffffffffffffffff8116811461246957600080fd5b600082601f83011261434357600080fd5b813560206143536142b683614250565b8281526060928302850182019282820191908785111561437257600080fd5b8387015b858110156143d45781818a03121561438e5760008081fd5b61439661418c565b81356143a181614144565b8152818601356143b081614274565b818701526040828101356143c381614314565b908201528452928401928101614376565b5090979650505050505050565b600067ffffffffffffffff8211156143fb576143fb614176565b50601f01601f191660200190565b600082601f83011261441a57600080fd5b81356144286142b6826143e1565b81815284602083860101111561443d57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261446b57600080fd5b8135602061447b6142b683614250565b82815260059290921b8401810191818101908684111561449a57600080fd5b8286015b848110156142fe57803567ffffffffffffffff8111156144be5760008081fd5b6144cc8986838b0101614409565b84525091830191830161449e565b600082601f8301126144eb57600080fd5b813560206144fb6142b683614250565b82815260059290921b8401810191818101908684111561451a57600080fd5b8286015b848110156142fe578035835291830191830161451e565b600060a0828403121561454757600080fd5b61454f6141b5565b9050813567ffffffffffffffff8082111561456957600080fd5b61457585838601614295565b8352602084013591508082111561458b57600080fd5b61459785838601614332565b602084015260408401359150808211156145b057600080fd5b6145bc8583860161445a565b604084015260608401359150808211156145d557600080fd5b506145e2848285016144da565b6060830152506080820135608082015292915050565b801515811461246957600080fd5b803561385d816145f8565b6000806040838503121561462457600080fd5b823567ffffffffffffffff81111561463b57600080fd5b61464785828601614535565b9250506020830135614658816145f8565b809150509250929050565b60006020828403121561467557600080fd5b8135612bb481614274565b634e487b7160e01b600052602160045260246000fd5b60208101600483106146b857634e487b7160e01b600052602160045260246000fd5b91905290565b6000606082840312156146d057600080fd5b6040516060810181811067ffffffffffffffff821117156146f3576146f3614176565b604052823561470181614144565b8152602083810135908201526040928301359281019290925250919050565b60005b8381101561473b578181015183820152602001614723565b8381111561474a576000848401525b50505050565b60008151808452614768816020860160208601614720565b601f01601f19169290920160200192915050565b602081526000612bb46020830184614750565b600060a082840312156147a157600080fd5b60405160a0810181811067ffffffffffffffff821117156147c4576147c4614176565b60405282356147d281614144565b8152602083013563ffffffff811681146147eb57600080fd5b602082015260408301356147fe81614274565b6040820152606083013561481181614274565b6060820152608083013561482481614274565b60808201529392505050565b600082601f83011261484157600080fd5b813560206148516142b683614250565b82815260059290921b8401810191818101908684111561487057600080fd5b8286015b848110156142fe57803561488781614144565b8352918301918301614874565b803560ff8116811461385d57600080fd5b60008060008060008060c087890312156148be57600080fd5b863567ffffffffffffffff808211156148d657600080fd5b6148e28a838b01614830565b975060208901359150808211156148f857600080fd5b6149048a838b01614830565b965061491260408a01614894565b9550606089013591508082111561492857600080fd5b6149348a838b01614409565b945061494260808a0161428a565b935060a089013591508082111561495857600080fd5b5061496589828a01614409565b9150509295509295509295565b60006020828403121561498457600080fd5b813567ffffffffffffffff81111561499b57600080fd5b820160808185031215612bb457600080fd5b600082601f8301126149be57600080fd5b813560206149ce6142b683614250565b82815260059290921b840181019181810190868411156149ed57600080fd5b8286015b848110156142fe578035614a0481614144565b83529183019183016149f1565b60008060408385031215614a2457600080fd5b823567ffffffffffffffff80821115614a3c57600080fd5b614a48868387016149ad565b9350602091508185013581811115614a5f57600080fd5b85019050601f81018613614a7257600080fd5b8035614a806142b682614250565b81815260059190911b82018301908381019088831115614a9f57600080fd5b928401925b82841015614abd57833582529284019290840190614aa4565b80955050505050509250929050565b600060208284031215614ade57600080fd5b813567ffffffffffffffff811115614af557600080fd5b614b01848285016149ad565b949350505050565b6020808252825182820181905260009190848201906040850190845b81811015614b4157835183529284019291840191600101614b25565b50909695505050505050565b600081518084526020808501945080840160005b83811015614b865781516001600160a01b031687529582019590820190600101614b61565b509495945050505050565b602081526000612bb46020830184614b4d565b6020808252825182820181905260009190848201906040850190845b81811015614b415783516001600160a01b031683529284019291840191600101614bc0565b60008060408385031215614bf857600080fd5b8235614c0381614144565b9150602083013561465881614144565b600082601f830112614c2457600080fd5b81356020614c346142b683614250565b82815260069290921b84018101918181019086841115614c5357600080fd5b8286015b848110156142fe5760408189031215614c705760008081fd5b614c786141d8565b8135614c8381614144565b81528185013585820152835291830191604001614c57565b60008060408385031215614cae57600080fd5b823567ffffffffffffffff80821115614cc657600080fd5b908401906101808287031215614cdb57600080fd5b614ce36141fb565b614cec8361428a565b8152614cfa6020840161428a565b602082015260408301356040820152614d1560608401614309565b6060820152614d266080840161428a565b608082015260a083013560a0820152614d4160c08401614606565b60c0820152614d5260e08401614309565b60e08201526101008084013583811115614d6b57600080fd5b614d7789828701614409565b8284015250506101208084013583811115614d9157600080fd5b614d9d89828701614c13565b8284015250506101409150614db3828401614309565b8282015261016091508183013582820152809450505050614dd660208401614606565b90509250929050565b60008083601f840112614df157600080fd5b50813567ffffffffffffffff811115614e0957600080fd5b6020830191508360208260051b8501011115614e2457600080fd5b9250929050565b60008060008060008060008060e0898b031215614e4757600080fd5b606089018a811115614e5857600080fd5b8998503567ffffffffffffffff80821115614e7257600080fd5b818b0191508b601f830112614e8657600080fd5b813581811115614e9557600080fd5b8c6020828501011115614ea757600080fd5b6020830199508098505060808b0135915080821115614ec557600080fd5b614ed18c838d01614ddf565b909750955060a08b0135915080821115614eea57600080fd5b50614ef78b828c01614ddf565b999c989b50969995989497949560c00135949350505050565b600060208284031215614f2257600080fd5b813567ffffffffffffffff811115614f3957600080fd5b614b0184828501614409565b600060208284031215614f5757600080fd5b813567ffffffffffffffff811115614f6e57600080fd5b614b0184828501614535565b805161385d816145f8565b600060208284031215614f9757600080fd5b8151612bb4816145f8565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615614fd257614fd2614fa2565b500290565b634e487b7160e01b600052603260045260246000fd5b6000600019820361500057615000614fa2565b5060010190565b600063ffffffff80831681810361502057615020614fa2565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b1660408501525080606084015261505a8184018a614b4d565b9050828103608084015261506e8189614b4d565b905060ff871660a084015282810360c084015261508b8187614750565b905067ffffffffffffffff851660e08401528281036101008401526150b08185614750565b9c9b505050505050505050505050565b6000828210156150d2576150d2614fa2565b500390565b634e487b7160e01b600052603160045260246000fd5b6000602082840312156150ff57600080fd5b8151612bb481614144565b6000821982111561511d5761511d614fa2565b500190565b600081518084526020808501945080840160005b83811015614b8657815180516001600160a01b031688528301518388015260409096019590820190600101615136565b6080815267ffffffffffffffff855116608082015260006020860151608060a0840152615197610100840182614750565b905060408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80808584030160c08601526151d38383614750565b925060608901519150808584030160e0860152506151f18282615122565b961515602085015250505060408101929092526001600160a01b0316606090910152919050565b600060ff821660ff84168060ff0382111561523557615235614fa2565b019392505050565b634e487b7160e01b600052601260045260246000fd5b600060ff8316806152665761526661523d565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b602080825282518282018190526000919060409081850190868401855b8281101561531257815180516001600160a01b031685528681015167ffffffffffffffff16878601528501516fffffffffffffffffffffffffffffffff1685850152606090930192908501906001016152be565b5091979650505050505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526153598285018b614b4d565b9150838203608085015261536d828a614b4d565b915060ff881660a085015283820360c085015261538a8288614750565b90861660e085015283810361010085015290506150b08185614750565b805161385d81614274565b600082601f8301126153c357600080fd5b815160206153d36142b683614250565b82815260059290921b840181019181810190868411156153f257600080fd5b8286015b848110156142fe57805161540981614274565b83529183019183016153f6565b805161385d81614144565b600082601f83011261543257600080fd5b815160206154426142b683614250565b8281526060928302850182019282820191908785111561546157600080fd5b8387015b858110156143d45781818a03121561547d5760008081fd5b61548561418c565b815161549081614144565b81528186015161549f81614274565b818701526040828101516154b281614314565b908201528452928401928101615465565b600082601f8301126154d457600080fd5b81516154e26142b6826143e1565b8181528460208386010111156154f757600080fd5b614b01826020830160208701614720565b600082601f83011261551957600080fd5b815160206155296142b683614250565b82815260059290921b8401810191818101908684111561554857600080fd5b8286015b848110156142fe57805167ffffffffffffffff81111561556c5760008081fd5b61557a8986838b01016154c3565b84525091830191830161554c565b600082601f83011261559957600080fd5b815160206155a96142b683614250565b82815260059290921b840181019181810190868411156155c857600080fd5b8286015b848110156142fe57805183529183019183016155cc565b6000602082840312156155f557600080fd5b815167ffffffffffffffff8082111561560d57600080fd5b9083019060a0828603121561562157600080fd5b6156296141b5565b82518281111561563857600080fd5b615644878286016153b2565b82525060208301518281111561565957600080fd5b61566587828601615421565b60208301525060408301518281111561567d57600080fd5b61568987828601615508565b6040830152506060830151828111156156a157600080fd5b6156ad87828601615588565b6060830152506080830151608082015280935050505092915050565b600082601f8301126156da57600080fd5b815160206156ea6142b683614250565b82815260069290921b8401810191818101908684111561570957600080fd5b8286015b848110156142fe57604081890312156157265760008081fd5b61572e6141d8565b815161573981614144565b8152818501518582015283529183019160400161570d565b60006020828403121561576357600080fd5b815167ffffffffffffffff8082111561577b57600080fd5b90830190610180828603121561579057600080fd5b6157986141fb565b6157a1836153a7565b81526157af602084016153a7565b6020820152604083015160408201526157ca60608401615416565b60608201526157db608084016153a7565b608082015260a083015160a08201526157f660c08401614f7a565b60c082015261580760e08401615416565b60e0820152610100808401518381111561582057600080fd5b61582c888287016154c3565b828401525050610120808401518381111561584657600080fd5b615852888287016156c9565b8284015250506101409150615868828401615416565b9181019190915261016091820151918101919091529392505050565b600067ffffffffffffffff8083168185168083038211156158a7576158a7614fa2565b01949350505050565b600067ffffffffffffffff80831681810361502057615020614fa2565b6000826158dc576158dc61523d565b500490565b602081526000612bb46020830184615122565b600081518084526020808501945080840160005b83811015614b8657815187529582019590820190600101615908565b60608152600061593760608301866158f4565b828103602084015261594981866158f4565b915050826040830152949350505050565b60006020828403121561596c57600080fd5b5051919050565b6040815261598e60408201845167ffffffffffffffff169052565b600060208401516159ab606084018267ffffffffffffffff169052565b506040840151608083015260608401516001600160a01b03811660a084015250608084015167ffffffffffffffff811660c08401525060a084015160e083015260c0840151610100615a008185018315159052565b60e08601519150610120615a1e818601846001600160a01b03169052565b81870151925061018091506101408281870152615a3f6101c0870185614750565b93508188015191506101607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030181880152615a7d8584615122565b9450818901519250615a99848801846001600160a01b03169052565b8801516101a087015250505083151560208401529050612459565b6000815160208301517fffffffff0000000000000000000000000000000000000000000000000000000080821693506004831015615afc5780818460040360031b1b83161693505b50505091905056fea164736f6c634300080f000a",
}

var EVM2EVMOffRampHelperABI = EVM2EVMOffRampHelperMetaData.ABI

var EVM2EVMOffRampHelperBin = EVM2EVMOffRampHelperMetaData.Bin

func DeployEVM2EVMOffRampHelper(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId uint64, chainId uint64, offRampConfig IEVM2EVMOffRampOffRampConfig, onRampAddress common.Address, commitStore common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig IAggregateRateLimiterRateLimiterConfig) (common.Address, *types.Transaction, *EVM2EVMOffRampHelper, error) {
	parsed, err := EVM2EVMOffRampHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMOffRampHelperBin), backend, sourceChainId, chainId, offRampConfig, onRampAddress, commitStore, afn, sourceTokens, pools, rateLimiterConfig)
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) CcipReceive(opts *bind.CallOpts, arg0 CommonAny2EVMMessage) error {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) CcipReceive(arg0 CommonAny2EVMMessage) error {
	return _EVM2EVMOffRampHelper.Contract.CcipReceive(&_EVM2EVMOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) CcipReceive(arg0 CommonAny2EVMMessage) error {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetCommitStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getCommitStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetCommitStore() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetCommitStore(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetCommitStore() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetCommitStore(&_EVM2EVMOffRampHelper.CallOpts)
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetRouter() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetRouter(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMOffRampHelper.Contract.GetRouter(&_EVM2EVMOffRampHelper.CallOpts)
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetCommitStore(opts *bind.TransactOpts, commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setCommitStore", commitStore)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetCommitStore(commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetCommitStore(&_EVM2EVMOffRampHelper.TransactOpts, commitStore)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetCommitStore(commitStore common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetCommitStore(&_EVM2EVMOffRampHelper.TransactOpts, commitStore)
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetRouter(&_EVM2EVMOffRampHelper.TransactOpts, router)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetRouter(&_EVM2EVMOffRampHelper.TransactOpts, router)
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
	Arg0 IEVM2EVMOffRampOffRampConfig
	Raw  types.Log
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

type EVM2EVMOffRampHelperOffRampRouterSetIterator struct {
	Event *EVM2EVMOffRampHelperOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperOffRampRouterSet)
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
		it.Event = new(EVM2EVMOffRampHelperOffRampRouterSet)
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

func (it *EVM2EVMOffRampHelperOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperOffRampRouterSet struct {
	Router        common.Address
	SourceChainId uint64
	OnRampAddress common.Address
	Raw           types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMOffRampHelperOffRampRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperOffRampRouterSetIterator{contract: _EVM2EVMOffRampHelper.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOffRampRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperOffRampRouterSet)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseOffRampRouterSet(log types.Log) (*EVM2EVMOffRampHelperOffRampRouterSet, error) {
	event := new(EVM2EVMOffRampHelperOffRampRouterSet)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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
	case _EVM2EVMOffRampHelper.abi.Events["OffRampRouterSet"].ID:
		return _EVM2EVMOffRampHelper.ParseOffRampRouterSet(log)
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
	return common.HexToHash("0x688cd64fc184dc976a5b5dda59c14123b4526a25ed379b766a8570f40b025709")
}

func (EVM2EVMOffRampHelperOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x052b5907be1d3ac35d571862117562e80ee743c01251e388dafb7dc4e92a726c")
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

	CcipReceive(opts *bind.CallOpts, arg0 CommonAny2EVMMessage) error

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

		error)

	GetCommitStore(opts *bind.CallOpts) (common.Address, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetOffRampConfig(opts *bind.CallOpts) (IEVM2EVMOffRampOffRampConfig, error)

	GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error)

	GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

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

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetCommitStore(opts *bind.TransactOpts, commitStore common.Address) (*types.Transaction, error)

	SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetOffRampConfig(opts *bind.TransactOpts, config IEVM2EVMOffRampOffRampConfig) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config IAggregateRateLimiterRateLimiterConfig) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

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

	FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMOffRampHelperOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOffRampRouterSet, router []common.Address) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*EVM2EVMOffRampHelperOffRampRouterSet, error)

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
