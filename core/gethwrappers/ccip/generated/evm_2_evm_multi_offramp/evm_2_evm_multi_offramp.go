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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumMultiOCR3Base.InvalidConfigErrorType\",\"name\":\"errorType\",\"type\":\"uint8\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.Interval\",\"name\":\"interval\",\"type\":\"tuple\"}],\"name\":\"InvalidInterval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"InvalidManualExecutionGasLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"messageDestChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidMessageDestChainSelector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"InvalidNewState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRoot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidStaticConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeavesCannotBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionGasLimitMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"errorReason\",\"type\":\"bytes\"}],\"name\":\"MessageValidationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonUniqueSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notPool\",\"type\":\"address\"}],\"name\":\"NotACompatiblePool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"RootAlreadyCommitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignaturesOutOfRegistration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StaleCommitReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"}],\"name\":\"StaticConfigCannotBeChanged\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"TokenDataMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"TokenHandlingError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedTokenData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongNumberOfSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroChainSelectorNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint224\",\"name\":\"usdPerToken\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint224\",\"name\":\"usdPerUnitGas\",\"type\":\"uint224\"}],\"internalType\":\"structInternal.GasPriceUpdate[]\",\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.PriceUpdates\",\"name\":\"priceUpdates\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.Interval\",\"name\":\"interval\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.MerkleRoot[]\",\"name\":\"merkleRoots\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.CommitReport\",\"name\":\"report\",\"type\":\"tuple\"}],\"name\":\"CommitReportAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"DynamicConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"RootRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"SkippedAlreadyExecutedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfig\",\"name\":\"sourceConfig\",\"type\":\"tuple\"}],\"name\":\"SourceChainConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"SourceChainSelectorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"}],\"name\":\"StaticConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfigArgs[]\",\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\"}],\"name\":\"applySourceChainConfigUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.RampTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.Any2EVMRampMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPriceSequenceNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"getSourceChainConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"minSeqNr\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"onRamp\",\"type\":\"bytes\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.SourceChainConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nonceManager\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"isBlessed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"}],\"name\":\"latestConfigDetails\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"n\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\"}],\"internalType\":\"structMultiOCR3Base.ConfigInfo\",\"name\":\"configInfo\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"}],\"internalType\":\"structMultiOCR3Base.OCRConfig\",\"name\":\"ocrConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"destChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structInternal.RampMessageHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destTokenAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.RampTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structInternal.Any2EVMRampMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReportSingleChain[]\",\"name\":\"reports\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"gasLimitOverrides\",\"type\":\"uint256[][]\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.UnblessedRoot[]\",\"name\":\"rootToReset\",\"type\":\"tuple[]\"}],\"name\":\"resetUnblessedRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxTokenTransferGas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPoolReleaseOrMintGas\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"messageValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMMultiOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"setDynamicConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"ocrPluginType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"F\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSignatureVerificationEnabled\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"}],\"internalType\":\"structMultiOCR3Base.OCRConfigArgs[]\",\"name\":\"ocrConfigArgs\",\"type\":\"tuple[]\"}],\"name\":\"setOCR3Configs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162006a6138038062006a618339810160408190526200003591620008e2565b33806000816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000bf57620000bf81620001fa565b5050466080525060208301516001600160a01b03161580620000ec575060408301516001600160a01b0316155b8062000103575060608301516001600160a01b0316155b1562000122576040516342bcdf7f60e11b815260040160405180910390fd5b82516001600160401b03166000036200014e5760405163c656089560e01b815260040160405180910390fd5b82516001600160401b0390811660a052602080850180516001600160a01b0390811660c05260408088018051831660e0526060808a01805185166101005283518b519098168852945184169587019590955251821690850152905116908201527f683eb52ee924eb817377cfa8f41f238f4bb7a877da5267869dfffbad85f564d89060800160405180910390a1620001e682620002a5565b620001f181620003fc565b50505062000c57565b336001600160a01b03821603620002545760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60a08101516001600160a01b03161580620002c8575080516001600160a01b0316155b15620002e7576040516342bcdf7f60e11b815260040160405180910390fd5b80516004805460208085018051604080880180516060808b0180516001600160a01b039b8c166001600160c01b0319909a168a17600160a01b63ffffffff98891602176001600160c01b0316600160c01b948816949094026001600160e01b031693909317600160e01b93871693909302929092179098556080808b018051600580546001600160a01b0319908116928e1692909217905560a0808e01805160068054909416908f161790925586519a8b5297518716988a0198909852925185169388019390935251909216958501959095525185169383019390935251909216908201527f0da37fd00459f4f5f0b8210d31525e4910ae674b8bab34b561d146bb45773a4c9060c00160405180910390a150565b60005b81518110156200064e57600082828151811062000420576200042062000a20565b60200260200101519050600081600001519050806001600160401b03166000036200045e5760405163c656089560e01b815260040160405180910390fd5b6001600160401b03811660009081526007602052604081206001810180549192916200048a9062000a36565b80601f0160208091040260200160405190810160405280929190818152602001828054620004b89062000a36565b8015620005095780601f10620004dd5761010080835404028352916020019162000509565b820191906000526020600020905b815481529060010190602001808311620004eb57829003601f168201915b505050505090506000846040015190508151600003620005ac57805160000362000546576040516342bcdf7f60e11b815260040160405180910390fd5b6001830162000556828262000ac7565b508254610100600160481b0319166101001783556040516001600160401b03851681527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb99060200160405180910390a1620005e7565b8080519060200120828051906020012014620005e75760405163c39a620560e01b81526001600160401b038516600482015260240162000083565b6020850151835460ff19169015151783556040516001600160401b038516907f4f49973170c548fddd4a48341b75e131818913f38f44d47af57e8735eee588ba906200063590869062000b93565b60405180910390a25050505050806001019050620003ff565b5050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156200068d576200068d62000652565b60405290565b604051608081016001600160401b03811182821017156200068d576200068d62000652565b60405160c081016001600160401b03811182821017156200068d576200068d62000652565b604051601f8201601f191681016001600160401b038111828210171562000708576200070862000652565b604052919050565b80516001600160401b03811681146200072857600080fd5b919050565b80516001600160a01b03811681146200072857600080fd5b805163ffffffff811681146200072857600080fd5b6000601f83601f8401126200076e57600080fd5b825160206001600160401b03808311156200078d576200078d62000652565b8260051b6200079e838201620006dd565b9384528681018301938381019089861115620007b957600080fd5b84890192505b85831015620008d557825184811115620007d95760008081fd5b89016060601f19828d038101821315620007f35760008081fd5b620007fd62000668565b6200080a89850162000710565b81526040808501518015158114620008225760008081fd5b828b01529284015192888411156200083a5760008081fd5b83850194508e603f8601126200085257600093508384fd5b898501519350888411156200086b576200086b62000652565b6200087c8a848e87011601620006dd565b92508383528e81858701011115620008945760008081fd5b60005b84811015620008b4578581018201518482018c01528a0162000897565b5060009383018a0193909352918201528352509184019190840190620007bf565b9998505050505050505050565b6000806000838503610160811215620008fa57600080fd5b60808112156200090957600080fd5b6200091362000693565b6200091e8662000710565b81526200092e602087016200072d565b602082015262000941604087016200072d565b604082015262000954606087016200072d565b6060820152935060c0607f19820112156200096e57600080fd5b5062000979620006b8565b62000987608086016200072d565b81526200099760a0860162000745565b6020820152620009aa60c0860162000745565b6040820152620009bd60e0860162000745565b6060820152620009d161010086016200072d565b6080820152620009e561012086016200072d565b60a08201526101408501519092506001600160401b0381111562000a0857600080fd5b62000a16868287016200075a565b9150509250925092565b634e487b7160e01b600052603260045260246000fd5b600181811c9082168062000a4b57607f821691505b60208210810362000a6c57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111562000ac2576000816000526020600020601f850160051c8101602086101562000a9d5750805b601f850160051c820191505b8181101562000abe5782815560010162000aa9565b5050505b505050565b81516001600160401b0381111562000ae35762000ae362000652565b62000afb8162000af4845462000a36565b8462000a72565b602080601f83116001811462000b33576000841562000b1a5750858301515b600019600386901b1c1916600185901b17855562000abe565b600085815260208120601f198616915b8281101562000b645788860151825594840194600190910190840162000b43565b508582101562000b835787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6020808252825460ff811615158383015260081c6001600160401b0316604083015260608083015260018084018054600093929190849062000bd58162000a36565b80608089015260a0600183166000811462000bf9576001811462000c165762000c48565b60ff19841660a08b015260a083151560051b8b0101945062000c48565b85600052602060002060005b8481101562000c3f5781548c820185015290880190890162000c22565b8b0160a0019550505b50929998505050505050505050565b60805160a05160c05160e05161010051615d9b62000cc66000396000818161023e0152612c1601526000818161020f0152612f000152600081816101e00152818161140401526114bb0152600081816101b0015261278a015260008181611c680152611cb40152615d9b6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806385572ffb116100d8578063d2a15d351161008c578063f2fde38b11610066578063f2fde38b146105b1578063f716f99f146105c4578063ff888fb1146105d757600080fd5b8063d2a15d351461056b578063e9d68a8e1461057e578063ece670b61461059e57600080fd5b8063a12a9870116100bd578063a12a9870146104f3578063c673e58414610506578063ccd37ba31461052657600080fd5b806385572ffb146104ca5780638da5cb5b146104d857600080fd5b8063403b2d631161012f5780637437ff9f116101145780637437ff9f1461038557806379ba5097146104af5780637d4eef60146104b757600080fd5b8063403b2d63146103525780635e36480c1461036557600080fd5b80632d04ab76116101605780632d04ab761461030e578063311cd513146103235780633f4b04aa1461033657600080fd5b806306285c691461017c578063181f5a77146102c5575b600080fd5b61026e60408051608081018252600080825260208201819052918101829052606081019190915260405180608001604052807f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b6040516102bc9190815167ffffffffffffffff1681526020808301516001600160a01b0390811691830191909152604080840151821690830152606092830151169181019190915260800190565b60405180910390f35b6103016040518060400160405280601d81526020017f45564d3245564d4d756c74694f666652616d7020312e362e302d64657600000081525081565b6040516102bc9190613e47565b61032161031c366004613ef2565b6105fa565b005b610321610331366004613fa5565b6109c0565b600a5460405167ffffffffffffffff90911681526020016102bc565b61032161036036600461412e565b610a29565b6103786103733660046141cd565b610a3d565b6040516102bc919061422a565b6104466040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c0810182526004546001600160a01b03808216835263ffffffff7401000000000000000000000000000000000000000083048116602085015278010000000000000000000000000000000000000000000000008304811694840194909452600160e01b90910490921660608201526005548216608082015260065490911660a082015290565b6040516102bc9190600060c0820190506001600160a01b03808451168352602084015163ffffffff808216602086015280604087015116604086015280606087015116606086015250508060808501511660808401528060a08501511660a08401525092915050565b610321610a93565b6103216104c536600461482e565b610b51565b610321610177366004614959565b6000546040516001600160a01b0390911681526020016102bc565b6103216105013660046149ad565b610cf6565b610519610514366004614ac7565b610d07565b6040516102bc9190614b27565b61055d610534366004614b9c565b67ffffffffffffffff919091166000908152600960209081526040808320938352929052205490565b6040519081526020016102bc565b610321610579366004614bc6565b610e65565b61059161058c366004614c3b565b610f1f565b6040516102bc9190614c56565b6103216105ac366004614c91565b611008565b6103216105bf366004614cf5565b61135b565b6103216105d2366004614d7a565b61136c565b6105ea6105e5366004614eb8565b6113ae565b60405190151581526020016102bc565b600061060887890189615041565b8051515190915015158061062157508051602001515115155b1561072157600a5460208a01359067ffffffffffffffff808316911610156106e057600a805467ffffffffffffffff191667ffffffffffffffff831617905560065482516040517f3937306f0000000000000000000000000000000000000000000000000000000081526001600160a01b0390921691633937306f916106a99160040161527f565b600060405180830381600087803b1580156106c357600080fd5b505af11580156106d7573d6000803e3d6000fd5b5050505061071f565b81602001515160000361071f576040517f2261116700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b60005b81602001515181101561090957600082602001518281518110610749576107496151ac565b602002602001015190506000816000015190506107658161146f565b600061077082611571565b602084015151815491925067ffffffffffffffff908116610100909204161415806107b2575060208084015190810151905167ffffffffffffffff9182169116115b156107fb57825160208401516040517feefb0cac0000000000000000000000000000000000000000000000000000000081526107f2929190600401615292565b60405180910390fd5b604083015180610837576040517f504570e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b835167ffffffffffffffff166000908152600960209081526040808320848452909152902054156108aa5783516040517f32cf0cbf00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602481018290526044016107f2565b60208085015101516108bd9060016152dd565b825468ffffffffffffffff00191661010067ffffffffffffffff928316021790925592511660009081526009602090815260408083209483529390529190912042905550600101610724565b507f3a3950e13dd607cc37980db0ef14266c40d2bba9c01b2e44bfe549808883095d816040516109399190615305565b60405180910390a16109b560008a8a8a8a8a8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c9182918501908490808284376000920191909152508b92506115d1915050565b505050505050505050565b610a006109cf828401846153a2565b60408051600080825260208201909252906109fa565b60608152602001906001900390816109e55790505b50611948565b604080516000808252602082019092529050610a236001858585858660006115d1565b50505050565b610a316119f8565b610a3a81611a54565b50565b6000610a4b600160046153d7565b6002610a58608085615400565b67ffffffffffffffff16610a6c9190615427565b610a768585611c1e565b901c166003811115610a8a57610a8a614200565b90505b92915050565b6001546001600160a01b03163314610aed5760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016107f2565b600080543373ffffffffffffffffffffffffffffffffffffffff19808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610b59611c65565b815181518114610b95576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b81811015610ce6576000848281518110610bb457610bb46151ac565b60200260200101519050600081602001515190506000858481518110610bdc57610bdc6151ac565b6020026020010151905080518214610c20576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b82811015610cd7576000828281518110610c3f57610c3f6151ac565b6020026020010151905080600014158015610c7a575084602001518281518110610c6b57610c6b6151ac565b60200260200101516080015181105b15610cce5784516040517fc8e9605100000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015260248101839052604481018290526064016107f2565b50600101610c23565b50505050806001019050610b98565b50610cf18383611948565b505050565b610cfe6119f8565b610a3a81611ce6565b610d4a6040805160e081019091526000606082018181526080830182905260a0830182905260c08301919091528190815260200160608152602001606081525090565b60ff808316600090815260026020818152604092839020835160e081018552815460608201908152600183015480881660808401526101008104881660a0840152620100009004909616151560c082015294855291820180548451818402810184019095528085529293858301939092830182828015610df357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610dd5575b5050505050815260200160038201805480602002602001604051908101604052809291908181526020018280548015610e5557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e37575b5050505050815250509050919050565b610e6d6119f8565b60005b81811015610cf1576000838383818110610e8c57610e8c6151ac565b905060400201803603810190610ea2919061543e565b9050610eb181602001516113ae565b610f1657805167ffffffffffffffff1660009081526009602090815260408083208285018051855290835281842093909355915191519182527f202f1139a3e334b6056064c0e9b19fd07e44a88d8f6e5ded571b24cf8c371f12910160405180910390a15b50600101610e70565b60408051606080820183526000808352602080840182905283850183905267ffffffffffffffff8681168352600782529185902085519384018652805460ff811615158552610100900490921690830152600181018054939492939192840191610f8890615477565b80601f0160208091040260200160405190810160405280929190818152602001828054610fb490615477565b8015610e555780601f10610fd657610100808354040283529160200191610e55565b820191906000526020600020905b815481529060010190602001808311610fe457505050919092525091949350505050565b333014611041576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516000808252602082019092528161107e565b60408051808201909152600080825260208201528152602001906001900390816110575790505b5060a084015151909150156110b1576110ae8360a001518460200151856060015186600001516020015186611f72565b90505b6040805160a0810182528451518152845160209081015167ffffffffffffffff16818301528086015183516000948401926110ed929101613e47565b60408051601f19818403018152918152908252868101516020830152018390526005549091506001600160a01b031680156111fa576040517f08d450a10000000000000000000000000000000000000000000000000000000081526001600160a01b038216906308d450a190611167908590600401615553565b600060405180830381600087803b15801561118157600080fd5b505af1925050508015611192575060015b6111fa573d8080156111c0576040519150601f19603f3d011682016040523d82523d6000602084013e6111c5565b606091505b50806040517f09c253250000000000000000000000000000000000000000000000000000000081526004016107f29190613e47565b60408501515115801561120f57506080850151155b80611226575060608501516001600160a01b03163b155b8061126657506060850151611264906001600160a01b03167f85572ffb00000000000000000000000000000000000000000000000000000000612050565b155b15611272575050505050565b60048054608087015160608801516040517f3cf9798300000000000000000000000000000000000000000000000000000000815260009485946001600160a01b031693633cf97983936112cd938a9361138893929101615566565b6000604051808303816000875af11580156112ec573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261131491908101906155a2565b50915091508161135257806040517f0a8d6e8c0000000000000000000000000000000000000000000000000000000081526004016107f29190613e47565b50505050505050565b6113636119f8565b610a3a8161206c565b6113746119f8565b60005b81518110156113aa576113a2828281518110611395576113956151ac565b6020026020010151612122565b600101611377565b5050565b6040805180820182523081526020810183815291517f4d61677100000000000000000000000000000000000000000000000000000000815290516001600160a01b039081166004830152915160248201526000917f00000000000000000000000000000000000000000000000000000000000000001690634d61677190604401602060405180830381865afa15801561144b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a8d9190615638565b6040517f2cbc26bb000000000000000000000000000000000000000000000000000000008152608082901b77ffffffffffffffff000000000000000000000000000000001660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632cbc26bb90602401602060405180830381865afa15801561150a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061152e9190615638565b15610a3a576040517ffdbd6a7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201526024016107f2565b67ffffffffffffffff81166000908152600760205260408120805460ff16610a8d576040517fed053c5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024016107f2565b60ff878116600090815260026020908152604080832081516080810183528154815260019091015480861693820193909352610100830485169181019190915262010000909104909216151560608301528735906116308760a4615655565b9050826060015115611678578451611649906020615427565b8651611656906020615427565b6116619060a0615655565b61166b9190615655565b6116759082615655565b90505b3681146116ba576040517f8e1192e1000000000000000000000000000000000000000000000000000000008152600481018290523660248201526044016107f2565b50815181146117025781516040517f93df584c0000000000000000000000000000000000000000000000000000000081526004810191909152602481018290526044016107f2565b61170a611c65565b60ff808a166000908152600360209081526040808320338452825280832081518083019092528054808616835293949193909284019161010090910416600281111561175857611758614200565b600281111561176957611769614200565b905250905060028160200151600281111561178657611786614200565b1480156117da5750600260008b60ff1660ff168152602001908152602001600020600301816000015160ff16815481106117c2576117c26151ac565b6000918252602090912001546001600160a01b031633145b611810576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b508160600151156118f257602082015161182b906001615668565b60ff16855114611867576040517f71253a2500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83518551146118a2576040517fa75d88af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600087876040516118b4929190615681565b6040519081900381206118cb918b90602001615691565b6040516020818303038152906040528051906020012090506118f08a82888888612466565b505b6040805182815260208a81013567ffffffffffffffff169082015260ff8b16917f198d6990ef96613a9026203077e422916918b03ff47f0be6bee7b02d8e139ef0910160405180910390a2505050505050505050565b8151600003611982576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805160408051600080825260208201909252911591905b84518110156119f1576119e98582815181106119b7576119b76151ac565b6020026020010151846119e3578583815181106119d6576119d66151ac565b602002602001015161267d565b8361267d565b600101611999565b5050505050565b6000546001600160a01b03163314611a525760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016107f2565b565b60a08101516001600160a01b03161580611a76575080516001600160a01b0316155b15611aad576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516004805460208085018051604080880180516060808b0180516001600160a01b039b8c167fffffffffffffffff000000000000000000000000000000000000000000000000909a168a177401000000000000000000000000000000000000000063ffffffff988916021777ffffffffffffffffffffffffffffffffffffffffffffffff167801000000000000000000000000000000000000000000000000948816949094026001600160e01b031693909317600160e01b93871693909302929092179098556080808b0180516005805473ffffffffffffffffffffffffffffffffffffffff19908116928e1692909217905560a0808e01805160068054909416908f161790925586519a8b5297518716988a0198909852925185169388019390935251909216958501959095525185169383019390935251909216908201527f0da37fd00459f4f5f0b8210d31525e4910ae674b8bab34b561d146bb45773a4c9060c00160405180910390a150565b67ffffffffffffffff8216600090815260086020526040812081611c436080856156a5565b67ffffffffffffffff1681526020810191909152604001600020549392505050565b467f000000000000000000000000000000000000000000000000000000000000000014611a52576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201524660248201526044016107f2565b60005b81518110156113aa576000828281518110611d0657611d066151ac565b602002602001015190506000816000015190508067ffffffffffffffff16600003611d5d576040517fc656089500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff81166000908152600760205260408120600181018054919291611d8890615477565b80601f0160208091040260200160405190810160405280929190818152602001828054611db490615477565b8015611e015780601f10611dd657610100808354040283529160200191611e01565b820191906000526020600020905b815481529060010190602001808311611de457829003601f168201915b505050505090506000846040015190508151600003611eba578051600003611e55576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018301611e638282615714565b50825468ffffffffffffffff00191661010017835560405167ffffffffffffffff851681527ff4c1390c70e5c0f491ae1ccbc06f9117cbbadf2767b247b3bc203280f24c0fb99060200160405180910390a1611f0d565b8080519060200120828051906020012014611f0d576040517fc39a620500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526024016107f2565b6020850151835460ff191690151517835560405167ffffffffffffffff8516907f4f49973170c548fddd4a48341b75e131818913f38f44d47af57e8735eee588ba90611f5a9086906157d4565b60405180910390a25050505050806001019050611ce9565b6060855167ffffffffffffffff811115611f8e57611f8e613ff9565b604051908082528060200260200182016040528015611fd357816020015b6040805180820190915260008082526020820152815260200190600190039081611fac5790505b50905060005b865181101561204657612021878281518110611ff757611ff76151ac565b6020026020010151878787878681518110612014576120146151ac565b6020026020010151612e9f565b828281518110612033576120336151ac565b6020908102919091010152600101611fd9565b5095945050505050565b600061205b8361327a565b8015610a8a5750610a8a83836132de565b336001600160a01b038216036120c45760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016107f2565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b806040015160ff1660000361214d576000604051631b3fab5160e11b81526004016107f29190615890565b60208082015160ff808216600090815260029093526040832060018101549293909283921690036121ba57606084015160018201805491151562010000027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff90921691909117905561220f565b6060840151600182015460ff620100009091041615159015151461220f576040517f87f6037c00000000000000000000000000000000000000000000000000000000815260ff841660048201526024016107f2565b60a08401518051601f60ff8216111561223e576001604051631b3fab5160e11b81526004016107f29190615890565b6122a4858560030180548060200260200160405190810160405280929190818152602001828054801561229a57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161227c575b5050505050613398565b8560600151156123d357612312858560020180548060200260200160405190810160405280929190818152602001828054801561229a576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161227c575050505050613398565b6080860151805161232c9060028701906020840190613d55565b5080516001850180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1661010060ff841690810291909117909155601f101561238c576002604051631b3fab5160e11b81526004016107f29190615890565b604088015161239c9060036158aa565b60ff168160ff16116123c4576003604051631b3fab5160e11b81526004016107f29190615890565b6123d087836001613401565b50505b6123df85836002613401565b81516123f49060038601906020850190613d55565b5060408681015160018501805460ff191660ff8316179055875180865560a089015192517fab8b1b57514019638d7b5ce9c638fe71366fe8e2be1c40a7a80f1733d0e9f5479361244d938a939260028b019291906158c6565b60405180910390a161245e85613581565b505050505050565b61246e613dc3565b835160005b81811015612673576000600188868460208110612492576124926151ac565b61249f91901a601b615668565b8985815181106124b1576124b16151ac565b60200260200101518986815181106124cb576124cb6151ac565b602002602001015160405160008152602001604052604051612509949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa15801561252b573d6000803e3d6000fd5b505060408051601f1981015160ff808e166000908152600360209081528582206001600160a01b0385168352815285822085870190965285548084168652939750909550929392840191610100900416600281111561258c5761258c614200565b600281111561259d5761259d614200565b90525090506001816020015160028111156125ba576125ba614200565b146125f1576040517fca31867a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051859060ff16601f8110612608576126086151ac565b602002015115612644576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600185826000015160ff16601f811061265f5761265f6151ac565b911515602090920201525050600101612473565b5050505050505050565b81516126888161146f565b600061269382611571565b60208501515190915060008190036126d6576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8460400151518114612714576040517f57e0e08300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff81111561272f5761272f613ff9565b604051908082528060200260200182016040528015612758578160200160208202803683370190505b50905060005b828110156128cd5760008760200151828151811061277e5761277e6151ac565b602002602001015190507f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681600001516040015167ffffffffffffffff161461281157805160409081015190517f38432a2200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016107f2565b6128a78186600101805461282490615477565b80601f016020809104026020016040519081016040528092919081815260200182805461285090615477565b801561289d5780601f106128725761010080835404028352916020019161289d565b820191906000526020600020905b81548152906001019060200180831161288057829003601f168201915b505050505061359d565b8383815181106128b9576128b96151ac565b60209081029190910101525060010161275e565b5060006128e4858389606001518a608001516136bf565b90508060000361292c576040517f7dd17a7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff861660048201526024016107f2565b8551151560005b848110156109b557600089602001518281518110612953576129536151ac565b60200260200101519050600061297189836000015160600151610a3d565b9050600281600381111561298757612987614200565b036129de578151606001516040805167ffffffffffffffff808d16825290921660208301527f3b575419319662b2a6f5e2467d84521517a3382b908eb3d557bb3fdb0c50e23c910160405180910390a15050612e97565b60008160038111156129f2576129f2614200565b1480612a0f57506003816003811115612a0d57612a0d614200565b145b612a60578151606001516040517f25507e7f00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808c16600483015290911660248201526044016107f2565b8315612b415760045460009074010000000000000000000000000000000000000000900463ffffffff16612a9487426153d7565b1190508080612ab457506003826003811115612ab257612ab2614200565b145b612af6576040517fa9cfc86200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b1660048201526024016107f2565b8a8481518110612b0857612b086151ac565b6020026020010151600014612b3b578a8481518110612b2957612b296151ac565b60200260200101518360800181815250505b50612ba7565b6000816003811115612b5557612b55614200565b14612ba7578151606001516040517f3ef2a99c00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808c16600483015290911660248201526044016107f2565b81516080015167ffffffffffffffff1615801590612bd657506000816003811115612bd457612bd4614200565b145b15612c9b5781516080015160208301516040517fe0e03cae0000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169263e0e03cae92612c4d928e92919060040161594c565b6020604051808303816000875af1158015612c6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c909190615638565b612c9b575050612e97565b60008b604001518481518110612cb357612cb36151ac565b6020026020010151905080518360a001515114612d17578251606001516040517f1cfe6d8b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808d16600483015290911660248201526044016107f2565b612d2b8a8460000151606001516001613715565b600080612d3885846137bd565b91509150612d4f8c86600001516060015184613715565b868015612d6d57506003826003811115612d6b57612d6b614200565b145b8015612d8b57506000846003811115612d8857612d88614200565b14155b15612dc8578451516040517f2b11b8d90000000000000000000000000000000000000000000000000000000081526107f291908390600401615979565b6003826003811115612ddc57612ddc614200565b14158015612dfc57506002826003811115612df957612df9614200565b14155b15612e3d578451606001516040517f926c5a3e0000000000000000000000000000000000000000000000000000000081526107f2918e918590600401615992565b8451805160609091015160405167ffffffffffffffff918216918f16907f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df290612e8990879087906159b8565b60405180910390a450505050505b600101612933565b60408051808201909152600080825260208201526000612ec28760200151613887565b6040517fbbe4f6db0000000000000000000000000000000000000000000000000000000081526001600160a01b0380831660048301529192506000917f0000000000000000000000000000000000000000000000000000000000000000169063bbe4f6db90602401602060405180830381865afa158015612f47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f6b91906159d8565b90506001600160a01b0381161580612fb35750612fb16001600160a01b0382167faff2afbf00000000000000000000000000000000000000000000000000000000612050565b155b15612ff5576040517fae9b4ce90000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016107f2565b6000806130dc633907753760e01b6040518061010001604052808c81526020018a67ffffffffffffffff1681526020018b6001600160a01b031681526020018d606001518152602001876001600160a01b031681526020018d6000015181526020018d6040015181526020018981525060405160240161307591906159f5565b60408051601f198184030181529190526020810180516001600160e01b03167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152600454859063ffffffff600160e01b9091041661138860846138e2565b50915091508161311a57806040517fe1cd55090000000000000000000000000000000000000000000000000000000081526004016107f29190613e47565b80516020146131625780516040517f78ef80240000000000000000000000000000000000000000000000000000000081526020600482015260248101919091526044016107f2565b6000818060200190518101906131789190615acc565b604080516001600160a01b038c16602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b03167fa9059cbb000000000000000000000000000000000000000000000000000000001790526004549192506132109187907801000000000000000000000000000000000000000000000000900463ffffffff1661138860846138e2565b5090935091508261324f57816040517fe1cd55090000000000000000000000000000000000000000000000000000000081526004016107f29190613e47565b604080518082019091526001600160a01b039095168552602085015250919250505095945050505050565b60006132a6827f01ffc9a7000000000000000000000000000000000000000000000000000000006132de565b8015610a8d57506132d7827fffffffff000000000000000000000000000000000000000000000000000000006132de565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180516001600160e01b03167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015613381575060208210155b801561338d5750600081115b979650505050505050565b60005b8151811015610cf15760ff8316600090815260036020526040812083519091908490849081106133cd576133cd6151ac565b6020908102919091018101516001600160a01b03168252810191909152604001600020805461ffff1916905560010161339b565b60005b82518160ff161015610a23576000838260ff1681518110613427576134276151ac565b602002602001015190506000600281111561344457613444614200565b60ff80871660009081526003602090815260408083206001600160a01b0387168452909152902054610100900416600281111561348357613483614200565b146134a4576004604051631b3fab5160e11b81526004016107f29190615890565b6001600160a01b0381166134e4576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405180604001604052808360ff16815260200184600281111561350a5761350a614200565b905260ff80871660009081526003602090815260408083206001600160a01b0387168452825290912083518154931660ff198416811782559184015190929091839161ffff19161761010083600281111561356757613567614200565b0217905550905050508061357a90615ae5565b9050613404565b60ff8116610a3a57600a805467ffffffffffffffff1916905550565b8151602080820151604092830151925160009384936135e3937f2425b0b9f9054c76ff151b0a175b18f37a4a4e82013a72e9f15c9caa095ed21f93909291889101615b04565b60408051601f1981840301815290829052805160209182012086518051888401516060808b0151908401516080808d0151950151959761362c9794969395929491939101615b37565b604051602081830303815290604052805190602001208560400151805190602001208660a001516040516020016136639190615c2e565b60408051601f198184030181528282528051602091820120908301969096528101939093526060830191909152608082015260a081019190915260c0015b60405160208183030381529060405280519060200120905092915050565b6000806136cd858585613a08565b90506136d8816113ae565b6136e657600091505061370d565b67ffffffffffffffff86166000908152600960209081526040808320938352929052205490505b949350505050565b60006002613724608085615400565b67ffffffffffffffff166137389190615427565b905060006137468585611c1e565b905081613755600160046153d7565b901b19168183600381111561376c5761376c614200565b67ffffffffffffffff871660009081526008602052604081209190921b9290921791829161379b6080886156a5565b67ffffffffffffffff1681526020810191909152604001600020555050505050565b6040517fece670b6000000000000000000000000000000000000000000000000000000008152600090606090309063ece670b6906138019087908790600401615c8e565b600060405180830381600087803b15801561381b57600080fd5b505af192505050801561382c575060015b61386b573d80801561385a576040519150601f19603f3d011682016040523d82523d6000602084013e61385f565b606091505b50600392509050613880565b50506040805160208101909152600081526002905b9250929050565b600081516020146138c657816040517f8d666f600000000000000000000000000000000000000000000000000000000081526004016107f29190613e47565b610a8d828060200190518101906138dd9190615acc565b613ca7565b6000606060008361ffff1667ffffffffffffffff81111561390557613905613ff9565b6040519080825280601f01601f19166020018201604052801561392f576020820181803683370190505b509150863b613962577f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b5a85811015613995577fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b85900360408104810387106139ce577f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b505a6000808a5160208c0160008c8cf193505a900390503d848111156139f15750835b808352806000602085013e50955095509592505050565b8251825160009190818303613a49576040517f11a6b26400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6101018211801590613a5d57506101018111155b613a7a576040516309bde33960e01b815260040160405180910390fd5b60001982820101610100811115613aa4576040516309bde33960e01b815260040160405180910390fd5b80600003613ad15786600081518110613abf57613abf6151ac565b60200260200101519350505050613ca0565b60008167ffffffffffffffff811115613aec57613aec613ff9565b604051908082528060200260200182016040528015613b15578160200160208202803683370190505b50905060008080805b85811015613c3f5760006001821b8b811603613b795788851015613b62578c5160018601958e918110613b5357613b536151ac565b60200260200101519050613b9b565b8551600185019487918110613b5357613b536151ac565b8b5160018401938d918110613b9057613b906151ac565b602002602001015190505b600089861015613bcb578d5160018701968f918110613bbc57613bbc6151ac565b60200260200101519050613bed565b8651600186019588918110613be257613be26151ac565b602002602001015190505b82851115613c0e576040516309bde33960e01b815260040160405180910390fd5b613c188282613d14565b878481518110613c2a57613c2a6151ac565b60209081029190910101525050600101613b1e565b506001850382148015613c5157508683145b8015613c5c57508581145b613c79576040516309bde33960e01b815260040160405180910390fd5b836001860381518110613c8e57613c8e6151ac565b60200260200101519750505050505050505b9392505050565b60006001600160a01b03821180613cbf575061040082105b15613d105760408051602081018490520160408051601f19818403018152908290527f8d666f600000000000000000000000000000000000000000000000000000000082526107f291600401613e47565b5090565b6000818310613d2c57613d278284613d32565b610a8a565b610a8a83835b6040805160016020820152908101839052606081018290526000906080016136a1565b828054828255906000526020600020908101928215613db7579160200282015b82811115613db7578251825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909116178255602090920191600190910190613d75565b50613d10929150613de2565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115613d105760008155600101613de3565b60005b83811015613e12578181015183820152602001613dfa565b50506000910152565b60008151808452613e33816020860160208601613df7565b601f01601f19169290920160200192915050565b602081526000610a8a6020830184613e1b565b8060608101831015610a8d57600080fd5b60008083601f840112613e7d57600080fd5b50813567ffffffffffffffff811115613e9557600080fd5b60208301915083602082850101111561388057600080fd5b60008083601f840112613ebf57600080fd5b50813567ffffffffffffffff811115613ed757600080fd5b6020830191508360208260051b850101111561388057600080fd5b60008060008060008060008060e0898b031215613f0e57600080fd5b613f188a8a613e5a565b9750606089013567ffffffffffffffff80821115613f3557600080fd5b613f418c838d01613e6b565b909950975060808b0135915080821115613f5a57600080fd5b613f668c838d01613ead565b909750955060a08b0135915080821115613f7f57600080fd5b50613f8c8b828c01613ead565b999c989b50969995989497949560c00135949350505050565b600080600060808486031215613fba57600080fd5b613fc48585613e5a565b9250606084013567ffffffffffffffff811115613fe057600080fd5b613fec86828701613e6b565b9497909650939450505050565b634e487b7160e01b600052604160045260246000fd5b60405160c0810167ffffffffffffffff8111828210171561403257614032613ff9565b60405290565b60405160a0810167ffffffffffffffff8111828210171561403257614032613ff9565b6040516080810167ffffffffffffffff8111828210171561403257614032613ff9565b6040516060810167ffffffffffffffff8111828210171561403257614032613ff9565b6040805190810167ffffffffffffffff8111828210171561403257614032613ff9565b604051601f8201601f1916810167ffffffffffffffff811182821017156140ed576140ed613ff9565b604052919050565b6001600160a01b0381168114610a3a57600080fd5b8035614115816140f5565b919050565b803563ffffffff8116811461411557600080fd5b600060c0828403121561414057600080fd5b61414861400f565b8235614153816140f5565b81526141616020840161411a565b60208201526141726040840161411a565b60408201526141836060840161411a565b60608201526080830135614196816140f5565b608082015260a08301356141a9816140f5565b60a08201529392505050565b803567ffffffffffffffff8116811461411557600080fd5b600080604083850312156141e057600080fd5b6141e9836141b5565b91506141f7602084016141b5565b90509250929050565b634e487b7160e01b600052602160045260246000fd5b6004811061422657614226614200565b9052565b60208101610a8d8284614216565b600067ffffffffffffffff82111561425257614252613ff9565b5060051b60200190565b600060a0828403121561426e57600080fd5b614276614038565b905081358152614288602083016141b5565b6020820152614299604083016141b5565b60408201526142aa606083016141b5565b60608201526142bb608083016141b5565b608082015292915050565b600067ffffffffffffffff8211156142e0576142e0613ff9565b50601f01601f191660200190565b600082601f8301126142ff57600080fd5b813561431261430d826142c6565b6140c4565b81815284602083860101111561432757600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261435557600080fd5b8135602061436561430d83614238565b82815260059290921b8401810191818101908684111561438457600080fd5b8286015b8481101561445a57803567ffffffffffffffff808211156143a95760008081fd5b8189019150608080601f19848d030112156143c45760008081fd5b6143cc61405b565b87840135838111156143de5760008081fd5b6143ec8d8a838801016142ee565b825250604080850135848111156144035760008081fd5b6144118e8b838901016142ee565b8a840152506060808601358581111561442a5760008081fd5b6144388f8c838a01016142ee565b9284019290925294909201359381019390935250508352918301918301614388565b509695505050505050565b6000610140828403121561447857600080fd5b61448061400f565b905061448c838361425c565b815260a082013567ffffffffffffffff808211156144a957600080fd5b6144b5858386016142ee565b602084015260c08401359150808211156144ce57600080fd5b6144da858386016142ee565b60408401526144eb60e0850161410a565b6060840152610100840135608084015261012084013591508082111561451057600080fd5b5061451d84828501614344565b60a08301525092915050565b600082601f83011261453a57600080fd5b8135602061454a61430d83614238565b82815260059290921b8401810191818101908684111561456957600080fd5b8286015b8481101561445a57803567ffffffffffffffff81111561458d5760008081fd5b61459b8986838b0101614465565b84525091830191830161456d565b600082601f8301126145ba57600080fd5b813560206145ca61430d83614238565b82815260059290921b840181019181810190868411156145e957600080fd5b8286015b8481101561445a57803567ffffffffffffffff81111561460d5760008081fd5b61461b8986838b01016142ee565b8452509183019183016145ed565b600082601f83011261463a57600080fd5b8135602061464a61430d83614238565b82815260059290921b8401810191818101908684111561466957600080fd5b8286015b8481101561445a57803567ffffffffffffffff81111561468d5760008081fd5b61469b8986838b01016145a9565b84525091830191830161466d565b600082601f8301126146ba57600080fd5b813560206146ca61430d83614238565b8083825260208201915060208460051b8701019350868411156146ec57600080fd5b602086015b8481101561445a57803583529183019183016146f1565b600082601f83011261471957600080fd5b8135602061472961430d83614238565b82815260059290921b8401810191818101908684111561474857600080fd5b8286015b8481101561445a57803567ffffffffffffffff8082111561476d5760008081fd5b818901915060a080601f19848d030112156147885760008081fd5b614790614038565b61479b8885016141b5565b8152604080850135848111156147b15760008081fd5b6147bf8e8b83890101614529565b8a84015250606080860135858111156147d85760008081fd5b6147e68f8c838a0101614629565b83850152506080915081860135858111156148015760008081fd5b61480f8f8c838a01016146a9565b918401919091525091909301359083015250835291830191830161474c565b600080604080848603121561484257600080fd5b833567ffffffffffffffff8082111561485a57600080fd5b61486687838801614708565b945060209150818601358181111561487d57600080fd5b8601601f8101881361488e57600080fd5b803561489c61430d82614238565b81815260059190911b8201840190848101908a8311156148bb57600080fd5b8584015b83811015614947578035868111156148d75760008081fd5b8501603f81018d136148e95760008081fd5b878101356148f961430d82614238565b81815260059190911b82018a0190898101908f8311156149195760008081fd5b928b01925b828410156149375783358252928a0192908a019061491e565b86525050509186019186016148bf565b50809750505050505050509250929050565b60006020828403121561496b57600080fd5b813567ffffffffffffffff81111561498257600080fd5b820160a08185031215613ca057600080fd5b8015158114610a3a57600080fd5b803561411581614994565b600060208083850312156149c057600080fd5b823567ffffffffffffffff808211156149d857600080fd5b818501915085601f8301126149ec57600080fd5b81356149fa61430d82614238565b81815260059190911b83018401908481019088831115614a1957600080fd5b8585015b83811015614aa957803585811115614a355760008081fd5b86016060818c03601f1901811315614a4d5760008081fd5b614a5561407e565b614a608a84016141b5565b8152604080840135614a7181614994565b828c0152918301359188831115614a885760008081fd5b614a968e8c858701016142ee565b9082015285525050918601918601614a1d565b5098975050505050505050565b803560ff8116811461411557600080fd5b600060208284031215614ad957600080fd5b610a8a82614ab6565b60008151808452602080850194506020840160005b83811015614b1c5781516001600160a01b031687529582019590820190600101614af7565b509495945050505050565b60208152600082518051602084015260ff602082015116604084015260ff604082015116606084015260608101511515608084015250602083015160c060a0840152614b7660e0840182614ae2565b90506040840151601f198483030160c0850152614b938282614ae2565b95945050505050565b60008060408385031215614baf57600080fd5b614bb8836141b5565b946020939093013593505050565b60008060208385031215614bd957600080fd5b823567ffffffffffffffff80821115614bf157600080fd5b818501915085601f830112614c0557600080fd5b813581811115614c1457600080fd5b8660208260061b8501011115614c2957600080fd5b60209290920196919550909350505050565b600060208284031215614c4d57600080fd5b610a8a826141b5565b6020815281511515602082015267ffffffffffffffff60208301511660408201526000604083015160608084015261370d6080840182613e1b565b60008060408385031215614ca457600080fd5b823567ffffffffffffffff80821115614cbc57600080fd5b614cc886838701614465565b93506020850135915080821115614cde57600080fd5b50614ceb858286016145a9565b9150509250929050565b600060208284031215614d0757600080fd5b8135613ca0816140f5565b600082601f830112614d2357600080fd5b81356020614d3361430d83614238565b8083825260208201915060208460051b870101935086841115614d5557600080fd5b602086015b8481101561445a578035614d6d816140f5565b8352918301918301614d5a565b60006020808385031215614d8d57600080fd5b823567ffffffffffffffff80821115614da557600080fd5b818501915085601f830112614db957600080fd5b8135614dc761430d82614238565b81815260059190911b83018401908481019088831115614de657600080fd5b8585015b83811015614aa957803585811115614e0157600080fd5b860160c0818c03601f19011215614e185760008081fd5b614e2061400f565b8882013581526040614e33818401614ab6565b8a8301526060614e44818501614ab6565b8284015260809150614e578285016149a2565b9083015260a08381013589811115614e6f5760008081fd5b614e7d8f8d83880101614d12565b838501525060c0840135915088821115614e975760008081fd5b614ea58e8c84870101614d12565b9083015250845250918601918601614dea565b600060208284031215614eca57600080fd5b5035919050565b80356001600160e01b038116811461411557600080fd5b600082601f830112614ef957600080fd5b81356020614f0961430d83614238565b82815260069290921b84018101918181019086841115614f2857600080fd5b8286015b8481101561445a5760408189031215614f455760008081fd5b614f4d6140a1565b614f56826141b5565b8152614f63858301614ed1565b81860152835291830191604001614f2c565b600082601f830112614f8657600080fd5b81356020614f9661430d83614238565b82815260079290921b84018101918181019086841115614fb557600080fd5b8286015b8481101561445a578088036080811215614fd35760008081fd5b614fdb61407e565b614fe4836141b5565b8152604080601f1984011215614ffa5760008081fd5b6150026140a1565b925061500f8785016141b5565b835261501c8185016141b5565b8388015281870192909252606083013591810191909152835291830191608001614fb9565b6000602080838503121561505457600080fd5b823567ffffffffffffffff8082111561506c57600080fd5b8185019150604080838803121561508257600080fd5b61508a6140a1565b83358381111561509957600080fd5b84016040818a0312156150ab57600080fd5b6150b36140a1565b8135858111156150c257600080fd5b8201601f81018b136150d357600080fd5b80356150e161430d82614238565b81815260069190911b8201890190898101908d83111561510057600080fd5b928a01925b828410156151505787848f03121561511d5760008081fd5b6151256140a1565b8435615130816140f5565b815261513d858d01614ed1565b818d0152825292870192908a0190615105565b84525050508187013593508484111561516857600080fd5b6151748a858401614ee8565b818801528252508385013591508282111561518e57600080fd5b61519a88838601614f75565b85820152809550505050505092915050565b634e487b7160e01b600052603260045260246000fd5b805160408084528151848201819052600092602091908201906060870190855b8181101561521957835180516001600160a01b031684528501516001600160e01b03168584015292840192918501916001016151e2565b50508583015187820388850152805180835290840192506000918401905b80831015615273578351805167ffffffffffffffff1683528501516001600160e01b031685830152928401926001929092019190850190615237565b50979650505050505050565b602081526000610a8a60208301846151c2565b67ffffffffffffffff8316815260608101613ca06020830184805167ffffffffffffffff908116835260209182015116910152565b634e487b7160e01b600052601160045260246000fd5b67ffffffffffffffff8181168382160190808211156152fe576152fe6152c7565b5092915050565b60006020808352606084516040808487015261532460608701836151c2565b87850151878203601f19016040890152805180835290860193506000918601905b80831015614aa957845167ffffffffffffffff81511683528781015161538489850182805167ffffffffffffffff908116835260209182015116910152565b50840151828701529386019360019290920191608090910190615345565b6000602082840312156153b457600080fd5b813567ffffffffffffffff8111156153cb57600080fd5b61370d84828501614708565b81810381811115610a8d57610a8d6152c7565b634e487b7160e01b600052601260045260246000fd5b600067ffffffffffffffff8084168061541b5761541b6153ea565b92169190910692915050565b8082028115828204841417610a8d57610a8d6152c7565b60006040828403121561545057600080fd5b6154586140a1565b615461836141b5565b8152602083013560208201528091505092915050565b600181811c9082168061548b57607f821691505b6020821081036154ab57634e487b7160e01b600052602260045260246000fd5b50919050565b805182526000602067ffffffffffffffff81840151168185015260408084015160a060408701526154e560a0870182613e1b565b9050606085015186820360608801526154fe8282613e1b565b608087810151898303918a01919091528051808352908601935060009250908501905b8083101561527357835180516001600160a01b0316835286015186830152928501926001929092019190840190615521565b602081526000610a8a60208301846154b1565b60808152600061557960808301876154b1565b61ffff9590951660208301525060408101929092526001600160a01b0316606090910152919050565b6000806000606084860312156155b757600080fd5b83516155c281614994565b602085015190935067ffffffffffffffff8111156155df57600080fd5b8401601f810186136155f057600080fd5b80516155fe61430d826142c6565b81815287602083850101111561561357600080fd5b615624826020830160208601613df7565b809450505050604084015190509250925092565b60006020828403121561564a57600080fd5b8151613ca081614994565b80820180821115610a8d57610a8d6152c7565b60ff8181168382160190811115610a8d57610a8d6152c7565b8183823760009101908152919050565b828152606082602083013760800192915050565b600067ffffffffffffffff808416806156c0576156c06153ea565b92169190910492915050565b601f821115610cf1576000816000526020600020601f850160051c810160208610156156f55750805b601f850160051c820191505b8181101561245e57828155600101615701565b815167ffffffffffffffff81111561572e5761572e613ff9565b6157428161573c8454615477565b846156cc565b602080601f831160018114615777576000841561575f5750858301515b600019600386901b1c1916600185901b17855561245e565b600085815260208120601f198616915b828110156157a657888601518255948401946001909101908401615787565b50858210156157c45787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60006020808352835460ff81161515602085015267ffffffffffffffff8160081c1660408501525060018085016060808601526000815461581481615477565b80608089015260a06001831660008114615835576001811461585157615881565b60ff19841660a08b015260a083151560051b8b01019450615881565b85600052602060002060005b848110156158785781548c820185015290880190890161585d565b8b0160a0019550505b50929998505050505050505050565b60208101600583106158a4576158a4614200565b91905290565b60ff81811683821602908116908181146152fe576152fe6152c7565b600060a0820160ff88168352602087602085015260a0604085015281875480845260c086019150886000526020600020935060005b818110156159205784546001600160a01b0316835260019485019492840192016158fb565b505084810360608601526159348188614ae2565b935050505060ff831660808301529695505050505050565b600067ffffffffffffffff808616835280851660208401525060606040830152614b936060830184613e1b565b82815260406020820152600061370d6040830184613e1b565b67ffffffffffffffff8481168252831660208201526060810161370d6040830184614216565b6159c28184614216565b60406020820152600061370d6040830184613e1b565b6000602082840312156159ea57600080fd5b8151613ca0816140f5565b6020815260008251610100806020850152615a14610120850183613e1b565b91506020850151615a31604086018267ffffffffffffffff169052565b5060408501516001600160a01b038116606086015250606085015160808501526080850151615a6b60a08601826001600160a01b03169052565b5060a0850151601f19808685030160c0870152615a888483613e1b565b935060c08701519150808685030160e0870152615aa58483613e1b565b935060e0870151915080868503018387015250615ac28382613e1b565b9695505050505050565b600060208284031215615ade57600080fd5b5051919050565b600060ff821660ff8103615afb57615afb6152c7565b60010192915050565b848152600067ffffffffffffffff808616602084015280851660408401525060806060830152615ac26080830184613e1b565b86815260c060208201526000615b5060c0830188613e1b565b6001600160a01b039690961660408301525067ffffffffffffffff9384166060820152608081019290925290911660a09091015292915050565b600082825180855260208086019550808260051b84010181860160005b84811015615c2157601f19868403018952815160808151818652615bcd82870182613e1b565b9150508582015185820387870152615be58282613e1b565b91505060408083015186830382880152615bff8382613e1b565b6060948501519790940196909652505098840198925090830190600101615ba7565b5090979650505050505050565b602081526000610a8a6020830184615b8a565b60008282518085526020808601955060208260051b8401016020860160005b84811015615c2157601f19868403018952615c7c838351613e1b565b98840198925090830190600101615c60565b604081526000835180516040840152602081015167ffffffffffffffff80821660608601528060408401511660808601528060608401511660a08601528060808401511660c086015250505060208401516101408060e0850152615cf6610180850183613e1b565b915060408601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08086850301610100870152615d338483613e1b565b935060608801519150615d526101208701836001600160a01b03169052565b60808801518387015260a0880151925080868503016101608701525050615d798282615b8a565b9150508281036020840152614b938185615c4156fea164736f6c6343000818000a",
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
