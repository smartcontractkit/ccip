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
	MaxDataBytes                            uint32
	MaxNumberOfTokensPerMsg                 uint16
	Router                                  common.Address
	MessageValidator                        common.Address
	MaxPoolReleaseOrMintGas                 uint32
	MaxTokenTransferGas                     uint32
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
	CommitStore        common.Address
	ChainSelector      uint64
	RmnProxy           common.Address
	TokenAdminRegistry common.Address
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"InvalidManualExecutionGasLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"InvalidMessageId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"InvalidNewState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidStaticConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionGasLimitMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"errorReason\",\"type\":\"bytes\"}],\"name\":\"MessageValidationError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notPool\",\"type\":\"address\"}],\"name\":\"NotACompatiblePool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"TokenDataMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"TokenHandlingError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedTokenData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"SkippedAlreadyExecutedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedSenderWithPreviousRampMessageInflight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfig\",\"name\":\"sourceConfig\",\"type\":\"tuple\"}],\"name\":\"SourceChainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainSelectorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\"}],\"name\":\"applySourceChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"getSourceChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReportSingleChain[]\",\"name\":\"reports\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"gasLimitOverrides\",\"type\":\"uint256[][]\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005fec38038062005fec833981016040819052620000359162000677565b33806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf8162000150565b5050466080525081516001600160a01b03161580620000e9575060608201516001600160a01b0316155b1562000108576040516342bcdf7f60e11b815260040160405180910390fd5b81516001600160a01b0390811660a05260208301516001600160401b031660c0526040830151811660e052606083015116610100526200014881620001fb565b505062000881565b336001600160a01b03821603620001aa5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60005b81518110156200055e5760008282815181106200021f576200021f620007f6565b60200260200101519050600081600001519050806001600160401b0316600003620002695760405163c39a620560e01b81526001600160401b038216600482015260240162000083565b60608201516001600160a01b031662000295576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160401b038116600090815260096020526040902060018101546001600160a01b03166200045d5760a0516040516374eb454760e11b81526001600160401b03841660048201526000916001600160a01b03169063e9d68a8e90602401606060405180830381865afa15801562000313573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200033991906200080c565b905083606001516001600160a01b031681604001516001600160a01b031614158062000371575060208101516001600160401b031615155b156200039c5760405163c39a620560e01b81526001600160401b038416600482015260240162000083565b620003d38385606001517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b36200056260201b60201c565b600283015560608401516001830180546001600160a01b0319166001600160a01b039283161790556040808601518454610100600160a81b0319166101009190931602919091178355516001600160401b03841681527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb99060200160405180910390a150620004c4565b606083015160018201546001600160a01b0390811691161415806200049957506040830151815461010090046001600160a01b03908116911614155b15620004c45760405163c39a620560e01b81526001600160401b038316600482015260240162000083565b6020830151815490151560ff199091161781556040516001600160401b038316907fdba8597411dc0624375cfff476f6173674609571f4d98d294dd3a47af07927849062000547908490815460ff81161515825260081c6001600160a01b0390811660208301526001830154166040820152600290910154606082015260800190565b60405180910390a2505050806001019050620001fe565b5050565b60c05160408051602081018490526001600160401b0380871692820192909252911660608201526001600160a01b038316608082015260009060a0016040516020818303038152906040528051906020012090509392505050565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715620005f857620005f8620005bd565b60405290565b604051601f8201601f191681016001600160401b0381118282101715620006295762000629620005bd565b604052919050565b80516001600160a01b03811681146200064957600080fd5b919050565b80516001600160401b03811681146200064957600080fd5b805180151581146200064957600080fd5b60008082840360a08112156200068c57600080fd5b6080808212156200069c57600080fd5b620006a6620005d3565b9150620006b38562000631565b82526020620006c48187016200064e565b818401526040620006d86040880162000631565b60408501526060620006ed6060890162000631565b606086015260808801519496506001600160401b03808611156200071057600080fd5b858901955089601f8701126200072557600080fd5b8551818111156200073a576200073a620005bd565b6200074a858260051b01620005fe565b818152858101925060079190911b87018501908b8211156200076b57600080fd5b968501965b81881015620007e45786888d0312156200078a5760008081fd5b62000794620005d3565b6200079f896200064e565b8152620007ae878a0162000666565b87820152620007bf868a0162000631565b86820152620007d0858a0162000631565b818601528352968601969185019162000770565b80985050505050505050509250929050565b634e487b7160e01b600052603260045260246000fd5b6000606082840312156200081f57600080fd5b604051606081016001600160401b0381118282101715620008445762000844620005bd565b604052620008528362000666565b815262000862602084016200064e565b6020820152620008756040840162000631565b60408201529392505050565b60805160a05160c05160e051610100516156d26200091a600039600081816102170152818161185e01526121b00152600081816101db0152818161183701526127410152600081816101ab0152818161180e015261318701526000818161016f015281816117e001528181611b420152612a36015260008181610abc01528181610b0801528181610f610152610fad01526156d26000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c806381ff7048116100b2578063afcb95d711610081578063e9d68a8e11610066578063e9d68a8e146104f3578063f2fde38b146105e5578063f52121a5146105f857600080fd5b8063afcb95d7146104c0578063b1dc65a4146104e057600080fd5b806381ff70481461042e57806385572ffb1461045e5780638b3643341461046c5780638da5cb5b1461049857600080fd5b80635e36480c116101095780637437ff9f116100ee5780637437ff9f1461031057806379ba5097146104135780637f63b7111461041b57600080fd5b80635e36480c146102db578063666cab8d146102fb57600080fd5b806306285c691461013b578063181f5a771461026a5780631ef38174146102b3578063542625af146102c8575b600080fd5b61025460408051608081018252600080825260208201819052918101829052606081019190915260405180608001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815250905090565b6040516102619190613c5e565b60405180910390f35b6102a66040518060400160405280601d81526020017f45564d3245564d4d756c74694f666652616d7020312e362e302d64657600000081525081565b6040516102619190613cff565b6102c66102c1366004613fc4565b61060b565b005b6102c66102d636600461456c565b610ab9565b6102ee6102e9366004614697565b610ce1565b604051610261919061473a565b610303610d75565b604051610261919061479a565b6104066040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152506040805160e08101825260075463ffffffff808216835264010000000082048116602084015261ffff680100000000000000008304169383019390935273ffffffffffffffffffffffffffffffffffffffff6a0100000000000000000000909104811660608301526008549081166080830152740100000000000000000000000000000000000000008104831660a08301527801000000000000000000000000000000000000000000000000900490911660c082015290565b60405161026191906147ad565b6102c6610de4565b6102c6610429366004614828565b610ee1565b6004546002546040805163ffffffff80851682526401000000009094049093166020840152820152606001610261565b6102c661013636600461490c565b61047f61047a366004614947565b610ef5565b60405167ffffffffffffffff9091168152602001610261565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610261565b604080516001815260006020820181905291810191909152606001610261565b6102c66104ee3660046149ba565b610f0b565b610595610501366004614a9f565b6040805160808101825260008082526020820181905291810182905260608101919091525067ffffffffffffffff166000908152600960209081526040918290208251608081018452815460ff81161515825273ffffffffffffffffffffffffffffffffffffffff610100909104811693820193909352600182015490921692820192909252600290910154606082015290565b6040805182511515815260208084015173ffffffffffffffffffffffffffffffffffffffff9081169183019190915283830151169181019190915260609182015191810191909152608001610261565b6102c66105f3366004614abc565b61119c565b6102c6610606366004614ad9565b6111ad565b84518460ff16601f821115610681576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f746f6f206d616e79207472616e736d697474657273000000000000000000000060448201526064015b60405180910390fd5b806000036106eb576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401610678565b6106f3611568565b6106fc856115eb565b60065460005b8181101561078057600560006006838154811061072157610721614b3d565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168352820192909252604001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000169055600101610702565b50875160005b818110156109765760008a82815181106107a2576107a2614b3d565b60200260200101519050600060028111156107bf576107bf6146d0565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260056020526040902054610100900460ff1660028111156107fe576107fe6146d0565b14610865576040517f89a6198900000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401610678565b73ffffffffffffffffffffffffffffffffffffffff81166108b2576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff83168152602081016002905273ffffffffffffffffffffffffffffffffffffffff821660009081526005602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001617610100836002811115610962576109626146d0565b021790555090505050806001019050610786565b50885161098a9060069060208c0190613bc8565b506003805460ff838116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909216908b161717905560048054610a109146913091906000906109e29063ffffffff16614b9b565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168d8d8d8d8d8d6118be565b600260000181905550600060048054906101000a900463ffffffff169050436004806101000a81548163ffffffff021916908363ffffffff1602179055507f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0581600260000154600460009054906101000a900463ffffffff168e8e8e8e8e8e604051610aa499989796959493929190614bbe565b60405180910390a15050505050505050505050565b467f000000000000000000000000000000000000000000000000000000000000000014610b44576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015267ffffffffffffffff46166024820152604401610678565b815181518114610b80576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b81811015610cd1576000848281518110610b9f57610b9f614b3d565b60200260200101519050600081602001515190506000858481518110610bc757610bc7614b3d565b6020026020010151905080518214610c0b576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b82811015610cc2576000828281518110610c2a57610c2a614b3d565b6020026020010151905080600014158015610c65575084602001518281518110610c5657610c56614b3d565b60200260200101516080015181105b15610cb95784516040517fc8e9605100000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024810183905260448101829052606401610678565b50600101610c0e565b50505050806001019050610b83565b50610cdc838361194b565b505050565b6000610cef60016004614c54565b6002610cfc608085614c96565b67ffffffffffffffff16610d109190614cbd565b67ffffffffffffffff85166000908152600b6020526040812090610d35608087614cd4565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002054901c166003811115610d6c57610d6c6146d0565b90505b92915050565b60606006805480602002602001604051908101604052809291908181526020018280548015610dda57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610daf575b5050505050905090565b60015473ffffffffffffffffffffffffffffffffffffffff163314610e65576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610678565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610ee9611568565b610ef2816119fb565b50565b600080610f028484611eac565b50949350505050565b610f158787611fdc565b600254883590808214610f5e576040517f93df584c0000000000000000000000000000000000000000000000000000000081526004810182905260248101839052604401610678565b467f000000000000000000000000000000000000000000000000000000000000000014610fdf576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006004820152466024820152604401610678565b6040805183815260208c81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a13360009081526005602090815260408083208151808301909252805460ff80821684529293919291840191610100909104166002811115611067576110676146d0565b6002811115611078576110786146d0565b9052509050600281602001516002811115611095576110956146d0565b1480156110dc57506006816000015160ff16815481106110b7576110b7614b3d565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b611112576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506000611120856020614cbd565b61112b886020614cbd565b6111378b610144614cfb565b6111419190614cfb565b61114b9190614cfb565b905036811461118f576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401610678565b5050505050505050505050565b6111a4611568565b610ef281612023565b3330146111e6576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160008082526020820190925281611223565b60408051808201909152600080825260208201528152602001906001900390816111fc5790505b5061014084015151909150156112be576101408301516040805160608101909152602085015173ffffffffffffffffffffffffffffffffffffffff1660808201526112bb91908060a0810160408051601f19818403018152918152908252875167ffffffffffffffff1660208301528781015173ffffffffffffffffffffffffffffffffffffffff1691015261016086015185612118565b90505b60006112ca848361263b565b60085490915073ffffffffffffffffffffffffffffffffffffffff1680156113d1576040517fa219f6e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82169063a219f6e59061133e908590600401614dcf565b600060405180830381600087803b15801561135857600080fd5b505af1925050508015611369575060015b6113d1573d808015611397576040519150601f19603f3d011682016040523d82523d6000602084013e61139c565b606091505b50806040517f09c253250000000000000000000000000000000000000000000000000000000081526004016106789190613cff565b610120850151511580156113e757506080850151155b8061140b5750604085015173ffffffffffffffffffffffffffffffffffffffff163b155b80611458575060408501516114569073ffffffffffffffffffffffffffffffffffffffff167f85572ffb000000000000000000000000000000000000000000000000000000006126eb565b155b15611464575050505050565b600754608086015160408088015190517f3cf9798300000000000000000000000000000000000000000000000000000000815260009384936a010000000000000000000090910473ffffffffffffffffffffffffffffffffffffffff1692633cf97983926114da92899261138892600401614de2565b6000604051808303816000875af11580156114f9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526115219190810190614e70565b50915091508161155f57806040517f0a8d6e8c0000000000000000000000000000000000000000000000000000000081526004016106789190613cff565b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146115e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610678565b565b6000818060200190518101906116019190614ee9565b606081015190915073ffffffffffffffffffffffffffffffffffffffff16611655576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516007805460208085015160408087015160608089015173ffffffffffffffffffffffffffffffffffffffff9081166a0100000000000000000000027fffff0000000000000000000000000000000000000000ffffffffffffffffffff61ffff9094166801000000000000000002939093167fffff00000000000000000000000000000000000000000000ffffffffffffffff63ffffffff968716640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009099169a87169a909a17979097179890981695909517179094556080808701516008805460a08a015160c08b015186167801000000000000000000000000000000000000000000000000027fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff9190961674010000000000000000000000000000000000000000027fffffffffffffffff000000000000000000000000000000000000000000000000909216938a1693909317179190911692909217909155835190810184527f00000000000000000000000000000000000000000000000000000000000000008516815267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016918101919091527f00000000000000000000000000000000000000000000000000000000000000008416818401527f000000000000000000000000000000000000000000000000000000000000000090931690830152517f8ee5eb870fc8e459cd27552173357785e2e9bbd45b41ee742a05610b0a728c7f916118b2918490614f84565b60405180910390a15050565b6000808a8a8a8a8a8a8a8a8a6040516020016118e299989796959493929190615050565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b8151600003611985576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160408051600080825260208201909252911591905b84518110156119f4576119ec8582815181106119ba576119ba614b3d565b6020026020010151846119e6578583815181106119d9576119d9614b3d565b6020026020010151612707565b83612707565b60010161199c565b5050505050565b60005b8151811015611ea8576000828281518110611a1b57611a1b614b3d565b602002602001015190506000816000015190508067ffffffffffffffff16600003611a7e576040517fc39a620500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610678565b606082015173ffffffffffffffffffffffffffffffffffffffff16611acf576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff81166000908152600960205260409020600181015473ffffffffffffffffffffffffffffffffffffffff16611d4c576040517fe9d68a8e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff831660048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063e9d68a8e90602401606060405180830381865afa158015611b9e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bc291906150e5565b9050836060015173ffffffffffffffffffffffffffffffffffffffff16816040015173ffffffffffffffffffffffffffffffffffffffff16141580611c145750602081015167ffffffffffffffff1615155b15611c57576040517fc39a620500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610678565b611c868385606001517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3613181565b600283015560608401516001830180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92831617905560408086015184547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010091909316029190911783555167ffffffffffffffff841681527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb99060200160405180910390a150611de4565b6060830151600182015473ffffffffffffffffffffffffffffffffffffffff9081169116141580611da1575060408301518154610100900473ffffffffffffffffffffffffffffffffffffffff908116911614155b15611de4576040517fc39a620500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff83166004820152602401610678565b602083015181549015157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090911617815560405167ffffffffffffffff8316907fdba8597411dc0624375cfff476f6173674609571f4d98d294dd3a47af079278490611e92908490815460ff81161515825260081c73ffffffffffffffffffffffffffffffffffffffff90811660208301526001830154166040820152600290910154606082015260800190565b60405180910390a25050508060010190506119fe565b5050565b67ffffffffffffffff8083166000908152600a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff861684529091528120549091829116808203611fce5767ffffffffffffffff8516600090815260096020526040902054610100900473ffffffffffffffffffffffffffffffffffffffff168015611fcc576040517f856c824700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015282169063856c824790602401602060405180830381865afa158015611f9b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fbf9190615139565b6001935093505050611fd5565b505b9150600090505b9250929050565b6000611fea82840184615156565b60408051600080825260208201909252919250610cdc91839161201d565b60608152602001906001900390816120085790505b5061194b565b3373ffffffffffffffffffffffffffffffffffffffff8216036120a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610678565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8360005b8551811015610f0257600084828151811061213957612139614b3d565b6020026020010151806020019051810190612154919061518b565b905060006121658260200151613211565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301529192506000917f0000000000000000000000000000000000000000000000000000000000000000169063bbe4f6db90602401602060405180830381865afa1580156121f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061221b9190615240565b905073ffffffffffffffffffffffffffffffffffffffff8116158061227d575061227b73ffffffffffffffffffffffffffffffffffffffff82167faff2afbf000000000000000000000000000000000000000000000000000000006126eb565b155b156122cc576040517fae9b4ce900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610678565b600080612414634059f55b60e01b6040518060e001604052808d6000015181526020018d6020015167ffffffffffffffff1681526020018d6040015173ffffffffffffffffffffffffffffffffffffffff1681526020018e8a8151811061233557612335614b3d565b602002602001015160200151815260200188600001518152602001886040015181526020018b8a8151811061236c5761236c614b3d565b6020026020010151815250604051602401612387919061525d565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152600854859063ffffffff7401000000000000000000000000000000000000000090910416611388608461326c565b50915091508161245257806040517fe1cd55090000000000000000000000000000000000000000000000000000000081526004016106789190613cff565b805160201461249a5780516040517f78ef8024000000000000000000000000000000000000000000000000000000008152602060048201526024810191909152604401610678565b6000818060200190518101906124b09190615319565b60408c810151815173ffffffffffffffffffffffffffffffffffffffff909116602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905260085491925061256f9187907801000000000000000000000000000000000000000000000000900463ffffffff16611388608461326c565b509093509150826125ae57816040517fe1cd55090000000000000000000000000000000000000000000000000000000081526004016106789190613cff565b848888815181106125c1576125c1614b3d565b60200260200101516000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508088888151811061261257612612614b3d565b6020026020010151602001818152505050505050505080600101905061211c565b949350505050565b6040805160a08101825260008082526020820152606091810182905281810182905260808101919091526040518060a001604052808461018001518152602001846000015167ffffffffffffffff16815260200184602001516040516020016126c0919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b6040516020818303038152906040528152602001846101200151815260200183815250905092915050565b60006126f683613392565b8015610d6c5750610d6c83836133f6565b81516040517f58babe3300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906358babe3390602401602060405180830381865afa15801561279d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127c19190615332565b15612804576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610678565b6020830151516000819003612844576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8360400151518114612882576040517f57e0e08300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff82166000908152600960205260409020805460ff166128e2576040517fed053c5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610678565b60008267ffffffffffffffff8111156128fd576128fd613d12565b604051908082528060200260200182016040528015612926578160200160208202803683370190505b50905060005b838110156129eb5760008760200151828151811061294c5761294c614b3d565b602002602001015190506129648185600201546134c5565b83838151811061297657612976614b3d565b60200260200101818152505080610180015183838151811061299a5761299a614b3d565b6020026020010151146129e2578061018001516040517f345039be00000000000000000000000000000000000000000000000000000000815260040161067891815260200190565b5060010161292c565b50606086015160808701516040517ffe41448f00000000000000000000000000000000000000000000000000000000815260009273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169263fe41448f92612a6d928a928892600401615380565b602060405180830381865afa158015612a8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612aae9190615319565b905080600003612af6576040517f7dd17a7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610678565b8551151560005b8581101561317657600089602001518281518110612b1d57612b1d614b3d565b602002602001015190506000612b37898360600151610ce1565b90506002816003811115612b4d57612b4d6146d0565b03612ba35760608201516040805167ffffffffffffffff808d16825290921660208301527f3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c910160405180910390a1505061316e565b6000816003811115612bb757612bb76146d0565b1480612bd457506003816003811115612bd257612bd26146d0565b145b612c245760608201516040517f25507e7f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808c1660048301529091166024820152604401610678565b8315612ced5760075460009063ffffffff16612c408742614c54565b1190508080612c6057506003826003811115612c5e57612c5e6146d0565b145b612ca2576040517fa9cfc86200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b166004820152602401610678565b8a8481518110612cb457612cb4614b3d565b6020026020010151600014612ce7578a8481518110612cd557612cd5614b3d565b60200260200101518360800181815250505b50612d52565b6000816003811115612d0157612d016146d0565b14612d525760608201516040517f3ef2a99c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808c1660048301529091166024820152604401610678565b600080612d638b8560200151611eac565b915091508015612e835760c084015167ffffffffffffffff16612d878360016153c7565b67ffffffffffffffff1614612e175760c084015160208501516040517f5444a3301c7c42dd164cbf6ba4b72bf02504f86c049b06a27fc2b662e334bdbd92612e06928f9267ffffffffffffffff938416815291909216602082015273ffffffffffffffffffffffffffffffffffffffff91909116604082015260600190565b60405180910390a15050505061316e565b67ffffffffffffffff8b81166000908152600a602090815260408083208883015173ffffffffffffffffffffffffffffffffffffffff168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000169184169190911790555b6000836003811115612e9757612e976146d0565b03612f355760c084015167ffffffffffffffff16612eb68360016153c7565b67ffffffffffffffff1614612f355760c084015160208501516040517f852dc8e405695593e311bd83991cf39b14a328f304935eac6d3d55617f911d8992612e06928f9267ffffffffffffffff938416815291909216602082015273ffffffffffffffffffffffffffffffffffffffff91909116604082015260600190565b60008d604001518681518110612f4d57612f4d614b3d565b60200260200101519050612f7b8561018001518d87606001518861014001515189610120015151865161362d565b612f8b8c8660600151600161374d565b600080612f98878461382b565b91509150612fab8e88606001518461374d565b888015612fc957506003826003811115612fc757612fc76146d0565b145b1561300957866101800151816040517f2b11b8d90000000000000000000000000000000000000000000000000000000081526004016106789291906153ef565b600382600381111561301d5761301d6146d0565b1415801561303d5750600282600381111561303a5761303a6146d0565b14155b1561307e578d8760600151836040517f926c5a3e00000000000000000000000000000000000000000000000000000000815260040161067893929190615408565b6000866003811115613092576130926146d0565b0361310d5767ffffffffffffffff808f166000908152600a602090815260408083208b83015173ffffffffffffffffffffffffffffffffffffffff1684529091528120805490921691906130e58361542e565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b866101800151876060015167ffffffffffffffff168f67ffffffffffffffff167f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df2858560405161315e92919061544b565b60405180910390a4505050505050505b600101612afd565b505050505050505050565b600081847f0000000000000000000000000000000000000000000000000000000000000000856040516020016131f1949392919093845267ffffffffffffffff92831660208501529116604083015273ffffffffffffffffffffffffffffffffffffffff16606082015260800190565b6040516020818303038152906040528051906020012090505b9392505050565b6000815160201461325057816040517f8d666f600000000000000000000000000000000000000000000000000000000081526004016106789190613cff565b610d6f828060200190518101906132679190615319565b613b4f565b6000606060008361ffff1667ffffffffffffffff81111561328f5761328f613d12565b6040519080825280601f01601f1916602001820160405280156132b9576020820181803683370190505b509150863b6132ec577f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b5a8581101561331f577fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b8590036040810481038710613358577f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b505a6000808a5160208c0160008c8cf193505a900390503d8481111561337b5750835b808352806000602085013e50955095509592505050565b60006133be827f01ffc9a7000000000000000000000000000000000000000000000000000000006133f6565b8015610d6f57506133ef827fffffffff000000000000000000000000000000000000000000000000000000006133f6565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d915060005190508280156134ae575060208210155b80156134ba5750600081115b979650505050505050565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b610100015160405160200161356898979695949392919073ffffffffffffffffffffffffffffffffffffffff9889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b60405160208183030381529060405280519060200120856101200151805190602001208661014001516040516020016135a1919061546b565b604051602081830303815290604052805190602001208761016001516040516020016135cd91906154d8565b60408051601f198184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b60075468010000000000000000900461ffff1683111561368d576040517fa1e5205a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808716600483015285166024820152604401610678565b8083146136da576040517f1cfe6d8b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808716600483015285166024820152604401610678565b600754640100000000900463ffffffff16821115613745576007546040517f1fd8fd040000000000000000000000000000000000000000000000000000000081526004810188905264010000000090910463ffffffff16602482015260448101839052606401610678565b505050505050565b6000600261375c608085614c96565b67ffffffffffffffff166137709190614cbd565b67ffffffffffffffff85166000908152600b602052604081209192509081613799608087614cd4565b67ffffffffffffffff1681526020810191909152604001600020549050816137c360016004614c54565b901b1916818360038111156137da576137da6146d0565b67ffffffffffffffff87166000908152600b602052604081209190921b92909217918291613809608088614cd4565b67ffffffffffffffff1681526020810191909152604001600020555050505050565b6040517ff52121a5000000000000000000000000000000000000000000000000000000008152600090606090309063f52121a59061386f90879087906004016154eb565b600060405180830381600087803b15801561388957600080fd5b505af192505050801561389a575060015b613b34573d8080156138c8576040519150601f19603f3d011682016040523d82523d6000602084013e6138cd565b606091505b5060006138d982615675565b90507f0a8d6e8c000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216148061396c57507fe1cd5509000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b806139b857507f8d666f60000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b80613a0457507f78ef8024000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b80613a5057507f0c3b563c000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b80613a9c57507fae9b4ce9000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b80613ae857507f09c25325000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b15613af95750600392509050611fd5565b856101800151826040517f2b11b8d90000000000000000000000000000000000000000000000000000000081526004016106789291906153ef565b50506040805160208101909152600081526002909250929050565b600073ffffffffffffffffffffffffffffffffffffffff821180613b735750600a82105b15613bc45760408051602081018490520160408051601f19818403018152908290527f8d666f6000000000000000000000000000000000000000000000000000000000825261067891600401613cff565b5090565b828054828255906000526020600020908101928215613c42579160200282015b82811115613c4257825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190613be8565b50613bc49291505b80821115613bc45760008155600101613c4a565b60808101610d6f8284805173ffffffffffffffffffffffffffffffffffffffff908116835260208083015167ffffffffffffffff169084015260408083015182169084015260609182015116910152565b60005b83811015613cca578181015183820152602001613cb2565b50506000910152565b60008151808452613ceb816020860160208601613caf565b601f01601f19169290920160200192915050565b602081526000610d6c6020830184613cd3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715613d6457613d64613d12565b60405290565b6040516101a0810167ffffffffffffffff81118282101715613d6457613d64613d12565b60405160a0810167ffffffffffffffff81118282101715613d6457613d64613d12565b6040516080810167ffffffffffffffff81118282101715613d6457613d64613d12565b60405160e0810167ffffffffffffffff81118282101715613d6457613d64613d12565b6040516060810167ffffffffffffffff81118282101715613d6457613d64613d12565b604051601f8201601f1916810167ffffffffffffffff81118282101715613e4357613e43613d12565b604052919050565b600067ffffffffffffffff821115613e6557613e65613d12565b5060051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81168114610ef257600080fd5b8035613e9c81613e6f565b919050565b600082601f830112613eb257600080fd5b81356020613ec7613ec283613e4b565b613e1a565b8083825260208201915060208460051b870101935086841115613ee957600080fd5b602086015b84811015613f0e578035613f0181613e6f565b8352918301918301613eee565b509695505050505050565b803560ff81168114613e9c57600080fd5b600067ffffffffffffffff821115613f4457613f44613d12565b50601f01601f191660200190565b600082601f830112613f6357600080fd5b8135613f71613ec282613f2a565b818152846020838601011115613f8657600080fd5b816020850160208301376000918101602001919091529392505050565b67ffffffffffffffff81168114610ef257600080fd5b8035613e9c81613fa3565b60008060008060008060c08789031215613fdd57600080fd5b863567ffffffffffffffff80821115613ff557600080fd5b6140018a838b01613ea1565b9750602089013591508082111561401757600080fd5b6140238a838b01613ea1565b965061403160408a01613f19565b9550606089013591508082111561404757600080fd5b6140538a838b01613f52565b945061406160808a01613fb9565b935060a089013591508082111561407757600080fd5b5061408489828a01613f52565b9150509295509295509295565b8015158114610ef257600080fd5b8035613e9c81614091565b600082601f8301126140bb57600080fd5b813560206140cb613ec283613e4b565b82815260069290921b840181019181810190868411156140ea57600080fd5b8286015b84811015613f0e57604081890312156141075760008081fd5b61410f613d41565b813561411a81613e6f565b815281850135858201528352918301916040016140ee565b600082601f83011261414357600080fd5b81356020614153613ec283613e4b565b82815260059290921b8401810191818101908684111561417257600080fd5b8286015b84811015613f0e57803567ffffffffffffffff8111156141965760008081fd5b6141a48986838b0101613f52565b845250918301918301614176565b60006101a082840312156141c557600080fd5b6141cd613d6a565b90506141d882613fb9565b81526141e660208301613e91565b60208201526141f760408301613e91565b604082015261420860608301613fb9565b60608201526080820135608082015261422360a0830161409f565b60a082015261423460c08301613fb9565b60c082015261424560e08301613e91565b60e082015261010082810135908201526101208083013567ffffffffffffffff8082111561427257600080fd5b61427e86838701613f52565b8385015261014092508285013591508082111561429a57600080fd5b6142a6868387016140aa565b838501526101609250828501359150808211156142c257600080fd5b506142cf85828601614132565b82840152505061018080830135818301525092915050565b600082601f8301126142f857600080fd5b81356020614308613ec283613e4b565b82815260059290921b8401810191818101908684111561432757600080fd5b8286015b84811015613f0e57803567ffffffffffffffff81111561434b5760008081fd5b6143598986838b01016141b2565b84525091830191830161432b565b600082601f83011261437857600080fd5b81356020614388613ec283613e4b565b82815260059290921b840181019181810190868411156143a757600080fd5b8286015b84811015613f0e57803567ffffffffffffffff8111156143cb5760008081fd5b6143d98986838b0101614132565b8452509183019183016143ab565b600082601f8301126143f857600080fd5b81356020614408613ec283613e4b565b8083825260208201915060208460051b87010193508684111561442a57600080fd5b602086015b84811015613f0e578035835291830191830161442f565b600082601f83011261445757600080fd5b81356020614467613ec283613e4b565b82815260059290921b8401810191818101908684111561448657600080fd5b8286015b84811015613f0e57803567ffffffffffffffff808211156144ab5760008081fd5b818901915060a080601f19848d030112156144c65760008081fd5b6144ce613d8e565b6144d9888501613fb9565b8152604080850135848111156144ef5760008081fd5b6144fd8e8b838901016142e7565b8a84015250606080860135858111156145165760008081fd5b6145248f8c838a0101614367565b838501525060809150818601358581111561453f5760008081fd5b61454d8f8c838a01016143e7565b918401919091525091909301359083015250835291830191830161448a565b600080604080848603121561458057600080fd5b833567ffffffffffffffff8082111561459857600080fd5b6145a487838801614446565b94506020915081860135818111156145bb57600080fd5b8601601f810188136145cc57600080fd5b80356145da613ec282613e4b565b81815260059190911b8201840190848101908a8311156145f957600080fd5b8584015b83811015614685578035868111156146155760008081fd5b8501603f81018d136146275760008081fd5b87810135614637613ec282613e4b565b81815260059190911b82018a0190898101908f8311156146575760008081fd5b928b01925b828410156146755783358252928a0192908a019061465c565b86525050509186019186016145fd565b50809750505050505050509250929050565b600080604083850312156146aa57600080fd5b82356146b581613fa3565b915060208301356146c581613fa3565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110614736577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b60208101610d6f82846146ff565b60008151808452602080850194506020840160005b8381101561478f57815173ffffffffffffffffffffffffffffffffffffffff168752958201959082019060010161475d565b509495945050505050565b602081526000610d6c6020830184614748565b60e08101610d6f828463ffffffff80825116835280602083015116602084015261ffff6040830151166040840152606082015173ffffffffffffffffffffffffffffffffffffffff808216606086015280608085015116608086015250508060a08301511660a08401528060c08301511660c0840152505050565b6000602080838503121561483b57600080fd5b823567ffffffffffffffff81111561485257600080fd5b8301601f8101851361486357600080fd5b8035614871613ec282613e4b565b81815260079190911b8201830190838101908783111561489057600080fd5b928401925b828410156134ba57608084890312156148ae5760008081fd5b6148b6613db1565b84356148c181613fa3565b8152848601356148d081614091565b818701526040858101356148e381613e6f565b908201526060858101356148f681613e6f565b9082015282526080939093019290840190614895565b60006020828403121561491e57600080fd5b813567ffffffffffffffff81111561493557600080fd5b820160a0818503121561320a57600080fd5b6000806040838503121561495a57600080fd5b823561496581613fa3565b915060208301356146c581613e6f565b60008083601f84011261498757600080fd5b50813567ffffffffffffffff81111561499f57600080fd5b6020830191508360208260051b8501011115611fd557600080fd5b60008060008060008060008060e0898b0312156149d657600080fd5b606089018a8111156149e757600080fd5b8998503567ffffffffffffffff80821115614a0157600080fd5b818b0191508b601f830112614a1557600080fd5b813581811115614a2457600080fd5b8c6020828501011115614a3657600080fd5b6020830199508098505060808b0135915080821115614a5457600080fd5b614a608c838d01614975565b909750955060a08b0135915080821115614a7957600080fd5b50614a868b828c01614975565b999c989b50969995989497949560c00135949350505050565b600060208284031215614ab157600080fd5b813561320a81613fa3565b600060208284031215614ace57600080fd5b813561320a81613e6f565b60008060408385031215614aec57600080fd5b823567ffffffffffffffff80821115614b0457600080fd5b614b10868387016141b2565b93506020850135915080821115614b2657600080fd5b50614b3385828601614132565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600063ffffffff808316818103614bb457614bb4614b6c565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b16604085015250806060840152614bee8184018a614748565b90508281036080840152614c028189614748565b905060ff871660a084015282810360c0840152614c1f8187613cd3565b905067ffffffffffffffff851660e0840152828103610100840152614c448185613cd3565b9c9b505050505050505050505050565b81810381811115610d6f57610d6f614b6c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff80841680614cb157614cb1614c67565b92169190910692915050565b8082028115828204841417610d6f57610d6f614b6c565b600067ffffffffffffffff80841680614cef57614cef614c67565b92169190910492915050565b80820180821115610d6f57610d6f614b6c565b60008151808452602080850194506020840160005b8381101561478f578151805173ffffffffffffffffffffffffffffffffffffffff1688528301518388015260409096019590820190600101614d23565b8051825267ffffffffffffffff60208201511660208301526000604082015160a06040850152614d9360a0850182613cd3565b905060608301518482036060860152614dac8282613cd3565b91505060808301518482036080860152614dc68282614d0e565b95945050505050565b602081526000610d6c6020830184614d60565b608081526000614df56080830187614d60565b61ffff95909516602083015250604081019290925273ffffffffffffffffffffffffffffffffffffffff16606090910152919050565b600082601f830112614e3c57600080fd5b8151614e4a613ec282613f2a565b818152846020838601011115614e5f57600080fd5b612633826020830160208701613caf565b600080600060608486031215614e8557600080fd5b8351614e9081614091565b602085015190935067ffffffffffffffff811115614ead57600080fd5b614eb986828701614e2b565b925050604084015190509250925092565b805163ffffffff81168114613e9c57600080fd5b8051613e9c81613e6f565b600060e08284031215614efb57600080fd5b614f03613dd4565b614f0c83614eca565b8152614f1a60208401614eca565b6020820152604083015161ffff81168114614f3457600080fd5b6040820152614f4560608401614ede565b6060820152614f5660808401614ede565b6080820152614f6760a08401614eca565b60a0820152614f7860c08401614eca565b60c08201529392505050565b6101608101614fd68285805173ffffffffffffffffffffffffffffffffffffffff908116835260208083015167ffffffffffffffff169084015260408083015182169084015260609182015116910152565b61320a608083018463ffffffff80825116835280602083015116602084015261ffff6040830151166040840152606082015173ffffffffffffffffffffffffffffffffffffffff808216606086015280608085015116608086015250508060a08301511660a08401528060c08301511660c0840152505050565b60006101208b835273ffffffffffffffffffffffffffffffffffffffff8b16602084015267ffffffffffffffff808b1660408501528160608501526150978285018b614748565b915083820360808501526150ab828a614748565b915060ff881660a085015283820360c08501526150c88288613cd3565b90861660e08501528381036101008501529050614c448185613cd3565b6000606082840312156150f757600080fd5b6150ff613df7565b825161510a81614091565b8152602083015161511a81613fa3565b6020820152604083015161512d81613e6f565b60408201529392505050565b60006020828403121561514b57600080fd5b815161320a81613fa3565b60006020828403121561516857600080fd5b813567ffffffffffffffff81111561517f57600080fd5b61263384828501614446565b60006020828403121561519d57600080fd5b815167ffffffffffffffff808211156151b557600080fd5b90830190606082860312156151c957600080fd5b6151d1613df7565b8251828111156151e057600080fd5b6151ec87828601614e2b565b82525060208301518281111561520157600080fd5b61520d87828601614e2b565b60208301525060408301518281111561522557600080fd5b61523187828601614e2b565b60408301525095945050505050565b60006020828403121561525257600080fd5b815161320a81613e6f565b602081526000825160e0602084015261527a610100840182613cd3565b905067ffffffffffffffff602085015116604084015260408401516152b7606085018273ffffffffffffffffffffffffffffffffffffffff169052565b50606084015160808401526080840151601f19808584030160a08601526152de8383613cd3565b925060a08601519150808584030160c08601526152fb8383613cd3565b925060c08601519150808584030160e086015250614dc68282613cd3565b60006020828403121561532b57600080fd5b5051919050565b60006020828403121561534457600080fd5b815161320a81614091565b60008151808452602080850194506020840160005b8381101561478f57815187529582019590820190600101615364565b67ffffffffffffffff851681526080602082015260006153a3608083018661534f565b82810360408401526153b5818661534f565b91505082606083015295945050505050565b67ffffffffffffffff8181168382160190808211156153e8576153e8614b6c565b5092915050565b8281526040602082015260006126336040830184613cd3565b67ffffffffffffffff8481168252831660208201526060810161263360408301846146ff565b600067ffffffffffffffff808316818103614bb457614bb4614b6c565b61545581846146ff565b6040602082015260006126336040830184613cd3565b602081526000610d6c6020830184614d0e565b60008282518085526020808601955060208260051b8401016020860160005b848110156154cb57601f198684030189526154b9838351613cd3565b9884019892509083019060010161549d565b5090979650505050505050565b602081526000610d6c602083018461547e565b6040815261550660408201845167ffffffffffffffff169052565b6000602084015161552f606084018273ffffffffffffffffffffffffffffffffffffffff169052565b50604084015173ffffffffffffffffffffffffffffffffffffffff8116608084015250606084015167ffffffffffffffff811660a084015250608084015160c083015260a084015180151560e08401525060c084015161010061559d8185018367ffffffffffffffff169052565b60e086015191506101206155c88186018473ffffffffffffffffffffffffffffffffffffffff169052565b81870151925061014091508282860152808701519250506101a061016081818701526155f86101e0870185613cd3565b93508288015192507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06101808188870301818901526156378686614d0e565b9550828a01519450818887030184890152615652868661547e565b9550808a01516101c089015250505050508281036020840152614dc6818561547e565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156156bd5780818460040360031b1b83161693505b50505091905056fea164736f6c6343000818000a",
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
	return common.HexToHash("0x8ee5eb870fc8e459cd27552173357785e2e9bbd45b41ee742a05610b0a728c7f")
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
