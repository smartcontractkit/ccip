// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package offramp

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

type IRMNV2Signature struct {
	R [32]byte
	S [32]byte
}

type InternalAny2EVMRampMessage struct {
	Header       InternalRampMessageHeader
	Sender       []byte
	Data         []byte
	Receiver     common.Address
	GasLimit     *big.Int
	TokenAmounts []InternalRampTokenAmount
}

type InternalExecutionReportSingleChain struct {
	SourceChainSelector uint64
	Messages            []InternalAny2EVMRampMessage
	OffchainTokenData   [][][]byte
	Proofs              [][32]byte
	ProofFlagBits       *big.Int
}

type InternalGasPriceUpdate struct {
	DestChainSelector uint64
	UsdPerUnitGas     *big.Int
}

type InternalMerkleRoot struct {
	SourceChainSelector uint64
	OnRampAddress       []byte
	MinSeqNr            uint64
	MaxSeqNr            uint64
	MerkleRoot          [32]byte
}

type InternalPriceUpdates struct {
	TokenPriceUpdates []InternalTokenPriceUpdate
	GasPriceUpdates   []InternalGasPriceUpdate
}

type InternalRampMessageHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

type InternalRampTokenAmount struct {
	SourcePoolAddress []byte
	DestTokenAddress  []byte
	ExtraData         []byte
	Amount            *big.Int
	DestExecData      []byte
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

type OffRampCommitReport struct {
	PriceUpdates  InternalPriceUpdates
	MerkleRoots   []InternalMerkleRoot
	RmnSignatures []IRMNV2Signature
}

type OffRampDynamicConfig struct {
	FeeQuoter                               common.Address
	PermissionLessExecutionThresholdSeconds uint32
	MaxTokenTransferGas                     uint32
	MaxPoolReleaseOrMintGas                 uint32
	MessageValidator                        common.Address
}

type OffRampSourceChainConfig struct {
	Router    common.Address
	IsEnabled bool
	MinSeqNr  uint64
	OnRamp    []byte
}

type OffRampSourceChainConfigArgs struct {
	Router              common.Address
	SourceChainSelector uint64
	IsEnabled           bool
	OnRamp              []byte
}

type OffRampStaticConfig struct {
	ChainSelector      uint64
	Rmn                common.Address
	TokenAdminRegistry common.Address
	NonceManager       common.Address
}

var OffRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"contractIRMNV2\",\"name\":\"rmn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"}],\"internalType\":\"structOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"feeQuoter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"}],\"internalType\":\"structOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractIRouter\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"internalType\":\"structOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumMultiOCR3Base.InvalidConfigErrorType\",\"name\":\"errorType\",\"type\":\"uint8\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"name\":\"InvalidInterval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"InvalidManualExecutionGasLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"messageDestChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidMessageDestChainSelector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"InvalidNewState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRoot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidStaticConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeavesCannotBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionGasLimitMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"errorReason\",\"type\":\"bytes\"}],\"name\":\"MessageValidationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notPool\",\"type\":\"address\"}],\"name\":\"NotACompatiblePool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReleased\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balancePre\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balancePost\",\"type\":\"uint256\"}],\"name\":\"ReleaseOrMintBalanceMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"RootAlreadyCommitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureVerificationDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"reportSourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"messageSourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainSelectorMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleCommitReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"}],\"name\":\"StaticConfigCannotBeChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"TokenDataMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"TokenHandlingError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedTokenData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroChainSelectorNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint224\",\"name\":\"usdPerToken\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint224\",\"name\":\"usdPerUnitGas\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.GasPriceUpdate[]\",\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.PriceUpdates\",\"name\":\"priceUpdates\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"onRampAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.MerkleRoot[]\",\"name\":\"merkleRoots\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIRMNV2.Signature[]\",\"name\":\"rmnSignatures\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structOffRamp.CommitReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"CommitReportAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeQuoter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"DynamicConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"RootRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"SkippedAlreadyExecutedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SkippedReportExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"contractIRouter\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structOffRamp.SourceChainConfig\",\"name\":\"sourceConfig\",\"type\":\"tuple\"}],\"name\":\"SourceChainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainSelectorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"contractIRMNV2\",\"name\":\"rmn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"}],\"name\":\"StaticConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIRouter\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"internalType\":\"structOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\"}],\"name\":\"applySourceChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destExecData\",\"type\":\"bytes\"}],\"internalType\":\"structInternal.RampTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.Any2EVMRampMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeQuoter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"}],\"internalType\":\"structOffRamp.DynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPriceSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"getSourceChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIRouter\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"internalType\":\"structOffRamp.SourceChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"contractIRMNV2\",\"name\":\"rmn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"}],\"internalType\":\"structOffRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"}],\"name\":\"latestConfigDetails\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"n\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\"}],\"internalType\":\"structMultiOCR3Base.ConfigInfo\",\"name\":\"configInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"}],\"internalType\":\"structMultiOCR3Base.OCRConfig\",\"name\":\"ocrConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"destExecData\",\"type\":\"bytes\"}],\"internalType\":\"structInternal.RampTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.Any2EVMRampMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReportSingleChain[]\",\"name\":\"reports\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"gasLimitOverrides\",\"type\":\"uint256[][]\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeQuoter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"}],\"internalType\":\"structOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"}],\"internalType\":\"structMultiOCR3Base.OCRConfigArgs[]\",\"name\":\"ocrConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"setOCR3Configs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162006b8038038062006b8083398101604081905262000035916200082b565b33806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf81620001fa565b5050466080525060208301516001600160a01b03161580620000ec575060408301516001600160a01b0316155b8062000103575060608301516001600160a01b0316155b1562000122576040516342bcdf7f60e11b815260040160405180910390fd5b82516001600160401b03166000036200014e5760405163c656089560e01b815260040160405180910390fd5b82516001600160401b0390811660a052602080850180516001600160a01b0390811660c05260408088018051831660e0526060808a01805185166101005283518b519098168852945184169587019590955251821690850152905116908201527f683eb52ee924eb817377cfa8f41f238f4bb7a877da5267869dfffbad85f564d89060800160405180910390a1620001e682620002a5565b620001f181620003c1565b50505062000bcb565b336001600160a01b03821603620002545760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0316620002ce576040516342bcdf7f60e11b815260040160405180910390fd5b80516004805460208085018051604080880180516060808b0180516001600160a01b039b8c166001600160c01b0319909a168a17600160a01b63ffffffff98891602176001600160c01b0316600160c01b948816949094026001600160e01b031693909317600160e01b93871693909302929092179098556080808b018051600580546001600160a01b031916918d169190911790558451988952955185169688019690965290518316918601919091525116938301939093529151909216908201527fa55bd56595c45f517e5967a3067f3dca684445a3080e7c04a4e0d5a40cda627d9060a00160405180910390a150565b60005b8151811015620005ca576000828281518110620003e557620003e562000981565b60200260200101519050600081602001519050806001600160401b0316600003620004235760405163c656089560e01b815260040160405180910390fd5b81516001600160a01b03166200044c576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160401b038116600090815260066020526040902060608301516001820180546200047a9062000997565b9050600003620004d8578154600160a81b600160e81b031916600160a81b1782556040516001600160401b03841681527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb99060200160405180910390a15b805115806200050e5750604080516000602082015201604051602081830303815290604052805190602001208180519060200120145b156200052d576040516342bcdf7f60e11b815260040160405180910390fd5b600182016200053d828262000a28565b50604080850151835486516001600160a01b03166001600160a01b0319921515600160a01b02929092166001600160a81b031990911617178355516001600160401b038416907f49f51971edd25182e97182d6ea372a0488ce2ab639f6a3a7ab4df0d2636fe56b90620005b290859062000af4565b60405180910390a250505050806001019050620003c4565b5050565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715620006095762000609620005ce565b60405290565b604051601f8201601f191681016001600160401b03811182821017156200063a576200063a620005ce565b604052919050565b80516001600160401b03811681146200065a57600080fd5b919050565b6001600160a01b03811681146200067557600080fd5b50565b805163ffffffff811681146200065a57600080fd5b6000601f83601f840112620006a157600080fd5b825160206001600160401b0380831115620006c057620006c0620005ce565b8260051b620006d18382016200060f565b9384528681018301938381019089861115620006ec57600080fd5b84890192505b858310156200081e578251848111156200070c5760008081fd5b89016080601f19828d038101821315620007265760008081fd5b62000730620005e4565b888401516200073f816200065f565b815260406200075085820162000642565b8a83015260608086015180151581146200076a5760008081fd5b83830152938501519389851115620007825760008081fd5b84860195508f603f8701126200079a57600094508485fd5b8a860151945089851115620007b357620007b3620005ce565b620007c48b858f880116016200060f565b93508484528f82868801011115620007dc5760008081fd5b60005b85811015620007fc578681018301518582018d01528b01620007df565b5060009484018b019490945250918201528352509184019190840190620006f2565b9998505050505050505050565b60008060008385036101408112156200084357600080fd5b60808112156200085257600080fd5b6200085c620005e4565b620008678662000642565b8152602086015162000879816200065f565b602082015260408601516200088e816200065f565b60408201526060860151620008a3816200065f565b6060820152935060a0607f1982011215620008bd57600080fd5b5060405160a081016001600160401b038082118383101715620008e457620008e4620005ce565b8160405260808701519150620008fa826200065f565b8183526200090b60a0880162000678565b60208401526200091e60c0880162000678565b60408401526200093160e0880162000678565b6060840152610100870151915062000949826200065f565b60808301829052610120870151929450808311156200096757600080fd5b505062000977868287016200068d565b9150509250925092565b634e487b7160e01b600052603260045260246000fd5b600181811c90821680620009ac57607f821691505b602082108103620009cd57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111562000a23576000816000526020600020601f850160051c81016020861015620009fe5750805b601f850160051c820191505b8181101562000a1f5782815560010162000a0a565b5050505b505050565b81516001600160401b0381111562000a445762000a44620005ce565b62000a5c8162000a55845462000997565b84620009d3565b602080601f83116001811462000a94576000841562000a7b5750858301515b600019600386901b1c1916600185901b17855562000a1f565b600085815260208120601f198616915b8281101562000ac55788860151825594840194600190910190840162000aa4565b508582101562000ae45787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b602080825282546001600160a01b0381168383015260a081901c60ff161515604084015260a81c6001600160401b0316606083015260808083015260018084018054600093929190849062000b498162000997565b8060a089015260c0600183166000811462000b6d576001811462000b8a5762000bbc565b60ff19841660c08b015260c083151560051b8b0101945062000bbc565b85600052602060002060005b8481101562000bb35781548c820185015290880190890162000b96565b8b0160c0019550505b50929998505050505050505050565b60805160a05160c05160e05161010051615f3f62000c416000396000818161024001526129bd0152600081816102110152612eb10152600081816101e20152818161077f01528181610986015261233f0152600081816101b201526125cc01526000818161172d01526117790152615f3f6000f3fe608060405234801561001057600080fd5b50600436106101515760003560e01c806379ba5097116100cd578063c673e58411610081578063e9d68a8e11610066578063e9d68a8e1461052a578063f2fde38b1461054a578063f716f99f1461055d57600080fd5b8063c673e584146104c5578063ccd37ba3146104e557600080fd5b80638da5cb5b116100b25780638da5cb5b14610484578063991a50181461049f578063a80036b4146104b257600080fd5b806379ba50971461046e57806385572ffb1461047657600080fd5b80632d04ab76116101245780633f4b04aa116101095780633f4b04aa146103365780635e36480c146103525780637437ff9f1461037257600080fd5b80632d04ab7614610310578063311cd5131461032357600080fd5b806304666f9c1461015657806305d938b51461016b57806306285c691461017e578063181f5a77146102c7575b600080fd5b61016961016436600461407d565b610570565b005b610169610179366004614707565b610584565b61027060408051608081018252600080825260208201819052918101829052606081019190915260405180608001604052807f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b6040516102be9190815167ffffffffffffffff1681526020808301516001600160a01b0390811691830191909152604080840151821690830152606092830151169181019190915260800190565b60405180910390f35b6103036040518060400160405280601181526020017f4f666652616d7020312e362e302d64657600000000000000000000000000000081525081565b6040516102be9190614882565b61016961031e36600461492d565b610729565b6101696103313660046149e0565b610cbf565b60095460405167ffffffffffffffff90911681526020016102be565b610365610360366004614a34565b610d28565b6040516102be9190614a91565b6104116040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a0810182526004546001600160a01b03808216835263ffffffff600160a01b83048116602085015278010000000000000000000000000000000000000000000000008304811694840194909452600160e01b9091049092166060820152600554909116608082015290565b6040516102be9190600060a0820190506001600160a01b03808451168352602084015163ffffffff808216602086015280604087015116604086015280606087015116606086015250508060808501511660808401525092915050565b610169610d7e565b610169610151366004614a9f565b6000546040516001600160a01b0390911681526020016102be565b6101696104ad366004614aee565b610e3c565b6101696104c0366004614b62565b610e4d565b6104d86104d3366004614bcf565b6111c0565b6040516102be9190614c2f565b61051c6104f3366004614ca4565b67ffffffffffffffff919091166000908152600860209081526040808320938352929052205490565b6040519081526020016102be565b61053d610538366004614cce565b61131e565b6040516102be9190614ce9565b610169610558366004614d37565b61142b565b61016961056b366004614dbc565b61143c565b61057861147e565b610581816114da565b50565b61058c61172a565b8151815181146105c8576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b818110156107195760008482815181106105e7576105e7614efa565b6020026020010151905060008160200151519050600085848151811061060f5761060f614efa565b6020026020010151905080518214610653576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8281101561070a57600082828151811061067257610672614efa565b6020026020010151905080600014610701578460200151828151811061069a5761069a614efa565b6020026020010151608001518110156107015784516040517fc8e9605100000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260248101839052604481018290526064015b60405180910390fd5b50600101610656565b505050508060010190506105cb565b5061072483836117ab565b505050565b600061073787890189615129565b602081015151909150156107e357602081015160408083015190517fccd5af000000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169263ccd5af00926107b29260040161538f565b60006040518083038186803b1580156107ca57600080fd5b505afa1580156107de573d6000803e3d6000fd5b505050505b805151511515806107f957508051602001515115155b156108f95760095460208a01359067ffffffffffffffff808316911610156108b8576009805467ffffffffffffffff191667ffffffffffffffff83161790556004805483516040517f3937306f0000000000000000000000000000000000000000000000000000000081526001600160a01b0390921692633937306f92610881929101615471565b600060405180830381600087803b15801561089b57600080fd5b505af11580156108af573d6000803e3d6000fd5b505050506108f7565b8160200151516000036108f7576040517f2261116700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b60005b816020015151811015610c085760008260200151828151811061092157610921614efa565b602090810291909101015180516040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608083901b166004820152919250906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690632cbc26bb90602401602060405180830381865afa1580156109cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109f19190615484565b15610a34576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016106f8565b6000610a3f8261185b565b6040840151815491925067ffffffffffffffff908116600160a81b90920416141580610a865750826060015167ffffffffffffffff16836040015167ffffffffffffffff16115b15610ae5578251604080850151606086015191517fd5e0f0d600000000000000000000000000000000000000000000000000000000815267ffffffffffffffff93841660048201529083166024820152911660448201526064016106f8565b608083015180610b21576040517f504570e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835167ffffffffffffffff16600090815260086020908152604080832084845290915290205415610b945783516040517f32cf0cbf00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602481018290526044016106f8565b6060840151610ba49060016154b7565b82547fffffff0000000000000000ffffffffffffffffffffffffffffffffffffffffff16600160a81b67ffffffffffffffff9283160217909255925116600090815260086020908152604080832094835293905291909120429055506001016108fc565b507f23bc80217a08968cec0790cd045b396fa7eea0a21af469e603329940b883d86d81604051610c3891906154df565b60405180910390a1610cb460008a8a8a8a8a8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c9182918501908490808284376000920191909152508b92506118c2915050565b505050505050505050565b610cff610cce82840184615537565b6040805160008082526020820190925290610cf9565b6060815260200190600190039081610ce45790505b506117ab565b604080516000808252602082019092529050610d226001858585858660006118c2565b50505050565b6000610d366001600461556c565b6002610d43608085615595565b67ffffffffffffffff16610d5791906155bc565b610d618585611c39565b901c166003811115610d7557610d75614a67565b90505b92915050565b6001546001600160a01b03163314610dd85760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016106f8565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610e4461147e565b61058181611c80565b333014610e86576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160008082526020820190925281610ec3565b6040805180820190915260008082526020820152815260200190600190039081610e9c5790505b5060a08501515190915015610ef757610ef48460a00151856020015186606001518760000151602001518787611de6565b90505b6040805160a0810182528551518152855160209081015167ffffffffffffffff1681830152808701518351600094840192610f33929101614882565b60408051601f19818403018152918152908252878101516020830152018390526005549091506001600160a01b03168015611040576040517f08d450a10000000000000000000000000000000000000000000000000000000081526001600160a01b038216906308d450a190610fad908590600401615675565b600060405180830381600087803b158015610fc757600080fd5b505af1925050508015610fd8575060015b611040573d808015611006576040519150601f19603f3d011682016040523d82523d6000602084013e61100b565b606091505b50806040517f09c253250000000000000000000000000000000000000000000000000000000081526004016106f89190614882565b60408601515115801561105557506080860151155b8061106c575060608601516001600160a01b03163b155b806110ac575060608601516110aa906001600160a01b03167f85572ffb00000000000000000000000000000000000000000000000000000000611f05565b155b156110b957505050505050565b855160209081015167ffffffffffffffff1660009081526006909152604080822054608089015160608a015192517f3cf9798300000000000000000000000000000000000000000000000000000000815284936001600160a01b0390931692633cf97983926111319289926113889291600401615688565b6000604051808303816000875af1158015611150573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261117891908101906156c4565b5091509150816111b657806040517f0a8d6e8c0000000000000000000000000000000000000000000000000000000081526004016106f89190614882565b5050505050505050565b6112036040805160e081019091526000606082018181526080830182905260a0830182905260c08301919091528190815260200160608152602001606081525090565b60ff808316600090815260026020818152604092839020835160e081018552815460608201908152600183015480881660808401526101008104881660a0840152620100009004909616151560c0820152948552918201805484518184028101840190955280855292938583019390928301828280156112ac57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161128e575b505050505081526020016003820180548060200260200160405190810160405280929190818152602001828054801561130e57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116112f0575b5050505050815250509050919050565b604080516080808201835260008083526020808401829052838501829052606080850181905267ffffffffffffffff878116845260068352928690208651948501875280546001600160a01b0381168652600160a01b810460ff16151593860193909352600160a81b9092049092169483019490945260018401805493949293918401916113ab9061575a565b80601f01602080910402602001604051908101604052809291908181526020018280546113d79061575a565b801561130e5780601f106113f95761010080835404028352916020019161130e565b820191906000526020600020905b81548152906001019060200180831161140757505050919092525091949350505050565b61143361147e565b61058181611f21565b61144461147e565b60005b815181101561147a5761147282828151811061146557611465614efa565b6020026020010151611fd7565b600101611447565b5050565b6000546001600160a01b031633146114d85760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016106f8565b565b60005b815181101561147a5760008282815181106114fa576114fa614efa565b602002602001015190506000816020015190508067ffffffffffffffff16600003611551576040517fc656089500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81516001600160a01b0316611579576040516342bcdf7f60e11b815260040160405180910390fd5b67ffffffffffffffff8116600090815260066020526040902060608301516001820180546115a69061575a565b90506000036116195781547fffffff0000000000000000ffffffffffffffffffffffffffffffffffffffffff16600160a81b17825560405167ffffffffffffffff841681527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb99060200160405180910390a15b8051158061164e5750604080516000602082015201604051602081830303815290604052805190602001208180519060200120145b1561166c576040516342bcdf7f60e11b815260040160405180910390fd5b6001820161167a82826157e4565b50604080850151835486516001600160a01b031673ffffffffffffffffffffffffffffffffffffffff19921515600160a01b02929092167fffffffffffffffffffffff000000000000000000000000000000000000000000909116171783555167ffffffffffffffff8416907f49f51971edd25182e97182d6ea372a0488ce2ab639f6a3a7ab4df0d2636fe56b906117139085906158a4565b60405180910390a2505050508060010190506114dd565b467f0000000000000000000000000000000000000000000000000000000000000000146114d8576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201524660248201526044016106f8565b81516000036117e5576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160408051600080825260208201909252911591905b84518110156118545761184c85828151811061181a5761181a614efa565b6020026020010151846118465785838151811061183957611839614efa565b60200260200101516122eb565b836122eb565b6001016117fc565b5050505050565b67ffffffffffffffff811660009081526006602052604081208054600160a01b900460ff16610d78576040517fed053c5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024016106f8565b60ff878116600090815260026020908152604080832081516080810183528154815260019091015480861693820193909352610100830485169181019190915262010000909104909216151560608301528735906119218760a4615972565b905082606001511561196957845161193a9060206155bc565b86516119479060206155bc565b6119529060a0615972565b61195c9190615972565b6119669082615972565b90505b3681146119ab576040517f8e1192e1000000000000000000000000000000000000000000000000000000008152600481018290523660248201526044016106f8565b50815181146119f35781516040517f93df584c0000000000000000000000000000000000000000000000000000000081526004810191909152602481018290526044016106f8565b6119fb61172a565b60ff808a1660009081526003602090815260408083203384528252808320815180830190925280548086168352939491939092840191610100909104166002811115611a4957611a49614a67565b6002811115611a5a57611a5a614a67565b9052509050600281602001516002811115611a7757611a77614a67565b148015611acb5750600260008b60ff1660ff168152602001908152602001600020600301816000015160ff1681548110611ab357611ab3614efa565b6000918252602090912001546001600160a01b031633145b611b01576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50816060015115611be3576020820151611b1c906001615985565b60ff16855114611b58576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8351855114611b93576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008787604051611ba592919061599e565b604051908190038120611bbc918b906020016159ae565b604051602081830303815290604052805190602001209050611be18a82888888612c6b565b505b6040805182815260208a81013567ffffffffffffffff169082015260ff8b16917f198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef0910160405180910390a2505050505050505050565b67ffffffffffffffff8216600090815260076020526040812081611c5e6080856159c2565b67ffffffffffffffff1681526020810191909152604001600020549392505050565b80516001600160a01b0316611ca8576040516342bcdf7f60e11b815260040160405180910390fd5b80516004805460208085018051604080880180516060808b0180516001600160a01b039b8c167fffffffffffffffff000000000000000000000000000000000000000000000000909a168a17600160a01b63ffffffff988916021777ffffffffffffffffffffffffffffffffffffffffffffffff167801000000000000000000000000000000000000000000000000948816949094026001600160e01b031693909317600160e01b93871693909302929092179098556080808b0180516005805473ffffffffffffffffffffffffffffffffffffffff1916918d169190911790558451988952955185169688019690965290518316918601919091525116938301939093529151909216908201527fa55bd56595c45f517e5967a3067f3dca684445a3080e7c04a4e0d5a40cda627d9060a00160405180910390a150565b6060865167ffffffffffffffff811115611e0257611e02613e94565b604051908082528060200260200182016040528015611e4757816020015b6040805180820190915260008082526020820152815260200190600190039081611e205790505b50905060005b8751811015611ef957611ed4888281518110611e6b57611e6b614efa565b6020026020010151888888888887818110611e8857611e88614efa565b9050602002810190611e9a91906159e9565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250612e5092505050565b828281518110611ee657611ee6614efa565b6020908102919091010152600101611e4d565b505b9695505050505050565b6000611f10836131f5565b8015610d755750610d758383613259565b336001600160a01b03821603611f795760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016106f8565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b806040015160ff16600003612002576000604051631b3fab5160e11b81526004016106f89190615a30565b60208082015160ff8082166000908152600290935260408320600181015492939092839216900361206f57606084015160018201805491151562010000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff9092169190911790556120c4565b6060840151600182015460ff62010000909104161515901515146120c4576040517f87f6037c00000000000000000000000000000000000000000000000000000000815260ff841660048201526024016106f8565b60a0840151805161010010156120f0576001604051631b3fab5160e11b81526004016106f89190615a30565b612156848460030180548060200260200160405190810160405280929190818152602001828054801561214c57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161212e575b5050505050613314565b846060015115612260576121c4848460020180548060200260200160405190810160405280929190818152602001828054801561214c576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161212e575050505050613314565b6080850151805161010010156121f0576002604051631b3fab5160e11b81526004016106f89190615a30565b6040860151612200906003615a4a565b60ff16815111612226576003604051631b3fab5160e11b81526004016106f89190615a30565b805160018401805461ff00191661010060ff8416021790556122519060028601906020840190613e0d565b5061225e8582600161337d565b505b61226c8482600261337d565b80516122819060038501906020840190613e0d565b5060408581015160018401805460ff191660ff8316179055865180855560a088015192517fab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f547936122da9389939260028a01929190615a66565b60405180910390a1611854846134f1565b815181516040517f2cbc26bb000000000000000000000000000000000000000000000000000000008152608083901b77ffffffffffffffff00000000000000000000000000000000166004820152901515907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632cbc26bb90602401602060405180830381865afa15801561238e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123b29190615484565b1561243e5780156123fb576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff831660048201526024016106f8565b60405167ffffffffffffffff831681527faab522ed53d887e56ed53dd37398a01aeef6a58e0fa77c2173beb9512d8949339060200160405180910390a150505050565b60006124498361185b565b60010180546124579061575a565b80601f01602080910402602001604051908101604052809291908181526020018280546124839061575a565b80156124d05780601f106124a5576101008083540402835291602001916124d0565b820191906000526020600020905b8154815290600101906020018083116124b357829003601f168201915b50505060208801515192935050506000819003612518576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8560400151518114612556576040517f57e0e08300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff81111561257157612571613e94565b60405190808252806020026020018201604052801561259a578160200160208202803683370190505b50905060005b828110156126f3576000886020015182815181106125c0576125c0614efa565b602002602001015190507f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681600001516040015167ffffffffffffffff161461265357805160409081015190517f38432a2200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016106f8565b8667ffffffffffffffff1681600001516020015167ffffffffffffffff16146126c3578051602001516040517f6c95f1eb00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808a16600483015290911660248201526044016106f8565b6126cd8186613561565b8383815181106126df576126df614efa565b6020908102919091010152506001016125a0565b50600061270a86838a606001518b60800151613683565b905080600003612752576040517f7dd17a7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff871660048201526024016106f8565b60005b83811015610cb45760005a905060008a60200151838151811061277a5761277a614efa565b6020026020010151905060006127988a836000015160600151610d28565b905060008160038111156127ae576127ae614a67565b14806127cb575060038160038111156127c9576127c9614a67565b145b612823578151606001516040805167ffffffffffffffff808e16825290921660208301527f3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c91015b60405180910390a1505050612c63565b88156128f357600454600090600160a01b900463ffffffff16612846874261556c565b11905080806128665750600382600381111561286457612864614a67565b145b6128a8576040517fa9cfc86200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526024016106f8565b8b85815181106128ba576128ba614efa565b60200260200101516000146128ed578b85815181106128db576128db614efa565b60200260200101518360800181815250505b50612954565b600081600381111561290757612907614a67565b14612954578151606001516040805167ffffffffffffffff808e16825290921660208301527f3ef2a99c550a751d4b0b261268f05a803dfb049ab43616a1ffb388f61fe651209101612813565b81516080015167ffffffffffffffff1615612a4357600081600381111561297d5761297d614a67565b03612a435781516080015160208301516040517fe0e03cae0000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169263e0e03cae926129f4928f929190600401615b12565b6020604051808303816000875af1158015612a13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a379190615484565b612a4357505050612c63565b60008c604001518581518110612a5b57612a5b614efa565b6020026020010151905080518360a001515114612abf578251606001516040517f1cfe6d8b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808e16600483015290911660248201526044016106f8565b612ad38b84600001516060015160016136c1565b600080612ae08584613769565b91509150612af78d866000015160600151846136c1565b8b15612b67576003826003811115612b1157612b11614a67565b03612b67576000846003811115612b2a57612b2a614a67565b14612b67578451516040517f2b11b8d90000000000000000000000000000000000000000000000000000000081526106f891908390600401615b3f565b6002826003811115612b7b57612b7b614a67565b14612bd5576003826003811115612b9457612b94614a67565b14612bd5578451606001516040517f926c5a3e0000000000000000000000000000000000000000000000000000000081526106f8918f918590600401615b58565b84600001516000015185600001516060015167ffffffffffffffff168e67ffffffffffffffff167f05665fe9ad095383d018353f4cbcba77e84db27dd215081bbf7cdf9ae6fbe48b8c8b81518110612c2f57612c2f614efa565b602002602001015186865a612c44908e61556c565b604051612c549493929190615b7e565b60405180910390a45050505050505b600101612755565b8251600090815b818110156111b6576000600188868460208110612c9157612c91614efa565b612c9e91901a601b615985565b898581518110612cb057612cb0614efa565b6020026020010151898681518110612cca57612cca614efa565b602002602001015160405160008152602001604052604051612d08949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015612d2a573d6000803e3d6000fd5b505060408051601f1981015160ff808e166000908152600360209081528582206001600160a01b03851683528152858220858701909652855480841686529397509095509293928401916101009004166002811115612d8b57612d8b614a67565b6002811115612d9c57612d9c614a67565b9052509050600181602001516002811115612db957612db9614a67565b14612df0576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600160ff9091161b851615612e33576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000015160ff166001901b851794505050806001019050612c72565b60408051808201909152600080825260208201526000612e738760200151613833565b6040517fbbe4f6db0000000000000000000000000000000000000000000000000000000081526001600160a01b0380831660048301529192506000917f0000000000000000000000000000000000000000000000000000000000000000169063bbe4f6db90602401602060405180830381865afa158015612ef8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f1c9190615bb5565b90506001600160a01b0381161580612f645750612f626001600160a01b0382167faff2afbf00000000000000000000000000000000000000000000000000000000611f05565b155b15612fa6576040517fae9b4ce90000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016106f8565b6004546000908190612fc89089908690600160e01b900463ffffffff166138d9565b9150915060008060006130956040518061010001604052808e81526020018c67ffffffffffffffff1681526020018d6001600160a01b031681526020018f606001518152602001896001600160a01b031681526020018f6000015181526020018f6040015181526020018b8152506040516024016130469190615bd2565b60408051601f198184030181529190526020810180516001600160e01b03167f390775370000000000000000000000000000000000000000000000000000000017905287866113886084613a07565b925092509250826130d457816040517fe1cd55090000000000000000000000000000000000000000000000000000000081526004016106f89190614882565b815160201461311c5781516040517f78ef80240000000000000000000000000000000000000000000000000000000081526020600482015260248101919091526044016106f8565b6000828060200190518101906131329190615c9f565b9050866001600160a01b03168c6001600160a01b0316146131c75760006131638d8a61315e868a61556c565b6138d9565b5090508681108061317d57508161317a888361556c565b14155b156131c5576040517fa966e21f0000000000000000000000000000000000000000000000000000000081526004810183905260248101889052604481018290526064016106f8565b505b604080518082019091526001600160a01b039098168852602088015250949550505050505095945050505050565b6000613221827f01ffc9a700000000000000000000000000000000000000000000000000000000613259565b8015610d785750613252827fffffffff00000000000000000000000000000000000000000000000000000000613259565b1592915050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000082166024820152600090819060440160408051601f19818403018152919052602080820180516001600160e01b03167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825192935060009283928392909183918a617530fa92503d915060005190508280156132fd575060208210155b80156133095750600081115b979650505050505050565b60005b81518110156107245760ff83166000908152600360205260408120835190919084908490811061334957613349614efa565b6020908102919091018101516001600160a01b03168252810191909152604001600020805461ffff19169055600101613317565b60005b8251811015610d2257600083828151811061339d5761339d614efa565b60200260200101519050600060028111156133ba576133ba614a67565b60ff80871660009081526003602090815260408083206001600160a01b038716845290915290205461010090041660028111156133f9576133f9614a67565b1461341a576004604051631b3fab5160e11b81526004016106f89190615a30565b6001600160a01b03811661345a576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405180604001604052808360ff16815260200184600281111561348057613480614a67565b905260ff80871660009081526003602090815260408083206001600160a01b0387168452825290912083518154931660ff198416811782559184015190929091839161ffff1916176101008360028111156134dd576134dd614a67565b021790555090505050806001019050613380565b60ff81166105815760ff8082166000908152600260205260409020600101546201000090041661354d576040517f7a3b4c9400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6009805467ffffffffffffffff1916905550565b8151602080820151604092830151925160009384936135a7937f2425b0b9f9054c76ff151b0a175b18f37a4a4e82013a72e9f15c9caa095ed21f93909291889101615cb8565b60408051601f1981840301815290829052805160209182012086518051888401516060808b0151908401516080808d015195015195976135f09794969395929491939101615ceb565b604051602081830303815290604052805190602001208560400151805190602001208660a001516040516020016136279190615df0565b60408051601f198184030181528282528051602091820120908301969096528101939093526060830191909152608082015260a081019190915260c0015b60405160208183030381529060405280519060200120905092915050565b600080613691858585613b2d565b67ffffffffffffffff8716600090815260086020908152604080832093835292905220549150505b949350505050565b600060026136d0608085615595565b67ffffffffffffffff166136e491906155bc565b905060006136f28585611c39565b9050816137016001600461556c565b901b19168183600381111561371857613718614a67565b67ffffffffffffffff871660009081526007602052604081209190921b929092179182916137476080886159c2565b67ffffffffffffffff1681526020810191909152604001600020555050505050565b6040517fa80036b4000000000000000000000000000000000000000000000000000000008152600090606090309063a80036b4906137ad9087908790600401615e50565b600060405180830381600087803b1580156137c757600080fd5b505af19250505080156137d8575060015b613817573d808015613806576040519150601f19603f3d011682016040523d82523d6000602084013e61380b565b606091505b5060039250905061382c565b50506040805160208101909152600081526002905b9250929050565b6000815160201461387257816040517f8d666f600000000000000000000000000000000000000000000000000000000081526004016106f89190614882565b6000828060200190518101906138889190615c9f565b90506001600160a01b038111806138a0575061040081105b15610d7857826040517f8d666f600000000000000000000000000000000000000000000000000000000081526004016106f89190614882565b60008060008060006139538860405160240161390491906001600160a01b0391909116815260200190565b60408051601f198184030181529190526020810180516001600160e01b03167f70a082310000000000000000000000000000000000000000000000000000000017905288886113886084613a07565b9250925092508261399257816040517fe1cd55090000000000000000000000000000000000000000000000000000000081526004016106f89190614882565b60208251146139da5781516040517f78ef80240000000000000000000000000000000000000000000000000000000081526020600482015260248101919091526044016106f8565b818060200190518101906139ee9190615c9f565b6139f8828861556c565b94509450505050935093915050565b6000606060008361ffff1667ffffffffffffffff811115613a2a57613a2a613e94565b6040519080825280601f01601f191660200182016040528015613a54576020820181803683370190505b509150863b613a87577f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b5a85811015613aba577fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b8590036040810481038710613af3577f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b505a6000808a5160208c0160008c8cf193505a900390503d84811115613b165750835b808352806000602085013e50955095509592505050565b8251825160009190818303613b6e576040517f11a6b26400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6101018211801590613b8257506101018111155b613b9f576040516309bde33960e01b815260040160405180910390fd5b60001982820101610100811115613bc9576040516309bde33960e01b815260040160405180910390fd5b80600003613bf65786600081518110613be457613be4614efa565b60200260200101519350505050613dc5565b60008167ffffffffffffffff811115613c1157613c11613e94565b604051908082528060200260200182016040528015613c3a578160200160208202803683370190505b50905060008080805b85811015613d645760006001821b8b811603613c9e5788851015613c87578c5160018601958e918110613c7857613c78614efa565b60200260200101519050613cc0565b8551600185019487918110613c7857613c78614efa565b8b5160018401938d918110613cb557613cb5614efa565b602002602001015190505b600089861015613cf0578d5160018701968f918110613ce157613ce1614efa565b60200260200101519050613d12565b8651600186019588918110613d0757613d07614efa565b602002602001015190505b82851115613d33576040516309bde33960e01b815260040160405180910390fd5b613d3d8282613dcc565b878481518110613d4f57613d4f614efa565b60209081029190910101525050600101613c43565b506001850382148015613d7657508683145b8015613d8157508581145b613d9e576040516309bde33960e01b815260040160405180910390fd5b836001860381518110613db357613db3614efa565b60200260200101519750505050505050505b9392505050565b6000818310613de457613ddf8284613dea565b610d75565b610d7583835b604080516001602082015290810183905260608101829052600090608001613665565b828054828255906000526020600020908101928215613e6f579160200282015b82811115613e6f578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190613e2d565b50613e7b929150613e7f565b5090565b5b80821115613e7b5760008155600101613e80565b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715613ecd57613ecd613e94565b60405290565b60405160a0810167ffffffffffffffff81118282101715613ecd57613ecd613e94565b60405160c0810167ffffffffffffffff81118282101715613ecd57613ecd613e94565b6040805190810167ffffffffffffffff81118282101715613ecd57613ecd613e94565b6040516060810167ffffffffffffffff81118282101715613ecd57613ecd613e94565b604051601f8201601f1916810167ffffffffffffffff81118282101715613f8857613f88613e94565b604052919050565b600067ffffffffffffffff821115613faa57613faa613e94565b5060051b60200190565b6001600160a01b038116811461058157600080fd5b803567ffffffffffffffff81168114613fe157600080fd5b919050565b801515811461058157600080fd5b8035613fe181613fe6565b600067ffffffffffffffff82111561401957614019613e94565b50601f01601f191660200190565b600082601f83011261403857600080fd5b813561404b61404682613fff565b613f5f565b81815284602083860101111561406057600080fd5b816020850160208301376000918101602001919091529392505050565b6000602080838503121561409057600080fd5b823567ffffffffffffffff808211156140a857600080fd5b818501915085601f8301126140bc57600080fd5b81356140ca61404682613f90565b81815260059190911b830184019084810190888311156140e957600080fd5b8585015b8381101561418f578035858111156141055760008081fd5b86016080818c03601f190181131561411d5760008081fd5b614125613eaa565b8983013561413281613fb4565b81526040614141848201613fc9565b8b83015260608085013561415481613fe6565b8383015292840135928984111561416d57600091508182fd5b61417b8f8d86880101614027565b9083015250855250509186019186016140ed565b5098975050505050505050565b600060a082840312156141ae57600080fd5b6141b6613ed3565b9050813581526141c860208301613fc9565b60208201526141d960408301613fc9565b60408201526141ea60608301613fc9565b60608201526141fb60808301613fc9565b608082015292915050565b8035613fe181613fb4565b600082601f83011261422257600080fd5b8135602061423261404683613f90565b82815260059290921b8401810191818101908684111561425157600080fd5b8286015b84811015611ef957803567ffffffffffffffff808211156142765760008081fd5b9088019060a0828b03601f19018113156142905760008081fd5b614298613ed3565b87840135838111156142aa5760008081fd5b6142b88d8a83880101614027565b825250604080850135848111156142cf5760008081fd5b6142dd8e8b83890101614027565b8a84015250606080860135858111156142f65760008081fd5b6143048f8c838a0101614027565b8385015250608091508186013581840152508285013592508383111561432a5760008081fd5b6143388d8a85880101614027565b908201528652505050918301918301614255565b6000610140828403121561435f57600080fd5b614367613ef6565b9050614373838361419c565b815260a082013567ffffffffffffffff8082111561439057600080fd5b61439c85838601614027565b602084015260c08401359150808211156143b557600080fd5b6143c185838601614027565b60408401526143d260e08501614206565b606084015261010084013560808401526101208401359150808211156143f757600080fd5b5061440484828501614211565b60a08301525092915050565b600082601f83011261442157600080fd5b8135602061443161404683613f90565b82815260059290921b8401810191818101908684111561445057600080fd5b8286015b84811015611ef957803567ffffffffffffffff8111156144745760008081fd5b6144828986838b010161434c565b845250918301918301614454565b600082601f8301126144a157600080fd5b813560206144b161404683613f90565b82815260059290921b840181019181810190868411156144d057600080fd5b8286015b84811015611ef957803567ffffffffffffffff808211156144f457600080fd5b818901915089603f83011261450857600080fd5b8582013561451861404682613f90565b81815260059190911b830160400190878101908c83111561453857600080fd5b604085015b838110156145715780358581111561455457600080fd5b6145638f6040838a0101614027565b84525091890191890161453d565b508752505050928401925083016144d4565b600082601f83011261459457600080fd5b813560206145a461404683613f90565b8083825260208201915060208460051b8701019350868411156145c657600080fd5b602086015b84811015611ef957803583529183019183016145cb565b600082601f8301126145f357600080fd5b8135602061460361404683613f90565b82815260059290921b8401810191818101908684111561462257600080fd5b8286015b84811015611ef957803567ffffffffffffffff808211156146475760008081fd5b9088019060a0828b03601f19018113156146615760008081fd5b614669613ed3565b614674888501613fc9565b81526040808501358481111561468a5760008081fd5b6146988e8b83890101614410565b8a84015250606080860135858111156146b15760008081fd5b6146bf8f8c838a0101614490565b83850152506080915081860135858111156146da5760008081fd5b6146e88f8c838a0101614583565b9184019190915250919093013590830152508352918301918301614626565b600080604080848603121561471b57600080fd5b833567ffffffffffffffff8082111561473357600080fd5b61473f878388016145e2565b945060209150818601358181111561475657600080fd5b8601601f8101881361476757600080fd5b803561477561404682613f90565b81815260059190911b8201840190848101908a83111561479457600080fd5b8584015b83811015614820578035868111156147b05760008081fd5b8501603f81018d136147c25760008081fd5b878101356147d261404682613f90565b81815260059190911b82018a0190898101908f8311156147f25760008081fd5b928b01925b828410156148105783358252928a0192908a01906147f7565b8652505050918601918601614798565b50809750505050505050509250929050565b60005b8381101561484d578181015183820152602001614835565b50506000910152565b6000815180845261486e816020860160208601614832565b601f01601f19169290920160200192915050565b602081526000610d756020830184614856565b8060608101831015610d7857600080fd5b60008083601f8401126148b857600080fd5b50813567ffffffffffffffff8111156148d057600080fd5b60208301915083602082850101111561382c57600080fd5b60008083601f8401126148fa57600080fd5b50813567ffffffffffffffff81111561491257600080fd5b6020830191508360208260051b850101111561382c57600080fd5b60008060008060008060008060e0898b03121561494957600080fd5b6149538a8a614895565b9750606089013567ffffffffffffffff8082111561497057600080fd5b61497c8c838d016148a6565b909950975060808b013591508082111561499557600080fd5b6149a18c838d016148e8565b909750955060a08b01359150808211156149ba57600080fd5b506149c78b828c016148e8565b999c989b50969995989497949560c00135949350505050565b6000806000608084860312156149f557600080fd5b6149ff8585614895565b9250606084013567ffffffffffffffff811115614a1b57600080fd5b614a27868287016148a6565b9497909650939450505050565b60008060408385031215614a4757600080fd5b614a5083613fc9565b9150614a5e60208401613fc9565b90509250929050565b634e487b7160e01b600052602160045260246000fd5b60048110614a8d57614a8d614a67565b9052565b60208101610d788284614a7d565b600060208284031215614ab157600080fd5b813567ffffffffffffffff811115614ac857600080fd5b820160a08185031215613dc557600080fd5b803563ffffffff81168114613fe157600080fd5b600060a08284031215614b0057600080fd5b614b08613ed3565b8235614b1381613fb4565b8152614b2160208401614ada565b6020820152614b3260408401614ada565b6040820152614b4360608401614ada565b60608201526080830135614b5681613fb4565b60808201529392505050565b600080600060408486031215614b7757600080fd5b833567ffffffffffffffff80821115614b8f57600080fd5b614b9b8783880161434c565b94506020860135915080821115614bb157600080fd5b50614a27868287016148e8565b803560ff81168114613fe157600080fd5b600060208284031215614be157600080fd5b610d7582614bbe565b60008151808452602080850194506020840160005b83811015614c245781516001600160a01b031687529582019590820190600101614bff565b509495945050505050565b60208152600082518051602084015260ff602082015116604084015260ff604082015116606084015260608101511515608084015250602083015160c060a0840152614c7e60e0840182614bea565b90506040840151601f198483030160c0850152614c9b8282614bea565b95945050505050565b60008060408385031215614cb757600080fd5b614cc083613fc9565b946020939093013593505050565b600060208284031215614ce057600080fd5b610d7582613fc9565b602081526001600160a01b03825116602082015260208201511515604082015267ffffffffffffffff6040830151166060820152600060608301516080808401526136b960a0840182614856565b600060208284031215614d4957600080fd5b8135613dc581613fb4565b600082601f830112614d6557600080fd5b81356020614d7561404683613f90565b8083825260208201915060208460051b870101935086841115614d9757600080fd5b602086015b84811015611ef9578035614daf81613fb4565b8352918301918301614d9c565b60006020808385031215614dcf57600080fd5b823567ffffffffffffffff80821115614de757600080fd5b818501915085601f830112614dfb57600080fd5b8135614e0961404682613f90565b81815260059190911b83018401908481019088831115614e2857600080fd5b8585015b8381101561418f57803585811115614e4357600080fd5b860160c0818c03601f19011215614e5a5760008081fd5b614e62613ef6565b8882013581526040614e75818401614bbe565b8a8301526060614e86818501614bbe565b8284015260809150614e99828501613ff4565b9083015260a08381013589811115614eb15760008081fd5b614ebf8f8d83880101614d54565b838501525060c0840135915088821115614ed95760008081fd5b614ee78e8c84870101614d54565b9083015250845250918601918601614e2c565b634e487b7160e01b600052603260045260246000fd5b80356001600160e01b0381168114613fe157600080fd5b600082601f830112614f3857600080fd5b81356020614f4861404683613f90565b82815260069290921b84018101918181019086841115614f6757600080fd5b8286015b84811015611ef95760408189031215614f845760008081fd5b614f8c613f19565b614f9582613fc9565b8152614fa2858301614f10565b81860152835291830191604001614f6b565b600082601f830112614fc557600080fd5b81356020614fd561404683613f90565b82815260059290921b84018101918181019086841115614ff457600080fd5b8286015b84811015611ef957803567ffffffffffffffff808211156150195760008081fd5b9088019060a0828b03601f19018113156150335760008081fd5b61503b613ed3565b615046888501613fc9565b81526040808501358481111561505c5760008081fd5b61506a8e8b83890101614027565b8a840152506060935061507e848601613fc9565b90820152608061508f858201613fc9565b93820193909352920135908201528352918301918301614ff8565b600082601f8301126150bb57600080fd5b813560206150cb61404683613f90565b82815260069290921b840181019181810190868411156150ea57600080fd5b8286015b84811015611ef957604081890312156151075760008081fd5b61510f613f19565b8135815284820135858201528352918301916040016150ee565b6000602080838503121561513c57600080fd5b823567ffffffffffffffff8082111561515457600080fd5b908401906060828703121561516857600080fd5b615170613f3c565b82358281111561517f57600080fd5b8301604081890381131561519257600080fd5b61519a613f19565b8235858111156151a957600080fd5b8301601f81018b136151ba57600080fd5b80356151c861404682613f90565b81815260069190911b8201890190898101908d8311156151e757600080fd5b928a01925b828410156152375785848f0312156152045760008081fd5b61520c613f19565b843561521781613fb4565b8152615224858d01614f10565b818d0152825292850192908a01906151ec565b84525050508287013591508482111561524f57600080fd5b61525b8a838501614f27565b8188015283525050828401358281111561527457600080fd5b61528088828601614fb4565b8583015250604083013593508184111561529957600080fd5b6152a5878585016150aa565b60408201529695505050505050565b600082825180855260208086019550808260051b84010181860160005b8481101561534657601f19868403018952815160a067ffffffffffffffff80835116865286830151828888015261530a83880182614856565b604085810151841690890152606080860151909316928801929092525060809283015192909501919091525097830197908301906001016152d1565b5090979650505050505050565b60008151808452602080850194506020840160005b83811015614c24578151805188528301518388015260409096019590820190600101615368565b6040815260006153a260408301856152b4565b8281036020840152614c9b8185615353565b805160408084528151848201819052600092602091908201906060870190855b8181101561540b57835180516001600160a01b031684528501516001600160e01b03168584015292840192918501916001016153d4565b50508583015187820388850152805180835290840192506000918401905b80831015615465578351805167ffffffffffffffff1683528501516001600160e01b031685830152928401926001929092019190850190615429565b50979650505050505050565b602081526000610d7560208301846153b4565b60006020828403121561549657600080fd5b8151613dc581613fe6565b634e487b7160e01b600052601160045260246000fd5b67ffffffffffffffff8181168382160190808211156154d8576154d86154a1565b5092915050565b6020815260008251606060208401526154fb60808401826153b4565b90506020840151601f198085840301604086015261551983836152b4565b9250604086015191508085840301606086015250614c9b8282615353565b60006020828403121561554957600080fd5b813567ffffffffffffffff81111561556057600080fd5b6136b9848285016145e2565b81810381811115610d7857610d786154a1565b634e487b7160e01b600052601260045260246000fd5b600067ffffffffffffffff808416806155b0576155b061557f565b92169190910692915050565b8082028115828204841417610d7857610d786154a1565b805182526000602067ffffffffffffffff81840151168185015260408084015160a0604087015261560760a0870182614856565b9050606085015186820360608801526156208282614856565b608087810151898303918a01919091528051808352908601935060009250908501905b8083101561546557835180516001600160a01b0316835286015186830152928501926001929092019190840190615643565b602081526000610d7560208301846155d3565b60808152600061569b60808301876155d3565b61ffff9590951660208301525060408101929092526001600160a01b0316606090910152919050565b6000806000606084860312156156d957600080fd5b83516156e481613fe6565b602085015190935067ffffffffffffffff81111561570157600080fd5b8401601f8101861361571257600080fd5b805161572061404682613fff565b81815287602083850101111561573557600080fd5b615746826020830160208601614832565b809450505050604084015190509250925092565b600181811c9082168061576e57607f821691505b60208210810361578e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610724576000816000526020600020601f850160051c810160208610156157bd5750805b601f850160051c820191505b818110156157dc578281556001016157c9565b505050505050565b815167ffffffffffffffff8111156157fe576157fe613e94565b6158128161580c845461575a565b84615794565b602080601f831160018114615847576000841561582f5750858301515b600019600386901b1c1916600185901b1785556157dc565b600085815260208120601f198616915b8281101561587657888601518255948401946001909101908401615857565b50858210156158945787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000602080835283546001600160a01b038116602085015260ff8160a01c161515604085015267ffffffffffffffff8160a81c166060850152506001808501608080860152600081546158f68161575a565b8060a089015260c06001831660008114615917576001811461593357615963565b60ff19841660c08b015260c083151560051b8b01019450615963565b85600052602060002060005b8481101561595a5781548c820185015290880190890161593f565b8b0160c0019550505b50929998505050505050505050565b80820180821115610d7857610d786154a1565b60ff8181168382160190811115610d7857610d786154a1565b8183823760009101908152919050565b828152606082602083013760800192915050565b600067ffffffffffffffff808416806159dd576159dd61557f565b92169190910492915050565b6000808335601e19843603018112615a0057600080fd5b83018035915067ffffffffffffffff821115615a1b57600080fd5b60200191503681900382131561382c57600080fd5b6020810160058310615a4457615a44614a67565b91905290565b60ff81811683821602908116908181146154d8576154d86154a1565b600060a0820160ff881683526020878185015260a0604085015281875480845260c0860191508860005282600020935060005b81811015615abe5784546001600160a01b031683526001948501949284019201615a99565b50508481036060860152865180825290820192508187019060005b81811015615afe5782516001600160a01b031685529383019391830191600101615ad9565b50505060ff85166080850152509050611efb565b600067ffffffffffffffff808616835280851660208401525060606040830152614c9b6060830184614856565b8281526040602082015260006136b96040830184614856565b67ffffffffffffffff848116825283166020820152606081016136b96040830184614a7d565b848152615b8e6020820185614a7d565b608060408201526000615ba46080830185614856565b905082606083015295945050505050565b600060208284031215615bc757600080fd5b8151613dc581613fb4565b6020815260008251610100806020850152615bf1610120850183614856565b91506020850151615c0e604086018267ffffffffffffffff169052565b5060408501516001600160a01b038116606086015250606085015160808501526080850151615c4860a08601826001600160a01b03169052565b5060a0850151601f19808685030160c0870152615c658483614856565b935060c08701519150808685030160e0870152615c828483614856565b935060e0870151915080868503018387015250611efb8382614856565b600060208284031215615cb157600080fd5b5051919050565b848152600067ffffffffffffffff808616602084015280851660408401525060806060830152611efb6080830184614856565b86815260c060208201526000615d0460c0830188614856565b6001600160a01b039690961660408301525067ffffffffffffffff9384166060820152608081019290925290911660a09091015292915050565b600082825180855260208086019550808260051b84010181860160005b8481101561534657601f19868403018952815160a08151818652615d8182870182614856565b9150508582015185820387870152615d998282614856565b91505060408083015186830382880152615db38382614856565b92505050606080830151818701525060808083015192508582038187015250615ddc8183614856565b9a86019a9450505090830190600101615d5b565b602081526000610d756020830184615d3e565b60008282518085526020808601955060208260051b8401016020860160005b8481101561534657601f19868403018952615e3e838351614856565b98840198925090830190600101615e22565b604081526000835180516040840152602081015167ffffffffffffffff80821660608601528060408401511660808601528060608401511660a08601528060808401511660c086015250505060208401516101408060e0850152615eb8610180850183614856565b91506040860151603f198086850301610100870152615ed78483614856565b935060608801519150615ef66101208701836001600160a01b03169052565b60808801518387015260a0880151925080868503016101608701525050615f1d8282615d3e565b9150508281036020840152614c9b8185615e0356fea164736f6c6343000818000a",
}

var OffRampABI = OffRampMetaData.ABI

var OffRampBin = OffRampMetaData.Bin

func DeployOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig OffRampStaticConfig, dynamicConfig OffRampDynamicConfig, sourceChainConfigs []OffRampSourceChainConfigArgs) (common.Address, *types.Transaction, *OffRamp, error) {
	parsed, err := OffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OffRampBin), backend, staticConfig, dynamicConfig, sourceChainConfigs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OffRamp{address: address, abi: *parsed, OffRampCaller: OffRampCaller{contract: contract}, OffRampTransactor: OffRampTransactor{contract: contract}, OffRampFilterer: OffRampFilterer{contract: contract}}, nil
}

type OffRamp struct {
	address common.Address
	abi     abi.ABI
	OffRampCaller
	OffRampTransactor
	OffRampFilterer
}

type OffRampCaller struct {
	contract *bind.BoundContract
}

type OffRampTransactor struct {
	contract *bind.BoundContract
}

type OffRampFilterer struct {
	contract *bind.BoundContract
}

type OffRampSession struct {
	Contract     *OffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OffRampCallerSession struct {
	Contract *OffRampCaller
	CallOpts bind.CallOpts
}

type OffRampTransactorSession struct {
	Contract     *OffRampTransactor
	TransactOpts bind.TransactOpts
}

type OffRampRaw struct {
	Contract *OffRamp
}

type OffRampCallerRaw struct {
	Contract *OffRampCaller
}

type OffRampTransactorRaw struct {
	Contract *OffRampTransactor
}

func NewOffRamp(address common.Address, backend bind.ContractBackend) (*OffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(OffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OffRamp{address: address, abi: abi, OffRampCaller: OffRampCaller{contract: contract}, OffRampTransactor: OffRampTransactor{contract: contract}, OffRampFilterer: OffRampFilterer{contract: contract}}, nil
}

func NewOffRampCaller(address common.Address, caller bind.ContractCaller) (*OffRampCaller, error) {
	contract, err := bindOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OffRampCaller{contract: contract}, nil
}

func NewOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*OffRampTransactor, error) {
	contract, err := bindOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OffRampTransactor{contract: contract}, nil
}

func NewOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*OffRampFilterer, error) {
	contract, err := bindOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OffRampFilterer{contract: contract}, nil
}

func bindOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OffRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OffRamp *OffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffRamp.Contract.OffRampCaller.contract.Call(opts, result, method, params...)
}

func (_OffRamp *OffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRamp.Contract.OffRampTransactor.contract.Transfer(opts)
}

func (_OffRamp *OffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffRamp.Contract.OffRampTransactor.contract.Transact(opts, method, params...)
}

func (_OffRamp *OffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_OffRamp *OffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRamp.Contract.contract.Transfer(opts)
}

func (_OffRamp *OffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_OffRamp *OffRampCaller) CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_OffRamp *OffRampSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _OffRamp.Contract.CcipReceive(&_OffRamp.CallOpts, arg0)
}

func (_OffRamp *OffRampCallerSession) CcipReceive(arg0 ClientAny2EVMMessage) error {
	return _OffRamp.Contract.CcipReceive(&_OffRamp.CallOpts, arg0)
}

func (_OffRamp *OffRampCaller) GetDynamicConfig(opts *bind.CallOpts) (OffRampDynamicConfig, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(OffRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OffRampDynamicConfig)).(*OffRampDynamicConfig)

	return out0, err

}

func (_OffRamp *OffRampSession) GetDynamicConfig() (OffRampDynamicConfig, error) {
	return _OffRamp.Contract.GetDynamicConfig(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetDynamicConfig() (OffRampDynamicConfig, error) {
	return _OffRamp.Contract.GetDynamicConfig(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getExecutionState", sourceChainSelector, sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_OffRamp *OffRampSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _OffRamp.Contract.GetExecutionState(&_OffRamp.CallOpts, sourceChainSelector, sequenceNumber)
}

func (_OffRamp *OffRampCallerSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _OffRamp.Contract.GetExecutionState(&_OffRamp.CallOpts, sourceChainSelector, sequenceNumber)
}

func (_OffRamp *OffRampCaller) GetLatestPriceSequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getLatestPriceSequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_OffRamp *OffRampSession) GetLatestPriceSequenceNumber() (uint64, error) {
	return _OffRamp.Contract.GetLatestPriceSequenceNumber(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetLatestPriceSequenceNumber() (uint64, error) {
	return _OffRamp.Contract.GetLatestPriceSequenceNumber(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) GetMerkleRoot(opts *bind.CallOpts, sourceChainSelector uint64, root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getMerkleRoot", sourceChainSelector, root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OffRamp *OffRampSession) GetMerkleRoot(sourceChainSelector uint64, root [32]byte) (*big.Int, error) {
	return _OffRamp.Contract.GetMerkleRoot(&_OffRamp.CallOpts, sourceChainSelector, root)
}

func (_OffRamp *OffRampCallerSession) GetMerkleRoot(sourceChainSelector uint64, root [32]byte) (*big.Int, error) {
	return _OffRamp.Contract.GetMerkleRoot(&_OffRamp.CallOpts, sourceChainSelector, root)
}

func (_OffRamp *OffRampCaller) GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (OffRampSourceChainConfig, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getSourceChainConfig", sourceChainSelector)

	if err != nil {
		return *new(OffRampSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OffRampSourceChainConfig)).(*OffRampSourceChainConfig)

	return out0, err

}

func (_OffRamp *OffRampSession) GetSourceChainConfig(sourceChainSelector uint64) (OffRampSourceChainConfig, error) {
	return _OffRamp.Contract.GetSourceChainConfig(&_OffRamp.CallOpts, sourceChainSelector)
}

func (_OffRamp *OffRampCallerSession) GetSourceChainConfig(sourceChainSelector uint64) (OffRampSourceChainConfig, error) {
	return _OffRamp.Contract.GetSourceChainConfig(&_OffRamp.CallOpts, sourceChainSelector)
}

func (_OffRamp *OffRampCaller) GetStaticConfig(opts *bind.CallOpts) (OffRampStaticConfig, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(OffRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OffRampStaticConfig)).(*OffRampStaticConfig)

	return out0, err

}

func (_OffRamp *OffRampSession) GetStaticConfig() (OffRampStaticConfig, error) {
	return _OffRamp.Contract.GetStaticConfig(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetStaticConfig() (OffRampStaticConfig, error) {
	return _OffRamp.Contract.GetStaticConfig(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) LatestConfigDetails(opts *bind.CallOpts, ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "latestConfigDetails", ocrPluginType)

	if err != nil {
		return *new(MultiOCR3BaseOCRConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(MultiOCR3BaseOCRConfig)).(*MultiOCR3BaseOCRConfig)

	return out0, err

}

func (_OffRamp *OffRampSession) LatestConfigDetails(ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error) {
	return _OffRamp.Contract.LatestConfigDetails(&_OffRamp.CallOpts, ocrPluginType)
}

func (_OffRamp *OffRampCallerSession) LatestConfigDetails(ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error) {
	return _OffRamp.Contract.LatestConfigDetails(&_OffRamp.CallOpts, ocrPluginType)
}

func (_OffRamp *OffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRamp *OffRampSession) Owner() (common.Address, error) {
	return _OffRamp.Contract.Owner(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) Owner() (common.Address, error) {
	return _OffRamp.Contract.Owner(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OffRamp *OffRampSession) TypeAndVersion() (string, error) {
	return _OffRamp.Contract.TypeAndVersion(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) TypeAndVersion() (string, error) {
	return _OffRamp.Contract.TypeAndVersion(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_OffRamp *OffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffRamp.Contract.AcceptOwnership(&_OffRamp.TransactOpts)
}

func (_OffRamp *OffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffRamp.Contract.AcceptOwnership(&_OffRamp.TransactOpts)
}

func (_OffRamp *OffRampTransactor) ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "applySourceChainConfigUpdates", sourceChainConfigUpdates)
}

func (_OffRamp *OffRampSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _OffRamp.Contract.ApplySourceChainConfigUpdates(&_OffRamp.TransactOpts, sourceChainConfigUpdates)
}

func (_OffRamp *OffRampTransactorSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _OffRamp.Contract.ApplySourceChainConfigUpdates(&_OffRamp.TransactOpts, sourceChainConfigUpdates)
}

func (_OffRamp *OffRampTransactor) Commit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "commit", reportContext, report, rs, ss, rawVs)
}

func (_OffRamp *OffRampSession) Commit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.Commit(&_OffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_OffRamp *OffRampTransactorSession) Commit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.Commit(&_OffRamp.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_OffRamp *OffRampTransactor) Execute(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "execute", reportContext, report)
}

func (_OffRamp *OffRampSession) Execute(reportContext [3][32]byte, report []byte) (*types.Transaction, error) {
	return _OffRamp.Contract.Execute(&_OffRamp.TransactOpts, reportContext, report)
}

func (_OffRamp *OffRampTransactorSession) Execute(reportContext [3][32]byte, report []byte) (*types.Transaction, error) {
	return _OffRamp.Contract.Execute(&_OffRamp.TransactOpts, reportContext, report)
}

func (_OffRamp *OffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMRampMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "executeSingleMessage", message, offchainTokenData)
}

func (_OffRamp *OffRampSession) ExecuteSingleMessage(message InternalAny2EVMRampMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.ExecuteSingleMessage(&_OffRamp.TransactOpts, message, offchainTokenData)
}

func (_OffRamp *OffRampTransactorSession) ExecuteSingleMessage(message InternalAny2EVMRampMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.ExecuteSingleMessage(&_OffRamp.TransactOpts, message, offchainTokenData)
}

func (_OffRamp *OffRampTransactor) ManuallyExecute(opts *bind.TransactOpts, reports []InternalExecutionReportSingleChain, gasLimitOverrides [][]*big.Int) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "manuallyExecute", reports, gasLimitOverrides)
}

func (_OffRamp *OffRampSession) ManuallyExecute(reports []InternalExecutionReportSingleChain, gasLimitOverrides [][]*big.Int) (*types.Transaction, error) {
	return _OffRamp.Contract.ManuallyExecute(&_OffRamp.TransactOpts, reports, gasLimitOverrides)
}

func (_OffRamp *OffRampTransactorSession) ManuallyExecute(reports []InternalExecutionReportSingleChain, gasLimitOverrides [][]*big.Int) (*types.Transaction, error) {
	return _OffRamp.Contract.ManuallyExecute(&_OffRamp.TransactOpts, reports, gasLimitOverrides)
}

func (_OffRamp *OffRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OffRampDynamicConfig) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_OffRamp *OffRampSession) SetDynamicConfig(dynamicConfig OffRampDynamicConfig) (*types.Transaction, error) {
	return _OffRamp.Contract.SetDynamicConfig(&_OffRamp.TransactOpts, dynamicConfig)
}

func (_OffRamp *OffRampTransactorSession) SetDynamicConfig(dynamicConfig OffRampDynamicConfig) (*types.Transaction, error) {
	return _OffRamp.Contract.SetDynamicConfig(&_OffRamp.TransactOpts, dynamicConfig)
}

func (_OffRamp *OffRampTransactor) SetOCR3Configs(opts *bind.TransactOpts, ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "setOCR3Configs", ocrConfigArgs)
}

func (_OffRamp *OffRampSession) SetOCR3Configs(ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
	return _OffRamp.Contract.SetOCR3Configs(&_OffRamp.TransactOpts, ocrConfigArgs)
}

func (_OffRamp *OffRampTransactorSession) SetOCR3Configs(ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
	return _OffRamp.Contract.SetOCR3Configs(&_OffRamp.TransactOpts, ocrConfigArgs)
}

func (_OffRamp *OffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_OffRamp *OffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.TransferOwnership(&_OffRamp.TransactOpts, to)
}

func (_OffRamp *OffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.TransferOwnership(&_OffRamp.TransactOpts, to)
}

type OffRampAlreadyAttemptedIterator struct {
	Event *OffRampAlreadyAttempted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampAlreadyAttemptedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampAlreadyAttempted)
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
		it.Event = new(OffRampAlreadyAttempted)
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

func (it *OffRampAlreadyAttemptedIterator) Error() error {
	return it.fail
}

func (it *OffRampAlreadyAttemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampAlreadyAttempted struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	Raw                 types.Log
}

func (_OffRamp *OffRampFilterer) FilterAlreadyAttempted(opts *bind.FilterOpts) (*OffRampAlreadyAttemptedIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "AlreadyAttempted")
	if err != nil {
		return nil, err
	}
	return &OffRampAlreadyAttemptedIterator{contract: _OffRamp.contract, event: "AlreadyAttempted", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchAlreadyAttempted(opts *bind.WatchOpts, sink chan<- *OffRampAlreadyAttempted) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "AlreadyAttempted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampAlreadyAttempted)
				if err := _OffRamp.contract.UnpackLog(event, "AlreadyAttempted", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseAlreadyAttempted(log types.Log) (*OffRampAlreadyAttempted, error) {
	event := new(OffRampAlreadyAttempted)
	if err := _OffRamp.contract.UnpackLog(event, "AlreadyAttempted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampCommitReportAcceptedIterator struct {
	Event *OffRampCommitReportAccepted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampCommitReportAcceptedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampCommitReportAccepted)
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
		it.Event = new(OffRampCommitReportAccepted)
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

func (it *OffRampCommitReportAcceptedIterator) Error() error {
	return it.fail
}

func (it *OffRampCommitReportAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampCommitReportAccepted struct {
	Report OffRampCommitReport
	Raw    types.Log
}

func (_OffRamp *OffRampFilterer) FilterCommitReportAccepted(opts *bind.FilterOpts) (*OffRampCommitReportAcceptedIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "CommitReportAccepted")
	if err != nil {
		return nil, err
	}
	return &OffRampCommitReportAcceptedIterator{contract: _OffRamp.contract, event: "CommitReportAccepted", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchCommitReportAccepted(opts *bind.WatchOpts, sink chan<- *OffRampCommitReportAccepted) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "CommitReportAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampCommitReportAccepted)
				if err := _OffRamp.contract.UnpackLog(event, "CommitReportAccepted", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseCommitReportAccepted(log types.Log) (*OffRampCommitReportAccepted, error) {
	event := new(OffRampCommitReportAccepted)
	if err := _OffRamp.contract.UnpackLog(event, "CommitReportAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampConfigSetIterator struct {
	Event *OffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampConfigSet)
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
		it.Event = new(OffRampConfigSet)
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

func (it *OffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampConfigSet struct {
	OcrPluginType uint8
	ConfigDigest  [32]byte
	Signers       []common.Address
	Transmitters  []common.Address
	F             uint8
	Raw           types.Log
}

func (_OffRamp *OffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OffRampConfigSetIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampConfigSetIterator{contract: _OffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampConfigSet)
				if err := _OffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseConfigSet(log types.Log) (*OffRampConfigSet, error) {
	event := new(OffRampConfigSet)
	if err := _OffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampDynamicConfigSetIterator struct {
	Event *OffRampDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampDynamicConfigSet)
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
		it.Event = new(OffRampDynamicConfigSet)
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

func (it *OffRampDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampDynamicConfigSet struct {
	DynamicConfig OffRampDynamicConfig
	Raw           types.Log
}

func (_OffRamp *OffRampFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*OffRampDynamicConfigSetIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampDynamicConfigSetIterator{contract: _OffRamp.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampDynamicConfigSet)
				if err := _OffRamp.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseDynamicConfigSet(log types.Log) (*OffRampDynamicConfigSet, error) {
	event := new(OffRampDynamicConfigSet)
	if err := _OffRamp.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampExecutionStateChangedIterator struct {
	Event *OffRampExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampExecutionStateChanged)
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
		it.Event = new(OffRampExecutionStateChanged)
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

func (it *OffRampExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *OffRampExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampExecutionStateChanged struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	MessageId           [32]byte
	MessageHash         [32]byte
	State               uint8
	ReturnData          []byte
	GasUsed             *big.Int
	Raw                 types.Log
}

func (_OffRamp *OffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*OffRampExecutionStateChangedIterator, error) {

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

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &OffRampExecutionStateChangedIterator{contract: _OffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *OffRampExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampExecutionStateChanged)
				if err := _OffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseExecutionStateChanged(log types.Log) (*OffRampExecutionStateChanged, error) {
	event := new(OffRampExecutionStateChanged)
	if err := _OffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOwnershipTransferRequestedIterator struct {
	Event *OffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOwnershipTransferRequested)
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
		it.Event = new(OffRampOwnershipTransferRequested)
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

func (it *OffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OffRamp *OffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffRampOwnershipTransferRequestedIterator{contract: _OffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOwnershipTransferRequested)
				if err := _OffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*OffRampOwnershipTransferRequested, error) {
	event := new(OffRampOwnershipTransferRequested)
	if err := _OffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOwnershipTransferredIterator struct {
	Event *OffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOwnershipTransferred)
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
		it.Event = new(OffRampOwnershipTransferred)
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

func (it *OffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OffRamp *OffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffRampOwnershipTransferredIterator{contract: _OffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOwnershipTransferred)
				if err := _OffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseOwnershipTransferred(log types.Log) (*OffRampOwnershipTransferred, error) {
	event := new(OffRampOwnershipTransferred)
	if err := _OffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampRootRemovedIterator struct {
	Event *OffRampRootRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampRootRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampRootRemoved)
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
		it.Event = new(OffRampRootRemoved)
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

func (it *OffRampRootRemovedIterator) Error() error {
	return it.fail
}

func (it *OffRampRootRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampRootRemoved struct {
	Root [32]byte
	Raw  types.Log
}

func (_OffRamp *OffRampFilterer) FilterRootRemoved(opts *bind.FilterOpts) (*OffRampRootRemovedIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "RootRemoved")
	if err != nil {
		return nil, err
	}
	return &OffRampRootRemovedIterator{contract: _OffRamp.contract, event: "RootRemoved", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchRootRemoved(opts *bind.WatchOpts, sink chan<- *OffRampRootRemoved) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "RootRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampRootRemoved)
				if err := _OffRamp.contract.UnpackLog(event, "RootRemoved", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseRootRemoved(log types.Log) (*OffRampRootRemoved, error) {
	event := new(OffRampRootRemoved)
	if err := _OffRamp.contract.UnpackLog(event, "RootRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampSkippedAlreadyExecutedMessageIterator struct {
	Event *OffRampSkippedAlreadyExecutedMessage

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampSkippedAlreadyExecutedMessageIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampSkippedAlreadyExecutedMessage)
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
		it.Event = new(OffRampSkippedAlreadyExecutedMessage)
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

func (it *OffRampSkippedAlreadyExecutedMessageIterator) Error() error {
	return it.fail
}

func (it *OffRampSkippedAlreadyExecutedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampSkippedAlreadyExecutedMessage struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	Raw                 types.Log
}

func (_OffRamp *OffRampFilterer) FilterSkippedAlreadyExecutedMessage(opts *bind.FilterOpts) (*OffRampSkippedAlreadyExecutedMessageIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "SkippedAlreadyExecutedMessage")
	if err != nil {
		return nil, err
	}
	return &OffRampSkippedAlreadyExecutedMessageIterator{contract: _OffRamp.contract, event: "SkippedAlreadyExecutedMessage", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchSkippedAlreadyExecutedMessage(opts *bind.WatchOpts, sink chan<- *OffRampSkippedAlreadyExecutedMessage) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "SkippedAlreadyExecutedMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampSkippedAlreadyExecutedMessage)
				if err := _OffRamp.contract.UnpackLog(event, "SkippedAlreadyExecutedMessage", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseSkippedAlreadyExecutedMessage(log types.Log) (*OffRampSkippedAlreadyExecutedMessage, error) {
	event := new(OffRampSkippedAlreadyExecutedMessage)
	if err := _OffRamp.contract.UnpackLog(event, "SkippedAlreadyExecutedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampSkippedReportExecutionIterator struct {
	Event *OffRampSkippedReportExecution

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampSkippedReportExecutionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampSkippedReportExecution)
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
		it.Event = new(OffRampSkippedReportExecution)
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

func (it *OffRampSkippedReportExecutionIterator) Error() error {
	return it.fail
}

func (it *OffRampSkippedReportExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampSkippedReportExecution struct {
	SourceChainSelector uint64
	Raw                 types.Log
}

func (_OffRamp *OffRampFilterer) FilterSkippedReportExecution(opts *bind.FilterOpts) (*OffRampSkippedReportExecutionIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "SkippedReportExecution")
	if err != nil {
		return nil, err
	}
	return &OffRampSkippedReportExecutionIterator{contract: _OffRamp.contract, event: "SkippedReportExecution", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchSkippedReportExecution(opts *bind.WatchOpts, sink chan<- *OffRampSkippedReportExecution) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "SkippedReportExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampSkippedReportExecution)
				if err := _OffRamp.contract.UnpackLog(event, "SkippedReportExecution", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseSkippedReportExecution(log types.Log) (*OffRampSkippedReportExecution, error) {
	event := new(OffRampSkippedReportExecution)
	if err := _OffRamp.contract.UnpackLog(event, "SkippedReportExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampSourceChainConfigSetIterator struct {
	Event *OffRampSourceChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampSourceChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampSourceChainConfigSet)
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
		it.Event = new(OffRampSourceChainConfigSet)
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

func (it *OffRampSourceChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampSourceChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampSourceChainConfigSet struct {
	SourceChainSelector uint64
	SourceConfig        OffRampSourceChainConfig
	Raw                 types.Log
}

func (_OffRamp *OffRampFilterer) FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*OffRampSourceChainConfigSetIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OffRampSourceChainConfigSetIterator{contract: _OffRamp.contract, event: "SourceChainConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampSourceChainConfigSet)
				if err := _OffRamp.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseSourceChainConfigSet(log types.Log) (*OffRampSourceChainConfigSet, error) {
	event := new(OffRampSourceChainConfigSet)
	if err := _OffRamp.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampSourceChainSelectorAddedIterator struct {
	Event *OffRampSourceChainSelectorAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampSourceChainSelectorAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampSourceChainSelectorAdded)
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
		it.Event = new(OffRampSourceChainSelectorAdded)
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

func (it *OffRampSourceChainSelectorAddedIterator) Error() error {
	return it.fail
}

func (it *OffRampSourceChainSelectorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampSourceChainSelectorAdded struct {
	SourceChainSelector uint64
	Raw                 types.Log
}

func (_OffRamp *OffRampFilterer) FilterSourceChainSelectorAdded(opts *bind.FilterOpts) (*OffRampSourceChainSelectorAddedIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "SourceChainSelectorAdded")
	if err != nil {
		return nil, err
	}
	return &OffRampSourceChainSelectorAddedIterator{contract: _OffRamp.contract, event: "SourceChainSelectorAdded", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchSourceChainSelectorAdded(opts *bind.WatchOpts, sink chan<- *OffRampSourceChainSelectorAdded) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "SourceChainSelectorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampSourceChainSelectorAdded)
				if err := _OffRamp.contract.UnpackLog(event, "SourceChainSelectorAdded", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseSourceChainSelectorAdded(log types.Log) (*OffRampSourceChainSelectorAdded, error) {
	event := new(OffRampSourceChainSelectorAdded)
	if err := _OffRamp.contract.UnpackLog(event, "SourceChainSelectorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampStaticConfigSetIterator struct {
	Event *OffRampStaticConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampStaticConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampStaticConfigSet)
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
		it.Event = new(OffRampStaticConfigSet)
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

func (it *OffRampStaticConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampStaticConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampStaticConfigSet struct {
	StaticConfig OffRampStaticConfig
	Raw          types.Log
}

func (_OffRamp *OffRampFilterer) FilterStaticConfigSet(opts *bind.FilterOpts) (*OffRampStaticConfigSetIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampStaticConfigSetIterator{contract: _OffRamp.contract, event: "StaticConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampStaticConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampStaticConfigSet)
				if err := _OffRamp.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseStaticConfigSet(log types.Log) (*OffRampStaticConfigSet, error) {
	event := new(OffRampStaticConfigSet)
	if err := _OffRamp.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampTransmittedIterator struct {
	Event *OffRampTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampTransmitted)
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
		it.Event = new(OffRampTransmitted)
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

func (it *OffRampTransmittedIterator) Error() error {
	return it.fail
}

func (it *OffRampTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampTransmitted struct {
	OcrPluginType  uint8
	ConfigDigest   [32]byte
	SequenceNumber uint64
	Raw            types.Log
}

func (_OffRamp *OffRampFilterer) FilterTransmitted(opts *bind.FilterOpts, ocrPluginType []uint8) (*OffRampTransmittedIterator, error) {

	var ocrPluginTypeRule []interface{}
	for _, ocrPluginTypeItem := range ocrPluginType {
		ocrPluginTypeRule = append(ocrPluginTypeRule, ocrPluginTypeItem)
	}

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "Transmitted", ocrPluginTypeRule)
	if err != nil {
		return nil, err
	}
	return &OffRampTransmittedIterator{contract: _OffRamp.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OffRampTransmitted, ocrPluginType []uint8) (event.Subscription, error) {

	var ocrPluginTypeRule []interface{}
	for _, ocrPluginTypeItem := range ocrPluginType {
		ocrPluginTypeRule = append(ocrPluginTypeRule, ocrPluginTypeItem)
	}

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "Transmitted", ocrPluginTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampTransmitted)
				if err := _OffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseTransmitted(log types.Log) (*OffRampTransmitted, error) {
	event := new(OffRampTransmitted)
	if err := _OffRamp.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_OffRamp *OffRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _OffRamp.abi.Events["AlreadyAttempted"].ID:
		return _OffRamp.ParseAlreadyAttempted(log)
	case _OffRamp.abi.Events["CommitReportAccepted"].ID:
		return _OffRamp.ParseCommitReportAccepted(log)
	case _OffRamp.abi.Events["ConfigSet"].ID:
		return _OffRamp.ParseConfigSet(log)
	case _OffRamp.abi.Events["DynamicConfigSet"].ID:
		return _OffRamp.ParseDynamicConfigSet(log)
	case _OffRamp.abi.Events["ExecutionStateChanged"].ID:
		return _OffRamp.ParseExecutionStateChanged(log)
	case _OffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _OffRamp.ParseOwnershipTransferRequested(log)
	case _OffRamp.abi.Events["OwnershipTransferred"].ID:
		return _OffRamp.ParseOwnershipTransferred(log)
	case _OffRamp.abi.Events["RootRemoved"].ID:
		return _OffRamp.ParseRootRemoved(log)
	case _OffRamp.abi.Events["SkippedAlreadyExecutedMessage"].ID:
		return _OffRamp.ParseSkippedAlreadyExecutedMessage(log)
	case _OffRamp.abi.Events["SkippedReportExecution"].ID:
		return _OffRamp.ParseSkippedReportExecution(log)
	case _OffRamp.abi.Events["SourceChainConfigSet"].ID:
		return _OffRamp.ParseSourceChainConfigSet(log)
	case _OffRamp.abi.Events["SourceChainSelectorAdded"].ID:
		return _OffRamp.ParseSourceChainSelectorAdded(log)
	case _OffRamp.abi.Events["StaticConfigSet"].ID:
		return _OffRamp.ParseStaticConfigSet(log)
	case _OffRamp.abi.Events["Transmitted"].ID:
		return _OffRamp.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (OffRampAlreadyAttempted) Topic() common.Hash {
	return common.HexToHash("0x3ef2a99c550a751d4b0b261268f05a803dfb049ab43616a1ffb388f61fe65120")
}

func (OffRampCommitReportAccepted) Topic() common.Hash {
	return common.HexToHash("0x23bc80217a08968cec0790cd045b396fa7eea0a21af469e603329940b883d86d")
}

func (OffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f547")
}

func (OffRampDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0xa55bd56595c45f517e5967a3067f3dca684445a3080e7c04a4e0d5a40cda627d")
}

func (OffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x05665fe9ad095383d018353f4cbcba77e84db27dd215081bbf7cdf9ae6fbe48b")
}

func (OffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OffRampRootRemoved) Topic() common.Hash {
	return common.HexToHash("0x202f1139a3e334b6056064c0e9b19fd07e44a88d8f6e5ded571b24cf8c371f12")
}

func (OffRampSkippedAlreadyExecutedMessage) Topic() common.Hash {
	return common.HexToHash("0x3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c")
}

func (OffRampSkippedReportExecution) Topic() common.Hash {
	return common.HexToHash("0xaab522ed53d887e56ed53dd37398a01aeef6a58e0fa77c2173beb9512d894933")
}

func (OffRampSourceChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x49f51971edd25182e97182d6ea372a0488ce2ab639f6a3a7ab4df0d2636fe56b")
}

func (OffRampSourceChainSelectorAdded) Topic() common.Hash {
	return common.HexToHash("0xf4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb9")
}

func (OffRampStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0x683eb52ee924eb817377cfa8f41f238f4bb7a877da5267869dfffbad85f564d8")
}

func (OffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0x198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef0")
}

func (_OffRamp *OffRamp) Address() common.Address {
	return _OffRamp.address
}

type OffRampInterface interface {
	CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error

	GetDynamicConfig(opts *bind.CallOpts) (OffRampDynamicConfig, error)

	GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error)

	GetLatestPriceSequenceNumber(opts *bind.CallOpts) (uint64, error)

	GetMerkleRoot(opts *bind.CallOpts, sourceChainSelector uint64, root [32]byte) (*big.Int, error)

	GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (OffRampSourceChainConfig, error)

	GetStaticConfig(opts *bind.CallOpts) (OffRampStaticConfig, error)

	LatestConfigDetails(opts *bind.CallOpts, ocrPluginType uint8) (MultiOCR3BaseOCRConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error)

	Commit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMRampMessage, offchainTokenData [][]byte) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, reports []InternalExecutionReportSingleChain, gasLimitOverrides [][]*big.Int) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OffRampDynamicConfig) (*types.Transaction, error)

	SetOCR3Configs(opts *bind.TransactOpts, ocrConfigArgs []MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAlreadyAttempted(opts *bind.FilterOpts) (*OffRampAlreadyAttemptedIterator, error)

	WatchAlreadyAttempted(opts *bind.WatchOpts, sink chan<- *OffRampAlreadyAttempted) (event.Subscription, error)

	ParseAlreadyAttempted(log types.Log) (*OffRampAlreadyAttempted, error)

	FilterCommitReportAccepted(opts *bind.FilterOpts) (*OffRampCommitReportAcceptedIterator, error)

	WatchCommitReportAccepted(opts *bind.WatchOpts, sink chan<- *OffRampCommitReportAccepted) (event.Subscription, error)

	ParseCommitReportAccepted(log types.Log) (*OffRampCommitReportAccepted, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OffRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OffRampConfigSet, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*OffRampDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*OffRampDynamicConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*OffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *OffRampExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*OffRampExecutionStateChanged, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OffRampOwnershipTransferred, error)

	FilterRootRemoved(opts *bind.FilterOpts) (*OffRampRootRemovedIterator, error)

	WatchRootRemoved(opts *bind.WatchOpts, sink chan<- *OffRampRootRemoved) (event.Subscription, error)

	ParseRootRemoved(log types.Log) (*OffRampRootRemoved, error)

	FilterSkippedAlreadyExecutedMessage(opts *bind.FilterOpts) (*OffRampSkippedAlreadyExecutedMessageIterator, error)

	WatchSkippedAlreadyExecutedMessage(opts *bind.WatchOpts, sink chan<- *OffRampSkippedAlreadyExecutedMessage) (event.Subscription, error)

	ParseSkippedAlreadyExecutedMessage(log types.Log) (*OffRampSkippedAlreadyExecutedMessage, error)

	FilterSkippedReportExecution(opts *bind.FilterOpts) (*OffRampSkippedReportExecutionIterator, error)

	WatchSkippedReportExecution(opts *bind.WatchOpts, sink chan<- *OffRampSkippedReportExecution) (event.Subscription, error)

	ParseSkippedReportExecution(log types.Log) (*OffRampSkippedReportExecution, error)

	FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*OffRampSourceChainConfigSetIterator, error)

	WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error)

	ParseSourceChainConfigSet(log types.Log) (*OffRampSourceChainConfigSet, error)

	FilterSourceChainSelectorAdded(opts *bind.FilterOpts) (*OffRampSourceChainSelectorAddedIterator, error)

	WatchSourceChainSelectorAdded(opts *bind.WatchOpts, sink chan<- *OffRampSourceChainSelectorAdded) (event.Subscription, error)

	ParseSourceChainSelectorAdded(log types.Log) (*OffRampSourceChainSelectorAdded, error)

	FilterStaticConfigSet(opts *bind.FilterOpts) (*OffRampStaticConfigSetIterator, error)

	WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampStaticConfigSet) (event.Subscription, error)

	ParseStaticConfigSet(log types.Log) (*OffRampStaticConfigSet, error)

	FilterTransmitted(opts *bind.FilterOpts, ocrPluginType []uint8) (*OffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OffRampTransmitted, ocrPluginType []uint8) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*OffRampTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
