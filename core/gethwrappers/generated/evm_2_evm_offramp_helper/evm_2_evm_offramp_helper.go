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

type EVM2EVMOffRampDynamicConfig struct {
	PermissionLessExecutionThresholdSeconds uint32
	Router                                  common.Address
	PriceRegistry                           common.Address
	Afn                                     common.Address
	MaxTokensLength                         uint16
	MaxDataSize                             uint32
}

type EVM2EVMOffRampStaticConfig struct {
	CommitStore   common.Address
	ChainId       uint64
	SourceChainId uint64
	OnRamp        common.Address
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
	Rate      *big.Int
	Capacity  *big.Int
}

type RateLimiterTokenBucket struct {
	Capacity    *big.Int
	Tokens      *big.Int
	Rate        *big.Int
	LastUpdated *big.Int
	IsEnabled   bool
}

var EVM2EVMOffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractIPool[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitStoreAlreadyInUse\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"ConsumingMoreThanMaxCapacity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"InvalidNewState\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOffRamp.DynamicConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"InvalidOffRampConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"waitInSeconds\",\"type\":\"uint256\"}],\"name\":\"RateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"TokenDataMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedTokenData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"removes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"internalType\":\"structInternal.PoolUpdate[]\",\"name\":\"adds\",\"type\":\"tuple[]\"}],\"name\":\"applyPoolUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdated\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"rep\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"destTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxTokensLength\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataSize\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMOffRamp.DynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"bitmapIndex\",\"type\":\"uint64\"}],\"name\":\"getExecutionStateBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"getPoolByDestToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPoolBySourceToken\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"sourceTokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"originalSender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"}],\"name\":\"releaseOrMintTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"setExecutionStateHelper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint208\",\"name\":\"rate\",\"type\":\"uint208\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b50604051620060d3380380620060d3833981016040819052620000359162000824565b83838383803380600081620000915760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c457620000c48162000400565b50506040805160a0810182528382018051808352905160208084018290528601516001600160d01b03169383018490524264ffffffffff1660608401819052955115156080909301839052600391909155600455600580547fff0000000000000000000000000000000000000000000000000000000000000016909217600160d01b909402939093176001600160f81b0316600160f81b90930292909217909155508151835114620001885760405162d8548360e71b815260040160405180910390fd5b60608401516001600160a01b03161580620001ab575083516001600160a01b0316155b15620001ca576040516342bcdf7f60e11b815260040160405180910390fd5b83600001516001600160a01b0316634120fccd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200020d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000233919062000910565b6001600160401b03166001146200025d57604051636fc2a20760e11b815260040160405180910390fd5b83516001600160a01b0390811660805260408501516001600160401b0390811660a05260208601511660c05260608501511660e052620002bd7fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a620004ab565b6101005260005b8351811015620003f15762000323848281518110620002e757620002e76200092e565b60200260200101518483815181106200030457620003046200092e565b6020026020010151600e6200051160201b62001a3d179092919060201c565b50620003dd8382815181106200033d576200033d6200092e565b60200260200101516001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000383573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003a9919062000944565b848381518110620003be57620003be6200092e565b602002602001015160116200051160201b62001a3d179092919060201c565b50620003e9816200096b565b9050620002c4565b50505050505050505062000993565b336001600160a01b038216036200045a5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000088565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008160a05160c05160e051604051602001620004f494939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b600062000535846001600160a01b038516846200053d602090811b62001a4f17901c565b949350505050565b60006200053584846001600160a01b03851660008281526002840160209081526040822083905562000535908590859062001a656200057c821b17901c565b60006200058a838362000593565b90505b92915050565b6000818152600183016020526040812054620005dc575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556200058d565b5060006200058d565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715620006205762000620620005e5565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620006515762000651620005e5565b604052919050565b6001600160a01b03811681146200066f57600080fd5b50565b80516001600160401b03811681146200068a57600080fd5b919050565b60006001600160401b03821115620006ab57620006ab620005e5565b5060051b60200190565b600082601f830112620006c757600080fd5b81516020620006e0620006da836200068f565b62000626565b82815260059290921b840181019181810190868411156200070057600080fd5b8286015b84811015620007285780516200071a8162000659565b835291830191830162000704565b509695505050505050565b600082601f8301126200074557600080fd5b8151602062000758620006da836200068f565b82815260059290921b840181019181810190868411156200077857600080fd5b8286015b8481101562000728578051620007928162000659565b83529183019183016200077c565b600060608284031215620007b357600080fd5b604051606081016001600160401b0381118282101715620007d857620007d8620005e5565b806040525080915082518015158114620007f157600080fd5b815260208301516001600160d01b03811681146200080e57600080fd5b6020820152604092830151920191909152919050565b6000806000808486036101208112156200083d57600080fd5b60808112156200084c57600080fd5b5062000857620005fb565b8551620008648162000659565b8152620008746020870162000672565b6020820152620008876040870162000672565b604082015260608601516200089c8162000659565b606082015260808601519094506001600160401b0380821115620008bf57600080fd5b620008cd88838901620006b5565b945060a0870151915080821115620008e457600080fd5b50620008f38782880162000733565b925050620009058660c08701620007a0565b905092959194509250565b6000602082840312156200092357600080fd5b6200058a8262000672565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156200095757600080fd5b8151620009648162000659565b9392505050565b6000600182016200098c57634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c05160e051610100516156b562000a1e60003960006126cb0152600081816102d301528181611d180152612cf401526000818161027301528181611cc60152612cd30152600081816102a301528181611cf00152818161210501528181612cb2015261345701526000818161024401528181611c98015261277b01526156b56000f3fe608060405234801561001057600080fd5b506004361061020b5760003560e01c80637437ff9f1161012a5780638fa4f53b116100bd578063b57671661161008c578063d3c7c2c711610071578063d3c7c2c714610730578063d7e2bb5014610738578063f2fde38b1461074b57600080fd5b8063b576716614610715578063c5a1d7f01461072857600080fd5b80638fa4f53b146106bc578063afcb95d7146106cf578063b1dc65a4146106ef578063b4069b311461070257600080fd5b806381ff7048116100f957806381ff70481461062157806385572ffb14610651578063856c82471461065f5780638da5cb5b146106ab57600080fd5b80637437ff9f14610517578063776e0ac4146105e657806379ba5097146106065780637c0e7c7a1461060e57600080fd5b80634a11d44a116101a25780635d86f141116101715780635d86f141146104c7578063666cab8d146104da578063681fba16146104ef578063704b6c021461050457600080fd5b80634a11d44a146103e857806350644972146103fb578063546719cd14610433578063599f6431146104a257600080fd5b80631ef38174116101de5780631ef381741461039a5780632dea00f3146103af5780633091aee7146103c25780633a87ac53146103d557600080fd5b806306285c6914610210578063142a98fc14610319578063147809b314610339578063181f5a7714610351575b600080fd5b61030360408051608081018252600080825260208201819052918101829052606081019190915260405180608001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b6040516103109190613bf8565b60405180910390f35b61032c610327366004613c6a565b61075e565b6040516103109190613cf1565b6103416107d9565b6040519015158152602001610310565b61038d6040518060400160405280601481526020017f45564d3245564d4f666652616d7020312e302e3000000000000000000000000081525081565b6040516103109190613d75565b6103ad6103a8366004613fd6565b610866565b005b6103ad6103bd3660046140a3565b610c3d565b6103ad6103d03660046140f9565b610c4b565b6103ad6103e3366004614214565b610cb6565b6103ad6103f6366004614394565b611046565b610425610409366004613c6a565b67ffffffffffffffff1660009081526015602052604090205490565b604051908152602001610310565b61043b611235565b6040516103109190600060a082019050825182526020830151602083015279ffffffffffffffffffffffffffffffffffffffffffffffffffff604084015116604083015264ffffffffff606084015116606083015260808301511515608083015292915050565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610310565b6104af6104d53660046144f9565b611304565b6104e2611360565b604051610310919061455a565b6104f76113c2565b604051610310919061456d565b6103ad6105123660046144f9565b61146e565b6105d96040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c081018252600b5463ffffffff80821683526001600160a01b0364010000000090920482166020840152600c54821693830193909352600d54908116606083015261ffff740100000000000000000000000000000000000000008204166080830152760100000000000000000000000000000000000000000000900490911660a082015290565b60405161031091906145ba565b6105f96105f43660046146a3565b611537565b6040516103109190614790565b6103ad611591565b6103ad61061c3660046149a5565b611674565b6008546006546040805163ffffffff80851682526401000000009094049093166020840152820152606001610310565b6103ad61020b3660046149ec565b61069261066d3660046144f9565b6001600160a01b031660009081526014602052604090205467ffffffffffffffff1690565b60405167ffffffffffffffff9091168152602001610310565b6000546001600160a01b03166104af565b6103ad6106ca366004614a27565b61167e565b604080516001815260006020820181905291810191909152606001610310565b6103ad6106fd366004614a5c565b611689565b6104af6107103660046144f9565b61188c565b6103ad610723366004614b13565b611940565b61042561194a565b6104f7611975565b6104af6107463660046144f9565b611a1d565b6103ad6107593660046144f9565b611a2c565b600061076c60016004614b84565b6002610779608085614bca565b67ffffffffffffffff1661078d9190614bf1565b6015600061079c608087614c2e565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002054901c1660038111156107d3576107d3613c87565b92915050565b600d54604080517f397796f700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163397796f79160048083019260209291908290030181865afa15801561083c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108609190614c60565b15905090565b8360ff16806000036108d9576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f736974697665000000000000000000000000000060448201526064015b60405180910390fd5b6108e1611a71565b6108ea84611ae7565b600a5460005b818110156109695760096000600a838154811061090f5761090f614c7d565b60009182526020808320909101546001600160a01b03168352820192909252604001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905561096281614cac565b90506108f0565b50865160005b81811015610afe57600089828151811061098b5761098b614c7d565b60200260200101519050600060028111156109a8576109a8613c87565b6001600160a01b038216600090815260096020526040902054610100900460ff1660028111156109da576109da613c87565b14610a41576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d697474657220616464726573730000000060448201526064016108d0565b6040805180820190915260ff8316815260208101600290526001600160a01b03821660009081526009602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610ae457610ae4613c87565b02179055509050505080610af790614cac565b905061096f565b508751610b1290600a9060208b0190613b6f565b506007805460ff838116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909216908a161717905560088054610b98914691309190600090610b6a9063ffffffff16614ce4565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168c8c8c8c8c8c611d7b565b6006819055600880544363ffffffff9081166401000000009081027fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff841681179094556040519083048216947f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0594610c299487949293918316921691909117908f908f908f908f908f908f90614d07565b60405180910390a150505050505050505050565b610c478282611e26565b5050565b6000546001600160a01b03163314801590610c7157506002546001600160a01b03163314155b15610ca8576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cb3600382611ed0565b50565b610cbe611a71565b60005b8251811015610e7d576000838281518110610cde57610cde614c7d565b60200260200101516000015190506000848381518110610d0057610d00614c7d565b6020026020010151602001519050610d2282600e611ff990919063ffffffff16565b610d58576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038116610d6d600e8461200e565b6001600160a01b031614610dad576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610db8600e83612023565b50610e26816001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dfa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1e9190614d9d565b601190612023565b50604080516001600160a01b038085168252831660208201527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a1505080610e7690614cac565b9050610cc1565b5060005b8151811015611041576000828281518110610e9e57610e9e614c7d565b60200260200101516000015190506000838381518110610ec057610ec0614c7d565b602002602001015160200151905060006001600160a01b0316826001600160a01b03161480610ef657506001600160a01b038116155b15610f2d576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f38600e83611ff9565b15610f6f576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f7b600e8383611a3d565b50610fea816001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fe19190614d9d565b60119083611a3d565b50604080516001600160a01b038085168252831660208201527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c910160405180910390a150508061103a90614cac565b9050610e81565b505050565b33301461107f576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051600080825260208201909252816110bc565b60408051808201909152600080825260208201528152602001906001900390816110955790505b50610120850151519091501561110e576101208401516060850151604080516001600160a01b03909216602083015261110b9291016040516020818303038152906040528660e0015186612038565b90505b60e08401516001600160a01b03163b158061115e575060e084015161115c906001600160a01b03167f85572ffb0000000000000000000000000000000000000000000000000000000061228e565b155b156111695750505050565b600b5464010000000090046001600160a01b0316635607b37561118c86846122aa565b848760a001518860e001516040518563ffffffff1660e01b81526004016111b69493929190614dba565b6020604051808303816000875af11580156111d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111f99190614c60565b61122f576040517fee4f4da800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526003548152600454602082015260055479ffffffffffffffffffffffffffffffffffffffffffffffffffff8116928201929092527a010000000000000000000000000000000000000000000000000000820464ffffffffff1660608201527f010000000000000000000000000000000000000000000000000000000000000090910460ff16151560808201526112ff9061234d565b905090565b60008080611313600e856123f0565b9150915081611359576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016108d0565b9392505050565b6060600a8054806020026020016040519081016040528092919081815260200182805480156113b857602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161139a575b5050505050905090565b60606113ce6011612412565b67ffffffffffffffff8111156113e6576113e6613d88565b60405190808252806020026020018201604052801561140f578160200160208202803683370190505b50905060005b815181101561146a57600061142b60118361241d565b5090508083838151811061144157611441614c7d565b6001600160a01b03909216602092830291909101909101525061146381614cac565b9050611415565b5090565b6000546001600160a01b0316331480159061149457506002546001600160a01b03163314155b156114cb576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c9060200160405180910390a150565b60606115868787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250899250611581915087905088614e7d565b612038565b979650505050505050565b6001546001600160a01b03163314611605576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016108d0565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610c478282612439565b610cb3816001612439565b6116938787612c97565b6006548835908082146116dc576040517f93df584c00000000000000000000000000000000000000000000000000000000815260048101829052602481018390526044016108d0565b6040805183815260208c81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a13360009081526009602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561176457611764613c87565b600281111561177557611775613c87565b905250905060028160200151600281111561179257611792613c87565b1480156117cc5750600a816000015160ff16815481106117b4576117b4614c7d565b6000918252602090912001546001600160a01b031633145b611802576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506000611810856020614bf1565b61181b886020614bf1565b6118278b610144614e8a565b6118319190614e8a565b61183b9190614e8a565b905036811461187f576040517f8e1192e1000000000000000000000000000000000000000000000000000000008152600481018290523660248201526044016108d0565b5050505050505050505050565b6000808061189b600e856123f0565b91509150816118d6576040517f9c8787c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611914573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119389190614d9d565b949350505050565b610c478282612c97565b60006112ff7fbdd59ac4dd1d82276c9a9c5d2656546346b9dcdb1f9b4204aed4ec15c23d7d3a612cad565b6060611981600e612412565b67ffffffffffffffff81111561199957611999613d88565b6040519080825280602002602001820160405280156119c2578160200160208202803683370190505b50905060005b815181101561146a5760006119de600e8361241d565b509050808383815181106119f4576119f4614c7d565b6001600160a01b039092166020928302919091019091015250611a1681614cac565b90506119c8565b600080806113136011856123f0565b611a34611a71565b610cb381612d6d565b6000611938846001600160a01b038516845b600061193884846001600160a01b038516612e48565b60006113598383612e65565b6000546001600160a01b03163314611ae5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016108d0565b565b600081806020019051810190611afd9190614ec1565b60208101519091506001600160a01b03161580611b25575060608101516001600160a01b0316155b15611b5e57806040517f9c1779a70000000000000000000000000000000000000000000000000000000081526004016108d091906145ba565b8051600b80546020808501516001600160a01b03908116640100000000027fffffffffffffffff00000000000000000000000000000000000000000000000090931663ffffffff9586161792909217909255604080850151600c80549184167fffffffffffffffffffffffff0000000000000000000000000000000000000000909216919091179055606080860151600d80546080808a015160a08b0151909916760100000000000000000000000000000000000000000000027fffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffffff61ffff909a1674010000000000000000000000000000000000000000027fffffffffffffffffffff00000000000000000000000000000000000000000000909316948816949094179190911797909716919091179055815194850182527f00000000000000000000000000000000000000000000000000000000000000008316855267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116948601949094527f0000000000000000000000000000000000000000000000000000000000000000909316848201527f000000000000000000000000000000000000000000000000000000000000000090911691830191909152517fb53c6ae8a09cb3c0b6a39aaf24999010a005da24252219c897301f8cd594d60191611d6f918490614f6f565b60405180910390a15050565b6000808a8a8a8a8a8a8a8a8a604051602001611d9f9998979695949392919061501d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b60006002611e35608085614bca565b67ffffffffffffffff16611e499190614bf1565b90506000601581611e5b608087614c2e565b67ffffffffffffffff168152602081019190915260400160002054905081611e8560016004614b84565b901b191681836003811115611e9c57611e9c613c87565b901b178060156000611eaf608088614c2e565b67ffffffffffffffff16815260208101919091526040016000205550505050565b611ed982612eb4565b8051600283018054604084015180865560208501517effffffffff00000000000000000000000000000000000000000000000000009092167f0100000000000000000000000000000000000000000000000000000000000000941515949094027fffffffffffff0000000000000000000000000000000000000000000000000000169390931779ffffffffffffffffffffffffffffffffffffffffffffffffffff9091161790556001830154611f8f9190613008565b60018301556040805182511515815260208084015179ffffffffffffffffffffffffffffffffffffffffffffffffffff169082015282820151918101919091527f44a2350342338075ac038f37b8d9e49e696e360492cb44cc6bc37fc117f19df890606001611d6f565b6000611359836001600160a01b03841661301e565b6000611359836001600160a01b03841661302a565b6000611359836001600160a01b038416613036565b60606000855167ffffffffffffffff81111561205657612056613d88565b60405190808252806020026020018201604052801561209b57816020015b60408051808201909152600080825260208201528152602001906001900390816120745790505b50905060005b865181101561226d5760006120d28883815181106120c1576120c1614c7d565b602002602001015160000151611304565b9050806001600160a01b0316638627fad688888b86815181106120f7576120f7614c7d565b6020026020010151602001517f00000000000000000000000000000000000000000000000000000000000000008a888151811061213657612136614c7d565b60200260200101516040518663ffffffff1660e01b815260040161215e9594939291906150a5565b600060405180830381600087803b15801561217857600080fd5b505af115801561218c573d6000803e3d6000fd5b50505050806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121f29190614d9d565b83838151811061220457612204614c7d565b60209081029190910101516001600160a01b039091169052875188908390811061223057612230614c7d565b60200260200101516020015183838151811061224e5761224e614c7d565b60209081029190910181015101525061226681614cac565b90506120a1565b50600c546122859082906001600160a01b0316613042565b95945050505050565b6000612299836131f3565b801561135957506113598383613257565b6040805160a08101825260008082526020820152606091810182905281810182905260808101919091526040518060a001604052808461016001518152602001846000015167ffffffffffffffff168152602001846060015160405160200161232291906001600160a01b0391909116815260200190565b6040516020818303038152906040528152602001846101000151815260200183815250905092915050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526000826060015164ffffffffff16426123919190614b84565b835160408501519192506123d9916123c59079ffffffffffffffffffffffffffffffffffffffffffffffffffff1684614bf1565b85602001516123d49190614e8a565b613008565b6020840152505064ffffffffff4216606082015290565b600080612406846001600160a01b038516613322565b915091505b9250929050565b60006107d382613331565b600080808061242c868661333c565b9097909650945050505050565b600d54604080517f397796f700000000000000000000000000000000000000000000000000000000815290516001600160a01b039092169163397796f7916004808201926020929091908290030181865afa15801561249c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124c09190614c60565b156124f6576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020820151516000819003612536576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8260400151518114612574576040517f57e0e08300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff81111561258f5761258f613d88565b6040519080825280602002602001820160405280156125b8578160200160208202803683370190505b50905060008267ffffffffffffffff8111156125d6576125d6613d88565b60405190808252806020026020018201604052801561268257816020015b60408051610180810182526000808252602080830182905292820181905260608083018290526080830182905260a0830182905260c0830182905260e083018290526101008301819052610120830152610140820181905261016082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092019101816125f45790505b50905060005b8381101561273d576000866020015182815181106126a8576126a8614c7d565b60200260200101518060200190518101906126c391906151d3565b90506126ef817f000000000000000000000000000000000000000000000000000000000000000061334b565b84838151811061270157612701614c7d565b6020026020010181815250508083838151811061272057612720614c7d565b6020026020010181905250508061273690614cac565b9050612688565b50606085015160808601516040517f320488750000000000000000000000000000000000000000000000000000000081526000926001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016926332048875926127b29288929091600401615336565b6020604051808303816000875af11580156127d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127f5919061536c565b905080600003612831576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b84811015612c8e57600083828151811061285057612850614c7d565b602002602001015190506000612869826020015161075e565b9050600081600381111561287f5761287f613c87565b148061289c5750600381600381111561289a5761289a613c87565b145b6128e45760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016108d0565b871561295c57600b5460009063ffffffff166129008642614b84565b11905080806129205750600382600381111561291e5761291e613c87565b145b612956576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506129b9565b600081600381111561297057612970613c87565b146129b95760208201516040517f67d9ba0f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016108d0565b60008160038111156129cd576129cd613c87565b03612a6c57608082015160608301516001600160a01b031660009081526014602052604090205467ffffffffffffffff91821691612a0d91166001615385565b67ffffffffffffffff1614612a6c5781606001516001600160a01b0316826080015167ffffffffffffffff167fd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf4123760405160405180910390a35050612c7e565b600089604001518481518110612a8457612a84614c7d565b60200260200101519050612a99838251613455565b612aa883602001516001611e26565b6000612ab584838c61362b565b9050612ac5846020015182611e26565b6003816003811115612ad957612ad9613c87565b14158015612af957506002816003811115612af657612af6613c87565b14155b15612b38578360200151816040517f9e2616030000000000000000000000000000000000000000000000000000000081526004016108d09291906153b1565b8360c0015115612bba576002816003811115612b5657612b56613c87565b03612bb55760608401516001600160a01b03166000908152601460205260408120805467ffffffffffffffff1691612b8d836153cf565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b612c2d565b6000836003811115612bce57612bce613c87565b03612c2d5760608401516001600160a01b03166000908152601460205260408120805467ffffffffffffffff1691612c05836153cf565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b836101600151846020015167ffffffffffffffff167f5e04b4755a4460aa6de4f3a906c4324a025c7449c02b52f5466659b5bfdfba5f83604051612c719190613cf1565b60405180910390a3505050505b612c8781614cac565b9050612834565b50505050505050565b610c47612ca682840184614a27565b6000612439565b6000817f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000604051602001612d50949392919093845267ffffffffffffffff9283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b336001600160a01b03821603612ddf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016108d0565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600082815260028401602052604081208290556119388484611a65565b6000818152600183016020526040812054612eac575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556107d3565b5060006107d3565b805460018201541480612ef05750600281015464ffffffffff7a0100000000000000000000000000000000000000000000000000009091041642145b15612ef85750565b805460018201541115612f37576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002810154600090612f6e907a010000000000000000000000000000000000000000000000000000900464ffffffffff1642614b84565b82546002840154919250612fb191612fa29079ffffffffffffffffffffffffffffffffffffffffffffffffffff1684614bf1565b84600101546123d49190614e8a565b60018301555060020180547fff0000000000ffffffffffffffffffffffffffffffffffffffffffffffffffff167a0100000000000000000000000000000000000000000000000000004264ffffffffff1602179055565b60008183106130175781611359565b5090919050565b60006113598383613766565b60006113598383613772565b600061135983836137fc565b81516000805b828110156131e7576000846001600160a01b031663d02641a087848151811061307357613073614c7d565b6020908102919091010151516040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b0390911660048201526024016040805180830381865afa1580156130da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130fe91906153ec565b51905077ffffffffffffffffffffffffffffffffffffffffffffffff811660000361317f5785828151811061313557613135614c7d565b6020908102919091010151516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201526024016108d0565b6131c986838151811061319457613194614c7d565b6020026020010151602001518277ffffffffffffffffffffffffffffffffffffffffffffffff1661381990919063ffffffff16565b6131d39084614e8a565b925050806131e090614cac565b9050613048565b5061122f600382613852565b600061321f827f01ffc9a700000000000000000000000000000000000000000000000000000000613257565b80156107d35750613250827fffffffff00000000000000000000000000000000000000000000000000000000613257565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d9150600051905082801561330f575060208210155b8015611586575015159695505050505050565b600080808061242c86866139a8565b60006107d3826139e2565b600080808061242c86866139ed565b60008060001b828460200151856080015186606001518760e00151886101000151805190602001208961012001516040516020016133899190614790565b604051602081830303815290604052805190602001208a60a001518b60c001518c61014001518d604001516040516020016134379c9b9a999897969594939291909b8c5260208c019a909a5267ffffffffffffffff98891660408c01529690971660608a01526001600160a01b0394851660808a015292841660a089015260c088019190915260e0870152610100860152911515610120850152166101408301526101608201526101800190565b60405160208183030381529060405280519060200120905092915050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16826000015167ffffffffffffffff16146134d55781516040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016108d0565b600d54610120830151517401000000000000000000000000000000000000000090910461ffff1610156135465760208201516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016108d0565b8082610120015151146135975760208201516040517f8808f8e700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016108d0565b600d546101008301515176010000000000000000000000000000000000000000000090910463ffffffff161015610c4757600d54610100830151516040517f8693378900000000000000000000000000000000000000000000000000000000815276010000000000000000000000000000000000000000000090920463ffffffff16600483015260248201526044016108d0565b6040517f4a11d44a0000000000000000000000000000000000000000000000000000000081526000903090634a11d44a9061366e908790879087906004016154bd565b600060405180830381600087803b15801561368857600080fd5b505af1925050508015613699575060015b61375c573d8080156136c7576040519150601f19603f3d011682016040523d82523d6000602084013e6136cc565b606091505b506136d681615615565b7fffffffff00000000000000000000000000000000000000000000000000000000167fee4f4da80000000000000000000000000000000000000000000000000000000003613728576003915050611359565b806040517fcf19edfd0000000000000000000000000000000000000000000000000000000081526004016108d09190613d75565b5060029392505050565b60006113598383613a18565b60008181526002830160205260408120548015158061379657506137968484613766565b611359576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f456e756d657261626c654d61703a206e6f6e6578697374656e74206b6579000060448201526064016108d0565b600081815260028301602052604081208190556113598383613a30565b6000670de0b6b3a76400006138488377ffffffffffffffffffffffffffffffffffffffffffffffff8616614bf1565b6113599190615665565b60028201547f0100000000000000000000000000000000000000000000000000000000000000900460ff161580613887575080155b15613890575050565b61389982612eb4565b81548111156138e15781546040517f48369c430000000000000000000000000000000000000000000000000000000081526004810191909152602481018290526044016108d0565b808260010154101561395f576002820154600183015479ffffffffffffffffffffffffffffffffffffffffffffffffffff909116906139209083614b84565b61392a9190615665565b6040517fdc96cefa0000000000000000000000000000000000000000000000000000000081526004016108d091815260200190565b808260010160008282546139739190614b84565b90915550506040518181527f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a90602001611d6f565b60008181526002830160205260408120548190806139d7576139ca8585613766565b92506000915061240b9050565b60019250905061240b565b60006107d382613a3c565b600080806139fb8585613a46565b600081815260029690960160205260409095205494959350505050565b60008181526001830160205260408120541515611359565b60006113598383613a52565b60006107d3825490565b60006113598383613b45565b60008181526001830160205260408120548015613b3b576000613a76600183614b84565b8554909150600090613a8a90600190614b84565b9050818114613aef576000866000018281548110613aaa57613aaa614c7d565b9060005260206000200154905080876000018481548110613acd57613acd614c7d565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080613b0057613b00615679565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506107d3565b60009150506107d3565b6000826000018281548110613b5c57613b5c614c7d565b9060005260206000200154905092915050565b828054828255906000526020600020908101928215613bdc579160200282015b82811115613bdc57825182547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03909116178255602090920191600190910190613b8f565b5061146a9291505b8082111561146a5760008155600101613be4565b608081016107d382846001600160a01b03808251168352602082015167ffffffffffffffff80821660208601528060408501511660408601525050806060830151166060840152505050565b67ffffffffffffffff81168114610cb357600080fd5b8035613c6581613c44565b919050565b600060208284031215613c7c57600080fd5b813561135981613c44565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110613ced577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b602081016107d38284613cb6565b60005b83811015613d1a578181015183820152602001613d02565b8381111561122f5750506000910152565b60008151808452613d43816020860160208601613cff565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006113596020830184613d2b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613dda57613dda613d88565b60405290565b604051610180810167ffffffffffffffff81118282101715613dda57613dda613d88565b60405160a0810167ffffffffffffffff81118282101715613dda57613dda613d88565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613e6e57613e6e613d88565b604052919050565b600067ffffffffffffffff821115613e9057613e90613d88565b5060051b60200190565b6001600160a01b0381168114610cb357600080fd5b8035613c6581613e9a565b600082601f830112613ecb57600080fd5b81356020613ee0613edb83613e76565b613e27565b82815260059290921b84018101918181019086841115613eff57600080fd5b8286015b84811015613f23578035613f1681613e9a565b8352918301918301613f03565b509695505050505050565b803560ff81168114613c6557600080fd5b600067ffffffffffffffff821115613f5957613f59613d88565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112613f9657600080fd5b8135613fa4613edb82613f3f565b818152846020838601011115613fb957600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c08789031215613fef57600080fd5b863567ffffffffffffffff8082111561400757600080fd5b6140138a838b01613eba565b9750602089013591508082111561402957600080fd5b6140358a838b01613eba565b965061404360408a01613f2e565b9550606089013591508082111561405957600080fd5b6140658a838b01613f85565b945061407360808a01613c5a565b935060a089013591508082111561408957600080fd5b5061409689828a01613f85565b9150509295509295509295565b600080604083850312156140b657600080fd5b82356140c181613c44565b91506020830135600481106140d557600080fd5b809150509250929050565b8015158114610cb357600080fd5b8035613c65816140e0565b60006060828403121561410b57600080fd5b6040516060810181811067ffffffffffffffff8211171561412e5761412e613d88565b604052823561413c816140e0565b8152602083013579ffffffffffffffffffffffffffffffffffffffffffffffffffff8116811461416b57600080fd5b60208201526040928301359281019290925250919050565b600082601f83011261419457600080fd5b813560206141a4613edb83613e76565b82815260069290921b840181019181810190868411156141c357600080fd5b8286015b84811015613f2357604081890312156141e05760008081fd5b6141e8613db7565b81356141f381613e9a565b81528185013561420281613e9a565b818601528352918301916040016141c7565b6000806040838503121561422757600080fd5b823567ffffffffffffffff8082111561423f57600080fd5b61424b86838701614183565b9350602085013591508082111561426157600080fd5b5061426e85828601614183565b9150509250929050565b600082601f83011261428957600080fd5b81356020614299613edb83613e76565b82815260069290921b840181019181810190868411156142b857600080fd5b8286015b84811015613f2357604081890312156142d55760008081fd5b6142dd613db7565b81356142e881613e9a565b815281850135858201528352918301916040016142bc565b600061430e613edb84613e76565b8381529050602080820190600585901b84018681111561432d57600080fd5b845b8181101561436957803567ffffffffffffffff81111561434f5760008081fd5b61435b89828901613f85565b85525092820192820161432f565b505050509392505050565b600082601f83011261438557600080fd5b61135983833560208501614300565b6000806000606084860312156143a957600080fd5b833567ffffffffffffffff808211156143c157600080fd5b9085019061018082880312156143d657600080fd5b6143de613de0565b6143e783613c5a565b81526143f560208401613c5a565b60208201526040830135604082015261441060608401613eaf565b606082015261442160808401613c5a565b608082015260a083013560a082015261443c60c084016140ee565b60c082015261444d60e08401613eaf565b60e0820152610100808401358381111561446657600080fd5b6144728a828701613f85565b828401525050610120808401358381111561448c57600080fd5b6144988a828701614278565b8284015250506101406144ac818501613eaf565b908201526101609283013592810192909252909350602085013590808211156144d457600080fd5b506144e186828701614374565b9250506144f0604085016140ee565b90509250925092565b60006020828403121561450b57600080fd5b813561135981613e9a565b600081518084526020808501945080840160005b8381101561454f5781516001600160a01b03168752958201959082019060010161452a565b509495945050505050565b6020815260006113596020830184614516565b6020808252825182820181905260009190848201906040850190845b818110156145ae5783516001600160a01b031683529284019291840191600101614589565b50909695505050505050565b60c081016107d3828463ffffffff80825116835260208201516001600160a01b038082166020860152806040850151166040860152806060850151166060860152505061ffff60808301511660808401528060a08301511660a0840152505050565b60008083601f84011261462e57600080fd5b50813567ffffffffffffffff81111561464657600080fd5b60208301915083602082850101111561240b57600080fd5b60008083601f84011261467057600080fd5b50813567ffffffffffffffff81111561468857600080fd5b6020830191508360208260051b850101111561240b57600080fd5b600080600080600080608087890312156146bc57600080fd5b863567ffffffffffffffff808211156146d457600080fd5b6146e08a838b01614278565b975060208901359150808211156146f657600080fd5b6147028a838b0161461c565b90975095506040890135915061471782613e9a565b9093506060880135908082111561472d57600080fd5b5061473a89828a0161465e565b979a9699509497509295939492505050565b600081518084526020808501945080840160005b8381101561454f57815180516001600160a01b031688528301518388015260409096019590820190600101614760565b602081526000611359602083018461474c565b600082601f8301126147b457600080fd5b813560206147c4613edb83613e76565b82815260059290921b840181019181810190868411156147e357600080fd5b8286015b84811015613f235780356147fa81613c44565b83529183019183016147e7565b600082601f83011261481857600080fd5b81356020614828613edb83613e76565b82815260059290921b8401810191818101908684111561484757600080fd5b8286015b84811015613f2357803567ffffffffffffffff81111561486b5760008081fd5b6148798986838b0101614374565b84525091830191830161484b565b600082601f83011261489857600080fd5b813560206148a8613edb83613e76565b82815260059290921b840181019181810190868411156148c757600080fd5b8286015b84811015613f2357803583529183019183016148cb565b600060a082840312156148f457600080fd5b6148fc613e04565b9050813567ffffffffffffffff8082111561491657600080fd5b614922858386016147a3565b8352602084013591508082111561493857600080fd5b61494485838601614374565b6020840152604084013591508082111561495d57600080fd5b61496985838601614807565b6040840152606084013591508082111561498257600080fd5b5061498f84828501614887565b6060830152506080820135608082015292915050565b600080604083850312156149b857600080fd5b823567ffffffffffffffff8111156149cf57600080fd5b6149db858286016148e2565b92505060208301356140d5816140e0565b6000602082840312156149fe57600080fd5b813567ffffffffffffffff811115614a1557600080fd5b820160a0818503121561135957600080fd5b600060208284031215614a3957600080fd5b813567ffffffffffffffff811115614a5057600080fd5b611938848285016148e2565b60008060008060008060008060e0898b031215614a7857600080fd5b606089018a811115614a8957600080fd5b8998503567ffffffffffffffff80821115614aa357600080fd5b614aaf8c838d0161461c565b909950975060808b0135915080821115614ac857600080fd5b614ad48c838d0161465e565b909750955060a08b0135915080821115614aed57600080fd5b50614afa8b828c0161465e565b999c989b50969995989497949560c00135949350505050565b60008060208385031215614b2657600080fd5b823567ffffffffffffffff811115614b3d57600080fd5b614b498582860161461c565b90969095509350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015614b9657614b96614b55565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff80841680614be557614be5614b9b565b92169190910692915050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614c2957614c29614b55565b500290565b600067ffffffffffffffff80841680614c4957614c49614b9b565b92169190910492915050565b8051613c65816140e0565b600060208284031215614c7257600080fd5b8151611359816140e0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614cdd57614cdd614b55565b5060010190565b600063ffffffff808316818103614cfd57614cfd614b55565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614d378184018a614516565b90508281036080840152614d4b8189614516565b905060ff871660a084015282810360c0840152614d688187613d2b565b905067ffffffffffffffff851660e0840152828103610100840152614d8d8185613d2b565b9c9b505050505050505050505050565b600060208284031215614daf57600080fd5b815161135981613e9a565b608081528451608082015267ffffffffffffffff60208601511660a08201526000604086015160a060c0840152614df5610120840182613d2b565b905060608701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80808584030160e0860152614e318383613d2b565b925060808901519150808584030161010086015250614e50828261474c565b92505050614e62602083018615159052565b83604083015261228560608301846001600160a01b03169052565b6000611359368484614300565b60008219821115614e9d57614e9d614b55565b500190565b805163ffffffff81168114613c6557600080fd5b8051613c6581613e9a565b600060c08284031215614ed357600080fd5b60405160c0810181811067ffffffffffffffff82111715614ef657614ef6613d88565b604052614f0283614ea2565b81526020830151614f1281613e9a565b60208201526040830151614f2581613e9a565b60408201526060830151614f3881613e9a565b6060820152608083015161ffff81168114614f5257600080fd5b6080820152614f6360a08401614ea2565b60a08201529392505050565b6101408101614fbc82856001600160a01b03808251168352602082015167ffffffffffffffff80821660208601528060408501511660408601525050806060830151166060840152505050565b611359608083018463ffffffff80825116835260208201516001600160a01b038082166020860152806040850151166040860152806060850151166060860152505061ffff60808301511660808401528060a08301511660a0840152505050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526150578285018b614516565b9150838203608085015261506b828a614516565b915060ff881660a085015283820360c08501526150888288613d2b565b90861660e08501528381036101008501529050614d8d8185613d2b565b60a0815260006150b860a0830188613d2b565b6001600160a01b038716602084015285604084015267ffffffffffffffff8516606084015282810360808401526150ef8185613d2b565b98975050505050505050565b8051613c6581613c44565b600082601f83011261511757600080fd5b8151615125613edb82613f3f565b81815284602083860101111561513a57600080fd5b611938826020830160208701613cff565b600082601f83011261515c57600080fd5b8151602061516c613edb83613e76565b82815260069290921b8401810191818101908684111561518b57600080fd5b8286015b84811015613f2357604081890312156151a85760008081fd5b6151b0613db7565b81516151bb81613e9a565b8152818501518582015283529183019160400161518f565b6000602082840312156151e557600080fd5b815167ffffffffffffffff808211156151fd57600080fd5b90830190610180828603121561521257600080fd5b61521a613de0565b615223836150fb565b8152615231602084016150fb565b60208201526040830151604082015261524c60608401614eb6565b606082015261525d608084016150fb565b608082015260a083015160a082015261527860c08401614c55565b60c082015261528960e08401614eb6565b60e082015261010080840151838111156152a257600080fd5b6152ae88828701615106565b82840152505061012080840151838111156152c857600080fd5b6152d48882870161514b565b82840152505061014091506152ea828401614eb6565b9181019190915261016091820151918101919091529392505050565b600081518084526020808501945080840160005b8381101561454f5781518752958201959082019060010161531a565b6060815260006153496060830186615306565b828103602084015261535b8186615306565b915050826040830152949350505050565b60006020828403121561537e57600080fd5b5051919050565b600067ffffffffffffffff8083168185168083038211156153a8576153a8614b55565b01949350505050565b67ffffffffffffffff83168152604081016113596020830184613cb6565b600067ffffffffffffffff808316818103614cfd57614cfd614b55565b6000604082840312156153fe57600080fd5b6040516040810181811067ffffffffffffffff8211171561542157615421613d88565b604052825177ffffffffffffffffffffffffffffffffffffffffffffffff8116811461544c57600080fd5b8152602083015161545c81613c44565b60208201529392505050565b600081518084526020808501808196508360051b8101915082860160005b858110156154b057828403895261549e848351613d2b565b98850198935090840190600101615486565b5091979650505050505050565b606081526154d860608201855167ffffffffffffffff169052565b600060208501516154f5608084018267ffffffffffffffff169052565b50604085015160a083015260608501516001600160a01b03811660c084015250608085015167ffffffffffffffff811660e08401525060a0850151610100818185015260c0870151915061012061554f8186018415159052565b60e0880151925061014061556d818701856001600160a01b03169052565b8289015193506101809250610160838188015261558e6101e0880186613d2b565b928a01518784037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa001858901529294506155c8858461474c565b9450818a015193506155e66101a08801856001600160a01b03169052565b8901516101c087015250505082810360208401526156048186615468565b915050611938604083018415159052565b6000815160208301517fffffffff000000000000000000000000000000000000000000000000000000008082169350600483101561565d5780818460040360031b1b83161693505b505050919050565b60008261567457615674614b9b565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c634300080f000a",
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) Execute(opts *bind.TransactOpts, rep InternalExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "execute", rep, manualExecution)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) Execute(rep InternalExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Execute(&_EVM2EVMOffRampHelper.TransactOpts, rep, manualExecution)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) Execute(rep InternalExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.Execute(&_EVM2EVMOffRampHelper.TransactOpts, rep, manualExecution)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "executeSingleMessage", message, offchainTokenData, manualExecution)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, offchainTokenData [][]byte, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMOffRampHelper.TransactOpts, message, offchainTokenData, manualExecution)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, offchainTokenData [][]byte, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMOffRampHelper.TransactOpts, message, offchainTokenData, manualExecution)
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

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactor) ReleaseOrMintTokens(opts *bind.TransactOpts, sourceTokenAmounts []ClientEVMTokenAmount, originalSender []byte, receiver common.Address, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.contract.Transact(opts, "releaseOrMintTokens", sourceTokenAmounts, originalSender, receiver, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperSession) ReleaseOrMintTokens(sourceTokenAmounts []ClientEVMTokenAmount, originalSender []byte, receiver common.Address, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ReleaseOrMintTokens(&_EVM2EVMOffRampHelper.TransactOpts, sourceTokenAmounts, originalSender, receiver, offchainTokenData)
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelperTransactorSession) ReleaseOrMintTokens(sourceTokenAmounts []ClientEVMTokenAmount, originalSender []byte, receiver common.Address, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMOffRampHelper.Contract.ReleaseOrMintTokens(&_EVM2EVMOffRampHelper.TransactOpts, sourceTokenAmounts, originalSender, receiver, offchainTokenData)
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
	return common.HexToHash("0xb53c6ae8a09cb3c0b6a39aaf24999010a005da24252219c897301f8cd594d601")
}

func (EVM2EVMOffRampHelperConfigSet0) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMOffRampHelperExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x5e04b4755a4460aa6de4f3a906c4324a025c7449c02b52f5466659b5bfdfba5f")
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

func (EVM2EVMOffRampHelperTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (_EVM2EVMOffRampHelper *EVM2EVMOffRampHelper) Address() common.Address {
	return _EVM2EVMOffRampHelper.address
}

type EVM2EVMOffRampHelperInterface interface {
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

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	MetadataHash(opts *bind.CallOpts) ([32]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyPoolUpdates(opts *bind.TransactOpts, removes []InternalPoolUpdate, adds []InternalPoolUpdate) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, rep InternalExecutionReport, manualExecution bool) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte, manualExecution bool) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report InternalExecutionReport) (*types.Transaction, error)

	ReleaseOrMintTokens(opts *bind.TransactOpts, sourceTokenAmounts []ClientEVMTokenAmount, originalSender []byte, receiver common.Address, offchainTokenData [][]byte) (*types.Transaction, error)

	Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error)

	SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetExecutionStateHelper(opts *bind.TransactOpts, sequenceNumber uint64, state uint8) (*types.Transaction, error)

	SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error)

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

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMOffRampHelperTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampHelperTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMOffRampHelperTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
