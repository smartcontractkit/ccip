// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_multi_offramp

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

type EVM2EVMMultiOffRampDynamicConfig struct {
	PermissionLessExecutionThresholdSeconds uint32
	Router                                  common.Address
	MaxNumberOfTokensPerMsg                 uint16
	MaxDataBytes                            uint32
	MaxPoolReleaseOrMintGas                 uint32
	MessageValidator                        common.Address
}

type EVM2EVMMultiOffRampSourceChainConfig struct {
	IsEnabled    bool
	PrevOffRamp  common.Address
	OnRamp       common.Address
	MetadataHash [32]byte
}

type EVM2EVMMultiOffRampSourceChainConfigArgs struct {
	SourceChainSelector uint64
	IsEnabled           bool
	PrevOffRamp         common.Address
	OnRamp              common.Address
}

type EVM2EVMMultiOffRampStaticConfig struct {
	CommitStore   common.Address
	ChainSelector uint64
	RmnProxy      common.Address
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

type InternalExecutionReportSingleChain struct {
	SourceChainSelector uint64
	Messages            []InternalEVM2EVMMessage
	OffchainTokenData   [][][]byte
	Proofs              [][32]byte
	ProofFlagBits       *big.Int
}

var EVM2EVMMultiOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"InvalidManualExecutionGasLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"InvalidMessageId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"InvalidNewState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidStaticConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionGasLimitMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"errorReason\",\"type\":\"bytes\"}],\"name\":\"MessageValidationError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notPool\",\"type\":\"address\"}],\"name\":\"NotACompatiblePool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"TokenDataMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"TokenHandlingError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedTokenData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"SkippedAlreadyExecutedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedSenderWithPreviousRampMessageInflight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfig\",\"name\":\"sourceConfig\",\"type\":\"tuple\"}],\"name\":\"SourceChainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainSelectorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\"}],\"name\":\"applySourceChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"getSourceChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReportSingleChain[]\",\"name\":\"reports\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"gasLimitOverrides\",\"type\":\"uint256[][]\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162005e3438038062005e34833981016040819052620000359162000679565b33806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf816200012c565b5050466080525081516001600160a01b0316620000ef576040516342bcdf7f60e11b815260040160405180910390fd5b81516001600160a01b0390811660a05260208301516001600160401b031660c05260408301511660e0526200012481620001d7565b505062000851565b336001600160a01b03821603620001865760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60005b81518110156200053b576000828281518110620001fb57620001fb620007e4565b60200260200101519050600081600001519050806001600160401b0316600003620002455760405163c39a620560e01b81526001600160401b038216600482015260240162000083565b60608201516001600160a01b031662000271576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160401b038116600090815260096020526040902060018101546001600160a01b03166200043a5760a0516040516374eb454760e11b81526001600160401b03841660048201526000916001600160a01b03169063e9d68a8e90602401606060405180830381865afa158015620002ef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003159190620007fa565b905083606001516001600160a01b031681604001516001600160a01b03161415806200034e5750600181602001516001600160401b0316115b15620003795760405163c39a620560e01b81526001600160401b038416600482015260240162000083565b620003b08385606001517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b36200053f60201b60201c565b600283015560608401516001830180546001600160a01b0319166001600160a01b039283161790556040808601518454610100600160a81b0319166101009190931602919091178355516001600160401b03841681527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb99060200160405180910390a150620004a1565b606083015160018201546001600160a01b0390811691161415806200047657506040830151815461010090046001600160a01b03908116911614155b15620004a15760405163c39a620560e01b81526001600160401b038316600482015260240162000083565b6020830151815490151560ff199091161781556040516001600160401b038316907fdba8597411dc0624375cfff476f6173674609571f4d98d294dd3a47af07927849062000524908490815460ff81161515825260081c6001600160a01b0390811660208301526001830154166040820152600290910154606082015260800190565b60405180910390a2505050806001019050620001da565b5050565b60c05160408051602081018490526001600160401b0380871692820192909252911660608201526001600160a01b038316608082015260009060a0016040516020818303038152906040528051906020012090509392505050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715620005d557620005d56200059a565b60405290565b604051608081016001600160401b0381118282101715620005d557620005d56200059a565b604051601f8201601f191681016001600160401b03811182821017156200062b576200062b6200059a565b604052919050565b80516001600160a01b03811681146200064b57600080fd5b919050565b80516001600160401b03811681146200064b57600080fd5b805180151581146200064b57600080fd5b6000808284036080808212156200068f57600080fd5b6060808312156200069f57600080fd5b620006a9620005b0565b9250620006b68662000633565b83526020620006c781880162000650565b818501526040620006db6040890162000633565b604086015260608801519496506001600160401b0380861115620006fe57600080fd5b858901955089601f8701126200071357600080fd5b8551818111156200072857620007286200059a565b62000738848260051b0162000600565b818152848101925060079190911b87018401908b8211156200075957600080fd5b968401965b81881015620007d25786888d031215620007785760008081fd5b62000782620005db565b6200078d8962000650565b81526200079c868a0162000668565b86820152620007ad858a0162000633565b85820152620007be878a0162000633565b81880152835296860196918401916200075e565b80985050505050505050509250929050565b634e487b7160e01b600052603260045260246000fd5b6000606082840312156200080d57600080fd5b62000817620005b0565b620008228362000668565b8152620008326020840162000650565b6020820152620008456040840162000633565b60408201529392505050565b60805160a05160c05160e051615563620008d1600039600081816101d4015281816117ac01526124ec0152600081816101a4015281816117860152612f320152600081816101680152818161175801528181611aad01526127e1015260008181610a5a01528181610aa601528181610eff0152610f4b01526155636000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c806381ff7048116100b2578063afcb95d711610081578063e9d68a8e11610066578063e9d68a8e14610491578063f2fde38b14610583578063f52121a51461059657600080fd5b8063afcb95d71461045e578063b1dc65a41461047e57600080fd5b806381ff7048146103cc57806385572ffb146103fc5780638b3643341461040a5780638da5cb5b1461043657600080fd5b80635e36480c116101095780637437ff9f116100ee5780637437ff9f146102cd57806379ba5097146103b15780637f63b711146103b957600080fd5b80635e36480c14610298578063666cab8d146102b857600080fd5b806306285c691461013b578063181f5a77146102275780631ef3817414610270578063542625af14610285575b600080fd5b610211604080516060810182526000808252602082018190529181019190915260405180606001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815250905090565b60405161021e9190613a81565b60405180910390f35b6102636040518060400160405280601d81526020017f45564d3245564d4d756c74694f666652616d7020312e362e302d64657600000081525081565b60405161021e9190613b34565b61028361027e366004613e12565b6105a9565b005b6102836102933660046143d8565b610a57565b6102ab6102a6366004614503565b610c7f565b60405161021e91906145a6565b6102c0610d13565b60405161021e9190614606565b6103a46040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260075463ffffffff808216835264010000000080830473ffffffffffffffffffffffffffffffffffffffff90811660208601527801000000000000000000000000000000000000000000000000840461ffff16958501959095527a0100000000000000000000000000000000000000000000000000009092048116606084015260085490811660808401520490911660a082015290565b60405161021e9190614619565b610283610d82565b6102836103c7366004614688565b610e7f565b6004546002546040805163ffffffff8085168252640100000000909404909316602084015282015260600161021e565b61028361013636600461476c565b61041d6104183660046147a7565b610e93565b60405167ffffffffffffffff909116815260200161021e565b60005460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161021e565b60408051600181526000602082018190529181019190915260600161021e565b61028361048c36600461481a565b610ea9565b61053361049f3660046148ff565b6040805160808101825260008082526020820181905291810182905260608101919091525067ffffffffffffffff166000908152600960209081526040918290208251608081018452815460ff81161515825273ffffffffffffffffffffffffffffffffffffffff610100909104811693820193909352600182015490921692820192909252600290910154606082015290565b6040805182511515815260208084015173ffffffffffffffffffffffffffffffffffffffff908116918301919091528383015116918101919091526060918201519181019190915260800161021e565b61028361059136600461491c565b61113a565b6102836105a4366004614939565b61114b565b84518460ff16601f82111561061f576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f746f6f206d616e79207472616e736d697474657273000000000000000000000060448201526064015b60405180910390fd5b80600003610689576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610616565b610691611544565b61069a856115c7565b60065460005b8181101561071e5760056000600683815481106106bf576106bf61499d565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001690556001016106a0565b50875160005b818110156109145760008a82815181106107405761074061499d565b602002602001015190506000600281111561075d5761075d61453c565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260056020526040902054610100900460ff16600281111561079c5761079c61453c565b14610803576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610616565b73ffffffffffffffffffffffffffffffffffffffff8116610850576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff83168152602081016002905273ffffffffffffffffffffffffffffffffffffffff821660009081526005602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156109005761090061453c565b021790555090505050806001019050610724565b5088516109289060069060208c01906139eb565b506003805460ff838116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909216908b1617179055600480546109ae9146913091906000906109809063ffffffff166149fb565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168d8d8d8d8d8d61180b565b600260000181905550600060048054906101000a900463ffffffff169050436004806101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581600260000154600460009054906101000a900463ffffffff168e8e8e8e8e8e604051610a4299989796959493929190614a1e565b60405180910390a15050505050505050505050565b467f000000000000000000000000000000000000000000000000000000000000000014610ae2576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015267ffffffffffffffff46166024820152604401610616565b815181518114610b1e576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b81811015610c6f576000848281518110610b3d57610b3d61499d565b60200260200101519050600081602001515190506000858481518110610b6557610b6561499d565b6020026020010151905080518214610ba9576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b82811015610c60576000828281518110610bc857610bc861499d565b6020026020010151905080600014158015610c03575084602001518281518110610bf457610bf461499d565b60200260200101516080015181105b15610c575784516040517fc8e9605100000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024810183905260448101829052606401610616565b50600101610bac565b50505050806001019050610b21565b50610c7a83836118b6565b505050565b6000610c8d60016004614ab4565b6002610c9a608085614af6565b67ffffffffffffffff16610cae9190614b1d565b67ffffffffffffffff85166000908152600b6020526040812090610cd3608087614b34565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002054901c166003811115610d0a57610d0a61453c565b90505b92915050565b60606006805480602002602001604051908101604052809291908181526020018280548015610d7857602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d4d575b5050505050905090565b60015473ffffffffffffffffffffffffffffffffffffffff163314610e03576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610616565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610e87611544565b610e9081611966565b50565b600080610ea08484611e18565b50949350505050565b610eb38787611f48565b600254883590808214610efc576040517f93df584c0000000000000000000000000000000000000000000000000000000081526004810182905260248101839052604401610616565b467f000000000000000000000000000000000000000000000000000000000000000014610f7d576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006004820152466024820152604401610616565b6040805183815260208c81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a13360009081526005602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156110055761100561453c565b60028111156110165761101661453c565b90525090506002816020015160028111156110335761103361453c565b14801561107a57506006816000015160ff16815481106110555761105561499d565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b6110b0576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060006110be856020614b1d565b6110c9886020614b1d565b6110d58b610144614b5b565b6110df9190614b5b565b6110e99190614b5b565b905036811461112d576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401610616565b5050505050505050505050565b611142611544565b610e9081611f8f565b333014611184576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051600080825260208201909252816111c1565b604080518082019091526000808252602082015281526020019060019003908161119a5790505b50610140840151519091501561127a576101408301516040805160608101909152602085015173ffffffffffffffffffffffffffffffffffffffff16608082015261127791908060a08101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152908252875167ffffffffffffffff1660208301528781015173ffffffffffffffffffffffffffffffffffffffff1691015261016086015185612084565b90505b600061128684836123e6565b600854909150640100000000900473ffffffffffffffffffffffffffffffffffffffff168015611395576040517fa219f6e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82169063a219f6e590611302908590600401614c2f565b600060405180830381600087803b15801561131c57600080fd5b505af192505050801561132d575060015b611395573d80801561135b576040519150601f19603f3d011682016040523d82523d6000602084013e611360565b606091505b50806040517f09c253250000000000000000000000000000000000000000000000000000000081526004016106169190613b34565b610120850151511580156113ab57506080850151155b806113cf5750604085015173ffffffffffffffffffffffffffffffffffffffff163b155b8061141c5750604085015161141a9073ffffffffffffffffffffffffffffffffffffffff167f85572ffb00000000000000000000000000000000000000000000000000000000612496565b155b15611428575050505050565b600754608086015160408088015190517f3cf97983000000000000000000000000000000000000000000000000000000008152600093849364010000000090910473ffffffffffffffffffffffffffffffffffffffff1692633cf979839261149892899261138892600401614c42565b6000604051808303816000875af11580156114b7573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526114fd9190810190614cd0565b50915091508161153b57806040517f0a8d6e8c0000000000000000000000000000000000000000000000000000000081526004016106169190613b34565b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146115c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610616565b565b6000818060200190518101906115dd9190614d3e565b602081015190915073ffffffffffffffffffffffffffffffffffffffff16611631576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516007805460208085015160408087015160608089015163ffffffff9081167a010000000000000000000000000000000000000000000000000000027fffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffffff61ffff909416780100000000000000000000000000000000000000000000000002939093167fffff000000000000ffffffffffffffffffffffffffffffffffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff9687166401000000009081027fffffffffffffffff0000000000000000000000000000000000000000000000009a8b169c85169c909c179b909b1716179290921790965560808801516008805460a08b015186169099029890961691161795909517909255835192830184527f00000000000000000000000000000000000000000000000000000000000000008216835267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016908301527f0000000000000000000000000000000000000000000000000000000000000000168183015290517f1c86233479b9d798a31091d42b1a954067dcd6c898553d63bf993757d2a90aff916117ff918490614dea565b60405180910390a15050565b6000808a8a8a8a8a8a8a8a8a60405160200161182f99989796959493929190614e9e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b81516000036118f0576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160408051600080825260208201909252911591905b845181101561195f576119578582815181106119255761192561499d565b602002602001015184611951578583815181106119445761194461499d565b60200260200101516124b2565b836124b2565b600101611907565b5050505050565b60005b8151811015611e145760008282815181106119865761198661499d565b602002602001015190506000816000015190508067ffffffffffffffff166000036119e9576040517fc39a620500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610616565b606082015173ffffffffffffffffffffffffffffffffffffffff16611a3a576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff81166000908152600960205260409020600181015473ffffffffffffffffffffffffffffffffffffffff16611cb8576040517fe9d68a8e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff831660048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063e9d68a8e90602401606060405180830381865afa158015611b09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b2d9190614f33565b9050836060015173ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff16141580611b8057506001816020015167ffffffffffffffff16115b15611bc3576040517fc39a620500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610616565b611bf28385606001517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3612f2c565b600283015560608401516001830180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92831617905560408086015184547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010091909316029190911783555167ffffffffffffffff841681527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb99060200160405180910390a150611d50565b6060830151600182015473ffffffffffffffffffffffffffffffffffffffff9081169116141580611d0d575060408301518154610100900473ffffffffffffffffffffffffffffffffffffffff908116911614155b15611d50576040517fc39a620500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff83166004820152602401610616565b602083015181549015157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090911617815560405167ffffffffffffffff8316907fdba8597411dc0624375cfff476f6173674609571f4d98d294dd3a47af079278490611dfe908490815460ff81161515825260081c73ffffffffffffffffffffffffffffffffffffffff90811660208301526001830154166040820152600290910154606082015260800190565b60405180910390a2505050806001019050611969565b5050565b67ffffffffffffffff8083166000908152600a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff861684529091528120549091829116808203611f3a5767ffffffffffffffff8516600090815260096020526040902054610100900473ffffffffffffffffffffffffffffffffffffffff168015611f38576040517f856c824700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015282169063856c824790602401602060405180830381865afa158015611f07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f2b9190614f87565b6001935093505050611f41565b505b9150600090505b9250929050565b6000611f5682840184614fa4565b60408051600080825260208201909252919250610c7a918391611f89565b6060815260200190600190039081611f745790505b506118b6565b3373ffffffffffffffffffffffffffffffffffffffff82160361200e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610616565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8360005b8551811015610ea05760008482815181106120a5576120a561499d565b60200260200101518060200190518101906120c09190614fd9565b905060006120d18260200151612fbc565b905061211373ffffffffffffffffffffffffffffffffffffffff82167faff2afbf00000000000000000000000000000000000000000000000000000000612496565b612161576040517fae9b4ce900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610616565b6000806122ae634059f55b60e01b6040518060e001604052808c6000015181526020018c6020015167ffffffffffffffff1681526020018c6040015173ffffffffffffffffffffffffffffffffffffffff1681526020018d89815181106121ca576121ca61499d565b602002602001015160200151815260200187600001518152602001876040015181526020018a89815181106122015761220161499d565b602002602001015181525060405160240161221c919061508e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152600854859063ffffffff166113886084613017565b5091509150816122ec57806040517fe1cd55090000000000000000000000000000000000000000000000000000000081526004016106169190613b34565b8051604014612336578051604080517f78ef802400000000000000000000000000000000000000000000000000000000815260048101919091526024810191909152604401610616565b6000808280602001905181019061234d9190615168565b9150915061235a8261313d565b88888151811061236c5761236c61499d565b60200260200101516000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808888815181106123bd576123bd61499d565b60200260200101516020018181525050505050505050806001019050612088565b949350505050565b6040805160a08101825260008082526020820152606091810182905281810182905260808101919091526040518060a001604052808461018001518152602001846000015167ffffffffffffffff168152602001846020015160405160200161246b919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b6040516020818303038152906040528152602001846101200151815260200183815250905092915050565b60006124a1836131d4565b8015610d0a5750610d0a8383613238565b81516040517f58babe3300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906358babe3390602401602060405180830381865afa158015612548573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061256c919061518c565b156125af576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610616565b60208301515160008190036125ef576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836040015151811461262d576040517f57e0e08300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff82166000908152600960205260409020805460ff1661268d576040517fed053c5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610616565b60008267ffffffffffffffff8111156126a8576126a8613b47565b6040519080825280602002602001820160405280156126d1578160200160208202803683370190505b50905060005b83811015612796576000876020015182815181106126f7576126f761499d565b6020026020010151905061270f818560020154613307565b8383815181106127215761272161499d565b6020026020010181815250508061018001518383815181106127455761274561499d565b60200260200101511461278d578061018001516040517f345039be00000000000000000000000000000000000000000000000000000000815260040161061691815260200190565b506001016126d7565b50606086015160808701516040517ffe41448f00000000000000000000000000000000000000000000000000000000815260009273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169263fe41448f92612818928a9288926004016151da565b602060405180830381865afa158015612835573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128599190615221565b9050806000036128a1576040517f7dd17a7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610616565b8551151560005b85811015612f21576000896020015182815181106128c8576128c861499d565b6020026020010151905060006128e2898360600151610c7f565b905060028160038111156128f8576128f861453c565b0361294e5760608201516040805167ffffffffffffffff808d16825290921660208301527f3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c910160405180910390a15050612f19565b60008160038111156129625761296261453c565b148061297f5750600381600381111561297d5761297d61453c565b145b6129cf5760608201516040517f25507e7f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808c1660048301529091166024820152604401610616565b8315612a985760075460009063ffffffff166129eb8742614ab4565b1190508080612a0b57506003826003811115612a0957612a0961453c565b145b612a4d576040517fa9cfc86200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b166004820152602401610616565b8a8481518110612a5f57612a5f61499d565b6020026020010151600014612a92578a8481518110612a8057612a8061499d565b60200260200101518360800181815250505b50612afd565b6000816003811115612aac57612aac61453c565b14612afd5760608201516040517f3ef2a99c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808c1660048301529091166024820152604401610616565b600080612b0e8b8560200151611e18565b915091508015612c2e5760c084015167ffffffffffffffff16612b3283600161523a565b67ffffffffffffffff1614612bc25760c084015160208501516040517f5444a3301c7c42dd164cbf6ba4b72bf02504f86c049b06a27fc2b662e334bdbd92612bb1928f9267ffffffffffffffff938416815291909216602082015273ffffffffffffffffffffffffffffffffffffffff91909116604082015260600190565b60405180910390a150505050612f19565b67ffffffffffffffff8b81166000908152600a602090815260408083208883015173ffffffffffffffffffffffffffffffffffffffff168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000169184169190911790555b6000836003811115612c4257612c4261453c565b03612ce05760c084015167ffffffffffffffff16612c6183600161523a565b67ffffffffffffffff1614612ce05760c084015160208501516040517f852dc8e405695593e311bd83991cf39b14a328f304935eac6d3d55617f911d8992612bb1928f9267ffffffffffffffff938416815291909216602082015273ffffffffffffffffffffffffffffffffffffffff91909116604082015260600190565b60008d604001518681518110612cf857612cf861499d565b60200260200101519050612d268561018001518d87606001518861014001515189610120015151865161348d565b612d368c866060015160016135e9565b600080612d4387846136c7565b91509150612d568e8860600151846135e9565b888015612d7457506003826003811115612d7257612d7261453c565b145b15612db457866101800151816040517f2b11b8d9000000000000000000000000000000000000000000000000000000008152600401610616929190615262565b6003826003811115612dc857612dc861453c565b14158015612de857506002826003811115612de557612de561453c565b14155b15612e29578d8760600151836040517f926c5a3e0000000000000000000000000000000000000000000000000000000081526004016106169392919061527b565b6000866003811115612e3d57612e3d61453c565b03612eb85767ffffffffffffffff808f166000908152600a602090815260408083208b83015173ffffffffffffffffffffffffffffffffffffffff168452909152812080549092169190612e90836152a1565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b866101800151876060015167ffffffffffffffff168f67ffffffffffffffff167f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df28585604051612f099291906152be565b60405180910390a4505050505050505b6001016128a8565b505050505050505050565b600081847f000000000000000000000000000000000000000000000000000000000000000085604051602001612f9c949392919093845267ffffffffffffffff92831660208501529116604083015273ffffffffffffffffffffffffffffffffffffffff16606082015260800190565b6040516020818303038152906040528051906020012090505b9392505050565b60008151602014612ffb57816040517f8d666f600000000000000000000000000000000000000000000000000000000081526004016106169190613b34565b610d0d828060200190518101906130129190615221565b61313d565b6000606060008361ffff1667ffffffffffffffff81111561303a5761303a613b47565b6040519080825280601f01601f191660200182016040528015613064576020820181803683370190505b509150863b613097577f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b5a858110156130ca577fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b8590036040810481038710613103577f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b505a6000808a5160208c0160008c8cf193505a900390503d848111156131265750835b808352806000602085013e50955095509592505050565b600073ffffffffffffffffffffffffffffffffffffffff8211806131615750600a82105b156131d057604080516020810184905201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f8d666f6000000000000000000000000000000000000000000000000000000000825261061691600401613b34565b5090565b6000613200827f01ffc9a700000000000000000000000000000000000000000000000000000000613238565b8015610d0d5750613231827fffffffff00000000000000000000000000000000000000000000000000000000613238565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d915060005190508280156132f0575060208210155b80156132fc5750600081115b979650505050505050565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b61010001516040516020016133aa98979695949392919073ffffffffffffffffffffffffffffffffffffffff9889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b60405160208183030381529060405280519060200120856101200151805190602001208661014001516040516020016133e391906152de565b6040516020818303038152906040528051906020012087610160015160405160200161340f9190615369565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b6007547801000000000000000000000000000000000000000000000000900461ffff168311156134fd576040517fa1e5205a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808716600483015285166024820152604401610616565b80831461354a576040517f1cfe6d8b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808716600483015285166024820152604401610616565b6007547a010000000000000000000000000000000000000000000000000000900463ffffffff168211156135e1576007546040517f1fd8fd04000000000000000000000000000000000000000000000000000000008152600481018890527a01000000000000000000000000000000000000000000000000000090910463ffffffff16602482015260448101839052606401610616565b505050505050565b600060026135f8608085614af6565b67ffffffffffffffff1661360c9190614b1d565b67ffffffffffffffff85166000908152600b602052604081209192509081613635608087614b34565b67ffffffffffffffff16815260208101919091526040016000205490508161365f60016004614ab4565b901b1916818360038111156136765761367661453c565b67ffffffffffffffff87166000908152600b602052604081209190921b929092179182916136a5608088614b34565b67ffffffffffffffff1681526020810191909152604001600020555050505050565b6040517ff52121a5000000000000000000000000000000000000000000000000000000008152600090606090309063f52121a59061370b908790879060040161537c565b600060405180830381600087803b15801561372557600080fd5b505af1925050508015613736575060015b6139d0573d808015613764576040519150601f19603f3d011682016040523d82523d6000602084013e613769565b606091505b50600061377582615506565b90507f0a8d6e8c000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216148061380857507fe1cd5509000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b8061385457507f8d666f60000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b806138a057507f78ef8024000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b806138ec57507f0c3b563c000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b8061393857507fae9b4ce9000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b8061398457507f09c25325000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b156139955750600392509050611f41565b856101800151826040517f2b11b8d9000000000000000000000000000000000000000000000000000000008152600401610616929190615262565b50506040805160208101909152600081526002909250929050565b828054828255906000526020600020908101928215613a65579160200282015b82811115613a6557825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190613a0b565b506131d09291505b808211156131d05760008155600101613a6d565b60608101610d0d8284805173ffffffffffffffffffffffffffffffffffffffff908116835260208083015167ffffffffffffffff169084015260409182015116910152565b60005b83811015613ae1578181015183820152602001613ac9565b50506000910152565b60008151808452613b02816020860160208601613ac6565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610d0a6020830184613aea565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613b9957613b99613b47565b60405290565b6040516101a0810167ffffffffffffffff81118282101715613b9957613b99613b47565b60405160a0810167ffffffffffffffff81118282101715613b9957613b99613b47565b6040516080810167ffffffffffffffff81118282101715613b9957613b99613b47565b6040516060810167ffffffffffffffff81118282101715613b9957613b99613b47565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613c7357613c73613b47565b604052919050565b600067ffffffffffffffff821115613c9557613c95613b47565b5060051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81168114610e9057600080fd5b8035613ccc81613c9f565b919050565b600082601f830112613ce257600080fd5b81356020613cf7613cf283613c7b565b613c2c565b8083825260208201915060208460051b870101935086841115613d1957600080fd5b602086015b84811015613d3e578035613d3181613c9f565b8352918301918301613d1e565b509695505050505050565b803560ff81168114613ccc57600080fd5b600067ffffffffffffffff821115613d7457613d74613b47565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112613db157600080fd5b8135613dbf613cf282613d5a565b818152846020838601011115613dd457600080fd5b816020850160208301376000918101602001919091529392505050565b67ffffffffffffffff81168114610e9057600080fd5b8035613ccc81613df1565b60008060008060008060c08789031215613e2b57600080fd5b863567ffffffffffffffff80821115613e4357600080fd5b613e4f8a838b01613cd1565b97506020890135915080821115613e6557600080fd5b613e718a838b01613cd1565b9650613e7f60408a01613d49565b95506060890135915080821115613e9557600080fd5b613ea18a838b01613da0565b9450613eaf60808a01613e07565b935060a0890135915080821115613ec557600080fd5b50613ed289828a01613da0565b9150509295509295509295565b8015158114610e9057600080fd5b8035613ccc81613edf565b600082601f830112613f0957600080fd5b81356020613f19613cf283613c7b565b82815260069290921b84018101918181019086841115613f3857600080fd5b8286015b84811015613d3e5760408189031215613f555760008081fd5b613f5d613b76565b8135613f6881613c9f565b81528185013585820152835291830191604001613f3c565b600082601f830112613f9157600080fd5b81356020613fa1613cf283613c7b565b82815260059290921b84018101918181019086841115613fc057600080fd5b8286015b84811015613d3e57803567ffffffffffffffff811115613fe45760008081fd5b613ff28986838b0101613da0565b845250918301918301613fc4565b60006101a0828403121561401357600080fd5b61401b613b9f565b905061402682613e07565b815261403460208301613cc1565b602082015261404560408301613cc1565b604082015261405660608301613e07565b60608201526080820135608082015261407160a08301613eed565b60a082015261408260c08301613e07565b60c082015261409360e08301613cc1565b60e082015261010082810135908201526101208083013567ffffffffffffffff808211156140c057600080fd5b6140cc86838701613da0565b838501526101409250828501359150808211156140e857600080fd5b6140f486838701613ef8565b8385015261016092508285013591508082111561411057600080fd5b5061411d85828601613f80565b82840152505061018080830135818301525092915050565b600082601f83011261414657600080fd5b81356020614156613cf283613c7b565b82815260059290921b8401810191818101908684111561417557600080fd5b8286015b84811015613d3e57803567ffffffffffffffff8111156141995760008081fd5b6141a78986838b0101614000565b845250918301918301614179565b600082601f8301126141c657600080fd5b813560206141d6613cf283613c7b565b82815260059290921b840181019181810190868411156141f557600080fd5b8286015b84811015613d3e57803567ffffffffffffffff8111156142195760008081fd5b6142278986838b0101613f80565b8452509183019183016141f9565b600082601f83011261424657600080fd5b81356020614256613cf283613c7b565b8083825260208201915060208460051b87010193508684111561427857600080fd5b602086015b84811015613d3e578035835291830191830161427d565b600082601f8301126142a557600080fd5b813560206142b5613cf283613c7b565b82815260059290921b840181019181810190868411156142d457600080fd5b8286015b84811015613d3e57803567ffffffffffffffff808211156142f95760008081fd5b818901915060a0807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848d030112156143325760008081fd5b61433a613bc3565b614345888501613e07565b81526040808501358481111561435b5760008081fd5b6143698e8b83890101614135565b8a84015250606080860135858111156143825760008081fd5b6143908f8c838a01016141b5565b83850152506080915081860135858111156143ab5760008081fd5b6143b98f8c838a0101614235565b91840191909152509190930135908301525083529183019183016142d8565b60008060408084860312156143ec57600080fd5b833567ffffffffffffffff8082111561440457600080fd5b61441087838801614294565b945060209150818601358181111561442757600080fd5b8601601f8101881361443857600080fd5b8035614446613cf282613c7b565b81815260059190911b8201840190848101908a83111561446557600080fd5b8584015b838110156144f1578035868111156144815760008081fd5b8501603f81018d136144935760008081fd5b878101356144a3613cf282613c7b565b81815260059190911b82018a0190898101908f8311156144c35760008081fd5b928b01925b828410156144e15783358252928a0192908a01906144c8565b8652505050918601918601614469565b50809750505050505050509250929050565b6000806040838503121561451657600080fd5b823561452181613df1565b9150602083013561453181613df1565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600481106145a2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b60208101610d0d828461456b565b60008151808452602080850194506020840160005b838110156145fb57815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016145c9565b509495945050505050565b602081526000610d0a60208301846145b4565b60c08101610d0d828463ffffffff808251168352602082015173ffffffffffffffffffffffffffffffffffffffff808216602086015261ffff60408501511660408601528260608501511660608601528260808501511660808601528060a08501511660a08601525050505050565b6000602080838503121561469b57600080fd5b823567ffffffffffffffff8111156146b257600080fd5b8301601f810185136146c357600080fd5b80356146d1613cf282613c7b565b81815260079190911b820183019083810190878311156146f057600080fd5b928401925b828410156132fc576080848903121561470e5760008081fd5b614716613be6565b843561472181613df1565b81528486013561473081613edf565b8187015260408581013561474381613c9f565b9082015260608581013561475681613c9f565b90820152825260809390930192908401906146f5565b60006020828403121561477e57600080fd5b813567ffffffffffffffff81111561479557600080fd5b820160a08185031215612fb557600080fd5b600080604083850312156147ba57600080fd5b82356147c581613df1565b9150602083013561453181613c9f565b60008083601f8401126147e757600080fd5b50813567ffffffffffffffff8111156147ff57600080fd5b6020830191508360208260051b8501011115611f4157600080fd5b60008060008060008060008060e0898b03121561483657600080fd5b606089018a81111561484757600080fd5b8998503567ffffffffffffffff8082111561486157600080fd5b818b0191508b601f83011261487557600080fd5b81358181111561488457600080fd5b8c602082850101111561489657600080fd5b6020830199508098505060808b01359150808211156148b457600080fd5b6148c08c838d016147d5565b909750955060a08b01359150808211156148d957600080fd5b506148e68b828c016147d5565b999c989b50969995989497949560c00135949350505050565b60006020828403121561491157600080fd5b8135612fb581613df1565b60006020828403121561492e57600080fd5b8135612fb581613c9f565b6000806040838503121561494c57600080fd5b823567ffffffffffffffff8082111561496457600080fd5b61497086838701614000565b9350602085013591508082111561498657600080fd5b5061499385828601613f80565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600063ffffffff808316818103614a1457614a146149cc565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614a4e8184018a6145b4565b90508281036080840152614a6281896145b4565b905060ff871660a084015282810360c0840152614a7f8187613aea565b905067ffffffffffffffff851660e0840152828103610100840152614aa48185613aea565b9c9b505050505050505050505050565b81810381811115610d0d57610d0d6149cc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff80841680614b1157614b11614ac7565b92169190910692915050565b8082028115828204841417610d0d57610d0d6149cc565b600067ffffffffffffffff80841680614b4f57614b4f614ac7565b92169190910492915050565b80820180821115610d0d57610d0d6149cc565b60008151808452602080850194506020840160005b838110156145fb578151805173ffffffffffffffffffffffffffffffffffffffff1688528301518388015260409096019590820190600101614b83565b8051825267ffffffffffffffff60208201511660208301526000604082015160a06040850152614bf360a0850182613aea565b905060608301518482036060860152614c0c8282613aea565b91505060808301518482036080860152614c268282614b6e565b95945050505050565b602081526000610d0a6020830184614bc0565b608081526000614c556080830187614bc0565b61ffff95909516602083015250604081019290925273ffffffffffffffffffffffffffffffffffffffff16606090910152919050565b600082601f830112614c9c57600080fd5b8151614caa613cf282613d5a565b818152846020838601011115614cbf57600080fd5b6123de826020830160208701613ac6565b600080600060608486031215614ce557600080fd5b8351614cf081613edf565b602085015190935067ffffffffffffffff811115614d0d57600080fd5b614d1986828701614c8b565b925050604084015190509250925092565b805163ffffffff81168114613ccc57600080fd5b600060c08284031215614d5057600080fd5b60405160c0810181811067ffffffffffffffff82111715614d7357614d73613b47565b604052614d7f83614d2a565b81526020830151614d8f81613c9f565b6020820152604083015161ffff81168114614da957600080fd5b6040820152614dba60608401614d2a565b6060820152614dcb60808401614d2a565b608082015260a0830151614dde81613c9f565b60a08201529392505050565b6101208101614e308285805173ffffffffffffffffffffffffffffffffffffffff908116835260208083015167ffffffffffffffff169084015260409182015116910152565b612fb5606083018463ffffffff808251168352602082015173ffffffffffffffffffffffffffffffffffffffff808216602086015261ffff60408501511660408601528260608501511660608601528260808501511660808601528060a08501511660a08601525050505050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b166040850152816060850152614ee58285018b6145b4565b91508382036080850152614ef9828a6145b4565b915060ff881660a085015283820360c0850152614f168288613aea565b90861660e08501528381036101008501529050614aa48185613aea565b600060608284031215614f4557600080fd5b614f4d613c09565b8251614f5881613edf565b81526020830151614f6881613df1565b60208201526040830151614f7b81613c9f565b60408201529392505050565b600060208284031215614f9957600080fd5b8151612fb581613df1565b600060208284031215614fb657600080fd5b813567ffffffffffffffff811115614fcd57600080fd5b6123de84828501614294565b600060208284031215614feb57600080fd5b815167ffffffffffffffff8082111561500357600080fd5b908301906060828603121561501757600080fd5b61501f613c09565b82518281111561502e57600080fd5b61503a87828601614c8b565b82525060208301518281111561504f57600080fd5b61505b87828601614c8b565b60208301525060408301518281111561507357600080fd5b61507f87828601614c8b565b60408301525095945050505050565b602081526000825160e060208401526150ab610100840182613aea565b905067ffffffffffffffff602085015116604084015260408401516150e8606085018273ffffffffffffffffffffffffffffffffffffffff169052565b506060840151608084015260808401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0808584030160a086015261512d8383613aea565b925060a08601519150808584030160c086015261514a8383613aea565b925060c08601519150808584030160e086015250614c268282613aea565b6000806040838503121561517b57600080fd5b505080516020909101519092909150565b60006020828403121561519e57600080fd5b8151612fb581613edf565b60008151808452602080850194506020840160005b838110156145fb578151875295820195908201906001016151be565b67ffffffffffffffff851681526080602082015260006151fd60808301866151a9565b828103604084015261520f81866151a9565b91505082606083015295945050505050565b60006020828403121561523357600080fd5b5051919050565b67ffffffffffffffff81811683821601908082111561525b5761525b6149cc565b5092915050565b8281526040602082015260006123de6040830184613aea565b67ffffffffffffffff848116825283166020820152606081016123de604083018461456b565b600067ffffffffffffffff808316818103614a1457614a146149cc565b6152c8818461456b565b6040602082015260006123de6040830184613aea565b602081526000610d0a6020830184614b6e565b60008282518085526020808601955060208260051b8401016020860160005b8481101561535c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086840301895261534a838351613aea565b98840198925090830190600101615310565b5090979650505050505050565b602081526000610d0a60208301846152f1565b6040815261539760408201845167ffffffffffffffff169052565b600060208401516153c0606084018273ffffffffffffffffffffffffffffffffffffffff169052565b50604084015173ffffffffffffffffffffffffffffffffffffffff8116608084015250606084015167ffffffffffffffff811660a084015250608084015160c083015260a084015180151560e08401525060c084015161010061542e8185018367ffffffffffffffff169052565b60e086015191506101206154598186018473ffffffffffffffffffffffffffffffffffffffff169052565b81870151925061014091508282860152808701519250506101a061016081818701526154896101e0870185613aea565b93508288015192507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06101808188870301818901526154c88686614b6e565b9550828a015194508188870301848901526154e386866152f1565b9550808a01516101c089015250505050508281036020840152614c2681856152f1565b6000815160208301517fffffffff000000000000000000000000000000000000000000000000000000008082169350600483101561554e5780818460040360031b1b83161693505b50505091905056fea164736f6c6343000818000a",
}

var EVM2EVMMultiOffRampABI = EVM2EVMMultiOffRampMetaData.ABI

var EVM2EVMMultiOffRampBin = EVM2EVMMultiOffRampMetaData.Bin

func DeployEVM2EVMMultiOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMMultiOffRampStaticConfig, sourceChainConfigs []EVM2EVMMultiOffRampSourceChainConfigArgs) (common.Address, *types.Transaction, *EVM2EVMMultiOffRamp, error) {
	parsed, err := EVM2EVMMultiOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMMultiOffRampBin), backend, staticConfig, sourceChainConfigs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMMultiOffRamp{address: address, abi: *parsed, EVM2EVMMultiOffRampCaller: EVM2EVMMultiOffRampCaller{contract: contract}, EVM2EVMMultiOffRampTransactor: EVM2EVMMultiOffRampTransactor{contract: contract}, EVM2EVMMultiOffRampFilterer: EVM2EVMMultiOffRampFilterer{contract: contract}}, nil
}

type EVM2EVMMultiOffRamp struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMMultiOffRampCaller
	EVM2EVMMultiOffRampTransactor
	EVM2EVMMultiOffRampFilterer
}

type EVM2EVMMultiOffRampCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMMultiOffRampTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMMultiOffRampFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMMultiOffRampSession struct {
	Contract     *EVM2EVMMultiOffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMMultiOffRampCallerSession struct {
	Contract *EVM2EVMMultiOffRampCaller
	CallOpts bind.CallOpts
}

type EVM2EVMMultiOffRampTransactorSession struct {
	Contract     *EVM2EVMMultiOffRampTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMMultiOffRampRaw struct {
	Contract *EVM2EVMMultiOffRamp
}

type EVM2EVMMultiOffRampCallerRaw struct {
	Contract *EVM2EVMMultiOffRampCaller
}

type EVM2EVMMultiOffRampTransactorRaw struct {
	Contract *EVM2EVMMultiOffRampTransactor
}

func NewEVM2EVMMultiOffRamp(address common.Address, backend bind.ContractBackend) (*EVM2EVMMultiOffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMMultiOffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMMultiOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRamp{address: address, abi: abi, EVM2EVMMultiOffRampCaller: EVM2EVMMultiOffRampCaller{contract: contract}, EVM2EVMMultiOffRampTransactor: EVM2EVMMultiOffRampTransactor{contract: contract}, EVM2EVMMultiOffRampFilterer: EVM2EVMMultiOffRampFilterer{contract: contract}}, nil
}

func NewEVM2EVMMultiOffRampCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMMultiOffRampCaller, error) {
	contract, err := bindEVM2EVMMultiOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampCaller{contract: contract}, nil
}

func NewEVM2EVMMultiOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMMultiOffRampTransactor, error) {
	contract, err := bindEVM2EVMMultiOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampTransactor{contract: contract}, nil
}

func NewEVM2EVMMultiOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMMultiOffRampFilterer, error) {
	contract, err := bindEVM2EVMMultiOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampFilterer{contract: contract}, nil
}

func bindEVM2EVMMultiOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EVM2EVMMultiOffRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMMultiOffRamp.Contract.EVM2EVMMultiOffRampCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.EVM2EVMMultiOffRampTransactor.contract.Transfer(opts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.EVM2EVMMultiOffRampTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMMultiOffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.contract.Transfer(opts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _EVM2EVMMultiOffRamp.Contract.CcipReceive(&_EVM2EVMMultiOffRamp.CallOpts, arg0)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _EVM2EVMMultiOffRamp.Contract.CcipReceive(&_EVM2EVMMultiOffRamp.CallOpts, arg0)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMMultiOffRampDynamicConfig, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(EVM2EVMMultiOffRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMMultiOffRampDynamicConfig)).(*EVM2EVMMultiOffRampDynamicConfig)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) GetDynamicConfig() (EVM2EVMMultiOffRampDynamicConfig, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetDynamicConfig(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) GetDynamicConfig() (EVM2EVMMultiOffRampDynamicConfig, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetDynamicConfig(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "getExecutionState", sourceChainSelector, sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetExecutionState(&_EVM2EVMMultiOffRamp.CallOpts, sourceChainSelector, sequenceNumber)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetExecutionState(&_EVM2EVMMultiOffRamp.CallOpts, sourceChainSelector, sequenceNumber)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) GetSenderNonce(opts *bind.CallOpts, sourceChainSelector uint64, sender common.Address) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "getSenderNonce", sourceChainSelector, sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) GetSenderNonce(sourceChainSelector uint64, sender common.Address) (uint64, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetSenderNonce(&_EVM2EVMMultiOffRamp.CallOpts, sourceChainSelector, sender)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) GetSenderNonce(sourceChainSelector uint64, sender common.Address) (uint64, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetSenderNonce(&_EVM2EVMMultiOffRamp.CallOpts, sourceChainSelector, sender)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (EVM2EVMMultiOffRampSourceChainConfig, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "getSourceChainConfig", sourceChainSelector)

	if err != nil {
		return *new(EVM2EVMMultiOffRampSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMMultiOffRampSourceChainConfig)).(*EVM2EVMMultiOffRampSourceChainConfig)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) GetSourceChainConfig(sourceChainSelector uint64) (EVM2EVMMultiOffRampSourceChainConfig, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetSourceChainConfig(&_EVM2EVMMultiOffRamp.CallOpts, sourceChainSelector)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) GetSourceChainConfig(sourceChainSelector uint64) (EVM2EVMMultiOffRampSourceChainConfig, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetSourceChainConfig(&_EVM2EVMMultiOffRamp.CallOpts, sourceChainSelector)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) GetStaticConfig(opts *bind.CallOpts) (EVM2EVMMultiOffRampStaticConfig, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(EVM2EVMMultiOffRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(EVM2EVMMultiOffRampStaticConfig)).(*EVM2EVMMultiOffRampStaticConfig)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) GetStaticConfig() (EVM2EVMMultiOffRampStaticConfig, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetStaticConfig(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) GetStaticConfig() (EVM2EVMMultiOffRampStaticConfig, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetStaticConfig(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) GetTransmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "getTransmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) GetTransmitters() ([]common.Address, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetTransmitters(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) GetTransmitters() ([]common.Address, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetTransmitters(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMMultiOffRamp.Contract.LatestConfigDetails(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMMultiOffRamp.Contract.LatestConfigDetails(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMMultiOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMMultiOffRamp.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) Owner() (common.Address, error) {
	return _EVM2EVMMultiOffRamp.Contract.Owner(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMMultiOffRamp.Contract.Owner(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) TypeAndVersion() (string, error) {
	return _EVM2EVMMultiOffRamp.Contract.TypeAndVersion(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMMultiOffRamp.Contract.TypeAndVersion(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.AcceptOwnership(&_EVM2EVMMultiOffRamp.TransactOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.AcceptOwnership(&_EVM2EVMMultiOffRamp.TransactOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []EVM2EVMMultiOffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "applySourceChainConfigUpdates", sourceChainConfigUpdates)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []EVM2EVMMultiOffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.ApplySourceChainConfigUpdates(&_EVM2EVMMultiOffRamp.TransactOpts, sourceChainConfigUpdates)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []EVM2EVMMultiOffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.ApplySourceChainConfigUpdates(&_EVM2EVMMultiOffRamp.TransactOpts, sourceChainConfigUpdates)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "executeSingleMessage", message, offchainTokenData)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMMultiOffRamp.TransactOpts, message, offchainTokenData)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMMultiOffRamp.TransactOpts, message, offchainTokenData)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) ManuallyExecute(opts *bind.TransactOpts, reports []InternalExecutionReportSingleChain, gasLimitOverrides [][]*big.Int) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "manuallyExecute", reports, gasLimitOverrides)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) ManuallyExecute(reports []InternalExecutionReportSingleChain, gasLimitOverrides [][]*big.Int) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.ManuallyExecute(&_EVM2EVMMultiOffRamp.TransactOpts, reports, gasLimitOverrides)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) ManuallyExecute(reports []InternalExecutionReportSingleChain, gasLimitOverrides [][]*big.Int) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.ManuallyExecute(&_EVM2EVMMultiOffRamp.TransactOpts, reports, gasLimitOverrides)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "setOCR2Config", signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.SetOCR2Config(&_EVM2EVMMultiOffRamp.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) SetOCR2Config(signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.SetOCR2Config(&_EVM2EVMMultiOffRamp.TransactOpts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.TransferOwnership(&_EVM2EVMMultiOffRamp.TransactOpts, to)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.TransferOwnership(&_EVM2EVMMultiOffRamp.TransactOpts, to)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "transmit", reportContext, report, rs, ss, arg4)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.Transmit(&_EVM2EVMMultiOffRamp.TransactOpts, reportContext, report, rs, ss, arg4)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.Transmit(&_EVM2EVMMultiOffRamp.TransactOpts, reportContext, report, rs, ss, arg4)
}

type EVM2EVMMultiOffRampConfigSetIterator struct {
	Event *EVM2EVMMultiOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampConfigSet)
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
		it.Event = new(EVM2EVMMultiOffRampConfigSet)
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

func (it *EVM2EVMMultiOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampConfigSet struct {
	StaticConfig  EVM2EVMMultiOffRampStaticConfig
	DynamicConfig EVM2EVMMultiOffRampDynamicConfig
	Raw           types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampConfigSetIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampConfigSet)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseConfigSet(log types.Log) (*EVM2EVMMultiOffRampConfigSet, error) {
	event := new(EVM2EVMMultiOffRampConfigSet)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOffRampConfigSet0Iterator struct {
	Event *EVM2EVMMultiOffRampConfigSet0

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampConfigSet0Iterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampConfigSet0)
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
		it.Event = new(EVM2EVMMultiOffRampConfigSet0)
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

func (it *EVM2EVMMultiOffRampConfigSet0Iterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampConfigSet0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampConfigSet0 struct {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterConfigSet0(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampConfigSet0Iterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "ConfigSet0")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampConfigSet0Iterator{contract: _EVM2EVMMultiOffRamp.contract, event: "ConfigSet0", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchConfigSet0(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampConfigSet0) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "ConfigSet0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampConfigSet0)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "ConfigSet0", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseConfigSet0(log types.Log) (*EVM2EVMMultiOffRampConfigSet0, error) {
	event := new(EVM2EVMMultiOffRampConfigSet0)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "ConfigSet0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOffRampExecutionStateChangedIterator struct {
	Event *EVM2EVMMultiOffRampExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampExecutionStateChanged)
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
		it.Event = new(EVM2EVMMultiOffRampExecutionStateChanged)
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

func (it *EVM2EVMMultiOffRampExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampExecutionStateChanged struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	MessageId           [32]byte
	State               uint8
	ReturnData          []byte
	Raw                 types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*EVM2EVMMultiOffRampExecutionStateChangedIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampExecutionStateChangedIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampExecutionStateChanged)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseExecutionStateChanged(log types.Log) (*EVM2EVMMultiOffRampExecutionStateChanged, error) {
	event := new(EVM2EVMMultiOffRampExecutionStateChanged)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOffRampOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMMultiOffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMMultiOffRampOwnershipTransferRequested)
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

func (it *EVM2EVMMultiOffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampOwnershipTransferRequestedIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampOwnershipTransferRequested)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMMultiOffRampOwnershipTransferRequested, error) {
	event := new(EVM2EVMMultiOffRampOwnershipTransferRequested)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOffRampOwnershipTransferredIterator struct {
	Event *EVM2EVMMultiOffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampOwnershipTransferred)
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
		it.Event = new(EVM2EVMMultiOffRampOwnershipTransferred)
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

func (it *EVM2EVMMultiOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampOwnershipTransferredIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampOwnershipTransferred)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMMultiOffRampOwnershipTransferred, error) {
	event := new(EVM2EVMMultiOffRampOwnershipTransferred)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOffRampSkippedAlreadyExecutedMessageIterator struct {
	Event *EVM2EVMMultiOffRampSkippedAlreadyExecutedMessage

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampSkippedAlreadyExecutedMessageIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampSkippedAlreadyExecutedMessage)
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
		it.Event = new(EVM2EVMMultiOffRampSkippedAlreadyExecutedMessage)
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

func (it *EVM2EVMMultiOffRampSkippedAlreadyExecutedMessageIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampSkippedAlreadyExecutedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampSkippedAlreadyExecutedMessage struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	Raw                 types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterSkippedAlreadyExecutedMessage(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampSkippedAlreadyExecutedMessageIterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "SkippedAlreadyExecutedMessage")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampSkippedAlreadyExecutedMessageIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "SkippedAlreadyExecutedMessage", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchSkippedAlreadyExecutedMessage(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampSkippedAlreadyExecutedMessage) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "SkippedAlreadyExecutedMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampSkippedAlreadyExecutedMessage)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "SkippedAlreadyExecutedMessage", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseSkippedAlreadyExecutedMessage(log types.Log) (*EVM2EVMMultiOffRampSkippedAlreadyExecutedMessage, error) {
	event := new(EVM2EVMMultiOffRampSkippedAlreadyExecutedMessage)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "SkippedAlreadyExecutedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOffRampSkippedIncorrectNonceIterator struct {
	Event *EVM2EVMMultiOffRampSkippedIncorrectNonce

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampSkippedIncorrectNonceIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampSkippedIncorrectNonce)
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
		it.Event = new(EVM2EVMMultiOffRampSkippedIncorrectNonce)
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

func (it *EVM2EVMMultiOffRampSkippedIncorrectNonceIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampSkippedIncorrectNonceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampSkippedIncorrectNonce struct {
	SourceChainSelector uint64
	Nonce               uint64
	Sender              common.Address
	Raw                 types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterSkippedIncorrectNonce(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampSkippedIncorrectNonceIterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "SkippedIncorrectNonce")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampSkippedIncorrectNonceIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "SkippedIncorrectNonce", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchSkippedIncorrectNonce(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampSkippedIncorrectNonce) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "SkippedIncorrectNonce")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampSkippedIncorrectNonce)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "SkippedIncorrectNonce", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseSkippedIncorrectNonce(log types.Log) (*EVM2EVMMultiOffRampSkippedIncorrectNonce, error) {
	event := new(EVM2EVMMultiOffRampSkippedIncorrectNonce)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "SkippedIncorrectNonce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflightIterator struct {
	Event *EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflight

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflightIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflight)
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
		it.Event = new(EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflight)
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

func (it *EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflightIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflight struct {
	SourceChainSelector uint64
	Nonce               uint64
	Sender              common.Address
	Raw                 types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterSkippedSenderWithPreviousRampMessageInflight(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflightIterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "SkippedSenderWithPreviousRampMessageInflight")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflightIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "SkippedSenderWithPreviousRampMessageInflight", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchSkippedSenderWithPreviousRampMessageInflight(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflight) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "SkippedSenderWithPreviousRampMessageInflight")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflight)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "SkippedSenderWithPreviousRampMessageInflight", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseSkippedSenderWithPreviousRampMessageInflight(log types.Log) (*EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflight, error) {
	event := new(EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflight)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "SkippedSenderWithPreviousRampMessageInflight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOffRampSourceChainConfigSetIterator struct {
	Event *EVM2EVMMultiOffRampSourceChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampSourceChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampSourceChainConfigSet)
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
		it.Event = new(EVM2EVMMultiOffRampSourceChainConfigSet)
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

func (it *EVM2EVMMultiOffRampSourceChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampSourceChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampSourceChainConfigSet struct {
	SourceChainSelector uint64
	SourceConfig        EVM2EVMMultiOffRampSourceChainConfig
	Raw                 types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*EVM2EVMMultiOffRampSourceChainConfigSetIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampSourceChainConfigSetIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "SourceChainConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampSourceChainConfigSet)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseSourceChainConfigSet(log types.Log) (*EVM2EVMMultiOffRampSourceChainConfigSet, error) {
	event := new(EVM2EVMMultiOffRampSourceChainConfigSet)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOffRampSourceChainSelectorAddedIterator struct {
	Event *EVM2EVMMultiOffRampSourceChainSelectorAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampSourceChainSelectorAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampSourceChainSelectorAdded)
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
		it.Event = new(EVM2EVMMultiOffRampSourceChainSelectorAdded)
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

func (it *EVM2EVMMultiOffRampSourceChainSelectorAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampSourceChainSelectorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampSourceChainSelectorAdded struct {
	SourceChainSelector uint64
	Raw                 types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterSourceChainSelectorAdded(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampSourceChainSelectorAddedIterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "SourceChainSelectorAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampSourceChainSelectorAddedIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "SourceChainSelectorAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchSourceChainSelectorAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampSourceChainSelectorAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "SourceChainSelectorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampSourceChainSelectorAdded)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "SourceChainSelectorAdded", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseSourceChainSelectorAdded(log types.Log) (*EVM2EVMMultiOffRampSourceChainSelectorAdded, error) {
	event := new(EVM2EVMMultiOffRampSourceChainSelectorAdded)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "SourceChainSelectorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOffRampTransmittedIterator struct {
	Event *EVM2EVMMultiOffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampTransmitted)
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
		it.Event = new(EVM2EVMMultiOffRampTransmitted)
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

func (it *EVM2EVMMultiOffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampTransmittedIterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampTransmittedIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampTransmitted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampTransmitted)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseTransmitted(log types.Log) (*EVM2EVMMultiOffRampTransmitted, error) {
	event := new(EVM2EVMMultiOffRampTransmitted)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMMultiOffRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMMultiOffRamp.ParseConfigSet(log)
	case _EVM2EVMMultiOffRamp.abi.Events["ConfigSet0"].ID:
		return _EVM2EVMMultiOffRamp.ParseConfigSet0(log)
	case _EVM2EVMMultiOffRamp.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMMultiOffRamp.ParseExecutionStateChanged(log)
	case _EVM2EVMMultiOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMMultiOffRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMMultiOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMMultiOffRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMMultiOffRamp.abi.Events["SkippedAlreadyExecutedMessage"].ID:
		return _EVM2EVMMultiOffRamp.ParseSkippedAlreadyExecutedMessage(log)
	case _EVM2EVMMultiOffRamp.abi.Events["SkippedIncorrectNonce"].ID:
		return _EVM2EVMMultiOffRamp.ParseSkippedIncorrectNonce(log)
	case _EVM2EVMMultiOffRamp.abi.Events["SkippedSenderWithPreviousRampMessageInflight"].ID:
		return _EVM2EVMMultiOffRamp.ParseSkippedSenderWithPreviousRampMessageInflight(log)
	case _EVM2EVMMultiOffRamp.abi.Events["SourceChainConfigSet"].ID:
		return _EVM2EVMMultiOffRamp.ParseSourceChainConfigSet(log)
	case _EVM2EVMMultiOffRamp.abi.Events["SourceChainSelectorAdded"].ID:
		return _EVM2EVMMultiOffRamp.ParseSourceChainSelectorAdded(log)
	case _EVM2EVMMultiOffRamp.abi.Events["Transmitted"].ID:
		return _EVM2EVMMultiOffRamp.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMMultiOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1c86233479b9d798a31091d42b1a954067dcd6c898553d63bf993757d2a90aff")
}

func (EVM2EVMMultiOffRampConfigSet0) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMMultiOffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df2")
}

func (EVM2EVMMultiOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMMultiOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMMultiOffRampSkippedAlreadyExecutedMessage) Topic() common.Hash {
	return common.HexToHash("0x3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c")
}

func (EVM2EVMMultiOffRampSkippedIncorrectNonce) Topic() common.Hash {
	return common.HexToHash("0x852dc8e405695593e311bd83991cf39b14a328f304935eac6d3d55617f911d89")
}

func (EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflight) Topic() common.Hash {
	return common.HexToHash("0x5444a3301c7c42dd164cbf6ba4b72bf02504f86c049b06a27fc2b662e334bdbd")
}

func (EVM2EVMMultiOffRampSourceChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xdba8597411dc0624375cfff476f6173674609571f4d98d294dd3a47af0792784")
}

func (EVM2EVMMultiOffRampSourceChainSelectorAdded) Topic() common.Hash {
	return common.HexToHash("0xf4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb9")
}

func (EVM2EVMMultiOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRamp) Address() common.Address {
	return _EVM2EVMMultiOffRamp.address
}

type EVM2EVMMultiOffRampInterface interface {
	CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMMultiOffRampDynamicConfig, error)

	GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error)

	GetSenderNonce(opts *bind.CallOpts, sourceChainSelector uint64, sender common.Address) (uint64, error)

	GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (EVM2EVMMultiOffRampSourceChainConfig, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMMultiOffRampStaticConfig, error)

	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []EVM2EVMMultiOffRampSourceChainConfigArgs) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, reports []InternalExecutionReportSingleChain, gasLimitOverrides [][]*big.Int) (*types.Transaction, error)

	SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMMultiOffRampConfigSet, error)

	FilterConfigSet0(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampConfigSet0Iterator, error)

	WatchConfigSet0(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampConfigSet0) (event.Subscription, error)

	ParseConfigSet0(log types.Log) (*EVM2EVMMultiOffRampConfigSet0, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*EVM2EVMMultiOffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMMultiOffRampExecutionStateChanged, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMMultiOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMMultiOffRampOwnershipTransferred, error)

	FilterSkippedAlreadyExecutedMessage(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampSkippedAlreadyExecutedMessageIterator, error)

	WatchSkippedAlreadyExecutedMessage(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampSkippedAlreadyExecutedMessage) (event.Subscription, error)

	ParseSkippedAlreadyExecutedMessage(log types.Log) (*EVM2EVMMultiOffRampSkippedAlreadyExecutedMessage, error)

	FilterSkippedIncorrectNonce(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampSkippedIncorrectNonceIterator, error)

	WatchSkippedIncorrectNonce(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampSkippedIncorrectNonce) (event.Subscription, error)

	ParseSkippedIncorrectNonce(log types.Log) (*EVM2EVMMultiOffRampSkippedIncorrectNonce, error)

	FilterSkippedSenderWithPreviousRampMessageInflight(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflightIterator, error)

	WatchSkippedSenderWithPreviousRampMessageInflight(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflight) (event.Subscription, error)

	ParseSkippedSenderWithPreviousRampMessageInflight(log types.Log) (*EVM2EVMMultiOffRampSkippedSenderWithPreviousRampMessageInflight, error)

	FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*EVM2EVMMultiOffRampSourceChainConfigSetIterator, error)

	WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error)

	ParseSourceChainConfigSet(log types.Log) (*EVM2EVMMultiOffRampSourceChainConfigSet, error)

	FilterSourceChainSelectorAdded(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampSourceChainSelectorAddedIterator, error)

	WatchSourceChainSelectorAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampSourceChainSelectorAdded) (event.Subscription, error)

	ParseSourceChainSelectorAdded(log types.Log) (*EVM2EVMMultiOffRampSourceChainSelectorAdded, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMMultiOffRampTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
