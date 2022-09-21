// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package any_2_evm_toll_offramp_helper

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
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated"
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
)

type AggregateRateLimiterInterfaceRateLimiterConfig struct {
	Rate     *big.Int
	Capacity *big.Int
}

type AggregateRateLimiterInterfaceTokenBucket struct {
	Rate        *big.Int
	Capacity    *big.Int
	Tokens      *big.Int
	LastUpdated *big.Int
}

type BaseOffRampInterfaceOffRampConfig struct {
	OnRampAddress                           common.Address
	PermissionLessExecutionThresholdSeconds uint32
	ExecutionDelaySeconds                   uint64
	MaxDataSize                             uint64
	MaxTokensLength                         uint64
}

type CCIPAny2EVMMessageFromSender struct {
	SourceChainId *big.Int
	Sender        []byte
	Receiver      common.Address
	Data          []byte
	DestTokens    []common.Address
	DestPools     []common.Address
	Amounts       []*big.Int
	GasLimit      *big.Int
}

type CCIPEVM2EVMTollMessage struct {
	SourceChainId  *big.Int
	SequenceNumber uint64
	Sender         common.Address
	Receiver       common.Address
	Data           []byte
	Tokens         []common.Address
	Amounts        []*big.Int
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	GasLimit       *big.Int
}

type CCIPExecutionReport struct {
	SequenceNumbers          []uint64
	TokenPerFeeCoinAddresses []common.Address
	TokenPerFeeCoin          []*big.Int
	EncodedMessages          [][]byte
	InnerProofs              [][32]byte
	InnerProofFlagBits       *big.Int
	OuterProofs              [][32]byte
	OuterProofFlagBits       *big.Int
}

var EVM2EVMTollOffRampHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"offRampConfig\",\"type\":\"tuple\"},{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"},{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"sourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"pools\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"tokenLimitsAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAFNSignal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadHealthConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"ExecutionError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"IncorrectNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"}],\"name\":\"InsufficientFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"}],\"name\":\"InvalidSourceChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenPoolConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManualExecutionNotYetEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"}],\"name\":\"MessageTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCoin\",\"type\":\"address\"}],\"name\":\"MissingFeeCoinPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoMessagesToExecute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PriceNotFoundForToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReceiverError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefillRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RootNotRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAndAmountMisMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenPoolMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensAndPriceLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"UnsupportedNumberOfTokens\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueExceedsAllowedThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"oldAFN\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractAFNInterface\",\"name\":\"newAFN\",\"type\":\"address\"}],\"name\":\"AFNSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"ExecutionStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"OffRampConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"OffRampRouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"TokenPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensRemovedFromBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateCurrentTokenBucketState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ccipReceive\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64[]\",\"name\":\"sequenceNumbers\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenPerFeeCoinAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenPerFeeCoin\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"encodedMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"innerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"innerProofFlagBits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"outerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"outerProofFlagBits\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.ExecutionReport\",\"name\":\"report\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"manualExecution\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"destTokens\",\"type\":\"address[]\"},{\"internalType\":\"contractPoolInterface[]\",\"name\":\"destPools\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.Any2EVMMessageFromSender\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"executeSingleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeTaken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAFN\",\"outputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlobVerifier\",\"outputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getDestinationToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDestinationTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getExecutionState\",\"outputs\":[{\"internalType\":\"enumCCIP.MessageExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"sourceToken\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"contractPoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getPricesForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenLimitAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAFNHealthy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"merkleGasShare\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"internalType\":\"structCCIP.EVM2EVMTollMessage\",\"name\":\"message\",\"type\":\"tuple\"}],\"name\":\"overheadGasToll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractPoolInterface\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"removePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"executableMessages\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAFNInterface\",\"name\":\"afn\",\"type\":\"address\"}],\"name\":\"setAFN\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBlobVerifierInterface\",\"name\":\"blobVerifier\",\"type\":\"address\"}],\"name\":\"setBlobVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"executionDelaySeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxDataSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTokensLength\",\"type\":\"uint64\"}],\"internalType\":\"structBaseOffRampInterface.OffRampConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"internalType\":\"structAggregateRateLimiterInterface.RateLimiterConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"setRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAny2EVMOffRampRouterInterface\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setTokenLimitAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620061a0380380620061a08339810160408190526200003491620006d4565b8888888888888888886001898989898989898989818185858833806000806000806101000a81548160ff02191690831515021790555060006001600160a01b0316826001600160a01b031603620000d25760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0380851661010002610100600160a81b0319909216919091179091558116156200010c576200010c816200041d565b5050506001600160a01b0381166200013757604051630958ef9b60e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b03929092169190911790558051825114620001795760405162d8548360e71b815260040160405180910390fd5b81516200018e906004906020850190620004ce565b5060005b825181101562000259576000828281518110620001b357620001b36200082e565b602002602001015190506040518060400160405280826001600160a01b03168152602001836001600160601b031681525060036000868581518110620001fd57620001fd6200082e565b6020908102919091018101516001600160a01b039081168352828201939093526040909101600020835193909101516001600160601b0316600160a01b029290911691909117905550620002518162000844565b905062000192565b5050600580546001600160a01b0319166001600160a01b03938416179055506040805160808101825283518082526020948501805195830186905251928201839052426060909201829052600855600993909355600a55600b91909155875116620002d7576040516342bcdf7f60e11b815260040160405180910390fd5b88608081815250508760a0818152505086600e60008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160000160146101000a81548163ffffffff021916908363ffffffff16021790555060408201518160000160186101000a8154816001600160401b0302191690836001600160401b0316021790555060608201518160010160006101000a8154816001600160401b0302191690836001600160401b0316021790555060808201518160010160086101000a8154816001600160401b0302191690836001600160401b0316021790555090505085600d60006101000a8154816001600160a01b0302191690836001600160a01b0316021790555050505050505050505080151560c081151581525050505050505050505050505050505050505050506200086c565b336001600160a01b03821603620004775760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000c9565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b82805482825590600052602060002090810192821562000526579160200282015b828111156200052657825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620004ef565b506200053492915062000538565b5090565b5b8082111562000534576000815560010162000539565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b03811182821017156200058a576200058a6200054f565b60405290565b6001600160a01b0381168114620005a657600080fd5b50565b8051620005b68162000590565b919050565b80516001600160401b0381168114620005b657600080fd5b600082601f830112620005e557600080fd5b815160206001600160401b03808311156200060457620006046200054f565b8260051b604051601f19603f830116810181811084821117156200062c576200062c6200054f565b6040529384528581018301938381019250878511156200064b57600080fd5b83870191505b8482101562000677578151620006678162000590565b8352918301919083019062000651565b979650505050505050565b6000604082840312156200069557600080fd5b604080519081016001600160401b0381118282101715620006ba57620006ba6200054f565b604052825181526020928301519281019290925250919050565b6000806000806000806000806000898b036101c0811215620006f557600080fd5b8a51995060208b0151985060a0603f19820112156200071357600080fd5b506200071e62000565565b60408b01516200072e8162000590565b815260608b015163ffffffff811681146200074857600080fd5b60208201526200075b60808c01620005bb565b60408201526200076e60a08c01620005bb565b60608201526200078160c08c01620005bb565b608082015296506200079660e08b01620005a9565b9550620007a76101008b01620005a9565b6101208b01519095506001600160401b0380821115620007c657600080fd5b620007d48d838e01620005d3565b95506101408c0151915080821115620007ec57600080fd5b50620007fb8c828d01620005d3565b9350506200080e8b6101608c0162000682565b91506200081f6101a08b01620005a9565b90509295985092959850929598565b634e487b7160e01b600052603260045260246000fd5b6000600182016200086557634e487b7160e01b600052601160045260246000fd5b5060010190565b60805160a05160c0516158e1620008bf60003960006115c701526000818161030c01528181611e49015261233b0152600081816102e901528181611e2401528181612316015261366c01526158e16000f3fe608060405234801561001057600080fd5b50600436106102de5760003560e01c80638456cb5911610186578063b4069b31116100e3578063c3f909d411610097578063e3d0e71211610071578063e3d0e712146107a5578063eb511dd4146107b8578063f2fde38b146107cb57600080fd5b8063c3f909d414610698578063c5a1d7f01461078a578063c9029f6a1461079257600080fd5b8063bbe4f6db116100c8578063bbe4f6db14610646578063be9b03f114610672578063c0d786551461068557600080fd5b8063b4069b3114610620578063b57671661461063357600080fd5b8063a639d1c01161013a578063afcb95d71161011f578063afcb95d7146105dc578063b0f479a1146105fc578063b1dc65a41461060d57600080fd5b8063a639d1c0146105a9578063a8b640c1146105bc57600080fd5b80638da5cb5b1161016b5780638da5cb5b1461054557806390c2339b1461055b578063918725431461059657600080fd5b80638456cb591461053557806389c065681461053d57600080fd5b80634741062e1161023f578063681fba16116101f357806379ba5097116101cd57806379ba5097146104f557806381411834146104fd57806381ff70481461050557600080fd5b8063681fba16146104ba578063694ec2b1146104cf578063744b92e2146104e257600080fd5b8063599f643111610224578063599f64311461048d5780635c975abb1461049e5780636133dc24146104a957600080fd5b80634741062e1461044c5780634dbcb0ef1461046c57600080fd5b80632222dd421161029657806339aa92641161027b57806339aa92641461041e5780633f4ba83a146104315780634352fa9f1461043957600080fd5b80632222dd42146103eb578063351f0faf1461041057600080fd5b8063142a98fc116102c7578063142a98fc14610350578063147809b31461038a578063181f5a77146103a257600080fd5b8063087ae6df146102e3578063108ee5fc1461033b575b600080fd5b604080517f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006020820152015b60405180910390f35b61034e610349366004614210565b6107de565b005b61037d61035e36600461424e565b67ffffffffffffffff1660009081526010602052604090205460ff1690565b6040516103329190614281565b610392610887565b6040519015158152602001610332565b6103de6040518060400160405280601881526020017f45564d3245564d546f6c6c4f666652616d7020312e302e30000000000000000081525081565b6040516103329190614301565b6002546001600160a01b03165b6040516001600160a01b039091168152602001610332565b61034e6102de366004614314565b61034e61042c366004614210565b610914565b61034e61093e565b61034e6104473660046144e3565b610950565b61045f61045a366004614547565b610b8c565b60405161033291906145bf565b61047f61047a36600461464b565b610c54565b604051908152602001610332565b6005546001600160a01b03166103f8565b60005460ff16610392565b600d546001600160a01b03166103f8565b6104c2610d03565b60405161033291906147ac565b61034e6104dd3660046147bf565b610dc8565b61034e6104f03660046148e5565b610e41565b61034e611112565b6104c26111ed565b6013546011546040805163ffffffff80851682526401000000009094049093166020840152820152606001610332565b61034e61124f565b6104c261125f565b60005461010090046001600160a01b03166103f8565b6105636112bf565b60405161033291908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b61034e6105a436600461491e565b611360565b61034e6105b7366004614210565b61147a565b61047f6105ca36600461496d565b60176020526000908152604090205481565b604080516001815260006020820181905291810191909152606001610332565b600c546001600160a01b03166103f8565b61034e61061b3660046149d2565b6114a4565b6103f861062e366004614210565b611a77565b61034e610641366004614ab7565b611b4c565b6103f8610654366004614210565b6001600160a01b039081166000908152600360205260409020541690565b61034e610680366004614be9565b611b58565b61034e610693366004614210565b61228c565b61072e6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506040805160a081018252600e546001600160a01b038116825263ffffffff600160a01b820416602083015267ffffffffffffffff600160c01b909104811692820192909252600f548083166060830152680100000000000000009004909116608082015290565b6040805182516001600160a01b0316815260208084015163ffffffff16908201528282015167ffffffffffffffff908116928201929092526060808401518316908201526080928301519091169181019190915260a001610332565b61047f6122de565b61034e6107a0366004614d32565b612391565b61034e6107b3366004614de4565b6124fb565b61034e6107c63660046148e5565b612c59565b61034e6107d9366004614210565b612e03565b6107e6612e14565b6001600160a01b038116610826576040517f0958ef9b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280546001600160a01b038381166001600160a01b031983168117909355604080519190921680825260208201939093527f2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28910160405180910390a15050565b600254604080517f46f8e6d700000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916346f8e6d79160048083019260209291908290030181865afa1580156108ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061090e9190614eb1565b15905090565b61091c612e14565b600580546001600160a01b0319166001600160a01b0392909216919091179055565b610946612e14565b61094e612e73565b565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561098d57506005546001600160a01b03163314155b156109ab576040516307b66ab160e51b815260040160405180910390fd5b8151815181146109e7576040517f3959163300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075460005b81811015610a41576006600060078381548110610a0c57610a0c614ece565b60009182526020808320909101546001600160a01b03168352820192909252604001812055610a3a81614efa565b90506109ed565b5060005b82811015610b71576000858281518110610a6157610a61614ece565b6020026020010151905060006001600160a01b0316816001600160a01b031603610ab7576040517fe622e04000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610ac957610ac9614ece565b602002602001015160066000836001600160a01b03166001600160a01b03168152602001908152602001600020819055507f4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f81868481518110610b2e57610b2e614ece565b6020026020010151604051610b589291906001600160a01b03929092168252602082015260400190565b60405180910390a150610b6a81614efa565b9050610a45565b508351610b85906007906020870190614166565b5050505050565b80516060908067ffffffffffffffff811115610baa57610baa61435b565b604051908082528060200260200182016040528015610bd3578160200160208202803683370190505b50915060005b81811015610c4d5760066000858381518110610bf757610bf7614ece565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054838281518110610c3257610c32614ece565b6020908102919091010152610c4681614efa565b9050610bd9565b5050919050565b6000808260800151518360a001515160206014610c719190614f13565b610c7b9190614f2b565b610c86906086614f13565b610c909190614f13565b90506000610c9f601083614f2b565b9050610a28611bbc8560a00151516001610cb99190614f13565b610cc590618aac614f2b565b6156b8610cd28986614f13565b610cdc9190614f13565b610ce69190614f13565b610cf09190614f13565b610cfa9190614f13565b95945050505050565b60045460609067ffffffffffffffff811115610d2157610d2161435b565b604051908082528060200260200182016040528015610d4a578160200160208202803683370190505b50905060005b600454811015610dc457610d8a60048281548110610d7057610d70614ece565b6000918252602090912001546001600160a01b0316611a77565b828281518110610d9c57610d9c614ece565b6001600160a01b0390921660209283029190910190910152610dbd81614efa565b9050610d50565b5090565b333014610e01576040517f371a732800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60808101515115610e3557610e1e81608001518260c00151612f0f565b610e358160a001518260c0015183604001516130a2565b610e3e8161313c565b50565b610e49612e14565b6004546000819003610e87576040517f6987841e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03838116600090815260036020908152604091829020825180840190935254928316808352600160a01b9093046bffffffffffffffffffffffff169082015290610eeb576040516302721e1f60e61b815260040160405180910390fd5b826001600160a01b031681600001516001600160a01b031614610f3a576040517f6cc7b99800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004610f49600185614f4a565b81548110610f5957610f59614ece565b9060005260206000200160009054906101000a90046001600160a01b03169050600482602001516bffffffffffffffffffffffff1681548110610f9e57610f9e614ece565b6000918252602090912001546001600160a01b03166004610fc0600186614f4a565b81548110610fd057610fd0614ece565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600483602001516bffffffffffffffffffffffff168154811061102457611024614ece565b6000918252602080832090910180546001600160a01b039485166001600160a01b03199091161790558481015184841683526003909152604090912080546bffffffffffffffffffffffff909216600160a01b0291909216179055600480548061109057611090614f61565b60008281526020808220600019908401810180546001600160a01b03191690559092019092556001600160a01b03878116808452600383526040808520949094558351908152908716918101919091527f987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c910160405180910390a15050505050565b6001546001600160a01b031633146111715760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054336101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff8416178455600180546001600160a01b03191690556040516001600160a01b03919093041692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6060601680548060200260200160405190810160405280929190818152602001828054801561124557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611227575b5050505050905090565b611257612e14565b61094e6131fb565b60606004805480602002602001604051908101604052809291908181526020018280548015611245576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611227575050505050905090565b6112ea6040518060800160405280600081526020016000815260200160008152602001600081525090565b6040805160808101825260085481526009546020820152600a5491810191909152600b546060820181905242906000906113249083614f4a565b602084015184519192506113509161133c9084614f2b565b856040015161134b9190614f13565b613283565b6040840152506060820152919050565b60005461010090046001600160a01b03166001600160a01b0316336001600160a01b03161415801561139d57506005546001600160a01b03163314155b156113bb576040516307b66ab160e51b815260040160405180910390fd5b805179ffffffffffffffffffffffffffffffffffffffffffffffffffff1161140f576040517f3d9cbdab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114196008613299565b602081015160098190558151600855600a546114359190613283565b600a55602081810151825160408051928352928201527f8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d91015b60405180910390a150565b611482612e14565b600d80546001600160a01b0319166001600160a01b0392909216919091179055565b60005a604080516020601f8b018190048102820181019092528981529192508a3591818c0135916114fa91849163ffffffff851691908e908e908190840183828082843760009201919091525061334692505050565b6040805183815262ffffff600884901c1660208201527fb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62910160405180910390a16040805160608101825260115480825260125460ff808216602085015261010090910416928201929092529083146115b55760405162461bcd60e51b815260206004820152601560248201527f636f6e666967446967657374206d69736d6174636800000000000000000000006044820152606401611168565b6115c38b8b8b8b8b8b6133bf565b60007f000000000000000000000000000000000000000000000000000000000000000015611620576002826020015183604001516116019190614f77565b61160b9190614fb2565b611616906001614f77565b60ff169050611636565b6020820151611630906001614f77565b60ff1690505b8881146116855760405162461bcd60e51b815260206004820152601a60248201527f77726f6e67206e756d626572206f66207369676e6174757265730000000000006044820152606401611168565b8887146116d45760405162461bcd60e51b815260206004820152601e60248201527f7369676e617475726573206f7574206f6620726567697374726174696f6e00006044820152606401611168565b3360009081526014602090815260408083208151808301909252805460ff808216845292939192918401916101009091041660028111156117175761171761426b565b60028111156117285761172861426b565b90525090506002816020015160028111156117455761174561426b565b14801561177f57506016816000015160ff168154811061176757611767614ece565b6000918252602090912001546001600160a01b031633145b6117cb5760405162461bcd60e51b815260206004820152601860248201527f756e617574686f72697a6564207472616e736d697474657200000000000000006044820152606401611168565b5050505050600088886040516117e2929190614fd4565b6040519081900381206117f9918c90602001614fe4565b6040516020818303038152906040528051906020012090506118196141c7565b604080518082019091526000808252602082015260005b88811015611a5557600060018588846020811061184f5761184f614ece565b61185c91901a601b614f77565b8d8d8681811061186e5761186e614ece565b905060200201358c8c8781811061188757611887614ece565b90506020020135604051600081526020016040526040516118c4949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa1580156118e6573d6000803e3d6000fd5b505060408051601f198101516001600160a01b03811660009081526014602090815290849020838501909452835460ff8082168552929650929450840191610100900416600281111561193b5761193b61426b565b600281111561194c5761194c61426b565b90525092506001836020015160028111156119695761196961426b565b146119b65760405162461bcd60e51b815260206004820152601e60248201527f61646472657373206e6f7420617574686f72697a656420746f207369676e00006044820152606401611168565b8251849060ff16601f81106119cd576119cd614ece565b602002015115611a1f5760405162461bcd60e51b815260206004820152601460248201527f6e6f6e2d756e69717565207369676e61747572650000000000000000000000006044820152606401611168565b600184846000015160ff16601f8110611a3a57611a3a614ece565b9115156020909202015250611a4e81614efa565b9050611830565b5050505063ffffffff8110611a6c57611a6c615000565b505050505050505050565b6001600160a01b0380821660009081526003602052604081205490911680611ab2576040516302721e1f60e61b815260040160405180910390fd5b6001600160a01b038084166000908152600360209081526040918290205482517f21df0da700000000000000000000000000000000000000000000000000000000815292519316926321df0da79260048082019392918290030181865afa158015611b21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b459190615021565b9392505050565b610e3e60008083613346565b60005460ff1615611bab5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611168565b600260009054906101000a90046001600160a01b03166001600160a01b03166346f8e6d76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611bfe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c229190614eb1565b15611c58576040517e7b22b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600c546001600160a01b0316611c9a576040517f179ce99f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060820151516000819003611cdb576040517f7a21217700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008167ffffffffffffffff811115611cf657611cf661435b565b604051908082528060200260200182016040528015611da257816020015b611d8f60405180610140016040528060008152602001600067ffffffffffffffff16815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160608152602001606081526020016060815260200160006001600160a01b0316815260200160008152602001600081525090565b815260200190600190039081611d145790505b50905060008267ffffffffffffffff811115611dc057611dc061435b565b604051908082528060200260200182016040528015611de9578160200160208202803683370190505b5090506000611e9a7f31a97e998befaf21ee85c8e3e1879003d67d46abb0b6b552882133e4c782986d600e54604080516020808201949094527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201526001600160a01b039092166080808401919091528151808403909101815260a09092019052805191012090565b905060005b84811015611f3457600087606001518281518110611ebf57611ebf614ece565b6020026020010151806020019051810190611eda919061514d565b9050611ee68184613453565b848381518110611ef857611ef8614ece565b60200260200101818152505080858381518110611f1757611f17614ece565b60200260200101819052505080611f2d90614efa565b9050611e9f565b50600080611f558489608001518a60a001518b60c001518c60e00151613570565b600e549193509150600090600160a01b900463ffffffff16611f778442614f4a565b11905060005b87811015612280576000878281518110611f9957611f99614ece565b602002602001015190506000611fcc826020015167ffffffffffffffff1660009081526010602052604090205460ff1690565b90506002816003811115611fe257611fe261426b565b0361202b5760208201516040517f50a6e05200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611168565b8a15806120355750835b806120515750600381600381111561204f5761204f61426b565b145b612087576040517f6358b0d000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61209082613668565b6000808260038111156120a5576120a561426b565b1480156120b057508b155b15612102576120cc8a51876120c59190615269565b8e856137af565b905060006120dd8460e00151613992565b90506120ea8183306139f4565b8184610100018181516120fd9190614f4a565b905250505b60008260038111156121165761211661426b565b146121505760208084015167ffffffffffffffff166000908152601790915260409020546101008401805161214c908390614f4a565b9052505b60208381015167ffffffffffffffff166000908152601090915260408120805460ff1916600117905561218a61218585613a42565b613cdc565b60208086015167ffffffffffffffff1660009081526010909152604090208054919250829160ff191660018360038111156121c7576121c761426b565b021790555060008360038111156121e0576121e061426b565b1480156121fe575060038160038111156121fc576121fc61426b565b145b156122255760208085015167ffffffffffffffff1660009081526017909152604090208290555b836020015167ffffffffffffffff167f06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf826040516122639190614281565b60405180910390a2505050508061227990614efa565b9050611f7d565b50505050505050505050565b612294612e14565b600c80546001600160a01b0319166001600160a01b0383169081179091556040517f993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d490600090a250565b600061238c7f31a97e998befaf21ee85c8e3e1879003d67d46abb0b6b552882133e4c782986d600e54604080516020808201949094527f0000000000000000000000000000000000000000000000000000000000000000818301527f000000000000000000000000000000000000000000000000000000000000000060608201526001600160a01b039092166080808401919091528151808403909101815260a09092019052805191012090565b905090565b612399612e14565b80516001600160a01b03166123da576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051600e805460208085018051604080880180516001600160a01b039098167fffffffffffffffff0000000000000000000000000000000000000000000000009096168617600160a01b63ffffffff948516021777ffffffffffffffffffffffffffffffffffffffffffffffff16600160c01b67ffffffffffffffff998a16021790965560608089018051600f80546080808e018051948e166fffffffffffffffffffffffffffffffff199093169290921768010000000000000000948e1694909402939093179091558451988952955190941695870195909552955187169085015291518516938301939093529151909216908201527fd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa09060a00161146f565b855185518560ff16601f8311156125545760405162461bcd60e51b815260206004820152601060248201527f746f6f206d616e79207369676e657273000000000000000000000000000000006044820152606401611168565b600081116125a45760405162461bcd60e51b815260206004820152601260248201527f66206d75737420626520706f73697469766500000000000000000000000000006044820152606401611168565b8183146126185760405162461bcd60e51b8152602060048201526024808201527f6f7261636c6520616464726573736573206f7574206f6620726567697374726160448201527f74696f6e000000000000000000000000000000000000000000000000000000006064820152608401611168565b612623816003614f2b565b83116126715760405162461bcd60e51b815260206004820152601860248201527f6661756c74792d6f7261636c65206620746f6f206869676800000000000000006044820152606401611168565b612679612e14565b6040805160c0810182528a8152602081018a905260ff8916918101919091526060810187905267ffffffffffffffff8616608082015260a081018590525b601554156127c9576015546000906126d190600190614f4a565b90506000601582815481106126e8576126e8614ece565b6000918252602082200154601680546001600160a01b039092169350908490811061271557612715614ece565b60009182526020808320909101546001600160a01b03858116845260149092526040808420805461ffff199081169091559290911680845292208054909116905560158054919250908061276b5761276b614f61565b600082815260209020810160001990810180546001600160a01b0319169055019055601680548061279e5761279e614f61565b600082815260209020810160001990810180546001600160a01b0319169055019055506126b7915050565b60005b815151811015612b0657600060146000846000015184815181106127f2576127f2614ece565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff16600281111561282f5761282f61426b565b1461287c5760405162461bcd60e51b815260206004820152601760248201527f7265706561746564207369676e657220616464726573730000000000000000006044820152606401611168565b6040805180820190915260ff821681526001602082015282518051601491600091859081106128ad576128ad614ece565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff1916176101008360028111156129065761290661426b565b0217905550600091506129169050565b601460008460200151848151811061293057612930614ece565b6020908102919091018101516001600160a01b0316825281019190915260400160002054610100900460ff16600281111561296d5761296d61426b565b146129ba5760405162461bcd60e51b815260206004820152601c60248201527f7265706561746564207472616e736d69747465722061646472657373000000006044820152606401611168565b6040805180820190915260ff8216815260208101600281525060146000846020015184815181106129ed576129ed614ece565b6020908102919091018101516001600160a01b03168252818101929092526040016000208251815460ff90911660ff19821681178355928401519192839161ffff191617610100836002811115612a4657612a4661426b565b021790555050825180516015925083908110612a6457612a64614ece565b602090810291909101810151825460018101845560009384529282902090920180546001600160a01b0319166001600160a01b039093169290921790915582015180516016919083908110612abb57612abb614ece565b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b03909216919091179055612aff81614efa565b90506127cc565b5060408101516012805460ff191660ff9092169190911790556013805467ffffffff0000000019811664010000000063ffffffff438116820292831785559083048116936001939092600092612b6392869290821691161761527d565b92506101000a81548163ffffffff021916908363ffffffff160217905550612bc24630601360009054906101000a900463ffffffff1663ffffffff16856000015186602001518760400151886060015189608001518a60a00151613df7565b6011819055825180516012805460ff9092166101000261ff001990921691909117905560135460208501516040808701516060880151608089015160a08a015193517f1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e0598612c44988b98919763ffffffff9092169690959194919391926152a5565b60405180910390a15050505050505050505050565b612c61612e14565b6001600160a01b0382161580612c7e57506001600160a01b038116155b15612cb5576040517f6c2a418000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116600090815260036020908152604091829020825180840190935254928316808352600160a01b9093046bffffffffffffffffffffffff16908201529015612d33576040517f3caf458500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03828116808352600480546bffffffffffffffffffffffff908116602080870191825288861660008181526003835260408082208a519551909616600160a01b0294909816939093179093558354600181018555939091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b90920180546001600160a01b031916821790558351908152908101919091527f95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c91015b60405180910390a1505050565b612e0b612e14565b610e3e81613e84565b60005461010090046001600160a01b0316331461094e5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401611168565b60005460ff16612ec55760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401611168565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6000805b835181101561300657600060066000868481518110612f3457612f34614ece565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054905080600003612fc357848281518110612f7957612f79614ece565b60200260200101516040517f9a655f7b00000000000000000000000000000000000000000000000000000000815260040161116891906001600160a01b0391909116815260200190565b838281518110612fd557612fd5614ece565b602002602001015181612fe89190614f2b565b612ff29084614f13565b92505080612fff90614efa565b9050612f13565b50801561309d576130176008613299565b600a54811115613053576040517f3bfa6f3800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600860020160008282546130689190614f4a565b90915550506040518181527fcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba37890602001612df6565b505050565b81518351146130dd576040517f7bdc0b2c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8351811015613136576131268482815181106130fe576130fe614ece565b602002602001015184838151811061311857613118614ece565b6020026020010151846139f4565b61312f81614efa565b90506130e0565b50505050565b60408101516001600160a01b03163b6131525750565b600c546040517f5c9fa11e0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690635c9fa11e9061319b90849060040161533b565b6020604051808303816000875af11580156131ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131de9190614eb1565b610e3e57604051631dc9e9b560e31b815260040160405180910390fd5b60005460ff161561324e5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401611168565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612ef23390565b60008183106132925781611b45565b5090919050565b60018101546002820154429114806132b45750808260030154145b156132bd575050565b8160010154826002015411156132ff576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008260030154826133119190614f4a565b60018401548454919250613338916133299084614f2b565b856002015461134b9190614f13565b600284015550600390910155565b306001600160a01b031663be9b03f18280602001905181019061336991906154f0565b60006040518363ffffffff1660e01b81526004016133889291906156b8565b600060405180830381600087803b1580156133a257600080fd5b505af11580156133b6573d6000803e3d6000fd5b50505050505050565b60006133cc826020614f2b565b6133d7856020614f2b565b6133e388610144614f13565b6133ed9190614f13565b6133f79190614f13565b613402906000614f13565b90503681146133b65760405162461bcd60e51b815260206004820152601860248201527f63616c6c64617461206c656e677468206d69736d6174636800000000000000006044820152606401611168565b60008060001b828460200151856040015186606001518760800151805190602001208860a0015160405160200161348a91906147ac565b604051602081830303815290604052805190602001208960c001516040516020016134b591906145bf565b604051602081830303815290604052805190602001208a61012001518b60e001518c61010001516040516020016135529b9a999897969594939291909a8b5260208b019990995267ffffffffffffffff9790971660408a01526001600160a01b0395861660608a0152938516608089015260a088019290925260c087015260e0860152610100850152166101208301526101408201526101600190565b60405160208183030381529060405280519060200120905092915050565b60008060005a600d546040517fe71e65ce0000000000000000000000000000000000000000000000000000000081529192506000916001600160a01b039091169063e71e65ce906135cd908c908c908c908c908c90600401615791565b6020604051808303816000875af11580156135ec573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061361091906157e3565b90506000811161364c576040517f894882b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805a6136589084614f4a565b9350935050509550959350505050565b80517f0000000000000000000000000000000000000000000000000000000000000000146136c85780516040517fd44bc9eb0000000000000000000000000000000000000000000000000000000081526004810191909152602401611168565b600f5460a0820151516801000000000000000090910467ffffffffffffffff1610806136fe57508060c00151518160a001515114155b156137475760208101516040517f099d3f7200000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401611168565b600f5460808201515167ffffffffffffffff9091161015610e3e57600f546080820151516040517f8693378900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301526024820152604401611168565b60008060006137c18460e00151613992565b6001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156137fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138229190615021565b905060005b85602001515181101561389b57816001600160a01b03168660200151828151811061385457613854614ece565b60200260200101516001600160a01b03160361388b578560400151818151811061388057613880614ece565b602002602001015192505b61389481614efa565b9050613827565b50816138de576040517fce480bcc0000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401611168565b6000670de0b6b3a7640000833a8761012001516138fb8b8a610c54565b6139059190614f13565b61390f9190614f2b565b6139199190614f2b565b6139239190615269565b90508461010001518111156139885760208501516101008601516040517f394a2c2700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9092166004830152602482018390526044820152606401611168565b9695505050505050565b6001600160a01b0381811660009081526003602052604090205416806139ef576040517fbf16aab60000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401611168565b919050565b6040517fea6192a20000000000000000000000000000000000000000000000000000000081526001600160a01b0382811660048301526024820184905284169063ea6192a290604401613388565b613a93604051806101000160405280600081526020016060815260200160006001600160a01b0316815260200160608152602001606081526020016060815260200160608152602001600081525090565b600080613ab38460a001518560c001518660e00151876101000151613f33565b8151919350915060008167ffffffffffffffff811115613ad557613ad561435b565b604051908082528060200260200182016040528015613afe578160200160208202803683370190505b50905060008267ffffffffffffffff811115613b1c57613b1c61435b565b604051908082528060200260200182016040528015613b45578160200160208202803683370190505b50905060005b83811015613c49576000613b77878381518110613b6a57613b6a614ece565b6020026020010151613992565b905080838381518110613b8c57613b8c614ece565b60200260200101906001600160a01b031690816001600160a01b031681525050806001600160a01b03166321df0da76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613bea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c0e9190615021565b848381518110613c2057613c20614ece565b6001600160a01b039092166020928302919091019091015250613c4281614efa565b9050613b4b565b50604051806101000160405280886000015181526020018860400151604051602001613c8491906001600160a01b0391909116815260200190565b604051602081830303815290604052815260200188606001516001600160a01b031681526020018860800151815260200183815260200182815260200185815260200188610120015181525095505050505050919050565b6040517f694ec2b1000000000000000000000000000000000000000000000000000000008152600090309063694ec2b190613d1b90859060040161533b565b600060405180830381600087803b158015613d3557600080fd5b505af1925050508015613d46575060015b613def573d808015613d74576040519150601f19603f3d011682016040523d82523d6000602084013e613d79565b606091505b50613d83816157fc565b7fffffffff0000000000000000000000000000000000000000000000000000000016631dc9e9b560e31b03613dbb5750600392915050565b806040517fcf19edfd0000000000000000000000000000000000000000000000000000000081526004016111689190614301565b506002919050565b6000808a8a8a8a8a8a8a8a8a604051602001613e1b9998979695949392919061584c565b60408051601f1981840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e01000000000000000000000000000000000000000000000000000000000000179150509998505050505050505050565b336001600160a01b03821603613edc5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401611168565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929361010090910416917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60608060005b8651811015613fbb57846001600160a01b0316878281518110613f5e57613f5e614ece565b60200260200101516001600160a01b031603613fab5783868281518110613f8757613f87614ece565b60200260200101818151613f9b9190614f13565b90525086925085915061415d9050565b613fb481614efa565b9050613f39565b50600086516001613fcc9190614f13565b67ffffffffffffffff811115613fe457613fe461435b565b60405190808252806020026020018201604052801561400d578160200160208202803683370190505b5090506000865160016140209190614f13565b67ffffffffffffffff8111156140385761403861435b565b604051908082528060200260200182016040528015614061578160200160208202803683370190505b50905060005b88518110156141035788818151811061408257614082614ece565b602002602001015183828151811061409c5761409c614ece565b60200260200101906001600160a01b031690816001600160a01b0316815250508781815181106140ce576140ce614ece565b60200260200101518282815181106140e8576140e8614ece565b60209081029190910101526140fc81614efa565b9050614067565b50858289518151811061411857614118614ece565b60200260200101906001600160a01b031690816001600160a01b031681525050848188518151811061414c5761414c614ece565b602090810291909101015290925090505b94509492505050565b8280548282559060005260206000209081019282156141bb579160200282015b828111156141bb57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614186565b50610dc49291506141e6565b604051806103e00160405280601f906020820280368337509192915050565b5b80821115610dc457600081556001016141e7565b6001600160a01b0381168114610e3e57600080fd5b60006020828403121561422257600080fd5b8135611b45816141fb565b67ffffffffffffffff81168114610e3e57600080fd5b80356139ef8161422d565b60006020828403121561426057600080fd5b8135611b458161422d565b634e487b7160e01b600052602160045260246000fd5b60208101600483106142a357634e487b7160e01b600052602160045260246000fd5b91905290565b60005b838110156142c45781810151838201526020016142ac565b838111156131365750506000910152565b600081518084526142ed8160208601602086016142a9565b601f01601f19169290920160200192915050565b602081526000611b4560208301846142d5565b60006020828403121561432657600080fd5b813567ffffffffffffffff81111561433d57600080fd5b82016101008185031215611b4557600080fd5b80356139ef816141fb565b634e487b7160e01b600052604160045260246000fd5b604051610140810167ffffffffffffffff811182821017156143955761439561435b565b60405290565b604051610100810167ffffffffffffffff811182821017156143955761439561435b565b604051601f8201601f1916810167ffffffffffffffff811182821017156143e8576143e861435b565b604052919050565b600067ffffffffffffffff82111561440a5761440a61435b565b5060051b60200190565b600082601f83011261442557600080fd5b8135602061443a614435836143f0565b6143bf565b82815260059290921b8401810191818101908684111561445957600080fd5b8286015b8481101561447d578035614470816141fb565b835291830191830161445d565b509695505050505050565b600082601f83011261449957600080fd5b813560206144a9614435836143f0565b82815260059290921b840181019181810190868411156144c857600080fd5b8286015b8481101561447d57803583529183019183016144cc565b600080604083850312156144f657600080fd5b823567ffffffffffffffff8082111561450e57600080fd5b61451a86838701614414565b9350602085013591508082111561453057600080fd5b5061453d85828601614488565b9150509250929050565b60006020828403121561455957600080fd5b813567ffffffffffffffff81111561457057600080fd5b61457c84828501614414565b949350505050565b600081518084526020808501945080840160005b838110156145b457815187529582019590820190600101614598565b509495945050505050565b602081526000611b456020830184614584565b600067ffffffffffffffff8211156145ec576145ec61435b565b50601f01601f191660200190565b600082601f83011261460b57600080fd5b8135614619614435826145d2565b81815284602083860101111561462e57600080fd5b816020850160208301376000918101602001919091529392505050565b6000806040838503121561465e57600080fd5b82359150602083013567ffffffffffffffff8082111561467d57600080fd5b90840190610140828703121561469257600080fd5b61469a614371565b823581526146aa60208401614243565b60208201526146bb60408401614350565b60408201526146cc60608401614350565b60608201526080830135828111156146e357600080fd5b6146ef888286016145fa565b60808301525060a08301358281111561470757600080fd5b61471388828601614414565b60a08301525060c08301358281111561472b57600080fd5b61473788828601614488565b60c08301525061474960e08401614350565b60e08201526101008381013590820152610120928301359281019290925250919491935090915050565b600081518084526020808501945080840160005b838110156145b45781516001600160a01b031687529582019590820190600101614787565b602081526000611b456020830184614773565b6000602082840312156147d157600080fd5b813567ffffffffffffffff808211156147e957600080fd5b9083019061010082860312156147fe57600080fd5b61480661439b565b8235815260208301358281111561481c57600080fd5b614828878286016145fa565b60208301525061483a60408401614350565b604082015260608301358281111561485157600080fd5b61485d878286016145fa565b60608301525060808301358281111561487557600080fd5b61488187828601614414565b60808301525060a08301358281111561489957600080fd5b6148a587828601614414565b60a08301525060c0830135828111156148bd57600080fd5b6148c987828601614488565b60c08301525060e083013560e082015280935050505092915050565b600080604083850312156148f857600080fd5b8235614903816141fb565b91506020830135614913816141fb565b809150509250929050565b60006040828403121561493057600080fd5b6040516040810181811067ffffffffffffffff821117156149535761495361435b565b604052823581526020928301359281019290925250919050565b60006020828403121561497f57600080fd5b5035919050565b60008083601f84011261499857600080fd5b50813567ffffffffffffffff8111156149b057600080fd5b6020830191508360208260051b85010111156149cb57600080fd5b9250929050565b60008060008060008060008060e0898b0312156149ee57600080fd5b606089018a8111156149ff57600080fd5b8998503567ffffffffffffffff80821115614a1957600080fd5b818b0191508b601f830112614a2d57600080fd5b813581811115614a3c57600080fd5b8c6020828501011115614a4e57600080fd5b6020830199508098505060808b0135915080821115614a6c57600080fd5b614a788c838d01614986565b909750955060a08b0135915080821115614a9157600080fd5b50614a9e8b828c01614986565b999c989b50969995989497949560c00135949350505050565b600060208284031215614ac957600080fd5b813567ffffffffffffffff811115614ae057600080fd5b61457c848285016145fa565b600082601f830112614afd57600080fd5b81356020614b0d614435836143f0565b82815260059290921b84018101918181019086841115614b2c57600080fd5b8286015b8481101561447d578035614b438161422d565b8352918301918301614b30565b600082601f830112614b6157600080fd5b81356020614b71614435836143f0565b82815260059290921b84018101918181019086841115614b9057600080fd5b8286015b8481101561447d57803567ffffffffffffffff811115614bb45760008081fd5b614bc28986838b01016145fa565b845250918301918301614b94565b8015158114610e3e57600080fd5b80356139ef81614bd0565b60008060408385031215614bfc57600080fd5b823567ffffffffffffffff80821115614c1457600080fd5b908401906101008287031215614c2957600080fd5b614c3161439b565b823582811115614c4057600080fd5b614c4c88828601614aec565b825250602083013582811115614c6157600080fd5b614c6d88828601614414565b602083015250604083013582811115614c8557600080fd5b614c9188828601614488565b604083015250606083013582811115614ca957600080fd5b614cb588828601614b50565b606083015250608083013582811115614ccd57600080fd5b614cd988828601614488565b60808301525060a083013560a082015260c083013582811115614cfb57600080fd5b614d0788828601614488565b60c08301525060e083013560e0820152809450505050614d2960208401614bde565b90509250929050565b600060a08284031215614d4457600080fd5b60405160a0810181811067ffffffffffffffff82111715614d6757614d6761435b565b6040528235614d75816141fb565b8152602083013563ffffffff81168114614d8e57600080fd5b60208201526040830135614da18161422d565b60408201526060830135614db48161422d565b60608201526080830135614dc78161422d565b60808201529392505050565b803560ff811681146139ef57600080fd5b60008060008060008060c08789031215614dfd57600080fd5b863567ffffffffffffffff80821115614e1557600080fd5b614e218a838b01614414565b97506020890135915080821115614e3757600080fd5b614e438a838b01614414565b9650614e5160408a01614dd3565b95506060890135915080821115614e6757600080fd5b614e738a838b016145fa565b9450614e8160808a01614243565b935060a0890135915080821115614e9757600080fd5b50614ea489828a016145fa565b9150509295509295509295565b600060208284031215614ec357600080fd5b8151611b4581614bd0565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201614f0c57614f0c614ee4565b5060010190565b60008219821115614f2657614f26614ee4565b500190565b6000816000190483118215151615614f4557614f45614ee4565b500290565b600082821015614f5c57614f5c614ee4565b500390565b634e487b7160e01b600052603160045260246000fd5b600060ff821660ff84168060ff03821115614f9457614f94614ee4565b019392505050565b634e487b7160e01b600052601260045260246000fd5b600060ff831680614fc557614fc5614f9c565b8060ff84160491505092915050565b8183823760009101908152919050565b8281526060826020830137600060809190910190815292915050565b634e487b7160e01b600052600160045260246000fd5b80516139ef816141fb565b60006020828403121561503357600080fd5b8151611b45816141fb565b80516139ef8161422d565b600082601f83011261505a57600080fd5b8151615068614435826145d2565b81815284602083860101111561507d57600080fd5b61457c8260208301602087016142a9565b600082601f83011261509f57600080fd5b815160206150af614435836143f0565b82815260059290921b840181019181810190868411156150ce57600080fd5b8286015b8481101561447d5780516150e5816141fb565b83529183019183016150d2565b600082601f83011261510357600080fd5b81516020615113614435836143f0565b82815260059290921b8401810191818101908684111561513257600080fd5b8286015b8481101561447d5780518352918301918301615136565b60006020828403121561515f57600080fd5b815167ffffffffffffffff8082111561517757600080fd5b90830190610140828603121561518c57600080fd5b615194614371565b825181526151a46020840161503e565b60208201526151b560408401615016565b60408201526151c660608401615016565b60608201526080830151828111156151dd57600080fd5b6151e987828601615049565b60808301525060a08301518281111561520157600080fd5b61520d8782860161508e565b60a08301525060c08301518281111561522557600080fd5b615231878286016150f2565b60c08301525061524360e08401615016565b60e082015261010083810151908201526101209283015192810192909252509392505050565b60008261527857615278614f9c565b500490565b600063ffffffff80831681851680830382111561529c5761529c614ee4565b01949350505050565b600061012063ffffffff808d1684528b6020850152808b166040850152508060608401526152d58184018a614773565b905082810360808401526152e98189614773565b905060ff871660a084015282810360c084015261530681876142d5565b905067ffffffffffffffff851660e084015282810361010084015261532b81856142d5565b9c9b505050505050505050505050565b6020815281516020820152600060208301516101008060408501526153646101208501836142d5565b9150604085015161538060608601826001600160a01b03169052565b506060850151601f198086850301608087015261539d84836142d5565b935060808701519150808685030160a08701526153ba8483614773565b935060a08701519150808685030160c08701526153d78483614773565b935060c08701519150808685030160e0870152506153f58382614584565b92505060e085015181850152508091505092915050565b600082601f83011261541d57600080fd5b8151602061542d614435836143f0565b82815260059290921b8401810191818101908684111561544c57600080fd5b8286015b8481101561447d5780516154638161422d565b8352918301918301615450565b600082601f83011261548157600080fd5b81516020615491614435836143f0565b82815260059290921b840181019181810190868411156154b057600080fd5b8286015b8481101561447d57805167ffffffffffffffff8111156154d45760008081fd5b6154e28986838b0101615049565b8452509183019183016154b4565b60006020828403121561550257600080fd5b815167ffffffffffffffff8082111561551a57600080fd5b90830190610100828603121561552f57600080fd5b61553761439b565b82518281111561554657600080fd5b6155528782860161540c565b82525060208301518281111561556757600080fd5b6155738782860161508e565b60208301525060408301518281111561558b57600080fd5b615597878286016150f2565b6040830152506060830151828111156155af57600080fd5b6155bb87828601615470565b6060830152506080830151828111156155d357600080fd5b6155df878286016150f2565b60808301525060a083015160a082015260c08301518281111561560157600080fd5b61560d878286016150f2565b60c08301525060e083015160e082015280935050505092915050565b600081518084526020808501945080840160005b838110156145b457815167ffffffffffffffff168752958201959082019060010161563d565b600081518084526020808501808196508360051b8101915082860160005b858110156156ab5782840389526156998483516142d5565b98850198935090840190600101615681565b5091979650505050505050565b60408152600083516101008060408501526156d7610140850183615629565b91506020860151603f19808685030160608701526156f58483614773565b935060408801519150808685030160808701526157128483614584565b935060608801519150808685030160a087015261572f8483615663565b935060808801519150808685030160c087015261574c8483614584565b935060a088015160e087015260c08801519150808685030183870152506157738382614584565b60e088015161012087015286151560208701529350611b4592505050565b60a0815260006157a460a0830188614584565b82810360208401526157b68188614584565b905085604084015282810360608401526157d08186614584565b9150508260808301529695505050505050565b6000602082840312156157f557600080fd5b5051919050565b6000815160208301517fffffffff00000000000000000000000000000000000000000000000000000000808216935060048310156158445780818460040360031b1b83161693505b505050919050565b60006101208b83526001600160a01b038b16602084015267ffffffffffffffff808b1660408501528160608501526158868285018b614773565b9150838203608085015261589a828a614773565b915060ff881660a085015283820360c08501526158b782886142d5565b90861660e0850152838103610100850152905061532b81856142d556fea164736f6c634300080f000a",
}

var EVM2EVMTollOffRampHelperABI = EVM2EVMTollOffRampHelperMetaData.ABI

var EVM2EVMTollOffRampHelperBin = EVM2EVMTollOffRampHelperMetaData.Bin

func DeployEVM2EVMTollOffRampHelper(auth *bind.TransactOpts, backend bind.ContractBackend, sourceChainId *big.Int, chainId *big.Int, offRampConfig BaseOffRampInterfaceOffRampConfig, blobVerifier common.Address, afn common.Address, sourceTokens []common.Address, pools []common.Address, rateLimiterConfig AggregateRateLimiterInterfaceRateLimiterConfig, tokenLimitsAdmin common.Address) (common.Address, *types.Transaction, *EVM2EVMTollOffRampHelper, error) {
	parsed, err := EVM2EVMTollOffRampHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EVM2EVMTollOffRampHelperBin), backend, sourceChainId, chainId, offRampConfig, blobVerifier, afn, sourceTokens, pools, rateLimiterConfig, tokenLimitsAdmin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EVM2EVMTollOffRampHelper{EVM2EVMTollOffRampHelperCaller: EVM2EVMTollOffRampHelperCaller{contract: contract}, EVM2EVMTollOffRampHelperTransactor: EVM2EVMTollOffRampHelperTransactor{contract: contract}, EVM2EVMTollOffRampHelperFilterer: EVM2EVMTollOffRampHelperFilterer{contract: contract}}, nil
}

type EVM2EVMTollOffRampHelper struct {
	address common.Address
	abi     abi.ABI
	EVM2EVMTollOffRampHelperCaller
	EVM2EVMTollOffRampHelperTransactor
	EVM2EVMTollOffRampHelperFilterer
}

type EVM2EVMTollOffRampHelperCaller struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampHelperTransactor struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampHelperFilterer struct {
	contract *bind.BoundContract
}

type EVM2EVMTollOffRampHelperSession struct {
	Contract     *EVM2EVMTollOffRampHelper
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOffRampHelperCallerSession struct {
	Contract *EVM2EVMTollOffRampHelperCaller
	CallOpts bind.CallOpts
}

type EVM2EVMTollOffRampHelperTransactorSession struct {
	Contract     *EVM2EVMTollOffRampHelperTransactor
	TransactOpts bind.TransactOpts
}

type EVM2EVMTollOffRampHelperRaw struct {
	Contract *EVM2EVMTollOffRampHelper
}

type EVM2EVMTollOffRampHelperCallerRaw struct {
	Contract *EVM2EVMTollOffRampHelperCaller
}

type EVM2EVMTollOffRampHelperTransactorRaw struct {
	Contract *EVM2EVMTollOffRampHelperTransactor
}

func NewEVM2EVMTollOffRampHelper(address common.Address, backend bind.ContractBackend) (*EVM2EVMTollOffRampHelper, error) {
	abi, err := abi.JSON(strings.NewReader(EVM2EVMTollOffRampHelperABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEVM2EVMTollOffRampHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelper{address: address, abi: abi, EVM2EVMTollOffRampHelperCaller: EVM2EVMTollOffRampHelperCaller{contract: contract}, EVM2EVMTollOffRampHelperTransactor: EVM2EVMTollOffRampHelperTransactor{contract: contract}, EVM2EVMTollOffRampHelperFilterer: EVM2EVMTollOffRampHelperFilterer{contract: contract}}, nil
}

func NewEVM2EVMTollOffRampHelperCaller(address common.Address, caller bind.ContractCaller) (*EVM2EVMTollOffRampHelperCaller, error) {
	contract, err := bindEVM2EVMTollOffRampHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperCaller{contract: contract}, nil
}

func NewEVM2EVMTollOffRampHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*EVM2EVMTollOffRampHelperTransactor, error) {
	contract, err := bindEVM2EVMTollOffRampHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperTransactor{contract: contract}, nil
}

func NewEVM2EVMTollOffRampHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*EVM2EVMTollOffRampHelperFilterer, error) {
	contract, err := bindEVM2EVMTollOffRampHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperFilterer{contract: contract}, nil
}

func bindEVM2EVMTollOffRampHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EVM2EVMTollOffRampHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOffRampHelper.Contract.EVM2EVMTollOffRampHelperCaller.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.EVM2EVMTollOffRampHelperTransactor.contract.Transfer(opts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.EVM2EVMTollOffRampHelperTransactor.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EVM2EVMTollOffRampHelper.Contract.contract.Call(opts, result, method, params...)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.contract.Transfer(opts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.contract.Transact(opts, method, params...)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "calculateCurrentTokenBucketState")

	if err != nil {
		return *new(AggregateRateLimiterInterfaceTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(AggregateRateLimiterInterfaceTokenBucket)).(*AggregateRateLimiterInterfaceTokenBucket)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMTollOffRampHelper.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) CalculateCurrentTokenBucketState() (AggregateRateLimiterInterfaceTokenBucket, error) {
	return _EVM2EVMTollOffRampHelper.Contract.CalculateCurrentTokenBucketState(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "ccipReceive", arg0)

	if err != nil {
		return err
	}

	return err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMTollOffRampHelper.Contract.CcipReceive(&_EVM2EVMTollOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) CcipReceive(arg0 CCIPAny2EVMMessageFromSender) error {
	return _EVM2EVMTollOffRampHelper.Contract.CcipReceive(&_EVM2EVMTollOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) FeeTaken(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "feeTaken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) FeeTaken(arg0 *big.Int) (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.FeeTaken(&_EVM2EVMTollOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) FeeTaken(arg0 *big.Int) (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.FeeTaken(&_EVM2EVMTollOffRampHelper.CallOpts, arg0)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetAFN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getAFN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetAFN(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetAFN() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetAFN(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetBlobVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getBlobVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetBlobVerifier() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetBlobVerifier(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetBlobVerifier() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetBlobVerifier(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getChainIDs")

	outstruct := new(GetChainIDs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetChainIDs(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetChainIDs() (GetChainIDs,

	error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetChainIDs(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(BaseOffRampInterfaceOffRampConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseOffRampInterfaceOffRampConfig)).(*BaseOffRampInterfaceOffRampConfig)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetConfig(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetConfig() (BaseOffRampInterfaceOffRampConfig, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetConfig(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getDestinationToken", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetDestinationToken(&_EVM2EVMTollOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetDestinationToken(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetDestinationToken(&_EVM2EVMTollOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getDestinationTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetDestinationTokens(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetDestinationTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetDestinationTokens(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getExecutionState", sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetExecutionState(&_EVM2EVMTollOffRampHelper.CallOpts, sequenceNumber)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetExecutionState(sequenceNumber uint64) (uint8, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetExecutionState(&_EVM2EVMTollOffRampHelper.CallOpts, sequenceNumber)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getPool", sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPool(&_EVM2EVMTollOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetPool(sourceToken common.Address) (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPool(&_EVM2EVMTollOffRampHelper.CallOpts, sourceToken)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getPoolTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPoolTokens(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetPoolTokens() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPoolTokens(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getPricesForTokens", tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPricesForTokens(&_EVM2EVMTollOffRampHelper.CallOpts, tokens)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetPricesForTokens(tokens []common.Address) ([]*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetPricesForTokens(&_EVM2EVMTollOffRampHelper.CallOpts, tokens)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetRouter(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetRouter() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetRouter(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "getTokenLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetTokenLimitAdmin(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) GetTokenLimitAdmin() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.GetTokenLimitAdmin(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) IsAFNHealthy(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "isAFNHealthy")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMTollOffRampHelper.Contract.IsAFNHealthy(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) IsAFNHealthy() (bool, error) {
	return _EVM2EVMTollOffRampHelper.Contract.IsAFNHealthy(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(LatestConfigDetails)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMTollOffRampHelper.Contract.LatestConfigDetails(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) LatestConfigDetails() (LatestConfigDetails,

	error) {
	return _EVM2EVMTollOffRampHelper.Contract.LatestConfigDetails(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

	error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(LatestConfigDigestAndEpoch)
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMTollOffRampHelper.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) LatestConfigDigestAndEpoch() (LatestConfigDigestAndEpoch,

	error) {
	return _EVM2EVMTollOffRampHelper.Contract.LatestConfigDigestAndEpoch(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) MetadataHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "metadataHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) MetadataHash() ([32]byte, error) {
	return _EVM2EVMTollOffRampHelper.Contract.MetadataHash(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) MetadataHash() ([32]byte, error) {
	return _EVM2EVMTollOffRampHelper.Contract.MetadataHash(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) OverheadGasToll(opts *bind.CallOpts, merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "overheadGasToll", merkleGasShare, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) OverheadGasToll(merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.OverheadGasToll(&_EVM2EVMTollOffRampHelper.CallOpts, merkleGasShare, message)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) OverheadGasToll(merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error) {
	return _EVM2EVMTollOffRampHelper.Contract.OverheadGasToll(&_EVM2EVMTollOffRampHelper.CallOpts, merkleGasShare, message)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Owner(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) Owner() (common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Owner(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Paused() (bool, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Paused(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) Paused() (bool, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Paused(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Transmitters(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) Transmitters() ([]common.Address, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Transmitters(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EVM2EVMTollOffRampHelper.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOffRampHelper.Contract.TypeAndVersion(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperCallerSession) TypeAndVersion() (string, error) {
	return _EVM2EVMTollOffRampHelper.Contract.TypeAndVersion(&_EVM2EVMTollOffRampHelper.CallOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "acceptOwnership")
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.AcceptOwnership(&_EVM2EVMTollOffRampHelper.TransactOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.AcceptOwnership(&_EVM2EVMTollOffRampHelper.TransactOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "addPool", token, pool)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.AddPool(&_EVM2EVMTollOffRampHelper.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) AddPool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.AddPool(&_EVM2EVMTollOffRampHelper.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) Execute(opts *bind.TransactOpts, report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "execute", report, manualExecution)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Execute(&_EVM2EVMTollOffRampHelper.TransactOpts, report, manualExecution)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) Execute(report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Execute(&_EVM2EVMTollOffRampHelper.TransactOpts, report, manualExecution)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "executeSingleMessage", message)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRampHelper.TransactOpts, message)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) ExecuteSingleMessage(message CCIPAny2EVMMessageFromSender) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.ExecuteSingleMessage(&_EVM2EVMTollOffRampHelper.TransactOpts, message)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "pause")
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Pause(&_EVM2EVMTollOffRampHelper.TransactOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) Pause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Pause(&_EVM2EVMTollOffRampHelper.TransactOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "removePool", token, pool)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.RemovePool(&_EVM2EVMTollOffRampHelper.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) RemovePool(token common.Address, pool common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.RemovePool(&_EVM2EVMTollOffRampHelper.TransactOpts, token, pool)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "report", executableMessages)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Report(executableMessages []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Report(&_EVM2EVMTollOffRampHelper.TransactOpts, executableMessages)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) Report(executableMessages []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Report(&_EVM2EVMTollOffRampHelper.TransactOpts, executableMessages)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setAFN", afn)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetAFN(&_EVM2EVMTollOffRampHelper.TransactOpts, afn)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetAFN(afn common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetAFN(&_EVM2EVMTollOffRampHelper.TransactOpts, afn)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetBlobVerifier(opts *bind.TransactOpts, blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setBlobVerifier", blobVerifier)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetBlobVerifier(&_EVM2EVMTollOffRampHelper.TransactOpts, blobVerifier)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetBlobVerifier(blobVerifier common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetBlobVerifier(&_EVM2EVMTollOffRampHelper.TransactOpts, blobVerifier)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetConfig(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setConfig", config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig(&_EVM2EVMTollOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetConfig(config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig(&_EVM2EVMTollOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetConfig0(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setConfig0", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig0(&_EVM2EVMTollOffRampHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetConfig0(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetConfig0(&_EVM2EVMTollOffRampHelper.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setPrices", tokens, prices)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetPrices(&_EVM2EVMTollOffRampHelper.TransactOpts, tokens, prices)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetPrices(tokens []common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetPrices(&_EVM2EVMTollOffRampHelper.TransactOpts, tokens, prices)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setRateLimiterConfig", config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetRateLimiterConfig(&_EVM2EVMTollOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetRateLimiterConfig(config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetRateLimiterConfig(&_EVM2EVMTollOffRampHelper.TransactOpts, config)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setRouter", router)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetRouter(&_EVM2EVMTollOffRampHelper.TransactOpts, router)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetRouter(router common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetRouter(&_EVM2EVMTollOffRampHelper.TransactOpts, router)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "setTokenLimitAdmin", newAdmin)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetTokenLimitAdmin(&_EVM2EVMTollOffRampHelper.TransactOpts, newAdmin)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) SetTokenLimitAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.SetTokenLimitAdmin(&_EVM2EVMTollOffRampHelper.TransactOpts, newAdmin)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "transferOwnership", to)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.TransferOwnership(&_EVM2EVMTollOffRampHelper.TransactOpts, to)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.TransferOwnership(&_EVM2EVMTollOffRampHelper.TransactOpts, to)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Transmit(&_EVM2EVMTollOffRampHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Transmit(&_EVM2EVMTollOffRampHelper.TransactOpts, reportContext, report, rs, ss, rawVs)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.contract.Transact(opts, "unpause")
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Unpause(&_EVM2EVMTollOffRampHelper.TransactOpts)
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperTransactorSession) Unpause() (*types.Transaction, error) {
	return _EVM2EVMTollOffRampHelper.Contract.Unpause(&_EVM2EVMTollOffRampHelper.TransactOpts)
}

type EVM2EVMTollOffRampHelperAFNSetIterator struct {
	Event *EVM2EVMTollOffRampHelperAFNSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperAFNSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperAFNSet)
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
		it.Event = new(EVM2EVMTollOffRampHelperAFNSet)
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

func (it *EVM2EVMTollOffRampHelperAFNSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperAFNSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperAFNSet struct {
	OldAFN common.Address
	NewAFN common.Address
	Raw    types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperAFNSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperAFNSetIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "AFNSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperAFNSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "AFNSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperAFNSet)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "AFNSet", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseAFNSet(log types.Log) (*EVM2EVMTollOffRampHelperAFNSet, error) {
	event := new(EVM2EVMTollOffRampHelperAFNSet)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "AFNSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperConfigChangedIterator struct {
	Event *EVM2EVMTollOffRampHelperConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperConfigChanged)
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
		it.Event = new(EVM2EVMTollOffRampHelperConfigChanged)
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

func (it *EVM2EVMTollOffRampHelperConfigChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperConfigChanged struct {
	Capacity *big.Int
	Rate     *big.Int
	Raw      types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperConfigChangedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperConfigChangedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperConfigChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperConfigChanged)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseConfigChanged(log types.Log) (*EVM2EVMTollOffRampHelperConfigChanged, error) {
	event := new(EVM2EVMTollOffRampHelperConfigChanged)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperConfigSetIterator struct {
	Event *EVM2EVMTollOffRampHelperConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperConfigSet)
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
		it.Event = new(EVM2EVMTollOffRampHelperConfigSet)
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

func (it *EVM2EVMTollOffRampHelperConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperConfigSet struct {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperConfigSetIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperConfigSet)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseConfigSet(log types.Log) (*EVM2EVMTollOffRampHelperConfigSet, error) {
	event := new(EVM2EVMTollOffRampHelperConfigSet)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperExecutionStateChangedIterator struct {
	Event *EVM2EVMTollOffRampHelperExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperExecutionStateChanged)
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
		it.Event = new(EVM2EVMTollOffRampHelperExecutionStateChanged)
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

func (it *EVM2EVMTollOffRampHelperExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperExecutionStateChanged struct {
	SequenceNumber uint64
	State          uint8
	Raw            types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMTollOffRampHelperExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperExecutionStateChangedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperExecutionStateChanged)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseExecutionStateChanged(log types.Log) (*EVM2EVMTollOffRampHelperExecutionStateChanged, error) {
	event := new(EVM2EVMTollOffRampHelperExecutionStateChanged)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperOffRampConfigSetIterator struct {
	Event *EVM2EVMTollOffRampHelperOffRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperOffRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperOffRampConfigSet)
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
		it.Event = new(EVM2EVMTollOffRampHelperOffRampConfigSet)
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

func (it *EVM2EVMTollOffRampHelperOffRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperOffRampConfigSet struct {
	Config BaseOffRampInterfaceOffRampConfig
	Raw    types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperOffRampConfigSetIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperOffRampConfigSetIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "OffRampConfigSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "OffRampConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperOffRampConfigSet)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseOffRampConfigSet(log types.Log) (*EVM2EVMTollOffRampHelperOffRampConfigSet, error) {
	event := new(EVM2EVMTollOffRampHelperOffRampConfigSet)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OffRampConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperOffRampRouterSetIterator struct {
	Event *EVM2EVMTollOffRampHelperOffRampRouterSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperOffRampRouterSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperOffRampRouterSet)
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
		it.Event = new(EVM2EVMTollOffRampHelperOffRampRouterSet)
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

func (it *EVM2EVMTollOffRampHelperOffRampRouterSetIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperOffRampRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperOffRampRouterSet struct {
	Router common.Address
	Raw    types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMTollOffRampHelperOffRampRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperOffRampRouterSetIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "OffRampRouterSet", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOffRampRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "OffRampRouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperOffRampRouterSet)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseOffRampRouterSet(log types.Log) (*EVM2EVMTollOffRampHelperOffRampRouterSet, error) {
	event := new(EVM2EVMTollOffRampHelperOffRampRouterSet)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OffRampRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator struct {
	Event *EVM2EVMTollOffRampHelperOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperOwnershipTransferRequested)
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
		it.Event = new(EVM2EVMTollOffRampHelperOwnershipTransferRequested)
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

func (it *EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperOwnershipTransferRequested)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOffRampHelperOwnershipTransferRequested, error) {
	event := new(EVM2EVMTollOffRampHelperOwnershipTransferRequested)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperOwnershipTransferredIterator struct {
	Event *EVM2EVMTollOffRampHelperOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperOwnershipTransferred)
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
		it.Event = new(EVM2EVMTollOffRampHelperOwnershipTransferred)
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

func (it *EVM2EVMTollOffRampHelperOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampHelperOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperOwnershipTransferredIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperOwnershipTransferred)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOffRampHelperOwnershipTransferred, error) {
	event := new(EVM2EVMTollOffRampHelperOwnershipTransferred)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperPausedIterator struct {
	Event *EVM2EVMTollOffRampHelperPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperPaused)
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
		it.Event = new(EVM2EVMTollOffRampHelperPaused)
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

func (it *EVM2EVMTollOffRampHelperPausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperPausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperPausedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperPaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperPaused)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParsePaused(log types.Log) (*EVM2EVMTollOffRampHelperPaused, error) {
	event := new(EVM2EVMTollOffRampHelperPaused)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperPoolAddedIterator struct {
	Event *EVM2EVMTollOffRampHelperPoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperPoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperPoolAdded)
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
		it.Event = new(EVM2EVMTollOffRampHelperPoolAdded)
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

func (it *EVM2EVMTollOffRampHelperPoolAddedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperPoolAdded struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperPoolAddedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperPoolAddedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperPoolAdded) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperPoolAdded)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParsePoolAdded(log types.Log) (*EVM2EVMTollOffRampHelperPoolAdded, error) {
	event := new(EVM2EVMTollOffRampHelperPoolAdded)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperPoolRemovedIterator struct {
	Event *EVM2EVMTollOffRampHelperPoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperPoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperPoolRemoved)
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
		it.Event = new(EVM2EVMTollOffRampHelperPoolRemoved)
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

func (it *EVM2EVMTollOffRampHelperPoolRemovedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperPoolRemoved struct {
	Token common.Address
	Pool  common.Address
	Raw   types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperPoolRemovedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperPoolRemovedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperPoolRemoved) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "PoolRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperPoolRemoved)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParsePoolRemoved(log types.Log) (*EVM2EVMTollOffRampHelperPoolRemoved, error) {
	event := new(EVM2EVMTollOffRampHelperPoolRemoved)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperTokenPriceChangedIterator struct {
	Event *EVM2EVMTollOffRampHelperTokenPriceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperTokenPriceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperTokenPriceChanged)
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
		it.Event = new(EVM2EVMTollOffRampHelperTokenPriceChanged)
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

func (it *EVM2EVMTollOffRampHelperTokenPriceChangedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperTokenPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperTokenPriceChanged struct {
	Token    common.Address
	NewPrice *big.Int
	Raw      types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperTokenPriceChangedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperTokenPriceChangedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "TokenPriceChanged", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperTokenPriceChanged) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "TokenPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperTokenPriceChanged)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseTokenPriceChanged(log types.Log) (*EVM2EVMTollOffRampHelperTokenPriceChanged, error) {
	event := new(EVM2EVMTollOffRampHelperTokenPriceChanged)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "TokenPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator struct {
	Event *EVM2EVMTollOffRampHelperTokensRemovedFromBucket

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperTokensRemovedFromBucket)
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
		it.Event = new(EVM2EVMTollOffRampHelperTokensRemovedFromBucket)
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

func (it *EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperTokensRemovedFromBucket struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "TokensRemovedFromBucket", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperTokensRemovedFromBucket) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "TokensRemovedFromBucket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperTokensRemovedFromBucket)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMTollOffRampHelperTokensRemovedFromBucket, error) {
	event := new(EVM2EVMTollOffRampHelperTokensRemovedFromBucket)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "TokensRemovedFromBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperTransmittedIterator struct {
	Event *EVM2EVMTollOffRampHelperTransmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperTransmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperTransmitted)
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
		it.Event = new(EVM2EVMTollOffRampHelperTransmitted)
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

func (it *EVM2EVMTollOffRampHelperTransmittedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperTransmittedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperTransmittedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperTransmitted) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperTransmitted)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseTransmitted(log types.Log) (*EVM2EVMTollOffRampHelperTransmitted, error) {
	event := new(EVM2EVMTollOffRampHelperTransmitted)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EVM2EVMTollOffRampHelperUnpausedIterator struct {
	Event *EVM2EVMTollOffRampHelperUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EVM2EVMTollOffRampHelperUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EVM2EVMTollOffRampHelperUnpaused)
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
		it.Event = new(EVM2EVMTollOffRampHelperUnpaused)
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

func (it *EVM2EVMTollOffRampHelperUnpausedIterator) Error() error {
	return it.fail
}

func (it *EVM2EVMTollOffRampHelperUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EVM2EVMTollOffRampHelperUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperUnpausedIterator, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EVM2EVMTollOffRampHelperUnpausedIterator{contract: _EVM2EVMTollOffRampHelper.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperUnpaused) (event.Subscription, error) {

	logs, sub, err := _EVM2EVMTollOffRampHelper.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EVM2EVMTollOffRampHelperUnpaused)
				if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelperFilterer) ParseUnpaused(log types.Log) (*EVM2EVMTollOffRampHelperUnpaused, error) {
	event := new(EVM2EVMTollOffRampHelperUnpaused)
	if err := _EVM2EVMTollOffRampHelper.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetChainIDs struct {
	SourceChainId *big.Int
	ChainId       *big.Int
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

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelper) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EVM2EVMTollOffRampHelper.abi.Events["AFNSet"].ID:
		return _EVM2EVMTollOffRampHelper.ParseAFNSet(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["ConfigChanged"].ID:
		return _EVM2EVMTollOffRampHelper.ParseConfigChanged(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["ConfigSet"].ID:
		return _EVM2EVMTollOffRampHelper.ParseConfigSet(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["ExecutionStateChanged"].ID:
		return _EVM2EVMTollOffRampHelper.ParseExecutionStateChanged(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["OffRampConfigSet"].ID:
		return _EVM2EVMTollOffRampHelper.ParseOffRampConfigSet(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["OffRampRouterSet"].ID:
		return _EVM2EVMTollOffRampHelper.ParseOffRampRouterSet(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["OwnershipTransferRequested"].ID:
		return _EVM2EVMTollOffRampHelper.ParseOwnershipTransferRequested(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["OwnershipTransferred"].ID:
		return _EVM2EVMTollOffRampHelper.ParseOwnershipTransferred(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["Paused"].ID:
		return _EVM2EVMTollOffRampHelper.ParsePaused(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["PoolAdded"].ID:
		return _EVM2EVMTollOffRampHelper.ParsePoolAdded(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["PoolRemoved"].ID:
		return _EVM2EVMTollOffRampHelper.ParsePoolRemoved(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["TokenPriceChanged"].ID:
		return _EVM2EVMTollOffRampHelper.ParseTokenPriceChanged(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["TokensRemovedFromBucket"].ID:
		return _EVM2EVMTollOffRampHelper.ParseTokensRemovedFromBucket(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["Transmitted"].ID:
		return _EVM2EVMTollOffRampHelper.ParseTransmitted(log)
	case _EVM2EVMTollOffRampHelper.abi.Events["Unpaused"].ID:
		return _EVM2EVMTollOffRampHelper.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (EVM2EVMTollOffRampHelperAFNSet) Topic() common.Hash {
	return common.HexToHash("0x2378f30feefb413d2caee0417ec344de95ab13977e41d6ce944d0a6d2d25bd28")
}

func (EVM2EVMTollOffRampHelperConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x8e012bd57e8109fb3513158da3ff482a86a1e3ff4d5be099be0945772547322d")
}

func (EVM2EVMTollOffRampHelperConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05")
}

func (EVM2EVMTollOffRampHelperExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x06d3f6de62d3b2a5b9679b586cacbb22580c79a7b682eabcd33b523ba208cfbf")
}

func (EVM2EVMTollOffRampHelperOffRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0xd0c6ebb2f4418da3c627fe6eefe153ea976e958e9f835284c55548a37fcf2aa0")
}

func (EVM2EVMTollOffRampHelperOffRampRouterSet) Topic() common.Hash {
	return common.HexToHash("0x993172116697b267c3e4c0884a97c58c6d6df4ff9f97c142ba57101a1e1ed4d4")
}

func (EVM2EVMTollOffRampHelperOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EVM2EVMTollOffRampHelperOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EVM2EVMTollOffRampHelperPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EVM2EVMTollOffRampHelperPoolAdded) Topic() common.Hash {
	return common.HexToHash("0x95f865c2808f8b2a85eea2611db7843150ee7835ef1403f9755918a97d76933c")
}

func (EVM2EVMTollOffRampHelperPoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x987eb3c2f78454541205f72f34839b434c306c9eaf4922efd7c0c3060fdb2e4c")
}

func (EVM2EVMTollOffRampHelperTokenPriceChanged) Topic() common.Hash {
	return common.HexToHash("0x4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
}

func (EVM2EVMTollOffRampHelperTokensRemovedFromBucket) Topic() common.Hash {
	return common.HexToHash("0xcecaabdf078137e9f3ffad598f679665628d62e269c3d929bd10fef8a22ba378")
}

func (EVM2EVMTollOffRampHelperTransmitted) Topic() common.Hash {
	return common.HexToHash("0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62")
}

func (EVM2EVMTollOffRampHelperUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EVM2EVMTollOffRampHelper *EVM2EVMTollOffRampHelper) Address() common.Address {
	return _EVM2EVMTollOffRampHelper.address
}

type EVM2EVMTollOffRampHelperInterface interface {
	CalculateCurrentTokenBucketState(opts *bind.CallOpts) (AggregateRateLimiterInterfaceTokenBucket, error)

	CcipReceive(opts *bind.CallOpts, arg0 CCIPAny2EVMMessageFromSender) error

	FeeTaken(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error)

	GetAFN(opts *bind.CallOpts) (common.Address, error)

	GetBlobVerifier(opts *bind.CallOpts) (common.Address, error)

	GetChainIDs(opts *bind.CallOpts) (GetChainIDs,

		error)

	GetConfig(opts *bind.CallOpts) (BaseOffRampInterfaceOffRampConfig, error)

	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetExecutionState(opts *bind.CallOpts, sequenceNumber uint64) (uint8, error)

	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)

	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)

	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetTokenLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	IsAFNHealthy(opts *bind.CallOpts) (bool, error)

	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails,

		error)

	LatestConfigDigestAndEpoch(opts *bind.CallOpts) (LatestConfigDigestAndEpoch,

		error)

	MetadataHash(opts *bind.CallOpts) ([32]byte, error)

	OverheadGasToll(opts *bind.CallOpts, merkleGasShare *big.Int, message CCIPEVM2EVMTollMessage) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	Transmitters(opts *bind.CallOpts) ([]common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, report CCIPExecutionReport, manualExecution bool) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message CCIPAny2EVMMessageFromSender) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)

	Report(opts *bind.TransactOpts, executableMessages []byte) (*types.Transaction, error)

	SetAFN(opts *bind.TransactOpts, afn common.Address) (*types.Transaction, error)

	SetBlobVerifier(opts *bind.TransactOpts, blobVerifier common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, config BaseOffRampInterfaceOffRampConfig) (*types.Transaction, error)

	SetConfig0(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error)

	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)

	SetRateLimiterConfig(opts *bind.TransactOpts, config AggregateRateLimiterInterfaceRateLimiterConfig) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, router common.Address) (*types.Transaction, error)

	SetTokenLimitAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAFNSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperAFNSetIterator, error)

	WatchAFNSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperAFNSet) (event.Subscription, error)

	ParseAFNSet(log types.Log) (*EVM2EVMTollOffRampHelperAFNSet, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*EVM2EVMTollOffRampHelperConfigChanged, error)

	FilterConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*EVM2EVMTollOffRampHelperConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64) (*EVM2EVMTollOffRampHelperExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperExecutionStateChanged, sequenceNumber []uint64) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*EVM2EVMTollOffRampHelperExecutionStateChanged, error)

	FilterOffRampConfigSet(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperOffRampConfigSetIterator, error)

	WatchOffRampConfigSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOffRampConfigSet) (event.Subscription, error)

	ParseOffRampConfigSet(log types.Log) (*EVM2EVMTollOffRampHelperOffRampConfigSet, error)

	FilterOffRampRouterSet(opts *bind.FilterOpts, router []common.Address) (*EVM2EVMTollOffRampHelperOffRampRouterSetIterator, error)

	WatchOffRampRouterSet(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOffRampRouterSet, router []common.Address) (event.Subscription, error)

	ParseOffRampRouterSet(log types.Log) (*EVM2EVMTollOffRampHelperOffRampRouterSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampHelperOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EVM2EVMTollOffRampHelperOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EVM2EVMTollOffRampHelperOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EVM2EVMTollOffRampHelperOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EVM2EVMTollOffRampHelperPaused, error)

	FilterPoolAdded(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperPoolAddedIterator, error)

	WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperPoolAdded) (event.Subscription, error)

	ParsePoolAdded(log types.Log) (*EVM2EVMTollOffRampHelperPoolAdded, error)

	FilterPoolRemoved(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperPoolRemovedIterator, error)

	WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperPoolRemoved) (event.Subscription, error)

	ParsePoolRemoved(log types.Log) (*EVM2EVMTollOffRampHelperPoolRemoved, error)

	FilterTokenPriceChanged(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperTokenPriceChangedIterator, error)

	WatchTokenPriceChanged(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperTokenPriceChanged) (event.Subscription, error)

	ParseTokenPriceChanged(log types.Log) (*EVM2EVMTollOffRampHelperTokenPriceChanged, error)

	FilterTokensRemovedFromBucket(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperTokensRemovedFromBucketIterator, error)

	WatchTokensRemovedFromBucket(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperTokensRemovedFromBucket) (event.Subscription, error)

	ParseTokensRemovedFromBucket(log types.Log) (*EVM2EVMTollOffRampHelperTokensRemovedFromBucket, error)

	FilterTransmitted(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperTransmittedIterator, error)

	WatchTransmitted(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperTransmitted) (event.Subscription, error)

	ParseTransmitted(log types.Log) (*EVM2EVMTollOffRampHelperTransmitted, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EVM2EVMTollOffRampHelperUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EVM2EVMTollOffRampHelperUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EVM2EVMTollOffRampHelperUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
