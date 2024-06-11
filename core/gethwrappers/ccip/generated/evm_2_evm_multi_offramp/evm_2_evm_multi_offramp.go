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
	MetadataHash [32]byte
}

type EVM2EVMMultiOffRampSourceChainConfigArgs struct {
	SourceChainSelector uint64
	IsEnabled           bool
	PrevOffRamp         common.Address
	MetadataHash        [32]byte
}

type EVM2EVMMultiOffRampStaticConfig struct {
	ChainSelector      uint64
	RmnProxy           common.Address
	TokenAdminRegistry common.Address
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
	IsSignatureVerificationEnabled bool
	Signers                        []common.Address
	Transmitters                   []common.Address
}

var EVM2EVMMultiOffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"DuplicateMetadataHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumMultiOCR3Base.InvalidConfigErrorType\",\"name\":\"errorType\",\"type\":\"uint8\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.Interval\",\"name\":\"interval\",\"type\":\"tuple\"}],\"name\":\"InvalidInterval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"InvalidManualExecutionGasLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"InvalidMessageId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"InvalidNewState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRoot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidStaticConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeavesCannotBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionGasLimitMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"errorReason\",\"type\":\"bytes\"}],\"name\":\"MessageValidationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notPool\",\"type\":\"address\"}],\"name\":\"NotACompatiblePool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"RootAlreadyCommitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleCommitReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"}],\"name\":\"StaticConfigCannotBeChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"TokenDataMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"TokenHandlingError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedTokenData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroChainSelectorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroMetadataHashNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint224\",\"name\":\"usdPerToken\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint224\",\"name\":\"usdPerUnitGas\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.GasPriceUpdate[]\",\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.PriceUpdates\",\"name\":\"priceUpdates\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.Interval\",\"name\":\"interval\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.MerkleRoot[]\",\"name\":\"merkleRoots\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.CommitReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"CommitReportAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"DynamicConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldSequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newSequenceNumber\",\"type\":\"uint64\"}],\"name\":\"LatestPriceSequenceNumberSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"RootRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"SkippedAlreadyExecutedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedSenderWithPreviousRampMessageInflight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfig\",\"name\":\"sourceConfig\",\"type\":\"tuple\"}],\"name\":\"SourceChainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainSelectorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"}],\"name\":\"StaticConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\"}],\"name\":\"applySourceChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPriceSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"getSourceChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"isBlessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"}],\"name\":\"latestConfigDetails\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"n\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\"}],\"internalType\":\"structMultiOCR3Base.ConfigInfo\",\"name\":\"configInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"}],\"internalType\":\"structMultiOCR3Base.OCRConfig\",\"name\":\"ocrConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReportSingleChain[]\",\"name\":\"reports\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"gasLimitOverrides\",\"type\":\"uint256[][]\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.UnblessedRoot[]\",\"name\":\"rootToReset\",\"type\":\"tuple[]\"}],\"name\":\"resetUnblessedRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"latestPriceSequenceNumber\",\"type\":\"uint64\"}],\"name\":\"setLatestPriceSequenceNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"}],\"internalType\":\"structMultiOCR3Base.OCRConfigArgs[]\",\"name\":\"ocrConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"setOCR3Configs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b5060405162006690380380620066908339810160408190526200003591620005cc565b33806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf81620001c1565b5050466080525060208201516001600160a01b03161580620000ec575060408201516001600160a01b0316155b156200010b576040516342bcdf7f60e11b815260040160405180910390fd5b81516001600160401b0316600003620001375760405163c656089560e01b815260040160405180910390fd5b81516001600160401b0390811660a052602080840180516001600160a01b0390811660c05260408087018051831660e052815188519096168652925182169385019390935290511682820152517f2f56698ec552a5d53d27d6f4b3dd8b6989f6426b6151a36aff61160c1d07efdf9181900360600190a1620001b9816200026c565b50506200074b565b336001600160a01b038216036200021b5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60005b8151811015620004fa57600082828151811062000290576200029062000735565b60200260200101519050600081600001519050806001600160401b0316600003620002ce5760405163c656089560e01b815260040160405180910390fd5b6001600160401b03811660009081526007602052604090206001810154620003f85760608301518062000314576040516330a68d4d60e01b815260040160405180910390fd5b60008181526008602052604090205460ff16156200034957604051630f940f8760e31b81526004810182905260240162000083565b60008181526008602052604090819020805460ff191660019081179091558301829055848101518354610100600160481b03196001600160a01b0390921669010000000000000000000291909116610100600160e81b031990911617610100178355517ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb990620003e99085906001600160401b0391909116815260200190565b60405180910390a1506200045a565b826060015181600101541415806200042f575060408301518154690100000000000000000090046001600160a01b03908116911614155b156200045a5760405163c39a620560e01b81526001600160401b038316600482015260240162000083565b6020830151815490151560ff199091161781556040516001600160401b038316907fb8f0c74385134334c728fcac437ec6c6397c9f2c1440532b0c44175a090b140190620004e3908490815460ff811615158252600881901c6001600160401b0316602083015260481c6001600160a01b03166040820152600190910154606082015260800190565b60405180910390a25050508060010190506200026f565b5050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715620005395762000539620004fe565b60405290565b604051608081016001600160401b0381118282101715620005395762000539620004fe565b604051601f8201601f191681016001600160401b03811182821017156200058f576200058f620004fe565b604052919050565b80516001600160401b0381168114620005af57600080fd5b919050565b80516001600160a01b0381168114620005af57600080fd5b600080828403608080821215620005e257600080fd5b606080831215620005f257600080fd5b620005fc62000514565b9250620006098662000597565b835260206200061a818801620005b4565b8185015260406200062e60408901620005b4565b604086015260608801519496506001600160401b03808611156200065157600080fd5b858901955089601f8701126200066657600080fd5b8551818111156200067b576200067b620004fe565b6200068b848260051b0162000564565b818152848101925060079190911b87018401908b821115620006ac57600080fd5b968401965b81881015620007235786888d031215620006cb5760008081fd5b620006d56200053f565b620006e08962000597565b8152858901518015158114620006f65760008081fd5b8187015262000707898601620005b4565b81860152888701518782015283529686019691840191620006b1565b80985050505050505050509250929050565b634e487b7160e01b600052603260045260246000fd5b60805160a05160c05160e051615eef620007a16000396000818161021e01526131420152600081816101ef015281816115d4015261168b015260006101bf015260008181611c270152611c730152615eef6000f3fe608060405234801561001057600080fd5b506004361061018d5760003560e01c806385572ffb116100e3578063d383e6d91161008c578063f52121a511610066578063f52121a51461065e578063f716f99f14610671578063ff888fb11461068457600080fd5b8063d383e6d914610552578063e9d68a8e14610565578063f2fde38b1461064b57600080fd5b8063c673e584116100bd578063c673e584146104da578063ccd37ba3146104fa578063d2a15d351461053f57600080fd5b806385572ffb1461049e5780638b364334146104ac5780638da5cb5b146104bf57600080fd5b8063403b2d631161014557806369600bca1161011f57806369600bca1461036e5780637437ff9f1461038157806379ba50971461049657600080fd5b8063403b2d6314610328578063542625af1461033b5780635e36480c1461034e57600080fd5b80632d04ab76116101765780632d04ab76146102d9578063311cd513146102ee5780633f4b04aa1461030157600080fd5b806306285c6914610192578063181f5a7714610290575b600080fd5b61024e604080516060810182526000808252602082018190529181019190915260405180606001604052807f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b60408051825167ffffffffffffffff1681526020808401516001600160a01b039081169183019190915292820151909216908201526060015b60405180910390f35b6102cc6040518060400160405280601d81526020017f45564d3245564d4d756c74694f666652616d7020312e362e302d64657600000081525081565b604051610287919061417f565b6102ec6102e736600461422a565b6106a7565b005b6102ec6102fc3660046142dd565b610a6d565b600c5467ffffffffffffffff165b60405167ffffffffffffffff9091168152602001610287565b6102ec61033636600461448a565b610ad6565b6102ec610349366004614aae565b610c93565b61036161035c366004614bd9565b610e38565b6040516102879190614c3c565b6102ec61037c366004614c4a565b610e8e565b61042d6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c0810182526004546001600160a01b03808216835263ffffffff74010000000000000000000000000000000000000000830481166020850152600160c01b8304811694840194909452600160e01b90910490921660608201526005548216608082015260065490911660a082015290565b6040516102879190600060c0820190506001600160a01b03808451168352602084015163ffffffff808216602086015280604087015116604086015280606087015116606086015250508060808501511660808401528060a08501511660a08401525092915050565b6102ec610ef9565b6102ec61018d366004614c67565b61030f6104ba366004614ca2565b610fb7565b6000546040516001600160a01b039091168152602001610287565b6104ed6104e8366004614ce1565b610fcd565b6040516102879190614d41565b610531610508366004614db6565b67ffffffffffffffff919091166000908152600b60209081526040808320938352929052205490565b604051908152602001610287565b6102ec61054d366004614de2565b61112b565b6102ec610560366004614e57565b6111e5565b610603610573366004614c4a565b6040805160808101825260008082526020820181905291810182905260608101919091525067ffffffffffffffff9081166000908152600760209081526040918290208251608081018452815460ff81161515825261010081049095169281019290925269010000000000000000009093046001600160a01b031691810191909152600190910154606082015290565b604051610287919081511515815260208083015167ffffffffffffffff16908201526040808301516001600160a01b0316908201526060918201519181019190915260800190565b6102ec610659366004614f31565b6111f9565b6102ec61066c366004614f4e565b61120a565b6102ec61067f36600461501a565b61153c565b610697610692366004615165565b61157e565b6040519015158152602001610287565b60006106b5878901896152f6565b805151519091501515806106ce57508051602001515115155b156107ce57600c5460208a01359067ffffffffffffffff8083169116101561078d57600c805467ffffffffffffffff191667ffffffffffffffff831617905560065482516040517f3937306f0000000000000000000000000000000000000000000000000000000081526001600160a01b0390921691633937306f9161075691600401615534565b600060405180830381600087803b15801561077057600080fd5b505af1158015610784573d6000803e3d6000fd5b505050506107cc565b8160200151516000036107cc576040517f2261116700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b60005b8160200151518110156109b6576000826020015182815181106107f6576107f6615461565b602002602001015190506000816000015190506108128161163f565b600061081d82611741565b602084015151815491925067ffffffffffffffff9081166101009092041614158061085f575060208084015190810151905167ffffffffffffffff9182169116115b156108a857825160208401516040517feefb0cac00000000000000000000000000000000000000000000000000000000815261089f929190600401615547565b60405180910390fd5b6040830151806108e4576040517f504570e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835167ffffffffffffffff166000908152600b60209081526040808320848452909152902054156109575783516040517f32cf0cbf00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024810182905260440161089f565b602080850151015161096a906001615592565b825468ffffffffffffffff00191661010067ffffffffffffffff92831602179092559251166000908152600b6020908152604080832094835293905291909120429055506001016107d1565b507f3a3950e13dd607cc37980db0ef14266c40d2bba9c01b2e44bfe549808883095d816040516109e691906155ba565b60405180910390a1610a6260008a8a8a8a8a8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c9182918501908490808284376000920191909152508b92506117a1915050565b505050505050505050565b610aad610a7c82840184615657565b6040805160008082526020820190925290610aa7565b6060815260200190600190039081610a925790505b50611b18565b604080516000808252602082019092529050610ad06001858585858660006117a1565b50505050565b610ade611bc8565b60a08101516001600160a01b03161580610b00575080516001600160a01b0316155b15610b37576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516004805460208085018051604080880180516060808b0180516001600160a01b039b8c167fffffffffffffffff000000000000000000000000000000000000000000000000909a168a177401000000000000000000000000000000000000000063ffffffff988916021777ffffffffffffffffffffffffffffffffffffffffffffffff16600160c01b948816949094026001600160e01b031693909317600160e01b93871693909302929092179098556080808b0180516005805473ffffffffffffffffffffffffffffffffffffffff19908116928e1692909217905560a0808e01805160068054909416908f161790925586519a8b5297518716988a0198909852925185169388019390935251909216958501959095525185169383019390935251909216908201527f0da37fd00459f4f5f0b8210d31525e4910ae674b8bab34b561d146bb45773a4c9060c00160405180910390a150565b610c9b611c24565b815181518114610cd7576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b81811015610e28576000848281518110610cf657610cf6615461565b60200260200101519050600081602001515190506000858481518110610d1e57610d1e615461565b6020026020010151905080518214610d62576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b82811015610e19576000828281518110610d8157610d81615461565b6020026020010151905080600014158015610dbc575084602001518281518110610dad57610dad615461565b60200260200101516080015181105b15610e105784516040517fc8e9605100000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602481018390526044810182905260640161089f565b50600101610d65565b50505050806001019050610cda565b50610e338383611b18565b505050565b6000610e466001600461568c565b6002610e536080856156b5565b67ffffffffffffffff16610e6791906156dc565b610e718585611ca5565b901c166003811115610e8557610e85614c12565b90505b92915050565b610e96611bc8565b600c805467ffffffffffffffff83811667ffffffffffffffff1983168117909355604080519190921680825260208201939093527f88ad9c61d6caf19a2af116a871802a03a31e680115a2dd20e8c08801d7c82f83910160405180910390a15050565b6001546001600160a01b03163314610f535760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161089f565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b600080610fc48484611cec565b50949350505050565b6110106040805160e081019091526000606082018181526080830182905260a0830182905260c08301919091528190815260200160608152602001606081525090565b60ff808316600090815260026020818152604092839020835160e081018552815460608201908152600183015480881660808401526101008104881660a0840152620100009004909616151560c0820152948552918201805484518184028101840190955280855292938583019390928301828280156110b957602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161109b575b505050505081526020016003820180548060200260200160405190810160405280929190818152602001828054801561111b57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116110fd575b5050505050815250509050919050565b611133611bc8565b60005b81811015610e3357600083838381811061115257611152615461565b90506040020180360381019061116891906156f3565b9050611177816020015161157e565b6111dc57805167ffffffffffffffff166000908152600b602090815260408083208285018051855290835281842093909355915191519182527f202f1139a3e334b6056064c0e9b19fd07e44a88d8f6e5ded571b24cf8c371f12910160405180910390a15b50600101611136565b6111ed611bc8565b6111f681611dfd565b50565b611201611bc8565b6111f6816120fd565b333014611243576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160008082526020820190925281611280565b60408051808201909152600080825260208201528152602001906001900390816112595790505b5061014084015151909150156112e1576112de83610140015184602001516040516020016112bd91906001600160a01b0391909116815260200190565b60408051601f198184030181529181528601518651610160880151876121b3565b90505b60006112ed8483612260565b6005549091506001600160a01b031680156113da576040517fa219f6e50000000000000000000000000000000000000000000000000000000081526001600160a01b0382169063a219f6e5906113479085906004016157db565b600060405180830381600087803b15801561136157600080fd5b505af1925050508015611372575060015b6113da573d8080156113a0576040519150601f19603f3d011682016040523d82523d6000602084013e6113a5565b606091505b50806040517f09c2532500000000000000000000000000000000000000000000000000000000815260040161089f919061417f565b610120850151511580156113f057506080850151155b80611407575060408501516001600160a01b03163b155b8061144757506040850151611445906001600160a01b03167f85572ffb00000000000000000000000000000000000000000000000000000000612303565b155b15611453575050505050565b60048054608087015160408089015190517f3cf9798300000000000000000000000000000000000000000000000000000000815260009485946001600160a01b031693633cf97983936114ae938a93611388939291016157ee565b6000604051808303816000875af11580156114cd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526114f5919081019061586f565b50915091508161153357806040517f0a8d6e8c00000000000000000000000000000000000000000000000000000000815260040161089f919061417f565b50505050505050565b611544611bc8565b60005b815181101561157a5761157282828151811061156557611565615461565b602002602001015161231f565b600101611547565b5050565b6040805180820182523081526020810183815291517f4d61677100000000000000000000000000000000000000000000000000000000815290516001600160a01b039081166004830152915160248201526000917f00000000000000000000000000000000000000000000000000000000000000001690634d61677190604401602060405180830381865afa15801561161b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8891906158c9565b6040517f2cbc26bb000000000000000000000000000000000000000000000000000000008152608082901b77ffffffffffffffff000000000000000000000000000000001660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632cbc26bb90602401602060405180830381865afa1580156116da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116fe91906158c9565b156111f6576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8216600482015260240161089f565b67ffffffffffffffff81166000908152600760205260408120805460ff16610e88576040517fed053c5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8416600482015260240161089f565b60ff878116600090815260026020908152604080832081516080810183528154815260019091015480861693820193909352610100830485169181019190915262010000909104909216151560608301528735906118008760a46158e6565b90508260600151156118485784516118199060206156dc565b86516118269060206156dc565b6118319060a06158e6565b61183b91906158e6565b61184590826158e6565b90505b36811461188a576040517f8e1192e10000000000000000000000000000000000000000000000000000000081526004810182905236602482015260440161089f565b50815181146118d25781516040517f93df584c00000000000000000000000000000000000000000000000000000000815260048101919091526024810182905260440161089f565b6118da611c24565b60ff808a166000908152600360209081526040808320338452825280832081518083019092528054808616835293949193909284019161010090910416600281111561192857611928614c12565b600281111561193957611939614c12565b905250905060028160200151600281111561195657611956614c12565b1480156119aa5750600260008b60ff1660ff168152602001908152602001600020600301816000015160ff168154811061199257611992615461565b6000918252602090912001546001600160a01b031633145b6119e0576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50816060015115611ac25760208201516119fb9060016158f9565b60ff16855114611a37576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8351855114611a72576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008787604051611a84929190615912565b604051908190038120611a9b918b90602001615922565b604051602081830303815290604052805190602001209050611ac08a8288888861262a565b505b6040805182815260208a81013567ffffffffffffffff169082015260ff8b16917f198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef0910160405180910390a2505050505050505050565b8151600003611b52576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160408051600080825260208201909252911591905b8451811015611bc157611bb9858281518110611b8757611b87615461565b602002602001015184611bb357858381518110611ba657611ba6615461565b6020026020010151612841565b83612841565b600101611b69565b5050505050565b6000546001600160a01b03163314611c225760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161089f565b565b467f000000000000000000000000000000000000000000000000000000000000000014611c22576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015246602482015260440161089f565b67ffffffffffffffff82166000908152600a6020526040812081611cca608085615936565b67ffffffffffffffff1681526020810191909152604001600020549392505050565b67ffffffffffffffff80831660009081526009602090815260408083206001600160a01b03861684529091528120549091829116808203611def5767ffffffffffffffff8516600090815260076020526040902054690100000000000000000090046001600160a01b03168015611ded576040517f856c82470000000000000000000000000000000000000000000000000000000081526001600160a01b03868116600483015282169063856c824790602401602060405180830381865afa158015611dbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611de0919061595d565b6001935093505050611df6565b505b9150600090505b9250929050565b60005b815181101561157a576000828281518110611e1d57611e1d615461565b602002602001015190506000816000015190508067ffffffffffffffff16600003611e74576040517fc656089500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff811660009081526007602052604090206001810154611fe457606083015180611ed2576040517f30a68d4d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008181526008602052604090205460ff1615611f1e576040517f7ca07c380000000000000000000000000000000000000000000000000000000081526004810182905260240161089f565b60008181526008602052604090819020805460ff19166001908117909155830182905584810151835468ffffffffffffffff00196001600160a01b03909216690100000000000000000002919091167fffffff00000000000000000000000000000000000000000000000000000000ff90911617610100178355517ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb990611fd690859067ffffffffffffffff91909116815260200190565b60405180910390a15061205d565b8260600151816001015414158061201a575060408301518154690100000000000000000090046001600160a01b03908116911614155b1561205d576040517fc39a620500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8316600482015260240161089f565b6020830151815490151560ff1990911617815560405167ffffffffffffffff8316907fb8f0c74385134334c728fcac437ec6c6397c9f2c1440532b0c44175a090b1401906120e7908490815460ff811615158252600881901c67ffffffffffffffff16602083015260481c6001600160a01b03166040820152600190910154606082015260800190565b60405180910390a2505050806001019050611e00565b336001600160a01b038216036121555760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161089f565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8560005b8751811015612255576122308882815181106121d5576121d5615461565b6020026020010151602001518888888886815181106121f6576121f6615461565b6020026020010151806020019051810190612211919061597a565b88878151811061222357612223615461565b60200260200101516130e1565b82828151811061224257612242615461565b60209081029190910101526001016121b7565b509695505050505050565b6040805160a08101825260008082526020820152606091810182905281810182905260808101919091526040518060a001604052808461018001518152602001846000015167ffffffffffffffff16815260200184602001516040516020016122d891906001600160a01b0391909116815260200190565b6040516020818303038152906040528152602001846101200151815260200183815250905092915050565b600061230e8361345a565b8015610e855750610e8583836134a6565b806040015160ff1660000361234a576000604051631b3fab5160e11b815260040161089f9190615a2f565b60208082015160ff8082166000908152600290935260408320600181015492939092839216900361239b576060840151600182018054911515620100000262ff0000199092169190911790556123f0565b6060840151600182015460ff62010000909104161515901515146123f0576040517f87f6037c00000000000000000000000000000000000000000000000000000000815260ff8416600482015260240161089f565b60a08401518051601f60ff8216111561241f576001604051631b3fab5160e11b815260040161089f9190615a2f565b612485858560030180548060200260200160405190810160405280929190818152602001828054801561247b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161245d575b5050505050613548565b856060015115612597576124f3858560020180548060200260200160405190810160405280929190818152602001828054801561247b576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161245d575050505050613548565b6080860151805161250d906002870190602084019061408d565b50805160018501805461ff00191661010060ff841690810291909117909155601f1015612550576002604051631b3fab5160e11b815260040161089f9190615a2f565b6040880151612560906003615a49565b60ff168160ff1611612588576003604051631b3fab5160e11b815260040161089f9190615a2f565b612594878360016135b1565b50505b6125a3858360026135b1565b81516125b8906003860190602085019061408d565b5060408681015160018501805460ff191660ff8316179055875180865560a089015192517fab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f54793612611938a939260028b01929190615a65565b60405180910390a161262285613731565b505050505050565b6126326140fb565b835160005b8181101561283757600060018886846020811061265657612656615461565b61266391901a601b6158f9565b89858151811061267557612675615461565b602002602001015189868151811061268f5761268f615461565b6020026020010151604051600081526020016040526040516126cd949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156126ef573d6000803e3d6000fd5b505060408051601f1981015160ff808e166000908152600360209081528582206001600160a01b0385168352815285822085870190965285548084168652939750909550929392840191610100900416600281111561275057612750614c12565b600281111561276157612761614c12565b905250905060018160200151600281111561277e5761277e614c12565b146127b5576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051859060ff16601f81106127cc576127cc615461565b602002015115612808576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600185826000015160ff16601f811061282357612823615461565b911515602090920201525050600101612637565b5050505050505050565b815161284c8161163f565b600061285782611741565b602085015151909150600081900361289a576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84604001515181146128d8576040517f57e0e08300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff8111156128f3576128f3614331565b60405190808252806020026020018201604052801561291c578160200160208202803683370190505b50905060005b828110156129e15760008760200151828151811061294257612942615461565b6020026020010151905061295a81866001015461374d565b83838151811061296c5761296c615461565b60200260200101818152505080610180015183838151811061299057612990615461565b6020026020010151146129d8578061018001516040517f345039be00000000000000000000000000000000000000000000000000000000815260040161089f91815260200190565b50600101612922565b5060006129f8858389606001518a608001516138a9565b905080600003612a40576040517f7dd17a7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8616600482015260240161089f565b8551151560005b84811015610a6257600089602001518281518110612a6757612a67615461565b602002602001015190506000612a81898360600151610e38565b90506002816003811115612a9757612a97614c12565b03612aed5760608201516040805167ffffffffffffffff808d16825290921660208301527f3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c910160405180910390a150506130d9565b6000816003811115612b0157612b01614c12565b1480612b1e57506003816003811115612b1c57612b1c614c12565b145b612b6e5760608201516040517f25507e7f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808c166004830152909116602482015260440161089f565b8315612c4f5760045460009074010000000000000000000000000000000000000000900463ffffffff16612ba2874261568c565b1190508080612bc257506003826003811115612bc057612bc0614c12565b145b612c04576040517fa9cfc86200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b16600482015260240161089f565b8a8481518110612c1657612c16615461565b6020026020010151600014612c49578a8481518110612c3757612c37615461565b60200260200101518360800181815250505b50612cb4565b6000816003811115612c6357612c63614c12565b14612cb45760608201516040517f3ef2a99c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808c166004830152909116602482015260440161089f565b600080612cc58b8560200151611cec565b915091508015612db45760c084015167ffffffffffffffff16612ce9836001615592565b67ffffffffffffffff1614612d6c5760c084015160208501516040517f5444a3301c7c42dd164cbf6ba4b72bf02504f86c049b06a27fc2b662e334bdbd92612d5b928f9267ffffffffffffffff93841681529190921660208201526001600160a01b0391909116604082015260600190565b60405180910390a1505050506130d9565b67ffffffffffffffff8b81166000908152600960209081526040808320888301516001600160a01b031684529091529020805467ffffffffffffffff19169184169190911790555b6000836003811115612dc857612dc8614c12565b03612e595760c084015167ffffffffffffffff16612de7836001615592565b67ffffffffffffffff1614612e595760c084015160208501516040517f852dc8e405695593e311bd83991cf39b14a328f304935eac6d3d55617f911d8992612d5b928f9267ffffffffffffffff93841681529190921660208201526001600160a01b0391909116604082015260600190565b60008d604001518681518110612e7157612e71615461565b6020026020010151905080518561014001515114612ed55760608501516040517f1cfe6d8b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808f166004830152909116602482015260440161089f565b612ee58c866060015160016138ff565b600080612ef287846139a7565b91509150612f058e8860600151846138ff565b888015612f2357506003826003811115612f2157612f21614c12565b145b8015612f4157506000866003811115612f3e57612f3e614c12565b14155b15612f8157866101800151816040517f2b11b8d900000000000000000000000000000000000000000000000000000000815260040161089f929190615aeb565b6003826003811115612f9557612f95614c12565b14158015612fb557506002826003811115612fb257612fb2614c12565b14155b15612ff6578d8760600151836040517f926c5a3e00000000000000000000000000000000000000000000000000000000815260040161089f93929190615b04565b600086600381111561300a5761300a614c12565b036130785767ffffffffffffffff808f1660009081526009602090815260408083208b8301516001600160a01b0316845290915281208054909216919061305083615b2a565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b866101800151876060015167ffffffffffffffff168f67ffffffffffffffff167f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df285856040516130c9929190615b51565b60405180910390a4505050505050505b600101612a47565b604080518082019091526000808252602082015260006131048460200151613bf1565b6040517fbbe4f6db0000000000000000000000000000000000000000000000000000000081526001600160a01b0380831660048301529192506000917f0000000000000000000000000000000000000000000000000000000000000000169063bbe4f6db90602401602060405180830381865afa158015613189573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131ad9190615b71565b90506001600160a01b03811615806131f557506131f36001600160a01b0382167faff2afbf00000000000000000000000000000000000000000000000000000000612303565b155b15613237576040517fae9b4ce90000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161089f565b600080613302633907753760e01b6040518061010001604052808d81526020018b67ffffffffffffffff1681526020018c6001600160a01b031681526020018e8152602001876001600160a01b031681526020018a6000015181526020018a604001518152602001898152506040516024016132b39190615b8e565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152600454859063ffffffff600160e01b909104166113886084613c33565b509150915081613327578060405163e1cd550960e01b815260040161089f919061417f565b805160201461336f5780516040517f78ef802400000000000000000000000000000000000000000000000000000000815260206004820152602481019190915260440161089f565b6000818060200190518101906133859190615c65565b604080516001600160a01b038d16602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b03167fa9059cbb00000000000000000000000000000000000000000000000000000000179052600454919250613408918790600160c01b900463ffffffff166113886084613c33565b5090935091508261342e578160405163e1cd550960e01b815260040161089f919061417f565b604080518082019091526001600160a01b03909516855260208501525091925050509695505050505050565b6000613486827f01ffc9a7000000000000000000000000000000000000000000000000000000006134a6565b8015610e88575061349f826001600160e01b03196134a6565b1592915050565b604080516001600160e01b03198316602480830191909152825180830390910181526044909101909152602080820180516001600160e01b03167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015613531575060208210155b801561353d5750600081115b979650505050505050565b60005b8151811015610e335760ff83166000908152600360205260408120835190919084908490811061357d5761357d615461565b6020908102919091018101516001600160a01b03168252810191909152604001600020805461ffff1916905560010161354b565b60005b82518160ff161015610ad0576000838260ff16815181106135d7576135d7615461565b60200260200101519050600060028111156135f4576135f4614c12565b60ff80871660009081526003602090815260408083206001600160a01b0387168452909152902054610100900416600281111561363357613633614c12565b14613654576004604051631b3fab5160e11b815260040161089f9190615a2f565b6001600160a01b038116613694576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405180604001604052808360ff1681526020018460028111156136ba576136ba614c12565b905260ff80871660009081526003602090815260408083206001600160a01b0387168452825290912083518154931660ff198416811782559184015190929091839161ffff19161761010083600281111561371757613717614c12565b0217905550905050508061372a90615c7e565b90506135b4565b60ff81166111f657600c805467ffffffffffffffff1916905550565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b61010001516040516020016137e39897969594939291906001600160a01b039889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b604051602081830303815290604052805190602001208561012001518051906020012086610140015160405160200161381c9190615c9d565b604051602081830303815290604052805190602001208761016001516040516020016138489190615d52565b60408051601f198184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e0015b60405160208183030381529060405280519060200120905092915050565b6000806138b7858585613d59565b90506138c28161157e565b6138d05760009150506138f7565b67ffffffffffffffff86166000908152600b60209081526040808320938352929052205490505b949350505050565b6000600261390e6080856156b5565b67ffffffffffffffff1661392291906156dc565b905060006139308585611ca5565b90508161393f6001600461568c565b901b19168183600381111561395657613956614c12565b67ffffffffffffffff87166000908152600a602052604081209190921b92909217918291613985608088615936565b67ffffffffffffffff1681526020810191909152604001600020555050505050565b6040517ff52121a5000000000000000000000000000000000000000000000000000000008152600090606090309063f52121a5906139eb9087908790600401615d65565b600060405180830381600087803b158015613a0557600080fd5b505af1925050508015613a16575060015b613bd6573d808015613a44576040519150601f19603f3d011682016040523d82523d6000602084013e613a49565b606091505b506000613a5582615eaa565b90507f0a8d6e8c000000000000000000000000000000000000000000000000000000006001600160e01b031982161480613a9f575063e1cd550960e01b6001600160e01b03198216145b80613aba575063046b337b60e51b6001600160e01b03198216145b80613aee57507f78ef8024000000000000000000000000000000000000000000000000000000006001600160e01b03198216145b80613b2257507f0c3b563c000000000000000000000000000000000000000000000000000000006001600160e01b03198216145b80613b5657507fae9b4ce9000000000000000000000000000000000000000000000000000000006001600160e01b03198216145b80613b8a57507f09c25325000000000000000000000000000000000000000000000000000000006001600160e01b03198216145b15613b9b5750600392509050611df6565b856101800151826040517f2b11b8d900000000000000000000000000000000000000000000000000000000815260040161089f929190615aeb565b50506040805160208101909152600081526002909250929050565b60008151602014613c17578160405163046b337b60e51b815260040161089f919061417f565b610e8882806020019051810190613c2e9190615c65565b613ff8565b6000606060008361ffff1667ffffffffffffffff811115613c5657613c56614331565b6040519080825280601f01601f191660200182016040528015613c80576020820181803683370190505b509150863b613cb3577f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b5a85811015613ce6577fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b8590036040810481038710613d1f577f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b505a6000808a5160208c0160008c8cf193505a900390503d84811115613d425750835b808352806000602085013e50955095509592505050565b8251825160009190818303613d9a576040517f11a6b26400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6101018211801590613dae57506101018111155b613dcb576040516309bde33960e01b815260040160405180910390fd5b60001982820101610100811115613df5576040516309bde33960e01b815260040160405180910390fd5b80600003613e225786600081518110613e1057613e10615461565b60200260200101519350505050613ff1565b60008167ffffffffffffffff811115613e3d57613e3d614331565b604051908082528060200260200182016040528015613e66578160200160208202803683370190505b50905060008080805b85811015613f905760006001821b8b811603613eca5788851015613eb3578c5160018601958e918110613ea457613ea4615461565b60200260200101519050613eec565b8551600185019487918110613ea457613ea4615461565b8b5160018401938d918110613ee157613ee1615461565b602002602001015190505b600089861015613f1c578d5160018701968f918110613f0d57613f0d615461565b60200260200101519050613f3e565b8651600186019588918110613f3357613f33615461565b602002602001015190505b82851115613f5f576040516309bde33960e01b815260040160405180910390fd5b613f69828261404c565b878481518110613f7b57613f7b615461565b60209081029190910101525050600101613e6f565b506001850382148015613fa257508683145b8015613fad57508581145b613fca576040516309bde33960e01b815260040160405180910390fd5b836001860381518110613fdf57613fdf615461565b60200260200101519750505050505050505b9392505050565b60006001600160a01b03821180614010575061040082105b156140485760408051602081018490520160408051601f198184030181529082905263046b337b60e51b825261089f9160040161417f565b5090565b60008183106140645761405f828461406a565b610e85565b610e8583835b60408051600160208201529081018390526060810182905260009060800161388b565b8280548282559060005260206000209081019282156140ef579160200282015b828111156140ef578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039091161782556020909201916001909101906140ad565b5061404892915061411a565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115614048576000815560010161411b565b60005b8381101561414a578181015183820152602001614132565b50506000910152565b6000815180845261416b81602086016020860161412f565b601f01601f19169290920160200192915050565b602081526000610e856020830184614153565b8060608101831015610e8857600080fd5b60008083601f8401126141b557600080fd5b50813567ffffffffffffffff8111156141cd57600080fd5b602083019150836020828501011115611df657600080fd5b60008083601f8401126141f757600080fd5b50813567ffffffffffffffff81111561420f57600080fd5b6020830191508360208260051b8501011115611df657600080fd5b60008060008060008060008060e0898b03121561424657600080fd5b6142508a8a614192565b9750606089013567ffffffffffffffff8082111561426d57600080fd5b6142798c838d016141a3565b909950975060808b013591508082111561429257600080fd5b61429e8c838d016141e5565b909750955060a08b01359150808211156142b757600080fd5b506142c48b828c016141e5565b999c989b50969995989497949560c00135949350505050565b6000806000608084860312156142f257600080fd5b6142fc8585614192565b9250606084013567ffffffffffffffff81111561431857600080fd5b614324868287016141a3565b9497909650939450505050565b634e487b7160e01b600052604160045260246000fd5b60405160c0810167ffffffffffffffff8111828210171561436a5761436a614331565b60405290565b6040805190810167ffffffffffffffff8111828210171561436a5761436a614331565b6040516101a0810167ffffffffffffffff8111828210171561436a5761436a614331565b60405160a0810167ffffffffffffffff8111828210171561436a5761436a614331565b6040516080810167ffffffffffffffff8111828210171561436a5761436a614331565b6040516060810167ffffffffffffffff8111828210171561436a5761436a614331565b604051601f8201601f1916810167ffffffffffffffff8111828210171561444957614449614331565b604052919050565b6001600160a01b03811681146111f657600080fd5b803561447181614451565b919050565b803563ffffffff8116811461447157600080fd5b600060c0828403121561449c57600080fd5b6144a4614347565b82356144af81614451565b81526144bd60208401614476565b60208201526144ce60408401614476565b60408201526144df60608401614476565b606082015260808301356144f281614451565b608082015260a083013561450581614451565b60a08201529392505050565b600067ffffffffffffffff82111561452b5761452b614331565b5060051b60200190565b67ffffffffffffffff811681146111f657600080fd5b803561447181614535565b80151581146111f657600080fd5b803561447181614556565b600067ffffffffffffffff82111561458957614589614331565b50601f01601f191660200190565b600082601f8301126145a857600080fd5b81356145bb6145b68261456f565b614420565b8181528460208386010111156145d057600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126145fe57600080fd5b8135602061460e6145b683614511565b82815260069290921b8401810191818101908684111561462d57600080fd5b8286015b84811015612255576040818903121561464a5760008081fd5b614652614370565b813561465d81614451565b81528185013585820152835291830191604001614631565b600082601f83011261468657600080fd5b813560206146966145b683614511565b82815260059290921b840181019181810190868411156146b557600080fd5b8286015b8481101561225557803567ffffffffffffffff8111156146d95760008081fd5b6146e78986838b0101614597565b8452509183019183016146b9565b60006101a0828403121561470857600080fd5b614710614393565b905061471b8261454b565b815261472960208301614466565b602082015261473a60408301614466565b604082015261474b6060830161454b565b60608201526080820135608082015261476660a08301614564565b60a082015261477760c0830161454b565b60c082015261478860e08301614466565b60e082015261010082810135908201526101208083013567ffffffffffffffff808211156147b557600080fd5b6147c186838701614597565b838501526101409250828501359150808211156147dd57600080fd5b6147e9868387016145ed565b8385015261016092508285013591508082111561480557600080fd5b5061481285828601614675565b82840152505061018080830135818301525092915050565b600082601f83011261483b57600080fd5b8135602061484b6145b683614511565b82815260059290921b8401810191818101908684111561486a57600080fd5b8286015b8481101561225557803567ffffffffffffffff81111561488e5760008081fd5b61489c8986838b01016146f5565b84525091830191830161486e565b600082601f8301126148bb57600080fd5b813560206148cb6145b683614511565b82815260059290921b840181019181810190868411156148ea57600080fd5b8286015b8481101561225557803567ffffffffffffffff81111561490e5760008081fd5b61491c8986838b0101614675565b8452509183019183016148ee565b600082601f83011261493b57600080fd5b8135602061494b6145b683614511565b8083825260208201915060208460051b87010193508684111561496d57600080fd5b602086015b848110156122555780358352918301918301614972565b600082601f83011261499a57600080fd5b813560206149aa6145b683614511565b82815260059290921b840181019181810190868411156149c957600080fd5b8286015b8481101561225557803567ffffffffffffffff808211156149ee5760008081fd5b9088019060a0828b03601f1901811315614a085760008081fd5b614a106143b7565b614a1b88850161454b565b815260408085013584811115614a315760008081fd5b614a3f8e8b8389010161482a565b8a8401525060608086013585811115614a585760008081fd5b614a668f8c838a01016148aa565b8385015250608091508186013585811115614a815760008081fd5b614a8f8f8c838a010161492a565b91840191909152509190930135908301525083529183019183016149cd565b6000806040808486031215614ac257600080fd5b833567ffffffffffffffff80821115614ada57600080fd5b614ae687838801614989565b9450602091508186013581811115614afd57600080fd5b8601601f81018813614b0e57600080fd5b8035614b1c6145b682614511565b81815260059190911b8201840190848101908a831115614b3b57600080fd5b8584015b83811015614bc757803586811115614b575760008081fd5b8501603f81018d13614b695760008081fd5b87810135614b796145b682614511565b81815260059190911b82018a0190898101908f831115614b995760008081fd5b928b01925b82841015614bb75783358252928a0192908a0190614b9e565b8652505050918601918601614b3f565b50809750505050505050509250929050565b60008060408385031215614bec57600080fd5b8235614bf781614535565b91506020830135614c0781614535565b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b60048110614c3857614c38614c12565b9052565b60208101610e888284614c28565b600060208284031215614c5c57600080fd5b8135613ff181614535565b600060208284031215614c7957600080fd5b813567ffffffffffffffff811115614c9057600080fd5b820160a08185031215613ff157600080fd5b60008060408385031215614cb557600080fd5b8235614cc081614535565b91506020830135614c0781614451565b803560ff8116811461447157600080fd5b600060208284031215614cf357600080fd5b610e8582614cd0565b60008151808452602080850194506020840160005b83811015614d365781516001600160a01b031687529582019590820190600101614d11565b509495945050505050565b60208152600082518051602084015260ff602082015116604084015260ff604082015116606084015260608101511515608084015250602083015160c060a0840152614d9060e0840182614cfc565b90506040840151601f198483030160c0850152614dad8282614cfc565b95945050505050565b60008060408385031215614dc957600080fd5b8235614dd481614535565b946020939093013593505050565b60008060208385031215614df557600080fd5b823567ffffffffffffffff80821115614e0d57600080fd5b818501915085601f830112614e2157600080fd5b813581811115614e3057600080fd5b8660208260061b8501011115614e4557600080fd5b60209290920196919550909350505050565b60006020808385031215614e6a57600080fd5b823567ffffffffffffffff811115614e8157600080fd5b8301601f81018513614e9257600080fd5b8035614ea06145b682614511565b81815260079190911b82018301908381019087831115614ebf57600080fd5b928401925b8284101561353d5760808489031215614edd5760008081fd5b614ee56143da565b8435614ef081614535565b815284860135614eff81614556565b81870152604085810135614f1281614451565b9082015260608581013590820152825260809093019290840190614ec4565b600060208284031215614f4357600080fd5b8135613ff181614451565b60008060408385031215614f6157600080fd5b823567ffffffffffffffff80821115614f7957600080fd5b614f85868387016146f5565b93506020850135915080821115614f9b57600080fd5b50614fa885828601614675565b9150509250929050565b600082601f830112614fc357600080fd5b81356020614fd36145b683614511565b8083825260208201915060208460051b870101935086841115614ff557600080fd5b602086015b8481101561225557803561500d81614451565b8352918301918301614ffa565b6000602080838503121561502d57600080fd5b823567ffffffffffffffff8082111561504557600080fd5b818501915085601f83011261505957600080fd5b81356150676145b682614511565b81815260059190911b8301840190848101908883111561508657600080fd5b8585015b83811015615158578035858111156150a157600080fd5b860160c0818c03601f190112156150b85760008081fd5b6150c0614347565b88820135815260406150d3818401614cd0565b8a83015260606150e4818501614cd0565b82840152608091506150f7828501614564565b9083015260a0838101358981111561510f5760008081fd5b61511d8f8d83880101614fb2565b838501525060c08401359150888211156151375760008081fd5b6151458e8c84870101614fb2565b908301525084525091860191860161508a565b5098975050505050505050565b60006020828403121561517757600080fd5b5035919050565b80356001600160e01b038116811461447157600080fd5b600082601f8301126151a657600080fd5b813560206151b66145b683614511565b82815260069290921b840181019181810190868411156151d557600080fd5b8286015b8481101561225557604081890312156151f25760008081fd5b6151fa614370565b813561520581614535565b815261521282860161517e565b818601528352918301916040016151d9565b600082601f83011261523557600080fd5b813560206152456145b683614511565b82815260079290921b8401810191818101908684111561526457600080fd5b8286015b848110156122555780880360808112156152825760008081fd5b61528a6143fd565b823561529581614535565b81526040601f1983018113156152ab5760008081fd5b6152b3614370565b9250868401356152c281614535565b8352838101356152d181614535565b8388015281870192909252606083013591810191909152835291830191608001615268565b6000602080838503121561530957600080fd5b823567ffffffffffffffff8082111561532157600080fd5b8185019150604080838803121561533757600080fd5b61533f614370565b83358381111561534e57600080fd5b84016040818a03121561536057600080fd5b615368614370565b81358581111561537757600080fd5b8201601f81018b1361538857600080fd5b80356153966145b682614511565b81815260069190911b8201890190898101908d8311156153b557600080fd5b928a01925b828410156154055787848f0312156153d25760008081fd5b6153da614370565b84356153e581614451565b81526153f2858d0161517e565b818d0152825292870192908a01906153ba565b84525050508187013593508484111561541d57600080fd5b6154298a858401615195565b818801528252508385013591508282111561544357600080fd5b61544f88838601615224565b85820152809550505050505092915050565b634e487b7160e01b600052603260045260246000fd5b805160408084528151848201819052600092602091908201906060870190855b818110156154ce57835180516001600160a01b031684528501516001600160e01b0316858401529284019291850191600101615497565b50508583015187820388850152805180835290840192506000918401905b80831015615528578351805167ffffffffffffffff1683528501516001600160e01b0316858301529284019260019290920191908501906154ec565b50979650505050505050565b602081526000610e856020830184615477565b67ffffffffffffffff8316815260608101613ff16020830184805167ffffffffffffffff908116835260209182015116910152565b634e487b7160e01b600052601160045260246000fd5b67ffffffffffffffff8181168382160190808211156155b3576155b361557c565b5092915050565b6000602080835260608451604080848701526155d96060870183615477565b87850151878203601f19016040890152805180835290860193506000918601905b8083101561515857845167ffffffffffffffff81511683528781015161563989850182805167ffffffffffffffff908116835260209182015116910152565b508401518287015293860193600192909201916080909101906155fa565b60006020828403121561566957600080fd5b813567ffffffffffffffff81111561568057600080fd5b6138f784828501614989565b81810381811115610e8857610e8861557c565b634e487b7160e01b600052601260045260246000fd5b600067ffffffffffffffff808416806156d0576156d061569f565b92169190910692915050565b8082028115828204841417610e8857610e8861557c565b60006040828403121561570557600080fd5b61570d614370565b823561571881614535565b81526020928301359281019290925250919050565b60008151808452602080850194506020840160005b83811015614d3657815180516001600160a01b031688526020908101519088015260408701965090820190600101615742565b8051825267ffffffffffffffff60208201511660208301526000604082015160a060408501526157a860a0850182614153565b9050606083015184820360608601526157c18282614153565b91505060808301518482036080860152614dad828261572d565b602081526000610e856020830184615775565b6080815260006158016080830187615775565b61ffff9590951660208301525060408101929092526001600160a01b0316606090910152919050565b600082601f83011261583b57600080fd5b81516158496145b68261456f565b81815284602083860101111561585e57600080fd5b6138f782602083016020870161412f565b60008060006060848603121561588457600080fd5b835161588f81614556565b602085015190935067ffffffffffffffff8111156158ac57600080fd5b6158b88682870161582a565b925050604084015190509250925092565b6000602082840312156158db57600080fd5b8151613ff181614556565b80820180821115610e8857610e8861557c565b60ff8181168382160190811115610e8857610e8861557c565b8183823760009101908152919050565b828152606082602083013760800192915050565b600067ffffffffffffffff808416806159515761595161569f565b92169190910492915050565b60006020828403121561596f57600080fd5b8151613ff181614535565b60006020828403121561598c57600080fd5b815167ffffffffffffffff808211156159a457600080fd5b90830190606082860312156159b857600080fd5b6159c06143fd565b8251828111156159cf57600080fd5b6159db8782860161582a565b8252506020830151828111156159f057600080fd5b6159fc8782860161582a565b602083015250604083015182811115615a1457600080fd5b615a208782860161582a565b60408301525095945050505050565b6020810160058310615a4357615a43614c12565b91905290565b60ff81811683821602908116908181146155b3576155b361557c565b600060a0820160ff88168352602087602085015260a0604085015281875480845260c086019150886000526020600020935060005b81811015615abf5784546001600160a01b031683526001948501949284019201615a9a565b50508481036060860152615ad38188614cfc565b935050505060ff831660808301529695505050505050565b8281526040602082015260006138f76040830184614153565b67ffffffffffffffff848116825283166020820152606081016138f76040830184614c28565b600067ffffffffffffffff808316818103615b4757615b4761557c565b6001019392505050565b615b5b8184614c28565b6040602082015260006138f76040830184614153565b600060208284031215615b8357600080fd5b8151613ff181614451565b6020815260008251610100806020850152615bad610120850183614153565b91506020850151615bca604086018267ffffffffffffffff169052565b5060408501516001600160a01b038116606086015250606085015160808501526080850151615c0460a08601826001600160a01b03169052565b5060a0850151601f19808685030160c0870152615c218483614153565b935060c08701519150808685030160e0870152615c3e8483614153565b935060e0870151915080868503018387015250615c5b8382614153565b9695505050505050565b600060208284031215615c7757600080fd5b5051919050565b600060ff821660ff8103615c9457615c9461557c565b60010192915050565b6020808252825182820181905260009190848201906040850190845b81811015615cec57835180516001600160a01b031684526020908101519084015260408301938501939250600101615cb9565b50909695505050505050565b60008282518085526020808601955060208260051b8401016020860160005b84811015615d4557601f19868403018952615d33838351614153565b98840198925090830190600101615d17565b5090979650505050505050565b602081526000610e856020830184615cf8565b60408152615d8060408201845167ffffffffffffffff169052565b60006020840151615d9c60608401826001600160a01b03169052565b5060408401516001600160a01b038116608084015250606084015167ffffffffffffffff811660a084015250608084015160c083015260a084015180151560e08401525060c0840151610100615dfd8185018367ffffffffffffffff169052565b60e08601519150610120615e1b818601846001600160a01b03169052565b81870151925061014091508282860152808701519250506101a06101608181870152615e4b6101e0870185614153565b9350828801519250603f19610180818887030181890152615e6c868661572d565b9550828a01519450818887030184890152615e878686615cf8565b9550808a01516101c089015250505050508281036020840152614dad8185615cf8565b6000815160208301516001600160e01b031980821693506004831015615eda5780818460040360031b1b83161693505b50505091905056fea164736f6c6343000818000a",
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCaller) GetLatestPriceSequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _EVM2EVMMultiOffRamp.contract.Call(opts, &out, "getLatestPriceSequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) GetLatestPriceSequenceNumber() (uint64, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetLatestPriceSequenceNumber(&_EVM2EVMMultiOffRamp.CallOpts)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampCallerSession) GetLatestPriceSequenceNumber() (uint64, error) {
	return _EVM2EVMMultiOffRamp.Contract.GetLatestPriceSequenceNumber(&_EVM2EVMMultiOffRamp.CallOpts)
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) SetLatestPriceSequenceNumber(opts *bind.TransactOpts, latestPriceSequenceNumber uint64) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "setLatestPriceSequenceNumber", latestPriceSequenceNumber)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) SetLatestPriceSequenceNumber(latestPriceSequenceNumber uint64) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.SetLatestPriceSequenceNumber(&_EVM2EVMMultiOffRamp.TransactOpts, latestPriceSequenceNumber)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) SetLatestPriceSequenceNumber(latestPriceSequenceNumber uint64) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.SetLatestPriceSequenceNumber(&_EVM2EVMMultiOffRamp.TransactOpts, latestPriceSequenceNumber)
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
	OcrPluginType uint8
	ConfigDigest  [32]byte
	Signers       []common.Address
	Transmitters  []common.Address
	F             uint8
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

type EVM2EVMMultiOffRampDynamicConfigSetIterator struct {
	Event *EVM2EVMMultiOffRampDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampDynamicConfigSet)
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
		it.Event = new(EVM2EVMMultiOffRampDynamicConfigSet)
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

func (it *EVM2EVMMultiOffRampDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampDynamicConfigSet struct {
	DynamicConfig EVM2EVMMultiOffRampDynamicConfig
	Raw           types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampDynamicConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampDynamicConfigSetIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampDynamicConfigSet)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseDynamicConfigSet(log types.Log) (*EVM2EVMMultiOffRampDynamicConfigSet, error) {
	event := new(EVM2EVMMultiOffRampDynamicConfigSet)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

type EVM2EVMMultiOffRampLatestPriceSequenceNumberSetIterator struct {
	Event *EVM2EVMMultiOffRampLatestPriceSequenceNumberSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampLatestPriceSequenceNumberSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampLatestPriceSequenceNumberSet)
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
		it.Event = new(EVM2EVMMultiOffRampLatestPriceSequenceNumberSet)
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

func (it *EVM2EVMMultiOffRampLatestPriceSequenceNumberSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampLatestPriceSequenceNumberSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampLatestPriceSequenceNumberSet struct {
	OldSequenceNumber uint64
	NewSequenceNumber uint64
	Raw               types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterLatestPriceSequenceNumberSet(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampLatestPriceSequenceNumberSetIterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "LatestPriceSequenceNumberSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampLatestPriceSequenceNumberSetIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "LatestPriceSequenceNumberSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchLatestPriceSequenceNumberSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampLatestPriceSequenceNumberSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "LatestPriceSequenceNumberSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampLatestPriceSequenceNumberSet)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "LatestPriceSequenceNumberSet", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseLatestPriceSequenceNumberSet(log types.Log) (*EVM2EVMMultiOffRampLatestPriceSequenceNumberSet, error) {
	event := new(EVM2EVMMultiOffRampLatestPriceSequenceNumberSet)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "LatestPriceSequenceNumberSet", log); err != nil {
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

type EVM2EVMMultiOffRampStaticConfigSetIterator struct {
	Event *EVM2EVMMultiOffRampStaticConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMMultiOffRampStaticConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMMultiOffRampStaticConfigSet)
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
		it.Event = new(EVM2EVMMultiOffRampStaticConfigSet)
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

func (it *EVM2EVMMultiOffRampStaticConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMMultiOffRampStaticConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMMultiOffRampStaticConfigSet struct {
	StaticConfig EVM2EVMMultiOffRampStaticConfig
	Raw          types.Log
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) FilterStaticConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampStaticConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.FilterLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMMultiOffRampStaticConfigSetIterator{contract: _EVM2EVMMultiOffRamp.contract, event: "StaticConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampStaticConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMMultiOffRamp.contract.WatchLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMMultiOffRampStaticConfigSet)
				if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampFilterer) ParseStaticConfigSet(log types.Log) (*EVM2EVMMultiOffRampStaticConfigSet, error) {
	event := new(EVM2EVMMultiOffRampStaticConfigSet)
	if err := _EVM2EVMMultiOffRamp.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMMultiOffRamp.abi.Events["CommitReportAccepted"].ID:
		return _EVM2EVMMultiOffRamp.ParseCommitReportAccepted(log)
	case _EVM2EVMMultiOffRamp.abi.Events["ConfigSet"].ID:
		return _EVM2EVMMultiOffRamp.ParseConfigSet(log)
	case _EVM2EVMMultiOffRamp.abi.Events["DynamicConfigSet"].ID:
		return _EVM2EVMMultiOffRamp.ParseDynamicConfigSet(log)
	case _EVM2EVMMultiOffRamp.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMMultiOffRamp.ParseExecutionStateChanged(log)
	case _EVM2EVMMultiOffRamp.abi.Events["LatestPriceSequenceNumberSet"].ID:
		return _EVM2EVMMultiOffRamp.ParseLatestPriceSequenceNumberSet(log)
	case _EVM2EVMMultiOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMMultiOffRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMMultiOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMMultiOffRamp.ParseOwnershipTransferred(log)
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
	case _EVM2EVMMultiOffRamp.abi.Events["StaticConfigSet"].ID:
		return _EVM2EVMMultiOffRamp.ParseStaticConfigSet(log)
	case _EVM2EVMMultiOffRamp.abi.Events["Transmitted"].ID:
		return _EVM2EVMMultiOffRamp.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMMultiOffRampCommitReportAccepted) Topic() common.Hash {
	return common.HexToHash("0x3a3950e13dd607cc37980db0ef14266c40d2bba9c01b2e44bfe549808883095d")
}

func (EVM2EVMMultiOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f547")
}

func (EVM2EVMMultiOffRampDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x0da37fd00459f4f5f0b8210d31525e4910ae674b8bab34b561d146bb45773a4c")
}

func (EVM2EVMMultiOffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df2")
}

func (EVM2EVMMultiOffRampLatestPriceSequenceNumberSet) Topic() common.Hash {
	return common.HexToHash("0x88ad9c61d6caf19a2af116a871802a03a31e680115a2dd20e8c08801d7c82f83")
}

func (EVM2EVMMultiOffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMMultiOffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
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
	return common.HexToHash("0xb8f0c74385134334c728fcac437ec6c6397c9f2c1440532b0c44175a090b1401")
}

func (EVM2EVMMultiOffRampSourceChainSelectorAdded) Topic() common.Hash {
	return common.HexToHash("0xf4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb9")
}

func (EVM2EVMMultiOffRampStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0x2f56698ec552a5d53d27d6f4b3dd8b6989f6426b6151a36aff61160c1d07efdf")
}

func (EVM2EVMMultiOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0x198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef0")
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRamp) Address() common.Address {
	return _EVM2EVMMultiOffRamp.address
}

type EVM2EVMMultiOffRampInterface interface {
	CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMMultiOffRampDynamicConfig, error)

	GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error)

	GetLatestPriceSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetMerkleRoot(opts *bind.CallOpts, sourceChainSelector uint64, root [32]byte) (*big.Int, error)

	GetSenderNonce(opts *bind.CallOpts, sourceChainSelector uint64, sender common.Address) (uint64, error)

	GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (EVM2EVMMultiOffRampSourceChainConfig, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMMultiOffRampStaticConfig, error)

	IsBlessed(opts *bind.CallOpts, root [32]byte) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts, ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []EVM2EVMMultiOffRampSourceChainConfigArgs) (*types.Transaction, error)

	Commit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, reports []InternalExecutionReportSingleChain, gasLimitOverrides [][]*big.Int) (*types.Transaction, error)

	ResetUnblessedRoots(opts *bind.TransactOpts, rootToReset []EVM2EVMMultiOffRampUnblessedRoot) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMMultiOffRampDynamicConfig) (*types.Transaction, error)

	SetLatestPriceSequenceNumber(opts *bind.TransactOpts, latestPriceSequenceNumber uint64) (*types.Transaction, error)

	SetOCR3Configs(opts *bind.TransactOpts, ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterCommitReportAccepted(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampCommitReportAcceptedIterator, error)

	WatchCommitReportAccepted(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampCommitReportAccepted) (event.Subscription, error)

	ParseCommitReportAccepted(log types.Log) (*EVM2EVMMultiOffRampCommitReportAccepted, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMMultiOffRampConfigSet, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*EVM2EVMMultiOffRampDynamicConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*EVM2EVMMultiOffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMMultiOffRampExecutionStateChanged, error)

	FilterLatestPriceSequenceNumberSet(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampLatestPriceSequenceNumberSetIterator, error)

	WatchLatestPriceSequenceNumberSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampLatestPriceSequenceNumberSet) (event.Subscription, error)

	ParseLatestPriceSequenceNumberSet(log types.Log) (*EVM2EVMMultiOffRampLatestPriceSequenceNumberSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMMultiOffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMMultiOffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMMultiOffRampOwnershipTransferred, error)

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

	FilterStaticConfigSet(opts *bind.FilterOpts) (*EVM2EVMMultiOffRampStaticConfigSetIterator, error)

	WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampStaticConfigSet) (event.Subscription, error)

	ParseStaticConfigSet(log types.Log) (*EVM2EVMMultiOffRampStaticConfigSet, error)

	FilterTransmitted(opts *bind.FilterOpts, ocrPluginType []uint8) (*EVM2EVMMultiOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMMultiOffRampTransmitted, ocrPluginType []uint8) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMMultiOffRampTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
