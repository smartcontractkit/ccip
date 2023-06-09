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
	Arm                                     common.Address
	MaxTokensLength                         uint16
	MaxDataSize                             uint32
}

type EVM2EVMOffRampStaticConfig struct {
	CommitStore         common.Address
	ChainSelector       uint64
	SourceChainSelector uint64
	OnRamp              common.Address
	PrevOffRamp         common.Address
}

type InternalEVM2EVMMessage struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	FeeTokenAmount      *big.Int
	Sender              common.Address
	Nonce               uint64
	GasLimit            *big.Int
	Strict              bool
	Receiver            common.Address
	Data                []byte
	TokenAmounts        []ClientEVMTokenAmount
	FeeToken            common.Address
	MessageId           [32]byte
}

type InternalExecutionReport struct {
	SequenceNumbers   []uint64
	EncodedMessages   [][]byte
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

var EVM2EVMOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadARMSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitStoreAlreadyInUse\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ConsumingMoreThanMaxCapacity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMessageId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"InvalidNewState\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"arm\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOffRamp.DynamicConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"InvalidOffRampConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"RateLimitReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"TokenDataMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedTokenData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"arm\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedSenderWithPreviousRampMessageInflight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"}],\"name\":\"applyPoolUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"destTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"arm\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOffRamp.DynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolByDestToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isARMHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b50604051620066a9380380620066a983398101604081905262000035916200087c565b8033806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c08162000464565b50506040805160a081018252602084810180516001600160801b039081168085524263ffffffff169385018490528751151585870181905292518216606086018190529790950151166080909301839052600380546001600160a01b031916909417600160801b9283021760ff60a01b1916600160a01b9091021790925502909117600455508151835114620001685760405162d8548360e71b815260040160405180910390fd5b60608401516001600160a01b031615806200018b575083516001600160a01b0316155b15620001aa576040516342bcdf7f60e11b815260040160405180910390fd5b83600001516001600160a01b0316634120fccd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200021391906200097d565b6001600160401b03166001146200023d57604051636fc2a20760e11b815260040160405180910390fd5b83516001600160a01b03908116608090815260408601516001600160401b0390811660a05260208701511660c0526060860151821660e0528501511661012052620002a87fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a6200050f565b6101005260005b8351811015620004595762000309848281518110620002d257620002d26200099b565b6020026020010151848381518110620002ef57620002ef6200099b565b6020026020010151600d6200057560201b9092919060201c565b50620003be8382815181106200032357620003236200099b565b60200260200101516001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000369573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200038f9190620009b1565b848381518110620003a457620003a46200099b565b602002602001015160106200057560201b9092919060201c565b507f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c848281518110620003f557620003f56200099b565b60200260200101518483815181106200041257620004126200099b565b60200260200101516040516200043e9291906001600160a01b0392831681529116602082015260400190565b60405180910390a16200045181620009d8565b9050620002af565b505050505062000a00565b336001600160a01b03821603620004be5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008160a05160c05160e0516040516020016200055894939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b60006200058d846001600160a01b0385168462000595565b949350505050565b60006200058d84846001600160a01b038516600082815260028401602052604081208290556200058d84846000620005ce8383620005d7565b90505b92915050565b60008181526001830160205260408120546200062057508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620005d1565b506000620005d1565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b038111828210171562000664576200066462000629565b60405290565b604051601f8201601f191681016001600160401b038111828210171562000695576200069562000629565b604052919050565b6001600160a01b0381168114620006b357600080fd5b50565b80516001600160401b0381168114620006ce57600080fd5b919050565b60006001600160401b03821115620006ef57620006ef62000629565b5060051b60200190565b600082601f8301126200070b57600080fd5b81516020620007246200071e83620006d3565b6200066a565b82815260059290921b840181019181810190868411156200074457600080fd5b8286015b848110156200076c5780516200075e816200069d565b835291830191830162000748565b509695505050505050565b600082601f8301126200078957600080fd5b815160206200079c6200071e83620006d3565b82815260059290921b84018101918181019086841115620007bc57600080fd5b8286015b848110156200076c578051620007d6816200069d565b8352918301918301620007c0565b80516001600160801b0381168114620006ce57600080fd5b6000606082840312156200080f57600080fd5b604051606081016001600160401b038111828210171562000834576200083462000629565b8060405250809150825180151581146200084d57600080fd5b81526200085d60208401620007e4565b60208201526200087060408401620007e4565b60408201525092915050565b6000806000808486036101408112156200089557600080fd5b60a0811215620008a457600080fd5b50620008af6200063f565b8551620008bc816200069d565b8152620008cc60208701620006b6565b6020820152620008df60408701620006b6565b60408201526060860151620008f4816200069d565b6060820152608086015162000909816200069d565b608082015260a08601519094506001600160401b03808211156200092c57600080fd5b6200093a88838901620006f9565b945060c08701519150808211156200095157600080fd5b50620009608782880162000777565b925050620009728660e08701620007fc565b905092959194509250565b6000602082840312156200099057600080fd5b620005ce82620006b6565b634e487b7160e01b600052603260045260246000fd5b600060208284031215620009c457600080fd5b8151620009d1816200069d565b9392505050565b600060018201620009f957634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e0516101005161012051615c0562000aa4600039600081816102d1015281816117180152818161179a01528181611f4601528181612b6c0152612bf3015260006127db0152600081816102950152611f1f0152600081816102350152611ecd01526000818161026501528181611ef7015281816121b701526137d60152600081816101f901528181611e9f01526128ef0152615c056000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c80637437ff9f116100f9578063afcb95d711610097578063c92b283211610071578063c92b283214610698578063d3c7c2c7146106ab578063d7e2bb50146106b3578063f2fde38b146106c657600080fd5b8063afcb95d714610652578063b1dc65a414610672578063b4069b311461068557600080fd5b806385572ffb116100d357806385572ffb146105e7578063856c8247146105f55780638da5cb5b146106215780638fa4f53b1461063f57600080fd5b80637437ff9f146104d357806379ba5097146105af57806381ff7048146105b757600080fd5b80634a11d44a116101665780635d86f141116101405780635d86f14114610483578063666cab8d14610496578063681fba16146104ab578063704b6c02146104c057600080fd5b80634a11d44a146103cd578063546719cd146103e0578063599f64311461044457600080fd5b8063181f5a7711610197578063181f5a771461035c5780631ef38174146103a55780633a87ac53146103ba57600080fd5b806306285c69146101be5780630a6cd30d14610324578063142a98fc1461033c575b600080fd5b61030e6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040518060a001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815250905090565b60405161031b9190614213565b60405180910390f35b61032c6106d9565b604051901515815260200161031b565b61034f61034a36600461429e565b610773565b60405161031b9190614325565b6103986040518060400160405280601481526020017f45564d3245564d4f666652616d7020312e302e3000000000000000000000000081525081565b60405161031b91906143a1565b6103b86103b336600461460f565b6107ee565b005b6103b86103c8366004614721565b610bec565b6103b86103db3660046148ae565b610ffc565b6103e86112aa565b60405161031b919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161031b565b61045e610491366004614a13565b61135f565b61049e6113c8565b60405161031b9190614a81565b6104b3611437565b60405161031b9190614a94565b6103b86104ce366004614a13565b6114f0565b6105a26040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c081018252600a5463ffffffff808216835273ffffffffffffffffffffffffffffffffffffffff64010000000090920482166020840152600b54821693830193909352600c54908116606083015261ffff740100000000000000000000000000000000000000008204166080830152760100000000000000000000000000000000000000000000900490911660a082015290565b60405161031b9190614aee565b6103b86115e0565b6007546005546040805163ffffffff8085168252640100000000909404909316602084015282015260600161031b565b6103b86101b9366004614b75565b610608610603366004614a13565b6116dd565b60405167ffffffffffffffff909116815260200161031b565b60005473ffffffffffffffffffffffffffffffffffffffff1661045e565b6103b861064d366004614b75565b611805565b60408051600181526000602082018190529181019190915260600161031b565b6103b8610680366004614bef565b61181b565b61045e610693366004614a13565b611a2b565b6103b86106a6366004614cf4565b611aec565b6104b3611b6e565b61045e6106c1366004614a13565b611c23565b6103b86106d4366004614a13565b611c32565b600c54604080517f397796f7000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163397796f79160048083019260209291908290030181865afa158015610749573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076d9190614d6d565b15905090565b600061078160016004614db9565b600261078e608085614dfb565b67ffffffffffffffff166107a29190614e22565b601460006107b1608087614e39565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002054901c1660038111156107e8576107e86142bb565b92915050565b8360ff1680600003610861576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064015b60405180910390fd5b610869611c43565b61087284611cc6565b60095460005b818110156108fe57600860006009838154811061089757610897614e60565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001690556108f781614e8f565b9050610878565b50865160005b81811015610aad57600089828151811061092057610920614e60565b602002602001015190506000600281111561093d5761093d6142bb565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260086020526040902054610100900460ff16600281111561097c5761097c6142bb565b146109e3576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610858565b6040805180820190915260ff83168152602081016002905273ffffffffffffffffffffffffffffffffffffffff821660009081526008602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610a9357610a936142bb565b02179055509050505080610aa690614e8f565b9050610904565b508751610ac19060099060208b019061417d565b506006805460ff838116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909216908a161717905560078054610b47914691309190600090610b199063ffffffff16614ec7565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168c8c8c8c8c8c611fa9565b6005819055600780544363ffffffff9081166401000000009081027fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff841681179094556040519083048216947f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0594610bd89487949293918316921691909117908f908f908f908f908f908f90614eea565b60405180910390a150505050505050505050565b610bf4611c43565b60005b83811015610df3576000858583818110610c1357610c13614e60565b610c299260206040909202019081019150614a13565b90506000868684818110610c3f57610c3f614e60565b9050604002016020016020810190610c579190614a13565b9050610c64600d83612054565b610c9a576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610cbc600d84612076565b73ffffffffffffffffffffffffffffffffffffffff1614610d09576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d14600d83612098565b50610d8f8173ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d879190614f80565b601090612098565b506040805173ffffffffffffffffffffffffffffffffffffffff8085168252831660208201527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a1505080610dec90614e8f565b9050610bf7565b5060005b81811015610ff5576000838383818110610e1357610e13614e60565b610e299260206040909202019081019150614a13565b90506000848484818110610e3f57610e3f614e60565b9050604002016020016020810190610e579190614a13565b905073ffffffffffffffffffffffffffffffffffffffff82161580610e90575073ffffffffffffffffffffffffffffffffffffffff8116155b15610ec7576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ed2600d83612054565b15610f09576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f15600d83836120ba565b50610f918173ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f889190614f80565b601090836120ba565b506040805173ffffffffffffffffffffffffffffffffffffffff8085168252831660208201527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a1505080610fee90614e8f565b9050610df7565b5050505050565b333014611035576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160008082526020820190925281611072565b604080518082019091526000808252602082015281526020019060019003908161104b5790505b5061012085015151909150156110d15761012084015160608501516040805173ffffffffffffffffffffffffffffffffffffffff90921660208301526110ce9291016040516020818303038152906040528660e00151866120dd565b90505b60e084015173ffffffffffffffffffffffffffffffffffffffff163b158061113b575060e08401516111399073ffffffffffffffffffffffffffffffffffffffff167f85572ffb00000000000000000000000000000000000000000000000000000000612367565b155b156111465750505050565b60a084015182156111a257610100850151516040906113889061116a906010614e22565b6111749190614f9d565b61117f906002614e22565b5a61118a9190614db9565b61119590603e614e22565b61119f9190614fb0565b90505b600a546000908190640100000000900473ffffffffffffffffffffffffffffffffffffffff16633cf979836111d78987612383565b611388868b60e001516040518563ffffffff1660e01b81526004016111ff9493929190615015565b6000604051808303816000875af115801561121e573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611264919081019061512c565b91509150816112a157806040517f0a8d6e8c00000000000000000000000000000000000000000000000000000000815260040161085891906143a1565b50505050505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526003546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff1660208501527401000000000000000000000000000000000000000090920460ff16151593830193909352600454808416606084015204909116608082015261135a90612433565b905090565b6000808061136e600d856124e5565b91509150816113c1576040517fbf16aab600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610858565b9392505050565b6060600980548060200260200160405190810160405280929190818152602001828054801561142d57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611402575b5050505050905090565b60606114436010612514565b67ffffffffffffffff81111561145b5761145b6143b4565b604051908082528060200260200182016040528015611484578160200160208202803683370190505b50905060005b81518110156114ec5760006114a060108361251f565b509050808383815181106114b6576114b6614e60565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152506114e581614e8f565b905061148a565b5090565b60005473ffffffffffffffffffffffffffffffffffffffff163314801590611530575060025473ffffffffffffffffffffffffffffffffffffffff163314155b15611567576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c9060200160405180910390a150565b60015473ffffffffffffffffffffffffffffffffffffffff163314611661576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610858565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b73ffffffffffffffffffffffffffffffffffffffff811660009081526013602052604081205467ffffffffffffffff168015801561175057507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1615155b156107e8576040517f856c824700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063856c824790602401602060405180830381865afa1580156117e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113c19190615188565b611818611811826153a7565b600161253b565b50565b6118258787613029565b60055488359080821461186e576040517f93df584c0000000000000000000000000000000000000000000000000000000081526004810182905260248101839052604401610858565b6040805183815260208c81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a13360009081526008602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156118f6576118f66142bb565b6002811115611907576119076142bb565b9052509050600281602001516002811115611924576119246142bb565b14801561196b57506009816000015160ff168154811061194657611946614e60565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b6119a1576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060006119af856020614e22565b6119ba886020614e22565b6119c68b610144614f9d565b6119d09190614f9d565b6119da9190614f9d565b9050368114611a1e576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401610858565b5050505050505050505050565b60008080611a3a600d856124e5565b9150915081611a75576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ac0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ae49190614f80565b949350505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314801590611b2c575060025473ffffffffffffffffffffffffffffffffffffffff163314155b15611b63576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611818600382613043565b6060611b7a600d612514565b67ffffffffffffffff811115611b9257611b926143b4565b604051908082528060200260200182016040528015611bbb578160200160208202803683370190505b50905060005b81518110156114ec576000611bd7600d8361251f565b50905080838381518110611bed57611bed614e60565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015250611c1c81614e8f565b9050611bc1565b6000808061136e6010856124e5565b611c3a611c43565b61181881613228565b60005473ffffffffffffffffffffffffffffffffffffffff163314611cc4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610858565b565b600081806020019051810190611cdc91906153d2565b602081015190915073ffffffffffffffffffffffffffffffffffffffff161580611d1e5750606081015173ffffffffffffffffffffffffffffffffffffffff16155b15611d5757806040517f9c1779a70000000000000000000000000000000000000000000000000000000081526004016108589190614aee565b8051600a805460208085015173ffffffffffffffffffffffffffffffffffffffff908116640100000000027fffffffffffffffff00000000000000000000000000000000000000000000000090931663ffffffff9586161792909217909255604080850151600b80549184167fffffffffffffffffffffffff0000000000000000000000000000000000000000909216919091179055606080860151600c80546080808a015160a0808c0151909a16760100000000000000000000000000000000000000000000027fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff61ffff90921674010000000000000000000000000000000000000000027fffffffffffffffffffff00000000000000000000000000000000000000000000909416958916959095179290921791909116929092179055825195860183527f00000000000000000000000000000000000000000000000000000000000000008416865267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116958701959095527f0000000000000000000000000000000000000000000000000000000000000000909416858301527f00000000000000000000000000000000000000000000000000000000000000008316908501527f000000000000000000000000000000000000000000000000000000000000000090911691830191909152517fccb0480fd9368458c55b5c2e24b62249e1f5eabb55ec2e42d37965a3ae1eaf8791611f9d918490615480565b60405180910390a15050565b6000808a8a8a8a8a8a8a8a8a604051602001611fcd99989796959493929190615551565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b60006113c18373ffffffffffffffffffffffffffffffffffffffff841661331d565b60006113c18373ffffffffffffffffffffffffffffffffffffffff8416613329565b60006113c18373ffffffffffffffffffffffffffffffffffffffff8416613335565b6000611ae48473ffffffffffffffffffffffffffffffffffffffff851684613341565b60606000855167ffffffffffffffff8111156120fb576120fb6143b4565b60405190808252806020026020018201604052801561214057816020015b60408051808201909152600080825260208201528152602001906001900390816121195790505b50905060005b865181101561233957600061217788838151811061216657612166614e60565b60200260200101516000015161135f565b90508073ffffffffffffffffffffffffffffffffffffffff16638627fad688888b86815181106121a9576121a9614e60565b6020026020010151602001517f00000000000000000000000000000000000000000000000000000000000000008a88815181106121e8576121e8614e60565b60200260200101516040518663ffffffff1660e01b81526004016122109594939291906155e6565b600060405180830381600087803b15801561222a57600080fd5b505af115801561223e573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa15801561228d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122b19190614f80565b8383815181106122c3576122c3614e60565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff909116905287518890839081106122fc576122fc614e60565b60200260200101516020015183838151811061231a5761231a614e60565b60209081029190910181015101525061233281614e8f565b9050612146565b50600b5461235e90829073ffffffffffffffffffffffffffffffffffffffff16613364565b95945050505050565b600061237283613542565b80156113c157506113c183836135a6565b6040805160a08101825260008082526020820152606091810182905281810182905260808101919091526040518060a001604052808461016001518152602001846000015167ffffffffffffffff1681526020018460600151604051602001612408919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b6040516020818303038152906040528152602001846101000151815260200183815250905092915050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526124c182606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff16426124a59190614db9565b85608001516fffffffffffffffffffffffffffffffff16613675565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b6000806125088473ffffffffffffffffffffffffffffffffffffffff8516613694565b915091505b9250929050565b60006107e8826136a3565b600080808061252e86866136ae565b9097909650945050505050565b600c54604080517f397796f7000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9092169163397796f7916004808201926020929091908290030181865afa1580156125ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125cf9190614d6d565b15612606576040517fc148371500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020820151516000819003612646576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8260400151518114612684576040517f57e0e08300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff81111561269f5761269f6143b4565b6040519080825280602002602001820160405280156126c8578160200160208202803683370190505b50905060008267ffffffffffffffff8111156126e6576126e66143b4565b60405190808252806020026020018201604052801561279257816020015b60408051610180810182526000808252602080830182905292820181905260608083018290526080830182905260a0830182905260c0830182905260e083018290526101008301819052610120830152610140820181905261016082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092019101816127045790505b50905060005b838110156128a4576000866020015182815181106127b8576127b8614e60565b60200260200101518060200190518101906127d391906156d1565b90506127ff817f00000000000000000000000000000000000000000000000000000000000000006136bd565b84838151811061281157612811614e60565b60200260200101818152505080610160015184838151811061283557612835614e60565b602002602001015114612874576040517f7185cf6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8083838151811061288757612887614e60565b6020026020010181905250508061289d90614e8f565b9050612798565b50606085015160808601516040517f3204887500000000000000000000000000000000000000000000000000000000815260009273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016926332048875926129269288929091600401615834565b6020604051808303816000875af1158015612945573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612969919061586a565b9050806000036129a5576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b848110156112a15760008382815181106129c4576129c4614e60565b6020026020010151905060006129dd8260200151610773565b905060008160038111156129f3576129f36142bb565b1480612a1057506003816003811115612a0e57612a0e6142bb565b145b612a585760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610858565b8715612ad057600a5460009063ffffffff16612a748642614db9565b1190508080612a9457506003826003811115612a9257612a926142bb565b145b612aca576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50612b2d565b6000816003811115612ae457612ae46142bb565b14612b2d5760208201516040517f67d9ba0f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610858565b606082015173ffffffffffffffffffffffffffffffffffffffff1660009081526013602052604090205467ffffffffffffffff1680158015612ba457507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1615155b15612d465760608301516040517f856c824700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201527f00000000000000000000000000000000000000000000000000000000000000009091169063856c824790602401602060405180830381865afa158015612c3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c609190615188565b608084015190915067ffffffffffffffff16612c7d826001615883565b67ffffffffffffffff1614612cea57826060015173ffffffffffffffffffffffffffffffffffffffff16836080015167ffffffffffffffff167fe44a20935573a783dd0d5991c92d7b6a0eb3173566530364db3ec10e9a990b5d60405160405180910390a3505050613019565b606083015173ffffffffffffffffffffffffffffffffffffffff16600090815260136020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83161790555b6000826003811115612d5a57612d5a6142bb565b03612de657608083015167ffffffffffffffff16612d79826001615883565b67ffffffffffffffff1614612de657826060015173ffffffffffffffffffffffffffffffffffffffff16836080015167ffffffffffffffff167fd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf4123760405160405180910390a3505050613019565b60008a604001518581518110612dfe57612dfe614e60565b60200260200101519050612e138482516137d4565b612e22846020015160016139aa565b600080612e3086848e613a54565b91509150612e428660200151836139aa565b6003826003811115612e5657612e566142bb565b14158015612e7657506002826003811115612e7357612e736142bb565b14155b15612eb5578560200151826040517f9e2616030000000000000000000000000000000000000000000000000000000081526004016108589291906158a4565b8560c0015115612f44576002826003811115612ed357612ed36142bb565b03612f3f57606086015173ffffffffffffffffffffffffffffffffffffffff166000908152601360205260408120805467ffffffffffffffff1691612f17836158c2565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b612fc4565b6000856003811115612f5857612f586142bb565b03612fc457606086015173ffffffffffffffffffffffffffffffffffffffff166000908152601360205260408120805467ffffffffffffffff1691612f9c836158c2565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b856101600151866020015167ffffffffffffffff167fd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65848460405161300a9291906158df565b60405180910390a35050505050505b61302281614e8f565b90506129a8565b61303f613038828401846158ff565b600061253b565b5050565b815460009061306c90700100000000000000000000000000000000900463ffffffff1642614db9565b9050801561310e57600183015483546130b4916fffffffffffffffffffffffffffffffff80821692811691859170010000000000000000000000000000000090910416613675565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354613134916fffffffffffffffffffffffffffffffff9081169116613ba6565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c199061321b9084908151151581526020808301516fffffffffffffffffffffffffffffffff90811691830191909152604092830151169181019190915260600190565b60405180910390a1505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036132a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610858565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006113c18383613bbc565b60006113c18383613bc8565b60006113c18383613c52565b6000611ae4848473ffffffffffffffffffffffffffffffffffffffff8516613c6f565b81516000805b828110156135305760008473ffffffffffffffffffffffffffffffffffffffff1663d02641a08784815181106133a2576133a2614e60565b6020908102919091010151516040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016040805180830381865afa158015613416573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061343a9190615934565b51905077ffffffffffffffffffffffffffffffffffffffffffffffff81166000036134c85785828151811061347157613471614e60565b6020908102919091010151516040517f9a655f7b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610858565b6135128683815181106134dd576134dd614e60565b6020026020010151602001518277ffffffffffffffffffffffffffffffffffffffffffffffff16613c8c90919063ffffffff16565b61351c9084614f9d565b9250508061352990614e8f565b905061336a565b5061353c600382613cc5565b50505050565b600061356e827f01ffc9a7000000000000000000000000000000000000000000000000000000006135a6565b80156107e8575061359f827fffffffff000000000000000000000000000000000000000000000000000000006135a6565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d9150600051905082801561365e575060208210155b801561366a5750600081115b979650505050505050565b600061235e856136858486614e22565b61368f9087614f9d565b613ba6565b600080808061252e8686613f54565b60006107e882613f8e565b600080808061252e8686613f99565b60008060001b828460200151856080015186606001518760e00151886101000151805190602001208961012001516040516020016136fb9190615992565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d604001516040516020016137b69c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a015273ffffffffffffffffffffffffffffffffffffffff94851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16826000015167ffffffffffffffff16146138545781516040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610858565b600c54610120830151517401000000000000000000000000000000000000000090910461ffff1610156138c55760208201516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610858565b8082610120015151146139165760208201516040517f8808f8e700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610858565b600c546101008301515176010000000000000000000000000000000000000000000090910463ffffffff16101561303f57600c54610100830151516040517f8693378900000000000000000000000000000000000000000000000000000000815276010000000000000000000000000000000000000000000090920463ffffffff1660048301526024820152604401610858565b600060026139b9608085614dfb565b67ffffffffffffffff166139cd9190614e22565b905060006014816139df608087614e39565b67ffffffffffffffff168152602081019190915260400160002054905081613a0960016004614db9565b901b191681836003811115613a2057613a206142bb565b901b178060146000613a33608088614e39565b67ffffffffffffffff16815260208101919091526040016000205550505050565b6040517f4a11d44a0000000000000000000000000000000000000000000000000000000081526000906060903090634a11d44a90613a9a908890889088906004016159fa565b600060405180830381600087803b158015613ab457600080fd5b505af1925050508015613ac5575060015b613b89573d808015613af3576040519150601f19603f3d011682016040523d82523d6000602084013e613af8565b606091505b50613b0281615b79565b7fffffffff00000000000000000000000000000000000000000000000000000000167f0a8d6e8c0000000000000000000000000000000000000000000000000000000003613b5557600392509050613b9e565b806040517fcf19edfd00000000000000000000000000000000000000000000000000000000815260040161085891906143a1565b50506040805160208101909152600081526002905b935093915050565b6000818310613bb557816113c1565b5090919050565b60006113c18383613fc4565b600081815260028301602052604081205480151580613bec5750613bec8484613bbc565b6113c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b657900006044820152606401610858565b600081815260028301602052604081208190556113c18383613fdc565b60008281526002840160205260408120829055611ae48484613fe8565b6000670de0b6b3a7640000613cbb8377ffffffffffffffffffffffffffffffffffffffffffffffff8616614e22565b6113c19190614fb0565b815474010000000000000000000000000000000000000000900460ff161580613cec575080155b15613cf5575050565b815460018301546fffffffffffffffffffffffffffffffff80831692911690600090613d3b90700100000000000000000000000000000000900463ffffffff1642614db9565b90508015613dfb5781831115613d7d576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001850154613db79083908590849070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16613675565b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000004263ffffffff160217865592505b83821015613e3f576040517f48369c430000000000000000000000000000000000000000000000000000000081526004810183905260248101859052604401610858565b83831015613ed35760018581015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16908190613e809082614db9565b613e8a8688614db9565b613e949190614f9d565b613e9e9190614fb0565b6040517fdc96cefa00000000000000000000000000000000000000000000000000000000815260040161085891815260200190565b613edd8484614db9565b85547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161786556040518581529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a15050505050565b6000818152600283016020526040812054819080613f8357613f768585613bbc565b92506000915061250d9050565b60019250905061250d565b60006107e882613ff4565b60008080613fa78585613ffe565b600081815260029690960160205260409095205494959350505050565b600081815260018301602052604081205415156113c1565b60006113c1838361400a565b60006113c18383614104565b60006107e8825490565b60006113c18383614153565b600081815260018301602052604081205480156140f357600061402e600183614db9565b855490915060009061404290600190614db9565b90508181146140a757600086600001828154811061406257614062614e60565b906000526020600020015490508087600001848154811061408557614085614e60565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806140b8576140b8615bc9565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506107e8565b60009150506107e8565b5092915050565b600081815260018301602052604081205461414b575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556107e8565b5060006107e8565b600082600001828154811061416a5761416a614e60565b9060005260206000200154905092915050565b8280548282559060005260206000209081019282156141f7579160200282015b828111156141f757825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90911617825560209092019160019091019061419d565b506114ec9291505b808211156114ec57600081556001016141ff565b60a081016107e8828473ffffffffffffffffffffffffffffffffffffffff808251168352602082015167ffffffffffffffff80821660208601528060408501511660408601525050806060830151166060840152806080830151166080840152505050565b67ffffffffffffffff8116811461181857600080fd5b803561429981614278565b919050565b6000602082840312156142b057600080fd5b81356113c181614278565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110614321577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b602081016107e882846142ea565b60005b8381101561434e578181015183820152602001614336565b50506000910152565b6000815180845261436f816020860160208601614333565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006113c16020830184614357565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715614406576144066143b4565b60405290565b604051610180810167ffffffffffffffff81118282101715614406576144066143b4565b60405160a0810167ffffffffffffffff81118282101715614406576144066143b4565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561449a5761449a6143b4565b604052919050565b600067ffffffffffffffff8211156144bc576144bc6143b4565b5060051b60200190565b73ffffffffffffffffffffffffffffffffffffffff8116811461181857600080fd5b8035614299816144c6565b600082601f83011261450457600080fd5b81356020614519614514836144a2565b614453565b82815260059290921b8401810191818101908684111561453857600080fd5b8286015b8481101561455c57803561454f816144c6565b835291830191830161453c565b509695505050505050565b803560ff8116811461429957600080fd5b600067ffffffffffffffff821115614592576145926143b4565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f8301126145cf57600080fd5b81356145dd61451482614578565b8181528460208386010111156145f257600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c0878903121561462857600080fd5b863567ffffffffffffffff8082111561464057600080fd5b61464c8a838b016144f3565b9750602089013591508082111561466257600080fd5b61466e8a838b016144f3565b965061467c60408a01614567565b9550606089013591508082111561469257600080fd5b61469e8a838b016145be565b94506146ac60808a0161428e565b935060a08901359150808211156146c257600080fd5b506146cf89828a016145be565b9150509295509295509295565b60008083601f8401126146ee57600080fd5b50813567ffffffffffffffff81111561470657600080fd5b6020830191508360208260061b850101111561250d57600080fd5b6000806000806040858703121561473757600080fd5b843567ffffffffffffffff8082111561474f57600080fd5b61475b888389016146dc565b9096509450602087013591508082111561477457600080fd5b50614781878288016146dc565b95989497509550505050565b801515811461181857600080fd5b80356142998161478d565b600082601f8301126147b757600080fd5b813560206147c7614514836144a2565b82815260069290921b840181019181810190868411156147e657600080fd5b8286015b8481101561455c57604081890312156148035760008081fd5b61480b6143e3565b8135614816816144c6565b815281850135858201528352918301916040016147ea565b600082601f83011261483f57600080fd5b8135602061484f614514836144a2565b82815260059290921b8401810191818101908684111561486e57600080fd5b8286015b8481101561455c57803567ffffffffffffffff8111156148925760008081fd5b6148a08986838b01016145be565b845250918301918301614872565b6000806000606084860312156148c357600080fd5b833567ffffffffffffffff808211156148db57600080fd5b9085019061018082880312156148f057600080fd5b6148f861440c565b6149018361428e565b815261490f6020840161428e565b60208201526040830135604082015261492a606084016144e8565b606082015261493b6080840161428e565b608082015260a083013560a082015261495660c0840161479b565b60c082015261496760e084016144e8565b60e0820152610100808401358381111561498057600080fd5b61498c8a8287016145be565b82840152505061012080840135838111156149a657600080fd5b6149b28a8287016147a6565b8284015250506101406149c68185016144e8565b908201526101609283013592810192909252909350602085013590808211156149ee57600080fd5b506149fb8682870161482e565b925050614a0a6040850161479b565b90509250925092565b600060208284031215614a2557600080fd5b81356113c1816144c6565b600081518084526020808501945080840160005b83811015614a7657815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101614a44565b509495945050505050565b6020815260006113c16020830184614a30565b6020808252825182820181905260009190848201906040850190845b81811015614ae257835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101614ab0565b50909695505050505050565b60c081016107e8828463ffffffff808251168352602082015173ffffffffffffffffffffffffffffffffffffffff8082166020860152806040850151166040860152806060850151166060860152505061ffff60808301511660808401528060a08301511660a0840152505050565b600060a08284031215614b6f57600080fd5b50919050565b600060208284031215614b8757600080fd5b813567ffffffffffffffff811115614b9e57600080fd5b611ae484828501614b5d565b60008083601f840112614bbc57600080fd5b50813567ffffffffffffffff811115614bd457600080fd5b6020830191508360208260051b850101111561250d57600080fd5b60008060008060008060008060e0898b031215614c0b57600080fd5b606089018a811115614c1c57600080fd5b8998503567ffffffffffffffff80821115614c3657600080fd5b818b0191508b601f830112614c4a57600080fd5b813581811115614c5957600080fd5b8c6020828501011115614c6b57600080fd5b6020830199508098505060808b0135915080821115614c8957600080fd5b614c958c838d01614baa565b909750955060a08b0135915080821115614cae57600080fd5b50614cbb8b828c01614baa565b999c989b50969995989497949560c00135949350505050565b80356fffffffffffffffffffffffffffffffff8116811461429957600080fd5b600060608284031215614d0657600080fd5b6040516060810181811067ffffffffffffffff82111715614d2957614d296143b4565b6040528235614d378161478d565b8152614d4560208401614cd4565b6020820152614d5660408401614cd4565b60408201529392505050565b80516142998161478d565b600060208284031215614d7f57600080fd5b81516113c18161478d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156107e8576107e8614d8a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff80841680614e1657614e16614dcc565b92169190910692915050565b80820281158282048414176107e8576107e8614d8a565b600067ffffffffffffffff80841680614e5457614e54614dcc565b92169190910492915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614ec057614ec0614d8a565b5060010190565b600063ffffffff808316818103614ee057614ee0614d8a565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614f1a8184018a614a30565b90508281036080840152614f2e8189614a30565b905060ff871660a084015282810360c0840152614f4b8187614357565b905067ffffffffffffffff851660e0840152828103610100840152614f708185614357565b9c9b505050505050505050505050565b600060208284031215614f9257600080fd5b81516113c1816144c6565b808201808211156107e8576107e8614d8a565b600082614fbf57614fbf614dcc565b500490565b600081518084526020808501945080840160005b83811015614a76578151805173ffffffffffffffffffffffffffffffffffffffff1688528301518388015260409096019590820190600101614fd8565b608081528451608082015267ffffffffffffffff60208601511660a08201526000604086015160a060c0840152615050610120840182614357565b905060608701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80808584030160e086015261508c8383614357565b9250608089015191508085840301610100860152506150ab8282614fc4565b925050506150bf602083018661ffff169052565b83604083015261235e606083018473ffffffffffffffffffffffffffffffffffffffff169052565b600082601f8301126150f857600080fd5b815161510661451482614578565b81815284602083860101111561511b57600080fd5b611ae4826020830160208701614333565b6000806040838503121561513f57600080fd5b825161514a8161478d565b602084015190925067ffffffffffffffff81111561516757600080fd5b615173858286016150e7565b9150509250929050565b805161429981614278565b60006020828403121561519a57600080fd5b81516113c181614278565b600082601f8301126151b657600080fd5b813560206151c6614514836144a2565b82815260059290921b840181019181810190868411156151e557600080fd5b8286015b8481101561455c5780356151fc81614278565b83529183019183016151e9565b600082601f83011261521a57600080fd5b8135602061522a614514836144a2565b82815260059290921b8401810191818101908684111561524957600080fd5b8286015b8481101561455c57803567ffffffffffffffff81111561526d5760008081fd5b61527b8986838b010161482e565b84525091830191830161524d565b600082601f83011261529a57600080fd5b813560206152aa614514836144a2565b82815260059290921b840181019181810190868411156152c957600080fd5b8286015b8481101561455c57803583529183019183016152cd565b600060a082840312156152f657600080fd5b6152fe614430565b9050813567ffffffffffffffff8082111561531857600080fd5b615324858386016151a5565b8352602084013591508082111561533a57600080fd5b6153468583860161482e565b6020840152604084013591508082111561535f57600080fd5b61536b85838601615209565b6040840152606084013591508082111561538457600080fd5b5061539184828501615289565b6060830152506080820135608082015292915050565b60006107e836836152e4565b805163ffffffff8116811461429957600080fd5b8051614299816144c6565b600060c082840312156153e457600080fd5b60405160c0810181811067ffffffffffffffff82111715615407576154076143b4565b604052615413836153b3565b81526020830151615423816144c6565b60208201526040830151615436816144c6565b60408201526060830151615449816144c6565b6060820152608083015161ffff8116811461546357600080fd5b608082015261547460a084016153b3565b60a08201529392505050565b61016081016154e6828573ffffffffffffffffffffffffffffffffffffffff808251168352602082015167ffffffffffffffff80821660208601528060408501511660408601525050806060830151166060840152806080830151166080840152505050565b825163ffffffff90811660a084810191909152602085015173ffffffffffffffffffffffffffffffffffffffff90811660c08601526040860151811660e0860152606086015116610100850152608085015161ffff16610120850152840151166101408301526113c1565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526155988285018b614a30565b915083820360808501526155ac828a614a30565b915060ff881660a085015283820360c08501526155c98288614357565b90861660e08501528381036101008501529050614f708185614357565b60a0815260006155f960a0830188614357565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015267ffffffffffffffff85166060840152828103608084015261563d8185614357565b98975050505050505050565b600082601f83011261565a57600080fd5b8151602061566a614514836144a2565b82815260069290921b8401810191818101908684111561568957600080fd5b8286015b8481101561455c57604081890312156156a65760008081fd5b6156ae6143e3565b81516156b9816144c6565b8152818501518582015283529183019160400161568d565b6000602082840312156156e357600080fd5b815167ffffffffffffffff808211156156fb57600080fd5b90830190610180828603121561571057600080fd5b61571861440c565b6157218361517d565b815261572f6020840161517d565b60208201526040830151604082015261574a606084016153c7565b606082015261575b6080840161517d565b608082015260a083015160a082015261577660c08401614d62565b60c082015261578760e084016153c7565b60e082015261010080840151838111156157a057600080fd5b6157ac888287016150e7565b82840152505061012080840151838111156157c657600080fd5b6157d288828701615649565b82840152505061014091506157e88284016153c7565b9181019190915261016091820151918101919091529392505050565b600081518084526020808501945080840160005b83811015614a7657815187529582019590820190600101615818565b6060815260006158476060830186615804565b82810360208401526158598186615804565b915050826040830152949350505050565b60006020828403121561587c57600080fd5b5051919050565b67ffffffffffffffff8181168382160190808211156140fd576140fd614d8a565b67ffffffffffffffff83168152604081016113c160208301846142ea565b600067ffffffffffffffff808316818103614ee057614ee0614d8a565b6158e981846142ea565b604060208201526000611ae46040830184614357565b60006020828403121561591157600080fd5b813567ffffffffffffffff81111561592857600080fd5b611ae4848285016152e4565b60006040828403121561594657600080fd5b61594e6143e3565b825177ffffffffffffffffffffffffffffffffffffffffffffffff8116811461597657600080fd5b8152602083015161598681614278565b60208201529392505050565b6020815260006113c16020830184614fc4565b600081518084526020808501808196508360051b8101915082860160005b858110156159ed5782840389526159db848351614357565b988501989350908401906001016159c3565b5091979650505050505050565b60608152615a1560608201855167ffffffffffffffff169052565b60006020850151615a32608084018267ffffffffffffffff169052565b50604085015160a0830152606085015173ffffffffffffffffffffffffffffffffffffffff811660c084015250608085015167ffffffffffffffff811660e08401525060a0850151610100818185015260c08701519150610120615a998186018415159052565b60e08801519250610140615ac48187018573ffffffffffffffffffffffffffffffffffffffff169052565b82890151935061018092506101608381880152615ae56101e0880186614357565b928a01518784037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa00185890152929450615b1f8584614fc4565b9450818a01519350615b4a6101a088018573ffffffffffffffffffffffffffffffffffffffff169052565b8901516101c08701525050508281036020840152615b6881866159a5565b915050611ae4604083018415159052565b6000815160208301517fffffffff0000000000000000000000000000000000000000000000000000000080821693506004831015615bc15780818460040360031b1b83161693505b505050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a",
}

var EVM2EVMOffRampABI = EVM2EVMOffRampMetaData.ABI

var EVM2EVMOffRampBin = EVM2EVMOffRampMetaData.Bin

func DeployEVM2EVMOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMOffRampStaticConfig, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig RateLimiterConfig) (common.Address, *types.Transaction, *EVM2EVMOffRamp, error) {
	parsed, err := EVM2EVMOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMOffRampBin), backend, staticConfig, sourceTokens, pools, rateLimiterConfig)
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
	parsed, err := EVM2EVMOffRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "currentRateLimiterState")

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _EVM2EVMOffRamp.Contract.CurrentRateLimiterState(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) CurrentRateLimiterState() (RateLimiterTokenBucket, error) {
	return _EVM2EVMOffRamp.Contract.CurrentRateLimiterState(&_EVM2EVMOffRamp.CallOpts)
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMOffRampDynamicConfig, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(EVM2EVMOffRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMOffRampDynamicConfig)).(*EVM2EVMOffRampDynamicConfig)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetDynamicConfig() (EVM2EVMOffRampDynamicConfig, error) {
	return _EVM2EVMOffRamp.Contract.GetDynamicConfig(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetDynamicConfig() (EVM2EVMOffRampDynamicConfig, error) {
	return _EVM2EVMOffRamp.Contract.GetDynamicConfig(&_EVM2EVMOffRamp.CallOpts)
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetStaticConfig(opts *bind.CallOpts) (EVM2EVMOffRampStaticConfig, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(EVM2EVMOffRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMOffRampStaticConfig)).(*EVM2EVMOffRampStaticConfig)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetStaticConfig() (EVM2EVMOffRampStaticConfig, error) {
	return _EVM2EVMOffRamp.Contract.GetStaticConfig(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetStaticConfig() (EVM2EVMOffRampStaticConfig, error) {
	return _EVM2EVMOffRamp.Contract.GetStaticConfig(&_EVM2EVMOffRamp.CallOpts)
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) IsARMHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "isARMHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) IsARMHealthy() (bool, error) {
	return _EVM2EVMOffRamp.Contract.IsARMHealthy(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) IsARMHealthy() (bool, error) {
	return _EVM2EVMOffRamp.Contract.IsARMHealthy(&_EVM2EVMOffRamp.CallOpts)
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) ApplyPoolUpdates(opts *bind.TransactOpts, removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "applyPoolUpdates", removes, adds)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) ApplyPoolUpdates(removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.ApplyPoolUpdates(&_EVM2EVMOffRamp.TransactOpts, removes, adds)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) ApplyPoolUpdates(removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.ApplyPoolUpdates(&_EVM2EVMOffRamp.TransactOpts, removes, adds)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "executeSingleMessage", message, offchainTokenData, manualExecution)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, offchainTokenData [][]byte, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMOffRamp.TransactOpts, message, offchainTokenData, manualExecution)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, offchainTokenData [][]byte, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMOffRamp.TransactOpts, message, offchainTokenData, manualExecution)
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "setAdmin", newAdmin)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetAdmin(&_EVM2EVMOffRamp.TransactOpts, newAdmin)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetAdmin(&_EVM2EVMOffRamp.TransactOpts, newAdmin)
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMOffRamp.TransactOpts, config)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) SetRateLimiterConfig(config RateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.SetRateLimiterConfig(&_EVM2EVMOffRamp.TransactOpts, config)
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "transmit", reportContext, report, rs, ss, arg4)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.Transmit(&_EVM2EVMOffRamp.TransactOpts, reportContext, report, rs, ss, arg4)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.Transmit(&_EVM2EVMOffRamp.TransactOpts, reportContext, report, rs, ss, arg4)
}

type EVM2EVMOffRampAdminSetIterator struct {
	Event *EVM2EVMOffRampAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampAdminSet)
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
		it.Event = new(EVM2EVMOffRampAdminSet)
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

func (it *EVM2EVMOffRampAdminSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampAdminSet struct {
	NewAdmin common.Address
	Raw      types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMOffRampAdminSetIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampAdminSetIterator{contract: _EVM2EVMOffRamp.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampAdminSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "AdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampAdminSet)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "AdminSet", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseAdminSet(log types.Log) (*EVM2EVMOffRampAdminSet, error) {
	event := new(EVM2EVMOffRampAdminSet)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "AdminSet", log); err != nil {
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
	StaticConfig  EVM2EVMOffRampStaticConfig
	DynamicConfig EVM2EVMOffRampDynamicConfig
	Raw           types.Log
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

type EVM2EVMOffRampConfigSet0Iterator struct {
	Event *EVM2EVMOffRampConfigSet0

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampConfigSet0Iterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampConfigSet0)
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
		it.Event = new(EVM2EVMOffRampConfigSet0)
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

func (it *EVM2EVMOffRampConfigSet0Iterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampConfigSet0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampConfigSet0 struct {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterConfigSet0(opts *bind.FilterOpts) (*EVM2EVMOffRampConfigSet0Iterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "ConfigSet0")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampConfigSet0Iterator{contract: _EVM2EVMOffRamp.contract, event: "ConfigSet0", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchConfigSet0(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampConfigSet0) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "ConfigSet0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampConfigSet0)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "ConfigSet0", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseConfigSet0(log types.Log) (*EVM2EVMOffRampConfigSet0, error) {
	event := new(EVM2EVMOffRampConfigSet0)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "ConfigSet0", log); err != nil {
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
	ReturnData     []byte
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

type EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflightIterator struct {
	Event *EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflightIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight)
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
		it.Event = new(EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight)
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

func (it *EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflightIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight struct {
	Nonce  uint64
	Sender common.Address
	Raw    types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterSkippedSenderWithPreviousRampMessageInflight(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflightIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "SkippedSenderWithPreviousRampMessageInflight", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflightIterator{contract: _EVM2EVMOffRamp.contract, event: "SkippedSenderWithPreviousRampMessageInflight", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchSkippedSenderWithPreviousRampMessageInflight(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight, nonce []uint64, sender []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "SkippedSenderWithPreviousRampMessageInflight", nonceRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "SkippedSenderWithPreviousRampMessageInflight", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseSkippedSenderWithPreviousRampMessageInflight(log types.Log) (*EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight, error) {
	event := new(EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "SkippedSenderWithPreviousRampMessageInflight", log); err != nil {
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
	case _EVM2EVMOffRamp.abi.Events["AdminSet"].ID:
		return _EVM2EVMOffRamp.ParseAdminSet(log)
	case _EVM2EVMOffRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMOffRamp.ParseConfigSet(log)
	case _EVM2EVMOffRamp.abi.Events["ConfigSet0"].ID:
		return _EVM2EVMOffRamp.ParseConfigSet0(log)
	case _EVM2EVMOffRamp.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMOffRamp.ParseExecutionStateChanged(log)
	case _EVM2EVMOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMOffRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMOffRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMOffRamp.abi.Events["PoolAdded"].ID:
		return _EVM2EVMOffRamp.ParsePoolAdded(log)
	case _EVM2EVMOffRamp.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMOffRamp.ParsePoolRemoved(log)
	case _EVM2EVMOffRamp.abi.Events["SkippedIncorrectNonce"].ID:
		return _EVM2EVMOffRamp.ParseSkippedIncorrectNonce(log)
	case _EVM2EVMOffRamp.abi.Events["SkippedSenderWithPreviousRampMessageInflight"].ID:
		return _EVM2EVMOffRamp.ParseSkippedSenderWithPreviousRampMessageInflight(log)
	case _EVM2EVMOffRamp.abi.Events["Transmitted"].ID:
		return _EVM2EVMOffRamp.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMOffRampAdminSet) Topic() common.Hash {
	return common.HexToHash("0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c")
}

func (EVM2EVMOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xccb0480fd9368458c55b5c2e24b62249e1f5eabb55ec2e42d37965a3ae1eaf87")
}

func (EVM2EVMOffRampConfigSet0) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMOffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0xd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65")
}

func (EVM2EVMOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
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

func (EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight) Topic() common.Hash {
	return common.HexToHash("0xe44a20935573a783dd0d5991c92d7b6a0eb3173566530364db3ec10e9a990b5d")
}

func (EVM2EVMOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (_EVM2EVMOffRamp *EVM2EVMOffRamp) Address() common.Address {
	return _EVM2EVMOffRamp.address
}

type EVM2EVMOffRampInterface interface {
	CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error

	CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMOffRampDynamicConfig, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetPoolByDestToken(opts *bind.CallOpts, destToken common.Address) (common.Address, error)

	GetPoolBySourceToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMOffRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	IsARMHealthy(opts *bind.CallOpts) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyPoolUpdates(opts *bind.TransactOpts, removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte, manualExecution bool) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report InternalExecutionReport) (*types.Transaction, error)

	SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error)

	FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMOffRampAdminSetIterator, error)

	WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampAdminSet) (event.Subscription, error)

	ParseAdminSet(log types.Log) (*EVM2EVMOffRampAdminSet, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMOffRampConfigSet, error)

	FilterConfigSet0(opts *bind.FilterOpts) (*EVM2EVMOffRampConfigSet0Iterator, error)

	WatchConfigSet0(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampConfigSet0) (event.Subscription, error)

	ParseConfigSet0(log types.Log) (*EVM2EVMOffRampConfigSet0, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64, messageId [][32]byte) (*EVM2EVMOffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampExecutionStateChanged, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMOffRampExecutionStateChanged, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMOffRampOwnershipTransferred, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMOffRampPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMOffRampPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMOffRampPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMOffRampPoolRemoved, error)

	FilterSkippedIncorrectNonce(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampSkippedIncorrectNonceIterator, error)

	WatchSkippedIncorrectNonce(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampSkippedIncorrectNonce, nonce []uint64, sender []common.Address) (event.Subscription, error)

	ParseSkippedIncorrectNonce(log types.Log) (*EVM2EVMOffRampSkippedIncorrectNonce, error)

	FilterSkippedSenderWithPreviousRampMessageInflight(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflightIterator, error)

	WatchSkippedSenderWithPreviousRampMessageInflight(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight, nonce []uint64, sender []common.Address) (event.Subscription, error)

	ParseSkippedSenderWithPreviousRampMessageInflight(log types.Log) (*EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMOffRampTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
