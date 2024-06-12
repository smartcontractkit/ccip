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

type EVM2EVMMultiOffRampCommitReport struct {
	PriceUpdates InternalPriceUpdates
	MerkleRoots  []EVM2EVMMultiOffRampMerkleRoot
}

type EVM2EVMMultiOffRampDynamicConfig struct {
	Router                                  common.Address
	PermissionLessExecutionThresholdSeconds uint32
	MaxTokenTransferGas                     uint32
	MaxPoolReleaseOrMintGas                 uint32
	MaxNumberOfTokensPerMsg                 uint16
	MaxDataBytes                            uint32
	MessageValidator                        common.Address
	PriceRegistry                           common.Address
}

type EVM2EVMMultiOffRampInterval struct {
	Min uint64
	Max uint64
}

type EVM2EVMMultiOffRampMerkleRoot struct {
	SourceChainSelector uint64
	Interval            EVM2EVMMultiOffRampInterval
	MerkleRoot          [32]byte
}

type EVM2EVMMultiOffRampSourceChainConfig struct {
	IsEnabled    bool
	MinSeqNr     uint64
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
	ChainSelector uint64
	RmnProxy      common.Address
}

type EVM2EVMMultiOffRampUnblessedRoot struct {
	SourceChainSelector uint64
	MerkleRoot          [32]byte
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

type InternalGasPriceUpdate struct {
	DestChainSelector uint64
	UsdPerUnitGas     *big.Int
}

type InternalPriceUpdates struct {
	TokenPriceUpdates []InternalTokenPriceUpdate
	GasPriceUpdates   []InternalGasPriceUpdate
}

type InternalTokenPriceUpdate struct {
	SourceToken common.Address
	UsdPerToken *big.Int
}

type MultiOCR3BaseConfigInfo struct {
	ConfigDigest                   [32]byte
	F                              uint8
	N                              uint8
	UniqueReports                  bool
	IsSignatureVerificationEnabled bool
}

type MultiOCR3BaseOCRConfig struct {
	ConfigInfo   MultiOCR3BaseConfigInfo
	Signers      []common.Address
	Transmitters []common.Address
}

type MultiOCR3BaseOCRConfigArgs struct {
	ConfigDigest                   [32]byte
	OcrPluginType                  uint8
	F                              uint8
	UniqueReports                  bool
	IsSignatureVerificationEnabled bool
	Signers                        []common.Address
	Transmitters                   []common.Address
}

var EVM2EVMMultiOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumMultiOCR3Base.InvalidConfigErrorType\",\"name\":\"errorType\",\"type\":\"uint8\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.Interval\",\"name\":\"interval\",\"type\":\"tuple\"}],\"name\":\"InvalidInterval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"InvalidManualExecutionGasLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"InvalidMessageId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"InvalidNewState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRoot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidStaticConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeavesCannotBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionGasLimitMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"errorReason\",\"type\":\"bytes\"}],\"name\":\"MessageValidationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notPool\",\"type\":\"address\"}],\"name\":\"NotACompatiblePool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausedError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"RootAlreadyCommitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleCommitReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"}],\"name\":\"StaticConfigCannotBeChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"TokenDataMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"TokenHandlingError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedTokenData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroChainSelectorNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint224\",\"name\":\"usdPerToken\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint224\",\"name\":\"usdPerUnitGas\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.GasPriceUpdate[]\",\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.PriceUpdates\",\"name\":\"priceUpdates\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.Interval\",\"name\":\"interval\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.MerkleRoot[]\",\"name\":\"merkleRoots\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.CommitReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"CommitReportAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"RootRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"SkippedAlreadyExecutedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedSenderWithPreviousRampMessageInflight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfig\",\"name\":\"sourceConfig\",\"type\":\"tuple\"}],\"name\":\"SourceChainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainSelectorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\"}],\"name\":\"applySourceChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPriceEpochAndRound\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"getSourceChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"isBlessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"isUnpausedAndNotCursed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"}],\"name\":\"latestConfigDetails\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"n\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"uniqueReports\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\"}],\"internalType\":\"structMultiOCR3Base.ConfigInfo\",\"name\":\"configInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"}],\"internalType\":\"structMultiOCR3Base.OCRConfig\",\"name\":\"ocrConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReportSingleChain[]\",\"name\":\"reports\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"gasLimitOverrides\",\"type\":\"uint256[][]\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.UnblessedRoot[]\",\"name\":\"rootToReset\",\"type\":\"tuple[]\"}],\"name\":\"resetUnblessedRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"latestPriceEpochAndRound\",\"type\":\"uint40\"}],\"name\":\"setLatestPriceEpochAndRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"uniqueReports\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"}],\"internalType\":\"structMultiOCR3Base.OCRConfigArgs[]\",\"name\":\"ocrConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"setOCR3Configs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e0604052600b805460ff60281b191690553480156200001e57600080fd5b5060405162007862380380620078628339810160408190526200004191620005e3565b3380600081620000985760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000cb57620000cb816200015c565b5050466080525060208201516001600160a01b0316620000fe576040516342bcdf7f60e11b815260040160405180910390fd5b81516001600160401b03166000036200012a5760405163c656089560e01b815260040160405180910390fd5b81516001600160401b031660a05260208201516001600160a01b031660c052620001548162000207565b505062000759565b336001600160a01b03821603620001b65760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200008f565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60005b8151811015620004b75760008282815181106200022b576200022b62000743565b60200260200101519050600081600001519050806001600160401b0316600003620002695760405163c656089560e01b815260040160405180910390fd5b60608201516001600160a01b031662000295576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160401b038116600090815260076020526040902060018101546001600160a01b03166200039b57620002f78284606001517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3620004bb60201b60201c565b600282015560608301516001820180546001600160a01b039283166001600160a01b03199091161790556040808501518354610100600160481b03199190931669010000000000000000000216610100600160e81b031990921691909117610100178255516001600160401b03831681527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb99060200160405180910390a16200040a565b606083015160018201546001600160a01b039081169116141580620003df575060408301518154690100000000000000000090046001600160a01b03908116911614155b156200040a5760405163c39a620560e01b81526001600160401b03831660048201526024016200008f565b6020830151815490151560ff199091161781556040516001600160401b038316907fa73c588738263db34ef8c1942db8f99559bc6696f6a812d42e76bafb4c0e8d3090620004a0908490815460ff811615158252600881901c6001600160401b0316602083015260481c6001600160a01b0390811660408301526001830154166060820152600290910154608082015260a00190565b60405180910390a25050508060010190506200020a565b5050565b60a0805160408051602081018590526001600160401b0380881692820192909252911660608201526001600160a01b0384166080820152600091016040516020818303038152906040528051906020012090509392505050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171562000550576200055062000515565b60405290565b604051608081016001600160401b038111828210171562000550576200055062000515565b604051601f8201601f191681016001600160401b0381118282101715620005a657620005a662000515565b604052919050565b80516001600160401b0381168114620005c657600080fd5b919050565b80516001600160a01b0381168114620005c657600080fd5b600080828403606080821215620005f957600080fd5b6040808312156200060957600080fd5b620006136200052b565b92506200062086620005ae565b8352602062000631818801620005cb565b8482015260408701519395506001600160401b03808511156200065357600080fd5b848801945088601f8601126200066857600080fd5b8451818111156200067d576200067d62000515565b6200068d838260051b016200057b565b818152838101925060079190911b86018301908a821115620006ae57600080fd5b958301955b8187101562000732576080878c031215620006ce5760008081fd5b620006d862000556565b620006e388620005ae565b8152848801518015158114620006f95760008081fd5b818601526200070a888701620005cb565b868201526200071b878901620005cb565b8188015283526080969096019591830191620006b3565b809750505050505050509250929050565b634e487b7160e01b600052603260045260246000fd5b60805160a05160c051617093620007cf600039600081816102330152818161099001528181610fe30152818161173901528181611a2e015261342c0152600081816101fa015281816109630152613dde015260008181610b5b01528181610ba701528181611e7e0152611eca01526170936000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c80637f63b711116100f9578063ccd37ba311610097578063e9d68a8e11610071578063e9d68a8e1461059e578063f2fde38b146106c6578063f52121a5146106d9578063ff888fb1146106ec57600080fd5b8063ccd37ba314610533578063d2a15d3514610578578063d783efe71461058b57600080fd5b80638b364334116100d35780638b364334146104c55780638da5cb5b146104d857806396c62bcc14610500578063c673e5841461051357600080fd5b80637f63b7111461049c5780638456cb59146104af57806385572ffb146104b757600080fd5b8063311cd513116101665780635c975abb116101405780635c975abb1461032e5780635e36480c1461034e5780637437ff9f1461036e57806379ba50971461049457600080fd5b8063311cd513146103005780633f4ba83a14610313578063542625af1461031b57600080fd5b8063181f5a7711610197578063181f5a771461029157806329b980e4146102da5780632d04ab76146102ed57600080fd5b806305a754ec146101be57806306285c69146101d357806310c374ed1461026d575b600080fd5b6101d16101cc36600461511b565b6106ff565b005b6040805180820182526000808252602091820152815180830190925267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016825273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169082015260405161026491906151c8565b60405180910390f35b600b5464ffffffffff165b60405167ffffffffffffffff9091168152602001610264565b6102cd6040518060400160405280601d81526020017f45564d3245564d4d756c74694f666652616d7020312e362e302d64657600000081525081565b604051610264919061524f565b6101d16102e8366004615262565b6109f2565b6101d16102fb366004615321565b610a32565b6101d161030e3660046153d4565b610abf565b6101d1610af2565b6101d16103293660046159cf565b610b58565b600b5465010000000000900460ff165b6040519015158152602001610264565b61036161035c366004615afa565b610d85565b6040516102649190615b76565b6104876040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915250604080516101008101825260045473ffffffffffffffffffffffffffffffffffffffff808216835263ffffffff74010000000000000000000000000000000000000000830481166020850152780100000000000000000000000000000000000000000000000083048116948401949094527c01000000000000000000000000000000000000000000000000000000009091048316606083015260055461ffff8116608084015262010000810490931660a08301526601000000000000909204821660c082015260065490911660e082015290565b6040516102649190615c33565b6101d1610e19565b6101d16104aa366004615c42565b610f16565b6101d1610f2a565b6101d16101b9366004615d22565b6102786104d3366004615d5d565b610f92565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610264565b61033e61050e366004615d92565b610fa8565b610526610521366004615dc0565b611081565b6040516102649190615e2d565b61056a610541366004615eaf565b67ffffffffffffffff919091166000908152600a60209081526040808320938352929052205490565b604051908152602001610264565b6101d1610586366004615edb565b611212565b6101d1610599366004615fb6565b6112cc565b61065c6105ac366004615d92565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091525067ffffffffffffffff908116600090815260076020908152604091829020825160a081018452815460ff81161515825261010081049095169281019290925273ffffffffffffffffffffffffffffffffffffffff69010000000000000000009094048416928201929092526001820154909216606083015260020154608082015290565b6040516102649190600060a08201905082511515825267ffffffffffffffff6020840151166020830152604083015173ffffffffffffffffffffffffffffffffffffffff808216604085015280606086015116606085015250506080830151608083015292915050565b6101d16106d4366004616116565b61130e565b6101d16106e7366004616131565b61131f565b61033e6106fa366004616195565b6116d6565b6107076117a4565b60e081015173ffffffffffffffffffffffffffffffffffffffff16610758576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805173ffffffffffffffffffffffffffffffffffffffff166107a6576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160048054602080850151604080870151606088015173ffffffffffffffffffffffffffffffffffffffff9788167fffffffffffffffff000000000000000000000000000000000000000000000000909616959095177401000000000000000000000000000000000000000063ffffffff948516021777ffffffffffffffffffffffffffffffffffffffffffffffff167801000000000000000000000000000000000000000000000000918416919091027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16177c0100000000000000000000000000000000000000000000000000000000948316949094029390931790935560808501516005805460a088015160c089015161ffff9094167fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000000090921691909117620100009190961602949094177fffffffffffff0000000000000000000000000000000000000000ffffffffffff166601000000000000918616919091021790925560e0840151600680547fffffffffffffffffffffffff000000000000000000000000000000000000000016918516919091179055805180820182527f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681527f000000000000000000000000000000000000000000000000000000000000000090931691830191909152517fe62019b871510800fececa3aa514e63605c7ef4d56739b2117759c95c4bb46bf916109e79184906161ae565b60405180910390a150565b6109fa6117a4565b600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000001664ffffffffff92909216919091179055565b610a41878760208b0135611827565b610ab5600089898989898080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808d0282810182019093528c82529093508c92508b9182918501908490808284376000920191909152508a9250611d3a915050565b5050505050505050565b610ac9828261218c565b604080516000808252602082019092529050610aec600185858585866000611d3a565b50505050565b610afa6117a4565b600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff1690556040513381527f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa906020015b60405180910390a1565b467f000000000000000000000000000000000000000000000000000000000000000014610be8576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015267ffffffffffffffff461660248201526044015b60405180910390fd5b815181518114610c24576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b81811015610d75576000848281518110610c4357610c436161ee565b60200260200101519050600081602001515190506000858481518110610c6b57610c6b6161ee565b6020026020010151905080518214610caf576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b82811015610d66576000828281518110610cce57610cce6161ee565b6020026020010151905080600014158015610d09575084602001518281518110610cfa57610cfa6161ee565b60200260200101516080015181105b15610d5d5784516040517fc8e9605100000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024810183905260448101829052606401610bdf565b50600101610cb2565b50505050806001019050610c27565b50610d8083836121c8565b505050565b6000610d936001600461624c565b6002610da060808561628e565b67ffffffffffffffff16610db491906162b5565b67ffffffffffffffff8516600090815260096020526040812090610dd96080876162cc565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002054901c166003811115610e1057610e10615b33565b90505b92915050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610e9a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610bdf565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610f1e6117a4565b610f2781612278565b50565b610f326117a4565b600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffff16650100000000001790556040513381527f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25890602001610b4e565b600080610f9f8484612616565b50949350505050565b6040517f58babe3300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906358babe3390602401602060405180830381865afa15801561103f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061106391906162f3565b158015610e135750600b5465010000000000900460ff161592915050565b6110cc6040805161010081019091526000606082018181526080830182905260a0830182905260c0830182905260e08301919091528190815260200160608152602001606081525090565b60ff8083166000908152600260208181526040928390208351610100808201865282546060830190815260018401548089166080850152918204881660a08401526201000082048816151560c08401526301000000909104909616151560e08201529485529182018054845181840281018401909552808552929385830193909283018282801561119357602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611168575b505050505081526020016003820180548060200260200160405190810160405280929190818152602001828054801561120257602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116111d7575b5050505050815250509050919050565b61121a6117a4565b60005b81811015610d80576000838383818110611239576112396161ee565b90506040020180360381019061124f9190616310565b905061125e81602001516116d6565b6112c357805167ffffffffffffffff166000908152600a602090815260408083208285018051855290835281842093909355915191519182527f202f1139a3e334b6056064c0e9b19fd07e44a88d8f6e5ded571b24cf8c371f12910160405180910390a15b5060010161121d565b6112d46117a4565b60005b815181101561130a576113028282815181106112f5576112f56161ee565b602002602001015161274e565b6001016112d7565b5050565b6113166117a4565b610f2781612b86565b333014611358576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160008082526020820190925281611395565b604080518082019091526000808252602082015281526020019060019003908161136e5790505b506101408401515190915015611430576101408301516040805160608101909152602085015173ffffffffffffffffffffffffffffffffffffffff16608082015261142d91908060a0810160408051601f19818403018152918152908252875167ffffffffffffffff1660208301528781015173ffffffffffffffffffffffffffffffffffffffff1691015261016086015185612c7b565b90505b600061143c848361310c565b6005549091506601000000000000900473ffffffffffffffffffffffffffffffffffffffff16801561154d576040517fa219f6e500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82169063a219f6e5906114ba908590600401616405565b600060405180830381600087803b1580156114d457600080fd5b505af19250505080156114e5575060015b61154d573d808015611513576040519150601f19603f3d011682016040523d82523d6000602084013e611518565b606091505b50806040517f09c25325000000000000000000000000000000000000000000000000000000008152600401610bdf919061524f565b6101208501515115801561156357506080850151155b806115875750604085015173ffffffffffffffffffffffffffffffffffffffff163b155b806115d4575060408501516115d29073ffffffffffffffffffffffffffffffffffffffff167f85572ffb000000000000000000000000000000000000000000000000000000006131bc565b155b156115e0575050505050565b60048054608087015160408089015190517f3cf97983000000000000000000000000000000000000000000000000000000008152600094859473ffffffffffffffffffffffffffffffffffffffff1693633cf9798393611648938a9361138893929101616418565b6000604051808303816000875af1158015611667573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261168f91908101906164a6565b5091509150816116cd57806040517f0a8d6e8c000000000000000000000000000000000000000000000000000000008152600401610bdf919061524f565b50505050505050565b6040805180820182523081526020810183815291517f4d616771000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9081166004830152915160248201526000917f00000000000000000000000000000000000000000000000000000000000000001690634d61677190604401602060405180830381865afa158015611780573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1391906162f3565b60005473ffffffffffffffffffffffffffffffffffffffff163314611825576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610bdf565b565b600b5465010000000000900460ff161561186d576040517feced32bc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061187b8385018561668d565b8051515190915015158061189457508051602001515115155b156119bd57600b5464ffffffffff8084169116101561197e57600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000001664ffffffffff841617905560065481516040517f3937306f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90921691633937306f91611934916004016168ea565b600060405180830381600087803b15801561194e57600080fd5b505af1158015611962573d6000803e3d6000fd5b505050508060200151516000036119795750505050565b6119bd565b8060200151516000036119bd576040517f2261116700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b816020015151811015611cfc576000826020015182815181106119e5576119e56161ee565b602090810291909101015180516040517f58babe3300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152919250907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906358babe3390602401602060405180830381865afa158015611a8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aae91906162f3565b15611af1576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610bdf565b67ffffffffffffffff81166000908152600760205260409020805460ff16611b51576040517fed053c5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff83166004820152602401610bdf565b6020830151518154610100900467ffffffffffffffff9081169116141580611b90575060208084015190810151905167ffffffffffffffff9182169116115b15611bd057825160208401516040517feefb0cac000000000000000000000000000000000000000000000000000000008152610bdf9291906004016168fd565b6040830151611c0b576040517f504570e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b825167ffffffffffffffff166000908152600a6020908152604080832081870151845290915290205415611c8457825160408085015190517f32cf0cbf00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301526024820152604401610bdf565b6020808401510151611c97906001616932565b81547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff1661010067ffffffffffffffff92831602179091558251166000908152600a6020908152604080832094810151835293905291909120429055506001016119c0565b507f3a3950e13dd607cc37980db0ef14266c40d2bba9c01b2e44bfe549808883095d81604051611d2c919061695a565b60405180910390a150505050565b60ff8781166000908152600260209081526040808320815160a08101835281548152600190910154808616938201939093526101008304851691810191909152620100008204841615156060820152630100000090910490921615156080830152873590611da98760a46169f7565b9050826080015115611df1578451611dc29060206162b5565b8651611dcf9060206162b5565b611dda9060a06169f7565b611de491906169f7565b611dee90826169f7565b90505b368114611e33576040517f8e1192e100000000000000000000000000000000000000000000000000000000815260048101829052366024820152604401610bdf565b5081518114611e7b5781516040517f93df584c000000000000000000000000000000000000000000000000000000008152600481019190915260248101829052604401610bdf565b467f000000000000000000000000000000000000000000000000000000000000000014611efc576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006004820152466024820152604401610bdf565b60ff808a1660009081526003602090815260408083203384528252808320815180830190925280548086168352939491939092840191610100909104166002811115611f4a57611f4a615b33565b6002811115611f5b57611f5b615b33565b9052509050600281602001516002811115611f7857611f78615b33565b148015611fd95750600260008b60ff1660ff168152602001908152602001600020600301816000015160ff1681548110611fb457611fb46161ee565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633145b61200f576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5081608001511561213757600082606001511561205b5760028360200151846040015161203c9190616a0a565b6120469190616a23565b612051906001616a0a565b60ff169050612071565b602083015161206b906001616a0a565b60ff1690505b808651146120ab576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84518651146120e6576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50600087876040516120f9929190616a45565b604051908190038120612110918b90602001616a55565b6040516020818303038152906040528051906020012090506121358a828888886131d8565b505b6040805182815260208a81013560081c63ffffffff169082015260ff8b16917f198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef0910160405180910390a2505050505050505050565b61130a61219b82840184616a69565b60408051600080825260208201909252906121c6565b60608152602001906001900390816121b15790505b505b8151600003612202576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160408051600080825260208201909252911591905b845181101561227157612269858281518110612237576122376161ee565b60200260200101518461226357858381518110612256576122566161ee565b60200260200101516133f2565b836133f2565b600101612219565b5050505050565b60005b815181101561130a576000828281518110612298576122986161ee565b602002602001015190506000816000015190508067ffffffffffffffff166000036122ef576040517fc656089500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606082015173ffffffffffffffffffffffffffffffffffffffff16612340576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff81166000908152600760205260409020600181015473ffffffffffffffffffffffffffffffffffffffff1661249e576123a78284606001517f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3613dd8565b6002820155606083015160018201805473ffffffffffffffffffffffffffffffffffffffff9283167fffffffffffffffffffffffff000000000000000000000000000000000000000090911617905560408085015183547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff91909316690100000000000000000002167fffffff00000000000000000000000000000000000000000000000000000000ff909216919091176101001782555167ffffffffffffffff831681527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb99060200160405180910390a161253e565b6060830151600182015473ffffffffffffffffffffffffffffffffffffffff90811691161415806124fb5750604083015181546901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff908116911614155b1561253e576040517fc39a620500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff83166004820152602401610bdf565b602083015181549015157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090911617815560405167ffffffffffffffff8316907fa73c588738263db34ef8c1942db8f99559bc6696f6a812d42e76bafb4c0e8d3090612600908490815460ff811615158252600881901c67ffffffffffffffff16602083015260481c73ffffffffffffffffffffffffffffffffffffffff90811660408301526001830154166060820152600290910154608082015260a00190565b60405180910390a250505080600101905061227b565b67ffffffffffffffff808316600090815260086020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616845290915281205490918291168082036127405767ffffffffffffffff85166000908152600760205260409020546901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff16801561273e576040517f856c824700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015282169063856c824790602401602060405180830381865afa15801561270d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127319190616a9e565b6001935093505050612747565b505b9150600090505b9250929050565b806040015160ff166000036127925760006040517f367f56a2000000000000000000000000000000000000000000000000000000008152600401610bdf9190616abb565b60208082015160ff80821660009081526002909352604083206001810154929390928392169003612832576060840151600182018054608087015115156301000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffff9315156201000002939093167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000ffff909116179190911790556128ac565b6060840151600182015460ff6201000090910416151590151514158061287057506080840151600182015460ff630100000090910416151590151514155b156128ac576040517f87f6037c00000000000000000000000000000000000000000000000000000000815260ff84166004820152602401610bdf565b60c08401518051601f60ff821611156128f45760016040517f367f56a2000000000000000000000000000000000000000000000000000000008152600401610bdf9190616abb565b612967858560030180548060200260200160405190810160405280929190818152602001828054801561295d57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311612932575b5050505050613e68565b856080015115612ad5576129e2858560020180548060200260200160405190810160405280929190818152602001828054801561295d5760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311612932575050505050613e68565b60a086015180516129fc9060028701906020840190614ec7565b5080516001850180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1661010060ff841690810291909117909155601f1015612a755760026040517f367f56a2000000000000000000000000000000000000000000000000000000008152600401610bdf9190616abb565b6040880151612a85906003616ad5565b60ff168160ff1611612ac65760036040517f367f56a2000000000000000000000000000000000000000000000000000000008152600401610bdf9190616abb565b612ad287836001613efb565b50505b612ae185836002613efb565b8151612af69060038601906020850190614ec7565b506040868101516001850180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff8316179055875180865560c089015192517fab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f54793612b6d938a939260028b01929190616af1565b60405180910390a1612b7e856140f6565b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603612c05576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610bdf565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8360005b8551811015610f9f576000848281518110612c9c57612c9c6161ee565b6020026020010151806020019051810190612cb79190616b84565b90506000612cc88260200151614129565b9050612d0a73ffffffffffffffffffffffffffffffffffffffff82167faff2afbf000000000000000000000000000000000000000000000000000000006131bc565b612d58576040517fae9b4ce900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610bdf565b600080612ea8634059f55b60e01b6040518060e001604052808c6000015181526020018c6020015167ffffffffffffffff1681526020018c6040015173ffffffffffffffffffffffffffffffffffffffff1681526020018d8981518110612dc157612dc16161ee565b602002602001015160200151815260200187600001518152602001876040015181526020018a8981518110612df857612df86161ee565b6020026020010151815250604051602401612e139190616c39565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152600454859063ffffffff7c0100000000000000000000000000000000000000000000000000000000909104166113886084614184565b509150915081612ee657806040517fe1cd5509000000000000000000000000000000000000000000000000000000008152600401610bdf919061524f565b8051604014612f30578051604080517f78ef802400000000000000000000000000000000000000000000000000000000815260048101919091526024810191909152604401610bdf565b60008082806020019051810190612f479190616cf5565b915091506000612f56836142aa565b60408d810151815173ffffffffffffffffffffffffffffffffffffffff909116602482015260448082018690528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526004549192506130159183907801000000000000000000000000000000000000000000000000900463ffffffff166113886084614184565b50909550935084158061304557506000845111801561304557508380602001905181019061304391906162f3565b155b1561307e57836040517fe1cd5509000000000000000000000000000000000000000000000000000000008152600401610bdf919061524f565b80898981518110613091576130916161ee565b60200260200101516000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050818989815181106130e2576130e26161ee565b6020026020010151602001818152505050505050505050806001019050612c7f565b949350505050565b6040805160a08101825260008082526020820152606091810182905281810182905260808101919091526040518060a001604052808461018001518152602001846000015167ffffffffffffffff1681526020018460200151604051602001613191919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b6040516020818303038152906040528152602001846101200151815260200183815250905092915050565b60006131c783614323565b8015610e105750610e108383614387565b6131e0614f4d565b835160005b81811015610ab5576000600188868460208110613204576132046161ee565b61321191901a601b616a0a565b898581518110613223576132236161ee565b602002602001015189868151811061323d5761323d6161ee565b60200260200101516040516000815260200160405260405161327b949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa15801561329d573d6000803e3d6000fd5b505060408051601f1981015160ff808e1660009081526003602090815285822073ffffffffffffffffffffffffffffffffffffffff85168352815285822085870190965285548084168652939750909550929392840191610100900416600281111561330b5761330b615b33565b600281111561331c5761331c615b33565b905250905060018160200151600281111561333957613339615b33565b14613370576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051859060ff16601f8110613387576133876161ee565b6020020151156133c3576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600185826000015160ff16601f81106133de576133de6161ee565b9115156020909202015250506001016131e5565b81516040517f58babe3300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906358babe3390602401602060405180830381865afa158015613488573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134ac91906162f3565b156134ef576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610bdf565b602083015151600081900361352f576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836040015151811461356d576040517f57e0e08300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff82166000908152600760205260409020805460ff166135cd576040517fed053c5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610bdf565b60008267ffffffffffffffff8111156135e8576135e8614f81565b604051908082528060200260200182016040528015613611578160200160208202803683370190505b50905060005b838110156136d657600087602001518281518110613637576136376161ee565b6020026020010151905061364f818560020154614456565b838381518110613661576136616161ee565b602002602001018181525050806101800151838381518110613685576136856161ee565b6020026020010151146136cd578061018001516040517f345039be000000000000000000000000000000000000000000000000000000008152600401610bdf91815260200190565b50600101613617565b5060006136ed858389606001518a608001516145bf565b905080600003613735576040517f7dd17a7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86166004820152602401610bdf565b8551151560005b85811015613dcd5760008960200151828151811061375c5761375c6161ee565b602002602001015190506000613776898360600151610d85565b9050600281600381111561378c5761378c615b33565b036137e25760608201516040805167ffffffffffffffff808d16825290921660208301527f3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c910160405180910390a15050613dc5565b60008160038111156137f6576137f6615b33565b14806138135750600381600381111561381157613811615b33565b145b6138635760608201516040517f25507e7f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808c1660048301529091166024820152604401610bdf565b83156139445760045460009074010000000000000000000000000000000000000000900463ffffffff16613897874261624c565b11905080806138b7575060038260038111156138b5576138b5615b33565b145b6138f9576040517fa9cfc86200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b166004820152602401610bdf565b8a848151811061390b5761390b6161ee565b602002602001015160001461393e578a848151811061392c5761392c6161ee565b60200260200101518360800181815250505b506139a9565b600081600381111561395857613958615b33565b146139a95760608201516040517f3ef2a99c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808c1660048301529091166024820152604401610bdf565b6000806139ba8b8560200151612616565b915091508015613ada5760c084015167ffffffffffffffff166139de836001616932565b67ffffffffffffffff1614613a6e5760c084015160208501516040517f5444a3301c7c42dd164cbf6ba4b72bf02504f86c049b06a27fc2b662e334bdbd92613a5d928f9267ffffffffffffffff938416815291909216602082015273ffffffffffffffffffffffffffffffffffffffff91909116604082015260600190565b60405180910390a150505050613dc5565b67ffffffffffffffff8b811660009081526008602090815260408083208883015173ffffffffffffffffffffffffffffffffffffffff168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000169184169190911790555b6000836003811115613aee57613aee615b33565b03613b8c5760c084015167ffffffffffffffff16613b0d836001616932565b67ffffffffffffffff1614613b8c5760c084015160208501516040517f852dc8e405695593e311bd83991cf39b14a328f304935eac6d3d55617f911d8992613a5d928f9267ffffffffffffffff938416815291909216602082015273ffffffffffffffffffffffffffffffffffffffff91909116604082015260600190565b60008d604001518681518110613ba457613ba46161ee565b60200260200101519050613bd28561018001518d87606001518861014001515189610120015151865161465c565b613be28c86606001516001614764565b600080613bef8784614842565b91509150613c028e886060015184614764565b888015613c2057506003826003811115613c1e57613c1e615b33565b145b15613c6057866101800151816040517f2b11b8d9000000000000000000000000000000000000000000000000000000008152600401610bdf929190616d19565b6003826003811115613c7457613c74615b33565b14158015613c9457506002826003811115613c9157613c91615b33565b14155b15613cd5578d8760600151836040517f926c5a3e000000000000000000000000000000000000000000000000000000008152600401610bdf93929190616d32565b6000866003811115613ce957613ce9615b33565b03613d645767ffffffffffffffff808f1660009081526008602090815260408083208b83015173ffffffffffffffffffffffffffffffffffffffff168452909152812080549092169190613d3c83616d58565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b866101800151876060015167ffffffffffffffff168f67ffffffffffffffff167f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df28585604051613db5929190616d7f565b60405180910390a4505050505050505b60010161373c565b505050505050505050565b600081847f000000000000000000000000000000000000000000000000000000000000000085604051602001613e48949392919093845267ffffffffffffffff92831660208501529116604083015273ffffffffffffffffffffffffffffffffffffffff16606082015260800190565b6040516020818303038152906040528051906020012090505b9392505050565b60005b8151811015610d805760ff831660009081526003602052604081208351909190849084908110613e9d57613e9d6161ee565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000169055600101613e6b565b60005b82518160ff161015610aec576000838260ff1681518110613f2157613f216161ee565b6020026020010151905060006002811115613f3e57613f3e615b33565b60ff808716600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091529020546101009004166002811115613f8a57613f8a615b33565b14613fc45760046040517f367f56a2000000000000000000000000000000000000000000000000000000008152600401610bdf9190616abb565b73ffffffffffffffffffffffffffffffffffffffff8116614011576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405180604001604052808360ff16815260200184600281111561403757614037615b33565b905260ff808716600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845282529091208351815493167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00841681178255918401519092909183917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016176101008360028111156140dc576140dc615b33565b021790555090505050806140ef90616d9f565b9050613efe565b60ff8116610f2757600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000016905550565b6000815160201461416857816040517f8d666f60000000000000000000000000000000000000000000000000000000008152600401610bdf919061524f565b610e138280602001905181019061417f9190616dbe565b6142aa565b6000606060008361ffff1667ffffffffffffffff8111156141a7576141a7614f81565b6040519080825280601f01601f1916602001820160405280156141d1576020820181803683370190505b509150863b614204577f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b5a85811015614237577fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b8590036040810481038710614270577f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b505a6000808a5160208c0160008c8cf193505a900390503d848111156142935750835b808352806000602085013e50955095509592505050565b600073ffffffffffffffffffffffffffffffffffffffff8211806142ce5750600a82105b1561431f5760408051602081018490520160408051601f19818403018152908290527f8d666f60000000000000000000000000000000000000000000000000000000008252610bdf9160040161524f565b5090565b600061434f827f01ffc9a700000000000000000000000000000000000000000000000000000000614387565b8015610e135750614380827fffffffff00000000000000000000000000000000000000000000000000000000614387565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d9150600051905082801561443f575060208210155b801561444b5750600081115b979650505050505050565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b61010001516040516020016144f998979695949392919073ffffffffffffffffffffffffffffffffffffffff9889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b60405160208183030381529060405280519060200120856101200151805190602001208661014001516040516020016145329190616dd7565b6040516020818303038152906040528051906020012087610160015160405160200161455e9190616e99565b60408051601f198184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e0015b60405160208183030381529060405280519060200120905092915050565b600b5460009065010000000000900460ff1615614608576040517feced32bc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000614615858585614b66565b9050614620816116d6565b61462e576000915050613104565b67ffffffffffffffff86166000908152600a6020908152604080832093835292905220549050949350505050565b60055461ffff168311156146b0576040517fa1e5205a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808716600483015285166024820152604401610bdf565b8083146146fd576040517f1cfe6d8b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808716600483015285166024820152604401610bdf565b60055462010000900463ffffffff16821115612b7e576005546040517f1fd8fd04000000000000000000000000000000000000000000000000000000008152600481018890526201000090910463ffffffff16602482015260448101839052606401610bdf565b6000600261477360808561628e565b67ffffffffffffffff1661478791906162b5565b67ffffffffffffffff8516600090815260096020526040812091925090816147b06080876162cc565b67ffffffffffffffff1681526020810191909152604001600020549050816147da6001600461624c565b901b1916818360038111156147f1576147f1615b33565b67ffffffffffffffff871660009081526009602052604081209190921b929092179182916148206080886162cc565b67ffffffffffffffff1681526020810191909152604001600020555050505050565b6040517ff52121a5000000000000000000000000000000000000000000000000000000008152600090606090309063f52121a5906148869087908790600401616eac565b600060405180830381600087803b1580156148a057600080fd5b505af19250505080156148b1575060015b614b4b573d8080156148df576040519150601f19603f3d011682016040523d82523d6000602084013e6148e4565b606091505b5060006148f082617036565b90507f0a8d6e8c000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216148061498357507fe1cd5509000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b806149cf57507f8d666f60000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b80614a1b57507f78ef8024000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b80614a6757507f0c3b563c000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b80614ab357507fae9b4ce9000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b80614aff57507f09c25325000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216145b15614b105750600392509050612747565b856101800151826040517f2b11b8d9000000000000000000000000000000000000000000000000000000008152600401610bdf929190616d19565b50506040805160208101909152600081526002909250929050565b8251825160009190818303614ba7576040517f11a6b26400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6101018211801590614bbb57506101018111155b614bf1576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82820101610100811115614c52576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600003614c7f5786600081518110614c6d57614c6d6161ee565b60200260200101519350505050613e61565b60008167ffffffffffffffff811115614c9a57614c9a614f81565b604051908082528060200260200182016040528015614cc3578160200160208202803683370190505b50905060008080805b85811015614e065760006001821b8b811603614d275788851015614d10578c5160018601958e918110614d0157614d016161ee565b60200260200101519050614d49565b8551600185019487918110614d0157614d016161ee565b8b5160018401938d918110614d3e57614d3e6161ee565b602002602001015190505b600089861015614d79578d5160018701968f918110614d6a57614d6a6161ee565b60200260200101519050614d9b565b8651600186019588918110614d9057614d906161ee565b602002602001015190505b82851115614dd5576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b614ddf8282614e86565b878481518110614df157614df16161ee565b60209081029190910101525050600101614ccc565b506001850382148015614e1857508683145b8015614e2357508581145b614e59576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836001860381518110614e6e57614e6e6161ee565b60200260200101519750505050505050509392505050565b6000818310614e9e57614e998284614ea4565b610e10565b610e1083835b6040805160016020820152908101839052606081018290526000906080016145a1565b828054828255906000526020600020908101928215614f41579160200282015b82811115614f4157825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190614ee7565b5061431f929150614f6c565b604051806103e00160405280601f906020820280368337509192915050565b5b8082111561431f5760008155600101614f6d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff81118282101715614fd457614fd4614f81565b60405290565b6040805190810167ffffffffffffffff81118282101715614fd457614fd4614f81565b6040516101a0810167ffffffffffffffff81118282101715614fd457614fd4614f81565b60405160a0810167ffffffffffffffff81118282101715614fd457614fd4614f81565b6040516080810167ffffffffffffffff81118282101715614fd457614fd4614f81565b60405160e0810167ffffffffffffffff81118282101715614fd457614fd4614f81565b6040516060810167ffffffffffffffff81118282101715614fd457614fd4614f81565b604051601f8201601f1916810167ffffffffffffffff811182821017156150d6576150d6614f81565b604052919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461510257600080fd5b919050565b803563ffffffff8116811461510257600080fd5b6000610100828403121561512e57600080fd5b615136614fb0565b61513f836150de565b815261514d60208401615107565b602082015261515e60408401615107565b604082015261516f60608401615107565b6060820152608083013561ffff8116811461518957600080fd5b608082015261519a60a08401615107565b60a08201526151ab60c084016150de565b60c08201526151bc60e084016150de565b60e08201529392505050565b815167ffffffffffffffff16815260208083015173ffffffffffffffffffffffffffffffffffffffff169082015260408101610e13565b60005b8381101561521a578181015183820152602001615202565b50506000910152565b6000815180845261523b8160208601602086016151ff565b601f01601f19169290920160200192915050565b602081526000610e106020830184615223565b60006020828403121561527457600080fd5b813564ffffffffff81168114613e6157600080fd5b8060608101831015610e1357600080fd5b60008083601f8401126152ac57600080fd5b50813567ffffffffffffffff8111156152c457600080fd5b60208301915083602082850101111561274757600080fd5b60008083601f8401126152ee57600080fd5b50813567ffffffffffffffff81111561530657600080fd5b6020830191508360208260051b850101111561274757600080fd5b60008060008060008060008060e0898b03121561533d57600080fd5b6153478a8a615289565b9750606089013567ffffffffffffffff8082111561536457600080fd5b6153708c838d0161529a565b909950975060808b013591508082111561538957600080fd5b6153958c838d016152dc565b909750955060a08b01359150808211156153ae57600080fd5b506153bb8b828c016152dc565b999c989b50969995989497949560c00135949350505050565b6000806000608084860312156153e957600080fd5b6153f38585615289565b9250606084013567ffffffffffffffff81111561540f57600080fd5b61541b8682870161529a565b9497909650939450505050565b600067ffffffffffffffff82111561544257615442614f81565b5060051b60200190565b67ffffffffffffffff81168114610f2757600080fd5b80356151028161544c565b8015158114610f2757600080fd5b80356151028161546d565b600067ffffffffffffffff8211156154a0576154a0614f81565b50601f01601f191660200190565b600082601f8301126154bf57600080fd5b81356154d26154cd82615486565b6150ad565b8181528460208386010111156154e757600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261551557600080fd5b813560206155256154cd83615428565b82815260069290921b8401810191818101908684111561554457600080fd5b8286015b8481101561558a57604081890312156155615760008081fd5b615569614fda565b615572826150de565b81528185013585820152835291830191604001615548565b509695505050505050565b600082601f8301126155a657600080fd5b813560206155b66154cd83615428565b82815260059290921b840181019181810190868411156155d557600080fd5b8286015b8481101561558a57803567ffffffffffffffff8111156155f95760008081fd5b6156078986838b01016154ae565b8452509183019183016155d9565b60006101a0828403121561562857600080fd5b615630614ffd565b905061563b82615462565b8152615649602083016150de565b602082015261565a604083016150de565b604082015261566b60608301615462565b60608201526080820135608082015261568660a0830161547b565b60a082015261569760c08301615462565b60c08201526156a860e083016150de565b60e082015261010082810135908201526101208083013567ffffffffffffffff808211156156d557600080fd5b6156e1868387016154ae565b838501526101409250828501359150808211156156fd57600080fd5b61570986838701615504565b8385015261016092508285013591508082111561572557600080fd5b5061573285828601615595565b82840152505061018080830135818301525092915050565b600082601f83011261575b57600080fd5b8135602061576b6154cd83615428565b82815260059290921b8401810191818101908684111561578a57600080fd5b8286015b8481101561558a57803567ffffffffffffffff8111156157ae5760008081fd5b6157bc8986838b0101615615565b84525091830191830161578e565b600082601f8301126157db57600080fd5b813560206157eb6154cd83615428565b82815260059290921b8401810191818101908684111561580a57600080fd5b8286015b8481101561558a57803567ffffffffffffffff81111561582e5760008081fd5b61583c8986838b0101615595565b84525091830191830161580e565b600082601f83011261585b57600080fd5b8135602061586b6154cd83615428565b8083825260208201915060208460051b87010193508684111561588d57600080fd5b602086015b8481101561558a5780358352918301918301615892565b600082601f8301126158ba57600080fd5b813560206158ca6154cd83615428565b82815260059290921b840181019181810190868411156158e957600080fd5b8286015b8481101561558a57803567ffffffffffffffff8082111561590e5760008081fd5b818901915060a080601f19848d030112156159295760008081fd5b615931615021565b61593c888501615462565b8152604080850135848111156159525760008081fd5b6159608e8b8389010161574a565b8a84015250606080860135858111156159795760008081fd5b6159878f8c838a01016157ca565b83850152506080915081860135858111156159a25760008081fd5b6159b08f8c838a010161584a565b91840191909152509190930135908301525083529183019183016158ed565b60008060408084860312156159e357600080fd5b833567ffffffffffffffff808211156159fb57600080fd5b615a07878388016158a9565b9450602091508186013581811115615a1e57600080fd5b8601601f81018813615a2f57600080fd5b8035615a3d6154cd82615428565b81815260059190911b8201840190848101908a831115615a5c57600080fd5b8584015b83811015615ae857803586811115615a785760008081fd5b8501603f81018d13615a8a5760008081fd5b87810135615a9a6154cd82615428565b81815260059190911b82018a0190898101908f831115615aba5760008081fd5b928b01925b82841015615ad85783358252928a0192908a0190615abf565b8652505050918601918601615a60565b50809750505050505050509250929050565b60008060408385031215615b0d57600080fd5b8235615b188161544c565b91506020830135615b288161544c565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110615b7257615b72615b33565b9052565b60208101610e138284615b62565b73ffffffffffffffffffffffffffffffffffffffff8151168252602081015163ffffffff808216602085015280604084015116604085015280606084015116606085015261ffff60808401511660808501528060a08401511660a0850152505060c0810151615c0b60c084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060e0810151610d8060e084018273ffffffffffffffffffffffffffffffffffffffff169052565b6101008101610e138284615b84565b60006020808385031215615c5557600080fd5b823567ffffffffffffffff811115615c6c57600080fd5b8301601f81018513615c7d57600080fd5b8035615c8b6154cd82615428565b81815260079190911b82018301908381019087831115615caa57600080fd5b928401925b8284101561444b5760808489031215615cc85760008081fd5b615cd0615044565b8435615cdb8161544c565b815284860135615cea8161546d565b818701526040615cfb8682016150de565b908201526060615d0c8682016150de565b9082015282526080939093019290840190615caf565b600060208284031215615d3457600080fd5b813567ffffffffffffffff811115615d4b57600080fd5b820160a08185031215613e6157600080fd5b60008060408385031215615d7057600080fd5b8235615d7b8161544c565b9150615d89602084016150de565b90509250929050565b600060208284031215615da457600080fd5b8135613e618161544c565b803560ff8116811461510257600080fd5b600060208284031215615dd257600080fd5b610e1082615daf565b60008151808452602080850194506020840160005b83811015615e2257815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101615df0565b509495945050505050565b60208152600082518051602084015260ff602082015116604084015260ff60408201511660608401526060810151151560808401526080810151151560a084015250602083015160e060c0840152615e89610100840182615ddb565b90506040840151601f198483030160e0850152615ea68282615ddb565b95945050505050565b60008060408385031215615ec257600080fd5b8235615ecd8161544c565b946020939093013593505050565b60008060208385031215615eee57600080fd5b823567ffffffffffffffff80821115615f0657600080fd5b818501915085601f830112615f1a57600080fd5b813581811115615f2957600080fd5b8660208260061b8501011115615f3e57600080fd5b60209290920196919550909350505050565b600082601f830112615f6157600080fd5b81356020615f716154cd83615428565b8083825260208201915060208460051b870101935086841115615f9357600080fd5b602086015b8481101561558a57615fa9816150de565b8352918301918301615f98565b60006020808385031215615fc957600080fd5b823567ffffffffffffffff80821115615fe157600080fd5b818501915085601f830112615ff557600080fd5b81356160036154cd82615428565b81815260059190911b8301840190848101908883111561602257600080fd5b8585015b838110156161095780358581111561603d57600080fd5b860160e0818c03601f190112156160545760008081fd5b61605c615067565b888201358152604061606f818401615daf565b8a8301526060616080818501615daf565b828401526080915061609382850161547b565b9083015260a06160a484820161547b565b8284015260c0915081840135898111156160be5760008081fd5b6160cc8f8d83880101615f50565b82850152505060e0830135888111156160e55760008081fd5b6160f38e8c83870101615f50565b9183019190915250845250918601918601616026565b5098975050505050505050565b60006020828403121561612857600080fd5b610e10826150de565b6000806040838503121561614457600080fd5b823567ffffffffffffffff8082111561615c57600080fd5b61616886838701615615565b9350602085013591508082111561617e57600080fd5b5061618b85828601615595565b9150509250929050565b6000602082840312156161a757600080fd5b5035919050565b825167ffffffffffffffff16815260208084015173ffffffffffffffffffffffffffffffffffffffff16908201526101408101613e616040830184615b84565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810381811115610e1357610e1361621d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff808416806162a9576162a961625f565b92169190910692915050565b8082028115828204841417610e1357610e1361621d565b600067ffffffffffffffff808416806162e7576162e761625f565b92169190910492915050565b60006020828403121561630557600080fd5b8151613e618161546d565b60006040828403121561632257600080fd5b61632a614fda565b82356163358161544c565b81526020928301359281019290925250919050565b60008151808452602080850194506020840160005b83811015615e22578151805173ffffffffffffffffffffffffffffffffffffffff168852602090810151908801526040870196509082019060010161635f565b8051825267ffffffffffffffff60208201511660208301526000604082015160a060408501526163d260a0850182615223565b9050606083015184820360608601526163eb8282615223565b91505060808301518482036080860152615ea6828261634a565b602081526000610e10602083018461639f565b60808152600061642b608083018761639f565b61ffff95909516602083015250604081019290925273ffffffffffffffffffffffffffffffffffffffff16606090910152919050565b600082601f83011261647257600080fd5b81516164806154cd82615486565b81815284602083860101111561649557600080fd5b6131048260208301602087016151ff565b6000806000606084860312156164bb57600080fd5b83516164c68161546d565b602085015190935067ffffffffffffffff8111156164e357600080fd5b6164ef86828701616461565b925050604084015190509250925092565b80357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116811461510257600080fd5b600082601f83011261653d57600080fd5b8135602061654d6154cd83615428565b82815260069290921b8401810191818101908684111561656c57600080fd5b8286015b8481101561558a57604081890312156165895760008081fd5b616591614fda565b813561659c8161544c565b81526165a9828601616500565b81860152835291830191604001616570565b600082601f8301126165cc57600080fd5b813560206165dc6154cd83615428565b82815260079290921b840181019181810190868411156165fb57600080fd5b8286015b8481101561558a5780880360808112156166195760008081fd5b61662161508a565b823561662c8161544c565b81526040601f1983018113156166425760008081fd5b61664a614fda565b9250868401356166598161544c565b8352838101356166688161544c565b83880152818701929092526060830135918101919091528352918301916080016165ff565b600060208083850312156166a057600080fd5b823567ffffffffffffffff808211156166b857600080fd5b818501915060408083880312156166ce57600080fd5b6166d6614fda565b8335838111156166e557600080fd5b84016040818a0312156166f757600080fd5b6166ff614fda565b81358581111561670e57600080fd5b8201601f81018b1361671f57600080fd5b803561672d6154cd82615428565b81815260069190911b8201890190898101908d83111561674c57600080fd5b928a01925b8284101561679a5787848f0312156167695760008081fd5b616771614fda565b61677a856150de565b81526167878c8601616500565b818d0152825292870192908a0190616751565b8452505050818701359350848411156167b257600080fd5b6167be8a85840161652c565b81880152825250838501359150828211156167d857600080fd5b6167e4888386016165bb565b85820152809550505050505092915050565b805160408084528151848201819052600092602091908201906060870190855b8181101561686f578351805173ffffffffffffffffffffffffffffffffffffffff1684528501517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16858401529284019291850191600101616816565b50508583015187820388850152805180835290840192506000918401905b808310156168de578351805167ffffffffffffffff1683528501517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168583015292840192600192909201919085019061688d565b50979650505050505050565b602081526000610e1060208301846167f6565b67ffffffffffffffff8316815260608101613e616020830184805167ffffffffffffffff908116835260209182015116910152565b67ffffffffffffffff8181168382160190808211156169535761695361621d565b5092915050565b60006020808352606084516040808487015261697960608701836167f6565b87850151878203601f19016040890152805180835290860193506000918601905b8083101561610957845167ffffffffffffffff8151168352878101516169d989850182805167ffffffffffffffff908116835260209182015116910152565b5084015182870152938601936001929092019160809091019061699a565b80820180821115610e1357610e1361621d565b60ff8181168382160190811115610e1357610e1361621d565b600060ff831680616a3657616a3661625f565b8060ff84160491505092915050565b8183823760009101908152919050565b828152606082602083013760800192915050565b600060208284031215616a7b57600080fd5b813567ffffffffffffffff811115616a9257600080fd5b613104848285016158a9565b600060208284031215616ab057600080fd5b8151613e618161544c565b6020810160058310616acf57616acf615b33565b91905290565b60ff81811683821602908116908181146169535761695361621d565b600060a0820160ff88168352602087602085015260a0604085015281875480845260c086019150886000526020600020935060005b81811015616b5857845473ffffffffffffffffffffffffffffffffffffffff1683526001948501949284019201616b26565b50508481036060860152616b6c8188615ddb565b935050505060ff831660808301529695505050505050565b600060208284031215616b9657600080fd5b815167ffffffffffffffff80821115616bae57600080fd5b9083019060608286031215616bc257600080fd5b616bca61508a565b825182811115616bd957600080fd5b616be587828601616461565b825250602083015182811115616bfa57600080fd5b616c0687828601616461565b602083015250604083015182811115616c1e57600080fd5b616c2a87828601616461565b60408301525095945050505050565b602081526000825160e06020840152616c56610100840182615223565b905067ffffffffffffffff60208501511660408401526040840151616c93606085018273ffffffffffffffffffffffffffffffffffffffff169052565b50606084015160808401526080840151601f19808584030160a0860152616cba8383615223565b925060a08601519150808584030160c0860152616cd78383615223565b925060c08601519150808584030160e086015250615ea68282615223565b60008060408385031215616d0857600080fd5b505080516020909101519092909150565b8281526040602082015260006131046040830184615223565b67ffffffffffffffff848116825283166020820152606081016131046040830184615b62565b600067ffffffffffffffff808316818103616d7557616d7561621d565b6001019392505050565b616d898184615b62565b6040602082015260006131046040830184615223565b600060ff821660ff8103616db557616db561621d565b60010192915050565b600060208284031215616dd057600080fd5b5051919050565b6020808252825182820181905260009190848201906040850190845b81811015616e33578351805173ffffffffffffffffffffffffffffffffffffffff1684526020908101519084015260408301938501939250600101616df3565b50909695505050505050565b60008282518085526020808601955060208260051b8401016020860160005b84811015616e8c57601f19868403018952616e7a838351615223565b98840198925090830190600101616e5e565b5090979650505050505050565b602081526000610e106020830184616e3f565b60408152616ec760408201845167ffffffffffffffff169052565b60006020840151616ef0606084018273ffffffffffffffffffffffffffffffffffffffff169052565b50604084015173ffffffffffffffffffffffffffffffffffffffff8116608084015250606084015167ffffffffffffffff811660a084015250608084015160c083015260a084015180151560e08401525060c0840151610100616f5e8185018367ffffffffffffffff169052565b60e08601519150610120616f898186018473ffffffffffffffffffffffffffffffffffffffff169052565b81870151925061014091508282860152808701519250506101a06101608181870152616fb96101e0870185615223565b93508288015192507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0610180818887030181890152616ff8868661634a565b9550828a015194508188870301848901526170138686616e3f565b9550808a01516101c089015250505050508281036020840152615ea68185616e3f565b6000815160208301517fffffffff000000000000000000000000000000000000000000000000000000008082169350600483101561707e5780818460040360031b1b83161693505b50505091905056fea164736f6c6343000818000a",
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) GetLatestPriceEpochAndRound(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "getLatestPriceEpochAndRound")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) GetLatestPriceEpochAndRound() (uint64, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetLatestPriceEpochAndRound(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) GetLatestPriceEpochAndRound() (uint64, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetLatestPriceEpochAndRound(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) GetMerkleRoot(opts *bind.CallOpts, sourceChainSelector uint64, root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "getMerkleRoot", sourceChainSelector, root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) GetMerkleRoot(sourceChainSelector uint64, root [32]byte) (*big.Int, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetMerkleRoot(&_EVM2EVMMultiOffRamp.CallOpts, sourceChainSelector, root)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) GetMerkleRoot(sourceChainSelector uint64, root [32]byte) (*big.Int, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetMerkleRoot(&_EVM2EVMMultiOffRamp.CallOpts, sourceChainSelector, root)
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) IsBlessed(opts *bind.CallOpts, root [32]byte) (bool, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "isBlessed", root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) IsBlessed(root [32]byte) (bool, error) {
	return _EVM2EVMMultiOffRamp.Contract.IsBlessed(&_EVM2EVMMultiOffRamp.CallOpts, root)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) IsBlessed(root [32]byte) (bool, error) {
	return _EVM2EVMMultiOffRamp.Contract.IsBlessed(&_EVM2EVMMultiOffRamp.CallOpts, root)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) IsUnpausedAndNotCursed(opts *bind.CallOpts, sourceChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "isUnpausedAndNotCursed", sourceChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) IsUnpausedAndNotCursed(sourceChainSelector uint64) (bool, error) {
	return _EVM2EVMMultiOffRamp.Contract.IsUnpausedAndNotCursed(&_EVM2EVMMultiOffRamp.CallOpts, sourceChainSelector)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) IsUnpausedAndNotCursed(sourceChainSelector uint64) (bool, error) {
	return _EVM2EVMMultiOffRamp.Contract.IsUnpausedAndNotCursed(&_EVM2EVMMultiOffRamp.CallOpts, sourceChainSelector)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) LatestConfigDetails(opts *bind.CallOpts, ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "latestConfigDetails", ocrPluginType)

	if err != nil {
		return *new(MultiOCR3BaseOCRConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(MultiOCR3BaseOCRConfig)).(*MultiOCR3BaseOCRConfig)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) LatestConfigDetails(ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error) {
	return _EVM2EVMMultiOffRamp.Contract.LatestConfigDetails(&_EVM2EVMMultiOffRamp.CallOpts, ocrPluginType)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) LatestConfigDetails(ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error) {
	return _EVM2EVMMultiOffRamp.Contract.LatestConfigDetails(&_EVM2EVMMultiOffRamp.CallOpts, ocrPluginType)
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) Paused() (bool, error) {
	return _EVM2EVMMultiOffRamp.Contract.Paused(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) Paused() (bool, error) {
	return _EVM2EVMMultiOffRamp.Contract.Paused(&_EVM2EVMMultiOffRamp.CallOpts)
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) Commit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "commit", reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) Commit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.Commit(&_EVM2EVMMultiOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) Commit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.Commit(&_EVM2EVMMultiOffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) Execute(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "execute", reportContext, report)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) Execute(reportContext [3][32]byte, report []byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.Execute(&_EVM2EVMMultiOffRamp.TransactOpts, reportContext, report)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) Execute(reportContext [3][32]byte, report []byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.Execute(&_EVM2EVMMultiOffRamp.TransactOpts, reportContext, report)
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "pause")
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.Pause(&_EVM2EVMMultiOffRamp.TransactOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.Pause(&_EVM2EVMMultiOffRamp.TransactOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) ResetUnblessedRoots(opts *bind.TransactOpts, rootToReset []EVM2EVMMultiOffRampUnblessedRoot) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "resetUnblessedRoots", rootToReset)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) ResetUnblessedRoots(rootToReset []EVM2EVMMultiOffRampUnblessedRoot) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.ResetUnblessedRoots(&_EVM2EVMMultiOffRamp.TransactOpts, rootToReset)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) ResetUnblessedRoots(rootToReset []EVM2EVMMultiOffRampUnblessedRoot) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.ResetUnblessedRoots(&_EVM2EVMMultiOffRamp.TransactOpts, rootToReset)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMMultiOffRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) SetDynamicConfig(dynamicConfig EVM2EVMMultiOffRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.SetDynamicConfig(&_EVM2EVMMultiOffRamp.TransactOpts, dynamicConfig)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) SetDynamicConfig(dynamicConfig EVM2EVMMultiOffRampDynamicConfig) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.SetDynamicConfig(&_EVM2EVMMultiOffRamp.TransactOpts, dynamicConfig)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) SetLatestPriceEpochAndRound(opts *bind.TransactOpts, latestPriceEpochAndRound *big.Int) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "setLatestPriceEpochAndRound", latestPriceEpochAndRound)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) SetLatestPriceEpochAndRound(latestPriceEpochAndRound *big.Int) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.SetLatestPriceEpochAndRound(&_EVM2EVMMultiOffRamp.TransactOpts, latestPriceEpochAndRound)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) SetLatestPriceEpochAndRound(latestPriceEpochAndRound *big.Int) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.SetLatestPriceEpochAndRound(&_EVM2EVMMultiOffRamp.TransactOpts, latestPriceEpochAndRound)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) SetOCR3Configs(opts *bind.TransactOpts, ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "setOCR3Configs", ocrConfigArgs)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) SetOCR3Configs(ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.SetOCR3Configs(&_EVM2EVMMultiOffRamp.TransactOpts, ocrConfigArgs)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) SetOCR3Configs(ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.SetOCR3Configs(&_EVM2EVMMultiOffRamp.TransactOpts, ocrConfigArgs)
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "unpause")
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.Unpause(&_EVM2EVMMultiOffRamp.TransactOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.Unpause(&_EVM2EVMMultiOffRamp.TransactOpts)
}

type EVM2EVMMultiOffRampCommitReportAcceptedIterator struct {
	Event *EVM2EVMMultiOffRampCommitReportAccepted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampCommitReportAcceptedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampCommitReportAccepted)
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
		it.Event = new(EVM2EVMMultiOffRampCommitReportAccepted)
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

func (it *EVM2EVMMultiOffRampCommitReportAcceptedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampCommitReportAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampCommitReportAccepted struct {
	Report EVM2EVMMultiOffRampCommitReport
	Raw    types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterCommitReportAccepted(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampCommitReportAcceptedIterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "CommitReportAccepted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampCommitReportAcceptedIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "CommitReportAccepted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchCommitReportAccepted(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampCommitReportAccepted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "CommitReportAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampCommitReportAccepted)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "CommitReportAccepted", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseCommitReportAccepted(log types.Log) (*EVM2EVMMultiOffRampCommitReportAccepted, error) {
	event := new(EVM2EVMMultiOffRampCommitReportAccepted)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "CommitReportAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	OcrPluginType uint8
	ConfigDigest  [32]byte
	Signers       []common.Address
	Transmitters  []common.Address
	F             uint8
	Raw           types.Log
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

type EVM2EVMMultiOffRampPausedIterator struct {
	Event *EVM2EVMMultiOffRampPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampPaused)
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
		it.Event = new(EVM2EVMMultiOffRampPaused)
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

func (it *EVM2EVMMultiOffRampPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampPausedIterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampPausedIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampPaused)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParsePaused(log types.Log) (*EVM2EVMMultiOffRampPaused, error) {
	event := new(EVM2EVMMultiOffRampPaused)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMMultiOffRampRootRemovedIterator struct {
	Event *EVM2EVMMultiOffRampRootRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampRootRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampRootRemoved)
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
		it.Event = new(EVM2EVMMultiOffRampRootRemoved)
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

func (it *EVM2EVMMultiOffRampRootRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampRootRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampRootRemoved struct {
	Root [32]byte
	Raw  types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterRootRemoved(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampRootRemovedIterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "RootRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampRootRemovedIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "RootRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchRootRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampRootRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "RootRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampRootRemoved)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "RootRemoved", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseRootRemoved(log types.Log) (*EVM2EVMMultiOffRampRootRemoved, error) {
	event := new(EVM2EVMMultiOffRampRootRemoved)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "RootRemoved", log); err != nil {
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
	OcrPluginType  uint8
	ConfigDigest   [32]byte
	SequenceNumber uint64
	Raw            types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterTransmitted(opts *bind.FilterOpts, ocrPluginType []uint8) (*EVM2EVMMultiOffRampTransmittedIterator, error) {

	var ocrPluginTypeRule []interface{}
	for _, ocrPluginTypeItem := range ocrPluginType {
		ocrPluginTypeRule = append(ocrPluginTypeRule, ocrPluginTypeItem)
	}

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "Transmitted", ocrPluginTypeRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampTransmittedIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampTransmitted, ocrPluginType []uint8) (event.Subscription, error) {

	var ocrPluginTypeRule []interface{}
	for _, ocrPluginTypeItem := range ocrPluginType {
		ocrPluginTypeRule = append(ocrPluginTypeRule, ocrPluginTypeItem)
	}

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "Transmitted", ocrPluginTypeRule)
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

type EVM2EVMMultiOffRampUnpausedIterator struct {
	Event *EVM2EVMMultiOffRampUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampUnpaused)
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
		it.Event = new(EVM2EVMMultiOffRampUnpaused)
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

func (it *EVM2EVMMultiOffRampUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampUnpausedIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampUnpaused)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseUnpaused(log types.Log) (*EVM2EVMMultiOffRampUnpaused, error) {
	event := new(EVM2EVMMultiOffRampUnpaused)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMMultiOffRamp.abi.Events["CommitReportAccepted"].ID:
		return _EVM2EVMMultiOffRamp.ParseCommitReportAccepted(log)
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
	case _EVM2EVMMultiOffRamp.abi.Events["Paused"].ID:
		return _EVM2EVMMultiOffRamp.ParsePaused(log)
	case _EVM2EVMMultiOffRamp.abi.Events["RootRemoved"].ID:
		return _EVM2EVMMultiOffRamp.ParseRootRemoved(log)
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
	case _EVM2EVMMultiOffRamp.abi.Events["Unpaused"].ID:
		return _EVM2EVMMultiOffRamp.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMMultiOffRampCommitReportAccepted) Topic() common.Hash {
	return common.HexToHash("0x3a3950e13dd607cc37980db0ef14266c40d2bba9c01b2e44bfe549808883095d")
}

func (EVM2EVMMultiOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xe62019b871510800fececa3aa514e63605c7ef4d56739b2117759c95c4bb46bf")
}

func (EVM2EVMMultiOffRampConfigSet0) Topic() common.Hash {
	return common.HexToHash("0xab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f547")
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

func (EVM2EVMMultiOffRampPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMMultiOffRampRootRemoved) Topic() common.Hash {
	return common.HexToHash("0x202f1139a3e334b6056064c0e9b19fd07e44a88d8f6e5ded571b24cf8c371f12")
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
	return common.HexToHash("0xa73c588738263db34ef8c1942db8f99559bc6696f6a812d42e76bafb4c0e8d30")
}

func (EVM2EVMMultiOffRampSourceChainSelectorAdded) Topic() common.Hash {
	return common.HexToHash("0xf4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb9")
}

func (EVM2EVMMultiOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0x198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef0")
}

func (EVM2EVMMultiOffRampUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRamp) Address() common.Address {
	return _EVM2EVMMultiOffRamp.address
}

type EVM2EVMMultiOffRampInterface interface {
	CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMMultiOffRampDynamicConfig, error)

	GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error)

	GetLatestPriceEpochAndRound(opts *bind.CallOpts) (uint64, error)

	GetMerkleRoot(opts *bind.CallOpts, sourceChainSelector uint64, root [32]byte) (*big.Int, error)

	GetSenderNonce(opts *bind.CallOpts, sourceChainSelector uint64, sender common.Address) (uint64, error)

	GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (EVM2EVMMultiOffRampSourceChainConfig, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMMultiOffRampStaticConfig, error)

	IsBlessed(opts *bind.CallOpts, root [32]byte) (bool, error)

	IsUnpausedAndNotCursed(opts *bind.CallOpts, sourceChainSelector uint64) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts, ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []EVM2EVMMultiOffRampSourceChainConfigArgs) (*types.Transaction, error)

	Commit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, reports []InternalExecutionReportSingleChain, gasLimitOverrides [][]*big.Int) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	ResetUnblessedRoots(opts *bind.TransactOpts, rootToReset []EVM2EVMMultiOffRampUnblessedRoot) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMMultiOffRampDynamicConfig) (*types.Transaction, error)

	SetLatestPriceEpochAndRound(opts *bind.TransactOpts, latestPriceEpochAndRound *big.Int) (*types.Transaction, error)

	SetOCR3Configs(opts *bind.TransactOpts, ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterCommitReportAccepted(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampCommitReportAcceptedIterator, error)

	WatchCommitReportAccepted(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampCommitReportAccepted) (event.Subscription, error)

	ParseCommitReportAccepted(log types.Log) (*EVM2EVMMultiOffRampCommitReportAccepted, error)

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

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMMultiOffRampPaused, error)

	FilterRootRemoved(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampRootRemovedIterator, error)

	WatchRootRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampRootRemoved) (event.Subscription, error)

	ParseRootRemoved(log types.Log) (*EVM2EVMMultiOffRampRootRemoved, error)

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

	FilterTransmitted(opts *bind.FilterOpts, ocrPluginType []uint8) (*EVM2EVMMultiOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampTransmitted, ocrPluginType []uint8) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMMultiOffRampTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMMultiOffRampUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
