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
	IsEnabled bool
	MinSeqNr  uint64
	OnRamp    []byte
}

type EVM2EVMMultiOffRampSourceChainConfigArgs struct {
	SourceChainSelector uint64
	IsEnabled           bool
	OnRamp              []byte
}

type EVM2EVMMultiOffRampStaticConfig struct {
	ChainSelector      uint64
	RmnProxy           common.Address
	TokenAdminRegistry common.Address
	NonceManager       common.Address
}

type EVM2EVMMultiOffRampUnblessedRoot struct {
	SourceChainSelector uint64
	MerkleRoot          [32]byte
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumMultiOCR3Base.InvalidConfigErrorType\",\"name\":\"errorType\",\"type\":\"uint8\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.Interval\",\"name\":\"interval\",\"type\":\"tuple\"}],\"name\":\"InvalidInterval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"InvalidManualExecutionGasLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"messageDestChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidMessageDestChainSelector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"InvalidNewState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRoot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidStaticConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeavesCannotBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionGasLimitMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"errorReason\",\"type\":\"bytes\"}],\"name\":\"MessageValidationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notPool\",\"type\":\"address\"}],\"name\":\"NotACompatiblePool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"RootAlreadyCommitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleCommitReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"}],\"name\":\"StaticConfigCannotBeChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"TokenDataMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"TokenHandlingError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedTokenData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroChainSelectorNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint224\",\"name\":\"usdPerToken\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint224\",\"name\":\"usdPerUnitGas\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.GasPriceUpdate[]\",\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.PriceUpdates\",\"name\":\"priceUpdates\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.Interval\",\"name\":\"interval\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.MerkleRoot[]\",\"name\":\"merkleRoots\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.CommitReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"CommitReportAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"DynamicConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"RootRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"SkippedAlreadyExecutedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfig\",\"name\":\"sourceConfig\",\"type\":\"tuple\"}],\"name\":\"SourceChainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainSelectorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"}],\"name\":\"StaticConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\"}],\"name\":\"applySourceChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.RampTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.Any2EVMRampMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPriceSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"getSourceChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"isBlessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"}],\"name\":\"latestConfigDetails\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"n\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\"}],\"internalType\":\"structMultiOCR3Base.ConfigInfo\",\"name\":\"configInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"}],\"internalType\":\"structMultiOCR3Base.OCRConfig\",\"name\":\"ocrConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.RampTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.Any2EVMRampMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReportSingleChain[]\",\"name\":\"reports\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"gasLimitOverrides\",\"type\":\"uint256[][]\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.UnblessedRoot[]\",\"name\":\"rootToReset\",\"type\":\"tuple[]\"}],\"name\":\"resetUnblessedRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"}],\"internalType\":\"structMultiOCR3Base.OCRConfigArgs[]\",\"name\":\"ocrConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"setOCR3Configs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162006cc238038062006cc28339810160408190526200003591620008e2565b33806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf81620001fa565b5050466080525060208301516001600160a01b03161580620000ec575060408301516001600160a01b0316155b8062000103575060608301516001600160a01b0316155b1562000122576040516342bcdf7f60e11b815260040160405180910390fd5b82516001600160401b03166000036200014e5760405163c656089560e01b815260040160405180910390fd5b82516001600160401b0390811660a052602080850180516001600160a01b0390811660c05260408088018051831660e0526060808a01805185166101005283518b519098168852945184169587019590955251821690850152905116908201527f683eb52ee924eb817377cfa8f41f238f4bb7a877da5267869dfffbad85f564d89060800160405180910390a1620001e682620002a5565b620001f181620003fc565b50505062000c57565b336001600160a01b03821603620002545760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60a08101516001600160a01b03161580620002c8575080516001600160a01b0316155b15620002e7576040516342bcdf7f60e11b815260040160405180910390fd5b80516004805460208085018051604080880180516060808b0180516001600160a01b039b8c166001600160c01b0319909a168a17600160a01b63ffffffff98891602176001600160c01b0316600160c01b948816949094026001600160e01b031693909317600160e01b93871693909302929092179098556080808b018051600580546001600160a01b0319908116928e1692909217905560a0808e01805160068054909416908f161790925586519a8b5297518716988a0198909852925185169388019390935251909216958501959095525185169383019390935251909216908201527f0da37fd00459f4f5f0b8210d31525e4910ae674b8bab34b561d146bb45773a4c9060c00160405180910390a150565b60005b81518110156200064e57600082828151811062000420576200042062000a20565b60200260200101519050600081600001519050806001600160401b03166000036200045e5760405163c656089560e01b815260040160405180910390fd5b6001600160401b03811660009081526007602052604081206001810180549192916200048a9062000a36565b80601f0160208091040260200160405190810160405280929190818152602001828054620004b89062000a36565b8015620005095780601f10620004dd5761010080835404028352916020019162000509565b820191906000526020600020905b815481529060010190602001808311620004eb57829003601f168201915b505050505090506000846040015190508151600003620005ac57805160000362000546576040516342bcdf7f60e11b815260040160405180910390fd5b6001830162000556828262000ac7565b508254610100600160481b0319166101001783556040516001600160401b03851681527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb99060200160405180910390a1620005e7565b8080519060200120828051906020012014620005e75760405163c39a620560e01b81526001600160401b038516600482015260240162000083565b6020850151835460ff19169015151783556040516001600160401b038516907f4f49973170c548fddd4a48341b75e131818913f38f44d47af57e8735eee588ba906200063590869062000b93565b60405180910390a25050505050806001019050620003ff565b5050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156200068d576200068d62000652565b60405290565b604051608081016001600160401b03811182821017156200068d576200068d62000652565b60405160c081016001600160401b03811182821017156200068d576200068d62000652565b604051601f8201601f191681016001600160401b038111828210171562000708576200070862000652565b604052919050565b80516001600160401b03811681146200072857600080fd5b919050565b80516001600160a01b03811681146200072857600080fd5b805163ffffffff811681146200072857600080fd5b6000601f83601f8401126200076e57600080fd5b825160206001600160401b03808311156200078d576200078d62000652565b8260051b6200079e838201620006dd565b9384528681018301938381019089861115620007b957600080fd5b84890192505b85831015620008d557825184811115620007d95760008081fd5b89016060601f19828d038101821315620007f35760008081fd5b620007fd62000668565b6200080a89850162000710565b81526040808501518015158114620008225760008081fd5b828b01529284015192888411156200083a5760008081fd5b83850194508e603f8601126200085257600093508384fd5b898501519350888411156200086b576200086b62000652565b6200087c8a848e87011601620006dd565b92508383528e81858701011115620008945760008081fd5b60005b84811015620008b4578581018201518482018c01528a0162000897565b5060009383018a0193909352918201528352509184019190840190620007bf565b9998505050505050505050565b6000806000838503610160811215620008fa57600080fd5b60808112156200090957600080fd5b6200091362000693565b6200091e8662000710565b81526200092e602087016200072d565b602082015262000941604087016200072d565b604082015262000954606087016200072d565b6060820152935060c0607f19820112156200096e57600080fd5b5062000979620006b8565b62000987608086016200072d565b81526200099760a0860162000745565b6020820152620009aa60c0860162000745565b6040820152620009bd60e0860162000745565b6060820152620009d161010086016200072d565b6080820152620009e561012086016200072d565b60a08201526101408501519092506001600160401b0381111562000a0857600080fd5b62000a16868287016200075a565b9150509250925092565b634e487b7160e01b600052603260045260246000fd5b600181811c9082168062000a4b57607f821691505b60208210810362000a6c57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111562000ac2576000816000526020600020601f850160051c8101602086101562000a9d5750805b601f850160051c820191505b8181101562000abe5782815560010162000aa9565b5050505b505050565b81516001600160401b0381111562000ae35762000ae362000652565b62000afb8162000af4845462000a36565b8462000a72565b602080601f83116001811462000b33576000841562000b1a5750858301515b600019600386901b1c1916600185901b17855562000abe565b600085815260208120601f198616915b8281101562000b645788860151825594840194600190910190840162000b43565b508582101562000b835787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6020808252825460ff811615158383015260081c6001600160401b0316604083015260608083015260018084018054600093929190849062000bd58162000a36565b80608089015260a0600183166000811462000bf9576001811462000c165762000c48565b60ff19841660a08b015260a083151560051b8b0101945062000c48565b85600052602060002060005b8481101562000c3f5781548c820185015290880190890162000c22565b8b0160a0019550505b50929998505050505050505050565b60805160a05160c05160e05161010051615ffc62000cc66000396000818161023e0152612c8b01526000818161020f0152612fa20152600081816101e001528181611488015261153f0152600081816101b00152612805015260008181611cd70152611d230152615ffc6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806385572ffb116100d8578063d2a15d351161008c578063f2fde38b11610066578063f2fde38b1461059c578063f716f99f146105af578063ff888fb1146105c257600080fd5b8063d2a15d3514610556578063e9d68a8e14610569578063ece670b61461058957600080fd5b8063a12a9870116100bd578063a12a9870146104de578063c673e584146104f1578063ccd37ba31461051157600080fd5b806385572ffb146104b55780638da5cb5b146104c357600080fd5b8063403b2d631161012f5780637437ff9f116101145780637437ff9f1461038557806379ba50971461049a5780637d4eef60146104a257600080fd5b8063403b2d63146103525780635e36480c1461036557600080fd5b80632d04ab76116101605780632d04ab761461030e578063311cd513146103235780633f4b04aa1461033657600080fd5b806306285c691461017c578063181f5a77146102c5575b600080fd5b61026e60408051608081018252600080825260208201819052918101829052606081019190915260405180608001604052807f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b6040516102bc9190815167ffffffffffffffff1681526020808301516001600160a01b0390811691830191909152604080840151821690830152606092830151169181019190915260800190565b60405180910390f35b6103016040518060400160405280601d81526020017f45564d3245564d4d756c74694f666652616d7020312e362e302d64657600000081525081565b6040516102bc9190613fae565b61032161031c366004614059565b6105e5565b005b61032161033136600461410c565b6109ab565b600a5460405167ffffffffffffffff90911681526020016102bc565b610321610360366004614295565b610a14565b610378610373366004614334565b610a28565b6040516102bc9190614391565b6104316040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c0810182526004546001600160a01b03808216835263ffffffff74010000000000000000000000000000000000000000830481166020850152600160c01b8304811694840194909452600160e01b90910490921660608201526005548216608082015260065490911660a082015290565b6040516102bc9190600060c0820190506001600160a01b03808451168352602084015163ffffffff808216602086015280604087015116604086015280606087015116606086015250508060808501511660808401528060a08501511660a08401525092915050565b610321610a7e565b6103216104b036600461497b565b610b3c565b610321610177366004614aa6565b6000546040516001600160a01b0390911681526020016102bc565b6103216104ec366004614afa565b610cdc565b6105046104ff366004614c07565b610ced565b6040516102bc9190614c67565b61054861051f366004614cdc565b67ffffffffffffffff919091166000908152600960209081526040808320938352929052205490565b6040519081526020016102bc565b610321610564366004614d06565b610e4b565b61057c610577366004614d7b565b610f05565b6040516102bc9190614d96565b610321610597366004614dd1565b610fee565b6103216105aa366004614e36565b6113df565b6103216105bd366004614ebb565b6113f0565b6105d56105d0366004614ff9565b611432565b60405190151581526020016102bc565b60006105f387890189615182565b8051515190915015158061060c57508051602001515115155b1561070c57600a5460208a01359067ffffffffffffffff808316911610156106cb57600a805467ffffffffffffffff191667ffffffffffffffff831617905560065482516040517f3937306f0000000000000000000000000000000000000000000000000000000081526001600160a01b0390921691633937306f91610694916004016153c0565b600060405180830381600087803b1580156106ae57600080fd5b505af11580156106c2573d6000803e3d6000fd5b5050505061070a565b81602001515160000361070a576040517f2261116700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b60005b8160200151518110156108f457600082602001518281518110610734576107346152ed565b60200260200101519050600081600001519050610750816114f3565b600061075b826115f5565b602084015151815491925067ffffffffffffffff9081166101009092041614158061079d575060208084015190810151905167ffffffffffffffff9182169116115b156107e657825160208401516040517feefb0cac0000000000000000000000000000000000000000000000000000000081526107dd9291906004016153d3565b60405180910390fd5b604083015180610822576040517f504570e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835167ffffffffffffffff166000908152600960209081526040808320848452909152902054156108955783516040517f32cf0cbf00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602481018290526044016107dd565b60208085015101516108a890600161541e565b825468ffffffffffffffff00191661010067ffffffffffffffff92831602179092559251166000908152600960209081526040808320948352939052919091204290555060010161070f565b507f3a3950e13dd607cc37980db0ef14266c40d2bba9c01b2e44bfe549808883095d816040516109249190615446565b60405180910390a16109a060008a8a8a8a8a8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c9182918501908490808284376000920191909152508b9250611655915050565b505050505050505050565b6109eb6109ba828401846154e3565b60408051600080825260208201909252906109e5565b60608152602001906001900390816109d05790505b506119cc565b604080516000808252602082019092529050610a0e600185858585866000611655565b50505050565b610a1c611a7c565b610a2581611ad8565b50565b6000610a3660016004615518565b6002610a43608085615541565b67ffffffffffffffff16610a579190615568565b610a618585611c8d565b901c166003811115610a7557610a75614367565b90505b92915050565b6001546001600160a01b03163314610ad85760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016107dd565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610b44611cd4565b815181518114610b80576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b81811015610ccc576000848281518110610b9f57610b9f6152ed565b60200260200101519050600081602001515190506000858481518110610bc757610bc76152ed565b6020026020010151905080518214610c0b576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b82811015610cbd576000828281518110610c2a57610c2a6152ed565b6020026020010151905080600014610cb45784602001518281518110610c5257610c526152ed565b602002602001015160800151811015610cb45784516040517fc8e9605100000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260248101839052604481018290526064016107dd565b50600101610c0e565b50505050806001019050610b83565b50610cd783836119cc565b505050565b610ce4611a7c565b610a2581611d55565b610d306040805160e081019091526000606082018181526080830182905260a0830182905260c08301919091528190815260200160608152602001606081525090565b60ff808316600090815260026020818152604092839020835160e081018552815460608201908152600183015480881660808401526101008104881660a0840152620100009004909616151560c082015294855291820180548451818402810184019095528085529293858301939092830182828015610dd957602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610dbb575b5050505050815260200160038201805480602002602001604051908101604052809291908181526020018280548015610e3b57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e1d575b5050505050815250509050919050565b610e53611a7c565b60005b81811015610cd7576000838383818110610e7257610e726152ed565b905060400201803603810190610e88919061557f565b9050610e978160200151611432565b610efc57805167ffffffffffffffff1660009081526009602090815260408083208285018051855290835281842093909355915191519182527f202f1139a3e334b6056064c0e9b19fd07e44a88d8f6e5ded571b24cf8c371f12910160405180910390a15b50600101610e56565b60408051606080820183526000808352602080840182905283850183905267ffffffffffffffff8681168352600782529185902085519384018652805460ff811615158552610100900490921690830152600181018054939492939192840191610f6e906155b8565b80601f0160208091040260200160405190810160405280929190818152602001828054610f9a906155b8565b8015610e3b5780601f10610fbc57610100808354040283529160200191610e3b565b820191906000526020600020905b815481529060010190602001808311610fca57505050919092525091949350505050565b333014611027576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160008082526020820190925281611064565b604080518082019091526000808252602082015281526020019060019003908161103d5790505b50905060006110776101208601866155f2565b905011156110c8576110c56110906101208601866155f2565b61109d60a088018861563c565b6110ae6101008a0160e08b01614e36565b6110be60408b0160208c01614d7b565b8989611fe1565b90505b6040805160a081018252853581526000916020808301916110ee91908901908901614d7b565b67ffffffffffffffff16815260200161110a60a088018861563c565b60405160200161111b929190615683565b60408051601f19818403018152919052815260200161113d60c088018861563c565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506020018390526005549091506001600160a01b03168015611267576040517f08d450a10000000000000000000000000000000000000000000000000000000081526001600160a01b038216906308d450a1906111d4908590600401615754565b600060405180830381600087803b1580156111ee57600080fd5b505af19250505080156111ff575060015b611267573d80801561122d576040519150601f19603f3d011682016040523d82523d6000602084013e611232565b606091505b50806040517f09c253250000000000000000000000000000000000000000000000000000000081526004016107dd9190613fae565b61127460c087018761563c565b15905080156112865750610100860135155b806112a9575061129d610100870160e08801614e36565b6001600160a01b03163b155b806112f557506112f37f85572ffb000000000000000000000000000000000000000000000000000000006112e4610100890160e08a01614e36565b6001600160a01b0316906120d5565b155b1561130257505050505050565b60045460009081906001600160a01b0316633cf97983856113886101008c018035906113319060e08f01614e36565b6040518563ffffffff1660e01b81526004016113509493929190615767565b6000604051808303816000875af115801561136f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261139791908101906157a3565b5091509150816113d557806040517f0a8d6e8c0000000000000000000000000000000000000000000000000000000081526004016107dd9190613fae565b5050505050505050565b6113e7611a7c565b610a25816120f1565b6113f8611a7c565b60005b815181101561142e57611426828281518110611419576114196152ed565b60200260200101516121a7565b6001016113fb565b5050565b6040805180820182523081526020810183815291517f4d61677100000000000000000000000000000000000000000000000000000000815290516001600160a01b039081166004830152915160248201526000917f00000000000000000000000000000000000000000000000000000000000000001690634d61677190604401602060405180830381865afa1580156114cf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a789190615839565b6040517f2cbc26bb000000000000000000000000000000000000000000000000000000008152608082901b77ffffffffffffffff000000000000000000000000000000001660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632cbc26bb90602401602060405180830381865afa15801561158e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115b29190615839565b15610a25576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016107dd565b67ffffffffffffffff81166000908152600760205260408120805460ff16610a78576040517fed053c5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024016107dd565b60ff878116600090815260026020908152604080832081516080810183528154815260019091015480861693820193909352610100830485169181019190915262010000909104909216151560608301528735906116b48760a4615856565b90508260600151156116fc5784516116cd906020615568565b86516116da906020615568565b6116e59060a0615856565b6116ef9190615856565b6116f99082615856565b90505b36811461173e576040517f8e1192e1000000000000000000000000000000000000000000000000000000008152600481018290523660248201526044016107dd565b50815181146117865781516040517f93df584c0000000000000000000000000000000000000000000000000000000081526004810191909152602481018290526044016107dd565b61178e611cd4565b60ff808a16600090815260036020908152604080832033845282528083208151808301909252805480861683529394919390928401916101009091041660028111156117dc576117dc614367565b60028111156117ed576117ed614367565b905250905060028160200151600281111561180a5761180a614367565b14801561185e5750600260008b60ff1660ff168152602001908152602001600020600301816000015160ff1681548110611846576118466152ed565b6000918252602090912001546001600160a01b031633145b611894576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b508160600151156119765760208201516118af906001615869565b60ff168551146118eb576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8351855114611926576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008787604051611938929190615882565b60405190819003812061194f918b90602001615892565b6040516020818303038152906040528051906020012090506119748a828888886124eb565b505b6040805182815260208a81013567ffffffffffffffff169082015260ff8b16917f198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef0910160405180910390a2505050505050505050565b8151600003611a06576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160408051600080825260208201909252911591905b8451811015611a7557611a6d858281518110611a3b57611a3b6152ed565b602002602001015184611a6757858381518110611a5a57611a5a6152ed565b60200260200101516126f8565b836126f8565b600101611a1d565b5050505050565b6000546001600160a01b03163314611ad65760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016107dd565b565b60a08101516001600160a01b03161580611afa575080516001600160a01b0316155b15611b31576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516004805460208085018051604080880180516060808b0180516001600160a01b039b8c167fffffffffffffffff000000000000000000000000000000000000000000000000909a168a177401000000000000000000000000000000000000000063ffffffff988916021777ffffffffffffffffffffffffffffffffffffffffffffffff16600160c01b948816949094026001600160e01b031693909317600160e01b93871693909302929092179098556080808b0180516005805473ffffffffffffffffffffffffffffffffffffffff19908116928e1692909217905560a0808e01805160068054909416908f161790925586519a8b5297518716988a0198909852925185169388019390935251909216958501959095525185169383019390935251909216908201527f0da37fd00459f4f5f0b8210d31525e4910ae674b8bab34b561d146bb45773a4c9060c00160405180910390a150565b67ffffffffffffffff8216600090815260086020526040812081611cb26080856158a6565b67ffffffffffffffff1681526020810191909152604001600020549392505050565b467f000000000000000000000000000000000000000000000000000000000000000014611ad6576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201524660248201526044016107dd565b60005b815181101561142e576000828281518110611d7557611d756152ed565b602002602001015190506000816000015190508067ffffffffffffffff16600003611dcc576040517fc656089500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff81166000908152600760205260408120600181018054919291611df7906155b8565b80601f0160208091040260200160405190810160405280929190818152602001828054611e23906155b8565b8015611e705780601f10611e4557610100808354040283529160200191611e70565b820191906000526020600020905b815481529060010190602001808311611e5357829003601f168201915b505050505090506000846040015190508151600003611f29578051600003611ec4576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018301611ed28282615915565b50825468ffffffffffffffff00191661010017835560405167ffffffffffffffff851681527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb99060200160405180910390a1611f7c565b8080519060200120828051906020012014611f7c576040517fc39a620500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526024016107dd565b6020850151835460ff191690151517835560405167ffffffffffffffff8516907f4f49973170c548fddd4a48341b75e131818913f38f44d47af57e8735eee588ba90611fc99086906159d5565b60405180910390a25050505050806001019050611d58565b60608767ffffffffffffffff811115611ffc57611ffc614160565b60405190808252806020026020018201604052801561204157816020015b604080518082019091526000808252602082015281526020019060019003908161201a5790505b50905060005b888110156120c8576120a38a8a83818110612064576120646152ed565b90506020028101906120769190615a91565b8989898989898881811061208c5761208c6152ed565b905060200281019061209e919061563c565b612f04565b8282815181106120b5576120b56152ed565b6020908102919091010152600101612047565b5098975050505050505050565b60006120e0836133ff565b8015610a755750610a758383613463565b336001600160a01b038216036121495760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016107dd565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b806040015160ff166000036121d2576000604051631b3fab5160e11b81526004016107dd9190615acf565b60208082015160ff8082166000908152600290935260408320600181015492939092839216900361223f57606084015160018201805491151562010000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff909216919091179055612294565b6060840151600182015460ff6201000090910416151590151514612294576040517f87f6037c00000000000000000000000000000000000000000000000000000000815260ff841660048201526024016107dd565b60a08401518051601f60ff821611156122c3576001604051631b3fab5160e11b81526004016107dd9190615acf565b612329858560030180548060200260200160405190810160405280929190818152602001828054801561231f57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612301575b505050505061351d565b85606001511561245857612397858560020180548060200260200160405190810160405280929190818152602001828054801561231f576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161230157505050505061351d565b608086015180516123b19060028701906020840190613eb8565b5080516001850180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1661010060ff841690810291909117909155601f1015612411576002604051631b3fab5160e11b81526004016107dd9190615acf565b6040880151612421906003615ae9565b60ff168160ff1611612449576003604051631b3fab5160e11b81526004016107dd9190615acf565b61245587836001613586565b50505b61246485836002613586565b81516124799060038601906020850190613eb8565b5060408681015160018501805460ff191660ff8316179055875180865560a089015192517fab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f547936124d2938a939260028b01929190615b05565b60405180910390a16124e385613706565b505050505050565b6124f3613f2a565b835160005b818110156113d5576000600188868460208110612517576125176152ed565b61252491901a601b615869565b898581518110612536576125366152ed565b6020026020010151898681518110612550576125506152ed565b60200260200101516040516000815260200160405260405161258e949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156125b0573d6000803e3d6000fd5b505060408051601f1981015160ff808e166000908152600360209081528582206001600160a01b0385168352815285822085870190965285548084168652939750909550929392840191610100900416600281111561261157612611614367565b600281111561262257612622614367565b905250905060018160200151600281111561263f5761263f614367565b14612676576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051859060ff16601f811061268d5761268d6152ed565b6020020151156126c9576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600185826000015160ff16601f81106126e4576126e46152ed565b9115156020909202015250506001016124f8565b8151612703816114f3565b600061270e826115f5565b6020850151519091506000819003612751576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b846040015151811461278f576040517f57e0e08300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff8111156127aa576127aa614160565b6040519080825280602002602001820160405280156127d3578160200160208202803683370190505b50905060005b82811015612948576000876020015182815181106127f9576127f96152ed565b602002602001015190507f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681600001516040015167ffffffffffffffff161461288c57805160409081015190517f38432a2200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016107dd565b6129228186600101805461289f906155b8565b80601f01602080910402602001604051908101604052809291908181526020018280546128cb906155b8565b80156129185780601f106128ed57610100808354040283529160200191612918565b820191906000526020600020905b8154815290600101906020018083116128fb57829003601f168201915b5050505050613722565b838381518110612934576129346152ed565b6020908102919091010152506001016127d9565b50600061295f858389606001518a60800151613844565b9050806000036129a7576040517f7dd17a7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff861660048201526024016107dd565b8551151560005b848110156109a0576000896020015182815181106129ce576129ce6152ed565b6020026020010151905060006129ec89836000015160600151610a28565b90506002816003811115612a0257612a02614367565b03612a59578151606001516040805167ffffffffffffffff808d16825290921660208301527f3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c910160405180910390a15050612efc565b6000816003811115612a6d57612a6d614367565b1480612a8a57506003816003811115612a8857612a88614367565b145b612adb578151606001516040517f25507e7f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808c16600483015290911660248201526044016107dd565b8315612bbc5760045460009074010000000000000000000000000000000000000000900463ffffffff16612b0f8742615518565b1190508080612b2f57506003826003811115612b2d57612b2d614367565b145b612b71576040517fa9cfc86200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b1660048201526024016107dd565b8a8481518110612b8357612b836152ed565b6020026020010151600014612bb6578a8481518110612ba457612ba46152ed565b60200260200101518360800181815250505b50612c22565b6000816003811115612bd057612bd0614367565b14612c22578151606001516040517f3ef2a99c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808c16600483015290911660248201526044016107dd565b81516080015167ffffffffffffffff1615612d10576000816003811115612c4b57612c4b614367565b03612d105781516080015160208301516040517fe0e03cae0000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169263e0e03cae92612cc2928e929190600401615bb7565b6020604051808303816000875af1158015612ce1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d059190615839565b612d10575050612efc565b60008b604001518481518110612d2857612d286152ed565b6020026020010151905080518360a001515114612d8c578251606001516040517f1cfe6d8b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808d16600483015290911660248201526044016107dd565b612da08a846000015160600151600161389a565b600080612dad8584613942565b91509150612dc48c8660000151606001518461389a565b8615612e34576003826003811115612dde57612dde614367565b03612e34576000846003811115612df757612df7614367565b14612e34578451516040517f2b11b8d90000000000000000000000000000000000000000000000000000000081526107dd91908390600401615be4565b6002826003811115612e4857612e48614367565b14612ea2576003826003811115612e6157612e61614367565b14612ea2578451606001516040517f926c5a3e0000000000000000000000000000000000000000000000000000000081526107dd918e918590600401615bfd565b8451805160609091015160405167ffffffffffffffff918216918f16907f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df290612eee9087908790615c23565b60405180910390a450505050505b6001016129ae565b60408051808201909152600080825260208201526000612f64612f2a60208b018b61563c565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613a0c92505050565b6040517fbbe4f6db0000000000000000000000000000000000000000000000000000000081526001600160a01b0380831660048301529192506000917f0000000000000000000000000000000000000000000000000000000000000000169063bbe4f6db90602401602060405180830381865afa158015612fe9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061300d9190615c43565b90506001600160a01b038116158061305557506130536001600160a01b0382167faff2afbf000000000000000000000000000000000000000000000000000000006120d5565b155b15613097576040517fae9b4ce90000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016107dd565b6000806132746040518061010001604052808d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018a67ffffffffffffffff1681526020018b6001600160a01b031681526020018e606001358152602001866001600160a01b031681526020018e806000019061313b919061563c565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018e8060400190613193919061563c565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f8c018190048102820181019092528a815291810191908b908b908190840183828082843760009201919091525050509152506040516132149190602401615c60565b60408051601f198184030181529190526020810180516001600160e01b03167f3907753700000000000000000000000000000000000000000000000000000000179052600454859063ffffffff600160e01b909104166113886084613ab2565b5091509150816132b257806040517fe1cd55090000000000000000000000000000000000000000000000000000000081526004016107dd9190613fae565b80516020146132fa5780516040517f78ef80240000000000000000000000000000000000000000000000000000000081526020600482015260248101919091526044016107dd565b6000818060200190518101906133109190615d2d565b6040516001600160a01b038c166024820152604481018290529091506133939060640160408051601f198184030181529190526020810180516001600160e01b03167fa9059cbb00000000000000000000000000000000000000000000000000000000179052600454879063ffffffff600160c01b909104166113886084613ab2565b509093509150826133d257816040517fe1cd55090000000000000000000000000000000000000000000000000000000081526004016107dd9190613fae565b604080518082019091526001600160a01b0390951685526020850152509192505050979650505050505050565b600061342b827f01ffc9a700000000000000000000000000000000000000000000000000000000613463565b8015610a78575061345c827fffffffff00000000000000000000000000000000000000000000000000000000613463565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180516001600160e01b03167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015613506575060208210155b80156135125750600081115b979650505050505050565b60005b8151811015610cd75760ff831660009081526003602052604081208351909190849084908110613552576135526152ed565b6020908102919091018101516001600160a01b03168252810191909152604001600020805461ffff19169055600101613520565b60005b82518160ff161015610a0e576000838260ff16815181106135ac576135ac6152ed565b60200260200101519050600060028111156135c9576135c9614367565b60ff80871660009081526003602090815260408083206001600160a01b0387168452909152902054610100900416600281111561360857613608614367565b14613629576004604051631b3fab5160e11b81526004016107dd9190615acf565b6001600160a01b038116613669576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405180604001604052808360ff16815260200184600281111561368f5761368f614367565b905260ff80871660009081526003602090815260408083206001600160a01b0387168452825290912083518154931660ff198416811782559184015190929091839161ffff1916176101008360028111156136ec576136ec614367565b021790555090505050806136ff90615d46565b9050613589565b60ff8116610a2557600a805467ffffffffffffffff1916905550565b815160208082015160409283015192516000938493613768937f2425b0b9f9054c76ff151b0a175b18f37a4a4e82013a72e9f15c9caa095ed21f93909291889101615d65565b60408051601f1981840301815290829052805160209182012086518051888401516060808b0151908401516080808d015195015195976137b19794969395929491939101615d98565b604051602081830303815290604052805190602001208560400151805190602001208660a001516040516020016137e89190615e8f565b60408051601f198184030181528282528051602091820120908301969096528101939093526060830191909152608082015260a081019190915260c0015b60405160208183030381529060405280519060200120905092915050565b600080613852858585613bd8565b905061385d81611432565b61386b576000915050613892565b67ffffffffffffffff86166000908152600960209081526040808320938352929052205490505b949350505050565b600060026138a9608085615541565b67ffffffffffffffff166138bd9190615568565b905060006138cb8585611c8d565b9050816138da60016004615518565b901b1916818360038111156138f1576138f1614367565b67ffffffffffffffff871660009081526008602052604081209190921b929092179182916139206080886158a6565b67ffffffffffffffff1681526020810191909152604001600020555050505050565b6040517fece670b6000000000000000000000000000000000000000000000000000000008152600090606090309063ece670b6906139869087908790600401615eef565b600060405180830381600087803b1580156139a057600080fd5b505af19250505080156139b1575060015b6139f0573d8080156139df576040519150601f19603f3d011682016040523d82523d6000602084013e6139e4565b606091505b50600392509050613a05565b50506040805160208101909152600081526002905b9250929050565b60008151602014613a4b57816040517f8d666f600000000000000000000000000000000000000000000000000000000081526004016107dd9190613fae565b600082806020019051810190613a619190615d2d565b90506001600160a01b03811180613a79575061040081105b15610a7857826040517f8d666f600000000000000000000000000000000000000000000000000000000081526004016107dd9190613fae565b6000606060008361ffff1667ffffffffffffffff811115613ad557613ad5614160565b6040519080825280601f01601f191660200182016040528015613aff576020820181803683370190505b509150863b613b32577f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b5a85811015613b65577fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b8590036040810481038710613b9e577f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b505a6000808a5160208c0160008c8cf193505a900390503d84811115613bc15750835b808352806000602085013e50955095509592505050565b8251825160009190818303613c19576040517f11a6b26400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6101018211801590613c2d57506101018111155b613c4a576040516309bde33960e01b815260040160405180910390fd5b60001982820101610100811115613c74576040516309bde33960e01b815260040160405180910390fd5b80600003613ca15786600081518110613c8f57613c8f6152ed565b60200260200101519350505050613e70565b60008167ffffffffffffffff811115613cbc57613cbc614160565b604051908082528060200260200182016040528015613ce5578160200160208202803683370190505b50905060008080805b85811015613e0f5760006001821b8b811603613d495788851015613d32578c5160018601958e918110613d2357613d236152ed565b60200260200101519050613d6b565b8551600185019487918110613d2357613d236152ed565b8b5160018401938d918110613d6057613d606152ed565b602002602001015190505b600089861015613d9b578d5160018701968f918110613d8c57613d8c6152ed565b60200260200101519050613dbd565b8651600186019588918110613db257613db26152ed565b602002602001015190505b82851115613dde576040516309bde33960e01b815260040160405180910390fd5b613de88282613e77565b878481518110613dfa57613dfa6152ed565b60209081029190910101525050600101613cee565b506001850382148015613e2157508683145b8015613e2c57508581145b613e49576040516309bde33960e01b815260040160405180910390fd5b836001860381518110613e5e57613e5e6152ed565b60200260200101519750505050505050505b9392505050565b6000818310613e8f57613e8a8284613e95565b610a75565b610a7583835b604080516001602082015290810183905260608101829052600090608001613826565b828054828255906000526020600020908101928215613f1a579160200282015b82811115613f1a578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190613ed8565b50613f26929150613f49565b5090565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115613f265760008155600101613f4a565b60005b83811015613f79578181015183820152602001613f61565b50506000910152565b60008151808452613f9a816020860160208601613f5e565b601f01601f19169290920160200192915050565b602081526000610a756020830184613f82565b8060608101831015610a7857600080fd5b60008083601f840112613fe457600080fd5b50813567ffffffffffffffff811115613ffc57600080fd5b602083019150836020828501011115613a0557600080fd5b60008083601f84011261402657600080fd5b50813567ffffffffffffffff81111561403e57600080fd5b6020830191508360208260051b8501011115613a0557600080fd5b60008060008060008060008060e0898b03121561407557600080fd5b61407f8a8a613fc1565b9750606089013567ffffffffffffffff8082111561409c57600080fd5b6140a88c838d01613fd2565b909950975060808b01359150808211156140c157600080fd5b6140cd8c838d01614014565b909750955060a08b01359150808211156140e657600080fd5b506140f38b828c01614014565b999c989b50969995989497949560c00135949350505050565b60008060006080848603121561412157600080fd5b61412b8585613fc1565b9250606084013567ffffffffffffffff81111561414757600080fd5b61415386828701613fd2565b9497909650939450505050565b634e487b7160e01b600052604160045260246000fd5b60405160c0810167ffffffffffffffff8111828210171561419957614199614160565b60405290565b60405160a0810167ffffffffffffffff8111828210171561419957614199614160565b6040516080810167ffffffffffffffff8111828210171561419957614199614160565b6040516060810167ffffffffffffffff8111828210171561419957614199614160565b6040805190810167ffffffffffffffff8111828210171561419957614199614160565b604051601f8201601f1916810167ffffffffffffffff8111828210171561425457614254614160565b604052919050565b6001600160a01b0381168114610a2557600080fd5b803561427c8161425c565b919050565b803563ffffffff8116811461427c57600080fd5b600060c082840312156142a757600080fd5b6142af614176565b82356142ba8161425c565b81526142c860208401614281565b60208201526142d960408401614281565b60408201526142ea60608401614281565b606082015260808301356142fd8161425c565b608082015260a08301356143108161425c565b60a08201529392505050565b803567ffffffffffffffff8116811461427c57600080fd5b6000806040838503121561434757600080fd5b6143508361431c565b915061435e6020840161431c565b90509250929050565b634e487b7160e01b600052602160045260246000fd5b6004811061438d5761438d614367565b9052565b60208101610a78828461437d565b600067ffffffffffffffff8211156143b9576143b9614160565b5060051b60200190565b600060a082840312156143d557600080fd5b6143dd61419f565b9050813581526143ef6020830161431c565b60208201526144006040830161431c565b60408201526144116060830161431c565b60608201526144226080830161431c565b608082015292915050565b600067ffffffffffffffff82111561444757614447614160565b50601f01601f191660200190565b600082601f83011261446657600080fd5b81356144796144748261442d565b61422b565b81815284602083860101111561448e57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126144bc57600080fd5b813560206144cc6144748361439f565b82815260059290921b840181019181810190868411156144eb57600080fd5b8286015b848110156145c157803567ffffffffffffffff808211156145105760008081fd5b8189019150608080601f19848d0301121561452b5760008081fd5b6145336141c2565b87840135838111156145455760008081fd5b6145538d8a83880101614455565b8252506040808501358481111561456a5760008081fd5b6145788e8b83890101614455565b8a84015250606080860135858111156145915760008081fd5b61459f8f8c838a0101614455565b92840192909252949092013593810193909352505083529183019183016144ef565b509695505050505050565b600082601f8301126145dd57600080fd5b813560206145ed6144748361439f565b82815260059290921b8401810191818101908684111561460c57600080fd5b8286015b848110156145c157803567ffffffffffffffff808211156146315760008081fd5b818901915061014080601f19848d0301121561464d5760008081fd5b614655614176565b6146618c8986016143c3565b815260c0840135838111156146765760008081fd5b6146848d8a83880101614455565b898301525060e08401358381111561469c5760008081fd5b6146aa8d8a83880101614455565b6040830152506146bd6101008501614271565b606082015261012084013560808201529083013590828211156146e05760008081fd5b6146ee8c89848701016144ab565b60a08201528652505050918301918301614610565b600082601f83011261471457600080fd5b813560206147246144748361439f565b82815260059290921b8401810191818101908684111561474357600080fd5b8286015b848110156145c157803567ffffffffffffffff8082111561476757600080fd5b818901915089603f83011261477b57600080fd5b8582013561478b6144748261439f565b81815260059190911b830160400190878101908c8311156147ab57600080fd5b604085015b838110156147e4578035858111156147c757600080fd5b6147d68f6040838a0101614455565b8452509189019189016147b0565b50875250505092840192508301614747565b600082601f83011261480757600080fd5b813560206148176144748361439f565b8083825260208201915060208460051b87010193508684111561483957600080fd5b602086015b848110156145c1578035835291830191830161483e565b600082601f83011261486657600080fd5b813560206148766144748361439f565b82815260059290921b8401810191818101908684111561489557600080fd5b8286015b848110156145c157803567ffffffffffffffff808211156148ba5760008081fd5b818901915060a080601f19848d030112156148d55760008081fd5b6148dd61419f565b6148e888850161431c565b8152604080850135848111156148fe5760008081fd5b61490c8e8b838901016145cc565b8a84015250606080860135858111156149255760008081fd5b6149338f8c838a0101614703565b838501525060809150818601358581111561494e5760008081fd5b61495c8f8c838a01016147f6565b9184019190915250919093013590830152508352918301918301614899565b600080604080848603121561498f57600080fd5b833567ffffffffffffffff808211156149a757600080fd5b6149b387838801614855565b94506020915081860135818111156149ca57600080fd5b8601601f810188136149db57600080fd5b80356149e96144748261439f565b81815260059190911b8201840190848101908a831115614a0857600080fd5b8584015b83811015614a9457803586811115614a245760008081fd5b8501603f81018d13614a365760008081fd5b87810135614a466144748261439f565b81815260059190911b82018a0190898101908f831115614a665760008081fd5b928b01925b82841015614a845783358252928a0192908a0190614a6b565b8652505050918601918601614a0c565b50809750505050505050509250929050565b600060208284031215614ab857600080fd5b813567ffffffffffffffff811115614acf57600080fd5b820160a08185031215613e7057600080fd5b8015158114610a2557600080fd5b803561427c81614ae1565b60006020808385031215614b0d57600080fd5b823567ffffffffffffffff80821115614b2557600080fd5b818501915085601f830112614b3957600080fd5b8135614b476144748261439f565b81815260059190911b83018401908481019088831115614b6657600080fd5b8585015b838110156120c857803585811115614b825760008081fd5b86016060818c03601f1901811315614b9a5760008081fd5b614ba26141e5565b614bad8a840161431c565b8152604080840135614bbe81614ae1565b828c0152918301359188831115614bd55760008081fd5b614be38e8c85870101614455565b9082015285525050918601918601614b6a565b803560ff8116811461427c57600080fd5b600060208284031215614c1957600080fd5b610a7582614bf6565b60008151808452602080850194506020840160005b83811015614c5c5781516001600160a01b031687529582019590820190600101614c37565b509495945050505050565b60208152600082518051602084015260ff602082015116604084015260ff604082015116606084015260608101511515608084015250602083015160c060a0840152614cb660e0840182614c22565b90506040840151601f198483030160c0850152614cd38282614c22565b95945050505050565b60008060408385031215614cef57600080fd5b614cf88361431c565b946020939093013593505050565b60008060208385031215614d1957600080fd5b823567ffffffffffffffff80821115614d3157600080fd5b818501915085601f830112614d4557600080fd5b813581811115614d5457600080fd5b8660208260061b8501011115614d6957600080fd5b60209290920196919550909350505050565b600060208284031215614d8d57600080fd5b610a758261431c565b6020815281511515602082015267ffffffffffffffff6020830151166040820152600060408301516060808401526138926080840182613f82565b600080600060408486031215614de657600080fd5b833567ffffffffffffffff80821115614dfe57600080fd5b908501906101408288031215614e1357600080fd5b90935060208501359080821115614e2957600080fd5b5061415386828701614014565b600060208284031215614e4857600080fd5b8135613e708161425c565b600082601f830112614e6457600080fd5b81356020614e746144748361439f565b8083825260208201915060208460051b870101935086841115614e9657600080fd5b602086015b848110156145c1578035614eae8161425c565b8352918301918301614e9b565b60006020808385031215614ece57600080fd5b823567ffffffffffffffff80821115614ee657600080fd5b818501915085601f830112614efa57600080fd5b8135614f086144748261439f565b81815260059190911b83018401908481019088831115614f2757600080fd5b8585015b838110156120c857803585811115614f4257600080fd5b860160c0818c03601f19011215614f595760008081fd5b614f61614176565b8882013581526040614f74818401614bf6565b8a8301526060614f85818501614bf6565b8284015260809150614f98828501614aef565b9083015260a08381013589811115614fb05760008081fd5b614fbe8f8d83880101614e53565b838501525060c0840135915088821115614fd85760008081fd5b614fe68e8c84870101614e53565b9083015250845250918601918601614f2b565b60006020828403121561500b57600080fd5b5035919050565b80356001600160e01b038116811461427c57600080fd5b600082601f83011261503a57600080fd5b8135602061504a6144748361439f565b82815260069290921b8401810191818101908684111561506957600080fd5b8286015b848110156145c157604081890312156150865760008081fd5b61508e614208565b6150978261431c565b81526150a4858301615012565b8186015283529183019160400161506d565b600082601f8301126150c757600080fd5b813560206150d76144748361439f565b82815260079290921b840181019181810190868411156150f657600080fd5b8286015b848110156145c15780880360808112156151145760008081fd5b61511c6141e5565b6151258361431c565b8152604080601f198401121561513b5760008081fd5b615143614208565b925061515087850161431c565b835261515d81850161431c565b83880152818701929092526060830135918101919091528352918301916080016150fa565b6000602080838503121561519557600080fd5b823567ffffffffffffffff808211156151ad57600080fd5b818501915060408083880312156151c357600080fd5b6151cb614208565b8335838111156151da57600080fd5b84016040818a0312156151ec57600080fd5b6151f4614208565b81358581111561520357600080fd5b8201601f81018b1361521457600080fd5b80356152226144748261439f565b81815260069190911b8201890190898101908d83111561524157600080fd5b928a01925b828410156152915787848f03121561525e5760008081fd5b615266614208565b84356152718161425c565b815261527e858d01615012565b818d0152825292870192908a0190615246565b8452505050818701359350848411156152a957600080fd5b6152b58a858401615029565b81880152825250838501359150828211156152cf57600080fd5b6152db888386016150b6565b85820152809550505050505092915050565b634e487b7160e01b600052603260045260246000fd5b805160408084528151848201819052600092602091908201906060870190855b8181101561535a57835180516001600160a01b031684528501516001600160e01b0316858401529284019291850191600101615323565b50508583015187820388850152805180835290840192506000918401905b808310156153b4578351805167ffffffffffffffff1683528501516001600160e01b031685830152928401926001929092019190850190615378565b50979650505050505050565b602081526000610a756020830184615303565b67ffffffffffffffff8316815260608101613e706020830184805167ffffffffffffffff908116835260209182015116910152565b634e487b7160e01b600052601160045260246000fd5b67ffffffffffffffff81811683821601908082111561543f5761543f615408565b5092915050565b6000602080835260608451604080848701526154656060870183615303565b87850151878203601f19016040890152805180835290860193506000918601905b808310156120c857845167ffffffffffffffff8151168352878101516154c589850182805167ffffffffffffffff908116835260209182015116910152565b50840151828701529386019360019290920191608090910190615486565b6000602082840312156154f557600080fd5b813567ffffffffffffffff81111561550c57600080fd5b61389284828501614855565b81810381811115610a7857610a78615408565b634e487b7160e01b600052601260045260246000fd5b600067ffffffffffffffff8084168061555c5761555c61552b565b92169190910692915050565b8082028115828204841417610a7857610a78615408565b60006040828403121561559157600080fd5b615599614208565b6155a28361431c565b8152602083013560208201528091505092915050565b600181811c908216806155cc57607f821691505b6020821081036155ec57634e487b7160e01b600052602260045260246000fd5b50919050565b6000808335601e1984360301811261560957600080fd5b83018035915067ffffffffffffffff82111561562457600080fd5b6020019150600581901b3603821315613a0557600080fd5b6000808335601e1984360301811261565357600080fd5b83018035915067ffffffffffffffff82111561566e57600080fd5b602001915036819003821315613a0557600080fd5b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b805182526000602067ffffffffffffffff81840151168185015260408084015160a060408701526156e660a0870182613f82565b9050606085015186820360608801526156ff8282613f82565b608087810151898303918a01919091528051808352908601935060009250908501905b808310156153b457835180516001600160a01b0316835286015186830152928501926001929092019190840190615722565b602081526000610a7560208301846156b2565b60808152600061577a60808301876156b2565b61ffff9590951660208301525060408101929092526001600160a01b0316606090910152919050565b6000806000606084860312156157b857600080fd5b83516157c381614ae1565b602085015190935067ffffffffffffffff8111156157e057600080fd5b8401601f810186136157f157600080fd5b80516157ff6144748261442d565b81815287602083850101111561581457600080fd5b615825826020830160208601613f5e565b809450505050604084015190509250925092565b60006020828403121561584b57600080fd5b8151613e7081614ae1565b80820180821115610a7857610a78615408565b60ff8181168382160190811115610a7857610a78615408565b8183823760009101908152919050565b828152606082602083013760800192915050565b600067ffffffffffffffff808416806158c1576158c161552b565b92169190910492915050565b601f821115610cd7576000816000526020600020601f850160051c810160208610156158f65750805b601f850160051c820191505b818110156124e357828155600101615902565b815167ffffffffffffffff81111561592f5761592f614160565b6159438161593d84546155b8565b846158cd565b602080601f83116001811461597857600084156159605750858301515b600019600386901b1c1916600185901b1785556124e3565b600085815260208120601f198616915b828110156159a757888601518255948401946001909101908401615988565b50858210156159c55787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60006020808352835460ff81161515602085015267ffffffffffffffff8160081c16604085015250600180850160608086015260008154615a15816155b8565b80608089015260a06001831660008114615a365760018114615a5257615a82565b60ff19841660a08b015260a083151560051b8b01019450615a82565b85600052602060002060005b84811015615a795781548c8201850152908801908901615a5e565b8b0160a0019550505b50929998505050505050505050565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112615ac557600080fd5b9190910192915050565b6020810160058310615ae357615ae3614367565b91905290565b60ff818116838216029081169081811461543f5761543f615408565b600060a0820160ff881683526020878185015260a0604085015281875480845260c0860191508860005282600020935060005b81811015615b5d5784546001600160a01b031683526001948501949284019201615b38565b50508481036060860152865180825290820192508187019060005b81811015615b9d5782516001600160a01b031685529383019391830191600101615b78565b50505060ff851660808501525090505b9695505050505050565b600067ffffffffffffffff808616835280851660208401525060606040830152614cd36060830184613f82565b8281526040602082015260006138926040830184613f82565b67ffffffffffffffff84811682528316602082015260608101613892604083018461437d565b615c2d818461437d565b6040602082015260006138926040830184613f82565b600060208284031215615c5557600080fd5b8151613e708161425c565b6020815260008251610100806020850152615c7f610120850183613f82565b91506020850151615c9c604086018267ffffffffffffffff169052565b5060408501516001600160a01b038116606086015250606085015160808501526080850151615cd660a08601826001600160a01b03169052565b5060a0850151601f19808685030160c0870152615cf38483613f82565b935060c08701519150808685030160e0870152615d108483613f82565b935060e0870151915080868503018387015250615bad8382613f82565b600060208284031215615d3f57600080fd5b5051919050565b600060ff821660ff8103615d5c57615d5c615408565b60010192915050565b848152600067ffffffffffffffff808616602084015280851660408401525060806060830152615bad6080830184613f82565b86815260c060208201526000615db160c0830188613f82565b6001600160a01b039690961660408301525067ffffffffffffffff9384166060820152608081019290925290911660a09091015292915050565b600082825180855260208086019550808260051b84010181860160005b84811015615e8257601f19868403018952815160808151818652615e2e82870182613f82565b9150508582015185820387870152615e468282613f82565b91505060408083015186830382880152615e608382613f82565b6060948501519790940196909652505098840198925090830190600101615e08565b5090979650505050505050565b602081526000610a756020830184615deb565b60008282518085526020808601955060208260051b8401016020860160005b84811015615e8257601f19868403018952615edd838351613f82565b98840198925090830190600101615ec1565b604081526000835180516040840152602081015167ffffffffffffffff80821660608601528060408401511660808601528060608401511660a08601528060808401511660c086015250505060208401516101408060e0850152615f57610180850183613f82565b915060408601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08086850301610100870152615f948483613f82565b935060608801519150615fb36101208701836001600160a01b03169052565b60808801518387015260a0880151925080868503016101608701525050615fda8282615deb565b9150508281036020840152614cd38185615ea256fea164736f6c6343000818000a",
}

var EVM2EVMMultiOffRampABI = EVM2EVMMultiOffRampMetaData.ABI

var EVM2EVMMultiOffRampBin = EVM2EVMMultiOffRampMetaData.Bin

func DeployEVM2EVMMultiOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMMultiOffRampStaticConfig, dynamicConfig EVM2EVMMultiOffRampDynamicConfig, sourceChainConfigs []EVM2EVMMultiOffRampSourceChainConfigArgs) (common.Address, *types.Transaction, *EVM2EVMMultiOffRamp, error) {
	parsed, err := EVM2EVMMultiOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMMultiOffRampBin), backend, staticConfig, dynamicConfig, sourceChainConfigs)
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

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMRampMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.contract.Transact(opts, "executeSingleMessage", message, offchainTokenData)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampSession) ExecuteSingleMessage(message InternalAny2EVMRampMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
	return _EVM2EVMMultiOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMMultiOffRamp.TransactOpts, message, offchainTokenData)
}

func (_EVM2EVMMultiOffRamp *EVM2EVMMultiOffRampTransactorSession) ExecuteSingleMessage(message InternalAny2EVMRampMessage, offchainTokenData [][]byte) (*types.Transaction, error) {
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
	case _EVM2EVMMultiOffRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMMultiOffRamp.ParseOwnershipTransferRequested(log)
	case _EVM2EVMMultiOffRamp.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMMultiOffRamp.ParseOwnershipTransferred(log)
	case _EVM2EVMMultiOffRamp.abi.Events["RootRemoved"].ID:
		return _EVM2EVMMultiOffRamp.ParseRootRemoved(log)
	case _EVM2EVMMultiOffRamp.abi.Events["SkippedAlreadyExecutedMessage"].ID:
		return _EVM2EVMMultiOffRamp.ParseSkippedAlreadyExecutedMessage(log)
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

func (EVM2EVMMultiOffRampSourceChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x4f49973170c548fddd4a48341b75e131818913f38f44d47af57e8735eee588ba")
}

func (EVM2EVMMultiOffRampSourceChainSelectorAdded) Topic() common.Hash {
	return common.HexToHash("0xf4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb9")
}

func (EVM2EVMMultiOffRampStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0x683eb52ee924eb817377cfa8f41f238f4bb7a877da5267869dfffbad85f564d8")
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

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMRampMessage, offchainTokenData [][]byte) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, reports []InternalExecutionReportSingleChain, gasLimitOverrides [][]*big.Int) (*types.Transaction, error)

	ResetUnblessedRoots(opts *bind.TransactOpts, rootToReset []EVM2EVMMultiOffRampUnblessedRoot) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig EVM2EVMMultiOffRampDynamicConfig) (*types.Transaction, error)

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
