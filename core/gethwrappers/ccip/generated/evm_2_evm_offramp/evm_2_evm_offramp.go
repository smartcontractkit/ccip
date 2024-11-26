// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package evm_2_evm_offramp

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
	zkSyncAccounts "github.com/zksync-sdk/zksync2-go/accounts"
	zkSyncClient "github.com/zksync-sdk/zksync2-go/clients"
	zktypes "github.com/zksync-sdk/zksync2-go/types"
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
	MaxDataBytes                            uint32
	MaxNumberOfTokensPerMsg                 uint16
	Router                                  common.Address
	PriceRegistry                           common.Address
}

type EVM2EVMOffRampGasLimitOverride struct {
	ReceiverExecutionGasLimit *big.Int
	TokenGasOverrides         []uint32
}

type EVM2EVMOffRampRateLimitToken struct {
	SourceToken common.Address
	DestToken   common.Address
}

type EVM2EVMOffRampStaticConfig struct {
	CommitStore         common.Address
	ChainSelector       uint64
	SourceChainSelector uint64
	OnRamp              common.Address
	PrevOffRamp         common.Address
	RmnProxy            common.Address
	TokenAdminRegistry  common.Address
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitStoreAlreadyInUse\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ConfigDigestMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"DestinationGasAmountCountMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyReport\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"ForkedChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumOCR2BaseNoChecks.InvalidConfigErrorType\",\"name\":\"errorType\",\"type\":\"uint8\"}],\"name\":\"InvalidConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidEVMAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"oldLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"InvalidManualExecutionGasLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMessageId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"InvalidNewState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokenIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenGasOverride\",\"type\":\"uint256\"}],\"name\":\"InvalidTokenGasOverride\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionGasLimitMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notPool\",\"type\":\"address\"}],\"name\":\"NotACompatiblePool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountReleased\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balancePre\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balancePost\",\"type\":\"uint256\"}],\"name\":\"ReleaseOrMintBalanceMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotCommitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"TokenDataMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"TokenHandlingError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedTransmitter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedTokenData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"WrongMessageLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyAttempted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"staticConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structEVM2EVMOffRamp.DynamicConfig\",\"name\":\"dynamicConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"SkippedAlreadyExecutedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedIncorrectNonce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SkippedSenderWithPreviousRampMessageInflight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"TokenAggregateRateLimitAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"name\":\"TokenAggregateRateLimitRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensConsumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\"}],\"internalType\":\"structClient.Any2EVMMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint32[]\",\"name\":\"tokenGasOverrides\",\"type\":\"uint32[]\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRateLimitTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"destTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDynamicConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxDataBytes\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOffRamp.DynamicConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumInternal.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSenderNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"commitStore\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"onRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prevOffRamp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAdminRegistry\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOffRamp.StaticConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"strict\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structClient.EVMTokenAmount[]\",\"name\":\"tokenAmounts\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sourceTokenData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"internalType\":\"structInternal.EVM2EVMMessage[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[][]\",\"name\":\"offchainTokenData\",\"type\":\"bytes[][]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"proofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structInternal.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"receiverExecutionGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint32[]\",\"name\":\"tokenGasOverrides\",\"type\":\"uint32[]\"}],\"internalType\":\"structEVM2EVMOffRamp.GasLimitOverride[]\",\"name\":\"gasLimitOverrides\",\"type\":\"tuple[]\"}],\"name\":\"manuallyExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setOCR2Config\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOffRamp.RateLimitToken[]\",\"name\":\"removes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"}],\"internalType\":\"structEVM2EVMOffRamp.RateLimitToken[]\",\"name\":\"adds\",\"type\":\"tuple[]\"}],\"name\":\"updateRateLimitTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101a06040523480156200001257600080fd5b506040516200660a3803806200660a8339810160408190526200003591620004ec565b8033806000816200008d5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c057620000c081620002ca565b50506040805160a081018252602084810180516001600160801b039081168085524263ffffffff169385018490528751151585870181905292518216606080870182905298909601519091166080948501819052600380546001600160a01b031916909217600160801b9485021760ff60a01b1916600160a01b90930292909217905502909117600455469052508201516001600160a01b031615806200016f575081516001600160a01b0316155b8062000186575060c08201516001600160a01b0316155b15620001a5576040516342bcdf7f60e11b815260040160405180910390fd5b81600001516001600160a01b0316634120fccd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015620001e8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200020e9190620005b5565b6001600160401b03166001146200023857604051636fc2a20760e11b815260040160405180910390fd5b81516001600160a01b0390811660a090815260408401516001600160401b0390811660c0908152602086015190911660e05260608501518316610100526080850151831661014052908401518216610160528301511661018052620002bd7f8acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b362000375565b6101205250620005da9050565b336001600160a01b03821603620003245760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000084565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60008160c05160e05161010051604051602001620003bf94939291909384526001600160401b039283166020850152911660408301526001600160a01b0316606082015260800190565b604051602081830303815290604052805190602001209050919050565b60405160e081016001600160401b03811182821017156200040d57634e487b7160e01b600052604160045260246000fd5b60405290565b80516001600160a01b03811681146200042b57600080fd5b919050565b80516001600160401b03811681146200042b57600080fd5b80516001600160801b03811681146200042b57600080fd5b6000606082840312156200047357600080fd5b604051606081016001600160401b0381118282101715620004a457634e487b7160e01b600052604160045260246000fd5b806040525080915082518015158114620004bd57600080fd5b8152620004cd6020840162000448565b6020820152620004e06040840162000448565b60408201525092915050565b6000808284036101408112156200050257600080fd5b60e08112156200051157600080fd5b506200051c620003dc565b620005278462000413565b8152620005376020850162000430565b60208201526200054a6040850162000430565b60408201526200055d6060850162000413565b6060820152620005706080850162000413565b60808201526200058360a0850162000413565b60a08201526200059660c0850162000413565b60c08201529150620005ac8460e0850162000460565b90509250929050565b600060208284031215620005c857600080fd5b620005d38262000430565b9392505050565b60805160a05160c05160e0516101005161012051610140516101605161018051615f41620006c9600039600081816102ec01528181611c6a01526133920152600081816102bd01528181611c420152611f2701526000818161028e01528181610d8d01528181610df201528181611c18015281816124a4015261250e015260006120c601526000818161025f0152611bee0152600081816101ff0152611b9201526000818161022f01528181611bc601528181611ee40152818161303b01526134bf0152600081816101d001528181611b6d01526121ac015260008181611e3e0152611e8a0152615f416000f3fe608060405234801561001057600080fd5b50600436106101825760003560e01c806385572ffb116100d8578063afcb95d71161008c578063c92b283211610066578063c92b2832146105f3578063f077b59214610606578063f2fde38b1461061c57600080fd5b8063afcb95d7146105ad578063b1dc65a4146105cd578063b6113fce146105e057600080fd5b8063873504d7116100bd578063873504d7146105765780638926c4ee146105895780638da5cb5b1461059c57600080fd5b806385572ffb1461053c578063856c82471461054a57600080fd5b8063599f64311161013a5780637437ff9f116101145780637437ff9f1461046157806379ba50971461050457806381ff70481461050c57600080fd5b8063599f643114610414578063666cab8d14610439578063704b6c021461044e57600080fd5b8063181f5a771161016b578063181f5a77146103525780631ef381741461039b578063546719cd146103b057600080fd5b806306285c6914610187578063142a98fc14610332575b600080fd5b61031c6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526040518060e001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815250905090565b6040516103299190614312565b60405180910390f35b6103456103403660046143a8565b61062f565b6040516103299190614408565b61038e6040518060400160405280601481526020017f45564d3245564d4f666652616d7020312e352e3000000000000000000000000081525081565b6040516103299190614466565b6103ae6103a936600461468f565b6106aa565b005b6103b8610a9e565b604051610329919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610329565b610441610b53565b60405161032991906147a1565b6103ae61045c3660046147b4565b610bb5565b6104f76040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a081018252600a5463ffffffff8082168352640100000000820416602083015268010000000000000000810461ffff16928201929092526a01000000000000000000009091046001600160a01b039081166060830152600b5416608082015290565b60405161032991906147d1565b6103ae610c7e565b6007546005546040805163ffffffff80851682526401000000009094049093166020840152820152606001610329565b6103ae610182366004614827565b61055d6105583660046147b4565b610d61565b60405167ffffffffffffffff9091168152602001610329565b6103ae6105843660046148f3565b610e64565b6103ae610597366004614e1e565b611037565b6000546001600160a01b0316610421565b604080516001815260006020820181905291810191909152606001610329565b6103ae6105db366004614f75565b61129c565b6103ae6105ee36600461505a565b6114a7565b6103ae610601366004615111565b6117e7565b61060e611852565b60405161032992919061517f565b6103ae61062a3660046147b4565b611978565b600061063d600160046151d3565b600261064a608085615215565b67ffffffffffffffff1661065e919061523c565b6010600061066d608087615253565b67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002054901c1660038111156106a4576106a46143c5565b92915050565b84518460ff16601f8211156106f75760016040517f367f56a20000000000000000000000000000000000000000000000000000000081526004016106ee919061527a565b60405180910390fd5b806000036107345760006040517f367f56a20000000000000000000000000000000000000000000000000000000081526004016106ee919061527a565b61073c611989565b610745856119ff565b60095460005b818110156107bc57600860006009838154811061076a5761076a615294565b60009182526020808320909101546001600160a01b03168352820192909252604001902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905560010161074b565b5050865160005b8181101561095f5760008982815181106107df576107df615294565b60200260200101519050600060028111156107fc576107fc6143c5565b6001600160a01b038216600090815260086020526040902054610100900460ff16600281111561082e5761082e6143c5565b146108685760026040517f367f56a20000000000000000000000000000000000000000000000000000000081526004016106ee919061527a565b6001600160a01b0381166108a8576040517fd6c62c9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805180820190915260ff8316815260208101600290526001600160a01b03821660009081526008602090815260409091208251815460ff9091167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082168117835592840151919283917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000161761010083600281111561094b5761094b6143c5565b0217905550905050508060010190506107c3565b5087516109739060099060208b0190614280565b506006805460ff838116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909216908a1617179055600780546109f99146913091906000906109cb9063ffffffff166152c3565b91906101000a81548163ffffffff021916908363ffffffff160217905563ffffffff168c8c8c8c8c8c611cc9565b6005819055600780544363ffffffff9081166401000000009081027fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff841681179094556040519083048216947f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0594610a8a9487949293918316921691909117908f908f908f908f908f908f906152e6565b60405180910390a150505050505050505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040805160a0810182526003546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff1660208501527401000000000000000000000000000000000000000090920460ff161515938301939093526004548084166060840152049091166080820152610b4e90611d56565b905090565b60606009805480602002602001604051908101604052809291908181526020018280548015610bab57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610b8d575b5050505050905090565b6000546001600160a01b03163314801590610bdb57506002546001600160a01b03163314155b15610c12576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383169081179091556040519081527f8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c9060200160405180910390a150565b6001546001600160a01b03163314610cf2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016106ee565b60008054337fffffffffffffffffffffffff0000000000000000000000000000000000000000808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6001600160a01b0381166000908152600f602052604081205467ffffffffffffffff168082036106a4577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316156106a4576040517f856c82470000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063856c824790602401602060405180830381865afa158015610e39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5d919061537c565b9392505050565b610e6c611989565b60005b8251811015610f3f57610ea9838281518110610e8d57610e8d615294565b602002602001015160200151600c611e0890919063ffffffff16565b15610f37577fcbf3cbeaed4ac1d605ed30f4af06c35acaeff2379db7f6146c9cceee83d58782838281518110610ee157610ee1615294565b602002602001015160000151848381518110610eff57610eff615294565b602002602001015160200151604051610f2e9291906001600160a01b0392831681529116602082015260400190565b60405180910390a15b600101610e6f565b5060005b815181101561103257610f9c828281518110610f6157610f61615294565b602002602001015160200151838381518110610f7f57610f7f615294565b602002602001015160000151600c611e1d9092919063ffffffff16565b1561102a577ffc23abf7ddbd3c02b1420dafa2355c56c1a06fbb8723862ac14d6bd74177361a828281518110610fd457610fd4615294565b602002602001015160000151838381518110610ff257610ff2615294565b6020026020010151602001516040516110219291906001600160a01b0392831681529116602082015260400190565b60405180910390a15b600101610f43565b505050565b61103f611e3b565b8151518151811461107c576040517f83e3f56400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b818110156112915760008460000151828151811061109f5761109f615294565b6020026020010151905060008483815181106110bd576110bd615294565b60209081029190910101518051909150801561112c57826080015181101561112c5761018083015160808401516040517f9c6db58d00000000000000000000000000000000000000000000000000000000815260048101929092526024820152604481018290526064016106ee565b816020015151836101400151511461118e5761018083015160608401516040517f85d2e5bf000000000000000000000000000000000000000000000000000000008152600481019290925267ffffffffffffffff1660248201526044016106ee565b61016083015160005b846101400151518110156112815760008287815181106111b9576111b9615294565b60200260200101518060200190518101906111d491906153de565b90506000856020015183815181106111ee576111ee615294565b602002602001015163ffffffff169050806000141580156112185750816060015163ffffffff1681105b156112775761018087015160608301516040517fef0c635200000000000000000000000000000000000000000000000000000000815260048101929092526024820185905263ffffffff166044820152606481018290526084016106ee565b5050600101611197565b505050505080600101905061107f565b506110328383611ebc565b6112a6878761290f565b6005548835908082146112ef576040517f93df584c00000000000000000000000000000000000000000000000000000000815260048101829052602481018390526044016106ee565b6112f7611e3b565b6040805183815260208c81013560081c63ffffffff16908201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a13360009081526008602090815260408083208151808301909252805460ff8082168452929391929184019161010090910416600281111561137f5761137f6143c5565b6002811115611390576113906143c5565b90525090506002816020015160028111156113ad576113ad6143c5565b1480156113e757506009816000015160ff16815481106113cf576113cf615294565b6000918252602090912001546001600160a01b031633145b61141d576040517fda0f08e800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50600061142b85602061523c565b61143688602061523c565b6114428b6101446154aa565b61144c91906154aa565b61145691906154aa565b905036811461149a576040517f8e1192e1000000000000000000000000000000000000000000000000000000008152600481018290523660248201526044016106ee565b5050505050505050505050565b3330146114e0576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516000808252602082019092528161151d565b60408051808201909152600080825260208201528152602001906001900390816114f65790505b50905060006115306101408701876154bd565b905011156115a7576115a46115496101408701876154bd565b6115596040890160208a016147b4565b604080516001600160a01b0390921660208301520160408051601f1981840301815291815261158e9060608b01908b016147b4565b61159c6101608b018b615525565b8a8a8a612966565b90505b6115b561012086018661558d565b15905080156115c657506080850135155b806115e857506115dc60608601604087016147b4565b6001600160a01b03163b155b8061163357506116317f85572ffb0000000000000000000000000000000000000000000000000000000061162260608801604089016147b4565b6001600160a01b031690612b8f565b155b1561163e57506117e1565b600a546040805160a08101909152610180870135815260009182916a01000000000000000000009091046001600160a01b031690633cf979839060208082019061168a908c018c6143a8565b67ffffffffffffffff1681526020018a60200160208101906116ac91906147b4565b604080516001600160a01b0390921660208301520160408051601f1981840301815291905281526020016116e46101208c018c61558d565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200186905261138860808b013561173960608d0160408e016147b4565b6040518563ffffffff1660e01b81526004016117589493929190615637565b6000604051808303816000875af1158015611777573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261179f91908101906156fc565b5091509150816117dd57806040517f0a8d6e8c0000000000000000000000000000000000000000000000000000000081526004016106ee9190614466565b5050505b50505050565b6000546001600160a01b0316331480159061180d57506002546001600160a01b03163314155b15611844576040517ff6cd562000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61184f600382612bab565b50565b6060806000611861600c612d90565b90508067ffffffffffffffff81111561187c5761187c614479565b6040519080825280602002602001820160405280156118a5578160200160208202803683370190505b5092508067ffffffffffffffff8111156118c1576118c1614479565b6040519080825280602002602001820160405280156118ea578160200160208202803683370190505b50915060005b8181101561197257600080611906600c84612d9b565b915091508086848151811061191d5761191d615294565b60200260200101906001600160a01b031690816001600160a01b0316815250508185848151811061195057611950615294565b6001600160a01b039092166020928302919091019091015250506001016118f0565b50509091565b611980611989565b61184f81612db9565b6000546001600160a01b031633146119fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016106ee565b565b600081806020019051810190611a159190615756565b60608101519091506001600160a01b0316611a5c576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600a805460208085015160408087015160608089015163ffffffff9889167fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000909716969096176401000000009890941697909702929092177fffff00000000000000000000000000000000000000000000ffffffffffffffff166801000000000000000061ffff909316929092027fffff0000000000000000000000000000000000000000ffffffffffffffffffff16919091176a01000000000000000000006001600160a01b039485160217909355608080860151600b80547fffffffffffffffffffffffff000000000000000000000000000000000000000016918516919091179055835160e0810185527f0000000000000000000000000000000000000000000000000000000000000000841681527f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff908116938201939093527f0000000000000000000000000000000000000000000000000000000000000000909216828501527f00000000000000000000000000000000000000000000000000000000000000008316948201949094527f00000000000000000000000000000000000000000000000000000000000000008216938101939093527f0000000000000000000000000000000000000000000000000000000000000000811660a08401527f00000000000000000000000000000000000000000000000000000000000000001660c0830152517f7879e20bb60a503429de4a2c912b5904f08a39f2af054c10fb46434b5d61126091611cbd9184906157f5565b60405180910390a15050565b6000808a8a8a8a8a8a8a8a8a604051602001611ced999897969594939291906158b7565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152611de482606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff1642611dc891906151d3565b85608001516fffffffffffffffffffffffffffffffff16612e94565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b6000610e5d836001600160a01b038416612ebc565b6000611e33846001600160a01b03851684612ec8565b949350505050565b467f0000000000000000000000000000000000000000000000000000000000000000146119fd576040517f0f01ce850000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201524660248201526044016106ee565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060801b77ffffffffffffffff000000000000000000000000000000001660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632cbc26bb90602401602060405180830381865afa158015611f76573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f9a919061593f565b15611fd1576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815151600081900361200e576040517ebf199700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826020015151811461204c576040517f57e0e08300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff81111561206757612067614479565b604051908082528060200260200182016040528015612090578160200160208202803683370190505b50905060005b82811015612168576000856000015182815181106120b6576120b6615294565b602002602001015190506120ea817f0000000000000000000000000000000000000000000000000000000000000000612ede565b8383815181106120fc576120fc615294565b60200260200101818152505080610180015183838151811061212057612120615294565b60200260200101511461215f576040517f7185cf6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50600101612096565b508251604080860151606087015191517f32048875000000000000000000000000000000000000000000000000000000008152921515926000926001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016926332048875926121e29288929160040161598d565b602060405180830381865afa1580156121ff573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061222391906159c3565b90508060000361225f576040517fea75680100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b848110156117dd5760008760000151828151811061228257612282615294565b60200260200101519050600061229b826060015161062f565b905060008160038111156122b1576122b16143c5565b14806122ce575060038160038111156122cc576122cc6143c5565b145b61231457816060015167ffffffffffffffff167fe3dd0bec917c965a133ddb2c84874725ee1e2fd8d763c19efa36d6a11cd82b1f60405160405180910390a25050612907565b606085156123fa5788848151811061232e5761232e615294565b6020908102919091018101510151600a5490915060009063ffffffff1661235587426151d3565b119050808061237557506003836003811115612373576123736143c5565b145b6123ab576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8985815181106123bd576123bd615294565b6020026020010151600001516000146123f4578985815181106123e2576123e2615294565b60209081029190910101515160808501525b5061245e565b600082600381111561240e5761240e6143c5565b1461245e57606083015160405167ffffffffffffffff90911681527f67d9ba0f63d427c482c2736300e6d5a34c6691dbcdea8ad35828a1f1ba47e8729060200160405180910390a1505050612907565b60c083015167ffffffffffffffff16156126df576020808401516001600160a01b03166000908152600f909152604081205467ffffffffffffffff1690819003612649577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316156126495760208401516040517f856c82470000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201527f00000000000000000000000000000000000000000000000000000000000000009091169063856c824790602401602060405180830381865afa158015612557573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061257b919061537c565b60c085015190915067ffffffffffffffff166125988260016159dc565b67ffffffffffffffff16146125f95783602001516001600160a01b03168460c0015167ffffffffffffffff167fe44a20935573a783dd0d5991c92d7b6a0eb3173566530364db3ec10e9a990b5d60405160405180910390a350505050612907565b6020848101516001600160a01b03166000908152600f9091526040902080547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83161790555b600083600381111561265d5761265d6143c5565b036126dd5760c084015167ffffffffffffffff1661267c8260016159dc565b67ffffffffffffffff16146126dd5783602001516001600160a01b03168460c0015167ffffffffffffffff167fd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf4123760405160405180910390a350505050612907565b505b60008a6020015185815181106126f7576126f7615294565b602002602001015190506127238460600151856000015186610140015151876101200151518551613039565b612732846060015160016131ba565b600080612740868486613264565b915091506127528660600151836131ba565b88156127be57600382600381111561276c5761276c6143c5565b036127be576000856003811115612785576127856143c5565b146127be57806040517fcf19edfd0000000000000000000000000000000000000000000000000000000081526004016106ee9190614466565b60028260038111156127d2576127d26143c5565b1461282a5760038260038111156127eb576127eb6143c5565b1461282a578560600151826040517f9e2616030000000000000000000000000000000000000000000000000000000081526004016106ee9291906159fd565b60c086015167ffffffffffffffff16156128b2576000856003811115612852576128526143c5565b036128b2576020808701516001600160a01b03166000908152600f90915260408120805467ffffffffffffffff169161288a83615a1b565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505b856101800151866060015167ffffffffffffffff167fd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef6584846040516128f8929190615a38565b60405180910390a35050505050505b600101612262565b61296261291e82840184615a58565b604080516000808252602082019092529061295c565b6040805180820190915260008152606060208201528152602001906001900390816129345790505b50611ebc565b5050565b60608989808060200260200160405190810160405280939291908181526020016000905b828210156129b6576129a760408302860136819003810190615a8d565b8152602001906001019061298a565b505050505090506000805b8a811015612b715760008888838181106129dd576129dd615294565b90506020028101906129ef919061558d565b8101906129fc9190615aa9565b90508451600014612a5757848281518110612a1957612a19615294565b602002602001015163ffffffff16600014612a5757848281518110612a4057612a40615294565b602090810291909101015163ffffffff1660608201525b612ad78d8d84818110612a6c57612a6c615294565b905060400201602001358c8c848b8b88818110612a8b57612a8b615294565b9050602002810190612a9d919061558d565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061333192505050565b848381518110612ae957612ae9615294565b6020026020010181905250612b25848381518110612b0957612b09615294565b602002602001015160000151600c6136fd90919063ffffffff16565b15612b6857612b5b848381518110612b3f57612b3f615294565b6020908102919091010151600b546001600160a01b0316613712565b612b6590846154aa565b92505b506001016129c1565b508015612b8157612b8181613833565b509998505050505050505050565b6000612b9a83613840565b8015610e5d5750610e5d83836138a4565b8154600090612bd490700100000000000000000000000000000000900463ffffffff16426151d3565b90508015612c765760018301548354612c1c916fffffffffffffffffffffffffffffffff80821692811691859170010000000000000000000000000000000090910416612e94565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354612c9c916fffffffffffffffffffffffffffffffff9081169116613974565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990612d839084908151151581526020808301516fffffffffffffffffffffffffffffffff90811691830191909152604092830151169181019190915260600190565b60405180910390a1505050565b60006106a48261398a565b6000808080612daa8686613995565b909450925050505b9250929050565b336001600160a01b03821603612e2b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016106ee565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000612eb385612ea4848661523c565b612eae90876154aa565b613974565b95945050505050565b6000610e5d83836139a4565b6000611e3384846001600160a01b0385166139c1565b60008060001b8284602001518560400151866060015187608001518860a001518960c001518a60e001518b6101000151604051602001612f749897969594939291906001600160a01b039889168152968816602088015267ffffffffffffffff95861660408801526060870194909452911515608086015290921660a0840152921660c082015260e08101919091526101000190565b6040516020818303038152906040528051906020012085610120015180519060200120866101400151604051602001612fad9190615b65565b60405160208183030381529060405280519060200120876101600151604051602001612fd99190615bd2565b60408051601f198184030181528282528051602091820120908301979097528101949094526060840192909252608083015260a082015260c081019190915260e00160405160208183030381529060405280519060200120905092915050565b7f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff168467ffffffffffffffff16146130b2576040517f1279ec8a00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526024016106ee565b600a5468010000000000000000900461ffff1683111561310a576040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff861660048201526024016106ee565b80831461314f576040517f8808f8e700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff861660048201526024016106ee565b600a54640100000000900463ffffffff168211156131b357600a546040517f8693378900000000000000000000000000000000000000000000000000000000815264010000000090910463ffffffff166004820152602481018390526044016106ee565b5050505050565b600060026131c9608085615215565b67ffffffffffffffff166131dd919061523c565b905060006010816131ef608087615253565b67ffffffffffffffff168152602081019190915260400160002054905081613219600160046151d3565b901b191681836003811115613230576132306143c5565b901b178060106000613243608088615253565b67ffffffffffffffff16815260208101919091526040016000205550505050565b6040517fb6113fce000000000000000000000000000000000000000000000000000000008152600090606090309063b6113fce906132aa90889088908890600401615c1c565b600060405180830381600087803b1580156132c457600080fd5b505af19250505080156132d5575060015b613314573d808015613303576040519150601f19603f3d011682016040523d82523d6000602084013e613308565b606091505b50600392509050613329565b50506040805160208101909152600081526002905b935093915050565b6040805180820190915260008082526020820152600061335484602001516139de565b6040517fbbe4f6db0000000000000000000000000000000000000000000000000000000081526001600160a01b0380831660048301529192506000917f0000000000000000000000000000000000000000000000000000000000000000169063bbe4f6db90602401602060405180830381865afa1580156133d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133fd9190615da5565b90506001600160a01b038116158061344557506134436001600160a01b0382167faff2afbf00000000000000000000000000000000000000000000000000000000612b8f565b155b15613487576040517fae9b4ce90000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526024016106ee565b60008061349f8885896060015163ffffffff16613a84565b91509150600080600061359d6040518061010001604052808e81526020017f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff1681526020018d6001600160a01b031681526020018f8152602001896001600160a01b031681526020018c6000015181526020018c6040015181526020018b8152506040516024016135399190615dc2565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f390775370000000000000000000000000000000000000000000000000000000017905287866113886084613bc7565b925092509250826135dc57816040517fe1cd55090000000000000000000000000000000000000000000000000000000081526004016106ee9190614466565b81516020146136245781516040517f78ef80240000000000000000000000000000000000000000000000000000000081526020600482015260248101919091526044016106ee565b60008280602001905181019061363a91906159c3565b9050866001600160a01b03168c6001600160a01b0316146136cf57600061366b8d8a613666868a6151d3565b613a84565b5090508681108061368557508161368288836151d3565b14155b156136cd576040517fa966e21f0000000000000000000000000000000000000000000000000000000081526004810183905260248101889052604481018290526064016106ee565b505b604080518082019091526001600160a01b039098168852602088015250949550505050505095945050505050565b6000610e5d836001600160a01b038416613ced565b81516040517fd02641a00000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009182919084169063d02641a0906024016040805180830381865afa158015613778573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061379c9190615e8f565b5190507bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81166000036138055783516040517f9a655f7b0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911660048201526024016106ee565b6020840151611e33907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690613cf9565b61184f6003826000613d36565b600061386c827f01ffc9a7000000000000000000000000000000000000000000000000000000006138a4565b80156106a4575061389d827fffffffff000000000000000000000000000000000000000000000000000000006138a4565b1592915050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000082166024820152600090819060440160408051601f19818403018152919052602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825192935060009283928392909183918a617530fa92503d9150600051905082801561395d575060208210155b80156139695750600081115b979650505050505050565b60008183106139835781610e5d565b5090919050565b60006106a482614085565b6000808080612daa8686614090565b60008181526002830160205260408120819055610e5d83836140bb565b60008281526002840160205260408120829055611e3384846140c7565b60008151602014613a1d57816040517f8d666f600000000000000000000000000000000000000000000000000000000081526004016106ee9190614466565b600082806020019051810190613a3391906159c3565b90506001600160a01b03811180613a4b575061040081105b156106a457826040517f8d666f600000000000000000000000000000000000000000000000000000000081526004016106ee9190614466565b6000806000806000613b1388604051602401613aaf91906001600160a01b0391909116815260200190565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f70a082310000000000000000000000000000000000000000000000000000000017905288886113886084613bc7565b92509250925082613b5257816040517fe1cd55090000000000000000000000000000000000000000000000000000000081526004016106ee9190614466565b6020825114613b9a5781516040517f78ef80240000000000000000000000000000000000000000000000000000000081526020600482015260248101919091526044016106ee565b81806020019051810190613bae91906159c3565b613bb882886151d3565b94509450505050935093915050565b6000606060008361ffff1667ffffffffffffffff811115613bea57613bea614479565b6040519080825280601f01601f191660200182016040528015613c14576020820181803683370190505b509150863b613c47577f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b5a85811015613c7a577fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b8590036040810481038710613cb3577f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b505a6000808a5160208c0160008c8cf193505a900390503d84811115613cd65750835b808352806000602085013e50955095509592505050565b6000610e5d83836140d3565b6000670de0b6b3a7640000613d2c837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff861661523c565b610e5d9190615ef1565b825474010000000000000000000000000000000000000000900460ff161580613d5d575081155b15613d6757505050565b825460018401546fffffffffffffffffffffffffffffffff80831692911690600090613dad90700100000000000000000000000000000000900463ffffffff16426151d3565b90508015613e6d5781831115613def576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001860154613e299083908590849070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16612e94565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000004263ffffffff160217875592505b84821015613f0a576001600160a01b038416613ebf576040517ff94ebcd100000000000000000000000000000000000000000000000000000000815260048101839052602481018690526044016106ee565b6040517f1a76572a00000000000000000000000000000000000000000000000000000000815260048101839052602481018690526001600160a01b03851660448201526064016106ee565b848310156140035760018681015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16906000908290613f4e90826151d3565b613f58878a6151d3565b613f6291906154aa565b613f6c9190615ef1565b90506001600160a01b038616613fb8576040517f15279c0800000000000000000000000000000000000000000000000000000000815260048101829052602481018690526044016106ee565b6040517fd0c8d23a00000000000000000000000000000000000000000000000000000000815260048101829052602481018690526001600160a01b03871660448201526064016106ee565b61400d85846151d3565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b60006106a4826140df565b6000808061409e85856140e9565b600081815260029690960160205260409095205494959350505050565b6000610e5d83836140f5565b6000610e5d83836141ef565b6000610e5d838361423e565b60006106a4825490565b6000610e5d8383614256565b600081815260018301602052604081205480156141de5760006141196001836151d3565b855490915060009061412d906001906151d3565b905081811461419257600086600001828154811061414d5761414d615294565b906000526020600020015490508087600001848154811061417057614170615294565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806141a3576141a3615f05565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506106a4565b60009150506106a4565b5092915050565b6000818152600183016020526040812054614236575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556106a4565b5060006106a4565b60008181526001830160205260408120541515610e5d565b600082600001828154811061426d5761426d615294565b9060005260206000200154905092915050565b8280548282559060005260206000209081019282156142ed579160200282015b828111156142ed57825182547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039091161782556020909201916001909101906142a0565b506142f99291506142fd565b5090565b5b808211156142f957600081556001016142fe565b60e081016106a482846001600160a01b03808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015250508060608301511660608401528060808301511660808401528060a08301511660a08401528060c08301511660c0840152505050565b67ffffffffffffffff8116811461184f57600080fd5b80356143a381614382565b919050565b6000602082840312156143ba57600080fd5b8135610e5d81614382565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60048110614404576144046143c5565b9052565b602081016106a482846143f4565b60005b83811015614431578181015183820152602001614419565b50506000910152565b60008151808452614452816020860160208601614416565b601f01601f19169290920160200192915050565b602081526000610e5d602083018461443a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156144cb576144cb614479565b60405290565b6040516080810167ffffffffffffffff811182821017156144cb576144cb614479565b6040516101a0810167ffffffffffffffff811182821017156144cb576144cb614479565b604051601f8201601f1916810167ffffffffffffffff8111828210171561454157614541614479565b604052919050565b600067ffffffffffffffff82111561456357614563614479565b5060051b60200190565b6001600160a01b038116811461184f57600080fd5b80356143a38161456d565b600082601f83011261459e57600080fd5b813560206145b36145ae83614549565b614518565b8083825260208201915060208460051b8701019350868411156145d557600080fd5b602086015b848110156145fa5780356145ed8161456d565b83529183019183016145da565b509695505050505050565b803560ff811681146143a357600080fd5b600067ffffffffffffffff82111561463057614630614479565b50601f01601f191660200190565b600082601f83011261464f57600080fd5b813561465d6145ae82614616565b81815284602083860101111561467257600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c087890312156146a857600080fd5b863567ffffffffffffffff808211156146c057600080fd5b6146cc8a838b0161458d565b975060208901359150808211156146e257600080fd5b6146ee8a838b0161458d565b96506146fc60408a01614605565b9550606089013591508082111561471257600080fd5b61471e8a838b0161463e565b945061472c60808a01614398565b935060a089013591508082111561474257600080fd5b5061474f89828a0161463e565b9150509295509295509295565b60008151808452602080850194506020840160005b838110156147965781516001600160a01b031687529582019590820190600101614771565b509495945050505050565b602081526000610e5d602083018461475c565b6000602082840312156147c657600080fd5b8135610e5d8161456d565b60a081016106a4828463ffffffff8082511683528060208301511660208401525061ffff604082015116604083015260608101516001600160a01b03808216606085015280608084015116608085015250505050565b60006020828403121561483957600080fd5b813567ffffffffffffffff81111561485057600080fd5b820160a08185031215610e5d57600080fd5b600082601f83011261487357600080fd5b813560206148836145ae83614549565b82815260069290921b840181019181810190868411156148a257600080fd5b8286015b848110156145fa57604081890312156148bf5760008081fd5b6148c76144a8565b81356148d28161456d565b8152818501356148e18161456d565b818601528352918301916040016148a6565b6000806040838503121561490657600080fd5b823567ffffffffffffffff8082111561491e57600080fd5b61492a86838701614862565b9350602085013591508082111561494057600080fd5b5061494d85828601614862565b9150509250929050565b801515811461184f57600080fd5b80356143a381614957565b60006040828403121561498257600080fd5b61498a6144a8565b905081356149978161456d565b808252506020820135602082015292915050565b600082601f8301126149bc57600080fd5b813560206149cc6145ae83614549565b8083825260208201915060208460061b8701019350868411156149ee57600080fd5b602086015b848110156145fa57614a058882614970565b8352918301916040016149f3565b600082601f830112614a2457600080fd5b81356020614a346145ae83614549565b82815260059290921b84018101918181019086841115614a5357600080fd5b8286015b848110156145fa57803567ffffffffffffffff811115614a775760008081fd5b614a858986838b010161463e565b845250918301918301614a57565b600082601f830112614aa457600080fd5b81356020614ab46145ae83614549565b82815260059290921b84018101918181019086841115614ad357600080fd5b8286015b848110156145fa57803567ffffffffffffffff811115614af75760008081fd5b614b058986838b0101614a13565b845250918301918301614ad7565b600082601f830112614b2457600080fd5b81356020614b346145ae83614549565b8083825260208201915060208460051b870101935086841115614b5657600080fd5b602086015b848110156145fa5780358352918301918301614b5b565b600060808284031215614b8457600080fd5b614b8c6144d1565b9050813567ffffffffffffffff80821115614ba657600080fd5b818401915084601f830112614bba57600080fd5b81356020614bca6145ae83614549565b82815260059290921b84018101918181019088841115614be957600080fd5b8286015b84811015614d4557803586811115614c0457600080fd5b87016101a0818c03601f19011215614c1b57600080fd5b614c236144f4565b614c2e868301614398565b8152614c3c60408301614582565b86820152614c4c60608301614582565b6040820152614c5d60808301614398565b606082015260a08201356080820152614c7860c08301614965565b60a0820152614c8960e08301614398565b60c0820152610100614c9c818401614582565b60e083015261012080840135828401526101409150818401358a811115614cc257600080fd5b614cd08f8a8388010161463e565b828501525050610160808401358a811115614cea57600080fd5b614cf88f8a838801016149ab565b83850152506101809150818401358a811115614d1357600080fd5b614d218f8a83880101614a13565b91840191909152506101a09290920135918101919091528352918301918301614bed565b5086525085810135935082841115614d5c57600080fd5b614d6887858801614a93565b90850152506040840135915080821115614d8157600080fd5b50614d8e84828501614b13565b6040830152506060820135606082015292915050565b63ffffffff8116811461184f57600080fd5b600082601f830112614dc757600080fd5b81356020614dd76145ae83614549565b8083825260208201915060208460051b870101935086841115614df957600080fd5b602086015b848110156145fa578035614e1181614da4565b8352918301918301614dfe565b6000806040808486031215614e3257600080fd5b833567ffffffffffffffff80821115614e4a57600080fd5b614e5687838801614b72565b9450602091508186013581811115614e6d57600080fd5b8601601f81018813614e7e57600080fd5b8035614e8c6145ae82614549565b81815260059190911b8201840190848101908a831115614eab57600080fd5b8584015b83811015614f1e57803586811115614ec75760008081fd5b8501808d03601f1901891315614edd5760008081fd5b614ee56144a8565b8882013581528982013588811115614efd5760008081fd5b614f0b8f8b83860101614db6565b828b015250845250918601918601614eaf565b50809750505050505050509250929050565b60008083601f840112614f4257600080fd5b50813567ffffffffffffffff811115614f5a57600080fd5b6020830191508360208260051b8501011115612db257600080fd5b60008060008060008060008060e0898b031215614f9157600080fd5b606089018a811115614fa257600080fd5b8998503567ffffffffffffffff80821115614fbc57600080fd5b818b0191508b601f830112614fd057600080fd5b813581811115614fdf57600080fd5b8c6020828501011115614ff157600080fd5b6020830199508098505060808b013591508082111561500f57600080fd5b61501b8c838d01614f30565b909750955060a08b013591508082111561503457600080fd5b506150418b828c01614f30565b999c989b50969995989497949560c00135949350505050565b6000806000806060858703121561507057600080fd5b843567ffffffffffffffff8082111561508857600080fd5b908601906101a0828903121561509d57600080fd5b909450602086013590808211156150b357600080fd5b6150bf88838901614f30565b909550935060408701359150808211156150d857600080fd5b506150e587828801614db6565b91505092959194509250565b80356fffffffffffffffffffffffffffffffff811681146143a357600080fd5b60006060828403121561512357600080fd5b6040516060810181811067ffffffffffffffff8211171561514657615146614479565b604052823561515481614957565b8152615162602084016150f1565b6020820152615173604084016150f1565b60408201529392505050565b604081526000615192604083018561475c565b8281036020840152612eb3818561475c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156106a4576106a46151a4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff80841680615230576152306151e6565b92169190910692915050565b80820281158282048414176106a4576106a46151a4565b600067ffffffffffffffff8084168061526e5761526e6151e6565b92169190910492915050565b602081016003831061528e5761528e6143c5565b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600063ffffffff8083168181036152dc576152dc6151a4565b6001019392505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526153168184018a61475c565b9050828103608084015261532a818961475c565b905060ff871660a084015282810360c0840152615347818761443a565b905067ffffffffffffffff851660e084015282810361010084015261536c818561443a565b9c9b505050505050505050505050565b60006020828403121561538e57600080fd5b8151610e5d81614382565b600082601f8301126153aa57600080fd5b81516153b86145ae82614616565b8181528460208386010111156153cd57600080fd5b611e33826020830160208701614416565b6000602082840312156153f057600080fd5b815167ffffffffffffffff8082111561540857600080fd5b908301906080828603121561541c57600080fd5b6154246144d1565b82518281111561543357600080fd5b61543f87828601615399565b82525060208301518281111561545457600080fd5b61546087828601615399565b60208301525060408301518281111561547857600080fd5b61548487828601615399565b6040830152506060830151925061549a83614da4565b6060810192909252509392505050565b808201808211156106a4576106a46151a4565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126154f257600080fd5b83018035915067ffffffffffffffff82111561550d57600080fd5b6020019150600681901b3603821315612db257600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261555a57600080fd5b83018035915067ffffffffffffffff82111561557557600080fd5b6020019150600581901b3603821315612db257600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126155c257600080fd5b83018035915067ffffffffffffffff8211156155dd57600080fd5b602001915036819003821315612db257600080fd5b60008151808452602080850194506020840160005b8381101561479657815180516001600160a01b031688528301518388015260409096019590820190600101615607565b608081528451608082015267ffffffffffffffff60208601511660a08201526000604086015160a060c084015261567261012084018261443a565b905060608701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80808584030160e08601526156ae838361443a565b9250608089015191508085840301610100860152506156cd82826155f2565b925050506156e1602083018661ffff169052565b836040830152612eb360608301846001600160a01b03169052565b60008060006060848603121561571157600080fd5b835161571c81614957565b602085015190935067ffffffffffffffff81111561573957600080fd5b61574586828701615399565b925050604084015190509250925092565b600060a0828403121561576857600080fd5b60405160a0810181811067ffffffffffffffff8211171561578b5761578b614479565b604052825161579981614da4565b815260208301516157a981614da4565b6020820152604083015161ffff811681146157c357600080fd5b604082015260608301516157d68161456d565b606082015260808301516157e98161456d565b60808201529392505050565b610180810161586682856001600160a01b03808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015250508060608301511660608401528060808301511660808401528060a08301511660a08401528060c08301511660c0840152505050565b825163ffffffff90811660e0840152602084015116610100830152604083015161ffff1661012083015260608301516001600160a01b03908116610140840152608084015116610160830152610e5d565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526158f18285018b61475c565b91508382036080850152615905828a61475c565b915060ff881660a085015283820360c0850152615922828861443a565b90861660e0850152838103610100850152905061536c818561443a565b60006020828403121561595157600080fd5b8151610e5d81614957565b60008151808452602080850194506020840160005b8381101561479657815187529582019590820190600101615971565b6060815260006159a0606083018661595c565b82810360208401526159b2818661595c565b915050826040830152949350505050565b6000602082840312156159d557600080fd5b5051919050565b67ffffffffffffffff8181168382160190808211156141e8576141e86151a4565b67ffffffffffffffff8316815260408101610e5d60208301846143f4565b600067ffffffffffffffff8083168181036152dc576152dc6151a4565b615a4281846143f4565b604060208201526000611e33604083018461443a565b600060208284031215615a6a57600080fd5b813567ffffffffffffffff811115615a8157600080fd5b611e3384828501614b72565b600060408284031215615a9f57600080fd5b610e5d8383614970565b600060208284031215615abb57600080fd5b813567ffffffffffffffff80821115615ad357600080fd5b9083019060808286031215615ae757600080fd5b615aef6144d1565b823582811115615afe57600080fd5b615b0a8782860161463e565b825250602083013582811115615b1f57600080fd5b615b2b8782860161463e565b602083015250604083013582811115615b4357600080fd5b615b4f8782860161463e565b6040830152506060830135925061549a83614da4565b602081526000610e5d60208301846155f2565b60008282518085526020808601955060208260051b8401016020860160005b84811015615bc557601f19868403018952615bb383835161443a565b98840198925090830190600101615b97565b5090979650505050505050565b602081526000610e5d6020830184615b78565b60008151808452602080850194506020840160005b8381101561479657815163ffffffff1687529582019590820190600101615bfa565b60608152615c3760608201855167ffffffffffffffff169052565b60006020850151615c5360808401826001600160a01b03169052565b5060408501516001600160a01b03811660a084015250606085015167ffffffffffffffff811660c084015250608085015160e083015260a0850151610100615c9e8185018315159052565b60c08701519150610120615cbd8186018467ffffffffffffffff169052565b60e08801519250610140615cdb818701856001600160a01b03169052565b828901519350610160925083838701528189015193506101a091506101808281880152615d0c61020088018661443a565b9450818a015191507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0808887030184890152615d4886846155f2565b948b01518886039091016101c0890152939450615d658585615b78565b9450808a01516101e0880152505050508281036020840152615d878186615b78565b90508281036040840152615d9b8185615be5565b9695505050505050565b600060208284031215615db757600080fd5b8151610e5d8161456d565b6020815260008251610100806020850152615de161012085018361443a565b91506020850151615dfe604086018267ffffffffffffffff169052565b5060408501516001600160a01b038116606086015250606085015160808501526080850151615e3860a08601826001600160a01b03169052565b5060a0850151601f19808685030160c0870152615e55848361443a565b935060c08701519150808685030160e0870152615e72848361443a565b935060e0870151915080868503018387015250615d9b838261443a565b600060408284031215615ea157600080fd5b615ea96144a8565b82517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81168114615ed557600080fd5b81526020830151615ee581614da4565b60208201529392505050565b600082615f0057615f006151e6565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000818000a",
}

var EVM2EVMOffRampABI = EVM2EVMOffRampMetaData.ABI

var EVM2EVMOffRampBin = EVM2EVMOffRampMetaData.Bin

func DeployEVM2EVMOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig EVM2EVMOffRampStaticConfig, rateLimiterConfig RateLimiterConfig) (common.Address, *CustomTransaction, *EVM2EVMOffRamp, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	switch chainId.Uint64() {

	case 324, 280, 300:
		return DeployZkSyncEVM2EVMOffRamp(auth, backend, staticConfig, rateLimiterConfig)
	}

	parsed, err := EVM2EVMOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMOffRampBin), backend, staticConfig, rateLimiterConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, &CustomTransaction{Transaction: tx, CustomHash: tx.Hash()}, &EVM2EVMOffRamp{address: address, abi: *parsed, EVM2EVMOffRampCaller: EVM2EVMOffRampCaller{contract: contract}, EVM2EVMOffRampTransactor: EVM2EVMOffRampTransactor{contract: contract}, EVM2EVMOffRampFilterer: EVM2EVMOffRampFilterer{contract: contract}}, nil
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampCaller) GetAllRateLimitTokens(opts *bind.CallOpts) (GetAllRateLimitTokens,

	error) {
	var out []interface{}
	err := _EVM2EVMOffRamp.contract.Call(opts, &out, "getAllRateLimitTokens")

	outstruct := new(GetAllRateLimitTokens)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceTokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.DestTokens = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) GetAllRateLimitTokens() (GetAllRateLimitTokens,

	error) {
	return _EVM2EVMOffRamp.Contract.GetAllRateLimitTokens(&_EVM2EVMOffRamp.CallOpts)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampCallerSession) GetAllRateLimitTokens() (GetAllRateLimitTokens,

	error) {
	return _EVM2EVMOffRamp.Contract.GetAllRateLimitTokens(&_EVM2EVMOffRamp.CallOpts)
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte, tokenGasOverrides []uint32) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "executeSingleMessage", message, offchainTokenData, tokenGasOverrides)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, offchainTokenData [][]byte, tokenGasOverrides []uint32) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMOffRamp.TransactOpts, message, offchainTokenData, tokenGasOverrides)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) ExecuteSingleMessage(message InternalEVM2EVMMessage, offchainTokenData [][]byte, tokenGasOverrides []uint32) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.ExecuteSingleMessage(&_EVM2EVMOffRamp.TransactOpts, message, offchainTokenData, tokenGasOverrides)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) ManuallyExecute(opts *bind.TransactOpts, report InternalExecutionReport, gasLimitOverrides []EVM2EVMOffRampGasLimitOverride) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "manuallyExecute", report, gasLimitOverrides)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) ManuallyExecute(report InternalExecutionReport, gasLimitOverrides []EVM2EVMOffRampGasLimitOverride) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.ManuallyExecute(&_EVM2EVMOffRamp.TransactOpts, report, gasLimitOverrides)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) ManuallyExecute(report InternalExecutionReport, gasLimitOverrides []EVM2EVMOffRampGasLimitOverride) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.ManuallyExecute(&_EVM2EVMOffRamp.TransactOpts, report, gasLimitOverrides)
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactor) UpdateRateLimitTokens(opts *bind.TransactOpts, removes []EVM2EVMOffRampRateLimitToken, adds []EVM2EVMOffRampRateLimitToken) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.contract.Transact(opts, "updateRateLimitTokens", removes, adds)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampSession) UpdateRateLimitTokens(removes []EVM2EVMOffRampRateLimitToken, adds []EVM2EVMOffRampRateLimitToken) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.UpdateRateLimitTokens(&_EVM2EVMOffRamp.TransactOpts, removes, adds)
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampTransactorSession) UpdateRateLimitTokens(removes []EVM2EVMOffRampRateLimitToken, adds []EVM2EVMOffRampRateLimitToken) (*types.Transaction, error) {
	return _EVM2EVMOffRamp.Contract.UpdateRateLimitTokens(&_EVM2EVMOffRamp.TransactOpts, removes, adds)
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

type EVM2EVMOffRampAlreadyAttemptedIterator struct {
	Event *EVM2EVMOffRampAlreadyAttempted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampAlreadyAttemptedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampAlreadyAttempted)
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
		it.Event = new(EVM2EVMOffRampAlreadyAttempted)
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

func (it *EVM2EVMOffRampAlreadyAttemptedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampAlreadyAttemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampAlreadyAttempted struct {
	SequenceNumber uint64
	Raw            types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterAlreadyAttempted(opts *bind.FilterOpts) (*EVM2EVMOffRampAlreadyAttemptedIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "AlreadyAttempted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampAlreadyAttemptedIterator{contract: _EVM2EVMOffRamp.contract, event: "AlreadyAttempted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchAlreadyAttempted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampAlreadyAttempted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "AlreadyAttempted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampAlreadyAttempted)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "AlreadyAttempted", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseAlreadyAttempted(log types.Log) (*EVM2EVMOffRampAlreadyAttempted, error) {
	event := new(EVM2EVMOffRampAlreadyAttempted)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "AlreadyAttempted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampConfigChangedIterator struct {
	Event *EVM2EVMOffRampConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampConfigChanged)
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
		it.Event = new(EVM2EVMOffRampConfigChanged)
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

func (it *EVM2EVMOffRampConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampConfigChangedIterator{contract: _EVM2EVMOffRamp.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampConfigChanged)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMOffRampConfigChanged, error) {
	event := new(EVM2EVMOffRampConfigChanged)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

type EVM2EVMOffRampSkippedAlreadyExecutedMessageIterator struct {
	Event *EVM2EVMOffRampSkippedAlreadyExecutedMessage

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampSkippedAlreadyExecutedMessageIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampSkippedAlreadyExecutedMessage)
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
		it.Event = new(EVM2EVMOffRampSkippedAlreadyExecutedMessage)
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

func (it *EVM2EVMOffRampSkippedAlreadyExecutedMessageIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampSkippedAlreadyExecutedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampSkippedAlreadyExecutedMessage struct {
	SequenceNumber uint64
	Raw            types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterSkippedAlreadyExecutedMessage(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMOffRampSkippedAlreadyExecutedMessageIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "SkippedAlreadyExecutedMessage", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampSkippedAlreadyExecutedMessageIterator{contract: _EVM2EVMOffRamp.contract, event: "SkippedAlreadyExecutedMessage", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchSkippedAlreadyExecutedMessage(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampSkippedAlreadyExecutedMessage, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "SkippedAlreadyExecutedMessage", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampSkippedAlreadyExecutedMessage)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "SkippedAlreadyExecutedMessage", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseSkippedAlreadyExecutedMessage(log types.Log) (*EVM2EVMOffRampSkippedAlreadyExecutedMessage, error) {
	event := new(EVM2EVMOffRampSkippedAlreadyExecutedMessage)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "SkippedAlreadyExecutedMessage", log); err != nil {
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

type EVM2EVMOffRampTokenAggregateRateLimitAddedIterator struct {
	Event *EVM2EVMOffRampTokenAggregateRateLimitAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampTokenAggregateRateLimitAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampTokenAggregateRateLimitAdded)
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
		it.Event = new(EVM2EVMOffRampTokenAggregateRateLimitAdded)
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

func (it *EVM2EVMOffRampTokenAggregateRateLimitAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampTokenAggregateRateLimitAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampTokenAggregateRateLimitAdded struct {
	SourceToken common.Address
	DestToken   common.Address
	Raw         types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterTokenAggregateRateLimitAdded(opts *bind.FilterOpts) (*EVM2EVMOffRampTokenAggregateRateLimitAddedIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "TokenAggregateRateLimitAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampTokenAggregateRateLimitAddedIterator{contract: _EVM2EVMOffRamp.contract, event: "TokenAggregateRateLimitAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchTokenAggregateRateLimitAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTokenAggregateRateLimitAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "TokenAggregateRateLimitAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampTokenAggregateRateLimitAdded)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "TokenAggregateRateLimitAdded", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseTokenAggregateRateLimitAdded(log types.Log) (*EVM2EVMOffRampTokenAggregateRateLimitAdded, error) {
	event := new(EVM2EVMOffRampTokenAggregateRateLimitAdded)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "TokenAggregateRateLimitAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampTokenAggregateRateLimitRemovedIterator struct {
	Event *EVM2EVMOffRampTokenAggregateRateLimitRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampTokenAggregateRateLimitRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampTokenAggregateRateLimitRemoved)
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
		it.Event = new(EVM2EVMOffRampTokenAggregateRateLimitRemoved)
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

func (it *EVM2EVMOffRampTokenAggregateRateLimitRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampTokenAggregateRateLimitRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampTokenAggregateRateLimitRemoved struct {
	SourceToken common.Address
	DestToken   common.Address
	Raw         types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterTokenAggregateRateLimitRemoved(opts *bind.FilterOpts) (*EVM2EVMOffRampTokenAggregateRateLimitRemovedIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "TokenAggregateRateLimitRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampTokenAggregateRateLimitRemovedIterator{contract: _EVM2EVMOffRamp.contract, event: "TokenAggregateRateLimitRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchTokenAggregateRateLimitRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTokenAggregateRateLimitRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "TokenAggregateRateLimitRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampTokenAggregateRateLimitRemoved)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "TokenAggregateRateLimitRemoved", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseTokenAggregateRateLimitRemoved(log types.Log) (*EVM2EVMOffRampTokenAggregateRateLimitRemoved, error) {
	event := new(EVM2EVMOffRampTokenAggregateRateLimitRemoved)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "TokenAggregateRateLimitRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMOffRampTokensConsumedIterator struct {
	Event *EVM2EVMOffRampTokensConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMOffRampTokensConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMOffRampTokensConsumed)
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
		it.Event = new(EVM2EVMOffRampTokensConsumed)
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

func (it *EVM2EVMOffRampTokensConsumedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMOffRampTokensConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMOffRampTokensConsumed struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) FilterTokensConsumed(opts *bind.FilterOpts) (*EVM2EVMOffRampTokensConsumedIterator, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.FilterLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMOffRampTokensConsumedIterator{contract: _EVM2EVMOffRamp.contract, event: "TokensConsumed", logs: logs, sub: sub}, nil
}

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTokensConsumed) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMOffRamp.contract.WatchLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMOffRampTokensConsumed)
				if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
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

func (_EVM2EVMOffRamp *EVM2EVMOffRampFilterer) ParseTokensConsumed(log types.Log) (*EVM2EVMOffRampTokensConsumed, error) {
	event := new(EVM2EVMOffRampTokensConsumed)
	if err := _EVM2EVMOffRamp.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
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

var EVM2EVMOffRampZkBin string = ("0x0004000000000002003400000000000200000060031002700000092b0030019d0000092b03300197000300000031035500020000000103550000000100200190000000d70000c13d0000008002000039000000400020043f000000040030008c000000f80000413d000000000201043b000000e002200270000009430020009c000001030000a13d000009440020009c000001a10000a13d000009450020009c000001cf0000213d000009490020009c000003d30000613d0000094a0020009c000005590000613d0000094b0020009c000000f80000c13d000000640030008c000000f80000413d0000000002000416000000000002004b000000f80000c13d0000000402100370000000000202043b002000000002001d000009300020009c000000f80000213d000000200230006a000009680020009c000000f80000213d000001a40020008c000000f80000413d0000002402100370000000000202043b001f00000002001d000009300020009c000000f80000213d0000001f020000290000002302200039000000000032004b000000f80000813d0000001f020000290000000402200039000000000221034f000000000202043b001e00000002001d000009300020009c000000f80000213d0000001f02000029001d00240020003d0000001e0200002900000005022002100000001d02200029000000000032004b000000f80000213d0000004402100370000000000202043b000009300020009c000000f80000213d0000002304200039000000000034004b000000f80000813d0000000404200039000000000441034f000000000504043b000009300050009c000000fd0000213d00000005045002100000003f064000390000095f06600197000009600060009c000000fd0000213d0000008006600039000000400060043f000000800050043f00000024022000390000000004240019000000000034004b000000f80000213d000000000005004b000000650000613d0000008003000039000000000521034f000000000505043b0000092b0050009c000000f80000213d000000200330003900000000005304350000002002200039000000000042004b0000005c0000413d0000000001000415001000000001001d000000400100043d001a00000001001d00000000010004100000000002000411000000000012004b00000aa30000c13d0000001a010000290000096a0010009c000000fd0000213d000000200200002900000004052000390000001a030000290000002001300039000000400010043f0000000000030435000201440020003d00000002010003670000000202100360000000000202043b0000000004000031000300000005001d00000000035400490000001f0730008a0000096b087001970000096b03200197000000000583013f000000000083004b00000000030000190000096b03004041000000000072004b00000000060000190000096b060080410000096b0050009c000000000306c019000000000003004b000000f80000c13d0000000302200029000000000321034f000000000303043b000900000003001d000009300030009c000000f80000213d00000009030000290000000606300210000000000564004900000020032000390000096b095001970000096b0a300197000000000b9a013f00000000009a004b00000000090000190000096b09004041000000000053004b00000000050000190000096b050020410000096b00b0009c000000000905c019000000000009004b000000f80000c13d000000090000006b00000c5d0000c13d0000000202000029001f0020002000920000001f02100360000000000202043b0000096b032001970000096b05700197000000000653013f000000000053004b00000000030000190000096b03004041000000000072004b00000000050000190000096b050080410000096b0060009c000000000305c019000000000003004b000000f80000c13d0000000303200029000000000231034f000000000202043b000009300020009c000000f80000213d00000000042400490000002003300039000000000043004b00000000050000190000096b050020410000096b044001970000096b03300197000000000643013f000000000043004b00000000030000190000096b030040410000096b0060009c000000000305c019000000000003004b000000f80000c13d000000000002004b0000147a0000c13d0000001f02000029000000a00220008a000000000221034f0000000003000415000000230330008a001e000500300218000000000202043b000000000002004b002300000000003d002300010000603d0000147e0000c13d000014970000013d000001a004000039000000400040043f0000000002000416000000000002004b000000f80000c13d0000001f023000390000092c02200197000001a002200039000000400020043f0000001f0530018f0000092d06300198000001a002600039000000e90000613d000000000701034f000000007807043c0000000004840436000000000024004b000000e50000c13d000000000005004b000000f60000613d000000000161034f0000000304500210000000000502043300000000054501cf000000000545022f000000000101043b0000010004400089000000000141022f00000000014101cf000000000151019f0000000000120435000001400030008c000000fa0000813d0000000001000019000024aa00010430000000400300043d0000092e0030009c000001450000a13d000009a401000041000000000010043f0000004101000039000000040010043f0000097201000041000024aa00010430000009520020009c000001ae0000213d000009590020009c000002240000a13d0000095a0020009c000003a80000613d0000095b0020009c000004260000613d0000095c0020009c000000f80000c13d0000000001000416000000000001004b000000f80000c13d000000800000043f000000a00000043f000000c00000043f000000e00000043f000001000000043f0000000301000039000000000101041a0000093202100197001f00000002001d000001200020043f00000080021002700000092b02200197002000000002001d000001400020043f00000987001001980000000001000039000000010100c039000001600010043f0000000401000039000000000101041a0000093202100197001e00000002001d000001800020043f0000008001100270000001a00010043f0000026001000039000000400010043f000001c00000043f000001e00000043f000002000000043f000002200000043f000002400000043f0000093501000041000000000010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000936011001c70000800b0200003924a824a30000040f000000010020019000001c500000613d000000000101043b0000002005000029000000000351004b000007940000813d000009a401000041000000000010043f0000001101000039000000040010043f0000097201000041000024aa00010430000000e001300039000000400010043f000001a00100043d0000092f0010009c000000f80000213d0000000001130436002000000001001d000001c00100043d000009300010009c000000f80000213d00000020020000290000000000120435000001e00100043d000009300010009c000000f80000213d0000004002300039001f00000002001d0000000000120435000002000100043d0000092f0010009c000000f80000213d0000006002300039001e00000002001d0000000000120435000002200100043d0000092f0010009c000000f80000213d0000008002300039001d00000002001d0000000000120435000002400100043d0000092f0010009c000000f80000213d000000a002300039001c00000002001d0000000000120435000002600100043d0000092f0010009c000000f80000213d000000c002300039001b00000002001d0000000000120435000000400100043d000009310010009c000000fd0000213d0000006002100039000000400020043f000002800400043d000000000004004b0000000002000039000000010200c039001a00000004001d000000000024004b000000f80000c13d0000001a020000290000000002210436000002a00400043d001900000004001d000009320040009c000000f80000213d00000019040000290000000000420435000002c00200043d001800000002001d000009320020009c000000f80000213d001600000003001d000000400110003900000018020000290000000000210435000000400100043d001700000001001d0000000001000411000000000001004b00000a4c0000c13d0000001703000029000000440130003900000940020000410000000000210435000000240130003900000018020000390000000000210435000009410100004100000000001304350000000401300039000000200200003900000000002104350000092b0030009c0000092b03008041000000400130021000000942011001c7000024aa000104300000094c0020009c000001ef0000a13d0000094d0020009c000002800000613d0000094e0020009c000003440000613d0000094f0020009c000000f80000c13d0000000001000416000000000001004b000000f80000c13d000000000100041a000006f60000013d000009530020009c000002540000a13d000009540020009c000003bc0000613d000009550020009c000005260000613d000009560020009c000000f80000c13d0000000001000416000000000001004b000000f80000c13d0000000101000039000000000201041a0000092f032001970000000006000411000000000036004b0000077d0000c13d000000000300041a0000093304300197000000000464019f000000000040041b0000093302200197000000000021041b00000000010004140000092f053001970000092b0010009c0000092b01008041000000c001100210000009ab011001c70000800d020000390000000303000039000009af04000041000007db0000013d000009460020009c000003dc0000613d000009470020009c000005f80000613d000009480020009c000000f80000c13d000000240030008c000000f80000413d0000000002000416000000000002004b000000f80000c13d0000000401100370000000000601043b0000092f0060009c000000f80000213d000000000100041a0000092f011001970000000005000411000000000015004b000007bd0000c13d000000000056004b000007ce0000c13d0000094101000041000000800010043f0000002001000039000000840010043f0000001701000039000000a40010043f000009ad01000041000000c40010043f000009aa01000041000024aa00010430000009500020009c000004170000613d000009510020009c000000f80000c13d000000240030008c000000f80000413d0000000002000416000000000002004b000000f80000c13d0000000401100370000000000101043b002000000001001d0000092f0010009c000000f80000213d0000002001000029000000000010043f0000000f01000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000000101043b000000000101041a0000093002100198000007c70000c13d0000097001000041000000000010044300000000010004120000000400100443000000c001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f000000010020019000001c500000613d000000400500043d000000000101043b0000092f02100198000008280000c13d00000000010500190000000002000019000007c80000013d0000095d0020009c000006590000613d0000095e0020009c000000f80000c13d000000240030008c000000f80000413d0000000002000416000000000002004b000000f80000c13d0000000401100370000000000101043b002000000001001d000009300010009c000000f80000213d00000020010000290000000701100270000000000010043f0000001001000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d00000020030000290000000102300210000000000101043b000000000101041a0000007f033001900000024a0000613d000000ff0420018f00000000033400d9000000020030008c0000013f0000c13d000000fe0220018f000000000121022f000000030110018f000000400200043d00000000001204350000092b0020009c0000092b020080410000004001200210000009a8011001c7000024a90001042e000009570020009c000006f10000613d000009580020009c000000f80000c13d0000000001000416000000000001004b000000f80000c13d0000000902000039000000000102041a000000800010043f000000000020043f0000002002000039000000000001004b0000026e0000613d000000a004000039000009b40200004100000000030000190000000005040019000000000402041a0000092f04400197000000000445043600000001022000390000000103300039000000000013004b000002650000413d000000600250008a000000800100003924a81d8f0000040f000000400200043d002000000002001d00000020010000390000000002120436000000800100003924a81dbf0000040f000000200200002900000000012100490000092b0010009c0000092b0100804100000060011002100000092b0020009c0000092b020080410000004002200210000000000121019f000024a90001042e000000440030008c000000f80000413d0000000002000416000000000002004b000000f80000c13d0000000402100370000000000202043b000009300020009c000000f80000213d0000002304200039000000000034004b000000f80000813d0000000404200039000000000441034f000000000504043b000009300050009c000000fd0000213d00000005045002100000003f044000390000095f04400197000009600040009c000000fd0000213d0000008004400039000000400040043f000000800050043f000000240220003900000006045002100000000004240019000000000034004b000000f80000213d000000000005004b0000080c0000c13d0000002402100370000000000202043b000009300020009c000000f80000213d0000002304200039000000000034004b00000000050000190000096b050080410000096b04400197000000000004004b00000000060000190000096b060040410000096b0040009c000000000605c019000000000006004b000000f80000c13d0000000404200039000000000441034f000000000504043b000009300050009c000000fd0000213d00000005045002100000003f044000390000095f04400197000000400600043d0000000004460019001d00000006001d000000000064004b00000000060000390000000106004039000009300040009c000000fd0000213d0000000100600190000000fd0000c13d000000400040043f0000001d040000290000000004540436001b00000004001d000000240220003900000006045002100000000004240019000000000034004b000000f80000213d000000000005004b00000b2e0000c13d000000000100041a0000092f011001970000000002000411000000000012004b000017bb0000c13d000000800100043d000000000001004b00000b4a0000c13d0000001d010000290000000001010433000000000001004b000007de0000613d002000000000001d000002e10000013d0000002002000029002000010020003d0000001d010000290000000001010433000000200010006b000007de0000813d000000200100002900000005011002100000001b01100029001c00000001001d00000000010104330000000012010434001e00000002001d00000000010104330000092f01100197001f00000001001d000000000010043f0000000e01000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000000101043b0000001e020000290000092f02200197000000000021041b0000001f01000029000000000010043f0000000d01000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000000101043b000000000101041a000000000001004b000002db0000c13d0000000c03000039000000000103041a000009300010009c000000fd0000213d0000000102100039000000000023041b000009610110009a0000001f02000029000000000021041b000000000103041a001e00000001001d000000000020043f0000000d01000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000000101043b0000001e02000029000000000021041b0000001d010000290000000001010433000000200010006c00001cca0000a13d0000001c010000290000000001010433000000002101043400000000020204330000092f02200197000000400300043d000000200430003900000000002404350000092f0110019700000000001304350000092b0030009c0000092b03008041000000400130021000000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f00000962011001c70000800d020000390000000103000039000009a60400004124a8249e0000040f0000000100200190000002db0000c13d000000f80000013d000000440030008c000000f80000413d0000000002000416000000000002004b000000f80000c13d0000000402100370000000000202043b001700000002001d000009300020009c000000f80000213d0000001702000029001600040020003d000000160230006a000009680020009c000000f80000213d000000800020008c000000f80000413d0000010002000039000000400020043f0000001602100360000000000202043b000009300020009c000000f80000213d0000001602200029001a00000002001d0000001f02200039000000000032004b000000f80000813d0000001a02100360000000000202043b000009300020009c000000fd0000213d00000005042002100000003f054000390000095f055001970000097e0050009c000000fd0000213d0000010005500039000000400050043f000001000020043f0000001a050000290000002006500039001900000064001d000000190030006b000000f80000213d000000000002004b000008d70000c13d0000010002000039000000800020043f00000016020000290000002002200039000000000221034f000000000202043b000009300020009c000000f80000213d0000001602200029001c00000002001d0000001f02200039000000000032004b00000000040000190000096b040080410000096b02200197000000000002004b00000000050000190000096b050040410000096b0020009c000000000504c019000000000005004b000000f80000c13d0000001c02100360000000000202043b000009300020009c000000fd0000213d00000005042002100000003f054000390000095f05500197000000400600043d0000000005560019001a00000006001d000000000065004b00000000060000390000000106004039000009300050009c000000fd0000213d0000000100600190000000fd0000c13d000000400050043f0000001a0500002900000000002504350000001c020000290000002005200039001b00000054001d0000001b0030006b000000f80000213d00000000020500190000001b0050006c00000bcb0000813d000000200900008a001e001a0000002d000007030000013d0000000001000416000000000001004b000000f80000c13d000000c001000039000000400010043f0000001401000039000000800010043f000009c501000041000000a00010043f0000002001000039000000c00010043f0000008001000039000000e00200003924a81dad0000040f000000c00110008a0000092b0010009c0000092b010080410000006001100210000009c6011001c7000024a90001042e000000240030008c000000f80000413d0000000002000416000000000002004b000000f80000c13d0000000401100370000000000101043b0000092f0010009c000000f80000213d0000000202000039000000000302041a000000000400041a0000092f054001970000000004000411000000000054004b000007870000613d0000092f05300197000000000054004b000007870000613d0000096301000041000000800010043f000009b101000041000024aa000104300000000001000416000000000001004b000000f80000c13d0000000101000039000000800010043f000000a00000043f000000c00000043f0000099c01000041000024a90001042e000000640030008c000000f80000413d0000000002000416000000000002004b000000f80000c13d000000e002000039000000400020043f0000000402100370000000000202043b000000000002004b0000000003000039000000010300c039000000000032004b000000f80000c13d000000800020043f0000002402100370000000000202043b000009320020009c000000f80000213d000000a00020043f0000004401100370000000000101043b000009320010009c000000f80000213d000000c00010043f000000000100041a0000092f021001970000000001000411000000000021004b000003ff0000613d0000000202000039000000000202041a0000092f02200197000000000021004b000008350000c13d0000000301000039000000000101041a002000000001001d0000093501000041000000000010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000936011001c70000800b0200003924a824a30000040f000000010020019000001c500000613d000000200600002900000080026002700000092b02200197000000000101043b000000000421004b0000013f0000413d000008770000c13d0000000304000039000000000104041a0000088c0000013d0000000001000416000000000001004b000000f80000c13d0000000501000039000000000101041a0000000702000039000000000202041a0000092b03200197000000800030043f00000020022002700000092b02200197000000a00020043f000000c00010043f0000099c01000041000024a90001042e000000c40030008c000000f80000413d0000000002000416000000000002004b000000f80000c13d0000000402100370000000000202043b000009300020009c000000f80000213d0000002304200039000000000034004b000000f80000813d0000000404200039000000000441034f000000000504043b000009300050009c000000fd0000213d00000005045002100000003f064000390000095f06600197000009600060009c000000fd0000213d0000008006600039000000400060043f000000800050043f00000024022000390000000004240019000000000034004b000000f80000213d000000000005004b0000044f0000613d0000008005000039000000000621034f000000000606043b0000092f0060009c000000f80000213d000000200550003900000000006504350000002002200039000000000042004b000004460000413d0000002402100370000000000202043b000009300020009c000000f80000213d0000002304200039000000000034004b00000000050000190000096b050080410000096b04400197000000000004004b00000000060000190000096b060040410000096b0040009c000000000605c019000000000006004b000000f80000c13d0000000404200039000000000441034f000000000404043b000009300040009c000000fd0000213d00000005054002100000003f065000390000095f06600197000000400700043d0000000006670019002000000007001d000000000076004b00000000070000390000000107004039000009300060009c000000fd0000213d0000000100700190000000fd0000c13d000000400060043f00000020060000290000000006460436001f00000006001d00000024022000390000000005250019000000000035004b000000f80000213d000000000004004b000004850000613d0000002004000029000000000621034f000000000606043b0000092f0060009c000000f80000213d000000200440003900000000006404350000002002200039000000000052004b0000047c0000413d0000004402100370000000000202043b001e00000002001d000000ff0020008c000000f80000213d0000006402100370000000000402043b000009300040009c000000f80000213d0000002302400039000000000032004b00000000050000190000096b050080410000096b02200197000000000002004b00000000060000190000096b060040410000096b0020009c000000000605c019000000000006004b000000f80000c13d0000000405400039000000000251034f000000000202043b000009300020009c000000fd0000213d0000001f06200039000009c8066001970000003f06600039000009c806600197000000400700043d0000000006670019001d00000007001d000000000076004b00000000070000390000000107004039000009300060009c000000fd0000213d0000000100700190000000fd0000c13d0000002404400039000000400060043f0000001d060000290000000006260436001c00000006001d0000000004420019000000000034004b000000f80000213d0000002004500039000000000541034f000009c8062001980000001f0720018f0000001c04600029000004c10000613d000000000805034f0000001c09000029000000008a08043c0000000009a90436000000000049004b000004bd0000c13d000000000007004b000004ce0000613d000000000565034f0000000306700210000000000704043300000000076701cf000000000767022f000000000505043b0000010006600089000000000565022f00000000056501cf000000000575019f00000000005404350000001c0220002900000000000204350000008402100370000000000202043b001b00000002001d000009300020009c000000f80000213d000000a402100370000000000402043b000009300040009c000000f80000213d0000002302400039000000000032004b00000000050000190000096b050080410000096b02200197000000000002004b00000000060000190000096b060040410000096b0020009c000000000605c019000000000006004b000000f80000c13d0000000405400039000000000251034f000000000202043b000009300020009c000000fd0000213d0000001f06200039000009c8066001970000003f06600039000009c806600197000000400700043d0000000006670019001a00000007001d000000000076004b00000000070000390000000107004039000009300060009c000000fd0000213d0000000100700190000000fd0000c13d0000002404400039000000400060043f0000001a060000290000000006260436001900000006001d0000000004420019000000000034004b000000f80000213d0000002003500039000000000331034f000009c8042001980000001f0520018f00000019014000290000050c0000613d000000000603034f0000001907000029000000006806043c0000000007870436000000000017004b000005080000c13d000000000005004b000005190000613d000000000343034f0000000304500210000000000501043300000000054501cf000000000545022f000000000303043b0000010004400089000000000343022f00000000034301cf000000000353019f00000000003104350000001901200029000000000001043500000020010000290000000001010433000000200010008c0000149e0000413d000000400100043d000009bb020000410000000000210435000000040210003900000001030000390000000000320435000014a50000013d0000000001000416000000000001004b000000f80000c13d24a81dcd0000040f000000400100043d002000000001001d24a81d840000040f0000000a01000039000000000101041a00000050021002700000092f0220019700000020060000290000006003600039000000000023043500000040021002700000ffff0220018f0000004004600039000000000024043500000020021002700000092b02200197000000200560003900000000002504350000092b01100197000000000016043500000080026000390000000b06000039000000000606041a0000092f066001970000000000620435000000400600043d000000000116043600000000050504330000092b05500197000000000051043500000000010404330000ffff0110018f0000004004600039000000000014043500000000010304330000092f011001970000006003600039000000000013043500000000010204330000092f01100197000000800260003900000000001204350000092b0060009c0000092b060080410000004001600210000009b0011001c7000024a90001042e000000e40030008c000000f80000413d0000000002000416000000000002004b000000f80000c13d0000006402100370000000000202043b000009300020009c000000f80000213d0000002304200039000000000034004b000000f80000813d0000000404200039000000000541034f000000000505043b001f00000005001d000009300050009c000000f80000213d00000024052000390020001f0050002d000000200030006b000000f80000213d0000008406100370000000000606043b000009300060009c000000f80000213d0000002307600039000000000037004b000000f80000813d0000000407600039000000000771034f000000000707043b000009300070009c000000f80000213d001600050070021800000016066000290000002406600039000000000036004b000000f80000213d000000a406100370000000000606043b000009300060009c000000f80000213d0000002307600039000000000037004b000000f80000813d0000000407600039000000000771034f000000000707043b000009300070009c000000f80000213d0000001f02000029000000200020008c000000f80000413d001500050070021800000015066000290000002406600039000000000036004b000000f80000213d0000002003400039000000000331034f000000000303043b000009300030009c000000f80000213d0000000003530019001400000003001d0000002003300069000009680030009c000000f80000213d000000800030008c000000f80000413d0000010002000039000000400020043f0000001403100360000000000303043b000009300030009c000000f80000213d0000001402300029001900000002001d0000001f03200039000000200030006c000000f80000813d0000001903100360000000000303043b000009300030009c000000fd0000213d00000005043002100000003f054000390000095f055001970000097e0050009c000000fd0000213d0000010005500039000000400050043f000001000030043f0000001902000029000000200220003900000000050200190000000004240019001800000004001d000000200040006c000000f80000213d000000000003004b000013010000c13d0000010002000039000000800020043f00000014020000290000002003200039000000000331034f000000000303043b000009300030009c000000f80000213d0000001402300029001b00000002001d0000001f022000390000002004000029000000000042004b00000000030000190000096b030080410000096b022001970000096b04400197000000000542013f000000000042004b00000000020000190000096b020040410000096b0050009c000000000203c019000000000002004b000000f80000c13d0000001b02100360000000000302043b000009300030009c000000fd0000213d00000005053002100000003f025000390000095f02200197000000400700043d0000000006270019001900000007001d000000000076004b00000000070000390000000107004039000009300060009c000000fd0000213d0000000100700190000000fd0000c13d000000400060043f000000190200002900000000003204350000001b02000029001d00200020003d0000001d03500029001a00000003001d000000200030006c000000f80000213d0000001a030000290000001d0030006b000017cc0000813d001c00190000002d00000ab50000013d0000000002000416000000000002004b000000f80000c13d0000000c02000039000000000202041a001e00000002001d000009300020009c000000fd0000213d0000001e0200002900000005042002100000003f024000390000095f05200197000009600050009c000000fd0000213d000000000131034f0000008002500039000000400020043f0000001e02000029000000800020043f0000001f0240018f000000000004004b000006150000613d000000a003400039000000a006000039000000000701034f000000007807043c0000000006860436000000000036004b000006110000c13d000000000002004b000000400600043d0000000003560019001d00000006001d000000000063004b00000000050000390000000105004039000009300030009c000000fd0000213d0000000100500190000000fd0000c13d000000400030043f0000001d080000290000001e030000290000000003380436001c00000003001d000000000004004b0000062d0000613d0000001c034000290000001c04000029000000001501043c0000000004540436000000000034004b000006290000c13d000000000002004b00000080070000390000001e0000006b000007e00000c13d000000400100043d000000400200003900000000022104360000004003100039000000800400043d00000000004304350000006003100039000000000004004b000006420000613d0000000005000019000000200770003900000000060704330000092f0660019700000000036304360000000105500039000000000045004b0000063b0000413d0000000004130049000000000042043500000000040804330000000002430436000000000004004b000006500000613d0000000003000019000000200880003900000000050804330000092f0550019700000000025204360000000103300039000000000043004b000006490000413d00000000021200490000092b0020009c0000092b0200804100000060022002100000092b0010009c0000092b010080410000004001100210000000000112019f000024a90001042e0000000001000416000000000001004b000000f80000c13d0000016001000039000000400010043f000000800000043f000000a00000043f000000c00000043f000000e00000043f000001000000043f000001200000043f000001400000043f24a81da10000040f0000000002000412003400000002001d003300200000003d002000000001001d000080050100003900000044030000390000000004000415000000340440008a0000000504400210000009700200004124a824800000040f0000092f02100197001e00000002001d00000020010000290000000001210436001f00000001001d0000000001000412003200000001001d003100600000003d0000000004000415000000320440008a000000050440021000008005010000390000097002000041000000440300003924a824800000040f00000930011001970000001f0200002900000000001204350000000001000412003000000001001d002f00400000003d0000000004000415000000300440008a000000050440021000008005010000390000097002000041000000440300003924a824800000040f000009300110019700000020020000290000004002200039001d00000002001d00000000001204350000000001000412002e00000001001d002d00800000003d00000000040004150000002e0440008a000000050440021000008005010000390000097002000041000000440300003924a824800000040f0000092f0110019700000020020000290000006002200039001c00000002001d00000000001204350000000001000412002c00000001001d002b00c00000003d00000000040004150000002c0440008a000000050440021000008005010000390000097002000041000000440300003924a824800000040f0000092f0110019700000020020000290000008002200039001b00000002001d00000000001204350000000001000412002a00000001001d002900e00000003d00000000040004150000002a0440008a000000050440021000008005010000390000097002000041000000440300003924a824800000040f0000092f011001970000002002000029000000a002200039001a00000002001d00000000001204350000000001000412002800000001001d002701000000003d0000000004000415000000280440008a000000050440021000008005010000390000097002000041000000440300003924a824800000040f0000002002000029000000c0022000390000092f011001970000000000120435000000400100043d0000001e0300002900000000033104360000001f040000290000000004040433000009300440019700000000004304350000001d0300002900000000030304330000093003300197000000400410003900000000003404350000001c0300002900000000030304330000092f03300197000000600410003900000000003404350000001b0300002900000000030304330000092f03300197000000800410003900000000003404350000001a0300002900000000030304330000092f03300197000000a004100039000000000034043500000000020204330000092f02200197000000c00310003900000000002304350000092b0010009c0000092b010080410000004001100210000009c7011001c7000024a90001042e0000000001000416000000000001004b000000f80000c13d0000000201000039000000000101041a0000092f01100197000000800010043f0000099d01000041000024a90001042e0000001e020000290000002002200039001e00000002001d0000001f0400002900000000004204350000001d0200002900000020022000390000001b0020006c00000bcb0000813d001d00000002001d000000000221034f000000000202043b000009300020009c000000f80000213d0000001c022000290000003f04200039000000000034004b00000000050000190000096b050080410000096b04400197000000000004004b00000000060000190000096b060040410000096b0040009c000000000605c019000000000006004b000000f80000c13d002000200020003d0000002004100360000000000404043b000009300040009c000000fd0000213d00000005054002100000003f065000390000095f06600197000000400700043d0000000006670019001f00000007001d000000000076004b00000000070000390000000107004039000009300060009c000000fd0000213d0000000100700190000000fd0000c13d000000400060043f0000001f060000290000000000460435000000400d2000390000000002d50019000000000032004b000000f80000213d00000000002d004b000006fa0000813d0000001f040000290000073a0000013d00000020044000390000000005ec001900000000000504350000000000f40435000000200dd0003900000000002d004b00000000090b0019000006fa0000813d0000000005d1034f000000000505043b000009300050009c000000f80000213d00000020065000290000003f05600039000000000035004b00000000070000190000096b070080410000096b05500197000000000005004b00000000080000190000096b080040410000096b0050009c000000000807c019000000000008004b000000f80000c13d0000002005600039000000000751034f000000000e07043b0000093000e0009c000000fd0000213d0000001f07e00039000000000797016f0000003f07700039000000000797016f000000400f00043d00000000087f00190000000000f8004b00000000070000390000000107004039000009300080009c000000fd0000213d0000000100700190000000fd0000c13d0000004006600039000000400080043f000000000cef043600000000066e0019000000000036004b000000f80000213d0000002005500039000000000751034f000000000b090019000000000a9e01700000000005ac00190000076f0000613d000000000807034f00000000060c0019000000008908043c0000000006960436000000000056004b0000076b0000c13d0000001f06e00190000007320000613d0000000007a7034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000007320000013d0000094101000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f000009ae01000041000000c40010043f000009aa01000041000024aa000104300000093303300197000000000313019f000000000032041b000000800010043f00000000010004140000092b0010009c0000092b01008041000000c001100210000009b2011001c70000800d020000390000000103000039000009b304000041000007db0000013d000001a00200043d000009320420019700000000023400a9000000000051004b0000079c0000613d00000000033200d9000000000043004b0000013f0000c13d0000001f0020002a0000013f0000413d0000001f022000290000001e0020006b00000000030200190000001e03004029000001200030043f0000092b01100197000001400010043f000000400100043d0000000002310436000001400300043d0000092b033001970000000000320435000001600200043d000000000002004b0000000002000039000000010200c03900000040031000390000000000230435000001800200043d000009320220019700000060031000390000000000230435000001a00200043d0000093202200197000000800310003900000000002304350000092b0010009c0000092b010080410000004001100210000009b0011001c7000024a90001042e0000094101000041000000800010043f0000002001000039000000840010043f0000001601000039000000a40010043f000009a901000041000000c40010043f000009aa01000041000024aa00010430000000400100043d00000000002104350000092b0010009c0000092b010080410000004001100210000009a8011001c7000024a90001042e0000000101000039000000000201041a0000093302200197000000000262019f000000000021041b00000000010004140000092b0010009c0000092b01008041000000c001100210000009ab011001c70000800d020000390000000303000039000009ac0400004124a8249e0000040f0000000100200190000000f80000613d0000000001000019000024a90001042e00000000040000190000000c01000039000000000101041a000000000041004b00001cca0000a13d002000000004001d000009610140009a000000000101041a001f00000001001d000000000010043f0000000e01000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000800200043d0000002004000029000000000042004b0000001d0800002900001cca0000a13d000000000101043b000000000101041a0000092f021001970000000501400210000000a00310003900000000002304350000000002080433000000000042004b00001cca0000a13d0000001c011000290000001f020000290000092f02200197000000000021043500000001044000390000001e0040006c0000008007000039000007e10000413d000006310000013d00000080050000390000000006230049000009680060009c000000f80000213d000000400060008c000000f80000413d000000400600043d0000096c0060009c000000fd0000213d0000004007600039000000400070043f000000000721034f000000000707043b0000092f0070009c000000f80000213d00000000077604360000002008200039000000000881034f000000000808043b0000092f0080009c000000f80000213d0000002005500039000000000087043500000000006504350000004002200039000000000042004b0000080d0000413d000002a00000013d000009a70100004100000000001504350000000401500039000000200300002900000000003104350000000001000414000000040020008c000008390000c13d0000000103000031000000200030008c00000020040000390000000004034019000008660000013d0000096301000041000000e00010043f0000096401000041000024aa000104300000092b0050009c0000092b03000041000000000305401900000040033002100000092b0010009c0000092b01008041000000c001100210000000000131019f00000972011001c7002000000005001d24a824a30000040f00000060031002700000092b03300197000000200030008c000000200400003900000000040340190000001f0640018f0000002007400190000000200b0000290000002005700029000008540000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b000008500000c13d000000000006004b000008610000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f00030000000103550000000100200190000008b90000613d00000000050b00190000001f01400039000000600210018f0000000001520019000000000021004b00000000020000390000000102004039000009300010009c000000fd0000213d0000000100200190000000fd0000c13d000000400010043f000000200030008c000000f80000413d0000000002050433000009300020009c000007c80000a13d000000f80000013d0000000402000039000000000202041a000000800520027000000000034500a900000000044300d9000000000054004b0000013f0000c13d0000093204600197000000000043001a0000013f0000413d00000000034300190000093202200197000000000032004b000000000203801900000080011002100000093801100197000000000112019f0000000304000039000000000204041a0000093302200197000000000121019f000000a00200043d00000932022001970000093203100197000000000032004b00000000030240190000096501100197000000000113019f000000800300043d000000000003004b0000000003000019000009370300c041000000000131019f000000000014041b000000c00100043d0000008001100210000000000121019f0000000402000039000000000012041b0000000001000039000000010100c039000000400200043d0000000001120436000000a00300043d00000932033001970000000000310435000000c00100043d0000093201100197000000400320003900000000001304350000092b0020009c0000092b02008041000000400120021000000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f00000966011001c70000800d020000390000000103000039000009670400004124a8249e0000040f0000000100200190000007de0000c13d000000f80000013d0000001f0530018f0000092d06300198000000400200043d0000000004620019000008c40000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000008c00000c13d000000000005004b000008d10000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f000000000014043500000060013002100000092b0020009c0000092b020080410000004002200210000000000112019f000024aa000104300000000002060019001c01200000003d0018002000300092000000200b00008a000008ed0000013d0000001f0500002900000160025000390000001d0400002900000000004204350000001e02000029000001a002200039000000000221034f000000000202043b000001800450003900000000002404350000001c020000290000000002520436001c00000002001d0000001b020000290000002002200039000000190020006c000003730000813d001b00000002001d000000000221034f000000000202043b000009300020009c000000f80000213d0000001a04200029001e00000004001d0000001802400069000009680020009c000000f80000213d000001a00020008c000000f80000413d000000400200043d001f00000002001d000009950020009c000000fd0000213d0000001f02000029000001a002200039000000400020043f0000001e020000290000002002200039000000000421034f000000000404043b000009300040009c000000f80000213d0000001f0500002900000000044504360000002002200039000000000521034f000000000505043b0000092f0050009c000000f80000213d00000000005404350000002002200039000000000421034f000000000404043b0000092f0040009c000000f80000213d0000001f05000029000000400550003900000000004504350000002002200039000000000421034f000000000404043b000009300040009c000000f80000213d0000001f06000029000000600560003900000000004504350000002004200039000000000441034f000000000404043b000000800560003900000000004504350000004002200039000000000421034f000000000404043b000000000004004b0000000005000039000000010500c039000000000054004b000000f80000c13d0000001f05000029000000a00550003900000000004504350000002002200039000000000421034f000000000404043b000009300040009c000000f80000213d0000001f05000029000000c00550003900000000004504350000002002200039000000000421034f000000000404043b0000092f0040009c000000f80000213d0000001f06000029000000e00560003900000000004504350000002004200039000000000441034f000000000404043b000001000560003900000000004504350000004002200039000000000421034f000000000404043b000009300040009c000000f80000213d0000001e064000290000003f04600039000000000034004b00000000050000190000096b050080410000096b04400197000000000004004b00000000070000190000096b070040410000096b0040009c000000000705c019000000000007004b000000f80000c13d0000002008600039000000000481034f000000000404043b000009300040009c000000fd0000213d0000001f054000390000000005b5016f0000003f055000390000000007b5016f000000400500043d0000000007750019000000000057004b00000000090000390000000109004039000009300070009c000000fd0000213d0000000100900190000000fd0000c13d0000004009600039000000400070043f00000000064504360000000007940019000000000037004b000000f80000213d0000002007800039000000000971034f0000000007b401700000000008760019000009780000613d000000000a09034f000000000c06001900000000ad0a043c000000000cdc043600000000008c004b000009740000c13d0000001f0a400190000009850000613d000000000779034f0000000309a00210000000000a080433000000000a9a01cf000000000a9a022f000000000707043b0000010009900089000000000797022f00000000079701cf0000000007a7019f0000000000780435000000000446001900000000000404350000001f04000029000001200440003900000000005404350000002002200039000000000421034f000000000404043b000009300040009c000000f80000213d0000001e054000290000003f04500039000000000034004b00000000060000190000096b060080410000096b04400197000000000004004b00000000070000190000096b070040410000096b0040009c000000000706c019000000000007004b000000f80000c13d0000002004500039000000000441034f000000000604043b000009300060009c000000fd0000213d00000005046002100000003f044000390000095f07400197000000400400043d0000000007740019000000000047004b00000000080000390000000108004039000009300070009c000000fd0000213d0000000100800190000000fd0000c13d0000004005500039000000400070043f000000000064043500000006066002100000000006560019000000000036004b000000f80000213d000000000056004b000009cf0000a13d00000000080400190000000007530049000009680070009c000000f80000213d000000400070008c000000f80000413d000000400700043d0000096c0070009c000000fd0000213d0000004009700039000000400090043f000000000951034f000000000909043b0000092f0090009c000000f80000213d00000020088000390000000009970436000000200a500039000000000aa1034f000000000a0a043b0000000000a9043500000000007804350000004005500039000000000065004b000009b70000413d0000001f05000029000001400550003900000000004504350000002002200039000000000221034f000000000202043b000009300020009c000000f80000213d0000001e022000290000003f04200039000000000034004b00000000050000190000096b050080410000096b04400197000000000004004b00000000060000190000096b060040410000096b0040009c000000000605c019000000000006004b000000f80000c13d002000200020003d0000002004100360000000000504043b000009300050009c000000fd0000213d00000005065002100000003f046000390000095f04400197000000400800043d0000000007480019001d00000008001d000000000087004b00000000040000390000000104004039000009300070009c000000fd0000213d0000000100400190000000fd0000c13d000000400070043f0000001d04000029000000000054043500000040052000390000000002560019000000000032004b000000f80000213d000000000025004b000008dc0000813d0000001d0f00002900000a090000013d000000200ff000390000000004640019000000000004043500000000008f04350000002005500039000000000025004b000000000b0e0019000008dc0000813d000000000451034f000000000404043b000009300040009c000000f80000213d000000200a4000290000003f04a00039000000000034004b00000000060000190000096b060080410000096b04400197000000000004004b00000000070000190000096b070040410000096b0040009c000000000706c019000000000007004b000000f80000c13d0000002009a00039000000000491034f000000000604043b000009300060009c000000fd0000213d0000001f046000390000000004b4016f0000003f044000390000000004b4016f000000400800043d0000000007480019000000000087004b00000000040000390000000104004039000009300070009c000000fd0000213d0000000100400190000000fd0000c13d000000400aa00039000000400070043f00000000046804360000000007a60019000000000037004b000000f80000213d0000002007900039000000000d71034f000000000e0b0019000000000cb601700000000009c4001900000a3e0000613d00000000070d034f000000000a040019000000007b07043c000000000aba043600000000009a004b00000a3a0000c13d0000001f0760019000000a010000613d000000000acd034f0000000307700210000000000b090433000000000b7b01cf000000000b7b022f000000000a0a043b0000010007700089000000000a7a022f00000000077a01cf0000000007b7019f000000000079043500000a010000013d000000000200041a0000093302200197000000000112019f000000000010041b0000001701000029000009340010009c000000fd0000213d0000001702000029000000a001200039000000400010043f00000019010000290000000001120436001500000001001d0000093501000041000000000010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000936011001c70000800b0200003924a824a30000040f000000010020019000001c500000613d000000000101043b0000092b0210019700000015030000290000000000230435000000170300002900000080023000390000001805000029000000000052043500000060023000390000001904000029000000000042043500000040023000390000001a030000290000000000320435000000000003004b0000000002000019000009370200c04100000080011002100000093801100197000000000112019f0000000302000039000000000302041a0000093903300197000000000131019f000000000141019f000000000012041b0000008001500210000000000141019f0000000402000039000000000012041b0000093a01000041000000000010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000936011001c70000800b0200003924a824a30000040f000000010020019000001c500000613d000000000101043b000000800010043f0000001e0100002900000000010104330000092f0010019800000a9b0000613d000000160100002900000000010104330000092f0210019800000a9b0000613d0000001b0100002900000000010104330000092f00100198000012510000c13d000000400100043d000009c40200004100000000002104350000092b0010009c0000092b0100804100000040011002100000093c011001c7000024aa0001043000000969010000410000001a0200002900000000001204350000092b0020009c0000092b0200804100000040012002100000093c011001c7000024aa000104300000001c020000290000002002200039001c00000002001d0000001e0300002900000000003204350000001d020000290000002002200039001d00000002001d0000001a0020006c000017cc0000813d0000001d02100360000000000302043b000009300030009c000000f80000213d0000001b033000290000003f02300039000000200020006c00000000050000190000096b050080410000096b02200197000000000642013f000000000042004b00000000020000190000096b020040410000096b0060009c000000000205c019000000000002004b000000f80000c13d000000200b3000390000000002b1034f000000000502043b000009300050009c000000fd0000213d00000005065002100000003f026000390000095f02200197000000400800043d0000000007280019001e00000008001d000000000087004b00000000080000390000000108004039000009300070009c000000fd0000213d0000000100800190000000fd0000c13d000000400070043f0000001e020000290000000000520435000000400d3000390000000003d60019000000200030006c000000f80000213d00000000003d004b00000aab0000813d0000001e0e00002900000aeb0000013d000000200ee000390000000002f90019000000000002043500000000006e0435000000200dd0003900000000003d004b00000aab0000813d0000000002d1034f000000000502043b000009300050009c000000f80000213d0000000005b500190000003f02500039000000200020006c00000000060000190000096b060080410000096b02200197000000000742013f000000000042004b00000000020000190000096b020040410000096b0070009c000000000206c019000000000002004b000000f80000c13d0000002008500039000000000281034f000000000f02043b0000093000f0009c000000fd0000213d0000001f02f00039000009c8022001970000003f02200039000009c802200197000000400600043d0000000007260019000000000067004b00000000090000390000000109004039000009300070009c000000fd0000213d0000000100900190000000fd0000c13d0000004002500039000000400070043f0000000009f6043600000000022f0019000000200020006c000000f80000213d0000002002800039000000000c21034f000009c807f00198000000000879001900000b200000613d000000000a0c034f000000000509001900000000a20a043c0000000005250436000000000085004b00000b1c0000c13d0000001f05f0019000000ae40000613d00000000027c034f0000000305500210000000000708043300000000075701cf000000000757022f000000000202043b0000010005500089000000000252022f00000000025201cf000000000272019f000000000028043500000ae40000013d0000001d050000290000000006230049000009680060009c000000f80000213d000000400060008c000000f80000413d000000400600043d0000096c0060009c000000fd0000213d0000004007600039000000400070043f000000000721034f000000000707043b0000092f0070009c000000f80000213d00000000077604360000002008200039000000000881034f000000000808043b0000092f0080009c000000f80000213d0000002005500039000000000087043500000000006504350000004002200039000000000042004b00000b2f0000413d000002cd0000013d000000000200001900000b510000013d0000001f020000290000000102200039000000800100043d000000000012004b000002d50000813d001f00000002001d0000000501200210000000a001100039001e00000001001d0000000001010433000000200110003900000000010104330000092f01100197002000000001001d000000000010043f0000000e01000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000000101043b000000000001041b0000002001000029000000000010043f0000000d01000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000000101043b000000000301041a000000000003004b00000b4c0000613d0000000c01000039000000000201041a000000000002004b0000013f0000613d000000010130008a000000000032004b00000b9b0000613d000000000012004b00001cca0000a13d000009a30130009a000009a30220009a000000000202041a000000000021041b000000000020043f0000000d01000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c70000801002000039001c00000003001d24a824a30000040f0000001c030000290000000100200190000000f80000613d000000000101043b000000000031041b0000000c01000039000000000301041a000000000003004b000012ef0000613d000000010130008a000009a30230009a000000000002041b0000000c02000039000000000012041b0000002001000029000000000010043f0000000d01000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000000101043b000000000001041b000000800100043d0000001f0010006c00001cca0000a13d0000001e010000290000000001010433000000002101043400000000020204330000092f02200197000000400300043d000000200430003900000000002404350000092f0110019700000000001304350000092b0030009c0000092b03008041000000400130021000000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f00000962011001c70000800d020000390000000103000039000009a50400004124a8249e0000040f000000010020019000000b4c0000c13d000000f80000013d0000001a02000029000000a00020043f00000017020000290000004402200039000000000421034f000000000404043b000009300040009c000000f80000213d00000016054000290000001f04500039000000000034004b00000000060000190000096b060080410000096b04400197000000000004004b00000000070000190000096b070040410000096b0040009c000000000706c019000000000007004b000000f80000c13d000000000451034f000000000604043b000009300060009c000000fd0000213d00000005076002100000003f047000390000095f08400197000000400400043d0000000008840019000000000048004b00000000090000390000000109004039000009300080009c000000fd0000213d0000000100900190000000fd0000c13d000000400080043f000000000064043500000020055000390000000006570019000000000036004b000000f80000213d000000000065004b00000c000000813d0000000007040019000000000851034f000000000808043b000000200770003900000000008704350000002005500039000000000065004b00000bf90000413d000000c00040043f0000002002200039000000000221034f000000000202043b000000e00020043f0000002402100370000000000202043b000009300020009c000000f80000213d0000002304200039000000000034004b00000000050000190000096b050080410000096b04400197000000000004004b00000000060000190000096b060040410000096b0040009c000000000605c019000000000006004b000000f80000c13d0000000404200039000000000441034f000000000604043b000009300060009c000000fd0000213d00000005056002100000003f045000390000095f04400197000000400700043d0000000004470019001b00000007001d000000000074004b00000000070000390000000107004039000009300040009c000000fd0000213d0000000100700190000000fd0000c13d000000400040043f0000001b040000290000000004640436001a00000004001d00000024042000390000000005450019000000000035004b000000f80000213d000000000006004b000018b10000c13d0000097001000041000000000010044300000000010004120000000400100443000000240000044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f000000010020019000001c500000613d000000000101043b002000000001001d0000093a01000041000000000010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000936011001c70000800b0200003924a824a30000040f000000010020019000001c500000613d000000000101043b000000200010006b0000194c0000c13d000000800100043d00000000020104330000001b010000290000000001010433001900000002001d000000000012004b000019540000c13d000000190000006b00001aa90000c13d00000080010000390000001b0200002924a81df70000040f0000000001000019000024a90001042e00000020050000290000002405500039000000000951034f000000000909043b0000092f0090009c000000f80000213d000000400b00043d000000200ab0003900000000009a0435000000200900003900000000009b043500080000000b001d0000096c00b0009c000000fd0000213d00000008090000290000004009900039000000400090043f0000002009500039000000000591034f000000000505043b0000092f0050009c000000f80000213d0000012009900039000000000991034f000000000909043b0000096b0a900197000000000b8a013f00000000008a004b00000000080000190000096b08004041000000000079004b00000000070000190000096b070080410000096b00b0009c000000000807c019000000000008004b000000f80000c13d000600030090002d0000000601100360000000000101043b000700000001001d000009300010009c000000f80000213d000000070100002900000005011002100000000001140049000000060400002900000020094000390000096b041001970000096b07900197000000000847013f000000000047004b00000000040000190000096b04004041000f00000009001d000000000019004b00000000010000190000096b010020410000096b0080009c000000000401c019000000000004004b000000f80000c13d0000000001000415000100000001001d000000090100002900000005011002100000003f011000390000095f01100197000000400700043d0000000004170019001a00000007001d000000000074004b00000000010000390000000101004039000009300040009c000000fd0000213d0000000100100190000000fd0000c13d0000000001000031000000400040043f0000001a0400002900000009070000290000000004740436000500000004001d0000000004360019000000000014004b000000f80000213d000000000034004b00000cd00000a13d00000002060003670000001a070000290000000008310049000009680080009c000000f80000213d000000400080008c000000f80000413d000000400800043d0000096c0080009c000000fd0000213d0000004009800039000000400090043f000000000936034f000000000909043b0000092f0090009c000000f80000213d00000020077000390000000009980436000000200a300039000000000aa6034f000000000a0a043b0000000000a9043500000000008704350000004003300039000000000043004b00000cb80000413d0012092f0050019b000400400020003d001b00000000001d000a00000000001d00000cda0000013d0000001b020000290000000102200039001b00000002001d000000090020006c00001bb60000813d0000001b02000029000000070020006c00001cca0000813d0000001b010000290000000502100210001400000002001d0000000f022000290000000201000367000000000221034f000000000302043b0000000002000031000000060420006a0000003f0440008a0000096b054001970000096b06300197000000000756013f000000000056004b00000000050000190000096b05004041000000000043004b00000000040000190000096b040080410000096b0070009c000000000504c019000000000005004b000000f80000c13d0000000f04300029000000000341034f000000000303043b000009300030009c000000f80000213d000000000632004900000020054000390000096b046001970000096b07500197000000000847013f000000000047004b00000000040000190000096b04004041000000000065004b00000000060000190000096b060020410000096b0080009c000000000406c019000000000004004b000000f80000c13d000000200030008c000000f80000413d000000000451034f000000000604043b000009300060009c000000f80000213d000000000453001900000000055600190000000003540049000009680030009c000000f80000213d000000800030008c000000f80000413d000000400300043d001900000003001d000009600030009c000000fd0000213d000000000651034f00000019030000290000008003300039000000400030043f000000000606043b000009300060009c000000f80000213d00000000085600190000001f06800039000000000046004b00000000070000190000096b070080410000096b096001970000096b06400197000000000a69013f000000000069004b00000000090000190000096b090040410000096b00a0009c000000000907c019000000000009004b000000f80000c13d000000000781034f000000000707043b000009300070009c000000fd0000213d0000001f09700039000009c8099001970000003f09900039000009c8099001970000000009390019000009300090009c000000fd0000213d0000002008800039000000400090043f00000000007304350000000009870019000000000049004b000000f80000213d000000000a81034f000009c80b7001980000001908000029000000a0088000390000000009b8001900000d4c0000613d000000000c0a034f000000000d08001900000000ce0c043c000000000ded043600000000009d004b00000d480000c13d0000001f0c70019000000d590000613d000000000aba034f000000030bc00210000000000c090433000000000cbc01cf000000000cbc022f000000000a0a043b000001000bb00089000000000aba022f000000000aba01cf000000000aca019f0000000000a904350000000007870019000000000007043500000019070000290000000003370436001c00000003001d0000002007500039000000000871034f000000000808043b000009300080009c000000f80000213d000000000a5800190000001f08a00039000000000048004b00000000090000190000096b090080410000096b08800197000000000b68013f000000000068004b00000000080000190000096b080040410000096b00b0009c000000000809c019000000000008004b000000f80000c13d0000000008a1034f000000000808043b000009300080009c000000fd0000213d0000001f09800039000009c8099001970000003f09900039000009c80b900197000000400900043d000000000bb9001900000000009b004b000000000c000039000000010c0040390000093000b0009c000000fd0000213d0000000100c00190000000fd0000c13d000000200ca000390000004000b0043f000000000a890436000000000bc8001900000000004b004b000000f80000213d000000000cc1034f000009c80d800198000000000bda001900000d920000613d000000000e0c034f000000000f0a001900000000e30e043c000000000f3f04360000000000bf004b00000d8e0000c13d0000001f0e80019000000d9f0000613d0000000003dc034f000000030ce00210000000000d0b0433000000000dcd01cf000000000dcd022f000000000303043b000001000cc000890000000003c3022f0000000003c301cf0000000003d3019f00000000003b043500000000038a001900000000000304350000001c0300002900000000009304350000002007700039000000000371034f000000000803043b000009300080009c000000f80000213d00000000085800190000001f03800039000000000043004b00000000050000190000096b050080410000096b03300197000000000963013f000000000063004b00000000030000190000096b030040410000096b0090009c000000000305c019000000000003004b000000f80000c13d000000000381034f000000000503043b000009300050009c000000fd0000213d0000001f03500039000009c8033001970000003f03300039000009c803300197000000400600043d0000000009360019000000000069004b000000000a000039000000010a004039000009300090009c000000fd0000213d0000000100a00190000000fd0000c13d000000200a800039000000400090043f00000000085604360000000003a50019000000000043004b000000f80000213d0000000009a1034f000009c80a5001980000000004a8001900000dd70000613d000000000b09034f000000000c08001900000000b30b043c000000000c3c043600000000004c004b00000dd30000c13d0000001f0b50019000000de40000613d0000000003a9034f0000000309b00210000000000a040433000000000a9a01cf000000000a9a022f000000000303043b0000010009900089000000000393022f00000000039301cf0000000003a3019f00000000003404350000000003580019000000000003043500000019030000290000004003300039000e00000003001d00000000006304350000002003700039000000000331034f000000000403043b0000092b0040009c000000f80000213d00000019030000290000006003300039001700000003001d0000000000430435000000800400043d000000000004004b00000dff0000613d0000001b0040006c00001cca0000a13d0000001403000029000000a00330003900000000030304330000092b0430019800000dff0000613d000000170300002900000000004304350000001b0500002900000006035002100000000403300029000000000331034f000000000303043b000d00000003001d0000001e0050006c00001cca0000813d0000001f0320006a00000014050000290000001d04500029000000000441034f000000000404043b000000430330008a0000096b053001970000096b06400197000000000756013f000000000056004b00000000050000190000096b05004041000000000034004b00000000030000190000096b030080410000096b0070009c000000000503c019000000000005004b000000f80000c13d0000001d05400029000000000351034f000000000403043b000009300040009c000000f80000213d000000000342004900000020065000390000096b053001970000096b07600197000000000857013f000000000057004b00000000050000190000096b05004041000000000036004b00000000030000190000096b030020410000096b0080009c000000000503c019000000000005004b000000f80000c13d0000001f03400039000009c8033001970000003f03300039000009c803300197000000400700043d0000000005370019001100000007001d000000000075004b00000000070000390000000107004039000009300050009c000000fd0000213d0000000100700190000000fd0000c13d000000400050043f000000110300002900000000054304360000000003640019000000000023004b000000f80000213d000000000261034f000009c806400198000000000165001900000e4c0000613d000000000702034f0000000008050019000000007307043c0000000008380436000000000018004b00000e480000c13d0000001f0740019000000e590000613d000000000262034f0000000303700210000000000601043300000000063601cf000000000636022f000000000202043b0000010003300089000000000232022f00000000023201cf000000000262019f0000000000210435000000000145001900000000000104350000000001000415000b00000001001d000000400100043d0000096c0010009c000000fd0000213d0000004002100039000000400020043f0000002002100039000000000002043500000000000104350000001c0100002900000000010104330000000023010434000000200030008c00001a940000c13d0000000002020433001800000002001d000004000220008a0000096d0020009c00001a940000213d000000400200043d0000096f010000410000000000120435001c00000002001d00000004012000390000001802000029000000000021043500000970010000410000000000100443000000000100041200000004001004430000010001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f000000010020019000001c500000613d000000000201043b00000000010004140000092f02200197000000040020008c00000e8f0000c13d0000000103000031000000200030008c0000002004000039000000000403401900000eb80000013d0000001c030000290000092b0030009c0000092b0300804100000040033002100000092b0010009c0000092b01008041000000c001100210000000000131019f00000972011001c724a824a30000040f00000060031002700000092b03300197000000200030008c0000002004000039000000000403401900000020064001900000001c0560002900000ea70000613d000000000701034f0000001c08000029000000007907043c0000000008980436000000000058004b00000ea30000c13d0000001f0740019000000eb40000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000001bf60000613d0000001f01400039000000600210018f0000001c01200029000000000021004b00000000020000390000000102004039000009300010009c000000fd0000213d0000000100200190000000fd0000c13d000000400010043f000000200030008c000000f80000413d0000001c010000290000000001010433001c00000001001d0000092f0010009c000000f80000213d0000001c0000006b00001aa30000613d0000000001000415001500000001001d000000400100043d0000002002100039000009730400004100000000004204350000002403100039000000000043043500000024030000390000000000310435000009310010009c000000fd0000213d0000006003100039000000400030043f0000001c03000029000000040030008c00000ee10000c13d0000000001020433000000000010043f000000010300003100000f0b0000013d0000092b0020009c0000092b02008041000000400220021000000000010104330000092b0010009c0000092b010080410000006001100210000000000121019f00000974011001c70000001c0200002924a824a30000040f00000060031002700000092b03300197000000200030008c00000020050000390000000005034019000000200450019000000ef90000613d000000000601034f0000000007000019000000006806043c0000000007870436000000000047004b00000ef50000c13d0000001f0550019000000f060000613d000000000641034f0000000305500210000000000704043300000000075701cf000000000757022f000000000606043b0000010005500089000000000656022f00000000055601cf000000000575019f0000000000540435000100000003001f0003000000010355000000010020019000001aa00000613d000000000100043d000000200030008c00001aa00000413d000000000001004b00001aa00000613d000000400100043d00000020021000390000097304000041000000000042043500000024041000390000097505000041000000000054043500000024040000390000000000410435000009310010009c000000fd0000213d0000006004100039000000400040043f0000001c04000029000000040040008c00000f220000c13d0000000001020433000000000010043f00000f530000013d0000092b0020009c0000092b02008041000000400220021000000000010104330000092b0010009c0000092b010080410000006001100210000000000121019f00000974011001c70000001c0200002924a824a30000040f00000060031002700000092b03300197000000200030008c00000020050000390000000005034019000000200450019000000f3a0000613d000000000601034f0000000007000019000000006806043c0000000007870436000000000047004b00000f360000c13d0000001f0550019000000f470000613d000000000641034f0000000305500210000000000704043300000000075701cf000000000757022f000000000606043b0000010005500089000000000656022f00000000055601cf000000000575019f00000000005404350003000000010355000100000003001f0000001f0030008c00000000010000390000000101002039000000000112016f0000000002000415000000260220008a0016000500200218000000010010008c00000f580000c13d000000000100043d0000000002000415000000250220008a0016000500200218000000000001004b00001aa00000c13d000000400100043d00000020021000390000097304000041000000000042043500000024041000390000097605000041000000000054043500000024040000390000000000410435000009310010009c000000fd0000213d0000006004100039000000400040043f0000001c04000029000000040040008c00000f6b0000c13d0000000001020433000000000010043f00000f950000013d0000092b0020009c0000092b02008041000000400220021000000000010104330000092b0010009c0000092b010080410000006001100210000000000121019f00000974011001c70000001c0200002924a824a30000040f00000060031002700000092b03300197000000200030008c00000020050000390000000005034019000000200450019000000f830000613d000000000601034f0000000007000019000000006806043c0000000007870436000000000047004b00000f7f0000c13d0000001f0550019000000f900000613d000000000641034f0000000305500210000000000704043300000000075701cf000000000757022f000000000606043b0000010005500089000000000656022f00000000055601cf000000000575019f0000000000540435000100000003001f0003000000010355000000010020019000001a9d0000613d000000000100043d000000200030008c00001a9d0000413d000000000001004b00000016010000290000000501100270000000000100003f000000010100c03f00000000010004150000001501100069000000000100000200001aa30000613d00000017010000290000000001010433000000400400043d00000020034000390000097702000041000c00000003001d000000000023043500000024024000390000001203000029000000000032043500000024020000390000000000240435001700000004001d000009310040009c000000fd0000213d00000017030000290000006002300039001600000002001d000000400020043f000009780030009c000000fd0000213d0015092b0010019b00000017040000290000012001400039000000400010043f000000840200003900000016030000290000000000230435000000800340003900000000020000310000000202200367001300000003001d000000002402043c0000000003430436000000000013004b00000fc00000c13d000009790100004100000000001004430000001801000029000000040010044300000000010004140000092b0010009c0000092b01008041000000c0011002100000097a011001c7000080020200003924a824a30000040f000000010020019000001c500000613d000000000101043b000000000001004b00001a7b0000613d0000000001000414000013880110008c00001a7f0000413d00000006021002700000000001210049000000150010006c00001a830000a13d0000000c010000290000092b0010009c0000092b0100804100000040011002100000000002000414000c00000002001d000000170200002900000000020204330000092b0020009c0000092b020080410000006002200210000000000112019f0000001502000029000000c002200210000000000121019f000000180200002924a8249e0000040f000300000001035500000060031002700001092b0030019d0000092b05300197000000840050008c0000008405008039000000000300041400000016040000290000000000540435000000e006500190000000130460002900000ffe0000613d000000000701034f0000001308000029000000007907043c0000000008980436000000000048004b00000ffa0000c13d0000001f055001900000100b0000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f0000000000140435000000010020019000001bc50000613d00000016010000290000000001010433000000200010008c00001a870000c13d0000000c0130006900150015001000730000013f0000413d000000400100043d001700000001001d0000097e0010009c000000fd0000213d0000000e010000290000000001010433000e00000001001d00000019010000290000000001010433001900000001001d00000013010000290000000001010433000c00000001001d00000017020000290000010001200039000000400010043f00000008010000290000000001120436001600000001001d00000970010000410000000000100443000000000100041200000004001004430000004001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f000000010020019000001c500000613d000000000601043b0000001709000029000000e00190003900000011020000290000000000210435000000c0029000390000000e030000290000000000320435000000a0039000390000001904000029000000000043043500000080049000390000001805000029000000000054043500000060059000390000000d070000290000000000750435000000400790003900000012080000290000000000870435000009300660019700000016080000290000000000680435000000400b00043d0000002008b000390000097f06000041001300000008001d00000000006804350000002406b000390000002008000039000000000086043500000000060904330000004408b0003900000100090000390000000000980435000001440ab00039000000009806043400000000008a043500190000000b001d0000016406b00039000000000008004b000010680000613d000000000a000019000000000b6a0019000000000ca90019000000000c0c04330000000000cb0435000000200aa0003900000000008a004b000010610000413d00000000096800190000000000090435000000160900002900000000090904330000093009900197000000190b000029000000640ab0003900000000009a043500000000070704330000092f077001970000008409b0003900000000007904350000000005050433000000a407b00039000000000057043500000000040404330000092f04400197000000c405b0003900000000004504350000001f04800039000009c804400197000000e405b00039000000000303043300000120074000390000000000750435000000000664001900000000540304340000000003460436000000000004004b0000108e0000613d000000000600001900000000073600190000000008650019000000000808043300000000008704350000002006600039000000000046004b000010870000413d000000000534001900000000000504350000001f04400039000009c804400197000000000534001900000019040000290000000003450049000000440330008a00000000020204330000010404400039000000000034043500000000430204340000000002350436000000000003004b000010a50000613d000000000500001900000000062500190000000007540019000000000707043300000000007604350000002005500039000000000035004b0000109e0000413d000000000423001900000000000404350000001f03300039000009c803300197000000000223001900000019040000290000000003420049000000440330008a00000000010104330000012404400039000000000034043500000000310104340000000002120436000000000001004b000010bc0000613d000000000400001900000000052400190000000006430019000000000606043300000000006504350000002004400039000000000014004b000010b50000413d00000000032100190000000000030435000000190400002900000000024200490000001f01100039000009c8011001970000000001210019000000200210008a00000000002404350000001f01100039000009c8011001970000000002410019000000000012004b00000000010000390000000101004039001700000002001d000009300020009c000000fd0000213d0000000100100190000000fd0000c13d0000001701000029000000400010043f000009800010009c000000fd0000213d0000001703000029000000c001300039000000400010043f0000008402000039000000000323043600000000020000310000000202200367001600000003001d000000002402043c0000000003430436000000000013004b000010dc0000c13d000009790100004100000000001004430000001c01000029000000040010044300000000010004140000092b0010009c0000092b01008041000000c0011002100000097a011001c7000080020200003924a824a30000040f000000010020019000001c500000613d000000000101043b000000000001004b00001a7b0000613d0000000001000414000013880110008c00001a7f0000413d00000006021002700000000001210049000000150010006c00001a830000a13d0000000001000414001100000001001d0000001c01000029000000040010008c000011000000c13d000000030100036700000001040000310000000002000019000011160000013d00000013010000290000092b0010009c0000092b010080410000004001100210000000190200002900000000020204330000092b0020009c0000092b020080410000006002200210000000000112019f00000015020000290000092b0020009c0000092b02008041000000c002200210000000000121019f0000001c0200002924a8249e0000040f000000010220015f000300000001035500000060031002700001092b0030019d0000092b043001970000000003000414000000840040008c000000840400803900000017050000290000000000450435000000e0064001900000001605600029000011240000613d000000000701034f0000001608000029000000007907043c0000000008980436000000000058004b000011200000c13d0000001f04400190000011310000613d000000000161034f0000000304400210000000000605043300000000064601cf000000000646022f000000000101043b0000010004400089000000000141022f00000000014101cf000000000161019f0000000000150435000000010020019000001bcf0000c13d00000017010000290000000001010433000000200010008c00001a870000c13d00000016010000290000000001010433001900000001001d0000001c02000029000000120020006b000011b90000613d000000110130006900160015001000730000013f0000413d000000400300043d00000020023000390000097701000041001300000002001d000000000012043500000024013000390000001202000029000000000021043500000024010000390000000000130435001c00000003001d000009310030009c000000fd0000213d0000001c020000290000006001200039001700000001001d000000400010043f000009780020009c000000fd0000213d0000001c040000290000012001400039000000400010043f000000840200003900000017030000290000000000230435000000800340003900000000020000310000000202200367001500000003001d000000002402043c0000000003430436000000000013004b0000115d0000c13d000009790100004100000000001004430000001801000029000000040010044300000000010004140000092b0010009c0000092b01008041000000c0011002100000097a011001c7000080020200003924a824a30000040f000000010020019000001c500000613d000000000101043b000000000001004b00001a7b0000613d0000000001000414000013880110008c00001a7f0000413d00000006021002700000000001210049000000160010006c00001a830000a13d00000013010000290000092b0010009c0000092b0100804100000040011002100000000002000414001300000002001d0000001c0200002900000000020204330000092b0020009c0000092b020080410000006002200210000000000112019f00000016020000290000092b0020009c0000092b02008041000000c002200210000000000121019f000000180200002924a8249e0000040f000300000001035500000060031002700001092b0030019d0000092b05300197000000840050008c0000008405008039000000000300041400000017040000290000000000540435000000e00650019000000015046000290000119d0000613d000000000701034f0000001508000029000000007907043c0000000008980436000000000048004b000011990000c13d0000001f05500190000011aa0000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f0000000000140435000000010020019000001c020000613d00000017010000290000000001010433000000200010008c00001a870000c13d0000001301300069000000160010006c0000013f0000213d000000150100002900000000010104330000000c0210006c00001c160000413d000000190020006c00001c160000c13d000000400100043d0000096c0010009c000000fd0000213d0000004002100039000000400020043f0000002002100039000000190300002900000000003204350000001802000029000000000021043500000000020004150000000b0220006900000000020000020000001a0200002900000000020204330000001b0020006c00001cca0000a13d00000014030000290000000502300029001c00000002001d00000000001204350000001a0200002900000000020204330000001b0020006c00001cca0000a13d00000000010104330000092f01100197000000000010043f0000000d01000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000000101043b000000000101041a000000000001004b00000cd50000613d0000001a0100002900000000010104330000001b0010006c00001cca0000a13d0000000b01000039000000000201041a0000001c010000290000000001010433001700000001001d0000000031010434001800000003001d000000400400043d00000983030000410000000003340436001900000003001d0000092f01100197001c00000004001d0000000403400039000000000013043500000000010004140000092f02200197000000040020008c000012000000c13d0000000103000031000000400030008c00000040040000390000000004034019000012290000013d0000001c030000290000092b0030009c0000092b0300804100000040033002100000092b0010009c0000092b01008041000000c001100210000000000131019f00000972011001c724a824a30000040f00000060031002700000092b03300197000000400030008c0000004004000039000000000403401900000060064001900000001c05600029000012180000613d000000000701034f0000001c08000029000000007907043c0000000008980436000000000058004b000012140000c13d0000001f07400190000012250000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f0003000000010355000000010020019000001c260000613d0000001f01400039000000e00210018f0000001c01200029000000000021004b00000000020000390000000102004039000009300010009c000000fd0000213d0000000100200190000000fd0000c13d000000400010043f000000400030008c000000f80000413d0000096c0010009c000000fd0000213d0000004002100039000000400020043f0000001c020000290000000002020433000009840020009c000000f80000213d0000000001210436000000190300002900000000030304330000092b0030009c000000f80000213d0000000000310435000000000002004b00001be90000613d0000001801000029000000000301043300000000012300a900000000022100d9000000000032004b0000013f0000c13d000009850110012a0000000a0010002a0000013f0000413d000a000a0010002d00000cd50000013d000000400300043d0000093b01000041001a00000003001d00000000001304350000000001000414000000040020008c0000125d0000c13d0000000103000031000000200030008c00000020040000390000000004034019000012870000013d0000001a030000290000092b0030009c0000092b0300804100000040033002100000092b0010009c0000092b01008041000000c001100210000000000131019f0000093c011001c724a824a30000040f00000060031002700000092b03300197000000200030008c000000200400003900000000040340190000001f0640018f00000020074001900000001a05700029000012760000613d000000000801034f0000001a09000029000000008a08043c0000000009a90436000000000059004b000012720000c13d000000000006004b000012830000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f00030000000103550000000100200190000012f50000613d0000001f01400039000000600110018f0000001a02100029000000000012004b00000000010000390000000101004039000009300020009c000000fd0000213d0000000100100190000000fd0000c13d000000400020043f000000200030008c000000f80000413d0000001a010000290000000001010433000009300010009c000000f80000213d000000010010008c0000149c0000c13d000000160100002900000000010104330000092f01100197000000a00010043f0000001f0100002900000000010104330000093001100197000000c00010043f000000200300002900000000030304330000093003300197000000e00030043f0000001e0400002900000000040404330000092f04400197000001000040043f0000001d0500002900000000050504330000092f05500197000001400050043f0000001c0500002900000000050504330000092f05500197000001600050043f0000001b0500002900000000050504330000092f05500197000001800050043f00000080052000390000000000450435000000600420003900000000003404350000004003200039000000000013043500000020012000390000093e03000041000000000031043500000080030000390000000000320435000009340020009c000000fd0000213d000000a003200039000000400030043f000000000202043324a8246b0000040f000001200010043f000000800200043d000001400000044300000160002004430000002002000039000000a00300043d0000018000200443000001a0003004430000004003000039000000c00400043d000001c000300443000001e0004004430000006003000039000000e00400043d00000200003004430000022000400443000001000300043d000000800400003900000240004004430000026000300443000000a0030000390000028000300443000002a000100443000000c001000039000001400300043d000002c000100443000002e000300443000000e001000039000001600300043d000003000010044300000320003004430000010001000039000001800300043d000003400010044300000360003004430000010000200443000000090100003900000120001004430000093f01000041000024a90001042e000009a401000041000000000010043f0000003101000039000000040010043f0000097201000041000024aa000104300000001f0530018f0000092d06300198000000400200043d0000000004620019000008c40000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000012fc0000c13d000008c40000013d0000000003050019001b01200000003d00000020020000290017002000200092000013170000013d0000001e0400002900000160024000390000001c0300002900000000003204350000001d02000029000001a002200039000000000221034f000000000202043b000001800340003900000000002304350000001b020000290000000002420436001b00000002001d0000001a030000290000002003300039000000180030006c000005c00000813d001a00000003001d000000000331034f000000000303043b000009300030009c000000f80000213d0000001903300029001d00000003001d0000001703300069000009680030009c000000f80000213d000001a00030008c000000f80000413d000000400200043d001e00000002001d000009950020009c000000fd0000213d0000001e02000029000001a003200039000000400030043f0000001d020000290000002003200039000000000431034f000000000404043b000009300040009c000000f80000213d0000001e0200002900000000044204360000002003300039000000000531034f000000000505043b0000092f0050009c000000f80000213d00000000005404350000002003300039000000000431034f000000000404043b0000092f0040009c000000f80000213d0000001e02000029000000400520003900000000004504350000002003300039000000000431034f000000000404043b000009300040009c000000f80000213d0000001e02000029000000600520003900000000004504350000002004300039000000000441034f000000000404043b000000800520003900000000004504350000004003300039000000000431034f000000000404043b000000000004004b0000000005000039000000010500c039000000000054004b000000f80000c13d0000001e02000029000000a00520003900000000004504350000002003300039000000000431034f000000000404043b000009300040009c000000f80000213d0000001e02000029000000c00520003900000000004504350000002003300039000000000431034f000000000404043b0000092f0040009c000000f80000213d0000001e02000029000000e00520003900000000004504350000002004300039000000000441034f000000000404043b000001000520003900000000004504350000004003300039000000000431034f000000000404043b000009300040009c000000f80000213d0000001d074000290000003f047000390000002002000029000000000024004b00000000050000190000096b050080410000096b044001970000096b0d2001970000000006d4013f0000000000d4004b00000000040000190000096b040040410000096b0060009c000000000405c019000000000004004b000000f80000c13d0000002006700039000000000461034f000000000404043b000009300040009c000000fd0000213d0000001f05400039000009c8055001970000003f05500039000009c808500197000000400500043d0000000008850019000000000058004b00000000090000390000000109004039000009300080009c000000fd0000213d0000000100900190000000fd0000c13d0000004009700039000000400080043f00000000074504360000000008940019000000200080006c000000f80000213d0000002006600039000000000661034f000009c8094001980000000008970019000013a50000613d000000000b06034f000000000c07001900000000ba0b043c000000000cac043600000000008c004b000013a10000c13d0000001f0a400190000013b20000613d000000000696034f0000000309a00210000000000a080433000000000a9a01cf000000000a9a022f000000000606043b0000010009900089000000000696022f00000000069601cf0000000006a6019f0000000000680435000000000447001900000000000404350000001e02000029000001200420003900000000005404350000002003300039000000000431034f000000000404043b000009300040009c000000f80000213d0000001d054000290000003f04500039000000200040006c00000000060000190000096b060080410000096b044001970000000007d4013f0000000000d4004b00000000040000190000096b040040410000096b0070009c000000000406c019000000000004004b000000f80000c13d0000002004500039000000000441034f000000000604043b000009300060009c000000fd0000213d00000005046002100000003f044000390000095f07400197000000400400043d0000000007740019000000000047004b00000000080000390000000108004039000009300070009c000000fd0000213d0000000100800190000000fd0000c13d0000004005500039000000400070043f000000000064043500000006066002100000000007560019000000200070006c000000f80000213d000000000057004b000013fd0000a13d00000000080400190000002006500069000009680060009c000000f80000213d000000400060008c000000f80000413d000000400600043d0000096c0060009c000000fd0000213d0000004009600039000000400090043f000000000951034f000000000909043b0000092f0090009c000000f80000213d00000020088000390000000009960436000000200a500039000000000aa1034f000000000a0a043b0000000000a9043500000000006804350000004005500039000000000075004b000013e50000413d0000001e02000029000001400520003900000000004504350000002003300039000000000331034f000000000303043b000009300030009c000000f80000213d0000001d033000290000003f04300039000000200040006c00000000050000190000096b050080410000096b044001970000000006d4013f0000000000d4004b00000000040000190000096b040040410000096b0060009c000000000405c019000000000004004b000000f80000c13d000000200e3000390000000004e1034f000000000404043b000009300040009c000000fd0000213d00000005064002100000003f056000390000095f05500197000000400200043d0000000005520019001c00000002001d000000000025004b00000000070000390000000107004039000009300050009c000000fd0000213d0000000100700190000000fd0000c13d000000400050043f0000001c02000029000000000042043500000040043000390000000003460019000000200030006c000000f80000213d000000000034004b000013060000813d0000001c0f000029000014370000013d000000200ff0003900000000027a0019000000000002043500000000009f04350000002004400039000000000034004b000013060000813d000000000541034f000000000505043b000009300050009c000000f80000213d0000000008e500190000003f05800039000000200050006c00000000060000190000096b060080410000096b055001970000000007d5013f0000000000d5004b00000000050000190000096b050040410000096b0070009c000000000506c019000000000005004b000000f80000c13d0000002006800039000000000561034f000000000705043b000009300070009c000000fd0000213d0000001f05700039000009c8055001970000003f05500039000009c805500197000000400900043d000000000b59001900000000009b004b000000000500003900000001050040390000093000b0009c000000fd0000213d0000000100500190000000fd0000c13d00000040058000390000004000b0043f000000000a7904360000000005570019000000200050006c000000f80000213d0000002005600039000000000551034f000009c80c7001980000000006ca00190000146c0000613d000000000b05034f00000000080a001900000000b20b043c0000000008280436000000000068004b000014680000c13d0000001f08700190000014300000613d0000000002c5034f0000000305800210000000000806043300000000085801cf000000000858022f000000000202043b0000010005500089000000000252022f00000000025201cf000000000282019f0000000000260435000014300000013d0000000002000415000000240220008a001e000500200218002400000000003d0000001f02000029001d00e0002000920000001d01100360000000000101043b0000092f0010009c000000f80000213d00000979020000410000000000200443000000040010044300000000010004140000092b0010009c0000092b01008041000000c0011002100000097a011001c7000080020200003924a824a30000040f000000010020019000001c500000613d000000000101043b000000000001004b0000001e010000290000000501100270000000000100003f000000010100603f000014aa0000c13d0000000001000415000000100110006900000000010000020000000001000019000024a90001042e0000093d0100004100000aa50000013d0000001e0000006b000014c50000c13d000000400100043d000009bb020000410000000000210435000000040210003900000000000204350000092b0010009c0000092b01008041000000400110021000000972011001c7000024aa000104300000001d010000290000000201100367000000000101043b001e00000001001d0000092f0010009c000000f80000213d0000000001000415001c00000001001d000000400100043d0000002002100039000009730300004100000000003204350000002404100039000000000034043500000024030000390000000000310435000009310010009c000000fd0000213d0000006003100039000000400030043f0000001e03000029000000040030008c0000186f0000c13d0000000001020433000000000010043f00000001030000310000189a0000013d000000000100041a0000092f011001970000000002000411000000000012004b000017bb0000c13d0000001d010000290000000001010433000009680010009c000000f80000213d000000a00010008c000000f80000413d000000400100043d001800000001001d000009340010009c000000fd0000213d0000001801000029000000a001100039000000400010043f0000001c0100002900000000010104330000092b0010009c000000f80000213d00000018020000290000000002120436001700000002001d0000001d02000029000000400220003900000000020204330000092b0020009c000000f80000213d000000170300002900000000002304350000001d03000029000000600330003900000000040304330000ffff0040008c000000f80000213d00000018030000290000004003300039001600000003001d00000000004304350000001d03000029000000800330003900000000050304330000092f0050009c000000f80000213d00000018030000290000006003300039001400000003001d00000000005304350000001d03000029000000a00330003900000000030304330000092f0030009c000000f80000213d00000018060000290000008006600039001300000006001d0000000000360435000000000005004b00000a9b0000613d0000004004400210000009b5044001970000002002200210000009b602200197000000000242019f0000005004500210000009b704400197000000000242019f0000000a04000039000000000504041a000009b805500197000000000252019f000000000112019f000000000014041b0000000b01000039000000000201041a0000093302200197000000000232019f000000000021041b000000400100043d001200000001001d0000092e0010009c000000fd0000213d0000001201000029000000e001100039000000400010043f00000970010000410000000000100443000000000100041200000004001004430000002001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f000000010020019000001c500000613d000000000101043b0000092f0110019700000012020000290000000001120436001100000001001d00000970010000410000000000100443000000000100041200000004001004430000006001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f000000010020019000001c500000613d000000000101043b00000930011001970000001102000029000000000012043500000970010000410000000000100443000000000100041200000004001004430000004001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f000000010020019000001c500000613d00000012020000290000004002200039000000000101043b0000093001100197001000000002001d000000000012043500000970010000410000000000100443000000000100041200000004001004430000008001000039001500000001001d000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f000000010020019000001c500000613d00000012020000290000006002200039000000000101043b0000092f01100197000f00000002001d00000000001204350000097001000041000000000010044300000000010004120000000400100443000000c001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f000000010020019000001c500000613d00000012020000290000008002200039000000000101043b0000092f01100197000e00000002001d00000000001204350000097001000041000000000010044300000000010004120000000400100443000000e001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f000000010020019000001c500000613d0000001202000029000000a002200039000000000101043b0000092f01100197000d00000002001d000000000012043500000970010000410000000000100443000000000100041200000004001004430000010001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f000000010020019000001c500000613d000000000101043b0000092f011001970000001203000029000000c002300039000000000012043500000000010304330000092f03100197000000400100043d00000000033104360000001104000029000000000404043300000930044001970000000000430435000000100300002900000000030304330000093003300197000000400410003900000000003404350000000f0300002900000000030304330000092f03300197000000600410003900000000003404350000000e0300002900000000030304330000092f03300197000000800410003900000000003404350000000d0300002900000000030304330000092f03300197000000a004100039000000000034043500000000020204330000092f02200197000000c0031000390000000000230435000000180200002900000000020204330000092b02200197000000e0031000390000000000230435000000170200002900000000020204330000092b0220019700000100031000390000000000230435000000160200002900000000020204330000ffff0220018f00000120031000390000000000230435000000140200002900000000020204330000092f0220019700000140031000390000000000230435000000130200002900000000020204330000092f02200197000001600310003900000000002304350000092b0010009c0000092b01008041000000400110021000000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f000009b9011001c70000800d020000390000000103000039000009ba0400004124a8249e0000040f0000000100200190000000f80000613d0000000901000039000000000101041a001700000001001d000000000001004b000016150000613d001800000000001d0000000901000039000000000101041a000000180010006c00001cca0000a13d0000001801000029000009990110009a000000000101041a0000092f01100197000000000010043f0000000801000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000000101043b000000000001041b00000018020000290000000102200039001800000002001d000000170020006c000015fa0000413d00000020010000290000000001010433001300000001001d000000000001004b0000000001000019000016710000613d001800000000001d00000020010000290000000001010433000000180010006c00001cca0000a13d000000180100002900000005011002100000001f0110002900000000010104330000092f01100197001600000001001d000000000010043f0000000801000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000000101043b000000000101041a0000000801100270000000ff0110018f000000030010008c000018690000813d000000400200043d001700000002001d000000000001004b00001cf20000c13d000000160000006b00001cfd0000613d00000017010000290000096c0010009c000000fd0000213d00000017020000290000004001200039000000400010043f0000001801000029000000ff0110018f00000000021204360000000201000039001400000002001d00000000001204350000001601000029000000000010043f0000000801000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000000101043b000000000201041a000009c90220019700000017030000290000000003030433000000ff0330018f000000000232019f000000000021041b00000014030000290000000003030433000000020030008c000018690000213d000009ca0220019700000008033002100000ff000330018f000000000223019f000000000021041b00000018020000290000000102200039001800000002001d000000130020006c0000161c0000413d00000020010000290000000001010433000009300010009c000000fd0000213d0000000902000039000000000302041a000000000012041b000000000031004b0000167e0000813d000009990210009a000009990330009a000000000032004b0000167e0000813d000000000002041b0000000102200039000000000032004b0000167a0000413d0000000902000039000000000020043f000000000001004b0000168c0000613d00000000020000190000002003000029000000200330003900000000040304330000092f04400197000009990520009a000000000045041b0000000102200039000000000012004b000016840000413d0000000601000039000000000201041a000009bd02200197000000130300002900000008033002100000ff000330018f000000000223019f0000001e022001af000000000021041b0000000701000039000000000101041a0000092b021001970000092b0020009c0000013f0000613d000009be01100197001f00010020003d0000001f011001af0000000702000039000000000012041b0000093a01000041000000000010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000936011001c70000800b0200003924a824a30000040f000000010020019000001c500000613d000000000301043b000000400100043d00000080021000390000012004000039000000000042043500000060021000390000001f04000029000000000042043500000000020004100000092f0220019700000040041000390000000000240435000000200210003900000000003204350000014003100039000000800400043d00000000004304350000016003100039000000000004004b000016c80000613d000000000500001900000015060000290000002006600039001500000006001d00000000060604330000092f0660019700000000036304360000000105500039000000000045004b000016bf0000413d0000000004130049000000200440008a000000a0051000390000000000450435000000200400002900000000040404330000000003430436000000000004004b000016da0000613d00000000050000190000002006000029000000200660003900000000070604330000092f0770019700000000037304360000000105500039000000000045004b000016d30000413d000000c0041000390000001e0500002900000000005404350000000004130049000000200440008a000000e00510003900000000004504350000001d0400002900000000040404330000000003430436000000000004004b000016ee0000613d000000000500001900000000063500190000001c07500029000000000707043300000000007604350000002005500039000000000045004b000016e70000413d000000000534001900000000000504350000001f04400039000009c80440019700000000043400190000000003140049000000200330008a000001200510003900000000003504350000001b0300002900000930053001970000010003100039001f00000005001d00000000005304350000001a0300002900000000030304330000000004340436000000000003004b000017090000613d000000000500001900000000064500190000001907500029000000000707043300000000007604350000002005500039000000000035004b000017020000413d0000000005430019000000000005043500000000041400490000001f03300039000009c8033001970000000003430019000000200430008a00000000004104350000001f03300039000009c8043001970000000003140019000000000043004b00000000040000390000000104004039000009300030009c000000fd0000213d0000000100400190000000fd0000c13d000000400030043f0000092b0020009c0000092b02008041000000400220021000000000010104330000092b0010009c0000092b010080410000006001100210000000000121019f00000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f000009ab011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000000101043b000009bf01100197000009c0011001c70000000502000039000000000012041b0000000701000039000000000101041a001b00000001001d000009c101000041000000000010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000936011001c70000800b0200003924a824a30000040f000000010020019000001c500000613d000000000101043b0000002001100210000009b6011001970000000704000039000000000204041a000009c203200197000000000113019f000000000014041b0000000501000039000000000301041a000000400100043d0000006004100039000001200500003900000000005404350000092b0220019700000040041000390000000000240435000000200210003900000000003204350000001b0200002900000020022002700000092b0220019700000000002104350000012002100039000000800300043d00000000003204350000014002100039000000000003004b000017670000613d00000080040000390000000005000019000000200440003900000000060404330000092f0660019700000000026204360000000105500039000000000035004b000017600000413d000000000312004900000080041000390000000000340435000000200300002900000000030304330000000002320436000000000003004b000017790000613d000000000400001900000020050000290000002005500039002000000005001d00000000050504330000092f0550019700000000025204360000000104400039000000000034004b000017700000413d0000000003120049000000c0041000390000000000340435000000a0031000390000001e0400002900000000004304350000001d0300002900000000030304330000000002320436000000000003004b0000178c0000613d000000000400001900000000052400190000001c06400029000000000606043300000000006504350000002004400039000000000034004b000017850000413d00000000042300190000000000040435000000e0041000390000001f0500002900000000005404350000001f03300039000009c80330019700000000022300190000000003120049000001000410003900000000003404350000001a0300002900000000030304330000000002320436000000000003004b000017a40000613d000000000400001900000000052400190000001906400029000000000606043300000000006504350000002004400039000000000034004b0000179d0000413d000000000423001900000000000404350000001f03300039000009c803300197000000000212004900000000023200190000092b0020009c0000092b0200804100000060022002100000092b0010009c0000092b010080410000004001100210000000000112019f00000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f000009ab011001c70000800d020000390000000103000039000009c304000041000007930000013d000000400100043d0000004402100039000009a9030000410000000000320435000000240210003900000016030000390000000000320435000009410200004100000000002104350000000402100039000000200300003900000000003204350000092b0010009c0000092b01008041000000400110021000000942011001c7000024aa000104300000001902000029000000a00020043f00000014020000290000004003200039000000000231034f000000000502043b000009300050009c000000f80000213d00000014055000290000001f02500039000000200020006c00000000060000190000096b060080410000096b02200197000000000742013f000000000042004b00000000020000190000096b020040410000096b0070009c000000000206c019000000000002004b000000f80000c13d000000000251034f000000000602043b000009300060009c000000fd0000213d00000005076002100000003f027000390000095f02200197000000400400043d0000000008240019000000000048004b00000000090000390000000109004039000009300080009c000000fd0000213d0000000100900190000000fd0000c13d000000400080043f000000000064043500000020055000390000000006570019000000200060006c000000f80000213d000000000065004b000018020000813d0000000002040019000000000751034f000000000707043b000000200220003900000000007204350000002005500039000000000065004b000017fb0000413d000000c00040043f0000002002300039000000000121034f000000000101043b000000e00010043f000000400200043d0000096a0020009c000000fd0000213d0000002001200039000000400010043f0000000000020435000000800100003924a81df70000040f0000000501000039000000000101041a00000004020000390000000202200367000000000202043b002000000002001d000000000021004b00001be00000c13d0000097001000041000000000010044300000000010004120000000400100443000000240000044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f000000010020019000001c500000613d000000000101043b001e00000001001d0000093a01000041000000000010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000936011001c70000800b0200003924a824a30000040f000000010020019000001c500000613d000000400200043d000000000101043b0000001e0010006b00001c510000c13d00000024010000390000000201100367000000000101043b00000008011002700000092b0110019700000020032000390000000000130435000000200100002900000000001204350000092b0020009c0000092b02008041000000400120021000000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f00000962011001c70000800d020000390000000103000039000009980400004124a8249e0000040f0000000100200190000000f80000613d00000000010004110000092f01100197000000000010043f0000000801000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000000f80000613d000000400200043d0000096c0020009c000000fd0000213d000000000101043b0000004003200039000000400030043f000000000301041a000000ff0130018f00000000021204360000000803300270000000ff0330018f000000020030008c00001caa0000a13d000009a401000041000000000010043f0000002101000039000000040010043f0000097201000041000024aa000104300000092b0020009c0000092b02008041000000400220021000000000010104330000092b0010009c0000092b010080410000006001100210000000000121019f00000974011001c70000001e0200002924a824a30000040f00000060031002700000092b03300197000000200030008c000000200400003900000000040340190000001f0540018f0000002004400190000018880000613d000000000601034f0000000007000019000000006806043c0000000007870436000000000047004b000018840000c13d000000000005004b000018950000613d000000000641034f0000000305500210000000000704043300000000075701cf000000000757022f000000000606043b0000010005500089000000000656022f00000000055601cf000000000575019f0000000000540435000100000003001f0003000000010355000000010020019000001a770000613d000000000100043d000000200030008c00001a770000413d000000000001004b00001a770000613d000000400100043d00000020021000390000097304000041000000000042043500000024041000390000097505000041000000000054043500000024040000390000000000410435000009310010009c000000fd0000213d0000006004100039000000400040043f0000001e04000029000000040040008c000019020000c13d0000000001020433000000000010043f000019340000013d000000240630008a0000001a07000029000018b90000013d0000000000a9043500000000078704360000002004400039000000000054004b00000c310000813d000000000841034f000000000808043b000009300080009c000000f80000213d000000000a2800190000000008a60049000009680080009c000000f80000213d000000400080008c000000f80000413d000000400800043d0000096c0080009c000000fd0000213d0000004009800039000000400090043f0000002409a00039000000000991034f000000000909043b0000000009980436000000440ba00039000000000bb1034f000000000b0b043b0000093000b0009c000000f80000213d000000000bab0019000000430ab0003900000000003a004b000000000c0000190000096b0c0080410000096b0aa0019700000000000a004b000000000d0000190000096b0d0040410000096b00a0009c000000000d0cc01900000000000d004b000000f80000c13d000000240ab00039000000000aa1034f000000000c0a043b0000093000c0009c000000fd0000213d000000050dc002100000003f0ad000390000095f0ea00197000000400a00043d000000000eea00190000000000ae004b000000000f000039000000010f0040390000093000e0009c000000fd0000213d0000000100f00190000000fd0000c13d0000004000e0043f0000000000ca0435000000440bb00039000000000cbd001900000000003c004b000000f80000213d0000000000cb004b000018b40000813d000000000d0a0019000000000eb1034f000000000e0e043b0000092b00e0009c000000f80000213d000000200dd000390000000000ed0435000000200bb000390000000000cb004b000018f80000413d000018b40000013d0000092b0020009c0000092b02008041000000400220021000000000010104330000092b0010009c0000092b010080410000006001100210000000000121019f00000974011001c70000001e0200002924a824a30000040f00000060031002700000092b03300197000000200030008c000000200400003900000000040340190000001f0540018f00000020044001900000191b0000613d000000000601034f0000000007000019000000006806043c0000000007870436000000000047004b000019170000c13d000000000005004b000019280000613d000000000641034f0000000305500210000000000704043300000000075701cf000000000757022f000000000606043b0000010005500089000000000656022f00000000055601cf000000000575019f00000000005404350003000000010355000100000003001f0000001f0030008c00000000010000390000000101002039000000000112016f0000000002000415000000220220008a001b000500200218000000010010008c000019390000c13d000000000100043d0000000002000415000000210220008a001b000500200218000000000001004b00001a770000c13d000000400100043d00000020021000390000097304000041000000000042043500000024041000390000099205000041000000000054043500000024040000390000000000410435000009310010009c000000fd0000213d0000006004100039000000400040043f0000001e04000029000000040040008c000019570000c13d0000000001020433000000000010043f000019820000013d000000400200043d00000024032000390000000000130435000009970100004100000000001204350000000401200039000000200300002900001a8e0000013d000000400100043d0000099e0200004100000a9d0000013d0000092b0020009c0000092b02008041000000400220021000000000010104330000092b0010009c0000092b010080410000006001100210000000000121019f00000974011001c70000001e0200002924a824a30000040f00000060031002700000092b03300197000000200030008c000000200400003900000000040340190000001f0540018f0000002004400190000019700000613d000000000601034f0000000007000019000000006806043c0000000007870436000000000047004b0000196c0000c13d000000000005004b0000197d0000613d000000000641034f0000000305500210000000000704043300000000075701cf000000000757022f000000000606043b0000010005500089000000000656022f00000000055601cf000000000575019f0000000000540435000100000003001f0003000000010355000000010020019000001a740000613d000000000100043d000000200030008c00001a740000413d000000000001004b0000001b010000290000000501100270000000000100003f000000010100c03f00000000010004150000001c011000690000000001000002000014970000613d0000001d01000029000000400310008a0000000202000367000000000432034f0000000a01000039000000000101041a000000000704043b000009300070009c000000f80000213d0000002003300039000000000332034f000000000303043b0000092f0030009c000000f80000213d000000400400043d00000020054000390000000000350435000000200300003900000000003404350000096c0040009c000000fd0000213d0000004003400039000000400030043f0000001f052003600000000009000031000000200690006a000000230660008a000000000505043b0000096b085001970000096b0a600197000000000ba8013f0000000000a8004b00000000080000190000096b08004041000000000065004b00000000060000190000096b060080410000096b00b0009c000000000806c019000000000008004b000000f80000c13d0000000306500029000000000562034f000000000505043b000009300050009c000000f80000213d0000000008590049000000200c6000390000096b068001970000096b0ac00197000000000b6a013f00000000006a004b00000000060000190000096b0600404100000000008c004b00000000080000190000096b080020410000096b00b0009c000000000608c019000000000006004b000000f80000c13d000009340030009c000000fd0000213d000000e006400039000000400060043f0000001f06000029000000600a6000390000000006a2034f000000000606043b000000000063043500000080064000390000000000460435000000600840003900000000007804350000001f07500039000009c8077001970000003f07700039000009c807700197000000400b00043d00000000077b00190000000000b7004b000000000d000039000000010d004039000009300070009c000000fd0000213d0000000100d00190000000fd0000c13d000000400070043f000000000d5b04360000000007c50019000000000097004b000000f80000213d0020000000c20353001f09c80050019c0000001f0e50018f0000001f09d00029000019f40000613d000000200f00035f00000000070d001900000000fc0f043c0000000007c70436000000000097004b000019f00000c13d00000000000e004b00001a020000613d0000001f0c0000290000002007c0035f000000030ce00210000000000e090433000000000ece01cf000000000ece022f000000000707043b000001000cc000890000000007c7022f0000000007c701cf0000000007e7019f000000000079043500000000055d00190000000000050435000000c0074000390000001a050000290000000000570435000000a0094000390000000000b904350000014005a0008a000000000452034f000000000404043b0000092f0040009c000000f80000213d000000400b00043d000009930a000041000000000aab0436001f0000000a001d0000004005500039000000000252034f000000000502043b0000000402b00039000000800a0000390000000000a2043500000000020304330000008403b00039000000000023043500000000020804330000093002200197000000a403b0003900000000002304350000000002060433000000c403b00039000000a00600003900000000006304350000012408b000390000000063020434000000000038043500200000000b001d0000014402b00039000000000003004b00001a320000613d0000000008000019000000000a280019000000000b860019000000000b0b04330000000000ba04350000002008800039000000000038004b00001a2b0000413d0000005001100270000000000623001900000000000604350000001f03300039000009c80330019700000000060904330000002008000029000000e408800039000000c0093000390000000000980435000000000323001900000000260604340000000003630436000000000006004b00001a490000613d00000000080000190000000009380019000000000a820019000000000a0a04330000000000a904350000002008800039000000000068004b00001a420000413d0000092f02100197000000000136001900000000000104350000001f01600039000009c801100197000000000131001900000020080000290000000003810049000000840630008a00000000030704330000010407800039000000000067043500000000060304330000000001610436000000000006004b00001a650000613d00000000070000190000002003300039000000000803043300000000980804340000092f0880019700000000088104360000000009090433000000000098043500000040011000390000000107700039000000000067004b00001a5a0000413d0000092f03400197000000200600002900000064046000390000000000340435000000440360003900000000005304350000002403600039000013880400003900000000004304350000000003000414000000040020008c00001d000000c13d0000000301000367000000010300003100001d140000013d0000001b010000290000000501100270000000000100003f00000000010004150000001c011000690000000001000002000014970000013d0000099001000041000000000010043f0000093c01000041000024aa000104300000098f01000041000000000010043f0000093c01000041000024aa000104300000098101000041000000000010043f0000093c01000041000024aa00010430000000400200043d000000240320003900000000001304350000097c0100004100000000001204350000000401200039000000200300003900000000003104350000092b0020009c0000092b0200804100000040012002100000097d011001c7000024aa00010430000000400400043d002000000004001d0000096e020000410000000000240435000000040240003900000020030000390000000000320435000000240240003900001c0b0000013d00000016010000290000000501100270000000000100003f000000000100041500000015011000690000000001000002000000400100043d0000099102000041000000000021043500000004021000390000001c03000029000005240000013d000000000300001900001aae0000013d0000000103300039000000190030006c00000c580000813d000000800400043d0000000001040433000000000031004b00001cca0000a13d0000001b010000290000000001010433000000000031004b00001cca0000a13d0000000505300210000000200150003900000000044100190000000004040433001c00000004001d0000001a0450002900000000040404330000000054040434002000000005001d000000000004004b00001ac60000613d0000001c0500002900000080055000390000000005050433000000000054004b00001ce50000413d0000001c040000290000014004400039001f00000004001d00000000040404330000000004040433000000200500002900000000050504330000000005050433000000000054004b00001cd00000c13d000000000004004b00001aab0000613d0000001c0400002900000160044000390000000004040433001e00000004001d001d00000014001d000000000700001900001adf0000013d00000001077000390000001f0100002900000000010104330000000001010433000000000017004b00001aab0000813d0000001e010000290000000001010433000000000031004b00001cca0000a13d0000001d0100002900000000010104330000000014010434000009680040009c000000f80000213d000000200040008c000000f80000413d0000000005010433000009300050009c000000f80000213d000000000c410019000000000a1500190000000001ac0049000009680010009c000000f80000213d000000800010008c000000f80000413d000000400b00043d0000096000b0009c000000fd0000213d000000800eb000390000004000e0043f00000000f10a0434000009300010009c000000f80000213d0000000001a100190000001f041000390000000000c4004b00000000050000190000096b050080410000096b044001970000096b0dc001970000000009d4013f0000000000d4004b00000000040000190000096b040040410000096b0090009c000000000405c019000000000004004b000000f80000c13d0000000041010434000009300010009c000000fd0000213d0000001f05100039000009c8055001970000003f05500039000009c8055001970000000005e50019000009300050009c000000fd0000213d000000400050043f00000000001e043500000000054100190000000000c5004b000000f80000213d000000a005b00039000000000001004b00001b250000613d000000000900001900000000085900190000000006490019000000000606043300000000006804350000002009900039000000000019004b00001b1e0000413d00000000015100190000000000010435000000000eeb043600000000010f0433000009300010009c000000f80000213d0000000001a100190000001f041000390000000000c4004b00000000050000190000096b050080410000096b044001970000000006d4013f0000000000d4004b00000000040000190000096b040040410000096b0060009c000000000405c019000000000004004b000000f80000c13d000000001f0104340000093000f0009c000000fd0000213d0000001f04f00039000009c8044001970000003f04400039000009c805400197000000400400043d0000000005540019000000000045004b00000000090000390000000109004039000009300050009c000000fd0000213d0000000100900190000000fd0000c13d000000400050043f0000000005f4043600000000061f00190000000000c6004b000000f80000213d00000000000f004b00001b580000613d0000000009000019000000000659001900000000081900190000000008080433000000000086043500000020099000390000000000f9004b00001b510000413d0000000001f50019000000000001043500000000004e04350000004001a000390000000001010433000009300010009c000000f80000213d0000000001a100190000001f041000390000000000c4004b00000000050000190000096b050080410000096b044001970000000006d4013f0000000000d4004b00000000040000190000096b040040410000096b0060009c000000000405c019000000000004004b000000f80000c13d000000001d0104340000093000d0009c000000fd0000213d0000001f04d00039000009c8044001970000003f04400039000009c804400197000000400e00043d00000000044e00190000000000e4004b00000000050000390000000105004039000009300040009c000000fd0000213d0000000100500190000000fd0000c13d000000400040043f0000000004de043600000000051d00190000000000c5004b000000f80000213d00000000000d004b00001b8c0000613d0000000005000019000000000645001900000000081500190000000008080433000000000086043500000020055000390000000000d5004b00001b850000413d0000000001d4001900000000000104350000004001b000390000000000e104350000006001a0003900000000010104330000092b0010009c000000f80000213d0000006004b000390000000000140435000000200400002900000000040404330000000005040433000000000075004b00001cca0000a13d00000005057002100000000004450019000000200440003900000000040404330000092b0440019800001ad90000613d000000000014004b00001ad90000813d0000001c0200002900000180022000390000000002020433000000400300043d000000640530003900000000004504350000004404300039000000000014043500000024013000390000000000710435000009a1010000410000000000130435000000040130003900000000002104350000092b0030009c0000092b030080410000004001300210000009a2011001c7000024aa000104300000000a0000006b00001bbd0000613d0000000301000039000000000101041a001f00000001001d000009870010019800001c320000c13d0000000001000415000000010110006900000000010000020000000004000031000000200240006a0000000201000367000000230720008a000000a40000013d000000400300043d002000000003001d0000097b0100004100000000001304350000000401300039000000200200003900000000002104350000002402300039000000160100002900001c0b0000013d000000400200043d002000000002001d0000097b0100004100000000001204350000000401200039000000170200002924a81de20000040f000000200200002900000000012100490000092b0010009c0000092b0100804100000060011002100000092b0020009c0000092b020080410000004002200210000000000121019f000024aa00010430000000400200043d000000240320003900000020040000290000000000430435000009960300004100000000003204350000000403200039000000000013043500001a8f0000013d00000017010000290000000001010433000000400200043d000009860300004100000000003204350000092f01100197000000040320003900000000001304350000092b0020009c0000092b02008041000000400120021000000972011001c7000024aa000104300000001f0530018f0000092d06300198000000400200043d0000000004620019000008c40000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001bfd0000c13d000008c40000013d000000400300043d002000000003001d0000097b0100004100000000001304350000000401300039000000200200003900000000002104350000002402300039000000170100002924a81dad0000040f000000200200002900000000012100490000092b0010009c0000092b010080410000092b0020009c0000092b0200804100000060011002100000004002200210000000000121019f000024aa00010430000000400200043d0000004403200039000000000013043500000024012000390000000c030000290000000000310435000009820100004100000000001204350000000401200039000000190300002900000000003104350000092b0020009c0000092b02008041000000400120021000000942011001c7000024aa000104300000001f0530018f0000092d06300198000000400200043d0000000004620019000008c40000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001c2d0000c13d000008c40000013d0000000401000039000000000101041a001e00000001001d0000093501000041000000000010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000936011001c70000800b0200003924a824a30000040f000000010020019000001c500000613d0000001f0200002900000080022002700000092b02200197000000000301043b000000000423004b0000013f0000413d0000001f0100002900000932011001970000001e02000029000009320220019700001c6a0000613d000000000021004b00001c580000a13d000000400100043d000009890200004100000a9d0000013d000000000001042f000000240320003900000000001304350000099701000041000000000012043500000004012000390000001e0300002900001a8e0000013d0000001e05000029000000800650027000000000056400a900000000044500d9000000000064004b0000013f0000c13d000000000015001a0000013f0000413d0000000001150019000000800330021000000938033001970000000305000039000000000405041a0000098804400197000000000334019f000000000035041b000000000012004b00000000010240190000000a0020006c00001c790000813d000000400100043d00000024031000390000000a0400002900000000004304350000098e030000410000000000310435000000040310003900000000002304350000092b0010009c0000092b0100804100000040011002100000097d011001c7000024aa000104300000000a0210006c00001c910000813d0000000402000039000000000202041a00000080022002720000013f0000613d0000000a041000690000000003240019000000010330008a000000000043004b0000013f0000413d000000400400043d000000240540003900000000001504350000098d01000041000000000014043500000000012300d9000000040240003900000000001204350000092b0040009c0000092b0400804100000040014002100000097d011001c7000024aa0001043000000932012001970000000303000039000000000203041a0000098a02200197000000000112019f000000000013041b000000400100043d0000000a0200002900000000002104350000092b0010009c0000092b01008041000000400110021000000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f0000098b011001c70000800d0200003900000001030000390000098c0400004124a8249e0000040f000000010020019000001bbd0000c13d000000f80000013d000000000032043500001ce20000c13d0000000902000039000000000302041a000000000013004b00001cca0000a13d000000000020043f000009990110009a000000000101041a0000092f011001970000000002000411000000000012004b00001ce20000c13d000001450100008a0000001f0010006b0000013f0000213d0000001f010000290000014401100039000000160010002a0000013f0000413d0000001601100029000000150010002a0000013f0000413d00000015011000290000000002000031000000000012004b000007de0000613d000000400300043d000000240430003900000000002404350000099b0200004100001cda0000013d000009a401000041000000000010043f0000003201000039000000040010043f0000097201000041000024aa000104300000001c0200002900000180012000390000000001010433000000600220003900000000020204330000093002200197000000400300043d00000024043000390000000000240435000009a0020000410000000000230435000000040230003900000000001204350000092b0030009c0000092b0300804100000040013002100000097d011001c7000024aa00010430000000400100043d0000099a0200004100000a9d0000013d0000001c0100002900000180011000390000000001010433000000400200043d00000044032000390000000000430435000000240320003900000000005304350000099f0300004100000000003204350000000403200039000000000013043500001c210000013d000009bb01000041000000170300002900000000001304350000000401300039000000020200003900000000002104350000092b0030009c0000092b03008041000000400130021000000972011001c7000024aa00010430000009bc01000041000000170200002900000aa50000013d000000200400002900000000014100490000092b0010009c0000092b0100804100000060011002100000092b0040009c0000092b040080410000004004400210000000000141019f0000092b0030009c0000092b03008041000000c003300210000000000113019f24a8249e0000040f00000060031002700001092b0030019d0000092b033001970003000000010355000000010020019000001d780000613d000009c8043001980000001f0530018f000000200240002900001d1e0000613d000000000601034f0000002007000029000000006806043c0000000007870436000000000027004b00001d1a0000c13d000000000005004b00001d2b0000613d000000000141034f0000000304500210000000000502043300000000054501cf000000000545022f000000000101043b0000010004400089000000000141022f00000000014101cf000000000151019f00000000001204350000001f01300039000009c8011001970000002002100029000000000012004b00000000010000390000000101004039000009300020009c000000fd0000213d0000000100100190000000fd0000c13d000000400020043f000009680030009c000000f80000213d000000600030008c000000f80000413d00000020010000290000000001010433000000000001004b0000000004000039000000010400c039000000000041004b000000f80000c13d0000001f040000290000000004040433000009300040009c000000f80000213d000000200530002900000020034000290000001f04300039000000000054004b00000000060000190000096b060080410000096b044001970000096b07500197000000000874013f000000000074004b00000000040000190000096b040040410000096b0080009c000000000406c019000000000004004b000000f80000c13d0000000043030434000009300030009c000000fd0000213d0000001f06300039000009c8066001970000003f06600039000009c8066001970000000006260019000009300060009c000000fd0000213d000000400060043f00000000063204360000000007430019000000000057004b000000f80000213d000000000003004b00001d6e0000613d000000000500001900000000076500190000000008450019000000000808043300000000008704350000002005500039000000000035004b00001d670000413d00000000036300190000000000030435000000000001004b000014970000c13d000000400300043d002000000003001d00000994010000410000000000130435000000040130003900001bd50000013d0000001f0530018f0000092d06300198000000400200043d0000000004620019000008c40000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b00001d7f0000c13d000008c40000013d000009cb0010009c00001d890000813d000000a001100039000000400010043f000000000001042d000009a401000041000000000010043f0000004101000039000000040010043f0000097201000041000024aa000104300000001f02200039000009c8022001970000000001120019000000000021004b00000000020000390000000102004039000009300010009c00001d9b0000213d000000010020019000001d9b0000c13d000000400010043f000000000001042d000009a401000041000000000010043f0000004101000039000000040010043f0000097201000041000024aa00010430000000400100043d000009cc0010009c00001da70000813d000000e002100039000000400020043f000000000001042d000009a401000041000000000010043f0000004101000039000000040010043f0000097201000041000024aa0001043000000000430104340000000001320436000000000003004b00001db90000613d000000000200001900000000051200190000000006240019000000000606043300000000006504350000002002200039000000000032004b00001db20000413d000000000213001900000000000204350000001f02300039000009c8022001970000000001210019000000000001042d000000000301001900000000040104330000000001420436000000000004004b00001dcc0000613d0000000002000019000000200330003900000000050304330000092f0550019700000000015104360000000102200039000000000042004b00001dc50000413d000000000001042d000000400100043d000009cb0010009c00001ddc0000813d000000a002100039000000400020043f000000800210003900000000000204350000006002100039000000000002043500000040021000390000000000020435000000200210003900000000000204350000000000010435000000000001042d000009a401000041000000000010043f0000004101000039000000040010043f0000097201000041000024aa0001043000000020030000390000000004310436000000003202043400000000002404350000004001100039000000000002004b00001df10000613d000000000400001900000000051400190000000006430019000000000606043300000000006504350000002004400039000000000024004b00001dea0000413d000000000312001900000000000304350000001f02200039000009c8022001970000000001120019000000000001042d0015000000000002000b00000002001d001200000001001d000000400200043d000009cd01000041001500000002001d000000000012043500000970010000410000000000100443000000000100041200000004001004430000004001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f0000000100200190000023ea0000613d00000015020000290000000402200039000000000101043b0000008001100210000009ce0110019700000000001204350000097001000041000000000010044300000000010004120000000400100443000000e001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f0000000100200190000023ea0000613d000000150b000029000000000201043b00000000010004140000092f02200197000000040020008c00001e2d0000c13d0000000103000031000000200030008c0000002004000039000000000403401900001e580000013d0000092b00b0009c0000092b0300004100000000030b401900000040033002100000092b0010009c0000092b01008041000000c001100210000000000131019f00000972011001c724a824a30000040f000000150b00002900000060031002700000092b03300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b001900001e470000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b00001e430000c13d000000000006004b00001e540000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f00030000000103550000000100200190000024400000613d0000001f01400039000000600110018f0000000002b10019000000000012004b00000000010000390000000101004039000d00000002001d000009300020009c000023de0000213d0000000100100190000023de0000c13d0000000d01000029000000400010043f0000001f0030008c000023d60000a13d00000000010b0433000000000001004b0000000002000039000000010200c039000000000021004b000023d60000c13d000000000001004b000024320000c13d000000120100002900000000210104340000000003010433000000000003004b000024340000613d000300000002001d00000000010204330000000001010433000000000013004b000024360000c13d001100000003001d000009300030009c000023de0000213d000000110100002900000005021002100000003f012000390000095f011001970000000d01100029000009300010009c000023de0000213d000000400010043f00000011010000290000000d030000290000000001130436000c00000001001d0000001f0120018f000000000002004b00001e930000613d0000000c04000029000000000224001900000000030000310000000203300367000000003503043c0000000004540436000000000024004b00001e8f0000c13d000000000001004b0000097001000041000000000010044300000000010004120000000400100443000000a001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f0000000100200190000023ea0000613d000000000101043b000a00000001001d0000000003000019000000120100002900000000010104330000000002010433000000000032004b000023d80000a13d001400000003001d0010000500300218000000100110002900000020011000390000000009010433000000200190003900000000010104330000004002900039000000000302043300000060029000390000000004020433000000a0029000390000000005020433000000c0029000390000000006020433000000e002900039000000000702043300000080029000390000000008020433001500000009001d00000100029000390000000009020433000000400200043d000001000a20003900000000009a0435000000800920003900000000008904350000092f07700197000000e00820003900000000007804350000093006600197000000c0072000390000000000670435000000000005004b0000000005000039000000010500c039000000a00620003900000000005604350000093004400197000000600520003900000000004504350000092f0330019700000040042000390000000000340435000001000300003900000000033204360000092f011001970000000000130435000009780020009c000023de0000213d0000012001200039000000400010043f0000092b0030009c0000092b03008041000000400130021000000000020204330000092b0020009c0000092b020080410000006002200210000000000112019f00000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f000009ab011001c7000080100200003924a824a30000040f0000000100200190000023d60000613d00000015020000290000012002200039000000000202043300000020032000390000092b0030009c0000092b03008041000000400330021000000000020204330000092b0020009c0000092b020080410000006002200210000000000232019f000000000101043b001300000001001d00000000010004140000092b0010009c0000092b01008041000000c001100210000000000121019f000009ab011001c7000080100200003924a824a30000040f0000000100200190000023d60000613d000000000601043b000000150100002900000140011000390000000003010433000000400100043d0000002002100039000000200400003900000000004204350000004004100039000000000503043300000000005404350000006004100039000000000005004b000f00000006001d00001f260000613d0000000006000019000000200900008a0000002003300039000000000703043300000000870704340000092f0770019700000000077404360000000008080433000000000087043500000040044000390000000106600039000000000056004b00001f1a0000413d00001f270000013d000000200900008a0000000003140049000000200430008a00000000004104350000001f03300039000000000493016f0000000003140019000000000043004b00000000040000390000000104004039000009300030009c000023de0000213d0000000100400190000023de0000c13d000000400030043f0000092b0020009c0000092b02008041000000400220021000000000010104330000092b0010009c0000092b010080410000006001100210000000000121019f00000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f000009ab011001c7000080100200003924a824a30000040f0000000100200190000023d60000613d000000000101043b000e00000001001d000000150100002900000160011000390000000003010433000000400100043d000000200210003900000020040000390000000000420435000000000403043300000005054002100000000006510019000000400510003900000000004504350000006007600039000000000004004b00001f760000613d0000000006000019000000200d00008a00001f630000013d000000000978001900000000000904350000001f088000390000000008d8016f00000000077800190000000106600039000000000046004b00001f770000813d0000000008170049000000600880008a000000200550003900000000008504350000002003300039000000000803043300000000980804340000000007870436000000000008004b00001f5b0000613d000000000a000019000000000b7a0019000000000ca90019000000000c0c04330000000000cb0435000000200aa0003900000000008a004b00001f6e0000413d00001f5b0000013d000000200d00008a0000000003170049000000200430008a00000000004104350000001f033000390000000004d3016f0000000003140019000000000043004b00000000040000390000000104004039000009300030009c000023de0000213d0000000100400190000023de0000c13d000000400030043f0000092b0020009c0000092b02008041000000400220021000000000010104330000092b0010009c0000092b010080410000006001100210000000000121019f00000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f000009ab011001c7000080100200003924a824a30000040f0000000100200190000023d60000613d000000000201043b000000400100043d000000c0031000390000000000230435000000a0021000390000000e03000029000000000032043500000080021000390000000f03000029000000000032043500000060021000390000001303000029000000000032043500000040021000390000000a030000290000000000320435000000c002000039000000000221043600000000000204350000092e0010009c000023de0000213d000000e003100039000000400030043f0000092b0020009c0000092b02008041000000400220021000000000010104330000092b0010009c0000092b010080410000006001100210000000000121019f00000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f000009ab011001c7000080100200003924a824a30000040f0000000100200190000023d60000613d000000000101043b0000000d0800002900000000020804330000001403000029000000000032004b000023d80000a13d00000010040000290000000c0240002900000000001204350000000002080433000000000032004b000023d80000a13d000000150200002900000180022000390000000002020433000000000021004b0000001101000029000023eb0000c13d0000000103300039000000000013004b00001ea60000413d000000120200002900000060012000390000000001010433000000400220003900000000020204330000000b030000290000000006030433000000400900043d000009d20300004100000000003904350000000403900039000000600400003900000000004304350000000005080433000000640490003900000000005404350000008404900039000000000006004b0000000006000039000000010600c039000400000006001d000000000005004b00001ff30000613d00000000060000190000002008800039000000000708043300000000047404360000000106600039000000000056004b00001fed0000413d00000000033400490000002405900039000000000035043500000000030204330000000006340436000000000003004b000020010000613d00000000040000190000002002200039000000000502043300000000065604360000000104400039000000000034004b00001ffb0000413d001400000006001d001500000009001d0000004402900039000000000012043500000970010000410000000000100443000000000100041200000004001004430000002001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f0000000100200190000023ea0000613d000000000201043b00000000010004140000092f02200197000000040020008c0000001403000029000020200000c13d0000000103000031000000200030008c00000020040000390000000004034019000000150b0000290000204f0000013d000000150400002900000000034300490000092b0030009c0000092b0300804100000060033002100000092b0040009c0000092b040080410000004004400210000000000343019f0000092b0010009c0000092b01008041000000c001100210000000000131019f24a824a30000040f000000150b00002900000060031002700000092b03300197000000200030008c000000200400003900000000040340190000001f0640018f000000200740019000000000057b00190000203e0000613d000000000801034f00000000090b0019000000008a08043c0000000009a90436000000000059004b0000203a0000c13d000000000006004b0000204b0000613d000000000771034f0000000306600210000000000805043300000000086801cf000000000868022f000000000707043b0000010006600089000000000767022f00000000066701cf000000000686019f0000000000650435000100000003001f000300000001035500000001002001900000244c0000613d0000001f01400039000000600210018f0000000001b20019000000000021004b00000000020000390000000102004039000009300010009c000023de0000213d0000000100200190000023de0000c13d000000400010043f000000200030008c000023d60000413d00000000020b0433000100000002001d000000000002004b0000243e0000613d000a00040000002d0000000003000019000020800000013d0000001f04200039000000000494016f0000000002320019000000000002043500000060024000390000092b0020009c0000092b0200804100000060022002100000092b0010009c0000092b010080410000004001100210000000000112019f00000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f000009ab011001c70000800d020000390000000303000039000009dc0400004124a8249e0000040f0000000100200190000a00040000002d000023d60000613d00000013030000290000000103300039000000110030006c000023d50000813d000000120100002900000000010104330000000002010433000000000032004b000023d80000a13d001300000003001d0000000502300210000e00200020003d0000000e011000290000000001010433001400000001001d0000006001100039001500000001001d0000000001010433001000000001001d0000093001100197000f00000001001d0000000701100270000000000010043f0000001001000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000023d60000613d0000000f020000290000000102200210000000000101043b000000000101041a00000010030000290000007f03300190000020a90000613d000000ff0420018f00000000033400d9000000020030008c000023e40000c13d000000fe0220018f000000000121022f0000000301100190000020c90000613d000000030010008c000022c80000c13d000f00000001001d0000000a0000006b0000001303000029000020ce0000c13d000000150100002900000000010104330000093001100197000000400200043d00000000001204350000092b0020009c0000092b02008041000000400120021000000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f0000098b011001c70000800d020000390000000103000039000009e00400004124a8249e0000040f0000000100200190000a00000000001d0000207c0000c13d000023d60000013d000f00000001001d0000000a0000006b000c00600000003d0000001303000029000020fc0000613d0000000b010000290000000001010433000000000031004b000023d80000a13d0000000e020000290000000b01200029001000000001001d000000000101043300000020011000390000000001010433000c00000001001d0000093501000041000000000010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000936011001c70000800b0200003924a824a30000040f0000000100200190000023ea0000613d000000000101043b000000010110006c000023e40000413d0000000f02000029000000030020008c000020ef0000613d0000000a02000039000000000202041a0000092b02200197000000000021004b0000240e0000a13d0000000b0100002900000000010104330000001303000029000000000031004b000023d80000a13d000000100100002900000000010104330000000001010433000000000001004b000020fc0000613d0000001402000029000000800220003900000000001204350000001402000029000000c001200039000d00000001001d00000000010104330000093000100198001000200020003d000021220000613d000000100100002900000000010104330000092f01100197000000000010043f0000000f01000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000023d60000613d000000000101043b000000000101041a0000093004100198000023150000613d0000000f0000006b0000001303000029000021220000c13d000009300040009c000023e40000613d0000000d01000029000000000101043300000930051001970000000101400039000000000051004b0000001303000029000023650000c13d000000030100002900000000010104330000000002010433000000000032004b000023d80000a13d0000000e01100029000000000301043300000015010000290000000001010433000200000001001d00000014020000290000012001200039000500000001001d00000000010104330000000001010433000a00000001001d0000000001020433000e00000001001d0000014001200039000600000001001d00000000010104330000000001010433000900000001001d000700000003001d0000000001030433000800000001001d00000970010000410000000000100443000000000100041200000004001004430000004001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f0000000100200190000023ea0000613d000000000101043b0000000e03000029000000000131013f0000093000100198000023ee0000c13d0000000a01000039000000000101041a00000040021002700000ffff0220018f0000000903000029000000000032004b0000000a04000029000023f30000413d0000000002030019000000080020006c000023f60000c13d00000020011002700000092b01100197000000000041004b000024020000413d00000015010000290000000002010433000009300120019700000001041002100000007f02200190000021690000613d000000ff0340018f00000000022300d9000000020020008c000023e40000c13d000e00000004001d0000000701100270000000000010043f0000001001000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000023d60000613d0000000e02000029000000fe0220018f000000030320020f000009e303300167000000000101043b000000000401041a000000000334016f000000010220020f000000000223019f000000000021041b000009790100004100000000001004430000000001000410000000040010044300000000010004140000092b0010009c0000092b01008041000000c0011002100000097a011001c7000080020200003924a824a30000040f0000000100200190000023ea0000613d000000000101043b000000000001004b000023d60000613d000000400e00043d000009da01000041000000000d1e04360000000401e00039000000600200003900000000002104350000001403000029000000000103043300000930011001970000006402e000390000000000120435000000100100002900000000010104330000092f011001970000008402e000390000000000120435000000400130003900000000010104330000092f01100197000000a402e000390000000000120435000000150100002900000000010104330000093001100197000000c402e00039000000000012043500000080013000390000000001010433000000e402e000390000000000120435000000a0013000390000000001010433000000000001004b0000000001000039000000010100c0390000010402e0003900000000001204350000000d01000029000000000101043300000930011001970000012402e000390000000000120435000000e00130003900000000010104330000092f011001970000014402e000390000000000120435000001000130003900000000010104330000016402e000390000000000120435000000050100002900000000010104330000018402e00039000001a00300003900000000003204350000020402e00039000000003101043400000000001204350000022402e00039000000000001004b000021d70000613d000000000400001900000000052400190000000006430019000000000606043300000000006504350000002004400039000000000014004b000021d00000413d000000000321001900000000000304350000001f01100039000009c80310019700000006010000290000000001010433000001a404e00039000001c0053000390000000000540435000000000223001900000000030104330000000002320436000000000003004b000021f10000613d00000000040000190000002001100039000000000501043300000000650504340000092f0550019700000000055204360000000006060433000000000065043500000040022000390000000104400039000000000034004b000021e60000413d0000000001e20049000000640110008a000000140300002900000160033000390000000003030433000001c404e00039000000000014043500000000040304330000000000420435000000050140021000000000011200190000002001100039000000000004004b0000221d0000613d000000000500001900000000060200190000220a0000013d000000000817001900000000000804350000001f07700039000009c80770019700000000011700190000000105500039000000000045004b0000221d0000813d0000000007210049000000200770008a000000200660003900000000007604350000002003300039000000000703043300000000870704340000000001710436000000000007004b000022020000613d0000000009000019000000000a190019000000000b980019000000000b0b04330000000000ba04350000002009900039000000000079004b000022150000413d000022020000013d0000000002e10049000000040220008a00000014030000290000018003300039000900000003001d00000000030304330000002404e000390000000000240435000001e402e000390000000000320435000000070b00002900000000020b04330000000000210435000000050320021000000000033100190000002005300039000000000002004b0000224d0000613d000000000300001900000000040100190000223a0000013d000000000756001900000000000704350000001f06600039000009c80660019700000000055600190000000103300039000000000023004b0000224d0000813d0000000006150049000000200660008a00000020044000390000000000640435000000200bb0003900000000060b043300000000760604340000000005650436000000000006004b000022320000613d00000000080000190000000009580019000000000a870019000000000a0a04330000000000a904350000002008800039000000000068004b000022450000413d000022320000013d0000000001e50049000000040110008a0000004402e0003900000000001204350000000c0100002900000000020104330000000001250436000000000002004b0000225f0000613d00000000030000190000000c05000029000000200550003900000000040504330000092b0440019700000000014104360000000103300039000000000023004b000022580000413d00000000020004140000000003000410000000040030008c0000227b0000613d0000000001e100490000092b0010009c0000092b0100804100000060011002100000092b00e0009c0000092b0300004100000000030e40190000004003300210000000000131019f0000092b0020009c0000092b02008041000000c002200210000000000121019f000000000200041000140000000d001d000e0000000e001d24a8249e0000040f0000000e0e000029000000140d00002900000060031002700001092b0030019d00030000000103550000000100200190000023370000613d0000093000e0009c000023de0000213d0000004000e0043f0000096a00e0009c000023de0000213d0000004000d0043f00000000000e043500000000050000190000000204000039000800000005001d000c00000004001d000e0000000e001d00140000000d001d00000015010000290000000002010433000009300120019700000001041002100000007f02200190000022920000613d000000ff0340018f00000000022300d9000000020020008c000023e40000c13d000700000004001d0000000701100270000a00000001001d000000000010043f0000001001000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000023d60000613d000000000101043b000000000101041a000600000001001d0000000a01000029000000000010043f0000001001000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000023d60000613d0000000702000029000000fe0220018f0000000c0700002900000000032701cf000000030220020f000009e302200167000000060220017f000000000232019f000000000101043b000000000021041b00000008010000290000000100100190000022d70000613d0000000f0000006b000000200900008a000000140a0000290000000e02000029000024160000c13d0000000d0100002900000000010104330000093000100198000022e00000c13d000022fc0000013d00000015010000290000000001010433000000000200041400000930051001970000092b0020009c0000092b02008041000000c001200210000009ab011001c70000800d020000390000000203000039000009d30400004124a8249e0000040f00000001002001900000207c0000c13d000023d60000013d0000000d0100002900000000010104330000093000100198000000200900008a000000140a0000290000000e02000029000022fc0000613d0000000f0000006b000022fc0000c13d000000100100002900000000010104330000092f01100197000000000010043f0000000f01000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c7000080100200003924a824a30000040f0000000100200190000023d60000613d000000000101043b000000000201041a0000093003200197000009300030009c000023e40000613d000009d6022001970000000103300039000000000223019f000000000021041b000000200900008a000000140a0000290000000e020000290000000c070000290000000901000029000000000601043300000015010000290000000004010433000000400100043d000000200510003900000040030000390000000000350435000000000071043500000000020204330000004003100039000000000023043500000060031000390000093005400197000000000002004b000020630000613d0000000004000019000000000734001900000000084a0019000000000808043300000000008704350000002004400039000000000024004b0000230d0000413d000020630000013d0000097001000041000000000010044300000000010004120000000400100443000000c001000039000000240010044300000000010004140000092b0010009c0000092b01008041000000c00110021000000971011001c7000080050200003924a824a30000040f0000000100200190000023ea0000613d000000000101043b0000092f02100198000023710000613d00000010010000290000000001010433000000400a00043d000009a70300004100000000003a04350000092f011001970000000403a0003900000000001304350000000001000414000000040020008c000023730000c13d0000000103000031000000200030008c000000200400003900000000040340190000239e0000013d0000092b023001980000000405000029000000800d0000390000000304000039000000600e000039000022840000613d0000001f032000390000092c033001970000003f03300039000009db03300197000000400e00043d00000000033e00190000000000e3004b00000000040000390000000104004039000009300030009c000023de0000213d0000000100400190000023de0000c13d000000400030043f000000000d2e04360000092d0420019800000000034d0019000023550000613d000000000501034f00000000060d0019000000005705043c0000000006760436000000000036004b000023510000c13d0000001f02200190000023620000613d000000000141034f0000000302200210000000000403043300000000042401cf000000000424022f000000000101043b0000010002200089000000000121022f00000000012101cf000000000141019f000000000013043500000004050000290000000304000039000022840000013d0000001001000029000000000101043300000000020004140000092b0020009c0000092b020080410000092f06100197000000c001200210000009ab011001c70000800d020000390000000303000039000009d704000041000022d30000013d0000000004000019000021160000013d0000092b00a0009c0000092b0300004100000000030a401900000040033002100000092b0010009c0000092b01008041000000c001100210000000000131019f00000972011001c700090000000a001d24a824a30000040f000000090a00002900000060031002700000092b03300197000000200030008c00000020040000390000000004034019000000200640019000000000056a00190000238d0000613d000000000701034f00000000080a0019000000007907043c0000000008980436000000000058004b000023890000c13d0000001f074001900000239a0000613d000000000661034f0000000307700210000000000805043300000000087801cf000000000878022f000000000606043b0000010007700089000000000676022f00000000067601cf000000000686019f0000000000650435000100000003001f00030000000103550000000100200190000024260000613d0000001f01400039000000600210018f0000000001a20019000000000021004b00000000020000390000000102004039000009300010009c000023de0000213d0000000100200190000023de0000c13d000000400010043f000000200030008c000023d60000413d00000000030a0433000009300030009c000023d60000213d000023e40000613d000000100100002900000000010104330000092f061001970000000d01000029000000000101043300000930051001970000000101300039000000000051004b000023cc0000c13d000000000060043f0000000f01000039000000200010043f00000000010004140000092b0010009c0000092b01008041000000c00110021000000962011001c70000801002000039000900000003001d24a824a30000040f00000009040000290000000100200190000023d60000613d000000000101043b000000000201041a000009d602200197000000000242019f000000000021041b000021160000013d00000000010004140000092b0010009c0000092b01008041000000c001100210000009ab011001c70000800d020000390000000303000039000009d504000041000022d30000013d000000000001042d0000000001000019000024aa00010430000009a401000041000000000010043f0000003201000039000000040010043f0000097201000041000024aa00010430000009a401000041000000000010043f0000004101000039000000040010043f0000097201000041000024aa00010430000009a401000041000000000010043f0000001101000039000000040010043f0000097201000041000024aa00010430000000000001042f000000400100043d000009d102000041000024100000013d000000400100043d000009d80200004100000000002104350000093002300197000023fb0000013d000000400100043d000009df02000041000023f80000013d000000400100043d000009d902000041000000000021043500000002020000290000093002200197000000040310003900000000002304350000092b0010009c0000092b01008041000000400110021000000972011001c7000024aa00010430000000400200043d00000024032000390000000000430435000009de030000410000000000320435000000040320003900000000001304350000092b0020009c0000092b0200804100000040012002100000097d011001c7000024aa00010430000000400100043d000009d40200004100000000002104350000092b0010009c0000092b0100804100000040011002100000093c011001c7000024aa00010430000000400300043d001500000003001d000009dd010000410000000000130435000000040130003924a81de20000040f000000150200002900000000012100490000092b0010009c0000092b0100804100000060011002100000092b0020009c0000092b020080410000004002200210000000000121019f000024aa000104300000001f0530018f0000092d06300198000000400200043d0000000004620019000024570000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b0000242d0000c13d000024570000013d000009cf01000041000024370000013d000009e201000041000024370000013d000009d0010000410000000d0200002900000000001204350000092b0020009c0000092b0200804100000040012002100000093c011001c7000024aa00010430000009e102000041000024100000013d0000001f0530018f0000092d06300198000000400200043d0000000004620019000024570000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000024470000c13d000024570000013d0000001f0530018f0000092d06300198000000400200043d0000000004620019000024570000613d000000000701034f0000000008020019000000007907043c0000000008980436000000000048004b000024530000c13d000000000005004b000024640000613d000000000161034f0000000305500210000000000604043300000000065601cf000000000656022f000000000101043b0000010005500089000000000151022f00000000015101cf000000000161019f000000000014043500000060013002100000092b0020009c0000092b020080410000004002200210000000000112019f000024aa00010430000000000001042f0000092b0010009c0000092b0100804100000040011002100000092b0020009c0000092b020080410000006002200210000000000112019f00000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f000009ab011001c7000080100200003924a824a30000040f00000001002001900000247e0000613d000000000101043b000000000001042d0000000001000019000024aa0001043000000000050100190000000000200443000000050030008c0000248e0000413d000000040100003900000000020000190000000506200210000000000664001900000005066002700000000006060031000000000161043a0000000102200039000000000031004b000024860000413d0000092b0030009c0000092b03008041000000600130021000000000020004140000092b0020009c0000092b02008041000000c002200210000000000112019f000009e4011001c7000000000205001924a824a30000040f00000001002001900000249d0000613d000000000101043b000000000001042d000000000001042f000024a1002104210000000102000039000000000001042d0000000002000019000000000001042d000024a6002104230000000102000039000000000001042d0000000002000019000000000001042d000024a800000432000024a90001042e000024aa00010430000000000000000000000000000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000000000000000000000000001ffffffe000000000000000000000000000000000000000000000000000000000ffffffe0000000000000000000000000000000000000000000000000ffffffffffffff1f000000000000000000000000ffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffff9f00000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffff5f796b89b91644bc98cd93958e4c9038275d622183e25ac5af08cc6b5d9553913202000002000000000000000000000000000000040000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000ffffffffffffffffffffff0000000000000000000000000000000000000000009a8a0592ac89c5ad3bc6df8224c17b485976f597df104ee20d0df415241f670b4120fccd000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000df85440e000000000000000000000000000000000000000000000000000000008acd72527118c8324937b1a42e02cd246697c3b633f1742f3cae11de233722b3000000020000000000000000000000000000028000000100000000000000000043616e6e6f7420736574206f776e657220746f207a65726f000000000000000008c379a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000000000000000000000081ff704700000000000000000000000000000000000000000000000000000000afcb95d600000000000000000000000000000000000000000000000000000000c92b283100000000000000000000000000000000000000000000000000000000c92b283200000000000000000000000000000000000000000000000000000000f077b59200000000000000000000000000000000000000000000000000000000f2fde38b00000000000000000000000000000000000000000000000000000000afcb95d700000000000000000000000000000000000000000000000000000000b1dc65a400000000000000000000000000000000000000000000000000000000b6113fce00000000000000000000000000000000000000000000000000000000873504d600000000000000000000000000000000000000000000000000000000873504d7000000000000000000000000000000000000000000000000000000008926c4ee000000000000000000000000000000000000000000000000000000008da5cb5b0000000000000000000000000000000000000000000000000000000081ff704800000000000000000000000000000000000000000000000000000000856c824700000000000000000000000000000000000000000000000000000000599f643000000000000000000000000000000000000000000000000000000000704b6c0100000000000000000000000000000000000000000000000000000000704b6c02000000000000000000000000000000000000000000000000000000007437ff9f0000000000000000000000000000000000000000000000000000000079ba509700000000000000000000000000000000000000000000000000000000599f643100000000000000000000000000000000000000000000000000000000666cab8d00000000000000000000000000000000000000000000000000000000181f5a7600000000000000000000000000000000000000000000000000000000181f5a77000000000000000000000000000000000000000000000000000000001ef3817400000000000000000000000000000000000000000000000000000000546719cd0000000000000000000000000000000000000000000000000000000006285c6900000000000000000000000000000000000000000000000000000000142a98fc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0000000000000000000000000000000000000000000000000ffffffffffffff7f209699368efae3c2ab13a6e9d9f9aceb6c5aebfb5ffd7bd0a9ff6281a30b57390200000000000000000000000000000000000040000000000000000000000000f6cd5620000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000e00000000000000000ffffffffffffffffffffff00ffffffff0000000000000000000000000000000002000000000000000000000000000000000000600000000000000000000000009ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c197fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff371a732800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffdf8000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffbf000000000000000000000000fffffffffffffffffffffffffffffffffffffbff8d666f6000000000000000000000000000000000000000000000000000000000bbe4f6db00000000000000000000000000000000000000000000000000000000310ab089e4439a4c15d089f94afb7896ff553aecb10793d0ab882de59d99a32e0200000200000000000000000000000000000044000000000000000000000000000000000000000000000000000000000000002400000000000000000000000001ffc9a7000000000000000000000000000000000000000000000000000000000000000000007530000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000000000000000000000000000aff2afbf0000000000000000000000000000000000000000000000000000000070a0823100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000fffffffffffffedf1806aa1896bbf26568e884a7374b41e002500962caba6a15023a8d90e8508b830200000200000000000000000000000000000024000000000000000000000000e1cd55090000000000000000000000000000000000000000000000000000000078ef8024000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000044000000000000000000000000000000000000000000000000000000000000000000000000fffffffffffffeff3907753700000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffff3f37c3be2900000000000000000000000000000000000000000000000000000000a966e21f00000000000000000000000000000000000000000000000000000000d02641a00000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000de0b6b3a76400009a655f7b000000000000000000000000000000000000000000000000000000000000000000000000000000ff0000000000000000000000000000000000000000ffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9725942a00000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffff0000000000000000000000000000000002000000000000000000000000000000000000200000000000000000000000001871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a15279c0800000000000000000000000000000000000000000000000000000000f94ebcd100000000000000000000000000000000000000000000000000000000afa32a2c000000000000000000000000000000000000000000000000000000000c3b563c00000000000000000000000000000000000000000000000000000000ae9b4ce90000000000000000000000000000000000000000000000000000000085572ffb000000000000000000000000000000000000000000000000000000003cf97983000000000000000000000000000000000000000000000000000000000a8d6e8c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000fffffffffffffe5f93df584c000000000000000000000000000000000000000000000000000000000f01ce8500000000000000000000000000000000000000000000000000000000b04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a6291eabfe8e493f369f48e58fdf2609ff8809506ce57440a6f25fddc25308a3851da0f08e8000000000000000000000000000000000000000000000000000000008e1192e1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060000000800000000000000000000000000000000000000000000000000000002000000080000000000000000083e3f564000000000000000000000000000000000000000000000000000000009c6db58d0000000000000000000000000000000000000000000000000000000085d2e5bf00000000000000000000000000000000000000000000000000000000ef0c6352000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000084000000000000000000000000209699368efae3c2ab13a6e9d9f9aceb6c5aebfb5ffd7bd0a9ff6281a30b573a4e487b7100000000000000000000000000000000000000000000000000000000cbf3cbeaed4ac1d605ed30f4af06c35acaeff2379db7f6146c9cceee83d58782fc23abf7ddbd3c02b1420dafa2355c56c1a06fbb8723862ac14d6bd74177361a856c82470000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000004f6e6c792063616c6c61626c65206279206f776e65720000000000000000000000000000000000000000000000000000000000640000008000000000000000000200000000000000000000000000000000000000000000000000000000000000ed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae127843616e6e6f74207472616e7366657220746f2073656c660000000000000000004d7573742062652070726f706f736564206f776e6572000000000000000000008be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000400000080000000000000000002000000000000000000000000000000000000200000008000000000000000008fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af00000000000000000000000000000000000000000000ffff0000000000000000000000000000000000000000000000000000000000000000ffffffff000000000000ffffffffffffffffffffffffffffffffffffffff00000000000000000000ffff00000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000001800000000000000000000000007879e20bb60a503429de4a2c912b5904f08a39f2af054c10fb46434b5d611260367f56a200000000000000000000000000000000000000000000000000000000d6c62c9b00000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000100000000000000000000000000000000000000000000000000000000000042cbb15ccdc3cad6266b0e7a08c0454b23bf29dc2df74b6f3c209e9336465bd1ffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e058579befe0000000000000000000000000000000000000000000000000000000045564d3245564d4f666652616d7020312e352e300000000000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000e0000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff000000000000000000000000000000000000000000000000ffffffffffffff60000000000000000000000000000000000000000000000000ffffffffffffff202cbc26bb000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffff0000000000000000000000000000000053ad11d80000000000000000000000000000000000000000000000000000000057e0e083000000000000000000000000000000000000000000000000000000007185cf6b000000000000000000000000000000000000000000000000000000003204887500000000000000000000000000000000000000000000000000000000e3dd0bec917c965a133ddb2c84874725ee1e2fd8d763c19efa36d6a11cd82b1f6358b0d000000000000000000000000000000000000000000000000000000000e44a20935573a783dd0d5991c92d7b6a0eb3173566530364db3ec10e9a990b5dffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000d32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf412371279ec8a000000000000000000000000000000000000000000000000000000008808f8e700000000000000000000000000000000000000000000000000000000b6113fce0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003ffffffe0d4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65cf19edfd000000000000000000000000000000000000000000000000000000008693378900000000000000000000000000000000000000000000000000000000099d3f720000000000000000000000000000000000000000000000000000000067d9ba0f63d427c482c2736300e6d5a34c6691dbcdea8ad35828a1f1ba47e872ea7568010000000000000000000000000000000000000000000000000000000000bf199700000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff020000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009168ba79fd88ade4089995c40afe9d661c5d5ec307701f4a5dc06541c12b1303")

type GetAllRateLimitTokens struct {
	SourceTokens []common.Address
	DestTokens   []common.Address
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
	case _EVM2EVMOffRamp.abi.Events["AlreadyAttempted"].ID:
		return _EVM2EVMOffRamp.ParseAlreadyAttempted(log)
	case _EVM2EVMOffRamp.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMOffRamp.ParseConfigChanged(log)
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
	case _EVM2EVMOffRamp.abi.Events["SkippedAlreadyExecutedMessage"].ID:
		return _EVM2EVMOffRamp.ParseSkippedAlreadyExecutedMessage(log)
	case _EVM2EVMOffRamp.abi.Events["SkippedIncorrectNonce"].ID:
		return _EVM2EVMOffRamp.ParseSkippedIncorrectNonce(log)
	case _EVM2EVMOffRamp.abi.Events["SkippedSenderWithPreviousRampMessageInflight"].ID:
		return _EVM2EVMOffRamp.ParseSkippedSenderWithPreviousRampMessageInflight(log)
	case _EVM2EVMOffRamp.abi.Events["TokenAggregateRateLimitAdded"].ID:
		return _EVM2EVMOffRamp.ParseTokenAggregateRateLimitAdded(log)
	case _EVM2EVMOffRamp.abi.Events["TokenAggregateRateLimitRemoved"].ID:
		return _EVM2EVMOffRamp.ParseTokenAggregateRateLimitRemoved(log)
	case _EVM2EVMOffRamp.abi.Events["TokensConsumed"].ID:
		return _EVM2EVMOffRamp.ParseTokensConsumed(log)
	case _EVM2EVMOffRamp.abi.Events["Transmitted"].ID:
		return _EVM2EVMOffRamp.ParseTransmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMOffRampAdminSet) Topic() common.Hash {
	return common.HexToHash("0x8fe72c3e0020beb3234e76ae6676fa576fbfcae600af1c4fea44784cf0db329c")
}

func (EVM2EVMOffRampAlreadyAttempted) Topic() common.Hash {
	return common.HexToHash("0x67d9ba0f63d427c482c2736300e6d5a34c6691dbcdea8ad35828a1f1ba47e872")
}

func (EVM2EVMOffRampConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (EVM2EVMOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x7879e20bb60a503429de4a2c912b5904f08a39f2af054c10fb46434b5d611260")
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

func (EVM2EVMOffRampSkippedAlreadyExecutedMessage) Topic() common.Hash {
	return common.HexToHash("0xe3dd0bec917c965a133ddb2c84874725ee1e2fd8d763c19efa36d6a11cd82b1f")
}

func (EVM2EVMOffRampSkippedIncorrectNonce) Topic() common.Hash {
	return common.HexToHash("0xd32ddb11d71e3d63411d37b09f9a8b28664f1cb1338bfd1413c173b0ebf41237")
}

func (EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight) Topic() common.Hash {
	return common.HexToHash("0xe44a20935573a783dd0d5991c92d7b6a0eb3173566530364db3ec10e9a990b5d")
}

func (EVM2EVMOffRampTokenAggregateRateLimitAdded) Topic() common.Hash {
	return common.HexToHash("0xfc23abf7ddbd3c02b1420dafa2355c56c1a06fbb8723862ac14d6bd74177361a")
}

func (EVM2EVMOffRampTokenAggregateRateLimitRemoved) Topic() common.Hash {
	return common.HexToHash("0xcbf3cbeaed4ac1d605ed30f4af06c35acaeff2379db7f6146c9cceee83d58782")
}

func (EVM2EVMOffRampTokensConsumed) Topic() common.Hash {
	return common.HexToHash("0x1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a")
}

func (EVM2EVMOffRampTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (_EVM2EVMOffRamp *EVM2EVMOffRamp) Address() common.Address {
	return _EVM2EVMOffRamp.address
}

type CustomTransaction struct {
	*types.Transaction
	CustomHash common.Hash
}

func (tx *CustomTransaction) Hash() common.Hash {
	return tx.CustomHash
}

func ConvertToTransaction(resp zktypes.TransactionResponse) *CustomTransaction {
	dtx := &types.DynamicFeeTx{
		ChainID:   resp.ChainID.ToInt(),
		Nonce:     uint64(resp.Nonce),
		GasTipCap: resp.MaxPriorityFeePerGas.ToInt(),
		GasFeeCap: resp.MaxFeePerGas.ToInt(),
		To:        &resp.To,
		Value:     resp.Value.ToInt(),
		Data:      resp.Data,
		Gas:       uint64(resp.Gas),
	}

	tx := types.NewTx(dtx)
	customTransaction := CustomTransaction{Transaction: tx, CustomHash: resp.Hash}
	return &customTransaction
}

func DeployZkSyncEVM2EVMOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, params ...interface{}) (common.Address, *CustomTransaction, *EVM2EVMOffRamp, error) {
	client, ok := backend.(*ethclient.Client)
	if !ok {
		return common.Address{}, nil, nil, errors.New("backend is not an ethclient")
	}
	
	zksyncClient := zkSyncClient.NewClient(client.Client())
	
	wallet := auth.Context.Value("wallet").(*zkSyncAccounts.Wallet)
	
	
	decodedBytes := common.FromHex(EVM2EVMOffRampZkBin)
	
	EVM2EVMOffRampAbi, err := EVM2EVMOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	constructor, err := EVM2EVMOffRampAbi.Pack("", params...)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	hash, err := wallet.DeployWithCreate(nil, zkSyncAccounts.CreateTransaction{
		Bytecode: decodedBytes,
		Calldata: constructor,
	})
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	
	receipt, err := zksyncClient.WaitMined(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	tx, _, err := zksyncClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	

	ethTx := ConvertToTransaction(*tx)
	address := receipt.ContractAddress

	parsed, err := EVM2EVMOffRampMetaData.GetAbi()
	contractBind := bind.NewBoundContract(address, *parsed, backend, backend, backend)

	contractReturn := &EVM2EVMOffRamp{address: address, abi: *parsed, EVM2EVMOffRampCaller: EVM2EVMOffRampCaller{contract: contractBind}, EVM2EVMOffRampTransactor: EVM2EVMOffRampTransactor{contract: contractBind}, EVM2EVMOffRampFilterer: EVM2EVMOffRampFilterer{contract: contractBind}}

	return address, ethTx, contractReturn, err
}

type EVM2EVMOffRampInterface interface {
	CcipReceive(opts *bind.CallOpts, arg0 ClientAny2EVMMessage) error

	CurrentRateLimiterState(opts *bind.CallOpts) (RateLimiterTokenBucket, error)

	GetAllRateLimitTokens(opts *bind.CallOpts) (GetAllRateLimitTokens,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (EVM2EVMOffRampDynamicConfig, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetSenderNonce(opts *bind.CallOpts, sender common.Address) (uint64, error)

	GetStaticConfig(opts *bind.CallOpts) (EVM2EVMOffRampStaticConfig, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetTransmitters(opts *bind.CallOpts) ([]common.Address, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalEVM2EVMMessage, offchainTokenData [][]byte, tokenGasOverrides []uint32) (*types.Transaction, error)

	ManuallyExecute(opts *bind.TransactOpts, report InternalExecutionReport, gasLimitOverrides []EVM2EVMOffRampGasLimitOverride) (*types.Transaction, error)

	SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	SetOCR2Config(opts *bind.TransactOpts, signers []common.Address, transmitters []common.Address, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config RateLimiterConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, arg4 [32]byte) (*types.Transaction, error)

	UpdateRateLimitTokens(opts *bind.TransactOpts, removes []EVM2EVMOffRampRateLimitToken, adds []EVM2EVMOffRampRateLimitToken) (*types.Transaction, error)

	FilterAdminSet(opts *bind.FilterOpts) (*EVM2EVMOffRampAdminSetIterator, error)

	WatchAdminSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampAdminSet) (event.Subscription, error)

	ParseAdminSet(log types.Log) (*EVM2EVMOffRampAdminSet, error)

	FilterAlreadyAttempted(opts *bind.FilterOpts) (*EVM2EVMOffRampAlreadyAttemptedIterator, error)

	WatchAlreadyAttempted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampAlreadyAttempted) (event.Subscription, error)

	ParseAlreadyAttempted(log types.Log) (*EVM2EVMOffRampAlreadyAttempted, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMOffRampConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMOffRampConfigChanged, error)

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

	FilterSkippedAlreadyExecutedMessage(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMOffRampSkippedAlreadyExecutedMessageIterator, error)

	WatchSkippedAlreadyExecutedMessage(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampSkippedAlreadyExecutedMessage, sequenceNumber []uint64) (event.Subscription, error)

	ParseSkippedAlreadyExecutedMessage(log types.Log) (*EVM2EVMOffRampSkippedAlreadyExecutedMessage, error)

	FilterSkippedIncorrectNonce(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampSkippedIncorrectNonceIterator, error)

	WatchSkippedIncorrectNonce(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampSkippedIncorrectNonce, nonce []uint64, sender []common.Address) (event.Subscription, error)

	ParseSkippedIncorrectNonce(log types.Log) (*EVM2EVMOffRampSkippedIncorrectNonce, error)

	FilterSkippedSenderWithPreviousRampMessageInflight(opts *bind.FilterOpts, nonce []uint64, sender []common.Address) (*EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflightIterator, error)

	WatchSkippedSenderWithPreviousRampMessageInflight(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight, nonce []uint64, sender []common.Address) (event.Subscription, error)

	ParseSkippedSenderWithPreviousRampMessageInflight(log types.Log) (*EVM2EVMOffRampSkippedSenderWithPreviousRampMessageInflight, error)

	FilterTokenAggregateRateLimitAdded(opts *bind.FilterOpts) (*EVM2EVMOffRampTokenAggregateRateLimitAddedIterator, error)

	WatchTokenAggregateRateLimitAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTokenAggregateRateLimitAdded) (event.Subscription, error)

	ParseTokenAggregateRateLimitAdded(log types.Log) (*EVM2EVMOffRampTokenAggregateRateLimitAdded, error)

	FilterTokenAggregateRateLimitRemoved(opts *bind.FilterOpts) (*EVM2EVMOffRampTokenAggregateRateLimitRemovedIterator, error)

	WatchTokenAggregateRateLimitRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTokenAggregateRateLimitRemoved) (event.Subscription, error)

	ParseTokenAggregateRateLimitRemoved(log types.Log) (*EVM2EVMOffRampTokenAggregateRateLimitRemoved, error)

	FilterTokensConsumed(opts *bind.FilterOpts) (*EVM2EVMOffRampTokensConsumedIterator, error)

	WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTokensConsumed) (event.Subscription, error)

	ParseTokensConsumed(log types.Log) (*EVM2EVMOffRampTokensConsumed, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMOffRampTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMOffRampTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMOffRampTransmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
