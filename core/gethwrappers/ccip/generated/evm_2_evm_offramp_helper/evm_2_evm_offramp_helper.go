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
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
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
	_ = abi.ConvertType
)

type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type EVM2EVMOffRampDynamicConfig struct {
	PermissionLessExecutionThresholdSeconds uint32
	Router                                  common.Address
	PriceRegistry                           common.Address
	MaxTokensLength                         uint16
	MaxDataSize                             uint32
	MaxPoolGas                              uint32
}

type EVM2EVMOffRampStaticConfig struct {
	CommitStore         common.Address
	ChainSelector       uint64
	SourceChainSelector uint64
	OnRamp              common.Address
	PrevOffRamp         common.Address
	ArmProxy            common.Address
}

type InternalEVM2EVMMessage struct {
	SourceChainSelector uint64
	Sender              common.Address
	Receiver            common.Address
	SequenceNumber      uint64
	GasLimit            *big.Int
	Strict              bool
	Nonce               uint64
	FeeToken            common.Address
	FeeTokenAmount      *big.Int
	Data                []byte
	TokenAmounts        []ClientEVMTokenAmount
	SourceTokenData     [][]byte
	MessageId           [32]byte
}

type InternalExecutionReport struct {
	Messages          []InternalEVM2EVMMessage
	OffchainTokenData [][][]byte
	Proofs            [][32]byte
	ProofFlagBits     *big.Int
}

type InternalPoolUpdate struct {
	Token common.Address
	Pool  common.Address
}

type RateLimiterConfig struct {
	IsEnabled bool
	Capacity  *big.Int
	Rate      *big.Int
}

type RateLimiterTokenBucket struct {
	Tokens      *big.Int
	LastUpdated uint32
	IsEnabled   bool
	Capacity    *big.Int
	Rate        *big.Int
}

var EVM2EVMOffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadARMSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitStoreAlreadyInUse\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"InvalidManualExecutionGasLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMessageId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"InvalidNewState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionGasLimitMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"TokenDataMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"TokenHandlingError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"TokenRateLimitError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedTokenData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolGas\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedSenderWithPreviousRampMessageInflight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"IS_SCRIPT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"}],\"name\":\"applyPoolUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"rep\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"manualExecGasLimits\",\"type\":\"uint256[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"destTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolGas\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOffRamp.DynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"bitmapIndex\",\"type\":\"uint64\"}],\"name\":\"getExecutionStateBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolByDestToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"armProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"gasLimitOverrides\",\"type\":\"uint256[]\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"originalSender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"}],\"name\":\"releaseOrMintTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"setExecutionStateHelper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"}],\"name\":\"trialExecute\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101806040526014805460ff191660011790553480156200001f57600080fd5b5060405162006fac38038062006fac8339810160408190526200004291620008a6565b838383838033806000816200009e5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000d157620000d1816200048d565b50506040805160a081018252602084810180516001600160801b039081168085524263ffffffff169385018490528751151585870181905292518216606086018190529790950151166080938401819052600380546001600160a01b031916909517600160801b9384021760ff60a01b1916600160a01b90920291909117909355909102909217600455504690528151835114620001815760405162d8548360e71b815260040160405180910390fd5b60608401516001600160a01b03161580620001a4575083516001600160a01b0316155b15620001c3576040516342bcdf7f60e11b815260040160405180910390fd5b83600001516001600160a01b0316634120fccd6040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000206573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200022c9190620009bd565b6001600160401b03166001146200025657604051636fc2a20760e11b815260040160405180910390fd5b83516001600160a01b0390811660a090815260408601516001600160401b0390811660c05260208701511660e052606086015182166101005260808601518216610140528501511661016052620002cd7f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b362000538565b6101205260005b83518110156200047e576200032e848281518110620002f757620002f7620009db565b6020026020010151848381518110620003145762000314620009db565b6020026020010151600c6200059f60201b9092919060201c565b50620003e3838281518110620003485762000348620009db565b60200260200101516001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200038e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003b49190620009f1565b848381518110620003c957620003c9620009db565b6020026020010151600f6200059f60201b9092919060201c565b507f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c8482815181106200041a576200041a620009db565b6020026020010151848381518110620004375762000437620009db565b6020026020010151604051620004639291906001600160a01b0392831681529116602082015260400190565b60405180910390a1620004768162000a18565b9050620002d4565b50505050505050505062000a40565b336001600160a01b03821603620004e75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000095565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008160c05160e051610100516040516020016200058294939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b6000620005b7846001600160a01b03851684620005bf565b949350505050565b6000620005b784846001600160a01b03851660008281526002840160205260408120829055620005b784846000620005f8838362000601565b90505b92915050565b60008181526001830160205260408120546200064a57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620005fb565b506000620005fb565b634e487b7160e01b600052604160045260246000fd5b60405160c081016001600160401b03811182821017156200068e576200068e62000653565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620006bf57620006bf62000653565b604052919050565b6001600160a01b0381168114620006dd57600080fd5b50565b80516001600160401b0381168114620006f857600080fd5b919050565b60006001600160401b0382111562000719576200071962000653565b5060051b60200190565b600082601f8301126200073557600080fd5b815160206200074e6200074883620006fd565b62000694565b82815260059290921b840181019181810190868411156200076e57600080fd5b8286015b84811015620007965780516200078881620006c7565b835291830191830162000772565b509695505050505050565b600082601f830112620007b357600080fd5b81516020620007c66200074883620006fd565b82815260059290921b84018101918181019086841115620007e657600080fd5b8286015b84811015620007965780516200080081620006c7565b8352918301918301620007ea565b80516001600160801b0381168114620006f857600080fd5b6000606082840312156200083957600080fd5b604051606081016001600160401b03811182821017156200085e576200085e62000653565b8060405250809150825180151581146200087757600080fd5b815262000887602084016200080e565b60208201526200089a604084016200080e565b60408201525092915050565b600080600080848603610160811215620008bf57600080fd5b60c0811215620008ce57600080fd5b50620008d962000669565b8551620008e681620006c7565b8152620008f660208701620006e0565b60208201526200090960408701620006e0565b604082015260608601516200091e81620006c7565b606082015260808601516200093381620006c7565b608082015260a08601516200094881620006c7565b60a082015260c08601519094506001600160401b03808211156200096b57600080fd5b620009798883890162000723565b945060e08701519150808211156200099057600080fd5b506200099f87828801620007a1565b925050620009b286610100870162000826565b905092959194509250565b600060208284031215620009d057600080fd5b620005f882620006e0565b634e487b7160e01b600052603260045260246000fd5b60006020828403121562000a0457600080fd5b815162000a1181620006c7565b9392505050565b60006001820162000a3957634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e0516101005161012051610140516101605161647a62000b3260003960008181610381015281816123950152612c0201526000818161034501528181611835015281816118b70152818161236e0152818161317801526131ff01526000612dbb01526000818161030901528181612344015261361a0152600081816102a9015281816122f201526135f90152600081816102d90152818161231c01528181612947015281816135d80152613f6801526000818161026d015281816122c40152612eb001526000818161157e015281816115ca0152818161197801526119c4015261647a6000f3fe608060405234801561001057600080fd5b50600436106102265760003560e01c80637437ff9f1161012a578063b4069b31116100bd578063d3c7c2c71161008c578063f2fde38b11610071578063f2fde38b14610820578063f52121a514610833578063f8ccbf471461084657600080fd5b8063d3c7c2c714610805578063d7e2bb501461080d57600080fd5b8063b4069b31146107c4578063b5767166146107d7578063c5a1d7f0146107ea578063c92b2832146107f257600080fd5b8063856c8247116100f9578063856c8247146107475780638da5cb5b14610773578063afcb95d714610791578063b1dc65a4146107b157600080fd5b80637437ff9f1461060a57806379ba50971461070157806381ff70481461070957806385572ffb1461073957600080fd5b8063546719cd116101bd578063681fba161161018c5780636c6bd845116101715780636c6bd845146105d1578063704b6c02146105e4578063740f4150146105f757600080fd5b8063681fba161461059c578063693928ae146105b157600080fd5b8063546719cd146104d1578063599f6431146105355780635d86f14114610574578063666cab8d1461058757600080fd5b80632dea00f3116101f95780632dea00f3146104525780633a87ac53146104655780634f9f03fe14610478578063506449721461049957600080fd5b806306285c691461022b578063142a98fc146103d4578063181f5a77146103f45780631ef381741461043d575b600080fd5b6103be6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a08101919091526040518060c001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815250905090565b6040516103cb91906149d4565b60405180910390f35b6103e76103e2366004614a6b565b610863565b6040516103cb9190614af2565b6104306040518060400160405280601481526020017f45564d3245564d4f666652616d7020312e322e3000000000000000000000000081525081565b6040516103cb9190614b6e565b61045061044b366004614ddc565b6108de565b005b610450610460366004614ea9565b610d9d565b610450610473366004614f2b565b610dab565b61048b610486366004615201565b6111bb565b6040516103cb929190615265565b6104c36104a7366004614a6b565b67ffffffffffffffff1660009081526013602052604090205490565b6040519081526020016103cb565b6104d96111d5565b6040516103cb919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016103cb565b61054f610582366004615285565b61128a565b61058f6112f3565b6040516103cb91906152f3565b6105a4611362565b6040516103cb9190615306565b6105c46105bf3660046153e7565b61141b565b6040516103cb919061550b565b6104506105df366004615707565b611481565b6104506105f2366004615285565b61148b565b610450610605366004615707565b61157b565b6106f46040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c081018252600a5463ffffffff808216835273ffffffffffffffffffffffffffffffffffffffff64010000000090920482166020840152600b549182169383019390935261ffff7401000000000000000000000000000000000000000082041660608301527601000000000000000000000000000000000000000000008104831660808301527a010000000000000000000000000000000000000000000000000000900490911660a082015290565b6040516103cb91906157c2565b6104506116fd565b6007546005546040805163ffffffff808516825264010000000090940490931660208401528201526060016103cb565b610450610226366004615831565b61075a610755366004615285565b6117fa565b60405167ffffffffffffffff90911681526020016103cb565b60005473ffffffffffffffffffffffffffffffffffffffff1661054f565b6040805160018152600060208201819052918101919091526060016103cb565b6104506107bf36600461586c565b611922565b61054f6107d2366004615285565b611bb3565b6104506107e5366004615923565b611c8c565b6104c3611c96565b610450610800366004615985565b611cc1565b6105a4611d46565b61054f61081b366004615285565b611dfb565b61045061082e366004615285565b611e0a565b610450610841366004615201565b611e1b565b6014546108539060ff1681565b60405190151581526020016103cb565b600061087160016004615a22565b600261087e608085615a64565b67ffffffffffffffff166108929190615a8b565b601360006108a1608087615aa2565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002054901c1660038111156108d8576108d8614a88565b92915050565b84518460ff16601f821115610954576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f746f6f206d616e79207472616e736d697474657273000000000000000000000060448201526064015b60405180910390fd5b806000036109be576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f7369746976650000000000000000000000000000604482015260640161094b565b6109c6612075565b6109cf856120f8565b60095460005b81811015610a5b5760086000600983815481106109f4576109f4615ac9565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000169055610a5481615af8565b90506109d5565b50875160005b81811015610c575760008a8281518110610a7d57610a7d615ac9565b6020026020010151905060006002811115610a9a57610a9a614a88565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260086020526040902054610100900460ff166002811115610ad957610ad9614a88565b14610b40576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d6974746572206164647265737300000000604482015260640161094b565b73ffffffffffffffffffffffffffffffffffffffff8116610b8d576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff83168152602081016002905273ffffffffffffffffffffffffffffffffffffffff821660009081526008602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610c3d57610c3d614a88565b02179055509050505080610c5090615af8565b9050610a61565b508851610c6b9060099060208c019061493e565b506006805460ff838116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909216908b161717905560078054610cf1914691309190600090610cc39063ffffffff16615b30565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168d8d8d8d8d8d6123f9565b6005600001819055506000600760049054906101000a900463ffffffff16905043600760046101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581600560000154600760009054906101000a900463ffffffff168e8e8e8e8e8e604051610d8899989796959493929190615b53565b60405180910390a15050505050505050505050565b610da782826124a4565b5050565b610db3612075565b60005b83811015610fb2576000858583818110610dd257610dd2615ac9565b610de89260206040909202019081019150615285565b90506000868684818110610dfe57610dfe615ac9565b9050604002016020016020810190610e169190615285565b9050610e23600c8361254e565b610e59576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610e7b600c84612570565b73ffffffffffffffffffffffffffffffffffffffff1614610ec8576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ed3600c83612592565b50610f4e8173ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f22573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f469190615be9565b600f90612592565b506040805173ffffffffffffffffffffffffffffffffffffffff8085168252831660208201527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a1505080610fab90615af8565b9050610db6565b5060005b818110156111b4576000838383818110610fd257610fd2615ac9565b610fe89260206040909202019081019150615285565b90506000848484818110610ffe57610ffe615ac9565b90506040020160200160208101906110169190615285565b905073ffffffffffffffffffffffffffffffffffffffff8216158061104f575073ffffffffffffffffffffffffffffffffffffffff8116155b15611086576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611091600c8361254e565b156110c8576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110d4600c83836125b4565b506111508173ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611123573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111479190615be9565b600f90836125b4565b506040805173ffffffffffffffffffffffffffffffffffffffff8085168252831660208201527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a15050806111ad90615af8565b9050610fb6565b5050505050565b600060606111c984846125d7565b915091505b9250929050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526003546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff1660208501527401000000000000000000000000000000000000000090920460ff1615159383019390935260045480841660608401520490911660808201526112859061277a565b905090565b60008080611299600c8561282c565b91509150816112ec576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161094b565b9392505050565b6060600980548060200260200160405190810160405280929190818152602001828054801561135857602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161132d575b5050505050905090565b606061136e600f61284f565b67ffffffffffffffff81111561138657611386614b81565b6040519080825280602002602001820160405280156113af578160200160208202803683370190505b50905060005b81518110156114175760006113cb600f8361285a565b509050808383815181106113e1576113e1615ac9565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101525061141081615af8565b90506113b5565b5090565b60606114748989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b925061146591508990508a615c06565b61146f8789615c06565b612876565b9998505050505050505050565b610da78282612c00565b60005473ffffffffffffffffffffffffffffffffffffffff1633148015906114cb575060025473ffffffffffffffffffffffffffffffffffffffff163314155b15611502576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c9060200160405180910390a150565b467f000000000000000000000000000000000000000000000000000000000000000014611606576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015267ffffffffffffffff4616602482015260440161094b565b81515181518114611643576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b818110156116ed57600083828151811061166257611662615ac9565b602002602001015190508060001415801561169b5750845180518390811061168c5761168c615ac9565b60200260200101516080015181105b156116dc576040517f085e39cf000000000000000000000000000000000000000000000000000000008152600481018390526024810182905260440161094b565b506116e681615af8565b9050611646565b506116f88383612c00565b505050565b60015473ffffffffffffffffffffffffffffffffffffffff16331461177e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161094b565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b73ffffffffffffffffffffffffffffffffffffffff811660009081526012602052604081205467ffffffffffffffff168015801561186d57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1615155b156108d8576040517f856c824700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063856c824790602401602060405180830381865afa1580156118fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112ec9190615c13565b61192c87876135b0565b600554883590808214611975576040517f93df584c000000000000000000000000000000000000000000000000000000008152600481018290526024810183905260440161094b565b467f0000000000000000000000000000000000000000000000000000000000000000146119f6576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015246602482015260440161094b565b6040805183815260208c81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a13360009081526008602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115611a7e57611a7e614a88565b6002811115611a8f57611a8f614a88565b9052509050600281602001516002811115611aac57611aac614a88565b148015611af357506009816000015160ff1681548110611ace57611ace615ac9565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b611b29576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506000611b37856020615a8b565b611b42886020615a8b565b611b4e8b610144615c30565b611b589190615c30565b611b629190615c30565b9050368114611ba6576040517f8e1192e10000000000000000000000000000000000000000000000000000000081526004810182905236602482015260440161094b565b5050505050505050505050565b60008080611bc2600c8561282c565b9150915081611c15576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161094b565b8073ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c849190615be9565b949350505050565b610da782826135b0565b60006112857f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b36135d3565b60005473ffffffffffffffffffffffffffffffffffffffff163314801590611d01575060025473ffffffffffffffffffffffffffffffffffffffff163314155b15611d38576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611d436003826136a0565b50565b6060611d52600c61284f565b67ffffffffffffffff811115611d6a57611d6a614b81565b604051908082528060200260200182016040528015611d93578160200160208202803683370190505b50905060005b8151811015611417576000611daf600c8361285a565b50905080838381518110611dc557611dc5615ac9565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015250611df481615af8565b9050611d99565b60008080611299600f8561282c565b611e12612075565b611d4381613885565b333014611e54576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160008082526020820190925281611e91565b6040805180820190915260008082526020820152815260200190600190039081611e6a5790505b506101408401515190915015611efe57611efb8361014001518460200151604051602001611edb919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b604051602081830303815290604052856040015186610160015186612876565b90505b604083015173ffffffffffffffffffffffffffffffffffffffff163b1580611f6857506040830151611f669073ffffffffffffffffffffffffffffffffffffffff167f85572ffb0000000000000000000000000000000000000000000000000000000061397a565b155b15611f7257505050565b600a546000908190640100000000900473ffffffffffffffffffffffffffffffffffffffff16633cf97983611fa78786613996565b611388886080015189604001516040518563ffffffff1660e01b8152600401611fd39493929190615c43565b6000604051808303816000875af1158015611ff2573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526120389190810190615d15565b91509150816111b457806040517f0a8d6e8c00000000000000000000000000000000000000000000000000000000815260040161094b9190614b6e565b60005473ffffffffffffffffffffffffffffffffffffffff1633146120f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161094b565b565b60008180602001905181019061210e9190615db6565b602081015190915073ffffffffffffffffffffffffffffffffffffffff16612162576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600a805460208085015173ffffffffffffffffffffffffffffffffffffffff908116640100000000027fffffffffffffffff00000000000000000000000000000000000000000000000090931663ffffffff9586161792909217909255604080850151600b80546060808901516080808b015160a0808d01518c167a010000000000000000000000000000000000000000000000000000027fffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffffff92909c1676010000000000000000000000000000000000000000000002919091167fffff0000000000000000ffffffffffffffffffffffffffffffffffffffffffff61ffff90941674010000000000000000000000000000000000000000027fffffffffffffffffffff00000000000000000000000000000000000000000000909616978a169790971794909417919091169490941797909717909155825160c0810184527f00000000000000000000000000000000000000000000000000000000000000008516815267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116968201969096527f0000000000000000000000000000000000000000000000000000000000000000909516858401527f00000000000000000000000000000000000000000000000000000000000000008416958501959095527f00000000000000000000000000000000000000000000000000000000000000008316908401527f00000000000000000000000000000000000000000000000000000000000000009091169282019290925290517fe668e1a4644c1a030b909bbfd837f5cfa914994ed5e0bb2e9c34a5c37753128a916123ed918490615e62565b60405180910390a15050565b6000808a8a8a8a8a8a8a8a8a60405160200161241d99989796959493929190615f3e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b600060026124b3608085615a64565b67ffffffffffffffff166124c79190615a8b565b905060006013816124d9608087615aa2565b67ffffffffffffffff16815260208101919091526040016000205490508161250360016004615a22565b901b19168183600381111561251a5761251a614a88565b901b17806013600061252d608088615aa2565b67ffffffffffffffff16815260208101919091526040016000205550505050565b60006112ec8373ffffffffffffffffffffffffffffffffffffffff8416613a46565b60006112ec8373ffffffffffffffffffffffffffffffffffffffff8416613a52565b60006112ec8373ffffffffffffffffffffffffffffffffffffffff8416613a5e565b6000611c848473ffffffffffffffffffffffffffffffffffffffff851684613a6a565b6040517ff52121a5000000000000000000000000000000000000000000000000000000008152600090606090309063f52121a59061261b9087908790600401616028565b600060405180830381600087803b15801561263557600080fd5b505af1925050508015612646575060015b61275f573d808015612674576040519150601f19603f3d011682016040523d82523d6000602084013e612679565b606091505b50612683816161b2565b7fffffffff00000000000000000000000000000000000000000000000000000000167f0a8d6e8c00000000000000000000000000000000000000000000000000000000148061271b57506126d6816161b2565b7fffffffff00000000000000000000000000000000000000000000000000000000167fe1cd550900000000000000000000000000000000000000000000000000000000145b1561272b576003925090506111ce565b806040517fcf19edfd00000000000000000000000000000000000000000000000000000000815260040161094b9190614b6e565b50506040805160208101909152600081526002909250929050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261280882606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff16426127ec9190615a22565b85608001516fffffffffffffffffffffffffffffffff16613a8d565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b6000806111c98473ffffffffffffffffffffffffffffffffffffffff8516613aac565b60006108d882613abb565b60008080806128698686613ac6565b9097909650945050505050565b60606000865167ffffffffffffffff81111561289457612894614b81565b6040519080825280602002602001820160405280156128d957816020015b60408051808201909152600080825260208201528152602001906001900390816128b25790505b50905060005b8751811015612bcf5760006129108983815181106128ff576128ff615ac9565b60200260200101516000015161128a565b9050600089838151811061292657612926615ac9565b6020026020010151602001519050600080612a9b638627fad660e01b8c8c867f00000000000000000000000000000000000000000000000000000000000000008e8b8151811061297857612978615ac9565b60200260200101518e8c8151811061299257612992615ac9565b60200260200101516040516020016129ab929190616202565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526129ea9594939291602401616227565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152600b54869063ffffffff7a010000000000000000000000000000000000000000000000000000909104166084611388613ad5565b9150915081612ad857806040517fe1cd550900000000000000000000000000000000000000000000000000000000815260040161094b9190614b6e565b8373ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612b23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b479190615be9565b868681518110612b5957612b59615ac9565b60200260200101516000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505082868681518110612baa57612baa615ac9565b602002602001015160200181815250505050505080612bc890615af8565b90506128df565b50600b54612bf490829073ffffffffffffffffffffffffffffffffffffffff16613bf8565b90505b95945050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663397796f76040518163ffffffff1660e01b8152600401602060405180830381865afa158015612c6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c8f919061628a565b15612cc6576040517fc148371500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151516000819003612d03576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8260200151518114612d41576040517f57e0e08300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115612d5c57612d5c614b81565b604051908082528060200260200182016040528015612d85578160200160208202803683370190505b50905060005b82811015612e6557600085600001518281518110612dab57612dab615ac9565b60200260200101519050612ddf817f0000000000000000000000000000000000000000000000000000000000000000613de0565b838381518110612df157612df1615ac9565b602002602001018181525050806101800151838381518110612e1557612e15615ac9565b602002602001015114612e54576040517f7185cf6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50612e5e81615af8565b9050612d8b565b50604080850151606086015191517f3204887500000000000000000000000000000000000000000000000000000000815260009273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692633204887592612ee6928792916004016162d7565b602060405180830381865afa158015612f03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f27919061630d565b905080600003612f63576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8351151560005b848110156135a757600087600001518281518110612f8a57612f8a615ac9565b602002602001015190506000612fa38260600151610863565b90506000816003811115612fb957612fb9614a88565b1480612fd657506003816003811115612fd457612fd4614a88565b145b61301e5760608201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161094b565b83156130db57600a5460009063ffffffff1661303a8742615a22565b119050808061305a5750600382600381111561305857613058614a88565b145b613090576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8884815181106130a2576130a2615ac9565b60200260200101516000146130d5578884815181106130c3576130c3615ac9565b60200260200101518360800181815250505b50613138565b60008160038111156130ef576130ef614a88565b146131385760608201516040517f67d9ba0f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161094b565b60208083015173ffffffffffffffffffffffffffffffffffffffff1660009081526012909152604090205467ffffffffffffffff16801580156131b057507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1615155b156133535760208301516040517f856c824700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201527f00000000000000000000000000000000000000000000000000000000000000009091169063856c824790602401602060405180830381865afa158015613248573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061326c9190615c13565b60c084015190915067ffffffffffffffff16613289826001616326565b67ffffffffffffffff16146132f657826020015173ffffffffffffffffffffffffffffffffffffffff168360c0015167ffffffffffffffff167fe44a20935573a783dd0d5991c92d7b6a0eb3173566530364db3ec10e9a990b5d60405160405180910390a3505050613597565b60208381015173ffffffffffffffffffffffffffffffffffffffff16600090815260129091526040902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83161790555b600082600381111561336757613367614a88565b036133f35760c083015167ffffffffffffffff16613386826001616326565b67ffffffffffffffff16146133f357826020015173ffffffffffffffffffffffffffffffffffffffff168360c0015167ffffffffffffffff167fd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf4123760405160405180910390a3505050613597565b60008a60200151858151811061340b5761340b615ac9565b60200260200101519050613420848251613f66565b61342f846060015160016124a4565b60008061343c86846125d7565b9150915061344e8660600151836124a4565b600382600381111561346257613462614a88565b141580156134825750600282600381111561347f5761347f614a88565b14155b156134c1578560600151826040517f9e26160300000000000000000000000000000000000000000000000000000000815260040161094b929190616347565b60008560038111156134d5576134d5614a88565b036135425760208087015173ffffffffffffffffffffffffffffffffffffffff166000908152601290915260408120805467ffffffffffffffff169161351a83616365565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b856101800151866060015167ffffffffffffffff167fd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef658484604051613588929190615265565b60405180910390a35050505050505b6135a081615af8565b9050612f6a565b50505050505050565b610da76135bf82840184616382565b604080516000815260208101909152612c00565b6000817f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000604051602001613683949392919093845267ffffffffffffffff92831660208501529116604083015273ffffffffffffffffffffffffffffffffffffffff16606082015260800190565b604051602081830303815290604052805190602001209050919050565b81546000906136c990700100000000000000000000000000000000900463ffffffff1642615a22565b9050801561376b5760018301548354613711916fffffffffffffffffffffffffffffffff80821692811691859170010000000000000000000000000000000090910416613a8d565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354613791916fffffffffffffffffffffffffffffffff908116911661413c565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19906138789084908151151581526020808301516fffffffffffffffffffffffffffffffff90811691830191909152604092830151169181019190915260600190565b60405180910390a1505050565b3373ffffffffffffffffffffffffffffffffffffffff821603613904576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161094b565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600061398583614152565b80156112ec57506112ec83836141b6565b6040805160a08101825260008082526020820152606091810182905281810182905260808101919091526040518060a001604052808461018001518152602001846000015167ffffffffffffffff1681526020018460200151604051602001613a1b919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b6040516020818303038152906040528152602001846101200151815260200183815250905092915050565b60006112ec8383614285565b60006112ec8383614291565b60006112ec838361431b565b6000611c84848473ffffffffffffffffffffffffffffffffffffffff8516614338565b6000612bf785613a9d8486615a8b565b613aa79087615c30565b61413c565b60008080806128698686614355565b60006108d88261438f565b6000808080612869868661439a565b600060608361ffff1667ffffffffffffffff811115613af657613af6614b81565b6040519080825280601f01601f191660200182016040528015613b20576020820181803683370190505b5090507f0c3b563c000000000000000000000000000000000000000000000000000000007fafa32a2c000000000000000000000000000000000000000000000000000000007f37c3be2900000000000000000000000000000000000000000000000000000000883b613b96578260005260046000fd5b5a86811015613ba9578260005260046000fd5b8690036040810481038910613bc2578160005260046000fd5b506000808b5160208d0160008d8df194503d87811115613bdf5750865b808552806000602087013e505050509550959350505050565b81516000805b82811015613dcc5760008473ffffffffffffffffffffffffffffffffffffffff1663d02641a0878481518110613c3657613c36615ac9565b6020908102919091010151516040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016040805180830381865afa158015613caa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613cce91906163b7565b5190507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116600003613d6057858281518110613d0957613d09615ac9565b6020908102919091010151516040517f9a655f7b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116600482015260240161094b565b613dae868381518110613d7557613d75615ac9565b602002602001015160200151827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166143c590919063ffffffff16565b613db89084615c30565b92505080613dc590615af8565b9050613bfe565b50613dda6003826000614402565b50505050565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b6101000151604051602001613e8398979695949392919073ffffffffffffffffffffffffffffffffffffffff9889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b6040516020818303038152906040528051906020012085610120015180519060200120866101400151604051602001613ebc919061550b565b60405160208183030381529060405280519060200120876101600151604051602001613ee89190616417565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16826000015167ffffffffffffffff1614613fe65781516040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161094b565b600b54610140830151517401000000000000000000000000000000000000000090910461ffff1610156140575760608201516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161094b565b8082610140015151146140a85760608201516040517f8808f8e700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260240161094b565b600b546101208301515176010000000000000000000000000000000000000000000090910463ffffffff161015610da757600b54610120830151516040517f8693378900000000000000000000000000000000000000000000000000000000815276010000000000000000000000000000000000000000000090920463ffffffff166004830152602482015260440161094b565b600081831061414b57816112ec565b5090919050565b600061417e827f01ffc9a7000000000000000000000000000000000000000000000000000000006141b6565b80156108d857506141af827fffffffff000000000000000000000000000000000000000000000000000000006141b6565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d9150600051905082801561426e575060208210155b801561427a5750600081115b979650505050505050565b60006112ec8383614785565b6000818152600283016020526040812054801515806142b557506142b58484614285565b6112ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b65790000604482015260640161094b565b600081815260028301602052604081208190556112ec838361479d565b60008281526002840160205260408120829055611c8484846147a9565b6000818152600283016020526040812054819080614384576143778585614285565b9250600091506111ce9050565b6001925090506111ce565b60006108d8826147b5565b600080806143a885856147bf565b600081815260029690960160205260409095205494959350505050565b6000670de0b6b3a76400006143f8837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8616615a8b565b6112ec919061642a565b825474010000000000000000000000000000000000000000900460ff161580614429575081155b1561443357505050565b825460018401546fffffffffffffffffffffffffffffffff8083169291169060009061447990700100000000000000000000000000000000900463ffffffff1642615a22565b9050801561453957818311156144bb576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018601546144f59083908590849070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16613a8d565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000004263ffffffff160217875592505b848210156145f05773ffffffffffffffffffffffffffffffffffffffff8416614598576040517ff94ebcd1000000000000000000000000000000000000000000000000000000008152600481018390526024810186905260440161094b565b6040517f1a76572a000000000000000000000000000000000000000000000000000000008152600481018390526024810186905273ffffffffffffffffffffffffffffffffffffffff8516604482015260640161094b565b848310156147035760018681015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff169060009082906146349082615a22565b61463e878a615a22565b6146489190615c30565b614652919061642a565b905073ffffffffffffffffffffffffffffffffffffffff86166146ab576040517f15279c08000000000000000000000000000000000000000000000000000000008152600481018290526024810186905260440161094b565b6040517fd0c8d23a000000000000000000000000000000000000000000000000000000008152600481018290526024810186905273ffffffffffffffffffffffffffffffffffffffff8716604482015260640161094b565b61470d8584615a22565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b600081815260018301602052604081205415156112ec565b60006112ec83836147cb565b60006112ec83836148c5565b60006108d8825490565b60006112ec8383614914565b600081815260018301602052604081205480156148b45760006147ef600183615a22565b855490915060009061480390600190615a22565b905081811461486857600086600001828154811061482357614823615ac9565b906000526020600020015490508087600001848154811061484657614846615ac9565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806148795761487961643e565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506108d8565b60009150506108d8565b5092915050565b600081815260018301602052604081205461490c575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556108d8565b5060006108d8565b600082600001828154811061492b5761492b615ac9565b9060005260206000200154905092915050565b8280548282559060005260206000209081019282156149b8579160200282015b828111156149b857825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90911617825560209092019160019091019061495e565b506114179291505b8082111561141757600081556001016149c0565b60c081016108d8828473ffffffffffffffffffffffffffffffffffffffff808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015250508060608301511660608401528060808301511660808401528060a08301511660a0840152505050565b67ffffffffffffffff81168114611d4357600080fd5b8035614a6681614a45565b919050565b600060208284031215614a7d57600080fd5b81356112ec81614a45565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110614aee577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b602081016108d88284614ab7565b60005b83811015614b1b578181015183820152602001614b03565b50506000910152565b60008151808452614b3c816020860160208601614b00565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006112ec6020830184614b24565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715614bd357614bd3614b81565b60405290565b6040516101a0810167ffffffffffffffff81118282101715614bd357614bd3614b81565b6040516080810167ffffffffffffffff81118282101715614bd357614bd3614b81565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715614c6757614c67614b81565b604052919050565b600067ffffffffffffffff821115614c8957614c89614b81565b5060051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81168114611d4357600080fd5b8035614a6681614c93565b600082601f830112614cd157600080fd5b81356020614ce6614ce183614c6f565b614c20565b82815260059290921b84018101918181019086841115614d0557600080fd5b8286015b84811015614d29578035614d1c81614c93565b8352918301918301614d09565b509695505050505050565b803560ff81168114614a6657600080fd5b600067ffffffffffffffff821115614d5f57614d5f614b81565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112614d9c57600080fd5b8135614daa614ce182614d45565b818152846020838601011115614dbf57600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c08789031215614df557600080fd5b863567ffffffffffffffff80821115614e0d57600080fd5b614e198a838b01614cc0565b97506020890135915080821115614e2f57600080fd5b614e3b8a838b01614cc0565b9650614e4960408a01614d34565b95506060890135915080821115614e5f57600080fd5b614e6b8a838b01614d8b565b9450614e7960808a01614a5b565b935060a0890135915080821115614e8f57600080fd5b50614e9c89828a01614d8b565b9150509295509295509295565b60008060408385031215614ebc57600080fd5b8235614ec781614a45565b9150602083013560048110614edb57600080fd5b809150509250929050565b60008083601f840112614ef857600080fd5b50813567ffffffffffffffff811115614f1057600080fd5b6020830191508360208260061b85010111156111ce57600080fd5b60008060008060408587031215614f4157600080fd5b843567ffffffffffffffff80821115614f5957600080fd5b614f6588838901614ee6565b90965094506020870135915080821115614f7e57600080fd5b50614f8b87828801614ee6565b95989497509550505050565b8015158114611d4357600080fd5b8035614a6681614f97565b600082601f830112614fc157600080fd5b81356020614fd1614ce183614c6f565b82815260069290921b84018101918181019086841115614ff057600080fd5b8286015b84811015614d29576040818903121561500d5760008081fd5b615015614bb0565b813561502081614c93565b81528185013585820152835291830191604001614ff4565b6000615046614ce184614c6f565b8381529050602080820190600585901b84018681111561506557600080fd5b845b818110156150a157803567ffffffffffffffff8111156150875760008081fd5b61509389828901614d8b565b855250928201928201615067565b505050509392505050565b600082601f8301126150bd57600080fd5b6112ec83833560208501615038565b60006101a082840312156150df57600080fd5b6150e7614bd9565b90506150f282614a5b565b815261510060208301614cb5565b602082015261511160408301614cb5565b604082015261512260608301614a5b565b60608201526080820135608082015261513d60a08301614fa5565b60a082015261514e60c08301614a5b565b60c082015261515f60e08301614cb5565b60e082015261010082810135908201526101208083013567ffffffffffffffff8082111561518c57600080fd5b61519886838701614d8b565b838501526101409250828501359150808211156151b457600080fd5b6151c086838701614fb0565b838501526101609250828501359150808211156151dc57600080fd5b506151e9858286016150ac565b82840152505061018080830135818301525092915050565b6000806040838503121561521457600080fd5b823567ffffffffffffffff8082111561522c57600080fd5b615238868387016150cc565b9350602085013591508082111561524e57600080fd5b5061525b858286016150ac565b9150509250929050565b61526f8184614ab7565b604060208201526000611c846040830184614b24565b60006020828403121561529757600080fd5b81356112ec81614c93565b600081518084526020808501945080840160005b838110156152e857815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016152b6565b509495945050505050565b6020815260006112ec60208301846152a2565b6020808252825182820181905260009190848201906040850190845b8181101561535457835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101615322565b50909695505050505050565b60008083601f84011261537257600080fd5b50813567ffffffffffffffff81111561538a57600080fd5b6020830191508360208285010111156111ce57600080fd5b60008083601f8401126153b457600080fd5b50813567ffffffffffffffff8111156153cc57600080fd5b6020830191508360208260051b85010111156111ce57600080fd5b60008060008060008060008060a0898b03121561540357600080fd5b883567ffffffffffffffff8082111561541b57600080fd5b6154278c838d01614fb0565b995060208b013591508082111561543d57600080fd5b6154498c838d01615360565b909950975060408b0135915061545e82614c93565b90955060608a0135908082111561547457600080fd5b6154808c838d016153a2565b909650945060808b013591508082111561549957600080fd5b506154a68b828c016153a2565b999c989b5096995094979396929594505050565b600081518084526020808501945080840160005b838110156152e8578151805173ffffffffffffffffffffffffffffffffffffffff16885283015183880152604090960195908201906001016154ce565b6020815260006112ec60208301846154ba565b600082601f83011261552f57600080fd5b8135602061553f614ce183614c6f565b82815260059290921b8401810191818101908684111561555e57600080fd5b8286015b84811015614d2957803567ffffffffffffffff8111156155825760008081fd5b6155908986838b01016150ac565b845250918301918301615562565b600082601f8301126155af57600080fd5b813560206155bf614ce183614c6f565b82815260059290921b840181019181810190868411156155de57600080fd5b8286015b84811015614d2957803583529183019183016155e2565b60006080828403121561560b57600080fd5b615613614bfd565b9050813567ffffffffffffffff8082111561562d57600080fd5b818401915084601f83011261564157600080fd5b81356020615651614ce183614c6f565b82815260059290921b8401810191818101908884111561567057600080fd5b8286015b848110156156a85780358681111561568c5760008081fd5b61569a8b86838b01016150cc565b845250918301918301615674565b50865250858101359350828411156156bf57600080fd5b6156cb8785880161551e565b908501525060408401359150808211156156e457600080fd5b506156f18482850161559e565b6040830152506060820135606082015292915050565b6000806040838503121561571a57600080fd5b823567ffffffffffffffff8082111561573257600080fd5b61573e868387016155f9565b935060209150818501358181111561575557600080fd5b85019050601f8101861361576857600080fd5b8035615776614ce182614c6f565b81815260059190911b8201830190838101908883111561579557600080fd5b928401925b828410156157b35783358252928401929084019061579a565b80955050505050509250929050565b60c081016108d8828463ffffffff808251168352602082015173ffffffffffffffffffffffffffffffffffffffff8082166020860152806040850151166040860152505061ffff60608301511660608401528060808301511660808401528060a08301511660a0840152505050565b60006020828403121561584357600080fd5b813567ffffffffffffffff81111561585a57600080fd5b820160a081850312156112ec57600080fd5b60008060008060008060008060e0898b03121561588857600080fd5b606089018a81111561589957600080fd5b8998503567ffffffffffffffff808211156158b357600080fd5b6158bf8c838d01615360565b909950975060808b01359150808211156158d857600080fd5b6158e48c838d016153a2565b909750955060a08b01359150808211156158fd57600080fd5b5061590a8b828c016153a2565b999c989b50969995989497949560c00135949350505050565b6000806020838503121561593657600080fd5b823567ffffffffffffffff81111561594d57600080fd5b61595985828601615360565b90969095509350505050565b80356fffffffffffffffffffffffffffffffff81168114614a6657600080fd5b60006060828403121561599757600080fd5b6040516060810181811067ffffffffffffffff821117156159ba576159ba614b81565b60405282356159c881614f97565b81526159d660208401615965565b60208201526159e760408401615965565b60408201529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156108d8576108d86159f3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff80841680615a7f57615a7f615a35565b92169190910692915050565b80820281158282048414176108d8576108d86159f3565b600067ffffffffffffffff80841680615abd57615abd615a35565b92169190910492915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203615b2957615b296159f3565b5060010190565b600063ffffffff808316818103615b4957615b496159f3565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152615b838184018a6152a2565b90508281036080840152615b9781896152a2565b905060ff871660a084015282810360c0840152615bb48187614b24565b905067ffffffffffffffff851660e0840152828103610100840152615bd98185614b24565b9c9b505050505050505050505050565b600060208284031215615bfb57600080fd5b81516112ec81614c93565b60006112ec368484615038565b600060208284031215615c2557600080fd5b81516112ec81614a45565b808201808211156108d8576108d86159f3565b608081528451608082015267ffffffffffffffff60208601511660a08201526000604086015160a060c0840152615c7e610120840182614b24565b905060608701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80808584030160e0860152615cba8383614b24565b925060808901519150808584030161010086015250615cd982826154ba565b92505050615ced602083018661ffff169052565b836040830152612bf7606083018473ffffffffffffffffffffffffffffffffffffffff169052565b60008060408385031215615d2857600080fd5b8251615d3381614f97565b602084015190925067ffffffffffffffff811115615d5057600080fd5b8301601f81018513615d6157600080fd5b8051615d6f614ce182614d45565b818152866020838501011115615d8457600080fd5b615d95826020830160208601614b00565b8093505050509250929050565b805163ffffffff81168114614a6657600080fd5b600060c08284031215615dc857600080fd5b60405160c0810181811067ffffffffffffffff82111715615deb57615deb614b81565b604052615df783615da2565b81526020830151615e0781614c93565b60208201526040830151615e1a81614c93565b6040820152606083015161ffff81168114615e3457600080fd5b6060820152615e4560808401615da2565b6080820152615e5660a08401615da2565b60a08201529392505050565b6101808101615ed4828573ffffffffffffffffffffffffffffffffffffffff808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015250508060608301511660608401528060808301511660808401528060a08301511660a0840152505050565b825163ffffffff90811660c0840152602084015173ffffffffffffffffffffffffffffffffffffffff90811660e0850152604085015116610100840152606084015161ffff166101208401526080840151811661014084015260a0840151166101608301526112ec565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b166040850152816060850152615f858285018b6152a2565b91508382036080850152615f99828a6152a2565b915060ff881660a085015283820360c0850152615fb68288614b24565b90861660e08501528381036101008501529050615bd98185614b24565b600081518084526020808501808196508360051b8101915082860160005b8581101561601b578284038952616009848351614b24565b98850198935090840190600101615ff1565b5091979650505050505050565b6040815261604360408201845167ffffffffffffffff169052565b6000602084015161606c606084018273ffffffffffffffffffffffffffffffffffffffff169052565b50604084015173ffffffffffffffffffffffffffffffffffffffff8116608084015250606084015167ffffffffffffffff811660a084015250608084015160c083015260a084015180151560e08401525060c08401516101006160da8185018367ffffffffffffffff169052565b60e086015191506101206161058186018473ffffffffffffffffffffffffffffffffffffffff169052565b81870151925061014091508282860152808701519250506101a061016081818701526161356101e0870185614b24565b93508288015192507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc061018081888703018189015261617486866154ba565b9550828a0151945081888703018489015261618f8686615fd3565b9550808a01516101c089015250505050508281036020840152612bf78185615fd3565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156161fa5780818460040360031b1b83161693505b505050919050565b6040815260006162156040830185614b24565b8281036020840152612bf78185614b24565b60a08152600061623a60a0830188614b24565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015267ffffffffffffffff85166060840152828103608084015261627e8185614b24565b98975050505050505050565b60006020828403121561629c57600080fd5b81516112ec81614f97565b600081518084526020808501945080840160005b838110156152e8578151875295820195908201906001016162bb565b6060815260006162ea60608301866162a7565b82810360208401526162fc81866162a7565b915050826040830152949350505050565b60006020828403121561631f57600080fd5b5051919050565b67ffffffffffffffff8181168382160190808211156148be576148be6159f3565b67ffffffffffffffff83168152604081016112ec6020830184614ab7565b600067ffffffffffffffff808316818103615b4957615b496159f3565b60006020828403121561639457600080fd5b813567ffffffffffffffff8111156163ab57600080fd5b611c84848285016155f9565b6000604082840312156163c957600080fd5b6163d1614bb0565b82517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811681146163fd57600080fd5b815261640b60208401615da2565b60208201529392505050565b6020815260006112ec6020830184615fd3565b60008261643957616439615a35565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a",
}

var EVM2EVMOffRampHelperABI = EVM2EVMOffRampHelperMetaData.ABI

var EVM2EVMOffRampHelperBin = EVM2EVMOffRampHelperMetaData.Bin

func DeployEVM2EVMOffRampHelper(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMOffRampStaticConfig, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig RateLimiterConfig) (common.Address, *types.Transaction, *EVM2EVMOffRampHelper, error) {
	parsed, err := EVM2EVMOffRampHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMOffRampHelperBin), backend, staticConfig, sourceTokens, pools, rateLimiterConfig)
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
	parsed, err := EVM2EVMOffRampHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) ISSCRIPT(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "IS_SCRIPT")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ISSCRIPT() (bool, error) {
	return _EVM2EVMOffRampHelper.Contract.ISSCRIPT(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) ISSCRIPT() (bool, error) {
	return _EVM2EVMOffRampHelper.Contract.ISSCRIPT(&_EVM2EVMOffRampHelper.CallOpts)
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "currentRateLimiterState")

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _EVM2EVMOffRampHelper.Contract.CurrentRateLimiterState(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _EVM2EVMOffRampHelper.Contract.CurrentRateLimiterState(&_EVM2EVMOffRampHelper.CallOpts)
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMOffRampDynamicConfig, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(EVM2EVMOffRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMOffRampDynamicConfig)).(*EVM2EVMOffRampDynamicConfig)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetDynamicConfig() (EVM2EVMOffRampDynamicConfig, error) {
	return _EVM2EVMOffRampHelper.Contract.GetDynamicConfig(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetDynamicConfig() (EVM2EVMOffRampDynamicConfig, error) {
	return _EVM2EVMOffRampHelper.Contract.GetDynamicConfig(&_EVM2EVMOffRampHelper.CallOpts)
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetExecutionStateBitMap(opts *bind.CallOpts, bitmapIndex uint64) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getExecutionStateBitMap", bitmapIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetExecutionStateBitMap(bitmapIndex uint64) (*big.Int, error) {
	return _EVM2EVMOffRampHelper.Contract.GetExecutionStateBitMap(&_EVM2EVMOffRampHelper.CallOpts, bitmapIndex)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetExecutionStateBitMap(bitmapIndex uint64) (*big.Int, error) {
	return _EVM2EVMOffRampHelper.Contract.GetExecutionStateBitMap(&_EVM2EVMOffRampHelper.CallOpts, bitmapIndex)
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCaller) GetStaticConfig(opts *bind.CallOpts) (EVM2EVMOffRampStaticConfig, error) {
	var out []interface{}
	err := _EVM2EVMOffRampHelper.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(EVM2EVMOffRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMOffRampStaticConfig)).(*EVM2EVMOffRampStaticConfig)

	return out0, err

}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) GetStaticConfig() (EVM2EVMOffRampStaticConfig, error) {
	return _EVM2EVMOffRampHelper.Contract.GetStaticConfig(&_EVM2EVMOffRampHelper.CallOpts)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperCallerSession) GetStaticConfig() (EVM2EVMOffRampStaticConfig, error) {
	return _EVM2EVMOffRampHelper.Contract.GetStaticConfig(&_EVM2EVMOffRampHelper.CallOpts)
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ApplyPoolUpdates(opts *bind.TransactOpts, removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "applyPoolUpdates", removes, adds)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ApplyPoolUpdates(removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ApplyPoolUpdates(&_EVM2EVMOffRampHelper.TransactOpts, removes, adds)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ApplyPoolUpdates(removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ApplyPoolUpdates(&_EVM2EVMOffRampHelper.TransactOpts, removes, adds)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) Execute(opts *bind.TransactOpts, rep InternalExecutionReport, manualExecGasLimits []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "execute", rep, manualExecGasLimits)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Execute(rep InternalExecutionReport, manualExecGasLimits []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Execute(&_EVM2EVMOffRampHelper.TransactOpts, rep, manualExecGasLimits)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) Execute(rep InternalExecutionReport, manualExecGasLimits []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Execute(&_EVM2EVMOffRampHelper.TransactOpts, rep, manualExecGasLimits)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "executeSingleMessage", message, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMOffRampHelper.TransactOpts, message, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMOffRampHelper.TransactOpts, message, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ManuallyExecute(opts *bind.TransactOpts, report InternalExecutionReport, gasLimitOverrides []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "manuallyExecute", report, gasLimitOverrides)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ManuallyExecute(report InternalExecutionReport, gasLimitOverrides []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ManuallyExecute(&_EVM2EVMOffRampHelper.TransactOpts, report, gasLimitOverrides)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ManuallyExecute(report InternalExecutionReport, gasLimitOverrides []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ManuallyExecute(&_EVM2EVMOffRampHelper.TransactOpts, report, gasLimitOverrides)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ReleaseOrMintTokens(opts *bind.TransactOpts, sourceTokenAmounts []ClientEVMTokenAmount, originalSender []byte, receiver common.Address, sourceTokenData [][]byte, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "releaseOrMintTokens", sourceTokenAmounts, originalSender, receiver, sourceTokenData, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ReleaseOrMintTokens(sourceTokenAmounts []ClientEVMTokenAmount, originalSender []byte, receiver common.Address, sourceTokenData [][]byte, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ReleaseOrMintTokens(&_EVM2EVMOffRampHelper.TransactOpts, sourceTokenAmounts, originalSender, receiver, sourceTokenData, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ReleaseOrMintTokens(sourceTokenAmounts []ClientEVMTokenAmount, originalSender []byte, receiver common.Address, sourceTokenData [][]byte, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ReleaseOrMintTokens(&_EVM2EVMOffRampHelper.TransactOpts, sourceTokenAmounts, originalSender, receiver, sourceTokenData, offchainTokenData)
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setAdmin", newAdmin)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetAdmin(&_EVM2EVMOffRampHelper.TransactOpts, newAdmin)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetAdmin(&_EVM2EVMOffRampHelper.TransactOpts, newAdmin)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetExecutionStateHelper(opts *bind.TransactOpts, sequenceNumber uint64, state uint8) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setExecutionStateHelper", sequenceNumber, state)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetExecutionStateHelper(sequenceNumber uint64, state uint8) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetExecutionStateHelper(&_EVM2EVMOffRampHelper.TransactOpts, sequenceNumber, state)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetExecutionStateHelper(sequenceNumber uint64, state uint8) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetExecutionStateHelper(&_EVM2EVMOffRampHelper.TransactOpts, sequenceNumber, state)
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetRateLimiterConfig(&_EVM2EVMOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.SetRateLimiterConfig(&_EVM2EVMOffRampHelper.TransactOpts, config)
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "transmit", reportContext, report, rs, ss, arg4)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Transmit(&_EVM2EVMOffRampHelper.TransactOpts, reportContext, report, rs, ss, arg4)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Transmit(&_EVM2EVMOffRampHelper.TransactOpts, reportContext, report, rs, ss, arg4)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) TrialExecute(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "trialExecute", message, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) TrialExecute(message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.TrialExecute(&_EVM2EVMOffRampHelper.TransactOpts, message, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) TrialExecute(message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.TrialExecute(&_EVM2EVMOffRampHelper.TransactOpts, message, offchainTokenData)
}

type EVM2EVMOffRampHelperAdminSetIterator struct {
	Event *EVM2EVMOffRampHelperAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperAdminSet)
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
		it.Event = new(EVM2EVMOffRampHelperAdminSet)
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

func (it *EVM2EVMOffRampHelperAdminSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperAdminSet struct {
	NewAdmin common.Address
	Raw      types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperAdminSetIterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperAdminSetIterator{contract: _EVM2EVMOffRampHelper.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperAdminSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperAdminSet)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseAdminSet(log types.Log) (*EVM2EVMOffRampHelperAdminSet, error) {
	event := new(EVM2EVMOffRampHelperAdminSet)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "AdminSet", log); err != nil {
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
	StaticConfig  EVM2EVMOffRampStaticConfig
	DynamicConfig EVM2EVMOffRampDynamicConfig
	Raw           types.Log
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

type EVM2EVMOffRampHelperConfigSet0Iterator struct {
	Event *EVM2EVMOffRampHelperConfigSet0

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperConfigSet0Iterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperConfigSet0)
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
		it.Event = new(EVM2EVMOffRampHelperConfigSet0)
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

func (it *EVM2EVMOffRampHelperConfigSet0Iterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperConfigSet0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperConfigSet0 struct {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterConfigSet0(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperConfigSet0Iterator, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "ConfigSet0")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperConfigSet0Iterator{contract: _EVM2EVMOffRampHelper.contract, event: "ConfigSet0", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchConfigSet0(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperConfigSet0) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "ConfigSet0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperConfigSet0)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ConfigSet0", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseConfigSet0(log types.Log) (*EVM2EVMOffRampHelperConfigSet0, error) {
	event := new(EVM2EVMOffRampHelperConfigSet0)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "ConfigSet0", log); err != nil {
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
	ReturnData     []byte
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

type EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator struct {
	Event *EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight)
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
		it.Event = new(EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight)
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

func (it *EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight struct {
	Nonce  uint64
	Sender common.Address
	Raw    types.Log
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) FilterSkippedSenderWithPreviousRampMessageInflight(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.FilterLogs(opts, "SkippedSenderWithPreviousRampMessageInflight", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator{contract: _EVM2EVMOffRampHelper.contract, event: "SkippedSenderWithPreviousRampMessageInflight", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) WatchSkippedSenderWithPreviousRampMessageInflight(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight, nonce []uint64, sender []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMOffRampHelper.contract.WatchLogs(opts, "SkippedSenderWithPreviousRampMessageInflight", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight)
				if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "SkippedSenderWithPreviousRampMessageInflight", log); err != nil {
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperFilterer) ParseSkippedSenderWithPreviousRampMessageInflight(log types.Log) (*EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight, error) {
	event := new(EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight)
	if err := _EVM2EVMOffRampHelper.contract.UnpackLog(event, "SkippedSenderWithPreviousRampMessageInflight", log); err != nil {
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
	case _EVM2EVMOffRampHelper.abi.Events["AdminSet"].ID:
		return _EVM2EVMOffRampHelper.ParseAdminSet(log)
	case _EVM2EVMOffRampHelper.abi.Events["ConfigSet"].ID:
		return _EVM2EVMOffRampHelper.ParseConfigSet(log)
	case _EVM2EVMOffRampHelper.abi.Events["ConfigSet0"].ID:
		return _EVM2EVMOffRampHelper.ParseConfigSet0(log)
	case _EVM2EVMOffRampHelper.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMOffRampHelper.ParseExecutionStateChanged(log)
	case _EVM2EVMOffRampHelper.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMOffRampHelper.ParseOwnershipTransferRequested(log)
	case _EVM2EVMOffRampHelper.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMOffRampHelper.ParseOwnershipTransferred(log)
	case _EVM2EVMOffRampHelper.abi.Events["PoolAdded"].ID:
		return _EVM2EVMOffRampHelper.ParsePoolAdded(log)
	case _EVM2EVMOffRampHelper.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMOffRampHelper.ParsePoolRemoved(log)
	case _EVM2EVMOffRampHelper.abi.Events["SkippedIncorrectNonce"].ID:
		return _EVM2EVMOffRampHelper.ParseSkippedIncorrectNonce(log)
	case _EVM2EVMOffRampHelper.abi.Events["SkippedSenderWithPreviousRampMessageInflight"].ID:
		return _EVM2EVMOffRampHelper.ParseSkippedSenderWithPreviousRampMessageInflight(log)
	case _EVM2EVMOffRampHelper.abi.Events["Transmitted"].ID:
		return _EVM2EVMOffRampHelper.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMOffRampHelperAdminSet) Topic() common.Hash {
	return common.HexToHash("0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c")
}

func (EVM2EVMOffRampHelperConfigSet) Topic() common.Hash {
	return common.HexToHash("0xe668e1a4644c1a030b909bbfd837f5cfa914994ed5e0bb2e9c34a5c37753128a")
}

func (EVM2EVMOffRampHelperConfigSet0) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMOffRampHelperExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0xd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65")
}

func (EVM2EVMOffRampHelperOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMOffRampHelperOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
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

func (EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight) Topic() common.Hash {
	return common.HexToHash("0xe44a20935573a783dd0d5991c92d7b6a0eb3173566530364db3ec10e9a990b5d")
}

func (EVM2EVMOffRampHelperTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelper) Address() common.Address {
	return _EVM2EVMOffRampHelper.address
}

type EVM2EVMOffRampHelperInterface interface {
	ISSCRIPT(opts *bind.CallOpts) (bool, error)

	CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error

	CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMOffRampDynamicConfig, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetExecutionStateBitMap(opts *bind.CallOpts, bitmapIndex uint64) (*big.Int, error)

	GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error)

	GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMOffRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	MetadataHash(opts *bind.CallOpts) ([32]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyPoolUpdates(opts *bind.TransactOpts, removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, rep InternalExecutionReport, manualExecGasLimits []*big.Int) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report InternalExecutionReport, gasLimitOverrides []*big.Int) (*types.Transaction, error)

	ReleaseOrMintTokens(opts *bind.TransactOpts, sourceTokenAmounts []ClientEVMTokenAmount, originalSender []byte, receiver common.Address, sourceTokenData [][]byte, offchainTokenData [][]byte) (*types.Transaction, error)

	Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error)

	SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetExecutionStateHelper(opts *bind.TransactOpts, sequenceNumber uint64, state uint8) (*types.Transaction, error)

	SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error)

	TrialExecute(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error)

	FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperAdminSetIterator, error)

	WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperAdminSet) (event.Subscription, error)

	ParseAdminSet(log types.Log) (*EVM2EVMOffRampHelperAdminSet, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMOffRampHelperConfigSet, error)

	FilterConfigSet0(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperConfigSet0Iterator, error)

	WatchConfigSet0(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperConfigSet0) (event.Subscription, error)

	ParseConfigSet0(log types.Log) (*EVM2EVMOffRampHelperConfigSet0, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64, messageId [][32]byte) (*EVM2EVMOffRampHelperExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperExecutionStateChanged, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMOffRampHelperExecutionStateChanged, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampHelperOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMOffRampHelperOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampHelperOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMOffRampHelperOwnershipTransferred, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMOffRampHelperPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMOffRampHelperPoolRemoved, error)

	FilterSkippedIncorrectNonce(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampHelperSkippedIncorrectNonceIterator, error)

	WatchSkippedIncorrectNonce(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperSkippedIncorrectNonce, nonce []uint64, sender []common.Address) (event.Subscription, error)

	ParseSkippedIncorrectNonce(log types.Log) (*EVM2EVMOffRampHelperSkippedIncorrectNonce, error)

	FilterSkippedSenderWithPreviousRampMessageInflight(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflightIterator, error)

	WatchSkippedSenderWithPreviousRampMessageInflight(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight, nonce []uint64, sender []common.Address) (event.Subscription, error)

	ParseSkippedSenderWithPreviousRampMessageInflight(log types.Log) (*EVM2EVMOffRampHelperSkippedSenderWithPreviousRampMessageInflight, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMOffRampHelperTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
